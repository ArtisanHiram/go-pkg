// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_b0e21776b9d45e13

type __obf_7361d3b42dd22c09 struct {
	__obf_2f2d498924addd14 uint
	__obf_96af297d4e88f620 uint
	__obf_0628d944dfc724b5 int
}

var __obf_2ea0d484844c6331 = __obf_7361d3b42dd22c09{23, 8, -127}
var __obf_0e52bd84f2e0c5e5 = __obf_7361d3b42dd22c09{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_96ba681e9a854a03(__obf_df803c2772cc0386 *__obf_b0e21776b9d45e13, __obf_22ce9de7558a3141 uint64, __obf_97e810ed39f8ac3f int, __obf_c19e228d4278835d *__obf_7361d3b42dd22c09,) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_22ce9de7558a3141 == 0 {
		__obf_df803c2772cc0386.__obf_2841543582033a23 = 0
		return
	}
	__obf_edd5970b7263bf62 :=// Compute upper and lower such that any decimal number
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
	__obf_c19e228d4278835d. // minimum possible exponent
	__obf_0628d944dfc724b5 + 1
	if __obf_97e810ed39f8ac3f > __obf_edd5970b7263bf62 && 332*(__obf_df803c2772cc0386.__obf_b4c2d4c5eeb9a307-__obf_df803c2772cc0386.__obf_2841543582033a23) >= 100*(__obf_97e810ed39f8ac3f-int(__obf_c19e228d4278835d.
	// The number is already shortest.
	__obf_2f2d498924addd14)) {

		return
	}
	__obf_df5a436834d581a3 :=// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	new(__obf_b0e21776b9d45e13)
	__obf_df5a436834d581a3.
		Assign(__obf_22ce9de7558a3141*2 + 1)
	__obf_df5a436834d581a3.
		Shift(__obf_97e810ed39f8ac3f - int(__obf_c19e228d4278835d.

		// d = mant << (exp - mantbits)
		// Next lowest floating point number is mant-1 << exp-mantbits,
		// unless mant-1 drops the significant bit and exp is not the minimum exp,
		// in which case the next lowest is mant*2-1 << exp-mantbits-1.
		// Either way, call it mantlo << explo-mantbits.
		// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
		__obf_2f2d498924addd14) - 1)

	var __obf_83a13514fbb817f8 uint64
	var __obf_b3d2d58036f46115 int
	if __obf_22ce9de7558a3141 > 1<<__obf_c19e228d4278835d.__obf_2f2d498924addd14 || __obf_97e810ed39f8ac3f == __obf_edd5970b7263bf62 {
		__obf_83a13514fbb817f8 = __obf_22ce9de7558a3141 - 1
		__obf_b3d2d58036f46115 = __obf_97e810ed39f8ac3f
	} else {
		__obf_83a13514fbb817f8 = __obf_22ce9de7558a3141*2 - 1
		__obf_b3d2d58036f46115 = __obf_97e810ed39f8ac3f - 1
	}
	__obf_40f9b405cbd46946 := new(__obf_b0e21776b9d45e13)
	__obf_40f9b405cbd46946.
		Assign(__obf_83a13514fbb817f8*2 + 1)
	__obf_40f9b405cbd46946.
		Shift(__obf_b3d2d58036f46115 - int(__obf_c19e228d4278835d.

		// The upper and lower bounds are possible outputs only if
		// the original mantissa is even, so that IEEE round-to-even
		// would round to the original mantissa and not the neighbors.
		__obf_2f2d498924addd14) - 1)
	__obf_8cd47ad8fbacff78 := __obf_22ce9de7558a3141%2 == 0

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
	var __obf_e8f899030e8c9120 uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_4629c791017603c0 := 0; ; __obf_4629c791017603c0++ {
		__obf_451e6a8232afef31 :=// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_4629c791017603c0 - __obf_df5a436834d581a3.__obf_b4c2d4c5eeb9a307 + __obf_df803c2772cc0386.__obf_b4c2d4c5eeb9a307
		if __obf_451e6a8232afef31 >= __obf_df803c2772cc0386.__obf_2841543582033a23 {
			break
		}
		__obf_04b95390bc141ba8 := __obf_4629c791017603c0 - __obf_df5a436834d581a3.__obf_b4c2d4c5eeb9a307 + __obf_40f9b405cbd46946.__obf_b4c2d4c5eeb9a307
		// lower digit
		__obf_9f7df88c606a79dc := byte('0')
		if __obf_04b95390bc141ba8 >= 0 && __obf_04b95390bc141ba8 < __obf_40f9b405cbd46946.__obf_2841543582033a23 {
			__obf_9f7df88c606a79dc = __obf_40f9b405cbd46946.__obf_df803c2772cc0386[__obf_04b95390bc141ba8]
		}
		__obf_cb9fa5e90307ec38 := byte('0') // middle digit
		if __obf_451e6a8232afef31 >= 0 {
			__obf_cb9fa5e90307ec38 = __obf_df803c2772cc0386.__obf_df803c2772cc0386[__obf_451e6a8232afef31]
		}
		__obf_7185e6ffce2bb589 := byte('0') // upper digit
		if __obf_4629c791017603c0 < __obf_df5a436834d581a3.__obf_2841543582033a23 {
			__obf_7185e6ffce2bb589 = __obf_df5a436834d581a3.

			// Okay to round down (truncate) if lower has a different digit
			// or if lower is inclusive and is exactly the result of rounding
			// down (i.e., and we have reached the final digit of lower).
			__obf_df803c2772cc0386[__obf_4629c791017603c0]
		}
		__obf_16b564be3fb20bed := __obf_9f7df88c606a79dc != __obf_cb9fa5e90307ec38 || __obf_8cd47ad8fbacff78 && __obf_04b95390bc141ba8+1 == __obf_40f9b405cbd46946.__obf_2841543582033a23

		switch {
		case __obf_e8f899030e8c9120 == 0 && __obf_cb9fa5e90307ec38+1 < __obf_7185e6ffce2bb589:
			__obf_e8f899030e8c9120 =// Example:
			// m = 12345xxx
			// u = 12347xxx
			2
		case __obf_e8f899030e8c9120 == 0 && __obf_cb9fa5e90307ec38 != __obf_7185e6ffce2bb589:
			__obf_e8f899030e8c9120 =// Example:
			// m = 12345xxx
			// u = 12346xxx
			1
		case __obf_e8f899030e8c9120 == 1 && (__obf_cb9fa5e90307ec38 != '9' || __obf_7185e6ffce2bb589 != '0'):
			__obf_e8f899030e8c9120 =// Example:
			// m = 1234598x
			// u = 1234600x
			2
		}
		__obf_300a37fd5252a419 :=// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_e8f899030e8c9120 > 0 && (__obf_8cd47ad8fbacff78 || __obf_e8f899030e8c9120 > 1 || __obf_4629c791017603c0+1 < __obf_df5a436834d581a3.

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		__obf_2841543582033a23)

		switch {
		case __obf_16b564be3fb20bed && __obf_300a37fd5252a419:
			__obf_df803c2772cc0386.
				Round(__obf_451e6a8232afef31 + 1)
			return
		case __obf_16b564be3fb20bed:
			__obf_df803c2772cc0386.
				RoundDown(__obf_451e6a8232afef31 + 1)
			return
		case __obf_300a37fd5252a419:
			__obf_df803c2772cc0386.
				RoundUp(__obf_451e6a8232afef31 + 1)
			return
		}
	}
}
