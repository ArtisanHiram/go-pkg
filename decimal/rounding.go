// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_ae16adf734cfe1aa

type __obf_94eeb44634f64ac2 struct {
	__obf_5991f9c502ad3b8d uint
	__obf_7361f35f98ac5740 uint
	__obf_bd6660fbe7220f81 int
}

var __obf_a47ee567a13f3274 = __obf_94eeb44634f64ac2{23, 8, -127}
var __obf_3201299bc3256294 = __obf_94eeb44634f64ac2{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_a9030787063799d3(__obf_19a369452cd5fbde *__obf_ae16adf734cfe1aa, __obf_6a3e060d23caf99c uint64, __obf_1ef84065a6a49f43 int, __obf_f9f981b2fd2cf5a1 *__obf_94eeb44634f64ac2) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_6a3e060d23caf99c == 0 {
		__obf_19a369452cd5fbde.__obf_04074ea54302ce17 = 0
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
	__obf_d2f7695c876e8976 := __obf_f9f981b2fd2cf5a1.__obf_bd6660fbe7220f81 + 1 // minimum possible exponent
	if __obf_1ef84065a6a49f43 > __obf_d2f7695c876e8976 && 332*(__obf_19a369452cd5fbde.__obf_45565779317643f9-__obf_19a369452cd5fbde.__obf_04074ea54302ce17) >= 100*(__obf_1ef84065a6a49f43-int(__obf_f9f981b2fd2cf5a1.__obf_5991f9c502ad3b8d)) {
		// The number is already shortest.
		return
	}

	// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	__obf_50b9a38cd6526423 := new(__obf_ae16adf734cfe1aa)
	__obf_50b9a38cd6526423.Assign(__obf_6a3e060d23caf99c*2 + 1)
	__obf_50b9a38cd6526423.Shift(__obf_1ef84065a6a49f43 - int(__obf_f9f981b2fd2cf5a1.__obf_5991f9c502ad3b8d) - 1)

	// d = mant << (exp - mantbits)
	// Next lowest floating point number is mant-1 << exp-mantbits,
	// unless mant-1 drops the significant bit and exp is not the minimum exp,
	// in which case the next lowest is mant*2-1 << exp-mantbits-1.
	// Either way, call it mantlo << explo-mantbits.
	// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
	var __obf_d40e13d06ea2b029 uint64
	var __obf_8529e110164c41e2 int
	if __obf_6a3e060d23caf99c > 1<<__obf_f9f981b2fd2cf5a1.__obf_5991f9c502ad3b8d || __obf_1ef84065a6a49f43 == __obf_d2f7695c876e8976 {
		__obf_d40e13d06ea2b029 = __obf_6a3e060d23caf99c - 1
		__obf_8529e110164c41e2 = __obf_1ef84065a6a49f43
	} else {
		__obf_d40e13d06ea2b029 = __obf_6a3e060d23caf99c*2 - 1
		__obf_8529e110164c41e2 = __obf_1ef84065a6a49f43 - 1
	}
	__obf_c221b01ae8c2a888 := new(__obf_ae16adf734cfe1aa)
	__obf_c221b01ae8c2a888.Assign(__obf_d40e13d06ea2b029*2 + 1)
	__obf_c221b01ae8c2a888.Shift(__obf_8529e110164c41e2 - int(__obf_f9f981b2fd2cf5a1.__obf_5991f9c502ad3b8d) - 1)

	// The upper and lower bounds are possible outputs only if
	// the original mantissa is even, so that IEEE round-to-even
	// would round to the original mantissa and not the neighbors.
	__obf_673f359a845e5b4e := __obf_6a3e060d23caf99c%2 == 0

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
	var __obf_10b4bba2401ed6aa uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_ec274a823f20839c := 0; ; __obf_ec274a823f20839c++ {
		// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_975d33c97396864c := __obf_ec274a823f20839c - __obf_50b9a38cd6526423.__obf_45565779317643f9 + __obf_19a369452cd5fbde.__obf_45565779317643f9
		if __obf_975d33c97396864c >= __obf_19a369452cd5fbde.__obf_04074ea54302ce17 {
			break
		}
		__obf_a9dc7e324a60ca96 := __obf_ec274a823f20839c - __obf_50b9a38cd6526423.__obf_45565779317643f9 + __obf_c221b01ae8c2a888.__obf_45565779317643f9
		__obf_a9cdd1166d877eb4 := byte('0') // lower digit
		if __obf_a9dc7e324a60ca96 >= 0 && __obf_a9dc7e324a60ca96 < __obf_c221b01ae8c2a888.__obf_04074ea54302ce17 {
			__obf_a9cdd1166d877eb4 = __obf_c221b01ae8c2a888.__obf_19a369452cd5fbde[__obf_a9dc7e324a60ca96]
		}
		__obf_3f873ba4f9076041 := byte('0') // middle digit
		if __obf_975d33c97396864c >= 0 {
			__obf_3f873ba4f9076041 = __obf_19a369452cd5fbde.__obf_19a369452cd5fbde[__obf_975d33c97396864c]
		}
		__obf_803f47736e5b323f := byte('0') // upper digit
		if __obf_ec274a823f20839c < __obf_50b9a38cd6526423.__obf_04074ea54302ce17 {
			__obf_803f47736e5b323f = __obf_50b9a38cd6526423.__obf_19a369452cd5fbde[__obf_ec274a823f20839c]
		}

		// Okay to round down (truncate) if lower has a different digit
		// or if lower is inclusive and is exactly the result of rounding
		// down (i.e., and we have reached the final digit of lower).
		__obf_f95e2da80ed45e5a := __obf_a9cdd1166d877eb4 != __obf_3f873ba4f9076041 || __obf_673f359a845e5b4e && __obf_a9dc7e324a60ca96+1 == __obf_c221b01ae8c2a888.__obf_04074ea54302ce17

		switch {
		case __obf_10b4bba2401ed6aa == 0 && __obf_3f873ba4f9076041+1 < __obf_803f47736e5b323f:
			// Example:
			// m = 12345xxx
			// u = 12347xxx
			__obf_10b4bba2401ed6aa = 2
		case __obf_10b4bba2401ed6aa == 0 && __obf_3f873ba4f9076041 != __obf_803f47736e5b323f:
			// Example:
			// m = 12345xxx
			// u = 12346xxx
			__obf_10b4bba2401ed6aa = 1
		case __obf_10b4bba2401ed6aa == 1 && (__obf_3f873ba4f9076041 != '9' || __obf_803f47736e5b323f != '0'):
			// Example:
			// m = 1234598x
			// u = 1234600x
			__obf_10b4bba2401ed6aa = 2
		}
		// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_3293f190f2566bd6 := __obf_10b4bba2401ed6aa > 0 && (__obf_673f359a845e5b4e || __obf_10b4bba2401ed6aa > 1 || __obf_ec274a823f20839c+1 < __obf_50b9a38cd6526423.__obf_04074ea54302ce17)

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		switch {
		case __obf_f95e2da80ed45e5a && __obf_3293f190f2566bd6:
			__obf_19a369452cd5fbde.Round(__obf_975d33c97396864c + 1)
			return
		case __obf_f95e2da80ed45e5a:
			__obf_19a369452cd5fbde.RoundDown(__obf_975d33c97396864c + 1)
			return
		case __obf_3293f190f2566bd6:
			__obf_19a369452cd5fbde.RoundUp(__obf_975d33c97396864c + 1)
			return
		}
	}
}
