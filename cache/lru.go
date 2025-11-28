package __obf_2f5c14e012cec42e

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
	OnEvicted func(__obf_f0b3ebeba048b5e4 string, __obf_67afdc23c3b56fab any)

	__obf_05e3d02624924845 *list.List
	__obf_2f5c14e012cec42e map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_8a371220abcf1f63 struct {
	__obf_f0b3ebeba048b5e4 string
	__obf_67afdc23c3b56fab []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_6f41af633739f5e7 int) Cache {
	return &Lru{
		MaxEntries:             __obf_6f41af633739f5e7,
		__obf_05e3d02624924845: list.New(),
		__obf_2f5c14e012cec42e: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_4379f98bc2ed9ecc *Lru) Has(__obf_f0b3ebeba048b5e4 string) (__obf_a39fb2a67596cf35 bool) {
	__obf_4379f98bc2ed9ecc.RLock()
	defer __obf_4379f98bc2ed9ecc.RUnlock()

	_, __obf_a39fb2a67596cf35 = __obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e[__obf_f0b3ebeba048b5e4]
	return __obf_a39fb2a67596cf35
}

// Delete key by prefix
func (__obf_4379f98bc2ed9ecc *Lru) Delete(__obf_68047b361b381cb2 string) {
	__obf_4379f98bc2ed9ecc.Lock()
	defer __obf_4379f98bc2ed9ecc.Unlock()
	for __obf_ffd22490bdfba93e, __obf_48da227d69fb34be := range __obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e {
		if strings.HasPrefix(__obf_ffd22490bdfba93e, __obf_68047b361b381cb2) {
			__obf_4379f98bc2ed9ecc.__obf_1aeb35ae68af2347(__obf_48da227d69fb34be)
		}
	}
}

func (__obf_4379f98bc2ed9ecc *Lru) Set(__obf_f0b3ebeba048b5e4 string, __obf_aee38dd09d94cdd5 any, __obf_6459304da6a62a40 time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_4379f98bc2ed9ecc *Lru) Add(__obf_f0b3ebeba048b5e4 string, __obf_67afdc23c3b56fab any) error {
	__obf_4379f98bc2ed9ecc.Lock()
	defer __obf_4379f98bc2ed9ecc.Unlock()
	if __obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e == nil {
		__obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e = make(map[string]*list.Element)
		__obf_4379f98bc2ed9ecc.__obf_05e3d02624924845 = list.New()
	}
	if __obf_26c3f9caa8b21c3e, __obf_a39fb2a67596cf35 := __obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e[__obf_f0b3ebeba048b5e4]; __obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_05e3d02624924845.MoveToFront(__obf_26c3f9caa8b21c3e)
		var __obf_956c4015cf7da152 error
		__obf_26c3f9caa8b21c3e.Value.(*__obf_8a371220abcf1f63).__obf_67afdc23c3b56fab, __obf_956c4015cf7da152 = util.AnyToBytes(__obf_67afdc23c3b56fab)
		return __obf_956c4015cf7da152
	}
	__obf_aee38dd09d94cdd5, __obf_956c4015cf7da152 := util.AnyToBytes(__obf_67afdc23c3b56fab)
	if __obf_956c4015cf7da152 != nil {
		return __obf_956c4015cf7da152
	}
	__obf_155da3dff3cb1162 := __obf_4379f98bc2ed9ecc.__obf_05e3d02624924845.PushFront(&__obf_8a371220abcf1f63{__obf_f0b3ebeba048b5e4, __obf_aee38dd09d94cdd5})
	__obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e[__obf_f0b3ebeba048b5e4] = __obf_155da3dff3cb1162
	if __obf_4379f98bc2ed9ecc.MaxEntries != 0 && __obf_4379f98bc2ed9ecc.__obf_05e3d02624924845.Len() > __obf_4379f98bc2ed9ecc.MaxEntries {
		__obf_4379f98bc2ed9ecc.RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_4379f98bc2ed9ecc *Lru) Get(__obf_f0b3ebeba048b5e4 string) (__obf_67afdc23c3b56fab []byte, __obf_956c4015cf7da152 error) {
	if __obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_155da3dff3cb1162, __obf_8895b0922be8b304 := __obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e[__obf_f0b3ebeba048b5e4]; __obf_8895b0922be8b304 {
		__obf_4379f98bc2ed9ecc.__obf_05e3d02624924845.MoveToFront(__obf_155da3dff3cb1162)
		return __obf_155da3dff3cb1162.Value.(*__obf_8a371220abcf1f63).__obf_67afdc23c3b56fab, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_f0b3ebeba048b5e4)
}

// Remove removes the provided key from the cache.
func (__obf_4379f98bc2ed9ecc *Lru) Remove(__obf_f0b3ebeba048b5e4 string) error {
	__obf_4379f98bc2ed9ecc.Lock()
	defer __obf_4379f98bc2ed9ecc.Unlock()
	if __obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e == nil {
		return ErrCacheEmpty
	}
	if __obf_155da3dff3cb1162, __obf_8895b0922be8b304 := __obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e[__obf_f0b3ebeba048b5e4]; __obf_8895b0922be8b304 {
		__obf_4379f98bc2ed9ecc.__obf_1aeb35ae68af2347(__obf_155da3dff3cb1162)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_4379f98bc2ed9ecc *Lru) RemoveOldest() {
	__obf_4379f98bc2ed9ecc.Lock()
	defer __obf_4379f98bc2ed9ecc.Unlock()
	if __obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e == nil {
		return
	}
	__obf_155da3dff3cb1162 := __obf_4379f98bc2ed9ecc.__obf_05e3d02624924845.Back()
	if __obf_155da3dff3cb1162 != nil {
		__obf_4379f98bc2ed9ecc.__obf_1aeb35ae68af2347(__obf_155da3dff3cb1162)
	}
}

func (__obf_4379f98bc2ed9ecc *Lru) __obf_1aeb35ae68af2347(__obf_5bdf35895068ad96 *list.Element) {
	__obf_4379f98bc2ed9ecc.__obf_05e3d02624924845.Remove(__obf_5bdf35895068ad96)
	__obf_6fa2fbbf74b8daf3 := __obf_5bdf35895068ad96.Value.(*__obf_8a371220abcf1f63)
	delete(__obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e, __obf_6fa2fbbf74b8daf3.__obf_f0b3ebeba048b5e4)
	if __obf_4379f98bc2ed9ecc.OnEvicted != nil {
		__obf_4379f98bc2ed9ecc.OnEvicted(__obf_6fa2fbbf74b8daf3.__obf_f0b3ebeba048b5e4, __obf_6fa2fbbf74b8daf3.__obf_67afdc23c3b56fab)
	}
}

// Len returns the number of items in the cache.
func (__obf_4379f98bc2ed9ecc *Lru) Len() int {
	if __obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e == nil {
		return 0
	}
	return __obf_4379f98bc2ed9ecc.__obf_05e3d02624924845.Len()
}

// Clear purges all stored items from the cache.
func (__obf_4379f98bc2ed9ecc *Lru) Clear() error {
	__obf_4379f98bc2ed9ecc.Lock()
	defer __obf_4379f98bc2ed9ecc.Unlock()
	if __obf_4379f98bc2ed9ecc.OnEvicted != nil {
		for _, __obf_5bdf35895068ad96 := range __obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e {
			__obf_6fa2fbbf74b8daf3 := __obf_5bdf35895068ad96.Value.(*__obf_8a371220abcf1f63)
			__obf_4379f98bc2ed9ecc.OnEvicted(__obf_6fa2fbbf74b8daf3.__obf_f0b3ebeba048b5e4, __obf_6fa2fbbf74b8daf3.__obf_67afdc23c3b56fab)
		}
	}
	__obf_4379f98bc2ed9ecc.__obf_05e3d02624924845 = nil
	__obf_4379f98bc2ed9ecc.__obf_2f5c14e012cec42e = nil
	return nil
}
