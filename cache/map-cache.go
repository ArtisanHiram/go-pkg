package __obf_79e7502d8563d901

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
func (__obf_a1788cdf3893cb5b Item) Expired() bool {
	if __obf_a1788cdf3893cb5b.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_a1788cdf3893cb5b.Expiration
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
	__obf_80ce7c8dcfb2ffa5 time.Duration
	__obf_7e212b5e6b308425 map[string]Item
	__obf_e1ec702455188bd2 sync.RWMutex
	__obf_fea7899e07bde3ee func(string, any)
	__obf_e9a8472a8856e943 *__obf_e9a8472a8856e943
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
func (__obf_db09942aea00505a *MapCache) Set(__obf_08ceb9044b83b8fd string, __obf_cda7749a15665783 any, __obf_fdac110b32b5e189 time.Duration) {
	// "Inlining" of set
	var __obf_8429338405294247 int64
	if __obf_fdac110b32b5e189 == DefaultExpiration {
		__obf_fdac110b32b5e189 = __obf_db09942aea00505a.__obf_80ce7c8dcfb2ffa5
	}
	if __obf_fdac110b32b5e189 > 0 {
		__obf_8429338405294247 = time.Now().Add(__obf_fdac110b32b5e189).UnixNano()
	}
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = Item{
		Object:     __obf_cda7749a15665783,
		Expiration: __obf_8429338405294247,
	}
	// TODO: Calls to mu.Unlock are currently not deferred because defer
	// adds ~200 ns (as of go1.)
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
}

func (__obf_db09942aea00505a *MapCache) __obf_65e9c3166f78e4b2(__obf_08ceb9044b83b8fd string, __obf_cda7749a15665783 any, __obf_fdac110b32b5e189 time.Duration) {
	var __obf_8429338405294247 int64
	if __obf_fdac110b32b5e189 == DefaultExpiration {
		__obf_fdac110b32b5e189 = __obf_db09942aea00505a.__obf_80ce7c8dcfb2ffa5
	}
	if __obf_fdac110b32b5e189 > 0 {
		__obf_8429338405294247 = time.Now().Add(__obf_fdac110b32b5e189).UnixNano()
	}
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = Item{
		Object:     __obf_cda7749a15665783,
		Expiration: __obf_8429338405294247,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_db09942aea00505a *MapCache) SetDefault(__obf_08ceb9044b83b8fd string, __obf_cda7749a15665783 any) {
	__obf_db09942aea00505a.Set(__obf_08ceb9044b83b8fd, __obf_cda7749a15665783, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_db09942aea00505a *MapCache) Add(__obf_08ceb9044b83b8fd string, __obf_cda7749a15665783 any, __obf_fdac110b32b5e189 time.Duration) error {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	_, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_c7d2682bcb03d15c(__obf_08ceb9044b83b8fd)
	if __obf_2dd2ca877d13cbbd {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return fmt.Errorf("Item %s already exists", __obf_08ceb9044b83b8fd)
	}
	__obf_db09942aea00505a.__obf_65e9c3166f78e4b2(__obf_08ceb9044b83b8fd, __obf_cda7749a15665783, __obf_fdac110b32b5e189)
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return nil
}

// Set a new value for the cache key only if it already exists, and the existing
// item hasn't expired. Returns an error otherwise.
func (__obf_db09942aea00505a *MapCache) Replace(__obf_08ceb9044b83b8fd string, __obf_cda7749a15665783 any, __obf_fdac110b32b5e189 time.Duration) error {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	_, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_c7d2682bcb03d15c(__obf_08ceb9044b83b8fd)
	if !__obf_2dd2ca877d13cbbd {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_08ceb9044b83b8fd)
	}
	__obf_db09942aea00505a.__obf_65e9c3166f78e4b2(__obf_08ceb9044b83b8fd, __obf_cda7749a15665783, __obf_fdac110b32b5e189)
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return nil
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
// whether the key was found.
func (__obf_db09942aea00505a *MapCache) Get(__obf_08ceb9044b83b8fd string) (any, bool) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.RLock()
	// "Inlining" of get and Expired
	__obf_a1788cdf3893cb5b, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.RUnlock()
		return nil, false
	}
	if __obf_a1788cdf3893cb5b.Expiration > 0 {
		if time.Now().UnixNano() > __obf_a1788cdf3893cb5b.Expiration {
			__obf_db09942aea00505a.__obf_e1ec702455188bd2.RUnlock()
			return nil, false
		}
	}
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.RUnlock()
	return __obf_a1788cdf3893cb5b.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_db09942aea00505a *MapCache) GetWithExpiration(__obf_08ceb9044b83b8fd string) (any, time.Time, bool) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.RLock()
	// "Inlining" of get and Expired
	__obf_a1788cdf3893cb5b, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_a1788cdf3893cb5b.Expiration > 0 {
		if time.Now().UnixNano() > __obf_a1788cdf3893cb5b.Expiration {
			__obf_db09942aea00505a.__obf_e1ec702455188bd2.RUnlock()
			return nil, time.Time{}, false
		}

		// Return the item and the expiration time
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.RUnlock()
		return __obf_a1788cdf3893cb5b.Object, time.Unix(0, __obf_a1788cdf3893cb5b.Expiration), true
	}

	// If expiration <= 0 (i.e. no expiration time set) then return the item
	// and a zeroed time.Time
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.RUnlock()
	return __obf_a1788cdf3893cb5b.Object, time.Time{}, true
}

func (__obf_db09942aea00505a *MapCache) __obf_c7d2682bcb03d15c(__obf_08ceb9044b83b8fd string) (any, bool) {
	__obf_a1788cdf3893cb5b, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_a1788cdf3893cb5b.Expiration > 0 {
		if time.Now().UnixNano() > __obf_a1788cdf3893cb5b.Expiration {
			return nil, false
		}
	}
	return __obf_a1788cdf3893cb5b.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_db09942aea00505a *MapCache) Increment(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int64) error {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	switch __obf_bfb4c8a1c77523af.Object.(type) {
	case int:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(int) + int(__obf_64ce7612e297509f)
	case int8:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(int8) + int8(__obf_64ce7612e297509f)
	case int16:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(int16) + int16(__obf_64ce7612e297509f)
	case int32:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(int32) + int32(__obf_64ce7612e297509f)
	case int64:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(int64) + __obf_64ce7612e297509f
	case uint:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uint) + uint(__obf_64ce7612e297509f)
	case uintptr:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uintptr) + uintptr(__obf_64ce7612e297509f)
	case uint8:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uint8) + uint8(__obf_64ce7612e297509f)
	case uint16:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uint16) + uint16(__obf_64ce7612e297509f)
	case uint32:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uint32) + uint32(__obf_64ce7612e297509f)
	case uint64:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uint64) + uint64(__obf_64ce7612e297509f)
	case float32:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(float32) + float32(__obf_64ce7612e297509f)
	case float64:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(float64) + float64(__obf_64ce7612e297509f)
	default:
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_08ceb9044b83b8fd)
	}
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return nil
}

// Increment an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to increment it by n. Pass a negative number to decrement the
// value. To retrieve the incremented value, use one of the specialized methods,
// e.g. IncrementFloat64.
func (__obf_db09942aea00505a *MapCache) IncrementFloat(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f float64) error {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	switch __obf_bfb4c8a1c77523af.Object.(type) {
	case float32:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(float32) + float32(__obf_64ce7612e297509f)
	case float64:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(float64) + __obf_64ce7612e297509f
	default:
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_08ceb9044b83b8fd)
	}
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return nil
}

// Increment an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementInt(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int) (int, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(int)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementInt8(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int8) (int8, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(int8)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementInt16(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int16) (int16, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(int16)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementInt32(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int32) (int32, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(int32)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementInt64(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int64) (int64, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(int64)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementUint(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uint) (uint, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uint)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementUintptr(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uintptr) (uintptr, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uintptr)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type uint8 by n. Returns an error if the item's value
// is not an uint8, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementUint8(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uint8) (uint8, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uint8)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementUint16(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uint16) (uint16, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uint16)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementUint32(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uint32) (uint32, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uint32)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementUint64(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uint64) (uint64, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uint64)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementFloat32(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f float32) (float32, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(float32)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Increment an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_db09942aea00505a *MapCache) IncrementFloat64(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f float64) (float64, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(float64)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 + __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to decrement it by n. To retrieve the decremented value, use one
// of the specialized methods, e.g. DecrementInt64.
func (__obf_db09942aea00505a *MapCache) Decrement(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int64) error {
	// TODO: Implement Increment and Decrement more cleanly.
	// (Cannot do Increment(k, n*-1) for uints.)
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_bfb4c8a1c77523af.Object.(type) {
	case int:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(int) - int(__obf_64ce7612e297509f)
	case int8:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(int8) - int8(__obf_64ce7612e297509f)
	case int16:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(int16) - int16(__obf_64ce7612e297509f)
	case int32:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(int32) - int32(__obf_64ce7612e297509f)
	case int64:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(int64) - __obf_64ce7612e297509f
	case uint:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uint) - uint(__obf_64ce7612e297509f)
	case uintptr:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uintptr) - uintptr(__obf_64ce7612e297509f)
	case uint8:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uint8) - uint8(__obf_64ce7612e297509f)
	case uint16:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uint16) - uint16(__obf_64ce7612e297509f)
	case uint32:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uint32) - uint32(__obf_64ce7612e297509f)
	case uint64:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(uint64) - uint64(__obf_64ce7612e297509f)
	case float32:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(float32) - float32(__obf_64ce7612e297509f)
	case float64:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(float64) - float64(__obf_64ce7612e297509f)
	default:
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_08ceb9044b83b8fd)
	}
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return nil
}

// Decrement an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to decrement it by n. Pass a negative number to decrement the
// value. To retrieve the decremented value, use one of the specialized methods,
// e.g. DecrementFloat64.
func (__obf_db09942aea00505a *MapCache) DecrementFloat(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f float64) error {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	switch __obf_bfb4c8a1c77523af.Object.(type) {
	case float32:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(float32) - float32(__obf_64ce7612e297509f)
	case float64:
		__obf_bfb4c8a1c77523af.Object = __obf_bfb4c8a1c77523af.Object.(float64) - __obf_64ce7612e297509f
	default:
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_08ceb9044b83b8fd)
	}
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return nil
}

// Decrement an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementInt(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int) (int, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(int)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementInt8(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int8) (int8, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(int8)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementInt16(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int16) (int16, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(int16)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementInt32(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int32) (int32, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(int32)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementInt64(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f int64) (int64, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(int64)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementUint(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uint) (uint, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uint)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementUintptr(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uintptr) (uintptr, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uintptr)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type uint8 by n. Returns an error if the item's value is
// not an uint8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementUint8(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uint8) (uint8, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uint8)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementUint16(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uint16) (uint16, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uint16)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementUint32(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uint32) (uint32, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uint32)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementUint64(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f uint64) (uint64, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(uint64)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementFloat32(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f float32) (float32, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(float32)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Decrement an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_db09942aea00505a *MapCache) DecrementFloat64(__obf_08ceb9044b83b8fd string, __obf_64ce7612e297509f float64) (float64, error) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
	if !__obf_2dd2ca877d13cbbd || __obf_bfb4c8a1c77523af.Expired() {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_08ceb9044b83b8fd)
	}
	__obf_1bfa7b98d4cc8853, __obf_7500628aeb1f47ab := __obf_bfb4c8a1c77523af.Object.(float64)
	if !__obf_7500628aeb1f47ab {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_08ceb9044b83b8fd)
	}
	__obf_94d71dacb8aaa08e := __obf_1bfa7b98d4cc8853 - __obf_64ce7612e297509f
	__obf_bfb4c8a1c77523af.Object = __obf_94d71dacb8aaa08e
	__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	return __obf_94d71dacb8aaa08e, nil
}

// Delete an item from the cache. Does nothing if the key is not in the cache.
func (__obf_db09942aea00505a *MapCache) Delete(__obf_08ceb9044b83b8fd string) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_bfb4c8a1c77523af, __obf_7a044bcf850f74bf := __obf_db09942aea00505a.delete(__obf_08ceb9044b83b8fd)
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	if __obf_7a044bcf850f74bf {
		__obf_db09942aea00505a.__obf_fea7899e07bde3ee(__obf_08ceb9044b83b8fd, __obf_bfb4c8a1c77523af)
	}
}

func (__obf_db09942aea00505a *MapCache) delete(__obf_08ceb9044b83b8fd string) (any, bool) {
	if __obf_db09942aea00505a.__obf_fea7899e07bde3ee != nil {
		if __obf_bfb4c8a1c77523af, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]; __obf_2dd2ca877d13cbbd {
			delete(__obf_db09942aea00505a.__obf_7e212b5e6b308425, __obf_08ceb9044b83b8fd)
			return __obf_bfb4c8a1c77523af.Object, true
		}
	}
	delete(__obf_db09942aea00505a.__obf_7e212b5e6b308425, __obf_08ceb9044b83b8fd)
	return nil, false
}

type __obf_905f749ba3689104 struct {
	__obf_50994613b7653a88 string
	__obf_c0aa5e8a46724fa3 any
}

// Delete all expired items from the cache.
func (__obf_db09942aea00505a *MapCache) DeleteExpired() {
	var __obf_21a2574b2c415dc7 []__obf_905f749ba3689104
	__obf_0a75a8dc85acfc9b := time.Now().UnixNano()
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	for __obf_08ceb9044b83b8fd, __obf_bfb4c8a1c77523af := range __obf_db09942aea00505a.__obf_7e212b5e6b308425 {
		// "Inlining" of expired
		if __obf_bfb4c8a1c77523af.Expiration > 0 && __obf_0a75a8dc85acfc9b > __obf_bfb4c8a1c77523af.Expiration {
			__obf_367ab49ba8e3ef20, __obf_7a044bcf850f74bf := __obf_db09942aea00505a.delete(__obf_08ceb9044b83b8fd)
			if __obf_7a044bcf850f74bf {
				__obf_21a2574b2c415dc7 = append(__obf_21a2574b2c415dc7, __obf_905f749ba3689104{__obf_08ceb9044b83b8fd, __obf_367ab49ba8e3ef20})
			}
		}
	}
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
	for _, __obf_bfb4c8a1c77523af := range __obf_21a2574b2c415dc7 {
		__obf_db09942aea00505a.__obf_fea7899e07bde3ee(__obf_bfb4c8a1c77523af.__obf_50994613b7653a88, __obf_bfb4c8a1c77523af.__obf_c0aa5e8a46724fa3)
	}
}

// Sets an (optional) function that is called with the key and value when an
// item is evicted from the cache. (Including when it is deleted manually, but
// not when it is overwritten.) Set to nil to disable.
func (__obf_db09942aea00505a *MapCache) OnEvicted(__obf_091eb7703f5d36d5 func(string, any)) {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_db09942aea00505a.__obf_fea7899e07bde3ee = __obf_091eb7703f5d36d5
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
}

// Write the cache's items (using Gob) to an io.Writer.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_db09942aea00505a *MapCache) Save(__obf_d18f4c2b1983c001 io.Writer) (__obf_8d46af2525fab46a error) {
	__obf_c99f1d2cb55f0274 := gob.NewEncoder(__obf_d18f4c2b1983c001)
	defer func() {
		if __obf_cda7749a15665783 := recover(); __obf_cda7749a15665783 != nil {
			__obf_8d46af2525fab46a = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.RLock()
	defer __obf_db09942aea00505a.__obf_e1ec702455188bd2.RUnlock()
	for _, __obf_bfb4c8a1c77523af := range __obf_db09942aea00505a.__obf_7e212b5e6b308425 {
		gob.Register(__obf_bfb4c8a1c77523af.Object)
	}
	__obf_8d46af2525fab46a = __obf_c99f1d2cb55f0274.Encode(&__obf_db09942aea00505a.__obf_7e212b5e6b308425)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_db09942aea00505a *MapCache) SaveFile(__obf_0795082e6a334958 string) error {
	__obf_f8ae3ec80bf41dd7, __obf_8d46af2525fab46a := os.Create(__obf_0795082e6a334958)
	if __obf_8d46af2525fab46a != nil {
		return __obf_8d46af2525fab46a
	}
	__obf_8d46af2525fab46a = __obf_db09942aea00505a.Save(__obf_f8ae3ec80bf41dd7)
	if __obf_8d46af2525fab46a != nil {
		__obf_f8ae3ec80bf41dd7.Close()
		return __obf_8d46af2525fab46a
	}
	return __obf_f8ae3ec80bf41dd7.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_db09942aea00505a *MapCache) Load(__obf_3038774b5f5f3925 io.Reader) error {
	__obf_2803cb7ab7b26a93 := gob.NewDecoder(__obf_3038774b5f5f3925)
	__obf_7e212b5e6b308425 := map[string]Item{}
	__obf_8d46af2525fab46a := __obf_2803cb7ab7b26a93.Decode(&__obf_7e212b5e6b308425)
	if __obf_8d46af2525fab46a == nil {
		__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
		defer __obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
		for __obf_08ceb9044b83b8fd, __obf_bfb4c8a1c77523af := range __obf_7e212b5e6b308425 {
			__obf_367ab49ba8e3ef20, __obf_2dd2ca877d13cbbd := __obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd]
			if !__obf_2dd2ca877d13cbbd || __obf_367ab49ba8e3ef20.Expired() {
				__obf_db09942aea00505a.__obf_7e212b5e6b308425[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
			}
		}
	}
	return __obf_8d46af2525fab46a
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_db09942aea00505a *MapCache) LoadFile(__obf_0795082e6a334958 string) error {
	__obf_f8ae3ec80bf41dd7, __obf_8d46af2525fab46a := os.Open(__obf_0795082e6a334958)
	if __obf_8d46af2525fab46a != nil {
		return __obf_8d46af2525fab46a
	}
	__obf_8d46af2525fab46a = __obf_db09942aea00505a.Load(__obf_f8ae3ec80bf41dd7)
	if __obf_8d46af2525fab46a != nil {
		__obf_f8ae3ec80bf41dd7.Close()
		return __obf_8d46af2525fab46a
	}
	return __obf_f8ae3ec80bf41dd7.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_db09942aea00505a *MapCache) Items() map[string]Item {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.RLock()
	defer __obf_db09942aea00505a.__obf_e1ec702455188bd2.RUnlock()
	__obf_bd80a351910d9377 := make(map[string]Item, len(__obf_db09942aea00505a.__obf_7e212b5e6b308425))
	__obf_0a75a8dc85acfc9b := time.Now().UnixNano()
	for __obf_08ceb9044b83b8fd, __obf_bfb4c8a1c77523af := range __obf_db09942aea00505a.__obf_7e212b5e6b308425 {
		// "Inlining" of Expired
		if __obf_bfb4c8a1c77523af.Expiration > 0 {
			if __obf_0a75a8dc85acfc9b > __obf_bfb4c8a1c77523af.Expiration {
				continue
			}
		}
		__obf_bd80a351910d9377[__obf_08ceb9044b83b8fd] = __obf_bfb4c8a1c77523af
	}
	return __obf_bd80a351910d9377
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_db09942aea00505a *MapCache) ItemCount() int {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.RLock()
	__obf_64ce7612e297509f := len(__obf_db09942aea00505a.__obf_7e212b5e6b308425)
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.RUnlock()
	return __obf_64ce7612e297509f
}

// Delete all items from the cache.
func (__obf_db09942aea00505a *MapCache) Flush() {
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Lock()
	__obf_db09942aea00505a.__obf_7e212b5e6b308425 = map[string]Item{}
	__obf_db09942aea00505a.__obf_e1ec702455188bd2.Unlock()
}

type __obf_e9a8472a8856e943 struct {
	Interval               time.Duration
	__obf_9a074c07a1f51392 chan bool
}

func (__obf_15e52c376e66e938 *__obf_e9a8472a8856e943) Run(__obf_db09942aea00505a *MapCache) {
	__obf_4f36fdfefba061d7 := time.NewTicker(__obf_15e52c376e66e938.Interval)
	for {
		select {
		case <-__obf_4f36fdfefba061d7.C:
			__obf_db09942aea00505a.DeleteExpired()
		case <-__obf_15e52c376e66e938.__obf_9a074c07a1f51392:
			__obf_4f36fdfefba061d7.Stop()
			return
		}
	}
}

func __obf_fea217aee1310560(__obf_db09942aea00505a *MapCache) {
	__obf_db09942aea00505a.__obf_e9a8472a8856e943.__obf_9a074c07a1f51392 <- true
}

func __obf_208b4e618b077b53(__obf_db09942aea00505a *MapCache, __obf_ca01c992d32e8279 time.Duration) {
	__obf_15e52c376e66e938 := &__obf_e9a8472a8856e943{
		Interval:               __obf_ca01c992d32e8279,
		__obf_9a074c07a1f51392: make(chan bool),
	}
	__obf_db09942aea00505a.__obf_e9a8472a8856e943 = __obf_15e52c376e66e938
	go __obf_15e52c376e66e938.Run(__obf_db09942aea00505a)
}

func __obf_25951e185666aa7b(__obf_ea9977e7cf8fd8cb time.Duration, __obf_bd80a351910d9377 map[string]Item) *MapCache {
	if __obf_ea9977e7cf8fd8cb == 0 {
		__obf_ea9977e7cf8fd8cb = -1
	}
	__obf_db09942aea00505a := &MapCache{
		__obf_80ce7c8dcfb2ffa5: __obf_ea9977e7cf8fd8cb,
		__obf_7e212b5e6b308425: __obf_bd80a351910d9377,
	}
	return __obf_db09942aea00505a
}

func __obf_ae943ac57bcb2041(__obf_ea9977e7cf8fd8cb time.Duration, __obf_ca01c992d32e8279 time.Duration, __obf_bd80a351910d9377 map[string]Item) *MapCache {
	__obf_db09942aea00505a := __obf_25951e185666aa7b(__obf_ea9977e7cf8fd8cb, __obf_bd80a351910d9377)
	// This trick ensures that the janitor goroutine (which--granted it
	// was enabled--is running DeleteExpired on c forever) does not keep
	// the returned C object from being garbage collected. When it is
	// garbage collected, the finalizer stops the janitor goroutine, after
	// which c can be collected.
	// C := &MapCache{c}
	if __obf_ca01c992d32e8279 > 0 {
		__obf_208b4e618b077b53(__obf_db09942aea00505a, __obf_ca01c992d32e8279)
		runtime.SetFinalizer(__obf_db09942aea00505a, __obf_fea217aee1310560)
	}
	return __obf_db09942aea00505a
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_80ce7c8dcfb2ffa5, __obf_ccfb5a471f2eefd6 time.Duration) *MapCache {
	__obf_7e212b5e6b308425 := make(map[string]Item)
	return __obf_ae943ac57bcb2041(__obf_80ce7c8dcfb2ffa5, __obf_ccfb5a471f2eefd6, __obf_7e212b5e6b308425)
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
func NewFrom(__obf_80ce7c8dcfb2ffa5, __obf_ccfb5a471f2eefd6 time.Duration, __obf_7e212b5e6b308425 map[string]Item) *MapCache {
	return __obf_ae943ac57bcb2041(__obf_80ce7c8dcfb2ffa5, __obf_ccfb5a471f2eefd6, __obf_7e212b5e6b308425)
}
