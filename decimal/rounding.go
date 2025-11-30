// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_ebbf7bb33b20bf6e

type __obf_8d7a820fe70b0941 struct {
	__obf_2824f9b56a37a620 uint
	__obf_a2f627ebab0e93ee uint
	__obf_69950ce091b60b98 int
}

var __obf_da91496795acf581 = __obf_8d7a820fe70b0941{23, 8, -127}
var __obf_3050fd41f9ad69c5 = __obf_8d7a820fe70b0941{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_ee1a55061233bcaf(__obf_3c5564ffc2911a1b *__obf_ebbf7bb33b20bf6e, __obf_a0aed3f02406911a uint64, __obf_4163ad30ac75204b int, __obf_5c718fb2fc061777 *__obf_8d7a820fe70b0941,) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_a0aed3f02406911a == 0 {
		__obf_3c5564ffc2911a1b.__obf_f75b34971d3492db = 0
		return
	}
	__obf_b1184b45635133ce :=// Compute upper and lower such that any decimal number
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
	__obf_5c718fb2fc061777. // minimum possible exponent
	__obf_69950ce091b60b98 + 1
	if __obf_4163ad30ac75204b > __obf_b1184b45635133ce && 332*(__obf_3c5564ffc2911a1b.__obf_43d811ed87289907-__obf_3c5564ffc2911a1b.__obf_f75b34971d3492db) >= 100*(__obf_4163ad30ac75204b-int(__obf_5c718fb2fc061777.
	// The number is already shortest.
	__obf_2824f9b56a37a620)) {

		return
	}
	__obf_d0b5ebdb95488430 :=// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	new(__obf_ebbf7bb33b20bf6e)
	__obf_d0b5ebdb95488430.
		Assign(__obf_a0aed3f02406911a*2 + 1)
	__obf_d0b5ebdb95488430.
		Shift(__obf_4163ad30ac75204b - int(__obf_5c718fb2fc061777.

		// d = mant << (exp - mantbits)
		// Next lowest floating point number is mant-1 << exp-mantbits,
		// unless mant-1 drops the significant bit and exp is not the minimum exp,
		// in which case the next lowest is mant*2-1 << exp-mantbits-1.
		// Either way, call it mantlo << explo-mantbits.
		// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
		__obf_2824f9b56a37a620) - 1)

	var __obf_fdf0d22bff700901 uint64
	var __obf_13944ec7d8ec486f int
	if __obf_a0aed3f02406911a > 1<<__obf_5c718fb2fc061777.__obf_2824f9b56a37a620 || __obf_4163ad30ac75204b == __obf_b1184b45635133ce {
		__obf_fdf0d22bff700901 = __obf_a0aed3f02406911a - 1
		__obf_13944ec7d8ec486f = __obf_4163ad30ac75204b
	} else {
		__obf_fdf0d22bff700901 = __obf_a0aed3f02406911a*2 - 1
		__obf_13944ec7d8ec486f = __obf_4163ad30ac75204b - 1
	}
	__obf_05c02d67747e10ae := new(__obf_ebbf7bb33b20bf6e)
	__obf_05c02d67747e10ae.
		Assign(__obf_fdf0d22bff700901*2 + 1)
	__obf_05c02d67747e10ae.
		Shift(__obf_13944ec7d8ec486f - int(__obf_5c718fb2fc061777.

		// The upper and lower bounds are possible outputs only if
		// the original mantissa is even, so that IEEE round-to-even
		// would round to the original mantissa and not the neighbors.
		__obf_2824f9b56a37a620) - 1)
	__obf_0e831bbc1106da46 := __obf_a0aed3f02406911a%2 == 0

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
	var __obf_c6480ac18087bfb0 uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_b6d42d879206c0b3 := 0; ; __obf_b6d42d879206c0b3++ {
		__obf_49d18c3040352697 :=// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_b6d42d879206c0b3 - __obf_d0b5ebdb95488430.__obf_43d811ed87289907 + __obf_3c5564ffc2911a1b.__obf_43d811ed87289907
		if __obf_49d18c3040352697 >= __obf_3c5564ffc2911a1b.__obf_f75b34971d3492db {
			break
		}
		__obf_35108c905a4b8649 := __obf_b6d42d879206c0b3 - __obf_d0b5ebdb95488430.__obf_43d811ed87289907 + __obf_05c02d67747e10ae.__obf_43d811ed87289907
		// lower digit
		__obf_70de2970ed80b361 := byte('0')
		if __obf_35108c905a4b8649 >= 0 && __obf_35108c905a4b8649 < __obf_05c02d67747e10ae.__obf_f75b34971d3492db {
			__obf_70de2970ed80b361 = __obf_05c02d67747e10ae.__obf_3c5564ffc2911a1b[__obf_35108c905a4b8649]
		}
		__obf_2c8d42b08bf1a2b7 := byte('0') // middle digit
		if __obf_49d18c3040352697 >= 0 {
			__obf_2c8d42b08bf1a2b7 = __obf_3c5564ffc2911a1b.__obf_3c5564ffc2911a1b[__obf_49d18c3040352697]
		}
		__obf_a2161c9cf8786527 := byte('0') // upper digit
		if __obf_b6d42d879206c0b3 < __obf_d0b5ebdb95488430.__obf_f75b34971d3492db {
			__obf_a2161c9cf8786527 = __obf_d0b5ebdb95488430.

			// Okay to round down (truncate) if lower has a different digit
			// or if lower is inclusive and is exactly the result of rounding
			// down (i.e., and we have reached the final digit of lower).
			__obf_3c5564ffc2911a1b[__obf_b6d42d879206c0b3]
		}
		__obf_bedc8ba440fcefe1 := __obf_70de2970ed80b361 != __obf_2c8d42b08bf1a2b7 || __obf_0e831bbc1106da46 && __obf_35108c905a4b8649+1 == __obf_05c02d67747e10ae.__obf_f75b34971d3492db

		switch {
		case __obf_c6480ac18087bfb0 == 0 && __obf_2c8d42b08bf1a2b7+1 < __obf_a2161c9cf8786527:
			__obf_c6480ac18087bfb0 =// Example:
			// m = 12345xxx
			// u = 12347xxx
			2
		case __obf_c6480ac18087bfb0 == 0 && __obf_2c8d42b08bf1a2b7 != __obf_a2161c9cf8786527:
			__obf_c6480ac18087bfb0 =// Example:
			// m = 12345xxx
			// u = 12346xxx
			1
		case __obf_c6480ac18087bfb0 == 1 && (__obf_2c8d42b08bf1a2b7 != '9' || __obf_a2161c9cf8786527 != '0'):
			__obf_c6480ac18087bfb0 =// Example:
			// m = 1234598x
			// u = 1234600x
			2
		}
		__obf_1f15ca8403f5bd63 :=// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_c6480ac18087bfb0 > 0 && (__obf_0e831bbc1106da46 || __obf_c6480ac18087bfb0 > 1 || __obf_b6d42d879206c0b3+1 < __obf_d0b5ebdb95488430.

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		__obf_f75b34971d3492db)

		switch {
		case __obf_bedc8ba440fcefe1 && __obf_1f15ca8403f5bd63:
			__obf_3c5564ffc2911a1b.
				Round(__obf_49d18c3040352697 + 1)
			return
		case __obf_bedc8ba440fcefe1:
			__obf_3c5564ffc2911a1b.
				RoundDown(__obf_49d18c3040352697 + 1)
			return
		case __obf_1f15ca8403f5bd63:
			__obf_3c5564ffc2911a1b.
				RoundUp(__obf_49d18c3040352697 + 1)
			return
		}
	}
}
