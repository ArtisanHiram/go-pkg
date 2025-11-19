// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

package __obf_17d0cfecf7e687b6

type __obf_ce035e3885d5f561 struct {
	__obf_714ef5fd93ab5467 uint
	__obf_54a54ea1a927766e uint
	__obf_2d0469f7185b0ad4 int
}

var __obf_b00f38b47daaa1ee = __obf_ce035e3885d5f561{23, 8, -127}
var __obf_9dee2c5a44d4ec98 = __obf_ce035e3885d5f561{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func __obf_8163cf85fce10314(__obf_8a994c90f60ad404 *__obf_17d0cfecf7e687b6, __obf_26f3bdcd35e04593 uint64, __obf_3e3f0912cbecfbf3 int, __obf_25725b9a4386f10b *__obf_ce035e3885d5f561) {
	// If mantissa is zero, the number is zero; stop now.
	if __obf_26f3bdcd35e04593 == 0 {
		__obf_8a994c90f60ad404.__obf_4673196fea260de2 = 0
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
	__obf_b78d2ac59239a7be := __obf_25725b9a4386f10b.__obf_2d0469f7185b0ad4 + 1 // minimum possible exponent
	if __obf_3e3f0912cbecfbf3 > __obf_b78d2ac59239a7be && 332*(__obf_8a994c90f60ad404.__obf_26344dec98644d1b-__obf_8a994c90f60ad404.__obf_4673196fea260de2) >= 100*(__obf_3e3f0912cbecfbf3-int(__obf_25725b9a4386f10b.__obf_714ef5fd93ab5467)) {
		// The number is already shortest.
		return
	}

	// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	__obf_883ccfb82b4d093d := new(__obf_17d0cfecf7e687b6)
	__obf_883ccfb82b4d093d.Assign(__obf_26f3bdcd35e04593*2 + 1)
	__obf_883ccfb82b4d093d.Shift(__obf_3e3f0912cbecfbf3 - int(__obf_25725b9a4386f10b.__obf_714ef5fd93ab5467) - 1)

	// d = mant << (exp - mantbits)
	// Next lowest floating point number is mant-1 << exp-mantbits,
	// unless mant-1 drops the significant bit and exp is not the minimum exp,
	// in which case the next lowest is mant*2-1 << exp-mantbits-1.
	// Either way, call it mantlo << explo-mantbits.
	// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
	var __obf_eac771f30279efc7 uint64
	var __obf_a4173e56ab9c94a6 int
	if __obf_26f3bdcd35e04593 > 1<<__obf_25725b9a4386f10b.__obf_714ef5fd93ab5467 || __obf_3e3f0912cbecfbf3 == __obf_b78d2ac59239a7be {
		__obf_eac771f30279efc7 = __obf_26f3bdcd35e04593 - 1
		__obf_a4173e56ab9c94a6 = __obf_3e3f0912cbecfbf3
	} else {
		__obf_eac771f30279efc7 = __obf_26f3bdcd35e04593*2 - 1
		__obf_a4173e56ab9c94a6 = __obf_3e3f0912cbecfbf3 - 1
	}
	__obf_2932d6419e93d9de := new(__obf_17d0cfecf7e687b6)
	__obf_2932d6419e93d9de.Assign(__obf_eac771f30279efc7*2 + 1)
	__obf_2932d6419e93d9de.Shift(__obf_a4173e56ab9c94a6 - int(__obf_25725b9a4386f10b.__obf_714ef5fd93ab5467) - 1)

	// The upper and lower bounds are possible outputs only if
	// the original mantissa is even, so that IEEE round-to-even
	// would round to the original mantissa and not the neighbors.
	__obf_40839d3f4b548a75 := __obf_26f3bdcd35e04593%2 == 0

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
	var __obf_b4a36813310ca8c2 uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for __obf_7b3350869fc2881c := 0; ; __obf_7b3350869fc2881c++ {
		// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		__obf_83117481bd754d5b := __obf_7b3350869fc2881c - __obf_883ccfb82b4d093d.__obf_26344dec98644d1b + __obf_8a994c90f60ad404.__obf_26344dec98644d1b
		if __obf_83117481bd754d5b >= __obf_8a994c90f60ad404.__obf_4673196fea260de2 {
			break
		}
		__obf_7f7e7ba5f44f19f6 := __obf_7b3350869fc2881c - __obf_883ccfb82b4d093d.__obf_26344dec98644d1b + __obf_2932d6419e93d9de.__obf_26344dec98644d1b
		__obf_ff8ac364fdc59939 := byte('0') // lower digit
		if __obf_7f7e7ba5f44f19f6 >= 0 && __obf_7f7e7ba5f44f19f6 < __obf_2932d6419e93d9de.__obf_4673196fea260de2 {
			__obf_ff8ac364fdc59939 = __obf_2932d6419e93d9de.__obf_8a994c90f60ad404[__obf_7f7e7ba5f44f19f6]
		}
		__obf_1a0547a84e8f209a := byte('0') // middle digit
		if __obf_83117481bd754d5b >= 0 {
			__obf_1a0547a84e8f209a = __obf_8a994c90f60ad404.__obf_8a994c90f60ad404[__obf_83117481bd754d5b]
		}
		__obf_4afb16ec76805068 := byte('0') // upper digit
		if __obf_7b3350869fc2881c < __obf_883ccfb82b4d093d.__obf_4673196fea260de2 {
			__obf_4afb16ec76805068 = __obf_883ccfb82b4d093d.__obf_8a994c90f60ad404[__obf_7b3350869fc2881c]
		}

		// Okay to round down (truncate) if lower has a different digit
		// or if lower is inclusive and is exactly the result of rounding
		// down (i.e., and we have reached the final digit of lower).
		__obf_761cec626026b41f := __obf_ff8ac364fdc59939 != __obf_1a0547a84e8f209a || __obf_40839d3f4b548a75 && __obf_7f7e7ba5f44f19f6+1 == __obf_2932d6419e93d9de.__obf_4673196fea260de2

		switch {
		case __obf_b4a36813310ca8c2 == 0 && __obf_1a0547a84e8f209a+1 < __obf_4afb16ec76805068:
			// Example:
			// m = 12345xxx
			// u = 12347xxx
			__obf_b4a36813310ca8c2 = 2
		case __obf_b4a36813310ca8c2 == 0 && __obf_1a0547a84e8f209a != __obf_4afb16ec76805068:
			// Example:
			// m = 12345xxx
			// u = 12346xxx
			__obf_b4a36813310ca8c2 = 1
		case __obf_b4a36813310ca8c2 == 1 && (__obf_1a0547a84e8f209a != '9' || __obf_4afb16ec76805068 != '0'):
			// Example:
			// m = 1234598x
			// u = 1234600x
			__obf_b4a36813310ca8c2 = 2
		}
		// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		__obf_cde7d7b0268abb7b := __obf_b4a36813310ca8c2 > 0 && (__obf_40839d3f4b548a75 || __obf_b4a36813310ca8c2 > 1 || __obf_7b3350869fc2881c+1 < __obf_883ccfb82b4d093d.__obf_4673196fea260de2)

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		switch {
		case __obf_761cec626026b41f && __obf_cde7d7b0268abb7b:
			__obf_8a994c90f60ad404.Round(__obf_83117481bd754d5b + 1)
			return
		case __obf_761cec626026b41f:
			__obf_8a994c90f60ad404.RoundDown(__obf_83117481bd754d5b + 1)
			return
		case __obf_cde7d7b0268abb7b:
			__obf_8a994c90f60ad404.RoundUp(__obf_83117481bd754d5b + 1)
			return
		}
	}
}
