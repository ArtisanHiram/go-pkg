// go-qrcode
// Copyright 2014 Tom Harwood

// Package reedsolomon provides error correction encoding for QR Code 2005.
//
// QR Code 2005 uses a Reed-Solomon error correcting code to detect and correct
// errors encountered during decoding.
//
// The generated RS codes are systematic, and consist of the input data with
// error correction bytes appended.
package __obf_13fa3c4d8cc2a5c1

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
func Encode(__obf_14fbbbd1ec149da3 *bitset.Bitset, __obf_87b30dc83d6a31bb int) *bitset.Bitset {
	__obf_297a6476b872cc45 := // Create a polynomial representing |data|.
		//
		// The bytes are interpreted as the sequence of coefficients of a polynomial.
		// The last byte's value becomes the x^0 coefficient, the second to last
		// becomes the x^1 coefficient and so on.
		__obf_b21c3e76a0cbcb0b(__obf_14fbbbd1ec149da3)
	__obf_297a6476b872cc45 = __obf_b680f3d50eea2b57(__obf_297a6476b872cc45, __obf_d07c34c23e218f0f(__obf_c56ed7ddf95cbae1,

		// Pick the generator polynomial.
		__obf_87b30dc83d6a31bb))
	__obf_e7de7ef68360e178 := __obf_55ddf87ae5b2d9e7(__obf_87b30dc83d6a31bb)
	__obf_7f08354c0e804ad1 := // Generate the error correction bytes.
		__obf_b5870e3900b6d078(__obf_297a6476b872cc45,

			// Combine the data & error correcting bytes.
			// The mathematically correct answer is:
			//
			//	result := gfPolyAdd(ecpoly, remainder).
			//
			// The encoding used by QR Code 2005 is slightly different this result: To
			// preserve the original |data| bit sequence exactly, the data and remainder
			// are combined manually below. This ensures any most significant zero bits
			// are preserved (and not optimised away).
			__obf_e7de7ef68360e178)
	__obf_d8a1da44e57d6011 := bitset.Clone(__obf_14fbbbd1ec149da3)
	__obf_d8a1da44e57d6011.
		AppendBytes(__obf_7f08354c0e804ad1.__obf_14fbbbd1ec149da3(__obf_87b30dc83d6a31bb))

	return __obf_d8a1da44e57d6011
}

// rsGeneratorPoly returns the Reed-Solomon generator polynomial with |degree|.
//
// The generator polynomial is calculated as:
// (x + a^0)(x + a^1)...(x + a^degree-1)
func __obf_55ddf87ae5b2d9e7(__obf_b41724be7f1bc628 int) __obf_087436cf3ba4af81 {
	if __obf_b41724be7f1bc628 < 2 {
		log.Panic("degree < 2")
	}
	__obf_e7de7ef68360e178 := __obf_087436cf3ba4af81{__obf_90c5f727b336f89b: []__obf_11b619b7d43c1a0d{1}}

	for __obf_c4dd66f72f8a479f := 0; __obf_c4dd66f72f8a479f < __obf_b41724be7f1bc628; __obf_c4dd66f72f8a479f++ {
		__obf_1001c64e194dd46a := __obf_087436cf3ba4af81{__obf_90c5f727b336f89b: []__obf_11b619b7d43c1a0d{__obf_279beb3968a66a07[__obf_c4dd66f72f8a479f], 1}}
		__obf_e7de7ef68360e178 = __obf_b680f3d50eea2b57(__obf_e7de7ef68360e178, __obf_1001c64e194dd46a)
	}

	return __obf_e7de7ef68360e178
}
