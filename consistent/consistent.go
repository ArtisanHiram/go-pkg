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
package __obf_6d605e4b82079612 // import "stathat.com/c/consistent"

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

type __obf_848abdb64e9e1f6a []uint32

// Len returns the length of the uints array.
func (__obf_f18336805ea9bb25 __obf_848abdb64e9e1f6a) Len() int { return len(__obf_f18336805ea9bb25) }

// Less returns true if element i is less than element j.
func (__obf_f18336805ea9bb25 __obf_848abdb64e9e1f6a) Less(__obf_c53644879e44d663, __obf_7466ff8079ff3aa1 int) bool {
	return __obf_f18336805ea9bb25[__obf_c53644879e44d663] < __obf_f18336805ea9bb25[__obf_7466ff8079ff3aa1]
}

// Swap exchanges elements i and j.
func (__obf_f18336805ea9bb25 __obf_848abdb64e9e1f6a) Swap(__obf_c53644879e44d663, __obf_7466ff8079ff3aa1 int) {
	__obf_f18336805ea9bb25[__obf_c53644879e44d663], __obf_f18336805ea9bb25[__obf_7466ff8079ff3aa1] = __obf_f18336805ea9bb25[__obf_7466ff8079ff3aa1], __obf_f18336805ea9bb25[__obf_c53644879e44d663]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_7d563a82a34c175d map[uint32]M
	__obf_94696ab114168f20 map[string]bool
	__obf_b75bdb0b287cd566 __obf_848abdb64e9e1f6a

	NumberOfReplicas       int
	__obf_271c3fd5bf629f29 int64
	__obf_eca85a9fd0767cf3 [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_aa40f13747012419 := new(Consistent[M])
	__obf_aa40f13747012419.
		NumberOfReplicas = 20
	__obf_aa40f13747012419.__obf_7d563a82a34c175d = make(map[uint32]M)
	__obf_aa40f13747012419.__obf_94696ab114168f20 = make(map[string]bool)
	return __obf_aa40f13747012419
}

// eltKey generates a string key for an element with an index.
func (__obf_aa40f13747012419 *Consistent[M]) __obf_ffd902bcb9fbc7b6(__obf_a409850a633d35b5 string, __obf_401a3bc8ec966689 int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_401a3bc8ec966689) + __obf_a409850a633d35b5
}

// Add inserts a string element in the consistent hash.
func (__obf_aa40f13747012419 *Consistent[M]) Add(__obf_a409850a633d35b5 M) {
	__obf_aa40f13747012419.
		Lock()
	defer __obf_aa40f13747012419.Unlock()
	__obf_aa40f13747012419.

		// need c.Lock() before calling
		__obf_68494662dc9180d8(__obf_a409850a633d35b5)
}

func (__obf_aa40f13747012419 *Consistent[M]) __obf_68494662dc9180d8(__obf_a409850a633d35b5 M) {
	for __obf_c53644879e44d663 := 0; __obf_c53644879e44d663 < __obf_aa40f13747012419.NumberOfReplicas; __obf_c53644879e44d663++ {
		__obf_aa40f13747012419.__obf_7d563a82a34c175d[__obf_aa40f13747012419.__obf_0f254dfcf68c19b1(__obf_aa40f13747012419.__obf_ffd902bcb9fbc7b6(__obf_a409850a633d35b5.String(), __obf_c53644879e44d663))] = __obf_a409850a633d35b5
	}
	__obf_aa40f13747012419.__obf_94696ab114168f20[__obf_a409850a633d35b5.String()] = true
	__obf_aa40f13747012419.__obf_e85602afe7615e81()
	__obf_aa40f13747012419.

		// Remove removes an element from the hash.
		__obf_271c3fd5bf629f29++
}

func (__obf_aa40f13747012419 *Consistent[M]) Remove(__obf_a409850a633d35b5 M) {
	__obf_aa40f13747012419.
		Lock()
	defer __obf_aa40f13747012419.Unlock()
	__obf_aa40f13747012419.__obf_df4e3c8988fad7e7(__obf_a409850a633d35b5.String())
}

// need c.Lock() before calling
func (__obf_aa40f13747012419 *Consistent[M]) __obf_df4e3c8988fad7e7(__obf_a409850a633d35b5 string) {
	for __obf_c53644879e44d663 := 0; __obf_c53644879e44d663 < __obf_aa40f13747012419.NumberOfReplicas; __obf_c53644879e44d663++ {
		delete(__obf_aa40f13747012419.__obf_7d563a82a34c175d, __obf_aa40f13747012419.__obf_0f254dfcf68c19b1(__obf_aa40f13747012419.__obf_ffd902bcb9fbc7b6(__obf_a409850a633d35b5, __obf_c53644879e44d663)))
	}
	delete(__obf_aa40f13747012419.__obf_94696ab114168f20, __obf_a409850a633d35b5)
	__obf_aa40f13747012419.__obf_e85602afe7615e81()
	__obf_aa40f13747012419.

		// Set sets all the elements in the hash.  If there are existing elements not
		// present in elts, they will be removed.
		__obf_271c3fd5bf629f29--
}

func (__obf_aa40f13747012419 *Consistent[M]) Set(__obf_c035362e8f9d1f57 []M) {
	__obf_aa40f13747012419.
		Lock()
	defer __obf_aa40f13747012419.Unlock()
	for __obf_56227607a6f80707 := range __obf_aa40f13747012419.__obf_94696ab114168f20 {
		__obf_bcbaa076b4d561dd := false
		for _, __obf_855168a493756c4f := range __obf_c035362e8f9d1f57 {
			if __obf_56227607a6f80707 == __obf_855168a493756c4f.String() {
				__obf_bcbaa076b4d561dd = true
				break
			}
		}
		if !__obf_bcbaa076b4d561dd {
			__obf_aa40f13747012419.__obf_df4e3c8988fad7e7(__obf_56227607a6f80707)
		}
	}
	for _, __obf_855168a493756c4f := range __obf_c035362e8f9d1f57 {
		_, __obf_6a692b27dfdf9798 := __obf_aa40f13747012419.__obf_94696ab114168f20[__obf_855168a493756c4f.String()]
		if __obf_6a692b27dfdf9798 {
			continue
		}
		__obf_aa40f13747012419.__obf_68494662dc9180d8(__obf_855168a493756c4f)
	}
}

func (__obf_aa40f13747012419 *Consistent[M]) Members() []string {
	__obf_aa40f13747012419.
		RLock()
	defer __obf_aa40f13747012419.RUnlock()
	var __obf_910264ae455d2054 []string
	for __obf_56227607a6f80707 := range __obf_aa40f13747012419.__obf_94696ab114168f20 {
		__obf_910264ae455d2054 = append(__obf_910264ae455d2054,

			// Get returns an element close to where name hashes to in the circle.
			__obf_56227607a6f80707)
	}
	return __obf_910264ae455d2054
}

func (__obf_aa40f13747012419 *Consistent[M]) Get(__obf_6e1bbb7735f334b4 string) (__obf_7d0e09fed385b540 M, __obf_d8d2fc5ed7e54dd6 error) {
	__obf_aa40f13747012419.
		RLock()
	defer __obf_aa40f13747012419.RUnlock()
	if len(__obf_aa40f13747012419.__obf_7d563a82a34c175d) == 0 {
		__obf_d8d2fc5ed7e54dd6 = ErrEmptyCircle
		return
	}
	__obf_ed9fc90c1e7fbdbf := __obf_aa40f13747012419.__obf_0f254dfcf68c19b1(__obf_6e1bbb7735f334b4)
	__obf_c53644879e44d663 := __obf_aa40f13747012419.__obf_a832d33e912dd16e(__obf_ed9fc90c1e7fbdbf)
	__obf_7d0e09fed385b540 = __obf_aa40f13747012419.__obf_7d563a82a34c175d[__obf_aa40f13747012419.__obf_b75bdb0b287cd566[__obf_c53644879e44d663]]
	return
}

func (__obf_aa40f13747012419 *Consistent[M]) __obf_a832d33e912dd16e(__obf_ed9fc90c1e7fbdbf uint32) (__obf_c53644879e44d663 int) {
	__obf_b44ebfb6b67b7245 := func(__obf_f18336805ea9bb25 int) bool {
		return __obf_aa40f13747012419.__obf_b75bdb0b287cd566[__obf_f18336805ea9bb25] > __obf_ed9fc90c1e7fbdbf
	}
	__obf_c53644879e44d663 = sort.Search(len(__obf_aa40f13747012419.__obf_b75bdb0b287cd566), __obf_b44ebfb6b67b7245)
	if __obf_c53644879e44d663 >= len(__obf_aa40f13747012419.__obf_b75bdb0b287cd566) {
		__obf_c53644879e44d663 = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_aa40f13747012419 *Consistent[M]) GetTwo(__obf_6e1bbb7735f334b4 string) (__obf_ed87d1cd75a77ab4 M, __obf_b9d78ab8e5febd46 M, __obf_d8d2fc5ed7e54dd6 error) {
	__obf_aa40f13747012419.
		RLock()
	defer __obf_aa40f13747012419.RUnlock()
	if len(__obf_aa40f13747012419.__obf_7d563a82a34c175d) == 0 {
		__obf_d8d2fc5ed7e54dd6 = ErrEmptyCircle
		return
	}
	__obf_ed9fc90c1e7fbdbf := __obf_aa40f13747012419.__obf_0f254dfcf68c19b1(__obf_6e1bbb7735f334b4)
	__obf_c53644879e44d663 := __obf_aa40f13747012419.__obf_a832d33e912dd16e(__obf_ed9fc90c1e7fbdbf)
	__obf_ed87d1cd75a77ab4 = __obf_aa40f13747012419.__obf_7d563a82a34c175d[__obf_aa40f13747012419.__obf_b75bdb0b287cd566[__obf_c53644879e44d663]]

	if __obf_aa40f13747012419.__obf_271c3fd5bf629f29 == 1 {
		return
	}
	__obf_cbfb7d3e767c76bb := __obf_c53644879e44d663
	for __obf_c53644879e44d663 = __obf_cbfb7d3e767c76bb + 1; __obf_c53644879e44d663 != __obf_cbfb7d3e767c76bb; __obf_c53644879e44d663++ {
		if __obf_c53644879e44d663 >= len(__obf_aa40f13747012419.__obf_b75bdb0b287cd566) {
			__obf_c53644879e44d663 = 0
		}
		__obf_b9d78ab8e5febd46 = __obf_aa40f13747012419.__obf_7d563a82a34c175d[__obf_aa40f13747012419.__obf_b75bdb0b287cd566[__obf_c53644879e44d663]]
		if __obf_b9d78ab8e5febd46.String() != __obf_ed87d1cd75a77ab4.String() {
			break
		}
	}
	return __obf_ed87d1cd75a77ab4,

		// GetN returns the N closest distinct elements to the name input in the circle.
		__obf_b9d78ab8e5febd46, nil
}

func (__obf_aa40f13747012419 *Consistent[M]) GetN(__obf_6e1bbb7735f334b4 string, __obf_a2dd6b6ed2807978 int) (__obf_7d0e09fed385b540 []M, __obf_d8d2fc5ed7e54dd6 error) {
	__obf_aa40f13747012419.
		RLock()
	defer __obf_aa40f13747012419.RUnlock()

	if len(__obf_aa40f13747012419.__obf_7d563a82a34c175d) == 0 {
		__obf_d8d2fc5ed7e54dd6 = ErrEmptyCircle
		return
	}

	if __obf_aa40f13747012419.__obf_271c3fd5bf629f29 < int64(__obf_a2dd6b6ed2807978) {
		__obf_a2dd6b6ed2807978 = int(__obf_aa40f13747012419.__obf_271c3fd5bf629f29)
	}

	var (
		__obf_ed9fc90c1e7fbdbf = __obf_aa40f13747012419.__obf_0f254dfcf68c19b1(__obf_6e1bbb7735f334b4)
		__obf_c53644879e44d663 = __obf_aa40f13747012419.__obf_a832d33e912dd16e(__obf_ed9fc90c1e7fbdbf)
		__obf_cbfb7d3e767c76bb = __obf_c53644879e44d663
		__obf_6478ed988457b2e7 = __obf_aa40f13747012419.__obf_7d563a82a34c175d[__obf_aa40f13747012419.__obf_b75bdb0b287cd566[__obf_c53644879e44d663]]
	)
	__obf_7d0e09fed385b540 = make([]M, 0, __obf_a2dd6b6ed2807978)
	__obf_7d0e09fed385b540 = append(__obf_7d0e09fed385b540, __obf_6478ed988457b2e7)

	if len(__obf_7d0e09fed385b540) == __obf_a2dd6b6ed2807978 {
		return __obf_7d0e09fed385b540, nil
	}

	for __obf_c53644879e44d663 = __obf_cbfb7d3e767c76bb + 1; __obf_c53644879e44d663 != __obf_cbfb7d3e767c76bb; __obf_c53644879e44d663++ {
		if __obf_c53644879e44d663 >= len(__obf_aa40f13747012419.__obf_b75bdb0b287cd566) {
			__obf_c53644879e44d663 = 0
		}
		__obf_6478ed988457b2e7 = __obf_aa40f13747012419.__obf_7d563a82a34c175d[__obf_aa40f13747012419.__obf_b75bdb0b287cd566[__obf_c53644879e44d663]]
		if !__obf_5af7c0cbcdd4f639(__obf_7d0e09fed385b540, __obf_6478ed988457b2e7) {
			__obf_7d0e09fed385b540 = append(__obf_7d0e09fed385b540, __obf_6478ed988457b2e7)
		}
		if len(__obf_7d0e09fed385b540) == __obf_a2dd6b6ed2807978 {
			break
		}
	}

	return __obf_7d0e09fed385b540, nil
}

func (__obf_aa40f13747012419 *Consistent[M]) __obf_0f254dfcf68c19b1(__obf_ed9fc90c1e7fbdbf string) uint32 {
	if __obf_aa40f13747012419.UseFnv {
		return __obf_aa40f13747012419.__obf_f784cdc682713e78(__obf_ed9fc90c1e7fbdbf)
	}
	return __obf_aa40f13747012419.__obf_1087f66f96ce2f73(__obf_ed9fc90c1e7fbdbf)
}

func (__obf_aa40f13747012419 *Consistent[M]) __obf_1087f66f96ce2f73(__obf_ed9fc90c1e7fbdbf string) uint32 {
	if len(__obf_ed9fc90c1e7fbdbf) < 64 {
		var __obf_eca85a9fd0767cf3 [64]byte
		copy(__obf_eca85a9fd0767cf3[:], __obf_ed9fc90c1e7fbdbf)
		return crc32.ChecksumIEEE(__obf_eca85a9fd0767cf3[:len(__obf_ed9fc90c1e7fbdbf)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_ed9fc90c1e7fbdbf))
}

func (__obf_aa40f13747012419 *Consistent[M]) __obf_f784cdc682713e78(__obf_ed9fc90c1e7fbdbf string) uint32 {
	__obf_30915fb6c7afe620 := fnv.New32a()
	__obf_30915fb6c7afe620.
		Write([]byte(__obf_ed9fc90c1e7fbdbf))
	return __obf_30915fb6c7afe620.Sum32()
}

func (__obf_aa40f13747012419 *Consistent[M]) __obf_e85602afe7615e81() {
	__obf_71de475ac5d06ca3 := __obf_aa40f13747012419.
		//reallocate if we're holding on to too much (1/4th)
		__obf_b75bdb0b287cd566[:0]

	if cap(__obf_aa40f13747012419.__obf_b75bdb0b287cd566)/(__obf_aa40f13747012419.NumberOfReplicas*4) > len(__obf_aa40f13747012419.__obf_7d563a82a34c175d) {
		__obf_71de475ac5d06ca3 = nil
	}
	for __obf_56227607a6f80707 := range __obf_aa40f13747012419.__obf_7d563a82a34c175d {
		__obf_71de475ac5d06ca3 = append(__obf_71de475ac5d06ca3, __obf_56227607a6f80707)
	}
	sort.Sort(__obf_71de475ac5d06ca3)
	__obf_aa40f13747012419.__obf_b75bdb0b287cd566 = __obf_71de475ac5d06ca3
}

func __obf_5af7c0cbcdd4f639[M Member](__obf_7ad317d3d7d0a731 []M, __obf_85b538c93099dc25 M) bool {
	for _, __obf_910264ae455d2054 := range __obf_7ad317d3d7d0a731 {
		if __obf_910264ae455d2054.String() == __obf_85b538c93099dc25.String() {
			return true
		}
	}
	return false
}
