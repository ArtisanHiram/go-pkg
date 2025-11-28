package __obf_32d927a1e06b7e71

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
	OnEvicted func(__obf_13113b328a6972ad string, __obf_b186e1095a1e7f1f any)

	__obf_df773dfec3c5d5dc *list.List
	__obf_32d927a1e06b7e71 map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_809b8fb1de2a7591 struct {
	__obf_13113b328a6972ad string
	__obf_b186e1095a1e7f1f []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_31c4ad7ded8e233a int) Cache {
	return &Lru{
		MaxEntries:             __obf_31c4ad7ded8e233a,
		__obf_df773dfec3c5d5dc: list.New(),
		__obf_32d927a1e06b7e71: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_443bb096afb82210 *Lru) Has(__obf_13113b328a6972ad string) (__obf_542c0401f22c9aec bool) {
	__obf_443bb096afb82210.RLock()
	defer __obf_443bb096afb82210.RUnlock()

	_, __obf_542c0401f22c9aec = __obf_443bb096afb82210.__obf_32d927a1e06b7e71[__obf_13113b328a6972ad]
	return __obf_542c0401f22c9aec
}

// Delete key by prefix
func (__obf_443bb096afb82210 *Lru) Delete(__obf_e5c9d9e65a3fa384 string) {
	__obf_443bb096afb82210.Lock()
	defer __obf_443bb096afb82210.Unlock()
	for __obf_4be74b34d9a5767c, __obf_cf59be873d0aaa34 := range __obf_443bb096afb82210.__obf_32d927a1e06b7e71 {
		if strings.HasPrefix(__obf_4be74b34d9a5767c, __obf_e5c9d9e65a3fa384) {
			__obf_443bb096afb82210.__obf_8f20e2ad394dc1b9(__obf_cf59be873d0aaa34)
		}
	}
}

func (__obf_443bb096afb82210 *Lru) Set(__obf_13113b328a6972ad string, __obf_a967d4ebf1531f4c any, __obf_481cbd2caaded2ed time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_443bb096afb82210 *Lru) Add(__obf_13113b328a6972ad string, __obf_b186e1095a1e7f1f any) error {
	__obf_443bb096afb82210.Lock()
	defer __obf_443bb096afb82210.Unlock()
	if __obf_443bb096afb82210.__obf_32d927a1e06b7e71 == nil {
		__obf_443bb096afb82210.__obf_32d927a1e06b7e71 = make(map[string]*list.Element)
		__obf_443bb096afb82210.__obf_df773dfec3c5d5dc = list.New()
	}
	if __obf_e9e779f3438d91dc, __obf_542c0401f22c9aec := __obf_443bb096afb82210.__obf_32d927a1e06b7e71[__obf_13113b328a6972ad]; __obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_df773dfec3c5d5dc.MoveToFront(__obf_e9e779f3438d91dc)
		var __obf_19fb5c4c25ff452a error
		__obf_e9e779f3438d91dc.Value.(*__obf_809b8fb1de2a7591).__obf_b186e1095a1e7f1f, __obf_19fb5c4c25ff452a = util.AnyToBytes(__obf_b186e1095a1e7f1f)
		return __obf_19fb5c4c25ff452a
	}
	__obf_a967d4ebf1531f4c, __obf_19fb5c4c25ff452a := util.AnyToBytes(__obf_b186e1095a1e7f1f)
	if __obf_19fb5c4c25ff452a != nil {
		return __obf_19fb5c4c25ff452a
	}
	__obf_6baf4c165dd0c90e := __obf_443bb096afb82210.__obf_df773dfec3c5d5dc.PushFront(&__obf_809b8fb1de2a7591{__obf_13113b328a6972ad, __obf_a967d4ebf1531f4c})
	__obf_443bb096afb82210.__obf_32d927a1e06b7e71[__obf_13113b328a6972ad] = __obf_6baf4c165dd0c90e
	if __obf_443bb096afb82210.MaxEntries != 0 && __obf_443bb096afb82210.__obf_df773dfec3c5d5dc.Len() > __obf_443bb096afb82210.MaxEntries {
		__obf_443bb096afb82210.RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_443bb096afb82210 *Lru) Get(__obf_13113b328a6972ad string) (__obf_b186e1095a1e7f1f []byte, __obf_19fb5c4c25ff452a error) {
	if __obf_443bb096afb82210.__obf_32d927a1e06b7e71 == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_6baf4c165dd0c90e, __obf_bcc7bf7ca8534fba := __obf_443bb096afb82210.__obf_32d927a1e06b7e71[__obf_13113b328a6972ad]; __obf_bcc7bf7ca8534fba {
		__obf_443bb096afb82210.__obf_df773dfec3c5d5dc.MoveToFront(__obf_6baf4c165dd0c90e)
		return __obf_6baf4c165dd0c90e.Value.(*__obf_809b8fb1de2a7591).__obf_b186e1095a1e7f1f, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_13113b328a6972ad)
}

// Remove removes the provided key from the cache.
func (__obf_443bb096afb82210 *Lru) Remove(__obf_13113b328a6972ad string) error {
	__obf_443bb096afb82210.Lock()
	defer __obf_443bb096afb82210.Unlock()
	if __obf_443bb096afb82210.__obf_32d927a1e06b7e71 == nil {
		return ErrCacheEmpty
	}
	if __obf_6baf4c165dd0c90e, __obf_bcc7bf7ca8534fba := __obf_443bb096afb82210.__obf_32d927a1e06b7e71[__obf_13113b328a6972ad]; __obf_bcc7bf7ca8534fba {
		__obf_443bb096afb82210.__obf_8f20e2ad394dc1b9(__obf_6baf4c165dd0c90e)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_443bb096afb82210 *Lru) RemoveOldest() {
	__obf_443bb096afb82210.Lock()
	defer __obf_443bb096afb82210.Unlock()
	if __obf_443bb096afb82210.__obf_32d927a1e06b7e71 == nil {
		return
	}
	__obf_6baf4c165dd0c90e := __obf_443bb096afb82210.__obf_df773dfec3c5d5dc.Back()
	if __obf_6baf4c165dd0c90e != nil {
		__obf_443bb096afb82210.__obf_8f20e2ad394dc1b9(__obf_6baf4c165dd0c90e)
	}
}

func (__obf_443bb096afb82210 *Lru) __obf_8f20e2ad394dc1b9(__obf_b9653b0201c59d0c *list.Element) {
	__obf_443bb096afb82210.__obf_df773dfec3c5d5dc.Remove(__obf_b9653b0201c59d0c)
	__obf_6df53d6ceb80c58c := __obf_b9653b0201c59d0c.Value.(*__obf_809b8fb1de2a7591)
	delete(__obf_443bb096afb82210.__obf_32d927a1e06b7e71, __obf_6df53d6ceb80c58c.__obf_13113b328a6972ad)
	if __obf_443bb096afb82210.OnEvicted != nil {
		__obf_443bb096afb82210.OnEvicted(__obf_6df53d6ceb80c58c.__obf_13113b328a6972ad, __obf_6df53d6ceb80c58c.__obf_b186e1095a1e7f1f)
	}
}

// Len returns the number of items in the cache.
func (__obf_443bb096afb82210 *Lru) Len() int {
	if __obf_443bb096afb82210.__obf_32d927a1e06b7e71 == nil {
		return 0
	}
	return __obf_443bb096afb82210.__obf_df773dfec3c5d5dc.Len()
}

// Clear purges all stored items from the cache.
func (__obf_443bb096afb82210 *Lru) Clear() error {
	__obf_443bb096afb82210.Lock()
	defer __obf_443bb096afb82210.Unlock()
	if __obf_443bb096afb82210.OnEvicted != nil {
		for _, __obf_b9653b0201c59d0c := range __obf_443bb096afb82210.__obf_32d927a1e06b7e71 {
			__obf_6df53d6ceb80c58c := __obf_b9653b0201c59d0c.Value.(*__obf_809b8fb1de2a7591)
			__obf_443bb096afb82210.OnEvicted(__obf_6df53d6ceb80c58c.__obf_13113b328a6972ad, __obf_6df53d6ceb80c58c.__obf_b186e1095a1e7f1f)
		}
	}
	__obf_443bb096afb82210.__obf_df773dfec3c5d5dc = nil
	__obf_443bb096afb82210.__obf_32d927a1e06b7e71 = nil
	return nil
}
