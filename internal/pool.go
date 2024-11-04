package internal

import (
	"fmt"
)

type Pool[GoCallback any, CCallback comparable] struct {
	pool        map[int]CCallback
	allocations map[int]GoCallback
}

func NewPool[GoCallback any, CCallback comparable](poolElements ...CCallback) *Pool[GoCallback, CCallback] {
	result := &Pool[GoCallback, CCallback]{
		pool:        make(map[int]CCallback),
		allocations: make(map[int]GoCallback),
	}

	for i := range len(poolElements) {
		result.pool[i] = poolElements[i]
	}

	return result
}

func (p *Pool[GoCallback, CCallback]) Get(id int) GoCallback {
	return p.allocations[id]
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
There are several possible solutions:
- You forgot to call Clear() for the pool (e.g. on the end/beginning of the frame)
- You've too many items calling Allocate without cleaning (more than %d).
  I this is the case, plelase report an issue ane we'll increase pool limit for the type.

Refer: https://github.com/AllenDang/cimgui-go/issues/224
`, callback, len(p.pool)))
}

func (p *Pool[GoCallback, CCallback]) Find(callback CCallback) GoCallback {
	for k, v := range p.pool {
		// if reflect.ValueOf(v).Pointer() == reflect.ValueOf(callback).Pointer() {
		if v == callback {
			return p.allocations[k]
		}
	}

	panic(fmt.Sprintf(`cimgui-go/internal.Pool: Find: Not Broken or Incorrectly Allocated C callback!

This means that you attempt to Find source function of type %T, but the function given as an argument to Find was either
allocated in a different way (not by this Pool) or was deallocated somehow.
Refer: https://github.com/AllenDang/cimgui-go/issues/224
`, callback))
}

func (p *Pool[GoCallback, CCallback]) Clear() {
	p.allocations = make(map[int]GoCallback)
}
