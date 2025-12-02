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
package __obf_bd46136174bdabe4

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

var __obf_1573d6eb6e796b01 = big.NewInt(0)
var __obf_b1e5dea21e0aeea5 = big.NewInt(1)
var __obf_53e3975b80a15b69 = big.NewInt(2)
var __obf_1f236215c06e3969 = big.NewInt(4)
var __obf_953cc735cf320980 = big.NewInt(5)
var __obf_263daee51f205bb2 = big.NewInt(10)
var __obf_58870f8992c18ef8 = big.NewInt(20)

var __obf_26acda08e4a64645 = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_4c2240e6f10eb22b *big.Int
	__obf_3ab33991ba7f9863 int32 // NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.

}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_4c2240e6f10eb22b int64, __obf_3ab33991ba7f9863 int32) Decimal {
	return Decimal{__obf_4c2240e6f10eb22b: big.NewInt(__obf_4c2240e6f10eb22b), __obf_3ab33991ba7f9863: __obf_3ab33991ba7f9863}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_4c2240e6f10eb22b int64) Decimal {
	return Decimal{__obf_4c2240e6f10eb22b: big.NewInt(__obf_4c2240e6f10eb22b), __obf_3ab33991ba7f9863: 0}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_4c2240e6f10eb22b int32) Decimal {
	return Decimal{__obf_4c2240e6f10eb22b: big.NewInt(int64(__obf_4c2240e6f10eb22b)), __obf_3ab33991ba7f9863: 0}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_4c2240e6f10eb22b uint64) Decimal {
	return Decimal{__obf_4c2240e6f10eb22b: new(big.Int).SetUint64(__obf_4c2240e6f10eb22b), __obf_3ab33991ba7f9863: 0}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_4c2240e6f10eb22b *big.Int, __obf_3ab33991ba7f9863 int32) Decimal {
	return Decimal{__obf_4c2240e6f10eb22b: new(big.Int).Set(__obf_4c2240e6f10eb22b), __obf_3ab33991ba7f9863: __obf_3ab33991ba7f9863}
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
func NewFromBigRat(__obf_4c2240e6f10eb22b *big.Rat, __obf_d55e3219d646b95d int32) Decimal {
	return Decimal{__obf_4c2240e6f10eb22b: new(big.Int).Set(__obf_4c2240e6f10eb22b.Num()), __obf_3ab33991ba7f9863: 0}.DivRound(Decimal{__obf_4c2240e6f10eb22b: new(big.Int).Set(__obf_4c2240e6f10eb22b.Denom()), __obf_3ab33991ba7f9863: 0}, __obf_d55e3219d646b95d)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_4c2240e6f10eb22b string) (Decimal, error) {
	__obf_b898e12f95d00f3f := __obf_4c2240e6f10eb22b
	var __obf_9dc2ddb2a2de52f6 string
	var __obf_3ab33991ba7f9863 int64
	__obf_cd6ea44df85ac387 := // Check if number is using scientific notation
		strings.IndexAny(__obf_4c2240e6f10eb22b, "Ee")
	if __obf_cd6ea44df85ac387 != -1 {
		__obf_1a5f585db9dc3920, __obf_1cf2c004a31cf720 := strconv.ParseInt(__obf_4c2240e6f10eb22b[__obf_cd6ea44df85ac387+1:], 10, 32)
		if __obf_1cf2c004a31cf720 != nil {
			if __obf_98c5f7349b099b51, __obf_58c4d7e4d27a6363 := __obf_1cf2c004a31cf720.(*strconv.NumError); __obf_58c4d7e4d27a6363 && __obf_98c5f7349b099b51.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_4c2240e6f10eb22b)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_4c2240e6f10eb22b)
		}
		__obf_4c2240e6f10eb22b = __obf_4c2240e6f10eb22b[:__obf_cd6ea44df85ac387]
		__obf_3ab33991ba7f9863 = __obf_1a5f585db9dc3920
	}
	__obf_8f3e1fd01893a8c9 := -1
	__obf_e46a04f647f87cd6 := len(__obf_4c2240e6f10eb22b)
	for __obf_fb920e988f68d242 := 0; __obf_fb920e988f68d242 < __obf_e46a04f647f87cd6; __obf_fb920e988f68d242++ {
		if __obf_4c2240e6f10eb22b[__obf_fb920e988f68d242] == '.' {
			if __obf_8f3e1fd01893a8c9 > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_4c2240e6f10eb22b)
			}
			__obf_8f3e1fd01893a8c9 = __obf_fb920e988f68d242
		}
	}

	if __obf_8f3e1fd01893a8c9 == -1 {
		__obf_9dc2ddb2a2de52f6 = // There is no decimal point, we can just parse the original string as
			// an int
			__obf_4c2240e6f10eb22b
	} else {
		if __obf_8f3e1fd01893a8c9+1 < __obf_e46a04f647f87cd6 {
			__obf_9dc2ddb2a2de52f6 = __obf_4c2240e6f10eb22b[:__obf_8f3e1fd01893a8c9] + __obf_4c2240e6f10eb22b[__obf_8f3e1fd01893a8c9+1:]
		} else {
			__obf_9dc2ddb2a2de52f6 = __obf_4c2240e6f10eb22b[:__obf_8f3e1fd01893a8c9]
		}
		__obf_1a5f585db9dc3920 := -len(__obf_4c2240e6f10eb22b[__obf_8f3e1fd01893a8c9+1:])
		__obf_3ab33991ba7f9863 += int64(__obf_1a5f585db9dc3920)
	}

	var __obf_7a3affb432853d53 *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_9dc2ddb2a2de52f6) <= 18 {
		__obf_5ac21bd6796d8805, __obf_1cf2c004a31cf720 := strconv.ParseInt(__obf_9dc2ddb2a2de52f6, 10, 64)
		if __obf_1cf2c004a31cf720 != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_4c2240e6f10eb22b)
		}
		__obf_7a3affb432853d53 = big.NewInt(__obf_5ac21bd6796d8805)
	} else {
		__obf_7a3affb432853d53 = new(big.Int)
		_, __obf_58c4d7e4d27a6363 := __obf_7a3affb432853d53.SetString(__obf_9dc2ddb2a2de52f6, 10)
		if !__obf_58c4d7e4d27a6363 {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_4c2240e6f10eb22b)
		}
	}

	if __obf_3ab33991ba7f9863 < math.MinInt32 || __obf_3ab33991ba7f9863 > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_b898e12f95d00f3f)
	}

	return Decimal{__obf_4c2240e6f10eb22b: __obf_7a3affb432853d53, __obf_3ab33991ba7f9863: int32(__obf_3ab33991ba7f9863)}, nil
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
func NewFromFormattedString(__obf_4c2240e6f10eb22b string, __obf_1d6a3fdc6f2210e6 *regexp.Regexp) (Decimal, error) {
	__obf_6588b739932f0d94 := __obf_1d6a3fdc6f2210e6.ReplaceAllString(__obf_4c2240e6f10eb22b, "")
	__obf_a1ffba48568e5107, __obf_1cf2c004a31cf720 := NewFromString(__obf_6588b739932f0d94)
	if __obf_1cf2c004a31cf720 != nil {
		return Decimal{}, __obf_1cf2c004a31cf720
	}
	return __obf_a1ffba48568e5107, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_4c2240e6f10eb22b string) Decimal {
	__obf_cad8fb067a1dbc4b, __obf_1cf2c004a31cf720 := NewFromString(__obf_4c2240e6f10eb22b)
	if __obf_1cf2c004a31cf720 != nil {
		panic(__obf_1cf2c004a31cf720)
	}
	return __obf_cad8fb067a1dbc4b
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
func NewFromFloat(__obf_4c2240e6f10eb22b float64) Decimal {
	if __obf_4c2240e6f10eb22b == 0 {
		return New(0, 0)
	}
	return __obf_71520ad26415235e(__obf_4c2240e6f10eb22b, math.Float64bits(__obf_4c2240e6f10eb22b), &__obf_5c7b21ace088503b)
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
func NewFromFloat32(__obf_4c2240e6f10eb22b float32) Decimal {
	if __obf_4c2240e6f10eb22b == 0 {
		return New(0, 0)
	}
	__obf_97404a94f9526684 := // XOR is workaround for https://github.com/golang/go/issues/26285
		math.Float32bits(__obf_4c2240e6f10eb22b) ^ 0x80808080
	return __obf_71520ad26415235e(float64(__obf_4c2240e6f10eb22b), uint64(__obf_97404a94f9526684)^0x80808080, &__obf_8177a6f362debfa1)
}

func __obf_71520ad26415235e(__obf_2ab63d30f41dca35 float64, __obf_588682ab1841f9f3 uint64, __obf_db4d85e8d4b69c3f *__obf_3510ef834411b273) Decimal {
	if math.IsNaN(__obf_2ab63d30f41dca35) || math.IsInf(__obf_2ab63d30f41dca35, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_2ab63d30f41dca35))
	}
	__obf_3ab33991ba7f9863 := int(__obf_588682ab1841f9f3>>__obf_db4d85e8d4b69c3f.__obf_350a4b63232b6662) & (1<<__obf_db4d85e8d4b69c3f.__obf_33dfabedaba76fac - 1)
	__obf_a1a9265f031cbbf2 := __obf_588682ab1841f9f3 & (uint64(1)<<__obf_db4d85e8d4b69c3f.__obf_350a4b63232b6662 - 1)

	switch __obf_3ab33991ba7f9863 {
	case 0:
		__obf_3ab33991ba7f9863++ // denormalized

	default:
		__obf_a1a9265f031cbbf2 |= // add implicit top bit
			uint64(1) << __obf_db4d85e8d4b69c3f.__obf_350a4b63232b6662
	}
	__obf_3ab33991ba7f9863 += __obf_db4d85e8d4b69c3f.__obf_d3ccd221dc097eda

	var __obf_a1ffba48568e5107 __obf_bd46136174bdabe4
	__obf_a1ffba48568e5107.
		Assign(__obf_a1a9265f031cbbf2)
	__obf_a1ffba48568e5107.
		Shift(__obf_3ab33991ba7f9863 - int(__obf_db4d85e8d4b69c3f.__obf_350a4b63232b6662))
	__obf_a1ffba48568e5107.__obf_c5a0f28fc9f4a6ce = __obf_588682ab1841f9f3>>(__obf_db4d85e8d4b69c3f.__obf_33dfabedaba76fac+__obf_db4d85e8d4b69c3f.__obf_350a4b63232b6662) != 0
	__obf_a3ec81b881a7faf1(&__obf_a1ffba48568e5107,
		// If less than 19 digits, we can do calculation in an int64.
		__obf_a1a9265f031cbbf2, __obf_3ab33991ba7f9863, __obf_db4d85e8d4b69c3f)

	if __obf_a1ffba48568e5107.__obf_3018d9650ee6ca39 < 19 {
		__obf_b79982d611e468ff := int64(0)
		__obf_a49032a74f3c7436 := int64(1)
		for __obf_fb920e988f68d242 := __obf_a1ffba48568e5107.__obf_3018d9650ee6ca39 - 1; __obf_fb920e988f68d242 >= 0; __obf_fb920e988f68d242-- {
			__obf_b79982d611e468ff += __obf_a49032a74f3c7436 * int64(__obf_a1ffba48568e5107.__obf_a1ffba48568e5107[__obf_fb920e988f68d242]-'0')
			__obf_a49032a74f3c7436 *= 10
		}
		if __obf_a1ffba48568e5107.__obf_c5a0f28fc9f4a6ce {
			__obf_b79982d611e468ff *= -1
		}
		return Decimal{__obf_4c2240e6f10eb22b: big.NewInt(__obf_b79982d611e468ff), __obf_3ab33991ba7f9863: int32(__obf_a1ffba48568e5107.__obf_c3fb49899f9ddeff) - int32(__obf_a1ffba48568e5107.__obf_3018d9650ee6ca39)}
	}
	__obf_7a3affb432853d53 := new(big.Int)
	__obf_7a3affb432853d53, __obf_58c4d7e4d27a6363 := __obf_7a3affb432853d53.SetString(string(__obf_a1ffba48568e5107.__obf_a1ffba48568e5107[:__obf_a1ffba48568e5107.__obf_3018d9650ee6ca39]), 10)
	if __obf_58c4d7e4d27a6363 {
		return Decimal{__obf_4c2240e6f10eb22b: __obf_7a3affb432853d53, __obf_3ab33991ba7f9863: int32(__obf_a1ffba48568e5107.__obf_c3fb49899f9ddeff) - int32(__obf_a1ffba48568e5107.__obf_3018d9650ee6ca39)}
	}

	return NewFromFloatWithExponent(__obf_2ab63d30f41dca35, int32(__obf_a1ffba48568e5107.

		// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
		// number of fractional digits.
		//
		// Example:
		//
		//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
		__obf_c3fb49899f9ddeff)-int32(__obf_a1ffba48568e5107.__obf_3018d9650ee6ca39))
}

func NewFromFloatWithExponent(__obf_4c2240e6f10eb22b float64, __obf_3ab33991ba7f9863 int32) Decimal {
	if math.IsNaN(__obf_4c2240e6f10eb22b) || math.IsInf(__obf_4c2240e6f10eb22b, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_4c2240e6f10eb22b))
	}
	__obf_588682ab1841f9f3 := math.Float64bits(__obf_4c2240e6f10eb22b)
	__obf_a1a9265f031cbbf2 := __obf_588682ab1841f9f3 & (1<<52 - 1)
	__obf_01bf3185de9e7e78 := int32((__obf_588682ab1841f9f3 >> 52) & (1<<11 - 1))
	__obf_6dbf9797d792aa4a := __obf_588682ab1841f9f3 >> 63

	if __obf_01bf3185de9e7e78 == 0 {
		// specials
		if __obf_a1a9265f031cbbf2 == 0 {
			return Decimal{}
		}
		__obf_01bf3185de9e7e78++ // subnormal

	} else {
		__obf_a1a9265f031cbbf2 |= // normal
			1 << 52
	}
	__obf_01bf3185de9e7e78 -= 1023 + 52

	// normalizing base-2 values
	for __obf_a1a9265f031cbbf2&1 == 0 {
		__obf_a1a9265f031cbbf2 = __obf_a1a9265f031cbbf2 >> 1
		__obf_01bf3185de9e7e78++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_3ab33991ba7f9863 < 0 && __obf_3ab33991ba7f9863 < __obf_01bf3185de9e7e78 {
		if __obf_01bf3185de9e7e78 < 0 {
			__obf_3ab33991ba7f9863 = __obf_01bf3185de9e7e78
		} else {
			__obf_3ab33991ba7f9863 = 0
		}
	}
	__obf_01bf3185de9e7e78 -= // representing 10^M * 2^N as 5^M * 2^(M+N)
		__obf_3ab33991ba7f9863
	__obf_39f57644ba1d6664 := big.NewInt(1)
	__obf_4400f577ffe4fca5 := big.NewInt(int64(__obf_a1a9265f031cbbf2))

	// applying 5^M
	if __obf_3ab33991ba7f9863 > 0 {
		__obf_39f57644ba1d6664 = __obf_39f57644ba1d6664.SetInt64(int64(__obf_3ab33991ba7f9863))
		__obf_39f57644ba1d6664 = __obf_39f57644ba1d6664.Exp(__obf_953cc735cf320980, __obf_39f57644ba1d6664, nil)
	} else if __obf_3ab33991ba7f9863 < 0 {
		__obf_39f57644ba1d6664 = __obf_39f57644ba1d6664.SetInt64(-int64(__obf_3ab33991ba7f9863))
		__obf_39f57644ba1d6664 = __obf_39f57644ba1d6664.Exp(__obf_953cc735cf320980, __obf_39f57644ba1d6664, nil)
		__obf_4400f577ffe4fca5 = __obf_4400f577ffe4fca5.Mul(__obf_4400f577ffe4fca5, __obf_39f57644ba1d6664)
		__obf_39f57644ba1d6664 = __obf_39f57644ba1d6664.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_01bf3185de9e7e78 > 0 {
		__obf_4400f577ffe4fca5 = __obf_4400f577ffe4fca5.Lsh(__obf_4400f577ffe4fca5, uint(__obf_01bf3185de9e7e78))
	} else if __obf_01bf3185de9e7e78 < 0 {
		__obf_39f57644ba1d6664 = __obf_39f57644ba1d6664.Lsh(__obf_39f57644ba1d6664, uint(-__obf_01bf3185de9e7e78))
	}

	// rounding and downscaling
	if __obf_3ab33991ba7f9863 > 0 || __obf_01bf3185de9e7e78 < 0 {
		__obf_4cab2f8bf2e5ef9b := new(big.Int).Rsh(__obf_39f57644ba1d6664, 1)
		__obf_4400f577ffe4fca5 = __obf_4400f577ffe4fca5.Add(__obf_4400f577ffe4fca5, __obf_4cab2f8bf2e5ef9b)
		__obf_4400f577ffe4fca5 = __obf_4400f577ffe4fca5.Quo(__obf_4400f577ffe4fca5, __obf_39f57644ba1d6664)
	}

	if __obf_6dbf9797d792aa4a == 1 {
		__obf_4400f577ffe4fca5 = __obf_4400f577ffe4fca5.Neg(__obf_4400f577ffe4fca5)
	}

	return Decimal{__obf_4c2240e6f10eb22b: __obf_4400f577ffe4fca5, __obf_3ab33991ba7f9863: __obf_3ab33991ba7f9863}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_a1ffba48568e5107 Decimal) Copy() Decimal {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()
	return Decimal{__obf_4c2240e6f10eb22b: new(big.Int).Set(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b), __obf_3ab33991ba7f9863: __obf_a1ffba48568e5107.

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
		__obf_3ab33991ba7f9863,
	}
}

func (__obf_a1ffba48568e5107 Decimal) __obf_09c56f1962a9d090(__obf_3ab33991ba7f9863 int32) Decimal {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()

	if __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 == __obf_3ab33991ba7f9863 {
		return Decimal{
			new(big.Int).Set(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b), __obf_a1ffba48568e5107.

				// NOTE(vadim): must convert exps to float64 before - to prevent overflow
				__obf_3ab33991ba7f9863,
		}
	}
	__obf_6ecededc473cc856 := math.Abs(float64(__obf_3ab33991ba7f9863) - float64(__obf_a1ffba48568e5107.__obf_3ab33991ba7f9863))
	__obf_4c2240e6f10eb22b := new(big.Int).Set(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b)
	__obf_4b83b5d53742a67e := new(big.Int).Exp(__obf_263daee51f205bb2, big.NewInt(int64(__obf_6ecededc473cc856)), nil)
	if __obf_3ab33991ba7f9863 > __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 {
		__obf_4c2240e6f10eb22b = __obf_4c2240e6f10eb22b.Quo(__obf_4c2240e6f10eb22b, __obf_4b83b5d53742a67e)
	} else if __obf_3ab33991ba7f9863 < __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 {
		__obf_4c2240e6f10eb22b = __obf_4c2240e6f10eb22b.Mul(__obf_4c2240e6f10eb22b, __obf_4b83b5d53742a67e)
	}

	return Decimal{__obf_4c2240e6f10eb22b: __obf_4c2240e6f10eb22b, __obf_3ab33991ba7f9863: __obf_3ab33991ba7f9863}
}

// Abs returns the absolute value of the decimal.
func (__obf_a1ffba48568e5107 Decimal) Abs() Decimal {
	if !__obf_a1ffba48568e5107.IsNegative() {
		return __obf_a1ffba48568e5107
	}
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()
	__obf_7c3715e5cc3d86b2 := new(big.Int).Abs(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b)
	return Decimal{__obf_4c2240e6f10eb22b: __obf_7c3715e5cc3d86b2, __obf_3ab33991ba7f9863: __obf_a1ffba48568e5107.

		// Add returns d + d2.
		__obf_3ab33991ba7f9863,
	}
}

func (__obf_a1ffba48568e5107 Decimal) Add(__obf_83d5b629b20f9f91 Decimal) Decimal {
	__obf_9fdd40229da129a1, __obf_f88a1c7011d37ae3 := RescalePair(__obf_a1ffba48568e5107, __obf_83d5b629b20f9f91)
	__obf_1ce7f065c84b27aa := new(big.Int).Add(__obf_9fdd40229da129a1.__obf_4c2240e6f10eb22b, __obf_f88a1c7011d37ae3.__obf_4c2240e6f10eb22b)
	return Decimal{__obf_4c2240e6f10eb22b: __obf_1ce7f065c84b27aa, __obf_3ab33991ba7f9863: __obf_9fdd40229da129a1.

		// Sub returns d - d2.
		__obf_3ab33991ba7f9863,
	}
}

func (__obf_a1ffba48568e5107 Decimal) Sub(__obf_83d5b629b20f9f91 Decimal) Decimal {
	__obf_9fdd40229da129a1, __obf_f88a1c7011d37ae3 := RescalePair(__obf_a1ffba48568e5107, __obf_83d5b629b20f9f91)
	__obf_1ce7f065c84b27aa := new(big.Int).Sub(__obf_9fdd40229da129a1.__obf_4c2240e6f10eb22b, __obf_f88a1c7011d37ae3.__obf_4c2240e6f10eb22b)
	return Decimal{__obf_4c2240e6f10eb22b: __obf_1ce7f065c84b27aa, __obf_3ab33991ba7f9863: __obf_9fdd40229da129a1.

		// Neg returns -d.
		__obf_3ab33991ba7f9863,
	}
}

func (__obf_a1ffba48568e5107 Decimal) Neg() Decimal {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()
	__obf_2ab63d30f41dca35 := new(big.Int).Neg(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b)
	return Decimal{__obf_4c2240e6f10eb22b: __obf_2ab63d30f41dca35, __obf_3ab33991ba7f9863: __obf_a1ffba48568e5107.

		// Mul returns d * d2.
		__obf_3ab33991ba7f9863,
	}
}

func (__obf_a1ffba48568e5107 Decimal) Mul(__obf_83d5b629b20f9f91 Decimal) Decimal {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()
	__obf_83d5b629b20f9f91.__obf_aa61179208fecfc6()
	__obf_120fc3446083fff7 := int64(__obf_a1ffba48568e5107.__obf_3ab33991ba7f9863) + int64(__obf_83d5b629b20f9f91.__obf_3ab33991ba7f9863)
	if __obf_120fc3446083fff7 > math.MaxInt32 || __obf_120fc3446083fff7 < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_120fc3446083fff7))
	}
	__obf_1ce7f065c84b27aa := new(big.Int).Mul(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b, __obf_83d5b629b20f9f91.__obf_4c2240e6f10eb22b)
	return Decimal{__obf_4c2240e6f10eb22b: __obf_1ce7f065c84b27aa, __obf_3ab33991ba7f9863: int32(__obf_120fc3446083fff7)}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_a1ffba48568e5107 Decimal) Shift(__obf_5dd4f6515e75ae70 int32) Decimal {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()
	return Decimal{__obf_4c2240e6f10eb22b: new(big.Int).Set(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b), __obf_3ab33991ba7f9863: __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 + __obf_5dd4f6515e75ae70}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_a1ffba48568e5107 Decimal) Div(__obf_83d5b629b20f9f91 Decimal) Decimal {
	return __obf_a1ffba48568e5107.DivRound(__obf_83d5b629b20f9f91, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_a1ffba48568e5107 Decimal) QuoRem(__obf_83d5b629b20f9f91 Decimal, __obf_d55e3219d646b95d int32) (Decimal, Decimal) {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()
	__obf_83d5b629b20f9f91.__obf_aa61179208fecfc6()
	if __obf_83d5b629b20f9f91.__obf_4c2240e6f10eb22b.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_de5937f4c05b37b2 := -__obf_d55e3219d646b95d
	__obf_98c5f7349b099b51 := int64(__obf_a1ffba48568e5107.__obf_3ab33991ba7f9863) - int64(__obf_83d5b629b20f9f91.__obf_3ab33991ba7f9863) - int64(__obf_de5937f4c05b37b2)
	if __obf_98c5f7349b099b51 > math.MaxInt32 || __obf_98c5f7349b099b51 < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_09d0a7d9a079073e, __obf_6387dcab62f31333,

		// d = a 10^ea
		// d2 = b 10^eb
		__obf_b71b7afef07bde96 big.Int
	var __obf_739c22215baa620b int32

	if __obf_98c5f7349b099b51 < 0 {
		__obf_09d0a7d9a079073e = *__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b
		__obf_b71b7afef07bde96.
			SetInt64(-__obf_98c5f7349b099b51)
		__obf_6387dcab62f31333.
			Exp(__obf_263daee51f205bb2, &__obf_b71b7afef07bde96, nil)
		__obf_6387dcab62f31333.
			Mul(__obf_83d5b629b20f9f91.__obf_4c2240e6f10eb22b, &__obf_6387dcab62f31333)
		__obf_739c22215baa620b = __obf_a1ffba48568e5107.
			// now aa = a
			//     bb = b 10^(scale + eb - ea)
			__obf_3ab33991ba7f9863

	} else {
		__obf_b71b7afef07bde96.
			SetInt64(__obf_98c5f7349b099b51)
		__obf_09d0a7d9a079073e.
			Exp(__obf_263daee51f205bb2, &__obf_b71b7afef07bde96, nil)
		__obf_09d0a7d9a079073e.
			Mul(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b, &__obf_09d0a7d9a079073e)
		__obf_6387dcab62f31333 = *__obf_83d5b629b20f9f91.__obf_4c2240e6f10eb22b

		// now aa = a ^ (ea - eb - scale)
		//     bb = b
		__obf_739c22215baa620b = __obf_de5937f4c05b37b2 + __obf_83d5b629b20f9f91.__obf_3ab33991ba7f9863

	}
	var __obf_d77d99c0876d6027, __obf_7314b6cb9c5fa307 big.Int
	__obf_d77d99c0876d6027.
		QuoRem(&__obf_09d0a7d9a079073e, &__obf_6387dcab62f31333, &__obf_7314b6cb9c5fa307)
	__obf_1a4116d826ac1c88 := Decimal{__obf_4c2240e6f10eb22b: &__obf_d77d99c0876d6027, __obf_3ab33991ba7f9863: __obf_de5937f4c05b37b2}
	__obf_a938e0061ce0478a := Decimal{__obf_4c2240e6f10eb22b: &__obf_7314b6cb9c5fa307, __obf_3ab33991ba7f9863: __obf_739c22215baa620b}
	return __obf_1a4116d826ac1c88,

		// DivRound divides and rounds to a given precision
		// i.e. to an integer multiple of 10^(-precision)
		//
		//	for a positive quotient digit 5 is rounded up, away from 0
		//	if the quotient is negative then digit 5 is rounded down, away from 0
		//
		// Note that precision<0 is allowed as input.
		__obf_a938e0061ce0478a
}

func (__obf_a1ffba48568e5107 Decimal) DivRound(__obf_83d5b629b20f9f91 Decimal, __obf_d55e3219d646b95d int32) Decimal {
	__obf_d77d99c0876d6027,
		// QuoRem already checks initialization
		__obf_7314b6cb9c5fa307 := __obf_a1ffba48568e5107.QuoRem(__obf_83d5b629b20f9f91,
		// the actual rounding decision is based on comparing r*10^precision and d2/2
		// instead compare 2 r 10 ^precision and d2
		__obf_d55e3219d646b95d)

	var __obf_65c097e9a1c47439 big.Int
	__obf_65c097e9a1c47439.
		Abs(__obf_7314b6cb9c5fa307.__obf_4c2240e6f10eb22b)
	__obf_65c097e9a1c47439.
		Lsh(&__obf_65c097e9a1c47439, 1)
	__obf_7f4464608d3e98ba := // now rv2 = abs(r.value) * 2
		Decimal{__obf_4c2240e6f10eb22b: &__obf_65c097e9a1c47439, __obf_3ab33991ba7f9863: __obf_7314b6cb9c5fa307.
			// r2 is now 2 * r * 10 ^ precision
			__obf_3ab33991ba7f9863 + __obf_d55e3219d646b95d}

	var __obf_4cc6d31ede5e315a = __obf_7f4464608d3e98ba.Cmp(__obf_83d5b629b20f9f91.Abs())

	if __obf_4cc6d31ede5e315a < 0 {
		return __obf_d77d99c0876d6027
	}

	if __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.Sign()*__obf_83d5b629b20f9f91.__obf_4c2240e6f10eb22b.Sign() < 0 {
		return __obf_d77d99c0876d6027.Sub(New(1, -__obf_d55e3219d646b95d))
	}

	return __obf_d77d99c0876d6027.Add(New(1, -__obf_d55e3219d646b95d))
}

// Mod returns d % d2.
func (__obf_a1ffba48568e5107 Decimal) Mod(__obf_83d5b629b20f9f91 Decimal) Decimal {
	_, __obf_7314b6cb9c5fa307 := __obf_a1ffba48568e5107.QuoRem(__obf_83d5b629b20f9f91, 0)
	return __obf_7314b6cb9c5fa307
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
func (__obf_a1ffba48568e5107 Decimal) Pow(__obf_83d5b629b20f9f91 Decimal) Decimal {
	__obf_03da3ca9b693fc27 := __obf_a1ffba48568e5107.Sign()
	__obf_e80e2d69e3fa8dbf := __obf_83d5b629b20f9f91.Sign()

	if __obf_03da3ca9b693fc27 == 0 {
		if __obf_e80e2d69e3fa8dbf == 0 {
			return Decimal{}
		}
		if __obf_e80e2d69e3fa8dbf == 1 {
			return Decimal{__obf_1573d6eb6e796b01, 0}
		}
		if __obf_e80e2d69e3fa8dbf == -1 {
			return Decimal{}
		}
	}

	if __obf_e80e2d69e3fa8dbf == 0 {
		return Decimal{__obf_b1e5dea21e0aeea5, 0}
	}
	__obf_bb2bbaf5ad72fade := // TODO: optimize extraction of fractional part
		Decimal{__obf_b1e5dea21e0aeea5, 0}
	__obf_0aed4c55bfd02d07, __obf_e556d4c0c2678d61 := __obf_83d5b629b20f9f91.QuoRem(__obf_bb2bbaf5ad72fade, 0)

	if __obf_03da3ca9b693fc27 == -1 && !__obf_e556d4c0c2678d61.IsZero() {
		return Decimal{}
	}
	__obf_af049c7749e14f94, _ := __obf_a1ffba48568e5107.PowBigInt(__obf_0aed4c55bfd02d07.

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_4c2240e6f10eb22b)

	if __obf_e556d4c0c2678d61.__obf_4c2240e6f10eb22b.Sign() == 0 {
		return __obf_af049c7749e14f94
	}
	__obf_c3c6dad7153ba371 := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_a1ffba48568e5107.NumDigits()
	__obf_c354c851af938e33 := __obf_83d5b629b20f9f91.NumDigits()
	__obf_d55e3219d646b95d := __obf_c3c6dad7153ba371

	if __obf_c354c851af938e33 > __obf_d55e3219d646b95d {
		__obf_d55e3219d646b95d += __obf_c354c851af938e33
	}
	__obf_d55e3219d646b95d += 6
	__obf_ebd4a5da1aa116e2,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_1cf2c004a31cf720 := __obf_a1ffba48568e5107.Abs().Ln(-__obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 + int32(__obf_d55e3219d646b95d))
	if __obf_1cf2c004a31cf720 != nil {
		return Decimal{}
	}
	__obf_ebd4a5da1aa116e2 = __obf_ebd4a5da1aa116e2.Mul(__obf_e556d4c0c2678d61)
	__obf_ebd4a5da1aa116e2, __obf_1cf2c004a31cf720 = __obf_ebd4a5da1aa116e2.ExpTaylor(-__obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 + int32(__obf_d55e3219d646b95d))
	if __obf_1cf2c004a31cf720 != nil {
		return Decimal{}
	}
	__obf_d45eb54903cb2ff9 := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_af049c7749e14f94.Mul(__obf_ebd4a5da1aa116e2)

	return __obf_d45eb54903cb2ff9
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
func (__obf_a1ffba48568e5107 Decimal) PowWithPrecision(__obf_83d5b629b20f9f91 Decimal, __obf_d55e3219d646b95d int32) (Decimal, error) {
	__obf_03da3ca9b693fc27 := __obf_a1ffba48568e5107.Sign()
	__obf_e80e2d69e3fa8dbf := __obf_83d5b629b20f9f91.Sign()

	if __obf_03da3ca9b693fc27 == 0 {
		if __obf_e80e2d69e3fa8dbf == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_e80e2d69e3fa8dbf == 1 {
			return Decimal{__obf_1573d6eb6e796b01, 0}, nil
		}
		if __obf_e80e2d69e3fa8dbf == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_e80e2d69e3fa8dbf == 0 {
		return Decimal{__obf_b1e5dea21e0aeea5, 0}, nil
	}
	__obf_bb2bbaf5ad72fade := // TODO: optimize extraction of fractional part
		Decimal{__obf_b1e5dea21e0aeea5, 0}
	__obf_0aed4c55bfd02d07, __obf_e556d4c0c2678d61 := __obf_83d5b629b20f9f91.QuoRem(__obf_bb2bbaf5ad72fade, 0)

	if __obf_03da3ca9b693fc27 == -1 && !__obf_e556d4c0c2678d61.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}
	__obf_af049c7749e14f94, _ := __obf_a1ffba48568e5107.__obf_bdd60e7777859473(__obf_0aed4c55bfd02d07.__obf_4c2240e6f10eb22b,

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_d55e3219d646b95d)

	if __obf_e556d4c0c2678d61.__obf_4c2240e6f10eb22b.Sign() == 0 {
		return __obf_af049c7749e14f94, nil
	}
	__obf_c3c6dad7153ba371 := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_a1ffba48568e5107.NumDigits()
	__obf_c354c851af938e33 := __obf_83d5b629b20f9f91.NumDigits()

	if int32(__obf_c3c6dad7153ba371) > __obf_d55e3219d646b95d {
		__obf_d55e3219d646b95d = int32(__obf_c3c6dad7153ba371)
	}
	if int32(__obf_c354c851af938e33) > __obf_d55e3219d646b95d {
		__obf_d55e3219d646b95d += int32(__obf_c354c851af938e33)
	}
	__obf_d55e3219d646b95d += // increase precision by 10 to compensate for errors in further calculations
		10
	__obf_ebd4a5da1aa116e2,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_1cf2c004a31cf720 := __obf_a1ffba48568e5107.Abs().Ln(__obf_d55e3219d646b95d)
	if __obf_1cf2c004a31cf720 != nil {
		return Decimal{}, __obf_1cf2c004a31cf720
	}
	__obf_ebd4a5da1aa116e2 = __obf_ebd4a5da1aa116e2.Mul(__obf_e556d4c0c2678d61)
	__obf_ebd4a5da1aa116e2, __obf_1cf2c004a31cf720 = __obf_ebd4a5da1aa116e2.ExpTaylor(__obf_d55e3219d646b95d)
	if __obf_1cf2c004a31cf720 != nil {
		return Decimal{}, __obf_1cf2c004a31cf720
	}
	__obf_d45eb54903cb2ff9 := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_af049c7749e14f94.Mul(__obf_ebd4a5da1aa116e2)

	return __obf_d45eb54903cb2ff9, nil
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
func (__obf_a1ffba48568e5107 Decimal) PowInt32(__obf_3ab33991ba7f9863 int32) (Decimal, error) {
	if __obf_a1ffba48568e5107.IsZero() && __obf_3ab33991ba7f9863 == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_d17e080daf2903f4 := __obf_3ab33991ba7f9863 < 0
	__obf_3ab33991ba7f9863 = __obf_d59b95d5897e7e7c(__obf_3ab33991ba7f9863)
	__obf_6a02be54118e0a20, __obf_5d6120ed52d2dc0a := __obf_a1ffba48568e5107, New(1, 0)

	for __obf_3ab33991ba7f9863 > 0 {
		if __obf_3ab33991ba7f9863%2 == 1 {
			__obf_5d6120ed52d2dc0a = __obf_5d6120ed52d2dc0a.Mul(__obf_6a02be54118e0a20)
		}
		__obf_3ab33991ba7f9863 /= 2

		if __obf_3ab33991ba7f9863 > 0 {
			__obf_6a02be54118e0a20 = __obf_6a02be54118e0a20.Mul(__obf_6a02be54118e0a20)
		}
	}

	if __obf_d17e080daf2903f4 {
		return New(1, 0).DivRound(__obf_5d6120ed52d2dc0a, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_5d6120ed52d2dc0a, nil
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
func (__obf_a1ffba48568e5107 Decimal) PowBigInt(__obf_3ab33991ba7f9863 *big.Int) (Decimal, error) {
	return __obf_a1ffba48568e5107.__obf_bdd60e7777859473(__obf_3ab33991ba7f9863, int32(PowPrecisionNegativeExponent))
}

func (__obf_a1ffba48568e5107 Decimal) __obf_bdd60e7777859473(__obf_3ab33991ba7f9863 *big.Int, __obf_d55e3219d646b95d int32) (Decimal, error) {
	if __obf_a1ffba48568e5107.IsZero() && __obf_3ab33991ba7f9863.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_afe1fccc344ef8b2 := new(big.Int).Set(__obf_3ab33991ba7f9863)
	__obf_d17e080daf2903f4 := __obf_3ab33991ba7f9863.Sign() < 0

	if __obf_d17e080daf2903f4 {
		__obf_afe1fccc344ef8b2.
			Abs(__obf_afe1fccc344ef8b2)
	}
	__obf_6a02be54118e0a20, __obf_5d6120ed52d2dc0a := __obf_a1ffba48568e5107, New(1, 0)

	for __obf_afe1fccc344ef8b2.Sign() > 0 {
		if __obf_afe1fccc344ef8b2.Bit(0) == 1 {
			__obf_5d6120ed52d2dc0a = __obf_5d6120ed52d2dc0a.Mul(__obf_6a02be54118e0a20)
		}
		__obf_afe1fccc344ef8b2.
			Rsh(__obf_afe1fccc344ef8b2, 1)

		if __obf_afe1fccc344ef8b2.Sign() > 0 {
			__obf_6a02be54118e0a20 = __obf_6a02be54118e0a20.Mul(__obf_6a02be54118e0a20)
		}
	}

	if __obf_d17e080daf2903f4 {
		return New(1, 0).DivRound(__obf_5d6120ed52d2dc0a, __obf_d55e3219d646b95d), nil
	}

	return __obf_5d6120ed52d2dc0a, nil
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
func (__obf_a1ffba48568e5107 Decimal) ExpHullAbrham(__obf_5db8394c87b56211 uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_a1ffba48568e5107.IsZero() {
		return Decimal{__obf_b1e5dea21e0aeea5, 0}, nil
	}
	__obf_0627340b851255bd := __obf_5db8394c87b56211

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_dc107516767dfe6c := __obf_a1ffba48568e5107.Abs().InexactFloat64()
	if __obf_d80648debc1b5218 := __obf_dc107516767dfe6c / 23; __obf_d80648debc1b5218 > float64(__obf_0627340b851255bd) && __obf_d80648debc1b5218 < float64(ExpMaxIterations) {
		__obf_0627340b851255bd = uint32(math.Ceil(__obf_d80648debc1b5218))
	}
	__obf_624d8e6829ade8f4 := // fail if abs(d) beyond an over/underflow threshold
		New(23*int64(__obf_0627340b851255bd), 0)
	if __obf_a1ffba48568e5107.Abs().Cmp(__obf_624d8e6829ade8f4) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}
	__obf_2437f45bc1996816 := // Return 1 if abs(d) small enough; this also avoids later over/underflow
		New(9, -int32(__obf_0627340b851255bd)-1)
	if __obf_a1ffba48568e5107.Abs().Cmp(__obf_2437f45bc1996816) <= 0 {
		return Decimal{__obf_b1e5dea21e0aeea5, __obf_a1ffba48568e5107.

			// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
			__obf_3ab33991ba7f9863}, nil
	}
	__obf_8007358663fe4d23 := __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 + int32(__obf_a1ffba48568e5107.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_8007358663fe4d23 < 0 {
		__obf_8007358663fe4d23 = 0
	}
	__obf_0020a03a38ba89e2 := New(1, __obf_8007358663fe4d23)
	__obf_7314b6cb9c5fa307 := // reduction factor
		Decimal{new(big.Int).Set(__obf_a1ffba48568e5107. // reduced argument
									__obf_4c2240e6f10eb22b), __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 - __obf_8007358663fe4d23}
	__obf_dd1d09a9386d1843 := int32(__obf_0627340b851255bd) + __obf_8007358663fe4d23 + 2
	__obf_31e9913868d1b855 := // precision for calculating the sum

		// Determine n, the number of therms for calculating sum
		// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
		// for solving appropriate equation, along with directed
		// roundings and simple rational bound for log10(p/abs(r))
		__obf_7314b6cb9c5fa307.Abs().InexactFloat64()
	__obf_ca25a840e80c0be2 := float64(__obf_dd1d09a9386d1843)
	__obf_4f539d9034b72303 := math.Ceil((1.453*__obf_ca25a840e80c0be2 - 1.182) / math.Log10(__obf_ca25a840e80c0be2/__obf_31e9913868d1b855))
	if __obf_4f539d9034b72303 > float64(ExpMaxIterations) || math.IsNaN(__obf_4f539d9034b72303) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_6a02be54118e0a20 := int64(__obf_4f539d9034b72303)
	__obf_b79982d611e468ff := New(0, 0)
	__obf_a022926e6dbb0d03 := New(1, 0)
	__obf_bb2bbaf5ad72fade := New(1, 0)
	for __obf_fb920e988f68d242 := __obf_6a02be54118e0a20 - 1; __obf_fb920e988f68d242 > 0; __obf_fb920e988f68d242-- {
		__obf_b79982d611e468ff.__obf_4c2240e6f10eb22b.
			SetInt64(__obf_fb920e988f68d242)
		__obf_a022926e6dbb0d03 = __obf_a022926e6dbb0d03.Mul(__obf_7314b6cb9c5fa307.DivRound(__obf_b79982d611e468ff, __obf_dd1d09a9386d1843))
		__obf_a022926e6dbb0d03 = __obf_a022926e6dbb0d03.Add(__obf_bb2bbaf5ad72fade)
	}
	__obf_203a8da5b2a3c81c := __obf_0020a03a38ba89e2.IntPart()
	__obf_d45eb54903cb2ff9 := New(1, 0)
	for __obf_fb920e988f68d242 := __obf_203a8da5b2a3c81c; __obf_fb920e988f68d242 > 0; __obf_fb920e988f68d242-- {
		__obf_d45eb54903cb2ff9 = __obf_d45eb54903cb2ff9.Mul(__obf_a022926e6dbb0d03)
	}
	__obf_c74d9203d9d43ae1 := int32(__obf_d45eb54903cb2ff9.NumDigits())

	var __obf_4800287bf7397896 int32
	if __obf_c74d9203d9d43ae1 > __obf_d59b95d5897e7e7c(__obf_d45eb54903cb2ff9.__obf_3ab33991ba7f9863) {
		__obf_4800287bf7397896 = int32(__obf_0627340b851255bd) - __obf_c74d9203d9d43ae1 - __obf_d45eb54903cb2ff9.__obf_3ab33991ba7f9863
	} else {
		__obf_4800287bf7397896 = int32(__obf_0627340b851255bd)
	}
	__obf_d45eb54903cb2ff9 = __obf_d45eb54903cb2ff9.Round(__obf_4800287bf7397896)

	return __obf_d45eb54903cb2ff9, nil
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
func (__obf_a1ffba48568e5107 Decimal) ExpTaylor(__obf_d55e3219d646b95d int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_a1ffba48568e5107.IsZero() {
		return Decimal{__obf_b1e5dea21e0aeea5, 0}.Round(__obf_d55e3219d646b95d), nil
	}

	var __obf_3be5dbb8fe389901 Decimal
	var __obf_9e800d5a4fcd8f43 int32
	if __obf_d55e3219d646b95d < 0 {
		__obf_3be5dbb8fe389901 = New(1, -1)
		__obf_9e800d5a4fcd8f43 = 8
	} else {
		__obf_3be5dbb8fe389901 = New(1, -__obf_d55e3219d646b95d-1)
		__obf_9e800d5a4fcd8f43 = __obf_d55e3219d646b95d + 1
	}
	__obf_1ec864ccd89cc1e2 := __obf_a1ffba48568e5107.Abs()
	__obf_af302cd96fe8c63c := __obf_a1ffba48568e5107.Abs()
	__obf_9548748c96d82681 := New(1, 0)
	__obf_5d6120ed52d2dc0a := New(1, 0)

	for __obf_fb920e988f68d242 := int64(1); ; {
		__obf_815b92cd689b8c05 := __obf_af302cd96fe8c63c.DivRound(__obf_9548748c96d82681, __obf_9e800d5a4fcd8f43)
		__obf_5d6120ed52d2dc0a = __obf_5d6120ed52d2dc0a.Add(__obf_815b92cd689b8c05)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_815b92cd689b8c05.Cmp(__obf_3be5dbb8fe389901) < 0 {
			break
		}
		__obf_af302cd96fe8c63c = __obf_af302cd96fe8c63c.Mul(__obf_1ec864ccd89cc1e2)
		__obf_fb920e988f68d242++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_26acda08e4a64645) >= int(__obf_fb920e988f68d242) && !__obf_26acda08e4a64645[__obf_fb920e988f68d242-1].IsZero() {
			__obf_9548748c96d82681 = __obf_26acda08e4a64645[__obf_fb920e988f68d242-1]
		} else {
			__obf_9548748c96d82681 = // To avoid any race conditions, firstly the zero value is appended to a slice to create
				// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
				// factorial using the index notation.
				__obf_26acda08e4a64645[__obf_fb920e988f68d242-2].Mul(New(__obf_fb920e988f68d242, 0))
			__obf_26acda08e4a64645 = append(__obf_26acda08e4a64645, Zero)
			__obf_26acda08e4a64645[__obf_fb920e988f68d242-1] = __obf_9548748c96d82681
		}
	}

	if __obf_a1ffba48568e5107.Sign() < 0 {
		__obf_5d6120ed52d2dc0a = New(1, 0).DivRound(__obf_5d6120ed52d2dc0a, __obf_d55e3219d646b95d+1)
	}
	__obf_5d6120ed52d2dc0a = __obf_5d6120ed52d2dc0a.Round(__obf_d55e3219d646b95d)
	return __obf_5d6120ed52d2dc0a, nil
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
func (__obf_a1ffba48568e5107 Decimal) Ln(__obf_d55e3219d646b95d int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_a1ffba48568e5107.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_a1ffba48568e5107.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}
	__obf_3b5beea964890a7c := __obf_d55e3219d646b95d + 2
	__obf_afa2d9a7dac2d065 := __obf_a1ffba48568e5107.Copy()

	var __obf_6eb27d25e6de812b, __obf_1fa0b1c6b41ba420, __obf_f78c0b969f39751a, __obf_e5eafcfc40394a1b, __obf_369a24f38e1ac9ce Decimal
	__obf_6eb27d25e6de812b = __obf_afa2d9a7dac2d065.Sub(Decimal{__obf_b1e5dea21e0aeea5, 0})
	__obf_1fa0b1c6b41ba420 = Decimal{__obf_b1e5dea21e0aeea5, -1}
	__obf_50dad0887807a28e := // for decimal in range [0.9, 1.1] where ln(d) is close to 0
		false

	if __obf_6eb27d25e6de812b.Abs().Cmp(__obf_1fa0b1c6b41ba420) <= 0 {
		__obf_50dad0887807a28e = true
	} else {
		__obf_af892703ca3392e5 := // reduce input decimal to range [0.1, 1)
			int32(__obf_afa2d9a7dac2d065.NumDigits()) + __obf_afa2d9a7dac2d065.__obf_3ab33991ba7f9863

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_afa2d9a7dac2d065.__obf_3ab33991ba7f9863 -= __obf_af892703ca3392e5
		__obf_a0f887999bf7be49 := __obf_a0f887999bf7be49.__obf_fc062c35aeb13736(__obf_3b5beea964890a7c)
		__obf_369a24f38e1ac9ce = NewFromInt32(__obf_af892703ca3392e5)
		__obf_369a24f38e1ac9ce = __obf_369a24f38e1ac9ce.Mul(__obf_a0f887999bf7be49)
		__obf_6eb27d25e6de812b = __obf_afa2d9a7dac2d065.Sub(Decimal{__obf_b1e5dea21e0aeea5, 0})

		if __obf_6eb27d25e6de812b.Abs().Cmp(__obf_1fa0b1c6b41ba420) <= 0 {
			__obf_50dad0887807a28e = true
		} else {
			__obf_f63ef0af26769878 := // initial estimate using floats
				__obf_afa2d9a7dac2d065.InexactFloat64()
			__obf_6eb27d25e6de812b = NewFromFloat(math.Log(__obf_f63ef0af26769878))
		}
	}
	__obf_3be5dbb8fe389901 := Decimal{__obf_b1e5dea21e0aeea5, -__obf_3b5beea964890a7c}

	if __obf_50dad0887807a28e {
		__obf_f78c0b969f39751a = // Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
			// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			// until the difference between current and next term is smaller than epsilon.
			// Coverage quite fast for decimals close to 1.0

			// z + 2
			__obf_6eb27d25e6de812b.Add(Decimal{__obf_53e3975b80a15b69, 0})
		__obf_1fa0b1c6b41ba420 = // z / (z + 2)
			__obf_6eb27d25e6de812b.DivRound(__obf_f78c0b969f39751a, __obf_3b5beea964890a7c)
		__obf_6eb27d25e6de812b = // 2 * (z / (z + 2))
			__obf_1fa0b1c6b41ba420.Add(__obf_1fa0b1c6b41ba420)
		__obf_f78c0b969f39751a = __obf_6eb27d25e6de812b.Copy()

		for __obf_6a02be54118e0a20 := 1; ; __obf_6a02be54118e0a20++ {
			__obf_f78c0b969f39751a = // 2 * (z / (z+2))^(2n+1)
				__obf_f78c0b969f39751a.Mul(__obf_1fa0b1c6b41ba420).Mul(__obf_1fa0b1c6b41ba420)
			__obf_e5eafcfc40394a1b = // 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
				NewFromInt(int64(2*__obf_6a02be54118e0a20 + 1))
			__obf_e5eafcfc40394a1b = __obf_f78c0b969f39751a.DivRound(__obf_e5eafcfc40394a1b, __obf_3b5beea964890a7c)
			__obf_6eb27d25e6de812b = // comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
				__obf_6eb27d25e6de812b.Add(__obf_e5eafcfc40394a1b)

			if __obf_e5eafcfc40394a1b.Abs().Cmp(__obf_3be5dbb8fe389901) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_5b14435a590cfe49 Decimal
		__obf_4518ebf28d866865 := __obf_3b5beea964890a7c*2 + 10

		for __obf_fb920e988f68d242 := int32(0); __obf_fb920e988f68d242 < __obf_4518ebf28d866865;
		// exp(a_n)
		__obf_fb920e988f68d242++ {
			__obf_1fa0b1c6b41ba420, _ = __obf_6eb27d25e6de812b.ExpTaylor(__obf_3b5beea964890a7c)
			__obf_f78c0b969f39751a = // exp(a_n) - z
				__obf_1fa0b1c6b41ba420.Sub(__obf_afa2d9a7dac2d065)
			__obf_f78c0b969f39751a = // 2 * (exp(a_n) - z)
				__obf_f78c0b969f39751a.Add(__obf_f78c0b969f39751a)
			__obf_e5eafcfc40394a1b = // exp(a_n) + z
				__obf_1fa0b1c6b41ba420.Add(__obf_afa2d9a7dac2d065)
			__obf_1fa0b1c6b41ba420 = // 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_f78c0b969f39751a.DivRound(__obf_e5eafcfc40394a1b, __obf_3b5beea964890a7c)
			__obf_6eb27d25e6de812b = // comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_6eb27d25e6de812b.Sub(__obf_1fa0b1c6b41ba420)

			if __obf_5b14435a590cfe49.Add(__obf_1fa0b1c6b41ba420).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_1fa0b1c6b41ba420.Abs().Cmp(__obf_3be5dbb8fe389901) <= 0 {
				break
			}
			__obf_5b14435a590cfe49 = __obf_1fa0b1c6b41ba420
		}
	}
	__obf_6eb27d25e6de812b = __obf_6eb27d25e6de812b.Add(__obf_369a24f38e1ac9ce)

	return __obf_6eb27d25e6de812b.Round(__obf_d55e3219d646b95d), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_a1ffba48568e5107 Decimal) NumDigits() int {
	if __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b == nil {
		return 1
	}

	if __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.IsInt64() {
		__obf_a617a84a4bf24c14 := __obf_a1ffba48568e5107.
			// restrict fast path to integers with exact conversion to float64
			__obf_4c2240e6f10eb22b.Int64()

		if __obf_a617a84a4bf24c14 <= (1<<53) && __obf_a617a84a4bf24c14 >= -(1<<53) {
			if __obf_a617a84a4bf24c14 == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_a617a84a4bf24c14)))) + 1
		}
	}
	__obf_7b4acc2c78427dda := int(float64(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.BitLen()) / math.Log2(10))
	__obf_f848f6520207de4b := // estimatedNumDigits (lg10) may be off by 1, need to verify
		big.NewInt(int64(__obf_7b4acc2c78427dda))
	__obf_ed6674b73f8aafd6 := __obf_f848f6520207de4b.Exp(__obf_263daee51f205bb2, __obf_f848f6520207de4b, nil)

	if __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.CmpAbs(__obf_ed6674b73f8aafd6) >= 0 {
		return __obf_7b4acc2c78427dda + 1
	}

	return __obf_7b4acc2c78427dda
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_a1ffba48568e5107 Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_7314b6cb9c5fa307 big.Int
	__obf_d77d99c0876d6027 := new(big.Int).Set(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b)
	for __obf_afa2d9a7dac2d065 := __obf_d59b95d5897e7e7c(__obf_a1ffba48568e5107.__obf_3ab33991ba7f9863); __obf_afa2d9a7dac2d065 > 0; __obf_afa2d9a7dac2d065-- {
		__obf_d77d99c0876d6027.
			QuoRem(__obf_d77d99c0876d6027, __obf_263daee51f205bb2, &__obf_7314b6cb9c5fa307)
		if __obf_7314b6cb9c5fa307.Cmp(__obf_1573d6eb6e796b01) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_d59b95d5897e7e7c(__obf_6a02be54118e0a20 int32) int32 {
	if __obf_6a02be54118e0a20 < 0 {
		return -__obf_6a02be54118e0a20
	}
	return __obf_6a02be54118e0a20
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_a1ffba48568e5107 Decimal) Cmp(__obf_83d5b629b20f9f91 Decimal) int {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()
	__obf_83d5b629b20f9f91.__obf_aa61179208fecfc6()

	if __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 == __obf_83d5b629b20f9f91.__obf_3ab33991ba7f9863 {
		return __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.Cmp(__obf_83d5b629b20f9f91.__obf_4c2240e6f10eb22b)
	}
	__obf_9fdd40229da129a1, __obf_f88a1c7011d37ae3 := RescalePair(__obf_a1ffba48568e5107, __obf_83d5b629b20f9f91)

	return __obf_9fdd40229da129a1.__obf_4c2240e6f10eb22b.Cmp(__obf_f88a1c7011d37ae3.

		// Compare compares the numbers represented by d and d2 and returns:
		//
		//	-1 if d <  d2
		//	 0 if d == d2
		//	+1 if d >  d2
		__obf_4c2240e6f10eb22b)
}

func (__obf_a1ffba48568e5107 Decimal) Compare(__obf_83d5b629b20f9f91 Decimal) int {
	return __obf_a1ffba48568e5107.Cmp(__obf_83d5b629b20f9f91)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_a1ffba48568e5107 Decimal) Equal(__obf_83d5b629b20f9f91 Decimal) bool {
	return __obf_a1ffba48568e5107.Cmp(__obf_83d5b629b20f9f91) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_a1ffba48568e5107 Decimal) Equals(__obf_83d5b629b20f9f91 Decimal) bool {
	return __obf_a1ffba48568e5107.Equal(__obf_83d5b629b20f9f91)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_a1ffba48568e5107 Decimal) GreaterThan(__obf_83d5b629b20f9f91 Decimal) bool {
	return __obf_a1ffba48568e5107.Cmp(__obf_83d5b629b20f9f91) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_a1ffba48568e5107 Decimal) GreaterThanOrEqual(__obf_83d5b629b20f9f91 Decimal) bool {
	__obf_44c66411987163f4 := __obf_a1ffba48568e5107.Cmp(__obf_83d5b629b20f9f91)
	return __obf_44c66411987163f4 == 1 || __obf_44c66411987163f4 == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_a1ffba48568e5107 Decimal) LessThan(__obf_83d5b629b20f9f91 Decimal) bool {
	return __obf_a1ffba48568e5107.Cmp(__obf_83d5b629b20f9f91) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_a1ffba48568e5107 Decimal) LessThanOrEqual(__obf_83d5b629b20f9f91 Decimal) bool {
	__obf_44c66411987163f4 := __obf_a1ffba48568e5107.Cmp(__obf_83d5b629b20f9f91)
	return __obf_44c66411987163f4 == -1 || __obf_44c66411987163f4 == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_a1ffba48568e5107 Decimal) Sign() int {
	if __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b == nil {
		return 0
	}
	return __obf_a1ffba48568e5107.

		// IsPositive return
		//
		//	true if d > 0
		//	false if d == 0
		//	false if d < 0
		__obf_4c2240e6f10eb22b.Sign()
}

func (__obf_a1ffba48568e5107 Decimal) IsPositive() bool {
	return __obf_a1ffba48568e5107.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_a1ffba48568e5107 Decimal) IsNegative() bool {
	return __obf_a1ffba48568e5107.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_a1ffba48568e5107 Decimal) IsZero() bool {
	return __obf_a1ffba48568e5107.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_a1ffba48568e5107 Decimal) Exponent() int32 {
	return __obf_a1ffba48568e5107.

		// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
		__obf_3ab33991ba7f9863
}

func (__obf_a1ffba48568e5107 Decimal) Coefficient() *big.Int {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_a1ffba48568e5107.

		// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
		// If coefficient cannot be represented in an int64, the result will be undefined.
		__obf_4c2240e6f10eb22b)
}

func (__obf_a1ffba48568e5107 Decimal) CoefficientInt64() int64 {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()
	return __obf_a1ffba48568e5107.

		// IntPart returns the integer component of the decimal.
		__obf_4c2240e6f10eb22b.Int64()
}

func (__obf_a1ffba48568e5107 Decimal) IntPart() int64 {
	__obf_d7cb3a9295d790f5 := __obf_a1ffba48568e5107.__obf_09c56f1962a9d090(0)
	return __obf_d7cb3a9295d790f5.__obf_4c2240e6f10eb22b.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_a1ffba48568e5107 Decimal) BigInt() *big.Int {
	__obf_d7cb3a9295d790f5 := __obf_a1ffba48568e5107.__obf_09c56f1962a9d090(0)
	return __obf_d7cb3a9295d790f5.

		// BigFloat returns decimal as BigFloat.
		// Be aware that casting decimal to BigFloat might cause a loss of precision.
		__obf_4c2240e6f10eb22b
}

func (__obf_a1ffba48568e5107 Decimal) BigFloat() *big.Float {
	__obf_dc107516767dfe6c := &big.Float{}
	__obf_dc107516767dfe6c.
		SetString(__obf_a1ffba48568e5107.String())
	return __obf_dc107516767dfe6c
}

// Rat returns a rational number representation of the decimal.
func (__obf_a1ffba48568e5107 Decimal) Rat() *big.Rat {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()
	if __obf_a1ffba48568e5107.
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_3ab33991ba7f9863 <= 0 {
		__obf_e33f274ec305c0ae := new(big.Int).Exp(__obf_263daee51f205bb2, big.NewInt(-int64(__obf_a1ffba48568e5107.__obf_3ab33991ba7f9863)), nil)
		return new(big.Rat).SetFrac(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b, __obf_e33f274ec305c0ae)
	}
	__obf_4135f12361329f7c := new(big.Int).Exp(__obf_263daee51f205bb2, big.NewInt(int64(__obf_a1ffba48568e5107.__obf_3ab33991ba7f9863)), nil)
	__obf_d82bfd7c0fea7547 := new(big.Int).Mul(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b, __obf_4135f12361329f7c)
	return new(big.Rat).SetFrac(__obf_d82bfd7c0fea7547,

		// Float64 returns the nearest float64 value for d and a bool indicating
		// whether f represents d exactly.
		// For more details, see the documentation for big.Rat.Float64
		__obf_b1e5dea21e0aeea5)
}

func (__obf_a1ffba48568e5107 Decimal) Float64() (__obf_dc107516767dfe6c float64, __obf_6f633dd615a89609 bool) {
	return __obf_a1ffba48568e5107.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_a1ffba48568e5107 Decimal) InexactFloat64() float64 {
	__obf_dc107516767dfe6c, _ := __obf_a1ffba48568e5107.Float64()
	return __obf_dc107516767dfe6c
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
func (__obf_a1ffba48568e5107 Decimal) String() string {
	return __obf_a1ffba48568e5107.string(true)
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
func (__obf_a1ffba48568e5107 Decimal) StringFixed(__obf_10bc0161e9bca3c9 int32) string {
	__obf_e89bd7abfa64e41c := __obf_a1ffba48568e5107.Round(__obf_10bc0161e9bca3c9)
	return __obf_e89bd7abfa64e41c.string(false)
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
func (__obf_a1ffba48568e5107 Decimal) StringFixedBank(__obf_10bc0161e9bca3c9 int32) string {
	__obf_e89bd7abfa64e41c := __obf_a1ffba48568e5107.RoundBank(__obf_10bc0161e9bca3c9)
	return __obf_e89bd7abfa64e41c.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_a1ffba48568e5107 Decimal) StringFixedCash(__obf_cdeac507726081a4 uint8) string {
	__obf_e89bd7abfa64e41c := __obf_a1ffba48568e5107.RoundCash(__obf_cdeac507726081a4)
	return __obf_e89bd7abfa64e41c.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_a1ffba48568e5107 Decimal) Round(__obf_10bc0161e9bca3c9 int32) Decimal {
	if __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 == -__obf_10bc0161e9bca3c9 {
		return __obf_a1ffba48568e5107
	}
	__obf_a9867477212e8b38 := // truncate to places + 1
		__obf_a1ffba48568e5107.__obf_09c56f1962a9d090(-__obf_10bc0161e9bca3c9 - 1)

	// add sign(d) * 0.5
	if __obf_a9867477212e8b38.__obf_4c2240e6f10eb22b.Sign() < 0 {
		__obf_a9867477212e8b38.__obf_4c2240e6f10eb22b.
			Sub(__obf_a9867477212e8b38.__obf_4c2240e6f10eb22b, __obf_953cc735cf320980)
	} else {
		__obf_a9867477212e8b38.__obf_4c2240e6f10eb22b.
			Add(__obf_a9867477212e8b38.__obf_4c2240e6f10eb22b,

				// floor for positive numbers, ceil for negative numbers
				__obf_953cc735cf320980)
	}

	_, __obf_a49032a74f3c7436 := __obf_a9867477212e8b38.__obf_4c2240e6f10eb22b.DivMod(__obf_a9867477212e8b38.__obf_4c2240e6f10eb22b, __obf_263daee51f205bb2, new(big.Int))
	__obf_a9867477212e8b38.__obf_3ab33991ba7f9863++
	if __obf_a9867477212e8b38.__obf_4c2240e6f10eb22b.Sign() < 0 && __obf_a49032a74f3c7436.Cmp(__obf_1573d6eb6e796b01) != 0 {
		__obf_a9867477212e8b38.__obf_4c2240e6f10eb22b.
			Add(__obf_a9867477212e8b38.__obf_4c2240e6f10eb22b,

				// RoundCeil rounds the decimal towards +infinity.
				//
				// Example:
				//
				//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
				//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
				//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
				//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
				__obf_b1e5dea21e0aeea5)
	}

	return __obf_a9867477212e8b38
}

func (__obf_a1ffba48568e5107 Decimal) RoundCeil(__obf_10bc0161e9bca3c9 int32) Decimal {
	if __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 >= -__obf_10bc0161e9bca3c9 {
		return __obf_a1ffba48568e5107
	}
	__obf_ae11d0ed5baaa6ca := __obf_a1ffba48568e5107.__obf_09c56f1962a9d090(-__obf_10bc0161e9bca3c9)
	if __obf_a1ffba48568e5107.Equal(__obf_ae11d0ed5baaa6ca) {
		return __obf_a1ffba48568e5107
	}

	if __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.Sign() > 0 {
		__obf_ae11d0ed5baaa6ca.__obf_4c2240e6f10eb22b.
			Add(__obf_ae11d0ed5baaa6ca.__obf_4c2240e6f10eb22b, __obf_b1e5dea21e0aeea5)
	}

	return __obf_ae11d0ed5baaa6ca
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_a1ffba48568e5107 Decimal) RoundFloor(__obf_10bc0161e9bca3c9 int32) Decimal {
	if __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 >= -__obf_10bc0161e9bca3c9 {
		return __obf_a1ffba48568e5107
	}
	__obf_ae11d0ed5baaa6ca := __obf_a1ffba48568e5107.__obf_09c56f1962a9d090(-__obf_10bc0161e9bca3c9)
	if __obf_a1ffba48568e5107.Equal(__obf_ae11d0ed5baaa6ca) {
		return __obf_a1ffba48568e5107
	}

	if __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.Sign() < 0 {
		__obf_ae11d0ed5baaa6ca.__obf_4c2240e6f10eb22b.
			Sub(__obf_ae11d0ed5baaa6ca.__obf_4c2240e6f10eb22b, __obf_b1e5dea21e0aeea5)
	}

	return __obf_ae11d0ed5baaa6ca
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_a1ffba48568e5107 Decimal) RoundUp(__obf_10bc0161e9bca3c9 int32) Decimal {
	if __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 >= -__obf_10bc0161e9bca3c9 {
		return __obf_a1ffba48568e5107
	}
	__obf_ae11d0ed5baaa6ca := __obf_a1ffba48568e5107.__obf_09c56f1962a9d090(-__obf_10bc0161e9bca3c9)
	if __obf_a1ffba48568e5107.Equal(__obf_ae11d0ed5baaa6ca) {
		return __obf_a1ffba48568e5107
	}

	if __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.Sign() > 0 {
		__obf_ae11d0ed5baaa6ca.__obf_4c2240e6f10eb22b.
			Add(__obf_ae11d0ed5baaa6ca.__obf_4c2240e6f10eb22b, __obf_b1e5dea21e0aeea5)
	} else if __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.Sign() < 0 {
		__obf_ae11d0ed5baaa6ca.__obf_4c2240e6f10eb22b.
			Sub(__obf_ae11d0ed5baaa6ca.__obf_4c2240e6f10eb22b, __obf_b1e5dea21e0aeea5)
	}

	return __obf_ae11d0ed5baaa6ca
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_a1ffba48568e5107 Decimal) RoundDown(__obf_10bc0161e9bca3c9 int32) Decimal {
	if __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 >= -__obf_10bc0161e9bca3c9 {
		return __obf_a1ffba48568e5107
	}
	__obf_ae11d0ed5baaa6ca := __obf_a1ffba48568e5107.__obf_09c56f1962a9d090(-__obf_10bc0161e9bca3c9)
	if __obf_a1ffba48568e5107.Equal(__obf_ae11d0ed5baaa6ca) {
		return __obf_a1ffba48568e5107
	}
	return __obf_ae11d0ed5baaa6ca
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
func (__obf_a1ffba48568e5107 Decimal) RoundBank(__obf_10bc0161e9bca3c9 int32) Decimal {
	__obf_14e512d7658e82e6 := __obf_a1ffba48568e5107.Round(__obf_10bc0161e9bca3c9)
	__obf_20e5e7579f8fb5b9 := __obf_a1ffba48568e5107.Sub(__obf_14e512d7658e82e6).Abs()
	__obf_e32d453b64e3e020 := New(5, -__obf_10bc0161e9bca3c9-1)
	if __obf_20e5e7579f8fb5b9.Cmp(__obf_e32d453b64e3e020) == 0 && __obf_14e512d7658e82e6.__obf_4c2240e6f10eb22b.Bit(0) != 0 {
		if __obf_14e512d7658e82e6.__obf_4c2240e6f10eb22b.Sign() < 0 {
			__obf_14e512d7658e82e6.__obf_4c2240e6f10eb22b.
				Add(__obf_14e512d7658e82e6.__obf_4c2240e6f10eb22b, __obf_b1e5dea21e0aeea5)
		} else {
			__obf_14e512d7658e82e6.__obf_4c2240e6f10eb22b.
				Sub(__obf_14e512d7658e82e6.__obf_4c2240e6f10eb22b, __obf_b1e5dea21e0aeea5)
		}
	}

	return __obf_14e512d7658e82e6
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
func (__obf_a1ffba48568e5107 Decimal) RoundCash(__obf_cdeac507726081a4 uint8) Decimal {
	var __obf_6687751773d688eb *big.Int
	switch __obf_cdeac507726081a4 {
	case 5:
		__obf_6687751773d688eb = __obf_58870f8992c18ef8
	case 10:
		__obf_6687751773d688eb = __obf_263daee51f205bb2
	case 25:
		__obf_6687751773d688eb = __obf_1f236215c06e3969
	case 50:
		__obf_6687751773d688eb = __obf_53e3975b80a15b69
	case 100:
		__obf_6687751773d688eb = __obf_b1e5dea21e0aeea5
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_cdeac507726081a4))
	}
	__obf_665b4151c218822a := Decimal{__obf_4c2240e6f10eb22b: __obf_6687751773d688eb}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_a1ffba48568e5107.Mul(__obf_665b4151c218822a).Round(0).Div(__obf_665b4151c218822a).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_a1ffba48568e5107 Decimal) Floor() Decimal {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()

	if __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 >= 0 {
		return __obf_a1ffba48568e5107
	}
	__obf_3ab33991ba7f9863 := big.NewInt(10)
	__obf_3ab33991ba7f9863.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_3ab33991ba7f9863, big.NewInt(-int64(__obf_a1ffba48568e5107.__obf_3ab33991ba7f9863)), nil)
	__obf_afa2d9a7dac2d065 := new(big.Int).Div(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b, __obf_3ab33991ba7f9863)
	return Decimal{__obf_4c2240e6f10eb22b: __obf_afa2d9a7dac2d065,

		// Ceil returns the nearest integer value greater than or equal to d.
		__obf_3ab33991ba7f9863: 0}
}

func (__obf_a1ffba48568e5107 Decimal) Ceil() Decimal {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()

	if __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 >= 0 {
		return __obf_a1ffba48568e5107
	}
	__obf_3ab33991ba7f9863 := big.NewInt(10)
	__obf_3ab33991ba7f9863.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_3ab33991ba7f9863, big.NewInt(-int64(__obf_a1ffba48568e5107.__obf_3ab33991ba7f9863)), nil)
	__obf_afa2d9a7dac2d065, __obf_a49032a74f3c7436 := new(big.Int).DivMod(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b, __obf_3ab33991ba7f9863, new(big.Int))
	if __obf_a49032a74f3c7436.Cmp(__obf_1573d6eb6e796b01) != 0 {
		__obf_afa2d9a7dac2d065.
			Add(__obf_afa2d9a7dac2d065, __obf_b1e5dea21e0aeea5)
	}
	return Decimal{__obf_4c2240e6f10eb22b: __obf_afa2d9a7dac2d065,

		// Truncate truncates off digits from the number, without rounding.
		//
		// NOTE: precision is the last digit that will not be truncated (must be >= 0).
		//
		// Example:
		//
		//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
		__obf_3ab33991ba7f9863: 0}
}

func (__obf_a1ffba48568e5107 Decimal) Truncate(__obf_d55e3219d646b95d int32) Decimal {
	__obf_a1ffba48568e5107.__obf_aa61179208fecfc6()
	if __obf_d55e3219d646b95d >= 0 && -__obf_d55e3219d646b95d > __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 {
		return __obf_a1ffba48568e5107.__obf_09c56f1962a9d090(-__obf_d55e3219d646b95d)
	}
	return __obf_a1ffba48568e5107
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_a1ffba48568e5107 *Decimal) UnmarshalJSON(__obf_a23a47d96e818b99 []byte) error {
	if string(__obf_a23a47d96e818b99) == "null" {
		return nil
	}
	__obf_c837be1551a2aa19, __obf_1cf2c004a31cf720 := __obf_3c3148ab2002d2ec(__obf_a23a47d96e818b99)
	if __obf_1cf2c004a31cf720 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_a23a47d96e818b99, __obf_1cf2c004a31cf720)
	}
	__obf_bd46136174bdabe4, __obf_1cf2c004a31cf720 := NewFromString(__obf_c837be1551a2aa19)
	*__obf_a1ffba48568e5107 = __obf_bd46136174bdabe4
	if __obf_1cf2c004a31cf720 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_c837be1551a2aa19, __obf_1cf2c004a31cf720)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_a1ffba48568e5107 Decimal) MarshalJSON() ([]byte, error) {
	var __obf_c837be1551a2aa19 string
	if MarshalJSONWithoutQuotes {
		__obf_c837be1551a2aa19 = __obf_a1ffba48568e5107.String()
	} else {
		__obf_c837be1551a2aa19 = "\"" + __obf_a1ffba48568e5107.String() + "\""
	}
	return []byte(__obf_c837be1551a2aa19), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_a1ffba48568e5107 *Decimal) UnmarshalBinary(__obf_41fa9fc5a9fd0535 []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_41fa9fc5a9fd0535) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_41fa9fc5a9fd0535, len(__obf_41fa9fc5a9fd0535))
	}
	__obf_a1ffba48568e5107.

		// Extract the exponent
		__obf_3ab33991ba7f9863 = int32(binary.BigEndian.Uint32(__obf_41fa9fc5a9fd0535[:4]))
	__obf_a1ffba48568e5107.

		// Extract the value
		__obf_4c2240e6f10eb22b = new(big.Int)
	if __obf_1cf2c004a31cf720 := __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.GobDecode(__obf_41fa9fc5a9fd0535[4:]); __obf_1cf2c004a31cf720 != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_41fa9fc5a9fd0535, __obf_1cf2c004a31cf720)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_a1ffba48568e5107 Decimal) MarshalBinary() (__obf_41fa9fc5a9fd0535 []byte, __obf_1cf2c004a31cf720 error) {
	// exp is written first, but encode value first to know output size
	var __obf_efc1ff30ada828f4 []byte
	if __obf_efc1ff30ada828f4, __obf_1cf2c004a31cf720 = __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.GobEncode(); __obf_1cf2c004a31cf720 != nil {
		return nil, __obf_1cf2c004a31cf720
	}
	__obf_b3da55e24ba1b0c0 := // Write the exponent in front, since it's a fixed size
		make([]byte, 4, len(__obf_efc1ff30ada828f4)+4)
	binary.BigEndian.PutUint32(__obf_b3da55e24ba1b0c0, uint32(__obf_a1ffba48568e5107.

		// Return the byte array
		__obf_3ab33991ba7f9863))

	return append(__obf_b3da55e24ba1b0c0, __obf_efc1ff30ada828f4...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_a1ffba48568e5107 *Decimal) Scan(__obf_4c2240e6f10eb22b any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_ad7c1a28e6ea5999 := __obf_4c2240e6f10eb22b.(type) {

	case float32:
		*__obf_a1ffba48568e5107 = NewFromFloat(float64(__obf_ad7c1a28e6ea5999))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_a1ffba48568e5107 = NewFromFloat(__obf_ad7c1a28e6ea5999)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_a1ffba48568e5107 = New(__obf_ad7c1a28e6ea5999, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_a1ffba48568e5107 = NewFromUint64(__obf_ad7c1a28e6ea5999)
		return nil

	default:
		__obf_c837be1551a2aa19,
			// default is trying to interpret value stored as string
			__obf_1cf2c004a31cf720 := __obf_3c3148ab2002d2ec(__obf_ad7c1a28e6ea5999)
		if __obf_1cf2c004a31cf720 != nil {
			return __obf_1cf2c004a31cf720
		}
		*__obf_a1ffba48568e5107, __obf_1cf2c004a31cf720 = NewFromString(__obf_c837be1551a2aa19)
		return __obf_1cf2c004a31cf720
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_a1ffba48568e5107 Decimal) Value() (driver.Value, error) {
	return __obf_a1ffba48568e5107.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_a1ffba48568e5107 *Decimal) UnmarshalText(__obf_db8b106b5394aca4 []byte) error {
	__obf_c837be1551a2aa19 := string(__obf_db8b106b5394aca4)
	__obf_cad8fb067a1dbc4b, __obf_1cf2c004a31cf720 := NewFromString(__obf_c837be1551a2aa19)
	*__obf_a1ffba48568e5107 = __obf_cad8fb067a1dbc4b
	if __obf_1cf2c004a31cf720 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_c837be1551a2aa19, __obf_1cf2c004a31cf720)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_a1ffba48568e5107 Decimal) MarshalText() (__obf_db8b106b5394aca4 []byte, __obf_1cf2c004a31cf720 error) {
	return []byte(__obf_a1ffba48568e5107.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_a1ffba48568e5107 Decimal) GobEncode() ([]byte, error) {
	return __obf_a1ffba48568e5107.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_a1ffba48568e5107 *Decimal) GobDecode(__obf_41fa9fc5a9fd0535 []byte) error {
	return __obf_a1ffba48568e5107.UnmarshalBinary(__obf_41fa9fc5a9fd0535)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_a1ffba48568e5107 Decimal) StringScaled(__obf_3ab33991ba7f9863 int32) string {
	return __obf_a1ffba48568e5107.__obf_09c56f1962a9d090(__obf_3ab33991ba7f9863).String()
}

func (__obf_a1ffba48568e5107 Decimal) string(__obf_48a28a189d2a9874 bool) string {
	if __obf_a1ffba48568e5107.__obf_3ab33991ba7f9863 >= 0 {
		return __obf_a1ffba48568e5107.__obf_09c56f1962a9d090(0).__obf_4c2240e6f10eb22b.String()
	}
	__obf_d59b95d5897e7e7c := new(big.Int).Abs(__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b)
	__obf_c837be1551a2aa19 := __obf_d59b95d5897e7e7c.String()

	var __obf_0a32afaac98d0d57, __obf_6ac3d5ef277c4d17 string
	__obf_0e18f7b03ef5d5ae := // NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
		// and you are on a 32-bit machine. Won't fix this super-edge case.
		int(__obf_a1ffba48568e5107.__obf_3ab33991ba7f9863)
	if len(__obf_c837be1551a2aa19) > -__obf_0e18f7b03ef5d5ae {
		__obf_0a32afaac98d0d57 = __obf_c837be1551a2aa19[:len(__obf_c837be1551a2aa19)+__obf_0e18f7b03ef5d5ae]
		__obf_6ac3d5ef277c4d17 = __obf_c837be1551a2aa19[len(__obf_c837be1551a2aa19)+__obf_0e18f7b03ef5d5ae:]
	} else {
		__obf_0a32afaac98d0d57 = "0"
		__obf_a7fc924ab4beb25a := -__obf_0e18f7b03ef5d5ae - len(__obf_c837be1551a2aa19)
		__obf_6ac3d5ef277c4d17 = strings.Repeat("0", __obf_a7fc924ab4beb25a) + __obf_c837be1551a2aa19
	}

	if __obf_48a28a189d2a9874 {
		__obf_fb920e988f68d242 := len(__obf_6ac3d5ef277c4d17) - 1
		for ; __obf_fb920e988f68d242 >= 0; __obf_fb920e988f68d242-- {
			if __obf_6ac3d5ef277c4d17[__obf_fb920e988f68d242] != '0' {
				break
			}
		}
		__obf_6ac3d5ef277c4d17 = __obf_6ac3d5ef277c4d17[:__obf_fb920e988f68d242+1]
	}
	__obf_fccc32ceb795482d := __obf_0a32afaac98d0d57
	if len(__obf_6ac3d5ef277c4d17) > 0 {
		__obf_fccc32ceb795482d += "." + __obf_6ac3d5ef277c4d17
	}

	if __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b.Sign() < 0 {
		return "-" + __obf_fccc32ceb795482d
	}

	return __obf_fccc32ceb795482d
}

func (__obf_a1ffba48568e5107 *Decimal) __obf_aa61179208fecfc6() {
	if __obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b == nil {
		__obf_a1ffba48568e5107.__obf_4c2240e6f10eb22b = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_6cdb9ea6a328922c Decimal, __obf_106c342aac7ddb90 ...Decimal) Decimal {
	__obf_ea197583a522b9ea := __obf_6cdb9ea6a328922c
	for _, __obf_d60fc7544f73befd := range __obf_106c342aac7ddb90 {
		if __obf_d60fc7544f73befd.Cmp(__obf_ea197583a522b9ea) < 0 {
			__obf_ea197583a522b9ea = __obf_d60fc7544f73befd
		}
	}
	return __obf_ea197583a522b9ea
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_6cdb9ea6a328922c Decimal, __obf_106c342aac7ddb90 ...Decimal) Decimal {
	__obf_ea197583a522b9ea := __obf_6cdb9ea6a328922c
	for _, __obf_d60fc7544f73befd := range __obf_106c342aac7ddb90 {
		if __obf_d60fc7544f73befd.Cmp(__obf_ea197583a522b9ea) > 0 {
			__obf_ea197583a522b9ea = __obf_d60fc7544f73befd
		}
	}
	return __obf_ea197583a522b9ea
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_6cdb9ea6a328922c Decimal, __obf_106c342aac7ddb90 ...Decimal) Decimal {
	__obf_7891600e7f6603c6 := __obf_6cdb9ea6a328922c
	for _, __obf_d60fc7544f73befd := range __obf_106c342aac7ddb90 {
		__obf_7891600e7f6603c6 = __obf_7891600e7f6603c6.Add(__obf_d60fc7544f73befd)
	}

	return __obf_7891600e7f6603c6
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_6cdb9ea6a328922c Decimal, __obf_106c342aac7ddb90 ...Decimal) Decimal {
	__obf_b9b5005968664708 := New(int64(len(__obf_106c342aac7ddb90)+1), 0)
	__obf_a022926e6dbb0d03 := Sum(__obf_6cdb9ea6a328922c, __obf_106c342aac7ddb90...)
	return __obf_a022926e6dbb0d03.Div(__obf_b9b5005968664708)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_07860bf3a92a0614 Decimal, __obf_83d5b629b20f9f91 Decimal) (Decimal, Decimal) {
	__obf_07860bf3a92a0614.__obf_aa61179208fecfc6()
	__obf_83d5b629b20f9f91.__obf_aa61179208fecfc6()

	if __obf_07860bf3a92a0614.__obf_3ab33991ba7f9863 < __obf_83d5b629b20f9f91.__obf_3ab33991ba7f9863 {
		return __obf_07860bf3a92a0614, __obf_83d5b629b20f9f91.__obf_09c56f1962a9d090(__obf_07860bf3a92a0614.__obf_3ab33991ba7f9863)
	} else if __obf_07860bf3a92a0614.__obf_3ab33991ba7f9863 > __obf_83d5b629b20f9f91.__obf_3ab33991ba7f9863 {
		return __obf_07860bf3a92a0614.__obf_09c56f1962a9d090(__obf_83d5b629b20f9f91.__obf_3ab33991ba7f9863), __obf_83d5b629b20f9f91
	}

	return __obf_07860bf3a92a0614, __obf_83d5b629b20f9f91
}

func __obf_3c3148ab2002d2ec(__obf_4c2240e6f10eb22b any) (string, error) {
	var __obf_25c5d5bdaa486d93 []byte

	switch __obf_ad7c1a28e6ea5999 := __obf_4c2240e6f10eb22b.(type) {
	case string:
		__obf_25c5d5bdaa486d93 = []byte(__obf_ad7c1a28e6ea5999)
	case []byte:
		__obf_25c5d5bdaa486d93 = __obf_ad7c1a28e6ea5999
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_4c2240e6f10eb22b,

			// If the amount is quoted, strip the quotes
			__obf_4c2240e6f10eb22b)
	}

	if len(__obf_25c5d5bdaa486d93) > 2 && __obf_25c5d5bdaa486d93[0] == '"' && __obf_25c5d5bdaa486d93[len(__obf_25c5d5bdaa486d93)-1] == '"' {
		__obf_25c5d5bdaa486d93 = __obf_25c5d5bdaa486d93[1 : len(__obf_25c5d5bdaa486d93)-1]
	}
	return string(__obf_25c5d5bdaa486d93), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the datautil.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_a1ffba48568e5107 Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_a1ffba48568e5107,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_a1ffba48568e5107 *NullDecimal) Scan(__obf_4c2240e6f10eb22b any) error {
	if __obf_4c2240e6f10eb22b == nil {
		__obf_a1ffba48568e5107.
			Valid = false
		return nil
	}
	__obf_a1ffba48568e5107.
		Valid = true
	return __obf_a1ffba48568e5107.Decimal.Scan(__obf_4c2240e6f10eb22b)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_a1ffba48568e5107 NullDecimal) Value() (driver.Value, error) {
	if !__obf_a1ffba48568e5107.Valid {
		return nil, nil
	}
	return __obf_a1ffba48568e5107.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_a1ffba48568e5107 *NullDecimal) UnmarshalJSON(__obf_a23a47d96e818b99 []byte) error {
	if string(__obf_a23a47d96e818b99) == "null" {
		__obf_a1ffba48568e5107.
			Valid = false
		return nil
	}
	__obf_a1ffba48568e5107.
		Valid = true
	return __obf_a1ffba48568e5107.Decimal.UnmarshalJSON(__obf_a23a47d96e818b99)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_a1ffba48568e5107 NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_a1ffba48568e5107.Valid {
		return []byte("null"), nil
	}
	return __obf_a1ffba48568e5107.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_a1ffba48568e5107 *NullDecimal) UnmarshalText(__obf_db8b106b5394aca4 []byte) error {
	__obf_c837be1551a2aa19 := string(__obf_db8b106b5394aca4)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_c837be1551a2aa19 == "" {
		__obf_a1ffba48568e5107.
			Valid = false
		return nil
	}
	if __obf_1cf2c004a31cf720 := __obf_a1ffba48568e5107.Decimal.UnmarshalText(__obf_db8b106b5394aca4); __obf_1cf2c004a31cf720 != nil {
		__obf_a1ffba48568e5107.
			Valid = false
		return __obf_1cf2c004a31cf720
	}
	__obf_a1ffba48568e5107.
		Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_a1ffba48568e5107 NullDecimal) MarshalText() (__obf_db8b106b5394aca4 []byte, __obf_1cf2c004a31cf720 error) {
	if !__obf_a1ffba48568e5107.Valid {
		return []byte{}, nil
	}
	return __obf_a1ffba48568e5107.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_a1ffba48568e5107 Decimal) Atan() Decimal {
	if __obf_a1ffba48568e5107.Equal(NewFromFloat(0.0)) {
		return __obf_a1ffba48568e5107
	}
	if __obf_a1ffba48568e5107.GreaterThan(NewFromFloat(0.0)) {
		return __obf_a1ffba48568e5107.__obf_52fb030f6e7dab51()
	}
	return __obf_a1ffba48568e5107.Neg().__obf_52fb030f6e7dab51().Neg()
}

func (__obf_a1ffba48568e5107 Decimal) __obf_e141471b8a43b06e() Decimal {
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
	__obf_afa2d9a7dac2d065 := __obf_a1ffba48568e5107.Mul(__obf_a1ffba48568e5107)
	__obf_2476dec57e4c4ca0 := P0.Mul(__obf_afa2d9a7dac2d065).Add(P1).Mul(__obf_afa2d9a7dac2d065).Add(P2).Mul(__obf_afa2d9a7dac2d065).Add(P3).Mul(__obf_afa2d9a7dac2d065).Add(P4).Mul(__obf_afa2d9a7dac2d065)
	__obf_082609a2c5a001a9 := __obf_afa2d9a7dac2d065.Add(Q0).Mul(__obf_afa2d9a7dac2d065).Add(Q1).Mul(__obf_afa2d9a7dac2d065).Add(Q2).Mul(__obf_afa2d9a7dac2d065).Add(Q3).Mul(__obf_afa2d9a7dac2d065).Add(Q4)
	__obf_afa2d9a7dac2d065 = __obf_2476dec57e4c4ca0.Div(__obf_082609a2c5a001a9)
	__obf_afa2d9a7dac2d065 = __obf_a1ffba48568e5107.Mul(__obf_afa2d9a7dac2d065).Add(__obf_a1ffba48568e5107)
	return __obf_afa2d9a7dac2d065
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_a1ffba48568e5107 Decimal) __obf_52fb030f6e7dab51() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)
	__obf_6f512a579e076148 := // tan(3*pi/8)
		NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_a1ffba48568e5107.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_a1ffba48568e5107.__obf_e141471b8a43b06e()
	}
	if __obf_a1ffba48568e5107.GreaterThan(Tan3pio8) {
		return __obf_6f512a579e076148.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_a1ffba48568e5107).__obf_e141471b8a43b06e()).Add(Morebits)
	}
	return __obf_6f512a579e076148.Div(NewFromFloat(4.0)).Add((__obf_a1ffba48568e5107.Sub(NewFromFloat(1.0)).Div(__obf_a1ffba48568e5107.Add(NewFromFloat(1.0)))).__obf_e141471b8a43b06e()).Add(NewFromFloat(0.5).Mul(Morebits))
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
func (__obf_a1ffba48568e5107 Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_a1ffba48568e5107.Equal(NewFromFloat(0.0)) {
		return __obf_a1ffba48568e5107
	}
	__obf_6dbf9797d792aa4a := // make argument positive but save the sign
		false
	if __obf_a1ffba48568e5107.LessThan(NewFromFloat(0.0)) {
		__obf_a1ffba48568e5107 = __obf_a1ffba48568e5107.Neg()
		__obf_6dbf9797d792aa4a = true
	}
	__obf_ff71c1fe24b055e0 := __obf_a1ffba48568e5107.Mul(M4PI).IntPart()
	__obf_01b5b96e8f080261 := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_ff71c1fe24b055e0)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_ff71c1fe24b055e0&1 == 1 {
		__obf_ff71c1fe24b055e0++
		__obf_01b5b96e8f080261 = __obf_01b5b96e8f080261.Add(NewFromFloat(1.0))
	}
	__obf_ff71c1fe24b055e0 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_ff71c1fe24b055e0 > 3 {
		__obf_6dbf9797d792aa4a = !__obf_6dbf9797d792aa4a
		__obf_ff71c1fe24b055e0 -= 4
	}
	__obf_afa2d9a7dac2d065 := __obf_a1ffba48568e5107.Sub(__obf_01b5b96e8f080261.Mul(PI4A)).Sub(__obf_01b5b96e8f080261.Mul(PI4B)).Sub(__obf_01b5b96e8f080261.Mul(PI4C))
	__obf_4dc2778d7f4c5c3c := // Extended precision modular arithmetic
		__obf_afa2d9a7dac2d065.Mul(__obf_afa2d9a7dac2d065)

	if __obf_ff71c1fe24b055e0 == 1 || __obf_ff71c1fe24b055e0 == 2 {
		__obf_717676e856f9f552 := __obf_4dc2778d7f4c5c3c.Mul(__obf_4dc2778d7f4c5c3c).Mul(_cos[0].Mul(__obf_4dc2778d7f4c5c3c).Add(_cos[1]).Mul(__obf_4dc2778d7f4c5c3c).Add(_cos[2]).Mul(__obf_4dc2778d7f4c5c3c).Add(_cos[3]).Mul(__obf_4dc2778d7f4c5c3c).Add(_cos[4]).Mul(__obf_4dc2778d7f4c5c3c).Add(_cos[5]))
		__obf_01b5b96e8f080261 = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_4dc2778d7f4c5c3c)).Add(__obf_717676e856f9f552)
	} else {
		__obf_01b5b96e8f080261 = __obf_afa2d9a7dac2d065.Add(__obf_afa2d9a7dac2d065.Mul(__obf_4dc2778d7f4c5c3c).Mul(_sin[0].Mul(__obf_4dc2778d7f4c5c3c).Add(_sin[1]).Mul(__obf_4dc2778d7f4c5c3c).Add(_sin[2]).Mul(__obf_4dc2778d7f4c5c3c).Add(_sin[3]).Mul(__obf_4dc2778d7f4c5c3c).Add(_sin[4]).Mul(__obf_4dc2778d7f4c5c3c).Add(_sin[5])))
	}
	if __obf_6dbf9797d792aa4a {
		__obf_01b5b96e8f080261 = __obf_01b5b96e8f080261.Neg()
	}
	return __obf_01b5b96e8f080261
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
func (__obf_a1ffba48568e5107 Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)  // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)  // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15) // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125)
	__obf_6dbf9797d792aa4a := // 4/pi

		// make argument positive
		false
	if __obf_a1ffba48568e5107.LessThan(NewFromFloat(0.0)) {
		__obf_a1ffba48568e5107 = __obf_a1ffba48568e5107.Neg()
	}
	__obf_ff71c1fe24b055e0 := __obf_a1ffba48568e5107.Mul(M4PI).IntPart()
	__obf_01b5b96e8f080261 := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_ff71c1fe24b055e0)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_ff71c1fe24b055e0&1 == 1 {
		__obf_ff71c1fe24b055e0++
		__obf_01b5b96e8f080261 = __obf_01b5b96e8f080261.Add(NewFromFloat(1.0))
	}
	__obf_ff71c1fe24b055e0 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_ff71c1fe24b055e0 > 3 {
		__obf_6dbf9797d792aa4a = !__obf_6dbf9797d792aa4a
		__obf_ff71c1fe24b055e0 -= 4
	}
	if __obf_ff71c1fe24b055e0 > 1 {
		__obf_6dbf9797d792aa4a = !__obf_6dbf9797d792aa4a
	}
	__obf_afa2d9a7dac2d065 := __obf_a1ffba48568e5107.Sub(__obf_01b5b96e8f080261.Mul(PI4A)).Sub(__obf_01b5b96e8f080261.Mul(PI4B)).Sub(__obf_01b5b96e8f080261.Mul(PI4C))
	__obf_4dc2778d7f4c5c3c := // Extended precision modular arithmetic
		__obf_afa2d9a7dac2d065.Mul(__obf_afa2d9a7dac2d065)

	if __obf_ff71c1fe24b055e0 == 1 || __obf_ff71c1fe24b055e0 == 2 {
		__obf_01b5b96e8f080261 = __obf_afa2d9a7dac2d065.Add(__obf_afa2d9a7dac2d065.Mul(__obf_4dc2778d7f4c5c3c).Mul(_sin[0].Mul(__obf_4dc2778d7f4c5c3c).Add(_sin[1]).Mul(__obf_4dc2778d7f4c5c3c).Add(_sin[2]).Mul(__obf_4dc2778d7f4c5c3c).Add(_sin[3]).Mul(__obf_4dc2778d7f4c5c3c).Add(_sin[4]).Mul(__obf_4dc2778d7f4c5c3c).Add(_sin[5])))
	} else {
		__obf_717676e856f9f552 := __obf_4dc2778d7f4c5c3c.Mul(__obf_4dc2778d7f4c5c3c).Mul(_cos[0].Mul(__obf_4dc2778d7f4c5c3c).Add(_cos[1]).Mul(__obf_4dc2778d7f4c5c3c).Add(_cos[2]).Mul(__obf_4dc2778d7f4c5c3c).Add(_cos[3]).Mul(__obf_4dc2778d7f4c5c3c).Add(_cos[4]).Mul(__obf_4dc2778d7f4c5c3c).Add(_cos[5]))
		__obf_01b5b96e8f080261 = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_4dc2778d7f4c5c3c)).Add(__obf_717676e856f9f552)
	}
	if __obf_6dbf9797d792aa4a {
		__obf_01b5b96e8f080261 = __obf_01b5b96e8f080261.Neg()
	}
	return __obf_01b5b96e8f080261
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
func (__obf_a1ffba48568e5107 Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_a1ffba48568e5107.Equal(NewFromFloat(0.0)) {
		return __obf_a1ffba48568e5107
	}
	__obf_6dbf9797d792aa4a := // make argument positive but save the sign
		false
	if __obf_a1ffba48568e5107.LessThan(NewFromFloat(0.0)) {
		__obf_a1ffba48568e5107 = __obf_a1ffba48568e5107.Neg()
		__obf_6dbf9797d792aa4a = true
	}
	__obf_ff71c1fe24b055e0 := __obf_a1ffba48568e5107.Mul(M4PI).IntPart()
	__obf_01b5b96e8f080261 := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_ff71c1fe24b055e0)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_ff71c1fe24b055e0&1 == 1 {
		__obf_ff71c1fe24b055e0++
		__obf_01b5b96e8f080261 = __obf_01b5b96e8f080261.Add(NewFromFloat(1.0))
	}
	__obf_afa2d9a7dac2d065 := __obf_a1ffba48568e5107.Sub(__obf_01b5b96e8f080261.Mul(PI4A)).Sub(__obf_01b5b96e8f080261.Mul(PI4B)).Sub(__obf_01b5b96e8f080261.Mul(PI4C))
	__obf_4dc2778d7f4c5c3c := // Extended precision modular arithmetic
		__obf_afa2d9a7dac2d065.Mul(__obf_afa2d9a7dac2d065)

	if __obf_4dc2778d7f4c5c3c.GreaterThan(NewFromFloat(1e-14)) {
		__obf_717676e856f9f552 := __obf_4dc2778d7f4c5c3c.Mul(_tanP[0].Mul(__obf_4dc2778d7f4c5c3c).Add(_tanP[1]).Mul(__obf_4dc2778d7f4c5c3c).Add(_tanP[2]))
		__obf_d691de659afcc030 := __obf_4dc2778d7f4c5c3c.Add(_tanQ[1]).Mul(__obf_4dc2778d7f4c5c3c).Add(_tanQ[2]).Mul(__obf_4dc2778d7f4c5c3c).Add(_tanQ[3]).Mul(__obf_4dc2778d7f4c5c3c).Add(_tanQ[4])
		__obf_01b5b96e8f080261 = __obf_afa2d9a7dac2d065.Add(__obf_afa2d9a7dac2d065.Mul(__obf_717676e856f9f552.Div(__obf_d691de659afcc030)))
	} else {
		__obf_01b5b96e8f080261 = __obf_afa2d9a7dac2d065
	}
	if __obf_ff71c1fe24b055e0&2 == 2 {
		__obf_01b5b96e8f080261 = NewFromFloat(-1.0).Div(__obf_01b5b96e8f080261)
	}
	if __obf_6dbf9797d792aa4a {
		__obf_01b5b96e8f080261 = __obf_01b5b96e8f080261.Neg()
	}
	return __obf_01b5b96e8f080261
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

type __obf_bd46136174bdabe4 struct {
	__obf_a1ffba48568e5107 [800]byte
	__obf_3018d9650ee6ca39 int  // digits, big-endian representation
	__obf_c3fb49899f9ddeff int  // number of digits used
	__obf_c5a0f28fc9f4a6ce bool // decimal point
	__obf_701edbf11344f61d bool // negative flag
	// discarded nonzero digits beyond d[:nd]
}

func (__obf_97404a94f9526684 *__obf_bd46136174bdabe4) String() string {
	__obf_6a02be54118e0a20 := 10 + __obf_97404a94f9526684.__obf_3018d9650ee6ca39
	if __obf_97404a94f9526684.__obf_c3fb49899f9ddeff > 0 {
		__obf_6a02be54118e0a20 += __obf_97404a94f9526684.__obf_c3fb49899f9ddeff
	}
	if __obf_97404a94f9526684.__obf_c3fb49899f9ddeff < 0 {
		__obf_6a02be54118e0a20 += -__obf_97404a94f9526684.__obf_c3fb49899f9ddeff
	}
	__obf_5ef9c62c1175aedc := make([]byte, __obf_6a02be54118e0a20)
	__obf_717676e856f9f552 := 0
	switch {
	case __obf_97404a94f9526684.__obf_3018d9650ee6ca39 == 0:
		return "0"

	case __obf_97404a94f9526684.
		// zeros fill space between decimal point and digits
		__obf_c3fb49899f9ddeff <= 0:
		__obf_5ef9c62c1175aedc[__obf_717676e856f9f552] = '0'
		__obf_717676e856f9f552++
		__obf_5ef9c62c1175aedc[__obf_717676e856f9f552] = '.'
		__obf_717676e856f9f552++
		__obf_717676e856f9f552 += __obf_d293eda1502d30d0(__obf_5ef9c62c1175aedc[__obf_717676e856f9f552 : __obf_717676e856f9f552+-__obf_97404a94f9526684.__obf_c3fb49899f9ddeff])
		__obf_717676e856f9f552 += copy(__obf_5ef9c62c1175aedc[__obf_717676e856f9f552:], __obf_97404a94f9526684.__obf_a1ffba48568e5107[0:__obf_97404a94f9526684.__obf_3018d9650ee6ca39])

	case __obf_97404a94f9526684.
		// decimal point in middle of digits
		__obf_c3fb49899f9ddeff < __obf_97404a94f9526684.__obf_3018d9650ee6ca39:
		__obf_717676e856f9f552 += copy(__obf_5ef9c62c1175aedc[__obf_717676e856f9f552:], __obf_97404a94f9526684.__obf_a1ffba48568e5107[0:__obf_97404a94f9526684.__obf_c3fb49899f9ddeff])
		__obf_5ef9c62c1175aedc[__obf_717676e856f9f552] = '.'
		__obf_717676e856f9f552++
		__obf_717676e856f9f552 += copy(__obf_5ef9c62c1175aedc[__obf_717676e856f9f552:], __obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_97404a94f9526684.__obf_c3fb49899f9ddeff:

		// zeros fill space between digits and decimal point
		__obf_97404a94f9526684.__obf_3018d9650ee6ca39])

	default:
		__obf_717676e856f9f552 += copy(__obf_5ef9c62c1175aedc[__obf_717676e856f9f552:], __obf_97404a94f9526684.__obf_a1ffba48568e5107[0:__obf_97404a94f9526684.__obf_3018d9650ee6ca39])
		__obf_717676e856f9f552 += __obf_d293eda1502d30d0(__obf_5ef9c62c1175aedc[__obf_717676e856f9f552 : __obf_717676e856f9f552+__obf_97404a94f9526684.__obf_c3fb49899f9ddeff-__obf_97404a94f9526684.__obf_3018d9650ee6ca39])
	}
	return string(__obf_5ef9c62c1175aedc[0:__obf_717676e856f9f552])
}

func __obf_d293eda1502d30d0(__obf_19eef2f96c90d2fb []byte) int {
	for __obf_fb920e988f68d242 := range __obf_19eef2f96c90d2fb {
		__obf_19eef2f96c90d2fb[__obf_fb920e988f68d242] = '0'
	}
	return len(__obf_19eef2f96c90d2fb)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_bc806f761c1943c6(__obf_97404a94f9526684 *__obf_bd46136174bdabe4) {
	for __obf_97404a94f9526684.__obf_3018d9650ee6ca39 > 0 && __obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_97404a94f9526684.__obf_3018d9650ee6ca39-1] == '0' {
		__obf_97404a94f9526684.__obf_3018d9650ee6ca39--
	}
	if __obf_97404a94f9526684.__obf_3018d9650ee6ca39 == 0 {
		__obf_97404a94f9526684.

			// Assign v to a.
			__obf_c3fb49899f9ddeff = 0
	}
}

func (__obf_97404a94f9526684 *__obf_bd46136174bdabe4) Assign(__obf_ad7c1a28e6ea5999 uint64) {
	var __obf_5ef9c62c1175aedc [24]byte
	__obf_6a02be54118e0a20 := // Write reversed decimal in buf.
		0
	for __obf_ad7c1a28e6ea5999 > 0 {
		__obf_426559321a715d9b := __obf_ad7c1a28e6ea5999 / 10
		__obf_ad7c1a28e6ea5999 -= 10 * __obf_426559321a715d9b
		__obf_5ef9c62c1175aedc[__obf_6a02be54118e0a20] = byte(__obf_ad7c1a28e6ea5999 + '0')
		__obf_6a02be54118e0a20++
		__obf_ad7c1a28e6ea5999 = __obf_426559321a715d9b
	}
	__obf_97404a94f9526684.

		// Reverse again to produce forward decimal in a.d.
		__obf_3018d9650ee6ca39 = 0
	for __obf_6a02be54118e0a20--; __obf_6a02be54118e0a20 >= 0; __obf_6a02be54118e0a20-- {
		__obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_97404a94f9526684.__obf_3018d9650ee6ca39] = __obf_5ef9c62c1175aedc[__obf_6a02be54118e0a20]
		__obf_97404a94f9526684.__obf_3018d9650ee6ca39++
	}
	__obf_97404a94f9526684.__obf_c3fb49899f9ddeff = __obf_97404a94f9526684.

		// Maximum shift that we can do in one pass without overflow.
		// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
		__obf_3018d9650ee6ca39
	__obf_bc806f761c1943c6(__obf_97404a94f9526684)
}

const __obf_f84c9c75f884a86f = 32 << (^uint(0) >> 63)
const __obf_c96a6ce315f5f6ea = __obf_f84c9c75f884a86f - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_d4f5ee819fc6c1e4(__obf_97404a94f9526684 *__obf_bd46136174bdabe4, __obf_0020a03a38ba89e2 uint) {
	__obf_7314b6cb9c5fa307 := 0
	__obf_717676e856f9f552 := // read pointer
		0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_6a02be54118e0a20 uint
	for ; __obf_6a02be54118e0a20>>__obf_0020a03a38ba89e2 == 0; __obf_7314b6cb9c5fa307++ {
		if __obf_7314b6cb9c5fa307 >= __obf_97404a94f9526684.__obf_3018d9650ee6ca39 {
			if __obf_6a02be54118e0a20 == 0 {
				__obf_97404a94f9526684.
					// a == 0; shouldn't get here, but handle anyway.
					__obf_3018d9650ee6ca39 = 0
				return
			}
			for __obf_6a02be54118e0a20>>__obf_0020a03a38ba89e2 == 0 {
				__obf_6a02be54118e0a20 = __obf_6a02be54118e0a20 * 10
				__obf_7314b6cb9c5fa307++
			}
			break
		}
		__obf_4cc6d31ede5e315a := uint(__obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_7314b6cb9c5fa307])
		__obf_6a02be54118e0a20 = __obf_6a02be54118e0a20*10 + __obf_4cc6d31ede5e315a - '0'
	}
	__obf_97404a94f9526684.__obf_c3fb49899f9ddeff -= __obf_7314b6cb9c5fa307 - 1

	var __obf_b8cf225d71ca4c1c uint = (1 << __obf_0020a03a38ba89e2) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_7314b6cb9c5fa307 < __obf_97404a94f9526684.__obf_3018d9650ee6ca39; __obf_7314b6cb9c5fa307++ {
		__obf_4cc6d31ede5e315a := uint(__obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_7314b6cb9c5fa307])
		__obf_69ee1cdaa89a2c6f := __obf_6a02be54118e0a20 >> __obf_0020a03a38ba89e2
		__obf_6a02be54118e0a20 &= __obf_b8cf225d71ca4c1c
		__obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_717676e856f9f552] = byte(__obf_69ee1cdaa89a2c6f + '0')
		__obf_717676e856f9f552++
		__obf_6a02be54118e0a20 = __obf_6a02be54118e0a20*10 + __obf_4cc6d31ede5e315a - '0'
	}

	// Put down extra digits.
	for __obf_6a02be54118e0a20 > 0 {
		__obf_69ee1cdaa89a2c6f := __obf_6a02be54118e0a20 >> __obf_0020a03a38ba89e2
		__obf_6a02be54118e0a20 &= __obf_b8cf225d71ca4c1c
		if __obf_717676e856f9f552 < len(__obf_97404a94f9526684.__obf_a1ffba48568e5107) {
			__obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_717676e856f9f552] = byte(__obf_69ee1cdaa89a2c6f + '0')
			__obf_717676e856f9f552++
		} else if __obf_69ee1cdaa89a2c6f > 0 {
			__obf_97404a94f9526684.__obf_701edbf11344f61d = true
		}
		__obf_6a02be54118e0a20 = __obf_6a02be54118e0a20 * 10
	}
	__obf_97404a94f9526684.__obf_3018d9650ee6ca39 = __obf_717676e856f9f552

	// Cheat sheet for left shift: table indexed by shift count giving
	// number of new digits that will be introduced by that shift.
	//
	// For example, leftcheats[4] = {2, "625"}.  That means that
	// if we are shifting by 4 (multiplying by 16), it will add 2 digits
	// when the string prefix is "625" through "999", and one fewer digit
	// if the string prefix is "000" through "624".
	//
	// Credit for this trick goes to Ken.
	__obf_bc806f761c1943c6(__obf_97404a94f9526684)
}

type __obf_3fa310871bcf1562 struct {
	__obf_e19cb14b57179ccb int
	__obf_706783e1d6409b75 string // number of new digits
	// minus one digit if original < a.
}

var __obf_c4bf526f3f67130d = []__obf_3fa310871bcf1562{
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
func __obf_45ffa59d871fdb61(__obf_84ea471b6b7e49a0 []byte, __obf_9de430966e5fb3a1 string) bool {
	for __obf_fb920e988f68d242 := 0; __obf_fb920e988f68d242 < len(__obf_9de430966e5fb3a1); __obf_fb920e988f68d242++ {
		if __obf_fb920e988f68d242 >= len(__obf_84ea471b6b7e49a0) {
			return true
		}
		if __obf_84ea471b6b7e49a0[__obf_fb920e988f68d242] != __obf_9de430966e5fb3a1[__obf_fb920e988f68d242] {
			return __obf_84ea471b6b7e49a0[__obf_fb920e988f68d242] < __obf_9de430966e5fb3a1[__obf_fb920e988f68d242]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_6265b208997ecde1(__obf_97404a94f9526684 *__obf_bd46136174bdabe4, __obf_0020a03a38ba89e2 uint) {
	__obf_e19cb14b57179ccb := __obf_c4bf526f3f67130d[__obf_0020a03a38ba89e2].__obf_e19cb14b57179ccb
	if __obf_45ffa59d871fdb61(__obf_97404a94f9526684.__obf_a1ffba48568e5107[0:__obf_97404a94f9526684.__obf_3018d9650ee6ca39], __obf_c4bf526f3f67130d[__obf_0020a03a38ba89e2].__obf_706783e1d6409b75) {
		__obf_e19cb14b57179ccb--
	}
	__obf_7314b6cb9c5fa307 := __obf_97404a94f9526684. // read index
								__obf_3018d9650ee6ca39
	__obf_717676e856f9f552 := __obf_97404a94f9526684. // write index
								__obf_3018d9650ee6ca39 + __obf_e19cb14b57179ccb

	// Pick up a digit, put down a digit.
	var __obf_6a02be54118e0a20 uint
	for __obf_7314b6cb9c5fa307--; __obf_7314b6cb9c5fa307 >= 0; __obf_7314b6cb9c5fa307-- {
		__obf_6a02be54118e0a20 += (uint(__obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_7314b6cb9c5fa307]) - '0') << __obf_0020a03a38ba89e2
		__obf_fb3225eed3e8f0c2 := __obf_6a02be54118e0a20 / 10
		__obf_86c33a114aa48f86 := __obf_6a02be54118e0a20 - 10*__obf_fb3225eed3e8f0c2
		__obf_717676e856f9f552--
		if __obf_717676e856f9f552 < len(__obf_97404a94f9526684.__obf_a1ffba48568e5107) {
			__obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_717676e856f9f552] = byte(__obf_86c33a114aa48f86 + '0')
		} else if __obf_86c33a114aa48f86 != 0 {
			__obf_97404a94f9526684.__obf_701edbf11344f61d = true
		}
		__obf_6a02be54118e0a20 = __obf_fb3225eed3e8f0c2
	}

	// Put down extra digits.
	for __obf_6a02be54118e0a20 > 0 {
		__obf_fb3225eed3e8f0c2 := __obf_6a02be54118e0a20 / 10
		__obf_86c33a114aa48f86 := __obf_6a02be54118e0a20 - 10*__obf_fb3225eed3e8f0c2
		__obf_717676e856f9f552--
		if __obf_717676e856f9f552 < len(__obf_97404a94f9526684.__obf_a1ffba48568e5107) {
			__obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_717676e856f9f552] = byte(__obf_86c33a114aa48f86 + '0')
		} else if __obf_86c33a114aa48f86 != 0 {
			__obf_97404a94f9526684.__obf_701edbf11344f61d = true
		}
		__obf_6a02be54118e0a20 = __obf_fb3225eed3e8f0c2
	}
	__obf_97404a94f9526684.__obf_3018d9650ee6ca39 += __obf_e19cb14b57179ccb
	if __obf_97404a94f9526684.__obf_3018d9650ee6ca39 >= len(__obf_97404a94f9526684.__obf_a1ffba48568e5107) {
		__obf_97404a94f9526684.__obf_3018d9650ee6ca39 = len(__obf_97404a94f9526684.__obf_a1ffba48568e5107)
	}
	__obf_97404a94f9526684.__obf_c3fb49899f9ddeff += __obf_e19cb14b57179ccb

	// Binary shift left (k > 0) or right (k < 0).
	__obf_bc806f761c1943c6(__obf_97404a94f9526684)
}

func (__obf_97404a94f9526684 *__obf_bd46136174bdabe4) Shift(__obf_0020a03a38ba89e2 int) {
	switch {
	case __obf_97404a94f9526684.
		// nothing to do: a == 0
		__obf_3018d9650ee6ca39 == 0:

	case __obf_0020a03a38ba89e2 > 0:
		for __obf_0020a03a38ba89e2 > __obf_c96a6ce315f5f6ea {
			__obf_6265b208997ecde1(__obf_97404a94f9526684, __obf_c96a6ce315f5f6ea)
			__obf_0020a03a38ba89e2 -= __obf_c96a6ce315f5f6ea
		}
		__obf_6265b208997ecde1(__obf_97404a94f9526684, uint(__obf_0020a03a38ba89e2))
	case __obf_0020a03a38ba89e2 < 0:
		for __obf_0020a03a38ba89e2 < -__obf_c96a6ce315f5f6ea {
			__obf_d4f5ee819fc6c1e4(__obf_97404a94f9526684, __obf_c96a6ce315f5f6ea)
			__obf_0020a03a38ba89e2 += __obf_c96a6ce315f5f6ea
		}
		__obf_d4f5ee819fc6c1e4(__obf_97404a94f9526684, uint(-__obf_0020a03a38ba89e2))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_b3547f99a88665be(__obf_97404a94f9526684 *__obf_bd46136174bdabe4, __obf_3018d9650ee6ca39 int) bool {
	if __obf_3018d9650ee6ca39 < 0 || __obf_3018d9650ee6ca39 >= __obf_97404a94f9526684.__obf_3018d9650ee6ca39 {
		return false
	}
	if __obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_3018d9650ee6ca39] == '5' && __obf_3018d9650ee6ca39+1 == __obf_97404a94f9526684. // exactly halfway - round to even
																		__obf_3018d9650ee6ca39 {
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_97404a94f9526684.__obf_701edbf11344f61d {
			return true
		}
		return __obf_3018d9650ee6ca39 > 0 && (__obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_3018d9650ee6ca39-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_97404a94f9526684.

		// Round a to nd digits (or fewer).
		// If nd is zero, it means we're rounding
		// just to the left of the digits, as in
		// 0.09 -> 0.1.
		__obf_a1ffba48568e5107[__obf_3018d9650ee6ca39] >= '5'
}

func (__obf_97404a94f9526684 *__obf_bd46136174bdabe4) Round(__obf_3018d9650ee6ca39 int) {
	if __obf_3018d9650ee6ca39 < 0 || __obf_3018d9650ee6ca39 >= __obf_97404a94f9526684.__obf_3018d9650ee6ca39 {
		return
	}
	if __obf_b3547f99a88665be(__obf_97404a94f9526684, __obf_3018d9650ee6ca39) {
		__obf_97404a94f9526684.
			RoundUp(__obf_3018d9650ee6ca39)
	} else {
		__obf_97404a94f9526684.
			RoundDown(__obf_3018d9650ee6ca39)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_97404a94f9526684 *__obf_bd46136174bdabe4) RoundDown(__obf_3018d9650ee6ca39 int) {
	if __obf_3018d9650ee6ca39 < 0 || __obf_3018d9650ee6ca39 >= __obf_97404a94f9526684.__obf_3018d9650ee6ca39 {
		return
	}
	__obf_97404a94f9526684.__obf_3018d9650ee6ca39 = __obf_3018d9650ee6ca39

	// Round a up to nd digits (or fewer).
	__obf_bc806f761c1943c6(__obf_97404a94f9526684)
}

func (__obf_97404a94f9526684 *__obf_bd46136174bdabe4) RoundUp(__obf_3018d9650ee6ca39 int) {
	if __obf_3018d9650ee6ca39 < 0 || __obf_3018d9650ee6ca39 >= __obf_97404a94f9526684.

		// round up
		__obf_3018d9650ee6ca39 {
		return
	}

	for __obf_fb920e988f68d242 := __obf_3018d9650ee6ca39 - 1; __obf_fb920e988f68d242 >= 0; __obf_fb920e988f68d242-- {
		__obf_4cc6d31ede5e315a := __obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_fb920e988f68d242]
		if __obf_4cc6d31ede5e315a < '9' {
			__obf_97404a94f9526684. // can stop after this digit
						__obf_a1ffba48568e5107[__obf_fb920e988f68d242]++
			__obf_97404a94f9526684.__obf_3018d9650ee6ca39 = __obf_fb920e988f68d242 + 1
			return
		}
	}
	__obf_97404a94f9526684.

		// Number is all 9s.
		// Change to single 1 with adjusted decimal point.
		__obf_a1ffba48568e5107[0] = '1'
	__obf_97404a94f9526684.

		// Extract integer part, rounded appropriately.
		// No guarantees about overflow.
		__obf_3018d9650ee6ca39 = 1
	__obf_97404a94f9526684.__obf_c3fb49899f9ddeff++
}

func (__obf_97404a94f9526684 *__obf_bd46136174bdabe4) RoundedInteger() uint64 {
	if __obf_97404a94f9526684.__obf_c3fb49899f9ddeff > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_fb920e988f68d242 int
	__obf_6a02be54118e0a20 := uint64(0)
	for __obf_fb920e988f68d242 = 0; __obf_fb920e988f68d242 < __obf_97404a94f9526684.__obf_c3fb49899f9ddeff && __obf_fb920e988f68d242 < __obf_97404a94f9526684.__obf_3018d9650ee6ca39; __obf_fb920e988f68d242++ {
		__obf_6a02be54118e0a20 = __obf_6a02be54118e0a20*10 + uint64(__obf_97404a94f9526684.__obf_a1ffba48568e5107[__obf_fb920e988f68d242]-'0')
	}
	for ; __obf_fb920e988f68d242 < __obf_97404a94f9526684.__obf_c3fb49899f9ddeff; __obf_fb920e988f68d242++ {
		__obf_6a02be54118e0a20 *= 10
	}
	if __obf_b3547f99a88665be(__obf_97404a94f9526684, __obf_97404a94f9526684.__obf_c3fb49899f9ddeff) {
		__obf_6a02be54118e0a20++
	}
	return __obf_6a02be54118e0a20
}
