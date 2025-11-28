package __obf_8281010a3a6d2ab0

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
func (__obf_c3889543117ef1cd Item) Expired() bool {
	if __obf_c3889543117ef1cd.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_c3889543117ef1cd.Expiration
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
	__obf_76fcab7a69caed43 time.Duration
	__obf_1db22143d25f2542 map[string]Item
	__obf_070e2f23960d0e82 sync.RWMutex
	__obf_e8c78dfcf674a750 func(string, any)
	__obf_c035083bd115908b *__obf_c035083bd115908b
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
func (__obf_868ba86bf2a4661f *MapCache) Set(__obf_f4db856e39cfcf16 string, __obf_037cfaf17c8e5b1a any, __obf_870ec3dd5d63746e time.Duration) {
	// "Inlining" of set
	var __obf_7444330fd960f532 int64
	if __obf_870ec3dd5d63746e == DefaultExpiration {
		__obf_870ec3dd5d63746e = __obf_868ba86bf2a4661f.__obf_76fcab7a69caed43
	}
	if __obf_870ec3dd5d63746e > 0 {
		__obf_7444330fd960f532 = time.Now().Add(__obf_870ec3dd5d63746e).UnixNano()
	}
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = Item{
		Object:     __obf_037cfaf17c8e5b1a,
		Expiration: __obf_7444330fd960f532,
	}
	// TODO: Calls to mu.Unlock are currently not deferred because defer
	// adds ~200 ns (as of go1.)
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
}

func (__obf_868ba86bf2a4661f *MapCache) __obf_974b2163a33277d1(__obf_f4db856e39cfcf16 string, __obf_037cfaf17c8e5b1a any, __obf_870ec3dd5d63746e time.Duration) {
	var __obf_7444330fd960f532 int64
	if __obf_870ec3dd5d63746e == DefaultExpiration {
		__obf_870ec3dd5d63746e = __obf_868ba86bf2a4661f.__obf_76fcab7a69caed43
	}
	if __obf_870ec3dd5d63746e > 0 {
		__obf_7444330fd960f532 = time.Now().Add(__obf_870ec3dd5d63746e).UnixNano()
	}
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = Item{
		Object:     __obf_037cfaf17c8e5b1a,
		Expiration: __obf_7444330fd960f532,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_868ba86bf2a4661f *MapCache) SetDefault(__obf_f4db856e39cfcf16 string, __obf_037cfaf17c8e5b1a any) {
	__obf_868ba86bf2a4661f.Set(__obf_f4db856e39cfcf16, __obf_037cfaf17c8e5b1a, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_868ba86bf2a4661f *MapCache) Add(__obf_f4db856e39cfcf16 string, __obf_037cfaf17c8e5b1a any, __obf_870ec3dd5d63746e time.Duration) error {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	_, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_4c388fe3e1d2c15a(__obf_f4db856e39cfcf16)
	if __obf_9005cf7c7e73c449 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return fmt.Errorf("Item %s already exists", __obf_f4db856e39cfcf16)
	}
	__obf_868ba86bf2a4661f.__obf_974b2163a33277d1(__obf_f4db856e39cfcf16, __obf_037cfaf17c8e5b1a, __obf_870ec3dd5d63746e)
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return nil
}

// Set a new value for the cache key only if it already exists, and the existing
// item hasn't expired. Returns an error otherwise.
func (__obf_868ba86bf2a4661f *MapCache) Replace(__obf_f4db856e39cfcf16 string, __obf_037cfaf17c8e5b1a any, __obf_870ec3dd5d63746e time.Duration) error {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	_, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_4c388fe3e1d2c15a(__obf_f4db856e39cfcf16)
	if !__obf_9005cf7c7e73c449 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_f4db856e39cfcf16)
	}
	__obf_868ba86bf2a4661f.__obf_974b2163a33277d1(__obf_f4db856e39cfcf16, __obf_037cfaf17c8e5b1a, __obf_870ec3dd5d63746e)
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return nil
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
// whether the key was found.
func (__obf_868ba86bf2a4661f *MapCache) Get(__obf_f4db856e39cfcf16 string) (any, bool) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RLock()
	// "Inlining" of get and Expired
	__obf_c3889543117ef1cd, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RUnlock()
		return nil, false
	}
	if __obf_c3889543117ef1cd.Expiration > 0 {
		if time.Now().UnixNano() > __obf_c3889543117ef1cd.Expiration {
			__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RUnlock()
			return nil, false
		}
	}
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RUnlock()
	return __obf_c3889543117ef1cd.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_868ba86bf2a4661f *MapCache) GetWithExpiration(__obf_f4db856e39cfcf16 string) (any, time.Time, bool) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RLock()
	// "Inlining" of get and Expired
	__obf_c3889543117ef1cd, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_c3889543117ef1cd.Expiration > 0 {
		if time.Now().UnixNano() > __obf_c3889543117ef1cd.Expiration {
			__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RUnlock()
			return nil, time.Time{}, false
		}

		// Return the item and the expiration time
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RUnlock()
		return __obf_c3889543117ef1cd.Object, time.Unix(0, __obf_c3889543117ef1cd.Expiration), true
	}

	// If expiration <= 0 (i.e. no expiration time set) then return the item
	// and a zeroed time.Time
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RUnlock()
	return __obf_c3889543117ef1cd.Object, time.Time{}, true
}

func (__obf_868ba86bf2a4661f *MapCache) __obf_4c388fe3e1d2c15a(__obf_f4db856e39cfcf16 string) (any, bool) {
	__obf_c3889543117ef1cd, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_c3889543117ef1cd.Expiration > 0 {
		if time.Now().UnixNano() > __obf_c3889543117ef1cd.Expiration {
			return nil, false
		}
	}
	return __obf_c3889543117ef1cd.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_868ba86bf2a4661f *MapCache) Increment(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int64) error {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	switch __obf_51fc6ca2aca54eed.Object.(type) {
	case int:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(int) + int(__obf_bafdae3e9f221694)
	case int8:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(int8) + int8(__obf_bafdae3e9f221694)
	case int16:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(int16) + int16(__obf_bafdae3e9f221694)
	case int32:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(int32) + int32(__obf_bafdae3e9f221694)
	case int64:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(int64) + __obf_bafdae3e9f221694
	case uint:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uint) + uint(__obf_bafdae3e9f221694)
	case uintptr:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uintptr) + uintptr(__obf_bafdae3e9f221694)
	case uint8:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uint8) + uint8(__obf_bafdae3e9f221694)
	case uint16:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uint16) + uint16(__obf_bafdae3e9f221694)
	case uint32:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uint32) + uint32(__obf_bafdae3e9f221694)
	case uint64:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uint64) + uint64(__obf_bafdae3e9f221694)
	case float32:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(float32) + float32(__obf_bafdae3e9f221694)
	case float64:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(float64) + float64(__obf_bafdae3e9f221694)
	default:
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_f4db856e39cfcf16)
	}
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return nil
}

// Increment an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to increment it by n. Pass a negative number to decrement the
// value. To retrieve the incremented value, use one of the specialized methods,
// e.g. IncrementFloat64.
func (__obf_868ba86bf2a4661f *MapCache) IncrementFloat(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 float64) error {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	switch __obf_51fc6ca2aca54eed.Object.(type) {
	case float32:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(float32) + float32(__obf_bafdae3e9f221694)
	case float64:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(float64) + __obf_bafdae3e9f221694
	default:
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_f4db856e39cfcf16)
	}
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return nil
}

// Increment an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementInt(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int) (int, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(int)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementInt8(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int8) (int8, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(int8)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementInt16(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int16) (int16, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(int16)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementInt32(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int32) (int32, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(int32)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementInt64(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int64) (int64, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(int64)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementUint(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uint) (uint, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uint)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementUintptr(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uintptr) (uintptr, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uintptr)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type uint8 by n. Returns an error if the item's value
// is not an uint8, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementUint8(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uint8) (uint8, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uint8)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementUint16(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uint16) (uint16, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uint16)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementUint32(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uint32) (uint32, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uint32)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementUint64(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uint64) (uint64, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uint64)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementFloat32(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 float32) (float32, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(float32)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Increment an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) IncrementFloat64(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 float64) (float64, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(float64)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 + __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to decrement it by n. To retrieve the decremented value, use one
// of the specialized methods, e.g. DecrementInt64.
func (__obf_868ba86bf2a4661f *MapCache) Decrement(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int64) error {
	// TODO: Implement Increment and Decrement more cleanly.
	// (Cannot do Increment(k, n*-1) for uints.)
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_51fc6ca2aca54eed.Object.(type) {
	case int:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(int) - int(__obf_bafdae3e9f221694)
	case int8:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(int8) - int8(__obf_bafdae3e9f221694)
	case int16:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(int16) - int16(__obf_bafdae3e9f221694)
	case int32:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(int32) - int32(__obf_bafdae3e9f221694)
	case int64:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(int64) - __obf_bafdae3e9f221694
	case uint:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uint) - uint(__obf_bafdae3e9f221694)
	case uintptr:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uintptr) - uintptr(__obf_bafdae3e9f221694)
	case uint8:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uint8) - uint8(__obf_bafdae3e9f221694)
	case uint16:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uint16) - uint16(__obf_bafdae3e9f221694)
	case uint32:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uint32) - uint32(__obf_bafdae3e9f221694)
	case uint64:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(uint64) - uint64(__obf_bafdae3e9f221694)
	case float32:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(float32) - float32(__obf_bafdae3e9f221694)
	case float64:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(float64) - float64(__obf_bafdae3e9f221694)
	default:
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_f4db856e39cfcf16)
	}
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return nil
}

// Decrement an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to decrement it by n. Pass a negative number to decrement the
// value. To retrieve the decremented value, use one of the specialized methods,
// e.g. DecrementFloat64.
func (__obf_868ba86bf2a4661f *MapCache) DecrementFloat(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 float64) error {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	switch __obf_51fc6ca2aca54eed.Object.(type) {
	case float32:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(float32) - float32(__obf_bafdae3e9f221694)
	case float64:
		__obf_51fc6ca2aca54eed.Object = __obf_51fc6ca2aca54eed.Object.(float64) - __obf_bafdae3e9f221694
	default:
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_f4db856e39cfcf16)
	}
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return nil
}

// Decrement an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementInt(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int) (int, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(int)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementInt8(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int8) (int8, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(int8)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementInt16(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int16) (int16, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(int16)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementInt32(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int32) (int32, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(int32)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementInt64(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 int64) (int64, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(int64)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementUint(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uint) (uint, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uint)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementUintptr(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uintptr) (uintptr, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uintptr)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type uint8 by n. Returns an error if the item's value is
// not an uint8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementUint8(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uint8) (uint8, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uint8)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementUint16(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uint16) (uint16, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uint16)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementUint32(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uint32) (uint32, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uint32)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementUint64(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 uint64) (uint64, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(uint64)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementFloat32(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 float32) (float32, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(float32)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Decrement an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_868ba86bf2a4661f *MapCache) DecrementFloat64(__obf_f4db856e39cfcf16 string, __obf_bafdae3e9f221694 float64) (float64, error) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
	if !__obf_9005cf7c7e73c449 || __obf_51fc6ca2aca54eed.Expired() {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f4db856e39cfcf16)
	}
	__obf_becee3cb16693426, __obf_102da286c8b8a571 := __obf_51fc6ca2aca54eed.Object.(float64)
	if !__obf_102da286c8b8a571 {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_f4db856e39cfcf16)
	}
	__obf_39ac721462ef7fcf := __obf_becee3cb16693426 - __obf_bafdae3e9f221694
	__obf_51fc6ca2aca54eed.Object = __obf_39ac721462ef7fcf
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	return __obf_39ac721462ef7fcf, nil
}

// Delete an item from the cache. Does nothing if the key is not in the cache.
func (__obf_868ba86bf2a4661f *MapCache) Delete(__obf_f4db856e39cfcf16 string) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_51fc6ca2aca54eed, __obf_0a5a61a78536f186 := __obf_868ba86bf2a4661f.delete(__obf_f4db856e39cfcf16)
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	if __obf_0a5a61a78536f186 {
		__obf_868ba86bf2a4661f.__obf_e8c78dfcf674a750(__obf_f4db856e39cfcf16, __obf_51fc6ca2aca54eed)
	}
}

func (__obf_868ba86bf2a4661f *MapCache) delete(__obf_f4db856e39cfcf16 string) (any, bool) {
	if __obf_868ba86bf2a4661f.__obf_e8c78dfcf674a750 != nil {
		if __obf_51fc6ca2aca54eed, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]; __obf_9005cf7c7e73c449 {
			delete(__obf_868ba86bf2a4661f.__obf_1db22143d25f2542, __obf_f4db856e39cfcf16)
			return __obf_51fc6ca2aca54eed.Object, true
		}
	}
	delete(__obf_868ba86bf2a4661f.__obf_1db22143d25f2542, __obf_f4db856e39cfcf16)
	return nil, false
}

type __obf_aa1c33634b9bc111 struct {
	__obf_805b9182f4a01a43 string
	__obf_b2563d7af6ddd73e any
}

// Delete all expired items from the cache.
func (__obf_868ba86bf2a4661f *MapCache) DeleteExpired() {
	var __obf_c058ea83c4d55c35 []__obf_aa1c33634b9bc111
	__obf_2dcdf2330b3562b2 := time.Now().UnixNano()
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	for __obf_f4db856e39cfcf16, __obf_51fc6ca2aca54eed := range __obf_868ba86bf2a4661f.__obf_1db22143d25f2542 {
		// "Inlining" of expired
		if __obf_51fc6ca2aca54eed.Expiration > 0 && __obf_2dcdf2330b3562b2 > __obf_51fc6ca2aca54eed.Expiration {
			__obf_b5f4331d33fcc004, __obf_0a5a61a78536f186 := __obf_868ba86bf2a4661f.delete(__obf_f4db856e39cfcf16)
			if __obf_0a5a61a78536f186 {
				__obf_c058ea83c4d55c35 = append(__obf_c058ea83c4d55c35, __obf_aa1c33634b9bc111{__obf_f4db856e39cfcf16, __obf_b5f4331d33fcc004})
			}
		}
	}
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
	for _, __obf_51fc6ca2aca54eed := range __obf_c058ea83c4d55c35 {
		__obf_868ba86bf2a4661f.__obf_e8c78dfcf674a750(__obf_51fc6ca2aca54eed.__obf_805b9182f4a01a43, __obf_51fc6ca2aca54eed.__obf_b2563d7af6ddd73e)
	}
}

// Sets an (optional) function that is called with the key and value when an
// item is evicted from the cache. (Including when it is deleted manually, but
// not when it is overwritten.) Set to nil to disable.
func (__obf_868ba86bf2a4661f *MapCache) OnEvicted(__obf_25fb48b25a74d179 func(string, any)) {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_868ba86bf2a4661f.__obf_e8c78dfcf674a750 = __obf_25fb48b25a74d179
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
}

// Write the cache's items (using Gob) to an io.Writer.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_868ba86bf2a4661f *MapCache) Save(__obf_2a6a4f11c85f6b29 io.Writer) (__obf_fd03130c6793cb3b error) {
	__obf_8d7b945f20d27032 := gob.NewEncoder(__obf_2a6a4f11c85f6b29)
	defer func() {
		if __obf_037cfaf17c8e5b1a := recover(); __obf_037cfaf17c8e5b1a != nil {
			__obf_fd03130c6793cb3b = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RLock()
	defer __obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RUnlock()
	for _, __obf_51fc6ca2aca54eed := range __obf_868ba86bf2a4661f.__obf_1db22143d25f2542 {
		gob.Register(__obf_51fc6ca2aca54eed.Object)
	}
	__obf_fd03130c6793cb3b = __obf_8d7b945f20d27032.Encode(&__obf_868ba86bf2a4661f.__obf_1db22143d25f2542)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_868ba86bf2a4661f *MapCache) SaveFile(__obf_a26cea21a5f6bdd4 string) error {
	__obf_5eae55f4b68692d3, __obf_fd03130c6793cb3b := os.Create(__obf_a26cea21a5f6bdd4)
	if __obf_fd03130c6793cb3b != nil {
		return __obf_fd03130c6793cb3b
	}
	__obf_fd03130c6793cb3b = __obf_868ba86bf2a4661f.Save(__obf_5eae55f4b68692d3)
	if __obf_fd03130c6793cb3b != nil {
		__obf_5eae55f4b68692d3.Close()
		return __obf_fd03130c6793cb3b
	}
	return __obf_5eae55f4b68692d3.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_868ba86bf2a4661f *MapCache) Load(__obf_bcf0de439a6c6928 io.Reader) error {
	__obf_ca4643bcd972a7ae := gob.NewDecoder(__obf_bcf0de439a6c6928)
	__obf_1db22143d25f2542 := map[string]Item{}
	__obf_fd03130c6793cb3b := __obf_ca4643bcd972a7ae.Decode(&__obf_1db22143d25f2542)
	if __obf_fd03130c6793cb3b == nil {
		__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
		defer __obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
		for __obf_f4db856e39cfcf16, __obf_51fc6ca2aca54eed := range __obf_1db22143d25f2542 {
			__obf_b5f4331d33fcc004, __obf_9005cf7c7e73c449 := __obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16]
			if !__obf_9005cf7c7e73c449 || __obf_b5f4331d33fcc004.Expired() {
				__obf_868ba86bf2a4661f.__obf_1db22143d25f2542[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
			}
		}
	}
	return __obf_fd03130c6793cb3b
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_868ba86bf2a4661f *MapCache) LoadFile(__obf_a26cea21a5f6bdd4 string) error {
	__obf_5eae55f4b68692d3, __obf_fd03130c6793cb3b := os.Open(__obf_a26cea21a5f6bdd4)
	if __obf_fd03130c6793cb3b != nil {
		return __obf_fd03130c6793cb3b
	}
	__obf_fd03130c6793cb3b = __obf_868ba86bf2a4661f.Load(__obf_5eae55f4b68692d3)
	if __obf_fd03130c6793cb3b != nil {
		__obf_5eae55f4b68692d3.Close()
		return __obf_fd03130c6793cb3b
	}
	return __obf_5eae55f4b68692d3.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_868ba86bf2a4661f *MapCache) Items() map[string]Item {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RLock()
	defer __obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RUnlock()
	__obf_384bcb41d6b01071 := make(map[string]Item, len(__obf_868ba86bf2a4661f.__obf_1db22143d25f2542))
	__obf_2dcdf2330b3562b2 := time.Now().UnixNano()
	for __obf_f4db856e39cfcf16, __obf_51fc6ca2aca54eed := range __obf_868ba86bf2a4661f.__obf_1db22143d25f2542 {
		// "Inlining" of Expired
		if __obf_51fc6ca2aca54eed.Expiration > 0 {
			if __obf_2dcdf2330b3562b2 > __obf_51fc6ca2aca54eed.Expiration {
				continue
			}
		}
		__obf_384bcb41d6b01071[__obf_f4db856e39cfcf16] = __obf_51fc6ca2aca54eed
	}
	return __obf_384bcb41d6b01071
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_868ba86bf2a4661f *MapCache) ItemCount() int {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RLock()
	__obf_bafdae3e9f221694 := len(__obf_868ba86bf2a4661f.__obf_1db22143d25f2542)
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.RUnlock()
	return __obf_bafdae3e9f221694
}

// Delete all items from the cache.
func (__obf_868ba86bf2a4661f *MapCache) Flush() {
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Lock()
	__obf_868ba86bf2a4661f.__obf_1db22143d25f2542 = map[string]Item{}
	__obf_868ba86bf2a4661f.__obf_070e2f23960d0e82.Unlock()
}

type __obf_c035083bd115908b struct {
	Interval               time.Duration
	__obf_d4ea4908ced9409c chan bool
}

func (__obf_19ac8e57bc9fca1d *__obf_c035083bd115908b) Run(__obf_868ba86bf2a4661f *MapCache) {
	__obf_66ca3461d0a96d1f := time.NewTicker(__obf_19ac8e57bc9fca1d.Interval)
	for {
		select {
		case <-__obf_66ca3461d0a96d1f.C:
			__obf_868ba86bf2a4661f.DeleteExpired()
		case <-__obf_19ac8e57bc9fca1d.__obf_d4ea4908ced9409c:
			__obf_66ca3461d0a96d1f.Stop()
			return
		}
	}
}

func __obf_04f5c9cc6ebe1d74(__obf_868ba86bf2a4661f *MapCache) {
	__obf_868ba86bf2a4661f.__obf_c035083bd115908b.__obf_d4ea4908ced9409c <- true
}

func __obf_b2aa2da5a4922e09(__obf_868ba86bf2a4661f *MapCache, __obf_56dd9789a91df99d time.Duration) {
	__obf_19ac8e57bc9fca1d := &__obf_c035083bd115908b{
		Interval:               __obf_56dd9789a91df99d,
		__obf_d4ea4908ced9409c: make(chan bool),
	}
	__obf_868ba86bf2a4661f.__obf_c035083bd115908b = __obf_19ac8e57bc9fca1d
	go __obf_19ac8e57bc9fca1d.Run(__obf_868ba86bf2a4661f)
}

func __obf_3659aada5fcb1428(__obf_d65ae60e5736cf48 time.Duration, __obf_384bcb41d6b01071 map[string]Item) *MapCache {
	if __obf_d65ae60e5736cf48 == 0 {
		__obf_d65ae60e5736cf48 = -1
	}
	__obf_868ba86bf2a4661f := &MapCache{
		__obf_76fcab7a69caed43: __obf_d65ae60e5736cf48,
		__obf_1db22143d25f2542: __obf_384bcb41d6b01071,
	}
	return __obf_868ba86bf2a4661f
}

func __obf_b60cad191d4e4b8f(__obf_d65ae60e5736cf48 time.Duration, __obf_56dd9789a91df99d time.Duration, __obf_384bcb41d6b01071 map[string]Item) *MapCache {
	__obf_868ba86bf2a4661f := __obf_3659aada5fcb1428(__obf_d65ae60e5736cf48, __obf_384bcb41d6b01071)
	// This trick ensures that the janitor goroutine (which--granted it
	// was enabled--is running DeleteExpired on c forever) does not keep
	// the returned C object from being garbage collected. When it is
	// garbage collected, the finalizer stops the janitor goroutine, after
	// which c can be collected.
	// C := &MapCache{c}
	if __obf_56dd9789a91df99d > 0 {
		__obf_b2aa2da5a4922e09(__obf_868ba86bf2a4661f, __obf_56dd9789a91df99d)
		runtime.SetFinalizer(__obf_868ba86bf2a4661f, __obf_04f5c9cc6ebe1d74)
	}
	return __obf_868ba86bf2a4661f
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_76fcab7a69caed43, __obf_12b98e13d377267d time.Duration) *MapCache {
	__obf_1db22143d25f2542 := make(map[string]Item)
	return __obf_b60cad191d4e4b8f(__obf_76fcab7a69caed43, __obf_12b98e13d377267d, __obf_1db22143d25f2542)
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
func NewFrom(__obf_76fcab7a69caed43, __obf_12b98e13d377267d time.Duration, __obf_1db22143d25f2542 map[string]Item) *MapCache {
	return __obf_b60cad191d4e4b8f(__obf_76fcab7a69caed43, __obf_12b98e13d377267d, __obf_1db22143d25f2542)
}
