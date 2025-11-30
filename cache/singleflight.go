package __obf_72d962f6a40bc95f

import "sync"

// call is an in-flight or completed Do call
type __obf_a4a515b885202b8c struct {
	__obf_cebba8baed09318f sync.WaitGroup
	__obf_1dbdf9c3e13789b0 any
	__obf_2498adaec5f4c8f1 error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_09437e388e205e38          sync.Mutex
	__obf_a5a77ce672c162a7          map// protects m
	[string]*__obf_a4a515b885202b8c // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_e33f1a7ee6731fa7 *Group) Do(__obf_73cb8148cbf55699 string, __obf_eb935ebd3578a707 func() (any, error)) (any, error) {
	__obf_e33f1a7ee6731fa7.__obf_09437e388e205e38.
		Lock()
	if __obf_e33f1a7ee6731fa7.__obf_a5a77ce672c162a7 == nil {
		__obf_e33f1a7ee6731fa7.__obf_a5a77ce672c162a7 = make(map[string]*__obf_a4a515b885202b8c)
	}
	if __obf_89bd5bf0a2a3aefa, __obf_97faa3bbfe732802 := __obf_e33f1a7ee6731fa7.__obf_a5a77ce672c162a7[__obf_73cb8148cbf55699]; __obf_97faa3bbfe732802 {
		__obf_e33f1a7ee6731fa7.__obf_09437e388e205e38.
			Unlock()
		__obf_89bd5bf0a2a3aefa.__obf_cebba8baed09318f.
			Wait()
		return __obf_89bd5bf0a2a3aefa.__obf_1dbdf9c3e13789b0, __obf_89bd5bf0a2a3aefa.__obf_2498adaec5f4c8f1
	}
	__obf_89bd5bf0a2a3aefa := new(__obf_a4a515b885202b8c)
	__obf_89bd5bf0a2a3aefa.__obf_cebba8baed09318f.
		Add(1)
	__obf_e33f1a7ee6731fa7.__obf_a5a77ce672c162a7[__obf_73cb8148cbf55699] = __obf_89bd5bf0a2a3aefa
	__obf_e33f1a7ee6731fa7.__obf_09437e388e205e38.
		Unlock()
	__obf_89bd5bf0a2a3aefa.__obf_1dbdf9c3e13789b0, __obf_89bd5bf0a2a3aefa.__obf_2498adaec5f4c8f1 = __obf_eb935ebd3578a707()
	__obf_89bd5bf0a2a3aefa.__obf_cebba8baed09318f.
		Done()
	__obf_e33f1a7ee6731fa7.__obf_09437e388e205e38.
		Lock()
	delete(__obf_e33f1a7ee6731fa7.__obf_a5a77ce672c162a7, __obf_73cb8148cbf55699)
	__obf_e33f1a7ee6731fa7.__obf_09437e388e205e38.
		Unlock()

	return __obf_89bd5bf0a2a3aefa.__obf_1dbdf9c3e13789b0, __obf_89bd5bf0a2a3aefa.__obf_2498adaec5f4c8f1
}
