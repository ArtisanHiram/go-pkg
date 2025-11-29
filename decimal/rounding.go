// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_4cc30aac3ca67486

type __obf_4db791a44d1e22db struct {
	__obf_85a8f036d45ff996 uint
	__obf_d0cd6e9a2f63cb34 uint
	__obf_d5618d67c0a6e865 int
}

var __obf_5344c617ea3140d5 = __obf_4db791a44d1e22db{23, 8, -127}
var __obf_f772cba36824086f = __obf_4db791a44d1e22db{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_711c33dc022ef1f9(__obf_aafc5e233dfea784 *__obf_4cc30aac3ca67486, __obf_7b4a9ee3891ebcd0 uint64, __obf_10eec1d7f8ad279f int, __obf_3313f66003031fa5 *__obf_4db791a44d1e22db,) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_7b4a9ee3891ebcd0 == 0 {
		__obf_aafc5e233dfea784.__obf_5c82a3d658ffb3ad = 0
		return
	}
	__obf_b1aa1a164ad22c97 :=// Compute upper and lower such that any decimal number
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
	__obf_3313f66003031fa5. // minimum possible exponent
	__obf_d5618d67c0a6e865 + 1
	if __obf_10eec1d7f8ad279f > __obf_b1aa1a164ad22c97 && 332*(__obf_aafc5e233dfea784.__obf_2b78448fb2f1bffe-__obf_aafc5e233dfea784.__obf_5c82a3d658ffb3ad) >= 100*(__obf_10eec1d7f8ad279f-int(__obf_3313f66003031fa5.
	// The number is already shortest.
	__obf_85a8f036d45ff996)) {

		return
	}
	__obf_ad9ffb383d5336c1 :=// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	new(__obf_4cc30aac3ca67486)
	__obf_ad9ffb383d5336c1.
		Assign(__obf_7b4a9ee3891ebcd0*2 + 1)
	__obf_ad9ffb383d5336c1.
		Shift(__obf_10eec1d7f8ad279f - int(__obf_3313f66003031fa5.

		// d = mant << (exp - mantbits)
		// Next lowest floating point number is mant-1 << exp-mantbits,
		// unless mant-1 drops the significant bit and exp is not the minimum exp,
		// in which case the next lowest is mant*2-1 << exp-mantbits-1.
		// Either way, call it mantlo << explo-mantbits.
		// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
		__obf_85a8f036d45ff996) - 1)

	var __obf_0757f9b10ade6ce3 uint64
	var __obf_f02efe1224fbbc80 int
	if __obf_7b4a9ee3891ebcd0 > 1<<__obf_3313f66003031fa5.__obf_85a8f036d45ff996 || __obf_10eec1d7f8ad279f == __obf_b1aa1a164ad22c97 {
		__obf_0757f9b10ade6ce3 = __obf_7b4a9ee3891ebcd0 - 1
		__obf_f02efe1224fbbc80 = __obf_10eec1d7f8ad279f
	} else {
		__obf_0757f9b10ade6ce3 = __obf_7b4a9ee3891ebcd0*2 - 1
		__obf_f02efe1224fbbc80 = __obf_10eec1d7f8ad279f - 1
	}
	__obf_1409a098d37feffc := new(__obf_4cc30aac3ca67486)
	__obf_1409a098d37feffc.
		Assign(__obf_0757f9b10ade6ce3*2 + 1)
	__obf_1409a098d37feffc.
		Shift(__obf_f02efe1224fbbc80 - int(__obf_3313f66003031fa5.

		// The upper and lower bounds are possible outputs only if
		// the original mantissa is even, so that IEEE round-to-even
		// would round to the original mantissa and not the neighbors.
		__obf_85a8f036d45ff996) - 1)
	__obf_e53135e894ecae0a := __obf_7b4a9ee3891ebcd0%2 == 0

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
	var __obf_f0db48a02dfea358 uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_c9955f0afdfb107e := 0; ; __obf_c9955f0afdfb107e++ {
		__obf_64036127cbf55b2f :=// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_c9955f0afdfb107e - __obf_ad9ffb383d5336c1.__obf_2b78448fb2f1bffe + __obf_aafc5e233dfea784.__obf_2b78448fb2f1bffe
		if __obf_64036127cbf55b2f >= __obf_aafc5e233dfea784.__obf_5c82a3d658ffb3ad {
			break
		}
		__obf_39643dec006b1b25 := __obf_c9955f0afdfb107e - __obf_ad9ffb383d5336c1.__obf_2b78448fb2f1bffe + __obf_1409a098d37feffc.__obf_2b78448fb2f1bffe
		// lower digit
		__obf_d971c9ed104c8581 := byte('0')
		if __obf_39643dec006b1b25 >= 0 && __obf_39643dec006b1b25 < __obf_1409a098d37feffc.__obf_5c82a3d658ffb3ad {
			__obf_d971c9ed104c8581 = __obf_1409a098d37feffc.__obf_aafc5e233dfea784[__obf_39643dec006b1b25]
		}
		__obf_ade33d4daf3b3fc1 := byte('0') // middle digit
		if __obf_64036127cbf55b2f >= 0 {
			__obf_ade33d4daf3b3fc1 = __obf_aafc5e233dfea784.__obf_aafc5e233dfea784[__obf_64036127cbf55b2f]
		}
		__obf_51e2f4b1c3fbb91f := byte('0') // upper digit
		if __obf_c9955f0afdfb107e < __obf_ad9ffb383d5336c1.__obf_5c82a3d658ffb3ad {
			__obf_51e2f4b1c3fbb91f = __obf_ad9ffb383d5336c1.

			// Okay to round down (truncate) if lower has a different digit
			// or if lower is inclusive and is exactly the result of rounding
			// down (i.e., and we have reached the final digit of lower).
			__obf_aafc5e233dfea784[__obf_c9955f0afdfb107e]
		}
		__obf_96433fd44d15c431 := __obf_d971c9ed104c8581 != __obf_ade33d4daf3b3fc1 || __obf_e53135e894ecae0a && __obf_39643dec006b1b25+1 == __obf_1409a098d37feffc.__obf_5c82a3d658ffb3ad

		switch {
		case __obf_f0db48a02dfea358 == 0 && __obf_ade33d4daf3b3fc1+1 < __obf_51e2f4b1c3fbb91f:
			__obf_f0db48a02dfea358 =// Example:
			// m = 12345xxx
			// u = 12347xxx
			2
		case __obf_f0db48a02dfea358 == 0 && __obf_ade33d4daf3b3fc1 != __obf_51e2f4b1c3fbb91f:
			__obf_f0db48a02dfea358 =// Example:
			// m = 12345xxx
			// u = 12346xxx
			1
		case __obf_f0db48a02dfea358 == 1 && (__obf_ade33d4daf3b3fc1 != '9' || __obf_51e2f4b1c3fbb91f != '0'):
			__obf_f0db48a02dfea358 =// Example:
			// m = 1234598x
			// u = 1234600x
			2
		}
		__obf_d0b9dbcb012a8d66 :=// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_f0db48a02dfea358 > 0 && (__obf_e53135e894ecae0a || __obf_f0db48a02dfea358 > 1 || __obf_c9955f0afdfb107e+1 < __obf_ad9ffb383d5336c1.

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		__obf_5c82a3d658ffb3ad)

		switch {
		case __obf_96433fd44d15c431 && __obf_d0b9dbcb012a8d66:
			__obf_aafc5e233dfea784.
				Round(__obf_64036127cbf55b2f + 1)
			return
		case __obf_96433fd44d15c431:
			__obf_aafc5e233dfea784.
				RoundDown(__obf_64036127cbf55b2f + 1)
			return
		case __obf_d0b9dbcb012a8d66:
			__obf_aafc5e233dfea784.
				RoundUp(__obf_64036127cbf55b2f + 1)
			return
		}
	}
}
