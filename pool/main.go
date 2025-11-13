package __obf_0f07cec62fc7f6fb

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_8ec7feddc98b0fc5 chan struct{}
	__obf_8aae87faaac9af4e *sync.WaitGroup
}

func NewPool(__obf_4aef4d0e6dab3599 int, __obf_8aae87faaac9af4e *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_4aef4d0e6dab3599), __obf_8aae87faaac9af4e}
}

func (__obf_2aeea9e08af2cce5 *WorkPool) Acquire() {
	defer __obf_2aeea9e08af2cce5.__obf_8aae87faaac9af4e.Add(1)
	__obf_2aeea9e08af2cce5.__obf_8ec7feddc98b0fc5 <- struct{}{}
}

func (__obf_2aeea9e08af2cce5 *WorkPool) Release() {
	defer __obf_2aeea9e08af2cce5.__obf_8aae87faaac9af4e.Done()
	<-__obf_2aeea9e08af2cce5.__obf_8ec7feddc98b0fc5
}

func (__obf_2aeea9e08af2cce5 *WorkPool) Wait() {
	defer close(__obf_2aeea9e08af2cce5.__obf_8ec7feddc98b0fc5)
	__obf_2aeea9e08af2cce5.__obf_8aae87faaac9af4e.Wait()
}
