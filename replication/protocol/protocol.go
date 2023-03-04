package protocol

import (
	"sync"

	"github.com/lemon-mint/super-system/replication/errno"
)

//go:generate stringer -type=MessageType
//go:generate gobe .

type MessageType uint8

type Status uint16

const (
	MT_Unknown MessageType = iota
	MT_Propose
	MT_Prepare
	MT_ProposeRejection
	MT_PrepareRejection
	MT_PrepareAcceptance
)

type Propose struct {
	ClientID       uint64
	SequenceNumber uint64
	Operation      []byte
}

type ProposeRejection struct {
	ClientID       uint64
	SequenceNumber uint64
	Reason         errno.Errno
}

type OperationEntry struct {
	ViewNumber      uint64
	OperationNumber uint64
	Operation       []byte
}

type Prepare struct {
	OperationEntry
	CommitNumber uint64
}

type PrepareRejection struct {
	ViewNumber      uint64
	OperationNumber uint64
	Reason          errno.Errno
}

type PrepareAcceptance struct {
	ViewNumber      uint64
	OperationNumber uint64
}

type Message struct {
	Source            uint64
	GroupID           uint64
	Type              MessageType
	Propose           `gobe_enum:"Type=MT_Propose"`
	Prepare           `gobe_enum:"Type=MT_Prepare"`
	PrepareAcceptance `gobe_enum:"Type=MT_PrepareAcceptance"`
	ProposeRejection  `gobe_enum:"Type=MT_ProposeRejection"`
	PrepareRejection  `gobe_enum:"Type=MT_PrepareRejection"`
}

var _MessagePool = sync.Pool{
	New: func() interface{} {
		return &Message{}
	},
}

func AcquireMessage() *Message {
	return _MessagePool.Get().(*Message)
}

func ReleaseMessage(m *Message) {
	*m = Message{} // Clear the message
	_MessagePool.Put(m)
}

func (m *Message) Free() {
	ReleaseMessage(m)
}
