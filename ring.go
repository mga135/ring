// Package ring
// Head() and Tail() are undefined when ring is size 0
package ring

import "errors"

var ErrInvalidCapacity = errors.New("capacity must be greater than 0")

type Ring[T any] struct {
	data []T
	idx  int
	size int
	cap  int
}

func New[T any](capacity int) (*Ring[T], error) {
	if capacity < 1 {
		return nil, ErrInvalidCapacity
	}

	return &Ring[T]{
		data: make([]T, capacity),
		cap:  capacity,
	}, nil
}

func (r *Ring[T]) Add(v T) {
	r.data[r.idx] = v
	r.idx = (r.idx + 1) % r.cap

	if r.size < r.cap {
		r.size++
	}
}

func (r *Ring[T]) Size() int {
	return r.size
}

func (r *Ring[T]) Cap() int {
	return r.cap
}

func (r *Ring[T]) Head() T {
	return r.data[(r.idx-r.size+r.cap)%r.cap]
}

func (r *Ring[T]) Tail() T {
	return r.data[(r.idx-1+r.cap)%r.cap]
}

func (r *Ring[T]) Data() []T {
	data := make([]T, r.size)
	start := (r.idx - r.size + r.cap) % r.cap

	for i := range r.size {
		data[i] = r.data[(start+i)%r.cap]
	}

	return data
}
