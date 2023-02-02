package uwal

import (
	"errors"
	"os"
	"time"

	"v8.run/go/supersystem/lemondb/uwal/uwalproto"
)

type Mode uint8

const (
	ModeRead Mode = iota
	ModeWrite
)

const UWAL_BLOCK_SIZE = 1 << 18

//
// UWAL File Format
// +-----------------+-----------------+-----------------+-----------------+
// |                            UWAL Header                                |
// +-----------------+-----------------+-----------------+-----------------+
// |                            Block Size Padding                         |
// +-----------------+-----------------+-----------------+-----------------+
// |                            UWAL Block0    						       |
// +-----------------+-----------------+-----------------+-----------------+
// |                            ...                                        |
// +-----------------+-----------------+-----------------+-----------------+
//
// UWAL Block Format
// +-----------------+-----------------+-----------------+-----------------+
// |                            Entry0                                     |
// +-----------------+-----------------+-----------------+-----------------+
// |                            ...                                        |
// +-----------------+-----------------+-----------------+-----------------+
// |                            EntryN                                     |
// +-----------------+-----------------+-----------------+-----------------+
// |                            Block Size Padding                         |
// +-----------------+-----------------+-----------------+-----------------+
//
// UWAL Entry Format
// +-----------------+-----------------+-----------------+-----------------+
// |  Key Size (4B)                                                        |
// +-----------------+-----------------+-----------------+-----------------+
// |  Value Size (4B)                                                      |
// +-----------------+-----------------+-----------------+-----------------+
// |  Log ID (4B)                                                          |
// +-----------------+-----------------+-----------------+-----------------+
// |  Type (1B)      |   Key Data                                          |
// +-----------------+-----------------+-----------------+-----------------+
// |  Value Data                                                           |
// +-----------------+-----------------+-----------------+-----------------+

type UWAL struct {
	f      *os.File
	header uwalproto.UWALHeader
	size   uint64

	mode Mode

	buffer []byte
	offset uint64
}

var (
	ErrInvalidUWALFormat = errors.New("invalid uwal format")
	ErrInvalidFileOffset = errors.New("invalid file offset")
)

func NewUWAL(f *os.File, id uint64) (*UWAL, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if stat.Size() != 0 {
		err = f.Truncate(0)
		if err != nil {
			return nil, err
		}
	}

	wal := &UWAL{
		f: f,
		header: uwalproto.UWALHeader{
			WALMagic: uwalproto.UWAL_MAGIC_V01,
			Version:  uwalproto.UWAL_VERSION_V01,
			FileID:   id,
			UWALTS:   uint64(time.Now().UnixNano()),
		},
		mode:   ModeWrite,
		buffer: make([]byte, UWAL_BLOCK_SIZE),
	}

	offset, err := f.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	if offset != 0 {
		return nil, ErrInvalidFileOffset
	}

	// Write the header
	wal.offset += wal.header.MarshalGOBE(wal.buffer[wal.offset:])
	_, err = wal.f.Write(wal.buffer)
	if err != nil {
		return nil, err
	}

	return wal, nil
}

func (w *UWAL) Close() error {
	return w.f.Close()
}

func OpenUWAL(f *os.File) (*UWAL, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if stat.Size() == 0 {
		return nil, ErrInvalidUWALFormat
	}

	buffer := make([]byte, UWAL_BLOCK_SIZE)
	n, err := f.Read(buffer)
	if err != nil {
		return nil, err
	}

	wal := &UWAL{
		f:      f,
		size:   uint64(stat.Size()),
		mode:   ModeRead,
		buffer: buffer,
	}

	_, ok := wal.header.UnmarshalGOBE(buffer[:n])
	if !ok {
		return nil, ErrInvalidUWALFormat
	}

	return wal, nil
}
