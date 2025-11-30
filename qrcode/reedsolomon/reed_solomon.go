// go-qrcode
// Copyright 2014 Tom Harwood

// Package reedsolomon provides error correction encoding for QR Code 2005.
//
// QR Code 2005 uses a Reed-Solomon error correcting code to detect and correct
// errors encountered during decoding.
//
// The generated RS codes are systematic, and consist of the input data with
// error correction bytes appended.
package __obf_eba3d68b9605015e

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
func Encode(__obf_4e04dfa10559e27c *bitset.Bitset, __obf_17c35e05e9dc2293 int) *bitset.Bitset {
	__obf_df8fb7bcb0b7004f := // Create a polynomial representing |data|.
		//
		// The bytes are interpreted as the sequence of coefficients of a polynomial.
		// The last byte's value becomes the x^0 coefficient, the second to last
		// becomes the x^1 coefficient and so on.
		__obf_f3a458a9c566d56a(__obf_4e04dfa10559e27c)
	__obf_df8fb7bcb0b7004f = __obf_613fea38d6c97b2d(__obf_df8fb7bcb0b7004f, __obf_a2a64cfa09775cd0(__obf_dfd4e77e4658e35b,

		// Pick the generator polynomial.
		__obf_17c35e05e9dc2293))
	__obf_b5231f3d7c46bd4a := __obf_e9e8efce04b892b3(__obf_17c35e05e9dc2293)
	__obf_9c353bc1b664ae8f := // Generate the error correction bytes.
		__obf_9b723898ea675104(__obf_df8fb7bcb0b7004f,

			// Combine the data & error correcting bytes.
			// The mathematically correct answer is:
			//
			//	result := gfPolyAdd(ecpoly, remainder).
			//
			// The encoding used by QR Code 2005 is slightly different this result: To
			// preserve the original |data| bit sequence exactly, the data and remainder
			// are combined manually below. This ensures any most significant zero bits
			// are preserved (and not optimised away).
			__obf_b5231f3d7c46bd4a)
	__obf_7fca30a86b604ba4 := bitset.Clone(__obf_4e04dfa10559e27c)
	__obf_7fca30a86b604ba4.
		AppendBytes(__obf_9c353bc1b664ae8f.__obf_4e04dfa10559e27c(__obf_17c35e05e9dc2293))

	return __obf_7fca30a86b604ba4
}

// rsGeneratorPoly returns the Reed-Solomon generator polynomial with |degree|.
//
// The generator polynomial is calculated as:
// (x + a^0)(x + a^1)...(x + a^degree-1)
func __obf_e9e8efce04b892b3(__obf_bf07404dd6aeed84 int) __obf_71eb0882383cb0ef {
	if __obf_bf07404dd6aeed84 < 2 {
		log.Panic("degree < 2")
	}
	__obf_b5231f3d7c46bd4a := __obf_71eb0882383cb0ef{__obf_ea46b5761247c020: []__obf_dbf34f1740ee9966{1}}

	for __obf_a2ccc63c916c0db9 := 0; __obf_a2ccc63c916c0db9 < __obf_bf07404dd6aeed84; __obf_a2ccc63c916c0db9++ {
		__obf_6d78e93227b888c2 := __obf_71eb0882383cb0ef{__obf_ea46b5761247c020: []__obf_dbf34f1740ee9966{__obf_4859dffeb395a9ff[__obf_a2ccc63c916c0db9], 1}}
		__obf_b5231f3d7c46bd4a = __obf_613fea38d6c97b2d(__obf_b5231f3d7c46bd4a, __obf_6d78e93227b888c2)
	}

	return __obf_b5231f3d7c46bd4a
}
