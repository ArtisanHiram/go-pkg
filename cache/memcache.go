package __obf_9e1ee87c6b054458

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_add8c597f733ba00 *memcache.Client
}

func (__obf_1185c88950520004 *Memcache) Has(__obf_3c13197612c6b39f string) bool {
	_, __obf_498673050542660c := __obf_1185c88950520004.__obf_add8c597f733ba00.Get(__obf_3c13197612c6b39f)
	return __obf_498673050542660c == nil
}

// Delete implements Cache.
func (__obf_1185c88950520004 *Memcache) Delete(__obf_14a292fb641d1b3d string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_07adfd25a67384d0 []string) Cache {
	return &Memcache{memcache.New(__obf_07adfd25a67384d0...)}
}

// Get get value from memcache.
func (__obf_1185c88950520004 *Memcache) Get(__obf_3c13197612c6b39f string) (__obf_870890d378c0cf36 []byte, __obf_498673050542660c error) {
	if __obf_755d49269b07cedb, __obf_498673050542660c := __obf_1185c88950520004.__obf_add8c597f733ba00.Get(__obf_3c13197612c6b39f); __obf_498673050542660c == nil {

		return __obf_755d49269b07cedb.Value, nil
	} else {
		return nil, __obf_498673050542660c
	}
}

func (__obf_1185c88950520004 *Memcache) Set(__obf_3c13197612c6b39f string, __obf_7d0a67e130b6e9b4 any, __obf_5a47a67359bdcb92 time.Duration) error {
	__obf_755d49269b07cedb := memcache.Item{Key: __obf_3c13197612c6b39f, Expiration: int32(__obf_5a47a67359bdcb92 / time.Second)}
	var __obf_498673050542660c error
	__obf_755d49269b07cedb.
		Value, __obf_498673050542660c = util.AnyToBytes(__obf_7d0a67e130b6e9b4)
	if __obf_498673050542660c != nil {
		return __obf_498673050542660c
	}
	return __obf_1185c88950520004.__obf_add8c597f733ba00.Set(&__obf_755d49269b07cedb)
}

func (__obf_1185c88950520004 *Memcache) Add(__obf_3c13197612c6b39f string, __obf_7d0a67e130b6e9b4 any) error {
	__obf_755d49269b07cedb := memcache.Item{Key: __obf_3c13197612c6b39f}

	var __obf_498673050542660c error
	__obf_755d49269b07cedb.
		Value, __obf_498673050542660c = util.AnyToBytes(__obf_7d0a67e130b6e9b4)
	if __obf_498673050542660c != nil {
		return __obf_498673050542660c
	}

	return __obf_1185c88950520004.__obf_add8c597f733ba00.Set(&__obf_755d49269b07cedb)
}

// ClearAll clears all cache in memcache.
func (__obf_1185c88950520004 *Memcache) Clear() error {
	return __obf_1185c88950520004.__obf_add8c597f733ba00.FlushAll()
}

// Incr increases counter.
func (__obf_1185c88950520004 *Memcache) Incr(__obf_3c13197612c6b39f string) error {
	_, __obf_498673050542660c := __obf_1185c88950520004.__obf_add8c597f733ba00.Increment(__obf_3c13197612c6b39f, 1)
	return __obf_498673050542660c
}

// Decr decreases counter.
func (__obf_1185c88950520004 *Memcache) Decr(__obf_3c13197612c6b39f string) error {
	_, __obf_498673050542660c := __obf_1185c88950520004.__obf_add8c597f733ba00.Decrement(__obf_3c13197612c6b39f, 1)
	return __obf_498673050542660c
}

func (__obf_1185c88950520004 *Memcache) Renew(__obf_3c13197612c6b39f string, __obf_5a47a67359bdcb92 time.Duration) error {
	if __obf_755d49269b07cedb, __obf_498673050542660c := __obf_1185c88950520004.__obf_add8c597f733ba00.Get(__obf_3c13197612c6b39f); __obf_498673050542660c == nil {
		return __obf_1185c88950520004.Set(__obf_3c13197612c6b39f, __obf_755d49269b07cedb, __obf_5a47a67359bdcb92)
	} else {
		return __obf_498673050542660c
	}
}

func (__obf_1185c88950520004 *Memcache) IsExist(__obf_3c13197612c6b39f string) bool {
	_, __obf_498673050542660c := __obf_1185c88950520004.__obf_add8c597f733ba00.Get(__obf_3c13197612c6b39f)
	return __obf_498673050542660c == nil
}

func (__obf_1185c88950520004 *Memcache) Remove(__obf_3c13197612c6b39f string) error {
	return __obf_1185c88950520004.__obf_add8c597f733ba00.Delete(__obf_3c13197612c6b39f)
}
