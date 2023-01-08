package replication

type ring[T any] struct {
	data []T
	r, w uint64
}

// Len returns the number of elements in the ring.
func (r *ring[T]) Len() uint64 {
	return r.w - r.r
}

// Write writes data to the ring. If the ring is full, ok is false.
func (r *ring[T]) Write(data T) (ok bool) {
	if r.w-r.r >= uint64(len(r.data)) {
		return false
	}
	r.data[r.w%uint64(len(r.data))] = data
	r.w++
	return true
}

// Read reads the next element from the ring. If the ring is empty, ok is false.
func (r *ring[T]) Read() (data T, ok bool) {
	if r.w == r.r {
		return
	}
	data = r.data[r.r%uint64(len(r.data))]
	r.r++
	return data, true
}

// Reset resets the ring to empty.
func (r *ring[T]) Reset() {
	r.r, r.w = 0, 0
}

// Cap returns the capacity of the ring.
func (r *ring[T]) Cap() uint64 {
	return uint64(len(r.data))
}

// Free returns the number of elements that can be written to the ring.
func (r *ring[T]) Free() uint64 {
	return r.Cap() - r.Len()
}

// At returns the element at index i, where i is relative to the read index.
func (r *ring[T]) At(i uint64) T {
	return r.data[(r.r+i)%r.Cap()]
}

// Drop drops the first n elements from the ring.
func (r *ring[T]) Drop(n uint64) {
	r.r += n
}
