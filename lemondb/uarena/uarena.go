package uarena

type UArena struct {
	Data   []byte
	offset uint64
	ref    uint64
}

func NewUArena(size uint64) *UArena {
	return &UArena{
		Data:   make([]byte, size),
		offset: 0,
		ref:    0,
	}
}

const (
	InvalidOffset = uint64(1<<64 - 1)
)

func (ua *UArena) Alloc(size uint64) (offset uint64) {
	if ua.offset+size > uint64(len(ua.Data)) {
		return InvalidOffset
	}
	offset = ua.offset
	ua.offset += size
	return
}

func (ua *UArena) Ref() uint64 {
	return ua.ref
}

func (ua *UArena) IncRef() {
	ua.ref++
}

func (ua *UArena) DecRef() (free bool) {
	ua.ref--
	return ua.ref == 0
}
