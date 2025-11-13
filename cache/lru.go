package __obf_f4d8558b35981340

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
	OnEvicted func(__obf_6208471da03258dd string, __obf_5c03cd51840607fb any)

	__obf_2c91823d76886560 *list.List
	__obf_f4d8558b35981340 map[string]*list.Element
	sync.RWMutex
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators

type __obf_9555bc6c57a529a1 struct {
	__obf_6208471da03258dd string
	__obf_5c03cd51840607fb []byte
}

// New creates a new Lru.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func NewLru(__obf_9a3de54b9ae2b5b7 int) Cache {
	return &Lru{
		MaxEntries:             __obf_9a3de54b9ae2b5b7,
		__obf_2c91823d76886560: list.New(),
		__obf_f4d8558b35981340: make(map[string]*list.Element),
	}
}

// Contains check the key exists
func (__obf_a744f58475749237 *Lru) Has(__obf_6208471da03258dd string) (__obf_ed10eafa7bc79f22 bool) {
	__obf_a744f58475749237.RLock()
	defer __obf_a744f58475749237.RUnlock()

	_, __obf_ed10eafa7bc79f22 = __obf_a744f58475749237.__obf_f4d8558b35981340[__obf_6208471da03258dd]
	return __obf_ed10eafa7bc79f22
}

// Delete key by prefix
func (__obf_a744f58475749237 *Lru) Delete(__obf_1261d18b5941c6bd string) {
	__obf_a744f58475749237.Lock()
	defer __obf_a744f58475749237.Unlock()
	for __obf_f038738651cdb77b, __obf_4153159a00c5e3bb := range __obf_a744f58475749237.__obf_f4d8558b35981340 {
		if strings.HasPrefix(__obf_f038738651cdb77b, __obf_1261d18b5941c6bd) {
			__obf_a744f58475749237.__obf_95c8f8691ff4cea6(__obf_4153159a00c5e3bb)
		}
	}
}

func (__obf_a744f58475749237 *Lru) Set(__obf_6208471da03258dd string, __obf_20a54be727e70f3c any, __obf_1e67295112a2a62d time.Duration) error {
	return nil
}

// Add adds a value to the cache.
func (__obf_a744f58475749237 *Lru) Add(__obf_6208471da03258dd string, __obf_5c03cd51840607fb any) error {
	__obf_a744f58475749237.Lock()
	defer __obf_a744f58475749237.Unlock()
	if __obf_a744f58475749237.__obf_f4d8558b35981340 == nil {
		__obf_a744f58475749237.__obf_f4d8558b35981340 = make(map[string]*list.Element)
		__obf_a744f58475749237.__obf_2c91823d76886560 = list.New()
	}
	if __obf_f30287642ff53a00, __obf_ed10eafa7bc79f22 := __obf_a744f58475749237.__obf_f4d8558b35981340[__obf_6208471da03258dd]; __obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_2c91823d76886560.MoveToFront(__obf_f30287642ff53a00)
		var __obf_2d5e4b42e300bedd error
		__obf_f30287642ff53a00.Value.(*__obf_9555bc6c57a529a1).__obf_5c03cd51840607fb, __obf_2d5e4b42e300bedd = util.AnyToBytes(__obf_5c03cd51840607fb)
		return __obf_2d5e4b42e300bedd
	}
	__obf_20a54be727e70f3c, __obf_2d5e4b42e300bedd := util.AnyToBytes(__obf_5c03cd51840607fb)
	if __obf_2d5e4b42e300bedd != nil {
		return __obf_2d5e4b42e300bedd
	}
	__obf_2941fb2f1e08bf9e := __obf_a744f58475749237.__obf_2c91823d76886560.PushFront(&__obf_9555bc6c57a529a1{__obf_6208471da03258dd, __obf_20a54be727e70f3c})
	__obf_a744f58475749237.__obf_f4d8558b35981340[__obf_6208471da03258dd] = __obf_2941fb2f1e08bf9e
	if __obf_a744f58475749237.MaxEntries != 0 && __obf_a744f58475749237.__obf_2c91823d76886560.Len() > __obf_a744f58475749237.MaxEntries {
		__obf_a744f58475749237.RemoveOldest()
	}
	return nil
}

// Get looks up a key's value from the cache.
func (__obf_a744f58475749237 *Lru) Get(__obf_6208471da03258dd string) (__obf_5c03cd51840607fb []byte, __obf_2d5e4b42e300bedd error) {
	if __obf_a744f58475749237.__obf_f4d8558b35981340 == nil {
		return nil, ErrCacheEmpty
	}
	if __obf_2941fb2f1e08bf9e, __obf_94d3e05f4cdf3e46 := __obf_a744f58475749237.__obf_f4d8558b35981340[__obf_6208471da03258dd]; __obf_94d3e05f4cdf3e46 {
		__obf_a744f58475749237.__obf_2c91823d76886560.MoveToFront(__obf_2941fb2f1e08bf9e)
		return __obf_2941fb2f1e08bf9e.Value.(*__obf_9555bc6c57a529a1).__obf_5c03cd51840607fb, nil
	}
	return nil, fmt.Errorf(ErrKeyNotExists.Error(), __obf_6208471da03258dd)
}

// Remove removes the provided key from the cache.
func (__obf_a744f58475749237 *Lru) Remove(__obf_6208471da03258dd string) error {
	__obf_a744f58475749237.Lock()
	defer __obf_a744f58475749237.Unlock()
	if __obf_a744f58475749237.__obf_f4d8558b35981340 == nil {
		return ErrCacheEmpty
	}
	if __obf_2941fb2f1e08bf9e, __obf_94d3e05f4cdf3e46 := __obf_a744f58475749237.__obf_f4d8558b35981340[__obf_6208471da03258dd]; __obf_94d3e05f4cdf3e46 {
		__obf_a744f58475749237.__obf_95c8f8691ff4cea6(__obf_2941fb2f1e08bf9e)
	}

	return nil
}

// RemoveOldest removes the oldest item from the cache.
func (__obf_a744f58475749237 *Lru) RemoveOldest() {
	__obf_a744f58475749237.Lock()
	defer __obf_a744f58475749237.Unlock()
	if __obf_a744f58475749237.__obf_f4d8558b35981340 == nil {
		return
	}
	__obf_2941fb2f1e08bf9e := __obf_a744f58475749237.__obf_2c91823d76886560.Back()
	if __obf_2941fb2f1e08bf9e != nil {
		__obf_a744f58475749237.__obf_95c8f8691ff4cea6(__obf_2941fb2f1e08bf9e)
	}
}

func (__obf_a744f58475749237 *Lru) __obf_95c8f8691ff4cea6(__obf_06e4a4e6e9269bd7 *list.Element) {
	__obf_a744f58475749237.__obf_2c91823d76886560.Remove(__obf_06e4a4e6e9269bd7)
	__obf_9d846bdf5efeecfd := __obf_06e4a4e6e9269bd7.Value.(*__obf_9555bc6c57a529a1)
	delete(__obf_a744f58475749237.__obf_f4d8558b35981340, __obf_9d846bdf5efeecfd.__obf_6208471da03258dd)
	if __obf_a744f58475749237.OnEvicted != nil {
		__obf_a744f58475749237.OnEvicted(__obf_9d846bdf5efeecfd.__obf_6208471da03258dd, __obf_9d846bdf5efeecfd.__obf_5c03cd51840607fb)
	}
}

// Len returns the number of items in the cache.
func (__obf_a744f58475749237 *Lru) Len() int {
	if __obf_a744f58475749237.__obf_f4d8558b35981340 == nil {
		return 0
	}
	return __obf_a744f58475749237.__obf_2c91823d76886560.Len()
}

// Clear purges all stored items from the cache.
func (__obf_a744f58475749237 *Lru) Clear() error {
	__obf_a744f58475749237.Lock()
	defer __obf_a744f58475749237.Unlock()
	if __obf_a744f58475749237.OnEvicted != nil {
		for _, __obf_06e4a4e6e9269bd7 := range __obf_a744f58475749237.__obf_f4d8558b35981340 {
			__obf_9d846bdf5efeecfd := __obf_06e4a4e6e9269bd7.Value.(*__obf_9555bc6c57a529a1)
			__obf_a744f58475749237.OnEvicted(__obf_9d846bdf5efeecfd.__obf_6208471da03258dd, __obf_9d846bdf5efeecfd.__obf_5c03cd51840607fb)
		}
	}
	__obf_a744f58475749237.__obf_2c91823d76886560 = nil
	__obf_a744f58475749237.__obf_f4d8558b35981340 = nil
	return nil
}
