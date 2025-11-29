// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_b0bebe5eb45b8ad6

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_d054d4311affb29a map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new: new, __obf_d054d4311affb29a: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_b8fc063d263e7f36 *Group) Get(__obf_9c15798bcb95be3e string) any {
	__obf_b8fc063d263e7f36.
		RLock()
	__obf_10522304adc21daf, __obf_e8c1fb9f7287beef := __obf_b8fc063d263e7f36.__obf_d054d4311affb29a[__obf_9c15798bcb95be3e]
	if __obf_e8c1fb9f7287beef {
		__obf_b8fc063d263e7f36.
			RUnlock()
		return __obf_10522304adc21daf
	}
	__obf_b8fc063d263e7f36.
		RUnlock()
	__obf_b8fc063d263e7f36.

		// slow path for group don`t have specified key value
		Lock()
	defer __obf_b8fc063d263e7f36.Unlock()
	__obf_10522304adc21daf, __obf_e8c1fb9f7287beef = __obf_b8fc063d263e7f36.__obf_d054d4311affb29a[__obf_9c15798bcb95be3e]
	if __obf_e8c1fb9f7287beef {
		return __obf_10522304adc21daf
	}
	__obf_10522304adc21daf = __obf_b8fc063d263e7f36.new()
	__obf_b8fc063d263e7f36.__obf_d054d4311affb29a[__obf_9c15798bcb95be3e] = __obf_10522304adc21daf
	return __obf_10522304adc21daf
}

// Reset resets the new function and deletes all existing objects.
func (__obf_b8fc063d263e7f36 *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_b8fc063d263e7f36.
		Lock()
	__obf_b8fc063d263e7f36.
		new = new
	__obf_b8fc063d263e7f36.
		Unlock()
	__obf_b8fc063d263e7f36.
		Clear()
}

// Clear deletes all objects.
func (__obf_b8fc063d263e7f36 *Group) Clear() {
	__obf_b8fc063d263e7f36.
		Lock()
	__obf_b8fc063d263e7f36.__obf_d054d4311affb29a = make(map[string]any)
	__obf_b8fc063d263e7f36.
		Unlock()
}
