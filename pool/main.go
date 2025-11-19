package __obf_3bba98d9f9f99c2a

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_fcf8fddf0e0de2f5 chan struct{}
	__obf_4fea3b30a1cf9f7a *sync.WaitGroup
}

func NewPool(__obf_6063ce7c774292f4 int, __obf_4fea3b30a1cf9f7a *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_6063ce7c774292f4), __obf_4fea3b30a1cf9f7a}
}

func (__obf_7338ae62ae267e06 *WorkPool) Acquire() {
	defer __obf_7338ae62ae267e06.__obf_4fea3b30a1cf9f7a.Add(1)
	__obf_7338ae62ae267e06.__obf_fcf8fddf0e0de2f5 <- struct{}{}
}

func (__obf_7338ae62ae267e06 *WorkPool) Release() {
	defer __obf_7338ae62ae267e06.__obf_4fea3b30a1cf9f7a.Done()
	<-__obf_7338ae62ae267e06.__obf_fcf8fddf0e0de2f5
}

func (__obf_7338ae62ae267e06 *WorkPool) Wait() {
	defer close(__obf_7338ae62ae267e06.__obf_fcf8fddf0e0de2f5)
	__obf_7338ae62ae267e06.__obf_4fea3b30a1cf9f7a.Wait()
}
