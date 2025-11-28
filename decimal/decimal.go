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
package __obf_af53385fca67d325

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

var __obf_ed0e7e529c925e39 = big.NewInt(0)
var __obf_acf4ecfba54e0451 = big.NewInt(1)
var __obf_7f85275c7a28a30a = big.NewInt(2)
var __obf_c43b7c699e959605 = big.NewInt(4)
var __obf_0fdf935c5e6da9d3 = big.NewInt(5)
var __obf_125c1d0eb0d00860 = big.NewInt(10)
var __obf_a370817798d4bdc4 = big.NewInt(20)

var __obf_d0eb5ea33f7ee79a = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_830e351d96ac1434 *big.Int

	// NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.
	__obf_e7725bcd1ee3b1c3 int32
}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_830e351d96ac1434 int64, __obf_e7725bcd1ee3b1c3 int32) Decimal {
	return Decimal{
		__obf_830e351d96ac1434: big.NewInt(__obf_830e351d96ac1434),
		__obf_e7725bcd1ee3b1c3: __obf_e7725bcd1ee3b1c3,
	}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_830e351d96ac1434 int64) Decimal {
	return Decimal{
		__obf_830e351d96ac1434: big.NewInt(__obf_830e351d96ac1434),
		__obf_e7725bcd1ee3b1c3: 0,
	}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_830e351d96ac1434 int32) Decimal {
	return Decimal{
		__obf_830e351d96ac1434: big.NewInt(int64(__obf_830e351d96ac1434)),
		__obf_e7725bcd1ee3b1c3: 0,
	}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_830e351d96ac1434 uint64) Decimal {
	return Decimal{
		__obf_830e351d96ac1434: new(big.Int).SetUint64(__obf_830e351d96ac1434),
		__obf_e7725bcd1ee3b1c3: 0,
	}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_830e351d96ac1434 *big.Int, __obf_e7725bcd1ee3b1c3 int32) Decimal {
	return Decimal{
		__obf_830e351d96ac1434: new(big.Int).Set(__obf_830e351d96ac1434),
		__obf_e7725bcd1ee3b1c3: __obf_e7725bcd1ee3b1c3,
	}
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
func NewFromBigRat(__obf_830e351d96ac1434 *big.Rat, __obf_a2144abdd719f349 int32) Decimal {
	return Decimal{
		__obf_830e351d96ac1434: new(big.Int).Set(__obf_830e351d96ac1434.Num()),
		__obf_e7725bcd1ee3b1c3: 0,
	}.DivRound(Decimal{
		__obf_830e351d96ac1434: new(big.Int).Set(__obf_830e351d96ac1434.Denom()),
		__obf_e7725bcd1ee3b1c3: 0,
	}, __obf_a2144abdd719f349)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_830e351d96ac1434 string) (Decimal, error) {
	__obf_711c6cab83504937 := __obf_830e351d96ac1434
	var __obf_f439f50f564dbf6f string
	var __obf_e7725bcd1ee3b1c3 int64

	// Check if number is using scientific notation
	__obf_78223a4bf32dfcf0 := strings.IndexAny(__obf_830e351d96ac1434, "Ee")
	if __obf_78223a4bf32dfcf0 != -1 {
		__obf_2d38e442cc75992b, __obf_95437c99969367dc := strconv.ParseInt(__obf_830e351d96ac1434[__obf_78223a4bf32dfcf0+1:], 10, 32)
		if __obf_95437c99969367dc != nil {
			if __obf_250e118ca646cc77, __obf_c26031e6d546c8ec := __obf_95437c99969367dc.(*strconv.NumError); __obf_c26031e6d546c8ec && __obf_250e118ca646cc77.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_830e351d96ac1434)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_830e351d96ac1434)
		}
		__obf_830e351d96ac1434 = __obf_830e351d96ac1434[:__obf_78223a4bf32dfcf0]
		__obf_e7725bcd1ee3b1c3 = __obf_2d38e442cc75992b
	}

	__obf_623524fdab8edd1d := -1
	__obf_9e9033047f843367 := len(__obf_830e351d96ac1434)
	for __obf_4aeff04bb586b943 := 0; __obf_4aeff04bb586b943 < __obf_9e9033047f843367; __obf_4aeff04bb586b943++ {
		if __obf_830e351d96ac1434[__obf_4aeff04bb586b943] == '.' {
			if __obf_623524fdab8edd1d > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_830e351d96ac1434)
			}
			__obf_623524fdab8edd1d = __obf_4aeff04bb586b943
		}
	}

	if __obf_623524fdab8edd1d == -1 {
		// There is no decimal point, we can just parse the original string as
		// an int
		__obf_f439f50f564dbf6f = __obf_830e351d96ac1434
	} else {
		if __obf_623524fdab8edd1d+1 < __obf_9e9033047f843367 {
			__obf_f439f50f564dbf6f = __obf_830e351d96ac1434[:__obf_623524fdab8edd1d] + __obf_830e351d96ac1434[__obf_623524fdab8edd1d+1:]
		} else {
			__obf_f439f50f564dbf6f = __obf_830e351d96ac1434[:__obf_623524fdab8edd1d]
		}
		__obf_2d38e442cc75992b := -len(__obf_830e351d96ac1434[__obf_623524fdab8edd1d+1:])
		__obf_e7725bcd1ee3b1c3 += int64(__obf_2d38e442cc75992b)
	}

	var __obf_c24121ecd509691f *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_f439f50f564dbf6f) <= 18 {
		__obf_9babc7f7b3f9914a, __obf_95437c99969367dc := strconv.ParseInt(__obf_f439f50f564dbf6f, 10, 64)
		if __obf_95437c99969367dc != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_830e351d96ac1434)
		}
		__obf_c24121ecd509691f = big.NewInt(__obf_9babc7f7b3f9914a)
	} else {
		__obf_c24121ecd509691f = new(big.Int)
		_, __obf_c26031e6d546c8ec := __obf_c24121ecd509691f.SetString(__obf_f439f50f564dbf6f, 10)
		if !__obf_c26031e6d546c8ec {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_830e351d96ac1434)
		}
	}

	if __obf_e7725bcd1ee3b1c3 < math.MinInt32 || __obf_e7725bcd1ee3b1c3 > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_711c6cab83504937)
	}

	return Decimal{
		__obf_830e351d96ac1434: __obf_c24121ecd509691f,
		__obf_e7725bcd1ee3b1c3: int32(__obf_e7725bcd1ee3b1c3),
	}, nil
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
func NewFromFormattedString(__obf_830e351d96ac1434 string, __obf_5b7fc5d5418a347d *regexp.Regexp) (Decimal, error) {
	__obf_a65d92bf32426a60 := __obf_5b7fc5d5418a347d.ReplaceAllString(__obf_830e351d96ac1434, "")
	__obf_c8d718debc5f6382, __obf_95437c99969367dc := NewFromString(__obf_a65d92bf32426a60)
	if __obf_95437c99969367dc != nil {
		return Decimal{}, __obf_95437c99969367dc
	}
	return __obf_c8d718debc5f6382, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_830e351d96ac1434 string) Decimal {
	__obf_e7bd985d39dcdcf8, __obf_95437c99969367dc := NewFromString(__obf_830e351d96ac1434)
	if __obf_95437c99969367dc != nil {
		panic(__obf_95437c99969367dc)
	}
	return __obf_e7bd985d39dcdcf8
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
func NewFromFloat(__obf_830e351d96ac1434 float64) Decimal {
	if __obf_830e351d96ac1434 == 0 {
		return New(0, 0)
	}
	return __obf_8e07336091d8c743(__obf_830e351d96ac1434, math.Float64bits(__obf_830e351d96ac1434), &__obf_0d845b58fe68fed9)
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
func NewFromFloat32(__obf_830e351d96ac1434 float32) Decimal {
	if __obf_830e351d96ac1434 == 0 {
		return New(0, 0)
	}
	// XOR is workaround for https://github.com/golang/go/issues/26285
	__obf_7e6fccb9375af12b := math.Float32bits(__obf_830e351d96ac1434) ^ 0x80808080
	return __obf_8e07336091d8c743(float64(__obf_830e351d96ac1434), uint64(__obf_7e6fccb9375af12b)^0x80808080, &__obf_e7a2312c3f2fd7ba)
}

func __obf_8e07336091d8c743(__obf_3ce5743c80ca37b3 float64, __obf_6e1762ca2200a050 uint64, __obf_0c69b8da42f5c255 *__obf_8f596f0fa62f0bb1) Decimal {
	if math.IsNaN(__obf_3ce5743c80ca37b3) || math.IsInf(__obf_3ce5743c80ca37b3, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_3ce5743c80ca37b3))
	}
	__obf_e7725bcd1ee3b1c3 := int(__obf_6e1762ca2200a050>>__obf_0c69b8da42f5c255.__obf_cdbd5d47d10193b8) & (1<<__obf_0c69b8da42f5c255.__obf_8259f1621247e8d5 - 1)
	__obf_2a9e5a4d768e50b1 := __obf_6e1762ca2200a050 & (uint64(1)<<__obf_0c69b8da42f5c255.__obf_cdbd5d47d10193b8 - 1)

	switch __obf_e7725bcd1ee3b1c3 {
	case 0:
		// denormalized
		__obf_e7725bcd1ee3b1c3++

	default:
		// add implicit top bit
		__obf_2a9e5a4d768e50b1 |= uint64(1) << __obf_0c69b8da42f5c255.__obf_cdbd5d47d10193b8
	}
	__obf_e7725bcd1ee3b1c3 += __obf_0c69b8da42f5c255.__obf_5ccd0fbc88864cea

	var __obf_c8d718debc5f6382 __obf_af53385fca67d325
	__obf_c8d718debc5f6382.Assign(__obf_2a9e5a4d768e50b1)
	__obf_c8d718debc5f6382.Shift(__obf_e7725bcd1ee3b1c3 - int(__obf_0c69b8da42f5c255.__obf_cdbd5d47d10193b8))
	__obf_c8d718debc5f6382.__obf_6bf80ec73009a641 = __obf_6e1762ca2200a050>>(__obf_0c69b8da42f5c255.__obf_8259f1621247e8d5+__obf_0c69b8da42f5c255.__obf_cdbd5d47d10193b8) != 0

	__obf_acb82e38a55277ae(&__obf_c8d718debc5f6382, __obf_2a9e5a4d768e50b1, __obf_e7725bcd1ee3b1c3, __obf_0c69b8da42f5c255)
	// If less than 19 digits, we can do calculation in an int64.
	if __obf_c8d718debc5f6382.__obf_b2da5b1e27e361da < 19 {
		__obf_247932334c18f488 := int64(0)
		__obf_4931dd9c612fb4d2 := int64(1)
		for __obf_4aeff04bb586b943 := __obf_c8d718debc5f6382.__obf_b2da5b1e27e361da - 1; __obf_4aeff04bb586b943 >= 0; __obf_4aeff04bb586b943-- {
			__obf_247932334c18f488 += __obf_4931dd9c612fb4d2 * int64(__obf_c8d718debc5f6382.__obf_c8d718debc5f6382[__obf_4aeff04bb586b943]-'0')
			__obf_4931dd9c612fb4d2 *= 10
		}
		if __obf_c8d718debc5f6382.__obf_6bf80ec73009a641 {
			__obf_247932334c18f488 *= -1
		}
		return Decimal{__obf_830e351d96ac1434: big.NewInt(__obf_247932334c18f488), __obf_e7725bcd1ee3b1c3: int32(__obf_c8d718debc5f6382.__obf_27d38fb675e867eb) - int32(__obf_c8d718debc5f6382.__obf_b2da5b1e27e361da)}
	}
	__obf_c24121ecd509691f := new(big.Int)
	__obf_c24121ecd509691f, __obf_c26031e6d546c8ec := __obf_c24121ecd509691f.SetString(string(__obf_c8d718debc5f6382.__obf_c8d718debc5f6382[:__obf_c8d718debc5f6382.__obf_b2da5b1e27e361da]), 10)
	if __obf_c26031e6d546c8ec {
		return Decimal{__obf_830e351d96ac1434: __obf_c24121ecd509691f, __obf_e7725bcd1ee3b1c3: int32(__obf_c8d718debc5f6382.__obf_27d38fb675e867eb) - int32(__obf_c8d718debc5f6382.__obf_b2da5b1e27e361da)}
	}

	return NewFromFloatWithExponent(__obf_3ce5743c80ca37b3, int32(__obf_c8d718debc5f6382.__obf_27d38fb675e867eb)-int32(__obf_c8d718debc5f6382.__obf_b2da5b1e27e361da))
}

// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
// number of fractional digits.
//
// Example:
//
//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
func NewFromFloatWithExponent(__obf_830e351d96ac1434 float64, __obf_e7725bcd1ee3b1c3 int32) Decimal {
	if math.IsNaN(__obf_830e351d96ac1434) || math.IsInf(__obf_830e351d96ac1434, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_830e351d96ac1434))
	}

	__obf_6e1762ca2200a050 := math.Float64bits(__obf_830e351d96ac1434)
	__obf_2a9e5a4d768e50b1 := __obf_6e1762ca2200a050 & (1<<52 - 1)
	__obf_f98b4172037efe19 := int32((__obf_6e1762ca2200a050 >> 52) & (1<<11 - 1))
	__obf_20d020ec018f3926 := __obf_6e1762ca2200a050 >> 63

	if __obf_f98b4172037efe19 == 0 {
		// specials
		if __obf_2a9e5a4d768e50b1 == 0 {
			return Decimal{}
		}
		// subnormal
		__obf_f98b4172037efe19++
	} else {
		// normal
		__obf_2a9e5a4d768e50b1 |= 1 << 52
	}

	__obf_f98b4172037efe19 -= 1023 + 52

	// normalizing base-2 values
	for __obf_2a9e5a4d768e50b1&1 == 0 {
		__obf_2a9e5a4d768e50b1 = __obf_2a9e5a4d768e50b1 >> 1
		__obf_f98b4172037efe19++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_e7725bcd1ee3b1c3 < 0 && __obf_e7725bcd1ee3b1c3 < __obf_f98b4172037efe19 {
		if __obf_f98b4172037efe19 < 0 {
			__obf_e7725bcd1ee3b1c3 = __obf_f98b4172037efe19
		} else {
			__obf_e7725bcd1ee3b1c3 = 0
		}
	}

	// representing 10^M * 2^N as 5^M * 2^(M+N)
	__obf_f98b4172037efe19 -= __obf_e7725bcd1ee3b1c3

	__obf_86f2e858a28847bb := big.NewInt(1)
	__obf_3e32987f5ea394eb := big.NewInt(int64(__obf_2a9e5a4d768e50b1))

	// applying 5^M
	if __obf_e7725bcd1ee3b1c3 > 0 {
		__obf_86f2e858a28847bb = __obf_86f2e858a28847bb.SetInt64(int64(__obf_e7725bcd1ee3b1c3))
		__obf_86f2e858a28847bb = __obf_86f2e858a28847bb.Exp(__obf_0fdf935c5e6da9d3, __obf_86f2e858a28847bb, nil)
	} else if __obf_e7725bcd1ee3b1c3 < 0 {
		__obf_86f2e858a28847bb = __obf_86f2e858a28847bb.SetInt64(-int64(__obf_e7725bcd1ee3b1c3))
		__obf_86f2e858a28847bb = __obf_86f2e858a28847bb.Exp(__obf_0fdf935c5e6da9d3, __obf_86f2e858a28847bb, nil)
		__obf_3e32987f5ea394eb = __obf_3e32987f5ea394eb.Mul(__obf_3e32987f5ea394eb, __obf_86f2e858a28847bb)
		__obf_86f2e858a28847bb = __obf_86f2e858a28847bb.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_f98b4172037efe19 > 0 {
		__obf_3e32987f5ea394eb = __obf_3e32987f5ea394eb.Lsh(__obf_3e32987f5ea394eb, uint(__obf_f98b4172037efe19))
	} else if __obf_f98b4172037efe19 < 0 {
		__obf_86f2e858a28847bb = __obf_86f2e858a28847bb.Lsh(__obf_86f2e858a28847bb, uint(-__obf_f98b4172037efe19))
	}

	// rounding and downscaling
	if __obf_e7725bcd1ee3b1c3 > 0 || __obf_f98b4172037efe19 < 0 {
		__obf_890bee35e8c85b50 := new(big.Int).Rsh(__obf_86f2e858a28847bb, 1)
		__obf_3e32987f5ea394eb = __obf_3e32987f5ea394eb.Add(__obf_3e32987f5ea394eb, __obf_890bee35e8c85b50)
		__obf_3e32987f5ea394eb = __obf_3e32987f5ea394eb.Quo(__obf_3e32987f5ea394eb, __obf_86f2e858a28847bb)
	}

	if __obf_20d020ec018f3926 == 1 {
		__obf_3e32987f5ea394eb = __obf_3e32987f5ea394eb.Neg(__obf_3e32987f5ea394eb)
	}

	return Decimal{
		__obf_830e351d96ac1434: __obf_3e32987f5ea394eb,
		__obf_e7725bcd1ee3b1c3: __obf_e7725bcd1ee3b1c3,
	}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_c8d718debc5f6382 Decimal) Copy() Decimal {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()
	return Decimal{
		__obf_830e351d96ac1434: new(big.Int).Set(__obf_c8d718debc5f6382.__obf_830e351d96ac1434),
		__obf_e7725bcd1ee3b1c3: __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3,
	}
}

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
func (__obf_c8d718debc5f6382 Decimal) __obf_e40699538ea665de(__obf_e7725bcd1ee3b1c3 int32) Decimal {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()

	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 == __obf_e7725bcd1ee3b1c3 {
		return Decimal{
			new(big.Int).Set(__obf_c8d718debc5f6382.__obf_830e351d96ac1434),
			__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3,
		}
	}

	// NOTE(vadim): must convert exps to float64 before - to prevent overflow
	__obf_d561e651075608c1 := math.Abs(float64(__obf_e7725bcd1ee3b1c3) - float64(__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3))
	__obf_830e351d96ac1434 := new(big.Int).Set(__obf_c8d718debc5f6382.__obf_830e351d96ac1434)

	__obf_77f610f5db136ddd := new(big.Int).Exp(__obf_125c1d0eb0d00860, big.NewInt(int64(__obf_d561e651075608c1)), nil)
	if __obf_e7725bcd1ee3b1c3 > __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 {
		__obf_830e351d96ac1434 = __obf_830e351d96ac1434.Quo(__obf_830e351d96ac1434, __obf_77f610f5db136ddd)
	} else if __obf_e7725bcd1ee3b1c3 < __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 {
		__obf_830e351d96ac1434 = __obf_830e351d96ac1434.Mul(__obf_830e351d96ac1434, __obf_77f610f5db136ddd)
	}

	return Decimal{
		__obf_830e351d96ac1434: __obf_830e351d96ac1434,
		__obf_e7725bcd1ee3b1c3: __obf_e7725bcd1ee3b1c3,
	}
}

// Abs returns the absolute value of the decimal.
func (__obf_c8d718debc5f6382 Decimal) Abs() Decimal {
	if !__obf_c8d718debc5f6382.IsNegative() {
		return __obf_c8d718debc5f6382
	}
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()
	__obf_a5bf9b5b656f36ad := new(big.Int).Abs(__obf_c8d718debc5f6382.__obf_830e351d96ac1434)
	return Decimal{
		__obf_830e351d96ac1434: __obf_a5bf9b5b656f36ad,
		__obf_e7725bcd1ee3b1c3: __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3,
	}
}

// Add returns d + d2.
func (__obf_c8d718debc5f6382 Decimal) Add(__obf_4b29119296c7019d Decimal) Decimal {
	__obf_1690aeca4bfa804c, __obf_4a2b3dfc8d0c7c14 := RescalePair(__obf_c8d718debc5f6382, __obf_4b29119296c7019d)

	__obf_3ce8553ac02932b7 := new(big.Int).Add(__obf_1690aeca4bfa804c.__obf_830e351d96ac1434, __obf_4a2b3dfc8d0c7c14.__obf_830e351d96ac1434)
	return Decimal{
		__obf_830e351d96ac1434: __obf_3ce8553ac02932b7,
		__obf_e7725bcd1ee3b1c3: __obf_1690aeca4bfa804c.__obf_e7725bcd1ee3b1c3,
	}
}

// Sub returns d - d2.
func (__obf_c8d718debc5f6382 Decimal) Sub(__obf_4b29119296c7019d Decimal) Decimal {
	__obf_1690aeca4bfa804c, __obf_4a2b3dfc8d0c7c14 := RescalePair(__obf_c8d718debc5f6382, __obf_4b29119296c7019d)

	__obf_3ce8553ac02932b7 := new(big.Int).Sub(__obf_1690aeca4bfa804c.__obf_830e351d96ac1434, __obf_4a2b3dfc8d0c7c14.__obf_830e351d96ac1434)
	return Decimal{
		__obf_830e351d96ac1434: __obf_3ce8553ac02932b7,
		__obf_e7725bcd1ee3b1c3: __obf_1690aeca4bfa804c.__obf_e7725bcd1ee3b1c3,
	}
}

// Neg returns -d.
func (__obf_c8d718debc5f6382 Decimal) Neg() Decimal {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()
	__obf_3ce5743c80ca37b3 := new(big.Int).Neg(__obf_c8d718debc5f6382.__obf_830e351d96ac1434)
	return Decimal{
		__obf_830e351d96ac1434: __obf_3ce5743c80ca37b3,
		__obf_e7725bcd1ee3b1c3: __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3,
	}
}

// Mul returns d * d2.
func (__obf_c8d718debc5f6382 Decimal) Mul(__obf_4b29119296c7019d Decimal) Decimal {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()
	__obf_4b29119296c7019d.__obf_b05475f1c7ee2e48()

	__obf_1b54183868fe131a := int64(__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3) + int64(__obf_4b29119296c7019d.__obf_e7725bcd1ee3b1c3)
	if __obf_1b54183868fe131a > math.MaxInt32 || __obf_1b54183868fe131a < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_1b54183868fe131a))
	}

	__obf_3ce8553ac02932b7 := new(big.Int).Mul(__obf_c8d718debc5f6382.__obf_830e351d96ac1434, __obf_4b29119296c7019d.__obf_830e351d96ac1434)
	return Decimal{
		__obf_830e351d96ac1434: __obf_3ce8553ac02932b7,
		__obf_e7725bcd1ee3b1c3: int32(__obf_1b54183868fe131a),
	}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_c8d718debc5f6382 Decimal) Shift(__obf_a0be95fc360bd94c int32) Decimal {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()
	return Decimal{
		__obf_830e351d96ac1434: new(big.Int).Set(__obf_c8d718debc5f6382.__obf_830e351d96ac1434),
		__obf_e7725bcd1ee3b1c3: __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 + __obf_a0be95fc360bd94c,
	}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_c8d718debc5f6382 Decimal) Div(__obf_4b29119296c7019d Decimal) Decimal {
	return __obf_c8d718debc5f6382.DivRound(__obf_4b29119296c7019d, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_c8d718debc5f6382 Decimal) QuoRem(__obf_4b29119296c7019d Decimal, __obf_a2144abdd719f349 int32) (Decimal, Decimal) {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()
	__obf_4b29119296c7019d.__obf_b05475f1c7ee2e48()
	if __obf_4b29119296c7019d.__obf_830e351d96ac1434.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_03569877ef966d99 := -__obf_a2144abdd719f349
	__obf_250e118ca646cc77 := int64(__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3) - int64(__obf_4b29119296c7019d.__obf_e7725bcd1ee3b1c3) - int64(__obf_03569877ef966d99)
	if __obf_250e118ca646cc77 > math.MaxInt32 || __obf_250e118ca646cc77 < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_ebcc664f54c058d8, __obf_3945df25de247d3b, __obf_9317f0c9da57ba0f big.Int
	var __obf_c5d73bc6c6d2c22f int32
	// d = a 10^ea
	// d2 = b 10^eb
	if __obf_250e118ca646cc77 < 0 {
		__obf_ebcc664f54c058d8 = *__obf_c8d718debc5f6382.__obf_830e351d96ac1434
		__obf_9317f0c9da57ba0f.SetInt64(-__obf_250e118ca646cc77)
		__obf_3945df25de247d3b.Exp(__obf_125c1d0eb0d00860, &__obf_9317f0c9da57ba0f, nil)
		__obf_3945df25de247d3b.Mul(__obf_4b29119296c7019d.__obf_830e351d96ac1434, &__obf_3945df25de247d3b)
		__obf_c5d73bc6c6d2c22f = __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3
		// now aa = a
		//     bb = b 10^(scale + eb - ea)
	} else {
		__obf_9317f0c9da57ba0f.SetInt64(__obf_250e118ca646cc77)
		__obf_ebcc664f54c058d8.Exp(__obf_125c1d0eb0d00860, &__obf_9317f0c9da57ba0f, nil)
		__obf_ebcc664f54c058d8.Mul(__obf_c8d718debc5f6382.__obf_830e351d96ac1434, &__obf_ebcc664f54c058d8)
		__obf_3945df25de247d3b = *__obf_4b29119296c7019d.__obf_830e351d96ac1434
		__obf_c5d73bc6c6d2c22f = __obf_03569877ef966d99 + __obf_4b29119296c7019d.__obf_e7725bcd1ee3b1c3
		// now aa = a ^ (ea - eb - scale)
		//     bb = b
	}
	var __obf_bdffe43eb2ff0705, __obf_791a40daaf73f23a big.Int
	__obf_bdffe43eb2ff0705.QuoRem(&__obf_ebcc664f54c058d8, &__obf_3945df25de247d3b, &__obf_791a40daaf73f23a)
	__obf_33c700b2e9f3fc41 := Decimal{__obf_830e351d96ac1434: &__obf_bdffe43eb2ff0705, __obf_e7725bcd1ee3b1c3: __obf_03569877ef966d99}
	__obf_15854ddd0bdd4927 := Decimal{__obf_830e351d96ac1434: &__obf_791a40daaf73f23a, __obf_e7725bcd1ee3b1c3: __obf_c5d73bc6c6d2c22f}
	return __obf_33c700b2e9f3fc41, __obf_15854ddd0bdd4927
}

// DivRound divides and rounds to a given precision
// i.e. to an integer multiple of 10^(-precision)
//
//	for a positive quotient digit 5 is rounded up, away from 0
//	if the quotient is negative then digit 5 is rounded down, away from 0
//
// Note that precision<0 is allowed as input.
func (__obf_c8d718debc5f6382 Decimal) DivRound(__obf_4b29119296c7019d Decimal, __obf_a2144abdd719f349 int32) Decimal {
	// QuoRem already checks initialization
	__obf_bdffe43eb2ff0705, __obf_791a40daaf73f23a := __obf_c8d718debc5f6382.QuoRem(__obf_4b29119296c7019d, __obf_a2144abdd719f349)
	// the actual rounding decision is based on comparing r*10^precision and d2/2
	// instead compare 2 r 10 ^precision and d2
	var __obf_8632563557db0ce8 big.Int
	__obf_8632563557db0ce8.Abs(__obf_791a40daaf73f23a.__obf_830e351d96ac1434)
	__obf_8632563557db0ce8.Lsh(&__obf_8632563557db0ce8, 1)
	// now rv2 = abs(r.value) * 2
	__obf_16e74080edf2b156 := Decimal{__obf_830e351d96ac1434: &__obf_8632563557db0ce8, __obf_e7725bcd1ee3b1c3: __obf_791a40daaf73f23a.__obf_e7725bcd1ee3b1c3 + __obf_a2144abdd719f349}
	// r2 is now 2 * r * 10 ^ precision
	var __obf_2b375914089c38b5 = __obf_16e74080edf2b156.Cmp(__obf_4b29119296c7019d.Abs())

	if __obf_2b375914089c38b5 < 0 {
		return __obf_bdffe43eb2ff0705
	}

	if __obf_c8d718debc5f6382.__obf_830e351d96ac1434.Sign()*__obf_4b29119296c7019d.__obf_830e351d96ac1434.Sign() < 0 {
		return __obf_bdffe43eb2ff0705.Sub(New(1, -__obf_a2144abdd719f349))
	}

	return __obf_bdffe43eb2ff0705.Add(New(1, -__obf_a2144abdd719f349))
}

// Mod returns d % d2.
func (__obf_c8d718debc5f6382 Decimal) Mod(__obf_4b29119296c7019d Decimal) Decimal {
	_, __obf_791a40daaf73f23a := __obf_c8d718debc5f6382.QuoRem(__obf_4b29119296c7019d, 0)
	return __obf_791a40daaf73f23a
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
func (__obf_c8d718debc5f6382 Decimal) Pow(__obf_4b29119296c7019d Decimal) Decimal {
	__obf_6b3a58839bcd3380 := __obf_c8d718debc5f6382.Sign()
	__obf_1941b5085d2c2cab := __obf_4b29119296c7019d.Sign()

	if __obf_6b3a58839bcd3380 == 0 {
		if __obf_1941b5085d2c2cab == 0 {
			return Decimal{}
		}
		if __obf_1941b5085d2c2cab == 1 {
			return Decimal{__obf_ed0e7e529c925e39, 0}
		}
		if __obf_1941b5085d2c2cab == -1 {
			return Decimal{}
		}
	}

	if __obf_1941b5085d2c2cab == 0 {
		return Decimal{__obf_acf4ecfba54e0451, 0}
	}

	// TODO: optimize extraction of fractional part
	__obf_49b2500e76899510 := Decimal{__obf_acf4ecfba54e0451, 0}
	__obf_6338039aa3687191, __obf_e393487c5544d68f := __obf_4b29119296c7019d.QuoRem(__obf_49b2500e76899510, 0)

	if __obf_6b3a58839bcd3380 == -1 && !__obf_e393487c5544d68f.IsZero() {
		return Decimal{}
	}

	__obf_7c8d793d37c25a52, _ := __obf_c8d718debc5f6382.PowBigInt(__obf_6338039aa3687191.__obf_830e351d96ac1434)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_e393487c5544d68f.__obf_830e351d96ac1434.Sign() == 0 {
		return __obf_7c8d793d37c25a52
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_ba52702546efe7f6 := __obf_c8d718debc5f6382.NumDigits()
	__obf_3fa6c9575bb81227 := __obf_4b29119296c7019d.NumDigits()

	__obf_a2144abdd719f349 := __obf_ba52702546efe7f6

	if __obf_3fa6c9575bb81227 > __obf_a2144abdd719f349 {
		__obf_a2144abdd719f349 += __obf_3fa6c9575bb81227
	}

	__obf_a2144abdd719f349 += 6

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_fe9c6e7cae23f453, __obf_95437c99969367dc := __obf_c8d718debc5f6382.Abs().Ln(-__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 + int32(__obf_a2144abdd719f349))
	if __obf_95437c99969367dc != nil {
		return Decimal{}
	}

	__obf_fe9c6e7cae23f453 = __obf_fe9c6e7cae23f453.Mul(__obf_e393487c5544d68f)

	__obf_fe9c6e7cae23f453, __obf_95437c99969367dc = __obf_fe9c6e7cae23f453.ExpTaylor(-__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 + int32(__obf_a2144abdd719f349))
	if __obf_95437c99969367dc != nil {
		return Decimal{}
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_53c898c8ea3bfead := __obf_7c8d793d37c25a52.Mul(__obf_fe9c6e7cae23f453)

	return __obf_53c898c8ea3bfead
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
func (__obf_c8d718debc5f6382 Decimal) PowWithPrecision(__obf_4b29119296c7019d Decimal, __obf_a2144abdd719f349 int32) (Decimal, error) {
	__obf_6b3a58839bcd3380 := __obf_c8d718debc5f6382.Sign()
	__obf_1941b5085d2c2cab := __obf_4b29119296c7019d.Sign()

	if __obf_6b3a58839bcd3380 == 0 {
		if __obf_1941b5085d2c2cab == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_1941b5085d2c2cab == 1 {
			return Decimal{__obf_ed0e7e529c925e39, 0}, nil
		}
		if __obf_1941b5085d2c2cab == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_1941b5085d2c2cab == 0 {
		return Decimal{__obf_acf4ecfba54e0451, 0}, nil
	}

	// TODO: optimize extraction of fractional part
	__obf_49b2500e76899510 := Decimal{__obf_acf4ecfba54e0451, 0}
	__obf_6338039aa3687191, __obf_e393487c5544d68f := __obf_4b29119296c7019d.QuoRem(__obf_49b2500e76899510, 0)

	if __obf_6b3a58839bcd3380 == -1 && !__obf_e393487c5544d68f.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}

	__obf_7c8d793d37c25a52, _ := __obf_c8d718debc5f6382.__obf_c2328fefbd9b5177(__obf_6338039aa3687191.__obf_830e351d96ac1434, __obf_a2144abdd719f349)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_e393487c5544d68f.__obf_830e351d96ac1434.Sign() == 0 {
		return __obf_7c8d793d37c25a52, nil
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_ba52702546efe7f6 := __obf_c8d718debc5f6382.NumDigits()
	__obf_3fa6c9575bb81227 := __obf_4b29119296c7019d.NumDigits()

	if int32(__obf_ba52702546efe7f6) > __obf_a2144abdd719f349 {
		__obf_a2144abdd719f349 = int32(__obf_ba52702546efe7f6)
	}
	if int32(__obf_3fa6c9575bb81227) > __obf_a2144abdd719f349 {
		__obf_a2144abdd719f349 += int32(__obf_3fa6c9575bb81227)
	}
	// increase precision by 10 to compensate for errors in further calculations
	__obf_a2144abdd719f349 += 10

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_fe9c6e7cae23f453, __obf_95437c99969367dc := __obf_c8d718debc5f6382.Abs().Ln(__obf_a2144abdd719f349)
	if __obf_95437c99969367dc != nil {
		return Decimal{}, __obf_95437c99969367dc
	}

	__obf_fe9c6e7cae23f453 = __obf_fe9c6e7cae23f453.Mul(__obf_e393487c5544d68f)

	__obf_fe9c6e7cae23f453, __obf_95437c99969367dc = __obf_fe9c6e7cae23f453.ExpTaylor(__obf_a2144abdd719f349)
	if __obf_95437c99969367dc != nil {
		return Decimal{}, __obf_95437c99969367dc
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_53c898c8ea3bfead := __obf_7c8d793d37c25a52.Mul(__obf_fe9c6e7cae23f453)

	return __obf_53c898c8ea3bfead, nil
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
func (__obf_c8d718debc5f6382 Decimal) PowInt32(__obf_e7725bcd1ee3b1c3 int32) (Decimal, error) {
	if __obf_c8d718debc5f6382.IsZero() && __obf_e7725bcd1ee3b1c3 == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_7fcad2fe2fef8e4a := __obf_e7725bcd1ee3b1c3 < 0
	__obf_e7725bcd1ee3b1c3 = __obf_8b38a12ace799621(__obf_e7725bcd1ee3b1c3)

	__obf_648608c066f38982, __obf_527bb05cffd104d5 := __obf_c8d718debc5f6382, New(1, 0)

	for __obf_e7725bcd1ee3b1c3 > 0 {
		if __obf_e7725bcd1ee3b1c3%2 == 1 {
			__obf_527bb05cffd104d5 = __obf_527bb05cffd104d5.Mul(__obf_648608c066f38982)
		}
		__obf_e7725bcd1ee3b1c3 /= 2

		if __obf_e7725bcd1ee3b1c3 > 0 {
			__obf_648608c066f38982 = __obf_648608c066f38982.Mul(__obf_648608c066f38982)
		}
	}

	if __obf_7fcad2fe2fef8e4a {
		return New(1, 0).DivRound(__obf_527bb05cffd104d5, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_527bb05cffd104d5, nil
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
func (__obf_c8d718debc5f6382 Decimal) PowBigInt(__obf_e7725bcd1ee3b1c3 *big.Int) (Decimal, error) {
	return __obf_c8d718debc5f6382.__obf_c2328fefbd9b5177(__obf_e7725bcd1ee3b1c3, int32(PowPrecisionNegativeExponent))
}

func (__obf_c8d718debc5f6382 Decimal) __obf_c2328fefbd9b5177(__obf_e7725bcd1ee3b1c3 *big.Int, __obf_a2144abdd719f349 int32) (Decimal, error) {
	if __obf_c8d718debc5f6382.IsZero() && __obf_e7725bcd1ee3b1c3.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_1731eef0664e4efd := new(big.Int).Set(__obf_e7725bcd1ee3b1c3)
	__obf_7fcad2fe2fef8e4a := __obf_e7725bcd1ee3b1c3.Sign() < 0

	if __obf_7fcad2fe2fef8e4a {
		__obf_1731eef0664e4efd.Abs(__obf_1731eef0664e4efd)
	}

	__obf_648608c066f38982, __obf_527bb05cffd104d5 := __obf_c8d718debc5f6382, New(1, 0)

	for __obf_1731eef0664e4efd.Sign() > 0 {
		if __obf_1731eef0664e4efd.Bit(0) == 1 {
			__obf_527bb05cffd104d5 = __obf_527bb05cffd104d5.Mul(__obf_648608c066f38982)
		}
		__obf_1731eef0664e4efd.Rsh(__obf_1731eef0664e4efd, 1)

		if __obf_1731eef0664e4efd.Sign() > 0 {
			__obf_648608c066f38982 = __obf_648608c066f38982.Mul(__obf_648608c066f38982)
		}
	}

	if __obf_7fcad2fe2fef8e4a {
		return New(1, 0).DivRound(__obf_527bb05cffd104d5, __obf_a2144abdd719f349), nil
	}

	return __obf_527bb05cffd104d5, nil
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
func (__obf_c8d718debc5f6382 Decimal) ExpHullAbrham(__obf_062edd20b4c2d151 uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_c8d718debc5f6382.IsZero() {
		return Decimal{__obf_acf4ecfba54e0451, 0}, nil
	}

	__obf_bf5b76aa2bf93256 := __obf_062edd20b4c2d151

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_957bb88d2d813f36 := __obf_c8d718debc5f6382.Abs().InexactFloat64()
	if __obf_685090547b10ea9c := __obf_957bb88d2d813f36 / 23; __obf_685090547b10ea9c > float64(__obf_bf5b76aa2bf93256) && __obf_685090547b10ea9c < float64(ExpMaxIterations) {
		__obf_bf5b76aa2bf93256 = uint32(math.Ceil(__obf_685090547b10ea9c))
	}

	// fail if abs(d) beyond an over/underflow threshold
	__obf_9e798128655c10e8 := New(23*int64(__obf_bf5b76aa2bf93256), 0)
	if __obf_c8d718debc5f6382.Abs().Cmp(__obf_9e798128655c10e8) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}

	// Return 1 if abs(d) small enough; this also avoids later over/underflow
	__obf_fdb8ed6dcd20268b := New(9, -int32(__obf_bf5b76aa2bf93256)-1)
	if __obf_c8d718debc5f6382.Abs().Cmp(__obf_fdb8ed6dcd20268b) <= 0 {
		return Decimal{__obf_acf4ecfba54e0451, __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3}, nil
	}

	// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
	__obf_6964cfeab73d0c1f := __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 + int32(__obf_c8d718debc5f6382.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_6964cfeab73d0c1f < 0 {
		__obf_6964cfeab73d0c1f = 0
	}

	__obf_88a1a18216771a9f := New(1, __obf_6964cfeab73d0c1f)                                                                                                                   // reduction factor
	__obf_791a40daaf73f23a := Decimal{new(big.Int).Set(__obf_c8d718debc5f6382.__obf_830e351d96ac1434), __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 - __obf_6964cfeab73d0c1f} // reduced argument
	__obf_9d14662cf0386368 := int32(__obf_bf5b76aa2bf93256) + __obf_6964cfeab73d0c1f + 2                                                                                       // precision for calculating the sum

	// Determine n, the number of therms for calculating sum
	// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
	// for solving appropriate equation, along with directed
	// roundings and simple rational bound for log10(p/abs(r))
	__obf_939cf35e64006d24 := __obf_791a40daaf73f23a.Abs().InexactFloat64()
	__obf_271a6b6fdbaf7fa5 := float64(__obf_9d14662cf0386368)
	__obf_f8a32b41a4387d03 := math.Ceil((1.453*__obf_271a6b6fdbaf7fa5 - 1.182) / math.Log10(__obf_271a6b6fdbaf7fa5/__obf_939cf35e64006d24))
	if __obf_f8a32b41a4387d03 > float64(ExpMaxIterations) || math.IsNaN(__obf_f8a32b41a4387d03) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_648608c066f38982 := int64(__obf_f8a32b41a4387d03)

	__obf_247932334c18f488 := New(0, 0)
	__obf_16b88190ca2a680e := New(1, 0)
	__obf_49b2500e76899510 := New(1, 0)
	for __obf_4aeff04bb586b943 := __obf_648608c066f38982 - 1; __obf_4aeff04bb586b943 > 0; __obf_4aeff04bb586b943-- {
		__obf_247932334c18f488.__obf_830e351d96ac1434.SetInt64(__obf_4aeff04bb586b943)
		__obf_16b88190ca2a680e = __obf_16b88190ca2a680e.Mul(__obf_791a40daaf73f23a.DivRound(__obf_247932334c18f488, __obf_9d14662cf0386368))
		__obf_16b88190ca2a680e = __obf_16b88190ca2a680e.Add(__obf_49b2500e76899510)
	}

	__obf_4040a1baaf192f28 := __obf_88a1a18216771a9f.IntPart()
	__obf_53c898c8ea3bfead := New(1, 0)
	for __obf_4aeff04bb586b943 := __obf_4040a1baaf192f28; __obf_4aeff04bb586b943 > 0; __obf_4aeff04bb586b943-- {
		__obf_53c898c8ea3bfead = __obf_53c898c8ea3bfead.Mul(__obf_16b88190ca2a680e)
	}

	__obf_6a1eaac72a6380c2 := int32(__obf_53c898c8ea3bfead.NumDigits())

	var __obf_26706e607619ea23 int32
	if __obf_6a1eaac72a6380c2 > __obf_8b38a12ace799621(__obf_53c898c8ea3bfead.__obf_e7725bcd1ee3b1c3) {
		__obf_26706e607619ea23 = int32(__obf_bf5b76aa2bf93256) - __obf_6a1eaac72a6380c2 - __obf_53c898c8ea3bfead.__obf_e7725bcd1ee3b1c3
	} else {
		__obf_26706e607619ea23 = int32(__obf_bf5b76aa2bf93256)
	}

	__obf_53c898c8ea3bfead = __obf_53c898c8ea3bfead.Round(__obf_26706e607619ea23)

	return __obf_53c898c8ea3bfead, nil
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
func (__obf_c8d718debc5f6382 Decimal) ExpTaylor(__obf_a2144abdd719f349 int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_c8d718debc5f6382.IsZero() {
		return Decimal{__obf_acf4ecfba54e0451, 0}.Round(__obf_a2144abdd719f349), nil
	}

	var __obf_25b977f06662c1d4 Decimal
	var __obf_1015a45bd85b806e int32
	if __obf_a2144abdd719f349 < 0 {
		__obf_25b977f06662c1d4 = New(1, -1)
		__obf_1015a45bd85b806e = 8
	} else {
		__obf_25b977f06662c1d4 = New(1, -__obf_a2144abdd719f349-1)
		__obf_1015a45bd85b806e = __obf_a2144abdd719f349 + 1
	}

	__obf_65eb522d705a325d := __obf_c8d718debc5f6382.Abs()
	__obf_bb547ae420243cd3 := __obf_c8d718debc5f6382.Abs()
	__obf_6569dbf4e50fb8b5 := New(1, 0)

	__obf_527bb05cffd104d5 := New(1, 0)

	for __obf_4aeff04bb586b943 := int64(1); ; {
		__obf_b34df5ca7a64efd3 := __obf_bb547ae420243cd3.DivRound(__obf_6569dbf4e50fb8b5, __obf_1015a45bd85b806e)
		__obf_527bb05cffd104d5 = __obf_527bb05cffd104d5.Add(__obf_b34df5ca7a64efd3)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_b34df5ca7a64efd3.Cmp(__obf_25b977f06662c1d4) < 0 {
			break
		}

		__obf_bb547ae420243cd3 = __obf_bb547ae420243cd3.Mul(__obf_65eb522d705a325d)

		__obf_4aeff04bb586b943++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_d0eb5ea33f7ee79a) >= int(__obf_4aeff04bb586b943) && !__obf_d0eb5ea33f7ee79a[__obf_4aeff04bb586b943-1].IsZero() {
			__obf_6569dbf4e50fb8b5 = __obf_d0eb5ea33f7ee79a[__obf_4aeff04bb586b943-1]
		} else {
			// To avoid any race conditions, firstly the zero value is appended to a slice to create
			// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
			// factorial using the index notation.
			__obf_6569dbf4e50fb8b5 = __obf_d0eb5ea33f7ee79a[__obf_4aeff04bb586b943-2].Mul(New(__obf_4aeff04bb586b943, 0))
			__obf_d0eb5ea33f7ee79a = append(__obf_d0eb5ea33f7ee79a, Zero)
			__obf_d0eb5ea33f7ee79a[__obf_4aeff04bb586b943-1] = __obf_6569dbf4e50fb8b5
		}
	}

	if __obf_c8d718debc5f6382.Sign() < 0 {
		__obf_527bb05cffd104d5 = New(1, 0).DivRound(__obf_527bb05cffd104d5, __obf_a2144abdd719f349+1)
	}

	__obf_527bb05cffd104d5 = __obf_527bb05cffd104d5.Round(__obf_a2144abdd719f349)
	return __obf_527bb05cffd104d5, nil
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
func (__obf_c8d718debc5f6382 Decimal) Ln(__obf_a2144abdd719f349 int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_c8d718debc5f6382.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_c8d718debc5f6382.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}

	__obf_cc399b1b2a4d04b3 := __obf_a2144abdd719f349 + 2
	__obf_ef9b2770cf8a72e8 := __obf_c8d718debc5f6382.Copy()

	var __obf_dbe8cbfd29ae1e55, __obf_ccf0b5b622aa3521, __obf_32590e343ca6e73f, __obf_11cf7704cd490b6c, __obf_229813d2ad96c3f2 Decimal
	__obf_dbe8cbfd29ae1e55 = __obf_ef9b2770cf8a72e8.Sub(Decimal{__obf_acf4ecfba54e0451, 0})
	__obf_ccf0b5b622aa3521 = Decimal{__obf_acf4ecfba54e0451, -1}

	// for decimal in range [0.9, 1.1] where ln(d) is close to 0
	__obf_209f5728159a9dad := false

	if __obf_dbe8cbfd29ae1e55.Abs().Cmp(__obf_ccf0b5b622aa3521) <= 0 {
		__obf_209f5728159a9dad = true
	} else {
		// reduce input decimal to range [0.1, 1)
		__obf_8ed261b9656feeec := int32(__obf_ef9b2770cf8a72e8.NumDigits()) + __obf_ef9b2770cf8a72e8.__obf_e7725bcd1ee3b1c3
		__obf_ef9b2770cf8a72e8.__obf_e7725bcd1ee3b1c3 -= __obf_8ed261b9656feeec

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_0064d0d25ab8c01a := __obf_0064d0d25ab8c01a.__obf_29ea37bb660d8abb(__obf_cc399b1b2a4d04b3)
		__obf_229813d2ad96c3f2 = NewFromInt32(__obf_8ed261b9656feeec)
		__obf_229813d2ad96c3f2 = __obf_229813d2ad96c3f2.Mul(__obf_0064d0d25ab8c01a)

		__obf_dbe8cbfd29ae1e55 = __obf_ef9b2770cf8a72e8.Sub(Decimal{__obf_acf4ecfba54e0451, 0})

		if __obf_dbe8cbfd29ae1e55.Abs().Cmp(__obf_ccf0b5b622aa3521) <= 0 {
			__obf_209f5728159a9dad = true
		} else {
			// initial estimate using floats
			__obf_da03c0f73663e9a1 := __obf_ef9b2770cf8a72e8.InexactFloat64()
			__obf_dbe8cbfd29ae1e55 = NewFromFloat(math.Log(__obf_da03c0f73663e9a1))
		}
	}

	__obf_25b977f06662c1d4 := Decimal{__obf_acf4ecfba54e0451, -__obf_cc399b1b2a4d04b3}

	if __obf_209f5728159a9dad {
		// Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
		// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
		// until the difference between current and next term is smaller than epsilon.
		// Coverage quite fast for decimals close to 1.0

		// z + 2
		__obf_32590e343ca6e73f = __obf_dbe8cbfd29ae1e55.Add(Decimal{__obf_7f85275c7a28a30a, 0})
		// z / (z + 2)
		__obf_ccf0b5b622aa3521 = __obf_dbe8cbfd29ae1e55.DivRound(__obf_32590e343ca6e73f, __obf_cc399b1b2a4d04b3)
		// 2 * (z / (z + 2))
		__obf_dbe8cbfd29ae1e55 = __obf_ccf0b5b622aa3521.Add(__obf_ccf0b5b622aa3521)
		__obf_32590e343ca6e73f = __obf_dbe8cbfd29ae1e55.Copy()

		for __obf_648608c066f38982 := 1; ; __obf_648608c066f38982++ {
			// 2 * (z / (z+2))^(2n+1)
			__obf_32590e343ca6e73f = __obf_32590e343ca6e73f.Mul(__obf_ccf0b5b622aa3521).Mul(__obf_ccf0b5b622aa3521)

			// 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
			__obf_11cf7704cd490b6c = NewFromInt(int64(2*__obf_648608c066f38982 + 1))
			__obf_11cf7704cd490b6c = __obf_32590e343ca6e73f.DivRound(__obf_11cf7704cd490b6c, __obf_cc399b1b2a4d04b3)

			// comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			__obf_dbe8cbfd29ae1e55 = __obf_dbe8cbfd29ae1e55.Add(__obf_11cf7704cd490b6c)

			if __obf_11cf7704cd490b6c.Abs().Cmp(__obf_25b977f06662c1d4) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_0823e297b30db025 Decimal
		__obf_c543f28c56744ffc := __obf_cc399b1b2a4d04b3*2 + 10

		for __obf_4aeff04bb586b943 := int32(0); __obf_4aeff04bb586b943 < __obf_c543f28c56744ffc; __obf_4aeff04bb586b943++ {
			// exp(a_n)
			__obf_ccf0b5b622aa3521, _ = __obf_dbe8cbfd29ae1e55.ExpTaylor(__obf_cc399b1b2a4d04b3)
			// exp(a_n) - z
			__obf_32590e343ca6e73f = __obf_ccf0b5b622aa3521.Sub(__obf_ef9b2770cf8a72e8)
			// 2 * (exp(a_n) - z)
			__obf_32590e343ca6e73f = __obf_32590e343ca6e73f.Add(__obf_32590e343ca6e73f)
			// exp(a_n) + z
			__obf_11cf7704cd490b6c = __obf_ccf0b5b622aa3521.Add(__obf_ef9b2770cf8a72e8)
			// 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_ccf0b5b622aa3521 = __obf_32590e343ca6e73f.DivRound(__obf_11cf7704cd490b6c, __obf_cc399b1b2a4d04b3)
			// comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_dbe8cbfd29ae1e55 = __obf_dbe8cbfd29ae1e55.Sub(__obf_ccf0b5b622aa3521)

			if __obf_0823e297b30db025.Add(__obf_ccf0b5b622aa3521).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_ccf0b5b622aa3521.Abs().Cmp(__obf_25b977f06662c1d4) <= 0 {
				break
			}

			__obf_0823e297b30db025 = __obf_ccf0b5b622aa3521
		}
	}

	__obf_dbe8cbfd29ae1e55 = __obf_dbe8cbfd29ae1e55.Add(__obf_229813d2ad96c3f2)

	return __obf_dbe8cbfd29ae1e55.Round(__obf_a2144abdd719f349), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_c8d718debc5f6382 Decimal) NumDigits() int {
	if __obf_c8d718debc5f6382.__obf_830e351d96ac1434 == nil {
		return 1
	}

	if __obf_c8d718debc5f6382.__obf_830e351d96ac1434.IsInt64() {
		__obf_6f1a29efac56765c := __obf_c8d718debc5f6382.__obf_830e351d96ac1434.Int64()
		// restrict fast path to integers with exact conversion to float64
		if __obf_6f1a29efac56765c <= (1<<53) && __obf_6f1a29efac56765c >= -(1<<53) {
			if __obf_6f1a29efac56765c == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_6f1a29efac56765c)))) + 1
		}
	}

	__obf_bccc68e48bce354b := int(float64(__obf_c8d718debc5f6382.__obf_830e351d96ac1434.BitLen()) / math.Log2(10))

	// estimatedNumDigits (lg10) may be off by 1, need to verify
	__obf_95972db099a0068a := big.NewInt(int64(__obf_bccc68e48bce354b))
	__obf_33ffabf59483a599 := __obf_95972db099a0068a.Exp(__obf_125c1d0eb0d00860, __obf_95972db099a0068a, nil)

	if __obf_c8d718debc5f6382.__obf_830e351d96ac1434.CmpAbs(__obf_33ffabf59483a599) >= 0 {
		return __obf_bccc68e48bce354b + 1
	}

	return __obf_bccc68e48bce354b
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_c8d718debc5f6382 Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_791a40daaf73f23a big.Int
	__obf_bdffe43eb2ff0705 := new(big.Int).Set(__obf_c8d718debc5f6382.__obf_830e351d96ac1434)
	for __obf_ef9b2770cf8a72e8 := __obf_8b38a12ace799621(__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3); __obf_ef9b2770cf8a72e8 > 0; __obf_ef9b2770cf8a72e8-- {
		__obf_bdffe43eb2ff0705.QuoRem(__obf_bdffe43eb2ff0705, __obf_125c1d0eb0d00860, &__obf_791a40daaf73f23a)
		if __obf_791a40daaf73f23a.Cmp(__obf_ed0e7e529c925e39) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_8b38a12ace799621(__obf_648608c066f38982 int32) int32 {
	if __obf_648608c066f38982 < 0 {
		return -__obf_648608c066f38982
	}
	return __obf_648608c066f38982
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_c8d718debc5f6382 Decimal) Cmp(__obf_4b29119296c7019d Decimal) int {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()
	__obf_4b29119296c7019d.__obf_b05475f1c7ee2e48()

	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 == __obf_4b29119296c7019d.__obf_e7725bcd1ee3b1c3 {
		return __obf_c8d718debc5f6382.__obf_830e351d96ac1434.Cmp(__obf_4b29119296c7019d.__obf_830e351d96ac1434)
	}

	__obf_1690aeca4bfa804c, __obf_4a2b3dfc8d0c7c14 := RescalePair(__obf_c8d718debc5f6382, __obf_4b29119296c7019d)

	return __obf_1690aeca4bfa804c.__obf_830e351d96ac1434.Cmp(__obf_4a2b3dfc8d0c7c14.__obf_830e351d96ac1434)
}

// Compare compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_c8d718debc5f6382 Decimal) Compare(__obf_4b29119296c7019d Decimal) int {
	return __obf_c8d718debc5f6382.Cmp(__obf_4b29119296c7019d)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_c8d718debc5f6382 Decimal) Equal(__obf_4b29119296c7019d Decimal) bool {
	return __obf_c8d718debc5f6382.Cmp(__obf_4b29119296c7019d) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_c8d718debc5f6382 Decimal) Equals(__obf_4b29119296c7019d Decimal) bool {
	return __obf_c8d718debc5f6382.Equal(__obf_4b29119296c7019d)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_c8d718debc5f6382 Decimal) GreaterThan(__obf_4b29119296c7019d Decimal) bool {
	return __obf_c8d718debc5f6382.Cmp(__obf_4b29119296c7019d) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_c8d718debc5f6382 Decimal) GreaterThanOrEqual(__obf_4b29119296c7019d Decimal) bool {
	__obf_c8d0708cb0496ff6 := __obf_c8d718debc5f6382.Cmp(__obf_4b29119296c7019d)
	return __obf_c8d0708cb0496ff6 == 1 || __obf_c8d0708cb0496ff6 == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_c8d718debc5f6382 Decimal) LessThan(__obf_4b29119296c7019d Decimal) bool {
	return __obf_c8d718debc5f6382.Cmp(__obf_4b29119296c7019d) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_c8d718debc5f6382 Decimal) LessThanOrEqual(__obf_4b29119296c7019d Decimal) bool {
	__obf_c8d0708cb0496ff6 := __obf_c8d718debc5f6382.Cmp(__obf_4b29119296c7019d)
	return __obf_c8d0708cb0496ff6 == -1 || __obf_c8d0708cb0496ff6 == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_c8d718debc5f6382 Decimal) Sign() int {
	if __obf_c8d718debc5f6382.__obf_830e351d96ac1434 == nil {
		return 0
	}
	return __obf_c8d718debc5f6382.__obf_830e351d96ac1434.Sign()
}

// IsPositive return
//
//	true if d > 0
//	false if d == 0
//	false if d < 0
func (__obf_c8d718debc5f6382 Decimal) IsPositive() bool {
	return __obf_c8d718debc5f6382.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_c8d718debc5f6382 Decimal) IsNegative() bool {
	return __obf_c8d718debc5f6382.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_c8d718debc5f6382 Decimal) IsZero() bool {
	return __obf_c8d718debc5f6382.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_c8d718debc5f6382 Decimal) Exponent() int32 {
	return __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3
}

// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
func (__obf_c8d718debc5f6382 Decimal) Coefficient() *big.Int {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_c8d718debc5f6382.__obf_830e351d96ac1434)
}

// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
// If coefficient cannot be represented in an int64, the result will be undefined.
func (__obf_c8d718debc5f6382 Decimal) CoefficientInt64() int64 {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()
	return __obf_c8d718debc5f6382.__obf_830e351d96ac1434.Int64()
}

// IntPart returns the integer component of the decimal.
func (__obf_c8d718debc5f6382 Decimal) IntPart() int64 {
	__obf_f32c24bd3578a970 := __obf_c8d718debc5f6382.__obf_e40699538ea665de(0)
	return __obf_f32c24bd3578a970.__obf_830e351d96ac1434.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_c8d718debc5f6382 Decimal) BigInt() *big.Int {
	__obf_f32c24bd3578a970 := __obf_c8d718debc5f6382.__obf_e40699538ea665de(0)
	return __obf_f32c24bd3578a970.__obf_830e351d96ac1434
}

// BigFloat returns decimal as BigFloat.
// Be aware that casting decimal to BigFloat might cause a loss of precision.
func (__obf_c8d718debc5f6382 Decimal) BigFloat() *big.Float {
	__obf_957bb88d2d813f36 := &big.Float{}
	__obf_957bb88d2d813f36.SetString(__obf_c8d718debc5f6382.String())
	return __obf_957bb88d2d813f36
}

// Rat returns a rational number representation of the decimal.
func (__obf_c8d718debc5f6382 Decimal) Rat() *big.Rat {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()
	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 <= 0 {
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_94b5aa61eb5be998 := new(big.Int).Exp(__obf_125c1d0eb0d00860, big.NewInt(-int64(__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3)), nil)
		return new(big.Rat).SetFrac(__obf_c8d718debc5f6382.__obf_830e351d96ac1434, __obf_94b5aa61eb5be998)
	}

	__obf_a8aabc0a3f68b1da := new(big.Int).Exp(__obf_125c1d0eb0d00860, big.NewInt(int64(__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3)), nil)
	__obf_7542327471b07e63 := new(big.Int).Mul(__obf_c8d718debc5f6382.__obf_830e351d96ac1434, __obf_a8aabc0a3f68b1da)
	return new(big.Rat).SetFrac(__obf_7542327471b07e63, __obf_acf4ecfba54e0451)
}

// Float64 returns the nearest float64 value for d and a bool indicating
// whether f represents d exactly.
// For more details, see the documentation for big.Rat.Float64
func (__obf_c8d718debc5f6382 Decimal) Float64() (__obf_957bb88d2d813f36 float64, __obf_498ddb57b3e80e48 bool) {
	return __obf_c8d718debc5f6382.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_c8d718debc5f6382 Decimal) InexactFloat64() float64 {
	__obf_957bb88d2d813f36, _ := __obf_c8d718debc5f6382.Float64()
	return __obf_957bb88d2d813f36
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
func (__obf_c8d718debc5f6382 Decimal) String() string {
	return __obf_c8d718debc5f6382.string(true)
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
func (__obf_c8d718debc5f6382 Decimal) StringFixed(__obf_60027c47d471a94d int32) string {
	__obf_421682ca894f51c9 := __obf_c8d718debc5f6382.Round(__obf_60027c47d471a94d)
	return __obf_421682ca894f51c9.string(false)
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
func (__obf_c8d718debc5f6382 Decimal) StringFixedBank(__obf_60027c47d471a94d int32) string {
	__obf_421682ca894f51c9 := __obf_c8d718debc5f6382.RoundBank(__obf_60027c47d471a94d)
	return __obf_421682ca894f51c9.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_c8d718debc5f6382 Decimal) StringFixedCash(__obf_2a03d45ad2a75a03 uint8) string {
	__obf_421682ca894f51c9 := __obf_c8d718debc5f6382.RoundCash(__obf_2a03d45ad2a75a03)
	return __obf_421682ca894f51c9.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_c8d718debc5f6382 Decimal) Round(__obf_60027c47d471a94d int32) Decimal {
	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 == -__obf_60027c47d471a94d {
		return __obf_c8d718debc5f6382
	}
	// truncate to places + 1
	__obf_8696cda2270d5f4b := __obf_c8d718debc5f6382.__obf_e40699538ea665de(-__obf_60027c47d471a94d - 1)

	// add sign(d) * 0.5
	if __obf_8696cda2270d5f4b.__obf_830e351d96ac1434.Sign() < 0 {
		__obf_8696cda2270d5f4b.__obf_830e351d96ac1434.Sub(__obf_8696cda2270d5f4b.__obf_830e351d96ac1434, __obf_0fdf935c5e6da9d3)
	} else {
		__obf_8696cda2270d5f4b.__obf_830e351d96ac1434.Add(__obf_8696cda2270d5f4b.__obf_830e351d96ac1434, __obf_0fdf935c5e6da9d3)
	}

	// floor for positive numbers, ceil for negative numbers
	_, __obf_4931dd9c612fb4d2 := __obf_8696cda2270d5f4b.__obf_830e351d96ac1434.DivMod(__obf_8696cda2270d5f4b.__obf_830e351d96ac1434, __obf_125c1d0eb0d00860, new(big.Int))
	__obf_8696cda2270d5f4b.__obf_e7725bcd1ee3b1c3++
	if __obf_8696cda2270d5f4b.__obf_830e351d96ac1434.Sign() < 0 && __obf_4931dd9c612fb4d2.Cmp(__obf_ed0e7e529c925e39) != 0 {
		__obf_8696cda2270d5f4b.__obf_830e351d96ac1434.Add(__obf_8696cda2270d5f4b.__obf_830e351d96ac1434, __obf_acf4ecfba54e0451)
	}

	return __obf_8696cda2270d5f4b
}

// RoundCeil rounds the decimal towards +infinity.
//
// Example:
//
//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
func (__obf_c8d718debc5f6382 Decimal) RoundCeil(__obf_60027c47d471a94d int32) Decimal {
	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 >= -__obf_60027c47d471a94d {
		return __obf_c8d718debc5f6382
	}

	__obf_2e37f945e84efeb2 := __obf_c8d718debc5f6382.__obf_e40699538ea665de(-__obf_60027c47d471a94d)
	if __obf_c8d718debc5f6382.Equal(__obf_2e37f945e84efeb2) {
		return __obf_c8d718debc5f6382
	}

	if __obf_c8d718debc5f6382.__obf_830e351d96ac1434.Sign() > 0 {
		__obf_2e37f945e84efeb2.__obf_830e351d96ac1434.Add(__obf_2e37f945e84efeb2.__obf_830e351d96ac1434, __obf_acf4ecfba54e0451)
	}

	return __obf_2e37f945e84efeb2
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_c8d718debc5f6382 Decimal) RoundFloor(__obf_60027c47d471a94d int32) Decimal {
	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 >= -__obf_60027c47d471a94d {
		return __obf_c8d718debc5f6382
	}

	__obf_2e37f945e84efeb2 := __obf_c8d718debc5f6382.__obf_e40699538ea665de(-__obf_60027c47d471a94d)
	if __obf_c8d718debc5f6382.Equal(__obf_2e37f945e84efeb2) {
		return __obf_c8d718debc5f6382
	}

	if __obf_c8d718debc5f6382.__obf_830e351d96ac1434.Sign() < 0 {
		__obf_2e37f945e84efeb2.__obf_830e351d96ac1434.Sub(__obf_2e37f945e84efeb2.__obf_830e351d96ac1434, __obf_acf4ecfba54e0451)
	}

	return __obf_2e37f945e84efeb2
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_c8d718debc5f6382 Decimal) RoundUp(__obf_60027c47d471a94d int32) Decimal {
	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 >= -__obf_60027c47d471a94d {
		return __obf_c8d718debc5f6382
	}

	__obf_2e37f945e84efeb2 := __obf_c8d718debc5f6382.__obf_e40699538ea665de(-__obf_60027c47d471a94d)
	if __obf_c8d718debc5f6382.Equal(__obf_2e37f945e84efeb2) {
		return __obf_c8d718debc5f6382
	}

	if __obf_c8d718debc5f6382.__obf_830e351d96ac1434.Sign() > 0 {
		__obf_2e37f945e84efeb2.__obf_830e351d96ac1434.Add(__obf_2e37f945e84efeb2.__obf_830e351d96ac1434, __obf_acf4ecfba54e0451)
	} else if __obf_c8d718debc5f6382.__obf_830e351d96ac1434.Sign() < 0 {
		__obf_2e37f945e84efeb2.__obf_830e351d96ac1434.Sub(__obf_2e37f945e84efeb2.__obf_830e351d96ac1434, __obf_acf4ecfba54e0451)
	}

	return __obf_2e37f945e84efeb2
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_c8d718debc5f6382 Decimal) RoundDown(__obf_60027c47d471a94d int32) Decimal {
	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 >= -__obf_60027c47d471a94d {
		return __obf_c8d718debc5f6382
	}

	__obf_2e37f945e84efeb2 := __obf_c8d718debc5f6382.__obf_e40699538ea665de(-__obf_60027c47d471a94d)
	if __obf_c8d718debc5f6382.Equal(__obf_2e37f945e84efeb2) {
		return __obf_c8d718debc5f6382
	}
	return __obf_2e37f945e84efeb2
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
func (__obf_c8d718debc5f6382 Decimal) RoundBank(__obf_60027c47d471a94d int32) Decimal {

	__obf_68da4d6990322abc := __obf_c8d718debc5f6382.Round(__obf_60027c47d471a94d)
	__obf_75d91a4485b563af := __obf_c8d718debc5f6382.Sub(__obf_68da4d6990322abc).Abs()

	__obf_fd169dd113b7681c := New(5, -__obf_60027c47d471a94d-1)
	if __obf_75d91a4485b563af.Cmp(__obf_fd169dd113b7681c) == 0 && __obf_68da4d6990322abc.__obf_830e351d96ac1434.Bit(0) != 0 {
		if __obf_68da4d6990322abc.__obf_830e351d96ac1434.Sign() < 0 {
			__obf_68da4d6990322abc.__obf_830e351d96ac1434.Add(__obf_68da4d6990322abc.__obf_830e351d96ac1434, __obf_acf4ecfba54e0451)
		} else {
			__obf_68da4d6990322abc.__obf_830e351d96ac1434.Sub(__obf_68da4d6990322abc.__obf_830e351d96ac1434, __obf_acf4ecfba54e0451)
		}
	}

	return __obf_68da4d6990322abc
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
func (__obf_c8d718debc5f6382 Decimal) RoundCash(__obf_2a03d45ad2a75a03 uint8) Decimal {
	var __obf_b23f4910670a27d7 *big.Int
	switch __obf_2a03d45ad2a75a03 {
	case 5:
		__obf_b23f4910670a27d7 = __obf_a370817798d4bdc4
	case 10:
		__obf_b23f4910670a27d7 = __obf_125c1d0eb0d00860
	case 25:
		__obf_b23f4910670a27d7 = __obf_c43b7c699e959605
	case 50:
		__obf_b23f4910670a27d7 = __obf_7f85275c7a28a30a
	case 100:
		__obf_b23f4910670a27d7 = __obf_acf4ecfba54e0451
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_2a03d45ad2a75a03))
	}
	__obf_986e7edebeba9ca9 := Decimal{
		__obf_830e351d96ac1434: __obf_b23f4910670a27d7,
	}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_c8d718debc5f6382.Mul(__obf_986e7edebeba9ca9).Round(0).Div(__obf_986e7edebeba9ca9).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_c8d718debc5f6382 Decimal) Floor() Decimal {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()

	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 >= 0 {
		return __obf_c8d718debc5f6382
	}

	__obf_e7725bcd1ee3b1c3 := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_e7725bcd1ee3b1c3.Exp(__obf_e7725bcd1ee3b1c3, big.NewInt(-int64(__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3)), nil)

	__obf_ef9b2770cf8a72e8 := new(big.Int).Div(__obf_c8d718debc5f6382.__obf_830e351d96ac1434, __obf_e7725bcd1ee3b1c3)
	return Decimal{__obf_830e351d96ac1434: __obf_ef9b2770cf8a72e8, __obf_e7725bcd1ee3b1c3: 0}
}

// Ceil returns the nearest integer value greater than or equal to d.
func (__obf_c8d718debc5f6382 Decimal) Ceil() Decimal {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()

	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 >= 0 {
		return __obf_c8d718debc5f6382
	}

	__obf_e7725bcd1ee3b1c3 := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_e7725bcd1ee3b1c3.Exp(__obf_e7725bcd1ee3b1c3, big.NewInt(-int64(__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3)), nil)

	__obf_ef9b2770cf8a72e8, __obf_4931dd9c612fb4d2 := new(big.Int).DivMod(__obf_c8d718debc5f6382.__obf_830e351d96ac1434, __obf_e7725bcd1ee3b1c3, new(big.Int))
	if __obf_4931dd9c612fb4d2.Cmp(__obf_ed0e7e529c925e39) != 0 {
		__obf_ef9b2770cf8a72e8.Add(__obf_ef9b2770cf8a72e8, __obf_acf4ecfba54e0451)
	}
	return Decimal{__obf_830e351d96ac1434: __obf_ef9b2770cf8a72e8, __obf_e7725bcd1ee3b1c3: 0}
}

// Truncate truncates off digits from the number, without rounding.
//
// NOTE: precision is the last digit that will not be truncated (must be >= 0).
//
// Example:
//
//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
func (__obf_c8d718debc5f6382 Decimal) Truncate(__obf_a2144abdd719f349 int32) Decimal {
	__obf_c8d718debc5f6382.__obf_b05475f1c7ee2e48()
	if __obf_a2144abdd719f349 >= 0 && -__obf_a2144abdd719f349 > __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 {
		return __obf_c8d718debc5f6382.__obf_e40699538ea665de(-__obf_a2144abdd719f349)
	}
	return __obf_c8d718debc5f6382
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_c8d718debc5f6382 *Decimal) UnmarshalJSON(__obf_3ffb444347d9973e []byte) error {
	if string(__obf_3ffb444347d9973e) == "null" {
		return nil
	}

	__obf_be4e99fccf2d81fa, __obf_95437c99969367dc := __obf_8ea3065ae03fb147(__obf_3ffb444347d9973e)
	if __obf_95437c99969367dc != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_3ffb444347d9973e, __obf_95437c99969367dc)
	}

	__obf_af53385fca67d325, __obf_95437c99969367dc := NewFromString(__obf_be4e99fccf2d81fa)
	*__obf_c8d718debc5f6382 = __obf_af53385fca67d325
	if __obf_95437c99969367dc != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_be4e99fccf2d81fa, __obf_95437c99969367dc)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_c8d718debc5f6382 Decimal) MarshalJSON() ([]byte, error) {
	var __obf_be4e99fccf2d81fa string
	if MarshalJSONWithoutQuotes {
		__obf_be4e99fccf2d81fa = __obf_c8d718debc5f6382.String()
	} else {
		__obf_be4e99fccf2d81fa = "\"" + __obf_c8d718debc5f6382.String() + "\""
	}
	return []byte(__obf_be4e99fccf2d81fa), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_c8d718debc5f6382 *Decimal) UnmarshalBinary(__obf_5c14e14c1f1c7843 []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_5c14e14c1f1c7843) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_5c14e14c1f1c7843, len(__obf_5c14e14c1f1c7843))
	}

	// Extract the exponent
	__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 = int32(binary.BigEndian.Uint32(__obf_5c14e14c1f1c7843[:4]))

	// Extract the value
	__obf_c8d718debc5f6382.__obf_830e351d96ac1434 = new(big.Int)
	if __obf_95437c99969367dc := __obf_c8d718debc5f6382.__obf_830e351d96ac1434.GobDecode(__obf_5c14e14c1f1c7843[4:]); __obf_95437c99969367dc != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_5c14e14c1f1c7843, __obf_95437c99969367dc)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_c8d718debc5f6382 Decimal) MarshalBinary() (__obf_5c14e14c1f1c7843 []byte, __obf_95437c99969367dc error) {
	// exp is written first, but encode value first to know output size
	var __obf_436e43bc299d9eaf []byte
	if __obf_436e43bc299d9eaf, __obf_95437c99969367dc = __obf_c8d718debc5f6382.__obf_830e351d96ac1434.GobEncode(); __obf_95437c99969367dc != nil {
		return nil, __obf_95437c99969367dc
	}

	// Write the exponent in front, since it's a fixed size
	__obf_51979dda959edabb := make([]byte, 4, len(__obf_436e43bc299d9eaf)+4)
	binary.BigEndian.PutUint32(__obf_51979dda959edabb, uint32(__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3))

	// Return the byte array
	return append(__obf_51979dda959edabb, __obf_436e43bc299d9eaf...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_c8d718debc5f6382 *Decimal) Scan(__obf_830e351d96ac1434 any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_41f2660d1e0d4d8e := __obf_830e351d96ac1434.(type) {

	case float32:
		*__obf_c8d718debc5f6382 = NewFromFloat(float64(__obf_41f2660d1e0d4d8e))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_c8d718debc5f6382 = NewFromFloat(__obf_41f2660d1e0d4d8e)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_c8d718debc5f6382 = New(__obf_41f2660d1e0d4d8e, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_c8d718debc5f6382 = NewFromUint64(__obf_41f2660d1e0d4d8e)
		return nil

	default:
		// default is trying to interpret value stored as string
		__obf_be4e99fccf2d81fa, __obf_95437c99969367dc := __obf_8ea3065ae03fb147(__obf_41f2660d1e0d4d8e)
		if __obf_95437c99969367dc != nil {
			return __obf_95437c99969367dc
		}
		*__obf_c8d718debc5f6382, __obf_95437c99969367dc = NewFromString(__obf_be4e99fccf2d81fa)
		return __obf_95437c99969367dc
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_c8d718debc5f6382 Decimal) Value() (driver.Value, error) {
	return __obf_c8d718debc5f6382.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_c8d718debc5f6382 *Decimal) UnmarshalText(__obf_b4dbd94a5835dd45 []byte) error {
	__obf_be4e99fccf2d81fa := string(__obf_b4dbd94a5835dd45)

	__obf_e7bd985d39dcdcf8, __obf_95437c99969367dc := NewFromString(__obf_be4e99fccf2d81fa)
	*__obf_c8d718debc5f6382 = __obf_e7bd985d39dcdcf8
	if __obf_95437c99969367dc != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_be4e99fccf2d81fa, __obf_95437c99969367dc)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_c8d718debc5f6382 Decimal) MarshalText() (__obf_b4dbd94a5835dd45 []byte, __obf_95437c99969367dc error) {
	return []byte(__obf_c8d718debc5f6382.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_c8d718debc5f6382 Decimal) GobEncode() ([]byte, error) {
	return __obf_c8d718debc5f6382.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_c8d718debc5f6382 *Decimal) GobDecode(__obf_5c14e14c1f1c7843 []byte) error {
	return __obf_c8d718debc5f6382.UnmarshalBinary(__obf_5c14e14c1f1c7843)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_c8d718debc5f6382 Decimal) StringScaled(__obf_e7725bcd1ee3b1c3 int32) string {
	return __obf_c8d718debc5f6382.__obf_e40699538ea665de(__obf_e7725bcd1ee3b1c3).String()
}

func (__obf_c8d718debc5f6382 Decimal) string(__obf_6b38990c5703d83c bool) string {
	if __obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3 >= 0 {
		return __obf_c8d718debc5f6382.__obf_e40699538ea665de(0).__obf_830e351d96ac1434.String()
	}

	__obf_8b38a12ace799621 := new(big.Int).Abs(__obf_c8d718debc5f6382.__obf_830e351d96ac1434)
	__obf_be4e99fccf2d81fa := __obf_8b38a12ace799621.String()

	var __obf_95ef1d102c878910, __obf_bd63aad168eb188a string

	// NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
	// and you are on a 32-bit machine. Won't fix this super-edge case.
	__obf_e44aaf11f7b9e182 := int(__obf_c8d718debc5f6382.__obf_e7725bcd1ee3b1c3)
	if len(__obf_be4e99fccf2d81fa) > -__obf_e44aaf11f7b9e182 {
		__obf_95ef1d102c878910 = __obf_be4e99fccf2d81fa[:len(__obf_be4e99fccf2d81fa)+__obf_e44aaf11f7b9e182]
		__obf_bd63aad168eb188a = __obf_be4e99fccf2d81fa[len(__obf_be4e99fccf2d81fa)+__obf_e44aaf11f7b9e182:]
	} else {
		__obf_95ef1d102c878910 = "0"

		__obf_efd56f54a801259a := -__obf_e44aaf11f7b9e182 - len(__obf_be4e99fccf2d81fa)
		__obf_bd63aad168eb188a = strings.Repeat("0", __obf_efd56f54a801259a) + __obf_be4e99fccf2d81fa
	}

	if __obf_6b38990c5703d83c {
		__obf_4aeff04bb586b943 := len(__obf_bd63aad168eb188a) - 1
		for ; __obf_4aeff04bb586b943 >= 0; __obf_4aeff04bb586b943-- {
			if __obf_bd63aad168eb188a[__obf_4aeff04bb586b943] != '0' {
				break
			}
		}
		__obf_bd63aad168eb188a = __obf_bd63aad168eb188a[:__obf_4aeff04bb586b943+1]
	}

	__obf_f711dd03dda2e586 := __obf_95ef1d102c878910
	if len(__obf_bd63aad168eb188a) > 0 {
		__obf_f711dd03dda2e586 += "." + __obf_bd63aad168eb188a
	}

	if __obf_c8d718debc5f6382.__obf_830e351d96ac1434.Sign() < 0 {
		return "-" + __obf_f711dd03dda2e586
	}

	return __obf_f711dd03dda2e586
}

func (__obf_c8d718debc5f6382 *Decimal) __obf_b05475f1c7ee2e48() {
	if __obf_c8d718debc5f6382.__obf_830e351d96ac1434 == nil {
		__obf_c8d718debc5f6382.__obf_830e351d96ac1434 = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_a3a66103cfff0c6d Decimal, __obf_33726d75893cc8d2 ...Decimal) Decimal {
	__obf_63eb5f12c8b6946c := __obf_a3a66103cfff0c6d
	for _, __obf_d97db6c26c6733a3 := range __obf_33726d75893cc8d2 {
		if __obf_d97db6c26c6733a3.Cmp(__obf_63eb5f12c8b6946c) < 0 {
			__obf_63eb5f12c8b6946c = __obf_d97db6c26c6733a3
		}
	}
	return __obf_63eb5f12c8b6946c
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_a3a66103cfff0c6d Decimal, __obf_33726d75893cc8d2 ...Decimal) Decimal {
	__obf_63eb5f12c8b6946c := __obf_a3a66103cfff0c6d
	for _, __obf_d97db6c26c6733a3 := range __obf_33726d75893cc8d2 {
		if __obf_d97db6c26c6733a3.Cmp(__obf_63eb5f12c8b6946c) > 0 {
			__obf_63eb5f12c8b6946c = __obf_d97db6c26c6733a3
		}
	}
	return __obf_63eb5f12c8b6946c
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_a3a66103cfff0c6d Decimal, __obf_33726d75893cc8d2 ...Decimal) Decimal {
	__obf_d9b6552619ae4ee6 := __obf_a3a66103cfff0c6d
	for _, __obf_d97db6c26c6733a3 := range __obf_33726d75893cc8d2 {
		__obf_d9b6552619ae4ee6 = __obf_d9b6552619ae4ee6.Add(__obf_d97db6c26c6733a3)
	}

	return __obf_d9b6552619ae4ee6
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_a3a66103cfff0c6d Decimal, __obf_33726d75893cc8d2 ...Decimal) Decimal {
	__obf_21b68de6b7f5c2bd := New(int64(len(__obf_33726d75893cc8d2)+1), 0)
	__obf_16b88190ca2a680e := Sum(__obf_a3a66103cfff0c6d, __obf_33726d75893cc8d2...)
	return __obf_16b88190ca2a680e.Div(__obf_21b68de6b7f5c2bd)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_b7439b49e0b40adb Decimal, __obf_4b29119296c7019d Decimal) (Decimal, Decimal) {
	__obf_b7439b49e0b40adb.__obf_b05475f1c7ee2e48()
	__obf_4b29119296c7019d.__obf_b05475f1c7ee2e48()

	if __obf_b7439b49e0b40adb.__obf_e7725bcd1ee3b1c3 < __obf_4b29119296c7019d.__obf_e7725bcd1ee3b1c3 {
		return __obf_b7439b49e0b40adb, __obf_4b29119296c7019d.__obf_e40699538ea665de(__obf_b7439b49e0b40adb.__obf_e7725bcd1ee3b1c3)
	} else if __obf_b7439b49e0b40adb.__obf_e7725bcd1ee3b1c3 > __obf_4b29119296c7019d.__obf_e7725bcd1ee3b1c3 {
		return __obf_b7439b49e0b40adb.__obf_e40699538ea665de(__obf_4b29119296c7019d.__obf_e7725bcd1ee3b1c3), __obf_4b29119296c7019d
	}

	return __obf_b7439b49e0b40adb, __obf_4b29119296c7019d
}

func __obf_8ea3065ae03fb147(__obf_830e351d96ac1434 any) (string, error) {
	var __obf_5ef2a5d7452960a0 []byte

	switch __obf_41f2660d1e0d4d8e := __obf_830e351d96ac1434.(type) {
	case string:
		__obf_5ef2a5d7452960a0 = []byte(__obf_41f2660d1e0d4d8e)
	case []byte:
		__obf_5ef2a5d7452960a0 = __obf_41f2660d1e0d4d8e
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_830e351d96ac1434, __obf_830e351d96ac1434)
	}

	// If the amount is quoted, strip the quotes
	if len(__obf_5ef2a5d7452960a0) > 2 && __obf_5ef2a5d7452960a0[0] == '"' && __obf_5ef2a5d7452960a0[len(__obf_5ef2a5d7452960a0)-1] == '"' {
		__obf_5ef2a5d7452960a0 = __obf_5ef2a5d7452960a0[1 : len(__obf_5ef2a5d7452960a0)-1]
	}
	return string(__obf_5ef2a5d7452960a0), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the datautil.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_c8d718debc5f6382 Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_c8d718debc5f6382,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_c8d718debc5f6382 *NullDecimal) Scan(__obf_830e351d96ac1434 any) error {
	if __obf_830e351d96ac1434 == nil {
		__obf_c8d718debc5f6382.Valid = false
		return nil
	}
	__obf_c8d718debc5f6382.Valid = true
	return __obf_c8d718debc5f6382.Decimal.Scan(__obf_830e351d96ac1434)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_c8d718debc5f6382 NullDecimal) Value() (driver.Value, error) {
	if !__obf_c8d718debc5f6382.Valid {
		return nil, nil
	}
	return __obf_c8d718debc5f6382.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_c8d718debc5f6382 *NullDecimal) UnmarshalJSON(__obf_3ffb444347d9973e []byte) error {
	if string(__obf_3ffb444347d9973e) == "null" {
		__obf_c8d718debc5f6382.Valid = false
		return nil
	}
	__obf_c8d718debc5f6382.Valid = true
	return __obf_c8d718debc5f6382.Decimal.UnmarshalJSON(__obf_3ffb444347d9973e)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_c8d718debc5f6382 NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_c8d718debc5f6382.Valid {
		return []byte("null"), nil
	}
	return __obf_c8d718debc5f6382.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_c8d718debc5f6382 *NullDecimal) UnmarshalText(__obf_b4dbd94a5835dd45 []byte) error {
	__obf_be4e99fccf2d81fa := string(__obf_b4dbd94a5835dd45)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_be4e99fccf2d81fa == "" {
		__obf_c8d718debc5f6382.Valid = false
		return nil
	}
	if __obf_95437c99969367dc := __obf_c8d718debc5f6382.Decimal.UnmarshalText(__obf_b4dbd94a5835dd45); __obf_95437c99969367dc != nil {
		__obf_c8d718debc5f6382.Valid = false
		return __obf_95437c99969367dc
	}
	__obf_c8d718debc5f6382.Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_c8d718debc5f6382 NullDecimal) MarshalText() (__obf_b4dbd94a5835dd45 []byte, __obf_95437c99969367dc error) {
	if !__obf_c8d718debc5f6382.Valid {
		return []byte{}, nil
	}
	return __obf_c8d718debc5f6382.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_c8d718debc5f6382 Decimal) Atan() Decimal {
	if __obf_c8d718debc5f6382.Equal(NewFromFloat(0.0)) {
		return __obf_c8d718debc5f6382
	}
	if __obf_c8d718debc5f6382.GreaterThan(NewFromFloat(0.0)) {
		return __obf_c8d718debc5f6382.__obf_6ab66d85fecb4c6e()
	}
	return __obf_c8d718debc5f6382.Neg().__obf_6ab66d85fecb4c6e().Neg()
}

func (__obf_c8d718debc5f6382 Decimal) __obf_6543e4c85f7e43c7() Decimal {
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
	__obf_ef9b2770cf8a72e8 := __obf_c8d718debc5f6382.Mul(__obf_c8d718debc5f6382)
	__obf_2cdc1821aea16b10 := P0.Mul(__obf_ef9b2770cf8a72e8).Add(P1).Mul(__obf_ef9b2770cf8a72e8).Add(P2).Mul(__obf_ef9b2770cf8a72e8).Add(P3).Mul(__obf_ef9b2770cf8a72e8).Add(P4).Mul(__obf_ef9b2770cf8a72e8)
	__obf_98d202aef8757c75 := __obf_ef9b2770cf8a72e8.Add(Q0).Mul(__obf_ef9b2770cf8a72e8).Add(Q1).Mul(__obf_ef9b2770cf8a72e8).Add(Q2).Mul(__obf_ef9b2770cf8a72e8).Add(Q3).Mul(__obf_ef9b2770cf8a72e8).Add(Q4)
	__obf_ef9b2770cf8a72e8 = __obf_2cdc1821aea16b10.Div(__obf_98d202aef8757c75)
	__obf_ef9b2770cf8a72e8 = __obf_c8d718debc5f6382.Mul(__obf_ef9b2770cf8a72e8).Add(__obf_c8d718debc5f6382)
	return __obf_ef9b2770cf8a72e8
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_c8d718debc5f6382 Decimal) __obf_6ab66d85fecb4c6e() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)      // tan(3*pi/8)
	__obf_a07a137e7c4adc75 := NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_c8d718debc5f6382.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_c8d718debc5f6382.__obf_6543e4c85f7e43c7()
	}
	if __obf_c8d718debc5f6382.GreaterThan(Tan3pio8) {
		return __obf_a07a137e7c4adc75.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_c8d718debc5f6382).__obf_6543e4c85f7e43c7()).Add(Morebits)
	}
	return __obf_a07a137e7c4adc75.Div(NewFromFloat(4.0)).Add((__obf_c8d718debc5f6382.Sub(NewFromFloat(1.0)).Div(__obf_c8d718debc5f6382.Add(NewFromFloat(1.0)))).__obf_6543e4c85f7e43c7()).Add(NewFromFloat(0.5).Mul(Morebits))
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
func (__obf_c8d718debc5f6382 Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_c8d718debc5f6382.Equal(NewFromFloat(0.0)) {
		return __obf_c8d718debc5f6382
	}
	// make argument positive but save the sign
	__obf_20d020ec018f3926 := false
	if __obf_c8d718debc5f6382.LessThan(NewFromFloat(0.0)) {
		__obf_c8d718debc5f6382 = __obf_c8d718debc5f6382.Neg()
		__obf_20d020ec018f3926 = true
	}

	__obf_544b469bd889a527 := __obf_c8d718debc5f6382.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_86e45493a546eb0f := NewFromFloat(float64(__obf_544b469bd889a527)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_544b469bd889a527&1 == 1 {
		__obf_544b469bd889a527++
		__obf_86e45493a546eb0f = __obf_86e45493a546eb0f.Add(NewFromFloat(1.0))
	}
	__obf_544b469bd889a527 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_544b469bd889a527 > 3 {
		__obf_20d020ec018f3926 = !__obf_20d020ec018f3926
		__obf_544b469bd889a527 -= 4
	}
	__obf_ef9b2770cf8a72e8 := __obf_c8d718debc5f6382.Sub(__obf_86e45493a546eb0f.Mul(PI4A)).Sub(__obf_86e45493a546eb0f.Mul(PI4B)).Sub(__obf_86e45493a546eb0f.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_79c31c93debc0112 := __obf_ef9b2770cf8a72e8.Mul(__obf_ef9b2770cf8a72e8)

	if __obf_544b469bd889a527 == 1 || __obf_544b469bd889a527 == 2 {
		__obf_3180ee6d5b6e1a09 := __obf_79c31c93debc0112.Mul(__obf_79c31c93debc0112).Mul(_cos[0].Mul(__obf_79c31c93debc0112).Add(_cos[1]).Mul(__obf_79c31c93debc0112).Add(_cos[2]).Mul(__obf_79c31c93debc0112).Add(_cos[3]).Mul(__obf_79c31c93debc0112).Add(_cos[4]).Mul(__obf_79c31c93debc0112).Add(_cos[5]))
		__obf_86e45493a546eb0f = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_79c31c93debc0112)).Add(__obf_3180ee6d5b6e1a09)
	} else {
		__obf_86e45493a546eb0f = __obf_ef9b2770cf8a72e8.Add(__obf_ef9b2770cf8a72e8.Mul(__obf_79c31c93debc0112).Mul(_sin[0].Mul(__obf_79c31c93debc0112).Add(_sin[1]).Mul(__obf_79c31c93debc0112).Add(_sin[2]).Mul(__obf_79c31c93debc0112).Add(_sin[3]).Mul(__obf_79c31c93debc0112).Add(_sin[4]).Mul(__obf_79c31c93debc0112).Add(_sin[5])))
	}
	if __obf_20d020ec018f3926 {
		__obf_86e45493a546eb0f = __obf_86e45493a546eb0f.Neg()
	}
	return __obf_86e45493a546eb0f
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
func (__obf_c8d718debc5f6382 Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	// make argument positive
	__obf_20d020ec018f3926 := false
	if __obf_c8d718debc5f6382.LessThan(NewFromFloat(0.0)) {
		__obf_c8d718debc5f6382 = __obf_c8d718debc5f6382.Neg()
	}

	__obf_544b469bd889a527 := __obf_c8d718debc5f6382.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_86e45493a546eb0f := NewFromFloat(float64(__obf_544b469bd889a527)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_544b469bd889a527&1 == 1 {
		__obf_544b469bd889a527++
		__obf_86e45493a546eb0f = __obf_86e45493a546eb0f.Add(NewFromFloat(1.0))
	}
	__obf_544b469bd889a527 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_544b469bd889a527 > 3 {
		__obf_20d020ec018f3926 = !__obf_20d020ec018f3926
		__obf_544b469bd889a527 -= 4
	}
	if __obf_544b469bd889a527 > 1 {
		__obf_20d020ec018f3926 = !__obf_20d020ec018f3926
	}

	__obf_ef9b2770cf8a72e8 := __obf_c8d718debc5f6382.Sub(__obf_86e45493a546eb0f.Mul(PI4A)).Sub(__obf_86e45493a546eb0f.Mul(PI4B)).Sub(__obf_86e45493a546eb0f.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_79c31c93debc0112 := __obf_ef9b2770cf8a72e8.Mul(__obf_ef9b2770cf8a72e8)

	if __obf_544b469bd889a527 == 1 || __obf_544b469bd889a527 == 2 {
		__obf_86e45493a546eb0f = __obf_ef9b2770cf8a72e8.Add(__obf_ef9b2770cf8a72e8.Mul(__obf_79c31c93debc0112).Mul(_sin[0].Mul(__obf_79c31c93debc0112).Add(_sin[1]).Mul(__obf_79c31c93debc0112).Add(_sin[2]).Mul(__obf_79c31c93debc0112).Add(_sin[3]).Mul(__obf_79c31c93debc0112).Add(_sin[4]).Mul(__obf_79c31c93debc0112).Add(_sin[5])))
	} else {
		__obf_3180ee6d5b6e1a09 := __obf_79c31c93debc0112.Mul(__obf_79c31c93debc0112).Mul(_cos[0].Mul(__obf_79c31c93debc0112).Add(_cos[1]).Mul(__obf_79c31c93debc0112).Add(_cos[2]).Mul(__obf_79c31c93debc0112).Add(_cos[3]).Mul(__obf_79c31c93debc0112).Add(_cos[4]).Mul(__obf_79c31c93debc0112).Add(_cos[5]))
		__obf_86e45493a546eb0f = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_79c31c93debc0112)).Add(__obf_3180ee6d5b6e1a09)
	}
	if __obf_20d020ec018f3926 {
		__obf_86e45493a546eb0f = __obf_86e45493a546eb0f.Neg()
	}
	return __obf_86e45493a546eb0f
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
func (__obf_c8d718debc5f6382 Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_c8d718debc5f6382.Equal(NewFromFloat(0.0)) {
		return __obf_c8d718debc5f6382
	}

	// make argument positive but save the sign
	__obf_20d020ec018f3926 := false
	if __obf_c8d718debc5f6382.LessThan(NewFromFloat(0.0)) {
		__obf_c8d718debc5f6382 = __obf_c8d718debc5f6382.Neg()
		__obf_20d020ec018f3926 = true
	}

	__obf_544b469bd889a527 := __obf_c8d718debc5f6382.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_86e45493a546eb0f := NewFromFloat(float64(__obf_544b469bd889a527)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_544b469bd889a527&1 == 1 {
		__obf_544b469bd889a527++
		__obf_86e45493a546eb0f = __obf_86e45493a546eb0f.Add(NewFromFloat(1.0))
	}

	__obf_ef9b2770cf8a72e8 := __obf_c8d718debc5f6382.Sub(__obf_86e45493a546eb0f.Mul(PI4A)).Sub(__obf_86e45493a546eb0f.Mul(PI4B)).Sub(__obf_86e45493a546eb0f.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_79c31c93debc0112 := __obf_ef9b2770cf8a72e8.Mul(__obf_ef9b2770cf8a72e8)

	if __obf_79c31c93debc0112.GreaterThan(NewFromFloat(1e-14)) {
		__obf_3180ee6d5b6e1a09 := __obf_79c31c93debc0112.Mul(_tanP[0].Mul(__obf_79c31c93debc0112).Add(_tanP[1]).Mul(__obf_79c31c93debc0112).Add(_tanP[2]))
		__obf_c755d92ea0063ee1 := __obf_79c31c93debc0112.Add(_tanQ[1]).Mul(__obf_79c31c93debc0112).Add(_tanQ[2]).Mul(__obf_79c31c93debc0112).Add(_tanQ[3]).Mul(__obf_79c31c93debc0112).Add(_tanQ[4])
		__obf_86e45493a546eb0f = __obf_ef9b2770cf8a72e8.Add(__obf_ef9b2770cf8a72e8.Mul(__obf_3180ee6d5b6e1a09.Div(__obf_c755d92ea0063ee1)))
	} else {
		__obf_86e45493a546eb0f = __obf_ef9b2770cf8a72e8
	}
	if __obf_544b469bd889a527&2 == 2 {
		__obf_86e45493a546eb0f = NewFromFloat(-1.0).Div(__obf_86e45493a546eb0f)
	}
	if __obf_20d020ec018f3926 {
		__obf_86e45493a546eb0f = __obf_86e45493a546eb0f.Neg()
	}
	return __obf_86e45493a546eb0f
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

type __obf_af53385fca67d325 struct {
	__obf_c8d718debc5f6382 [800]byte // digits, big-endian representation
	__obf_b2da5b1e27e361da int       // number of digits used
	__obf_27d38fb675e867eb int       // decimal point
	__obf_6bf80ec73009a641 bool      // negative flag
	__obf_bfcdbd91a38f8023 bool      // discarded nonzero digits beyond d[:nd]
}

func (__obf_7e6fccb9375af12b *__obf_af53385fca67d325) String() string {
	__obf_648608c066f38982 := 10 + __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da
	if __obf_7e6fccb9375af12b.__obf_27d38fb675e867eb > 0 {
		__obf_648608c066f38982 += __obf_7e6fccb9375af12b.__obf_27d38fb675e867eb
	}
	if __obf_7e6fccb9375af12b.__obf_27d38fb675e867eb < 0 {
		__obf_648608c066f38982 += -__obf_7e6fccb9375af12b.__obf_27d38fb675e867eb
	}

	__obf_4e90c18b99bf124d := make([]byte, __obf_648608c066f38982)
	__obf_3180ee6d5b6e1a09 := 0
	switch {
	case __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da == 0:
		return "0"

	case __obf_7e6fccb9375af12b.__obf_27d38fb675e867eb <= 0:
		// zeros fill space between decimal point and digits
		__obf_4e90c18b99bf124d[__obf_3180ee6d5b6e1a09] = '0'
		__obf_3180ee6d5b6e1a09++
		__obf_4e90c18b99bf124d[__obf_3180ee6d5b6e1a09] = '.'
		__obf_3180ee6d5b6e1a09++
		__obf_3180ee6d5b6e1a09 += __obf_af258bd5df546e11(__obf_4e90c18b99bf124d[__obf_3180ee6d5b6e1a09 : __obf_3180ee6d5b6e1a09+-__obf_7e6fccb9375af12b.__obf_27d38fb675e867eb])
		__obf_3180ee6d5b6e1a09 += copy(__obf_4e90c18b99bf124d[__obf_3180ee6d5b6e1a09:], __obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[0:__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da])

	case __obf_7e6fccb9375af12b.__obf_27d38fb675e867eb < __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da:
		// decimal point in middle of digits
		__obf_3180ee6d5b6e1a09 += copy(__obf_4e90c18b99bf124d[__obf_3180ee6d5b6e1a09:], __obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[0:__obf_7e6fccb9375af12b.__obf_27d38fb675e867eb])
		__obf_4e90c18b99bf124d[__obf_3180ee6d5b6e1a09] = '.'
		__obf_3180ee6d5b6e1a09++
		__obf_3180ee6d5b6e1a09 += copy(__obf_4e90c18b99bf124d[__obf_3180ee6d5b6e1a09:], __obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_7e6fccb9375af12b.__obf_27d38fb675e867eb:__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da])

	default:
		// zeros fill space between digits and decimal point
		__obf_3180ee6d5b6e1a09 += copy(__obf_4e90c18b99bf124d[__obf_3180ee6d5b6e1a09:], __obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[0:__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da])
		__obf_3180ee6d5b6e1a09 += __obf_af258bd5df546e11(__obf_4e90c18b99bf124d[__obf_3180ee6d5b6e1a09 : __obf_3180ee6d5b6e1a09+__obf_7e6fccb9375af12b.__obf_27d38fb675e867eb-__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da])
	}
	return string(__obf_4e90c18b99bf124d[0:__obf_3180ee6d5b6e1a09])
}

func __obf_af258bd5df546e11(__obf_0a781472a065cf6c []byte) int {
	for __obf_4aeff04bb586b943 := range __obf_0a781472a065cf6c {
		__obf_0a781472a065cf6c[__obf_4aeff04bb586b943] = '0'
	}
	return len(__obf_0a781472a065cf6c)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_a1e6598e370d18e9(__obf_7e6fccb9375af12b *__obf_af53385fca67d325) {
	for __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da > 0 && __obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da-1] == '0' {
		__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da--
	}
	if __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da == 0 {
		__obf_7e6fccb9375af12b.__obf_27d38fb675e867eb = 0
	}
}

// Assign v to a.
func (__obf_7e6fccb9375af12b *__obf_af53385fca67d325) Assign(__obf_41f2660d1e0d4d8e uint64) {
	var __obf_4e90c18b99bf124d [24]byte

	// Write reversed decimal in buf.
	__obf_648608c066f38982 := 0
	for __obf_41f2660d1e0d4d8e > 0 {
		__obf_227c27f468975f39 := __obf_41f2660d1e0d4d8e / 10
		__obf_41f2660d1e0d4d8e -= 10 * __obf_227c27f468975f39
		__obf_4e90c18b99bf124d[__obf_648608c066f38982] = byte(__obf_41f2660d1e0d4d8e + '0')
		__obf_648608c066f38982++
		__obf_41f2660d1e0d4d8e = __obf_227c27f468975f39
	}

	// Reverse again to produce forward decimal in a.d.
	__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da = 0
	for __obf_648608c066f38982--; __obf_648608c066f38982 >= 0; __obf_648608c066f38982-- {
		__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da] = __obf_4e90c18b99bf124d[__obf_648608c066f38982]
		__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da++
	}
	__obf_7e6fccb9375af12b.__obf_27d38fb675e867eb = __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da
	__obf_a1e6598e370d18e9(__obf_7e6fccb9375af12b)
}

// Maximum shift that we can do in one pass without overflow.
// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
const __obf_92c267f3c3102799 = 32 << (^uint(0) >> 63)
const __obf_8d718f05fde0c798 = __obf_92c267f3c3102799 - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_83c71f4bc1c48920(__obf_7e6fccb9375af12b *__obf_af53385fca67d325, __obf_88a1a18216771a9f uint) {
	__obf_791a40daaf73f23a := 0 // read pointer
	__obf_3180ee6d5b6e1a09 := 0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_648608c066f38982 uint
	for ; __obf_648608c066f38982>>__obf_88a1a18216771a9f == 0; __obf_791a40daaf73f23a++ {
		if __obf_791a40daaf73f23a >= __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da {
			if __obf_648608c066f38982 == 0 {
				// a == 0; shouldn't get here, but handle anyway.
				__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da = 0
				return
			}
			for __obf_648608c066f38982>>__obf_88a1a18216771a9f == 0 {
				__obf_648608c066f38982 = __obf_648608c066f38982 * 10
				__obf_791a40daaf73f23a++
			}
			break
		}
		__obf_2b375914089c38b5 := uint(__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_791a40daaf73f23a])
		__obf_648608c066f38982 = __obf_648608c066f38982*10 + __obf_2b375914089c38b5 - '0'
	}
	__obf_7e6fccb9375af12b.__obf_27d38fb675e867eb -= __obf_791a40daaf73f23a - 1

	var __obf_e8de0afbdf926e12 uint = (1 << __obf_88a1a18216771a9f) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_791a40daaf73f23a < __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da; __obf_791a40daaf73f23a++ {
		__obf_2b375914089c38b5 := uint(__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_791a40daaf73f23a])
		__obf_2057b9b2094a9285 := __obf_648608c066f38982 >> __obf_88a1a18216771a9f
		__obf_648608c066f38982 &= __obf_e8de0afbdf926e12
		__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_3180ee6d5b6e1a09] = byte(__obf_2057b9b2094a9285 + '0')
		__obf_3180ee6d5b6e1a09++
		__obf_648608c066f38982 = __obf_648608c066f38982*10 + __obf_2b375914089c38b5 - '0'
	}

	// Put down extra digits.
	for __obf_648608c066f38982 > 0 {
		__obf_2057b9b2094a9285 := __obf_648608c066f38982 >> __obf_88a1a18216771a9f
		__obf_648608c066f38982 &= __obf_e8de0afbdf926e12
		if __obf_3180ee6d5b6e1a09 < len(__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382) {
			__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_3180ee6d5b6e1a09] = byte(__obf_2057b9b2094a9285 + '0')
			__obf_3180ee6d5b6e1a09++
		} else if __obf_2057b9b2094a9285 > 0 {
			__obf_7e6fccb9375af12b.__obf_bfcdbd91a38f8023 = true
		}
		__obf_648608c066f38982 = __obf_648608c066f38982 * 10
	}

	__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da = __obf_3180ee6d5b6e1a09
	__obf_a1e6598e370d18e9(__obf_7e6fccb9375af12b)
}

// Cheat sheet for left shift: table indexed by shift count giving
// number of new digits that will be introduced by that shift.
//
// For example, leftcheats[4] = {2, "625"}.  That means that
// if we are shifting by 4 (multiplying by 16), it will add 2 digits
// when the string prefix is "625" through "999", and one fewer digit
// if the string prefix is "000" through "624".
//
// Credit for this trick goes to Ken.

type __obf_4530066a4cab2b37 struct {
	__obf_a20ddcf3b283bf88 int    // number of new digits
	__obf_ba0cbf932608724b string // minus one digit if original < a.
}

var __obf_82d0243396b1b20c = []__obf_4530066a4cab2b37{
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
func __obf_48fdf39b2766ebf8(__obf_a311363f116585b2 []byte, __obf_aa21bb8619396adc string) bool {
	for __obf_4aeff04bb586b943 := 0; __obf_4aeff04bb586b943 < len(__obf_aa21bb8619396adc); __obf_4aeff04bb586b943++ {
		if __obf_4aeff04bb586b943 >= len(__obf_a311363f116585b2) {
			return true
		}
		if __obf_a311363f116585b2[__obf_4aeff04bb586b943] != __obf_aa21bb8619396adc[__obf_4aeff04bb586b943] {
			return __obf_a311363f116585b2[__obf_4aeff04bb586b943] < __obf_aa21bb8619396adc[__obf_4aeff04bb586b943]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_433f441578eb03a8(__obf_7e6fccb9375af12b *__obf_af53385fca67d325, __obf_88a1a18216771a9f uint) {
	__obf_a20ddcf3b283bf88 := __obf_82d0243396b1b20c[__obf_88a1a18216771a9f].__obf_a20ddcf3b283bf88
	if __obf_48fdf39b2766ebf8(__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[0:__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da], __obf_82d0243396b1b20c[__obf_88a1a18216771a9f].__obf_ba0cbf932608724b) {
		__obf_a20ddcf3b283bf88--
	}

	__obf_791a40daaf73f23a := __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da                          // read index
	__obf_3180ee6d5b6e1a09 := __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da + __obf_a20ddcf3b283bf88 // write index

	// Pick up a digit, put down a digit.
	var __obf_648608c066f38982 uint
	for __obf_791a40daaf73f23a--; __obf_791a40daaf73f23a >= 0; __obf_791a40daaf73f23a-- {
		__obf_648608c066f38982 += (uint(__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_791a40daaf73f23a]) - '0') << __obf_88a1a18216771a9f
		__obf_af2ba60d95bf130f := __obf_648608c066f38982 / 10
		__obf_b565ce6f1dbaa587 := __obf_648608c066f38982 - 10*__obf_af2ba60d95bf130f
		__obf_3180ee6d5b6e1a09--
		if __obf_3180ee6d5b6e1a09 < len(__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382) {
			__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_3180ee6d5b6e1a09] = byte(__obf_b565ce6f1dbaa587 + '0')
		} else if __obf_b565ce6f1dbaa587 != 0 {
			__obf_7e6fccb9375af12b.__obf_bfcdbd91a38f8023 = true
		}
		__obf_648608c066f38982 = __obf_af2ba60d95bf130f
	}

	// Put down extra digits.
	for __obf_648608c066f38982 > 0 {
		__obf_af2ba60d95bf130f := __obf_648608c066f38982 / 10
		__obf_b565ce6f1dbaa587 := __obf_648608c066f38982 - 10*__obf_af2ba60d95bf130f
		__obf_3180ee6d5b6e1a09--
		if __obf_3180ee6d5b6e1a09 < len(__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382) {
			__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_3180ee6d5b6e1a09] = byte(__obf_b565ce6f1dbaa587 + '0')
		} else if __obf_b565ce6f1dbaa587 != 0 {
			__obf_7e6fccb9375af12b.__obf_bfcdbd91a38f8023 = true
		}
		__obf_648608c066f38982 = __obf_af2ba60d95bf130f
	}

	__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da += __obf_a20ddcf3b283bf88
	if __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da >= len(__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382) {
		__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da = len(__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382)
	}
	__obf_7e6fccb9375af12b.__obf_27d38fb675e867eb += __obf_a20ddcf3b283bf88
	__obf_a1e6598e370d18e9(__obf_7e6fccb9375af12b)
}

// Binary shift left (k > 0) or right (k < 0).
func (__obf_7e6fccb9375af12b *__obf_af53385fca67d325) Shift(__obf_88a1a18216771a9f int) {
	switch {
	case __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da == 0:
		// nothing to do: a == 0
	case __obf_88a1a18216771a9f > 0:
		for __obf_88a1a18216771a9f > __obf_8d718f05fde0c798 {
			__obf_433f441578eb03a8(__obf_7e6fccb9375af12b, __obf_8d718f05fde0c798)
			__obf_88a1a18216771a9f -= __obf_8d718f05fde0c798
		}
		__obf_433f441578eb03a8(__obf_7e6fccb9375af12b, uint(__obf_88a1a18216771a9f))
	case __obf_88a1a18216771a9f < 0:
		for __obf_88a1a18216771a9f < -__obf_8d718f05fde0c798 {
			__obf_83c71f4bc1c48920(__obf_7e6fccb9375af12b, __obf_8d718f05fde0c798)
			__obf_88a1a18216771a9f += __obf_8d718f05fde0c798
		}
		__obf_83c71f4bc1c48920(__obf_7e6fccb9375af12b, uint(-__obf_88a1a18216771a9f))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_a8d464be22aa495b(__obf_7e6fccb9375af12b *__obf_af53385fca67d325, __obf_b2da5b1e27e361da int) bool {
	if __obf_b2da5b1e27e361da < 0 || __obf_b2da5b1e27e361da >= __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da {
		return false
	}
	if __obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_b2da5b1e27e361da] == '5' && __obf_b2da5b1e27e361da+1 == __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da { // exactly halfway - round to even
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_7e6fccb9375af12b.__obf_bfcdbd91a38f8023 {
			return true
		}
		return __obf_b2da5b1e27e361da > 0 && (__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_b2da5b1e27e361da-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_b2da5b1e27e361da] >= '5'
}

// Round a to nd digits (or fewer).
// If nd is zero, it means we're rounding
// just to the left of the digits, as in
// 0.09 -> 0.1.
func (__obf_7e6fccb9375af12b *__obf_af53385fca67d325) Round(__obf_b2da5b1e27e361da int) {
	if __obf_b2da5b1e27e361da < 0 || __obf_b2da5b1e27e361da >= __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da {
		return
	}
	if __obf_a8d464be22aa495b(__obf_7e6fccb9375af12b, __obf_b2da5b1e27e361da) {
		__obf_7e6fccb9375af12b.RoundUp(__obf_b2da5b1e27e361da)
	} else {
		__obf_7e6fccb9375af12b.RoundDown(__obf_b2da5b1e27e361da)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_7e6fccb9375af12b *__obf_af53385fca67d325) RoundDown(__obf_b2da5b1e27e361da int) {
	if __obf_b2da5b1e27e361da < 0 || __obf_b2da5b1e27e361da >= __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da {
		return
	}
	__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da = __obf_b2da5b1e27e361da
	__obf_a1e6598e370d18e9(__obf_7e6fccb9375af12b)
}

// Round a up to nd digits (or fewer).
func (__obf_7e6fccb9375af12b *__obf_af53385fca67d325) RoundUp(__obf_b2da5b1e27e361da int) {
	if __obf_b2da5b1e27e361da < 0 || __obf_b2da5b1e27e361da >= __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da {
		return
	}

	// round up
	for __obf_4aeff04bb586b943 := __obf_b2da5b1e27e361da - 1; __obf_4aeff04bb586b943 >= 0; __obf_4aeff04bb586b943-- {
		__obf_2b375914089c38b5 := __obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_4aeff04bb586b943]
		if __obf_2b375914089c38b5 < '9' { // can stop after this digit
			__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_4aeff04bb586b943]++
			__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da = __obf_4aeff04bb586b943 + 1
			return
		}
	}

	// Number is all 9s.
	// Change to single 1 with adjusted decimal point.
	__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[0] = '1'
	__obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da = 1
	__obf_7e6fccb9375af12b.__obf_27d38fb675e867eb++
}

// Extract integer part, rounded appropriately.
// No guarantees about overflow.
func (__obf_7e6fccb9375af12b *__obf_af53385fca67d325) RoundedInteger() uint64 {
	if __obf_7e6fccb9375af12b.__obf_27d38fb675e867eb > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_4aeff04bb586b943 int
	__obf_648608c066f38982 := uint64(0)
	for __obf_4aeff04bb586b943 = 0; __obf_4aeff04bb586b943 < __obf_7e6fccb9375af12b.__obf_27d38fb675e867eb && __obf_4aeff04bb586b943 < __obf_7e6fccb9375af12b.__obf_b2da5b1e27e361da; __obf_4aeff04bb586b943++ {
		__obf_648608c066f38982 = __obf_648608c066f38982*10 + uint64(__obf_7e6fccb9375af12b.__obf_c8d718debc5f6382[__obf_4aeff04bb586b943]-'0')
	}
	for ; __obf_4aeff04bb586b943 < __obf_7e6fccb9375af12b.__obf_27d38fb675e867eb; __obf_4aeff04bb586b943++ {
		__obf_648608c066f38982 *= 10
	}
	if __obf_a8d464be22aa495b(__obf_7e6fccb9375af12b, __obf_7e6fccb9375af12b.__obf_27d38fb675e867eb) {
		__obf_648608c066f38982++
	}
	return __obf_648608c066f38982
}
