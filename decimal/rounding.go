// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_3e0a215d3ac5298d

type __obf_b9da0d5925fb006e struct {
	__obf_179bd2684cdde9cd uint
	__obf_1eecfa2ca0d4455f uint
	__obf_b1391270660d97a8 int
}

var __obf_9381b591902a5ca1 = __obf_b9da0d5925fb006e{23, 8, -127}
var __obf_dc4c42ad9bdcae6e = __obf_b9da0d5925fb006e{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_4e452f472f313f6b(__obf_e498fd1d3feac2c4 *__obf_3e0a215d3ac5298d, __obf_a8350a62e8fe3cad uint64, __obf_4b33f2812f2bcc7e int, __obf_414bae728c6159e7 *__obf_b9da0d5925fb006e,) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_a8350a62e8fe3cad == 0 {
		__obf_e498fd1d3feac2c4.__obf_7a088d2e01e4cd05 = 0
		return
	}
	__obf_5f93232b668d88e3 :=// Compute upper and lower such that any decimal number
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
	__obf_414bae728c6159e7. // minimum possible exponent
	__obf_b1391270660d97a8 + 1
	if __obf_4b33f2812f2bcc7e > __obf_5f93232b668d88e3 && 332*(__obf_e498fd1d3feac2c4.__obf_164d07da7c832f2c-__obf_e498fd1d3feac2c4.__obf_7a088d2e01e4cd05) >= 100*(__obf_4b33f2812f2bcc7e-int(__obf_414bae728c6159e7.
	// The number is already shortest.
	__obf_179bd2684cdde9cd)) {

		return
	}
	__obf_88a335dba8c72a3b :=// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	new(__obf_3e0a215d3ac5298d)
	__obf_88a335dba8c72a3b.
		Assign(__obf_a8350a62e8fe3cad*2 + 1)
	__obf_88a335dba8c72a3b.
		Shift(__obf_4b33f2812f2bcc7e - int(__obf_414bae728c6159e7.

		// d = mant << (exp - mantbits)
		// Next lowest floating point number is mant-1 << exp-mantbits,
		// unless mant-1 drops the significant bit and exp is not the minimum exp,
		// in which case the next lowest is mant*2-1 << exp-mantbits-1.
		// Either way, call it mantlo << explo-mantbits.
		// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
		__obf_179bd2684cdde9cd) - 1)

	var __obf_d2ada3d19efc81ac uint64
	var __obf_81245a83477ecf38 int
	if __obf_a8350a62e8fe3cad > 1<<__obf_414bae728c6159e7.__obf_179bd2684cdde9cd || __obf_4b33f2812f2bcc7e == __obf_5f93232b668d88e3 {
		__obf_d2ada3d19efc81ac = __obf_a8350a62e8fe3cad - 1
		__obf_81245a83477ecf38 = __obf_4b33f2812f2bcc7e
	} else {
		__obf_d2ada3d19efc81ac = __obf_a8350a62e8fe3cad*2 - 1
		__obf_81245a83477ecf38 = __obf_4b33f2812f2bcc7e - 1
	}
	__obf_5a1148a0ba21d5a8 := new(__obf_3e0a215d3ac5298d)
	__obf_5a1148a0ba21d5a8.
		Assign(__obf_d2ada3d19efc81ac*2 + 1)
	__obf_5a1148a0ba21d5a8.
		Shift(__obf_81245a83477ecf38 - int(__obf_414bae728c6159e7.

		// The upper and lower bounds are possible outputs only if
		// the original mantissa is even, so that IEEE round-to-even
		// would round to the original mantissa and not the neighbors.
		__obf_179bd2684cdde9cd) - 1)
	__obf_04e21ea4ee469d56 := __obf_a8350a62e8fe3cad%2 == 0

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
	var __obf_812d0bfa6f2e2d24 uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_a5a24b56813aa7d6 := 0; ; __obf_a5a24b56813aa7d6++ {
		__obf_d45c5cf2fd7ba95c :=// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_a5a24b56813aa7d6 - __obf_88a335dba8c72a3b.__obf_164d07da7c832f2c + __obf_e498fd1d3feac2c4.__obf_164d07da7c832f2c
		if __obf_d45c5cf2fd7ba95c >= __obf_e498fd1d3feac2c4.__obf_7a088d2e01e4cd05 {
			break
		}
		__obf_49b77fb32eef3139 := __obf_a5a24b56813aa7d6 - __obf_88a335dba8c72a3b.__obf_164d07da7c832f2c + __obf_5a1148a0ba21d5a8.__obf_164d07da7c832f2c
		// lower digit
		__obf_2555792e3e804add := byte('0')
		if __obf_49b77fb32eef3139 >= 0 && __obf_49b77fb32eef3139 < __obf_5a1148a0ba21d5a8.__obf_7a088d2e01e4cd05 {
			__obf_2555792e3e804add = __obf_5a1148a0ba21d5a8.__obf_e498fd1d3feac2c4[__obf_49b77fb32eef3139]
		}
		__obf_e524d5ae197a0bbb := byte('0') // middle digit
		if __obf_d45c5cf2fd7ba95c >= 0 {
			__obf_e524d5ae197a0bbb = __obf_e498fd1d3feac2c4.__obf_e498fd1d3feac2c4[__obf_d45c5cf2fd7ba95c]
		}
		__obf_b9289f6e07c1c649 := byte('0') // upper digit
		if __obf_a5a24b56813aa7d6 < __obf_88a335dba8c72a3b.__obf_7a088d2e01e4cd05 {
			__obf_b9289f6e07c1c649 = __obf_88a335dba8c72a3b.

			// Okay to round down (truncate) if lower has a different digit
			// or if lower is inclusive and is exactly the result of rounding
			// down (i.e., and we have reached the final digit of lower).
			__obf_e498fd1d3feac2c4[__obf_a5a24b56813aa7d6]
		}
		__obf_d5e2dd50a172b254 := __obf_2555792e3e804add != __obf_e524d5ae197a0bbb || __obf_04e21ea4ee469d56 && __obf_49b77fb32eef3139+1 == __obf_5a1148a0ba21d5a8.__obf_7a088d2e01e4cd05

		switch {
		case __obf_812d0bfa6f2e2d24 == 0 && __obf_e524d5ae197a0bbb+1 < __obf_b9289f6e07c1c649:
			__obf_812d0bfa6f2e2d24 =// Example:
			// m = 12345xxx
			// u = 12347xxx
			2
		case __obf_812d0bfa6f2e2d24 == 0 && __obf_e524d5ae197a0bbb != __obf_b9289f6e07c1c649:
			__obf_812d0bfa6f2e2d24 =// Example:
			// m = 12345xxx
			// u = 12346xxx
			1
		case __obf_812d0bfa6f2e2d24 == 1 && (__obf_e524d5ae197a0bbb != '9' || __obf_b9289f6e07c1c649 != '0'):
			__obf_812d0bfa6f2e2d24 =// Example:
			// m = 1234598x
			// u = 1234600x
			2
		}
		__obf_1083bb983af0dd7b :=// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_812d0bfa6f2e2d24 > 0 && (__obf_04e21ea4ee469d56 || __obf_812d0bfa6f2e2d24 > 1 || __obf_a5a24b56813aa7d6+1 < __obf_88a335dba8c72a3b.

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		__obf_7a088d2e01e4cd05)

		switch {
		case __obf_d5e2dd50a172b254 && __obf_1083bb983af0dd7b:
			__obf_e498fd1d3feac2c4.
				Round(__obf_d45c5cf2fd7ba95c + 1)
			return
		case __obf_d5e2dd50a172b254:
			__obf_e498fd1d3feac2c4.
				RoundDown(__obf_d45c5cf2fd7ba95c + 1)
			return
		case __obf_1083bb983af0dd7b:
			__obf_e498fd1d3feac2c4.
				RoundUp(__obf_d45c5cf2fd7ba95c + 1)
			return
		}
	}
}
