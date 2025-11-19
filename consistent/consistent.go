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
package __obf_13e392706e0afef4 // import "stathat.com/c/consistent"

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

type __obf_386734d66d8f1224 []uint32

// Len returns the length of the uints array.
func (__obf_a16d53b4d84e2d86 __obf_386734d66d8f1224) Len() int { return len(__obf_a16d53b4d84e2d86) }

// Less returns true if element i is less than element j.
func (__obf_a16d53b4d84e2d86 __obf_386734d66d8f1224) Less(__obf_752c63ce61fb53df, __obf_6d867bb9f7f30ab3 int) bool {
	return __obf_a16d53b4d84e2d86[__obf_752c63ce61fb53df] < __obf_a16d53b4d84e2d86[__obf_6d867bb9f7f30ab3]
}

// Swap exchanges elements i and j.
func (__obf_a16d53b4d84e2d86 __obf_386734d66d8f1224) Swap(__obf_752c63ce61fb53df, __obf_6d867bb9f7f30ab3 int) {
	__obf_a16d53b4d84e2d86[__obf_752c63ce61fb53df], __obf_a16d53b4d84e2d86[__obf_6d867bb9f7f30ab3] = __obf_a16d53b4d84e2d86[__obf_6d867bb9f7f30ab3], __obf_a16d53b4d84e2d86[__obf_752c63ce61fb53df]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_0dd44ad07a2e74ae map[uint32]M
	__obf_71822fdcbf4577a0 map[string]bool
	__obf_2098dfe2f1462732 __obf_386734d66d8f1224
	NumberOfReplicas       int
	__obf_0a369e0fe346acf9 int64
	__obf_8753cad51edc4c0f [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_0dc399645195e6c8 := new(Consistent[M])
	__obf_0dc399645195e6c8.NumberOfReplicas = 20
	__obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae = make(map[uint32]M)
	__obf_0dc399645195e6c8.__obf_71822fdcbf4577a0 = make(map[string]bool)
	return __obf_0dc399645195e6c8
}

// eltKey generates a string key for an element with an index.
func (__obf_0dc399645195e6c8 *Consistent[M]) __obf_c9d61d49fa22adf8(__obf_86fd24133742eaf7 string, __obf_f5a982aa1f663948 int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_f5a982aa1f663948) + __obf_86fd24133742eaf7
}

// Add inserts a string element in the consistent hash.
func (__obf_0dc399645195e6c8 *Consistent[M]) Add(__obf_86fd24133742eaf7 M) {
	__obf_0dc399645195e6c8.Lock()
	defer __obf_0dc399645195e6c8.Unlock()
	__obf_0dc399645195e6c8.__obf_95ecc864cdc3d61e(__obf_86fd24133742eaf7)
}

// need c.Lock() before calling
func (__obf_0dc399645195e6c8 *Consistent[M]) __obf_95ecc864cdc3d61e(__obf_86fd24133742eaf7 M) {
	for __obf_752c63ce61fb53df := 0; __obf_752c63ce61fb53df < __obf_0dc399645195e6c8.NumberOfReplicas; __obf_752c63ce61fb53df++ {
		__obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae[__obf_0dc399645195e6c8.__obf_35ebcd48d872850a(__obf_0dc399645195e6c8.__obf_c9d61d49fa22adf8(__obf_86fd24133742eaf7.String(), __obf_752c63ce61fb53df))] = __obf_86fd24133742eaf7
	}
	__obf_0dc399645195e6c8.__obf_71822fdcbf4577a0[__obf_86fd24133742eaf7.String()] = true
	__obf_0dc399645195e6c8.__obf_a72ccf45466c63bf()
	__obf_0dc399645195e6c8.__obf_0a369e0fe346acf9++
}

// Remove removes an element from the hash.
func (__obf_0dc399645195e6c8 *Consistent[M]) Remove(__obf_86fd24133742eaf7 M) {
	__obf_0dc399645195e6c8.Lock()
	defer __obf_0dc399645195e6c8.Unlock()
	__obf_0dc399645195e6c8.__obf_e6baebd8b63e5af6(__obf_86fd24133742eaf7.String())
}

// need c.Lock() before calling
func (__obf_0dc399645195e6c8 *Consistent[M]) __obf_e6baebd8b63e5af6(__obf_86fd24133742eaf7 string) {
	for __obf_752c63ce61fb53df := 0; __obf_752c63ce61fb53df < __obf_0dc399645195e6c8.NumberOfReplicas; __obf_752c63ce61fb53df++ {
		delete(__obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae, __obf_0dc399645195e6c8.__obf_35ebcd48d872850a(__obf_0dc399645195e6c8.__obf_c9d61d49fa22adf8(__obf_86fd24133742eaf7, __obf_752c63ce61fb53df)))
	}
	delete(__obf_0dc399645195e6c8.__obf_71822fdcbf4577a0, __obf_86fd24133742eaf7)
	__obf_0dc399645195e6c8.__obf_a72ccf45466c63bf()
	__obf_0dc399645195e6c8.__obf_0a369e0fe346acf9--
}

// Set sets all the elements in the hash.  If there are existing elements not
// present in elts, they will be removed.
func (__obf_0dc399645195e6c8 *Consistent[M]) Set(__obf_0a23168f509068da []M) {
	__obf_0dc399645195e6c8.Lock()
	defer __obf_0dc399645195e6c8.Unlock()
	for __obf_a309b1293482f43a := range __obf_0dc399645195e6c8.__obf_71822fdcbf4577a0 {
		__obf_7cec256377a1b05f := false
		for _, __obf_522b152a21ef3f9b := range __obf_0a23168f509068da {
			if __obf_a309b1293482f43a == __obf_522b152a21ef3f9b.String() {
				__obf_7cec256377a1b05f = true
				break
			}
		}
		if !__obf_7cec256377a1b05f {
			__obf_0dc399645195e6c8.__obf_e6baebd8b63e5af6(__obf_a309b1293482f43a)
		}
	}
	for _, __obf_522b152a21ef3f9b := range __obf_0a23168f509068da {
		_, __obf_3dcc651784a57de7 := __obf_0dc399645195e6c8.__obf_71822fdcbf4577a0[__obf_522b152a21ef3f9b.String()]
		if __obf_3dcc651784a57de7 {
			continue
		}
		__obf_0dc399645195e6c8.__obf_95ecc864cdc3d61e(__obf_522b152a21ef3f9b)
	}
}

func (__obf_0dc399645195e6c8 *Consistent[M]) Members() []string {
	__obf_0dc399645195e6c8.RLock()
	defer __obf_0dc399645195e6c8.RUnlock()
	var __obf_2c51d8f25d13347b []string
	for __obf_a309b1293482f43a := range __obf_0dc399645195e6c8.__obf_71822fdcbf4577a0 {
		__obf_2c51d8f25d13347b = append(__obf_2c51d8f25d13347b, __obf_a309b1293482f43a)
	}
	return __obf_2c51d8f25d13347b
}

// Get returns an element close to where name hashes to in the circle.
func (__obf_0dc399645195e6c8 *Consistent[M]) Get(__obf_dffe96d76c2d47c2 string) (__obf_b566dd9d15432f80 M, __obf_2de13886252b7cc1 error) {
	__obf_0dc399645195e6c8.RLock()
	defer __obf_0dc399645195e6c8.RUnlock()
	if len(__obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae) == 0 {
		__obf_2de13886252b7cc1 = ErrEmptyCircle
		return
	}
	__obf_aff96e1c015cdc58 := __obf_0dc399645195e6c8.__obf_35ebcd48d872850a(__obf_dffe96d76c2d47c2)
	__obf_752c63ce61fb53df := __obf_0dc399645195e6c8.__obf_67c83c627e28556c(__obf_aff96e1c015cdc58)
	__obf_b566dd9d15432f80 = __obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae[__obf_0dc399645195e6c8.__obf_2098dfe2f1462732[__obf_752c63ce61fb53df]]
	return
}

func (__obf_0dc399645195e6c8 *Consistent[M]) __obf_67c83c627e28556c(__obf_aff96e1c015cdc58 uint32) (__obf_752c63ce61fb53df int) {
	__obf_fc783cd906417da1 := func(__obf_a16d53b4d84e2d86 int) bool {
		return __obf_0dc399645195e6c8.__obf_2098dfe2f1462732[__obf_a16d53b4d84e2d86] > __obf_aff96e1c015cdc58
	}
	__obf_752c63ce61fb53df = sort.Search(len(__obf_0dc399645195e6c8.__obf_2098dfe2f1462732), __obf_fc783cd906417da1)
	if __obf_752c63ce61fb53df >= len(__obf_0dc399645195e6c8.__obf_2098dfe2f1462732) {
		__obf_752c63ce61fb53df = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_0dc399645195e6c8 *Consistent[M]) GetTwo(__obf_dffe96d76c2d47c2 string) (__obf_0217e324f491d16b M, __obf_8636920731655a20 M, __obf_2de13886252b7cc1 error) {
	__obf_0dc399645195e6c8.RLock()
	defer __obf_0dc399645195e6c8.RUnlock()
	if len(__obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae) == 0 {
		__obf_2de13886252b7cc1 = ErrEmptyCircle
		return
	}
	__obf_aff96e1c015cdc58 := __obf_0dc399645195e6c8.__obf_35ebcd48d872850a(__obf_dffe96d76c2d47c2)
	__obf_752c63ce61fb53df := __obf_0dc399645195e6c8.__obf_67c83c627e28556c(__obf_aff96e1c015cdc58)
	__obf_0217e324f491d16b = __obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae[__obf_0dc399645195e6c8.__obf_2098dfe2f1462732[__obf_752c63ce61fb53df]]

	if __obf_0dc399645195e6c8.__obf_0a369e0fe346acf9 == 1 {
		return
	}

	__obf_688af3952a19cbc1 := __obf_752c63ce61fb53df
	for __obf_752c63ce61fb53df = __obf_688af3952a19cbc1 + 1; __obf_752c63ce61fb53df != __obf_688af3952a19cbc1; __obf_752c63ce61fb53df++ {
		if __obf_752c63ce61fb53df >= len(__obf_0dc399645195e6c8.__obf_2098dfe2f1462732) {
			__obf_752c63ce61fb53df = 0
		}
		__obf_8636920731655a20 = __obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae[__obf_0dc399645195e6c8.__obf_2098dfe2f1462732[__obf_752c63ce61fb53df]]
		if __obf_8636920731655a20.String() != __obf_0217e324f491d16b.String() {
			break
		}
	}
	return __obf_0217e324f491d16b, __obf_8636920731655a20, nil
}

// GetN returns the N closest distinct elements to the name input in the circle.
func (__obf_0dc399645195e6c8 *Consistent[M]) GetN(__obf_dffe96d76c2d47c2 string, __obf_5aa6e958b1876e77 int) (__obf_b566dd9d15432f80 []M, __obf_2de13886252b7cc1 error) {
	__obf_0dc399645195e6c8.RLock()
	defer __obf_0dc399645195e6c8.RUnlock()

	if len(__obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae) == 0 {
		__obf_2de13886252b7cc1 = ErrEmptyCircle
		return
	}

	if __obf_0dc399645195e6c8.__obf_0a369e0fe346acf9 < int64(__obf_5aa6e958b1876e77) {
		__obf_5aa6e958b1876e77 = int(__obf_0dc399645195e6c8.__obf_0a369e0fe346acf9)
	}

	var (
		__obf_aff96e1c015cdc58 = __obf_0dc399645195e6c8.__obf_35ebcd48d872850a(__obf_dffe96d76c2d47c2)
		__obf_752c63ce61fb53df = __obf_0dc399645195e6c8.__obf_67c83c627e28556c(__obf_aff96e1c015cdc58)
		__obf_688af3952a19cbc1 = __obf_752c63ce61fb53df
		__obf_b4959245983c0234 = __obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae[__obf_0dc399645195e6c8.__obf_2098dfe2f1462732[__obf_752c63ce61fb53df]]
	)
	__obf_b566dd9d15432f80 = make([]M, 0, __obf_5aa6e958b1876e77)
	__obf_b566dd9d15432f80 = append(__obf_b566dd9d15432f80, __obf_b4959245983c0234)

	if len(__obf_b566dd9d15432f80) == __obf_5aa6e958b1876e77 {
		return __obf_b566dd9d15432f80, nil
	}

	for __obf_752c63ce61fb53df = __obf_688af3952a19cbc1 + 1; __obf_752c63ce61fb53df != __obf_688af3952a19cbc1; __obf_752c63ce61fb53df++ {
		if __obf_752c63ce61fb53df >= len(__obf_0dc399645195e6c8.__obf_2098dfe2f1462732) {
			__obf_752c63ce61fb53df = 0
		}
		__obf_b4959245983c0234 = __obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae[__obf_0dc399645195e6c8.__obf_2098dfe2f1462732[__obf_752c63ce61fb53df]]
		if !__obf_a695557906aae573(__obf_b566dd9d15432f80, __obf_b4959245983c0234) {
			__obf_b566dd9d15432f80 = append(__obf_b566dd9d15432f80, __obf_b4959245983c0234)
		}
		if len(__obf_b566dd9d15432f80) == __obf_5aa6e958b1876e77 {
			break
		}
	}

	return __obf_b566dd9d15432f80, nil
}

func (__obf_0dc399645195e6c8 *Consistent[M]) __obf_35ebcd48d872850a(__obf_aff96e1c015cdc58 string) uint32 {
	if __obf_0dc399645195e6c8.UseFnv {
		return __obf_0dc399645195e6c8.__obf_46be3eece4f1b783(__obf_aff96e1c015cdc58)
	}
	return __obf_0dc399645195e6c8.__obf_24e10131bf88436c(__obf_aff96e1c015cdc58)
}

func (__obf_0dc399645195e6c8 *Consistent[M]) __obf_24e10131bf88436c(__obf_aff96e1c015cdc58 string) uint32 {
	if len(__obf_aff96e1c015cdc58) < 64 {
		var __obf_8753cad51edc4c0f [64]byte
		copy(__obf_8753cad51edc4c0f[:], __obf_aff96e1c015cdc58)
		return crc32.ChecksumIEEE(__obf_8753cad51edc4c0f[:len(__obf_aff96e1c015cdc58)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_aff96e1c015cdc58))
}

func (__obf_0dc399645195e6c8 *Consistent[M]) __obf_46be3eece4f1b783(__obf_aff96e1c015cdc58 string) uint32 {
	__obf_3253bac028fd816d := fnv.New32a()
	__obf_3253bac028fd816d.Write([]byte(__obf_aff96e1c015cdc58))
	return __obf_3253bac028fd816d.Sum32()
}

func (__obf_0dc399645195e6c8 *Consistent[M]) __obf_a72ccf45466c63bf() {
	__obf_93e966fd14b50ee0 := __obf_0dc399645195e6c8.__obf_2098dfe2f1462732[:0]
	//reallocate if we're holding on to too much (1/4th)
	if cap(__obf_0dc399645195e6c8.__obf_2098dfe2f1462732)/(__obf_0dc399645195e6c8.NumberOfReplicas*4) > len(__obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae) {
		__obf_93e966fd14b50ee0 = nil
	}
	for __obf_a309b1293482f43a := range __obf_0dc399645195e6c8.__obf_0dd44ad07a2e74ae {
		__obf_93e966fd14b50ee0 = append(__obf_93e966fd14b50ee0, __obf_a309b1293482f43a)
	}
	sort.Sort(__obf_93e966fd14b50ee0)
	__obf_0dc399645195e6c8.__obf_2098dfe2f1462732 = __obf_93e966fd14b50ee0
}

func __obf_a695557906aae573[M Member](__obf_a62d77c68de9fdd8 []M, __obf_1e235d63d343716b M) bool {
	for _, __obf_2c51d8f25d13347b := range __obf_a62d77c68de9fdd8 {
		if __obf_2c51d8f25d13347b.String() == __obf_1e235d63d343716b.String() {
			return true
		}
	}
	return false
}
