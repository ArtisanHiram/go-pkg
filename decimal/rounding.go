// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_0962dc77c6b6239b

type __obf_e46206765e5dbb45 struct {
	__obf_f89d4a69e09b5dc8 uint
	__obf_44fa0a4671bd9290 uint
	__obf_279e6fa356ffbd95 int
}

var __obf_27426ff0b890a183 = __obf_e46206765e5dbb45{23, 8, -127}
var __obf_b3a4121c853c30da = __obf_e46206765e5dbb45{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_f513f8a0934968e3(__obf_d09f058c0a390c93 *__obf_0962dc77c6b6239b, __obf_a54bc9ed25a8de66 uint64, __obf_406325483a83b1fa int, __obf_8b0328b22b58eb64 *__obf_e46206765e5dbb45) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_a54bc9ed25a8de66 == 0 {
		__obf_d09f058c0a390c93.__obf_ea04243a9869d25e = 0
		return
	}

	// Compute upper and lower such that any decimal number
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
	__obf_0bfdcbfcf4dc323e := __obf_8b0328b22b58eb64.__obf_279e6fa356ffbd95 + 1 // minimum possible exponent
	if __obf_406325483a83b1fa > __obf_0bfdcbfcf4dc323e && 332*(__obf_d09f058c0a390c93.__obf_cb19c1fd6e41a0d7-__obf_d09f058c0a390c93.__obf_ea04243a9869d25e) >= 100*(__obf_406325483a83b1fa-int(__obf_8b0328b22b58eb64.__obf_f89d4a69e09b5dc8)) {
		// The number is already shortest.
		return
	}

	// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	__obf_7d5427ccd092cf6a := new(__obf_0962dc77c6b6239b)
	__obf_7d5427ccd092cf6a.Assign(__obf_a54bc9ed25a8de66*2 + 1)
	__obf_7d5427ccd092cf6a.Shift(__obf_406325483a83b1fa - int(__obf_8b0328b22b58eb64.__obf_f89d4a69e09b5dc8) - 1)

	// d = mant << (exp - mantbits)
	// Next lowest floating point number is mant-1 << exp-mantbits,
	// unless mant-1 drops the significant bit and exp is not the minimum exp,
	// in which case the next lowest is mant*2-1 << exp-mantbits-1.
	// Either way, call it mantlo << explo-mantbits.
	// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
	var __obf_784e6fa10477d8d2 uint64
	var __obf_08114cd5bfe63ddc int
	if __obf_a54bc9ed25a8de66 > 1<<__obf_8b0328b22b58eb64.__obf_f89d4a69e09b5dc8 || __obf_406325483a83b1fa == __obf_0bfdcbfcf4dc323e {
		__obf_784e6fa10477d8d2 = __obf_a54bc9ed25a8de66 - 1
		__obf_08114cd5bfe63ddc = __obf_406325483a83b1fa
	} else {
		__obf_784e6fa10477d8d2 = __obf_a54bc9ed25a8de66*2 - 1
		__obf_08114cd5bfe63ddc = __obf_406325483a83b1fa - 1
	}
	__obf_799c435fe541b169 := new(__obf_0962dc77c6b6239b)
	__obf_799c435fe541b169.Assign(__obf_784e6fa10477d8d2*2 + 1)
	__obf_799c435fe541b169.Shift(__obf_08114cd5bfe63ddc - int(__obf_8b0328b22b58eb64.__obf_f89d4a69e09b5dc8) - 1)

	// The upper and lower bounds are possible outputs only if
	// the original mantissa is even, so that IEEE round-to-even
	// would round to the original mantissa and not the neighbors.
	__obf_1b954a927c83a39e := __obf_a54bc9ed25a8de66%2 == 0

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
	var __obf_5117cbcb9b277049 uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_5a3a71d25725e3bc := 0; ; __obf_5a3a71d25725e3bc++ {
		// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_30dafc63db10db19 := __obf_5a3a71d25725e3bc - __obf_7d5427ccd092cf6a.__obf_cb19c1fd6e41a0d7 + __obf_d09f058c0a390c93.__obf_cb19c1fd6e41a0d7
		if __obf_30dafc63db10db19 >= __obf_d09f058c0a390c93.__obf_ea04243a9869d25e {
			break
		}
		__obf_b47a577ef017b433 := __obf_5a3a71d25725e3bc - __obf_7d5427ccd092cf6a.__obf_cb19c1fd6e41a0d7 + __obf_799c435fe541b169.__obf_cb19c1fd6e41a0d7
		__obf_9a4da8ddc71b9383 := byte('0') // lower digit
		if __obf_b47a577ef017b433 >= 0 && __obf_b47a577ef017b433 < __obf_799c435fe541b169.__obf_ea04243a9869d25e {
			__obf_9a4da8ddc71b9383 = __obf_799c435fe541b169.__obf_d09f058c0a390c93[__obf_b47a577ef017b433]
		}
		__obf_b041e2c85e138a64 := byte('0') // middle digit
		if __obf_30dafc63db10db19 >= 0 {
			__obf_b041e2c85e138a64 = __obf_d09f058c0a390c93.__obf_d09f058c0a390c93[__obf_30dafc63db10db19]
		}
		__obf_8c956dfc0c2456b1 := byte('0') // upper digit
		if __obf_5a3a71d25725e3bc < __obf_7d5427ccd092cf6a.__obf_ea04243a9869d25e {
			__obf_8c956dfc0c2456b1 = __obf_7d5427ccd092cf6a.__obf_d09f058c0a390c93[__obf_5a3a71d25725e3bc]
		}

		// Okay to round down (truncate) if lower has a different digit
		// or if lower is inclusive and is exactly the result of rounding
		// down (i.e., and we have reached the final digit of lower).
		__obf_f7cef50b7a10fd39 := __obf_9a4da8ddc71b9383 != __obf_b041e2c85e138a64 || __obf_1b954a927c83a39e && __obf_b47a577ef017b433+1 == __obf_799c435fe541b169.__obf_ea04243a9869d25e

		switch {
		case __obf_5117cbcb9b277049 == 0 && __obf_b041e2c85e138a64+1 < __obf_8c956dfc0c2456b1:
			// Example:
			// m = 12345xxx
			// u = 12347xxx
			__obf_5117cbcb9b277049 = 2
		case __obf_5117cbcb9b277049 == 0 && __obf_b041e2c85e138a64 != __obf_8c956dfc0c2456b1:
			// Example:
			// m = 12345xxx
			// u = 12346xxx
			__obf_5117cbcb9b277049 = 1
		case __obf_5117cbcb9b277049 == 1 && (__obf_b041e2c85e138a64 != '9' || __obf_8c956dfc0c2456b1 != '0'):
			// Example:
			// m = 1234598x
			// u = 1234600x
			__obf_5117cbcb9b277049 = 2
		}
		// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_aeb370d7fba30acb := __obf_5117cbcb9b277049 > 0 && (__obf_1b954a927c83a39e || __obf_5117cbcb9b277049 > 1 || __obf_5a3a71d25725e3bc+1 < __obf_7d5427ccd092cf6a.__obf_ea04243a9869d25e)

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		switch {
		case __obf_f7cef50b7a10fd39 && __obf_aeb370d7fba30acb:
			__obf_d09f058c0a390c93.Round(__obf_30dafc63db10db19 + 1)
			return
		case __obf_f7cef50b7a10fd39:
			__obf_d09f058c0a390c93.RoundDown(__obf_30dafc63db10db19 + 1)
			return
		case __obf_aeb370d7fba30acb:
			__obf_d09f058c0a390c93.RoundUp(__obf_30dafc63db10db19 + 1)
			return
		}
	}
}
