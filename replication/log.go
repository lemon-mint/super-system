package replication

import (
	"errors"
)

type LogStore[T any] struct {
	logs   []T
	offset uint64
}

var (
	ErrInvalidIndex = errors.New("LogStore[T]: Invalid index")
)

func (ls *LogStore[T]) Append(msg T, id uint64) error {
	// Check if the id is the next one
	if id != ls.offset+uint64(len(ls.logs)) {
		return ErrInvalidIndex
	}
	ls.logs = append(ls.logs, msg)
	return nil
}

func (ls *LogStore[T]) Get(index uint64) (T, error) {
	var zero T
	if index < ls.offset {
		return zero, ErrInvalidIndex
	}
	if index >= ls.offset+uint64(len(ls.logs)) {
		return zero, ErrInvalidIndex
	}
	return ls.logs[index-ls.offset], nil
}

func (ls *LogStore[T]) LastIndex() (uint64, error) {
	return ls.offset + uint64(len(ls.logs)) - 1, nil
}

func (ls *LogStore[T]) FirstIndex() (uint64, error) {
	return ls.offset, nil
}

func (ls *LogStore[T]) Truncate(index uint64) error {
	if index < ls.offset {
		return ErrInvalidIndex
	}

	if index >= ls.offset+uint64(len(ls.logs)) {
		return ErrInvalidIndex
	}

	tt := index - ls.offset
	newLen := len(ls.logs) - int(tt)
	copy(ls.logs, ls.logs[tt:])
	ls.logs = ls.logs[:newLen]
	ls.offset = index

	return nil
}

func (ls *LogStore[T]) Len() int {
	return len(ls.logs)
}

func (ls *LogStore[T]) UnsafeSetOffset(index uint64) {
	ls.offset = index
}

func (ls *LogStore[T]) UnsafeSetLogs(log []T) {
	ls.logs = log
}

func (ls *LogStore[T]) UnsafeGetLogs() []T {
	return ls.logs
}
