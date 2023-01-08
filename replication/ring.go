package replication

type ring[T any] struct {
	data []T
	r, w uint64
}

func newring[T any](size uint64) *ring[T] {
	return &ring[T]{data: make([]T, size)}
}

func (r *ring[T]) Len() uint64 {
	return r.w - r.r
}

func (r *ring[T]) Write(data T) (ok bool) {
	if r.w-r.r >= uint64(len(r.data)) {
		return false
	}
	r.data[r.w%uint64(len(r.data))] = data
	r.w++
	return true
}

func (r *ring[T]) Read() (data T, ok bool) {
	if r.w == r.r {
		return
	}
	data = r.data[r.r%uint64(len(r.data))]
	r.r++
	return data, true
}

func (r *ring[T]) Reset() {
	r.r, r.w = 0, 0
}

func (r *ring[T]) Cap() uint64 {
	return uint64(len(r.data))
}

func (r *ring[T]) Free() uint64 {
	return r.Cap() - r.Len()
}
