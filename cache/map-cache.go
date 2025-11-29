package __obf_6fd4fe34e3f784df

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
func (__obf_90b4f954b51e95a7 Item) Expired() bool {
	if __obf_90b4f954b51e95a7.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_90b4f954b51e95a7.Expiration
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
	__obf_0a3ee88650ce44a7 time. // *cache
		// If this is confusing, see the comment at the bottom of New()
		Duration
	__obf_0c0fad92722f4264 map[string]Item
	__obf_53983106240f1755 sync.RWMutex
	__obf_de42ab8bb23c2bee func(string, any)
	__obf_c39f2111b362a882 *__obf_c39f2111b362a882
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
func (__obf_3cd39772de5fd3ee *MapCache) Set(__obf_2a0493c028b94053 string, __obf_c9582df27c8088ea any, __obf_de081202d8435645 time.Duration) {
	// "Inlining" of set
	var __obf_0f22edc65e150c28 int64
	if __obf_de081202d8435645 == DefaultExpiration {
		__obf_de081202d8435645 = __obf_3cd39772de5fd3ee.__obf_0a3ee88650ce44a7
	}
	if __obf_de081202d8435645 > 0 {
		__obf_0f22edc65e150c28 = time.Now().Add(__obf_de081202d8435645).UnixNano()
	}
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = Item{
		Object:     __obf_c9582df27c8088ea,
		Expiration: __obf_0f22edc65e150c28,
	}
	__obf_3cd39772de5fd3ee.
		// TODO: Calls to mu.Unlock are currently not deferred because defer
		// adds ~200 ns (as of go1.)
		__obf_53983106240f1755.
		Unlock()
}

func (__obf_3cd39772de5fd3ee *MapCache) __obf_29584f996a520211(__obf_2a0493c028b94053 string, __obf_c9582df27c8088ea any, __obf_de081202d8435645 time.Duration) {
	var __obf_0f22edc65e150c28 int64
	if __obf_de081202d8435645 == DefaultExpiration {
		__obf_de081202d8435645 = __obf_3cd39772de5fd3ee.__obf_0a3ee88650ce44a7
	}
	if __obf_de081202d8435645 > 0 {
		__obf_0f22edc65e150c28 = time.Now().Add(__obf_de081202d8435645).UnixNano()
	}
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = Item{
		Object:     __obf_c9582df27c8088ea,
		Expiration: __obf_0f22edc65e150c28,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_3cd39772de5fd3ee *MapCache) SetDefault(__obf_2a0493c028b94053 string, __obf_c9582df27c8088ea any) {
	__obf_3cd39772de5fd3ee.
		Set(__obf_2a0493c028b94053, __obf_c9582df27c8088ea, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_3cd39772de5fd3ee *MapCache) Add(__obf_2a0493c028b94053 string, __obf_c9582df27c8088ea any, __obf_de081202d8435645 time.Duration) error {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	_, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_b51286810c5a4f46(__obf_2a0493c028b94053)
	if __obf_4228f4fcd56ac4f8 {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return fmt.Errorf("Item %s already exists", __obf_2a0493c028b94053)
	}
	__obf_3cd39772de5fd3ee.__obf_29584f996a520211(__obf_2a0493c028b94053, __obf_c9582df27c8088ea,

		// Set a new value for the cache key only if it already exists, and the existing
		// item hasn't expired. Returns an error otherwise.
		__obf_de081202d8435645)
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Unlock()
	return nil
}

func (__obf_3cd39772de5fd3ee *MapCache) Replace(__obf_2a0493c028b94053 string, __obf_c9582df27c8088ea any, __obf_de081202d8435645 time.Duration) error {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	_, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_b51286810c5a4f46(__obf_2a0493c028b94053)
	if !__obf_4228f4fcd56ac4f8 {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_2a0493c028b94053)
	}
	__obf_3cd39772de5fd3ee.__obf_29584f996a520211(__obf_2a0493c028b94053, __obf_c9582df27c8088ea,

		// Get an item from the cache. Returns the item or nil, and a bool indicating
		// whether the key was found.
		__obf_de081202d8435645)
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Unlock()
	return nil
}

func (__obf_3cd39772de5fd3ee *MapCache) Get(__obf_2a0493c028b94053 string) (any, bool) {
	__obf_3cd39772de5fd3ee.

		// "Inlining" of get and Expired
		__obf_53983106240f1755.
		RLock()
	__obf_90b4f954b51e95a7, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			RUnlock()
		return nil, false
	}
	if __obf_90b4f954b51e95a7.Expiration > 0 {
		if time.Now().UnixNano() > __obf_90b4f954b51e95a7.Expiration {
			__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
				RUnlock()
			return nil, false
		}
	}
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		RUnlock()
	return __obf_90b4f954b51e95a7.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_3cd39772de5fd3ee *MapCache) GetWithExpiration(__obf_2a0493c028b94053 string) (any, time.Time, bool) {
	__obf_3cd39772de5fd3ee.

		// "Inlining" of get and Expired
		__obf_53983106240f1755.
		RLock()
	__obf_90b4f954b51e95a7, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_90b4f954b51e95a7.Expiration > 0 {
		if time.Now().UnixNano() > __obf_90b4f954b51e95a7.Expiration {
			__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
				RUnlock()
			return nil, time.Time{}, false
		}
		__obf_3cd39772de5fd3ee.

			// Return the item and the expiration time
			__obf_53983106240f1755.
			RUnlock()
		return __obf_90b4f954b51e95a7.Object, time.Unix(0, __obf_90b4f954b51e95a7.Expiration), true
	}
	__obf_3cd39772de5fd3ee.

		// If expiration <= 0 (i.e. no expiration time set) then return the item
		// and a zeroed time.Time
		__obf_53983106240f1755.
		RUnlock()
	return __obf_90b4f954b51e95a7.Object, time.Time{}, true
}

func (__obf_3cd39772de5fd3ee *MapCache) __obf_b51286810c5a4f46(__obf_2a0493c028b94053 string) (any, bool) {
	__obf_90b4f954b51e95a7, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_90b4f954b51e95a7.Expiration > 0 {
		if time.Now().UnixNano() > __obf_90b4f954b51e95a7.Expiration {
			return nil, false
		}
	}
	return __obf_90b4f954b51e95a7.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_3cd39772de5fd3ee *MapCache) Increment(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int64) error {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	switch __obf_b1ecf81133de0ff6.Object.(type) {
	case int:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(int) + int(__obf_a091a23e60b9ef54)
	case int8:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(int8) + int8(__obf_a091a23e60b9ef54)
	case int16:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(int16) + int16(__obf_a091a23e60b9ef54)
	case int32:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(int32) + int32(__obf_a091a23e60b9ef54)
	case int64:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(int64) + __obf_a091a23e60b9ef54
	case uint:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uint) + uint(__obf_a091a23e60b9ef54)
	case uintptr:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uintptr) + uintptr(__obf_a091a23e60b9ef54)
	case uint8:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uint8) + uint8(__obf_a091a23e60b9ef54)
	case uint16:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uint16) + uint16(__obf_a091a23e60b9ef54)
	case uint32:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uint32) + uint32(__obf_a091a23e60b9ef54)
	case uint64:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uint64) + uint64(__obf_a091a23e60b9ef54)
	case float32:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(float32) + float32(__obf_a091a23e60b9ef54)
	case float64:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(float64) + float64(__obf_a091a23e60b9ef54)
	default:
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_2a0493c028b94053)
	}
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to increment it by n. Pass a negative number to decrement the
		// value. To retrieve the incremented value, use one of the specialized methods,
		// e.g. IncrementFloat64.
		__obf_53983106240f1755.
		Unlock()
	return nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementFloat(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 float64) error {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	switch __obf_b1ecf81133de0ff6.Object.(type) {
	case float32:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(float32) + float32(__obf_a091a23e60b9ef54)
	case float64:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(float64) + __obf_a091a23e60b9ef54
	default:
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_2a0493c028b94053)
	}
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementInt(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int) (int, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(int)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementInt8(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int8) (int8, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(int8)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementInt16(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int16) (int16, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(int16)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementInt32(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int32) (int32, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(int32)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementInt64(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int64) (int64, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(int64)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementUint(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uint) (uint, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uint)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementUintptr(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uintptr) (uintptr, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uintptr)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type uint8 by n. Returns an error if the item's value
		// is not an uint8, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementUint8(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uint8) (uint8, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uint8)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementUint16(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uint16) (uint16, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uint16)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementUint32(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uint32) (uint32, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uint32)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementUint64(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uint64) (uint64, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uint64)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementFloat32(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 float32) (float32, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(float32)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Increment an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) IncrementFloat64(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 float64) (float64, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(float64)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 + __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
		// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
		// item's value is not an integer, if it was not found, or if it is not
		// possible to decrement it by n. To retrieve the decremented value, use one
		// of the specialized methods, e.g. DecrementInt64.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) Decrement(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int64) error {
	__obf_3cd39772de5fd3ee.
		// TODO: Implement Increment and Decrement more cleanly.
		// (Cannot do Increment(k, n*-1) for uints.)
		__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_b1ecf81133de0ff6.Object.(type) {
	case int:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(int) - int(__obf_a091a23e60b9ef54)
	case int8:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(int8) - int8(__obf_a091a23e60b9ef54)
	case int16:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(int16) - int16(__obf_a091a23e60b9ef54)
	case int32:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(int32) - int32(__obf_a091a23e60b9ef54)
	case int64:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(int64) - __obf_a091a23e60b9ef54
	case uint:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uint) - uint(__obf_a091a23e60b9ef54)
	case uintptr:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uintptr) - uintptr(__obf_a091a23e60b9ef54)
	case uint8:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uint8) - uint8(__obf_a091a23e60b9ef54)
	case uint16:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uint16) - uint16(__obf_a091a23e60b9ef54)
	case uint32:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uint32) - uint32(__obf_a091a23e60b9ef54)
	case uint64:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(uint64) - uint64(__obf_a091a23e60b9ef54)
	case float32:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(float32) - float32(__obf_a091a23e60b9ef54)
	case float64:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(float64) - float64(__obf_a091a23e60b9ef54)
	default:
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_2a0493c028b94053)
	}
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to decrement it by n. Pass a negative number to decrement the
		// value. To retrieve the decremented value, use one of the specialized methods,
		// e.g. DecrementFloat64.
		__obf_53983106240f1755.
		Unlock()
	return nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementFloat(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 float64) error {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	switch __obf_b1ecf81133de0ff6.Object.(type) {
	case float32:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(float32) - float32(__obf_a091a23e60b9ef54)
	case float64:
		__obf_b1ecf81133de0ff6.
			Object = __obf_b1ecf81133de0ff6.Object.(float64) - __obf_a091a23e60b9ef54
	default:
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_2a0493c028b94053)
	}
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementInt(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int) (int, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(int)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementInt8(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int8) (int8, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(int8)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementInt16(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int16) (int16, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(int16)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementInt32(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int32) (int32, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(int32)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementInt64(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 int64) (int64, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(int64)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementUint(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uint) (uint, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uint)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementUintptr(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uintptr) (uintptr, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uintptr)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type uint8 by n. Returns an error if the item's value is
		// not an uint8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementUint8(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uint8) (uint8, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uint8)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementUint16(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uint16) (uint16, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uint16)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementUint32(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uint32) (uint32, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uint32)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementUint64(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 uint64) (uint64, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(uint64)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementFloat32(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 float32) (float32, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(float32)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Decrement an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) DecrementFloat64(__obf_2a0493c028b94053 string, __obf_a091a23e60b9ef54 float64) (float64, error) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
	if !__obf_4228f4fcd56ac4f8 || __obf_b1ecf81133de0ff6.Expired() {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_2a0493c028b94053)
	}
	__obf_79b96bbb9cac9854, __obf_949c761ffd90d5fa := __obf_b1ecf81133de0ff6.Object.(float64)
	if !__obf_949c761ffd90d5fa {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_2a0493c028b94053)
	}
	__obf_30ee4d62b3698152 := __obf_79b96bbb9cac9854 - __obf_a091a23e60b9ef54
	__obf_b1ecf81133de0ff6.
		Object = __obf_30ee4d62b3698152
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	__obf_3cd39772de5fd3ee.

		// Delete an item from the cache. Does nothing if the key is not in the cache.
		__obf_53983106240f1755.
		Unlock()
	return __obf_30ee4d62b3698152, nil
}

func (__obf_3cd39772de5fd3ee *MapCache) Delete(__obf_2a0493c028b94053 string) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_b1ecf81133de0ff6, __obf_7aa921d16239fdd7 := __obf_3cd39772de5fd3ee.delete(__obf_2a0493c028b94053)
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Unlock()
	if __obf_7aa921d16239fdd7 {
		__obf_3cd39772de5fd3ee.__obf_de42ab8bb23c2bee(__obf_2a0493c028b94053, __obf_b1ecf81133de0ff6)
	}
}

func (__obf_3cd39772de5fd3ee *MapCache) delete(__obf_2a0493c028b94053 string) (any, bool) {
	if __obf_3cd39772de5fd3ee.__obf_de42ab8bb23c2bee != nil {
		if __obf_b1ecf81133de0ff6, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]; __obf_4228f4fcd56ac4f8 {
			delete(__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264, __obf_2a0493c028b94053)
			return __obf_b1ecf81133de0ff6.Object, true
		}
	}
	delete(__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264, __obf_2a0493c028b94053)
	return nil, false
}

type __obf_2921bfb172b2987a struct {
	__obf_ca986a38d1f8fbbb string
	__obf_e35c6f9510020aab any
}

// Delete all expired items from the cache.
func (__obf_3cd39772de5fd3ee *MapCache) DeleteExpired() {
	var __obf_3916a1bdfe6850ce []__obf_2921bfb172b2987a
	__obf_8150b358628c5292 := time.Now().UnixNano()
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	for __obf_2a0493c028b94053, __obf_b1ecf81133de0ff6 := range __obf_3cd39772de5fd3ee.
		// "Inlining" of expired
		__obf_0c0fad92722f4264 {

		if __obf_b1ecf81133de0ff6.Expiration > 0 && __obf_8150b358628c5292 > __obf_b1ecf81133de0ff6.Expiration {
			__obf_cd9346a423f5ce9d, __obf_7aa921d16239fdd7 := __obf_3cd39772de5fd3ee.delete(__obf_2a0493c028b94053)
			if __obf_7aa921d16239fdd7 {
				__obf_3916a1bdfe6850ce = append(__obf_3916a1bdfe6850ce, __obf_2921bfb172b2987a{__obf_2a0493c028b94053, __obf_cd9346a423f5ce9d})
			}
		}
	}
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Unlock()
	for _, __obf_b1ecf81133de0ff6 := range __obf_3916a1bdfe6850ce {
		__obf_3cd39772de5fd3ee.__obf_de42ab8bb23c2bee(__obf_b1ecf81133de0ff6.__obf_ca986a38d1f8fbbb,

			// Sets an (optional) function that is called with the key and value when an
			// item is evicted from the cache. (Including when it is deleted manually, but
			// not when it is overwritten.) Set to nil to disable.
			__obf_b1ecf81133de0ff6.__obf_e35c6f9510020aab)
	}
}

func (__obf_3cd39772de5fd3ee *MapCache) OnEvicted(__obf_d3bd16c592c47888 func(string, any)) {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_3cd39772de5fd3ee.__obf_de42ab8bb23c2bee = __obf_d3bd16c592c47888
	__obf_3cd39772de5fd3ee.

		// Write the cache's items (using Gob) to an io.Writer.
		//
		// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
		// documentation for NewFrom().)
		__obf_53983106240f1755.
		Unlock()
}

func (__obf_3cd39772de5fd3ee *MapCache) Save(__obf_7e941a6263b99ea3 io.Writer) (__obf_7f8f5f5c173fa58e error) {
	__obf_4ab283981a4deb49 := gob.NewEncoder(__obf_7e941a6263b99ea3)
	defer func() {
		if __obf_c9582df27c8088ea := recover(); __obf_c9582df27c8088ea != nil {
			__obf_7f8f5f5c173fa58e = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		RLock()
	defer __obf_3cd39772de5fd3ee.__obf_53983106240f1755.RUnlock()
	for _, __obf_b1ecf81133de0ff6 := range __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264 {
		gob.Register(__obf_b1ecf81133de0ff6.Object)
	}
	__obf_7f8f5f5c173fa58e = __obf_4ab283981a4deb49.Encode(&__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_3cd39772de5fd3ee *MapCache) SaveFile(__obf_6d6606cc565f833e string) error {
	__obf_52e794309333fcfb, __obf_7f8f5f5c173fa58e := os.Create(__obf_6d6606cc565f833e)
	if __obf_7f8f5f5c173fa58e != nil {
		return __obf_7f8f5f5c173fa58e
	}
	__obf_7f8f5f5c173fa58e = __obf_3cd39772de5fd3ee.Save(__obf_52e794309333fcfb)
	if __obf_7f8f5f5c173fa58e != nil {
		__obf_52e794309333fcfb.
			Close()
		return __obf_7f8f5f5c173fa58e
	}
	return __obf_52e794309333fcfb.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_3cd39772de5fd3ee *MapCache) Load(__obf_129832cddf20a553 io.Reader) error {
	__obf_f4175b2c269ef6e4 := gob.NewDecoder(__obf_129832cddf20a553)
	__obf_0c0fad92722f4264 := map[string]Item{}
	__obf_7f8f5f5c173fa58e := __obf_f4175b2c269ef6e4.Decode(&__obf_0c0fad92722f4264)
	if __obf_7f8f5f5c173fa58e == nil {
		__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
			Lock()
		defer __obf_3cd39772de5fd3ee.__obf_53983106240f1755.Unlock()
		for __obf_2a0493c028b94053, __obf_b1ecf81133de0ff6 := range __obf_0c0fad92722f4264 {
			__obf_cd9346a423f5ce9d, __obf_4228f4fcd56ac4f8 := __obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053]
			if !__obf_4228f4fcd56ac4f8 || __obf_cd9346a423f5ce9d.Expired() {
				__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
			}
		}
	}
	return __obf_7f8f5f5c173fa58e
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_3cd39772de5fd3ee *MapCache) LoadFile(__obf_6d6606cc565f833e string) error {
	__obf_52e794309333fcfb, __obf_7f8f5f5c173fa58e := os.Open(__obf_6d6606cc565f833e)
	if __obf_7f8f5f5c173fa58e != nil {
		return __obf_7f8f5f5c173fa58e
	}
	__obf_7f8f5f5c173fa58e = __obf_3cd39772de5fd3ee.Load(__obf_52e794309333fcfb)
	if __obf_7f8f5f5c173fa58e != nil {
		__obf_52e794309333fcfb.
			Close()
		return __obf_7f8f5f5c173fa58e
	}
	return __obf_52e794309333fcfb.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_3cd39772de5fd3ee *MapCache) Items() map[string]Item {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		RLock()
	defer __obf_3cd39772de5fd3ee.__obf_53983106240f1755.RUnlock()
	__obf_1674045fb710ff20 := make(map[string]Item, len(__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264))
	__obf_8150b358628c5292 := time.Now().UnixNano()
	for __obf_2a0493c028b94053, __obf_b1ecf81133de0ff6 := range __obf_3cd39772de5fd3ee.
		// "Inlining" of Expired
		__obf_0c0fad92722f4264 {

		if __obf_b1ecf81133de0ff6.Expiration > 0 {
			if __obf_8150b358628c5292 > __obf_b1ecf81133de0ff6.Expiration {
				continue
			}
		}
		__obf_1674045fb710ff20[__obf_2a0493c028b94053] = __obf_b1ecf81133de0ff6
	}
	return __obf_1674045fb710ff20
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_3cd39772de5fd3ee *MapCache) ItemCount() int {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		RLock()
	__obf_a091a23e60b9ef54 := len(__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264)
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		RUnlock()
	return __obf_a091a23e60b9ef54
}

// Delete all items from the cache.
func (__obf_3cd39772de5fd3ee *MapCache) Flush() {
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Lock()
	__obf_3cd39772de5fd3ee.__obf_0c0fad92722f4264 = map[string]Item{}
	__obf_3cd39772de5fd3ee.__obf_53983106240f1755.
		Unlock()
}

type __obf_c39f2111b362a882 struct {
	Interval               time.Duration
	__obf_bd9f25b055551994 chan bool
}

func (__obf_9711887cb685839e *__obf_c39f2111b362a882) Run(__obf_3cd39772de5fd3ee *MapCache) {
	__obf_d5f2d701df66c0a9 := time.NewTicker(__obf_9711887cb685839e.Interval)
	for {
		select {
		case <-__obf_d5f2d701df66c0a9.C:
			__obf_3cd39772de5fd3ee.
				DeleteExpired()
		case <-__obf_9711887cb685839e.__obf_bd9f25b055551994:
			__obf_d5f2d701df66c0a9.
				Stop()
			return
		}
	}
}

func __obf_7f2548c3fcfdfbd0(__obf_3cd39772de5fd3ee *MapCache) {
	__obf_3cd39772de5fd3ee.__obf_c39f2111b362a882.__obf_bd9f25b055551994 <- true
}

func __obf_01da4d2c441d6b6c(__obf_3cd39772de5fd3ee *MapCache, __obf_1c25c70f26dc0270 time.Duration) {
	__obf_9711887cb685839e := &__obf_c39f2111b362a882{
		Interval: __obf_1c25c70f26dc0270, __obf_bd9f25b055551994: make(chan bool),
	}
	__obf_3cd39772de5fd3ee.__obf_c39f2111b362a882 = __obf_9711887cb685839e
	go __obf_9711887cb685839e.Run(__obf_3cd39772de5fd3ee)
}

func __obf_dc3e062b073bd1d8(__obf_7bb25afd62a456b2 time.Duration, __obf_1674045fb710ff20 map[string]Item) *MapCache {
	if __obf_7bb25afd62a456b2 == 0 {
		__obf_7bb25afd62a456b2 = -1
	}
	__obf_3cd39772de5fd3ee := &MapCache{__obf_0a3ee88650ce44a7: __obf_7bb25afd62a456b2, __obf_0c0fad92722f4264: __obf_1674045fb710ff20}
	return __obf_3cd39772de5fd3ee
}

func __obf_da51d3d4467ffd14(__obf_7bb25afd62a456b2 time.Duration, __obf_1c25c70f26dc0270 time.Duration, __obf_1674045fb710ff20 map[string]Item) *MapCache {
	__obf_3cd39772de5fd3ee := __obf_dc3e062b073bd1d8(__obf_7bb25afd62a456b2,
		// This trick ensures that the janitor goroutine (which--granted it
		// was enabled--is running DeleteExpired on c forever) does not keep
		// the returned C object from being garbage collected. When it is
		// garbage collected, the finalizer stops the janitor goroutine, after
		// which c can be collected.
		// C := &MapCache{c}
		__obf_1674045fb710ff20)

	if __obf_1c25c70f26dc0270 > 0 {
		__obf_01da4d2c441d6b6c(__obf_3cd39772de5fd3ee, __obf_1c25c70f26dc0270)
		runtime.SetFinalizer(__obf_3cd39772de5fd3ee, __obf_7f2548c3fcfdfbd0)
	}
	return __obf_3cd39772de5fd3ee
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_0a3ee88650ce44a7, __obf_1d2bd7605226f5ba time.Duration) *MapCache {
	__obf_0c0fad92722f4264 := make(map[string]Item)
	return __obf_da51d3d4467ffd14(__obf_0a3ee88650ce44a7, __obf_1d2bd7605226f5ba, __obf_0c0fad92722f4264)
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
func NewFrom(__obf_0a3ee88650ce44a7, __obf_1d2bd7605226f5ba time.Duration, __obf_0c0fad92722f4264 map[string]Item) *MapCache {
	return __obf_da51d3d4467ffd14(__obf_0a3ee88650ce44a7, __obf_1d2bd7605226f5ba, __obf_0c0fad92722f4264)
}
