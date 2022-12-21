package replication

type Node struct {
	ID      uint64
	Address string
}

type Config struct {
	ViewChangeTimeout uint64
	HeartbeatTimeout  uint64
}

type LogEntry struct {
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

	Nodes           []Node
	OperationNumber uint64
	ViewNumber      uint64
	Log             []LogEntry
	CommitNumber    uint64
	ClientTable     []uint64
}

func (rg *ReplicationGroup) Quorum() int {
	return len(rg.Nodes)/2 + 1
}
