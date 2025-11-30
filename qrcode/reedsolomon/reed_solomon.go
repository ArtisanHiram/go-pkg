// go-qrcode
// Copyright 2014 Tom Harwood

// Package reedsolomon provides error correction encoding for QR Code 2005.
//
// QR Code 2005 uses a Reed-Solomon error correcting code to detect and correct
// errors encountered during decoding.
//
// The generated RS codes are systematic, and consist of the input data with
// error correction bytes appended.
package __obf_0a6b19e606a79505

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
func Encode(__obf_28ed8d33917ceb06 *bitset.Bitset, __obf_82a5714b935a061f int) *bitset.Bitset {
	__obf_2fa0330128234d19 := // Create a polynomial representing |data|.
		//
		// The bytes are interpreted as the sequence of coefficients of a polynomial.
		// The last byte's value becomes the x^0 coefficient, the second to last
		// becomes the x^1 coefficient and so on.
		__obf_c64ac0e17e6e6c02(__obf_28ed8d33917ceb06)
	__obf_2fa0330128234d19 = __obf_19e74ed5348a919f(__obf_2fa0330128234d19, __obf_e9fb1cfcff10030e(__obf_942f4a8f8e23d198,

		// Pick the generator polynomial.
		__obf_82a5714b935a061f))
	__obf_6852816da53e6b55 := __obf_32a531d245aec6b6(__obf_82a5714b935a061f)
	__obf_9d5e81d223c58052 := // Generate the error correction bytes.
		__obf_8e9e781c42fba024(__obf_2fa0330128234d19,

			// Combine the data & error correcting bytes.
			// The mathematically correct answer is:
			//
			//	result := gfPolyAdd(ecpoly, remainder).
			//
			// The encoding used by QR Code 2005 is slightly different this result: To
			// preserve the original |data| bit sequence exactly, the data and remainder
			// are combined manually below. This ensures any most significant zero bits
			// are preserved (and not optimised away).
			__obf_6852816da53e6b55)
	__obf_1d5fad7dc8e9e1a7 := bitset.Clone(__obf_28ed8d33917ceb06)
	__obf_1d5fad7dc8e9e1a7.
		AppendBytes(__obf_9d5e81d223c58052.__obf_28ed8d33917ceb06(__obf_82a5714b935a061f))

	return __obf_1d5fad7dc8e9e1a7
}

// rsGeneratorPoly returns the Reed-Solomon generator polynomial with |degree|.
//
// The generator polynomial is calculated as:
// (x + a^0)(x + a^1)...(x + a^degree-1)
func __obf_32a531d245aec6b6(__obf_60b1039746a15fe5 int) __obf_c6021ed06d3a96ab {
	if __obf_60b1039746a15fe5 < 2 {
		log.Panic("degree < 2")
	}
	__obf_6852816da53e6b55 := __obf_c6021ed06d3a96ab{__obf_930e137791b175f5: []__obf_1005b52d0b467f2a{1}}

	for __obf_18c84b23c99ce50d := 0; __obf_18c84b23c99ce50d < __obf_60b1039746a15fe5; __obf_18c84b23c99ce50d++ {
		__obf_10b3094e78abf0fb := __obf_c6021ed06d3a96ab{__obf_930e137791b175f5: []__obf_1005b52d0b467f2a{__obf_dd9a0020a431b944[__obf_18c84b23c99ce50d], 1}}
		__obf_6852816da53e6b55 = __obf_19e74ed5348a919f(__obf_6852816da53e6b55, __obf_10b3094e78abf0fb)
	}

	return __obf_6852816da53e6b55
}
