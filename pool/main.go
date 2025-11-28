package __obf_3a682104ea13968a

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_b0ef2be2ecfdadf1 chan struct{}
	__obf_e92d17632258f5d6 *sync.WaitGroup
}

func NewPool(__obf_fb7e7db6b27f1057 int, __obf_e92d17632258f5d6 *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_fb7e7db6b27f1057), __obf_e92d17632258f5d6}
}

func (__obf_bef39b9dde02e27e *WorkPool) Acquire() {
	defer __obf_bef39b9dde02e27e.__obf_e92d17632258f5d6.Add(1)
	__obf_bef39b9dde02e27e.__obf_b0ef2be2ecfdadf1 <- struct{}{}
}

func (__obf_bef39b9dde02e27e *WorkPool) Release() {
	defer __obf_bef39b9dde02e27e.__obf_e92d17632258f5d6.Done()
	<-__obf_bef39b9dde02e27e.__obf_b0ef2be2ecfdadf1
}

func (__obf_bef39b9dde02e27e *WorkPool) Wait() {
	defer close(__obf_bef39b9dde02e27e.__obf_b0ef2be2ecfdadf1)
	__obf_bef39b9dde02e27e.__obf_e92d17632258f5d6.Wait()
}
