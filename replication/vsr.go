package replication

import (
	"errors"

	"github.com/lemon-mint/super-system/replication/errno"
	"github.com/lemon-mint/super-system/replication/protocol"
)

type Config struct {
	GroupID           uint64
	Peers             []uint64
	ViewChangeTimeout uint64
	HeartbeatTimeout  uint64
	MessageBus        MessageBus
}

type Status = protocol.Status

const (
	Status_Normal Status = 1 << iota
	Status_ViewChange
	Status_Recovering
)

var (
	ErrStatusNotNormal = errors.New("Not in normal state")
	ErrNotALeader      = errors.New("Not a leader")
	ErrInvalidView     = errors.New("Invalid view")
	ErrReplay          = errors.New("Replay")
)

type ClientEntry struct {
	SequenceNumber uint64
	ResponseCache  CacheEntry
}

type VSRState struct {
	Configuration Config
	NodeID        uint64
	ReplicaNumber uint64

	ViewNumber uint64
	StableView uint64 // StableView is the highest view number that has been stable

	Status Status // Normal, ViewChange, Recovering

	OperationNumber uint64
	CommitNumberMAX uint64
	CommitNumberMIN uint64

	OpLog       MemoryLog[protocol.OperationEntry] // The log of operations
	ClientTable map[uint64]ClientEntry

	HeartbeatTick  uint64
	ViewChangeTick uint64
}

func (v *VSRState) Leader() uint64 {
	return v.Configuration.Peers[v.ViewNumber%uint64(len(v.Configuration.Peers))]
}

func (v *VSRState) OnPropose(m *protocol.Message) (opn uint64, e errno.Errno) {
	if m.Type != protocol.MT_Propose {
		panic("Invalid message type, (unreachable)")
	}

	// Assert: v.Status == Status_Normal
	if v.Status != Status_Normal {
		e = errno.ERRNO_STATUSNOTNORMAL
		return
	}

	// Assert: v.ViewNumber == v.StableView
	if v.ViewNumber != v.StableView {
		e = errno.ERRNO_NOTINSTABLEVIEW
		return
	}

	// Assert: v.Leader() == v.NodeID
	if v.Leader() != v.NodeID {
		e = errno.ERRNO_NOTLEADER
		return
		// TODO: redirect to leader
	}

	// Check if client exists
	if _, ok := v.ClientTable[m.Propose.ClientID]; !ok {
		v.ClientTable[m.Propose.ClientID] = ClientEntry{m.Propose.ClientID, CacheEntry{0, 0}}
	}

	client := v.ClientTable[m.Propose.ClientID]

	// Assert: seq > client.SequenceNumber
	if m.Propose.SequenceNumber <= client.SequenceNumber {
		e = errno.ERRNO_REPLAY
		return
	}

	// Update client sequence number
	client.SequenceNumber = m.Propose.SequenceNumber
	v.ClientTable[m.Propose.ClientID] = client

	// Increment operation number
	v.OperationNumber++
	// Assign operation number to operation
	opn = v.OperationNumber

	msg := protocol.AcquireMessage()
	msg.GroupID = v.Configuration.GroupID
	msg.Type = protocol.MT_Prepare
	msg.Prepare.CommitNumber = v.CommitNumberMAX
	msg.Prepare.OperationEntry.OperationNumber = opn
	msg.Prepare.OperationEntry.Operation = m.Propose.Operation
	msg.Prepare.OperationEntry.ViewNumber = v.ViewNumber

	// Append operation to log
	err := v.OpLog.Append(msg.Prepare.OperationEntry, opn)
	if err != nil {
		// TODO: Handle error
		panic(err)
	}

	// Send operation to all replicas
	v.Configuration.MessageBus.BroadcastMessageToReplicas(msg)

	return
}

func (v *VSRState) OnPrepare(m *protocol.Message) (e errno.Errno) {
	if m.Type != protocol.MT_Prepare {
		panic("Invalid message type, (unreachable)")
	}

	// Assert: v.Status == Status_Normal
	if v.Status != Status_Normal {
		e = errno.ERRNO_STATUSNOTNORMAL
		return
	}

	// Assert: v.ViewNumber == v.StableView
	if v.ViewNumber != v.StableView {
		e = errno.ERRNO_NOTINSTABLEVIEW
		return
	}

	// Assert: view == v.ViewNumber
	if m.Prepare.ViewNumber != v.ViewNumber {
		e = errno.ERRNO_VIEWMISMATCH
		return
	}

	// Assert: src == v.Leader()
	if m.Source != v.Leader() {
		e = errno.ERRNO_NOTLEADER
		return
	}

	// Assert: v.CommitNumberMAX <= m.Prepare.CommitNumber
	if v.CommitNumberMAX > m.Prepare.CommitNumber {
		e = errno.ERRNO_LATECOMMIT
		return
	}

	// Assert: m.Prepare.CommitNumber <= v.OperationNumber
	if m.Prepare.CommitNumber > v.OperationNumber {
		e = errno.ERRNO_EARLYCOMMIT
		return
	}

	// Check if all the earlier operations exist in the log
	// TODO: Handle zero-length log
	last, err := v.OpLog.LastIndex()
	if err != nil {
		return
	}

	if last >= m.Prepare.OperationNumber {
		// Already have this operation
		return
	}

	// Append operation to log
	v.OpLog.Append(m.Prepare.OperationEntry, m.Prepare.OperationEntry.OperationNumber)

	// Commit all operations up to commit number
	v.CommitNumberMAX = m.Prepare.CommitNumber
	v.CommitNumberMIN = m.Prepare.CommitNumber

	msg := protocol.AcquireMessage()
	msg.GroupID = v.Configuration.GroupID
	msg.Type = protocol.MT_PrepareAcceptance
	msg.PrepareAcceptance.ViewNumber = v.ViewNumber
	msg.PrepareAcceptance.OperationNumber = m.Prepare.OperationEntry.OperationNumber
	v.Configuration.MessageBus.SendMessageToReplica(m.Source, msg)

	return
}

func (v *VSRState) OnPrepareAcceptance(m *protocol.Message) (e errno.Errno) {
	if m.Type != protocol.MT_PrepareAcceptance {
		panic("Invalid message type, (unreachable)")
	}

	// Assert: v.Status == Status_Normal
	if v.Status != Status_Normal {
		e = errno.ERRNO_STATUSNOTNORMAL
		return
	}

	// Assert: v.ViewNumber == v.StableView
	if v.ViewNumber != v.StableView {
		e = errno.ERRNO_NOTINSTABLEVIEW
		return
	}

	// Assert: view == v.ViewNumber
	if m.PrepareAcceptance.ViewNumber != v.ViewNumber {
		e = errno.ERRNO_VIEWMISMATCH
		return
	}

	// Assert: v.NodeID == v.Leader()
	if v.NodeID != v.Leader() {
		e = errno.ERRNO_NOTLEADER
		return
	}

	return
}
