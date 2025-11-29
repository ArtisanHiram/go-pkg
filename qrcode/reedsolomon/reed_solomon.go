// go-qrcode
// Copyright 2014 Tom Harwood

// Package reedsolomon provides error correction encoding for QR Code 2005.
//
// QR Code 2005 uses a Reed-Solomon error correcting code to detect and correct
// errors encountered during decoding.
//
// The generated RS codes are systematic, and consist of the input data with
// error correction bytes appended.
package __obf_79572f72dbe37a0e

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
func Encode(__obf_a8518aca69865c96 *bitset.Bitset, __obf_d4688bbb42d68c95 int) *bitset.Bitset {
	__obf_4caa8c4f164888fc := // Create a polynomial representing |data|.
		//
		// The bytes are interpreted as the sequence of coefficients of a polynomial.
		// The last byte's value becomes the x^0 coefficient, the second to last
		// becomes the x^1 coefficient and so on.
		__obf_1c8448dc38dc8598(__obf_a8518aca69865c96)
	__obf_4caa8c4f164888fc = __obf_5bf6d069e50e9834(__obf_4caa8c4f164888fc, __obf_ff98f1989ce99555(__obf_99954ad5772cea02,

		// Pick the generator polynomial.
		__obf_d4688bbb42d68c95))
	__obf_d181413fb4038907 := __obf_7a43c13fc16e630f(__obf_d4688bbb42d68c95)
	__obf_f435504415a4dac9 := // Generate the error correction bytes.
		__obf_27b66f0b936c98aa(__obf_4caa8c4f164888fc,

			// Combine the data & error correcting bytes.
			// The mathematically correct answer is:
			//
			//	result := gfPolyAdd(ecpoly, remainder).
			//
			// The encoding used by QR Code 2005 is slightly different this result: To
			// preserve the original |data| bit sequence exactly, the data and remainder
			// are combined manually below. This ensures any most significant zero bits
			// are preserved (and not optimised away).
			__obf_d181413fb4038907)
	__obf_8214529cf7250f23 := bitset.Clone(__obf_a8518aca69865c96)
	__obf_8214529cf7250f23.
		AppendBytes(__obf_f435504415a4dac9.__obf_a8518aca69865c96(__obf_d4688bbb42d68c95))

	return __obf_8214529cf7250f23
}

// rsGeneratorPoly returns the Reed-Solomon generator polynomial with |degree|.
//
// The generator polynomial is calculated as:
// (x + a^0)(x + a^1)...(x + a^degree-1)
func __obf_7a43c13fc16e630f(__obf_a8bb449e9bfb0f73 int) __obf_c9418fcf8fc08b00 {
	if __obf_a8bb449e9bfb0f73 < 2 {
		log.Panic("degree < 2")
	}
	__obf_d181413fb4038907 := __obf_c9418fcf8fc08b00{__obf_887f3c7bd44a27c9: []__obf_21ecc2aea3abd05b{1}}

	for __obf_d9d01cdc0f623246 := 0; __obf_d9d01cdc0f623246 < __obf_a8bb449e9bfb0f73; __obf_d9d01cdc0f623246++ {
		__obf_320631a2c3fc8260 := __obf_c9418fcf8fc08b00{__obf_887f3c7bd44a27c9: []__obf_21ecc2aea3abd05b{__obf_b5d4b82062805337[__obf_d9d01cdc0f623246], 1}}
		__obf_d181413fb4038907 = __obf_5bf6d069e50e9834(__obf_d181413fb4038907, __obf_320631a2c3fc8260)
	}

	return __obf_d181413fb4038907
}
