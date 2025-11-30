package __obf_9e1ee87c6b054458

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
func (__obf_755d49269b07cedb Item) Expired() bool {
	if __obf_755d49269b07cedb.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > __obf_755d49269b07cedb.Expiration
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
	__obf_87cea4da8de7826c time. // *cache
		// If this is confusing, see the comment at the bottom of New()
		Duration
	__obf_0d00aca7befe11fb map[string]Item
	__obf_26007347eed65c2a sync.RWMutex
	__obf_531f779b45352a46 func(string, any)
	__obf_47c6aae2e51f622e *__obf_47c6aae2e51f622e
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
func (__obf_8af2676446f64c49 *MapCache) Set(__obf_ffc285490d834437 string, __obf_75038cdd74d6258f any, __obf_d094b5585f6de199 time.Duration) {
	// "Inlining" of set
	var __obf_4f22c5826c5e34df int64
	if __obf_d094b5585f6de199 == DefaultExpiration {
		__obf_d094b5585f6de199 = __obf_8af2676446f64c49.__obf_87cea4da8de7826c
	}
	if __obf_d094b5585f6de199 > 0 {
		__obf_4f22c5826c5e34df = time.Now().Add(__obf_d094b5585f6de199).UnixNano()
	}
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = Item{
		Object:     __obf_75038cdd74d6258f,
		Expiration: __obf_4f22c5826c5e34df,
	}
	__obf_8af2676446f64c49.
		// TODO: Calls to mu.Unlock are currently not deferred because defer
		// adds ~200 ns (as of go1.)
		__obf_26007347eed65c2a.
		Unlock()
}

func (__obf_8af2676446f64c49 *MapCache) __obf_c0e0c1dea0ef9c51(__obf_ffc285490d834437 string, __obf_75038cdd74d6258f any, __obf_d094b5585f6de199 time.Duration) {
	var __obf_4f22c5826c5e34df int64
	if __obf_d094b5585f6de199 == DefaultExpiration {
		__obf_d094b5585f6de199 = __obf_8af2676446f64c49.__obf_87cea4da8de7826c
	}
	if __obf_d094b5585f6de199 > 0 {
		__obf_4f22c5826c5e34df = time.Now().Add(__obf_d094b5585f6de199).UnixNano()
	}
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = Item{
		Object:     __obf_75038cdd74d6258f,
		Expiration: __obf_4f22c5826c5e34df,
	}
}

// Add an item to the cache, replacing any existing item, using the default
// expiration.
func (__obf_8af2676446f64c49 *MapCache) SetDefault(__obf_ffc285490d834437 string, __obf_75038cdd74d6258f any) {
	__obf_8af2676446f64c49.
		Set(__obf_ffc285490d834437, __obf_75038cdd74d6258f, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (__obf_8af2676446f64c49 *MapCache) Add(__obf_ffc285490d834437 string, __obf_75038cdd74d6258f any, __obf_d094b5585f6de199 time.Duration) error {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	_, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_e32536d1802b806b(__obf_ffc285490d834437)
	if __obf_e527ca4e2a4ed527 {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return fmt.Errorf("Item %s already exists", __obf_ffc285490d834437)
	}
	__obf_8af2676446f64c49.__obf_c0e0c1dea0ef9c51(__obf_ffc285490d834437, __obf_75038cdd74d6258f,

		// Set a new value for the cache key only if it already exists, and the existing
		// item hasn't expired. Returns an error otherwise.
		__obf_d094b5585f6de199)
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Unlock()
	return nil
}

func (__obf_8af2676446f64c49 *MapCache) Replace(__obf_ffc285490d834437 string, __obf_75038cdd74d6258f any, __obf_d094b5585f6de199 time.Duration) error {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	_, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_e32536d1802b806b(__obf_ffc285490d834437)
	if !__obf_e527ca4e2a4ed527 {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return fmt.Errorf("Item %s doesn't exist", __obf_ffc285490d834437)
	}
	__obf_8af2676446f64c49.__obf_c0e0c1dea0ef9c51(__obf_ffc285490d834437, __obf_75038cdd74d6258f,

		// Get an item from the cache. Returns the item or nil, and a bool indicating
		// whether the key was found.
		__obf_d094b5585f6de199)
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Unlock()
	return nil
}

func (__obf_8af2676446f64c49 *MapCache) Get(__obf_ffc285490d834437 string) (any, bool) {
	__obf_8af2676446f64c49.

		// "Inlining" of get and Expired
		__obf_26007347eed65c2a.
		RLock()
	__obf_755d49269b07cedb, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			RUnlock()
		return nil, false
	}
	if __obf_755d49269b07cedb.Expiration > 0 {
		if time.Now().UnixNano() > __obf_755d49269b07cedb.Expiration {
			__obf_8af2676446f64c49.__obf_26007347eed65c2a.
				RUnlock()
			return nil, false
		}
	}
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		RUnlock()
	return __obf_755d49269b07cedb.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (__obf_8af2676446f64c49 *MapCache) GetWithExpiration(__obf_ffc285490d834437 string) (any, time.Time, bool) {
	__obf_8af2676446f64c49.

		// "Inlining" of get and Expired
		__obf_26007347eed65c2a.
		RLock()
	__obf_755d49269b07cedb, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			RUnlock()
		return nil, time.Time{}, false
	}

	if __obf_755d49269b07cedb.Expiration > 0 {
		if time.Now().UnixNano() > __obf_755d49269b07cedb.Expiration {
			__obf_8af2676446f64c49.__obf_26007347eed65c2a.
				RUnlock()
			return nil, time.Time{}, false
		}
		__obf_8af2676446f64c49.

			// Return the item and the expiration time
			__obf_26007347eed65c2a.
			RUnlock()
		return __obf_755d49269b07cedb.Object, time.Unix(0, __obf_755d49269b07cedb.Expiration), true
	}
	__obf_8af2676446f64c49.

		// If expiration <= 0 (i.e. no expiration time set) then return the item
		// and a zeroed time.Time
		__obf_26007347eed65c2a.
		RUnlock()
	return __obf_755d49269b07cedb.Object, time.Time{}, true
}

func (__obf_8af2676446f64c49 *MapCache) __obf_e32536d1802b806b(__obf_ffc285490d834437 string) (any, bool) {
	__obf_755d49269b07cedb, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 {
		return nil, false
	}
	// "Inlining" of Expired
	if __obf_755d49269b07cedb.Expiration > 0 {
		if time.Now().UnixNano() > __obf_755d49269b07cedb.Expiration {
			return nil, false
		}
	}
	return __obf_755d49269b07cedb.Object, true
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (__obf_8af2676446f64c49 *MapCache) Increment(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int64) error {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	switch __obf_aa24543fddec7b80.Object.(type) {
	case int:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(int) + int(__obf_34e919d33567e903)
	case int8:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(int8) + int8(__obf_34e919d33567e903)
	case int16:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(int16) + int16(__obf_34e919d33567e903)
	case int32:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(int32) + int32(__obf_34e919d33567e903)
	case int64:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(int64) + __obf_34e919d33567e903
	case uint:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uint) + uint(__obf_34e919d33567e903)
	case uintptr:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uintptr) + uintptr(__obf_34e919d33567e903)
	case uint8:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uint8) + uint8(__obf_34e919d33567e903)
	case uint16:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uint16) + uint16(__obf_34e919d33567e903)
	case uint32:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uint32) + uint32(__obf_34e919d33567e903)
	case uint64:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uint64) + uint64(__obf_34e919d33567e903)
	case float32:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(float32) + float32(__obf_34e919d33567e903)
	case float64:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(float64) + float64(__obf_34e919d33567e903)
	default:
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_ffc285490d834437)
	}
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to increment it by n. Pass a negative number to decrement the
		// value. To retrieve the incremented value, use one of the specialized methods,
		// e.g. IncrementFloat64.
		__obf_26007347eed65c2a.
		Unlock()
	return nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementFloat(__obf_ffc285490d834437 string, __obf_34e919d33567e903 float64) error {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	switch __obf_aa24543fddec7b80.Object.(type) {
	case float32:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(float32) + float32(__obf_34e919d33567e903)
	case float64:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(float64) + __obf_34e919d33567e903
	default:
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_ffc285490d834437)
	}
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementInt(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int) (int, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(int)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementInt8(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int8) (int8, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(int8)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementInt16(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int16) (int16, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(int16)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementInt32(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int32) (int32, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(int32)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementInt64(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int64) (int64, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(int64)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the incremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementUint(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uint) (uint, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uint)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementUintptr(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uintptr) (uintptr, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uintptr)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type uint8 by n. Returns an error if the item's value
		// is not an uint8, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementUint8(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uint8) (uint8, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uint8)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementUint16(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uint16) (uint16, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uint16)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementUint32(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uint32) (uint32, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uint32)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementUint64(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uint64) (uint64, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uint64)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementFloat32(__obf_ffc285490d834437 string, __obf_34e919d33567e903 float32) (float32, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(float32)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Increment an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// incremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) IncrementFloat64(__obf_ffc285490d834437 string, __obf_34e919d33567e903 float64) (float64, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(float64)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a + __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
		// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
		// item's value is not an integer, if it was not found, or if it is not
		// possible to decrement it by n. To retrieve the decremented value, use one
		// of the specialized methods, e.g. DecrementInt64.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) Decrement(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int64) error {
	__obf_8af2676446f64c49.
		// TODO: Implement Increment and Decrement more cleanly.
		// (Cannot do Increment(k, n*-1) for uints.)
		__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return fmt.Errorf("Item not found")
	}
	switch __obf_aa24543fddec7b80.Object.(type) {
	case int:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(int) - int(__obf_34e919d33567e903)
	case int8:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(int8) - int8(__obf_34e919d33567e903)
	case int16:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(int16) - int16(__obf_34e919d33567e903)
	case int32:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(int32) - int32(__obf_34e919d33567e903)
	case int64:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(int64) - __obf_34e919d33567e903
	case uint:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uint) - uint(__obf_34e919d33567e903)
	case uintptr:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uintptr) - uintptr(__obf_34e919d33567e903)
	case uint8:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uint8) - uint8(__obf_34e919d33567e903)
	case uint16:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uint16) - uint16(__obf_34e919d33567e903)
	case uint32:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uint32) - uint32(__obf_34e919d33567e903)
	case uint64:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(uint64) - uint64(__obf_34e919d33567e903)
	case float32:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(float32) - float32(__obf_34e919d33567e903)
	case float64:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(float64) - float64(__obf_34e919d33567e903)
	default:
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return fmt.Errorf("The value for %s is not an integer", __obf_ffc285490d834437)
	}
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type float32 or float64 by n. Returns an error if the
		// item's value is not floating point, if it was not found, or if it is not
		// possible to decrement it by n. Pass a negative number to decrement the
		// value. To retrieve the decremented value, use one of the specialized methods,
		// e.g. DecrementFloat64.
		__obf_26007347eed65c2a.
		Unlock()
	return nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementFloat(__obf_ffc285490d834437 string, __obf_34e919d33567e903 float64) error {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	switch __obf_aa24543fddec7b80.Object.(type) {
	case float32:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(float32) - float32(__obf_34e919d33567e903)
	case float64:
		__obf_aa24543fddec7b80.
			Object = __obf_aa24543fddec7b80.Object.(float64) - __obf_34e919d33567e903
	default:
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return fmt.Errorf("The value for %s does not have type float32 or float64", __obf_ffc285490d834437)
	}
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type int by n. Returns an error if the item's value is
		// not an int, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementInt(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int) (int, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(int)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type int8 by n. Returns an error if the item's value is
		// not an int8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementInt8(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int8) (int8, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(int8)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int8", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type int16 by n. Returns an error if the item's value is
		// not an int16, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementInt16(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int16) (int16, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(int16)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int16", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type int32 by n. Returns an error if the item's value is
		// not an int32, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementInt32(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int32) (int32, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(int32)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int32", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type int64 by n. Returns an error if the item's value is
		// not an int64, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementInt64(__obf_ffc285490d834437 string, __obf_34e919d33567e903 int64) (int64, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(int64)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an int64", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type uint by n. Returns an error if the item's value is
		// not an uint, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementUint(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uint) (uint, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uint)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type uintptr by n. Returns an error if the item's value
		// is not an uintptr, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementUintptr(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uintptr) (uintptr, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uintptr)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uintptr", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type uint8 by n. Returns an error if the item's value is
		// not an uint8, or if it was not found. If there is no error, the decremented
		// value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementUint8(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uint8) (uint8, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uint8)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint8", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type uint16 by n. Returns an error if the item's value
		// is not an uint16, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementUint16(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uint16) (uint16, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uint16)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint16", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type uint32 by n. Returns an error if the item's value
		// is not an uint32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementUint32(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uint32) (uint32, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uint32)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint32", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type uint64 by n. Returns an error if the item's value
		// is not an uint64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementUint64(__obf_ffc285490d834437 string, __obf_34e919d33567e903 uint64) (uint64, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(uint64)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an uint64", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type float32 by n. Returns an error if the item's value
		// is not an float32, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementFloat32(__obf_ffc285490d834437 string, __obf_34e919d33567e903 float32) (float32, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(float32)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float32", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Decrement an item of type float64 by n. Returns an error if the item's value
		// is not an float64, or if it was not found. If there is no error, the
		// decremented value is returned.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) DecrementFloat64(__obf_ffc285490d834437 string, __obf_34e919d33567e903 float64) (float64, error) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
	if !__obf_e527ca4e2a4ed527 || __obf_aa24543fddec7b80.Expired() {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("Item %s not found", __obf_ffc285490d834437)
	}
	__obf_2892b89e1ae9974a, __obf_40eed10588aa6cec := __obf_aa24543fddec7b80.Object.(float64)
	if !__obf_40eed10588aa6cec {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Unlock()
		return 0, fmt.Errorf("The value for %s is not an float64", __obf_ffc285490d834437)
	}
	__obf_b654d85a275e5c2b := __obf_2892b89e1ae9974a - __obf_34e919d33567e903
	__obf_aa24543fddec7b80.
		Object = __obf_b654d85a275e5c2b
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	__obf_8af2676446f64c49.

		// Delete an item from the cache. Does nothing if the key is not in the cache.
		__obf_26007347eed65c2a.
		Unlock()
	return __obf_b654d85a275e5c2b, nil
}

func (__obf_8af2676446f64c49 *MapCache) Delete(__obf_ffc285490d834437 string) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_aa24543fddec7b80, __obf_90748537928ce52a := __obf_8af2676446f64c49.delete(__obf_ffc285490d834437)
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Unlock()
	if __obf_90748537928ce52a {
		__obf_8af2676446f64c49.__obf_531f779b45352a46(__obf_ffc285490d834437, __obf_aa24543fddec7b80)
	}
}

func (__obf_8af2676446f64c49 *MapCache) delete(__obf_ffc285490d834437 string) (any, bool) {
	if __obf_8af2676446f64c49.__obf_531f779b45352a46 != nil {
		if __obf_aa24543fddec7b80, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]; __obf_e527ca4e2a4ed527 {
			delete(__obf_8af2676446f64c49.__obf_0d00aca7befe11fb, __obf_ffc285490d834437)
			return __obf_aa24543fddec7b80.Object, true
		}
	}
	delete(__obf_8af2676446f64c49.__obf_0d00aca7befe11fb, __obf_ffc285490d834437)
	return nil, false
}

type __obf_c9b6f8337598824b struct {
	__obf_3c13197612c6b39f string
	__obf_2d0927453cc08d1b any
}

// Delete all expired items from the cache.
func (__obf_8af2676446f64c49 *MapCache) DeleteExpired() {
	var __obf_5ff94b00b12d9bf8 []__obf_c9b6f8337598824b
	__obf_f7969abcc9371dff := time.Now().UnixNano()
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	for __obf_ffc285490d834437, __obf_aa24543fddec7b80 := range __obf_8af2676446f64c49.
		// "Inlining" of expired
		__obf_0d00aca7befe11fb {

		if __obf_aa24543fddec7b80.Expiration > 0 && __obf_f7969abcc9371dff > __obf_aa24543fddec7b80.Expiration {
			__obf_68cec354e917bec4, __obf_90748537928ce52a := __obf_8af2676446f64c49.delete(__obf_ffc285490d834437)
			if __obf_90748537928ce52a {
				__obf_5ff94b00b12d9bf8 = append(__obf_5ff94b00b12d9bf8, __obf_c9b6f8337598824b{__obf_ffc285490d834437, __obf_68cec354e917bec4})
			}
		}
	}
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Unlock()
	for _, __obf_aa24543fddec7b80 := range __obf_5ff94b00b12d9bf8 {
		__obf_8af2676446f64c49.__obf_531f779b45352a46(__obf_aa24543fddec7b80.__obf_3c13197612c6b39f,

			// Sets an (optional) function that is called with the key and value when an
			// item is evicted from the cache. (Including when it is deleted manually, but
			// not when it is overwritten.) Set to nil to disable.
			__obf_aa24543fddec7b80.__obf_2d0927453cc08d1b)
	}
}

func (__obf_8af2676446f64c49 *MapCache) OnEvicted(__obf_961fc3b1f2fdc2d4 func(string, any)) {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_8af2676446f64c49.__obf_531f779b45352a46 = __obf_961fc3b1f2fdc2d4
	__obf_8af2676446f64c49.

		// Write the cache's items (using Gob) to an io.Writer.
		//
		// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
		// documentation for NewFrom().)
		__obf_26007347eed65c2a.
		Unlock()
}

func (__obf_8af2676446f64c49 *MapCache) Save(__obf_539ae185966de514 io.Writer) (__obf_498673050542660c error) {
	__obf_3695aef862da81d9 := gob.NewEncoder(__obf_539ae185966de514)
	defer func() {
		if __obf_75038cdd74d6258f := recover(); __obf_75038cdd74d6258f != nil {
			__obf_498673050542660c = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		RLock()
	defer __obf_8af2676446f64c49.__obf_26007347eed65c2a.RUnlock()
	for _, __obf_aa24543fddec7b80 := range __obf_8af2676446f64c49.__obf_0d00aca7befe11fb {
		gob.Register(__obf_aa24543fddec7b80.Object)
	}
	__obf_498673050542660c = __obf_3695aef862da81d9.Encode(&__obf_8af2676446f64c49.__obf_0d00aca7befe11fb)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_8af2676446f64c49 *MapCache) SaveFile(__obf_5d78d0dfb2fbd637 string) error {
	__obf_0fe34552f18261af, __obf_498673050542660c := os.Create(__obf_5d78d0dfb2fbd637)
	if __obf_498673050542660c != nil {
		return __obf_498673050542660c
	}
	__obf_498673050542660c = __obf_8af2676446f64c49.Save(__obf_0fe34552f18261af)
	if __obf_498673050542660c != nil {
		__obf_0fe34552f18261af.
			Close()
		return __obf_498673050542660c
	}
	return __obf_0fe34552f18261af.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_8af2676446f64c49 *MapCache) Load(__obf_171f9f68d8286408 io.Reader) error {
	__obf_1ea1f1ed012f482d := gob.NewDecoder(__obf_171f9f68d8286408)
	__obf_0d00aca7befe11fb := map[string]Item{}
	__obf_498673050542660c := __obf_1ea1f1ed012f482d.Decode(&__obf_0d00aca7befe11fb)
	if __obf_498673050542660c == nil {
		__obf_8af2676446f64c49.__obf_26007347eed65c2a.
			Lock()
		defer __obf_8af2676446f64c49.__obf_26007347eed65c2a.Unlock()
		for __obf_ffc285490d834437, __obf_aa24543fddec7b80 := range __obf_0d00aca7befe11fb {
			__obf_68cec354e917bec4, __obf_e527ca4e2a4ed527 := __obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437]
			if !__obf_e527ca4e2a4ed527 || __obf_68cec354e917bec4.Expired() {
				__obf_8af2676446f64c49.__obf_0d00aca7befe11fb[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
			}
		}
	}
	return __obf_498673050542660c
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (__obf_8af2676446f64c49 *MapCache) LoadFile(__obf_5d78d0dfb2fbd637 string) error {
	__obf_0fe34552f18261af, __obf_498673050542660c := os.Open(__obf_5d78d0dfb2fbd637)
	if __obf_498673050542660c != nil {
		return __obf_498673050542660c
	}
	__obf_498673050542660c = __obf_8af2676446f64c49.Load(__obf_0fe34552f18261af)
	if __obf_498673050542660c != nil {
		__obf_0fe34552f18261af.
			Close()
		return __obf_498673050542660c
	}
	return __obf_0fe34552f18261af.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (__obf_8af2676446f64c49 *MapCache) Items() map[string]Item {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		RLock()
	defer __obf_8af2676446f64c49.__obf_26007347eed65c2a.RUnlock()
	__obf_2c6bba2eafff0a17 := make(map[string]Item, len(__obf_8af2676446f64c49.__obf_0d00aca7befe11fb))
	__obf_f7969abcc9371dff := time.Now().UnixNano()
	for __obf_ffc285490d834437, __obf_aa24543fddec7b80 := range __obf_8af2676446f64c49.
		// "Inlining" of Expired
		__obf_0d00aca7befe11fb {

		if __obf_aa24543fddec7b80.Expiration > 0 {
			if __obf_f7969abcc9371dff > __obf_aa24543fddec7b80.Expiration {
				continue
			}
		}
		__obf_2c6bba2eafff0a17[__obf_ffc285490d834437] = __obf_aa24543fddec7b80
	}
	return __obf_2c6bba2eafff0a17
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (__obf_8af2676446f64c49 *MapCache) ItemCount() int {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		RLock()
	__obf_34e919d33567e903 := len(__obf_8af2676446f64c49.__obf_0d00aca7befe11fb)
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		RUnlock()
	return __obf_34e919d33567e903
}

// Delete all items from the cache.
func (__obf_8af2676446f64c49 *MapCache) Flush() {
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Lock()
	__obf_8af2676446f64c49.__obf_0d00aca7befe11fb = map[string]Item{}
	__obf_8af2676446f64c49.__obf_26007347eed65c2a.
		Unlock()
}

type __obf_47c6aae2e51f622e struct {
	Interval               time.Duration
	__obf_ece7b0f4ae4c1147 chan bool
}

func (__obf_5a930e2ba2f701ef *__obf_47c6aae2e51f622e) Run(__obf_8af2676446f64c49 *MapCache) {
	__obf_baa63daca99df408 := time.NewTicker(__obf_5a930e2ba2f701ef.Interval)
	for {
		select {
		case <-__obf_baa63daca99df408.C:
			__obf_8af2676446f64c49.
				DeleteExpired()
		case <-__obf_5a930e2ba2f701ef.__obf_ece7b0f4ae4c1147:
			__obf_baa63daca99df408.
				Stop()
			return
		}
	}
}

func __obf_d2d46ba2ec0cfa5b(__obf_8af2676446f64c49 *MapCache) {
	__obf_8af2676446f64c49.__obf_47c6aae2e51f622e.__obf_ece7b0f4ae4c1147 <- true
}

func __obf_4a4b54f55d28267b(__obf_8af2676446f64c49 *MapCache, __obf_13d41bb70f264951 time.Duration) {
	__obf_5a930e2ba2f701ef := &__obf_47c6aae2e51f622e{
		Interval: __obf_13d41bb70f264951, __obf_ece7b0f4ae4c1147: make(chan bool),
	}
	__obf_8af2676446f64c49.__obf_47c6aae2e51f622e = __obf_5a930e2ba2f701ef
	go __obf_5a930e2ba2f701ef.Run(__obf_8af2676446f64c49)
}

func __obf_26041b41ef3eb9df(__obf_5351fb68bfc558c5 time.Duration, __obf_2c6bba2eafff0a17 map[string]Item) *MapCache {
	if __obf_5351fb68bfc558c5 == 0 {
		__obf_5351fb68bfc558c5 = -1
	}
	__obf_8af2676446f64c49 := &MapCache{__obf_87cea4da8de7826c: __obf_5351fb68bfc558c5, __obf_0d00aca7befe11fb: __obf_2c6bba2eafff0a17}
	return __obf_8af2676446f64c49
}

func __obf_18fd55b4bac2f8b2(__obf_5351fb68bfc558c5 time.Duration, __obf_13d41bb70f264951 time.Duration, __obf_2c6bba2eafff0a17 map[string]Item) *MapCache {
	__obf_8af2676446f64c49 := __obf_26041b41ef3eb9df(__obf_5351fb68bfc558c5,
		// This trick ensures that the janitor goroutine (which--granted it
		// was enabled--is running DeleteExpired on c forever) does not keep
		// the returned C object from being garbage collected. When it is
		// garbage collected, the finalizer stops the janitor goroutine, after
		// which c can be collected.
		// C := &MapCache{c}
		__obf_2c6bba2eafff0a17)

	if __obf_13d41bb70f264951 > 0 {
		__obf_4a4b54f55d28267b(__obf_8af2676446f64c49, __obf_13d41bb70f264951)
		runtime.SetFinalizer(__obf_8af2676446f64c49, __obf_d2d46ba2ec0cfa5b)
	}
	return __obf_8af2676446f64c49
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(__obf_87cea4da8de7826c, __obf_715327e8eeb7b279 time.Duration) *MapCache {
	__obf_0d00aca7befe11fb := make(map[string]Item)
	return __obf_18fd55b4bac2f8b2(__obf_87cea4da8de7826c, __obf_715327e8eeb7b279, __obf_0d00aca7befe11fb)
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
func NewFrom(__obf_87cea4da8de7826c, __obf_715327e8eeb7b279 time.Duration, __obf_0d00aca7befe11fb map[string]Item) *MapCache {
	return __obf_18fd55b4bac2f8b2(__obf_87cea4da8de7826c, __obf_715327e8eeb7b279, __obf_0d00aca7befe11fb)
}
