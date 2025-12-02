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
package __obf_4e34da8c68f3e562 // import "stathat.com/c/consistent"

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

type __obf_5d7b70c8d54be23c []uint32

// Len returns the length of the uints array.
func (__obf_98984a85e9115c26 __obf_5d7b70c8d54be23c) Len() int { return len(__obf_98984a85e9115c26) }

// Less returns true if element i is less than element j.
func (__obf_98984a85e9115c26 __obf_5d7b70c8d54be23c) Less(__obf_415752691d5c9316, __obf_16149d855f54e402 int) bool {
	return __obf_98984a85e9115c26[__obf_415752691d5c9316] < __obf_98984a85e9115c26[__obf_16149d855f54e402]
}

// Swap exchanges elements i and j.
func (__obf_98984a85e9115c26 __obf_5d7b70c8d54be23c) Swap(__obf_415752691d5c9316, __obf_16149d855f54e402 int) {
	__obf_98984a85e9115c26[__obf_415752691d5c9316], __obf_98984a85e9115c26[__obf_16149d855f54e402] = __obf_98984a85e9115c26[__obf_16149d855f54e402], __obf_98984a85e9115c26[__obf_415752691d5c9316]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_a818bbc94d2dee5d map[uint32]M
	__obf_d82cb5f62465e84e map[string]bool
	__obf_939f971304175ab5 __obf_5d7b70c8d54be23c

	NumberOfReplicas       int
	__obf_e249ba1a0abfbca1 int64
	__obf_43876eaa82c3005f [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_46fda7b1298990c7 := new(Consistent[M])
	__obf_46fda7b1298990c7.
		NumberOfReplicas = 20
	__obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d = make(map[uint32]M)
	__obf_46fda7b1298990c7.__obf_d82cb5f62465e84e = make(map[string]bool)
	return __obf_46fda7b1298990c7
}

// eltKey generates a string key for an element with an index.
func (__obf_46fda7b1298990c7 *Consistent[M]) __obf_9052c919cdf735f3(__obf_f348338e1fd82042 string, __obf_afe07540a2345e7a int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_afe07540a2345e7a) + __obf_f348338e1fd82042
}

// Add inserts a string element in the consistent hash.
func (__obf_46fda7b1298990c7 *Consistent[M]) Add(__obf_f348338e1fd82042 M) {
	__obf_46fda7b1298990c7.
		Lock()
	defer __obf_46fda7b1298990c7.Unlock()
	__obf_46fda7b1298990c7.

		// need c.Lock() before calling
		__obf_89bbea5cdfb188a0(__obf_f348338e1fd82042)
}

func (__obf_46fda7b1298990c7 *Consistent[M]) __obf_89bbea5cdfb188a0(__obf_f348338e1fd82042 M) {
	for __obf_415752691d5c9316 := 0; __obf_415752691d5c9316 < __obf_46fda7b1298990c7.NumberOfReplicas; __obf_415752691d5c9316++ {
		__obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d[__obf_46fda7b1298990c7.__obf_640f515be599ad00(__obf_46fda7b1298990c7.__obf_9052c919cdf735f3(__obf_f348338e1fd82042.String(), __obf_415752691d5c9316))] = __obf_f348338e1fd82042
	}
	__obf_46fda7b1298990c7.__obf_d82cb5f62465e84e[__obf_f348338e1fd82042.String()] = true
	__obf_46fda7b1298990c7.__obf_d2ad337730c98c3b()
	__obf_46fda7b1298990c7.

		// Remove removes an element from the hash.
		__obf_e249ba1a0abfbca1++
}

func (__obf_46fda7b1298990c7 *Consistent[M]) Remove(__obf_f348338e1fd82042 M) {
	__obf_46fda7b1298990c7.
		Lock()
	defer __obf_46fda7b1298990c7.Unlock()
	__obf_46fda7b1298990c7.__obf_3fe7cf5409327224(__obf_f348338e1fd82042.String())
}

// need c.Lock() before calling
func (__obf_46fda7b1298990c7 *Consistent[M]) __obf_3fe7cf5409327224(__obf_f348338e1fd82042 string) {
	for __obf_415752691d5c9316 := 0; __obf_415752691d5c9316 < __obf_46fda7b1298990c7.NumberOfReplicas; __obf_415752691d5c9316++ {
		delete(__obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d, __obf_46fda7b1298990c7.__obf_640f515be599ad00(__obf_46fda7b1298990c7.__obf_9052c919cdf735f3(__obf_f348338e1fd82042, __obf_415752691d5c9316)))
	}
	delete(__obf_46fda7b1298990c7.__obf_d82cb5f62465e84e, __obf_f348338e1fd82042)
	__obf_46fda7b1298990c7.__obf_d2ad337730c98c3b()
	__obf_46fda7b1298990c7.

		// Set sets all the elements in the hash.  If there are existing elements not
		// present in elts, they will be removed.
		__obf_e249ba1a0abfbca1--
}

func (__obf_46fda7b1298990c7 *Consistent[M]) Set(__obf_06a4c3bfbda2a7ac []M) {
	__obf_46fda7b1298990c7.
		Lock()
	defer __obf_46fda7b1298990c7.Unlock()
	for __obf_72b35587496e1b8e := range __obf_46fda7b1298990c7.__obf_d82cb5f62465e84e {
		__obf_e988654abe050183 := false
		for _, __obf_784254a3461324fb := range __obf_06a4c3bfbda2a7ac {
			if __obf_72b35587496e1b8e == __obf_784254a3461324fb.String() {
				__obf_e988654abe050183 = true
				break
			}
		}
		if !__obf_e988654abe050183 {
			__obf_46fda7b1298990c7.__obf_3fe7cf5409327224(__obf_72b35587496e1b8e)
		}
	}
	for _, __obf_784254a3461324fb := range __obf_06a4c3bfbda2a7ac {
		_, __obf_f23278670c7cbfe2 := __obf_46fda7b1298990c7.__obf_d82cb5f62465e84e[__obf_784254a3461324fb.String()]
		if __obf_f23278670c7cbfe2 {
			continue
		}
		__obf_46fda7b1298990c7.__obf_89bbea5cdfb188a0(__obf_784254a3461324fb)
	}
}

func (__obf_46fda7b1298990c7 *Consistent[M]) Members() []string {
	__obf_46fda7b1298990c7.
		RLock()
	defer __obf_46fda7b1298990c7.RUnlock()
	var __obf_6e45f2b870910723 []string
	for __obf_72b35587496e1b8e := range __obf_46fda7b1298990c7.__obf_d82cb5f62465e84e {
		__obf_6e45f2b870910723 = append(__obf_6e45f2b870910723,

			// Get returns an element close to where name hashes to in the circle.
			__obf_72b35587496e1b8e)
	}
	return __obf_6e45f2b870910723
}

func (__obf_46fda7b1298990c7 *Consistent[M]) Get(__obf_dbac0194ffa6f253 string) (__obf_4ca1ed24f24e0e11 M, __obf_269384ef9b9f4e90 error) {
	__obf_46fda7b1298990c7.
		RLock()
	defer __obf_46fda7b1298990c7.RUnlock()
	if len(__obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d) == 0 {
		__obf_269384ef9b9f4e90 = ErrEmptyCircle
		return
	}
	__obf_530221ffc85407ea := __obf_46fda7b1298990c7.__obf_640f515be599ad00(__obf_dbac0194ffa6f253)
	__obf_415752691d5c9316 := __obf_46fda7b1298990c7.__obf_008ae5e6f013c269(__obf_530221ffc85407ea)
	__obf_4ca1ed24f24e0e11 = __obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d[__obf_46fda7b1298990c7.__obf_939f971304175ab5[__obf_415752691d5c9316]]
	return
}

func (__obf_46fda7b1298990c7 *Consistent[M]) __obf_008ae5e6f013c269(__obf_530221ffc85407ea uint32) (__obf_415752691d5c9316 int) {
	__obf_fda4c9a1bb183d94 := func(__obf_98984a85e9115c26 int) bool {
		return __obf_46fda7b1298990c7.__obf_939f971304175ab5[__obf_98984a85e9115c26] > __obf_530221ffc85407ea
	}
	__obf_415752691d5c9316 = sort.Search(len(__obf_46fda7b1298990c7.__obf_939f971304175ab5), __obf_fda4c9a1bb183d94)
	if __obf_415752691d5c9316 >= len(__obf_46fda7b1298990c7.__obf_939f971304175ab5) {
		__obf_415752691d5c9316 = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_46fda7b1298990c7 *Consistent[M]) GetTwo(__obf_dbac0194ffa6f253 string) (__obf_e55af27bc09e7d1b M, __obf_7df26e36a9360ff0 M, __obf_269384ef9b9f4e90 error) {
	__obf_46fda7b1298990c7.
		RLock()
	defer __obf_46fda7b1298990c7.RUnlock()
	if len(__obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d) == 0 {
		__obf_269384ef9b9f4e90 = ErrEmptyCircle
		return
	}
	__obf_530221ffc85407ea := __obf_46fda7b1298990c7.__obf_640f515be599ad00(__obf_dbac0194ffa6f253)
	__obf_415752691d5c9316 := __obf_46fda7b1298990c7.__obf_008ae5e6f013c269(__obf_530221ffc85407ea)
	__obf_e55af27bc09e7d1b = __obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d[__obf_46fda7b1298990c7.__obf_939f971304175ab5[__obf_415752691d5c9316]]

	if __obf_46fda7b1298990c7.__obf_e249ba1a0abfbca1 == 1 {
		return
	}
	__obf_475bedd3866e0987 := __obf_415752691d5c9316
	for __obf_415752691d5c9316 = __obf_475bedd3866e0987 + 1; __obf_415752691d5c9316 != __obf_475bedd3866e0987; __obf_415752691d5c9316++ {
		if __obf_415752691d5c9316 >= len(__obf_46fda7b1298990c7.__obf_939f971304175ab5) {
			__obf_415752691d5c9316 = 0
		}
		__obf_7df26e36a9360ff0 = __obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d[__obf_46fda7b1298990c7.__obf_939f971304175ab5[__obf_415752691d5c9316]]
		if __obf_7df26e36a9360ff0.String() != __obf_e55af27bc09e7d1b.String() {
			break
		}
	}
	return __obf_e55af27bc09e7d1b,

		// GetN returns the N closest distinct elements to the name input in the circle.
		__obf_7df26e36a9360ff0, nil
}

func (__obf_46fda7b1298990c7 *Consistent[M]) GetN(__obf_dbac0194ffa6f253 string, __obf_c87052c701cf313c int) (__obf_4ca1ed24f24e0e11 []M, __obf_269384ef9b9f4e90 error) {
	__obf_46fda7b1298990c7.
		RLock()
	defer __obf_46fda7b1298990c7.RUnlock()

	if len(__obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d) == 0 {
		__obf_269384ef9b9f4e90 = ErrEmptyCircle
		return
	}

	if __obf_46fda7b1298990c7.__obf_e249ba1a0abfbca1 < int64(__obf_c87052c701cf313c) {
		__obf_c87052c701cf313c = int(__obf_46fda7b1298990c7.__obf_e249ba1a0abfbca1)
	}

	var (
		__obf_530221ffc85407ea = __obf_46fda7b1298990c7.__obf_640f515be599ad00(__obf_dbac0194ffa6f253)
		__obf_415752691d5c9316 = __obf_46fda7b1298990c7.__obf_008ae5e6f013c269(__obf_530221ffc85407ea)
		__obf_475bedd3866e0987 = __obf_415752691d5c9316
		__obf_dfc455f707cc951e = __obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d[__obf_46fda7b1298990c7.__obf_939f971304175ab5[__obf_415752691d5c9316]]
	)
	__obf_4ca1ed24f24e0e11 = make([]M, 0, __obf_c87052c701cf313c)
	__obf_4ca1ed24f24e0e11 = append(__obf_4ca1ed24f24e0e11, __obf_dfc455f707cc951e)

	if len(__obf_4ca1ed24f24e0e11) == __obf_c87052c701cf313c {
		return __obf_4ca1ed24f24e0e11, nil
	}

	for __obf_415752691d5c9316 = __obf_475bedd3866e0987 + 1; __obf_415752691d5c9316 != __obf_475bedd3866e0987; __obf_415752691d5c9316++ {
		if __obf_415752691d5c9316 >= len(__obf_46fda7b1298990c7.__obf_939f971304175ab5) {
			__obf_415752691d5c9316 = 0
		}
		__obf_dfc455f707cc951e = __obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d[__obf_46fda7b1298990c7.__obf_939f971304175ab5[__obf_415752691d5c9316]]
		if !__obf_21530fd1bc3e35e0(__obf_4ca1ed24f24e0e11, __obf_dfc455f707cc951e) {
			__obf_4ca1ed24f24e0e11 = append(__obf_4ca1ed24f24e0e11, __obf_dfc455f707cc951e)
		}
		if len(__obf_4ca1ed24f24e0e11) == __obf_c87052c701cf313c {
			break
		}
	}

	return __obf_4ca1ed24f24e0e11, nil
}

func (__obf_46fda7b1298990c7 *Consistent[M]) __obf_640f515be599ad00(__obf_530221ffc85407ea string) uint32 {
	if __obf_46fda7b1298990c7.UseFnv {
		return __obf_46fda7b1298990c7.__obf_f10576f2130249ea(__obf_530221ffc85407ea)
	}
	return __obf_46fda7b1298990c7.__obf_09db042508544082(__obf_530221ffc85407ea)
}

func (__obf_46fda7b1298990c7 *Consistent[M]) __obf_09db042508544082(__obf_530221ffc85407ea string) uint32 {
	if len(__obf_530221ffc85407ea) < 64 {
		var __obf_43876eaa82c3005f [64]byte
		copy(__obf_43876eaa82c3005f[:], __obf_530221ffc85407ea)
		return crc32.ChecksumIEEE(__obf_43876eaa82c3005f[:len(__obf_530221ffc85407ea)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_530221ffc85407ea))
}

func (__obf_46fda7b1298990c7 *Consistent[M]) __obf_f10576f2130249ea(__obf_530221ffc85407ea string) uint32 {
	__obf_ff60f3a7ab989769 := fnv.New32a()
	__obf_ff60f3a7ab989769.
		Write([]byte(__obf_530221ffc85407ea))
	return __obf_ff60f3a7ab989769.Sum32()
}

func (__obf_46fda7b1298990c7 *Consistent[M]) __obf_d2ad337730c98c3b() {
	__obf_eceaf98b5abecf05 := __obf_46fda7b1298990c7.
		//reallocate if we're holding on to too much (1/4th)
		__obf_939f971304175ab5[:0]

	if cap(__obf_46fda7b1298990c7.__obf_939f971304175ab5)/(__obf_46fda7b1298990c7.NumberOfReplicas*4) > len(__obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d) {
		__obf_eceaf98b5abecf05 = nil
	}
	for __obf_72b35587496e1b8e := range __obf_46fda7b1298990c7.__obf_a818bbc94d2dee5d {
		__obf_eceaf98b5abecf05 = append(__obf_eceaf98b5abecf05, __obf_72b35587496e1b8e)
	}
	sort.Sort(__obf_eceaf98b5abecf05)
	__obf_46fda7b1298990c7.__obf_939f971304175ab5 = __obf_eceaf98b5abecf05
}

func __obf_21530fd1bc3e35e0[M Member](__obf_28b5581bb1fd7696 []M, __obf_1e5ce82914f4e295 M) bool {
	for _, __obf_6e45f2b870910723 := range __obf_28b5581bb1fd7696 {
		if __obf_6e45f2b870910723.String() == __obf_1e5ce82914f4e295.String() {
			return true
		}
	}
	return false
}
