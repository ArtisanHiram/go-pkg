package __obf_72d962f6a40bc95f

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
	OnEvicted              func(__obf_73cb8148cbf55699 string, __obf_82bf2d5ce3534205 any)
	__obf_238b932d03e3f451 *list.List
	__obf_72d962f6a40bc95f map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_953b8b04400c10db struct {
	__obf_73cb8148cbf55699 string
	__obf_82bf2d5ce3534205 []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_9b9f5d2ea74e99ff int) Cache {
	return &Lru{
		MaxEntries: __obf_9b9f5d2ea74e99ff, __obf_238b932d03e3f451: list.New(), __obf_72d962f6a40bc95f: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_89bd5bf0a2a3aefa *Lru) Has(__obf_73cb8148cbf55699 string) (__obf_97faa3bbfe732802 bool) {
	__obf_89bd5bf0a2a3aefa.
		RLock()
	defer __obf_89bd5bf0a2a3aefa.RUnlock()

	_, __obf_97faa3bbfe732802 = __obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f[__obf_73cb8148cbf55699]
	return __obf_97faa3bbfe732802
}

// Delete key by prefix
func (__obf_89bd5bf0a2a3aefa *Lru) Delete(__obf_939173229b00ad97 string) {
	__obf_89bd5bf0a2a3aefa.
		Lock()
	defer __obf_89bd5bf0a2a3aefa.Unlock()
	for __obf_41bce9a58f43c231, __obf_b33efcc537be1290 := range __obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f {
		if strings.HasPrefix(__obf_41bce9a58f43c231, __obf_939173229b00ad97) {
			__obf_89bd5bf0a2a3aefa.__obf_651948667f8af266(__obf_b33efcc537be1290)
		}
	}
}

func (__obf_89bd5bf0a2a3aefa *Lru) Set(__obf_73cb8148cbf55699 string, __obf_1dbdf9c3e13789b0 any, __obf_c6c24327f124e623 time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_89bd5bf0a2a3aefa *Lru) Add(__obf_73cb8148cbf55699 string, __obf_82bf2d5ce3534205 any) error {
	__obf_89bd5bf0a2a3aefa.
		Lock()
	defer __obf_89bd5bf0a2a3aefa.Unlock()
	if __obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f == nil {
		__obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f = make(map[string]*list.Element)
		__obf_89bd5bf0a2a3aefa.__obf_238b932d03e3f451 = list.New()
	}
	if __obf_531948f208cbec12, __obf_97faa3bbfe732802 := __obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f[__obf_73cb8148cbf55699]; __obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_238b932d03e3f451.
			MoveToFront(__obf_531948f208cbec12)
		var __obf_2498adaec5f4c8f1 error
		__obf_531948f208cbec12.
			Value.(*__obf_953b8b04400c10db).__obf_82bf2d5ce3534205, __obf_2498adaec5f4c8f1 = util.AnyToBytes(__obf_82bf2d5ce3534205)
		return __obf_2498adaec5f4c8f1
	}
	__obf_1dbdf9c3e13789b0, __obf_2498adaec5f4c8f1 := util.AnyToBytes(__obf_82bf2d5ce3534205)
	if __obf_2498adaec5f4c8f1 != nil {
		return __obf_2498adaec5f4c8f1
	}
	__obf_e575c7c6fa901909 := __obf_89bd5bf0a2a3aefa.__obf_238b932d03e3f451.PushFront(&__obf_953b8b04400c10db{__obf_73cb8148cbf55699, __obf_1dbdf9c3e13789b0})
	__obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f[__obf_73cb8148cbf55699] = __obf_e575c7c6fa901909
	if __obf_89bd5bf0a2a3aefa.MaxEntries != 0 && __obf_89bd5bf0a2a3aefa.__obf_238b932d03e3f451.Len() > __obf_89bd5bf0a2a3aefa.MaxEntries {
		__obf_89bd5bf0a2a3aefa.
			RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_89bd5bf0a2a3aefa *Lru) Get(__obf_73cb8148cbf55699 string) (__obf_82bf2d5ce3534205 []byte, __obf_2498adaec5f4c8f1 error) {
	if __obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_e575c7c6fa901909, __obf_64562781d198733c := __obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f[__obf_73cb8148cbf55699]; __obf_64562781d198733c {
		__obf_89bd5bf0a2a3aefa.__obf_238b932d03e3f451.
			MoveToFront(__obf_e575c7c6fa901909)
		return __obf_e575c7c6fa901909.Value.(*__obf_953b8b04400c10db).__obf_82bf2d5ce3534205, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_73cb8148cbf55699)
}

// Remove removes the provided key from the cache.
func (__obf_89bd5bf0a2a3aefa *Lru) Remove(__obf_73cb8148cbf55699 string) error {
	__obf_89bd5bf0a2a3aefa.
		Lock()
	defer __obf_89bd5bf0a2a3aefa.Unlock()
	if __obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f == nil {
		return ErrCacheEmpty
	}
	if __obf_e575c7c6fa901909, __obf_64562781d198733c := __obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f[__obf_73cb8148cbf55699]; __obf_64562781d198733c {
		__obf_89bd5bf0a2a3aefa.__obf_651948667f8af266(__obf_e575c7c6fa901909)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_89bd5bf0a2a3aefa *Lru) RemoveOldest() {
	__obf_89bd5bf0a2a3aefa.
		Lock()
	defer __obf_89bd5bf0a2a3aefa.Unlock()
	if __obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f == nil {
		return
	}
	__obf_e575c7c6fa901909 := __obf_89bd5bf0a2a3aefa.__obf_238b932d03e3f451.Back()
	if __obf_e575c7c6fa901909 != nil {
		__obf_89bd5bf0a2a3aefa.__obf_651948667f8af266(__obf_e575c7c6fa901909)
	}
}

func (__obf_89bd5bf0a2a3aefa *Lru) __obf_651948667f8af266(__obf_d1618b4c9656c93c *list.Element) {
	__obf_89bd5bf0a2a3aefa.__obf_238b932d03e3f451.
		Remove(__obf_d1618b4c9656c93c)
	__obf_b5991b3d1eb715fa := __obf_d1618b4c9656c93c.Value.(*__obf_953b8b04400c10db)
	delete(__obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f, __obf_b5991b3d1eb715fa.__obf_73cb8148cbf55699)
	if __obf_89bd5bf0a2a3aefa.OnEvicted != nil {
		__obf_89bd5bf0a2a3aefa.
			OnEvicted(__obf_b5991b3d1eb715fa.__obf_73cb8148cbf55699,

				// Len returns the number of items in the cache.
				__obf_b5991b3d1eb715fa.__obf_82bf2d5ce3534205)
	}
}

func (__obf_89bd5bf0a2a3aefa *Lru) Len() int {
	if __obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f == nil {
		return 0
	}
	return __obf_89bd5bf0a2a3aefa.

		// Clear purges all stored items from the cache.
		__obf_238b932d03e3f451.Len()
}

func (__obf_89bd5bf0a2a3aefa *Lru) Clear() error {
	__obf_89bd5bf0a2a3aefa.
		Lock()
	defer __obf_89bd5bf0a2a3aefa.Unlock()
	if __obf_89bd5bf0a2a3aefa.OnEvicted != nil {
		for _, __obf_d1618b4c9656c93c := range __obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f {
			__obf_b5991b3d1eb715fa := __obf_d1618b4c9656c93c.Value.(*__obf_953b8b04400c10db)
			__obf_89bd5bf0a2a3aefa.
				OnEvicted(__obf_b5991b3d1eb715fa.__obf_73cb8148cbf55699, __obf_b5991b3d1eb715fa.__obf_82bf2d5ce3534205)
		}
	}
	__obf_89bd5bf0a2a3aefa.__obf_238b932d03e3f451 = nil
	__obf_89bd5bf0a2a3aefa.__obf_72d962f6a40bc95f = nil
	return nil
}
