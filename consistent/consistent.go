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
package __obf_38fc7a436cbc0f56 // import "stathat.com/c/consistent"

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

type __obf_7ba9f0825579b1a0 []uint32

// Len returns the length of the uints array.
func (__obf_cbc1ca5a8d4ccdf0 __obf_7ba9f0825579b1a0) Len() int { return len(__obf_cbc1ca5a8d4ccdf0) }

// Less returns true if element i is less than element j.
func (__obf_cbc1ca5a8d4ccdf0 __obf_7ba9f0825579b1a0) Less(__obf_1fd367ef8f4a1660, __obf_1afcd43b69f81f2f int) bool {
	return __obf_cbc1ca5a8d4ccdf0[__obf_1fd367ef8f4a1660] < __obf_cbc1ca5a8d4ccdf0[__obf_1afcd43b69f81f2f]
}

// Swap exchanges elements i and j.
func (__obf_cbc1ca5a8d4ccdf0 __obf_7ba9f0825579b1a0) Swap(__obf_1fd367ef8f4a1660, __obf_1afcd43b69f81f2f int) {
	__obf_cbc1ca5a8d4ccdf0[__obf_1fd367ef8f4a1660], __obf_cbc1ca5a8d4ccdf0[__obf_1afcd43b69f81f2f] = __obf_cbc1ca5a8d4ccdf0[__obf_1afcd43b69f81f2f], __obf_cbc1ca5a8d4ccdf0[__obf_1fd367ef8f4a1660]
}

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent[M Member] struct {
	__obf_61b149eb4dc77c12 map[uint32]M
	__obf_a4e0651483591fce map[string]bool
	__obf_c456021f7d1597fe __obf_7ba9f0825579b1a0
	NumberOfReplicas       int
	__obf_1bce143f4f40c872 int64
	__obf_959ebc5d8ede29cc [64]byte
	UseFnv                 bool
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New[M Member]() *Consistent[M] {
	__obf_e287beb4e6aad1d1 := new(Consistent[M])
	__obf_e287beb4e6aad1d1.NumberOfReplicas = 20
	__obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12 = make(map[uint32]M)
	__obf_e287beb4e6aad1d1.__obf_a4e0651483591fce = make(map[string]bool)
	return __obf_e287beb4e6aad1d1
}

// eltKey generates a string key for an element with an index.
func (__obf_e287beb4e6aad1d1 *Consistent[M]) __obf_ea6cbb53f414c78b(__obf_f960d3fa1b496fd7 string, __obf_8bbd5f773daf0213 int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(__obf_8bbd5f773daf0213) + __obf_f960d3fa1b496fd7
}

// Add inserts a string element in the consistent hash.
func (__obf_e287beb4e6aad1d1 *Consistent[M]) Add(__obf_f960d3fa1b496fd7 M) {
	__obf_e287beb4e6aad1d1.Lock()
	defer __obf_e287beb4e6aad1d1.Unlock()
	__obf_e287beb4e6aad1d1.__obf_9eeb07371307cb88(__obf_f960d3fa1b496fd7)
}

// need c.Lock() before calling
func (__obf_e287beb4e6aad1d1 *Consistent[M]) __obf_9eeb07371307cb88(__obf_f960d3fa1b496fd7 M) {
	for __obf_1fd367ef8f4a1660 := 0; __obf_1fd367ef8f4a1660 < __obf_e287beb4e6aad1d1.NumberOfReplicas; __obf_1fd367ef8f4a1660++ {
		__obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12[__obf_e287beb4e6aad1d1.__obf_6780f01db35b26b5(__obf_e287beb4e6aad1d1.__obf_ea6cbb53f414c78b(__obf_f960d3fa1b496fd7.String(), __obf_1fd367ef8f4a1660))] = __obf_f960d3fa1b496fd7
	}
	__obf_e287beb4e6aad1d1.__obf_a4e0651483591fce[__obf_f960d3fa1b496fd7.String()] = true
	__obf_e287beb4e6aad1d1.__obf_f7cadff3405eac05()
	__obf_e287beb4e6aad1d1.__obf_1bce143f4f40c872++
}

// Remove removes an element from the hash.
func (__obf_e287beb4e6aad1d1 *Consistent[M]) Remove(__obf_f960d3fa1b496fd7 M) {
	__obf_e287beb4e6aad1d1.Lock()
	defer __obf_e287beb4e6aad1d1.Unlock()
	__obf_e287beb4e6aad1d1.__obf_c1bc2dead409313a(__obf_f960d3fa1b496fd7.String())
}

// need c.Lock() before calling
func (__obf_e287beb4e6aad1d1 *Consistent[M]) __obf_c1bc2dead409313a(__obf_f960d3fa1b496fd7 string) {
	for __obf_1fd367ef8f4a1660 := 0; __obf_1fd367ef8f4a1660 < __obf_e287beb4e6aad1d1.NumberOfReplicas; __obf_1fd367ef8f4a1660++ {
		delete(__obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12, __obf_e287beb4e6aad1d1.__obf_6780f01db35b26b5(__obf_e287beb4e6aad1d1.__obf_ea6cbb53f414c78b(__obf_f960d3fa1b496fd7, __obf_1fd367ef8f4a1660)))
	}
	delete(__obf_e287beb4e6aad1d1.__obf_a4e0651483591fce, __obf_f960d3fa1b496fd7)
	__obf_e287beb4e6aad1d1.__obf_f7cadff3405eac05()
	__obf_e287beb4e6aad1d1.__obf_1bce143f4f40c872--
}

// Set sets all the elements in the hash.  If there are existing elements not
// present in elts, they will be removed.
func (__obf_e287beb4e6aad1d1 *Consistent[M]) Set(__obf_2ff8bd4ded7d1936 []M) {
	__obf_e287beb4e6aad1d1.Lock()
	defer __obf_e287beb4e6aad1d1.Unlock()
	for __obf_78117545f587262b := range __obf_e287beb4e6aad1d1.__obf_a4e0651483591fce {
		__obf_55eed423f89a0538 := false
		for _, __obf_2eb952e853accbf5 := range __obf_2ff8bd4ded7d1936 {
			if __obf_78117545f587262b == __obf_2eb952e853accbf5.String() {
				__obf_55eed423f89a0538 = true
				break
			}
		}
		if !__obf_55eed423f89a0538 {
			__obf_e287beb4e6aad1d1.__obf_c1bc2dead409313a(__obf_78117545f587262b)
		}
	}
	for _, __obf_2eb952e853accbf5 := range __obf_2ff8bd4ded7d1936 {
		_, __obf_063a317b9db0cfc9 := __obf_e287beb4e6aad1d1.__obf_a4e0651483591fce[__obf_2eb952e853accbf5.String()]
		if __obf_063a317b9db0cfc9 {
			continue
		}
		__obf_e287beb4e6aad1d1.__obf_9eeb07371307cb88(__obf_2eb952e853accbf5)
	}
}

func (__obf_e287beb4e6aad1d1 *Consistent[M]) Members() []string {
	__obf_e287beb4e6aad1d1.RLock()
	defer __obf_e287beb4e6aad1d1.RUnlock()
	var __obf_1ff8e530188db8a5 []string
	for __obf_78117545f587262b := range __obf_e287beb4e6aad1d1.__obf_a4e0651483591fce {
		__obf_1ff8e530188db8a5 = append(__obf_1ff8e530188db8a5, __obf_78117545f587262b)
	}
	return __obf_1ff8e530188db8a5
}

// Get returns an element close to where name hashes to in the circle.
func (__obf_e287beb4e6aad1d1 *Consistent[M]) Get(__obf_248f176fb4ac9302 string) (__obf_a70bc68166ec2d45 M, __obf_dca4c0aff351dbe8 error) {
	__obf_e287beb4e6aad1d1.RLock()
	defer __obf_e287beb4e6aad1d1.RUnlock()
	if len(__obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12) == 0 {
		__obf_dca4c0aff351dbe8 = ErrEmptyCircle
		return
	}
	__obf_622100f206720bb9 := __obf_e287beb4e6aad1d1.__obf_6780f01db35b26b5(__obf_248f176fb4ac9302)
	__obf_1fd367ef8f4a1660 := __obf_e287beb4e6aad1d1.__obf_ccd6048cd4101a56(__obf_622100f206720bb9)
	__obf_a70bc68166ec2d45 = __obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12[__obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe[__obf_1fd367ef8f4a1660]]
	return
}

func (__obf_e287beb4e6aad1d1 *Consistent[M]) __obf_ccd6048cd4101a56(__obf_622100f206720bb9 uint32) (__obf_1fd367ef8f4a1660 int) {
	__obf_6cd30f95e369992c := func(__obf_cbc1ca5a8d4ccdf0 int) bool {
		return __obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe[__obf_cbc1ca5a8d4ccdf0] > __obf_622100f206720bb9
	}
	__obf_1fd367ef8f4a1660 = sort.Search(len(__obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe), __obf_6cd30f95e369992c)
	if __obf_1fd367ef8f4a1660 >= len(__obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe) {
		__obf_1fd367ef8f4a1660 = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (__obf_e287beb4e6aad1d1 *Consistent[M]) GetTwo(__obf_248f176fb4ac9302 string) (__obf_27f5d0309e6ba9eb M, __obf_9cf8f25e0d557843 M, __obf_dca4c0aff351dbe8 error) {
	__obf_e287beb4e6aad1d1.RLock()
	defer __obf_e287beb4e6aad1d1.RUnlock()
	if len(__obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12) == 0 {
		__obf_dca4c0aff351dbe8 = ErrEmptyCircle
		return
	}
	__obf_622100f206720bb9 := __obf_e287beb4e6aad1d1.__obf_6780f01db35b26b5(__obf_248f176fb4ac9302)
	__obf_1fd367ef8f4a1660 := __obf_e287beb4e6aad1d1.__obf_ccd6048cd4101a56(__obf_622100f206720bb9)
	__obf_27f5d0309e6ba9eb = __obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12[__obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe[__obf_1fd367ef8f4a1660]]

	if __obf_e287beb4e6aad1d1.__obf_1bce143f4f40c872 == 1 {
		return
	}

	__obf_6376460501e53632 := __obf_1fd367ef8f4a1660
	for __obf_1fd367ef8f4a1660 = __obf_6376460501e53632 + 1; __obf_1fd367ef8f4a1660 != __obf_6376460501e53632; __obf_1fd367ef8f4a1660++ {
		if __obf_1fd367ef8f4a1660 >= len(__obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe) {
			__obf_1fd367ef8f4a1660 = 0
		}
		__obf_9cf8f25e0d557843 = __obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12[__obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe[__obf_1fd367ef8f4a1660]]
		if __obf_9cf8f25e0d557843.String() != __obf_27f5d0309e6ba9eb.String() {
			break
		}
	}
	return __obf_27f5d0309e6ba9eb, __obf_9cf8f25e0d557843, nil
}

// GetN returns the N closest distinct elements to the name input in the circle.
func (__obf_e287beb4e6aad1d1 *Consistent[M]) GetN(__obf_248f176fb4ac9302 string, __obf_e71ef9f2ccd9d4ad int) (__obf_a70bc68166ec2d45 []M, __obf_dca4c0aff351dbe8 error) {
	__obf_e287beb4e6aad1d1.RLock()
	defer __obf_e287beb4e6aad1d1.RUnlock()

	if len(__obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12) == 0 {
		__obf_dca4c0aff351dbe8 = ErrEmptyCircle
		return
	}

	if __obf_e287beb4e6aad1d1.__obf_1bce143f4f40c872 < int64(__obf_e71ef9f2ccd9d4ad) {
		__obf_e71ef9f2ccd9d4ad = int(__obf_e287beb4e6aad1d1.__obf_1bce143f4f40c872)
	}

	var (
		__obf_622100f206720bb9 = __obf_e287beb4e6aad1d1.__obf_6780f01db35b26b5(__obf_248f176fb4ac9302)
		__obf_1fd367ef8f4a1660 = __obf_e287beb4e6aad1d1.__obf_ccd6048cd4101a56(__obf_622100f206720bb9)
		__obf_6376460501e53632 = __obf_1fd367ef8f4a1660
		__obf_ec37e7e0e874f451 = __obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12[__obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe[__obf_1fd367ef8f4a1660]]
	)
	__obf_a70bc68166ec2d45 = make([]M, 0, __obf_e71ef9f2ccd9d4ad)
	__obf_a70bc68166ec2d45 = append(__obf_a70bc68166ec2d45, __obf_ec37e7e0e874f451)

	if len(__obf_a70bc68166ec2d45) == __obf_e71ef9f2ccd9d4ad {
		return __obf_a70bc68166ec2d45, nil
	}

	for __obf_1fd367ef8f4a1660 = __obf_6376460501e53632 + 1; __obf_1fd367ef8f4a1660 != __obf_6376460501e53632; __obf_1fd367ef8f4a1660++ {
		if __obf_1fd367ef8f4a1660 >= len(__obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe) {
			__obf_1fd367ef8f4a1660 = 0
		}
		__obf_ec37e7e0e874f451 = __obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12[__obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe[__obf_1fd367ef8f4a1660]]
		if !__obf_1ac762003054a190(__obf_a70bc68166ec2d45, __obf_ec37e7e0e874f451) {
			__obf_a70bc68166ec2d45 = append(__obf_a70bc68166ec2d45, __obf_ec37e7e0e874f451)
		}
		if len(__obf_a70bc68166ec2d45) == __obf_e71ef9f2ccd9d4ad {
			break
		}
	}

	return __obf_a70bc68166ec2d45, nil
}

func (__obf_e287beb4e6aad1d1 *Consistent[M]) __obf_6780f01db35b26b5(__obf_622100f206720bb9 string) uint32 {
	if __obf_e287beb4e6aad1d1.UseFnv {
		return __obf_e287beb4e6aad1d1.__obf_f12e4083df8d0860(__obf_622100f206720bb9)
	}
	return __obf_e287beb4e6aad1d1.__obf_e0f855ed45a864d9(__obf_622100f206720bb9)
}

func (__obf_e287beb4e6aad1d1 *Consistent[M]) __obf_e0f855ed45a864d9(__obf_622100f206720bb9 string) uint32 {
	if len(__obf_622100f206720bb9) < 64 {
		var __obf_959ebc5d8ede29cc [64]byte
		copy(__obf_959ebc5d8ede29cc[:], __obf_622100f206720bb9)
		return crc32.ChecksumIEEE(__obf_959ebc5d8ede29cc[:len(__obf_622100f206720bb9)])
	}
	return crc32.ChecksumIEEE([]byte(__obf_622100f206720bb9))
}

func (__obf_e287beb4e6aad1d1 *Consistent[M]) __obf_f12e4083df8d0860(__obf_622100f206720bb9 string) uint32 {
	__obf_392f060d13d1ec89 := fnv.New32a()
	__obf_392f060d13d1ec89.Write([]byte(__obf_622100f206720bb9))
	return __obf_392f060d13d1ec89.Sum32()
}

func (__obf_e287beb4e6aad1d1 *Consistent[M]) __obf_f7cadff3405eac05() {
	__obf_b6137ed37fb83b25 := __obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe[:0]
	//reallocate if we're holding on to too much (1/4th)
	if cap(__obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe)/(__obf_e287beb4e6aad1d1.NumberOfReplicas*4) > len(__obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12) {
		__obf_b6137ed37fb83b25 = nil
	}
	for __obf_78117545f587262b := range __obf_e287beb4e6aad1d1.__obf_61b149eb4dc77c12 {
		__obf_b6137ed37fb83b25 = append(__obf_b6137ed37fb83b25, __obf_78117545f587262b)
	}
	sort.Sort(__obf_b6137ed37fb83b25)
	__obf_e287beb4e6aad1d1.__obf_c456021f7d1597fe = __obf_b6137ed37fb83b25
}

func __obf_1ac762003054a190[M Member](__obf_5976fe0667497184 []M, __obf_c91a82afd5cfbad2 M) bool {
	for _, __obf_1ff8e530188db8a5 := range __obf_5976fe0667497184 {
		if __obf_1ff8e530188db8a5.String() == __obf_c91a82afd5cfbad2.String() {
			return true
		}
	}
	return false
}
