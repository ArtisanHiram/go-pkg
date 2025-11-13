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
package __obf_adf9c80aee40a9d8

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

var __obf_31e6998552abcf35 = big.NewInt(0)
var __obf_6dcccdf86673e16e = big.NewInt(1)
var __obf_97304e664e9a31ae = big.NewInt(2)
var __obf_c16b63cd3d3ece01 = big.NewInt(4)
var __obf_84ba5bd50a08ec11 = big.NewInt(5)
var __obf_a9bf241bb08f75b7 = big.NewInt(10)
var __obf_d3e2e1fd11685001 = big.NewInt(20)

var __obf_f94a2e25be307681 = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_a1a825437ca7e6ad *big.Int

	// NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.
	__obf_fed6fd50f0fcc834 int32
}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_a1a825437ca7e6ad int64, __obf_fed6fd50f0fcc834 int32) Decimal {
	return Decimal{
		__obf_a1a825437ca7e6ad: big.NewInt(__obf_a1a825437ca7e6ad),
		__obf_fed6fd50f0fcc834: __obf_fed6fd50f0fcc834,
	}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_a1a825437ca7e6ad int64) Decimal {
	return Decimal{
		__obf_a1a825437ca7e6ad: big.NewInt(__obf_a1a825437ca7e6ad),
		__obf_fed6fd50f0fcc834: 0,
	}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_a1a825437ca7e6ad int32) Decimal {
	return Decimal{
		__obf_a1a825437ca7e6ad: big.NewInt(int64(__obf_a1a825437ca7e6ad)),
		__obf_fed6fd50f0fcc834: 0,
	}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_a1a825437ca7e6ad uint64) Decimal {
	return Decimal{
		__obf_a1a825437ca7e6ad: new(big.Int).SetUint64(__obf_a1a825437ca7e6ad),
		__obf_fed6fd50f0fcc834: 0,
	}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_a1a825437ca7e6ad *big.Int, __obf_fed6fd50f0fcc834 int32) Decimal {
	return Decimal{
		__obf_a1a825437ca7e6ad: new(big.Int).Set(__obf_a1a825437ca7e6ad),
		__obf_fed6fd50f0fcc834: __obf_fed6fd50f0fcc834,
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
func NewFromBigRat(__obf_a1a825437ca7e6ad *big.Rat, __obf_60fb543883293d44 int32) Decimal {
	return Decimal{
		__obf_a1a825437ca7e6ad: new(big.Int).Set(__obf_a1a825437ca7e6ad.Num()),
		__obf_fed6fd50f0fcc834: 0,
	}.DivRound(Decimal{
		__obf_a1a825437ca7e6ad: new(big.Int).Set(__obf_a1a825437ca7e6ad.Denom()),
		__obf_fed6fd50f0fcc834: 0,
	}, __obf_60fb543883293d44)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_a1a825437ca7e6ad string) (Decimal, error) {
	__obf_9273c838eea41d8f := __obf_a1a825437ca7e6ad
	var __obf_69c1a6c444c381c9 string
	var __obf_fed6fd50f0fcc834 int64

	// Check if number is using scientific notation
	__obf_bf71caa27716e3b7 := strings.IndexAny(__obf_a1a825437ca7e6ad, "Ee")
	if __obf_bf71caa27716e3b7 != -1 {
		__obf_ff50a3675fd3eaaf, __obf_db664194ea89d0f0 := strconv.ParseInt(__obf_a1a825437ca7e6ad[__obf_bf71caa27716e3b7+1:], 10, 32)
		if __obf_db664194ea89d0f0 != nil {
			if __obf_b88acbdaaf930d31, __obf_4322384e3e2aafa1 := __obf_db664194ea89d0f0.(*strconv.NumError); __obf_4322384e3e2aafa1 && __obf_b88acbdaaf930d31.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_a1a825437ca7e6ad)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_a1a825437ca7e6ad)
		}
		__obf_a1a825437ca7e6ad = __obf_a1a825437ca7e6ad[:__obf_bf71caa27716e3b7]
		__obf_fed6fd50f0fcc834 = __obf_ff50a3675fd3eaaf
	}

	__obf_8dd318fa1e2ddd50 := -1
	__obf_fcccdd0a86d38268 := len(__obf_a1a825437ca7e6ad)
	for __obf_159ec70197834006 := 0; __obf_159ec70197834006 < __obf_fcccdd0a86d38268; __obf_159ec70197834006++ {
		if __obf_a1a825437ca7e6ad[__obf_159ec70197834006] == '.' {
			if __obf_8dd318fa1e2ddd50 > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_a1a825437ca7e6ad)
			}
			__obf_8dd318fa1e2ddd50 = __obf_159ec70197834006
		}
	}

	if __obf_8dd318fa1e2ddd50 == -1 {
		// There is no decimal point, we can just parse the original string as
		// an int
		__obf_69c1a6c444c381c9 = __obf_a1a825437ca7e6ad
	} else {
		if __obf_8dd318fa1e2ddd50+1 < __obf_fcccdd0a86d38268 {
			__obf_69c1a6c444c381c9 = __obf_a1a825437ca7e6ad[:__obf_8dd318fa1e2ddd50] + __obf_a1a825437ca7e6ad[__obf_8dd318fa1e2ddd50+1:]
		} else {
			__obf_69c1a6c444c381c9 = __obf_a1a825437ca7e6ad[:__obf_8dd318fa1e2ddd50]
		}
		__obf_ff50a3675fd3eaaf := -len(__obf_a1a825437ca7e6ad[__obf_8dd318fa1e2ddd50+1:])
		__obf_fed6fd50f0fcc834 += int64(__obf_ff50a3675fd3eaaf)
	}

	var __obf_a37fa1d8a4fe3682 *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_69c1a6c444c381c9) <= 18 {
		__obf_b22f421b61330b0a, __obf_db664194ea89d0f0 := strconv.ParseInt(__obf_69c1a6c444c381c9, 10, 64)
		if __obf_db664194ea89d0f0 != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_a1a825437ca7e6ad)
		}
		__obf_a37fa1d8a4fe3682 = big.NewInt(__obf_b22f421b61330b0a)
	} else {
		__obf_a37fa1d8a4fe3682 = new(big.Int)
		_, __obf_4322384e3e2aafa1 := __obf_a37fa1d8a4fe3682.SetString(__obf_69c1a6c444c381c9, 10)
		if !__obf_4322384e3e2aafa1 {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_a1a825437ca7e6ad)
		}
	}

	if __obf_fed6fd50f0fcc834 < math.MinInt32 || __obf_fed6fd50f0fcc834 > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_9273c838eea41d8f)
	}

	return Decimal{
		__obf_a1a825437ca7e6ad: __obf_a37fa1d8a4fe3682,
		__obf_fed6fd50f0fcc834: int32(__obf_fed6fd50f0fcc834),
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
func NewFromFormattedString(__obf_a1a825437ca7e6ad string, __obf_3d048b685f838d49 *regexp.Regexp) (Decimal, error) {
	__obf_66e8f58fdca33842 := __obf_3d048b685f838d49.ReplaceAllString(__obf_a1a825437ca7e6ad, "")
	__obf_c4701b3bb28cd2ae, __obf_db664194ea89d0f0 := NewFromString(__obf_66e8f58fdca33842)
	if __obf_db664194ea89d0f0 != nil {
		return Decimal{}, __obf_db664194ea89d0f0
	}
	return __obf_c4701b3bb28cd2ae, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_a1a825437ca7e6ad string) Decimal {
	__obf_aaf3c5b40a6a1b7f, __obf_db664194ea89d0f0 := NewFromString(__obf_a1a825437ca7e6ad)
	if __obf_db664194ea89d0f0 != nil {
		panic(__obf_db664194ea89d0f0)
	}
	return __obf_aaf3c5b40a6a1b7f
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
func NewFromFloat(__obf_a1a825437ca7e6ad float64) Decimal {
	if __obf_a1a825437ca7e6ad == 0 {
		return New(0, 0)
	}
	return __obf_24fa4939d72941df(__obf_a1a825437ca7e6ad, math.Float64bits(__obf_a1a825437ca7e6ad), &__obf_b6eec24024f8885c)
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
func NewFromFloat32(__obf_a1a825437ca7e6ad float32) Decimal {
	if __obf_a1a825437ca7e6ad == 0 {
		return New(0, 0)
	}
	// XOR is workaround for https://github.com/golang/go/issues/26285
	__obf_935771d361ffde13 := math.Float32bits(__obf_a1a825437ca7e6ad) ^ 0x80808080
	return __obf_24fa4939d72941df(float64(__obf_a1a825437ca7e6ad), uint64(__obf_935771d361ffde13)^0x80808080, &__obf_90dee1d4d7d96500)
}

func __obf_24fa4939d72941df(__obf_8eb59897cac4cd31 float64, __obf_5e67f5dc968af230 uint64, __obf_cdb3e6cf09ad181f *__obf_84acfd4cf591f3ed) Decimal {
	if math.IsNaN(__obf_8eb59897cac4cd31) || math.IsInf(__obf_8eb59897cac4cd31, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_8eb59897cac4cd31))
	}
	__obf_fed6fd50f0fcc834 := int(__obf_5e67f5dc968af230>>__obf_cdb3e6cf09ad181f.__obf_415e8862356d404c) & (1<<__obf_cdb3e6cf09ad181f.__obf_926f241eb4411d53 - 1)
	__obf_15edc752c636a4a0 := __obf_5e67f5dc968af230 & (uint64(1)<<__obf_cdb3e6cf09ad181f.__obf_415e8862356d404c - 1)

	switch __obf_fed6fd50f0fcc834 {
	case 0:
		// denormalized
		__obf_fed6fd50f0fcc834++

	default:
		// add implicit top bit
		__obf_15edc752c636a4a0 |= uint64(1) << __obf_cdb3e6cf09ad181f.__obf_415e8862356d404c
	}
	__obf_fed6fd50f0fcc834 += __obf_cdb3e6cf09ad181f.__obf_3ff00e3b0682ef4a

	var __obf_c4701b3bb28cd2ae __obf_adf9c80aee40a9d8
	__obf_c4701b3bb28cd2ae.Assign(__obf_15edc752c636a4a0)
	__obf_c4701b3bb28cd2ae.Shift(__obf_fed6fd50f0fcc834 - int(__obf_cdb3e6cf09ad181f.__obf_415e8862356d404c))
	__obf_c4701b3bb28cd2ae.__obf_468f39cd022d1bd4 = __obf_5e67f5dc968af230>>(__obf_cdb3e6cf09ad181f.__obf_926f241eb4411d53+__obf_cdb3e6cf09ad181f.__obf_415e8862356d404c) != 0

	__obf_da99d1421cbe52f7(&__obf_c4701b3bb28cd2ae, __obf_15edc752c636a4a0, __obf_fed6fd50f0fcc834, __obf_cdb3e6cf09ad181f)
	// If less than 19 digits, we can do calculation in an int64.
	if __obf_c4701b3bb28cd2ae.__obf_087eecd15d19676e < 19 {
		__obf_fd8e5bcfddaceac6 := int64(0)
		__obf_4739b75ffb6404bf := int64(1)
		for __obf_159ec70197834006 := __obf_c4701b3bb28cd2ae.__obf_087eecd15d19676e - 1; __obf_159ec70197834006 >= 0; __obf_159ec70197834006-- {
			__obf_fd8e5bcfddaceac6 += __obf_4739b75ffb6404bf * int64(__obf_c4701b3bb28cd2ae.__obf_c4701b3bb28cd2ae[__obf_159ec70197834006]-'0')
			__obf_4739b75ffb6404bf *= 10
		}
		if __obf_c4701b3bb28cd2ae.__obf_468f39cd022d1bd4 {
			__obf_fd8e5bcfddaceac6 *= -1
		}
		return Decimal{__obf_a1a825437ca7e6ad: big.NewInt(__obf_fd8e5bcfddaceac6), __obf_fed6fd50f0fcc834: int32(__obf_c4701b3bb28cd2ae.__obf_d30835ea2dfed4ba) - int32(__obf_c4701b3bb28cd2ae.__obf_087eecd15d19676e)}
	}
	__obf_a37fa1d8a4fe3682 := new(big.Int)
	__obf_a37fa1d8a4fe3682, __obf_4322384e3e2aafa1 := __obf_a37fa1d8a4fe3682.SetString(string(__obf_c4701b3bb28cd2ae.__obf_c4701b3bb28cd2ae[:__obf_c4701b3bb28cd2ae.__obf_087eecd15d19676e]), 10)
	if __obf_4322384e3e2aafa1 {
		return Decimal{__obf_a1a825437ca7e6ad: __obf_a37fa1d8a4fe3682, __obf_fed6fd50f0fcc834: int32(__obf_c4701b3bb28cd2ae.__obf_d30835ea2dfed4ba) - int32(__obf_c4701b3bb28cd2ae.__obf_087eecd15d19676e)}
	}

	return NewFromFloatWithExponent(__obf_8eb59897cac4cd31, int32(__obf_c4701b3bb28cd2ae.__obf_d30835ea2dfed4ba)-int32(__obf_c4701b3bb28cd2ae.__obf_087eecd15d19676e))
}

// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
// number of fractional digits.
//
// Example:
//
//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
func NewFromFloatWithExponent(__obf_a1a825437ca7e6ad float64, __obf_fed6fd50f0fcc834 int32) Decimal {
	if math.IsNaN(__obf_a1a825437ca7e6ad) || math.IsInf(__obf_a1a825437ca7e6ad, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_a1a825437ca7e6ad))
	}

	__obf_5e67f5dc968af230 := math.Float64bits(__obf_a1a825437ca7e6ad)
	__obf_15edc752c636a4a0 := __obf_5e67f5dc968af230 & (1<<52 - 1)
	__obf_9371df33d5dc1013 := int32((__obf_5e67f5dc968af230 >> 52) & (1<<11 - 1))
	__obf_81ae3aeca8ee3212 := __obf_5e67f5dc968af230 >> 63

	if __obf_9371df33d5dc1013 == 0 {
		// specials
		if __obf_15edc752c636a4a0 == 0 {
			return Decimal{}
		}
		// subnormal
		__obf_9371df33d5dc1013++
	} else {
		// normal
		__obf_15edc752c636a4a0 |= 1 << 52
	}

	__obf_9371df33d5dc1013 -= 1023 + 52

	// normalizing base-2 values
	for __obf_15edc752c636a4a0&1 == 0 {
		__obf_15edc752c636a4a0 = __obf_15edc752c636a4a0 >> 1
		__obf_9371df33d5dc1013++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_fed6fd50f0fcc834 < 0 && __obf_fed6fd50f0fcc834 < __obf_9371df33d5dc1013 {
		if __obf_9371df33d5dc1013 < 0 {
			__obf_fed6fd50f0fcc834 = __obf_9371df33d5dc1013
		} else {
			__obf_fed6fd50f0fcc834 = 0
		}
	}

	// representing 10^M * 2^N as 5^M * 2^(M+N)
	__obf_9371df33d5dc1013 -= __obf_fed6fd50f0fcc834

	__obf_df99691a60199660 := big.NewInt(1)
	__obf_6f15d1b04c797c54 := big.NewInt(int64(__obf_15edc752c636a4a0))

	// applying 5^M
	if __obf_fed6fd50f0fcc834 > 0 {
		__obf_df99691a60199660 = __obf_df99691a60199660.SetInt64(int64(__obf_fed6fd50f0fcc834))
		__obf_df99691a60199660 = __obf_df99691a60199660.Exp(__obf_84ba5bd50a08ec11, __obf_df99691a60199660, nil)
	} else if __obf_fed6fd50f0fcc834 < 0 {
		__obf_df99691a60199660 = __obf_df99691a60199660.SetInt64(-int64(__obf_fed6fd50f0fcc834))
		__obf_df99691a60199660 = __obf_df99691a60199660.Exp(__obf_84ba5bd50a08ec11, __obf_df99691a60199660, nil)
		__obf_6f15d1b04c797c54 = __obf_6f15d1b04c797c54.Mul(__obf_6f15d1b04c797c54, __obf_df99691a60199660)
		__obf_df99691a60199660 = __obf_df99691a60199660.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_9371df33d5dc1013 > 0 {
		__obf_6f15d1b04c797c54 = __obf_6f15d1b04c797c54.Lsh(__obf_6f15d1b04c797c54, uint(__obf_9371df33d5dc1013))
	} else if __obf_9371df33d5dc1013 < 0 {
		__obf_df99691a60199660 = __obf_df99691a60199660.Lsh(__obf_df99691a60199660, uint(-__obf_9371df33d5dc1013))
	}

	// rounding and downscaling
	if __obf_fed6fd50f0fcc834 > 0 || __obf_9371df33d5dc1013 < 0 {
		__obf_870e500da8c714a2 := new(big.Int).Rsh(__obf_df99691a60199660, 1)
		__obf_6f15d1b04c797c54 = __obf_6f15d1b04c797c54.Add(__obf_6f15d1b04c797c54, __obf_870e500da8c714a2)
		__obf_6f15d1b04c797c54 = __obf_6f15d1b04c797c54.Quo(__obf_6f15d1b04c797c54, __obf_df99691a60199660)
	}

	if __obf_81ae3aeca8ee3212 == 1 {
		__obf_6f15d1b04c797c54 = __obf_6f15d1b04c797c54.Neg(__obf_6f15d1b04c797c54)
	}

	return Decimal{
		__obf_a1a825437ca7e6ad: __obf_6f15d1b04c797c54,
		__obf_fed6fd50f0fcc834: __obf_fed6fd50f0fcc834,
	}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_c4701b3bb28cd2ae Decimal) Copy() Decimal {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()
	return Decimal{
		__obf_a1a825437ca7e6ad: new(big.Int).Set(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad),
		__obf_fed6fd50f0fcc834: __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834,
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
func (__obf_c4701b3bb28cd2ae Decimal) __obf_452282ed04deedc9(__obf_fed6fd50f0fcc834 int32) Decimal {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()

	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 == __obf_fed6fd50f0fcc834 {
		return Decimal{
			new(big.Int).Set(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad),
			__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834,
		}
	}

	// NOTE(vadim): must convert exps to float64 before - to prevent overflow
	__obf_78903e6b3f42a1ec := math.Abs(float64(__obf_fed6fd50f0fcc834) - float64(__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834))
	__obf_a1a825437ca7e6ad := new(big.Int).Set(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad)

	__obf_2511d71159986c18 := new(big.Int).Exp(__obf_a9bf241bb08f75b7, big.NewInt(int64(__obf_78903e6b3f42a1ec)), nil)
	if __obf_fed6fd50f0fcc834 > __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 {
		__obf_a1a825437ca7e6ad = __obf_a1a825437ca7e6ad.Quo(__obf_a1a825437ca7e6ad, __obf_2511d71159986c18)
	} else if __obf_fed6fd50f0fcc834 < __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 {
		__obf_a1a825437ca7e6ad = __obf_a1a825437ca7e6ad.Mul(__obf_a1a825437ca7e6ad, __obf_2511d71159986c18)
	}

	return Decimal{
		__obf_a1a825437ca7e6ad: __obf_a1a825437ca7e6ad,
		__obf_fed6fd50f0fcc834: __obf_fed6fd50f0fcc834,
	}
}

// Abs returns the absolute value of the decimal.
func (__obf_c4701b3bb28cd2ae Decimal) Abs() Decimal {
	if !__obf_c4701b3bb28cd2ae.IsNegative() {
		return __obf_c4701b3bb28cd2ae
	}
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()
	__obf_d7a610286c5dc89c := new(big.Int).Abs(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad)
	return Decimal{
		__obf_a1a825437ca7e6ad: __obf_d7a610286c5dc89c,
		__obf_fed6fd50f0fcc834: __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834,
	}
}

// Add returns d + d2.
func (__obf_c4701b3bb28cd2ae Decimal) Add(__obf_5c7cf1148c3170fc Decimal) Decimal {
	__obf_cd15ee72b8cc5c32, __obf_ed4756049744a46c := RescalePair(__obf_c4701b3bb28cd2ae, __obf_5c7cf1148c3170fc)

	__obf_38b70c8bf2400b40 := new(big.Int).Add(__obf_cd15ee72b8cc5c32.__obf_a1a825437ca7e6ad, __obf_ed4756049744a46c.__obf_a1a825437ca7e6ad)
	return Decimal{
		__obf_a1a825437ca7e6ad: __obf_38b70c8bf2400b40,
		__obf_fed6fd50f0fcc834: __obf_cd15ee72b8cc5c32.__obf_fed6fd50f0fcc834,
	}
}

// Sub returns d - d2.
func (__obf_c4701b3bb28cd2ae Decimal) Sub(__obf_5c7cf1148c3170fc Decimal) Decimal {
	__obf_cd15ee72b8cc5c32, __obf_ed4756049744a46c := RescalePair(__obf_c4701b3bb28cd2ae, __obf_5c7cf1148c3170fc)

	__obf_38b70c8bf2400b40 := new(big.Int).Sub(__obf_cd15ee72b8cc5c32.__obf_a1a825437ca7e6ad, __obf_ed4756049744a46c.__obf_a1a825437ca7e6ad)
	return Decimal{
		__obf_a1a825437ca7e6ad: __obf_38b70c8bf2400b40,
		__obf_fed6fd50f0fcc834: __obf_cd15ee72b8cc5c32.__obf_fed6fd50f0fcc834,
	}
}

// Neg returns -d.
func (__obf_c4701b3bb28cd2ae Decimal) Neg() Decimal {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()
	__obf_8eb59897cac4cd31 := new(big.Int).Neg(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad)
	return Decimal{
		__obf_a1a825437ca7e6ad: __obf_8eb59897cac4cd31,
		__obf_fed6fd50f0fcc834: __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834,
	}
}

// Mul returns d * d2.
func (__obf_c4701b3bb28cd2ae Decimal) Mul(__obf_5c7cf1148c3170fc Decimal) Decimal {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()
	__obf_5c7cf1148c3170fc.__obf_4b849233a95eee86()

	__obf_8c12ef89201a21f7 := int64(__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834) + int64(__obf_5c7cf1148c3170fc.__obf_fed6fd50f0fcc834)
	if __obf_8c12ef89201a21f7 > math.MaxInt32 || __obf_8c12ef89201a21f7 < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_8c12ef89201a21f7))
	}

	__obf_38b70c8bf2400b40 := new(big.Int).Mul(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad, __obf_5c7cf1148c3170fc.__obf_a1a825437ca7e6ad)
	return Decimal{
		__obf_a1a825437ca7e6ad: __obf_38b70c8bf2400b40,
		__obf_fed6fd50f0fcc834: int32(__obf_8c12ef89201a21f7),
	}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_c4701b3bb28cd2ae Decimal) Shift(__obf_8a36add632ffb70c int32) Decimal {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()
	return Decimal{
		__obf_a1a825437ca7e6ad: new(big.Int).Set(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad),
		__obf_fed6fd50f0fcc834: __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 + __obf_8a36add632ffb70c,
	}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_c4701b3bb28cd2ae Decimal) Div(__obf_5c7cf1148c3170fc Decimal) Decimal {
	return __obf_c4701b3bb28cd2ae.DivRound(__obf_5c7cf1148c3170fc, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_c4701b3bb28cd2ae Decimal) QuoRem(__obf_5c7cf1148c3170fc Decimal, __obf_60fb543883293d44 int32) (Decimal, Decimal) {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()
	__obf_5c7cf1148c3170fc.__obf_4b849233a95eee86()
	if __obf_5c7cf1148c3170fc.__obf_a1a825437ca7e6ad.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_8b82908fc680b217 := -__obf_60fb543883293d44
	__obf_b88acbdaaf930d31 := int64(__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834) - int64(__obf_5c7cf1148c3170fc.__obf_fed6fd50f0fcc834) - int64(__obf_8b82908fc680b217)
	if __obf_b88acbdaaf930d31 > math.MaxInt32 || __obf_b88acbdaaf930d31 < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_69d9a22283a7fe10, __obf_f07cc24b854412e9, __obf_7204f82fef28f3b5 big.Int
	var __obf_2402977cd24e43af int32
	// d = a 10^ea
	// d2 = b 10^eb
	if __obf_b88acbdaaf930d31 < 0 {
		__obf_69d9a22283a7fe10 = *__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad
		__obf_7204f82fef28f3b5.SetInt64(-__obf_b88acbdaaf930d31)
		__obf_f07cc24b854412e9.Exp(__obf_a9bf241bb08f75b7, &__obf_7204f82fef28f3b5, nil)
		__obf_f07cc24b854412e9.Mul(__obf_5c7cf1148c3170fc.__obf_a1a825437ca7e6ad, &__obf_f07cc24b854412e9)
		__obf_2402977cd24e43af = __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834
		// now aa = a
		//     bb = b 10^(scale + eb - ea)
	} else {
		__obf_7204f82fef28f3b5.SetInt64(__obf_b88acbdaaf930d31)
		__obf_69d9a22283a7fe10.Exp(__obf_a9bf241bb08f75b7, &__obf_7204f82fef28f3b5, nil)
		__obf_69d9a22283a7fe10.Mul(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad, &__obf_69d9a22283a7fe10)
		__obf_f07cc24b854412e9 = *__obf_5c7cf1148c3170fc.__obf_a1a825437ca7e6ad
		__obf_2402977cd24e43af = __obf_8b82908fc680b217 + __obf_5c7cf1148c3170fc.__obf_fed6fd50f0fcc834
		// now aa = a ^ (ea - eb - scale)
		//     bb = b
	}
	var __obf_9d9315f9ed27915a, __obf_e1eef1cd0f9a218e big.Int
	__obf_9d9315f9ed27915a.QuoRem(&__obf_69d9a22283a7fe10, &__obf_f07cc24b854412e9, &__obf_e1eef1cd0f9a218e)
	__obf_92168af107e3f632 := Decimal{__obf_a1a825437ca7e6ad: &__obf_9d9315f9ed27915a, __obf_fed6fd50f0fcc834: __obf_8b82908fc680b217}
	__obf_37221d27765843ad := Decimal{__obf_a1a825437ca7e6ad: &__obf_e1eef1cd0f9a218e, __obf_fed6fd50f0fcc834: __obf_2402977cd24e43af}
	return __obf_92168af107e3f632, __obf_37221d27765843ad
}

// DivRound divides and rounds to a given precision
// i.e. to an integer multiple of 10^(-precision)
//
//	for a positive quotient digit 5 is rounded up, away from 0
//	if the quotient is negative then digit 5 is rounded down, away from 0
//
// Note that precision<0 is allowed as input.
func (__obf_c4701b3bb28cd2ae Decimal) DivRound(__obf_5c7cf1148c3170fc Decimal, __obf_60fb543883293d44 int32) Decimal {
	// QuoRem already checks initialization
	__obf_9d9315f9ed27915a, __obf_e1eef1cd0f9a218e := __obf_c4701b3bb28cd2ae.QuoRem(__obf_5c7cf1148c3170fc, __obf_60fb543883293d44)
	// the actual rounding decision is based on comparing r*10^precision and d2/2
	// instead compare 2 r 10 ^precision and d2
	var __obf_498eb496e947f2f7 big.Int
	__obf_498eb496e947f2f7.Abs(__obf_e1eef1cd0f9a218e.__obf_a1a825437ca7e6ad)
	__obf_498eb496e947f2f7.Lsh(&__obf_498eb496e947f2f7, 1)
	// now rv2 = abs(r.value) * 2
	__obf_fd550d0c4fe19d92 := Decimal{__obf_a1a825437ca7e6ad: &__obf_498eb496e947f2f7, __obf_fed6fd50f0fcc834: __obf_e1eef1cd0f9a218e.__obf_fed6fd50f0fcc834 + __obf_60fb543883293d44}
	// r2 is now 2 * r * 10 ^ precision
	var __obf_49577c83f724294c = __obf_fd550d0c4fe19d92.Cmp(__obf_5c7cf1148c3170fc.Abs())

	if __obf_49577c83f724294c < 0 {
		return __obf_9d9315f9ed27915a
	}

	if __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.Sign()*__obf_5c7cf1148c3170fc.__obf_a1a825437ca7e6ad.Sign() < 0 {
		return __obf_9d9315f9ed27915a.Sub(New(1, -__obf_60fb543883293d44))
	}

	return __obf_9d9315f9ed27915a.Add(New(1, -__obf_60fb543883293d44))
}

// Mod returns d % d2.
func (__obf_c4701b3bb28cd2ae Decimal) Mod(__obf_5c7cf1148c3170fc Decimal) Decimal {
	_, __obf_e1eef1cd0f9a218e := __obf_c4701b3bb28cd2ae.QuoRem(__obf_5c7cf1148c3170fc, 0)
	return __obf_e1eef1cd0f9a218e
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
func (__obf_c4701b3bb28cd2ae Decimal) Pow(__obf_5c7cf1148c3170fc Decimal) Decimal {
	__obf_f3c9b4b2356d7cc9 := __obf_c4701b3bb28cd2ae.Sign()
	__obf_6d860b7cd06941cc := __obf_5c7cf1148c3170fc.Sign()

	if __obf_f3c9b4b2356d7cc9 == 0 {
		if __obf_6d860b7cd06941cc == 0 {
			return Decimal{}
		}
		if __obf_6d860b7cd06941cc == 1 {
			return Decimal{__obf_31e6998552abcf35, 0}
		}
		if __obf_6d860b7cd06941cc == -1 {
			return Decimal{}
		}
	}

	if __obf_6d860b7cd06941cc == 0 {
		return Decimal{__obf_6dcccdf86673e16e, 0}
	}

	// TODO: optimize extraction of fractional part
	__obf_144c40e8322c0b92 := Decimal{__obf_6dcccdf86673e16e, 0}
	__obf_1235c9530aecc29b, __obf_63b94ab9e37045ad := __obf_5c7cf1148c3170fc.QuoRem(__obf_144c40e8322c0b92, 0)

	if __obf_f3c9b4b2356d7cc9 == -1 && !__obf_63b94ab9e37045ad.IsZero() {
		return Decimal{}
	}

	__obf_21271cbee6f28e58, _ := __obf_c4701b3bb28cd2ae.PowBigInt(__obf_1235c9530aecc29b.__obf_a1a825437ca7e6ad)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_63b94ab9e37045ad.__obf_a1a825437ca7e6ad.Sign() == 0 {
		return __obf_21271cbee6f28e58
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_df164c01bd219b41 := __obf_c4701b3bb28cd2ae.NumDigits()
	__obf_a6902edc3d438234 := __obf_5c7cf1148c3170fc.NumDigits()

	__obf_60fb543883293d44 := __obf_df164c01bd219b41

	if __obf_a6902edc3d438234 > __obf_60fb543883293d44 {
		__obf_60fb543883293d44 += __obf_a6902edc3d438234
	}

	__obf_60fb543883293d44 += 6

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_21394d176152e7ad, __obf_db664194ea89d0f0 := __obf_c4701b3bb28cd2ae.Abs().Ln(-__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 + int32(__obf_60fb543883293d44))
	if __obf_db664194ea89d0f0 != nil {
		return Decimal{}
	}

	__obf_21394d176152e7ad = __obf_21394d176152e7ad.Mul(__obf_63b94ab9e37045ad)

	__obf_21394d176152e7ad, __obf_db664194ea89d0f0 = __obf_21394d176152e7ad.ExpTaylor(-__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 + int32(__obf_60fb543883293d44))
	if __obf_db664194ea89d0f0 != nil {
		return Decimal{}
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_b48de9a28bbf0a58 := __obf_21271cbee6f28e58.Mul(__obf_21394d176152e7ad)

	return __obf_b48de9a28bbf0a58
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
func (__obf_c4701b3bb28cd2ae Decimal) PowWithPrecision(__obf_5c7cf1148c3170fc Decimal, __obf_60fb543883293d44 int32) (Decimal, error) {
	__obf_f3c9b4b2356d7cc9 := __obf_c4701b3bb28cd2ae.Sign()
	__obf_6d860b7cd06941cc := __obf_5c7cf1148c3170fc.Sign()

	if __obf_f3c9b4b2356d7cc9 == 0 {
		if __obf_6d860b7cd06941cc == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_6d860b7cd06941cc == 1 {
			return Decimal{__obf_31e6998552abcf35, 0}, nil
		}
		if __obf_6d860b7cd06941cc == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_6d860b7cd06941cc == 0 {
		return Decimal{__obf_6dcccdf86673e16e, 0}, nil
	}

	// TODO: optimize extraction of fractional part
	__obf_144c40e8322c0b92 := Decimal{__obf_6dcccdf86673e16e, 0}
	__obf_1235c9530aecc29b, __obf_63b94ab9e37045ad := __obf_5c7cf1148c3170fc.QuoRem(__obf_144c40e8322c0b92, 0)

	if __obf_f3c9b4b2356d7cc9 == -1 && !__obf_63b94ab9e37045ad.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}

	__obf_21271cbee6f28e58, _ := __obf_c4701b3bb28cd2ae.__obf_8ba1660137a7d859(__obf_1235c9530aecc29b.__obf_a1a825437ca7e6ad, __obf_60fb543883293d44)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_63b94ab9e37045ad.__obf_a1a825437ca7e6ad.Sign() == 0 {
		return __obf_21271cbee6f28e58, nil
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_df164c01bd219b41 := __obf_c4701b3bb28cd2ae.NumDigits()
	__obf_a6902edc3d438234 := __obf_5c7cf1148c3170fc.NumDigits()

	if int32(__obf_df164c01bd219b41) > __obf_60fb543883293d44 {
		__obf_60fb543883293d44 = int32(__obf_df164c01bd219b41)
	}
	if int32(__obf_a6902edc3d438234) > __obf_60fb543883293d44 {
		__obf_60fb543883293d44 += int32(__obf_a6902edc3d438234)
	}
	// increase precision by 10 to compensate for errors in further calculations
	__obf_60fb543883293d44 += 10

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_21394d176152e7ad, __obf_db664194ea89d0f0 := __obf_c4701b3bb28cd2ae.Abs().Ln(__obf_60fb543883293d44)
	if __obf_db664194ea89d0f0 != nil {
		return Decimal{}, __obf_db664194ea89d0f0
	}

	__obf_21394d176152e7ad = __obf_21394d176152e7ad.Mul(__obf_63b94ab9e37045ad)

	__obf_21394d176152e7ad, __obf_db664194ea89d0f0 = __obf_21394d176152e7ad.ExpTaylor(__obf_60fb543883293d44)
	if __obf_db664194ea89d0f0 != nil {
		return Decimal{}, __obf_db664194ea89d0f0
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_b48de9a28bbf0a58 := __obf_21271cbee6f28e58.Mul(__obf_21394d176152e7ad)

	return __obf_b48de9a28bbf0a58, nil
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
func (__obf_c4701b3bb28cd2ae Decimal) PowInt32(__obf_fed6fd50f0fcc834 int32) (Decimal, error) {
	if __obf_c4701b3bb28cd2ae.IsZero() && __obf_fed6fd50f0fcc834 == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_27dc795875d42527 := __obf_fed6fd50f0fcc834 < 0
	__obf_fed6fd50f0fcc834 = __obf_7b3894d563bc937f(__obf_fed6fd50f0fcc834)

	__obf_ee70f56e2047ecea, __obf_f742ee674d679d2d := __obf_c4701b3bb28cd2ae, New(1, 0)

	for __obf_fed6fd50f0fcc834 > 0 {
		if __obf_fed6fd50f0fcc834%2 == 1 {
			__obf_f742ee674d679d2d = __obf_f742ee674d679d2d.Mul(__obf_ee70f56e2047ecea)
		}
		__obf_fed6fd50f0fcc834 /= 2

		if __obf_fed6fd50f0fcc834 > 0 {
			__obf_ee70f56e2047ecea = __obf_ee70f56e2047ecea.Mul(__obf_ee70f56e2047ecea)
		}
	}

	if __obf_27dc795875d42527 {
		return New(1, 0).DivRound(__obf_f742ee674d679d2d, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_f742ee674d679d2d, nil
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
func (__obf_c4701b3bb28cd2ae Decimal) PowBigInt(__obf_fed6fd50f0fcc834 *big.Int) (Decimal, error) {
	return __obf_c4701b3bb28cd2ae.__obf_8ba1660137a7d859(__obf_fed6fd50f0fcc834, int32(PowPrecisionNegativeExponent))
}

func (__obf_c4701b3bb28cd2ae Decimal) __obf_8ba1660137a7d859(__obf_fed6fd50f0fcc834 *big.Int, __obf_60fb543883293d44 int32) (Decimal, error) {
	if __obf_c4701b3bb28cd2ae.IsZero() && __obf_fed6fd50f0fcc834.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_8468bb0b58969db1 := new(big.Int).Set(__obf_fed6fd50f0fcc834)
	__obf_27dc795875d42527 := __obf_fed6fd50f0fcc834.Sign() < 0

	if __obf_27dc795875d42527 {
		__obf_8468bb0b58969db1.Abs(__obf_8468bb0b58969db1)
	}

	__obf_ee70f56e2047ecea, __obf_f742ee674d679d2d := __obf_c4701b3bb28cd2ae, New(1, 0)

	for __obf_8468bb0b58969db1.Sign() > 0 {
		if __obf_8468bb0b58969db1.Bit(0) == 1 {
			__obf_f742ee674d679d2d = __obf_f742ee674d679d2d.Mul(__obf_ee70f56e2047ecea)
		}
		__obf_8468bb0b58969db1.Rsh(__obf_8468bb0b58969db1, 1)

		if __obf_8468bb0b58969db1.Sign() > 0 {
			__obf_ee70f56e2047ecea = __obf_ee70f56e2047ecea.Mul(__obf_ee70f56e2047ecea)
		}
	}

	if __obf_27dc795875d42527 {
		return New(1, 0).DivRound(__obf_f742ee674d679d2d, __obf_60fb543883293d44), nil
	}

	return __obf_f742ee674d679d2d, nil
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
func (__obf_c4701b3bb28cd2ae Decimal) ExpHullAbrham(__obf_6080ae6c2714fcbc uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_c4701b3bb28cd2ae.IsZero() {
		return Decimal{__obf_6dcccdf86673e16e, 0}, nil
	}

	__obf_1e5f88dd5d359662 := __obf_6080ae6c2714fcbc

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_966341667278a6f8 := __obf_c4701b3bb28cd2ae.Abs().InexactFloat64()
	if __obf_837145ee82c7c44b := __obf_966341667278a6f8 / 23; __obf_837145ee82c7c44b > float64(__obf_1e5f88dd5d359662) && __obf_837145ee82c7c44b < float64(ExpMaxIterations) {
		__obf_1e5f88dd5d359662 = uint32(math.Ceil(__obf_837145ee82c7c44b))
	}

	// fail if abs(d) beyond an over/underflow threshold
	__obf_56c91bd29ae1750f := New(23*int64(__obf_1e5f88dd5d359662), 0)
	if __obf_c4701b3bb28cd2ae.Abs().Cmp(__obf_56c91bd29ae1750f) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}

	// Return 1 if abs(d) small enough; this also avoids later over/underflow
	__obf_7b5dafdfb3a91b86 := New(9, -int32(__obf_1e5f88dd5d359662)-1)
	if __obf_c4701b3bb28cd2ae.Abs().Cmp(__obf_7b5dafdfb3a91b86) <= 0 {
		return Decimal{__obf_6dcccdf86673e16e, __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834}, nil
	}

	// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
	__obf_92443747cf082639 := __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 + int32(__obf_c4701b3bb28cd2ae.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_92443747cf082639 < 0 {
		__obf_92443747cf082639 = 0
	}

	__obf_5e8ea31acf2647c9 := New(1, __obf_92443747cf082639)                                                                                                                   // reduction factor
	__obf_e1eef1cd0f9a218e := Decimal{new(big.Int).Set(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad), __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 - __obf_92443747cf082639} // reduced argument
	__obf_90649e37a9bea2dd := int32(__obf_1e5f88dd5d359662) + __obf_92443747cf082639 + 2                                                                                       // precision for calculating the sum

	// Determine n, the number of therms for calculating sum
	// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
	// for solving appropriate equation, along with directed
	// roundings and simple rational bound for log10(p/abs(r))
	__obf_dc2a45d118fdff30 := __obf_e1eef1cd0f9a218e.Abs().InexactFloat64()
	__obf_b1210e9636d46151 := float64(__obf_90649e37a9bea2dd)
	__obf_c8bfed5d7d27e9ca := math.Ceil((1.453*__obf_b1210e9636d46151 - 1.182) / math.Log10(__obf_b1210e9636d46151/__obf_dc2a45d118fdff30))
	if __obf_c8bfed5d7d27e9ca > float64(ExpMaxIterations) || math.IsNaN(__obf_c8bfed5d7d27e9ca) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_ee70f56e2047ecea := int64(__obf_c8bfed5d7d27e9ca)

	__obf_fd8e5bcfddaceac6 := New(0, 0)
	__obf_512e9f06910cca45 := New(1, 0)
	__obf_144c40e8322c0b92 := New(1, 0)
	for __obf_159ec70197834006 := __obf_ee70f56e2047ecea - 1; __obf_159ec70197834006 > 0; __obf_159ec70197834006-- {
		__obf_fd8e5bcfddaceac6.__obf_a1a825437ca7e6ad.SetInt64(__obf_159ec70197834006)
		__obf_512e9f06910cca45 = __obf_512e9f06910cca45.Mul(__obf_e1eef1cd0f9a218e.DivRound(__obf_fd8e5bcfddaceac6, __obf_90649e37a9bea2dd))
		__obf_512e9f06910cca45 = __obf_512e9f06910cca45.Add(__obf_144c40e8322c0b92)
	}

	__obf_937992cda17e915a := __obf_5e8ea31acf2647c9.IntPart()
	__obf_b48de9a28bbf0a58 := New(1, 0)
	for __obf_159ec70197834006 := __obf_937992cda17e915a; __obf_159ec70197834006 > 0; __obf_159ec70197834006-- {
		__obf_b48de9a28bbf0a58 = __obf_b48de9a28bbf0a58.Mul(__obf_512e9f06910cca45)
	}

	__obf_fcd23926fb05722e := int32(__obf_b48de9a28bbf0a58.NumDigits())

	var __obf_26eb8a6d3017fde8 int32
	if __obf_fcd23926fb05722e > __obf_7b3894d563bc937f(__obf_b48de9a28bbf0a58.__obf_fed6fd50f0fcc834) {
		__obf_26eb8a6d3017fde8 = int32(__obf_1e5f88dd5d359662) - __obf_fcd23926fb05722e - __obf_b48de9a28bbf0a58.__obf_fed6fd50f0fcc834
	} else {
		__obf_26eb8a6d3017fde8 = int32(__obf_1e5f88dd5d359662)
	}

	__obf_b48de9a28bbf0a58 = __obf_b48de9a28bbf0a58.Round(__obf_26eb8a6d3017fde8)

	return __obf_b48de9a28bbf0a58, nil
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
func (__obf_c4701b3bb28cd2ae Decimal) ExpTaylor(__obf_60fb543883293d44 int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_c4701b3bb28cd2ae.IsZero() {
		return Decimal{__obf_6dcccdf86673e16e, 0}.Round(__obf_60fb543883293d44), nil
	}

	var __obf_0551e7396b41e252 Decimal
	var __obf_cd0d477d7123bfcb int32
	if __obf_60fb543883293d44 < 0 {
		__obf_0551e7396b41e252 = New(1, -1)
		__obf_cd0d477d7123bfcb = 8
	} else {
		__obf_0551e7396b41e252 = New(1, -__obf_60fb543883293d44-1)
		__obf_cd0d477d7123bfcb = __obf_60fb543883293d44 + 1
	}

	__obf_1157353202e58f35 := __obf_c4701b3bb28cd2ae.Abs()
	__obf_7a53500472d4614c := __obf_c4701b3bb28cd2ae.Abs()
	__obf_ee7fbc840509611b := New(1, 0)

	__obf_f742ee674d679d2d := New(1, 0)

	for __obf_159ec70197834006 := int64(1); ; {
		__obf_6bc4bac2c49c90db := __obf_7a53500472d4614c.DivRound(__obf_ee7fbc840509611b, __obf_cd0d477d7123bfcb)
		__obf_f742ee674d679d2d = __obf_f742ee674d679d2d.Add(__obf_6bc4bac2c49c90db)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_6bc4bac2c49c90db.Cmp(__obf_0551e7396b41e252) < 0 {
			break
		}

		__obf_7a53500472d4614c = __obf_7a53500472d4614c.Mul(__obf_1157353202e58f35)

		__obf_159ec70197834006++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_f94a2e25be307681) >= int(__obf_159ec70197834006) && !__obf_f94a2e25be307681[__obf_159ec70197834006-1].IsZero() {
			__obf_ee7fbc840509611b = __obf_f94a2e25be307681[__obf_159ec70197834006-1]
		} else {
			// To avoid any race conditions, firstly the zero value is appended to a slice to create
			// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
			// factorial using the index notation.
			__obf_ee7fbc840509611b = __obf_f94a2e25be307681[__obf_159ec70197834006-2].Mul(New(__obf_159ec70197834006, 0))
			__obf_f94a2e25be307681 = append(__obf_f94a2e25be307681, Zero)
			__obf_f94a2e25be307681[__obf_159ec70197834006-1] = __obf_ee7fbc840509611b
		}
	}

	if __obf_c4701b3bb28cd2ae.Sign() < 0 {
		__obf_f742ee674d679d2d = New(1, 0).DivRound(__obf_f742ee674d679d2d, __obf_60fb543883293d44+1)
	}

	__obf_f742ee674d679d2d = __obf_f742ee674d679d2d.Round(__obf_60fb543883293d44)
	return __obf_f742ee674d679d2d, nil
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
func (__obf_c4701b3bb28cd2ae Decimal) Ln(__obf_60fb543883293d44 int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_c4701b3bb28cd2ae.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_c4701b3bb28cd2ae.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}

	__obf_afb41485c15984cd := __obf_60fb543883293d44 + 2
	__obf_a3d7b54ee2bf67c1 := __obf_c4701b3bb28cd2ae.Copy()

	var __obf_a94ba4bdf01b795d, __obf_bc048d9520754d9f, __obf_4e4f54a4f9b91f43, __obf_d018f50ec02e76d6, __obf_8fa7bfc983285841 Decimal
	__obf_a94ba4bdf01b795d = __obf_a3d7b54ee2bf67c1.Sub(Decimal{__obf_6dcccdf86673e16e, 0})
	__obf_bc048d9520754d9f = Decimal{__obf_6dcccdf86673e16e, -1}

	// for decimal in range [0.9, 1.1] where ln(d) is close to 0
	__obf_38079d1a31326a8e := false

	if __obf_a94ba4bdf01b795d.Abs().Cmp(__obf_bc048d9520754d9f) <= 0 {
		__obf_38079d1a31326a8e = true
	} else {
		// reduce input decimal to range [0.1, 1)
		__obf_38d082d12979b586 := int32(__obf_a3d7b54ee2bf67c1.NumDigits()) + __obf_a3d7b54ee2bf67c1.__obf_fed6fd50f0fcc834
		__obf_a3d7b54ee2bf67c1.__obf_fed6fd50f0fcc834 -= __obf_38d082d12979b586

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_cc8f6d2f40466d6c := __obf_cc8f6d2f40466d6c.__obf_88c6d523c9ba6f93(__obf_afb41485c15984cd)
		__obf_8fa7bfc983285841 = NewFromInt32(__obf_38d082d12979b586)
		__obf_8fa7bfc983285841 = __obf_8fa7bfc983285841.Mul(__obf_cc8f6d2f40466d6c)

		__obf_a94ba4bdf01b795d = __obf_a3d7b54ee2bf67c1.Sub(Decimal{__obf_6dcccdf86673e16e, 0})

		if __obf_a94ba4bdf01b795d.Abs().Cmp(__obf_bc048d9520754d9f) <= 0 {
			__obf_38079d1a31326a8e = true
		} else {
			// initial estimate using floats
			__obf_c1a3c5cb55dd1921 := __obf_a3d7b54ee2bf67c1.InexactFloat64()
			__obf_a94ba4bdf01b795d = NewFromFloat(math.Log(__obf_c1a3c5cb55dd1921))
		}
	}

	__obf_0551e7396b41e252 := Decimal{__obf_6dcccdf86673e16e, -__obf_afb41485c15984cd}

	if __obf_38079d1a31326a8e {
		// Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
		// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
		// until the difference between current and next term is smaller than epsilon.
		// Coverage quite fast for decimals close to 1.0

		// z + 2
		__obf_4e4f54a4f9b91f43 = __obf_a94ba4bdf01b795d.Add(Decimal{__obf_97304e664e9a31ae, 0})
		// z / (z + 2)
		__obf_bc048d9520754d9f = __obf_a94ba4bdf01b795d.DivRound(__obf_4e4f54a4f9b91f43, __obf_afb41485c15984cd)
		// 2 * (z / (z + 2))
		__obf_a94ba4bdf01b795d = __obf_bc048d9520754d9f.Add(__obf_bc048d9520754d9f)
		__obf_4e4f54a4f9b91f43 = __obf_a94ba4bdf01b795d.Copy()

		for __obf_ee70f56e2047ecea := 1; ; __obf_ee70f56e2047ecea++ {
			// 2 * (z / (z+2))^(2n+1)
			__obf_4e4f54a4f9b91f43 = __obf_4e4f54a4f9b91f43.Mul(__obf_bc048d9520754d9f).Mul(__obf_bc048d9520754d9f)

			// 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
			__obf_d018f50ec02e76d6 = NewFromInt(int64(2*__obf_ee70f56e2047ecea + 1))
			__obf_d018f50ec02e76d6 = __obf_4e4f54a4f9b91f43.DivRound(__obf_d018f50ec02e76d6, __obf_afb41485c15984cd)

			// comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			__obf_a94ba4bdf01b795d = __obf_a94ba4bdf01b795d.Add(__obf_d018f50ec02e76d6)

			if __obf_d018f50ec02e76d6.Abs().Cmp(__obf_0551e7396b41e252) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_bcc05041e2332651 Decimal
		__obf_54cdec8177befb58 := __obf_afb41485c15984cd*2 + 10

		for __obf_159ec70197834006 := int32(0); __obf_159ec70197834006 < __obf_54cdec8177befb58; __obf_159ec70197834006++ {
			// exp(a_n)
			__obf_bc048d9520754d9f, _ = __obf_a94ba4bdf01b795d.ExpTaylor(__obf_afb41485c15984cd)
			// exp(a_n) - z
			__obf_4e4f54a4f9b91f43 = __obf_bc048d9520754d9f.Sub(__obf_a3d7b54ee2bf67c1)
			// 2 * (exp(a_n) - z)
			__obf_4e4f54a4f9b91f43 = __obf_4e4f54a4f9b91f43.Add(__obf_4e4f54a4f9b91f43)
			// exp(a_n) + z
			__obf_d018f50ec02e76d6 = __obf_bc048d9520754d9f.Add(__obf_a3d7b54ee2bf67c1)
			// 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_bc048d9520754d9f = __obf_4e4f54a4f9b91f43.DivRound(__obf_d018f50ec02e76d6, __obf_afb41485c15984cd)
			// comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_a94ba4bdf01b795d = __obf_a94ba4bdf01b795d.Sub(__obf_bc048d9520754d9f)

			if __obf_bcc05041e2332651.Add(__obf_bc048d9520754d9f).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_bc048d9520754d9f.Abs().Cmp(__obf_0551e7396b41e252) <= 0 {
				break
			}

			__obf_bcc05041e2332651 = __obf_bc048d9520754d9f
		}
	}

	__obf_a94ba4bdf01b795d = __obf_a94ba4bdf01b795d.Add(__obf_8fa7bfc983285841)

	return __obf_a94ba4bdf01b795d.Round(__obf_60fb543883293d44), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_c4701b3bb28cd2ae Decimal) NumDigits() int {
	if __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad == nil {
		return 1
	}

	if __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.IsInt64() {
		__obf_d7556b2f5ea0bbd6 := __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.Int64()
		// restrict fast path to integers with exact conversion to float64
		if __obf_d7556b2f5ea0bbd6 <= (1<<53) && __obf_d7556b2f5ea0bbd6 >= -(1<<53) {
			if __obf_d7556b2f5ea0bbd6 == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_d7556b2f5ea0bbd6)))) + 1
		}
	}

	__obf_b75af5a1cb1c5946 := int(float64(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.BitLen()) / math.Log2(10))

	// estimatedNumDigits (lg10) may be off by 1, need to verify
	__obf_264796df0476859b := big.NewInt(int64(__obf_b75af5a1cb1c5946))
	__obf_dbe14b52cdb2a8e3 := __obf_264796df0476859b.Exp(__obf_a9bf241bb08f75b7, __obf_264796df0476859b, nil)

	if __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.CmpAbs(__obf_dbe14b52cdb2a8e3) >= 0 {
		return __obf_b75af5a1cb1c5946 + 1
	}

	return __obf_b75af5a1cb1c5946
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_c4701b3bb28cd2ae Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_e1eef1cd0f9a218e big.Int
	__obf_9d9315f9ed27915a := new(big.Int).Set(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad)
	for __obf_a3d7b54ee2bf67c1 := __obf_7b3894d563bc937f(__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834); __obf_a3d7b54ee2bf67c1 > 0; __obf_a3d7b54ee2bf67c1-- {
		__obf_9d9315f9ed27915a.QuoRem(__obf_9d9315f9ed27915a, __obf_a9bf241bb08f75b7, &__obf_e1eef1cd0f9a218e)
		if __obf_e1eef1cd0f9a218e.Cmp(__obf_31e6998552abcf35) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_7b3894d563bc937f(__obf_ee70f56e2047ecea int32) int32 {
	if __obf_ee70f56e2047ecea < 0 {
		return -__obf_ee70f56e2047ecea
	}
	return __obf_ee70f56e2047ecea
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_c4701b3bb28cd2ae Decimal) Cmp(__obf_5c7cf1148c3170fc Decimal) int {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()
	__obf_5c7cf1148c3170fc.__obf_4b849233a95eee86()

	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 == __obf_5c7cf1148c3170fc.__obf_fed6fd50f0fcc834 {
		return __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.Cmp(__obf_5c7cf1148c3170fc.__obf_a1a825437ca7e6ad)
	}

	__obf_cd15ee72b8cc5c32, __obf_ed4756049744a46c := RescalePair(__obf_c4701b3bb28cd2ae, __obf_5c7cf1148c3170fc)

	return __obf_cd15ee72b8cc5c32.__obf_a1a825437ca7e6ad.Cmp(__obf_ed4756049744a46c.__obf_a1a825437ca7e6ad)
}

// Compare compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_c4701b3bb28cd2ae Decimal) Compare(__obf_5c7cf1148c3170fc Decimal) int {
	return __obf_c4701b3bb28cd2ae.Cmp(__obf_5c7cf1148c3170fc)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_c4701b3bb28cd2ae Decimal) Equal(__obf_5c7cf1148c3170fc Decimal) bool {
	return __obf_c4701b3bb28cd2ae.Cmp(__obf_5c7cf1148c3170fc) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_c4701b3bb28cd2ae Decimal) Equals(__obf_5c7cf1148c3170fc Decimal) bool {
	return __obf_c4701b3bb28cd2ae.Equal(__obf_5c7cf1148c3170fc)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_c4701b3bb28cd2ae Decimal) GreaterThan(__obf_5c7cf1148c3170fc Decimal) bool {
	return __obf_c4701b3bb28cd2ae.Cmp(__obf_5c7cf1148c3170fc) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_c4701b3bb28cd2ae Decimal) GreaterThanOrEqual(__obf_5c7cf1148c3170fc Decimal) bool {
	__obf_24ca3a0e936311c7 := __obf_c4701b3bb28cd2ae.Cmp(__obf_5c7cf1148c3170fc)
	return __obf_24ca3a0e936311c7 == 1 || __obf_24ca3a0e936311c7 == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_c4701b3bb28cd2ae Decimal) LessThan(__obf_5c7cf1148c3170fc Decimal) bool {
	return __obf_c4701b3bb28cd2ae.Cmp(__obf_5c7cf1148c3170fc) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_c4701b3bb28cd2ae Decimal) LessThanOrEqual(__obf_5c7cf1148c3170fc Decimal) bool {
	__obf_24ca3a0e936311c7 := __obf_c4701b3bb28cd2ae.Cmp(__obf_5c7cf1148c3170fc)
	return __obf_24ca3a0e936311c7 == -1 || __obf_24ca3a0e936311c7 == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_c4701b3bb28cd2ae Decimal) Sign() int {
	if __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad == nil {
		return 0
	}
	return __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.Sign()
}

// IsPositive return
//
//	true if d > 0
//	false if d == 0
//	false if d < 0
func (__obf_c4701b3bb28cd2ae Decimal) IsPositive() bool {
	return __obf_c4701b3bb28cd2ae.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_c4701b3bb28cd2ae Decimal) IsNegative() bool {
	return __obf_c4701b3bb28cd2ae.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_c4701b3bb28cd2ae Decimal) IsZero() bool {
	return __obf_c4701b3bb28cd2ae.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_c4701b3bb28cd2ae Decimal) Exponent() int32 {
	return __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834
}

// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
func (__obf_c4701b3bb28cd2ae Decimal) Coefficient() *big.Int {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad)
}

// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
// If coefficient cannot be represented in an int64, the result will be undefined.
func (__obf_c4701b3bb28cd2ae Decimal) CoefficientInt64() int64 {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()
	return __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.Int64()
}

// IntPart returns the integer component of the decimal.
func (__obf_c4701b3bb28cd2ae Decimal) IntPart() int64 {
	__obf_eda1d501b5799992 := __obf_c4701b3bb28cd2ae.__obf_452282ed04deedc9(0)
	return __obf_eda1d501b5799992.__obf_a1a825437ca7e6ad.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_c4701b3bb28cd2ae Decimal) BigInt() *big.Int {
	__obf_eda1d501b5799992 := __obf_c4701b3bb28cd2ae.__obf_452282ed04deedc9(0)
	return __obf_eda1d501b5799992.__obf_a1a825437ca7e6ad
}

// BigFloat returns decimal as BigFloat.
// Be aware that casting decimal to BigFloat might cause a loss of precision.
func (__obf_c4701b3bb28cd2ae Decimal) BigFloat() *big.Float {
	__obf_966341667278a6f8 := &big.Float{}
	__obf_966341667278a6f8.SetString(__obf_c4701b3bb28cd2ae.String())
	return __obf_966341667278a6f8
}

// Rat returns a rational number representation of the decimal.
func (__obf_c4701b3bb28cd2ae Decimal) Rat() *big.Rat {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()
	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 <= 0 {
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_d0e346cf8df7568c := new(big.Int).Exp(__obf_a9bf241bb08f75b7, big.NewInt(-int64(__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834)), nil)
		return new(big.Rat).SetFrac(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad, __obf_d0e346cf8df7568c)
	}

	__obf_ae5624a77ad217ca := new(big.Int).Exp(__obf_a9bf241bb08f75b7, big.NewInt(int64(__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834)), nil)
	__obf_ed072fef3ae941a7 := new(big.Int).Mul(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad, __obf_ae5624a77ad217ca)
	return new(big.Rat).SetFrac(__obf_ed072fef3ae941a7, __obf_6dcccdf86673e16e)
}

// Float64 returns the nearest float64 value for d and a bool indicating
// whether f represents d exactly.
// For more details, see the documentation for big.Rat.Float64
func (__obf_c4701b3bb28cd2ae Decimal) Float64() (__obf_966341667278a6f8 float64, __obf_a418cf7d98486aeb bool) {
	return __obf_c4701b3bb28cd2ae.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_c4701b3bb28cd2ae Decimal) InexactFloat64() float64 {
	__obf_966341667278a6f8, _ := __obf_c4701b3bb28cd2ae.Float64()
	return __obf_966341667278a6f8
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
func (__obf_c4701b3bb28cd2ae Decimal) String() string {
	return __obf_c4701b3bb28cd2ae.string(true)
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
func (__obf_c4701b3bb28cd2ae Decimal) StringFixed(__obf_cc5404f88394000e int32) string {
	__obf_584c44704b53e8ef := __obf_c4701b3bb28cd2ae.Round(__obf_cc5404f88394000e)
	return __obf_584c44704b53e8ef.string(false)
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
func (__obf_c4701b3bb28cd2ae Decimal) StringFixedBank(__obf_cc5404f88394000e int32) string {
	__obf_584c44704b53e8ef := __obf_c4701b3bb28cd2ae.RoundBank(__obf_cc5404f88394000e)
	return __obf_584c44704b53e8ef.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_c4701b3bb28cd2ae Decimal) StringFixedCash(__obf_56d9063089782e43 uint8) string {
	__obf_584c44704b53e8ef := __obf_c4701b3bb28cd2ae.RoundCash(__obf_56d9063089782e43)
	return __obf_584c44704b53e8ef.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_c4701b3bb28cd2ae Decimal) Round(__obf_cc5404f88394000e int32) Decimal {
	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 == -__obf_cc5404f88394000e {
		return __obf_c4701b3bb28cd2ae
	}
	// truncate to places + 1
	__obf_b8acb06d607160a4 := __obf_c4701b3bb28cd2ae.__obf_452282ed04deedc9(-__obf_cc5404f88394000e - 1)

	// add sign(d) * 0.5
	if __obf_b8acb06d607160a4.__obf_a1a825437ca7e6ad.Sign() < 0 {
		__obf_b8acb06d607160a4.__obf_a1a825437ca7e6ad.Sub(__obf_b8acb06d607160a4.__obf_a1a825437ca7e6ad, __obf_84ba5bd50a08ec11)
	} else {
		__obf_b8acb06d607160a4.__obf_a1a825437ca7e6ad.Add(__obf_b8acb06d607160a4.__obf_a1a825437ca7e6ad, __obf_84ba5bd50a08ec11)
	}

	// floor for positive numbers, ceil for negative numbers
	_, __obf_4739b75ffb6404bf := __obf_b8acb06d607160a4.__obf_a1a825437ca7e6ad.DivMod(__obf_b8acb06d607160a4.__obf_a1a825437ca7e6ad, __obf_a9bf241bb08f75b7, new(big.Int))
	__obf_b8acb06d607160a4.__obf_fed6fd50f0fcc834++
	if __obf_b8acb06d607160a4.__obf_a1a825437ca7e6ad.Sign() < 0 && __obf_4739b75ffb6404bf.Cmp(__obf_31e6998552abcf35) != 0 {
		__obf_b8acb06d607160a4.__obf_a1a825437ca7e6ad.Add(__obf_b8acb06d607160a4.__obf_a1a825437ca7e6ad, __obf_6dcccdf86673e16e)
	}

	return __obf_b8acb06d607160a4
}

// RoundCeil rounds the decimal towards +infinity.
//
// Example:
//
//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
func (__obf_c4701b3bb28cd2ae Decimal) RoundCeil(__obf_cc5404f88394000e int32) Decimal {
	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 >= -__obf_cc5404f88394000e {
		return __obf_c4701b3bb28cd2ae
	}

	__obf_5b9388b52376e3f5 := __obf_c4701b3bb28cd2ae.__obf_452282ed04deedc9(-__obf_cc5404f88394000e)
	if __obf_c4701b3bb28cd2ae.Equal(__obf_5b9388b52376e3f5) {
		return __obf_c4701b3bb28cd2ae
	}

	if __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.Sign() > 0 {
		__obf_5b9388b52376e3f5.__obf_a1a825437ca7e6ad.Add(__obf_5b9388b52376e3f5.__obf_a1a825437ca7e6ad, __obf_6dcccdf86673e16e)
	}

	return __obf_5b9388b52376e3f5
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_c4701b3bb28cd2ae Decimal) RoundFloor(__obf_cc5404f88394000e int32) Decimal {
	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 >= -__obf_cc5404f88394000e {
		return __obf_c4701b3bb28cd2ae
	}

	__obf_5b9388b52376e3f5 := __obf_c4701b3bb28cd2ae.__obf_452282ed04deedc9(-__obf_cc5404f88394000e)
	if __obf_c4701b3bb28cd2ae.Equal(__obf_5b9388b52376e3f5) {
		return __obf_c4701b3bb28cd2ae
	}

	if __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.Sign() < 0 {
		__obf_5b9388b52376e3f5.__obf_a1a825437ca7e6ad.Sub(__obf_5b9388b52376e3f5.__obf_a1a825437ca7e6ad, __obf_6dcccdf86673e16e)
	}

	return __obf_5b9388b52376e3f5
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_c4701b3bb28cd2ae Decimal) RoundUp(__obf_cc5404f88394000e int32) Decimal {
	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 >= -__obf_cc5404f88394000e {
		return __obf_c4701b3bb28cd2ae
	}

	__obf_5b9388b52376e3f5 := __obf_c4701b3bb28cd2ae.__obf_452282ed04deedc9(-__obf_cc5404f88394000e)
	if __obf_c4701b3bb28cd2ae.Equal(__obf_5b9388b52376e3f5) {
		return __obf_c4701b3bb28cd2ae
	}

	if __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.Sign() > 0 {
		__obf_5b9388b52376e3f5.__obf_a1a825437ca7e6ad.Add(__obf_5b9388b52376e3f5.__obf_a1a825437ca7e6ad, __obf_6dcccdf86673e16e)
	} else if __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.Sign() < 0 {
		__obf_5b9388b52376e3f5.__obf_a1a825437ca7e6ad.Sub(__obf_5b9388b52376e3f5.__obf_a1a825437ca7e6ad, __obf_6dcccdf86673e16e)
	}

	return __obf_5b9388b52376e3f5
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_c4701b3bb28cd2ae Decimal) RoundDown(__obf_cc5404f88394000e int32) Decimal {
	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 >= -__obf_cc5404f88394000e {
		return __obf_c4701b3bb28cd2ae
	}

	__obf_5b9388b52376e3f5 := __obf_c4701b3bb28cd2ae.__obf_452282ed04deedc9(-__obf_cc5404f88394000e)
	if __obf_c4701b3bb28cd2ae.Equal(__obf_5b9388b52376e3f5) {
		return __obf_c4701b3bb28cd2ae
	}
	return __obf_5b9388b52376e3f5
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
func (__obf_c4701b3bb28cd2ae Decimal) RoundBank(__obf_cc5404f88394000e int32) Decimal {

	__obf_97fc9da5a53d2b14 := __obf_c4701b3bb28cd2ae.Round(__obf_cc5404f88394000e)
	__obf_5da5e6b8f0ca5e73 := __obf_c4701b3bb28cd2ae.Sub(__obf_97fc9da5a53d2b14).Abs()

	__obf_8d072b2f58f3054d := New(5, -__obf_cc5404f88394000e-1)
	if __obf_5da5e6b8f0ca5e73.Cmp(__obf_8d072b2f58f3054d) == 0 && __obf_97fc9da5a53d2b14.__obf_a1a825437ca7e6ad.Bit(0) != 0 {
		if __obf_97fc9da5a53d2b14.__obf_a1a825437ca7e6ad.Sign() < 0 {
			__obf_97fc9da5a53d2b14.__obf_a1a825437ca7e6ad.Add(__obf_97fc9da5a53d2b14.__obf_a1a825437ca7e6ad, __obf_6dcccdf86673e16e)
		} else {
			__obf_97fc9da5a53d2b14.__obf_a1a825437ca7e6ad.Sub(__obf_97fc9da5a53d2b14.__obf_a1a825437ca7e6ad, __obf_6dcccdf86673e16e)
		}
	}

	return __obf_97fc9da5a53d2b14
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
func (__obf_c4701b3bb28cd2ae Decimal) RoundCash(__obf_56d9063089782e43 uint8) Decimal {
	var __obf_4728d963b1c41e39 *big.Int
	switch __obf_56d9063089782e43 {
	case 5:
		__obf_4728d963b1c41e39 = __obf_d3e2e1fd11685001
	case 10:
		__obf_4728d963b1c41e39 = __obf_a9bf241bb08f75b7
	case 25:
		__obf_4728d963b1c41e39 = __obf_c16b63cd3d3ece01
	case 50:
		__obf_4728d963b1c41e39 = __obf_97304e664e9a31ae
	case 100:
		__obf_4728d963b1c41e39 = __obf_6dcccdf86673e16e
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_56d9063089782e43))
	}
	__obf_5f0f2f9f58c5c3c0 := Decimal{
		__obf_a1a825437ca7e6ad: __obf_4728d963b1c41e39,
	}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_c4701b3bb28cd2ae.Mul(__obf_5f0f2f9f58c5c3c0).Round(0).Div(__obf_5f0f2f9f58c5c3c0).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_c4701b3bb28cd2ae Decimal) Floor() Decimal {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()

	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 >= 0 {
		return __obf_c4701b3bb28cd2ae
	}

	__obf_fed6fd50f0fcc834 := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_fed6fd50f0fcc834.Exp(__obf_fed6fd50f0fcc834, big.NewInt(-int64(__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834)), nil)

	__obf_a3d7b54ee2bf67c1 := new(big.Int).Div(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad, __obf_fed6fd50f0fcc834)
	return Decimal{__obf_a1a825437ca7e6ad: __obf_a3d7b54ee2bf67c1, __obf_fed6fd50f0fcc834: 0}
}

// Ceil returns the nearest integer value greater than or equal to d.
func (__obf_c4701b3bb28cd2ae Decimal) Ceil() Decimal {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()

	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 >= 0 {
		return __obf_c4701b3bb28cd2ae
	}

	__obf_fed6fd50f0fcc834 := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_fed6fd50f0fcc834.Exp(__obf_fed6fd50f0fcc834, big.NewInt(-int64(__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834)), nil)

	__obf_a3d7b54ee2bf67c1, __obf_4739b75ffb6404bf := new(big.Int).DivMod(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad, __obf_fed6fd50f0fcc834, new(big.Int))
	if __obf_4739b75ffb6404bf.Cmp(__obf_31e6998552abcf35) != 0 {
		__obf_a3d7b54ee2bf67c1.Add(__obf_a3d7b54ee2bf67c1, __obf_6dcccdf86673e16e)
	}
	return Decimal{__obf_a1a825437ca7e6ad: __obf_a3d7b54ee2bf67c1, __obf_fed6fd50f0fcc834: 0}
}

// Truncate truncates off digits from the number, without rounding.
//
// NOTE: precision is the last digit that will not be truncated (must be >= 0).
//
// Example:
//
//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
func (__obf_c4701b3bb28cd2ae Decimal) Truncate(__obf_60fb543883293d44 int32) Decimal {
	__obf_c4701b3bb28cd2ae.__obf_4b849233a95eee86()
	if __obf_60fb543883293d44 >= 0 && -__obf_60fb543883293d44 > __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 {
		return __obf_c4701b3bb28cd2ae.__obf_452282ed04deedc9(-__obf_60fb543883293d44)
	}
	return __obf_c4701b3bb28cd2ae
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_c4701b3bb28cd2ae *Decimal) UnmarshalJSON(__obf_9571474f72c71159 []byte) error {
	if string(__obf_9571474f72c71159) == "null" {
		return nil
	}

	__obf_4f9552183401314b, __obf_db664194ea89d0f0 := __obf_422c9f22dd4dfc68(__obf_9571474f72c71159)
	if __obf_db664194ea89d0f0 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_9571474f72c71159, __obf_db664194ea89d0f0)
	}

	__obf_adf9c80aee40a9d8, __obf_db664194ea89d0f0 := NewFromString(__obf_4f9552183401314b)
	*__obf_c4701b3bb28cd2ae = __obf_adf9c80aee40a9d8
	if __obf_db664194ea89d0f0 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_4f9552183401314b, __obf_db664194ea89d0f0)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_c4701b3bb28cd2ae Decimal) MarshalJSON() ([]byte, error) {
	var __obf_4f9552183401314b string
	if MarshalJSONWithoutQuotes {
		__obf_4f9552183401314b = __obf_c4701b3bb28cd2ae.String()
	} else {
		__obf_4f9552183401314b = "\"" + __obf_c4701b3bb28cd2ae.String() + "\""
	}
	return []byte(__obf_4f9552183401314b), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_c4701b3bb28cd2ae *Decimal) UnmarshalBinary(__obf_eb399de1be669279 []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_eb399de1be669279) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_eb399de1be669279, len(__obf_eb399de1be669279))
	}

	// Extract the exponent
	__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 = int32(binary.BigEndian.Uint32(__obf_eb399de1be669279[:4]))

	// Extract the value
	__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad = new(big.Int)
	if __obf_db664194ea89d0f0 := __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.GobDecode(__obf_eb399de1be669279[4:]); __obf_db664194ea89d0f0 != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_eb399de1be669279, __obf_db664194ea89d0f0)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_c4701b3bb28cd2ae Decimal) MarshalBinary() (__obf_eb399de1be669279 []byte, __obf_db664194ea89d0f0 error) {
	// exp is written first, but encode value first to know output size
	var __obf_519cbef7abfe9e2b []byte
	if __obf_519cbef7abfe9e2b, __obf_db664194ea89d0f0 = __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.GobEncode(); __obf_db664194ea89d0f0 != nil {
		return nil, __obf_db664194ea89d0f0
	}

	// Write the exponent in front, since it's a fixed size
	__obf_f1a454baab9c3462 := make([]byte, 4, len(__obf_519cbef7abfe9e2b)+4)
	binary.BigEndian.PutUint32(__obf_f1a454baab9c3462, uint32(__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834))

	// Return the byte array
	return append(__obf_f1a454baab9c3462, __obf_519cbef7abfe9e2b...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_c4701b3bb28cd2ae *Decimal) Scan(__obf_a1a825437ca7e6ad any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_d056d9dcafa97698 := __obf_a1a825437ca7e6ad.(type) {

	case float32:
		*__obf_c4701b3bb28cd2ae = NewFromFloat(float64(__obf_d056d9dcafa97698))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_c4701b3bb28cd2ae = NewFromFloat(__obf_d056d9dcafa97698)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_c4701b3bb28cd2ae = New(__obf_d056d9dcafa97698, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_c4701b3bb28cd2ae = NewFromUint64(__obf_d056d9dcafa97698)
		return nil

	default:
		// default is trying to interpret value stored as string
		__obf_4f9552183401314b, __obf_db664194ea89d0f0 := __obf_422c9f22dd4dfc68(__obf_d056d9dcafa97698)
		if __obf_db664194ea89d0f0 != nil {
			return __obf_db664194ea89d0f0
		}
		*__obf_c4701b3bb28cd2ae, __obf_db664194ea89d0f0 = NewFromString(__obf_4f9552183401314b)
		return __obf_db664194ea89d0f0
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_c4701b3bb28cd2ae Decimal) Value() (driver.Value, error) {
	return __obf_c4701b3bb28cd2ae.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_c4701b3bb28cd2ae *Decimal) UnmarshalText(__obf_13dc8aff84147625 []byte) error {
	__obf_4f9552183401314b := string(__obf_13dc8aff84147625)

	__obf_aaf3c5b40a6a1b7f, __obf_db664194ea89d0f0 := NewFromString(__obf_4f9552183401314b)
	*__obf_c4701b3bb28cd2ae = __obf_aaf3c5b40a6a1b7f
	if __obf_db664194ea89d0f0 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_4f9552183401314b, __obf_db664194ea89d0f0)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_c4701b3bb28cd2ae Decimal) MarshalText() (__obf_13dc8aff84147625 []byte, __obf_db664194ea89d0f0 error) {
	return []byte(__obf_c4701b3bb28cd2ae.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_c4701b3bb28cd2ae Decimal) GobEncode() ([]byte, error) {
	return __obf_c4701b3bb28cd2ae.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_c4701b3bb28cd2ae *Decimal) GobDecode(__obf_eb399de1be669279 []byte) error {
	return __obf_c4701b3bb28cd2ae.UnmarshalBinary(__obf_eb399de1be669279)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_c4701b3bb28cd2ae Decimal) StringScaled(__obf_fed6fd50f0fcc834 int32) string {
	return __obf_c4701b3bb28cd2ae.__obf_452282ed04deedc9(__obf_fed6fd50f0fcc834).String()
}

func (__obf_c4701b3bb28cd2ae Decimal) string(__obf_3d88ac59294a2641 bool) string {
	if __obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834 >= 0 {
		return __obf_c4701b3bb28cd2ae.__obf_452282ed04deedc9(0).__obf_a1a825437ca7e6ad.String()
	}

	__obf_7b3894d563bc937f := new(big.Int).Abs(__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad)
	__obf_4f9552183401314b := __obf_7b3894d563bc937f.String()

	var __obf_bcb73fc952fc18db, __obf_13f8a64cb9beba59 string

	// NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
	// and you are on a 32-bit machine. Won't fix this super-edge case.
	__obf_67044e9b1899efec := int(__obf_c4701b3bb28cd2ae.__obf_fed6fd50f0fcc834)
	if len(__obf_4f9552183401314b) > -__obf_67044e9b1899efec {
		__obf_bcb73fc952fc18db = __obf_4f9552183401314b[:len(__obf_4f9552183401314b)+__obf_67044e9b1899efec]
		__obf_13f8a64cb9beba59 = __obf_4f9552183401314b[len(__obf_4f9552183401314b)+__obf_67044e9b1899efec:]
	} else {
		__obf_bcb73fc952fc18db = "0"

		__obf_beb79442d0368778 := -__obf_67044e9b1899efec - len(__obf_4f9552183401314b)
		__obf_13f8a64cb9beba59 = strings.Repeat("0", __obf_beb79442d0368778) + __obf_4f9552183401314b
	}

	if __obf_3d88ac59294a2641 {
		__obf_159ec70197834006 := len(__obf_13f8a64cb9beba59) - 1
		for ; __obf_159ec70197834006 >= 0; __obf_159ec70197834006-- {
			if __obf_13f8a64cb9beba59[__obf_159ec70197834006] != '0' {
				break
			}
		}
		__obf_13f8a64cb9beba59 = __obf_13f8a64cb9beba59[:__obf_159ec70197834006+1]
	}

	__obf_2a0dd5426e9907ea := __obf_bcb73fc952fc18db
	if len(__obf_13f8a64cb9beba59) > 0 {
		__obf_2a0dd5426e9907ea += "." + __obf_13f8a64cb9beba59
	}

	if __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad.Sign() < 0 {
		return "-" + __obf_2a0dd5426e9907ea
	}

	return __obf_2a0dd5426e9907ea
}

func (__obf_c4701b3bb28cd2ae *Decimal) __obf_4b849233a95eee86() {
	if __obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad == nil {
		__obf_c4701b3bb28cd2ae.__obf_a1a825437ca7e6ad = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_8c5e71978702f5ff Decimal, __obf_496337d4762bc775 ...Decimal) Decimal {
	__obf_edfb45ae797263ef := __obf_8c5e71978702f5ff
	for _, __obf_817531fc9afb059d := range __obf_496337d4762bc775 {
		if __obf_817531fc9afb059d.Cmp(__obf_edfb45ae797263ef) < 0 {
			__obf_edfb45ae797263ef = __obf_817531fc9afb059d
		}
	}
	return __obf_edfb45ae797263ef
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_8c5e71978702f5ff Decimal, __obf_496337d4762bc775 ...Decimal) Decimal {
	__obf_edfb45ae797263ef := __obf_8c5e71978702f5ff
	for _, __obf_817531fc9afb059d := range __obf_496337d4762bc775 {
		if __obf_817531fc9afb059d.Cmp(__obf_edfb45ae797263ef) > 0 {
			__obf_edfb45ae797263ef = __obf_817531fc9afb059d
		}
	}
	return __obf_edfb45ae797263ef
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_8c5e71978702f5ff Decimal, __obf_496337d4762bc775 ...Decimal) Decimal {
	__obf_9338e3e7cd64adde := __obf_8c5e71978702f5ff
	for _, __obf_817531fc9afb059d := range __obf_496337d4762bc775 {
		__obf_9338e3e7cd64adde = __obf_9338e3e7cd64adde.Add(__obf_817531fc9afb059d)
	}

	return __obf_9338e3e7cd64adde
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_8c5e71978702f5ff Decimal, __obf_496337d4762bc775 ...Decimal) Decimal {
	__obf_b98868b73db53827 := New(int64(len(__obf_496337d4762bc775)+1), 0)
	__obf_512e9f06910cca45 := Sum(__obf_8c5e71978702f5ff, __obf_496337d4762bc775...)
	return __obf_512e9f06910cca45.Div(__obf_b98868b73db53827)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_d5adf70c420dc23a Decimal, __obf_5c7cf1148c3170fc Decimal) (Decimal, Decimal) {
	__obf_d5adf70c420dc23a.__obf_4b849233a95eee86()
	__obf_5c7cf1148c3170fc.__obf_4b849233a95eee86()

	if __obf_d5adf70c420dc23a.__obf_fed6fd50f0fcc834 < __obf_5c7cf1148c3170fc.__obf_fed6fd50f0fcc834 {
		return __obf_d5adf70c420dc23a, __obf_5c7cf1148c3170fc.__obf_452282ed04deedc9(__obf_d5adf70c420dc23a.__obf_fed6fd50f0fcc834)
	} else if __obf_d5adf70c420dc23a.__obf_fed6fd50f0fcc834 > __obf_5c7cf1148c3170fc.__obf_fed6fd50f0fcc834 {
		return __obf_d5adf70c420dc23a.__obf_452282ed04deedc9(__obf_5c7cf1148c3170fc.__obf_fed6fd50f0fcc834), __obf_5c7cf1148c3170fc
	}

	return __obf_d5adf70c420dc23a, __obf_5c7cf1148c3170fc
}

func __obf_422c9f22dd4dfc68(__obf_a1a825437ca7e6ad any) (string, error) {
	var __obf_3c68acfde9d69829 []byte

	switch __obf_d056d9dcafa97698 := __obf_a1a825437ca7e6ad.(type) {
	case string:
		__obf_3c68acfde9d69829 = []byte(__obf_d056d9dcafa97698)
	case []byte:
		__obf_3c68acfde9d69829 = __obf_d056d9dcafa97698
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_a1a825437ca7e6ad, __obf_a1a825437ca7e6ad)
	}

	// If the amount is quoted, strip the quotes
	if len(__obf_3c68acfde9d69829) > 2 && __obf_3c68acfde9d69829[0] == '"' && __obf_3c68acfde9d69829[len(__obf_3c68acfde9d69829)-1] == '"' {
		__obf_3c68acfde9d69829 = __obf_3c68acfde9d69829[1 : len(__obf_3c68acfde9d69829)-1]
	}
	return string(__obf_3c68acfde9d69829), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the database.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_c4701b3bb28cd2ae Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_c4701b3bb28cd2ae,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_c4701b3bb28cd2ae *NullDecimal) Scan(__obf_a1a825437ca7e6ad any) error {
	if __obf_a1a825437ca7e6ad == nil {
		__obf_c4701b3bb28cd2ae.Valid = false
		return nil
	}
	__obf_c4701b3bb28cd2ae.Valid = true
	return __obf_c4701b3bb28cd2ae.Decimal.Scan(__obf_a1a825437ca7e6ad)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_c4701b3bb28cd2ae NullDecimal) Value() (driver.Value, error) {
	if !__obf_c4701b3bb28cd2ae.Valid {
		return nil, nil
	}
	return __obf_c4701b3bb28cd2ae.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_c4701b3bb28cd2ae *NullDecimal) UnmarshalJSON(__obf_9571474f72c71159 []byte) error {
	if string(__obf_9571474f72c71159) == "null" {
		__obf_c4701b3bb28cd2ae.Valid = false
		return nil
	}
	__obf_c4701b3bb28cd2ae.Valid = true
	return __obf_c4701b3bb28cd2ae.Decimal.UnmarshalJSON(__obf_9571474f72c71159)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_c4701b3bb28cd2ae NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_c4701b3bb28cd2ae.Valid {
		return []byte("null"), nil
	}
	return __obf_c4701b3bb28cd2ae.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_c4701b3bb28cd2ae *NullDecimal) UnmarshalText(__obf_13dc8aff84147625 []byte) error {
	__obf_4f9552183401314b := string(__obf_13dc8aff84147625)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_4f9552183401314b == "" {
		__obf_c4701b3bb28cd2ae.Valid = false
		return nil
	}
	if __obf_db664194ea89d0f0 := __obf_c4701b3bb28cd2ae.Decimal.UnmarshalText(__obf_13dc8aff84147625); __obf_db664194ea89d0f0 != nil {
		__obf_c4701b3bb28cd2ae.Valid = false
		return __obf_db664194ea89d0f0
	}
	__obf_c4701b3bb28cd2ae.Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_c4701b3bb28cd2ae NullDecimal) MarshalText() (__obf_13dc8aff84147625 []byte, __obf_db664194ea89d0f0 error) {
	if !__obf_c4701b3bb28cd2ae.Valid {
		return []byte{}, nil
	}
	return __obf_c4701b3bb28cd2ae.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_c4701b3bb28cd2ae Decimal) Atan() Decimal {
	if __obf_c4701b3bb28cd2ae.Equal(NewFromFloat(0.0)) {
		return __obf_c4701b3bb28cd2ae
	}
	if __obf_c4701b3bb28cd2ae.GreaterThan(NewFromFloat(0.0)) {
		return __obf_c4701b3bb28cd2ae.__obf_feefee19a4e86d6f()
	}
	return __obf_c4701b3bb28cd2ae.Neg().__obf_feefee19a4e86d6f().Neg()
}

func (__obf_c4701b3bb28cd2ae Decimal) __obf_9b10bdc0b70efcce() Decimal {
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
	__obf_a3d7b54ee2bf67c1 := __obf_c4701b3bb28cd2ae.Mul(__obf_c4701b3bb28cd2ae)
	__obf_51d5e8287cfa93a4 := P0.Mul(__obf_a3d7b54ee2bf67c1).Add(P1).Mul(__obf_a3d7b54ee2bf67c1).Add(P2).Mul(__obf_a3d7b54ee2bf67c1).Add(P3).Mul(__obf_a3d7b54ee2bf67c1).Add(P4).Mul(__obf_a3d7b54ee2bf67c1)
	__obf_11e48796e09ab512 := __obf_a3d7b54ee2bf67c1.Add(Q0).Mul(__obf_a3d7b54ee2bf67c1).Add(Q1).Mul(__obf_a3d7b54ee2bf67c1).Add(Q2).Mul(__obf_a3d7b54ee2bf67c1).Add(Q3).Mul(__obf_a3d7b54ee2bf67c1).Add(Q4)
	__obf_a3d7b54ee2bf67c1 = __obf_51d5e8287cfa93a4.Div(__obf_11e48796e09ab512)
	__obf_a3d7b54ee2bf67c1 = __obf_c4701b3bb28cd2ae.Mul(__obf_a3d7b54ee2bf67c1).Add(__obf_c4701b3bb28cd2ae)
	return __obf_a3d7b54ee2bf67c1
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_c4701b3bb28cd2ae Decimal) __obf_feefee19a4e86d6f() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)      // tan(3*pi/8)
	__obf_6960b1e99def3323 := NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_c4701b3bb28cd2ae.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_c4701b3bb28cd2ae.__obf_9b10bdc0b70efcce()
	}
	if __obf_c4701b3bb28cd2ae.GreaterThan(Tan3pio8) {
		return __obf_6960b1e99def3323.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_c4701b3bb28cd2ae).__obf_9b10bdc0b70efcce()).Add(Morebits)
	}
	return __obf_6960b1e99def3323.Div(NewFromFloat(4.0)).Add((__obf_c4701b3bb28cd2ae.Sub(NewFromFloat(1.0)).Div(__obf_c4701b3bb28cd2ae.Add(NewFromFloat(1.0)))).__obf_9b10bdc0b70efcce()).Add(NewFromFloat(0.5).Mul(Morebits))
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
func (__obf_c4701b3bb28cd2ae Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_c4701b3bb28cd2ae.Equal(NewFromFloat(0.0)) {
		return __obf_c4701b3bb28cd2ae
	}
	// make argument positive but save the sign
	__obf_81ae3aeca8ee3212 := false
	if __obf_c4701b3bb28cd2ae.LessThan(NewFromFloat(0.0)) {
		__obf_c4701b3bb28cd2ae = __obf_c4701b3bb28cd2ae.Neg()
		__obf_81ae3aeca8ee3212 = true
	}

	__obf_3da1d28da36e0ae0 := __obf_c4701b3bb28cd2ae.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_328c7a071989cb79 := NewFromFloat(float64(__obf_3da1d28da36e0ae0)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_3da1d28da36e0ae0&1 == 1 {
		__obf_3da1d28da36e0ae0++
		__obf_328c7a071989cb79 = __obf_328c7a071989cb79.Add(NewFromFloat(1.0))
	}
	__obf_3da1d28da36e0ae0 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_3da1d28da36e0ae0 > 3 {
		__obf_81ae3aeca8ee3212 = !__obf_81ae3aeca8ee3212
		__obf_3da1d28da36e0ae0 -= 4
	}
	__obf_a3d7b54ee2bf67c1 := __obf_c4701b3bb28cd2ae.Sub(__obf_328c7a071989cb79.Mul(PI4A)).Sub(__obf_328c7a071989cb79.Mul(PI4B)).Sub(__obf_328c7a071989cb79.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_d0882bb783f82b54 := __obf_a3d7b54ee2bf67c1.Mul(__obf_a3d7b54ee2bf67c1)

	if __obf_3da1d28da36e0ae0 == 1 || __obf_3da1d28da36e0ae0 == 2 {
		__obf_a3aa9136d792c9d2 := __obf_d0882bb783f82b54.Mul(__obf_d0882bb783f82b54).Mul(_cos[0].Mul(__obf_d0882bb783f82b54).Add(_cos[1]).Mul(__obf_d0882bb783f82b54).Add(_cos[2]).Mul(__obf_d0882bb783f82b54).Add(_cos[3]).Mul(__obf_d0882bb783f82b54).Add(_cos[4]).Mul(__obf_d0882bb783f82b54).Add(_cos[5]))
		__obf_328c7a071989cb79 = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_d0882bb783f82b54)).Add(__obf_a3aa9136d792c9d2)
	} else {
		__obf_328c7a071989cb79 = __obf_a3d7b54ee2bf67c1.Add(__obf_a3d7b54ee2bf67c1.Mul(__obf_d0882bb783f82b54).Mul(_sin[0].Mul(__obf_d0882bb783f82b54).Add(_sin[1]).Mul(__obf_d0882bb783f82b54).Add(_sin[2]).Mul(__obf_d0882bb783f82b54).Add(_sin[3]).Mul(__obf_d0882bb783f82b54).Add(_sin[4]).Mul(__obf_d0882bb783f82b54).Add(_sin[5])))
	}
	if __obf_81ae3aeca8ee3212 {
		__obf_328c7a071989cb79 = __obf_328c7a071989cb79.Neg()
	}
	return __obf_328c7a071989cb79
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
func (__obf_c4701b3bb28cd2ae Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	// make argument positive
	__obf_81ae3aeca8ee3212 := false
	if __obf_c4701b3bb28cd2ae.LessThan(NewFromFloat(0.0)) {
		__obf_c4701b3bb28cd2ae = __obf_c4701b3bb28cd2ae.Neg()
	}

	__obf_3da1d28da36e0ae0 := __obf_c4701b3bb28cd2ae.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_328c7a071989cb79 := NewFromFloat(float64(__obf_3da1d28da36e0ae0)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_3da1d28da36e0ae0&1 == 1 {
		__obf_3da1d28da36e0ae0++
		__obf_328c7a071989cb79 = __obf_328c7a071989cb79.Add(NewFromFloat(1.0))
	}
	__obf_3da1d28da36e0ae0 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_3da1d28da36e0ae0 > 3 {
		__obf_81ae3aeca8ee3212 = !__obf_81ae3aeca8ee3212
		__obf_3da1d28da36e0ae0 -= 4
	}
	if __obf_3da1d28da36e0ae0 > 1 {
		__obf_81ae3aeca8ee3212 = !__obf_81ae3aeca8ee3212
	}

	__obf_a3d7b54ee2bf67c1 := __obf_c4701b3bb28cd2ae.Sub(__obf_328c7a071989cb79.Mul(PI4A)).Sub(__obf_328c7a071989cb79.Mul(PI4B)).Sub(__obf_328c7a071989cb79.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_d0882bb783f82b54 := __obf_a3d7b54ee2bf67c1.Mul(__obf_a3d7b54ee2bf67c1)

	if __obf_3da1d28da36e0ae0 == 1 || __obf_3da1d28da36e0ae0 == 2 {
		__obf_328c7a071989cb79 = __obf_a3d7b54ee2bf67c1.Add(__obf_a3d7b54ee2bf67c1.Mul(__obf_d0882bb783f82b54).Mul(_sin[0].Mul(__obf_d0882bb783f82b54).Add(_sin[1]).Mul(__obf_d0882bb783f82b54).Add(_sin[2]).Mul(__obf_d0882bb783f82b54).Add(_sin[3]).Mul(__obf_d0882bb783f82b54).Add(_sin[4]).Mul(__obf_d0882bb783f82b54).Add(_sin[5])))
	} else {
		__obf_a3aa9136d792c9d2 := __obf_d0882bb783f82b54.Mul(__obf_d0882bb783f82b54).Mul(_cos[0].Mul(__obf_d0882bb783f82b54).Add(_cos[1]).Mul(__obf_d0882bb783f82b54).Add(_cos[2]).Mul(__obf_d0882bb783f82b54).Add(_cos[3]).Mul(__obf_d0882bb783f82b54).Add(_cos[4]).Mul(__obf_d0882bb783f82b54).Add(_cos[5]))
		__obf_328c7a071989cb79 = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_d0882bb783f82b54)).Add(__obf_a3aa9136d792c9d2)
	}
	if __obf_81ae3aeca8ee3212 {
		__obf_328c7a071989cb79 = __obf_328c7a071989cb79.Neg()
	}
	return __obf_328c7a071989cb79
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
func (__obf_c4701b3bb28cd2ae Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_c4701b3bb28cd2ae.Equal(NewFromFloat(0.0)) {
		return __obf_c4701b3bb28cd2ae
	}

	// make argument positive but save the sign
	__obf_81ae3aeca8ee3212 := false
	if __obf_c4701b3bb28cd2ae.LessThan(NewFromFloat(0.0)) {
		__obf_c4701b3bb28cd2ae = __obf_c4701b3bb28cd2ae.Neg()
		__obf_81ae3aeca8ee3212 = true
	}

	__obf_3da1d28da36e0ae0 := __obf_c4701b3bb28cd2ae.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_328c7a071989cb79 := NewFromFloat(float64(__obf_3da1d28da36e0ae0)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_3da1d28da36e0ae0&1 == 1 {
		__obf_3da1d28da36e0ae0++
		__obf_328c7a071989cb79 = __obf_328c7a071989cb79.Add(NewFromFloat(1.0))
	}

	__obf_a3d7b54ee2bf67c1 := __obf_c4701b3bb28cd2ae.Sub(__obf_328c7a071989cb79.Mul(PI4A)).Sub(__obf_328c7a071989cb79.Mul(PI4B)).Sub(__obf_328c7a071989cb79.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_d0882bb783f82b54 := __obf_a3d7b54ee2bf67c1.Mul(__obf_a3d7b54ee2bf67c1)

	if __obf_d0882bb783f82b54.GreaterThan(NewFromFloat(1e-14)) {
		__obf_a3aa9136d792c9d2 := __obf_d0882bb783f82b54.Mul(_tanP[0].Mul(__obf_d0882bb783f82b54).Add(_tanP[1]).Mul(__obf_d0882bb783f82b54).Add(_tanP[2]))
		__obf_2a4c30c5f6fc7eb1 := __obf_d0882bb783f82b54.Add(_tanQ[1]).Mul(__obf_d0882bb783f82b54).Add(_tanQ[2]).Mul(__obf_d0882bb783f82b54).Add(_tanQ[3]).Mul(__obf_d0882bb783f82b54).Add(_tanQ[4])
		__obf_328c7a071989cb79 = __obf_a3d7b54ee2bf67c1.Add(__obf_a3d7b54ee2bf67c1.Mul(__obf_a3aa9136d792c9d2.Div(__obf_2a4c30c5f6fc7eb1)))
	} else {
		__obf_328c7a071989cb79 = __obf_a3d7b54ee2bf67c1
	}
	if __obf_3da1d28da36e0ae0&2 == 2 {
		__obf_328c7a071989cb79 = NewFromFloat(-1.0).Div(__obf_328c7a071989cb79)
	}
	if __obf_81ae3aeca8ee3212 {
		__obf_328c7a071989cb79 = __obf_328c7a071989cb79.Neg()
	}
	return __obf_328c7a071989cb79
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

type __obf_adf9c80aee40a9d8 struct {
	__obf_c4701b3bb28cd2ae [800]byte // digits, big-endian representation
	__obf_087eecd15d19676e int       // number of digits used
	__obf_d30835ea2dfed4ba int       // decimal point
	__obf_468f39cd022d1bd4 bool      // negative flag
	__obf_61c081fcd0e60518 bool      // discarded nonzero digits beyond d[:nd]
}

func (__obf_935771d361ffde13 *__obf_adf9c80aee40a9d8) String() string {
	__obf_ee70f56e2047ecea := 10 + __obf_935771d361ffde13.__obf_087eecd15d19676e
	if __obf_935771d361ffde13.__obf_d30835ea2dfed4ba > 0 {
		__obf_ee70f56e2047ecea += __obf_935771d361ffde13.__obf_d30835ea2dfed4ba
	}
	if __obf_935771d361ffde13.__obf_d30835ea2dfed4ba < 0 {
		__obf_ee70f56e2047ecea += -__obf_935771d361ffde13.__obf_d30835ea2dfed4ba
	}

	__obf_bfcbbe05e5ecf4e8 := make([]byte, __obf_ee70f56e2047ecea)
	__obf_a3aa9136d792c9d2 := 0
	switch {
	case __obf_935771d361ffde13.__obf_087eecd15d19676e == 0:
		return "0"

	case __obf_935771d361ffde13.__obf_d30835ea2dfed4ba <= 0:
		// zeros fill space between decimal point and digits
		__obf_bfcbbe05e5ecf4e8[__obf_a3aa9136d792c9d2] = '0'
		__obf_a3aa9136d792c9d2++
		__obf_bfcbbe05e5ecf4e8[__obf_a3aa9136d792c9d2] = '.'
		__obf_a3aa9136d792c9d2++
		__obf_a3aa9136d792c9d2 += __obf_74fc1dfb8fd8cefc(__obf_bfcbbe05e5ecf4e8[__obf_a3aa9136d792c9d2 : __obf_a3aa9136d792c9d2+-__obf_935771d361ffde13.__obf_d30835ea2dfed4ba])
		__obf_a3aa9136d792c9d2 += copy(__obf_bfcbbe05e5ecf4e8[__obf_a3aa9136d792c9d2:], __obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[0:__obf_935771d361ffde13.__obf_087eecd15d19676e])

	case __obf_935771d361ffde13.__obf_d30835ea2dfed4ba < __obf_935771d361ffde13.__obf_087eecd15d19676e:
		// decimal point in middle of digits
		__obf_a3aa9136d792c9d2 += copy(__obf_bfcbbe05e5ecf4e8[__obf_a3aa9136d792c9d2:], __obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[0:__obf_935771d361ffde13.__obf_d30835ea2dfed4ba])
		__obf_bfcbbe05e5ecf4e8[__obf_a3aa9136d792c9d2] = '.'
		__obf_a3aa9136d792c9d2++
		__obf_a3aa9136d792c9d2 += copy(__obf_bfcbbe05e5ecf4e8[__obf_a3aa9136d792c9d2:], __obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_935771d361ffde13.__obf_d30835ea2dfed4ba:__obf_935771d361ffde13.__obf_087eecd15d19676e])

	default:
		// zeros fill space between digits and decimal point
		__obf_a3aa9136d792c9d2 += copy(__obf_bfcbbe05e5ecf4e8[__obf_a3aa9136d792c9d2:], __obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[0:__obf_935771d361ffde13.__obf_087eecd15d19676e])
		__obf_a3aa9136d792c9d2 += __obf_74fc1dfb8fd8cefc(__obf_bfcbbe05e5ecf4e8[__obf_a3aa9136d792c9d2 : __obf_a3aa9136d792c9d2+__obf_935771d361ffde13.__obf_d30835ea2dfed4ba-__obf_935771d361ffde13.__obf_087eecd15d19676e])
	}
	return string(__obf_bfcbbe05e5ecf4e8[0:__obf_a3aa9136d792c9d2])
}

func __obf_74fc1dfb8fd8cefc(__obf_bc31ca34245c46b2 []byte) int {
	for __obf_159ec70197834006 := range __obf_bc31ca34245c46b2 {
		__obf_bc31ca34245c46b2[__obf_159ec70197834006] = '0'
	}
	return len(__obf_bc31ca34245c46b2)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_5629d577bcd80015(__obf_935771d361ffde13 *__obf_adf9c80aee40a9d8) {
	for __obf_935771d361ffde13.__obf_087eecd15d19676e > 0 && __obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_935771d361ffde13.__obf_087eecd15d19676e-1] == '0' {
		__obf_935771d361ffde13.__obf_087eecd15d19676e--
	}
	if __obf_935771d361ffde13.__obf_087eecd15d19676e == 0 {
		__obf_935771d361ffde13.__obf_d30835ea2dfed4ba = 0
	}
}

// Assign v to a.
func (__obf_935771d361ffde13 *__obf_adf9c80aee40a9d8) Assign(__obf_d056d9dcafa97698 uint64) {
	var __obf_bfcbbe05e5ecf4e8 [24]byte

	// Write reversed decimal in buf.
	__obf_ee70f56e2047ecea := 0
	for __obf_d056d9dcafa97698 > 0 {
		__obf_1a9155ade8f45275 := __obf_d056d9dcafa97698 / 10
		__obf_d056d9dcafa97698 -= 10 * __obf_1a9155ade8f45275
		__obf_bfcbbe05e5ecf4e8[__obf_ee70f56e2047ecea] = byte(__obf_d056d9dcafa97698 + '0')
		__obf_ee70f56e2047ecea++
		__obf_d056d9dcafa97698 = __obf_1a9155ade8f45275
	}

	// Reverse again to produce forward decimal in a.d.
	__obf_935771d361ffde13.__obf_087eecd15d19676e = 0
	for __obf_ee70f56e2047ecea--; __obf_ee70f56e2047ecea >= 0; __obf_ee70f56e2047ecea-- {
		__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_935771d361ffde13.__obf_087eecd15d19676e] = __obf_bfcbbe05e5ecf4e8[__obf_ee70f56e2047ecea]
		__obf_935771d361ffde13.__obf_087eecd15d19676e++
	}
	__obf_935771d361ffde13.__obf_d30835ea2dfed4ba = __obf_935771d361ffde13.__obf_087eecd15d19676e
	__obf_5629d577bcd80015(__obf_935771d361ffde13)
}

// Maximum shift that we can do in one pass without overflow.
// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
const __obf_d2d0d2b688b2c2d0 = 32 << (^uint(0) >> 63)
const __obf_e0dd74a58ba1d50b = __obf_d2d0d2b688b2c2d0 - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_b0367e2a3d347c80(__obf_935771d361ffde13 *__obf_adf9c80aee40a9d8, __obf_5e8ea31acf2647c9 uint) {
	__obf_e1eef1cd0f9a218e := 0 // read pointer
	__obf_a3aa9136d792c9d2 := 0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_ee70f56e2047ecea uint
	for ; __obf_ee70f56e2047ecea>>__obf_5e8ea31acf2647c9 == 0; __obf_e1eef1cd0f9a218e++ {
		if __obf_e1eef1cd0f9a218e >= __obf_935771d361ffde13.__obf_087eecd15d19676e {
			if __obf_ee70f56e2047ecea == 0 {
				// a == 0; shouldn't get here, but handle anyway.
				__obf_935771d361ffde13.__obf_087eecd15d19676e = 0
				return
			}
			for __obf_ee70f56e2047ecea>>__obf_5e8ea31acf2647c9 == 0 {
				__obf_ee70f56e2047ecea = __obf_ee70f56e2047ecea * 10
				__obf_e1eef1cd0f9a218e++
			}
			break
		}
		__obf_49577c83f724294c := uint(__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_e1eef1cd0f9a218e])
		__obf_ee70f56e2047ecea = __obf_ee70f56e2047ecea*10 + __obf_49577c83f724294c - '0'
	}
	__obf_935771d361ffde13.__obf_d30835ea2dfed4ba -= __obf_e1eef1cd0f9a218e - 1

	var __obf_525bb2180cac7175 uint = (1 << __obf_5e8ea31acf2647c9) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_e1eef1cd0f9a218e < __obf_935771d361ffde13.__obf_087eecd15d19676e; __obf_e1eef1cd0f9a218e++ {
		__obf_49577c83f724294c := uint(__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_e1eef1cd0f9a218e])
		__obf_eccea68eb780e8e4 := __obf_ee70f56e2047ecea >> __obf_5e8ea31acf2647c9
		__obf_ee70f56e2047ecea &= __obf_525bb2180cac7175
		__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_a3aa9136d792c9d2] = byte(__obf_eccea68eb780e8e4 + '0')
		__obf_a3aa9136d792c9d2++
		__obf_ee70f56e2047ecea = __obf_ee70f56e2047ecea*10 + __obf_49577c83f724294c - '0'
	}

	// Put down extra digits.
	for __obf_ee70f56e2047ecea > 0 {
		__obf_eccea68eb780e8e4 := __obf_ee70f56e2047ecea >> __obf_5e8ea31acf2647c9
		__obf_ee70f56e2047ecea &= __obf_525bb2180cac7175
		if __obf_a3aa9136d792c9d2 < len(__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae) {
			__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_a3aa9136d792c9d2] = byte(__obf_eccea68eb780e8e4 + '0')
			__obf_a3aa9136d792c9d2++
		} else if __obf_eccea68eb780e8e4 > 0 {
			__obf_935771d361ffde13.__obf_61c081fcd0e60518 = true
		}
		__obf_ee70f56e2047ecea = __obf_ee70f56e2047ecea * 10
	}

	__obf_935771d361ffde13.__obf_087eecd15d19676e = __obf_a3aa9136d792c9d2
	__obf_5629d577bcd80015(__obf_935771d361ffde13)
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

type __obf_42279be0dabef80b struct {
	__obf_50e3c0ac578b1723 int    // number of new digits
	__obf_3eaef284baf6f77f string // minus one digit if original < a.
}

var __obf_6a6135b4029744de = []__obf_42279be0dabef80b{
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
func __obf_c4576f7872364318(__obf_83e64ebb51b2b0f6 []byte, __obf_9a1986da7e3bc21f string) bool {
	for __obf_159ec70197834006 := 0; __obf_159ec70197834006 < len(__obf_9a1986da7e3bc21f); __obf_159ec70197834006++ {
		if __obf_159ec70197834006 >= len(__obf_83e64ebb51b2b0f6) {
			return true
		}
		if __obf_83e64ebb51b2b0f6[__obf_159ec70197834006] != __obf_9a1986da7e3bc21f[__obf_159ec70197834006] {
			return __obf_83e64ebb51b2b0f6[__obf_159ec70197834006] < __obf_9a1986da7e3bc21f[__obf_159ec70197834006]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_904c814ac089b9ba(__obf_935771d361ffde13 *__obf_adf9c80aee40a9d8, __obf_5e8ea31acf2647c9 uint) {
	__obf_50e3c0ac578b1723 := __obf_6a6135b4029744de[__obf_5e8ea31acf2647c9].__obf_50e3c0ac578b1723
	if __obf_c4576f7872364318(__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[0:__obf_935771d361ffde13.__obf_087eecd15d19676e], __obf_6a6135b4029744de[__obf_5e8ea31acf2647c9].__obf_3eaef284baf6f77f) {
		__obf_50e3c0ac578b1723--
	}

	__obf_e1eef1cd0f9a218e := __obf_935771d361ffde13.__obf_087eecd15d19676e                          // read index
	__obf_a3aa9136d792c9d2 := __obf_935771d361ffde13.__obf_087eecd15d19676e + __obf_50e3c0ac578b1723 // write index

	// Pick up a digit, put down a digit.
	var __obf_ee70f56e2047ecea uint
	for __obf_e1eef1cd0f9a218e--; __obf_e1eef1cd0f9a218e >= 0; __obf_e1eef1cd0f9a218e-- {
		__obf_ee70f56e2047ecea += (uint(__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_e1eef1cd0f9a218e]) - '0') << __obf_5e8ea31acf2647c9
		__obf_a0b0f6dbd513f468 := __obf_ee70f56e2047ecea / 10
		__obf_a6e216b8f5e28306 := __obf_ee70f56e2047ecea - 10*__obf_a0b0f6dbd513f468
		__obf_a3aa9136d792c9d2--
		if __obf_a3aa9136d792c9d2 < len(__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae) {
			__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_a3aa9136d792c9d2] = byte(__obf_a6e216b8f5e28306 + '0')
		} else if __obf_a6e216b8f5e28306 != 0 {
			__obf_935771d361ffde13.__obf_61c081fcd0e60518 = true
		}
		__obf_ee70f56e2047ecea = __obf_a0b0f6dbd513f468
	}

	// Put down extra digits.
	for __obf_ee70f56e2047ecea > 0 {
		__obf_a0b0f6dbd513f468 := __obf_ee70f56e2047ecea / 10
		__obf_a6e216b8f5e28306 := __obf_ee70f56e2047ecea - 10*__obf_a0b0f6dbd513f468
		__obf_a3aa9136d792c9d2--
		if __obf_a3aa9136d792c9d2 < len(__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae) {
			__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_a3aa9136d792c9d2] = byte(__obf_a6e216b8f5e28306 + '0')
		} else if __obf_a6e216b8f5e28306 != 0 {
			__obf_935771d361ffde13.__obf_61c081fcd0e60518 = true
		}
		__obf_ee70f56e2047ecea = __obf_a0b0f6dbd513f468
	}

	__obf_935771d361ffde13.__obf_087eecd15d19676e += __obf_50e3c0ac578b1723
	if __obf_935771d361ffde13.__obf_087eecd15d19676e >= len(__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae) {
		__obf_935771d361ffde13.__obf_087eecd15d19676e = len(__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae)
	}
	__obf_935771d361ffde13.__obf_d30835ea2dfed4ba += __obf_50e3c0ac578b1723
	__obf_5629d577bcd80015(__obf_935771d361ffde13)
}

// Binary shift left (k > 0) or right (k < 0).
func (__obf_935771d361ffde13 *__obf_adf9c80aee40a9d8) Shift(__obf_5e8ea31acf2647c9 int) {
	switch {
	case __obf_935771d361ffde13.__obf_087eecd15d19676e == 0:
		// nothing to do: a == 0
	case __obf_5e8ea31acf2647c9 > 0:
		for __obf_5e8ea31acf2647c9 > __obf_e0dd74a58ba1d50b {
			__obf_904c814ac089b9ba(__obf_935771d361ffde13, __obf_e0dd74a58ba1d50b)
			__obf_5e8ea31acf2647c9 -= __obf_e0dd74a58ba1d50b
		}
		__obf_904c814ac089b9ba(__obf_935771d361ffde13, uint(__obf_5e8ea31acf2647c9))
	case __obf_5e8ea31acf2647c9 < 0:
		for __obf_5e8ea31acf2647c9 < -__obf_e0dd74a58ba1d50b {
			__obf_b0367e2a3d347c80(__obf_935771d361ffde13, __obf_e0dd74a58ba1d50b)
			__obf_5e8ea31acf2647c9 += __obf_e0dd74a58ba1d50b
		}
		__obf_b0367e2a3d347c80(__obf_935771d361ffde13, uint(-__obf_5e8ea31acf2647c9))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_4baaf85f4fa15018(__obf_935771d361ffde13 *__obf_adf9c80aee40a9d8, __obf_087eecd15d19676e int) bool {
	if __obf_087eecd15d19676e < 0 || __obf_087eecd15d19676e >= __obf_935771d361ffde13.__obf_087eecd15d19676e {
		return false
	}
	if __obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_087eecd15d19676e] == '5' && __obf_087eecd15d19676e+1 == __obf_935771d361ffde13.__obf_087eecd15d19676e { // exactly halfway - round to even
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_935771d361ffde13.__obf_61c081fcd0e60518 {
			return true
		}
		return __obf_087eecd15d19676e > 0 && (__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_087eecd15d19676e-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_087eecd15d19676e] >= '5'
}

// Round a to nd digits (or fewer).
// If nd is zero, it means we're rounding
// just to the left of the digits, as in
// 0.09 -> 0.1.
func (__obf_935771d361ffde13 *__obf_adf9c80aee40a9d8) Round(__obf_087eecd15d19676e int) {
	if __obf_087eecd15d19676e < 0 || __obf_087eecd15d19676e >= __obf_935771d361ffde13.__obf_087eecd15d19676e {
		return
	}
	if __obf_4baaf85f4fa15018(__obf_935771d361ffde13, __obf_087eecd15d19676e) {
		__obf_935771d361ffde13.RoundUp(__obf_087eecd15d19676e)
	} else {
		__obf_935771d361ffde13.RoundDown(__obf_087eecd15d19676e)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_935771d361ffde13 *__obf_adf9c80aee40a9d8) RoundDown(__obf_087eecd15d19676e int) {
	if __obf_087eecd15d19676e < 0 || __obf_087eecd15d19676e >= __obf_935771d361ffde13.__obf_087eecd15d19676e {
		return
	}
	__obf_935771d361ffde13.__obf_087eecd15d19676e = __obf_087eecd15d19676e
	__obf_5629d577bcd80015(__obf_935771d361ffde13)
}

// Round a up to nd digits (or fewer).
func (__obf_935771d361ffde13 *__obf_adf9c80aee40a9d8) RoundUp(__obf_087eecd15d19676e int) {
	if __obf_087eecd15d19676e < 0 || __obf_087eecd15d19676e >= __obf_935771d361ffde13.__obf_087eecd15d19676e {
		return
	}

	// round up
	for __obf_159ec70197834006 := __obf_087eecd15d19676e - 1; __obf_159ec70197834006 >= 0; __obf_159ec70197834006-- {
		__obf_49577c83f724294c := __obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_159ec70197834006]
		if __obf_49577c83f724294c < '9' { // can stop after this digit
			__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_159ec70197834006]++
			__obf_935771d361ffde13.__obf_087eecd15d19676e = __obf_159ec70197834006 + 1
			return
		}
	}

	// Number is all 9s.
	// Change to single 1 with adjusted decimal point.
	__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[0] = '1'
	__obf_935771d361ffde13.__obf_087eecd15d19676e = 1
	__obf_935771d361ffde13.__obf_d30835ea2dfed4ba++
}

// Extract integer part, rounded appropriately.
// No guarantees about overflow.
func (__obf_935771d361ffde13 *__obf_adf9c80aee40a9d8) RoundedInteger() uint64 {
	if __obf_935771d361ffde13.__obf_d30835ea2dfed4ba > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_159ec70197834006 int
	__obf_ee70f56e2047ecea := uint64(0)
	for __obf_159ec70197834006 = 0; __obf_159ec70197834006 < __obf_935771d361ffde13.__obf_d30835ea2dfed4ba && __obf_159ec70197834006 < __obf_935771d361ffde13.__obf_087eecd15d19676e; __obf_159ec70197834006++ {
		__obf_ee70f56e2047ecea = __obf_ee70f56e2047ecea*10 + uint64(__obf_935771d361ffde13.__obf_c4701b3bb28cd2ae[__obf_159ec70197834006]-'0')
	}
	for ; __obf_159ec70197834006 < __obf_935771d361ffde13.__obf_d30835ea2dfed4ba; __obf_159ec70197834006++ {
		__obf_ee70f56e2047ecea *= 10
	}
	if __obf_4baaf85f4fa15018(__obf_935771d361ffde13, __obf_935771d361ffde13.__obf_d30835ea2dfed4ba) {
		__obf_ee70f56e2047ecea++
	}
	return __obf_ee70f56e2047ecea
}
