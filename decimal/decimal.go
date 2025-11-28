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
package __obf_ae16adf734cfe1aa

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

var __obf_91bffcce19efb7d8 = big.NewInt(0)
var __obf_6b7012530e205f39 = big.NewInt(1)
var __obf_3338f87aa2a5f65a = big.NewInt(2)
var __obf_7a3c2c07b99ffb48 = big.NewInt(4)
var __obf_a02c1e3273085f0d = big.NewInt(5)
var __obf_6e3308650b479548 = big.NewInt(10)
var __obf_cc9bfe65d22c82e4 = big.NewInt(20)

var __obf_5004c7371a2d0179 = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	__obf_6e5f77decbc21dea *big.Int

	// NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.
	__obf_1ef84065a6a49f43 int32
}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(__obf_6e5f77decbc21dea int64, __obf_1ef84065a6a49f43 int32) Decimal {
	return Decimal{
		__obf_6e5f77decbc21dea: big.NewInt(__obf_6e5f77decbc21dea),
		__obf_1ef84065a6a49f43: __obf_1ef84065a6a49f43,
	}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(__obf_6e5f77decbc21dea int64) Decimal {
	return Decimal{
		__obf_6e5f77decbc21dea: big.NewInt(__obf_6e5f77decbc21dea),
		__obf_1ef84065a6a49f43: 0,
	}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(__obf_6e5f77decbc21dea int32) Decimal {
	return Decimal{
		__obf_6e5f77decbc21dea: big.NewInt(int64(__obf_6e5f77decbc21dea)),
		__obf_1ef84065a6a49f43: 0,
	}
}

// NewFromUint64 converts an uint64 to Decimal.
//
// Example:
//
//	NewFromUint64(123).String() // output: "123"
func NewFromUint64(__obf_6e5f77decbc21dea uint64) Decimal {
	return Decimal{
		__obf_6e5f77decbc21dea: new(big.Int).SetUint64(__obf_6e5f77decbc21dea),
		__obf_1ef84065a6a49f43: 0,
	}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(__obf_6e5f77decbc21dea *big.Int, __obf_1ef84065a6a49f43 int32) Decimal {
	return Decimal{
		__obf_6e5f77decbc21dea: new(big.Int).Set(__obf_6e5f77decbc21dea),
		__obf_1ef84065a6a49f43: __obf_1ef84065a6a49f43,
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
func NewFromBigRat(__obf_6e5f77decbc21dea *big.Rat, __obf_fb38b245d2cdb3ee int32) Decimal {
	return Decimal{
		__obf_6e5f77decbc21dea: new(big.Int).Set(__obf_6e5f77decbc21dea.Num()),
		__obf_1ef84065a6a49f43: 0,
	}.DivRound(Decimal{
		__obf_6e5f77decbc21dea: new(big.Int).Set(__obf_6e5f77decbc21dea.Denom()),
		__obf_1ef84065a6a49f43: 0,
	}, __obf_fb38b245d2cdb3ee)
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(__obf_6e5f77decbc21dea string) (Decimal, error) {
	__obf_cc7cb376e25aa332 := __obf_6e5f77decbc21dea
	var __obf_0ba2a9c4efd07969 string
	var __obf_1ef84065a6a49f43 int64

	// Check if number is using scientific notation
	__obf_d0e3791be54e7bad := strings.IndexAny(__obf_6e5f77decbc21dea, "Ee")
	if __obf_d0e3791be54e7bad != -1 {
		__obf_123517f3c1fa834b, __obf_68bcf986f9cdf0c0 := strconv.ParseInt(__obf_6e5f77decbc21dea[__obf_d0e3791be54e7bad+1:], 10, 32)
		if __obf_68bcf986f9cdf0c0 != nil {
			if __obf_c1ef4fea8d3ecd63, __obf_1ef0d34491e1175a := __obf_68bcf986f9cdf0c0.(*strconv.NumError); __obf_1ef0d34491e1175a && __obf_c1ef4fea8d3ecd63.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_6e5f77decbc21dea)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", __obf_6e5f77decbc21dea)
		}
		__obf_6e5f77decbc21dea = __obf_6e5f77decbc21dea[:__obf_d0e3791be54e7bad]
		__obf_1ef84065a6a49f43 = __obf_123517f3c1fa834b
	}

	__obf_5ff66ef4e01e4d48 := -1
	__obf_73d9032530b4dc24 := len(__obf_6e5f77decbc21dea)
	for __obf_40720372b6d983c7 := 0; __obf_40720372b6d983c7 < __obf_73d9032530b4dc24; __obf_40720372b6d983c7++ {
		if __obf_6e5f77decbc21dea[__obf_40720372b6d983c7] == '.' {
			if __obf_5ff66ef4e01e4d48 > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", __obf_6e5f77decbc21dea)
			}
			__obf_5ff66ef4e01e4d48 = __obf_40720372b6d983c7
		}
	}

	if __obf_5ff66ef4e01e4d48 == -1 {
		// There is no decimal point, we can just parse the original string as
		// an int
		__obf_0ba2a9c4efd07969 = __obf_6e5f77decbc21dea
	} else {
		if __obf_5ff66ef4e01e4d48+1 < __obf_73d9032530b4dc24 {
			__obf_0ba2a9c4efd07969 = __obf_6e5f77decbc21dea[:__obf_5ff66ef4e01e4d48] + __obf_6e5f77decbc21dea[__obf_5ff66ef4e01e4d48+1:]
		} else {
			__obf_0ba2a9c4efd07969 = __obf_6e5f77decbc21dea[:__obf_5ff66ef4e01e4d48]
		}
		__obf_123517f3c1fa834b := -len(__obf_6e5f77decbc21dea[__obf_5ff66ef4e01e4d48+1:])
		__obf_1ef84065a6a49f43 += int64(__obf_123517f3c1fa834b)
	}

	var __obf_75dd2f5b4248b696 *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(__obf_0ba2a9c4efd07969) <= 18 {
		__obf_6138b75a3c5ef6ec, __obf_68bcf986f9cdf0c0 := strconv.ParseInt(__obf_0ba2a9c4efd07969, 10, 64)
		if __obf_68bcf986f9cdf0c0 != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_6e5f77decbc21dea)
		}
		__obf_75dd2f5b4248b696 = big.NewInt(__obf_6138b75a3c5ef6ec)
	} else {
		__obf_75dd2f5b4248b696 = new(big.Int)
		_, __obf_1ef0d34491e1175a := __obf_75dd2f5b4248b696.SetString(__obf_0ba2a9c4efd07969, 10)
		if !__obf_1ef0d34491e1175a {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", __obf_6e5f77decbc21dea)
		}
	}

	if __obf_1ef84065a6a49f43 < math.MinInt32 || __obf_1ef84065a6a49f43 > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", __obf_cc7cb376e25aa332)
	}

	return Decimal{
		__obf_6e5f77decbc21dea: __obf_75dd2f5b4248b696,
		__obf_1ef84065a6a49f43: int32(__obf_1ef84065a6a49f43),
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
func NewFromFormattedString(__obf_6e5f77decbc21dea string, __obf_4de269304e9eb3c6 *regexp.Regexp) (Decimal, error) {
	__obf_a38f167b5062b0c0 := __obf_4de269304e9eb3c6.ReplaceAllString(__obf_6e5f77decbc21dea, "")
	__obf_19a369452cd5fbde, __obf_68bcf986f9cdf0c0 := NewFromString(__obf_a38f167b5062b0c0)
	if __obf_68bcf986f9cdf0c0 != nil {
		return Decimal{}, __obf_68bcf986f9cdf0c0
	}
	return __obf_19a369452cd5fbde, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(__obf_6e5f77decbc21dea string) Decimal {
	__obf_e8ecc3b172d7ce84, __obf_68bcf986f9cdf0c0 := NewFromString(__obf_6e5f77decbc21dea)
	if __obf_68bcf986f9cdf0c0 != nil {
		panic(__obf_68bcf986f9cdf0c0)
	}
	return __obf_e8ecc3b172d7ce84
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
func NewFromFloat(__obf_6e5f77decbc21dea float64) Decimal {
	if __obf_6e5f77decbc21dea == 0 {
		return New(0, 0)
	}
	return __obf_b25cfa86d7c21099(__obf_6e5f77decbc21dea, math.Float64bits(__obf_6e5f77decbc21dea), &__obf_3201299bc3256294)
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
func NewFromFloat32(__obf_6e5f77decbc21dea float32) Decimal {
	if __obf_6e5f77decbc21dea == 0 {
		return New(0, 0)
	}
	// XOR is workaround for https://github.com/golang/go/issues/26285
	__obf_364e38cb96f1c266 := math.Float32bits(__obf_6e5f77decbc21dea) ^ 0x80808080
	return __obf_b25cfa86d7c21099(float64(__obf_6e5f77decbc21dea), uint64(__obf_364e38cb96f1c266)^0x80808080, &__obf_a47ee567a13f3274)
}

func __obf_b25cfa86d7c21099(__obf_3fcee5abdf286e44 float64, __obf_ca6a2e6cd90372b0 uint64, __obf_f9f981b2fd2cf5a1 *__obf_94eeb44634f64ac2) Decimal {
	if math.IsNaN(__obf_3fcee5abdf286e44) || math.IsInf(__obf_3fcee5abdf286e44, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_3fcee5abdf286e44))
	}
	__obf_1ef84065a6a49f43 := int(__obf_ca6a2e6cd90372b0>>__obf_f9f981b2fd2cf5a1.__obf_5991f9c502ad3b8d) & (1<<__obf_f9f981b2fd2cf5a1.__obf_7361f35f98ac5740 - 1)
	__obf_6a3e060d23caf99c := __obf_ca6a2e6cd90372b0 & (uint64(1)<<__obf_f9f981b2fd2cf5a1.__obf_5991f9c502ad3b8d - 1)

	switch __obf_1ef84065a6a49f43 {
	case 0:
		// denormalized
		__obf_1ef84065a6a49f43++

	default:
		// add implicit top bit
		__obf_6a3e060d23caf99c |= uint64(1) << __obf_f9f981b2fd2cf5a1.__obf_5991f9c502ad3b8d
	}
	__obf_1ef84065a6a49f43 += __obf_f9f981b2fd2cf5a1.__obf_bd6660fbe7220f81

	var __obf_19a369452cd5fbde __obf_ae16adf734cfe1aa
	__obf_19a369452cd5fbde.Assign(__obf_6a3e060d23caf99c)
	__obf_19a369452cd5fbde.Shift(__obf_1ef84065a6a49f43 - int(__obf_f9f981b2fd2cf5a1.__obf_5991f9c502ad3b8d))
	__obf_19a369452cd5fbde.__obf_d12c1496d23e3503 = __obf_ca6a2e6cd90372b0>>(__obf_f9f981b2fd2cf5a1.__obf_7361f35f98ac5740+__obf_f9f981b2fd2cf5a1.__obf_5991f9c502ad3b8d) != 0

	__obf_a9030787063799d3(&__obf_19a369452cd5fbde, __obf_6a3e060d23caf99c, __obf_1ef84065a6a49f43, __obf_f9f981b2fd2cf5a1)
	// If less than 19 digits, we can do calculation in an int64.
	if __obf_19a369452cd5fbde.__obf_04074ea54302ce17 < 19 {
		__obf_e059bdc33d645894 := int64(0)
		__obf_3f873ba4f9076041 := int64(1)
		for __obf_40720372b6d983c7 := __obf_19a369452cd5fbde.__obf_04074ea54302ce17 - 1; __obf_40720372b6d983c7 >= 0; __obf_40720372b6d983c7-- {
			__obf_e059bdc33d645894 += __obf_3f873ba4f9076041 * int64(__obf_19a369452cd5fbde.__obf_19a369452cd5fbde[__obf_40720372b6d983c7]-'0')
			__obf_3f873ba4f9076041 *= 10
		}
		if __obf_19a369452cd5fbde.__obf_d12c1496d23e3503 {
			__obf_e059bdc33d645894 *= -1
		}
		return Decimal{__obf_6e5f77decbc21dea: big.NewInt(__obf_e059bdc33d645894), __obf_1ef84065a6a49f43: int32(__obf_19a369452cd5fbde.__obf_45565779317643f9) - int32(__obf_19a369452cd5fbde.__obf_04074ea54302ce17)}
	}
	__obf_75dd2f5b4248b696 := new(big.Int)
	__obf_75dd2f5b4248b696, __obf_1ef0d34491e1175a := __obf_75dd2f5b4248b696.SetString(string(__obf_19a369452cd5fbde.__obf_19a369452cd5fbde[:__obf_19a369452cd5fbde.__obf_04074ea54302ce17]), 10)
	if __obf_1ef0d34491e1175a {
		return Decimal{__obf_6e5f77decbc21dea: __obf_75dd2f5b4248b696, __obf_1ef84065a6a49f43: int32(__obf_19a369452cd5fbde.__obf_45565779317643f9) - int32(__obf_19a369452cd5fbde.__obf_04074ea54302ce17)}
	}

	return NewFromFloatWithExponent(__obf_3fcee5abdf286e44, int32(__obf_19a369452cd5fbde.__obf_45565779317643f9)-int32(__obf_19a369452cd5fbde.__obf_04074ea54302ce17))
}

// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
// number of fractional digits.
//
// Example:
//
//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
func NewFromFloatWithExponent(__obf_6e5f77decbc21dea float64, __obf_1ef84065a6a49f43 int32) Decimal {
	if math.IsNaN(__obf_6e5f77decbc21dea) || math.IsInf(__obf_6e5f77decbc21dea, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", __obf_6e5f77decbc21dea))
	}

	__obf_ca6a2e6cd90372b0 := math.Float64bits(__obf_6e5f77decbc21dea)
	__obf_6a3e060d23caf99c := __obf_ca6a2e6cd90372b0 & (1<<52 - 1)
	__obf_eb56d6ae62080b35 := int32((__obf_ca6a2e6cd90372b0 >> 52) & (1<<11 - 1))
	__obf_0a87f1adff24c937 := __obf_ca6a2e6cd90372b0 >> 63

	if __obf_eb56d6ae62080b35 == 0 {
		// specials
		if __obf_6a3e060d23caf99c == 0 {
			return Decimal{}
		}
		// subnormal
		__obf_eb56d6ae62080b35++
	} else {
		// normal
		__obf_6a3e060d23caf99c |= 1 << 52
	}

	__obf_eb56d6ae62080b35 -= 1023 + 52

	// normalizing base-2 values
	for __obf_6a3e060d23caf99c&1 == 0 {
		__obf_6a3e060d23caf99c = __obf_6a3e060d23caf99c >> 1
		__obf_eb56d6ae62080b35++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if __obf_1ef84065a6a49f43 < 0 && __obf_1ef84065a6a49f43 < __obf_eb56d6ae62080b35 {
		if __obf_eb56d6ae62080b35 < 0 {
			__obf_1ef84065a6a49f43 = __obf_eb56d6ae62080b35
		} else {
			__obf_1ef84065a6a49f43 = 0
		}
	}

	// representing 10^M * 2^N as 5^M * 2^(M+N)
	__obf_eb56d6ae62080b35 -= __obf_1ef84065a6a49f43

	__obf_78c292e95bf52fd1 := big.NewInt(1)
	__obf_ce93d52d1f0f5002 := big.NewInt(int64(__obf_6a3e060d23caf99c))

	// applying 5^M
	if __obf_1ef84065a6a49f43 > 0 {
		__obf_78c292e95bf52fd1 = __obf_78c292e95bf52fd1.SetInt64(int64(__obf_1ef84065a6a49f43))
		__obf_78c292e95bf52fd1 = __obf_78c292e95bf52fd1.Exp(__obf_a02c1e3273085f0d, __obf_78c292e95bf52fd1, nil)
	} else if __obf_1ef84065a6a49f43 < 0 {
		__obf_78c292e95bf52fd1 = __obf_78c292e95bf52fd1.SetInt64(-int64(__obf_1ef84065a6a49f43))
		__obf_78c292e95bf52fd1 = __obf_78c292e95bf52fd1.Exp(__obf_a02c1e3273085f0d, __obf_78c292e95bf52fd1, nil)
		__obf_ce93d52d1f0f5002 = __obf_ce93d52d1f0f5002.Mul(__obf_ce93d52d1f0f5002, __obf_78c292e95bf52fd1)
		__obf_78c292e95bf52fd1 = __obf_78c292e95bf52fd1.SetUint64(1)
	}

	// applying 2^(M+N)
	if __obf_eb56d6ae62080b35 > 0 {
		__obf_ce93d52d1f0f5002 = __obf_ce93d52d1f0f5002.Lsh(__obf_ce93d52d1f0f5002, uint(__obf_eb56d6ae62080b35))
	} else if __obf_eb56d6ae62080b35 < 0 {
		__obf_78c292e95bf52fd1 = __obf_78c292e95bf52fd1.Lsh(__obf_78c292e95bf52fd1, uint(-__obf_eb56d6ae62080b35))
	}

	// rounding and downscaling
	if __obf_1ef84065a6a49f43 > 0 || __obf_eb56d6ae62080b35 < 0 {
		__obf_938e9a4f95b74e87 := new(big.Int).Rsh(__obf_78c292e95bf52fd1, 1)
		__obf_ce93d52d1f0f5002 = __obf_ce93d52d1f0f5002.Add(__obf_ce93d52d1f0f5002, __obf_938e9a4f95b74e87)
		__obf_ce93d52d1f0f5002 = __obf_ce93d52d1f0f5002.Quo(__obf_ce93d52d1f0f5002, __obf_78c292e95bf52fd1)
	}

	if __obf_0a87f1adff24c937 == 1 {
		__obf_ce93d52d1f0f5002 = __obf_ce93d52d1f0f5002.Neg(__obf_ce93d52d1f0f5002)
	}

	return Decimal{
		__obf_6e5f77decbc21dea: __obf_ce93d52d1f0f5002,
		__obf_1ef84065a6a49f43: __obf_1ef84065a6a49f43,
	}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (__obf_19a369452cd5fbde Decimal) Copy() Decimal {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()
	return Decimal{
		__obf_6e5f77decbc21dea: new(big.Int).Set(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea),
		__obf_1ef84065a6a49f43: __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43,
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
func (__obf_19a369452cd5fbde Decimal) __obf_957a7b5240163386(__obf_1ef84065a6a49f43 int32) Decimal {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()

	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 == __obf_1ef84065a6a49f43 {
		return Decimal{
			new(big.Int).Set(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea),
			__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43,
		}
	}

	// NOTE(vadim): must convert exps to float64 before - to prevent overflow
	__obf_19b4ab1fa58f1937 := math.Abs(float64(__obf_1ef84065a6a49f43) - float64(__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43))
	__obf_6e5f77decbc21dea := new(big.Int).Set(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea)

	__obf_c469517f5dd0980f := new(big.Int).Exp(__obf_6e3308650b479548, big.NewInt(int64(__obf_19b4ab1fa58f1937)), nil)
	if __obf_1ef84065a6a49f43 > __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 {
		__obf_6e5f77decbc21dea = __obf_6e5f77decbc21dea.Quo(__obf_6e5f77decbc21dea, __obf_c469517f5dd0980f)
	} else if __obf_1ef84065a6a49f43 < __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 {
		__obf_6e5f77decbc21dea = __obf_6e5f77decbc21dea.Mul(__obf_6e5f77decbc21dea, __obf_c469517f5dd0980f)
	}

	return Decimal{
		__obf_6e5f77decbc21dea: __obf_6e5f77decbc21dea,
		__obf_1ef84065a6a49f43: __obf_1ef84065a6a49f43,
	}
}

// Abs returns the absolute value of the decimal.
func (__obf_19a369452cd5fbde Decimal) Abs() Decimal {
	if !__obf_19a369452cd5fbde.IsNegative() {
		return __obf_19a369452cd5fbde
	}
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()
	__obf_dcd405a8970df614 := new(big.Int).Abs(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea)
	return Decimal{
		__obf_6e5f77decbc21dea: __obf_dcd405a8970df614,
		__obf_1ef84065a6a49f43: __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43,
	}
}

// Add returns d + d2.
func (__obf_19a369452cd5fbde Decimal) Add(__obf_e5610ee7ba67ce83 Decimal) Decimal {
	__obf_e569aa78d09dc408, __obf_f8f48e5992257be1 := RescalePair(__obf_19a369452cd5fbde, __obf_e5610ee7ba67ce83)

	__obf_84a93c33bbb4a4c7 := new(big.Int).Add(__obf_e569aa78d09dc408.__obf_6e5f77decbc21dea, __obf_f8f48e5992257be1.__obf_6e5f77decbc21dea)
	return Decimal{
		__obf_6e5f77decbc21dea: __obf_84a93c33bbb4a4c7,
		__obf_1ef84065a6a49f43: __obf_e569aa78d09dc408.__obf_1ef84065a6a49f43,
	}
}

// Sub returns d - d2.
func (__obf_19a369452cd5fbde Decimal) Sub(__obf_e5610ee7ba67ce83 Decimal) Decimal {
	__obf_e569aa78d09dc408, __obf_f8f48e5992257be1 := RescalePair(__obf_19a369452cd5fbde, __obf_e5610ee7ba67ce83)

	__obf_84a93c33bbb4a4c7 := new(big.Int).Sub(__obf_e569aa78d09dc408.__obf_6e5f77decbc21dea, __obf_f8f48e5992257be1.__obf_6e5f77decbc21dea)
	return Decimal{
		__obf_6e5f77decbc21dea: __obf_84a93c33bbb4a4c7,
		__obf_1ef84065a6a49f43: __obf_e569aa78d09dc408.__obf_1ef84065a6a49f43,
	}
}

// Neg returns -d.
func (__obf_19a369452cd5fbde Decimal) Neg() Decimal {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()
	__obf_3fcee5abdf286e44 := new(big.Int).Neg(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea)
	return Decimal{
		__obf_6e5f77decbc21dea: __obf_3fcee5abdf286e44,
		__obf_1ef84065a6a49f43: __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43,
	}
}

// Mul returns d * d2.
func (__obf_19a369452cd5fbde Decimal) Mul(__obf_e5610ee7ba67ce83 Decimal) Decimal {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()
	__obf_e5610ee7ba67ce83.__obf_b1325bfc63b249a7()

	__obf_2771e4da350786fa := int64(__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43) + int64(__obf_e5610ee7ba67ce83.__obf_1ef84065a6a49f43)
	if __obf_2771e4da350786fa > math.MaxInt32 || __obf_2771e4da350786fa < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", __obf_2771e4da350786fa))
	}

	__obf_84a93c33bbb4a4c7 := new(big.Int).Mul(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea, __obf_e5610ee7ba67ce83.__obf_6e5f77decbc21dea)
	return Decimal{
		__obf_6e5f77decbc21dea: __obf_84a93c33bbb4a4c7,
		__obf_1ef84065a6a49f43: int32(__obf_2771e4da350786fa),
	}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (__obf_19a369452cd5fbde Decimal) Shift(__obf_ebae5be0d75c387b int32) Decimal {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()
	return Decimal{
		__obf_6e5f77decbc21dea: new(big.Int).Set(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea),
		__obf_1ef84065a6a49f43: __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 + __obf_ebae5be0d75c387b,
	}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (__obf_19a369452cd5fbde Decimal) Div(__obf_e5610ee7ba67ce83 Decimal) Decimal {
	return __obf_19a369452cd5fbde.DivRound(__obf_e5610ee7ba67ce83, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (__obf_19a369452cd5fbde Decimal) QuoRem(__obf_e5610ee7ba67ce83 Decimal, __obf_fb38b245d2cdb3ee int32) (Decimal, Decimal) {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()
	__obf_e5610ee7ba67ce83.__obf_b1325bfc63b249a7()
	if __obf_e5610ee7ba67ce83.__obf_6e5f77decbc21dea.Sign() == 0 {
		panic("decimal division by 0")
	}
	__obf_f24b6cbe98a42168 := -__obf_fb38b245d2cdb3ee
	__obf_c1ef4fea8d3ecd63 := int64(__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43) - int64(__obf_e5610ee7ba67ce83.__obf_1ef84065a6a49f43) - int64(__obf_f24b6cbe98a42168)
	if __obf_c1ef4fea8d3ecd63 > math.MaxInt32 || __obf_c1ef4fea8d3ecd63 < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var __obf_7f652b2fe88ad74e, __obf_9ab697160cd0e872, __obf_50708d4010a00d9a big.Int
	var __obf_16e479196b43af13 int32
	// d = a 10^ea
	// d2 = b 10^eb
	if __obf_c1ef4fea8d3ecd63 < 0 {
		__obf_7f652b2fe88ad74e = *__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea
		__obf_50708d4010a00d9a.SetInt64(-__obf_c1ef4fea8d3ecd63)
		__obf_9ab697160cd0e872.Exp(__obf_6e3308650b479548, &__obf_50708d4010a00d9a, nil)
		__obf_9ab697160cd0e872.Mul(__obf_e5610ee7ba67ce83.__obf_6e5f77decbc21dea, &__obf_9ab697160cd0e872)
		__obf_16e479196b43af13 = __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43
		// now aa = a
		//     bb = b 10^(scale + eb - ea)
	} else {
		__obf_50708d4010a00d9a.SetInt64(__obf_c1ef4fea8d3ecd63)
		__obf_7f652b2fe88ad74e.Exp(__obf_6e3308650b479548, &__obf_50708d4010a00d9a, nil)
		__obf_7f652b2fe88ad74e.Mul(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea, &__obf_7f652b2fe88ad74e)
		__obf_9ab697160cd0e872 = *__obf_e5610ee7ba67ce83.__obf_6e5f77decbc21dea
		__obf_16e479196b43af13 = __obf_f24b6cbe98a42168 + __obf_e5610ee7ba67ce83.__obf_1ef84065a6a49f43
		// now aa = a ^ (ea - eb - scale)
		//     bb = b
	}
	var __obf_99a04cdb50494bb6, __obf_0feccac4ef34d0a4 big.Int
	__obf_99a04cdb50494bb6.QuoRem(&__obf_7f652b2fe88ad74e, &__obf_9ab697160cd0e872, &__obf_0feccac4ef34d0a4)
	__obf_20068c1e4acf6fcb := Decimal{__obf_6e5f77decbc21dea: &__obf_99a04cdb50494bb6, __obf_1ef84065a6a49f43: __obf_f24b6cbe98a42168}
	__obf_51b99efbe12bffee := Decimal{__obf_6e5f77decbc21dea: &__obf_0feccac4ef34d0a4, __obf_1ef84065a6a49f43: __obf_16e479196b43af13}
	return __obf_20068c1e4acf6fcb, __obf_51b99efbe12bffee
}

// DivRound divides and rounds to a given precision
// i.e. to an integer multiple of 10^(-precision)
//
//	for a positive quotient digit 5 is rounded up, away from 0
//	if the quotient is negative then digit 5 is rounded down, away from 0
//
// Note that precision<0 is allowed as input.
func (__obf_19a369452cd5fbde Decimal) DivRound(__obf_e5610ee7ba67ce83 Decimal, __obf_fb38b245d2cdb3ee int32) Decimal {
	// QuoRem already checks initialization
	__obf_99a04cdb50494bb6, __obf_0feccac4ef34d0a4 := __obf_19a369452cd5fbde.QuoRem(__obf_e5610ee7ba67ce83, __obf_fb38b245d2cdb3ee)
	// the actual rounding decision is based on comparing r*10^precision and d2/2
	// instead compare 2 r 10 ^precision and d2
	var __obf_aabcdbb90e686acd big.Int
	__obf_aabcdbb90e686acd.Abs(__obf_0feccac4ef34d0a4.__obf_6e5f77decbc21dea)
	__obf_aabcdbb90e686acd.Lsh(&__obf_aabcdbb90e686acd, 1)
	// now rv2 = abs(r.value) * 2
	__obf_2d6f459f7f9d2700 := Decimal{__obf_6e5f77decbc21dea: &__obf_aabcdbb90e686acd, __obf_1ef84065a6a49f43: __obf_0feccac4ef34d0a4.__obf_1ef84065a6a49f43 + __obf_fb38b245d2cdb3ee}
	// r2 is now 2 * r * 10 ^ precision
	var __obf_6dfba1dc52bc2b42 = __obf_2d6f459f7f9d2700.Cmp(__obf_e5610ee7ba67ce83.Abs())

	if __obf_6dfba1dc52bc2b42 < 0 {
		return __obf_99a04cdb50494bb6
	}

	if __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.Sign()*__obf_e5610ee7ba67ce83.__obf_6e5f77decbc21dea.Sign() < 0 {
		return __obf_99a04cdb50494bb6.Sub(New(1, -__obf_fb38b245d2cdb3ee))
	}

	return __obf_99a04cdb50494bb6.Add(New(1, -__obf_fb38b245d2cdb3ee))
}

// Mod returns d % d2.
func (__obf_19a369452cd5fbde Decimal) Mod(__obf_e5610ee7ba67ce83 Decimal) Decimal {
	_, __obf_0feccac4ef34d0a4 := __obf_19a369452cd5fbde.QuoRem(__obf_e5610ee7ba67ce83, 0)
	return __obf_0feccac4ef34d0a4
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
func (__obf_19a369452cd5fbde Decimal) Pow(__obf_e5610ee7ba67ce83 Decimal) Decimal {
	__obf_e2d91f0337623bec := __obf_19a369452cd5fbde.Sign()
	__obf_1b2d52d258f49740 := __obf_e5610ee7ba67ce83.Sign()

	if __obf_e2d91f0337623bec == 0 {
		if __obf_1b2d52d258f49740 == 0 {
			return Decimal{}
		}
		if __obf_1b2d52d258f49740 == 1 {
			return Decimal{__obf_91bffcce19efb7d8, 0}
		}
		if __obf_1b2d52d258f49740 == -1 {
			return Decimal{}
		}
	}

	if __obf_1b2d52d258f49740 == 0 {
		return Decimal{__obf_6b7012530e205f39, 0}
	}

	// TODO: optimize extraction of fractional part
	__obf_3b803b95d5fc8b6b := Decimal{__obf_6b7012530e205f39, 0}
	__obf_a15fe04a384142dd, __obf_f3a8aaed56831f9d := __obf_e5610ee7ba67ce83.QuoRem(__obf_3b803b95d5fc8b6b, 0)

	if __obf_e2d91f0337623bec == -1 && !__obf_f3a8aaed56831f9d.IsZero() {
		return Decimal{}
	}

	__obf_09d2bbe0858a15c0, _ := __obf_19a369452cd5fbde.PowBigInt(__obf_a15fe04a384142dd.__obf_6e5f77decbc21dea)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_f3a8aaed56831f9d.__obf_6e5f77decbc21dea.Sign() == 0 {
		return __obf_09d2bbe0858a15c0
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_49d8c84b3237ddf4 := __obf_19a369452cd5fbde.NumDigits()
	__obf_388de7c089524859 := __obf_e5610ee7ba67ce83.NumDigits()

	__obf_fb38b245d2cdb3ee := __obf_49d8c84b3237ddf4

	if __obf_388de7c089524859 > __obf_fb38b245d2cdb3ee {
		__obf_fb38b245d2cdb3ee += __obf_388de7c089524859
	}

	__obf_fb38b245d2cdb3ee += 6

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_7ca7f3589705fc1e, __obf_68bcf986f9cdf0c0 := __obf_19a369452cd5fbde.Abs().Ln(-__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 + int32(__obf_fb38b245d2cdb3ee))
	if __obf_68bcf986f9cdf0c0 != nil {
		return Decimal{}
	}

	__obf_7ca7f3589705fc1e = __obf_7ca7f3589705fc1e.Mul(__obf_f3a8aaed56831f9d)

	__obf_7ca7f3589705fc1e, __obf_68bcf986f9cdf0c0 = __obf_7ca7f3589705fc1e.ExpTaylor(-__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 + int32(__obf_fb38b245d2cdb3ee))
	if __obf_68bcf986f9cdf0c0 != nil {
		return Decimal{}
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_a01a49266a86f271 := __obf_09d2bbe0858a15c0.Mul(__obf_7ca7f3589705fc1e)

	return __obf_a01a49266a86f271
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
func (__obf_19a369452cd5fbde Decimal) PowWithPrecision(__obf_e5610ee7ba67ce83 Decimal, __obf_fb38b245d2cdb3ee int32) (Decimal, error) {
	__obf_e2d91f0337623bec := __obf_19a369452cd5fbde.Sign()
	__obf_1b2d52d258f49740 := __obf_e5610ee7ba67ce83.Sign()

	if __obf_e2d91f0337623bec == 0 {
		if __obf_1b2d52d258f49740 == 0 {
			return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
		}
		if __obf_1b2d52d258f49740 == 1 {
			return Decimal{__obf_91bffcce19efb7d8, 0}, nil
		}
		if __obf_1b2d52d258f49740 == -1 {
			return Decimal{}, fmt.Errorf("cannot represent infinity value of 0 ** y, where y < 0")
		}
	}

	if __obf_1b2d52d258f49740 == 0 {
		return Decimal{__obf_6b7012530e205f39, 0}, nil
	}

	// TODO: optimize extraction of fractional part
	__obf_3b803b95d5fc8b6b := Decimal{__obf_6b7012530e205f39, 0}
	__obf_a15fe04a384142dd, __obf_f3a8aaed56831f9d := __obf_e5610ee7ba67ce83.QuoRem(__obf_3b803b95d5fc8b6b, 0)

	if __obf_e2d91f0337623bec == -1 && !__obf_f3a8aaed56831f9d.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent imaginary value of x ** y, where x < 0 and y is non-integer decimal")
	}

	__obf_09d2bbe0858a15c0, _ := __obf_19a369452cd5fbde.__obf_9ca5ed499a711b19(__obf_a15fe04a384142dd.__obf_6e5f77decbc21dea, __obf_fb38b245d2cdb3ee)

	// if exponent is an integer we don't need to calculate d1**frac(d2)
	if __obf_f3a8aaed56831f9d.__obf_6e5f77decbc21dea.Sign() == 0 {
		return __obf_09d2bbe0858a15c0, nil
	}

	// TODO: optimize NumDigits for more performant precision adjustment
	__obf_49d8c84b3237ddf4 := __obf_19a369452cd5fbde.NumDigits()
	__obf_388de7c089524859 := __obf_e5610ee7ba67ce83.NumDigits()

	if int32(__obf_49d8c84b3237ddf4) > __obf_fb38b245d2cdb3ee {
		__obf_fb38b245d2cdb3ee = int32(__obf_49d8c84b3237ddf4)
	}
	if int32(__obf_388de7c089524859) > __obf_fb38b245d2cdb3ee {
		__obf_fb38b245d2cdb3ee += int32(__obf_388de7c089524859)
	}
	// increase precision by 10 to compensate for errors in further calculations
	__obf_fb38b245d2cdb3ee += 10

	// Calculate x ** frac(y), where
	// x ** frac(y) = exp(ln(x ** frac(y)) = exp(ln(x) * frac(y))
	__obf_7ca7f3589705fc1e, __obf_68bcf986f9cdf0c0 := __obf_19a369452cd5fbde.Abs().Ln(__obf_fb38b245d2cdb3ee)
	if __obf_68bcf986f9cdf0c0 != nil {
		return Decimal{}, __obf_68bcf986f9cdf0c0
	}

	__obf_7ca7f3589705fc1e = __obf_7ca7f3589705fc1e.Mul(__obf_f3a8aaed56831f9d)

	__obf_7ca7f3589705fc1e, __obf_68bcf986f9cdf0c0 = __obf_7ca7f3589705fc1e.ExpTaylor(__obf_fb38b245d2cdb3ee)
	if __obf_68bcf986f9cdf0c0 != nil {
		return Decimal{}, __obf_68bcf986f9cdf0c0
	}

	// Join integer and fractional part,
	// base ** (expBase + expFrac) = base ** expBase * base ** expFrac
	__obf_a01a49266a86f271 := __obf_09d2bbe0858a15c0.Mul(__obf_7ca7f3589705fc1e)

	return __obf_a01a49266a86f271, nil
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
func (__obf_19a369452cd5fbde Decimal) PowInt32(__obf_1ef84065a6a49f43 int32) (Decimal, error) {
	if __obf_19a369452cd5fbde.IsZero() && __obf_1ef84065a6a49f43 == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_c520382c5238e14a := __obf_1ef84065a6a49f43 < 0
	__obf_1ef84065a6a49f43 = __obf_8f6522a0df8c3107(__obf_1ef84065a6a49f43)

	__obf_3edab89cf83bab1c, __obf_5a37cbbecbaac4ed := __obf_19a369452cd5fbde, New(1, 0)

	for __obf_1ef84065a6a49f43 > 0 {
		if __obf_1ef84065a6a49f43%2 == 1 {
			__obf_5a37cbbecbaac4ed = __obf_5a37cbbecbaac4ed.Mul(__obf_3edab89cf83bab1c)
		}
		__obf_1ef84065a6a49f43 /= 2

		if __obf_1ef84065a6a49f43 > 0 {
			__obf_3edab89cf83bab1c = __obf_3edab89cf83bab1c.Mul(__obf_3edab89cf83bab1c)
		}
	}

	if __obf_c520382c5238e14a {
		return New(1, 0).DivRound(__obf_5a37cbbecbaac4ed, int32(PowPrecisionNegativeExponent)), nil
	}

	return __obf_5a37cbbecbaac4ed, nil
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
func (__obf_19a369452cd5fbde Decimal) PowBigInt(__obf_1ef84065a6a49f43 *big.Int) (Decimal, error) {
	return __obf_19a369452cd5fbde.__obf_9ca5ed499a711b19(__obf_1ef84065a6a49f43, int32(PowPrecisionNegativeExponent))
}

func (__obf_19a369452cd5fbde Decimal) __obf_9ca5ed499a711b19(__obf_1ef84065a6a49f43 *big.Int, __obf_fb38b245d2cdb3ee int32) (Decimal, error) {
	if __obf_19a369452cd5fbde.IsZero() && __obf_1ef84065a6a49f43.Sign() == 0 {
		return Decimal{}, fmt.Errorf("cannot represent undefined value of 0**0")
	}

	__obf_9b28b1f06a86b092 := new(big.Int).Set(__obf_1ef84065a6a49f43)
	__obf_c520382c5238e14a := __obf_1ef84065a6a49f43.Sign() < 0

	if __obf_c520382c5238e14a {
		__obf_9b28b1f06a86b092.Abs(__obf_9b28b1f06a86b092)
	}

	__obf_3edab89cf83bab1c, __obf_5a37cbbecbaac4ed := __obf_19a369452cd5fbde, New(1, 0)

	for __obf_9b28b1f06a86b092.Sign() > 0 {
		if __obf_9b28b1f06a86b092.Bit(0) == 1 {
			__obf_5a37cbbecbaac4ed = __obf_5a37cbbecbaac4ed.Mul(__obf_3edab89cf83bab1c)
		}
		__obf_9b28b1f06a86b092.Rsh(__obf_9b28b1f06a86b092, 1)

		if __obf_9b28b1f06a86b092.Sign() > 0 {
			__obf_3edab89cf83bab1c = __obf_3edab89cf83bab1c.Mul(__obf_3edab89cf83bab1c)
		}
	}

	if __obf_c520382c5238e14a {
		return New(1, 0).DivRound(__obf_5a37cbbecbaac4ed, __obf_fb38b245d2cdb3ee), nil
	}

	return __obf_5a37cbbecbaac4ed, nil
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
func (__obf_19a369452cd5fbde Decimal) ExpHullAbrham(__obf_bfc48a42dcc5a5cc uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if __obf_19a369452cd5fbde.IsZero() {
		return Decimal{__obf_6b7012530e205f39, 0}, nil
	}

	__obf_b7f6e437e07cf74a := __obf_bfc48a42dcc5a5cc

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	__obf_fd7d6ef724312b50 := __obf_19a369452cd5fbde.Abs().InexactFloat64()
	if __obf_aca08fd18006ae92 := __obf_fd7d6ef724312b50 / 23; __obf_aca08fd18006ae92 > float64(__obf_b7f6e437e07cf74a) && __obf_aca08fd18006ae92 < float64(ExpMaxIterations) {
		__obf_b7f6e437e07cf74a = uint32(math.Ceil(__obf_aca08fd18006ae92))
	}

	// fail if abs(d) beyond an over/underflow threshold
	__obf_d63340ff75fc06f7 := New(23*int64(__obf_b7f6e437e07cf74a), 0)
	if __obf_19a369452cd5fbde.Abs().Cmp(__obf_d63340ff75fc06f7) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}

	// Return 1 if abs(d) small enough; this also avoids later over/underflow
	__obf_e844928fcd460900 := New(9, -int32(__obf_b7f6e437e07cf74a)-1)
	if __obf_19a369452cd5fbde.Abs().Cmp(__obf_e844928fcd460900) <= 0 {
		return Decimal{__obf_6b7012530e205f39, __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43}, nil
	}

	// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
	__obf_0e45b010b2b05e0d := __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 + int32(__obf_19a369452cd5fbde.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if __obf_0e45b010b2b05e0d < 0 {
		__obf_0e45b010b2b05e0d = 0
	}

	__obf_ec401ce06db88a9f := New(1, __obf_0e45b010b2b05e0d)                                                                                                                   // reduction factor
	__obf_0feccac4ef34d0a4 := Decimal{new(big.Int).Set(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea), __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 - __obf_0e45b010b2b05e0d} // reduced argument
	__obf_f00a4e020967bf11 := int32(__obf_b7f6e437e07cf74a) + __obf_0e45b010b2b05e0d + 2                                                                                       // precision for calculating the sum

	// Determine n, the number of therms for calculating sum
	// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
	// for solving appropriate equation, along with directed
	// roundings and simple rational bound for log10(p/abs(r))
	__obf_cfeb02a08e0decd8 := __obf_0feccac4ef34d0a4.Abs().InexactFloat64()
	__obf_35a16f64f7e5af88 := float64(__obf_f00a4e020967bf11)
	__obf_d14c538f3ac0369c := math.Ceil((1.453*__obf_35a16f64f7e5af88 - 1.182) / math.Log10(__obf_35a16f64f7e5af88/__obf_cfeb02a08e0decd8))
	if __obf_d14c538f3ac0369c > float64(ExpMaxIterations) || math.IsNaN(__obf_d14c538f3ac0369c) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	__obf_3edab89cf83bab1c := int64(__obf_d14c538f3ac0369c)

	__obf_e059bdc33d645894 := New(0, 0)
	__obf_1882daa5a5fd23a5 := New(1, 0)
	__obf_3b803b95d5fc8b6b := New(1, 0)
	for __obf_40720372b6d983c7 := __obf_3edab89cf83bab1c - 1; __obf_40720372b6d983c7 > 0; __obf_40720372b6d983c7-- {
		__obf_e059bdc33d645894.__obf_6e5f77decbc21dea.SetInt64(__obf_40720372b6d983c7)
		__obf_1882daa5a5fd23a5 = __obf_1882daa5a5fd23a5.Mul(__obf_0feccac4ef34d0a4.DivRound(__obf_e059bdc33d645894, __obf_f00a4e020967bf11))
		__obf_1882daa5a5fd23a5 = __obf_1882daa5a5fd23a5.Add(__obf_3b803b95d5fc8b6b)
	}

	__obf_398ac2a491ab93f4 := __obf_ec401ce06db88a9f.IntPart()
	__obf_a01a49266a86f271 := New(1, 0)
	for __obf_40720372b6d983c7 := __obf_398ac2a491ab93f4; __obf_40720372b6d983c7 > 0; __obf_40720372b6d983c7-- {
		__obf_a01a49266a86f271 = __obf_a01a49266a86f271.Mul(__obf_1882daa5a5fd23a5)
	}

	__obf_195eb4899c090619 := int32(__obf_a01a49266a86f271.NumDigits())

	var __obf_86c41a5c5717f0f5 int32
	if __obf_195eb4899c090619 > __obf_8f6522a0df8c3107(__obf_a01a49266a86f271.__obf_1ef84065a6a49f43) {
		__obf_86c41a5c5717f0f5 = int32(__obf_b7f6e437e07cf74a) - __obf_195eb4899c090619 - __obf_a01a49266a86f271.__obf_1ef84065a6a49f43
	} else {
		__obf_86c41a5c5717f0f5 = int32(__obf_b7f6e437e07cf74a)
	}

	__obf_a01a49266a86f271 = __obf_a01a49266a86f271.Round(__obf_86c41a5c5717f0f5)

	return __obf_a01a49266a86f271, nil
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
func (__obf_19a369452cd5fbde Decimal) ExpTaylor(__obf_fb38b245d2cdb3ee int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if __obf_19a369452cd5fbde.IsZero() {
		return Decimal{__obf_6b7012530e205f39, 0}.Round(__obf_fb38b245d2cdb3ee), nil
	}

	var __obf_cf5de1db9e631594 Decimal
	var __obf_25ab290f5081581a int32
	if __obf_fb38b245d2cdb3ee < 0 {
		__obf_cf5de1db9e631594 = New(1, -1)
		__obf_25ab290f5081581a = 8
	} else {
		__obf_cf5de1db9e631594 = New(1, -__obf_fb38b245d2cdb3ee-1)
		__obf_25ab290f5081581a = __obf_fb38b245d2cdb3ee + 1
	}

	__obf_6d46c256a4bb47b8 := __obf_19a369452cd5fbde.Abs()
	__obf_b954d39f1598c041 := __obf_19a369452cd5fbde.Abs()
	__obf_5794cb2cfcaed684 := New(1, 0)

	__obf_5a37cbbecbaac4ed := New(1, 0)

	for __obf_40720372b6d983c7 := int64(1); ; {
		__obf_20e97b9a2de85df5 := __obf_b954d39f1598c041.DivRound(__obf_5794cb2cfcaed684, __obf_25ab290f5081581a)
		__obf_5a37cbbecbaac4ed = __obf_5a37cbbecbaac4ed.Add(__obf_20e97b9a2de85df5)

		// Stop Taylor series when current step is smaller than epsilon
		if __obf_20e97b9a2de85df5.Cmp(__obf_cf5de1db9e631594) < 0 {
			break
		}

		__obf_b954d39f1598c041 = __obf_b954d39f1598c041.Mul(__obf_6d46c256a4bb47b8)

		__obf_40720372b6d983c7++

		// Calculate next factorial number or retrieve cached value
		if len(__obf_5004c7371a2d0179) >= int(__obf_40720372b6d983c7) && !__obf_5004c7371a2d0179[__obf_40720372b6d983c7-1].IsZero() {
			__obf_5794cb2cfcaed684 = __obf_5004c7371a2d0179[__obf_40720372b6d983c7-1]
		} else {
			// To avoid any race conditions, firstly the zero value is appended to a slice to create
			// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
			// factorial using the index notation.
			__obf_5794cb2cfcaed684 = __obf_5004c7371a2d0179[__obf_40720372b6d983c7-2].Mul(New(__obf_40720372b6d983c7, 0))
			__obf_5004c7371a2d0179 = append(__obf_5004c7371a2d0179, Zero)
			__obf_5004c7371a2d0179[__obf_40720372b6d983c7-1] = __obf_5794cb2cfcaed684
		}
	}

	if __obf_19a369452cd5fbde.Sign() < 0 {
		__obf_5a37cbbecbaac4ed = New(1, 0).DivRound(__obf_5a37cbbecbaac4ed, __obf_fb38b245d2cdb3ee+1)
	}

	__obf_5a37cbbecbaac4ed = __obf_5a37cbbecbaac4ed.Round(__obf_fb38b245d2cdb3ee)
	return __obf_5a37cbbecbaac4ed, nil
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
func (__obf_19a369452cd5fbde Decimal) Ln(__obf_fb38b245d2cdb3ee int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if __obf_19a369452cd5fbde.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if __obf_19a369452cd5fbde.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}

	__obf_8c06c377a56a9277 := __obf_fb38b245d2cdb3ee + 2
	__obf_f0a3b6dafda7d395 := __obf_19a369452cd5fbde.Copy()

	var __obf_291675a83df87dce, __obf_72676d7e0f36a12e, __obf_110691455bfd43fc, __obf_fdc0dd5534b24a23, __obf_bf7c636ae74a6766 Decimal
	__obf_291675a83df87dce = __obf_f0a3b6dafda7d395.Sub(Decimal{__obf_6b7012530e205f39, 0})
	__obf_72676d7e0f36a12e = Decimal{__obf_6b7012530e205f39, -1}

	// for decimal in range [0.9, 1.1] where ln(d) is close to 0
	__obf_38c47c160fcb6992 := false

	if __obf_291675a83df87dce.Abs().Cmp(__obf_72676d7e0f36a12e) <= 0 {
		__obf_38c47c160fcb6992 = true
	} else {
		// reduce input decimal to range [0.1, 1)
		__obf_08915d087dbf986e := int32(__obf_f0a3b6dafda7d395.NumDigits()) + __obf_f0a3b6dafda7d395.__obf_1ef84065a6a49f43
		__obf_f0a3b6dafda7d395.__obf_1ef84065a6a49f43 -= __obf_08915d087dbf986e

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		__obf_d03443580eb472ff := __obf_d03443580eb472ff.__obf_a43f9e51ea49d1d6(__obf_8c06c377a56a9277)
		__obf_bf7c636ae74a6766 = NewFromInt32(__obf_08915d087dbf986e)
		__obf_bf7c636ae74a6766 = __obf_bf7c636ae74a6766.Mul(__obf_d03443580eb472ff)

		__obf_291675a83df87dce = __obf_f0a3b6dafda7d395.Sub(Decimal{__obf_6b7012530e205f39, 0})

		if __obf_291675a83df87dce.Abs().Cmp(__obf_72676d7e0f36a12e) <= 0 {
			__obf_38c47c160fcb6992 = true
		} else {
			// initial estimate using floats
			__obf_77c091eaef07d7d3 := __obf_f0a3b6dafda7d395.InexactFloat64()
			__obf_291675a83df87dce = NewFromFloat(math.Log(__obf_77c091eaef07d7d3))
		}
	}

	__obf_cf5de1db9e631594 := Decimal{__obf_6b7012530e205f39, -__obf_8c06c377a56a9277}

	if __obf_38c47c160fcb6992 {
		// Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
		// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
		// until the difference between current and next term is smaller than epsilon.
		// Coverage quite fast for decimals close to 1.0

		// z + 2
		__obf_110691455bfd43fc = __obf_291675a83df87dce.Add(Decimal{__obf_3338f87aa2a5f65a, 0})
		// z / (z + 2)
		__obf_72676d7e0f36a12e = __obf_291675a83df87dce.DivRound(__obf_110691455bfd43fc, __obf_8c06c377a56a9277)
		// 2 * (z / (z + 2))
		__obf_291675a83df87dce = __obf_72676d7e0f36a12e.Add(__obf_72676d7e0f36a12e)
		__obf_110691455bfd43fc = __obf_291675a83df87dce.Copy()

		for __obf_3edab89cf83bab1c := 1; ; __obf_3edab89cf83bab1c++ {
			// 2 * (z / (z+2))^(2n+1)
			__obf_110691455bfd43fc = __obf_110691455bfd43fc.Mul(__obf_72676d7e0f36a12e).Mul(__obf_72676d7e0f36a12e)

			// 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
			__obf_fdc0dd5534b24a23 = NewFromInt(int64(2*__obf_3edab89cf83bab1c + 1))
			__obf_fdc0dd5534b24a23 = __obf_110691455bfd43fc.DivRound(__obf_fdc0dd5534b24a23, __obf_8c06c377a56a9277)

			// comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			__obf_291675a83df87dce = __obf_291675a83df87dce.Add(__obf_fdc0dd5534b24a23)

			if __obf_fdc0dd5534b24a23.Abs().Cmp(__obf_cf5de1db9e631594) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		var __obf_eb489493ed7f4e19 Decimal
		__obf_f9780f2e47b88006 := __obf_8c06c377a56a9277*2 + 10

		for __obf_40720372b6d983c7 := int32(0); __obf_40720372b6d983c7 < __obf_f9780f2e47b88006; __obf_40720372b6d983c7++ {
			// exp(a_n)
			__obf_72676d7e0f36a12e, _ = __obf_291675a83df87dce.ExpTaylor(__obf_8c06c377a56a9277)
			// exp(a_n) - z
			__obf_110691455bfd43fc = __obf_72676d7e0f36a12e.Sub(__obf_f0a3b6dafda7d395)
			// 2 * (exp(a_n) - z)
			__obf_110691455bfd43fc = __obf_110691455bfd43fc.Add(__obf_110691455bfd43fc)
			// exp(a_n) + z
			__obf_fdc0dd5534b24a23 = __obf_72676d7e0f36a12e.Add(__obf_f0a3b6dafda7d395)
			// 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_72676d7e0f36a12e = __obf_110691455bfd43fc.DivRound(__obf_fdc0dd5534b24a23, __obf_8c06c377a56a9277)
			// comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
			__obf_291675a83df87dce = __obf_291675a83df87dce.Sub(__obf_72676d7e0f36a12e)

			if __obf_eb489493ed7f4e19.Add(__obf_72676d7e0f36a12e).IsZero() {
				// If iteration steps oscillate we should return early and prevent an infinity loop
				// NOTE(mwoss): This should be quite a rare case, returning error is not necessary
				break
			}

			if __obf_72676d7e0f36a12e.Abs().Cmp(__obf_cf5de1db9e631594) <= 0 {
				break
			}

			__obf_eb489493ed7f4e19 = __obf_72676d7e0f36a12e
		}
	}

	__obf_291675a83df87dce = __obf_291675a83df87dce.Add(__obf_bf7c636ae74a6766)

	return __obf_291675a83df87dce.Round(__obf_fb38b245d2cdb3ee), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
func (__obf_19a369452cd5fbde Decimal) NumDigits() int {
	if __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea == nil {
		return 1
	}

	if __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.IsInt64() {
		__obf_bd0a6195fd130370 := __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.Int64()
		// restrict fast path to integers with exact conversion to float64
		if __obf_bd0a6195fd130370 <= (1<<53) && __obf_bd0a6195fd130370 >= -(1<<53) {
			if __obf_bd0a6195fd130370 == 0 {
				return 1
			}
			return int(math.Log10(math.Abs(float64(__obf_bd0a6195fd130370)))) + 1
		}
	}

	__obf_e73455e4b039eecd := int(float64(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.BitLen()) / math.Log2(10))

	// estimatedNumDigits (lg10) may be off by 1, need to verify
	__obf_89bbd71996a3d55d := big.NewInt(int64(__obf_e73455e4b039eecd))
	__obf_d414c33699262c20 := __obf_89bbd71996a3d55d.Exp(__obf_6e3308650b479548, __obf_89bbd71996a3d55d, nil)

	if __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.CmpAbs(__obf_d414c33699262c20) >= 0 {
		return __obf_e73455e4b039eecd + 1
	}

	return __obf_e73455e4b039eecd
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (__obf_19a369452cd5fbde Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var __obf_0feccac4ef34d0a4 big.Int
	__obf_99a04cdb50494bb6 := new(big.Int).Set(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea)
	for __obf_f0a3b6dafda7d395 := __obf_8f6522a0df8c3107(__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43); __obf_f0a3b6dafda7d395 > 0; __obf_f0a3b6dafda7d395-- {
		__obf_99a04cdb50494bb6.QuoRem(__obf_99a04cdb50494bb6, __obf_6e3308650b479548, &__obf_0feccac4ef34d0a4)
		if __obf_0feccac4ef34d0a4.Cmp(__obf_91bffcce19efb7d8) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func __obf_8f6522a0df8c3107(__obf_3edab89cf83bab1c int32) int32 {
	if __obf_3edab89cf83bab1c < 0 {
		return -__obf_3edab89cf83bab1c
	}
	return __obf_3edab89cf83bab1c
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_19a369452cd5fbde Decimal) Cmp(__obf_e5610ee7ba67ce83 Decimal) int {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()
	__obf_e5610ee7ba67ce83.__obf_b1325bfc63b249a7()

	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 == __obf_e5610ee7ba67ce83.__obf_1ef84065a6a49f43 {
		return __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.Cmp(__obf_e5610ee7ba67ce83.__obf_6e5f77decbc21dea)
	}

	__obf_e569aa78d09dc408, __obf_f8f48e5992257be1 := RescalePair(__obf_19a369452cd5fbde, __obf_e5610ee7ba67ce83)

	return __obf_e569aa78d09dc408.__obf_6e5f77decbc21dea.Cmp(__obf_f8f48e5992257be1.__obf_6e5f77decbc21dea)
}

// Compare compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (__obf_19a369452cd5fbde Decimal) Compare(__obf_e5610ee7ba67ce83 Decimal) int {
	return __obf_19a369452cd5fbde.Cmp(__obf_e5610ee7ba67ce83)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (__obf_19a369452cd5fbde Decimal) Equal(__obf_e5610ee7ba67ce83 Decimal) bool {
	return __obf_19a369452cd5fbde.Cmp(__obf_e5610ee7ba67ce83) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (__obf_19a369452cd5fbde Decimal) Equals(__obf_e5610ee7ba67ce83 Decimal) bool {
	return __obf_19a369452cd5fbde.Equal(__obf_e5610ee7ba67ce83)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (__obf_19a369452cd5fbde Decimal) GreaterThan(__obf_e5610ee7ba67ce83 Decimal) bool {
	return __obf_19a369452cd5fbde.Cmp(__obf_e5610ee7ba67ce83) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (__obf_19a369452cd5fbde Decimal) GreaterThanOrEqual(__obf_e5610ee7ba67ce83 Decimal) bool {
	__obf_633ebead09af5abb := __obf_19a369452cd5fbde.Cmp(__obf_e5610ee7ba67ce83)
	return __obf_633ebead09af5abb == 1 || __obf_633ebead09af5abb == 0
}

// LessThan (LT) returns true when d is less than d2.
func (__obf_19a369452cd5fbde Decimal) LessThan(__obf_e5610ee7ba67ce83 Decimal) bool {
	return __obf_19a369452cd5fbde.Cmp(__obf_e5610ee7ba67ce83) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (__obf_19a369452cd5fbde Decimal) LessThanOrEqual(__obf_e5610ee7ba67ce83 Decimal) bool {
	__obf_633ebead09af5abb := __obf_19a369452cd5fbde.Cmp(__obf_e5610ee7ba67ce83)
	return __obf_633ebead09af5abb == -1 || __obf_633ebead09af5abb == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (__obf_19a369452cd5fbde Decimal) Sign() int {
	if __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea == nil {
		return 0
	}
	return __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.Sign()
}

// IsPositive return
//
//	true if d > 0
//	false if d == 0
//	false if d < 0
func (__obf_19a369452cd5fbde Decimal) IsPositive() bool {
	return __obf_19a369452cd5fbde.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (__obf_19a369452cd5fbde Decimal) IsNegative() bool {
	return __obf_19a369452cd5fbde.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (__obf_19a369452cd5fbde Decimal) IsZero() bool {
	return __obf_19a369452cd5fbde.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (__obf_19a369452cd5fbde Decimal) Exponent() int32 {
	return __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43
}

// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
func (__obf_19a369452cd5fbde Decimal) Coefficient() *big.Int {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea)
}

// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
// If coefficient cannot be represented in an int64, the result will be undefined.
func (__obf_19a369452cd5fbde Decimal) CoefficientInt64() int64 {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()
	return __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.Int64()
}

// IntPart returns the integer component of the decimal.
func (__obf_19a369452cd5fbde Decimal) IntPart() int64 {
	__obf_fe4a29a2006b0e73 := __obf_19a369452cd5fbde.__obf_957a7b5240163386(0)
	return __obf_fe4a29a2006b0e73.__obf_6e5f77decbc21dea.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (__obf_19a369452cd5fbde Decimal) BigInt() *big.Int {
	__obf_fe4a29a2006b0e73 := __obf_19a369452cd5fbde.__obf_957a7b5240163386(0)
	return __obf_fe4a29a2006b0e73.__obf_6e5f77decbc21dea
}

// BigFloat returns decimal as BigFloat.
// Be aware that casting decimal to BigFloat might cause a loss of precision.
func (__obf_19a369452cd5fbde Decimal) BigFloat() *big.Float {
	__obf_fd7d6ef724312b50 := &big.Float{}
	__obf_fd7d6ef724312b50.SetString(__obf_19a369452cd5fbde.String())
	return __obf_fd7d6ef724312b50
}

// Rat returns a rational number representation of the decimal.
func (__obf_19a369452cd5fbde Decimal) Rat() *big.Rat {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()
	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 <= 0 {
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		__obf_dec3313ff2430ad1 := new(big.Int).Exp(__obf_6e3308650b479548, big.NewInt(-int64(__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43)), nil)
		return new(big.Rat).SetFrac(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea, __obf_dec3313ff2430ad1)
	}

	__obf_5f9f7cde1c1cb6b0 := new(big.Int).Exp(__obf_6e3308650b479548, big.NewInt(int64(__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43)), nil)
	__obf_f33935c2c4a07945 := new(big.Int).Mul(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea, __obf_5f9f7cde1c1cb6b0)
	return new(big.Rat).SetFrac(__obf_f33935c2c4a07945, __obf_6b7012530e205f39)
}

// Float64 returns the nearest float64 value for d and a bool indicating
// whether f represents d exactly.
// For more details, see the documentation for big.Rat.Float64
func (__obf_19a369452cd5fbde Decimal) Float64() (__obf_fd7d6ef724312b50 float64, __obf_8fed91263b695030 bool) {
	return __obf_19a369452cd5fbde.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (__obf_19a369452cd5fbde Decimal) InexactFloat64() float64 {
	__obf_fd7d6ef724312b50, _ := __obf_19a369452cd5fbde.Float64()
	return __obf_fd7d6ef724312b50
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
func (__obf_19a369452cd5fbde Decimal) String() string {
	return __obf_19a369452cd5fbde.string(true)
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
func (__obf_19a369452cd5fbde Decimal) StringFixed(__obf_eb137eafa1a56d73 int32) string {
	__obf_cd9f5961f510f6b4 := __obf_19a369452cd5fbde.Round(__obf_eb137eafa1a56d73)
	return __obf_cd9f5961f510f6b4.string(false)
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
func (__obf_19a369452cd5fbde Decimal) StringFixedBank(__obf_eb137eafa1a56d73 int32) string {
	__obf_cd9f5961f510f6b4 := __obf_19a369452cd5fbde.RoundBank(__obf_eb137eafa1a56d73)
	return __obf_cd9f5961f510f6b4.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (__obf_19a369452cd5fbde Decimal) StringFixedCash(__obf_0f11bf0455c26e02 uint8) string {
	__obf_cd9f5961f510f6b4 := __obf_19a369452cd5fbde.RoundCash(__obf_0f11bf0455c26e02)
	return __obf_cd9f5961f510f6b4.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (__obf_19a369452cd5fbde Decimal) Round(__obf_eb137eafa1a56d73 int32) Decimal {
	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 == -__obf_eb137eafa1a56d73 {
		return __obf_19a369452cd5fbde
	}
	// truncate to places + 1
	__obf_0f6065acafe14405 := __obf_19a369452cd5fbde.__obf_957a7b5240163386(-__obf_eb137eafa1a56d73 - 1)

	// add sign(d) * 0.5
	if __obf_0f6065acafe14405.__obf_6e5f77decbc21dea.Sign() < 0 {
		__obf_0f6065acafe14405.__obf_6e5f77decbc21dea.Sub(__obf_0f6065acafe14405.__obf_6e5f77decbc21dea, __obf_a02c1e3273085f0d)
	} else {
		__obf_0f6065acafe14405.__obf_6e5f77decbc21dea.Add(__obf_0f6065acafe14405.__obf_6e5f77decbc21dea, __obf_a02c1e3273085f0d)
	}

	// floor for positive numbers, ceil for negative numbers
	_, __obf_3f873ba4f9076041 := __obf_0f6065acafe14405.__obf_6e5f77decbc21dea.DivMod(__obf_0f6065acafe14405.__obf_6e5f77decbc21dea, __obf_6e3308650b479548, new(big.Int))
	__obf_0f6065acafe14405.__obf_1ef84065a6a49f43++
	if __obf_0f6065acafe14405.__obf_6e5f77decbc21dea.Sign() < 0 && __obf_3f873ba4f9076041.Cmp(__obf_91bffcce19efb7d8) != 0 {
		__obf_0f6065acafe14405.__obf_6e5f77decbc21dea.Add(__obf_0f6065acafe14405.__obf_6e5f77decbc21dea, __obf_6b7012530e205f39)
	}

	return __obf_0f6065acafe14405
}

// RoundCeil rounds the decimal towards +infinity.
//
// Example:
//
//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
func (__obf_19a369452cd5fbde Decimal) RoundCeil(__obf_eb137eafa1a56d73 int32) Decimal {
	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 >= -__obf_eb137eafa1a56d73 {
		return __obf_19a369452cd5fbde
	}

	__obf_73ec129015c1fb73 := __obf_19a369452cd5fbde.__obf_957a7b5240163386(-__obf_eb137eafa1a56d73)
	if __obf_19a369452cd5fbde.Equal(__obf_73ec129015c1fb73) {
		return __obf_19a369452cd5fbde
	}

	if __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.Sign() > 0 {
		__obf_73ec129015c1fb73.__obf_6e5f77decbc21dea.Add(__obf_73ec129015c1fb73.__obf_6e5f77decbc21dea, __obf_6b7012530e205f39)
	}

	return __obf_73ec129015c1fb73
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (__obf_19a369452cd5fbde Decimal) RoundFloor(__obf_eb137eafa1a56d73 int32) Decimal {
	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 >= -__obf_eb137eafa1a56d73 {
		return __obf_19a369452cd5fbde
	}

	__obf_73ec129015c1fb73 := __obf_19a369452cd5fbde.__obf_957a7b5240163386(-__obf_eb137eafa1a56d73)
	if __obf_19a369452cd5fbde.Equal(__obf_73ec129015c1fb73) {
		return __obf_19a369452cd5fbde
	}

	if __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.Sign() < 0 {
		__obf_73ec129015c1fb73.__obf_6e5f77decbc21dea.Sub(__obf_73ec129015c1fb73.__obf_6e5f77decbc21dea, __obf_6b7012530e205f39)
	}

	return __obf_73ec129015c1fb73
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (__obf_19a369452cd5fbde Decimal) RoundUp(__obf_eb137eafa1a56d73 int32) Decimal {
	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 >= -__obf_eb137eafa1a56d73 {
		return __obf_19a369452cd5fbde
	}

	__obf_73ec129015c1fb73 := __obf_19a369452cd5fbde.__obf_957a7b5240163386(-__obf_eb137eafa1a56d73)
	if __obf_19a369452cd5fbde.Equal(__obf_73ec129015c1fb73) {
		return __obf_19a369452cd5fbde
	}

	if __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.Sign() > 0 {
		__obf_73ec129015c1fb73.__obf_6e5f77decbc21dea.Add(__obf_73ec129015c1fb73.__obf_6e5f77decbc21dea, __obf_6b7012530e205f39)
	} else if __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.Sign() < 0 {
		__obf_73ec129015c1fb73.__obf_6e5f77decbc21dea.Sub(__obf_73ec129015c1fb73.__obf_6e5f77decbc21dea, __obf_6b7012530e205f39)
	}

	return __obf_73ec129015c1fb73
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (__obf_19a369452cd5fbde Decimal) RoundDown(__obf_eb137eafa1a56d73 int32) Decimal {
	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 >= -__obf_eb137eafa1a56d73 {
		return __obf_19a369452cd5fbde
	}

	__obf_73ec129015c1fb73 := __obf_19a369452cd5fbde.__obf_957a7b5240163386(-__obf_eb137eafa1a56d73)
	if __obf_19a369452cd5fbde.Equal(__obf_73ec129015c1fb73) {
		return __obf_19a369452cd5fbde
	}
	return __obf_73ec129015c1fb73
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
func (__obf_19a369452cd5fbde Decimal) RoundBank(__obf_eb137eafa1a56d73 int32) Decimal {

	__obf_27ea7adcf083f578 := __obf_19a369452cd5fbde.Round(__obf_eb137eafa1a56d73)
	__obf_e01abbc36ecc041e := __obf_19a369452cd5fbde.Sub(__obf_27ea7adcf083f578).Abs()

	__obf_51d13fb3aa70db27 := New(5, -__obf_eb137eafa1a56d73-1)
	if __obf_e01abbc36ecc041e.Cmp(__obf_51d13fb3aa70db27) == 0 && __obf_27ea7adcf083f578.__obf_6e5f77decbc21dea.Bit(0) != 0 {
		if __obf_27ea7adcf083f578.__obf_6e5f77decbc21dea.Sign() < 0 {
			__obf_27ea7adcf083f578.__obf_6e5f77decbc21dea.Add(__obf_27ea7adcf083f578.__obf_6e5f77decbc21dea, __obf_6b7012530e205f39)
		} else {
			__obf_27ea7adcf083f578.__obf_6e5f77decbc21dea.Sub(__obf_27ea7adcf083f578.__obf_6e5f77decbc21dea, __obf_6b7012530e205f39)
		}
	}

	return __obf_27ea7adcf083f578
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
func (__obf_19a369452cd5fbde Decimal) RoundCash(__obf_0f11bf0455c26e02 uint8) Decimal {
	var __obf_2ef6e82509e87a46 *big.Int
	switch __obf_0f11bf0455c26e02 {
	case 5:
		__obf_2ef6e82509e87a46 = __obf_cc9bfe65d22c82e4
	case 10:
		__obf_2ef6e82509e87a46 = __obf_6e3308650b479548
	case 25:
		__obf_2ef6e82509e87a46 = __obf_7a3c2c07b99ffb48
	case 50:
		__obf_2ef6e82509e87a46 = __obf_3338f87aa2a5f65a
	case 100:
		__obf_2ef6e82509e87a46 = __obf_6b7012530e205f39
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", __obf_0f11bf0455c26e02))
	}
	__obf_473188eb89a369e4 := Decimal{
		__obf_6e5f77decbc21dea: __obf_2ef6e82509e87a46,
	}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return __obf_19a369452cd5fbde.Mul(__obf_473188eb89a369e4).Round(0).Div(__obf_473188eb89a369e4).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (__obf_19a369452cd5fbde Decimal) Floor() Decimal {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()

	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 >= 0 {
		return __obf_19a369452cd5fbde
	}

	__obf_1ef84065a6a49f43 := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_1ef84065a6a49f43.Exp(__obf_1ef84065a6a49f43, big.NewInt(-int64(__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43)), nil)

	__obf_f0a3b6dafda7d395 := new(big.Int).Div(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea, __obf_1ef84065a6a49f43)
	return Decimal{__obf_6e5f77decbc21dea: __obf_f0a3b6dafda7d395, __obf_1ef84065a6a49f43: 0}
}

// Ceil returns the nearest integer value greater than or equal to d.
func (__obf_19a369452cd5fbde Decimal) Ceil() Decimal {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()

	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 >= 0 {
		return __obf_19a369452cd5fbde
	}

	__obf_1ef84065a6a49f43 := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	__obf_1ef84065a6a49f43.Exp(__obf_1ef84065a6a49f43, big.NewInt(-int64(__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43)), nil)

	__obf_f0a3b6dafda7d395, __obf_3f873ba4f9076041 := new(big.Int).DivMod(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea, __obf_1ef84065a6a49f43, new(big.Int))
	if __obf_3f873ba4f9076041.Cmp(__obf_91bffcce19efb7d8) != 0 {
		__obf_f0a3b6dafda7d395.Add(__obf_f0a3b6dafda7d395, __obf_6b7012530e205f39)
	}
	return Decimal{__obf_6e5f77decbc21dea: __obf_f0a3b6dafda7d395, __obf_1ef84065a6a49f43: 0}
}

// Truncate truncates off digits from the number, without rounding.
//
// NOTE: precision is the last digit that will not be truncated (must be >= 0).
//
// Example:
//
//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
func (__obf_19a369452cd5fbde Decimal) Truncate(__obf_fb38b245d2cdb3ee int32) Decimal {
	__obf_19a369452cd5fbde.__obf_b1325bfc63b249a7()
	if __obf_fb38b245d2cdb3ee >= 0 && -__obf_fb38b245d2cdb3ee > __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 {
		return __obf_19a369452cd5fbde.__obf_957a7b5240163386(-__obf_fb38b245d2cdb3ee)
	}
	return __obf_19a369452cd5fbde
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_19a369452cd5fbde *Decimal) UnmarshalJSON(__obf_bb0607f4ee95f9e9 []byte) error {
	if string(__obf_bb0607f4ee95f9e9) == "null" {
		return nil
	}

	__obf_1e6e3de76535e1bf, __obf_68bcf986f9cdf0c0 := __obf_b4b0c2ff19759ebc(__obf_bb0607f4ee95f9e9)
	if __obf_68bcf986f9cdf0c0 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_bb0607f4ee95f9e9, __obf_68bcf986f9cdf0c0)
	}

	__obf_ae16adf734cfe1aa, __obf_68bcf986f9cdf0c0 := NewFromString(__obf_1e6e3de76535e1bf)
	*__obf_19a369452cd5fbde = __obf_ae16adf734cfe1aa
	if __obf_68bcf986f9cdf0c0 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_1e6e3de76535e1bf, __obf_68bcf986f9cdf0c0)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_19a369452cd5fbde Decimal) MarshalJSON() ([]byte, error) {
	var __obf_1e6e3de76535e1bf string
	if MarshalJSONWithoutQuotes {
		__obf_1e6e3de76535e1bf = __obf_19a369452cd5fbde.String()
	} else {
		__obf_1e6e3de76535e1bf = "\"" + __obf_19a369452cd5fbde.String() + "\""
	}
	return []byte(__obf_1e6e3de76535e1bf), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (__obf_19a369452cd5fbde *Decimal) UnmarshalBinary(__obf_b774665a4676567d []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(__obf_b774665a4676567d) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", __obf_b774665a4676567d, len(__obf_b774665a4676567d))
	}

	// Extract the exponent
	__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 = int32(binary.BigEndian.Uint32(__obf_b774665a4676567d[:4]))

	// Extract the value
	__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea = new(big.Int)
	if __obf_68bcf986f9cdf0c0 := __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.GobDecode(__obf_b774665a4676567d[4:]); __obf_68bcf986f9cdf0c0 != nil {
		return fmt.Errorf("error decoding binary %v: %s", __obf_b774665a4676567d, __obf_68bcf986f9cdf0c0)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (__obf_19a369452cd5fbde Decimal) MarshalBinary() (__obf_b774665a4676567d []byte, __obf_68bcf986f9cdf0c0 error) {
	// exp is written first, but encode value first to know output size
	var __obf_551e83f3d41c5c97 []byte
	if __obf_551e83f3d41c5c97, __obf_68bcf986f9cdf0c0 = __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.GobEncode(); __obf_68bcf986f9cdf0c0 != nil {
		return nil, __obf_68bcf986f9cdf0c0
	}

	// Write the exponent in front, since it's a fixed size
	__obf_d9b150080a98dfe6 := make([]byte, 4, len(__obf_551e83f3d41c5c97)+4)
	binary.BigEndian.PutUint32(__obf_d9b150080a98dfe6, uint32(__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43))

	// Return the byte array
	return append(__obf_d9b150080a98dfe6, __obf_551e83f3d41c5c97...), nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_19a369452cd5fbde *Decimal) Scan(__obf_6e5f77decbc21dea any) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch __obf_847c12b47a12274f := __obf_6e5f77decbc21dea.(type) {

	case float32:
		*__obf_19a369452cd5fbde = NewFromFloat(float64(__obf_847c12b47a12274f))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*__obf_19a369452cd5fbde = NewFromFloat(__obf_847c12b47a12274f)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*__obf_19a369452cd5fbde = New(__obf_847c12b47a12274f, 0)
		return nil

	case uint64:
		// while clickhouse may send 0 in db as uint64
		*__obf_19a369452cd5fbde = NewFromUint64(__obf_847c12b47a12274f)
		return nil

	default:
		// default is trying to interpret value stored as string
		__obf_1e6e3de76535e1bf, __obf_68bcf986f9cdf0c0 := __obf_b4b0c2ff19759ebc(__obf_847c12b47a12274f)
		if __obf_68bcf986f9cdf0c0 != nil {
			return __obf_68bcf986f9cdf0c0
		}
		*__obf_19a369452cd5fbde, __obf_68bcf986f9cdf0c0 = NewFromString(__obf_1e6e3de76535e1bf)
		return __obf_68bcf986f9cdf0c0
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_19a369452cd5fbde Decimal) Value() (driver.Value, error) {
	return __obf_19a369452cd5fbde.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (__obf_19a369452cd5fbde *Decimal) UnmarshalText(__obf_206229668f99ca6d []byte) error {
	__obf_1e6e3de76535e1bf := string(__obf_206229668f99ca6d)

	__obf_e8ecc3b172d7ce84, __obf_68bcf986f9cdf0c0 := NewFromString(__obf_1e6e3de76535e1bf)
	*__obf_19a369452cd5fbde = __obf_e8ecc3b172d7ce84
	if __obf_68bcf986f9cdf0c0 != nil {
		return fmt.Errorf("error decoding string '%s': %s", __obf_1e6e3de76535e1bf, __obf_68bcf986f9cdf0c0)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_19a369452cd5fbde Decimal) MarshalText() (__obf_206229668f99ca6d []byte, __obf_68bcf986f9cdf0c0 error) {
	return []byte(__obf_19a369452cd5fbde.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (__obf_19a369452cd5fbde Decimal) GobEncode() ([]byte, error) {
	return __obf_19a369452cd5fbde.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (__obf_19a369452cd5fbde *Decimal) GobDecode(__obf_b774665a4676567d []byte) error {
	return __obf_19a369452cd5fbde.UnmarshalBinary(__obf_b774665a4676567d)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (__obf_19a369452cd5fbde Decimal) StringScaled(__obf_1ef84065a6a49f43 int32) string {
	return __obf_19a369452cd5fbde.__obf_957a7b5240163386(__obf_1ef84065a6a49f43).String()
}

func (__obf_19a369452cd5fbde Decimal) string(__obf_06dced229dc88d1a bool) string {
	if __obf_19a369452cd5fbde.__obf_1ef84065a6a49f43 >= 0 {
		return __obf_19a369452cd5fbde.__obf_957a7b5240163386(0).__obf_6e5f77decbc21dea.String()
	}

	__obf_8f6522a0df8c3107 := new(big.Int).Abs(__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea)
	__obf_1e6e3de76535e1bf := __obf_8f6522a0df8c3107.String()

	var __obf_65b0c23f761b6d4d, __obf_5e4d7dd61d551a43 string

	// NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
	// and you are on a 32-bit machine. Won't fix this super-edge case.
	__obf_d26c2280c86a4213 := int(__obf_19a369452cd5fbde.__obf_1ef84065a6a49f43)
	if len(__obf_1e6e3de76535e1bf) > -__obf_d26c2280c86a4213 {
		__obf_65b0c23f761b6d4d = __obf_1e6e3de76535e1bf[:len(__obf_1e6e3de76535e1bf)+__obf_d26c2280c86a4213]
		__obf_5e4d7dd61d551a43 = __obf_1e6e3de76535e1bf[len(__obf_1e6e3de76535e1bf)+__obf_d26c2280c86a4213:]
	} else {
		__obf_65b0c23f761b6d4d = "0"

		__obf_6ffe472be3a9bba1 := -__obf_d26c2280c86a4213 - len(__obf_1e6e3de76535e1bf)
		__obf_5e4d7dd61d551a43 = strings.Repeat("0", __obf_6ffe472be3a9bba1) + __obf_1e6e3de76535e1bf
	}

	if __obf_06dced229dc88d1a {
		__obf_40720372b6d983c7 := len(__obf_5e4d7dd61d551a43) - 1
		for ; __obf_40720372b6d983c7 >= 0; __obf_40720372b6d983c7-- {
			if __obf_5e4d7dd61d551a43[__obf_40720372b6d983c7] != '0' {
				break
			}
		}
		__obf_5e4d7dd61d551a43 = __obf_5e4d7dd61d551a43[:__obf_40720372b6d983c7+1]
	}

	__obf_a78756e7843a968b := __obf_65b0c23f761b6d4d
	if len(__obf_5e4d7dd61d551a43) > 0 {
		__obf_a78756e7843a968b += "." + __obf_5e4d7dd61d551a43
	}

	if __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea.Sign() < 0 {
		return "-" + __obf_a78756e7843a968b
	}

	return __obf_a78756e7843a968b
}

func (__obf_19a369452cd5fbde *Decimal) __obf_b1325bfc63b249a7() {
	if __obf_19a369452cd5fbde.__obf_6e5f77decbc21dea == nil {
		__obf_19a369452cd5fbde.__obf_6e5f77decbc21dea = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(__obf_a957bd021d745ab7 Decimal, __obf_9f2f92fe3d31076a ...Decimal) Decimal {
	__obf_6cf964822d1a9f94 := __obf_a957bd021d745ab7
	for _, __obf_244d5a292e67910b := range __obf_9f2f92fe3d31076a {
		if __obf_244d5a292e67910b.Cmp(__obf_6cf964822d1a9f94) < 0 {
			__obf_6cf964822d1a9f94 = __obf_244d5a292e67910b
		}
	}
	return __obf_6cf964822d1a9f94
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(__obf_a957bd021d745ab7 Decimal, __obf_9f2f92fe3d31076a ...Decimal) Decimal {
	__obf_6cf964822d1a9f94 := __obf_a957bd021d745ab7
	for _, __obf_244d5a292e67910b := range __obf_9f2f92fe3d31076a {
		if __obf_244d5a292e67910b.Cmp(__obf_6cf964822d1a9f94) > 0 {
			__obf_6cf964822d1a9f94 = __obf_244d5a292e67910b
		}
	}
	return __obf_6cf964822d1a9f94
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(__obf_a957bd021d745ab7 Decimal, __obf_9f2f92fe3d31076a ...Decimal) Decimal {
	__obf_f6935ddaf75a9429 := __obf_a957bd021d745ab7
	for _, __obf_244d5a292e67910b := range __obf_9f2f92fe3d31076a {
		__obf_f6935ddaf75a9429 = __obf_f6935ddaf75a9429.Add(__obf_244d5a292e67910b)
	}

	return __obf_f6935ddaf75a9429
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(__obf_a957bd021d745ab7 Decimal, __obf_9f2f92fe3d31076a ...Decimal) Decimal {
	__obf_386b7d21e8a8de4a := New(int64(len(__obf_9f2f92fe3d31076a)+1), 0)
	__obf_1882daa5a5fd23a5 := Sum(__obf_a957bd021d745ab7, __obf_9f2f92fe3d31076a...)
	return __obf_1882daa5a5fd23a5.Div(__obf_386b7d21e8a8de4a)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(__obf_c5110493d022b03b Decimal, __obf_e5610ee7ba67ce83 Decimal) (Decimal, Decimal) {
	__obf_c5110493d022b03b.__obf_b1325bfc63b249a7()
	__obf_e5610ee7ba67ce83.__obf_b1325bfc63b249a7()

	if __obf_c5110493d022b03b.__obf_1ef84065a6a49f43 < __obf_e5610ee7ba67ce83.__obf_1ef84065a6a49f43 {
		return __obf_c5110493d022b03b, __obf_e5610ee7ba67ce83.__obf_957a7b5240163386(__obf_c5110493d022b03b.__obf_1ef84065a6a49f43)
	} else if __obf_c5110493d022b03b.__obf_1ef84065a6a49f43 > __obf_e5610ee7ba67ce83.__obf_1ef84065a6a49f43 {
		return __obf_c5110493d022b03b.__obf_957a7b5240163386(__obf_e5610ee7ba67ce83.__obf_1ef84065a6a49f43), __obf_e5610ee7ba67ce83
	}

	return __obf_c5110493d022b03b, __obf_e5610ee7ba67ce83
}

func __obf_b4b0c2ff19759ebc(__obf_6e5f77decbc21dea any) (string, error) {
	var __obf_ccf6444bfd6ef423 []byte

	switch __obf_847c12b47a12274f := __obf_6e5f77decbc21dea.(type) {
	case string:
		__obf_ccf6444bfd6ef423 = []byte(__obf_847c12b47a12274f)
	case []byte:
		__obf_ccf6444bfd6ef423 = __obf_847c12b47a12274f
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", __obf_6e5f77decbc21dea, __obf_6e5f77decbc21dea)
	}

	// If the amount is quoted, strip the quotes
	if len(__obf_ccf6444bfd6ef423) > 2 && __obf_ccf6444bfd6ef423[0] == '"' && __obf_ccf6444bfd6ef423[len(__obf_ccf6444bfd6ef423)-1] == '"' {
		__obf_ccf6444bfd6ef423 = __obf_ccf6444bfd6ef423[1 : len(__obf_ccf6444bfd6ef423)-1]
	}
	return string(__obf_ccf6444bfd6ef423), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the datautil.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(__obf_19a369452cd5fbde Decimal) NullDecimal {
	return NullDecimal{
		Decimal: __obf_19a369452cd5fbde,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (__obf_19a369452cd5fbde *NullDecimal) Scan(__obf_6e5f77decbc21dea any) error {
	if __obf_6e5f77decbc21dea == nil {
		__obf_19a369452cd5fbde.Valid = false
		return nil
	}
	__obf_19a369452cd5fbde.Valid = true
	return __obf_19a369452cd5fbde.Decimal.Scan(__obf_6e5f77decbc21dea)
}

// Value implements the driver.Valuer interface for database serialization.
func (__obf_19a369452cd5fbde NullDecimal) Value() (driver.Value, error) {
	if !__obf_19a369452cd5fbde.Valid {
		return nil, nil
	}
	return __obf_19a369452cd5fbde.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (__obf_19a369452cd5fbde *NullDecimal) UnmarshalJSON(__obf_bb0607f4ee95f9e9 []byte) error {
	if string(__obf_bb0607f4ee95f9e9) == "null" {
		__obf_19a369452cd5fbde.Valid = false
		return nil
	}
	__obf_19a369452cd5fbde.Valid = true
	return __obf_19a369452cd5fbde.Decimal.UnmarshalJSON(__obf_bb0607f4ee95f9e9)
}

// MarshalJSON implements the json.Marshaler interface.
func (__obf_19a369452cd5fbde NullDecimal) MarshalJSON() ([]byte, error) {
	if !__obf_19a369452cd5fbde.Valid {
		return []byte("null"), nil
	}
	return __obf_19a369452cd5fbde.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (__obf_19a369452cd5fbde *NullDecimal) UnmarshalText(__obf_206229668f99ca6d []byte) error {
	__obf_1e6e3de76535e1bf := string(__obf_206229668f99ca6d)

	// check for empty XML or XML without body e.g., <tag></tag>
	if __obf_1e6e3de76535e1bf == "" {
		__obf_19a369452cd5fbde.Valid = false
		return nil
	}
	if __obf_68bcf986f9cdf0c0 := __obf_19a369452cd5fbde.Decimal.UnmarshalText(__obf_206229668f99ca6d); __obf_68bcf986f9cdf0c0 != nil {
		__obf_19a369452cd5fbde.Valid = false
		return __obf_68bcf986f9cdf0c0
	}
	__obf_19a369452cd5fbde.Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (__obf_19a369452cd5fbde NullDecimal) MarshalText() (__obf_206229668f99ca6d []byte, __obf_68bcf986f9cdf0c0 error) {
	if !__obf_19a369452cd5fbde.Valid {
		return []byte{}, nil
	}
	return __obf_19a369452cd5fbde.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (__obf_19a369452cd5fbde Decimal) Atan() Decimal {
	if __obf_19a369452cd5fbde.Equal(NewFromFloat(0.0)) {
		return __obf_19a369452cd5fbde
	}
	if __obf_19a369452cd5fbde.GreaterThan(NewFromFloat(0.0)) {
		return __obf_19a369452cd5fbde.__obf_e2a1a7ef085b8edd()
	}
	return __obf_19a369452cd5fbde.Neg().__obf_e2a1a7ef085b8edd().Neg()
}

func (__obf_19a369452cd5fbde Decimal) __obf_bfc57a3079197eb5() Decimal {
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
	__obf_f0a3b6dafda7d395 := __obf_19a369452cd5fbde.Mul(__obf_19a369452cd5fbde)
	__obf_225e316cf6199eff := P0.Mul(__obf_f0a3b6dafda7d395).Add(P1).Mul(__obf_f0a3b6dafda7d395).Add(P2).Mul(__obf_f0a3b6dafda7d395).Add(P3).Mul(__obf_f0a3b6dafda7d395).Add(P4).Mul(__obf_f0a3b6dafda7d395)
	__obf_32166e2f302bb9e6 := __obf_f0a3b6dafda7d395.Add(Q0).Mul(__obf_f0a3b6dafda7d395).Add(Q1).Mul(__obf_f0a3b6dafda7d395).Add(Q2).Mul(__obf_f0a3b6dafda7d395).Add(Q3).Mul(__obf_f0a3b6dafda7d395).Add(Q4)
	__obf_f0a3b6dafda7d395 = __obf_225e316cf6199eff.Div(__obf_32166e2f302bb9e6)
	__obf_f0a3b6dafda7d395 = __obf_19a369452cd5fbde.Mul(__obf_f0a3b6dafda7d395).Add(__obf_19a369452cd5fbde)
	return __obf_f0a3b6dafda7d395
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (__obf_19a369452cd5fbde Decimal) __obf_e2a1a7ef085b8edd() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)      // tan(3*pi/8)
	__obf_62e68c99e102320d := NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if __obf_19a369452cd5fbde.LessThanOrEqual(NewFromFloat(0.66)) {
		return __obf_19a369452cd5fbde.__obf_bfc57a3079197eb5()
	}
	if __obf_19a369452cd5fbde.GreaterThan(Tan3pio8) {
		return __obf_62e68c99e102320d.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(__obf_19a369452cd5fbde).__obf_bfc57a3079197eb5()).Add(Morebits)
	}
	return __obf_62e68c99e102320d.Div(NewFromFloat(4.0)).Add((__obf_19a369452cd5fbde.Sub(NewFromFloat(1.0)).Div(__obf_19a369452cd5fbde.Add(NewFromFloat(1.0)))).__obf_bfc57a3079197eb5()).Add(NewFromFloat(0.5).Mul(Morebits))
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
func (__obf_19a369452cd5fbde Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_19a369452cd5fbde.Equal(NewFromFloat(0.0)) {
		return __obf_19a369452cd5fbde
	}
	// make argument positive but save the sign
	__obf_0a87f1adff24c937 := false
	if __obf_19a369452cd5fbde.LessThan(NewFromFloat(0.0)) {
		__obf_19a369452cd5fbde = __obf_19a369452cd5fbde.Neg()
		__obf_0a87f1adff24c937 = true
	}

	__obf_13685bbda520f0c9 := __obf_19a369452cd5fbde.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_aba9ed83ce5f99ca := NewFromFloat(float64(__obf_13685bbda520f0c9)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_13685bbda520f0c9&1 == 1 {
		__obf_13685bbda520f0c9++
		__obf_aba9ed83ce5f99ca = __obf_aba9ed83ce5f99ca.Add(NewFromFloat(1.0))
	}
	__obf_13685bbda520f0c9 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_13685bbda520f0c9 > 3 {
		__obf_0a87f1adff24c937 = !__obf_0a87f1adff24c937
		__obf_13685bbda520f0c9 -= 4
	}
	__obf_f0a3b6dafda7d395 := __obf_19a369452cd5fbde.Sub(__obf_aba9ed83ce5f99ca.Mul(PI4A)).Sub(__obf_aba9ed83ce5f99ca.Mul(PI4B)).Sub(__obf_aba9ed83ce5f99ca.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_fac411bb87698bdc := __obf_f0a3b6dafda7d395.Mul(__obf_f0a3b6dafda7d395)

	if __obf_13685bbda520f0c9 == 1 || __obf_13685bbda520f0c9 == 2 {
		__obf_9c7c0223c400a310 := __obf_fac411bb87698bdc.Mul(__obf_fac411bb87698bdc).Mul(_cos[0].Mul(__obf_fac411bb87698bdc).Add(_cos[1]).Mul(__obf_fac411bb87698bdc).Add(_cos[2]).Mul(__obf_fac411bb87698bdc).Add(_cos[3]).Mul(__obf_fac411bb87698bdc).Add(_cos[4]).Mul(__obf_fac411bb87698bdc).Add(_cos[5]))
		__obf_aba9ed83ce5f99ca = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_fac411bb87698bdc)).Add(__obf_9c7c0223c400a310)
	} else {
		__obf_aba9ed83ce5f99ca = __obf_f0a3b6dafda7d395.Add(__obf_f0a3b6dafda7d395.Mul(__obf_fac411bb87698bdc).Mul(_sin[0].Mul(__obf_fac411bb87698bdc).Add(_sin[1]).Mul(__obf_fac411bb87698bdc).Add(_sin[2]).Mul(__obf_fac411bb87698bdc).Add(_sin[3]).Mul(__obf_fac411bb87698bdc).Add(_sin[4]).Mul(__obf_fac411bb87698bdc).Add(_sin[5])))
	}
	if __obf_0a87f1adff24c937 {
		__obf_aba9ed83ce5f99ca = __obf_aba9ed83ce5f99ca.Neg()
	}
	return __obf_aba9ed83ce5f99ca
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
func (__obf_19a369452cd5fbde Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	// make argument positive
	__obf_0a87f1adff24c937 := false
	if __obf_19a369452cd5fbde.LessThan(NewFromFloat(0.0)) {
		__obf_19a369452cd5fbde = __obf_19a369452cd5fbde.Neg()
	}

	__obf_13685bbda520f0c9 := __obf_19a369452cd5fbde.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_aba9ed83ce5f99ca := NewFromFloat(float64(__obf_13685bbda520f0c9)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_13685bbda520f0c9&1 == 1 {
		__obf_13685bbda520f0c9++
		__obf_aba9ed83ce5f99ca = __obf_aba9ed83ce5f99ca.Add(NewFromFloat(1.0))
	}
	__obf_13685bbda520f0c9 &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if __obf_13685bbda520f0c9 > 3 {
		__obf_0a87f1adff24c937 = !__obf_0a87f1adff24c937
		__obf_13685bbda520f0c9 -= 4
	}
	if __obf_13685bbda520f0c9 > 1 {
		__obf_0a87f1adff24c937 = !__obf_0a87f1adff24c937
	}

	__obf_f0a3b6dafda7d395 := __obf_19a369452cd5fbde.Sub(__obf_aba9ed83ce5f99ca.Mul(PI4A)).Sub(__obf_aba9ed83ce5f99ca.Mul(PI4B)).Sub(__obf_aba9ed83ce5f99ca.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_fac411bb87698bdc := __obf_f0a3b6dafda7d395.Mul(__obf_f0a3b6dafda7d395)

	if __obf_13685bbda520f0c9 == 1 || __obf_13685bbda520f0c9 == 2 {
		__obf_aba9ed83ce5f99ca = __obf_f0a3b6dafda7d395.Add(__obf_f0a3b6dafda7d395.Mul(__obf_fac411bb87698bdc).Mul(_sin[0].Mul(__obf_fac411bb87698bdc).Add(_sin[1]).Mul(__obf_fac411bb87698bdc).Add(_sin[2]).Mul(__obf_fac411bb87698bdc).Add(_sin[3]).Mul(__obf_fac411bb87698bdc).Add(_sin[4]).Mul(__obf_fac411bb87698bdc).Add(_sin[5])))
	} else {
		__obf_9c7c0223c400a310 := __obf_fac411bb87698bdc.Mul(__obf_fac411bb87698bdc).Mul(_cos[0].Mul(__obf_fac411bb87698bdc).Add(_cos[1]).Mul(__obf_fac411bb87698bdc).Add(_cos[2]).Mul(__obf_fac411bb87698bdc).Add(_cos[3]).Mul(__obf_fac411bb87698bdc).Add(_cos[4]).Mul(__obf_fac411bb87698bdc).Add(_cos[5]))
		__obf_aba9ed83ce5f99ca = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(__obf_fac411bb87698bdc)).Add(__obf_9c7c0223c400a310)
	}
	if __obf_0a87f1adff24c937 {
		__obf_aba9ed83ce5f99ca = __obf_aba9ed83ce5f99ca.Neg()
	}
	return __obf_aba9ed83ce5f99ca
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
func (__obf_19a369452cd5fbde Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if __obf_19a369452cd5fbde.Equal(NewFromFloat(0.0)) {
		return __obf_19a369452cd5fbde
	}

	// make argument positive but save the sign
	__obf_0a87f1adff24c937 := false
	if __obf_19a369452cd5fbde.LessThan(NewFromFloat(0.0)) {
		__obf_19a369452cd5fbde = __obf_19a369452cd5fbde.Neg()
		__obf_0a87f1adff24c937 = true
	}

	__obf_13685bbda520f0c9 := __obf_19a369452cd5fbde.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	__obf_aba9ed83ce5f99ca := NewFromFloat(float64(__obf_13685bbda520f0c9)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if __obf_13685bbda520f0c9&1 == 1 {
		__obf_13685bbda520f0c9++
		__obf_aba9ed83ce5f99ca = __obf_aba9ed83ce5f99ca.Add(NewFromFloat(1.0))
	}

	__obf_f0a3b6dafda7d395 := __obf_19a369452cd5fbde.Sub(__obf_aba9ed83ce5f99ca.Mul(PI4A)).Sub(__obf_aba9ed83ce5f99ca.Mul(PI4B)).Sub(__obf_aba9ed83ce5f99ca.Mul(PI4C)) // Extended precision modular arithmetic
	__obf_fac411bb87698bdc := __obf_f0a3b6dafda7d395.Mul(__obf_f0a3b6dafda7d395)

	if __obf_fac411bb87698bdc.GreaterThan(NewFromFloat(1e-14)) {
		__obf_9c7c0223c400a310 := __obf_fac411bb87698bdc.Mul(_tanP[0].Mul(__obf_fac411bb87698bdc).Add(_tanP[1]).Mul(__obf_fac411bb87698bdc).Add(_tanP[2]))
		__obf_cfc89f6d692b6000 := __obf_fac411bb87698bdc.Add(_tanQ[1]).Mul(__obf_fac411bb87698bdc).Add(_tanQ[2]).Mul(__obf_fac411bb87698bdc).Add(_tanQ[3]).Mul(__obf_fac411bb87698bdc).Add(_tanQ[4])
		__obf_aba9ed83ce5f99ca = __obf_f0a3b6dafda7d395.Add(__obf_f0a3b6dafda7d395.Mul(__obf_9c7c0223c400a310.Div(__obf_cfc89f6d692b6000)))
	} else {
		__obf_aba9ed83ce5f99ca = __obf_f0a3b6dafda7d395
	}
	if __obf_13685bbda520f0c9&2 == 2 {
		__obf_aba9ed83ce5f99ca = NewFromFloat(-1.0).Div(__obf_aba9ed83ce5f99ca)
	}
	if __obf_0a87f1adff24c937 {
		__obf_aba9ed83ce5f99ca = __obf_aba9ed83ce5f99ca.Neg()
	}
	return __obf_aba9ed83ce5f99ca
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

type __obf_ae16adf734cfe1aa struct {
	__obf_19a369452cd5fbde [800]byte // digits, big-endian representation
	__obf_04074ea54302ce17 int       // number of digits used
	__obf_45565779317643f9 int       // decimal point
	__obf_d12c1496d23e3503 bool      // negative flag
	__obf_a9f8af80e19992b6 bool      // discarded nonzero digits beyond d[:nd]
}

func (__obf_364e38cb96f1c266 *__obf_ae16adf734cfe1aa) String() string {
	__obf_3edab89cf83bab1c := 10 + __obf_364e38cb96f1c266.__obf_04074ea54302ce17
	if __obf_364e38cb96f1c266.__obf_45565779317643f9 > 0 {
		__obf_3edab89cf83bab1c += __obf_364e38cb96f1c266.__obf_45565779317643f9
	}
	if __obf_364e38cb96f1c266.__obf_45565779317643f9 < 0 {
		__obf_3edab89cf83bab1c += -__obf_364e38cb96f1c266.__obf_45565779317643f9
	}

	__obf_27adff64d3e60f06 := make([]byte, __obf_3edab89cf83bab1c)
	__obf_9c7c0223c400a310 := 0
	switch {
	case __obf_364e38cb96f1c266.__obf_04074ea54302ce17 == 0:
		return "0"

	case __obf_364e38cb96f1c266.__obf_45565779317643f9 <= 0:
		// zeros fill space between decimal point and digits
		__obf_27adff64d3e60f06[__obf_9c7c0223c400a310] = '0'
		__obf_9c7c0223c400a310++
		__obf_27adff64d3e60f06[__obf_9c7c0223c400a310] = '.'
		__obf_9c7c0223c400a310++
		__obf_9c7c0223c400a310 += __obf_e2f74946f0b4300c(__obf_27adff64d3e60f06[__obf_9c7c0223c400a310 : __obf_9c7c0223c400a310+-__obf_364e38cb96f1c266.__obf_45565779317643f9])
		__obf_9c7c0223c400a310 += copy(__obf_27adff64d3e60f06[__obf_9c7c0223c400a310:], __obf_364e38cb96f1c266.__obf_19a369452cd5fbde[0:__obf_364e38cb96f1c266.__obf_04074ea54302ce17])

	case __obf_364e38cb96f1c266.__obf_45565779317643f9 < __obf_364e38cb96f1c266.__obf_04074ea54302ce17:
		// decimal point in middle of digits
		__obf_9c7c0223c400a310 += copy(__obf_27adff64d3e60f06[__obf_9c7c0223c400a310:], __obf_364e38cb96f1c266.__obf_19a369452cd5fbde[0:__obf_364e38cb96f1c266.__obf_45565779317643f9])
		__obf_27adff64d3e60f06[__obf_9c7c0223c400a310] = '.'
		__obf_9c7c0223c400a310++
		__obf_9c7c0223c400a310 += copy(__obf_27adff64d3e60f06[__obf_9c7c0223c400a310:], __obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_364e38cb96f1c266.__obf_45565779317643f9:__obf_364e38cb96f1c266.__obf_04074ea54302ce17])

	default:
		// zeros fill space between digits and decimal point
		__obf_9c7c0223c400a310 += copy(__obf_27adff64d3e60f06[__obf_9c7c0223c400a310:], __obf_364e38cb96f1c266.__obf_19a369452cd5fbde[0:__obf_364e38cb96f1c266.__obf_04074ea54302ce17])
		__obf_9c7c0223c400a310 += __obf_e2f74946f0b4300c(__obf_27adff64d3e60f06[__obf_9c7c0223c400a310 : __obf_9c7c0223c400a310+__obf_364e38cb96f1c266.__obf_45565779317643f9-__obf_364e38cb96f1c266.__obf_04074ea54302ce17])
	}
	return string(__obf_27adff64d3e60f06[0:__obf_9c7c0223c400a310])
}

func __obf_e2f74946f0b4300c(__obf_1eac2c84174deaf7 []byte) int {
	for __obf_40720372b6d983c7 := range __obf_1eac2c84174deaf7 {
		__obf_1eac2c84174deaf7[__obf_40720372b6d983c7] = '0'
	}
	return len(__obf_1eac2c84174deaf7)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func __obf_0e49a0dae06ff7fc(__obf_364e38cb96f1c266 *__obf_ae16adf734cfe1aa) {
	for __obf_364e38cb96f1c266.__obf_04074ea54302ce17 > 0 && __obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_364e38cb96f1c266.__obf_04074ea54302ce17-1] == '0' {
		__obf_364e38cb96f1c266.__obf_04074ea54302ce17--
	}
	if __obf_364e38cb96f1c266.__obf_04074ea54302ce17 == 0 {
		__obf_364e38cb96f1c266.__obf_45565779317643f9 = 0
	}
}

// Assign v to a.
func (__obf_364e38cb96f1c266 *__obf_ae16adf734cfe1aa) Assign(__obf_847c12b47a12274f uint64) {
	var __obf_27adff64d3e60f06 [24]byte

	// Write reversed decimal in buf.
	__obf_3edab89cf83bab1c := 0
	for __obf_847c12b47a12274f > 0 {
		__obf_8cf82387c82987e3 := __obf_847c12b47a12274f / 10
		__obf_847c12b47a12274f -= 10 * __obf_8cf82387c82987e3
		__obf_27adff64d3e60f06[__obf_3edab89cf83bab1c] = byte(__obf_847c12b47a12274f + '0')
		__obf_3edab89cf83bab1c++
		__obf_847c12b47a12274f = __obf_8cf82387c82987e3
	}

	// Reverse again to produce forward decimal in a.d.
	__obf_364e38cb96f1c266.__obf_04074ea54302ce17 = 0
	for __obf_3edab89cf83bab1c--; __obf_3edab89cf83bab1c >= 0; __obf_3edab89cf83bab1c-- {
		__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_364e38cb96f1c266.__obf_04074ea54302ce17] = __obf_27adff64d3e60f06[__obf_3edab89cf83bab1c]
		__obf_364e38cb96f1c266.__obf_04074ea54302ce17++
	}
	__obf_364e38cb96f1c266.__obf_45565779317643f9 = __obf_364e38cb96f1c266.__obf_04074ea54302ce17
	__obf_0e49a0dae06ff7fc(__obf_364e38cb96f1c266)
}

// Maximum shift that we can do in one pass without overflow.
// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
const __obf_072540258d9bc049 = 32 << (^uint(0) >> 63)
const __obf_19a0abec1376a547 = __obf_072540258d9bc049 - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_5a5b314cdd262b1a(__obf_364e38cb96f1c266 *__obf_ae16adf734cfe1aa, __obf_ec401ce06db88a9f uint) {
	__obf_0feccac4ef34d0a4 := 0 // read pointer
	__obf_9c7c0223c400a310 := 0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var __obf_3edab89cf83bab1c uint
	for ; __obf_3edab89cf83bab1c>>__obf_ec401ce06db88a9f == 0; __obf_0feccac4ef34d0a4++ {
		if __obf_0feccac4ef34d0a4 >= __obf_364e38cb96f1c266.__obf_04074ea54302ce17 {
			if __obf_3edab89cf83bab1c == 0 {
				// a == 0; shouldn't get here, but handle anyway.
				__obf_364e38cb96f1c266.__obf_04074ea54302ce17 = 0
				return
			}
			for __obf_3edab89cf83bab1c>>__obf_ec401ce06db88a9f == 0 {
				__obf_3edab89cf83bab1c = __obf_3edab89cf83bab1c * 10
				__obf_0feccac4ef34d0a4++
			}
			break
		}
		__obf_6dfba1dc52bc2b42 := uint(__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_0feccac4ef34d0a4])
		__obf_3edab89cf83bab1c = __obf_3edab89cf83bab1c*10 + __obf_6dfba1dc52bc2b42 - '0'
	}
	__obf_364e38cb96f1c266.__obf_45565779317643f9 -= __obf_0feccac4ef34d0a4 - 1

	var __obf_2960a71b685f1638 uint = (1 << __obf_ec401ce06db88a9f) - 1

	// Pick up a digit, put down a digit.
	for ; __obf_0feccac4ef34d0a4 < __obf_364e38cb96f1c266.__obf_04074ea54302ce17; __obf_0feccac4ef34d0a4++ {
		__obf_6dfba1dc52bc2b42 := uint(__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_0feccac4ef34d0a4])
		__obf_742bea9a421ff878 := __obf_3edab89cf83bab1c >> __obf_ec401ce06db88a9f
		__obf_3edab89cf83bab1c &= __obf_2960a71b685f1638
		__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_9c7c0223c400a310] = byte(__obf_742bea9a421ff878 + '0')
		__obf_9c7c0223c400a310++
		__obf_3edab89cf83bab1c = __obf_3edab89cf83bab1c*10 + __obf_6dfba1dc52bc2b42 - '0'
	}

	// Put down extra digits.
	for __obf_3edab89cf83bab1c > 0 {
		__obf_742bea9a421ff878 := __obf_3edab89cf83bab1c >> __obf_ec401ce06db88a9f
		__obf_3edab89cf83bab1c &= __obf_2960a71b685f1638
		if __obf_9c7c0223c400a310 < len(__obf_364e38cb96f1c266.__obf_19a369452cd5fbde) {
			__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_9c7c0223c400a310] = byte(__obf_742bea9a421ff878 + '0')
			__obf_9c7c0223c400a310++
		} else if __obf_742bea9a421ff878 > 0 {
			__obf_364e38cb96f1c266.__obf_a9f8af80e19992b6 = true
		}
		__obf_3edab89cf83bab1c = __obf_3edab89cf83bab1c * 10
	}

	__obf_364e38cb96f1c266.__obf_04074ea54302ce17 = __obf_9c7c0223c400a310
	__obf_0e49a0dae06ff7fc(__obf_364e38cb96f1c266)
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

type __obf_2312a386ebb40631 struct {
	__obf_c0eae4d6d233536d int    // number of new digits
	__obf_faa129282fbe4c56 string // minus one digit if original < a.
}

var __obf_e664b6f58cdf7a41 = []__obf_2312a386ebb40631{
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
func __obf_cb98672515456c4f(__obf_2ea797d6029d6922 []byte, __obf_1d21ccd15f74d9eb string) bool {
	for __obf_40720372b6d983c7 := 0; __obf_40720372b6d983c7 < len(__obf_1d21ccd15f74d9eb); __obf_40720372b6d983c7++ {
		if __obf_40720372b6d983c7 >= len(__obf_2ea797d6029d6922) {
			return true
		}
		if __obf_2ea797d6029d6922[__obf_40720372b6d983c7] != __obf_1d21ccd15f74d9eb[__obf_40720372b6d983c7] {
			return __obf_2ea797d6029d6922[__obf_40720372b6d983c7] < __obf_1d21ccd15f74d9eb[__obf_40720372b6d983c7]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func __obf_bfe89c7811ef030d(__obf_364e38cb96f1c266 *__obf_ae16adf734cfe1aa, __obf_ec401ce06db88a9f uint) {
	__obf_c0eae4d6d233536d := __obf_e664b6f58cdf7a41[__obf_ec401ce06db88a9f].__obf_c0eae4d6d233536d
	if __obf_cb98672515456c4f(__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[0:__obf_364e38cb96f1c266.__obf_04074ea54302ce17], __obf_e664b6f58cdf7a41[__obf_ec401ce06db88a9f].__obf_faa129282fbe4c56) {
		__obf_c0eae4d6d233536d--
	}

	__obf_0feccac4ef34d0a4 := __obf_364e38cb96f1c266.__obf_04074ea54302ce17                          // read index
	__obf_9c7c0223c400a310 := __obf_364e38cb96f1c266.__obf_04074ea54302ce17 + __obf_c0eae4d6d233536d // write index

	// Pick up a digit, put down a digit.
	var __obf_3edab89cf83bab1c uint
	for __obf_0feccac4ef34d0a4--; __obf_0feccac4ef34d0a4 >= 0; __obf_0feccac4ef34d0a4-- {
		__obf_3edab89cf83bab1c += (uint(__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_0feccac4ef34d0a4]) - '0') << __obf_ec401ce06db88a9f
		__obf_e2ed4a0db8d049a4 := __obf_3edab89cf83bab1c / 10
		__obf_b64a2b410302a63f := __obf_3edab89cf83bab1c - 10*__obf_e2ed4a0db8d049a4
		__obf_9c7c0223c400a310--
		if __obf_9c7c0223c400a310 < len(__obf_364e38cb96f1c266.__obf_19a369452cd5fbde) {
			__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_9c7c0223c400a310] = byte(__obf_b64a2b410302a63f + '0')
		} else if __obf_b64a2b410302a63f != 0 {
			__obf_364e38cb96f1c266.__obf_a9f8af80e19992b6 = true
		}
		__obf_3edab89cf83bab1c = __obf_e2ed4a0db8d049a4
	}

	// Put down extra digits.
	for __obf_3edab89cf83bab1c > 0 {
		__obf_e2ed4a0db8d049a4 := __obf_3edab89cf83bab1c / 10
		__obf_b64a2b410302a63f := __obf_3edab89cf83bab1c - 10*__obf_e2ed4a0db8d049a4
		__obf_9c7c0223c400a310--
		if __obf_9c7c0223c400a310 < len(__obf_364e38cb96f1c266.__obf_19a369452cd5fbde) {
			__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_9c7c0223c400a310] = byte(__obf_b64a2b410302a63f + '0')
		} else if __obf_b64a2b410302a63f != 0 {
			__obf_364e38cb96f1c266.__obf_a9f8af80e19992b6 = true
		}
		__obf_3edab89cf83bab1c = __obf_e2ed4a0db8d049a4
	}

	__obf_364e38cb96f1c266.__obf_04074ea54302ce17 += __obf_c0eae4d6d233536d
	if __obf_364e38cb96f1c266.__obf_04074ea54302ce17 >= len(__obf_364e38cb96f1c266.__obf_19a369452cd5fbde) {
		__obf_364e38cb96f1c266.__obf_04074ea54302ce17 = len(__obf_364e38cb96f1c266.__obf_19a369452cd5fbde)
	}
	__obf_364e38cb96f1c266.__obf_45565779317643f9 += __obf_c0eae4d6d233536d
	__obf_0e49a0dae06ff7fc(__obf_364e38cb96f1c266)
}

// Binary shift left (k > 0) or right (k < 0).
func (__obf_364e38cb96f1c266 *__obf_ae16adf734cfe1aa) Shift(__obf_ec401ce06db88a9f int) {
	switch {
	case __obf_364e38cb96f1c266.__obf_04074ea54302ce17 == 0:
		// nothing to do: a == 0
	case __obf_ec401ce06db88a9f > 0:
		for __obf_ec401ce06db88a9f > __obf_19a0abec1376a547 {
			__obf_bfe89c7811ef030d(__obf_364e38cb96f1c266, __obf_19a0abec1376a547)
			__obf_ec401ce06db88a9f -= __obf_19a0abec1376a547
		}
		__obf_bfe89c7811ef030d(__obf_364e38cb96f1c266, uint(__obf_ec401ce06db88a9f))
	case __obf_ec401ce06db88a9f < 0:
		for __obf_ec401ce06db88a9f < -__obf_19a0abec1376a547 {
			__obf_5a5b314cdd262b1a(__obf_364e38cb96f1c266, __obf_19a0abec1376a547)
			__obf_ec401ce06db88a9f += __obf_19a0abec1376a547
		}
		__obf_5a5b314cdd262b1a(__obf_364e38cb96f1c266, uint(-__obf_ec401ce06db88a9f))
	}
}

// If we chop a at nd digits, should we round up?
func __obf_1f47eb3e5c6a4efc(__obf_364e38cb96f1c266 *__obf_ae16adf734cfe1aa, __obf_04074ea54302ce17 int) bool {
	if __obf_04074ea54302ce17 < 0 || __obf_04074ea54302ce17 >= __obf_364e38cb96f1c266.__obf_04074ea54302ce17 {
		return false
	}
	if __obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_04074ea54302ce17] == '5' && __obf_04074ea54302ce17+1 == __obf_364e38cb96f1c266.__obf_04074ea54302ce17 { // exactly halfway - round to even
		// if we truncated, a little higher than what's recorded - always round up
		if __obf_364e38cb96f1c266.__obf_a9f8af80e19992b6 {
			return true
		}
		return __obf_04074ea54302ce17 > 0 && (__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_04074ea54302ce17-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return __obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_04074ea54302ce17] >= '5'
}

// Round a to nd digits (or fewer).
// If nd is zero, it means we're rounding
// just to the left of the digits, as in
// 0.09 -> 0.1.
func (__obf_364e38cb96f1c266 *__obf_ae16adf734cfe1aa) Round(__obf_04074ea54302ce17 int) {
	if __obf_04074ea54302ce17 < 0 || __obf_04074ea54302ce17 >= __obf_364e38cb96f1c266.__obf_04074ea54302ce17 {
		return
	}
	if __obf_1f47eb3e5c6a4efc(__obf_364e38cb96f1c266, __obf_04074ea54302ce17) {
		__obf_364e38cb96f1c266.RoundUp(__obf_04074ea54302ce17)
	} else {
		__obf_364e38cb96f1c266.RoundDown(__obf_04074ea54302ce17)
	}
}

// Round a down to nd digits (or fewer).
func (__obf_364e38cb96f1c266 *__obf_ae16adf734cfe1aa) RoundDown(__obf_04074ea54302ce17 int) {
	if __obf_04074ea54302ce17 < 0 || __obf_04074ea54302ce17 >= __obf_364e38cb96f1c266.__obf_04074ea54302ce17 {
		return
	}
	__obf_364e38cb96f1c266.__obf_04074ea54302ce17 = __obf_04074ea54302ce17
	__obf_0e49a0dae06ff7fc(__obf_364e38cb96f1c266)
}

// Round a up to nd digits (or fewer).
func (__obf_364e38cb96f1c266 *__obf_ae16adf734cfe1aa) RoundUp(__obf_04074ea54302ce17 int) {
	if __obf_04074ea54302ce17 < 0 || __obf_04074ea54302ce17 >= __obf_364e38cb96f1c266.__obf_04074ea54302ce17 {
		return
	}

	// round up
	for __obf_40720372b6d983c7 := __obf_04074ea54302ce17 - 1; __obf_40720372b6d983c7 >= 0; __obf_40720372b6d983c7-- {
		__obf_6dfba1dc52bc2b42 := __obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_40720372b6d983c7]
		if __obf_6dfba1dc52bc2b42 < '9' { // can stop after this digit
			__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_40720372b6d983c7]++
			__obf_364e38cb96f1c266.__obf_04074ea54302ce17 = __obf_40720372b6d983c7 + 1
			return
		}
	}

	// Number is all 9s.
	// Change to single 1 with adjusted decimal point.
	__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[0] = '1'
	__obf_364e38cb96f1c266.__obf_04074ea54302ce17 = 1
	__obf_364e38cb96f1c266.__obf_45565779317643f9++
}

// Extract integer part, rounded appropriately.
// No guarantees about overflow.
func (__obf_364e38cb96f1c266 *__obf_ae16adf734cfe1aa) RoundedInteger() uint64 {
	if __obf_364e38cb96f1c266.__obf_45565779317643f9 > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var __obf_40720372b6d983c7 int
	__obf_3edab89cf83bab1c := uint64(0)
	for __obf_40720372b6d983c7 = 0; __obf_40720372b6d983c7 < __obf_364e38cb96f1c266.__obf_45565779317643f9 && __obf_40720372b6d983c7 < __obf_364e38cb96f1c266.__obf_04074ea54302ce17; __obf_40720372b6d983c7++ {
		__obf_3edab89cf83bab1c = __obf_3edab89cf83bab1c*10 + uint64(__obf_364e38cb96f1c266.__obf_19a369452cd5fbde[__obf_40720372b6d983c7]-'0')
	}
	for ; __obf_40720372b6d983c7 < __obf_364e38cb96f1c266.__obf_45565779317643f9; __obf_40720372b6d983c7++ {
		__obf_3edab89cf83bab1c *= 10
	}
	if __obf_1f47eb3e5c6a4efc(__obf_364e38cb96f1c266, __obf_364e38cb96f1c266.__obf_45565779317643f9) {
		__obf_3edab89cf83bab1c++
	}
	return __obf_3edab89cf83bab1c
}
