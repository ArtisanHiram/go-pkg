package __obf_6b306490bf7221d3

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_e4a967102e02cabb *memcache.Client
}

func (__obf_28d12717306d5141 *Memcache) Has(__obf_fa3a380c35ada5d9 string) bool {
	_, __obf_9ffdd5bbb9f1dc61 := __obf_28d12717306d5141.__obf_e4a967102e02cabb.Get(__obf_fa3a380c35ada5d9)
	return __obf_9ffdd5bbb9f1dc61 == nil
}

// Delete implements Cache.
func (__obf_28d12717306d5141 *Memcache) Delete(__obf_7e6d8e57607bcb1d string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_1bffdd54f6019aa1 []string) Cache {
	return &Memcache{memcache.New(__obf_1bffdd54f6019aa1...)}
}

// Get get value from memcache.
func (__obf_28d12717306d5141 *Memcache) Get(__obf_fa3a380c35ada5d9 string) (__obf_0881f02c55ffa410 []byte, __obf_9ffdd5bbb9f1dc61 error) {
	if __obf_54b1e5958ef673f0, __obf_9ffdd5bbb9f1dc61 := __obf_28d12717306d5141.__obf_e4a967102e02cabb.Get(__obf_fa3a380c35ada5d9); __obf_9ffdd5bbb9f1dc61 == nil {

		return __obf_54b1e5958ef673f0.Value, nil
	} else {
		return nil, __obf_9ffdd5bbb9f1dc61
	}
}

func (__obf_28d12717306d5141 *Memcache) Set(__obf_fa3a380c35ada5d9 string, __obf_17a4af14df8bae4f any, __obf_a15e4803e6944fc5 time.Duration) error {
	__obf_54b1e5958ef673f0 := memcache.Item{Key: __obf_fa3a380c35ada5d9, Expiration: int32(__obf_a15e4803e6944fc5 / time.Second)}
	var __obf_9ffdd5bbb9f1dc61 error
	__obf_54b1e5958ef673f0.
		Value, __obf_9ffdd5bbb9f1dc61 = util.AnyToBytes(__obf_17a4af14df8bae4f)
	if __obf_9ffdd5bbb9f1dc61 != nil {
		return __obf_9ffdd5bbb9f1dc61
	}
	return __obf_28d12717306d5141.__obf_e4a967102e02cabb.Set(&__obf_54b1e5958ef673f0)
}

func (__obf_28d12717306d5141 *Memcache) Add(__obf_fa3a380c35ada5d9 string, __obf_17a4af14df8bae4f any) error {
	__obf_54b1e5958ef673f0 := memcache.Item{Key: __obf_fa3a380c35ada5d9}

	var __obf_9ffdd5bbb9f1dc61 error
	__obf_54b1e5958ef673f0.
		Value, __obf_9ffdd5bbb9f1dc61 = util.AnyToBytes(__obf_17a4af14df8bae4f)
	if __obf_9ffdd5bbb9f1dc61 != nil {
		return __obf_9ffdd5bbb9f1dc61
	}

	return __obf_28d12717306d5141.__obf_e4a967102e02cabb.Set(&__obf_54b1e5958ef673f0)
}

// ClearAll clears all cache in memcache.
func (__obf_28d12717306d5141 *Memcache) Clear() error {
	return __obf_28d12717306d5141.__obf_e4a967102e02cabb.FlushAll()
}

// Incr increases counter.
func (__obf_28d12717306d5141 *Memcache) Incr(__obf_fa3a380c35ada5d9 string) error {
	_, __obf_9ffdd5bbb9f1dc61 := __obf_28d12717306d5141.__obf_e4a967102e02cabb.Increment(__obf_fa3a380c35ada5d9, 1)
	return __obf_9ffdd5bbb9f1dc61
}

// Decr decreases counter.
func (__obf_28d12717306d5141 *Memcache) Decr(__obf_fa3a380c35ada5d9 string) error {
	_, __obf_9ffdd5bbb9f1dc61 := __obf_28d12717306d5141.__obf_e4a967102e02cabb.Decrement(__obf_fa3a380c35ada5d9, 1)
	return __obf_9ffdd5bbb9f1dc61
}

func (__obf_28d12717306d5141 *Memcache) Renew(__obf_fa3a380c35ada5d9 string, __obf_a15e4803e6944fc5 time.Duration) error {
	if __obf_54b1e5958ef673f0, __obf_9ffdd5bbb9f1dc61 := __obf_28d12717306d5141.__obf_e4a967102e02cabb.Get(__obf_fa3a380c35ada5d9); __obf_9ffdd5bbb9f1dc61 == nil {
		return __obf_28d12717306d5141.Set(__obf_fa3a380c35ada5d9, __obf_54b1e5958ef673f0, __obf_a15e4803e6944fc5)
	} else {
		return __obf_9ffdd5bbb9f1dc61
	}
}

func (__obf_28d12717306d5141 *Memcache) IsExist(__obf_fa3a380c35ada5d9 string) bool {
	_, __obf_9ffdd5bbb9f1dc61 := __obf_28d12717306d5141.__obf_e4a967102e02cabb.Get(__obf_fa3a380c35ada5d9)
	return __obf_9ffdd5bbb9f1dc61 == nil
}

func (__obf_28d12717306d5141 *Memcache) Remove(__obf_fa3a380c35ada5d9 string) error {
	return __obf_28d12717306d5141.__obf_e4a967102e02cabb.Delete(__obf_fa3a380c35ada5d9)
}
