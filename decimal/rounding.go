// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_adf9c80aee40a9d8

type __obf_84acfd4cf591f3ed struct {
	__obf_415e8862356d404c uint
	__obf_926f241eb4411d53 uint
	__obf_3ff00e3b0682ef4a int
}

var __obf_90dee1d4d7d96500 = __obf_84acfd4cf591f3ed{23, 8, -127}
var __obf_b6eec24024f8885c = __obf_84acfd4cf591f3ed{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_da99d1421cbe52f7(__obf_c4701b3bb28cd2ae *__obf_adf9c80aee40a9d8, __obf_15edc752c636a4a0 uint64, __obf_fed6fd50f0fcc834 int, __obf_cdb3e6cf09ad181f *__obf_84acfd4cf591f3ed) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_15edc752c636a4a0 == 0 {
		__obf_c4701b3bb28cd2ae.__obf_087eecd15d19676e = 0
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
	__obf_b8f9f6f8e536b469 := __obf_cdb3e6cf09ad181f.__obf_3ff00e3b0682ef4a + 1 // minimum possible exponent
	if __obf_fed6fd50f0fcc834 > __obf_b8f9f6f8e536b469 && 332*(__obf_c4701b3bb28cd2ae.__obf_d30835ea2dfed4ba-__obf_c4701b3bb28cd2ae.__obf_087eecd15d19676e) >= 100*(__obf_fed6fd50f0fcc834-int(__obf_cdb3e6cf09ad181f.__obf_415e8862356d404c)) {
		// The number is already shortest.
		return
	}

	// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	__obf_176a296e341c9131 := new(__obf_adf9c80aee40a9d8)
	__obf_176a296e341c9131.Assign(__obf_15edc752c636a4a0*2 + 1)
	__obf_176a296e341c9131.Shift(__obf_fed6fd50f0fcc834 - int(__obf_cdb3e6cf09ad181f.__obf_415e8862356d404c) - 1)

	// d = mant << (exp - mantbits)
	// Next lowest floating point number is mant-1 << exp-mantbits,
	// unless mant-1 drops the significant bit and exp is not the minimum exp,
	// in which case the next lowest is mant*2-1 << exp-mantbits-1.
	// Either way, call it mantlo << explo-mantbits.
	// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
	var __obf_506f449fa2292857 uint64
	var __obf_447283d1931030d9 int
	if __obf_15edc752c636a4a0 > 1<<__obf_cdb3e6cf09ad181f.__obf_415e8862356d404c || __obf_fed6fd50f0fcc834 == __obf_b8f9f6f8e536b469 {
		__obf_506f449fa2292857 = __obf_15edc752c636a4a0 - 1
		__obf_447283d1931030d9 = __obf_fed6fd50f0fcc834
	} else {
		__obf_506f449fa2292857 = __obf_15edc752c636a4a0*2 - 1
		__obf_447283d1931030d9 = __obf_fed6fd50f0fcc834 - 1
	}
	__obf_0c2beda650c0a346 := new(__obf_adf9c80aee40a9d8)
	__obf_0c2beda650c0a346.Assign(__obf_506f449fa2292857*2 + 1)
	__obf_0c2beda650c0a346.Shift(__obf_447283d1931030d9 - int(__obf_cdb3e6cf09ad181f.__obf_415e8862356d404c) - 1)

	// The upper and lower bounds are possible outputs only if
	// the original mantissa is even, so that IEEE round-to-even
	// would round to the original mantissa and not the neighbors.
	__obf_268e61ca1c919b3a := __obf_15edc752c636a4a0%2 == 0

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
	var __obf_7d66af59ae5531b6 uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_c6baf9ca2480264f := 0; ; __obf_c6baf9ca2480264f++ {
		// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_a1d8263aff745fa0 := __obf_c6baf9ca2480264f - __obf_176a296e341c9131.__obf_d30835ea2dfed4ba + __obf_c4701b3bb28cd2ae.__obf_d30835ea2dfed4ba
		if __obf_a1d8263aff745fa0 >= __obf_c4701b3bb28cd2ae.__obf_087eecd15d19676e {
			break
		}
		__obf_c58859aabfcf97da := __obf_c6baf9ca2480264f - __obf_176a296e341c9131.__obf_d30835ea2dfed4ba + __obf_0c2beda650c0a346.__obf_d30835ea2dfed4ba
		__obf_397799876e0b47ec := byte('0') // lower digit
		if __obf_c58859aabfcf97da >= 0 && __obf_c58859aabfcf97da < __obf_0c2beda650c0a346.__obf_087eecd15d19676e {
			__obf_397799876e0b47ec = __obf_0c2beda650c0a346.__obf_c4701b3bb28cd2ae[__obf_c58859aabfcf97da]
		}
		__obf_4739b75ffb6404bf := byte('0') // middle digit
		if __obf_a1d8263aff745fa0 >= 0 {
			__obf_4739b75ffb6404bf = __obf_c4701b3bb28cd2ae.__obf_c4701b3bb28cd2ae[__obf_a1d8263aff745fa0]
		}
		__obf_4b28f68e230383a4 := byte('0') // upper digit
		if __obf_c6baf9ca2480264f < __obf_176a296e341c9131.__obf_087eecd15d19676e {
			__obf_4b28f68e230383a4 = __obf_176a296e341c9131.__obf_c4701b3bb28cd2ae[__obf_c6baf9ca2480264f]
		}

		// Okay to round down (truncate) if lower has a different digit
		// or if lower is inclusive and is exactly the result of rounding
		// down (i.e., and we have reached the final digit of lower).
		__obf_1cb267adcffd3e79 := __obf_397799876e0b47ec != __obf_4739b75ffb6404bf || __obf_268e61ca1c919b3a && __obf_c58859aabfcf97da+1 == __obf_0c2beda650c0a346.__obf_087eecd15d19676e

		switch {
		case __obf_7d66af59ae5531b6 == 0 && __obf_4739b75ffb6404bf+1 < __obf_4b28f68e230383a4:
			// Example:
			// m = 12345xxx
			// u = 12347xxx
			__obf_7d66af59ae5531b6 = 2
		case __obf_7d66af59ae5531b6 == 0 && __obf_4739b75ffb6404bf != __obf_4b28f68e230383a4:
			// Example:
			// m = 12345xxx
			// u = 12346xxx
			__obf_7d66af59ae5531b6 = 1
		case __obf_7d66af59ae5531b6 == 1 && (__obf_4739b75ffb6404bf != '9' || __obf_4b28f68e230383a4 != '0'):
			// Example:
			// m = 1234598x
			// u = 1234600x
			__obf_7d66af59ae5531b6 = 2
		}
		// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_22d9862dd6427e1d := __obf_7d66af59ae5531b6 > 0 && (__obf_268e61ca1c919b3a || __obf_7d66af59ae5531b6 > 1 || __obf_c6baf9ca2480264f+1 < __obf_176a296e341c9131.__obf_087eecd15d19676e)

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		switch {
		case __obf_1cb267adcffd3e79 && __obf_22d9862dd6427e1d:
			__obf_c4701b3bb28cd2ae.Round(__obf_a1d8263aff745fa0 + 1)
			return
		case __obf_1cb267adcffd3e79:
			__obf_c4701b3bb28cd2ae.RoundDown(__obf_a1d8263aff745fa0 + 1)
			return
		case __obf_22d9862dd6427e1d:
			__obf_c4701b3bb28cd2ae.RoundUp(__obf_a1d8263aff745fa0 + 1)
			return
		}
	}
}
