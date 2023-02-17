package uwal

import (
	"encoding/binary"
	"errors"
	"os"
	"time"

	"github.com/lemon-mint/super-system/lemondb/uwal/uwalproto"
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
// |  Log Type (1B)  | Entry Type (1B) | Key Data                          |
// +-----------------+-----------------+-----------------+-----------------+
// |  Value Data                                                           |
// +-----------------+-----------------+-----------------+-----------------+
//

const terminatorSize = 4 + 4 + 4 + 1 + 1
const entryHeaserSize = terminatorSize

type LogType uint8

const (
	LogStart LogType = iota
	LogMark
	LogCommit
	LogAbort
	LogTerminate
)

type EntryType uint8

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
	ErrTooLargeEntry     = errors.New("too large entry")
	ErrReadOnly          = errors.New("read only")
)

func NewUWAL(f *os.File, id uint32) (*UWAL, error) {
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

type Entry struct {
	Key       []byte
	Value     []byte
	LogType   LogType
	EntryType EntryType
}

func (w *UWAL) WriteEntry(e *Entry) error {
	if w.mode != ModeWrite {
		return ErrReadOnly
	}

	if len(e.Key) > 0xFFFFFFFF || len(e.Value) > 0xFFFFFFFF ||
		entryHeaserSize+len(e.Key)+len(e.Value)+terminatorSize > UWAL_BLOCK_SIZE {
		return ErrTooLargeEntry
	}

	// Write the entry
	size := entryHeaserSize + uint64(len(e.Key)) + uint64(len(e.Value))
	if w.offset+size+terminatorSize > UWAL_BLOCK_SIZE {
		err := w.FlushBlock()
		if err != nil {
			return err
		}
	}

	_ = w.buffer[w.offset+entryHeaserSize+uint64(len(e.Key))+uint64(len(e.Value))] // bounds check hint to compiler
	_ = w.buffer[w.offset+13]                                                      // bounds check hint to compiler
	binary.LittleEndian.PutUint32(w.buffer[w.offset:], uint32(len(e.Key)))         // key size
	binary.LittleEndian.PutUint32(w.buffer[w.offset+4:], uint32(len(e.Value)))     // value size
	binary.LittleEndian.PutUint32(w.buffer[w.offset+8:], w.header.FileID)          // log id
	w.buffer[w.offset+12] = byte(e.LogType)                                        // log type
	w.buffer[w.offset+13] = byte(e.EntryType)                                      // entry type
	copy(w.buffer[w.offset+entryHeaserSize:], e.Key)                               // key data
	copy(w.buffer[w.offset+entryHeaserSize+uint64(len(e.Key)):], e.Value)          // value data
	w.offset += size

	return nil
}

func (w *UWAL) FlushBlock() error {
	if w.mode != ModeWrite {
		return ErrReadOnly
	}

	if w.offset > 0 {
		if w.offset+terminatorSize > UWAL_BLOCK_SIZE {
			return ErrTooLargeEntry
		}

		// Write the terminator
		binary.LittleEndian.PutUint32(w.buffer[w.offset:], 0)                 // key size
		binary.LittleEndian.PutUint32(w.buffer[w.offset+4:], 0)               // value size
		binary.LittleEndian.PutUint32(w.buffer[w.offset+8:], w.header.FileID) // log id
		w.buffer[w.offset+12] = byte(LogTerminate)                            // log type
		w.buffer[w.offset+13] = 0                                             // entry type
		w.offset += terminatorSize

		zero := w.buffer[w.offset:]
		for i := range zero {
			zero[i] = 0
		}

		// Write the block
		_, err := w.f.Write(w.buffer)
		if err != nil {
			return err
		}
		w.offset = 0
	}

	return nil
}

func (w *UWAL) Sync() error {
	return w.f.Sync()
}
