package __obf_f36428c4eaad9d20

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_d13ba12356730919 chan struct{}
	__obf_e2a6b3c150ee6356 *sync.WaitGroup
}

func NewPool(__obf_90b81aad1e98b8e8 int, __obf_e2a6b3c150ee6356 *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_90b81aad1e98b8e8), __obf_e2a6b3c150ee6356}
}

func (__obf_37fbd405a22e6324 *WorkPool) Acquire() {
	defer __obf_37fbd405a22e6324.__obf_e2a6b3c150ee6356.Add(1)
	__obf_37fbd405a22e6324.__obf_d13ba12356730919 <- struct{}{}
}

func (__obf_37fbd405a22e6324 *WorkPool) Release() {
	defer __obf_37fbd405a22e6324.__obf_e2a6b3c150ee6356.Done()
	<-__obf_37fbd405a22e6324.__obf_d13ba12356730919
}

func (__obf_37fbd405a22e6324 *WorkPool) Wait() {
	defer close(__obf_37fbd405a22e6324.__obf_d13ba12356730919)
	__obf_37fbd405a22e6324.__obf_e2a6b3c150ee6356.
		Wait()
}
