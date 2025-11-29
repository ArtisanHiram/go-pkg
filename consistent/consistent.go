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
package __obf_e372f6d452d8c255 // import "stathat.com/c/consistent"

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

type __obf_9581786378fdb14c []uint32

// Len returns the length of the uints array.
func (__obf_da31fab042146cb1 __obf_9581786378fdb14c) Len() int { return len(__obf_da31fab042146cb1) }

// Less returns true if element i is less than element j.
func (__obf_da31fab042146cb1 __obf_9581786378fdb14c) Less(__obf_12d456a52dd0fa73, __obf_779e24c2512009b1 int) bool {
	return __obf_da31fab042146cb1[__obf_12d456a52dd0fa73] < __obf_da31fab042146cb1[__obf_779e24c2512009b1]
}

// Swap exchanges elements i and j.
func (__obf_da31fab042146cb1 __obf_9581786378fdb14c) Swap(__obf_12d456a52dd0fa73, __obf_779e24c2512009b1 int) {
	__obf_da31fab042146cb1[__obf_12d456a52dd0fa73], __obf_da31fab042146cb1[__obf_779e24c2512009b1] = __obf_da31fab042146cb1[__obf_779e24c2512009b1], __obf_da31fab042146cb1[__obf_12d456a52dd0fa73]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_d3b7e2fbbbb195b6 map[uint32]M
	__obf_85674369c1accc7e map[string]bool
	__obf_27e0d051531f1b97 __obf_9581786378fdb14c

	NumberOfReplicas       int
	__obf_1dd9a0f958f3655a int64
	__obf_adfc779ee349fcc6 [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_0dc283c37429c9a2 := new(Consistent[M])
	__obf_0dc283c37429c9a2.
		NumberOfReplicas = 20
	__obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6 = make(map[uint32]M)
	__obf_0dc283c37429c9a2.__obf_85674369c1accc7e = make(map[string]bool)
	return __obf_0dc283c37429c9a2
}

// eltKey generates a string key for an element with an index.
func (__obf_0dc283c37429c9a2 *Consistent[M]) __obf_65668e3cbc59c904(__obf_f426cd329c8bc4ca string, __obf_169345b4a30d7b48 int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_169345b4a30d7b48) + __obf_f426cd329c8bc4ca
}

// Add inserts a string element in the consistent hash.
func (__obf_0dc283c37429c9a2 *Consistent[M]) Add(__obf_f426cd329c8bc4ca M) {
	__obf_0dc283c37429c9a2.
		Lock()
	defer __obf_0dc283c37429c9a2.Unlock()
	__obf_0dc283c37429c9a2.

		// need c.Lock() before calling
		__obf_5212d0eff70abb57(__obf_f426cd329c8bc4ca)
}

func (__obf_0dc283c37429c9a2 *Consistent[M]) __obf_5212d0eff70abb57(__obf_f426cd329c8bc4ca M) {
	for __obf_12d456a52dd0fa73 := 0; __obf_12d456a52dd0fa73 < __obf_0dc283c37429c9a2.NumberOfReplicas; __obf_12d456a52dd0fa73++ {
		__obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6[__obf_0dc283c37429c9a2.__obf_6fe9ea8987a7e652(__obf_0dc283c37429c9a2.__obf_65668e3cbc59c904(__obf_f426cd329c8bc4ca.String(), __obf_12d456a52dd0fa73))] = __obf_f426cd329c8bc4ca
	}
	__obf_0dc283c37429c9a2.__obf_85674369c1accc7e[__obf_f426cd329c8bc4ca.String()] = true
	__obf_0dc283c37429c9a2.__obf_baaa01c5a4bbbc64()
	__obf_0dc283c37429c9a2.

		// Remove removes an element from the hash.
		__obf_1dd9a0f958f3655a++
}

func (__obf_0dc283c37429c9a2 *Consistent[M]) Remove(__obf_f426cd329c8bc4ca M) {
	__obf_0dc283c37429c9a2.
		Lock()
	defer __obf_0dc283c37429c9a2.Unlock()
	__obf_0dc283c37429c9a2.__obf_9b920888e935efa5(__obf_f426cd329c8bc4ca.String())
}

// need c.Lock() before calling
func (__obf_0dc283c37429c9a2 *Consistent[M]) __obf_9b920888e935efa5(__obf_f426cd329c8bc4ca string) {
	for __obf_12d456a52dd0fa73 := 0; __obf_12d456a52dd0fa73 < __obf_0dc283c37429c9a2.NumberOfReplicas; __obf_12d456a52dd0fa73++ {
		delete(__obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6, __obf_0dc283c37429c9a2.__obf_6fe9ea8987a7e652(__obf_0dc283c37429c9a2.__obf_65668e3cbc59c904(__obf_f426cd329c8bc4ca, __obf_12d456a52dd0fa73)))
	}
	delete(__obf_0dc283c37429c9a2.__obf_85674369c1accc7e, __obf_f426cd329c8bc4ca)
	__obf_0dc283c37429c9a2.__obf_baaa01c5a4bbbc64()
	__obf_0dc283c37429c9a2.

		// Set sets all the elements in the hash.  If there are existing elements not
		// present in elts, they will be removed.
		__obf_1dd9a0f958f3655a--
}

func (__obf_0dc283c37429c9a2 *Consistent[M]) Set(__obf_cc91ec6b8be5a64a []M) {
	__obf_0dc283c37429c9a2.
		Lock()
	defer __obf_0dc283c37429c9a2.Unlock()
	for __obf_068640d7e6c7dd1c := range __obf_0dc283c37429c9a2.__obf_85674369c1accc7e {
		__obf_f514f80953d31ffd := false
		for _, __obf_4b61a0b74d31e850 := range __obf_cc91ec6b8be5a64a {
			if __obf_068640d7e6c7dd1c == __obf_4b61a0b74d31e850.String() {
				__obf_f514f80953d31ffd = true
				break
			}
		}
		if !__obf_f514f80953d31ffd {
			__obf_0dc283c37429c9a2.__obf_9b920888e935efa5(__obf_068640d7e6c7dd1c)
		}
	}
	for _, __obf_4b61a0b74d31e850 := range __obf_cc91ec6b8be5a64a {
		_, __obf_10150f36154b3ca6 := __obf_0dc283c37429c9a2.__obf_85674369c1accc7e[__obf_4b61a0b74d31e850.String()]
		if __obf_10150f36154b3ca6 {
			continue
		}
		__obf_0dc283c37429c9a2.__obf_5212d0eff70abb57(__obf_4b61a0b74d31e850)
	}
}

func (__obf_0dc283c37429c9a2 *Consistent[M]) Members() []string {
	__obf_0dc283c37429c9a2.
		RLock()
	defer __obf_0dc283c37429c9a2.RUnlock()
	var __obf_842e0740cba494d9 []string
	for __obf_068640d7e6c7dd1c := range __obf_0dc283c37429c9a2.__obf_85674369c1accc7e {
		__obf_842e0740cba494d9 = append(__obf_842e0740cba494d9,

			// Get returns an element close to where name hashes to in the circle.
			__obf_068640d7e6c7dd1c)
	}
	return __obf_842e0740cba494d9
}

func (__obf_0dc283c37429c9a2 *Consistent[M]) Get(__obf_c63cb994d5c4f83e string) (__obf_095ebebc00a3e9d4 M, __obf_7d07d34c920943b0 error) {
	__obf_0dc283c37429c9a2.
		RLock()
	defer __obf_0dc283c37429c9a2.RUnlock()
	if len(__obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6) == 0 {
		__obf_7d07d34c920943b0 = ErrEmptyCircle
		return
	}
	__obf_274c443874c5ae86 := __obf_0dc283c37429c9a2.__obf_6fe9ea8987a7e652(__obf_c63cb994d5c4f83e)
	__obf_12d456a52dd0fa73 := __obf_0dc283c37429c9a2.__obf_8e3ca5edc141df4a(__obf_274c443874c5ae86)
	__obf_095ebebc00a3e9d4 = __obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6[__obf_0dc283c37429c9a2.__obf_27e0d051531f1b97[__obf_12d456a52dd0fa73]]
	return
}

func (__obf_0dc283c37429c9a2 *Consistent[M]) __obf_8e3ca5edc141df4a(__obf_274c443874c5ae86 uint32) (__obf_12d456a52dd0fa73 int) {
	__obf_58dc7566374827fe := func(__obf_da31fab042146cb1 int) bool {
		return __obf_0dc283c37429c9a2.__obf_27e0d051531f1b97[__obf_da31fab042146cb1] > __obf_274c443874c5ae86
	}
	__obf_12d456a52dd0fa73 = sort.Search(len(__obf_0dc283c37429c9a2.__obf_27e0d051531f1b97), __obf_58dc7566374827fe)
	if __obf_12d456a52dd0fa73 >= len(__obf_0dc283c37429c9a2.__obf_27e0d051531f1b97) {
		__obf_12d456a52dd0fa73 = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_0dc283c37429c9a2 *Consistent[M]) GetTwo(__obf_c63cb994d5c4f83e string) (__obf_0a92a28efb64ce3b M, __obf_47bebbd6687661e7 M, __obf_7d07d34c920943b0 error) {
	__obf_0dc283c37429c9a2.
		RLock()
	defer __obf_0dc283c37429c9a2.RUnlock()
	if len(__obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6) == 0 {
		__obf_7d07d34c920943b0 = ErrEmptyCircle
		return
	}
	__obf_274c443874c5ae86 := __obf_0dc283c37429c9a2.__obf_6fe9ea8987a7e652(__obf_c63cb994d5c4f83e)
	__obf_12d456a52dd0fa73 := __obf_0dc283c37429c9a2.__obf_8e3ca5edc141df4a(__obf_274c443874c5ae86)
	__obf_0a92a28efb64ce3b = __obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6[__obf_0dc283c37429c9a2.__obf_27e0d051531f1b97[__obf_12d456a52dd0fa73]]

	if __obf_0dc283c37429c9a2.__obf_1dd9a0f958f3655a == 1 {
		return
	}
	__obf_47f1019286523e5e := __obf_12d456a52dd0fa73
	for __obf_12d456a52dd0fa73 = __obf_47f1019286523e5e + 1; __obf_12d456a52dd0fa73 != __obf_47f1019286523e5e; __obf_12d456a52dd0fa73++ {
		if __obf_12d456a52dd0fa73 >= len(__obf_0dc283c37429c9a2.__obf_27e0d051531f1b97) {
			__obf_12d456a52dd0fa73 = 0
		}
		__obf_47bebbd6687661e7 = __obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6[__obf_0dc283c37429c9a2.__obf_27e0d051531f1b97[__obf_12d456a52dd0fa73]]
		if __obf_47bebbd6687661e7.String() != __obf_0a92a28efb64ce3b.String() {
			break
		}
	}
	return __obf_0a92a28efb64ce3b,

		// GetN returns the N closest distinct elements to the name input in the circle.
		__obf_47bebbd6687661e7, nil
}

func (__obf_0dc283c37429c9a2 *Consistent[M]) GetN(__obf_c63cb994d5c4f83e string, __obf_9610bb14d9eebcde int) (__obf_095ebebc00a3e9d4 []M, __obf_7d07d34c920943b0 error) {
	__obf_0dc283c37429c9a2.
		RLock()
	defer __obf_0dc283c37429c9a2.RUnlock()

	if len(__obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6) == 0 {
		__obf_7d07d34c920943b0 = ErrEmptyCircle
		return
	}

	if __obf_0dc283c37429c9a2.__obf_1dd9a0f958f3655a < int64(__obf_9610bb14d9eebcde) {
		__obf_9610bb14d9eebcde = int(__obf_0dc283c37429c9a2.__obf_1dd9a0f958f3655a)
	}

	var (
		__obf_274c443874c5ae86 = __obf_0dc283c37429c9a2.__obf_6fe9ea8987a7e652(__obf_c63cb994d5c4f83e)
		__obf_12d456a52dd0fa73 = __obf_0dc283c37429c9a2.__obf_8e3ca5edc141df4a(__obf_274c443874c5ae86)
		__obf_47f1019286523e5e = __obf_12d456a52dd0fa73
		__obf_e613172f3fd23f33 = __obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6[__obf_0dc283c37429c9a2.__obf_27e0d051531f1b97[__obf_12d456a52dd0fa73]]
	)
	__obf_095ebebc00a3e9d4 = make([]M, 0, __obf_9610bb14d9eebcde)
	__obf_095ebebc00a3e9d4 = append(__obf_095ebebc00a3e9d4, __obf_e613172f3fd23f33)

	if len(__obf_095ebebc00a3e9d4) == __obf_9610bb14d9eebcde {
		return __obf_095ebebc00a3e9d4, nil
	}

	for __obf_12d456a52dd0fa73 = __obf_47f1019286523e5e + 1; __obf_12d456a52dd0fa73 != __obf_47f1019286523e5e; __obf_12d456a52dd0fa73++ {
		if __obf_12d456a52dd0fa73 >= len(__obf_0dc283c37429c9a2.__obf_27e0d051531f1b97) {
			__obf_12d456a52dd0fa73 = 0
		}
		__obf_e613172f3fd23f33 = __obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6[__obf_0dc283c37429c9a2.__obf_27e0d051531f1b97[__obf_12d456a52dd0fa73]]
		if !__obf_90cd2405eacba76d(__obf_095ebebc00a3e9d4, __obf_e613172f3fd23f33) {
			__obf_095ebebc00a3e9d4 = append(__obf_095ebebc00a3e9d4, __obf_e613172f3fd23f33)
		}
		if len(__obf_095ebebc00a3e9d4) == __obf_9610bb14d9eebcde {
			break
		}
	}

	return __obf_095ebebc00a3e9d4, nil
}

func (__obf_0dc283c37429c9a2 *Consistent[M]) __obf_6fe9ea8987a7e652(__obf_274c443874c5ae86 string) uint32 {
	if __obf_0dc283c37429c9a2.UseFnv {
		return __obf_0dc283c37429c9a2.__obf_2d39a95f5b829713(__obf_274c443874c5ae86)
	}
	return __obf_0dc283c37429c9a2.__obf_d97873f09e3206d5(__obf_274c443874c5ae86)
}

func (__obf_0dc283c37429c9a2 *Consistent[M]) __obf_d97873f09e3206d5(__obf_274c443874c5ae86 string) uint32 {
	if len(__obf_274c443874c5ae86) < 64 {
		var __obf_adfc779ee349fcc6 [64]byte
		copy(__obf_adfc779ee349fcc6[:], __obf_274c443874c5ae86)
		return crc32.ChecksumIEEE(__obf_adfc779ee349fcc6[:len(__obf_274c443874c5ae86)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_274c443874c5ae86))
}

func (__obf_0dc283c37429c9a2 *Consistent[M]) __obf_2d39a95f5b829713(__obf_274c443874c5ae86 string) uint32 {
	__obf_1017822b9d440de0 := fnv.New32a()
	__obf_1017822b9d440de0.
		Write([]byte(__obf_274c443874c5ae86))
	return __obf_1017822b9d440de0.Sum32()
}

func (__obf_0dc283c37429c9a2 *Consistent[M]) __obf_baaa01c5a4bbbc64() {
	__obf_20e36a127df3a3b8 := __obf_0dc283c37429c9a2.
		//reallocate if we're holding on to too much (1/4th)
		__obf_27e0d051531f1b97[:0]

	if cap(__obf_0dc283c37429c9a2.__obf_27e0d051531f1b97)/(__obf_0dc283c37429c9a2.NumberOfReplicas*4) > len(__obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6) {
		__obf_20e36a127df3a3b8 = nil
	}
	for __obf_068640d7e6c7dd1c := range __obf_0dc283c37429c9a2.__obf_d3b7e2fbbbb195b6 {
		__obf_20e36a127df3a3b8 = append(__obf_20e36a127df3a3b8, __obf_068640d7e6c7dd1c)
	}
	sort.Sort(__obf_20e36a127df3a3b8)
	__obf_0dc283c37429c9a2.__obf_27e0d051531f1b97 = __obf_20e36a127df3a3b8
}

func __obf_90cd2405eacba76d[M Member](__obf_5dff040934d98fec []M, __obf_edca33fae637d907 M) bool {
	for _, __obf_842e0740cba494d9 := range __obf_5dff040934d98fec {
		if __obf_842e0740cba494d9.String() == __obf_edca33fae637d907.String() {
			return true
		}
	}
	return false
}
