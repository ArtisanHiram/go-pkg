package __obf_3b8640e918b7e3ff

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
	OnEvicted              func(__obf_3405c14f70aaa4d0 string, __obf_20c3560b9ea02863 any)
	__obf_399eaf03ba0d2f3b *list.List
	__obf_3b8640e918b7e3ff map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_96f16e1e0a329561 struct {
	__obf_3405c14f70aaa4d0 string
	__obf_20c3560b9ea02863 []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_a5264fe8127e11bf int) Cache {
	return &Lru{
		MaxEntries: __obf_a5264fe8127e11bf, __obf_399eaf03ba0d2f3b: list.New(), __obf_3b8640e918b7e3ff: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_5b354994d2d4bcf1 *Lru) Has(__obf_3405c14f70aaa4d0 string) (__obf_cddbfb0aefdf4145 bool) {
	__obf_5b354994d2d4bcf1.
		RLock()
	defer __obf_5b354994d2d4bcf1.RUnlock()

	_, __obf_cddbfb0aefdf4145 = __obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff[__obf_3405c14f70aaa4d0]
	return __obf_cddbfb0aefdf4145
}

// Delete key by prefix
func (__obf_5b354994d2d4bcf1 *Lru) Delete(__obf_4298344c079705aa string) {
	__obf_5b354994d2d4bcf1.
		Lock()
	defer __obf_5b354994d2d4bcf1.Unlock()
	for __obf_ebdd4526b9b834d1, __obf_fe4eccc509e2ba8c := range __obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff {
		if strings.HasPrefix(__obf_ebdd4526b9b834d1, __obf_4298344c079705aa) {
			__obf_5b354994d2d4bcf1.__obf_faee6a7da5b50044(__obf_fe4eccc509e2ba8c)
		}
	}
}

func (__obf_5b354994d2d4bcf1 *Lru) Set(__obf_3405c14f70aaa4d0 string, __obf_13458218654a7f13 any, __obf_433bae995cf2f1ae time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_5b354994d2d4bcf1 *Lru) Add(__obf_3405c14f70aaa4d0 string, __obf_20c3560b9ea02863 any) error {
	__obf_5b354994d2d4bcf1.
		Lock()
	defer __obf_5b354994d2d4bcf1.Unlock()
	if __obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff == nil {
		__obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff = make(map[string]*list.Element)
		__obf_5b354994d2d4bcf1.__obf_399eaf03ba0d2f3b = list.New()
	}
	if __obf_e4188f7be8716cf0, __obf_cddbfb0aefdf4145 := __obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff[__obf_3405c14f70aaa4d0]; __obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_399eaf03ba0d2f3b.
			MoveToFront(__obf_e4188f7be8716cf0)
		var __obf_8c6b39ef87b4061f error
		__obf_e4188f7be8716cf0.
			Value.(*__obf_96f16e1e0a329561).__obf_20c3560b9ea02863, __obf_8c6b39ef87b4061f = util.AnyToBytes(__obf_20c3560b9ea02863)
		return __obf_8c6b39ef87b4061f
	}
	__obf_13458218654a7f13, __obf_8c6b39ef87b4061f := util.AnyToBytes(__obf_20c3560b9ea02863)
	if __obf_8c6b39ef87b4061f != nil {
		return __obf_8c6b39ef87b4061f
	}
	__obf_60f2443029f7f5b5 := __obf_5b354994d2d4bcf1.__obf_399eaf03ba0d2f3b.PushFront(&__obf_96f16e1e0a329561{__obf_3405c14f70aaa4d0, __obf_13458218654a7f13})
	__obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff[__obf_3405c14f70aaa4d0] = __obf_60f2443029f7f5b5
	if __obf_5b354994d2d4bcf1.MaxEntries != 0 && __obf_5b354994d2d4bcf1.__obf_399eaf03ba0d2f3b.Len() > __obf_5b354994d2d4bcf1.MaxEntries {
		__obf_5b354994d2d4bcf1.
			RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_5b354994d2d4bcf1 *Lru) Get(__obf_3405c14f70aaa4d0 string) (__obf_20c3560b9ea02863 []byte, __obf_8c6b39ef87b4061f error) {
	if __obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_60f2443029f7f5b5, __obf_1cdcb9f2f3a09a45 := __obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff[__obf_3405c14f70aaa4d0]; __obf_1cdcb9f2f3a09a45 {
		__obf_5b354994d2d4bcf1.__obf_399eaf03ba0d2f3b.
			MoveToFront(__obf_60f2443029f7f5b5)
		return __obf_60f2443029f7f5b5.Value.(*__obf_96f16e1e0a329561).__obf_20c3560b9ea02863, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_3405c14f70aaa4d0)
}

// Remove removes the provided key from the cache.
func (__obf_5b354994d2d4bcf1 *Lru) Remove(__obf_3405c14f70aaa4d0 string) error {
	__obf_5b354994d2d4bcf1.
		Lock()
	defer __obf_5b354994d2d4bcf1.Unlock()
	if __obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff == nil {
		return ErrCacheEmpty
	}
	if __obf_60f2443029f7f5b5, __obf_1cdcb9f2f3a09a45 := __obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff[__obf_3405c14f70aaa4d0]; __obf_1cdcb9f2f3a09a45 {
		__obf_5b354994d2d4bcf1.__obf_faee6a7da5b50044(__obf_60f2443029f7f5b5)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_5b354994d2d4bcf1 *Lru) RemoveOldest() {
	__obf_5b354994d2d4bcf1.
		Lock()
	defer __obf_5b354994d2d4bcf1.Unlock()
	if __obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff == nil {
		return
	}
	__obf_60f2443029f7f5b5 := __obf_5b354994d2d4bcf1.__obf_399eaf03ba0d2f3b.Back()
	if __obf_60f2443029f7f5b5 != nil {
		__obf_5b354994d2d4bcf1.__obf_faee6a7da5b50044(__obf_60f2443029f7f5b5)
	}
}

func (__obf_5b354994d2d4bcf1 *Lru) __obf_faee6a7da5b50044(__obf_43b169a18e2edb05 *list.Element) {
	__obf_5b354994d2d4bcf1.__obf_399eaf03ba0d2f3b.
		Remove(__obf_43b169a18e2edb05)
	__obf_a5cb9a623deb9798 := __obf_43b169a18e2edb05.Value.(*__obf_96f16e1e0a329561)
	delete(__obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff, __obf_a5cb9a623deb9798.__obf_3405c14f70aaa4d0)
	if __obf_5b354994d2d4bcf1.OnEvicted != nil {
		__obf_5b354994d2d4bcf1.
			OnEvicted(__obf_a5cb9a623deb9798.__obf_3405c14f70aaa4d0,

				// Len returns the number of items in the cache.
				__obf_a5cb9a623deb9798.__obf_20c3560b9ea02863)
	}
}

func (__obf_5b354994d2d4bcf1 *Lru) Len() int {
	if __obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff == nil {
		return 0
	}
	return __obf_5b354994d2d4bcf1.

		// Clear purges all stored items from the cache.
		__obf_399eaf03ba0d2f3b.Len()
}

func (__obf_5b354994d2d4bcf1 *Lru) Clear() error {
	__obf_5b354994d2d4bcf1.
		Lock()
	defer __obf_5b354994d2d4bcf1.Unlock()
	if __obf_5b354994d2d4bcf1.OnEvicted != nil {
		for _, __obf_43b169a18e2edb05 := range __obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff {
			__obf_a5cb9a623deb9798 := __obf_43b169a18e2edb05.Value.(*__obf_96f16e1e0a329561)
			__obf_5b354994d2d4bcf1.
				OnEvicted(__obf_a5cb9a623deb9798.__obf_3405c14f70aaa4d0, __obf_a5cb9a623deb9798.__obf_20c3560b9ea02863)
		}
	}
	__obf_5b354994d2d4bcf1.__obf_399eaf03ba0d2f3b = nil
	__obf_5b354994d2d4bcf1.__obf_3b8640e918b7e3ff = nil
	return nil
}
