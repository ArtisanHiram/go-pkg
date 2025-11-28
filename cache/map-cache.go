package __obf_2f5c14e012cec42e

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
func (__obf_c82ba4cf197d7074 Item) Expired() bool {
	if __obf_c82ba4cf197d7074.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_c82ba4cf197d7074.Expiration
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
	__obf_60d547018349aee8 time.Duration
	__obf_1f0b0509f51e115b map[string]Item
	__obf_1791af550e1052ce sync.RWMutex
	__obf_f9fffb9cc4a9e4d3 func(string, any)
	__obf_927553cc46984b1c *__obf_927553cc46984b1c
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
func (__obf_4379f98bc2ed9ecc *MapCache) Set(__obf_ffd22490bdfba93e string, __obf_0aaefdb26a71e0aa any, __obf_6faad9fdc75cdecf time.Duration) {
	// "Inlining" of set
	var __obf_5bdf35895068ad96 int64
	if __obf_6faad9fdc75cdecf == DefaultExpiration {
		__obf_6faad9fdc75cdecf = __obf_4379f98bc2ed9ecc.__obf_60d547018349aee8
	}
	if __obf_6faad9fdc75cdecf > 0 {
		__obf_5bdf35895068ad96 = time.Now().Add(__obf_6faad9fdc75cdecf).UnixNano()
	}
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = Item{
		Object:     __obf_0aaefdb26a71e0aa,
		Expiration: __obf_5bdf35895068ad96,
	}
	// TODO: Calls to mu.Unlock are currently not deferred because defer
	// adds ~200 ns (as of go1.)
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
}

func (__obf_4379f98bc2ed9ecc *MapCache) __obf_cfa024e39c175d91(__obf_ffd22490bdfba93e string, __obf_0aaefdb26a71e0aa any, __obf_6faad9fdc75cdecf time.Duration) {
	var __obf_5bdf35895068ad96 int64
	if __obf_6faad9fdc75cdecf == DefaultExpiration {
		__obf_6faad9fdc75cdecf = __obf_4379f98bc2ed9ecc.__obf_60d547018349aee8
	}
	if __obf_6faad9fdc75cdecf > 0 {
		__obf_5bdf35895068ad96 = time.Now().Add(__obf_6faad9fdc75cdecf).UnixNano()
	}
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = Item{
		Object:     __obf_0aaefdb26a71e0aa,
		Expiration: __obf_5bdf35895068ad96,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_4379f98bc2ed9ecc *MapCache) SetDefault(__obf_ffd22490bdfba93e string, __obf_0aaefdb26a71e0aa any) {
	__obf_4379f98bc2ed9ecc.Set(__obf_ffd22490bdfba93e, __obf_0aaefdb26a71e0aa, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_4379f98bc2ed9ecc *MapCache) Add(__obf_ffd22490bdfba93e string, __obf_0aaefdb26a71e0aa any, __obf_6faad9fdc75cdecf time.Duration) error {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	_, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_7d2097b1098a198c(__obf_ffd22490bdfba93e)
	if __obf_b21fad572c6460d6 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return fmt.Errorf("Item %s already exists", __obf_ffd22490bdfba93e)
	}
	__obf_4379f98bc2ed9ecc.__obf_cfa024e39c175d91(__obf_ffd22490bdfba93e, __obf_0aaefdb26a71e0aa, __obf_6faad9fdc75cdecf)
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return nil
}

// Set a new value for the cache key only if it already exists, and the existing
// item hasn't expired. Returns an error otherwise.
func (__obf_4379f98bc2ed9ecc *MapCache) Replace(__obf_ffd22490bdfba93e string, __obf_0aaefdb26a71e0aa any, __obf_6faad9fdc75cdecf time.Duration) error {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	_, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_7d2097b1098a198c(__obf_ffd22490bdfba93e)
	if !__obf_b21fad572c6460d6 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_ffd22490bdfba93e)
	}
	__obf_4379f98bc2ed9ecc.__obf_cfa024e39c175d91(__obf_ffd22490bdfba93e, __obf_0aaefdb26a71e0aa, __obf_6faad9fdc75cdecf)
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return nil
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
// whether the key was found.
func (__obf_4379f98bc2ed9ecc *MapCache) Get(__obf_ffd22490bdfba93e string) (any, bool) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RLock()
	// "Inlining" of get and Expired
	__obf_c82ba4cf197d7074, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RUnlock()
		return nil, false
	}
	if __obf_c82ba4cf197d7074.Expiration > 0 {
		if time.Now().UnixNano() > __obf_c82ba4cf197d7074.Expiration {
			__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RUnlock()
			return nil, false
		}
	}
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RUnlock()
	return __obf_c82ba4cf197d7074.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_4379f98bc2ed9ecc *MapCache) GetWithExpiration(__obf_ffd22490bdfba93e string) (any, time.Time, bool) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RLock()
	// "Inlining" of get and Expired
	__obf_c82ba4cf197d7074, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_c82ba4cf197d7074.Expiration > 0 {
		if time.Now().UnixNano() > __obf_c82ba4cf197d7074.Expiration {
			__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RUnlock()
			return nil, time.Time{}, false
		}

		// Return the item and the expiration time
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RUnlock()
		return __obf_c82ba4cf197d7074.Object, time.Unix(0, __obf_c82ba4cf197d7074.Expiration), true
	}

	// If expiration <= 0 (i.e. no expiration time set) then return the item
	// and a zeroed time.Time
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RUnlock()
	return __obf_c82ba4cf197d7074.Object, time.Time{}, true
}

func (__obf_4379f98bc2ed9ecc *MapCache) __obf_7d2097b1098a198c(__obf_ffd22490bdfba93e string) (any, bool) {
	__obf_c82ba4cf197d7074, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_c82ba4cf197d7074.Expiration > 0 {
		if time.Now().UnixNano() > __obf_c82ba4cf197d7074.Expiration {
			return nil, false
		}
	}
	return __obf_c82ba4cf197d7074.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_4379f98bc2ed9ecc *MapCache) Increment(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int64) error {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	switch __obf_48da227d69fb34be.Object.(type) {
	case int:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(int) + int(__obf_9844e3ac0a185e1e)
	case int8:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(int8) + int8(__obf_9844e3ac0a185e1e)
	case int16:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(int16) + int16(__obf_9844e3ac0a185e1e)
	case int32:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(int32) + int32(__obf_9844e3ac0a185e1e)
	case int64:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(int64) + __obf_9844e3ac0a185e1e
	case uint:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uint) + uint(__obf_9844e3ac0a185e1e)
	case uintptr:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uintptr) + uintptr(__obf_9844e3ac0a185e1e)
	case uint8:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uint8) + uint8(__obf_9844e3ac0a185e1e)
	case uint16:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uint16) + uint16(__obf_9844e3ac0a185e1e)
	case uint32:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uint32) + uint32(__obf_9844e3ac0a185e1e)
	case uint64:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uint64) + uint64(__obf_9844e3ac0a185e1e)
	case float32:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(float32) + float32(__obf_9844e3ac0a185e1e)
	case float64:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(float64) + float64(__obf_9844e3ac0a185e1e)
	default:
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_ffd22490bdfba93e)
	}
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return nil
}

// Increment an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to increment it by n. Pass a negative number to decrement the
// value. To retrieve the incremented value, use one of the specialized methods,
// e.g. IncrementFloat64.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementFloat(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e float64) error {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	switch __obf_48da227d69fb34be.Object.(type) {
	case float32:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(float32) + float32(__obf_9844e3ac0a185e1e)
	case float64:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(float64) + __obf_9844e3ac0a185e1e
	default:
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_ffd22490bdfba93e)
	}
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return nil
}

// Increment an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementInt(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int) (int, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(int)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementInt8(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int8) (int8, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(int8)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementInt16(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int16) (int16, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(int16)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementInt32(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int32) (int32, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(int32)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementInt64(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int64) (int64, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(int64)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementUint(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uint) (uint, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uint)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementUintptr(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uintptr) (uintptr, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uintptr)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type uint8 by n. Returns an error if the item's value
// is not an uint8, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementUint8(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uint8) (uint8, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uint8)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementUint16(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uint16) (uint16, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uint16)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementUint32(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uint32) (uint32, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uint32)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementUint64(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uint64) (uint64, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uint64)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementFloat32(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e float32) (float32, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(float32)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Increment an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) IncrementFloat64(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e float64) (float64, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(float64)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f + __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to decrement it by n. To retrieve the decremented value, use one
// of the specialized methods, e.g. DecrementInt64.
func (__obf_4379f98bc2ed9ecc *MapCache) Decrement(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int64) error {
	// TODO: Implement Increment and Decrement more cleanly.
	// (Cannot do Increment(k, n*-1) for uints.)
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_48da227d69fb34be.Object.(type) {
	case int:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(int) - int(__obf_9844e3ac0a185e1e)
	case int8:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(int8) - int8(__obf_9844e3ac0a185e1e)
	case int16:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(int16) - int16(__obf_9844e3ac0a185e1e)
	case int32:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(int32) - int32(__obf_9844e3ac0a185e1e)
	case int64:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(int64) - __obf_9844e3ac0a185e1e
	case uint:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uint) - uint(__obf_9844e3ac0a185e1e)
	case uintptr:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uintptr) - uintptr(__obf_9844e3ac0a185e1e)
	case uint8:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uint8) - uint8(__obf_9844e3ac0a185e1e)
	case uint16:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uint16) - uint16(__obf_9844e3ac0a185e1e)
	case uint32:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uint32) - uint32(__obf_9844e3ac0a185e1e)
	case uint64:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(uint64) - uint64(__obf_9844e3ac0a185e1e)
	case float32:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(float32) - float32(__obf_9844e3ac0a185e1e)
	case float64:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(float64) - float64(__obf_9844e3ac0a185e1e)
	default:
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_ffd22490bdfba93e)
	}
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return nil
}

// Decrement an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to decrement it by n. Pass a negative number to decrement the
// value. To retrieve the decremented value, use one of the specialized methods,
// e.g. DecrementFloat64.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementFloat(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e float64) error {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	switch __obf_48da227d69fb34be.Object.(type) {
	case float32:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(float32) - float32(__obf_9844e3ac0a185e1e)
	case float64:
		__obf_48da227d69fb34be.Object = __obf_48da227d69fb34be.Object.(float64) - __obf_9844e3ac0a185e1e
	default:
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_ffd22490bdfba93e)
	}
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return nil
}

// Decrement an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementInt(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int) (int, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(int)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementInt8(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int8) (int8, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(int8)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementInt16(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int16) (int16, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(int16)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementInt32(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int32) (int32, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(int32)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementInt64(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e int64) (int64, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(int64)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementUint(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uint) (uint, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uint)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementUintptr(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uintptr) (uintptr, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uintptr)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type uint8 by n. Returns an error if the item's value is
// not an uint8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementUint8(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uint8) (uint8, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uint8)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementUint16(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uint16) (uint16, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uint16)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementUint32(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uint32) (uint32, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uint32)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementUint64(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e uint64) (uint64, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(uint64)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementFloat32(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e float32) (float32, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(float32)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Decrement an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_4379f98bc2ed9ecc *MapCache) DecrementFloat64(__obf_ffd22490bdfba93e string, __obf_9844e3ac0a185e1e float64) (float64, error) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
	if !__obf_b21fad572c6460d6 || __obf_48da227d69fb34be.Expired() {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffd22490bdfba93e)
	}
	__obf_a2ce163d6505135f, __obf_a39fb2a67596cf35 := __obf_48da227d69fb34be.Object.(float64)
	if !__obf_a39fb2a67596cf35 {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_ffd22490bdfba93e)
	}
	__obf_6c33f4d4599dd3c5 := __obf_a2ce163d6505135f - __obf_9844e3ac0a185e1e
	__obf_48da227d69fb34be.Object = __obf_6c33f4d4599dd3c5
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	return __obf_6c33f4d4599dd3c5, nil
}

// Delete an item from the cache. Does nothing if the key is not in the cache.
func (__obf_4379f98bc2ed9ecc *MapCache) Delete(__obf_ffd22490bdfba93e string) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_48da227d69fb34be, __obf_026e0eda53f0bab1 := __obf_4379f98bc2ed9ecc.delete(__obf_ffd22490bdfba93e)
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	if __obf_026e0eda53f0bab1 {
		__obf_4379f98bc2ed9ecc.__obf_f9fffb9cc4a9e4d3(__obf_ffd22490bdfba93e, __obf_48da227d69fb34be)
	}
}

func (__obf_4379f98bc2ed9ecc *MapCache) delete(__obf_ffd22490bdfba93e string) (any, bool) {
	if __obf_4379f98bc2ed9ecc.__obf_f9fffb9cc4a9e4d3 != nil {
		if __obf_48da227d69fb34be, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]; __obf_b21fad572c6460d6 {
			delete(__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b, __obf_ffd22490bdfba93e)
			return __obf_48da227d69fb34be.Object, true
		}
	}
	delete(__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b, __obf_ffd22490bdfba93e)
	return nil, false
}

type __obf_2ac198141cdbeed1 struct {
	__obf_f0b3ebeba048b5e4 string
	__obf_67afdc23c3b56fab any
}

// Delete all expired items from the cache.
func (__obf_4379f98bc2ed9ecc *MapCache) DeleteExpired() {
	var __obf_db9afc69fe2fe9c5 []__obf_2ac198141cdbeed1
	__obf_265fb276057b3ae8 := time.Now().UnixNano()
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	for __obf_ffd22490bdfba93e, __obf_48da227d69fb34be := range __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b {
		// "Inlining" of expired
		if __obf_48da227d69fb34be.Expiration > 0 && __obf_265fb276057b3ae8 > __obf_48da227d69fb34be.Expiration {
			__obf_18fc7be52c7a8c07, __obf_026e0eda53f0bab1 := __obf_4379f98bc2ed9ecc.delete(__obf_ffd22490bdfba93e)
			if __obf_026e0eda53f0bab1 {
				__obf_db9afc69fe2fe9c5 = append(__obf_db9afc69fe2fe9c5, __obf_2ac198141cdbeed1{__obf_ffd22490bdfba93e, __obf_18fc7be52c7a8c07})
			}
		}
	}
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
	for _, __obf_48da227d69fb34be := range __obf_db9afc69fe2fe9c5 {
		__obf_4379f98bc2ed9ecc.__obf_f9fffb9cc4a9e4d3(__obf_48da227d69fb34be.__obf_f0b3ebeba048b5e4, __obf_48da227d69fb34be.__obf_67afdc23c3b56fab)
	}
}

// Sets an (optional) function that is called with the key and value when an
// item is evicted from the cache. (Including when it is deleted manually, but
// not when it is overwritten.) Set to nil to disable.
func (__obf_4379f98bc2ed9ecc *MapCache) OnEvicted(__obf_21dabb6b735960ad func(string, any)) {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_4379f98bc2ed9ecc.__obf_f9fffb9cc4a9e4d3 = __obf_21dabb6b735960ad
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
}

// Write the cache's items (using Gob) to an io.Writer.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_4379f98bc2ed9ecc *MapCache) Save(__obf_c8adc7023bb4962d io.Writer) (__obf_956c4015cf7da152 error) {
	__obf_ce3d12e532d54a7f := gob.NewEncoder(__obf_c8adc7023bb4962d)
	defer func() {
		if __obf_0aaefdb26a71e0aa := recover(); __obf_0aaefdb26a71e0aa != nil {
			__obf_956c4015cf7da152 = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RLock()
	defer __obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RUnlock()
	for _, __obf_48da227d69fb34be := range __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b {
		gob.Register(__obf_48da227d69fb34be.Object)
	}
	__obf_956c4015cf7da152 = __obf_ce3d12e532d54a7f.Encode(&__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_4379f98bc2ed9ecc *MapCache) SaveFile(__obf_e9ecb0e7af59c68e string) error {
	__obf_664d11e3092a429f, __obf_956c4015cf7da152 := os.Create(__obf_e9ecb0e7af59c68e)
	if __obf_956c4015cf7da152 != nil {
		return __obf_956c4015cf7da152
	}
	__obf_956c4015cf7da152 = __obf_4379f98bc2ed9ecc.Save(__obf_664d11e3092a429f)
	if __obf_956c4015cf7da152 != nil {
		__obf_664d11e3092a429f.Close()
		return __obf_956c4015cf7da152
	}
	return __obf_664d11e3092a429f.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_4379f98bc2ed9ecc *MapCache) Load(__obf_79efa7595252dd83 io.Reader) error {
	__obf_84ab220dd6afe32a := gob.NewDecoder(__obf_79efa7595252dd83)
	__obf_1f0b0509f51e115b := map[string]Item{}
	__obf_956c4015cf7da152 := __obf_84ab220dd6afe32a.Decode(&__obf_1f0b0509f51e115b)
	if __obf_956c4015cf7da152 == nil {
		__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
		defer __obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
		for __obf_ffd22490bdfba93e, __obf_48da227d69fb34be := range __obf_1f0b0509f51e115b {
			__obf_18fc7be52c7a8c07, __obf_b21fad572c6460d6 := __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e]
			if !__obf_b21fad572c6460d6 || __obf_18fc7be52c7a8c07.Expired() {
				__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
			}
		}
	}
	return __obf_956c4015cf7da152
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_4379f98bc2ed9ecc *MapCache) LoadFile(__obf_e9ecb0e7af59c68e string) error {
	__obf_664d11e3092a429f, __obf_956c4015cf7da152 := os.Open(__obf_e9ecb0e7af59c68e)
	if __obf_956c4015cf7da152 != nil {
		return __obf_956c4015cf7da152
	}
	__obf_956c4015cf7da152 = __obf_4379f98bc2ed9ecc.Load(__obf_664d11e3092a429f)
	if __obf_956c4015cf7da152 != nil {
		__obf_664d11e3092a429f.Close()
		return __obf_956c4015cf7da152
	}
	return __obf_664d11e3092a429f.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_4379f98bc2ed9ecc *MapCache) Items() map[string]Item {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RLock()
	defer __obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RUnlock()
	__obf_b127f8671154d05f := make(map[string]Item, len(__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b))
	__obf_265fb276057b3ae8 := time.Now().UnixNano()
	for __obf_ffd22490bdfba93e, __obf_48da227d69fb34be := range __obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b {
		// "Inlining" of Expired
		if __obf_48da227d69fb34be.Expiration > 0 {
			if __obf_265fb276057b3ae8 > __obf_48da227d69fb34be.Expiration {
				continue
			}
		}
		__obf_b127f8671154d05f[__obf_ffd22490bdfba93e] = __obf_48da227d69fb34be
	}
	return __obf_b127f8671154d05f
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_4379f98bc2ed9ecc *MapCache) ItemCount() int {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RLock()
	__obf_9844e3ac0a185e1e := len(__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b)
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.RUnlock()
	return __obf_9844e3ac0a185e1e
}

// Delete all items from the cache.
func (__obf_4379f98bc2ed9ecc *MapCache) Flush() {
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Lock()
	__obf_4379f98bc2ed9ecc.__obf_1f0b0509f51e115b = map[string]Item{}
	__obf_4379f98bc2ed9ecc.__obf_1791af550e1052ce.Unlock()
}

type __obf_927553cc46984b1c struct {
	Interval               time.Duration
	__obf_09568715bd3160ac chan bool
}

func (__obf_d1e651c15abff19c *__obf_927553cc46984b1c) Run(__obf_4379f98bc2ed9ecc *MapCache) {
	__obf_c661ab57795bb0d1 := time.NewTicker(__obf_d1e651c15abff19c.Interval)
	for {
		select {
		case <-__obf_c661ab57795bb0d1.C:
			__obf_4379f98bc2ed9ecc.DeleteExpired()
		case <-__obf_d1e651c15abff19c.__obf_09568715bd3160ac:
			__obf_c661ab57795bb0d1.Stop()
			return
		}
	}
}

func __obf_8abcbdf2802232e9(__obf_4379f98bc2ed9ecc *MapCache) {
	__obf_4379f98bc2ed9ecc.__obf_927553cc46984b1c.__obf_09568715bd3160ac <- true
}

func __obf_5e2c16346a0f2ec0(__obf_4379f98bc2ed9ecc *MapCache, __obf_379aaa7f9d68a02c time.Duration) {
	__obf_d1e651c15abff19c := &__obf_927553cc46984b1c{
		Interval:               __obf_379aaa7f9d68a02c,
		__obf_09568715bd3160ac: make(chan bool),
	}
	__obf_4379f98bc2ed9ecc.__obf_927553cc46984b1c = __obf_d1e651c15abff19c
	go __obf_d1e651c15abff19c.Run(__obf_4379f98bc2ed9ecc)
}

func __obf_dec5b075a2e044d9(__obf_6b89a931139659ca time.Duration, __obf_b127f8671154d05f map[string]Item) *MapCache {
	if __obf_6b89a931139659ca == 0 {
		__obf_6b89a931139659ca = -1
	}
	__obf_4379f98bc2ed9ecc := &MapCache{
		__obf_60d547018349aee8: __obf_6b89a931139659ca,
		__obf_1f0b0509f51e115b: __obf_b127f8671154d05f,
	}
	return __obf_4379f98bc2ed9ecc
}

func __obf_6a8efcc6f17f2f96(__obf_6b89a931139659ca time.Duration, __obf_379aaa7f9d68a02c time.Duration, __obf_b127f8671154d05f map[string]Item) *MapCache {
	__obf_4379f98bc2ed9ecc := __obf_dec5b075a2e044d9(__obf_6b89a931139659ca, __obf_b127f8671154d05f)
	// This trick ensures that the janitor goroutine (which--granted it
	// was enabled--is running DeleteExpired on c forever) does not keep
	// the returned C object from being garbage collected. When it is
	// garbage collected, the finalizer stops the janitor goroutine, after
	// which c can be collected.
	// C := &MapCache{c}
	if __obf_379aaa7f9d68a02c > 0 {
		__obf_5e2c16346a0f2ec0(__obf_4379f98bc2ed9ecc, __obf_379aaa7f9d68a02c)
		runtime.SetFinalizer(__obf_4379f98bc2ed9ecc, __obf_8abcbdf2802232e9)
	}
	return __obf_4379f98bc2ed9ecc
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_60d547018349aee8, __obf_07a4b4595e9225f5 time.Duration) *MapCache {
	__obf_1f0b0509f51e115b := make(map[string]Item)
	return __obf_6a8efcc6f17f2f96(__obf_60d547018349aee8, __obf_07a4b4595e9225f5, __obf_1f0b0509f51e115b)
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
func NewFrom(__obf_60d547018349aee8, __obf_07a4b4595e9225f5 time.Duration, __obf_1f0b0509f51e115b map[string]Item) *MapCache {
	return __obf_6a8efcc6f17f2f96(__obf_60d547018349aee8, __obf_07a4b4595e9225f5, __obf_1f0b0509f51e115b)
}
