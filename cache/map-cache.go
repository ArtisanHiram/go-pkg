package __obf_32d927a1e06b7e71

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

type Item struct {
	Object     any
	Expiration int64
}

// Returns true if the item has expired.
func (__obf_868a39cf4c3076e5 Item) Expired() bool {
	if __obf_868a39cf4c3076e5.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_868a39cf4c3076e5.Expiration
}

const (
	// For use with functions that take an expiration time.
	NoExpiration time.Duration = -1
	// For use with functions that take an expiration time. Equivalent to
	// passing in the same expiration duration as was given to New() or
	// NewFrom() when the cache was created (e.g. 5 minutes.)
	DefaultExpiration time.Duration = 0
)

type MapCache struct {
	// *cache
	// If this is confusing, see the comment at the bottom of New()
	__obf_db88d00ab46a15a4 time.Duration
	__obf_14629efa3245d903 map[string]Item
	__obf_3123ffce698833d2 sync.RWMutex
	__obf_782d2b7723a6368d func(string, any)
	__obf_a074ec3814b550e3 *__obf_a074ec3814b550e3
}

// type cache struct {
// 	defaultExpiration time.Duration
// 	items             map[string]Item
// 	mu                sync.RWMutex
// 	onEvicted         func(string, any)
// 	janitor           *janitor
// }

// Add an item to the cache, replacing any existing item. If the duration is 0
// (DefaultExpiration), the cache's default expiration time is used. If it is -1
// (NoExpiration), the item never expires.
func (__obf_443bb096afb82210 *MapCache) Set(__obf_4be74b34d9a5767c string, __obf_54f5f063b9803b5c any, __obf_27675eb5bbef6055 time.Duration) {
	// "Inlining" of set
	var __obf_b9653b0201c59d0c int64
	if __obf_27675eb5bbef6055 == DefaultExpiration {
		__obf_27675eb5bbef6055 = __obf_443bb096afb82210.__obf_db88d00ab46a15a4
	}
	if __obf_27675eb5bbef6055 > 0 {
		__obf_b9653b0201c59d0c = time.Now().Add(__obf_27675eb5bbef6055).UnixNano()
	}
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = Item{
		Object:     __obf_54f5f063b9803b5c,
		Expiration: __obf_b9653b0201c59d0c,
	}
	// TODO: Calls to mu.Unlock are currently not deferred because defer
	// adds ~200 ns (as of go1.)
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
}

func (__obf_443bb096afb82210 *MapCache) __obf_5c648bef92716950(__obf_4be74b34d9a5767c string, __obf_54f5f063b9803b5c any, __obf_27675eb5bbef6055 time.Duration) {
	var __obf_b9653b0201c59d0c int64
	if __obf_27675eb5bbef6055 == DefaultExpiration {
		__obf_27675eb5bbef6055 = __obf_443bb096afb82210.__obf_db88d00ab46a15a4
	}
	if __obf_27675eb5bbef6055 > 0 {
		__obf_b9653b0201c59d0c = time.Now().Add(__obf_27675eb5bbef6055).UnixNano()
	}
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = Item{
		Object:     __obf_54f5f063b9803b5c,
		Expiration: __obf_b9653b0201c59d0c,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_443bb096afb82210 *MapCache) SetDefault(__obf_4be74b34d9a5767c string, __obf_54f5f063b9803b5c any) {
	__obf_443bb096afb82210.Set(__obf_4be74b34d9a5767c, __obf_54f5f063b9803b5c, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_443bb096afb82210 *MapCache) Add(__obf_4be74b34d9a5767c string, __obf_54f5f063b9803b5c any, __obf_27675eb5bbef6055 time.Duration) error {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	_, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_594099c9beb08605(__obf_4be74b34d9a5767c)
	if __obf_cd9e67238052e880 {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return fmt.Errorf("Item %s already exists", __obf_4be74b34d9a5767c)
	}
	__obf_443bb096afb82210.__obf_5c648bef92716950(__obf_4be74b34d9a5767c, __obf_54f5f063b9803b5c, __obf_27675eb5bbef6055)
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return nil
}

// Set a new value for the cache key only if it already exists, and the existing
// item hasn't expired. Returns an error otherwise.
func (__obf_443bb096afb82210 *MapCache) Replace(__obf_4be74b34d9a5767c string, __obf_54f5f063b9803b5c any, __obf_27675eb5bbef6055 time.Duration) error {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	_, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_594099c9beb08605(__obf_4be74b34d9a5767c)
	if !__obf_cd9e67238052e880 {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_4be74b34d9a5767c)
	}
	__obf_443bb096afb82210.__obf_5c648bef92716950(__obf_4be74b34d9a5767c, __obf_54f5f063b9803b5c, __obf_27675eb5bbef6055)
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return nil
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
// whether the key was found.
func (__obf_443bb096afb82210 *MapCache) Get(__obf_4be74b34d9a5767c string) (any, bool) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.RLock()
	// "Inlining" of get and Expired
	__obf_868a39cf4c3076e5, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.RUnlock()
		return nil, false
	}
	if __obf_868a39cf4c3076e5.Expiration > 0 {
		if time.Now().UnixNano() > __obf_868a39cf4c3076e5.Expiration {
			__obf_443bb096afb82210.__obf_3123ffce698833d2.RUnlock()
			return nil, false
		}
	}
	__obf_443bb096afb82210.__obf_3123ffce698833d2.RUnlock()
	return __obf_868a39cf4c3076e5.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_443bb096afb82210 *MapCache) GetWithExpiration(__obf_4be74b34d9a5767c string) (any, time.Time, bool) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.RLock()
	// "Inlining" of get and Expired
	__obf_868a39cf4c3076e5, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_868a39cf4c3076e5.Expiration > 0 {
		if time.Now().UnixNano() > __obf_868a39cf4c3076e5.Expiration {
			__obf_443bb096afb82210.__obf_3123ffce698833d2.RUnlock()
			return nil, time.Time{}, false
		}

		// Return the item and the expiration time
		__obf_443bb096afb82210.__obf_3123ffce698833d2.RUnlock()
		return __obf_868a39cf4c3076e5.Object, time.Unix(0, __obf_868a39cf4c3076e5.Expiration), true
	}

	// If expiration <= 0 (i.e. no expiration time set) then return the item
	// and a zeroed time.Time
	__obf_443bb096afb82210.__obf_3123ffce698833d2.RUnlock()
	return __obf_868a39cf4c3076e5.Object, time.Time{}, true
}

func (__obf_443bb096afb82210 *MapCache) __obf_594099c9beb08605(__obf_4be74b34d9a5767c string) (any, bool) {
	__obf_868a39cf4c3076e5, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_868a39cf4c3076e5.Expiration > 0 {
		if time.Now().UnixNano() > __obf_868a39cf4c3076e5.Expiration {
			return nil, false
		}
	}
	return __obf_868a39cf4c3076e5.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_443bb096afb82210 *MapCache) Increment(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int64) error {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	switch __obf_cf59be873d0aaa34.Object.(type) {
	case int:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(int) + int(__obf_4b6270952db86018)
	case int8:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(int8) + int8(__obf_4b6270952db86018)
	case int16:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(int16) + int16(__obf_4b6270952db86018)
	case int32:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(int32) + int32(__obf_4b6270952db86018)
	case int64:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(int64) + __obf_4b6270952db86018
	case uint:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uint) + uint(__obf_4b6270952db86018)
	case uintptr:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uintptr) + uintptr(__obf_4b6270952db86018)
	case uint8:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uint8) + uint8(__obf_4b6270952db86018)
	case uint16:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uint16) + uint16(__obf_4b6270952db86018)
	case uint32:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uint32) + uint32(__obf_4b6270952db86018)
	case uint64:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uint64) + uint64(__obf_4b6270952db86018)
	case float32:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(float32) + float32(__obf_4b6270952db86018)
	case float64:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(float64) + float64(__obf_4b6270952db86018)
	default:
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_4be74b34d9a5767c)
	}
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return nil
}

// Increment an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to increment it by n. Pass a negative number to decrement the
// value. To retrieve the incremented value, use one of the specialized methods,
// e.g. IncrementFloat64.
func (__obf_443bb096afb82210 *MapCache) IncrementFloat(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 float64) error {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	switch __obf_cf59be873d0aaa34.Object.(type) {
	case float32:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(float32) + float32(__obf_4b6270952db86018)
	case float64:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(float64) + __obf_4b6270952db86018
	default:
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_4be74b34d9a5767c)
	}
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return nil
}

// Increment an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementInt(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int) (int, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(int)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementInt8(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int8) (int8, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(int8)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementInt16(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int16) (int16, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(int16)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementInt32(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int32) (int32, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(int32)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementInt64(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int64) (int64, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(int64)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementUint(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uint) (uint, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uint)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementUintptr(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uintptr) (uintptr, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uintptr)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type uint8 by n. Returns an error if the item's value
// is not an uint8, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementUint8(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uint8) (uint8, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uint8)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementUint16(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uint16) (uint16, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uint16)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementUint32(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uint32) (uint32, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uint32)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementUint64(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uint64) (uint64, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uint64)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementFloat32(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 float32) (float32, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(float32)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Increment an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_443bb096afb82210 *MapCache) IncrementFloat64(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 float64) (float64, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(float64)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 + __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to decrement it by n. To retrieve the decremented value, use one
// of the specialized methods, e.g. DecrementInt64.
func (__obf_443bb096afb82210 *MapCache) Decrement(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int64) error {
	// TODO: Implement Increment and Decrement more cleanly.
	// (Cannot do Increment(k, n*-1) for uints.)
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_cf59be873d0aaa34.Object.(type) {
	case int:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(int) - int(__obf_4b6270952db86018)
	case int8:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(int8) - int8(__obf_4b6270952db86018)
	case int16:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(int16) - int16(__obf_4b6270952db86018)
	case int32:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(int32) - int32(__obf_4b6270952db86018)
	case int64:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(int64) - __obf_4b6270952db86018
	case uint:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uint) - uint(__obf_4b6270952db86018)
	case uintptr:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uintptr) - uintptr(__obf_4b6270952db86018)
	case uint8:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uint8) - uint8(__obf_4b6270952db86018)
	case uint16:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uint16) - uint16(__obf_4b6270952db86018)
	case uint32:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uint32) - uint32(__obf_4b6270952db86018)
	case uint64:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(uint64) - uint64(__obf_4b6270952db86018)
	case float32:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(float32) - float32(__obf_4b6270952db86018)
	case float64:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(float64) - float64(__obf_4b6270952db86018)
	default:
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_4be74b34d9a5767c)
	}
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return nil
}

// Decrement an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to decrement it by n. Pass a negative number to decrement the
// value. To retrieve the decremented value, use one of the specialized methods,
// e.g. DecrementFloat64.
func (__obf_443bb096afb82210 *MapCache) DecrementFloat(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 float64) error {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	switch __obf_cf59be873d0aaa34.Object.(type) {
	case float32:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(float32) - float32(__obf_4b6270952db86018)
	case float64:
		__obf_cf59be873d0aaa34.Object = __obf_cf59be873d0aaa34.Object.(float64) - __obf_4b6270952db86018
	default:
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_4be74b34d9a5767c)
	}
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return nil
}

// Decrement an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementInt(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int) (int, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(int)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementInt8(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int8) (int8, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(int8)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementInt16(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int16) (int16, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(int16)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementInt32(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int32) (int32, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(int32)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementInt64(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 int64) (int64, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(int64)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementUint(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uint) (uint, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uint)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementUintptr(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uintptr) (uintptr, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uintptr)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type uint8 by n. Returns an error if the item's value is
// not an uint8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementUint8(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uint8) (uint8, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uint8)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementUint16(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uint16) (uint16, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uint16)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementUint32(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uint32) (uint32, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uint32)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementUint64(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 uint64) (uint64, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(uint64)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementFloat32(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 float32) (float32, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(float32)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Decrement an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_443bb096afb82210 *MapCache) DecrementFloat64(__obf_4be74b34d9a5767c string, __obf_4b6270952db86018 float64) (float64, error) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
	if !__obf_cd9e67238052e880 || __obf_cf59be873d0aaa34.Expired() {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_4be74b34d9a5767c)
	}
	__obf_57ac83e1e1acf2e4, __obf_542c0401f22c9aec := __obf_cf59be873d0aaa34.Object.(float64)
	if !__obf_542c0401f22c9aec {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_4be74b34d9a5767c)
	}
	__obf_73ccc1ca914e623a := __obf_57ac83e1e1acf2e4 - __obf_4b6270952db86018
	__obf_cf59be873d0aaa34.Object = __obf_73ccc1ca914e623a
	__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	return __obf_73ccc1ca914e623a, nil
}

// Delete an item from the cache. Does nothing if the key is not in the cache.
func (__obf_443bb096afb82210 *MapCache) Delete(__obf_4be74b34d9a5767c string) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_cf59be873d0aaa34, __obf_978eed27e88f5b86 := __obf_443bb096afb82210.delete(__obf_4be74b34d9a5767c)
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	if __obf_978eed27e88f5b86 {
		__obf_443bb096afb82210.__obf_782d2b7723a6368d(__obf_4be74b34d9a5767c, __obf_cf59be873d0aaa34)
	}
}

func (__obf_443bb096afb82210 *MapCache) delete(__obf_4be74b34d9a5767c string) (any, bool) {
	if __obf_443bb096afb82210.__obf_782d2b7723a6368d != nil {
		if __obf_cf59be873d0aaa34, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]; __obf_cd9e67238052e880 {
			delete(__obf_443bb096afb82210.__obf_14629efa3245d903, __obf_4be74b34d9a5767c)
			return __obf_cf59be873d0aaa34.Object, true
		}
	}
	delete(__obf_443bb096afb82210.__obf_14629efa3245d903, __obf_4be74b34d9a5767c)
	return nil, false
}

type __obf_8e5b2813349391a3 struct {
	__obf_13113b328a6972ad string
	__obf_b186e1095a1e7f1f any
}

// Delete all expired items from the cache.
func (__obf_443bb096afb82210 *MapCache) DeleteExpired() {
	var __obf_c4915adbeb007cfc []__obf_8e5b2813349391a3
	__obf_a47f8a22d5b883e0 := time.Now().UnixNano()
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	for __obf_4be74b34d9a5767c, __obf_cf59be873d0aaa34 := range __obf_443bb096afb82210.__obf_14629efa3245d903 {
		// "Inlining" of expired
		if __obf_cf59be873d0aaa34.Expiration > 0 && __obf_a47f8a22d5b883e0 > __obf_cf59be873d0aaa34.Expiration {
			__obf_d2708969ff7deb50, __obf_978eed27e88f5b86 := __obf_443bb096afb82210.delete(__obf_4be74b34d9a5767c)
			if __obf_978eed27e88f5b86 {
				__obf_c4915adbeb007cfc = append(__obf_c4915adbeb007cfc, __obf_8e5b2813349391a3{__obf_4be74b34d9a5767c, __obf_d2708969ff7deb50})
			}
		}
	}
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
	for _, __obf_cf59be873d0aaa34 := range __obf_c4915adbeb007cfc {
		__obf_443bb096afb82210.__obf_782d2b7723a6368d(__obf_cf59be873d0aaa34.__obf_13113b328a6972ad, __obf_cf59be873d0aaa34.__obf_b186e1095a1e7f1f)
	}
}

// Sets an (optional) function that is called with the key and value when an
// item is evicted from the cache. (Including when it is deleted manually, but
// not when it is overwritten.) Set to nil to disable.
func (__obf_443bb096afb82210 *MapCache) OnEvicted(__obf_4e6623a0ec6d4355 func(string, any)) {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_443bb096afb82210.__obf_782d2b7723a6368d = __obf_4e6623a0ec6d4355
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
}

// Write the cache's items (using Gob) to an io.Writer.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_443bb096afb82210 *MapCache) Save(__obf_bfad442833941127 io.Writer) (__obf_19fb5c4c25ff452a error) {
	__obf_dd9075bb72f68036 := gob.NewEncoder(__obf_bfad442833941127)
	defer func() {
		if __obf_54f5f063b9803b5c := recover(); __obf_54f5f063b9803b5c != nil {
			__obf_19fb5c4c25ff452a = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_443bb096afb82210.__obf_3123ffce698833d2.RLock()
	defer __obf_443bb096afb82210.__obf_3123ffce698833d2.RUnlock()
	for _, __obf_cf59be873d0aaa34 := range __obf_443bb096afb82210.__obf_14629efa3245d903 {
		gob.Register(__obf_cf59be873d0aaa34.Object)
	}
	__obf_19fb5c4c25ff452a = __obf_dd9075bb72f68036.Encode(&__obf_443bb096afb82210.__obf_14629efa3245d903)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_443bb096afb82210 *MapCache) SaveFile(__obf_1ecc776790fc2973 string) error {
	__obf_1bc58fe6bfc1fa80, __obf_19fb5c4c25ff452a := os.Create(__obf_1ecc776790fc2973)
	if __obf_19fb5c4c25ff452a != nil {
		return __obf_19fb5c4c25ff452a
	}
	__obf_19fb5c4c25ff452a = __obf_443bb096afb82210.Save(__obf_1bc58fe6bfc1fa80)
	if __obf_19fb5c4c25ff452a != nil {
		__obf_1bc58fe6bfc1fa80.Close()
		return __obf_19fb5c4c25ff452a
	}
	return __obf_1bc58fe6bfc1fa80.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_443bb096afb82210 *MapCache) Load(__obf_bb9a2535b3ac3a55 io.Reader) error {
	__obf_8e9650542bb573fe := gob.NewDecoder(__obf_bb9a2535b3ac3a55)
	__obf_14629efa3245d903 := map[string]Item{}
	__obf_19fb5c4c25ff452a := __obf_8e9650542bb573fe.Decode(&__obf_14629efa3245d903)
	if __obf_19fb5c4c25ff452a == nil {
		__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
		defer __obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
		for __obf_4be74b34d9a5767c, __obf_cf59be873d0aaa34 := range __obf_14629efa3245d903 {
			__obf_d2708969ff7deb50, __obf_cd9e67238052e880 := __obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c]
			if !__obf_cd9e67238052e880 || __obf_d2708969ff7deb50.Expired() {
				__obf_443bb096afb82210.__obf_14629efa3245d903[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
			}
		}
	}
	return __obf_19fb5c4c25ff452a
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_443bb096afb82210 *MapCache) LoadFile(__obf_1ecc776790fc2973 string) error {
	__obf_1bc58fe6bfc1fa80, __obf_19fb5c4c25ff452a := os.Open(__obf_1ecc776790fc2973)
	if __obf_19fb5c4c25ff452a != nil {
		return __obf_19fb5c4c25ff452a
	}
	__obf_19fb5c4c25ff452a = __obf_443bb096afb82210.Load(__obf_1bc58fe6bfc1fa80)
	if __obf_19fb5c4c25ff452a != nil {
		__obf_1bc58fe6bfc1fa80.Close()
		return __obf_19fb5c4c25ff452a
	}
	return __obf_1bc58fe6bfc1fa80.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_443bb096afb82210 *MapCache) Items() map[string]Item {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.RLock()
	defer __obf_443bb096afb82210.__obf_3123ffce698833d2.RUnlock()
	__obf_549c2dd367f7ff13 := make(map[string]Item, len(__obf_443bb096afb82210.__obf_14629efa3245d903))
	__obf_a47f8a22d5b883e0 := time.Now().UnixNano()
	for __obf_4be74b34d9a5767c, __obf_cf59be873d0aaa34 := range __obf_443bb096afb82210.__obf_14629efa3245d903 {
		// "Inlining" of Expired
		if __obf_cf59be873d0aaa34.Expiration > 0 {
			if __obf_a47f8a22d5b883e0 > __obf_cf59be873d0aaa34.Expiration {
				continue
			}
		}
		__obf_549c2dd367f7ff13[__obf_4be74b34d9a5767c] = __obf_cf59be873d0aaa34
	}
	return __obf_549c2dd367f7ff13
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_443bb096afb82210 *MapCache) ItemCount() int {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.RLock()
	__obf_4b6270952db86018 := len(__obf_443bb096afb82210.__obf_14629efa3245d903)
	__obf_443bb096afb82210.__obf_3123ffce698833d2.RUnlock()
	return __obf_4b6270952db86018
}

// Delete all items from the cache.
func (__obf_443bb096afb82210 *MapCache) Flush() {
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Lock()
	__obf_443bb096afb82210.__obf_14629efa3245d903 = map[string]Item{}
	__obf_443bb096afb82210.__obf_3123ffce698833d2.Unlock()
}

type __obf_a074ec3814b550e3 struct {
	Interval               time.Duration
	__obf_05b6b50e7f6bc991 chan bool
}

func (__obf_03304024c3a3b5a6 *__obf_a074ec3814b550e3) Run(__obf_443bb096afb82210 *MapCache) {
	__obf_001bd0323c7ce5dc := time.NewTicker(__obf_03304024c3a3b5a6.Interval)
	for {
		select {
		case <-__obf_001bd0323c7ce5dc.C:
			__obf_443bb096afb82210.DeleteExpired()
		case <-__obf_03304024c3a3b5a6.__obf_05b6b50e7f6bc991:
			__obf_001bd0323c7ce5dc.Stop()
			return
		}
	}
}

func __obf_aa52e06794c8fe18(__obf_443bb096afb82210 *MapCache) {
	__obf_443bb096afb82210.__obf_a074ec3814b550e3.__obf_05b6b50e7f6bc991 <- true
}

func __obf_7bd8b1c31d2a891f(__obf_443bb096afb82210 *MapCache, __obf_1bfb5ef8e94f0502 time.Duration) {
	__obf_03304024c3a3b5a6 := &__obf_a074ec3814b550e3{
		Interval:               __obf_1bfb5ef8e94f0502,
		__obf_05b6b50e7f6bc991: make(chan bool),
	}
	__obf_443bb096afb82210.__obf_a074ec3814b550e3 = __obf_03304024c3a3b5a6
	go __obf_03304024c3a3b5a6.Run(__obf_443bb096afb82210)
}

func __obf_6383fd5cbfff3cd1(__obf_08a58213b12ccfde time.Duration, __obf_549c2dd367f7ff13 map[string]Item) *MapCache {
	if __obf_08a58213b12ccfde == 0 {
		__obf_08a58213b12ccfde = -1
	}
	__obf_443bb096afb82210 := &MapCache{
		__obf_db88d00ab46a15a4: __obf_08a58213b12ccfde,
		__obf_14629efa3245d903: __obf_549c2dd367f7ff13,
	}
	return __obf_443bb096afb82210
}

func __obf_055b9791558bd35d(__obf_08a58213b12ccfde time.Duration, __obf_1bfb5ef8e94f0502 time.Duration, __obf_549c2dd367f7ff13 map[string]Item) *MapCache {
	__obf_443bb096afb82210 := __obf_6383fd5cbfff3cd1(__obf_08a58213b12ccfde, __obf_549c2dd367f7ff13)
	// This trick ensures that the janitor goroutine (which--granted it
	// was enabled--is running DeleteExpired on c forever) does not keep
	// the returned C object from being garbage collected. When it is
	// garbage collected, the finalizer stops the janitor goroutine, after
	// which c can be collected.
	// C := &MapCache{c}
	if __obf_1bfb5ef8e94f0502 > 0 {
		__obf_7bd8b1c31d2a891f(__obf_443bb096afb82210, __obf_1bfb5ef8e94f0502)
		runtime.SetFinalizer(__obf_443bb096afb82210, __obf_aa52e06794c8fe18)
	}
	return __obf_443bb096afb82210
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_db88d00ab46a15a4, __obf_ecb4c884999e3b8c time.Duration) *MapCache {
	__obf_14629efa3245d903 := make(map[string]Item)
	return __obf_055b9791558bd35d(__obf_db88d00ab46a15a4, __obf_ecb4c884999e3b8c, __obf_14629efa3245d903)
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
//
// NewFrom() also accepts an items map which will serve as the underlying map
// for the cache. This is useful for starting from a deserialized cache
// (serialized using e.g. gob.Encode() on c.Items()), or passing in e.g.
// make(map[string]Item, 500) to improve startup performance when the cache
// is expected to reach a certain minimum size.
//
// Only the cache's methods synchronize access to this map, so it is not
// recommended to keep any references to the map around after creating a cache.
// If need be, the map can be accessed at a later point using c.Items() (subject
// to the same caveat.)
//
// Note regarding serialization: When using e.g. gob, make sure to
// gob.Register() the individual types stored in the cache before encoding a
// map retrieved with c.Items(), and to register those same types before
// decoding a blob containing an items map.
func NewFrom(__obf_db88d00ab46a15a4, __obf_ecb4c884999e3b8c time.Duration, __obf_14629efa3245d903 map[string]Item) *MapCache {
	return __obf_055b9791558bd35d(__obf_db88d00ab46a15a4, __obf_ecb4c884999e3b8c, __obf_14629efa3245d903)
}
