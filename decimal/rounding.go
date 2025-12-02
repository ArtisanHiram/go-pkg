// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_bd46136174bdabe4

type __obf_3510ef834411b273 struct {
	__obf_350a4b63232b6662 uint
	__obf_33dfabedaba76fac uint
	__obf_d3ccd221dc097eda int
}

var __obf_8177a6f362debfa1 = __obf_3510ef834411b273{23, 8, -127}
var __obf_5c7b21ace088503b = __obf_3510ef834411b273{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_a3ec81b881a7faf1(__obf_a1ffba48568e5107 *__obf_bd46136174bdabe4, __obf_a1a9265f031cbbf2 uint64, __obf_3ab33991ba7f9863 int, __obf_db4d85e8d4b69c3f *__obf_3510ef834411b273,) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_a1a9265f031cbbf2 == 0 {
		__obf_a1ffba48568e5107.__obf_3018d9650ee6ca39 = 0
		return
	}
	__obf_6ce7f2752ac44142 :=// Compute upper and lower such that any decimal number
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
	__obf_db4d85e8d4b69c3f. // minimum possible exponent
	__obf_d3ccd221dc097eda + 1
	if __obf_3ab33991ba7f9863 > __obf_6ce7f2752ac44142 && 332*(__obf_a1ffba48568e5107.__obf_c3fb49899f9ddeff-__obf_a1ffba48568e5107.__obf_3018d9650ee6ca39) >= 100*(__obf_3ab33991ba7f9863-int(__obf_db4d85e8d4b69c3f.
	// The number is already shortest.
	__obf_350a4b63232b6662)) {

		return
	}
	__obf_a675830d4e9ee80d :=// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	new(__obf_bd46136174bdabe4)
	__obf_a675830d4e9ee80d.
		Assign(__obf_a1a9265f031cbbf2*2 + 1)
	__obf_a675830d4e9ee80d.
		Shift(__obf_3ab33991ba7f9863 - int(__obf_db4d85e8d4b69c3f.

		// d = mant << (exp - mantbits)
		// Next lowest floating point number is mant-1 << exp-mantbits,
		// unless mant-1 drops the significant bit and exp is not the minimum exp,
		// in which case the next lowest is mant*2-1 << exp-mantbits-1.
		// Either way, call it mantlo << explo-mantbits.
		// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
		__obf_350a4b63232b6662) - 1)

	var __obf_e8e7fe73ab5cba0f uint64
	var __obf_2121839fe402984e int
	if __obf_a1a9265f031cbbf2 > 1<<__obf_db4d85e8d4b69c3f.__obf_350a4b63232b6662 || __obf_3ab33991ba7f9863 == __obf_6ce7f2752ac44142 {
		__obf_e8e7fe73ab5cba0f = __obf_a1a9265f031cbbf2 - 1
		__obf_2121839fe402984e = __obf_3ab33991ba7f9863
	} else {
		__obf_e8e7fe73ab5cba0f = __obf_a1a9265f031cbbf2*2 - 1
		__obf_2121839fe402984e = __obf_3ab33991ba7f9863 - 1
	}
	__obf_6a1452d0110d0a19 := new(__obf_bd46136174bdabe4)
	__obf_6a1452d0110d0a19.
		Assign(__obf_e8e7fe73ab5cba0f*2 + 1)
	__obf_6a1452d0110d0a19.
		Shift(__obf_2121839fe402984e - int(__obf_db4d85e8d4b69c3f.

		// The upper and lower bounds are possible outputs only if
		// the original mantissa is even, so that IEEE round-to-even
		// would round to the original mantissa and not the neighbors.
		__obf_350a4b63232b6662) - 1)
	__obf_771501e25f11487e := __obf_a1a9265f031cbbf2%2 == 0

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
	var __obf_ad99cf5408bd992c uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_88717fe81374ccf8 := 0; ; __obf_88717fe81374ccf8++ {
		__obf_120a15cd1c761900 :=// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_88717fe81374ccf8 - __obf_a675830d4e9ee80d.__obf_c3fb49899f9ddeff + __obf_a1ffba48568e5107.__obf_c3fb49899f9ddeff
		if __obf_120a15cd1c761900 >= __obf_a1ffba48568e5107.__obf_3018d9650ee6ca39 {
			break
		}
		__obf_ea735b87f8fbbac8 := __obf_88717fe81374ccf8 - __obf_a675830d4e9ee80d.__obf_c3fb49899f9ddeff + __obf_6a1452d0110d0a19.__obf_c3fb49899f9ddeff
		// lower digit
		__obf_395939c10a12f68e := byte('0')
		if __obf_ea735b87f8fbbac8 >= 0 && __obf_ea735b87f8fbbac8 < __obf_6a1452d0110d0a19.__obf_3018d9650ee6ca39 {
			__obf_395939c10a12f68e = __obf_6a1452d0110d0a19.__obf_a1ffba48568e5107[__obf_ea735b87f8fbbac8]
		}
		__obf_a49032a74f3c7436 := byte('0') // middle digit
		if __obf_120a15cd1c761900 >= 0 {
			__obf_a49032a74f3c7436 = __obf_a1ffba48568e5107.__obf_a1ffba48568e5107[__obf_120a15cd1c761900]
		}
		__obf_9dbb6f343bc63a35 := byte('0') // upper digit
		if __obf_88717fe81374ccf8 < __obf_a675830d4e9ee80d.__obf_3018d9650ee6ca39 {
			__obf_9dbb6f343bc63a35 = __obf_a675830d4e9ee80d.

			// Okay to round down (truncate) if lower has a different digit
			// or if lower is inclusive and is exactly the result of rounding
			// down (i.e., and we have reached the final digit of lower).
			__obf_a1ffba48568e5107[__obf_88717fe81374ccf8]
		}
		__obf_623704defa0f9b24 := __obf_395939c10a12f68e != __obf_a49032a74f3c7436 || __obf_771501e25f11487e && __obf_ea735b87f8fbbac8+1 == __obf_6a1452d0110d0a19.__obf_3018d9650ee6ca39

		switch {
		case __obf_ad99cf5408bd992c == 0 && __obf_a49032a74f3c7436+1 < __obf_9dbb6f343bc63a35:
			__obf_ad99cf5408bd992c =// Example:
			// m = 12345xxx
			// u = 12347xxx
			2
		case __obf_ad99cf5408bd992c == 0 && __obf_a49032a74f3c7436 != __obf_9dbb6f343bc63a35:
			__obf_ad99cf5408bd992c =// Example:
			// m = 12345xxx
			// u = 12346xxx
			1
		case __obf_ad99cf5408bd992c == 1 && (__obf_a49032a74f3c7436 != '9' || __obf_9dbb6f343bc63a35 != '0'):
			__obf_ad99cf5408bd992c =// Example:
			// m = 1234598x
			// u = 1234600x
			2
		}
		__obf_eb25397190129ede :=// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_ad99cf5408bd992c > 0 && (__obf_771501e25f11487e || __obf_ad99cf5408bd992c > 1 || __obf_88717fe81374ccf8+1 < __obf_a675830d4e9ee80d.

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		__obf_3018d9650ee6ca39)

		switch {
		case __obf_623704defa0f9b24 && __obf_eb25397190129ede:
			__obf_a1ffba48568e5107.
				Round(__obf_120a15cd1c761900 + 1)
			return
		case __obf_623704defa0f9b24:
			__obf_a1ffba48568e5107.
				RoundDown(__obf_120a15cd1c761900 + 1)
			return
		case __obf_eb25397190129ede:
			__obf_a1ffba48568e5107.
				RoundUp(__obf_120a15cd1c761900 + 1)
			return
		}
	}
}
