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
package __obf_ed5d4cf47c78219f // import "stathat.com/c/consistent"

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

type __obf_7c2e129065462a8f []uint32

// Len returns the length of the uints array.
func (__obf_658ec30b86cbad50 __obf_7c2e129065462a8f) Len() int { return len(__obf_658ec30b86cbad50) }

// Less returns true if element i is less than element j.
func (__obf_658ec30b86cbad50 __obf_7c2e129065462a8f) Less(__obf_59cad04ec9ba116c, __obf_529ee647fcdceaf4 int) bool {
	return __obf_658ec30b86cbad50[__obf_59cad04ec9ba116c] < __obf_658ec30b86cbad50[__obf_529ee647fcdceaf4]
}

// Swap exchanges elements i and j.
func (__obf_658ec30b86cbad50 __obf_7c2e129065462a8f) Swap(__obf_59cad04ec9ba116c, __obf_529ee647fcdceaf4 int) {
	__obf_658ec30b86cbad50[__obf_59cad04ec9ba116c], __obf_658ec30b86cbad50[__obf_529ee647fcdceaf4] = __obf_658ec30b86cbad50[__obf_529ee647fcdceaf4], __obf_658ec30b86cbad50[__obf_59cad04ec9ba116c]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_b32431a0085e4728 map[uint32]M
	__obf_573b9f4adff76913 map[string]bool
	__obf_f8dad1fa545d0b9c __obf_7c2e129065462a8f

	NumberOfReplicas       int
	__obf_525f930b687a428b int64
	__obf_8f1a9a3d1d85b5b1 [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_2cc5392205ea6f7e := new(Consistent[M])
	__obf_2cc5392205ea6f7e.
		NumberOfReplicas = 20
	__obf_2cc5392205ea6f7e.__obf_b32431a0085e4728 = make(map[uint32]M)
	__obf_2cc5392205ea6f7e.__obf_573b9f4adff76913 = make(map[string]bool)
	return __obf_2cc5392205ea6f7e
}

// eltKey generates a string key for an element with an index.
func (__obf_2cc5392205ea6f7e *Consistent[M]) __obf_b63f0b5603aed972(__obf_fd70529d25a00143 string, __obf_649047368042f632 int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_649047368042f632) + __obf_fd70529d25a00143
}

// Add inserts a string element in the consistent hash.
func (__obf_2cc5392205ea6f7e *Consistent[M]) Add(__obf_fd70529d25a00143 M) {
	__obf_2cc5392205ea6f7e.
		Lock()
	defer __obf_2cc5392205ea6f7e.Unlock()
	__obf_2cc5392205ea6f7e.

		// need c.Lock() before calling
		__obf_87f6e684e30b943c(__obf_fd70529d25a00143)
}

func (__obf_2cc5392205ea6f7e *Consistent[M]) __obf_87f6e684e30b943c(__obf_fd70529d25a00143 M) {
	for __obf_59cad04ec9ba116c := 0; __obf_59cad04ec9ba116c < __obf_2cc5392205ea6f7e.NumberOfReplicas; __obf_59cad04ec9ba116c++ {
		__obf_2cc5392205ea6f7e.__obf_b32431a0085e4728[__obf_2cc5392205ea6f7e.__obf_219b02727add8ba0(__obf_2cc5392205ea6f7e.__obf_b63f0b5603aed972(__obf_fd70529d25a00143.String(), __obf_59cad04ec9ba116c))] = __obf_fd70529d25a00143
	}
	__obf_2cc5392205ea6f7e.__obf_573b9f4adff76913[__obf_fd70529d25a00143.String()] = true
	__obf_2cc5392205ea6f7e.__obf_4b41ceae11e6daee()
	__obf_2cc5392205ea6f7e.

		// Remove removes an element from the hash.
		__obf_525f930b687a428b++
}

func (__obf_2cc5392205ea6f7e *Consistent[M]) Remove(__obf_fd70529d25a00143 M) {
	__obf_2cc5392205ea6f7e.
		Lock()
	defer __obf_2cc5392205ea6f7e.Unlock()
	__obf_2cc5392205ea6f7e.__obf_b8a676a317f4e96d(__obf_fd70529d25a00143.String())
}

// need c.Lock() before calling
func (__obf_2cc5392205ea6f7e *Consistent[M]) __obf_b8a676a317f4e96d(__obf_fd70529d25a00143 string) {
	for __obf_59cad04ec9ba116c := 0; __obf_59cad04ec9ba116c < __obf_2cc5392205ea6f7e.NumberOfReplicas; __obf_59cad04ec9ba116c++ {
		delete(__obf_2cc5392205ea6f7e.__obf_b32431a0085e4728, __obf_2cc5392205ea6f7e.__obf_219b02727add8ba0(__obf_2cc5392205ea6f7e.__obf_b63f0b5603aed972(__obf_fd70529d25a00143, __obf_59cad04ec9ba116c)))
	}
	delete(__obf_2cc5392205ea6f7e.__obf_573b9f4adff76913, __obf_fd70529d25a00143)
	__obf_2cc5392205ea6f7e.__obf_4b41ceae11e6daee()
	__obf_2cc5392205ea6f7e.

		// Set sets all the elements in the hash.  If there are existing elements not
		// present in elts, they will be removed.
		__obf_525f930b687a428b--
}

func (__obf_2cc5392205ea6f7e *Consistent[M]) Set(__obf_353b9985010f81d5 []M) {
	__obf_2cc5392205ea6f7e.
		Lock()
	defer __obf_2cc5392205ea6f7e.Unlock()
	for __obf_9c16974b86e0c1f1 := range __obf_2cc5392205ea6f7e.__obf_573b9f4adff76913 {
		__obf_eb9b5981c723327a := false
		for _, __obf_d924236c0c664a5c := range __obf_353b9985010f81d5 {
			if __obf_9c16974b86e0c1f1 == __obf_d924236c0c664a5c.String() {
				__obf_eb9b5981c723327a = true
				break
			}
		}
		if !__obf_eb9b5981c723327a {
			__obf_2cc5392205ea6f7e.__obf_b8a676a317f4e96d(__obf_9c16974b86e0c1f1)
		}
	}
	for _, __obf_d924236c0c664a5c := range __obf_353b9985010f81d5 {
		_, __obf_43923daa16e98436 := __obf_2cc5392205ea6f7e.__obf_573b9f4adff76913[__obf_d924236c0c664a5c.String()]
		if __obf_43923daa16e98436 {
			continue
		}
		__obf_2cc5392205ea6f7e.__obf_87f6e684e30b943c(__obf_d924236c0c664a5c)
	}
}

func (__obf_2cc5392205ea6f7e *Consistent[M]) Members() []string {
	__obf_2cc5392205ea6f7e.
		RLock()
	defer __obf_2cc5392205ea6f7e.RUnlock()
	var __obf_5012e0ddf72cce0d []string
	for __obf_9c16974b86e0c1f1 := range __obf_2cc5392205ea6f7e.__obf_573b9f4adff76913 {
		__obf_5012e0ddf72cce0d = append(__obf_5012e0ddf72cce0d,

			// Get returns an element close to where name hashes to in the circle.
			__obf_9c16974b86e0c1f1)
	}
	return __obf_5012e0ddf72cce0d
}

func (__obf_2cc5392205ea6f7e *Consistent[M]) Get(__obf_bacbbe0483def87e string) (__obf_6a15cf9c00ce4b33 M, __obf_af02760b7a79d19c error) {
	__obf_2cc5392205ea6f7e.
		RLock()
	defer __obf_2cc5392205ea6f7e.RUnlock()
	if len(__obf_2cc5392205ea6f7e.__obf_b32431a0085e4728) == 0 {
		__obf_af02760b7a79d19c = ErrEmptyCircle
		return
	}
	__obf_0db93f791e5a5975 := __obf_2cc5392205ea6f7e.__obf_219b02727add8ba0(__obf_bacbbe0483def87e)
	__obf_59cad04ec9ba116c := __obf_2cc5392205ea6f7e.__obf_700638f80f8460fd(__obf_0db93f791e5a5975)
	__obf_6a15cf9c00ce4b33 = __obf_2cc5392205ea6f7e.__obf_b32431a0085e4728[__obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c[__obf_59cad04ec9ba116c]]
	return
}

func (__obf_2cc5392205ea6f7e *Consistent[M]) __obf_700638f80f8460fd(__obf_0db93f791e5a5975 uint32) (__obf_59cad04ec9ba116c int) {
	__obf_44ba288d561a39d7 := func(__obf_658ec30b86cbad50 int) bool {
		return __obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c[__obf_658ec30b86cbad50] > __obf_0db93f791e5a5975
	}
	__obf_59cad04ec9ba116c = sort.Search(len(__obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c), __obf_44ba288d561a39d7)
	if __obf_59cad04ec9ba116c >= len(__obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c) {
		__obf_59cad04ec9ba116c = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_2cc5392205ea6f7e *Consistent[M]) GetTwo(__obf_bacbbe0483def87e string) (__obf_ac4c60c2b1f7684d M, __obf_c8b15a51b3fd2b39 M, __obf_af02760b7a79d19c error) {
	__obf_2cc5392205ea6f7e.
		RLock()
	defer __obf_2cc5392205ea6f7e.RUnlock()
	if len(__obf_2cc5392205ea6f7e.__obf_b32431a0085e4728) == 0 {
		__obf_af02760b7a79d19c = ErrEmptyCircle
		return
	}
	__obf_0db93f791e5a5975 := __obf_2cc5392205ea6f7e.__obf_219b02727add8ba0(__obf_bacbbe0483def87e)
	__obf_59cad04ec9ba116c := __obf_2cc5392205ea6f7e.__obf_700638f80f8460fd(__obf_0db93f791e5a5975)
	__obf_ac4c60c2b1f7684d = __obf_2cc5392205ea6f7e.__obf_b32431a0085e4728[__obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c[__obf_59cad04ec9ba116c]]

	if __obf_2cc5392205ea6f7e.__obf_525f930b687a428b == 1 {
		return
	}
	__obf_ca597e6fd49c2796 := __obf_59cad04ec9ba116c
	for __obf_59cad04ec9ba116c = __obf_ca597e6fd49c2796 + 1; __obf_59cad04ec9ba116c != __obf_ca597e6fd49c2796; __obf_59cad04ec9ba116c++ {
		if __obf_59cad04ec9ba116c >= len(__obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c) {
			__obf_59cad04ec9ba116c = 0
		}
		__obf_c8b15a51b3fd2b39 = __obf_2cc5392205ea6f7e.__obf_b32431a0085e4728[__obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c[__obf_59cad04ec9ba116c]]
		if __obf_c8b15a51b3fd2b39.String() != __obf_ac4c60c2b1f7684d.String() {
			break
		}
	}
	return __obf_ac4c60c2b1f7684d,

		// GetN returns the N closest distinct elements to the name input in the circle.
		__obf_c8b15a51b3fd2b39, nil
}

func (__obf_2cc5392205ea6f7e *Consistent[M]) GetN(__obf_bacbbe0483def87e string, __obf_98a1209e16c2f849 int) (__obf_6a15cf9c00ce4b33 []M, __obf_af02760b7a79d19c error) {
	__obf_2cc5392205ea6f7e.
		RLock()
	defer __obf_2cc5392205ea6f7e.RUnlock()

	if len(__obf_2cc5392205ea6f7e.__obf_b32431a0085e4728) == 0 {
		__obf_af02760b7a79d19c = ErrEmptyCircle
		return
	}

	if __obf_2cc5392205ea6f7e.__obf_525f930b687a428b < int64(__obf_98a1209e16c2f849) {
		__obf_98a1209e16c2f849 = int(__obf_2cc5392205ea6f7e.__obf_525f930b687a428b)
	}

	var (
		__obf_0db93f791e5a5975 = __obf_2cc5392205ea6f7e.__obf_219b02727add8ba0(__obf_bacbbe0483def87e)
		__obf_59cad04ec9ba116c = __obf_2cc5392205ea6f7e.__obf_700638f80f8460fd(__obf_0db93f791e5a5975)
		__obf_ca597e6fd49c2796 = __obf_59cad04ec9ba116c
		__obf_885ffbab7c03da55 = __obf_2cc5392205ea6f7e.__obf_b32431a0085e4728[__obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c[__obf_59cad04ec9ba116c]]
	)
	__obf_6a15cf9c00ce4b33 = make([]M, 0, __obf_98a1209e16c2f849)
	__obf_6a15cf9c00ce4b33 = append(__obf_6a15cf9c00ce4b33, __obf_885ffbab7c03da55)

	if len(__obf_6a15cf9c00ce4b33) == __obf_98a1209e16c2f849 {
		return __obf_6a15cf9c00ce4b33, nil
	}

	for __obf_59cad04ec9ba116c = __obf_ca597e6fd49c2796 + 1; __obf_59cad04ec9ba116c != __obf_ca597e6fd49c2796; __obf_59cad04ec9ba116c++ {
		if __obf_59cad04ec9ba116c >= len(__obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c) {
			__obf_59cad04ec9ba116c = 0
		}
		__obf_885ffbab7c03da55 = __obf_2cc5392205ea6f7e.__obf_b32431a0085e4728[__obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c[__obf_59cad04ec9ba116c]]
		if !__obf_89f065dcc001902b(__obf_6a15cf9c00ce4b33, __obf_885ffbab7c03da55) {
			__obf_6a15cf9c00ce4b33 = append(__obf_6a15cf9c00ce4b33, __obf_885ffbab7c03da55)
		}
		if len(__obf_6a15cf9c00ce4b33) == __obf_98a1209e16c2f849 {
			break
		}
	}

	return __obf_6a15cf9c00ce4b33, nil
}

func (__obf_2cc5392205ea6f7e *Consistent[M]) __obf_219b02727add8ba0(__obf_0db93f791e5a5975 string) uint32 {
	if __obf_2cc5392205ea6f7e.UseFnv {
		return __obf_2cc5392205ea6f7e.__obf_296b02d9f085fe2d(__obf_0db93f791e5a5975)
	}
	return __obf_2cc5392205ea6f7e.__obf_44d05b53f628a1aa(__obf_0db93f791e5a5975)
}

func (__obf_2cc5392205ea6f7e *Consistent[M]) __obf_44d05b53f628a1aa(__obf_0db93f791e5a5975 string) uint32 {
	if len(__obf_0db93f791e5a5975) < 64 {
		var __obf_8f1a9a3d1d85b5b1 [64]byte
		copy(__obf_8f1a9a3d1d85b5b1[:], __obf_0db93f791e5a5975)
		return crc32.ChecksumIEEE(__obf_8f1a9a3d1d85b5b1[:len(__obf_0db93f791e5a5975)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_0db93f791e5a5975))
}

func (__obf_2cc5392205ea6f7e *Consistent[M]) __obf_296b02d9f085fe2d(__obf_0db93f791e5a5975 string) uint32 {
	__obf_7e903bae48be7c2f := fnv.New32a()
	__obf_7e903bae48be7c2f.
		Write([]byte(__obf_0db93f791e5a5975))
	return __obf_7e903bae48be7c2f.Sum32()
}

func (__obf_2cc5392205ea6f7e *Consistent[M]) __obf_4b41ceae11e6daee() {
	__obf_821d67ba57c19d43 := __obf_2cc5392205ea6f7e.
		//reallocate if we're holding on to too much (1/4th)
		__obf_f8dad1fa545d0b9c[:0]

	if cap(__obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c)/(__obf_2cc5392205ea6f7e.NumberOfReplicas*4) > len(__obf_2cc5392205ea6f7e.__obf_b32431a0085e4728) {
		__obf_821d67ba57c19d43 = nil
	}
	for __obf_9c16974b86e0c1f1 := range __obf_2cc5392205ea6f7e.__obf_b32431a0085e4728 {
		__obf_821d67ba57c19d43 = append(__obf_821d67ba57c19d43, __obf_9c16974b86e0c1f1)
	}
	sort.Sort(__obf_821d67ba57c19d43)
	__obf_2cc5392205ea6f7e.__obf_f8dad1fa545d0b9c = __obf_821d67ba57c19d43
}

func __obf_89f065dcc001902b[M Member](__obf_b44aef5e8423bc09 []M, __obf_b4f280cad7f00673 M) bool {
	for _, __obf_5012e0ddf72cce0d := range __obf_b44aef5e8423bc09 {
		if __obf_5012e0ddf72cce0d.String() == __obf_b4f280cad7f00673.String() {
			return true
		}
	}
	return false
}
