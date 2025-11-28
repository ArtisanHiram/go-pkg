// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_210b94d2c4b23455

type __obf_73320c94d3fe82cc struct {
	__obf_b197d08c88c50ba7 uint
	__obf_2e8a1cc5b4179be2 uint
	__obf_3dc7e69ee4097b24 int
}

var __obf_75b4ac7d1ab2e9d4 = __obf_73320c94d3fe82cc{23, 8, -127}
var __obf_7e93547175be045c = __obf_73320c94d3fe82cc{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_705af1f438fcfe9e(__obf_45d3205e59c45c8b *__obf_210b94d2c4b23455, __obf_5f6f09d40b1d194a uint64, __obf_e060caa13549e714 int, __obf_c7a132b4ff25bccf *__obf_73320c94d3fe82cc) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_5f6f09d40b1d194a == 0 {
		__obf_45d3205e59c45c8b.__obf_2298ab2cdf3cb6bd = 0
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
	__obf_c63676764dbcbdc6 := __obf_c7a132b4ff25bccf.__obf_3dc7e69ee4097b24 + 1 // minimum possible exponent
	if __obf_e060caa13549e714 > __obf_c63676764dbcbdc6 && 332*(__obf_45d3205e59c45c8b.__obf_44d25d407353f76c-__obf_45d3205e59c45c8b.__obf_2298ab2cdf3cb6bd) >= 100*(__obf_e060caa13549e714-int(__obf_c7a132b4ff25bccf.__obf_b197d08c88c50ba7)) {
		// The number is already shortest.
		return
	}

	// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	__obf_5d3f01212578fea8 := new(__obf_210b94d2c4b23455)
	__obf_5d3f01212578fea8.Assign(__obf_5f6f09d40b1d194a*2 + 1)
	__obf_5d3f01212578fea8.Shift(__obf_e060caa13549e714 - int(__obf_c7a132b4ff25bccf.__obf_b197d08c88c50ba7) - 1)

	// d = mant << (exp - mantbits)
	// Next lowest floating point number is mant-1 << exp-mantbits,
	// unless mant-1 drops the significant bit and exp is not the minimum exp,
	// in which case the next lowest is mant*2-1 << exp-mantbits-1.
	// Either way, call it mantlo << explo-mantbits.
	// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
	var __obf_1ae5e35f3ea4b3a2 uint64
	var __obf_7cd4d6f8421e7aa9 int
	if __obf_5f6f09d40b1d194a > 1<<__obf_c7a132b4ff25bccf.__obf_b197d08c88c50ba7 || __obf_e060caa13549e714 == __obf_c63676764dbcbdc6 {
		__obf_1ae5e35f3ea4b3a2 = __obf_5f6f09d40b1d194a - 1
		__obf_7cd4d6f8421e7aa9 = __obf_e060caa13549e714
	} else {
		__obf_1ae5e35f3ea4b3a2 = __obf_5f6f09d40b1d194a*2 - 1
		__obf_7cd4d6f8421e7aa9 = __obf_e060caa13549e714 - 1
	}
	__obf_df8e209aa96b6bd8 := new(__obf_210b94d2c4b23455)
	__obf_df8e209aa96b6bd8.Assign(__obf_1ae5e35f3ea4b3a2*2 + 1)
	__obf_df8e209aa96b6bd8.Shift(__obf_7cd4d6f8421e7aa9 - int(__obf_c7a132b4ff25bccf.__obf_b197d08c88c50ba7) - 1)

	// The upper and lower bounds are possible outputs only if
	// the original mantissa is even, so that IEEE round-to-even
	// would round to the original mantissa and not the neighbors.
	__obf_81b823809b73cc4d := __obf_5f6f09d40b1d194a%2 == 0

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
	var __obf_46523e1e2b4f27d7 uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_22a50dcf65de6546 := 0; ; __obf_22a50dcf65de6546++ {
		// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_091d5e6528d1126c := __obf_22a50dcf65de6546 - __obf_5d3f01212578fea8.__obf_44d25d407353f76c + __obf_45d3205e59c45c8b.__obf_44d25d407353f76c
		if __obf_091d5e6528d1126c >= __obf_45d3205e59c45c8b.__obf_2298ab2cdf3cb6bd {
			break
		}
		__obf_28380bd6dee04d6d := __obf_22a50dcf65de6546 - __obf_5d3f01212578fea8.__obf_44d25d407353f76c + __obf_df8e209aa96b6bd8.__obf_44d25d407353f76c
		__obf_b2371e9cc964aa35 := byte('0') // lower digit
		if __obf_28380bd6dee04d6d >= 0 && __obf_28380bd6dee04d6d < __obf_df8e209aa96b6bd8.__obf_2298ab2cdf3cb6bd {
			__obf_b2371e9cc964aa35 = __obf_df8e209aa96b6bd8.__obf_45d3205e59c45c8b[__obf_28380bd6dee04d6d]
		}
		__obf_763dd0cca8ce709c := byte('0') // middle digit
		if __obf_091d5e6528d1126c >= 0 {
			__obf_763dd0cca8ce709c = __obf_45d3205e59c45c8b.__obf_45d3205e59c45c8b[__obf_091d5e6528d1126c]
		}
		__obf_34f291ba009def63 := byte('0') // upper digit
		if __obf_22a50dcf65de6546 < __obf_5d3f01212578fea8.__obf_2298ab2cdf3cb6bd {
			__obf_34f291ba009def63 = __obf_5d3f01212578fea8.__obf_45d3205e59c45c8b[__obf_22a50dcf65de6546]
		}

		// Okay to round down (truncate) if lower has a different digit
		// or if lower is inclusive and is exactly the result of rounding
		// down (i.e., and we have reached the final digit of lower).
		__obf_cdb3317f49b74dc5 := __obf_b2371e9cc964aa35 != __obf_763dd0cca8ce709c || __obf_81b823809b73cc4d && __obf_28380bd6dee04d6d+1 == __obf_df8e209aa96b6bd8.__obf_2298ab2cdf3cb6bd

		switch {
		case __obf_46523e1e2b4f27d7 == 0 && __obf_763dd0cca8ce709c+1 < __obf_34f291ba009def63:
			// Example:
			// m = 12345xxx
			// u = 12347xxx
			__obf_46523e1e2b4f27d7 = 2
		case __obf_46523e1e2b4f27d7 == 0 && __obf_763dd0cca8ce709c != __obf_34f291ba009def63:
			// Example:
			// m = 12345xxx
			// u = 12346xxx
			__obf_46523e1e2b4f27d7 = 1
		case __obf_46523e1e2b4f27d7 == 1 && (__obf_763dd0cca8ce709c != '9' || __obf_34f291ba009def63 != '0'):
			// Example:
			// m = 1234598x
			// u = 1234600x
			__obf_46523e1e2b4f27d7 = 2
		}
		// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_7ea5a3386d73ad9d := __obf_46523e1e2b4f27d7 > 0 && (__obf_81b823809b73cc4d || __obf_46523e1e2b4f27d7 > 1 || __obf_22a50dcf65de6546+1 < __obf_5d3f01212578fea8.__obf_2298ab2cdf3cb6bd)

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		switch {
		case __obf_cdb3317f49b74dc5 && __obf_7ea5a3386d73ad9d:
			__obf_45d3205e59c45c8b.Round(__obf_091d5e6528d1126c + 1)
			return
		case __obf_cdb3317f49b74dc5:
			__obf_45d3205e59c45c8b.RoundDown(__obf_091d5e6528d1126c + 1)
			return
		case __obf_7ea5a3386d73ad9d:
			__obf_45d3205e59c45c8b.RoundUp(__obf_091d5e6528d1126c + 1)
			return
		}
	}
}
