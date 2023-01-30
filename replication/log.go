package replication

import (
	"errors"
)

type MemoryLog[T any] struct {
	ring   ring[T]
	offset uint64
}

func NewMemoryLog[T any](size uint64) *MemoryLog[T] {
	return &MemoryLog[T]{
		ring: ring[T]{
			data: make([]T, size),
		},
	}
}

var (
	ErrInvalidIndex = errors.New("LogStore[T]: Invalid index")
	ErrRingFull     = errors.New("LogStore[T]: Ring is full")
)

func (m *MemoryLog[T]) Append(msg T, id uint64) error {
	// Check if the id is the next one
	if id != m.offset+m.ring.Len() {
		return ErrInvalidIndex
	}
	if !m.ring.Write(msg) {
		return ErrRingFull
	}
	return nil
}

func (m *MemoryLog[T]) Get(index uint64) (T, error) {
	var zero T
	if index < m.offset {
		return zero, ErrInvalidIndex
	}
	if index >= m.offset+m.ring.Len() {
		return zero, ErrInvalidIndex
	}
	return m.ring.At(index - m.offset), nil
}

func (m *MemoryLog[T]) LastIndex() (uint64, error) {
	if m.ring.Len() == 0 {
		return 0, ErrInvalidIndex
	}
	return m.offset + m.ring.Len() - 1, nil
}

func (m *MemoryLog[T]) FirstIndex() (uint64, error) {
	return m.offset, nil
}

func (m *MemoryLog[T]) Truncate(index uint64) error {
	if index < m.offset {
		return ErrInvalidIndex
	}

	if index >= m.offset+m.ring.Len() {
		return ErrInvalidIndex
	}

	m.ring.Drop(index - m.offset)
	m.offset = index

	return nil
}

func (m *MemoryLog[T]) Len() uint64 {
	return m.ring.Len()
}

func (ls *MemoryLog[T]) UnsafeSetOffset(index uint64) {
	ls.offset = index
}
