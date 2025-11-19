package __obf_62df4de078c8d208

import "sync"

// call is an in-flight or completed Do call
type __obf_5135b30c4a1c5011 struct {
	__obf_fdaf6d2017942e97 sync.WaitGroup
	__obf_676d75836c094b83 any
	__obf_0560c6b8f27080cc error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_bc58716fda429cf7 sync.Mutex                         // protects m
	__obf_17fa7ee2c9072df2 map[string]*__obf_5135b30c4a1c5011 // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_3c380430c3aa39c9 *Group) Do(__obf_4aecf3c737bbe5e8 string, __obf_0175bbcffc34863e func() (any, error)) (any, error) {
	__obf_3c380430c3aa39c9.__obf_bc58716fda429cf7.Lock()
	if __obf_3c380430c3aa39c9.__obf_17fa7ee2c9072df2 == nil {
		__obf_3c380430c3aa39c9.__obf_17fa7ee2c9072df2 = make(map[string]*__obf_5135b30c4a1c5011)
	}
	if __obf_1e8c1f17e8f9e9cd, __obf_4bca5d4f1d82d040 := __obf_3c380430c3aa39c9.__obf_17fa7ee2c9072df2[__obf_4aecf3c737bbe5e8]; __obf_4bca5d4f1d82d040 {
		__obf_3c380430c3aa39c9.__obf_bc58716fda429cf7.Unlock()
		__obf_1e8c1f17e8f9e9cd.__obf_fdaf6d2017942e97.Wait()
		return __obf_1e8c1f17e8f9e9cd.__obf_676d75836c094b83, __obf_1e8c1f17e8f9e9cd.__obf_0560c6b8f27080cc
	}
	__obf_1e8c1f17e8f9e9cd := new(__obf_5135b30c4a1c5011)
	__obf_1e8c1f17e8f9e9cd.__obf_fdaf6d2017942e97.Add(1)
	__obf_3c380430c3aa39c9.__obf_17fa7ee2c9072df2[__obf_4aecf3c737bbe5e8] = __obf_1e8c1f17e8f9e9cd
	__obf_3c380430c3aa39c9.__obf_bc58716fda429cf7.Unlock()

	__obf_1e8c1f17e8f9e9cd.__obf_676d75836c094b83, __obf_1e8c1f17e8f9e9cd.__obf_0560c6b8f27080cc = __obf_0175bbcffc34863e()
	__obf_1e8c1f17e8f9e9cd.__obf_fdaf6d2017942e97.Done()

	__obf_3c380430c3aa39c9.__obf_bc58716fda429cf7.Lock()
	delete(__obf_3c380430c3aa39c9.__obf_17fa7ee2c9072df2, __obf_4aecf3c737bbe5e8)
	__obf_3c380430c3aa39c9.__obf_bc58716fda429cf7.Unlock()

	return __obf_1e8c1f17e8f9e9cd.__obf_676d75836c094b83, __obf_1e8c1f17e8f9e9cd.__obf_0560c6b8f27080cc
}
