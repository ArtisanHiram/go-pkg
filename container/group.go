// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_9861fa13140c30a3

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_010e40a589eedb1e map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new:                    new,
		__obf_010e40a589eedb1e: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_5cf43db6e628181f *Group) Get(__obf_43194ec765d86867 string) any {
	__obf_5cf43db6e628181f.RLock()
	__obf_478af39a6b826f97, __obf_da2280290c7dd975 := __obf_5cf43db6e628181f.__obf_010e40a589eedb1e[__obf_43194ec765d86867]
	if __obf_da2280290c7dd975 {
		__obf_5cf43db6e628181f.RUnlock()
		return __obf_478af39a6b826f97
	}
	__obf_5cf43db6e628181f.RUnlock()

	// slow path for group don`t have specified key value
	__obf_5cf43db6e628181f.Lock()
	defer __obf_5cf43db6e628181f.Unlock()
	__obf_478af39a6b826f97, __obf_da2280290c7dd975 = __obf_5cf43db6e628181f.__obf_010e40a589eedb1e[__obf_43194ec765d86867]
	if __obf_da2280290c7dd975 {
		return __obf_478af39a6b826f97
	}
	__obf_478af39a6b826f97 = __obf_5cf43db6e628181f.new()
	__obf_5cf43db6e628181f.__obf_010e40a589eedb1e[__obf_43194ec765d86867] = __obf_478af39a6b826f97
	return __obf_478af39a6b826f97
}

// Reset resets the new function and deletes all existing objects.
func (__obf_5cf43db6e628181f *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_5cf43db6e628181f.Lock()
	__obf_5cf43db6e628181f.new = new
	__obf_5cf43db6e628181f.Unlock()
	__obf_5cf43db6e628181f.Clear()
}

// Clear deletes all objects.
func (__obf_5cf43db6e628181f *Group) Clear() {
	__obf_5cf43db6e628181f.Lock()
	__obf_5cf43db6e628181f.__obf_010e40a589eedb1e = make(map[string]any)
	__obf_5cf43db6e628181f.Unlock()
}
