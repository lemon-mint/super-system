package uwal

import (
	"bufio"
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

type UWAL struct {
	f      *os.File
	header uwalproto.UWALHeader
	size   uint64

	mode Mode
	br   *bufio.Reader
	bw   *bufio.Writer
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
		mode: ModeWrite,
	}

	offset, err := f.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	if offset != 0 {
		return nil, ErrInvalidFileOffset
	}

	// Write the header
	header_size := wal.header.SizeGOBE()
	header_buf := make([]byte, header_size)
	wal.header.MarshalGOBE(header_buf)
	_, err = wal.f.Write(header_buf)
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

	return nil, nil
}
