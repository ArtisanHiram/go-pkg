package __obf_2f5c14e012cec42e

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_b698cbb59ce9bf03 *memcache.Client
}

func (__obf_cf1467a859f2b0bd *Memcache) Has(__obf_f0b3ebeba048b5e4 string) bool {
	_, __obf_956c4015cf7da152 := __obf_cf1467a859f2b0bd.__obf_b698cbb59ce9bf03.Get(__obf_f0b3ebeba048b5e4)
	return __obf_956c4015cf7da152 == nil
}

// Delete implements Cache.
func (__obf_cf1467a859f2b0bd *Memcache) Delete(__obf_68047b361b381cb2 string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_e8f21a2a3828cf39 []string) Cache {
	return &Memcache{memcache.New(__obf_e8f21a2a3828cf39...)}
}

// Get get value from memcache.
func (__obf_cf1467a859f2b0bd *Memcache) Get(__obf_f0b3ebeba048b5e4 string) (__obf_e6853b3bab37acbf []byte, __obf_956c4015cf7da152 error) {
	if __obf_c82ba4cf197d7074, __obf_956c4015cf7da152 := __obf_cf1467a859f2b0bd.__obf_b698cbb59ce9bf03.Get(__obf_f0b3ebeba048b5e4); __obf_956c4015cf7da152 == nil {

		return __obf_c82ba4cf197d7074.Value, nil
	} else {
		return nil, __obf_956c4015cf7da152
	}
}

func (__obf_cf1467a859f2b0bd *Memcache) Set(__obf_f0b3ebeba048b5e4 string, __obf_aee38dd09d94cdd5 any, __obf_6459304da6a62a40 time.Duration) error {
	__obf_c82ba4cf197d7074 := memcache.Item{Key: __obf_f0b3ebeba048b5e4, Expiration: int32(__obf_6459304da6a62a40 / time.Second)}
	var __obf_956c4015cf7da152 error
	__obf_c82ba4cf197d7074.Value, __obf_956c4015cf7da152 = util.AnyToBytes(__obf_aee38dd09d94cdd5)
	if __obf_956c4015cf7da152 != nil {
		return __obf_956c4015cf7da152
	}
	return __obf_cf1467a859f2b0bd.__obf_b698cbb59ce9bf03.Set(&__obf_c82ba4cf197d7074)
}

func (__obf_cf1467a859f2b0bd *Memcache) Add(__obf_f0b3ebeba048b5e4 string, __obf_aee38dd09d94cdd5 any) error {
	__obf_c82ba4cf197d7074 := memcache.Item{Key: __obf_f0b3ebeba048b5e4}

	var __obf_956c4015cf7da152 error
	__obf_c82ba4cf197d7074.Value, __obf_956c4015cf7da152 = util.AnyToBytes(__obf_aee38dd09d94cdd5)
	if __obf_956c4015cf7da152 != nil {
		return __obf_956c4015cf7da152
	}

	return __obf_cf1467a859f2b0bd.__obf_b698cbb59ce9bf03.Set(&__obf_c82ba4cf197d7074)
}

// ClearAll clears all cache in memcache.
func (__obf_cf1467a859f2b0bd *Memcache) Clear() error {
	return __obf_cf1467a859f2b0bd.__obf_b698cbb59ce9bf03.FlushAll()
}

// Incr increases counter.
func (__obf_cf1467a859f2b0bd *Memcache) Incr(__obf_f0b3ebeba048b5e4 string) error {
	_, __obf_956c4015cf7da152 := __obf_cf1467a859f2b0bd.__obf_b698cbb59ce9bf03.Increment(__obf_f0b3ebeba048b5e4, 1)
	return __obf_956c4015cf7da152
}

// Decr decreases counter.
func (__obf_cf1467a859f2b0bd *Memcache) Decr(__obf_f0b3ebeba048b5e4 string) error {
	_, __obf_956c4015cf7da152 := __obf_cf1467a859f2b0bd.__obf_b698cbb59ce9bf03.Decrement(__obf_f0b3ebeba048b5e4, 1)
	return __obf_956c4015cf7da152
}

func (__obf_cf1467a859f2b0bd *Memcache) Renew(__obf_f0b3ebeba048b5e4 string, __obf_6459304da6a62a40 time.Duration) error {
	if __obf_c82ba4cf197d7074, __obf_956c4015cf7da152 := __obf_cf1467a859f2b0bd.__obf_b698cbb59ce9bf03.Get(__obf_f0b3ebeba048b5e4); __obf_956c4015cf7da152 == nil {
		return __obf_cf1467a859f2b0bd.Set(__obf_f0b3ebeba048b5e4, __obf_c82ba4cf197d7074, __obf_6459304da6a62a40)
	} else {
		return __obf_956c4015cf7da152
	}
}

func (__obf_cf1467a859f2b0bd *Memcache) IsExist(__obf_f0b3ebeba048b5e4 string) bool {
	_, __obf_956c4015cf7da152 := __obf_cf1467a859f2b0bd.__obf_b698cbb59ce9bf03.Get(__obf_f0b3ebeba048b5e4)
	return __obf_956c4015cf7da152 == nil
}

func (__obf_cf1467a859f2b0bd *Memcache) Remove(__obf_f0b3ebeba048b5e4 string) error {
	return __obf_cf1467a859f2b0bd.__obf_b698cbb59ce9bf03.Delete(__obf_f0b3ebeba048b5e4)
}
