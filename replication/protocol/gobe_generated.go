package protocol

func (ns25519 *Message) SizeGOBE() uint64 {
	var ns25520 uint64

	// ZZ: (struct{Type ..MessageType; Payload []byte})(ns25519)

	// ZZ: (..MessageType)(ns25519.Type)

	// ZZ: (uint8)(ns25519.Type)
	ns25520 += 1

	// ZZ: ([]byte)(ns25519.Payload)
	ns25520 += 8 + uint64(len(ns25519.Payload))

	return ns25520
}

func (ns25521 *Message) MarshalGOBE(dst []byte) uint64 {
	var ns25522 uint64

	// ZZ: (struct{Type ..MessageType; Payload []byte})(ns25521)

	// ZZ: (..MessageType)(ns25521.Type)

	// ZZ: (uint8)(ns25521.Type)
	dst[ns25522+0] = byte(ns25521.Type >> 0)
	ns25522++

	// ZZ: ([]byte)(ns25521.Payload)
	var ns25523 uint64 = uint64(len(ns25521.Payload))
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25523 >> 0)
	dst[ns25522+1] = byte(ns25523 >> 8)
	dst[ns25522+2] = byte(ns25523 >> 16)
	dst[ns25522+3] = byte(ns25523 >> 24)
	dst[ns25522+4] = byte(ns25523 >> 32)
	dst[ns25522+5] = byte(ns25523 >> 40)
	dst[ns25522+6] = byte(ns25523 >> 48)
	dst[ns25522+7] = byte(ns25523 >> 56)
	copy(dst[ns25522+8:], ns25521.Payload)
	ns25522 += 8 + ns25523

	return ns25522
}

func (ns25524 *Message) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Type ..MessageType; Payload []byte})(ns25524)

	// ZZ: (..MessageType)(ns25524.Type)

	// ZZ: (uint8)(ns25524.Type)
	if uint64(len(src)) < offset+1 {
		return
	}
	_ = src[offset+0]
	ns25524.Type = MessageType(
		uint8(src[offset+0]) << 0)
	offset += 1

	// ZZ: ([]byte)(ns25524.Payload)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25525 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25525 {
		return
	}
	ns25524.Payload = src[offset : offset+ns25525]
	offset += ns25525

	ok = true
	return
}

func (ns25526 *MessageType) SizeGOBE() uint64 {
	var ns25527 uint64

	// ZZ: (..MessageType)(*ns25526)

	// ZZ: (uint8)(*ns25526)
	ns25527 += 1

	return ns25527
}

func (ns25528 *MessageType) MarshalGOBE(dst []byte) uint64 {
	var ns25529 uint64

	// ZZ: (..MessageType)(*ns25528)

	// ZZ: (uint8)(*ns25528)
	dst[ns25529+0] = byte(*ns25528 >> 0)
	ns25529++

	return ns25529
}

func (ns25530 *MessageType) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (..MessageType)(*ns25530)

	// ZZ: (uint8)(*ns25530)
	if uint64(len(src)) < offset+1 {
		return
	}
	_ = src[offset+0]
	*ns25530 = MessageType(
		uint8(src[offset+0]) << 0)
	offset += 1

	ok = true
	return
}
