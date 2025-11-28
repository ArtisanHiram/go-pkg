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
package __obf_849f8354176204bc // import "stathat.com/c/consistent"

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

type __obf_b36abd4368a43704 []uint32

// Len returns the length of the uints array.
func (__obf_6c1b84d1658f9956 __obf_b36abd4368a43704) Len() int { return len(__obf_6c1b84d1658f9956) }

// Less returns true if element i is less than element j.
func (__obf_6c1b84d1658f9956 __obf_b36abd4368a43704) Less(__obf_746419697917435d, __obf_d0e1b90125d7d48a int) bool {
	return __obf_6c1b84d1658f9956[__obf_746419697917435d] < __obf_6c1b84d1658f9956[__obf_d0e1b90125d7d48a]
}

// Swap exchanges elements i and j.
func (__obf_6c1b84d1658f9956 __obf_b36abd4368a43704) Swap(__obf_746419697917435d, __obf_d0e1b90125d7d48a int) {
	__obf_6c1b84d1658f9956[__obf_746419697917435d], __obf_6c1b84d1658f9956[__obf_d0e1b90125d7d48a] = __obf_6c1b84d1658f9956[__obf_d0e1b90125d7d48a], __obf_6c1b84d1658f9956[__obf_746419697917435d]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_1da2c0edc4733d31 map[uint32]M
	__obf_75a48131fe662d8c map[string]bool
	__obf_e2b02186e95df7eb __obf_b36abd4368a43704
	NumberOfReplicas       int
	__obf_b60c2e053d30256b int64
	__obf_5199462a95084636 [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_cd8444e0360a1344 := new(Consistent[M])
	__obf_cd8444e0360a1344.NumberOfReplicas = 20
	__obf_cd8444e0360a1344.__obf_1da2c0edc4733d31 = make(map[uint32]M)
	__obf_cd8444e0360a1344.__obf_75a48131fe662d8c = make(map[string]bool)
	return __obf_cd8444e0360a1344
}

// eltKey generates a string key for an element with an index.
func (__obf_cd8444e0360a1344 *Consistent[M]) __obf_bdbc535b5267d147(__obf_e58d1b2ff6ed1ba1 string, __obf_4746ce89668b9e58 int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_4746ce89668b9e58) + __obf_e58d1b2ff6ed1ba1
}

// Add inserts a string element in the consistent hash.
func (__obf_cd8444e0360a1344 *Consistent[M]) Add(__obf_e58d1b2ff6ed1ba1 M) {
	__obf_cd8444e0360a1344.Lock()
	defer __obf_cd8444e0360a1344.Unlock()
	__obf_cd8444e0360a1344.__obf_b096b975c3b23625(__obf_e58d1b2ff6ed1ba1)
}

// need c.Lock() before calling
func (__obf_cd8444e0360a1344 *Consistent[M]) __obf_b096b975c3b23625(__obf_e58d1b2ff6ed1ba1 M) {
	for __obf_746419697917435d := 0; __obf_746419697917435d < __obf_cd8444e0360a1344.NumberOfReplicas; __obf_746419697917435d++ {
		__obf_cd8444e0360a1344.__obf_1da2c0edc4733d31[__obf_cd8444e0360a1344.__obf_803a53778fcdb987(__obf_cd8444e0360a1344.__obf_bdbc535b5267d147(__obf_e58d1b2ff6ed1ba1.String(), __obf_746419697917435d))] = __obf_e58d1b2ff6ed1ba1
	}
	__obf_cd8444e0360a1344.__obf_75a48131fe662d8c[__obf_e58d1b2ff6ed1ba1.String()] = true
	__obf_cd8444e0360a1344.__obf_2e6fbf9f9d8032df()
	__obf_cd8444e0360a1344.__obf_b60c2e053d30256b++
}

// Remove removes an element from the hash.
func (__obf_cd8444e0360a1344 *Consistent[M]) Remove(__obf_e58d1b2ff6ed1ba1 M) {
	__obf_cd8444e0360a1344.Lock()
	defer __obf_cd8444e0360a1344.Unlock()
	__obf_cd8444e0360a1344.__obf_d7635c56f4765ba6(__obf_e58d1b2ff6ed1ba1.String())
}

// need c.Lock() before calling
func (__obf_cd8444e0360a1344 *Consistent[M]) __obf_d7635c56f4765ba6(__obf_e58d1b2ff6ed1ba1 string) {
	for __obf_746419697917435d := 0; __obf_746419697917435d < __obf_cd8444e0360a1344.NumberOfReplicas; __obf_746419697917435d++ {
		delete(__obf_cd8444e0360a1344.__obf_1da2c0edc4733d31, __obf_cd8444e0360a1344.__obf_803a53778fcdb987(__obf_cd8444e0360a1344.__obf_bdbc535b5267d147(__obf_e58d1b2ff6ed1ba1, __obf_746419697917435d)))
	}
	delete(__obf_cd8444e0360a1344.__obf_75a48131fe662d8c, __obf_e58d1b2ff6ed1ba1)
	__obf_cd8444e0360a1344.__obf_2e6fbf9f9d8032df()
	__obf_cd8444e0360a1344.__obf_b60c2e053d30256b--
}

// Set sets all the elements in the hash.  If there are existing elements not
// present in elts, they will be removed.
func (__obf_cd8444e0360a1344 *Consistent[M]) Set(__obf_03a464f124693d4e []M) {
	__obf_cd8444e0360a1344.Lock()
	defer __obf_cd8444e0360a1344.Unlock()
	for __obf_bc470d1d2d751fec := range __obf_cd8444e0360a1344.__obf_75a48131fe662d8c {
		__obf_da7485193b0e34ac := false
		for _, __obf_dbce546d3476033d := range __obf_03a464f124693d4e {
			if __obf_bc470d1d2d751fec == __obf_dbce546d3476033d.String() {
				__obf_da7485193b0e34ac = true
				break
			}
		}
		if !__obf_da7485193b0e34ac {
			__obf_cd8444e0360a1344.__obf_d7635c56f4765ba6(__obf_bc470d1d2d751fec)
		}
	}
	for _, __obf_dbce546d3476033d := range __obf_03a464f124693d4e {
		_, __obf_e0a71e103f3085fa := __obf_cd8444e0360a1344.__obf_75a48131fe662d8c[__obf_dbce546d3476033d.String()]
		if __obf_e0a71e103f3085fa {
			continue
		}
		__obf_cd8444e0360a1344.__obf_b096b975c3b23625(__obf_dbce546d3476033d)
	}
}

func (__obf_cd8444e0360a1344 *Consistent[M]) Members() []string {
	__obf_cd8444e0360a1344.RLock()
	defer __obf_cd8444e0360a1344.RUnlock()
	var __obf_ba34ef5ec45a58ed []string
	for __obf_bc470d1d2d751fec := range __obf_cd8444e0360a1344.__obf_75a48131fe662d8c {
		__obf_ba34ef5ec45a58ed = append(__obf_ba34ef5ec45a58ed, __obf_bc470d1d2d751fec)
	}
	return __obf_ba34ef5ec45a58ed
}

// Get returns an element close to where name hashes to in the circle.
func (__obf_cd8444e0360a1344 *Consistent[M]) Get(__obf_fc30a0ffde27b11a string) (__obf_99920f89cccf7aac M, __obf_7378bafad3e22943 error) {
	__obf_cd8444e0360a1344.RLock()
	defer __obf_cd8444e0360a1344.RUnlock()
	if len(__obf_cd8444e0360a1344.__obf_1da2c0edc4733d31) == 0 {
		__obf_7378bafad3e22943 = ErrEmptyCircle
		return
	}
	__obf_cf3642e971002f87 := __obf_cd8444e0360a1344.__obf_803a53778fcdb987(__obf_fc30a0ffde27b11a)
	__obf_746419697917435d := __obf_cd8444e0360a1344.__obf_03483bc73a66245a(__obf_cf3642e971002f87)
	__obf_99920f89cccf7aac = __obf_cd8444e0360a1344.__obf_1da2c0edc4733d31[__obf_cd8444e0360a1344.__obf_e2b02186e95df7eb[__obf_746419697917435d]]
	return
}

func (__obf_cd8444e0360a1344 *Consistent[M]) __obf_03483bc73a66245a(__obf_cf3642e971002f87 uint32) (__obf_746419697917435d int) {
	__obf_e17611e690ba715c := func(__obf_6c1b84d1658f9956 int) bool {
		return __obf_cd8444e0360a1344.__obf_e2b02186e95df7eb[__obf_6c1b84d1658f9956] > __obf_cf3642e971002f87
	}
	__obf_746419697917435d = sort.Search(len(__obf_cd8444e0360a1344.__obf_e2b02186e95df7eb), __obf_e17611e690ba715c)
	if __obf_746419697917435d >= len(__obf_cd8444e0360a1344.__obf_e2b02186e95df7eb) {
		__obf_746419697917435d = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_cd8444e0360a1344 *Consistent[M]) GetTwo(__obf_fc30a0ffde27b11a string) (__obf_ceb650b4b6602a88 M, __obf_fd3ff99aa948e491 M, __obf_7378bafad3e22943 error) {
	__obf_cd8444e0360a1344.RLock()
	defer __obf_cd8444e0360a1344.RUnlock()
	if len(__obf_cd8444e0360a1344.__obf_1da2c0edc4733d31) == 0 {
		__obf_7378bafad3e22943 = ErrEmptyCircle
		return
	}
	__obf_cf3642e971002f87 := __obf_cd8444e0360a1344.__obf_803a53778fcdb987(__obf_fc30a0ffde27b11a)
	__obf_746419697917435d := __obf_cd8444e0360a1344.__obf_03483bc73a66245a(__obf_cf3642e971002f87)
	__obf_ceb650b4b6602a88 = __obf_cd8444e0360a1344.__obf_1da2c0edc4733d31[__obf_cd8444e0360a1344.__obf_e2b02186e95df7eb[__obf_746419697917435d]]

	if __obf_cd8444e0360a1344.__obf_b60c2e053d30256b == 1 {
		return
	}

	__obf_a15ecf3e28411da7 := __obf_746419697917435d
	for __obf_746419697917435d = __obf_a15ecf3e28411da7 + 1; __obf_746419697917435d != __obf_a15ecf3e28411da7; __obf_746419697917435d++ {
		if __obf_746419697917435d >= len(__obf_cd8444e0360a1344.__obf_e2b02186e95df7eb) {
			__obf_746419697917435d = 0
		}
		__obf_fd3ff99aa948e491 = __obf_cd8444e0360a1344.__obf_1da2c0edc4733d31[__obf_cd8444e0360a1344.__obf_e2b02186e95df7eb[__obf_746419697917435d]]
		if __obf_fd3ff99aa948e491.String() != __obf_ceb650b4b6602a88.String() {
			break
		}
	}
	return __obf_ceb650b4b6602a88, __obf_fd3ff99aa948e491, nil
}

// GetN returns the N closest distinct elements to the name input in the circle.
func (__obf_cd8444e0360a1344 *Consistent[M]) GetN(__obf_fc30a0ffde27b11a string, __obf_2be6b486cd304c24 int) (__obf_99920f89cccf7aac []M, __obf_7378bafad3e22943 error) {
	__obf_cd8444e0360a1344.RLock()
	defer __obf_cd8444e0360a1344.RUnlock()

	if len(__obf_cd8444e0360a1344.__obf_1da2c0edc4733d31) == 0 {
		__obf_7378bafad3e22943 = ErrEmptyCircle
		return
	}

	if __obf_cd8444e0360a1344.__obf_b60c2e053d30256b < int64(__obf_2be6b486cd304c24) {
		__obf_2be6b486cd304c24 = int(__obf_cd8444e0360a1344.__obf_b60c2e053d30256b)
	}

	var (
		__obf_cf3642e971002f87 = __obf_cd8444e0360a1344.__obf_803a53778fcdb987(__obf_fc30a0ffde27b11a)
		__obf_746419697917435d = __obf_cd8444e0360a1344.__obf_03483bc73a66245a(__obf_cf3642e971002f87)
		__obf_a15ecf3e28411da7 = __obf_746419697917435d
		__obf_2af4f9aa4a6dfc38 = __obf_cd8444e0360a1344.__obf_1da2c0edc4733d31[__obf_cd8444e0360a1344.__obf_e2b02186e95df7eb[__obf_746419697917435d]]
	)
	__obf_99920f89cccf7aac = make([]M, 0, __obf_2be6b486cd304c24)
	__obf_99920f89cccf7aac = append(__obf_99920f89cccf7aac, __obf_2af4f9aa4a6dfc38)

	if len(__obf_99920f89cccf7aac) == __obf_2be6b486cd304c24 {
		return __obf_99920f89cccf7aac, nil
	}

	for __obf_746419697917435d = __obf_a15ecf3e28411da7 + 1; __obf_746419697917435d != __obf_a15ecf3e28411da7; __obf_746419697917435d++ {
		if __obf_746419697917435d >= len(__obf_cd8444e0360a1344.__obf_e2b02186e95df7eb) {
			__obf_746419697917435d = 0
		}
		__obf_2af4f9aa4a6dfc38 = __obf_cd8444e0360a1344.__obf_1da2c0edc4733d31[__obf_cd8444e0360a1344.__obf_e2b02186e95df7eb[__obf_746419697917435d]]
		if !__obf_1418ff6558e8d181(__obf_99920f89cccf7aac, __obf_2af4f9aa4a6dfc38) {
			__obf_99920f89cccf7aac = append(__obf_99920f89cccf7aac, __obf_2af4f9aa4a6dfc38)
		}
		if len(__obf_99920f89cccf7aac) == __obf_2be6b486cd304c24 {
			break
		}
	}

	return __obf_99920f89cccf7aac, nil
}

func (__obf_cd8444e0360a1344 *Consistent[M]) __obf_803a53778fcdb987(__obf_cf3642e971002f87 string) uint32 {
	if __obf_cd8444e0360a1344.UseFnv {
		return __obf_cd8444e0360a1344.__obf_06d7614c181dc8b1(__obf_cf3642e971002f87)
	}
	return __obf_cd8444e0360a1344.__obf_14200fe05c4cda20(__obf_cf3642e971002f87)
}

func (__obf_cd8444e0360a1344 *Consistent[M]) __obf_14200fe05c4cda20(__obf_cf3642e971002f87 string) uint32 {
	if len(__obf_cf3642e971002f87) < 64 {
		var __obf_5199462a95084636 [64]byte
		copy(__obf_5199462a95084636[:], __obf_cf3642e971002f87)
		return crc32.ChecksumIEEE(__obf_5199462a95084636[:len(__obf_cf3642e971002f87)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_cf3642e971002f87))
}

func (__obf_cd8444e0360a1344 *Consistent[M]) __obf_06d7614c181dc8b1(__obf_cf3642e971002f87 string) uint32 {
	__obf_d9f245a1c87558c5 := fnv.New32a()
	__obf_d9f245a1c87558c5.Write([]byte(__obf_cf3642e971002f87))
	return __obf_d9f245a1c87558c5.Sum32()
}

func (__obf_cd8444e0360a1344 *Consistent[M]) __obf_2e6fbf9f9d8032df() {
	__obf_18f6058abd23c009 := __obf_cd8444e0360a1344.__obf_e2b02186e95df7eb[:0]
	//reallocate if we're holding on to too much (1/4th)
	if cap(__obf_cd8444e0360a1344.__obf_e2b02186e95df7eb)/(__obf_cd8444e0360a1344.NumberOfReplicas*4) > len(__obf_cd8444e0360a1344.__obf_1da2c0edc4733d31) {
		__obf_18f6058abd23c009 = nil
	}
	for __obf_bc470d1d2d751fec := range __obf_cd8444e0360a1344.__obf_1da2c0edc4733d31 {
		__obf_18f6058abd23c009 = append(__obf_18f6058abd23c009, __obf_bc470d1d2d751fec)
	}
	sort.Sort(__obf_18f6058abd23c009)
	__obf_cd8444e0360a1344.__obf_e2b02186e95df7eb = __obf_18f6058abd23c009
}

func __obf_1418ff6558e8d181[M Member](__obf_23122e564ceea8a2 []M, __obf_e719fb6c88c2330c M) bool {
	for _, __obf_ba34ef5ec45a58ed := range __obf_23122e564ceea8a2 {
		if __obf_ba34ef5ec45a58ed.String() == __obf_e719fb6c88c2330c.String() {
			return true
		}
	}
	return false
}
