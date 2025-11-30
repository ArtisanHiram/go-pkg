package __obf_72d962f6a40bc95f

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
func (__obf_cdc238f327ed4e5c Item) Expired() bool {
	if __obf_cdc238f327ed4e5c.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_cdc238f327ed4e5c.Expiration
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
	__obf_2a700263752d2e00 time. // *cache
		// If this is confusing, see the comment at the bottom of New()
		Duration
	__obf_2c5da83e1dd3805c map[string]Item
	__obf_09437e388e205e38 sync.RWMutex
	__obf_ef9bf920f1ed755d func(string, any)
	__obf_ed816f9c2df1d218 *__obf_ed816f9c2df1d218
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
func (__obf_89bd5bf0a2a3aefa *MapCache) Set(__obf_41bce9a58f43c231 string, __obf_4c4a641aafd95cda any, __obf_57a090cfb5a3e8b0 time.Duration) {
	// "Inlining" of set
	var __obf_d1618b4c9656c93c int64
	if __obf_57a090cfb5a3e8b0 == DefaultExpiration {
		__obf_57a090cfb5a3e8b0 = __obf_89bd5bf0a2a3aefa.__obf_2a700263752d2e00
	}
	if __obf_57a090cfb5a3e8b0 > 0 {
		__obf_d1618b4c9656c93c = time.Now().Add(__obf_57a090cfb5a3e8b0).UnixNano()
	}
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = Item{
		Object:     __obf_4c4a641aafd95cda,
		Expiration: __obf_d1618b4c9656c93c,
	}
	__obf_89bd5bf0a2a3aefa.
		// TODO: Calls to mu.Unlock are currently not deferred because defer
		// adds ~200 ns (as of go1.)
		__obf_09437e388e205e38.
		Unlock()
}

func (__obf_89bd5bf0a2a3aefa *MapCache) __obf_3d57127afdc41dbd(__obf_41bce9a58f43c231 string, __obf_4c4a641aafd95cda any, __obf_57a090cfb5a3e8b0 time.Duration) {
	var __obf_d1618b4c9656c93c int64
	if __obf_57a090cfb5a3e8b0 == DefaultExpiration {
		__obf_57a090cfb5a3e8b0 = __obf_89bd5bf0a2a3aefa.__obf_2a700263752d2e00
	}
	if __obf_57a090cfb5a3e8b0 > 0 {
		__obf_d1618b4c9656c93c = time.Now().Add(__obf_57a090cfb5a3e8b0).UnixNano()
	}
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = Item{
		Object:     __obf_4c4a641aafd95cda,
		Expiration: __obf_d1618b4c9656c93c,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_89bd5bf0a2a3aefa *MapCache) SetDefault(__obf_41bce9a58f43c231 string, __obf_4c4a641aafd95cda any) {
	__obf_89bd5bf0a2a3aefa.
		Set(__obf_41bce9a58f43c231, __obf_4c4a641aafd95cda, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_89bd5bf0a2a3aefa *MapCache) Add(__obf_41bce9a58f43c231 string, __obf_4c4a641aafd95cda any, __obf_57a090cfb5a3e8b0 time.Duration) error {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	_, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_f0343616f43c52e5(__obf_41bce9a58f43c231)
	if __obf_29bd8183b37d202a {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return fmt.Errorf("Item %s already exists", __obf_41bce9a58f43c231)
	}
	__obf_89bd5bf0a2a3aefa.__obf_3d57127afdc41dbd(__obf_41bce9a58f43c231, __obf_4c4a641aafd95cda,

		// Set a new value for the cache key only if it already exists, and the existing
		// item hasn't expired. Returns an error otherwise.
		__obf_57a090cfb5a3e8b0)
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Unlock()
	return nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) Replace(__obf_41bce9a58f43c231 string, __obf_4c4a641aafd95cda any, __obf_57a090cfb5a3e8b0 time.Duration) error {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	_, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_f0343616f43c52e5(__obf_41bce9a58f43c231)
	if !__obf_29bd8183b37d202a {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_41bce9a58f43c231)
	}
	__obf_89bd5bf0a2a3aefa.__obf_3d57127afdc41dbd(__obf_41bce9a58f43c231, __obf_4c4a641aafd95cda,

		// Get an item from the cache. Returns the item or nil, and a bool indicating
		// whether the key was found.
		__obf_57a090cfb5a3e8b0)
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Unlock()
	return nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) Get(__obf_41bce9a58f43c231 string) (any, bool) {
	__obf_89bd5bf0a2a3aefa.

		// "Inlining" of get and Expired
		__obf_09437e388e205e38.
		RLock()
	__obf_cdc238f327ed4e5c, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			RUnlock()
		return nil, false
	}
	if __obf_cdc238f327ed4e5c.Expiration > 0 {
		if time.Now().UnixNano() > __obf_cdc238f327ed4e5c.Expiration {
			__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
				RUnlock()
			return nil, false
		}
	}
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		RUnlock()
	return __obf_cdc238f327ed4e5c.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_89bd5bf0a2a3aefa *MapCache) GetWithExpiration(__obf_41bce9a58f43c231 string) (any, time.Time, bool) {
	__obf_89bd5bf0a2a3aefa.

		// "Inlining" of get and Expired
		__obf_09437e388e205e38.
		RLock()
	__obf_cdc238f327ed4e5c, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_cdc238f327ed4e5c.Expiration > 0 {
		if time.Now().UnixNano() > __obf_cdc238f327ed4e5c.Expiration {
			__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
				RUnlock()
			return nil, time.Time{}, false
		}
		__obf_89bd5bf0a2a3aefa.

			// Return the item and the expiration time
			__obf_09437e388e205e38.
			RUnlock()
		return __obf_cdc238f327ed4e5c.Object, time.Unix(0, __obf_cdc238f327ed4e5c.Expiration), true
	}
	__obf_89bd5bf0a2a3aefa.

		// If expiration <= 0 (i.e. no expiration time set) then return the item
		// and a zeroed time.Time
		__obf_09437e388e205e38.
		RUnlock()
	return __obf_cdc238f327ed4e5c.Object, time.Time{}, true
}

func (__obf_89bd5bf0a2a3aefa *MapCache) __obf_f0343616f43c52e5(__obf_41bce9a58f43c231 string) (any, bool) {
	__obf_cdc238f327ed4e5c, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_cdc238f327ed4e5c.Expiration > 0 {
		if time.Now().UnixNano() > __obf_cdc238f327ed4e5c.Expiration {
			return nil, false
		}
	}
	return __obf_cdc238f327ed4e5c.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_89bd5bf0a2a3aefa *MapCache) Increment(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int64) error {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	switch __obf_b33efcc537be1290.Object.(type) {
	case int:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(int) + int(__obf_2bf733a50ac047cc)
	case int8:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(int8) + int8(__obf_2bf733a50ac047cc)
	case int16:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(int16) + int16(__obf_2bf733a50ac047cc)
	case int32:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(int32) + int32(__obf_2bf733a50ac047cc)
	case int64:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(int64) + __obf_2bf733a50ac047cc
	case uint:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uint) + uint(__obf_2bf733a50ac047cc)
	case uintptr:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uintptr) + uintptr(__obf_2bf733a50ac047cc)
	case uint8:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uint8) + uint8(__obf_2bf733a50ac047cc)
	case uint16:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uint16) + uint16(__obf_2bf733a50ac047cc)
	case uint32:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uint32) + uint32(__obf_2bf733a50ac047cc)
	case uint64:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uint64) + uint64(__obf_2bf733a50ac047cc)
	case float32:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(float32) + float32(__obf_2bf733a50ac047cc)
	case float64:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(float64) + float64(__obf_2bf733a50ac047cc)
	default:
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_41bce9a58f43c231)
	}
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to increment it by n. Pass a negative number to decrement the
		// value. To retrieve the incremented value, use one of the specialized methods,
		// e.g. IncrementFloat64.
		__obf_09437e388e205e38.
		Unlock()
	return nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementFloat(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc float64) error {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	switch __obf_b33efcc537be1290.Object.(type) {
	case float32:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(float32) + float32(__obf_2bf733a50ac047cc)
	case float64:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(float64) + __obf_2bf733a50ac047cc
	default:
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_41bce9a58f43c231)
	}
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementInt(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int) (int, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(int)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementInt8(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int8) (int8, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(int8)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementInt16(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int16) (int16, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(int16)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementInt32(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int32) (int32, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(int32)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementInt64(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int64) (int64, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(int64)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementUint(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uint) (uint, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uint)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementUintptr(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uintptr) (uintptr, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uintptr)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type uint8 by n. Returns an error if the item's value
		// is not an uint8, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementUint8(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uint8) (uint8, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uint8)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementUint16(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uint16) (uint16, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uint16)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementUint32(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uint32) (uint32, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uint32)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementUint64(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uint64) (uint64, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uint64)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementFloat32(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc float32) (float32, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(float32)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Increment an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) IncrementFloat64(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc float64) (float64, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(float64)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 + __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
		// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
		// item's value is not an integer, if it was not found, or if it is not
		// possible to decrement it by n. To retrieve the decremented value, use one
		// of the specialized methods, e.g. DecrementInt64.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) Decrement(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int64) error {
	__obf_89bd5bf0a2a3aefa.
		// TODO: Implement Increment and Decrement more cleanly.
		// (Cannot do Increment(k, n*-1) for uints.)
		__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_b33efcc537be1290.Object.(type) {
	case int:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(int) - int(__obf_2bf733a50ac047cc)
	case int8:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(int8) - int8(__obf_2bf733a50ac047cc)
	case int16:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(int16) - int16(__obf_2bf733a50ac047cc)
	case int32:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(int32) - int32(__obf_2bf733a50ac047cc)
	case int64:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(int64) - __obf_2bf733a50ac047cc
	case uint:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uint) - uint(__obf_2bf733a50ac047cc)
	case uintptr:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uintptr) - uintptr(__obf_2bf733a50ac047cc)
	case uint8:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uint8) - uint8(__obf_2bf733a50ac047cc)
	case uint16:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uint16) - uint16(__obf_2bf733a50ac047cc)
	case uint32:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uint32) - uint32(__obf_2bf733a50ac047cc)
	case uint64:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(uint64) - uint64(__obf_2bf733a50ac047cc)
	case float32:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(float32) - float32(__obf_2bf733a50ac047cc)
	case float64:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(float64) - float64(__obf_2bf733a50ac047cc)
	default:
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_41bce9a58f43c231)
	}
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to decrement it by n. Pass a negative number to decrement the
		// value. To retrieve the decremented value, use one of the specialized methods,
		// e.g. DecrementFloat64.
		__obf_09437e388e205e38.
		Unlock()
	return nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementFloat(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc float64) error {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	switch __obf_b33efcc537be1290.Object.(type) {
	case float32:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(float32) - float32(__obf_2bf733a50ac047cc)
	case float64:
		__obf_b33efcc537be1290.
			Object = __obf_b33efcc537be1290.Object.(float64) - __obf_2bf733a50ac047cc
	default:
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_41bce9a58f43c231)
	}
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementInt(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int) (int, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(int)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementInt8(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int8) (int8, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(int8)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementInt16(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int16) (int16, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(int16)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementInt32(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int32) (int32, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(int32)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementInt64(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc int64) (int64, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(int64)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementUint(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uint) (uint, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uint)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementUintptr(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uintptr) (uintptr, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uintptr)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type uint8 by n. Returns an error if the item's value is
		// not an uint8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementUint8(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uint8) (uint8, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uint8)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementUint16(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uint16) (uint16, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uint16)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementUint32(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uint32) (uint32, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uint32)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementUint64(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc uint64) (uint64, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(uint64)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementFloat32(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc float32) (float32, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(float32)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Decrement an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) DecrementFloat64(__obf_41bce9a58f43c231 string, __obf_2bf733a50ac047cc float64) (float64, error) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
	if !__obf_29bd8183b37d202a || __obf_b33efcc537be1290.Expired() {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_41bce9a58f43c231)
	}
	__obf_7387261227c81695, __obf_97faa3bbfe732802 := __obf_b33efcc537be1290.Object.(float64)
	if !__obf_97faa3bbfe732802 {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_41bce9a58f43c231)
	}
	__obf_046fb6b33a57df1a := __obf_7387261227c81695 - __obf_2bf733a50ac047cc
	__obf_b33efcc537be1290.
		Object = __obf_046fb6b33a57df1a
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	__obf_89bd5bf0a2a3aefa.

		// Delete an item from the cache. Does nothing if the key is not in the cache.
		__obf_09437e388e205e38.
		Unlock()
	return __obf_046fb6b33a57df1a, nil
}

func (__obf_89bd5bf0a2a3aefa *MapCache) Delete(__obf_41bce9a58f43c231 string) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_b33efcc537be1290, __obf_7bc1b66a7503946d := __obf_89bd5bf0a2a3aefa.delete(__obf_41bce9a58f43c231)
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Unlock()
	if __obf_7bc1b66a7503946d {
		__obf_89bd5bf0a2a3aefa.__obf_ef9bf920f1ed755d(__obf_41bce9a58f43c231, __obf_b33efcc537be1290)
	}
}

func (__obf_89bd5bf0a2a3aefa *MapCache) delete(__obf_41bce9a58f43c231 string) (any, bool) {
	if __obf_89bd5bf0a2a3aefa.__obf_ef9bf920f1ed755d != nil {
		if __obf_b33efcc537be1290, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]; __obf_29bd8183b37d202a {
			delete(__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c, __obf_41bce9a58f43c231)
			return __obf_b33efcc537be1290.Object, true
		}
	}
	delete(__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c, __obf_41bce9a58f43c231)
	return nil, false
}

type __obf_735a0727bc120993 struct {
	__obf_73cb8148cbf55699 string
	__obf_82bf2d5ce3534205 any
}

// Delete all expired items from the cache.
func (__obf_89bd5bf0a2a3aefa *MapCache) DeleteExpired() {
	var __obf_dcc8a508d47b6469 []__obf_735a0727bc120993
	__obf_77d958f49d42b249 := time.Now().UnixNano()
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	for __obf_41bce9a58f43c231, __obf_b33efcc537be1290 := range __obf_89bd5bf0a2a3aefa.
		// "Inlining" of expired
		__obf_2c5da83e1dd3805c {

		if __obf_b33efcc537be1290.Expiration > 0 && __obf_77d958f49d42b249 > __obf_b33efcc537be1290.Expiration {
			__obf_7b2bfb30845a2dbb, __obf_7bc1b66a7503946d := __obf_89bd5bf0a2a3aefa.delete(__obf_41bce9a58f43c231)
			if __obf_7bc1b66a7503946d {
				__obf_dcc8a508d47b6469 = append(__obf_dcc8a508d47b6469, __obf_735a0727bc120993{__obf_41bce9a58f43c231, __obf_7b2bfb30845a2dbb})
			}
		}
	}
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Unlock()
	for _, __obf_b33efcc537be1290 := range __obf_dcc8a508d47b6469 {
		__obf_89bd5bf0a2a3aefa.__obf_ef9bf920f1ed755d(__obf_b33efcc537be1290.__obf_73cb8148cbf55699,

			// Sets an (optional) function that is called with the key and value when an
			// item is evicted from the cache. (Including when it is deleted manually, but
			// not when it is overwritten.) Set to nil to disable.
			__obf_b33efcc537be1290.__obf_82bf2d5ce3534205)
	}
}

func (__obf_89bd5bf0a2a3aefa *MapCache) OnEvicted(__obf_1bfd87875172a06e func(string, any)) {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_89bd5bf0a2a3aefa.__obf_ef9bf920f1ed755d = __obf_1bfd87875172a06e
	__obf_89bd5bf0a2a3aefa.

		// Write the cache's items (using Gob) to an io.Writer.
		//
		// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
		// documentation for NewFrom().)
		__obf_09437e388e205e38.
		Unlock()
}

func (__obf_89bd5bf0a2a3aefa *MapCache) Save(__obf_62798b54046b78ec io.Writer) (__obf_2498adaec5f4c8f1 error) {
	__obf_a1c49ffc65961848 := gob.NewEncoder(__obf_62798b54046b78ec)
	defer func() {
		if __obf_4c4a641aafd95cda := recover(); __obf_4c4a641aafd95cda != nil {
			__obf_2498adaec5f4c8f1 = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		RLock()
	defer __obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.RUnlock()
	for _, __obf_b33efcc537be1290 := range __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c {
		gob.Register(__obf_b33efcc537be1290.Object)
	}
	__obf_2498adaec5f4c8f1 = __obf_a1c49ffc65961848.Encode(&__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_89bd5bf0a2a3aefa *MapCache) SaveFile(__obf_613374bd45123e6e string) error {
	__obf_4719b5d8aafb8e31, __obf_2498adaec5f4c8f1 := os.Create(__obf_613374bd45123e6e)
	if __obf_2498adaec5f4c8f1 != nil {
		return __obf_2498adaec5f4c8f1
	}
	__obf_2498adaec5f4c8f1 = __obf_89bd5bf0a2a3aefa.Save(__obf_4719b5d8aafb8e31)
	if __obf_2498adaec5f4c8f1 != nil {
		__obf_4719b5d8aafb8e31.
			Close()
		return __obf_2498adaec5f4c8f1
	}
	return __obf_4719b5d8aafb8e31.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_89bd5bf0a2a3aefa *MapCache) Load(__obf_d20240e4761bf874 io.Reader) error {
	__obf_b4c775e6df07ccbd := gob.NewDecoder(__obf_d20240e4761bf874)
	__obf_2c5da83e1dd3805c := map[string]Item{}
	__obf_2498adaec5f4c8f1 := __obf_b4c775e6df07ccbd.Decode(&__obf_2c5da83e1dd3805c)
	if __obf_2498adaec5f4c8f1 == nil {
		__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
			Lock()
		defer __obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.Unlock()
		for __obf_41bce9a58f43c231, __obf_b33efcc537be1290 := range __obf_2c5da83e1dd3805c {
			__obf_7b2bfb30845a2dbb, __obf_29bd8183b37d202a := __obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231]
			if !__obf_29bd8183b37d202a || __obf_7b2bfb30845a2dbb.Expired() {
				__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
			}
		}
	}
	return __obf_2498adaec5f4c8f1
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_89bd5bf0a2a3aefa *MapCache) LoadFile(__obf_613374bd45123e6e string) error {
	__obf_4719b5d8aafb8e31, __obf_2498adaec5f4c8f1 := os.Open(__obf_613374bd45123e6e)
	if __obf_2498adaec5f4c8f1 != nil {
		return __obf_2498adaec5f4c8f1
	}
	__obf_2498adaec5f4c8f1 = __obf_89bd5bf0a2a3aefa.Load(__obf_4719b5d8aafb8e31)
	if __obf_2498adaec5f4c8f1 != nil {
		__obf_4719b5d8aafb8e31.
			Close()
		return __obf_2498adaec5f4c8f1
	}
	return __obf_4719b5d8aafb8e31.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_89bd5bf0a2a3aefa *MapCache) Items() map[string]Item {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		RLock()
	defer __obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.RUnlock()
	__obf_a5a77ce672c162a7 := make(map[string]Item, len(__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c))
	__obf_77d958f49d42b249 := time.Now().UnixNano()
	for __obf_41bce9a58f43c231, __obf_b33efcc537be1290 := range __obf_89bd5bf0a2a3aefa.
		// "Inlining" of Expired
		__obf_2c5da83e1dd3805c {

		if __obf_b33efcc537be1290.Expiration > 0 {
			if __obf_77d958f49d42b249 > __obf_b33efcc537be1290.Expiration {
				continue
			}
		}
		__obf_a5a77ce672c162a7[__obf_41bce9a58f43c231] = __obf_b33efcc537be1290
	}
	return __obf_a5a77ce672c162a7
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_89bd5bf0a2a3aefa *MapCache) ItemCount() int {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		RLock()
	__obf_2bf733a50ac047cc := len(__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c)
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		RUnlock()
	return __obf_2bf733a50ac047cc
}

// Delete all items from the cache.
func (__obf_89bd5bf0a2a3aefa *MapCache) Flush() {
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Lock()
	__obf_89bd5bf0a2a3aefa.__obf_2c5da83e1dd3805c = map[string]Item{}
	__obf_89bd5bf0a2a3aefa.__obf_09437e388e205e38.
		Unlock()
}

type __obf_ed816f9c2df1d218 struct {
	Interval               time.Duration
	__obf_c9c4e600abcc9229 chan bool
}

func (__obf_c703205e5c1d8a6a *__obf_ed816f9c2df1d218) Run(__obf_89bd5bf0a2a3aefa *MapCache) {
	__obf_6dc21ba030136b32 := time.NewTicker(__obf_c703205e5c1d8a6a.Interval)
	for {
		select {
		case <-__obf_6dc21ba030136b32.C:
			__obf_89bd5bf0a2a3aefa.
				DeleteExpired()
		case <-__obf_c703205e5c1d8a6a.__obf_c9c4e600abcc9229:
			__obf_6dc21ba030136b32.
				Stop()
			return
		}
	}
}

func __obf_5bf862fe7cebc2f8(__obf_89bd5bf0a2a3aefa *MapCache) {
	__obf_89bd5bf0a2a3aefa.__obf_ed816f9c2df1d218.__obf_c9c4e600abcc9229 <- true
}

func __obf_2941414c774e9264(__obf_89bd5bf0a2a3aefa *MapCache, __obf_ac9020ec4f83a6e5 time.Duration) {
	__obf_c703205e5c1d8a6a := &__obf_ed816f9c2df1d218{
		Interval: __obf_ac9020ec4f83a6e5, __obf_c9c4e600abcc9229: make(chan bool),
	}
	__obf_89bd5bf0a2a3aefa.__obf_ed816f9c2df1d218 = __obf_c703205e5c1d8a6a
	go __obf_c703205e5c1d8a6a.Run(__obf_89bd5bf0a2a3aefa)
}

func __obf_cf785d2b8d303ce1(__obf_64c99bd794d570f1 time.Duration, __obf_a5a77ce672c162a7 map[string]Item) *MapCache {
	if __obf_64c99bd794d570f1 == 0 {
		__obf_64c99bd794d570f1 = -1
	}
	__obf_89bd5bf0a2a3aefa := &MapCache{__obf_2a700263752d2e00: __obf_64c99bd794d570f1, __obf_2c5da83e1dd3805c: __obf_a5a77ce672c162a7}
	return __obf_89bd5bf0a2a3aefa
}

func __obf_d483f71dd7cfc3bd(__obf_64c99bd794d570f1 time.Duration, __obf_ac9020ec4f83a6e5 time.Duration, __obf_a5a77ce672c162a7 map[string]Item) *MapCache {
	__obf_89bd5bf0a2a3aefa := __obf_cf785d2b8d303ce1(__obf_64c99bd794d570f1,
		// This trick ensures that the janitor goroutine (which--granted it
		// was enabled--is running DeleteExpired on c forever) does not keep
		// the returned C object from being garbage collected. When it is
		// garbage collected, the finalizer stops the janitor goroutine, after
		// which c can be collected.
		// C := &MapCache{c}
		__obf_a5a77ce672c162a7)

	if __obf_ac9020ec4f83a6e5 > 0 {
		__obf_2941414c774e9264(__obf_89bd5bf0a2a3aefa, __obf_ac9020ec4f83a6e5)
		runtime.SetFinalizer(__obf_89bd5bf0a2a3aefa, __obf_5bf862fe7cebc2f8)
	}
	return __obf_89bd5bf0a2a3aefa
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_2a700263752d2e00, __obf_df2f08adafb6e156 time.Duration) *MapCache {
	__obf_2c5da83e1dd3805c := make(map[string]Item)
	return __obf_d483f71dd7cfc3bd(__obf_2a700263752d2e00, __obf_df2f08adafb6e156, __obf_2c5da83e1dd3805c)
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
func NewFrom(__obf_2a700263752d2e00, __obf_df2f08adafb6e156 time.Duration, __obf_2c5da83e1dd3805c map[string]Item) *MapCache {
	return __obf_d483f71dd7cfc3bd(__obf_2a700263752d2e00, __obf_df2f08adafb6e156, __obf_2c5da83e1dd3805c)
}
