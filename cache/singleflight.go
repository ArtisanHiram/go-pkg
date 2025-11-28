package __obf_32d927a1e06b7e71

import "sync"

// call is an in-flight or completed Do call
type __obf_10f94ce554e86430 struct {
	__obf_d3c273cd2f08c66e sync.WaitGroup
	__obf_a967d4ebf1531f4c any
	__obf_19fb5c4c25ff452a error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_3123ffce698833d2 sync.Mutex                         // protects m
	__obf_549c2dd367f7ff13 map[string]*__obf_10f94ce554e86430 // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_da812586baf1c809 *Group) Do(__obf_13113b328a6972ad string, __obf_9ac055d43ce03c4f func() (any, error)) (any, error) {
	__obf_da812586baf1c809.__obf_3123ffce698833d2.Lock()
	if __obf_da812586baf1c809.__obf_549c2dd367f7ff13 == nil {
		__obf_da812586baf1c809.__obf_549c2dd367f7ff13 = make(map[string]*__obf_10f94ce554e86430)
	}
	if __obf_443bb096afb82210, __obf_542c0401f22c9aec := __obf_da812586baf1c809.__obf_549c2dd367f7ff13[__obf_13113b328a6972ad]; __obf_542c0401f22c9aec {
		__obf_da812586baf1c809.__obf_3123ffce698833d2.Unlock()
		__obf_443bb096afb82210.__obf_d3c273cd2f08c66e.Wait()
		return __obf_443bb096afb82210.__obf_a967d4ebf1531f4c, __obf_443bb096afb82210.__obf_19fb5c4c25ff452a
	}
	__obf_443bb096afb82210 := new(__obf_10f94ce554e86430)
	__obf_443bb096afb82210.__obf_d3c273cd2f08c66e.Add(1)
	__obf_da812586baf1c809.__obf_549c2dd367f7ff13[__obf_13113b328a6972ad] = __obf_443bb096afb82210
	__obf_da812586baf1c809.__obf_3123ffce698833d2.Unlock()

	__obf_443bb096afb82210.__obf_a967d4ebf1531f4c, __obf_443bb096afb82210.__obf_19fb5c4c25ff452a = __obf_9ac055d43ce03c4f()
	__obf_443bb096afb82210.__obf_d3c273cd2f08c66e.Done()

	__obf_da812586baf1c809.__obf_3123ffce698833d2.Lock()
	delete(__obf_da812586baf1c809.__obf_549c2dd367f7ff13, __obf_13113b328a6972ad)
	__obf_da812586baf1c809.__obf_3123ffce698833d2.Unlock()

	return __obf_443bb096afb82210.__obf_a967d4ebf1531f4c, __obf_443bb096afb82210.__obf_19fb5c4c25ff452a
}
