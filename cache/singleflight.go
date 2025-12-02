package __obf_6b306490bf7221d3

import "sync"

// call is an in-flight or completed Do call
type __obf_1a8df4824b3def6d struct {
	__obf_6274ab57ec38f0eb sync.WaitGroup
	__obf_17a4af14df8bae4f any
	__obf_9ffdd5bbb9f1dc61 error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_2aacabd3f44c4b79          sync.Mutex
	__obf_2a1b11e5bfb41cce          map// protects m
	[string]*__obf_1a8df4824b3def6d // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_bc0f0089338403ce *Group) Do(__obf_fa3a380c35ada5d9 string, __obf_b615d054acd2440e func() (any, error)) (any, error) {
	__obf_bc0f0089338403ce.__obf_2aacabd3f44c4b79.
		Lock()
	if __obf_bc0f0089338403ce.__obf_2a1b11e5bfb41cce == nil {
		__obf_bc0f0089338403ce.__obf_2a1b11e5bfb41cce = make(map[string]*__obf_1a8df4824b3def6d)
	}
	if __obf_6ccac4bb8646a806, __obf_8761a3e94b6603ac := __obf_bc0f0089338403ce.__obf_2a1b11e5bfb41cce[__obf_fa3a380c35ada5d9]; __obf_8761a3e94b6603ac {
		__obf_bc0f0089338403ce.__obf_2aacabd3f44c4b79.
			Unlock()
		__obf_6ccac4bb8646a806.__obf_6274ab57ec38f0eb.
			Wait()
		return __obf_6ccac4bb8646a806.__obf_17a4af14df8bae4f, __obf_6ccac4bb8646a806.__obf_9ffdd5bbb9f1dc61
	}
	__obf_6ccac4bb8646a806 := new(__obf_1a8df4824b3def6d)
	__obf_6ccac4bb8646a806.__obf_6274ab57ec38f0eb.
		Add(1)
	__obf_bc0f0089338403ce.__obf_2a1b11e5bfb41cce[__obf_fa3a380c35ada5d9] = __obf_6ccac4bb8646a806
	__obf_bc0f0089338403ce.__obf_2aacabd3f44c4b79.
		Unlock()
	__obf_6ccac4bb8646a806.__obf_17a4af14df8bae4f, __obf_6ccac4bb8646a806.__obf_9ffdd5bbb9f1dc61 = __obf_b615d054acd2440e()
	__obf_6ccac4bb8646a806.__obf_6274ab57ec38f0eb.
		Done()
	__obf_bc0f0089338403ce.__obf_2aacabd3f44c4b79.
		Lock()
	delete(__obf_bc0f0089338403ce.__obf_2a1b11e5bfb41cce, __obf_fa3a380c35ada5d9)
	__obf_bc0f0089338403ce.__obf_2aacabd3f44c4b79.
		Unlock()

	return __obf_6ccac4bb8646a806.__obf_17a4af14df8bae4f, __obf_6ccac4bb8646a806.__obf_9ffdd5bbb9f1dc61
}
