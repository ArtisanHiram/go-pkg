package __obf_9e1ee87c6b054458

import (
	"container/list"
	"fmt"
	util "github.com/ArtisanHiram/go-pkg/util"
	"strings"
	"sync"
	"time"
)

// Lru is an LRU cache. It is not safe for concurrent access.
type Lru struct {

	// MaxEntries is the maximum number of cache entries before
	// an item is evicted. Zero means no limit.
	MaxEntries int

	// OnEvicted optionally specifies a callback function to be
	// executed when an entry is purged from the cache.
	OnEvicted              func(__obf_3c13197612c6b39f string, __obf_2d0927453cc08d1b any)
	__obf_3cf6a609b9a5b8d0 *list.List
	__obf_9e1ee87c6b054458 map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_81f8848256aa9749 struct {
	__obf_3c13197612c6b39f string
	__obf_2d0927453cc08d1b []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_9c8431183e549025 int) Cache {
	return &Lru{
		MaxEntries: __obf_9c8431183e549025, __obf_3cf6a609b9a5b8d0: list.New(), __obf_9e1ee87c6b054458: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_8af2676446f64c49 *Lru) Has(__obf_3c13197612c6b39f string) (__obf_40eed10588aa6cec bool) {
	__obf_8af2676446f64c49.
		RLock()
	defer __obf_8af2676446f64c49.RUnlock()

	_, __obf_40eed10588aa6cec = __obf_8af2676446f64c49.__obf_9e1ee87c6b054458[__obf_3c13197612c6b39f]
	return __obf_40eed10588aa6cec
}

// Delete key by prefix
func (__obf_8af2676446f64c49 *Lru) Delete(__obf_14a292fb641d1b3d string) {
	__obf_8af2676446f64c49.
		Lock()
	defer __obf_8af2676446f64c49.Unlock()
	for __obf_ffc285490d834437, __obf_aa24543fddec7b80 := range __obf_8af2676446f64c49.__obf_9e1ee87c6b054458 {
		if strings.HasPrefix(__obf_ffc285490d834437, __obf_14a292fb641d1b3d) {
			__obf_8af2676446f64c49.__obf_0af7dba24b2ffeb7(__obf_aa24543fddec7b80)
		}
	}
}

func (__obf_8af2676446f64c49 *Lru) Set(__obf_3c13197612c6b39f string, __obf_7d0a67e130b6e9b4 any, __obf_5a47a67359bdcb92 time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_8af2676446f64c49 *Lru) Add(__obf_3c13197612c6b39f string, __obf_2d0927453cc08d1b any) error {
	__obf_8af2676446f64c49.
		Lock()
	defer __obf_8af2676446f64c49.Unlock()
	if __obf_8af2676446f64c49.__obf_9e1ee87c6b054458 == nil {
		__obf_8af2676446f64c49.__obf_9e1ee87c6b054458 = make(map[string]*list.Element)
		__obf_8af2676446f64c49.__obf_3cf6a609b9a5b8d0 = list.New()
	}
	if __obf_8465441720639bab, __obf_40eed10588aa6cec := __obf_8af2676446f64c49.__obf_9e1ee87c6b054458[__obf_3c13197612c6b39f]; __obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_3cf6a609b9a5b8d0.
			MoveToFront(__obf_8465441720639bab)
		var __obf_498673050542660c error
		__obf_8465441720639bab.
			Value.(*__obf_81f8848256aa9749).__obf_2d0927453cc08d1b, __obf_498673050542660c = util.AnyToBytes(__obf_2d0927453cc08d1b)
		return __obf_498673050542660c
	}
	__obf_7d0a67e130b6e9b4, __obf_498673050542660c := util.AnyToBytes(__obf_2d0927453cc08d1b)
	if __obf_498673050542660c != nil {
		return __obf_498673050542660c
	}
	__obf_97f63dc1928252ba := __obf_8af2676446f64c49.__obf_3cf6a609b9a5b8d0.PushFront(&__obf_81f8848256aa9749{__obf_3c13197612c6b39f, __obf_7d0a67e130b6e9b4})
	__obf_8af2676446f64c49.__obf_9e1ee87c6b054458[__obf_3c13197612c6b39f] = __obf_97f63dc1928252ba
	if __obf_8af2676446f64c49.MaxEntries != 0 && __obf_8af2676446f64c49.__obf_3cf6a609b9a5b8d0.Len() > __obf_8af2676446f64c49.MaxEntries {
		__obf_8af2676446f64c49.
			RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_8af2676446f64c49 *Lru) Get(__obf_3c13197612c6b39f string) (__obf_2d0927453cc08d1b []byte, __obf_498673050542660c error) {
	if __obf_8af2676446f64c49.__obf_9e1ee87c6b054458 == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_97f63dc1928252ba, __obf_abbf56da3e7cf93a := __obf_8af2676446f64c49.__obf_9e1ee87c6b054458[__obf_3c13197612c6b39f]; __obf_abbf56da3e7cf93a {
		__obf_8af2676446f64c49.__obf_3cf6a609b9a5b8d0.
			MoveToFront(__obf_97f63dc1928252ba)
		return __obf_97f63dc1928252ba.Value.(*__obf_81f8848256aa9749).__obf_2d0927453cc08d1b, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_3c13197612c6b39f)
}

// Remove removes the provided key from the cache.
func (__obf_8af2676446f64c49 *Lru) Remove(__obf_3c13197612c6b39f string) error {
	__obf_8af2676446f64c49.
		Lock()
	defer __obf_8af2676446f64c49.Unlock()
	if __obf_8af2676446f64c49.__obf_9e1ee87c6b054458 == nil {
		return ErrCacheEmpty
	}
	if __obf_97f63dc1928252ba, __obf_abbf56da3e7cf93a := __obf_8af2676446f64c49.__obf_9e1ee87c6b054458[__obf_3c13197612c6b39f]; __obf_abbf56da3e7cf93a {
		__obf_8af2676446f64c49.__obf_0af7dba24b2ffeb7(__obf_97f63dc1928252ba)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_8af2676446f64c49 *Lru) RemoveOldest() {
	__obf_8af2676446f64c49.
		Lock()
	defer __obf_8af2676446f64c49.Unlock()
	if __obf_8af2676446f64c49.__obf_9e1ee87c6b054458 == nil {
		return
	}
	__obf_97f63dc1928252ba := __obf_8af2676446f64c49.__obf_3cf6a609b9a5b8d0.Back()
	if __obf_97f63dc1928252ba != nil {
		__obf_8af2676446f64c49.__obf_0af7dba24b2ffeb7(__obf_97f63dc1928252ba)
	}
}

func (__obf_8af2676446f64c49 *Lru) __obf_0af7dba24b2ffeb7(__obf_4f22c5826c5e34df *list.Element) {
	__obf_8af2676446f64c49.__obf_3cf6a609b9a5b8d0.
		Remove(__obf_4f22c5826c5e34df)
	__obf_09ad994f24ec1125 := __obf_4f22c5826c5e34df.Value.(*__obf_81f8848256aa9749)
	delete(__obf_8af2676446f64c49.__obf_9e1ee87c6b054458, __obf_09ad994f24ec1125.__obf_3c13197612c6b39f)
	if __obf_8af2676446f64c49.OnEvicted != nil {
		__obf_8af2676446f64c49.
			OnEvicted(__obf_09ad994f24ec1125.__obf_3c13197612c6b39f,

				// Len returns the number of items in the cache.
				__obf_09ad994f24ec1125.__obf_2d0927453cc08d1b)
	}
}

func (__obf_8af2676446f64c49 *Lru) Len() int {
	if __obf_8af2676446f64c49.__obf_9e1ee87c6b054458 == nil {
		return 0
	}
	return __obf_8af2676446f64c49.

		// Clear purges all stored items from the cache.
		__obf_3cf6a609b9a5b8d0.Len()
}

func (__obf_8af2676446f64c49 *Lru) Clear() error {
	__obf_8af2676446f64c49.
		Lock()
	defer __obf_8af2676446f64c49.Unlock()
	if __obf_8af2676446f64c49.OnEvicted != nil {
		for _, __obf_4f22c5826c5e34df := range __obf_8af2676446f64c49.__obf_9e1ee87c6b054458 {
			__obf_09ad994f24ec1125 := __obf_4f22c5826c5e34df.Value.(*__obf_81f8848256aa9749)
			__obf_8af2676446f64c49.
				OnEvicted(__obf_09ad994f24ec1125.__obf_3c13197612c6b39f, __obf_09ad994f24ec1125.__obf_2d0927453cc08d1b)
		}
	}
	__obf_8af2676446f64c49.__obf_3cf6a609b9a5b8d0 = nil
	__obf_8af2676446f64c49.__obf_9e1ee87c6b054458 = nil
	return nil
}
