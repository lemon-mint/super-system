package replication

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
	CommitNumber    uint64

	OpLog       MemoryLog[ClientEntry] // The log of operations
	ClientTable map[uint64]ClientEntry

	HeartbeatTick  uint64
	ViewChangeTick uint64
}
