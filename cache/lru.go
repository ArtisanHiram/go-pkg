package __obf_62df4de078c8d208

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
	OnEvicted func(__obf_4aecf3c737bbe5e8 string, __obf_b99b98b3128d77a4 any)

	__obf_1726174d38a0b3b0 *list.List
	__obf_62df4de078c8d208 map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_8ca3bda6ee9c1b85 struct {
	__obf_4aecf3c737bbe5e8 string
	__obf_b99b98b3128d77a4 []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_4cdafb7d71158829 int) Cache {
	return &Lru{
		MaxEntries:             __obf_4cdafb7d71158829,
		__obf_1726174d38a0b3b0: list.New(),
		__obf_62df4de078c8d208: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_1e8c1f17e8f9e9cd *Lru) Has(__obf_4aecf3c737bbe5e8 string) (__obf_4bca5d4f1d82d040 bool) {
	__obf_1e8c1f17e8f9e9cd.RLock()
	defer __obf_1e8c1f17e8f9e9cd.RUnlock()

	_, __obf_4bca5d4f1d82d040 = __obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208[__obf_4aecf3c737bbe5e8]
	return __obf_4bca5d4f1d82d040
}

// Delete key by prefix
func (__obf_1e8c1f17e8f9e9cd *Lru) Delete(__obf_0c00b35780000392 string) {
	__obf_1e8c1f17e8f9e9cd.Lock()
	defer __obf_1e8c1f17e8f9e9cd.Unlock()
	for __obf_c0676b3110071e57, __obf_298afaabecd63bb9 := range __obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208 {
		if strings.HasPrefix(__obf_c0676b3110071e57, __obf_0c00b35780000392) {
			__obf_1e8c1f17e8f9e9cd.__obf_91180fb18d0699f5(__obf_298afaabecd63bb9)
		}
	}
}

func (__obf_1e8c1f17e8f9e9cd *Lru) Set(__obf_4aecf3c737bbe5e8 string, __obf_676d75836c094b83 any, __obf_24d92102aed9ce02 time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_1e8c1f17e8f9e9cd *Lru) Add(__obf_4aecf3c737bbe5e8 string, __obf_b99b98b3128d77a4 any) error {
	__obf_1e8c1f17e8f9e9cd.Lock()
	defer __obf_1e8c1f17e8f9e9cd.Unlock()
	if __obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208 == nil {
		__obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208 = make(map[string]*list.Element)
		__obf_1e8c1f17e8f9e9cd.__obf_1726174d38a0b3b0 = list.New()
	}
	if __obf_e80208bbfa1d0a79, __obf_4bca5d4f1d82d040 := __obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208[__obf_4aecf3c737bbe5e8]; __obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_1726174d38a0b3b0.MoveToFront(__obf_e80208bbfa1d0a79)
		var __obf_0560c6b8f27080cc error
		__obf_e80208bbfa1d0a79.Value.(*__obf_8ca3bda6ee9c1b85).__obf_b99b98b3128d77a4, __obf_0560c6b8f27080cc = util.AnyToBytes(__obf_b99b98b3128d77a4)
		return __obf_0560c6b8f27080cc
	}
	__obf_676d75836c094b83, __obf_0560c6b8f27080cc := util.AnyToBytes(__obf_b99b98b3128d77a4)
	if __obf_0560c6b8f27080cc != nil {
		return __obf_0560c6b8f27080cc
	}
	__obf_2a8198aab01675a0 := __obf_1e8c1f17e8f9e9cd.__obf_1726174d38a0b3b0.PushFront(&__obf_8ca3bda6ee9c1b85{__obf_4aecf3c737bbe5e8, __obf_676d75836c094b83})
	__obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208[__obf_4aecf3c737bbe5e8] = __obf_2a8198aab01675a0
	if __obf_1e8c1f17e8f9e9cd.MaxEntries != 0 && __obf_1e8c1f17e8f9e9cd.__obf_1726174d38a0b3b0.Len() > __obf_1e8c1f17e8f9e9cd.MaxEntries {
		__obf_1e8c1f17e8f9e9cd.RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_1e8c1f17e8f9e9cd *Lru) Get(__obf_4aecf3c737bbe5e8 string) (__obf_b99b98b3128d77a4 []byte, __obf_0560c6b8f27080cc error) {
	if __obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208 == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_2a8198aab01675a0, __obf_898cdd7de12971be := __obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208[__obf_4aecf3c737bbe5e8]; __obf_898cdd7de12971be {
		__obf_1e8c1f17e8f9e9cd.__obf_1726174d38a0b3b0.MoveToFront(__obf_2a8198aab01675a0)
		return __obf_2a8198aab01675a0.Value.(*__obf_8ca3bda6ee9c1b85).__obf_b99b98b3128d77a4, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_4aecf3c737bbe5e8)
}

// Remove removes the provided key from the cache.
func (__obf_1e8c1f17e8f9e9cd *Lru) Remove(__obf_4aecf3c737bbe5e8 string) error {
	__obf_1e8c1f17e8f9e9cd.Lock()
	defer __obf_1e8c1f17e8f9e9cd.Unlock()
	if __obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208 == nil {
		return ErrCacheEmpty
	}
	if __obf_2a8198aab01675a0, __obf_898cdd7de12971be := __obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208[__obf_4aecf3c737bbe5e8]; __obf_898cdd7de12971be {
		__obf_1e8c1f17e8f9e9cd.__obf_91180fb18d0699f5(__obf_2a8198aab01675a0)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_1e8c1f17e8f9e9cd *Lru) RemoveOldest() {
	__obf_1e8c1f17e8f9e9cd.Lock()
	defer __obf_1e8c1f17e8f9e9cd.Unlock()
	if __obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208 == nil {
		return
	}
	__obf_2a8198aab01675a0 := __obf_1e8c1f17e8f9e9cd.__obf_1726174d38a0b3b0.Back()
	if __obf_2a8198aab01675a0 != nil {
		__obf_1e8c1f17e8f9e9cd.__obf_91180fb18d0699f5(__obf_2a8198aab01675a0)
	}
}

func (__obf_1e8c1f17e8f9e9cd *Lru) __obf_91180fb18d0699f5(__obf_68f2d1855583b558 *list.Element) {
	__obf_1e8c1f17e8f9e9cd.__obf_1726174d38a0b3b0.Remove(__obf_68f2d1855583b558)
	__obf_dff7da468445ab94 := __obf_68f2d1855583b558.Value.(*__obf_8ca3bda6ee9c1b85)
	delete(__obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208, __obf_dff7da468445ab94.__obf_4aecf3c737bbe5e8)
	if __obf_1e8c1f17e8f9e9cd.OnEvicted != nil {
		__obf_1e8c1f17e8f9e9cd.OnEvicted(__obf_dff7da468445ab94.__obf_4aecf3c737bbe5e8, __obf_dff7da468445ab94.__obf_b99b98b3128d77a4)
	}
}

// Len returns the number of items in the cache.
func (__obf_1e8c1f17e8f9e9cd *Lru) Len() int {
	if __obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208 == nil {
		return 0
	}
	return __obf_1e8c1f17e8f9e9cd.__obf_1726174d38a0b3b0.Len()
}

// Clear purges all stored items from the cache.
func (__obf_1e8c1f17e8f9e9cd *Lru) Clear() error {
	__obf_1e8c1f17e8f9e9cd.Lock()
	defer __obf_1e8c1f17e8f9e9cd.Unlock()
	if __obf_1e8c1f17e8f9e9cd.OnEvicted != nil {
		for _, __obf_68f2d1855583b558 := range __obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208 {
			__obf_dff7da468445ab94 := __obf_68f2d1855583b558.Value.(*__obf_8ca3bda6ee9c1b85)
			__obf_1e8c1f17e8f9e9cd.OnEvicted(__obf_dff7da468445ab94.__obf_4aecf3c737bbe5e8, __obf_dff7da468445ab94.__obf_b99b98b3128d77a4)
		}
	}
	__obf_1e8c1f17e8f9e9cd.__obf_1726174d38a0b3b0 = nil
	__obf_1e8c1f17e8f9e9cd.__obf_62df4de078c8d208 = nil
	return nil
}
