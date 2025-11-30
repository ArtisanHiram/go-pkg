// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_038560a94647875f

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_fcda9cff02fa0ecd map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new: new, __obf_fcda9cff02fa0ecd: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_6b36aa7c49562329 *Group) Get(__obf_52e73bb48e810dd2 string) any {
	__obf_6b36aa7c49562329.
		RLock()
	__obf_5ec12fe59cf82410, __obf_5904187971b4cb86 := __obf_6b36aa7c49562329.__obf_fcda9cff02fa0ecd[__obf_52e73bb48e810dd2]
	if __obf_5904187971b4cb86 {
		__obf_6b36aa7c49562329.
			RUnlock()
		return __obf_5ec12fe59cf82410
	}
	__obf_6b36aa7c49562329.
		RUnlock()
	__obf_6b36aa7c49562329.

		// slow path for group don`t have specified key value
		Lock()
	defer __obf_6b36aa7c49562329.Unlock()
	__obf_5ec12fe59cf82410, __obf_5904187971b4cb86 = __obf_6b36aa7c49562329.__obf_fcda9cff02fa0ecd[__obf_52e73bb48e810dd2]
	if __obf_5904187971b4cb86 {
		return __obf_5ec12fe59cf82410
	}
	__obf_5ec12fe59cf82410 = __obf_6b36aa7c49562329.new()
	__obf_6b36aa7c49562329.__obf_fcda9cff02fa0ecd[__obf_52e73bb48e810dd2] = __obf_5ec12fe59cf82410
	return __obf_5ec12fe59cf82410
}

// Reset resets the new function and deletes all existing objects.
func (__obf_6b36aa7c49562329 *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_6b36aa7c49562329.
		Lock()
	__obf_6b36aa7c49562329.
		new = new
	__obf_6b36aa7c49562329.
		Unlock()
	__obf_6b36aa7c49562329.
		Clear()
}

// Clear deletes all objects.
func (__obf_6b36aa7c49562329 *Group) Clear() {
	__obf_6b36aa7c49562329.
		Lock()
	__obf_6b36aa7c49562329.__obf_fcda9cff02fa0ecd = make(map[string]any)
	__obf_6b36aa7c49562329.
		Unlock()
}
