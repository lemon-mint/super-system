package protocol

type MessageType uint8

const (
	MT_Unknown MessageType = iota
)

type Message struct {
	Type    MessageType
	Payload []byte
}
