package replication

import (
	"errors"
	"sync"

	"v8.run/go/supersystem/replication/vsrproto"
)

type Node struct {
	ID      uint64
	Address string
}

type Config struct {
	NodeID            uint64
	ViewChangeTimeout uint64
	HeartbeatTimeout  uint64
}

type Status uint16

const (
	Status_Normal Status = 1 << iota
	Status_ViewChange
	Status_Recovering
)

type ClientT struct {
	SequenceNumber uint64
	Last           *vsrproto.Message
}

type CommitT struct {
	Index uint64
	N     uint64
}

type CommitTable struct {
	Table            []CommitT
	CommitNumber_MAX uint64
	CommitNumber_MIN uint64
	Offset           uint64
}

func (ct *CommitTable) Add(index uint64, n uint64) {
	ct.Table = append(ct.Table, CommitT{index, n})
}

func (ct *CommitTable) update(quorum uint64) {
	// Update CommitNumber_MIN
	for m := ct.CommitNumber_MIN + 1; m <= ct.CommitNumber_MAX; m++ {
		if uint64(len(ct.Table)) <= m-ct.Offset {
			break
		}

		if ct.Table[m-ct.Offset].Index != m {
			panic("unreachable")
		}

		if ct.Table[m-ct.Offset].N < quorum {
			break
		}
		ct.CommitNumber_MIN = m
	}
}

func (ct *CommitTable) Commit(index uint64, quorum uint64) bool {
	if uint64(len(ct.Table)) <= index-ct.Offset {
		return false
	}

	if ct.Table[index-ct.Offset].Index != index {
		panic("unreachable")
	}
	ct.Table[index-ct.Offset].N++

	if ct.Table[index-ct.Offset].N >= quorum {
		if index > ct.CommitNumber_MAX {
			ct.CommitNumber_MAX = index
		}
		ct.update(quorum)
		return true
	}

	return false
}

type Transport interface {
	Send(msg *vsrproto.Message, dst uint64) error
	Broadcast(msg *vsrproto.Message) error
}

type ReplicationGroup struct {
	GroupID uint64
	Config  *Config

	Log          LogStore[[]byte]
	Clock        Clock
	Lock         sync.Mutex
	Loopback     []*vsrproto.Message
	MessageQueue []*vsrproto.Message

	Transport Transport

	Nodes           []Node
	Status          Status
	OperationNumber uint64
	ViewNumber      uint64

	CommitTable CommitTable
	ClientTable map[uint64]*ClientT

	HeartbeatDeadline  uint64
	ViewChangeDeadline uint64
}

func (rg *ReplicationGroup) Quorum() uint64 {
	return uint64(len(rg.Nodes)/2 + 1)
}

func (rg *ReplicationGroup) Leader() uint64 {
	return rg.Nodes[rg.ViewNumber%uint64(len(rg.Nodes))].ID
}

func (rg *ReplicationGroup) Tick() error {
	rg.Lock.Lock()
	defer rg.Lock.Unlock()

	rg.Clock.Tick()
	switch rg.Status {
	case Status_Normal:
		if rg.Clock.NowTicks() >= rg.HeartbeatDeadline {
			rg.Status = Status_ViewChange
			rg.ViewChangeDeadline = rg.Clock.NowTicks() + rg.Config.ViewChangeTimeout
			// TODO: Broadcast StartViewChange
		}
	case Status_ViewChange:
		if rg.Clock.NowTicks() >= rg.ViewChangeDeadline {
			rg.ViewNumber++
			rg.ViewChangeDeadline = rg.Clock.NowTicks() + rg.Config.ViewChangeTimeout
		}
	}

	return nil
}

var (
	ErrInvalidMessage = errors.New("invalid message")
	ErrInvalidClient  = errors.New("invalid client")
)

func (rg *ReplicationGroup) processMessage(msg *vsrproto.Message) error {
	switch msg.Type {
	case vsrproto.MessageType_MPropose:
		if rg.Leader() != rg.Config.NodeID {
			return nil
		}

		propose := msg.Propose
		if propose == nil {
			return ErrInvalidMessage
		}

		client, ok := rg.ClientTable[propose.ClientID]
		if !ok {
			return ErrInvalidClient
		}

		if client.SequenceNumber >= propose.SequenceNumber {
			if client.SequenceNumber == propose.SequenceNumber &&
				client.Last != nil {
				rg.Transport.Send(client.Last, propose.ClientID)
				return nil
			}
			duplicate := &vsrproto.Message{
				GroupID: rg.GroupID,
				Source:  rg.Config.NodeID,
				Type:    vsrproto.MessageType_MProposeReject,
				ProposeReject: &vsrproto.ProposeReject{
					ClientID:       propose.ClientID,
					SequenceNumber: propose.SequenceNumber,
					Error:          vsrproto.Error_EDupPropose,
				},
			}
			rg.Transport.Send(duplicate, propose.ClientID)
			return nil
		}

		rg.OperationNumber++
		n := rg.OperationNumber

		prepare := &vsrproto.Message{
			GroupID: rg.GroupID,
			Source:  rg.Config.NodeID,

			Type: vsrproto.MessageType_MPrepare,
			Prepare: &vsrproto.Prepare{
				ViewNumber:       rg.ViewNumber,
				OperationNumber:  n,
				Propose:          propose,
				CommitNumber_MIN: rg.CommitTable.CommitNumber_MIN,
				CommitNumber_MAX: rg.CommitTable.CommitNumber_MAX,
			},
		}

		err := rg.Log.Append(propose.Operation, n)
		if err != nil {
			return err
		}
		client.SequenceNumber = propose.SequenceNumber

		rg.Loopback = append(rg.Loopback, prepare)
		rg.Transport.Broadcast(prepare)
	case vsrproto.MessageType_MPrepare:
		if rg.Status != Status_Normal {
			// Ignore prepares during view change or recovery
			return nil
		}

		prepare := msg.Prepare
		if prepare == nil {
			return ErrInvalidMessage
		}

		if prepare.ViewNumber != rg.ViewNumber {
			// Ignore prepares from other views
			return nil
		}

		if prepare.OperationNumber <= rg.CommitTable.CommitNumber_MIN {
			// Ignore prepares for already committed operations
			return nil
		}

		err := rg.Log.Append(prepare.Propose.Operation, prepare.OperationNumber)
		if err != nil {
			// TODO: use state transfer to recover missing operations
			return err
		}

		prepareOK := &vsrproto.Message{
			GroupID: rg.GroupID,
			Source:  rg.Config.NodeID,
			Type:    vsrproto.MessageType_MPrepareOK,
			PrepareOK: &vsrproto.PrepareOK{
				ViewNumber:      rg.ViewNumber,
				OperationNumber: prepare.OperationNumber,
			},
		}

		if rg.Leader() == rg.Config.NodeID {
			rg.Loopback = append(rg.Loopback, prepareOK)
		} else {
			rg.Transport.Send(prepareOK, rg.Leader())
		}

		// Reset the heartbeat deadline
		rg.HeartbeatDeadline = rg.Clock.NowTicks() + rg.Config.HeartbeatTimeout
	case vsrproto.MessageType_MPrepareOK:
		if rg.Status != Status_Normal {
			// Ignore prepares during view change or recovery
			return nil
		}

		if rg.Leader() != rg.Config.NodeID {
			return nil
		}

		prepareOK := msg.PrepareOK
		if prepareOK == nil {
			return ErrInvalidMessage
		}

		if prepareOK.ViewNumber != rg.ViewNumber {
			// Ignore prepares from other views
			return nil
		}

		rg.CommitTable.Commit(prepareOK.OperationNumber, uint64(rg.Quorum()))
	case vsrproto.MessageType_MCommit:
		if rg.Status != Status_Normal {
			// Ignore commits during view change or recovery
			return nil
		}

		commit := msg.Commit
		if commit == nil {
			return ErrInvalidMessage
		}

		if commit.ViewNumber != rg.ViewNumber {
			// Ignore commits from other views
			return nil
		}

		if commit.CommitNumber_MIN <= rg.CommitTable.CommitNumber_MIN {
			// Ignore commits for already committed operations
			return nil
		}

		last, err := rg.Log.LastIndex()
		if err != nil {
			return err
		}

		if last < commit.CommitNumber_MIN {
			// TODO: use state transfer to recover missing operations
			return nil
		}

		if rg.Leader() != rg.Config.NodeID {
			rg.CommitTable.CommitNumber_MAX = commit.CommitNumber_MAX
			rg.CommitTable.CommitNumber_MIN = commit.CommitNumber_MIN
			return nil
		}
	}

	return nil
}
