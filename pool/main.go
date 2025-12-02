package __obf_e6ddedfdf75d4ad4

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_a4f8dafd4aac040d chan struct{}
	__obf_3e2c27920b82c6eb *sync.WaitGroup
}

func NewPool(__obf_82f851b1c0d3bb28 int, __obf_3e2c27920b82c6eb *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_82f851b1c0d3bb28), __obf_3e2c27920b82c6eb}
}

func (__obf_8518e1f60174e441 *WorkPool) Acquire() {
	defer __obf_8518e1f60174e441.__obf_3e2c27920b82c6eb.Add(1)
	__obf_8518e1f60174e441.__obf_a4f8dafd4aac040d <- struct{}{}
}

func (__obf_8518e1f60174e441 *WorkPool) Release() {
	defer __obf_8518e1f60174e441.__obf_3e2c27920b82c6eb.Done()
	<-__obf_8518e1f60174e441.__obf_a4f8dafd4aac040d
}

func (__obf_8518e1f60174e441 *WorkPool) Wait() {
	defer close(__obf_8518e1f60174e441.__obf_a4f8dafd4aac040d)
	__obf_8518e1f60174e441.__obf_3e2c27920b82c6eb.
		Wait()
}
