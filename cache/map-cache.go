package __obf_a05682be1c6caf18

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
func (__obf_0bb9085921bdea16 Item) Expired() bool {
	if __obf_0bb9085921bdea16.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_0bb9085921bdea16.Expiration
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
	__obf_29a3a77862d6cd97 time. // *cache
		// If this is confusing, see the comment at the bottom of New()
		Duration
	__obf_8b7eb177de0a1f38 map[string]Item
	__obf_04ccec3be03c0675 sync.RWMutex
	__obf_21b4f1ae81335d24 func(string, any)
	__obf_9fe35822577b52fe *__obf_9fe35822577b52fe
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
func (__obf_365ba253fb24199c *MapCache) Set(__obf_b0c62418c29009a4 string, __obf_979ed7a4b7626a7e any, __obf_7e5fcb88db4cded2 time.Duration) {
	// "Inlining" of set
	var __obf_1086d0f8dfd43ca1 int64
	if __obf_7e5fcb88db4cded2 == DefaultExpiration {
		__obf_7e5fcb88db4cded2 = __obf_365ba253fb24199c.__obf_29a3a77862d6cd97
	}
	if __obf_7e5fcb88db4cded2 > 0 {
		__obf_1086d0f8dfd43ca1 = time.Now().Add(__obf_7e5fcb88db4cded2).UnixNano()
	}
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = Item{
		Object:     __obf_979ed7a4b7626a7e,
		Expiration: __obf_1086d0f8dfd43ca1,
	}
	__obf_365ba253fb24199c.
		// TODO: Calls to mu.Unlock are currently not deferred because defer
		// adds ~200 ns (as of go1.)
		__obf_04ccec3be03c0675.
		Unlock()
}

func (__obf_365ba253fb24199c *MapCache) __obf_f5c46f0cb915dd0c(__obf_b0c62418c29009a4 string, __obf_979ed7a4b7626a7e any, __obf_7e5fcb88db4cded2 time.Duration) {
	var __obf_1086d0f8dfd43ca1 int64
	if __obf_7e5fcb88db4cded2 == DefaultExpiration {
		__obf_7e5fcb88db4cded2 = __obf_365ba253fb24199c.__obf_29a3a77862d6cd97
	}
	if __obf_7e5fcb88db4cded2 > 0 {
		__obf_1086d0f8dfd43ca1 = time.Now().Add(__obf_7e5fcb88db4cded2).UnixNano()
	}
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = Item{
		Object:     __obf_979ed7a4b7626a7e,
		Expiration: __obf_1086d0f8dfd43ca1,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_365ba253fb24199c *MapCache) SetDefault(__obf_b0c62418c29009a4 string, __obf_979ed7a4b7626a7e any) {
	__obf_365ba253fb24199c.
		Set(__obf_b0c62418c29009a4, __obf_979ed7a4b7626a7e, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_365ba253fb24199c *MapCache) Add(__obf_b0c62418c29009a4 string, __obf_979ed7a4b7626a7e any, __obf_7e5fcb88db4cded2 time.Duration) error {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	_, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_d2195395afc8e655(__obf_b0c62418c29009a4)
	if __obf_83ebd96a05d52c0b {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return fmt.Errorf("Item %s already exists", __obf_b0c62418c29009a4)
	}
	__obf_365ba253fb24199c.__obf_f5c46f0cb915dd0c(__obf_b0c62418c29009a4, __obf_979ed7a4b7626a7e,

		// Set a new value for the cache key only if it already exists, and the existing
		// item hasn't expired. Returns an error otherwise.
		__obf_7e5fcb88db4cded2)
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Unlock()
	return nil
}

func (__obf_365ba253fb24199c *MapCache) Replace(__obf_b0c62418c29009a4 string, __obf_979ed7a4b7626a7e any, __obf_7e5fcb88db4cded2 time.Duration) error {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	_, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_d2195395afc8e655(__obf_b0c62418c29009a4)
	if !__obf_83ebd96a05d52c0b {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_b0c62418c29009a4)
	}
	__obf_365ba253fb24199c.__obf_f5c46f0cb915dd0c(__obf_b0c62418c29009a4, __obf_979ed7a4b7626a7e,

		// Get an item from the cache. Returns the item or nil, and a bool indicating
		// whether the key was found.
		__obf_7e5fcb88db4cded2)
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Unlock()
	return nil
}

func (__obf_365ba253fb24199c *MapCache) Get(__obf_b0c62418c29009a4 string) (any, bool) {
	__obf_365ba253fb24199c.

		// "Inlining" of get and Expired
		__obf_04ccec3be03c0675.
		RLock()
	__obf_0bb9085921bdea16, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			RUnlock()
		return nil, false
	}
	if __obf_0bb9085921bdea16.Expiration > 0 {
		if time.Now().UnixNano() > __obf_0bb9085921bdea16.Expiration {
			__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
				RUnlock()
			return nil, false
		}
	}
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		RUnlock()
	return __obf_0bb9085921bdea16.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_365ba253fb24199c *MapCache) GetWithExpiration(__obf_b0c62418c29009a4 string) (any, time.Time, bool) {
	__obf_365ba253fb24199c.

		// "Inlining" of get and Expired
		__obf_04ccec3be03c0675.
		RLock()
	__obf_0bb9085921bdea16, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_0bb9085921bdea16.Expiration > 0 {
		if time.Now().UnixNano() > __obf_0bb9085921bdea16.Expiration {
			__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
				RUnlock()
			return nil, time.Time{}, false
		}
		__obf_365ba253fb24199c.

			// Return the item and the expiration time
			__obf_04ccec3be03c0675.
			RUnlock()
		return __obf_0bb9085921bdea16.Object, time.Unix(0, __obf_0bb9085921bdea16.Expiration), true
	}
	__obf_365ba253fb24199c.

		// If expiration <= 0 (i.e. no expiration time set) then return the item
		// and a zeroed time.Time
		__obf_04ccec3be03c0675.
		RUnlock()
	return __obf_0bb9085921bdea16.Object, time.Time{}, true
}

func (__obf_365ba253fb24199c *MapCache) __obf_d2195395afc8e655(__obf_b0c62418c29009a4 string) (any, bool) {
	__obf_0bb9085921bdea16, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_0bb9085921bdea16.Expiration > 0 {
		if time.Now().UnixNano() > __obf_0bb9085921bdea16.Expiration {
			return nil, false
		}
	}
	return __obf_0bb9085921bdea16.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_365ba253fb24199c *MapCache) Increment(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int64) error {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	switch __obf_40394ecac45b9c79.Object.(type) {
	case int:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(int) + int(__obf_e2a6516886eac1c4)
	case int8:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(int8) + int8(__obf_e2a6516886eac1c4)
	case int16:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(int16) + int16(__obf_e2a6516886eac1c4)
	case int32:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(int32) + int32(__obf_e2a6516886eac1c4)
	case int64:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(int64) + __obf_e2a6516886eac1c4
	case uint:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uint) + uint(__obf_e2a6516886eac1c4)
	case uintptr:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uintptr) + uintptr(__obf_e2a6516886eac1c4)
	case uint8:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uint8) + uint8(__obf_e2a6516886eac1c4)
	case uint16:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uint16) + uint16(__obf_e2a6516886eac1c4)
	case uint32:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uint32) + uint32(__obf_e2a6516886eac1c4)
	case uint64:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uint64) + uint64(__obf_e2a6516886eac1c4)
	case float32:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(float32) + float32(__obf_e2a6516886eac1c4)
	case float64:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(float64) + float64(__obf_e2a6516886eac1c4)
	default:
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_b0c62418c29009a4)
	}
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to increment it by n. Pass a negative number to decrement the
		// value. To retrieve the incremented value, use one of the specialized methods,
		// e.g. IncrementFloat64.
		__obf_04ccec3be03c0675.
		Unlock()
	return nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementFloat(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 float64) error {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	switch __obf_40394ecac45b9c79.Object.(type) {
	case float32:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(float32) + float32(__obf_e2a6516886eac1c4)
	case float64:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(float64) + __obf_e2a6516886eac1c4
	default:
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_b0c62418c29009a4)
	}
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementInt(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int) (int, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(int)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementInt8(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int8) (int8, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(int8)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementInt16(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int16) (int16, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(int16)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementInt32(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int32) (int32, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(int32)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementInt64(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int64) (int64, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(int64)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementUint(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uint) (uint, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uint)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementUintptr(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uintptr) (uintptr, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uintptr)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type uint8 by n. Returns an error if the item's value
		// is not an uint8, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementUint8(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uint8) (uint8, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uint8)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementUint16(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uint16) (uint16, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uint16)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementUint32(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uint32) (uint32, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uint32)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementUint64(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uint64) (uint64, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uint64)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementFloat32(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 float32) (float32, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(float32)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Increment an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) IncrementFloat64(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 float64) (float64, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(float64)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb + __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
		// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
		// item's value is not an integer, if it was not found, or if it is not
		// possible to decrement it by n. To retrieve the decremented value, use one
		// of the specialized methods, e.g. DecrementInt64.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) Decrement(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int64) error {
	__obf_365ba253fb24199c.
		// TODO: Implement Increment and Decrement more cleanly.
		// (Cannot do Increment(k, n*-1) for uints.)
		__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_40394ecac45b9c79.Object.(type) {
	case int:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(int) - int(__obf_e2a6516886eac1c4)
	case int8:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(int8) - int8(__obf_e2a6516886eac1c4)
	case int16:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(int16) - int16(__obf_e2a6516886eac1c4)
	case int32:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(int32) - int32(__obf_e2a6516886eac1c4)
	case int64:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(int64) - __obf_e2a6516886eac1c4
	case uint:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uint) - uint(__obf_e2a6516886eac1c4)
	case uintptr:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uintptr) - uintptr(__obf_e2a6516886eac1c4)
	case uint8:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uint8) - uint8(__obf_e2a6516886eac1c4)
	case uint16:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uint16) - uint16(__obf_e2a6516886eac1c4)
	case uint32:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uint32) - uint32(__obf_e2a6516886eac1c4)
	case uint64:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(uint64) - uint64(__obf_e2a6516886eac1c4)
	case float32:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(float32) - float32(__obf_e2a6516886eac1c4)
	case float64:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(float64) - float64(__obf_e2a6516886eac1c4)
	default:
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_b0c62418c29009a4)
	}
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to decrement it by n. Pass a negative number to decrement the
		// value. To retrieve the decremented value, use one of the specialized methods,
		// e.g. DecrementFloat64.
		__obf_04ccec3be03c0675.
		Unlock()
	return nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementFloat(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 float64) error {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	switch __obf_40394ecac45b9c79.Object.(type) {
	case float32:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(float32) - float32(__obf_e2a6516886eac1c4)
	case float64:
		__obf_40394ecac45b9c79.
			Object = __obf_40394ecac45b9c79.Object.(float64) - __obf_e2a6516886eac1c4
	default:
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_b0c62418c29009a4)
	}
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementInt(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int) (int, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(int)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementInt8(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int8) (int8, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(int8)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementInt16(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int16) (int16, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(int16)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementInt32(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int32) (int32, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(int32)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementInt64(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 int64) (int64, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(int64)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementUint(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uint) (uint, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uint)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementUintptr(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uintptr) (uintptr, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uintptr)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type uint8 by n. Returns an error if the item's value is
		// not an uint8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementUint8(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uint8) (uint8, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uint8)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementUint16(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uint16) (uint16, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uint16)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementUint32(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uint32) (uint32, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uint32)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementUint64(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 uint64) (uint64, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(uint64)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementFloat32(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 float32) (float32, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(float32)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Decrement an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) DecrementFloat64(__obf_b0c62418c29009a4 string, __obf_e2a6516886eac1c4 float64) (float64, error) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
	if !__obf_83ebd96a05d52c0b || __obf_40394ecac45b9c79.Expired() {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_b0c62418c29009a4)
	}
	__obf_ea636629b5d77aeb, __obf_19211ff80aaad9b4 := __obf_40394ecac45b9c79.Object.(float64)
	if !__obf_19211ff80aaad9b4 {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_b0c62418c29009a4)
	}
	__obf_5de34280e081d218 := __obf_ea636629b5d77aeb - __obf_e2a6516886eac1c4
	__obf_40394ecac45b9c79.
		Object = __obf_5de34280e081d218
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	__obf_365ba253fb24199c.

		// Delete an item from the cache. Does nothing if the key is not in the cache.
		__obf_04ccec3be03c0675.
		Unlock()
	return __obf_5de34280e081d218, nil
}

func (__obf_365ba253fb24199c *MapCache) Delete(__obf_b0c62418c29009a4 string) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_40394ecac45b9c79, __obf_4743304b3246fd40 := __obf_365ba253fb24199c.delete(__obf_b0c62418c29009a4)
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Unlock()
	if __obf_4743304b3246fd40 {
		__obf_365ba253fb24199c.__obf_21b4f1ae81335d24(__obf_b0c62418c29009a4, __obf_40394ecac45b9c79)
	}
}

func (__obf_365ba253fb24199c *MapCache) delete(__obf_b0c62418c29009a4 string) (any, bool) {
	if __obf_365ba253fb24199c.__obf_21b4f1ae81335d24 != nil {
		if __obf_40394ecac45b9c79, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]; __obf_83ebd96a05d52c0b {
			delete(__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38, __obf_b0c62418c29009a4)
			return __obf_40394ecac45b9c79.Object, true
		}
	}
	delete(__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38, __obf_b0c62418c29009a4)
	return nil, false
}

type __obf_708703d8dbd6b76a struct {
	__obf_a0e56915a05e8d99 string
	__obf_106924f0283e3cd8 any
}

// Delete all expired items from the cache.
func (__obf_365ba253fb24199c *MapCache) DeleteExpired() {
	var __obf_026805bbb6779be7 []__obf_708703d8dbd6b76a
	__obf_bda7719e9501545d := time.Now().UnixNano()
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	for __obf_b0c62418c29009a4, __obf_40394ecac45b9c79 := range __obf_365ba253fb24199c.
		// "Inlining" of expired
		__obf_8b7eb177de0a1f38 {

		if __obf_40394ecac45b9c79.Expiration > 0 && __obf_bda7719e9501545d > __obf_40394ecac45b9c79.Expiration {
			__obf_1bbe1fa0395e69e1, __obf_4743304b3246fd40 := __obf_365ba253fb24199c.delete(__obf_b0c62418c29009a4)
			if __obf_4743304b3246fd40 {
				__obf_026805bbb6779be7 = append(__obf_026805bbb6779be7, __obf_708703d8dbd6b76a{__obf_b0c62418c29009a4, __obf_1bbe1fa0395e69e1})
			}
		}
	}
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Unlock()
	for _, __obf_40394ecac45b9c79 := range __obf_026805bbb6779be7 {
		__obf_365ba253fb24199c.__obf_21b4f1ae81335d24(__obf_40394ecac45b9c79.__obf_a0e56915a05e8d99,

			// Sets an (optional) function that is called with the key and value when an
			// item is evicted from the cache. (Including when it is deleted manually, but
			// not when it is overwritten.) Set to nil to disable.
			__obf_40394ecac45b9c79.__obf_106924f0283e3cd8)
	}
}

func (__obf_365ba253fb24199c *MapCache) OnEvicted(__obf_69411c8828870f29 func(string, any)) {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_365ba253fb24199c.__obf_21b4f1ae81335d24 = __obf_69411c8828870f29
	__obf_365ba253fb24199c.

		// Write the cache's items (using Gob) to an io.Writer.
		//
		// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
		// documentation for NewFrom().)
		__obf_04ccec3be03c0675.
		Unlock()
}

func (__obf_365ba253fb24199c *MapCache) Save(__obf_94ffec342aa4bca7 io.Writer) (__obf_94de95cb433b50f1 error) {
	__obf_542e8bfbb2403f45 := gob.NewEncoder(__obf_94ffec342aa4bca7)
	defer func() {
		if __obf_979ed7a4b7626a7e := recover(); __obf_979ed7a4b7626a7e != nil {
			__obf_94de95cb433b50f1 = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		RLock()
	defer __obf_365ba253fb24199c.__obf_04ccec3be03c0675.RUnlock()
	for _, __obf_40394ecac45b9c79 := range __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38 {
		gob.Register(__obf_40394ecac45b9c79.Object)
	}
	__obf_94de95cb433b50f1 = __obf_542e8bfbb2403f45.Encode(&__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_365ba253fb24199c *MapCache) SaveFile(__obf_657f0e07c675957f string) error {
	__obf_8708c93cb30b6dee, __obf_94de95cb433b50f1 := os.Create(__obf_657f0e07c675957f)
	if __obf_94de95cb433b50f1 != nil {
		return __obf_94de95cb433b50f1
	}
	__obf_94de95cb433b50f1 = __obf_365ba253fb24199c.Save(__obf_8708c93cb30b6dee)
	if __obf_94de95cb433b50f1 != nil {
		__obf_8708c93cb30b6dee.
			Close()
		return __obf_94de95cb433b50f1
	}
	return __obf_8708c93cb30b6dee.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_365ba253fb24199c *MapCache) Load(__obf_3f3723f674c63e34 io.Reader) error {
	__obf_2d9dbe1eb9d4b78b := gob.NewDecoder(__obf_3f3723f674c63e34)
	__obf_8b7eb177de0a1f38 := map[string]Item{}
	__obf_94de95cb433b50f1 := __obf_2d9dbe1eb9d4b78b.Decode(&__obf_8b7eb177de0a1f38)
	if __obf_94de95cb433b50f1 == nil {
		__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
			Lock()
		defer __obf_365ba253fb24199c.__obf_04ccec3be03c0675.Unlock()
		for __obf_b0c62418c29009a4, __obf_40394ecac45b9c79 := range __obf_8b7eb177de0a1f38 {
			__obf_1bbe1fa0395e69e1, __obf_83ebd96a05d52c0b := __obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4]
			if !__obf_83ebd96a05d52c0b || __obf_1bbe1fa0395e69e1.Expired() {
				__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
			}
		}
	}
	return __obf_94de95cb433b50f1
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_365ba253fb24199c *MapCache) LoadFile(__obf_657f0e07c675957f string) error {
	__obf_8708c93cb30b6dee, __obf_94de95cb433b50f1 := os.Open(__obf_657f0e07c675957f)
	if __obf_94de95cb433b50f1 != nil {
		return __obf_94de95cb433b50f1
	}
	__obf_94de95cb433b50f1 = __obf_365ba253fb24199c.Load(__obf_8708c93cb30b6dee)
	if __obf_94de95cb433b50f1 != nil {
		__obf_8708c93cb30b6dee.
			Close()
		return __obf_94de95cb433b50f1
	}
	return __obf_8708c93cb30b6dee.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_365ba253fb24199c *MapCache) Items() map[string]Item {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		RLock()
	defer __obf_365ba253fb24199c.__obf_04ccec3be03c0675.RUnlock()
	__obf_cf817044cb0e2a7b := make(map[string]Item, len(__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38))
	__obf_bda7719e9501545d := time.Now().UnixNano()
	for __obf_b0c62418c29009a4, __obf_40394ecac45b9c79 := range __obf_365ba253fb24199c.
		// "Inlining" of Expired
		__obf_8b7eb177de0a1f38 {

		if __obf_40394ecac45b9c79.Expiration > 0 {
			if __obf_bda7719e9501545d > __obf_40394ecac45b9c79.Expiration {
				continue
			}
		}
		__obf_cf817044cb0e2a7b[__obf_b0c62418c29009a4] = __obf_40394ecac45b9c79
	}
	return __obf_cf817044cb0e2a7b
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_365ba253fb24199c *MapCache) ItemCount() int {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		RLock()
	__obf_e2a6516886eac1c4 := len(__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38)
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		RUnlock()
	return __obf_e2a6516886eac1c4
}

// Delete all items from the cache.
func (__obf_365ba253fb24199c *MapCache) Flush() {
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Lock()
	__obf_365ba253fb24199c.__obf_8b7eb177de0a1f38 = map[string]Item{}
	__obf_365ba253fb24199c.__obf_04ccec3be03c0675.
		Unlock()
}

type __obf_9fe35822577b52fe struct {
	Interval               time.Duration
	__obf_85d4128797a92178 chan bool
}

func (__obf_37abed8ef92cb1a3 *__obf_9fe35822577b52fe) Run(__obf_365ba253fb24199c *MapCache) {
	__obf_f0b7604aa5585fc1 := time.NewTicker(__obf_37abed8ef92cb1a3.Interval)
	for {
		select {
		case <-__obf_f0b7604aa5585fc1.C:
			__obf_365ba253fb24199c.
				DeleteExpired()
		case <-__obf_37abed8ef92cb1a3.__obf_85d4128797a92178:
			__obf_f0b7604aa5585fc1.
				Stop()
			return
		}
	}
}

func __obf_688c980a5860bfff(__obf_365ba253fb24199c *MapCache) {
	__obf_365ba253fb24199c.__obf_9fe35822577b52fe.__obf_85d4128797a92178 <- true
}

func __obf_5cc774f5a96ce8d5(__obf_365ba253fb24199c *MapCache, __obf_36a13ebc2c511cf5 time.Duration) {
	__obf_37abed8ef92cb1a3 := &__obf_9fe35822577b52fe{
		Interval: __obf_36a13ebc2c511cf5, __obf_85d4128797a92178: make(chan bool),
	}
	__obf_365ba253fb24199c.__obf_9fe35822577b52fe = __obf_37abed8ef92cb1a3
	go __obf_37abed8ef92cb1a3.Run(__obf_365ba253fb24199c)
}

func __obf_d4860ce03ea3440c(__obf_4d58f23e4ee9539a time.Duration, __obf_cf817044cb0e2a7b map[string]Item) *MapCache {
	if __obf_4d58f23e4ee9539a == 0 {
		__obf_4d58f23e4ee9539a = -1
	}
	__obf_365ba253fb24199c := &MapCache{__obf_29a3a77862d6cd97: __obf_4d58f23e4ee9539a, __obf_8b7eb177de0a1f38: __obf_cf817044cb0e2a7b}
	return __obf_365ba253fb24199c
}

func __obf_50f9257c419c5b7b(__obf_4d58f23e4ee9539a time.Duration, __obf_36a13ebc2c511cf5 time.Duration, __obf_cf817044cb0e2a7b map[string]Item) *MapCache {
	__obf_365ba253fb24199c := __obf_d4860ce03ea3440c(__obf_4d58f23e4ee9539a,
		// This trick ensures that the janitor goroutine (which--granted it
		// was enabled--is running DeleteExpired on c forever) does not keep
		// the returned C object from being garbage collected. When it is
		// garbage collected, the finalizer stops the janitor goroutine, after
		// which c can be collected.
		// C := &MapCache{c}
		__obf_cf817044cb0e2a7b)

	if __obf_36a13ebc2c511cf5 > 0 {
		__obf_5cc774f5a96ce8d5(__obf_365ba253fb24199c, __obf_36a13ebc2c511cf5)
		runtime.SetFinalizer(__obf_365ba253fb24199c, __obf_688c980a5860bfff)
	}
	return __obf_365ba253fb24199c
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_29a3a77862d6cd97, __obf_e8c1ab90c946acde time.Duration) *MapCache {
	__obf_8b7eb177de0a1f38 := make(map[string]Item)
	return __obf_50f9257c419c5b7b(__obf_29a3a77862d6cd97, __obf_e8c1ab90c946acde, __obf_8b7eb177de0a1f38)
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
func NewFrom(__obf_29a3a77862d6cd97, __obf_e8c1ab90c946acde time.Duration, __obf_8b7eb177de0a1f38 map[string]Item) *MapCache {
	return __obf_50f9257c419c5b7b(__obf_29a3a77862d6cd97, __obf_e8c1ab90c946acde, __obf_8b7eb177de0a1f38)
}
