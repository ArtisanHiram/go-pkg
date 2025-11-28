package __obf_b9a96f7b2f3b936d

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_a09f34ec301bd9df chan struct{}
	__obf_b5013a4935858e0e *sync.WaitGroup
}

func NewPool(__obf_9a0ed1ec851808f9 int, __obf_b5013a4935858e0e *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_9a0ed1ec851808f9), __obf_b5013a4935858e0e}
}

func (__obf_6f8d8e2a1ddc3814 *WorkPool) Acquire() {
	defer __obf_6f8d8e2a1ddc3814.__obf_b5013a4935858e0e.Add(1)
	__obf_6f8d8e2a1ddc3814.__obf_a09f34ec301bd9df <- struct{}{}
}

func (__obf_6f8d8e2a1ddc3814 *WorkPool) Release() {
	defer __obf_6f8d8e2a1ddc3814.__obf_b5013a4935858e0e.Done()
	<-__obf_6f8d8e2a1ddc3814.__obf_a09f34ec301bd9df
}

func (__obf_6f8d8e2a1ddc3814 *WorkPool) Wait() {
	defer close(__obf_6f8d8e2a1ddc3814.__obf_a09f34ec301bd9df)
	__obf_6f8d8e2a1ddc3814.__obf_b5013a4935858e0e.Wait()
}
