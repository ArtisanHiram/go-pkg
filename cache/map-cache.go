package __obf_6b306490bf7221d3

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
func (__obf_54b1e5958ef673f0 Item) Expired() bool {
	if __obf_54b1e5958ef673f0.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_54b1e5958ef673f0.Expiration
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
	__obf_efa9e45df3c18f44 time. // *cache
		// If this is confusing, see the comment at the bottom of New()
		Duration
	__obf_b4d587b1629307c0 map[string]Item
	__obf_2aacabd3f44c4b79 sync.RWMutex
	__obf_8c76d4f25d47695a func(string, any)
	__obf_cc3d7a1795ac42ab *__obf_cc3d7a1795ac42ab
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
func (__obf_6ccac4bb8646a806 *MapCache) Set(__obf_afff83fa0e852ce0 string, __obf_3ab13e197ed6fe6b any, __obf_b39c53bb7d5c0767 time.Duration) {
	// "Inlining" of set
	var __obf_be37eed188feee9f int64
	if __obf_b39c53bb7d5c0767 == DefaultExpiration {
		__obf_b39c53bb7d5c0767 = __obf_6ccac4bb8646a806.__obf_efa9e45df3c18f44
	}
	if __obf_b39c53bb7d5c0767 > 0 {
		__obf_be37eed188feee9f = time.Now().Add(__obf_b39c53bb7d5c0767).UnixNano()
	}
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = Item{
		Object:     __obf_3ab13e197ed6fe6b,
		Expiration: __obf_be37eed188feee9f,
	}
	__obf_6ccac4bb8646a806.
		// TODO: Calls to mu.Unlock are currently not deferred because defer
		// adds ~200 ns (as of go1.)
		__obf_2aacabd3f44c4b79.
		Unlock()
}

func (__obf_6ccac4bb8646a806 *MapCache) __obf_930bee944f85c808(__obf_afff83fa0e852ce0 string, __obf_3ab13e197ed6fe6b any, __obf_b39c53bb7d5c0767 time.Duration) {
	var __obf_be37eed188feee9f int64
	if __obf_b39c53bb7d5c0767 == DefaultExpiration {
		__obf_b39c53bb7d5c0767 = __obf_6ccac4bb8646a806.__obf_efa9e45df3c18f44
	}
	if __obf_b39c53bb7d5c0767 > 0 {
		__obf_be37eed188feee9f = time.Now().Add(__obf_b39c53bb7d5c0767).UnixNano()
	}
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = Item{
		Object:     __obf_3ab13e197ed6fe6b,
		Expiration: __obf_be37eed188feee9f,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_6ccac4bb8646a806 *MapCache) SetDefault(__obf_afff83fa0e852ce0 string, __obf_3ab13e197ed6fe6b any) {
	__obf_6ccac4bb8646a806.
		Set(__obf_afff83fa0e852ce0, __obf_3ab13e197ed6fe6b, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_6ccac4bb8646a806 *MapCache) Add(__obf_afff83fa0e852ce0 string, __obf_3ab13e197ed6fe6b any, __obf_b39c53bb7d5c0767 time.Duration) error {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	_, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b66283acccf8c10b(__obf_afff83fa0e852ce0)
	if __obf_69b7b10d07db29c0 {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return fmt.Errorf("Item %s already exists", __obf_afff83fa0e852ce0)
	}
	__obf_6ccac4bb8646a806.__obf_930bee944f85c808(__obf_afff83fa0e852ce0, __obf_3ab13e197ed6fe6b,

		// Set a new value for the cache key only if it already exists, and the existing
		// item hasn't expired. Returns an error otherwise.
		__obf_b39c53bb7d5c0767)
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Unlock()
	return nil
}

func (__obf_6ccac4bb8646a806 *MapCache) Replace(__obf_afff83fa0e852ce0 string, __obf_3ab13e197ed6fe6b any, __obf_b39c53bb7d5c0767 time.Duration) error {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	_, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b66283acccf8c10b(__obf_afff83fa0e852ce0)
	if !__obf_69b7b10d07db29c0 {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_afff83fa0e852ce0)
	}
	__obf_6ccac4bb8646a806.__obf_930bee944f85c808(__obf_afff83fa0e852ce0, __obf_3ab13e197ed6fe6b,

		// Get an item from the cache. Returns the item or nil, and a bool indicating
		// whether the key was found.
		__obf_b39c53bb7d5c0767)
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Unlock()
	return nil
}

func (__obf_6ccac4bb8646a806 *MapCache) Get(__obf_afff83fa0e852ce0 string) (any, bool) {
	__obf_6ccac4bb8646a806.

		// "Inlining" of get and Expired
		__obf_2aacabd3f44c4b79.
		RLock()
	__obf_54b1e5958ef673f0, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			RUnlock()
		return nil, false
	}
	if __obf_54b1e5958ef673f0.Expiration > 0 {
		if time.Now().UnixNano() > __obf_54b1e5958ef673f0.Expiration {
			__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
				RUnlock()
			return nil, false
		}
	}
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		RUnlock()
	return __obf_54b1e5958ef673f0.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_6ccac4bb8646a806 *MapCache) GetWithExpiration(__obf_afff83fa0e852ce0 string) (any, time.Time, bool) {
	__obf_6ccac4bb8646a806.

		// "Inlining" of get and Expired
		__obf_2aacabd3f44c4b79.
		RLock()
	__obf_54b1e5958ef673f0, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_54b1e5958ef673f0.Expiration > 0 {
		if time.Now().UnixNano() > __obf_54b1e5958ef673f0.Expiration {
			__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
				RUnlock()
			return nil, time.Time{}, false
		}
		__obf_6ccac4bb8646a806.

			// Return the item and the expiration time
			__obf_2aacabd3f44c4b79.
			RUnlock()
		return __obf_54b1e5958ef673f0.Object, time.Unix(0, __obf_54b1e5958ef673f0.Expiration), true
	}
	__obf_6ccac4bb8646a806.

		// If expiration <= 0 (i.e. no expiration time set) then return the item
		// and a zeroed time.Time
		__obf_2aacabd3f44c4b79.
		RUnlock()
	return __obf_54b1e5958ef673f0.Object, time.Time{}, true
}

func (__obf_6ccac4bb8646a806 *MapCache) __obf_b66283acccf8c10b(__obf_afff83fa0e852ce0 string) (any, bool) {
	__obf_54b1e5958ef673f0, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_54b1e5958ef673f0.Expiration > 0 {
		if time.Now().UnixNano() > __obf_54b1e5958ef673f0.Expiration {
			return nil, false
		}
	}
	return __obf_54b1e5958ef673f0.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_6ccac4bb8646a806 *MapCache) Increment(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int64) error {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	switch __obf_6208501794a51081.Object.(type) {
	case int:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(int) + int(__obf_96fe4f800f83b51e)
	case int8:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(int8) + int8(__obf_96fe4f800f83b51e)
	case int16:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(int16) + int16(__obf_96fe4f800f83b51e)
	case int32:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(int32) + int32(__obf_96fe4f800f83b51e)
	case int64:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(int64) + __obf_96fe4f800f83b51e
	case uint:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uint) + uint(__obf_96fe4f800f83b51e)
	case uintptr:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uintptr) + uintptr(__obf_96fe4f800f83b51e)
	case uint8:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uint8) + uint8(__obf_96fe4f800f83b51e)
	case uint16:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uint16) + uint16(__obf_96fe4f800f83b51e)
	case uint32:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uint32) + uint32(__obf_96fe4f800f83b51e)
	case uint64:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uint64) + uint64(__obf_96fe4f800f83b51e)
	case float32:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(float32) + float32(__obf_96fe4f800f83b51e)
	case float64:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(float64) + float64(__obf_96fe4f800f83b51e)
	default:
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_afff83fa0e852ce0)
	}
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to increment it by n. Pass a negative number to decrement the
		// value. To retrieve the incremented value, use one of the specialized methods,
		// e.g. IncrementFloat64.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementFloat(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e float64) error {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	switch __obf_6208501794a51081.Object.(type) {
	case float32:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(float32) + float32(__obf_96fe4f800f83b51e)
	case float64:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(float64) + __obf_96fe4f800f83b51e
	default:
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_afff83fa0e852ce0)
	}
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementInt(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int) (int, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(int)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementInt8(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int8) (int8, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(int8)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementInt16(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int16) (int16, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(int16)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementInt32(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int32) (int32, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(int32)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementInt64(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int64) (int64, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(int64)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementUint(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uint) (uint, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uint)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementUintptr(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uintptr) (uintptr, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uintptr)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type uint8 by n. Returns an error if the item's value
		// is not an uint8, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementUint8(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uint8) (uint8, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uint8)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementUint16(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uint16) (uint16, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uint16)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementUint32(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uint32) (uint32, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uint32)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementUint64(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uint64) (uint64, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uint64)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementFloat32(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e float32) (float32, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(float32)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Increment an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) IncrementFloat64(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e float64) (float64, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(float64)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 + __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
		// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
		// item's value is not an integer, if it was not found, or if it is not
		// possible to decrement it by n. To retrieve the decremented value, use one
		// of the specialized methods, e.g. DecrementInt64.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) Decrement(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int64) error {
	__obf_6ccac4bb8646a806.
		// TODO: Implement Increment and Decrement more cleanly.
		// (Cannot do Increment(k, n*-1) for uints.)
		__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_6208501794a51081.Object.(type) {
	case int:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(int) - int(__obf_96fe4f800f83b51e)
	case int8:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(int8) - int8(__obf_96fe4f800f83b51e)
	case int16:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(int16) - int16(__obf_96fe4f800f83b51e)
	case int32:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(int32) - int32(__obf_96fe4f800f83b51e)
	case int64:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(int64) - __obf_96fe4f800f83b51e
	case uint:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uint) - uint(__obf_96fe4f800f83b51e)
	case uintptr:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uintptr) - uintptr(__obf_96fe4f800f83b51e)
	case uint8:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uint8) - uint8(__obf_96fe4f800f83b51e)
	case uint16:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uint16) - uint16(__obf_96fe4f800f83b51e)
	case uint32:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uint32) - uint32(__obf_96fe4f800f83b51e)
	case uint64:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(uint64) - uint64(__obf_96fe4f800f83b51e)
	case float32:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(float32) - float32(__obf_96fe4f800f83b51e)
	case float64:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(float64) - float64(__obf_96fe4f800f83b51e)
	default:
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_afff83fa0e852ce0)
	}
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to decrement it by n. Pass a negative number to decrement the
		// value. To retrieve the decremented value, use one of the specialized methods,
		// e.g. DecrementFloat64.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementFloat(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e float64) error {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	switch __obf_6208501794a51081.Object.(type) {
	case float32:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(float32) - float32(__obf_96fe4f800f83b51e)
	case float64:
		__obf_6208501794a51081.
			Object = __obf_6208501794a51081.Object.(float64) - __obf_96fe4f800f83b51e
	default:
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_afff83fa0e852ce0)
	}
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementInt(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int) (int, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(int)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementInt8(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int8) (int8, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(int8)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementInt16(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int16) (int16, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(int16)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementInt32(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int32) (int32, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(int32)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementInt64(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e int64) (int64, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(int64)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementUint(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uint) (uint, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uint)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementUintptr(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uintptr) (uintptr, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uintptr)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type uint8 by n. Returns an error if the item's value is
		// not an uint8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementUint8(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uint8) (uint8, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uint8)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementUint16(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uint16) (uint16, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uint16)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementUint32(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uint32) (uint32, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uint32)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementUint64(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e uint64) (uint64, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(uint64)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementFloat32(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e float32) (float32, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(float32)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Decrement an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) DecrementFloat64(__obf_afff83fa0e852ce0 string, __obf_96fe4f800f83b51e float64) (float64, error) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
	if !__obf_69b7b10d07db29c0 || __obf_6208501794a51081.Expired() {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_afff83fa0e852ce0)
	}
	__obf_65ea976e127830d6, __obf_8761a3e94b6603ac := __obf_6208501794a51081.Object.(float64)
	if !__obf_8761a3e94b6603ac {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_afff83fa0e852ce0)
	}
	__obf_6499bdb79d3727ab := __obf_65ea976e127830d6 - __obf_96fe4f800f83b51e
	__obf_6208501794a51081.
		Object = __obf_6499bdb79d3727ab
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	__obf_6ccac4bb8646a806.

		// Delete an item from the cache. Does nothing if the key is not in the cache.
		__obf_2aacabd3f44c4b79.
		Unlock()
	return __obf_6499bdb79d3727ab, nil
}

func (__obf_6ccac4bb8646a806 *MapCache) Delete(__obf_afff83fa0e852ce0 string) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6208501794a51081, __obf_dfdbdbd83228e3a1 := __obf_6ccac4bb8646a806.delete(__obf_afff83fa0e852ce0)
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Unlock()
	if __obf_dfdbdbd83228e3a1 {
		__obf_6ccac4bb8646a806.__obf_8c76d4f25d47695a(__obf_afff83fa0e852ce0, __obf_6208501794a51081)
	}
}

func (__obf_6ccac4bb8646a806 *MapCache) delete(__obf_afff83fa0e852ce0 string) (any, bool) {
	if __obf_6ccac4bb8646a806.__obf_8c76d4f25d47695a != nil {
		if __obf_6208501794a51081, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]; __obf_69b7b10d07db29c0 {
			delete(__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0, __obf_afff83fa0e852ce0)
			return __obf_6208501794a51081.Object, true
		}
	}
	delete(__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0, __obf_afff83fa0e852ce0)
	return nil, false
}

type __obf_2fb0503bf5ec10b9 struct {
	__obf_fa3a380c35ada5d9 string
	__obf_c2b0e78296a7a084 any
}

// Delete all expired items from the cache.
func (__obf_6ccac4bb8646a806 *MapCache) DeleteExpired() {
	var __obf_05abb24f81ed38de []__obf_2fb0503bf5ec10b9
	__obf_0ee6663e8d6c2336 := time.Now().UnixNano()
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	for __obf_afff83fa0e852ce0, __obf_6208501794a51081 := range __obf_6ccac4bb8646a806.
		// "Inlining" of expired
		__obf_b4d587b1629307c0 {

		if __obf_6208501794a51081.Expiration > 0 && __obf_0ee6663e8d6c2336 > __obf_6208501794a51081.Expiration {
			__obf_be8218dc64ff1212, __obf_dfdbdbd83228e3a1 := __obf_6ccac4bb8646a806.delete(__obf_afff83fa0e852ce0)
			if __obf_dfdbdbd83228e3a1 {
				__obf_05abb24f81ed38de = append(__obf_05abb24f81ed38de, __obf_2fb0503bf5ec10b9{__obf_afff83fa0e852ce0, __obf_be8218dc64ff1212})
			}
		}
	}
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Unlock()
	for _, __obf_6208501794a51081 := range __obf_05abb24f81ed38de {
		__obf_6ccac4bb8646a806.__obf_8c76d4f25d47695a(__obf_6208501794a51081.__obf_fa3a380c35ada5d9,

			// Sets an (optional) function that is called with the key and value when an
			// item is evicted from the cache. (Including when it is deleted manually, but
			// not when it is overwritten.) Set to nil to disable.
			__obf_6208501794a51081.__obf_c2b0e78296a7a084)
	}
}

func (__obf_6ccac4bb8646a806 *MapCache) OnEvicted(__obf_7a0c782b6fdea76a func(string, any)) {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6ccac4bb8646a806.__obf_8c76d4f25d47695a = __obf_7a0c782b6fdea76a
	__obf_6ccac4bb8646a806.

		// Write the cache's items (using Gob) to an io.Writer.
		//
		// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
		// documentation for NewFrom().)
		__obf_2aacabd3f44c4b79.
		Unlock()
}

func (__obf_6ccac4bb8646a806 *MapCache) Save(__obf_9d7803536a66fb33 io.Writer) (__obf_9ffdd5bbb9f1dc61 error) {
	__obf_95030e14a18986fb := gob.NewEncoder(__obf_9d7803536a66fb33)
	defer func() {
		if __obf_3ab13e197ed6fe6b := recover(); __obf_3ab13e197ed6fe6b != nil {
			__obf_9ffdd5bbb9f1dc61 = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		RLock()
	defer __obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.RUnlock()
	for _, __obf_6208501794a51081 := range __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0 {
		gob.Register(__obf_6208501794a51081.Object)
	}
	__obf_9ffdd5bbb9f1dc61 = __obf_95030e14a18986fb.Encode(&__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_6ccac4bb8646a806 *MapCache) SaveFile(__obf_ad8db1cf4fdf963b string) error {
	__obf_21ba25ffd5ba2109, __obf_9ffdd5bbb9f1dc61 := os.Create(__obf_ad8db1cf4fdf963b)
	if __obf_9ffdd5bbb9f1dc61 != nil {
		return __obf_9ffdd5bbb9f1dc61
	}
	__obf_9ffdd5bbb9f1dc61 = __obf_6ccac4bb8646a806.Save(__obf_21ba25ffd5ba2109)
	if __obf_9ffdd5bbb9f1dc61 != nil {
		__obf_21ba25ffd5ba2109.
			Close()
		return __obf_9ffdd5bbb9f1dc61
	}
	return __obf_21ba25ffd5ba2109.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_6ccac4bb8646a806 *MapCache) Load(__obf_3a5bbe1f7e27c46d io.Reader) error {
	__obf_177b67818a43cefb := gob.NewDecoder(__obf_3a5bbe1f7e27c46d)
	__obf_b4d587b1629307c0 := map[string]Item{}
	__obf_9ffdd5bbb9f1dc61 := __obf_177b67818a43cefb.Decode(&__obf_b4d587b1629307c0)
	if __obf_9ffdd5bbb9f1dc61 == nil {
		__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
			Lock()
		defer __obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.Unlock()
		for __obf_afff83fa0e852ce0, __obf_6208501794a51081 := range __obf_b4d587b1629307c0 {
			__obf_be8218dc64ff1212, __obf_69b7b10d07db29c0 := __obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0]
			if !__obf_69b7b10d07db29c0 || __obf_be8218dc64ff1212.Expired() {
				__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
			}
		}
	}
	return __obf_9ffdd5bbb9f1dc61
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_6ccac4bb8646a806 *MapCache) LoadFile(__obf_ad8db1cf4fdf963b string) error {
	__obf_21ba25ffd5ba2109, __obf_9ffdd5bbb9f1dc61 := os.Open(__obf_ad8db1cf4fdf963b)
	if __obf_9ffdd5bbb9f1dc61 != nil {
		return __obf_9ffdd5bbb9f1dc61
	}
	__obf_9ffdd5bbb9f1dc61 = __obf_6ccac4bb8646a806.Load(__obf_21ba25ffd5ba2109)
	if __obf_9ffdd5bbb9f1dc61 != nil {
		__obf_21ba25ffd5ba2109.
			Close()
		return __obf_9ffdd5bbb9f1dc61
	}
	return __obf_21ba25ffd5ba2109.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_6ccac4bb8646a806 *MapCache) Items() map[string]Item {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		RLock()
	defer __obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.RUnlock()
	__obf_2a1b11e5bfb41cce := make(map[string]Item, len(__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0))
	__obf_0ee6663e8d6c2336 := time.Now().UnixNano()
	for __obf_afff83fa0e852ce0, __obf_6208501794a51081 := range __obf_6ccac4bb8646a806.
		// "Inlining" of Expired
		__obf_b4d587b1629307c0 {

		if __obf_6208501794a51081.Expiration > 0 {
			if __obf_0ee6663e8d6c2336 > __obf_6208501794a51081.Expiration {
				continue
			}
		}
		__obf_2a1b11e5bfb41cce[__obf_afff83fa0e852ce0] = __obf_6208501794a51081
	}
	return __obf_2a1b11e5bfb41cce
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_6ccac4bb8646a806 *MapCache) ItemCount() int {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		RLock()
	__obf_96fe4f800f83b51e := len(__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0)
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		RUnlock()
	return __obf_96fe4f800f83b51e
}

// Delete all items from the cache.
func (__obf_6ccac4bb8646a806 *MapCache) Flush() {
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Lock()
	__obf_6ccac4bb8646a806.__obf_b4d587b1629307c0 = map[string]Item{}
	__obf_6ccac4bb8646a806.__obf_2aacabd3f44c4b79.
		Unlock()
}

type __obf_cc3d7a1795ac42ab struct {
	Interval               time.Duration
	__obf_8f1e925268e30aa1 chan bool
}

func (__obf_09cdea3fba79bc71 *__obf_cc3d7a1795ac42ab) Run(__obf_6ccac4bb8646a806 *MapCache) {
	__obf_8a7409aabc601b0b := time.NewTicker(__obf_09cdea3fba79bc71.Interval)
	for {
		select {
		case <-__obf_8a7409aabc601b0b.C:
			__obf_6ccac4bb8646a806.
				DeleteExpired()
		case <-__obf_09cdea3fba79bc71.__obf_8f1e925268e30aa1:
			__obf_8a7409aabc601b0b.
				Stop()
			return
		}
	}
}

func __obf_a3e6a008a6623414(__obf_6ccac4bb8646a806 *MapCache) {
	__obf_6ccac4bb8646a806.__obf_cc3d7a1795ac42ab.__obf_8f1e925268e30aa1 <- true
}

func __obf_8095453f70d2e66e(__obf_6ccac4bb8646a806 *MapCache, __obf_cc6d28d8792da3c1 time.Duration) {
	__obf_09cdea3fba79bc71 := &__obf_cc3d7a1795ac42ab{
		Interval: __obf_cc6d28d8792da3c1, __obf_8f1e925268e30aa1: make(chan bool),
	}
	__obf_6ccac4bb8646a806.__obf_cc3d7a1795ac42ab = __obf_09cdea3fba79bc71
	go __obf_09cdea3fba79bc71.Run(__obf_6ccac4bb8646a806)
}

func __obf_f89c6c96404c32ad(__obf_6aaa919faad23937 time.Duration, __obf_2a1b11e5bfb41cce map[string]Item) *MapCache {
	if __obf_6aaa919faad23937 == 0 {
		__obf_6aaa919faad23937 = -1
	}
	__obf_6ccac4bb8646a806 := &MapCache{__obf_efa9e45df3c18f44: __obf_6aaa919faad23937, __obf_b4d587b1629307c0: __obf_2a1b11e5bfb41cce}
	return __obf_6ccac4bb8646a806
}

func __obf_5998b9814aa3c88a(__obf_6aaa919faad23937 time.Duration, __obf_cc6d28d8792da3c1 time.Duration, __obf_2a1b11e5bfb41cce map[string]Item) *MapCache {
	__obf_6ccac4bb8646a806 := __obf_f89c6c96404c32ad(__obf_6aaa919faad23937,
		// This trick ensures that the janitor goroutine (which--granted it
		// was enabled--is running DeleteExpired on c forever) does not keep
		// the returned C object from being garbage collected. When it is
		// garbage collected, the finalizer stops the janitor goroutine, after
		// which c can be collected.
		// C := &MapCache{c}
		__obf_2a1b11e5bfb41cce)

	if __obf_cc6d28d8792da3c1 > 0 {
		__obf_8095453f70d2e66e(__obf_6ccac4bb8646a806, __obf_cc6d28d8792da3c1)
		runtime.SetFinalizer(__obf_6ccac4bb8646a806, __obf_a3e6a008a6623414)
	}
	return __obf_6ccac4bb8646a806
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_efa9e45df3c18f44, __obf_f5dc766af3690723 time.Duration) *MapCache {
	__obf_b4d587b1629307c0 := make(map[string]Item)
	return __obf_5998b9814aa3c88a(__obf_efa9e45df3c18f44, __obf_f5dc766af3690723, __obf_b4d587b1629307c0)
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
func NewFrom(__obf_efa9e45df3c18f44, __obf_f5dc766af3690723 time.Duration, __obf_b4d587b1629307c0 map[string]Item) *MapCache {
	return __obf_5998b9814aa3c88a(__obf_efa9e45df3c18f44, __obf_f5dc766af3690723, __obf_b4d587b1629307c0)
}
