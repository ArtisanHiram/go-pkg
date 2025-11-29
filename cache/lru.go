package __obf_6fd4fe34e3f784df

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
	OnEvicted              func(__obf_ca986a38d1f8fbbb string, __obf_e35c6f9510020aab any)
	__obf_9e593b9cad5acad6 *list.List
	__obf_6fd4fe34e3f784df map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_d34c1cc41e44647a struct {
	__obf_ca986a38d1f8fbbb string
	__obf_e35c6f9510020aab []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_b8db617683455a34 int) Cache {
	return &Lru{
		MaxEntries: __obf_b8db617683455a34, __obf_9e593b9cad5acad6: list.New(), __obf_6fd4fe34e3f784df: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_3cd39772de5fd3ee *Lru) Has(__obf_ca986a38d1f8fbbb string) (__obf_949c761ffd90d5fa bool) {
	__obf_3cd39772de5fd3ee.
		RLock()
	defer __obf_3cd39772de5fd3ee.RUnlock()

	_, __obf_949c761ffd90d5fa = __obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df[__obf_ca986a38d1f8fbbb]
	return __obf_949c761ffd90d5fa
}

// Delete key by prefix
func (__obf_3cd39772de5fd3ee *Lru) Delete(__obf_c53a458d6b02aeb9 string) {
	__obf_3cd39772de5fd3ee.
		Lock()
	defer __obf_3cd39772de5fd3ee.Unlock()
	for __obf_2a0493c028b94053, __obf_b1ecf81133de0ff6 := range __obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df {
		if strings.HasPrefix(__obf_2a0493c028b94053, __obf_c53a458d6b02aeb9) {
			__obf_3cd39772de5fd3ee.__obf_6a29b44d819fc0b4(__obf_b1ecf81133de0ff6)
		}
	}
}

func (__obf_3cd39772de5fd3ee *Lru) Set(__obf_ca986a38d1f8fbbb string, __obf_e1d0258259217620 any, __obf_bfbeaaea582ee60a time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_3cd39772de5fd3ee *Lru) Add(__obf_ca986a38d1f8fbbb string, __obf_e35c6f9510020aab any) error {
	__obf_3cd39772de5fd3ee.
		Lock()
	defer __obf_3cd39772de5fd3ee.Unlock()
	if __obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df == nil {
		__obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df = make(map[string]*list.Element)
		__obf_3cd39772de5fd3ee.__obf_9e593b9cad5acad6 = list.New()
	}
	if __obf_4a29e522ccf1492e, __obf_949c761ffd90d5fa := __obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df[__obf_ca986a38d1f8fbbb]; __obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_9e593b9cad5acad6.
			MoveToFront(__obf_4a29e522ccf1492e)
		var __obf_7f8f5f5c173fa58e error
		__obf_4a29e522ccf1492e.
			Value.(*__obf_d34c1cc41e44647a).__obf_e35c6f9510020aab, __obf_7f8f5f5c173fa58e = util.AnyToBytes(__obf_e35c6f9510020aab)
		return __obf_7f8f5f5c173fa58e
	}
	__obf_e1d0258259217620, __obf_7f8f5f5c173fa58e := util.AnyToBytes(__obf_e35c6f9510020aab)
	if __obf_7f8f5f5c173fa58e != nil {
		return __obf_7f8f5f5c173fa58e
	}
	__obf_0e1e569a5ebcbb0c := __obf_3cd39772de5fd3ee.__obf_9e593b9cad5acad6.PushFront(&__obf_d34c1cc41e44647a{__obf_ca986a38d1f8fbbb, __obf_e1d0258259217620})
	__obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df[__obf_ca986a38d1f8fbbb] = __obf_0e1e569a5ebcbb0c
	if __obf_3cd39772de5fd3ee.MaxEntries != 0 && __obf_3cd39772de5fd3ee.__obf_9e593b9cad5acad6.Len() > __obf_3cd39772de5fd3ee.MaxEntries {
		__obf_3cd39772de5fd3ee.
			RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_3cd39772de5fd3ee *Lru) Get(__obf_ca986a38d1f8fbbb string) (__obf_e35c6f9510020aab []byte, __obf_7f8f5f5c173fa58e error) {
	if __obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_0e1e569a5ebcbb0c, __obf_5ead06b5d7f835a2 := __obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df[__obf_ca986a38d1f8fbbb]; __obf_5ead06b5d7f835a2 {
		__obf_3cd39772de5fd3ee.__obf_9e593b9cad5acad6.
			MoveToFront(__obf_0e1e569a5ebcbb0c)
		return __obf_0e1e569a5ebcbb0c.Value.(*__obf_d34c1cc41e44647a).__obf_e35c6f9510020aab, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_ca986a38d1f8fbbb)
}

// Remove removes the provided key from the cache.
func (__obf_3cd39772de5fd3ee *Lru) Remove(__obf_ca986a38d1f8fbbb string) error {
	__obf_3cd39772de5fd3ee.
		Lock()
	defer __obf_3cd39772de5fd3ee.Unlock()
	if __obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df == nil {
		return ErrCacheEmpty
	}
	if __obf_0e1e569a5ebcbb0c, __obf_5ead06b5d7f835a2 := __obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df[__obf_ca986a38d1f8fbbb]; __obf_5ead06b5d7f835a2 {
		__obf_3cd39772de5fd3ee.__obf_6a29b44d819fc0b4(__obf_0e1e569a5ebcbb0c)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_3cd39772de5fd3ee *Lru) RemoveOldest() {
	__obf_3cd39772de5fd3ee.
		Lock()
	defer __obf_3cd39772de5fd3ee.Unlock()
	if __obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df == nil {
		return
	}
	__obf_0e1e569a5ebcbb0c := __obf_3cd39772de5fd3ee.__obf_9e593b9cad5acad6.Back()
	if __obf_0e1e569a5ebcbb0c != nil {
		__obf_3cd39772de5fd3ee.__obf_6a29b44d819fc0b4(__obf_0e1e569a5ebcbb0c)
	}
}

func (__obf_3cd39772de5fd3ee *Lru) __obf_6a29b44d819fc0b4(__obf_0f22edc65e150c28 *list.Element) {
	__obf_3cd39772de5fd3ee.__obf_9e593b9cad5acad6.
		Remove(__obf_0f22edc65e150c28)
	__obf_10e482d747d89898 := __obf_0f22edc65e150c28.Value.(*__obf_d34c1cc41e44647a)
	delete(__obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df, __obf_10e482d747d89898.__obf_ca986a38d1f8fbbb)
	if __obf_3cd39772de5fd3ee.OnEvicted != nil {
		__obf_3cd39772de5fd3ee.
			OnEvicted(__obf_10e482d747d89898.__obf_ca986a38d1f8fbbb,

				// Len returns the number of items in the cache.
				__obf_10e482d747d89898.__obf_e35c6f9510020aab)
	}
}

func (__obf_3cd39772de5fd3ee *Lru) Len() int {
	if __obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df == nil {
		return 0
	}
	return __obf_3cd39772de5fd3ee.

		// Clear purges all stored items from the cache.
		__obf_9e593b9cad5acad6.Len()
}

func (__obf_3cd39772de5fd3ee *Lru) Clear() error {
	__obf_3cd39772de5fd3ee.
		Lock()
	defer __obf_3cd39772de5fd3ee.Unlock()
	if __obf_3cd39772de5fd3ee.OnEvicted != nil {
		for _, __obf_0f22edc65e150c28 := range __obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df {
			__obf_10e482d747d89898 := __obf_0f22edc65e150c28.Value.(*__obf_d34c1cc41e44647a)
			__obf_3cd39772de5fd3ee.
				OnEvicted(__obf_10e482d747d89898.__obf_ca986a38d1f8fbbb, __obf_10e482d747d89898.__obf_e35c6f9510020aab)
		}
	}
	__obf_3cd39772de5fd3ee.__obf_9e593b9cad5acad6 = nil
	__obf_3cd39772de5fd3ee.__obf_6fd4fe34e3f784df = nil
	return nil
}
