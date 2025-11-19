package __obf_62df4de078c8d208

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_bfabc99365c8a45e *memcache.Client
}

func (__obf_7c1e1258eb05b926 *Memcache) Has(__obf_4aecf3c737bbe5e8 string) bool {
	_, __obf_0560c6b8f27080cc := __obf_7c1e1258eb05b926.__obf_bfabc99365c8a45e.Get(__obf_4aecf3c737bbe5e8)
	return __obf_0560c6b8f27080cc == nil
}

// Delete implements Cache.
func (__obf_7c1e1258eb05b926 *Memcache) Delete(__obf_0c00b35780000392 string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_f65c89f179a19985 []string) Cache {
	return &Memcache{memcache.New(__obf_f65c89f179a19985...)}
}

// Get get value from memcache.
func (__obf_7c1e1258eb05b926 *Memcache) Get(__obf_4aecf3c737bbe5e8 string) (__obf_f93359fb1c645d1a []byte, __obf_0560c6b8f27080cc error) {
	if __obf_76c3f3755f4b29a1, __obf_0560c6b8f27080cc := __obf_7c1e1258eb05b926.__obf_bfabc99365c8a45e.Get(__obf_4aecf3c737bbe5e8); __obf_0560c6b8f27080cc == nil {

		return __obf_76c3f3755f4b29a1.Value, nil
	} else {
		return nil, __obf_0560c6b8f27080cc
	}
}

func (__obf_7c1e1258eb05b926 *Memcache) Set(__obf_4aecf3c737bbe5e8 string, __obf_676d75836c094b83 any, __obf_24d92102aed9ce02 time.Duration) error {
	__obf_76c3f3755f4b29a1 := memcache.Item{Key: __obf_4aecf3c737bbe5e8, Expiration: int32(__obf_24d92102aed9ce02 / time.Second)}
	var __obf_0560c6b8f27080cc error
	__obf_76c3f3755f4b29a1.Value, __obf_0560c6b8f27080cc = util.AnyToBytes(__obf_676d75836c094b83)
	if __obf_0560c6b8f27080cc != nil {
		return __obf_0560c6b8f27080cc
	}
	return __obf_7c1e1258eb05b926.__obf_bfabc99365c8a45e.Set(&__obf_76c3f3755f4b29a1)
}

func (__obf_7c1e1258eb05b926 *Memcache) Add(__obf_4aecf3c737bbe5e8 string, __obf_676d75836c094b83 any) error {
	__obf_76c3f3755f4b29a1 := memcache.Item{Key: __obf_4aecf3c737bbe5e8}

	var __obf_0560c6b8f27080cc error
	__obf_76c3f3755f4b29a1.Value, __obf_0560c6b8f27080cc = util.AnyToBytes(__obf_676d75836c094b83)
	if __obf_0560c6b8f27080cc != nil {
		return __obf_0560c6b8f27080cc
	}

	return __obf_7c1e1258eb05b926.__obf_bfabc99365c8a45e.Set(&__obf_76c3f3755f4b29a1)
}

// ClearAll clears all cache in memcache.
func (__obf_7c1e1258eb05b926 *Memcache) Clear() error {
	return __obf_7c1e1258eb05b926.__obf_bfabc99365c8a45e.FlushAll()
}

// Incr increases counter.
func (__obf_7c1e1258eb05b926 *Memcache) Incr(__obf_4aecf3c737bbe5e8 string) error {
	_, __obf_0560c6b8f27080cc := __obf_7c1e1258eb05b926.__obf_bfabc99365c8a45e.Increment(__obf_4aecf3c737bbe5e8, 1)
	return __obf_0560c6b8f27080cc
}

// Decr decreases counter.
func (__obf_7c1e1258eb05b926 *Memcache) Decr(__obf_4aecf3c737bbe5e8 string) error {
	_, __obf_0560c6b8f27080cc := __obf_7c1e1258eb05b926.__obf_bfabc99365c8a45e.Decrement(__obf_4aecf3c737bbe5e8, 1)
	return __obf_0560c6b8f27080cc
}

func (__obf_7c1e1258eb05b926 *Memcache) Renew(__obf_4aecf3c737bbe5e8 string, __obf_24d92102aed9ce02 time.Duration) error {
	if __obf_76c3f3755f4b29a1, __obf_0560c6b8f27080cc := __obf_7c1e1258eb05b926.__obf_bfabc99365c8a45e.Get(__obf_4aecf3c737bbe5e8); __obf_0560c6b8f27080cc == nil {
		return __obf_7c1e1258eb05b926.Set(__obf_4aecf3c737bbe5e8, __obf_76c3f3755f4b29a1, __obf_24d92102aed9ce02)
	} else {
		return __obf_0560c6b8f27080cc
	}
}

func (__obf_7c1e1258eb05b926 *Memcache) IsExist(__obf_4aecf3c737bbe5e8 string) bool {
	_, __obf_0560c6b8f27080cc := __obf_7c1e1258eb05b926.__obf_bfabc99365c8a45e.Get(__obf_4aecf3c737bbe5e8)
	return __obf_0560c6b8f27080cc == nil
}

func (__obf_7c1e1258eb05b926 *Memcache) Remove(__obf_4aecf3c737bbe5e8 string) error {
	return __obf_7c1e1258eb05b926.__obf_bfabc99365c8a45e.Delete(__obf_4aecf3c737bbe5e8)
}
