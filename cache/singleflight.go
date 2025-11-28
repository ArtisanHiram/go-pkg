package __obf_8281010a3a6d2ab0

import "sync"

// call is an in-flight or completed Do call
type __obf_9c1b0405b73fafb7 struct {
	__obf_6fea15d9f85fd424 sync.WaitGroup
	__obf_a31843a6764aecf9 any
	__obf_fd03130c6793cb3b error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_070e2f23960d0e82 sync.Mutex                         // protects m
	__obf_384bcb41d6b01071 map[string]*__obf_9c1b0405b73fafb7 // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_a96b5f1031c4d274 *Group) Do(__obf_805b9182f4a01a43 string, __obf_ba28f4cbe819430e func() (any, error)) (any, error) {
	__obf_a96b5f1031c4d274.__obf_070e2f23960d0e82.Lock()
	if __obf_a96b5f1031c4d274.__obf_384bcb41d6b01071 == nil {
		__obf_a96b5f1031c4d274.__obf_384bcb41d6b01071 = make(map[string]*__obf_9c1b0405b73fafb7)
	}
	if __obf_868ba86bf2a4661f, __obf_102da286c8b8a571 := __obf_a96b5f1031c4d274.__obf_384bcb41d6b01071[__obf_805b9182f4a01a43]; __obf_102da286c8b8a571 {
		__obf_a96b5f1031c4d274.__obf_070e2f23960d0e82.Unlock()
		__obf_868ba86bf2a4661f.__obf_6fea15d9f85fd424.Wait()
		return __obf_868ba86bf2a4661f.__obf_a31843a6764aecf9, __obf_868ba86bf2a4661f.__obf_fd03130c6793cb3b
	}
	__obf_868ba86bf2a4661f := new(__obf_9c1b0405b73fafb7)
	__obf_868ba86bf2a4661f.__obf_6fea15d9f85fd424.Add(1)
	__obf_a96b5f1031c4d274.__obf_384bcb41d6b01071[__obf_805b9182f4a01a43] = __obf_868ba86bf2a4661f
	__obf_a96b5f1031c4d274.__obf_070e2f23960d0e82.Unlock()

	__obf_868ba86bf2a4661f.__obf_a31843a6764aecf9, __obf_868ba86bf2a4661f.__obf_fd03130c6793cb3b = __obf_ba28f4cbe819430e()
	__obf_868ba86bf2a4661f.__obf_6fea15d9f85fd424.Done()

	__obf_a96b5f1031c4d274.__obf_070e2f23960d0e82.Lock()
	delete(__obf_a96b5f1031c4d274.__obf_384bcb41d6b01071, __obf_805b9182f4a01a43)
	__obf_a96b5f1031c4d274.__obf_070e2f23960d0e82.Unlock()

	return __obf_868ba86bf2a4661f.__obf_a31843a6764aecf9, __obf_868ba86bf2a4661f.__obf_fd03130c6793cb3b
}
