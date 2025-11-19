package __obf_62df4de078c8d208

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
func (__obf_76c3f3755f4b29a1 Item) Expired() bool {
	if __obf_76c3f3755f4b29a1.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_76c3f3755f4b29a1.Expiration
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
	__obf_bc351f72e9074c26 time.Duration
	__obf_1d5d1defae6f4d55 map[string]Item
	__obf_bc58716fda429cf7 sync.RWMutex
	__obf_f8b7842824568810 func(string, any)
	__obf_4336a7018604d0e4 *__obf_4336a7018604d0e4
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
func (__obf_1e8c1f17e8f9e9cd *MapCache) Set(__obf_c0676b3110071e57 string, __obf_901366c1b9787938 any, __obf_fc398c812f75a944 time.Duration) {
	// "Inlining" of set
	var __obf_68f2d1855583b558 int64
	if __obf_fc398c812f75a944 == DefaultExpiration {
		__obf_fc398c812f75a944 = __obf_1e8c1f17e8f9e9cd.__obf_bc351f72e9074c26
	}
	if __obf_fc398c812f75a944 > 0 {
		__obf_68f2d1855583b558 = time.Now().Add(__obf_fc398c812f75a944).UnixNano()
	}
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = Item{
		Object:     __obf_901366c1b9787938,
		Expiration: __obf_68f2d1855583b558,
	}
	// TODO: Calls to mu.Unlock are currently not deferred because defer
	// adds ~200 ns (as of go1.)
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
}

func (__obf_1e8c1f17e8f9e9cd *MapCache) __obf_3a3da05e88cba2a3(__obf_c0676b3110071e57 string, __obf_901366c1b9787938 any, __obf_fc398c812f75a944 time.Duration) {
	var __obf_68f2d1855583b558 int64
	if __obf_fc398c812f75a944 == DefaultExpiration {
		__obf_fc398c812f75a944 = __obf_1e8c1f17e8f9e9cd.__obf_bc351f72e9074c26
	}
	if __obf_fc398c812f75a944 > 0 {
		__obf_68f2d1855583b558 = time.Now().Add(__obf_fc398c812f75a944).UnixNano()
	}
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = Item{
		Object:     __obf_901366c1b9787938,
		Expiration: __obf_68f2d1855583b558,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_1e8c1f17e8f9e9cd *MapCache) SetDefault(__obf_c0676b3110071e57 string, __obf_901366c1b9787938 any) {
	__obf_1e8c1f17e8f9e9cd.Set(__obf_c0676b3110071e57, __obf_901366c1b9787938, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_1e8c1f17e8f9e9cd *MapCache) Add(__obf_c0676b3110071e57 string, __obf_901366c1b9787938 any, __obf_fc398c812f75a944 time.Duration) error {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	_, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_5ee661b99bc4737b(__obf_c0676b3110071e57)
	if __obf_889b1fa74a37b35f {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return fmt.Errorf("Item %s already exists", __obf_c0676b3110071e57)
	}
	__obf_1e8c1f17e8f9e9cd.__obf_3a3da05e88cba2a3(__obf_c0676b3110071e57, __obf_901366c1b9787938, __obf_fc398c812f75a944)
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return nil
}

// Set a new value for the cache key only if it already exists, and the existing
// item hasn't expired. Returns an error otherwise.
func (__obf_1e8c1f17e8f9e9cd *MapCache) Replace(__obf_c0676b3110071e57 string, __obf_901366c1b9787938 any, __obf_fc398c812f75a944 time.Duration) error {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	_, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_5ee661b99bc4737b(__obf_c0676b3110071e57)
	if !__obf_889b1fa74a37b35f {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_c0676b3110071e57)
	}
	__obf_1e8c1f17e8f9e9cd.__obf_3a3da05e88cba2a3(__obf_c0676b3110071e57, __obf_901366c1b9787938, __obf_fc398c812f75a944)
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return nil
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
// whether the key was found.
func (__obf_1e8c1f17e8f9e9cd *MapCache) Get(__obf_c0676b3110071e57 string) (any, bool) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RLock()
	// "Inlining" of get and Expired
	__obf_76c3f3755f4b29a1, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RUnlock()
		return nil, false
	}
	if __obf_76c3f3755f4b29a1.Expiration > 0 {
		if time.Now().UnixNano() > __obf_76c3f3755f4b29a1.Expiration {
			__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RUnlock()
			return nil, false
		}
	}
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RUnlock()
	return __obf_76c3f3755f4b29a1.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_1e8c1f17e8f9e9cd *MapCache) GetWithExpiration(__obf_c0676b3110071e57 string) (any, time.Time, bool) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RLock()
	// "Inlining" of get and Expired
	__obf_76c3f3755f4b29a1, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_76c3f3755f4b29a1.Expiration > 0 {
		if time.Now().UnixNano() > __obf_76c3f3755f4b29a1.Expiration {
			__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RUnlock()
			return nil, time.Time{}, false
		}

		// Return the item and the expiration time
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RUnlock()
		return __obf_76c3f3755f4b29a1.Object, time.Unix(0, __obf_76c3f3755f4b29a1.Expiration), true
	}

	// If expiration <= 0 (i.e. no expiration time set) then return the item
	// and a zeroed time.Time
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RUnlock()
	return __obf_76c3f3755f4b29a1.Object, time.Time{}, true
}

func (__obf_1e8c1f17e8f9e9cd *MapCache) __obf_5ee661b99bc4737b(__obf_c0676b3110071e57 string) (any, bool) {
	__obf_76c3f3755f4b29a1, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_76c3f3755f4b29a1.Expiration > 0 {
		if time.Now().UnixNano() > __obf_76c3f3755f4b29a1.Expiration {
			return nil, false
		}
	}
	return __obf_76c3f3755f4b29a1.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_1e8c1f17e8f9e9cd *MapCache) Increment(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int64) error {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	switch __obf_298afaabecd63bb9.Object.(type) {
	case int:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(int) + int(__obf_4029f0efebfc66c6)
	case int8:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(int8) + int8(__obf_4029f0efebfc66c6)
	case int16:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(int16) + int16(__obf_4029f0efebfc66c6)
	case int32:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(int32) + int32(__obf_4029f0efebfc66c6)
	case int64:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(int64) + __obf_4029f0efebfc66c6
	case uint:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uint) + uint(__obf_4029f0efebfc66c6)
	case uintptr:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uintptr) + uintptr(__obf_4029f0efebfc66c6)
	case uint8:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uint8) + uint8(__obf_4029f0efebfc66c6)
	case uint16:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uint16) + uint16(__obf_4029f0efebfc66c6)
	case uint32:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uint32) + uint32(__obf_4029f0efebfc66c6)
	case uint64:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uint64) + uint64(__obf_4029f0efebfc66c6)
	case float32:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(float32) + float32(__obf_4029f0efebfc66c6)
	case float64:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(float64) + float64(__obf_4029f0efebfc66c6)
	default:
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_c0676b3110071e57)
	}
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return nil
}

// Increment an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to increment it by n. Pass a negative number to decrement the
// value. To retrieve the incremented value, use one of the specialized methods,
// e.g. IncrementFloat64.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementFloat(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 float64) error {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	switch __obf_298afaabecd63bb9.Object.(type) {
	case float32:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(float32) + float32(__obf_4029f0efebfc66c6)
	case float64:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(float64) + __obf_4029f0efebfc66c6
	default:
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_c0676b3110071e57)
	}
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return nil
}

// Increment an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementInt(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int) (int, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(int)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementInt8(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int8) (int8, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(int8)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementInt16(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int16) (int16, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(int16)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementInt32(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int32) (int32, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(int32)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementInt64(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int64) (int64, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(int64)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementUint(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uint) (uint, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uint)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementUintptr(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uintptr) (uintptr, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uintptr)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type uint8 by n. Returns an error if the item's value
// is not an uint8, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementUint8(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uint8) (uint8, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uint8)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementUint16(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uint16) (uint16, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uint16)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementUint32(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uint32) (uint32, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uint32)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementUint64(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uint64) (uint64, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uint64)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementFloat32(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 float32) (float32, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(float32)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Increment an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) IncrementFloat64(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 float64) (float64, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(float64)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a + __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to decrement it by n. To retrieve the decremented value, use one
// of the specialized methods, e.g. DecrementInt64.
func (__obf_1e8c1f17e8f9e9cd *MapCache) Decrement(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int64) error {
	// TODO: Implement Increment and Decrement more cleanly.
	// (Cannot do Increment(k, n*-1) for uints.)
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_298afaabecd63bb9.Object.(type) {
	case int:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(int) - int(__obf_4029f0efebfc66c6)
	case int8:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(int8) - int8(__obf_4029f0efebfc66c6)
	case int16:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(int16) - int16(__obf_4029f0efebfc66c6)
	case int32:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(int32) - int32(__obf_4029f0efebfc66c6)
	case int64:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(int64) - __obf_4029f0efebfc66c6
	case uint:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uint) - uint(__obf_4029f0efebfc66c6)
	case uintptr:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uintptr) - uintptr(__obf_4029f0efebfc66c6)
	case uint8:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uint8) - uint8(__obf_4029f0efebfc66c6)
	case uint16:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uint16) - uint16(__obf_4029f0efebfc66c6)
	case uint32:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uint32) - uint32(__obf_4029f0efebfc66c6)
	case uint64:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(uint64) - uint64(__obf_4029f0efebfc66c6)
	case float32:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(float32) - float32(__obf_4029f0efebfc66c6)
	case float64:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(float64) - float64(__obf_4029f0efebfc66c6)
	default:
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_c0676b3110071e57)
	}
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return nil
}

// Decrement an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to decrement it by n. Pass a negative number to decrement the
// value. To retrieve the decremented value, use one of the specialized methods,
// e.g. DecrementFloat64.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementFloat(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 float64) error {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	switch __obf_298afaabecd63bb9.Object.(type) {
	case float32:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(float32) - float32(__obf_4029f0efebfc66c6)
	case float64:
		__obf_298afaabecd63bb9.Object = __obf_298afaabecd63bb9.Object.(float64) - __obf_4029f0efebfc66c6
	default:
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_c0676b3110071e57)
	}
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return nil
}

// Decrement an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementInt(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int) (int, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(int)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementInt8(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int8) (int8, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(int8)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementInt16(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int16) (int16, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(int16)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementInt32(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int32) (int32, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(int32)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementInt64(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 int64) (int64, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(int64)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementUint(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uint) (uint, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uint)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementUintptr(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uintptr) (uintptr, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uintptr)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type uint8 by n. Returns an error if the item's value is
// not an uint8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementUint8(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uint8) (uint8, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uint8)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementUint16(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uint16) (uint16, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uint16)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementUint32(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uint32) (uint32, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uint32)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementUint64(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 uint64) (uint64, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(uint64)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementFloat32(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 float32) (float32, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(float32)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Decrement an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DecrementFloat64(__obf_c0676b3110071e57 string, __obf_4029f0efebfc66c6 float64) (float64, error) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
	if !__obf_889b1fa74a37b35f || __obf_298afaabecd63bb9.Expired() {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_c0676b3110071e57)
	}
	__obf_ac1e250d3e0db38a, __obf_4bca5d4f1d82d040 := __obf_298afaabecd63bb9.Object.(float64)
	if !__obf_4bca5d4f1d82d040 {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_c0676b3110071e57)
	}
	__obf_6125f7eddefe2170 := __obf_ac1e250d3e0db38a - __obf_4029f0efebfc66c6
	__obf_298afaabecd63bb9.Object = __obf_6125f7eddefe2170
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	return __obf_6125f7eddefe2170, nil
}

// Delete an item from the cache. Does nothing if the key is not in the cache.
func (__obf_1e8c1f17e8f9e9cd *MapCache) Delete(__obf_c0676b3110071e57 string) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_298afaabecd63bb9, __obf_02de68c5ead58da0 := __obf_1e8c1f17e8f9e9cd.delete(__obf_c0676b3110071e57)
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	if __obf_02de68c5ead58da0 {
		__obf_1e8c1f17e8f9e9cd.__obf_f8b7842824568810(__obf_c0676b3110071e57, __obf_298afaabecd63bb9)
	}
}

func (__obf_1e8c1f17e8f9e9cd *MapCache) delete(__obf_c0676b3110071e57 string) (any, bool) {
	if __obf_1e8c1f17e8f9e9cd.__obf_f8b7842824568810 != nil {
		if __obf_298afaabecd63bb9, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]; __obf_889b1fa74a37b35f {
			delete(__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55, __obf_c0676b3110071e57)
			return __obf_298afaabecd63bb9.Object, true
		}
	}
	delete(__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55, __obf_c0676b3110071e57)
	return nil, false
}

type __obf_cb01e7d7e10da71a struct {
	__obf_4aecf3c737bbe5e8 string
	__obf_b99b98b3128d77a4 any
}

// Delete all expired items from the cache.
func (__obf_1e8c1f17e8f9e9cd *MapCache) DeleteExpired() {
	var __obf_f888b8e41822f034 []__obf_cb01e7d7e10da71a
	__obf_c5d09205c91947fd := time.Now().UnixNano()
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	for __obf_c0676b3110071e57, __obf_298afaabecd63bb9 := range __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55 {
		// "Inlining" of expired
		if __obf_298afaabecd63bb9.Expiration > 0 && __obf_c5d09205c91947fd > __obf_298afaabecd63bb9.Expiration {
			__obf_5d110475fa35a207, __obf_02de68c5ead58da0 := __obf_1e8c1f17e8f9e9cd.delete(__obf_c0676b3110071e57)
			if __obf_02de68c5ead58da0 {
				__obf_f888b8e41822f034 = append(__obf_f888b8e41822f034, __obf_cb01e7d7e10da71a{__obf_c0676b3110071e57, __obf_5d110475fa35a207})
			}
		}
	}
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
	for _, __obf_298afaabecd63bb9 := range __obf_f888b8e41822f034 {
		__obf_1e8c1f17e8f9e9cd.__obf_f8b7842824568810(__obf_298afaabecd63bb9.__obf_4aecf3c737bbe5e8, __obf_298afaabecd63bb9.__obf_b99b98b3128d77a4)
	}
}

// Sets an (optional) function that is called with the key and value when an
// item is evicted from the cache. (Including when it is deleted manually, but
// not when it is overwritten.) Set to nil to disable.
func (__obf_1e8c1f17e8f9e9cd *MapCache) OnEvicted(__obf_e0505be006a4623d func(string, any)) {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_1e8c1f17e8f9e9cd.__obf_f8b7842824568810 = __obf_e0505be006a4623d
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
}

// Write the cache's items (using Gob) to an io.Writer.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_1e8c1f17e8f9e9cd *MapCache) Save(__obf_a6e67633216a641b io.Writer) (__obf_0560c6b8f27080cc error) {
	__obf_fc591f18406f4079 := gob.NewEncoder(__obf_a6e67633216a641b)
	defer func() {
		if __obf_901366c1b9787938 := recover(); __obf_901366c1b9787938 != nil {
			__obf_0560c6b8f27080cc = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RLock()
	defer __obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RUnlock()
	for _, __obf_298afaabecd63bb9 := range __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55 {
		gob.Register(__obf_298afaabecd63bb9.Object)
	}
	__obf_0560c6b8f27080cc = __obf_fc591f18406f4079.Encode(&__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_1e8c1f17e8f9e9cd *MapCache) SaveFile(__obf_64b4286d3b68414f string) error {
	__obf_151073dbfed2e274, __obf_0560c6b8f27080cc := os.Create(__obf_64b4286d3b68414f)
	if __obf_0560c6b8f27080cc != nil {
		return __obf_0560c6b8f27080cc
	}
	__obf_0560c6b8f27080cc = __obf_1e8c1f17e8f9e9cd.Save(__obf_151073dbfed2e274)
	if __obf_0560c6b8f27080cc != nil {
		__obf_151073dbfed2e274.Close()
		return __obf_0560c6b8f27080cc
	}
	return __obf_151073dbfed2e274.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_1e8c1f17e8f9e9cd *MapCache) Load(__obf_2d21271215a2e06f io.Reader) error {
	__obf_525d77ca92aebf39 := gob.NewDecoder(__obf_2d21271215a2e06f)
	__obf_1d5d1defae6f4d55 := map[string]Item{}
	__obf_0560c6b8f27080cc := __obf_525d77ca92aebf39.Decode(&__obf_1d5d1defae6f4d55)
	if __obf_0560c6b8f27080cc == nil {
		__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
		defer __obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
		for __obf_c0676b3110071e57, __obf_298afaabecd63bb9 := range __obf_1d5d1defae6f4d55 {
			__obf_5d110475fa35a207, __obf_889b1fa74a37b35f := __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57]
			if !__obf_889b1fa74a37b35f || __obf_5d110475fa35a207.Expired() {
				__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
			}
		}
	}
	return __obf_0560c6b8f27080cc
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_1e8c1f17e8f9e9cd *MapCache) LoadFile(__obf_64b4286d3b68414f string) error {
	__obf_151073dbfed2e274, __obf_0560c6b8f27080cc := os.Open(__obf_64b4286d3b68414f)
	if __obf_0560c6b8f27080cc != nil {
		return __obf_0560c6b8f27080cc
	}
	__obf_0560c6b8f27080cc = __obf_1e8c1f17e8f9e9cd.Load(__obf_151073dbfed2e274)
	if __obf_0560c6b8f27080cc != nil {
		__obf_151073dbfed2e274.Close()
		return __obf_0560c6b8f27080cc
	}
	return __obf_151073dbfed2e274.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_1e8c1f17e8f9e9cd *MapCache) Items() map[string]Item {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RLock()
	defer __obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RUnlock()
	__obf_17fa7ee2c9072df2 := make(map[string]Item, len(__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55))
	__obf_c5d09205c91947fd := time.Now().UnixNano()
	for __obf_c0676b3110071e57, __obf_298afaabecd63bb9 := range __obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55 {
		// "Inlining" of Expired
		if __obf_298afaabecd63bb9.Expiration > 0 {
			if __obf_c5d09205c91947fd > __obf_298afaabecd63bb9.Expiration {
				continue
			}
		}
		__obf_17fa7ee2c9072df2[__obf_c0676b3110071e57] = __obf_298afaabecd63bb9
	}
	return __obf_17fa7ee2c9072df2
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_1e8c1f17e8f9e9cd *MapCache) ItemCount() int {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RLock()
	__obf_4029f0efebfc66c6 := len(__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55)
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.RUnlock()
	return __obf_4029f0efebfc66c6
}

// Delete all items from the cache.
func (__obf_1e8c1f17e8f9e9cd *MapCache) Flush() {
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Lock()
	__obf_1e8c1f17e8f9e9cd.__obf_1d5d1defae6f4d55 = map[string]Item{}
	__obf_1e8c1f17e8f9e9cd.__obf_bc58716fda429cf7.Unlock()
}

type __obf_4336a7018604d0e4 struct {
	Interval               time.Duration
	__obf_fa671b0e1a5b525c chan bool
}

func (__obf_7a518f537eee8b51 *__obf_4336a7018604d0e4) Run(__obf_1e8c1f17e8f9e9cd *MapCache) {
	__obf_80ac7a191daad925 := time.NewTicker(__obf_7a518f537eee8b51.Interval)
	for {
		select {
		case <-__obf_80ac7a191daad925.C:
			__obf_1e8c1f17e8f9e9cd.DeleteExpired()
		case <-__obf_7a518f537eee8b51.__obf_fa671b0e1a5b525c:
			__obf_80ac7a191daad925.Stop()
			return
		}
	}
}

func __obf_e13d534b20b3cd03(__obf_1e8c1f17e8f9e9cd *MapCache) {
	__obf_1e8c1f17e8f9e9cd.__obf_4336a7018604d0e4.__obf_fa671b0e1a5b525c <- true
}

func __obf_fc9c5934131c5305(__obf_1e8c1f17e8f9e9cd *MapCache, __obf_1cad2f68cf3c415c time.Duration) {
	__obf_7a518f537eee8b51 := &__obf_4336a7018604d0e4{
		Interval:               __obf_1cad2f68cf3c415c,
		__obf_fa671b0e1a5b525c: make(chan bool),
	}
	__obf_1e8c1f17e8f9e9cd.__obf_4336a7018604d0e4 = __obf_7a518f537eee8b51
	go __obf_7a518f537eee8b51.Run(__obf_1e8c1f17e8f9e9cd)
}

func __obf_927b36b4760ff129(__obf_1594582d7bf2e9f4 time.Duration, __obf_17fa7ee2c9072df2 map[string]Item) *MapCache {
	if __obf_1594582d7bf2e9f4 == 0 {
		__obf_1594582d7bf2e9f4 = -1
	}
	__obf_1e8c1f17e8f9e9cd := &MapCache{
		__obf_bc351f72e9074c26: __obf_1594582d7bf2e9f4,
		__obf_1d5d1defae6f4d55: __obf_17fa7ee2c9072df2,
	}
	return __obf_1e8c1f17e8f9e9cd
}

func __obf_13c405efc9fcc35c(__obf_1594582d7bf2e9f4 time.Duration, __obf_1cad2f68cf3c415c time.Duration, __obf_17fa7ee2c9072df2 map[string]Item) *MapCache {
	__obf_1e8c1f17e8f9e9cd := __obf_927b36b4760ff129(__obf_1594582d7bf2e9f4, __obf_17fa7ee2c9072df2)
	// This trick ensures that the janitor goroutine (which--granted it
	// was enabled--is running DeleteExpired on c forever) does not keep
	// the returned C object from being garbage collected. When it is
	// garbage collected, the finalizer stops the janitor goroutine, after
	// which c can be collected.
	// C := &MapCache{c}
	if __obf_1cad2f68cf3c415c > 0 {
		__obf_fc9c5934131c5305(__obf_1e8c1f17e8f9e9cd, __obf_1cad2f68cf3c415c)
		runtime.SetFinalizer(__obf_1e8c1f17e8f9e9cd, __obf_e13d534b20b3cd03)
	}
	return __obf_1e8c1f17e8f9e9cd
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_bc351f72e9074c26, __obf_b99284a895226c6b time.Duration) *MapCache {
	__obf_1d5d1defae6f4d55 := make(map[string]Item)
	return __obf_13c405efc9fcc35c(__obf_bc351f72e9074c26, __obf_b99284a895226c6b, __obf_1d5d1defae6f4d55)
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
func NewFrom(__obf_bc351f72e9074c26, __obf_b99284a895226c6b time.Duration, __obf_1d5d1defae6f4d55 map[string]Item) *MapCache {
	return __obf_13c405efc9fcc35c(__obf_bc351f72e9074c26, __obf_b99284a895226c6b, __obf_1d5d1defae6f4d55)
}
