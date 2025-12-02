// go-qrcode
// Copyright 2014 Tom Harwood

// Package reedsolomon provides error correction encoding for QR Code 2005.
//
// QR Code 2005 uses a Reed-Solomon error correcting code to detect and correct
// errors encountered during decoding.
//
// The generated RS codes are systematic, and consist of the input data with
// error correction bytes appended.
package __obf_e1f3dcd8d719d52b

import (
	"log"

	bitset "github.com/ArtisanHiram/go-pkg/qrcode/bitset"
)

// Encode data for QR Code 2005 using the appropriate Reed-Solomon code.
//
// numECBytes is the number of error correction bytes to append, and is
// determined by the target QR Code's version and error correction level.
//
// ISO/IEC 18004 table 9 specifies the numECBytes required. e.g. a 1-L code has
// numECBytes=7.
func Encode(__obf_e84c8ad5f9fe4eba *bitset.Bitset, __obf_3ac76bf083414f56 int) *bitset.Bitset {
	__obf_f515a4277019914c := // Create a polynomial representing |data|.
		//
		// The bytes are interpreted as the sequence of coefficients of a polynomial.
		// The last byte's value becomes the x^0 coefficient, the second to last
		// becomes the x^1 coefficient and so on.
		__obf_8c8e622b7e5d1808(__obf_e84c8ad5f9fe4eba)
	__obf_f515a4277019914c = __obf_64c6d4eebca914e2(__obf_f515a4277019914c, __obf_52f952075352cc74(__obf_98a3f50d23acf486,

		// Pick the generator polynomial.
		__obf_3ac76bf083414f56))
	__obf_133515636299b453 := __obf_c6892ee1b96d137d(__obf_3ac76bf083414f56)
	__obf_0fed35ba81544ce5 := // Generate the error correction bytes.
		__obf_b38b77e4cdfbc62a(__obf_f515a4277019914c,

			// Combine the data & error correcting bytes.
			// The mathematically correct answer is:
			//
			//	result := gfPolyAdd(ecpoly, remainder).
			//
			// The encoding used by QR Code 2005 is slightly different this result: To
			// preserve the original |data| bit sequence exactly, the data and remainder
			// are combined manually below. This ensures any most significant zero bits
			// are preserved (and not optimised away).
			__obf_133515636299b453)
	__obf_903beff2487e2326 := bitset.Clone(__obf_e84c8ad5f9fe4eba)
	__obf_903beff2487e2326.
		AppendBytes(__obf_0fed35ba81544ce5.__obf_e84c8ad5f9fe4eba(__obf_3ac76bf083414f56))

	return __obf_903beff2487e2326
}

// rsGeneratorPoly returns the Reed-Solomon generator polynomial with |degree|.
//
// The generator polynomial is calculated as:
// (x + a^0)(x + a^1)...(x + a^degree-1)
func __obf_c6892ee1b96d137d(__obf_eee26701a2608609 int) __obf_7f0a87df1252e0f1 {
	if __obf_eee26701a2608609 < 2 {
		log.Panic("degree < 2")
	}
	__obf_133515636299b453 := __obf_7f0a87df1252e0f1{__obf_dc2d266987a80f84: []__obf_69935e5d788798dd{1}}

	for __obf_13b2631ac20e0422 := 0; __obf_13b2631ac20e0422 < __obf_eee26701a2608609; __obf_13b2631ac20e0422++ {
		__obf_2b1b0b6f6f1d2f46 := __obf_7f0a87df1252e0f1{__obf_dc2d266987a80f84: []__obf_69935e5d788798dd{__obf_0548bc2533e0f272[__obf_13b2631ac20e0422], 1}}
		__obf_133515636299b453 = __obf_64c6d4eebca914e2(__obf_133515636299b453, __obf_2b1b0b6f6f1d2f46)
	}

	return __obf_133515636299b453
}
