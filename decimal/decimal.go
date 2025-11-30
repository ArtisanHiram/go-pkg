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
package __obf_ebbf7bb33b20bf6e

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

var __obf_6f124f5735d8ea59 = big.NewInt(0)
var __obf_eb38f1915e7fe761 = big.NewInt(1)
var __obf_b39aaa52669b613d = big.NewInt(2)
var __obf_ac06b4a01c95937c = big.NewInt(4)
var __obf_e52ae739ff41f438 = big.NewInt(5)
var __obf_068b23c822520b43 = big.NewInt(10)
var __obf_f19af932fedccc2d = big.NewInt(20)

var __obf_1f497da7032f6620 = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_356f14ec9671961e *big.Int
	__obf_4163ad30ac75204b int32 // NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.

}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_356f14ec9671961e int64, __obf_4163ad30ac75204b int32) Decimal {
	return Decimal{__obf_356f14ec9671961e: big.NewInt(__obf_356f14ec9671961e), __obf_4163ad30ac75204b: __obf_4163ad30ac75204b}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_356f14ec9671961e int64) Decimal {
	return Decimal{__obf_356f14ec9671961e: big.NewInt(__obf_356f14ec9671961e), __obf_4163ad30ac75204b: 0}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_356f14ec9671961e int32) Decimal {
	return Decimal{__obf_356f14ec9671961e: big.NewInt(int64(__obf_356f14ec9671961e)), __obf_4163ad30ac75204b: 0}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_356f14ec9671961e uint64) Decimal {
	return Decimal{__obf_356f14ec9671961e: new(big.Int).SetUint64(__obf_356f14ec9671961e), __obf_4163ad30ac75204b: 0}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_356f14ec9671961e *big.Int, __obf_4163ad30ac75204b int32) Decimal {
	return Decimal{__obf_356f14ec9671961e: new(big.Int).Set(__obf_356f14ec9671961e), __obf_4163ad30ac75204b: __obf_4163ad30ac75204b}
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
func NewFromBigRat(__obf_356f14ec9671961e *big.Rat, __obf_c0882397220e65d2 int32) Decimal {
	return Decimal{__obf_356f14ec9671961e: new(big.Int).Set(__obf_356f14ec9671961e.Num()), __obf_4163ad30ac75204b: 0}.DivRound(Decimal{__obf_356f14ec9671961e: new(big.Int).Set(__obf_356f14ec9671961e.Denom()), __obf_4163ad30ac75204b: 0}, __obf_c0882397220e65d2)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_356f14ec9671961e string) (Decimal, error) {
	__obf_c13857c6f91c4932 := __obf_356f14ec9671961e
	var __obf_d135efe13b8bbe30 string
	var __obf_4163ad30ac75204b int64
	__obf_6dd46a4fd2c1ae77 := // Check if number is using scientific notation
		strings.IndexAny(__obf_356f14ec9671961e, "Ee")
	if __obf_6dd46a4fd2c1ae77 != -1 {
		__obf_c7906bb4463f2920, __obf_25b3ae434b8d1db5 := strconv.ParseInt(__obf_356f14ec9671961e[__obf_6dd46a4fd2c1ae77+1:], 10, 32)
		if __obf_25b3ae434b8d1db5 != nil {
			if __obf_6d6c98fd6986850d, __obf_87dc3845739761cd := __obf_25b3ae434b8d1db5.(*strconv.NumError); __obf_87dc3845739761cd && __obf_6d6c98fd6986850d.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_356f14ec9671961e)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_356f14ec9671961e)
		}
		__obf_356f14ec9671961e = __obf_356f14ec9671961e[:__obf_6dd46a4fd2c1ae77]
		__obf_4163ad30ac75204b = __obf_c7906bb4463f2920
	}
	__obf_b58ee5d8cae558cc := -1
	__obf_b27138cffc88e78e := len(__obf_356f14ec9671961e)
	for __obf_6ee7f4c10f6c9355 := 0; __obf_6ee7f4c10f6c9355 < __obf_b27138cffc88e78e; __obf_6ee7f4c10f6c9355++ {
		if __obf_356f14ec9671961e[__obf_6ee7f4c10f6c9355] == '.' {
			if __obf_b58ee5d8cae558cc > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_356f14ec9671961e)
			}
			__obf_b58ee5d8cae558cc = __obf_6ee7f4c10f6c9355
		}
	}

	if __obf_b58ee5d8cae558cc == -1 {
		__obf_d135efe13b8bbe30 = // There is no decimal point, we can just parse the original string as
			// an int
			__obf_356f14ec9671961e
	} else {
		if __obf_b58ee5d8cae558cc+1 < __obf_b27138cffc88e78e {
			__obf_d135efe13b8bbe30 = __obf_356f14ec9671961e[:__obf_b58ee5d8cae558cc] + __obf_356f14ec9671961e[__obf_b58ee5d8cae558cc+1:]
		} else {
			__obf_d135efe13b8bbe30 = __obf_356f14ec9671961e[:__obf_b58ee5d8cae558cc]
		}
		__obf_c7906bb4463f2920 := -len(__obf_356f14ec9671961e[__obf_b58ee5d8cae558cc+1:])
		__obf_4163ad30ac75204b += int64(__obf_c7906bb4463f2920)
	}

	var __obf_1ae974a4483f33be *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_d135efe13b8bbe30) <= 18 {
		__obf_f99f42e04aeb5a69, __obf_25b3ae434b8d1db5 := strconv.ParseInt(__obf_d135efe13b8bbe30, 10, 64)
		if __obf_25b3ae434b8d1db5 != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_356f14ec9671961e)
		}
		__obf_1ae974a4483f33be = big.NewInt(__obf_f99f42e04aeb5a69)
	} else {
		__obf_1ae974a4483f33be = new(big.Int)
		_, __obf_87dc3845739761cd := __obf_1ae974a4483f33be.SetString(__obf_d135efe13b8bbe30, 10)
		if !__obf_87dc3845739761cd {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_356f14ec9671961e)
		}
	}

	if __obf_4163ad30ac75204b < math.MinInt32 || __obf_4163ad30ac75204b > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_c13857c6f91c4932)
	}

	return Decimal{__obf_356f14ec9671961e: __obf_1ae974a4483f33be, __obf_4163ad30ac75204b: int32(__obf_4163ad30ac75204b)}, nil
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
func NewFromFormattedString(__obf_356f14ec9671961e string, __obf_5ef4b135ff0d10a4 *regexp.Regexp) (Decimal, error) {
	__obf_aaec34fcf7f5ee6d := __obf_5ef4b135ff0d10a4.ReplaceAllString(__obf_356f14ec9671961e, "")
	__obf_3c5564ffc2911a1b, __obf_25b3ae434b8d1db5 := NewFromString(__obf_aaec34fcf7f5ee6d)
	if __obf_25b3ae434b8d1db5 != nil {
		return Decimal{}, __obf_25b3ae434b8d1db5
	}
	return __obf_3c5564ffc2911a1b, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_356f14ec9671961e string) Decimal {
	__obf_9397222af603e4cf, __obf_25b3ae434b8d1db5 := NewFromString(__obf_356f14ec9671961e)
	if __obf_25b3ae434b8d1db5 != nil {
		panic(__obf_25b3ae434b8d1db5)
	}
	return __obf_9397222af603e4cf
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
func NewFromFloat(__obf_356f14ec9671961e float64) Decimal {
	if __obf_356f14ec9671961e == 0 {
		return New(0, 0)
	}
	return __obf_97f48c64896e1fab(__obf_356f14ec9671961e, math.Float64bits(__obf_356f14ec9671961e), &__obf_3050fd41f9ad69c5)
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
func NewFromFloat32(__obf_356f14ec9671961e float32) Decimal {
	if __obf_356f14ec9671961e == 0 {
		return New(0, 0)
	}
	__obf_b93edf907a7818bf := // XOR is workaround for https://github.com/golang/go/issues/26285
		math.Float32bits(__obf_356f14ec9671961e) ^ 0x80808080
	return __obf_97f48c64896e1fab(float64(__obf_356f14ec9671961e), uint64(__obf_b93edf907a7818bf)^0x80808080, &__obf_da91496795acf581)
}

func __obf_97f48c64896e1fab(__obf_f2a33aba05a016a2 float64, __obf_87103b480525605f uint64, __obf_5c718fb2fc061777 *__obf_8d7a820fe70b0941) Decimal {
	if math.IsNaN(__obf_f2a33aba05a016a2) || math.IsInf(__obf_f2a33aba05a016a2, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_f2a33aba05a016a2))
	}
	__obf_4163ad30ac75204b := int(__obf_87103b480525605f>>__obf_5c718fb2fc061777.__obf_2824f9b56a37a620) & (1<<__obf_5c718fb2fc061777.__obf_a2f627ebab0e93ee - 1)
	__obf_a0aed3f02406911a := __obf_87103b480525605f & (uint64(1)<<__obf_5c718fb2fc061777.__obf_2824f9b56a37a620 - 1)

	switch __obf_4163ad30ac75204b {
	case 0:
		__obf_4163ad30ac75204b++ // denormalized

	default:
		__obf_a0aed3f02406911a |= // add implicit top bit
			uint64(1) << __obf_5c718fb2fc061777.__obf_2824f9b56a37a620
	}
	__obf_4163ad30ac75204b += __obf_5c718fb2fc061777.__obf_69950ce091b60b98

	var __obf_3c5564ffc2911a1b __obf_ebbf7bb33b20bf6e
	__obf_3c5564ffc2911a1b.
		Assign(__obf_a0aed3f02406911a)
	__obf_3c5564ffc2911a1b.
		Shift(__obf_4163ad30ac75204b - int(__obf_5c718fb2fc061777.__obf_2824f9b56a37a620))
	__obf_3c5564ffc2911a1b.__obf_d46708a8a96b3d7e = __obf_87103b480525605f>>(__obf_5c718fb2fc061777.__obf_a2f627ebab0e93ee+__obf_5c718fb2fc061777.__obf_2824f9b56a37a620) != 0
	__obf_ee1a55061233bcaf(&__obf_3c5564ffc2911a1b,
		// If less than 19 digits, we can do calculation in an int64.
		__obf_a0aed3f02406911a, __obf_4163ad30ac75204b, __obf_5c718fb2fc061777)

	if __obf_3c5564ffc2911a1b.__obf_f75b34971d3492db < 19 {
		__obf_c65f80ebe797c070 := int64(0)
		__obf_2c8d42b08bf1a2b7 := int64(1)
		for __obf_6ee7f4c10f6c9355 := __obf_3c5564ffc2911a1b.__obf_f75b34971d3492db - 1; __obf_6ee7f4c10f6c9355 >= 0; __obf_6ee7f4c10f6c9355-- {
			__obf_c65f80ebe797c070 += __obf_2c8d42b08bf1a2b7 * int64(__obf_3c5564ffc2911a1b.__obf_3c5564ffc2911a1b[__obf_6ee7f4c10f6c9355]-'0')
			__obf_2c8d42b08bf1a2b7 *= 10
		}
		if __obf_3c5564ffc2911a1b.__obf_d46708a8a96b3d7e {
			__obf_c65f80ebe797c070 *= -1
		}
		return Decimal{__obf_356f14ec9671961e: big.NewInt(__obf_c65f80ebe797c070), __obf_4163ad30ac75204b: int32(__obf_3c5564ffc2911a1b.__obf_43d811ed87289907) - int32(__obf_3c5564ffc2911a1b.__obf_f75b34971d3492db)}
	}
	__obf_1ae974a4483f33be := new(big.Int)
	__obf_1ae974a4483f33be, __obf_87dc3845739761cd := __obf_1ae974a4483f33be.SetString(string(__obf_3c5564ffc2911a1b.__obf_3c5564ffc2911a1b[:__obf_3c5564ffc2911a1b.__obf_f75b34971d3492db]), 10)
	if __obf_87dc3845739761cd {
		return Decimal{__obf_356f14ec9671961e: __obf_1ae974a4483f33be, __obf_4163ad30ac75204b: int32(__obf_3c5564ffc2911a1b.__obf_43d811ed87289907) - int32(__obf_3c5564ffc2911a1b.__obf_f75b34971d3492db)}
	}

	return NewFromFloatWithExponent(__obf_f2a33aba05a016a2, int32(__obf_3c5564ffc2911a1b.

		// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
		// number of fractional digits.
		//
		// Example:
		//
		//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
		__obf_43d811ed87289907)-int32(__obf_3c5564ffc2911a1b.__obf_f75b34971d3492db))
}

func NewFromFloatWithExponent(__obf_356f14ec9671961e float64, __obf_4163ad30ac75204b int32) Decimal {
	if math.IsNaN(__obf_356f14ec9671961e) || math.IsInf(__obf_356f14ec9671961e, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_356f14ec9671961e))
	}
	__obf_87103b480525605f := math.Float64bits(__obf_356f14ec9671961e)
	__obf_a0aed3f02406911a := __obf_87103b480525605f & (1<<52 - 1)
	__obf_a5c989bc65639fc4 := int32((__obf_87103b480525605f >> 52) & (1<<11 - 1))
	__obf_07c356752a4027b3 := __obf_87103b480525605f >> 63

	if __obf_a5c989bc65639fc4 == 0 {
		// specials
		if __obf_a0aed3f02406911a == 0 {
			return Decimal{}
		}
		__obf_a5c989bc65639fc4++ // subnormal

	} else {
		__obf_a0aed3f02406911a |= // normal
			1 << 52
	}
	__obf_a5c989bc65639fc4 -= 1023 + 52

	// normalizing base-2 values
	for __obf_a0aed3f02406911a&1 == 0 {
		__obf_a0aed3f02406911a = __obf_a0aed3f02406911a >> 1
		__obf_a5c989bc65639fc4++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_4163ad30ac75204b < 0 && __obf_4163ad30ac75204b < __obf_a5c989bc65639fc4 {
		if __obf_a5c989bc65639fc4 < 0 {
			__obf_4163ad30ac75204b = __obf_a5c989bc65639fc4
		} else {
			__obf_4163ad30ac75204b = 0
		}
	}
	__obf_a5c989bc65639fc4 -= // representing 10^M * 2^N as 5^M * 2^(M+N)
		__obf_4163ad30ac75204b
	__obf_97a055e54a81ce50 := big.NewInt(1)
	__obf_d049224562f52d09 := big.NewInt(int64(__obf_a0aed3f02406911a))

	// applying 5^M
	if __obf_4163ad30ac75204b > 0 {
		__obf_97a055e54a81ce50 = __obf_97a055e54a81ce50.SetInt64(int64(__obf_4163ad30ac75204b))
		__obf_97a055e54a81ce50 = __obf_97a055e54a81ce50.Exp(__obf_e52ae739ff41f438, __obf_97a055e54a81ce50, nil)
	} else if __obf_4163ad30ac75204b < 0 {
		__obf_97a055e54a81ce50 = __obf_97a055e54a81ce50.SetInt64(-int64(__obf_4163ad30ac75204b))
		__obf_97a055e54a81ce50 = __obf_97a055e54a81ce50.Exp(__obf_e52ae739ff41f438, __obf_97a055e54a81ce50, nil)
		__obf_d049224562f52d09 = __obf_d049224562f52d09.Mul(__obf_d049224562f52d09, __obf_97a055e54a81ce50)
		__obf_97a055e54a81ce50 = __obf_97a055e54a81ce50.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_a5c989bc65639fc4 > 0 {
		__obf_d049224562f52d09 = __obf_d049224562f52d09.Lsh(__obf_d049224562f52d09, uint(__obf_a5c989bc65639fc4))
	} else if __obf_a5c989bc65639fc4 < 0 {
		__obf_97a055e54a81ce50 = __obf_97a055e54a81ce50.Lsh(__obf_97a055e54a81ce50, uint(-__obf_a5c989bc65639fc4))
	}

	// rounding and downscaling
	if __obf_4163ad30ac75204b > 0 || __obf_a5c989bc65639fc4 < 0 {
		__obf_a15db1dc23afcb5f := new(big.Int).Rsh(__obf_97a055e54a81ce50, 1)
		__obf_d049224562f52d09 = __obf_d049224562f52d09.Add(__obf_d049224562f52d09, __obf_a15db1dc23afcb5f)
		__obf_d049224562f52d09 = __obf_d049224562f52d09.Quo(__obf_d049224562f52d09, __obf_97a055e54a81ce50)
	}

	if __obf_07c356752a4027b3 == 1 {
		__obf_d049224562f52d09 = __obf_d049224562f52d09.Neg(__obf_d049224562f52d09)
	}

	return Decimal{__obf_356f14ec9671961e: __obf_d049224562f52d09, __obf_4163ad30ac75204b: __obf_4163ad30ac75204b}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_3c5564ffc2911a1b Decimal) Copy() Decimal {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()
	return Decimal{__obf_356f14ec9671961e: new(big.Int).Set(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e), __obf_4163ad30ac75204b: __obf_3c5564ffc2911a1b.

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
		__obf_4163ad30ac75204b,
	}
}

func (__obf_3c5564ffc2911a1b Decimal) __obf_cfec7d8af9246d7b(__obf_4163ad30ac75204b int32) Decimal {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()

	if __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b == __obf_4163ad30ac75204b {
		return Decimal{
			new(big.Int).Set(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e), __obf_3c5564ffc2911a1b.

				// NOTE(vadim): must convert exps to float64 before - to prevent overflow
				__obf_4163ad30ac75204b,
		}
	}
	__obf_80308aacd8589f04 := math.Abs(float64(__obf_4163ad30ac75204b) - float64(__obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b))
	__obf_356f14ec9671961e := new(big.Int).Set(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e)
	__obf_06518a185ef40ccf := new(big.Int).Exp(__obf_068b23c822520b43, big.NewInt(int64(__obf_80308aacd8589f04)), nil)
	if __obf_4163ad30ac75204b > __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b {
		__obf_356f14ec9671961e = __obf_356f14ec9671961e.Quo(__obf_356f14ec9671961e, __obf_06518a185ef40ccf)
	} else if __obf_4163ad30ac75204b < __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b {
		__obf_356f14ec9671961e = __obf_356f14ec9671961e.Mul(__obf_356f14ec9671961e, __obf_06518a185ef40ccf)
	}

	return Decimal{__obf_356f14ec9671961e: __obf_356f14ec9671961e, __obf_4163ad30ac75204b: __obf_4163ad30ac75204b}
}

// Abs returns the absolute value of the decimal.
func (__obf_3c5564ffc2911a1b Decimal) Abs() Decimal {
	if !__obf_3c5564ffc2911a1b.IsNegative() {
		return __obf_3c5564ffc2911a1b
	}
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()
	__obf_ad3cf29659514d86 := new(big.Int).Abs(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e)
	return Decimal{__obf_356f14ec9671961e: __obf_ad3cf29659514d86, __obf_4163ad30ac75204b: __obf_3c5564ffc2911a1b.

		// Add returns d + d2.
		__obf_4163ad30ac75204b,
	}
}

func (__obf_3c5564ffc2911a1b Decimal) Add(__obf_718cb7c55117fa47 Decimal) Decimal {
	__obf_9f7dbbc061b0788c, __obf_ef5ba9ca72c05c92 := RescalePair(__obf_3c5564ffc2911a1b, __obf_718cb7c55117fa47)
	__obf_8c7d24b8e02673f9 := new(big.Int).Add(__obf_9f7dbbc061b0788c.__obf_356f14ec9671961e, __obf_ef5ba9ca72c05c92.__obf_356f14ec9671961e)
	return Decimal{__obf_356f14ec9671961e: __obf_8c7d24b8e02673f9, __obf_4163ad30ac75204b: __obf_9f7dbbc061b0788c.

		// Sub returns d - d2.
		__obf_4163ad30ac75204b,
	}
}

func (__obf_3c5564ffc2911a1b Decimal) Sub(__obf_718cb7c55117fa47 Decimal) Decimal {
	__obf_9f7dbbc061b0788c, __obf_ef5ba9ca72c05c92 := RescalePair(__obf_3c5564ffc2911a1b, __obf_718cb7c55117fa47)
	__obf_8c7d24b8e02673f9 := new(big.Int).Sub(__obf_9f7dbbc061b0788c.__obf_356f14ec9671961e, __obf_ef5ba9ca72c05c92.__obf_356f14ec9671961e)
	return Decimal{__obf_356f14ec9671961e: __obf_8c7d24b8e02673f9, __obf_4163ad30ac75204b: __obf_9f7dbbc061b0788c.

		// Neg returns -d.
		__obf_4163ad30ac75204b,
	}
}

func (__obf_3c5564ffc2911a1b Decimal) Neg() Decimal {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()
	__obf_f2a33aba05a016a2 := new(big.Int).Neg(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e)
	return Decimal{__obf_356f14ec9671961e: __obf_f2a33aba05a016a2, __obf_4163ad30ac75204b: __obf_3c5564ffc2911a1b.

		// Mul returns d * d2.
		__obf_4163ad30ac75204b,
	}
}

func (__obf_3c5564ffc2911a1b Decimal) Mul(__obf_718cb7c55117fa47 Decimal) Decimal {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()
	__obf_718cb7c55117fa47.__obf_4d97a450ece4f814()
	__obf_9f2aa0bc0164bf0a := int64(__obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b) + int64(__obf_718cb7c55117fa47.__obf_4163ad30ac75204b)
	if __obf_9f2aa0bc0164bf0a > math.MaxInt32 || __obf_9f2aa0bc0164bf0a < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_9f2aa0bc0164bf0a))
	}
	__obf_8c7d24b8e02673f9 := new(big.Int).Mul(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e, __obf_718cb7c55117fa47.__obf_356f14ec9671961e)
	return Decimal{__obf_356f14ec9671961e: __obf_8c7d24b8e02673f9, __obf_4163ad30ac75204b: int32(__obf_9f2aa0bc0164bf0a)}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_3c5564ffc2911a1b Decimal) Shift(__obf_3d20977dcaeff532 int32) Decimal {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()
	return Decimal{__obf_356f14ec9671961e: new(big.Int).Set(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e), __obf_4163ad30ac75204b: __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b + __obf_3d20977dcaeff532}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_3c5564ffc2911a1b Decimal) Div(__obf_718cb7c55117fa47 Decimal) Decimal {
	return __obf_3c5564ffc2911a1b.DivRound(__obf_718cb7c55117fa47, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_3c5564ffc2911a1b Decimal) QuoRem(__obf_718cb7c55117fa47 Decimal, __obf_c0882397220e65d2 int32) (Decimal, Decimal) {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()
	__obf_718cb7c55117fa47.__obf_4d97a450ece4f814()
	if __obf_718cb7c55117fa47.__obf_356f14ec9671961e.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_53aa31d93fab941a := -__obf_c0882397220e65d2
	__obf_6d6c98fd6986850d := int64(__obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b) - int64(__obf_718cb7c55117fa47.__obf_4163ad30ac75204b) - int64(__obf_53aa31d93fab941a)
	if __obf_6d6c98fd6986850d > math.MaxInt32 || __obf_6d6c98fd6986850d < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_306cfadc1d14a972, __obf_b8467b0738beae97,

		// d = a 10^ea
		// d2 = b 10^eb
		__obf_c35dc4b9a0af0ae5 big.Int
	var __obf_020cb5b1606e6d54 int32

	if __obf_6d6c98fd6986850d < 0 {
		__obf_306cfadc1d14a972 = *__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e
		__obf_c35dc4b9a0af0ae5.
			SetInt64(-__obf_6d6c98fd6986850d)
		__obf_b8467b0738beae97.
			Exp(__obf_068b23c822520b43, &__obf_c35dc4b9a0af0ae5, nil)
		__obf_b8467b0738beae97.
			Mul(__obf_718cb7c55117fa47.__obf_356f14ec9671961e, &__obf_b8467b0738beae97)
		__obf_020cb5b1606e6d54 = __obf_3c5564ffc2911a1b.
			// now aa = a
			//     bb = b 10^(scale + eb - ea)
			__obf_4163ad30ac75204b

	} else {
		__obf_c35dc4b9a0af0ae5.
			SetInt64(__obf_6d6c98fd6986850d)
		__obf_306cfadc1d14a972.
			Exp(__obf_068b23c822520b43, &__obf_c35dc4b9a0af0ae5, nil)
		__obf_306cfadc1d14a972.
			Mul(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e, &__obf_306cfadc1d14a972)
		__obf_b8467b0738beae97 = *__obf_718cb7c55117fa47.__obf_356f14ec9671961e

		// now aa = a ^ (ea - eb - scale)
		//     bb = b
		__obf_020cb5b1606e6d54 = __obf_53aa31d93fab941a + __obf_718cb7c55117fa47.__obf_4163ad30ac75204b

	}
	var __obf_58dd48c2daf72a44, __obf_2b7b4dcf959f4b21 big.Int
	__obf_58dd48c2daf72a44.
		QuoRem(&__obf_306cfadc1d14a972, &__obf_b8467b0738beae97, &__obf_2b7b4dcf959f4b21)
	__obf_5c8ab66b25579fce := Decimal{__obf_356f14ec9671961e: &__obf_58dd48c2daf72a44, __obf_4163ad30ac75204b: __obf_53aa31d93fab941a}
	__obf_2ab6d38ba3fb0990 := Decimal{__obf_356f14ec9671961e: &__obf_2b7b4dcf959f4b21, __obf_4163ad30ac75204b: __obf_020cb5b1606e6d54}
	return __obf_5c8ab66b25579fce,

		// DivRound divides and rounds to a given precision
		// i.e. to an integer multiple of 10^(-precision)
		//
		//	for a positive quotient digit 5 is rounded up, away from 0
		//	if the quotient is negative then digit 5 is rounded down, away from 0
		//
		// Note that precision<0 is allowed as input.
		__obf_2ab6d38ba3fb0990
}

func (__obf_3c5564ffc2911a1b Decimal) DivRound(__obf_718cb7c55117fa47 Decimal, __obf_c0882397220e65d2 int32) Decimal {
	__obf_58dd48c2daf72a44,
		// QuoRem already checks initialization
		__obf_2b7b4dcf959f4b21 := __obf_3c5564ffc2911a1b.QuoRem(__obf_718cb7c55117fa47,
		// the actual rounding decision is based on comparing r*10^precision and d2/2
		// instead compare 2 r 10 ^precision and d2
		__obf_c0882397220e65d2)

	var __obf_615fbf3eb28a1cab big.Int
	__obf_615fbf3eb28a1cab.
		Abs(__obf_2b7b4dcf959f4b21.__obf_356f14ec9671961e)
	__obf_615fbf3eb28a1cab.
		Lsh(&__obf_615fbf3eb28a1cab, 1)
	__obf_3f5f173e15237696 := // now rv2 = abs(r.value) * 2
		Decimal{__obf_356f14ec9671961e: &__obf_615fbf3eb28a1cab, __obf_4163ad30ac75204b: __obf_2b7b4dcf959f4b21.
			// r2 is now 2 * r * 10 ^ precision
			__obf_4163ad30ac75204b + __obf_c0882397220e65d2}

	var __obf_a1cf2ec9c989d244 = __obf_3f5f173e15237696.Cmp(__obf_718cb7c55117fa47.Abs())

	if __obf_a1cf2ec9c989d244 < 0 {
		return __obf_58dd48c2daf72a44
	}

	if __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.Sign()*__obf_718cb7c55117fa47.__obf_356f14ec9671961e.Sign() < 0 {
		return __obf_58dd48c2daf72a44.Sub(New(1, -__obf_c0882397220e65d2))
	}

	return __obf_58dd48c2daf72a44.Add(New(1, -__obf_c0882397220e65d2))
}

// Mod returns d % d2.
func (__obf_3c5564ffc2911a1b Decimal) Mod(__obf_718cb7c55117fa47 Decimal) Decimal {
	_, __obf_2b7b4dcf959f4b21 := __obf_3c5564ffc2911a1b.QuoRem(__obf_718cb7c55117fa47, 0)
	return __obf_2b7b4dcf959f4b21
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
func (__obf_3c5564ffc2911a1b Decimal) Pow(__obf_718cb7c55117fa47 Decimal) Decimal {
	__obf_09dbe2a3aac0b561 := __obf_3c5564ffc2911a1b.Sign()
	__obf_39bbb905d2cda1d2 := __obf_718cb7c55117fa47.Sign()

	if __obf_09dbe2a3aac0b561 == 0 {
		if __obf_39bbb905d2cda1d2 == 0 {
			return Decimal{}
		}
		if __obf_39bbb905d2cda1d2 == 1 {
			return Decimal{__obf_6f124f5735d8ea59, 0}
		}
		if __obf_39bbb905d2cda1d2 == -1 {
			return Decimal{}
		}
	}

	if __obf_39bbb905d2cda1d2 == 0 {
		return Decimal{__obf_eb38f1915e7fe761, 0}
	}
	__obf_98940868b0fe9101 := // TODO: optimize extraction of fractional part
		Decimal{__obf_eb38f1915e7fe761, 0}
	__obf_161acd91aab9dce4, __obf_ea118a9a47665a3e := __obf_718cb7c55117fa47.QuoRem(__obf_98940868b0fe9101, 0)

	if __obf_09dbe2a3aac0b561 == -1 && !__obf_ea118a9a47665a3e.IsZero() {
		return Decimal{}
	}
	__obf_7c46fd50ca27ee8a, _ := __obf_3c5564ffc2911a1b.PowBigInt(__obf_161acd91aab9dce4.

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_356f14ec9671961e)

	if __obf_ea118a9a47665a3e.__obf_356f14ec9671961e.Sign() == 0 {
		return __obf_7c46fd50ca27ee8a
	}
	__obf_c4867538582305d9 := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_3c5564ffc2911a1b.NumDigits()
	__obf_4387f4cf8730030c := __obf_718cb7c55117fa47.NumDigits()
	__obf_c0882397220e65d2 := __obf_c4867538582305d9

	if __obf_4387f4cf8730030c > __obf_c0882397220e65d2 {
		__obf_c0882397220e65d2 += __obf_4387f4cf8730030c
	}
	__obf_c0882397220e65d2 += 6
	__obf_0b0cc1f26a40d814,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_25b3ae434b8d1db5 := __obf_3c5564ffc2911a1b.Abs().Ln(-__obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b + int32(__obf_c0882397220e65d2))
	if __obf_25b3ae434b8d1db5 != nil {
		return Decimal{}
	}
	__obf_0b0cc1f26a40d814 = __obf_0b0cc1f26a40d814.Mul(__obf_ea118a9a47665a3e)
	__obf_0b0cc1f26a40d814, __obf_25b3ae434b8d1db5 = __obf_0b0cc1f26a40d814.ExpTaylor(-__obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b + int32(__obf_c0882397220e65d2))
	if __obf_25b3ae434b8d1db5 != nil {
		return Decimal{}
	}
	__obf_bbd65286a9fac425 := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_7c46fd50ca27ee8a.Mul(__obf_0b0cc1f26a40d814)

	return __obf_bbd65286a9fac425
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
func (__obf_3c5564ffc2911a1b Decimal) PowWithPrecision(__obf_718cb7c55117fa47 Decimal, __obf_c0882397220e65d2 int32) (Decimal, error) {
	__obf_09dbe2a3aac0b561 := __obf_3c5564ffc2911a1b.Sign()
	__obf_39bbb905d2cda1d2 := __obf_718cb7c55117fa47.Sign()

	if __obf_09dbe2a3aac0b561 == 0 {
		if __obf_39bbb905d2cda1d2 == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_39bbb905d2cda1d2 == 1 {
			return Decimal{__obf_6f124f5735d8ea59, 0}, nil
		}
		if __obf_39bbb905d2cda1d2 == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_39bbb905d2cda1d2 == 0 {
		return Decimal{__obf_eb38f1915e7fe761, 0}, nil
	}
	__obf_98940868b0fe9101 := // TODO: optimize extraction of fractional part
		Decimal{__obf_eb38f1915e7fe761, 0}
	__obf_161acd91aab9dce4, __obf_ea118a9a47665a3e := __obf_718cb7c55117fa47.QuoRem(__obf_98940868b0fe9101, 0)

	if __obf_09dbe2a3aac0b561 == -1 && !__obf_ea118a9a47665a3e.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}
	__obf_7c46fd50ca27ee8a, _ := __obf_3c5564ffc2911a1b.__obf_1633b02533c4a79c(__obf_161acd91aab9dce4.__obf_356f14ec9671961e,

		// if exponent is an integer we don't need to calculate d1**frac(d2)
		__obf_c0882397220e65d2)

	if __obf_ea118a9a47665a3e.__obf_356f14ec9671961e.Sign() == 0 {
		return __obf_7c46fd50ca27ee8a, nil
	}
	__obf_c4867538582305d9 := // TODO: optimize NumDigits for more performant precision adjustment
		__obf_3c5564ffc2911a1b.NumDigits()
	__obf_4387f4cf8730030c := __obf_718cb7c55117fa47.NumDigits()

	if int32(__obf_c4867538582305d9) > __obf_c0882397220e65d2 {
		__obf_c0882397220e65d2 = int32(__obf_c4867538582305d9)
	}
	if int32(__obf_4387f4cf8730030c) > __obf_c0882397220e65d2 {
		__obf_c0882397220e65d2 += int32(__obf_4387f4cf8730030c)
	}
	__obf_c0882397220e65d2 += // increase precision by 10 to compensate for errors in further calculations
		10
	__obf_0b0cc1f26a40d814,

		// Calculate x ** frac(y), where
		// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
		__obf_25b3ae434b8d1db5 := __obf_3c5564ffc2911a1b.Abs().Ln(__obf_c0882397220e65d2)
	if __obf_25b3ae434b8d1db5 != nil {
		return Decimal{}, __obf_25b3ae434b8d1db5
	}
	__obf_0b0cc1f26a40d814 = __obf_0b0cc1f26a40d814.Mul(__obf_ea118a9a47665a3e)
	__obf_0b0cc1f26a40d814, __obf_25b3ae434b8d1db5 = __obf_0b0cc1f26a40d814.ExpTaylor(__obf_c0882397220e65d2)
	if __obf_25b3ae434b8d1db5 != nil {
		return Decimal{}, __obf_25b3ae434b8d1db5
	}
	__obf_bbd65286a9fac425 := // Join integer and fractional part,
		// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
		__obf_7c46fd50ca27ee8a.Mul(__obf_0b0cc1f26a40d814)

	return __obf_bbd65286a9fac425, nil
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
func (__obf_3c5564ffc2911a1b Decimal) PowInt32(__obf_4163ad30ac75204b int32) (Decimal, error) {
	if __obf_3c5564ffc2911a1b.IsZero() && __obf_4163ad30ac75204b == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_683ee08a7afe15da := __obf_4163ad30ac75204b < 0
	__obf_4163ad30ac75204b = __obf_d44c132c275234aa(__obf_4163ad30ac75204b)
	__obf_0657fff5d9f5aeef, __obf_15a03af38514d25c := __obf_3c5564ffc2911a1b, New(1, 0)

	for __obf_4163ad30ac75204b > 0 {
		if __obf_4163ad30ac75204b%2 == 1 {
			__obf_15a03af38514d25c = __obf_15a03af38514d25c.Mul(__obf_0657fff5d9f5aeef)
		}
		__obf_4163ad30ac75204b /= 2

		if __obf_4163ad30ac75204b > 0 {
			__obf_0657fff5d9f5aeef = __obf_0657fff5d9f5aeef.Mul(__obf_0657fff5d9f5aeef)
		}
	}

	if __obf_683ee08a7afe15da {
		return New(1, 0).DivRound(__obf_15a03af38514d25c, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_15a03af38514d25c, nil
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
func (__obf_3c5564ffc2911a1b Decimal) PowBigInt(__obf_4163ad30ac75204b *big.Int) (Decimal, error) {
	return __obf_3c5564ffc2911a1b.__obf_1633b02533c4a79c(__obf_4163ad30ac75204b, int32(PowPrecisionNegativeExponent))
}

func (__obf_3c5564ffc2911a1b Decimal) __obf_1633b02533c4a79c(__obf_4163ad30ac75204b *big.Int, __obf_c0882397220e65d2 int32) (Decimal, error) {
	if __obf_3c5564ffc2911a1b.IsZero() && __obf_4163ad30ac75204b.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}
	__obf_5e8b93059fc00047 := new(big.Int).Set(__obf_4163ad30ac75204b)
	__obf_683ee08a7afe15da := __obf_4163ad30ac75204b.Sign() < 0

	if __obf_683ee08a7afe15da {
		__obf_5e8b93059fc00047.
			Abs(__obf_5e8b93059fc00047)
	}
	__obf_0657fff5d9f5aeef, __obf_15a03af38514d25c := __obf_3c5564ffc2911a1b, New(1, 0)

	for __obf_5e8b93059fc00047.Sign() > 0 {
		if __obf_5e8b93059fc00047.Bit(0) == 1 {
			__obf_15a03af38514d25c = __obf_15a03af38514d25c.Mul(__obf_0657fff5d9f5aeef)
		}
		__obf_5e8b93059fc00047.
			Rsh(__obf_5e8b93059fc00047, 1)

		if __obf_5e8b93059fc00047.Sign() > 0 {
			__obf_0657fff5d9f5aeef = __obf_0657fff5d9f5aeef.Mul(__obf_0657fff5d9f5aeef)
		}
	}

	if __obf_683ee08a7afe15da {
		return New(1, 0).DivRound(__obf_15a03af38514d25c, __obf_c0882397220e65d2), nil
	}

	return __obf_15a03af38514d25c, nil
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
func (__obf_3c5564ffc2911a1b Decimal) ExpHullAbrham(__obf_1b2724d718d3a085 uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_3c5564ffc2911a1b.IsZero() {
		return Decimal{__obf_eb38f1915e7fe761, 0}, nil
	}
	__obf_d7825254824b100e := __obf_1b2724d718d3a085

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_cafcc62d0a40df85 := __obf_3c5564ffc2911a1b.Abs().InexactFloat64()
	if __obf_9c917c9d91b5da79 := __obf_cafcc62d0a40df85 / 23; __obf_9c917c9d91b5da79 > float64(__obf_d7825254824b100e) && __obf_9c917c9d91b5da79 < float64(ExpMaxIterations) {
		__obf_d7825254824b100e = uint32(math.Ceil(__obf_9c917c9d91b5da79))
	}
	__obf_f2c11776b3daf5e2 := // fail if abs(d) beyond an over/underflow threshold
		New(23*int64(__obf_d7825254824b100e), 0)
	if __obf_3c5564ffc2911a1b.Abs().Cmp(__obf_f2c11776b3daf5e2) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}
	__obf_ca165653b40274b8 := // Return 1 if abs(d) small enough; this also avoids later over/underflow
		New(9, -int32(__obf_d7825254824b100e)-1)
	if __obf_3c5564ffc2911a1b.Abs().Cmp(__obf_ca165653b40274b8) <= 0 {
		return Decimal{__obf_eb38f1915e7fe761, __obf_3c5564ffc2911a1b.

			// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
			__obf_4163ad30ac75204b}, nil
	}
	__obf_c26204f4870c50c6 := __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b + int32(__obf_3c5564ffc2911a1b.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_c26204f4870c50c6 < 0 {
		__obf_c26204f4870c50c6 = 0
	}
	__obf_55924e50211870ef := New(1, __obf_c26204f4870c50c6)
	__obf_2b7b4dcf959f4b21 := // reduction factor
		Decimal{new(big.Int).Set(__obf_3c5564ffc2911a1b. // reduced argument
									__obf_356f14ec9671961e), __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b - __obf_c26204f4870c50c6}
	__obf_0a8f7728610094e3 := int32(__obf_d7825254824b100e) + __obf_c26204f4870c50c6 + 2
	__obf_e4ae44929fd6c4e0 := // precision for calculating the sum

		// Determine n, the number of therms for calculating sum
		// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
		// for solving appropriate equation, along with directed
		// roundings and simple rational bound for log10(p/abs(r))
		__obf_2b7b4dcf959f4b21.Abs().InexactFloat64()
	__obf_73e31dfb35ee7500 := float64(__obf_0a8f7728610094e3)
	__obf_25dd64ab1c007767 := math.Ceil((1.453*__obf_73e31dfb35ee7500 - 1.182) / math.Log10(__obf_73e31dfb35ee7500/__obf_e4ae44929fd6c4e0))
	if __obf_25dd64ab1c007767 > float64(ExpMaxIterations) || math.IsNaN(__obf_25dd64ab1c007767) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_0657fff5d9f5aeef := int64(__obf_25dd64ab1c007767)
	__obf_c65f80ebe797c070 := New(0, 0)
	__obf_f9decec89e95425e := New(1, 0)
	__obf_98940868b0fe9101 := New(1, 0)
	for __obf_6ee7f4c10f6c9355 := __obf_0657fff5d9f5aeef - 1; __obf_6ee7f4c10f6c9355 > 0; __obf_6ee7f4c10f6c9355-- {
		__obf_c65f80ebe797c070.__obf_356f14ec9671961e.
			SetInt64(__obf_6ee7f4c10f6c9355)
		__obf_f9decec89e95425e = __obf_f9decec89e95425e.Mul(__obf_2b7b4dcf959f4b21.DivRound(__obf_c65f80ebe797c070, __obf_0a8f7728610094e3))
		__obf_f9decec89e95425e = __obf_f9decec89e95425e.Add(__obf_98940868b0fe9101)
	}
	__obf_dce9b7aa96fb1638 := __obf_55924e50211870ef.IntPart()
	__obf_bbd65286a9fac425 := New(1, 0)
	for __obf_6ee7f4c10f6c9355 := __obf_dce9b7aa96fb1638; __obf_6ee7f4c10f6c9355 > 0; __obf_6ee7f4c10f6c9355-- {
		__obf_bbd65286a9fac425 = __obf_bbd65286a9fac425.Mul(__obf_f9decec89e95425e)
	}
	__obf_432a3fe2b3c6e456 := int32(__obf_bbd65286a9fac425.NumDigits())

	var __obf_c0f55061607ec10b int32
	if __obf_432a3fe2b3c6e456 > __obf_d44c132c275234aa(__obf_bbd65286a9fac425.__obf_4163ad30ac75204b) {
		__obf_c0f55061607ec10b = int32(__obf_d7825254824b100e) - __obf_432a3fe2b3c6e456 - __obf_bbd65286a9fac425.__obf_4163ad30ac75204b
	} else {
		__obf_c0f55061607ec10b = int32(__obf_d7825254824b100e)
	}
	__obf_bbd65286a9fac425 = __obf_bbd65286a9fac425.Round(__obf_c0f55061607ec10b)

	return __obf_bbd65286a9fac425, nil
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
func (__obf_3c5564ffc2911a1b Decimal) ExpTaylor(__obf_c0882397220e65d2 int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_3c5564ffc2911a1b.IsZero() {
		return Decimal{__obf_eb38f1915e7fe761, 0}.Round(__obf_c0882397220e65d2), nil
	}

	var __obf_0fa610753799e4e3 Decimal
	var __obf_d31393e4a6da759b int32
	if __obf_c0882397220e65d2 < 0 {
		__obf_0fa610753799e4e3 = New(1, -1)
		__obf_d31393e4a6da759b = 8
	} else {
		__obf_0fa610753799e4e3 = New(1, -__obf_c0882397220e65d2-1)
		__obf_d31393e4a6da759b = __obf_c0882397220e65d2 + 1
	}
	__obf_b839a7e44c62034d := __obf_3c5564ffc2911a1b.Abs()
	__obf_9107901db918d2ca := __obf_3c5564ffc2911a1b.Abs()
	__obf_733f5975b9b9a989 := New(1, 0)
	__obf_15a03af38514d25c := New(1, 0)

	for __obf_6ee7f4c10f6c9355 := int64(1); ; {
		__obf_140a030835afdd9b := __obf_9107901db918d2ca.DivRound(__obf_733f5975b9b9a989, __obf_d31393e4a6da759b)
		__obf_15a03af38514d25c = __obf_15a03af38514d25c.Add(__obf_140a030835afdd9b)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_140a030835afdd9b.Cmp(__obf_0fa610753799e4e3) < 0 {
			break
		}
		__obf_9107901db918d2ca = __obf_9107901db918d2ca.Mul(__obf_b839a7e44c62034d)
		__obf_6ee7f4c10f6c9355++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_1f497da7032f6620) >= int(__obf_6ee7f4c10f6c9355) && !__obf_1f497da7032f6620[__obf_6ee7f4c10f6c9355-1].IsZero() {
			__obf_733f5975b9b9a989 = __obf_1f497da7032f6620[__obf_6ee7f4c10f6c9355-1]
		} else {
			__obf_733f5975b9b9a989 = // To avoid any race conditions, firstly the zero value is appended to a slice to create
				// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
				// factorial using the index notation.
				__obf_1f497da7032f6620[__obf_6ee7f4c10f6c9355-2].Mul(New(__obf_6ee7f4c10f6c9355, 0))
			__obf_1f497da7032f6620 = append(__obf_1f497da7032f6620, Zero)
			__obf_1f497da7032f6620[__obf_6ee7f4c10f6c9355-1] = __obf_733f5975b9b9a989
		}
	}

	if __obf_3c5564ffc2911a1b.Sign() < 0 {
		__obf_15a03af38514d25c = New(1, 0).DivRound(__obf_15a03af38514d25c, __obf_c0882397220e65d2+1)
	}
	__obf_15a03af38514d25c = __obf_15a03af38514d25c.Round(__obf_c0882397220e65d2)
	return __obf_15a03af38514d25c, nil
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
func (__obf_3c5564ffc2911a1b Decimal) Ln(__obf_c0882397220e65d2 int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_3c5564ffc2911a1b.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_3c5564ffc2911a1b.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}
	__obf_4f913605f41d2b61 := __obf_c0882397220e65d2 + 2
	__obf_46a4df43b872c9a9 := __obf_3c5564ffc2911a1b.Copy()

	var __obf_5f8b4d9cc6f2d809, __obf_ed0e8a2ad05321a8, __obf_c5db8e7cd68d347e, __obf_41c59c0c47043825, __obf_bbc6b6df4f5ec458 Decimal
	__obf_5f8b4d9cc6f2d809 = __obf_46a4df43b872c9a9.Sub(Decimal{__obf_eb38f1915e7fe761, 0})
	__obf_ed0e8a2ad05321a8 = Decimal{__obf_eb38f1915e7fe761, -1}
	__obf_7f73a7c20c5bf6e7 := // for decimal in range [0.9, 1.1] where ln(d) is close to 0
		false

	if __obf_5f8b4d9cc6f2d809.Abs().Cmp(__obf_ed0e8a2ad05321a8) <= 0 {
		__obf_7f73a7c20c5bf6e7 = true
	} else {
		__obf_df3ee13d9daef15d := // reduce input decimal to range [0.1, 1)
			int32(__obf_46a4df43b872c9a9.NumDigits()) + __obf_46a4df43b872c9a9.__obf_4163ad30ac75204b

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_46a4df43b872c9a9.__obf_4163ad30ac75204b -= __obf_df3ee13d9daef15d
		__obf_7b8c5b5ebb6c483c := __obf_7b8c5b5ebb6c483c.__obf_9e84e5a13cae3117(__obf_4f913605f41d2b61)
		__obf_bbc6b6df4f5ec458 = NewFromInt32(__obf_df3ee13d9daef15d)
		__obf_bbc6b6df4f5ec458 = __obf_bbc6b6df4f5ec458.Mul(__obf_7b8c5b5ebb6c483c)
		__obf_5f8b4d9cc6f2d809 = __obf_46a4df43b872c9a9.Sub(Decimal{__obf_eb38f1915e7fe761, 0})

		if __obf_5f8b4d9cc6f2d809.Abs().Cmp(__obf_ed0e8a2ad05321a8) <= 0 {
			__obf_7f73a7c20c5bf6e7 = true
		} else {
			__obf_45fa0fe850bf3b81 := // initial estimate using floats
				__obf_46a4df43b872c9a9.InexactFloat64()
			__obf_5f8b4d9cc6f2d809 = NewFromFloat(math.Log(__obf_45fa0fe850bf3b81))
		}
	}
	__obf_0fa610753799e4e3 := Decimal{__obf_eb38f1915e7fe761, -__obf_4f913605f41d2b61}

	if __obf_7f73a7c20c5bf6e7 {
		__obf_c5db8e7cd68d347e = // Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
			// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			// until the difference between current and next term is smaller than epsilon.
			// Coverage quite fast for decimals close to 1.0

			// z + 2
			__obf_5f8b4d9cc6f2d809.Add(Decimal{__obf_b39aaa52669b613d, 0})
		__obf_ed0e8a2ad05321a8 = // z / (z + 2)
			__obf_5f8b4d9cc6f2d809.DivRound(__obf_c5db8e7cd68d347e, __obf_4f913605f41d2b61)
		__obf_5f8b4d9cc6f2d809 = // 2 * (z / (z + 2))
			__obf_ed0e8a2ad05321a8.Add(__obf_ed0e8a2ad05321a8)
		__obf_c5db8e7cd68d347e = __obf_5f8b4d9cc6f2d809.Copy()

		for __obf_0657fff5d9f5aeef := 1; ; __obf_0657fff5d9f5aeef++ {
			__obf_c5db8e7cd68d347e = // 2 * (z / (z+2))^(2n+1)
				__obf_c5db8e7cd68d347e.Mul(__obf_ed0e8a2ad05321a8).Mul(__obf_ed0e8a2ad05321a8)
			__obf_41c59c0c47043825 = // 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
				NewFromInt(int64(2*__obf_0657fff5d9f5aeef + 1))
			__obf_41c59c0c47043825 = __obf_c5db8e7cd68d347e.DivRound(__obf_41c59c0c47043825, __obf_4f913605f41d2b61)
			__obf_5f8b4d9cc6f2d809 = // comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
				__obf_5f8b4d9cc6f2d809.Add(__obf_41c59c0c47043825)

			if __obf_41c59c0c47043825.Abs().Cmp(__obf_0fa610753799e4e3) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_ab69a931ff58dbd8 Decimal
		__obf_ace62018caacd225 := __obf_4f913605f41d2b61*2 + 10

		for __obf_6ee7f4c10f6c9355 := int32(0); __obf_6ee7f4c10f6c9355 < __obf_ace62018caacd225;
		// exp(a_n)
		__obf_6ee7f4c10f6c9355++ {
			__obf_ed0e8a2ad05321a8, _ = __obf_5f8b4d9cc6f2d809.ExpTaylor(__obf_4f913605f41d2b61)
			__obf_c5db8e7cd68d347e = // exp(a_n) - z
				__obf_ed0e8a2ad05321a8.Sub(__obf_46a4df43b872c9a9)
			__obf_c5db8e7cd68d347e = // 2 * (exp(a_n) - z)
				__obf_c5db8e7cd68d347e.Add(__obf_c5db8e7cd68d347e)
			__obf_41c59c0c47043825 = // exp(a_n) + z
				__obf_ed0e8a2ad05321a8.Add(__obf_46a4df43b872c9a9)
			__obf_ed0e8a2ad05321a8 = // 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_c5db8e7cd68d347e.DivRound(__obf_41c59c0c47043825, __obf_4f913605f41d2b61)
			__obf_5f8b4d9cc6f2d809 = // comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
				__obf_5f8b4d9cc6f2d809.Sub(__obf_ed0e8a2ad05321a8)

			if __obf_ab69a931ff58dbd8.Add(__obf_ed0e8a2ad05321a8).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_ed0e8a2ad05321a8.Abs().Cmp(__obf_0fa610753799e4e3) <= 0 {
				break
			}
			__obf_ab69a931ff58dbd8 = __obf_ed0e8a2ad05321a8
		}
	}
	__obf_5f8b4d9cc6f2d809 = __obf_5f8b4d9cc6f2d809.Add(__obf_bbc6b6df4f5ec458)

	return __obf_5f8b4d9cc6f2d809.Round(__obf_c0882397220e65d2), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_3c5564ffc2911a1b Decimal) NumDigits() int {
	if __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e == nil {
		return 1
	}

	if __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.IsInt64() {
		__obf_7be63703093491f0 := __obf_3c5564ffc2911a1b.
			// restrict fast path to integers with exact conversion to float64
			__obf_356f14ec9671961e.Int64()

		if __obf_7be63703093491f0 <= (1<<53) && __obf_7be63703093491f0 >= -(1<<53) {
			if __obf_7be63703093491f0 == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_7be63703093491f0)))) + 1
		}
	}
	__obf_774da70220d34073 := int(float64(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.BitLen()) / math.Log2(10))
	__obf_c0670d78df42aaa5 := // estimatedNumDigits (lg10) may be off by 1, need to verify
		big.NewInt(int64(__obf_774da70220d34073))
	__obf_8061fd17e47b0191 := __obf_c0670d78df42aaa5.Exp(__obf_068b23c822520b43, __obf_c0670d78df42aaa5, nil)

	if __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.CmpAbs(__obf_8061fd17e47b0191) >= 0 {
		return __obf_774da70220d34073 + 1
	}

	return __obf_774da70220d34073
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_3c5564ffc2911a1b Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_2b7b4dcf959f4b21 big.Int
	__obf_58dd48c2daf72a44 := new(big.Int).Set(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e)
	for __obf_46a4df43b872c9a9 := __obf_d44c132c275234aa(__obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b); __obf_46a4df43b872c9a9 > 0; __obf_46a4df43b872c9a9-- {
		__obf_58dd48c2daf72a44.
			QuoRem(__obf_58dd48c2daf72a44, __obf_068b23c822520b43, &__obf_2b7b4dcf959f4b21)
		if __obf_2b7b4dcf959f4b21.Cmp(__obf_6f124f5735d8ea59) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_d44c132c275234aa(__obf_0657fff5d9f5aeef int32) int32 {
	if __obf_0657fff5d9f5aeef < 0 {
		return -__obf_0657fff5d9f5aeef
	}
	return __obf_0657fff5d9f5aeef
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_3c5564ffc2911a1b Decimal) Cmp(__obf_718cb7c55117fa47 Decimal) int {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()
	__obf_718cb7c55117fa47.__obf_4d97a450ece4f814()

	if __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b == __obf_718cb7c55117fa47.__obf_4163ad30ac75204b {
		return __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.Cmp(__obf_718cb7c55117fa47.__obf_356f14ec9671961e)
	}
	__obf_9f7dbbc061b0788c, __obf_ef5ba9ca72c05c92 := RescalePair(__obf_3c5564ffc2911a1b, __obf_718cb7c55117fa47)

	return __obf_9f7dbbc061b0788c.__obf_356f14ec9671961e.Cmp(__obf_ef5ba9ca72c05c92.

		// Compare compares the numbers represented by d and d2 and returns:
		//
		//	-1 if d <  d2
		//	 0 if d == d2
		//	+1 if d >  d2
		__obf_356f14ec9671961e)
}

func (__obf_3c5564ffc2911a1b Decimal) Compare(__obf_718cb7c55117fa47 Decimal) int {
	return __obf_3c5564ffc2911a1b.Cmp(__obf_718cb7c55117fa47)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_3c5564ffc2911a1b Decimal) Equal(__obf_718cb7c55117fa47 Decimal) bool {
	return __obf_3c5564ffc2911a1b.Cmp(__obf_718cb7c55117fa47) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_3c5564ffc2911a1b Decimal) Equals(__obf_718cb7c55117fa47 Decimal) bool {
	return __obf_3c5564ffc2911a1b.Equal(__obf_718cb7c55117fa47)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_3c5564ffc2911a1b Decimal) GreaterThan(__obf_718cb7c55117fa47 Decimal) bool {
	return __obf_3c5564ffc2911a1b.Cmp(__obf_718cb7c55117fa47) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_3c5564ffc2911a1b Decimal) GreaterThanOrEqual(__obf_718cb7c55117fa47 Decimal) bool {
	__obf_bdfac7c27eb78da2 := __obf_3c5564ffc2911a1b.Cmp(__obf_718cb7c55117fa47)
	return __obf_bdfac7c27eb78da2 == 1 || __obf_bdfac7c27eb78da2 == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_3c5564ffc2911a1b Decimal) LessThan(__obf_718cb7c55117fa47 Decimal) bool {
	return __obf_3c5564ffc2911a1b.Cmp(__obf_718cb7c55117fa47) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_3c5564ffc2911a1b Decimal) LessThanOrEqual(__obf_718cb7c55117fa47 Decimal) bool {
	__obf_bdfac7c27eb78da2 := __obf_3c5564ffc2911a1b.Cmp(__obf_718cb7c55117fa47)
	return __obf_bdfac7c27eb78da2 == -1 || __obf_bdfac7c27eb78da2 == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_3c5564ffc2911a1b Decimal) Sign() int {
	if __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e == nil {
		return 0
	}
	return __obf_3c5564ffc2911a1b.

		// IsPositive return
		//
		//	true if d > 0
		//	false if d == 0
		//	false if d < 0
		__obf_356f14ec9671961e.Sign()
}

func (__obf_3c5564ffc2911a1b Decimal) IsPositive() bool {
	return __obf_3c5564ffc2911a1b.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_3c5564ffc2911a1b Decimal) IsNegative() bool {
	return __obf_3c5564ffc2911a1b.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_3c5564ffc2911a1b Decimal) IsZero() bool {
	return __obf_3c5564ffc2911a1b.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_3c5564ffc2911a1b Decimal) Exponent() int32 {
	return __obf_3c5564ffc2911a1b.

		// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
		__obf_4163ad30ac75204b
}

func (__obf_3c5564ffc2911a1b Decimal) Coefficient() *big.Int {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_3c5564ffc2911a1b.

		// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
		// If coefficient cannot be represented in an int64, the result will be undefined.
		__obf_356f14ec9671961e)
}

func (__obf_3c5564ffc2911a1b Decimal) CoefficientInt64() int64 {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()
	return __obf_3c5564ffc2911a1b.

		// IntPart returns the integer component of the decimal.
		__obf_356f14ec9671961e.Int64()
}

func (__obf_3c5564ffc2911a1b Decimal) IntPart() int64 {
	__obf_94f2cdaa03c907e8 := __obf_3c5564ffc2911a1b.__obf_cfec7d8af9246d7b(0)
	return __obf_94f2cdaa03c907e8.__obf_356f14ec9671961e.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_3c5564ffc2911a1b Decimal) BigInt() *big.Int {
	__obf_94f2cdaa03c907e8 := __obf_3c5564ffc2911a1b.__obf_cfec7d8af9246d7b(0)
	return __obf_94f2cdaa03c907e8.

		// BigFloat returns decimal as BigFloat.
		// Be aware that casting decimal to BigFloat might cause a loss of precision.
		__obf_356f14ec9671961e
}

func (__obf_3c5564ffc2911a1b Decimal) BigFloat() *big.Float {
	__obf_cafcc62d0a40df85 := &big.Float{}
	__obf_cafcc62d0a40df85.
		SetString(__obf_3c5564ffc2911a1b.String())
	return __obf_cafcc62d0a40df85
}

// Rat returns a rational number representation of the decimal.
func (__obf_3c5564ffc2911a1b Decimal) Rat() *big.Rat {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()
	if __obf_3c5564ffc2911a1b.
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_4163ad30ac75204b <= 0 {
		__obf_1ab9caaa34315e83 := new(big.Int).Exp(__obf_068b23c822520b43, big.NewInt(-int64(__obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b)), nil)
		return new(big.Rat).SetFrac(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e, __obf_1ab9caaa34315e83)
	}
	__obf_35f5ed2a742d8d53 := new(big.Int).Exp(__obf_068b23c822520b43, big.NewInt(int64(__obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b)), nil)
	__obf_1f1ce9118a448fdb := new(big.Int).Mul(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e, __obf_35f5ed2a742d8d53)
	return new(big.Rat).SetFrac(__obf_1f1ce9118a448fdb,

		// Float64 returns the nearest float64 value for d and a bool indicating
		// whether f represents d exactly.
		// For more details, see the documentation for big.Rat.Float64
		__obf_eb38f1915e7fe761)
}

func (__obf_3c5564ffc2911a1b Decimal) Float64() (__obf_cafcc62d0a40df85 float64, __obf_da5dd743f3db388d bool) {
	return __obf_3c5564ffc2911a1b.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_3c5564ffc2911a1b Decimal) InexactFloat64() float64 {
	__obf_cafcc62d0a40df85, _ := __obf_3c5564ffc2911a1b.Float64()
	return __obf_cafcc62d0a40df85
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
func (__obf_3c5564ffc2911a1b Decimal) String() string {
	return __obf_3c5564ffc2911a1b.string(true)
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
func (__obf_3c5564ffc2911a1b Decimal) StringFixed(__obf_a335286ee0a281fb int32) string {
	__obf_1a00a389a971fa77 := __obf_3c5564ffc2911a1b.Round(__obf_a335286ee0a281fb)
	return __obf_1a00a389a971fa77.string(false)
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
func (__obf_3c5564ffc2911a1b Decimal) StringFixedBank(__obf_a335286ee0a281fb int32) string {
	__obf_1a00a389a971fa77 := __obf_3c5564ffc2911a1b.RoundBank(__obf_a335286ee0a281fb)
	return __obf_1a00a389a971fa77.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_3c5564ffc2911a1b Decimal) StringFixedCash(__obf_b869fcfe822fcf3c uint8) string {
	__obf_1a00a389a971fa77 := __obf_3c5564ffc2911a1b.RoundCash(__obf_b869fcfe822fcf3c)
	return __obf_1a00a389a971fa77.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_3c5564ffc2911a1b Decimal) Round(__obf_a335286ee0a281fb int32) Decimal {
	if __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b == -__obf_a335286ee0a281fb {
		return __obf_3c5564ffc2911a1b
	}
	__obf_f03362e651cba3bf := // truncate to places + 1
		__obf_3c5564ffc2911a1b.__obf_cfec7d8af9246d7b(-__obf_a335286ee0a281fb - 1)

	// add sign(d) * 0.5
	if __obf_f03362e651cba3bf.__obf_356f14ec9671961e.Sign() < 0 {
		__obf_f03362e651cba3bf.__obf_356f14ec9671961e.
			Sub(__obf_f03362e651cba3bf.__obf_356f14ec9671961e, __obf_e52ae739ff41f438)
	} else {
		__obf_f03362e651cba3bf.__obf_356f14ec9671961e.
			Add(__obf_f03362e651cba3bf.__obf_356f14ec9671961e,

				// floor for positive numbers, ceil for negative numbers
				__obf_e52ae739ff41f438)
	}

	_, __obf_2c8d42b08bf1a2b7 := __obf_f03362e651cba3bf.__obf_356f14ec9671961e.DivMod(__obf_f03362e651cba3bf.__obf_356f14ec9671961e, __obf_068b23c822520b43, new(big.Int))
	__obf_f03362e651cba3bf.__obf_4163ad30ac75204b++
	if __obf_f03362e651cba3bf.__obf_356f14ec9671961e.Sign() < 0 && __obf_2c8d42b08bf1a2b7.Cmp(__obf_6f124f5735d8ea59) != 0 {
		__obf_f03362e651cba3bf.__obf_356f14ec9671961e.
			Add(__obf_f03362e651cba3bf.__obf_356f14ec9671961e,

				// RoundCeil rounds the decimal towards +infinity.
				//
				// Example:
				//
				//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
				//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
				//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
				//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
				__obf_eb38f1915e7fe761)
	}

	return __obf_f03362e651cba3bf
}

func (__obf_3c5564ffc2911a1b Decimal) RoundCeil(__obf_a335286ee0a281fb int32) Decimal {
	if __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b >= -__obf_a335286ee0a281fb {
		return __obf_3c5564ffc2911a1b
	}
	__obf_1d3824ab17ee71a7 := __obf_3c5564ffc2911a1b.__obf_cfec7d8af9246d7b(-__obf_a335286ee0a281fb)
	if __obf_3c5564ffc2911a1b.Equal(__obf_1d3824ab17ee71a7) {
		return __obf_3c5564ffc2911a1b
	}

	if __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.Sign() > 0 {
		__obf_1d3824ab17ee71a7.__obf_356f14ec9671961e.
			Add(__obf_1d3824ab17ee71a7.__obf_356f14ec9671961e, __obf_eb38f1915e7fe761)
	}

	return __obf_1d3824ab17ee71a7
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_3c5564ffc2911a1b Decimal) RoundFloor(__obf_a335286ee0a281fb int32) Decimal {
	if __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b >= -__obf_a335286ee0a281fb {
		return __obf_3c5564ffc2911a1b
	}
	__obf_1d3824ab17ee71a7 := __obf_3c5564ffc2911a1b.__obf_cfec7d8af9246d7b(-__obf_a335286ee0a281fb)
	if __obf_3c5564ffc2911a1b.Equal(__obf_1d3824ab17ee71a7) {
		return __obf_3c5564ffc2911a1b
	}

	if __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.Sign() < 0 {
		__obf_1d3824ab17ee71a7.__obf_356f14ec9671961e.
			Sub(__obf_1d3824ab17ee71a7.__obf_356f14ec9671961e, __obf_eb38f1915e7fe761)
	}

	return __obf_1d3824ab17ee71a7
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_3c5564ffc2911a1b Decimal) RoundUp(__obf_a335286ee0a281fb int32) Decimal {
	if __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b >= -__obf_a335286ee0a281fb {
		return __obf_3c5564ffc2911a1b
	}
	__obf_1d3824ab17ee71a7 := __obf_3c5564ffc2911a1b.__obf_cfec7d8af9246d7b(-__obf_a335286ee0a281fb)
	if __obf_3c5564ffc2911a1b.Equal(__obf_1d3824ab17ee71a7) {
		return __obf_3c5564ffc2911a1b
	}

	if __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.Sign() > 0 {
		__obf_1d3824ab17ee71a7.__obf_356f14ec9671961e.
			Add(__obf_1d3824ab17ee71a7.__obf_356f14ec9671961e, __obf_eb38f1915e7fe761)
	} else if __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.Sign() < 0 {
		__obf_1d3824ab17ee71a7.__obf_356f14ec9671961e.
			Sub(__obf_1d3824ab17ee71a7.__obf_356f14ec9671961e, __obf_eb38f1915e7fe761)
	}

	return __obf_1d3824ab17ee71a7
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_3c5564ffc2911a1b Decimal) RoundDown(__obf_a335286ee0a281fb int32) Decimal {
	if __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b >= -__obf_a335286ee0a281fb {
		return __obf_3c5564ffc2911a1b
	}
	__obf_1d3824ab17ee71a7 := __obf_3c5564ffc2911a1b.__obf_cfec7d8af9246d7b(-__obf_a335286ee0a281fb)
	if __obf_3c5564ffc2911a1b.Equal(__obf_1d3824ab17ee71a7) {
		return __obf_3c5564ffc2911a1b
	}
	return __obf_1d3824ab17ee71a7
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
func (__obf_3c5564ffc2911a1b Decimal) RoundBank(__obf_a335286ee0a281fb int32) Decimal {
	__obf_ac6cc9813268c94c := __obf_3c5564ffc2911a1b.Round(__obf_a335286ee0a281fb)
	__obf_41dc39418e97b0bd := __obf_3c5564ffc2911a1b.Sub(__obf_ac6cc9813268c94c).Abs()
	__obf_49328a7dfcb24113 := New(5, -__obf_a335286ee0a281fb-1)
	if __obf_41dc39418e97b0bd.Cmp(__obf_49328a7dfcb24113) == 0 && __obf_ac6cc9813268c94c.__obf_356f14ec9671961e.Bit(0) != 0 {
		if __obf_ac6cc9813268c94c.__obf_356f14ec9671961e.Sign() < 0 {
			__obf_ac6cc9813268c94c.__obf_356f14ec9671961e.
				Add(__obf_ac6cc9813268c94c.__obf_356f14ec9671961e, __obf_eb38f1915e7fe761)
		} else {
			__obf_ac6cc9813268c94c.__obf_356f14ec9671961e.
				Sub(__obf_ac6cc9813268c94c.__obf_356f14ec9671961e, __obf_eb38f1915e7fe761)
		}
	}

	return __obf_ac6cc9813268c94c
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
func (__obf_3c5564ffc2911a1b Decimal) RoundCash(__obf_b869fcfe822fcf3c uint8) Decimal {
	var __obf_03a92b73be553c97 *big.Int
	switch __obf_b869fcfe822fcf3c {
	case 5:
		__obf_03a92b73be553c97 = __obf_f19af932fedccc2d
	case 10:
		__obf_03a92b73be553c97 = __obf_068b23c822520b43
	case 25:
		__obf_03a92b73be553c97 = __obf_ac06b4a01c95937c
	case 50:
		__obf_03a92b73be553c97 = __obf_b39aaa52669b613d
	case 100:
		__obf_03a92b73be553c97 = __obf_eb38f1915e7fe761
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_b869fcfe822fcf3c))
	}
	__obf_3985bb55463148e6 := Decimal{__obf_356f14ec9671961e: __obf_03a92b73be553c97}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_3c5564ffc2911a1b.Mul(__obf_3985bb55463148e6).Round(0).Div(__obf_3985bb55463148e6).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_3c5564ffc2911a1b Decimal) Floor() Decimal {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()

	if __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b >= 0 {
		return __obf_3c5564ffc2911a1b
	}
	__obf_4163ad30ac75204b := big.NewInt(10)
	__obf_4163ad30ac75204b.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_4163ad30ac75204b, big.NewInt(-int64(__obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b)), nil)
	__obf_46a4df43b872c9a9 := new(big.Int).Div(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e, __obf_4163ad30ac75204b)
	return Decimal{__obf_356f14ec9671961e: __obf_46a4df43b872c9a9,

		// Ceil returns the nearest integer value greater than or equal to d.
		__obf_4163ad30ac75204b: 0}
}

func (__obf_3c5564ffc2911a1b Decimal) Ceil() Decimal {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()

	if __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b >= 0 {
		return __obf_3c5564ffc2911a1b
	}
	__obf_4163ad30ac75204b := big.NewInt(10)
	__obf_4163ad30ac75204b.

		// NOTE(vadim): must negate after casting to prevent int32 overflow
		Exp(__obf_4163ad30ac75204b, big.NewInt(-int64(__obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b)), nil)
	__obf_46a4df43b872c9a9, __obf_2c8d42b08bf1a2b7 := new(big.Int).DivMod(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e, __obf_4163ad30ac75204b, new(big.Int))
	if __obf_2c8d42b08bf1a2b7.Cmp(__obf_6f124f5735d8ea59) != 0 {
		__obf_46a4df43b872c9a9.
			Add(__obf_46a4df43b872c9a9, __obf_eb38f1915e7fe761)
	}
	return Decimal{__obf_356f14ec9671961e: __obf_46a4df43b872c9a9,

		// Truncate truncates off digits from the number, without rounding.
		//
		// NOTE: precision is the last digit that will not be truncated (must be >= 0).
		//
		// Example:
		//
		//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
		__obf_4163ad30ac75204b: 0}
}

func (__obf_3c5564ffc2911a1b Decimal) Truncate(__obf_c0882397220e65d2 int32) Decimal {
	__obf_3c5564ffc2911a1b.__obf_4d97a450ece4f814()
	if __obf_c0882397220e65d2 >= 0 && -__obf_c0882397220e65d2 > __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b {
		return __obf_3c5564ffc2911a1b.__obf_cfec7d8af9246d7b(-__obf_c0882397220e65d2)
	}
	return __obf_3c5564ffc2911a1b
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_3c5564ffc2911a1b *Decimal) UnmarshalJSON(__obf_215ab4f8efdc7615 []byte) error {
	if string(__obf_215ab4f8efdc7615) == "null" {
		return nil
	}
	__obf_fa7765e795d22460, __obf_25b3ae434b8d1db5 := __obf_af31ce3ac4ddea5d(__obf_215ab4f8efdc7615)
	if __obf_25b3ae434b8d1db5 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_215ab4f8efdc7615, __obf_25b3ae434b8d1db5)
	}
	__obf_ebbf7bb33b20bf6e, __obf_25b3ae434b8d1db5 := NewFromString(__obf_fa7765e795d22460)
	*__obf_3c5564ffc2911a1b = __obf_ebbf7bb33b20bf6e
	if __obf_25b3ae434b8d1db5 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_fa7765e795d22460, __obf_25b3ae434b8d1db5)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_3c5564ffc2911a1b Decimal) MarshalJSON() ([]byte, error) {
	var __obf_fa7765e795d22460 string
	if MarshalJSONWithoutQuotes {
		__obf_fa7765e795d22460 = __obf_3c5564ffc2911a1b.String()
	} else {
		__obf_fa7765e795d22460 = "\"" + __obf_3c5564ffc2911a1b.String() + "\""
	}
	return []byte(__obf_fa7765e795d22460), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_3c5564ffc2911a1b *Decimal) UnmarshalBinary(__obf_1518897804a8de82 []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_1518897804a8de82) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_1518897804a8de82, len(__obf_1518897804a8de82))
	}
	__obf_3c5564ffc2911a1b.

		// Extract the exponent
		__obf_4163ad30ac75204b = int32(binary.BigEndian.Uint32(__obf_1518897804a8de82[:4]))
	__obf_3c5564ffc2911a1b.

		// Extract the value
		__obf_356f14ec9671961e = new(big.Int)
	if __obf_25b3ae434b8d1db5 := __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.GobDecode(__obf_1518897804a8de82[4:]); __obf_25b3ae434b8d1db5 != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_1518897804a8de82, __obf_25b3ae434b8d1db5)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_3c5564ffc2911a1b Decimal) MarshalBinary() (__obf_1518897804a8de82 []byte, __obf_25b3ae434b8d1db5 error) {
	// exp is written first, but encode value first to know output size
	var __obf_7992a00e8ab2b9fd []byte
	if __obf_7992a00e8ab2b9fd, __obf_25b3ae434b8d1db5 = __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.GobEncode(); __obf_25b3ae434b8d1db5 != nil {
		return nil, __obf_25b3ae434b8d1db5
	}
	__obf_c9f694ad1d7986e5 := // Write the exponent in front, since it's a fixed size
		make([]byte, 4, len(__obf_7992a00e8ab2b9fd)+4)
	binary.BigEndian.PutUint32(__obf_c9f694ad1d7986e5, uint32(__obf_3c5564ffc2911a1b.

		// Return the byte array
		__obf_4163ad30ac75204b))

	return append(__obf_c9f694ad1d7986e5, __obf_7992a00e8ab2b9fd...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_3c5564ffc2911a1b *Decimal) Scan(__obf_356f14ec9671961e any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_3785ba6483ed0c4c := __obf_356f14ec9671961e.(type) {

	case float32:
		*__obf_3c5564ffc2911a1b = NewFromFloat(float64(__obf_3785ba6483ed0c4c))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_3c5564ffc2911a1b = NewFromFloat(__obf_3785ba6483ed0c4c)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_3c5564ffc2911a1b = New(__obf_3785ba6483ed0c4c, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_3c5564ffc2911a1b = NewFromUint64(__obf_3785ba6483ed0c4c)
		return nil

	default:
		__obf_fa7765e795d22460,
			// default is trying to interpret value stored as string
			__obf_25b3ae434b8d1db5 := __obf_af31ce3ac4ddea5d(__obf_3785ba6483ed0c4c)
		if __obf_25b3ae434b8d1db5 != nil {
			return __obf_25b3ae434b8d1db5
		}
		*__obf_3c5564ffc2911a1b, __obf_25b3ae434b8d1db5 = NewFromString(__obf_fa7765e795d22460)
		return __obf_25b3ae434b8d1db5
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_3c5564ffc2911a1b Decimal) Value() (driver.Value, error) {
	return __obf_3c5564ffc2911a1b.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_3c5564ffc2911a1b *Decimal) UnmarshalText(__obf_9ae31ba998bbea2d []byte) error {
	__obf_fa7765e795d22460 := string(__obf_9ae31ba998bbea2d)
	__obf_9397222af603e4cf, __obf_25b3ae434b8d1db5 := NewFromString(__obf_fa7765e795d22460)
	*__obf_3c5564ffc2911a1b = __obf_9397222af603e4cf
	if __obf_25b3ae434b8d1db5 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_fa7765e795d22460, __obf_25b3ae434b8d1db5)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_3c5564ffc2911a1b Decimal) MarshalText() (__obf_9ae31ba998bbea2d []byte, __obf_25b3ae434b8d1db5 error) {
	return []byte(__obf_3c5564ffc2911a1b.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_3c5564ffc2911a1b Decimal) GobEncode() ([]byte, error) {
	return __obf_3c5564ffc2911a1b.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_3c5564ffc2911a1b *Decimal) GobDecode(__obf_1518897804a8de82 []byte) error {
	return __obf_3c5564ffc2911a1b.UnmarshalBinary(__obf_1518897804a8de82)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_3c5564ffc2911a1b Decimal) StringScaled(__obf_4163ad30ac75204b int32) string {
	return __obf_3c5564ffc2911a1b.__obf_cfec7d8af9246d7b(__obf_4163ad30ac75204b).String()
}

func (__obf_3c5564ffc2911a1b Decimal) string(__obf_1e56127232cf5ee4 bool) string {
	if __obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b >= 0 {
		return __obf_3c5564ffc2911a1b.__obf_cfec7d8af9246d7b(0).__obf_356f14ec9671961e.String()
	}
	__obf_d44c132c275234aa := new(big.Int).Abs(__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e)
	__obf_fa7765e795d22460 := __obf_d44c132c275234aa.String()

	var __obf_d58be7493886a7ca, __obf_e51ac97d44b21eef string
	__obf_86f312229a76f67d := // NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
		// and you are on a 32-bit machine. Won't fix this super-edge case.
		int(__obf_3c5564ffc2911a1b.__obf_4163ad30ac75204b)
	if len(__obf_fa7765e795d22460) > -__obf_86f312229a76f67d {
		__obf_d58be7493886a7ca = __obf_fa7765e795d22460[:len(__obf_fa7765e795d22460)+__obf_86f312229a76f67d]
		__obf_e51ac97d44b21eef = __obf_fa7765e795d22460[len(__obf_fa7765e795d22460)+__obf_86f312229a76f67d:]
	} else {
		__obf_d58be7493886a7ca = "0"
		__obf_a4ad13ad5529de71 := -__obf_86f312229a76f67d - len(__obf_fa7765e795d22460)
		__obf_e51ac97d44b21eef = strings.Repeat("0", __obf_a4ad13ad5529de71) + __obf_fa7765e795d22460
	}

	if __obf_1e56127232cf5ee4 {
		__obf_6ee7f4c10f6c9355 := len(__obf_e51ac97d44b21eef) - 1
		for ; __obf_6ee7f4c10f6c9355 >= 0; __obf_6ee7f4c10f6c9355-- {
			if __obf_e51ac97d44b21eef[__obf_6ee7f4c10f6c9355] != '0' {
				break
			}
		}
		__obf_e51ac97d44b21eef = __obf_e51ac97d44b21eef[:__obf_6ee7f4c10f6c9355+1]
	}
	__obf_682e71c86d819364 := __obf_d58be7493886a7ca
	if len(__obf_e51ac97d44b21eef) > 0 {
		__obf_682e71c86d819364 += "." + __obf_e51ac97d44b21eef
	}

	if __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e.Sign() < 0 {
		return "-" + __obf_682e71c86d819364
	}

	return __obf_682e71c86d819364
}

func (__obf_3c5564ffc2911a1b *Decimal) __obf_4d97a450ece4f814() {
	if __obf_3c5564ffc2911a1b.__obf_356f14ec9671961e == nil {
		__obf_3c5564ffc2911a1b.__obf_356f14ec9671961e = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_53bc756179e4e207 Decimal, __obf_8408781cd4562b46 ...Decimal) Decimal {
	__obf_061e950f3bd6226c := __obf_53bc756179e4e207
	for _, __obf_504a31427ce6ef66 := range __obf_8408781cd4562b46 {
		if __obf_504a31427ce6ef66.Cmp(__obf_061e950f3bd6226c) < 0 {
			__obf_061e950f3bd6226c = __obf_504a31427ce6ef66
		}
	}
	return __obf_061e950f3bd6226c
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_53bc756179e4e207 Decimal, __obf_8408781cd4562b46 ...Decimal) Decimal {
	__obf_061e950f3bd6226c := __obf_53bc756179e4e207
	for _, __obf_504a31427ce6ef66 := range __obf_8408781cd4562b46 {
		if __obf_504a31427ce6ef66.Cmp(__obf_061e950f3bd6226c) > 0 {
			__obf_061e950f3bd6226c = __obf_504a31427ce6ef66
		}
	}
	return __obf_061e950f3bd6226c
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_53bc756179e4e207 Decimal, __obf_8408781cd4562b46 ...Decimal) Decimal {
	__obf_af44409a01922ef9 := __obf_53bc756179e4e207
	for _, __obf_504a31427ce6ef66 := range __obf_8408781cd4562b46 {
		__obf_af44409a01922ef9 = __obf_af44409a01922ef9.Add(__obf_504a31427ce6ef66)
	}

	return __obf_af44409a01922ef9
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_53bc756179e4e207 Decimal, __obf_8408781cd4562b46 ...Decimal) Decimal {
	__obf_288377529a03d116 := New(int64(len(__obf_8408781cd4562b46)+1), 0)
	__obf_f9decec89e95425e := Sum(__obf_53bc756179e4e207, __obf_8408781cd4562b46...)
	return __obf_f9decec89e95425e.Div(__obf_288377529a03d116)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_b327d68f66614e91 Decimal, __obf_718cb7c55117fa47 Decimal) (Decimal, Decimal) {
	__obf_b327d68f66614e91.__obf_4d97a450ece4f814()
	__obf_718cb7c55117fa47.__obf_4d97a450ece4f814()

	if __obf_b327d68f66614e91.__obf_4163ad30ac75204b < __obf_718cb7c55117fa47.__obf_4163ad30ac75204b {
		return __obf_b327d68f66614e91, __obf_718cb7c55117fa47.__obf_cfec7d8af9246d7b(__obf_b327d68f66614e91.__obf_4163ad30ac75204b)
	} else if __obf_b327d68f66614e91.__obf_4163ad30ac75204b > __obf_718cb7c55117fa47.__obf_4163ad30ac75204b {
		return __obf_b327d68f66614e91.__obf_cfec7d8af9246d7b(__obf_718cb7c55117fa47.__obf_4163ad30ac75204b), __obf_718cb7c55117fa47
	}

	return __obf_b327d68f66614e91, __obf_718cb7c55117fa47
}

func __obf_af31ce3ac4ddea5d(__obf_356f14ec9671961e any) (string, error) {
	var __obf_20ce5e27d41aed81 []byte

	switch __obf_3785ba6483ed0c4c := __obf_356f14ec9671961e.(type) {
	case string:
		__obf_20ce5e27d41aed81 = []byte(__obf_3785ba6483ed0c4c)
	case []byte:
		__obf_20ce5e27d41aed81 = __obf_3785ba6483ed0c4c
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_356f14ec9671961e,

			// If the amount is quoted, strip the quotes
			__obf_356f14ec9671961e)
	}

	if len(__obf_20ce5e27d41aed81) > 2 && __obf_20ce5e27d41aed81[0] == '"' && __obf_20ce5e27d41aed81[len(__obf_20ce5e27d41aed81)-1] == '"' {
		__obf_20ce5e27d41aed81 = __obf_20ce5e27d41aed81[1 : len(__obf_20ce5e27d41aed81)-1]
	}
	return string(__obf_20ce5e27d41aed81), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the datautil.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_3c5564ffc2911a1b Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_3c5564ffc2911a1b,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_3c5564ffc2911a1b *NullDecimal) Scan(__obf_356f14ec9671961e any) error {
	if __obf_356f14ec9671961e == nil {
		__obf_3c5564ffc2911a1b.
			Valid = false
		return nil
	}
	__obf_3c5564ffc2911a1b.
		Valid = true
	return __obf_3c5564ffc2911a1b.Decimal.Scan(__obf_356f14ec9671961e)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_3c5564ffc2911a1b NullDecimal) Value() (driver.Value, error) {
	if !__obf_3c5564ffc2911a1b.Valid {
		return nil, nil
	}
	return __obf_3c5564ffc2911a1b.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_3c5564ffc2911a1b *NullDecimal) UnmarshalJSON(__obf_215ab4f8efdc7615 []byte) error {
	if string(__obf_215ab4f8efdc7615) == "null" {
		__obf_3c5564ffc2911a1b.
			Valid = false
		return nil
	}
	__obf_3c5564ffc2911a1b.
		Valid = true
	return __obf_3c5564ffc2911a1b.Decimal.UnmarshalJSON(__obf_215ab4f8efdc7615)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_3c5564ffc2911a1b NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_3c5564ffc2911a1b.Valid {
		return []byte("null"), nil
	}
	return __obf_3c5564ffc2911a1b.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_3c5564ffc2911a1b *NullDecimal) UnmarshalText(__obf_9ae31ba998bbea2d []byte) error {
	__obf_fa7765e795d22460 := string(__obf_9ae31ba998bbea2d)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_fa7765e795d22460 == "" {
		__obf_3c5564ffc2911a1b.
			Valid = false
		return nil
	}
	if __obf_25b3ae434b8d1db5 := __obf_3c5564ffc2911a1b.Decimal.UnmarshalText(__obf_9ae31ba998bbea2d); __obf_25b3ae434b8d1db5 != nil {
		__obf_3c5564ffc2911a1b.
			Valid = false
		return __obf_25b3ae434b8d1db5
	}
	__obf_3c5564ffc2911a1b.
		Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_3c5564ffc2911a1b NullDecimal) MarshalText() (__obf_9ae31ba998bbea2d []byte, __obf_25b3ae434b8d1db5 error) {
	if !__obf_3c5564ffc2911a1b.Valid {
		return []byte{}, nil
	}
	return __obf_3c5564ffc2911a1b.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_3c5564ffc2911a1b Decimal) Atan() Decimal {
	if __obf_3c5564ffc2911a1b.Equal(NewFromFloat(0.0)) {
		return __obf_3c5564ffc2911a1b
	}
	if __obf_3c5564ffc2911a1b.GreaterThan(NewFromFloat(0.0)) {
		return __obf_3c5564ffc2911a1b.__obf_3bbdbbd26514948a()
	}
	return __obf_3c5564ffc2911a1b.Neg().__obf_3bbdbbd26514948a().Neg()
}

func (__obf_3c5564ffc2911a1b Decimal) __obf_919dfbc73168f177() Decimal {
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
	__obf_46a4df43b872c9a9 := __obf_3c5564ffc2911a1b.Mul(__obf_3c5564ffc2911a1b)
	__obf_ff7a67ba9ac4d95e := P0.Mul(__obf_46a4df43b872c9a9).Add(P1).Mul(__obf_46a4df43b872c9a9).Add(P2).Mul(__obf_46a4df43b872c9a9).Add(P3).Mul(__obf_46a4df43b872c9a9).Add(P4).Mul(__obf_46a4df43b872c9a9)
	__obf_276841098379287d := __obf_46a4df43b872c9a9.Add(Q0).Mul(__obf_46a4df43b872c9a9).Add(Q1).Mul(__obf_46a4df43b872c9a9).Add(Q2).Mul(__obf_46a4df43b872c9a9).Add(Q3).Mul(__obf_46a4df43b872c9a9).Add(Q4)
	__obf_46a4df43b872c9a9 = __obf_ff7a67ba9ac4d95e.Div(__obf_276841098379287d)
	__obf_46a4df43b872c9a9 = __obf_3c5564ffc2911a1b.Mul(__obf_46a4df43b872c9a9).Add(__obf_3c5564ffc2911a1b)
	return __obf_46a4df43b872c9a9
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_3c5564ffc2911a1b Decimal) __obf_3bbdbbd26514948a() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)
	__obf_d2f0c6b765359ab0 := // tan(3*pi/8)
		NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_3c5564ffc2911a1b.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_3c5564ffc2911a1b.__obf_919dfbc73168f177()
	}
	if __obf_3c5564ffc2911a1b.GreaterThan(Tan3pio8) {
		return __obf_d2f0c6b765359ab0.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_3c5564ffc2911a1b).__obf_919dfbc73168f177()).Add(Morebits)
	}
	return __obf_d2f0c6b765359ab0.Div(NewFromFloat(4.0)).Add((__obf_3c5564ffc2911a1b.Sub(NewFromFloat(1.0)).Div(__obf_3c5564ffc2911a1b.Add(NewFromFloat(1.0)))).__obf_919dfbc73168f177()).Add(NewFromFloat(0.5).Mul(Morebits))
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
func (__obf_3c5564ffc2911a1b Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_3c5564ffc2911a1b.Equal(NewFromFloat(0.0)) {
		return __obf_3c5564ffc2911a1b
	}
	__obf_07c356752a4027b3 := // make argument positive but save the sign
		false
	if __obf_3c5564ffc2911a1b.LessThan(NewFromFloat(0.0)) {
		__obf_3c5564ffc2911a1b = __obf_3c5564ffc2911a1b.Neg()
		__obf_07c356752a4027b3 = true
	}
	__obf_e4e686f2e7c57b0a := __obf_3c5564ffc2911a1b.Mul(M4PI).IntPart()
	__obf_c9ce1b5e622f992c := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_e4e686f2e7c57b0a)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_e4e686f2e7c57b0a&1 == 1 {
		__obf_e4e686f2e7c57b0a++
		__obf_c9ce1b5e622f992c = __obf_c9ce1b5e622f992c.Add(NewFromFloat(1.0))
	}
	__obf_e4e686f2e7c57b0a &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_e4e686f2e7c57b0a > 3 {
		__obf_07c356752a4027b3 = !__obf_07c356752a4027b3
		__obf_e4e686f2e7c57b0a -= 4
	}
	__obf_46a4df43b872c9a9 := __obf_3c5564ffc2911a1b.Sub(__obf_c9ce1b5e622f992c.Mul(PI4A)).Sub(__obf_c9ce1b5e622f992c.Mul(PI4B)).Sub(__obf_c9ce1b5e622f992c.Mul(PI4C))
	__obf_ef73b20430ae076f := // Extended precision modular arithmetic
		__obf_46a4df43b872c9a9.Mul(__obf_46a4df43b872c9a9)

	if __obf_e4e686f2e7c57b0a == 1 || __obf_e4e686f2e7c57b0a == 2 {
		__obf_d6ae9e541a607f97 := __obf_ef73b20430ae076f.Mul(__obf_ef73b20430ae076f).Mul(_cos[0].Mul(__obf_ef73b20430ae076f).Add(_cos[1]).Mul(__obf_ef73b20430ae076f).Add(_cos[2]).Mul(__obf_ef73b20430ae076f).Add(_cos[3]).Mul(__obf_ef73b20430ae076f).Add(_cos[4]).Mul(__obf_ef73b20430ae076f).Add(_cos[5]))
		__obf_c9ce1b5e622f992c = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_ef73b20430ae076f)).Add(__obf_d6ae9e541a607f97)
	} else {
		__obf_c9ce1b5e622f992c = __obf_46a4df43b872c9a9.Add(__obf_46a4df43b872c9a9.Mul(__obf_ef73b20430ae076f).Mul(_sin[0].Mul(__obf_ef73b20430ae076f).Add(_sin[1]).Mul(__obf_ef73b20430ae076f).Add(_sin[2]).Mul(__obf_ef73b20430ae076f).Add(_sin[3]).Mul(__obf_ef73b20430ae076f).Add(_sin[4]).Mul(__obf_ef73b20430ae076f).Add(_sin[5])))
	}
	if __obf_07c356752a4027b3 {
		__obf_c9ce1b5e622f992c = __obf_c9ce1b5e622f992c.Neg()
	}
	return __obf_c9ce1b5e622f992c
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
func (__obf_3c5564ffc2911a1b Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)  // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)  // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15) // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125)
	__obf_07c356752a4027b3 := // 4/pi

		// make argument positive
		false
	if __obf_3c5564ffc2911a1b.LessThan(NewFromFloat(0.0)) {
		__obf_3c5564ffc2911a1b = __obf_3c5564ffc2911a1b.Neg()
	}
	__obf_e4e686f2e7c57b0a := __obf_3c5564ffc2911a1b.Mul(M4PI).IntPart()
	__obf_c9ce1b5e622f992c := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_e4e686f2e7c57b0a)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_e4e686f2e7c57b0a&1 == 1 {
		__obf_e4e686f2e7c57b0a++
		__obf_c9ce1b5e622f992c = __obf_c9ce1b5e622f992c.Add(NewFromFloat(1.0))
	}
	__obf_e4e686f2e7c57b0a &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_e4e686f2e7c57b0a > 3 {
		__obf_07c356752a4027b3 = !__obf_07c356752a4027b3
		__obf_e4e686f2e7c57b0a -= 4
	}
	if __obf_e4e686f2e7c57b0a > 1 {
		__obf_07c356752a4027b3 = !__obf_07c356752a4027b3
	}
	__obf_46a4df43b872c9a9 := __obf_3c5564ffc2911a1b.Sub(__obf_c9ce1b5e622f992c.Mul(PI4A)).Sub(__obf_c9ce1b5e622f992c.Mul(PI4B)).Sub(__obf_c9ce1b5e622f992c.Mul(PI4C))
	__obf_ef73b20430ae076f := // Extended precision modular arithmetic
		__obf_46a4df43b872c9a9.Mul(__obf_46a4df43b872c9a9)

	if __obf_e4e686f2e7c57b0a == 1 || __obf_e4e686f2e7c57b0a == 2 {
		__obf_c9ce1b5e622f992c = __obf_46a4df43b872c9a9.Add(__obf_46a4df43b872c9a9.Mul(__obf_ef73b20430ae076f).Mul(_sin[0].Mul(__obf_ef73b20430ae076f).Add(_sin[1]).Mul(__obf_ef73b20430ae076f).Add(_sin[2]).Mul(__obf_ef73b20430ae076f).Add(_sin[3]).Mul(__obf_ef73b20430ae076f).Add(_sin[4]).Mul(__obf_ef73b20430ae076f).Add(_sin[5])))
	} else {
		__obf_d6ae9e541a607f97 := __obf_ef73b20430ae076f.Mul(__obf_ef73b20430ae076f).Mul(_cos[0].Mul(__obf_ef73b20430ae076f).Add(_cos[1]).Mul(__obf_ef73b20430ae076f).Add(_cos[2]).Mul(__obf_ef73b20430ae076f).Add(_cos[3]).Mul(__obf_ef73b20430ae076f).Add(_cos[4]).Mul(__obf_ef73b20430ae076f).Add(_cos[5]))
		__obf_c9ce1b5e622f992c = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_ef73b20430ae076f)).Add(__obf_d6ae9e541a607f97)
	}
	if __obf_07c356752a4027b3 {
		__obf_c9ce1b5e622f992c = __obf_c9ce1b5e622f992c.Neg()
	}
	return __obf_c9ce1b5e622f992c
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
func (__obf_3c5564ffc2911a1b Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_3c5564ffc2911a1b.Equal(NewFromFloat(0.0)) {
		return __obf_3c5564ffc2911a1b
	}
	__obf_07c356752a4027b3 := // make argument positive but save the sign
		false
	if __obf_3c5564ffc2911a1b.LessThan(NewFromFloat(0.0)) {
		__obf_3c5564ffc2911a1b = __obf_3c5564ffc2911a1b.Neg()
		__obf_07c356752a4027b3 = true
	}
	__obf_e4e686f2e7c57b0a := __obf_3c5564ffc2911a1b.Mul(M4PI).IntPart()
	__obf_c9ce1b5e622f992c := // integer part of x/(Pi/4), as integer for tests on the phase angle
		NewFromFloat(float64(__obf_e4e686f2e7c57b0a)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_e4e686f2e7c57b0a&1 == 1 {
		__obf_e4e686f2e7c57b0a++
		__obf_c9ce1b5e622f992c = __obf_c9ce1b5e622f992c.Add(NewFromFloat(1.0))
	}
	__obf_46a4df43b872c9a9 := __obf_3c5564ffc2911a1b.Sub(__obf_c9ce1b5e622f992c.Mul(PI4A)).Sub(__obf_c9ce1b5e622f992c.Mul(PI4B)).Sub(__obf_c9ce1b5e622f992c.Mul(PI4C))
	__obf_ef73b20430ae076f := // Extended precision modular arithmetic
		__obf_46a4df43b872c9a9.Mul(__obf_46a4df43b872c9a9)

	if __obf_ef73b20430ae076f.GreaterThan(NewFromFloat(1e-14)) {
		__obf_d6ae9e541a607f97 := __obf_ef73b20430ae076f.Mul(_tanP[0].Mul(__obf_ef73b20430ae076f).Add(_tanP[1]).Mul(__obf_ef73b20430ae076f).Add(_tanP[2]))
		__obf_7608654877e0d64e := __obf_ef73b20430ae076f.Add(_tanQ[1]).Mul(__obf_ef73b20430ae076f).Add(_tanQ[2]).Mul(__obf_ef73b20430ae076f).Add(_tanQ[3]).Mul(__obf_ef73b20430ae076f).Add(_tanQ[4])
		__obf_c9ce1b5e622f992c = __obf_46a4df43b872c9a9.Add(__obf_46a4df43b872c9a9.Mul(__obf_d6ae9e541a607f97.Div(__obf_7608654877e0d64e)))
	} else {
		__obf_c9ce1b5e622f992c = __obf_46a4df43b872c9a9
	}
	if __obf_e4e686f2e7c57b0a&2 == 2 {
		__obf_c9ce1b5e622f992c = NewFromFloat(-1.0).Div(__obf_c9ce1b5e622f992c)
	}
	if __obf_07c356752a4027b3 {
		__obf_c9ce1b5e622f992c = __obf_c9ce1b5e622f992c.Neg()
	}
	return __obf_c9ce1b5e622f992c
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

type __obf_ebbf7bb33b20bf6e struct {
	__obf_3c5564ffc2911a1b [800]byte
	__obf_f75b34971d3492db int  // digits, big-endian representation
	__obf_43d811ed87289907 int  // number of digits used
	__obf_d46708a8a96b3d7e bool // decimal point
	__obf_3ff6e68af28bf41f bool // negative flag
	// discarded nonzero digits beyond d[:nd]
}

func (__obf_b93edf907a7818bf *__obf_ebbf7bb33b20bf6e) String() string {
	__obf_0657fff5d9f5aeef := 10 + __obf_b93edf907a7818bf.__obf_f75b34971d3492db
	if __obf_b93edf907a7818bf.__obf_43d811ed87289907 > 0 {
		__obf_0657fff5d9f5aeef += __obf_b93edf907a7818bf.__obf_43d811ed87289907
	}
	if __obf_b93edf907a7818bf.__obf_43d811ed87289907 < 0 {
		__obf_0657fff5d9f5aeef += -__obf_b93edf907a7818bf.__obf_43d811ed87289907
	}
	__obf_f55ab7584695a22d := make([]byte, __obf_0657fff5d9f5aeef)
	__obf_d6ae9e541a607f97 := 0
	switch {
	case __obf_b93edf907a7818bf.__obf_f75b34971d3492db == 0:
		return "0"

	case __obf_b93edf907a7818bf.
		// zeros fill space between decimal point and digits
		__obf_43d811ed87289907 <= 0:
		__obf_f55ab7584695a22d[__obf_d6ae9e541a607f97] = '0'
		__obf_d6ae9e541a607f97++
		__obf_f55ab7584695a22d[__obf_d6ae9e541a607f97] = '.'
		__obf_d6ae9e541a607f97++
		__obf_d6ae9e541a607f97 += __obf_2e9be9268f327d4c(__obf_f55ab7584695a22d[__obf_d6ae9e541a607f97 : __obf_d6ae9e541a607f97+-__obf_b93edf907a7818bf.__obf_43d811ed87289907])
		__obf_d6ae9e541a607f97 += copy(__obf_f55ab7584695a22d[__obf_d6ae9e541a607f97:], __obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[0:__obf_b93edf907a7818bf.__obf_f75b34971d3492db])

	case __obf_b93edf907a7818bf.
		// decimal point in middle of digits
		__obf_43d811ed87289907 < __obf_b93edf907a7818bf.__obf_f75b34971d3492db:
		__obf_d6ae9e541a607f97 += copy(__obf_f55ab7584695a22d[__obf_d6ae9e541a607f97:], __obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[0:__obf_b93edf907a7818bf.__obf_43d811ed87289907])
		__obf_f55ab7584695a22d[__obf_d6ae9e541a607f97] = '.'
		__obf_d6ae9e541a607f97++
		__obf_d6ae9e541a607f97 += copy(__obf_f55ab7584695a22d[__obf_d6ae9e541a607f97:], __obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_b93edf907a7818bf.__obf_43d811ed87289907:

		// zeros fill space between digits and decimal point
		__obf_b93edf907a7818bf.__obf_f75b34971d3492db])

	default:
		__obf_d6ae9e541a607f97 += copy(__obf_f55ab7584695a22d[__obf_d6ae9e541a607f97:], __obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[0:__obf_b93edf907a7818bf.__obf_f75b34971d3492db])
		__obf_d6ae9e541a607f97 += __obf_2e9be9268f327d4c(__obf_f55ab7584695a22d[__obf_d6ae9e541a607f97 : __obf_d6ae9e541a607f97+__obf_b93edf907a7818bf.__obf_43d811ed87289907-__obf_b93edf907a7818bf.__obf_f75b34971d3492db])
	}
	return string(__obf_f55ab7584695a22d[0:__obf_d6ae9e541a607f97])
}

func __obf_2e9be9268f327d4c(__obf_a7a6039de2566b41 []byte) int {
	for __obf_6ee7f4c10f6c9355 := range __obf_a7a6039de2566b41 {
		__obf_a7a6039de2566b41[__obf_6ee7f4c10f6c9355] = '0'
	}
	return len(__obf_a7a6039de2566b41)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_cb9e4eca275b51a6(__obf_b93edf907a7818bf *__obf_ebbf7bb33b20bf6e) {
	for __obf_b93edf907a7818bf.__obf_f75b34971d3492db > 0 && __obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_b93edf907a7818bf.__obf_f75b34971d3492db-1] == '0' {
		__obf_b93edf907a7818bf.__obf_f75b34971d3492db--
	}
	if __obf_b93edf907a7818bf.__obf_f75b34971d3492db == 0 {
		__obf_b93edf907a7818bf.

			// Assign v to a.
			__obf_43d811ed87289907 = 0
	}
}

func (__obf_b93edf907a7818bf *__obf_ebbf7bb33b20bf6e) Assign(__obf_3785ba6483ed0c4c uint64) {
	var __obf_f55ab7584695a22d [24]byte
	__obf_0657fff5d9f5aeef := // Write reversed decimal in buf.
		0
	for __obf_3785ba6483ed0c4c > 0 {
		__obf_92487c8646abada5 := __obf_3785ba6483ed0c4c / 10
		__obf_3785ba6483ed0c4c -= 10 * __obf_92487c8646abada5
		__obf_f55ab7584695a22d[__obf_0657fff5d9f5aeef] = byte(__obf_3785ba6483ed0c4c + '0')
		__obf_0657fff5d9f5aeef++
		__obf_3785ba6483ed0c4c = __obf_92487c8646abada5
	}
	__obf_b93edf907a7818bf.

		// Reverse again to produce forward decimal in a.d.
		__obf_f75b34971d3492db = 0
	for __obf_0657fff5d9f5aeef--; __obf_0657fff5d9f5aeef >= 0; __obf_0657fff5d9f5aeef-- {
		__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_b93edf907a7818bf.__obf_f75b34971d3492db] = __obf_f55ab7584695a22d[__obf_0657fff5d9f5aeef]
		__obf_b93edf907a7818bf.__obf_f75b34971d3492db++
	}
	__obf_b93edf907a7818bf.__obf_43d811ed87289907 = __obf_b93edf907a7818bf.

		// Maximum shift that we can do in one pass without overflow.
		// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
		__obf_f75b34971d3492db
	__obf_cb9e4eca275b51a6(__obf_b93edf907a7818bf)
}

const __obf_17cdfce21997612b = 32 << (^uint(0) >> 63)
const __obf_813e8542a2969939 = __obf_17cdfce21997612b - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_a75f62e4c39cabd6(__obf_b93edf907a7818bf *__obf_ebbf7bb33b20bf6e, __obf_55924e50211870ef uint) {
	__obf_2b7b4dcf959f4b21 := 0
	__obf_d6ae9e541a607f97 := // read pointer
		0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_0657fff5d9f5aeef uint
	for ; __obf_0657fff5d9f5aeef>>__obf_55924e50211870ef == 0; __obf_2b7b4dcf959f4b21++ {
		if __obf_2b7b4dcf959f4b21 >= __obf_b93edf907a7818bf.__obf_f75b34971d3492db {
			if __obf_0657fff5d9f5aeef == 0 {
				__obf_b93edf907a7818bf.
					// a == 0; shouldn't get here, but handle anyway.
					__obf_f75b34971d3492db = 0
				return
			}
			for __obf_0657fff5d9f5aeef>>__obf_55924e50211870ef == 0 {
				__obf_0657fff5d9f5aeef = __obf_0657fff5d9f5aeef * 10
				__obf_2b7b4dcf959f4b21++
			}
			break
		}
		__obf_a1cf2ec9c989d244 := uint(__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_2b7b4dcf959f4b21])
		__obf_0657fff5d9f5aeef = __obf_0657fff5d9f5aeef*10 + __obf_a1cf2ec9c989d244 - '0'
	}
	__obf_b93edf907a7818bf.__obf_43d811ed87289907 -= __obf_2b7b4dcf959f4b21 - 1

	var __obf_61784c4923054f6a uint = (1 << __obf_55924e50211870ef) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_2b7b4dcf959f4b21 < __obf_b93edf907a7818bf.__obf_f75b34971d3492db; __obf_2b7b4dcf959f4b21++ {
		__obf_a1cf2ec9c989d244 := uint(__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_2b7b4dcf959f4b21])
		__obf_d6cfd060005c859d := __obf_0657fff5d9f5aeef >> __obf_55924e50211870ef
		__obf_0657fff5d9f5aeef &= __obf_61784c4923054f6a
		__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_d6ae9e541a607f97] = byte(__obf_d6cfd060005c859d + '0')
		__obf_d6ae9e541a607f97++
		__obf_0657fff5d9f5aeef = __obf_0657fff5d9f5aeef*10 + __obf_a1cf2ec9c989d244 - '0'
	}

	// Put down extra digits.
	for __obf_0657fff5d9f5aeef > 0 {
		__obf_d6cfd060005c859d := __obf_0657fff5d9f5aeef >> __obf_55924e50211870ef
		__obf_0657fff5d9f5aeef &= __obf_61784c4923054f6a
		if __obf_d6ae9e541a607f97 < len(__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b) {
			__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_d6ae9e541a607f97] = byte(__obf_d6cfd060005c859d + '0')
			__obf_d6ae9e541a607f97++
		} else if __obf_d6cfd060005c859d > 0 {
			__obf_b93edf907a7818bf.__obf_3ff6e68af28bf41f = true
		}
		__obf_0657fff5d9f5aeef = __obf_0657fff5d9f5aeef * 10
	}
	__obf_b93edf907a7818bf.__obf_f75b34971d3492db = __obf_d6ae9e541a607f97

	// Cheat sheet for left shift: table indexed by shift count giving
	// number of new digits that will be introduced by that shift.
	//
	// For example, leftcheats[4] = {2, "625"}.  That means that
	// if we are shifting by 4 (multiplying by 16), it will add 2 digits
	// when the string prefix is "625" through "999", and one fewer digit
	// if the string prefix is "000" through "624".
	//
	// Credit for this trick goes to Ken.
	__obf_cb9e4eca275b51a6(__obf_b93edf907a7818bf)
}

type __obf_4c6ada58ca3cc953 struct {
	__obf_8706633970def1d1 int
	__obf_b80443e8e9b7d4c9 string // number of new digits
	// minus one digit if original < a.
}

var __obf_47af59ee8bae17e9 = []__obf_4c6ada58ca3cc953{
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
func __obf_a4aa5ada2ce9b8a1(__obf_63152a35c2a3d9f4 []byte, __obf_88ca45e2afd102b9 string) bool {
	for __obf_6ee7f4c10f6c9355 := 0; __obf_6ee7f4c10f6c9355 < len(__obf_88ca45e2afd102b9); __obf_6ee7f4c10f6c9355++ {
		if __obf_6ee7f4c10f6c9355 >= len(__obf_63152a35c2a3d9f4) {
			return true
		}
		if __obf_63152a35c2a3d9f4[__obf_6ee7f4c10f6c9355] != __obf_88ca45e2afd102b9[__obf_6ee7f4c10f6c9355] {
			return __obf_63152a35c2a3d9f4[__obf_6ee7f4c10f6c9355] < __obf_88ca45e2afd102b9[__obf_6ee7f4c10f6c9355]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_14c5fb06a609cf5a(__obf_b93edf907a7818bf *__obf_ebbf7bb33b20bf6e, __obf_55924e50211870ef uint) {
	__obf_8706633970def1d1 := __obf_47af59ee8bae17e9[__obf_55924e50211870ef].__obf_8706633970def1d1
	if __obf_a4aa5ada2ce9b8a1(__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[0:__obf_b93edf907a7818bf.__obf_f75b34971d3492db], __obf_47af59ee8bae17e9[__obf_55924e50211870ef].__obf_b80443e8e9b7d4c9) {
		__obf_8706633970def1d1--
	}
	__obf_2b7b4dcf959f4b21 := __obf_b93edf907a7818bf. // read index
								__obf_f75b34971d3492db
	__obf_d6ae9e541a607f97 := __obf_b93edf907a7818bf. // write index
								__obf_f75b34971d3492db + __obf_8706633970def1d1

	// Pick up a digit, put down a digit.
	var __obf_0657fff5d9f5aeef uint
	for __obf_2b7b4dcf959f4b21--; __obf_2b7b4dcf959f4b21 >= 0; __obf_2b7b4dcf959f4b21-- {
		__obf_0657fff5d9f5aeef += (uint(__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_2b7b4dcf959f4b21]) - '0') << __obf_55924e50211870ef
		__obf_28336601dd09705e := __obf_0657fff5d9f5aeef / 10
		__obf_e3a62147921e75a7 := __obf_0657fff5d9f5aeef - 10*__obf_28336601dd09705e
		__obf_d6ae9e541a607f97--
		if __obf_d6ae9e541a607f97 < len(__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b) {
			__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_d6ae9e541a607f97] = byte(__obf_e3a62147921e75a7 + '0')
		} else if __obf_e3a62147921e75a7 != 0 {
			__obf_b93edf907a7818bf.__obf_3ff6e68af28bf41f = true
		}
		__obf_0657fff5d9f5aeef = __obf_28336601dd09705e
	}

	// Put down extra digits.
	for __obf_0657fff5d9f5aeef > 0 {
		__obf_28336601dd09705e := __obf_0657fff5d9f5aeef / 10
		__obf_e3a62147921e75a7 := __obf_0657fff5d9f5aeef - 10*__obf_28336601dd09705e
		__obf_d6ae9e541a607f97--
		if __obf_d6ae9e541a607f97 < len(__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b) {
			__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_d6ae9e541a607f97] = byte(__obf_e3a62147921e75a7 + '0')
		} else if __obf_e3a62147921e75a7 != 0 {
			__obf_b93edf907a7818bf.__obf_3ff6e68af28bf41f = true
		}
		__obf_0657fff5d9f5aeef = __obf_28336601dd09705e
	}
	__obf_b93edf907a7818bf.__obf_f75b34971d3492db += __obf_8706633970def1d1
	if __obf_b93edf907a7818bf.__obf_f75b34971d3492db >= len(__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b) {
		__obf_b93edf907a7818bf.__obf_f75b34971d3492db = len(__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b)
	}
	__obf_b93edf907a7818bf.__obf_43d811ed87289907 += __obf_8706633970def1d1

	// Binary shift left (k > 0) or right (k < 0).
	__obf_cb9e4eca275b51a6(__obf_b93edf907a7818bf)
}

func (__obf_b93edf907a7818bf *__obf_ebbf7bb33b20bf6e) Shift(__obf_55924e50211870ef int) {
	switch {
	case __obf_b93edf907a7818bf.
		// nothing to do: a == 0
		__obf_f75b34971d3492db == 0:

	case __obf_55924e50211870ef > 0:
		for __obf_55924e50211870ef > __obf_813e8542a2969939 {
			__obf_14c5fb06a609cf5a(__obf_b93edf907a7818bf, __obf_813e8542a2969939)
			__obf_55924e50211870ef -= __obf_813e8542a2969939
		}
		__obf_14c5fb06a609cf5a(__obf_b93edf907a7818bf, uint(__obf_55924e50211870ef))
	case __obf_55924e50211870ef < 0:
		for __obf_55924e50211870ef < -__obf_813e8542a2969939 {
			__obf_a75f62e4c39cabd6(__obf_b93edf907a7818bf, __obf_813e8542a2969939)
			__obf_55924e50211870ef += __obf_813e8542a2969939
		}
		__obf_a75f62e4c39cabd6(__obf_b93edf907a7818bf, uint(-__obf_55924e50211870ef))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_ac876bf9b95252b0(__obf_b93edf907a7818bf *__obf_ebbf7bb33b20bf6e, __obf_f75b34971d3492db int) bool {
	if __obf_f75b34971d3492db < 0 || __obf_f75b34971d3492db >= __obf_b93edf907a7818bf.__obf_f75b34971d3492db {
		return false
	}
	if __obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_f75b34971d3492db] == '5' && __obf_f75b34971d3492db+1 == __obf_b93edf907a7818bf. // exactly halfway - round to even
																		__obf_f75b34971d3492db {
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_b93edf907a7818bf.__obf_3ff6e68af28bf41f {
			return true
		}
		return __obf_f75b34971d3492db > 0 && (__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_f75b34971d3492db-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_b93edf907a7818bf.

		// Round a to nd digits (or fewer).
		// If nd is zero, it means we're rounding
		// just to the left of the digits, as in
		// 0.09 -> 0.1.
		__obf_3c5564ffc2911a1b[__obf_f75b34971d3492db] >= '5'
}

func (__obf_b93edf907a7818bf *__obf_ebbf7bb33b20bf6e) Round(__obf_f75b34971d3492db int) {
	if __obf_f75b34971d3492db < 0 || __obf_f75b34971d3492db >= __obf_b93edf907a7818bf.__obf_f75b34971d3492db {
		return
	}
	if __obf_ac876bf9b95252b0(__obf_b93edf907a7818bf, __obf_f75b34971d3492db) {
		__obf_b93edf907a7818bf.
			RoundUp(__obf_f75b34971d3492db)
	} else {
		__obf_b93edf907a7818bf.
			RoundDown(__obf_f75b34971d3492db)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_b93edf907a7818bf *__obf_ebbf7bb33b20bf6e) RoundDown(__obf_f75b34971d3492db int) {
	if __obf_f75b34971d3492db < 0 || __obf_f75b34971d3492db >= __obf_b93edf907a7818bf.__obf_f75b34971d3492db {
		return
	}
	__obf_b93edf907a7818bf.__obf_f75b34971d3492db = __obf_f75b34971d3492db

	// Round a up to nd digits (or fewer).
	__obf_cb9e4eca275b51a6(__obf_b93edf907a7818bf)
}

func (__obf_b93edf907a7818bf *__obf_ebbf7bb33b20bf6e) RoundUp(__obf_f75b34971d3492db int) {
	if __obf_f75b34971d3492db < 0 || __obf_f75b34971d3492db >= __obf_b93edf907a7818bf.

		// round up
		__obf_f75b34971d3492db {
		return
	}

	for __obf_6ee7f4c10f6c9355 := __obf_f75b34971d3492db - 1; __obf_6ee7f4c10f6c9355 >= 0; __obf_6ee7f4c10f6c9355-- {
		__obf_a1cf2ec9c989d244 := __obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_6ee7f4c10f6c9355]
		if __obf_a1cf2ec9c989d244 < '9' {
			__obf_b93edf907a7818bf. // can stop after this digit
						__obf_3c5564ffc2911a1b[__obf_6ee7f4c10f6c9355]++
			__obf_b93edf907a7818bf.__obf_f75b34971d3492db = __obf_6ee7f4c10f6c9355 + 1
			return
		}
	}
	__obf_b93edf907a7818bf.

		// Number is all 9s.
		// Change to single 1 with adjusted decimal point.
		__obf_3c5564ffc2911a1b[0] = '1'
	__obf_b93edf907a7818bf.

		// Extract integer part, rounded appropriately.
		// No guarantees about overflow.
		__obf_f75b34971d3492db = 1
	__obf_b93edf907a7818bf.__obf_43d811ed87289907++
}

func (__obf_b93edf907a7818bf *__obf_ebbf7bb33b20bf6e) RoundedInteger() uint64 {
	if __obf_b93edf907a7818bf.__obf_43d811ed87289907 > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_6ee7f4c10f6c9355 int
	__obf_0657fff5d9f5aeef := uint64(0)
	for __obf_6ee7f4c10f6c9355 = 0; __obf_6ee7f4c10f6c9355 < __obf_b93edf907a7818bf.__obf_43d811ed87289907 && __obf_6ee7f4c10f6c9355 < __obf_b93edf907a7818bf.__obf_f75b34971d3492db; __obf_6ee7f4c10f6c9355++ {
		__obf_0657fff5d9f5aeef = __obf_0657fff5d9f5aeef*10 + uint64(__obf_b93edf907a7818bf.__obf_3c5564ffc2911a1b[__obf_6ee7f4c10f6c9355]-'0')
	}
	for ; __obf_6ee7f4c10f6c9355 < __obf_b93edf907a7818bf.__obf_43d811ed87289907; __obf_6ee7f4c10f6c9355++ {
		__obf_0657fff5d9f5aeef *= 10
	}
	if __obf_ac876bf9b95252b0(__obf_b93edf907a7818bf, __obf_b93edf907a7818bf.__obf_43d811ed87289907) {
		__obf_0657fff5d9f5aeef++
	}
	return __obf_0657fff5d9f5aeef
}
