package __obf_60cfa5c95ac1adcc

import (
	"sync"
)

// type Work func()

type WorkPool struct {
	__obf_a4ca110b11a5428e chan struct{}
	__obf_b29c031fcc75637d *sync.WaitGroup
}

func NewPool(__obf_a5f019e44e055a4e int, __obf_b29c031fcc75637d *sync.WaitGroup) *WorkPool {
	return &WorkPool{make(chan struct{}, __obf_a5f019e44e055a4e), __obf_b29c031fcc75637d}
}

func (__obf_f9cc47107db288dd *WorkPool) Acquire() {
	defer __obf_f9cc47107db288dd.__obf_b29c031fcc75637d.Add(1)
	__obf_f9cc47107db288dd.__obf_a4ca110b11a5428e <- struct{}{}
}

func (__obf_f9cc47107db288dd *WorkPool) Release() {
	defer __obf_f9cc47107db288dd.__obf_b29c031fcc75637d.Done()
	<-__obf_f9cc47107db288dd.__obf_a4ca110b11a5428e
}

func (__obf_f9cc47107db288dd *WorkPool) Wait() {
	defer close(__obf_f9cc47107db288dd.__obf_a4ca110b11a5428e)
	__obf_f9cc47107db288dd.__obf_b29c031fcc75637d.Wait()
}
