package __obf_3b8640e918b7e3ff

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
func (__obf_64869904d1dc9f43 Item) Expired() bool {
	if __obf_64869904d1dc9f43.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_64869904d1dc9f43.Expiration
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
	__obf_113e9a23713e45c1 time. // *cache
		// If this is confusing, see the comment at the bottom of New()
		Duration
	__obf_717158131d1e5e98 map[string]Item
	__obf_d8684d60a98e1a8c sync.RWMutex
	__obf_2f7516313eb1f286 func(string, any)
	__obf_b14857644ea2159c *__obf_b14857644ea2159c
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
func (__obf_5b354994d2d4bcf1 *MapCache) Set(__obf_ebdd4526b9b834d1 string, __obf_5c75aa5f1a5ca219 any, __obf_cad7cc30a9b8797e time.Duration) {
	// "Inlining" of set
	var __obf_43b169a18e2edb05 int64
	if __obf_cad7cc30a9b8797e == DefaultExpiration {
		__obf_cad7cc30a9b8797e = __obf_5b354994d2d4bcf1.__obf_113e9a23713e45c1
	}
	if __obf_cad7cc30a9b8797e > 0 {
		__obf_43b169a18e2edb05 = time.Now().Add(__obf_cad7cc30a9b8797e).UnixNano()
	}
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = Item{
		Object:     __obf_5c75aa5f1a5ca219,
		Expiration: __obf_43b169a18e2edb05,
	}
	__obf_5b354994d2d4bcf1.
		// TODO: Calls to mu.Unlock are currently not deferred because defer
		// adds ~200 ns (as of go1.)
		__obf_d8684d60a98e1a8c.
		Unlock()
}

func (__obf_5b354994d2d4bcf1 *MapCache) __obf_37106b9a1b5b0c1f(__obf_ebdd4526b9b834d1 string, __obf_5c75aa5f1a5ca219 any, __obf_cad7cc30a9b8797e time.Duration) {
	var __obf_43b169a18e2edb05 int64
	if __obf_cad7cc30a9b8797e == DefaultExpiration {
		__obf_cad7cc30a9b8797e = __obf_5b354994d2d4bcf1.__obf_113e9a23713e45c1
	}
	if __obf_cad7cc30a9b8797e > 0 {
		__obf_43b169a18e2edb05 = time.Now().Add(__obf_cad7cc30a9b8797e).UnixNano()
	}
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = Item{
		Object:     __obf_5c75aa5f1a5ca219,
		Expiration: __obf_43b169a18e2edb05,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_5b354994d2d4bcf1 *MapCache) SetDefault(__obf_ebdd4526b9b834d1 string, __obf_5c75aa5f1a5ca219 any) {
	__obf_5b354994d2d4bcf1.
		Set(__obf_ebdd4526b9b834d1, __obf_5c75aa5f1a5ca219, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_5b354994d2d4bcf1 *MapCache) Add(__obf_ebdd4526b9b834d1 string, __obf_5c75aa5f1a5ca219 any, __obf_cad7cc30a9b8797e time.Duration) error {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	_, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_5a909fe103f858fd(__obf_ebdd4526b9b834d1)
	if __obf_bbc8a02c23dc299e {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return fmt.Errorf("Item %s already exists", __obf_ebdd4526b9b834d1)
	}
	__obf_5b354994d2d4bcf1.__obf_37106b9a1b5b0c1f(__obf_ebdd4526b9b834d1, __obf_5c75aa5f1a5ca219,

		// Set a new value for the cache key only if it already exists, and the existing
		// item hasn't expired. Returns an error otherwise.
		__obf_cad7cc30a9b8797e)
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Unlock()
	return nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) Replace(__obf_ebdd4526b9b834d1 string, __obf_5c75aa5f1a5ca219 any, __obf_cad7cc30a9b8797e time.Duration) error {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	_, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_5a909fe103f858fd(__obf_ebdd4526b9b834d1)
	if !__obf_bbc8a02c23dc299e {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_ebdd4526b9b834d1)
	}
	__obf_5b354994d2d4bcf1.__obf_37106b9a1b5b0c1f(__obf_ebdd4526b9b834d1, __obf_5c75aa5f1a5ca219,

		// Get an item from the cache. Returns the item or nil, and a bool indicating
		// whether the key was found.
		__obf_cad7cc30a9b8797e)
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Unlock()
	return nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) Get(__obf_ebdd4526b9b834d1 string) (any, bool) {
	__obf_5b354994d2d4bcf1.

		// "Inlining" of get and Expired
		__obf_d8684d60a98e1a8c.
		RLock()
	__obf_64869904d1dc9f43, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			RUnlock()
		return nil, false
	}
	if __obf_64869904d1dc9f43.Expiration > 0 {
		if time.Now().UnixNano() > __obf_64869904d1dc9f43.Expiration {
			__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
				RUnlock()
			return nil, false
		}
	}
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		RUnlock()
	return __obf_64869904d1dc9f43.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_5b354994d2d4bcf1 *MapCache) GetWithExpiration(__obf_ebdd4526b9b834d1 string) (any, time.Time, bool) {
	__obf_5b354994d2d4bcf1.

		// "Inlining" of get and Expired
		__obf_d8684d60a98e1a8c.
		RLock()
	__obf_64869904d1dc9f43, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_64869904d1dc9f43.Expiration > 0 {
		if time.Now().UnixNano() > __obf_64869904d1dc9f43.Expiration {
			__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
				RUnlock()
			return nil, time.Time{}, false
		}
		__obf_5b354994d2d4bcf1.

			// Return the item and the expiration time
			__obf_d8684d60a98e1a8c.
			RUnlock()
		return __obf_64869904d1dc9f43.Object, time.Unix(0, __obf_64869904d1dc9f43.Expiration), true
	}
	__obf_5b354994d2d4bcf1.

		// If expiration <= 0 (i.e. no expiration time set) then return the item
		// and a zeroed time.Time
		__obf_d8684d60a98e1a8c.
		RUnlock()
	return __obf_64869904d1dc9f43.Object, time.Time{}, true
}

func (__obf_5b354994d2d4bcf1 *MapCache) __obf_5a909fe103f858fd(__obf_ebdd4526b9b834d1 string) (any, bool) {
	__obf_64869904d1dc9f43, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_64869904d1dc9f43.Expiration > 0 {
		if time.Now().UnixNano() > __obf_64869904d1dc9f43.Expiration {
			return nil, false
		}
	}
	return __obf_64869904d1dc9f43.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_5b354994d2d4bcf1 *MapCache) Increment(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int64) error {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	switch __obf_fe4eccc509e2ba8c.Object.(type) {
	case int:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(int) + int(__obf_2b5886b254c22600)
	case int8:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(int8) + int8(__obf_2b5886b254c22600)
	case int16:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(int16) + int16(__obf_2b5886b254c22600)
	case int32:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(int32) + int32(__obf_2b5886b254c22600)
	case int64:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(int64) + __obf_2b5886b254c22600
	case uint:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uint) + uint(__obf_2b5886b254c22600)
	case uintptr:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uintptr) + uintptr(__obf_2b5886b254c22600)
	case uint8:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uint8) + uint8(__obf_2b5886b254c22600)
	case uint16:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uint16) + uint16(__obf_2b5886b254c22600)
	case uint32:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uint32) + uint32(__obf_2b5886b254c22600)
	case uint64:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uint64) + uint64(__obf_2b5886b254c22600)
	case float32:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(float32) + float32(__obf_2b5886b254c22600)
	case float64:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(float64) + float64(__obf_2b5886b254c22600)
	default:
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_ebdd4526b9b834d1)
	}
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to increment it by n. Pass a negative number to decrement the
		// value. To retrieve the incremented value, use one of the specialized methods,
		// e.g. IncrementFloat64.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementFloat(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 float64) error {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	switch __obf_fe4eccc509e2ba8c.Object.(type) {
	case float32:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(float32) + float32(__obf_2b5886b254c22600)
	case float64:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(float64) + __obf_2b5886b254c22600
	default:
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_ebdd4526b9b834d1)
	}
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementInt(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int) (int, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(int)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementInt8(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int8) (int8, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(int8)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementInt16(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int16) (int16, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(int16)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementInt32(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int32) (int32, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(int32)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementInt64(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int64) (int64, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(int64)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementUint(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uint) (uint, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uint)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementUintptr(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uintptr) (uintptr, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uintptr)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type uint8 by n. Returns an error if the item's value
		// is not an uint8, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementUint8(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uint8) (uint8, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uint8)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementUint16(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uint16) (uint16, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uint16)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementUint32(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uint32) (uint32, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uint32)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementUint64(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uint64) (uint64, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uint64)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementFloat32(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 float32) (float32, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(float32)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Increment an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) IncrementFloat64(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 float64) (float64, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(float64)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 + __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
		// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
		// item's value is not an integer, if it was not found, or if it is not
		// possible to decrement it by n. To retrieve the decremented value, use one
		// of the specialized methods, e.g. DecrementInt64.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) Decrement(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int64) error {
	__obf_5b354994d2d4bcf1.
		// TODO: Implement Increment and Decrement more cleanly.
		// (Cannot do Increment(k, n*-1) for uints.)
		__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_fe4eccc509e2ba8c.Object.(type) {
	case int:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(int) - int(__obf_2b5886b254c22600)
	case int8:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(int8) - int8(__obf_2b5886b254c22600)
	case int16:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(int16) - int16(__obf_2b5886b254c22600)
	case int32:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(int32) - int32(__obf_2b5886b254c22600)
	case int64:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(int64) - __obf_2b5886b254c22600
	case uint:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uint) - uint(__obf_2b5886b254c22600)
	case uintptr:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uintptr) - uintptr(__obf_2b5886b254c22600)
	case uint8:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uint8) - uint8(__obf_2b5886b254c22600)
	case uint16:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uint16) - uint16(__obf_2b5886b254c22600)
	case uint32:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uint32) - uint32(__obf_2b5886b254c22600)
	case uint64:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(uint64) - uint64(__obf_2b5886b254c22600)
	case float32:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(float32) - float32(__obf_2b5886b254c22600)
	case float64:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(float64) - float64(__obf_2b5886b254c22600)
	default:
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_ebdd4526b9b834d1)
	}
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to decrement it by n. Pass a negative number to decrement the
		// value. To retrieve the decremented value, use one of the specialized methods,
		// e.g. DecrementFloat64.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementFloat(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 float64) error {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	switch __obf_fe4eccc509e2ba8c.Object.(type) {
	case float32:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(float32) - float32(__obf_2b5886b254c22600)
	case float64:
		__obf_fe4eccc509e2ba8c.
			Object = __obf_fe4eccc509e2ba8c.Object.(float64) - __obf_2b5886b254c22600
	default:
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_ebdd4526b9b834d1)
	}
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementInt(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int) (int, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(int)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementInt8(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int8) (int8, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(int8)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementInt16(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int16) (int16, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(int16)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementInt32(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int32) (int32, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(int32)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementInt64(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 int64) (int64, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(int64)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementUint(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uint) (uint, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uint)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementUintptr(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uintptr) (uintptr, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uintptr)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type uint8 by n. Returns an error if the item's value is
		// not an uint8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementUint8(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uint8) (uint8, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uint8)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementUint16(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uint16) (uint16, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uint16)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementUint32(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uint32) (uint32, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uint32)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementUint64(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 uint64) (uint64, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(uint64)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementFloat32(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 float32) (float32, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(float32)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Decrement an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) DecrementFloat64(__obf_ebdd4526b9b834d1 string, __obf_2b5886b254c22600 float64) (float64, error) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
	if !__obf_bbc8a02c23dc299e || __obf_fe4eccc509e2ba8c.Expired() {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ebdd4526b9b834d1)
	}
	__obf_eae05f58ae327cd2, __obf_cddbfb0aefdf4145 := __obf_fe4eccc509e2ba8c.Object.(float64)
	if !__obf_cddbfb0aefdf4145 {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_ebdd4526b9b834d1)
	}
	__obf_3649b3b5ddd0b17f := __obf_eae05f58ae327cd2 - __obf_2b5886b254c22600
	__obf_fe4eccc509e2ba8c.
		Object = __obf_3649b3b5ddd0b17f
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	__obf_5b354994d2d4bcf1.

		// Delete an item from the cache. Does nothing if the key is not in the cache.
		__obf_d8684d60a98e1a8c.
		Unlock()
	return __obf_3649b3b5ddd0b17f, nil
}

func (__obf_5b354994d2d4bcf1 *MapCache) Delete(__obf_ebdd4526b9b834d1 string) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_fe4eccc509e2ba8c, __obf_623bffab51b760c4 := __obf_5b354994d2d4bcf1.delete(__obf_ebdd4526b9b834d1)
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Unlock()
	if __obf_623bffab51b760c4 {
		__obf_5b354994d2d4bcf1.__obf_2f7516313eb1f286(__obf_ebdd4526b9b834d1, __obf_fe4eccc509e2ba8c)
	}
}

func (__obf_5b354994d2d4bcf1 *MapCache) delete(__obf_ebdd4526b9b834d1 string) (any, bool) {
	if __obf_5b354994d2d4bcf1.__obf_2f7516313eb1f286 != nil {
		if __obf_fe4eccc509e2ba8c, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]; __obf_bbc8a02c23dc299e {
			delete(__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98, __obf_ebdd4526b9b834d1)
			return __obf_fe4eccc509e2ba8c.Object, true
		}
	}
	delete(__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98, __obf_ebdd4526b9b834d1)
	return nil, false
}

type __obf_d92bf28da7b167f5 struct {
	__obf_3405c14f70aaa4d0 string
	__obf_20c3560b9ea02863 any
}

// Delete all expired items from the cache.
func (__obf_5b354994d2d4bcf1 *MapCache) DeleteExpired() {
	var __obf_59c6f51ae9717b32 []__obf_d92bf28da7b167f5
	__obf_a083ea27f648e6e9 := time.Now().UnixNano()
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	for __obf_ebdd4526b9b834d1, __obf_fe4eccc509e2ba8c := range __obf_5b354994d2d4bcf1.
		// "Inlining" of expired
		__obf_717158131d1e5e98 {

		if __obf_fe4eccc509e2ba8c.Expiration > 0 && __obf_a083ea27f648e6e9 > __obf_fe4eccc509e2ba8c.Expiration {
			__obf_2bac026612517f46, __obf_623bffab51b760c4 := __obf_5b354994d2d4bcf1.delete(__obf_ebdd4526b9b834d1)
			if __obf_623bffab51b760c4 {
				__obf_59c6f51ae9717b32 = append(__obf_59c6f51ae9717b32, __obf_d92bf28da7b167f5{__obf_ebdd4526b9b834d1, __obf_2bac026612517f46})
			}
		}
	}
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Unlock()
	for _, __obf_fe4eccc509e2ba8c := range __obf_59c6f51ae9717b32 {
		__obf_5b354994d2d4bcf1.__obf_2f7516313eb1f286(__obf_fe4eccc509e2ba8c.__obf_3405c14f70aaa4d0,

			// Sets an (optional) function that is called with the key and value when an
			// item is evicted from the cache. (Including when it is deleted manually, but
			// not when it is overwritten.) Set to nil to disable.
			__obf_fe4eccc509e2ba8c.__obf_20c3560b9ea02863)
	}
}

func (__obf_5b354994d2d4bcf1 *MapCache) OnEvicted(__obf_8d4805f0b73cf3cb func(string, any)) {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_5b354994d2d4bcf1.__obf_2f7516313eb1f286 = __obf_8d4805f0b73cf3cb
	__obf_5b354994d2d4bcf1.

		// Write the cache's items (using Gob) to an io.Writer.
		//
		// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
		// documentation for NewFrom().)
		__obf_d8684d60a98e1a8c.
		Unlock()
}

func (__obf_5b354994d2d4bcf1 *MapCache) Save(__obf_f4b55b696e5f4905 io.Writer) (__obf_8c6b39ef87b4061f error) {
	__obf_a00d7036ec3b1731 := gob.NewEncoder(__obf_f4b55b696e5f4905)
	defer func() {
		if __obf_5c75aa5f1a5ca219 := recover(); __obf_5c75aa5f1a5ca219 != nil {
			__obf_8c6b39ef87b4061f = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		RLock()
	defer __obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.RUnlock()
	for _, __obf_fe4eccc509e2ba8c := range __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98 {
		gob.Register(__obf_fe4eccc509e2ba8c.Object)
	}
	__obf_8c6b39ef87b4061f = __obf_a00d7036ec3b1731.Encode(&__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_5b354994d2d4bcf1 *MapCache) SaveFile(__obf_7a6be275b828c84c string) error {
	__obf_e78f50d97182661a, __obf_8c6b39ef87b4061f := os.Create(__obf_7a6be275b828c84c)
	if __obf_8c6b39ef87b4061f != nil {
		return __obf_8c6b39ef87b4061f
	}
	__obf_8c6b39ef87b4061f = __obf_5b354994d2d4bcf1.Save(__obf_e78f50d97182661a)
	if __obf_8c6b39ef87b4061f != nil {
		__obf_e78f50d97182661a.
			Close()
		return __obf_8c6b39ef87b4061f
	}
	return __obf_e78f50d97182661a.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_5b354994d2d4bcf1 *MapCache) Load(__obf_f3a80f3795c3e8d5 io.Reader) error {
	__obf_69c075eafbe6aec2 := gob.NewDecoder(__obf_f3a80f3795c3e8d5)
	__obf_717158131d1e5e98 := map[string]Item{}
	__obf_8c6b39ef87b4061f := __obf_69c075eafbe6aec2.Decode(&__obf_717158131d1e5e98)
	if __obf_8c6b39ef87b4061f == nil {
		__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
			Lock()
		defer __obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.Unlock()
		for __obf_ebdd4526b9b834d1, __obf_fe4eccc509e2ba8c := range __obf_717158131d1e5e98 {
			__obf_2bac026612517f46, __obf_bbc8a02c23dc299e := __obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1]
			if !__obf_bbc8a02c23dc299e || __obf_2bac026612517f46.Expired() {
				__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
			}
		}
	}
	return __obf_8c6b39ef87b4061f
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_5b354994d2d4bcf1 *MapCache) LoadFile(__obf_7a6be275b828c84c string) error {
	__obf_e78f50d97182661a, __obf_8c6b39ef87b4061f := os.Open(__obf_7a6be275b828c84c)
	if __obf_8c6b39ef87b4061f != nil {
		return __obf_8c6b39ef87b4061f
	}
	__obf_8c6b39ef87b4061f = __obf_5b354994d2d4bcf1.Load(__obf_e78f50d97182661a)
	if __obf_8c6b39ef87b4061f != nil {
		__obf_e78f50d97182661a.
			Close()
		return __obf_8c6b39ef87b4061f
	}
	return __obf_e78f50d97182661a.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_5b354994d2d4bcf1 *MapCache) Items() map[string]Item {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		RLock()
	defer __obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.RUnlock()
	__obf_885f88c1b206dbea := make(map[string]Item, len(__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98))
	__obf_a083ea27f648e6e9 := time.Now().UnixNano()
	for __obf_ebdd4526b9b834d1, __obf_fe4eccc509e2ba8c := range __obf_5b354994d2d4bcf1.
		// "Inlining" of Expired
		__obf_717158131d1e5e98 {

		if __obf_fe4eccc509e2ba8c.Expiration > 0 {
			if __obf_a083ea27f648e6e9 > __obf_fe4eccc509e2ba8c.Expiration {
				continue
			}
		}
		__obf_885f88c1b206dbea[__obf_ebdd4526b9b834d1] = __obf_fe4eccc509e2ba8c
	}
	return __obf_885f88c1b206dbea
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_5b354994d2d4bcf1 *MapCache) ItemCount() int {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		RLock()
	__obf_2b5886b254c22600 := len(__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98)
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		RUnlock()
	return __obf_2b5886b254c22600
}

// Delete all items from the cache.
func (__obf_5b354994d2d4bcf1 *MapCache) Flush() {
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Lock()
	__obf_5b354994d2d4bcf1.__obf_717158131d1e5e98 = map[string]Item{}
	__obf_5b354994d2d4bcf1.__obf_d8684d60a98e1a8c.
		Unlock()
}

type __obf_b14857644ea2159c struct {
	Interval               time.Duration
	__obf_4db19207dc6c2b45 chan bool
}

func (__obf_28f5603e4a9e37df *__obf_b14857644ea2159c) Run(__obf_5b354994d2d4bcf1 *MapCache) {
	__obf_1186e195ec614c36 := time.NewTicker(__obf_28f5603e4a9e37df.Interval)
	for {
		select {
		case <-__obf_1186e195ec614c36.C:
			__obf_5b354994d2d4bcf1.
				DeleteExpired()
		case <-__obf_28f5603e4a9e37df.__obf_4db19207dc6c2b45:
			__obf_1186e195ec614c36.
				Stop()
			return
		}
	}
}

func __obf_332562983dd7d3ad(__obf_5b354994d2d4bcf1 *MapCache) {
	__obf_5b354994d2d4bcf1.__obf_b14857644ea2159c.__obf_4db19207dc6c2b45 <- true
}

func __obf_43a3254f787aba60(__obf_5b354994d2d4bcf1 *MapCache, __obf_6d01edbf1aed1678 time.Duration) {
	__obf_28f5603e4a9e37df := &__obf_b14857644ea2159c{
		Interval: __obf_6d01edbf1aed1678, __obf_4db19207dc6c2b45: make(chan bool),
	}
	__obf_5b354994d2d4bcf1.__obf_b14857644ea2159c = __obf_28f5603e4a9e37df
	go __obf_28f5603e4a9e37df.Run(__obf_5b354994d2d4bcf1)
}

func __obf_960123d08fbebbc2(__obf_3945db68ae5aaa00 time.Duration, __obf_885f88c1b206dbea map[string]Item) *MapCache {
	if __obf_3945db68ae5aaa00 == 0 {
		__obf_3945db68ae5aaa00 = -1
	}
	__obf_5b354994d2d4bcf1 := &MapCache{__obf_113e9a23713e45c1: __obf_3945db68ae5aaa00, __obf_717158131d1e5e98: __obf_885f88c1b206dbea}
	return __obf_5b354994d2d4bcf1
}

func __obf_d7ac0beec78a8ed5(__obf_3945db68ae5aaa00 time.Duration, __obf_6d01edbf1aed1678 time.Duration, __obf_885f88c1b206dbea map[string]Item) *MapCache {
	__obf_5b354994d2d4bcf1 := __obf_960123d08fbebbc2(__obf_3945db68ae5aaa00,
		// This trick ensures that the janitor goroutine (which--granted it
		// was enabled--is running DeleteExpired on c forever) does not keep
		// the returned C object from being garbage collected. When it is
		// garbage collected, the finalizer stops the janitor goroutine, after
		// which c can be collected.
		// C := &MapCache{c}
		__obf_885f88c1b206dbea)

	if __obf_6d01edbf1aed1678 > 0 {
		__obf_43a3254f787aba60(__obf_5b354994d2d4bcf1, __obf_6d01edbf1aed1678)
		runtime.SetFinalizer(__obf_5b354994d2d4bcf1, __obf_332562983dd7d3ad)
	}
	return __obf_5b354994d2d4bcf1
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_113e9a23713e45c1, __obf_c716a6869b1c3396 time.Duration) *MapCache {
	__obf_717158131d1e5e98 := make(map[string]Item)
	return __obf_d7ac0beec78a8ed5(__obf_113e9a23713e45c1, __obf_c716a6869b1c3396, __obf_717158131d1e5e98)
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
func NewFrom(__obf_113e9a23713e45c1, __obf_c716a6869b1c3396 time.Duration, __obf_717158131d1e5e98 map[string]Item) *MapCache {
	return __obf_d7ac0beec78a8ed5(__obf_113e9a23713e45c1, __obf_c716a6869b1c3396, __obf_717158131d1e5e98)
}
