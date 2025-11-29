package __obf_a05682be1c6caf18

import "sync"

// call is an in-flight or completed Do call
type __obf_2675da139a84b934 struct {
	__obf_215eff0b76a00887 sync.WaitGroup
	__obf_6030b68d8a95172f any
	__obf_94de95cb433b50f1 error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_04ccec3be03c0675          sync.Mutex
	__obf_cf817044cb0e2a7b          map// protects m
	[string]*__obf_2675da139a84b934 // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_ba7d7da9d5a9ab18 *Group) Do(__obf_a0e56915a05e8d99 string, __obf_3bb5af003a0af361 func() (any, error)) (any, error) {
	__obf_ba7d7da9d5a9ab18.__obf_04ccec3be03c0675.
		Lock()
	if __obf_ba7d7da9d5a9ab18.__obf_cf817044cb0e2a7b == nil {
		__obf_ba7d7da9d5a9ab18.__obf_cf817044cb0e2a7b = make(map[string]*__obf_2675da139a84b934)
	}
	if __obf_365ba253fb24199c, __obf_19211ff80aaad9b4 := __obf_ba7d7da9d5a9ab18.__obf_cf817044cb0e2a7b[__obf_a0e56915a05e8d99]; __obf_19211ff80aaad9b4 {
		__obf_ba7d7da9d5a9ab18.__obf_04ccec3be03c0675.
			Unlock()
		__obf_365ba253fb24199c.__obf_215eff0b76a00887.
			Wait()
		return __obf_365ba253fb24199c.__obf_6030b68d8a95172f, __obf_365ba253fb24199c.__obf_94de95cb433b50f1
	}
	__obf_365ba253fb24199c := new(__obf_2675da139a84b934)
	__obf_365ba253fb24199c.__obf_215eff0b76a00887.
		Add(1)
	__obf_ba7d7da9d5a9ab18.__obf_cf817044cb0e2a7b[__obf_a0e56915a05e8d99] = __obf_365ba253fb24199c
	__obf_ba7d7da9d5a9ab18.__obf_04ccec3be03c0675.
		Unlock()
	__obf_365ba253fb24199c.__obf_6030b68d8a95172f, __obf_365ba253fb24199c.__obf_94de95cb433b50f1 = __obf_3bb5af003a0af361()
	__obf_365ba253fb24199c.__obf_215eff0b76a00887.
		Done()
	__obf_ba7d7da9d5a9ab18.__obf_04ccec3be03c0675.
		Lock()
	delete(__obf_ba7d7da9d5a9ab18.__obf_cf817044cb0e2a7b, __obf_a0e56915a05e8d99)
	__obf_ba7d7da9d5a9ab18.__obf_04ccec3be03c0675.
		Unlock()

	return __obf_365ba253fb24199c.__obf_6030b68d8a95172f, __obf_365ba253fb24199c.__obf_94de95cb433b50f1
}
