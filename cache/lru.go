package __obf_79e7502d8563d901

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
	OnEvicted func(__obf_50994613b7653a88 string, __obf_c0aa5e8a46724fa3 any)

	__obf_73a96acf05836f19 *list.List
	__obf_79e7502d8563d901 map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_66b90efe482cfb66 struct {
	__obf_50994613b7653a88 string
	__obf_c0aa5e8a46724fa3 []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_d7556e7b7b1dc12c int) Cache {
	return &Lru{
		MaxEntries:             __obf_d7556e7b7b1dc12c,
		__obf_73a96acf05836f19: list.New(),
		__obf_79e7502d8563d901: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_db09942aea00505a *Lru) Has(__obf_50994613b7653a88 string) (__obf_7500628aeb1f47ab bool) {
	__obf_db09942aea00505a.RLock()
	defer __obf_db09942aea00505a.RUnlock()

	_, __obf_7500628aeb1f47ab = __obf_db09942aea00505a.__obf_79e7502d8563d901[__obf_50994613b7653a88]
	return __obf_7500628aeb1f47ab
}

// Delete key by prefix
func (__obf_db09942aea00505a *Lru) Delete(__obf_bac226386ec983d7 string) {
	__obf_db09942aea00505a.Lock()
	defer __obf_db09942aea00505a.Unlock()
	for __obf_08ceb9044b83b8fd, __obf_bfb4c8a1c77523af := range __obf_db09942aea00505a.__obf_79e7502d8563d901 {
		if strings.HasPrefix(__obf_08ceb9044b83b8fd, __obf_bac226386ec983d7) {
			__obf_db09942aea00505a.__obf_a2e1b538a936b750(__obf_bfb4c8a1c77523af)
		}
	}
}

func (__obf_db09942aea00505a *Lru) Set(__obf_50994613b7653a88 string, __obf_e043381e0771feca any, __obf_4ff6d8752a4fea92 time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_db09942aea00505a *Lru) Add(__obf_50994613b7653a88 string, __obf_c0aa5e8a46724fa3 any) error {
	__obf_db09942aea00505a.Lock()
	defer __obf_db09942aea00505a.Unlock()
	if __obf_db09942aea00505a.__obf_79e7502d8563d901 == nil {
		__obf_db09942aea00505a.__obf_79e7502d8563d901 = make(map[string]*list.Element)
		__obf_db09942aea00505a.__obf_73a96acf05836f19 = list.New()
	}
	if __obf_f021144e52e8aaa6, __obf_7500628aeb1f47ab := __obf_db09942aea00505a.__obf_79e7502d8563d901[__obf_50994613b7653a88]; __obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_73a96acf05836f19.MoveToFront(__obf_f021144e52e8aaa6)
		var __obf_8d46af2525fab46a error
		__obf_f021144e52e8aaa6.Value.(*__obf_66b90efe482cfb66).__obf_c0aa5e8a46724fa3, __obf_8d46af2525fab46a = util.AnyToBytes(__obf_c0aa5e8a46724fa3)
		return __obf_8d46af2525fab46a
	}
	__obf_e043381e0771feca, __obf_8d46af2525fab46a := util.AnyToBytes(__obf_c0aa5e8a46724fa3)
	if __obf_8d46af2525fab46a != nil {
		return __obf_8d46af2525fab46a
	}
	__obf_83a81e04e46ea7b6 := __obf_db09942aea00505a.__obf_73a96acf05836f19.PushFront(&__obf_66b90efe482cfb66{__obf_50994613b7653a88, __obf_e043381e0771feca})
	__obf_db09942aea00505a.__obf_79e7502d8563d901[__obf_50994613b7653a88] = __obf_83a81e04e46ea7b6
	if __obf_db09942aea00505a.MaxEntries != 0 && __obf_db09942aea00505a.__obf_73a96acf05836f19.Len() > __obf_db09942aea00505a.MaxEntries {
		__obf_db09942aea00505a.RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_db09942aea00505a *Lru) Get(__obf_50994613b7653a88 string) (__obf_c0aa5e8a46724fa3 []byte, __obf_8d46af2525fab46a error) {
	if __obf_db09942aea00505a.__obf_79e7502d8563d901 == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_83a81e04e46ea7b6, __obf_4e903a24167b6665 := __obf_db09942aea00505a.__obf_79e7502d8563d901[__obf_50994613b7653a88]; __obf_4e903a24167b6665 {
		__obf_db09942aea00505a.__obf_73a96acf05836f19.MoveToFront(__obf_83a81e04e46ea7b6)
		return __obf_83a81e04e46ea7b6.Value.(*__obf_66b90efe482cfb66).__obf_c0aa5e8a46724fa3, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_50994613b7653a88)
}

// Remove removes the provided key from the cache.
func (__obf_db09942aea00505a *Lru) Remove(__obf_50994613b7653a88 string) error {
	__obf_db09942aea00505a.Lock()
	defer __obf_db09942aea00505a.Unlock()
	if __obf_db09942aea00505a.__obf_79e7502d8563d901 == nil {
		return ErrCacheEmpty
	}
	if __obf_83a81e04e46ea7b6, __obf_4e903a24167b6665 := __obf_db09942aea00505a.__obf_79e7502d8563d901[__obf_50994613b7653a88]; __obf_4e903a24167b6665 {
		__obf_db09942aea00505a.__obf_a2e1b538a936b750(__obf_83a81e04e46ea7b6)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_db09942aea00505a *Lru) RemoveOldest() {
	__obf_db09942aea00505a.Lock()
	defer __obf_db09942aea00505a.Unlock()
	if __obf_db09942aea00505a.__obf_79e7502d8563d901 == nil {
		return
	}
	__obf_83a81e04e46ea7b6 := __obf_db09942aea00505a.__obf_73a96acf05836f19.Back()
	if __obf_83a81e04e46ea7b6 != nil {
		__obf_db09942aea00505a.__obf_a2e1b538a936b750(__obf_83a81e04e46ea7b6)
	}
}

func (__obf_db09942aea00505a *Lru) __obf_a2e1b538a936b750(__obf_8429338405294247 *list.Element) {
	__obf_db09942aea00505a.__obf_73a96acf05836f19.Remove(__obf_8429338405294247)
	__obf_863db98f55e26f91 := __obf_8429338405294247.Value.(*__obf_66b90efe482cfb66)
	delete(__obf_db09942aea00505a.__obf_79e7502d8563d901, __obf_863db98f55e26f91.__obf_50994613b7653a88)
	if __obf_db09942aea00505a.OnEvicted != nil {
		__obf_db09942aea00505a.OnEvicted(__obf_863db98f55e26f91.__obf_50994613b7653a88, __obf_863db98f55e26f91.__obf_c0aa5e8a46724fa3)
	}
}

// Len returns the number of items in the cache.
func (__obf_db09942aea00505a *Lru) Len() int {
	if __obf_db09942aea00505a.__obf_79e7502d8563d901 == nil {
		return 0
	}
	return __obf_db09942aea00505a.__obf_73a96acf05836f19.Len()
}

// Clear purges all stored items from the cache.
func (__obf_db09942aea00505a *Lru) Clear() error {
	__obf_db09942aea00505a.Lock()
	defer __obf_db09942aea00505a.Unlock()
	if __obf_db09942aea00505a.OnEvicted != nil {
		for _, __obf_8429338405294247 := range __obf_db09942aea00505a.__obf_79e7502d8563d901 {
			__obf_863db98f55e26f91 := __obf_8429338405294247.Value.(*__obf_66b90efe482cfb66)
			__obf_db09942aea00505a.OnEvicted(__obf_863db98f55e26f91.__obf_50994613b7653a88, __obf_863db98f55e26f91.__obf_c0aa5e8a46724fa3)
		}
	}
	__obf_db09942aea00505a.__obf_73a96acf05836f19 = nil
	__obf_db09942aea00505a.__obf_79e7502d8563d901 = nil
	return nil
}
