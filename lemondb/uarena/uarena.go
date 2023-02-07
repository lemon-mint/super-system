package uarena

import "sync/atomic"

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
	OutOfMemory = uint32(1<<32 - 1)
)

func (ua *UArena) Alloc(size uint32) (addr uint32) {
	size = ((size + 7) >> 3) << 3
	addr = atomic.AddUint32(&ua.offset, size)
	if addr >= uint32(len(ua.Data)) {
		return OutOfMemory
	}
	return addr - size
}

// Reset resets the offset of the arena to 0.
// This is useful when you want to reuse the arena.
// This method requires synchronization.
func (ua *UArena) Reset() {
	ua.offset = 0

	// memset(ua.Data, 0, len(ua.Data))
	for i := range ua.Data {
		ua.Data[i] = 0
	}
}
