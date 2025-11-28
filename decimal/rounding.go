// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_af53385fca67d325

type __obf_8f596f0fa62f0bb1 struct {
	__obf_cdbd5d47d10193b8 uint
	__obf_8259f1621247e8d5 uint
	__obf_5ccd0fbc88864cea int
}

var __obf_e7a2312c3f2fd7ba = __obf_8f596f0fa62f0bb1{23, 8, -127}
var __obf_0d845b58fe68fed9 = __obf_8f596f0fa62f0bb1{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_acb82e38a55277ae(__obf_c8d718debc5f6382 *__obf_af53385fca67d325, __obf_2a9e5a4d768e50b1 uint64, __obf_e7725bcd1ee3b1c3 int, __obf_0c69b8da42f5c255 *__obf_8f596f0fa62f0bb1) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_2a9e5a4d768e50b1 == 0 {
		__obf_c8d718debc5f6382.__obf_b2da5b1e27e361da = 0
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
	__obf_2ef7b4722bf9751f := __obf_0c69b8da42f5c255.__obf_5ccd0fbc88864cea + 1 // minimum possible exponent
	if __obf_e7725bcd1ee3b1c3 > __obf_2ef7b4722bf9751f && 332*(__obf_c8d718debc5f6382.__obf_27d38fb675e867eb-__obf_c8d718debc5f6382.__obf_b2da5b1e27e361da) >= 100*(__obf_e7725bcd1ee3b1c3-int(__obf_0c69b8da42f5c255.__obf_cdbd5d47d10193b8)) {
		// The number is already shortest.
		return
	}

	// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	__obf_90f2c8c29de4cc37 := new(__obf_af53385fca67d325)
	__obf_90f2c8c29de4cc37.Assign(__obf_2a9e5a4d768e50b1*2 + 1)
	__obf_90f2c8c29de4cc37.Shift(__obf_e7725bcd1ee3b1c3 - int(__obf_0c69b8da42f5c255.__obf_cdbd5d47d10193b8) - 1)

	// d = mant << (exp - mantbits)
	// Next lowest floating point number is mant-1 << exp-mantbits,
	// unless mant-1 drops the significant bit and exp is not the minimum exp,
	// in which case the next lowest is mant*2-1 << exp-mantbits-1.
	// Either way, call it mantlo << explo-mantbits.
	// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
	var __obf_9c2badb762cb15ec uint64
	var __obf_fd4e9eb204c609cf int
	if __obf_2a9e5a4d768e50b1 > 1<<__obf_0c69b8da42f5c255.__obf_cdbd5d47d10193b8 || __obf_e7725bcd1ee3b1c3 == __obf_2ef7b4722bf9751f {
		__obf_9c2badb762cb15ec = __obf_2a9e5a4d768e50b1 - 1
		__obf_fd4e9eb204c609cf = __obf_e7725bcd1ee3b1c3
	} else {
		__obf_9c2badb762cb15ec = __obf_2a9e5a4d768e50b1*2 - 1
		__obf_fd4e9eb204c609cf = __obf_e7725bcd1ee3b1c3 - 1
	}
	__obf_21a2467097d2d16c := new(__obf_af53385fca67d325)
	__obf_21a2467097d2d16c.Assign(__obf_9c2badb762cb15ec*2 + 1)
	__obf_21a2467097d2d16c.Shift(__obf_fd4e9eb204c609cf - int(__obf_0c69b8da42f5c255.__obf_cdbd5d47d10193b8) - 1)

	// The upper and lower bounds are possible outputs only if
	// the original mantissa is even, so that IEEE round-to-even
	// would round to the original mantissa and not the neighbors.
	__obf_bd1515ae4889f05c := __obf_2a9e5a4d768e50b1%2 == 0

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
	var __obf_a0d3a15a57b8f534 uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_9af8085e2f463f67 := 0; ; __obf_9af8085e2f463f67++ {
		// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_e93a4e8e37d31be7 := __obf_9af8085e2f463f67 - __obf_90f2c8c29de4cc37.__obf_27d38fb675e867eb + __obf_c8d718debc5f6382.__obf_27d38fb675e867eb
		if __obf_e93a4e8e37d31be7 >= __obf_c8d718debc5f6382.__obf_b2da5b1e27e361da {
			break
		}
		__obf_339e1b9225d39a73 := __obf_9af8085e2f463f67 - __obf_90f2c8c29de4cc37.__obf_27d38fb675e867eb + __obf_21a2467097d2d16c.__obf_27d38fb675e867eb
		__obf_e06b79e322ecf9f3 := byte('0') // lower digit
		if __obf_339e1b9225d39a73 >= 0 && __obf_339e1b9225d39a73 < __obf_21a2467097d2d16c.__obf_b2da5b1e27e361da {
			__obf_e06b79e322ecf9f3 = __obf_21a2467097d2d16c.__obf_c8d718debc5f6382[__obf_339e1b9225d39a73]
		}
		__obf_4931dd9c612fb4d2 := byte('0') // middle digit
		if __obf_e93a4e8e37d31be7 >= 0 {
			__obf_4931dd9c612fb4d2 = __obf_c8d718debc5f6382.__obf_c8d718debc5f6382[__obf_e93a4e8e37d31be7]
		}
		__obf_8a366d706524a312 := byte('0') // upper digit
		if __obf_9af8085e2f463f67 < __obf_90f2c8c29de4cc37.__obf_b2da5b1e27e361da {
			__obf_8a366d706524a312 = __obf_90f2c8c29de4cc37.__obf_c8d718debc5f6382[__obf_9af8085e2f463f67]
		}

		// Okay to round down (truncate) if lower has a different digit
		// or if lower is inclusive and is exactly the result of rounding
		// down (i.e., and we have reached the final digit of lower).
		__obf_590fa13995ac33eb := __obf_e06b79e322ecf9f3 != __obf_4931dd9c612fb4d2 || __obf_bd1515ae4889f05c && __obf_339e1b9225d39a73+1 == __obf_21a2467097d2d16c.__obf_b2da5b1e27e361da

		switch {
		case __obf_a0d3a15a57b8f534 == 0 && __obf_4931dd9c612fb4d2+1 < __obf_8a366d706524a312:
			// Example:
			// m = 12345xxx
			// u = 12347xxx
			__obf_a0d3a15a57b8f534 = 2
		case __obf_a0d3a15a57b8f534 == 0 && __obf_4931dd9c612fb4d2 != __obf_8a366d706524a312:
			// Example:
			// m = 12345xxx
			// u = 12346xxx
			__obf_a0d3a15a57b8f534 = 1
		case __obf_a0d3a15a57b8f534 == 1 && (__obf_4931dd9c612fb4d2 != '9' || __obf_8a366d706524a312 != '0'):
			// Example:
			// m = 1234598x
			// u = 1234600x
			__obf_a0d3a15a57b8f534 = 2
		}
		// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_38c66db458d88c02 := __obf_a0d3a15a57b8f534 > 0 && (__obf_bd1515ae4889f05c || __obf_a0d3a15a57b8f534 > 1 || __obf_9af8085e2f463f67+1 < __obf_90f2c8c29de4cc37.__obf_b2da5b1e27e361da)

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		switch {
		case __obf_590fa13995ac33eb && __obf_38c66db458d88c02:
			__obf_c8d718debc5f6382.Round(__obf_e93a4e8e37d31be7 + 1)
			return
		case __obf_590fa13995ac33eb:
			__obf_c8d718debc5f6382.RoundDown(__obf_e93a4e8e37d31be7 + 1)
			return
		case __obf_38c66db458d88c02:
			__obf_c8d718debc5f6382.RoundUp(__obf_e93a4e8e37d31be7 + 1)
			return
		}
	}
}
