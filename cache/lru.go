package __obf_8281010a3a6d2ab0

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
	OnEvicted func(__obf_805b9182f4a01a43 string, __obf_b2563d7af6ddd73e any)

	__obf_ed3e58a32dfb4c90 *list.List
	__obf_8281010a3a6d2ab0 map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_9457550a2cf07a7b struct {
	__obf_805b9182f4a01a43 string
	__obf_b2563d7af6ddd73e []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_072799e30b1b9896 int) Cache {
	return &Lru{
		MaxEntries:             __obf_072799e30b1b9896,
		__obf_ed3e58a32dfb4c90: list.New(),
		__obf_8281010a3a6d2ab0: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_868ba86bf2a4661f *Lru) Has(__obf_805b9182f4a01a43 string) (__obf_102da286c8b8a571 bool) {
	__obf_868ba86bf2a4661f.RLock()
	defer __obf_868ba86bf2a4661f.RUnlock()

	_, __obf_102da286c8b8a571 = __obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0[__obf_805b9182f4a01a43]
	return __obf_102da286c8b8a571
}

// Delete key by prefix
func (__obf_868ba86bf2a4661f *Lru) Delete(__obf_c2739d84a72f2e49 string) {
	__obf_868ba86bf2a4661f.Lock()
	defer __obf_868ba86bf2a4661f.Unlock()
	for __obf_f4db856e39cfcf16, __obf_51fc6ca2aca54eed := range __obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0 {
		if strings.HasPrefix(__obf_f4db856e39cfcf16, __obf_c2739d84a72f2e49) {
			__obf_868ba86bf2a4661f.__obf_e5ee73cda1d28ed2(__obf_51fc6ca2aca54eed)
		}
	}
}

func (__obf_868ba86bf2a4661f *Lru) Set(__obf_805b9182f4a01a43 string, __obf_a31843a6764aecf9 any, __obf_e2e393700839b297 time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_868ba86bf2a4661f *Lru) Add(__obf_805b9182f4a01a43 string, __obf_b2563d7af6ddd73e any) error {
	__obf_868ba86bf2a4661f.Lock()
	defer __obf_868ba86bf2a4661f.Unlock()
	if __obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0 == nil {
		__obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0 = make(map[string]*list.Element)
		__obf_868ba86bf2a4661f.__obf_ed3e58a32dfb4c90 = list.New()
	}
	if __obf_7286d64f3273ba10, __obf_102da286c8b8a571 := __obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0[__obf_805b9182f4a01a43]; __obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_ed3e58a32dfb4c90.MoveToFront(__obf_7286d64f3273ba10)
		var __obf_fd03130c6793cb3b error
		__obf_7286d64f3273ba10.Value.(*__obf_9457550a2cf07a7b).__obf_b2563d7af6ddd73e, __obf_fd03130c6793cb3b = util.AnyToBytes(__obf_b2563d7af6ddd73e)
		return __obf_fd03130c6793cb3b
	}
	__obf_a31843a6764aecf9, __obf_fd03130c6793cb3b := util.AnyToBytes(__obf_b2563d7af6ddd73e)
	if __obf_fd03130c6793cb3b != nil {
		return __obf_fd03130c6793cb3b
	}
	__obf_3d9e07396c32509d := __obf_868ba86bf2a4661f.__obf_ed3e58a32dfb4c90.PushFront(&__obf_9457550a2cf07a7b{__obf_805b9182f4a01a43, __obf_a31843a6764aecf9})
	__obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0[__obf_805b9182f4a01a43] = __obf_3d9e07396c32509d
	if __obf_868ba86bf2a4661f.MaxEntries != 0 && __obf_868ba86bf2a4661f.__obf_ed3e58a32dfb4c90.Len() > __obf_868ba86bf2a4661f.MaxEntries {
		__obf_868ba86bf2a4661f.RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_868ba86bf2a4661f *Lru) Get(__obf_805b9182f4a01a43 string) (__obf_b2563d7af6ddd73e []byte, __obf_fd03130c6793cb3b error) {
	if __obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0 == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_3d9e07396c32509d, __obf_9c5f6a8cbf061ca9 := __obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0[__obf_805b9182f4a01a43]; __obf_9c5f6a8cbf061ca9 {
		__obf_868ba86bf2a4661f.__obf_ed3e58a32dfb4c90.MoveToFront(__obf_3d9e07396c32509d)
		return __obf_3d9e07396c32509d.Value.(*__obf_9457550a2cf07a7b).__obf_b2563d7af6ddd73e, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_805b9182f4a01a43)
}

// Remove removes the provided key from the cache.
func (__obf_868ba86bf2a4661f *Lru) Remove(__obf_805b9182f4a01a43 string) error {
	__obf_868ba86bf2a4661f.Lock()
	defer __obf_868ba86bf2a4661f.Unlock()
	if __obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0 == nil {
		return ErrCacheEmpty
	}
	if __obf_3d9e07396c32509d, __obf_9c5f6a8cbf061ca9 := __obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0[__obf_805b9182f4a01a43]; __obf_9c5f6a8cbf061ca9 {
		__obf_868ba86bf2a4661f.__obf_e5ee73cda1d28ed2(__obf_3d9e07396c32509d)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_868ba86bf2a4661f *Lru) RemoveOldest() {
	__obf_868ba86bf2a4661f.Lock()
	defer __obf_868ba86bf2a4661f.Unlock()
	if __obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0 == nil {
		return
	}
	__obf_3d9e07396c32509d := __obf_868ba86bf2a4661f.__obf_ed3e58a32dfb4c90.Back()
	if __obf_3d9e07396c32509d != nil {
		__obf_868ba86bf2a4661f.__obf_e5ee73cda1d28ed2(__obf_3d9e07396c32509d)
	}
}

func (__obf_868ba86bf2a4661f *Lru) __obf_e5ee73cda1d28ed2(__obf_7444330fd960f532 *list.Element) {
	__obf_868ba86bf2a4661f.__obf_ed3e58a32dfb4c90.Remove(__obf_7444330fd960f532)
	__obf_bc4a8815f44960aa := __obf_7444330fd960f532.Value.(*__obf_9457550a2cf07a7b)
	delete(__obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0, __obf_bc4a8815f44960aa.__obf_805b9182f4a01a43)
	if __obf_868ba86bf2a4661f.OnEvicted != nil {
		__obf_868ba86bf2a4661f.OnEvicted(__obf_bc4a8815f44960aa.__obf_805b9182f4a01a43, __obf_bc4a8815f44960aa.__obf_b2563d7af6ddd73e)
	}
}

// Len returns the number of items in the cache.
func (__obf_868ba86bf2a4661f *Lru) Len() int {
	if __obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0 == nil {
		return 0
	}
	return __obf_868ba86bf2a4661f.__obf_ed3e58a32dfb4c90.Len()
}

// Clear purges all stored items from the cache.
func (__obf_868ba86bf2a4661f *Lru) Clear() error {
	__obf_868ba86bf2a4661f.Lock()
	defer __obf_868ba86bf2a4661f.Unlock()
	if __obf_868ba86bf2a4661f.OnEvicted != nil {
		for _, __obf_7444330fd960f532 := range __obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0 {
			__obf_bc4a8815f44960aa := __obf_7444330fd960f532.Value.(*__obf_9457550a2cf07a7b)
			__obf_868ba86bf2a4661f.OnEvicted(__obf_bc4a8815f44960aa.__obf_805b9182f4a01a43, __obf_bc4a8815f44960aa.__obf_b2563d7af6ddd73e)
		}
	}
	__obf_868ba86bf2a4661f.__obf_ed3e58a32dfb4c90 = nil
	__obf_868ba86bf2a4661f.__obf_8281010a3a6d2ab0 = nil
	return nil
}
