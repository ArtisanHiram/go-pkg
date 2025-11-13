package __obf_f4d8558b35981340

import "sync"

// call is an in-flight or completed Do call
type __obf_d4fbb1153fff816b struct {
	__obf_82c21c5dcc11fbda sync.WaitGroup
	__obf_20a54be727e70f3c any
	__obf_2d5e4b42e300bedd error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	__obf_03f513e727a97a9a sync.Mutex                         // protects m
	__obf_1f6d94b11ce336b6 map[string]*__obf_d4fbb1153fff816b // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (__obf_b652a32a46f6cdb5 *Group) Do(__obf_6208471da03258dd string, __obf_e01a2258903c9396 func() (any, error)) (any, error) {
	__obf_b652a32a46f6cdb5.__obf_03f513e727a97a9a.Lock()
	if __obf_b652a32a46f6cdb5.__obf_1f6d94b11ce336b6 == nil {
		__obf_b652a32a46f6cdb5.__obf_1f6d94b11ce336b6 = make(map[string]*__obf_d4fbb1153fff816b)
	}
	if __obf_a744f58475749237, __obf_ed10eafa7bc79f22 := __obf_b652a32a46f6cdb5.__obf_1f6d94b11ce336b6[__obf_6208471da03258dd]; __obf_ed10eafa7bc79f22 {
		__obf_b652a32a46f6cdb5.__obf_03f513e727a97a9a.Unlock()
		__obf_a744f58475749237.__obf_82c21c5dcc11fbda.Wait()
		return __obf_a744f58475749237.__obf_20a54be727e70f3c, __obf_a744f58475749237.__obf_2d5e4b42e300bedd
	}
	__obf_a744f58475749237 := new(__obf_d4fbb1153fff816b)
	__obf_a744f58475749237.__obf_82c21c5dcc11fbda.Add(1)
	__obf_b652a32a46f6cdb5.__obf_1f6d94b11ce336b6[__obf_6208471da03258dd] = __obf_a744f58475749237
	__obf_b652a32a46f6cdb5.__obf_03f513e727a97a9a.Unlock()

	__obf_a744f58475749237.__obf_20a54be727e70f3c, __obf_a744f58475749237.__obf_2d5e4b42e300bedd = __obf_e01a2258903c9396()
	__obf_a744f58475749237.__obf_82c21c5dcc11fbda.Done()

	__obf_b652a32a46f6cdb5.__obf_03f513e727a97a9a.Lock()
	delete(__obf_b652a32a46f6cdb5.__obf_1f6d94b11ce336b6, __obf_6208471da03258dd)
	__obf_b652a32a46f6cdb5.__obf_03f513e727a97a9a.Unlock()

	return __obf_a744f58475749237.__obf_20a54be727e70f3c, __obf_a744f58475749237.__obf_2d5e4b42e300bedd
}
