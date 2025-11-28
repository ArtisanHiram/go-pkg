package __obf_9efafe65e1e91d58

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_8c0d97f5a029292b chan struct{}
	__obf_7abec4e6261cf211 *sync.WaitGroup
}

func NewPool(__obf_506223d90d71dde9 int, __obf_7abec4e6261cf211 *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_506223d90d71dde9), __obf_7abec4e6261cf211}
}

func (__obf_8c77e7a0bab4f733 *WorkPool) Acquire() {
	defer __obf_8c77e7a0bab4f733.__obf_7abec4e6261cf211.Add(1)
	__obf_8c77e7a0bab4f733.__obf_8c0d97f5a029292b <- struct{}{}
}

func (__obf_8c77e7a0bab4f733 *WorkPool) Release() {
	defer __obf_8c77e7a0bab4f733.__obf_7abec4e6261cf211.Done()
	<-__obf_8c77e7a0bab4f733.__obf_8c0d97f5a029292b
}

func (__obf_8c77e7a0bab4f733 *WorkPool) Wait() {
	defer close(__obf_8c77e7a0bab4f733.__obf_8c0d97f5a029292b)
	__obf_8c77e7a0bab4f733.__obf_7abec4e6261cf211.Wait()
}
