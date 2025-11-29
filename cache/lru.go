package __obf_a05682be1c6caf18

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
	OnEvicted              func(__obf_a0e56915a05e8d99 string, __obf_106924f0283e3cd8 any)
	__obf_a168c4f57e6d2b85 *list.List
	__obf_a05682be1c6caf18 map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_a39ae807d5334dfc struct {
	__obf_a0e56915a05e8d99 string
	__obf_106924f0283e3cd8 []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_b7c87f601bc04ebb int) Cache {
	return &Lru{
		MaxEntries: __obf_b7c87f601bc04ebb, __obf_a168c4f57e6d2b85: list.New(), __obf_a05682be1c6caf18: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_365ba253fb24199c *Lru) Has(__obf_a0e56915a05e8d99 string) (__obf_19211ff80aaad9b4 bool) {
	__obf_365ba253fb24199c.
		RLock()
	defer __obf_365ba253fb24199c.RUnlock()

	_, __obf_19211ff80aaad9b4 = __obf_365ba253fb24199c.__obf_a05682be1c6caf18[__obf_a0e56915a05e8d99]
	return __obf_19211ff80aaad9b4
}

// Delete key by prefix
func (__obf_365ba253fb24199c *Lru) Delete(__obf_715bf114d0b286fb string) {
	__obf_365ba253fb24199c.
		Lock()
	defer __obf_365ba253fb24199c.Unlock()
	for __obf_b0c62418c29009a4, __obf_40394ecac45b9c79 := range __obf_365ba253fb24199c.__obf_a05682be1c6caf18 {
		if strings.HasPrefix(__obf_b0c62418c29009a4, __obf_715bf114d0b286fb) {
			__obf_365ba253fb24199c.__obf_53bc00c565921611(__obf_40394ecac45b9c79)
		}
	}
}

func (__obf_365ba253fb24199c *Lru) Set(__obf_a0e56915a05e8d99 string, __obf_6030b68d8a95172f any, __obf_c8947601a564de1b time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_365ba253fb24199c *Lru) Add(__obf_a0e56915a05e8d99 string, __obf_106924f0283e3cd8 any) error {
	__obf_365ba253fb24199c.
		Lock()
	defer __obf_365ba253fb24199c.Unlock()
	if __obf_365ba253fb24199c.__obf_a05682be1c6caf18 == nil {
		__obf_365ba253fb24199c.__obf_a05682be1c6caf18 = make(map[string]*list.Element)
		__obf_365ba253fb24199c.__obf_a168c4f57e6d2b85 = list.New()
	}
	if __obf_b4ba67c7cc1659de, __obf_19211ff80aaad9b4 := __obf_365ba253fb24199c.__obf_a05682be1c6caf18[__obf_a0e56915a05e8d99]; __obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_a168c4f57e6d2b85.
			MoveToFront(__obf_b4ba67c7cc1659de)
		var __obf_94de95cb433b50f1 error
		__obf_b4ba67c7cc1659de.
			Value.(*__obf_a39ae807d5334dfc).__obf_106924f0283e3cd8, __obf_94de95cb433b50f1 = util.AnyToBytes(__obf_106924f0283e3cd8)
		return __obf_94de95cb433b50f1
	}
	__obf_6030b68d8a95172f, __obf_94de95cb433b50f1 := util.AnyToBytes(__obf_106924f0283e3cd8)
	if __obf_94de95cb433b50f1 != nil {
		return __obf_94de95cb433b50f1
	}
	__obf_9980e6db3383b4b2 := __obf_365ba253fb24199c.__obf_a168c4f57e6d2b85.PushFront(&__obf_a39ae807d5334dfc{__obf_a0e56915a05e8d99, __obf_6030b68d8a95172f})
	__obf_365ba253fb24199c.__obf_a05682be1c6caf18[__obf_a0e56915a05e8d99] = __obf_9980e6db3383b4b2
	if __obf_365ba253fb24199c.MaxEntries != 0 && __obf_365ba253fb24199c.__obf_a168c4f57e6d2b85.Len() > __obf_365ba253fb24199c.MaxEntries {
		__obf_365ba253fb24199c.
			RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_365ba253fb24199c *Lru) Get(__obf_a0e56915a05e8d99 string) (__obf_106924f0283e3cd8 []byte, __obf_94de95cb433b50f1 error) {
	if __obf_365ba253fb24199c.__obf_a05682be1c6caf18 == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_9980e6db3383b4b2, __obf_39a769536fddae1b := __obf_365ba253fb24199c.__obf_a05682be1c6caf18[__obf_a0e56915a05e8d99]; __obf_39a769536fddae1b {
		__obf_365ba253fb24199c.__obf_a168c4f57e6d2b85.
			MoveToFront(__obf_9980e6db3383b4b2)
		return __obf_9980e6db3383b4b2.Value.(*__obf_a39ae807d5334dfc).__obf_106924f0283e3cd8, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_a0e56915a05e8d99)
}

// Remove removes the provided key from the cache.
func (__obf_365ba253fb24199c *Lru) Remove(__obf_a0e56915a05e8d99 string) error {
	__obf_365ba253fb24199c.
		Lock()
	defer __obf_365ba253fb24199c.Unlock()
	if __obf_365ba253fb24199c.__obf_a05682be1c6caf18 == nil {
		return ErrCacheEmpty
	}
	if __obf_9980e6db3383b4b2, __obf_39a769536fddae1b := __obf_365ba253fb24199c.__obf_a05682be1c6caf18[__obf_a0e56915a05e8d99]; __obf_39a769536fddae1b {
		__obf_365ba253fb24199c.__obf_53bc00c565921611(__obf_9980e6db3383b4b2)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_365ba253fb24199c *Lru) RemoveOldest() {
	__obf_365ba253fb24199c.
		Lock()
	defer __obf_365ba253fb24199c.Unlock()
	if __obf_365ba253fb24199c.__obf_a05682be1c6caf18 == nil {
		return
	}
	__obf_9980e6db3383b4b2 := __obf_365ba253fb24199c.__obf_a168c4f57e6d2b85.Back()
	if __obf_9980e6db3383b4b2 != nil {
		__obf_365ba253fb24199c.__obf_53bc00c565921611(__obf_9980e6db3383b4b2)
	}
}

func (__obf_365ba253fb24199c *Lru) __obf_53bc00c565921611(__obf_1086d0f8dfd43ca1 *list.Element) {
	__obf_365ba253fb24199c.__obf_a168c4f57e6d2b85.
		Remove(__obf_1086d0f8dfd43ca1)
	__obf_aaf429529a90d259 := __obf_1086d0f8dfd43ca1.Value.(*__obf_a39ae807d5334dfc)
	delete(__obf_365ba253fb24199c.__obf_a05682be1c6caf18, __obf_aaf429529a90d259.__obf_a0e56915a05e8d99)
	if __obf_365ba253fb24199c.OnEvicted != nil {
		__obf_365ba253fb24199c.
			OnEvicted(__obf_aaf429529a90d259.__obf_a0e56915a05e8d99,

				// Len returns the number of items in the cache.
				__obf_aaf429529a90d259.__obf_106924f0283e3cd8)
	}
}

func (__obf_365ba253fb24199c *Lru) Len() int {
	if __obf_365ba253fb24199c.__obf_a05682be1c6caf18 == nil {
		return 0
	}
	return __obf_365ba253fb24199c.

		// Clear purges all stored items from the cache.
		__obf_a168c4f57e6d2b85.Len()
}

func (__obf_365ba253fb24199c *Lru) Clear() error {
	__obf_365ba253fb24199c.
		Lock()
	defer __obf_365ba253fb24199c.Unlock()
	if __obf_365ba253fb24199c.OnEvicted != nil {
		for _, __obf_1086d0f8dfd43ca1 := range __obf_365ba253fb24199c.__obf_a05682be1c6caf18 {
			__obf_aaf429529a90d259 := __obf_1086d0f8dfd43ca1.Value.(*__obf_a39ae807d5334dfc)
			__obf_365ba253fb24199c.
				OnEvicted(__obf_aaf429529a90d259.__obf_a0e56915a05e8d99, __obf_aaf429529a90d259.__obf_106924f0283e3cd8)
		}
	}
	__obf_365ba253fb24199c.__obf_a168c4f57e6d2b85 = nil
	__obf_365ba253fb24199c.__obf_a05682be1c6caf18 = nil
	return nil
}
