package replication

import (
	"errors"
)

type Config struct {
	Peers             []uint64
	ViewChangeTimeout uint64
	HeartbeatTimeout  uint64
	MessageBus        MessageBus
}

type Status uint16

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

type OperationEntry struct {
	OperationNumber uint64
	Operation       []byte
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

	OpLog       MemoryLog[OperationEntry] // The log of operations
	ClientTable map[uint64]ClientEntry

	HeartbeatTick  uint64
	ViewChangeTick uint64
}

func (v *VSRState) Leader() uint64 {
	return v.Configuration.Peers[0]
}

func (v *VSRState) OnPropose(ClientID uint64, seq uint64, op []byte) (opn uint64, err error) {
	// Assert: v.Status == Status_Normal
	if v.Status != Status_Normal {
		err = ErrStatusNotNormal
		return
	}

	// Assert: v.ViewNumber == v.StableView
	if v.ViewNumber != v.StableView {
		err = ErrStatusNotNormal
		return
	}

	// Assert: v.Leader() == v.NodeID
	if v.Leader() != v.NodeID {
		err = ErrNotALeader
		return
		// TODO: redirect to leader
	}

	// Check if client exists
	if _, ok := v.ClientTable[ClientID]; !ok {
		v.ClientTable[ClientID] = ClientEntry{0, CacheEntry{0, 0}}
	}

	client := v.ClientTable[ClientID]

	// Assert: seq > client.SequenceNumber
	if seq <= client.SequenceNumber {
		err = ErrReplay
		return
	}

	// Update client sequence number
	client.SequenceNumber = seq
	v.ClientTable[ClientID] = client

	// Increment operation number
	v.OperationNumber++
	// Assign operation number to operation
	opn = v.OperationNumber

	// Append operation to log
	v.OpLog.Append(OperationEntry{opn, op}, opn)

	// Send operation to all replicas
	// TODO: send to all replicas
	return
}

func (v *VSRState) OnPrepare(src uint64, view uint64, op OperationEntry) (err error) {
	// Assert: v.Status == Status_Normal
	if v.Status != Status_Normal {
		err = ErrStatusNotNormal
		return
	}

	// Assert: v.ViewNumber == v.StableView
	if v.ViewNumber != v.StableView {
		err = ErrStatusNotNormal
		return
	}

	// Assert: view == v.ViewNumber
	if view != v.ViewNumber {
		err = ErrInvalidView
		return
	}

	// Assert: src == v.Leader()
	if src != v.Leader() {
		err = ErrInvalidView
		return
	}

	// Check if all the earlier operations exist in the log
	last, err := v.OpLog.LastIndex()
	if err != nil {
		return
	}

	if last >= op.OperationNumber {
		// Already have this operation
		return
	}

	return
}
