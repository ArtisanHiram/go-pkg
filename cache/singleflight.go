package __obf_79e7502d8563d901

import "sync"

// call is an in-flight or completed Do call
type __obf_10fc309f91e0076a struct {
	__obf_52aa963c71306d0c sync.WaitGroup
	__obf_e043381e0771feca any
	__obf_8d46af2525fab46a error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_e1ec702455188bd2 sync.Mutex                         // protects m
	__obf_bd80a351910d9377 map[string]*__obf_10fc309f91e0076a // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_7c78f0e75bb54bc8 *Group) Do(__obf_50994613b7653a88 string, __obf_a8d39957f3282147 func() (any, error)) (any, error) {
	__obf_7c78f0e75bb54bc8.__obf_e1ec702455188bd2.Lock()
	if __obf_7c78f0e75bb54bc8.__obf_bd80a351910d9377 == nil {
		__obf_7c78f0e75bb54bc8.__obf_bd80a351910d9377 = make(map[string]*__obf_10fc309f91e0076a)
	}
	if __obf_db09942aea00505a, __obf_7500628aeb1f47ab := __obf_7c78f0e75bb54bc8.__obf_bd80a351910d9377[__obf_50994613b7653a88]; __obf_7500628aeb1f47ab {
		__obf_7c78f0e75bb54bc8.__obf_e1ec702455188bd2.Unlock()
		__obf_db09942aea00505a.__obf_52aa963c71306d0c.Wait()
		return __obf_db09942aea00505a.__obf_e043381e0771feca, __obf_db09942aea00505a.__obf_8d46af2525fab46a
	}
	__obf_db09942aea00505a := new(__obf_10fc309f91e0076a)
	__obf_db09942aea00505a.__obf_52aa963c71306d0c.Add(1)
	__obf_7c78f0e75bb54bc8.__obf_bd80a351910d9377[__obf_50994613b7653a88] = __obf_db09942aea00505a
	__obf_7c78f0e75bb54bc8.__obf_e1ec702455188bd2.Unlock()

	__obf_db09942aea00505a.__obf_e043381e0771feca, __obf_db09942aea00505a.__obf_8d46af2525fab46a = __obf_a8d39957f3282147()
	__obf_db09942aea00505a.__obf_52aa963c71306d0c.Done()

	__obf_7c78f0e75bb54bc8.__obf_e1ec702455188bd2.Lock()
	delete(__obf_7c78f0e75bb54bc8.__obf_bd80a351910d9377, __obf_50994613b7653a88)
	__obf_7c78f0e75bb54bc8.__obf_e1ec702455188bd2.Unlock()

	return __obf_db09942aea00505a.__obf_e043381e0771feca, __obf_db09942aea00505a.__obf_8d46af2525fab46a
}
