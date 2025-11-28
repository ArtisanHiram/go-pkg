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
package __obf_210b94d2c4b23455

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

var __obf_22b4ed810afb99f1 = big.NewInt(0)
var __obf_3f4a47d8e0bfca14 = big.NewInt(1)
var __obf_c2d07968aecbaf08 = big.NewInt(2)
var __obf_9ae61dd5c563398c = big.NewInt(4)
var __obf_8266beb7fdfb4bd8 = big.NewInt(5)
var __obf_db69efbd472ddd1a = big.NewInt(10)
var __obf_b1624ddba33f3a22 = big.NewInt(20)

var __obf_cbf1348e181861b8 = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_d129e62d54e19e72 *big.Int

	// NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.
	__obf_e060caa13549e714 int32
}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_d129e62d54e19e72 int64, __obf_e060caa13549e714 int32) Decimal {
	return Decimal{
		__obf_d129e62d54e19e72: big.NewInt(__obf_d129e62d54e19e72),
		__obf_e060caa13549e714: __obf_e060caa13549e714,
	}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_d129e62d54e19e72 int64) Decimal {
	return Decimal{
		__obf_d129e62d54e19e72: big.NewInt(__obf_d129e62d54e19e72),
		__obf_e060caa13549e714: 0,
	}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_d129e62d54e19e72 int32) Decimal {
	return Decimal{
		__obf_d129e62d54e19e72: big.NewInt(int64(__obf_d129e62d54e19e72)),
		__obf_e060caa13549e714: 0,
	}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_d129e62d54e19e72 uint64) Decimal {
	return Decimal{
		__obf_d129e62d54e19e72: new(big.Int).SetUint64(__obf_d129e62d54e19e72),
		__obf_e060caa13549e714: 0,
	}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_d129e62d54e19e72 *big.Int, __obf_e060caa13549e714 int32) Decimal {
	return Decimal{
		__obf_d129e62d54e19e72: new(big.Int).Set(__obf_d129e62d54e19e72),
		__obf_e060caa13549e714: __obf_e060caa13549e714,
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
func NewFromBigRat(__obf_d129e62d54e19e72 *big.Rat, __obf_11e1ee221b72ac06 int32) Decimal {
	return Decimal{
		__obf_d129e62d54e19e72: new(big.Int).Set(__obf_d129e62d54e19e72.Num()),
		__obf_e060caa13549e714: 0,
	}.DivRound(Decimal{
		__obf_d129e62d54e19e72: new(big.Int).Set(__obf_d129e62d54e19e72.Denom()),
		__obf_e060caa13549e714: 0,
	}, __obf_11e1ee221b72ac06)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_d129e62d54e19e72 string) (Decimal, error) {
	__obf_be0d3274bea5aaca := __obf_d129e62d54e19e72
	var __obf_c3292eb373737f56 string
	var __obf_e060caa13549e714 int64

	// Check if number is using scientific notation
	__obf_3025d432b76cd800 := strings.IndexAny(__obf_d129e62d54e19e72, "Ee")
	if __obf_3025d432b76cd800 != -1 {
		__obf_fe82159dcd1d1cd8, __obf_4ff66c6003689543 := strconv.ParseInt(__obf_d129e62d54e19e72[__obf_3025d432b76cd800+1:], 10, 32)
		if __obf_4ff66c6003689543 != nil {
			if __obf_e3bbc3fb50b5b7f4, __obf_88b906f77e2efee5 := __obf_4ff66c6003689543.(*strconv.NumError); __obf_88b906f77e2efee5 && __obf_e3bbc3fb50b5b7f4.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_d129e62d54e19e72)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_d129e62d54e19e72)
		}
		__obf_d129e62d54e19e72 = __obf_d129e62d54e19e72[:__obf_3025d432b76cd800]
		__obf_e060caa13549e714 = __obf_fe82159dcd1d1cd8
	}

	__obf_e319e26529181031 := -1
	__obf_b6db106f2c68c1ab := len(__obf_d129e62d54e19e72)
	for __obf_dec02b860bfb4fdd := 0; __obf_dec02b860bfb4fdd < __obf_b6db106f2c68c1ab; __obf_dec02b860bfb4fdd++ {
		if __obf_d129e62d54e19e72[__obf_dec02b860bfb4fdd] == '.' {
			if __obf_e319e26529181031 > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_d129e62d54e19e72)
			}
			__obf_e319e26529181031 = __obf_dec02b860bfb4fdd
		}
	}

	if __obf_e319e26529181031 == -1 {
		// There is no decimal point, we can just parse the original string as
		// an int
		__obf_c3292eb373737f56 = __obf_d129e62d54e19e72
	} else {
		if __obf_e319e26529181031+1 < __obf_b6db106f2c68c1ab {
			__obf_c3292eb373737f56 = __obf_d129e62d54e19e72[:__obf_e319e26529181031] + __obf_d129e62d54e19e72[__obf_e319e26529181031+1:]
		} else {
			__obf_c3292eb373737f56 = __obf_d129e62d54e19e72[:__obf_e319e26529181031]
		}
		__obf_fe82159dcd1d1cd8 := -len(__obf_d129e62d54e19e72[__obf_e319e26529181031+1:])
		__obf_e060caa13549e714 += int64(__obf_fe82159dcd1d1cd8)
	}

	var __obf_1172a5bbbc2c5e00 *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_c3292eb373737f56) <= 18 {
		__obf_af6981f57535aad3, __obf_4ff66c6003689543 := strconv.ParseInt(__obf_c3292eb373737f56, 10, 64)
		if __obf_4ff66c6003689543 != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_d129e62d54e19e72)
		}
		__obf_1172a5bbbc2c5e00 = big.NewInt(__obf_af6981f57535aad3)
	} else {
		__obf_1172a5bbbc2c5e00 = new(big.Int)
		_, __obf_88b906f77e2efee5 := __obf_1172a5bbbc2c5e00.SetString(__obf_c3292eb373737f56, 10)
		if !__obf_88b906f77e2efee5 {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_d129e62d54e19e72)
		}
	}

	if __obf_e060caa13549e714 < math.MinInt32 || __obf_e060caa13549e714 > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_be0d3274bea5aaca)
	}

	return Decimal{
		__obf_d129e62d54e19e72: __obf_1172a5bbbc2c5e00,
		__obf_e060caa13549e714: int32(__obf_e060caa13549e714),
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
func NewFromFormattedString(__obf_d129e62d54e19e72 string, __obf_b28a579441acb812 *regexp.Regexp) (Decimal, error) {
	__obf_7ce7bfa69b6271d8 := __obf_b28a579441acb812.ReplaceAllString(__obf_d129e62d54e19e72, "")
	__obf_45d3205e59c45c8b, __obf_4ff66c6003689543 := NewFromString(__obf_7ce7bfa69b6271d8)
	if __obf_4ff66c6003689543 != nil {
		return Decimal{}, __obf_4ff66c6003689543
	}
	return __obf_45d3205e59c45c8b, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_d129e62d54e19e72 string) Decimal {
	__obf_b99754db9573b7e4, __obf_4ff66c6003689543 := NewFromString(__obf_d129e62d54e19e72)
	if __obf_4ff66c6003689543 != nil {
		panic(__obf_4ff66c6003689543)
	}
	return __obf_b99754db9573b7e4
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
func NewFromFloat(__obf_d129e62d54e19e72 float64) Decimal {
	if __obf_d129e62d54e19e72 == 0 {
		return New(0, 0)
	}
	return __obf_9ce1240f6180b1b1(__obf_d129e62d54e19e72, math.Float64bits(__obf_d129e62d54e19e72), &__obf_7e93547175be045c)
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
func NewFromFloat32(__obf_d129e62d54e19e72 float32) Decimal {
	if __obf_d129e62d54e19e72 == 0 {
		return New(0, 0)
	}
	// XOR is workaround for https://github.com/golang/go/issues/26285
	__obf_a7affdfda169fd03 := math.Float32bits(__obf_d129e62d54e19e72) ^ 0x80808080
	return __obf_9ce1240f6180b1b1(float64(__obf_d129e62d54e19e72), uint64(__obf_a7affdfda169fd03)^0x80808080, &__obf_75b4ac7d1ab2e9d4)
}

func __obf_9ce1240f6180b1b1(__obf_bff636421fe05c92 float64, __obf_af7dea8b5f2166d7 uint64, __obf_c7a132b4ff25bccf *__obf_73320c94d3fe82cc) Decimal {
	if math.IsNaN(__obf_bff636421fe05c92) || math.IsInf(__obf_bff636421fe05c92, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_bff636421fe05c92))
	}
	__obf_e060caa13549e714 := int(__obf_af7dea8b5f2166d7>>__obf_c7a132b4ff25bccf.__obf_b197d08c88c50ba7) & (1<<__obf_c7a132b4ff25bccf.__obf_2e8a1cc5b4179be2 - 1)
	__obf_5f6f09d40b1d194a := __obf_af7dea8b5f2166d7 & (uint64(1)<<__obf_c7a132b4ff25bccf.__obf_b197d08c88c50ba7 - 1)

	switch __obf_e060caa13549e714 {
	case 0:
		// denormalized
		__obf_e060caa13549e714++

	default:
		// add implicit top bit
		__obf_5f6f09d40b1d194a |= uint64(1) << __obf_c7a132b4ff25bccf.__obf_b197d08c88c50ba7
	}
	__obf_e060caa13549e714 += __obf_c7a132b4ff25bccf.__obf_3dc7e69ee4097b24

	var __obf_45d3205e59c45c8b __obf_210b94d2c4b23455
	__obf_45d3205e59c45c8b.Assign(__obf_5f6f09d40b1d194a)
	__obf_45d3205e59c45c8b.Shift(__obf_e060caa13549e714 - int(__obf_c7a132b4ff25bccf.__obf_b197d08c88c50ba7))
	__obf_45d3205e59c45c8b.__obf_4ae2cd7e4db15eb2 = __obf_af7dea8b5f2166d7>>(__obf_c7a132b4ff25bccf.__obf_2e8a1cc5b4179be2+__obf_c7a132b4ff25bccf.__obf_b197d08c88c50ba7) != 0

	__obf_705af1f438fcfe9e(&__obf_45d3205e59c45c8b, __obf_5f6f09d40b1d194a, __obf_e060caa13549e714, __obf_c7a132b4ff25bccf)
	// If less than 19 digits, we can do calculation in an int64.
	if __obf_45d3205e59c45c8b.__obf_2298ab2cdf3cb6bd < 19 {
		__obf_5f72611079aef925 := int64(0)
		__obf_763dd0cca8ce709c := int64(1)
		for __obf_dec02b860bfb4fdd := __obf_45d3205e59c45c8b.__obf_2298ab2cdf3cb6bd - 1; __obf_dec02b860bfb4fdd >= 0; __obf_dec02b860bfb4fdd-- {
			__obf_5f72611079aef925 += __obf_763dd0cca8ce709c * int64(__obf_45d3205e59c45c8b.__obf_45d3205e59c45c8b[__obf_dec02b860bfb4fdd]-'0')
			__obf_763dd0cca8ce709c *= 10
		}
		if __obf_45d3205e59c45c8b.__obf_4ae2cd7e4db15eb2 {
			__obf_5f72611079aef925 *= -1
		}
		return Decimal{__obf_d129e62d54e19e72: big.NewInt(__obf_5f72611079aef925), __obf_e060caa13549e714: int32(__obf_45d3205e59c45c8b.__obf_44d25d407353f76c) - int32(__obf_45d3205e59c45c8b.__obf_2298ab2cdf3cb6bd)}
	}
	__obf_1172a5bbbc2c5e00 := new(big.Int)
	__obf_1172a5bbbc2c5e00, __obf_88b906f77e2efee5 := __obf_1172a5bbbc2c5e00.SetString(string(__obf_45d3205e59c45c8b.__obf_45d3205e59c45c8b[:__obf_45d3205e59c45c8b.__obf_2298ab2cdf3cb6bd]), 10)
	if __obf_88b906f77e2efee5 {
		return Decimal{__obf_d129e62d54e19e72: __obf_1172a5bbbc2c5e00, __obf_e060caa13549e714: int32(__obf_45d3205e59c45c8b.__obf_44d25d407353f76c) - int32(__obf_45d3205e59c45c8b.__obf_2298ab2cdf3cb6bd)}
	}

	return NewFromFloatWithExponent(__obf_bff636421fe05c92, int32(__obf_45d3205e59c45c8b.__obf_44d25d407353f76c)-int32(__obf_45d3205e59c45c8b.__obf_2298ab2cdf3cb6bd))
}

// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
// number of fractional digits.
//
// Example:
//
//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
func NewFromFloatWithExponent(__obf_d129e62d54e19e72 float64, __obf_e060caa13549e714 int32) Decimal {
	if math.IsNaN(__obf_d129e62d54e19e72) || math.IsInf(__obf_d129e62d54e19e72, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_d129e62d54e19e72))
	}

	__obf_af7dea8b5f2166d7 := math.Float64bits(__obf_d129e62d54e19e72)
	__obf_5f6f09d40b1d194a := __obf_af7dea8b5f2166d7 & (1<<52 - 1)
	__obf_4c484a1a2b826adc := int32((__obf_af7dea8b5f2166d7 >> 52) & (1<<11 - 1))
	__obf_8c897683b6557fb7 := __obf_af7dea8b5f2166d7 >> 63

	if __obf_4c484a1a2b826adc == 0 {
		// specials
		if __obf_5f6f09d40b1d194a == 0 {
			return Decimal{}
		}
		// subnormal
		__obf_4c484a1a2b826adc++
	} else {
		// normal
		__obf_5f6f09d40b1d194a |= 1 << 52
	}

	__obf_4c484a1a2b826adc -= 1023 + 52

	// normalizing base-2 values
	for __obf_5f6f09d40b1d194a&1 == 0 {
		__obf_5f6f09d40b1d194a = __obf_5f6f09d40b1d194a >> 1
		__obf_4c484a1a2b826adc++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_e060caa13549e714 < 0 && __obf_e060caa13549e714 < __obf_4c484a1a2b826adc {
		if __obf_4c484a1a2b826adc < 0 {
			__obf_e060caa13549e714 = __obf_4c484a1a2b826adc
		} else {
			__obf_e060caa13549e714 = 0
		}
	}

	// representing 10^M * 2^N as 5^M * 2^(M+N)
	__obf_4c484a1a2b826adc -= __obf_e060caa13549e714

	__obf_d279bc1d6f7d239f := big.NewInt(1)
	__obf_ea3cca6b0d01a7ab := big.NewInt(int64(__obf_5f6f09d40b1d194a))

	// applying 5^M
	if __obf_e060caa13549e714 > 0 {
		__obf_d279bc1d6f7d239f = __obf_d279bc1d6f7d239f.SetInt64(int64(__obf_e060caa13549e714))
		__obf_d279bc1d6f7d239f = __obf_d279bc1d6f7d239f.Exp(__obf_8266beb7fdfb4bd8, __obf_d279bc1d6f7d239f, nil)
	} else if __obf_e060caa13549e714 < 0 {
		__obf_d279bc1d6f7d239f = __obf_d279bc1d6f7d239f.SetInt64(-int64(__obf_e060caa13549e714))
		__obf_d279bc1d6f7d239f = __obf_d279bc1d6f7d239f.Exp(__obf_8266beb7fdfb4bd8, __obf_d279bc1d6f7d239f, nil)
		__obf_ea3cca6b0d01a7ab = __obf_ea3cca6b0d01a7ab.Mul(__obf_ea3cca6b0d01a7ab, __obf_d279bc1d6f7d239f)
		__obf_d279bc1d6f7d239f = __obf_d279bc1d6f7d239f.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_4c484a1a2b826adc > 0 {
		__obf_ea3cca6b0d01a7ab = __obf_ea3cca6b0d01a7ab.Lsh(__obf_ea3cca6b0d01a7ab, uint(__obf_4c484a1a2b826adc))
	} else if __obf_4c484a1a2b826adc < 0 {
		__obf_d279bc1d6f7d239f = __obf_d279bc1d6f7d239f.Lsh(__obf_d279bc1d6f7d239f, uint(-__obf_4c484a1a2b826adc))
	}

	// rounding and downscaling
	if __obf_e060caa13549e714 > 0 || __obf_4c484a1a2b826adc < 0 {
		__obf_197dd2c9f2d22564 := new(big.Int).Rsh(__obf_d279bc1d6f7d239f, 1)
		__obf_ea3cca6b0d01a7ab = __obf_ea3cca6b0d01a7ab.Add(__obf_ea3cca6b0d01a7ab, __obf_197dd2c9f2d22564)
		__obf_ea3cca6b0d01a7ab = __obf_ea3cca6b0d01a7ab.Quo(__obf_ea3cca6b0d01a7ab, __obf_d279bc1d6f7d239f)
	}

	if __obf_8c897683b6557fb7 == 1 {
		__obf_ea3cca6b0d01a7ab = __obf_ea3cca6b0d01a7ab.Neg(__obf_ea3cca6b0d01a7ab)
	}

	return Decimal{
		__obf_d129e62d54e19e72: __obf_ea3cca6b0d01a7ab,
		__obf_e060caa13549e714: __obf_e060caa13549e714,
	}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_45d3205e59c45c8b Decimal) Copy() Decimal {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()
	return Decimal{
		__obf_d129e62d54e19e72: new(big.Int).Set(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72),
		__obf_e060caa13549e714: __obf_45d3205e59c45c8b.__obf_e060caa13549e714,
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
func (__obf_45d3205e59c45c8b Decimal) __obf_df6a3d24ae564ff2(__obf_e060caa13549e714 int32) Decimal {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()

	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 == __obf_e060caa13549e714 {
		return Decimal{
			new(big.Int).Set(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72),
			__obf_45d3205e59c45c8b.__obf_e060caa13549e714,
		}
	}

	// NOTE(vadim): must convert exps to float64 before - to prevent overflow
	__obf_ca0879c0043346e6 := math.Abs(float64(__obf_e060caa13549e714) - float64(__obf_45d3205e59c45c8b.__obf_e060caa13549e714))
	__obf_d129e62d54e19e72 := new(big.Int).Set(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72)

	__obf_1bd06f29b40b12a8 := new(big.Int).Exp(__obf_db69efbd472ddd1a, big.NewInt(int64(__obf_ca0879c0043346e6)), nil)
	if __obf_e060caa13549e714 > __obf_45d3205e59c45c8b.__obf_e060caa13549e714 {
		__obf_d129e62d54e19e72 = __obf_d129e62d54e19e72.Quo(__obf_d129e62d54e19e72, __obf_1bd06f29b40b12a8)
	} else if __obf_e060caa13549e714 < __obf_45d3205e59c45c8b.__obf_e060caa13549e714 {
		__obf_d129e62d54e19e72 = __obf_d129e62d54e19e72.Mul(__obf_d129e62d54e19e72, __obf_1bd06f29b40b12a8)
	}

	return Decimal{
		__obf_d129e62d54e19e72: __obf_d129e62d54e19e72,
		__obf_e060caa13549e714: __obf_e060caa13549e714,
	}
}

// Abs returns the absolute value of the decimal.
func (__obf_45d3205e59c45c8b Decimal) Abs() Decimal {
	if !__obf_45d3205e59c45c8b.IsNegative() {
		return __obf_45d3205e59c45c8b
	}
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()
	__obf_d076b88ce9dd34a3 := new(big.Int).Abs(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72)
	return Decimal{
		__obf_d129e62d54e19e72: __obf_d076b88ce9dd34a3,
		__obf_e060caa13549e714: __obf_45d3205e59c45c8b.__obf_e060caa13549e714,
	}
}

// Add returns d + d2.
func (__obf_45d3205e59c45c8b Decimal) Add(__obf_0cf0ba5fcabed996 Decimal) Decimal {
	__obf_9f54e2cc2cf3003a, __obf_94a58003f07ed165 := RescalePair(__obf_45d3205e59c45c8b, __obf_0cf0ba5fcabed996)

	__obf_365e52066c3c9a6d := new(big.Int).Add(__obf_9f54e2cc2cf3003a.__obf_d129e62d54e19e72, __obf_94a58003f07ed165.__obf_d129e62d54e19e72)
	return Decimal{
		__obf_d129e62d54e19e72: __obf_365e52066c3c9a6d,
		__obf_e060caa13549e714: __obf_9f54e2cc2cf3003a.__obf_e060caa13549e714,
	}
}

// Sub returns d - d2.
func (__obf_45d3205e59c45c8b Decimal) Sub(__obf_0cf0ba5fcabed996 Decimal) Decimal {
	__obf_9f54e2cc2cf3003a, __obf_94a58003f07ed165 := RescalePair(__obf_45d3205e59c45c8b, __obf_0cf0ba5fcabed996)

	__obf_365e52066c3c9a6d := new(big.Int).Sub(__obf_9f54e2cc2cf3003a.__obf_d129e62d54e19e72, __obf_94a58003f07ed165.__obf_d129e62d54e19e72)
	return Decimal{
		__obf_d129e62d54e19e72: __obf_365e52066c3c9a6d,
		__obf_e060caa13549e714: __obf_9f54e2cc2cf3003a.__obf_e060caa13549e714,
	}
}

// Neg returns -d.
func (__obf_45d3205e59c45c8b Decimal) Neg() Decimal {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()
	__obf_bff636421fe05c92 := new(big.Int).Neg(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72)
	return Decimal{
		__obf_d129e62d54e19e72: __obf_bff636421fe05c92,
		__obf_e060caa13549e714: __obf_45d3205e59c45c8b.__obf_e060caa13549e714,
	}
}

// Mul returns d * d2.
func (__obf_45d3205e59c45c8b Decimal) Mul(__obf_0cf0ba5fcabed996 Decimal) Decimal {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()
	__obf_0cf0ba5fcabed996.__obf_f2008d48a41fad81()

	__obf_8f308ea14bce0ad8 := int64(__obf_45d3205e59c45c8b.__obf_e060caa13549e714) + int64(__obf_0cf0ba5fcabed996.__obf_e060caa13549e714)
	if __obf_8f308ea14bce0ad8 > math.MaxInt32 || __obf_8f308ea14bce0ad8 < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_8f308ea14bce0ad8))
	}

	__obf_365e52066c3c9a6d := new(big.Int).Mul(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72, __obf_0cf0ba5fcabed996.__obf_d129e62d54e19e72)
	return Decimal{
		__obf_d129e62d54e19e72: __obf_365e52066c3c9a6d,
		__obf_e060caa13549e714: int32(__obf_8f308ea14bce0ad8),
	}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_45d3205e59c45c8b Decimal) Shift(__obf_3fb6d47fc6364f57 int32) Decimal {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()
	return Decimal{
		__obf_d129e62d54e19e72: new(big.Int).Set(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72),
		__obf_e060caa13549e714: __obf_45d3205e59c45c8b.__obf_e060caa13549e714 + __obf_3fb6d47fc6364f57,
	}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_45d3205e59c45c8b Decimal) Div(__obf_0cf0ba5fcabed996 Decimal) Decimal {
	return __obf_45d3205e59c45c8b.DivRound(__obf_0cf0ba5fcabed996, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_45d3205e59c45c8b Decimal) QuoRem(__obf_0cf0ba5fcabed996 Decimal, __obf_11e1ee221b72ac06 int32) (Decimal, Decimal) {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()
	__obf_0cf0ba5fcabed996.__obf_f2008d48a41fad81()
	if __obf_0cf0ba5fcabed996.__obf_d129e62d54e19e72.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_d5b204fdd13393fe := -__obf_11e1ee221b72ac06
	__obf_e3bbc3fb50b5b7f4 := int64(__obf_45d3205e59c45c8b.__obf_e060caa13549e714) - int64(__obf_0cf0ba5fcabed996.__obf_e060caa13549e714) - int64(__obf_d5b204fdd13393fe)
	if __obf_e3bbc3fb50b5b7f4 > math.MaxInt32 || __obf_e3bbc3fb50b5b7f4 < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_dbab0e049a5a449a, __obf_188653a6283f3920, __obf_8db7cc363bfa156b big.Int
	var __obf_86ca07af5a58f6d2 int32
	// d = a 10^ea
	// d2 = b 10^eb
	if __obf_e3bbc3fb50b5b7f4 < 0 {
		__obf_dbab0e049a5a449a = *__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72
		__obf_8db7cc363bfa156b.SetInt64(-__obf_e3bbc3fb50b5b7f4)
		__obf_188653a6283f3920.Exp(__obf_db69efbd472ddd1a, &__obf_8db7cc363bfa156b, nil)
		__obf_188653a6283f3920.Mul(__obf_0cf0ba5fcabed996.__obf_d129e62d54e19e72, &__obf_188653a6283f3920)
		__obf_86ca07af5a58f6d2 = __obf_45d3205e59c45c8b.__obf_e060caa13549e714
		// now aa = a
		//     bb = b 10^(scale + eb - ea)
	} else {
		__obf_8db7cc363bfa156b.SetInt64(__obf_e3bbc3fb50b5b7f4)
		__obf_dbab0e049a5a449a.Exp(__obf_db69efbd472ddd1a, &__obf_8db7cc363bfa156b, nil)
		__obf_dbab0e049a5a449a.Mul(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72, &__obf_dbab0e049a5a449a)
		__obf_188653a6283f3920 = *__obf_0cf0ba5fcabed996.__obf_d129e62d54e19e72
		__obf_86ca07af5a58f6d2 = __obf_d5b204fdd13393fe + __obf_0cf0ba5fcabed996.__obf_e060caa13549e714
		// now aa = a ^ (ea - eb - scale)
		//     bb = b
	}
	var __obf_606198092cfca255, __obf_8bfc98156daaf406 big.Int
	__obf_606198092cfca255.QuoRem(&__obf_dbab0e049a5a449a, &__obf_188653a6283f3920, &__obf_8bfc98156daaf406)
	__obf_1d3f163dea4891f7 := Decimal{__obf_d129e62d54e19e72: &__obf_606198092cfca255, __obf_e060caa13549e714: __obf_d5b204fdd13393fe}
	__obf_e0822ef1e566e731 := Decimal{__obf_d129e62d54e19e72: &__obf_8bfc98156daaf406, __obf_e060caa13549e714: __obf_86ca07af5a58f6d2}
	return __obf_1d3f163dea4891f7, __obf_e0822ef1e566e731
}

// DivRound divides and rounds to a given precision
// i.e. to an integer multiple of 10^(-precision)
//
//	for a positive quotient digit 5 is rounded up, away from 0
//	if the quotient is negative then digit 5 is rounded down, away from 0
//
// Note that precision<0 is allowed as input.
func (__obf_45d3205e59c45c8b Decimal) DivRound(__obf_0cf0ba5fcabed996 Decimal, __obf_11e1ee221b72ac06 int32) Decimal {
	// QuoRem already checks initialization
	__obf_606198092cfca255, __obf_8bfc98156daaf406 := __obf_45d3205e59c45c8b.QuoRem(__obf_0cf0ba5fcabed996, __obf_11e1ee221b72ac06)
	// the actual rounding decision is based on comparing r*10^precision and d2/2
	// instead compare 2 r 10 ^precision and d2
	var __obf_206d5ca2e1f1a4d8 big.Int
	__obf_206d5ca2e1f1a4d8.Abs(__obf_8bfc98156daaf406.__obf_d129e62d54e19e72)
	__obf_206d5ca2e1f1a4d8.Lsh(&__obf_206d5ca2e1f1a4d8, 1)
	// now rv2 = abs(r.value) * 2
	__obf_d5c51636d4d3d119 := Decimal{__obf_d129e62d54e19e72: &__obf_206d5ca2e1f1a4d8, __obf_e060caa13549e714: __obf_8bfc98156daaf406.__obf_e060caa13549e714 + __obf_11e1ee221b72ac06}
	// r2 is now 2 * r * 10 ^ precision
	var __obf_820da776a4f8225c = __obf_d5c51636d4d3d119.Cmp(__obf_0cf0ba5fcabed996.Abs())

	if __obf_820da776a4f8225c < 0 {
		return __obf_606198092cfca255
	}

	if __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.Sign()*__obf_0cf0ba5fcabed996.__obf_d129e62d54e19e72.Sign() < 0 {
		return __obf_606198092cfca255.Sub(New(1, -__obf_11e1ee221b72ac06))
	}

	return __obf_606198092cfca255.Add(New(1, -__obf_11e1ee221b72ac06))
}

// Mod returns d % d2.
func (__obf_45d3205e59c45c8b Decimal) Mod(__obf_0cf0ba5fcabed996 Decimal) Decimal {
	_, __obf_8bfc98156daaf406 := __obf_45d3205e59c45c8b.QuoRem(__obf_0cf0ba5fcabed996, 0)
	return __obf_8bfc98156daaf406
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
func (__obf_45d3205e59c45c8b Decimal) Pow(__obf_0cf0ba5fcabed996 Decimal) Decimal {
	__obf_eb483018f4ad84ba := __obf_45d3205e59c45c8b.Sign()
	__obf_7cee98855a794973 := __obf_0cf0ba5fcabed996.Sign()

	if __obf_eb483018f4ad84ba == 0 {
		if __obf_7cee98855a794973 == 0 {
			return Decimal{}
		}
		if __obf_7cee98855a794973 == 1 {
			return Decimal{__obf_22b4ed810afb99f1, 0}
		}
		if __obf_7cee98855a794973 == -1 {
			return Decimal{}
		}
	}

	if __obf_7cee98855a794973 == 0 {
		return Decimal{__obf_3f4a47d8e0bfca14, 0}
	}

	// TODO: optimize extraction of fractional part
	__obf_0d80c9d64cb9ba53 := Decimal{__obf_3f4a47d8e0bfca14, 0}
	__obf_5a8d81b38a707822, __obf_d765cbeb0fad3eb1 := __obf_0cf0ba5fcabed996.QuoRem(__obf_0d80c9d64cb9ba53, 0)

	if __obf_eb483018f4ad84ba == -1 && !__obf_d765cbeb0fad3eb1.IsZero() {
		return Decimal{}
	}

	__obf_1c0d2ef70e063f12, _ := __obf_45d3205e59c45c8b.PowBigInt(__obf_5a8d81b38a707822.__obf_d129e62d54e19e72)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_d765cbeb0fad3eb1.__obf_d129e62d54e19e72.Sign() == 0 {
		return __obf_1c0d2ef70e063f12
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_31936d0a141531ab := __obf_45d3205e59c45c8b.NumDigits()
	__obf_81a59cc026252429 := __obf_0cf0ba5fcabed996.NumDigits()

	__obf_11e1ee221b72ac06 := __obf_31936d0a141531ab

	if __obf_81a59cc026252429 > __obf_11e1ee221b72ac06 {
		__obf_11e1ee221b72ac06 += __obf_81a59cc026252429
	}

	__obf_11e1ee221b72ac06 += 6

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_a2f20404d9c0413f, __obf_4ff66c6003689543 := __obf_45d3205e59c45c8b.Abs().Ln(-__obf_45d3205e59c45c8b.__obf_e060caa13549e714 + int32(__obf_11e1ee221b72ac06))
	if __obf_4ff66c6003689543 != nil {
		return Decimal{}
	}

	__obf_a2f20404d9c0413f = __obf_a2f20404d9c0413f.Mul(__obf_d765cbeb0fad3eb1)

	__obf_a2f20404d9c0413f, __obf_4ff66c6003689543 = __obf_a2f20404d9c0413f.ExpTaylor(-__obf_45d3205e59c45c8b.__obf_e060caa13549e714 + int32(__obf_11e1ee221b72ac06))
	if __obf_4ff66c6003689543 != nil {
		return Decimal{}
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_8106b464d75159d5 := __obf_1c0d2ef70e063f12.Mul(__obf_a2f20404d9c0413f)

	return __obf_8106b464d75159d5
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
func (__obf_45d3205e59c45c8b Decimal) PowWithPrecision(__obf_0cf0ba5fcabed996 Decimal, __obf_11e1ee221b72ac06 int32) (Decimal, error) {
	__obf_eb483018f4ad84ba := __obf_45d3205e59c45c8b.Sign()
	__obf_7cee98855a794973 := __obf_0cf0ba5fcabed996.Sign()

	if __obf_eb483018f4ad84ba == 0 {
		if __obf_7cee98855a794973 == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_7cee98855a794973 == 1 {
			return Decimal{__obf_22b4ed810afb99f1, 0}, nil
		}
		if __obf_7cee98855a794973 == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_7cee98855a794973 == 0 {
		return Decimal{__obf_3f4a47d8e0bfca14, 0}, nil
	}

	// TODO: optimize extraction of fractional part
	__obf_0d80c9d64cb9ba53 := Decimal{__obf_3f4a47d8e0bfca14, 0}
	__obf_5a8d81b38a707822, __obf_d765cbeb0fad3eb1 := __obf_0cf0ba5fcabed996.QuoRem(__obf_0d80c9d64cb9ba53, 0)

	if __obf_eb483018f4ad84ba == -1 && !__obf_d765cbeb0fad3eb1.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}

	__obf_1c0d2ef70e063f12, _ := __obf_45d3205e59c45c8b.__obf_50ed535e01749f93(__obf_5a8d81b38a707822.__obf_d129e62d54e19e72, __obf_11e1ee221b72ac06)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_d765cbeb0fad3eb1.__obf_d129e62d54e19e72.Sign() == 0 {
		return __obf_1c0d2ef70e063f12, nil
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_31936d0a141531ab := __obf_45d3205e59c45c8b.NumDigits()
	__obf_81a59cc026252429 := __obf_0cf0ba5fcabed996.NumDigits()

	if int32(__obf_31936d0a141531ab) > __obf_11e1ee221b72ac06 {
		__obf_11e1ee221b72ac06 = int32(__obf_31936d0a141531ab)
	}
	if int32(__obf_81a59cc026252429) > __obf_11e1ee221b72ac06 {
		__obf_11e1ee221b72ac06 += int32(__obf_81a59cc026252429)
	}
	// increase precision by 10 to compensate for errors in further calculations
	__obf_11e1ee221b72ac06 += 10

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_a2f20404d9c0413f, __obf_4ff66c6003689543 := __obf_45d3205e59c45c8b.Abs().Ln(__obf_11e1ee221b72ac06)
	if __obf_4ff66c6003689543 != nil {
		return Decimal{}, __obf_4ff66c6003689543
	}

	__obf_a2f20404d9c0413f = __obf_a2f20404d9c0413f.Mul(__obf_d765cbeb0fad3eb1)

	__obf_a2f20404d9c0413f, __obf_4ff66c6003689543 = __obf_a2f20404d9c0413f.ExpTaylor(__obf_11e1ee221b72ac06)
	if __obf_4ff66c6003689543 != nil {
		return Decimal{}, __obf_4ff66c6003689543
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_8106b464d75159d5 := __obf_1c0d2ef70e063f12.Mul(__obf_a2f20404d9c0413f)

	return __obf_8106b464d75159d5, nil
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
func (__obf_45d3205e59c45c8b Decimal) PowInt32(__obf_e060caa13549e714 int32) (Decimal, error) {
	if __obf_45d3205e59c45c8b.IsZero() && __obf_e060caa13549e714 == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_1154313d67e84cf1 := __obf_e060caa13549e714 < 0
	__obf_e060caa13549e714 = __obf_c318d0106586e799(__obf_e060caa13549e714)

	__obf_46e3648cfcdd2b28, __obf_99b8e31c695cfff0 := __obf_45d3205e59c45c8b, New(1, 0)

	for __obf_e060caa13549e714 > 0 {
		if __obf_e060caa13549e714%2 == 1 {
			__obf_99b8e31c695cfff0 = __obf_99b8e31c695cfff0.Mul(__obf_46e3648cfcdd2b28)
		}
		__obf_e060caa13549e714 /= 2

		if __obf_e060caa13549e714 > 0 {
			__obf_46e3648cfcdd2b28 = __obf_46e3648cfcdd2b28.Mul(__obf_46e3648cfcdd2b28)
		}
	}

	if __obf_1154313d67e84cf1 {
		return New(1, 0).DivRound(__obf_99b8e31c695cfff0, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_99b8e31c695cfff0, nil
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
func (__obf_45d3205e59c45c8b Decimal) PowBigInt(__obf_e060caa13549e714 *big.Int) (Decimal, error) {
	return __obf_45d3205e59c45c8b.__obf_50ed535e01749f93(__obf_e060caa13549e714, int32(PowPrecisionNegativeExponent))
}

func (__obf_45d3205e59c45c8b Decimal) __obf_50ed535e01749f93(__obf_e060caa13549e714 *big.Int, __obf_11e1ee221b72ac06 int32) (Decimal, error) {
	if __obf_45d3205e59c45c8b.IsZero() && __obf_e060caa13549e714.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_19123f771cd80da3 := new(big.Int).Set(__obf_e060caa13549e714)
	__obf_1154313d67e84cf1 := __obf_e060caa13549e714.Sign() < 0

	if __obf_1154313d67e84cf1 {
		__obf_19123f771cd80da3.Abs(__obf_19123f771cd80da3)
	}

	__obf_46e3648cfcdd2b28, __obf_99b8e31c695cfff0 := __obf_45d3205e59c45c8b, New(1, 0)

	for __obf_19123f771cd80da3.Sign() > 0 {
		if __obf_19123f771cd80da3.Bit(0) == 1 {
			__obf_99b8e31c695cfff0 = __obf_99b8e31c695cfff0.Mul(__obf_46e3648cfcdd2b28)
		}
		__obf_19123f771cd80da3.Rsh(__obf_19123f771cd80da3, 1)

		if __obf_19123f771cd80da3.Sign() > 0 {
			__obf_46e3648cfcdd2b28 = __obf_46e3648cfcdd2b28.Mul(__obf_46e3648cfcdd2b28)
		}
	}

	if __obf_1154313d67e84cf1 {
		return New(1, 0).DivRound(__obf_99b8e31c695cfff0, __obf_11e1ee221b72ac06), nil
	}

	return __obf_99b8e31c695cfff0, nil
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
func (__obf_45d3205e59c45c8b Decimal) ExpHullAbrham(__obf_8993eb9b55fdbfb4 uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_45d3205e59c45c8b.IsZero() {
		return Decimal{__obf_3f4a47d8e0bfca14, 0}, nil
	}

	__obf_71f34b1b424ef299 := __obf_8993eb9b55fdbfb4

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_1217a04e566bbfb7 := __obf_45d3205e59c45c8b.Abs().InexactFloat64()
	if __obf_e5803c2b2a45bf2a := __obf_1217a04e566bbfb7 / 23; __obf_e5803c2b2a45bf2a > float64(__obf_71f34b1b424ef299) && __obf_e5803c2b2a45bf2a < float64(ExpMaxIterations) {
		__obf_71f34b1b424ef299 = uint32(math.Ceil(__obf_e5803c2b2a45bf2a))
	}

	// fail if abs(d) beyond an over/underflow threshold
	__obf_0b61eca7a9f758b6 := New(23*int64(__obf_71f34b1b424ef299), 0)
	if __obf_45d3205e59c45c8b.Abs().Cmp(__obf_0b61eca7a9f758b6) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}

	// Return 1 if abs(d) small enough; this also avoids later over/underflow
	__obf_26b04c7026a42a4b := New(9, -int32(__obf_71f34b1b424ef299)-1)
	if __obf_45d3205e59c45c8b.Abs().Cmp(__obf_26b04c7026a42a4b) <= 0 {
		return Decimal{__obf_3f4a47d8e0bfca14, __obf_45d3205e59c45c8b.__obf_e060caa13549e714}, nil
	}

	// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
	__obf_e0a0c6edd92b188c := __obf_45d3205e59c45c8b.__obf_e060caa13549e714 + int32(__obf_45d3205e59c45c8b.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_e0a0c6edd92b188c < 0 {
		__obf_e0a0c6edd92b188c = 0
	}

	__obf_8431217a5a792855 := New(1, __obf_e0a0c6edd92b188c)                                                                                                                   // reduction factor
	__obf_8bfc98156daaf406 := Decimal{new(big.Int).Set(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72), __obf_45d3205e59c45c8b.__obf_e060caa13549e714 - __obf_e0a0c6edd92b188c} // reduced argument
	__obf_93bb0b0ed036819b := int32(__obf_71f34b1b424ef299) + __obf_e0a0c6edd92b188c + 2                                                                                       // precision for calculating the sum

	// Determine n, the number of therms for calculating sum
	// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
	// for solving appropriate equation, along with directed
	// roundings and simple rational bound for log10(p/abs(r))
	__obf_2fbe70a80f35d99d := __obf_8bfc98156daaf406.Abs().InexactFloat64()
	__obf_1fcc885cf7ceebaf := float64(__obf_93bb0b0ed036819b)
	__obf_d9ff312bc342a7de := math.Ceil((1.453*__obf_1fcc885cf7ceebaf - 1.182) / math.Log10(__obf_1fcc885cf7ceebaf/__obf_2fbe70a80f35d99d))
	if __obf_d9ff312bc342a7de > float64(ExpMaxIterations) || math.IsNaN(__obf_d9ff312bc342a7de) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_46e3648cfcdd2b28 := int64(__obf_d9ff312bc342a7de)

	__obf_5f72611079aef925 := New(0, 0)
	__obf_f26c5cf9e7a07902 := New(1, 0)
	__obf_0d80c9d64cb9ba53 := New(1, 0)
	for __obf_dec02b860bfb4fdd := __obf_46e3648cfcdd2b28 - 1; __obf_dec02b860bfb4fdd > 0; __obf_dec02b860bfb4fdd-- {
		__obf_5f72611079aef925.__obf_d129e62d54e19e72.SetInt64(__obf_dec02b860bfb4fdd)
		__obf_f26c5cf9e7a07902 = __obf_f26c5cf9e7a07902.Mul(__obf_8bfc98156daaf406.DivRound(__obf_5f72611079aef925, __obf_93bb0b0ed036819b))
		__obf_f26c5cf9e7a07902 = __obf_f26c5cf9e7a07902.Add(__obf_0d80c9d64cb9ba53)
	}

	__obf_1dddca309c1fa7ac := __obf_8431217a5a792855.IntPart()
	__obf_8106b464d75159d5 := New(1, 0)
	for __obf_dec02b860bfb4fdd := __obf_1dddca309c1fa7ac; __obf_dec02b860bfb4fdd > 0; __obf_dec02b860bfb4fdd-- {
		__obf_8106b464d75159d5 = __obf_8106b464d75159d5.Mul(__obf_f26c5cf9e7a07902)
	}

	__obf_87f268692f984466 := int32(__obf_8106b464d75159d5.NumDigits())

	var __obf_8fc063294e03ee10 int32
	if __obf_87f268692f984466 > __obf_c318d0106586e799(__obf_8106b464d75159d5.__obf_e060caa13549e714) {
		__obf_8fc063294e03ee10 = int32(__obf_71f34b1b424ef299) - __obf_87f268692f984466 - __obf_8106b464d75159d5.__obf_e060caa13549e714
	} else {
		__obf_8fc063294e03ee10 = int32(__obf_71f34b1b424ef299)
	}

	__obf_8106b464d75159d5 = __obf_8106b464d75159d5.Round(__obf_8fc063294e03ee10)

	return __obf_8106b464d75159d5, nil
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
func (__obf_45d3205e59c45c8b Decimal) ExpTaylor(__obf_11e1ee221b72ac06 int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_45d3205e59c45c8b.IsZero() {
		return Decimal{__obf_3f4a47d8e0bfca14, 0}.Round(__obf_11e1ee221b72ac06), nil
	}

	var __obf_a4f9306f4df24c6e Decimal
	var __obf_0c7ab54e121fbbb8 int32
	if __obf_11e1ee221b72ac06 < 0 {
		__obf_a4f9306f4df24c6e = New(1, -1)
		__obf_0c7ab54e121fbbb8 = 8
	} else {
		__obf_a4f9306f4df24c6e = New(1, -__obf_11e1ee221b72ac06-1)
		__obf_0c7ab54e121fbbb8 = __obf_11e1ee221b72ac06 + 1
	}

	__obf_415e32c31c0eb51e := __obf_45d3205e59c45c8b.Abs()
	__obf_0095034b64e8288d := __obf_45d3205e59c45c8b.Abs()
	__obf_d070f1fb485ea741 := New(1, 0)

	__obf_99b8e31c695cfff0 := New(1, 0)

	for __obf_dec02b860bfb4fdd := int64(1); ; {
		__obf_79c052c820f4705e := __obf_0095034b64e8288d.DivRound(__obf_d070f1fb485ea741, __obf_0c7ab54e121fbbb8)
		__obf_99b8e31c695cfff0 = __obf_99b8e31c695cfff0.Add(__obf_79c052c820f4705e)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_79c052c820f4705e.Cmp(__obf_a4f9306f4df24c6e) < 0 {
			break
		}

		__obf_0095034b64e8288d = __obf_0095034b64e8288d.Mul(__obf_415e32c31c0eb51e)

		__obf_dec02b860bfb4fdd++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_cbf1348e181861b8) >= int(__obf_dec02b860bfb4fdd) && !__obf_cbf1348e181861b8[__obf_dec02b860bfb4fdd-1].IsZero() {
			__obf_d070f1fb485ea741 = __obf_cbf1348e181861b8[__obf_dec02b860bfb4fdd-1]
		} else {
			// To avoid any race conditions, firstly the zero value is appended to a slice to create
			// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
			// factorial using the index notation.
			__obf_d070f1fb485ea741 = __obf_cbf1348e181861b8[__obf_dec02b860bfb4fdd-2].Mul(New(__obf_dec02b860bfb4fdd, 0))
			__obf_cbf1348e181861b8 = append(__obf_cbf1348e181861b8, Zero)
			__obf_cbf1348e181861b8[__obf_dec02b860bfb4fdd-1] = __obf_d070f1fb485ea741
		}
	}

	if __obf_45d3205e59c45c8b.Sign() < 0 {
		__obf_99b8e31c695cfff0 = New(1, 0).DivRound(__obf_99b8e31c695cfff0, __obf_11e1ee221b72ac06+1)
	}

	__obf_99b8e31c695cfff0 = __obf_99b8e31c695cfff0.Round(__obf_11e1ee221b72ac06)
	return __obf_99b8e31c695cfff0, nil
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
func (__obf_45d3205e59c45c8b Decimal) Ln(__obf_11e1ee221b72ac06 int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_45d3205e59c45c8b.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_45d3205e59c45c8b.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}

	__obf_0e15112aae14f75c := __obf_11e1ee221b72ac06 + 2
	__obf_dd4d042995f3d916 := __obf_45d3205e59c45c8b.Copy()

	var __obf_6ce8b65676a1569f, __obf_d9de1d044999b491, __obf_e777acb30d1a000d, __obf_88948d050b02633c, __obf_b9befd7e781ea088 Decimal
	__obf_6ce8b65676a1569f = __obf_dd4d042995f3d916.Sub(Decimal{__obf_3f4a47d8e0bfca14, 0})
	__obf_d9de1d044999b491 = Decimal{__obf_3f4a47d8e0bfca14, -1}

	// for decimal in range [0.9, 1.1] where ln(d) is close to 0
	__obf_9dd8fd870a8c3bd8 := false

	if __obf_6ce8b65676a1569f.Abs().Cmp(__obf_d9de1d044999b491) <= 0 {
		__obf_9dd8fd870a8c3bd8 = true
	} else {
		// reduce input decimal to range [0.1, 1)
		__obf_5f4a790be9658e76 := int32(__obf_dd4d042995f3d916.NumDigits()) + __obf_dd4d042995f3d916.__obf_e060caa13549e714
		__obf_dd4d042995f3d916.__obf_e060caa13549e714 -= __obf_5f4a790be9658e76

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_3b4375c651a23d83 := __obf_3b4375c651a23d83.__obf_334478ff4f48102f(__obf_0e15112aae14f75c)
		__obf_b9befd7e781ea088 = NewFromInt32(__obf_5f4a790be9658e76)
		__obf_b9befd7e781ea088 = __obf_b9befd7e781ea088.Mul(__obf_3b4375c651a23d83)

		__obf_6ce8b65676a1569f = __obf_dd4d042995f3d916.Sub(Decimal{__obf_3f4a47d8e0bfca14, 0})

		if __obf_6ce8b65676a1569f.Abs().Cmp(__obf_d9de1d044999b491) <= 0 {
			__obf_9dd8fd870a8c3bd8 = true
		} else {
			// initial estimate using floats
			__obf_6ae39b957b89813e := __obf_dd4d042995f3d916.InexactFloat64()
			__obf_6ce8b65676a1569f = NewFromFloat(math.Log(__obf_6ae39b957b89813e))
		}
	}

	__obf_a4f9306f4df24c6e := Decimal{__obf_3f4a47d8e0bfca14, -__obf_0e15112aae14f75c}

	if __obf_9dd8fd870a8c3bd8 {
		// Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
		// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
		// until the difference between current and next term is smaller than epsilon.
		// Coverage quite fast for decimals close to 1.0

		// z + 2
		__obf_e777acb30d1a000d = __obf_6ce8b65676a1569f.Add(Decimal{__obf_c2d07968aecbaf08, 0})
		// z / (z + 2)
		__obf_d9de1d044999b491 = __obf_6ce8b65676a1569f.DivRound(__obf_e777acb30d1a000d, __obf_0e15112aae14f75c)
		// 2 * (z / (z + 2))
		__obf_6ce8b65676a1569f = __obf_d9de1d044999b491.Add(__obf_d9de1d044999b491)
		__obf_e777acb30d1a000d = __obf_6ce8b65676a1569f.Copy()

		for __obf_46e3648cfcdd2b28 := 1; ; __obf_46e3648cfcdd2b28++ {
			// 2 * (z / (z+2))^(2n+1)
			__obf_e777acb30d1a000d = __obf_e777acb30d1a000d.Mul(__obf_d9de1d044999b491).Mul(__obf_d9de1d044999b491)

			// 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
			__obf_88948d050b02633c = NewFromInt(int64(2*__obf_46e3648cfcdd2b28 + 1))
			__obf_88948d050b02633c = __obf_e777acb30d1a000d.DivRound(__obf_88948d050b02633c, __obf_0e15112aae14f75c)

			// comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			__obf_6ce8b65676a1569f = __obf_6ce8b65676a1569f.Add(__obf_88948d050b02633c)

			if __obf_88948d050b02633c.Abs().Cmp(__obf_a4f9306f4df24c6e) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_4606edbb26850d5e Decimal
		__obf_8d4a0eb4711ebfc5 := __obf_0e15112aae14f75c*2 + 10

		for __obf_dec02b860bfb4fdd := int32(0); __obf_dec02b860bfb4fdd < __obf_8d4a0eb4711ebfc5; __obf_dec02b860bfb4fdd++ {
			// exp(a_n)
			__obf_d9de1d044999b491, _ = __obf_6ce8b65676a1569f.ExpTaylor(__obf_0e15112aae14f75c)
			// exp(a_n) - z
			__obf_e777acb30d1a000d = __obf_d9de1d044999b491.Sub(__obf_dd4d042995f3d916)
			// 2 * (exp(a_n) - z)
			__obf_e777acb30d1a000d = __obf_e777acb30d1a000d.Add(__obf_e777acb30d1a000d)
			// exp(a_n) + z
			__obf_88948d050b02633c = __obf_d9de1d044999b491.Add(__obf_dd4d042995f3d916)
			// 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_d9de1d044999b491 = __obf_e777acb30d1a000d.DivRound(__obf_88948d050b02633c, __obf_0e15112aae14f75c)
			// comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_6ce8b65676a1569f = __obf_6ce8b65676a1569f.Sub(__obf_d9de1d044999b491)

			if __obf_4606edbb26850d5e.Add(__obf_d9de1d044999b491).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_d9de1d044999b491.Abs().Cmp(__obf_a4f9306f4df24c6e) <= 0 {
				break
			}

			__obf_4606edbb26850d5e = __obf_d9de1d044999b491
		}
	}

	__obf_6ce8b65676a1569f = __obf_6ce8b65676a1569f.Add(__obf_b9befd7e781ea088)

	return __obf_6ce8b65676a1569f.Round(__obf_11e1ee221b72ac06), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_45d3205e59c45c8b Decimal) NumDigits() int {
	if __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72 == nil {
		return 1
	}

	if __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.IsInt64() {
		__obf_c26e329faa193eb7 := __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.Int64()
		// restrict fast path to integers with exact conversion to float64
		if __obf_c26e329faa193eb7 <= (1<<53) && __obf_c26e329faa193eb7 >= -(1<<53) {
			if __obf_c26e329faa193eb7 == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_c26e329faa193eb7)))) + 1
		}
	}

	__obf_0dc36b0cdf127452 := int(float64(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.BitLen()) / math.Log2(10))

	// estimatedNumDigits (lg10) may be off by 1, need to verify
	__obf_5d5c3270b08f3a6d := big.NewInt(int64(__obf_0dc36b0cdf127452))
	__obf_9ab05f2a895086a4 := __obf_5d5c3270b08f3a6d.Exp(__obf_db69efbd472ddd1a, __obf_5d5c3270b08f3a6d, nil)

	if __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.CmpAbs(__obf_9ab05f2a895086a4) >= 0 {
		return __obf_0dc36b0cdf127452 + 1
	}

	return __obf_0dc36b0cdf127452
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_45d3205e59c45c8b Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_8bfc98156daaf406 big.Int
	__obf_606198092cfca255 := new(big.Int).Set(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72)
	for __obf_dd4d042995f3d916 := __obf_c318d0106586e799(__obf_45d3205e59c45c8b.__obf_e060caa13549e714); __obf_dd4d042995f3d916 > 0; __obf_dd4d042995f3d916-- {
		__obf_606198092cfca255.QuoRem(__obf_606198092cfca255, __obf_db69efbd472ddd1a, &__obf_8bfc98156daaf406)
		if __obf_8bfc98156daaf406.Cmp(__obf_22b4ed810afb99f1) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_c318d0106586e799(__obf_46e3648cfcdd2b28 int32) int32 {
	if __obf_46e3648cfcdd2b28 < 0 {
		return -__obf_46e3648cfcdd2b28
	}
	return __obf_46e3648cfcdd2b28
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_45d3205e59c45c8b Decimal) Cmp(__obf_0cf0ba5fcabed996 Decimal) int {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()
	__obf_0cf0ba5fcabed996.__obf_f2008d48a41fad81()

	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 == __obf_0cf0ba5fcabed996.__obf_e060caa13549e714 {
		return __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.Cmp(__obf_0cf0ba5fcabed996.__obf_d129e62d54e19e72)
	}

	__obf_9f54e2cc2cf3003a, __obf_94a58003f07ed165 := RescalePair(__obf_45d3205e59c45c8b, __obf_0cf0ba5fcabed996)

	return __obf_9f54e2cc2cf3003a.__obf_d129e62d54e19e72.Cmp(__obf_94a58003f07ed165.__obf_d129e62d54e19e72)
}

// Compare compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_45d3205e59c45c8b Decimal) Compare(__obf_0cf0ba5fcabed996 Decimal) int {
	return __obf_45d3205e59c45c8b.Cmp(__obf_0cf0ba5fcabed996)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_45d3205e59c45c8b Decimal) Equal(__obf_0cf0ba5fcabed996 Decimal) bool {
	return __obf_45d3205e59c45c8b.Cmp(__obf_0cf0ba5fcabed996) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_45d3205e59c45c8b Decimal) Equals(__obf_0cf0ba5fcabed996 Decimal) bool {
	return __obf_45d3205e59c45c8b.Equal(__obf_0cf0ba5fcabed996)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_45d3205e59c45c8b Decimal) GreaterThan(__obf_0cf0ba5fcabed996 Decimal) bool {
	return __obf_45d3205e59c45c8b.Cmp(__obf_0cf0ba5fcabed996) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_45d3205e59c45c8b Decimal) GreaterThanOrEqual(__obf_0cf0ba5fcabed996 Decimal) bool {
	__obf_b56e230ff646fed8 := __obf_45d3205e59c45c8b.Cmp(__obf_0cf0ba5fcabed996)
	return __obf_b56e230ff646fed8 == 1 || __obf_b56e230ff646fed8 == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_45d3205e59c45c8b Decimal) LessThan(__obf_0cf0ba5fcabed996 Decimal) bool {
	return __obf_45d3205e59c45c8b.Cmp(__obf_0cf0ba5fcabed996) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_45d3205e59c45c8b Decimal) LessThanOrEqual(__obf_0cf0ba5fcabed996 Decimal) bool {
	__obf_b56e230ff646fed8 := __obf_45d3205e59c45c8b.Cmp(__obf_0cf0ba5fcabed996)
	return __obf_b56e230ff646fed8 == -1 || __obf_b56e230ff646fed8 == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_45d3205e59c45c8b Decimal) Sign() int {
	if __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72 == nil {
		return 0
	}
	return __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.Sign()
}

// IsPositive return
//
//	true if d > 0
//	false if d == 0
//	false if d < 0
func (__obf_45d3205e59c45c8b Decimal) IsPositive() bool {
	return __obf_45d3205e59c45c8b.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_45d3205e59c45c8b Decimal) IsNegative() bool {
	return __obf_45d3205e59c45c8b.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_45d3205e59c45c8b Decimal) IsZero() bool {
	return __obf_45d3205e59c45c8b.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_45d3205e59c45c8b Decimal) Exponent() int32 {
	return __obf_45d3205e59c45c8b.__obf_e060caa13549e714
}

// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
func (__obf_45d3205e59c45c8b Decimal) Coefficient() *big.Int {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72)
}

// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
// If coefficient cannot be represented in an int64, the result will be undefined.
func (__obf_45d3205e59c45c8b Decimal) CoefficientInt64() int64 {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()
	return __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.Int64()
}

// IntPart returns the integer component of the decimal.
func (__obf_45d3205e59c45c8b Decimal) IntPart() int64 {
	__obf_b18d5e466e04b5f2 := __obf_45d3205e59c45c8b.__obf_df6a3d24ae564ff2(0)
	return __obf_b18d5e466e04b5f2.__obf_d129e62d54e19e72.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_45d3205e59c45c8b Decimal) BigInt() *big.Int {
	__obf_b18d5e466e04b5f2 := __obf_45d3205e59c45c8b.__obf_df6a3d24ae564ff2(0)
	return __obf_b18d5e466e04b5f2.__obf_d129e62d54e19e72
}

// BigFloat returns decimal as BigFloat.
// Be aware that casting decimal to BigFloat might cause a loss of precision.
func (__obf_45d3205e59c45c8b Decimal) BigFloat() *big.Float {
	__obf_1217a04e566bbfb7 := &big.Float{}
	__obf_1217a04e566bbfb7.SetString(__obf_45d3205e59c45c8b.String())
	return __obf_1217a04e566bbfb7
}

// Rat returns a rational number representation of the decimal.
func (__obf_45d3205e59c45c8b Decimal) Rat() *big.Rat {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()
	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 <= 0 {
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_f41320ea5f9398f6 := new(big.Int).Exp(__obf_db69efbd472ddd1a, big.NewInt(-int64(__obf_45d3205e59c45c8b.__obf_e060caa13549e714)), nil)
		return new(big.Rat).SetFrac(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72, __obf_f41320ea5f9398f6)
	}

	__obf_479c7faca7381835 := new(big.Int).Exp(__obf_db69efbd472ddd1a, big.NewInt(int64(__obf_45d3205e59c45c8b.__obf_e060caa13549e714)), nil)
	__obf_ac64b5c069c3518a := new(big.Int).Mul(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72, __obf_479c7faca7381835)
	return new(big.Rat).SetFrac(__obf_ac64b5c069c3518a, __obf_3f4a47d8e0bfca14)
}

// Float64 returns the nearest float64 value for d and a bool indicating
// whether f represents d exactly.
// For more details, see the documentation for big.Rat.Float64
func (__obf_45d3205e59c45c8b Decimal) Float64() (__obf_1217a04e566bbfb7 float64, __obf_65db996ef77ff2eb bool) {
	return __obf_45d3205e59c45c8b.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_45d3205e59c45c8b Decimal) InexactFloat64() float64 {
	__obf_1217a04e566bbfb7, _ := __obf_45d3205e59c45c8b.Float64()
	return __obf_1217a04e566bbfb7
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
func (__obf_45d3205e59c45c8b Decimal) String() string {
	return __obf_45d3205e59c45c8b.string(true)
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
func (__obf_45d3205e59c45c8b Decimal) StringFixed(__obf_4f23fb2a2d558c71 int32) string {
	__obf_fd7525ed420899da := __obf_45d3205e59c45c8b.Round(__obf_4f23fb2a2d558c71)
	return __obf_fd7525ed420899da.string(false)
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
func (__obf_45d3205e59c45c8b Decimal) StringFixedBank(__obf_4f23fb2a2d558c71 int32) string {
	__obf_fd7525ed420899da := __obf_45d3205e59c45c8b.RoundBank(__obf_4f23fb2a2d558c71)
	return __obf_fd7525ed420899da.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_45d3205e59c45c8b Decimal) StringFixedCash(__obf_e8a18117b7dd2702 uint8) string {
	__obf_fd7525ed420899da := __obf_45d3205e59c45c8b.RoundCash(__obf_e8a18117b7dd2702)
	return __obf_fd7525ed420899da.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_45d3205e59c45c8b Decimal) Round(__obf_4f23fb2a2d558c71 int32) Decimal {
	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 == -__obf_4f23fb2a2d558c71 {
		return __obf_45d3205e59c45c8b
	}
	// truncate to places + 1
	__obf_5fb7c595c26c2313 := __obf_45d3205e59c45c8b.__obf_df6a3d24ae564ff2(-__obf_4f23fb2a2d558c71 - 1)

	// add sign(d) * 0.5
	if __obf_5fb7c595c26c2313.__obf_d129e62d54e19e72.Sign() < 0 {
		__obf_5fb7c595c26c2313.__obf_d129e62d54e19e72.Sub(__obf_5fb7c595c26c2313.__obf_d129e62d54e19e72, __obf_8266beb7fdfb4bd8)
	} else {
		__obf_5fb7c595c26c2313.__obf_d129e62d54e19e72.Add(__obf_5fb7c595c26c2313.__obf_d129e62d54e19e72, __obf_8266beb7fdfb4bd8)
	}

	// floor for positive numbers, ceil for negative numbers
	_, __obf_763dd0cca8ce709c := __obf_5fb7c595c26c2313.__obf_d129e62d54e19e72.DivMod(__obf_5fb7c595c26c2313.__obf_d129e62d54e19e72, __obf_db69efbd472ddd1a, new(big.Int))
	__obf_5fb7c595c26c2313.__obf_e060caa13549e714++
	if __obf_5fb7c595c26c2313.__obf_d129e62d54e19e72.Sign() < 0 && __obf_763dd0cca8ce709c.Cmp(__obf_22b4ed810afb99f1) != 0 {
		__obf_5fb7c595c26c2313.__obf_d129e62d54e19e72.Add(__obf_5fb7c595c26c2313.__obf_d129e62d54e19e72, __obf_3f4a47d8e0bfca14)
	}

	return __obf_5fb7c595c26c2313
}

// RoundCeil rounds the decimal towards +infinity.
//
// Example:
//
//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
func (__obf_45d3205e59c45c8b Decimal) RoundCeil(__obf_4f23fb2a2d558c71 int32) Decimal {
	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 >= -__obf_4f23fb2a2d558c71 {
		return __obf_45d3205e59c45c8b
	}

	__obf_38cd5438308df2cd := __obf_45d3205e59c45c8b.__obf_df6a3d24ae564ff2(-__obf_4f23fb2a2d558c71)
	if __obf_45d3205e59c45c8b.Equal(__obf_38cd5438308df2cd) {
		return __obf_45d3205e59c45c8b
	}

	if __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.Sign() > 0 {
		__obf_38cd5438308df2cd.__obf_d129e62d54e19e72.Add(__obf_38cd5438308df2cd.__obf_d129e62d54e19e72, __obf_3f4a47d8e0bfca14)
	}

	return __obf_38cd5438308df2cd
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_45d3205e59c45c8b Decimal) RoundFloor(__obf_4f23fb2a2d558c71 int32) Decimal {
	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 >= -__obf_4f23fb2a2d558c71 {
		return __obf_45d3205e59c45c8b
	}

	__obf_38cd5438308df2cd := __obf_45d3205e59c45c8b.__obf_df6a3d24ae564ff2(-__obf_4f23fb2a2d558c71)
	if __obf_45d3205e59c45c8b.Equal(__obf_38cd5438308df2cd) {
		return __obf_45d3205e59c45c8b
	}

	if __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.Sign() < 0 {
		__obf_38cd5438308df2cd.__obf_d129e62d54e19e72.Sub(__obf_38cd5438308df2cd.__obf_d129e62d54e19e72, __obf_3f4a47d8e0bfca14)
	}

	return __obf_38cd5438308df2cd
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_45d3205e59c45c8b Decimal) RoundUp(__obf_4f23fb2a2d558c71 int32) Decimal {
	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 >= -__obf_4f23fb2a2d558c71 {
		return __obf_45d3205e59c45c8b
	}

	__obf_38cd5438308df2cd := __obf_45d3205e59c45c8b.__obf_df6a3d24ae564ff2(-__obf_4f23fb2a2d558c71)
	if __obf_45d3205e59c45c8b.Equal(__obf_38cd5438308df2cd) {
		return __obf_45d3205e59c45c8b
	}

	if __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.Sign() > 0 {
		__obf_38cd5438308df2cd.__obf_d129e62d54e19e72.Add(__obf_38cd5438308df2cd.__obf_d129e62d54e19e72, __obf_3f4a47d8e0bfca14)
	} else if __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.Sign() < 0 {
		__obf_38cd5438308df2cd.__obf_d129e62d54e19e72.Sub(__obf_38cd5438308df2cd.__obf_d129e62d54e19e72, __obf_3f4a47d8e0bfca14)
	}

	return __obf_38cd5438308df2cd
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_45d3205e59c45c8b Decimal) RoundDown(__obf_4f23fb2a2d558c71 int32) Decimal {
	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 >= -__obf_4f23fb2a2d558c71 {
		return __obf_45d3205e59c45c8b
	}

	__obf_38cd5438308df2cd := __obf_45d3205e59c45c8b.__obf_df6a3d24ae564ff2(-__obf_4f23fb2a2d558c71)
	if __obf_45d3205e59c45c8b.Equal(__obf_38cd5438308df2cd) {
		return __obf_45d3205e59c45c8b
	}
	return __obf_38cd5438308df2cd
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
func (__obf_45d3205e59c45c8b Decimal) RoundBank(__obf_4f23fb2a2d558c71 int32) Decimal {

	__obf_242a336c43673540 := __obf_45d3205e59c45c8b.Round(__obf_4f23fb2a2d558c71)
	__obf_7eab4344a9e58361 := __obf_45d3205e59c45c8b.Sub(__obf_242a336c43673540).Abs()

	__obf_eaf3d488ce525b2f := New(5, -__obf_4f23fb2a2d558c71-1)
	if __obf_7eab4344a9e58361.Cmp(__obf_eaf3d488ce525b2f) == 0 && __obf_242a336c43673540.__obf_d129e62d54e19e72.Bit(0) != 0 {
		if __obf_242a336c43673540.__obf_d129e62d54e19e72.Sign() < 0 {
			__obf_242a336c43673540.__obf_d129e62d54e19e72.Add(__obf_242a336c43673540.__obf_d129e62d54e19e72, __obf_3f4a47d8e0bfca14)
		} else {
			__obf_242a336c43673540.__obf_d129e62d54e19e72.Sub(__obf_242a336c43673540.__obf_d129e62d54e19e72, __obf_3f4a47d8e0bfca14)
		}
	}

	return __obf_242a336c43673540
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
func (__obf_45d3205e59c45c8b Decimal) RoundCash(__obf_e8a18117b7dd2702 uint8) Decimal {
	var __obf_b59f61ce957d7b13 *big.Int
	switch __obf_e8a18117b7dd2702 {
	case 5:
		__obf_b59f61ce957d7b13 = __obf_b1624ddba33f3a22
	case 10:
		__obf_b59f61ce957d7b13 = __obf_db69efbd472ddd1a
	case 25:
		__obf_b59f61ce957d7b13 = __obf_9ae61dd5c563398c
	case 50:
		__obf_b59f61ce957d7b13 = __obf_c2d07968aecbaf08
	case 100:
		__obf_b59f61ce957d7b13 = __obf_3f4a47d8e0bfca14
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_e8a18117b7dd2702))
	}
	__obf_9bb402b5fd6b8e6d := Decimal{
		__obf_d129e62d54e19e72: __obf_b59f61ce957d7b13,
	}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_45d3205e59c45c8b.Mul(__obf_9bb402b5fd6b8e6d).Round(0).Div(__obf_9bb402b5fd6b8e6d).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_45d3205e59c45c8b Decimal) Floor() Decimal {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()

	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 >= 0 {
		return __obf_45d3205e59c45c8b
	}

	__obf_e060caa13549e714 := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_e060caa13549e714.Exp(__obf_e060caa13549e714, big.NewInt(-int64(__obf_45d3205e59c45c8b.__obf_e060caa13549e714)), nil)

	__obf_dd4d042995f3d916 := new(big.Int).Div(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72, __obf_e060caa13549e714)
	return Decimal{__obf_d129e62d54e19e72: __obf_dd4d042995f3d916, __obf_e060caa13549e714: 0}
}

// Ceil returns the nearest integer value greater than or equal to d.
func (__obf_45d3205e59c45c8b Decimal) Ceil() Decimal {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()

	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 >= 0 {
		return __obf_45d3205e59c45c8b
	}

	__obf_e060caa13549e714 := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_e060caa13549e714.Exp(__obf_e060caa13549e714, big.NewInt(-int64(__obf_45d3205e59c45c8b.__obf_e060caa13549e714)), nil)

	__obf_dd4d042995f3d916, __obf_763dd0cca8ce709c := new(big.Int).DivMod(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72, __obf_e060caa13549e714, new(big.Int))
	if __obf_763dd0cca8ce709c.Cmp(__obf_22b4ed810afb99f1) != 0 {
		__obf_dd4d042995f3d916.Add(__obf_dd4d042995f3d916, __obf_3f4a47d8e0bfca14)
	}
	return Decimal{__obf_d129e62d54e19e72: __obf_dd4d042995f3d916, __obf_e060caa13549e714: 0}
}

// Truncate truncates off digits from the number, without rounding.
//
// NOTE: precision is the last digit that will not be truncated (must be >= 0).
//
// Example:
//
//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
func (__obf_45d3205e59c45c8b Decimal) Truncate(__obf_11e1ee221b72ac06 int32) Decimal {
	__obf_45d3205e59c45c8b.__obf_f2008d48a41fad81()
	if __obf_11e1ee221b72ac06 >= 0 && -__obf_11e1ee221b72ac06 > __obf_45d3205e59c45c8b.__obf_e060caa13549e714 {
		return __obf_45d3205e59c45c8b.__obf_df6a3d24ae564ff2(-__obf_11e1ee221b72ac06)
	}
	return __obf_45d3205e59c45c8b
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_45d3205e59c45c8b *Decimal) UnmarshalJSON(__obf_4ff56fd714f33a1b []byte) error {
	if string(__obf_4ff56fd714f33a1b) == "null" {
		return nil
	}

	__obf_0905e1b7d0466bde, __obf_4ff66c6003689543 := __obf_e6592800dd967b05(__obf_4ff56fd714f33a1b)
	if __obf_4ff66c6003689543 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_4ff56fd714f33a1b, __obf_4ff66c6003689543)
	}

	__obf_210b94d2c4b23455, __obf_4ff66c6003689543 := NewFromString(__obf_0905e1b7d0466bde)
	*__obf_45d3205e59c45c8b = __obf_210b94d2c4b23455
	if __obf_4ff66c6003689543 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_0905e1b7d0466bde, __obf_4ff66c6003689543)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_45d3205e59c45c8b Decimal) MarshalJSON() ([]byte, error) {
	var __obf_0905e1b7d0466bde string
	if MarshalJSONWithoutQuotes {
		__obf_0905e1b7d0466bde = __obf_45d3205e59c45c8b.String()
	} else {
		__obf_0905e1b7d0466bde = "\"" + __obf_45d3205e59c45c8b.String() + "\""
	}
	return []byte(__obf_0905e1b7d0466bde), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_45d3205e59c45c8b *Decimal) UnmarshalBinary(__obf_5f9ac580dcfdd6a4 []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_5f9ac580dcfdd6a4) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_5f9ac580dcfdd6a4, len(__obf_5f9ac580dcfdd6a4))
	}

	// Extract the exponent
	__obf_45d3205e59c45c8b.__obf_e060caa13549e714 = int32(binary.BigEndian.Uint32(__obf_5f9ac580dcfdd6a4[:4]))

	// Extract the value
	__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72 = new(big.Int)
	if __obf_4ff66c6003689543 := __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.GobDecode(__obf_5f9ac580dcfdd6a4[4:]); __obf_4ff66c6003689543 != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_5f9ac580dcfdd6a4, __obf_4ff66c6003689543)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_45d3205e59c45c8b Decimal) MarshalBinary() (__obf_5f9ac580dcfdd6a4 []byte, __obf_4ff66c6003689543 error) {
	// exp is written first, but encode value first to know output size
	var __obf_b9436d0eaeeaa9e1 []byte
	if __obf_b9436d0eaeeaa9e1, __obf_4ff66c6003689543 = __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.GobEncode(); __obf_4ff66c6003689543 != nil {
		return nil, __obf_4ff66c6003689543
	}

	// Write the exponent in front, since it's a fixed size
	__obf_da202dba15e05015 := make([]byte, 4, len(__obf_b9436d0eaeeaa9e1)+4)
	binary.BigEndian.PutUint32(__obf_da202dba15e05015, uint32(__obf_45d3205e59c45c8b.__obf_e060caa13549e714))

	// Return the byte array
	return append(__obf_da202dba15e05015, __obf_b9436d0eaeeaa9e1...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_45d3205e59c45c8b *Decimal) Scan(__obf_d129e62d54e19e72 any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_6afdad92fd9f820c := __obf_d129e62d54e19e72.(type) {

	case float32:
		*__obf_45d3205e59c45c8b = NewFromFloat(float64(__obf_6afdad92fd9f820c))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_45d3205e59c45c8b = NewFromFloat(__obf_6afdad92fd9f820c)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_45d3205e59c45c8b = New(__obf_6afdad92fd9f820c, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_45d3205e59c45c8b = NewFromUint64(__obf_6afdad92fd9f820c)
		return nil

	default:
		// default is trying to interpret value stored as string
		__obf_0905e1b7d0466bde, __obf_4ff66c6003689543 := __obf_e6592800dd967b05(__obf_6afdad92fd9f820c)
		if __obf_4ff66c6003689543 != nil {
			return __obf_4ff66c6003689543
		}
		*__obf_45d3205e59c45c8b, __obf_4ff66c6003689543 = NewFromString(__obf_0905e1b7d0466bde)
		return __obf_4ff66c6003689543
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_45d3205e59c45c8b Decimal) Value() (driver.Value, error) {
	return __obf_45d3205e59c45c8b.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_45d3205e59c45c8b *Decimal) UnmarshalText(__obf_e5ae08f7a172c680 []byte) error {
	__obf_0905e1b7d0466bde := string(__obf_e5ae08f7a172c680)

	__obf_b99754db9573b7e4, __obf_4ff66c6003689543 := NewFromString(__obf_0905e1b7d0466bde)
	*__obf_45d3205e59c45c8b = __obf_b99754db9573b7e4
	if __obf_4ff66c6003689543 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_0905e1b7d0466bde, __obf_4ff66c6003689543)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_45d3205e59c45c8b Decimal) MarshalText() (__obf_e5ae08f7a172c680 []byte, __obf_4ff66c6003689543 error) {
	return []byte(__obf_45d3205e59c45c8b.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_45d3205e59c45c8b Decimal) GobEncode() ([]byte, error) {
	return __obf_45d3205e59c45c8b.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_45d3205e59c45c8b *Decimal) GobDecode(__obf_5f9ac580dcfdd6a4 []byte) error {
	return __obf_45d3205e59c45c8b.UnmarshalBinary(__obf_5f9ac580dcfdd6a4)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_45d3205e59c45c8b Decimal) StringScaled(__obf_e060caa13549e714 int32) string {
	return __obf_45d3205e59c45c8b.__obf_df6a3d24ae564ff2(__obf_e060caa13549e714).String()
}

func (__obf_45d3205e59c45c8b Decimal) string(__obf_a08a038e5f1b7693 bool) string {
	if __obf_45d3205e59c45c8b.__obf_e060caa13549e714 >= 0 {
		return __obf_45d3205e59c45c8b.__obf_df6a3d24ae564ff2(0).__obf_d129e62d54e19e72.String()
	}

	__obf_c318d0106586e799 := new(big.Int).Abs(__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72)
	__obf_0905e1b7d0466bde := __obf_c318d0106586e799.String()

	var __obf_7a6acf8df2e65feb, __obf_6909e7d2f53445c5 string

	// NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
	// and you are on a 32-bit machine. Won't fix this super-edge case.
	__obf_5b7a9b9dda71f5d2 := int(__obf_45d3205e59c45c8b.__obf_e060caa13549e714)
	if len(__obf_0905e1b7d0466bde) > -__obf_5b7a9b9dda71f5d2 {
		__obf_7a6acf8df2e65feb = __obf_0905e1b7d0466bde[:len(__obf_0905e1b7d0466bde)+__obf_5b7a9b9dda71f5d2]
		__obf_6909e7d2f53445c5 = __obf_0905e1b7d0466bde[len(__obf_0905e1b7d0466bde)+__obf_5b7a9b9dda71f5d2:]
	} else {
		__obf_7a6acf8df2e65feb = "0"

		__obf_2b725ae37c345bb0 := -__obf_5b7a9b9dda71f5d2 - len(__obf_0905e1b7d0466bde)
		__obf_6909e7d2f53445c5 = strings.Repeat("0", __obf_2b725ae37c345bb0) + __obf_0905e1b7d0466bde
	}

	if __obf_a08a038e5f1b7693 {
		__obf_dec02b860bfb4fdd := len(__obf_6909e7d2f53445c5) - 1
		for ; __obf_dec02b860bfb4fdd >= 0; __obf_dec02b860bfb4fdd-- {
			if __obf_6909e7d2f53445c5[__obf_dec02b860bfb4fdd] != '0' {
				break
			}
		}
		__obf_6909e7d2f53445c5 = __obf_6909e7d2f53445c5[:__obf_dec02b860bfb4fdd+1]
	}

	__obf_5060b255dcae8b47 := __obf_7a6acf8df2e65feb
	if len(__obf_6909e7d2f53445c5) > 0 {
		__obf_5060b255dcae8b47 += "." + __obf_6909e7d2f53445c5
	}

	if __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72.Sign() < 0 {
		return "-" + __obf_5060b255dcae8b47
	}

	return __obf_5060b255dcae8b47
}

func (__obf_45d3205e59c45c8b *Decimal) __obf_f2008d48a41fad81() {
	if __obf_45d3205e59c45c8b.__obf_d129e62d54e19e72 == nil {
		__obf_45d3205e59c45c8b.__obf_d129e62d54e19e72 = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_e3d00598af30d024 Decimal, __obf_7c513b37dc1fad45 ...Decimal) Decimal {
	__obf_9df656b40587d3b6 := __obf_e3d00598af30d024
	for _, __obf_0ec9adcbd207742f := range __obf_7c513b37dc1fad45 {
		if __obf_0ec9adcbd207742f.Cmp(__obf_9df656b40587d3b6) < 0 {
			__obf_9df656b40587d3b6 = __obf_0ec9adcbd207742f
		}
	}
	return __obf_9df656b40587d3b6
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_e3d00598af30d024 Decimal, __obf_7c513b37dc1fad45 ...Decimal) Decimal {
	__obf_9df656b40587d3b6 := __obf_e3d00598af30d024
	for _, __obf_0ec9adcbd207742f := range __obf_7c513b37dc1fad45 {
		if __obf_0ec9adcbd207742f.Cmp(__obf_9df656b40587d3b6) > 0 {
			__obf_9df656b40587d3b6 = __obf_0ec9adcbd207742f
		}
	}
	return __obf_9df656b40587d3b6
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_e3d00598af30d024 Decimal, __obf_7c513b37dc1fad45 ...Decimal) Decimal {
	__obf_a93a63026301ead5 := __obf_e3d00598af30d024
	for _, __obf_0ec9adcbd207742f := range __obf_7c513b37dc1fad45 {
		__obf_a93a63026301ead5 = __obf_a93a63026301ead5.Add(__obf_0ec9adcbd207742f)
	}

	return __obf_a93a63026301ead5
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_e3d00598af30d024 Decimal, __obf_7c513b37dc1fad45 ...Decimal) Decimal {
	__obf_30bbd3baeeb5ee19 := New(int64(len(__obf_7c513b37dc1fad45)+1), 0)
	__obf_f26c5cf9e7a07902 := Sum(__obf_e3d00598af30d024, __obf_7c513b37dc1fad45...)
	return __obf_f26c5cf9e7a07902.Div(__obf_30bbd3baeeb5ee19)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_59df625c6740d736 Decimal, __obf_0cf0ba5fcabed996 Decimal) (Decimal, Decimal) {
	__obf_59df625c6740d736.__obf_f2008d48a41fad81()
	__obf_0cf0ba5fcabed996.__obf_f2008d48a41fad81()

	if __obf_59df625c6740d736.__obf_e060caa13549e714 < __obf_0cf0ba5fcabed996.__obf_e060caa13549e714 {
		return __obf_59df625c6740d736, __obf_0cf0ba5fcabed996.__obf_df6a3d24ae564ff2(__obf_59df625c6740d736.__obf_e060caa13549e714)
	} else if __obf_59df625c6740d736.__obf_e060caa13549e714 > __obf_0cf0ba5fcabed996.__obf_e060caa13549e714 {
		return __obf_59df625c6740d736.__obf_df6a3d24ae564ff2(__obf_0cf0ba5fcabed996.__obf_e060caa13549e714), __obf_0cf0ba5fcabed996
	}

	return __obf_59df625c6740d736, __obf_0cf0ba5fcabed996
}

func __obf_e6592800dd967b05(__obf_d129e62d54e19e72 any) (string, error) {
	var __obf_dc396c66542e829b []byte

	switch __obf_6afdad92fd9f820c := __obf_d129e62d54e19e72.(type) {
	case string:
		__obf_dc396c66542e829b = []byte(__obf_6afdad92fd9f820c)
	case []byte:
		__obf_dc396c66542e829b = __obf_6afdad92fd9f820c
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_d129e62d54e19e72, __obf_d129e62d54e19e72)
	}

	// If the amount is quoted, strip the quotes
	if len(__obf_dc396c66542e829b) > 2 && __obf_dc396c66542e829b[0] == '"' && __obf_dc396c66542e829b[len(__obf_dc396c66542e829b)-1] == '"' {
		__obf_dc396c66542e829b = __obf_dc396c66542e829b[1 : len(__obf_dc396c66542e829b)-1]
	}
	return string(__obf_dc396c66542e829b), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the datautil.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_45d3205e59c45c8b Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_45d3205e59c45c8b,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_45d3205e59c45c8b *NullDecimal) Scan(__obf_d129e62d54e19e72 any) error {
	if __obf_d129e62d54e19e72 == nil {
		__obf_45d3205e59c45c8b.Valid = false
		return nil
	}
	__obf_45d3205e59c45c8b.Valid = true
	return __obf_45d3205e59c45c8b.Decimal.Scan(__obf_d129e62d54e19e72)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_45d3205e59c45c8b NullDecimal) Value() (driver.Value, error) {
	if !__obf_45d3205e59c45c8b.Valid {
		return nil, nil
	}
	return __obf_45d3205e59c45c8b.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_45d3205e59c45c8b *NullDecimal) UnmarshalJSON(__obf_4ff56fd714f33a1b []byte) error {
	if string(__obf_4ff56fd714f33a1b) == "null" {
		__obf_45d3205e59c45c8b.Valid = false
		return nil
	}
	__obf_45d3205e59c45c8b.Valid = true
	return __obf_45d3205e59c45c8b.Decimal.UnmarshalJSON(__obf_4ff56fd714f33a1b)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_45d3205e59c45c8b NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_45d3205e59c45c8b.Valid {
		return []byte("null"), nil
	}
	return __obf_45d3205e59c45c8b.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_45d3205e59c45c8b *NullDecimal) UnmarshalText(__obf_e5ae08f7a172c680 []byte) error {
	__obf_0905e1b7d0466bde := string(__obf_e5ae08f7a172c680)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_0905e1b7d0466bde == "" {
		__obf_45d3205e59c45c8b.Valid = false
		return nil
	}
	if __obf_4ff66c6003689543 := __obf_45d3205e59c45c8b.Decimal.UnmarshalText(__obf_e5ae08f7a172c680); __obf_4ff66c6003689543 != nil {
		__obf_45d3205e59c45c8b.Valid = false
		return __obf_4ff66c6003689543
	}
	__obf_45d3205e59c45c8b.Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_45d3205e59c45c8b NullDecimal) MarshalText() (__obf_e5ae08f7a172c680 []byte, __obf_4ff66c6003689543 error) {
	if !__obf_45d3205e59c45c8b.Valid {
		return []byte{}, nil
	}
	return __obf_45d3205e59c45c8b.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_45d3205e59c45c8b Decimal) Atan() Decimal {
	if __obf_45d3205e59c45c8b.Equal(NewFromFloat(0.0)) {
		return __obf_45d3205e59c45c8b
	}
	if __obf_45d3205e59c45c8b.GreaterThan(NewFromFloat(0.0)) {
		return __obf_45d3205e59c45c8b.__obf_29694f48ac7f39e5()
	}
	return __obf_45d3205e59c45c8b.Neg().__obf_29694f48ac7f39e5().Neg()
}

func (__obf_45d3205e59c45c8b Decimal) __obf_ac5e82a8ed97f7b7() Decimal {
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
	__obf_dd4d042995f3d916 := __obf_45d3205e59c45c8b.Mul(__obf_45d3205e59c45c8b)
	__obf_afc9b2db13b64d6e := P0.Mul(__obf_dd4d042995f3d916).Add(P1).Mul(__obf_dd4d042995f3d916).Add(P2).Mul(__obf_dd4d042995f3d916).Add(P3).Mul(__obf_dd4d042995f3d916).Add(P4).Mul(__obf_dd4d042995f3d916)
	__obf_1de020e196dfce0d := __obf_dd4d042995f3d916.Add(Q0).Mul(__obf_dd4d042995f3d916).Add(Q1).Mul(__obf_dd4d042995f3d916).Add(Q2).Mul(__obf_dd4d042995f3d916).Add(Q3).Mul(__obf_dd4d042995f3d916).Add(Q4)
	__obf_dd4d042995f3d916 = __obf_afc9b2db13b64d6e.Div(__obf_1de020e196dfce0d)
	__obf_dd4d042995f3d916 = __obf_45d3205e59c45c8b.Mul(__obf_dd4d042995f3d916).Add(__obf_45d3205e59c45c8b)
	return __obf_dd4d042995f3d916
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_45d3205e59c45c8b Decimal) __obf_29694f48ac7f39e5() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)      // tan(3*pi/8)
	__obf_e9292a70790fc77e := NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_45d3205e59c45c8b.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_45d3205e59c45c8b.__obf_ac5e82a8ed97f7b7()
	}
	if __obf_45d3205e59c45c8b.GreaterThan(Tan3pio8) {
		return __obf_e9292a70790fc77e.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_45d3205e59c45c8b).__obf_ac5e82a8ed97f7b7()).Add(Morebits)
	}
	return __obf_e9292a70790fc77e.Div(NewFromFloat(4.0)).Add((__obf_45d3205e59c45c8b.Sub(NewFromFloat(1.0)).Div(__obf_45d3205e59c45c8b.Add(NewFromFloat(1.0)))).__obf_ac5e82a8ed97f7b7()).Add(NewFromFloat(0.5).Mul(Morebits))
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
func (__obf_45d3205e59c45c8b Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_45d3205e59c45c8b.Equal(NewFromFloat(0.0)) {
		return __obf_45d3205e59c45c8b
	}
	// make argument positive but save the sign
	__obf_8c897683b6557fb7 := false
	if __obf_45d3205e59c45c8b.LessThan(NewFromFloat(0.0)) {
		__obf_45d3205e59c45c8b = __obf_45d3205e59c45c8b.Neg()
		__obf_8c897683b6557fb7 = true
	}

	__obf_37be0a5811b831ee := __obf_45d3205e59c45c8b.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_877f288b62c7ab1d := NewFromFloat(float64(__obf_37be0a5811b831ee)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_37be0a5811b831ee&1 == 1 {
		__obf_37be0a5811b831ee++
		__obf_877f288b62c7ab1d = __obf_877f288b62c7ab1d.Add(NewFromFloat(1.0))
	}
	__obf_37be0a5811b831ee &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_37be0a5811b831ee > 3 {
		__obf_8c897683b6557fb7 = !__obf_8c897683b6557fb7
		__obf_37be0a5811b831ee -= 4
	}
	__obf_dd4d042995f3d916 := __obf_45d3205e59c45c8b.Sub(__obf_877f288b62c7ab1d.Mul(PI4A)).Sub(__obf_877f288b62c7ab1d.Mul(PI4B)).Sub(__obf_877f288b62c7ab1d.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_cdd5b637faef6146 := __obf_dd4d042995f3d916.Mul(__obf_dd4d042995f3d916)

	if __obf_37be0a5811b831ee == 1 || __obf_37be0a5811b831ee == 2 {
		__obf_b6c7f3c8efed3bad := __obf_cdd5b637faef6146.Mul(__obf_cdd5b637faef6146).Mul(_cos[0].Mul(__obf_cdd5b637faef6146).Add(_cos[1]).Mul(__obf_cdd5b637faef6146).Add(_cos[2]).Mul(__obf_cdd5b637faef6146).Add(_cos[3]).Mul(__obf_cdd5b637faef6146).Add(_cos[4]).Mul(__obf_cdd5b637faef6146).Add(_cos[5]))
		__obf_877f288b62c7ab1d = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_cdd5b637faef6146)).Add(__obf_b6c7f3c8efed3bad)
	} else {
		__obf_877f288b62c7ab1d = __obf_dd4d042995f3d916.Add(__obf_dd4d042995f3d916.Mul(__obf_cdd5b637faef6146).Mul(_sin[0].Mul(__obf_cdd5b637faef6146).Add(_sin[1]).Mul(__obf_cdd5b637faef6146).Add(_sin[2]).Mul(__obf_cdd5b637faef6146).Add(_sin[3]).Mul(__obf_cdd5b637faef6146).Add(_sin[4]).Mul(__obf_cdd5b637faef6146).Add(_sin[5])))
	}
	if __obf_8c897683b6557fb7 {
		__obf_877f288b62c7ab1d = __obf_877f288b62c7ab1d.Neg()
	}
	return __obf_877f288b62c7ab1d
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
func (__obf_45d3205e59c45c8b Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	// make argument positive
	__obf_8c897683b6557fb7 := false
	if __obf_45d3205e59c45c8b.LessThan(NewFromFloat(0.0)) {
		__obf_45d3205e59c45c8b = __obf_45d3205e59c45c8b.Neg()
	}

	__obf_37be0a5811b831ee := __obf_45d3205e59c45c8b.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_877f288b62c7ab1d := NewFromFloat(float64(__obf_37be0a5811b831ee)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_37be0a5811b831ee&1 == 1 {
		__obf_37be0a5811b831ee++
		__obf_877f288b62c7ab1d = __obf_877f288b62c7ab1d.Add(NewFromFloat(1.0))
	}
	__obf_37be0a5811b831ee &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_37be0a5811b831ee > 3 {
		__obf_8c897683b6557fb7 = !__obf_8c897683b6557fb7
		__obf_37be0a5811b831ee -= 4
	}
	if __obf_37be0a5811b831ee > 1 {
		__obf_8c897683b6557fb7 = !__obf_8c897683b6557fb7
	}

	__obf_dd4d042995f3d916 := __obf_45d3205e59c45c8b.Sub(__obf_877f288b62c7ab1d.Mul(PI4A)).Sub(__obf_877f288b62c7ab1d.Mul(PI4B)).Sub(__obf_877f288b62c7ab1d.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_cdd5b637faef6146 := __obf_dd4d042995f3d916.Mul(__obf_dd4d042995f3d916)

	if __obf_37be0a5811b831ee == 1 || __obf_37be0a5811b831ee == 2 {
		__obf_877f288b62c7ab1d = __obf_dd4d042995f3d916.Add(__obf_dd4d042995f3d916.Mul(__obf_cdd5b637faef6146).Mul(_sin[0].Mul(__obf_cdd5b637faef6146).Add(_sin[1]).Mul(__obf_cdd5b637faef6146).Add(_sin[2]).Mul(__obf_cdd5b637faef6146).Add(_sin[3]).Mul(__obf_cdd5b637faef6146).Add(_sin[4]).Mul(__obf_cdd5b637faef6146).Add(_sin[5])))
	} else {
		__obf_b6c7f3c8efed3bad := __obf_cdd5b637faef6146.Mul(__obf_cdd5b637faef6146).Mul(_cos[0].Mul(__obf_cdd5b637faef6146).Add(_cos[1]).Mul(__obf_cdd5b637faef6146).Add(_cos[2]).Mul(__obf_cdd5b637faef6146).Add(_cos[3]).Mul(__obf_cdd5b637faef6146).Add(_cos[4]).Mul(__obf_cdd5b637faef6146).Add(_cos[5]))
		__obf_877f288b62c7ab1d = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_cdd5b637faef6146)).Add(__obf_b6c7f3c8efed3bad)
	}
	if __obf_8c897683b6557fb7 {
		__obf_877f288b62c7ab1d = __obf_877f288b62c7ab1d.Neg()
	}
	return __obf_877f288b62c7ab1d
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
func (__obf_45d3205e59c45c8b Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_45d3205e59c45c8b.Equal(NewFromFloat(0.0)) {
		return __obf_45d3205e59c45c8b
	}

	// make argument positive but save the sign
	__obf_8c897683b6557fb7 := false
	if __obf_45d3205e59c45c8b.LessThan(NewFromFloat(0.0)) {
		__obf_45d3205e59c45c8b = __obf_45d3205e59c45c8b.Neg()
		__obf_8c897683b6557fb7 = true
	}

	__obf_37be0a5811b831ee := __obf_45d3205e59c45c8b.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_877f288b62c7ab1d := NewFromFloat(float64(__obf_37be0a5811b831ee)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_37be0a5811b831ee&1 == 1 {
		__obf_37be0a5811b831ee++
		__obf_877f288b62c7ab1d = __obf_877f288b62c7ab1d.Add(NewFromFloat(1.0))
	}

	__obf_dd4d042995f3d916 := __obf_45d3205e59c45c8b.Sub(__obf_877f288b62c7ab1d.Mul(PI4A)).Sub(__obf_877f288b62c7ab1d.Mul(PI4B)).Sub(__obf_877f288b62c7ab1d.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_cdd5b637faef6146 := __obf_dd4d042995f3d916.Mul(__obf_dd4d042995f3d916)

	if __obf_cdd5b637faef6146.GreaterThan(NewFromFloat(1e-14)) {
		__obf_b6c7f3c8efed3bad := __obf_cdd5b637faef6146.Mul(_tanP[0].Mul(__obf_cdd5b637faef6146).Add(_tanP[1]).Mul(__obf_cdd5b637faef6146).Add(_tanP[2]))
		__obf_4210667d93928829 := __obf_cdd5b637faef6146.Add(_tanQ[1]).Mul(__obf_cdd5b637faef6146).Add(_tanQ[2]).Mul(__obf_cdd5b637faef6146).Add(_tanQ[3]).Mul(__obf_cdd5b637faef6146).Add(_tanQ[4])
		__obf_877f288b62c7ab1d = __obf_dd4d042995f3d916.Add(__obf_dd4d042995f3d916.Mul(__obf_b6c7f3c8efed3bad.Div(__obf_4210667d93928829)))
	} else {
		__obf_877f288b62c7ab1d = __obf_dd4d042995f3d916
	}
	if __obf_37be0a5811b831ee&2 == 2 {
		__obf_877f288b62c7ab1d = NewFromFloat(-1.0).Div(__obf_877f288b62c7ab1d)
	}
	if __obf_8c897683b6557fb7 {
		__obf_877f288b62c7ab1d = __obf_877f288b62c7ab1d.Neg()
	}
	return __obf_877f288b62c7ab1d
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

type __obf_210b94d2c4b23455 struct {
	__obf_45d3205e59c45c8b [800]byte // digits, big-endian representation
	__obf_2298ab2cdf3cb6bd int       // number of digits used
	__obf_44d25d407353f76c int       // decimal point
	__obf_4ae2cd7e4db15eb2 bool      // negative flag
	__obf_b5f84f6e6f423be6 bool      // discarded nonzero digits beyond d[:nd]
}

func (__obf_a7affdfda169fd03 *__obf_210b94d2c4b23455) String() string {
	__obf_46e3648cfcdd2b28 := 10 + __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd
	if __obf_a7affdfda169fd03.__obf_44d25d407353f76c > 0 {
		__obf_46e3648cfcdd2b28 += __obf_a7affdfda169fd03.__obf_44d25d407353f76c
	}
	if __obf_a7affdfda169fd03.__obf_44d25d407353f76c < 0 {
		__obf_46e3648cfcdd2b28 += -__obf_a7affdfda169fd03.__obf_44d25d407353f76c
	}

	__obf_ef174b7d24f6be30 := make([]byte, __obf_46e3648cfcdd2b28)
	__obf_b6c7f3c8efed3bad := 0
	switch {
	case __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd == 0:
		return "0"

	case __obf_a7affdfda169fd03.__obf_44d25d407353f76c <= 0:
		// zeros fill space between decimal point and digits
		__obf_ef174b7d24f6be30[__obf_b6c7f3c8efed3bad] = '0'
		__obf_b6c7f3c8efed3bad++
		__obf_ef174b7d24f6be30[__obf_b6c7f3c8efed3bad] = '.'
		__obf_b6c7f3c8efed3bad++
		__obf_b6c7f3c8efed3bad += __obf_4045c456a66d75e7(__obf_ef174b7d24f6be30[__obf_b6c7f3c8efed3bad : __obf_b6c7f3c8efed3bad+-__obf_a7affdfda169fd03.__obf_44d25d407353f76c])
		__obf_b6c7f3c8efed3bad += copy(__obf_ef174b7d24f6be30[__obf_b6c7f3c8efed3bad:], __obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[0:__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd])

	case __obf_a7affdfda169fd03.__obf_44d25d407353f76c < __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd:
		// decimal point in middle of digits
		__obf_b6c7f3c8efed3bad += copy(__obf_ef174b7d24f6be30[__obf_b6c7f3c8efed3bad:], __obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[0:__obf_a7affdfda169fd03.__obf_44d25d407353f76c])
		__obf_ef174b7d24f6be30[__obf_b6c7f3c8efed3bad] = '.'
		__obf_b6c7f3c8efed3bad++
		__obf_b6c7f3c8efed3bad += copy(__obf_ef174b7d24f6be30[__obf_b6c7f3c8efed3bad:], __obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_a7affdfda169fd03.__obf_44d25d407353f76c:__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd])

	default:
		// zeros fill space between digits and decimal point
		__obf_b6c7f3c8efed3bad += copy(__obf_ef174b7d24f6be30[__obf_b6c7f3c8efed3bad:], __obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[0:__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd])
		__obf_b6c7f3c8efed3bad += __obf_4045c456a66d75e7(__obf_ef174b7d24f6be30[__obf_b6c7f3c8efed3bad : __obf_b6c7f3c8efed3bad+__obf_a7affdfda169fd03.__obf_44d25d407353f76c-__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd])
	}
	return string(__obf_ef174b7d24f6be30[0:__obf_b6c7f3c8efed3bad])
}

func __obf_4045c456a66d75e7(__obf_f5ecaf3aa26075a6 []byte) int {
	for __obf_dec02b860bfb4fdd := range __obf_f5ecaf3aa26075a6 {
		__obf_f5ecaf3aa26075a6[__obf_dec02b860bfb4fdd] = '0'
	}
	return len(__obf_f5ecaf3aa26075a6)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_54823070883b28a6(__obf_a7affdfda169fd03 *__obf_210b94d2c4b23455) {
	for __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd > 0 && __obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd-1] == '0' {
		__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd--
	}
	if __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd == 0 {
		__obf_a7affdfda169fd03.__obf_44d25d407353f76c = 0
	}
}

// Assign v to a.
func (__obf_a7affdfda169fd03 *__obf_210b94d2c4b23455) Assign(__obf_6afdad92fd9f820c uint64) {
	var __obf_ef174b7d24f6be30 [24]byte

	// Write reversed decimal in buf.
	__obf_46e3648cfcdd2b28 := 0
	for __obf_6afdad92fd9f820c > 0 {
		__obf_afd170dcd2204738 := __obf_6afdad92fd9f820c / 10
		__obf_6afdad92fd9f820c -= 10 * __obf_afd170dcd2204738
		__obf_ef174b7d24f6be30[__obf_46e3648cfcdd2b28] = byte(__obf_6afdad92fd9f820c + '0')
		__obf_46e3648cfcdd2b28++
		__obf_6afdad92fd9f820c = __obf_afd170dcd2204738
	}

	// Reverse again to produce forward decimal in a.d.
	__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd = 0
	for __obf_46e3648cfcdd2b28--; __obf_46e3648cfcdd2b28 >= 0; __obf_46e3648cfcdd2b28-- {
		__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd] = __obf_ef174b7d24f6be30[__obf_46e3648cfcdd2b28]
		__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd++
	}
	__obf_a7affdfda169fd03.__obf_44d25d407353f76c = __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd
	__obf_54823070883b28a6(__obf_a7affdfda169fd03)
}

// Maximum shift that we can do in one pass without overflow.
// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
const __obf_754f7ac97e2e4958 = 32 << (^uint(0) >> 63)
const __obf_989e62b587026dc9 = __obf_754f7ac97e2e4958 - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_655d7c295cab5a7b(__obf_a7affdfda169fd03 *__obf_210b94d2c4b23455, __obf_8431217a5a792855 uint) {
	__obf_8bfc98156daaf406 := 0 // read pointer
	__obf_b6c7f3c8efed3bad := 0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_46e3648cfcdd2b28 uint
	for ; __obf_46e3648cfcdd2b28>>__obf_8431217a5a792855 == 0; __obf_8bfc98156daaf406++ {
		if __obf_8bfc98156daaf406 >= __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd {
			if __obf_46e3648cfcdd2b28 == 0 {
				// a == 0; shouldn't get here, but handle anyway.
				__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd = 0
				return
			}
			for __obf_46e3648cfcdd2b28>>__obf_8431217a5a792855 == 0 {
				__obf_46e3648cfcdd2b28 = __obf_46e3648cfcdd2b28 * 10
				__obf_8bfc98156daaf406++
			}
			break
		}
		__obf_820da776a4f8225c := uint(__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_8bfc98156daaf406])
		__obf_46e3648cfcdd2b28 = __obf_46e3648cfcdd2b28*10 + __obf_820da776a4f8225c - '0'
	}
	__obf_a7affdfda169fd03.__obf_44d25d407353f76c -= __obf_8bfc98156daaf406 - 1

	var __obf_60932481bc80ca9b uint = (1 << __obf_8431217a5a792855) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_8bfc98156daaf406 < __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd; __obf_8bfc98156daaf406++ {
		__obf_820da776a4f8225c := uint(__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_8bfc98156daaf406])
		__obf_2cac2e2423ea4c13 := __obf_46e3648cfcdd2b28 >> __obf_8431217a5a792855
		__obf_46e3648cfcdd2b28 &= __obf_60932481bc80ca9b
		__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_b6c7f3c8efed3bad] = byte(__obf_2cac2e2423ea4c13 + '0')
		__obf_b6c7f3c8efed3bad++
		__obf_46e3648cfcdd2b28 = __obf_46e3648cfcdd2b28*10 + __obf_820da776a4f8225c - '0'
	}

	// Put down extra digits.
	for __obf_46e3648cfcdd2b28 > 0 {
		__obf_2cac2e2423ea4c13 := __obf_46e3648cfcdd2b28 >> __obf_8431217a5a792855
		__obf_46e3648cfcdd2b28 &= __obf_60932481bc80ca9b
		if __obf_b6c7f3c8efed3bad < len(__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b) {
			__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_b6c7f3c8efed3bad] = byte(__obf_2cac2e2423ea4c13 + '0')
			__obf_b6c7f3c8efed3bad++
		} else if __obf_2cac2e2423ea4c13 > 0 {
			__obf_a7affdfda169fd03.__obf_b5f84f6e6f423be6 = true
		}
		__obf_46e3648cfcdd2b28 = __obf_46e3648cfcdd2b28 * 10
	}

	__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd = __obf_b6c7f3c8efed3bad
	__obf_54823070883b28a6(__obf_a7affdfda169fd03)
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

type __obf_6115c10bcdac26b6 struct {
	__obf_8eac66a2db598c7e int    // number of new digits
	__obf_e4b1dd1fbe0b0b36 string // minus one digit if original < a.
}

var __obf_3cd9f25374bd0e24 = []__obf_6115c10bcdac26b6{
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
func __obf_fbdf3ddf250b1512(__obf_5366d320659b7735 []byte, __obf_5c83bf0ab27310a3 string) bool {
	for __obf_dec02b860bfb4fdd := 0; __obf_dec02b860bfb4fdd < len(__obf_5c83bf0ab27310a3); __obf_dec02b860bfb4fdd++ {
		if __obf_dec02b860bfb4fdd >= len(__obf_5366d320659b7735) {
			return true
		}
		if __obf_5366d320659b7735[__obf_dec02b860bfb4fdd] != __obf_5c83bf0ab27310a3[__obf_dec02b860bfb4fdd] {
			return __obf_5366d320659b7735[__obf_dec02b860bfb4fdd] < __obf_5c83bf0ab27310a3[__obf_dec02b860bfb4fdd]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_1a4a04e60c1d5b28(__obf_a7affdfda169fd03 *__obf_210b94d2c4b23455, __obf_8431217a5a792855 uint) {
	__obf_8eac66a2db598c7e := __obf_3cd9f25374bd0e24[__obf_8431217a5a792855].__obf_8eac66a2db598c7e
	if __obf_fbdf3ddf250b1512(__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[0:__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd], __obf_3cd9f25374bd0e24[__obf_8431217a5a792855].__obf_e4b1dd1fbe0b0b36) {
		__obf_8eac66a2db598c7e--
	}

	__obf_8bfc98156daaf406 := __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd                          // read index
	__obf_b6c7f3c8efed3bad := __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd + __obf_8eac66a2db598c7e // write index

	// Pick up a digit, put down a digit.
	var __obf_46e3648cfcdd2b28 uint
	for __obf_8bfc98156daaf406--; __obf_8bfc98156daaf406 >= 0; __obf_8bfc98156daaf406-- {
		__obf_46e3648cfcdd2b28 += (uint(__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_8bfc98156daaf406]) - '0') << __obf_8431217a5a792855
		__obf_e0bc07dee70b2943 := __obf_46e3648cfcdd2b28 / 10
		__obf_9de558a3a25b0f69 := __obf_46e3648cfcdd2b28 - 10*__obf_e0bc07dee70b2943
		__obf_b6c7f3c8efed3bad--
		if __obf_b6c7f3c8efed3bad < len(__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b) {
			__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_b6c7f3c8efed3bad] = byte(__obf_9de558a3a25b0f69 + '0')
		} else if __obf_9de558a3a25b0f69 != 0 {
			__obf_a7affdfda169fd03.__obf_b5f84f6e6f423be6 = true
		}
		__obf_46e3648cfcdd2b28 = __obf_e0bc07dee70b2943
	}

	// Put down extra digits.
	for __obf_46e3648cfcdd2b28 > 0 {
		__obf_e0bc07dee70b2943 := __obf_46e3648cfcdd2b28 / 10
		__obf_9de558a3a25b0f69 := __obf_46e3648cfcdd2b28 - 10*__obf_e0bc07dee70b2943
		__obf_b6c7f3c8efed3bad--
		if __obf_b6c7f3c8efed3bad < len(__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b) {
			__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_b6c7f3c8efed3bad] = byte(__obf_9de558a3a25b0f69 + '0')
		} else if __obf_9de558a3a25b0f69 != 0 {
			__obf_a7affdfda169fd03.__obf_b5f84f6e6f423be6 = true
		}
		__obf_46e3648cfcdd2b28 = __obf_e0bc07dee70b2943
	}

	__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd += __obf_8eac66a2db598c7e
	if __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd >= len(__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b) {
		__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd = len(__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b)
	}
	__obf_a7affdfda169fd03.__obf_44d25d407353f76c += __obf_8eac66a2db598c7e
	__obf_54823070883b28a6(__obf_a7affdfda169fd03)
}

// Binary shift left (k > 0) or right (k < 0).
func (__obf_a7affdfda169fd03 *__obf_210b94d2c4b23455) Shift(__obf_8431217a5a792855 int) {
	switch {
	case __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd == 0:
		// nothing to do: a == 0
	case __obf_8431217a5a792855 > 0:
		for __obf_8431217a5a792855 > __obf_989e62b587026dc9 {
			__obf_1a4a04e60c1d5b28(__obf_a7affdfda169fd03, __obf_989e62b587026dc9)
			__obf_8431217a5a792855 -= __obf_989e62b587026dc9
		}
		__obf_1a4a04e60c1d5b28(__obf_a7affdfda169fd03, uint(__obf_8431217a5a792855))
	case __obf_8431217a5a792855 < 0:
		for __obf_8431217a5a792855 < -__obf_989e62b587026dc9 {
			__obf_655d7c295cab5a7b(__obf_a7affdfda169fd03, __obf_989e62b587026dc9)
			__obf_8431217a5a792855 += __obf_989e62b587026dc9
		}
		__obf_655d7c295cab5a7b(__obf_a7affdfda169fd03, uint(-__obf_8431217a5a792855))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_e61bee8b27eb8a58(__obf_a7affdfda169fd03 *__obf_210b94d2c4b23455, __obf_2298ab2cdf3cb6bd int) bool {
	if __obf_2298ab2cdf3cb6bd < 0 || __obf_2298ab2cdf3cb6bd >= __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd {
		return false
	}
	if __obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_2298ab2cdf3cb6bd] == '5' && __obf_2298ab2cdf3cb6bd+1 == __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd { // exactly halfway - round to even
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_a7affdfda169fd03.__obf_b5f84f6e6f423be6 {
			return true
		}
		return __obf_2298ab2cdf3cb6bd > 0 && (__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_2298ab2cdf3cb6bd-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_2298ab2cdf3cb6bd] >= '5'
}

// Round a to nd digits (or fewer).
// If nd is zero, it means we're rounding
// just to the left of the digits, as in
// 0.09 -> 0.1.
func (__obf_a7affdfda169fd03 *__obf_210b94d2c4b23455) Round(__obf_2298ab2cdf3cb6bd int) {
	if __obf_2298ab2cdf3cb6bd < 0 || __obf_2298ab2cdf3cb6bd >= __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd {
		return
	}
	if __obf_e61bee8b27eb8a58(__obf_a7affdfda169fd03, __obf_2298ab2cdf3cb6bd) {
		__obf_a7affdfda169fd03.RoundUp(__obf_2298ab2cdf3cb6bd)
	} else {
		__obf_a7affdfda169fd03.RoundDown(__obf_2298ab2cdf3cb6bd)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_a7affdfda169fd03 *__obf_210b94d2c4b23455) RoundDown(__obf_2298ab2cdf3cb6bd int) {
	if __obf_2298ab2cdf3cb6bd < 0 || __obf_2298ab2cdf3cb6bd >= __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd {
		return
	}
	__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd = __obf_2298ab2cdf3cb6bd
	__obf_54823070883b28a6(__obf_a7affdfda169fd03)
}

// Round a up to nd digits (or fewer).
func (__obf_a7affdfda169fd03 *__obf_210b94d2c4b23455) RoundUp(__obf_2298ab2cdf3cb6bd int) {
	if __obf_2298ab2cdf3cb6bd < 0 || __obf_2298ab2cdf3cb6bd >= __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd {
		return
	}

	// round up
	for __obf_dec02b860bfb4fdd := __obf_2298ab2cdf3cb6bd - 1; __obf_dec02b860bfb4fdd >= 0; __obf_dec02b860bfb4fdd-- {
		__obf_820da776a4f8225c := __obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_dec02b860bfb4fdd]
		if __obf_820da776a4f8225c < '9' { // can stop after this digit
			__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_dec02b860bfb4fdd]++
			__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd = __obf_dec02b860bfb4fdd + 1
			return
		}
	}

	// Number is all 9s.
	// Change to single 1 with adjusted decimal point.
	__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[0] = '1'
	__obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd = 1
	__obf_a7affdfda169fd03.__obf_44d25d407353f76c++
}

// Extract integer part, rounded appropriately.
// No guarantees about overflow.
func (__obf_a7affdfda169fd03 *__obf_210b94d2c4b23455) RoundedInteger() uint64 {
	if __obf_a7affdfda169fd03.__obf_44d25d407353f76c > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_dec02b860bfb4fdd int
	__obf_46e3648cfcdd2b28 := uint64(0)
	for __obf_dec02b860bfb4fdd = 0; __obf_dec02b860bfb4fdd < __obf_a7affdfda169fd03.__obf_44d25d407353f76c && __obf_dec02b860bfb4fdd < __obf_a7affdfda169fd03.__obf_2298ab2cdf3cb6bd; __obf_dec02b860bfb4fdd++ {
		__obf_46e3648cfcdd2b28 = __obf_46e3648cfcdd2b28*10 + uint64(__obf_a7affdfda169fd03.__obf_45d3205e59c45c8b[__obf_dec02b860bfb4fdd]-'0')
	}
	for ; __obf_dec02b860bfb4fdd < __obf_a7affdfda169fd03.__obf_44d25d407353f76c; __obf_dec02b860bfb4fdd++ {
		__obf_46e3648cfcdd2b28 *= 10
	}
	if __obf_e61bee8b27eb8a58(__obf_a7affdfda169fd03, __obf_a7affdfda169fd03.__obf_44d25d407353f76c) {
		__obf_46e3648cfcdd2b28++
	}
	return __obf_46e3648cfcdd2b28
}
