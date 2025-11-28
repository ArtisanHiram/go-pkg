// Copyright (C) 2012 Numerotron Inc.
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package consistent provides a consistent hashing function.
//
// Consistent hashing is often used to distribute requests to a changing set of servers.  For example,
// say you have some cache servers cacheA, cacheB, and cacheC.  You want to decide which cache server
// to use to look up information on a user.
//
// You could use a typical hash table and hash the user id
// to one of cacheA, cacheB, or cacheC.  But with a typical hash table, if you add or remove a server,
// almost all keys will get remapped to different results, which basically could bring your service
// to a grinding halt while the caches get rebuilt.
//
// With a consistent hash, adding or removing a server drastically reduces the number of keys that
// get remapped.
//
// Read more about consistent hashing on wikipedia:  http://en.wikipedia.org/wiki/Consistent_hashing
package __obf_7f8161d9c33aa9f1 // import "stathat.com/c/consistent"

import (
	"errors"
	"hash/crc32"
	"hash/fnv"
	"sort"
	"strconv"
	"sync"
)

type Member interface {
	String() string
}

type __obf_d29d3fb427d21ec7 []uint32

// Len returns the length of the uints array.
func (__obf_67e9a335fe07ee98 __obf_d29d3fb427d21ec7) Len() int { return len(__obf_67e9a335fe07ee98) }

// Less returns true if element i is less than element j.
func (__obf_67e9a335fe07ee98 __obf_d29d3fb427d21ec7) Less(__obf_29fac533cbfc20ae, __obf_f680f955f86eb149 int) bool {
	return __obf_67e9a335fe07ee98[__obf_29fac533cbfc20ae] < __obf_67e9a335fe07ee98[__obf_f680f955f86eb149]
}

// Swap exchanges elements i and j.
func (__obf_67e9a335fe07ee98 __obf_d29d3fb427d21ec7) Swap(__obf_29fac533cbfc20ae, __obf_f680f955f86eb149 int) {
	__obf_67e9a335fe07ee98[__obf_29fac533cbfc20ae], __obf_67e9a335fe07ee98[__obf_f680f955f86eb149] = __obf_67e9a335fe07ee98[__obf_f680f955f86eb149], __obf_67e9a335fe07ee98[__obf_29fac533cbfc20ae]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_7f7a8a3949a7e000 map[uint32]M
	__obf_b81d9faee107a3fa map[string]bool
	__obf_b97f1d9eaa7755f2 __obf_d29d3fb427d21ec7
	NumberOfReplicas       int
	__obf_9f481bcb2119bb31 int64
	__obf_f0f9cd65f36f89f5 [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_ef13691bd4d43783 := new(Consistent[M])
	__obf_ef13691bd4d43783.NumberOfReplicas = 20
	__obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000 = make(map[uint32]M)
	__obf_ef13691bd4d43783.__obf_b81d9faee107a3fa = make(map[string]bool)
	return __obf_ef13691bd4d43783
}

// eltKey generates a string key for an element with an index.
func (__obf_ef13691bd4d43783 *Consistent[M]) __obf_686d94cab328cc90(__obf_a5efe47fd419da20 string, __obf_b08f1fac49664224 int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_b08f1fac49664224) + __obf_a5efe47fd419da20
}

// Add inserts a string element in the consistent hash.
func (__obf_ef13691bd4d43783 *Consistent[M]) Add(__obf_a5efe47fd419da20 M) {
	__obf_ef13691bd4d43783.Lock()
	defer __obf_ef13691bd4d43783.Unlock()
	__obf_ef13691bd4d43783.__obf_cc9f9f8300ebe80d(__obf_a5efe47fd419da20)
}

// need c.Lock() before calling
func (__obf_ef13691bd4d43783 *Consistent[M]) __obf_cc9f9f8300ebe80d(__obf_a5efe47fd419da20 M) {
	for __obf_29fac533cbfc20ae := 0; __obf_29fac533cbfc20ae < __obf_ef13691bd4d43783.NumberOfReplicas; __obf_29fac533cbfc20ae++ {
		__obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000[__obf_ef13691bd4d43783.__obf_e676a3fe3cb0439e(__obf_ef13691bd4d43783.__obf_686d94cab328cc90(__obf_a5efe47fd419da20.String(), __obf_29fac533cbfc20ae))] = __obf_a5efe47fd419da20
	}
	__obf_ef13691bd4d43783.__obf_b81d9faee107a3fa[__obf_a5efe47fd419da20.String()] = true
	__obf_ef13691bd4d43783.__obf_1dc4fde087144b28()
	__obf_ef13691bd4d43783.__obf_9f481bcb2119bb31++
}

// Remove removes an element from the hash.
func (__obf_ef13691bd4d43783 *Consistent[M]) Remove(__obf_a5efe47fd419da20 M) {
	__obf_ef13691bd4d43783.Lock()
	defer __obf_ef13691bd4d43783.Unlock()
	__obf_ef13691bd4d43783.__obf_12685e2c21b61642(__obf_a5efe47fd419da20.String())
}

// need c.Lock() before calling
func (__obf_ef13691bd4d43783 *Consistent[M]) __obf_12685e2c21b61642(__obf_a5efe47fd419da20 string) {
	for __obf_29fac533cbfc20ae := 0; __obf_29fac533cbfc20ae < __obf_ef13691bd4d43783.NumberOfReplicas; __obf_29fac533cbfc20ae++ {
		delete(__obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000, __obf_ef13691bd4d43783.__obf_e676a3fe3cb0439e(__obf_ef13691bd4d43783.__obf_686d94cab328cc90(__obf_a5efe47fd419da20, __obf_29fac533cbfc20ae)))
	}
	delete(__obf_ef13691bd4d43783.__obf_b81d9faee107a3fa, __obf_a5efe47fd419da20)
	__obf_ef13691bd4d43783.__obf_1dc4fde087144b28()
	__obf_ef13691bd4d43783.__obf_9f481bcb2119bb31--
}

// Set sets all the elements in the hash.  If there are existing elements not
// present in elts, they will be removed.
func (__obf_ef13691bd4d43783 *Consistent[M]) Set(__obf_9b79afb29e289224 []M) {
	__obf_ef13691bd4d43783.Lock()
	defer __obf_ef13691bd4d43783.Unlock()
	for __obf_dc282976f5b462d7 := range __obf_ef13691bd4d43783.__obf_b81d9faee107a3fa {
		__obf_0fd521ad81ed2060 := false
		for _, __obf_6516a398acb7ade2 := range __obf_9b79afb29e289224 {
			if __obf_dc282976f5b462d7 == __obf_6516a398acb7ade2.String() {
				__obf_0fd521ad81ed2060 = true
				break
			}
		}
		if !__obf_0fd521ad81ed2060 {
			__obf_ef13691bd4d43783.__obf_12685e2c21b61642(__obf_dc282976f5b462d7)
		}
	}
	for _, __obf_6516a398acb7ade2 := range __obf_9b79afb29e289224 {
		_, __obf_6d28d61fc6145691 := __obf_ef13691bd4d43783.__obf_b81d9faee107a3fa[__obf_6516a398acb7ade2.String()]
		if __obf_6d28d61fc6145691 {
			continue
		}
		__obf_ef13691bd4d43783.__obf_cc9f9f8300ebe80d(__obf_6516a398acb7ade2)
	}
}

func (__obf_ef13691bd4d43783 *Consistent[M]) Members() []string {
	__obf_ef13691bd4d43783.RLock()
	defer __obf_ef13691bd4d43783.RUnlock()
	var __obf_d312093d87264870 []string
	for __obf_dc282976f5b462d7 := range __obf_ef13691bd4d43783.__obf_b81d9faee107a3fa {
		__obf_d312093d87264870 = append(__obf_d312093d87264870, __obf_dc282976f5b462d7)
	}
	return __obf_d312093d87264870
}

// Get returns an element close to where name hashes to in the circle.
func (__obf_ef13691bd4d43783 *Consistent[M]) Get(__obf_bc11348cd2922609 string) (__obf_b8a7ce890003a3ed M, __obf_9758acbd047c69f9 error) {
	__obf_ef13691bd4d43783.RLock()
	defer __obf_ef13691bd4d43783.RUnlock()
	if len(__obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000) == 0 {
		__obf_9758acbd047c69f9 = ErrEmptyCircle
		return
	}
	__obf_59bb96d2c3c56f5d := __obf_ef13691bd4d43783.__obf_e676a3fe3cb0439e(__obf_bc11348cd2922609)
	__obf_29fac533cbfc20ae := __obf_ef13691bd4d43783.__obf_b6a8103cf28f0bd0(__obf_59bb96d2c3c56f5d)
	__obf_b8a7ce890003a3ed = __obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000[__obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2[__obf_29fac533cbfc20ae]]
	return
}

func (__obf_ef13691bd4d43783 *Consistent[M]) __obf_b6a8103cf28f0bd0(__obf_59bb96d2c3c56f5d uint32) (__obf_29fac533cbfc20ae int) {
	__obf_cc88324d2b954f1e := func(__obf_67e9a335fe07ee98 int) bool {
		return __obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2[__obf_67e9a335fe07ee98] > __obf_59bb96d2c3c56f5d
	}
	__obf_29fac533cbfc20ae = sort.Search(len(__obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2), __obf_cc88324d2b954f1e)
	if __obf_29fac533cbfc20ae >= len(__obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2) {
		__obf_29fac533cbfc20ae = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_ef13691bd4d43783 *Consistent[M]) GetTwo(__obf_bc11348cd2922609 string) (__obf_c81a42470361c64b M, __obf_469b206113527cdf M, __obf_9758acbd047c69f9 error) {
	__obf_ef13691bd4d43783.RLock()
	defer __obf_ef13691bd4d43783.RUnlock()
	if len(__obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000) == 0 {
		__obf_9758acbd047c69f9 = ErrEmptyCircle
		return
	}
	__obf_59bb96d2c3c56f5d := __obf_ef13691bd4d43783.__obf_e676a3fe3cb0439e(__obf_bc11348cd2922609)
	__obf_29fac533cbfc20ae := __obf_ef13691bd4d43783.__obf_b6a8103cf28f0bd0(__obf_59bb96d2c3c56f5d)
	__obf_c81a42470361c64b = __obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000[__obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2[__obf_29fac533cbfc20ae]]

	if __obf_ef13691bd4d43783.__obf_9f481bcb2119bb31 == 1 {
		return
	}

	__obf_f8160f87405f098c := __obf_29fac533cbfc20ae
	for __obf_29fac533cbfc20ae = __obf_f8160f87405f098c + 1; __obf_29fac533cbfc20ae != __obf_f8160f87405f098c; __obf_29fac533cbfc20ae++ {
		if __obf_29fac533cbfc20ae >= len(__obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2) {
			__obf_29fac533cbfc20ae = 0
		}
		__obf_469b206113527cdf = __obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000[__obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2[__obf_29fac533cbfc20ae]]
		if __obf_469b206113527cdf.String() != __obf_c81a42470361c64b.String() {
			break
		}
	}
	return __obf_c81a42470361c64b, __obf_469b206113527cdf, nil
}

// GetN returns the N closest distinct elements to the name input in the circle.
func (__obf_ef13691bd4d43783 *Consistent[M]) GetN(__obf_bc11348cd2922609 string, __obf_de11922cf1006187 int) (__obf_b8a7ce890003a3ed []M, __obf_9758acbd047c69f9 error) {
	__obf_ef13691bd4d43783.RLock()
	defer __obf_ef13691bd4d43783.RUnlock()

	if len(__obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000) == 0 {
		__obf_9758acbd047c69f9 = ErrEmptyCircle
		return
	}

	if __obf_ef13691bd4d43783.__obf_9f481bcb2119bb31 < int64(__obf_de11922cf1006187) {
		__obf_de11922cf1006187 = int(__obf_ef13691bd4d43783.__obf_9f481bcb2119bb31)
	}

	var (
		__obf_59bb96d2c3c56f5d = __obf_ef13691bd4d43783.__obf_e676a3fe3cb0439e(__obf_bc11348cd2922609)
		__obf_29fac533cbfc20ae = __obf_ef13691bd4d43783.__obf_b6a8103cf28f0bd0(__obf_59bb96d2c3c56f5d)
		__obf_f8160f87405f098c = __obf_29fac533cbfc20ae
		__obf_4f964c3f131b150c = __obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000[__obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2[__obf_29fac533cbfc20ae]]
	)
	__obf_b8a7ce890003a3ed = make([]M, 0, __obf_de11922cf1006187)
	__obf_b8a7ce890003a3ed = append(__obf_b8a7ce890003a3ed, __obf_4f964c3f131b150c)

	if len(__obf_b8a7ce890003a3ed) == __obf_de11922cf1006187 {
		return __obf_b8a7ce890003a3ed, nil
	}

	for __obf_29fac533cbfc20ae = __obf_f8160f87405f098c + 1; __obf_29fac533cbfc20ae != __obf_f8160f87405f098c; __obf_29fac533cbfc20ae++ {
		if __obf_29fac533cbfc20ae >= len(__obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2) {
			__obf_29fac533cbfc20ae = 0
		}
		__obf_4f964c3f131b150c = __obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000[__obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2[__obf_29fac533cbfc20ae]]
		if !__obf_10c8f222ab061458(__obf_b8a7ce890003a3ed, __obf_4f964c3f131b150c) {
			__obf_b8a7ce890003a3ed = append(__obf_b8a7ce890003a3ed, __obf_4f964c3f131b150c)
		}
		if len(__obf_b8a7ce890003a3ed) == __obf_de11922cf1006187 {
			break
		}
	}

	return __obf_b8a7ce890003a3ed, nil
}

func (__obf_ef13691bd4d43783 *Consistent[M]) __obf_e676a3fe3cb0439e(__obf_59bb96d2c3c56f5d string) uint32 {
	if __obf_ef13691bd4d43783.UseFnv {
		return __obf_ef13691bd4d43783.__obf_910adb9596117e58(__obf_59bb96d2c3c56f5d)
	}
	return __obf_ef13691bd4d43783.__obf_2d73374aac790bd1(__obf_59bb96d2c3c56f5d)
}

func (__obf_ef13691bd4d43783 *Consistent[M]) __obf_2d73374aac790bd1(__obf_59bb96d2c3c56f5d string) uint32 {
	if len(__obf_59bb96d2c3c56f5d) < 64 {
		var __obf_f0f9cd65f36f89f5 [64]byte
		copy(__obf_f0f9cd65f36f89f5[:], __obf_59bb96d2c3c56f5d)
		return crc32.ChecksumIEEE(__obf_f0f9cd65f36f89f5[:len(__obf_59bb96d2c3c56f5d)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_59bb96d2c3c56f5d))
}

func (__obf_ef13691bd4d43783 *Consistent[M]) __obf_910adb9596117e58(__obf_59bb96d2c3c56f5d string) uint32 {
	__obf_3374197ae552360a := fnv.New32a()
	__obf_3374197ae552360a.Write([]byte(__obf_59bb96d2c3c56f5d))
	return __obf_3374197ae552360a.Sum32()
}

func (__obf_ef13691bd4d43783 *Consistent[M]) __obf_1dc4fde087144b28() {
	__obf_eff3542d3aacbc67 := __obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2[:0]
	//reallocate if we're holding on to too much (1/4th)
	if cap(__obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2)/(__obf_ef13691bd4d43783.NumberOfReplicas*4) > len(__obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000) {
		__obf_eff3542d3aacbc67 = nil
	}
	for __obf_dc282976f5b462d7 := range __obf_ef13691bd4d43783.__obf_7f7a8a3949a7e000 {
		__obf_eff3542d3aacbc67 = append(__obf_eff3542d3aacbc67, __obf_dc282976f5b462d7)
	}
	sort.Sort(__obf_eff3542d3aacbc67)
	__obf_ef13691bd4d43783.__obf_b97f1d9eaa7755f2 = __obf_eff3542d3aacbc67
}

func __obf_10c8f222ab061458[M Member](__obf_425682f8c7b31741 []M, __obf_d76822e7159b6f40 M) bool {
	for _, __obf_d312093d87264870 := range __obf_425682f8c7b31741 {
		if __obf_d312093d87264870.String() == __obf_d76822e7159b6f40.String() {
			return true
		}
	}
	return false
}
