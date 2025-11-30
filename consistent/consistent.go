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
package __obf_9277bffea8420bd4 // import "stathat.com/c/consistent"

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

type __obf_68d50b2236191dc3 []uint32

// Len returns the length of the uints array.
func (__obf_228581d32283deb7 __obf_68d50b2236191dc3) Len() int { return len(__obf_228581d32283deb7) }

// Less returns true if element i is less than element j.
func (__obf_228581d32283deb7 __obf_68d50b2236191dc3) Less(__obf_913b2a5190233a37, __obf_2e8a081074c9dcab int) bool {
	return __obf_228581d32283deb7[__obf_913b2a5190233a37] < __obf_228581d32283deb7[__obf_2e8a081074c9dcab]
}

// Swap exchanges elements i and j.
func (__obf_228581d32283deb7 __obf_68d50b2236191dc3) Swap(__obf_913b2a5190233a37, __obf_2e8a081074c9dcab int) {
	__obf_228581d32283deb7[__obf_913b2a5190233a37], __obf_228581d32283deb7[__obf_2e8a081074c9dcab] = __obf_228581d32283deb7[__obf_2e8a081074c9dcab], __obf_228581d32283deb7[__obf_913b2a5190233a37]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_7bb82b47b88dcb5d map[uint32]M
	__obf_6e41b50a2e5ea414 map[string]bool
	__obf_7c4698fb08cd6e25 __obf_68d50b2236191dc3

	NumberOfReplicas       int
	__obf_d1c16c89938deed3 int64
	__obf_4f579cbaef60e2de [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_d0e79efaea5a2ff2 := new(Consistent[M])
	__obf_d0e79efaea5a2ff2.
		NumberOfReplicas = 20
	__obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d = make(map[uint32]M)
	__obf_d0e79efaea5a2ff2.__obf_6e41b50a2e5ea414 = make(map[string]bool)
	return __obf_d0e79efaea5a2ff2
}

// eltKey generates a string key for an element with an index.
func (__obf_d0e79efaea5a2ff2 *Consistent[M]) __obf_b8643669311aee06(__obf_467d2d19232edb24 string, __obf_b1f7477195560652 int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_b1f7477195560652) + __obf_467d2d19232edb24
}

// Add inserts a string element in the consistent hash.
func (__obf_d0e79efaea5a2ff2 *Consistent[M]) Add(__obf_467d2d19232edb24 M) {
	__obf_d0e79efaea5a2ff2.
		Lock()
	defer __obf_d0e79efaea5a2ff2.Unlock()
	__obf_d0e79efaea5a2ff2.

		// need c.Lock() before calling
		__obf_9290ca2626a7350b(__obf_467d2d19232edb24)
}

func (__obf_d0e79efaea5a2ff2 *Consistent[M]) __obf_9290ca2626a7350b(__obf_467d2d19232edb24 M) {
	for __obf_913b2a5190233a37 := 0; __obf_913b2a5190233a37 < __obf_d0e79efaea5a2ff2.NumberOfReplicas; __obf_913b2a5190233a37++ {
		__obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d[__obf_d0e79efaea5a2ff2.__obf_a23735341aef034f(__obf_d0e79efaea5a2ff2.__obf_b8643669311aee06(__obf_467d2d19232edb24.String(), __obf_913b2a5190233a37))] = __obf_467d2d19232edb24
	}
	__obf_d0e79efaea5a2ff2.__obf_6e41b50a2e5ea414[__obf_467d2d19232edb24.String()] = true
	__obf_d0e79efaea5a2ff2.__obf_6789f4d8e0d3910e()
	__obf_d0e79efaea5a2ff2.

		// Remove removes an element from the hash.
		__obf_d1c16c89938deed3++
}

func (__obf_d0e79efaea5a2ff2 *Consistent[M]) Remove(__obf_467d2d19232edb24 M) {
	__obf_d0e79efaea5a2ff2.
		Lock()
	defer __obf_d0e79efaea5a2ff2.Unlock()
	__obf_d0e79efaea5a2ff2.__obf_ffbc34e324c0fa13(__obf_467d2d19232edb24.String())
}

// need c.Lock() before calling
func (__obf_d0e79efaea5a2ff2 *Consistent[M]) __obf_ffbc34e324c0fa13(__obf_467d2d19232edb24 string) {
	for __obf_913b2a5190233a37 := 0; __obf_913b2a5190233a37 < __obf_d0e79efaea5a2ff2.NumberOfReplicas; __obf_913b2a5190233a37++ {
		delete(__obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d, __obf_d0e79efaea5a2ff2.__obf_a23735341aef034f(__obf_d0e79efaea5a2ff2.__obf_b8643669311aee06(__obf_467d2d19232edb24, __obf_913b2a5190233a37)))
	}
	delete(__obf_d0e79efaea5a2ff2.__obf_6e41b50a2e5ea414, __obf_467d2d19232edb24)
	__obf_d0e79efaea5a2ff2.__obf_6789f4d8e0d3910e()
	__obf_d0e79efaea5a2ff2.

		// Set sets all the elements in the hash.  If there are existing elements not
		// present in elts, they will be removed.
		__obf_d1c16c89938deed3--
}

func (__obf_d0e79efaea5a2ff2 *Consistent[M]) Set(__obf_932218e1b1835803 []M) {
	__obf_d0e79efaea5a2ff2.
		Lock()
	defer __obf_d0e79efaea5a2ff2.Unlock()
	for __obf_676630c6e2242fed := range __obf_d0e79efaea5a2ff2.__obf_6e41b50a2e5ea414 {
		__obf_287e6d20bce04316 := false
		for _, __obf_0bfa4f923de8e0c0 := range __obf_932218e1b1835803 {
			if __obf_676630c6e2242fed == __obf_0bfa4f923de8e0c0.String() {
				__obf_287e6d20bce04316 = true
				break
			}
		}
		if !__obf_287e6d20bce04316 {
			__obf_d0e79efaea5a2ff2.__obf_ffbc34e324c0fa13(__obf_676630c6e2242fed)
		}
	}
	for _, __obf_0bfa4f923de8e0c0 := range __obf_932218e1b1835803 {
		_, __obf_c179880452e74fd4 := __obf_d0e79efaea5a2ff2.__obf_6e41b50a2e5ea414[__obf_0bfa4f923de8e0c0.String()]
		if __obf_c179880452e74fd4 {
			continue
		}
		__obf_d0e79efaea5a2ff2.__obf_9290ca2626a7350b(__obf_0bfa4f923de8e0c0)
	}
}

func (__obf_d0e79efaea5a2ff2 *Consistent[M]) Members() []string {
	__obf_d0e79efaea5a2ff2.
		RLock()
	defer __obf_d0e79efaea5a2ff2.RUnlock()
	var __obf_6fd254ac00619146 []string
	for __obf_676630c6e2242fed := range __obf_d0e79efaea5a2ff2.__obf_6e41b50a2e5ea414 {
		__obf_6fd254ac00619146 = append(__obf_6fd254ac00619146,

			// Get returns an element close to where name hashes to in the circle.
			__obf_676630c6e2242fed)
	}
	return __obf_6fd254ac00619146
}

func (__obf_d0e79efaea5a2ff2 *Consistent[M]) Get(__obf_bf266d4d4903d78b string) (__obf_69de6d6c3c5e18b0 M, __obf_ef58c2a9a7e5f2c6 error) {
	__obf_d0e79efaea5a2ff2.
		RLock()
	defer __obf_d0e79efaea5a2ff2.RUnlock()
	if len(__obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d) == 0 {
		__obf_ef58c2a9a7e5f2c6 = ErrEmptyCircle
		return
	}
	__obf_113a9a94b9d72957 := __obf_d0e79efaea5a2ff2.__obf_a23735341aef034f(__obf_bf266d4d4903d78b)
	__obf_913b2a5190233a37 := __obf_d0e79efaea5a2ff2.__obf_e3a758e2669db3c8(__obf_113a9a94b9d72957)
	__obf_69de6d6c3c5e18b0 = __obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d[__obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25[__obf_913b2a5190233a37]]
	return
}

func (__obf_d0e79efaea5a2ff2 *Consistent[M]) __obf_e3a758e2669db3c8(__obf_113a9a94b9d72957 uint32) (__obf_913b2a5190233a37 int) {
	__obf_6495762012ca18fc := func(__obf_228581d32283deb7 int) bool {
		return __obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25[__obf_228581d32283deb7] > __obf_113a9a94b9d72957
	}
	__obf_913b2a5190233a37 = sort.Search(len(__obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25), __obf_6495762012ca18fc)
	if __obf_913b2a5190233a37 >= len(__obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25) {
		__obf_913b2a5190233a37 = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_d0e79efaea5a2ff2 *Consistent[M]) GetTwo(__obf_bf266d4d4903d78b string) (__obf_15342c54f580ab95 M, __obf_53d50ee6ea4a838e M, __obf_ef58c2a9a7e5f2c6 error) {
	__obf_d0e79efaea5a2ff2.
		RLock()
	defer __obf_d0e79efaea5a2ff2.RUnlock()
	if len(__obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d) == 0 {
		__obf_ef58c2a9a7e5f2c6 = ErrEmptyCircle
		return
	}
	__obf_113a9a94b9d72957 := __obf_d0e79efaea5a2ff2.__obf_a23735341aef034f(__obf_bf266d4d4903d78b)
	__obf_913b2a5190233a37 := __obf_d0e79efaea5a2ff2.__obf_e3a758e2669db3c8(__obf_113a9a94b9d72957)
	__obf_15342c54f580ab95 = __obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d[__obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25[__obf_913b2a5190233a37]]

	if __obf_d0e79efaea5a2ff2.__obf_d1c16c89938deed3 == 1 {
		return
	}
	__obf_2bfdb67479ce8179 := __obf_913b2a5190233a37
	for __obf_913b2a5190233a37 = __obf_2bfdb67479ce8179 + 1; __obf_913b2a5190233a37 != __obf_2bfdb67479ce8179; __obf_913b2a5190233a37++ {
		if __obf_913b2a5190233a37 >= len(__obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25) {
			__obf_913b2a5190233a37 = 0
		}
		__obf_53d50ee6ea4a838e = __obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d[__obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25[__obf_913b2a5190233a37]]
		if __obf_53d50ee6ea4a838e.String() != __obf_15342c54f580ab95.String() {
			break
		}
	}
	return __obf_15342c54f580ab95,

		// GetN returns the N closest distinct elements to the name input in the circle.
		__obf_53d50ee6ea4a838e, nil
}

func (__obf_d0e79efaea5a2ff2 *Consistent[M]) GetN(__obf_bf266d4d4903d78b string, __obf_a817f467b95994f1 int) (__obf_69de6d6c3c5e18b0 []M, __obf_ef58c2a9a7e5f2c6 error) {
	__obf_d0e79efaea5a2ff2.
		RLock()
	defer __obf_d0e79efaea5a2ff2.RUnlock()

	if len(__obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d) == 0 {
		__obf_ef58c2a9a7e5f2c6 = ErrEmptyCircle
		return
	}

	if __obf_d0e79efaea5a2ff2.__obf_d1c16c89938deed3 < int64(__obf_a817f467b95994f1) {
		__obf_a817f467b95994f1 = int(__obf_d0e79efaea5a2ff2.__obf_d1c16c89938deed3)
	}

	var (
		__obf_113a9a94b9d72957 = __obf_d0e79efaea5a2ff2.__obf_a23735341aef034f(__obf_bf266d4d4903d78b)
		__obf_913b2a5190233a37 = __obf_d0e79efaea5a2ff2.__obf_e3a758e2669db3c8(__obf_113a9a94b9d72957)
		__obf_2bfdb67479ce8179 = __obf_913b2a5190233a37
		__obf_c991381d23761b9c = __obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d[__obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25[__obf_913b2a5190233a37]]
	)
	__obf_69de6d6c3c5e18b0 = make([]M, 0, __obf_a817f467b95994f1)
	__obf_69de6d6c3c5e18b0 = append(__obf_69de6d6c3c5e18b0, __obf_c991381d23761b9c)

	if len(__obf_69de6d6c3c5e18b0) == __obf_a817f467b95994f1 {
		return __obf_69de6d6c3c5e18b0, nil
	}

	for __obf_913b2a5190233a37 = __obf_2bfdb67479ce8179 + 1; __obf_913b2a5190233a37 != __obf_2bfdb67479ce8179; __obf_913b2a5190233a37++ {
		if __obf_913b2a5190233a37 >= len(__obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25) {
			__obf_913b2a5190233a37 = 0
		}
		__obf_c991381d23761b9c = __obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d[__obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25[__obf_913b2a5190233a37]]
		if !__obf_f7504a0e2c0e8394(__obf_69de6d6c3c5e18b0, __obf_c991381d23761b9c) {
			__obf_69de6d6c3c5e18b0 = append(__obf_69de6d6c3c5e18b0, __obf_c991381d23761b9c)
		}
		if len(__obf_69de6d6c3c5e18b0) == __obf_a817f467b95994f1 {
			break
		}
	}

	return __obf_69de6d6c3c5e18b0, nil
}

func (__obf_d0e79efaea5a2ff2 *Consistent[M]) __obf_a23735341aef034f(__obf_113a9a94b9d72957 string) uint32 {
	if __obf_d0e79efaea5a2ff2.UseFnv {
		return __obf_d0e79efaea5a2ff2.__obf_419b32a25a7bf89a(__obf_113a9a94b9d72957)
	}
	return __obf_d0e79efaea5a2ff2.__obf_5dadcf56b2c56497(__obf_113a9a94b9d72957)
}

func (__obf_d0e79efaea5a2ff2 *Consistent[M]) __obf_5dadcf56b2c56497(__obf_113a9a94b9d72957 string) uint32 {
	if len(__obf_113a9a94b9d72957) < 64 {
		var __obf_4f579cbaef60e2de [64]byte
		copy(__obf_4f579cbaef60e2de[:], __obf_113a9a94b9d72957)
		return crc32.ChecksumIEEE(__obf_4f579cbaef60e2de[:len(__obf_113a9a94b9d72957)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_113a9a94b9d72957))
}

func (__obf_d0e79efaea5a2ff2 *Consistent[M]) __obf_419b32a25a7bf89a(__obf_113a9a94b9d72957 string) uint32 {
	__obf_7a84944abdcf4a3b := fnv.New32a()
	__obf_7a84944abdcf4a3b.
		Write([]byte(__obf_113a9a94b9d72957))
	return __obf_7a84944abdcf4a3b.Sum32()
}

func (__obf_d0e79efaea5a2ff2 *Consistent[M]) __obf_6789f4d8e0d3910e() {
	__obf_930950fe17492e7a := __obf_d0e79efaea5a2ff2.
		//reallocate if we're holding on to too much (1/4th)
		__obf_7c4698fb08cd6e25[:0]

	if cap(__obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25)/(__obf_d0e79efaea5a2ff2.NumberOfReplicas*4) > len(__obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d) {
		__obf_930950fe17492e7a = nil
	}
	for __obf_676630c6e2242fed := range __obf_d0e79efaea5a2ff2.__obf_7bb82b47b88dcb5d {
		__obf_930950fe17492e7a = append(__obf_930950fe17492e7a, __obf_676630c6e2242fed)
	}
	sort.Sort(__obf_930950fe17492e7a)
	__obf_d0e79efaea5a2ff2.__obf_7c4698fb08cd6e25 = __obf_930950fe17492e7a
}

func __obf_f7504a0e2c0e8394[M Member](__obf_748cbce3a9073095 []M, __obf_90ed8dd04b44da60 M) bool {
	for _, __obf_6fd254ac00619146 := range __obf_748cbce3a9073095 {
		if __obf_6fd254ac00619146.String() == __obf_90ed8dd04b44da60.String() {
			return true
		}
	}
	return false
}
