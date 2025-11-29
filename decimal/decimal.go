// Package decimal implements an arbitrary precision fixed-point decimal.
//
// The zero-value of a Decimal is 0, as you would expect.
//
// The best way to create a new Decimal is to use decimal.NewFromString, ex:
//
//	n, err := decimal.NewFromString("-123.4567")
//	n.String() // output: "-123.4567"
//
// To use Decimal as part of a struct:
//
//	type StructName struct {
//	    Number Decimal
//	}
//
// Note: This can "only" represent numbers with a maximum of 2^31 digits after the decimal point.
package __obf_4cc30aac3ca67486

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

// DivisionPrecision is the number of decimal places in the result when it
// doesn't divide exactly.
//
// Example:
//
//	d1 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3))
//	d1.String() // output: "0.6666666666666667"
//	d2 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(30000))
//	d2.String() // output: "0.0000666666666667"
//	d3 := decimal.NewFromFloat(20000).Div(decimal.NewFromFloat(3))
//	d3.String() // output: "6666.6666666666666667"
//	decimal.DivisionPrecision = 3
//	d4 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3))
//	d4.String() // output: "0.667"
var DivisionPrecision = 16

// PowPrecisionNegativeExponent specifies the maximum precision of the result (digits after decimal point)
// when calculating decimal power. Only used for cases where the exponent is a negative number.
// This constant applies to Pow, PowInt32 and PowBigInt methods, PowWithPrecision method is not constrained by it.
//
// Example:
//
//	d1, err := decimal.NewFromFloat(15.2).PowInt32(-2)
//	d1.String() // output: "0.0043282548476454"
//
//	decimal.PowPrecisionNegativeExponent = 24
//	d2, err := decimal.NewFromFloat(15.2).PowInt32(-2)
//	d2.String() // output: "0.004328254847645429362881"
var PowPrecisionNegativeExponent = 16

// MarshalJSONWithoutQuotes should be set to true if you want the decimal to
// be JSON marshaled as a number, instead of as a string.
// WARNING: this is dangerous for decimals with many digits, since many JSON
// unmarshallers (ex: Javascript's) will unmarshal JSON numbers to IEEE 754
// double-precision floating point numbers, which means you can potentially
// silently lose precision.
var MarshalJSONWithoutQuotes = false

// ExpMaxIterations specifies the maximum number of iterations needed to calculate
// precise natural exponent value using ExpHullAbrham method.
var ExpMaxIterations = 1000

// Zero constant, to make computations faster.
// Zero should never be compared with == or != directly, please use decimal.Equal or decimal.Cmp instead.
var Zero = New(0, 1)

var __obf_f0e1aaf9670d8421 = big.NewInt(0)
var __obf_f136d7fff3e7d6f7 = big.NewInt(1)
var __obf_3f8f48b9238d224e = big.NewInt(2)
var __obf_bee298c3e91651d7 = big.NewInt(4)
var __obf_3f87d09b59f6435e = big.NewInt(5)
var __obf_cd9358bd6469dc36 = big.NewInt(10)
var __obf_abb4b37cdf2c9d6e = big.NewInt(20)

var __obf_2b7a96df87ff7b2b = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_22dafd20a48df74e *big.Int
	__obf_10eec1d7f8ad279f int32 // NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.

}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_22dafd20a48df74e int64, __obf_10eec1d7f8ad279f int32) Decimal {
	return Decimal{__obf_22dafd20a48df74e: big.NewInt(__obf_22dafd20a48df74e), __obf_10eec1d7f8ad279f: __obf_10eec1d7f8ad279f}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_22dafd20a48df74e int64) Decimal {
	return Decimal{__obf_22dafd20a48df74e: big.NewInt(__obf_22dafd20a48df74e), __obf_10eec1d7f8ad279f: 0}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_22dafd20a48df74e int32) Decimal {
	return Decimal{__obf_22dafd20a48df74e: big.NewInt(int64(__obf_22dafd20a48df74e)), __obf_10eec1d7f8ad279f: 0}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_22dafd20a48df74e uint64) Decimal {
	return Decimal{__obf_22dafd20a48df74e: new(big.Int).SetUint64(__obf_22dafd20a48df74e), __obf_10eec1d7f8ad279f: 0}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_22dafd20a48df74e *big.Int, __obf_10eec1d7f8ad279f int32) Decimal {
	return Decimal{__obf_22dafd20a48df74e: new(big.Int).Set(__obf_22dafd20a48df74e), __obf_10eec1d7f8ad279f: __obf_10eec1d7f8ad279f}
}

// NewFromBigRat returns a new Decimal from a big.Rat. The numerator and
// denominator are divided and rounded to the given precision.
//
// Example:
//
//	d1 := NewFromBigRat(big.NewRat(0, 1), 0)    // output: "0"
//	d2 := NewFromBigRat(big.NewRat(4, 5), 1)    // output: "0.8"
//	d3 := NewFromBigRat(big.NewRat(1000, 3), 3) // output: "333.333"
//	d4 := NewFromBigRat(big.NewRat(2, 7), 4)    // output: "0.2857"
func NewFromBigRat(__obf_22dafd20a48df74e *big.Rat, __obf_6eb04b4176b00c83 int32) Decimal {
	return Decimal{__obf_22dafd20a48df74e: new(big.Int).Set(__obf_22dafd20a48df74e.Num()), __obf_10eec1d7f8ad279f: 0}.DivRound(Decimal{__obf_22dafd20a48df74e: new(big.Int).Set(__obf_22dafd20a48df74e.Denom()), __obf_10eec1d7f8ad279f: 0}, __obf_6eb04b4176b00c83)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_22dafd20a48df74e string) (Decimal, error) {
	__obf_21f07081234a3c7d := __obf_22dafd20a48df74e
	var __obf_15f348dac6330403 string
	var __obf_10eec1d7f8ad279f int64
	__obf_5b84dc812de10589 := // Check if number is using scientific notation
		strings.IndexAny(__obf_22dafd20a48df74e, "Ee")
	if __obf_5b84dc812de10589 != -1 {
		__obf_1a43225f7f4e9492, __obf_2e2f1e21b557821f := strconv.ParseInt(__obf_22dafd20a48df74e[__obf_5b84dc812de10589+1:], 10, 32)
		if __obf_2e2f1e21b557821f != nil {
			if __obf_e338d3262f43bafa, __obf_12475deb05c894e6 := __obf_2e2f1e21b557821f.(*strconv.NumError); __obf_12475deb05c894e6 && __obf_e338d3262f43bafa.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_22dafd20a48df74e)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_22dafd20a48df74e)
		}
		__obf_22dafd20a48df74e = __obf_22dafd20a48df74e[:__obf_5b84dc812de10589]
		__obf_10eec1d7f8ad279f = __obf_1a43225f7f4e9492
	}
	__obf_ddb5b67eb1194b03 := -1
	__obf_86f5fcbbe5428372 := len(__obf_22dafd20a48df74e)
	for __obf_03985402d384eb80 := 0; __obf_03985402d384eb80 < __obf_86f5fcbbe5428372; __obf_03985402d384eb80++ {
		if __obf_22dafd20a48df74e[__obf_03985402d384eb80] == '.' {
			if __obf_ddb5b67eb1194b03 > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_22dafd20a48df74e)
			}
			__obf_ddb5b67eb1194b03 = __obf_03985402d384eb80
		}
	}

	if __obf_ddb5b67eb1194b03 == -1 {
		__obf_15f348dac6330403 = // There is no decimal point, we can just parse the original string as
			// an int
			__obf_22dafd20a48df74e
	} else {
		if __obf_ddb5b67eb1194b03+1 < __obf_86f5fcbbe5428372 {
			__obf_15f348dac6330403 = __obf_22dafd20a48df74e[:__obf_ddb5b67eb1194b03] + __obf_22dafd20a48df74e[__obf_ddb5b67eb1194b03+1:]
		} else {
			__obf_15f348dac6330403 = __obf_22dafd20a48df74e[:__obf_ddb5b67eb1194b03]
		}
		__obf_1a43225f7f4e9492 := -len(__obf_22dafd20a48df74e[__obf_ddb5b67eb1194b03+1:])
		__obf_10eec1d7f8ad279f += int64(__obf_1a43225f7f4e9492)
	}

	var __obf_ab9cfcacb27251b9 *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_15f348dac6330403) <= 18 {
		__obf_5d47bb6623ed7dcd, __obf_2e2f1e21b557821f := strconv.ParseInt(__obf_15f348dac6330403, 10, 64)
		if __obf_2e2f1e21b557821f != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_22dafd20a48df74e)
		}
		__obf_ab9cfcacb27251b9 = big.NewInt(__obf_5d47bb6623ed7dcd)
	} else {
		__obf_ab9cfcacb27251b9 = new(big.Int)
		_, __obf_12475deb05c894e6 := __obf_ab9cfcacb27251b9.SetString(__obf_15f348dac6330403, 10)
		if !__obf_12475deb05c894e6 {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_22dafd20a48df74e)
		}
	}

	if __obf_10eec1d7f8ad279f < math.MinInt32 || __obf_10eec1d7f8ad279f > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_21f07081234a3c7d)
	}

	return Decimal{__obf_22dafd20a48df74e: __obf_ab9cfcacb27251b9, __obf_10eec1d7f8ad279f: int32(__obf_10eec1d7f8ad279f)}, nil
}

// NewFromFormattedString returns a new Decimal from a formatted string representation.
// The second argument - replRegexp, is a regular expression that is used to find characters that should be
// removed from given decimal string representation. All matched characters will be replaced with an empty string.
//
// Example:
//
//	r := regexp.MustCompile("[$,]")
//	d1, err := NewFromFormattedString("$5,125.99", r)
//
//	r2 := regexp.MustCompile("[_]")
//	d2, err := NewFromFormattedString("1_000_000", r2)
//
//	r3 := regexp.MustCompile("[USD\\s]")
//	d3, err := NewFromFormattedString("5000 USD", r3)
func NewFromFormattedString(__obf_22dafd20a48df74e string, __obf_08cdc2611608dfac *regexp.Regexp) (Decimal, error) {
	__obf_f80cc1f8deba4b67 := __obf_08cdc2611608dfac.ReplaceAllString(__obf_22dafd20a48df74e, "")
	__obf_aafc5e233dfea784, __obf_2e2f1e21b557821f := NewFromString(__obf_f80cc1f8deba4b67)
	if __obf_2e2f1e21b557821f != nil {
		return Decimal{}, __obf_2e2f1e21b557821f
	}
	return __obf_aafc5e233dfea784, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_22dafd20a48df74e string) Decimal {
	__obf_012da267317dc449, __obf_2e2f1e21b557821f := NewFromString(__obf_22dafd20a48df74e)
	if __obf_2e2f1e21b557821f != nil {
		panic(__obf_2e2f1e21b557821f)
	}
	return __obf_012da267317dc449
}

// NewFromFloat converts a float64 to Decimal.
//
// The converted number will contain the number of significant digits that can be
// represented in a float with reliable roundtrip.
// This is typically 15 digits, but may be more in some cases.
// See https://www.exploringbinary.com/decimal-precision-of-binary-floating-point-numbers/ for more information.
//
// For slightly faster conversion, use NewFromFloatWithExponent where you can specify the precision in absolute terms.
//
// NOTE: this will panic on NaN, +/-inf
func NewFromFloat(__obf_22dafd20a48df74e float64) Decimal {
	if __obf_22dafd20a48df74e == 0 {
		return New(0, 0)
	}
	return __obf_c01e755ba8d11bc9(__obf_22dafd20a48df74e, math.Float64bits(__obf_22dafd20a48df74e), &__obf_f772cba36824086f)
}

// NewFromFloat32 converts a float32 to Decimal.
//
// The converted number will contain the number of significant digits that can be
// represented in a float with reliable roundtrip.
// This is typically 6-8 digits depending on the input.
// See https://www.exploringbinary.com/decimal-precision-of-binary-floating-point-numbers/ for more information.
//
// For slightly faster conversion, use NewFromFloatWithExponent where you can specify the precision in absolute terms.
//
// NOTE: this will panic on NaN, +/-inf
func NewFromFloat32(__obf_22dafd20a48df74e float32) Decimal {
	if __obf_22dafd20a48df74e == 0 {
		return New(0, 0)
	}
	__obf_c09a2b51ef153b53 := // XOR is workaround for https://github.com/golang/go/issues/26285
		math.Float32bits(__obf_22dafd20a48df74e) ^ 0x80808080
	return __obf_c01e755ba8d11bc9(float64(__obf_22dafd20a48df74e), uint64(__obf_c09a2b51ef153b53)^0x80808080, &__obf_5344c617ea3140d5)
}

func __obf_c01e755ba8d11bc9(__obf_3ea3ab1a55b06161 float64, __obf_7cb46ac41c42b72f uint64, __obf_3313f66003031fa5 *__obf_4db791a44d1e22db) Decimal {
	if math.IsNaN(__obf_3ea3ab1a55b06161) || math.IsInf(__obf_3ea3ab1a55b06161, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_3ea3ab1a55b06161))
	}
	__obf_10eec1d7f8ad279f := int(__obf_7cb46ac41c42b72f>>__obf_3313f66003031fa5.__obf_85a8f036d45ff996) & (1<<__obf_3313f66003031fa5.__obf_d0cd6e9a2f63cb34 - 1)
	__obf_7b4a9ee3891ebcd0 := __obf_7cb46ac41c42b72f & (uint64(1)<<__obf_3313f66003031fa5.__obf_85a8f036d45ff996 - 1)

	switch __obf_10eec1d7f8ad279f {
	case 0:
		__obf_10eec1d7f8ad279f++ // denormalized

	default:
		__obf_7b4a9ee3891ebcd0 |= // add implicit top bit
			uint64(1) << __obf_3313f66003031fa5.__obf_85a8f036d45ff996
	}
	__obf_10eec1d7f8ad279f += __obf_3313f66003031fa5.__obf_d5618d67c0a6e865

	var __obf_aafc5e233dfea784 __obf_4cc30aac3ca67486
	__obf_aafc5e233dfea784.
		Assign(__obf_7b4a9ee3891ebcd0)
	__obf_aafc5e233dfea784.
		Shift(__obf_10eec1d7f8ad279f - int(__obf_3313f66003031fa5.__obf_85a8f036d45ff996))
	__obf_aafc5e233dfea784.__obf_ef34c5939301e47a = __obf_7cb46ac41c42b72f>>(__obf_3313f66003031fa5.__obf_d0cd6e9a2f63cb34+__obf_3313f66003031fa5.__obf_85a8f036d45ff996) != 0
	__obf_711c33dc022ef1f9(&__obf_aafc5e233dfea784,
		// If less than 19 digits, we can do calculation in an int64.
		__obf_7b4a9ee3891ebcd0, __obf_10eec1d7f8ad279f, __obf_3313f66003031fa5)

	if __obf_aafc5e233dfea784.__obf_5c82a3d658ffb3ad < 19 {
		__obf_6893e4bf8945b770 := int64(0)
		__obf_ade33d4daf3b3fc1 := int64(1)
		for __obf_03985402d384eb80 := __obf_aafc5e233dfea784.__obf_5c82a3d658ffb3ad - 1; __obf_03985402d384eb80 >= 0; __obf_03985402d384eb80-- {
			__obf_6893e4bf8945b770 += __obf_ade33d4daf3b3fc1 * int64(__obf_aafc5e233dfea784.__obf_aafc5e233dfea784[__obf_03985402d384eb80]-'0')
			__obf_ade33d4daf3b3fc1 *= 10
		}
		if __obf_aafc5e233dfea784.__obf_ef34c5939301e47a {
			__obf_6893e4bf8945b770 *= -1
		}
		return Decimal{__obf_22dafd20a48df74e: big.NewInt(__obf_6893e4bf8945b770), __obf_10eec1d7f8ad279f: int32(__obf_aafc5e233dfea784.__obf_2b78448fb2f1bffe) - int32(__obf_aafc5e233dfea784.__obf_5c82a3d658ffb3ad)}
	}
	__obf_ab9cfcacb27251b9 := new(big.Int)
	__obf_ab9cfcacb27251b9, __obf_12475deb05c894e6 := __obf_ab9cfcacb27251b9.SetString(string(__obf_aafc5e233dfea784.__obf_aafc5e233dfea784[:__obf_aafc5e233dfea784.__obf_5c82a3d658ffb3ad]), 10)
	if __obf_12475deb05c894e6 {
		return Decimal{__obf_22dafd20a48df74e: __obf_ab9cfcacb27251b9, __obf_10eec1d7f8ad279f: int32(__obf_aafc5e233dfea784.__obf_2b78448fb2f1bffe) - int32(__obf_aafc5e233dfea784.__obf_5c82a3d658ffb3ad)}
	}

	return NewFromFloatWithExponent(__obf_3ea3ab1a55b06161, int32(__obf_aafc5e233dfea784.

		// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
		// number of fractional digits.
		//
		// Example:
		//
		//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
		__obf_2b78448fb2f1bffe)-int32(__obf_aafc5e233dfea784.__obf_5c82a3d658ffb3ad))
}

func NewFromFloatWithExponent(__obf_22dafd20a48df74e float64, __obf_10eec1d7f8ad279f int32) Decimal {
	if math.IsNaN(__obf_22dafd20a48df74e) || math.IsInf(__obf_22dafd20a48df74e, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_22dafd20a48df74e))
	}
	__obf_7cb46ac41c42b72f := math.Float64bits(__obf_22dafd20a48df74e)
	__obf_7b4a9ee3891ebcd0 := __obf_7cb46ac41c42b72f & (1<<52 - 1)
	__obf_a0458288142b1969 := int32((__obf_7cb46ac41c42b72f >> 52) & (1<<11 - 1))
	__obf_38ae79a7ec63ef24 := __obf_7cb46ac41c42b72f >> 63

	if __obf_a0458288142b1969 == 0 {
		// specials
		if __obf_7b4a9ee3891ebcd0 == 0 {
			return Decimal{}
		}
		__obf_a0458288142b1969++ // subnormal

	} else {
		__obf_7b4a9ee3891ebcd0 |= // normal
			1 << 52
	}
	__obf_a0458288142b1969 -= 1023 + 52

	// normalizing base-2 values
	for __obf_7b4a9ee3891ebcd0&1 == 0 {
		__obf_7b4a9ee3891ebcd0 = __obf_7b4a9ee3891ebcd0 >> 1
		__obf_a0458288142b1969++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_10eec1d7f8ad279f < 0 && __obf_10eec1d7f8ad279f < __obf_a0458288142b1969 {
		if __obf_a0458288142b1969 < 0 {
			__obf_10eec1d7f8ad279f = __obf_a0458288142b1969
		} else {
			__obf_10eec1d7f8ad279f = 0
		}
	}
	__obf_a0458288142b1969 -= // representing 10^M * 2^N as 5^M * 2^(M+N)
		__obf_10eec1d7f8ad279f
	__obf_f662c272469b1aab := big.NewInt(1)
	__obf_d6942bd962ff16b6 := big.NewInt(int64(__obf_7b4a9ee3891ebcd0))

	// applying 5^M
	if __obf_10eec1d7f8ad279f > 0 {
		__obf_f662c272469b1aab = __obf_f662c272469b1aab.SetInt64(int64(__obf_10eec1d7f8ad279f))
		__obf_f662c272469b1aab = __obf_f662c272469b1aab.Exp(__obf_3f87d09b59f6435e, __obf_f662c272469b1aab, nil)
	} else if __obf_10eec1d7f8ad279f < 0 {
		__obf_f662c272469b1aab = __obf_f662c272469b1aab.SetInt64(-int64(__obf_10eec1d7f8ad279f))
		__obf_f662c272469b1aab = __obf_f662c272469b1aab.Exp(__obf_3f87d09b59f6435e, __obf_f662c272469b1aab, nil)
		__obf_d6942bd962ff16b6 = __obf_d6942bd962ff16b6.Mul(__obf_d6942bd962ff16b6, __obf_f662c272469b1aab)
		__obf_f662c272469b1aab = __obf_f662c272469b1aab.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_a0458288142b1969 > 0 {
		__obf_d6942bd962ff16b6 = __obf_d6942bd962ff16b6.Lsh(__obf_d6942bd962ff16b6, uint(__obf_a0458288142b1969))
	} else if __obf_a0458288142b1969 < 0 {
		__obf_f662c272469b1aab = __obf_f662c272469b1aab.Lsh(__obf_f662c272469b1aab, uint(-__obf_a0458288142b1969))
	}

	// rounding and downscaling
	if __obf_10eec1d7f8ad279f > 0 || __obf_a0458288142b1969 < 0 {
		__obf_d371fa7fd1d7dcc6 := new(big.Int).Rsh(__obf_f662c272469b1aab, 1)
		__obf_d6942bd962ff16b6 = __obf_d6942bd962ff16b6.Add(__obf_d6942bd962ff16b6, __obf_d371fa7fd1d7dcc6)
		__obf_d6942bd962ff16b6 = __obf_d6942bd962ff16b6.Quo(__obf_d6942bd962ff16b6, __obf_f662c272469b1aab)
	}

	if __obf_38ae79a7ec63ef24 == 1 {
		__obf_d6942bd962ff16b6 = __obf_d6942bd962ff16b6.Neg(__obf_d6942bd962ff16b6)
	}

	return Decimal{__obf_22dafd20a48df74e: __obf_d6942bd962ff16b6, __obf_10eec1d7f8ad279f: __obf_10eec1d7f8ad279f}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_aafc5e233dfea784 Decimal) Copy() Decimal {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()
	return Decimal{__obf_22dafd20a48df74e: new(big.Int).Set(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e), __obf_10eec1d7f8ad279f: __obf_aafc5e233dfea784.

		// rescale returns a rescaled version of the decimal. Returned
		// decimal may be less precise if the given exponent is bigger
		// than the initial exponent of the Decimal.
		// NOTE: this will truncate, NOT round
		//
		// Example:
		//
		//	d := New(12345, -4)
		//	d2 := d.rescale(-1)
		//	d3 := d2.rescale(-4)
		//	println(d1)
		//	println(d2)
		//	println(d3)
		//
		// Output:
		//
		//	1.2345
		//	1.2
		//	1.2000
		__obf_10eec1d7f8ad279f,
	}
}

func (__obf_aafc5e233dfea784 Decimal) __obf_51b05428a428dd2d(__obf_10eec1d7f8ad279f int32) Decimal {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()

	if __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f == __obf_10eec1d7f8ad279f {
		return Decimal{
			new(big.Int).Set(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e), __obf_aafc5e233dfea784.

				// NOTE(vadim): must convert exps to float64 before - to prevent overflow
				__obf_10eec1d7f8ad279f,
		}
	}
	__obf_49c750d8ab1e1e10 := math.Abs(float64(__obf_10eec1d7f8ad279f) - float64(__obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f))
	__obf_22dafd20a48df74e := new(big.Int).Set(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e)
	__obf_d2a65d296a5b025e := new(big.Int).Exp(__obf_cd9358bd6469dc36, big.NewInt(int64(__obf_49c750d8ab1e1e10)), nil)
	if __obf_10eec1d7f8ad279f > __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f {
		__obf_22dafd20a48df74e = __obf_22dafd20a48df74e.Quo(__obf_22dafd20a48df74e, __obf_d2a65d296a5b025e)
	} else if __obf_10eec1d7f8ad279f < __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f {
		__obf_22dafd20a48df74e = __obf_22dafd20a48df74e.Mul(__obf_22dafd20a48df74e, __obf_d2a65d296a5b025e)
	}

	return Decimal{__obf_22dafd20a48df74e: __obf_22dafd20a48df74e, __obf_10eec1d7f8ad279f: __obf_10eec1d7f8ad279f}
}

// Abs returns the absolute value of the decimal.
func (__obf_aafc5e233dfea784 Decimal) Abs() Decimal {
	if !__obf_aafc5e233dfea784.IsNegative() {
		return __obf_aafc5e233dfea784
	}
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()
	__obf_448d00243c7f1aa5 := new(big.Int).Abs(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e)
	return Decimal{__obf_22dafd20a48df74e: __obf_448d00243c7f1aa5, __obf_10eec1d7f8ad279f: __obf_aafc5e233dfea784.

		// Add returns d + d2.
		__obf_10eec1d7f8ad279f,
	}
}

func (__obf_aafc5e233dfea784 Decimal) Add(__obf_dea290dc9968f164 Decimal) Decimal {
	__obf_ad773f9fb30e9e69, __obf_39520b512e5d9ede := RescalePair(__obf_aafc5e233dfea784, __obf_dea290dc9968f164)
	__obf_fa86311d12e9b97d := new(big.Int).Add(__obf_ad773f9fb30e9e69.__obf_22dafd20a48df74e, __obf_39520b512e5d9ede.__obf_22dafd20a48df74e)
	return Decimal{__obf_22dafd20a48df74e: __obf_fa86311d12e9b97d, __obf_10eec1d7f8ad279f: __obf_ad773f9fb30e9e69.

		// Sub returns d - d2.
		__obf_10eec1d7f8ad279f,
	}
}

func (__obf_aafc5e233dfea784 Decimal) Sub(__obf_dea290dc9968f164 Decimal) Decimal {
	__obf_ad773f9fb30e9e69, __obf_39520b512e5d9ede := RescalePair(__obf_aafc5e233dfea784, __obf_dea290dc9968f164)
	__obf_fa86311d12e9b97d := new(big.Int).Sub(__obf_ad773f9fb30e9e69.__obf_22dafd20a48df74e, __obf_39520b512e5d9ede.__obf_22dafd20a48df74e)
	return Decimal{__obf_22dafd20a48df74e: __obf_fa86311d12e9b97d, __obf_10eec1d7f8ad279f: __obf_ad773f9fb30e9e69.

		// Neg returns -d.
		__obf_10eec1d7f8ad279f,
	}
}

func (__obf_aafc5e233dfea784 Decimal) Neg() Decimal {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()
	__obf_3ea3ab1a55b06161 := new(big.Int).Neg(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e)
	return Decimal{__obf_22dafd20a48df74e: __obf_3ea3ab1a55b06161, __obf_10eec1d7f8ad279f: __obf_aafc5e233dfea784.

		// Mul returns d * d2.
		__obf_10eec1d7f8ad279f,
	}
}

func (__obf_aafc5e233dfea784 Decimal) Mul(__obf_dea290dc9968f164 Decimal) Decimal {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()
	__obf_dea290dc9968f164.__obf_04511f44d644c8ae()
	__obf_360ded2578a57aec := int64(__obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f) + int64(__obf_dea290dc9968f164.__obf_10eec1d7f8ad279f)
	if __obf_360ded2578a57aec > math.MaxInt32 || __obf_360ded2578a57aec < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_360ded2578a57aec))
	}
	__obf_fa86311d12e9b97d := new(big.Int).Mul(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e, __obf_dea290dc9968f164.__obf_22dafd20a48df74e)
	return Decimal{__obf_22dafd20a48df74e: __obf_fa86311d12e9b97d, __obf_10eec1d7f8ad279f: int32(__obf_360ded2578a57aec)}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_aafc5e233dfea784 Decimal) Shift(__obf_261649088c036bf4 int32) Decimal {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()
	return Decimal{__obf_22dafd20a48df74e: new(big.Int).Set(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e), __obf_10eec1d7f8ad279f: __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f + __obf_261649088c036bf4}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_aafc5e233dfea784 Decimal) Div(__obf_dea290dc9968f164 Decimal) Decimal {
	return __obf_aafc5e233dfea784.DivRound(__obf_dea290dc9968f164, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_aafc5e233dfea784 Decimal) QuoRem(__obf_dea290dc9968f164 Decimal, __obf_6eb04b4176b00c83 int32) (Decimal, Decimal) {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()
	__obf_dea290dc9968f164.__obf_04511f44d644c8ae()
	if __obf_dea290dc9968f164.__obf_22dafd20a48df74e.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_747ea4f165e354c6 := -__obf_6eb04b4176b00c83
	__obf_e338d3262f43bafa := int64(__obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f) - int64(__obf_dea290dc9968f164.__obf_10eec1d7f8ad279f) - int64(__obf_747ea4f165e354c6)
	if __obf_e338d3262f43bafa > math.MaxInt32 || __obf_e338d3262f43bafa < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_b7f82eacfa0ae060, __obf_d7c76a47b6e59383,

		// d = a 10^ea
		// d2 = b 10^eb
		__obf_263c80375d27b0b6 big.Int
	var __obf_8e28ce2a47d25072 int32

	if __obf_e338d3262f43bafa < 0 {
		__obf_b7f82eacfa0ae060 = *__obf_aafc5e233dfea784.__obf_22dafd20a48df74e
		__obf_263c80375d27b0b6.
			SetInt64(-__obf_e338d3262f43bafa)
		__obf_d7c76a47b6e59383.
			Exp(__obf_cd9358bd6469dc36, &__obf_263c80375d27b0b6, nil)
		__obf_d7c76a47b6e59383.
			Mul(__obf_dea290dc9968f164.__obf_22dafd20a48df74e, &__obf_d7c76a47b6e59383)
		__obf_8e28ce2a47d25072 = __obf_aafc5e233dfea784.
			// now aa = a
			//     bb = b 10^(scale + eb - ea)
			__obf_10eec1d7f8ad279f

	} else {
		__obf_263c80375d27b0b6.
			SetInt64(__obf_e338d3262f43bafa)
		__obf_b7f82eacfa0ae060.
			Exp(__obf_cd9358bd6469dc36, &__obf_263c80375d27b0b6, nil)
		__obf_b7f82eacfa0ae060.
			Mul(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e, &__obf_b7f82eacfa0ae060)
		__obf_d7c76a47b6e59383 = *__obf_dea290dc9968f164.__obf_22dafd20a48df74e

		// now aa = a ^ (ea - eb - scale)
		//     bb = b
		__obf_8e28ce2a47d25072 = __obf_747ea4f165e354c6 + __obf_dea290dc9968f164.__obf_10eec1d7f8ad279f

	}
	var __obf_ce37968bd6fef53d, __obf_5039aeab743c839a big.Int
	__obf_ce37968bd6fef53d.
		QuoRem(&__obf_b7f82eacfa0ae060, &__obf_d7c76a47b6e59383, &__obf_5039aeab743c839a)
	__obf_de2f94c9804d321e := Decimal{__obf_22dafd20a48df74e: &__obf_ce37968bd6fef53d, __obf_10eec1d7f8ad279f: __obf_747ea4f165e354c6}
	__obf_224b48664002e729 := Decimal{__obf_22dafd20a48df74e: &__obf_5039aeab743c839a, __obf_10eec1d7f8ad279f: __obf_8e28ce2a47d25072}
	return __obf_de2f94c9804d321e,

		// DivRound divides and rounds to a given precision
		// i.e. to an integer multiple of 10^(-precision)
		//
		//	for a positive quotient digit 5 is rounded up, away from 0
		//	if the quotient is negative then digit 5 is rounded down, away from 0
		//
		// Note that precision<0 is allowed as input.
		__obf_224b48664002e729
}

func (__obf_aafc5e233dfea784 Decimal) DivRound(__obf_dea290dc9968f164 Decimal, __obf_6eb04b4176b00c83 int32) Decimal {
	__obf_ce37968bd6fef53d,
		// QuoRem already checks initialization
		__obf_5039aeab743c839a := __obf_aafc5e233dfea784.QuoRem(__obf_dea290dc9968f164,
		// the actual rounding decision is based on comparing r*10^precision and d2/2
		// instead compare 2 r 10 ^precision and d2
		__obf_6eb04b4176b00c83)

	var __obf_19eb1c6620305eae big.Int
	__obf_19eb1c6620305eae.
		Abs(__obf_5039aeab743c839a.__obf_22dafd20a48df74e)
	__obf_19eb1c6620305eae.
		Lsh(&__obf_19eb1c6620305eae, 1)
	__obf_0bc8b3799321eb31 := // now rv2 = abs(r.value) * 2
		Decimal{__obf_22dafd20a48df74e: &__obf_19eb1c6620305eae, __obf_10eec1d7f8ad279f: __obf_5039aeab743c839a.
			// r2 is now 2 * r * 10 ^ precision
			__obf_10eec1d7f8ad279f + __obf_6eb04b4176b00c83}

	var __obf_d520aaab80df92cd = __obf_0bc8b3799321eb31.Cmp(__obf_dea290dc9968f164.Abs())

	if __obf_d520aaab80df92cd < 0 {
		return __obf_ce37968bd6fef53d
	}

	if __obf_aafc5e233dfea784.__obf_22dafd20a48df74e.Sign()*__obf_dea290dc9968f164.__obf_22dafd20a48df74e.Sign() < 0 {
		return __obf_ce37968bd6fef53d.Sub(New(1, -__obf_6eb04b4176b00c83))
	}

	return __obf_ce37968bd6fef53d.Add(New(1, -__obf_6eb04b4176b00c83))
}

// Mod returns d % d2.
func (__obf_aafc5e233dfea784 Decimal) Mod(__obf_dea290dc9968f164 Decimal) Decimal {
	_, __obf_5039aeab743c839a := __obf_aafc5e233dfea784.QuoRem(__obf_dea290dc9968f164, 0)
	return __obf_5039aeab743c839a
}

// Pow returns d to the power of d2.
// When exponent is negative the returned decimal will have maximum precision of PowPrecisionNegativeExponent places after decimal point.
//
// Pow returns 0 (zero-value of Decimal) instead of error for power operation edge cases, to handle those edge cases use PowWithPrecision
// Edge cases not handled by Pow:
//   - 0 ** 0 => undefined value
//   - 0 ** y, where y < 0 => infinity
//   - x ** y, where x < 0 and y is non-integer decimal => imaginary value
//
// Example:
//
//	d1 := decimal.NewFromFloat(4.0)
//	d2 := decimal.NewFromFloat(4.0)
//	res1 := d1.Pow(d2)
//	res1.String() // output: "256"
//
//	d3 := decimal.NewFromFloat(5.0)
//	d4 := decimal.NewFromFloat(5.73)
//	res2 := d3.Pow(d4)
//	res2.String() // output: "10118.08037125"
func (__obf_aafc5e233dfea784 Decimal) Pow(__obf_dea290dc9968f164 Decimal) Decimal {
	__obf_f1f2031d4bf0e733 := __obf_aafc5e233dfea784.Sign()
	__obf_040b06dd84a695fd := __obf_dea290dc9968f164.Sign()

	if __obf_f1f2031d4bf0e733 == 0 {
		if __obf_040b06dd84a695fd == 0 {
			return Decimal{}
		}
		if __obf_040b06dd84a695fd == 1 {
			return Decimal{__obf_f0e1aaf9670d8421, 0}
		}
		if __obf_040b06dd84a695fd == -1 {
			return Decimal{}
		}
	}

	if __obf_040b06dd84a695fd == 0 {
		return Decimal{__obf_f136d7fff3e7d6f7, 0}
	}
	__obf_c4890ca4b5a82ab2 := // TODO: optimize extraction of fractional part
		Decimal{__obf_f136d7fff3e7d6f7, 0}
	__obf_eda303fc15e51a15, __obf_0ccb8203e910baca := __obf_dea290dc9968f164.QuoRem(__obf_c4890ca4b5a82ab2, 0)

	if __obf_f1f2031d4bf0e733 == -1 && !__obf_0ccb8203e910baca.IsZero() {
		return Decimal{}
	}
	__obf_7bcc3203cb912799, _ := __obf_aafc5e233dfea784.PowBigInt(__obf_eda303fc15e51a15.

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_22dafd20a48df74e)

	if __obf_0ccb8203e910baca.__obf_22dafd20a48df74e.Sign() == 0 {
		return __obf_7bcc3203cb912799
	}
	__obf_d993878aaae88d79 := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_aafc5e233dfea784.NumDigits()
	__obf_319b0586c2852361 := __obf_dea290dc9968f164.NumDigits()
	__obf_6eb04b4176b00c83 := __obf_d993878aaae88d79

	if __obf_319b0586c2852361 > __obf_6eb04b4176b00c83 {
		__obf_6eb04b4176b00c83 += __obf_319b0586c2852361
	}
	__obf_6eb04b4176b00c83 += 6
	__obf_7a222794f606b2f2,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_2e2f1e21b557821f := __obf_aafc5e233dfea784.Abs().Ln(-__obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f + int32(__obf_6eb04b4176b00c83))
	if __obf_2e2f1e21b557821f != nil {
		return Decimal{}
	}
	__obf_7a222794f606b2f2 = __obf_7a222794f606b2f2.Mul(__obf_0ccb8203e910baca)
	__obf_7a222794f606b2f2, __obf_2e2f1e21b557821f = __obf_7a222794f606b2f2.ExpTaylor(-__obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f + int32(__obf_6eb04b4176b00c83))
	if __obf_2e2f1e21b557821f != nil {
		return Decimal{}
	}
	__obf_2e701e573b417e0a := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_7bcc3203cb912799.Mul(__obf_7a222794f606b2f2)

	return __obf_2e701e573b417e0a
}

// PowWithPrecision returns d to the power of d2.
// Precision parameter specifies minimum precision of the result (digits after decimal point).
// Returned decimal is not rounded to 'precision' places after decimal point.
//
// PowWithPrecision returns error when:
//   - 0 ** 0 => undefined value
//   - 0 ** y, where y < 0 => infinity
//   - x ** y, where x < 0 and y is non-integer decimal => imaginary value
//
// Example:
//
//	d1 := decimal.NewFromFloat(4.0)
//	d2 := decimal.NewFromFloat(4.0)
//	res1, err := d1.PowWithPrecision(d2, 2)
//	res1.String() // output: "256"
//
//	d3 := decimal.NewFromFloat(5.0)
//	d4 := decimal.NewFromFloat(5.73)
//	res2, err := d3.PowWithPrecision(d4, 5)
//	res2.String() // output: "10118.080371595015625"
//
//	d5 := decimal.NewFromFloat(-3.0)
//	d6 := decimal.NewFromFloat(-6.0)
//	res3, err := d5.PowWithPrecision(d6, 10)
//	res3.String() // output: "0.0013717421"
func (__obf_aafc5e233dfea784 Decimal) PowWithPrecision(__obf_dea290dc9968f164 Decimal, __obf_6eb04b4176b00c83 int32) (Decimal, error) {
	__obf_f1f2031d4bf0e733 := __obf_aafc5e233dfea784.Sign()
	__obf_040b06dd84a695fd := __obf_dea290dc9968f164.Sign()

	if __obf_f1f2031d4bf0e733 == 0 {
		if __obf_040b06dd84a695fd == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_040b06dd84a695fd == 1 {
			return Decimal{__obf_f0e1aaf9670d8421, 0}, nil
		}
		if __obf_040b06dd84a695fd == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_040b06dd84a695fd == 0 {
		return Decimal{__obf_f136d7fff3e7d6f7, 0}, nil
	}
	__obf_c4890ca4b5a82ab2 := // TODO: optimize extraction of fractional part
		Decimal{__obf_f136d7fff3e7d6f7, 0}
	__obf_eda303fc15e51a15, __obf_0ccb8203e910baca := __obf_dea290dc9968f164.QuoRem(__obf_c4890ca4b5a82ab2, 0)

	if __obf_f1f2031d4bf0e733 == -1 && !__obf_0ccb8203e910baca.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}
	__obf_7bcc3203cb912799, _ := __obf_aafc5e233dfea784.__obf_c31eb46d74e4f686(__obf_eda303fc15e51a15.__obf_22dafd20a48df74e,

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_6eb04b4176b00c83)

	if __obf_0ccb8203e910baca.__obf_22dafd20a48df74e.Sign() == 0 {
		return __obf_7bcc3203cb912799, nil
	}
	__obf_d993878aaae88d79 := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_aafc5e233dfea784.NumDigits()
	__obf_319b0586c2852361 := __obf_dea290dc9968f164.NumDigits()

	if int32(__obf_d993878aaae88d79) > __obf_6eb04b4176b00c83 {
		__obf_6eb04b4176b00c83 = int32(__obf_d993878aaae88d79)
	}
	if int32(__obf_319b0586c2852361) > __obf_6eb04b4176b00c83 {
		__obf_6eb04b4176b00c83 += int32(__obf_319b0586c2852361)
	}
	__obf_6eb04b4176b00c83 += // increase precision by 10 to compensate for errors in further calculations
		10
	__obf_7a222794f606b2f2,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_2e2f1e21b557821f := __obf_aafc5e233dfea784.Abs().Ln(__obf_6eb04b4176b00c83)
	if __obf_2e2f1e21b557821f != nil {
		return Decimal{}, __obf_2e2f1e21b557821f
	}
	__obf_7a222794f606b2f2 = __obf_7a222794f606b2f2.Mul(__obf_0ccb8203e910baca)
	__obf_7a222794f606b2f2, __obf_2e2f1e21b557821f = __obf_7a222794f606b2f2.ExpTaylor(__obf_6eb04b4176b00c83)
	if __obf_2e2f1e21b557821f != nil {
		return Decimal{}, __obf_2e2f1e21b557821f
	}
	__obf_2e701e573b417e0a := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_7bcc3203cb912799.Mul(__obf_7a222794f606b2f2)

	return __obf_2e701e573b417e0a, nil
}

// PowInt32 returns d to the power of exp, where exp is int32.
// Only returns error when d and exp is 0, thus result is undefined.
//
// When exponent is negative the returned decimal will have maximum precision of PowPrecisionNegativeExponent places after decimal point.
//
// Example:
//
//	d1, err := decimal.NewFromFloat(4.0).PowInt32(4)
//	d1.String() // output: "256"
//
//	d2, err := decimal.NewFromFloat(3.13).PowInt32(5)
//	d2.String() // output: "300.4150512793"
func (__obf_aafc5e233dfea784 Decimal) PowInt32(__obf_10eec1d7f8ad279f int32) (Decimal, error) {
	if __obf_aafc5e233dfea784.IsZero() && __obf_10eec1d7f8ad279f == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_547dd75f620497be := __obf_10eec1d7f8ad279f < 0
	__obf_10eec1d7f8ad279f = __obf_1d9c7a78d7fb520b(__obf_10eec1d7f8ad279f)
	__obf_d6441e69f8e999be, __obf_004ab4ca1d179e81 := __obf_aafc5e233dfea784, New(1, 0)

	for __obf_10eec1d7f8ad279f > 0 {
		if __obf_10eec1d7f8ad279f%2 == 1 {
			__obf_004ab4ca1d179e81 = __obf_004ab4ca1d179e81.Mul(__obf_d6441e69f8e999be)
		}
		__obf_10eec1d7f8ad279f /= 2

		if __obf_10eec1d7f8ad279f > 0 {
			__obf_d6441e69f8e999be = __obf_d6441e69f8e999be.Mul(__obf_d6441e69f8e999be)
		}
	}

	if __obf_547dd75f620497be {
		return New(1, 0).DivRound(__obf_004ab4ca1d179e81, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_004ab4ca1d179e81, nil
}

// PowBigInt returns d to the power of exp, where exp is big.Int.
// Only returns error when d and exp is 0, thus result is undefined.
//
// When exponent is negative the returned decimal will have maximum precision of PowPrecisionNegativeExponent places after decimal point.
//
// Example:
//
//	d1, err := decimal.NewFromFloat(3.0).PowBigInt(big.NewInt(3))
//	d1.String() // output: "27"
//
//	d2, err := decimal.NewFromFloat(629.25).PowBigInt(big.NewInt(5))
//	d2.String() // output: "98654323103449.5673828125"
func (__obf_aafc5e233dfea784 Decimal) PowBigInt(__obf_10eec1d7f8ad279f *big.Int) (Decimal, error) {
	return __obf_aafc5e233dfea784.__obf_c31eb46d74e4f686(__obf_10eec1d7f8ad279f, int32(PowPrecisionNegativeExponent))
}

func (__obf_aafc5e233dfea784 Decimal) __obf_c31eb46d74e4f686(__obf_10eec1d7f8ad279f *big.Int, __obf_6eb04b4176b00c83 int32) (Decimal, error) {
	if __obf_aafc5e233dfea784.IsZero() && __obf_10eec1d7f8ad279f.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_4de7aabe6135af48 := new(big.Int).Set(__obf_10eec1d7f8ad279f)
	__obf_547dd75f620497be := __obf_10eec1d7f8ad279f.Sign() < 0

	if __obf_547dd75f620497be {
		__obf_4de7aabe6135af48.
			Abs(__obf_4de7aabe6135af48)
	}
	__obf_d6441e69f8e999be, __obf_004ab4ca1d179e81 := __obf_aafc5e233dfea784, New(1, 0)

	for __obf_4de7aabe6135af48.Sign() > 0 {
		if __obf_4de7aabe6135af48.Bit(0) == 1 {
			__obf_004ab4ca1d179e81 = __obf_004ab4ca1d179e81.Mul(__obf_d6441e69f8e999be)
		}
		__obf_4de7aabe6135af48.
			Rsh(__obf_4de7aabe6135af48, 1)

		if __obf_4de7aabe6135af48.Sign() > 0 {
			__obf_d6441e69f8e999be = __obf_d6441e69f8e999be.Mul(__obf_d6441e69f8e999be)
		}
	}

	if __obf_547dd75f620497be {
		return New(1, 0).DivRound(__obf_004ab4ca1d179e81, __obf_6eb04b4176b00c83), nil
	}

	return __obf_004ab4ca1d179e81, nil
}

// ExpHullAbrham calculates the natural exponent of decimal (e to the power of d) using Hull-Abraham algorithm.
// OverallPrecision argument specifies the overall precision of the result (integer part + decimal part).
//
// ExpHullAbrham is faster than ExpTaylor for small precision values, but it is much slower for large precision values.
//
// Example:
//
//	NewFromFloat(26.1).ExpHullAbrham(2).String()    // output: "220000000000"
//	NewFromFloat(26.1).ExpHullAbrham(20).String()   // output: "216314672147.05767284"
func (__obf_aafc5e233dfea784 Decimal) ExpHullAbrham(__obf_5568bfff1b98f6c0 uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_aafc5e233dfea784.IsZero() {
		return Decimal{__obf_f136d7fff3e7d6f7, 0}, nil
	}
	__obf_c017efb5d66d56c6 := __obf_5568bfff1b98f6c0

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_605681be6110cd23 := __obf_aafc5e233dfea784.Abs().InexactFloat64()
	if __obf_c85a7837cd766c82 := __obf_605681be6110cd23 / 23; __obf_c85a7837cd766c82 > float64(__obf_c017efb5d66d56c6) && __obf_c85a7837cd766c82 < float64(ExpMaxIterations) {
		__obf_c017efb5d66d56c6 = uint32(math.Ceil(__obf_c85a7837cd766c82))
	}
	__obf_94e7255a52edc483 := // fail if abs(d) beyond an over/underflow threshold
		New(23*int64(__obf_c017efb5d66d56c6), 0)
	if __obf_aafc5e233dfea784.Abs().Cmp(__obf_94e7255a52edc483) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}
	__obf_19687bd84099fd23 := // Return 1 if abs(d) small enough; this also avoids later over/underflow
		New(9, -int32(__obf_c017efb5d66d56c6)-1)
	if __obf_aafc5e233dfea784.Abs().Cmp(__obf_19687bd84099fd23) <= 0 {
		return Decimal{__obf_f136d7fff3e7d6f7, __obf_aafc5e233dfea784.

			// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
			__obf_10eec1d7f8ad279f}, nil
	}
	__obf_a2e0258e0cd11e39 := __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f + int32(__obf_aafc5e233dfea784.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_a2e0258e0cd11e39 < 0 {
		__obf_a2e0258e0cd11e39 = 0
	}
	__obf_8f1bcc26019920eb := New(1, __obf_a2e0258e0cd11e39)
	__obf_5039aeab743c839a := // reduction factor
		Decimal{new(big.Int).Set(__obf_aafc5e233dfea784. // reduced argument
									__obf_22dafd20a48df74e), __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f - __obf_a2e0258e0cd11e39}
	__obf_bd96a373e94ba477 := int32(__obf_c017efb5d66d56c6) + __obf_a2e0258e0cd11e39 + 2
	__obf_a8afd930744fed47 := // precision for calculating the sum

		// Determine n, the number of therms for calculating sum
		// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
		// for solving appropriate equation, along with directed
		// roundings and simple rational bound for log10(p/abs(r))
		__obf_5039aeab743c839a.Abs().InexactFloat64()
	__obf_0f3fd9bf1d1bb82d := float64(__obf_bd96a373e94ba477)
	__obf_0e5cd1434460b1f1 := math.Ceil((1.453*__obf_0f3fd9bf1d1bb82d - 1.182) / math.Log10(__obf_0f3fd9bf1d1bb82d/__obf_a8afd930744fed47))
	if __obf_0e5cd1434460b1f1 > float64(ExpMaxIterations) || math.IsNaN(__obf_0e5cd1434460b1f1) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_d6441e69f8e999be := int64(__obf_0e5cd1434460b1f1)
	__obf_6893e4bf8945b770 := New(0, 0)
	__obf_6cebf47a1809ac55 := New(1, 0)
	__obf_c4890ca4b5a82ab2 := New(1, 0)
	for __obf_03985402d384eb80 := __obf_d6441e69f8e999be - 1; __obf_03985402d384eb80 > 0; __obf_03985402d384eb80-- {
		__obf_6893e4bf8945b770.__obf_22dafd20a48df74e.
			SetInt64(__obf_03985402d384eb80)
		__obf_6cebf47a1809ac55 = __obf_6cebf47a1809ac55.Mul(__obf_5039aeab743c839a.DivRound(__obf_6893e4bf8945b770, __obf_bd96a373e94ba477))
		__obf_6cebf47a1809ac55 = __obf_6cebf47a1809ac55.Add(__obf_c4890ca4b5a82ab2)
	}
	__obf_b702b3fb6ad488bb := __obf_8f1bcc26019920eb.IntPart()
	__obf_2e701e573b417e0a := New(1, 0)
	for __obf_03985402d384eb80 := __obf_b702b3fb6ad488bb; __obf_03985402d384eb80 > 0; __obf_03985402d384eb80-- {
		__obf_2e701e573b417e0a = __obf_2e701e573b417e0a.Mul(__obf_6cebf47a1809ac55)
	}
	__obf_b9d8d94609fbf2a3 := int32(__obf_2e701e573b417e0a.NumDigits())

	var __obf_34b7878916442502 int32
	if __obf_b9d8d94609fbf2a3 > __obf_1d9c7a78d7fb520b(__obf_2e701e573b417e0a.__obf_10eec1d7f8ad279f) {
		__obf_34b7878916442502 = int32(__obf_c017efb5d66d56c6) - __obf_b9d8d94609fbf2a3 - __obf_2e701e573b417e0a.__obf_10eec1d7f8ad279f
	} else {
		__obf_34b7878916442502 = int32(__obf_c017efb5d66d56c6)
	}
	__obf_2e701e573b417e0a = __obf_2e701e573b417e0a.Round(__obf_34b7878916442502)

	return __obf_2e701e573b417e0a, nil
}

// ExpTaylor calculates the natural exponent of decimal (e to the power of d) using Taylor series expansion.
// Precision argument specifies how precise the result must be (number of digits after decimal point).
// Negative precision is allowed.
//
// ExpTaylor is much faster for large precision values than ExpHullAbrham.
//
// Example:
//
//	d, err := NewFromFloat(26.1).ExpTaylor(2).String()
//	d.String()  // output: "216314672147.06"
//
//	NewFromFloat(26.1).ExpTaylor(20).String()
//	d.String()  // output: "216314672147.05767284062928674083"
//
//	NewFromFloat(26.1).ExpTaylor(-10).String()
//	d.String()  // output: "220000000000"
func (__obf_aafc5e233dfea784 Decimal) ExpTaylor(__obf_6eb04b4176b00c83 int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_aafc5e233dfea784.IsZero() {
		return Decimal{__obf_f136d7fff3e7d6f7, 0}.Round(__obf_6eb04b4176b00c83), nil
	}

	var __obf_6c0667e52134e183 Decimal
	var __obf_33fb5ddee6560b28 int32
	if __obf_6eb04b4176b00c83 < 0 {
		__obf_6c0667e52134e183 = New(1, -1)
		__obf_33fb5ddee6560b28 = 8
	} else {
		__obf_6c0667e52134e183 = New(1, -__obf_6eb04b4176b00c83-1)
		__obf_33fb5ddee6560b28 = __obf_6eb04b4176b00c83 + 1
	}
	__obf_1f486aff8c848bcf := __obf_aafc5e233dfea784.Abs()
	__obf_28220357deacf118 := __obf_aafc5e233dfea784.Abs()
	__obf_aac878564a4c20df := New(1, 0)
	__obf_004ab4ca1d179e81 := New(1, 0)

	for __obf_03985402d384eb80 := int64(1); ; {
		__obf_d94550437ab03d54 := __obf_28220357deacf118.DivRound(__obf_aac878564a4c20df, __obf_33fb5ddee6560b28)
		__obf_004ab4ca1d179e81 = __obf_004ab4ca1d179e81.Add(__obf_d94550437ab03d54)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_d94550437ab03d54.Cmp(__obf_6c0667e52134e183) < 0 {
			break
		}
		__obf_28220357deacf118 = __obf_28220357deacf118.Mul(__obf_1f486aff8c848bcf)
		__obf_03985402d384eb80++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_2b7a96df87ff7b2b) >= int(__obf_03985402d384eb80) && !__obf_2b7a96df87ff7b2b[__obf_03985402d384eb80-1].IsZero() {
			__obf_aac878564a4c20df = __obf_2b7a96df87ff7b2b[__obf_03985402d384eb80-1]
		} else {
			__obf_aac878564a4c20df = // To avoid any race conditions, firstly the zero value is appended to a slice to create
				// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
				// factorial using the index notation.
				__obf_2b7a96df87ff7b2b[__obf_03985402d384eb80-2].Mul(New(__obf_03985402d384eb80, 0))
			__obf_2b7a96df87ff7b2b = append(__obf_2b7a96df87ff7b2b, Zero)
			__obf_2b7a96df87ff7b2b[__obf_03985402d384eb80-1] = __obf_aac878564a4c20df
		}
	}

	if __obf_aafc5e233dfea784.Sign() < 0 {
		__obf_004ab4ca1d179e81 = New(1, 0).DivRound(__obf_004ab4ca1d179e81, __obf_6eb04b4176b00c83+1)
	}
	__obf_004ab4ca1d179e81 = __obf_004ab4ca1d179e81.Round(__obf_6eb04b4176b00c83)
	return __obf_004ab4ca1d179e81, nil
}

// Ln calculates natural logarithm of d.
// Precision argument specifies how precise the result must be (number of digits after decimal point).
// Negative precision is allowed.
//
// Example:
//
//	d1, err := NewFromFloat(13.3).Ln(2)
//	d1.String()  // output: "2.59"
//
//	d2, err := NewFromFloat(579.161).Ln(10)
//	d2.String()  // output: "6.3615805046"
func (__obf_aafc5e233dfea784 Decimal) Ln(__obf_6eb04b4176b00c83 int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_aafc5e233dfea784.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_aafc5e233dfea784.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}
	__obf_7bb9c12451674b54 := __obf_6eb04b4176b00c83 + 2
	__obf_6f45f5fffe86c965 := __obf_aafc5e233dfea784.Copy()

	var __obf_d2663d096f3420f4, __obf_e1e07161f161b085, __obf_765128adc8370199, __obf_4ebb53d79228e8bd, __obf_11c3dff308c0a52d Decimal
	__obf_d2663d096f3420f4 = __obf_6f45f5fffe86c965.Sub(Decimal{__obf_f136d7fff3e7d6f7, 0})
	__obf_e1e07161f161b085 = Decimal{__obf_f136d7fff3e7d6f7, -1}
	__obf_8928c6666334073e := // for decimal in range [0.9, 1.1] where ln(d) is close to 0
		false

	if __obf_d2663d096f3420f4.Abs().Cmp(__obf_e1e07161f161b085) <= 0 {
		__obf_8928c6666334073e = true
	} else {
		__obf_764e302576902489 := // reduce input decimal to range [0.1, 1)
			int32(__obf_6f45f5fffe86c965.NumDigits()) + __obf_6f45f5fffe86c965.__obf_10eec1d7f8ad279f

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_6f45f5fffe86c965.__obf_10eec1d7f8ad279f -= __obf_764e302576902489
		__obf_0eed81a9c3bb0153 := __obf_0eed81a9c3bb0153.__obf_bac49944c6644e36(__obf_7bb9c12451674b54)
		__obf_11c3dff308c0a52d = NewFromInt32(__obf_764e302576902489)
		__obf_11c3dff308c0a52d = __obf_11c3dff308c0a52d.Mul(__obf_0eed81a9c3bb0153)
		__obf_d2663d096f3420f4 = __obf_6f45f5fffe86c965.Sub(Decimal{__obf_f136d7fff3e7d6f7, 0})

		if __obf_d2663d096f3420f4.Abs().Cmp(__obf_e1e07161f161b085) <= 0 {
			__obf_8928c6666334073e = true
		} else {
			__obf_bdde578ec8e56dbc := // initial estimate using floats
				__obf_6f45f5fffe86c965.InexactFloat64()
			__obf_d2663d096f3420f4 = NewFromFloat(math.Log(__obf_bdde578ec8e56dbc))
		}
	}
	__obf_6c0667e52134e183 := Decimal{__obf_f136d7fff3e7d6f7, -__obf_7bb9c12451674b54}

	if __obf_8928c6666334073e {
		__obf_765128adc8370199 = // Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
			// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			// until the difference between current and next term is smaller than epsilon.
			// Coverage quite fast for decimals close to 1.0

			// z + 2
			__obf_d2663d096f3420f4.Add(Decimal{__obf_3f8f48b9238d224e, 0})
		__obf_e1e07161f161b085 = // z / (z + 2)
			__obf_d2663d096f3420f4.DivRound(__obf_765128adc8370199, __obf_7bb9c12451674b54)
		__obf_d2663d096f3420f4 = // 2 * (z / (z + 2))
			__obf_e1e07161f161b085.Add(__obf_e1e07161f161b085)
		__obf_765128adc8370199 = __obf_d2663d096f3420f4.Copy()

		for __obf_d6441e69f8e999be := 1; ; __obf_d6441e69f8e999be++ {
			__obf_765128adc8370199 = // 2 * (z / (z+2))^(2n+1)
				__obf_765128adc8370199.Mul(__obf_e1e07161f161b085).Mul(__obf_e1e07161f161b085)
			__obf_4ebb53d79228e8bd = // 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
				NewFromInt(int64(2*__obf_d6441e69f8e999be + 1))
			__obf_4ebb53d79228e8bd = __obf_765128adc8370199.DivRound(__obf_4ebb53d79228e8bd, __obf_7bb9c12451674b54)
			__obf_d2663d096f3420f4 = // comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
				__obf_d2663d096f3420f4.Add(__obf_4ebb53d79228e8bd)

			if __obf_4ebb53d79228e8bd.Abs().Cmp(__obf_6c0667e52134e183) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_ec799caa43d0e017 Decimal
		__obf_09588f35e01a70b1 := __obf_7bb9c12451674b54*2 + 10

		for __obf_03985402d384eb80 := int32(0); __obf_03985402d384eb80 < __obf_09588f35e01a70b1;
		// exp(a_n)
		__obf_03985402d384eb80++ {
			__obf_e1e07161f161b085, _ = __obf_d2663d096f3420f4.ExpTaylor(__obf_7bb9c12451674b54)
			__obf_765128adc8370199 = // exp(a_n) - z
				__obf_e1e07161f161b085.Sub(__obf_6f45f5fffe86c965)
			__obf_765128adc8370199 = // 2 * (exp(a_n) - z)
				__obf_765128adc8370199.Add(__obf_765128adc8370199)
			__obf_4ebb53d79228e8bd = // exp(a_n) + z
				__obf_e1e07161f161b085.Add(__obf_6f45f5fffe86c965)
			__obf_e1e07161f161b085 = // 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_765128adc8370199.DivRound(__obf_4ebb53d79228e8bd, __obf_7bb9c12451674b54)
			__obf_d2663d096f3420f4 = // comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_d2663d096f3420f4.Sub(__obf_e1e07161f161b085)

			if __obf_ec799caa43d0e017.Add(__obf_e1e07161f161b085).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_e1e07161f161b085.Abs().Cmp(__obf_6c0667e52134e183) <= 0 {
				break
			}
			__obf_ec799caa43d0e017 = __obf_e1e07161f161b085
		}
	}
	__obf_d2663d096f3420f4 = __obf_d2663d096f3420f4.Add(__obf_11c3dff308c0a52d)

	return __obf_d2663d096f3420f4.Round(__obf_6eb04b4176b00c83), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_aafc5e233dfea784 Decimal) NumDigits() int {
	if __obf_aafc5e233dfea784.__obf_22dafd20a48df74e == nil {
		return 1
	}

	if __obf_aafc5e233dfea784.__obf_22dafd20a48df74e.IsInt64() {
		__obf_7b0ba8356755aea8 := __obf_aafc5e233dfea784.
			// restrict fast path to integers with exact conversion to float64
			__obf_22dafd20a48df74e.Int64()

		if __obf_7b0ba8356755aea8 <= (1<<53) && __obf_7b0ba8356755aea8 >= -(1<<53) {
			if __obf_7b0ba8356755aea8 == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_7b0ba8356755aea8)))) + 1
		}
	}
	__obf_2448cd922194bed7 := int(float64(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e.BitLen()) / math.Log2(10))
	__obf_a471bcaa3ce60cee := // estimatedNumDigits (lg10) may be off by 1, need to verify
		big.NewInt(int64(__obf_2448cd922194bed7))
	__obf_90e8f4247ad47955 := __obf_a471bcaa3ce60cee.Exp(__obf_cd9358bd6469dc36, __obf_a471bcaa3ce60cee, nil)

	if __obf_aafc5e233dfea784.__obf_22dafd20a48df74e.CmpAbs(__obf_90e8f4247ad47955) >= 0 {
		return __obf_2448cd922194bed7 + 1
	}

	return __obf_2448cd922194bed7
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_aafc5e233dfea784 Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_5039aeab743c839a big.Int
	__obf_ce37968bd6fef53d := new(big.Int).Set(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e)
	for __obf_6f45f5fffe86c965 := __obf_1d9c7a78d7fb520b(__obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f); __obf_6f45f5fffe86c965 > 0; __obf_6f45f5fffe86c965-- {
		__obf_ce37968bd6fef53d.
			QuoRem(__obf_ce37968bd6fef53d, __obf_cd9358bd6469dc36, &__obf_5039aeab743c839a)
		if __obf_5039aeab743c839a.Cmp(__obf_f0e1aaf9670d8421) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_1d9c7a78d7fb520b(__obf_d6441e69f8e999be int32) int32 {
	if __obf_d6441e69f8e999be < 0 {
		return -__obf_d6441e69f8e999be
	}
	return __obf_d6441e69f8e999be
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_aafc5e233dfea784 Decimal) Cmp(__obf_dea290dc9968f164 Decimal) int {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()
	__obf_dea290dc9968f164.__obf_04511f44d644c8ae()

	if __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f == __obf_dea290dc9968f164.__obf_10eec1d7f8ad279f {
		return __obf_aafc5e233dfea784.__obf_22dafd20a48df74e.Cmp(__obf_dea290dc9968f164.__obf_22dafd20a48df74e)
	}
	__obf_ad773f9fb30e9e69, __obf_39520b512e5d9ede := RescalePair(__obf_aafc5e233dfea784, __obf_dea290dc9968f164)

	return __obf_ad773f9fb30e9e69.__obf_22dafd20a48df74e.Cmp(__obf_39520b512e5d9ede.

		// Compare compares the numbers represented by d and d2 and returns:
		//
		//	-1 if d <  d2
		//	 0 if d == d2
		//	+1 if d >  d2
		__obf_22dafd20a48df74e)
}

func (__obf_aafc5e233dfea784 Decimal) Compare(__obf_dea290dc9968f164 Decimal) int {
	return __obf_aafc5e233dfea784.Cmp(__obf_dea290dc9968f164)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_aafc5e233dfea784 Decimal) Equal(__obf_dea290dc9968f164 Decimal) bool {
	return __obf_aafc5e233dfea784.Cmp(__obf_dea290dc9968f164) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_aafc5e233dfea784 Decimal) Equals(__obf_dea290dc9968f164 Decimal) bool {
	return __obf_aafc5e233dfea784.Equal(__obf_dea290dc9968f164)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_aafc5e233dfea784 Decimal) GreaterThan(__obf_dea290dc9968f164 Decimal) bool {
	return __obf_aafc5e233dfea784.Cmp(__obf_dea290dc9968f164) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_aafc5e233dfea784 Decimal) GreaterThanOrEqual(__obf_dea290dc9968f164 Decimal) bool {
	__obf_38bfef7747f85827 := __obf_aafc5e233dfea784.Cmp(__obf_dea290dc9968f164)
	return __obf_38bfef7747f85827 == 1 || __obf_38bfef7747f85827 == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_aafc5e233dfea784 Decimal) LessThan(__obf_dea290dc9968f164 Decimal) bool {
	return __obf_aafc5e233dfea784.Cmp(__obf_dea290dc9968f164) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_aafc5e233dfea784 Decimal) LessThanOrEqual(__obf_dea290dc9968f164 Decimal) bool {
	__obf_38bfef7747f85827 := __obf_aafc5e233dfea784.Cmp(__obf_dea290dc9968f164)
	return __obf_38bfef7747f85827 == -1 || __obf_38bfef7747f85827 == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_aafc5e233dfea784 Decimal) Sign() int {
	if __obf_aafc5e233dfea784.__obf_22dafd20a48df74e == nil {
		return 0
	}
	return __obf_aafc5e233dfea784.

		// IsPositive return
		//
		//	true if d > 0
		//	false if d == 0
		//	false if d < 0
		__obf_22dafd20a48df74e.Sign()
}

func (__obf_aafc5e233dfea784 Decimal) IsPositive() bool {
	return __obf_aafc5e233dfea784.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_aafc5e233dfea784 Decimal) IsNegative() bool {
	return __obf_aafc5e233dfea784.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_aafc5e233dfea784 Decimal) IsZero() bool {
	return __obf_aafc5e233dfea784.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_aafc5e233dfea784 Decimal) Exponent() int32 {
	return __obf_aafc5e233dfea784.

		// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
		__obf_10eec1d7f8ad279f
}

func (__obf_aafc5e233dfea784 Decimal) Coefficient() *big.Int {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_aafc5e233dfea784.

		// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
		// If coefficient cannot be represented in an int64, the result will be undefined.
		__obf_22dafd20a48df74e)
}

func (__obf_aafc5e233dfea784 Decimal) CoefficientInt64() int64 {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()
	return __obf_aafc5e233dfea784.

		// IntPart returns the integer component of the decimal.
		__obf_22dafd20a48df74e.Int64()
}

func (__obf_aafc5e233dfea784 Decimal) IntPart() int64 {
	__obf_b9d8762a0ca17f81 := __obf_aafc5e233dfea784.__obf_51b05428a428dd2d(0)
	return __obf_b9d8762a0ca17f81.__obf_22dafd20a48df74e.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_aafc5e233dfea784 Decimal) BigInt() *big.Int {
	__obf_b9d8762a0ca17f81 := __obf_aafc5e233dfea784.__obf_51b05428a428dd2d(0)
	return __obf_b9d8762a0ca17f81.

		// BigFloat returns decimal as BigFloat.
		// Be aware that casting decimal to BigFloat might cause a loss of precision.
		__obf_22dafd20a48df74e
}

func (__obf_aafc5e233dfea784 Decimal) BigFloat() *big.Float {
	__obf_605681be6110cd23 := &big.Float{}
	__obf_605681be6110cd23.
		SetString(__obf_aafc5e233dfea784.String())
	return __obf_605681be6110cd23
}

// Rat returns a rational number representation of the decimal.
func (__obf_aafc5e233dfea784 Decimal) Rat() *big.Rat {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()
	if __obf_aafc5e233dfea784.
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_10eec1d7f8ad279f <= 0 {
		__obf_eea4e16c91456535 := new(big.Int).Exp(__obf_cd9358bd6469dc36, big.NewInt(-int64(__obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f)), nil)
		return new(big.Rat).SetFrac(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e, __obf_eea4e16c91456535)
	}
	__obf_285796e4f42a2a11 := new(big.Int).Exp(__obf_cd9358bd6469dc36, big.NewInt(int64(__obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f)), nil)
	__obf_22156b7400d3ade3 := new(big.Int).Mul(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e, __obf_285796e4f42a2a11)
	return new(big.Rat).SetFrac(__obf_22156b7400d3ade3,

		// Float64 returns the nearest float64 value for d and a bool indicating
		// whether f represents d exactly.
		// For more details, see the documentation for big.Rat.Float64
		__obf_f136d7fff3e7d6f7)
}

func (__obf_aafc5e233dfea784 Decimal) Float64() (__obf_605681be6110cd23 float64, __obf_bc58cb6c1ce1b650 bool) {
	return __obf_aafc5e233dfea784.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_aafc5e233dfea784 Decimal) InexactFloat64() float64 {
	__obf_605681be6110cd23, _ := __obf_aafc5e233dfea784.Float64()
	return __obf_605681be6110cd23
}

// String returns the string representation of the decimal
// with the fixed point.
//
// Example:
//
//	d := New(-12345, -3)
//	println(d.String())
//
// Output:
//
//	-12.345
func (__obf_aafc5e233dfea784 Decimal) String() string {
	return __obf_aafc5e233dfea784.string(true)
}

// StringFixed returns a rounded fixed-point string with places digits after
// the decimal point.
//
// Example:
//
//	NewFromFloat(0).StringFixed(2) // output: "0.00"
//	NewFromFloat(0).StringFixed(0) // output: "0"
//	NewFromFloat(5.45).StringFixed(0) // output: "5"
//	NewFromFloat(5.45).StringFixed(1) // output: "5.5"
//	NewFromFloat(5.45).StringFixed(2) // output: "5.45"
//	NewFromFloat(5.45).StringFixed(3) // output: "5.450"
//	NewFromFloat(545).StringFixed(-1) // output: "550"
func (__obf_aafc5e233dfea784 Decimal) StringFixed(__obf_1259b5c4caf8d34d int32) string {
	__obf_a07a49b3566eb27b := __obf_aafc5e233dfea784.Round(__obf_1259b5c4caf8d34d)
	return __obf_a07a49b3566eb27b.string(false)
}

// StringFixedBank returns a banker rounded fixed-point string with places digits
// after the decimal point.
//
// Example:
//
//	NewFromFloat(0).StringFixedBank(2) // output: "0.00"
//	NewFromFloat(0).StringFixedBank(0) // output: "0"
//	NewFromFloat(5.45).StringFixedBank(0) // output: "5"
//	NewFromFloat(5.45).StringFixedBank(1) // output: "5.4"
//	NewFromFloat(5.45).StringFixedBank(2) // output: "5.45"
//	NewFromFloat(5.45).StringFixedBank(3) // output: "5.450"
//	NewFromFloat(545).StringFixedBank(-1) // output: "540"
func (__obf_aafc5e233dfea784 Decimal) StringFixedBank(__obf_1259b5c4caf8d34d int32) string {
	__obf_a07a49b3566eb27b := __obf_aafc5e233dfea784.RoundBank(__obf_1259b5c4caf8d34d)
	return __obf_a07a49b3566eb27b.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_aafc5e233dfea784 Decimal) StringFixedCash(__obf_f5b39b07684a644b uint8) string {
	__obf_a07a49b3566eb27b := __obf_aafc5e233dfea784.RoundCash(__obf_f5b39b07684a644b)
	return __obf_a07a49b3566eb27b.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_aafc5e233dfea784 Decimal) Round(__obf_1259b5c4caf8d34d int32) Decimal {
	if __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f == -__obf_1259b5c4caf8d34d {
		return __obf_aafc5e233dfea784
	}
	__obf_561feaee4f1013f9 := // truncate to places + 1
		__obf_aafc5e233dfea784.__obf_51b05428a428dd2d(-__obf_1259b5c4caf8d34d - 1)

	// add sign(d) * 0.5
	if __obf_561feaee4f1013f9.__obf_22dafd20a48df74e.Sign() < 0 {
		__obf_561feaee4f1013f9.__obf_22dafd20a48df74e.
			Sub(__obf_561feaee4f1013f9.__obf_22dafd20a48df74e, __obf_3f87d09b59f6435e)
	} else {
		__obf_561feaee4f1013f9.__obf_22dafd20a48df74e.
			Add(__obf_561feaee4f1013f9.__obf_22dafd20a48df74e,

				// floor for positive numbers, ceil for negative numbers
				__obf_3f87d09b59f6435e)
	}

	_, __obf_ade33d4daf3b3fc1 := __obf_561feaee4f1013f9.__obf_22dafd20a48df74e.DivMod(__obf_561feaee4f1013f9.__obf_22dafd20a48df74e, __obf_cd9358bd6469dc36, new(big.Int))
	__obf_561feaee4f1013f9.__obf_10eec1d7f8ad279f++
	if __obf_561feaee4f1013f9.__obf_22dafd20a48df74e.Sign() < 0 && __obf_ade33d4daf3b3fc1.Cmp(__obf_f0e1aaf9670d8421) != 0 {
		__obf_561feaee4f1013f9.__obf_22dafd20a48df74e.
			Add(__obf_561feaee4f1013f9.__obf_22dafd20a48df74e,

				// RoundCeil rounds the decimal towards +infinity.
				//
				// Example:
				//
				//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
				//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
				//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
				//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
				__obf_f136d7fff3e7d6f7)
	}

	return __obf_561feaee4f1013f9
}

func (__obf_aafc5e233dfea784 Decimal) RoundCeil(__obf_1259b5c4caf8d34d int32) Decimal {
	if __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f >= -__obf_1259b5c4caf8d34d {
		return __obf_aafc5e233dfea784
	}
	__obf_9400c6e971841675 := __obf_aafc5e233dfea784.__obf_51b05428a428dd2d(-__obf_1259b5c4caf8d34d)
	if __obf_aafc5e233dfea784.Equal(__obf_9400c6e971841675) {
		return __obf_aafc5e233dfea784
	}

	if __obf_aafc5e233dfea784.__obf_22dafd20a48df74e.Sign() > 0 {
		__obf_9400c6e971841675.__obf_22dafd20a48df74e.
			Add(__obf_9400c6e971841675.__obf_22dafd20a48df74e, __obf_f136d7fff3e7d6f7)
	}

	return __obf_9400c6e971841675
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_aafc5e233dfea784 Decimal) RoundFloor(__obf_1259b5c4caf8d34d int32) Decimal {
	if __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f >= -__obf_1259b5c4caf8d34d {
		return __obf_aafc5e233dfea784
	}
	__obf_9400c6e971841675 := __obf_aafc5e233dfea784.__obf_51b05428a428dd2d(-__obf_1259b5c4caf8d34d)
	if __obf_aafc5e233dfea784.Equal(__obf_9400c6e971841675) {
		return __obf_aafc5e233dfea784
	}

	if __obf_aafc5e233dfea784.__obf_22dafd20a48df74e.Sign() < 0 {
		__obf_9400c6e971841675.__obf_22dafd20a48df74e.
			Sub(__obf_9400c6e971841675.__obf_22dafd20a48df74e, __obf_f136d7fff3e7d6f7)
	}

	return __obf_9400c6e971841675
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_aafc5e233dfea784 Decimal) RoundUp(__obf_1259b5c4caf8d34d int32) Decimal {
	if __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f >= -__obf_1259b5c4caf8d34d {
		return __obf_aafc5e233dfea784
	}
	__obf_9400c6e971841675 := __obf_aafc5e233dfea784.__obf_51b05428a428dd2d(-__obf_1259b5c4caf8d34d)
	if __obf_aafc5e233dfea784.Equal(__obf_9400c6e971841675) {
		return __obf_aafc5e233dfea784
	}

	if __obf_aafc5e233dfea784.__obf_22dafd20a48df74e.Sign() > 0 {
		__obf_9400c6e971841675.__obf_22dafd20a48df74e.
			Add(__obf_9400c6e971841675.__obf_22dafd20a48df74e, __obf_f136d7fff3e7d6f7)
	} else if __obf_aafc5e233dfea784.__obf_22dafd20a48df74e.Sign() < 0 {
		__obf_9400c6e971841675.__obf_22dafd20a48df74e.
			Sub(__obf_9400c6e971841675.__obf_22dafd20a48df74e, __obf_f136d7fff3e7d6f7)
	}

	return __obf_9400c6e971841675
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_aafc5e233dfea784 Decimal) RoundDown(__obf_1259b5c4caf8d34d int32) Decimal {
	if __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f >= -__obf_1259b5c4caf8d34d {
		return __obf_aafc5e233dfea784
	}
	__obf_9400c6e971841675 := __obf_aafc5e233dfea784.__obf_51b05428a428dd2d(-__obf_1259b5c4caf8d34d)
	if __obf_aafc5e233dfea784.Equal(__obf_9400c6e971841675) {
		return __obf_aafc5e233dfea784
	}
	return __obf_9400c6e971841675
}

// RoundBank rounds the decimal to places decimal places.
// If the final digit to round is equidistant from the nearest two integers the
// rounded value is taken as the even number
//
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Examples:
//
//	NewFromFloat(5.45).RoundBank(1).String() // output: "5.4"
//	NewFromFloat(545).RoundBank(-1).String() // output: "540"
//	NewFromFloat(5.46).RoundBank(1).String() // output: "5.5"
//	NewFromFloat(546).RoundBank(-1).String() // output: "550"
//	NewFromFloat(5.55).RoundBank(1).String() // output: "5.6"
//	NewFromFloat(555).RoundBank(-1).String() // output: "560"
func (__obf_aafc5e233dfea784 Decimal) RoundBank(__obf_1259b5c4caf8d34d int32) Decimal {
	__obf_b046628ba347e8d2 := __obf_aafc5e233dfea784.Round(__obf_1259b5c4caf8d34d)
	__obf_72949ebdd739e997 := __obf_aafc5e233dfea784.Sub(__obf_b046628ba347e8d2).Abs()
	__obf_03cda501e2ed87dc := New(5, -__obf_1259b5c4caf8d34d-1)
	if __obf_72949ebdd739e997.Cmp(__obf_03cda501e2ed87dc) == 0 && __obf_b046628ba347e8d2.__obf_22dafd20a48df74e.Bit(0) != 0 {
		if __obf_b046628ba347e8d2.__obf_22dafd20a48df74e.Sign() < 0 {
			__obf_b046628ba347e8d2.__obf_22dafd20a48df74e.
				Add(__obf_b046628ba347e8d2.__obf_22dafd20a48df74e, __obf_f136d7fff3e7d6f7)
		} else {
			__obf_b046628ba347e8d2.__obf_22dafd20a48df74e.
				Sub(__obf_b046628ba347e8d2.__obf_22dafd20a48df74e, __obf_f136d7fff3e7d6f7)
		}
	}

	return __obf_b046628ba347e8d2
}

// RoundCash aka Cash/Penny/öre rounding rounds decimal to a specific
// interval. The amount payable for a cash transaction is rounded to the nearest
// multiple of the minimum currency unit available. The following intervals are
// available: 5, 10, 25, 50 and 100; any other number throws a panic.
//
//	  5:   5 cent rounding 3.43 => 3.45
//	 10:  10 cent rounding 3.45 => 3.50 (5 gets rounded up)
//	 25:  25 cent rounding 3.41 => 3.50
//	 50:  50 cent rounding 3.75 => 4.00
//	100: 100 cent rounding 3.50 => 4.00
//
// For more details: https://en.wikipedia.org/wiki/Cash_rounding
func (__obf_aafc5e233dfea784 Decimal) RoundCash(__obf_f5b39b07684a644b uint8) Decimal {
	var __obf_d89d3221cfb69699 *big.Int
	switch __obf_f5b39b07684a644b {
	case 5:
		__obf_d89d3221cfb69699 = __obf_abb4b37cdf2c9d6e
	case 10:
		__obf_d89d3221cfb69699 = __obf_cd9358bd6469dc36
	case 25:
		__obf_d89d3221cfb69699 = __obf_bee298c3e91651d7
	case 50:
		__obf_d89d3221cfb69699 = __obf_3f8f48b9238d224e
	case 100:
		__obf_d89d3221cfb69699 = __obf_f136d7fff3e7d6f7
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_f5b39b07684a644b))
	}
	__obf_e8e4b4c978fe6621 := Decimal{__obf_22dafd20a48df74e: __obf_d89d3221cfb69699}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_aafc5e233dfea784.Mul(__obf_e8e4b4c978fe6621).Round(0).Div(__obf_e8e4b4c978fe6621).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_aafc5e233dfea784 Decimal) Floor() Decimal {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()

	if __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f >= 0 {
		return __obf_aafc5e233dfea784
	}
	__obf_10eec1d7f8ad279f := big.NewInt(10)
	__obf_10eec1d7f8ad279f.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_10eec1d7f8ad279f, big.NewInt(-int64(__obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f)), nil)
	__obf_6f45f5fffe86c965 := new(big.Int).Div(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e, __obf_10eec1d7f8ad279f)
	return Decimal{__obf_22dafd20a48df74e: __obf_6f45f5fffe86c965,

		// Ceil returns the nearest integer value greater than or equal to d.
		__obf_10eec1d7f8ad279f: 0}
}

func (__obf_aafc5e233dfea784 Decimal) Ceil() Decimal {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()

	if __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f >= 0 {
		return __obf_aafc5e233dfea784
	}
	__obf_10eec1d7f8ad279f := big.NewInt(10)
	__obf_10eec1d7f8ad279f.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_10eec1d7f8ad279f, big.NewInt(-int64(__obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f)), nil)
	__obf_6f45f5fffe86c965, __obf_ade33d4daf3b3fc1 := new(big.Int).DivMod(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e, __obf_10eec1d7f8ad279f, new(big.Int))
	if __obf_ade33d4daf3b3fc1.Cmp(__obf_f0e1aaf9670d8421) != 0 {
		__obf_6f45f5fffe86c965.
			Add(__obf_6f45f5fffe86c965, __obf_f136d7fff3e7d6f7)
	}
	return Decimal{__obf_22dafd20a48df74e: __obf_6f45f5fffe86c965,

		// Truncate truncates off digits from the number, without rounding.
		//
		// NOTE: precision is the last digit that will not be truncated (must be >= 0).
		//
		// Example:
		//
		//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
		__obf_10eec1d7f8ad279f: 0}
}

func (__obf_aafc5e233dfea784 Decimal) Truncate(__obf_6eb04b4176b00c83 int32) Decimal {
	__obf_aafc5e233dfea784.__obf_04511f44d644c8ae()
	if __obf_6eb04b4176b00c83 >= 0 && -__obf_6eb04b4176b00c83 > __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f {
		return __obf_aafc5e233dfea784.__obf_51b05428a428dd2d(-__obf_6eb04b4176b00c83)
	}
	return __obf_aafc5e233dfea784
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_aafc5e233dfea784 *Decimal) UnmarshalJSON(__obf_588b0d05ec5b802e []byte) error {
	if string(__obf_588b0d05ec5b802e) == "null" {
		return nil
	}
	__obf_8966ceb6be5a7f08, __obf_2e2f1e21b557821f := __obf_f7d5b68e22816626(__obf_588b0d05ec5b802e)
	if __obf_2e2f1e21b557821f != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_588b0d05ec5b802e, __obf_2e2f1e21b557821f)
	}
	__obf_4cc30aac3ca67486, __obf_2e2f1e21b557821f := NewFromString(__obf_8966ceb6be5a7f08)
	*__obf_aafc5e233dfea784 = __obf_4cc30aac3ca67486
	if __obf_2e2f1e21b557821f != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_8966ceb6be5a7f08, __obf_2e2f1e21b557821f)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_aafc5e233dfea784 Decimal) MarshalJSON() ([]byte, error) {
	var __obf_8966ceb6be5a7f08 string
	if MarshalJSONWithoutQuotes {
		__obf_8966ceb6be5a7f08 = __obf_aafc5e233dfea784.String()
	} else {
		__obf_8966ceb6be5a7f08 = "\"" + __obf_aafc5e233dfea784.String() + "\""
	}
	return []byte(__obf_8966ceb6be5a7f08), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_aafc5e233dfea784 *Decimal) UnmarshalBinary(__obf_0449bc69685b54b7 []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_0449bc69685b54b7) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_0449bc69685b54b7, len(__obf_0449bc69685b54b7))
	}
	__obf_aafc5e233dfea784.

		// Extract the exponent
		__obf_10eec1d7f8ad279f = int32(binary.BigEndian.Uint32(__obf_0449bc69685b54b7[:4]))
	__obf_aafc5e233dfea784.

		// Extract the value
		__obf_22dafd20a48df74e = new(big.Int)
	if __obf_2e2f1e21b557821f := __obf_aafc5e233dfea784.__obf_22dafd20a48df74e.GobDecode(__obf_0449bc69685b54b7[4:]); __obf_2e2f1e21b557821f != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_0449bc69685b54b7, __obf_2e2f1e21b557821f)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_aafc5e233dfea784 Decimal) MarshalBinary() (__obf_0449bc69685b54b7 []byte, __obf_2e2f1e21b557821f error) {
	// exp is written first, but encode value first to know output size
	var __obf_1b1a6f0c546c6e66 []byte
	if __obf_1b1a6f0c546c6e66, __obf_2e2f1e21b557821f = __obf_aafc5e233dfea784.__obf_22dafd20a48df74e.GobEncode(); __obf_2e2f1e21b557821f != nil {
		return nil, __obf_2e2f1e21b557821f
	}
	__obf_a5b2dc9cadc92ce5 := // Write the exponent in front, since it's a fixed size
		make([]byte, 4, len(__obf_1b1a6f0c546c6e66)+4)
	binary.BigEndian.PutUint32(__obf_a5b2dc9cadc92ce5, uint32(__obf_aafc5e233dfea784.

		// Return the byte array
		__obf_10eec1d7f8ad279f))

	return append(__obf_a5b2dc9cadc92ce5, __obf_1b1a6f0c546c6e66...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_aafc5e233dfea784 *Decimal) Scan(__obf_22dafd20a48df74e any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_58d61934b6551bac := __obf_22dafd20a48df74e.(type) {

	case float32:
		*__obf_aafc5e233dfea784 = NewFromFloat(float64(__obf_58d61934b6551bac))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_aafc5e233dfea784 = NewFromFloat(__obf_58d61934b6551bac)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_aafc5e233dfea784 = New(__obf_58d61934b6551bac, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_aafc5e233dfea784 = NewFromUint64(__obf_58d61934b6551bac)
		return nil

	default:
		__obf_8966ceb6be5a7f08,
			// default is trying to interpret value stored as string
			__obf_2e2f1e21b557821f := __obf_f7d5b68e22816626(__obf_58d61934b6551bac)
		if __obf_2e2f1e21b557821f != nil {
			return __obf_2e2f1e21b557821f
		}
		*__obf_aafc5e233dfea784, __obf_2e2f1e21b557821f = NewFromString(__obf_8966ceb6be5a7f08)
		return __obf_2e2f1e21b557821f
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_aafc5e233dfea784 Decimal) Value() (driver.Value, error) {
	return __obf_aafc5e233dfea784.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_aafc5e233dfea784 *Decimal) UnmarshalText(__obf_aa9c5208346a835f []byte) error {
	__obf_8966ceb6be5a7f08 := string(__obf_aa9c5208346a835f)
	__obf_012da267317dc449, __obf_2e2f1e21b557821f := NewFromString(__obf_8966ceb6be5a7f08)
	*__obf_aafc5e233dfea784 = __obf_012da267317dc449
	if __obf_2e2f1e21b557821f != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_8966ceb6be5a7f08, __obf_2e2f1e21b557821f)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_aafc5e233dfea784 Decimal) MarshalText() (__obf_aa9c5208346a835f []byte, __obf_2e2f1e21b557821f error) {
	return []byte(__obf_aafc5e233dfea784.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_aafc5e233dfea784 Decimal) GobEncode() ([]byte, error) {
	return __obf_aafc5e233dfea784.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_aafc5e233dfea784 *Decimal) GobDecode(__obf_0449bc69685b54b7 []byte) error {
	return __obf_aafc5e233dfea784.UnmarshalBinary(__obf_0449bc69685b54b7)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_aafc5e233dfea784 Decimal) StringScaled(__obf_10eec1d7f8ad279f int32) string {
	return __obf_aafc5e233dfea784.__obf_51b05428a428dd2d(__obf_10eec1d7f8ad279f).String()
}

func (__obf_aafc5e233dfea784 Decimal) string(__obf_1a90b7c83c28a8ae bool) string {
	if __obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f >= 0 {
		return __obf_aafc5e233dfea784.__obf_51b05428a428dd2d(0).__obf_22dafd20a48df74e.String()
	}
	__obf_1d9c7a78d7fb520b := new(big.Int).Abs(__obf_aafc5e233dfea784.__obf_22dafd20a48df74e)
	__obf_8966ceb6be5a7f08 := __obf_1d9c7a78d7fb520b.String()

	var __obf_61a32cbfc444f697, __obf_a827aa0132e7fa72 string
	__obf_d8bedcdfbd16349b := // NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
		// and you are on a 32-bit machine. Won't fix this super-edge case.
		int(__obf_aafc5e233dfea784.__obf_10eec1d7f8ad279f)
	if len(__obf_8966ceb6be5a7f08) > -__obf_d8bedcdfbd16349b {
		__obf_61a32cbfc444f697 = __obf_8966ceb6be5a7f08[:len(__obf_8966ceb6be5a7f08)+__obf_d8bedcdfbd16349b]
		__obf_a827aa0132e7fa72 = __obf_8966ceb6be5a7f08[len(__obf_8966ceb6be5a7f08)+__obf_d8bedcdfbd16349b:]
	} else {
		__obf_61a32cbfc444f697 = "0"
		__obf_7e569e830a691203 := -__obf_d8bedcdfbd16349b - len(__obf_8966ceb6be5a7f08)
		__obf_a827aa0132e7fa72 = strings.Repeat("0", __obf_7e569e830a691203) + __obf_8966ceb6be5a7f08
	}

	if __obf_1a90b7c83c28a8ae {
		__obf_03985402d384eb80 := len(__obf_a827aa0132e7fa72) - 1
		for ; __obf_03985402d384eb80 >= 0; __obf_03985402d384eb80-- {
			if __obf_a827aa0132e7fa72[__obf_03985402d384eb80] != '0' {
				break
			}
		}
		__obf_a827aa0132e7fa72 = __obf_a827aa0132e7fa72[:__obf_03985402d384eb80+1]
	}
	__obf_098a2e9c0d0c6740 := __obf_61a32cbfc444f697
	if len(__obf_a827aa0132e7fa72) > 0 {
		__obf_098a2e9c0d0c6740 += "." + __obf_a827aa0132e7fa72
	}

	if __obf_aafc5e233dfea784.__obf_22dafd20a48df74e.Sign() < 0 {
		return "-" + __obf_098a2e9c0d0c6740
	}

	return __obf_098a2e9c0d0c6740
}

func (__obf_aafc5e233dfea784 *Decimal) __obf_04511f44d644c8ae() {
	if __obf_aafc5e233dfea784.__obf_22dafd20a48df74e == nil {
		__obf_aafc5e233dfea784.__obf_22dafd20a48df74e = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_70fba8925f23aca8 Decimal, __obf_a2d5c43554fb647f ...Decimal) Decimal {
	__obf_7c4ba3960b4b8cfe := __obf_70fba8925f23aca8
	for _, __obf_175f7a0b5f412205 := range __obf_a2d5c43554fb647f {
		if __obf_175f7a0b5f412205.Cmp(__obf_7c4ba3960b4b8cfe) < 0 {
			__obf_7c4ba3960b4b8cfe = __obf_175f7a0b5f412205
		}
	}
	return __obf_7c4ba3960b4b8cfe
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_70fba8925f23aca8 Decimal, __obf_a2d5c43554fb647f ...Decimal) Decimal {
	__obf_7c4ba3960b4b8cfe := __obf_70fba8925f23aca8
	for _, __obf_175f7a0b5f412205 := range __obf_a2d5c43554fb647f {
		if __obf_175f7a0b5f412205.Cmp(__obf_7c4ba3960b4b8cfe) > 0 {
			__obf_7c4ba3960b4b8cfe = __obf_175f7a0b5f412205
		}
	}
	return __obf_7c4ba3960b4b8cfe
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_70fba8925f23aca8 Decimal, __obf_a2d5c43554fb647f ...Decimal) Decimal {
	__obf_fd8d51e688f8cebc := __obf_70fba8925f23aca8
	for _, __obf_175f7a0b5f412205 := range __obf_a2d5c43554fb647f {
		__obf_fd8d51e688f8cebc = __obf_fd8d51e688f8cebc.Add(__obf_175f7a0b5f412205)
	}

	return __obf_fd8d51e688f8cebc
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_70fba8925f23aca8 Decimal, __obf_a2d5c43554fb647f ...Decimal) Decimal {
	__obf_29c3b43377d473e1 := New(int64(len(__obf_a2d5c43554fb647f)+1), 0)
	__obf_6cebf47a1809ac55 := Sum(__obf_70fba8925f23aca8, __obf_a2d5c43554fb647f...)
	return __obf_6cebf47a1809ac55.Div(__obf_29c3b43377d473e1)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_1f14eb770089284a Decimal, __obf_dea290dc9968f164 Decimal) (Decimal, Decimal) {
	__obf_1f14eb770089284a.__obf_04511f44d644c8ae()
	__obf_dea290dc9968f164.__obf_04511f44d644c8ae()

	if __obf_1f14eb770089284a.__obf_10eec1d7f8ad279f < __obf_dea290dc9968f164.__obf_10eec1d7f8ad279f {
		return __obf_1f14eb770089284a, __obf_dea290dc9968f164.__obf_51b05428a428dd2d(__obf_1f14eb770089284a.__obf_10eec1d7f8ad279f)
	} else if __obf_1f14eb770089284a.__obf_10eec1d7f8ad279f > __obf_dea290dc9968f164.__obf_10eec1d7f8ad279f {
		return __obf_1f14eb770089284a.__obf_51b05428a428dd2d(__obf_dea290dc9968f164.__obf_10eec1d7f8ad279f), __obf_dea290dc9968f164
	}

	return __obf_1f14eb770089284a, __obf_dea290dc9968f164
}

func __obf_f7d5b68e22816626(__obf_22dafd20a48df74e any) (string, error) {
	var __obf_c318805e404b06be []byte

	switch __obf_58d61934b6551bac := __obf_22dafd20a48df74e.(type) {
	case string:
		__obf_c318805e404b06be = []byte(__obf_58d61934b6551bac)
	case []byte:
		__obf_c318805e404b06be = __obf_58d61934b6551bac
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_22dafd20a48df74e,

			// If the amount is quoted, strip the quotes
			__obf_22dafd20a48df74e)
	}

	if len(__obf_c318805e404b06be) > 2 && __obf_c318805e404b06be[0] == '"' && __obf_c318805e404b06be[len(__obf_c318805e404b06be)-1] == '"' {
		__obf_c318805e404b06be = __obf_c318805e404b06be[1 : len(__obf_c318805e404b06be)-1]
	}
	return string(__obf_c318805e404b06be), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the datautil.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_aafc5e233dfea784 Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_aafc5e233dfea784,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_aafc5e233dfea784 *NullDecimal) Scan(__obf_22dafd20a48df74e any) error {
	if __obf_22dafd20a48df74e == nil {
		__obf_aafc5e233dfea784.
			Valid = false
		return nil
	}
	__obf_aafc5e233dfea784.
		Valid = true
	return __obf_aafc5e233dfea784.Decimal.Scan(__obf_22dafd20a48df74e)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_aafc5e233dfea784 NullDecimal) Value() (driver.Value, error) {
	if !__obf_aafc5e233dfea784.Valid {
		return nil, nil
	}
	return __obf_aafc5e233dfea784.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_aafc5e233dfea784 *NullDecimal) UnmarshalJSON(__obf_588b0d05ec5b802e []byte) error {
	if string(__obf_588b0d05ec5b802e) == "null" {
		__obf_aafc5e233dfea784.
			Valid = false
		return nil
	}
	__obf_aafc5e233dfea784.
		Valid = true
	return __obf_aafc5e233dfea784.Decimal.UnmarshalJSON(__obf_588b0d05ec5b802e)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_aafc5e233dfea784 NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_aafc5e233dfea784.Valid {
		return []byte("null"), nil
	}
	return __obf_aafc5e233dfea784.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_aafc5e233dfea784 *NullDecimal) UnmarshalText(__obf_aa9c5208346a835f []byte) error {
	__obf_8966ceb6be5a7f08 := string(__obf_aa9c5208346a835f)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_8966ceb6be5a7f08 == "" {
		__obf_aafc5e233dfea784.
			Valid = false
		return nil
	}
	if __obf_2e2f1e21b557821f := __obf_aafc5e233dfea784.Decimal.UnmarshalText(__obf_aa9c5208346a835f); __obf_2e2f1e21b557821f != nil {
		__obf_aafc5e233dfea784.
			Valid = false
		return __obf_2e2f1e21b557821f
	}
	__obf_aafc5e233dfea784.
		Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_aafc5e233dfea784 NullDecimal) MarshalText() (__obf_aa9c5208346a835f []byte, __obf_2e2f1e21b557821f error) {
	if !__obf_aafc5e233dfea784.Valid {
		return []byte{}, nil
	}
	return __obf_aafc5e233dfea784.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_aafc5e233dfea784 Decimal) Atan() Decimal {
	if __obf_aafc5e233dfea784.Equal(NewFromFloat(0.0)) {
		return __obf_aafc5e233dfea784
	}
	if __obf_aafc5e233dfea784.GreaterThan(NewFromFloat(0.0)) {
		return __obf_aafc5e233dfea784.__obf_10b17920d67fb403()
	}
	return __obf_aafc5e233dfea784.Neg().__obf_10b17920d67fb403().Neg()
}

func (__obf_aafc5e233dfea784 Decimal) __obf_5018ca138b2dbf36() Decimal {
	P0 := NewFromFloat(-8.750608600031904122785e-01)
	P1 := NewFromFloat(-1.615753718733365076637e+01)
	P2 := NewFromFloat(-7.500855792314704667340e+01)
	P3 := NewFromFloat(-1.228866684490136173410e+02)
	P4 := NewFromFloat(-6.485021904942025371773e+01)
	Q0 := NewFromFloat(2.485846490142306297962e+01)
	Q1 := NewFromFloat(1.650270098316988542046e+02)
	Q2 := NewFromFloat(4.328810604912902668951e+02)
	Q3 := NewFromFloat(4.853903996359136964868e+02)
	Q4 := NewFromFloat(1.945506571482613964425e+02)
	__obf_6f45f5fffe86c965 := __obf_aafc5e233dfea784.Mul(__obf_aafc5e233dfea784)
	__obf_1a7cf90a0681f730 := P0.Mul(__obf_6f45f5fffe86c965).Add(P1).Mul(__obf_6f45f5fffe86c965).Add(P2).Mul(__obf_6f45f5fffe86c965).Add(P3).Mul(__obf_6f45f5fffe86c965).Add(P4).Mul(__obf_6f45f5fffe86c965)
	__obf_d612ece1ba6a7ef0 := __obf_6f45f5fffe86c965.Add(Q0).Mul(__obf_6f45f5fffe86c965).Add(Q1).Mul(__obf_6f45f5fffe86c965).Add(Q2).Mul(__obf_6f45f5fffe86c965).Add(Q3).Mul(__obf_6f45f5fffe86c965).Add(Q4)
	__obf_6f45f5fffe86c965 = __obf_1a7cf90a0681f730.Div(__obf_d612ece1ba6a7ef0)
	__obf_6f45f5fffe86c965 = __obf_aafc5e233dfea784.Mul(__obf_6f45f5fffe86c965).Add(__obf_aafc5e233dfea784)
	return __obf_6f45f5fffe86c965
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_aafc5e233dfea784 Decimal) __obf_10b17920d67fb403() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)
	__obf_746ea9129e6690eb := // tan(3*pi/8)
		NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_aafc5e233dfea784.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_aafc5e233dfea784.__obf_5018ca138b2dbf36()
	}
	if __obf_aafc5e233dfea784.GreaterThan(Tan3pio8) {
		return __obf_746ea9129e6690eb.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_aafc5e233dfea784).__obf_5018ca138b2dbf36()).Add(Morebits)
	}
	return __obf_746ea9129e6690eb.Div(NewFromFloat(4.0)).Add((__obf_aafc5e233dfea784.Sub(NewFromFloat(1.0)).Div(__obf_aafc5e233dfea784.Add(NewFromFloat(1.0)))).__obf_5018ca138b2dbf36()).Add(NewFromFloat(0.5).Mul(Morebits))
}

// sin coefficients
var _sin = [...]Decimal{
	NewFromFloat(1.58962301576546568060e-10), // 0x3de5d8fd1fd19ccd
	NewFromFloat(-2.50507477628578072866e-8), // 0xbe5ae5e5a9291f5d
	NewFromFloat(2.75573136213857245213e-6),  // 0x3ec71de3567d48a1
	NewFromFloat(-1.98412698295895385996e-4), // 0xbf2a01a019bfdf03
	NewFromFloat(8.33333333332211858878e-3),  // 0x3f8111111110f7d0
	NewFromFloat(-1.66666666666666307295e-1), // 0xbfc5555555555548
}

// Sin returns the sine of the radian argument x.
func (__obf_aafc5e233dfea784 Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_aafc5e233dfea784.Equal(NewFromFloat(0.0)) {
		return __obf_aafc5e233dfea784
	}
	__obf_38ae79a7ec63ef24 := // make argument positive but save the sign
		false
	if __obf_aafc5e233dfea784.LessThan(NewFromFloat(0.0)) {
		__obf_aafc5e233dfea784 = __obf_aafc5e233dfea784.Neg()
		__obf_38ae79a7ec63ef24 = true
	}
	__obf_629f6b9626aeca5f := __obf_aafc5e233dfea784.Mul(M4PI).IntPart()
	__obf_b704283d9d1cd34a := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_629f6b9626aeca5f)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_629f6b9626aeca5f&1 == 1 {
		__obf_629f6b9626aeca5f++
		__obf_b704283d9d1cd34a = __obf_b704283d9d1cd34a.Add(NewFromFloat(1.0))
	}
	__obf_629f6b9626aeca5f &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_629f6b9626aeca5f > 3 {
		__obf_38ae79a7ec63ef24 = !__obf_38ae79a7ec63ef24
		__obf_629f6b9626aeca5f -= 4
	}
	__obf_6f45f5fffe86c965 := __obf_aafc5e233dfea784.Sub(__obf_b704283d9d1cd34a.Mul(PI4A)).Sub(__obf_b704283d9d1cd34a.Mul(PI4B)).Sub(__obf_b704283d9d1cd34a.Mul(PI4C))
	__obf_76f604259ec8caea := // Extended precision modular arithmetic
		__obf_6f45f5fffe86c965.Mul(__obf_6f45f5fffe86c965)

	if __obf_629f6b9626aeca5f == 1 || __obf_629f6b9626aeca5f == 2 {
		__obf_83094a3f8cd288f0 := __obf_76f604259ec8caea.Mul(__obf_76f604259ec8caea).Mul(_cos[0].Mul(__obf_76f604259ec8caea).Add(_cos[1]).Mul(__obf_76f604259ec8caea).Add(_cos[2]).Mul(__obf_76f604259ec8caea).Add(_cos[3]).Mul(__obf_76f604259ec8caea).Add(_cos[4]).Mul(__obf_76f604259ec8caea).Add(_cos[5]))
		__obf_b704283d9d1cd34a = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_76f604259ec8caea)).Add(__obf_83094a3f8cd288f0)
	} else {
		__obf_b704283d9d1cd34a = __obf_6f45f5fffe86c965.Add(__obf_6f45f5fffe86c965.Mul(__obf_76f604259ec8caea).Mul(_sin[0].Mul(__obf_76f604259ec8caea).Add(_sin[1]).Mul(__obf_76f604259ec8caea).Add(_sin[2]).Mul(__obf_76f604259ec8caea).Add(_sin[3]).Mul(__obf_76f604259ec8caea).Add(_sin[4]).Mul(__obf_76f604259ec8caea).Add(_sin[5])))
	}
	if __obf_38ae79a7ec63ef24 {
		__obf_b704283d9d1cd34a = __obf_b704283d9d1cd34a.Neg()
	}
	return __obf_b704283d9d1cd34a
}

// cos coefficients
var _cos = [...]Decimal{
	NewFromFloat(-1.13585365213876817300e-11), // 0xbda8fa49a0861a9b
	NewFromFloat(2.08757008419747316778e-9),   // 0x3e21ee9d7b4e3f05
	NewFromFloat(-2.75573141792967388112e-7),  // 0xbe927e4f7eac4bc6
	NewFromFloat(2.48015872888517045348e-5),   // 0x3efa01a019c844f5
	NewFromFloat(-1.38888888888730564116e-3),  // 0xbf56c16c16c14f91
	NewFromFloat(4.16666666666665929218e-2),   // 0x3fa555555555554b
}

// Cos returns the cosine of the radian argument x.
func (__obf_aafc5e233dfea784 Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)  // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)  // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15) // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125)
	__obf_38ae79a7ec63ef24 := // 4/pi

		// make argument positive
		false
	if __obf_aafc5e233dfea784.LessThan(NewFromFloat(0.0)) {
		__obf_aafc5e233dfea784 = __obf_aafc5e233dfea784.Neg()
	}
	__obf_629f6b9626aeca5f := __obf_aafc5e233dfea784.Mul(M4PI).IntPart()
	__obf_b704283d9d1cd34a := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_629f6b9626aeca5f)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_629f6b9626aeca5f&1 == 1 {
		__obf_629f6b9626aeca5f++
		__obf_b704283d9d1cd34a = __obf_b704283d9d1cd34a.Add(NewFromFloat(1.0))
	}
	__obf_629f6b9626aeca5f &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_629f6b9626aeca5f > 3 {
		__obf_38ae79a7ec63ef24 = !__obf_38ae79a7ec63ef24
		__obf_629f6b9626aeca5f -= 4
	}
	if __obf_629f6b9626aeca5f > 1 {
		__obf_38ae79a7ec63ef24 = !__obf_38ae79a7ec63ef24
	}
	__obf_6f45f5fffe86c965 := __obf_aafc5e233dfea784.Sub(__obf_b704283d9d1cd34a.Mul(PI4A)).Sub(__obf_b704283d9d1cd34a.Mul(PI4B)).Sub(__obf_b704283d9d1cd34a.Mul(PI4C))
	__obf_76f604259ec8caea := // Extended precision modular arithmetic
		__obf_6f45f5fffe86c965.Mul(__obf_6f45f5fffe86c965)

	if __obf_629f6b9626aeca5f == 1 || __obf_629f6b9626aeca5f == 2 {
		__obf_b704283d9d1cd34a = __obf_6f45f5fffe86c965.Add(__obf_6f45f5fffe86c965.Mul(__obf_76f604259ec8caea).Mul(_sin[0].Mul(__obf_76f604259ec8caea).Add(_sin[1]).Mul(__obf_76f604259ec8caea).Add(_sin[2]).Mul(__obf_76f604259ec8caea).Add(_sin[3]).Mul(__obf_76f604259ec8caea).Add(_sin[4]).Mul(__obf_76f604259ec8caea).Add(_sin[5])))
	} else {
		__obf_83094a3f8cd288f0 := __obf_76f604259ec8caea.Mul(__obf_76f604259ec8caea).Mul(_cos[0].Mul(__obf_76f604259ec8caea).Add(_cos[1]).Mul(__obf_76f604259ec8caea).Add(_cos[2]).Mul(__obf_76f604259ec8caea).Add(_cos[3]).Mul(__obf_76f604259ec8caea).Add(_cos[4]).Mul(__obf_76f604259ec8caea).Add(_cos[5]))
		__obf_b704283d9d1cd34a = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_76f604259ec8caea)).Add(__obf_83094a3f8cd288f0)
	}
	if __obf_38ae79a7ec63ef24 {
		__obf_b704283d9d1cd34a = __obf_b704283d9d1cd34a.Neg()
	}
	return __obf_b704283d9d1cd34a
}

var _tanP = [...]Decimal{
	NewFromFloat(-1.30936939181383777646e+4), // 0xc0c992d8d24f3f38
	NewFromFloat(1.15351664838587416140e+6),  // 0x413199eca5fc9ddd
	NewFromFloat(-1.79565251976484877988e+7), // 0xc1711fead3299176
}
var _tanQ = [...]Decimal{
	NewFromFloat(1.00000000000000000000e+0),
	NewFromFloat(1.36812963470692954678e+4),  //0x40cab8a5eeb36572
	NewFromFloat(-1.32089234440210967447e+6), //0xc13427bc582abc96
	NewFromFloat(2.50083801823357915839e+7),  //0x4177d98fc2ead8ef
	NewFromFloat(-5.38695755929454629881e+7), //0xc189afe03cbe5a31
}

// Tan returns the tangent of the radian argument x.
func (__obf_aafc5e233dfea784 Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_aafc5e233dfea784.Equal(NewFromFloat(0.0)) {
		return __obf_aafc5e233dfea784
	}
	__obf_38ae79a7ec63ef24 := // make argument positive but save the sign
		false
	if __obf_aafc5e233dfea784.LessThan(NewFromFloat(0.0)) {
		__obf_aafc5e233dfea784 = __obf_aafc5e233dfea784.Neg()
		__obf_38ae79a7ec63ef24 = true
	}
	__obf_629f6b9626aeca5f := __obf_aafc5e233dfea784.Mul(M4PI).IntPart()
	__obf_b704283d9d1cd34a := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_629f6b9626aeca5f)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_629f6b9626aeca5f&1 == 1 {
		__obf_629f6b9626aeca5f++
		__obf_b704283d9d1cd34a = __obf_b704283d9d1cd34a.Add(NewFromFloat(1.0))
	}
	__obf_6f45f5fffe86c965 := __obf_aafc5e233dfea784.Sub(__obf_b704283d9d1cd34a.Mul(PI4A)).Sub(__obf_b704283d9d1cd34a.Mul(PI4B)).Sub(__obf_b704283d9d1cd34a.Mul(PI4C))
	__obf_76f604259ec8caea := // Extended precision modular arithmetic
		__obf_6f45f5fffe86c965.Mul(__obf_6f45f5fffe86c965)

	if __obf_76f604259ec8caea.GreaterThan(NewFromFloat(1e-14)) {
		__obf_83094a3f8cd288f0 := __obf_76f604259ec8caea.Mul(_tanP[0].Mul(__obf_76f604259ec8caea).Add(_tanP[1]).Mul(__obf_76f604259ec8caea).Add(_tanP[2]))
		__obf_7f301de8ed5f0d7f := __obf_76f604259ec8caea.Add(_tanQ[1]).Mul(__obf_76f604259ec8caea).Add(_tanQ[2]).Mul(__obf_76f604259ec8caea).Add(_tanQ[3]).Mul(__obf_76f604259ec8caea).Add(_tanQ[4])
		__obf_b704283d9d1cd34a = __obf_6f45f5fffe86c965.Add(__obf_6f45f5fffe86c965.Mul(__obf_83094a3f8cd288f0.Div(__obf_7f301de8ed5f0d7f)))
	} else {
		__obf_b704283d9d1cd34a = __obf_6f45f5fffe86c965
	}
	if __obf_629f6b9626aeca5f&2 == 2 {
		__obf_b704283d9d1cd34a = NewFromFloat(-1.0).Div(__obf_b704283d9d1cd34a)
	}
	if __obf_38ae79a7ec63ef24 {
		__obf_b704283d9d1cd34a = __obf_b704283d9d1cd34a.Neg()
	}
	return __obf_b704283d9d1cd34a
}

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

type __obf_4cc30aac3ca67486 struct {
	__obf_aafc5e233dfea784 [800]byte
	__obf_5c82a3d658ffb3ad int  // digits, big-endian representation
	__obf_2b78448fb2f1bffe int  // number of digits used
	__obf_ef34c5939301e47a bool // decimal point
	__obf_0b587bdd951615f7 bool // negative flag
	// discarded nonzero digits beyond d[:nd]
}

func (__obf_c09a2b51ef153b53 *__obf_4cc30aac3ca67486) String() string {
	__obf_d6441e69f8e999be := 10 + __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad
	if __obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe > 0 {
		__obf_d6441e69f8e999be += __obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe
	}
	if __obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe < 0 {
		__obf_d6441e69f8e999be += -__obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe
	}
	__obf_2fa64d0aef13da4a := make([]byte, __obf_d6441e69f8e999be)
	__obf_83094a3f8cd288f0 := 0
	switch {
	case __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad == 0:
		return "0"

	case __obf_c09a2b51ef153b53.
		// zeros fill space between decimal point and digits
		__obf_2b78448fb2f1bffe <= 0:
		__obf_2fa64d0aef13da4a[__obf_83094a3f8cd288f0] = '0'
		__obf_83094a3f8cd288f0++
		__obf_2fa64d0aef13da4a[__obf_83094a3f8cd288f0] = '.'
		__obf_83094a3f8cd288f0++
		__obf_83094a3f8cd288f0 += __obf_6cefa1f454d13d94(__obf_2fa64d0aef13da4a[__obf_83094a3f8cd288f0 : __obf_83094a3f8cd288f0+-__obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe])
		__obf_83094a3f8cd288f0 += copy(__obf_2fa64d0aef13da4a[__obf_83094a3f8cd288f0:], __obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[0:__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad])

	case __obf_c09a2b51ef153b53.
		// decimal point in middle of digits
		__obf_2b78448fb2f1bffe < __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad:
		__obf_83094a3f8cd288f0 += copy(__obf_2fa64d0aef13da4a[__obf_83094a3f8cd288f0:], __obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[0:__obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe])
		__obf_2fa64d0aef13da4a[__obf_83094a3f8cd288f0] = '.'
		__obf_83094a3f8cd288f0++
		__obf_83094a3f8cd288f0 += copy(__obf_2fa64d0aef13da4a[__obf_83094a3f8cd288f0:], __obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe:

		// zeros fill space between digits and decimal point
		__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad])

	default:
		__obf_83094a3f8cd288f0 += copy(__obf_2fa64d0aef13da4a[__obf_83094a3f8cd288f0:], __obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[0:__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad])
		__obf_83094a3f8cd288f0 += __obf_6cefa1f454d13d94(__obf_2fa64d0aef13da4a[__obf_83094a3f8cd288f0 : __obf_83094a3f8cd288f0+__obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe-__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad])
	}
	return string(__obf_2fa64d0aef13da4a[0:__obf_83094a3f8cd288f0])
}

func __obf_6cefa1f454d13d94(__obf_61bf364499524f13 []byte) int {
	for __obf_03985402d384eb80 := range __obf_61bf364499524f13 {
		__obf_61bf364499524f13[__obf_03985402d384eb80] = '0'
	}
	return len(__obf_61bf364499524f13)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_31d3f5cacc8faba5(__obf_c09a2b51ef153b53 *__obf_4cc30aac3ca67486) {
	for __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad > 0 && __obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad-1] == '0' {
		__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad--
	}
	if __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad == 0 {
		__obf_c09a2b51ef153b53.

			// Assign v to a.
			__obf_2b78448fb2f1bffe = 0
	}
}

func (__obf_c09a2b51ef153b53 *__obf_4cc30aac3ca67486) Assign(__obf_58d61934b6551bac uint64) {
	var __obf_2fa64d0aef13da4a [24]byte
	__obf_d6441e69f8e999be := // Write reversed decimal in buf.
		0
	for __obf_58d61934b6551bac > 0 {
		__obf_1619133bba33e0be := __obf_58d61934b6551bac / 10
		__obf_58d61934b6551bac -= 10 * __obf_1619133bba33e0be
		__obf_2fa64d0aef13da4a[__obf_d6441e69f8e999be] = byte(__obf_58d61934b6551bac + '0')
		__obf_d6441e69f8e999be++
		__obf_58d61934b6551bac = __obf_1619133bba33e0be
	}
	__obf_c09a2b51ef153b53.

		// Reverse again to produce forward decimal in a.d.
		__obf_5c82a3d658ffb3ad = 0
	for __obf_d6441e69f8e999be--; __obf_d6441e69f8e999be >= 0; __obf_d6441e69f8e999be-- {
		__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad] = __obf_2fa64d0aef13da4a[__obf_d6441e69f8e999be]
		__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad++
	}
	__obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe = __obf_c09a2b51ef153b53.

		// Maximum shift that we can do in one pass without overflow.
		// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
		__obf_5c82a3d658ffb3ad
	__obf_31d3f5cacc8faba5(__obf_c09a2b51ef153b53)
}

const __obf_68532746ac5ed4e4 = 32 << (^uint(0) >> 63)
const __obf_a399770894bf00d0 = __obf_68532746ac5ed4e4 - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_f04759d2e55de434(__obf_c09a2b51ef153b53 *__obf_4cc30aac3ca67486, __obf_8f1bcc26019920eb uint) {
	__obf_5039aeab743c839a := 0
	__obf_83094a3f8cd288f0 := // read pointer
		0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_d6441e69f8e999be uint
	for ; __obf_d6441e69f8e999be>>__obf_8f1bcc26019920eb == 0; __obf_5039aeab743c839a++ {
		if __obf_5039aeab743c839a >= __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad {
			if __obf_d6441e69f8e999be == 0 {
				__obf_c09a2b51ef153b53.
					// a == 0; shouldn't get here, but handle anyway.
					__obf_5c82a3d658ffb3ad = 0
				return
			}
			for __obf_d6441e69f8e999be>>__obf_8f1bcc26019920eb == 0 {
				__obf_d6441e69f8e999be = __obf_d6441e69f8e999be * 10
				__obf_5039aeab743c839a++
			}
			break
		}
		__obf_d520aaab80df92cd := uint(__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_5039aeab743c839a])
		__obf_d6441e69f8e999be = __obf_d6441e69f8e999be*10 + __obf_d520aaab80df92cd - '0'
	}
	__obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe -= __obf_5039aeab743c839a - 1

	var __obf_06b5f862243c1836 uint = (1 << __obf_8f1bcc26019920eb) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_5039aeab743c839a < __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad; __obf_5039aeab743c839a++ {
		__obf_d520aaab80df92cd := uint(__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_5039aeab743c839a])
		__obf_06503bb8ab96d782 := __obf_d6441e69f8e999be >> __obf_8f1bcc26019920eb
		__obf_d6441e69f8e999be &= __obf_06b5f862243c1836
		__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_83094a3f8cd288f0] = byte(__obf_06503bb8ab96d782 + '0')
		__obf_83094a3f8cd288f0++
		__obf_d6441e69f8e999be = __obf_d6441e69f8e999be*10 + __obf_d520aaab80df92cd - '0'
	}

	// Put down extra digits.
	for __obf_d6441e69f8e999be > 0 {
		__obf_06503bb8ab96d782 := __obf_d6441e69f8e999be >> __obf_8f1bcc26019920eb
		__obf_d6441e69f8e999be &= __obf_06b5f862243c1836
		if __obf_83094a3f8cd288f0 < len(__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784) {
			__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_83094a3f8cd288f0] = byte(__obf_06503bb8ab96d782 + '0')
			__obf_83094a3f8cd288f0++
		} else if __obf_06503bb8ab96d782 > 0 {
			__obf_c09a2b51ef153b53.__obf_0b587bdd951615f7 = true
		}
		__obf_d6441e69f8e999be = __obf_d6441e69f8e999be * 10
	}
	__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad = __obf_83094a3f8cd288f0

	// Cheat sheet for left shift: table indexed by shift count giving
	// number of new digits that will be introduced by that shift.
	//
	// For example, leftcheats[4] = {2, "625"}.  That means that
	// if we are shifting by 4 (multiplying by 16), it will add 2 digits
	// when the string prefix is "625" through "999", and one fewer digit
	// if the string prefix is "000" through "624".
	//
	// Credit for this trick goes to Ken.
	__obf_31d3f5cacc8faba5(__obf_c09a2b51ef153b53)
}

type __obf_4026db5e3b742331 struct {
	__obf_e5a54521b4bee072 int
	__obf_c7cdd8a8bd7164e4 string // number of new digits
	// minus one digit if original < a.
}

var __obf_105383f0a7aaa6f9 = []__obf_4026db5e3b742331{
	// Leading digits of 1/2^i = 5^i.
	// 5^23 is not an exact 64-bit floating point number,
	// so have to use bc for the math.
	// Go up to 60 to be large enough for 32bit and 64bit platforms.
	/*
		seq 60 | sed 's/^/5^/' | bc |
		awk 'BEGIN{ print "\t{ 0, \"\" }," }
		{
			log2 = log(2)/log(10)
			printf("\t{ %d, \"%s\" },\t// * %d\n",
				int(log2*NR+1), $0, 2**NR)
		}'
	*/
	{0, ""},
	{1, "5"},                                           // * 2
	{1, "25"},                                          // * 4
	{1, "125"},                                         // * 8
	{2, "625"},                                         // * 16
	{2, "3125"},                                        // * 32
	{2, "15625"},                                       // * 64
	{3, "78125"},                                       // * 128
	{3, "390625"},                                      // * 256
	{3, "1953125"},                                     // * 512
	{4, "9765625"},                                     // * 1024
	{4, "48828125"},                                    // * 2048
	{4, "244140625"},                                   // * 4096
	{4, "1220703125"},                                  // * 8192
	{5, "6103515625"},                                  // * 16384
	{5, "30517578125"},                                 // * 32768
	{5, "152587890625"},                                // * 65536
	{6, "762939453125"},                                // * 131072
	{6, "3814697265625"},                               // * 262144
	{6, "19073486328125"},                              // * 524288
	{7, "95367431640625"},                              // * 1048576
	{7, "476837158203125"},                             // * 2097152
	{7, "2384185791015625"},                            // * 4194304
	{7, "11920928955078125"},                           // * 8388608
	{8, "59604644775390625"},                           // * 16777216
	{8, "298023223876953125"},                          // * 33554432
	{8, "1490116119384765625"},                         // * 67108864
	{9, "7450580596923828125"},                         // * 134217728
	{9, "37252902984619140625"},                        // * 268435456
	{9, "186264514923095703125"},                       // * 536870912
	{10, "931322574615478515625"},                      // * 1073741824
	{10, "4656612873077392578125"},                     // * 2147483648
	{10, "23283064365386962890625"},                    // * 4294967296
	{10, "116415321826934814453125"},                   // * 8589934592
	{11, "582076609134674072265625"},                   // * 17179869184
	{11, "2910383045673370361328125"},                  // * 34359738368
	{11, "14551915228366851806640625"},                 // * 68719476736
	{12, "72759576141834259033203125"},                 // * 137438953472
	{12, "363797880709171295166015625"},                // * 274877906944
	{12, "1818989403545856475830078125"},               // * 549755813888
	{13, "9094947017729282379150390625"},               // * 1099511627776
	{13, "45474735088646411895751953125"},              // * 2199023255552
	{13, "227373675443232059478759765625"},             // * 4398046511104
	{13, "1136868377216160297393798828125"},            // * 8796093022208
	{14, "5684341886080801486968994140625"},            // * 17592186044416
	{14, "28421709430404007434844970703125"},           // * 35184372088832
	{14, "142108547152020037174224853515625"},          // * 70368744177664
	{15, "710542735760100185871124267578125"},          // * 140737488355328
	{15, "3552713678800500929355621337890625"},         // * 281474976710656
	{15, "17763568394002504646778106689453125"},        // * 562949953421312
	{16, "88817841970012523233890533447265625"},        // * 1125899906842624
	{16, "444089209850062616169452667236328125"},       // * 2251799813685248
	{16, "2220446049250313080847263336181640625"},      // * 4503599627370496
	{16, "11102230246251565404236316680908203125"},     // * 9007199254740992
	{17, "55511151231257827021181583404541015625"},     // * 18014398509481984
	{17, "277555756156289135105907917022705078125"},    // * 36028797018963968
	{17, "1387778780781445675529539585113525390625"},   // * 72057594037927936
	{18, "6938893903907228377647697925567626953125"},   // * 144115188075855872
	{18, "34694469519536141888238489627838134765625"},  // * 288230376151711744
	{18, "173472347597680709441192448139190673828125"}, // * 576460752303423488
	{19, "867361737988403547205962240695953369140625"}, // * 1152921504606846976
}

// Is the leading prefix of b lexicographically less than s?
func __obf_b1925d22af4aeed3(__obf_3a574d7202216d22 []byte, __obf_8becf347d3c74472 string) bool {
	for __obf_03985402d384eb80 := 0; __obf_03985402d384eb80 < len(__obf_8becf347d3c74472); __obf_03985402d384eb80++ {
		if __obf_03985402d384eb80 >= len(__obf_3a574d7202216d22) {
			return true
		}
		if __obf_3a574d7202216d22[__obf_03985402d384eb80] != __obf_8becf347d3c74472[__obf_03985402d384eb80] {
			return __obf_3a574d7202216d22[__obf_03985402d384eb80] < __obf_8becf347d3c74472[__obf_03985402d384eb80]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_f82bba3e85a33798(__obf_c09a2b51ef153b53 *__obf_4cc30aac3ca67486, __obf_8f1bcc26019920eb uint) {
	__obf_e5a54521b4bee072 := __obf_105383f0a7aaa6f9[__obf_8f1bcc26019920eb].__obf_e5a54521b4bee072
	if __obf_b1925d22af4aeed3(__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[0:__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad], __obf_105383f0a7aaa6f9[__obf_8f1bcc26019920eb].__obf_c7cdd8a8bd7164e4) {
		__obf_e5a54521b4bee072--
	}
	__obf_5039aeab743c839a := __obf_c09a2b51ef153b53. // read index
								__obf_5c82a3d658ffb3ad
	__obf_83094a3f8cd288f0 := __obf_c09a2b51ef153b53. // write index
								__obf_5c82a3d658ffb3ad + __obf_e5a54521b4bee072

	// Pick up a digit, put down a digit.
	var __obf_d6441e69f8e999be uint
	for __obf_5039aeab743c839a--; __obf_5039aeab743c839a >= 0; __obf_5039aeab743c839a-- {
		__obf_d6441e69f8e999be += (uint(__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_5039aeab743c839a]) - '0') << __obf_8f1bcc26019920eb
		__obf_3eba6f9658764c89 := __obf_d6441e69f8e999be / 10
		__obf_e4cf45b477e444a7 := __obf_d6441e69f8e999be - 10*__obf_3eba6f9658764c89
		__obf_83094a3f8cd288f0--
		if __obf_83094a3f8cd288f0 < len(__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784) {
			__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_83094a3f8cd288f0] = byte(__obf_e4cf45b477e444a7 + '0')
		} else if __obf_e4cf45b477e444a7 != 0 {
			__obf_c09a2b51ef153b53.__obf_0b587bdd951615f7 = true
		}
		__obf_d6441e69f8e999be = __obf_3eba6f9658764c89
	}

	// Put down extra digits.
	for __obf_d6441e69f8e999be > 0 {
		__obf_3eba6f9658764c89 := __obf_d6441e69f8e999be / 10
		__obf_e4cf45b477e444a7 := __obf_d6441e69f8e999be - 10*__obf_3eba6f9658764c89
		__obf_83094a3f8cd288f0--
		if __obf_83094a3f8cd288f0 < len(__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784) {
			__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_83094a3f8cd288f0] = byte(__obf_e4cf45b477e444a7 + '0')
		} else if __obf_e4cf45b477e444a7 != 0 {
			__obf_c09a2b51ef153b53.__obf_0b587bdd951615f7 = true
		}
		__obf_d6441e69f8e999be = __obf_3eba6f9658764c89
	}
	__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad += __obf_e5a54521b4bee072
	if __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad >= len(__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784) {
		__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad = len(__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784)
	}
	__obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe += __obf_e5a54521b4bee072

	// Binary shift left (k > 0) or right (k < 0).
	__obf_31d3f5cacc8faba5(__obf_c09a2b51ef153b53)
}

func (__obf_c09a2b51ef153b53 *__obf_4cc30aac3ca67486) Shift(__obf_8f1bcc26019920eb int) {
	switch {
	case __obf_c09a2b51ef153b53.
		// nothing to do: a == 0
		__obf_5c82a3d658ffb3ad == 0:

	case __obf_8f1bcc26019920eb > 0:
		for __obf_8f1bcc26019920eb > __obf_a399770894bf00d0 {
			__obf_f82bba3e85a33798(__obf_c09a2b51ef153b53, __obf_a399770894bf00d0)
			__obf_8f1bcc26019920eb -= __obf_a399770894bf00d0
		}
		__obf_f82bba3e85a33798(__obf_c09a2b51ef153b53, uint(__obf_8f1bcc26019920eb))
	case __obf_8f1bcc26019920eb < 0:
		for __obf_8f1bcc26019920eb < -__obf_a399770894bf00d0 {
			__obf_f04759d2e55de434(__obf_c09a2b51ef153b53, __obf_a399770894bf00d0)
			__obf_8f1bcc26019920eb += __obf_a399770894bf00d0
		}
		__obf_f04759d2e55de434(__obf_c09a2b51ef153b53, uint(-__obf_8f1bcc26019920eb))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_5d1c750d31406111(__obf_c09a2b51ef153b53 *__obf_4cc30aac3ca67486, __obf_5c82a3d658ffb3ad int) bool {
	if __obf_5c82a3d658ffb3ad < 0 || __obf_5c82a3d658ffb3ad >= __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad {
		return false
	}
	if __obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_5c82a3d658ffb3ad] == '5' && __obf_5c82a3d658ffb3ad+1 == __obf_c09a2b51ef153b53. // exactly halfway - round to even
																		__obf_5c82a3d658ffb3ad {
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_c09a2b51ef153b53.__obf_0b587bdd951615f7 {
			return true
		}
		return __obf_5c82a3d658ffb3ad > 0 && (__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_5c82a3d658ffb3ad-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_c09a2b51ef153b53.

		// Round a to nd digits (or fewer).
		// If nd is zero, it means we're rounding
		// just to the left of the digits, as in
		// 0.09 -> 0.1.
		__obf_aafc5e233dfea784[__obf_5c82a3d658ffb3ad] >= '5'
}

func (__obf_c09a2b51ef153b53 *__obf_4cc30aac3ca67486) Round(__obf_5c82a3d658ffb3ad int) {
	if __obf_5c82a3d658ffb3ad < 0 || __obf_5c82a3d658ffb3ad >= __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad {
		return
	}
	if __obf_5d1c750d31406111(__obf_c09a2b51ef153b53, __obf_5c82a3d658ffb3ad) {
		__obf_c09a2b51ef153b53.
			RoundUp(__obf_5c82a3d658ffb3ad)
	} else {
		__obf_c09a2b51ef153b53.
			RoundDown(__obf_5c82a3d658ffb3ad)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_c09a2b51ef153b53 *__obf_4cc30aac3ca67486) RoundDown(__obf_5c82a3d658ffb3ad int) {
	if __obf_5c82a3d658ffb3ad < 0 || __obf_5c82a3d658ffb3ad >= __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad {
		return
	}
	__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad = __obf_5c82a3d658ffb3ad

	// Round a up to nd digits (or fewer).
	__obf_31d3f5cacc8faba5(__obf_c09a2b51ef153b53)
}

func (__obf_c09a2b51ef153b53 *__obf_4cc30aac3ca67486) RoundUp(__obf_5c82a3d658ffb3ad int) {
	if __obf_5c82a3d658ffb3ad < 0 || __obf_5c82a3d658ffb3ad >= __obf_c09a2b51ef153b53.

		// round up
		__obf_5c82a3d658ffb3ad {
		return
	}

	for __obf_03985402d384eb80 := __obf_5c82a3d658ffb3ad - 1; __obf_03985402d384eb80 >= 0; __obf_03985402d384eb80-- {
		__obf_d520aaab80df92cd := __obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_03985402d384eb80]
		if __obf_d520aaab80df92cd < '9' {
			__obf_c09a2b51ef153b53. // can stop after this digit
						__obf_aafc5e233dfea784[__obf_03985402d384eb80]++
			__obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad = __obf_03985402d384eb80 + 1
			return
		}
	}
	__obf_c09a2b51ef153b53.

		// Number is all 9s.
		// Change to single 1 with adjusted decimal point.
		__obf_aafc5e233dfea784[0] = '1'
	__obf_c09a2b51ef153b53.

		// Extract integer part, rounded appropriately.
		// No guarantees about overflow.
		__obf_5c82a3d658ffb3ad = 1
	__obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe++
}

func (__obf_c09a2b51ef153b53 *__obf_4cc30aac3ca67486) RoundedInteger() uint64 {
	if __obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_03985402d384eb80 int
	__obf_d6441e69f8e999be := uint64(0)
	for __obf_03985402d384eb80 = 0; __obf_03985402d384eb80 < __obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe && __obf_03985402d384eb80 < __obf_c09a2b51ef153b53.__obf_5c82a3d658ffb3ad; __obf_03985402d384eb80++ {
		__obf_d6441e69f8e999be = __obf_d6441e69f8e999be*10 + uint64(__obf_c09a2b51ef153b53.__obf_aafc5e233dfea784[__obf_03985402d384eb80]-'0')
	}
	for ; __obf_03985402d384eb80 < __obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe; __obf_03985402d384eb80++ {
		__obf_d6441e69f8e999be *= 10
	}
	if __obf_5d1c750d31406111(__obf_c09a2b51ef153b53, __obf_c09a2b51ef153b53.__obf_2b78448fb2f1bffe) {
		__obf_d6441e69f8e999be++
	}
	return __obf_d6441e69f8e999be
}
