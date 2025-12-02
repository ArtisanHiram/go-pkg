// Package group provides a sample lazy load container.
// The group only creating a new object not until the object is needed by user.
// And it will cache all the objects to reduce the creation of object.
package __obf_4a16ef421fb74992

import "sync"

// Group is a lazy load container.
type Group struct {
	new                    func() any
	__obf_4f59e8768782d4bd map[string]any
	sync.RWMutex
}

// NewGroup news a group container.
func NewGroup(new func() any) *Group {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	return &Group{
		new: new, __obf_4f59e8768782d4bd: make(map[string]any),
	}
}

// Get gets the object by the given key.
func (__obf_1c490bd02b52fdeb *Group) Get(__obf_d1a1f98ae0fb119e string) any {
	__obf_1c490bd02b52fdeb.
		RLock()
	__obf_b3b7dced15e3b450, __obf_0a7b144ecfd706f0 := __obf_1c490bd02b52fdeb.__obf_4f59e8768782d4bd[__obf_d1a1f98ae0fb119e]
	if __obf_0a7b144ecfd706f0 {
		__obf_1c490bd02b52fdeb.
			RUnlock()
		return __obf_b3b7dced15e3b450
	}
	__obf_1c490bd02b52fdeb.
		RUnlock()
	__obf_1c490bd02b52fdeb.

		// slow path for group don`t have specified key value
		Lock()
	defer __obf_1c490bd02b52fdeb.Unlock()
	__obf_b3b7dced15e3b450, __obf_0a7b144ecfd706f0 = __obf_1c490bd02b52fdeb.__obf_4f59e8768782d4bd[__obf_d1a1f98ae0fb119e]
	if __obf_0a7b144ecfd706f0 {
		return __obf_b3b7dced15e3b450
	}
	__obf_b3b7dced15e3b450 = __obf_1c490bd02b52fdeb.new()
	__obf_1c490bd02b52fdeb.__obf_4f59e8768782d4bd[__obf_d1a1f98ae0fb119e] = __obf_b3b7dced15e3b450
	return __obf_b3b7dced15e3b450
}

// Reset resets the new function and deletes all existing objects.
func (__obf_1c490bd02b52fdeb *Group) Reset(new func() any) {
	if new == nil {
		panic("container.group: can't assign a nil to the new function")
	}
	__obf_1c490bd02b52fdeb.
		Lock()
	__obf_1c490bd02b52fdeb.
		new = new
	__obf_1c490bd02b52fdeb.
		Unlock()
	__obf_1c490bd02b52fdeb.
		Clear()
}

// Clear deletes all objects.
func (__obf_1c490bd02b52fdeb *Group) Clear() {
	__obf_1c490bd02b52fdeb.
		Lock()
	__obf_1c490bd02b52fdeb.__obf_4f59e8768782d4bd = make(map[string]any)
	__obf_1c490bd02b52fdeb.
		Unlock()
}
