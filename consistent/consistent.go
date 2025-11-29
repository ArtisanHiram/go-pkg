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
package __obf_9be990879c21f0f5 // import "stathat.com/c/consistent"

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

type __obf_4ff6e007d0ade68a []uint32

// Len returns the length of the uints array.
func (__obf_4a9efc42239e810b __obf_4ff6e007d0ade68a) Len() int { return len(__obf_4a9efc42239e810b) }

// Less returns true if element i is less than element j.
func (__obf_4a9efc42239e810b __obf_4ff6e007d0ade68a) Less(__obf_be4841bc046b7409, __obf_803adb4c6b818c7f int) bool {
	return __obf_4a9efc42239e810b[__obf_be4841bc046b7409] < __obf_4a9efc42239e810b[__obf_803adb4c6b818c7f]
}

// Swap exchanges elements i and j.
func (__obf_4a9efc42239e810b __obf_4ff6e007d0ade68a) Swap(__obf_be4841bc046b7409, __obf_803adb4c6b818c7f int) {
	__obf_4a9efc42239e810b[__obf_be4841bc046b7409], __obf_4a9efc42239e810b[__obf_803adb4c6b818c7f] = __obf_4a9efc42239e810b[__obf_803adb4c6b818c7f], __obf_4a9efc42239e810b[__obf_be4841bc046b7409]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_ea72712c57531268 map[uint32]M
	__obf_6bd17618b08a636d map[string]bool
	__obf_0ecc6febb4ce9c06 __obf_4ff6e007d0ade68a

	NumberOfReplicas       int
	__obf_012d48a0f6aceb42 int64
	__obf_bf809ed7c0664b97 [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_1b12a052f41c70af := new(Consistent[M])
	__obf_1b12a052f41c70af.
		NumberOfReplicas = 20
	__obf_1b12a052f41c70af.__obf_ea72712c57531268 = make(map[uint32]M)
	__obf_1b12a052f41c70af.__obf_6bd17618b08a636d = make(map[string]bool)
	return __obf_1b12a052f41c70af
}

// eltKey generates a string key for an element with an index.
func (__obf_1b12a052f41c70af *Consistent[M]) __obf_c2a44d6cb9d090de(__obf_718c74ba293c0403 string, __obf_e378a41570b13b52 int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_e378a41570b13b52) + __obf_718c74ba293c0403
}

// Add inserts a string element in the consistent hash.
func (__obf_1b12a052f41c70af *Consistent[M]) Add(__obf_718c74ba293c0403 M) {
	__obf_1b12a052f41c70af.
		Lock()
	defer __obf_1b12a052f41c70af.Unlock()
	__obf_1b12a052f41c70af.

		// need c.Lock() before calling
		__obf_daf4e94362cc1ff7(__obf_718c74ba293c0403)
}

func (__obf_1b12a052f41c70af *Consistent[M]) __obf_daf4e94362cc1ff7(__obf_718c74ba293c0403 M) {
	for __obf_be4841bc046b7409 := 0; __obf_be4841bc046b7409 < __obf_1b12a052f41c70af.NumberOfReplicas; __obf_be4841bc046b7409++ {
		__obf_1b12a052f41c70af.__obf_ea72712c57531268[__obf_1b12a052f41c70af.__obf_6024c04436db7272(__obf_1b12a052f41c70af.__obf_c2a44d6cb9d090de(__obf_718c74ba293c0403.String(), __obf_be4841bc046b7409))] = __obf_718c74ba293c0403
	}
	__obf_1b12a052f41c70af.__obf_6bd17618b08a636d[__obf_718c74ba293c0403.String()] = true
	__obf_1b12a052f41c70af.__obf_a6d070f3962c7614()
	__obf_1b12a052f41c70af.

		// Remove removes an element from the hash.
		__obf_012d48a0f6aceb42++
}

func (__obf_1b12a052f41c70af *Consistent[M]) Remove(__obf_718c74ba293c0403 M) {
	__obf_1b12a052f41c70af.
		Lock()
	defer __obf_1b12a052f41c70af.Unlock()
	__obf_1b12a052f41c70af.__obf_10efe00a7b71a486(__obf_718c74ba293c0403.String())
}

// need c.Lock() before calling
func (__obf_1b12a052f41c70af *Consistent[M]) __obf_10efe00a7b71a486(__obf_718c74ba293c0403 string) {
	for __obf_be4841bc046b7409 := 0; __obf_be4841bc046b7409 < __obf_1b12a052f41c70af.NumberOfReplicas; __obf_be4841bc046b7409++ {
		delete(__obf_1b12a052f41c70af.__obf_ea72712c57531268, __obf_1b12a052f41c70af.__obf_6024c04436db7272(__obf_1b12a052f41c70af.__obf_c2a44d6cb9d090de(__obf_718c74ba293c0403, __obf_be4841bc046b7409)))
	}
	delete(__obf_1b12a052f41c70af.__obf_6bd17618b08a636d, __obf_718c74ba293c0403)
	__obf_1b12a052f41c70af.__obf_a6d070f3962c7614()
	__obf_1b12a052f41c70af.

		// Set sets all the elements in the hash.  If there are existing elements not
		// present in elts, they will be removed.
		__obf_012d48a0f6aceb42--
}

func (__obf_1b12a052f41c70af *Consistent[M]) Set(__obf_408a9f8bceece723 []M) {
	__obf_1b12a052f41c70af.
		Lock()
	defer __obf_1b12a052f41c70af.Unlock()
	for __obf_f91d42de35549de5 := range __obf_1b12a052f41c70af.__obf_6bd17618b08a636d {
		__obf_71dd4b108537d5ef := false
		for _, __obf_17ad8dd03a4094bc := range __obf_408a9f8bceece723 {
			if __obf_f91d42de35549de5 == __obf_17ad8dd03a4094bc.String() {
				__obf_71dd4b108537d5ef = true
				break
			}
		}
		if !__obf_71dd4b108537d5ef {
			__obf_1b12a052f41c70af.__obf_10efe00a7b71a486(__obf_f91d42de35549de5)
		}
	}
	for _, __obf_17ad8dd03a4094bc := range __obf_408a9f8bceece723 {
		_, __obf_e5c842ebd893d90b := __obf_1b12a052f41c70af.__obf_6bd17618b08a636d[__obf_17ad8dd03a4094bc.String()]
		if __obf_e5c842ebd893d90b {
			continue
		}
		__obf_1b12a052f41c70af.__obf_daf4e94362cc1ff7(__obf_17ad8dd03a4094bc)
	}
}

func (__obf_1b12a052f41c70af *Consistent[M]) Members() []string {
	__obf_1b12a052f41c70af.
		RLock()
	defer __obf_1b12a052f41c70af.RUnlock()
	var __obf_fa0edd2e7be81b10 []string
	for __obf_f91d42de35549de5 := range __obf_1b12a052f41c70af.__obf_6bd17618b08a636d {
		__obf_fa0edd2e7be81b10 = append(__obf_fa0edd2e7be81b10,

			// Get returns an element close to where name hashes to in the circle.
			__obf_f91d42de35549de5)
	}
	return __obf_fa0edd2e7be81b10
}

func (__obf_1b12a052f41c70af *Consistent[M]) Get(__obf_f580e7ad61b6d51d string) (__obf_47af78e7a654ee0d M, __obf_50a8d4a1a88c1b23 error) {
	__obf_1b12a052f41c70af.
		RLock()
	defer __obf_1b12a052f41c70af.RUnlock()
	if len(__obf_1b12a052f41c70af.__obf_ea72712c57531268) == 0 {
		__obf_50a8d4a1a88c1b23 = ErrEmptyCircle
		return
	}
	__obf_0ea9fdd2182c6f78 := __obf_1b12a052f41c70af.__obf_6024c04436db7272(__obf_f580e7ad61b6d51d)
	__obf_be4841bc046b7409 := __obf_1b12a052f41c70af.__obf_d45447c909083df3(__obf_0ea9fdd2182c6f78)
	__obf_47af78e7a654ee0d = __obf_1b12a052f41c70af.__obf_ea72712c57531268[__obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06[__obf_be4841bc046b7409]]
	return
}

func (__obf_1b12a052f41c70af *Consistent[M]) __obf_d45447c909083df3(__obf_0ea9fdd2182c6f78 uint32) (__obf_be4841bc046b7409 int) {
	__obf_54588510064453a5 := func(__obf_4a9efc42239e810b int) bool {
		return __obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06[__obf_4a9efc42239e810b] > __obf_0ea9fdd2182c6f78
	}
	__obf_be4841bc046b7409 = sort.Search(len(__obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06), __obf_54588510064453a5)
	if __obf_be4841bc046b7409 >= len(__obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06) {
		__obf_be4841bc046b7409 = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_1b12a052f41c70af *Consistent[M]) GetTwo(__obf_f580e7ad61b6d51d string) (__obf_b93b3da85b6b9f9d M, __obf_52ddc059ee9b5ecc M, __obf_50a8d4a1a88c1b23 error) {
	__obf_1b12a052f41c70af.
		RLock()
	defer __obf_1b12a052f41c70af.RUnlock()
	if len(__obf_1b12a052f41c70af.__obf_ea72712c57531268) == 0 {
		__obf_50a8d4a1a88c1b23 = ErrEmptyCircle
		return
	}
	__obf_0ea9fdd2182c6f78 := __obf_1b12a052f41c70af.__obf_6024c04436db7272(__obf_f580e7ad61b6d51d)
	__obf_be4841bc046b7409 := __obf_1b12a052f41c70af.__obf_d45447c909083df3(__obf_0ea9fdd2182c6f78)
	__obf_b93b3da85b6b9f9d = __obf_1b12a052f41c70af.__obf_ea72712c57531268[__obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06[__obf_be4841bc046b7409]]

	if __obf_1b12a052f41c70af.__obf_012d48a0f6aceb42 == 1 {
		return
	}
	__obf_26c3deba3d228fbf := __obf_be4841bc046b7409
	for __obf_be4841bc046b7409 = __obf_26c3deba3d228fbf + 1; __obf_be4841bc046b7409 != __obf_26c3deba3d228fbf; __obf_be4841bc046b7409++ {
		if __obf_be4841bc046b7409 >= len(__obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06) {
			__obf_be4841bc046b7409 = 0
		}
		__obf_52ddc059ee9b5ecc = __obf_1b12a052f41c70af.__obf_ea72712c57531268[__obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06[__obf_be4841bc046b7409]]
		if __obf_52ddc059ee9b5ecc.String() != __obf_b93b3da85b6b9f9d.String() {
			break
		}
	}
	return __obf_b93b3da85b6b9f9d,

		// GetN returns the N closest distinct elements to the name input in the circle.
		__obf_52ddc059ee9b5ecc, nil
}

func (__obf_1b12a052f41c70af *Consistent[M]) GetN(__obf_f580e7ad61b6d51d string, __obf_45b5f557da5fa9ca int) (__obf_47af78e7a654ee0d []M, __obf_50a8d4a1a88c1b23 error) {
	__obf_1b12a052f41c70af.
		RLock()
	defer __obf_1b12a052f41c70af.RUnlock()

	if len(__obf_1b12a052f41c70af.__obf_ea72712c57531268) == 0 {
		__obf_50a8d4a1a88c1b23 = ErrEmptyCircle
		return
	}

	if __obf_1b12a052f41c70af.__obf_012d48a0f6aceb42 < int64(__obf_45b5f557da5fa9ca) {
		__obf_45b5f557da5fa9ca = int(__obf_1b12a052f41c70af.__obf_012d48a0f6aceb42)
	}

	var (
		__obf_0ea9fdd2182c6f78 = __obf_1b12a052f41c70af.__obf_6024c04436db7272(__obf_f580e7ad61b6d51d)
		__obf_be4841bc046b7409 = __obf_1b12a052f41c70af.__obf_d45447c909083df3(__obf_0ea9fdd2182c6f78)
		__obf_26c3deba3d228fbf = __obf_be4841bc046b7409
		__obf_dfcec79e8fc3a1c9 = __obf_1b12a052f41c70af.__obf_ea72712c57531268[__obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06[__obf_be4841bc046b7409]]
	)
	__obf_47af78e7a654ee0d = make([]M, 0, __obf_45b5f557da5fa9ca)
	__obf_47af78e7a654ee0d = append(__obf_47af78e7a654ee0d, __obf_dfcec79e8fc3a1c9)

	if len(__obf_47af78e7a654ee0d) == __obf_45b5f557da5fa9ca {
		return __obf_47af78e7a654ee0d, nil
	}

	for __obf_be4841bc046b7409 = __obf_26c3deba3d228fbf + 1; __obf_be4841bc046b7409 != __obf_26c3deba3d228fbf; __obf_be4841bc046b7409++ {
		if __obf_be4841bc046b7409 >= len(__obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06) {
			__obf_be4841bc046b7409 = 0
		}
		__obf_dfcec79e8fc3a1c9 = __obf_1b12a052f41c70af.__obf_ea72712c57531268[__obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06[__obf_be4841bc046b7409]]
		if !__obf_86e281cf8cd86033(__obf_47af78e7a654ee0d, __obf_dfcec79e8fc3a1c9) {
			__obf_47af78e7a654ee0d = append(__obf_47af78e7a654ee0d, __obf_dfcec79e8fc3a1c9)
		}
		if len(__obf_47af78e7a654ee0d) == __obf_45b5f557da5fa9ca {
			break
		}
	}

	return __obf_47af78e7a654ee0d, nil
}

func (__obf_1b12a052f41c70af *Consistent[M]) __obf_6024c04436db7272(__obf_0ea9fdd2182c6f78 string) uint32 {
	if __obf_1b12a052f41c70af.UseFnv {
		return __obf_1b12a052f41c70af.__obf_da18db77131f92f1(__obf_0ea9fdd2182c6f78)
	}
	return __obf_1b12a052f41c70af.__obf_de411f5c18a58a2f(__obf_0ea9fdd2182c6f78)
}

func (__obf_1b12a052f41c70af *Consistent[M]) __obf_de411f5c18a58a2f(__obf_0ea9fdd2182c6f78 string) uint32 {
	if len(__obf_0ea9fdd2182c6f78) < 64 {
		var __obf_bf809ed7c0664b97 [64]byte
		copy(__obf_bf809ed7c0664b97[:], __obf_0ea9fdd2182c6f78)
		return crc32.ChecksumIEEE(__obf_bf809ed7c0664b97[:len(__obf_0ea9fdd2182c6f78)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_0ea9fdd2182c6f78))
}

func (__obf_1b12a052f41c70af *Consistent[M]) __obf_da18db77131f92f1(__obf_0ea9fdd2182c6f78 string) uint32 {
	__obf_e920fad6e4f5ccb0 := fnv.New32a()
	__obf_e920fad6e4f5ccb0.
		Write([]byte(__obf_0ea9fdd2182c6f78))
	return __obf_e920fad6e4f5ccb0.Sum32()
}

func (__obf_1b12a052f41c70af *Consistent[M]) __obf_a6d070f3962c7614() {
	__obf_37fc2c16f5e7ef83 := __obf_1b12a052f41c70af.
		//reallocate if we're holding on to too much (1/4th)
		__obf_0ecc6febb4ce9c06[:0]

	if cap(__obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06)/(__obf_1b12a052f41c70af.NumberOfReplicas*4) > len(__obf_1b12a052f41c70af.__obf_ea72712c57531268) {
		__obf_37fc2c16f5e7ef83 = nil
	}
	for __obf_f91d42de35549de5 := range __obf_1b12a052f41c70af.__obf_ea72712c57531268 {
		__obf_37fc2c16f5e7ef83 = append(__obf_37fc2c16f5e7ef83, __obf_f91d42de35549de5)
	}
	sort.Sort(__obf_37fc2c16f5e7ef83)
	__obf_1b12a052f41c70af.__obf_0ecc6febb4ce9c06 = __obf_37fc2c16f5e7ef83
}

func __obf_86e281cf8cd86033[M Member](__obf_7e22552c857cfb3b []M, __obf_72d51fa7b0537a4f M) bool {
	for _, __obf_fa0edd2e7be81b10 := range __obf_7e22552c857cfb3b {
		if __obf_fa0edd2e7be81b10.String() == __obf_72d51fa7b0537a4f.String() {
			return true
		}
	}
	return false
}
