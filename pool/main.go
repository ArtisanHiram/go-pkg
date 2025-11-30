package __obf_5ee0ed5311c3f8e6

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_31c1875c16c1811b chan struct{}
	__obf_15ae9a7c79af352e *sync.WaitGroup
}

func NewPool(__obf_d5d37558f41bbf8f int, __obf_15ae9a7c79af352e *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_d5d37558f41bbf8f), __obf_15ae9a7c79af352e}
}

func (__obf_a241047fea0d393a *WorkPool) Acquire() {
	defer __obf_a241047fea0d393a.__obf_15ae9a7c79af352e.Add(1)
	__obf_a241047fea0d393a.__obf_31c1875c16c1811b <- struct{}{}
}

func (__obf_a241047fea0d393a *WorkPool) Release() {
	defer __obf_a241047fea0d393a.__obf_15ae9a7c79af352e.Done()
	<-__obf_a241047fea0d393a.__obf_31c1875c16c1811b
}

func (__obf_a241047fea0d393a *WorkPool) Wait() {
	defer close(__obf_a241047fea0d393a.__obf_31c1875c16c1811b)
	__obf_a241047fea0d393a.__obf_15ae9a7c79af352e.
		Wait()
}
