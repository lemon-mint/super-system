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
	return v.Configuration.Peers[0]
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
	msg.Prepare.ViewNumber = v.ViewNumber
	msg.Prepare.OperationNumber = opn
	msg.Prepare.CommitNumber = v.CommitNumberMAX
	msg.Prepare.Operation = m.Propose.Operation

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

func (v *VSRState) OnPrepare(src uint64, view uint64, ope protocol.OperationEntry) (e errno.Errno) {
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
	if view != v.ViewNumber {
		e = errno.ERRNO_VIEWMISMATCH
		return
	}

	// Assert: src == v.Leader()
	if src != v.Leader() {
		e = errno.ERRNO_NOTLEADER
		return
	}

	// Check if all the earlier operations exist in the log
	// TODO: Handle zero-length log
	last, err := v.OpLog.LastIndex()
	if err != nil {
		return
	}

	if last >= ope.OperationNumber {
		// Already have this operation
		return
	}

	// Append operation to log
	v.OpLog.Append(ope, ope.OperationNumber)

	msg := protocol.AcquireMessage()
	msg.GroupID = v.Configuration.GroupID
	msg.Type = protocol.MT_PrepareAcceptance
	msg.PrepareAcceptance.ViewNumber = v.ViewNumber
	msg.PrepareAcceptance.OperationNumber = ope.OperationNumber
	v.Configuration.MessageBus.SendMessageToReplica(src, msg)

	return
}
