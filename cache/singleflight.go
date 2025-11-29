package __obf_6fd4fe34e3f784df

import "sync"

// call is an in-flight or completed Do call
type __obf_56215d966f991469 struct {
	__obf_fc2a37b7622a7ea3 sync.WaitGroup
	__obf_e1d0258259217620 any
	__obf_7f8f5f5c173fa58e error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_53983106240f1755          sync.Mutex
	__obf_1674045fb710ff20          map// protects m
	[string]*__obf_56215d966f991469 // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_1b7c463ccb263688 *Group) Do(__obf_ca986a38d1f8fbbb string, __obf_011b3cff87e3a2ba func() (any, error)) (any, error) {
	__obf_1b7c463ccb263688.__obf_53983106240f1755.
		Lock()
	if __obf_1b7c463ccb263688.__obf_1674045fb710ff20 == nil {
		__obf_1b7c463ccb263688.__obf_1674045fb710ff20 = make(map[string]*__obf_56215d966f991469)
	}
	if __obf_3cd39772de5fd3ee, __obf_949c761ffd90d5fa := __obf_1b7c463ccb263688.__obf_1674045fb710ff20[__obf_ca986a38d1f8fbbb]; __obf_949c761ffd90d5fa {
		__obf_1b7c463ccb263688.__obf_53983106240f1755.
			Unlock()
		__obf_3cd39772de5fd3ee.__obf_fc2a37b7622a7ea3.
			Wait()
		return __obf_3cd39772de5fd3ee.__obf_e1d0258259217620, __obf_3cd39772de5fd3ee.__obf_7f8f5f5c173fa58e
	}
	__obf_3cd39772de5fd3ee := new(__obf_56215d966f991469)
	__obf_3cd39772de5fd3ee.__obf_fc2a37b7622a7ea3.
		Add(1)
	__obf_1b7c463ccb263688.__obf_1674045fb710ff20[__obf_ca986a38d1f8fbbb] = __obf_3cd39772de5fd3ee
	__obf_1b7c463ccb263688.__obf_53983106240f1755.
		Unlock()
	__obf_3cd39772de5fd3ee.__obf_e1d0258259217620, __obf_3cd39772de5fd3ee.__obf_7f8f5f5c173fa58e = __obf_011b3cff87e3a2ba()
	__obf_3cd39772de5fd3ee.__obf_fc2a37b7622a7ea3.
		Done()
	__obf_1b7c463ccb263688.__obf_53983106240f1755.
		Lock()
	delete(__obf_1b7c463ccb263688.__obf_1674045fb710ff20, __obf_ca986a38d1f8fbbb)
	__obf_1b7c463ccb263688.__obf_53983106240f1755.
		Unlock()

	return __obf_3cd39772de5fd3ee.__obf_e1d0258259217620, __obf_3cd39772de5fd3ee.__obf_7f8f5f5c173fa58e
}
