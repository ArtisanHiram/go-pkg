package __obf_3b8640e918b7e3ff

import "sync"

// call is an in-flight or completed Do call
type __obf_1161b9d88827e444 struct {
	__obf_4158c88150019551 sync.WaitGroup
	__obf_13458218654a7f13 any
	__obf_8c6b39ef87b4061f error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_d8684d60a98e1a8c          sync.Mutex
	__obf_885f88c1b206dbea          map// protects m
	[string]*__obf_1161b9d88827e444 // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_2e59c5f7b481103b *Group) Do(__obf_3405c14f70aaa4d0 string, __obf_c0917fd9c4f84d86 func() (any, error)) (any, error) {
	__obf_2e59c5f7b481103b.__obf_d8684d60a98e1a8c.
		Lock()
	if __obf_2e59c5f7b481103b.__obf_885f88c1b206dbea == nil {
		__obf_2e59c5f7b481103b.__obf_885f88c1b206dbea = make(map[string]*__obf_1161b9d88827e444)
	}
	if __obf_5b354994d2d4bcf1, __obf_cddbfb0aefdf4145 := __obf_2e59c5f7b481103b.__obf_885f88c1b206dbea[__obf_3405c14f70aaa4d0]; __obf_cddbfb0aefdf4145 {
		__obf_2e59c5f7b481103b.__obf_d8684d60a98e1a8c.
			Unlock()
		__obf_5b354994d2d4bcf1.__obf_4158c88150019551.
			Wait()
		return __obf_5b354994d2d4bcf1.__obf_13458218654a7f13, __obf_5b354994d2d4bcf1.__obf_8c6b39ef87b4061f
	}
	__obf_5b354994d2d4bcf1 := new(__obf_1161b9d88827e444)
	__obf_5b354994d2d4bcf1.__obf_4158c88150019551.
		Add(1)
	__obf_2e59c5f7b481103b.__obf_885f88c1b206dbea[__obf_3405c14f70aaa4d0] = __obf_5b354994d2d4bcf1
	__obf_2e59c5f7b481103b.__obf_d8684d60a98e1a8c.
		Unlock()
	__obf_5b354994d2d4bcf1.__obf_13458218654a7f13, __obf_5b354994d2d4bcf1.__obf_8c6b39ef87b4061f = __obf_c0917fd9c4f84d86()
	__obf_5b354994d2d4bcf1.__obf_4158c88150019551.
		Done()
	__obf_2e59c5f7b481103b.__obf_d8684d60a98e1a8c.
		Lock()
	delete(__obf_2e59c5f7b481103b.__obf_885f88c1b206dbea, __obf_3405c14f70aaa4d0)
	__obf_2e59c5f7b481103b.__obf_d8684d60a98e1a8c.
		Unlock()

	return __obf_5b354994d2d4bcf1.__obf_13458218654a7f13, __obf_5b354994d2d4bcf1.__obf_8c6b39ef87b4061f
}
