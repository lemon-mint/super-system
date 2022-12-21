package replication

type LogEntry struct {
}

type ReplicationGroup struct {
	GroupID    uint64
	ViewNumber uint64
	Log        []LogEntry
}
