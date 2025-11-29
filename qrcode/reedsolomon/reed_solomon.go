// go-qrcode
// Copyright 2014 Tom Harwood

// Package reedsolomon provides error correction encoding for QR Code 2005.
//
// QR Code 2005 uses a Reed-Solomon error correcting code to detect and correct
// errors encountered during decoding.
//
// The generated RS codes are systematic, and consist of the input data with
// error correction bytes appended.
package __obf_d5be46fdf9a047aa

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
func Encode(__obf_6a71e09d65c25d89 *bitset.Bitset, __obf_17782ba2a2600ecf int) *bitset.Bitset {
	__obf_e9a7a4633f9c15b0 := // Create a polynomial representing |data|.
		//
		// The bytes are interpreted as the sequence of coefficients of a polynomial.
		// The last byte's value becomes the x^0 coefficient, the second to last
		// becomes the x^1 coefficient and so on.
		__obf_b72e69aa5faca050(__obf_6a71e09d65c25d89)
	__obf_e9a7a4633f9c15b0 = __obf_83e2645b1ae7eb30(__obf_e9a7a4633f9c15b0, __obf_0498c51b59c427a4(__obf_336334426df17940,

		// Pick the generator polynomial.
		__obf_17782ba2a2600ecf))
	__obf_4d0a707493e77e8b := __obf_37f9ae29ad3cfa29(__obf_17782ba2a2600ecf)
	__obf_d87cecb5fbe51cf0 := // Generate the error correction bytes.
		__obf_c170f4171b722bb2(__obf_e9a7a4633f9c15b0,

			// Combine the data & error correcting bytes.
			// The mathematically correct answer is:
			//
			//	result := gfPolyAdd(ecpoly, remainder).
			//
			// The encoding used by QR Code 2005 is slightly different this result: To
			// preserve the original |data| bit sequence exactly, the data and remainder
			// are combined manually below. This ensures any most significant zero bits
			// are preserved (and not optimised away).
			__obf_4d0a707493e77e8b)
	__obf_2efda911f3d2d3c1 := bitset.Clone(__obf_6a71e09d65c25d89)
	__obf_2efda911f3d2d3c1.
		AppendBytes(__obf_d87cecb5fbe51cf0.__obf_6a71e09d65c25d89(__obf_17782ba2a2600ecf))

	return __obf_2efda911f3d2d3c1
}

// rsGeneratorPoly returns the Reed-Solomon generator polynomial with |degree|.
//
// The generator polynomial is calculated as:
// (x + a^0)(x + a^1)...(x + a^degree-1)
func __obf_37f9ae29ad3cfa29(__obf_bce1f4d9d31e3eda int) __obf_681065f079dec5b9 {
	if __obf_bce1f4d9d31e3eda < 2 {
		log.Panic("degree < 2")
	}
	__obf_4d0a707493e77e8b := __obf_681065f079dec5b9{__obf_f7ff1cffece1f4f8: []__obf_94bafdc0d5ce6bf6{1}}

	for __obf_183af01b71412ca5 := 0; __obf_183af01b71412ca5 < __obf_bce1f4d9d31e3eda; __obf_183af01b71412ca5++ {
		__obf_e9847248c3cfd651 := __obf_681065f079dec5b9{__obf_f7ff1cffece1f4f8: []__obf_94bafdc0d5ce6bf6{__obf_2b1bf999fbc85162[__obf_183af01b71412ca5], 1}}
		__obf_4d0a707493e77e8b = __obf_83e2645b1ae7eb30(__obf_4d0a707493e77e8b, __obf_e9847248c3cfd651)
	}

	return __obf_4d0a707493e77e8b
}
