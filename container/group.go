// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_e72ce603d10d02a1

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_7395e77052e9e60b map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new: new, __obf_7395e77052e9e60b: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_bdcddb3344666a74 *Group) Get(__obf_9dd08479fe906662 string) any {
	__obf_bdcddb3344666a74.
		RLock()
	__obf_ebe2552a06e91009, __obf_b95641e1ff7de824 := __obf_bdcddb3344666a74.__obf_7395e77052e9e60b[__obf_9dd08479fe906662]
	if __obf_b95641e1ff7de824 {
		__obf_bdcddb3344666a74.
			RUnlock()
		return __obf_ebe2552a06e91009
	}
	__obf_bdcddb3344666a74.
		RUnlock()
	__obf_bdcddb3344666a74.

		// slow path for group don`t have specified key value
		Lock()
	defer __obf_bdcddb3344666a74.Unlock()
	__obf_ebe2552a06e91009, __obf_b95641e1ff7de824 = __obf_bdcddb3344666a74.__obf_7395e77052e9e60b[__obf_9dd08479fe906662]
	if __obf_b95641e1ff7de824 {
		return __obf_ebe2552a06e91009
	}
	__obf_ebe2552a06e91009 = __obf_bdcddb3344666a74.new()
	__obf_bdcddb3344666a74.__obf_7395e77052e9e60b[__obf_9dd08479fe906662] = __obf_ebe2552a06e91009
	return __obf_ebe2552a06e91009
}

// Reset resets the new function and deletes all existing objects.
func (__obf_bdcddb3344666a74 *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_bdcddb3344666a74.
		Lock()
	__obf_bdcddb3344666a74.
		new = new
	__obf_bdcddb3344666a74.
		Unlock()
	__obf_bdcddb3344666a74.
		Clear()
}

// Clear deletes all objects.
func (__obf_bdcddb3344666a74 *Group) Clear() {
	__obf_bdcddb3344666a74.
		Lock()
	__obf_bdcddb3344666a74.__obf_7395e77052e9e60b = make(map[string]any)
	__obf_bdcddb3344666a74.
		Unlock()
}
