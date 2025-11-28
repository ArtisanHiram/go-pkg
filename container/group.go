// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_1fda7fbdeda52f1e

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_c6f1eaada6fb59d9 map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new:                    new,
		__obf_c6f1eaada6fb59d9: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_bff32981c1e1be21 *Group) Get(__obf_95c4c09a56c1a070 string) any {
	__obf_bff32981c1e1be21.RLock()
	__obf_9371496cccb864c5, __obf_d24c61455fa2e4db := __obf_bff32981c1e1be21.__obf_c6f1eaada6fb59d9[__obf_95c4c09a56c1a070]
	if __obf_d24c61455fa2e4db {
		__obf_bff32981c1e1be21.RUnlock()
		return __obf_9371496cccb864c5
	}
	__obf_bff32981c1e1be21.RUnlock()

	// slow path for group don`t have specified key value
	__obf_bff32981c1e1be21.Lock()
	defer __obf_bff32981c1e1be21.Unlock()
	__obf_9371496cccb864c5, __obf_d24c61455fa2e4db = __obf_bff32981c1e1be21.__obf_c6f1eaada6fb59d9[__obf_95c4c09a56c1a070]
	if __obf_d24c61455fa2e4db {
		return __obf_9371496cccb864c5
	}
	__obf_9371496cccb864c5 = __obf_bff32981c1e1be21.new()
	__obf_bff32981c1e1be21.__obf_c6f1eaada6fb59d9[__obf_95c4c09a56c1a070] = __obf_9371496cccb864c5
	return __obf_9371496cccb864c5
}

// Reset resets the new function and deletes all existing objects.
func (__obf_bff32981c1e1be21 *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_bff32981c1e1be21.Lock()
	__obf_bff32981c1e1be21.new = new
	__obf_bff32981c1e1be21.Unlock()
	__obf_bff32981c1e1be21.Clear()
}

// Clear deletes all objects.
func (__obf_bff32981c1e1be21 *Group) Clear() {
	__obf_bff32981c1e1be21.Lock()
	__obf_bff32981c1e1be21.__obf_c6f1eaada6fb59d9 = make(map[string]any)
	__obf_bff32981c1e1be21.Unlock()
}
