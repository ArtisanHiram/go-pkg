package __obf_32d927a1e06b7e71

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_edb55ded92750bb8 *memcache.Client
}

func (__obf_16b47678040c4328 *Memcache) Has(__obf_13113b328a6972ad string) bool {
	_, __obf_19fb5c4c25ff452a := __obf_16b47678040c4328.__obf_edb55ded92750bb8.Get(__obf_13113b328a6972ad)
	return __obf_19fb5c4c25ff452a == nil
}

// Delete implements Cache.
func (__obf_16b47678040c4328 *Memcache) Delete(__obf_e5c9d9e65a3fa384 string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_e25c2423b16da714 []string) Cache {
	return &Memcache{memcache.New(__obf_e25c2423b16da714...)}
}

// Get get value from memcache.
func (__obf_16b47678040c4328 *Memcache) Get(__obf_13113b328a6972ad string) (__obf_0b932686dfd4cc38 []byte, __obf_19fb5c4c25ff452a error) {
	if __obf_868a39cf4c3076e5, __obf_19fb5c4c25ff452a := __obf_16b47678040c4328.__obf_edb55ded92750bb8.Get(__obf_13113b328a6972ad); __obf_19fb5c4c25ff452a == nil {

		return __obf_868a39cf4c3076e5.Value, nil
	} else {
		return nil, __obf_19fb5c4c25ff452a
	}
}

func (__obf_16b47678040c4328 *Memcache) Set(__obf_13113b328a6972ad string, __obf_a967d4ebf1531f4c any, __obf_481cbd2caaded2ed time.Duration) error {
	__obf_868a39cf4c3076e5 := memcache.Item{Key: __obf_13113b328a6972ad, Expiration: int32(__obf_481cbd2caaded2ed / time.Second)}
	var __obf_19fb5c4c25ff452a error
	__obf_868a39cf4c3076e5.Value, __obf_19fb5c4c25ff452a = util.AnyToBytes(__obf_a967d4ebf1531f4c)
	if __obf_19fb5c4c25ff452a != nil {
		return __obf_19fb5c4c25ff452a
	}
	return __obf_16b47678040c4328.__obf_edb55ded92750bb8.Set(&__obf_868a39cf4c3076e5)
}

func (__obf_16b47678040c4328 *Memcache) Add(__obf_13113b328a6972ad string, __obf_a967d4ebf1531f4c any) error {
	__obf_868a39cf4c3076e5 := memcache.Item{Key: __obf_13113b328a6972ad}

	var __obf_19fb5c4c25ff452a error
	__obf_868a39cf4c3076e5.Value, __obf_19fb5c4c25ff452a = util.AnyToBytes(__obf_a967d4ebf1531f4c)
	if __obf_19fb5c4c25ff452a != nil {
		return __obf_19fb5c4c25ff452a
	}

	return __obf_16b47678040c4328.__obf_edb55ded92750bb8.Set(&__obf_868a39cf4c3076e5)
}

// ClearAll clears all cache in memcache.
func (__obf_16b47678040c4328 *Memcache) Clear() error {
	return __obf_16b47678040c4328.__obf_edb55ded92750bb8.FlushAll()
}

// Incr increases counter.
func (__obf_16b47678040c4328 *Memcache) Incr(__obf_13113b328a6972ad string) error {
	_, __obf_19fb5c4c25ff452a := __obf_16b47678040c4328.__obf_edb55ded92750bb8.Increment(__obf_13113b328a6972ad, 1)
	return __obf_19fb5c4c25ff452a
}

// Decr decreases counter.
func (__obf_16b47678040c4328 *Memcache) Decr(__obf_13113b328a6972ad string) error {
	_, __obf_19fb5c4c25ff452a := __obf_16b47678040c4328.__obf_edb55ded92750bb8.Decrement(__obf_13113b328a6972ad, 1)
	return __obf_19fb5c4c25ff452a
}

func (__obf_16b47678040c4328 *Memcache) Renew(__obf_13113b328a6972ad string, __obf_481cbd2caaded2ed time.Duration) error {
	if __obf_868a39cf4c3076e5, __obf_19fb5c4c25ff452a := __obf_16b47678040c4328.__obf_edb55ded92750bb8.Get(__obf_13113b328a6972ad); __obf_19fb5c4c25ff452a == nil {
		return __obf_16b47678040c4328.Set(__obf_13113b328a6972ad, __obf_868a39cf4c3076e5, __obf_481cbd2caaded2ed)
	} else {
		return __obf_19fb5c4c25ff452a
	}
}

func (__obf_16b47678040c4328 *Memcache) IsExist(__obf_13113b328a6972ad string) bool {
	_, __obf_19fb5c4c25ff452a := __obf_16b47678040c4328.__obf_edb55ded92750bb8.Get(__obf_13113b328a6972ad)
	return __obf_19fb5c4c25ff452a == nil
}

func (__obf_16b47678040c4328 *Memcache) Remove(__obf_13113b328a6972ad string) error {
	return __obf_16b47678040c4328.__obf_edb55ded92750bb8.Delete(__obf_13113b328a6972ad)
}
