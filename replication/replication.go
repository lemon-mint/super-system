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

type ReplicationGroup struct {
	GroupID uint64
	Config  *Config

	Clock    Clock
	Lock     sync.Mutex
	Loopback []*vsrproto.Message

	Nodes            []Node
	Status           Status
	OperationNumber  uint64
	ViewNumber       uint64
	CommitNumber_MIN uint64
	CommitNumber_MAX uint64
	ClientTable      map[uint64]*ClientT
}

func (rg *ReplicationGroup) Quorum() int {
	return len(rg.Nodes)/2 + 1
}

func (rg *ReplicationGroup) Tick() error {
	rg.Lock.Lock()
	defer rg.Lock.Unlock()

	rg.Clock.Tick()

	return nil
}

var (
	ErrInvalidMessage = errors.New("invalid message")
	ErrInvalidClient  = errors.New("invalid client")
)

func (rg *ReplicationGroup) processMessage(msg *vsrproto.Message) error {
	switch msg.Type {
	case vsrproto.MessageType_MPropose:
		propose := msg.Propose
		if propose == nil {
			return ErrInvalidMessage
		}
		clientid := propose.ClientID
		if clientid >= uint64(len(rg.ClientTable)) {
			return ErrInvalidClient
		}

		if rg.ClientTable[clientid].SequenceNumber >= propose.SequenceNumber {
			// TODO: Reject duplicate proposals
			return nil
		}
		rg.ClientTable[clientid].SequenceNumber = propose.SequenceNumber

		rg.OperationNumber++
		n := rg.OperationNumber
		prepare := &vsrproto.Message{
			GroupID: rg.GroupID,
			Source:  rg.Config.NodeID,

			Type: vsrproto.MessageType_MPrepare,
			Prepare: &vsrproto.Prepare{
				ViewNumber:      rg.ViewNumber,
				OperationNumber: n,
				Propose:         propose,
				CommitNumber:    rg.CommitNumber_MIN,
			},
		}

		// TODO: Send prepare to all nodes
		rg.Loopback = append(rg.Loopback, prepare)

		// TODO: Wait for quorum of prepares
		// Commit
		rg.CommitNumber_MAX = n
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

		if prepare.OperationNumber <= rg.CommitNumber_MIN {
			// Ignore prepares for already committed operations
			return nil
		}
	}

	return nil
}
