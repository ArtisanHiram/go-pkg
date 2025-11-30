package __obf_3b8640e918b7e3ff

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_59534d8923f168d9 *memcache.Client
}

func (__obf_0a017f229603a007 *Memcache) Has(__obf_3405c14f70aaa4d0 string) bool {
	_, __obf_8c6b39ef87b4061f := __obf_0a017f229603a007.__obf_59534d8923f168d9.Get(__obf_3405c14f70aaa4d0)
	return __obf_8c6b39ef87b4061f == nil
}

// Delete implements Cache.
func (__obf_0a017f229603a007 *Memcache) Delete(__obf_4298344c079705aa string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_abb2cf3daed6c1dd []string) Cache {
	return &Memcache{memcache.New(__obf_abb2cf3daed6c1dd...)}
}

// Get get value from memcache.
func (__obf_0a017f229603a007 *Memcache) Get(__obf_3405c14f70aaa4d0 string) (__obf_ac534eb14660defa []byte, __obf_8c6b39ef87b4061f error) {
	if __obf_64869904d1dc9f43, __obf_8c6b39ef87b4061f := __obf_0a017f229603a007.__obf_59534d8923f168d9.Get(__obf_3405c14f70aaa4d0); __obf_8c6b39ef87b4061f == nil {

		return __obf_64869904d1dc9f43.Value, nil
	} else {
		return nil, __obf_8c6b39ef87b4061f
	}
}

func (__obf_0a017f229603a007 *Memcache) Set(__obf_3405c14f70aaa4d0 string, __obf_13458218654a7f13 any, __obf_433bae995cf2f1ae time.Duration) error {
	__obf_64869904d1dc9f43 := memcache.Item{Key: __obf_3405c14f70aaa4d0, Expiration: int32(__obf_433bae995cf2f1ae / time.Second)}
	var __obf_8c6b39ef87b4061f error
	__obf_64869904d1dc9f43.
		Value, __obf_8c6b39ef87b4061f = util.AnyToBytes(__obf_13458218654a7f13)
	if __obf_8c6b39ef87b4061f != nil {
		return __obf_8c6b39ef87b4061f
	}
	return __obf_0a017f229603a007.__obf_59534d8923f168d9.Set(&__obf_64869904d1dc9f43)
}

func (__obf_0a017f229603a007 *Memcache) Add(__obf_3405c14f70aaa4d0 string, __obf_13458218654a7f13 any) error {
	__obf_64869904d1dc9f43 := memcache.Item{Key: __obf_3405c14f70aaa4d0}

	var __obf_8c6b39ef87b4061f error
	__obf_64869904d1dc9f43.
		Value, __obf_8c6b39ef87b4061f = util.AnyToBytes(__obf_13458218654a7f13)
	if __obf_8c6b39ef87b4061f != nil {
		return __obf_8c6b39ef87b4061f
	}

	return __obf_0a017f229603a007.__obf_59534d8923f168d9.Set(&__obf_64869904d1dc9f43)
}

// ClearAll clears all cache in memcache.
func (__obf_0a017f229603a007 *Memcache) Clear() error {
	return __obf_0a017f229603a007.__obf_59534d8923f168d9.FlushAll()
}

// Incr increases counter.
func (__obf_0a017f229603a007 *Memcache) Incr(__obf_3405c14f70aaa4d0 string) error {
	_, __obf_8c6b39ef87b4061f := __obf_0a017f229603a007.__obf_59534d8923f168d9.Increment(__obf_3405c14f70aaa4d0, 1)
	return __obf_8c6b39ef87b4061f
}

// Decr decreases counter.
func (__obf_0a017f229603a007 *Memcache) Decr(__obf_3405c14f70aaa4d0 string) error {
	_, __obf_8c6b39ef87b4061f := __obf_0a017f229603a007.__obf_59534d8923f168d9.Decrement(__obf_3405c14f70aaa4d0, 1)
	return __obf_8c6b39ef87b4061f
}

func (__obf_0a017f229603a007 *Memcache) Renew(__obf_3405c14f70aaa4d0 string, __obf_433bae995cf2f1ae time.Duration) error {
	if __obf_64869904d1dc9f43, __obf_8c6b39ef87b4061f := __obf_0a017f229603a007.__obf_59534d8923f168d9.Get(__obf_3405c14f70aaa4d0); __obf_8c6b39ef87b4061f == nil {
		return __obf_0a017f229603a007.Set(__obf_3405c14f70aaa4d0, __obf_64869904d1dc9f43, __obf_433bae995cf2f1ae)
	} else {
		return __obf_8c6b39ef87b4061f
	}
}

func (__obf_0a017f229603a007 *Memcache) IsExist(__obf_3405c14f70aaa4d0 string) bool {
	_, __obf_8c6b39ef87b4061f := __obf_0a017f229603a007.__obf_59534d8923f168d9.Get(__obf_3405c14f70aaa4d0)
	return __obf_8c6b39ef87b4061f == nil
}

func (__obf_0a017f229603a007 *Memcache) Remove(__obf_3405c14f70aaa4d0 string) error {
	return __obf_0a017f229603a007.__obf_59534d8923f168d9.Delete(__obf_3405c14f70aaa4d0)
}
