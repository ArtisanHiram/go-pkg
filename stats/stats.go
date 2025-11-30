package __obf_5865d4dc61043b11

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_b0bc5503bc2661cc []float64) (max float64, __obf_786467db099882d9 error) {
	__obf_5b6b1c8b9be46734 := len(__obf_b0bc5503bc2661cc)

	// Return an error if there are no numbers
	if __obf_5b6b1c8b9be46734 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_b0bc5503bc2661cc[0]

	// Loop and replace higher values
	for __obf_ccfe5db1ac8ff7e8 := 1; __obf_ccfe5db1ac8ff7e8 < __obf_5b6b1c8b9be46734; __obf_ccfe5db1ac8ff7e8++ {
		if __obf_b0bc5503bc2661cc[__obf_ccfe5db1ac8ff7e8] > max {
			max = __obf_b0bc5503bc2661cc[__obf_ccfe5db1ac8ff7e8]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_b0bc5503bc2661cc []float64) (min float64, __obf_786467db099882d9 error) {
	__obf_5b6b1c8b9be46734 := // Get the count of numbers in the slice
		len(__obf_b0bc5503bc2661cc)

	// Return an error if there are no numbers
	if __obf_5b6b1c8b9be46734 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_b0bc5503bc2661cc[0]

	// Iterate until done checking for a lower value
	for __obf_ccfe5db1ac8ff7e8 := 1; __obf_ccfe5db1ac8ff7e8 < __obf_5b6b1c8b9be46734; __obf_ccfe5db1ac8ff7e8++ {
		if __obf_b0bc5503bc2661cc[__obf_ccfe5db1ac8ff7e8] < min {
			min = __obf_b0bc5503bc2661cc[__obf_ccfe5db1ac8ff7e8]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_b0bc5503bc2661cc float64, __obf_a2481e66a215b354 int) (__obf_b076a52f280bbb81 float64, __obf_786467db099882d9 error) {

	// If the float is not a number
	if math.IsNaN(__obf_b0bc5503bc2661cc) {
		return math.NaN(), ErrNaN
	}
	__obf_c9e77a60cebb022b := // Find out the actual sign and correct the input for later
		1.0
	if __obf_b0bc5503bc2661cc < 0 {
		__obf_c9e77a60cebb022b = -1
		__obf_b0bc5503bc2661cc *= -1
	}
	__obf_53ad2f411acf035e := // Use the places arg to get the amount of precision wanted
		math.Pow(10, float64(__obf_a2481e66a215b354))
	__obf_e8280063f8991632 := // Find the decimal place we are looking to round
		__obf_b0bc5503bc2661cc * __obf_53ad2f411acf035e

	// Get the actual decimal number as a fraction to be compared
	_, __obf_d44e02ab73cb16a1 := math.Modf(__obf_e8280063f8991632)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_d44e02ab73cb16a1 >= 0.5 {
		__obf_b076a52f280bbb81 = math.Ceil(__obf_e8280063f8991632)
	} else {
		__obf_b076a52f280bbb81 = math.Floor(__obf_e8280063f8991632)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_b076a52f280bbb81 / __obf_53ad2f411acf035e * __obf_c9e77a60cebb022b, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_b0bc5503bc2661cc []float64) (float64, error) {
	__obf_5b6b1c8b9be46734 := len(__obf_b0bc5503bc2661cc)
	if __obf_5b6b1c8b9be46734 == 0 {
		return math.NaN(), ErrEmptyInput
	}
	__obf_e18ce02c318caa61, _ := Sum(__obf_b0bc5503bc2661cc)

	return __obf_e18ce02c318caa61 / float64(__obf_5b6b1c8b9be46734), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_b0bc5503bc2661cc []float64) (__obf_797b078d577c8f53 float64, __obf_786467db099882d9 error) {
	__obf_7f20ce810423128d := // Start by sorting a copy of the slice
		__obf_713a257c97207840(__obf_b0bc5503bc2661cc)
	__obf_5b6b1c8b9be46734 := // No math is needed if there are no numbers
		// For even numbers we add the two middle numbers
		// and divide by two using the mean function above
		// For odd numbers we just use the middle number
		len(__obf_7f20ce810423128d)
	if __obf_5b6b1c8b9be46734 == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_5b6b1c8b9be46734%2 == 0 {
		__obf_797b078d577c8f53, _ = Mean(__obf_7f20ce810423128d[__obf_5b6b1c8b9be46734/2-1 : __obf_5b6b1c8b9be46734/2+1])
	} else {
		__obf_797b078d577c8f53 = __obf_7f20ce810423128d[__obf_5b6b1c8b9be46734/2]
	}

	return __obf_797b078d577c8f53, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_b0bc5503bc2661cc []float64) (__obf_e18ce02c318caa61 float64, __obf_786467db099882d9 error) {

	if len(__obf_b0bc5503bc2661cc) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_f80a47d9b46e761f := range __obf_b0bc5503bc2661cc {
		__obf_e18ce02c318caa61 += __obf_f80a47d9b46e761f
	}

	return __obf_e18ce02c318caa61, nil
}

// 得到一个区间map
func RangeCount(__obf_be713ad08f43086a float64, __obf_01b126c9cdd8cf11 int, __obf_b0bc5503bc2661cc []float64) (__obf_c15216f611507214 map[string]int, __obf_786467db099882d9 error) {
	if __obf_01b126c9cdd8cf11 <= 0 {
		return nil, ErrIllegalNum
	}
	__obf_91a271e1531ed183 := // total, _ := Max(input)
		__obf_f602652101083886(__obf_be713ad08f43086a / float64(__obf_01b126c9cdd8cf11))
	__obf_bb74de268bf49833 := func(__obf_25bca95541d417ad float64) string {
		for __obf_ccfe5db1ac8ff7e8 := 1; __obf_ccfe5db1ac8ff7e8 < __obf_01b126c9cdd8cf11; __obf_ccfe5db1ac8ff7e8++ {
			if __obf_25bca95541d417ad < float64(__obf_ccfe5db1ac8ff7e8*__obf_91a271e1531ed183) {
				return fmt.Sprintf("%d~%d", (__obf_ccfe5db1ac8ff7e8-1)*__obf_91a271e1531ed183, __obf_ccfe5db1ac8ff7e8*__obf_91a271e1531ed183)
			}
		}
		return fmt.Sprintf("%d~", (__obf_01b126c9cdd8cf11-1)*__obf_91a271e1531ed183)
	}
	__obf_c15216f611507214 = make(map[string]int)
	for _, __obf_50f3cb13f870bcff := range __obf_b0bc5503bc2661cc {
		__obf_c15216f611507214[__obf_bb74de268bf49833(__obf_50f3cb13f870bcff)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_ccfe5db1ac8ff7e8 uint) (__obf_7f20ce810423128d uint) {
	if __obf_ccfe5db1ac8ff7e8 == 0 {
		return 1
	}
	for __obf_ccfe5db1ac8ff7e8 != 0 {
		__obf_ccfe5db1ac8ff7e8 /= 10
		__obf_7f20ce810423128d++
	}
	return __obf_7f20ce810423128d
}

// 计算阶乘: n!
func Factorial(__obf_f80a47d9b46e761f int) (__obf_f939a8ba3b59ac8f int) {
	if __obf_f80a47d9b46e761f > 0 {
		__obf_f939a8ba3b59ac8f = 1
		for __obf_ccfe5db1ac8ff7e8 := 1; __obf_ccfe5db1ac8ff7e8 <= __obf_f80a47d9b46e761f; __obf_ccfe5db1ac8ff7e8++ {
			__obf_f939a8ba3b59ac8f *= __obf_ccfe5db1ac8ff7e8
		}
	}

	return __obf_f939a8ba3b59ac8f
}

// base 的 exp 次方
func IntPow(__obf_b9ad1c2a6153ff5c, __obf_671685fa69cc2a55 int) int {
	__obf_f72f204a6da9cc48 := 1
	for {
		if __obf_671685fa69cc2a55&1 == 1 {
			__obf_f72f204a6da9cc48 *= __obf_b9ad1c2a6153ff5c
		}
		__obf_671685fa69cc2a55 >>= 1
		if __obf_671685fa69cc2a55 == 0 {
			break
		}
		__obf_b9ad1c2a6153ff5c *= __obf_b9ad1c2a6153ff5c
	}

	return __obf_f72f204a6da9cc48
}
