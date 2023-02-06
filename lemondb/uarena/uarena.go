package uarena

type UArena struct {
	Data   []byte
	offset uint32
}

func NewUArena(size uint64) *UArena {
	return &UArena{
		Data:   make([]byte, size),
		offset: 0,
	}
}

const (
	InvalidOffset = uint32(1<<32 - 1)
)

func (ua *UArena) Alloc(size uint32) (offset uint32) {
	if ua.offset+size > uint32(len(ua.Data)) {
		return InvalidOffset
	}
	offset = ua.offset
	ua.offset += size
	return
}

func (ua *UArena) Bytes(offset uint32, size uint32) []byte {
	return ua.Data[offset : offset+size]
}
