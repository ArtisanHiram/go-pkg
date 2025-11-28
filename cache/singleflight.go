package __obf_2f5c14e012cec42e

import "sync"

// call is an in-flight or completed Do call
type __obf_956e636928c02864 struct {
	__obf_6a4878a8ca84762f sync.WaitGroup
	__obf_aee38dd09d94cdd5 any
	__obf_956c4015cf7da152 error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_1791af550e1052ce sync.Mutex                         // protects m
	__obf_b127f8671154d05f map[string]*__obf_956e636928c02864 // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_2d4acf5eae184044 *Group) Do(__obf_f0b3ebeba048b5e4 string, __obf_a8209df76c3bc8e8 func() (any, error)) (any, error) {
	__obf_2d4acf5eae184044.__obf_1791af550e1052ce.Lock()
	if __obf_2d4acf5eae184044.__obf_b127f8671154d05f == nil {
		__obf_2d4acf5eae184044.__obf_b127f8671154d05f = make(map[string]*__obf_956e636928c02864)
	}
	if __obf_4379f98bc2ed9ecc, __obf_a39fb2a67596cf35 := __obf_2d4acf5eae184044.__obf_b127f8671154d05f[__obf_f0b3ebeba048b5e4]; __obf_a39fb2a67596cf35 {
		__obf_2d4acf5eae184044.__obf_1791af550e1052ce.Unlock()
		__obf_4379f98bc2ed9ecc.__obf_6a4878a8ca84762f.Wait()
		return __obf_4379f98bc2ed9ecc.__obf_aee38dd09d94cdd5, __obf_4379f98bc2ed9ecc.__obf_956c4015cf7da152
	}
	__obf_4379f98bc2ed9ecc := new(__obf_956e636928c02864)
	__obf_4379f98bc2ed9ecc.__obf_6a4878a8ca84762f.Add(1)
	__obf_2d4acf5eae184044.__obf_b127f8671154d05f[__obf_f0b3ebeba048b5e4] = __obf_4379f98bc2ed9ecc
	__obf_2d4acf5eae184044.__obf_1791af550e1052ce.Unlock()

	__obf_4379f98bc2ed9ecc.__obf_aee38dd09d94cdd5, __obf_4379f98bc2ed9ecc.__obf_956c4015cf7da152 = __obf_a8209df76c3bc8e8()
	__obf_4379f98bc2ed9ecc.__obf_6a4878a8ca84762f.Done()

	__obf_2d4acf5eae184044.__obf_1791af550e1052ce.Lock()
	delete(__obf_2d4acf5eae184044.__obf_b127f8671154d05f, __obf_f0b3ebeba048b5e4)
	__obf_2d4acf5eae184044.__obf_1791af550e1052ce.Unlock()

	return __obf_4379f98bc2ed9ecc.__obf_aee38dd09d94cdd5, __obf_4379f98bc2ed9ecc.__obf_956c4015cf7da152
}
