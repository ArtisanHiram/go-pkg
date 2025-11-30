package __obf_9e1ee87c6b054458

import "sync"

// call is an in-flight or completed Do call
type __obf_d2b44bf28ebe38c4 struct {
	__obf_4df297942177b9a3 sync.WaitGroup
	__obf_7d0a67e130b6e9b4 any
	__obf_498673050542660c error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_26007347eed65c2a          sync.Mutex
	__obf_2c6bba2eafff0a17          map// protects m
	[string]*__obf_d2b44bf28ebe38c4 // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_7a48d4f4afee7674 *Group) Do(__obf_3c13197612c6b39f string, __obf_6d9f33b69b402b6e func() (any, error)) (any, error) {
	__obf_7a48d4f4afee7674.__obf_26007347eed65c2a.
		Lock()
	if __obf_7a48d4f4afee7674.__obf_2c6bba2eafff0a17 == nil {
		__obf_7a48d4f4afee7674.__obf_2c6bba2eafff0a17 = make(map[string]*__obf_d2b44bf28ebe38c4)
	}
	if __obf_8af2676446f64c49, __obf_40eed10588aa6cec := __obf_7a48d4f4afee7674.__obf_2c6bba2eafff0a17[__obf_3c13197612c6b39f]; __obf_40eed10588aa6cec {
		__obf_7a48d4f4afee7674.__obf_26007347eed65c2a.
			Unlock()
		__obf_8af2676446f64c49.__obf_4df297942177b9a3.
			Wait()
		return __obf_8af2676446f64c49.__obf_7d0a67e130b6e9b4, __obf_8af2676446f64c49.__obf_498673050542660c
	}
	__obf_8af2676446f64c49 := new(__obf_d2b44bf28ebe38c4)
	__obf_8af2676446f64c49.__obf_4df297942177b9a3.
		Add(1)
	__obf_7a48d4f4afee7674.__obf_2c6bba2eafff0a17[__obf_3c13197612c6b39f] = __obf_8af2676446f64c49
	__obf_7a48d4f4afee7674.__obf_26007347eed65c2a.
		Unlock()
	__obf_8af2676446f64c49.__obf_7d0a67e130b6e9b4, __obf_8af2676446f64c49.__obf_498673050542660c = __obf_6d9f33b69b402b6e()
	__obf_8af2676446f64c49.__obf_4df297942177b9a3.
		Done()
	__obf_7a48d4f4afee7674.__obf_26007347eed65c2a.
		Lock()
	delete(__obf_7a48d4f4afee7674.__obf_2c6bba2eafff0a17, __obf_3c13197612c6b39f)
	__obf_7a48d4f4afee7674.__obf_26007347eed65c2a.
		Unlock()

	return __obf_8af2676446f64c49.__obf_7d0a67e130b6e9b4, __obf_8af2676446f64c49.__obf_498673050542660c
}
