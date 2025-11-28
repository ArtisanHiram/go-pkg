// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_4f13ac5aa043b5ce

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_ad12e06d8e8771c7 map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new:                    new,
		__obf_ad12e06d8e8771c7: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_cc822c2ab8d4fc6e *Group) Get(__obf_797707a3dac0ebb7 string) any {
	__obf_cc822c2ab8d4fc6e.RLock()
	__obf_1e75fd93d2a55aa7, __obf_c7b01e58066bbb7a := __obf_cc822c2ab8d4fc6e.__obf_ad12e06d8e8771c7[__obf_797707a3dac0ebb7]
	if __obf_c7b01e58066bbb7a {
		__obf_cc822c2ab8d4fc6e.RUnlock()
		return __obf_1e75fd93d2a55aa7
	}
	__obf_cc822c2ab8d4fc6e.RUnlock()

	// slow path for group don`t have specified key value
	__obf_cc822c2ab8d4fc6e.Lock()
	defer __obf_cc822c2ab8d4fc6e.Unlock()
	__obf_1e75fd93d2a55aa7, __obf_c7b01e58066bbb7a = __obf_cc822c2ab8d4fc6e.__obf_ad12e06d8e8771c7[__obf_797707a3dac0ebb7]
	if __obf_c7b01e58066bbb7a {
		return __obf_1e75fd93d2a55aa7
	}
	__obf_1e75fd93d2a55aa7 = __obf_cc822c2ab8d4fc6e.new()
	__obf_cc822c2ab8d4fc6e.__obf_ad12e06d8e8771c7[__obf_797707a3dac0ebb7] = __obf_1e75fd93d2a55aa7
	return __obf_1e75fd93d2a55aa7
}

// Reset resets the new function and deletes all existing objects.
func (__obf_cc822c2ab8d4fc6e *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_cc822c2ab8d4fc6e.Lock()
	__obf_cc822c2ab8d4fc6e.new = new
	__obf_cc822c2ab8d4fc6e.Unlock()
	__obf_cc822c2ab8d4fc6e.Clear()
}

// Clear deletes all objects.
func (__obf_cc822c2ab8d4fc6e *Group) Clear() {
	__obf_cc822c2ab8d4fc6e.Lock()
	__obf_cc822c2ab8d4fc6e.__obf_ad12e06d8e8771c7 = make(map[string]any)
	__obf_cc822c2ab8d4fc6e.Unlock()
}
