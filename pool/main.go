package __obf_12d0e02977a9c646

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_1401f7bbc0823f7d chan struct{}
	__obf_196bb5cec7375d42 *sync.WaitGroup
}

func NewPool(__obf_3625536e8a873901 int, __obf_196bb5cec7375d42 *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_3625536e8a873901), __obf_196bb5cec7375d42}
}

func (__obf_5398bcc0b3bb848f *WorkPool) Acquire() {
	defer __obf_5398bcc0b3bb848f.__obf_196bb5cec7375d42.Add(1)
	__obf_5398bcc0b3bb848f.__obf_1401f7bbc0823f7d <- struct{}{}
}

func (__obf_5398bcc0b3bb848f *WorkPool) Release() {
	defer __obf_5398bcc0b3bb848f.__obf_196bb5cec7375d42.Done()
	<-__obf_5398bcc0b3bb848f.__obf_1401f7bbc0823f7d
}

func (__obf_5398bcc0b3bb848f *WorkPool) Wait() {
	defer close(__obf_5398bcc0b3bb848f.__obf_1401f7bbc0823f7d)
	__obf_5398bcc0b3bb848f.__obf_196bb5cec7375d42.
		Wait()
}
