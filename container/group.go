// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_76599ab96a208083

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_f2c3ba527a3a24ba map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new:                    new,
		__obf_f2c3ba527a3a24ba: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_4711ef69cdf6c898 *Group) Get(__obf_4368715d0a6d4f0b string) any {
	__obf_4711ef69cdf6c898.RLock()
	__obf_93d90cff8a0252e8, __obf_d47b43f0df7ca990 := __obf_4711ef69cdf6c898.__obf_f2c3ba527a3a24ba[__obf_4368715d0a6d4f0b]
	if __obf_d47b43f0df7ca990 {
		__obf_4711ef69cdf6c898.RUnlock()
		return __obf_93d90cff8a0252e8
	}
	__obf_4711ef69cdf6c898.RUnlock()

	// slow path for group don`t have specified key value
	__obf_4711ef69cdf6c898.Lock()
	defer __obf_4711ef69cdf6c898.Unlock()
	__obf_93d90cff8a0252e8, __obf_d47b43f0df7ca990 = __obf_4711ef69cdf6c898.__obf_f2c3ba527a3a24ba[__obf_4368715d0a6d4f0b]
	if __obf_d47b43f0df7ca990 {
		return __obf_93d90cff8a0252e8
	}
	__obf_93d90cff8a0252e8 = __obf_4711ef69cdf6c898.new()
	__obf_4711ef69cdf6c898.__obf_f2c3ba527a3a24ba[__obf_4368715d0a6d4f0b] = __obf_93d90cff8a0252e8
	return __obf_93d90cff8a0252e8
}

// Reset resets the new function and deletes all existing objects.
func (__obf_4711ef69cdf6c898 *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_4711ef69cdf6c898.Lock()
	__obf_4711ef69cdf6c898.new = new
	__obf_4711ef69cdf6c898.Unlock()
	__obf_4711ef69cdf6c898.Clear()
}

// Clear deletes all objects.
func (__obf_4711ef69cdf6c898 *Group) Clear() {
	__obf_4711ef69cdf6c898.Lock()
	__obf_4711ef69cdf6c898.__obf_f2c3ba527a3a24ba = make(map[string]any)
	__obf_4711ef69cdf6c898.Unlock()
}
