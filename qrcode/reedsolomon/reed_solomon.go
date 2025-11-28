// go-qrcode
// Copyright 2014 Tom Harwood

// Package reedsolomon provides error correction encoding for QR Code 2005.
//
// QR Code 2005 uses a Reed-Solomon error correcting code to detect and correct
// errors encountered during decoding.
//
// The generated RS codes are systematic, and consist of the input data with
// error correction bytes appended.
package __obf_6d5da6b2ca65c82e

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
func Encode(__obf_92f02e566ccfb5ff *bitset.Bitset, __obf_c2cbca6726bbb9ae int) *bitset.Bitset {
	// Create a polynomial representing |data|.
	//
	// The bytes are interpreted as the sequence of coefficients of a polynomial.
	// The last byte's value becomes the x^0 coefficient, the second to last
	// becomes the x^1 coefficient and so on.
	__obf_3380eb824cdf1232 := __obf_daafd0c6eef1f175(__obf_92f02e566ccfb5ff)
	__obf_3380eb824cdf1232 = __obf_227da40f9c6c41f3(__obf_3380eb824cdf1232, __obf_3f8089776f332a11(__obf_2a87e72cea132955, __obf_c2cbca6726bbb9ae))

	// Pick the generator polynomial.
	__obf_07a2d9551d55e893 := __obf_ce04f35f81f52030(__obf_c2cbca6726bbb9ae)

	// Generate the error correction bytes.
	__obf_8977186fbcc7b181 := __obf_888a3a462522e873(__obf_3380eb824cdf1232, __obf_07a2d9551d55e893)

	// Combine the data & error correcting bytes.
	// The mathematically correct answer is:
	//
	//	result := gfPolyAdd(ecpoly, remainder).
	//
	// The encoding used by QR Code 2005 is slightly different this result: To
	// preserve the original |data| bit sequence exactly, the data and remainder
	// are combined manually below. This ensures any most significant zero bits
	// are preserved (and not optimised away).
	__obf_b8d4231c43bc645b := bitset.Clone(__obf_92f02e566ccfb5ff)
	__obf_b8d4231c43bc645b.AppendBytes(__obf_8977186fbcc7b181.__obf_92f02e566ccfb5ff(__obf_c2cbca6726bbb9ae))

	return __obf_b8d4231c43bc645b
}

// rsGeneratorPoly returns the Reed-Solomon generator polynomial with |degree|.
//
// The generator polynomial is calculated as:
// (x + a^0)(x + a^1)...(x + a^degree-1)
func __obf_ce04f35f81f52030(__obf_d68be180ca1279c1 int) __obf_40186e28e54c6d61 {
	if __obf_d68be180ca1279c1 < 2 {
		log.Panic("degree < 2")
	}

	__obf_07a2d9551d55e893 := __obf_40186e28e54c6d61{__obf_747aa4fb15a9f618: []__obf_29e803ad00f0d12b{1}}

	for __obf_85aff0c19ced1f98 := 0; __obf_85aff0c19ced1f98 < __obf_d68be180ca1279c1; __obf_85aff0c19ced1f98++ {
		__obf_f6436e10c9ad932b := __obf_40186e28e54c6d61{__obf_747aa4fb15a9f618: []__obf_29e803ad00f0d12b{__obf_654399568b579db1[__obf_85aff0c19ced1f98], 1}}
		__obf_07a2d9551d55e893 = __obf_227da40f9c6c41f3(__obf_07a2d9551d55e893, __obf_f6436e10c9ad932b)
	}

	return __obf_07a2d9551d55e893
}
