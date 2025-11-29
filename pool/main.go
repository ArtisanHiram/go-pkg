package __obf_b03a5a347fe87e1b

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_f406c578c78b0bfb chan struct{}
	__obf_dae01da1d4d17401 *sync.WaitGroup
}

func NewPool(__obf_8bb41c6a88dcb36c int, __obf_dae01da1d4d17401 *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_8bb41c6a88dcb36c), __obf_dae01da1d4d17401}
}

func (__obf_d6d57eb5ecb2a804 *WorkPool) Acquire() {
	defer __obf_d6d57eb5ecb2a804.__obf_dae01da1d4d17401.Add(1)
	__obf_d6d57eb5ecb2a804.__obf_f406c578c78b0bfb <- struct{}{}
}

func (__obf_d6d57eb5ecb2a804 *WorkPool) Release() {
	defer __obf_d6d57eb5ecb2a804.__obf_dae01da1d4d17401.Done()
	<-__obf_d6d57eb5ecb2a804.__obf_f406c578c78b0bfb
}

func (__obf_d6d57eb5ecb2a804 *WorkPool) Wait() {
	defer close(__obf_d6d57eb5ecb2a804.__obf_f406c578c78b0bfb)
	__obf_d6d57eb5ecb2a804.__obf_dae01da1d4d17401.
		Wait()
}
