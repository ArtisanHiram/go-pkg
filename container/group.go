// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_90a4883a02d0b41c

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_c20d68262da83750 map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new: new, __obf_c20d68262da83750: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_f5cee4c381a6964f *Group) Get(__obf_5bba24c1758bbf28 string) any {
	__obf_f5cee4c381a6964f.
		RLock()
	__obf_1c35129ed9e54186, __obf_c16a0dd3a76f078a := __obf_f5cee4c381a6964f.__obf_c20d68262da83750[__obf_5bba24c1758bbf28]
	if __obf_c16a0dd3a76f078a {
		__obf_f5cee4c381a6964f.
			RUnlock()
		return __obf_1c35129ed9e54186
	}
	__obf_f5cee4c381a6964f.
		RUnlock()
	__obf_f5cee4c381a6964f.

		// slow path for group don`t have specified key value
		Lock()
	defer __obf_f5cee4c381a6964f.Unlock()
	__obf_1c35129ed9e54186, __obf_c16a0dd3a76f078a = __obf_f5cee4c381a6964f.__obf_c20d68262da83750[__obf_5bba24c1758bbf28]
	if __obf_c16a0dd3a76f078a {
		return __obf_1c35129ed9e54186
	}
	__obf_1c35129ed9e54186 = __obf_f5cee4c381a6964f.new()
	__obf_f5cee4c381a6964f.__obf_c20d68262da83750[__obf_5bba24c1758bbf28] = __obf_1c35129ed9e54186
	return __obf_1c35129ed9e54186
}

// Reset resets the new function and deletes all existing objects.
func (__obf_f5cee4c381a6964f *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_f5cee4c381a6964f.
		Lock()
	__obf_f5cee4c381a6964f.
		new = new
	__obf_f5cee4c381a6964f.
		Unlock()
	__obf_f5cee4c381a6964f.
		Clear()
}

// Clear deletes all objects.
func (__obf_f5cee4c381a6964f *Group) Clear() {
	__obf_f5cee4c381a6964f.
		Lock()
	__obf_f5cee4c381a6964f.__obf_c20d68262da83750 = make(map[string]any)
	__obf_f5cee4c381a6964f.
		Unlock()
}
