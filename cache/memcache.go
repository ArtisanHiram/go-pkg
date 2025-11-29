package __obf_6fd4fe34e3f784df

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_6c8d97bce271df8c *memcache.Client
}

func (__obf_cdfad3dc535716c4 *Memcache) Has(__obf_ca986a38d1f8fbbb string) bool {
	_, __obf_7f8f5f5c173fa58e := __obf_cdfad3dc535716c4.__obf_6c8d97bce271df8c.Get(__obf_ca986a38d1f8fbbb)
	return __obf_7f8f5f5c173fa58e == nil
}

// Delete implements Cache.
func (__obf_cdfad3dc535716c4 *Memcache) Delete(__obf_c53a458d6b02aeb9 string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_7c06fb7b778400e6 []string) Cache {
	return &Memcache{memcache.New(__obf_7c06fb7b778400e6...)}
}

// Get get value from memcache.
func (__obf_cdfad3dc535716c4 *Memcache) Get(__obf_ca986a38d1f8fbbb string) (__obf_1c110c32e55884ce []byte, __obf_7f8f5f5c173fa58e error) {
	if __obf_90b4f954b51e95a7, __obf_7f8f5f5c173fa58e := __obf_cdfad3dc535716c4.__obf_6c8d97bce271df8c.Get(__obf_ca986a38d1f8fbbb); __obf_7f8f5f5c173fa58e == nil {

		return __obf_90b4f954b51e95a7.Value, nil
	} else {
		return nil, __obf_7f8f5f5c173fa58e
	}
}

func (__obf_cdfad3dc535716c4 *Memcache) Set(__obf_ca986a38d1f8fbbb string, __obf_e1d0258259217620 any, __obf_bfbeaaea582ee60a time.Duration) error {
	__obf_90b4f954b51e95a7 := memcache.Item{Key: __obf_ca986a38d1f8fbbb, Expiration: int32(__obf_bfbeaaea582ee60a / time.Second)}
	var __obf_7f8f5f5c173fa58e error
	__obf_90b4f954b51e95a7.
		Value, __obf_7f8f5f5c173fa58e = util.AnyToBytes(__obf_e1d0258259217620)
	if __obf_7f8f5f5c173fa58e != nil {
		return __obf_7f8f5f5c173fa58e
	}
	return __obf_cdfad3dc535716c4.__obf_6c8d97bce271df8c.Set(&__obf_90b4f954b51e95a7)
}

func (__obf_cdfad3dc535716c4 *Memcache) Add(__obf_ca986a38d1f8fbbb string, __obf_e1d0258259217620 any) error {
	__obf_90b4f954b51e95a7 := memcache.Item{Key: __obf_ca986a38d1f8fbbb}

	var __obf_7f8f5f5c173fa58e error
	__obf_90b4f954b51e95a7.
		Value, __obf_7f8f5f5c173fa58e = util.AnyToBytes(__obf_e1d0258259217620)
	if __obf_7f8f5f5c173fa58e != nil {
		return __obf_7f8f5f5c173fa58e
	}

	return __obf_cdfad3dc535716c4.__obf_6c8d97bce271df8c.Set(&__obf_90b4f954b51e95a7)
}

// ClearAll clears all cache in memcache.
func (__obf_cdfad3dc535716c4 *Memcache) Clear() error {
	return __obf_cdfad3dc535716c4.__obf_6c8d97bce271df8c.FlushAll()
}

// Incr increases counter.
func (__obf_cdfad3dc535716c4 *Memcache) Incr(__obf_ca986a38d1f8fbbb string) error {
	_, __obf_7f8f5f5c173fa58e := __obf_cdfad3dc535716c4.__obf_6c8d97bce271df8c.Increment(__obf_ca986a38d1f8fbbb, 1)
	return __obf_7f8f5f5c173fa58e
}

// Decr decreases counter.
func (__obf_cdfad3dc535716c4 *Memcache) Decr(__obf_ca986a38d1f8fbbb string) error {
	_, __obf_7f8f5f5c173fa58e := __obf_cdfad3dc535716c4.__obf_6c8d97bce271df8c.Decrement(__obf_ca986a38d1f8fbbb, 1)
	return __obf_7f8f5f5c173fa58e
}

func (__obf_cdfad3dc535716c4 *Memcache) Renew(__obf_ca986a38d1f8fbbb string, __obf_bfbeaaea582ee60a time.Duration) error {
	if __obf_90b4f954b51e95a7, __obf_7f8f5f5c173fa58e := __obf_cdfad3dc535716c4.__obf_6c8d97bce271df8c.Get(__obf_ca986a38d1f8fbbb); __obf_7f8f5f5c173fa58e == nil {
		return __obf_cdfad3dc535716c4.Set(__obf_ca986a38d1f8fbbb, __obf_90b4f954b51e95a7, __obf_bfbeaaea582ee60a)
	} else {
		return __obf_7f8f5f5c173fa58e
	}
}

func (__obf_cdfad3dc535716c4 *Memcache) IsExist(__obf_ca986a38d1f8fbbb string) bool {
	_, __obf_7f8f5f5c173fa58e := __obf_cdfad3dc535716c4.__obf_6c8d97bce271df8c.Get(__obf_ca986a38d1f8fbbb)
	return __obf_7f8f5f5c173fa58e == nil
}

func (__obf_cdfad3dc535716c4 *Memcache) Remove(__obf_ca986a38d1f8fbbb string) error {
	return __obf_cdfad3dc535716c4.__obf_6c8d97bce271df8c.Delete(__obf_ca986a38d1f8fbbb)
}
