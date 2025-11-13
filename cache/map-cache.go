package __obf_f4d8558b35981340

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
func (__obf_a7d9bc40b5640158 Item) Expired() bool {
	if __obf_a7d9bc40b5640158.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_a7d9bc40b5640158.Expiration
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
	__obf_cf2ab0a31fea39d5 time.Duration
	__obf_d71a79b74fa6f45f map[string]Item
	__obf_03f513e727a97a9a sync.RWMutex
	__obf_79fd0a449fb86f69 func(string, any)
	__obf_19deb58e2726fc71 *__obf_19deb58e2726fc71
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
func (__obf_a744f58475749237 *MapCache) Set(__obf_f038738651cdb77b string, __obf_916d170f0073a688 any, __obf_dc93050659c4ccb6 time.Duration) {
	// "Inlining" of set
	var __obf_06e4a4e6e9269bd7 int64
	if __obf_dc93050659c4ccb6 == DefaultExpiration {
		__obf_dc93050659c4ccb6 = __obf_a744f58475749237.__obf_cf2ab0a31fea39d5
	}
	if __obf_dc93050659c4ccb6 > 0 {
		__obf_06e4a4e6e9269bd7 = time.Now().Add(__obf_dc93050659c4ccb6).UnixNano()
	}
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = Item{
		Object:     __obf_916d170f0073a688,
		Expiration: __obf_06e4a4e6e9269bd7,
	}
	// TODO: Calls to mu.Unlock are currently not deferred because defer
	// adds ~200 ns (as of go1.)
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
}

func (__obf_a744f58475749237 *MapCache) __obf_c1e574bfca2b9bfd(__obf_f038738651cdb77b string, __obf_916d170f0073a688 any, __obf_dc93050659c4ccb6 time.Duration) {
	var __obf_06e4a4e6e9269bd7 int64
	if __obf_dc93050659c4ccb6 == DefaultExpiration {
		__obf_dc93050659c4ccb6 = __obf_a744f58475749237.__obf_cf2ab0a31fea39d5
	}
	if __obf_dc93050659c4ccb6 > 0 {
		__obf_06e4a4e6e9269bd7 = time.Now().Add(__obf_dc93050659c4ccb6).UnixNano()
	}
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = Item{
		Object:     __obf_916d170f0073a688,
		Expiration: __obf_06e4a4e6e9269bd7,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_a744f58475749237 *MapCache) SetDefault(__obf_f038738651cdb77b string, __obf_916d170f0073a688 any) {
	__obf_a744f58475749237.Set(__obf_f038738651cdb77b, __obf_916d170f0073a688, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_a744f58475749237 *MapCache) Add(__obf_f038738651cdb77b string, __obf_916d170f0073a688 any, __obf_dc93050659c4ccb6 time.Duration) error {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	_, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_7587f321ec867980(__obf_f038738651cdb77b)
	if __obf_9d182609c2b1ed7d {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return fmt.Errorf("Item %s already exists", __obf_f038738651cdb77b)
	}
	__obf_a744f58475749237.__obf_c1e574bfca2b9bfd(__obf_f038738651cdb77b, __obf_916d170f0073a688, __obf_dc93050659c4ccb6)
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return nil
}

// Set a new value for the cache key only if it already exists, and the existing
// item hasn't expired. Returns an error otherwise.
func (__obf_a744f58475749237 *MapCache) Replace(__obf_f038738651cdb77b string, __obf_916d170f0073a688 any, __obf_dc93050659c4ccb6 time.Duration) error {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	_, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_7587f321ec867980(__obf_f038738651cdb77b)
	if !__obf_9d182609c2b1ed7d {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_f038738651cdb77b)
	}
	__obf_a744f58475749237.__obf_c1e574bfca2b9bfd(__obf_f038738651cdb77b, __obf_916d170f0073a688, __obf_dc93050659c4ccb6)
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return nil
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
// whether the key was found.
func (__obf_a744f58475749237 *MapCache) Get(__obf_f038738651cdb77b string) (any, bool) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.RLock()
	// "Inlining" of get and Expired
	__obf_a7d9bc40b5640158, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.RUnlock()
		return nil, false
	}
	if __obf_a7d9bc40b5640158.Expiration > 0 {
		if time.Now().UnixNano() > __obf_a7d9bc40b5640158.Expiration {
			__obf_a744f58475749237.__obf_03f513e727a97a9a.RUnlock()
			return nil, false
		}
	}
	__obf_a744f58475749237.__obf_03f513e727a97a9a.RUnlock()
	return __obf_a7d9bc40b5640158.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_a744f58475749237 *MapCache) GetWithExpiration(__obf_f038738651cdb77b string) (any, time.Time, bool) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.RLock()
	// "Inlining" of get and Expired
	__obf_a7d9bc40b5640158, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_a7d9bc40b5640158.Expiration > 0 {
		if time.Now().UnixNano() > __obf_a7d9bc40b5640158.Expiration {
			__obf_a744f58475749237.__obf_03f513e727a97a9a.RUnlock()
			return nil, time.Time{}, false
		}

		// Return the item and the expiration time
		__obf_a744f58475749237.__obf_03f513e727a97a9a.RUnlock()
		return __obf_a7d9bc40b5640158.Object, time.Unix(0, __obf_a7d9bc40b5640158.Expiration), true
	}

	// If expiration <= 0 (i.e. no expiration time set) then return the item
	// and a zeroed time.Time
	__obf_a744f58475749237.__obf_03f513e727a97a9a.RUnlock()
	return __obf_a7d9bc40b5640158.Object, time.Time{}, true
}

func (__obf_a744f58475749237 *MapCache) __obf_7587f321ec867980(__obf_f038738651cdb77b string) (any, bool) {
	__obf_a7d9bc40b5640158, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_a7d9bc40b5640158.Expiration > 0 {
		if time.Now().UnixNano() > __obf_a7d9bc40b5640158.Expiration {
			return nil, false
		}
	}
	return __obf_a7d9bc40b5640158.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_a744f58475749237 *MapCache) Increment(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int64) error {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	switch __obf_4153159a00c5e3bb.Object.(type) {
	case int:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(int) + int(__obf_97c11419d5893183)
	case int8:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(int8) + int8(__obf_97c11419d5893183)
	case int16:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(int16) + int16(__obf_97c11419d5893183)
	case int32:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(int32) + int32(__obf_97c11419d5893183)
	case int64:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(int64) + __obf_97c11419d5893183
	case uint:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uint) + uint(__obf_97c11419d5893183)
	case uintptr:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uintptr) + uintptr(__obf_97c11419d5893183)
	case uint8:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uint8) + uint8(__obf_97c11419d5893183)
	case uint16:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uint16) + uint16(__obf_97c11419d5893183)
	case uint32:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uint32) + uint32(__obf_97c11419d5893183)
	case uint64:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uint64) + uint64(__obf_97c11419d5893183)
	case float32:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(float32) + float32(__obf_97c11419d5893183)
	case float64:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(float64) + float64(__obf_97c11419d5893183)
	default:
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_f038738651cdb77b)
	}
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return nil
}

// Increment an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to increment it by n. Pass a negative number to decrement the
// value. To retrieve the incremented value, use one of the specialized methods,
// e.g. IncrementFloat64.
func (__obf_a744f58475749237 *MapCache) IncrementFloat(__obf_f038738651cdb77b string, __obf_97c11419d5893183 float64) error {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	switch __obf_4153159a00c5e3bb.Object.(type) {
	case float32:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(float32) + float32(__obf_97c11419d5893183)
	case float64:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(float64) + __obf_97c11419d5893183
	default:
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_f038738651cdb77b)
	}
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return nil
}

// Increment an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementInt(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int) (int, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(int)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementInt8(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int8) (int8, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(int8)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementInt16(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int16) (int16, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(int16)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementInt32(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int32) (int32, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(int32)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementInt64(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int64) (int64, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(int64)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the incremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementUint(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uint) (uint, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uint)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementUintptr(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uintptr) (uintptr, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uintptr)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type uint8 by n. Returns an error if the item's value
// is not an uint8, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementUint8(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uint8) (uint8, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uint8)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementUint16(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uint16) (uint16, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uint16)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementUint32(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uint32) (uint32, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uint32)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementUint64(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uint64) (uint64, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uint64)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementFloat32(__obf_f038738651cdb77b string, __obf_97c11419d5893183 float32) (float32, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(float32)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Increment an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// incremented value is returned.
func (__obf_a744f58475749237 *MapCache) IncrementFloat64(__obf_f038738651cdb77b string, __obf_97c11419d5893183 float64) (float64, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(float64)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 + __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to decrement it by n. To retrieve the decremented value, use one
// of the specialized methods, e.g. DecrementInt64.
func (__obf_a744f58475749237 *MapCache) Decrement(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int64) error {
	// TODO: Implement Increment and Decrement more cleanly.
	// (Cannot do Increment(k, n*-1) for uints.)
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_4153159a00c5e3bb.Object.(type) {
	case int:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(int) - int(__obf_97c11419d5893183)
	case int8:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(int8) - int8(__obf_97c11419d5893183)
	case int16:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(int16) - int16(__obf_97c11419d5893183)
	case int32:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(int32) - int32(__obf_97c11419d5893183)
	case int64:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(int64) - __obf_97c11419d5893183
	case uint:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uint) - uint(__obf_97c11419d5893183)
	case uintptr:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uintptr) - uintptr(__obf_97c11419d5893183)
	case uint8:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uint8) - uint8(__obf_97c11419d5893183)
	case uint16:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uint16) - uint16(__obf_97c11419d5893183)
	case uint32:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uint32) - uint32(__obf_97c11419d5893183)
	case uint64:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(uint64) - uint64(__obf_97c11419d5893183)
	case float32:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(float32) - float32(__obf_97c11419d5893183)
	case float64:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(float64) - float64(__obf_97c11419d5893183)
	default:
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_f038738651cdb77b)
	}
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return nil
}

// Decrement an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to decrement it by n. Pass a negative number to decrement the
// value. To retrieve the decremented value, use one of the specialized methods,
// e.g. DecrementFloat64.
func (__obf_a744f58475749237 *MapCache) DecrementFloat(__obf_f038738651cdb77b string, __obf_97c11419d5893183 float64) error {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	switch __obf_4153159a00c5e3bb.Object.(type) {
	case float32:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(float32) - float32(__obf_97c11419d5893183)
	case float64:
		__obf_4153159a00c5e3bb.Object = __obf_4153159a00c5e3bb.Object.(float64) - __obf_97c11419d5893183
	default:
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_f038738651cdb77b)
	}
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return nil
}

// Decrement an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementInt(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int) (int, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(int)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementInt8(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int8) (int8, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(int8)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementInt16(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int16) (int16, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(int16)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementInt32(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int32) (int32, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(int32)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementInt64(__obf_f038738651cdb77b string, __obf_97c11419d5893183 int64) (int64, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(int64)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementUint(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uint) (uint, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uint)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementUintptr(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uintptr) (uintptr, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uintptr)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type uint8 by n. Returns an error if the item's value is
// not an uint8, or if it was not found. If there is no error, the decremented
// value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementUint8(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uint8) (uint8, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uint8)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementUint16(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uint16) (uint16, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uint16)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementUint32(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uint32) (uint32, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uint32)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementUint64(__obf_f038738651cdb77b string, __obf_97c11419d5893183 uint64) (uint64, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(uint64)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementFloat32(__obf_f038738651cdb77b string, __obf_97c11419d5893183 float32) (float32, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(float32)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Decrement an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// decremented value is returned.
func (__obf_a744f58475749237 *MapCache) DecrementFloat64(__obf_f038738651cdb77b string, __obf_97c11419d5893183 float64) (float64, error) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
	if !__obf_9d182609c2b1ed7d || __obf_4153159a00c5e3bb.Expired() {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_f038738651cdb77b)
	}
	__obf_e56652c25521a1c4, __obf_ed10eafa7bc79f22 := __obf_4153159a00c5e3bb.Object.(float64)
	if !__obf_ed10eafa7bc79f22 {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_f038738651cdb77b)
	}
	__obf_9314eb419322e5c3 := __obf_e56652c25521a1c4 - __obf_97c11419d5893183
	__obf_4153159a00c5e3bb.Object = __obf_9314eb419322e5c3
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	return __obf_9314eb419322e5c3, nil
}

// Delete an item from the cache. Does nothing if the key is not in the cache.
func (__obf_a744f58475749237 *MapCache) Delete(__obf_f038738651cdb77b string) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_4153159a00c5e3bb, __obf_043116a1e6a75d57 := __obf_a744f58475749237.delete(__obf_f038738651cdb77b)
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	if __obf_043116a1e6a75d57 {
		__obf_a744f58475749237.__obf_79fd0a449fb86f69(__obf_f038738651cdb77b, __obf_4153159a00c5e3bb)
	}
}

func (__obf_a744f58475749237 *MapCache) delete(__obf_f038738651cdb77b string) (any, bool) {
	if __obf_a744f58475749237.__obf_79fd0a449fb86f69 != nil {
		if __obf_4153159a00c5e3bb, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]; __obf_9d182609c2b1ed7d {
			delete(__obf_a744f58475749237.__obf_d71a79b74fa6f45f, __obf_f038738651cdb77b)
			return __obf_4153159a00c5e3bb.Object, true
		}
	}
	delete(__obf_a744f58475749237.__obf_d71a79b74fa6f45f, __obf_f038738651cdb77b)
	return nil, false
}

type __obf_3947fde927b0f446 struct {
	__obf_6208471da03258dd string
	__obf_5c03cd51840607fb any
}

// Delete all expired items from the cache.
func (__obf_a744f58475749237 *MapCache) DeleteExpired() {
	var __obf_4e3e6b2b6f770188 []__obf_3947fde927b0f446
	__obf_7bd993b82a2d9d80 := time.Now().UnixNano()
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	for __obf_f038738651cdb77b, __obf_4153159a00c5e3bb := range __obf_a744f58475749237.__obf_d71a79b74fa6f45f {
		// "Inlining" of expired
		if __obf_4153159a00c5e3bb.Expiration > 0 && __obf_7bd993b82a2d9d80 > __obf_4153159a00c5e3bb.Expiration {
			__obf_38f821202457eb68, __obf_043116a1e6a75d57 := __obf_a744f58475749237.delete(__obf_f038738651cdb77b)
			if __obf_043116a1e6a75d57 {
				__obf_4e3e6b2b6f770188 = append(__obf_4e3e6b2b6f770188, __obf_3947fde927b0f446{__obf_f038738651cdb77b, __obf_38f821202457eb68})
			}
		}
	}
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
	for _, __obf_4153159a00c5e3bb := range __obf_4e3e6b2b6f770188 {
		__obf_a744f58475749237.__obf_79fd0a449fb86f69(__obf_4153159a00c5e3bb.__obf_6208471da03258dd, __obf_4153159a00c5e3bb.__obf_5c03cd51840607fb)
	}
}

// Sets an (optional) function that is called with the key and value when an
// item is evicted from the cache. (Including when it is deleted manually, but
// not when it is overwritten.) Set to nil to disable.
func (__obf_a744f58475749237 *MapCache) OnEvicted(__obf_d1304c07ee90999d func(string, any)) {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_a744f58475749237.__obf_79fd0a449fb86f69 = __obf_d1304c07ee90999d
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
}

// Write the cache's items (using Gob) to an io.Writer.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_a744f58475749237 *MapCache) Save(__obf_250f503e3b4e0efe io.Writer) (__obf_2d5e4b42e300bedd error) {
	__obf_ed8ff5b064352a0d := gob.NewEncoder(__obf_250f503e3b4e0efe)
	defer func() {
		if __obf_916d170f0073a688 := recover(); __obf_916d170f0073a688 != nil {
			__obf_2d5e4b42e300bedd = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_a744f58475749237.__obf_03f513e727a97a9a.RLock()
	defer __obf_a744f58475749237.__obf_03f513e727a97a9a.RUnlock()
	for _, __obf_4153159a00c5e3bb := range __obf_a744f58475749237.__obf_d71a79b74fa6f45f {
		gob.Register(__obf_4153159a00c5e3bb.Object)
	}
	__obf_2d5e4b42e300bedd = __obf_ed8ff5b064352a0d.Encode(&__obf_a744f58475749237.__obf_d71a79b74fa6f45f)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_a744f58475749237 *MapCache) SaveFile(__obf_7061d7d6bba30c80 string) error {
	__obf_22d2b5af70e18063, __obf_2d5e4b42e300bedd := os.Create(__obf_7061d7d6bba30c80)
	if __obf_2d5e4b42e300bedd != nil {
		return __obf_2d5e4b42e300bedd
	}
	__obf_2d5e4b42e300bedd = __obf_a744f58475749237.Save(__obf_22d2b5af70e18063)
	if __obf_2d5e4b42e300bedd != nil {
		__obf_22d2b5af70e18063.Close()
		return __obf_2d5e4b42e300bedd
	}
	return __obf_22d2b5af70e18063.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_a744f58475749237 *MapCache) Load(__obf_ed249696e6d62c67 io.Reader) error {
	__obf_2f08859deefeafb8 := gob.NewDecoder(__obf_ed249696e6d62c67)
	__obf_d71a79b74fa6f45f := map[string]Item{}
	__obf_2d5e4b42e300bedd := __obf_2f08859deefeafb8.Decode(&__obf_d71a79b74fa6f45f)
	if __obf_2d5e4b42e300bedd == nil {
		__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
		defer __obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
		for __obf_f038738651cdb77b, __obf_4153159a00c5e3bb := range __obf_d71a79b74fa6f45f {
			__obf_38f821202457eb68, __obf_9d182609c2b1ed7d := __obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b]
			if !__obf_9d182609c2b1ed7d || __obf_38f821202457eb68.Expired() {
				__obf_a744f58475749237.__obf_d71a79b74fa6f45f[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
			}
		}
	}
	return __obf_2d5e4b42e300bedd
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_a744f58475749237 *MapCache) LoadFile(__obf_7061d7d6bba30c80 string) error {
	__obf_22d2b5af70e18063, __obf_2d5e4b42e300bedd := os.Open(__obf_7061d7d6bba30c80)
	if __obf_2d5e4b42e300bedd != nil {
		return __obf_2d5e4b42e300bedd
	}
	__obf_2d5e4b42e300bedd = __obf_a744f58475749237.Load(__obf_22d2b5af70e18063)
	if __obf_2d5e4b42e300bedd != nil {
		__obf_22d2b5af70e18063.Close()
		return __obf_2d5e4b42e300bedd
	}
	return __obf_22d2b5af70e18063.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_a744f58475749237 *MapCache) Items() map[string]Item {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.RLock()
	defer __obf_a744f58475749237.__obf_03f513e727a97a9a.RUnlock()
	__obf_1f6d94b11ce336b6 := make(map[string]Item, len(__obf_a744f58475749237.__obf_d71a79b74fa6f45f))
	__obf_7bd993b82a2d9d80 := time.Now().UnixNano()
	for __obf_f038738651cdb77b, __obf_4153159a00c5e3bb := range __obf_a744f58475749237.__obf_d71a79b74fa6f45f {
		// "Inlining" of Expired
		if __obf_4153159a00c5e3bb.Expiration > 0 {
			if __obf_7bd993b82a2d9d80 > __obf_4153159a00c5e3bb.Expiration {
				continue
			}
		}
		__obf_1f6d94b11ce336b6[__obf_f038738651cdb77b] = __obf_4153159a00c5e3bb
	}
	return __obf_1f6d94b11ce336b6
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_a744f58475749237 *MapCache) ItemCount() int {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.RLock()
	__obf_97c11419d5893183 := len(__obf_a744f58475749237.__obf_d71a79b74fa6f45f)
	__obf_a744f58475749237.__obf_03f513e727a97a9a.RUnlock()
	return __obf_97c11419d5893183
}

// Delete all items from the cache.
func (__obf_a744f58475749237 *MapCache) Flush() {
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Lock()
	__obf_a744f58475749237.__obf_d71a79b74fa6f45f = map[string]Item{}
	__obf_a744f58475749237.__obf_03f513e727a97a9a.Unlock()
}

type __obf_19deb58e2726fc71 struct {
	Interval               time.Duration
	__obf_18aeb9387671549f chan bool
}

func (__obf_4dbcdaa111fffea2 *__obf_19deb58e2726fc71) Run(__obf_a744f58475749237 *MapCache) {
	__obf_701fbf68de3b6fb0 := time.NewTicker(__obf_4dbcdaa111fffea2.Interval)
	for {
		select {
		case <-__obf_701fbf68de3b6fb0.C:
			__obf_a744f58475749237.DeleteExpired()
		case <-__obf_4dbcdaa111fffea2.__obf_18aeb9387671549f:
			__obf_701fbf68de3b6fb0.Stop()
			return
		}
	}
}

func __obf_c776d70f31a46b61(__obf_a744f58475749237 *MapCache) {
	__obf_a744f58475749237.__obf_19deb58e2726fc71.__obf_18aeb9387671549f <- true
}

func __obf_1b9130a90a60be3d(__obf_a744f58475749237 *MapCache, __obf_731e158dd27c4bb2 time.Duration) {
	__obf_4dbcdaa111fffea2 := &__obf_19deb58e2726fc71{
		Interval:               __obf_731e158dd27c4bb2,
		__obf_18aeb9387671549f: make(chan bool),
	}
	__obf_a744f58475749237.__obf_19deb58e2726fc71 = __obf_4dbcdaa111fffea2
	go __obf_4dbcdaa111fffea2.Run(__obf_a744f58475749237)
}

func __obf_35daf1d33cf0d6e0(__obf_c7b483f992c088c6 time.Duration, __obf_1f6d94b11ce336b6 map[string]Item) *MapCache {
	if __obf_c7b483f992c088c6 == 0 {
		__obf_c7b483f992c088c6 = -1
	}
	__obf_a744f58475749237 := &MapCache{
		__obf_cf2ab0a31fea39d5: __obf_c7b483f992c088c6,
		__obf_d71a79b74fa6f45f: __obf_1f6d94b11ce336b6,
	}
	return __obf_a744f58475749237
}

func __obf_13b9aa2d454254f7(__obf_c7b483f992c088c6 time.Duration, __obf_731e158dd27c4bb2 time.Duration, __obf_1f6d94b11ce336b6 map[string]Item) *MapCache {
	__obf_a744f58475749237 := __obf_35daf1d33cf0d6e0(__obf_c7b483f992c088c6, __obf_1f6d94b11ce336b6)
	// This trick ensures that the janitor goroutine (which--granted it
	// was enabled--is running DeleteExpired on c forever) does not keep
	// the returned C object from being garbage collected. When it is
	// garbage collected, the finalizer stops the janitor goroutine, after
	// which c can be collected.
	// C := &MapCache{c}
	if __obf_731e158dd27c4bb2 > 0 {
		__obf_1b9130a90a60be3d(__obf_a744f58475749237, __obf_731e158dd27c4bb2)
		runtime.SetFinalizer(__obf_a744f58475749237, __obf_c776d70f31a46b61)
	}
	return __obf_a744f58475749237
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_cf2ab0a31fea39d5, __obf_59111e969b33b396 time.Duration) *MapCache {
	__obf_d71a79b74fa6f45f := make(map[string]Item)
	return __obf_13b9aa2d454254f7(__obf_cf2ab0a31fea39d5, __obf_59111e969b33b396, __obf_d71a79b74fa6f45f)
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
func NewFrom(__obf_cf2ab0a31fea39d5, __obf_59111e969b33b396 time.Duration, __obf_d71a79b74fa6f45f map[string]Item) *MapCache {
	return __obf_13b9aa2d454254f7(__obf_cf2ab0a31fea39d5, __obf_59111e969b33b396, __obf_d71a79b74fa6f45f)
}
