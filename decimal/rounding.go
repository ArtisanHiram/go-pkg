// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_a5d0662d16479f1d

type __obf_ff127edd42fa495d struct {
	__obf_c0ccdfc479ac2e96 uint
	__obf_8ffbe185d6c7002e uint
	__obf_0570419913d32c75 int
}

var __obf_8dea1c634edf06b9 = __obf_ff127edd42fa495d{23, 8, -127}
var __obf_3d88be66e25681be = __obf_ff127edd42fa495d{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_1e54fb2a7fbae193(__obf_5debc2a983e87a03 *__obf_a5d0662d16479f1d, __obf_7ede9834c2850305 uint64, __obf_d5f3440c0176542c int, __obf_24b406b733eb006b *__obf_ff127edd42fa495d,) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_7ede9834c2850305 == 0 {
		__obf_5debc2a983e87a03.__obf_d5ebe437c13ed047 = 0
		return
	}
	__obf_488a0a68434c0581 :=// Compute upper and lower such that any decimal number
	// between upper and lower (possibly inclusive)
	// will round to the original floating point number.

	// We may see at once that the number is already shortest.
	//
	// Suppose d is not denormal, so that 2^exp <= d < 10^dp.
	// The closest shorter number is at least 10^(dp-nd) away.
	// The lower/upper bounds computed below are at distance
	// at most 2^(exp-mantbits).
	//
	// So the number is already shortest if 10^(dp-nd) > 2^(exp-mantbits),
	// or equivalently log2(10)*(dp-nd) > exp-mantbits.
	// It is true if 332/100*(dp-nd) >= exp-mantbits (log2(10) > 3.32).
	__obf_24b406b733eb006b. // minimum possible exponent
	__obf_0570419913d32c75 + 1
	if __obf_d5f3440c0176542c > __obf_488a0a68434c0581 && 332*(__obf_5debc2a983e87a03.__obf_4e162ef443f77409-__obf_5debc2a983e87a03.__obf_d5ebe437c13ed047) >= 100*(__obf_d5f3440c0176542c-int(__obf_24b406b733eb006b.
	// The number is already shortest.
	__obf_c0ccdfc479ac2e96)) {

		return
	}
	__obf_8b2afe27ede33cd5 :=// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	new(__obf_a5d0662d16479f1d)
	__obf_8b2afe27ede33cd5.
		Assign(__obf_7ede9834c2850305*2 + 1)
	__obf_8b2afe27ede33cd5.
		Shift(__obf_d5f3440c0176542c - int(__obf_24b406b733eb006b.

		// d = mant << (exp - mantbits)
		// Next lowest floating point number is mant-1 << exp-mantbits,
		// unless mant-1 drops the significant bit and exp is not the minimum exp,
		// in which case the next lowest is mant*2-1 << exp-mantbits-1.
		// Either way, call it mantlo << explo-mantbits.
		// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
		__obf_c0ccdfc479ac2e96) - 1)

	var __obf_29bc92412c84f450 uint64
	var __obf_2281d268fa9095df int
	if __obf_7ede9834c2850305 > 1<<__obf_24b406b733eb006b.__obf_c0ccdfc479ac2e96 || __obf_d5f3440c0176542c == __obf_488a0a68434c0581 {
		__obf_29bc92412c84f450 = __obf_7ede9834c2850305 - 1
		__obf_2281d268fa9095df = __obf_d5f3440c0176542c
	} else {
		__obf_29bc92412c84f450 = __obf_7ede9834c2850305*2 - 1
		__obf_2281d268fa9095df = __obf_d5f3440c0176542c - 1
	}
	__obf_649ca2858bee3f79 := new(__obf_a5d0662d16479f1d)
	__obf_649ca2858bee3f79.
		Assign(__obf_29bc92412c84f450*2 + 1)
	__obf_649ca2858bee3f79.
		Shift(__obf_2281d268fa9095df - int(__obf_24b406b733eb006b.

		// The upper and lower bounds are possible outputs only if
		// the original mantissa is even, so that IEEE round-to-even
		// would round to the original mantissa and not the neighbors.
		__obf_c0ccdfc479ac2e96) - 1)
	__obf_5ef4a655e080804f := __obf_7ede9834c2850305%2 == 0

	// As we walk the digits we want to know whether rounding up would fall
	// within the upper bound. This is tracked by upperdelta:
	//
	// If upperdelta == 0, the digits of d and upper are the same so far.
	//
	// If upperdelta == 1, we saw a difference of 1 between d and upper on a
	// previous digit and subsequently only 9s for d and 0s for upper.
	// (Thus rounding up may fall outside the bound, if it is exclusive.)
	//
	// If upperdelta == 2, then the difference is greater than 1
	// and we know that rounding up falls within the bound.
	var __obf_44d2ba6749a1da02 uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_d1cb4adc28b22184 := 0; ; __obf_d1cb4adc28b22184++ {
		__obf_694e52f88a3d43cb :=// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_d1cb4adc28b22184 - __obf_8b2afe27ede33cd5.__obf_4e162ef443f77409 + __obf_5debc2a983e87a03.__obf_4e162ef443f77409
		if __obf_694e52f88a3d43cb >= __obf_5debc2a983e87a03.__obf_d5ebe437c13ed047 {
			break
		}
		__obf_6ad88b134215ad67 := __obf_d1cb4adc28b22184 - __obf_8b2afe27ede33cd5.__obf_4e162ef443f77409 + __obf_649ca2858bee3f79.__obf_4e162ef443f77409
		// lower digit
		__obf_1b903128c72fcef5 := byte('0')
		if __obf_6ad88b134215ad67 >= 0 && __obf_6ad88b134215ad67 < __obf_649ca2858bee3f79.__obf_d5ebe437c13ed047 {
			__obf_1b903128c72fcef5 = __obf_649ca2858bee3f79.__obf_5debc2a983e87a03[__obf_6ad88b134215ad67]
		}
		__obf_dfa47acb5d50aedb := byte('0') // middle digit
		if __obf_694e52f88a3d43cb >= 0 {
			__obf_dfa47acb5d50aedb = __obf_5debc2a983e87a03.__obf_5debc2a983e87a03[__obf_694e52f88a3d43cb]
		}
		__obf_735615af25fd6fe2 := byte('0') // upper digit
		if __obf_d1cb4adc28b22184 < __obf_8b2afe27ede33cd5.__obf_d5ebe437c13ed047 {
			__obf_735615af25fd6fe2 = __obf_8b2afe27ede33cd5.

			// Okay to round down (truncate) if lower has a different digit
			// or if lower is inclusive and is exactly the result of rounding
			// down (i.e., and we have reached the final digit of lower).
			__obf_5debc2a983e87a03[__obf_d1cb4adc28b22184]
		}
		__obf_f459f09d7f8c9203 := __obf_1b903128c72fcef5 != __obf_dfa47acb5d50aedb || __obf_5ef4a655e080804f && __obf_6ad88b134215ad67+1 == __obf_649ca2858bee3f79.__obf_d5ebe437c13ed047

		switch {
		case __obf_44d2ba6749a1da02 == 0 && __obf_dfa47acb5d50aedb+1 < __obf_735615af25fd6fe2:
			__obf_44d2ba6749a1da02 =// Example:
			// m = 12345xxx
			// u = 12347xxx
			2
		case __obf_44d2ba6749a1da02 == 0 && __obf_dfa47acb5d50aedb != __obf_735615af25fd6fe2:
			__obf_44d2ba6749a1da02 =// Example:
			// m = 12345xxx
			// u = 12346xxx
			1
		case __obf_44d2ba6749a1da02 == 1 && (__obf_dfa47acb5d50aedb != '9' || __obf_735615af25fd6fe2 != '0'):
			__obf_44d2ba6749a1da02 =// Example:
			// m = 1234598x
			// u = 1234600x
			2
		}
		__obf_6eead63449174932 :=// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_44d2ba6749a1da02 > 0 && (__obf_5ef4a655e080804f || __obf_44d2ba6749a1da02 > 1 || __obf_d1cb4adc28b22184+1 < __obf_8b2afe27ede33cd5.

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		__obf_d5ebe437c13ed047)

		switch {
		case __obf_f459f09d7f8c9203 && __obf_6eead63449174932:
			__obf_5debc2a983e87a03.
				Round(__obf_694e52f88a3d43cb + 1)
			return
		case __obf_f459f09d7f8c9203:
			__obf_5debc2a983e87a03.
				RoundDown(__obf_694e52f88a3d43cb + 1)
			return
		case __obf_6eead63449174932:
			__obf_5debc2a983e87a03.
				RoundUp(__obf_694e52f88a3d43cb + 1)
			return
		}
	}
}
