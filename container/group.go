// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_9004b47202f9204b

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_489f8053148d8299 map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new: new, __obf_489f8053148d8299: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_890fc8f79511b521 *Group) Get(__obf_355e7c922e433678 string) any {
	__obf_890fc8f79511b521.
		RLock()
	__obf_8616b2f4d0fa5a6a, __obf_545047dd3a6c3c8d := __obf_890fc8f79511b521.__obf_489f8053148d8299[__obf_355e7c922e433678]
	if __obf_545047dd3a6c3c8d {
		__obf_890fc8f79511b521.
			RUnlock()
		return __obf_8616b2f4d0fa5a6a
	}
	__obf_890fc8f79511b521.
		RUnlock()
	__obf_890fc8f79511b521.

		// slow path for group don`t have specified key value
		Lock()
	defer __obf_890fc8f79511b521.Unlock()
	__obf_8616b2f4d0fa5a6a, __obf_545047dd3a6c3c8d = __obf_890fc8f79511b521.__obf_489f8053148d8299[__obf_355e7c922e433678]
	if __obf_545047dd3a6c3c8d {
		return __obf_8616b2f4d0fa5a6a
	}
	__obf_8616b2f4d0fa5a6a = __obf_890fc8f79511b521.new()
	__obf_890fc8f79511b521.__obf_489f8053148d8299[__obf_355e7c922e433678] = __obf_8616b2f4d0fa5a6a
	return __obf_8616b2f4d0fa5a6a
}

// Reset resets the new function and deletes all existing objects.
func (__obf_890fc8f79511b521 *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_890fc8f79511b521.
		Lock()
	__obf_890fc8f79511b521.
		new = new
	__obf_890fc8f79511b521.
		Unlock()
	__obf_890fc8f79511b521.
		Clear()
}

// Clear deletes all objects.
func (__obf_890fc8f79511b521 *Group) Clear() {
	__obf_890fc8f79511b521.
		Lock()
	__obf_890fc8f79511b521.__obf_489f8053148d8299 = make(map[string]any)
	__obf_890fc8f79511b521.
		Unlock()
}
