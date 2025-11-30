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
package __obf_a5d0662d16479f1d

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

var __obf_304be536ded8753b = big.NewInt(0)
var __obf_0ca31adf9d5bc652 = big.NewInt(1)
var __obf_51443023c3dfc1b9 = big.NewInt(2)
var __obf_e62aec2193220910 = big.NewInt(4)
var __obf_656e224a8e2fd8a8 = big.NewInt(5)
var __obf_8f06f98f079e4122 = big.NewInt(10)
var __obf_015884ffa111f28b = big.NewInt(20)

var __obf_edd8d74bea9c2de4 = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_ab2c02ab751b9f42 *big.Int
	__obf_d5f3440c0176542c int32 // NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.

}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_ab2c02ab751b9f42 int64, __obf_d5f3440c0176542c int32) Decimal {
	return Decimal{__obf_ab2c02ab751b9f42: big.NewInt(__obf_ab2c02ab751b9f42), __obf_d5f3440c0176542c: __obf_d5f3440c0176542c}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_ab2c02ab751b9f42 int64) Decimal {
	return Decimal{__obf_ab2c02ab751b9f42: big.NewInt(__obf_ab2c02ab751b9f42), __obf_d5f3440c0176542c: 0}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_ab2c02ab751b9f42 int32) Decimal {
	return Decimal{__obf_ab2c02ab751b9f42: big.NewInt(int64(__obf_ab2c02ab751b9f42)), __obf_d5f3440c0176542c: 0}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_ab2c02ab751b9f42 uint64) Decimal {
	return Decimal{__obf_ab2c02ab751b9f42: new(big.Int).SetUint64(__obf_ab2c02ab751b9f42), __obf_d5f3440c0176542c: 0}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_ab2c02ab751b9f42 *big.Int, __obf_d5f3440c0176542c int32) Decimal {
	return Decimal{__obf_ab2c02ab751b9f42: new(big.Int).Set(__obf_ab2c02ab751b9f42), __obf_d5f3440c0176542c: __obf_d5f3440c0176542c}
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
func NewFromBigRat(__obf_ab2c02ab751b9f42 *big.Rat, __obf_9e4f7aff89415e13 int32) Decimal {
	return Decimal{__obf_ab2c02ab751b9f42: new(big.Int).Set(__obf_ab2c02ab751b9f42.Num()), __obf_d5f3440c0176542c: 0}.DivRound(Decimal{__obf_ab2c02ab751b9f42: new(big.Int).Set(__obf_ab2c02ab751b9f42.Denom()), __obf_d5f3440c0176542c: 0}, __obf_9e4f7aff89415e13)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_ab2c02ab751b9f42 string) (Decimal, error) {
	__obf_dbef8dc5b9b01a79 := __obf_ab2c02ab751b9f42
	var __obf_e7754e83c555dd6a string
	var __obf_d5f3440c0176542c int64
	__obf_fb75bb33203e1202 := // Check if number is using scientific notation
		strings.IndexAny(__obf_ab2c02ab751b9f42, "Ee")
	if __obf_fb75bb33203e1202 != -1 {
		__obf_c7d02b17693b9530, __obf_239001ea6a8e4288 := strconv.ParseInt(__obf_ab2c02ab751b9f42[__obf_fb75bb33203e1202+1:], 10, 32)
		if __obf_239001ea6a8e4288 != nil {
			if __obf_8237856756c7a8ca, __obf_4c4c4c9db3c018b9 := __obf_239001ea6a8e4288.(*strconv.NumError); __obf_4c4c4c9db3c018b9 && __obf_8237856756c7a8ca.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_ab2c02ab751b9f42)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_ab2c02ab751b9f42)
		}
		__obf_ab2c02ab751b9f42 = __obf_ab2c02ab751b9f42[:__obf_fb75bb33203e1202]
		__obf_d5f3440c0176542c = __obf_c7d02b17693b9530
	}
	__obf_6a6c17d2e5f611af := -1
	__obf_de74b75498d53d7d := len(__obf_ab2c02ab751b9f42)
	for __obf_283995368e36d3d0 := 0; __obf_283995368e36d3d0 < __obf_de74b75498d53d7d; __obf_283995368e36d3d0++ {
		if __obf_ab2c02ab751b9f42[__obf_283995368e36d3d0] == '.' {
			if __obf_6a6c17d2e5f611af > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_ab2c02ab751b9f42)
			}
			__obf_6a6c17d2e5f611af = __obf_283995368e36d3d0
		}
	}

	if __obf_6a6c17d2e5f611af == -1 {
		__obf_e7754e83c555dd6a = // There is no decimal point, we can just parse the original string as
			// an int
			__obf_ab2c02ab751b9f42
	} else {
		if __obf_6a6c17d2e5f611af+1 < __obf_de74b75498d53d7d {
			__obf_e7754e83c555dd6a = __obf_ab2c02ab751b9f42[:__obf_6a6c17d2e5f611af] + __obf_ab2c02ab751b9f42[__obf_6a6c17d2e5f611af+1:]
		} else {
			__obf_e7754e83c555dd6a = __obf_ab2c02ab751b9f42[:__obf_6a6c17d2e5f611af]
		}
		__obf_c7d02b17693b9530 := -len(__obf_ab2c02ab751b9f42[__obf_6a6c17d2e5f611af+1:])
		__obf_d5f3440c0176542c += int64(__obf_c7d02b17693b9530)
	}

	var __obf_74f2531a3ddc5373 *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_e7754e83c555dd6a) <= 18 {
		__obf_35e65d688faf6d4b, __obf_239001ea6a8e4288 := strconv.ParseInt(__obf_e7754e83c555dd6a, 10, 64)
		if __obf_239001ea6a8e4288 != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_ab2c02ab751b9f42)
		}
		__obf_74f2531a3ddc5373 = big.NewInt(__obf_35e65d688faf6d4b)
	} else {
		__obf_74f2531a3ddc5373 = new(big.Int)
		_, __obf_4c4c4c9db3c018b9 := __obf_74f2531a3ddc5373.SetString(__obf_e7754e83c555dd6a, 10)
		if !__obf_4c4c4c9db3c018b9 {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_ab2c02ab751b9f42)
		}
	}

	if __obf_d5f3440c0176542c < math.MinInt32 || __obf_d5f3440c0176542c > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_dbef8dc5b9b01a79)
	}

	return Decimal{__obf_ab2c02ab751b9f42: __obf_74f2531a3ddc5373, __obf_d5f3440c0176542c: int32(__obf_d5f3440c0176542c)}, nil
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
func NewFromFormattedString(__obf_ab2c02ab751b9f42 string, __obf_fceaeb31231b8f8f *regexp.Regexp) (Decimal, error) {
	__obf_a6725d731f16653e := __obf_fceaeb31231b8f8f.ReplaceAllString(__obf_ab2c02ab751b9f42, "")
	__obf_5debc2a983e87a03, __obf_239001ea6a8e4288 := NewFromString(__obf_a6725d731f16653e)
	if __obf_239001ea6a8e4288 != nil {
		return Decimal{}, __obf_239001ea6a8e4288
	}
	return __obf_5debc2a983e87a03, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_ab2c02ab751b9f42 string) Decimal {
	__obf_5b220531a71314a4, __obf_239001ea6a8e4288 := NewFromString(__obf_ab2c02ab751b9f42)
	if __obf_239001ea6a8e4288 != nil {
		panic(__obf_239001ea6a8e4288)
	}
	return __obf_5b220531a71314a4
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
func NewFromFloat(__obf_ab2c02ab751b9f42 float64) Decimal {
	if __obf_ab2c02ab751b9f42 == 0 {
		return New(0, 0)
	}
	return __obf_c594f98a81b30522(__obf_ab2c02ab751b9f42, math.Float64bits(__obf_ab2c02ab751b9f42), &__obf_3d88be66e25681be)
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
func NewFromFloat32(__obf_ab2c02ab751b9f42 float32) Decimal {
	if __obf_ab2c02ab751b9f42 == 0 {
		return New(0, 0)
	}
	__obf_53ac9a541c64b86f := // XOR is workaround for https://github.com/golang/go/issues/26285
		math.Float32bits(__obf_ab2c02ab751b9f42) ^ 0x80808080
	return __obf_c594f98a81b30522(float64(__obf_ab2c02ab751b9f42), uint64(__obf_53ac9a541c64b86f)^0x80808080, &__obf_8dea1c634edf06b9)
}

func __obf_c594f98a81b30522(__obf_b487ac8d24688d3e float64, __obf_cbcc7239028a71ca uint64, __obf_24b406b733eb006b *__obf_ff127edd42fa495d) Decimal {
	if math.IsNaN(__obf_b487ac8d24688d3e) || math.IsInf(__obf_b487ac8d24688d3e, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_b487ac8d24688d3e))
	}
	__obf_d5f3440c0176542c := int(__obf_cbcc7239028a71ca>>__obf_24b406b733eb006b.__obf_c0ccdfc479ac2e96) & (1<<__obf_24b406b733eb006b.__obf_8ffbe185d6c7002e - 1)
	__obf_7ede9834c2850305 := __obf_cbcc7239028a71ca & (uint64(1)<<__obf_24b406b733eb006b.__obf_c0ccdfc479ac2e96 - 1)

	switch __obf_d5f3440c0176542c {
	case 0:
		__obf_d5f3440c0176542c++ // denormalized

	default:
		__obf_7ede9834c2850305 |= // add implicit top bit
			uint64(1) << __obf_24b406b733eb006b.__obf_c0ccdfc479ac2e96
	}
	__obf_d5f3440c0176542c += __obf_24b406b733eb006b.__obf_0570419913d32c75

	var __obf_5debc2a983e87a03 __obf_a5d0662d16479f1d
	__obf_5debc2a983e87a03.
		Assign(__obf_7ede9834c2850305)
	__obf_5debc2a983e87a03.
		Shift(__obf_d5f3440c0176542c - int(__obf_24b406b733eb006b.__obf_c0ccdfc479ac2e96))
	__obf_5debc2a983e87a03.__obf_5084f9ae1c148972 = __obf_cbcc7239028a71ca>>(__obf_24b406b733eb006b.__obf_8ffbe185d6c7002e+__obf_24b406b733eb006b.__obf_c0ccdfc479ac2e96) != 0
	__obf_1e54fb2a7fbae193(&__obf_5debc2a983e87a03,
		// If less than 19 digits, we can do calculation in an int64.
		__obf_7ede9834c2850305, __obf_d5f3440c0176542c, __obf_24b406b733eb006b)

	if __obf_5debc2a983e87a03.__obf_d5ebe437c13ed047 < 19 {
		__obf_ec10a27b4c6805fe := int64(0)
		__obf_dfa47acb5d50aedb := int64(1)
		for __obf_283995368e36d3d0 := __obf_5debc2a983e87a03.__obf_d5ebe437c13ed047 - 1; __obf_283995368e36d3d0 >= 0; __obf_283995368e36d3d0-- {
			__obf_ec10a27b4c6805fe += __obf_dfa47acb5d50aedb * int64(__obf_5debc2a983e87a03.__obf_5debc2a983e87a03[__obf_283995368e36d3d0]-'0')
			__obf_dfa47acb5d50aedb *= 10
		}
		if __obf_5debc2a983e87a03.__obf_5084f9ae1c148972 {
			__obf_ec10a27b4c6805fe *= -1
		}
		return Decimal{__obf_ab2c02ab751b9f42: big.NewInt(__obf_ec10a27b4c6805fe), __obf_d5f3440c0176542c: int32(__obf_5debc2a983e87a03.__obf_4e162ef443f77409) - int32(__obf_5debc2a983e87a03.__obf_d5ebe437c13ed047)}
	}
	__obf_74f2531a3ddc5373 := new(big.Int)
	__obf_74f2531a3ddc5373, __obf_4c4c4c9db3c018b9 := __obf_74f2531a3ddc5373.SetString(string(__obf_5debc2a983e87a03.__obf_5debc2a983e87a03[:__obf_5debc2a983e87a03.__obf_d5ebe437c13ed047]), 10)
	if __obf_4c4c4c9db3c018b9 {
		return Decimal{__obf_ab2c02ab751b9f42: __obf_74f2531a3ddc5373, __obf_d5f3440c0176542c: int32(__obf_5debc2a983e87a03.__obf_4e162ef443f77409) - int32(__obf_5debc2a983e87a03.__obf_d5ebe437c13ed047)}
	}

	return NewFromFloatWithExponent(__obf_b487ac8d24688d3e, int32(__obf_5debc2a983e87a03.

		// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
		// number of fractional digits.
		//
		// Example:
		//
		//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
		__obf_4e162ef443f77409)-int32(__obf_5debc2a983e87a03.__obf_d5ebe437c13ed047))
}

func NewFromFloatWithExponent(__obf_ab2c02ab751b9f42 float64, __obf_d5f3440c0176542c int32) Decimal {
	if math.IsNaN(__obf_ab2c02ab751b9f42) || math.IsInf(__obf_ab2c02ab751b9f42, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_ab2c02ab751b9f42))
	}
	__obf_cbcc7239028a71ca := math.Float64bits(__obf_ab2c02ab751b9f42)
	__obf_7ede9834c2850305 := __obf_cbcc7239028a71ca & (1<<52 - 1)
	__obf_e0fb9b0bfbf2d093 := int32((__obf_cbcc7239028a71ca >> 52) & (1<<11 - 1))
	__obf_f49c051f0678a1f6 := __obf_cbcc7239028a71ca >> 63

	if __obf_e0fb9b0bfbf2d093 == 0 {
		// specials
		if __obf_7ede9834c2850305 == 0 {
			return Decimal{}
		}
		__obf_e0fb9b0bfbf2d093++ // subnormal

	} else {
		__obf_7ede9834c2850305 |= // normal
			1 << 52
	}
	__obf_e0fb9b0bfbf2d093 -= 1023 + 52

	// normalizing base-2 values
	for __obf_7ede9834c2850305&1 == 0 {
		__obf_7ede9834c2850305 = __obf_7ede9834c2850305 >> 1
		__obf_e0fb9b0bfbf2d093++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_d5f3440c0176542c < 0 && __obf_d5f3440c0176542c < __obf_e0fb9b0bfbf2d093 {
		if __obf_e0fb9b0bfbf2d093 < 0 {
			__obf_d5f3440c0176542c = __obf_e0fb9b0bfbf2d093
		} else {
			__obf_d5f3440c0176542c = 0
		}
	}
	__obf_e0fb9b0bfbf2d093 -= // representing 10^M * 2^N as 5^M * 2^(M+N)
		__obf_d5f3440c0176542c
	__obf_f3adf77928d7e5c2 := big.NewInt(1)
	__obf_6b08695a00fc6ffd := big.NewInt(int64(__obf_7ede9834c2850305))

	// applying 5^M
	if __obf_d5f3440c0176542c > 0 {
		__obf_f3adf77928d7e5c2 = __obf_f3adf77928d7e5c2.SetInt64(int64(__obf_d5f3440c0176542c))
		__obf_f3adf77928d7e5c2 = __obf_f3adf77928d7e5c2.Exp(__obf_656e224a8e2fd8a8, __obf_f3adf77928d7e5c2, nil)
	} else if __obf_d5f3440c0176542c < 0 {
		__obf_f3adf77928d7e5c2 = __obf_f3adf77928d7e5c2.SetInt64(-int64(__obf_d5f3440c0176542c))
		__obf_f3adf77928d7e5c2 = __obf_f3adf77928d7e5c2.Exp(__obf_656e224a8e2fd8a8, __obf_f3adf77928d7e5c2, nil)
		__obf_6b08695a00fc6ffd = __obf_6b08695a00fc6ffd.Mul(__obf_6b08695a00fc6ffd, __obf_f3adf77928d7e5c2)
		__obf_f3adf77928d7e5c2 = __obf_f3adf77928d7e5c2.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_e0fb9b0bfbf2d093 > 0 {
		__obf_6b08695a00fc6ffd = __obf_6b08695a00fc6ffd.Lsh(__obf_6b08695a00fc6ffd, uint(__obf_e0fb9b0bfbf2d093))
	} else if __obf_e0fb9b0bfbf2d093 < 0 {
		__obf_f3adf77928d7e5c2 = __obf_f3adf77928d7e5c2.Lsh(__obf_f3adf77928d7e5c2, uint(-__obf_e0fb9b0bfbf2d093))
	}

	// rounding and downscaling
	if __obf_d5f3440c0176542c > 0 || __obf_e0fb9b0bfbf2d093 < 0 {
		__obf_8816681d9b7e9f19 := new(big.Int).Rsh(__obf_f3adf77928d7e5c2, 1)
		__obf_6b08695a00fc6ffd = __obf_6b08695a00fc6ffd.Add(__obf_6b08695a00fc6ffd, __obf_8816681d9b7e9f19)
		__obf_6b08695a00fc6ffd = __obf_6b08695a00fc6ffd.Quo(__obf_6b08695a00fc6ffd, __obf_f3adf77928d7e5c2)
	}

	if __obf_f49c051f0678a1f6 == 1 {
		__obf_6b08695a00fc6ffd = __obf_6b08695a00fc6ffd.Neg(__obf_6b08695a00fc6ffd)
	}

	return Decimal{__obf_ab2c02ab751b9f42: __obf_6b08695a00fc6ffd, __obf_d5f3440c0176542c: __obf_d5f3440c0176542c}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_5debc2a983e87a03 Decimal) Copy() Decimal {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()
	return Decimal{__obf_ab2c02ab751b9f42: new(big.Int).Set(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42), __obf_d5f3440c0176542c: __obf_5debc2a983e87a03.

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
		__obf_d5f3440c0176542c,
	}
}

func (__obf_5debc2a983e87a03 Decimal) __obf_c2312e5df65a128a(__obf_d5f3440c0176542c int32) Decimal {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()

	if __obf_5debc2a983e87a03.__obf_d5f3440c0176542c == __obf_d5f3440c0176542c {
		return Decimal{
			new(big.Int).Set(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42), __obf_5debc2a983e87a03.

				// NOTE(vadim): must convert exps to float64 before - to prevent overflow
				__obf_d5f3440c0176542c,
		}
	}
	__obf_29444ac194cf9e44 := math.Abs(float64(__obf_d5f3440c0176542c) - float64(__obf_5debc2a983e87a03.__obf_d5f3440c0176542c))
	__obf_ab2c02ab751b9f42 := new(big.Int).Set(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42)
	__obf_895b041c888f0508 := new(big.Int).Exp(__obf_8f06f98f079e4122, big.NewInt(int64(__obf_29444ac194cf9e44)), nil)
	if __obf_d5f3440c0176542c > __obf_5debc2a983e87a03.__obf_d5f3440c0176542c {
		__obf_ab2c02ab751b9f42 = __obf_ab2c02ab751b9f42.Quo(__obf_ab2c02ab751b9f42, __obf_895b041c888f0508)
	} else if __obf_d5f3440c0176542c < __obf_5debc2a983e87a03.__obf_d5f3440c0176542c {
		__obf_ab2c02ab751b9f42 = __obf_ab2c02ab751b9f42.Mul(__obf_ab2c02ab751b9f42, __obf_895b041c888f0508)
	}

	return Decimal{__obf_ab2c02ab751b9f42: __obf_ab2c02ab751b9f42, __obf_d5f3440c0176542c: __obf_d5f3440c0176542c}
}

// Abs returns the absolute value of the decimal.
func (__obf_5debc2a983e87a03 Decimal) Abs() Decimal {
	if !__obf_5debc2a983e87a03.IsNegative() {
		return __obf_5debc2a983e87a03
	}
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()
	__obf_df6031c97b595c47 := new(big.Int).Abs(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42)
	return Decimal{__obf_ab2c02ab751b9f42: __obf_df6031c97b595c47, __obf_d5f3440c0176542c: __obf_5debc2a983e87a03.

		// Add returns d + d2.
		__obf_d5f3440c0176542c,
	}
}

func (__obf_5debc2a983e87a03 Decimal) Add(__obf_e06b6478467c95fe Decimal) Decimal {
	__obf_85728b8dfdaeed90, __obf_c83c6d3225e9675c := RescalePair(__obf_5debc2a983e87a03, __obf_e06b6478467c95fe)
	__obf_8bd58a39cba2a46b := new(big.Int).Add(__obf_85728b8dfdaeed90.__obf_ab2c02ab751b9f42, __obf_c83c6d3225e9675c.__obf_ab2c02ab751b9f42)
	return Decimal{__obf_ab2c02ab751b9f42: __obf_8bd58a39cba2a46b, __obf_d5f3440c0176542c: __obf_85728b8dfdaeed90.

		// Sub returns d - d2.
		__obf_d5f3440c0176542c,
	}
}

func (__obf_5debc2a983e87a03 Decimal) Sub(__obf_e06b6478467c95fe Decimal) Decimal {
	__obf_85728b8dfdaeed90, __obf_c83c6d3225e9675c := RescalePair(__obf_5debc2a983e87a03, __obf_e06b6478467c95fe)
	__obf_8bd58a39cba2a46b := new(big.Int).Sub(__obf_85728b8dfdaeed90.__obf_ab2c02ab751b9f42, __obf_c83c6d3225e9675c.__obf_ab2c02ab751b9f42)
	return Decimal{__obf_ab2c02ab751b9f42: __obf_8bd58a39cba2a46b, __obf_d5f3440c0176542c: __obf_85728b8dfdaeed90.

		// Neg returns -d.
		__obf_d5f3440c0176542c,
	}
}

func (__obf_5debc2a983e87a03 Decimal) Neg() Decimal {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()
	__obf_b487ac8d24688d3e := new(big.Int).Neg(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42)
	return Decimal{__obf_ab2c02ab751b9f42: __obf_b487ac8d24688d3e, __obf_d5f3440c0176542c: __obf_5debc2a983e87a03.

		// Mul returns d * d2.
		__obf_d5f3440c0176542c,
	}
}

func (__obf_5debc2a983e87a03 Decimal) Mul(__obf_e06b6478467c95fe Decimal) Decimal {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()
	__obf_e06b6478467c95fe.__obf_23af5656aedaa490()
	__obf_224113570f282cff := int64(__obf_5debc2a983e87a03.__obf_d5f3440c0176542c) + int64(__obf_e06b6478467c95fe.__obf_d5f3440c0176542c)
	if __obf_224113570f282cff > math.MaxInt32 || __obf_224113570f282cff < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_224113570f282cff))
	}
	__obf_8bd58a39cba2a46b := new(big.Int).Mul(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42, __obf_e06b6478467c95fe.__obf_ab2c02ab751b9f42)
	return Decimal{__obf_ab2c02ab751b9f42: __obf_8bd58a39cba2a46b, __obf_d5f3440c0176542c: int32(__obf_224113570f282cff)}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_5debc2a983e87a03 Decimal) Shift(__obf_0646db04f8a6ec3c int32) Decimal {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()
	return Decimal{__obf_ab2c02ab751b9f42: new(big.Int).Set(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42), __obf_d5f3440c0176542c: __obf_5debc2a983e87a03.__obf_d5f3440c0176542c + __obf_0646db04f8a6ec3c}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_5debc2a983e87a03 Decimal) Div(__obf_e06b6478467c95fe Decimal) Decimal {
	return __obf_5debc2a983e87a03.DivRound(__obf_e06b6478467c95fe, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_5debc2a983e87a03 Decimal) QuoRem(__obf_e06b6478467c95fe Decimal, __obf_9e4f7aff89415e13 int32) (Decimal, Decimal) {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()
	__obf_e06b6478467c95fe.__obf_23af5656aedaa490()
	if __obf_e06b6478467c95fe.__obf_ab2c02ab751b9f42.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_6bd4ec679d7f6a29 := -__obf_9e4f7aff89415e13
	__obf_8237856756c7a8ca := int64(__obf_5debc2a983e87a03.__obf_d5f3440c0176542c) - int64(__obf_e06b6478467c95fe.__obf_d5f3440c0176542c) - int64(__obf_6bd4ec679d7f6a29)
	if __obf_8237856756c7a8ca > math.MaxInt32 || __obf_8237856756c7a8ca < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_ab1e10cf0c8efe44, __obf_353cfaa2df394f47,

		// d = a 10^ea
		// d2 = b 10^eb
		__obf_1ad08a1c7538e89d big.Int
	var __obf_b3df56c27155d601 int32

	if __obf_8237856756c7a8ca < 0 {
		__obf_ab1e10cf0c8efe44 = *__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42
		__obf_1ad08a1c7538e89d.
			SetInt64(-__obf_8237856756c7a8ca)
		__obf_353cfaa2df394f47.
			Exp(__obf_8f06f98f079e4122, &__obf_1ad08a1c7538e89d, nil)
		__obf_353cfaa2df394f47.
			Mul(__obf_e06b6478467c95fe.__obf_ab2c02ab751b9f42, &__obf_353cfaa2df394f47)
		__obf_b3df56c27155d601 = __obf_5debc2a983e87a03.
			// now aa = a
			//     bb = b 10^(scale + eb - ea)
			__obf_d5f3440c0176542c

	} else {
		__obf_1ad08a1c7538e89d.
			SetInt64(__obf_8237856756c7a8ca)
		__obf_ab1e10cf0c8efe44.
			Exp(__obf_8f06f98f079e4122, &__obf_1ad08a1c7538e89d, nil)
		__obf_ab1e10cf0c8efe44.
			Mul(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42, &__obf_ab1e10cf0c8efe44)
		__obf_353cfaa2df394f47 = *__obf_e06b6478467c95fe.__obf_ab2c02ab751b9f42

		// now aa = a ^ (ea - eb - scale)
		//     bb = b
		__obf_b3df56c27155d601 = __obf_6bd4ec679d7f6a29 + __obf_e06b6478467c95fe.__obf_d5f3440c0176542c

	}
	var __obf_8316632570bb1426, __obf_1a925cec70c560aa big.Int
	__obf_8316632570bb1426.
		QuoRem(&__obf_ab1e10cf0c8efe44, &__obf_353cfaa2df394f47, &__obf_1a925cec70c560aa)
	__obf_8a9609e95948f792 := Decimal{__obf_ab2c02ab751b9f42: &__obf_8316632570bb1426, __obf_d5f3440c0176542c: __obf_6bd4ec679d7f6a29}
	__obf_acfd567e48cd7dab := Decimal{__obf_ab2c02ab751b9f42: &__obf_1a925cec70c560aa, __obf_d5f3440c0176542c: __obf_b3df56c27155d601}
	return __obf_8a9609e95948f792,

		// DivRound divides and rounds to a given precision
		// i.e. to an integer multiple of 10^(-precision)
		//
		//	for a positive quotient digit 5 is rounded up, away from 0
		//	if the quotient is negative then digit 5 is rounded down, away from 0
		//
		// Note that precision<0 is allowed as input.
		__obf_acfd567e48cd7dab
}

func (__obf_5debc2a983e87a03 Decimal) DivRound(__obf_e06b6478467c95fe Decimal, __obf_9e4f7aff89415e13 int32) Decimal {
	__obf_8316632570bb1426,
		// QuoRem already checks initialization
		__obf_1a925cec70c560aa := __obf_5debc2a983e87a03.QuoRem(__obf_e06b6478467c95fe,
		// the actual rounding decision is based on comparing r*10^precision and d2/2
		// instead compare 2 r 10 ^precision and d2
		__obf_9e4f7aff89415e13)

	var __obf_b813fad7a5620fef big.Int
	__obf_b813fad7a5620fef.
		Abs(__obf_1a925cec70c560aa.__obf_ab2c02ab751b9f42)
	__obf_b813fad7a5620fef.
		Lsh(&__obf_b813fad7a5620fef, 1)
	__obf_1454a1dda4e7d236 := // now rv2 = abs(r.value) * 2
		Decimal{__obf_ab2c02ab751b9f42: &__obf_b813fad7a5620fef, __obf_d5f3440c0176542c: __obf_1a925cec70c560aa.
			// r2 is now 2 * r * 10 ^ precision
			__obf_d5f3440c0176542c + __obf_9e4f7aff89415e13}

	var __obf_307e3ae1887fcdd1 = __obf_1454a1dda4e7d236.Cmp(__obf_e06b6478467c95fe.Abs())

	if __obf_307e3ae1887fcdd1 < 0 {
		return __obf_8316632570bb1426
	}

	if __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.Sign()*__obf_e06b6478467c95fe.__obf_ab2c02ab751b9f42.Sign() < 0 {
		return __obf_8316632570bb1426.Sub(New(1, -__obf_9e4f7aff89415e13))
	}

	return __obf_8316632570bb1426.Add(New(1, -__obf_9e4f7aff89415e13))
}

// Mod returns d % d2.
func (__obf_5debc2a983e87a03 Decimal) Mod(__obf_e06b6478467c95fe Decimal) Decimal {
	_, __obf_1a925cec70c560aa := __obf_5debc2a983e87a03.QuoRem(__obf_e06b6478467c95fe, 0)
	return __obf_1a925cec70c560aa
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
func (__obf_5debc2a983e87a03 Decimal) Pow(__obf_e06b6478467c95fe Decimal) Decimal {
	__obf_5e3b817b9bb1b94c := __obf_5debc2a983e87a03.Sign()
	__obf_8f7c8bfa26837118 := __obf_e06b6478467c95fe.Sign()

	if __obf_5e3b817b9bb1b94c == 0 {
		if __obf_8f7c8bfa26837118 == 0 {
			return Decimal{}
		}
		if __obf_8f7c8bfa26837118 == 1 {
			return Decimal{__obf_304be536ded8753b, 0}
		}
		if __obf_8f7c8bfa26837118 == -1 {
			return Decimal{}
		}
	}

	if __obf_8f7c8bfa26837118 == 0 {
		return Decimal{__obf_0ca31adf9d5bc652, 0}
	}
	__obf_b75a9f95ea711bee := // TODO: optimize extraction of fractional part
		Decimal{__obf_0ca31adf9d5bc652, 0}
	__obf_b4081d143e4c03ba, __obf_c28e2082dfdc56dd := __obf_e06b6478467c95fe.QuoRem(__obf_b75a9f95ea711bee, 0)

	if __obf_5e3b817b9bb1b94c == -1 && !__obf_c28e2082dfdc56dd.IsZero() {
		return Decimal{}
	}
	__obf_7a51617fd0db4516, _ := __obf_5debc2a983e87a03.PowBigInt(__obf_b4081d143e4c03ba.

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_ab2c02ab751b9f42)

	if __obf_c28e2082dfdc56dd.__obf_ab2c02ab751b9f42.Sign() == 0 {
		return __obf_7a51617fd0db4516
	}
	__obf_0c8a1d19ff1ecbf4 := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_5debc2a983e87a03.NumDigits()
	__obf_f3cf95ebe8aaf50f := __obf_e06b6478467c95fe.NumDigits()
	__obf_9e4f7aff89415e13 := __obf_0c8a1d19ff1ecbf4

	if __obf_f3cf95ebe8aaf50f > __obf_9e4f7aff89415e13 {
		__obf_9e4f7aff89415e13 += __obf_f3cf95ebe8aaf50f
	}
	__obf_9e4f7aff89415e13 += 6
	__obf_6c1045358c680cf4,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_239001ea6a8e4288 := __obf_5debc2a983e87a03.Abs().Ln(-__obf_5debc2a983e87a03.__obf_d5f3440c0176542c + int32(__obf_9e4f7aff89415e13))
	if __obf_239001ea6a8e4288 != nil {
		return Decimal{}
	}
	__obf_6c1045358c680cf4 = __obf_6c1045358c680cf4.Mul(__obf_c28e2082dfdc56dd)
	__obf_6c1045358c680cf4, __obf_239001ea6a8e4288 = __obf_6c1045358c680cf4.ExpTaylor(-__obf_5debc2a983e87a03.__obf_d5f3440c0176542c + int32(__obf_9e4f7aff89415e13))
	if __obf_239001ea6a8e4288 != nil {
		return Decimal{}
	}
	__obf_41d953c797b1a497 := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_7a51617fd0db4516.Mul(__obf_6c1045358c680cf4)

	return __obf_41d953c797b1a497
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
func (__obf_5debc2a983e87a03 Decimal) PowWithPrecision(__obf_e06b6478467c95fe Decimal, __obf_9e4f7aff89415e13 int32) (Decimal, error) {
	__obf_5e3b817b9bb1b94c := __obf_5debc2a983e87a03.Sign()
	__obf_8f7c8bfa26837118 := __obf_e06b6478467c95fe.Sign()

	if __obf_5e3b817b9bb1b94c == 0 {
		if __obf_8f7c8bfa26837118 == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_8f7c8bfa26837118 == 1 {
			return Decimal{__obf_304be536ded8753b, 0}, nil
		}
		if __obf_8f7c8bfa26837118 == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_8f7c8bfa26837118 == 0 {
		return Decimal{__obf_0ca31adf9d5bc652, 0}, nil
	}
	__obf_b75a9f95ea711bee := // TODO: optimize extraction of fractional part
		Decimal{__obf_0ca31adf9d5bc652, 0}
	__obf_b4081d143e4c03ba, __obf_c28e2082dfdc56dd := __obf_e06b6478467c95fe.QuoRem(__obf_b75a9f95ea711bee, 0)

	if __obf_5e3b817b9bb1b94c == -1 && !__obf_c28e2082dfdc56dd.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}
	__obf_7a51617fd0db4516, _ := __obf_5debc2a983e87a03.__obf_2d81bb673cbad92f(__obf_b4081d143e4c03ba.__obf_ab2c02ab751b9f42,

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_9e4f7aff89415e13)

	if __obf_c28e2082dfdc56dd.__obf_ab2c02ab751b9f42.Sign() == 0 {
		return __obf_7a51617fd0db4516, nil
	}
	__obf_0c8a1d19ff1ecbf4 := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_5debc2a983e87a03.NumDigits()
	__obf_f3cf95ebe8aaf50f := __obf_e06b6478467c95fe.NumDigits()

	if int32(__obf_0c8a1d19ff1ecbf4) > __obf_9e4f7aff89415e13 {
		__obf_9e4f7aff89415e13 = int32(__obf_0c8a1d19ff1ecbf4)
	}
	if int32(__obf_f3cf95ebe8aaf50f) > __obf_9e4f7aff89415e13 {
		__obf_9e4f7aff89415e13 += int32(__obf_f3cf95ebe8aaf50f)
	}
	__obf_9e4f7aff89415e13 += // increase precision by 10 to compensate for errors in further calculations
		10
	__obf_6c1045358c680cf4,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_239001ea6a8e4288 := __obf_5debc2a983e87a03.Abs().Ln(__obf_9e4f7aff89415e13)
	if __obf_239001ea6a8e4288 != nil {
		return Decimal{}, __obf_239001ea6a8e4288
	}
	__obf_6c1045358c680cf4 = __obf_6c1045358c680cf4.Mul(__obf_c28e2082dfdc56dd)
	__obf_6c1045358c680cf4, __obf_239001ea6a8e4288 = __obf_6c1045358c680cf4.ExpTaylor(__obf_9e4f7aff89415e13)
	if __obf_239001ea6a8e4288 != nil {
		return Decimal{}, __obf_239001ea6a8e4288
	}
	__obf_41d953c797b1a497 := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_7a51617fd0db4516.Mul(__obf_6c1045358c680cf4)

	return __obf_41d953c797b1a497, nil
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
func (__obf_5debc2a983e87a03 Decimal) PowInt32(__obf_d5f3440c0176542c int32) (Decimal, error) {
	if __obf_5debc2a983e87a03.IsZero() && __obf_d5f3440c0176542c == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_7fab1e64bf4323b8 := __obf_d5f3440c0176542c < 0
	__obf_d5f3440c0176542c = __obf_b8bca49ac727c366(__obf_d5f3440c0176542c)
	__obf_5618ed758d2ff827, __obf_e531afa5ffc5fea3 := __obf_5debc2a983e87a03, New(1, 0)

	for __obf_d5f3440c0176542c > 0 {
		if __obf_d5f3440c0176542c%2 == 1 {
			__obf_e531afa5ffc5fea3 = __obf_e531afa5ffc5fea3.Mul(__obf_5618ed758d2ff827)
		}
		__obf_d5f3440c0176542c /= 2

		if __obf_d5f3440c0176542c > 0 {
			__obf_5618ed758d2ff827 = __obf_5618ed758d2ff827.Mul(__obf_5618ed758d2ff827)
		}
	}

	if __obf_7fab1e64bf4323b8 {
		return New(1, 0).DivRound(__obf_e531afa5ffc5fea3, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_e531afa5ffc5fea3, nil
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
func (__obf_5debc2a983e87a03 Decimal) PowBigInt(__obf_d5f3440c0176542c *big.Int) (Decimal, error) {
	return __obf_5debc2a983e87a03.__obf_2d81bb673cbad92f(__obf_d5f3440c0176542c, int32(PowPrecisionNegativeExponent))
}

func (__obf_5debc2a983e87a03 Decimal) __obf_2d81bb673cbad92f(__obf_d5f3440c0176542c *big.Int, __obf_9e4f7aff89415e13 int32) (Decimal, error) {
	if __obf_5debc2a983e87a03.IsZero() && __obf_d5f3440c0176542c.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_367d7ac98383333e := new(big.Int).Set(__obf_d5f3440c0176542c)
	__obf_7fab1e64bf4323b8 := __obf_d5f3440c0176542c.Sign() < 0

	if __obf_7fab1e64bf4323b8 {
		__obf_367d7ac98383333e.
			Abs(__obf_367d7ac98383333e)
	}
	__obf_5618ed758d2ff827, __obf_e531afa5ffc5fea3 := __obf_5debc2a983e87a03, New(1, 0)

	for __obf_367d7ac98383333e.Sign() > 0 {
		if __obf_367d7ac98383333e.Bit(0) == 1 {
			__obf_e531afa5ffc5fea3 = __obf_e531afa5ffc5fea3.Mul(__obf_5618ed758d2ff827)
		}
		__obf_367d7ac98383333e.
			Rsh(__obf_367d7ac98383333e, 1)

		if __obf_367d7ac98383333e.Sign() > 0 {
			__obf_5618ed758d2ff827 = __obf_5618ed758d2ff827.Mul(__obf_5618ed758d2ff827)
		}
	}

	if __obf_7fab1e64bf4323b8 {
		return New(1, 0).DivRound(__obf_e531afa5ffc5fea3, __obf_9e4f7aff89415e13), nil
	}

	return __obf_e531afa5ffc5fea3, nil
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
func (__obf_5debc2a983e87a03 Decimal) ExpHullAbrham(__obf_30f65108f97c2645 uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_5debc2a983e87a03.IsZero() {
		return Decimal{__obf_0ca31adf9d5bc652, 0}, nil
	}
	__obf_2ecb8e6711382f9c := __obf_30f65108f97c2645

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_42f0cd6ce08f2588 := __obf_5debc2a983e87a03.Abs().InexactFloat64()
	if __obf_db58a136143abcae := __obf_42f0cd6ce08f2588 / 23; __obf_db58a136143abcae > float64(__obf_2ecb8e6711382f9c) && __obf_db58a136143abcae < float64(ExpMaxIterations) {
		__obf_2ecb8e6711382f9c = uint32(math.Ceil(__obf_db58a136143abcae))
	}
	__obf_49257a1790243e44 := // fail if abs(d) beyond an over/underflow threshold
		New(23*int64(__obf_2ecb8e6711382f9c), 0)
	if __obf_5debc2a983e87a03.Abs().Cmp(__obf_49257a1790243e44) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}
	__obf_6c4e40e84fb533b8 := // Return 1 if abs(d) small enough; this also avoids later over/underflow
		New(9, -int32(__obf_2ecb8e6711382f9c)-1)
	if __obf_5debc2a983e87a03.Abs().Cmp(__obf_6c4e40e84fb533b8) <= 0 {
		return Decimal{__obf_0ca31adf9d5bc652, __obf_5debc2a983e87a03.

			// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
			__obf_d5f3440c0176542c}, nil
	}
	__obf_8133ee5c5b86aacf := __obf_5debc2a983e87a03.__obf_d5f3440c0176542c + int32(__obf_5debc2a983e87a03.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_8133ee5c5b86aacf < 0 {
		__obf_8133ee5c5b86aacf = 0
	}
	__obf_c0abc1dc124d3675 := New(1, __obf_8133ee5c5b86aacf)
	__obf_1a925cec70c560aa := // reduction factor
		Decimal{new(big.Int).Set(__obf_5debc2a983e87a03. // reduced argument
									__obf_ab2c02ab751b9f42), __obf_5debc2a983e87a03.__obf_d5f3440c0176542c - __obf_8133ee5c5b86aacf}
	__obf_6b72c778919ad94b := int32(__obf_2ecb8e6711382f9c) + __obf_8133ee5c5b86aacf + 2
	__obf_2b648debb6648663 := // precision for calculating the sum

		// Determine n, the number of therms for calculating sum
		// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
		// for solving appropriate equation, along with directed
		// roundings and simple rational bound for log10(p/abs(r))
		__obf_1a925cec70c560aa.Abs().InexactFloat64()
	__obf_8fd318a777b5b46d := float64(__obf_6b72c778919ad94b)
	__obf_a1435455e84ea966 := math.Ceil((1.453*__obf_8fd318a777b5b46d - 1.182) / math.Log10(__obf_8fd318a777b5b46d/__obf_2b648debb6648663))
	if __obf_a1435455e84ea966 > float64(ExpMaxIterations) || math.IsNaN(__obf_a1435455e84ea966) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_5618ed758d2ff827 := int64(__obf_a1435455e84ea966)
	__obf_ec10a27b4c6805fe := New(0, 0)
	__obf_272fc79b42d8fb7b := New(1, 0)
	__obf_b75a9f95ea711bee := New(1, 0)
	for __obf_283995368e36d3d0 := __obf_5618ed758d2ff827 - 1; __obf_283995368e36d3d0 > 0; __obf_283995368e36d3d0-- {
		__obf_ec10a27b4c6805fe.__obf_ab2c02ab751b9f42.
			SetInt64(__obf_283995368e36d3d0)
		__obf_272fc79b42d8fb7b = __obf_272fc79b42d8fb7b.Mul(__obf_1a925cec70c560aa.DivRound(__obf_ec10a27b4c6805fe, __obf_6b72c778919ad94b))
		__obf_272fc79b42d8fb7b = __obf_272fc79b42d8fb7b.Add(__obf_b75a9f95ea711bee)
	}
	__obf_197e6eb327366326 := __obf_c0abc1dc124d3675.IntPart()
	__obf_41d953c797b1a497 := New(1, 0)
	for __obf_283995368e36d3d0 := __obf_197e6eb327366326; __obf_283995368e36d3d0 > 0; __obf_283995368e36d3d0-- {
		__obf_41d953c797b1a497 = __obf_41d953c797b1a497.Mul(__obf_272fc79b42d8fb7b)
	}
	__obf_41d96c267929605b := int32(__obf_41d953c797b1a497.NumDigits())

	var __obf_954ecbc3af32171c int32
	if __obf_41d96c267929605b > __obf_b8bca49ac727c366(__obf_41d953c797b1a497.__obf_d5f3440c0176542c) {
		__obf_954ecbc3af32171c = int32(__obf_2ecb8e6711382f9c) - __obf_41d96c267929605b - __obf_41d953c797b1a497.__obf_d5f3440c0176542c
	} else {
		__obf_954ecbc3af32171c = int32(__obf_2ecb8e6711382f9c)
	}
	__obf_41d953c797b1a497 = __obf_41d953c797b1a497.Round(__obf_954ecbc3af32171c)

	return __obf_41d953c797b1a497, nil
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
func (__obf_5debc2a983e87a03 Decimal) ExpTaylor(__obf_9e4f7aff89415e13 int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_5debc2a983e87a03.IsZero() {
		return Decimal{__obf_0ca31adf9d5bc652, 0}.Round(__obf_9e4f7aff89415e13), nil
	}

	var __obf_b4bbd1adfe42f527 Decimal
	var __obf_7a157af777137562 int32
	if __obf_9e4f7aff89415e13 < 0 {
		__obf_b4bbd1adfe42f527 = New(1, -1)
		__obf_7a157af777137562 = 8
	} else {
		__obf_b4bbd1adfe42f527 = New(1, -__obf_9e4f7aff89415e13-1)
		__obf_7a157af777137562 = __obf_9e4f7aff89415e13 + 1
	}
	__obf_03824ef84782b144 := __obf_5debc2a983e87a03.Abs()
	__obf_a3e59393b46f5d7c := __obf_5debc2a983e87a03.Abs()
	__obf_5d3bc4baa190a804 := New(1, 0)
	__obf_e531afa5ffc5fea3 := New(1, 0)

	for __obf_283995368e36d3d0 := int64(1); ; {
		__obf_6a69ce2371844749 := __obf_a3e59393b46f5d7c.DivRound(__obf_5d3bc4baa190a804, __obf_7a157af777137562)
		__obf_e531afa5ffc5fea3 = __obf_e531afa5ffc5fea3.Add(__obf_6a69ce2371844749)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_6a69ce2371844749.Cmp(__obf_b4bbd1adfe42f527) < 0 {
			break
		}
		__obf_a3e59393b46f5d7c = __obf_a3e59393b46f5d7c.Mul(__obf_03824ef84782b144)
		__obf_283995368e36d3d0++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_edd8d74bea9c2de4) >= int(__obf_283995368e36d3d0) && !__obf_edd8d74bea9c2de4[__obf_283995368e36d3d0-1].IsZero() {
			__obf_5d3bc4baa190a804 = __obf_edd8d74bea9c2de4[__obf_283995368e36d3d0-1]
		} else {
			__obf_5d3bc4baa190a804 = // To avoid any race conditions, firstly the zero value is appended to a slice to create
				// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
				// factorial using the index notation.
				__obf_edd8d74bea9c2de4[__obf_283995368e36d3d0-2].Mul(New(__obf_283995368e36d3d0, 0))
			__obf_edd8d74bea9c2de4 = append(__obf_edd8d74bea9c2de4, Zero)
			__obf_edd8d74bea9c2de4[__obf_283995368e36d3d0-1] = __obf_5d3bc4baa190a804
		}
	}

	if __obf_5debc2a983e87a03.Sign() < 0 {
		__obf_e531afa5ffc5fea3 = New(1, 0).DivRound(__obf_e531afa5ffc5fea3, __obf_9e4f7aff89415e13+1)
	}
	__obf_e531afa5ffc5fea3 = __obf_e531afa5ffc5fea3.Round(__obf_9e4f7aff89415e13)
	return __obf_e531afa5ffc5fea3, nil
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
func (__obf_5debc2a983e87a03 Decimal) Ln(__obf_9e4f7aff89415e13 int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_5debc2a983e87a03.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_5debc2a983e87a03.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}
	__obf_0c3df84cc39e5815 := __obf_9e4f7aff89415e13 + 2
	__obf_25508e8597181510 := __obf_5debc2a983e87a03.Copy()

	var __obf_1bdc4788b863aed0, __obf_4647a81d1a2c2b1f, __obf_1658299eb7ff54fa, __obf_62b94eebda4a31b9, __obf_a35a5e9e2799b84b Decimal
	__obf_1bdc4788b863aed0 = __obf_25508e8597181510.Sub(Decimal{__obf_0ca31adf9d5bc652, 0})
	__obf_4647a81d1a2c2b1f = Decimal{__obf_0ca31adf9d5bc652, -1}
	__obf_5d8d5f2e6cfbe20f := // for decimal in range [0.9, 1.1] where ln(d) is close to 0
		false

	if __obf_1bdc4788b863aed0.Abs().Cmp(__obf_4647a81d1a2c2b1f) <= 0 {
		__obf_5d8d5f2e6cfbe20f = true
	} else {
		__obf_aebfd24ad70f18ed := // reduce input decimal to range [0.1, 1)
			int32(__obf_25508e8597181510.NumDigits()) + __obf_25508e8597181510.__obf_d5f3440c0176542c

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_25508e8597181510.__obf_d5f3440c0176542c -= __obf_aebfd24ad70f18ed
		__obf_49222726fcf2a30d := __obf_49222726fcf2a30d.__obf_a7135e6903b7ead2(__obf_0c3df84cc39e5815)
		__obf_a35a5e9e2799b84b = NewFromInt32(__obf_aebfd24ad70f18ed)
		__obf_a35a5e9e2799b84b = __obf_a35a5e9e2799b84b.Mul(__obf_49222726fcf2a30d)
		__obf_1bdc4788b863aed0 = __obf_25508e8597181510.Sub(Decimal{__obf_0ca31adf9d5bc652, 0})

		if __obf_1bdc4788b863aed0.Abs().Cmp(__obf_4647a81d1a2c2b1f) <= 0 {
			__obf_5d8d5f2e6cfbe20f = true
		} else {
			__obf_5cc753e3cc7b3545 := // initial estimate using floats
				__obf_25508e8597181510.InexactFloat64()
			__obf_1bdc4788b863aed0 = NewFromFloat(math.Log(__obf_5cc753e3cc7b3545))
		}
	}
	__obf_b4bbd1adfe42f527 := Decimal{__obf_0ca31adf9d5bc652, -__obf_0c3df84cc39e5815}

	if __obf_5d8d5f2e6cfbe20f {
		__obf_1658299eb7ff54fa = // Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
			// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			// until the difference between current and next term is smaller than epsilon.
			// Coverage quite fast for decimals close to 1.0

			// z + 2
			__obf_1bdc4788b863aed0.Add(Decimal{__obf_51443023c3dfc1b9, 0})
		__obf_4647a81d1a2c2b1f = // z / (z + 2)
			__obf_1bdc4788b863aed0.DivRound(__obf_1658299eb7ff54fa, __obf_0c3df84cc39e5815)
		__obf_1bdc4788b863aed0 = // 2 * (z / (z + 2))
			__obf_4647a81d1a2c2b1f.Add(__obf_4647a81d1a2c2b1f)
		__obf_1658299eb7ff54fa = __obf_1bdc4788b863aed0.Copy()

		for __obf_5618ed758d2ff827 := 1; ; __obf_5618ed758d2ff827++ {
			__obf_1658299eb7ff54fa = // 2 * (z / (z+2))^(2n+1)
				__obf_1658299eb7ff54fa.Mul(__obf_4647a81d1a2c2b1f).Mul(__obf_4647a81d1a2c2b1f)
			__obf_62b94eebda4a31b9 = // 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
				NewFromInt(int64(2*__obf_5618ed758d2ff827 + 1))
			__obf_62b94eebda4a31b9 = __obf_1658299eb7ff54fa.DivRound(__obf_62b94eebda4a31b9, __obf_0c3df84cc39e5815)
			__obf_1bdc4788b863aed0 = // comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
				__obf_1bdc4788b863aed0.Add(__obf_62b94eebda4a31b9)

			if __obf_62b94eebda4a31b9.Abs().Cmp(__obf_b4bbd1adfe42f527) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_e245a155f462b39c Decimal
		__obf_1037074597e427a1 := __obf_0c3df84cc39e5815*2 + 10

		for __obf_283995368e36d3d0 := int32(0); __obf_283995368e36d3d0 < __obf_1037074597e427a1;
		// exp(a_n)
		__obf_283995368e36d3d0++ {
			__obf_4647a81d1a2c2b1f, _ = __obf_1bdc4788b863aed0.ExpTaylor(__obf_0c3df84cc39e5815)
			__obf_1658299eb7ff54fa = // exp(a_n) - z
				__obf_4647a81d1a2c2b1f.Sub(__obf_25508e8597181510)
			__obf_1658299eb7ff54fa = // 2 * (exp(a_n) - z)
				__obf_1658299eb7ff54fa.Add(__obf_1658299eb7ff54fa)
			__obf_62b94eebda4a31b9 = // exp(a_n) + z
				__obf_4647a81d1a2c2b1f.Add(__obf_25508e8597181510)
			__obf_4647a81d1a2c2b1f = // 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_1658299eb7ff54fa.DivRound(__obf_62b94eebda4a31b9, __obf_0c3df84cc39e5815)
			__obf_1bdc4788b863aed0 = // comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_1bdc4788b863aed0.Sub(__obf_4647a81d1a2c2b1f)

			if __obf_e245a155f462b39c.Add(__obf_4647a81d1a2c2b1f).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_4647a81d1a2c2b1f.Abs().Cmp(__obf_b4bbd1adfe42f527) <= 0 {
				break
			}
			__obf_e245a155f462b39c = __obf_4647a81d1a2c2b1f
		}
	}
	__obf_1bdc4788b863aed0 = __obf_1bdc4788b863aed0.Add(__obf_a35a5e9e2799b84b)

	return __obf_1bdc4788b863aed0.Round(__obf_9e4f7aff89415e13), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_5debc2a983e87a03 Decimal) NumDigits() int {
	if __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42 == nil {
		return 1
	}

	if __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.IsInt64() {
		__obf_bc7d75022022aa58 := __obf_5debc2a983e87a03.
			// restrict fast path to integers with exact conversion to float64
			__obf_ab2c02ab751b9f42.Int64()

		if __obf_bc7d75022022aa58 <= (1<<53) && __obf_bc7d75022022aa58 >= -(1<<53) {
			if __obf_bc7d75022022aa58 == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_bc7d75022022aa58)))) + 1
		}
	}
	__obf_cc99aec1558f54ab := int(float64(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.BitLen()) / math.Log2(10))
	__obf_20ad68a699486e06 := // estimatedNumDigits (lg10) may be off by 1, need to verify
		big.NewInt(int64(__obf_cc99aec1558f54ab))
	__obf_cea743b96d1f1c96 := __obf_20ad68a699486e06.Exp(__obf_8f06f98f079e4122, __obf_20ad68a699486e06, nil)

	if __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.CmpAbs(__obf_cea743b96d1f1c96) >= 0 {
		return __obf_cc99aec1558f54ab + 1
	}

	return __obf_cc99aec1558f54ab
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_5debc2a983e87a03 Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_5debc2a983e87a03.__obf_d5f3440c0176542c >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_1a925cec70c560aa big.Int
	__obf_8316632570bb1426 := new(big.Int).Set(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42)
	for __obf_25508e8597181510 := __obf_b8bca49ac727c366(__obf_5debc2a983e87a03.__obf_d5f3440c0176542c); __obf_25508e8597181510 > 0; __obf_25508e8597181510-- {
		__obf_8316632570bb1426.
			QuoRem(__obf_8316632570bb1426, __obf_8f06f98f079e4122, &__obf_1a925cec70c560aa)
		if __obf_1a925cec70c560aa.Cmp(__obf_304be536ded8753b) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_b8bca49ac727c366(__obf_5618ed758d2ff827 int32) int32 {
	if __obf_5618ed758d2ff827 < 0 {
		return -__obf_5618ed758d2ff827
	}
	return __obf_5618ed758d2ff827
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_5debc2a983e87a03 Decimal) Cmp(__obf_e06b6478467c95fe Decimal) int {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()
	__obf_e06b6478467c95fe.__obf_23af5656aedaa490()

	if __obf_5debc2a983e87a03.__obf_d5f3440c0176542c == __obf_e06b6478467c95fe.__obf_d5f3440c0176542c {
		return __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.Cmp(__obf_e06b6478467c95fe.__obf_ab2c02ab751b9f42)
	}
	__obf_85728b8dfdaeed90, __obf_c83c6d3225e9675c := RescalePair(__obf_5debc2a983e87a03, __obf_e06b6478467c95fe)

	return __obf_85728b8dfdaeed90.__obf_ab2c02ab751b9f42.Cmp(__obf_c83c6d3225e9675c.

		// Compare compares the numbers represented by d and d2 and returns:
		//
		//	-1 if d <  d2
		//	 0 if d == d2
		//	+1 if d >  d2
		__obf_ab2c02ab751b9f42)
}

func (__obf_5debc2a983e87a03 Decimal) Compare(__obf_e06b6478467c95fe Decimal) int {
	return __obf_5debc2a983e87a03.Cmp(__obf_e06b6478467c95fe)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_5debc2a983e87a03 Decimal) Equal(__obf_e06b6478467c95fe Decimal) bool {
	return __obf_5debc2a983e87a03.Cmp(__obf_e06b6478467c95fe) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_5debc2a983e87a03 Decimal) Equals(__obf_e06b6478467c95fe Decimal) bool {
	return __obf_5debc2a983e87a03.Equal(__obf_e06b6478467c95fe)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_5debc2a983e87a03 Decimal) GreaterThan(__obf_e06b6478467c95fe Decimal) bool {
	return __obf_5debc2a983e87a03.Cmp(__obf_e06b6478467c95fe) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_5debc2a983e87a03 Decimal) GreaterThanOrEqual(__obf_e06b6478467c95fe Decimal) bool {
	__obf_28d3007bcc0cd5e1 := __obf_5debc2a983e87a03.Cmp(__obf_e06b6478467c95fe)
	return __obf_28d3007bcc0cd5e1 == 1 || __obf_28d3007bcc0cd5e1 == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_5debc2a983e87a03 Decimal) LessThan(__obf_e06b6478467c95fe Decimal) bool {
	return __obf_5debc2a983e87a03.Cmp(__obf_e06b6478467c95fe) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_5debc2a983e87a03 Decimal) LessThanOrEqual(__obf_e06b6478467c95fe Decimal) bool {
	__obf_28d3007bcc0cd5e1 := __obf_5debc2a983e87a03.Cmp(__obf_e06b6478467c95fe)
	return __obf_28d3007bcc0cd5e1 == -1 || __obf_28d3007bcc0cd5e1 == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_5debc2a983e87a03 Decimal) Sign() int {
	if __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42 == nil {
		return 0
	}
	return __obf_5debc2a983e87a03.

		// IsPositive return
		//
		//	true if d > 0
		//	false if d == 0
		//	false if d < 0
		__obf_ab2c02ab751b9f42.Sign()
}

func (__obf_5debc2a983e87a03 Decimal) IsPositive() bool {
	return __obf_5debc2a983e87a03.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_5debc2a983e87a03 Decimal) IsNegative() bool {
	return __obf_5debc2a983e87a03.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_5debc2a983e87a03 Decimal) IsZero() bool {
	return __obf_5debc2a983e87a03.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_5debc2a983e87a03 Decimal) Exponent() int32 {
	return __obf_5debc2a983e87a03.

		// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
		__obf_d5f3440c0176542c
}

func (__obf_5debc2a983e87a03 Decimal) Coefficient() *big.Int {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_5debc2a983e87a03.

		// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
		// If coefficient cannot be represented in an int64, the result will be undefined.
		__obf_ab2c02ab751b9f42)
}

func (__obf_5debc2a983e87a03 Decimal) CoefficientInt64() int64 {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()
	return __obf_5debc2a983e87a03.

		// IntPart returns the integer component of the decimal.
		__obf_ab2c02ab751b9f42.Int64()
}

func (__obf_5debc2a983e87a03 Decimal) IntPart() int64 {
	__obf_10844ef3f8d466fd := __obf_5debc2a983e87a03.__obf_c2312e5df65a128a(0)
	return __obf_10844ef3f8d466fd.__obf_ab2c02ab751b9f42.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_5debc2a983e87a03 Decimal) BigInt() *big.Int {
	__obf_10844ef3f8d466fd := __obf_5debc2a983e87a03.__obf_c2312e5df65a128a(0)
	return __obf_10844ef3f8d466fd.

		// BigFloat returns decimal as BigFloat.
		// Be aware that casting decimal to BigFloat might cause a loss of precision.
		__obf_ab2c02ab751b9f42
}

func (__obf_5debc2a983e87a03 Decimal) BigFloat() *big.Float {
	__obf_42f0cd6ce08f2588 := &big.Float{}
	__obf_42f0cd6ce08f2588.
		SetString(__obf_5debc2a983e87a03.String())
	return __obf_42f0cd6ce08f2588
}

// Rat returns a rational number representation of the decimal.
func (__obf_5debc2a983e87a03 Decimal) Rat() *big.Rat {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()
	if __obf_5debc2a983e87a03.
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_d5f3440c0176542c <= 0 {
		__obf_80d8366688ae52ce := new(big.Int).Exp(__obf_8f06f98f079e4122, big.NewInt(-int64(__obf_5debc2a983e87a03.__obf_d5f3440c0176542c)), nil)
		return new(big.Rat).SetFrac(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42, __obf_80d8366688ae52ce)
	}
	__obf_7497049ae08a8517 := new(big.Int).Exp(__obf_8f06f98f079e4122, big.NewInt(int64(__obf_5debc2a983e87a03.__obf_d5f3440c0176542c)), nil)
	__obf_b1005423fd5ead82 := new(big.Int).Mul(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42, __obf_7497049ae08a8517)
	return new(big.Rat).SetFrac(__obf_b1005423fd5ead82,

		// Float64 returns the nearest float64 value for d and a bool indicating
		// whether f represents d exactly.
		// For more details, see the documentation for big.Rat.Float64
		__obf_0ca31adf9d5bc652)
}

func (__obf_5debc2a983e87a03 Decimal) Float64() (__obf_42f0cd6ce08f2588 float64, __obf_5bf1b9c7afc7551b bool) {
	return __obf_5debc2a983e87a03.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_5debc2a983e87a03 Decimal) InexactFloat64() float64 {
	__obf_42f0cd6ce08f2588, _ := __obf_5debc2a983e87a03.Float64()
	return __obf_42f0cd6ce08f2588
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
func (__obf_5debc2a983e87a03 Decimal) String() string {
	return __obf_5debc2a983e87a03.string(true)
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
func (__obf_5debc2a983e87a03 Decimal) StringFixed(__obf_75015b2f7f038799 int32) string {
	__obf_18428e68cc3cefd9 := __obf_5debc2a983e87a03.Round(__obf_75015b2f7f038799)
	return __obf_18428e68cc3cefd9.string(false)
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
func (__obf_5debc2a983e87a03 Decimal) StringFixedBank(__obf_75015b2f7f038799 int32) string {
	__obf_18428e68cc3cefd9 := __obf_5debc2a983e87a03.RoundBank(__obf_75015b2f7f038799)
	return __obf_18428e68cc3cefd9.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_5debc2a983e87a03 Decimal) StringFixedCash(__obf_d903963582c887a0 uint8) string {
	__obf_18428e68cc3cefd9 := __obf_5debc2a983e87a03.RoundCash(__obf_d903963582c887a0)
	return __obf_18428e68cc3cefd9.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_5debc2a983e87a03 Decimal) Round(__obf_75015b2f7f038799 int32) Decimal {
	if __obf_5debc2a983e87a03.__obf_d5f3440c0176542c == -__obf_75015b2f7f038799 {
		return __obf_5debc2a983e87a03
	}
	__obf_0c78f20897d4a732 := // truncate to places + 1
		__obf_5debc2a983e87a03.__obf_c2312e5df65a128a(-__obf_75015b2f7f038799 - 1)

	// add sign(d) * 0.5
	if __obf_0c78f20897d4a732.__obf_ab2c02ab751b9f42.Sign() < 0 {
		__obf_0c78f20897d4a732.__obf_ab2c02ab751b9f42.
			Sub(__obf_0c78f20897d4a732.__obf_ab2c02ab751b9f42, __obf_656e224a8e2fd8a8)
	} else {
		__obf_0c78f20897d4a732.__obf_ab2c02ab751b9f42.
			Add(__obf_0c78f20897d4a732.__obf_ab2c02ab751b9f42,

				// floor for positive numbers, ceil for negative numbers
				__obf_656e224a8e2fd8a8)
	}

	_, __obf_dfa47acb5d50aedb := __obf_0c78f20897d4a732.__obf_ab2c02ab751b9f42.DivMod(__obf_0c78f20897d4a732.__obf_ab2c02ab751b9f42, __obf_8f06f98f079e4122, new(big.Int))
	__obf_0c78f20897d4a732.__obf_d5f3440c0176542c++
	if __obf_0c78f20897d4a732.__obf_ab2c02ab751b9f42.Sign() < 0 && __obf_dfa47acb5d50aedb.Cmp(__obf_304be536ded8753b) != 0 {
		__obf_0c78f20897d4a732.__obf_ab2c02ab751b9f42.
			Add(__obf_0c78f20897d4a732.__obf_ab2c02ab751b9f42,

				// RoundCeil rounds the decimal towards +infinity.
				//
				// Example:
				//
				//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
				//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
				//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
				//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
				__obf_0ca31adf9d5bc652)
	}

	return __obf_0c78f20897d4a732
}

func (__obf_5debc2a983e87a03 Decimal) RoundCeil(__obf_75015b2f7f038799 int32) Decimal {
	if __obf_5debc2a983e87a03.__obf_d5f3440c0176542c >= -__obf_75015b2f7f038799 {
		return __obf_5debc2a983e87a03
	}
	__obf_765815f339325114 := __obf_5debc2a983e87a03.__obf_c2312e5df65a128a(-__obf_75015b2f7f038799)
	if __obf_5debc2a983e87a03.Equal(__obf_765815f339325114) {
		return __obf_5debc2a983e87a03
	}

	if __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.Sign() > 0 {
		__obf_765815f339325114.__obf_ab2c02ab751b9f42.
			Add(__obf_765815f339325114.__obf_ab2c02ab751b9f42, __obf_0ca31adf9d5bc652)
	}

	return __obf_765815f339325114
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_5debc2a983e87a03 Decimal) RoundFloor(__obf_75015b2f7f038799 int32) Decimal {
	if __obf_5debc2a983e87a03.__obf_d5f3440c0176542c >= -__obf_75015b2f7f038799 {
		return __obf_5debc2a983e87a03
	}
	__obf_765815f339325114 := __obf_5debc2a983e87a03.__obf_c2312e5df65a128a(-__obf_75015b2f7f038799)
	if __obf_5debc2a983e87a03.Equal(__obf_765815f339325114) {
		return __obf_5debc2a983e87a03
	}

	if __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.Sign() < 0 {
		__obf_765815f339325114.__obf_ab2c02ab751b9f42.
			Sub(__obf_765815f339325114.__obf_ab2c02ab751b9f42, __obf_0ca31adf9d5bc652)
	}

	return __obf_765815f339325114
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_5debc2a983e87a03 Decimal) RoundUp(__obf_75015b2f7f038799 int32) Decimal {
	if __obf_5debc2a983e87a03.__obf_d5f3440c0176542c >= -__obf_75015b2f7f038799 {
		return __obf_5debc2a983e87a03
	}
	__obf_765815f339325114 := __obf_5debc2a983e87a03.__obf_c2312e5df65a128a(-__obf_75015b2f7f038799)
	if __obf_5debc2a983e87a03.Equal(__obf_765815f339325114) {
		return __obf_5debc2a983e87a03
	}

	if __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.Sign() > 0 {
		__obf_765815f339325114.__obf_ab2c02ab751b9f42.
			Add(__obf_765815f339325114.__obf_ab2c02ab751b9f42, __obf_0ca31adf9d5bc652)
	} else if __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.Sign() < 0 {
		__obf_765815f339325114.__obf_ab2c02ab751b9f42.
			Sub(__obf_765815f339325114.__obf_ab2c02ab751b9f42, __obf_0ca31adf9d5bc652)
	}

	return __obf_765815f339325114
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_5debc2a983e87a03 Decimal) RoundDown(__obf_75015b2f7f038799 int32) Decimal {
	if __obf_5debc2a983e87a03.__obf_d5f3440c0176542c >= -__obf_75015b2f7f038799 {
		return __obf_5debc2a983e87a03
	}
	__obf_765815f339325114 := __obf_5debc2a983e87a03.__obf_c2312e5df65a128a(-__obf_75015b2f7f038799)
	if __obf_5debc2a983e87a03.Equal(__obf_765815f339325114) {
		return __obf_5debc2a983e87a03
	}
	return __obf_765815f339325114
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
func (__obf_5debc2a983e87a03 Decimal) RoundBank(__obf_75015b2f7f038799 int32) Decimal {
	__obf_68beb60c8424efc2 := __obf_5debc2a983e87a03.Round(__obf_75015b2f7f038799)
	__obf_b19449e9bb087ae5 := __obf_5debc2a983e87a03.Sub(__obf_68beb60c8424efc2).Abs()
	__obf_e765f21ab284a5fd := New(5, -__obf_75015b2f7f038799-1)
	if __obf_b19449e9bb087ae5.Cmp(__obf_e765f21ab284a5fd) == 0 && __obf_68beb60c8424efc2.__obf_ab2c02ab751b9f42.Bit(0) != 0 {
		if __obf_68beb60c8424efc2.__obf_ab2c02ab751b9f42.Sign() < 0 {
			__obf_68beb60c8424efc2.__obf_ab2c02ab751b9f42.
				Add(__obf_68beb60c8424efc2.__obf_ab2c02ab751b9f42, __obf_0ca31adf9d5bc652)
		} else {
			__obf_68beb60c8424efc2.__obf_ab2c02ab751b9f42.
				Sub(__obf_68beb60c8424efc2.__obf_ab2c02ab751b9f42, __obf_0ca31adf9d5bc652)
		}
	}

	return __obf_68beb60c8424efc2
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
func (__obf_5debc2a983e87a03 Decimal) RoundCash(__obf_d903963582c887a0 uint8) Decimal {
	var __obf_721d743dfb64b7c2 *big.Int
	switch __obf_d903963582c887a0 {
	case 5:
		__obf_721d743dfb64b7c2 = __obf_015884ffa111f28b
	case 10:
		__obf_721d743dfb64b7c2 = __obf_8f06f98f079e4122
	case 25:
		__obf_721d743dfb64b7c2 = __obf_e62aec2193220910
	case 50:
		__obf_721d743dfb64b7c2 = __obf_51443023c3dfc1b9
	case 100:
		__obf_721d743dfb64b7c2 = __obf_0ca31adf9d5bc652
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_d903963582c887a0))
	}
	__obf_6a191a995d8f4733 := Decimal{__obf_ab2c02ab751b9f42: __obf_721d743dfb64b7c2}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_5debc2a983e87a03.Mul(__obf_6a191a995d8f4733).Round(0).Div(__obf_6a191a995d8f4733).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_5debc2a983e87a03 Decimal) Floor() Decimal {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()

	if __obf_5debc2a983e87a03.__obf_d5f3440c0176542c >= 0 {
		return __obf_5debc2a983e87a03
	}
	__obf_d5f3440c0176542c := big.NewInt(10)
	__obf_d5f3440c0176542c.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_d5f3440c0176542c, big.NewInt(-int64(__obf_5debc2a983e87a03.__obf_d5f3440c0176542c)), nil)
	__obf_25508e8597181510 := new(big.Int).Div(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42, __obf_d5f3440c0176542c)
	return Decimal{__obf_ab2c02ab751b9f42: __obf_25508e8597181510,

		// Ceil returns the nearest integer value greater than or equal to d.
		__obf_d5f3440c0176542c: 0}
}

func (__obf_5debc2a983e87a03 Decimal) Ceil() Decimal {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()

	if __obf_5debc2a983e87a03.__obf_d5f3440c0176542c >= 0 {
		return __obf_5debc2a983e87a03
	}
	__obf_d5f3440c0176542c := big.NewInt(10)
	__obf_d5f3440c0176542c.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_d5f3440c0176542c, big.NewInt(-int64(__obf_5debc2a983e87a03.__obf_d5f3440c0176542c)), nil)
	__obf_25508e8597181510, __obf_dfa47acb5d50aedb := new(big.Int).DivMod(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42, __obf_d5f3440c0176542c, new(big.Int))
	if __obf_dfa47acb5d50aedb.Cmp(__obf_304be536ded8753b) != 0 {
		__obf_25508e8597181510.
			Add(__obf_25508e8597181510, __obf_0ca31adf9d5bc652)
	}
	return Decimal{__obf_ab2c02ab751b9f42: __obf_25508e8597181510,

		// Truncate truncates off digits from the number, without rounding.
		//
		// NOTE: precision is the last digit that will not be truncated (must be >= 0).
		//
		// Example:
		//
		//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
		__obf_d5f3440c0176542c: 0}
}

func (__obf_5debc2a983e87a03 Decimal) Truncate(__obf_9e4f7aff89415e13 int32) Decimal {
	__obf_5debc2a983e87a03.__obf_23af5656aedaa490()
	if __obf_9e4f7aff89415e13 >= 0 && -__obf_9e4f7aff89415e13 > __obf_5debc2a983e87a03.__obf_d5f3440c0176542c {
		return __obf_5debc2a983e87a03.__obf_c2312e5df65a128a(-__obf_9e4f7aff89415e13)
	}
	return __obf_5debc2a983e87a03
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_5debc2a983e87a03 *Decimal) UnmarshalJSON(__obf_28e009fc002656b2 []byte) error {
	if string(__obf_28e009fc002656b2) == "null" {
		return nil
	}
	__obf_e516f3570d706736, __obf_239001ea6a8e4288 := __obf_985175b6d3df6456(__obf_28e009fc002656b2)
	if __obf_239001ea6a8e4288 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_28e009fc002656b2, __obf_239001ea6a8e4288)
	}
	__obf_a5d0662d16479f1d, __obf_239001ea6a8e4288 := NewFromString(__obf_e516f3570d706736)
	*__obf_5debc2a983e87a03 = __obf_a5d0662d16479f1d
	if __obf_239001ea6a8e4288 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_e516f3570d706736, __obf_239001ea6a8e4288)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_5debc2a983e87a03 Decimal) MarshalJSON() ([]byte, error) {
	var __obf_e516f3570d706736 string
	if MarshalJSONWithoutQuotes {
		__obf_e516f3570d706736 = __obf_5debc2a983e87a03.String()
	} else {
		__obf_e516f3570d706736 = "\"" + __obf_5debc2a983e87a03.String() + "\""
	}
	return []byte(__obf_e516f3570d706736), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_5debc2a983e87a03 *Decimal) UnmarshalBinary(__obf_e7192161ccc808e1 []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_e7192161ccc808e1) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_e7192161ccc808e1, len(__obf_e7192161ccc808e1))
	}
	__obf_5debc2a983e87a03.

		// Extract the exponent
		__obf_d5f3440c0176542c = int32(binary.BigEndian.Uint32(__obf_e7192161ccc808e1[:4]))
	__obf_5debc2a983e87a03.

		// Extract the value
		__obf_ab2c02ab751b9f42 = new(big.Int)
	if __obf_239001ea6a8e4288 := __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.GobDecode(__obf_e7192161ccc808e1[4:]); __obf_239001ea6a8e4288 != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_e7192161ccc808e1, __obf_239001ea6a8e4288)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_5debc2a983e87a03 Decimal) MarshalBinary() (__obf_e7192161ccc808e1 []byte, __obf_239001ea6a8e4288 error) {
	// exp is written first, but encode value first to know output size
	var __obf_819f8698f3197616 []byte
	if __obf_819f8698f3197616, __obf_239001ea6a8e4288 = __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.GobEncode(); __obf_239001ea6a8e4288 != nil {
		return nil, __obf_239001ea6a8e4288
	}
	__obf_1ac4d7a26cbb3166 := // Write the exponent in front, since it's a fixed size
		make([]byte, 4, len(__obf_819f8698f3197616)+4)
	binary.BigEndian.PutUint32(__obf_1ac4d7a26cbb3166, uint32(__obf_5debc2a983e87a03.

		// Return the byte array
		__obf_d5f3440c0176542c))

	return append(__obf_1ac4d7a26cbb3166, __obf_819f8698f3197616...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_5debc2a983e87a03 *Decimal) Scan(__obf_ab2c02ab751b9f42 any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_3d81c055a6608685 := __obf_ab2c02ab751b9f42.(type) {

	case float32:
		*__obf_5debc2a983e87a03 = NewFromFloat(float64(__obf_3d81c055a6608685))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_5debc2a983e87a03 = NewFromFloat(__obf_3d81c055a6608685)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_5debc2a983e87a03 = New(__obf_3d81c055a6608685, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_5debc2a983e87a03 = NewFromUint64(__obf_3d81c055a6608685)
		return nil

	default:
		__obf_e516f3570d706736,
			// default is trying to interpret value stored as string
			__obf_239001ea6a8e4288 := __obf_985175b6d3df6456(__obf_3d81c055a6608685)
		if __obf_239001ea6a8e4288 != nil {
			return __obf_239001ea6a8e4288
		}
		*__obf_5debc2a983e87a03, __obf_239001ea6a8e4288 = NewFromString(__obf_e516f3570d706736)
		return __obf_239001ea6a8e4288
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_5debc2a983e87a03 Decimal) Value() (driver.Value, error) {
	return __obf_5debc2a983e87a03.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_5debc2a983e87a03 *Decimal) UnmarshalText(__obf_66be6589b070a7f7 []byte) error {
	__obf_e516f3570d706736 := string(__obf_66be6589b070a7f7)
	__obf_5b220531a71314a4, __obf_239001ea6a8e4288 := NewFromString(__obf_e516f3570d706736)
	*__obf_5debc2a983e87a03 = __obf_5b220531a71314a4
	if __obf_239001ea6a8e4288 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_e516f3570d706736, __obf_239001ea6a8e4288)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_5debc2a983e87a03 Decimal) MarshalText() (__obf_66be6589b070a7f7 []byte, __obf_239001ea6a8e4288 error) {
	return []byte(__obf_5debc2a983e87a03.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_5debc2a983e87a03 Decimal) GobEncode() ([]byte, error) {
	return __obf_5debc2a983e87a03.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_5debc2a983e87a03 *Decimal) GobDecode(__obf_e7192161ccc808e1 []byte) error {
	return __obf_5debc2a983e87a03.UnmarshalBinary(__obf_e7192161ccc808e1)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_5debc2a983e87a03 Decimal) StringScaled(__obf_d5f3440c0176542c int32) string {
	return __obf_5debc2a983e87a03.__obf_c2312e5df65a128a(__obf_d5f3440c0176542c).String()
}

func (__obf_5debc2a983e87a03 Decimal) string(__obf_278e4fc57e567d24 bool) string {
	if __obf_5debc2a983e87a03.__obf_d5f3440c0176542c >= 0 {
		return __obf_5debc2a983e87a03.__obf_c2312e5df65a128a(0).__obf_ab2c02ab751b9f42.String()
	}
	__obf_b8bca49ac727c366 := new(big.Int).Abs(__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42)
	__obf_e516f3570d706736 := __obf_b8bca49ac727c366.String()

	var __obf_e0b7ab871822d70f, __obf_cb76c5d2681653de string
	__obf_fcb4f30ebc868559 := // NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
		// and you are on a 32-bit machine. Won't fix this super-edge case.
		int(__obf_5debc2a983e87a03.__obf_d5f3440c0176542c)
	if len(__obf_e516f3570d706736) > -__obf_fcb4f30ebc868559 {
		__obf_e0b7ab871822d70f = __obf_e516f3570d706736[:len(__obf_e516f3570d706736)+__obf_fcb4f30ebc868559]
		__obf_cb76c5d2681653de = __obf_e516f3570d706736[len(__obf_e516f3570d706736)+__obf_fcb4f30ebc868559:]
	} else {
		__obf_e0b7ab871822d70f = "0"
		__obf_a29e960303b03738 := -__obf_fcb4f30ebc868559 - len(__obf_e516f3570d706736)
		__obf_cb76c5d2681653de = strings.Repeat("0", __obf_a29e960303b03738) + __obf_e516f3570d706736
	}

	if __obf_278e4fc57e567d24 {
		__obf_283995368e36d3d0 := len(__obf_cb76c5d2681653de) - 1
		for ; __obf_283995368e36d3d0 >= 0; __obf_283995368e36d3d0-- {
			if __obf_cb76c5d2681653de[__obf_283995368e36d3d0] != '0' {
				break
			}
		}
		__obf_cb76c5d2681653de = __obf_cb76c5d2681653de[:__obf_283995368e36d3d0+1]
	}
	__obf_0a0327a35314c231 := __obf_e0b7ab871822d70f
	if len(__obf_cb76c5d2681653de) > 0 {
		__obf_0a0327a35314c231 += "." + __obf_cb76c5d2681653de
	}

	if __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42.Sign() < 0 {
		return "-" + __obf_0a0327a35314c231
	}

	return __obf_0a0327a35314c231
}

func (__obf_5debc2a983e87a03 *Decimal) __obf_23af5656aedaa490() {
	if __obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42 == nil {
		__obf_5debc2a983e87a03.__obf_ab2c02ab751b9f42 = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_0dc4b9bde0a0bdaa Decimal, __obf_8ef89cb57cc9acac ...Decimal) Decimal {
	__obf_e26478d44a24982c := __obf_0dc4b9bde0a0bdaa
	for _, __obf_96af40f0ab2e4220 := range __obf_8ef89cb57cc9acac {
		if __obf_96af40f0ab2e4220.Cmp(__obf_e26478d44a24982c) < 0 {
			__obf_e26478d44a24982c = __obf_96af40f0ab2e4220
		}
	}
	return __obf_e26478d44a24982c
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_0dc4b9bde0a0bdaa Decimal, __obf_8ef89cb57cc9acac ...Decimal) Decimal {
	__obf_e26478d44a24982c := __obf_0dc4b9bde0a0bdaa
	for _, __obf_96af40f0ab2e4220 := range __obf_8ef89cb57cc9acac {
		if __obf_96af40f0ab2e4220.Cmp(__obf_e26478d44a24982c) > 0 {
			__obf_e26478d44a24982c = __obf_96af40f0ab2e4220
		}
	}
	return __obf_e26478d44a24982c
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_0dc4b9bde0a0bdaa Decimal, __obf_8ef89cb57cc9acac ...Decimal) Decimal {
	__obf_855b8cbe2cbd85c6 := __obf_0dc4b9bde0a0bdaa
	for _, __obf_96af40f0ab2e4220 := range __obf_8ef89cb57cc9acac {
		__obf_855b8cbe2cbd85c6 = __obf_855b8cbe2cbd85c6.Add(__obf_96af40f0ab2e4220)
	}

	return __obf_855b8cbe2cbd85c6
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_0dc4b9bde0a0bdaa Decimal, __obf_8ef89cb57cc9acac ...Decimal) Decimal {
	__obf_969d7bf073f1b996 := New(int64(len(__obf_8ef89cb57cc9acac)+1), 0)
	__obf_272fc79b42d8fb7b := Sum(__obf_0dc4b9bde0a0bdaa, __obf_8ef89cb57cc9acac...)
	return __obf_272fc79b42d8fb7b.Div(__obf_969d7bf073f1b996)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_99b4277faaad444c Decimal, __obf_e06b6478467c95fe Decimal) (Decimal, Decimal) {
	__obf_99b4277faaad444c.__obf_23af5656aedaa490()
	__obf_e06b6478467c95fe.__obf_23af5656aedaa490()

	if __obf_99b4277faaad444c.__obf_d5f3440c0176542c < __obf_e06b6478467c95fe.__obf_d5f3440c0176542c {
		return __obf_99b4277faaad444c, __obf_e06b6478467c95fe.__obf_c2312e5df65a128a(__obf_99b4277faaad444c.__obf_d5f3440c0176542c)
	} else if __obf_99b4277faaad444c.__obf_d5f3440c0176542c > __obf_e06b6478467c95fe.__obf_d5f3440c0176542c {
		return __obf_99b4277faaad444c.__obf_c2312e5df65a128a(__obf_e06b6478467c95fe.__obf_d5f3440c0176542c), __obf_e06b6478467c95fe
	}

	return __obf_99b4277faaad444c, __obf_e06b6478467c95fe
}

func __obf_985175b6d3df6456(__obf_ab2c02ab751b9f42 any) (string, error) {
	var __obf_ab8d869129b68052 []byte

	switch __obf_3d81c055a6608685 := __obf_ab2c02ab751b9f42.(type) {
	case string:
		__obf_ab8d869129b68052 = []byte(__obf_3d81c055a6608685)
	case []byte:
		__obf_ab8d869129b68052 = __obf_3d81c055a6608685
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_ab2c02ab751b9f42,

			// If the amount is quoted, strip the quotes
			__obf_ab2c02ab751b9f42)
	}

	if len(__obf_ab8d869129b68052) > 2 && __obf_ab8d869129b68052[0] == '"' && __obf_ab8d869129b68052[len(__obf_ab8d869129b68052)-1] == '"' {
		__obf_ab8d869129b68052 = __obf_ab8d869129b68052[1 : len(__obf_ab8d869129b68052)-1]
	}
	return string(__obf_ab8d869129b68052), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the datautil.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_5debc2a983e87a03 Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_5debc2a983e87a03,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_5debc2a983e87a03 *NullDecimal) Scan(__obf_ab2c02ab751b9f42 any) error {
	if __obf_ab2c02ab751b9f42 == nil {
		__obf_5debc2a983e87a03.
			Valid = false
		return nil
	}
	__obf_5debc2a983e87a03.
		Valid = true
	return __obf_5debc2a983e87a03.Decimal.Scan(__obf_ab2c02ab751b9f42)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_5debc2a983e87a03 NullDecimal) Value() (driver.Value, error) {
	if !__obf_5debc2a983e87a03.Valid {
		return nil, nil
	}
	return __obf_5debc2a983e87a03.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_5debc2a983e87a03 *NullDecimal) UnmarshalJSON(__obf_28e009fc002656b2 []byte) error {
	if string(__obf_28e009fc002656b2) == "null" {
		__obf_5debc2a983e87a03.
			Valid = false
		return nil
	}
	__obf_5debc2a983e87a03.
		Valid = true
	return __obf_5debc2a983e87a03.Decimal.UnmarshalJSON(__obf_28e009fc002656b2)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_5debc2a983e87a03 NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_5debc2a983e87a03.Valid {
		return []byte("null"), nil
	}
	return __obf_5debc2a983e87a03.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_5debc2a983e87a03 *NullDecimal) UnmarshalText(__obf_66be6589b070a7f7 []byte) error {
	__obf_e516f3570d706736 := string(__obf_66be6589b070a7f7)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_e516f3570d706736 == "" {
		__obf_5debc2a983e87a03.
			Valid = false
		return nil
	}
	if __obf_239001ea6a8e4288 := __obf_5debc2a983e87a03.Decimal.UnmarshalText(__obf_66be6589b070a7f7); __obf_239001ea6a8e4288 != nil {
		__obf_5debc2a983e87a03.
			Valid = false
		return __obf_239001ea6a8e4288
	}
	__obf_5debc2a983e87a03.
		Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_5debc2a983e87a03 NullDecimal) MarshalText() (__obf_66be6589b070a7f7 []byte, __obf_239001ea6a8e4288 error) {
	if !__obf_5debc2a983e87a03.Valid {
		return []byte{}, nil
	}
	return __obf_5debc2a983e87a03.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_5debc2a983e87a03 Decimal) Atan() Decimal {
	if __obf_5debc2a983e87a03.Equal(NewFromFloat(0.0)) {
		return __obf_5debc2a983e87a03
	}
	if __obf_5debc2a983e87a03.GreaterThan(NewFromFloat(0.0)) {
		return __obf_5debc2a983e87a03.__obf_7f9d92d19e9f1691()
	}
	return __obf_5debc2a983e87a03.Neg().__obf_7f9d92d19e9f1691().Neg()
}

func (__obf_5debc2a983e87a03 Decimal) __obf_ef1fa7eeb76bc692() Decimal {
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
	__obf_25508e8597181510 := __obf_5debc2a983e87a03.Mul(__obf_5debc2a983e87a03)
	__obf_356fa5357dd192a1 := P0.Mul(__obf_25508e8597181510).Add(P1).Mul(__obf_25508e8597181510).Add(P2).Mul(__obf_25508e8597181510).Add(P3).Mul(__obf_25508e8597181510).Add(P4).Mul(__obf_25508e8597181510)
	__obf_874a9b5bc7a587fa := __obf_25508e8597181510.Add(Q0).Mul(__obf_25508e8597181510).Add(Q1).Mul(__obf_25508e8597181510).Add(Q2).Mul(__obf_25508e8597181510).Add(Q3).Mul(__obf_25508e8597181510).Add(Q4)
	__obf_25508e8597181510 = __obf_356fa5357dd192a1.Div(__obf_874a9b5bc7a587fa)
	__obf_25508e8597181510 = __obf_5debc2a983e87a03.Mul(__obf_25508e8597181510).Add(__obf_5debc2a983e87a03)
	return __obf_25508e8597181510
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_5debc2a983e87a03 Decimal) __obf_7f9d92d19e9f1691() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)
	__obf_b8985fd99922e1c4 := // tan(3*pi/8)
		NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_5debc2a983e87a03.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_5debc2a983e87a03.__obf_ef1fa7eeb76bc692()
	}
	if __obf_5debc2a983e87a03.GreaterThan(Tan3pio8) {
		return __obf_b8985fd99922e1c4.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_5debc2a983e87a03).__obf_ef1fa7eeb76bc692()).Add(Morebits)
	}
	return __obf_b8985fd99922e1c4.Div(NewFromFloat(4.0)).Add((__obf_5debc2a983e87a03.Sub(NewFromFloat(1.0)).Div(__obf_5debc2a983e87a03.Add(NewFromFloat(1.0)))).__obf_ef1fa7eeb76bc692()).Add(NewFromFloat(0.5).Mul(Morebits))
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
func (__obf_5debc2a983e87a03 Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_5debc2a983e87a03.Equal(NewFromFloat(0.0)) {
		return __obf_5debc2a983e87a03
	}
	__obf_f49c051f0678a1f6 := // make argument positive but save the sign
		false
	if __obf_5debc2a983e87a03.LessThan(NewFromFloat(0.0)) {
		__obf_5debc2a983e87a03 = __obf_5debc2a983e87a03.Neg()
		__obf_f49c051f0678a1f6 = true
	}
	__obf_da55bcf58963da1c := __obf_5debc2a983e87a03.Mul(M4PI).IntPart()
	__obf_2e2f27af7d607942 := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_da55bcf58963da1c)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_da55bcf58963da1c&1 == 1 {
		__obf_da55bcf58963da1c++
		__obf_2e2f27af7d607942 = __obf_2e2f27af7d607942.Add(NewFromFloat(1.0))
	}
	__obf_da55bcf58963da1c &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_da55bcf58963da1c > 3 {
		__obf_f49c051f0678a1f6 = !__obf_f49c051f0678a1f6
		__obf_da55bcf58963da1c -= 4
	}
	__obf_25508e8597181510 := __obf_5debc2a983e87a03.Sub(__obf_2e2f27af7d607942.Mul(PI4A)).Sub(__obf_2e2f27af7d607942.Mul(PI4B)).Sub(__obf_2e2f27af7d607942.Mul(PI4C))
	__obf_c3dc0170b9308efd := // Extended precision modular arithmetic
		__obf_25508e8597181510.Mul(__obf_25508e8597181510)

	if __obf_da55bcf58963da1c == 1 || __obf_da55bcf58963da1c == 2 {
		__obf_b3105e0528f6edc5 := __obf_c3dc0170b9308efd.Mul(__obf_c3dc0170b9308efd).Mul(_cos[0].Mul(__obf_c3dc0170b9308efd).Add(_cos[1]).Mul(__obf_c3dc0170b9308efd).Add(_cos[2]).Mul(__obf_c3dc0170b9308efd).Add(_cos[3]).Mul(__obf_c3dc0170b9308efd).Add(_cos[4]).Mul(__obf_c3dc0170b9308efd).Add(_cos[5]))
		__obf_2e2f27af7d607942 = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_c3dc0170b9308efd)).Add(__obf_b3105e0528f6edc5)
	} else {
		__obf_2e2f27af7d607942 = __obf_25508e8597181510.Add(__obf_25508e8597181510.Mul(__obf_c3dc0170b9308efd).Mul(_sin[0].Mul(__obf_c3dc0170b9308efd).Add(_sin[1]).Mul(__obf_c3dc0170b9308efd).Add(_sin[2]).Mul(__obf_c3dc0170b9308efd).Add(_sin[3]).Mul(__obf_c3dc0170b9308efd).Add(_sin[4]).Mul(__obf_c3dc0170b9308efd).Add(_sin[5])))
	}
	if __obf_f49c051f0678a1f6 {
		__obf_2e2f27af7d607942 = __obf_2e2f27af7d607942.Neg()
	}
	return __obf_2e2f27af7d607942
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
func (__obf_5debc2a983e87a03 Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)  // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)  // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15) // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125)
	__obf_f49c051f0678a1f6 := // 4/pi

		// make argument positive
		false
	if __obf_5debc2a983e87a03.LessThan(NewFromFloat(0.0)) {
		__obf_5debc2a983e87a03 = __obf_5debc2a983e87a03.Neg()
	}
	__obf_da55bcf58963da1c := __obf_5debc2a983e87a03.Mul(M4PI).IntPart()
	__obf_2e2f27af7d607942 := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_da55bcf58963da1c)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_da55bcf58963da1c&1 == 1 {
		__obf_da55bcf58963da1c++
		__obf_2e2f27af7d607942 = __obf_2e2f27af7d607942.Add(NewFromFloat(1.0))
	}
	__obf_da55bcf58963da1c &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_da55bcf58963da1c > 3 {
		__obf_f49c051f0678a1f6 = !__obf_f49c051f0678a1f6
		__obf_da55bcf58963da1c -= 4
	}
	if __obf_da55bcf58963da1c > 1 {
		__obf_f49c051f0678a1f6 = !__obf_f49c051f0678a1f6
	}
	__obf_25508e8597181510 := __obf_5debc2a983e87a03.Sub(__obf_2e2f27af7d607942.Mul(PI4A)).Sub(__obf_2e2f27af7d607942.Mul(PI4B)).Sub(__obf_2e2f27af7d607942.Mul(PI4C))
	__obf_c3dc0170b9308efd := // Extended precision modular arithmetic
		__obf_25508e8597181510.Mul(__obf_25508e8597181510)

	if __obf_da55bcf58963da1c == 1 || __obf_da55bcf58963da1c == 2 {
		__obf_2e2f27af7d607942 = __obf_25508e8597181510.Add(__obf_25508e8597181510.Mul(__obf_c3dc0170b9308efd).Mul(_sin[0].Mul(__obf_c3dc0170b9308efd).Add(_sin[1]).Mul(__obf_c3dc0170b9308efd).Add(_sin[2]).Mul(__obf_c3dc0170b9308efd).Add(_sin[3]).Mul(__obf_c3dc0170b9308efd).Add(_sin[4]).Mul(__obf_c3dc0170b9308efd).Add(_sin[5])))
	} else {
		__obf_b3105e0528f6edc5 := __obf_c3dc0170b9308efd.Mul(__obf_c3dc0170b9308efd).Mul(_cos[0].Mul(__obf_c3dc0170b9308efd).Add(_cos[1]).Mul(__obf_c3dc0170b9308efd).Add(_cos[2]).Mul(__obf_c3dc0170b9308efd).Add(_cos[3]).Mul(__obf_c3dc0170b9308efd).Add(_cos[4]).Mul(__obf_c3dc0170b9308efd).Add(_cos[5]))
		__obf_2e2f27af7d607942 = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_c3dc0170b9308efd)).Add(__obf_b3105e0528f6edc5)
	}
	if __obf_f49c051f0678a1f6 {
		__obf_2e2f27af7d607942 = __obf_2e2f27af7d607942.Neg()
	}
	return __obf_2e2f27af7d607942
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
func (__obf_5debc2a983e87a03 Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_5debc2a983e87a03.Equal(NewFromFloat(0.0)) {
		return __obf_5debc2a983e87a03
	}
	__obf_f49c051f0678a1f6 := // make argument positive but save the sign
		false
	if __obf_5debc2a983e87a03.LessThan(NewFromFloat(0.0)) {
		__obf_5debc2a983e87a03 = __obf_5debc2a983e87a03.Neg()
		__obf_f49c051f0678a1f6 = true
	}
	__obf_da55bcf58963da1c := __obf_5debc2a983e87a03.Mul(M4PI).IntPart()
	__obf_2e2f27af7d607942 := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_da55bcf58963da1c)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_da55bcf58963da1c&1 == 1 {
		__obf_da55bcf58963da1c++
		__obf_2e2f27af7d607942 = __obf_2e2f27af7d607942.Add(NewFromFloat(1.0))
	}
	__obf_25508e8597181510 := __obf_5debc2a983e87a03.Sub(__obf_2e2f27af7d607942.Mul(PI4A)).Sub(__obf_2e2f27af7d607942.Mul(PI4B)).Sub(__obf_2e2f27af7d607942.Mul(PI4C))
	__obf_c3dc0170b9308efd := // Extended precision modular arithmetic
		__obf_25508e8597181510.Mul(__obf_25508e8597181510)

	if __obf_c3dc0170b9308efd.GreaterThan(NewFromFloat(1e-14)) {
		__obf_b3105e0528f6edc5 := __obf_c3dc0170b9308efd.Mul(_tanP[0].Mul(__obf_c3dc0170b9308efd).Add(_tanP[1]).Mul(__obf_c3dc0170b9308efd).Add(_tanP[2]))
		__obf_f1727b8ff3895dec := __obf_c3dc0170b9308efd.Add(_tanQ[1]).Mul(__obf_c3dc0170b9308efd).Add(_tanQ[2]).Mul(__obf_c3dc0170b9308efd).Add(_tanQ[3]).Mul(__obf_c3dc0170b9308efd).Add(_tanQ[4])
		__obf_2e2f27af7d607942 = __obf_25508e8597181510.Add(__obf_25508e8597181510.Mul(__obf_b3105e0528f6edc5.Div(__obf_f1727b8ff3895dec)))
	} else {
		__obf_2e2f27af7d607942 = __obf_25508e8597181510
	}
	if __obf_da55bcf58963da1c&2 == 2 {
		__obf_2e2f27af7d607942 = NewFromFloat(-1.0).Div(__obf_2e2f27af7d607942)
	}
	if __obf_f49c051f0678a1f6 {
		__obf_2e2f27af7d607942 = __obf_2e2f27af7d607942.Neg()
	}
	return __obf_2e2f27af7d607942
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

type __obf_a5d0662d16479f1d struct {
	__obf_5debc2a983e87a03 [800]byte
	__obf_d5ebe437c13ed047 int  // digits, big-endian representation
	__obf_4e162ef443f77409 int  // number of digits used
	__obf_5084f9ae1c148972 bool // decimal point
	__obf_bd845c2e13a22c30 bool // negative flag
	// discarded nonzero digits beyond d[:nd]
}

func (__obf_53ac9a541c64b86f *__obf_a5d0662d16479f1d) String() string {
	__obf_5618ed758d2ff827 := 10 + __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047
	if __obf_53ac9a541c64b86f.__obf_4e162ef443f77409 > 0 {
		__obf_5618ed758d2ff827 += __obf_53ac9a541c64b86f.__obf_4e162ef443f77409
	}
	if __obf_53ac9a541c64b86f.__obf_4e162ef443f77409 < 0 {
		__obf_5618ed758d2ff827 += -__obf_53ac9a541c64b86f.__obf_4e162ef443f77409
	}
	__obf_7f731835d5cdcc94 := make([]byte, __obf_5618ed758d2ff827)
	__obf_b3105e0528f6edc5 := 0
	switch {
	case __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 == 0:
		return "0"

	case __obf_53ac9a541c64b86f.
		// zeros fill space between decimal point and digits
		__obf_4e162ef443f77409 <= 0:
		__obf_7f731835d5cdcc94[__obf_b3105e0528f6edc5] = '0'
		__obf_b3105e0528f6edc5++
		__obf_7f731835d5cdcc94[__obf_b3105e0528f6edc5] = '.'
		__obf_b3105e0528f6edc5++
		__obf_b3105e0528f6edc5 += __obf_778b33ac757de0d7(__obf_7f731835d5cdcc94[__obf_b3105e0528f6edc5 : __obf_b3105e0528f6edc5+-__obf_53ac9a541c64b86f.__obf_4e162ef443f77409])
		__obf_b3105e0528f6edc5 += copy(__obf_7f731835d5cdcc94[__obf_b3105e0528f6edc5:], __obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[0:__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047])

	case __obf_53ac9a541c64b86f.
		// decimal point in middle of digits
		__obf_4e162ef443f77409 < __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047:
		__obf_b3105e0528f6edc5 += copy(__obf_7f731835d5cdcc94[__obf_b3105e0528f6edc5:], __obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[0:__obf_53ac9a541c64b86f.__obf_4e162ef443f77409])
		__obf_7f731835d5cdcc94[__obf_b3105e0528f6edc5] = '.'
		__obf_b3105e0528f6edc5++
		__obf_b3105e0528f6edc5 += copy(__obf_7f731835d5cdcc94[__obf_b3105e0528f6edc5:], __obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_53ac9a541c64b86f.__obf_4e162ef443f77409:

		// zeros fill space between digits and decimal point
		__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047])

	default:
		__obf_b3105e0528f6edc5 += copy(__obf_7f731835d5cdcc94[__obf_b3105e0528f6edc5:], __obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[0:__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047])
		__obf_b3105e0528f6edc5 += __obf_778b33ac757de0d7(__obf_7f731835d5cdcc94[__obf_b3105e0528f6edc5 : __obf_b3105e0528f6edc5+__obf_53ac9a541c64b86f.__obf_4e162ef443f77409-__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047])
	}
	return string(__obf_7f731835d5cdcc94[0:__obf_b3105e0528f6edc5])
}

func __obf_778b33ac757de0d7(__obf_d94dd6fa70f06b05 []byte) int {
	for __obf_283995368e36d3d0 := range __obf_d94dd6fa70f06b05 {
		__obf_d94dd6fa70f06b05[__obf_283995368e36d3d0] = '0'
	}
	return len(__obf_d94dd6fa70f06b05)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_1db2cd3f69354294(__obf_53ac9a541c64b86f *__obf_a5d0662d16479f1d) {
	for __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 > 0 && __obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047-1] == '0' {
		__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047--
	}
	if __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 == 0 {
		__obf_53ac9a541c64b86f.

			// Assign v to a.
			__obf_4e162ef443f77409 = 0
	}
}

func (__obf_53ac9a541c64b86f *__obf_a5d0662d16479f1d) Assign(__obf_3d81c055a6608685 uint64) {
	var __obf_7f731835d5cdcc94 [24]byte
	__obf_5618ed758d2ff827 := // Write reversed decimal in buf.
		0
	for __obf_3d81c055a6608685 > 0 {
		__obf_1f27c79843b89c11 := __obf_3d81c055a6608685 / 10
		__obf_3d81c055a6608685 -= 10 * __obf_1f27c79843b89c11
		__obf_7f731835d5cdcc94[__obf_5618ed758d2ff827] = byte(__obf_3d81c055a6608685 + '0')
		__obf_5618ed758d2ff827++
		__obf_3d81c055a6608685 = __obf_1f27c79843b89c11
	}
	__obf_53ac9a541c64b86f.

		// Reverse again to produce forward decimal in a.d.
		__obf_d5ebe437c13ed047 = 0
	for __obf_5618ed758d2ff827--; __obf_5618ed758d2ff827 >= 0; __obf_5618ed758d2ff827-- {
		__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047] = __obf_7f731835d5cdcc94[__obf_5618ed758d2ff827]
		__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047++
	}
	__obf_53ac9a541c64b86f.__obf_4e162ef443f77409 = __obf_53ac9a541c64b86f.

		// Maximum shift that we can do in one pass without overflow.
		// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
		__obf_d5ebe437c13ed047
	__obf_1db2cd3f69354294(__obf_53ac9a541c64b86f)
}

const __obf_4b057de50c3a38f0 = 32 << (^uint(0) >> 63)
const __obf_27869bee05058955 = __obf_4b057de50c3a38f0 - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_984d9d20ea260dae(__obf_53ac9a541c64b86f *__obf_a5d0662d16479f1d, __obf_c0abc1dc124d3675 uint) {
	__obf_1a925cec70c560aa := 0
	__obf_b3105e0528f6edc5 := // read pointer
		0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_5618ed758d2ff827 uint
	for ; __obf_5618ed758d2ff827>>__obf_c0abc1dc124d3675 == 0; __obf_1a925cec70c560aa++ {
		if __obf_1a925cec70c560aa >= __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 {
			if __obf_5618ed758d2ff827 == 0 {
				__obf_53ac9a541c64b86f.
					// a == 0; shouldn't get here, but handle anyway.
					__obf_d5ebe437c13ed047 = 0
				return
			}
			for __obf_5618ed758d2ff827>>__obf_c0abc1dc124d3675 == 0 {
				__obf_5618ed758d2ff827 = __obf_5618ed758d2ff827 * 10
				__obf_1a925cec70c560aa++
			}
			break
		}
		__obf_307e3ae1887fcdd1 := uint(__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_1a925cec70c560aa])
		__obf_5618ed758d2ff827 = __obf_5618ed758d2ff827*10 + __obf_307e3ae1887fcdd1 - '0'
	}
	__obf_53ac9a541c64b86f.__obf_4e162ef443f77409 -= __obf_1a925cec70c560aa - 1

	var __obf_3dcc5d272c13b411 uint = (1 << __obf_c0abc1dc124d3675) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_1a925cec70c560aa < __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047; __obf_1a925cec70c560aa++ {
		__obf_307e3ae1887fcdd1 := uint(__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_1a925cec70c560aa])
		__obf_f0374d0c797e3363 := __obf_5618ed758d2ff827 >> __obf_c0abc1dc124d3675
		__obf_5618ed758d2ff827 &= __obf_3dcc5d272c13b411
		__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_b3105e0528f6edc5] = byte(__obf_f0374d0c797e3363 + '0')
		__obf_b3105e0528f6edc5++
		__obf_5618ed758d2ff827 = __obf_5618ed758d2ff827*10 + __obf_307e3ae1887fcdd1 - '0'
	}

	// Put down extra digits.
	for __obf_5618ed758d2ff827 > 0 {
		__obf_f0374d0c797e3363 := __obf_5618ed758d2ff827 >> __obf_c0abc1dc124d3675
		__obf_5618ed758d2ff827 &= __obf_3dcc5d272c13b411
		if __obf_b3105e0528f6edc5 < len(__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03) {
			__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_b3105e0528f6edc5] = byte(__obf_f0374d0c797e3363 + '0')
			__obf_b3105e0528f6edc5++
		} else if __obf_f0374d0c797e3363 > 0 {
			__obf_53ac9a541c64b86f.__obf_bd845c2e13a22c30 = true
		}
		__obf_5618ed758d2ff827 = __obf_5618ed758d2ff827 * 10
	}
	__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 = __obf_b3105e0528f6edc5

	// Cheat sheet for left shift: table indexed by shift count giving
	// number of new digits that will be introduced by that shift.
	//
	// For example, leftcheats[4] = {2, "625"}.  That means that
	// if we are shifting by 4 (multiplying by 16), it will add 2 digits
	// when the string prefix is "625" through "999", and one fewer digit
	// if the string prefix is "000" through "624".
	//
	// Credit for this trick goes to Ken.
	__obf_1db2cd3f69354294(__obf_53ac9a541c64b86f)
}

type __obf_03e54402a33a9325 struct {
	__obf_40ab3ce93e682d6e int
	__obf_ecad1568f60514f6 string // number of new digits
	// minus one digit if original < a.
}

var __obf_ca0894ee1768183a = []__obf_03e54402a33a9325{
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
func __obf_faefe470cd0b13b9(__obf_edabb4f52a6b6a1f []byte, __obf_f484e81814ac5ff0 string) bool {
	for __obf_283995368e36d3d0 := 0; __obf_283995368e36d3d0 < len(__obf_f484e81814ac5ff0); __obf_283995368e36d3d0++ {
		if __obf_283995368e36d3d0 >= len(__obf_edabb4f52a6b6a1f) {
			return true
		}
		if __obf_edabb4f52a6b6a1f[__obf_283995368e36d3d0] != __obf_f484e81814ac5ff0[__obf_283995368e36d3d0] {
			return __obf_edabb4f52a6b6a1f[__obf_283995368e36d3d0] < __obf_f484e81814ac5ff0[__obf_283995368e36d3d0]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_aa1e7945f76d9f9e(__obf_53ac9a541c64b86f *__obf_a5d0662d16479f1d, __obf_c0abc1dc124d3675 uint) {
	__obf_40ab3ce93e682d6e := __obf_ca0894ee1768183a[__obf_c0abc1dc124d3675].__obf_40ab3ce93e682d6e
	if __obf_faefe470cd0b13b9(__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[0:__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047], __obf_ca0894ee1768183a[__obf_c0abc1dc124d3675].__obf_ecad1568f60514f6) {
		__obf_40ab3ce93e682d6e--
	}
	__obf_1a925cec70c560aa := __obf_53ac9a541c64b86f. // read index
								__obf_d5ebe437c13ed047
	__obf_b3105e0528f6edc5 := __obf_53ac9a541c64b86f. // write index
								__obf_d5ebe437c13ed047 + __obf_40ab3ce93e682d6e

	// Pick up a digit, put down a digit.
	var __obf_5618ed758d2ff827 uint
	for __obf_1a925cec70c560aa--; __obf_1a925cec70c560aa >= 0; __obf_1a925cec70c560aa-- {
		__obf_5618ed758d2ff827 += (uint(__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_1a925cec70c560aa]) - '0') << __obf_c0abc1dc124d3675
		__obf_f2bb065cdbfca60d := __obf_5618ed758d2ff827 / 10
		__obf_0a6e50fc50122fc9 := __obf_5618ed758d2ff827 - 10*__obf_f2bb065cdbfca60d
		__obf_b3105e0528f6edc5--
		if __obf_b3105e0528f6edc5 < len(__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03) {
			__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_b3105e0528f6edc5] = byte(__obf_0a6e50fc50122fc9 + '0')
		} else if __obf_0a6e50fc50122fc9 != 0 {
			__obf_53ac9a541c64b86f.__obf_bd845c2e13a22c30 = true
		}
		__obf_5618ed758d2ff827 = __obf_f2bb065cdbfca60d
	}

	// Put down extra digits.
	for __obf_5618ed758d2ff827 > 0 {
		__obf_f2bb065cdbfca60d := __obf_5618ed758d2ff827 / 10
		__obf_0a6e50fc50122fc9 := __obf_5618ed758d2ff827 - 10*__obf_f2bb065cdbfca60d
		__obf_b3105e0528f6edc5--
		if __obf_b3105e0528f6edc5 < len(__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03) {
			__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_b3105e0528f6edc5] = byte(__obf_0a6e50fc50122fc9 + '0')
		} else if __obf_0a6e50fc50122fc9 != 0 {
			__obf_53ac9a541c64b86f.__obf_bd845c2e13a22c30 = true
		}
		__obf_5618ed758d2ff827 = __obf_f2bb065cdbfca60d
	}
	__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 += __obf_40ab3ce93e682d6e
	if __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 >= len(__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03) {
		__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 = len(__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03)
	}
	__obf_53ac9a541c64b86f.__obf_4e162ef443f77409 += __obf_40ab3ce93e682d6e

	// Binary shift left (k > 0) or right (k < 0).
	__obf_1db2cd3f69354294(__obf_53ac9a541c64b86f)
}

func (__obf_53ac9a541c64b86f *__obf_a5d0662d16479f1d) Shift(__obf_c0abc1dc124d3675 int) {
	switch {
	case __obf_53ac9a541c64b86f.
		// nothing to do: a == 0
		__obf_d5ebe437c13ed047 == 0:

	case __obf_c0abc1dc124d3675 > 0:
		for __obf_c0abc1dc124d3675 > __obf_27869bee05058955 {
			__obf_aa1e7945f76d9f9e(__obf_53ac9a541c64b86f, __obf_27869bee05058955)
			__obf_c0abc1dc124d3675 -= __obf_27869bee05058955
		}
		__obf_aa1e7945f76d9f9e(__obf_53ac9a541c64b86f, uint(__obf_c0abc1dc124d3675))
	case __obf_c0abc1dc124d3675 < 0:
		for __obf_c0abc1dc124d3675 < -__obf_27869bee05058955 {
			__obf_984d9d20ea260dae(__obf_53ac9a541c64b86f, __obf_27869bee05058955)
			__obf_c0abc1dc124d3675 += __obf_27869bee05058955
		}
		__obf_984d9d20ea260dae(__obf_53ac9a541c64b86f, uint(-__obf_c0abc1dc124d3675))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_ea591fe8aac22108(__obf_53ac9a541c64b86f *__obf_a5d0662d16479f1d, __obf_d5ebe437c13ed047 int) bool {
	if __obf_d5ebe437c13ed047 < 0 || __obf_d5ebe437c13ed047 >= __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 {
		return false
	}
	if __obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_d5ebe437c13ed047] == '5' && __obf_d5ebe437c13ed047+1 == __obf_53ac9a541c64b86f. // exactly halfway - round to even
																		__obf_d5ebe437c13ed047 {
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_53ac9a541c64b86f.__obf_bd845c2e13a22c30 {
			return true
		}
		return __obf_d5ebe437c13ed047 > 0 && (__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_d5ebe437c13ed047-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_53ac9a541c64b86f.

		// Round a to nd digits (or fewer).
		// If nd is zero, it means we're rounding
		// just to the left of the digits, as in
		// 0.09 -> 0.1.
		__obf_5debc2a983e87a03[__obf_d5ebe437c13ed047] >= '5'
}

func (__obf_53ac9a541c64b86f *__obf_a5d0662d16479f1d) Round(__obf_d5ebe437c13ed047 int) {
	if __obf_d5ebe437c13ed047 < 0 || __obf_d5ebe437c13ed047 >= __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 {
		return
	}
	if __obf_ea591fe8aac22108(__obf_53ac9a541c64b86f, __obf_d5ebe437c13ed047) {
		__obf_53ac9a541c64b86f.
			RoundUp(__obf_d5ebe437c13ed047)
	} else {
		__obf_53ac9a541c64b86f.
			RoundDown(__obf_d5ebe437c13ed047)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_53ac9a541c64b86f *__obf_a5d0662d16479f1d) RoundDown(__obf_d5ebe437c13ed047 int) {
	if __obf_d5ebe437c13ed047 < 0 || __obf_d5ebe437c13ed047 >= __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 {
		return
	}
	__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 = __obf_d5ebe437c13ed047

	// Round a up to nd digits (or fewer).
	__obf_1db2cd3f69354294(__obf_53ac9a541c64b86f)
}

func (__obf_53ac9a541c64b86f *__obf_a5d0662d16479f1d) RoundUp(__obf_d5ebe437c13ed047 int) {
	if __obf_d5ebe437c13ed047 < 0 || __obf_d5ebe437c13ed047 >= __obf_53ac9a541c64b86f.

		// round up
		__obf_d5ebe437c13ed047 {
		return
	}

	for __obf_283995368e36d3d0 := __obf_d5ebe437c13ed047 - 1; __obf_283995368e36d3d0 >= 0; __obf_283995368e36d3d0-- {
		__obf_307e3ae1887fcdd1 := __obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_283995368e36d3d0]
		if __obf_307e3ae1887fcdd1 < '9' {
			__obf_53ac9a541c64b86f. // can stop after this digit
						__obf_5debc2a983e87a03[__obf_283995368e36d3d0]++
			__obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047 = __obf_283995368e36d3d0 + 1
			return
		}
	}
	__obf_53ac9a541c64b86f.

		// Number is all 9s.
		// Change to single 1 with adjusted decimal point.
		__obf_5debc2a983e87a03[0] = '1'
	__obf_53ac9a541c64b86f.

		// Extract integer part, rounded appropriately.
		// No guarantees about overflow.
		__obf_d5ebe437c13ed047 = 1
	__obf_53ac9a541c64b86f.__obf_4e162ef443f77409++
}

func (__obf_53ac9a541c64b86f *__obf_a5d0662d16479f1d) RoundedInteger() uint64 {
	if __obf_53ac9a541c64b86f.__obf_4e162ef443f77409 > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_283995368e36d3d0 int
	__obf_5618ed758d2ff827 := uint64(0)
	for __obf_283995368e36d3d0 = 0; __obf_283995368e36d3d0 < __obf_53ac9a541c64b86f.__obf_4e162ef443f77409 && __obf_283995368e36d3d0 < __obf_53ac9a541c64b86f.__obf_d5ebe437c13ed047; __obf_283995368e36d3d0++ {
		__obf_5618ed758d2ff827 = __obf_5618ed758d2ff827*10 + uint64(__obf_53ac9a541c64b86f.__obf_5debc2a983e87a03[__obf_283995368e36d3d0]-'0')
	}
	for ; __obf_283995368e36d3d0 < __obf_53ac9a541c64b86f.__obf_4e162ef443f77409; __obf_283995368e36d3d0++ {
		__obf_5618ed758d2ff827 *= 10
	}
	if __obf_ea591fe8aac22108(__obf_53ac9a541c64b86f, __obf_53ac9a541c64b86f.__obf_4e162ef443f77409) {
		__obf_5618ed758d2ff827++
	}
	return __obf_5618ed758d2ff827
}
