package internal

import "fmt"

type Pool[GoCallback, CCallback any] struct {
	pool        map[int]CCallback
	allocations map[int]GoCallback
}

func NewPool[GoCallback, CCallback any](poolElements ...CCallback) *Pool[GoCallback, CCallback] {
	result := &Pool[GoCallback, CCallback]{
		pool:        make(map[int]CCallback),
		allocations: make(map[int]GoCallback),
	}

	for i := range len(poolElements) {
		result.pool[i] = poolElements[i]
	}

	return result
}

func (p *Pool[GoCallback, CCallback]) Get(id int) CCallback {
	return p.pool[id]
}

func (p *Pool[GoCallback, CCallback]) Allocate(callback GoCallback) CCallback {
	for k, v := range p.pool {
		if _, ok := p.allocations[k]; !ok {
			p.allocations[k] = callback
			return v
		}
	}

	panic(fmt.Sprintf(`cimgui-go/internal.Pool: Allocate: Out of space!

This means that you attempt to allocate too many callbacks for %T.
Refer: https://github.com/AllenDang/cimgui-go/issues/224
`, callback))
}
