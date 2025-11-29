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
package __obf_3e0a215d3ac5298d

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

var __obf_8e244cd834e11609 = big.NewInt(0)
var __obf_8ddf5479a62fb2a1 = big.NewInt(1)
var __obf_013f5033320a21e5 = big.NewInt(2)
var __obf_49fac42b5d78351d = big.NewInt(4)
var __obf_40db68f6d1a5418a = big.NewInt(5)
var __obf_8fd4b9ddee2a2f0a = big.NewInt(10)
var __obf_20c5d1868eb7c331 = big.NewInt(20)

var __obf_3184c27ef4f96eb3 = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_980b45647b76adba *big.Int
	__obf_4b33f2812f2bcc7e int32 // NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.

}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_980b45647b76adba int64, __obf_4b33f2812f2bcc7e int32) Decimal {
	return Decimal{__obf_980b45647b76adba: big.NewInt(__obf_980b45647b76adba), __obf_4b33f2812f2bcc7e: __obf_4b33f2812f2bcc7e}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_980b45647b76adba int64) Decimal {
	return Decimal{__obf_980b45647b76adba: big.NewInt(__obf_980b45647b76adba), __obf_4b33f2812f2bcc7e: 0}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_980b45647b76adba int32) Decimal {
	return Decimal{__obf_980b45647b76adba: big.NewInt(int64(__obf_980b45647b76adba)), __obf_4b33f2812f2bcc7e: 0}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_980b45647b76adba uint64) Decimal {
	return Decimal{__obf_980b45647b76adba: new(big.Int).SetUint64(__obf_980b45647b76adba), __obf_4b33f2812f2bcc7e: 0}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_980b45647b76adba *big.Int, __obf_4b33f2812f2bcc7e int32) Decimal {
	return Decimal{__obf_980b45647b76adba: new(big.Int).Set(__obf_980b45647b76adba), __obf_4b33f2812f2bcc7e: __obf_4b33f2812f2bcc7e}
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
func NewFromBigRat(__obf_980b45647b76adba *big.Rat, __obf_02e5a6c8816cdbf7 int32) Decimal {
	return Decimal{__obf_980b45647b76adba: new(big.Int).Set(__obf_980b45647b76adba.Num()), __obf_4b33f2812f2bcc7e: 0}.DivRound(Decimal{__obf_980b45647b76adba: new(big.Int).Set(__obf_980b45647b76adba.Denom()), __obf_4b33f2812f2bcc7e: 0}, __obf_02e5a6c8816cdbf7)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_980b45647b76adba string) (Decimal, error) {
	__obf_7f49d9bbb4694573 := __obf_980b45647b76adba
	var __obf_d3bd696a5bbb6bac string
	var __obf_4b33f2812f2bcc7e int64
	__obf_d6204b55dfb240a6 := // Check if number is using scientific notation
		strings.IndexAny(__obf_980b45647b76adba, "Ee")
	if __obf_d6204b55dfb240a6 != -1 {
		__obf_a46f4b89a28d08f8, __obf_3200dc10096fdd51 := strconv.ParseInt(__obf_980b45647b76adba[__obf_d6204b55dfb240a6+1:], 10, 32)
		if __obf_3200dc10096fdd51 != nil {
			if __obf_94cf3676209f21ab, __obf_38cbfc2225f19927 := __obf_3200dc10096fdd51.(*strconv.NumError); __obf_38cbfc2225f19927 && __obf_94cf3676209f21ab.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_980b45647b76adba)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_980b45647b76adba)
		}
		__obf_980b45647b76adba = __obf_980b45647b76adba[:__obf_d6204b55dfb240a6]
		__obf_4b33f2812f2bcc7e = __obf_a46f4b89a28d08f8
	}
	__obf_edc2e265513eef2a := -1
	__obf_14d8ac2b1e00f1c4 := len(__obf_980b45647b76adba)
	for __obf_a32504ffc5fafcfe := 0; __obf_a32504ffc5fafcfe < __obf_14d8ac2b1e00f1c4; __obf_a32504ffc5fafcfe++ {
		if __obf_980b45647b76adba[__obf_a32504ffc5fafcfe] == '.' {
			if __obf_edc2e265513eef2a > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_980b45647b76adba)
			}
			__obf_edc2e265513eef2a = __obf_a32504ffc5fafcfe
		}
	}

	if __obf_edc2e265513eef2a == -1 {
		__obf_d3bd696a5bbb6bac = // There is no decimal point, we can just parse the original string as
			// an int
			__obf_980b45647b76adba
	} else {
		if __obf_edc2e265513eef2a+1 < __obf_14d8ac2b1e00f1c4 {
			__obf_d3bd696a5bbb6bac = __obf_980b45647b76adba[:__obf_edc2e265513eef2a] + __obf_980b45647b76adba[__obf_edc2e265513eef2a+1:]
		} else {
			__obf_d3bd696a5bbb6bac = __obf_980b45647b76adba[:__obf_edc2e265513eef2a]
		}
		__obf_a46f4b89a28d08f8 := -len(__obf_980b45647b76adba[__obf_edc2e265513eef2a+1:])
		__obf_4b33f2812f2bcc7e += int64(__obf_a46f4b89a28d08f8)
	}

	var __obf_1dbd038f3e2365a5 *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_d3bd696a5bbb6bac) <= 18 {
		__obf_103f7fb6047e5f05, __obf_3200dc10096fdd51 := strconv.ParseInt(__obf_d3bd696a5bbb6bac, 10, 64)
		if __obf_3200dc10096fdd51 != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_980b45647b76adba)
		}
		__obf_1dbd038f3e2365a5 = big.NewInt(__obf_103f7fb6047e5f05)
	} else {
		__obf_1dbd038f3e2365a5 = new(big.Int)
		_, __obf_38cbfc2225f19927 := __obf_1dbd038f3e2365a5.SetString(__obf_d3bd696a5bbb6bac, 10)
		if !__obf_38cbfc2225f19927 {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_980b45647b76adba)
		}
	}

	if __obf_4b33f2812f2bcc7e < math.MinInt32 || __obf_4b33f2812f2bcc7e > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_7f49d9bbb4694573)
	}

	return Decimal{__obf_980b45647b76adba: __obf_1dbd038f3e2365a5, __obf_4b33f2812f2bcc7e: int32(__obf_4b33f2812f2bcc7e)}, nil
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
func NewFromFormattedString(__obf_980b45647b76adba string, __obf_3c8b31dd1f605eea *regexp.Regexp) (Decimal, error) {
	__obf_2d544ba25e4967b1 := __obf_3c8b31dd1f605eea.ReplaceAllString(__obf_980b45647b76adba, "")
	__obf_e498fd1d3feac2c4, __obf_3200dc10096fdd51 := NewFromString(__obf_2d544ba25e4967b1)
	if __obf_3200dc10096fdd51 != nil {
		return Decimal{}, __obf_3200dc10096fdd51
	}
	return __obf_e498fd1d3feac2c4, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_980b45647b76adba string) Decimal {
	__obf_2aecda69462f7dfa, __obf_3200dc10096fdd51 := NewFromString(__obf_980b45647b76adba)
	if __obf_3200dc10096fdd51 != nil {
		panic(__obf_3200dc10096fdd51)
	}
	return __obf_2aecda69462f7dfa
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
func NewFromFloat(__obf_980b45647b76adba float64) Decimal {
	if __obf_980b45647b76adba == 0 {
		return New(0, 0)
	}
	return __obf_2a35f9e5d0740ea8(__obf_980b45647b76adba, math.Float64bits(__obf_980b45647b76adba), &__obf_dc4c42ad9bdcae6e)
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
func NewFromFloat32(__obf_980b45647b76adba float32) Decimal {
	if __obf_980b45647b76adba == 0 {
		return New(0, 0)
	}
	__obf_55d13866fd49a7b4 := // XOR is workaround for https://github.com/golang/go/issues/26285
		math.Float32bits(__obf_980b45647b76adba) ^ 0x80808080
	return __obf_2a35f9e5d0740ea8(float64(__obf_980b45647b76adba), uint64(__obf_55d13866fd49a7b4)^0x80808080, &__obf_9381b591902a5ca1)
}

func __obf_2a35f9e5d0740ea8(__obf_6bd018db5b97cc18 float64, __obf_caac7889db6548df uint64, __obf_414bae728c6159e7 *__obf_b9da0d5925fb006e) Decimal {
	if math.IsNaN(__obf_6bd018db5b97cc18) || math.IsInf(__obf_6bd018db5b97cc18, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_6bd018db5b97cc18))
	}
	__obf_4b33f2812f2bcc7e := int(__obf_caac7889db6548df>>__obf_414bae728c6159e7.__obf_179bd2684cdde9cd) & (1<<__obf_414bae728c6159e7.__obf_1eecfa2ca0d4455f - 1)
	__obf_a8350a62e8fe3cad := __obf_caac7889db6548df & (uint64(1)<<__obf_414bae728c6159e7.__obf_179bd2684cdde9cd - 1)

	switch __obf_4b33f2812f2bcc7e {
	case 0:
		__obf_4b33f2812f2bcc7e++ // denormalized

	default:
		__obf_a8350a62e8fe3cad |= // add implicit top bit
			uint64(1) << __obf_414bae728c6159e7.__obf_179bd2684cdde9cd
	}
	__obf_4b33f2812f2bcc7e += __obf_414bae728c6159e7.__obf_b1391270660d97a8

	var __obf_e498fd1d3feac2c4 __obf_3e0a215d3ac5298d
	__obf_e498fd1d3feac2c4.
		Assign(__obf_a8350a62e8fe3cad)
	__obf_e498fd1d3feac2c4.
		Shift(__obf_4b33f2812f2bcc7e - int(__obf_414bae728c6159e7.__obf_179bd2684cdde9cd))
	__obf_e498fd1d3feac2c4.__obf_20103b61d9019c90 = __obf_caac7889db6548df>>(__obf_414bae728c6159e7.__obf_1eecfa2ca0d4455f+__obf_414bae728c6159e7.__obf_179bd2684cdde9cd) != 0
	__obf_4e452f472f313f6b(&__obf_e498fd1d3feac2c4,
		// If less than 19 digits, we can do calculation in an int64.
		__obf_a8350a62e8fe3cad, __obf_4b33f2812f2bcc7e, __obf_414bae728c6159e7)

	if __obf_e498fd1d3feac2c4.__obf_7a088d2e01e4cd05 < 19 {
		__obf_d7f93af02d364e15 := int64(0)
		__obf_e524d5ae197a0bbb := int64(1)
		for __obf_a32504ffc5fafcfe := __obf_e498fd1d3feac2c4.__obf_7a088d2e01e4cd05 - 1; __obf_a32504ffc5fafcfe >= 0; __obf_a32504ffc5fafcfe-- {
			__obf_d7f93af02d364e15 += __obf_e524d5ae197a0bbb * int64(__obf_e498fd1d3feac2c4.__obf_e498fd1d3feac2c4[__obf_a32504ffc5fafcfe]-'0')
			__obf_e524d5ae197a0bbb *= 10
		}
		if __obf_e498fd1d3feac2c4.__obf_20103b61d9019c90 {
			__obf_d7f93af02d364e15 *= -1
		}
		return Decimal{__obf_980b45647b76adba: big.NewInt(__obf_d7f93af02d364e15), __obf_4b33f2812f2bcc7e: int32(__obf_e498fd1d3feac2c4.__obf_164d07da7c832f2c) - int32(__obf_e498fd1d3feac2c4.__obf_7a088d2e01e4cd05)}
	}
	__obf_1dbd038f3e2365a5 := new(big.Int)
	__obf_1dbd038f3e2365a5, __obf_38cbfc2225f19927 := __obf_1dbd038f3e2365a5.SetString(string(__obf_e498fd1d3feac2c4.__obf_e498fd1d3feac2c4[:__obf_e498fd1d3feac2c4.__obf_7a088d2e01e4cd05]), 10)
	if __obf_38cbfc2225f19927 {
		return Decimal{__obf_980b45647b76adba: __obf_1dbd038f3e2365a5, __obf_4b33f2812f2bcc7e: int32(__obf_e498fd1d3feac2c4.__obf_164d07da7c832f2c) - int32(__obf_e498fd1d3feac2c4.__obf_7a088d2e01e4cd05)}
	}

	return NewFromFloatWithExponent(__obf_6bd018db5b97cc18, int32(__obf_e498fd1d3feac2c4.

		// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
		// number of fractional digits.
		//
		// Example:
		//
		//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
		__obf_164d07da7c832f2c)-int32(__obf_e498fd1d3feac2c4.__obf_7a088d2e01e4cd05))
}

func NewFromFloatWithExponent(__obf_980b45647b76adba float64, __obf_4b33f2812f2bcc7e int32) Decimal {
	if math.IsNaN(__obf_980b45647b76adba) || math.IsInf(__obf_980b45647b76adba, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_980b45647b76adba))
	}
	__obf_caac7889db6548df := math.Float64bits(__obf_980b45647b76adba)
	__obf_a8350a62e8fe3cad := __obf_caac7889db6548df & (1<<52 - 1)
	__obf_ea52fee6c6b6b144 := int32((__obf_caac7889db6548df >> 52) & (1<<11 - 1))
	__obf_ae515fc8976e420b := __obf_caac7889db6548df >> 63

	if __obf_ea52fee6c6b6b144 == 0 {
		// specials
		if __obf_a8350a62e8fe3cad == 0 {
			return Decimal{}
		}
		__obf_ea52fee6c6b6b144++ // subnormal

	} else {
		__obf_a8350a62e8fe3cad |= // normal
			1 << 52
	}
	__obf_ea52fee6c6b6b144 -= 1023 + 52

	// normalizing base-2 values
	for __obf_a8350a62e8fe3cad&1 == 0 {
		__obf_a8350a62e8fe3cad = __obf_a8350a62e8fe3cad >> 1
		__obf_ea52fee6c6b6b144++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_4b33f2812f2bcc7e < 0 && __obf_4b33f2812f2bcc7e < __obf_ea52fee6c6b6b144 {
		if __obf_ea52fee6c6b6b144 < 0 {
			__obf_4b33f2812f2bcc7e = __obf_ea52fee6c6b6b144
		} else {
			__obf_4b33f2812f2bcc7e = 0
		}
	}
	__obf_ea52fee6c6b6b144 -= // representing 10^M * 2^N as 5^M * 2^(M+N)
		__obf_4b33f2812f2bcc7e
	__obf_ec0ed206003885e9 := big.NewInt(1)
	__obf_3afe2fe6e8605cdf := big.NewInt(int64(__obf_a8350a62e8fe3cad))

	// applying 5^M
	if __obf_4b33f2812f2bcc7e > 0 {
		__obf_ec0ed206003885e9 = __obf_ec0ed206003885e9.SetInt64(int64(__obf_4b33f2812f2bcc7e))
		__obf_ec0ed206003885e9 = __obf_ec0ed206003885e9.Exp(__obf_40db68f6d1a5418a, __obf_ec0ed206003885e9, nil)
	} else if __obf_4b33f2812f2bcc7e < 0 {
		__obf_ec0ed206003885e9 = __obf_ec0ed206003885e9.SetInt64(-int64(__obf_4b33f2812f2bcc7e))
		__obf_ec0ed206003885e9 = __obf_ec0ed206003885e9.Exp(__obf_40db68f6d1a5418a, __obf_ec0ed206003885e9, nil)
		__obf_3afe2fe6e8605cdf = __obf_3afe2fe6e8605cdf.Mul(__obf_3afe2fe6e8605cdf, __obf_ec0ed206003885e9)
		__obf_ec0ed206003885e9 = __obf_ec0ed206003885e9.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_ea52fee6c6b6b144 > 0 {
		__obf_3afe2fe6e8605cdf = __obf_3afe2fe6e8605cdf.Lsh(__obf_3afe2fe6e8605cdf, uint(__obf_ea52fee6c6b6b144))
	} else if __obf_ea52fee6c6b6b144 < 0 {
		__obf_ec0ed206003885e9 = __obf_ec0ed206003885e9.Lsh(__obf_ec0ed206003885e9, uint(-__obf_ea52fee6c6b6b144))
	}

	// rounding and downscaling
	if __obf_4b33f2812f2bcc7e > 0 || __obf_ea52fee6c6b6b144 < 0 {
		__obf_476af5c381c1b786 := new(big.Int).Rsh(__obf_ec0ed206003885e9, 1)
		__obf_3afe2fe6e8605cdf = __obf_3afe2fe6e8605cdf.Add(__obf_3afe2fe6e8605cdf, __obf_476af5c381c1b786)
		__obf_3afe2fe6e8605cdf = __obf_3afe2fe6e8605cdf.Quo(__obf_3afe2fe6e8605cdf, __obf_ec0ed206003885e9)
	}

	if __obf_ae515fc8976e420b == 1 {
		__obf_3afe2fe6e8605cdf = __obf_3afe2fe6e8605cdf.Neg(__obf_3afe2fe6e8605cdf)
	}

	return Decimal{__obf_980b45647b76adba: __obf_3afe2fe6e8605cdf, __obf_4b33f2812f2bcc7e: __obf_4b33f2812f2bcc7e}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_e498fd1d3feac2c4 Decimal) Copy() Decimal {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()
	return Decimal{__obf_980b45647b76adba: new(big.Int).Set(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba), __obf_4b33f2812f2bcc7e: __obf_e498fd1d3feac2c4.

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
		__obf_4b33f2812f2bcc7e,
	}
}

func (__obf_e498fd1d3feac2c4 Decimal) __obf_b297b59a32692f76(__obf_4b33f2812f2bcc7e int32) Decimal {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()

	if __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e == __obf_4b33f2812f2bcc7e {
		return Decimal{
			new(big.Int).Set(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba), __obf_e498fd1d3feac2c4.

				// NOTE(vadim): must convert exps to float64 before - to prevent overflow
				__obf_4b33f2812f2bcc7e,
		}
	}
	__obf_d2e82c3abf930c78 := math.Abs(float64(__obf_4b33f2812f2bcc7e) - float64(__obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e))
	__obf_980b45647b76adba := new(big.Int).Set(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba)
	__obf_26416f22601b0983 := new(big.Int).Exp(__obf_8fd4b9ddee2a2f0a, big.NewInt(int64(__obf_d2e82c3abf930c78)), nil)
	if __obf_4b33f2812f2bcc7e > __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e {
		__obf_980b45647b76adba = __obf_980b45647b76adba.Quo(__obf_980b45647b76adba, __obf_26416f22601b0983)
	} else if __obf_4b33f2812f2bcc7e < __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e {
		__obf_980b45647b76adba = __obf_980b45647b76adba.Mul(__obf_980b45647b76adba, __obf_26416f22601b0983)
	}

	return Decimal{__obf_980b45647b76adba: __obf_980b45647b76adba, __obf_4b33f2812f2bcc7e: __obf_4b33f2812f2bcc7e}
}

// Abs returns the absolute value of the decimal.
func (__obf_e498fd1d3feac2c4 Decimal) Abs() Decimal {
	if !__obf_e498fd1d3feac2c4.IsNegative() {
		return __obf_e498fd1d3feac2c4
	}
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()
	__obf_379aae7d7fb7305e := new(big.Int).Abs(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba)
	return Decimal{__obf_980b45647b76adba: __obf_379aae7d7fb7305e, __obf_4b33f2812f2bcc7e: __obf_e498fd1d3feac2c4.

		// Add returns d + d2.
		__obf_4b33f2812f2bcc7e,
	}
}

func (__obf_e498fd1d3feac2c4 Decimal) Add(__obf_e48c8e69bc461c40 Decimal) Decimal {
	__obf_4e9cf9f19030e561, __obf_31e081dce7d8bd2b := RescalePair(__obf_e498fd1d3feac2c4, __obf_e48c8e69bc461c40)
	__obf_858eabb9c407b31f := new(big.Int).Add(__obf_4e9cf9f19030e561.__obf_980b45647b76adba, __obf_31e081dce7d8bd2b.__obf_980b45647b76adba)
	return Decimal{__obf_980b45647b76adba: __obf_858eabb9c407b31f, __obf_4b33f2812f2bcc7e: __obf_4e9cf9f19030e561.

		// Sub returns d - d2.
		__obf_4b33f2812f2bcc7e,
	}
}

func (__obf_e498fd1d3feac2c4 Decimal) Sub(__obf_e48c8e69bc461c40 Decimal) Decimal {
	__obf_4e9cf9f19030e561, __obf_31e081dce7d8bd2b := RescalePair(__obf_e498fd1d3feac2c4, __obf_e48c8e69bc461c40)
	__obf_858eabb9c407b31f := new(big.Int).Sub(__obf_4e9cf9f19030e561.__obf_980b45647b76adba, __obf_31e081dce7d8bd2b.__obf_980b45647b76adba)
	return Decimal{__obf_980b45647b76adba: __obf_858eabb9c407b31f, __obf_4b33f2812f2bcc7e: __obf_4e9cf9f19030e561.

		// Neg returns -d.
		__obf_4b33f2812f2bcc7e,
	}
}

func (__obf_e498fd1d3feac2c4 Decimal) Neg() Decimal {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()
	__obf_6bd018db5b97cc18 := new(big.Int).Neg(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba)
	return Decimal{__obf_980b45647b76adba: __obf_6bd018db5b97cc18, __obf_4b33f2812f2bcc7e: __obf_e498fd1d3feac2c4.

		// Mul returns d * d2.
		__obf_4b33f2812f2bcc7e,
	}
}

func (__obf_e498fd1d3feac2c4 Decimal) Mul(__obf_e48c8e69bc461c40 Decimal) Decimal {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()
	__obf_e48c8e69bc461c40.__obf_a167cff196883330()
	__obf_d04c4703bdfcce99 := int64(__obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e) + int64(__obf_e48c8e69bc461c40.__obf_4b33f2812f2bcc7e)
	if __obf_d04c4703bdfcce99 > math.MaxInt32 || __obf_d04c4703bdfcce99 < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_d04c4703bdfcce99))
	}
	__obf_858eabb9c407b31f := new(big.Int).Mul(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba, __obf_e48c8e69bc461c40.__obf_980b45647b76adba)
	return Decimal{__obf_980b45647b76adba: __obf_858eabb9c407b31f, __obf_4b33f2812f2bcc7e: int32(__obf_d04c4703bdfcce99)}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_e498fd1d3feac2c4 Decimal) Shift(__obf_233037f91159f9d7 int32) Decimal {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()
	return Decimal{__obf_980b45647b76adba: new(big.Int).Set(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba), __obf_4b33f2812f2bcc7e: __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e + __obf_233037f91159f9d7}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_e498fd1d3feac2c4 Decimal) Div(__obf_e48c8e69bc461c40 Decimal) Decimal {
	return __obf_e498fd1d3feac2c4.DivRound(__obf_e48c8e69bc461c40, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_e498fd1d3feac2c4 Decimal) QuoRem(__obf_e48c8e69bc461c40 Decimal, __obf_02e5a6c8816cdbf7 int32) (Decimal, Decimal) {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()
	__obf_e48c8e69bc461c40.__obf_a167cff196883330()
	if __obf_e48c8e69bc461c40.__obf_980b45647b76adba.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_3505d65c8939662e := -__obf_02e5a6c8816cdbf7
	__obf_94cf3676209f21ab := int64(__obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e) - int64(__obf_e48c8e69bc461c40.__obf_4b33f2812f2bcc7e) - int64(__obf_3505d65c8939662e)
	if __obf_94cf3676209f21ab > math.MaxInt32 || __obf_94cf3676209f21ab < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_7f9ecbec2cdb72a5, __obf_f7709720f8581805,

		// d = a 10^ea
		// d2 = b 10^eb
		__obf_f96ce57b80f82ef5 big.Int
	var __obf_dbcede39fc475511 int32

	if __obf_94cf3676209f21ab < 0 {
		__obf_7f9ecbec2cdb72a5 = *__obf_e498fd1d3feac2c4.__obf_980b45647b76adba
		__obf_f96ce57b80f82ef5.
			SetInt64(-__obf_94cf3676209f21ab)
		__obf_f7709720f8581805.
			Exp(__obf_8fd4b9ddee2a2f0a, &__obf_f96ce57b80f82ef5, nil)
		__obf_f7709720f8581805.
			Mul(__obf_e48c8e69bc461c40.__obf_980b45647b76adba, &__obf_f7709720f8581805)
		__obf_dbcede39fc475511 = __obf_e498fd1d3feac2c4.
			// now aa = a
			//     bb = b 10^(scale + eb - ea)
			__obf_4b33f2812f2bcc7e

	} else {
		__obf_f96ce57b80f82ef5.
			SetInt64(__obf_94cf3676209f21ab)
		__obf_7f9ecbec2cdb72a5.
			Exp(__obf_8fd4b9ddee2a2f0a, &__obf_f96ce57b80f82ef5, nil)
		__obf_7f9ecbec2cdb72a5.
			Mul(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba, &__obf_7f9ecbec2cdb72a5)
		__obf_f7709720f8581805 = *__obf_e48c8e69bc461c40.__obf_980b45647b76adba

		// now aa = a ^ (ea - eb - scale)
		//     bb = b
		__obf_dbcede39fc475511 = __obf_3505d65c8939662e + __obf_e48c8e69bc461c40.__obf_4b33f2812f2bcc7e

	}
	var __obf_209c7ec826089def, __obf_b827ad869fe015ad big.Int
	__obf_209c7ec826089def.
		QuoRem(&__obf_7f9ecbec2cdb72a5, &__obf_f7709720f8581805, &__obf_b827ad869fe015ad)
	__obf_7aa08c29e83d7810 := Decimal{__obf_980b45647b76adba: &__obf_209c7ec826089def, __obf_4b33f2812f2bcc7e: __obf_3505d65c8939662e}
	__obf_05aa2c418f53b886 := Decimal{__obf_980b45647b76adba: &__obf_b827ad869fe015ad, __obf_4b33f2812f2bcc7e: __obf_dbcede39fc475511}
	return __obf_7aa08c29e83d7810,

		// DivRound divides and rounds to a given precision
		// i.e. to an integer multiple of 10^(-precision)
		//
		//	for a positive quotient digit 5 is rounded up, away from 0
		//	if the quotient is negative then digit 5 is rounded down, away from 0
		//
		// Note that precision<0 is allowed as input.
		__obf_05aa2c418f53b886
}

func (__obf_e498fd1d3feac2c4 Decimal) DivRound(__obf_e48c8e69bc461c40 Decimal, __obf_02e5a6c8816cdbf7 int32) Decimal {
	__obf_209c7ec826089def,
		// QuoRem already checks initialization
		__obf_b827ad869fe015ad := __obf_e498fd1d3feac2c4.QuoRem(__obf_e48c8e69bc461c40,
		// the actual rounding decision is based on comparing r*10^precision and d2/2
		// instead compare 2 r 10 ^precision and d2
		__obf_02e5a6c8816cdbf7)

	var __obf_630ec67007f687c6 big.Int
	__obf_630ec67007f687c6.
		Abs(__obf_b827ad869fe015ad.__obf_980b45647b76adba)
	__obf_630ec67007f687c6.
		Lsh(&__obf_630ec67007f687c6, 1)
	__obf_dfdd582c2ad36291 := // now rv2 = abs(r.value) * 2
		Decimal{__obf_980b45647b76adba: &__obf_630ec67007f687c6, __obf_4b33f2812f2bcc7e: __obf_b827ad869fe015ad.
			// r2 is now 2 * r * 10 ^ precision
			__obf_4b33f2812f2bcc7e + __obf_02e5a6c8816cdbf7}

	var __obf_d0f3a30e016877c1 = __obf_dfdd582c2ad36291.Cmp(__obf_e48c8e69bc461c40.Abs())

	if __obf_d0f3a30e016877c1 < 0 {
		return __obf_209c7ec826089def
	}

	if __obf_e498fd1d3feac2c4.__obf_980b45647b76adba.Sign()*__obf_e48c8e69bc461c40.__obf_980b45647b76adba.Sign() < 0 {
		return __obf_209c7ec826089def.Sub(New(1, -__obf_02e5a6c8816cdbf7))
	}

	return __obf_209c7ec826089def.Add(New(1, -__obf_02e5a6c8816cdbf7))
}

// Mod returns d % d2.
func (__obf_e498fd1d3feac2c4 Decimal) Mod(__obf_e48c8e69bc461c40 Decimal) Decimal {
	_, __obf_b827ad869fe015ad := __obf_e498fd1d3feac2c4.QuoRem(__obf_e48c8e69bc461c40, 0)
	return __obf_b827ad869fe015ad
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
func (__obf_e498fd1d3feac2c4 Decimal) Pow(__obf_e48c8e69bc461c40 Decimal) Decimal {
	__obf_31c4d1451161da09 := __obf_e498fd1d3feac2c4.Sign()
	__obf_ba1d61b8c07d6eb7 := __obf_e48c8e69bc461c40.Sign()

	if __obf_31c4d1451161da09 == 0 {
		if __obf_ba1d61b8c07d6eb7 == 0 {
			return Decimal{}
		}
		if __obf_ba1d61b8c07d6eb7 == 1 {
			return Decimal{__obf_8e244cd834e11609, 0}
		}
		if __obf_ba1d61b8c07d6eb7 == -1 {
			return Decimal{}
		}
	}

	if __obf_ba1d61b8c07d6eb7 == 0 {
		return Decimal{__obf_8ddf5479a62fb2a1, 0}
	}
	__obf_1ff12bb01961e39d := // TODO: optimize extraction of fractional part
		Decimal{__obf_8ddf5479a62fb2a1, 0}
	__obf_6f5cebac3a6f59be, __obf_2213a8585c2c9fbc := __obf_e48c8e69bc461c40.QuoRem(__obf_1ff12bb01961e39d, 0)

	if __obf_31c4d1451161da09 == -1 && !__obf_2213a8585c2c9fbc.IsZero() {
		return Decimal{}
	}
	__obf_6f29ba54bff4302f, _ := __obf_e498fd1d3feac2c4.PowBigInt(__obf_6f5cebac3a6f59be.

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_980b45647b76adba)

	if __obf_2213a8585c2c9fbc.__obf_980b45647b76adba.Sign() == 0 {
		return __obf_6f29ba54bff4302f
	}
	__obf_2e7f5f35b33a3ac1 := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_e498fd1d3feac2c4.NumDigits()
	__obf_569c92be1f9221eb := __obf_e48c8e69bc461c40.NumDigits()
	__obf_02e5a6c8816cdbf7 := __obf_2e7f5f35b33a3ac1

	if __obf_569c92be1f9221eb > __obf_02e5a6c8816cdbf7 {
		__obf_02e5a6c8816cdbf7 += __obf_569c92be1f9221eb
	}
	__obf_02e5a6c8816cdbf7 += 6
	__obf_874ae39867794a19,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_3200dc10096fdd51 := __obf_e498fd1d3feac2c4.Abs().Ln(-__obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e + int32(__obf_02e5a6c8816cdbf7))
	if __obf_3200dc10096fdd51 != nil {
		return Decimal{}
	}
	__obf_874ae39867794a19 = __obf_874ae39867794a19.Mul(__obf_2213a8585c2c9fbc)
	__obf_874ae39867794a19, __obf_3200dc10096fdd51 = __obf_874ae39867794a19.ExpTaylor(-__obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e + int32(__obf_02e5a6c8816cdbf7))
	if __obf_3200dc10096fdd51 != nil {
		return Decimal{}
	}
	__obf_b11445a694bc2930 := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_6f29ba54bff4302f.Mul(__obf_874ae39867794a19)

	return __obf_b11445a694bc2930
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
func (__obf_e498fd1d3feac2c4 Decimal) PowWithPrecision(__obf_e48c8e69bc461c40 Decimal, __obf_02e5a6c8816cdbf7 int32) (Decimal, error) {
	__obf_31c4d1451161da09 := __obf_e498fd1d3feac2c4.Sign()
	__obf_ba1d61b8c07d6eb7 := __obf_e48c8e69bc461c40.Sign()

	if __obf_31c4d1451161da09 == 0 {
		if __obf_ba1d61b8c07d6eb7 == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_ba1d61b8c07d6eb7 == 1 {
			return Decimal{__obf_8e244cd834e11609, 0}, nil
		}
		if __obf_ba1d61b8c07d6eb7 == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_ba1d61b8c07d6eb7 == 0 {
		return Decimal{__obf_8ddf5479a62fb2a1, 0}, nil
	}
	__obf_1ff12bb01961e39d := // TODO: optimize extraction of fractional part
		Decimal{__obf_8ddf5479a62fb2a1, 0}
	__obf_6f5cebac3a6f59be, __obf_2213a8585c2c9fbc := __obf_e48c8e69bc461c40.QuoRem(__obf_1ff12bb01961e39d, 0)

	if __obf_31c4d1451161da09 == -1 && !__obf_2213a8585c2c9fbc.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}
	__obf_6f29ba54bff4302f, _ := __obf_e498fd1d3feac2c4.__obf_b9d4029375d69979(__obf_6f5cebac3a6f59be.__obf_980b45647b76adba,

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_02e5a6c8816cdbf7)

	if __obf_2213a8585c2c9fbc.__obf_980b45647b76adba.Sign() == 0 {
		return __obf_6f29ba54bff4302f, nil
	}
	__obf_2e7f5f35b33a3ac1 := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_e498fd1d3feac2c4.NumDigits()
	__obf_569c92be1f9221eb := __obf_e48c8e69bc461c40.NumDigits()

	if int32(__obf_2e7f5f35b33a3ac1) > __obf_02e5a6c8816cdbf7 {
		__obf_02e5a6c8816cdbf7 = int32(__obf_2e7f5f35b33a3ac1)
	}
	if int32(__obf_569c92be1f9221eb) > __obf_02e5a6c8816cdbf7 {
		__obf_02e5a6c8816cdbf7 += int32(__obf_569c92be1f9221eb)
	}
	__obf_02e5a6c8816cdbf7 += // increase precision by 10 to compensate for errors in further calculations
		10
	__obf_874ae39867794a19,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_3200dc10096fdd51 := __obf_e498fd1d3feac2c4.Abs().Ln(__obf_02e5a6c8816cdbf7)
	if __obf_3200dc10096fdd51 != nil {
		return Decimal{}, __obf_3200dc10096fdd51
	}
	__obf_874ae39867794a19 = __obf_874ae39867794a19.Mul(__obf_2213a8585c2c9fbc)
	__obf_874ae39867794a19, __obf_3200dc10096fdd51 = __obf_874ae39867794a19.ExpTaylor(__obf_02e5a6c8816cdbf7)
	if __obf_3200dc10096fdd51 != nil {
		return Decimal{}, __obf_3200dc10096fdd51
	}
	__obf_b11445a694bc2930 := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_6f29ba54bff4302f.Mul(__obf_874ae39867794a19)

	return __obf_b11445a694bc2930, nil
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
func (__obf_e498fd1d3feac2c4 Decimal) PowInt32(__obf_4b33f2812f2bcc7e int32) (Decimal, error) {
	if __obf_e498fd1d3feac2c4.IsZero() && __obf_4b33f2812f2bcc7e == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_a2813f27f5b054c4 := __obf_4b33f2812f2bcc7e < 0
	__obf_4b33f2812f2bcc7e = __obf_e85d6a20e6805469(__obf_4b33f2812f2bcc7e)
	__obf_70b1c35e7203372c, __obf_7c54c7a651b9edcb := __obf_e498fd1d3feac2c4, New(1, 0)

	for __obf_4b33f2812f2bcc7e > 0 {
		if __obf_4b33f2812f2bcc7e%2 == 1 {
			__obf_7c54c7a651b9edcb = __obf_7c54c7a651b9edcb.Mul(__obf_70b1c35e7203372c)
		}
		__obf_4b33f2812f2bcc7e /= 2

		if __obf_4b33f2812f2bcc7e > 0 {
			__obf_70b1c35e7203372c = __obf_70b1c35e7203372c.Mul(__obf_70b1c35e7203372c)
		}
	}

	if __obf_a2813f27f5b054c4 {
		return New(1, 0).DivRound(__obf_7c54c7a651b9edcb, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_7c54c7a651b9edcb, nil
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
func (__obf_e498fd1d3feac2c4 Decimal) PowBigInt(__obf_4b33f2812f2bcc7e *big.Int) (Decimal, error) {
	return __obf_e498fd1d3feac2c4.__obf_b9d4029375d69979(__obf_4b33f2812f2bcc7e, int32(PowPrecisionNegativeExponent))
}

func (__obf_e498fd1d3feac2c4 Decimal) __obf_b9d4029375d69979(__obf_4b33f2812f2bcc7e *big.Int, __obf_02e5a6c8816cdbf7 int32) (Decimal, error) {
	if __obf_e498fd1d3feac2c4.IsZero() && __obf_4b33f2812f2bcc7e.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_5b771213aa9dd567 := new(big.Int).Set(__obf_4b33f2812f2bcc7e)
	__obf_a2813f27f5b054c4 := __obf_4b33f2812f2bcc7e.Sign() < 0

	if __obf_a2813f27f5b054c4 {
		__obf_5b771213aa9dd567.
			Abs(__obf_5b771213aa9dd567)
	}
	__obf_70b1c35e7203372c, __obf_7c54c7a651b9edcb := __obf_e498fd1d3feac2c4, New(1, 0)

	for __obf_5b771213aa9dd567.Sign() > 0 {
		if __obf_5b771213aa9dd567.Bit(0) == 1 {
			__obf_7c54c7a651b9edcb = __obf_7c54c7a651b9edcb.Mul(__obf_70b1c35e7203372c)
		}
		__obf_5b771213aa9dd567.
			Rsh(__obf_5b771213aa9dd567, 1)

		if __obf_5b771213aa9dd567.Sign() > 0 {
			__obf_70b1c35e7203372c = __obf_70b1c35e7203372c.Mul(__obf_70b1c35e7203372c)
		}
	}

	if __obf_a2813f27f5b054c4 {
		return New(1, 0).DivRound(__obf_7c54c7a651b9edcb, __obf_02e5a6c8816cdbf7), nil
	}

	return __obf_7c54c7a651b9edcb, nil
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
func (__obf_e498fd1d3feac2c4 Decimal) ExpHullAbrham(__obf_b38a8677ffd3b9f9 uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_e498fd1d3feac2c4.IsZero() {
		return Decimal{__obf_8ddf5479a62fb2a1, 0}, nil
	}
	__obf_125d88bd6f747292 := __obf_b38a8677ffd3b9f9

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_1cee75b7ac144682 := __obf_e498fd1d3feac2c4.Abs().InexactFloat64()
	if __obf_ff2eb166fe15e576 := __obf_1cee75b7ac144682 / 23; __obf_ff2eb166fe15e576 > float64(__obf_125d88bd6f747292) && __obf_ff2eb166fe15e576 < float64(ExpMaxIterations) {
		__obf_125d88bd6f747292 = uint32(math.Ceil(__obf_ff2eb166fe15e576))
	}
	__obf_6ecf80100457210d := // fail if abs(d) beyond an over/underflow threshold
		New(23*int64(__obf_125d88bd6f747292), 0)
	if __obf_e498fd1d3feac2c4.Abs().Cmp(__obf_6ecf80100457210d) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}
	__obf_6e3d56504ef90ce1 := // Return 1 if abs(d) small enough; this also avoids later over/underflow
		New(9, -int32(__obf_125d88bd6f747292)-1)
	if __obf_e498fd1d3feac2c4.Abs().Cmp(__obf_6e3d56504ef90ce1) <= 0 {
		return Decimal{__obf_8ddf5479a62fb2a1, __obf_e498fd1d3feac2c4.

			// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
			__obf_4b33f2812f2bcc7e}, nil
	}
	__obf_8adb0e032a322da8 := __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e + int32(__obf_e498fd1d3feac2c4.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_8adb0e032a322da8 < 0 {
		__obf_8adb0e032a322da8 = 0
	}
	__obf_0fefbf9a88cfc16b := New(1, __obf_8adb0e032a322da8)
	__obf_b827ad869fe015ad := // reduction factor
		Decimal{new(big.Int).Set(__obf_e498fd1d3feac2c4. // reduced argument
									__obf_980b45647b76adba), __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e - __obf_8adb0e032a322da8}
	__obf_594b8bb0de96dda4 := int32(__obf_125d88bd6f747292) + __obf_8adb0e032a322da8 + 2
	__obf_bd04a971f774331d := // precision for calculating the sum

		// Determine n, the number of therms for calculating sum
		// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
		// for solving appropriate equation, along with directed
		// roundings and simple rational bound for log10(p/abs(r))
		__obf_b827ad869fe015ad.Abs().InexactFloat64()
	__obf_0de5fe5bde1c42ce := float64(__obf_594b8bb0de96dda4)
	__obf_d3323498f6015b51 := math.Ceil((1.453*__obf_0de5fe5bde1c42ce - 1.182) / math.Log10(__obf_0de5fe5bde1c42ce/__obf_bd04a971f774331d))
	if __obf_d3323498f6015b51 > float64(ExpMaxIterations) || math.IsNaN(__obf_d3323498f6015b51) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_70b1c35e7203372c := int64(__obf_d3323498f6015b51)
	__obf_d7f93af02d364e15 := New(0, 0)
	__obf_422f54d84cbf9f8b := New(1, 0)
	__obf_1ff12bb01961e39d := New(1, 0)
	for __obf_a32504ffc5fafcfe := __obf_70b1c35e7203372c - 1; __obf_a32504ffc5fafcfe > 0; __obf_a32504ffc5fafcfe-- {
		__obf_d7f93af02d364e15.__obf_980b45647b76adba.
			SetInt64(__obf_a32504ffc5fafcfe)
		__obf_422f54d84cbf9f8b = __obf_422f54d84cbf9f8b.Mul(__obf_b827ad869fe015ad.DivRound(__obf_d7f93af02d364e15, __obf_594b8bb0de96dda4))
		__obf_422f54d84cbf9f8b = __obf_422f54d84cbf9f8b.Add(__obf_1ff12bb01961e39d)
	}
	__obf_38a43d81bbc3d298 := __obf_0fefbf9a88cfc16b.IntPart()
	__obf_b11445a694bc2930 := New(1, 0)
	for __obf_a32504ffc5fafcfe := __obf_38a43d81bbc3d298; __obf_a32504ffc5fafcfe > 0; __obf_a32504ffc5fafcfe-- {
		__obf_b11445a694bc2930 = __obf_b11445a694bc2930.Mul(__obf_422f54d84cbf9f8b)
	}
	__obf_0d178cd344c73373 := int32(__obf_b11445a694bc2930.NumDigits())

	var __obf_94c05a27a8a9887c int32
	if __obf_0d178cd344c73373 > __obf_e85d6a20e6805469(__obf_b11445a694bc2930.__obf_4b33f2812f2bcc7e) {
		__obf_94c05a27a8a9887c = int32(__obf_125d88bd6f747292) - __obf_0d178cd344c73373 - __obf_b11445a694bc2930.__obf_4b33f2812f2bcc7e
	} else {
		__obf_94c05a27a8a9887c = int32(__obf_125d88bd6f747292)
	}
	__obf_b11445a694bc2930 = __obf_b11445a694bc2930.Round(__obf_94c05a27a8a9887c)

	return __obf_b11445a694bc2930, nil
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
func (__obf_e498fd1d3feac2c4 Decimal) ExpTaylor(__obf_02e5a6c8816cdbf7 int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_e498fd1d3feac2c4.IsZero() {
		return Decimal{__obf_8ddf5479a62fb2a1, 0}.Round(__obf_02e5a6c8816cdbf7), nil
	}

	var __obf_17ac81c5d033c6c4 Decimal
	var __obf_e493dd2f213b65e3 int32
	if __obf_02e5a6c8816cdbf7 < 0 {
		__obf_17ac81c5d033c6c4 = New(1, -1)
		__obf_e493dd2f213b65e3 = 8
	} else {
		__obf_17ac81c5d033c6c4 = New(1, -__obf_02e5a6c8816cdbf7-1)
		__obf_e493dd2f213b65e3 = __obf_02e5a6c8816cdbf7 + 1
	}
	__obf_00c68320069dc073 := __obf_e498fd1d3feac2c4.Abs()
	__obf_2abdef815c2d4d03 := __obf_e498fd1d3feac2c4.Abs()
	__obf_b452315ca4e34aeb := New(1, 0)
	__obf_7c54c7a651b9edcb := New(1, 0)

	for __obf_a32504ffc5fafcfe := int64(1); ; {
		__obf_a89d23a3ccf0f69c := __obf_2abdef815c2d4d03.DivRound(__obf_b452315ca4e34aeb, __obf_e493dd2f213b65e3)
		__obf_7c54c7a651b9edcb = __obf_7c54c7a651b9edcb.Add(__obf_a89d23a3ccf0f69c)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_a89d23a3ccf0f69c.Cmp(__obf_17ac81c5d033c6c4) < 0 {
			break
		}
		__obf_2abdef815c2d4d03 = __obf_2abdef815c2d4d03.Mul(__obf_00c68320069dc073)
		__obf_a32504ffc5fafcfe++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_3184c27ef4f96eb3) >= int(__obf_a32504ffc5fafcfe) && !__obf_3184c27ef4f96eb3[__obf_a32504ffc5fafcfe-1].IsZero() {
			__obf_b452315ca4e34aeb = __obf_3184c27ef4f96eb3[__obf_a32504ffc5fafcfe-1]
		} else {
			__obf_b452315ca4e34aeb = // To avoid any race conditions, firstly the zero value is appended to a slice to create
				// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
				// factorial using the index notation.
				__obf_3184c27ef4f96eb3[__obf_a32504ffc5fafcfe-2].Mul(New(__obf_a32504ffc5fafcfe, 0))
			__obf_3184c27ef4f96eb3 = append(__obf_3184c27ef4f96eb3, Zero)
			__obf_3184c27ef4f96eb3[__obf_a32504ffc5fafcfe-1] = __obf_b452315ca4e34aeb
		}
	}

	if __obf_e498fd1d3feac2c4.Sign() < 0 {
		__obf_7c54c7a651b9edcb = New(1, 0).DivRound(__obf_7c54c7a651b9edcb, __obf_02e5a6c8816cdbf7+1)
	}
	__obf_7c54c7a651b9edcb = __obf_7c54c7a651b9edcb.Round(__obf_02e5a6c8816cdbf7)
	return __obf_7c54c7a651b9edcb, nil
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
func (__obf_e498fd1d3feac2c4 Decimal) Ln(__obf_02e5a6c8816cdbf7 int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_e498fd1d3feac2c4.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_e498fd1d3feac2c4.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}
	__obf_c56f06bc84be0ae8 := __obf_02e5a6c8816cdbf7 + 2
	__obf_3b3ef24c3805160d := __obf_e498fd1d3feac2c4.Copy()

	var __obf_3a0b292e10e4ffb7, __obf_ba85cb4f962d4b24, __obf_7f01c31aab0451c6, __obf_3e865660f88e266c, __obf_6b01e4b9bb112e56 Decimal
	__obf_3a0b292e10e4ffb7 = __obf_3b3ef24c3805160d.Sub(Decimal{__obf_8ddf5479a62fb2a1, 0})
	__obf_ba85cb4f962d4b24 = Decimal{__obf_8ddf5479a62fb2a1, -1}
	__obf_5c2046a7e051eacd := // for decimal in range [0.9, 1.1] where ln(d) is close to 0
		false

	if __obf_3a0b292e10e4ffb7.Abs().Cmp(__obf_ba85cb4f962d4b24) <= 0 {
		__obf_5c2046a7e051eacd = true
	} else {
		__obf_e3e8c330d5f3d713 := // reduce input decimal to range [0.1, 1)
			int32(__obf_3b3ef24c3805160d.NumDigits()) + __obf_3b3ef24c3805160d.__obf_4b33f2812f2bcc7e

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_3b3ef24c3805160d.__obf_4b33f2812f2bcc7e -= __obf_e3e8c330d5f3d713
		__obf_fb8ecd4a11c50c13 := __obf_fb8ecd4a11c50c13.__obf_54869c391870cab9(__obf_c56f06bc84be0ae8)
		__obf_6b01e4b9bb112e56 = NewFromInt32(__obf_e3e8c330d5f3d713)
		__obf_6b01e4b9bb112e56 = __obf_6b01e4b9bb112e56.Mul(__obf_fb8ecd4a11c50c13)
		__obf_3a0b292e10e4ffb7 = __obf_3b3ef24c3805160d.Sub(Decimal{__obf_8ddf5479a62fb2a1, 0})

		if __obf_3a0b292e10e4ffb7.Abs().Cmp(__obf_ba85cb4f962d4b24) <= 0 {
			__obf_5c2046a7e051eacd = true
		} else {
			__obf_c7fb9b7a5f5059de := // initial estimate using floats
				__obf_3b3ef24c3805160d.InexactFloat64()
			__obf_3a0b292e10e4ffb7 = NewFromFloat(math.Log(__obf_c7fb9b7a5f5059de))
		}
	}
	__obf_17ac81c5d033c6c4 := Decimal{__obf_8ddf5479a62fb2a1, -__obf_c56f06bc84be0ae8}

	if __obf_5c2046a7e051eacd {
		__obf_7f01c31aab0451c6 = // Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
			// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			// until the difference between current and next term is smaller than epsilon.
			// Coverage quite fast for decimals close to 1.0

			// z + 2
			__obf_3a0b292e10e4ffb7.Add(Decimal{__obf_013f5033320a21e5, 0})
		__obf_ba85cb4f962d4b24 = // z / (z + 2)
			__obf_3a0b292e10e4ffb7.DivRound(__obf_7f01c31aab0451c6, __obf_c56f06bc84be0ae8)
		__obf_3a0b292e10e4ffb7 = // 2 * (z / (z + 2))
			__obf_ba85cb4f962d4b24.Add(__obf_ba85cb4f962d4b24)
		__obf_7f01c31aab0451c6 = __obf_3a0b292e10e4ffb7.Copy()

		for __obf_70b1c35e7203372c := 1; ; __obf_70b1c35e7203372c++ {
			__obf_7f01c31aab0451c6 = // 2 * (z / (z+2))^(2n+1)
				__obf_7f01c31aab0451c6.Mul(__obf_ba85cb4f962d4b24).Mul(__obf_ba85cb4f962d4b24)
			__obf_3e865660f88e266c = // 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
				NewFromInt(int64(2*__obf_70b1c35e7203372c + 1))
			__obf_3e865660f88e266c = __obf_7f01c31aab0451c6.DivRound(__obf_3e865660f88e266c, __obf_c56f06bc84be0ae8)
			__obf_3a0b292e10e4ffb7 = // comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
				__obf_3a0b292e10e4ffb7.Add(__obf_3e865660f88e266c)

			if __obf_3e865660f88e266c.Abs().Cmp(__obf_17ac81c5d033c6c4) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_f894626ef670eff9 Decimal
		__obf_c2f4ae18447d8eb4 := __obf_c56f06bc84be0ae8*2 + 10

		for __obf_a32504ffc5fafcfe := int32(0); __obf_a32504ffc5fafcfe < __obf_c2f4ae18447d8eb4;
		// exp(a_n)
		__obf_a32504ffc5fafcfe++ {
			__obf_ba85cb4f962d4b24, _ = __obf_3a0b292e10e4ffb7.ExpTaylor(__obf_c56f06bc84be0ae8)
			__obf_7f01c31aab0451c6 = // exp(a_n) - z
				__obf_ba85cb4f962d4b24.Sub(__obf_3b3ef24c3805160d)
			__obf_7f01c31aab0451c6 = // 2 * (exp(a_n) - z)
				__obf_7f01c31aab0451c6.Add(__obf_7f01c31aab0451c6)
			__obf_3e865660f88e266c = // exp(a_n) + z
				__obf_ba85cb4f962d4b24.Add(__obf_3b3ef24c3805160d)
			__obf_ba85cb4f962d4b24 = // 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_7f01c31aab0451c6.DivRound(__obf_3e865660f88e266c, __obf_c56f06bc84be0ae8)
			__obf_3a0b292e10e4ffb7 = // comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_3a0b292e10e4ffb7.Sub(__obf_ba85cb4f962d4b24)

			if __obf_f894626ef670eff9.Add(__obf_ba85cb4f962d4b24).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_ba85cb4f962d4b24.Abs().Cmp(__obf_17ac81c5d033c6c4) <= 0 {
				break
			}
			__obf_f894626ef670eff9 = __obf_ba85cb4f962d4b24
		}
	}
	__obf_3a0b292e10e4ffb7 = __obf_3a0b292e10e4ffb7.Add(__obf_6b01e4b9bb112e56)

	return __obf_3a0b292e10e4ffb7.Round(__obf_02e5a6c8816cdbf7), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_e498fd1d3feac2c4 Decimal) NumDigits() int {
	if __obf_e498fd1d3feac2c4.__obf_980b45647b76adba == nil {
		return 1
	}

	if __obf_e498fd1d3feac2c4.__obf_980b45647b76adba.IsInt64() {
		__obf_adbd09a8342f6c82 := __obf_e498fd1d3feac2c4.
			// restrict fast path to integers with exact conversion to float64
			__obf_980b45647b76adba.Int64()

		if __obf_adbd09a8342f6c82 <= (1<<53) && __obf_adbd09a8342f6c82 >= -(1<<53) {
			if __obf_adbd09a8342f6c82 == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_adbd09a8342f6c82)))) + 1
		}
	}
	__obf_6c66098e433a111c := int(float64(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba.BitLen()) / math.Log2(10))
	__obf_88cbedd911582ea1 := // estimatedNumDigits (lg10) may be off by 1, need to verify
		big.NewInt(int64(__obf_6c66098e433a111c))
	__obf_4bb4ab386c78850f := __obf_88cbedd911582ea1.Exp(__obf_8fd4b9ddee2a2f0a, __obf_88cbedd911582ea1, nil)

	if __obf_e498fd1d3feac2c4.__obf_980b45647b76adba.CmpAbs(__obf_4bb4ab386c78850f) >= 0 {
		return __obf_6c66098e433a111c + 1
	}

	return __obf_6c66098e433a111c
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_e498fd1d3feac2c4 Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_b827ad869fe015ad big.Int
	__obf_209c7ec826089def := new(big.Int).Set(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba)
	for __obf_3b3ef24c3805160d := __obf_e85d6a20e6805469(__obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e); __obf_3b3ef24c3805160d > 0; __obf_3b3ef24c3805160d-- {
		__obf_209c7ec826089def.
			QuoRem(__obf_209c7ec826089def, __obf_8fd4b9ddee2a2f0a, &__obf_b827ad869fe015ad)
		if __obf_b827ad869fe015ad.Cmp(__obf_8e244cd834e11609) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_e85d6a20e6805469(__obf_70b1c35e7203372c int32) int32 {
	if __obf_70b1c35e7203372c < 0 {
		return -__obf_70b1c35e7203372c
	}
	return __obf_70b1c35e7203372c
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_e498fd1d3feac2c4 Decimal) Cmp(__obf_e48c8e69bc461c40 Decimal) int {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()
	__obf_e48c8e69bc461c40.__obf_a167cff196883330()

	if __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e == __obf_e48c8e69bc461c40.__obf_4b33f2812f2bcc7e {
		return __obf_e498fd1d3feac2c4.__obf_980b45647b76adba.Cmp(__obf_e48c8e69bc461c40.__obf_980b45647b76adba)
	}
	__obf_4e9cf9f19030e561, __obf_31e081dce7d8bd2b := RescalePair(__obf_e498fd1d3feac2c4, __obf_e48c8e69bc461c40)

	return __obf_4e9cf9f19030e561.__obf_980b45647b76adba.Cmp(__obf_31e081dce7d8bd2b.

		// Compare compares the numbers represented by d and d2 and returns:
		//
		//	-1 if d <  d2
		//	 0 if d == d2
		//	+1 if d >  d2
		__obf_980b45647b76adba)
}

func (__obf_e498fd1d3feac2c4 Decimal) Compare(__obf_e48c8e69bc461c40 Decimal) int {
	return __obf_e498fd1d3feac2c4.Cmp(__obf_e48c8e69bc461c40)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_e498fd1d3feac2c4 Decimal) Equal(__obf_e48c8e69bc461c40 Decimal) bool {
	return __obf_e498fd1d3feac2c4.Cmp(__obf_e48c8e69bc461c40) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_e498fd1d3feac2c4 Decimal) Equals(__obf_e48c8e69bc461c40 Decimal) bool {
	return __obf_e498fd1d3feac2c4.Equal(__obf_e48c8e69bc461c40)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_e498fd1d3feac2c4 Decimal) GreaterThan(__obf_e48c8e69bc461c40 Decimal) bool {
	return __obf_e498fd1d3feac2c4.Cmp(__obf_e48c8e69bc461c40) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_e498fd1d3feac2c4 Decimal) GreaterThanOrEqual(__obf_e48c8e69bc461c40 Decimal) bool {
	__obf_53979e1b9bb13773 := __obf_e498fd1d3feac2c4.Cmp(__obf_e48c8e69bc461c40)
	return __obf_53979e1b9bb13773 == 1 || __obf_53979e1b9bb13773 == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_e498fd1d3feac2c4 Decimal) LessThan(__obf_e48c8e69bc461c40 Decimal) bool {
	return __obf_e498fd1d3feac2c4.Cmp(__obf_e48c8e69bc461c40) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_e498fd1d3feac2c4 Decimal) LessThanOrEqual(__obf_e48c8e69bc461c40 Decimal) bool {
	__obf_53979e1b9bb13773 := __obf_e498fd1d3feac2c4.Cmp(__obf_e48c8e69bc461c40)
	return __obf_53979e1b9bb13773 == -1 || __obf_53979e1b9bb13773 == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_e498fd1d3feac2c4 Decimal) Sign() int {
	if __obf_e498fd1d3feac2c4.__obf_980b45647b76adba == nil {
		return 0
	}
	return __obf_e498fd1d3feac2c4.

		// IsPositive return
		//
		//	true if d > 0
		//	false if d == 0
		//	false if d < 0
		__obf_980b45647b76adba.Sign()
}

func (__obf_e498fd1d3feac2c4 Decimal) IsPositive() bool {
	return __obf_e498fd1d3feac2c4.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_e498fd1d3feac2c4 Decimal) IsNegative() bool {
	return __obf_e498fd1d3feac2c4.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_e498fd1d3feac2c4 Decimal) IsZero() bool {
	return __obf_e498fd1d3feac2c4.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_e498fd1d3feac2c4 Decimal) Exponent() int32 {
	return __obf_e498fd1d3feac2c4.

		// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
		__obf_4b33f2812f2bcc7e
}

func (__obf_e498fd1d3feac2c4 Decimal) Coefficient() *big.Int {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_e498fd1d3feac2c4.

		// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
		// If coefficient cannot be represented in an int64, the result will be undefined.
		__obf_980b45647b76adba)
}

func (__obf_e498fd1d3feac2c4 Decimal) CoefficientInt64() int64 {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()
	return __obf_e498fd1d3feac2c4.

		// IntPart returns the integer component of the decimal.
		__obf_980b45647b76adba.Int64()
}

func (__obf_e498fd1d3feac2c4 Decimal) IntPart() int64 {
	__obf_3d8aa8e04c317c97 := __obf_e498fd1d3feac2c4.__obf_b297b59a32692f76(0)
	return __obf_3d8aa8e04c317c97.__obf_980b45647b76adba.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_e498fd1d3feac2c4 Decimal) BigInt() *big.Int {
	__obf_3d8aa8e04c317c97 := __obf_e498fd1d3feac2c4.__obf_b297b59a32692f76(0)
	return __obf_3d8aa8e04c317c97.

		// BigFloat returns decimal as BigFloat.
		// Be aware that casting decimal to BigFloat might cause a loss of precision.
		__obf_980b45647b76adba
}

func (__obf_e498fd1d3feac2c4 Decimal) BigFloat() *big.Float {
	__obf_1cee75b7ac144682 := &big.Float{}
	__obf_1cee75b7ac144682.
		SetString(__obf_e498fd1d3feac2c4.String())
	return __obf_1cee75b7ac144682
}

// Rat returns a rational number representation of the decimal.
func (__obf_e498fd1d3feac2c4 Decimal) Rat() *big.Rat {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()
	if __obf_e498fd1d3feac2c4.
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_4b33f2812f2bcc7e <= 0 {
		__obf_e2c07910c9e9de7f := new(big.Int).Exp(__obf_8fd4b9ddee2a2f0a, big.NewInt(-int64(__obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e)), nil)
		return new(big.Rat).SetFrac(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba, __obf_e2c07910c9e9de7f)
	}
	__obf_5ae3215f5ae8edd8 := new(big.Int).Exp(__obf_8fd4b9ddee2a2f0a, big.NewInt(int64(__obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e)), nil)
	__obf_6841a0e73e0bc596 := new(big.Int).Mul(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba, __obf_5ae3215f5ae8edd8)
	return new(big.Rat).SetFrac(__obf_6841a0e73e0bc596,

		// Float64 returns the nearest float64 value for d and a bool indicating
		// whether f represents d exactly.
		// For more details, see the documentation for big.Rat.Float64
		__obf_8ddf5479a62fb2a1)
}

func (__obf_e498fd1d3feac2c4 Decimal) Float64() (__obf_1cee75b7ac144682 float64, __obf_6ad8a44fbf349967 bool) {
	return __obf_e498fd1d3feac2c4.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_e498fd1d3feac2c4 Decimal) InexactFloat64() float64 {
	__obf_1cee75b7ac144682, _ := __obf_e498fd1d3feac2c4.Float64()
	return __obf_1cee75b7ac144682
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
func (__obf_e498fd1d3feac2c4 Decimal) String() string {
	return __obf_e498fd1d3feac2c4.string(true)
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
func (__obf_e498fd1d3feac2c4 Decimal) StringFixed(__obf_95a0813b7432a4cc int32) string {
	__obf_d55c8be8c6e05e70 := __obf_e498fd1d3feac2c4.Round(__obf_95a0813b7432a4cc)
	return __obf_d55c8be8c6e05e70.string(false)
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
func (__obf_e498fd1d3feac2c4 Decimal) StringFixedBank(__obf_95a0813b7432a4cc int32) string {
	__obf_d55c8be8c6e05e70 := __obf_e498fd1d3feac2c4.RoundBank(__obf_95a0813b7432a4cc)
	return __obf_d55c8be8c6e05e70.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_e498fd1d3feac2c4 Decimal) StringFixedCash(__obf_c33f50d41fa0e23a uint8) string {
	__obf_d55c8be8c6e05e70 := __obf_e498fd1d3feac2c4.RoundCash(__obf_c33f50d41fa0e23a)
	return __obf_d55c8be8c6e05e70.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_e498fd1d3feac2c4 Decimal) Round(__obf_95a0813b7432a4cc int32) Decimal {
	if __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e == -__obf_95a0813b7432a4cc {
		return __obf_e498fd1d3feac2c4
	}
	__obf_39698b8ed5518b25 := // truncate to places + 1
		__obf_e498fd1d3feac2c4.__obf_b297b59a32692f76(-__obf_95a0813b7432a4cc - 1)

	// add sign(d) * 0.5
	if __obf_39698b8ed5518b25.__obf_980b45647b76adba.Sign() < 0 {
		__obf_39698b8ed5518b25.__obf_980b45647b76adba.
			Sub(__obf_39698b8ed5518b25.__obf_980b45647b76adba, __obf_40db68f6d1a5418a)
	} else {
		__obf_39698b8ed5518b25.__obf_980b45647b76adba.
			Add(__obf_39698b8ed5518b25.__obf_980b45647b76adba,

				// floor for positive numbers, ceil for negative numbers
				__obf_40db68f6d1a5418a)
	}

	_, __obf_e524d5ae197a0bbb := __obf_39698b8ed5518b25.__obf_980b45647b76adba.DivMod(__obf_39698b8ed5518b25.__obf_980b45647b76adba, __obf_8fd4b9ddee2a2f0a, new(big.Int))
	__obf_39698b8ed5518b25.__obf_4b33f2812f2bcc7e++
	if __obf_39698b8ed5518b25.__obf_980b45647b76adba.Sign() < 0 && __obf_e524d5ae197a0bbb.Cmp(__obf_8e244cd834e11609) != 0 {
		__obf_39698b8ed5518b25.__obf_980b45647b76adba.
			Add(__obf_39698b8ed5518b25.__obf_980b45647b76adba,

				// RoundCeil rounds the decimal towards +infinity.
				//
				// Example:
				//
				//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
				//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
				//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
				//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
				__obf_8ddf5479a62fb2a1)
	}

	return __obf_39698b8ed5518b25
}

func (__obf_e498fd1d3feac2c4 Decimal) RoundCeil(__obf_95a0813b7432a4cc int32) Decimal {
	if __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e >= -__obf_95a0813b7432a4cc {
		return __obf_e498fd1d3feac2c4
	}
	__obf_c8567c8482b2c5b4 := __obf_e498fd1d3feac2c4.__obf_b297b59a32692f76(-__obf_95a0813b7432a4cc)
	if __obf_e498fd1d3feac2c4.Equal(__obf_c8567c8482b2c5b4) {
		return __obf_e498fd1d3feac2c4
	}

	if __obf_e498fd1d3feac2c4.__obf_980b45647b76adba.Sign() > 0 {
		__obf_c8567c8482b2c5b4.__obf_980b45647b76adba.
			Add(__obf_c8567c8482b2c5b4.__obf_980b45647b76adba, __obf_8ddf5479a62fb2a1)
	}

	return __obf_c8567c8482b2c5b4
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_e498fd1d3feac2c4 Decimal) RoundFloor(__obf_95a0813b7432a4cc int32) Decimal {
	if __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e >= -__obf_95a0813b7432a4cc {
		return __obf_e498fd1d3feac2c4
	}
	__obf_c8567c8482b2c5b4 := __obf_e498fd1d3feac2c4.__obf_b297b59a32692f76(-__obf_95a0813b7432a4cc)
	if __obf_e498fd1d3feac2c4.Equal(__obf_c8567c8482b2c5b4) {
		return __obf_e498fd1d3feac2c4
	}

	if __obf_e498fd1d3feac2c4.__obf_980b45647b76adba.Sign() < 0 {
		__obf_c8567c8482b2c5b4.__obf_980b45647b76adba.
			Sub(__obf_c8567c8482b2c5b4.__obf_980b45647b76adba, __obf_8ddf5479a62fb2a1)
	}

	return __obf_c8567c8482b2c5b4
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_e498fd1d3feac2c4 Decimal) RoundUp(__obf_95a0813b7432a4cc int32) Decimal {
	if __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e >= -__obf_95a0813b7432a4cc {
		return __obf_e498fd1d3feac2c4
	}
	__obf_c8567c8482b2c5b4 := __obf_e498fd1d3feac2c4.__obf_b297b59a32692f76(-__obf_95a0813b7432a4cc)
	if __obf_e498fd1d3feac2c4.Equal(__obf_c8567c8482b2c5b4) {
		return __obf_e498fd1d3feac2c4
	}

	if __obf_e498fd1d3feac2c4.__obf_980b45647b76adba.Sign() > 0 {
		__obf_c8567c8482b2c5b4.__obf_980b45647b76adba.
			Add(__obf_c8567c8482b2c5b4.__obf_980b45647b76adba, __obf_8ddf5479a62fb2a1)
	} else if __obf_e498fd1d3feac2c4.__obf_980b45647b76adba.Sign() < 0 {
		__obf_c8567c8482b2c5b4.__obf_980b45647b76adba.
			Sub(__obf_c8567c8482b2c5b4.__obf_980b45647b76adba, __obf_8ddf5479a62fb2a1)
	}

	return __obf_c8567c8482b2c5b4
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_e498fd1d3feac2c4 Decimal) RoundDown(__obf_95a0813b7432a4cc int32) Decimal {
	if __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e >= -__obf_95a0813b7432a4cc {
		return __obf_e498fd1d3feac2c4
	}
	__obf_c8567c8482b2c5b4 := __obf_e498fd1d3feac2c4.__obf_b297b59a32692f76(-__obf_95a0813b7432a4cc)
	if __obf_e498fd1d3feac2c4.Equal(__obf_c8567c8482b2c5b4) {
		return __obf_e498fd1d3feac2c4
	}
	return __obf_c8567c8482b2c5b4
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
func (__obf_e498fd1d3feac2c4 Decimal) RoundBank(__obf_95a0813b7432a4cc int32) Decimal {
	__obf_1ea5813fc65ab676 := __obf_e498fd1d3feac2c4.Round(__obf_95a0813b7432a4cc)
	__obf_834b51f7b19b3057 := __obf_e498fd1d3feac2c4.Sub(__obf_1ea5813fc65ab676).Abs()
	__obf_181a4cf4112664f2 := New(5, -__obf_95a0813b7432a4cc-1)
	if __obf_834b51f7b19b3057.Cmp(__obf_181a4cf4112664f2) == 0 && __obf_1ea5813fc65ab676.__obf_980b45647b76adba.Bit(0) != 0 {
		if __obf_1ea5813fc65ab676.__obf_980b45647b76adba.Sign() < 0 {
			__obf_1ea5813fc65ab676.__obf_980b45647b76adba.
				Add(__obf_1ea5813fc65ab676.__obf_980b45647b76adba, __obf_8ddf5479a62fb2a1)
		} else {
			__obf_1ea5813fc65ab676.__obf_980b45647b76adba.
				Sub(__obf_1ea5813fc65ab676.__obf_980b45647b76adba, __obf_8ddf5479a62fb2a1)
		}
	}

	return __obf_1ea5813fc65ab676
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
func (__obf_e498fd1d3feac2c4 Decimal) RoundCash(__obf_c33f50d41fa0e23a uint8) Decimal {
	var __obf_61d8de6478bf7a06 *big.Int
	switch __obf_c33f50d41fa0e23a {
	case 5:
		__obf_61d8de6478bf7a06 = __obf_20c5d1868eb7c331
	case 10:
		__obf_61d8de6478bf7a06 = __obf_8fd4b9ddee2a2f0a
	case 25:
		__obf_61d8de6478bf7a06 = __obf_49fac42b5d78351d
	case 50:
		__obf_61d8de6478bf7a06 = __obf_013f5033320a21e5
	case 100:
		__obf_61d8de6478bf7a06 = __obf_8ddf5479a62fb2a1
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_c33f50d41fa0e23a))
	}
	__obf_5b137796217a11a9 := Decimal{__obf_980b45647b76adba: __obf_61d8de6478bf7a06}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_e498fd1d3feac2c4.Mul(__obf_5b137796217a11a9).Round(0).Div(__obf_5b137796217a11a9).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_e498fd1d3feac2c4 Decimal) Floor() Decimal {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()

	if __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e >= 0 {
		return __obf_e498fd1d3feac2c4
	}
	__obf_4b33f2812f2bcc7e := big.NewInt(10)
	__obf_4b33f2812f2bcc7e.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_4b33f2812f2bcc7e, big.NewInt(-int64(__obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e)), nil)
	__obf_3b3ef24c3805160d := new(big.Int).Div(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba, __obf_4b33f2812f2bcc7e)
	return Decimal{__obf_980b45647b76adba: __obf_3b3ef24c3805160d,

		// Ceil returns the nearest integer value greater than or equal to d.
		__obf_4b33f2812f2bcc7e: 0}
}

func (__obf_e498fd1d3feac2c4 Decimal) Ceil() Decimal {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()

	if __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e >= 0 {
		return __obf_e498fd1d3feac2c4
	}
	__obf_4b33f2812f2bcc7e := big.NewInt(10)
	__obf_4b33f2812f2bcc7e.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_4b33f2812f2bcc7e, big.NewInt(-int64(__obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e)), nil)
	__obf_3b3ef24c3805160d, __obf_e524d5ae197a0bbb := new(big.Int).DivMod(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba, __obf_4b33f2812f2bcc7e, new(big.Int))
	if __obf_e524d5ae197a0bbb.Cmp(__obf_8e244cd834e11609) != 0 {
		__obf_3b3ef24c3805160d.
			Add(__obf_3b3ef24c3805160d, __obf_8ddf5479a62fb2a1)
	}
	return Decimal{__obf_980b45647b76adba: __obf_3b3ef24c3805160d,

		// Truncate truncates off digits from the number, without rounding.
		//
		// NOTE: precision is the last digit that will not be truncated (must be >= 0).
		//
		// Example:
		//
		//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
		__obf_4b33f2812f2bcc7e: 0}
}

func (__obf_e498fd1d3feac2c4 Decimal) Truncate(__obf_02e5a6c8816cdbf7 int32) Decimal {
	__obf_e498fd1d3feac2c4.__obf_a167cff196883330()
	if __obf_02e5a6c8816cdbf7 >= 0 && -__obf_02e5a6c8816cdbf7 > __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e {
		return __obf_e498fd1d3feac2c4.__obf_b297b59a32692f76(-__obf_02e5a6c8816cdbf7)
	}
	return __obf_e498fd1d3feac2c4
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_e498fd1d3feac2c4 *Decimal) UnmarshalJSON(__obf_cc70c9aaa9dae22e []byte) error {
	if string(__obf_cc70c9aaa9dae22e) == "null" {
		return nil
	}
	__obf_bb7ee06215c1f13d, __obf_3200dc10096fdd51 := __obf_233626bc5b72721d(__obf_cc70c9aaa9dae22e)
	if __obf_3200dc10096fdd51 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_cc70c9aaa9dae22e, __obf_3200dc10096fdd51)
	}
	__obf_3e0a215d3ac5298d, __obf_3200dc10096fdd51 := NewFromString(__obf_bb7ee06215c1f13d)
	*__obf_e498fd1d3feac2c4 = __obf_3e0a215d3ac5298d
	if __obf_3200dc10096fdd51 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_bb7ee06215c1f13d, __obf_3200dc10096fdd51)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_e498fd1d3feac2c4 Decimal) MarshalJSON() ([]byte, error) {
	var __obf_bb7ee06215c1f13d string
	if MarshalJSONWithoutQuotes {
		__obf_bb7ee06215c1f13d = __obf_e498fd1d3feac2c4.String()
	} else {
		__obf_bb7ee06215c1f13d = "\"" + __obf_e498fd1d3feac2c4.String() + "\""
	}
	return []byte(__obf_bb7ee06215c1f13d), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_e498fd1d3feac2c4 *Decimal) UnmarshalBinary(__obf_6ac936a750f74600 []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_6ac936a750f74600) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_6ac936a750f74600, len(__obf_6ac936a750f74600))
	}
	__obf_e498fd1d3feac2c4.

		// Extract the exponent
		__obf_4b33f2812f2bcc7e = int32(binary.BigEndian.Uint32(__obf_6ac936a750f74600[:4]))
	__obf_e498fd1d3feac2c4.

		// Extract the value
		__obf_980b45647b76adba = new(big.Int)
	if __obf_3200dc10096fdd51 := __obf_e498fd1d3feac2c4.__obf_980b45647b76adba.GobDecode(__obf_6ac936a750f74600[4:]); __obf_3200dc10096fdd51 != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_6ac936a750f74600, __obf_3200dc10096fdd51)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_e498fd1d3feac2c4 Decimal) MarshalBinary() (__obf_6ac936a750f74600 []byte, __obf_3200dc10096fdd51 error) {
	// exp is written first, but encode value first to know output size
	var __obf_ff2c34025efd3c97 []byte
	if __obf_ff2c34025efd3c97, __obf_3200dc10096fdd51 = __obf_e498fd1d3feac2c4.__obf_980b45647b76adba.GobEncode(); __obf_3200dc10096fdd51 != nil {
		return nil, __obf_3200dc10096fdd51
	}
	__obf_3c6f2897876fc500 := // Write the exponent in front, since it's a fixed size
		make([]byte, 4, len(__obf_ff2c34025efd3c97)+4)
	binary.BigEndian.PutUint32(__obf_3c6f2897876fc500, uint32(__obf_e498fd1d3feac2c4.

		// Return the byte array
		__obf_4b33f2812f2bcc7e))

	return append(__obf_3c6f2897876fc500, __obf_ff2c34025efd3c97...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_e498fd1d3feac2c4 *Decimal) Scan(__obf_980b45647b76adba any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_f94a661e904c6c4f := __obf_980b45647b76adba.(type) {

	case float32:
		*__obf_e498fd1d3feac2c4 = NewFromFloat(float64(__obf_f94a661e904c6c4f))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_e498fd1d3feac2c4 = NewFromFloat(__obf_f94a661e904c6c4f)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_e498fd1d3feac2c4 = New(__obf_f94a661e904c6c4f, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_e498fd1d3feac2c4 = NewFromUint64(__obf_f94a661e904c6c4f)
		return nil

	default:
		__obf_bb7ee06215c1f13d,
			// default is trying to interpret value stored as string
			__obf_3200dc10096fdd51 := __obf_233626bc5b72721d(__obf_f94a661e904c6c4f)
		if __obf_3200dc10096fdd51 != nil {
			return __obf_3200dc10096fdd51
		}
		*__obf_e498fd1d3feac2c4, __obf_3200dc10096fdd51 = NewFromString(__obf_bb7ee06215c1f13d)
		return __obf_3200dc10096fdd51
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_e498fd1d3feac2c4 Decimal) Value() (driver.Value, error) {
	return __obf_e498fd1d3feac2c4.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_e498fd1d3feac2c4 *Decimal) UnmarshalText(__obf_a7e519919672fad0 []byte) error {
	__obf_bb7ee06215c1f13d := string(__obf_a7e519919672fad0)
	__obf_2aecda69462f7dfa, __obf_3200dc10096fdd51 := NewFromString(__obf_bb7ee06215c1f13d)
	*__obf_e498fd1d3feac2c4 = __obf_2aecda69462f7dfa
	if __obf_3200dc10096fdd51 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_bb7ee06215c1f13d, __obf_3200dc10096fdd51)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_e498fd1d3feac2c4 Decimal) MarshalText() (__obf_a7e519919672fad0 []byte, __obf_3200dc10096fdd51 error) {
	return []byte(__obf_e498fd1d3feac2c4.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_e498fd1d3feac2c4 Decimal) GobEncode() ([]byte, error) {
	return __obf_e498fd1d3feac2c4.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_e498fd1d3feac2c4 *Decimal) GobDecode(__obf_6ac936a750f74600 []byte) error {
	return __obf_e498fd1d3feac2c4.UnmarshalBinary(__obf_6ac936a750f74600)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_e498fd1d3feac2c4 Decimal) StringScaled(__obf_4b33f2812f2bcc7e int32) string {
	return __obf_e498fd1d3feac2c4.__obf_b297b59a32692f76(__obf_4b33f2812f2bcc7e).String()
}

func (__obf_e498fd1d3feac2c4 Decimal) string(__obf_e3f2f56799b074d8 bool) string {
	if __obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e >= 0 {
		return __obf_e498fd1d3feac2c4.__obf_b297b59a32692f76(0).__obf_980b45647b76adba.String()
	}
	__obf_e85d6a20e6805469 := new(big.Int).Abs(__obf_e498fd1d3feac2c4.__obf_980b45647b76adba)
	__obf_bb7ee06215c1f13d := __obf_e85d6a20e6805469.String()

	var __obf_67b1941cda1eefde, __obf_bbfdc1ee5dabe412 string
	__obf_24f5f06b168fc9c4 := // NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
		// and you are on a 32-bit machine. Won't fix this super-edge case.
		int(__obf_e498fd1d3feac2c4.__obf_4b33f2812f2bcc7e)
	if len(__obf_bb7ee06215c1f13d) > -__obf_24f5f06b168fc9c4 {
		__obf_67b1941cda1eefde = __obf_bb7ee06215c1f13d[:len(__obf_bb7ee06215c1f13d)+__obf_24f5f06b168fc9c4]
		__obf_bbfdc1ee5dabe412 = __obf_bb7ee06215c1f13d[len(__obf_bb7ee06215c1f13d)+__obf_24f5f06b168fc9c4:]
	} else {
		__obf_67b1941cda1eefde = "0"
		__obf_e80a4388d4f36593 := -__obf_24f5f06b168fc9c4 - len(__obf_bb7ee06215c1f13d)
		__obf_bbfdc1ee5dabe412 = strings.Repeat("0", __obf_e80a4388d4f36593) + __obf_bb7ee06215c1f13d
	}

	if __obf_e3f2f56799b074d8 {
		__obf_a32504ffc5fafcfe := len(__obf_bbfdc1ee5dabe412) - 1
		for ; __obf_a32504ffc5fafcfe >= 0; __obf_a32504ffc5fafcfe-- {
			if __obf_bbfdc1ee5dabe412[__obf_a32504ffc5fafcfe] != '0' {
				break
			}
		}
		__obf_bbfdc1ee5dabe412 = __obf_bbfdc1ee5dabe412[:__obf_a32504ffc5fafcfe+1]
	}
	__obf_5070a4220b083e32 := __obf_67b1941cda1eefde
	if len(__obf_bbfdc1ee5dabe412) > 0 {
		__obf_5070a4220b083e32 += "." + __obf_bbfdc1ee5dabe412
	}

	if __obf_e498fd1d3feac2c4.__obf_980b45647b76adba.Sign() < 0 {
		return "-" + __obf_5070a4220b083e32
	}

	return __obf_5070a4220b083e32
}

func (__obf_e498fd1d3feac2c4 *Decimal) __obf_a167cff196883330() {
	if __obf_e498fd1d3feac2c4.__obf_980b45647b76adba == nil {
		__obf_e498fd1d3feac2c4.__obf_980b45647b76adba = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_fb5fd6aab5d398d0 Decimal, __obf_b90230115ec92d35 ...Decimal) Decimal {
	__obf_29a7593c68f7592f := __obf_fb5fd6aab5d398d0
	for _, __obf_b6f6d5e3672d16b3 := range __obf_b90230115ec92d35 {
		if __obf_b6f6d5e3672d16b3.Cmp(__obf_29a7593c68f7592f) < 0 {
			__obf_29a7593c68f7592f = __obf_b6f6d5e3672d16b3
		}
	}
	return __obf_29a7593c68f7592f
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_fb5fd6aab5d398d0 Decimal, __obf_b90230115ec92d35 ...Decimal) Decimal {
	__obf_29a7593c68f7592f := __obf_fb5fd6aab5d398d0
	for _, __obf_b6f6d5e3672d16b3 := range __obf_b90230115ec92d35 {
		if __obf_b6f6d5e3672d16b3.Cmp(__obf_29a7593c68f7592f) > 0 {
			__obf_29a7593c68f7592f = __obf_b6f6d5e3672d16b3
		}
	}
	return __obf_29a7593c68f7592f
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_fb5fd6aab5d398d0 Decimal, __obf_b90230115ec92d35 ...Decimal) Decimal {
	__obf_bc9baae2cc60df49 := __obf_fb5fd6aab5d398d0
	for _, __obf_b6f6d5e3672d16b3 := range __obf_b90230115ec92d35 {
		__obf_bc9baae2cc60df49 = __obf_bc9baae2cc60df49.Add(__obf_b6f6d5e3672d16b3)
	}

	return __obf_bc9baae2cc60df49
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_fb5fd6aab5d398d0 Decimal, __obf_b90230115ec92d35 ...Decimal) Decimal {
	__obf_30a147404e4df4c1 := New(int64(len(__obf_b90230115ec92d35)+1), 0)
	__obf_422f54d84cbf9f8b := Sum(__obf_fb5fd6aab5d398d0, __obf_b90230115ec92d35...)
	return __obf_422f54d84cbf9f8b.Div(__obf_30a147404e4df4c1)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_e7022a3c77dbba3b Decimal, __obf_e48c8e69bc461c40 Decimal) (Decimal, Decimal) {
	__obf_e7022a3c77dbba3b.__obf_a167cff196883330()
	__obf_e48c8e69bc461c40.__obf_a167cff196883330()

	if __obf_e7022a3c77dbba3b.__obf_4b33f2812f2bcc7e < __obf_e48c8e69bc461c40.__obf_4b33f2812f2bcc7e {
		return __obf_e7022a3c77dbba3b, __obf_e48c8e69bc461c40.__obf_b297b59a32692f76(__obf_e7022a3c77dbba3b.__obf_4b33f2812f2bcc7e)
	} else if __obf_e7022a3c77dbba3b.__obf_4b33f2812f2bcc7e > __obf_e48c8e69bc461c40.__obf_4b33f2812f2bcc7e {
		return __obf_e7022a3c77dbba3b.__obf_b297b59a32692f76(__obf_e48c8e69bc461c40.__obf_4b33f2812f2bcc7e), __obf_e48c8e69bc461c40
	}

	return __obf_e7022a3c77dbba3b, __obf_e48c8e69bc461c40
}

func __obf_233626bc5b72721d(__obf_980b45647b76adba any) (string, error) {
	var __obf_09208a81dbf9a465 []byte

	switch __obf_f94a661e904c6c4f := __obf_980b45647b76adba.(type) {
	case string:
		__obf_09208a81dbf9a465 = []byte(__obf_f94a661e904c6c4f)
	case []byte:
		__obf_09208a81dbf9a465 = __obf_f94a661e904c6c4f
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_980b45647b76adba,

			// If the amount is quoted, strip the quotes
			__obf_980b45647b76adba)
	}

	if len(__obf_09208a81dbf9a465) > 2 && __obf_09208a81dbf9a465[0] == '"' && __obf_09208a81dbf9a465[len(__obf_09208a81dbf9a465)-1] == '"' {
		__obf_09208a81dbf9a465 = __obf_09208a81dbf9a465[1 : len(__obf_09208a81dbf9a465)-1]
	}
	return string(__obf_09208a81dbf9a465), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the datautil.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_e498fd1d3feac2c4 Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_e498fd1d3feac2c4,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_e498fd1d3feac2c4 *NullDecimal) Scan(__obf_980b45647b76adba any) error {
	if __obf_980b45647b76adba == nil {
		__obf_e498fd1d3feac2c4.
			Valid = false
		return nil
	}
	__obf_e498fd1d3feac2c4.
		Valid = true
	return __obf_e498fd1d3feac2c4.Decimal.Scan(__obf_980b45647b76adba)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_e498fd1d3feac2c4 NullDecimal) Value() (driver.Value, error) {
	if !__obf_e498fd1d3feac2c4.Valid {
		return nil, nil
	}
	return __obf_e498fd1d3feac2c4.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_e498fd1d3feac2c4 *NullDecimal) UnmarshalJSON(__obf_cc70c9aaa9dae22e []byte) error {
	if string(__obf_cc70c9aaa9dae22e) == "null" {
		__obf_e498fd1d3feac2c4.
			Valid = false
		return nil
	}
	__obf_e498fd1d3feac2c4.
		Valid = true
	return __obf_e498fd1d3feac2c4.Decimal.UnmarshalJSON(__obf_cc70c9aaa9dae22e)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_e498fd1d3feac2c4 NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_e498fd1d3feac2c4.Valid {
		return []byte("null"), nil
	}
	return __obf_e498fd1d3feac2c4.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_e498fd1d3feac2c4 *NullDecimal) UnmarshalText(__obf_a7e519919672fad0 []byte) error {
	__obf_bb7ee06215c1f13d := string(__obf_a7e519919672fad0)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_bb7ee06215c1f13d == "" {
		__obf_e498fd1d3feac2c4.
			Valid = false
		return nil
	}
	if __obf_3200dc10096fdd51 := __obf_e498fd1d3feac2c4.Decimal.UnmarshalText(__obf_a7e519919672fad0); __obf_3200dc10096fdd51 != nil {
		__obf_e498fd1d3feac2c4.
			Valid = false
		return __obf_3200dc10096fdd51
	}
	__obf_e498fd1d3feac2c4.
		Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_e498fd1d3feac2c4 NullDecimal) MarshalText() (__obf_a7e519919672fad0 []byte, __obf_3200dc10096fdd51 error) {
	if !__obf_e498fd1d3feac2c4.Valid {
		return []byte{}, nil
	}
	return __obf_e498fd1d3feac2c4.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_e498fd1d3feac2c4 Decimal) Atan() Decimal {
	if __obf_e498fd1d3feac2c4.Equal(NewFromFloat(0.0)) {
		return __obf_e498fd1d3feac2c4
	}
	if __obf_e498fd1d3feac2c4.GreaterThan(NewFromFloat(0.0)) {
		return __obf_e498fd1d3feac2c4.__obf_abb2ea115c1fecb2()
	}
	return __obf_e498fd1d3feac2c4.Neg().__obf_abb2ea115c1fecb2().Neg()
}

func (__obf_e498fd1d3feac2c4 Decimal) __obf_aa748c2f31637f3a() Decimal {
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
	__obf_3b3ef24c3805160d := __obf_e498fd1d3feac2c4.Mul(__obf_e498fd1d3feac2c4)
	__obf_4616086baf4cb5fe := P0.Mul(__obf_3b3ef24c3805160d).Add(P1).Mul(__obf_3b3ef24c3805160d).Add(P2).Mul(__obf_3b3ef24c3805160d).Add(P3).Mul(__obf_3b3ef24c3805160d).Add(P4).Mul(__obf_3b3ef24c3805160d)
	__obf_e7210d49ad636f15 := __obf_3b3ef24c3805160d.Add(Q0).Mul(__obf_3b3ef24c3805160d).Add(Q1).Mul(__obf_3b3ef24c3805160d).Add(Q2).Mul(__obf_3b3ef24c3805160d).Add(Q3).Mul(__obf_3b3ef24c3805160d).Add(Q4)
	__obf_3b3ef24c3805160d = __obf_4616086baf4cb5fe.Div(__obf_e7210d49ad636f15)
	__obf_3b3ef24c3805160d = __obf_e498fd1d3feac2c4.Mul(__obf_3b3ef24c3805160d).Add(__obf_e498fd1d3feac2c4)
	return __obf_3b3ef24c3805160d
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_e498fd1d3feac2c4 Decimal) __obf_abb2ea115c1fecb2() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)
	__obf_3a30440738843731 := // tan(3*pi/8)
		NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_e498fd1d3feac2c4.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_e498fd1d3feac2c4.__obf_aa748c2f31637f3a()
	}
	if __obf_e498fd1d3feac2c4.GreaterThan(Tan3pio8) {
		return __obf_3a30440738843731.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_e498fd1d3feac2c4).__obf_aa748c2f31637f3a()).Add(Morebits)
	}
	return __obf_3a30440738843731.Div(NewFromFloat(4.0)).Add((__obf_e498fd1d3feac2c4.Sub(NewFromFloat(1.0)).Div(__obf_e498fd1d3feac2c4.Add(NewFromFloat(1.0)))).__obf_aa748c2f31637f3a()).Add(NewFromFloat(0.5).Mul(Morebits))
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
func (__obf_e498fd1d3feac2c4 Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_e498fd1d3feac2c4.Equal(NewFromFloat(0.0)) {
		return __obf_e498fd1d3feac2c4
	}
	__obf_ae515fc8976e420b := // make argument positive but save the sign
		false
	if __obf_e498fd1d3feac2c4.LessThan(NewFromFloat(0.0)) {
		__obf_e498fd1d3feac2c4 = __obf_e498fd1d3feac2c4.Neg()
		__obf_ae515fc8976e420b = true
	}
	__obf_6f021b32b781c111 := __obf_e498fd1d3feac2c4.Mul(M4PI).IntPart()
	__obf_efa18593310aecde := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_6f021b32b781c111)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_6f021b32b781c111&1 == 1 {
		__obf_6f021b32b781c111++
		__obf_efa18593310aecde = __obf_efa18593310aecde.Add(NewFromFloat(1.0))
	}
	__obf_6f021b32b781c111 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_6f021b32b781c111 > 3 {
		__obf_ae515fc8976e420b = !__obf_ae515fc8976e420b
		__obf_6f021b32b781c111 -= 4
	}
	__obf_3b3ef24c3805160d := __obf_e498fd1d3feac2c4.Sub(__obf_efa18593310aecde.Mul(PI4A)).Sub(__obf_efa18593310aecde.Mul(PI4B)).Sub(__obf_efa18593310aecde.Mul(PI4C))
	__obf_21220ad6f032b9a8 := // Extended precision modular arithmetic
		__obf_3b3ef24c3805160d.Mul(__obf_3b3ef24c3805160d)

	if __obf_6f021b32b781c111 == 1 || __obf_6f021b32b781c111 == 2 {
		__obf_f85c5aae9b59cb4e := __obf_21220ad6f032b9a8.Mul(__obf_21220ad6f032b9a8).Mul(_cos[0].Mul(__obf_21220ad6f032b9a8).Add(_cos[1]).Mul(__obf_21220ad6f032b9a8).Add(_cos[2]).Mul(__obf_21220ad6f032b9a8).Add(_cos[3]).Mul(__obf_21220ad6f032b9a8).Add(_cos[4]).Mul(__obf_21220ad6f032b9a8).Add(_cos[5]))
		__obf_efa18593310aecde = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_21220ad6f032b9a8)).Add(__obf_f85c5aae9b59cb4e)
	} else {
		__obf_efa18593310aecde = __obf_3b3ef24c3805160d.Add(__obf_3b3ef24c3805160d.Mul(__obf_21220ad6f032b9a8).Mul(_sin[0].Mul(__obf_21220ad6f032b9a8).Add(_sin[1]).Mul(__obf_21220ad6f032b9a8).Add(_sin[2]).Mul(__obf_21220ad6f032b9a8).Add(_sin[3]).Mul(__obf_21220ad6f032b9a8).Add(_sin[4]).Mul(__obf_21220ad6f032b9a8).Add(_sin[5])))
	}
	if __obf_ae515fc8976e420b {
		__obf_efa18593310aecde = __obf_efa18593310aecde.Neg()
	}
	return __obf_efa18593310aecde
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
func (__obf_e498fd1d3feac2c4 Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)  // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)  // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15) // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125)
	__obf_ae515fc8976e420b := // 4/pi

		// make argument positive
		false
	if __obf_e498fd1d3feac2c4.LessThan(NewFromFloat(0.0)) {
		__obf_e498fd1d3feac2c4 = __obf_e498fd1d3feac2c4.Neg()
	}
	__obf_6f021b32b781c111 := __obf_e498fd1d3feac2c4.Mul(M4PI).IntPart()
	__obf_efa18593310aecde := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_6f021b32b781c111)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_6f021b32b781c111&1 == 1 {
		__obf_6f021b32b781c111++
		__obf_efa18593310aecde = __obf_efa18593310aecde.Add(NewFromFloat(1.0))
	}
	__obf_6f021b32b781c111 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_6f021b32b781c111 > 3 {
		__obf_ae515fc8976e420b = !__obf_ae515fc8976e420b
		__obf_6f021b32b781c111 -= 4
	}
	if __obf_6f021b32b781c111 > 1 {
		__obf_ae515fc8976e420b = !__obf_ae515fc8976e420b
	}
	__obf_3b3ef24c3805160d := __obf_e498fd1d3feac2c4.Sub(__obf_efa18593310aecde.Mul(PI4A)).Sub(__obf_efa18593310aecde.Mul(PI4B)).Sub(__obf_efa18593310aecde.Mul(PI4C))
	__obf_21220ad6f032b9a8 := // Extended precision modular arithmetic
		__obf_3b3ef24c3805160d.Mul(__obf_3b3ef24c3805160d)

	if __obf_6f021b32b781c111 == 1 || __obf_6f021b32b781c111 == 2 {
		__obf_efa18593310aecde = __obf_3b3ef24c3805160d.Add(__obf_3b3ef24c3805160d.Mul(__obf_21220ad6f032b9a8).Mul(_sin[0].Mul(__obf_21220ad6f032b9a8).Add(_sin[1]).Mul(__obf_21220ad6f032b9a8).Add(_sin[2]).Mul(__obf_21220ad6f032b9a8).Add(_sin[3]).Mul(__obf_21220ad6f032b9a8).Add(_sin[4]).Mul(__obf_21220ad6f032b9a8).Add(_sin[5])))
	} else {
		__obf_f85c5aae9b59cb4e := __obf_21220ad6f032b9a8.Mul(__obf_21220ad6f032b9a8).Mul(_cos[0].Mul(__obf_21220ad6f032b9a8).Add(_cos[1]).Mul(__obf_21220ad6f032b9a8).Add(_cos[2]).Mul(__obf_21220ad6f032b9a8).Add(_cos[3]).Mul(__obf_21220ad6f032b9a8).Add(_cos[4]).Mul(__obf_21220ad6f032b9a8).Add(_cos[5]))
		__obf_efa18593310aecde = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_21220ad6f032b9a8)).Add(__obf_f85c5aae9b59cb4e)
	}
	if __obf_ae515fc8976e420b {
		__obf_efa18593310aecde = __obf_efa18593310aecde.Neg()
	}
	return __obf_efa18593310aecde
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
func (__obf_e498fd1d3feac2c4 Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_e498fd1d3feac2c4.Equal(NewFromFloat(0.0)) {
		return __obf_e498fd1d3feac2c4
	}
	__obf_ae515fc8976e420b := // make argument positive but save the sign
		false
	if __obf_e498fd1d3feac2c4.LessThan(NewFromFloat(0.0)) {
		__obf_e498fd1d3feac2c4 = __obf_e498fd1d3feac2c4.Neg()
		__obf_ae515fc8976e420b = true
	}
	__obf_6f021b32b781c111 := __obf_e498fd1d3feac2c4.Mul(M4PI).IntPart()
	__obf_efa18593310aecde := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_6f021b32b781c111)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_6f021b32b781c111&1 == 1 {
		__obf_6f021b32b781c111++
		__obf_efa18593310aecde = __obf_efa18593310aecde.Add(NewFromFloat(1.0))
	}
	__obf_3b3ef24c3805160d := __obf_e498fd1d3feac2c4.Sub(__obf_efa18593310aecde.Mul(PI4A)).Sub(__obf_efa18593310aecde.Mul(PI4B)).Sub(__obf_efa18593310aecde.Mul(PI4C))
	__obf_21220ad6f032b9a8 := // Extended precision modular arithmetic
		__obf_3b3ef24c3805160d.Mul(__obf_3b3ef24c3805160d)

	if __obf_21220ad6f032b9a8.GreaterThan(NewFromFloat(1e-14)) {
		__obf_f85c5aae9b59cb4e := __obf_21220ad6f032b9a8.Mul(_tanP[0].Mul(__obf_21220ad6f032b9a8).Add(_tanP[1]).Mul(__obf_21220ad6f032b9a8).Add(_tanP[2]))
		__obf_ee5e2564a3975b5d := __obf_21220ad6f032b9a8.Add(_tanQ[1]).Mul(__obf_21220ad6f032b9a8).Add(_tanQ[2]).Mul(__obf_21220ad6f032b9a8).Add(_tanQ[3]).Mul(__obf_21220ad6f032b9a8).Add(_tanQ[4])
		__obf_efa18593310aecde = __obf_3b3ef24c3805160d.Add(__obf_3b3ef24c3805160d.Mul(__obf_f85c5aae9b59cb4e.Div(__obf_ee5e2564a3975b5d)))
	} else {
		__obf_efa18593310aecde = __obf_3b3ef24c3805160d
	}
	if __obf_6f021b32b781c111&2 == 2 {
		__obf_efa18593310aecde = NewFromFloat(-1.0).Div(__obf_efa18593310aecde)
	}
	if __obf_ae515fc8976e420b {
		__obf_efa18593310aecde = __obf_efa18593310aecde.Neg()
	}
	return __obf_efa18593310aecde
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

type __obf_3e0a215d3ac5298d struct {
	__obf_e498fd1d3feac2c4 [800]byte
	__obf_7a088d2e01e4cd05 int  // digits, big-endian representation
	__obf_164d07da7c832f2c int  // number of digits used
	__obf_20103b61d9019c90 bool // decimal point
	__obf_3a6f83bf0d7dafe4 bool // negative flag
	// discarded nonzero digits beyond d[:nd]
}

func (__obf_55d13866fd49a7b4 *__obf_3e0a215d3ac5298d) String() string {
	__obf_70b1c35e7203372c := 10 + __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05
	if __obf_55d13866fd49a7b4.__obf_164d07da7c832f2c > 0 {
		__obf_70b1c35e7203372c += __obf_55d13866fd49a7b4.__obf_164d07da7c832f2c
	}
	if __obf_55d13866fd49a7b4.__obf_164d07da7c832f2c < 0 {
		__obf_70b1c35e7203372c += -__obf_55d13866fd49a7b4.__obf_164d07da7c832f2c
	}
	__obf_95d0e7ba75c6df01 := make([]byte, __obf_70b1c35e7203372c)
	__obf_f85c5aae9b59cb4e := 0
	switch {
	case __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 == 0:
		return "0"

	case __obf_55d13866fd49a7b4.
		// zeros fill space between decimal point and digits
		__obf_164d07da7c832f2c <= 0:
		__obf_95d0e7ba75c6df01[__obf_f85c5aae9b59cb4e] = '0'
		__obf_f85c5aae9b59cb4e++
		__obf_95d0e7ba75c6df01[__obf_f85c5aae9b59cb4e] = '.'
		__obf_f85c5aae9b59cb4e++
		__obf_f85c5aae9b59cb4e += __obf_a039bea22f17a712(__obf_95d0e7ba75c6df01[__obf_f85c5aae9b59cb4e : __obf_f85c5aae9b59cb4e+-__obf_55d13866fd49a7b4.__obf_164d07da7c832f2c])
		__obf_f85c5aae9b59cb4e += copy(__obf_95d0e7ba75c6df01[__obf_f85c5aae9b59cb4e:], __obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[0:__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05])

	case __obf_55d13866fd49a7b4.
		// decimal point in middle of digits
		__obf_164d07da7c832f2c < __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05:
		__obf_f85c5aae9b59cb4e += copy(__obf_95d0e7ba75c6df01[__obf_f85c5aae9b59cb4e:], __obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[0:__obf_55d13866fd49a7b4.__obf_164d07da7c832f2c])
		__obf_95d0e7ba75c6df01[__obf_f85c5aae9b59cb4e] = '.'
		__obf_f85c5aae9b59cb4e++
		__obf_f85c5aae9b59cb4e += copy(__obf_95d0e7ba75c6df01[__obf_f85c5aae9b59cb4e:], __obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_55d13866fd49a7b4.__obf_164d07da7c832f2c:

		// zeros fill space between digits and decimal point
		__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05])

	default:
		__obf_f85c5aae9b59cb4e += copy(__obf_95d0e7ba75c6df01[__obf_f85c5aae9b59cb4e:], __obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[0:__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05])
		__obf_f85c5aae9b59cb4e += __obf_a039bea22f17a712(__obf_95d0e7ba75c6df01[__obf_f85c5aae9b59cb4e : __obf_f85c5aae9b59cb4e+__obf_55d13866fd49a7b4.__obf_164d07da7c832f2c-__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05])
	}
	return string(__obf_95d0e7ba75c6df01[0:__obf_f85c5aae9b59cb4e])
}

func __obf_a039bea22f17a712(__obf_0be6265144bce178 []byte) int {
	for __obf_a32504ffc5fafcfe := range __obf_0be6265144bce178 {
		__obf_0be6265144bce178[__obf_a32504ffc5fafcfe] = '0'
	}
	return len(__obf_0be6265144bce178)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_800f8c75bc7c90e0(__obf_55d13866fd49a7b4 *__obf_3e0a215d3ac5298d) {
	for __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 > 0 && __obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05-1] == '0' {
		__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05--
	}
	if __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 == 0 {
		__obf_55d13866fd49a7b4.

			// Assign v to a.
			__obf_164d07da7c832f2c = 0
	}
}

func (__obf_55d13866fd49a7b4 *__obf_3e0a215d3ac5298d) Assign(__obf_f94a661e904c6c4f uint64) {
	var __obf_95d0e7ba75c6df01 [24]byte
	__obf_70b1c35e7203372c := // Write reversed decimal in buf.
		0
	for __obf_f94a661e904c6c4f > 0 {
		__obf_a296c2e3f9528c45 := __obf_f94a661e904c6c4f / 10
		__obf_f94a661e904c6c4f -= 10 * __obf_a296c2e3f9528c45
		__obf_95d0e7ba75c6df01[__obf_70b1c35e7203372c] = byte(__obf_f94a661e904c6c4f + '0')
		__obf_70b1c35e7203372c++
		__obf_f94a661e904c6c4f = __obf_a296c2e3f9528c45
	}
	__obf_55d13866fd49a7b4.

		// Reverse again to produce forward decimal in a.d.
		__obf_7a088d2e01e4cd05 = 0
	for __obf_70b1c35e7203372c--; __obf_70b1c35e7203372c >= 0; __obf_70b1c35e7203372c-- {
		__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05] = __obf_95d0e7ba75c6df01[__obf_70b1c35e7203372c]
		__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05++
	}
	__obf_55d13866fd49a7b4.__obf_164d07da7c832f2c = __obf_55d13866fd49a7b4.

		// Maximum shift that we can do in one pass without overflow.
		// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
		__obf_7a088d2e01e4cd05
	__obf_800f8c75bc7c90e0(__obf_55d13866fd49a7b4)
}

const __obf_73cc5297a4513b3f = 32 << (^uint(0) >> 63)
const __obf_b51cc6d999d130c7 = __obf_73cc5297a4513b3f - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_9ef64b0acb08f65b(__obf_55d13866fd49a7b4 *__obf_3e0a215d3ac5298d, __obf_0fefbf9a88cfc16b uint) {
	__obf_b827ad869fe015ad := 0
	__obf_f85c5aae9b59cb4e := // read pointer
		0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_70b1c35e7203372c uint
	for ; __obf_70b1c35e7203372c>>__obf_0fefbf9a88cfc16b == 0; __obf_b827ad869fe015ad++ {
		if __obf_b827ad869fe015ad >= __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 {
			if __obf_70b1c35e7203372c == 0 {
				__obf_55d13866fd49a7b4.
					// a == 0; shouldn't get here, but handle anyway.
					__obf_7a088d2e01e4cd05 = 0
				return
			}
			for __obf_70b1c35e7203372c>>__obf_0fefbf9a88cfc16b == 0 {
				__obf_70b1c35e7203372c = __obf_70b1c35e7203372c * 10
				__obf_b827ad869fe015ad++
			}
			break
		}
		__obf_d0f3a30e016877c1 := uint(__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_b827ad869fe015ad])
		__obf_70b1c35e7203372c = __obf_70b1c35e7203372c*10 + __obf_d0f3a30e016877c1 - '0'
	}
	__obf_55d13866fd49a7b4.__obf_164d07da7c832f2c -= __obf_b827ad869fe015ad - 1

	var __obf_19501d889006ff82 uint = (1 << __obf_0fefbf9a88cfc16b) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_b827ad869fe015ad < __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05; __obf_b827ad869fe015ad++ {
		__obf_d0f3a30e016877c1 := uint(__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_b827ad869fe015ad])
		__obf_a5cda4bf92925ecf := __obf_70b1c35e7203372c >> __obf_0fefbf9a88cfc16b
		__obf_70b1c35e7203372c &= __obf_19501d889006ff82
		__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_f85c5aae9b59cb4e] = byte(__obf_a5cda4bf92925ecf + '0')
		__obf_f85c5aae9b59cb4e++
		__obf_70b1c35e7203372c = __obf_70b1c35e7203372c*10 + __obf_d0f3a30e016877c1 - '0'
	}

	// Put down extra digits.
	for __obf_70b1c35e7203372c > 0 {
		__obf_a5cda4bf92925ecf := __obf_70b1c35e7203372c >> __obf_0fefbf9a88cfc16b
		__obf_70b1c35e7203372c &= __obf_19501d889006ff82
		if __obf_f85c5aae9b59cb4e < len(__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4) {
			__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_f85c5aae9b59cb4e] = byte(__obf_a5cda4bf92925ecf + '0')
			__obf_f85c5aae9b59cb4e++
		} else if __obf_a5cda4bf92925ecf > 0 {
			__obf_55d13866fd49a7b4.__obf_3a6f83bf0d7dafe4 = true
		}
		__obf_70b1c35e7203372c = __obf_70b1c35e7203372c * 10
	}
	__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 = __obf_f85c5aae9b59cb4e

	// Cheat sheet for left shift: table indexed by shift count giving
	// number of new digits that will be introduced by that shift.
	//
	// For example, leftcheats[4] = {2, "625"}.  That means that
	// if we are shifting by 4 (multiplying by 16), it will add 2 digits
	// when the string prefix is "625" through "999", and one fewer digit
	// if the string prefix is "000" through "624".
	//
	// Credit for this trick goes to Ken.
	__obf_800f8c75bc7c90e0(__obf_55d13866fd49a7b4)
}

type __obf_286e69d226617442 struct {
	__obf_c47fe4ed44a7322b int
	__obf_d0b99eb1688117dc string // number of new digits
	// minus one digit if original < a.
}

var __obf_d5a34869205af9a3 = []__obf_286e69d226617442{
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
func __obf_85b6f910353450fa(__obf_046424310fede70c []byte, __obf_584552b633cee99d string) bool {
	for __obf_a32504ffc5fafcfe := 0; __obf_a32504ffc5fafcfe < len(__obf_584552b633cee99d); __obf_a32504ffc5fafcfe++ {
		if __obf_a32504ffc5fafcfe >= len(__obf_046424310fede70c) {
			return true
		}
		if __obf_046424310fede70c[__obf_a32504ffc5fafcfe] != __obf_584552b633cee99d[__obf_a32504ffc5fafcfe] {
			return __obf_046424310fede70c[__obf_a32504ffc5fafcfe] < __obf_584552b633cee99d[__obf_a32504ffc5fafcfe]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_291f816b97326a6a(__obf_55d13866fd49a7b4 *__obf_3e0a215d3ac5298d, __obf_0fefbf9a88cfc16b uint) {
	__obf_c47fe4ed44a7322b := __obf_d5a34869205af9a3[__obf_0fefbf9a88cfc16b].__obf_c47fe4ed44a7322b
	if __obf_85b6f910353450fa(__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[0:__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05], __obf_d5a34869205af9a3[__obf_0fefbf9a88cfc16b].__obf_d0b99eb1688117dc) {
		__obf_c47fe4ed44a7322b--
	}
	__obf_b827ad869fe015ad := __obf_55d13866fd49a7b4. // read index
								__obf_7a088d2e01e4cd05
	__obf_f85c5aae9b59cb4e := __obf_55d13866fd49a7b4. // write index
								__obf_7a088d2e01e4cd05 + __obf_c47fe4ed44a7322b

	// Pick up a digit, put down a digit.
	var __obf_70b1c35e7203372c uint
	for __obf_b827ad869fe015ad--; __obf_b827ad869fe015ad >= 0; __obf_b827ad869fe015ad-- {
		__obf_70b1c35e7203372c += (uint(__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_b827ad869fe015ad]) - '0') << __obf_0fefbf9a88cfc16b
		__obf_25fbf8e7c24f4f7d := __obf_70b1c35e7203372c / 10
		__obf_09560398ee91545f := __obf_70b1c35e7203372c - 10*__obf_25fbf8e7c24f4f7d
		__obf_f85c5aae9b59cb4e--
		if __obf_f85c5aae9b59cb4e < len(__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4) {
			__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_f85c5aae9b59cb4e] = byte(__obf_09560398ee91545f + '0')
		} else if __obf_09560398ee91545f != 0 {
			__obf_55d13866fd49a7b4.__obf_3a6f83bf0d7dafe4 = true
		}
		__obf_70b1c35e7203372c = __obf_25fbf8e7c24f4f7d
	}

	// Put down extra digits.
	for __obf_70b1c35e7203372c > 0 {
		__obf_25fbf8e7c24f4f7d := __obf_70b1c35e7203372c / 10
		__obf_09560398ee91545f := __obf_70b1c35e7203372c - 10*__obf_25fbf8e7c24f4f7d
		__obf_f85c5aae9b59cb4e--
		if __obf_f85c5aae9b59cb4e < len(__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4) {
			__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_f85c5aae9b59cb4e] = byte(__obf_09560398ee91545f + '0')
		} else if __obf_09560398ee91545f != 0 {
			__obf_55d13866fd49a7b4.__obf_3a6f83bf0d7dafe4 = true
		}
		__obf_70b1c35e7203372c = __obf_25fbf8e7c24f4f7d
	}
	__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 += __obf_c47fe4ed44a7322b
	if __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 >= len(__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4) {
		__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 = len(__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4)
	}
	__obf_55d13866fd49a7b4.__obf_164d07da7c832f2c += __obf_c47fe4ed44a7322b

	// Binary shift left (k > 0) or right (k < 0).
	__obf_800f8c75bc7c90e0(__obf_55d13866fd49a7b4)
}

func (__obf_55d13866fd49a7b4 *__obf_3e0a215d3ac5298d) Shift(__obf_0fefbf9a88cfc16b int) {
	switch {
	case __obf_55d13866fd49a7b4.
		// nothing to do: a == 0
		__obf_7a088d2e01e4cd05 == 0:

	case __obf_0fefbf9a88cfc16b > 0:
		for __obf_0fefbf9a88cfc16b > __obf_b51cc6d999d130c7 {
			__obf_291f816b97326a6a(__obf_55d13866fd49a7b4, __obf_b51cc6d999d130c7)
			__obf_0fefbf9a88cfc16b -= __obf_b51cc6d999d130c7
		}
		__obf_291f816b97326a6a(__obf_55d13866fd49a7b4, uint(__obf_0fefbf9a88cfc16b))
	case __obf_0fefbf9a88cfc16b < 0:
		for __obf_0fefbf9a88cfc16b < -__obf_b51cc6d999d130c7 {
			__obf_9ef64b0acb08f65b(__obf_55d13866fd49a7b4, __obf_b51cc6d999d130c7)
			__obf_0fefbf9a88cfc16b += __obf_b51cc6d999d130c7
		}
		__obf_9ef64b0acb08f65b(__obf_55d13866fd49a7b4, uint(-__obf_0fefbf9a88cfc16b))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_38e87e7479153a4c(__obf_55d13866fd49a7b4 *__obf_3e0a215d3ac5298d, __obf_7a088d2e01e4cd05 int) bool {
	if __obf_7a088d2e01e4cd05 < 0 || __obf_7a088d2e01e4cd05 >= __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 {
		return false
	}
	if __obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_7a088d2e01e4cd05] == '5' && __obf_7a088d2e01e4cd05+1 == __obf_55d13866fd49a7b4. // exactly halfway - round to even
																		__obf_7a088d2e01e4cd05 {
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_55d13866fd49a7b4.__obf_3a6f83bf0d7dafe4 {
			return true
		}
		return __obf_7a088d2e01e4cd05 > 0 && (__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_7a088d2e01e4cd05-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_55d13866fd49a7b4.

		// Round a to nd digits (or fewer).
		// If nd is zero, it means we're rounding
		// just to the left of the digits, as in
		// 0.09 -> 0.1.
		__obf_e498fd1d3feac2c4[__obf_7a088d2e01e4cd05] >= '5'
}

func (__obf_55d13866fd49a7b4 *__obf_3e0a215d3ac5298d) Round(__obf_7a088d2e01e4cd05 int) {
	if __obf_7a088d2e01e4cd05 < 0 || __obf_7a088d2e01e4cd05 >= __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 {
		return
	}
	if __obf_38e87e7479153a4c(__obf_55d13866fd49a7b4, __obf_7a088d2e01e4cd05) {
		__obf_55d13866fd49a7b4.
			RoundUp(__obf_7a088d2e01e4cd05)
	} else {
		__obf_55d13866fd49a7b4.
			RoundDown(__obf_7a088d2e01e4cd05)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_55d13866fd49a7b4 *__obf_3e0a215d3ac5298d) RoundDown(__obf_7a088d2e01e4cd05 int) {
	if __obf_7a088d2e01e4cd05 < 0 || __obf_7a088d2e01e4cd05 >= __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 {
		return
	}
	__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 = __obf_7a088d2e01e4cd05

	// Round a up to nd digits (or fewer).
	__obf_800f8c75bc7c90e0(__obf_55d13866fd49a7b4)
}

func (__obf_55d13866fd49a7b4 *__obf_3e0a215d3ac5298d) RoundUp(__obf_7a088d2e01e4cd05 int) {
	if __obf_7a088d2e01e4cd05 < 0 || __obf_7a088d2e01e4cd05 >= __obf_55d13866fd49a7b4.

		// round up
		__obf_7a088d2e01e4cd05 {
		return
	}

	for __obf_a32504ffc5fafcfe := __obf_7a088d2e01e4cd05 - 1; __obf_a32504ffc5fafcfe >= 0; __obf_a32504ffc5fafcfe-- {
		__obf_d0f3a30e016877c1 := __obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_a32504ffc5fafcfe]
		if __obf_d0f3a30e016877c1 < '9' {
			__obf_55d13866fd49a7b4. // can stop after this digit
						__obf_e498fd1d3feac2c4[__obf_a32504ffc5fafcfe]++
			__obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05 = __obf_a32504ffc5fafcfe + 1
			return
		}
	}
	__obf_55d13866fd49a7b4.

		// Number is all 9s.
		// Change to single 1 with adjusted decimal point.
		__obf_e498fd1d3feac2c4[0] = '1'
	__obf_55d13866fd49a7b4.

		// Extract integer part, rounded appropriately.
		// No guarantees about overflow.
		__obf_7a088d2e01e4cd05 = 1
	__obf_55d13866fd49a7b4.__obf_164d07da7c832f2c++
}

func (__obf_55d13866fd49a7b4 *__obf_3e0a215d3ac5298d) RoundedInteger() uint64 {
	if __obf_55d13866fd49a7b4.__obf_164d07da7c832f2c > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_a32504ffc5fafcfe int
	__obf_70b1c35e7203372c := uint64(0)
	for __obf_a32504ffc5fafcfe = 0; __obf_a32504ffc5fafcfe < __obf_55d13866fd49a7b4.__obf_164d07da7c832f2c && __obf_a32504ffc5fafcfe < __obf_55d13866fd49a7b4.__obf_7a088d2e01e4cd05; __obf_a32504ffc5fafcfe++ {
		__obf_70b1c35e7203372c = __obf_70b1c35e7203372c*10 + uint64(__obf_55d13866fd49a7b4.__obf_e498fd1d3feac2c4[__obf_a32504ffc5fafcfe]-'0')
	}
	for ; __obf_a32504ffc5fafcfe < __obf_55d13866fd49a7b4.__obf_164d07da7c832f2c; __obf_a32504ffc5fafcfe++ {
		__obf_70b1c35e7203372c *= 10
	}
	if __obf_38e87e7479153a4c(__obf_55d13866fd49a7b4, __obf_55d13866fd49a7b4.__obf_164d07da7c832f2c) {
		__obf_70b1c35e7203372c++
	}
	return __obf_70b1c35e7203372c
}
