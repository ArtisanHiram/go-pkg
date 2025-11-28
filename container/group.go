// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_62eba4024f8fa381

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_cfff6b37b91c3027 map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new:                    new,
		__obf_cfff6b37b91c3027: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_85a23f58f48352a4 *Group) Get(__obf_df070ab4c712506c string) any {
	__obf_85a23f58f48352a4.RLock()
	__obf_f69c1647001964ce, __obf_b47c9f51c1edd10c := __obf_85a23f58f48352a4.__obf_cfff6b37b91c3027[__obf_df070ab4c712506c]
	if __obf_b47c9f51c1edd10c {
		__obf_85a23f58f48352a4.RUnlock()
		return __obf_f69c1647001964ce
	}
	__obf_85a23f58f48352a4.RUnlock()

	// slow path for group don`t have specified key value
	__obf_85a23f58f48352a4.Lock()
	defer __obf_85a23f58f48352a4.Unlock()
	__obf_f69c1647001964ce, __obf_b47c9f51c1edd10c = __obf_85a23f58f48352a4.__obf_cfff6b37b91c3027[__obf_df070ab4c712506c]
	if __obf_b47c9f51c1edd10c {
		return __obf_f69c1647001964ce
	}
	__obf_f69c1647001964ce = __obf_85a23f58f48352a4.new()
	__obf_85a23f58f48352a4.__obf_cfff6b37b91c3027[__obf_df070ab4c712506c] = __obf_f69c1647001964ce
	return __obf_f69c1647001964ce
}

// Reset resets the new function and deletes all existing objects.
func (__obf_85a23f58f48352a4 *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_85a23f58f48352a4.Lock()
	__obf_85a23f58f48352a4.new = new
	__obf_85a23f58f48352a4.Unlock()
	__obf_85a23f58f48352a4.Clear()
}

// Clear deletes all objects.
func (__obf_85a23f58f48352a4 *Group) Clear() {
	__obf_85a23f58f48352a4.Lock()
	__obf_85a23f58f48352a4.__obf_cfff6b37b91c3027 = make(map[string]any)
	__obf_85a23f58f48352a4.Unlock()
}
