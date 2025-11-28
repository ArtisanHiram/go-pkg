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
package __obf_0301f59fb993362c // import "stathat.com/c/consistent"

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

type __obf_ea7d854be32dd080 []uint32

// Len returns the length of the uints array.
func (__obf_879ace077dd4e54b __obf_ea7d854be32dd080) Len() int { return len(__obf_879ace077dd4e54b) }

// Less returns true if element i is less than element j.
func (__obf_879ace077dd4e54b __obf_ea7d854be32dd080) Less(__obf_ad18e1a934138e00, __obf_158f24ef7a477bf6 int) bool {
	return __obf_879ace077dd4e54b[__obf_ad18e1a934138e00] < __obf_879ace077dd4e54b[__obf_158f24ef7a477bf6]
}

// Swap exchanges elements i and j.
func (__obf_879ace077dd4e54b __obf_ea7d854be32dd080) Swap(__obf_ad18e1a934138e00, __obf_158f24ef7a477bf6 int) {
	__obf_879ace077dd4e54b[__obf_ad18e1a934138e00], __obf_879ace077dd4e54b[__obf_158f24ef7a477bf6] = __obf_879ace077dd4e54b[__obf_158f24ef7a477bf6], __obf_879ace077dd4e54b[__obf_ad18e1a934138e00]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_f676dba18a27126b map[uint32]M
	__obf_f61740a2045b4df0 map[string]bool
	__obf_09f053a2f4d8c850 __obf_ea7d854be32dd080
	NumberOfReplicas       int
	__obf_d8ee2469570179b3 int64
	__obf_66728e2dd958dbeb [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_45f3e74a1d2883d1 := new(Consistent[M])
	__obf_45f3e74a1d2883d1.NumberOfReplicas = 20
	__obf_45f3e74a1d2883d1.__obf_f676dba18a27126b = make(map[uint32]M)
	__obf_45f3e74a1d2883d1.__obf_f61740a2045b4df0 = make(map[string]bool)
	return __obf_45f3e74a1d2883d1
}

// eltKey generates a string key for an element with an index.
func (__obf_45f3e74a1d2883d1 *Consistent[M]) __obf_c11f08586873b18e(__obf_3e361bdabbb7eb25 string, __obf_9451c6e3ec8eab54 int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_9451c6e3ec8eab54) + __obf_3e361bdabbb7eb25
}

// Add inserts a string element in the consistent hash.
func (__obf_45f3e74a1d2883d1 *Consistent[M]) Add(__obf_3e361bdabbb7eb25 M) {
	__obf_45f3e74a1d2883d1.Lock()
	defer __obf_45f3e74a1d2883d1.Unlock()
	__obf_45f3e74a1d2883d1.__obf_1533697d3b6d0abb(__obf_3e361bdabbb7eb25)
}

// need c.Lock() before calling
func (__obf_45f3e74a1d2883d1 *Consistent[M]) __obf_1533697d3b6d0abb(__obf_3e361bdabbb7eb25 M) {
	for __obf_ad18e1a934138e00 := 0; __obf_ad18e1a934138e00 < __obf_45f3e74a1d2883d1.NumberOfReplicas; __obf_ad18e1a934138e00++ {
		__obf_45f3e74a1d2883d1.__obf_f676dba18a27126b[__obf_45f3e74a1d2883d1.__obf_cdd4cca3812913ef(__obf_45f3e74a1d2883d1.__obf_c11f08586873b18e(__obf_3e361bdabbb7eb25.String(), __obf_ad18e1a934138e00))] = __obf_3e361bdabbb7eb25
	}
	__obf_45f3e74a1d2883d1.__obf_f61740a2045b4df0[__obf_3e361bdabbb7eb25.String()] = true
	__obf_45f3e74a1d2883d1.__obf_6c550ed99091584a()
	__obf_45f3e74a1d2883d1.__obf_d8ee2469570179b3++
}

// Remove removes an element from the hash.
func (__obf_45f3e74a1d2883d1 *Consistent[M]) Remove(__obf_3e361bdabbb7eb25 M) {
	__obf_45f3e74a1d2883d1.Lock()
	defer __obf_45f3e74a1d2883d1.Unlock()
	__obf_45f3e74a1d2883d1.__obf_e6c376bc5d8d1fea(__obf_3e361bdabbb7eb25.String())
}

// need c.Lock() before calling
func (__obf_45f3e74a1d2883d1 *Consistent[M]) __obf_e6c376bc5d8d1fea(__obf_3e361bdabbb7eb25 string) {
	for __obf_ad18e1a934138e00 := 0; __obf_ad18e1a934138e00 < __obf_45f3e74a1d2883d1.NumberOfReplicas; __obf_ad18e1a934138e00++ {
		delete(__obf_45f3e74a1d2883d1.__obf_f676dba18a27126b, __obf_45f3e74a1d2883d1.__obf_cdd4cca3812913ef(__obf_45f3e74a1d2883d1.__obf_c11f08586873b18e(__obf_3e361bdabbb7eb25, __obf_ad18e1a934138e00)))
	}
	delete(__obf_45f3e74a1d2883d1.__obf_f61740a2045b4df0, __obf_3e361bdabbb7eb25)
	__obf_45f3e74a1d2883d1.__obf_6c550ed99091584a()
	__obf_45f3e74a1d2883d1.__obf_d8ee2469570179b3--
}

// Set sets all the elements in the hash.  If there are existing elements not
// present in elts, they will be removed.
func (__obf_45f3e74a1d2883d1 *Consistent[M]) Set(__obf_781e19fecc6baf94 []M) {
	__obf_45f3e74a1d2883d1.Lock()
	defer __obf_45f3e74a1d2883d1.Unlock()
	for __obf_68b8821c15b38b11 := range __obf_45f3e74a1d2883d1.__obf_f61740a2045b4df0 {
		__obf_d894b115c623f2f7 := false
		for _, __obf_a38164fb2f2e9688 := range __obf_781e19fecc6baf94 {
			if __obf_68b8821c15b38b11 == __obf_a38164fb2f2e9688.String() {
				__obf_d894b115c623f2f7 = true
				break
			}
		}
		if !__obf_d894b115c623f2f7 {
			__obf_45f3e74a1d2883d1.__obf_e6c376bc5d8d1fea(__obf_68b8821c15b38b11)
		}
	}
	for _, __obf_a38164fb2f2e9688 := range __obf_781e19fecc6baf94 {
		_, __obf_61c8a08170380307 := __obf_45f3e74a1d2883d1.__obf_f61740a2045b4df0[__obf_a38164fb2f2e9688.String()]
		if __obf_61c8a08170380307 {
			continue
		}
		__obf_45f3e74a1d2883d1.__obf_1533697d3b6d0abb(__obf_a38164fb2f2e9688)
	}
}

func (__obf_45f3e74a1d2883d1 *Consistent[M]) Members() []string {
	__obf_45f3e74a1d2883d1.RLock()
	defer __obf_45f3e74a1d2883d1.RUnlock()
	var __obf_2bcdbc79239bddad []string
	for __obf_68b8821c15b38b11 := range __obf_45f3e74a1d2883d1.__obf_f61740a2045b4df0 {
		__obf_2bcdbc79239bddad = append(__obf_2bcdbc79239bddad, __obf_68b8821c15b38b11)
	}
	return __obf_2bcdbc79239bddad
}

// Get returns an element close to where name hashes to in the circle.
func (__obf_45f3e74a1d2883d1 *Consistent[M]) Get(__obf_b437c8c482d48be5 string) (__obf_7b6425723213d03c M, __obf_e3cd0f441610b7cc error) {
	__obf_45f3e74a1d2883d1.RLock()
	defer __obf_45f3e74a1d2883d1.RUnlock()
	if len(__obf_45f3e74a1d2883d1.__obf_f676dba18a27126b) == 0 {
		__obf_e3cd0f441610b7cc = ErrEmptyCircle
		return
	}
	__obf_c60784b512777679 := __obf_45f3e74a1d2883d1.__obf_cdd4cca3812913ef(__obf_b437c8c482d48be5)
	__obf_ad18e1a934138e00 := __obf_45f3e74a1d2883d1.__obf_d1ce1f26a8b0175a(__obf_c60784b512777679)
	__obf_7b6425723213d03c = __obf_45f3e74a1d2883d1.__obf_f676dba18a27126b[__obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850[__obf_ad18e1a934138e00]]
	return
}

func (__obf_45f3e74a1d2883d1 *Consistent[M]) __obf_d1ce1f26a8b0175a(__obf_c60784b512777679 uint32) (__obf_ad18e1a934138e00 int) {
	__obf_5d4f2a00e0e1b89a := func(__obf_879ace077dd4e54b int) bool {
		return __obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850[__obf_879ace077dd4e54b] > __obf_c60784b512777679
	}
	__obf_ad18e1a934138e00 = sort.Search(len(__obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850), __obf_5d4f2a00e0e1b89a)
	if __obf_ad18e1a934138e00 >= len(__obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850) {
		__obf_ad18e1a934138e00 = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_45f3e74a1d2883d1 *Consistent[M]) GetTwo(__obf_b437c8c482d48be5 string) (__obf_3a88a37147a7932d M, __obf_fae98bd6447babcb M, __obf_e3cd0f441610b7cc error) {
	__obf_45f3e74a1d2883d1.RLock()
	defer __obf_45f3e74a1d2883d1.RUnlock()
	if len(__obf_45f3e74a1d2883d1.__obf_f676dba18a27126b) == 0 {
		__obf_e3cd0f441610b7cc = ErrEmptyCircle
		return
	}
	__obf_c60784b512777679 := __obf_45f3e74a1d2883d1.__obf_cdd4cca3812913ef(__obf_b437c8c482d48be5)
	__obf_ad18e1a934138e00 := __obf_45f3e74a1d2883d1.__obf_d1ce1f26a8b0175a(__obf_c60784b512777679)
	__obf_3a88a37147a7932d = __obf_45f3e74a1d2883d1.__obf_f676dba18a27126b[__obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850[__obf_ad18e1a934138e00]]

	if __obf_45f3e74a1d2883d1.__obf_d8ee2469570179b3 == 1 {
		return
	}

	__obf_b6d8a02ce2b1f839 := __obf_ad18e1a934138e00
	for __obf_ad18e1a934138e00 = __obf_b6d8a02ce2b1f839 + 1; __obf_ad18e1a934138e00 != __obf_b6d8a02ce2b1f839; __obf_ad18e1a934138e00++ {
		if __obf_ad18e1a934138e00 >= len(__obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850) {
			__obf_ad18e1a934138e00 = 0
		}
		__obf_fae98bd6447babcb = __obf_45f3e74a1d2883d1.__obf_f676dba18a27126b[__obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850[__obf_ad18e1a934138e00]]
		if __obf_fae98bd6447babcb.String() != __obf_3a88a37147a7932d.String() {
			break
		}
	}
	return __obf_3a88a37147a7932d, __obf_fae98bd6447babcb, nil
}

// GetN returns the N closest distinct elements to the name input in the circle.
func (__obf_45f3e74a1d2883d1 *Consistent[M]) GetN(__obf_b437c8c482d48be5 string, __obf_fe29c92e32f3b303 int) (__obf_7b6425723213d03c []M, __obf_e3cd0f441610b7cc error) {
	__obf_45f3e74a1d2883d1.RLock()
	defer __obf_45f3e74a1d2883d1.RUnlock()

	if len(__obf_45f3e74a1d2883d1.__obf_f676dba18a27126b) == 0 {
		__obf_e3cd0f441610b7cc = ErrEmptyCircle
		return
	}

	if __obf_45f3e74a1d2883d1.__obf_d8ee2469570179b3 < int64(__obf_fe29c92e32f3b303) {
		__obf_fe29c92e32f3b303 = int(__obf_45f3e74a1d2883d1.__obf_d8ee2469570179b3)
	}

	var (
		__obf_c60784b512777679 = __obf_45f3e74a1d2883d1.__obf_cdd4cca3812913ef(__obf_b437c8c482d48be5)
		__obf_ad18e1a934138e00 = __obf_45f3e74a1d2883d1.__obf_d1ce1f26a8b0175a(__obf_c60784b512777679)
		__obf_b6d8a02ce2b1f839 = __obf_ad18e1a934138e00
		__obf_1151e60760466a83 = __obf_45f3e74a1d2883d1.__obf_f676dba18a27126b[__obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850[__obf_ad18e1a934138e00]]
	)
	__obf_7b6425723213d03c = make([]M, 0, __obf_fe29c92e32f3b303)
	__obf_7b6425723213d03c = append(__obf_7b6425723213d03c, __obf_1151e60760466a83)

	if len(__obf_7b6425723213d03c) == __obf_fe29c92e32f3b303 {
		return __obf_7b6425723213d03c, nil
	}

	for __obf_ad18e1a934138e00 = __obf_b6d8a02ce2b1f839 + 1; __obf_ad18e1a934138e00 != __obf_b6d8a02ce2b1f839; __obf_ad18e1a934138e00++ {
		if __obf_ad18e1a934138e00 >= len(__obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850) {
			__obf_ad18e1a934138e00 = 0
		}
		__obf_1151e60760466a83 = __obf_45f3e74a1d2883d1.__obf_f676dba18a27126b[__obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850[__obf_ad18e1a934138e00]]
		if !__obf_a85ceb86d583435b(__obf_7b6425723213d03c, __obf_1151e60760466a83) {
			__obf_7b6425723213d03c = append(__obf_7b6425723213d03c, __obf_1151e60760466a83)
		}
		if len(__obf_7b6425723213d03c) == __obf_fe29c92e32f3b303 {
			break
		}
	}

	return __obf_7b6425723213d03c, nil
}

func (__obf_45f3e74a1d2883d1 *Consistent[M]) __obf_cdd4cca3812913ef(__obf_c60784b512777679 string) uint32 {
	if __obf_45f3e74a1d2883d1.UseFnv {
		return __obf_45f3e74a1d2883d1.__obf_9db1052ccb1861a5(__obf_c60784b512777679)
	}
	return __obf_45f3e74a1d2883d1.__obf_042491e190f1981e(__obf_c60784b512777679)
}

func (__obf_45f3e74a1d2883d1 *Consistent[M]) __obf_042491e190f1981e(__obf_c60784b512777679 string) uint32 {
	if len(__obf_c60784b512777679) < 64 {
		var __obf_66728e2dd958dbeb [64]byte
		copy(__obf_66728e2dd958dbeb[:], __obf_c60784b512777679)
		return crc32.ChecksumIEEE(__obf_66728e2dd958dbeb[:len(__obf_c60784b512777679)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_c60784b512777679))
}

func (__obf_45f3e74a1d2883d1 *Consistent[M]) __obf_9db1052ccb1861a5(__obf_c60784b512777679 string) uint32 {
	__obf_eeb66db60ecb4d3c := fnv.New32a()
	__obf_eeb66db60ecb4d3c.Write([]byte(__obf_c60784b512777679))
	return __obf_eeb66db60ecb4d3c.Sum32()
}

func (__obf_45f3e74a1d2883d1 *Consistent[M]) __obf_6c550ed99091584a() {
	__obf_a8f3ace2a1f8b68e := __obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850[:0]
	//reallocate if we're holding on to too much (1/4th)
	if cap(__obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850)/(__obf_45f3e74a1d2883d1.NumberOfReplicas*4) > len(__obf_45f3e74a1d2883d1.__obf_f676dba18a27126b) {
		__obf_a8f3ace2a1f8b68e = nil
	}
	for __obf_68b8821c15b38b11 := range __obf_45f3e74a1d2883d1.__obf_f676dba18a27126b {
		__obf_a8f3ace2a1f8b68e = append(__obf_a8f3ace2a1f8b68e, __obf_68b8821c15b38b11)
	}
	sort.Sort(__obf_a8f3ace2a1f8b68e)
	__obf_45f3e74a1d2883d1.__obf_09f053a2f4d8c850 = __obf_a8f3ace2a1f8b68e
}

func __obf_a85ceb86d583435b[M Member](__obf_a826451fdc738304 []M, __obf_a8a7598880f4a3c4 M) bool {
	for _, __obf_2bcdbc79239bddad := range __obf_a826451fdc738304 {
		if __obf_2bcdbc79239bddad.String() == __obf_a8a7598880f4a3c4.String() {
			return true
		}
	}
	return false
}
