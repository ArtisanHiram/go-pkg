package __obf_a05682be1c6caf18

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_643796552747948e *memcache.Client
}

func (__obf_fd3028a2c5c10262 *Memcache) Has(__obf_a0e56915a05e8d99 string) bool {
	_, __obf_94de95cb433b50f1 := __obf_fd3028a2c5c10262.__obf_643796552747948e.Get(__obf_a0e56915a05e8d99)
	return __obf_94de95cb433b50f1 == nil
}

// Delete implements Cache.
func (__obf_fd3028a2c5c10262 *Memcache) Delete(__obf_715bf114d0b286fb string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_da89f1ff5c3f9f41 []string) Cache {
	return &Memcache{memcache.New(__obf_da89f1ff5c3f9f41...)}
}

// Get get value from memcache.
func (__obf_fd3028a2c5c10262 *Memcache) Get(__obf_a0e56915a05e8d99 string) (__obf_08736824dd78a6ad []byte, __obf_94de95cb433b50f1 error) {
	if __obf_0bb9085921bdea16, __obf_94de95cb433b50f1 := __obf_fd3028a2c5c10262.__obf_643796552747948e.Get(__obf_a0e56915a05e8d99); __obf_94de95cb433b50f1 == nil {

		return __obf_0bb9085921bdea16.Value, nil
	} else {
		return nil, __obf_94de95cb433b50f1
	}
}

func (__obf_fd3028a2c5c10262 *Memcache) Set(__obf_a0e56915a05e8d99 string, __obf_6030b68d8a95172f any, __obf_c8947601a564de1b time.Duration) error {
	__obf_0bb9085921bdea16 := memcache.Item{Key: __obf_a0e56915a05e8d99, Expiration: int32(__obf_c8947601a564de1b / time.Second)}
	var __obf_94de95cb433b50f1 error
	__obf_0bb9085921bdea16.
		Value, __obf_94de95cb433b50f1 = util.AnyToBytes(__obf_6030b68d8a95172f)
	if __obf_94de95cb433b50f1 != nil {
		return __obf_94de95cb433b50f1
	}
	return __obf_fd3028a2c5c10262.__obf_643796552747948e.Set(&__obf_0bb9085921bdea16)
}

func (__obf_fd3028a2c5c10262 *Memcache) Add(__obf_a0e56915a05e8d99 string, __obf_6030b68d8a95172f any) error {
	__obf_0bb9085921bdea16 := memcache.Item{Key: __obf_a0e56915a05e8d99}

	var __obf_94de95cb433b50f1 error
	__obf_0bb9085921bdea16.
		Value, __obf_94de95cb433b50f1 = util.AnyToBytes(__obf_6030b68d8a95172f)
	if __obf_94de95cb433b50f1 != nil {
		return __obf_94de95cb433b50f1
	}

	return __obf_fd3028a2c5c10262.__obf_643796552747948e.Set(&__obf_0bb9085921bdea16)
}

// ClearAll clears all cache in memcache.
func (__obf_fd3028a2c5c10262 *Memcache) Clear() error {
	return __obf_fd3028a2c5c10262.__obf_643796552747948e.FlushAll()
}

// Incr increases counter.
func (__obf_fd3028a2c5c10262 *Memcache) Incr(__obf_a0e56915a05e8d99 string) error {
	_, __obf_94de95cb433b50f1 := __obf_fd3028a2c5c10262.__obf_643796552747948e.Increment(__obf_a0e56915a05e8d99, 1)
	return __obf_94de95cb433b50f1
}

// Decr decreases counter.
func (__obf_fd3028a2c5c10262 *Memcache) Decr(__obf_a0e56915a05e8d99 string) error {
	_, __obf_94de95cb433b50f1 := __obf_fd3028a2c5c10262.__obf_643796552747948e.Decrement(__obf_a0e56915a05e8d99, 1)
	return __obf_94de95cb433b50f1
}

func (__obf_fd3028a2c5c10262 *Memcache) Renew(__obf_a0e56915a05e8d99 string, __obf_c8947601a564de1b time.Duration) error {
	if __obf_0bb9085921bdea16, __obf_94de95cb433b50f1 := __obf_fd3028a2c5c10262.__obf_643796552747948e.Get(__obf_a0e56915a05e8d99); __obf_94de95cb433b50f1 == nil {
		return __obf_fd3028a2c5c10262.Set(__obf_a0e56915a05e8d99, __obf_0bb9085921bdea16, __obf_c8947601a564de1b)
	} else {
		return __obf_94de95cb433b50f1
	}
}

func (__obf_fd3028a2c5c10262 *Memcache) IsExist(__obf_a0e56915a05e8d99 string) bool {
	_, __obf_94de95cb433b50f1 := __obf_fd3028a2c5c10262.__obf_643796552747948e.Get(__obf_a0e56915a05e8d99)
	return __obf_94de95cb433b50f1 == nil
}

func (__obf_fd3028a2c5c10262 *Memcache) Remove(__obf_a0e56915a05e8d99 string) error {
	return __obf_fd3028a2c5c10262.__obf_643796552747948e.Delete(__obf_a0e56915a05e8d99)
}
