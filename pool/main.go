package __obf_445fe3b2702a2779

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_a16e8b25cc5e7f79 chan struct{}
	__obf_16b79d5ee2c58f10 *sync.WaitGroup
}

func NewPool(__obf_0374cc8b62d44737 int, __obf_16b79d5ee2c58f10 *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_0374cc8b62d44737), __obf_16b79d5ee2c58f10}
}

func (__obf_4e4963664aefd48c *WorkPool) Acquire() {
	defer __obf_4e4963664aefd48c.__obf_16b79d5ee2c58f10.Add(1)
	__obf_4e4963664aefd48c.__obf_a16e8b25cc5e7f79 <- struct{}{}
}

func (__obf_4e4963664aefd48c *WorkPool) Release() {
	defer __obf_4e4963664aefd48c.__obf_16b79d5ee2c58f10.Done()
	<-__obf_4e4963664aefd48c.__obf_a16e8b25cc5e7f79
}

func (__obf_4e4963664aefd48c *WorkPool) Wait() {
	defer close(__obf_4e4963664aefd48c.__obf_a16e8b25cc5e7f79)
	__obf_4e4963664aefd48c.__obf_16b79d5ee2c58f10.
		Wait()
}
