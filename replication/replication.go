package replication

import (
	"sync"

	"v8.run/go/supersystem/replication/vsrproto"
)

type Node struct {
	ID      uint64
	Address string
}

type Config struct {
	ViewChangeTimeout uint64
	HeartbeatTimeout  uint64
}

type Status uint16

const (
	Status_Normal Status = 1 << iota
	Status_ViewChange
	Status_Recovering
)

type ReplicationGroup struct {
	GroupID uint64
	Config  *Config

	Clock    Clock
	Lock     sync.Mutex
	Loopback []*vsrproto.Message

	Nodes           []Node
	OperationNumber uint64
	ViewNumber      uint64
	CommitNumber    uint64
	ClientTable     []uint64
	ClientTimeout   []uint64
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
