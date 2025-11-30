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
package __obf_b0e21776b9d45e13

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

var __obf_ec129f5235aa8458 = big.NewInt(0)
var __obf_c8643a8d2b2d6a40 = big.NewInt(1)
var __obf_f741b5f088b7b4b7 = big.NewInt(2)
var __obf_55b28070f37dbd50 = big.NewInt(4)
var __obf_4374a770c15c9bd3 = big.NewInt(5)
var __obf_d4d5e0f4e5fc42f5 = big.NewInt(10)
var __obf_9195a8e64ca7e21a = big.NewInt(20)

var __obf_a7f8febfb20b3c1e = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_8f559e1327c9ee84 *big.Int
	__obf_97e810ed39f8ac3f int32 // NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.

}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_8f559e1327c9ee84 int64, __obf_97e810ed39f8ac3f int32) Decimal {
	return Decimal{__obf_8f559e1327c9ee84: big.NewInt(__obf_8f559e1327c9ee84), __obf_97e810ed39f8ac3f: __obf_97e810ed39f8ac3f}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_8f559e1327c9ee84 int64) Decimal {
	return Decimal{__obf_8f559e1327c9ee84: big.NewInt(__obf_8f559e1327c9ee84), __obf_97e810ed39f8ac3f: 0}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_8f559e1327c9ee84 int32) Decimal {
	return Decimal{__obf_8f559e1327c9ee84: big.NewInt(int64(__obf_8f559e1327c9ee84)), __obf_97e810ed39f8ac3f: 0}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_8f559e1327c9ee84 uint64) Decimal {
	return Decimal{__obf_8f559e1327c9ee84: new(big.Int).SetUint64(__obf_8f559e1327c9ee84), __obf_97e810ed39f8ac3f: 0}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_8f559e1327c9ee84 *big.Int, __obf_97e810ed39f8ac3f int32) Decimal {
	return Decimal{__obf_8f559e1327c9ee84: new(big.Int).Set(__obf_8f559e1327c9ee84), __obf_97e810ed39f8ac3f: __obf_97e810ed39f8ac3f}
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
func NewFromBigRat(__obf_8f559e1327c9ee84 *big.Rat, __obf_e69ecd4c32f1420b int32) Decimal {
	return Decimal{__obf_8f559e1327c9ee84: new(big.Int).Set(__obf_8f559e1327c9ee84.Num()), __obf_97e810ed39f8ac3f: 0}.DivRound(Decimal{__obf_8f559e1327c9ee84: new(big.Int).Set(__obf_8f559e1327c9ee84.Denom()), __obf_97e810ed39f8ac3f: 0}, __obf_e69ecd4c32f1420b)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_8f559e1327c9ee84 string) (Decimal, error) {
	__obf_3914498306978f51 := __obf_8f559e1327c9ee84
	var __obf_db8076f917984df3 string
	var __obf_97e810ed39f8ac3f int64
	__obf_ff164edb3bedf438 := // Check if number is using scientific notation
		strings.IndexAny(__obf_8f559e1327c9ee84, "Ee")
	if __obf_ff164edb3bedf438 != -1 {
		__obf_9b5105062f538c41, __obf_b2a10ebf0533880b := strconv.ParseInt(__obf_8f559e1327c9ee84[__obf_ff164edb3bedf438+1:], 10, 32)
		if __obf_b2a10ebf0533880b != nil {
			if __obf_7321ee49f99531d7, __obf_e9230491278a66dd := __obf_b2a10ebf0533880b.(*strconv.NumError); __obf_e9230491278a66dd && __obf_7321ee49f99531d7.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_8f559e1327c9ee84)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_8f559e1327c9ee84)
		}
		__obf_8f559e1327c9ee84 = __obf_8f559e1327c9ee84[:__obf_ff164edb3bedf438]
		__obf_97e810ed39f8ac3f = __obf_9b5105062f538c41
	}
	__obf_6f711b95cae72c21 := -1
	__obf_bac7005559143f05 := len(__obf_8f559e1327c9ee84)
	for __obf_0038051e77a34ba0 := 0; __obf_0038051e77a34ba0 < __obf_bac7005559143f05; __obf_0038051e77a34ba0++ {
		if __obf_8f559e1327c9ee84[__obf_0038051e77a34ba0] == '.' {
			if __obf_6f711b95cae72c21 > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_8f559e1327c9ee84)
			}
			__obf_6f711b95cae72c21 = __obf_0038051e77a34ba0
		}
	}

	if __obf_6f711b95cae72c21 == -1 {
		__obf_db8076f917984df3 = // There is no decimal point, we can just parse the original string as
			// an int
			__obf_8f559e1327c9ee84
	} else {
		if __obf_6f711b95cae72c21+1 < __obf_bac7005559143f05 {
			__obf_db8076f917984df3 = __obf_8f559e1327c9ee84[:__obf_6f711b95cae72c21] + __obf_8f559e1327c9ee84[__obf_6f711b95cae72c21+1:]
		} else {
			__obf_db8076f917984df3 = __obf_8f559e1327c9ee84[:__obf_6f711b95cae72c21]
		}
		__obf_9b5105062f538c41 := -len(__obf_8f559e1327c9ee84[__obf_6f711b95cae72c21+1:])
		__obf_97e810ed39f8ac3f += int64(__obf_9b5105062f538c41)
	}

	var __obf_0e957e65aadc5bd2 *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_db8076f917984df3) <= 18 {
		__obf_e80989271d934117, __obf_b2a10ebf0533880b := strconv.ParseInt(__obf_db8076f917984df3, 10, 64)
		if __obf_b2a10ebf0533880b != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_8f559e1327c9ee84)
		}
		__obf_0e957e65aadc5bd2 = big.NewInt(__obf_e80989271d934117)
	} else {
		__obf_0e957e65aadc5bd2 = new(big.Int)
		_, __obf_e9230491278a66dd := __obf_0e957e65aadc5bd2.SetString(__obf_db8076f917984df3, 10)
		if !__obf_e9230491278a66dd {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_8f559e1327c9ee84)
		}
	}

	if __obf_97e810ed39f8ac3f < math.MinInt32 || __obf_97e810ed39f8ac3f > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_3914498306978f51)
	}

	return Decimal{__obf_8f559e1327c9ee84: __obf_0e957e65aadc5bd2, __obf_97e810ed39f8ac3f: int32(__obf_97e810ed39f8ac3f)}, nil
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
func NewFromFormattedString(__obf_8f559e1327c9ee84 string, __obf_df66916210c70cec *regexp.Regexp) (Decimal, error) {
	__obf_aa401f44bdd3c143 := __obf_df66916210c70cec.ReplaceAllString(__obf_8f559e1327c9ee84, "")
	__obf_df803c2772cc0386, __obf_b2a10ebf0533880b := NewFromString(__obf_aa401f44bdd3c143)
	if __obf_b2a10ebf0533880b != nil {
		return Decimal{}, __obf_b2a10ebf0533880b
	}
	return __obf_df803c2772cc0386, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_8f559e1327c9ee84 string) Decimal {
	__obf_428f69476005da25, __obf_b2a10ebf0533880b := NewFromString(__obf_8f559e1327c9ee84)
	if __obf_b2a10ebf0533880b != nil {
		panic(__obf_b2a10ebf0533880b)
	}
	return __obf_428f69476005da25
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
func NewFromFloat(__obf_8f559e1327c9ee84 float64) Decimal {
	if __obf_8f559e1327c9ee84 == 0 {
		return New(0, 0)
	}
	return __obf_a46e0fbfcc7309f8(__obf_8f559e1327c9ee84, math.Float64bits(__obf_8f559e1327c9ee84), &__obf_0e52bd84f2e0c5e5)
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
func NewFromFloat32(__obf_8f559e1327c9ee84 float32) Decimal {
	if __obf_8f559e1327c9ee84 == 0 {
		return New(0, 0)
	}
	__obf_ca960a67e11dc734 := // XOR is workaround for https://github.com/golang/go/issues/26285
		math.Float32bits(__obf_8f559e1327c9ee84) ^ 0x80808080
	return __obf_a46e0fbfcc7309f8(float64(__obf_8f559e1327c9ee84), uint64(__obf_ca960a67e11dc734)^0x80808080, &__obf_2ea0d484844c6331)
}

func __obf_a46e0fbfcc7309f8(__obf_b2485bbdce5d739f float64, __obf_a5095c98ff2818d7 uint64, __obf_c19e228d4278835d *__obf_7361d3b42dd22c09) Decimal {
	if math.IsNaN(__obf_b2485bbdce5d739f) || math.IsInf(__obf_b2485bbdce5d739f, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_b2485bbdce5d739f))
	}
	__obf_97e810ed39f8ac3f := int(__obf_a5095c98ff2818d7>>__obf_c19e228d4278835d.__obf_2f2d498924addd14) & (1<<__obf_c19e228d4278835d.__obf_96af297d4e88f620 - 1)
	__obf_22ce9de7558a3141 := __obf_a5095c98ff2818d7 & (uint64(1)<<__obf_c19e228d4278835d.__obf_2f2d498924addd14 - 1)

	switch __obf_97e810ed39f8ac3f {
	case 0:
		__obf_97e810ed39f8ac3f++ // denormalized

	default:
		__obf_22ce9de7558a3141 |= // add implicit top bit
			uint64(1) << __obf_c19e228d4278835d.__obf_2f2d498924addd14
	}
	__obf_97e810ed39f8ac3f += __obf_c19e228d4278835d.__obf_0628d944dfc724b5

	var __obf_df803c2772cc0386 __obf_b0e21776b9d45e13
	__obf_df803c2772cc0386.
		Assign(__obf_22ce9de7558a3141)
	__obf_df803c2772cc0386.
		Shift(__obf_97e810ed39f8ac3f - int(__obf_c19e228d4278835d.__obf_2f2d498924addd14))
	__obf_df803c2772cc0386.__obf_d3daa27031f7f18b = __obf_a5095c98ff2818d7>>(__obf_c19e228d4278835d.__obf_96af297d4e88f620+__obf_c19e228d4278835d.__obf_2f2d498924addd14) != 0
	__obf_96ba681e9a854a03(&__obf_df803c2772cc0386,
		// If less than 19 digits, we can do calculation in an int64.
		__obf_22ce9de7558a3141, __obf_97e810ed39f8ac3f, __obf_c19e228d4278835d)

	if __obf_df803c2772cc0386.__obf_2841543582033a23 < 19 {
		__obf_ba9c7603fdcb2493 := int64(0)
		__obf_cb9fa5e90307ec38 := int64(1)
		for __obf_0038051e77a34ba0 := __obf_df803c2772cc0386.__obf_2841543582033a23 - 1; __obf_0038051e77a34ba0 >= 0; __obf_0038051e77a34ba0-- {
			__obf_ba9c7603fdcb2493 += __obf_cb9fa5e90307ec38 * int64(__obf_df803c2772cc0386.__obf_df803c2772cc0386[__obf_0038051e77a34ba0]-'0')
			__obf_cb9fa5e90307ec38 *= 10
		}
		if __obf_df803c2772cc0386.__obf_d3daa27031f7f18b {
			__obf_ba9c7603fdcb2493 *= -1
		}
		return Decimal{__obf_8f559e1327c9ee84: big.NewInt(__obf_ba9c7603fdcb2493), __obf_97e810ed39f8ac3f: int32(__obf_df803c2772cc0386.__obf_b4c2d4c5eeb9a307) - int32(__obf_df803c2772cc0386.__obf_2841543582033a23)}
	}
	__obf_0e957e65aadc5bd2 := new(big.Int)
	__obf_0e957e65aadc5bd2, __obf_e9230491278a66dd := __obf_0e957e65aadc5bd2.SetString(string(__obf_df803c2772cc0386.__obf_df803c2772cc0386[:__obf_df803c2772cc0386.__obf_2841543582033a23]), 10)
	if __obf_e9230491278a66dd {
		return Decimal{__obf_8f559e1327c9ee84: __obf_0e957e65aadc5bd2, __obf_97e810ed39f8ac3f: int32(__obf_df803c2772cc0386.__obf_b4c2d4c5eeb9a307) - int32(__obf_df803c2772cc0386.__obf_2841543582033a23)}
	}

	return NewFromFloatWithExponent(__obf_b2485bbdce5d739f, int32(__obf_df803c2772cc0386.

		// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
		// number of fractional digits.
		//
		// Example:
		//
		//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
		__obf_b4c2d4c5eeb9a307)-int32(__obf_df803c2772cc0386.__obf_2841543582033a23))
}

func NewFromFloatWithExponent(__obf_8f559e1327c9ee84 float64, __obf_97e810ed39f8ac3f int32) Decimal {
	if math.IsNaN(__obf_8f559e1327c9ee84) || math.IsInf(__obf_8f559e1327c9ee84, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_8f559e1327c9ee84))
	}
	__obf_a5095c98ff2818d7 := math.Float64bits(__obf_8f559e1327c9ee84)
	__obf_22ce9de7558a3141 := __obf_a5095c98ff2818d7 & (1<<52 - 1)
	__obf_54a7787674fe4d4b := int32((__obf_a5095c98ff2818d7 >> 52) & (1<<11 - 1))
	__obf_cf03f6d30f3621c5 := __obf_a5095c98ff2818d7 >> 63

	if __obf_54a7787674fe4d4b == 0 {
		// specials
		if __obf_22ce9de7558a3141 == 0 {
			return Decimal{}
		}
		__obf_54a7787674fe4d4b++ // subnormal

	} else {
		__obf_22ce9de7558a3141 |= // normal
			1 << 52
	}
	__obf_54a7787674fe4d4b -= 1023 + 52

	// normalizing base-2 values
	for __obf_22ce9de7558a3141&1 == 0 {
		__obf_22ce9de7558a3141 = __obf_22ce9de7558a3141 >> 1
		__obf_54a7787674fe4d4b++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_97e810ed39f8ac3f < 0 && __obf_97e810ed39f8ac3f < __obf_54a7787674fe4d4b {
		if __obf_54a7787674fe4d4b < 0 {
			__obf_97e810ed39f8ac3f = __obf_54a7787674fe4d4b
		} else {
			__obf_97e810ed39f8ac3f = 0
		}
	}
	__obf_54a7787674fe4d4b -= // representing 10^M * 2^N as 5^M * 2^(M+N)
		__obf_97e810ed39f8ac3f
	__obf_39d906e3a3077b25 := big.NewInt(1)
	__obf_fface1e9292aa9c2 := big.NewInt(int64(__obf_22ce9de7558a3141))

	// applying 5^M
	if __obf_97e810ed39f8ac3f > 0 {
		__obf_39d906e3a3077b25 = __obf_39d906e3a3077b25.SetInt64(int64(__obf_97e810ed39f8ac3f))
		__obf_39d906e3a3077b25 = __obf_39d906e3a3077b25.Exp(__obf_4374a770c15c9bd3, __obf_39d906e3a3077b25, nil)
	} else if __obf_97e810ed39f8ac3f < 0 {
		__obf_39d906e3a3077b25 = __obf_39d906e3a3077b25.SetInt64(-int64(__obf_97e810ed39f8ac3f))
		__obf_39d906e3a3077b25 = __obf_39d906e3a3077b25.Exp(__obf_4374a770c15c9bd3, __obf_39d906e3a3077b25, nil)
		__obf_fface1e9292aa9c2 = __obf_fface1e9292aa9c2.Mul(__obf_fface1e9292aa9c2, __obf_39d906e3a3077b25)
		__obf_39d906e3a3077b25 = __obf_39d906e3a3077b25.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_54a7787674fe4d4b > 0 {
		__obf_fface1e9292aa9c2 = __obf_fface1e9292aa9c2.Lsh(__obf_fface1e9292aa9c2, uint(__obf_54a7787674fe4d4b))
	} else if __obf_54a7787674fe4d4b < 0 {
		__obf_39d906e3a3077b25 = __obf_39d906e3a3077b25.Lsh(__obf_39d906e3a3077b25, uint(-__obf_54a7787674fe4d4b))
	}

	// rounding and downscaling
	if __obf_97e810ed39f8ac3f > 0 || __obf_54a7787674fe4d4b < 0 {
		__obf_e9a18fd092cf6d14 := new(big.Int).Rsh(__obf_39d906e3a3077b25, 1)
		__obf_fface1e9292aa9c2 = __obf_fface1e9292aa9c2.Add(__obf_fface1e9292aa9c2, __obf_e9a18fd092cf6d14)
		__obf_fface1e9292aa9c2 = __obf_fface1e9292aa9c2.Quo(__obf_fface1e9292aa9c2, __obf_39d906e3a3077b25)
	}

	if __obf_cf03f6d30f3621c5 == 1 {
		__obf_fface1e9292aa9c2 = __obf_fface1e9292aa9c2.Neg(__obf_fface1e9292aa9c2)
	}

	return Decimal{__obf_8f559e1327c9ee84: __obf_fface1e9292aa9c2, __obf_97e810ed39f8ac3f: __obf_97e810ed39f8ac3f}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_df803c2772cc0386 Decimal) Copy() Decimal {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()
	return Decimal{__obf_8f559e1327c9ee84: new(big.Int).Set(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84), __obf_97e810ed39f8ac3f: __obf_df803c2772cc0386.

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
		__obf_97e810ed39f8ac3f,
	}
}

func (__obf_df803c2772cc0386 Decimal) __obf_749ebd7071bc390b(__obf_97e810ed39f8ac3f int32) Decimal {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()

	if __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f == __obf_97e810ed39f8ac3f {
		return Decimal{
			new(big.Int).Set(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84), __obf_df803c2772cc0386.

				// NOTE(vadim): must convert exps to float64 before - to prevent overflow
				__obf_97e810ed39f8ac3f,
		}
	}
	__obf_bc4ea5161de22c13 := math.Abs(float64(__obf_97e810ed39f8ac3f) - float64(__obf_df803c2772cc0386.__obf_97e810ed39f8ac3f))
	__obf_8f559e1327c9ee84 := new(big.Int).Set(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84)
	__obf_802ffb6d6f8751be := new(big.Int).Exp(__obf_d4d5e0f4e5fc42f5, big.NewInt(int64(__obf_bc4ea5161de22c13)), nil)
	if __obf_97e810ed39f8ac3f > __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f {
		__obf_8f559e1327c9ee84 = __obf_8f559e1327c9ee84.Quo(__obf_8f559e1327c9ee84, __obf_802ffb6d6f8751be)
	} else if __obf_97e810ed39f8ac3f < __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f {
		__obf_8f559e1327c9ee84 = __obf_8f559e1327c9ee84.Mul(__obf_8f559e1327c9ee84, __obf_802ffb6d6f8751be)
	}

	return Decimal{__obf_8f559e1327c9ee84: __obf_8f559e1327c9ee84, __obf_97e810ed39f8ac3f: __obf_97e810ed39f8ac3f}
}

// Abs returns the absolute value of the decimal.
func (__obf_df803c2772cc0386 Decimal) Abs() Decimal {
	if !__obf_df803c2772cc0386.IsNegative() {
		return __obf_df803c2772cc0386
	}
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()
	__obf_0aadda5c120d0258 := new(big.Int).Abs(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84)
	return Decimal{__obf_8f559e1327c9ee84: __obf_0aadda5c120d0258, __obf_97e810ed39f8ac3f: __obf_df803c2772cc0386.

		// Add returns d + d2.
		__obf_97e810ed39f8ac3f,
	}
}

func (__obf_df803c2772cc0386 Decimal) Add(__obf_b6bc08d52d165029 Decimal) Decimal {
	__obf_ce5f1b4c4db8f0f4, __obf_b9159a371b250a0c := RescalePair(__obf_df803c2772cc0386, __obf_b6bc08d52d165029)
	__obf_0b7bfb7ced19b5b1 := new(big.Int).Add(__obf_ce5f1b4c4db8f0f4.__obf_8f559e1327c9ee84, __obf_b9159a371b250a0c.__obf_8f559e1327c9ee84)
	return Decimal{__obf_8f559e1327c9ee84: __obf_0b7bfb7ced19b5b1, __obf_97e810ed39f8ac3f: __obf_ce5f1b4c4db8f0f4.

		// Sub returns d - d2.
		__obf_97e810ed39f8ac3f,
	}
}

func (__obf_df803c2772cc0386 Decimal) Sub(__obf_b6bc08d52d165029 Decimal) Decimal {
	__obf_ce5f1b4c4db8f0f4, __obf_b9159a371b250a0c := RescalePair(__obf_df803c2772cc0386, __obf_b6bc08d52d165029)
	__obf_0b7bfb7ced19b5b1 := new(big.Int).Sub(__obf_ce5f1b4c4db8f0f4.__obf_8f559e1327c9ee84, __obf_b9159a371b250a0c.__obf_8f559e1327c9ee84)
	return Decimal{__obf_8f559e1327c9ee84: __obf_0b7bfb7ced19b5b1, __obf_97e810ed39f8ac3f: __obf_ce5f1b4c4db8f0f4.

		// Neg returns -d.
		__obf_97e810ed39f8ac3f,
	}
}

func (__obf_df803c2772cc0386 Decimal) Neg() Decimal {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()
	__obf_b2485bbdce5d739f := new(big.Int).Neg(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84)
	return Decimal{__obf_8f559e1327c9ee84: __obf_b2485bbdce5d739f, __obf_97e810ed39f8ac3f: __obf_df803c2772cc0386.

		// Mul returns d * d2.
		__obf_97e810ed39f8ac3f,
	}
}

func (__obf_df803c2772cc0386 Decimal) Mul(__obf_b6bc08d52d165029 Decimal) Decimal {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()
	__obf_b6bc08d52d165029.__obf_a6a609bb966fabfb()
	__obf_45b1ec27ef5278e2 := int64(__obf_df803c2772cc0386.__obf_97e810ed39f8ac3f) + int64(__obf_b6bc08d52d165029.__obf_97e810ed39f8ac3f)
	if __obf_45b1ec27ef5278e2 > math.MaxInt32 || __obf_45b1ec27ef5278e2 < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_45b1ec27ef5278e2))
	}
	__obf_0b7bfb7ced19b5b1 := new(big.Int).Mul(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84, __obf_b6bc08d52d165029.__obf_8f559e1327c9ee84)
	return Decimal{__obf_8f559e1327c9ee84: __obf_0b7bfb7ced19b5b1, __obf_97e810ed39f8ac3f: int32(__obf_45b1ec27ef5278e2)}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_df803c2772cc0386 Decimal) Shift(__obf_2d852567572af160 int32) Decimal {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()
	return Decimal{__obf_8f559e1327c9ee84: new(big.Int).Set(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84), __obf_97e810ed39f8ac3f: __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f + __obf_2d852567572af160}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_df803c2772cc0386 Decimal) Div(__obf_b6bc08d52d165029 Decimal) Decimal {
	return __obf_df803c2772cc0386.DivRound(__obf_b6bc08d52d165029, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_df803c2772cc0386 Decimal) QuoRem(__obf_b6bc08d52d165029 Decimal, __obf_e69ecd4c32f1420b int32) (Decimal, Decimal) {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()
	__obf_b6bc08d52d165029.__obf_a6a609bb966fabfb()
	if __obf_b6bc08d52d165029.__obf_8f559e1327c9ee84.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_3d471e082178e843 := -__obf_e69ecd4c32f1420b
	__obf_7321ee49f99531d7 := int64(__obf_df803c2772cc0386.__obf_97e810ed39f8ac3f) - int64(__obf_b6bc08d52d165029.__obf_97e810ed39f8ac3f) - int64(__obf_3d471e082178e843)
	if __obf_7321ee49f99531d7 > math.MaxInt32 || __obf_7321ee49f99531d7 < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_ac5ce983b98c0896, __obf_531bef95e5daa7e0,

		// d = a 10^ea
		// d2 = b 10^eb
		__obf_ee2534f9678f0e09 big.Int
	var __obf_cedf36ec853d2333 int32

	if __obf_7321ee49f99531d7 < 0 {
		__obf_ac5ce983b98c0896 = *__obf_df803c2772cc0386.__obf_8f559e1327c9ee84
		__obf_ee2534f9678f0e09.
			SetInt64(-__obf_7321ee49f99531d7)
		__obf_531bef95e5daa7e0.
			Exp(__obf_d4d5e0f4e5fc42f5, &__obf_ee2534f9678f0e09, nil)
		__obf_531bef95e5daa7e0.
			Mul(__obf_b6bc08d52d165029.__obf_8f559e1327c9ee84, &__obf_531bef95e5daa7e0)
		__obf_cedf36ec853d2333 = __obf_df803c2772cc0386.
			// now aa = a
			//     bb = b 10^(scale + eb - ea)
			__obf_97e810ed39f8ac3f

	} else {
		__obf_ee2534f9678f0e09.
			SetInt64(__obf_7321ee49f99531d7)
		__obf_ac5ce983b98c0896.
			Exp(__obf_d4d5e0f4e5fc42f5, &__obf_ee2534f9678f0e09, nil)
		__obf_ac5ce983b98c0896.
			Mul(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84, &__obf_ac5ce983b98c0896)
		__obf_531bef95e5daa7e0 = *__obf_b6bc08d52d165029.__obf_8f559e1327c9ee84

		// now aa = a ^ (ea - eb - scale)
		//     bb = b
		__obf_cedf36ec853d2333 = __obf_3d471e082178e843 + __obf_b6bc08d52d165029.__obf_97e810ed39f8ac3f

	}
	var __obf_2a31e4af66195793, __obf_69a72b316864499d big.Int
	__obf_2a31e4af66195793.
		QuoRem(&__obf_ac5ce983b98c0896, &__obf_531bef95e5daa7e0, &__obf_69a72b316864499d)
	__obf_a3d3511e97fd2f6b := Decimal{__obf_8f559e1327c9ee84: &__obf_2a31e4af66195793, __obf_97e810ed39f8ac3f: __obf_3d471e082178e843}
	__obf_c276dcd472891965 := Decimal{__obf_8f559e1327c9ee84: &__obf_69a72b316864499d, __obf_97e810ed39f8ac3f: __obf_cedf36ec853d2333}
	return __obf_a3d3511e97fd2f6b,

		// DivRound divides and rounds to a given precision
		// i.e. to an integer multiple of 10^(-precision)
		//
		//	for a positive quotient digit 5 is rounded up, away from 0
		//	if the quotient is negative then digit 5 is rounded down, away from 0
		//
		// Note that precision<0 is allowed as input.
		__obf_c276dcd472891965
}

func (__obf_df803c2772cc0386 Decimal) DivRound(__obf_b6bc08d52d165029 Decimal, __obf_e69ecd4c32f1420b int32) Decimal {
	__obf_2a31e4af66195793,
		// QuoRem already checks initialization
		__obf_69a72b316864499d := __obf_df803c2772cc0386.QuoRem(__obf_b6bc08d52d165029,
		// the actual rounding decision is based on comparing r*10^precision and d2/2
		// instead compare 2 r 10 ^precision and d2
		__obf_e69ecd4c32f1420b)

	var __obf_93ea491676fae347 big.Int
	__obf_93ea491676fae347.
		Abs(__obf_69a72b316864499d.__obf_8f559e1327c9ee84)
	__obf_93ea491676fae347.
		Lsh(&__obf_93ea491676fae347, 1)
	__obf_73d7ad7482861598 := // now rv2 = abs(r.value) * 2
		Decimal{__obf_8f559e1327c9ee84: &__obf_93ea491676fae347, __obf_97e810ed39f8ac3f: __obf_69a72b316864499d.
			// r2 is now 2 * r * 10 ^ precision
			__obf_97e810ed39f8ac3f + __obf_e69ecd4c32f1420b}

	var __obf_253219491e4be33b = __obf_73d7ad7482861598.Cmp(__obf_b6bc08d52d165029.Abs())

	if __obf_253219491e4be33b < 0 {
		return __obf_2a31e4af66195793
	}

	if __obf_df803c2772cc0386.__obf_8f559e1327c9ee84.Sign()*__obf_b6bc08d52d165029.__obf_8f559e1327c9ee84.Sign() < 0 {
		return __obf_2a31e4af66195793.Sub(New(1, -__obf_e69ecd4c32f1420b))
	}

	return __obf_2a31e4af66195793.Add(New(1, -__obf_e69ecd4c32f1420b))
}

// Mod returns d % d2.
func (__obf_df803c2772cc0386 Decimal) Mod(__obf_b6bc08d52d165029 Decimal) Decimal {
	_, __obf_69a72b316864499d := __obf_df803c2772cc0386.QuoRem(__obf_b6bc08d52d165029, 0)
	return __obf_69a72b316864499d
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
func (__obf_df803c2772cc0386 Decimal) Pow(__obf_b6bc08d52d165029 Decimal) Decimal {
	__obf_a6c4d873d619640b := __obf_df803c2772cc0386.Sign()
	__obf_52c74f74cb4b4844 := __obf_b6bc08d52d165029.Sign()

	if __obf_a6c4d873d619640b == 0 {
		if __obf_52c74f74cb4b4844 == 0 {
			return Decimal{}
		}
		if __obf_52c74f74cb4b4844 == 1 {
			return Decimal{__obf_ec129f5235aa8458, 0}
		}
		if __obf_52c74f74cb4b4844 == -1 {
			return Decimal{}
		}
	}

	if __obf_52c74f74cb4b4844 == 0 {
		return Decimal{__obf_c8643a8d2b2d6a40, 0}
	}
	__obf_4aae8d20fd4eaf66 := // TODO: optimize extraction of fractional part
		Decimal{__obf_c8643a8d2b2d6a40, 0}
	__obf_ddb81b725bab291a, __obf_6a73025256814cba := __obf_b6bc08d52d165029.QuoRem(__obf_4aae8d20fd4eaf66, 0)

	if __obf_a6c4d873d619640b == -1 && !__obf_6a73025256814cba.IsZero() {
		return Decimal{}
	}
	__obf_dd70e8929c261ac7, _ := __obf_df803c2772cc0386.PowBigInt(__obf_ddb81b725bab291a.

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_8f559e1327c9ee84)

	if __obf_6a73025256814cba.__obf_8f559e1327c9ee84.Sign() == 0 {
		return __obf_dd70e8929c261ac7
	}
	__obf_51c5f2b2b2662b7e := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_df803c2772cc0386.NumDigits()
	__obf_038c585447669ee5 := __obf_b6bc08d52d165029.NumDigits()
	__obf_e69ecd4c32f1420b := __obf_51c5f2b2b2662b7e

	if __obf_038c585447669ee5 > __obf_e69ecd4c32f1420b {
		__obf_e69ecd4c32f1420b += __obf_038c585447669ee5
	}
	__obf_e69ecd4c32f1420b += 6
	__obf_07d089e0aaf4fa6e,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_b2a10ebf0533880b := __obf_df803c2772cc0386.Abs().Ln(-__obf_df803c2772cc0386.__obf_97e810ed39f8ac3f + int32(__obf_e69ecd4c32f1420b))
	if __obf_b2a10ebf0533880b != nil {
		return Decimal{}
	}
	__obf_07d089e0aaf4fa6e = __obf_07d089e0aaf4fa6e.Mul(__obf_6a73025256814cba)
	__obf_07d089e0aaf4fa6e, __obf_b2a10ebf0533880b = __obf_07d089e0aaf4fa6e.ExpTaylor(-__obf_df803c2772cc0386.__obf_97e810ed39f8ac3f + int32(__obf_e69ecd4c32f1420b))
	if __obf_b2a10ebf0533880b != nil {
		return Decimal{}
	}
	__obf_aee88fdbf6e4041b := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_dd70e8929c261ac7.Mul(__obf_07d089e0aaf4fa6e)

	return __obf_aee88fdbf6e4041b
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
func (__obf_df803c2772cc0386 Decimal) PowWithPrecision(__obf_b6bc08d52d165029 Decimal, __obf_e69ecd4c32f1420b int32) (Decimal, error) {
	__obf_a6c4d873d619640b := __obf_df803c2772cc0386.Sign()
	__obf_52c74f74cb4b4844 := __obf_b6bc08d52d165029.Sign()

	if __obf_a6c4d873d619640b == 0 {
		if __obf_52c74f74cb4b4844 == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_52c74f74cb4b4844 == 1 {
			return Decimal{__obf_ec129f5235aa8458, 0}, nil
		}
		if __obf_52c74f74cb4b4844 == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_52c74f74cb4b4844 == 0 {
		return Decimal{__obf_c8643a8d2b2d6a40, 0}, nil
	}
	__obf_4aae8d20fd4eaf66 := // TODO: optimize extraction of fractional part
		Decimal{__obf_c8643a8d2b2d6a40, 0}
	__obf_ddb81b725bab291a, __obf_6a73025256814cba := __obf_b6bc08d52d165029.QuoRem(__obf_4aae8d20fd4eaf66, 0)

	if __obf_a6c4d873d619640b == -1 && !__obf_6a73025256814cba.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}
	__obf_dd70e8929c261ac7, _ := __obf_df803c2772cc0386.__obf_2174f2309638b62a(__obf_ddb81b725bab291a.__obf_8f559e1327c9ee84,

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_e69ecd4c32f1420b)

	if __obf_6a73025256814cba.__obf_8f559e1327c9ee84.Sign() == 0 {
		return __obf_dd70e8929c261ac7, nil
	}
	__obf_51c5f2b2b2662b7e := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_df803c2772cc0386.NumDigits()
	__obf_038c585447669ee5 := __obf_b6bc08d52d165029.NumDigits()

	if int32(__obf_51c5f2b2b2662b7e) > __obf_e69ecd4c32f1420b {
		__obf_e69ecd4c32f1420b = int32(__obf_51c5f2b2b2662b7e)
	}
	if int32(__obf_038c585447669ee5) > __obf_e69ecd4c32f1420b {
		__obf_e69ecd4c32f1420b += int32(__obf_038c585447669ee5)
	}
	__obf_e69ecd4c32f1420b += // increase precision by 10 to compensate for errors in further calculations
		10
	__obf_07d089e0aaf4fa6e,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_b2a10ebf0533880b := __obf_df803c2772cc0386.Abs().Ln(__obf_e69ecd4c32f1420b)
	if __obf_b2a10ebf0533880b != nil {
		return Decimal{}, __obf_b2a10ebf0533880b
	}
	__obf_07d089e0aaf4fa6e = __obf_07d089e0aaf4fa6e.Mul(__obf_6a73025256814cba)
	__obf_07d089e0aaf4fa6e, __obf_b2a10ebf0533880b = __obf_07d089e0aaf4fa6e.ExpTaylor(__obf_e69ecd4c32f1420b)
	if __obf_b2a10ebf0533880b != nil {
		return Decimal{}, __obf_b2a10ebf0533880b
	}
	__obf_aee88fdbf6e4041b := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_dd70e8929c261ac7.Mul(__obf_07d089e0aaf4fa6e)

	return __obf_aee88fdbf6e4041b, nil
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
func (__obf_df803c2772cc0386 Decimal) PowInt32(__obf_97e810ed39f8ac3f int32) (Decimal, error) {
	if __obf_df803c2772cc0386.IsZero() && __obf_97e810ed39f8ac3f == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_aa595ea29b9e86ac := __obf_97e810ed39f8ac3f < 0
	__obf_97e810ed39f8ac3f = __obf_10bc131e66c76f40(__obf_97e810ed39f8ac3f)
	__obf_519018557784958f, __obf_16d1c91cb4076a0c := __obf_df803c2772cc0386, New(1, 0)

	for __obf_97e810ed39f8ac3f > 0 {
		if __obf_97e810ed39f8ac3f%2 == 1 {
			__obf_16d1c91cb4076a0c = __obf_16d1c91cb4076a0c.Mul(__obf_519018557784958f)
		}
		__obf_97e810ed39f8ac3f /= 2

		if __obf_97e810ed39f8ac3f > 0 {
			__obf_519018557784958f = __obf_519018557784958f.Mul(__obf_519018557784958f)
		}
	}

	if __obf_aa595ea29b9e86ac {
		return New(1, 0).DivRound(__obf_16d1c91cb4076a0c, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_16d1c91cb4076a0c, nil
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
func (__obf_df803c2772cc0386 Decimal) PowBigInt(__obf_97e810ed39f8ac3f *big.Int) (Decimal, error) {
	return __obf_df803c2772cc0386.__obf_2174f2309638b62a(__obf_97e810ed39f8ac3f, int32(PowPrecisionNegativeExponent))
}

func (__obf_df803c2772cc0386 Decimal) __obf_2174f2309638b62a(__obf_97e810ed39f8ac3f *big.Int, __obf_e69ecd4c32f1420b int32) (Decimal, error) {
	if __obf_df803c2772cc0386.IsZero() && __obf_97e810ed39f8ac3f.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_be9889c82fde2242 := new(big.Int).Set(__obf_97e810ed39f8ac3f)
	__obf_aa595ea29b9e86ac := __obf_97e810ed39f8ac3f.Sign() < 0

	if __obf_aa595ea29b9e86ac {
		__obf_be9889c82fde2242.
			Abs(__obf_be9889c82fde2242)
	}
	__obf_519018557784958f, __obf_16d1c91cb4076a0c := __obf_df803c2772cc0386, New(1, 0)

	for __obf_be9889c82fde2242.Sign() > 0 {
		if __obf_be9889c82fde2242.Bit(0) == 1 {
			__obf_16d1c91cb4076a0c = __obf_16d1c91cb4076a0c.Mul(__obf_519018557784958f)
		}
		__obf_be9889c82fde2242.
			Rsh(__obf_be9889c82fde2242, 1)

		if __obf_be9889c82fde2242.Sign() > 0 {
			__obf_519018557784958f = __obf_519018557784958f.Mul(__obf_519018557784958f)
		}
	}

	if __obf_aa595ea29b9e86ac {
		return New(1, 0).DivRound(__obf_16d1c91cb4076a0c, __obf_e69ecd4c32f1420b), nil
	}

	return __obf_16d1c91cb4076a0c, nil
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
func (__obf_df803c2772cc0386 Decimal) ExpHullAbrham(__obf_fa7ea26b503ef3e1 uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_df803c2772cc0386.IsZero() {
		return Decimal{__obf_c8643a8d2b2d6a40, 0}, nil
	}
	__obf_3c56b50a4a14ae32 := __obf_fa7ea26b503ef3e1

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_047f3bcbc753afbe := __obf_df803c2772cc0386.Abs().InexactFloat64()
	if __obf_8bc3b5efe8592354 := __obf_047f3bcbc753afbe / 23; __obf_8bc3b5efe8592354 > float64(__obf_3c56b50a4a14ae32) && __obf_8bc3b5efe8592354 < float64(ExpMaxIterations) {
		__obf_3c56b50a4a14ae32 = uint32(math.Ceil(__obf_8bc3b5efe8592354))
	}
	__obf_85dd3e722e422ed6 := // fail if abs(d) beyond an over/underflow threshold
		New(23*int64(__obf_3c56b50a4a14ae32), 0)
	if __obf_df803c2772cc0386.Abs().Cmp(__obf_85dd3e722e422ed6) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}
	__obf_9f25f8c29e3a541b := // Return 1 if abs(d) small enough; this also avoids later over/underflow
		New(9, -int32(__obf_3c56b50a4a14ae32)-1)
	if __obf_df803c2772cc0386.Abs().Cmp(__obf_9f25f8c29e3a541b) <= 0 {
		return Decimal{__obf_c8643a8d2b2d6a40, __obf_df803c2772cc0386.

			// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
			__obf_97e810ed39f8ac3f}, nil
	}
	__obf_8d7b8d6cca70e2b2 := __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f + int32(__obf_df803c2772cc0386.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_8d7b8d6cca70e2b2 < 0 {
		__obf_8d7b8d6cca70e2b2 = 0
	}
	__obf_fb209a5c6c42bbe3 := New(1, __obf_8d7b8d6cca70e2b2)
	__obf_69a72b316864499d := // reduction factor
		Decimal{new(big.Int).Set(__obf_df803c2772cc0386. // reduced argument
									__obf_8f559e1327c9ee84), __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f - __obf_8d7b8d6cca70e2b2}
	__obf_aa343990ddb4add0 := int32(__obf_3c56b50a4a14ae32) + __obf_8d7b8d6cca70e2b2 + 2
	__obf_8a89433e6918bd32 := // precision for calculating the sum

		// Determine n, the number of therms for calculating sum
		// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
		// for solving appropriate equation, along with directed
		// roundings and simple rational bound for log10(p/abs(r))
		__obf_69a72b316864499d.Abs().InexactFloat64()
	__obf_aaabdf1979c5e455 := float64(__obf_aa343990ddb4add0)
	__obf_3e13207cc01cfb67 := math.Ceil((1.453*__obf_aaabdf1979c5e455 - 1.182) / math.Log10(__obf_aaabdf1979c5e455/__obf_8a89433e6918bd32))
	if __obf_3e13207cc01cfb67 > float64(ExpMaxIterations) || math.IsNaN(__obf_3e13207cc01cfb67) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_519018557784958f := int64(__obf_3e13207cc01cfb67)
	__obf_ba9c7603fdcb2493 := New(0, 0)
	__obf_ee1eb81404c3703d := New(1, 0)
	__obf_4aae8d20fd4eaf66 := New(1, 0)
	for __obf_0038051e77a34ba0 := __obf_519018557784958f - 1; __obf_0038051e77a34ba0 > 0; __obf_0038051e77a34ba0-- {
		__obf_ba9c7603fdcb2493.__obf_8f559e1327c9ee84.
			SetInt64(__obf_0038051e77a34ba0)
		__obf_ee1eb81404c3703d = __obf_ee1eb81404c3703d.Mul(__obf_69a72b316864499d.DivRound(__obf_ba9c7603fdcb2493, __obf_aa343990ddb4add0))
		__obf_ee1eb81404c3703d = __obf_ee1eb81404c3703d.Add(__obf_4aae8d20fd4eaf66)
	}
	__obf_6cd00f677cac2158 := __obf_fb209a5c6c42bbe3.IntPart()
	__obf_aee88fdbf6e4041b := New(1, 0)
	for __obf_0038051e77a34ba0 := __obf_6cd00f677cac2158; __obf_0038051e77a34ba0 > 0; __obf_0038051e77a34ba0-- {
		__obf_aee88fdbf6e4041b = __obf_aee88fdbf6e4041b.Mul(__obf_ee1eb81404c3703d)
	}
	__obf_e9ac145450a2dc3d := int32(__obf_aee88fdbf6e4041b.NumDigits())

	var __obf_75f262ad156de50c int32
	if __obf_e9ac145450a2dc3d > __obf_10bc131e66c76f40(__obf_aee88fdbf6e4041b.__obf_97e810ed39f8ac3f) {
		__obf_75f262ad156de50c = int32(__obf_3c56b50a4a14ae32) - __obf_e9ac145450a2dc3d - __obf_aee88fdbf6e4041b.__obf_97e810ed39f8ac3f
	} else {
		__obf_75f262ad156de50c = int32(__obf_3c56b50a4a14ae32)
	}
	__obf_aee88fdbf6e4041b = __obf_aee88fdbf6e4041b.Round(__obf_75f262ad156de50c)

	return __obf_aee88fdbf6e4041b, nil
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
func (__obf_df803c2772cc0386 Decimal) ExpTaylor(__obf_e69ecd4c32f1420b int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_df803c2772cc0386.IsZero() {
		return Decimal{__obf_c8643a8d2b2d6a40, 0}.Round(__obf_e69ecd4c32f1420b), nil
	}

	var __obf_0914ef8b977edfc9 Decimal
	var __obf_1789feb8b392f9e2 int32
	if __obf_e69ecd4c32f1420b < 0 {
		__obf_0914ef8b977edfc9 = New(1, -1)
		__obf_1789feb8b392f9e2 = 8
	} else {
		__obf_0914ef8b977edfc9 = New(1, -__obf_e69ecd4c32f1420b-1)
		__obf_1789feb8b392f9e2 = __obf_e69ecd4c32f1420b + 1
	}
	__obf_cb27c97bba1191f5 := __obf_df803c2772cc0386.Abs()
	__obf_0587ecb24ae8cf92 := __obf_df803c2772cc0386.Abs()
	__obf_7de627171a30939c := New(1, 0)
	__obf_16d1c91cb4076a0c := New(1, 0)

	for __obf_0038051e77a34ba0 := int64(1); ; {
		__obf_44d4d96115e6a8b2 := __obf_0587ecb24ae8cf92.DivRound(__obf_7de627171a30939c, __obf_1789feb8b392f9e2)
		__obf_16d1c91cb4076a0c = __obf_16d1c91cb4076a0c.Add(__obf_44d4d96115e6a8b2)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_44d4d96115e6a8b2.Cmp(__obf_0914ef8b977edfc9) < 0 {
			break
		}
		__obf_0587ecb24ae8cf92 = __obf_0587ecb24ae8cf92.Mul(__obf_cb27c97bba1191f5)
		__obf_0038051e77a34ba0++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_a7f8febfb20b3c1e) >= int(__obf_0038051e77a34ba0) && !__obf_a7f8febfb20b3c1e[__obf_0038051e77a34ba0-1].IsZero() {
			__obf_7de627171a30939c = __obf_a7f8febfb20b3c1e[__obf_0038051e77a34ba0-1]
		} else {
			__obf_7de627171a30939c = // To avoid any race conditions, firstly the zero value is appended to a slice to create
				// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
				// factorial using the index notation.
				__obf_a7f8febfb20b3c1e[__obf_0038051e77a34ba0-2].Mul(New(__obf_0038051e77a34ba0, 0))
			__obf_a7f8febfb20b3c1e = append(__obf_a7f8febfb20b3c1e, Zero)
			__obf_a7f8febfb20b3c1e[__obf_0038051e77a34ba0-1] = __obf_7de627171a30939c
		}
	}

	if __obf_df803c2772cc0386.Sign() < 0 {
		__obf_16d1c91cb4076a0c = New(1, 0).DivRound(__obf_16d1c91cb4076a0c, __obf_e69ecd4c32f1420b+1)
	}
	__obf_16d1c91cb4076a0c = __obf_16d1c91cb4076a0c.Round(__obf_e69ecd4c32f1420b)
	return __obf_16d1c91cb4076a0c, nil
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
func (__obf_df803c2772cc0386 Decimal) Ln(__obf_e69ecd4c32f1420b int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_df803c2772cc0386.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_df803c2772cc0386.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}
	__obf_498694f5e0e5ee19 := __obf_e69ecd4c32f1420b + 2
	__obf_9f51474a44562371 := __obf_df803c2772cc0386.Copy()

	var __obf_69a9dd8c3e7c03be, __obf_f4ca130b57e71b36, __obf_302e7ae0d904ebf8, __obf_23c55a6becee1145, __obf_98d8e8f0bd42156e Decimal
	__obf_69a9dd8c3e7c03be = __obf_9f51474a44562371.Sub(Decimal{__obf_c8643a8d2b2d6a40, 0})
	__obf_f4ca130b57e71b36 = Decimal{__obf_c8643a8d2b2d6a40, -1}
	__obf_081dba7565cd4f79 := // for decimal in range [0.9, 1.1] where ln(d) is close to 0
		false

	if __obf_69a9dd8c3e7c03be.Abs().Cmp(__obf_f4ca130b57e71b36) <= 0 {
		__obf_081dba7565cd4f79 = true
	} else {
		__obf_d32896ffdd32ed47 := // reduce input decimal to range [0.1, 1)
			int32(__obf_9f51474a44562371.NumDigits()) + __obf_9f51474a44562371.__obf_97e810ed39f8ac3f

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_9f51474a44562371.__obf_97e810ed39f8ac3f -= __obf_d32896ffdd32ed47
		__obf_3865966de6fb72b2 := __obf_3865966de6fb72b2.__obf_fd94e66514481507(__obf_498694f5e0e5ee19)
		__obf_98d8e8f0bd42156e = NewFromInt32(__obf_d32896ffdd32ed47)
		__obf_98d8e8f0bd42156e = __obf_98d8e8f0bd42156e.Mul(__obf_3865966de6fb72b2)
		__obf_69a9dd8c3e7c03be = __obf_9f51474a44562371.Sub(Decimal{__obf_c8643a8d2b2d6a40, 0})

		if __obf_69a9dd8c3e7c03be.Abs().Cmp(__obf_f4ca130b57e71b36) <= 0 {
			__obf_081dba7565cd4f79 = true
		} else {
			__obf_90e831fbd4825f23 := // initial estimate using floats
				__obf_9f51474a44562371.InexactFloat64()
			__obf_69a9dd8c3e7c03be = NewFromFloat(math.Log(__obf_90e831fbd4825f23))
		}
	}
	__obf_0914ef8b977edfc9 := Decimal{__obf_c8643a8d2b2d6a40, -__obf_498694f5e0e5ee19}

	if __obf_081dba7565cd4f79 {
		__obf_302e7ae0d904ebf8 = // Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
			// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			// until the difference between current and next term is smaller than epsilon.
			// Coverage quite fast for decimals close to 1.0

			// z + 2
			__obf_69a9dd8c3e7c03be.Add(Decimal{__obf_f741b5f088b7b4b7, 0})
		__obf_f4ca130b57e71b36 = // z / (z + 2)
			__obf_69a9dd8c3e7c03be.DivRound(__obf_302e7ae0d904ebf8, __obf_498694f5e0e5ee19)
		__obf_69a9dd8c3e7c03be = // 2 * (z / (z + 2))
			__obf_f4ca130b57e71b36.Add(__obf_f4ca130b57e71b36)
		__obf_302e7ae0d904ebf8 = __obf_69a9dd8c3e7c03be.Copy()

		for __obf_519018557784958f := 1; ; __obf_519018557784958f++ {
			__obf_302e7ae0d904ebf8 = // 2 * (z / (z+2))^(2n+1)
				__obf_302e7ae0d904ebf8.Mul(__obf_f4ca130b57e71b36).Mul(__obf_f4ca130b57e71b36)
			__obf_23c55a6becee1145 = // 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
				NewFromInt(int64(2*__obf_519018557784958f + 1))
			__obf_23c55a6becee1145 = __obf_302e7ae0d904ebf8.DivRound(__obf_23c55a6becee1145, __obf_498694f5e0e5ee19)
			__obf_69a9dd8c3e7c03be = // comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
				__obf_69a9dd8c3e7c03be.Add(__obf_23c55a6becee1145)

			if __obf_23c55a6becee1145.Abs().Cmp(__obf_0914ef8b977edfc9) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_89383cab1a537792 Decimal
		__obf_83e4e421e289635d := __obf_498694f5e0e5ee19*2 + 10

		for __obf_0038051e77a34ba0 := int32(0); __obf_0038051e77a34ba0 < __obf_83e4e421e289635d;
		// exp(a_n)
		__obf_0038051e77a34ba0++ {
			__obf_f4ca130b57e71b36, _ = __obf_69a9dd8c3e7c03be.ExpTaylor(__obf_498694f5e0e5ee19)
			__obf_302e7ae0d904ebf8 = // exp(a_n) - z
				__obf_f4ca130b57e71b36.Sub(__obf_9f51474a44562371)
			__obf_302e7ae0d904ebf8 = // 2 * (exp(a_n) - z)
				__obf_302e7ae0d904ebf8.Add(__obf_302e7ae0d904ebf8)
			__obf_23c55a6becee1145 = // exp(a_n) + z
				__obf_f4ca130b57e71b36.Add(__obf_9f51474a44562371)
			__obf_f4ca130b57e71b36 = // 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_302e7ae0d904ebf8.DivRound(__obf_23c55a6becee1145, __obf_498694f5e0e5ee19)
			__obf_69a9dd8c3e7c03be = // comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_69a9dd8c3e7c03be.Sub(__obf_f4ca130b57e71b36)

			if __obf_89383cab1a537792.Add(__obf_f4ca130b57e71b36).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_f4ca130b57e71b36.Abs().Cmp(__obf_0914ef8b977edfc9) <= 0 {
				break
			}
			__obf_89383cab1a537792 = __obf_f4ca130b57e71b36
		}
	}
	__obf_69a9dd8c3e7c03be = __obf_69a9dd8c3e7c03be.Add(__obf_98d8e8f0bd42156e)

	return __obf_69a9dd8c3e7c03be.Round(__obf_e69ecd4c32f1420b), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_df803c2772cc0386 Decimal) NumDigits() int {
	if __obf_df803c2772cc0386.__obf_8f559e1327c9ee84 == nil {
		return 1
	}

	if __obf_df803c2772cc0386.__obf_8f559e1327c9ee84.IsInt64() {
		__obf_7b50cbb1d16fcf58 := __obf_df803c2772cc0386.
			// restrict fast path to integers with exact conversion to float64
			__obf_8f559e1327c9ee84.Int64()

		if __obf_7b50cbb1d16fcf58 <= (1<<53) && __obf_7b50cbb1d16fcf58 >= -(1<<53) {
			if __obf_7b50cbb1d16fcf58 == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_7b50cbb1d16fcf58)))) + 1
		}
	}
	__obf_71e03f080d8dda57 := int(float64(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84.BitLen()) / math.Log2(10))
	__obf_a0c2f72f201efcba := // estimatedNumDigits (lg10) may be off by 1, need to verify
		big.NewInt(int64(__obf_71e03f080d8dda57))
	__obf_d8160dd41ebf11f0 := __obf_a0c2f72f201efcba.Exp(__obf_d4d5e0f4e5fc42f5, __obf_a0c2f72f201efcba, nil)

	if __obf_df803c2772cc0386.__obf_8f559e1327c9ee84.CmpAbs(__obf_d8160dd41ebf11f0) >= 0 {
		return __obf_71e03f080d8dda57 + 1
	}

	return __obf_71e03f080d8dda57
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_df803c2772cc0386 Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_69a72b316864499d big.Int
	__obf_2a31e4af66195793 := new(big.Int).Set(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84)
	for __obf_9f51474a44562371 := __obf_10bc131e66c76f40(__obf_df803c2772cc0386.__obf_97e810ed39f8ac3f); __obf_9f51474a44562371 > 0; __obf_9f51474a44562371-- {
		__obf_2a31e4af66195793.
			QuoRem(__obf_2a31e4af66195793, __obf_d4d5e0f4e5fc42f5, &__obf_69a72b316864499d)
		if __obf_69a72b316864499d.Cmp(__obf_ec129f5235aa8458) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_10bc131e66c76f40(__obf_519018557784958f int32) int32 {
	if __obf_519018557784958f < 0 {
		return -__obf_519018557784958f
	}
	return __obf_519018557784958f
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_df803c2772cc0386 Decimal) Cmp(__obf_b6bc08d52d165029 Decimal) int {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()
	__obf_b6bc08d52d165029.__obf_a6a609bb966fabfb()

	if __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f == __obf_b6bc08d52d165029.__obf_97e810ed39f8ac3f {
		return __obf_df803c2772cc0386.__obf_8f559e1327c9ee84.Cmp(__obf_b6bc08d52d165029.__obf_8f559e1327c9ee84)
	}
	__obf_ce5f1b4c4db8f0f4, __obf_b9159a371b250a0c := RescalePair(__obf_df803c2772cc0386, __obf_b6bc08d52d165029)

	return __obf_ce5f1b4c4db8f0f4.__obf_8f559e1327c9ee84.Cmp(__obf_b9159a371b250a0c.

		// Compare compares the numbers represented by d and d2 and returns:
		//
		//	-1 if d <  d2
		//	 0 if d == d2
		//	+1 if d >  d2
		__obf_8f559e1327c9ee84)
}

func (__obf_df803c2772cc0386 Decimal) Compare(__obf_b6bc08d52d165029 Decimal) int {
	return __obf_df803c2772cc0386.Cmp(__obf_b6bc08d52d165029)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_df803c2772cc0386 Decimal) Equal(__obf_b6bc08d52d165029 Decimal) bool {
	return __obf_df803c2772cc0386.Cmp(__obf_b6bc08d52d165029) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_df803c2772cc0386 Decimal) Equals(__obf_b6bc08d52d165029 Decimal) bool {
	return __obf_df803c2772cc0386.Equal(__obf_b6bc08d52d165029)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_df803c2772cc0386 Decimal) GreaterThan(__obf_b6bc08d52d165029 Decimal) bool {
	return __obf_df803c2772cc0386.Cmp(__obf_b6bc08d52d165029) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_df803c2772cc0386 Decimal) GreaterThanOrEqual(__obf_b6bc08d52d165029 Decimal) bool {
	__obf_0ea848c2276951cd := __obf_df803c2772cc0386.Cmp(__obf_b6bc08d52d165029)
	return __obf_0ea848c2276951cd == 1 || __obf_0ea848c2276951cd == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_df803c2772cc0386 Decimal) LessThan(__obf_b6bc08d52d165029 Decimal) bool {
	return __obf_df803c2772cc0386.Cmp(__obf_b6bc08d52d165029) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_df803c2772cc0386 Decimal) LessThanOrEqual(__obf_b6bc08d52d165029 Decimal) bool {
	__obf_0ea848c2276951cd := __obf_df803c2772cc0386.Cmp(__obf_b6bc08d52d165029)
	return __obf_0ea848c2276951cd == -1 || __obf_0ea848c2276951cd == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_df803c2772cc0386 Decimal) Sign() int {
	if __obf_df803c2772cc0386.__obf_8f559e1327c9ee84 == nil {
		return 0
	}
	return __obf_df803c2772cc0386.

		// IsPositive return
		//
		//	true if d > 0
		//	false if d == 0
		//	false if d < 0
		__obf_8f559e1327c9ee84.Sign()
}

func (__obf_df803c2772cc0386 Decimal) IsPositive() bool {
	return __obf_df803c2772cc0386.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_df803c2772cc0386 Decimal) IsNegative() bool {
	return __obf_df803c2772cc0386.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_df803c2772cc0386 Decimal) IsZero() bool {
	return __obf_df803c2772cc0386.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_df803c2772cc0386 Decimal) Exponent() int32 {
	return __obf_df803c2772cc0386.

		// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
		__obf_97e810ed39f8ac3f
}

func (__obf_df803c2772cc0386 Decimal) Coefficient() *big.Int {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_df803c2772cc0386.

		// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
		// If coefficient cannot be represented in an int64, the result will be undefined.
		__obf_8f559e1327c9ee84)
}

func (__obf_df803c2772cc0386 Decimal) CoefficientInt64() int64 {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()
	return __obf_df803c2772cc0386.

		// IntPart returns the integer component of the decimal.
		__obf_8f559e1327c9ee84.Int64()
}

func (__obf_df803c2772cc0386 Decimal) IntPart() int64 {
	__obf_287112ad0a09cdfa := __obf_df803c2772cc0386.__obf_749ebd7071bc390b(0)
	return __obf_287112ad0a09cdfa.__obf_8f559e1327c9ee84.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_df803c2772cc0386 Decimal) BigInt() *big.Int {
	__obf_287112ad0a09cdfa := __obf_df803c2772cc0386.__obf_749ebd7071bc390b(0)
	return __obf_287112ad0a09cdfa.

		// BigFloat returns decimal as BigFloat.
		// Be aware that casting decimal to BigFloat might cause a loss of precision.
		__obf_8f559e1327c9ee84
}

func (__obf_df803c2772cc0386 Decimal) BigFloat() *big.Float {
	__obf_047f3bcbc753afbe := &big.Float{}
	__obf_047f3bcbc753afbe.
		SetString(__obf_df803c2772cc0386.String())
	return __obf_047f3bcbc753afbe
}

// Rat returns a rational number representation of the decimal.
func (__obf_df803c2772cc0386 Decimal) Rat() *big.Rat {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()
	if __obf_df803c2772cc0386.
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_97e810ed39f8ac3f <= 0 {
		__obf_a7a77930c9d1eea7 := new(big.Int).Exp(__obf_d4d5e0f4e5fc42f5, big.NewInt(-int64(__obf_df803c2772cc0386.__obf_97e810ed39f8ac3f)), nil)
		return new(big.Rat).SetFrac(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84, __obf_a7a77930c9d1eea7)
	}
	__obf_ba0bdedb5973d5fe := new(big.Int).Exp(__obf_d4d5e0f4e5fc42f5, big.NewInt(int64(__obf_df803c2772cc0386.__obf_97e810ed39f8ac3f)), nil)
	__obf_5ea998046095c3b8 := new(big.Int).Mul(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84, __obf_ba0bdedb5973d5fe)
	return new(big.Rat).SetFrac(__obf_5ea998046095c3b8,

		// Float64 returns the nearest float64 value for d and a bool indicating
		// whether f represents d exactly.
		// For more details, see the documentation for big.Rat.Float64
		__obf_c8643a8d2b2d6a40)
}

func (__obf_df803c2772cc0386 Decimal) Float64() (__obf_047f3bcbc753afbe float64, __obf_6855428b730a35ea bool) {
	return __obf_df803c2772cc0386.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_df803c2772cc0386 Decimal) InexactFloat64() float64 {
	__obf_047f3bcbc753afbe, _ := __obf_df803c2772cc0386.Float64()
	return __obf_047f3bcbc753afbe
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
func (__obf_df803c2772cc0386 Decimal) String() string {
	return __obf_df803c2772cc0386.string(true)
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
func (__obf_df803c2772cc0386 Decimal) StringFixed(__obf_7720d31c8a8713af int32) string {
	__obf_148eee7180dcb9c9 := __obf_df803c2772cc0386.Round(__obf_7720d31c8a8713af)
	return __obf_148eee7180dcb9c9.string(false)
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
func (__obf_df803c2772cc0386 Decimal) StringFixedBank(__obf_7720d31c8a8713af int32) string {
	__obf_148eee7180dcb9c9 := __obf_df803c2772cc0386.RoundBank(__obf_7720d31c8a8713af)
	return __obf_148eee7180dcb9c9.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_df803c2772cc0386 Decimal) StringFixedCash(__obf_9bd715875a8ab5f4 uint8) string {
	__obf_148eee7180dcb9c9 := __obf_df803c2772cc0386.RoundCash(__obf_9bd715875a8ab5f4)
	return __obf_148eee7180dcb9c9.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_df803c2772cc0386 Decimal) Round(__obf_7720d31c8a8713af int32) Decimal {
	if __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f == -__obf_7720d31c8a8713af {
		return __obf_df803c2772cc0386
	}
	__obf_3940c0a66218f180 := // truncate to places + 1
		__obf_df803c2772cc0386.__obf_749ebd7071bc390b(-__obf_7720d31c8a8713af - 1)

	// add sign(d) * 0.5
	if __obf_3940c0a66218f180.__obf_8f559e1327c9ee84.Sign() < 0 {
		__obf_3940c0a66218f180.__obf_8f559e1327c9ee84.
			Sub(__obf_3940c0a66218f180.__obf_8f559e1327c9ee84, __obf_4374a770c15c9bd3)
	} else {
		__obf_3940c0a66218f180.__obf_8f559e1327c9ee84.
			Add(__obf_3940c0a66218f180.__obf_8f559e1327c9ee84,

				// floor for positive numbers, ceil for negative numbers
				__obf_4374a770c15c9bd3)
	}

	_, __obf_cb9fa5e90307ec38 := __obf_3940c0a66218f180.__obf_8f559e1327c9ee84.DivMod(__obf_3940c0a66218f180.__obf_8f559e1327c9ee84, __obf_d4d5e0f4e5fc42f5, new(big.Int))
	__obf_3940c0a66218f180.__obf_97e810ed39f8ac3f++
	if __obf_3940c0a66218f180.__obf_8f559e1327c9ee84.Sign() < 0 && __obf_cb9fa5e90307ec38.Cmp(__obf_ec129f5235aa8458) != 0 {
		__obf_3940c0a66218f180.__obf_8f559e1327c9ee84.
			Add(__obf_3940c0a66218f180.__obf_8f559e1327c9ee84,

				// RoundCeil rounds the decimal towards +infinity.
				//
				// Example:
				//
				//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
				//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
				//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
				//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
				__obf_c8643a8d2b2d6a40)
	}

	return __obf_3940c0a66218f180
}

func (__obf_df803c2772cc0386 Decimal) RoundCeil(__obf_7720d31c8a8713af int32) Decimal {
	if __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f >= -__obf_7720d31c8a8713af {
		return __obf_df803c2772cc0386
	}
	__obf_728aa6b7aa2a955b := __obf_df803c2772cc0386.__obf_749ebd7071bc390b(-__obf_7720d31c8a8713af)
	if __obf_df803c2772cc0386.Equal(__obf_728aa6b7aa2a955b) {
		return __obf_df803c2772cc0386
	}

	if __obf_df803c2772cc0386.__obf_8f559e1327c9ee84.Sign() > 0 {
		__obf_728aa6b7aa2a955b.__obf_8f559e1327c9ee84.
			Add(__obf_728aa6b7aa2a955b.__obf_8f559e1327c9ee84, __obf_c8643a8d2b2d6a40)
	}

	return __obf_728aa6b7aa2a955b
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_df803c2772cc0386 Decimal) RoundFloor(__obf_7720d31c8a8713af int32) Decimal {
	if __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f >= -__obf_7720d31c8a8713af {
		return __obf_df803c2772cc0386
	}
	__obf_728aa6b7aa2a955b := __obf_df803c2772cc0386.__obf_749ebd7071bc390b(-__obf_7720d31c8a8713af)
	if __obf_df803c2772cc0386.Equal(__obf_728aa6b7aa2a955b) {
		return __obf_df803c2772cc0386
	}

	if __obf_df803c2772cc0386.__obf_8f559e1327c9ee84.Sign() < 0 {
		__obf_728aa6b7aa2a955b.__obf_8f559e1327c9ee84.
			Sub(__obf_728aa6b7aa2a955b.__obf_8f559e1327c9ee84, __obf_c8643a8d2b2d6a40)
	}

	return __obf_728aa6b7aa2a955b
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_df803c2772cc0386 Decimal) RoundUp(__obf_7720d31c8a8713af int32) Decimal {
	if __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f >= -__obf_7720d31c8a8713af {
		return __obf_df803c2772cc0386
	}
	__obf_728aa6b7aa2a955b := __obf_df803c2772cc0386.__obf_749ebd7071bc390b(-__obf_7720d31c8a8713af)
	if __obf_df803c2772cc0386.Equal(__obf_728aa6b7aa2a955b) {
		return __obf_df803c2772cc0386
	}

	if __obf_df803c2772cc0386.__obf_8f559e1327c9ee84.Sign() > 0 {
		__obf_728aa6b7aa2a955b.__obf_8f559e1327c9ee84.
			Add(__obf_728aa6b7aa2a955b.__obf_8f559e1327c9ee84, __obf_c8643a8d2b2d6a40)
	} else if __obf_df803c2772cc0386.__obf_8f559e1327c9ee84.Sign() < 0 {
		__obf_728aa6b7aa2a955b.__obf_8f559e1327c9ee84.
			Sub(__obf_728aa6b7aa2a955b.__obf_8f559e1327c9ee84, __obf_c8643a8d2b2d6a40)
	}

	return __obf_728aa6b7aa2a955b
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_df803c2772cc0386 Decimal) RoundDown(__obf_7720d31c8a8713af int32) Decimal {
	if __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f >= -__obf_7720d31c8a8713af {
		return __obf_df803c2772cc0386
	}
	__obf_728aa6b7aa2a955b := __obf_df803c2772cc0386.__obf_749ebd7071bc390b(-__obf_7720d31c8a8713af)
	if __obf_df803c2772cc0386.Equal(__obf_728aa6b7aa2a955b) {
		return __obf_df803c2772cc0386
	}
	return __obf_728aa6b7aa2a955b
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
func (__obf_df803c2772cc0386 Decimal) RoundBank(__obf_7720d31c8a8713af int32) Decimal {
	__obf_fc448115c899b670 := __obf_df803c2772cc0386.Round(__obf_7720d31c8a8713af)
	__obf_f60cde1e62406605 := __obf_df803c2772cc0386.Sub(__obf_fc448115c899b670).Abs()
	__obf_c95eaea2b1d5fdfd := New(5, -__obf_7720d31c8a8713af-1)
	if __obf_f60cde1e62406605.Cmp(__obf_c95eaea2b1d5fdfd) == 0 && __obf_fc448115c899b670.__obf_8f559e1327c9ee84.Bit(0) != 0 {
		if __obf_fc448115c899b670.__obf_8f559e1327c9ee84.Sign() < 0 {
			__obf_fc448115c899b670.__obf_8f559e1327c9ee84.
				Add(__obf_fc448115c899b670.__obf_8f559e1327c9ee84, __obf_c8643a8d2b2d6a40)
		} else {
			__obf_fc448115c899b670.__obf_8f559e1327c9ee84.
				Sub(__obf_fc448115c899b670.__obf_8f559e1327c9ee84, __obf_c8643a8d2b2d6a40)
		}
	}

	return __obf_fc448115c899b670
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
func (__obf_df803c2772cc0386 Decimal) RoundCash(__obf_9bd715875a8ab5f4 uint8) Decimal {
	var __obf_d665fc20ced50c30 *big.Int
	switch __obf_9bd715875a8ab5f4 {
	case 5:
		__obf_d665fc20ced50c30 = __obf_9195a8e64ca7e21a
	case 10:
		__obf_d665fc20ced50c30 = __obf_d4d5e0f4e5fc42f5
	case 25:
		__obf_d665fc20ced50c30 = __obf_55b28070f37dbd50
	case 50:
		__obf_d665fc20ced50c30 = __obf_f741b5f088b7b4b7
	case 100:
		__obf_d665fc20ced50c30 = __obf_c8643a8d2b2d6a40
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_9bd715875a8ab5f4))
	}
	__obf_084d1d7b80198719 := Decimal{__obf_8f559e1327c9ee84: __obf_d665fc20ced50c30}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_df803c2772cc0386.Mul(__obf_084d1d7b80198719).Round(0).Div(__obf_084d1d7b80198719).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_df803c2772cc0386 Decimal) Floor() Decimal {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()

	if __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f >= 0 {
		return __obf_df803c2772cc0386
	}
	__obf_97e810ed39f8ac3f := big.NewInt(10)
	__obf_97e810ed39f8ac3f.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_97e810ed39f8ac3f, big.NewInt(-int64(__obf_df803c2772cc0386.__obf_97e810ed39f8ac3f)), nil)
	__obf_9f51474a44562371 := new(big.Int).Div(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84, __obf_97e810ed39f8ac3f)
	return Decimal{__obf_8f559e1327c9ee84: __obf_9f51474a44562371,

		// Ceil returns the nearest integer value greater than or equal to d.
		__obf_97e810ed39f8ac3f: 0}
}

func (__obf_df803c2772cc0386 Decimal) Ceil() Decimal {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()

	if __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f >= 0 {
		return __obf_df803c2772cc0386
	}
	__obf_97e810ed39f8ac3f := big.NewInt(10)
	__obf_97e810ed39f8ac3f.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_97e810ed39f8ac3f, big.NewInt(-int64(__obf_df803c2772cc0386.__obf_97e810ed39f8ac3f)), nil)
	__obf_9f51474a44562371, __obf_cb9fa5e90307ec38 := new(big.Int).DivMod(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84, __obf_97e810ed39f8ac3f, new(big.Int))
	if __obf_cb9fa5e90307ec38.Cmp(__obf_ec129f5235aa8458) != 0 {
		__obf_9f51474a44562371.
			Add(__obf_9f51474a44562371, __obf_c8643a8d2b2d6a40)
	}
	return Decimal{__obf_8f559e1327c9ee84: __obf_9f51474a44562371,

		// Truncate truncates off digits from the number, without rounding.
		//
		// NOTE: precision is the last digit that will not be truncated (must be >= 0).
		//
		// Example:
		//
		//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
		__obf_97e810ed39f8ac3f: 0}
}

func (__obf_df803c2772cc0386 Decimal) Truncate(__obf_e69ecd4c32f1420b int32) Decimal {
	__obf_df803c2772cc0386.__obf_a6a609bb966fabfb()
	if __obf_e69ecd4c32f1420b >= 0 && -__obf_e69ecd4c32f1420b > __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f {
		return __obf_df803c2772cc0386.__obf_749ebd7071bc390b(-__obf_e69ecd4c32f1420b)
	}
	return __obf_df803c2772cc0386
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_df803c2772cc0386 *Decimal) UnmarshalJSON(__obf_7b9bd66f12aa0215 []byte) error {
	if string(__obf_7b9bd66f12aa0215) == "null" {
		return nil
	}
	__obf_b3226814fd131eda, __obf_b2a10ebf0533880b := __obf_e7f7278a29822adc(__obf_7b9bd66f12aa0215)
	if __obf_b2a10ebf0533880b != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_7b9bd66f12aa0215, __obf_b2a10ebf0533880b)
	}
	__obf_b0e21776b9d45e13, __obf_b2a10ebf0533880b := NewFromString(__obf_b3226814fd131eda)
	*__obf_df803c2772cc0386 = __obf_b0e21776b9d45e13
	if __obf_b2a10ebf0533880b != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_b3226814fd131eda, __obf_b2a10ebf0533880b)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_df803c2772cc0386 Decimal) MarshalJSON() ([]byte, error) {
	var __obf_b3226814fd131eda string
	if MarshalJSONWithoutQuotes {
		__obf_b3226814fd131eda = __obf_df803c2772cc0386.String()
	} else {
		__obf_b3226814fd131eda = "\"" + __obf_df803c2772cc0386.String() + "\""
	}
	return []byte(__obf_b3226814fd131eda), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_df803c2772cc0386 *Decimal) UnmarshalBinary(__obf_47676632f1a99ad8 []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_47676632f1a99ad8) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_47676632f1a99ad8, len(__obf_47676632f1a99ad8))
	}
	__obf_df803c2772cc0386.

		// Extract the exponent
		__obf_97e810ed39f8ac3f = int32(binary.BigEndian.Uint32(__obf_47676632f1a99ad8[:4]))
	__obf_df803c2772cc0386.

		// Extract the value
		__obf_8f559e1327c9ee84 = new(big.Int)
	if __obf_b2a10ebf0533880b := __obf_df803c2772cc0386.__obf_8f559e1327c9ee84.GobDecode(__obf_47676632f1a99ad8[4:]); __obf_b2a10ebf0533880b != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_47676632f1a99ad8, __obf_b2a10ebf0533880b)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_df803c2772cc0386 Decimal) MarshalBinary() (__obf_47676632f1a99ad8 []byte, __obf_b2a10ebf0533880b error) {
	// exp is written first, but encode value first to know output size
	var __obf_9c2c651271a83fad []byte
	if __obf_9c2c651271a83fad, __obf_b2a10ebf0533880b = __obf_df803c2772cc0386.__obf_8f559e1327c9ee84.GobEncode(); __obf_b2a10ebf0533880b != nil {
		return nil, __obf_b2a10ebf0533880b
	}
	__obf_eb9578d334aad456 := // Write the exponent in front, since it's a fixed size
		make([]byte, 4, len(__obf_9c2c651271a83fad)+4)
	binary.BigEndian.PutUint32(__obf_eb9578d334aad456, uint32(__obf_df803c2772cc0386.

		// Return the byte array
		__obf_97e810ed39f8ac3f))

	return append(__obf_eb9578d334aad456, __obf_9c2c651271a83fad...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_df803c2772cc0386 *Decimal) Scan(__obf_8f559e1327c9ee84 any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_7c58e6101912d0e0 := __obf_8f559e1327c9ee84.(type) {

	case float32:
		*__obf_df803c2772cc0386 = NewFromFloat(float64(__obf_7c58e6101912d0e0))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_df803c2772cc0386 = NewFromFloat(__obf_7c58e6101912d0e0)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_df803c2772cc0386 = New(__obf_7c58e6101912d0e0, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_df803c2772cc0386 = NewFromUint64(__obf_7c58e6101912d0e0)
		return nil

	default:
		__obf_b3226814fd131eda,
			// default is trying to interpret value stored as string
			__obf_b2a10ebf0533880b := __obf_e7f7278a29822adc(__obf_7c58e6101912d0e0)
		if __obf_b2a10ebf0533880b != nil {
			return __obf_b2a10ebf0533880b
		}
		*__obf_df803c2772cc0386, __obf_b2a10ebf0533880b = NewFromString(__obf_b3226814fd131eda)
		return __obf_b2a10ebf0533880b
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_df803c2772cc0386 Decimal) Value() (driver.Value, error) {
	return __obf_df803c2772cc0386.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_df803c2772cc0386 *Decimal) UnmarshalText(__obf_4104aedfa89212c3 []byte) error {
	__obf_b3226814fd131eda := string(__obf_4104aedfa89212c3)
	__obf_428f69476005da25, __obf_b2a10ebf0533880b := NewFromString(__obf_b3226814fd131eda)
	*__obf_df803c2772cc0386 = __obf_428f69476005da25
	if __obf_b2a10ebf0533880b != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_b3226814fd131eda, __obf_b2a10ebf0533880b)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_df803c2772cc0386 Decimal) MarshalText() (__obf_4104aedfa89212c3 []byte, __obf_b2a10ebf0533880b error) {
	return []byte(__obf_df803c2772cc0386.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_df803c2772cc0386 Decimal) GobEncode() ([]byte, error) {
	return __obf_df803c2772cc0386.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_df803c2772cc0386 *Decimal) GobDecode(__obf_47676632f1a99ad8 []byte) error {
	return __obf_df803c2772cc0386.UnmarshalBinary(__obf_47676632f1a99ad8)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_df803c2772cc0386 Decimal) StringScaled(__obf_97e810ed39f8ac3f int32) string {
	return __obf_df803c2772cc0386.__obf_749ebd7071bc390b(__obf_97e810ed39f8ac3f).String()
}

func (__obf_df803c2772cc0386 Decimal) string(__obf_6c66a8bbfe1f681a bool) string {
	if __obf_df803c2772cc0386.__obf_97e810ed39f8ac3f >= 0 {
		return __obf_df803c2772cc0386.__obf_749ebd7071bc390b(0).__obf_8f559e1327c9ee84.String()
	}
	__obf_10bc131e66c76f40 := new(big.Int).Abs(__obf_df803c2772cc0386.__obf_8f559e1327c9ee84)
	__obf_b3226814fd131eda := __obf_10bc131e66c76f40.String()

	var __obf_2cd91842aa2c42c5, __obf_f5a2ccd7d2c2599c string
	__obf_b656472b92ac0372 := // NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
		// and you are on a 32-bit machine. Won't fix this super-edge case.
		int(__obf_df803c2772cc0386.__obf_97e810ed39f8ac3f)
	if len(__obf_b3226814fd131eda) > -__obf_b656472b92ac0372 {
		__obf_2cd91842aa2c42c5 = __obf_b3226814fd131eda[:len(__obf_b3226814fd131eda)+__obf_b656472b92ac0372]
		__obf_f5a2ccd7d2c2599c = __obf_b3226814fd131eda[len(__obf_b3226814fd131eda)+__obf_b656472b92ac0372:]
	} else {
		__obf_2cd91842aa2c42c5 = "0"
		__obf_7dfa9fee1ca1bad7 := -__obf_b656472b92ac0372 - len(__obf_b3226814fd131eda)
		__obf_f5a2ccd7d2c2599c = strings.Repeat("0", __obf_7dfa9fee1ca1bad7) + __obf_b3226814fd131eda
	}

	if __obf_6c66a8bbfe1f681a {
		__obf_0038051e77a34ba0 := len(__obf_f5a2ccd7d2c2599c) - 1
		for ; __obf_0038051e77a34ba0 >= 0; __obf_0038051e77a34ba0-- {
			if __obf_f5a2ccd7d2c2599c[__obf_0038051e77a34ba0] != '0' {
				break
			}
		}
		__obf_f5a2ccd7d2c2599c = __obf_f5a2ccd7d2c2599c[:__obf_0038051e77a34ba0+1]
	}
	__obf_35bdb6d9709e131b := __obf_2cd91842aa2c42c5
	if len(__obf_f5a2ccd7d2c2599c) > 0 {
		__obf_35bdb6d9709e131b += "." + __obf_f5a2ccd7d2c2599c
	}

	if __obf_df803c2772cc0386.__obf_8f559e1327c9ee84.Sign() < 0 {
		return "-" + __obf_35bdb6d9709e131b
	}

	return __obf_35bdb6d9709e131b
}

func (__obf_df803c2772cc0386 *Decimal) __obf_a6a609bb966fabfb() {
	if __obf_df803c2772cc0386.__obf_8f559e1327c9ee84 == nil {
		__obf_df803c2772cc0386.__obf_8f559e1327c9ee84 = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_18a01609ffa3a43d Decimal, __obf_07623eecd9e07bf0 ...Decimal) Decimal {
	__obf_be96e7ce4d860194 := __obf_18a01609ffa3a43d
	for _, __obf_2319a318d53adef2 := range __obf_07623eecd9e07bf0 {
		if __obf_2319a318d53adef2.Cmp(__obf_be96e7ce4d860194) < 0 {
			__obf_be96e7ce4d860194 = __obf_2319a318d53adef2
		}
	}
	return __obf_be96e7ce4d860194
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_18a01609ffa3a43d Decimal, __obf_07623eecd9e07bf0 ...Decimal) Decimal {
	__obf_be96e7ce4d860194 := __obf_18a01609ffa3a43d
	for _, __obf_2319a318d53adef2 := range __obf_07623eecd9e07bf0 {
		if __obf_2319a318d53adef2.Cmp(__obf_be96e7ce4d860194) > 0 {
			__obf_be96e7ce4d860194 = __obf_2319a318d53adef2
		}
	}
	return __obf_be96e7ce4d860194
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_18a01609ffa3a43d Decimal, __obf_07623eecd9e07bf0 ...Decimal) Decimal {
	__obf_77a0e9752004cc30 := __obf_18a01609ffa3a43d
	for _, __obf_2319a318d53adef2 := range __obf_07623eecd9e07bf0 {
		__obf_77a0e9752004cc30 = __obf_77a0e9752004cc30.Add(__obf_2319a318d53adef2)
	}

	return __obf_77a0e9752004cc30
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_18a01609ffa3a43d Decimal, __obf_07623eecd9e07bf0 ...Decimal) Decimal {
	__obf_c9ac136623c99281 := New(int64(len(__obf_07623eecd9e07bf0)+1), 0)
	__obf_ee1eb81404c3703d := Sum(__obf_18a01609ffa3a43d, __obf_07623eecd9e07bf0...)
	return __obf_ee1eb81404c3703d.Div(__obf_c9ac136623c99281)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_11a00294e6b5e174 Decimal, __obf_b6bc08d52d165029 Decimal) (Decimal, Decimal) {
	__obf_11a00294e6b5e174.__obf_a6a609bb966fabfb()
	__obf_b6bc08d52d165029.__obf_a6a609bb966fabfb()

	if __obf_11a00294e6b5e174.__obf_97e810ed39f8ac3f < __obf_b6bc08d52d165029.__obf_97e810ed39f8ac3f {
		return __obf_11a00294e6b5e174, __obf_b6bc08d52d165029.__obf_749ebd7071bc390b(__obf_11a00294e6b5e174.__obf_97e810ed39f8ac3f)
	} else if __obf_11a00294e6b5e174.__obf_97e810ed39f8ac3f > __obf_b6bc08d52d165029.__obf_97e810ed39f8ac3f {
		return __obf_11a00294e6b5e174.__obf_749ebd7071bc390b(__obf_b6bc08d52d165029.__obf_97e810ed39f8ac3f), __obf_b6bc08d52d165029
	}

	return __obf_11a00294e6b5e174, __obf_b6bc08d52d165029
}

func __obf_e7f7278a29822adc(__obf_8f559e1327c9ee84 any) (string, error) {
	var __obf_dad74d30783b6a95 []byte

	switch __obf_7c58e6101912d0e0 := __obf_8f559e1327c9ee84.(type) {
	case string:
		__obf_dad74d30783b6a95 = []byte(__obf_7c58e6101912d0e0)
	case []byte:
		__obf_dad74d30783b6a95 = __obf_7c58e6101912d0e0
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_8f559e1327c9ee84,

			// If the amount is quoted, strip the quotes
			__obf_8f559e1327c9ee84)
	}

	if len(__obf_dad74d30783b6a95) > 2 && __obf_dad74d30783b6a95[0] == '"' && __obf_dad74d30783b6a95[len(__obf_dad74d30783b6a95)-1] == '"' {
		__obf_dad74d30783b6a95 = __obf_dad74d30783b6a95[1 : len(__obf_dad74d30783b6a95)-1]
	}
	return string(__obf_dad74d30783b6a95), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the datautil.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_df803c2772cc0386 Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_df803c2772cc0386,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_df803c2772cc0386 *NullDecimal) Scan(__obf_8f559e1327c9ee84 any) error {
	if __obf_8f559e1327c9ee84 == nil {
		__obf_df803c2772cc0386.
			Valid = false
		return nil
	}
	__obf_df803c2772cc0386.
		Valid = true
	return __obf_df803c2772cc0386.Decimal.Scan(__obf_8f559e1327c9ee84)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_df803c2772cc0386 NullDecimal) Value() (driver.Value, error) {
	if !__obf_df803c2772cc0386.Valid {
		return nil, nil
	}
	return __obf_df803c2772cc0386.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_df803c2772cc0386 *NullDecimal) UnmarshalJSON(__obf_7b9bd66f12aa0215 []byte) error {
	if string(__obf_7b9bd66f12aa0215) == "null" {
		__obf_df803c2772cc0386.
			Valid = false
		return nil
	}
	__obf_df803c2772cc0386.
		Valid = true
	return __obf_df803c2772cc0386.Decimal.UnmarshalJSON(__obf_7b9bd66f12aa0215)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_df803c2772cc0386 NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_df803c2772cc0386.Valid {
		return []byte("null"), nil
	}
	return __obf_df803c2772cc0386.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_df803c2772cc0386 *NullDecimal) UnmarshalText(__obf_4104aedfa89212c3 []byte) error {
	__obf_b3226814fd131eda := string(__obf_4104aedfa89212c3)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_b3226814fd131eda == "" {
		__obf_df803c2772cc0386.
			Valid = false
		return nil
	}
	if __obf_b2a10ebf0533880b := __obf_df803c2772cc0386.Decimal.UnmarshalText(__obf_4104aedfa89212c3); __obf_b2a10ebf0533880b != nil {
		__obf_df803c2772cc0386.
			Valid = false
		return __obf_b2a10ebf0533880b
	}
	__obf_df803c2772cc0386.
		Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_df803c2772cc0386 NullDecimal) MarshalText() (__obf_4104aedfa89212c3 []byte, __obf_b2a10ebf0533880b error) {
	if !__obf_df803c2772cc0386.Valid {
		return []byte{}, nil
	}
	return __obf_df803c2772cc0386.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_df803c2772cc0386 Decimal) Atan() Decimal {
	if __obf_df803c2772cc0386.Equal(NewFromFloat(0.0)) {
		return __obf_df803c2772cc0386
	}
	if __obf_df803c2772cc0386.GreaterThan(NewFromFloat(0.0)) {
		return __obf_df803c2772cc0386.__obf_39e16b36a5f2d23e()
	}
	return __obf_df803c2772cc0386.Neg().__obf_39e16b36a5f2d23e().Neg()
}

func (__obf_df803c2772cc0386 Decimal) __obf_e0604005c6e22ac3() Decimal {
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
	__obf_9f51474a44562371 := __obf_df803c2772cc0386.Mul(__obf_df803c2772cc0386)
	__obf_d75120ce94ea6947 := P0.Mul(__obf_9f51474a44562371).Add(P1).Mul(__obf_9f51474a44562371).Add(P2).Mul(__obf_9f51474a44562371).Add(P3).Mul(__obf_9f51474a44562371).Add(P4).Mul(__obf_9f51474a44562371)
	__obf_2d56f42d61b43605 := __obf_9f51474a44562371.Add(Q0).Mul(__obf_9f51474a44562371).Add(Q1).Mul(__obf_9f51474a44562371).Add(Q2).Mul(__obf_9f51474a44562371).Add(Q3).Mul(__obf_9f51474a44562371).Add(Q4)
	__obf_9f51474a44562371 = __obf_d75120ce94ea6947.Div(__obf_2d56f42d61b43605)
	__obf_9f51474a44562371 = __obf_df803c2772cc0386.Mul(__obf_9f51474a44562371).Add(__obf_df803c2772cc0386)
	return __obf_9f51474a44562371
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_df803c2772cc0386 Decimal) __obf_39e16b36a5f2d23e() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)
	__obf_c8325a326585f9fb := // tan(3*pi/8)
		NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_df803c2772cc0386.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_df803c2772cc0386.__obf_e0604005c6e22ac3()
	}
	if __obf_df803c2772cc0386.GreaterThan(Tan3pio8) {
		return __obf_c8325a326585f9fb.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_df803c2772cc0386).__obf_e0604005c6e22ac3()).Add(Morebits)
	}
	return __obf_c8325a326585f9fb.Div(NewFromFloat(4.0)).Add((__obf_df803c2772cc0386.Sub(NewFromFloat(1.0)).Div(__obf_df803c2772cc0386.Add(NewFromFloat(1.0)))).__obf_e0604005c6e22ac3()).Add(NewFromFloat(0.5).Mul(Morebits))
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
func (__obf_df803c2772cc0386 Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_df803c2772cc0386.Equal(NewFromFloat(0.0)) {
		return __obf_df803c2772cc0386
	}
	__obf_cf03f6d30f3621c5 := // make argument positive but save the sign
		false
	if __obf_df803c2772cc0386.LessThan(NewFromFloat(0.0)) {
		__obf_df803c2772cc0386 = __obf_df803c2772cc0386.Neg()
		__obf_cf03f6d30f3621c5 = true
	}
	__obf_a3bbf189b761f48c := __obf_df803c2772cc0386.Mul(M4PI).IntPart()
	__obf_9c398b8b5eba5c9d := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_a3bbf189b761f48c)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_a3bbf189b761f48c&1 == 1 {
		__obf_a3bbf189b761f48c++
		__obf_9c398b8b5eba5c9d = __obf_9c398b8b5eba5c9d.Add(NewFromFloat(1.0))
	}
	__obf_a3bbf189b761f48c &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_a3bbf189b761f48c > 3 {
		__obf_cf03f6d30f3621c5 = !__obf_cf03f6d30f3621c5
		__obf_a3bbf189b761f48c -= 4
	}
	__obf_9f51474a44562371 := __obf_df803c2772cc0386.Sub(__obf_9c398b8b5eba5c9d.Mul(PI4A)).Sub(__obf_9c398b8b5eba5c9d.Mul(PI4B)).Sub(__obf_9c398b8b5eba5c9d.Mul(PI4C))
	__obf_c4a0445e9148b061 := // Extended precision modular arithmetic
		__obf_9f51474a44562371.Mul(__obf_9f51474a44562371)

	if __obf_a3bbf189b761f48c == 1 || __obf_a3bbf189b761f48c == 2 {
		__obf_ab99dbfaf01372b8 := __obf_c4a0445e9148b061.Mul(__obf_c4a0445e9148b061).Mul(_cos[0].Mul(__obf_c4a0445e9148b061).Add(_cos[1]).Mul(__obf_c4a0445e9148b061).Add(_cos[2]).Mul(__obf_c4a0445e9148b061).Add(_cos[3]).Mul(__obf_c4a0445e9148b061).Add(_cos[4]).Mul(__obf_c4a0445e9148b061).Add(_cos[5]))
		__obf_9c398b8b5eba5c9d = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_c4a0445e9148b061)).Add(__obf_ab99dbfaf01372b8)
	} else {
		__obf_9c398b8b5eba5c9d = __obf_9f51474a44562371.Add(__obf_9f51474a44562371.Mul(__obf_c4a0445e9148b061).Mul(_sin[0].Mul(__obf_c4a0445e9148b061).Add(_sin[1]).Mul(__obf_c4a0445e9148b061).Add(_sin[2]).Mul(__obf_c4a0445e9148b061).Add(_sin[3]).Mul(__obf_c4a0445e9148b061).Add(_sin[4]).Mul(__obf_c4a0445e9148b061).Add(_sin[5])))
	}
	if __obf_cf03f6d30f3621c5 {
		__obf_9c398b8b5eba5c9d = __obf_9c398b8b5eba5c9d.Neg()
	}
	return __obf_9c398b8b5eba5c9d
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
func (__obf_df803c2772cc0386 Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)  // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)  // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15) // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125)
	__obf_cf03f6d30f3621c5 := // 4/pi

		// make argument positive
		false
	if __obf_df803c2772cc0386.LessThan(NewFromFloat(0.0)) {
		__obf_df803c2772cc0386 = __obf_df803c2772cc0386.Neg()
	}
	__obf_a3bbf189b761f48c := __obf_df803c2772cc0386.Mul(M4PI).IntPart()
	__obf_9c398b8b5eba5c9d := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_a3bbf189b761f48c)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_a3bbf189b761f48c&1 == 1 {
		__obf_a3bbf189b761f48c++
		__obf_9c398b8b5eba5c9d = __obf_9c398b8b5eba5c9d.Add(NewFromFloat(1.0))
	}
	__obf_a3bbf189b761f48c &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_a3bbf189b761f48c > 3 {
		__obf_cf03f6d30f3621c5 = !__obf_cf03f6d30f3621c5
		__obf_a3bbf189b761f48c -= 4
	}
	if __obf_a3bbf189b761f48c > 1 {
		__obf_cf03f6d30f3621c5 = !__obf_cf03f6d30f3621c5
	}
	__obf_9f51474a44562371 := __obf_df803c2772cc0386.Sub(__obf_9c398b8b5eba5c9d.Mul(PI4A)).Sub(__obf_9c398b8b5eba5c9d.Mul(PI4B)).Sub(__obf_9c398b8b5eba5c9d.Mul(PI4C))
	__obf_c4a0445e9148b061 := // Extended precision modular arithmetic
		__obf_9f51474a44562371.Mul(__obf_9f51474a44562371)

	if __obf_a3bbf189b761f48c == 1 || __obf_a3bbf189b761f48c == 2 {
		__obf_9c398b8b5eba5c9d = __obf_9f51474a44562371.Add(__obf_9f51474a44562371.Mul(__obf_c4a0445e9148b061).Mul(_sin[0].Mul(__obf_c4a0445e9148b061).Add(_sin[1]).Mul(__obf_c4a0445e9148b061).Add(_sin[2]).Mul(__obf_c4a0445e9148b061).Add(_sin[3]).Mul(__obf_c4a0445e9148b061).Add(_sin[4]).Mul(__obf_c4a0445e9148b061).Add(_sin[5])))
	} else {
		__obf_ab99dbfaf01372b8 := __obf_c4a0445e9148b061.Mul(__obf_c4a0445e9148b061).Mul(_cos[0].Mul(__obf_c4a0445e9148b061).Add(_cos[1]).Mul(__obf_c4a0445e9148b061).Add(_cos[2]).Mul(__obf_c4a0445e9148b061).Add(_cos[3]).Mul(__obf_c4a0445e9148b061).Add(_cos[4]).Mul(__obf_c4a0445e9148b061).Add(_cos[5]))
		__obf_9c398b8b5eba5c9d = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_c4a0445e9148b061)).Add(__obf_ab99dbfaf01372b8)
	}
	if __obf_cf03f6d30f3621c5 {
		__obf_9c398b8b5eba5c9d = __obf_9c398b8b5eba5c9d.Neg()
	}
	return __obf_9c398b8b5eba5c9d
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
func (__obf_df803c2772cc0386 Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_df803c2772cc0386.Equal(NewFromFloat(0.0)) {
		return __obf_df803c2772cc0386
	}
	__obf_cf03f6d30f3621c5 := // make argument positive but save the sign
		false
	if __obf_df803c2772cc0386.LessThan(NewFromFloat(0.0)) {
		__obf_df803c2772cc0386 = __obf_df803c2772cc0386.Neg()
		__obf_cf03f6d30f3621c5 = true
	}
	__obf_a3bbf189b761f48c := __obf_df803c2772cc0386.Mul(M4PI).IntPart()
	__obf_9c398b8b5eba5c9d := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_a3bbf189b761f48c)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_a3bbf189b761f48c&1 == 1 {
		__obf_a3bbf189b761f48c++
		__obf_9c398b8b5eba5c9d = __obf_9c398b8b5eba5c9d.Add(NewFromFloat(1.0))
	}
	__obf_9f51474a44562371 := __obf_df803c2772cc0386.Sub(__obf_9c398b8b5eba5c9d.Mul(PI4A)).Sub(__obf_9c398b8b5eba5c9d.Mul(PI4B)).Sub(__obf_9c398b8b5eba5c9d.Mul(PI4C))
	__obf_c4a0445e9148b061 := // Extended precision modular arithmetic
		__obf_9f51474a44562371.Mul(__obf_9f51474a44562371)

	if __obf_c4a0445e9148b061.GreaterThan(NewFromFloat(1e-14)) {
		__obf_ab99dbfaf01372b8 := __obf_c4a0445e9148b061.Mul(_tanP[0].Mul(__obf_c4a0445e9148b061).Add(_tanP[1]).Mul(__obf_c4a0445e9148b061).Add(_tanP[2]))
		__obf_5ac346f5c5af0973 := __obf_c4a0445e9148b061.Add(_tanQ[1]).Mul(__obf_c4a0445e9148b061).Add(_tanQ[2]).Mul(__obf_c4a0445e9148b061).Add(_tanQ[3]).Mul(__obf_c4a0445e9148b061).Add(_tanQ[4])
		__obf_9c398b8b5eba5c9d = __obf_9f51474a44562371.Add(__obf_9f51474a44562371.Mul(__obf_ab99dbfaf01372b8.Div(__obf_5ac346f5c5af0973)))
	} else {
		__obf_9c398b8b5eba5c9d = __obf_9f51474a44562371
	}
	if __obf_a3bbf189b761f48c&2 == 2 {
		__obf_9c398b8b5eba5c9d = NewFromFloat(-1.0).Div(__obf_9c398b8b5eba5c9d)
	}
	if __obf_cf03f6d30f3621c5 {
		__obf_9c398b8b5eba5c9d = __obf_9c398b8b5eba5c9d.Neg()
	}
	return __obf_9c398b8b5eba5c9d
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

type __obf_b0e21776b9d45e13 struct {
	__obf_df803c2772cc0386 [800]byte
	__obf_2841543582033a23 int  // digits, big-endian representation
	__obf_b4c2d4c5eeb9a307 int  // number of digits used
	__obf_d3daa27031f7f18b bool // decimal point
	__obf_36cee10909116c00 bool // negative flag
	// discarded nonzero digits beyond d[:nd]
}

func (__obf_ca960a67e11dc734 *__obf_b0e21776b9d45e13) String() string {
	__obf_519018557784958f := 10 + __obf_ca960a67e11dc734.__obf_2841543582033a23
	if __obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307 > 0 {
		__obf_519018557784958f += __obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307
	}
	if __obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307 < 0 {
		__obf_519018557784958f += -__obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307
	}
	__obf_b66375793300bef9 := make([]byte, __obf_519018557784958f)
	__obf_ab99dbfaf01372b8 := 0
	switch {
	case __obf_ca960a67e11dc734.__obf_2841543582033a23 == 0:
		return "0"

	case __obf_ca960a67e11dc734.
		// zeros fill space between decimal point and digits
		__obf_b4c2d4c5eeb9a307 <= 0:
		__obf_b66375793300bef9[__obf_ab99dbfaf01372b8] = '0'
		__obf_ab99dbfaf01372b8++
		__obf_b66375793300bef9[__obf_ab99dbfaf01372b8] = '.'
		__obf_ab99dbfaf01372b8++
		__obf_ab99dbfaf01372b8 += __obf_55d39280c974f286(__obf_b66375793300bef9[__obf_ab99dbfaf01372b8 : __obf_ab99dbfaf01372b8+-__obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307])
		__obf_ab99dbfaf01372b8 += copy(__obf_b66375793300bef9[__obf_ab99dbfaf01372b8:], __obf_ca960a67e11dc734.__obf_df803c2772cc0386[0:__obf_ca960a67e11dc734.__obf_2841543582033a23])

	case __obf_ca960a67e11dc734.
		// decimal point in middle of digits
		__obf_b4c2d4c5eeb9a307 < __obf_ca960a67e11dc734.__obf_2841543582033a23:
		__obf_ab99dbfaf01372b8 += copy(__obf_b66375793300bef9[__obf_ab99dbfaf01372b8:], __obf_ca960a67e11dc734.__obf_df803c2772cc0386[0:__obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307])
		__obf_b66375793300bef9[__obf_ab99dbfaf01372b8] = '.'
		__obf_ab99dbfaf01372b8++
		__obf_ab99dbfaf01372b8 += copy(__obf_b66375793300bef9[__obf_ab99dbfaf01372b8:], __obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307:

		// zeros fill space between digits and decimal point
		__obf_ca960a67e11dc734.__obf_2841543582033a23])

	default:
		__obf_ab99dbfaf01372b8 += copy(__obf_b66375793300bef9[__obf_ab99dbfaf01372b8:], __obf_ca960a67e11dc734.__obf_df803c2772cc0386[0:__obf_ca960a67e11dc734.__obf_2841543582033a23])
		__obf_ab99dbfaf01372b8 += __obf_55d39280c974f286(__obf_b66375793300bef9[__obf_ab99dbfaf01372b8 : __obf_ab99dbfaf01372b8+__obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307-__obf_ca960a67e11dc734.__obf_2841543582033a23])
	}
	return string(__obf_b66375793300bef9[0:__obf_ab99dbfaf01372b8])
}

func __obf_55d39280c974f286(__obf_7fa0b91c9ec42af5 []byte) int {
	for __obf_0038051e77a34ba0 := range __obf_7fa0b91c9ec42af5 {
		__obf_7fa0b91c9ec42af5[__obf_0038051e77a34ba0] = '0'
	}
	return len(__obf_7fa0b91c9ec42af5)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_6cb7846618aff8a5(__obf_ca960a67e11dc734 *__obf_b0e21776b9d45e13) {
	for __obf_ca960a67e11dc734.__obf_2841543582033a23 > 0 && __obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_ca960a67e11dc734.__obf_2841543582033a23-1] == '0' {
		__obf_ca960a67e11dc734.__obf_2841543582033a23--
	}
	if __obf_ca960a67e11dc734.__obf_2841543582033a23 == 0 {
		__obf_ca960a67e11dc734.

			// Assign v to a.
			__obf_b4c2d4c5eeb9a307 = 0
	}
}

func (__obf_ca960a67e11dc734 *__obf_b0e21776b9d45e13) Assign(__obf_7c58e6101912d0e0 uint64) {
	var __obf_b66375793300bef9 [24]byte
	__obf_519018557784958f := // Write reversed decimal in buf.
		0
	for __obf_7c58e6101912d0e0 > 0 {
		__obf_90989c4ac9f58746 := __obf_7c58e6101912d0e0 / 10
		__obf_7c58e6101912d0e0 -= 10 * __obf_90989c4ac9f58746
		__obf_b66375793300bef9[__obf_519018557784958f] = byte(__obf_7c58e6101912d0e0 + '0')
		__obf_519018557784958f++
		__obf_7c58e6101912d0e0 = __obf_90989c4ac9f58746
	}
	__obf_ca960a67e11dc734.

		// Reverse again to produce forward decimal in a.d.
		__obf_2841543582033a23 = 0
	for __obf_519018557784958f--; __obf_519018557784958f >= 0; __obf_519018557784958f-- {
		__obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_ca960a67e11dc734.__obf_2841543582033a23] = __obf_b66375793300bef9[__obf_519018557784958f]
		__obf_ca960a67e11dc734.__obf_2841543582033a23++
	}
	__obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307 = __obf_ca960a67e11dc734.

		// Maximum shift that we can do in one pass without overflow.
		// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
		__obf_2841543582033a23
	__obf_6cb7846618aff8a5(__obf_ca960a67e11dc734)
}

const __obf_744c254c10b6483a = 32 << (^uint(0) >> 63)
const __obf_64f20795be4a8a8c = __obf_744c254c10b6483a - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_acd3d09b9944a451(__obf_ca960a67e11dc734 *__obf_b0e21776b9d45e13, __obf_fb209a5c6c42bbe3 uint) {
	__obf_69a72b316864499d := 0
	__obf_ab99dbfaf01372b8 := // read pointer
		0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_519018557784958f uint
	for ; __obf_519018557784958f>>__obf_fb209a5c6c42bbe3 == 0; __obf_69a72b316864499d++ {
		if __obf_69a72b316864499d >= __obf_ca960a67e11dc734.__obf_2841543582033a23 {
			if __obf_519018557784958f == 0 {
				__obf_ca960a67e11dc734.
					// a == 0; shouldn't get here, but handle anyway.
					__obf_2841543582033a23 = 0
				return
			}
			for __obf_519018557784958f>>__obf_fb209a5c6c42bbe3 == 0 {
				__obf_519018557784958f = __obf_519018557784958f * 10
				__obf_69a72b316864499d++
			}
			break
		}
		__obf_253219491e4be33b := uint(__obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_69a72b316864499d])
		__obf_519018557784958f = __obf_519018557784958f*10 + __obf_253219491e4be33b - '0'
	}
	__obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307 -= __obf_69a72b316864499d - 1

	var __obf_aebc47ae3b3b0f57 uint = (1 << __obf_fb209a5c6c42bbe3) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_69a72b316864499d < __obf_ca960a67e11dc734.__obf_2841543582033a23; __obf_69a72b316864499d++ {
		__obf_253219491e4be33b := uint(__obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_69a72b316864499d])
		__obf_efddb1169d85408f := __obf_519018557784958f >> __obf_fb209a5c6c42bbe3
		__obf_519018557784958f &= __obf_aebc47ae3b3b0f57
		__obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_ab99dbfaf01372b8] = byte(__obf_efddb1169d85408f + '0')
		__obf_ab99dbfaf01372b8++
		__obf_519018557784958f = __obf_519018557784958f*10 + __obf_253219491e4be33b - '0'
	}

	// Put down extra digits.
	for __obf_519018557784958f > 0 {
		__obf_efddb1169d85408f := __obf_519018557784958f >> __obf_fb209a5c6c42bbe3
		__obf_519018557784958f &= __obf_aebc47ae3b3b0f57
		if __obf_ab99dbfaf01372b8 < len(__obf_ca960a67e11dc734.__obf_df803c2772cc0386) {
			__obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_ab99dbfaf01372b8] = byte(__obf_efddb1169d85408f + '0')
			__obf_ab99dbfaf01372b8++
		} else if __obf_efddb1169d85408f > 0 {
			__obf_ca960a67e11dc734.__obf_36cee10909116c00 = true
		}
		__obf_519018557784958f = __obf_519018557784958f * 10
	}
	__obf_ca960a67e11dc734.__obf_2841543582033a23 = __obf_ab99dbfaf01372b8

	// Cheat sheet for left shift: table indexed by shift count giving
	// number of new digits that will be introduced by that shift.
	//
	// For example, leftcheats[4] = {2, "625"}.  That means that
	// if we are shifting by 4 (multiplying by 16), it will add 2 digits
	// when the string prefix is "625" through "999", and one fewer digit
	// if the string prefix is "000" through "624".
	//
	// Credit for this trick goes to Ken.
	__obf_6cb7846618aff8a5(__obf_ca960a67e11dc734)
}

type __obf_bf03e9327b751a61 struct {
	__obf_605d6af40f0d831b int
	__obf_9f9268373a114d09 string // number of new digits
	// minus one digit if original < a.
}

var __obf_5b9aab3c7aa5a764 = []__obf_bf03e9327b751a61{
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
func __obf_a9e596524245a3be(__obf_a46556e609360f0e []byte, __obf_5a851afcb64b57aa string) bool {
	for __obf_0038051e77a34ba0 := 0; __obf_0038051e77a34ba0 < len(__obf_5a851afcb64b57aa); __obf_0038051e77a34ba0++ {
		if __obf_0038051e77a34ba0 >= len(__obf_a46556e609360f0e) {
			return true
		}
		if __obf_a46556e609360f0e[__obf_0038051e77a34ba0] != __obf_5a851afcb64b57aa[__obf_0038051e77a34ba0] {
			return __obf_a46556e609360f0e[__obf_0038051e77a34ba0] < __obf_5a851afcb64b57aa[__obf_0038051e77a34ba0]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_7491abcbcc2008e2(__obf_ca960a67e11dc734 *__obf_b0e21776b9d45e13, __obf_fb209a5c6c42bbe3 uint) {
	__obf_605d6af40f0d831b := __obf_5b9aab3c7aa5a764[__obf_fb209a5c6c42bbe3].__obf_605d6af40f0d831b
	if __obf_a9e596524245a3be(__obf_ca960a67e11dc734.__obf_df803c2772cc0386[0:__obf_ca960a67e11dc734.__obf_2841543582033a23], __obf_5b9aab3c7aa5a764[__obf_fb209a5c6c42bbe3].__obf_9f9268373a114d09) {
		__obf_605d6af40f0d831b--
	}
	__obf_69a72b316864499d := __obf_ca960a67e11dc734. // read index
								__obf_2841543582033a23
	__obf_ab99dbfaf01372b8 := __obf_ca960a67e11dc734. // write index
								__obf_2841543582033a23 + __obf_605d6af40f0d831b

	// Pick up a digit, put down a digit.
	var __obf_519018557784958f uint
	for __obf_69a72b316864499d--; __obf_69a72b316864499d >= 0; __obf_69a72b316864499d-- {
		__obf_519018557784958f += (uint(__obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_69a72b316864499d]) - '0') << __obf_fb209a5c6c42bbe3
		__obf_03ba54e1ad1f8005 := __obf_519018557784958f / 10
		__obf_ba39bde05a96fe04 := __obf_519018557784958f - 10*__obf_03ba54e1ad1f8005
		__obf_ab99dbfaf01372b8--
		if __obf_ab99dbfaf01372b8 < len(__obf_ca960a67e11dc734.__obf_df803c2772cc0386) {
			__obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_ab99dbfaf01372b8] = byte(__obf_ba39bde05a96fe04 + '0')
		} else if __obf_ba39bde05a96fe04 != 0 {
			__obf_ca960a67e11dc734.__obf_36cee10909116c00 = true
		}
		__obf_519018557784958f = __obf_03ba54e1ad1f8005
	}

	// Put down extra digits.
	for __obf_519018557784958f > 0 {
		__obf_03ba54e1ad1f8005 := __obf_519018557784958f / 10
		__obf_ba39bde05a96fe04 := __obf_519018557784958f - 10*__obf_03ba54e1ad1f8005
		__obf_ab99dbfaf01372b8--
		if __obf_ab99dbfaf01372b8 < len(__obf_ca960a67e11dc734.__obf_df803c2772cc0386) {
			__obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_ab99dbfaf01372b8] = byte(__obf_ba39bde05a96fe04 + '0')
		} else if __obf_ba39bde05a96fe04 != 0 {
			__obf_ca960a67e11dc734.__obf_36cee10909116c00 = true
		}
		__obf_519018557784958f = __obf_03ba54e1ad1f8005
	}
	__obf_ca960a67e11dc734.__obf_2841543582033a23 += __obf_605d6af40f0d831b
	if __obf_ca960a67e11dc734.__obf_2841543582033a23 >= len(__obf_ca960a67e11dc734.__obf_df803c2772cc0386) {
		__obf_ca960a67e11dc734.__obf_2841543582033a23 = len(__obf_ca960a67e11dc734.__obf_df803c2772cc0386)
	}
	__obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307 += __obf_605d6af40f0d831b

	// Binary shift left (k > 0) or right (k < 0).
	__obf_6cb7846618aff8a5(__obf_ca960a67e11dc734)
}

func (__obf_ca960a67e11dc734 *__obf_b0e21776b9d45e13) Shift(__obf_fb209a5c6c42bbe3 int) {
	switch {
	case __obf_ca960a67e11dc734.
		// nothing to do: a == 0
		__obf_2841543582033a23 == 0:

	case __obf_fb209a5c6c42bbe3 > 0:
		for __obf_fb209a5c6c42bbe3 > __obf_64f20795be4a8a8c {
			__obf_7491abcbcc2008e2(__obf_ca960a67e11dc734, __obf_64f20795be4a8a8c)
			__obf_fb209a5c6c42bbe3 -= __obf_64f20795be4a8a8c
		}
		__obf_7491abcbcc2008e2(__obf_ca960a67e11dc734, uint(__obf_fb209a5c6c42bbe3))
	case __obf_fb209a5c6c42bbe3 < 0:
		for __obf_fb209a5c6c42bbe3 < -__obf_64f20795be4a8a8c {
			__obf_acd3d09b9944a451(__obf_ca960a67e11dc734, __obf_64f20795be4a8a8c)
			__obf_fb209a5c6c42bbe3 += __obf_64f20795be4a8a8c
		}
		__obf_acd3d09b9944a451(__obf_ca960a67e11dc734, uint(-__obf_fb209a5c6c42bbe3))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_771537d76950cee5(__obf_ca960a67e11dc734 *__obf_b0e21776b9d45e13, __obf_2841543582033a23 int) bool {
	if __obf_2841543582033a23 < 0 || __obf_2841543582033a23 >= __obf_ca960a67e11dc734.__obf_2841543582033a23 {
		return false
	}
	if __obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_2841543582033a23] == '5' && __obf_2841543582033a23+1 == __obf_ca960a67e11dc734. // exactly halfway - round to even
																		__obf_2841543582033a23 {
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_ca960a67e11dc734.__obf_36cee10909116c00 {
			return true
		}
		return __obf_2841543582033a23 > 0 && (__obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_2841543582033a23-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_ca960a67e11dc734.

		// Round a to nd digits (or fewer).
		// If nd is zero, it means we're rounding
		// just to the left of the digits, as in
		// 0.09 -> 0.1.
		__obf_df803c2772cc0386[__obf_2841543582033a23] >= '5'
}

func (__obf_ca960a67e11dc734 *__obf_b0e21776b9d45e13) Round(__obf_2841543582033a23 int) {
	if __obf_2841543582033a23 < 0 || __obf_2841543582033a23 >= __obf_ca960a67e11dc734.__obf_2841543582033a23 {
		return
	}
	if __obf_771537d76950cee5(__obf_ca960a67e11dc734, __obf_2841543582033a23) {
		__obf_ca960a67e11dc734.
			RoundUp(__obf_2841543582033a23)
	} else {
		__obf_ca960a67e11dc734.
			RoundDown(__obf_2841543582033a23)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_ca960a67e11dc734 *__obf_b0e21776b9d45e13) RoundDown(__obf_2841543582033a23 int) {
	if __obf_2841543582033a23 < 0 || __obf_2841543582033a23 >= __obf_ca960a67e11dc734.__obf_2841543582033a23 {
		return
	}
	__obf_ca960a67e11dc734.__obf_2841543582033a23 = __obf_2841543582033a23

	// Round a up to nd digits (or fewer).
	__obf_6cb7846618aff8a5(__obf_ca960a67e11dc734)
}

func (__obf_ca960a67e11dc734 *__obf_b0e21776b9d45e13) RoundUp(__obf_2841543582033a23 int) {
	if __obf_2841543582033a23 < 0 || __obf_2841543582033a23 >= __obf_ca960a67e11dc734.

		// round up
		__obf_2841543582033a23 {
		return
	}

	for __obf_0038051e77a34ba0 := __obf_2841543582033a23 - 1; __obf_0038051e77a34ba0 >= 0; __obf_0038051e77a34ba0-- {
		__obf_253219491e4be33b := __obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_0038051e77a34ba0]
		if __obf_253219491e4be33b < '9' {
			__obf_ca960a67e11dc734. // can stop after this digit
						__obf_df803c2772cc0386[__obf_0038051e77a34ba0]++
			__obf_ca960a67e11dc734.__obf_2841543582033a23 = __obf_0038051e77a34ba0 + 1
			return
		}
	}
	__obf_ca960a67e11dc734.

		// Number is all 9s.
		// Change to single 1 with adjusted decimal point.
		__obf_df803c2772cc0386[0] = '1'
	__obf_ca960a67e11dc734.

		// Extract integer part, rounded appropriately.
		// No guarantees about overflow.
		__obf_2841543582033a23 = 1
	__obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307++
}

func (__obf_ca960a67e11dc734 *__obf_b0e21776b9d45e13) RoundedInteger() uint64 {
	if __obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307 > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_0038051e77a34ba0 int
	__obf_519018557784958f := uint64(0)
	for __obf_0038051e77a34ba0 = 0; __obf_0038051e77a34ba0 < __obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307 && __obf_0038051e77a34ba0 < __obf_ca960a67e11dc734.__obf_2841543582033a23; __obf_0038051e77a34ba0++ {
		__obf_519018557784958f = __obf_519018557784958f*10 + uint64(__obf_ca960a67e11dc734.__obf_df803c2772cc0386[__obf_0038051e77a34ba0]-'0')
	}
	for ; __obf_0038051e77a34ba0 < __obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307; __obf_0038051e77a34ba0++ {
		__obf_519018557784958f *= 10
	}
	if __obf_771537d76950cee5(__obf_ca960a67e11dc734, __obf_ca960a67e11dc734.__obf_b4c2d4c5eeb9a307) {
		__obf_519018557784958f++
	}
	return __obf_519018557784958f
}
