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
package __obf_17d0cfecf7e687b6

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

var __obf_f9353e10ec98e7cd = big.NewInt(0)
var __obf_38f9fa21a076c52c = big.NewInt(1)
var __obf_e9d2424d8d05f51b = big.NewInt(2)
var __obf_7e76cab4ab253fc7 = big.NewInt(4)
var __obf_ac970d134d03d297 = big.NewInt(5)
var __obf_4a348d545ccd57bf = big.NewInt(10)
var __obf_1fe59b8bee7770a1 = big.NewInt(20)

var __obf_a71570cb39e71d12 = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_04aabd427bafd586 *big.Int

	// NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.
	__obf_3e3f0912cbecfbf3 int32
}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_04aabd427bafd586 int64, __obf_3e3f0912cbecfbf3 int32) Decimal {
	return Decimal{
		__obf_04aabd427bafd586: big.NewInt(__obf_04aabd427bafd586),
		__obf_3e3f0912cbecfbf3: __obf_3e3f0912cbecfbf3,
	}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_04aabd427bafd586 int64) Decimal {
	return Decimal{
		__obf_04aabd427bafd586: big.NewInt(__obf_04aabd427bafd586),
		__obf_3e3f0912cbecfbf3: 0,
	}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_04aabd427bafd586 int32) Decimal {
	return Decimal{
		__obf_04aabd427bafd586: big.NewInt(int64(__obf_04aabd427bafd586)),
		__obf_3e3f0912cbecfbf3: 0,
	}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_04aabd427bafd586 uint64) Decimal {
	return Decimal{
		__obf_04aabd427bafd586: new(big.Int).SetUint64(__obf_04aabd427bafd586),
		__obf_3e3f0912cbecfbf3: 0,
	}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_04aabd427bafd586 *big.Int, __obf_3e3f0912cbecfbf3 int32) Decimal {
	return Decimal{
		__obf_04aabd427bafd586: new(big.Int).Set(__obf_04aabd427bafd586),
		__obf_3e3f0912cbecfbf3: __obf_3e3f0912cbecfbf3,
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
func NewFromBigRat(__obf_04aabd427bafd586 *big.Rat, __obf_c51dca51652f0a68 int32) Decimal {
	return Decimal{
		__obf_04aabd427bafd586: new(big.Int).Set(__obf_04aabd427bafd586.Num()),
		__obf_3e3f0912cbecfbf3: 0,
	}.DivRound(Decimal{
		__obf_04aabd427bafd586: new(big.Int).Set(__obf_04aabd427bafd586.Denom()),
		__obf_3e3f0912cbecfbf3: 0,
	}, __obf_c51dca51652f0a68)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_04aabd427bafd586 string) (Decimal, error) {
	__obf_47efb98752fd6eb7 := __obf_04aabd427bafd586
	var __obf_758b84e1207f8aae string
	var __obf_3e3f0912cbecfbf3 int64

	// Check if number is using scientific notation
	__obf_2c07019bf7738640 := strings.IndexAny(__obf_04aabd427bafd586, "Ee")
	if __obf_2c07019bf7738640 != -1 {
		__obf_53b1e48bd6fb616e, __obf_8d1e8d31f090b7f5 := strconv.ParseInt(__obf_04aabd427bafd586[__obf_2c07019bf7738640+1:], 10, 32)
		if __obf_8d1e8d31f090b7f5 != nil {
			if __obf_2832c1f1d480fa8e, __obf_613e1c247a70a3df := __obf_8d1e8d31f090b7f5.(*strconv.NumError); __obf_613e1c247a70a3df && __obf_2832c1f1d480fa8e.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_04aabd427bafd586)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_04aabd427bafd586)
		}
		__obf_04aabd427bafd586 = __obf_04aabd427bafd586[:__obf_2c07019bf7738640]
		__obf_3e3f0912cbecfbf3 = __obf_53b1e48bd6fb616e
	}

	__obf_f13795abf4f1f903 := -1
	__obf_7a99ff0fe7ee04b5 := len(__obf_04aabd427bafd586)
	for __obf_bf8f4a859c474f4d := 0; __obf_bf8f4a859c474f4d < __obf_7a99ff0fe7ee04b5; __obf_bf8f4a859c474f4d++ {
		if __obf_04aabd427bafd586[__obf_bf8f4a859c474f4d] == '.' {
			if __obf_f13795abf4f1f903 > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_04aabd427bafd586)
			}
			__obf_f13795abf4f1f903 = __obf_bf8f4a859c474f4d
		}
	}

	if __obf_f13795abf4f1f903 == -1 {
		// There is no decimal point, we can just parse the original string as
		// an int
		__obf_758b84e1207f8aae = __obf_04aabd427bafd586
	} else {
		if __obf_f13795abf4f1f903+1 < __obf_7a99ff0fe7ee04b5 {
			__obf_758b84e1207f8aae = __obf_04aabd427bafd586[:__obf_f13795abf4f1f903] + __obf_04aabd427bafd586[__obf_f13795abf4f1f903+1:]
		} else {
			__obf_758b84e1207f8aae = __obf_04aabd427bafd586[:__obf_f13795abf4f1f903]
		}
		__obf_53b1e48bd6fb616e := -len(__obf_04aabd427bafd586[__obf_f13795abf4f1f903+1:])
		__obf_3e3f0912cbecfbf3 += int64(__obf_53b1e48bd6fb616e)
	}

	var __obf_c0f2c2835752f02a *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_758b84e1207f8aae) <= 18 {
		__obf_7043ce2aef7320f7, __obf_8d1e8d31f090b7f5 := strconv.ParseInt(__obf_758b84e1207f8aae, 10, 64)
		if __obf_8d1e8d31f090b7f5 != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_04aabd427bafd586)
		}
		__obf_c0f2c2835752f02a = big.NewInt(__obf_7043ce2aef7320f7)
	} else {
		__obf_c0f2c2835752f02a = new(big.Int)
		_, __obf_613e1c247a70a3df := __obf_c0f2c2835752f02a.SetString(__obf_758b84e1207f8aae, 10)
		if !__obf_613e1c247a70a3df {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_04aabd427bafd586)
		}
	}

	if __obf_3e3f0912cbecfbf3 < math.MinInt32 || __obf_3e3f0912cbecfbf3 > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_47efb98752fd6eb7)
	}

	return Decimal{
		__obf_04aabd427bafd586: __obf_c0f2c2835752f02a,
		__obf_3e3f0912cbecfbf3: int32(__obf_3e3f0912cbecfbf3),
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
func NewFromFormattedString(__obf_04aabd427bafd586 string, __obf_51dba675b5fbab17 *regexp.Regexp) (Decimal, error) {
	__obf_2e4714248ec616ef := __obf_51dba675b5fbab17.ReplaceAllString(__obf_04aabd427bafd586, "")
	__obf_8a994c90f60ad404, __obf_8d1e8d31f090b7f5 := NewFromString(__obf_2e4714248ec616ef)
	if __obf_8d1e8d31f090b7f5 != nil {
		return Decimal{}, __obf_8d1e8d31f090b7f5
	}
	return __obf_8a994c90f60ad404, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_04aabd427bafd586 string) Decimal {
	__obf_89a3a56c47f41acb, __obf_8d1e8d31f090b7f5 := NewFromString(__obf_04aabd427bafd586)
	if __obf_8d1e8d31f090b7f5 != nil {
		panic(__obf_8d1e8d31f090b7f5)
	}
	return __obf_89a3a56c47f41acb
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
func NewFromFloat(__obf_04aabd427bafd586 float64) Decimal {
	if __obf_04aabd427bafd586 == 0 {
		return New(0, 0)
	}
	return __obf_51b9ba2e9375ad2c(__obf_04aabd427bafd586, math.Float64bits(__obf_04aabd427bafd586), &__obf_9dee2c5a44d4ec98)
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
func NewFromFloat32(__obf_04aabd427bafd586 float32) Decimal {
	if __obf_04aabd427bafd586 == 0 {
		return New(0, 0)
	}
	// XOR is workaround for https://github.com/golang/go/issues/26285
	__obf_b0487b5cb853a320 := math.Float32bits(__obf_04aabd427bafd586) ^ 0x80808080
	return __obf_51b9ba2e9375ad2c(float64(__obf_04aabd427bafd586), uint64(__obf_b0487b5cb853a320)^0x80808080, &__obf_b00f38b47daaa1ee)
}

func __obf_51b9ba2e9375ad2c(__obf_4d7c699332f4c605 float64, __obf_eddf34d110373ed9 uint64, __obf_25725b9a4386f10b *__obf_ce035e3885d5f561) Decimal {
	if math.IsNaN(__obf_4d7c699332f4c605) || math.IsInf(__obf_4d7c699332f4c605, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_4d7c699332f4c605))
	}
	__obf_3e3f0912cbecfbf3 := int(__obf_eddf34d110373ed9>>__obf_25725b9a4386f10b.__obf_714ef5fd93ab5467) & (1<<__obf_25725b9a4386f10b.__obf_54a54ea1a927766e - 1)
	__obf_26f3bdcd35e04593 := __obf_eddf34d110373ed9 & (uint64(1)<<__obf_25725b9a4386f10b.__obf_714ef5fd93ab5467 - 1)

	switch __obf_3e3f0912cbecfbf3 {
	case 0:
		// denormalized
		__obf_3e3f0912cbecfbf3++

	default:
		// add implicit top bit
		__obf_26f3bdcd35e04593 |= uint64(1) << __obf_25725b9a4386f10b.__obf_714ef5fd93ab5467
	}
	__obf_3e3f0912cbecfbf3 += __obf_25725b9a4386f10b.__obf_2d0469f7185b0ad4

	var __obf_8a994c90f60ad404 __obf_17d0cfecf7e687b6
	__obf_8a994c90f60ad404.Assign(__obf_26f3bdcd35e04593)
	__obf_8a994c90f60ad404.Shift(__obf_3e3f0912cbecfbf3 - int(__obf_25725b9a4386f10b.__obf_714ef5fd93ab5467))
	__obf_8a994c90f60ad404.__obf_f7e560ef2994ea5a = __obf_eddf34d110373ed9>>(__obf_25725b9a4386f10b.__obf_54a54ea1a927766e+__obf_25725b9a4386f10b.__obf_714ef5fd93ab5467) != 0

	__obf_8163cf85fce10314(&__obf_8a994c90f60ad404, __obf_26f3bdcd35e04593, __obf_3e3f0912cbecfbf3, __obf_25725b9a4386f10b)
	// If less than 19 digits, we can do calculation in an int64.
	if __obf_8a994c90f60ad404.__obf_4673196fea260de2 < 19 {
		__obf_4bbdbb28098ff3a4 := int64(0)
		__obf_1a0547a84e8f209a := int64(1)
		for __obf_bf8f4a859c474f4d := __obf_8a994c90f60ad404.__obf_4673196fea260de2 - 1; __obf_bf8f4a859c474f4d >= 0; __obf_bf8f4a859c474f4d-- {
			__obf_4bbdbb28098ff3a4 += __obf_1a0547a84e8f209a * int64(__obf_8a994c90f60ad404.__obf_8a994c90f60ad404[__obf_bf8f4a859c474f4d]-'0')
			__obf_1a0547a84e8f209a *= 10
		}
		if __obf_8a994c90f60ad404.__obf_f7e560ef2994ea5a {
			__obf_4bbdbb28098ff3a4 *= -1
		}
		return Decimal{__obf_04aabd427bafd586: big.NewInt(__obf_4bbdbb28098ff3a4), __obf_3e3f0912cbecfbf3: int32(__obf_8a994c90f60ad404.__obf_26344dec98644d1b) - int32(__obf_8a994c90f60ad404.__obf_4673196fea260de2)}
	}
	__obf_c0f2c2835752f02a := new(big.Int)
	__obf_c0f2c2835752f02a, __obf_613e1c247a70a3df := __obf_c0f2c2835752f02a.SetString(string(__obf_8a994c90f60ad404.__obf_8a994c90f60ad404[:__obf_8a994c90f60ad404.__obf_4673196fea260de2]), 10)
	if __obf_613e1c247a70a3df {
		return Decimal{__obf_04aabd427bafd586: __obf_c0f2c2835752f02a, __obf_3e3f0912cbecfbf3: int32(__obf_8a994c90f60ad404.__obf_26344dec98644d1b) - int32(__obf_8a994c90f60ad404.__obf_4673196fea260de2)}
	}

	return NewFromFloatWithExponent(__obf_4d7c699332f4c605, int32(__obf_8a994c90f60ad404.__obf_26344dec98644d1b)-int32(__obf_8a994c90f60ad404.__obf_4673196fea260de2))
}

// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
// number of fractional digits.
//
// Example:
//
//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
func NewFromFloatWithExponent(__obf_04aabd427bafd586 float64, __obf_3e3f0912cbecfbf3 int32) Decimal {
	if math.IsNaN(__obf_04aabd427bafd586) || math.IsInf(__obf_04aabd427bafd586, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_04aabd427bafd586))
	}

	__obf_eddf34d110373ed9 := math.Float64bits(__obf_04aabd427bafd586)
	__obf_26f3bdcd35e04593 := __obf_eddf34d110373ed9 & (1<<52 - 1)
	__obf_71f3d6c74e9e688b := int32((__obf_eddf34d110373ed9 >> 52) & (1<<11 - 1))
	__obf_54a1419c0ed802d9 := __obf_eddf34d110373ed9 >> 63

	if __obf_71f3d6c74e9e688b == 0 {
		// specials
		if __obf_26f3bdcd35e04593 == 0 {
			return Decimal{}
		}
		// subnormal
		__obf_71f3d6c74e9e688b++
	} else {
		// normal
		__obf_26f3bdcd35e04593 |= 1 << 52
	}

	__obf_71f3d6c74e9e688b -= 1023 + 52

	// normalizing base-2 values
	for __obf_26f3bdcd35e04593&1 == 0 {
		__obf_26f3bdcd35e04593 = __obf_26f3bdcd35e04593 >> 1
		__obf_71f3d6c74e9e688b++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_3e3f0912cbecfbf3 < 0 && __obf_3e3f0912cbecfbf3 < __obf_71f3d6c74e9e688b {
		if __obf_71f3d6c74e9e688b < 0 {
			__obf_3e3f0912cbecfbf3 = __obf_71f3d6c74e9e688b
		} else {
			__obf_3e3f0912cbecfbf3 = 0
		}
	}

	// representing 10^M * 2^N as 5^M * 2^(M+N)
	__obf_71f3d6c74e9e688b -= __obf_3e3f0912cbecfbf3

	__obf_b553773c069e3b75 := big.NewInt(1)
	__obf_b689a16adca06d79 := big.NewInt(int64(__obf_26f3bdcd35e04593))

	// applying 5^M
	if __obf_3e3f0912cbecfbf3 > 0 {
		__obf_b553773c069e3b75 = __obf_b553773c069e3b75.SetInt64(int64(__obf_3e3f0912cbecfbf3))
		__obf_b553773c069e3b75 = __obf_b553773c069e3b75.Exp(__obf_ac970d134d03d297, __obf_b553773c069e3b75, nil)
	} else if __obf_3e3f0912cbecfbf3 < 0 {
		__obf_b553773c069e3b75 = __obf_b553773c069e3b75.SetInt64(-int64(__obf_3e3f0912cbecfbf3))
		__obf_b553773c069e3b75 = __obf_b553773c069e3b75.Exp(__obf_ac970d134d03d297, __obf_b553773c069e3b75, nil)
		__obf_b689a16adca06d79 = __obf_b689a16adca06d79.Mul(__obf_b689a16adca06d79, __obf_b553773c069e3b75)
		__obf_b553773c069e3b75 = __obf_b553773c069e3b75.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_71f3d6c74e9e688b > 0 {
		__obf_b689a16adca06d79 = __obf_b689a16adca06d79.Lsh(__obf_b689a16adca06d79, uint(__obf_71f3d6c74e9e688b))
	} else if __obf_71f3d6c74e9e688b < 0 {
		__obf_b553773c069e3b75 = __obf_b553773c069e3b75.Lsh(__obf_b553773c069e3b75, uint(-__obf_71f3d6c74e9e688b))
	}

	// rounding and downscaling
	if __obf_3e3f0912cbecfbf3 > 0 || __obf_71f3d6c74e9e688b < 0 {
		__obf_4842d92e86b612c8 := new(big.Int).Rsh(__obf_b553773c069e3b75, 1)
		__obf_b689a16adca06d79 = __obf_b689a16adca06d79.Add(__obf_b689a16adca06d79, __obf_4842d92e86b612c8)
		__obf_b689a16adca06d79 = __obf_b689a16adca06d79.Quo(__obf_b689a16adca06d79, __obf_b553773c069e3b75)
	}

	if __obf_54a1419c0ed802d9 == 1 {
		__obf_b689a16adca06d79 = __obf_b689a16adca06d79.Neg(__obf_b689a16adca06d79)
	}

	return Decimal{
		__obf_04aabd427bafd586: __obf_b689a16adca06d79,
		__obf_3e3f0912cbecfbf3: __obf_3e3f0912cbecfbf3,
	}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_8a994c90f60ad404 Decimal) Copy() Decimal {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()
	return Decimal{
		__obf_04aabd427bafd586: new(big.Int).Set(__obf_8a994c90f60ad404.__obf_04aabd427bafd586),
		__obf_3e3f0912cbecfbf3: __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3,
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
func (__obf_8a994c90f60ad404 Decimal) __obf_f614630d6ca91251(__obf_3e3f0912cbecfbf3 int32) Decimal {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()

	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 == __obf_3e3f0912cbecfbf3 {
		return Decimal{
			new(big.Int).Set(__obf_8a994c90f60ad404.__obf_04aabd427bafd586),
			__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3,
		}
	}

	// NOTE(vadim): must convert exps to float64 before - to prevent overflow
	__obf_73e30d96abae8e1e := math.Abs(float64(__obf_3e3f0912cbecfbf3) - float64(__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3))
	__obf_04aabd427bafd586 := new(big.Int).Set(__obf_8a994c90f60ad404.__obf_04aabd427bafd586)

	__obf_6af1498cc0e1039e := new(big.Int).Exp(__obf_4a348d545ccd57bf, big.NewInt(int64(__obf_73e30d96abae8e1e)), nil)
	if __obf_3e3f0912cbecfbf3 > __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 {
		__obf_04aabd427bafd586 = __obf_04aabd427bafd586.Quo(__obf_04aabd427bafd586, __obf_6af1498cc0e1039e)
	} else if __obf_3e3f0912cbecfbf3 < __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 {
		__obf_04aabd427bafd586 = __obf_04aabd427bafd586.Mul(__obf_04aabd427bafd586, __obf_6af1498cc0e1039e)
	}

	return Decimal{
		__obf_04aabd427bafd586: __obf_04aabd427bafd586,
		__obf_3e3f0912cbecfbf3: __obf_3e3f0912cbecfbf3,
	}
}

// Abs returns the absolute value of the decimal.
func (__obf_8a994c90f60ad404 Decimal) Abs() Decimal {
	if !__obf_8a994c90f60ad404.IsNegative() {
		return __obf_8a994c90f60ad404
	}
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()
	__obf_527e3a75c860b953 := new(big.Int).Abs(__obf_8a994c90f60ad404.__obf_04aabd427bafd586)
	return Decimal{
		__obf_04aabd427bafd586: __obf_527e3a75c860b953,
		__obf_3e3f0912cbecfbf3: __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3,
	}
}

// Add returns d + d2.
func (__obf_8a994c90f60ad404 Decimal) Add(__obf_14eacb0a6cf534d0 Decimal) Decimal {
	__obf_40a67261674d8aeb, __obf_5ffe231c2020dfa3 := RescalePair(__obf_8a994c90f60ad404, __obf_14eacb0a6cf534d0)

	__obf_818d126a6efbc0f1 := new(big.Int).Add(__obf_40a67261674d8aeb.__obf_04aabd427bafd586, __obf_5ffe231c2020dfa3.__obf_04aabd427bafd586)
	return Decimal{
		__obf_04aabd427bafd586: __obf_818d126a6efbc0f1,
		__obf_3e3f0912cbecfbf3: __obf_40a67261674d8aeb.__obf_3e3f0912cbecfbf3,
	}
}

// Sub returns d - d2.
func (__obf_8a994c90f60ad404 Decimal) Sub(__obf_14eacb0a6cf534d0 Decimal) Decimal {
	__obf_40a67261674d8aeb, __obf_5ffe231c2020dfa3 := RescalePair(__obf_8a994c90f60ad404, __obf_14eacb0a6cf534d0)

	__obf_818d126a6efbc0f1 := new(big.Int).Sub(__obf_40a67261674d8aeb.__obf_04aabd427bafd586, __obf_5ffe231c2020dfa3.__obf_04aabd427bafd586)
	return Decimal{
		__obf_04aabd427bafd586: __obf_818d126a6efbc0f1,
		__obf_3e3f0912cbecfbf3: __obf_40a67261674d8aeb.__obf_3e3f0912cbecfbf3,
	}
}

// Neg returns -d.
func (__obf_8a994c90f60ad404 Decimal) Neg() Decimal {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()
	__obf_4d7c699332f4c605 := new(big.Int).Neg(__obf_8a994c90f60ad404.__obf_04aabd427bafd586)
	return Decimal{
		__obf_04aabd427bafd586: __obf_4d7c699332f4c605,
		__obf_3e3f0912cbecfbf3: __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3,
	}
}

// Mul returns d * d2.
func (__obf_8a994c90f60ad404 Decimal) Mul(__obf_14eacb0a6cf534d0 Decimal) Decimal {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()
	__obf_14eacb0a6cf534d0.__obf_90aaf6af73b21543()

	__obf_06dc95aee41225f3 := int64(__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3) + int64(__obf_14eacb0a6cf534d0.__obf_3e3f0912cbecfbf3)
	if __obf_06dc95aee41225f3 > math.MaxInt32 || __obf_06dc95aee41225f3 < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_06dc95aee41225f3))
	}

	__obf_818d126a6efbc0f1 := new(big.Int).Mul(__obf_8a994c90f60ad404.__obf_04aabd427bafd586, __obf_14eacb0a6cf534d0.__obf_04aabd427bafd586)
	return Decimal{
		__obf_04aabd427bafd586: __obf_818d126a6efbc0f1,
		__obf_3e3f0912cbecfbf3: int32(__obf_06dc95aee41225f3),
	}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_8a994c90f60ad404 Decimal) Shift(__obf_a3ff5c4d5644c72b int32) Decimal {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()
	return Decimal{
		__obf_04aabd427bafd586: new(big.Int).Set(__obf_8a994c90f60ad404.__obf_04aabd427bafd586),
		__obf_3e3f0912cbecfbf3: __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 + __obf_a3ff5c4d5644c72b,
	}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_8a994c90f60ad404 Decimal) Div(__obf_14eacb0a6cf534d0 Decimal) Decimal {
	return __obf_8a994c90f60ad404.DivRound(__obf_14eacb0a6cf534d0, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_8a994c90f60ad404 Decimal) QuoRem(__obf_14eacb0a6cf534d0 Decimal, __obf_c51dca51652f0a68 int32) (Decimal, Decimal) {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()
	__obf_14eacb0a6cf534d0.__obf_90aaf6af73b21543()
	if __obf_14eacb0a6cf534d0.__obf_04aabd427bafd586.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_512377f7f842cfc3 := -__obf_c51dca51652f0a68
	__obf_2832c1f1d480fa8e := int64(__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3) - int64(__obf_14eacb0a6cf534d0.__obf_3e3f0912cbecfbf3) - int64(__obf_512377f7f842cfc3)
	if __obf_2832c1f1d480fa8e > math.MaxInt32 || __obf_2832c1f1d480fa8e < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_e3d495b7bdbb1a3f, __obf_9a26f0b0872857aa, __obf_40bb7b9d68c51b43 big.Int
	var __obf_6eac871b50f7e631 int32
	// d = a 10^ea
	// d2 = b 10^eb
	if __obf_2832c1f1d480fa8e < 0 {
		__obf_e3d495b7bdbb1a3f = *__obf_8a994c90f60ad404.__obf_04aabd427bafd586
		__obf_40bb7b9d68c51b43.SetInt64(-__obf_2832c1f1d480fa8e)
		__obf_9a26f0b0872857aa.Exp(__obf_4a348d545ccd57bf, &__obf_40bb7b9d68c51b43, nil)
		__obf_9a26f0b0872857aa.Mul(__obf_14eacb0a6cf534d0.__obf_04aabd427bafd586, &__obf_9a26f0b0872857aa)
		__obf_6eac871b50f7e631 = __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3
		// now aa = a
		//     bb = b 10^(scale + eb - ea)
	} else {
		__obf_40bb7b9d68c51b43.SetInt64(__obf_2832c1f1d480fa8e)
		__obf_e3d495b7bdbb1a3f.Exp(__obf_4a348d545ccd57bf, &__obf_40bb7b9d68c51b43, nil)
		__obf_e3d495b7bdbb1a3f.Mul(__obf_8a994c90f60ad404.__obf_04aabd427bafd586, &__obf_e3d495b7bdbb1a3f)
		__obf_9a26f0b0872857aa = *__obf_14eacb0a6cf534d0.__obf_04aabd427bafd586
		__obf_6eac871b50f7e631 = __obf_512377f7f842cfc3 + __obf_14eacb0a6cf534d0.__obf_3e3f0912cbecfbf3
		// now aa = a ^ (ea - eb - scale)
		//     bb = b
	}
	var __obf_9a7e8237c0ecc118, __obf_a61a90a65de9f74c big.Int
	__obf_9a7e8237c0ecc118.QuoRem(&__obf_e3d495b7bdbb1a3f, &__obf_9a26f0b0872857aa, &__obf_a61a90a65de9f74c)
	__obf_27bf969a6d7c1094 := Decimal{__obf_04aabd427bafd586: &__obf_9a7e8237c0ecc118, __obf_3e3f0912cbecfbf3: __obf_512377f7f842cfc3}
	__obf_d689ea74359145d3 := Decimal{__obf_04aabd427bafd586: &__obf_a61a90a65de9f74c, __obf_3e3f0912cbecfbf3: __obf_6eac871b50f7e631}
	return __obf_27bf969a6d7c1094, __obf_d689ea74359145d3
}

// DivRound divides and rounds to a given precision
// i.e. to an integer multiple of 10^(-precision)
//
//	for a positive quotient digit 5 is rounded up, away from 0
//	if the quotient is negative then digit 5 is rounded down, away from 0
//
// Note that precision<0 is allowed as input.
func (__obf_8a994c90f60ad404 Decimal) DivRound(__obf_14eacb0a6cf534d0 Decimal, __obf_c51dca51652f0a68 int32) Decimal {
	// QuoRem already checks initialization
	__obf_9a7e8237c0ecc118, __obf_a61a90a65de9f74c := __obf_8a994c90f60ad404.QuoRem(__obf_14eacb0a6cf534d0, __obf_c51dca51652f0a68)
	// the actual rounding decision is based on comparing r*10^precision and d2/2
	// instead compare 2 r 10 ^precision and d2
	var __obf_41575a0debfedb74 big.Int
	__obf_41575a0debfedb74.Abs(__obf_a61a90a65de9f74c.__obf_04aabd427bafd586)
	__obf_41575a0debfedb74.Lsh(&__obf_41575a0debfedb74, 1)
	// now rv2 = abs(r.value) * 2
	__obf_d91f93e9388a102d := Decimal{__obf_04aabd427bafd586: &__obf_41575a0debfedb74, __obf_3e3f0912cbecfbf3: __obf_a61a90a65de9f74c.__obf_3e3f0912cbecfbf3 + __obf_c51dca51652f0a68}
	// r2 is now 2 * r * 10 ^ precision
	var __obf_82bb736e86a88e3a = __obf_d91f93e9388a102d.Cmp(__obf_14eacb0a6cf534d0.Abs())

	if __obf_82bb736e86a88e3a < 0 {
		return __obf_9a7e8237c0ecc118
	}

	if __obf_8a994c90f60ad404.__obf_04aabd427bafd586.Sign()*__obf_14eacb0a6cf534d0.__obf_04aabd427bafd586.Sign() < 0 {
		return __obf_9a7e8237c0ecc118.Sub(New(1, -__obf_c51dca51652f0a68))
	}

	return __obf_9a7e8237c0ecc118.Add(New(1, -__obf_c51dca51652f0a68))
}

// Mod returns d % d2.
func (__obf_8a994c90f60ad404 Decimal) Mod(__obf_14eacb0a6cf534d0 Decimal) Decimal {
	_, __obf_a61a90a65de9f74c := __obf_8a994c90f60ad404.QuoRem(__obf_14eacb0a6cf534d0, 0)
	return __obf_a61a90a65de9f74c
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
func (__obf_8a994c90f60ad404 Decimal) Pow(__obf_14eacb0a6cf534d0 Decimal) Decimal {
	__obf_40baf7ff8fcc0a8d := __obf_8a994c90f60ad404.Sign()
	__obf_541cf5499009a3e9 := __obf_14eacb0a6cf534d0.Sign()

	if __obf_40baf7ff8fcc0a8d == 0 {
		if __obf_541cf5499009a3e9 == 0 {
			return Decimal{}
		}
		if __obf_541cf5499009a3e9 == 1 {
			return Decimal{__obf_f9353e10ec98e7cd, 0}
		}
		if __obf_541cf5499009a3e9 == -1 {
			return Decimal{}
		}
	}

	if __obf_541cf5499009a3e9 == 0 {
		return Decimal{__obf_38f9fa21a076c52c, 0}
	}

	// TODO: optimize extraction of fractional part
	__obf_734773e2a8ed9719 := Decimal{__obf_38f9fa21a076c52c, 0}
	__obf_863119a3079a25e4, __obf_e5d6d59b7e387f65 := __obf_14eacb0a6cf534d0.QuoRem(__obf_734773e2a8ed9719, 0)

	if __obf_40baf7ff8fcc0a8d == -1 && !__obf_e5d6d59b7e387f65.IsZero() {
		return Decimal{}
	}

	__obf_036f5ce65ebf609d, _ := __obf_8a994c90f60ad404.PowBigInt(__obf_863119a3079a25e4.__obf_04aabd427bafd586)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_e5d6d59b7e387f65.__obf_04aabd427bafd586.Sign() == 0 {
		return __obf_036f5ce65ebf609d
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_9193c83dd34b8218 := __obf_8a994c90f60ad404.NumDigits()
	__obf_bb417b83731cfefa := __obf_14eacb0a6cf534d0.NumDigits()

	__obf_c51dca51652f0a68 := __obf_9193c83dd34b8218

	if __obf_bb417b83731cfefa > __obf_c51dca51652f0a68 {
		__obf_c51dca51652f0a68 += __obf_bb417b83731cfefa
	}

	__obf_c51dca51652f0a68 += 6

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_72b5e03bbff9234e, __obf_8d1e8d31f090b7f5 := __obf_8a994c90f60ad404.Abs().Ln(-__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 + int32(__obf_c51dca51652f0a68))
	if __obf_8d1e8d31f090b7f5 != nil {
		return Decimal{}
	}

	__obf_72b5e03bbff9234e = __obf_72b5e03bbff9234e.Mul(__obf_e5d6d59b7e387f65)

	__obf_72b5e03bbff9234e, __obf_8d1e8d31f090b7f5 = __obf_72b5e03bbff9234e.ExpTaylor(-__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 + int32(__obf_c51dca51652f0a68))
	if __obf_8d1e8d31f090b7f5 != nil {
		return Decimal{}
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_4a2d7558c41c4a84 := __obf_036f5ce65ebf609d.Mul(__obf_72b5e03bbff9234e)

	return __obf_4a2d7558c41c4a84
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
func (__obf_8a994c90f60ad404 Decimal) PowWithPrecision(__obf_14eacb0a6cf534d0 Decimal, __obf_c51dca51652f0a68 int32) (Decimal, error) {
	__obf_40baf7ff8fcc0a8d := __obf_8a994c90f60ad404.Sign()
	__obf_541cf5499009a3e9 := __obf_14eacb0a6cf534d0.Sign()

	if __obf_40baf7ff8fcc0a8d == 0 {
		if __obf_541cf5499009a3e9 == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_541cf5499009a3e9 == 1 {
			return Decimal{__obf_f9353e10ec98e7cd, 0}, nil
		}
		if __obf_541cf5499009a3e9 == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_541cf5499009a3e9 == 0 {
		return Decimal{__obf_38f9fa21a076c52c, 0}, nil
	}

	// TODO: optimize extraction of fractional part
	__obf_734773e2a8ed9719 := Decimal{__obf_38f9fa21a076c52c, 0}
	__obf_863119a3079a25e4, __obf_e5d6d59b7e387f65 := __obf_14eacb0a6cf534d0.QuoRem(__obf_734773e2a8ed9719, 0)

	if __obf_40baf7ff8fcc0a8d == -1 && !__obf_e5d6d59b7e387f65.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}

	__obf_036f5ce65ebf609d, _ := __obf_8a994c90f60ad404.__obf_f4f6557ee687fa80(__obf_863119a3079a25e4.__obf_04aabd427bafd586, __obf_c51dca51652f0a68)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_e5d6d59b7e387f65.__obf_04aabd427bafd586.Sign() == 0 {
		return __obf_036f5ce65ebf609d, nil
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_9193c83dd34b8218 := __obf_8a994c90f60ad404.NumDigits()
	__obf_bb417b83731cfefa := __obf_14eacb0a6cf534d0.NumDigits()

	if int32(__obf_9193c83dd34b8218) > __obf_c51dca51652f0a68 {
		__obf_c51dca51652f0a68 = int32(__obf_9193c83dd34b8218)
	}
	if int32(__obf_bb417b83731cfefa) > __obf_c51dca51652f0a68 {
		__obf_c51dca51652f0a68 += int32(__obf_bb417b83731cfefa)
	}
	// increase precision by 10 to compensate for errors in further calculations
	__obf_c51dca51652f0a68 += 10

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_72b5e03bbff9234e, __obf_8d1e8d31f090b7f5 := __obf_8a994c90f60ad404.Abs().Ln(__obf_c51dca51652f0a68)
	if __obf_8d1e8d31f090b7f5 != nil {
		return Decimal{}, __obf_8d1e8d31f090b7f5
	}

	__obf_72b5e03bbff9234e = __obf_72b5e03bbff9234e.Mul(__obf_e5d6d59b7e387f65)

	__obf_72b5e03bbff9234e, __obf_8d1e8d31f090b7f5 = __obf_72b5e03bbff9234e.ExpTaylor(__obf_c51dca51652f0a68)
	if __obf_8d1e8d31f090b7f5 != nil {
		return Decimal{}, __obf_8d1e8d31f090b7f5
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_4a2d7558c41c4a84 := __obf_036f5ce65ebf609d.Mul(__obf_72b5e03bbff9234e)

	return __obf_4a2d7558c41c4a84, nil
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
func (__obf_8a994c90f60ad404 Decimal) PowInt32(__obf_3e3f0912cbecfbf3 int32) (Decimal, error) {
	if __obf_8a994c90f60ad404.IsZero() && __obf_3e3f0912cbecfbf3 == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_e9a34ee74ab0e190 := __obf_3e3f0912cbecfbf3 < 0
	__obf_3e3f0912cbecfbf3 = __obf_d1bb30f067c5e1eb(__obf_3e3f0912cbecfbf3)

	__obf_ea676f25508dcb39, __obf_7bd66370dd4f01aa := __obf_8a994c90f60ad404, New(1, 0)

	for __obf_3e3f0912cbecfbf3 > 0 {
		if __obf_3e3f0912cbecfbf3%2 == 1 {
			__obf_7bd66370dd4f01aa = __obf_7bd66370dd4f01aa.Mul(__obf_ea676f25508dcb39)
		}
		__obf_3e3f0912cbecfbf3 /= 2

		if __obf_3e3f0912cbecfbf3 > 0 {
			__obf_ea676f25508dcb39 = __obf_ea676f25508dcb39.Mul(__obf_ea676f25508dcb39)
		}
	}

	if __obf_e9a34ee74ab0e190 {
		return New(1, 0).DivRound(__obf_7bd66370dd4f01aa, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_7bd66370dd4f01aa, nil
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
func (__obf_8a994c90f60ad404 Decimal) PowBigInt(__obf_3e3f0912cbecfbf3 *big.Int) (Decimal, error) {
	return __obf_8a994c90f60ad404.__obf_f4f6557ee687fa80(__obf_3e3f0912cbecfbf3, int32(PowPrecisionNegativeExponent))
}

func (__obf_8a994c90f60ad404 Decimal) __obf_f4f6557ee687fa80(__obf_3e3f0912cbecfbf3 *big.Int, __obf_c51dca51652f0a68 int32) (Decimal, error) {
	if __obf_8a994c90f60ad404.IsZero() && __obf_3e3f0912cbecfbf3.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_78bbe3860d2d759f := new(big.Int).Set(__obf_3e3f0912cbecfbf3)
	__obf_e9a34ee74ab0e190 := __obf_3e3f0912cbecfbf3.Sign() < 0

	if __obf_e9a34ee74ab0e190 {
		__obf_78bbe3860d2d759f.Abs(__obf_78bbe3860d2d759f)
	}

	__obf_ea676f25508dcb39, __obf_7bd66370dd4f01aa := __obf_8a994c90f60ad404, New(1, 0)

	for __obf_78bbe3860d2d759f.Sign() > 0 {
		if __obf_78bbe3860d2d759f.Bit(0) == 1 {
			__obf_7bd66370dd4f01aa = __obf_7bd66370dd4f01aa.Mul(__obf_ea676f25508dcb39)
		}
		__obf_78bbe3860d2d759f.Rsh(__obf_78bbe3860d2d759f, 1)

		if __obf_78bbe3860d2d759f.Sign() > 0 {
			__obf_ea676f25508dcb39 = __obf_ea676f25508dcb39.Mul(__obf_ea676f25508dcb39)
		}
	}

	if __obf_e9a34ee74ab0e190 {
		return New(1, 0).DivRound(__obf_7bd66370dd4f01aa, __obf_c51dca51652f0a68), nil
	}

	return __obf_7bd66370dd4f01aa, nil
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
func (__obf_8a994c90f60ad404 Decimal) ExpHullAbrham(__obf_4d778d3eef3a5ac9 uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_8a994c90f60ad404.IsZero() {
		return Decimal{__obf_38f9fa21a076c52c, 0}, nil
	}

	__obf_a6b8661d3f6961f3 := __obf_4d778d3eef3a5ac9

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_f2668f829072e434 := __obf_8a994c90f60ad404.Abs().InexactFloat64()
	if __obf_65b7482c8167c5ce := __obf_f2668f829072e434 / 23; __obf_65b7482c8167c5ce > float64(__obf_a6b8661d3f6961f3) && __obf_65b7482c8167c5ce < float64(ExpMaxIterations) {
		__obf_a6b8661d3f6961f3 = uint32(math.Ceil(__obf_65b7482c8167c5ce))
	}

	// fail if abs(d) beyond an over/underflow threshold
	__obf_902a12b8a766b8dd := New(23*int64(__obf_a6b8661d3f6961f3), 0)
	if __obf_8a994c90f60ad404.Abs().Cmp(__obf_902a12b8a766b8dd) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}

	// Return 1 if abs(d) small enough; this also avoids later over/underflow
	__obf_a5ec84c2fca190f7 := New(9, -int32(__obf_a6b8661d3f6961f3)-1)
	if __obf_8a994c90f60ad404.Abs().Cmp(__obf_a5ec84c2fca190f7) <= 0 {
		return Decimal{__obf_38f9fa21a076c52c, __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3}, nil
	}

	// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
	__obf_1a0a96d61f3a8378 := __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 + int32(__obf_8a994c90f60ad404.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_1a0a96d61f3a8378 < 0 {
		__obf_1a0a96d61f3a8378 = 0
	}

	__obf_27199f055c617c49 := New(1, __obf_1a0a96d61f3a8378)                                                                                                                   // reduction factor
	__obf_a61a90a65de9f74c := Decimal{new(big.Int).Set(__obf_8a994c90f60ad404.__obf_04aabd427bafd586), __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 - __obf_1a0a96d61f3a8378} // reduced argument
	__obf_3e70371e552b2926 := int32(__obf_a6b8661d3f6961f3) + __obf_1a0a96d61f3a8378 + 2                                                                                       // precision for calculating the sum

	// Determine n, the number of therms for calculating sum
	// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
	// for solving appropriate equation, along with directed
	// roundings and simple rational bound for log10(p/abs(r))
	__obf_cb3fe21cd5dbabc9 := __obf_a61a90a65de9f74c.Abs().InexactFloat64()
	__obf_2b9cfa9f4c63e652 := float64(__obf_3e70371e552b2926)
	__obf_13e9abb47c88870e := math.Ceil((1.453*__obf_2b9cfa9f4c63e652 - 1.182) / math.Log10(__obf_2b9cfa9f4c63e652/__obf_cb3fe21cd5dbabc9))
	if __obf_13e9abb47c88870e > float64(ExpMaxIterations) || math.IsNaN(__obf_13e9abb47c88870e) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_ea676f25508dcb39 := int64(__obf_13e9abb47c88870e)

	__obf_4bbdbb28098ff3a4 := New(0, 0)
	__obf_691733ca8f4c39c1 := New(1, 0)
	__obf_734773e2a8ed9719 := New(1, 0)
	for __obf_bf8f4a859c474f4d := __obf_ea676f25508dcb39 - 1; __obf_bf8f4a859c474f4d > 0; __obf_bf8f4a859c474f4d-- {
		__obf_4bbdbb28098ff3a4.__obf_04aabd427bafd586.SetInt64(__obf_bf8f4a859c474f4d)
		__obf_691733ca8f4c39c1 = __obf_691733ca8f4c39c1.Mul(__obf_a61a90a65de9f74c.DivRound(__obf_4bbdbb28098ff3a4, __obf_3e70371e552b2926))
		__obf_691733ca8f4c39c1 = __obf_691733ca8f4c39c1.Add(__obf_734773e2a8ed9719)
	}

	__obf_fcfbf9daaf83aff2 := __obf_27199f055c617c49.IntPart()
	__obf_4a2d7558c41c4a84 := New(1, 0)
	for __obf_bf8f4a859c474f4d := __obf_fcfbf9daaf83aff2; __obf_bf8f4a859c474f4d > 0; __obf_bf8f4a859c474f4d-- {
		__obf_4a2d7558c41c4a84 = __obf_4a2d7558c41c4a84.Mul(__obf_691733ca8f4c39c1)
	}

	__obf_9f79c267fc4a4404 := int32(__obf_4a2d7558c41c4a84.NumDigits())

	var __obf_6f2f248ad93258c9 int32
	if __obf_9f79c267fc4a4404 > __obf_d1bb30f067c5e1eb(__obf_4a2d7558c41c4a84.__obf_3e3f0912cbecfbf3) {
		__obf_6f2f248ad93258c9 = int32(__obf_a6b8661d3f6961f3) - __obf_9f79c267fc4a4404 - __obf_4a2d7558c41c4a84.__obf_3e3f0912cbecfbf3
	} else {
		__obf_6f2f248ad93258c9 = int32(__obf_a6b8661d3f6961f3)
	}

	__obf_4a2d7558c41c4a84 = __obf_4a2d7558c41c4a84.Round(__obf_6f2f248ad93258c9)

	return __obf_4a2d7558c41c4a84, nil
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
func (__obf_8a994c90f60ad404 Decimal) ExpTaylor(__obf_c51dca51652f0a68 int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_8a994c90f60ad404.IsZero() {
		return Decimal{__obf_38f9fa21a076c52c, 0}.Round(__obf_c51dca51652f0a68), nil
	}

	var __obf_fc0485686ed36867 Decimal
	var __obf_6ad2eeaed50ba607 int32
	if __obf_c51dca51652f0a68 < 0 {
		__obf_fc0485686ed36867 = New(1, -1)
		__obf_6ad2eeaed50ba607 = 8
	} else {
		__obf_fc0485686ed36867 = New(1, -__obf_c51dca51652f0a68-1)
		__obf_6ad2eeaed50ba607 = __obf_c51dca51652f0a68 + 1
	}

	__obf_523614d833771815 := __obf_8a994c90f60ad404.Abs()
	__obf_fc33a22417ce6c4d := __obf_8a994c90f60ad404.Abs()
	__obf_f8c8eecc413389f0 := New(1, 0)

	__obf_7bd66370dd4f01aa := New(1, 0)

	for __obf_bf8f4a859c474f4d := int64(1); ; {
		__obf_02163a811a15f548 := __obf_fc33a22417ce6c4d.DivRound(__obf_f8c8eecc413389f0, __obf_6ad2eeaed50ba607)
		__obf_7bd66370dd4f01aa = __obf_7bd66370dd4f01aa.Add(__obf_02163a811a15f548)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_02163a811a15f548.Cmp(__obf_fc0485686ed36867) < 0 {
			break
		}

		__obf_fc33a22417ce6c4d = __obf_fc33a22417ce6c4d.Mul(__obf_523614d833771815)

		__obf_bf8f4a859c474f4d++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_a71570cb39e71d12) >= int(__obf_bf8f4a859c474f4d) && !__obf_a71570cb39e71d12[__obf_bf8f4a859c474f4d-1].IsZero() {
			__obf_f8c8eecc413389f0 = __obf_a71570cb39e71d12[__obf_bf8f4a859c474f4d-1]
		} else {
			// To avoid any race conditions, firstly the zero value is appended to a slice to create
			// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
			// factorial using the index notation.
			__obf_f8c8eecc413389f0 = __obf_a71570cb39e71d12[__obf_bf8f4a859c474f4d-2].Mul(New(__obf_bf8f4a859c474f4d, 0))
			__obf_a71570cb39e71d12 = append(__obf_a71570cb39e71d12, Zero)
			__obf_a71570cb39e71d12[__obf_bf8f4a859c474f4d-1] = __obf_f8c8eecc413389f0
		}
	}

	if __obf_8a994c90f60ad404.Sign() < 0 {
		__obf_7bd66370dd4f01aa = New(1, 0).DivRound(__obf_7bd66370dd4f01aa, __obf_c51dca51652f0a68+1)
	}

	__obf_7bd66370dd4f01aa = __obf_7bd66370dd4f01aa.Round(__obf_c51dca51652f0a68)
	return __obf_7bd66370dd4f01aa, nil
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
func (__obf_8a994c90f60ad404 Decimal) Ln(__obf_c51dca51652f0a68 int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_8a994c90f60ad404.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_8a994c90f60ad404.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}

	__obf_4b09eb070bb29fe7 := __obf_c51dca51652f0a68 + 2
	__obf_6ac07a49686a1017 := __obf_8a994c90f60ad404.Copy()

	var __obf_202934d5348a4981, __obf_9b8fdf50686333f1, __obf_4951045606e0dedb, __obf_acf591e11b749f46, __obf_db1860422526c57f Decimal
	__obf_202934d5348a4981 = __obf_6ac07a49686a1017.Sub(Decimal{__obf_38f9fa21a076c52c, 0})
	__obf_9b8fdf50686333f1 = Decimal{__obf_38f9fa21a076c52c, -1}

	// for decimal in range [0.9, 1.1] where ln(d) is close to 0
	__obf_d5d7f4cc27185f30 := false

	if __obf_202934d5348a4981.Abs().Cmp(__obf_9b8fdf50686333f1) <= 0 {
		__obf_d5d7f4cc27185f30 = true
	} else {
		// reduce input decimal to range [0.1, 1)
		__obf_510f5cd3dcb26e5a := int32(__obf_6ac07a49686a1017.NumDigits()) + __obf_6ac07a49686a1017.__obf_3e3f0912cbecfbf3
		__obf_6ac07a49686a1017.__obf_3e3f0912cbecfbf3 -= __obf_510f5cd3dcb26e5a

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_927e9bc47c799b14 := __obf_927e9bc47c799b14.__obf_dec2b65a8a915614(__obf_4b09eb070bb29fe7)
		__obf_db1860422526c57f = NewFromInt32(__obf_510f5cd3dcb26e5a)
		__obf_db1860422526c57f = __obf_db1860422526c57f.Mul(__obf_927e9bc47c799b14)

		__obf_202934d5348a4981 = __obf_6ac07a49686a1017.Sub(Decimal{__obf_38f9fa21a076c52c, 0})

		if __obf_202934d5348a4981.Abs().Cmp(__obf_9b8fdf50686333f1) <= 0 {
			__obf_d5d7f4cc27185f30 = true
		} else {
			// initial estimate using floats
			__obf_31bcb6f928f3a0ab := __obf_6ac07a49686a1017.InexactFloat64()
			__obf_202934d5348a4981 = NewFromFloat(math.Log(__obf_31bcb6f928f3a0ab))
		}
	}

	__obf_fc0485686ed36867 := Decimal{__obf_38f9fa21a076c52c, -__obf_4b09eb070bb29fe7}

	if __obf_d5d7f4cc27185f30 {
		// Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
		// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
		// until the difference between current and next term is smaller than epsilon.
		// Coverage quite fast for decimals close to 1.0

		// z + 2
		__obf_4951045606e0dedb = __obf_202934d5348a4981.Add(Decimal{__obf_e9d2424d8d05f51b, 0})
		// z / (z + 2)
		__obf_9b8fdf50686333f1 = __obf_202934d5348a4981.DivRound(__obf_4951045606e0dedb, __obf_4b09eb070bb29fe7)
		// 2 * (z / (z + 2))
		__obf_202934d5348a4981 = __obf_9b8fdf50686333f1.Add(__obf_9b8fdf50686333f1)
		__obf_4951045606e0dedb = __obf_202934d5348a4981.Copy()

		for __obf_ea676f25508dcb39 := 1; ; __obf_ea676f25508dcb39++ {
			// 2 * (z / (z+2))^(2n+1)
			__obf_4951045606e0dedb = __obf_4951045606e0dedb.Mul(__obf_9b8fdf50686333f1).Mul(__obf_9b8fdf50686333f1)

			// 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
			__obf_acf591e11b749f46 = NewFromInt(int64(2*__obf_ea676f25508dcb39 + 1))
			__obf_acf591e11b749f46 = __obf_4951045606e0dedb.DivRound(__obf_acf591e11b749f46, __obf_4b09eb070bb29fe7)

			// comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			__obf_202934d5348a4981 = __obf_202934d5348a4981.Add(__obf_acf591e11b749f46)

			if __obf_acf591e11b749f46.Abs().Cmp(__obf_fc0485686ed36867) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_d124630e9c0b3359 Decimal
		__obf_8f1704694623a711 := __obf_4b09eb070bb29fe7*2 + 10

		for __obf_bf8f4a859c474f4d := int32(0); __obf_bf8f4a859c474f4d < __obf_8f1704694623a711; __obf_bf8f4a859c474f4d++ {
			// exp(a_n)
			__obf_9b8fdf50686333f1, _ = __obf_202934d5348a4981.ExpTaylor(__obf_4b09eb070bb29fe7)
			// exp(a_n) - z
			__obf_4951045606e0dedb = __obf_9b8fdf50686333f1.Sub(__obf_6ac07a49686a1017)
			// 2 * (exp(a_n) - z)
			__obf_4951045606e0dedb = __obf_4951045606e0dedb.Add(__obf_4951045606e0dedb)
			// exp(a_n) + z
			__obf_acf591e11b749f46 = __obf_9b8fdf50686333f1.Add(__obf_6ac07a49686a1017)
			// 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_9b8fdf50686333f1 = __obf_4951045606e0dedb.DivRound(__obf_acf591e11b749f46, __obf_4b09eb070bb29fe7)
			// comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_202934d5348a4981 = __obf_202934d5348a4981.Sub(__obf_9b8fdf50686333f1)

			if __obf_d124630e9c0b3359.Add(__obf_9b8fdf50686333f1).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_9b8fdf50686333f1.Abs().Cmp(__obf_fc0485686ed36867) <= 0 {
				break
			}

			__obf_d124630e9c0b3359 = __obf_9b8fdf50686333f1
		}
	}

	__obf_202934d5348a4981 = __obf_202934d5348a4981.Add(__obf_db1860422526c57f)

	return __obf_202934d5348a4981.Round(__obf_c51dca51652f0a68), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_8a994c90f60ad404 Decimal) NumDigits() int {
	if __obf_8a994c90f60ad404.__obf_04aabd427bafd586 == nil {
		return 1
	}

	if __obf_8a994c90f60ad404.__obf_04aabd427bafd586.IsInt64() {
		__obf_1bd7ef64de055664 := __obf_8a994c90f60ad404.__obf_04aabd427bafd586.Int64()
		// restrict fast path to integers with exact conversion to float64
		if __obf_1bd7ef64de055664 <= (1<<53) && __obf_1bd7ef64de055664 >= -(1<<53) {
			if __obf_1bd7ef64de055664 == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_1bd7ef64de055664)))) + 1
		}
	}

	__obf_fc2198b5e39a1d73 := int(float64(__obf_8a994c90f60ad404.__obf_04aabd427bafd586.BitLen()) / math.Log2(10))

	// estimatedNumDigits (lg10) may be off by 1, need to verify
	__obf_e6bacc48734402d4 := big.NewInt(int64(__obf_fc2198b5e39a1d73))
	__obf_2c3c86514bc3f38e := __obf_e6bacc48734402d4.Exp(__obf_4a348d545ccd57bf, __obf_e6bacc48734402d4, nil)

	if __obf_8a994c90f60ad404.__obf_04aabd427bafd586.CmpAbs(__obf_2c3c86514bc3f38e) >= 0 {
		return __obf_fc2198b5e39a1d73 + 1
	}

	return __obf_fc2198b5e39a1d73
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_8a994c90f60ad404 Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_a61a90a65de9f74c big.Int
	__obf_9a7e8237c0ecc118 := new(big.Int).Set(__obf_8a994c90f60ad404.__obf_04aabd427bafd586)
	for __obf_6ac07a49686a1017 := __obf_d1bb30f067c5e1eb(__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3); __obf_6ac07a49686a1017 > 0; __obf_6ac07a49686a1017-- {
		__obf_9a7e8237c0ecc118.QuoRem(__obf_9a7e8237c0ecc118, __obf_4a348d545ccd57bf, &__obf_a61a90a65de9f74c)
		if __obf_a61a90a65de9f74c.Cmp(__obf_f9353e10ec98e7cd) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_d1bb30f067c5e1eb(__obf_ea676f25508dcb39 int32) int32 {
	if __obf_ea676f25508dcb39 < 0 {
		return -__obf_ea676f25508dcb39
	}
	return __obf_ea676f25508dcb39
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_8a994c90f60ad404 Decimal) Cmp(__obf_14eacb0a6cf534d0 Decimal) int {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()
	__obf_14eacb0a6cf534d0.__obf_90aaf6af73b21543()

	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 == __obf_14eacb0a6cf534d0.__obf_3e3f0912cbecfbf3 {
		return __obf_8a994c90f60ad404.__obf_04aabd427bafd586.Cmp(__obf_14eacb0a6cf534d0.__obf_04aabd427bafd586)
	}

	__obf_40a67261674d8aeb, __obf_5ffe231c2020dfa3 := RescalePair(__obf_8a994c90f60ad404, __obf_14eacb0a6cf534d0)

	return __obf_40a67261674d8aeb.__obf_04aabd427bafd586.Cmp(__obf_5ffe231c2020dfa3.__obf_04aabd427bafd586)
}

// Compare compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_8a994c90f60ad404 Decimal) Compare(__obf_14eacb0a6cf534d0 Decimal) int {
	return __obf_8a994c90f60ad404.Cmp(__obf_14eacb0a6cf534d0)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_8a994c90f60ad404 Decimal) Equal(__obf_14eacb0a6cf534d0 Decimal) bool {
	return __obf_8a994c90f60ad404.Cmp(__obf_14eacb0a6cf534d0) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_8a994c90f60ad404 Decimal) Equals(__obf_14eacb0a6cf534d0 Decimal) bool {
	return __obf_8a994c90f60ad404.Equal(__obf_14eacb0a6cf534d0)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_8a994c90f60ad404 Decimal) GreaterThan(__obf_14eacb0a6cf534d0 Decimal) bool {
	return __obf_8a994c90f60ad404.Cmp(__obf_14eacb0a6cf534d0) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_8a994c90f60ad404 Decimal) GreaterThanOrEqual(__obf_14eacb0a6cf534d0 Decimal) bool {
	__obf_402cff7ed14536cf := __obf_8a994c90f60ad404.Cmp(__obf_14eacb0a6cf534d0)
	return __obf_402cff7ed14536cf == 1 || __obf_402cff7ed14536cf == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_8a994c90f60ad404 Decimal) LessThan(__obf_14eacb0a6cf534d0 Decimal) bool {
	return __obf_8a994c90f60ad404.Cmp(__obf_14eacb0a6cf534d0) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_8a994c90f60ad404 Decimal) LessThanOrEqual(__obf_14eacb0a6cf534d0 Decimal) bool {
	__obf_402cff7ed14536cf := __obf_8a994c90f60ad404.Cmp(__obf_14eacb0a6cf534d0)
	return __obf_402cff7ed14536cf == -1 || __obf_402cff7ed14536cf == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_8a994c90f60ad404 Decimal) Sign() int {
	if __obf_8a994c90f60ad404.__obf_04aabd427bafd586 == nil {
		return 0
	}
	return __obf_8a994c90f60ad404.__obf_04aabd427bafd586.Sign()
}

// IsPositive return
//
//	true if d > 0
//	false if d == 0
//	false if d < 0
func (__obf_8a994c90f60ad404 Decimal) IsPositive() bool {
	return __obf_8a994c90f60ad404.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_8a994c90f60ad404 Decimal) IsNegative() bool {
	return __obf_8a994c90f60ad404.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_8a994c90f60ad404 Decimal) IsZero() bool {
	return __obf_8a994c90f60ad404.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_8a994c90f60ad404 Decimal) Exponent() int32 {
	return __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3
}

// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
func (__obf_8a994c90f60ad404 Decimal) Coefficient() *big.Int {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_8a994c90f60ad404.__obf_04aabd427bafd586)
}

// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
// If coefficient cannot be represented in an int64, the result will be undefined.
func (__obf_8a994c90f60ad404 Decimal) CoefficientInt64() int64 {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()
	return __obf_8a994c90f60ad404.__obf_04aabd427bafd586.Int64()
}

// IntPart returns the integer component of the decimal.
func (__obf_8a994c90f60ad404 Decimal) IntPart() int64 {
	__obf_f456f54bd8c52f5a := __obf_8a994c90f60ad404.__obf_f614630d6ca91251(0)
	return __obf_f456f54bd8c52f5a.__obf_04aabd427bafd586.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_8a994c90f60ad404 Decimal) BigInt() *big.Int {
	__obf_f456f54bd8c52f5a := __obf_8a994c90f60ad404.__obf_f614630d6ca91251(0)
	return __obf_f456f54bd8c52f5a.__obf_04aabd427bafd586
}

// BigFloat returns decimal as BigFloat.
// Be aware that casting decimal to BigFloat might cause a loss of precision.
func (__obf_8a994c90f60ad404 Decimal) BigFloat() *big.Float {
	__obf_f2668f829072e434 := &big.Float{}
	__obf_f2668f829072e434.SetString(__obf_8a994c90f60ad404.String())
	return __obf_f2668f829072e434
}

// Rat returns a rational number representation of the decimal.
func (__obf_8a994c90f60ad404 Decimal) Rat() *big.Rat {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()
	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 <= 0 {
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_43ca97e7a3313558 := new(big.Int).Exp(__obf_4a348d545ccd57bf, big.NewInt(-int64(__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3)), nil)
		return new(big.Rat).SetFrac(__obf_8a994c90f60ad404.__obf_04aabd427bafd586, __obf_43ca97e7a3313558)
	}

	__obf_2d14997694bc0709 := new(big.Int).Exp(__obf_4a348d545ccd57bf, big.NewInt(int64(__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3)), nil)
	__obf_3fc5ece3729ae968 := new(big.Int).Mul(__obf_8a994c90f60ad404.__obf_04aabd427bafd586, __obf_2d14997694bc0709)
	return new(big.Rat).SetFrac(__obf_3fc5ece3729ae968, __obf_38f9fa21a076c52c)
}

// Float64 returns the nearest float64 value for d and a bool indicating
// whether f represents d exactly.
// For more details, see the documentation for big.Rat.Float64
func (__obf_8a994c90f60ad404 Decimal) Float64() (__obf_f2668f829072e434 float64, __obf_702ca521c7cb7d88 bool) {
	return __obf_8a994c90f60ad404.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_8a994c90f60ad404 Decimal) InexactFloat64() float64 {
	__obf_f2668f829072e434, _ := __obf_8a994c90f60ad404.Float64()
	return __obf_f2668f829072e434
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
func (__obf_8a994c90f60ad404 Decimal) String() string {
	return __obf_8a994c90f60ad404.string(true)
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
func (__obf_8a994c90f60ad404 Decimal) StringFixed(__obf_3277a2f5cc1a9377 int32) string {
	__obf_443e7c61f45bcc55 := __obf_8a994c90f60ad404.Round(__obf_3277a2f5cc1a9377)
	return __obf_443e7c61f45bcc55.string(false)
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
func (__obf_8a994c90f60ad404 Decimal) StringFixedBank(__obf_3277a2f5cc1a9377 int32) string {
	__obf_443e7c61f45bcc55 := __obf_8a994c90f60ad404.RoundBank(__obf_3277a2f5cc1a9377)
	return __obf_443e7c61f45bcc55.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_8a994c90f60ad404 Decimal) StringFixedCash(__obf_c208b02a41f708c1 uint8) string {
	__obf_443e7c61f45bcc55 := __obf_8a994c90f60ad404.RoundCash(__obf_c208b02a41f708c1)
	return __obf_443e7c61f45bcc55.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_8a994c90f60ad404 Decimal) Round(__obf_3277a2f5cc1a9377 int32) Decimal {
	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 == -__obf_3277a2f5cc1a9377 {
		return __obf_8a994c90f60ad404
	}
	// truncate to places + 1
	__obf_3400e90a81b5e878 := __obf_8a994c90f60ad404.__obf_f614630d6ca91251(-__obf_3277a2f5cc1a9377 - 1)

	// add sign(d) * 0.5
	if __obf_3400e90a81b5e878.__obf_04aabd427bafd586.Sign() < 0 {
		__obf_3400e90a81b5e878.__obf_04aabd427bafd586.Sub(__obf_3400e90a81b5e878.__obf_04aabd427bafd586, __obf_ac970d134d03d297)
	} else {
		__obf_3400e90a81b5e878.__obf_04aabd427bafd586.Add(__obf_3400e90a81b5e878.__obf_04aabd427bafd586, __obf_ac970d134d03d297)
	}

	// floor for positive numbers, ceil for negative numbers
	_, __obf_1a0547a84e8f209a := __obf_3400e90a81b5e878.__obf_04aabd427bafd586.DivMod(__obf_3400e90a81b5e878.__obf_04aabd427bafd586, __obf_4a348d545ccd57bf, new(big.Int))
	__obf_3400e90a81b5e878.__obf_3e3f0912cbecfbf3++
	if __obf_3400e90a81b5e878.__obf_04aabd427bafd586.Sign() < 0 && __obf_1a0547a84e8f209a.Cmp(__obf_f9353e10ec98e7cd) != 0 {
		__obf_3400e90a81b5e878.__obf_04aabd427bafd586.Add(__obf_3400e90a81b5e878.__obf_04aabd427bafd586, __obf_38f9fa21a076c52c)
	}

	return __obf_3400e90a81b5e878
}

// RoundCeil rounds the decimal towards +infinity.
//
// Example:
//
//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
func (__obf_8a994c90f60ad404 Decimal) RoundCeil(__obf_3277a2f5cc1a9377 int32) Decimal {
	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 >= -__obf_3277a2f5cc1a9377 {
		return __obf_8a994c90f60ad404
	}

	__obf_b5ef24df54cf4dce := __obf_8a994c90f60ad404.__obf_f614630d6ca91251(-__obf_3277a2f5cc1a9377)
	if __obf_8a994c90f60ad404.Equal(__obf_b5ef24df54cf4dce) {
		return __obf_8a994c90f60ad404
	}

	if __obf_8a994c90f60ad404.__obf_04aabd427bafd586.Sign() > 0 {
		__obf_b5ef24df54cf4dce.__obf_04aabd427bafd586.Add(__obf_b5ef24df54cf4dce.__obf_04aabd427bafd586, __obf_38f9fa21a076c52c)
	}

	return __obf_b5ef24df54cf4dce
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_8a994c90f60ad404 Decimal) RoundFloor(__obf_3277a2f5cc1a9377 int32) Decimal {
	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 >= -__obf_3277a2f5cc1a9377 {
		return __obf_8a994c90f60ad404
	}

	__obf_b5ef24df54cf4dce := __obf_8a994c90f60ad404.__obf_f614630d6ca91251(-__obf_3277a2f5cc1a9377)
	if __obf_8a994c90f60ad404.Equal(__obf_b5ef24df54cf4dce) {
		return __obf_8a994c90f60ad404
	}

	if __obf_8a994c90f60ad404.__obf_04aabd427bafd586.Sign() < 0 {
		__obf_b5ef24df54cf4dce.__obf_04aabd427bafd586.Sub(__obf_b5ef24df54cf4dce.__obf_04aabd427bafd586, __obf_38f9fa21a076c52c)
	}

	return __obf_b5ef24df54cf4dce
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_8a994c90f60ad404 Decimal) RoundUp(__obf_3277a2f5cc1a9377 int32) Decimal {
	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 >= -__obf_3277a2f5cc1a9377 {
		return __obf_8a994c90f60ad404
	}

	__obf_b5ef24df54cf4dce := __obf_8a994c90f60ad404.__obf_f614630d6ca91251(-__obf_3277a2f5cc1a9377)
	if __obf_8a994c90f60ad404.Equal(__obf_b5ef24df54cf4dce) {
		return __obf_8a994c90f60ad404
	}

	if __obf_8a994c90f60ad404.__obf_04aabd427bafd586.Sign() > 0 {
		__obf_b5ef24df54cf4dce.__obf_04aabd427bafd586.Add(__obf_b5ef24df54cf4dce.__obf_04aabd427bafd586, __obf_38f9fa21a076c52c)
	} else if __obf_8a994c90f60ad404.__obf_04aabd427bafd586.Sign() < 0 {
		__obf_b5ef24df54cf4dce.__obf_04aabd427bafd586.Sub(__obf_b5ef24df54cf4dce.__obf_04aabd427bafd586, __obf_38f9fa21a076c52c)
	}

	return __obf_b5ef24df54cf4dce
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_8a994c90f60ad404 Decimal) RoundDown(__obf_3277a2f5cc1a9377 int32) Decimal {
	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 >= -__obf_3277a2f5cc1a9377 {
		return __obf_8a994c90f60ad404
	}

	__obf_b5ef24df54cf4dce := __obf_8a994c90f60ad404.__obf_f614630d6ca91251(-__obf_3277a2f5cc1a9377)
	if __obf_8a994c90f60ad404.Equal(__obf_b5ef24df54cf4dce) {
		return __obf_8a994c90f60ad404
	}
	return __obf_b5ef24df54cf4dce
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
func (__obf_8a994c90f60ad404 Decimal) RoundBank(__obf_3277a2f5cc1a9377 int32) Decimal {

	__obf_1287f1689fea2166 := __obf_8a994c90f60ad404.Round(__obf_3277a2f5cc1a9377)
	__obf_3d0c61b892811871 := __obf_8a994c90f60ad404.Sub(__obf_1287f1689fea2166).Abs()

	__obf_f7504e9163879525 := New(5, -__obf_3277a2f5cc1a9377-1)
	if __obf_3d0c61b892811871.Cmp(__obf_f7504e9163879525) == 0 && __obf_1287f1689fea2166.__obf_04aabd427bafd586.Bit(0) != 0 {
		if __obf_1287f1689fea2166.__obf_04aabd427bafd586.Sign() < 0 {
			__obf_1287f1689fea2166.__obf_04aabd427bafd586.Add(__obf_1287f1689fea2166.__obf_04aabd427bafd586, __obf_38f9fa21a076c52c)
		} else {
			__obf_1287f1689fea2166.__obf_04aabd427bafd586.Sub(__obf_1287f1689fea2166.__obf_04aabd427bafd586, __obf_38f9fa21a076c52c)
		}
	}

	return __obf_1287f1689fea2166
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
func (__obf_8a994c90f60ad404 Decimal) RoundCash(__obf_c208b02a41f708c1 uint8) Decimal {
	var __obf_dfd39af689d4bcf3 *big.Int
	switch __obf_c208b02a41f708c1 {
	case 5:
		__obf_dfd39af689d4bcf3 = __obf_1fe59b8bee7770a1
	case 10:
		__obf_dfd39af689d4bcf3 = __obf_4a348d545ccd57bf
	case 25:
		__obf_dfd39af689d4bcf3 = __obf_7e76cab4ab253fc7
	case 50:
		__obf_dfd39af689d4bcf3 = __obf_e9d2424d8d05f51b
	case 100:
		__obf_dfd39af689d4bcf3 = __obf_38f9fa21a076c52c
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_c208b02a41f708c1))
	}
	__obf_4570877839be3e45 := Decimal{
		__obf_04aabd427bafd586: __obf_dfd39af689d4bcf3,
	}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_8a994c90f60ad404.Mul(__obf_4570877839be3e45).Round(0).Div(__obf_4570877839be3e45).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_8a994c90f60ad404 Decimal) Floor() Decimal {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()

	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 >= 0 {
		return __obf_8a994c90f60ad404
	}

	__obf_3e3f0912cbecfbf3 := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_3e3f0912cbecfbf3.Exp(__obf_3e3f0912cbecfbf3, big.NewInt(-int64(__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3)), nil)

	__obf_6ac07a49686a1017 := new(big.Int).Div(__obf_8a994c90f60ad404.__obf_04aabd427bafd586, __obf_3e3f0912cbecfbf3)
	return Decimal{__obf_04aabd427bafd586: __obf_6ac07a49686a1017, __obf_3e3f0912cbecfbf3: 0}
}

// Ceil returns the nearest integer value greater than or equal to d.
func (__obf_8a994c90f60ad404 Decimal) Ceil() Decimal {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()

	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 >= 0 {
		return __obf_8a994c90f60ad404
	}

	__obf_3e3f0912cbecfbf3 := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_3e3f0912cbecfbf3.Exp(__obf_3e3f0912cbecfbf3, big.NewInt(-int64(__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3)), nil)

	__obf_6ac07a49686a1017, __obf_1a0547a84e8f209a := new(big.Int).DivMod(__obf_8a994c90f60ad404.__obf_04aabd427bafd586, __obf_3e3f0912cbecfbf3, new(big.Int))
	if __obf_1a0547a84e8f209a.Cmp(__obf_f9353e10ec98e7cd) != 0 {
		__obf_6ac07a49686a1017.Add(__obf_6ac07a49686a1017, __obf_38f9fa21a076c52c)
	}
	return Decimal{__obf_04aabd427bafd586: __obf_6ac07a49686a1017, __obf_3e3f0912cbecfbf3: 0}
}

// Truncate truncates off digits from the number, without rounding.
//
// NOTE: precision is the last digit that will not be truncated (must be >= 0).
//
// Example:
//
//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
func (__obf_8a994c90f60ad404 Decimal) Truncate(__obf_c51dca51652f0a68 int32) Decimal {
	__obf_8a994c90f60ad404.__obf_90aaf6af73b21543()
	if __obf_c51dca51652f0a68 >= 0 && -__obf_c51dca51652f0a68 > __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 {
		return __obf_8a994c90f60ad404.__obf_f614630d6ca91251(-__obf_c51dca51652f0a68)
	}
	return __obf_8a994c90f60ad404
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_8a994c90f60ad404 *Decimal) UnmarshalJSON(__obf_3d300da1dc338eae []byte) error {
	if string(__obf_3d300da1dc338eae) == "null" {
		return nil
	}

	__obf_9c64256143a865c3, __obf_8d1e8d31f090b7f5 := __obf_997794d0d833bd78(__obf_3d300da1dc338eae)
	if __obf_8d1e8d31f090b7f5 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_3d300da1dc338eae, __obf_8d1e8d31f090b7f5)
	}

	__obf_17d0cfecf7e687b6, __obf_8d1e8d31f090b7f5 := NewFromString(__obf_9c64256143a865c3)
	*__obf_8a994c90f60ad404 = __obf_17d0cfecf7e687b6
	if __obf_8d1e8d31f090b7f5 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_9c64256143a865c3, __obf_8d1e8d31f090b7f5)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_8a994c90f60ad404 Decimal) MarshalJSON() ([]byte, error) {
	var __obf_9c64256143a865c3 string
	if MarshalJSONWithoutQuotes {
		__obf_9c64256143a865c3 = __obf_8a994c90f60ad404.String()
	} else {
		__obf_9c64256143a865c3 = "\"" + __obf_8a994c90f60ad404.String() + "\""
	}
	return []byte(__obf_9c64256143a865c3), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_8a994c90f60ad404 *Decimal) UnmarshalBinary(__obf_0f9c180aa38c34e4 []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_0f9c180aa38c34e4) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_0f9c180aa38c34e4, len(__obf_0f9c180aa38c34e4))
	}

	// Extract the exponent
	__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 = int32(binary.BigEndian.Uint32(__obf_0f9c180aa38c34e4[:4]))

	// Extract the value
	__obf_8a994c90f60ad404.__obf_04aabd427bafd586 = new(big.Int)
	if __obf_8d1e8d31f090b7f5 := __obf_8a994c90f60ad404.__obf_04aabd427bafd586.GobDecode(__obf_0f9c180aa38c34e4[4:]); __obf_8d1e8d31f090b7f5 != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_0f9c180aa38c34e4, __obf_8d1e8d31f090b7f5)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_8a994c90f60ad404 Decimal) MarshalBinary() (__obf_0f9c180aa38c34e4 []byte, __obf_8d1e8d31f090b7f5 error) {
	// exp is written first, but encode value first to know output size
	var __obf_5ceda58a95b0d829 []byte
	if __obf_5ceda58a95b0d829, __obf_8d1e8d31f090b7f5 = __obf_8a994c90f60ad404.__obf_04aabd427bafd586.GobEncode(); __obf_8d1e8d31f090b7f5 != nil {
		return nil, __obf_8d1e8d31f090b7f5
	}

	// Write the exponent in front, since it's a fixed size
	__obf_6fc977816f32bf91 := make([]byte, 4, len(__obf_5ceda58a95b0d829)+4)
	binary.BigEndian.PutUint32(__obf_6fc977816f32bf91, uint32(__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3))

	// Return the byte array
	return append(__obf_6fc977816f32bf91, __obf_5ceda58a95b0d829...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_8a994c90f60ad404 *Decimal) Scan(__obf_04aabd427bafd586 any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_aff74d9405ab4816 := __obf_04aabd427bafd586.(type) {

	case float32:
		*__obf_8a994c90f60ad404 = NewFromFloat(float64(__obf_aff74d9405ab4816))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_8a994c90f60ad404 = NewFromFloat(__obf_aff74d9405ab4816)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_8a994c90f60ad404 = New(__obf_aff74d9405ab4816, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_8a994c90f60ad404 = NewFromUint64(__obf_aff74d9405ab4816)
		return nil

	default:
		// default is trying to interpret value stored as string
		__obf_9c64256143a865c3, __obf_8d1e8d31f090b7f5 := __obf_997794d0d833bd78(__obf_aff74d9405ab4816)
		if __obf_8d1e8d31f090b7f5 != nil {
			return __obf_8d1e8d31f090b7f5
		}
		*__obf_8a994c90f60ad404, __obf_8d1e8d31f090b7f5 = NewFromString(__obf_9c64256143a865c3)
		return __obf_8d1e8d31f090b7f5
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_8a994c90f60ad404 Decimal) Value() (driver.Value, error) {
	return __obf_8a994c90f60ad404.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_8a994c90f60ad404 *Decimal) UnmarshalText(__obf_42eb62fa580e29c5 []byte) error {
	__obf_9c64256143a865c3 := string(__obf_42eb62fa580e29c5)

	__obf_89a3a56c47f41acb, __obf_8d1e8d31f090b7f5 := NewFromString(__obf_9c64256143a865c3)
	*__obf_8a994c90f60ad404 = __obf_89a3a56c47f41acb
	if __obf_8d1e8d31f090b7f5 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_9c64256143a865c3, __obf_8d1e8d31f090b7f5)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_8a994c90f60ad404 Decimal) MarshalText() (__obf_42eb62fa580e29c5 []byte, __obf_8d1e8d31f090b7f5 error) {
	return []byte(__obf_8a994c90f60ad404.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_8a994c90f60ad404 Decimal) GobEncode() ([]byte, error) {
	return __obf_8a994c90f60ad404.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_8a994c90f60ad404 *Decimal) GobDecode(__obf_0f9c180aa38c34e4 []byte) error {
	return __obf_8a994c90f60ad404.UnmarshalBinary(__obf_0f9c180aa38c34e4)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_8a994c90f60ad404 Decimal) StringScaled(__obf_3e3f0912cbecfbf3 int32) string {
	return __obf_8a994c90f60ad404.__obf_f614630d6ca91251(__obf_3e3f0912cbecfbf3).String()
}

func (__obf_8a994c90f60ad404 Decimal) string(__obf_4aa8633f9dffd9d0 bool) string {
	if __obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3 >= 0 {
		return __obf_8a994c90f60ad404.__obf_f614630d6ca91251(0).__obf_04aabd427bafd586.String()
	}

	__obf_d1bb30f067c5e1eb := new(big.Int).Abs(__obf_8a994c90f60ad404.__obf_04aabd427bafd586)
	__obf_9c64256143a865c3 := __obf_d1bb30f067c5e1eb.String()

	var __obf_a325b93216c0f37b, __obf_517a92251c86ce6b string

	// NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
	// and you are on a 32-bit machine. Won't fix this super-edge case.
	__obf_825cab7a6027db70 := int(__obf_8a994c90f60ad404.__obf_3e3f0912cbecfbf3)
	if len(__obf_9c64256143a865c3) > -__obf_825cab7a6027db70 {
		__obf_a325b93216c0f37b = __obf_9c64256143a865c3[:len(__obf_9c64256143a865c3)+__obf_825cab7a6027db70]
		__obf_517a92251c86ce6b = __obf_9c64256143a865c3[len(__obf_9c64256143a865c3)+__obf_825cab7a6027db70:]
	} else {
		__obf_a325b93216c0f37b = "0"

		__obf_1a07e5092b637e1f := -__obf_825cab7a6027db70 - len(__obf_9c64256143a865c3)
		__obf_517a92251c86ce6b = strings.Repeat("0", __obf_1a07e5092b637e1f) + __obf_9c64256143a865c3
	}

	if __obf_4aa8633f9dffd9d0 {
		__obf_bf8f4a859c474f4d := len(__obf_517a92251c86ce6b) - 1
		for ; __obf_bf8f4a859c474f4d >= 0; __obf_bf8f4a859c474f4d-- {
			if __obf_517a92251c86ce6b[__obf_bf8f4a859c474f4d] != '0' {
				break
			}
		}
		__obf_517a92251c86ce6b = __obf_517a92251c86ce6b[:__obf_bf8f4a859c474f4d+1]
	}

	__obf_4a396a5326dea78f := __obf_a325b93216c0f37b
	if len(__obf_517a92251c86ce6b) > 0 {
		__obf_4a396a5326dea78f += "." + __obf_517a92251c86ce6b
	}

	if __obf_8a994c90f60ad404.__obf_04aabd427bafd586.Sign() < 0 {
		return "-" + __obf_4a396a5326dea78f
	}

	return __obf_4a396a5326dea78f
}

func (__obf_8a994c90f60ad404 *Decimal) __obf_90aaf6af73b21543() {
	if __obf_8a994c90f60ad404.__obf_04aabd427bafd586 == nil {
		__obf_8a994c90f60ad404.__obf_04aabd427bafd586 = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_a80b413745c693f7 Decimal, __obf_04acf37c7b11980b ...Decimal) Decimal {
	__obf_d17525b2ae3d2a00 := __obf_a80b413745c693f7
	for _, __obf_032164d8462a548c := range __obf_04acf37c7b11980b {
		if __obf_032164d8462a548c.Cmp(__obf_d17525b2ae3d2a00) < 0 {
			__obf_d17525b2ae3d2a00 = __obf_032164d8462a548c
		}
	}
	return __obf_d17525b2ae3d2a00
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_a80b413745c693f7 Decimal, __obf_04acf37c7b11980b ...Decimal) Decimal {
	__obf_d17525b2ae3d2a00 := __obf_a80b413745c693f7
	for _, __obf_032164d8462a548c := range __obf_04acf37c7b11980b {
		if __obf_032164d8462a548c.Cmp(__obf_d17525b2ae3d2a00) > 0 {
			__obf_d17525b2ae3d2a00 = __obf_032164d8462a548c
		}
	}
	return __obf_d17525b2ae3d2a00
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_a80b413745c693f7 Decimal, __obf_04acf37c7b11980b ...Decimal) Decimal {
	__obf_ea55ca34522424cb := __obf_a80b413745c693f7
	for _, __obf_032164d8462a548c := range __obf_04acf37c7b11980b {
		__obf_ea55ca34522424cb = __obf_ea55ca34522424cb.Add(__obf_032164d8462a548c)
	}

	return __obf_ea55ca34522424cb
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_a80b413745c693f7 Decimal, __obf_04acf37c7b11980b ...Decimal) Decimal {
	__obf_b9ea44f541228482 := New(int64(len(__obf_04acf37c7b11980b)+1), 0)
	__obf_691733ca8f4c39c1 := Sum(__obf_a80b413745c693f7, __obf_04acf37c7b11980b...)
	return __obf_691733ca8f4c39c1.Div(__obf_b9ea44f541228482)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_0b105432c1ffc5c3 Decimal, __obf_14eacb0a6cf534d0 Decimal) (Decimal, Decimal) {
	__obf_0b105432c1ffc5c3.__obf_90aaf6af73b21543()
	__obf_14eacb0a6cf534d0.__obf_90aaf6af73b21543()

	if __obf_0b105432c1ffc5c3.__obf_3e3f0912cbecfbf3 < __obf_14eacb0a6cf534d0.__obf_3e3f0912cbecfbf3 {
		return __obf_0b105432c1ffc5c3, __obf_14eacb0a6cf534d0.__obf_f614630d6ca91251(__obf_0b105432c1ffc5c3.__obf_3e3f0912cbecfbf3)
	} else if __obf_0b105432c1ffc5c3.__obf_3e3f0912cbecfbf3 > __obf_14eacb0a6cf534d0.__obf_3e3f0912cbecfbf3 {
		return __obf_0b105432c1ffc5c3.__obf_f614630d6ca91251(__obf_14eacb0a6cf534d0.__obf_3e3f0912cbecfbf3), __obf_14eacb0a6cf534d0
	}

	return __obf_0b105432c1ffc5c3, __obf_14eacb0a6cf534d0
}

func __obf_997794d0d833bd78(__obf_04aabd427bafd586 any) (string, error) {
	var __obf_f21be4ed7fec7558 []byte

	switch __obf_aff74d9405ab4816 := __obf_04aabd427bafd586.(type) {
	case string:
		__obf_f21be4ed7fec7558 = []byte(__obf_aff74d9405ab4816)
	case []byte:
		__obf_f21be4ed7fec7558 = __obf_aff74d9405ab4816
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_04aabd427bafd586, __obf_04aabd427bafd586)
	}

	// If the amount is quoted, strip the quotes
	if len(__obf_f21be4ed7fec7558) > 2 && __obf_f21be4ed7fec7558[0] == '"' && __obf_f21be4ed7fec7558[len(__obf_f21be4ed7fec7558)-1] == '"' {
		__obf_f21be4ed7fec7558 = __obf_f21be4ed7fec7558[1 : len(__obf_f21be4ed7fec7558)-1]
	}
	return string(__obf_f21be4ed7fec7558), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the datautil.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_8a994c90f60ad404 Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_8a994c90f60ad404,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_8a994c90f60ad404 *NullDecimal) Scan(__obf_04aabd427bafd586 any) error {
	if __obf_04aabd427bafd586 == nil {
		__obf_8a994c90f60ad404.Valid = false
		return nil
	}
	__obf_8a994c90f60ad404.Valid = true
	return __obf_8a994c90f60ad404.Decimal.Scan(__obf_04aabd427bafd586)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_8a994c90f60ad404 NullDecimal) Value() (driver.Value, error) {
	if !__obf_8a994c90f60ad404.Valid {
		return nil, nil
	}
	return __obf_8a994c90f60ad404.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_8a994c90f60ad404 *NullDecimal) UnmarshalJSON(__obf_3d300da1dc338eae []byte) error {
	if string(__obf_3d300da1dc338eae) == "null" {
		__obf_8a994c90f60ad404.Valid = false
		return nil
	}
	__obf_8a994c90f60ad404.Valid = true
	return __obf_8a994c90f60ad404.Decimal.UnmarshalJSON(__obf_3d300da1dc338eae)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_8a994c90f60ad404 NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_8a994c90f60ad404.Valid {
		return []byte("null"), nil
	}
	return __obf_8a994c90f60ad404.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_8a994c90f60ad404 *NullDecimal) UnmarshalText(__obf_42eb62fa580e29c5 []byte) error {
	__obf_9c64256143a865c3 := string(__obf_42eb62fa580e29c5)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_9c64256143a865c3 == "" {
		__obf_8a994c90f60ad404.Valid = false
		return nil
	}
	if __obf_8d1e8d31f090b7f5 := __obf_8a994c90f60ad404.Decimal.UnmarshalText(__obf_42eb62fa580e29c5); __obf_8d1e8d31f090b7f5 != nil {
		__obf_8a994c90f60ad404.Valid = false
		return __obf_8d1e8d31f090b7f5
	}
	__obf_8a994c90f60ad404.Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_8a994c90f60ad404 NullDecimal) MarshalText() (__obf_42eb62fa580e29c5 []byte, __obf_8d1e8d31f090b7f5 error) {
	if !__obf_8a994c90f60ad404.Valid {
		return []byte{}, nil
	}
	return __obf_8a994c90f60ad404.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_8a994c90f60ad404 Decimal) Atan() Decimal {
	if __obf_8a994c90f60ad404.Equal(NewFromFloat(0.0)) {
		return __obf_8a994c90f60ad404
	}
	if __obf_8a994c90f60ad404.GreaterThan(NewFromFloat(0.0)) {
		return __obf_8a994c90f60ad404.__obf_6b712f300bc1b02c()
	}
	return __obf_8a994c90f60ad404.Neg().__obf_6b712f300bc1b02c().Neg()
}

func (__obf_8a994c90f60ad404 Decimal) __obf_82e378658ccfe677() Decimal {
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
	__obf_6ac07a49686a1017 := __obf_8a994c90f60ad404.Mul(__obf_8a994c90f60ad404)
	__obf_6df28454c8f22423 := P0.Mul(__obf_6ac07a49686a1017).Add(P1).Mul(__obf_6ac07a49686a1017).Add(P2).Mul(__obf_6ac07a49686a1017).Add(P3).Mul(__obf_6ac07a49686a1017).Add(P4).Mul(__obf_6ac07a49686a1017)
	__obf_e12562b377c7844d := __obf_6ac07a49686a1017.Add(Q0).Mul(__obf_6ac07a49686a1017).Add(Q1).Mul(__obf_6ac07a49686a1017).Add(Q2).Mul(__obf_6ac07a49686a1017).Add(Q3).Mul(__obf_6ac07a49686a1017).Add(Q4)
	__obf_6ac07a49686a1017 = __obf_6df28454c8f22423.Div(__obf_e12562b377c7844d)
	__obf_6ac07a49686a1017 = __obf_8a994c90f60ad404.Mul(__obf_6ac07a49686a1017).Add(__obf_8a994c90f60ad404)
	return __obf_6ac07a49686a1017
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_8a994c90f60ad404 Decimal) __obf_6b712f300bc1b02c() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)      // tan(3*pi/8)
	__obf_25ab1d460e094b78 := NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_8a994c90f60ad404.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_8a994c90f60ad404.__obf_82e378658ccfe677()
	}
	if __obf_8a994c90f60ad404.GreaterThan(Tan3pio8) {
		return __obf_25ab1d460e094b78.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_8a994c90f60ad404).__obf_82e378658ccfe677()).Add(Morebits)
	}
	return __obf_25ab1d460e094b78.Div(NewFromFloat(4.0)).Add((__obf_8a994c90f60ad404.Sub(NewFromFloat(1.0)).Div(__obf_8a994c90f60ad404.Add(NewFromFloat(1.0)))).__obf_82e378658ccfe677()).Add(NewFromFloat(0.5).Mul(Morebits))
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
func (__obf_8a994c90f60ad404 Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_8a994c90f60ad404.Equal(NewFromFloat(0.0)) {
		return __obf_8a994c90f60ad404
	}
	// make argument positive but save the sign
	__obf_54a1419c0ed802d9 := false
	if __obf_8a994c90f60ad404.LessThan(NewFromFloat(0.0)) {
		__obf_8a994c90f60ad404 = __obf_8a994c90f60ad404.Neg()
		__obf_54a1419c0ed802d9 = true
	}

	__obf_9c694461c6544600 := __obf_8a994c90f60ad404.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_266d290f733ffa0d := NewFromFloat(float64(__obf_9c694461c6544600)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_9c694461c6544600&1 == 1 {
		__obf_9c694461c6544600++
		__obf_266d290f733ffa0d = __obf_266d290f733ffa0d.Add(NewFromFloat(1.0))
	}
	__obf_9c694461c6544600 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_9c694461c6544600 > 3 {
		__obf_54a1419c0ed802d9 = !__obf_54a1419c0ed802d9
		__obf_9c694461c6544600 -= 4
	}
	__obf_6ac07a49686a1017 := __obf_8a994c90f60ad404.Sub(__obf_266d290f733ffa0d.Mul(PI4A)).Sub(__obf_266d290f733ffa0d.Mul(PI4B)).Sub(__obf_266d290f733ffa0d.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_72008bc30dfcfac8 := __obf_6ac07a49686a1017.Mul(__obf_6ac07a49686a1017)

	if __obf_9c694461c6544600 == 1 || __obf_9c694461c6544600 == 2 {
		__obf_fa5568a61932fb47 := __obf_72008bc30dfcfac8.Mul(__obf_72008bc30dfcfac8).Mul(_cos[0].Mul(__obf_72008bc30dfcfac8).Add(_cos[1]).Mul(__obf_72008bc30dfcfac8).Add(_cos[2]).Mul(__obf_72008bc30dfcfac8).Add(_cos[3]).Mul(__obf_72008bc30dfcfac8).Add(_cos[4]).Mul(__obf_72008bc30dfcfac8).Add(_cos[5]))
		__obf_266d290f733ffa0d = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_72008bc30dfcfac8)).Add(__obf_fa5568a61932fb47)
	} else {
		__obf_266d290f733ffa0d = __obf_6ac07a49686a1017.Add(__obf_6ac07a49686a1017.Mul(__obf_72008bc30dfcfac8).Mul(_sin[0].Mul(__obf_72008bc30dfcfac8).Add(_sin[1]).Mul(__obf_72008bc30dfcfac8).Add(_sin[2]).Mul(__obf_72008bc30dfcfac8).Add(_sin[3]).Mul(__obf_72008bc30dfcfac8).Add(_sin[4]).Mul(__obf_72008bc30dfcfac8).Add(_sin[5])))
	}
	if __obf_54a1419c0ed802d9 {
		__obf_266d290f733ffa0d = __obf_266d290f733ffa0d.Neg()
	}
	return __obf_266d290f733ffa0d
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
func (__obf_8a994c90f60ad404 Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	// make argument positive
	__obf_54a1419c0ed802d9 := false
	if __obf_8a994c90f60ad404.LessThan(NewFromFloat(0.0)) {
		__obf_8a994c90f60ad404 = __obf_8a994c90f60ad404.Neg()
	}

	__obf_9c694461c6544600 := __obf_8a994c90f60ad404.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_266d290f733ffa0d := NewFromFloat(float64(__obf_9c694461c6544600)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_9c694461c6544600&1 == 1 {
		__obf_9c694461c6544600++
		__obf_266d290f733ffa0d = __obf_266d290f733ffa0d.Add(NewFromFloat(1.0))
	}
	__obf_9c694461c6544600 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_9c694461c6544600 > 3 {
		__obf_54a1419c0ed802d9 = !__obf_54a1419c0ed802d9
		__obf_9c694461c6544600 -= 4
	}
	if __obf_9c694461c6544600 > 1 {
		__obf_54a1419c0ed802d9 = !__obf_54a1419c0ed802d9
	}

	__obf_6ac07a49686a1017 := __obf_8a994c90f60ad404.Sub(__obf_266d290f733ffa0d.Mul(PI4A)).Sub(__obf_266d290f733ffa0d.Mul(PI4B)).Sub(__obf_266d290f733ffa0d.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_72008bc30dfcfac8 := __obf_6ac07a49686a1017.Mul(__obf_6ac07a49686a1017)

	if __obf_9c694461c6544600 == 1 || __obf_9c694461c6544600 == 2 {
		__obf_266d290f733ffa0d = __obf_6ac07a49686a1017.Add(__obf_6ac07a49686a1017.Mul(__obf_72008bc30dfcfac8).Mul(_sin[0].Mul(__obf_72008bc30dfcfac8).Add(_sin[1]).Mul(__obf_72008bc30dfcfac8).Add(_sin[2]).Mul(__obf_72008bc30dfcfac8).Add(_sin[3]).Mul(__obf_72008bc30dfcfac8).Add(_sin[4]).Mul(__obf_72008bc30dfcfac8).Add(_sin[5])))
	} else {
		__obf_fa5568a61932fb47 := __obf_72008bc30dfcfac8.Mul(__obf_72008bc30dfcfac8).Mul(_cos[0].Mul(__obf_72008bc30dfcfac8).Add(_cos[1]).Mul(__obf_72008bc30dfcfac8).Add(_cos[2]).Mul(__obf_72008bc30dfcfac8).Add(_cos[3]).Mul(__obf_72008bc30dfcfac8).Add(_cos[4]).Mul(__obf_72008bc30dfcfac8).Add(_cos[5]))
		__obf_266d290f733ffa0d = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_72008bc30dfcfac8)).Add(__obf_fa5568a61932fb47)
	}
	if __obf_54a1419c0ed802d9 {
		__obf_266d290f733ffa0d = __obf_266d290f733ffa0d.Neg()
	}
	return __obf_266d290f733ffa0d
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
func (__obf_8a994c90f60ad404 Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_8a994c90f60ad404.Equal(NewFromFloat(0.0)) {
		return __obf_8a994c90f60ad404
	}

	// make argument positive but save the sign
	__obf_54a1419c0ed802d9 := false
	if __obf_8a994c90f60ad404.LessThan(NewFromFloat(0.0)) {
		__obf_8a994c90f60ad404 = __obf_8a994c90f60ad404.Neg()
		__obf_54a1419c0ed802d9 = true
	}

	__obf_9c694461c6544600 := __obf_8a994c90f60ad404.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_266d290f733ffa0d := NewFromFloat(float64(__obf_9c694461c6544600)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_9c694461c6544600&1 == 1 {
		__obf_9c694461c6544600++
		__obf_266d290f733ffa0d = __obf_266d290f733ffa0d.Add(NewFromFloat(1.0))
	}

	__obf_6ac07a49686a1017 := __obf_8a994c90f60ad404.Sub(__obf_266d290f733ffa0d.Mul(PI4A)).Sub(__obf_266d290f733ffa0d.Mul(PI4B)).Sub(__obf_266d290f733ffa0d.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_72008bc30dfcfac8 := __obf_6ac07a49686a1017.Mul(__obf_6ac07a49686a1017)

	if __obf_72008bc30dfcfac8.GreaterThan(NewFromFloat(1e-14)) {
		__obf_fa5568a61932fb47 := __obf_72008bc30dfcfac8.Mul(_tanP[0].Mul(__obf_72008bc30dfcfac8).Add(_tanP[1]).Mul(__obf_72008bc30dfcfac8).Add(_tanP[2]))
		__obf_88ec5486ff0ec693 := __obf_72008bc30dfcfac8.Add(_tanQ[1]).Mul(__obf_72008bc30dfcfac8).Add(_tanQ[2]).Mul(__obf_72008bc30dfcfac8).Add(_tanQ[3]).Mul(__obf_72008bc30dfcfac8).Add(_tanQ[4])
		__obf_266d290f733ffa0d = __obf_6ac07a49686a1017.Add(__obf_6ac07a49686a1017.Mul(__obf_fa5568a61932fb47.Div(__obf_88ec5486ff0ec693)))
	} else {
		__obf_266d290f733ffa0d = __obf_6ac07a49686a1017
	}
	if __obf_9c694461c6544600&2 == 2 {
		__obf_266d290f733ffa0d = NewFromFloat(-1.0).Div(__obf_266d290f733ffa0d)
	}
	if __obf_54a1419c0ed802d9 {
		__obf_266d290f733ffa0d = __obf_266d290f733ffa0d.Neg()
	}
	return __obf_266d290f733ffa0d
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

type __obf_17d0cfecf7e687b6 struct {
	__obf_8a994c90f60ad404 [800]byte // digits, big-endian representation
	__obf_4673196fea260de2 int       // number of digits used
	__obf_26344dec98644d1b int       // decimal point
	__obf_f7e560ef2994ea5a bool      // negative flag
	__obf_b97622edbc30ab7f bool      // discarded nonzero digits beyond d[:nd]
}

func (__obf_b0487b5cb853a320 *__obf_17d0cfecf7e687b6) String() string {
	__obf_ea676f25508dcb39 := 10 + __obf_b0487b5cb853a320.__obf_4673196fea260de2
	if __obf_b0487b5cb853a320.__obf_26344dec98644d1b > 0 {
		__obf_ea676f25508dcb39 += __obf_b0487b5cb853a320.__obf_26344dec98644d1b
	}
	if __obf_b0487b5cb853a320.__obf_26344dec98644d1b < 0 {
		__obf_ea676f25508dcb39 += -__obf_b0487b5cb853a320.__obf_26344dec98644d1b
	}

	__obf_9be8ff2c056b2b33 := make([]byte, __obf_ea676f25508dcb39)
	__obf_fa5568a61932fb47 := 0
	switch {
	case __obf_b0487b5cb853a320.__obf_4673196fea260de2 == 0:
		return "0"

	case __obf_b0487b5cb853a320.__obf_26344dec98644d1b <= 0:
		// zeros fill space between decimal point and digits
		__obf_9be8ff2c056b2b33[__obf_fa5568a61932fb47] = '0'
		__obf_fa5568a61932fb47++
		__obf_9be8ff2c056b2b33[__obf_fa5568a61932fb47] = '.'
		__obf_fa5568a61932fb47++
		__obf_fa5568a61932fb47 += __obf_6a73efd54926dbe2(__obf_9be8ff2c056b2b33[__obf_fa5568a61932fb47 : __obf_fa5568a61932fb47+-__obf_b0487b5cb853a320.__obf_26344dec98644d1b])
		__obf_fa5568a61932fb47 += copy(__obf_9be8ff2c056b2b33[__obf_fa5568a61932fb47:], __obf_b0487b5cb853a320.__obf_8a994c90f60ad404[0:__obf_b0487b5cb853a320.__obf_4673196fea260de2])

	case __obf_b0487b5cb853a320.__obf_26344dec98644d1b < __obf_b0487b5cb853a320.__obf_4673196fea260de2:
		// decimal point in middle of digits
		__obf_fa5568a61932fb47 += copy(__obf_9be8ff2c056b2b33[__obf_fa5568a61932fb47:], __obf_b0487b5cb853a320.__obf_8a994c90f60ad404[0:__obf_b0487b5cb853a320.__obf_26344dec98644d1b])
		__obf_9be8ff2c056b2b33[__obf_fa5568a61932fb47] = '.'
		__obf_fa5568a61932fb47++
		__obf_fa5568a61932fb47 += copy(__obf_9be8ff2c056b2b33[__obf_fa5568a61932fb47:], __obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_b0487b5cb853a320.__obf_26344dec98644d1b:__obf_b0487b5cb853a320.__obf_4673196fea260de2])

	default:
		// zeros fill space between digits and decimal point
		__obf_fa5568a61932fb47 += copy(__obf_9be8ff2c056b2b33[__obf_fa5568a61932fb47:], __obf_b0487b5cb853a320.__obf_8a994c90f60ad404[0:__obf_b0487b5cb853a320.__obf_4673196fea260de2])
		__obf_fa5568a61932fb47 += __obf_6a73efd54926dbe2(__obf_9be8ff2c056b2b33[__obf_fa5568a61932fb47 : __obf_fa5568a61932fb47+__obf_b0487b5cb853a320.__obf_26344dec98644d1b-__obf_b0487b5cb853a320.__obf_4673196fea260de2])
	}
	return string(__obf_9be8ff2c056b2b33[0:__obf_fa5568a61932fb47])
}

func __obf_6a73efd54926dbe2(__obf_536aa54373a32abd []byte) int {
	for __obf_bf8f4a859c474f4d := range __obf_536aa54373a32abd {
		__obf_536aa54373a32abd[__obf_bf8f4a859c474f4d] = '0'
	}
	return len(__obf_536aa54373a32abd)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_ab200a3c820d9438(__obf_b0487b5cb853a320 *__obf_17d0cfecf7e687b6) {
	for __obf_b0487b5cb853a320.__obf_4673196fea260de2 > 0 && __obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_b0487b5cb853a320.__obf_4673196fea260de2-1] == '0' {
		__obf_b0487b5cb853a320.__obf_4673196fea260de2--
	}
	if __obf_b0487b5cb853a320.__obf_4673196fea260de2 == 0 {
		__obf_b0487b5cb853a320.__obf_26344dec98644d1b = 0
	}
}

// Assign v to a.
func (__obf_b0487b5cb853a320 *__obf_17d0cfecf7e687b6) Assign(__obf_aff74d9405ab4816 uint64) {
	var __obf_9be8ff2c056b2b33 [24]byte

	// Write reversed decimal in buf.
	__obf_ea676f25508dcb39 := 0
	for __obf_aff74d9405ab4816 > 0 {
		__obf_6eca85239a96792d := __obf_aff74d9405ab4816 / 10
		__obf_aff74d9405ab4816 -= 10 * __obf_6eca85239a96792d
		__obf_9be8ff2c056b2b33[__obf_ea676f25508dcb39] = byte(__obf_aff74d9405ab4816 + '0')
		__obf_ea676f25508dcb39++
		__obf_aff74d9405ab4816 = __obf_6eca85239a96792d
	}

	// Reverse again to produce forward decimal in a.d.
	__obf_b0487b5cb853a320.__obf_4673196fea260de2 = 0
	for __obf_ea676f25508dcb39--; __obf_ea676f25508dcb39 >= 0; __obf_ea676f25508dcb39-- {
		__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_b0487b5cb853a320.__obf_4673196fea260de2] = __obf_9be8ff2c056b2b33[__obf_ea676f25508dcb39]
		__obf_b0487b5cb853a320.__obf_4673196fea260de2++
	}
	__obf_b0487b5cb853a320.__obf_26344dec98644d1b = __obf_b0487b5cb853a320.__obf_4673196fea260de2
	__obf_ab200a3c820d9438(__obf_b0487b5cb853a320)
}

// Maximum shift that we can do in one pass without overflow.
// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
const __obf_02263a00df308ec8 = 32 << (^uint(0) >> 63)
const __obf_43e7800794da4616 = __obf_02263a00df308ec8 - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_52b3eb4384bcfb52(__obf_b0487b5cb853a320 *__obf_17d0cfecf7e687b6, __obf_27199f055c617c49 uint) {
	__obf_a61a90a65de9f74c := 0 // read pointer
	__obf_fa5568a61932fb47 := 0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_ea676f25508dcb39 uint
	for ; __obf_ea676f25508dcb39>>__obf_27199f055c617c49 == 0; __obf_a61a90a65de9f74c++ {
		if __obf_a61a90a65de9f74c >= __obf_b0487b5cb853a320.__obf_4673196fea260de2 {
			if __obf_ea676f25508dcb39 == 0 {
				// a == 0; shouldn't get here, but handle anyway.
				__obf_b0487b5cb853a320.__obf_4673196fea260de2 = 0
				return
			}
			for __obf_ea676f25508dcb39>>__obf_27199f055c617c49 == 0 {
				__obf_ea676f25508dcb39 = __obf_ea676f25508dcb39 * 10
				__obf_a61a90a65de9f74c++
			}
			break
		}
		__obf_82bb736e86a88e3a := uint(__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_a61a90a65de9f74c])
		__obf_ea676f25508dcb39 = __obf_ea676f25508dcb39*10 + __obf_82bb736e86a88e3a - '0'
	}
	__obf_b0487b5cb853a320.__obf_26344dec98644d1b -= __obf_a61a90a65de9f74c - 1

	var __obf_d6e5ed4cc09e821c uint = (1 << __obf_27199f055c617c49) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_a61a90a65de9f74c < __obf_b0487b5cb853a320.__obf_4673196fea260de2; __obf_a61a90a65de9f74c++ {
		__obf_82bb736e86a88e3a := uint(__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_a61a90a65de9f74c])
		__obf_7683115b39fa9900 := __obf_ea676f25508dcb39 >> __obf_27199f055c617c49
		__obf_ea676f25508dcb39 &= __obf_d6e5ed4cc09e821c
		__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_fa5568a61932fb47] = byte(__obf_7683115b39fa9900 + '0')
		__obf_fa5568a61932fb47++
		__obf_ea676f25508dcb39 = __obf_ea676f25508dcb39*10 + __obf_82bb736e86a88e3a - '0'
	}

	// Put down extra digits.
	for __obf_ea676f25508dcb39 > 0 {
		__obf_7683115b39fa9900 := __obf_ea676f25508dcb39 >> __obf_27199f055c617c49
		__obf_ea676f25508dcb39 &= __obf_d6e5ed4cc09e821c
		if __obf_fa5568a61932fb47 < len(__obf_b0487b5cb853a320.__obf_8a994c90f60ad404) {
			__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_fa5568a61932fb47] = byte(__obf_7683115b39fa9900 + '0')
			__obf_fa5568a61932fb47++
		} else if __obf_7683115b39fa9900 > 0 {
			__obf_b0487b5cb853a320.__obf_b97622edbc30ab7f = true
		}
		__obf_ea676f25508dcb39 = __obf_ea676f25508dcb39 * 10
	}

	__obf_b0487b5cb853a320.__obf_4673196fea260de2 = __obf_fa5568a61932fb47
	__obf_ab200a3c820d9438(__obf_b0487b5cb853a320)
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

type __obf_2062fe5f910c29b3 struct {
	__obf_c979817411052385 int    // number of new digits
	__obf_24a74a47fb0fefcc string // minus one digit if original < a.
}

var __obf_77b35c8cc5ff7919 = []__obf_2062fe5f910c29b3{
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
func __obf_baa9aa0685e961e3(__obf_fa061c05d0a4acb0 []byte, __obf_ddc1971c7184cc03 string) bool {
	for __obf_bf8f4a859c474f4d := 0; __obf_bf8f4a859c474f4d < len(__obf_ddc1971c7184cc03); __obf_bf8f4a859c474f4d++ {
		if __obf_bf8f4a859c474f4d >= len(__obf_fa061c05d0a4acb0) {
			return true
		}
		if __obf_fa061c05d0a4acb0[__obf_bf8f4a859c474f4d] != __obf_ddc1971c7184cc03[__obf_bf8f4a859c474f4d] {
			return __obf_fa061c05d0a4acb0[__obf_bf8f4a859c474f4d] < __obf_ddc1971c7184cc03[__obf_bf8f4a859c474f4d]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_90bd0f2e717d26b7(__obf_b0487b5cb853a320 *__obf_17d0cfecf7e687b6, __obf_27199f055c617c49 uint) {
	__obf_c979817411052385 := __obf_77b35c8cc5ff7919[__obf_27199f055c617c49].__obf_c979817411052385
	if __obf_baa9aa0685e961e3(__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[0:__obf_b0487b5cb853a320.__obf_4673196fea260de2], __obf_77b35c8cc5ff7919[__obf_27199f055c617c49].__obf_24a74a47fb0fefcc) {
		__obf_c979817411052385--
	}

	__obf_a61a90a65de9f74c := __obf_b0487b5cb853a320.__obf_4673196fea260de2                          // read index
	__obf_fa5568a61932fb47 := __obf_b0487b5cb853a320.__obf_4673196fea260de2 + __obf_c979817411052385 // write index

	// Pick up a digit, put down a digit.
	var __obf_ea676f25508dcb39 uint
	for __obf_a61a90a65de9f74c--; __obf_a61a90a65de9f74c >= 0; __obf_a61a90a65de9f74c-- {
		__obf_ea676f25508dcb39 += (uint(__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_a61a90a65de9f74c]) - '0') << __obf_27199f055c617c49
		__obf_a88f31b3480cfcc9 := __obf_ea676f25508dcb39 / 10
		__obf_86c50d43b5d02220 := __obf_ea676f25508dcb39 - 10*__obf_a88f31b3480cfcc9
		__obf_fa5568a61932fb47--
		if __obf_fa5568a61932fb47 < len(__obf_b0487b5cb853a320.__obf_8a994c90f60ad404) {
			__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_fa5568a61932fb47] = byte(__obf_86c50d43b5d02220 + '0')
		} else if __obf_86c50d43b5d02220 != 0 {
			__obf_b0487b5cb853a320.__obf_b97622edbc30ab7f = true
		}
		__obf_ea676f25508dcb39 = __obf_a88f31b3480cfcc9
	}

	// Put down extra digits.
	for __obf_ea676f25508dcb39 > 0 {
		__obf_a88f31b3480cfcc9 := __obf_ea676f25508dcb39 / 10
		__obf_86c50d43b5d02220 := __obf_ea676f25508dcb39 - 10*__obf_a88f31b3480cfcc9
		__obf_fa5568a61932fb47--
		if __obf_fa5568a61932fb47 < len(__obf_b0487b5cb853a320.__obf_8a994c90f60ad404) {
			__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_fa5568a61932fb47] = byte(__obf_86c50d43b5d02220 + '0')
		} else if __obf_86c50d43b5d02220 != 0 {
			__obf_b0487b5cb853a320.__obf_b97622edbc30ab7f = true
		}
		__obf_ea676f25508dcb39 = __obf_a88f31b3480cfcc9
	}

	__obf_b0487b5cb853a320.__obf_4673196fea260de2 += __obf_c979817411052385
	if __obf_b0487b5cb853a320.__obf_4673196fea260de2 >= len(__obf_b0487b5cb853a320.__obf_8a994c90f60ad404) {
		__obf_b0487b5cb853a320.__obf_4673196fea260de2 = len(__obf_b0487b5cb853a320.__obf_8a994c90f60ad404)
	}
	__obf_b0487b5cb853a320.__obf_26344dec98644d1b += __obf_c979817411052385
	__obf_ab200a3c820d9438(__obf_b0487b5cb853a320)
}

// Binary shift left (k > 0) or right (k < 0).
func (__obf_b0487b5cb853a320 *__obf_17d0cfecf7e687b6) Shift(__obf_27199f055c617c49 int) {
	switch {
	case __obf_b0487b5cb853a320.__obf_4673196fea260de2 == 0:
		// nothing to do: a == 0
	case __obf_27199f055c617c49 > 0:
		for __obf_27199f055c617c49 > __obf_43e7800794da4616 {
			__obf_90bd0f2e717d26b7(__obf_b0487b5cb853a320, __obf_43e7800794da4616)
			__obf_27199f055c617c49 -= __obf_43e7800794da4616
		}
		__obf_90bd0f2e717d26b7(__obf_b0487b5cb853a320, uint(__obf_27199f055c617c49))
	case __obf_27199f055c617c49 < 0:
		for __obf_27199f055c617c49 < -__obf_43e7800794da4616 {
			__obf_52b3eb4384bcfb52(__obf_b0487b5cb853a320, __obf_43e7800794da4616)
			__obf_27199f055c617c49 += __obf_43e7800794da4616
		}
		__obf_52b3eb4384bcfb52(__obf_b0487b5cb853a320, uint(-__obf_27199f055c617c49))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_db6a4a6638e82e43(__obf_b0487b5cb853a320 *__obf_17d0cfecf7e687b6, __obf_4673196fea260de2 int) bool {
	if __obf_4673196fea260de2 < 0 || __obf_4673196fea260de2 >= __obf_b0487b5cb853a320.__obf_4673196fea260de2 {
		return false
	}
	if __obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_4673196fea260de2] == '5' && __obf_4673196fea260de2+1 == __obf_b0487b5cb853a320.__obf_4673196fea260de2 { // exactly halfway - round to even
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_b0487b5cb853a320.__obf_b97622edbc30ab7f {
			return true
		}
		return __obf_4673196fea260de2 > 0 && (__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_4673196fea260de2-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_4673196fea260de2] >= '5'
}

// Round a to nd digits (or fewer).
// If nd is zero, it means we're rounding
// just to the left of the digits, as in
// 0.09 -> 0.1.
func (__obf_b0487b5cb853a320 *__obf_17d0cfecf7e687b6) Round(__obf_4673196fea260de2 int) {
	if __obf_4673196fea260de2 < 0 || __obf_4673196fea260de2 >= __obf_b0487b5cb853a320.__obf_4673196fea260de2 {
		return
	}
	if __obf_db6a4a6638e82e43(__obf_b0487b5cb853a320, __obf_4673196fea260de2) {
		__obf_b0487b5cb853a320.RoundUp(__obf_4673196fea260de2)
	} else {
		__obf_b0487b5cb853a320.RoundDown(__obf_4673196fea260de2)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_b0487b5cb853a320 *__obf_17d0cfecf7e687b6) RoundDown(__obf_4673196fea260de2 int) {
	if __obf_4673196fea260de2 < 0 || __obf_4673196fea260de2 >= __obf_b0487b5cb853a320.__obf_4673196fea260de2 {
		return
	}
	__obf_b0487b5cb853a320.__obf_4673196fea260de2 = __obf_4673196fea260de2
	__obf_ab200a3c820d9438(__obf_b0487b5cb853a320)
}

// Round a up to nd digits (or fewer).
func (__obf_b0487b5cb853a320 *__obf_17d0cfecf7e687b6) RoundUp(__obf_4673196fea260de2 int) {
	if __obf_4673196fea260de2 < 0 || __obf_4673196fea260de2 >= __obf_b0487b5cb853a320.__obf_4673196fea260de2 {
		return
	}

	// round up
	for __obf_bf8f4a859c474f4d := __obf_4673196fea260de2 - 1; __obf_bf8f4a859c474f4d >= 0; __obf_bf8f4a859c474f4d-- {
		__obf_82bb736e86a88e3a := __obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_bf8f4a859c474f4d]
		if __obf_82bb736e86a88e3a < '9' { // can stop after this digit
			__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_bf8f4a859c474f4d]++
			__obf_b0487b5cb853a320.__obf_4673196fea260de2 = __obf_bf8f4a859c474f4d + 1
			return
		}
	}

	// Number is all 9s.
	// Change to single 1 with adjusted decimal point.
	__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[0] = '1'
	__obf_b0487b5cb853a320.__obf_4673196fea260de2 = 1
	__obf_b0487b5cb853a320.__obf_26344dec98644d1b++
}

// Extract integer part, rounded appropriately.
// No guarantees about overflow.
func (__obf_b0487b5cb853a320 *__obf_17d0cfecf7e687b6) RoundedInteger() uint64 {
	if __obf_b0487b5cb853a320.__obf_26344dec98644d1b > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_bf8f4a859c474f4d int
	__obf_ea676f25508dcb39 := uint64(0)
	for __obf_bf8f4a859c474f4d = 0; __obf_bf8f4a859c474f4d < __obf_b0487b5cb853a320.__obf_26344dec98644d1b && __obf_bf8f4a859c474f4d < __obf_b0487b5cb853a320.__obf_4673196fea260de2; __obf_bf8f4a859c474f4d++ {
		__obf_ea676f25508dcb39 = __obf_ea676f25508dcb39*10 + uint64(__obf_b0487b5cb853a320.__obf_8a994c90f60ad404[__obf_bf8f4a859c474f4d]-'0')
	}
	for ; __obf_bf8f4a859c474f4d < __obf_b0487b5cb853a320.__obf_26344dec98644d1b; __obf_bf8f4a859c474f4d++ {
		__obf_ea676f25508dcb39 *= 10
	}
	if __obf_db6a4a6638e82e43(__obf_b0487b5cb853a320, __obf_b0487b5cb853a320.__obf_26344dec98644d1b) {
		__obf_ea676f25508dcb39++
	}
	return __obf_ea676f25508dcb39
}
