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
package __obf_8f3c509a07205b31 // import "stathat.com/c/consistent"

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

type __obf_c6bee850db210df2 []uint32

// Len returns the length of the uints array.
func (__obf_0f902eeba149524b __obf_c6bee850db210df2) Len() int { return len(__obf_0f902eeba149524b) }

// Less returns true if element i is less than element j.
func (__obf_0f902eeba149524b __obf_c6bee850db210df2) Less(__obf_ef8bfb29ca7bb0e7, __obf_8023d95e9e2b70cb int) bool {
	return __obf_0f902eeba149524b[__obf_ef8bfb29ca7bb0e7] < __obf_0f902eeba149524b[__obf_8023d95e9e2b70cb]
}

// Swap exchanges elements i and j.
func (__obf_0f902eeba149524b __obf_c6bee850db210df2) Swap(__obf_ef8bfb29ca7bb0e7, __obf_8023d95e9e2b70cb int) {
	__obf_0f902eeba149524b[__obf_ef8bfb29ca7bb0e7], __obf_0f902eeba149524b[__obf_8023d95e9e2b70cb] = __obf_0f902eeba149524b[__obf_8023d95e9e2b70cb], __obf_0f902eeba149524b[__obf_ef8bfb29ca7bb0e7]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_0749833d0072201f map[uint32]M
	__obf_d77b34c229e62938 map[string]bool
	__obf_dcd802f772ef1e85 __obf_c6bee850db210df2
	NumberOfReplicas       int
	__obf_e6a16545ab1bb25e int64
	__obf_93c8d5a5be91c57f [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_96e8d883e607c580 := new(Consistent[M])
	__obf_96e8d883e607c580.NumberOfReplicas = 20
	__obf_96e8d883e607c580.__obf_0749833d0072201f = make(map[uint32]M)
	__obf_96e8d883e607c580.__obf_d77b34c229e62938 = make(map[string]bool)
	return __obf_96e8d883e607c580
}

// eltKey generates a string key for an element with an index.
func (__obf_96e8d883e607c580 *Consistent[M]) __obf_dfd75ab4b22c01a5(__obf_a285e9ff0d8f1e0e string, __obf_d12821b8cd0ac7e2 int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_d12821b8cd0ac7e2) + __obf_a285e9ff0d8f1e0e
}

// Add inserts a string element in the consistent hash.
func (__obf_96e8d883e607c580 *Consistent[M]) Add(__obf_a285e9ff0d8f1e0e M) {
	__obf_96e8d883e607c580.Lock()
	defer __obf_96e8d883e607c580.Unlock()
	__obf_96e8d883e607c580.__obf_9f2b40cf84f5673a(__obf_a285e9ff0d8f1e0e)
}

// need c.Lock() before calling
func (__obf_96e8d883e607c580 *Consistent[M]) __obf_9f2b40cf84f5673a(__obf_a285e9ff0d8f1e0e M) {
	for __obf_ef8bfb29ca7bb0e7 := 0; __obf_ef8bfb29ca7bb0e7 < __obf_96e8d883e607c580.NumberOfReplicas; __obf_ef8bfb29ca7bb0e7++ {
		__obf_96e8d883e607c580.__obf_0749833d0072201f[__obf_96e8d883e607c580.__obf_56ea79bb300d2d00(__obf_96e8d883e607c580.__obf_dfd75ab4b22c01a5(__obf_a285e9ff0d8f1e0e.String(), __obf_ef8bfb29ca7bb0e7))] = __obf_a285e9ff0d8f1e0e
	}
	__obf_96e8d883e607c580.__obf_d77b34c229e62938[__obf_a285e9ff0d8f1e0e.String()] = true
	__obf_96e8d883e607c580.__obf_33374dcacc20fcd6()
	__obf_96e8d883e607c580.__obf_e6a16545ab1bb25e++
}

// Remove removes an element from the hash.
func (__obf_96e8d883e607c580 *Consistent[M]) Remove(__obf_a285e9ff0d8f1e0e M) {
	__obf_96e8d883e607c580.Lock()
	defer __obf_96e8d883e607c580.Unlock()
	__obf_96e8d883e607c580.__obf_5e220b0f9a913164(__obf_a285e9ff0d8f1e0e.String())
}

// need c.Lock() before calling
func (__obf_96e8d883e607c580 *Consistent[M]) __obf_5e220b0f9a913164(__obf_a285e9ff0d8f1e0e string) {
	for __obf_ef8bfb29ca7bb0e7 := 0; __obf_ef8bfb29ca7bb0e7 < __obf_96e8d883e607c580.NumberOfReplicas; __obf_ef8bfb29ca7bb0e7++ {
		delete(__obf_96e8d883e607c580.__obf_0749833d0072201f, __obf_96e8d883e607c580.__obf_56ea79bb300d2d00(__obf_96e8d883e607c580.__obf_dfd75ab4b22c01a5(__obf_a285e9ff0d8f1e0e, __obf_ef8bfb29ca7bb0e7)))
	}
	delete(__obf_96e8d883e607c580.__obf_d77b34c229e62938, __obf_a285e9ff0d8f1e0e)
	__obf_96e8d883e607c580.__obf_33374dcacc20fcd6()
	__obf_96e8d883e607c580.__obf_e6a16545ab1bb25e--
}

// Set sets all the elements in the hash.  If there are existing elements not
// present in elts, they will be removed.
func (__obf_96e8d883e607c580 *Consistent[M]) Set(__obf_4b5241775a3e8e91 []M) {
	__obf_96e8d883e607c580.Lock()
	defer __obf_96e8d883e607c580.Unlock()
	for __obf_1e3322cd8cce88d9 := range __obf_96e8d883e607c580.__obf_d77b34c229e62938 {
		__obf_68a7277be5c0a20c := false
		for _, __obf_fd725b9b9ace9362 := range __obf_4b5241775a3e8e91 {
			if __obf_1e3322cd8cce88d9 == __obf_fd725b9b9ace9362.String() {
				__obf_68a7277be5c0a20c = true
				break
			}
		}
		if !__obf_68a7277be5c0a20c {
			__obf_96e8d883e607c580.__obf_5e220b0f9a913164(__obf_1e3322cd8cce88d9)
		}
	}
	for _, __obf_fd725b9b9ace9362 := range __obf_4b5241775a3e8e91 {
		_, __obf_85ac2f5253937ed3 := __obf_96e8d883e607c580.__obf_d77b34c229e62938[__obf_fd725b9b9ace9362.String()]
		if __obf_85ac2f5253937ed3 {
			continue
		}
		__obf_96e8d883e607c580.__obf_9f2b40cf84f5673a(__obf_fd725b9b9ace9362)
	}
}

func (__obf_96e8d883e607c580 *Consistent[M]) Members() []string {
	__obf_96e8d883e607c580.RLock()
	defer __obf_96e8d883e607c580.RUnlock()
	var __obf_5151484d7675b9c5 []string
	for __obf_1e3322cd8cce88d9 := range __obf_96e8d883e607c580.__obf_d77b34c229e62938 {
		__obf_5151484d7675b9c5 = append(__obf_5151484d7675b9c5, __obf_1e3322cd8cce88d9)
	}
	return __obf_5151484d7675b9c5
}

// Get returns an element close to where name hashes to in the circle.
func (__obf_96e8d883e607c580 *Consistent[M]) Get(__obf_e53789d526f8eefa string) (__obf_f775428bffb8e930 M, __obf_4515281320f6bc24 error) {
	__obf_96e8d883e607c580.RLock()
	defer __obf_96e8d883e607c580.RUnlock()
	if len(__obf_96e8d883e607c580.__obf_0749833d0072201f) == 0 {
		__obf_4515281320f6bc24 = ErrEmptyCircle
		return
	}
	__obf_8007fcbcf485f7d5 := __obf_96e8d883e607c580.__obf_56ea79bb300d2d00(__obf_e53789d526f8eefa)
	__obf_ef8bfb29ca7bb0e7 := __obf_96e8d883e607c580.__obf_987e975ab8f9fa62(__obf_8007fcbcf485f7d5)
	__obf_f775428bffb8e930 = __obf_96e8d883e607c580.__obf_0749833d0072201f[__obf_96e8d883e607c580.__obf_dcd802f772ef1e85[__obf_ef8bfb29ca7bb0e7]]
	return
}

func (__obf_96e8d883e607c580 *Consistent[M]) __obf_987e975ab8f9fa62(__obf_8007fcbcf485f7d5 uint32) (__obf_ef8bfb29ca7bb0e7 int) {
	__obf_03bae068df7551dd := func(__obf_0f902eeba149524b int) bool {
		return __obf_96e8d883e607c580.__obf_dcd802f772ef1e85[__obf_0f902eeba149524b] > __obf_8007fcbcf485f7d5
	}
	__obf_ef8bfb29ca7bb0e7 = sort.Search(len(__obf_96e8d883e607c580.__obf_dcd802f772ef1e85), __obf_03bae068df7551dd)
	if __obf_ef8bfb29ca7bb0e7 >= len(__obf_96e8d883e607c580.__obf_dcd802f772ef1e85) {
		__obf_ef8bfb29ca7bb0e7 = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_96e8d883e607c580 *Consistent[M]) GetTwo(__obf_e53789d526f8eefa string) (__obf_e4978263edc2faa6 M, __obf_4f5d0acaef39cf69 M, __obf_4515281320f6bc24 error) {
	__obf_96e8d883e607c580.RLock()
	defer __obf_96e8d883e607c580.RUnlock()
	if len(__obf_96e8d883e607c580.__obf_0749833d0072201f) == 0 {
		__obf_4515281320f6bc24 = ErrEmptyCircle
		return
	}
	__obf_8007fcbcf485f7d5 := __obf_96e8d883e607c580.__obf_56ea79bb300d2d00(__obf_e53789d526f8eefa)
	__obf_ef8bfb29ca7bb0e7 := __obf_96e8d883e607c580.__obf_987e975ab8f9fa62(__obf_8007fcbcf485f7d5)
	__obf_e4978263edc2faa6 = __obf_96e8d883e607c580.__obf_0749833d0072201f[__obf_96e8d883e607c580.__obf_dcd802f772ef1e85[__obf_ef8bfb29ca7bb0e7]]

	if __obf_96e8d883e607c580.__obf_e6a16545ab1bb25e == 1 {
		return
	}

	__obf_449382ef4cde0d91 := __obf_ef8bfb29ca7bb0e7
	for __obf_ef8bfb29ca7bb0e7 = __obf_449382ef4cde0d91 + 1; __obf_ef8bfb29ca7bb0e7 != __obf_449382ef4cde0d91; __obf_ef8bfb29ca7bb0e7++ {
		if __obf_ef8bfb29ca7bb0e7 >= len(__obf_96e8d883e607c580.__obf_dcd802f772ef1e85) {
			__obf_ef8bfb29ca7bb0e7 = 0
		}
		__obf_4f5d0acaef39cf69 = __obf_96e8d883e607c580.__obf_0749833d0072201f[__obf_96e8d883e607c580.__obf_dcd802f772ef1e85[__obf_ef8bfb29ca7bb0e7]]
		if __obf_4f5d0acaef39cf69.String() != __obf_e4978263edc2faa6.String() {
			break
		}
	}
	return __obf_e4978263edc2faa6, __obf_4f5d0acaef39cf69, nil
}

// GetN returns the N closest distinct elements to the name input in the circle.
func (__obf_96e8d883e607c580 *Consistent[M]) GetN(__obf_e53789d526f8eefa string, __obf_d8ce5564b2d36bf2 int) (__obf_f775428bffb8e930 []M, __obf_4515281320f6bc24 error) {
	__obf_96e8d883e607c580.RLock()
	defer __obf_96e8d883e607c580.RUnlock()

	if len(__obf_96e8d883e607c580.__obf_0749833d0072201f) == 0 {
		__obf_4515281320f6bc24 = ErrEmptyCircle
		return
	}

	if __obf_96e8d883e607c580.__obf_e6a16545ab1bb25e < int64(__obf_d8ce5564b2d36bf2) {
		__obf_d8ce5564b2d36bf2 = int(__obf_96e8d883e607c580.__obf_e6a16545ab1bb25e)
	}

	var (
		__obf_8007fcbcf485f7d5 = __obf_96e8d883e607c580.__obf_56ea79bb300d2d00(__obf_e53789d526f8eefa)
		__obf_ef8bfb29ca7bb0e7 = __obf_96e8d883e607c580.__obf_987e975ab8f9fa62(__obf_8007fcbcf485f7d5)
		__obf_449382ef4cde0d91 = __obf_ef8bfb29ca7bb0e7
		__obf_02b15fec94d63c4b = __obf_96e8d883e607c580.__obf_0749833d0072201f[__obf_96e8d883e607c580.__obf_dcd802f772ef1e85[__obf_ef8bfb29ca7bb0e7]]
	)
	__obf_f775428bffb8e930 = make([]M, 0, __obf_d8ce5564b2d36bf2)
	__obf_f775428bffb8e930 = append(__obf_f775428bffb8e930, __obf_02b15fec94d63c4b)

	if len(__obf_f775428bffb8e930) == __obf_d8ce5564b2d36bf2 {
		return __obf_f775428bffb8e930, nil
	}

	for __obf_ef8bfb29ca7bb0e7 = __obf_449382ef4cde0d91 + 1; __obf_ef8bfb29ca7bb0e7 != __obf_449382ef4cde0d91; __obf_ef8bfb29ca7bb0e7++ {
		if __obf_ef8bfb29ca7bb0e7 >= len(__obf_96e8d883e607c580.__obf_dcd802f772ef1e85) {
			__obf_ef8bfb29ca7bb0e7 = 0
		}
		__obf_02b15fec94d63c4b = __obf_96e8d883e607c580.__obf_0749833d0072201f[__obf_96e8d883e607c580.__obf_dcd802f772ef1e85[__obf_ef8bfb29ca7bb0e7]]
		if !__obf_9d3ab9eee5e3ee01(__obf_f775428bffb8e930, __obf_02b15fec94d63c4b) {
			__obf_f775428bffb8e930 = append(__obf_f775428bffb8e930, __obf_02b15fec94d63c4b)
		}
		if len(__obf_f775428bffb8e930) == __obf_d8ce5564b2d36bf2 {
			break
		}
	}

	return __obf_f775428bffb8e930, nil
}

func (__obf_96e8d883e607c580 *Consistent[M]) __obf_56ea79bb300d2d00(__obf_8007fcbcf485f7d5 string) uint32 {
	if __obf_96e8d883e607c580.UseFnv {
		return __obf_96e8d883e607c580.__obf_ec712e08e222ec7e(__obf_8007fcbcf485f7d5)
	}
	return __obf_96e8d883e607c580.__obf_8090921b8d3635dc(__obf_8007fcbcf485f7d5)
}

func (__obf_96e8d883e607c580 *Consistent[M]) __obf_8090921b8d3635dc(__obf_8007fcbcf485f7d5 string) uint32 {
	if len(__obf_8007fcbcf485f7d5) < 64 {
		var __obf_93c8d5a5be91c57f [64]byte
		copy(__obf_93c8d5a5be91c57f[:], __obf_8007fcbcf485f7d5)
		return crc32.ChecksumIEEE(__obf_93c8d5a5be91c57f[:len(__obf_8007fcbcf485f7d5)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_8007fcbcf485f7d5))
}

func (__obf_96e8d883e607c580 *Consistent[M]) __obf_ec712e08e222ec7e(__obf_8007fcbcf485f7d5 string) uint32 {
	__obf_7daf57395622cd78 := fnv.New32a()
	__obf_7daf57395622cd78.Write([]byte(__obf_8007fcbcf485f7d5))
	return __obf_7daf57395622cd78.Sum32()
}

func (__obf_96e8d883e607c580 *Consistent[M]) __obf_33374dcacc20fcd6() {
	__obf_0789923633e562ac := __obf_96e8d883e607c580.__obf_dcd802f772ef1e85[:0]
	//reallocate if we're holding on to too much (1/4th)
	if cap(__obf_96e8d883e607c580.__obf_dcd802f772ef1e85)/(__obf_96e8d883e607c580.NumberOfReplicas*4) > len(__obf_96e8d883e607c580.__obf_0749833d0072201f) {
		__obf_0789923633e562ac = nil
	}
	for __obf_1e3322cd8cce88d9 := range __obf_96e8d883e607c580.__obf_0749833d0072201f {
		__obf_0789923633e562ac = append(__obf_0789923633e562ac, __obf_1e3322cd8cce88d9)
	}
	sort.Sort(__obf_0789923633e562ac)
	__obf_96e8d883e607c580.__obf_dcd802f772ef1e85 = __obf_0789923633e562ac
}

func __obf_9d3ab9eee5e3ee01[M Member](__obf_545831dcd906cc9a []M, __obf_e1a94a4626212de8 M) bool {
	for _, __obf_5151484d7675b9c5 := range __obf_545831dcd906cc9a {
		if __obf_5151484d7675b9c5.String() == __obf_e1a94a4626212de8.String() {
			return true
		}
	}
	return false
}
