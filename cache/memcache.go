package __obf_79e7502d8563d901

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_d8b078708650cb3f *memcache.Client
}

func (__obf_8147b797bbed6dde *Memcache) Has(__obf_50994613b7653a88 string) bool {
	_, __obf_8d46af2525fab46a := __obf_8147b797bbed6dde.__obf_d8b078708650cb3f.Get(__obf_50994613b7653a88)
	return __obf_8d46af2525fab46a == nil
}

// Delete implements Cache.
func (__obf_8147b797bbed6dde *Memcache) Delete(__obf_bac226386ec983d7 string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_8041a8def04ca050 []string) Cache {
	return &Memcache{memcache.New(__obf_8041a8def04ca050...)}
}

// Get get value from memcache.
func (__obf_8147b797bbed6dde *Memcache) Get(__obf_50994613b7653a88 string) (__obf_03f4c5222d46b04b []byte, __obf_8d46af2525fab46a error) {
	if __obf_a1788cdf3893cb5b, __obf_8d46af2525fab46a := __obf_8147b797bbed6dde.__obf_d8b078708650cb3f.Get(__obf_50994613b7653a88); __obf_8d46af2525fab46a == nil {

		return __obf_a1788cdf3893cb5b.Value, nil
	} else {
		return nil, __obf_8d46af2525fab46a
	}
}

func (__obf_8147b797bbed6dde *Memcache) Set(__obf_50994613b7653a88 string, __obf_e043381e0771feca any, __obf_4ff6d8752a4fea92 time.Duration) error {
	__obf_a1788cdf3893cb5b := memcache.Item{Key: __obf_50994613b7653a88, Expiration: int32(__obf_4ff6d8752a4fea92 / time.Second)}
	var __obf_8d46af2525fab46a error
	__obf_a1788cdf3893cb5b.Value, __obf_8d46af2525fab46a = util.AnyToBytes(__obf_e043381e0771feca)
	if __obf_8d46af2525fab46a != nil {
		return __obf_8d46af2525fab46a
	}
	return __obf_8147b797bbed6dde.__obf_d8b078708650cb3f.Set(&__obf_a1788cdf3893cb5b)
}

func (__obf_8147b797bbed6dde *Memcache) Add(__obf_50994613b7653a88 string, __obf_e043381e0771feca any) error {
	__obf_a1788cdf3893cb5b := memcache.Item{Key: __obf_50994613b7653a88}

	var __obf_8d46af2525fab46a error
	__obf_a1788cdf3893cb5b.Value, __obf_8d46af2525fab46a = util.AnyToBytes(__obf_e043381e0771feca)
	if __obf_8d46af2525fab46a != nil {
		return __obf_8d46af2525fab46a
	}

	return __obf_8147b797bbed6dde.__obf_d8b078708650cb3f.Set(&__obf_a1788cdf3893cb5b)
}

// ClearAll clears all cache in memcache.
func (__obf_8147b797bbed6dde *Memcache) Clear() error {
	return __obf_8147b797bbed6dde.__obf_d8b078708650cb3f.FlushAll()
}

// Incr increases counter.
func (__obf_8147b797bbed6dde *Memcache) Incr(__obf_50994613b7653a88 string) error {
	_, __obf_8d46af2525fab46a := __obf_8147b797bbed6dde.__obf_d8b078708650cb3f.Increment(__obf_50994613b7653a88, 1)
	return __obf_8d46af2525fab46a
}

// Decr decreases counter.
func (__obf_8147b797bbed6dde *Memcache) Decr(__obf_50994613b7653a88 string) error {
	_, __obf_8d46af2525fab46a := __obf_8147b797bbed6dde.__obf_d8b078708650cb3f.Decrement(__obf_50994613b7653a88, 1)
	return __obf_8d46af2525fab46a
}

func (__obf_8147b797bbed6dde *Memcache) Renew(__obf_50994613b7653a88 string, __obf_4ff6d8752a4fea92 time.Duration) error {
	if __obf_a1788cdf3893cb5b, __obf_8d46af2525fab46a := __obf_8147b797bbed6dde.__obf_d8b078708650cb3f.Get(__obf_50994613b7653a88); __obf_8d46af2525fab46a == nil {
		return __obf_8147b797bbed6dde.Set(__obf_50994613b7653a88, __obf_a1788cdf3893cb5b, __obf_4ff6d8752a4fea92)
	} else {
		return __obf_8d46af2525fab46a
	}
}

func (__obf_8147b797bbed6dde *Memcache) IsExist(__obf_50994613b7653a88 string) bool {
	_, __obf_8d46af2525fab46a := __obf_8147b797bbed6dde.__obf_d8b078708650cb3f.Get(__obf_50994613b7653a88)
	return __obf_8d46af2525fab46a == nil
}

func (__obf_8147b797bbed6dde *Memcache) Remove(__obf_50994613b7653a88 string) error {
	return __obf_8147b797bbed6dde.__obf_d8b078708650cb3f.Delete(__obf_50994613b7653a88)
}
