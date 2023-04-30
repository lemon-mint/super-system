package replication

import (
	"io"

	"github.com/lemon-mint/super-system/replication/protocol"
)

// StateMachine API is based on the github.com/hashicorp/raft library

type SnapshotSink interface {
	io.WriteCloser

	SnapshotID() string
	Cancel() error
}

type Snapshot interface {
	Save(w SnapshotSink) error
	Release()
}

type StateMachine interface {
	ApplyOperations(logs []protocol.OperationEntry) error

	Snapshot() (Snapshot, error)
	Restore(r io.ReadCloser) error
}
