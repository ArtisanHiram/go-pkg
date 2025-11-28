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
package __obf_0962dc77c6b6239b

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

var __obf_1bb56bff7931ff30 = big.NewInt(0)
var __obf_25f16531cde494de = big.NewInt(1)
var __obf_3cbd576e784ed478 = big.NewInt(2)
var __obf_e1dd4002d73ad6cd = big.NewInt(4)
var __obf_dc2da84e9cc8100a = big.NewInt(5)
var __obf_44ff59a8ed385a0e = big.NewInt(10)
var __obf_3316a3319b2a8b33 = big.NewInt(20)

var __obf_b83126d172f97ebd = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_c36b36e18c228e7f *big.Int

	// NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.
	__obf_406325483a83b1fa int32
}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_c36b36e18c228e7f int64, __obf_406325483a83b1fa int32) Decimal {
	return Decimal{
		__obf_c36b36e18c228e7f: big.NewInt(__obf_c36b36e18c228e7f),
		__obf_406325483a83b1fa: __obf_406325483a83b1fa,
	}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_c36b36e18c228e7f int64) Decimal {
	return Decimal{
		__obf_c36b36e18c228e7f: big.NewInt(__obf_c36b36e18c228e7f),
		__obf_406325483a83b1fa: 0,
	}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_c36b36e18c228e7f int32) Decimal {
	return Decimal{
		__obf_c36b36e18c228e7f: big.NewInt(int64(__obf_c36b36e18c228e7f)),
		__obf_406325483a83b1fa: 0,
	}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_c36b36e18c228e7f uint64) Decimal {
	return Decimal{
		__obf_c36b36e18c228e7f: new(big.Int).SetUint64(__obf_c36b36e18c228e7f),
		__obf_406325483a83b1fa: 0,
	}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_c36b36e18c228e7f *big.Int, __obf_406325483a83b1fa int32) Decimal {
	return Decimal{
		__obf_c36b36e18c228e7f: new(big.Int).Set(__obf_c36b36e18c228e7f),
		__obf_406325483a83b1fa: __obf_406325483a83b1fa,
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
func NewFromBigRat(__obf_c36b36e18c228e7f *big.Rat, __obf_7efed86618b2f272 int32) Decimal {
	return Decimal{
		__obf_c36b36e18c228e7f: new(big.Int).Set(__obf_c36b36e18c228e7f.Num()),
		__obf_406325483a83b1fa: 0,
	}.DivRound(Decimal{
		__obf_c36b36e18c228e7f: new(big.Int).Set(__obf_c36b36e18c228e7f.Denom()),
		__obf_406325483a83b1fa: 0,
	}, __obf_7efed86618b2f272)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_c36b36e18c228e7f string) (Decimal, error) {
	__obf_f4fe036c72616c63 := __obf_c36b36e18c228e7f
	var __obf_fae6fc384188eff2 string
	var __obf_406325483a83b1fa int64

	// Check if number is using scientific notation
	__obf_fb702628f4231304 := strings.IndexAny(__obf_c36b36e18c228e7f, "Ee")
	if __obf_fb702628f4231304 != -1 {
		__obf_59a5d6e7dc9ae100, __obf_c1d612f6bfc234d7 := strconv.ParseInt(__obf_c36b36e18c228e7f[__obf_fb702628f4231304+1:], 10, 32)
		if __obf_c1d612f6bfc234d7 != nil {
			if __obf_6b31de26b3f4abad, __obf_5a1cc71347798e7b := __obf_c1d612f6bfc234d7.(*strconv.NumError); __obf_5a1cc71347798e7b && __obf_6b31de26b3f4abad.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_c36b36e18c228e7f)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_c36b36e18c228e7f)
		}
		__obf_c36b36e18c228e7f = __obf_c36b36e18c228e7f[:__obf_fb702628f4231304]
		__obf_406325483a83b1fa = __obf_59a5d6e7dc9ae100
	}

	__obf_8eb05059edeaaa5c := -1
	__obf_53c198495e2c6dcf := len(__obf_c36b36e18c228e7f)
	for __obf_a6733ea50196cc53 := 0; __obf_a6733ea50196cc53 < __obf_53c198495e2c6dcf; __obf_a6733ea50196cc53++ {
		if __obf_c36b36e18c228e7f[__obf_a6733ea50196cc53] == '.' {
			if __obf_8eb05059edeaaa5c > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_c36b36e18c228e7f)
			}
			__obf_8eb05059edeaaa5c = __obf_a6733ea50196cc53
		}
	}

	if __obf_8eb05059edeaaa5c == -1 {
		// There is no decimal point, we can just parse the original string as
		// an int
		__obf_fae6fc384188eff2 = __obf_c36b36e18c228e7f
	} else {
		if __obf_8eb05059edeaaa5c+1 < __obf_53c198495e2c6dcf {
			__obf_fae6fc384188eff2 = __obf_c36b36e18c228e7f[:__obf_8eb05059edeaaa5c] + __obf_c36b36e18c228e7f[__obf_8eb05059edeaaa5c+1:]
		} else {
			__obf_fae6fc384188eff2 = __obf_c36b36e18c228e7f[:__obf_8eb05059edeaaa5c]
		}
		__obf_59a5d6e7dc9ae100 := -len(__obf_c36b36e18c228e7f[__obf_8eb05059edeaaa5c+1:])
		__obf_406325483a83b1fa += int64(__obf_59a5d6e7dc9ae100)
	}

	var __obf_68e723775b28a699 *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_fae6fc384188eff2) <= 18 {
		__obf_57eef8a374887423, __obf_c1d612f6bfc234d7 := strconv.ParseInt(__obf_fae6fc384188eff2, 10, 64)
		if __obf_c1d612f6bfc234d7 != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_c36b36e18c228e7f)
		}
		__obf_68e723775b28a699 = big.NewInt(__obf_57eef8a374887423)
	} else {
		__obf_68e723775b28a699 = new(big.Int)
		_, __obf_5a1cc71347798e7b := __obf_68e723775b28a699.SetString(__obf_fae6fc384188eff2, 10)
		if !__obf_5a1cc71347798e7b {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_c36b36e18c228e7f)
		}
	}

	if __obf_406325483a83b1fa < math.MinInt32 || __obf_406325483a83b1fa > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_f4fe036c72616c63)
	}

	return Decimal{
		__obf_c36b36e18c228e7f: __obf_68e723775b28a699,
		__obf_406325483a83b1fa: int32(__obf_406325483a83b1fa),
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
func NewFromFormattedString(__obf_c36b36e18c228e7f string, __obf_0a98142fb9fdee24 *regexp.Regexp) (Decimal, error) {
	__obf_a22955a460c354ca := __obf_0a98142fb9fdee24.ReplaceAllString(__obf_c36b36e18c228e7f, "")
	__obf_d09f058c0a390c93, __obf_c1d612f6bfc234d7 := NewFromString(__obf_a22955a460c354ca)
	if __obf_c1d612f6bfc234d7 != nil {
		return Decimal{}, __obf_c1d612f6bfc234d7
	}
	return __obf_d09f058c0a390c93, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_c36b36e18c228e7f string) Decimal {
	__obf_bc636ce96a8a7277, __obf_c1d612f6bfc234d7 := NewFromString(__obf_c36b36e18c228e7f)
	if __obf_c1d612f6bfc234d7 != nil {
		panic(__obf_c1d612f6bfc234d7)
	}
	return __obf_bc636ce96a8a7277
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
func NewFromFloat(__obf_c36b36e18c228e7f float64) Decimal {
	if __obf_c36b36e18c228e7f == 0 {
		return New(0, 0)
	}
	return __obf_428cf55ae033ae83(__obf_c36b36e18c228e7f, math.Float64bits(__obf_c36b36e18c228e7f), &__obf_b3a4121c853c30da)
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
func NewFromFloat32(__obf_c36b36e18c228e7f float32) Decimal {
	if __obf_c36b36e18c228e7f == 0 {
		return New(0, 0)
	}
	// XOR is workaround for https://github.com/golang/go/issues/26285
	__obf_5097a5efcc046ddc := math.Float32bits(__obf_c36b36e18c228e7f) ^ 0x80808080
	return __obf_428cf55ae033ae83(float64(__obf_c36b36e18c228e7f), uint64(__obf_5097a5efcc046ddc)^0x80808080, &__obf_27426ff0b890a183)
}

func __obf_428cf55ae033ae83(__obf_4c6accf7c36a21ec float64, __obf_0e017b837b95afdb uint64, __obf_8b0328b22b58eb64 *__obf_e46206765e5dbb45) Decimal {
	if math.IsNaN(__obf_4c6accf7c36a21ec) || math.IsInf(__obf_4c6accf7c36a21ec, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_4c6accf7c36a21ec))
	}
	__obf_406325483a83b1fa := int(__obf_0e017b837b95afdb>>__obf_8b0328b22b58eb64.__obf_f89d4a69e09b5dc8) & (1<<__obf_8b0328b22b58eb64.__obf_44fa0a4671bd9290 - 1)
	__obf_a54bc9ed25a8de66 := __obf_0e017b837b95afdb & (uint64(1)<<__obf_8b0328b22b58eb64.__obf_f89d4a69e09b5dc8 - 1)

	switch __obf_406325483a83b1fa {
	case 0:
		// denormalized
		__obf_406325483a83b1fa++

	default:
		// add implicit top bit
		__obf_a54bc9ed25a8de66 |= uint64(1) << __obf_8b0328b22b58eb64.__obf_f89d4a69e09b5dc8
	}
	__obf_406325483a83b1fa += __obf_8b0328b22b58eb64.__obf_279e6fa356ffbd95

	var __obf_d09f058c0a390c93 __obf_0962dc77c6b6239b
	__obf_d09f058c0a390c93.Assign(__obf_a54bc9ed25a8de66)
	__obf_d09f058c0a390c93.Shift(__obf_406325483a83b1fa - int(__obf_8b0328b22b58eb64.__obf_f89d4a69e09b5dc8))
	__obf_d09f058c0a390c93.__obf_f0af453d1e176af6 = __obf_0e017b837b95afdb>>(__obf_8b0328b22b58eb64.__obf_44fa0a4671bd9290+__obf_8b0328b22b58eb64.__obf_f89d4a69e09b5dc8) != 0

	__obf_f513f8a0934968e3(&__obf_d09f058c0a390c93, __obf_a54bc9ed25a8de66, __obf_406325483a83b1fa, __obf_8b0328b22b58eb64)
	// If less than 19 digits, we can do calculation in an int64.
	if __obf_d09f058c0a390c93.__obf_ea04243a9869d25e < 19 {
		__obf_bab5f70189aa167a := int64(0)
		__obf_b041e2c85e138a64 := int64(1)
		for __obf_a6733ea50196cc53 := __obf_d09f058c0a390c93.__obf_ea04243a9869d25e - 1; __obf_a6733ea50196cc53 >= 0; __obf_a6733ea50196cc53-- {
			__obf_bab5f70189aa167a += __obf_b041e2c85e138a64 * int64(__obf_d09f058c0a390c93.__obf_d09f058c0a390c93[__obf_a6733ea50196cc53]-'0')
			__obf_b041e2c85e138a64 *= 10
		}
		if __obf_d09f058c0a390c93.__obf_f0af453d1e176af6 {
			__obf_bab5f70189aa167a *= -1
		}
		return Decimal{__obf_c36b36e18c228e7f: big.NewInt(__obf_bab5f70189aa167a), __obf_406325483a83b1fa: int32(__obf_d09f058c0a390c93.__obf_cb19c1fd6e41a0d7) - int32(__obf_d09f058c0a390c93.__obf_ea04243a9869d25e)}
	}
	__obf_68e723775b28a699 := new(big.Int)
	__obf_68e723775b28a699, __obf_5a1cc71347798e7b := __obf_68e723775b28a699.SetString(string(__obf_d09f058c0a390c93.__obf_d09f058c0a390c93[:__obf_d09f058c0a390c93.__obf_ea04243a9869d25e]), 10)
	if __obf_5a1cc71347798e7b {
		return Decimal{__obf_c36b36e18c228e7f: __obf_68e723775b28a699, __obf_406325483a83b1fa: int32(__obf_d09f058c0a390c93.__obf_cb19c1fd6e41a0d7) - int32(__obf_d09f058c0a390c93.__obf_ea04243a9869d25e)}
	}

	return NewFromFloatWithExponent(__obf_4c6accf7c36a21ec, int32(__obf_d09f058c0a390c93.__obf_cb19c1fd6e41a0d7)-int32(__obf_d09f058c0a390c93.__obf_ea04243a9869d25e))
}

// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
// number of fractional digits.
//
// Example:
//
//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
func NewFromFloatWithExponent(__obf_c36b36e18c228e7f float64, __obf_406325483a83b1fa int32) Decimal {
	if math.IsNaN(__obf_c36b36e18c228e7f) || math.IsInf(__obf_c36b36e18c228e7f, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_c36b36e18c228e7f))
	}

	__obf_0e017b837b95afdb := math.Float64bits(__obf_c36b36e18c228e7f)
	__obf_a54bc9ed25a8de66 := __obf_0e017b837b95afdb & (1<<52 - 1)
	__obf_739f278b145b2380 := int32((__obf_0e017b837b95afdb >> 52) & (1<<11 - 1))
	__obf_87165cc42c6992a3 := __obf_0e017b837b95afdb >> 63

	if __obf_739f278b145b2380 == 0 {
		// specials
		if __obf_a54bc9ed25a8de66 == 0 {
			return Decimal{}
		}
		// subnormal
		__obf_739f278b145b2380++
	} else {
		// normal
		__obf_a54bc9ed25a8de66 |= 1 << 52
	}

	__obf_739f278b145b2380 -= 1023 + 52

	// normalizing base-2 values
	for __obf_a54bc9ed25a8de66&1 == 0 {
		__obf_a54bc9ed25a8de66 = __obf_a54bc9ed25a8de66 >> 1
		__obf_739f278b145b2380++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_406325483a83b1fa < 0 && __obf_406325483a83b1fa < __obf_739f278b145b2380 {
		if __obf_739f278b145b2380 < 0 {
			__obf_406325483a83b1fa = __obf_739f278b145b2380
		} else {
			__obf_406325483a83b1fa = 0
		}
	}

	// representing 10^M * 2^N as 5^M * 2^(M+N)
	__obf_739f278b145b2380 -= __obf_406325483a83b1fa

	__obf_02a8739fa97e954e := big.NewInt(1)
	__obf_9178141ce7345815 := big.NewInt(int64(__obf_a54bc9ed25a8de66))

	// applying 5^M
	if __obf_406325483a83b1fa > 0 {
		__obf_02a8739fa97e954e = __obf_02a8739fa97e954e.SetInt64(int64(__obf_406325483a83b1fa))
		__obf_02a8739fa97e954e = __obf_02a8739fa97e954e.Exp(__obf_dc2da84e9cc8100a, __obf_02a8739fa97e954e, nil)
	} else if __obf_406325483a83b1fa < 0 {
		__obf_02a8739fa97e954e = __obf_02a8739fa97e954e.SetInt64(-int64(__obf_406325483a83b1fa))
		__obf_02a8739fa97e954e = __obf_02a8739fa97e954e.Exp(__obf_dc2da84e9cc8100a, __obf_02a8739fa97e954e, nil)
		__obf_9178141ce7345815 = __obf_9178141ce7345815.Mul(__obf_9178141ce7345815, __obf_02a8739fa97e954e)
		__obf_02a8739fa97e954e = __obf_02a8739fa97e954e.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_739f278b145b2380 > 0 {
		__obf_9178141ce7345815 = __obf_9178141ce7345815.Lsh(__obf_9178141ce7345815, uint(__obf_739f278b145b2380))
	} else if __obf_739f278b145b2380 < 0 {
		__obf_02a8739fa97e954e = __obf_02a8739fa97e954e.Lsh(__obf_02a8739fa97e954e, uint(-__obf_739f278b145b2380))
	}

	// rounding and downscaling
	if __obf_406325483a83b1fa > 0 || __obf_739f278b145b2380 < 0 {
		__obf_ce7d20e6f5981c53 := new(big.Int).Rsh(__obf_02a8739fa97e954e, 1)
		__obf_9178141ce7345815 = __obf_9178141ce7345815.Add(__obf_9178141ce7345815, __obf_ce7d20e6f5981c53)
		__obf_9178141ce7345815 = __obf_9178141ce7345815.Quo(__obf_9178141ce7345815, __obf_02a8739fa97e954e)
	}

	if __obf_87165cc42c6992a3 == 1 {
		__obf_9178141ce7345815 = __obf_9178141ce7345815.Neg(__obf_9178141ce7345815)
	}

	return Decimal{
		__obf_c36b36e18c228e7f: __obf_9178141ce7345815,
		__obf_406325483a83b1fa: __obf_406325483a83b1fa,
	}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_d09f058c0a390c93 Decimal) Copy() Decimal {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()
	return Decimal{
		__obf_c36b36e18c228e7f: new(big.Int).Set(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f),
		__obf_406325483a83b1fa: __obf_d09f058c0a390c93.__obf_406325483a83b1fa,
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
func (__obf_d09f058c0a390c93 Decimal) __obf_d44df924a1b98873(__obf_406325483a83b1fa int32) Decimal {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()

	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa == __obf_406325483a83b1fa {
		return Decimal{
			new(big.Int).Set(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f),
			__obf_d09f058c0a390c93.__obf_406325483a83b1fa,
		}
	}

	// NOTE(vadim): must convert exps to float64 before - to prevent overflow
	__obf_b023c1714ba1c58d := math.Abs(float64(__obf_406325483a83b1fa) - float64(__obf_d09f058c0a390c93.__obf_406325483a83b1fa))
	__obf_c36b36e18c228e7f := new(big.Int).Set(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f)

	__obf_a981a455fcecbe85 := new(big.Int).Exp(__obf_44ff59a8ed385a0e, big.NewInt(int64(__obf_b023c1714ba1c58d)), nil)
	if __obf_406325483a83b1fa > __obf_d09f058c0a390c93.__obf_406325483a83b1fa {
		__obf_c36b36e18c228e7f = __obf_c36b36e18c228e7f.Quo(__obf_c36b36e18c228e7f, __obf_a981a455fcecbe85)
	} else if __obf_406325483a83b1fa < __obf_d09f058c0a390c93.__obf_406325483a83b1fa {
		__obf_c36b36e18c228e7f = __obf_c36b36e18c228e7f.Mul(__obf_c36b36e18c228e7f, __obf_a981a455fcecbe85)
	}

	return Decimal{
		__obf_c36b36e18c228e7f: __obf_c36b36e18c228e7f,
		__obf_406325483a83b1fa: __obf_406325483a83b1fa,
	}
}

// Abs returns the absolute value of the decimal.
func (__obf_d09f058c0a390c93 Decimal) Abs() Decimal {
	if !__obf_d09f058c0a390c93.IsNegative() {
		return __obf_d09f058c0a390c93
	}
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()
	__obf_7b9a006e7f3823b4 := new(big.Int).Abs(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f)
	return Decimal{
		__obf_c36b36e18c228e7f: __obf_7b9a006e7f3823b4,
		__obf_406325483a83b1fa: __obf_d09f058c0a390c93.__obf_406325483a83b1fa,
	}
}

// Add returns d + d2.
func (__obf_d09f058c0a390c93 Decimal) Add(__obf_f3234e094c8ed488 Decimal) Decimal {
	__obf_c1fa73dffb7b10fa, __obf_3e0ad0062bf745cb := RescalePair(__obf_d09f058c0a390c93, __obf_f3234e094c8ed488)

	__obf_8df00b326bd0e7c3 := new(big.Int).Add(__obf_c1fa73dffb7b10fa.__obf_c36b36e18c228e7f, __obf_3e0ad0062bf745cb.__obf_c36b36e18c228e7f)
	return Decimal{
		__obf_c36b36e18c228e7f: __obf_8df00b326bd0e7c3,
		__obf_406325483a83b1fa: __obf_c1fa73dffb7b10fa.__obf_406325483a83b1fa,
	}
}

// Sub returns d - d2.
func (__obf_d09f058c0a390c93 Decimal) Sub(__obf_f3234e094c8ed488 Decimal) Decimal {
	__obf_c1fa73dffb7b10fa, __obf_3e0ad0062bf745cb := RescalePair(__obf_d09f058c0a390c93, __obf_f3234e094c8ed488)

	__obf_8df00b326bd0e7c3 := new(big.Int).Sub(__obf_c1fa73dffb7b10fa.__obf_c36b36e18c228e7f, __obf_3e0ad0062bf745cb.__obf_c36b36e18c228e7f)
	return Decimal{
		__obf_c36b36e18c228e7f: __obf_8df00b326bd0e7c3,
		__obf_406325483a83b1fa: __obf_c1fa73dffb7b10fa.__obf_406325483a83b1fa,
	}
}

// Neg returns -d.
func (__obf_d09f058c0a390c93 Decimal) Neg() Decimal {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()
	__obf_4c6accf7c36a21ec := new(big.Int).Neg(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f)
	return Decimal{
		__obf_c36b36e18c228e7f: __obf_4c6accf7c36a21ec,
		__obf_406325483a83b1fa: __obf_d09f058c0a390c93.__obf_406325483a83b1fa,
	}
}

// Mul returns d * d2.
func (__obf_d09f058c0a390c93 Decimal) Mul(__obf_f3234e094c8ed488 Decimal) Decimal {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()
	__obf_f3234e094c8ed488.__obf_b119326c257b1b2a()

	__obf_5587418c09b335be := int64(__obf_d09f058c0a390c93.__obf_406325483a83b1fa) + int64(__obf_f3234e094c8ed488.__obf_406325483a83b1fa)
	if __obf_5587418c09b335be > math.MaxInt32 || __obf_5587418c09b335be < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_5587418c09b335be))
	}

	__obf_8df00b326bd0e7c3 := new(big.Int).Mul(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f, __obf_f3234e094c8ed488.__obf_c36b36e18c228e7f)
	return Decimal{
		__obf_c36b36e18c228e7f: __obf_8df00b326bd0e7c3,
		__obf_406325483a83b1fa: int32(__obf_5587418c09b335be),
	}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_d09f058c0a390c93 Decimal) Shift(__obf_f054861756521c97 int32) Decimal {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()
	return Decimal{
		__obf_c36b36e18c228e7f: new(big.Int).Set(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f),
		__obf_406325483a83b1fa: __obf_d09f058c0a390c93.__obf_406325483a83b1fa + __obf_f054861756521c97,
	}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_d09f058c0a390c93 Decimal) Div(__obf_f3234e094c8ed488 Decimal) Decimal {
	return __obf_d09f058c0a390c93.DivRound(__obf_f3234e094c8ed488, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_d09f058c0a390c93 Decimal) QuoRem(__obf_f3234e094c8ed488 Decimal, __obf_7efed86618b2f272 int32) (Decimal, Decimal) {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()
	__obf_f3234e094c8ed488.__obf_b119326c257b1b2a()
	if __obf_f3234e094c8ed488.__obf_c36b36e18c228e7f.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_7916ec76510088c1 := -__obf_7efed86618b2f272
	__obf_6b31de26b3f4abad := int64(__obf_d09f058c0a390c93.__obf_406325483a83b1fa) - int64(__obf_f3234e094c8ed488.__obf_406325483a83b1fa) - int64(__obf_7916ec76510088c1)
	if __obf_6b31de26b3f4abad > math.MaxInt32 || __obf_6b31de26b3f4abad < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_15218c7dca47c57c, __obf_6fa38131d19e1364, __obf_9fb732b11a24b3c8 big.Int
	var __obf_c208697a844b8def int32
	// d = a 10^ea
	// d2 = b 10^eb
	if __obf_6b31de26b3f4abad < 0 {
		__obf_15218c7dca47c57c = *__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f
		__obf_9fb732b11a24b3c8.SetInt64(-__obf_6b31de26b3f4abad)
		__obf_6fa38131d19e1364.Exp(__obf_44ff59a8ed385a0e, &__obf_9fb732b11a24b3c8, nil)
		__obf_6fa38131d19e1364.Mul(__obf_f3234e094c8ed488.__obf_c36b36e18c228e7f, &__obf_6fa38131d19e1364)
		__obf_c208697a844b8def = __obf_d09f058c0a390c93.__obf_406325483a83b1fa
		// now aa = a
		//     bb = b 10^(scale + eb - ea)
	} else {
		__obf_9fb732b11a24b3c8.SetInt64(__obf_6b31de26b3f4abad)
		__obf_15218c7dca47c57c.Exp(__obf_44ff59a8ed385a0e, &__obf_9fb732b11a24b3c8, nil)
		__obf_15218c7dca47c57c.Mul(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f, &__obf_15218c7dca47c57c)
		__obf_6fa38131d19e1364 = *__obf_f3234e094c8ed488.__obf_c36b36e18c228e7f
		__obf_c208697a844b8def = __obf_7916ec76510088c1 + __obf_f3234e094c8ed488.__obf_406325483a83b1fa
		// now aa = a ^ (ea - eb - scale)
		//     bb = b
	}
	var __obf_1b9e13ef114085d1, __obf_a934ca815da1fc1a big.Int
	__obf_1b9e13ef114085d1.QuoRem(&__obf_15218c7dca47c57c, &__obf_6fa38131d19e1364, &__obf_a934ca815da1fc1a)
	__obf_0fb9f51f7f599bf6 := Decimal{__obf_c36b36e18c228e7f: &__obf_1b9e13ef114085d1, __obf_406325483a83b1fa: __obf_7916ec76510088c1}
	__obf_851402996a7905d1 := Decimal{__obf_c36b36e18c228e7f: &__obf_a934ca815da1fc1a, __obf_406325483a83b1fa: __obf_c208697a844b8def}
	return __obf_0fb9f51f7f599bf6, __obf_851402996a7905d1
}

// DivRound divides and rounds to a given precision
// i.e. to an integer multiple of 10^(-precision)
//
//	for a positive quotient digit 5 is rounded up, away from 0
//	if the quotient is negative then digit 5 is rounded down, away from 0
//
// Note that precision<0 is allowed as input.
func (__obf_d09f058c0a390c93 Decimal) DivRound(__obf_f3234e094c8ed488 Decimal, __obf_7efed86618b2f272 int32) Decimal {
	// QuoRem already checks initialization
	__obf_1b9e13ef114085d1, __obf_a934ca815da1fc1a := __obf_d09f058c0a390c93.QuoRem(__obf_f3234e094c8ed488, __obf_7efed86618b2f272)
	// the actual rounding decision is based on comparing r*10^precision and d2/2
	// instead compare 2 r 10 ^precision and d2
	var __obf_e486c59cbc93e241 big.Int
	__obf_e486c59cbc93e241.Abs(__obf_a934ca815da1fc1a.__obf_c36b36e18c228e7f)
	__obf_e486c59cbc93e241.Lsh(&__obf_e486c59cbc93e241, 1)
	// now rv2 = abs(r.value) * 2
	__obf_a39cd17941b9200c := Decimal{__obf_c36b36e18c228e7f: &__obf_e486c59cbc93e241, __obf_406325483a83b1fa: __obf_a934ca815da1fc1a.__obf_406325483a83b1fa + __obf_7efed86618b2f272}
	// r2 is now 2 * r * 10 ^ precision
	var __obf_ec75e4ddc71ce85e = __obf_a39cd17941b9200c.Cmp(__obf_f3234e094c8ed488.Abs())

	if __obf_ec75e4ddc71ce85e < 0 {
		return __obf_1b9e13ef114085d1
	}

	if __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.Sign()*__obf_f3234e094c8ed488.__obf_c36b36e18c228e7f.Sign() < 0 {
		return __obf_1b9e13ef114085d1.Sub(New(1, -__obf_7efed86618b2f272))
	}

	return __obf_1b9e13ef114085d1.Add(New(1, -__obf_7efed86618b2f272))
}

// Mod returns d % d2.
func (__obf_d09f058c0a390c93 Decimal) Mod(__obf_f3234e094c8ed488 Decimal) Decimal {
	_, __obf_a934ca815da1fc1a := __obf_d09f058c0a390c93.QuoRem(__obf_f3234e094c8ed488, 0)
	return __obf_a934ca815da1fc1a
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
func (__obf_d09f058c0a390c93 Decimal) Pow(__obf_f3234e094c8ed488 Decimal) Decimal {
	__obf_450569c7fe284678 := __obf_d09f058c0a390c93.Sign()
	__obf_609ea661285d6196 := __obf_f3234e094c8ed488.Sign()

	if __obf_450569c7fe284678 == 0 {
		if __obf_609ea661285d6196 == 0 {
			return Decimal{}
		}
		if __obf_609ea661285d6196 == 1 {
			return Decimal{__obf_1bb56bff7931ff30, 0}
		}
		if __obf_609ea661285d6196 == -1 {
			return Decimal{}
		}
	}

	if __obf_609ea661285d6196 == 0 {
		return Decimal{__obf_25f16531cde494de, 0}
	}

	// TODO: optimize extraction of fractional part
	__obf_ffb29ce7762ffc55 := Decimal{__obf_25f16531cde494de, 0}
	__obf_60ba2b229d4ee5c6, __obf_3d1eececbebb1a33 := __obf_f3234e094c8ed488.QuoRem(__obf_ffb29ce7762ffc55, 0)

	if __obf_450569c7fe284678 == -1 && !__obf_3d1eececbebb1a33.IsZero() {
		return Decimal{}
	}

	__obf_10242900ad553b4b, _ := __obf_d09f058c0a390c93.PowBigInt(__obf_60ba2b229d4ee5c6.__obf_c36b36e18c228e7f)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_3d1eececbebb1a33.__obf_c36b36e18c228e7f.Sign() == 0 {
		return __obf_10242900ad553b4b
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_badbacb99efba1da := __obf_d09f058c0a390c93.NumDigits()
	__obf_9e5c7fbe16e98a1b := __obf_f3234e094c8ed488.NumDigits()

	__obf_7efed86618b2f272 := __obf_badbacb99efba1da

	if __obf_9e5c7fbe16e98a1b > __obf_7efed86618b2f272 {
		__obf_7efed86618b2f272 += __obf_9e5c7fbe16e98a1b
	}

	__obf_7efed86618b2f272 += 6

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_c1b085d7f5661a8b, __obf_c1d612f6bfc234d7 := __obf_d09f058c0a390c93.Abs().Ln(-__obf_d09f058c0a390c93.__obf_406325483a83b1fa + int32(__obf_7efed86618b2f272))
	if __obf_c1d612f6bfc234d7 != nil {
		return Decimal{}
	}

	__obf_c1b085d7f5661a8b = __obf_c1b085d7f5661a8b.Mul(__obf_3d1eececbebb1a33)

	__obf_c1b085d7f5661a8b, __obf_c1d612f6bfc234d7 = __obf_c1b085d7f5661a8b.ExpTaylor(-__obf_d09f058c0a390c93.__obf_406325483a83b1fa + int32(__obf_7efed86618b2f272))
	if __obf_c1d612f6bfc234d7 != nil {
		return Decimal{}
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_df948d36f2e10e0e := __obf_10242900ad553b4b.Mul(__obf_c1b085d7f5661a8b)

	return __obf_df948d36f2e10e0e
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
func (__obf_d09f058c0a390c93 Decimal) PowWithPrecision(__obf_f3234e094c8ed488 Decimal, __obf_7efed86618b2f272 int32) (Decimal, error) {
	__obf_450569c7fe284678 := __obf_d09f058c0a390c93.Sign()
	__obf_609ea661285d6196 := __obf_f3234e094c8ed488.Sign()

	if __obf_450569c7fe284678 == 0 {
		if __obf_609ea661285d6196 == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_609ea661285d6196 == 1 {
			return Decimal{__obf_1bb56bff7931ff30, 0}, nil
		}
		if __obf_609ea661285d6196 == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_609ea661285d6196 == 0 {
		return Decimal{__obf_25f16531cde494de, 0}, nil
	}

	// TODO: optimize extraction of fractional part
	__obf_ffb29ce7762ffc55 := Decimal{__obf_25f16531cde494de, 0}
	__obf_60ba2b229d4ee5c6, __obf_3d1eececbebb1a33 := __obf_f3234e094c8ed488.QuoRem(__obf_ffb29ce7762ffc55, 0)

	if __obf_450569c7fe284678 == -1 && !__obf_3d1eececbebb1a33.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}

	__obf_10242900ad553b4b, _ := __obf_d09f058c0a390c93.__obf_d206a3f11e17859e(__obf_60ba2b229d4ee5c6.__obf_c36b36e18c228e7f, __obf_7efed86618b2f272)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_3d1eececbebb1a33.__obf_c36b36e18c228e7f.Sign() == 0 {
		return __obf_10242900ad553b4b, nil
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_badbacb99efba1da := __obf_d09f058c0a390c93.NumDigits()
	__obf_9e5c7fbe16e98a1b := __obf_f3234e094c8ed488.NumDigits()

	if int32(__obf_badbacb99efba1da) > __obf_7efed86618b2f272 {
		__obf_7efed86618b2f272 = int32(__obf_badbacb99efba1da)
	}
	if int32(__obf_9e5c7fbe16e98a1b) > __obf_7efed86618b2f272 {
		__obf_7efed86618b2f272 += int32(__obf_9e5c7fbe16e98a1b)
	}
	// increase precision by 10 to compensate for errors in further calculations
	__obf_7efed86618b2f272 += 10

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_c1b085d7f5661a8b, __obf_c1d612f6bfc234d7 := __obf_d09f058c0a390c93.Abs().Ln(__obf_7efed86618b2f272)
	if __obf_c1d612f6bfc234d7 != nil {
		return Decimal{}, __obf_c1d612f6bfc234d7
	}

	__obf_c1b085d7f5661a8b = __obf_c1b085d7f5661a8b.Mul(__obf_3d1eececbebb1a33)

	__obf_c1b085d7f5661a8b, __obf_c1d612f6bfc234d7 = __obf_c1b085d7f5661a8b.ExpTaylor(__obf_7efed86618b2f272)
	if __obf_c1d612f6bfc234d7 != nil {
		return Decimal{}, __obf_c1d612f6bfc234d7
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_df948d36f2e10e0e := __obf_10242900ad553b4b.Mul(__obf_c1b085d7f5661a8b)

	return __obf_df948d36f2e10e0e, nil
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
func (__obf_d09f058c0a390c93 Decimal) PowInt32(__obf_406325483a83b1fa int32) (Decimal, error) {
	if __obf_d09f058c0a390c93.IsZero() && __obf_406325483a83b1fa == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_8502632695ccbc17 := __obf_406325483a83b1fa < 0
	__obf_406325483a83b1fa = __obf_37fad7391972833c(__obf_406325483a83b1fa)

	__obf_45b69da0ab68e425, __obf_f1febca7bf443d8f := __obf_d09f058c0a390c93, New(1, 0)

	for __obf_406325483a83b1fa > 0 {
		if __obf_406325483a83b1fa%2 == 1 {
			__obf_f1febca7bf443d8f = __obf_f1febca7bf443d8f.Mul(__obf_45b69da0ab68e425)
		}
		__obf_406325483a83b1fa /= 2

		if __obf_406325483a83b1fa > 0 {
			__obf_45b69da0ab68e425 = __obf_45b69da0ab68e425.Mul(__obf_45b69da0ab68e425)
		}
	}

	if __obf_8502632695ccbc17 {
		return New(1, 0).DivRound(__obf_f1febca7bf443d8f, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_f1febca7bf443d8f, nil
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
func (__obf_d09f058c0a390c93 Decimal) PowBigInt(__obf_406325483a83b1fa *big.Int) (Decimal, error) {
	return __obf_d09f058c0a390c93.__obf_d206a3f11e17859e(__obf_406325483a83b1fa, int32(PowPrecisionNegativeExponent))
}

func (__obf_d09f058c0a390c93 Decimal) __obf_d206a3f11e17859e(__obf_406325483a83b1fa *big.Int, __obf_7efed86618b2f272 int32) (Decimal, error) {
	if __obf_d09f058c0a390c93.IsZero() && __obf_406325483a83b1fa.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_041dde05bc98a284 := new(big.Int).Set(__obf_406325483a83b1fa)
	__obf_8502632695ccbc17 := __obf_406325483a83b1fa.Sign() < 0

	if __obf_8502632695ccbc17 {
		__obf_041dde05bc98a284.Abs(__obf_041dde05bc98a284)
	}

	__obf_45b69da0ab68e425, __obf_f1febca7bf443d8f := __obf_d09f058c0a390c93, New(1, 0)

	for __obf_041dde05bc98a284.Sign() > 0 {
		if __obf_041dde05bc98a284.Bit(0) == 1 {
			__obf_f1febca7bf443d8f = __obf_f1febca7bf443d8f.Mul(__obf_45b69da0ab68e425)
		}
		__obf_041dde05bc98a284.Rsh(__obf_041dde05bc98a284, 1)

		if __obf_041dde05bc98a284.Sign() > 0 {
			__obf_45b69da0ab68e425 = __obf_45b69da0ab68e425.Mul(__obf_45b69da0ab68e425)
		}
	}

	if __obf_8502632695ccbc17 {
		return New(1, 0).DivRound(__obf_f1febca7bf443d8f, __obf_7efed86618b2f272), nil
	}

	return __obf_f1febca7bf443d8f, nil
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
func (__obf_d09f058c0a390c93 Decimal) ExpHullAbrham(__obf_6d8f1071e6a9333f uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_d09f058c0a390c93.IsZero() {
		return Decimal{__obf_25f16531cde494de, 0}, nil
	}

	__obf_4790853bbbd98160 := __obf_6d8f1071e6a9333f

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_cb430c29f24deb1e := __obf_d09f058c0a390c93.Abs().InexactFloat64()
	if __obf_8e6183ef0b08fd20 := __obf_cb430c29f24deb1e / 23; __obf_8e6183ef0b08fd20 > float64(__obf_4790853bbbd98160) && __obf_8e6183ef0b08fd20 < float64(ExpMaxIterations) {
		__obf_4790853bbbd98160 = uint32(math.Ceil(__obf_8e6183ef0b08fd20))
	}

	// fail if abs(d) beyond an over/underflow threshold
	__obf_0d0699f070216fe8 := New(23*int64(__obf_4790853bbbd98160), 0)
	if __obf_d09f058c0a390c93.Abs().Cmp(__obf_0d0699f070216fe8) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}

	// Return 1 if abs(d) small enough; this also avoids later over/underflow
	__obf_f3c86764321764b6 := New(9, -int32(__obf_4790853bbbd98160)-1)
	if __obf_d09f058c0a390c93.Abs().Cmp(__obf_f3c86764321764b6) <= 0 {
		return Decimal{__obf_25f16531cde494de, __obf_d09f058c0a390c93.__obf_406325483a83b1fa}, nil
	}

	// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
	__obf_d8a1f54969413649 := __obf_d09f058c0a390c93.__obf_406325483a83b1fa + int32(__obf_d09f058c0a390c93.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_d8a1f54969413649 < 0 {
		__obf_d8a1f54969413649 = 0
	}

	__obf_c6eff4cc2bd046d5 := New(1, __obf_d8a1f54969413649)                                                                                                                   // reduction factor
	__obf_a934ca815da1fc1a := Decimal{new(big.Int).Set(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f), __obf_d09f058c0a390c93.__obf_406325483a83b1fa - __obf_d8a1f54969413649} // reduced argument
	__obf_43f93d52162c34e4 := int32(__obf_4790853bbbd98160) + __obf_d8a1f54969413649 + 2                                                                                       // precision for calculating the sum

	// Determine n, the number of therms for calculating sum
	// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
	// for solving appropriate equation, along with directed
	// roundings and simple rational bound for log10(p/abs(r))
	__obf_8d3ba6e2ccd60e6d := __obf_a934ca815da1fc1a.Abs().InexactFloat64()
	__obf_17b5bf4105c51c4a := float64(__obf_43f93d52162c34e4)
	__obf_c7f3db0810afc00e := math.Ceil((1.453*__obf_17b5bf4105c51c4a - 1.182) / math.Log10(__obf_17b5bf4105c51c4a/__obf_8d3ba6e2ccd60e6d))
	if __obf_c7f3db0810afc00e > float64(ExpMaxIterations) || math.IsNaN(__obf_c7f3db0810afc00e) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_45b69da0ab68e425 := int64(__obf_c7f3db0810afc00e)

	__obf_bab5f70189aa167a := New(0, 0)
	__obf_0a03d5c951bd51ca := New(1, 0)
	__obf_ffb29ce7762ffc55 := New(1, 0)
	for __obf_a6733ea50196cc53 := __obf_45b69da0ab68e425 - 1; __obf_a6733ea50196cc53 > 0; __obf_a6733ea50196cc53-- {
		__obf_bab5f70189aa167a.__obf_c36b36e18c228e7f.SetInt64(__obf_a6733ea50196cc53)
		__obf_0a03d5c951bd51ca = __obf_0a03d5c951bd51ca.Mul(__obf_a934ca815da1fc1a.DivRound(__obf_bab5f70189aa167a, __obf_43f93d52162c34e4))
		__obf_0a03d5c951bd51ca = __obf_0a03d5c951bd51ca.Add(__obf_ffb29ce7762ffc55)
	}

	__obf_3ee8d6ce3886b697 := __obf_c6eff4cc2bd046d5.IntPart()
	__obf_df948d36f2e10e0e := New(1, 0)
	for __obf_a6733ea50196cc53 := __obf_3ee8d6ce3886b697; __obf_a6733ea50196cc53 > 0; __obf_a6733ea50196cc53-- {
		__obf_df948d36f2e10e0e = __obf_df948d36f2e10e0e.Mul(__obf_0a03d5c951bd51ca)
	}

	__obf_03b4b32eb90e503d := int32(__obf_df948d36f2e10e0e.NumDigits())

	var __obf_9e00dd9d89629d72 int32
	if __obf_03b4b32eb90e503d > __obf_37fad7391972833c(__obf_df948d36f2e10e0e.__obf_406325483a83b1fa) {
		__obf_9e00dd9d89629d72 = int32(__obf_4790853bbbd98160) - __obf_03b4b32eb90e503d - __obf_df948d36f2e10e0e.__obf_406325483a83b1fa
	} else {
		__obf_9e00dd9d89629d72 = int32(__obf_4790853bbbd98160)
	}

	__obf_df948d36f2e10e0e = __obf_df948d36f2e10e0e.Round(__obf_9e00dd9d89629d72)

	return __obf_df948d36f2e10e0e, nil
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
func (__obf_d09f058c0a390c93 Decimal) ExpTaylor(__obf_7efed86618b2f272 int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_d09f058c0a390c93.IsZero() {
		return Decimal{__obf_25f16531cde494de, 0}.Round(__obf_7efed86618b2f272), nil
	}

	var __obf_082ee8b1a1ba9d2f Decimal
	var __obf_5f57099ccd3b913e int32
	if __obf_7efed86618b2f272 < 0 {
		__obf_082ee8b1a1ba9d2f = New(1, -1)
		__obf_5f57099ccd3b913e = 8
	} else {
		__obf_082ee8b1a1ba9d2f = New(1, -__obf_7efed86618b2f272-1)
		__obf_5f57099ccd3b913e = __obf_7efed86618b2f272 + 1
	}

	__obf_f4fc06761296e9bb := __obf_d09f058c0a390c93.Abs()
	__obf_c7885e2e751af9ba := __obf_d09f058c0a390c93.Abs()
	__obf_c44ec3d1e6ba1f29 := New(1, 0)

	__obf_f1febca7bf443d8f := New(1, 0)

	for __obf_a6733ea50196cc53 := int64(1); ; {
		__obf_b27e1a204c229d5a := __obf_c7885e2e751af9ba.DivRound(__obf_c44ec3d1e6ba1f29, __obf_5f57099ccd3b913e)
		__obf_f1febca7bf443d8f = __obf_f1febca7bf443d8f.Add(__obf_b27e1a204c229d5a)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_b27e1a204c229d5a.Cmp(__obf_082ee8b1a1ba9d2f) < 0 {
			break
		}

		__obf_c7885e2e751af9ba = __obf_c7885e2e751af9ba.Mul(__obf_f4fc06761296e9bb)

		__obf_a6733ea50196cc53++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_b83126d172f97ebd) >= int(__obf_a6733ea50196cc53) && !__obf_b83126d172f97ebd[__obf_a6733ea50196cc53-1].IsZero() {
			__obf_c44ec3d1e6ba1f29 = __obf_b83126d172f97ebd[__obf_a6733ea50196cc53-1]
		} else {
			// To avoid any race conditions, firstly the zero value is appended to a slice to create
			// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
			// factorial using the index notation.
			__obf_c44ec3d1e6ba1f29 = __obf_b83126d172f97ebd[__obf_a6733ea50196cc53-2].Mul(New(__obf_a6733ea50196cc53, 0))
			__obf_b83126d172f97ebd = append(__obf_b83126d172f97ebd, Zero)
			__obf_b83126d172f97ebd[__obf_a6733ea50196cc53-1] = __obf_c44ec3d1e6ba1f29
		}
	}

	if __obf_d09f058c0a390c93.Sign() < 0 {
		__obf_f1febca7bf443d8f = New(1, 0).DivRound(__obf_f1febca7bf443d8f, __obf_7efed86618b2f272+1)
	}

	__obf_f1febca7bf443d8f = __obf_f1febca7bf443d8f.Round(__obf_7efed86618b2f272)
	return __obf_f1febca7bf443d8f, nil
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
func (__obf_d09f058c0a390c93 Decimal) Ln(__obf_7efed86618b2f272 int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_d09f058c0a390c93.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_d09f058c0a390c93.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}

	__obf_1314bb6b90f4d9a1 := __obf_7efed86618b2f272 + 2
	__obf_cfabb27034b7f97a := __obf_d09f058c0a390c93.Copy()

	var __obf_33ed815004e03065, __obf_83031f42b615dc28, __obf_eff104be57722508, __obf_0bd67fe03d37cd38, __obf_6f28deb5fd45d0e8 Decimal
	__obf_33ed815004e03065 = __obf_cfabb27034b7f97a.Sub(Decimal{__obf_25f16531cde494de, 0})
	__obf_83031f42b615dc28 = Decimal{__obf_25f16531cde494de, -1}

	// for decimal in range [0.9, 1.1] where ln(d) is close to 0
	__obf_8245ac3846d79b39 := false

	if __obf_33ed815004e03065.Abs().Cmp(__obf_83031f42b615dc28) <= 0 {
		__obf_8245ac3846d79b39 = true
	} else {
		// reduce input decimal to range [0.1, 1)
		__obf_ad064f0e2e7bced1 := int32(__obf_cfabb27034b7f97a.NumDigits()) + __obf_cfabb27034b7f97a.__obf_406325483a83b1fa
		__obf_cfabb27034b7f97a.__obf_406325483a83b1fa -= __obf_ad064f0e2e7bced1

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_5a26e8654d026182 := __obf_5a26e8654d026182.__obf_eb4cce5de048f0ba(__obf_1314bb6b90f4d9a1)
		__obf_6f28deb5fd45d0e8 = NewFromInt32(__obf_ad064f0e2e7bced1)
		__obf_6f28deb5fd45d0e8 = __obf_6f28deb5fd45d0e8.Mul(__obf_5a26e8654d026182)

		__obf_33ed815004e03065 = __obf_cfabb27034b7f97a.Sub(Decimal{__obf_25f16531cde494de, 0})

		if __obf_33ed815004e03065.Abs().Cmp(__obf_83031f42b615dc28) <= 0 {
			__obf_8245ac3846d79b39 = true
		} else {
			// initial estimate using floats
			__obf_0fe558412d516e5d := __obf_cfabb27034b7f97a.InexactFloat64()
			__obf_33ed815004e03065 = NewFromFloat(math.Log(__obf_0fe558412d516e5d))
		}
	}

	__obf_082ee8b1a1ba9d2f := Decimal{__obf_25f16531cde494de, -__obf_1314bb6b90f4d9a1}

	if __obf_8245ac3846d79b39 {
		// Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
		// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
		// until the difference between current and next term is smaller than epsilon.
		// Coverage quite fast for decimals close to 1.0

		// z + 2
		__obf_eff104be57722508 = __obf_33ed815004e03065.Add(Decimal{__obf_3cbd576e784ed478, 0})
		// z / (z + 2)
		__obf_83031f42b615dc28 = __obf_33ed815004e03065.DivRound(__obf_eff104be57722508, __obf_1314bb6b90f4d9a1)
		// 2 * (z / (z + 2))
		__obf_33ed815004e03065 = __obf_83031f42b615dc28.Add(__obf_83031f42b615dc28)
		__obf_eff104be57722508 = __obf_33ed815004e03065.Copy()

		for __obf_45b69da0ab68e425 := 1; ; __obf_45b69da0ab68e425++ {
			// 2 * (z / (z+2))^(2n+1)
			__obf_eff104be57722508 = __obf_eff104be57722508.Mul(__obf_83031f42b615dc28).Mul(__obf_83031f42b615dc28)

			// 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
			__obf_0bd67fe03d37cd38 = NewFromInt(int64(2*__obf_45b69da0ab68e425 + 1))
			__obf_0bd67fe03d37cd38 = __obf_eff104be57722508.DivRound(__obf_0bd67fe03d37cd38, __obf_1314bb6b90f4d9a1)

			// comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			__obf_33ed815004e03065 = __obf_33ed815004e03065.Add(__obf_0bd67fe03d37cd38)

			if __obf_0bd67fe03d37cd38.Abs().Cmp(__obf_082ee8b1a1ba9d2f) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_63a91f9fa58cdee0 Decimal
		__obf_3fb8bd9af117c09c := __obf_1314bb6b90f4d9a1*2 + 10

		for __obf_a6733ea50196cc53 := int32(0); __obf_a6733ea50196cc53 < __obf_3fb8bd9af117c09c; __obf_a6733ea50196cc53++ {
			// exp(a_n)
			__obf_83031f42b615dc28, _ = __obf_33ed815004e03065.ExpTaylor(__obf_1314bb6b90f4d9a1)
			// exp(a_n) - z
			__obf_eff104be57722508 = __obf_83031f42b615dc28.Sub(__obf_cfabb27034b7f97a)
			// 2 * (exp(a_n) - z)
			__obf_eff104be57722508 = __obf_eff104be57722508.Add(__obf_eff104be57722508)
			// exp(a_n) + z
			__obf_0bd67fe03d37cd38 = __obf_83031f42b615dc28.Add(__obf_cfabb27034b7f97a)
			// 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_83031f42b615dc28 = __obf_eff104be57722508.DivRound(__obf_0bd67fe03d37cd38, __obf_1314bb6b90f4d9a1)
			// comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_33ed815004e03065 = __obf_33ed815004e03065.Sub(__obf_83031f42b615dc28)

			if __obf_63a91f9fa58cdee0.Add(__obf_83031f42b615dc28).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_83031f42b615dc28.Abs().Cmp(__obf_082ee8b1a1ba9d2f) <= 0 {
				break
			}

			__obf_63a91f9fa58cdee0 = __obf_83031f42b615dc28
		}
	}

	__obf_33ed815004e03065 = __obf_33ed815004e03065.Add(__obf_6f28deb5fd45d0e8)

	return __obf_33ed815004e03065.Round(__obf_7efed86618b2f272), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_d09f058c0a390c93 Decimal) NumDigits() int {
	if __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f == nil {
		return 1
	}

	if __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.IsInt64() {
		__obf_3e637496a1a1f908 := __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.Int64()
		// restrict fast path to integers with exact conversion to float64
		if __obf_3e637496a1a1f908 <= (1<<53) && __obf_3e637496a1a1f908 >= -(1<<53) {
			if __obf_3e637496a1a1f908 == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_3e637496a1a1f908)))) + 1
		}
	}

	__obf_e823e6d5b3406aa7 := int(float64(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.BitLen()) / math.Log2(10))

	// estimatedNumDigits (lg10) may be off by 1, need to verify
	__obf_a279893fa3ef34b6 := big.NewInt(int64(__obf_e823e6d5b3406aa7))
	__obf_3fa5f0f6cf1d9ba7 := __obf_a279893fa3ef34b6.Exp(__obf_44ff59a8ed385a0e, __obf_a279893fa3ef34b6, nil)

	if __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.CmpAbs(__obf_3fa5f0f6cf1d9ba7) >= 0 {
		return __obf_e823e6d5b3406aa7 + 1
	}

	return __obf_e823e6d5b3406aa7
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_d09f058c0a390c93 Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_a934ca815da1fc1a big.Int
	__obf_1b9e13ef114085d1 := new(big.Int).Set(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f)
	for __obf_cfabb27034b7f97a := __obf_37fad7391972833c(__obf_d09f058c0a390c93.__obf_406325483a83b1fa); __obf_cfabb27034b7f97a > 0; __obf_cfabb27034b7f97a-- {
		__obf_1b9e13ef114085d1.QuoRem(__obf_1b9e13ef114085d1, __obf_44ff59a8ed385a0e, &__obf_a934ca815da1fc1a)
		if __obf_a934ca815da1fc1a.Cmp(__obf_1bb56bff7931ff30) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_37fad7391972833c(__obf_45b69da0ab68e425 int32) int32 {
	if __obf_45b69da0ab68e425 < 0 {
		return -__obf_45b69da0ab68e425
	}
	return __obf_45b69da0ab68e425
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_d09f058c0a390c93 Decimal) Cmp(__obf_f3234e094c8ed488 Decimal) int {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()
	__obf_f3234e094c8ed488.__obf_b119326c257b1b2a()

	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa == __obf_f3234e094c8ed488.__obf_406325483a83b1fa {
		return __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.Cmp(__obf_f3234e094c8ed488.__obf_c36b36e18c228e7f)
	}

	__obf_c1fa73dffb7b10fa, __obf_3e0ad0062bf745cb := RescalePair(__obf_d09f058c0a390c93, __obf_f3234e094c8ed488)

	return __obf_c1fa73dffb7b10fa.__obf_c36b36e18c228e7f.Cmp(__obf_3e0ad0062bf745cb.__obf_c36b36e18c228e7f)
}

// Compare compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_d09f058c0a390c93 Decimal) Compare(__obf_f3234e094c8ed488 Decimal) int {
	return __obf_d09f058c0a390c93.Cmp(__obf_f3234e094c8ed488)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_d09f058c0a390c93 Decimal) Equal(__obf_f3234e094c8ed488 Decimal) bool {
	return __obf_d09f058c0a390c93.Cmp(__obf_f3234e094c8ed488) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_d09f058c0a390c93 Decimal) Equals(__obf_f3234e094c8ed488 Decimal) bool {
	return __obf_d09f058c0a390c93.Equal(__obf_f3234e094c8ed488)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_d09f058c0a390c93 Decimal) GreaterThan(__obf_f3234e094c8ed488 Decimal) bool {
	return __obf_d09f058c0a390c93.Cmp(__obf_f3234e094c8ed488) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_d09f058c0a390c93 Decimal) GreaterThanOrEqual(__obf_f3234e094c8ed488 Decimal) bool {
	__obf_fabbd5845ceb048e := __obf_d09f058c0a390c93.Cmp(__obf_f3234e094c8ed488)
	return __obf_fabbd5845ceb048e == 1 || __obf_fabbd5845ceb048e == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_d09f058c0a390c93 Decimal) LessThan(__obf_f3234e094c8ed488 Decimal) bool {
	return __obf_d09f058c0a390c93.Cmp(__obf_f3234e094c8ed488) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_d09f058c0a390c93 Decimal) LessThanOrEqual(__obf_f3234e094c8ed488 Decimal) bool {
	__obf_fabbd5845ceb048e := __obf_d09f058c0a390c93.Cmp(__obf_f3234e094c8ed488)
	return __obf_fabbd5845ceb048e == -1 || __obf_fabbd5845ceb048e == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_d09f058c0a390c93 Decimal) Sign() int {
	if __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f == nil {
		return 0
	}
	return __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.Sign()
}

// IsPositive return
//
//	true if d > 0
//	false if d == 0
//	false if d < 0
func (__obf_d09f058c0a390c93 Decimal) IsPositive() bool {
	return __obf_d09f058c0a390c93.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_d09f058c0a390c93 Decimal) IsNegative() bool {
	return __obf_d09f058c0a390c93.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_d09f058c0a390c93 Decimal) IsZero() bool {
	return __obf_d09f058c0a390c93.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_d09f058c0a390c93 Decimal) Exponent() int32 {
	return __obf_d09f058c0a390c93.__obf_406325483a83b1fa
}

// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
func (__obf_d09f058c0a390c93 Decimal) Coefficient() *big.Int {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f)
}

// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
// If coefficient cannot be represented in an int64, the result will be undefined.
func (__obf_d09f058c0a390c93 Decimal) CoefficientInt64() int64 {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()
	return __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.Int64()
}

// IntPart returns the integer component of the decimal.
func (__obf_d09f058c0a390c93 Decimal) IntPart() int64 {
	__obf_1319cc31f48d3a28 := __obf_d09f058c0a390c93.__obf_d44df924a1b98873(0)
	return __obf_1319cc31f48d3a28.__obf_c36b36e18c228e7f.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_d09f058c0a390c93 Decimal) BigInt() *big.Int {
	__obf_1319cc31f48d3a28 := __obf_d09f058c0a390c93.__obf_d44df924a1b98873(0)
	return __obf_1319cc31f48d3a28.__obf_c36b36e18c228e7f
}

// BigFloat returns decimal as BigFloat.
// Be aware that casting decimal to BigFloat might cause a loss of precision.
func (__obf_d09f058c0a390c93 Decimal) BigFloat() *big.Float {
	__obf_cb430c29f24deb1e := &big.Float{}
	__obf_cb430c29f24deb1e.SetString(__obf_d09f058c0a390c93.String())
	return __obf_cb430c29f24deb1e
}

// Rat returns a rational number representation of the decimal.
func (__obf_d09f058c0a390c93 Decimal) Rat() *big.Rat {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()
	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa <= 0 {
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_3fd45fa1be8c5bee := new(big.Int).Exp(__obf_44ff59a8ed385a0e, big.NewInt(-int64(__obf_d09f058c0a390c93.__obf_406325483a83b1fa)), nil)
		return new(big.Rat).SetFrac(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f, __obf_3fd45fa1be8c5bee)
	}

	__obf_06d411512358dedf := new(big.Int).Exp(__obf_44ff59a8ed385a0e, big.NewInt(int64(__obf_d09f058c0a390c93.__obf_406325483a83b1fa)), nil)
	__obf_e7f08852b714628e := new(big.Int).Mul(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f, __obf_06d411512358dedf)
	return new(big.Rat).SetFrac(__obf_e7f08852b714628e, __obf_25f16531cde494de)
}

// Float64 returns the nearest float64 value for d and a bool indicating
// whether f represents d exactly.
// For more details, see the documentation for big.Rat.Float64
func (__obf_d09f058c0a390c93 Decimal) Float64() (__obf_cb430c29f24deb1e float64, __obf_1d7338aa878aa2d3 bool) {
	return __obf_d09f058c0a390c93.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_d09f058c0a390c93 Decimal) InexactFloat64() float64 {
	__obf_cb430c29f24deb1e, _ := __obf_d09f058c0a390c93.Float64()
	return __obf_cb430c29f24deb1e
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
func (__obf_d09f058c0a390c93 Decimal) String() string {
	return __obf_d09f058c0a390c93.string(true)
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
func (__obf_d09f058c0a390c93 Decimal) StringFixed(__obf_c97f9440e7da0c32 int32) string {
	__obf_2a6386939d30aa57 := __obf_d09f058c0a390c93.Round(__obf_c97f9440e7da0c32)
	return __obf_2a6386939d30aa57.string(false)
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
func (__obf_d09f058c0a390c93 Decimal) StringFixedBank(__obf_c97f9440e7da0c32 int32) string {
	__obf_2a6386939d30aa57 := __obf_d09f058c0a390c93.RoundBank(__obf_c97f9440e7da0c32)
	return __obf_2a6386939d30aa57.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_d09f058c0a390c93 Decimal) StringFixedCash(__obf_62b88200d4864d59 uint8) string {
	__obf_2a6386939d30aa57 := __obf_d09f058c0a390c93.RoundCash(__obf_62b88200d4864d59)
	return __obf_2a6386939d30aa57.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_d09f058c0a390c93 Decimal) Round(__obf_c97f9440e7da0c32 int32) Decimal {
	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa == -__obf_c97f9440e7da0c32 {
		return __obf_d09f058c0a390c93
	}
	// truncate to places + 1
	__obf_afc53812d03631be := __obf_d09f058c0a390c93.__obf_d44df924a1b98873(-__obf_c97f9440e7da0c32 - 1)

	// add sign(d) * 0.5
	if __obf_afc53812d03631be.__obf_c36b36e18c228e7f.Sign() < 0 {
		__obf_afc53812d03631be.__obf_c36b36e18c228e7f.Sub(__obf_afc53812d03631be.__obf_c36b36e18c228e7f, __obf_dc2da84e9cc8100a)
	} else {
		__obf_afc53812d03631be.__obf_c36b36e18c228e7f.Add(__obf_afc53812d03631be.__obf_c36b36e18c228e7f, __obf_dc2da84e9cc8100a)
	}

	// floor for positive numbers, ceil for negative numbers
	_, __obf_b041e2c85e138a64 := __obf_afc53812d03631be.__obf_c36b36e18c228e7f.DivMod(__obf_afc53812d03631be.__obf_c36b36e18c228e7f, __obf_44ff59a8ed385a0e, new(big.Int))
	__obf_afc53812d03631be.__obf_406325483a83b1fa++
	if __obf_afc53812d03631be.__obf_c36b36e18c228e7f.Sign() < 0 && __obf_b041e2c85e138a64.Cmp(__obf_1bb56bff7931ff30) != 0 {
		__obf_afc53812d03631be.__obf_c36b36e18c228e7f.Add(__obf_afc53812d03631be.__obf_c36b36e18c228e7f, __obf_25f16531cde494de)
	}

	return __obf_afc53812d03631be
}

// RoundCeil rounds the decimal towards +infinity.
//
// Example:
//
//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
func (__obf_d09f058c0a390c93 Decimal) RoundCeil(__obf_c97f9440e7da0c32 int32) Decimal {
	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa >= -__obf_c97f9440e7da0c32 {
		return __obf_d09f058c0a390c93
	}

	__obf_0b19174f046cd670 := __obf_d09f058c0a390c93.__obf_d44df924a1b98873(-__obf_c97f9440e7da0c32)
	if __obf_d09f058c0a390c93.Equal(__obf_0b19174f046cd670) {
		return __obf_d09f058c0a390c93
	}

	if __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.Sign() > 0 {
		__obf_0b19174f046cd670.__obf_c36b36e18c228e7f.Add(__obf_0b19174f046cd670.__obf_c36b36e18c228e7f, __obf_25f16531cde494de)
	}

	return __obf_0b19174f046cd670
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_d09f058c0a390c93 Decimal) RoundFloor(__obf_c97f9440e7da0c32 int32) Decimal {
	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa >= -__obf_c97f9440e7da0c32 {
		return __obf_d09f058c0a390c93
	}

	__obf_0b19174f046cd670 := __obf_d09f058c0a390c93.__obf_d44df924a1b98873(-__obf_c97f9440e7da0c32)
	if __obf_d09f058c0a390c93.Equal(__obf_0b19174f046cd670) {
		return __obf_d09f058c0a390c93
	}

	if __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.Sign() < 0 {
		__obf_0b19174f046cd670.__obf_c36b36e18c228e7f.Sub(__obf_0b19174f046cd670.__obf_c36b36e18c228e7f, __obf_25f16531cde494de)
	}

	return __obf_0b19174f046cd670
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_d09f058c0a390c93 Decimal) RoundUp(__obf_c97f9440e7da0c32 int32) Decimal {
	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa >= -__obf_c97f9440e7da0c32 {
		return __obf_d09f058c0a390c93
	}

	__obf_0b19174f046cd670 := __obf_d09f058c0a390c93.__obf_d44df924a1b98873(-__obf_c97f9440e7da0c32)
	if __obf_d09f058c0a390c93.Equal(__obf_0b19174f046cd670) {
		return __obf_d09f058c0a390c93
	}

	if __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.Sign() > 0 {
		__obf_0b19174f046cd670.__obf_c36b36e18c228e7f.Add(__obf_0b19174f046cd670.__obf_c36b36e18c228e7f, __obf_25f16531cde494de)
	} else if __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.Sign() < 0 {
		__obf_0b19174f046cd670.__obf_c36b36e18c228e7f.Sub(__obf_0b19174f046cd670.__obf_c36b36e18c228e7f, __obf_25f16531cde494de)
	}

	return __obf_0b19174f046cd670
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_d09f058c0a390c93 Decimal) RoundDown(__obf_c97f9440e7da0c32 int32) Decimal {
	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa >= -__obf_c97f9440e7da0c32 {
		return __obf_d09f058c0a390c93
	}

	__obf_0b19174f046cd670 := __obf_d09f058c0a390c93.__obf_d44df924a1b98873(-__obf_c97f9440e7da0c32)
	if __obf_d09f058c0a390c93.Equal(__obf_0b19174f046cd670) {
		return __obf_d09f058c0a390c93
	}
	return __obf_0b19174f046cd670
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
func (__obf_d09f058c0a390c93 Decimal) RoundBank(__obf_c97f9440e7da0c32 int32) Decimal {

	__obf_603401f5bd2df87a := __obf_d09f058c0a390c93.Round(__obf_c97f9440e7da0c32)
	__obf_0b98c67f785f60dc := __obf_d09f058c0a390c93.Sub(__obf_603401f5bd2df87a).Abs()

	__obf_34d8fed473f6c6f8 := New(5, -__obf_c97f9440e7da0c32-1)
	if __obf_0b98c67f785f60dc.Cmp(__obf_34d8fed473f6c6f8) == 0 && __obf_603401f5bd2df87a.__obf_c36b36e18c228e7f.Bit(0) != 0 {
		if __obf_603401f5bd2df87a.__obf_c36b36e18c228e7f.Sign() < 0 {
			__obf_603401f5bd2df87a.__obf_c36b36e18c228e7f.Add(__obf_603401f5bd2df87a.__obf_c36b36e18c228e7f, __obf_25f16531cde494de)
		} else {
			__obf_603401f5bd2df87a.__obf_c36b36e18c228e7f.Sub(__obf_603401f5bd2df87a.__obf_c36b36e18c228e7f, __obf_25f16531cde494de)
		}
	}

	return __obf_603401f5bd2df87a
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
func (__obf_d09f058c0a390c93 Decimal) RoundCash(__obf_62b88200d4864d59 uint8) Decimal {
	var __obf_4bc619935b3fdcd5 *big.Int
	switch __obf_62b88200d4864d59 {
	case 5:
		__obf_4bc619935b3fdcd5 = __obf_3316a3319b2a8b33
	case 10:
		__obf_4bc619935b3fdcd5 = __obf_44ff59a8ed385a0e
	case 25:
		__obf_4bc619935b3fdcd5 = __obf_e1dd4002d73ad6cd
	case 50:
		__obf_4bc619935b3fdcd5 = __obf_3cbd576e784ed478
	case 100:
		__obf_4bc619935b3fdcd5 = __obf_25f16531cde494de
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_62b88200d4864d59))
	}
	__obf_ff9aa3fb34ef1c3e := Decimal{
		__obf_c36b36e18c228e7f: __obf_4bc619935b3fdcd5,
	}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_d09f058c0a390c93.Mul(__obf_ff9aa3fb34ef1c3e).Round(0).Div(__obf_ff9aa3fb34ef1c3e).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_d09f058c0a390c93 Decimal) Floor() Decimal {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()

	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa >= 0 {
		return __obf_d09f058c0a390c93
	}

	__obf_406325483a83b1fa := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_406325483a83b1fa.Exp(__obf_406325483a83b1fa, big.NewInt(-int64(__obf_d09f058c0a390c93.__obf_406325483a83b1fa)), nil)

	__obf_cfabb27034b7f97a := new(big.Int).Div(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f, __obf_406325483a83b1fa)
	return Decimal{__obf_c36b36e18c228e7f: __obf_cfabb27034b7f97a, __obf_406325483a83b1fa: 0}
}

// Ceil returns the nearest integer value greater than or equal to d.
func (__obf_d09f058c0a390c93 Decimal) Ceil() Decimal {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()

	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa >= 0 {
		return __obf_d09f058c0a390c93
	}

	__obf_406325483a83b1fa := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_406325483a83b1fa.Exp(__obf_406325483a83b1fa, big.NewInt(-int64(__obf_d09f058c0a390c93.__obf_406325483a83b1fa)), nil)

	__obf_cfabb27034b7f97a, __obf_b041e2c85e138a64 := new(big.Int).DivMod(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f, __obf_406325483a83b1fa, new(big.Int))
	if __obf_b041e2c85e138a64.Cmp(__obf_1bb56bff7931ff30) != 0 {
		__obf_cfabb27034b7f97a.Add(__obf_cfabb27034b7f97a, __obf_25f16531cde494de)
	}
	return Decimal{__obf_c36b36e18c228e7f: __obf_cfabb27034b7f97a, __obf_406325483a83b1fa: 0}
}

// Truncate truncates off digits from the number, without rounding.
//
// NOTE: precision is the last digit that will not be truncated (must be >= 0).
//
// Example:
//
//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
func (__obf_d09f058c0a390c93 Decimal) Truncate(__obf_7efed86618b2f272 int32) Decimal {
	__obf_d09f058c0a390c93.__obf_b119326c257b1b2a()
	if __obf_7efed86618b2f272 >= 0 && -__obf_7efed86618b2f272 > __obf_d09f058c0a390c93.__obf_406325483a83b1fa {
		return __obf_d09f058c0a390c93.__obf_d44df924a1b98873(-__obf_7efed86618b2f272)
	}
	return __obf_d09f058c0a390c93
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_d09f058c0a390c93 *Decimal) UnmarshalJSON(__obf_e8fd8b794f91b308 []byte) error {
	if string(__obf_e8fd8b794f91b308) == "null" {
		return nil
	}

	__obf_d27522cdebc9fff3, __obf_c1d612f6bfc234d7 := __obf_4404c46e45e621bb(__obf_e8fd8b794f91b308)
	if __obf_c1d612f6bfc234d7 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_e8fd8b794f91b308, __obf_c1d612f6bfc234d7)
	}

	__obf_0962dc77c6b6239b, __obf_c1d612f6bfc234d7 := NewFromString(__obf_d27522cdebc9fff3)
	*__obf_d09f058c0a390c93 = __obf_0962dc77c6b6239b
	if __obf_c1d612f6bfc234d7 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_d27522cdebc9fff3, __obf_c1d612f6bfc234d7)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_d09f058c0a390c93 Decimal) MarshalJSON() ([]byte, error) {
	var __obf_d27522cdebc9fff3 string
	if MarshalJSONWithoutQuotes {
		__obf_d27522cdebc9fff3 = __obf_d09f058c0a390c93.String()
	} else {
		__obf_d27522cdebc9fff3 = "\"" + __obf_d09f058c0a390c93.String() + "\""
	}
	return []byte(__obf_d27522cdebc9fff3), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_d09f058c0a390c93 *Decimal) UnmarshalBinary(__obf_76bba392325e4533 []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_76bba392325e4533) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_76bba392325e4533, len(__obf_76bba392325e4533))
	}

	// Extract the exponent
	__obf_d09f058c0a390c93.__obf_406325483a83b1fa = int32(binary.BigEndian.Uint32(__obf_76bba392325e4533[:4]))

	// Extract the value
	__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f = new(big.Int)
	if __obf_c1d612f6bfc234d7 := __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.GobDecode(__obf_76bba392325e4533[4:]); __obf_c1d612f6bfc234d7 != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_76bba392325e4533, __obf_c1d612f6bfc234d7)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_d09f058c0a390c93 Decimal) MarshalBinary() (__obf_76bba392325e4533 []byte, __obf_c1d612f6bfc234d7 error) {
	// exp is written first, but encode value first to know output size
	var __obf_53881788b7c6f89e []byte
	if __obf_53881788b7c6f89e, __obf_c1d612f6bfc234d7 = __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.GobEncode(); __obf_c1d612f6bfc234d7 != nil {
		return nil, __obf_c1d612f6bfc234d7
	}

	// Write the exponent in front, since it's a fixed size
	__obf_0110d9f08cbe5049 := make([]byte, 4, len(__obf_53881788b7c6f89e)+4)
	binary.BigEndian.PutUint32(__obf_0110d9f08cbe5049, uint32(__obf_d09f058c0a390c93.__obf_406325483a83b1fa))

	// Return the byte array
	return append(__obf_0110d9f08cbe5049, __obf_53881788b7c6f89e...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_d09f058c0a390c93 *Decimal) Scan(__obf_c36b36e18c228e7f any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_a07140d89ee7df26 := __obf_c36b36e18c228e7f.(type) {

	case float32:
		*__obf_d09f058c0a390c93 = NewFromFloat(float64(__obf_a07140d89ee7df26))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_d09f058c0a390c93 = NewFromFloat(__obf_a07140d89ee7df26)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_d09f058c0a390c93 = New(__obf_a07140d89ee7df26, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_d09f058c0a390c93 = NewFromUint64(__obf_a07140d89ee7df26)
		return nil

	default:
		// default is trying to interpret value stored as string
		__obf_d27522cdebc9fff3, __obf_c1d612f6bfc234d7 := __obf_4404c46e45e621bb(__obf_a07140d89ee7df26)
		if __obf_c1d612f6bfc234d7 != nil {
			return __obf_c1d612f6bfc234d7
		}
		*__obf_d09f058c0a390c93, __obf_c1d612f6bfc234d7 = NewFromString(__obf_d27522cdebc9fff3)
		return __obf_c1d612f6bfc234d7
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_d09f058c0a390c93 Decimal) Value() (driver.Value, error) {
	return __obf_d09f058c0a390c93.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_d09f058c0a390c93 *Decimal) UnmarshalText(__obf_56af27e60e31824b []byte) error {
	__obf_d27522cdebc9fff3 := string(__obf_56af27e60e31824b)

	__obf_bc636ce96a8a7277, __obf_c1d612f6bfc234d7 := NewFromString(__obf_d27522cdebc9fff3)
	*__obf_d09f058c0a390c93 = __obf_bc636ce96a8a7277
	if __obf_c1d612f6bfc234d7 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_d27522cdebc9fff3, __obf_c1d612f6bfc234d7)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_d09f058c0a390c93 Decimal) MarshalText() (__obf_56af27e60e31824b []byte, __obf_c1d612f6bfc234d7 error) {
	return []byte(__obf_d09f058c0a390c93.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_d09f058c0a390c93 Decimal) GobEncode() ([]byte, error) {
	return __obf_d09f058c0a390c93.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_d09f058c0a390c93 *Decimal) GobDecode(__obf_76bba392325e4533 []byte) error {
	return __obf_d09f058c0a390c93.UnmarshalBinary(__obf_76bba392325e4533)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_d09f058c0a390c93 Decimal) StringScaled(__obf_406325483a83b1fa int32) string {
	return __obf_d09f058c0a390c93.__obf_d44df924a1b98873(__obf_406325483a83b1fa).String()
}

func (__obf_d09f058c0a390c93 Decimal) string(__obf_02867e68dfa271d5 bool) string {
	if __obf_d09f058c0a390c93.__obf_406325483a83b1fa >= 0 {
		return __obf_d09f058c0a390c93.__obf_d44df924a1b98873(0).__obf_c36b36e18c228e7f.String()
	}

	__obf_37fad7391972833c := new(big.Int).Abs(__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f)
	__obf_d27522cdebc9fff3 := __obf_37fad7391972833c.String()

	var __obf_b6fa4ff85902af25, __obf_ac722d2c47f80856 string

	// NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
	// and you are on a 32-bit machine. Won't fix this super-edge case.
	__obf_edfebbae2a495a9f := int(__obf_d09f058c0a390c93.__obf_406325483a83b1fa)
	if len(__obf_d27522cdebc9fff3) > -__obf_edfebbae2a495a9f {
		__obf_b6fa4ff85902af25 = __obf_d27522cdebc9fff3[:len(__obf_d27522cdebc9fff3)+__obf_edfebbae2a495a9f]
		__obf_ac722d2c47f80856 = __obf_d27522cdebc9fff3[len(__obf_d27522cdebc9fff3)+__obf_edfebbae2a495a9f:]
	} else {
		__obf_b6fa4ff85902af25 = "0"

		__obf_307eb9c0766f3eee := -__obf_edfebbae2a495a9f - len(__obf_d27522cdebc9fff3)
		__obf_ac722d2c47f80856 = strings.Repeat("0", __obf_307eb9c0766f3eee) + __obf_d27522cdebc9fff3
	}

	if __obf_02867e68dfa271d5 {
		__obf_a6733ea50196cc53 := len(__obf_ac722d2c47f80856) - 1
		for ; __obf_a6733ea50196cc53 >= 0; __obf_a6733ea50196cc53-- {
			if __obf_ac722d2c47f80856[__obf_a6733ea50196cc53] != '0' {
				break
			}
		}
		__obf_ac722d2c47f80856 = __obf_ac722d2c47f80856[:__obf_a6733ea50196cc53+1]
	}

	__obf_dcfebf40f008ef2a := __obf_b6fa4ff85902af25
	if len(__obf_ac722d2c47f80856) > 0 {
		__obf_dcfebf40f008ef2a += "." + __obf_ac722d2c47f80856
	}

	if __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f.Sign() < 0 {
		return "-" + __obf_dcfebf40f008ef2a
	}

	return __obf_dcfebf40f008ef2a
}

func (__obf_d09f058c0a390c93 *Decimal) __obf_b119326c257b1b2a() {
	if __obf_d09f058c0a390c93.__obf_c36b36e18c228e7f == nil {
		__obf_d09f058c0a390c93.__obf_c36b36e18c228e7f = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_601c4c6a2f69ebaf Decimal, __obf_a6bd19a4a2b2d12a ...Decimal) Decimal {
	__obf_37bb74e984e7327e := __obf_601c4c6a2f69ebaf
	for _, __obf_e10a3284e09c0ec8 := range __obf_a6bd19a4a2b2d12a {
		if __obf_e10a3284e09c0ec8.Cmp(__obf_37bb74e984e7327e) < 0 {
			__obf_37bb74e984e7327e = __obf_e10a3284e09c0ec8
		}
	}
	return __obf_37bb74e984e7327e
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_601c4c6a2f69ebaf Decimal, __obf_a6bd19a4a2b2d12a ...Decimal) Decimal {
	__obf_37bb74e984e7327e := __obf_601c4c6a2f69ebaf
	for _, __obf_e10a3284e09c0ec8 := range __obf_a6bd19a4a2b2d12a {
		if __obf_e10a3284e09c0ec8.Cmp(__obf_37bb74e984e7327e) > 0 {
			__obf_37bb74e984e7327e = __obf_e10a3284e09c0ec8
		}
	}
	return __obf_37bb74e984e7327e
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_601c4c6a2f69ebaf Decimal, __obf_a6bd19a4a2b2d12a ...Decimal) Decimal {
	__obf_a63308b9298ca4e5 := __obf_601c4c6a2f69ebaf
	for _, __obf_e10a3284e09c0ec8 := range __obf_a6bd19a4a2b2d12a {
		__obf_a63308b9298ca4e5 = __obf_a63308b9298ca4e5.Add(__obf_e10a3284e09c0ec8)
	}

	return __obf_a63308b9298ca4e5
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_601c4c6a2f69ebaf Decimal, __obf_a6bd19a4a2b2d12a ...Decimal) Decimal {
	__obf_3a8999938925f58d := New(int64(len(__obf_a6bd19a4a2b2d12a)+1), 0)
	__obf_0a03d5c951bd51ca := Sum(__obf_601c4c6a2f69ebaf, __obf_a6bd19a4a2b2d12a...)
	return __obf_0a03d5c951bd51ca.Div(__obf_3a8999938925f58d)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_00e4676552627ab3 Decimal, __obf_f3234e094c8ed488 Decimal) (Decimal, Decimal) {
	__obf_00e4676552627ab3.__obf_b119326c257b1b2a()
	__obf_f3234e094c8ed488.__obf_b119326c257b1b2a()

	if __obf_00e4676552627ab3.__obf_406325483a83b1fa < __obf_f3234e094c8ed488.__obf_406325483a83b1fa {
		return __obf_00e4676552627ab3, __obf_f3234e094c8ed488.__obf_d44df924a1b98873(__obf_00e4676552627ab3.__obf_406325483a83b1fa)
	} else if __obf_00e4676552627ab3.__obf_406325483a83b1fa > __obf_f3234e094c8ed488.__obf_406325483a83b1fa {
		return __obf_00e4676552627ab3.__obf_d44df924a1b98873(__obf_f3234e094c8ed488.__obf_406325483a83b1fa), __obf_f3234e094c8ed488
	}

	return __obf_00e4676552627ab3, __obf_f3234e094c8ed488
}

func __obf_4404c46e45e621bb(__obf_c36b36e18c228e7f any) (string, error) {
	var __obf_677b5d29e57d52c8 []byte

	switch __obf_a07140d89ee7df26 := __obf_c36b36e18c228e7f.(type) {
	case string:
		__obf_677b5d29e57d52c8 = []byte(__obf_a07140d89ee7df26)
	case []byte:
		__obf_677b5d29e57d52c8 = __obf_a07140d89ee7df26
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_c36b36e18c228e7f, __obf_c36b36e18c228e7f)
	}

	// If the amount is quoted, strip the quotes
	if len(__obf_677b5d29e57d52c8) > 2 && __obf_677b5d29e57d52c8[0] == '"' && __obf_677b5d29e57d52c8[len(__obf_677b5d29e57d52c8)-1] == '"' {
		__obf_677b5d29e57d52c8 = __obf_677b5d29e57d52c8[1 : len(__obf_677b5d29e57d52c8)-1]
	}
	return string(__obf_677b5d29e57d52c8), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the datautil.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_d09f058c0a390c93 Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_d09f058c0a390c93,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_d09f058c0a390c93 *NullDecimal) Scan(__obf_c36b36e18c228e7f any) error {
	if __obf_c36b36e18c228e7f == nil {
		__obf_d09f058c0a390c93.Valid = false
		return nil
	}
	__obf_d09f058c0a390c93.Valid = true
	return __obf_d09f058c0a390c93.Decimal.Scan(__obf_c36b36e18c228e7f)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_d09f058c0a390c93 NullDecimal) Value() (driver.Value, error) {
	if !__obf_d09f058c0a390c93.Valid {
		return nil, nil
	}
	return __obf_d09f058c0a390c93.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_d09f058c0a390c93 *NullDecimal) UnmarshalJSON(__obf_e8fd8b794f91b308 []byte) error {
	if string(__obf_e8fd8b794f91b308) == "null" {
		__obf_d09f058c0a390c93.Valid = false
		return nil
	}
	__obf_d09f058c0a390c93.Valid = true
	return __obf_d09f058c0a390c93.Decimal.UnmarshalJSON(__obf_e8fd8b794f91b308)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_d09f058c0a390c93 NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_d09f058c0a390c93.Valid {
		return []byte("null"), nil
	}
	return __obf_d09f058c0a390c93.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_d09f058c0a390c93 *NullDecimal) UnmarshalText(__obf_56af27e60e31824b []byte) error {
	__obf_d27522cdebc9fff3 := string(__obf_56af27e60e31824b)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_d27522cdebc9fff3 == "" {
		__obf_d09f058c0a390c93.Valid = false
		return nil
	}
	if __obf_c1d612f6bfc234d7 := __obf_d09f058c0a390c93.Decimal.UnmarshalText(__obf_56af27e60e31824b); __obf_c1d612f6bfc234d7 != nil {
		__obf_d09f058c0a390c93.Valid = false
		return __obf_c1d612f6bfc234d7
	}
	__obf_d09f058c0a390c93.Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_d09f058c0a390c93 NullDecimal) MarshalText() (__obf_56af27e60e31824b []byte, __obf_c1d612f6bfc234d7 error) {
	if !__obf_d09f058c0a390c93.Valid {
		return []byte{}, nil
	}
	return __obf_d09f058c0a390c93.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_d09f058c0a390c93 Decimal) Atan() Decimal {
	if __obf_d09f058c0a390c93.Equal(NewFromFloat(0.0)) {
		return __obf_d09f058c0a390c93
	}
	if __obf_d09f058c0a390c93.GreaterThan(NewFromFloat(0.0)) {
		return __obf_d09f058c0a390c93.__obf_c211ad731ffd0563()
	}
	return __obf_d09f058c0a390c93.Neg().__obf_c211ad731ffd0563().Neg()
}

func (__obf_d09f058c0a390c93 Decimal) __obf_65e84b2d0f9bf309() Decimal {
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
	__obf_cfabb27034b7f97a := __obf_d09f058c0a390c93.Mul(__obf_d09f058c0a390c93)
	__obf_7d5ecff6d2d5c88d := P0.Mul(__obf_cfabb27034b7f97a).Add(P1).Mul(__obf_cfabb27034b7f97a).Add(P2).Mul(__obf_cfabb27034b7f97a).Add(P3).Mul(__obf_cfabb27034b7f97a).Add(P4).Mul(__obf_cfabb27034b7f97a)
	__obf_64e7317a703fe325 := __obf_cfabb27034b7f97a.Add(Q0).Mul(__obf_cfabb27034b7f97a).Add(Q1).Mul(__obf_cfabb27034b7f97a).Add(Q2).Mul(__obf_cfabb27034b7f97a).Add(Q3).Mul(__obf_cfabb27034b7f97a).Add(Q4)
	__obf_cfabb27034b7f97a = __obf_7d5ecff6d2d5c88d.Div(__obf_64e7317a703fe325)
	__obf_cfabb27034b7f97a = __obf_d09f058c0a390c93.Mul(__obf_cfabb27034b7f97a).Add(__obf_d09f058c0a390c93)
	return __obf_cfabb27034b7f97a
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_d09f058c0a390c93 Decimal) __obf_c211ad731ffd0563() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)      // tan(3*pi/8)
	__obf_daf4071b4f8ca84e := NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_d09f058c0a390c93.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_d09f058c0a390c93.__obf_65e84b2d0f9bf309()
	}
	if __obf_d09f058c0a390c93.GreaterThan(Tan3pio8) {
		return __obf_daf4071b4f8ca84e.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_d09f058c0a390c93).__obf_65e84b2d0f9bf309()).Add(Morebits)
	}
	return __obf_daf4071b4f8ca84e.Div(NewFromFloat(4.0)).Add((__obf_d09f058c0a390c93.Sub(NewFromFloat(1.0)).Div(__obf_d09f058c0a390c93.Add(NewFromFloat(1.0)))).__obf_65e84b2d0f9bf309()).Add(NewFromFloat(0.5).Mul(Morebits))
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
func (__obf_d09f058c0a390c93 Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_d09f058c0a390c93.Equal(NewFromFloat(0.0)) {
		return __obf_d09f058c0a390c93
	}
	// make argument positive but save the sign
	__obf_87165cc42c6992a3 := false
	if __obf_d09f058c0a390c93.LessThan(NewFromFloat(0.0)) {
		__obf_d09f058c0a390c93 = __obf_d09f058c0a390c93.Neg()
		__obf_87165cc42c6992a3 = true
	}

	__obf_2ada872b5e39595f := __obf_d09f058c0a390c93.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_cc34602f8a37fb57 := NewFromFloat(float64(__obf_2ada872b5e39595f)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_2ada872b5e39595f&1 == 1 {
		__obf_2ada872b5e39595f++
		__obf_cc34602f8a37fb57 = __obf_cc34602f8a37fb57.Add(NewFromFloat(1.0))
	}
	__obf_2ada872b5e39595f &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_2ada872b5e39595f > 3 {
		__obf_87165cc42c6992a3 = !__obf_87165cc42c6992a3
		__obf_2ada872b5e39595f -= 4
	}
	__obf_cfabb27034b7f97a := __obf_d09f058c0a390c93.Sub(__obf_cc34602f8a37fb57.Mul(PI4A)).Sub(__obf_cc34602f8a37fb57.Mul(PI4B)).Sub(__obf_cc34602f8a37fb57.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_1b188d7167364580 := __obf_cfabb27034b7f97a.Mul(__obf_cfabb27034b7f97a)

	if __obf_2ada872b5e39595f == 1 || __obf_2ada872b5e39595f == 2 {
		__obf_c981e663ad88d103 := __obf_1b188d7167364580.Mul(__obf_1b188d7167364580).Mul(_cos[0].Mul(__obf_1b188d7167364580).Add(_cos[1]).Mul(__obf_1b188d7167364580).Add(_cos[2]).Mul(__obf_1b188d7167364580).Add(_cos[3]).Mul(__obf_1b188d7167364580).Add(_cos[4]).Mul(__obf_1b188d7167364580).Add(_cos[5]))
		__obf_cc34602f8a37fb57 = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_1b188d7167364580)).Add(__obf_c981e663ad88d103)
	} else {
		__obf_cc34602f8a37fb57 = __obf_cfabb27034b7f97a.Add(__obf_cfabb27034b7f97a.Mul(__obf_1b188d7167364580).Mul(_sin[0].Mul(__obf_1b188d7167364580).Add(_sin[1]).Mul(__obf_1b188d7167364580).Add(_sin[2]).Mul(__obf_1b188d7167364580).Add(_sin[3]).Mul(__obf_1b188d7167364580).Add(_sin[4]).Mul(__obf_1b188d7167364580).Add(_sin[5])))
	}
	if __obf_87165cc42c6992a3 {
		__obf_cc34602f8a37fb57 = __obf_cc34602f8a37fb57.Neg()
	}
	return __obf_cc34602f8a37fb57
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
func (__obf_d09f058c0a390c93 Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	// make argument positive
	__obf_87165cc42c6992a3 := false
	if __obf_d09f058c0a390c93.LessThan(NewFromFloat(0.0)) {
		__obf_d09f058c0a390c93 = __obf_d09f058c0a390c93.Neg()
	}

	__obf_2ada872b5e39595f := __obf_d09f058c0a390c93.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_cc34602f8a37fb57 := NewFromFloat(float64(__obf_2ada872b5e39595f)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_2ada872b5e39595f&1 == 1 {
		__obf_2ada872b5e39595f++
		__obf_cc34602f8a37fb57 = __obf_cc34602f8a37fb57.Add(NewFromFloat(1.0))
	}
	__obf_2ada872b5e39595f &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_2ada872b5e39595f > 3 {
		__obf_87165cc42c6992a3 = !__obf_87165cc42c6992a3
		__obf_2ada872b5e39595f -= 4
	}
	if __obf_2ada872b5e39595f > 1 {
		__obf_87165cc42c6992a3 = !__obf_87165cc42c6992a3
	}

	__obf_cfabb27034b7f97a := __obf_d09f058c0a390c93.Sub(__obf_cc34602f8a37fb57.Mul(PI4A)).Sub(__obf_cc34602f8a37fb57.Mul(PI4B)).Sub(__obf_cc34602f8a37fb57.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_1b188d7167364580 := __obf_cfabb27034b7f97a.Mul(__obf_cfabb27034b7f97a)

	if __obf_2ada872b5e39595f == 1 || __obf_2ada872b5e39595f == 2 {
		__obf_cc34602f8a37fb57 = __obf_cfabb27034b7f97a.Add(__obf_cfabb27034b7f97a.Mul(__obf_1b188d7167364580).Mul(_sin[0].Mul(__obf_1b188d7167364580).Add(_sin[1]).Mul(__obf_1b188d7167364580).Add(_sin[2]).Mul(__obf_1b188d7167364580).Add(_sin[3]).Mul(__obf_1b188d7167364580).Add(_sin[4]).Mul(__obf_1b188d7167364580).Add(_sin[5])))
	} else {
		__obf_c981e663ad88d103 := __obf_1b188d7167364580.Mul(__obf_1b188d7167364580).Mul(_cos[0].Mul(__obf_1b188d7167364580).Add(_cos[1]).Mul(__obf_1b188d7167364580).Add(_cos[2]).Mul(__obf_1b188d7167364580).Add(_cos[3]).Mul(__obf_1b188d7167364580).Add(_cos[4]).Mul(__obf_1b188d7167364580).Add(_cos[5]))
		__obf_cc34602f8a37fb57 = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_1b188d7167364580)).Add(__obf_c981e663ad88d103)
	}
	if __obf_87165cc42c6992a3 {
		__obf_cc34602f8a37fb57 = __obf_cc34602f8a37fb57.Neg()
	}
	return __obf_cc34602f8a37fb57
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
func (__obf_d09f058c0a390c93 Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_d09f058c0a390c93.Equal(NewFromFloat(0.0)) {
		return __obf_d09f058c0a390c93
	}

	// make argument positive but save the sign
	__obf_87165cc42c6992a3 := false
	if __obf_d09f058c0a390c93.LessThan(NewFromFloat(0.0)) {
		__obf_d09f058c0a390c93 = __obf_d09f058c0a390c93.Neg()
		__obf_87165cc42c6992a3 = true
	}

	__obf_2ada872b5e39595f := __obf_d09f058c0a390c93.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_cc34602f8a37fb57 := NewFromFloat(float64(__obf_2ada872b5e39595f)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_2ada872b5e39595f&1 == 1 {
		__obf_2ada872b5e39595f++
		__obf_cc34602f8a37fb57 = __obf_cc34602f8a37fb57.Add(NewFromFloat(1.0))
	}

	__obf_cfabb27034b7f97a := __obf_d09f058c0a390c93.Sub(__obf_cc34602f8a37fb57.Mul(PI4A)).Sub(__obf_cc34602f8a37fb57.Mul(PI4B)).Sub(__obf_cc34602f8a37fb57.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_1b188d7167364580 := __obf_cfabb27034b7f97a.Mul(__obf_cfabb27034b7f97a)

	if __obf_1b188d7167364580.GreaterThan(NewFromFloat(1e-14)) {
		__obf_c981e663ad88d103 := __obf_1b188d7167364580.Mul(_tanP[0].Mul(__obf_1b188d7167364580).Add(_tanP[1]).Mul(__obf_1b188d7167364580).Add(_tanP[2]))
		__obf_1c048090fbae2177 := __obf_1b188d7167364580.Add(_tanQ[1]).Mul(__obf_1b188d7167364580).Add(_tanQ[2]).Mul(__obf_1b188d7167364580).Add(_tanQ[3]).Mul(__obf_1b188d7167364580).Add(_tanQ[4])
		__obf_cc34602f8a37fb57 = __obf_cfabb27034b7f97a.Add(__obf_cfabb27034b7f97a.Mul(__obf_c981e663ad88d103.Div(__obf_1c048090fbae2177)))
	} else {
		__obf_cc34602f8a37fb57 = __obf_cfabb27034b7f97a
	}
	if __obf_2ada872b5e39595f&2 == 2 {
		__obf_cc34602f8a37fb57 = NewFromFloat(-1.0).Div(__obf_cc34602f8a37fb57)
	}
	if __obf_87165cc42c6992a3 {
		__obf_cc34602f8a37fb57 = __obf_cc34602f8a37fb57.Neg()
	}
	return __obf_cc34602f8a37fb57
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

type __obf_0962dc77c6b6239b struct {
	__obf_d09f058c0a390c93 [800]byte // digits, big-endian representation
	__obf_ea04243a9869d25e int       // number of digits used
	__obf_cb19c1fd6e41a0d7 int       // decimal point
	__obf_f0af453d1e176af6 bool      // negative flag
	__obf_6d99ec3c0920e3eb bool      // discarded nonzero digits beyond d[:nd]
}

func (__obf_5097a5efcc046ddc *__obf_0962dc77c6b6239b) String() string {
	__obf_45b69da0ab68e425 := 10 + __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e
	if __obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7 > 0 {
		__obf_45b69da0ab68e425 += __obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7
	}
	if __obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7 < 0 {
		__obf_45b69da0ab68e425 += -__obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7
	}

	__obf_93542e6c18fdd6ca := make([]byte, __obf_45b69da0ab68e425)
	__obf_c981e663ad88d103 := 0
	switch {
	case __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e == 0:
		return "0"

	case __obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7 <= 0:
		// zeros fill space between decimal point and digits
		__obf_93542e6c18fdd6ca[__obf_c981e663ad88d103] = '0'
		__obf_c981e663ad88d103++
		__obf_93542e6c18fdd6ca[__obf_c981e663ad88d103] = '.'
		__obf_c981e663ad88d103++
		__obf_c981e663ad88d103 += __obf_50fe429044a97b04(__obf_93542e6c18fdd6ca[__obf_c981e663ad88d103 : __obf_c981e663ad88d103+-__obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7])
		__obf_c981e663ad88d103 += copy(__obf_93542e6c18fdd6ca[__obf_c981e663ad88d103:], __obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[0:__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e])

	case __obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7 < __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e:
		// decimal point in middle of digits
		__obf_c981e663ad88d103 += copy(__obf_93542e6c18fdd6ca[__obf_c981e663ad88d103:], __obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[0:__obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7])
		__obf_93542e6c18fdd6ca[__obf_c981e663ad88d103] = '.'
		__obf_c981e663ad88d103++
		__obf_c981e663ad88d103 += copy(__obf_93542e6c18fdd6ca[__obf_c981e663ad88d103:], __obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7:__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e])

	default:
		// zeros fill space between digits and decimal point
		__obf_c981e663ad88d103 += copy(__obf_93542e6c18fdd6ca[__obf_c981e663ad88d103:], __obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[0:__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e])
		__obf_c981e663ad88d103 += __obf_50fe429044a97b04(__obf_93542e6c18fdd6ca[__obf_c981e663ad88d103 : __obf_c981e663ad88d103+__obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7-__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e])
	}
	return string(__obf_93542e6c18fdd6ca[0:__obf_c981e663ad88d103])
}

func __obf_50fe429044a97b04(__obf_d8e4295fe41206b1 []byte) int {
	for __obf_a6733ea50196cc53 := range __obf_d8e4295fe41206b1 {
		__obf_d8e4295fe41206b1[__obf_a6733ea50196cc53] = '0'
	}
	return len(__obf_d8e4295fe41206b1)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_80a150c45f6664a5(__obf_5097a5efcc046ddc *__obf_0962dc77c6b6239b) {
	for __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e > 0 && __obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e-1] == '0' {
		__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e--
	}
	if __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e == 0 {
		__obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7 = 0
	}
}

// Assign v to a.
func (__obf_5097a5efcc046ddc *__obf_0962dc77c6b6239b) Assign(__obf_a07140d89ee7df26 uint64) {
	var __obf_93542e6c18fdd6ca [24]byte

	// Write reversed decimal in buf.
	__obf_45b69da0ab68e425 := 0
	for __obf_a07140d89ee7df26 > 0 {
		__obf_b83d0db5c6f6bede := __obf_a07140d89ee7df26 / 10
		__obf_a07140d89ee7df26 -= 10 * __obf_b83d0db5c6f6bede
		__obf_93542e6c18fdd6ca[__obf_45b69da0ab68e425] = byte(__obf_a07140d89ee7df26 + '0')
		__obf_45b69da0ab68e425++
		__obf_a07140d89ee7df26 = __obf_b83d0db5c6f6bede
	}

	// Reverse again to produce forward decimal in a.d.
	__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e = 0
	for __obf_45b69da0ab68e425--; __obf_45b69da0ab68e425 >= 0; __obf_45b69da0ab68e425-- {
		__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e] = __obf_93542e6c18fdd6ca[__obf_45b69da0ab68e425]
		__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e++
	}
	__obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7 = __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e
	__obf_80a150c45f6664a5(__obf_5097a5efcc046ddc)
}

// Maximum shift that we can do in one pass without overflow.
// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
const __obf_6af2b3468807a8d6 = 32 << (^uint(0) >> 63)
const __obf_809dd9d021cb9743 = __obf_6af2b3468807a8d6 - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_5e1efe9fc84c5b3e(__obf_5097a5efcc046ddc *__obf_0962dc77c6b6239b, __obf_c6eff4cc2bd046d5 uint) {
	__obf_a934ca815da1fc1a := 0 // read pointer
	__obf_c981e663ad88d103 := 0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_45b69da0ab68e425 uint
	for ; __obf_45b69da0ab68e425>>__obf_c6eff4cc2bd046d5 == 0; __obf_a934ca815da1fc1a++ {
		if __obf_a934ca815da1fc1a >= __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e {
			if __obf_45b69da0ab68e425 == 0 {
				// a == 0; shouldn't get here, but handle anyway.
				__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e = 0
				return
			}
			for __obf_45b69da0ab68e425>>__obf_c6eff4cc2bd046d5 == 0 {
				__obf_45b69da0ab68e425 = __obf_45b69da0ab68e425 * 10
				__obf_a934ca815da1fc1a++
			}
			break
		}
		__obf_ec75e4ddc71ce85e := uint(__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_a934ca815da1fc1a])
		__obf_45b69da0ab68e425 = __obf_45b69da0ab68e425*10 + __obf_ec75e4ddc71ce85e - '0'
	}
	__obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7 -= __obf_a934ca815da1fc1a - 1

	var __obf_7bb5289c9c89dc1c uint = (1 << __obf_c6eff4cc2bd046d5) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_a934ca815da1fc1a < __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e; __obf_a934ca815da1fc1a++ {
		__obf_ec75e4ddc71ce85e := uint(__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_a934ca815da1fc1a])
		__obf_d19bf61c6ae7a20e := __obf_45b69da0ab68e425 >> __obf_c6eff4cc2bd046d5
		__obf_45b69da0ab68e425 &= __obf_7bb5289c9c89dc1c
		__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_c981e663ad88d103] = byte(__obf_d19bf61c6ae7a20e + '0')
		__obf_c981e663ad88d103++
		__obf_45b69da0ab68e425 = __obf_45b69da0ab68e425*10 + __obf_ec75e4ddc71ce85e - '0'
	}

	// Put down extra digits.
	for __obf_45b69da0ab68e425 > 0 {
		__obf_d19bf61c6ae7a20e := __obf_45b69da0ab68e425 >> __obf_c6eff4cc2bd046d5
		__obf_45b69da0ab68e425 &= __obf_7bb5289c9c89dc1c
		if __obf_c981e663ad88d103 < len(__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93) {
			__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_c981e663ad88d103] = byte(__obf_d19bf61c6ae7a20e + '0')
			__obf_c981e663ad88d103++
		} else if __obf_d19bf61c6ae7a20e > 0 {
			__obf_5097a5efcc046ddc.__obf_6d99ec3c0920e3eb = true
		}
		__obf_45b69da0ab68e425 = __obf_45b69da0ab68e425 * 10
	}

	__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e = __obf_c981e663ad88d103
	__obf_80a150c45f6664a5(__obf_5097a5efcc046ddc)
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

type __obf_60b087b69279ef70 struct {
	__obf_e25b9a4ddc40a521 int    // number of new digits
	__obf_ac7aaf94727b9dec string // minus one digit if original < a.
}

var __obf_e6b729c8a19aa74f = []__obf_60b087b69279ef70{
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
func __obf_0c709a396e42b86e(__obf_5acc254c98f3fd5e []byte, __obf_539943ae97e52628 string) bool {
	for __obf_a6733ea50196cc53 := 0; __obf_a6733ea50196cc53 < len(__obf_539943ae97e52628); __obf_a6733ea50196cc53++ {
		if __obf_a6733ea50196cc53 >= len(__obf_5acc254c98f3fd5e) {
			return true
		}
		if __obf_5acc254c98f3fd5e[__obf_a6733ea50196cc53] != __obf_539943ae97e52628[__obf_a6733ea50196cc53] {
			return __obf_5acc254c98f3fd5e[__obf_a6733ea50196cc53] < __obf_539943ae97e52628[__obf_a6733ea50196cc53]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_7f131afdfb9c0003(__obf_5097a5efcc046ddc *__obf_0962dc77c6b6239b, __obf_c6eff4cc2bd046d5 uint) {
	__obf_e25b9a4ddc40a521 := __obf_e6b729c8a19aa74f[__obf_c6eff4cc2bd046d5].__obf_e25b9a4ddc40a521
	if __obf_0c709a396e42b86e(__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[0:__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e], __obf_e6b729c8a19aa74f[__obf_c6eff4cc2bd046d5].__obf_ac7aaf94727b9dec) {
		__obf_e25b9a4ddc40a521--
	}

	__obf_a934ca815da1fc1a := __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e                          // read index
	__obf_c981e663ad88d103 := __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e + __obf_e25b9a4ddc40a521 // write index

	// Pick up a digit, put down a digit.
	var __obf_45b69da0ab68e425 uint
	for __obf_a934ca815da1fc1a--; __obf_a934ca815da1fc1a >= 0; __obf_a934ca815da1fc1a-- {
		__obf_45b69da0ab68e425 += (uint(__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_a934ca815da1fc1a]) - '0') << __obf_c6eff4cc2bd046d5
		__obf_2133c767fa1629c8 := __obf_45b69da0ab68e425 / 10
		__obf_7c7926c5ab5f5cee := __obf_45b69da0ab68e425 - 10*__obf_2133c767fa1629c8
		__obf_c981e663ad88d103--
		if __obf_c981e663ad88d103 < len(__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93) {
			__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_c981e663ad88d103] = byte(__obf_7c7926c5ab5f5cee + '0')
		} else if __obf_7c7926c5ab5f5cee != 0 {
			__obf_5097a5efcc046ddc.__obf_6d99ec3c0920e3eb = true
		}
		__obf_45b69da0ab68e425 = __obf_2133c767fa1629c8
	}

	// Put down extra digits.
	for __obf_45b69da0ab68e425 > 0 {
		__obf_2133c767fa1629c8 := __obf_45b69da0ab68e425 / 10
		__obf_7c7926c5ab5f5cee := __obf_45b69da0ab68e425 - 10*__obf_2133c767fa1629c8
		__obf_c981e663ad88d103--
		if __obf_c981e663ad88d103 < len(__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93) {
			__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_c981e663ad88d103] = byte(__obf_7c7926c5ab5f5cee + '0')
		} else if __obf_7c7926c5ab5f5cee != 0 {
			__obf_5097a5efcc046ddc.__obf_6d99ec3c0920e3eb = true
		}
		__obf_45b69da0ab68e425 = __obf_2133c767fa1629c8
	}

	__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e += __obf_e25b9a4ddc40a521
	if __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e >= len(__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93) {
		__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e = len(__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93)
	}
	__obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7 += __obf_e25b9a4ddc40a521
	__obf_80a150c45f6664a5(__obf_5097a5efcc046ddc)
}

// Binary shift left (k > 0) or right (k < 0).
func (__obf_5097a5efcc046ddc *__obf_0962dc77c6b6239b) Shift(__obf_c6eff4cc2bd046d5 int) {
	switch {
	case __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e == 0:
		// nothing to do: a == 0
	case __obf_c6eff4cc2bd046d5 > 0:
		for __obf_c6eff4cc2bd046d5 > __obf_809dd9d021cb9743 {
			__obf_7f131afdfb9c0003(__obf_5097a5efcc046ddc, __obf_809dd9d021cb9743)
			__obf_c6eff4cc2bd046d5 -= __obf_809dd9d021cb9743
		}
		__obf_7f131afdfb9c0003(__obf_5097a5efcc046ddc, uint(__obf_c6eff4cc2bd046d5))
	case __obf_c6eff4cc2bd046d5 < 0:
		for __obf_c6eff4cc2bd046d5 < -__obf_809dd9d021cb9743 {
			__obf_5e1efe9fc84c5b3e(__obf_5097a5efcc046ddc, __obf_809dd9d021cb9743)
			__obf_c6eff4cc2bd046d5 += __obf_809dd9d021cb9743
		}
		__obf_5e1efe9fc84c5b3e(__obf_5097a5efcc046ddc, uint(-__obf_c6eff4cc2bd046d5))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_5366ed222e70c315(__obf_5097a5efcc046ddc *__obf_0962dc77c6b6239b, __obf_ea04243a9869d25e int) bool {
	if __obf_ea04243a9869d25e < 0 || __obf_ea04243a9869d25e >= __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e {
		return false
	}
	if __obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_ea04243a9869d25e] == '5' && __obf_ea04243a9869d25e+1 == __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e { // exactly halfway - round to even
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_5097a5efcc046ddc.__obf_6d99ec3c0920e3eb {
			return true
		}
		return __obf_ea04243a9869d25e > 0 && (__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_ea04243a9869d25e-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_ea04243a9869d25e] >= '5'
}

// Round a to nd digits (or fewer).
// If nd is zero, it means we're rounding
// just to the left of the digits, as in
// 0.09 -> 0.1.
func (__obf_5097a5efcc046ddc *__obf_0962dc77c6b6239b) Round(__obf_ea04243a9869d25e int) {
	if __obf_ea04243a9869d25e < 0 || __obf_ea04243a9869d25e >= __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e {
		return
	}
	if __obf_5366ed222e70c315(__obf_5097a5efcc046ddc, __obf_ea04243a9869d25e) {
		__obf_5097a5efcc046ddc.RoundUp(__obf_ea04243a9869d25e)
	} else {
		__obf_5097a5efcc046ddc.RoundDown(__obf_ea04243a9869d25e)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_5097a5efcc046ddc *__obf_0962dc77c6b6239b) RoundDown(__obf_ea04243a9869d25e int) {
	if __obf_ea04243a9869d25e < 0 || __obf_ea04243a9869d25e >= __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e {
		return
	}
	__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e = __obf_ea04243a9869d25e
	__obf_80a150c45f6664a5(__obf_5097a5efcc046ddc)
}

// Round a up to nd digits (or fewer).
func (__obf_5097a5efcc046ddc *__obf_0962dc77c6b6239b) RoundUp(__obf_ea04243a9869d25e int) {
	if __obf_ea04243a9869d25e < 0 || __obf_ea04243a9869d25e >= __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e {
		return
	}

	// round up
	for __obf_a6733ea50196cc53 := __obf_ea04243a9869d25e - 1; __obf_a6733ea50196cc53 >= 0; __obf_a6733ea50196cc53-- {
		__obf_ec75e4ddc71ce85e := __obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_a6733ea50196cc53]
		if __obf_ec75e4ddc71ce85e < '9' { // can stop after this digit
			__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_a6733ea50196cc53]++
			__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e = __obf_a6733ea50196cc53 + 1
			return
		}
	}

	// Number is all 9s.
	// Change to single 1 with adjusted decimal point.
	__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[0] = '1'
	__obf_5097a5efcc046ddc.__obf_ea04243a9869d25e = 1
	__obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7++
}

// Extract integer part, rounded appropriately.
// No guarantees about overflow.
func (__obf_5097a5efcc046ddc *__obf_0962dc77c6b6239b) RoundedInteger() uint64 {
	if __obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7 > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_a6733ea50196cc53 int
	__obf_45b69da0ab68e425 := uint64(0)
	for __obf_a6733ea50196cc53 = 0; __obf_a6733ea50196cc53 < __obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7 && __obf_a6733ea50196cc53 < __obf_5097a5efcc046ddc.__obf_ea04243a9869d25e; __obf_a6733ea50196cc53++ {
		__obf_45b69da0ab68e425 = __obf_45b69da0ab68e425*10 + uint64(__obf_5097a5efcc046ddc.__obf_d09f058c0a390c93[__obf_a6733ea50196cc53]-'0')
	}
	for ; __obf_a6733ea50196cc53 < __obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7; __obf_a6733ea50196cc53++ {
		__obf_45b69da0ab68e425 *= 10
	}
	if __obf_5366ed222e70c315(__obf_5097a5efcc046ddc, __obf_5097a5efcc046ddc.__obf_cb19c1fd6e41a0d7) {
		__obf_45b69da0ab68e425++
	}
	return __obf_45b69da0ab68e425
}
