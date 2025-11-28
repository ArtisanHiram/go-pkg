// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_af42fb6cde2beed6

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_ecffc655bb698ea6 map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new:                    new,
		__obf_ecffc655bb698ea6: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_a21c61750a893db5 *Group) Get(__obf_55cde42f6d47c5be string) any {
	__obf_a21c61750a893db5.RLock()
	__obf_289dd915247ca25d, __obf_021756766e1db5da := __obf_a21c61750a893db5.__obf_ecffc655bb698ea6[__obf_55cde42f6d47c5be]
	if __obf_021756766e1db5da {
		__obf_a21c61750a893db5.RUnlock()
		return __obf_289dd915247ca25d
	}
	__obf_a21c61750a893db5.RUnlock()

	// slow path for group don`t have specified key value
	__obf_a21c61750a893db5.Lock()
	defer __obf_a21c61750a893db5.Unlock()
	__obf_289dd915247ca25d, __obf_021756766e1db5da = __obf_a21c61750a893db5.__obf_ecffc655bb698ea6[__obf_55cde42f6d47c5be]
	if __obf_021756766e1db5da {
		return __obf_289dd915247ca25d
	}
	__obf_289dd915247ca25d = __obf_a21c61750a893db5.new()
	__obf_a21c61750a893db5.__obf_ecffc655bb698ea6[__obf_55cde42f6d47c5be] = __obf_289dd915247ca25d
	return __obf_289dd915247ca25d
}

// Reset resets the new function and deletes all existing objects.
func (__obf_a21c61750a893db5 *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_a21c61750a893db5.Lock()
	__obf_a21c61750a893db5.new = new
	__obf_a21c61750a893db5.Unlock()
	__obf_a21c61750a893db5.Clear()
}

// Clear deletes all objects.
func (__obf_a21c61750a893db5 *Group) Clear() {
	__obf_a21c61750a893db5.Lock()
	__obf_a21c61750a893db5.__obf_ecffc655bb698ea6 = make(map[string]any)
	__obf_a21c61750a893db5.Unlock()
}
