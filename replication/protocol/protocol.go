package protocol

//go:generate stringer -type=MessageType
//go:generate gobe .

type MessageType uint8

const (
	MT_Unknown MessageType = iota
)

type Message struct {
	Type    MessageType
	Payload []byte
}
