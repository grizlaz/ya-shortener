package pool

import (
	"sync"
)

type Resetter interface {
	Reset()
}

type Pool[T Resetter] struct {
	mu    sync.Mutex
	items []T
	New   func() T
}

func New[T Resetter]() *Pool[T] {
	return &Pool[T]{}
}

func (p *Pool[T]) Get() T {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.items) > 0 {
		last := len(p.items) - 1
		v := p.items[last]
		p.items = p.items[:last]
		return v
	}

	if p.New != nil {
		return p.New()
	}

	var v T
	return v
}

func (p *Pool[T]) Put(v T) {
	v.Reset()

	p.mu.Lock()
	defer p.mu.Unlock()
	p.items = append(p.items, v)
}
