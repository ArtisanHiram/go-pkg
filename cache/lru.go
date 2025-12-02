package __obf_6b306490bf7221d3

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
	OnEvicted              func(__obf_fa3a380c35ada5d9 string, __obf_c2b0e78296a7a084 any)
	__obf_d4fb3b985ed0a167 *list.List
	__obf_6b306490bf7221d3 map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_35bbaec790d3686a struct {
	__obf_fa3a380c35ada5d9 string
	__obf_c2b0e78296a7a084 []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_b4d4509cf8a0fcb3 int) Cache {
	return &Lru{
		MaxEntries: __obf_b4d4509cf8a0fcb3, __obf_d4fb3b985ed0a167: list.New(), __obf_6b306490bf7221d3: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_6ccac4bb8646a806 *Lru) Has(__obf_fa3a380c35ada5d9 string) (__obf_8761a3e94b6603ac bool) {
	__obf_6ccac4bb8646a806.
		RLock()
	defer __obf_6ccac4bb8646a806.RUnlock()

	_, __obf_8761a3e94b6603ac = __obf_6ccac4bb8646a806.__obf_6b306490bf7221d3[__obf_fa3a380c35ada5d9]
	return __obf_8761a3e94b6603ac
}

// Delete key by prefix
func (__obf_6ccac4bb8646a806 *Lru) Delete(__obf_7e6d8e57607bcb1d string) {
	__obf_6ccac4bb8646a806.
		Lock()
	defer __obf_6ccac4bb8646a806.Unlock()
	for __obf_afff83fa0e852ce0, __obf_6208501794a51081 := range __obf_6ccac4bb8646a806.__obf_6b306490bf7221d3 {
		if strings.HasPrefix(__obf_afff83fa0e852ce0, __obf_7e6d8e57607bcb1d) {
			__obf_6ccac4bb8646a806.__obf_28ff5fe49401f5fd(__obf_6208501794a51081)
		}
	}
}

func (__obf_6ccac4bb8646a806 *Lru) Set(__obf_fa3a380c35ada5d9 string, __obf_17a4af14df8bae4f any, __obf_a15e4803e6944fc5 time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_6ccac4bb8646a806 *Lru) Add(__obf_fa3a380c35ada5d9 string, __obf_c2b0e78296a7a084 any) error {
	__obf_6ccac4bb8646a806.
		Lock()
	defer __obf_6ccac4bb8646a806.Unlock()
	if __obf_6ccac4bb8646a806.__obf_6b306490bf7221d3 == nil {
		__obf_6ccac4bb8646a806.__obf_6b306490bf7221d3 = make(map[string]*list.Element)
		__obf_6ccac4bb8646a806.__obf_d4fb3b985ed0a167 = list.New()
	}
	if __obf_0655f271d7b7a057, __obf_8761a3e94b6603ac := __obf_6ccac4bb8646a806.__obf_6b306490bf7221d3[__obf_fa3a380c35ada5d9]; __obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_d4fb3b985ed0a167.
			MoveToFront(__obf_0655f271d7b7a057)
		var __obf_9ffdd5bbb9f1dc61 error
		__obf_0655f271d7b7a057.
			Value.(*__obf_35bbaec790d3686a).__obf_c2b0e78296a7a084, __obf_9ffdd5bbb9f1dc61 = util.AnyToBytes(__obf_c2b0e78296a7a084)
		return __obf_9ffdd5bbb9f1dc61
	}
	__obf_17a4af14df8bae4f, __obf_9ffdd5bbb9f1dc61 := util.AnyToBytes(__obf_c2b0e78296a7a084)
	if __obf_9ffdd5bbb9f1dc61 != nil {
		return __obf_9ffdd5bbb9f1dc61
	}
	__obf_25155d4b8607d876 := __obf_6ccac4bb8646a806.__obf_d4fb3b985ed0a167.PushFront(&__obf_35bbaec790d3686a{__obf_fa3a380c35ada5d9, __obf_17a4af14df8bae4f})
	__obf_6ccac4bb8646a806.__obf_6b306490bf7221d3[__obf_fa3a380c35ada5d9] = __obf_25155d4b8607d876
	if __obf_6ccac4bb8646a806.MaxEntries != 0 && __obf_6ccac4bb8646a806.__obf_d4fb3b985ed0a167.Len() > __obf_6ccac4bb8646a806.MaxEntries {
		__obf_6ccac4bb8646a806.
			RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_6ccac4bb8646a806 *Lru) Get(__obf_fa3a380c35ada5d9 string) (__obf_c2b0e78296a7a084 []byte, __obf_9ffdd5bbb9f1dc61 error) {
	if __obf_6ccac4bb8646a806.__obf_6b306490bf7221d3 == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_25155d4b8607d876, __obf_807cd92cadf68aba := __obf_6ccac4bb8646a806.__obf_6b306490bf7221d3[__obf_fa3a380c35ada5d9]; __obf_807cd92cadf68aba {
		__obf_6ccac4bb8646a806.__obf_d4fb3b985ed0a167.
			MoveToFront(__obf_25155d4b8607d876)
		return __obf_25155d4b8607d876.Value.(*__obf_35bbaec790d3686a).__obf_c2b0e78296a7a084, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_fa3a380c35ada5d9)
}

// Remove removes the provided key from the cache.
func (__obf_6ccac4bb8646a806 *Lru) Remove(__obf_fa3a380c35ada5d9 string) error {
	__obf_6ccac4bb8646a806.
		Lock()
	defer __obf_6ccac4bb8646a806.Unlock()
	if __obf_6ccac4bb8646a806.__obf_6b306490bf7221d3 == nil {
		return ErrCacheEmpty
	}
	if __obf_25155d4b8607d876, __obf_807cd92cadf68aba := __obf_6ccac4bb8646a806.__obf_6b306490bf7221d3[__obf_fa3a380c35ada5d9]; __obf_807cd92cadf68aba {
		__obf_6ccac4bb8646a806.__obf_28ff5fe49401f5fd(__obf_25155d4b8607d876)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_6ccac4bb8646a806 *Lru) RemoveOldest() {
	__obf_6ccac4bb8646a806.
		Lock()
	defer __obf_6ccac4bb8646a806.Unlock()
	if __obf_6ccac4bb8646a806.__obf_6b306490bf7221d3 == nil {
		return
	}
	__obf_25155d4b8607d876 := __obf_6ccac4bb8646a806.__obf_d4fb3b985ed0a167.Back()
	if __obf_25155d4b8607d876 != nil {
		__obf_6ccac4bb8646a806.__obf_28ff5fe49401f5fd(__obf_25155d4b8607d876)
	}
}

func (__obf_6ccac4bb8646a806 *Lru) __obf_28ff5fe49401f5fd(__obf_be37eed188feee9f *list.Element) {
	__obf_6ccac4bb8646a806.__obf_d4fb3b985ed0a167.
		Remove(__obf_be37eed188feee9f)
	__obf_eda25f326b28841b := __obf_be37eed188feee9f.Value.(*__obf_35bbaec790d3686a)
	delete(__obf_6ccac4bb8646a806.__obf_6b306490bf7221d3, __obf_eda25f326b28841b.__obf_fa3a380c35ada5d9)
	if __obf_6ccac4bb8646a806.OnEvicted != nil {
		__obf_6ccac4bb8646a806.
			OnEvicted(__obf_eda25f326b28841b.__obf_fa3a380c35ada5d9,

				// Len returns the number of items in the cache.
				__obf_eda25f326b28841b.__obf_c2b0e78296a7a084)
	}
}

func (__obf_6ccac4bb8646a806 *Lru) Len() int {
	if __obf_6ccac4bb8646a806.__obf_6b306490bf7221d3 == nil {
		return 0
	}
	return __obf_6ccac4bb8646a806.

		// Clear purges all stored items from the cache.
		__obf_d4fb3b985ed0a167.Len()
}

func (__obf_6ccac4bb8646a806 *Lru) Clear() error {
	__obf_6ccac4bb8646a806.
		Lock()
	defer __obf_6ccac4bb8646a806.Unlock()
	if __obf_6ccac4bb8646a806.OnEvicted != nil {
		for _, __obf_be37eed188feee9f := range __obf_6ccac4bb8646a806.__obf_6b306490bf7221d3 {
			__obf_eda25f326b28841b := __obf_be37eed188feee9f.Value.(*__obf_35bbaec790d3686a)
			__obf_6ccac4bb8646a806.
				OnEvicted(__obf_eda25f326b28841b.__obf_fa3a380c35ada5d9, __obf_eda25f326b28841b.__obf_c2b0e78296a7a084)
		}
	}
	__obf_6ccac4bb8646a806.__obf_d4fb3b985ed0a167 = nil
	__obf_6ccac4bb8646a806.__obf_6b306490bf7221d3 = nil
	return nil
}
