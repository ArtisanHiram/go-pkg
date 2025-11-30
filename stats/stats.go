package __obf_bae661114bf2a601

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_c5a1fa438321e828 []float64) (max float64, __obf_af23b0de43c08000 error) {
	__obf_48489c395a33a241 := len(__obf_c5a1fa438321e828)

	// Return an error if there are no numbers
	if __obf_48489c395a33a241 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_c5a1fa438321e828[0]

	// Loop and replace higher values
	for __obf_91e29e63b8ab7e94 := 1; __obf_91e29e63b8ab7e94 < __obf_48489c395a33a241; __obf_91e29e63b8ab7e94++ {
		if __obf_c5a1fa438321e828[__obf_91e29e63b8ab7e94] > max {
			max = __obf_c5a1fa438321e828[__obf_91e29e63b8ab7e94]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_c5a1fa438321e828 []float64) (min float64, __obf_af23b0de43c08000 error) {
	__obf_48489c395a33a241 := // Get the count of numbers in the slice
		len(__obf_c5a1fa438321e828)

	// Return an error if there are no numbers
	if __obf_48489c395a33a241 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_c5a1fa438321e828[0]

	// Iterate until done checking for a lower value
	for __obf_91e29e63b8ab7e94 := 1; __obf_91e29e63b8ab7e94 < __obf_48489c395a33a241; __obf_91e29e63b8ab7e94++ {
		if __obf_c5a1fa438321e828[__obf_91e29e63b8ab7e94] < min {
			min = __obf_c5a1fa438321e828[__obf_91e29e63b8ab7e94]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_c5a1fa438321e828 float64, __obf_c5993402abb644f4 int) (__obf_0276cf85cae798e4 float64, __obf_af23b0de43c08000 error) {

	// If the float is not a number
	if math.IsNaN(__obf_c5a1fa438321e828) {
		return math.NaN(), ErrNaN
	}
	__obf_91a4a84501a5099c := // Find out the actual sign and correct the input for later
		1.0
	if __obf_c5a1fa438321e828 < 0 {
		__obf_91a4a84501a5099c = -1
		__obf_c5a1fa438321e828 *= -1
	}
	__obf_1b2a35f18142ef55 := // Use the places arg to get the amount of precision wanted
		math.Pow(10, float64(__obf_c5993402abb644f4))
	__obf_08b9a23183032a85 := // Find the decimal place we are looking to round
		__obf_c5a1fa438321e828 * __obf_1b2a35f18142ef55

	// Get the actual decimal number as a fraction to be compared
	_, __obf_f6010eab5dc27a4f := math.Modf(__obf_08b9a23183032a85)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_f6010eab5dc27a4f >= 0.5 {
		__obf_0276cf85cae798e4 = math.Ceil(__obf_08b9a23183032a85)
	} else {
		__obf_0276cf85cae798e4 = math.Floor(__obf_08b9a23183032a85)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_0276cf85cae798e4 / __obf_1b2a35f18142ef55 * __obf_91a4a84501a5099c, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_c5a1fa438321e828 []float64) (float64, error) {
	__obf_48489c395a33a241 := len(__obf_c5a1fa438321e828)
	if __obf_48489c395a33a241 == 0 {
		return math.NaN(), ErrEmptyInput
	}
	__obf_b3c13c6c68804fa0, _ := Sum(__obf_c5a1fa438321e828)

	return __obf_b3c13c6c68804fa0 / float64(__obf_48489c395a33a241), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_c5a1fa438321e828 []float64) (__obf_8c07381fa4189c1b float64, __obf_af23b0de43c08000 error) {
	__obf_bb96b0cd6c8fa7bc := // Start by sorting a copy of the slice
		__obf_083545a4cd904204(__obf_c5a1fa438321e828)
	__obf_48489c395a33a241 := // No math is needed if there are no numbers
		// For even numbers we add the two middle numbers
		// and divide by two using the mean function above
		// For odd numbers we just use the middle number
		len(__obf_bb96b0cd6c8fa7bc)
	if __obf_48489c395a33a241 == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_48489c395a33a241%2 == 0 {
		__obf_8c07381fa4189c1b, _ = Mean(__obf_bb96b0cd6c8fa7bc[__obf_48489c395a33a241/2-1 : __obf_48489c395a33a241/2+1])
	} else {
		__obf_8c07381fa4189c1b = __obf_bb96b0cd6c8fa7bc[__obf_48489c395a33a241/2]
	}

	return __obf_8c07381fa4189c1b, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_c5a1fa438321e828 []float64) (__obf_b3c13c6c68804fa0 float64, __obf_af23b0de43c08000 error) {

	if len(__obf_c5a1fa438321e828) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_cea35a19fa824f05 := range __obf_c5a1fa438321e828 {
		__obf_b3c13c6c68804fa0 += __obf_cea35a19fa824f05
	}

	return __obf_b3c13c6c68804fa0, nil
}

// 得到一个区间map
func RangeCount(__obf_afeb6a6e72874971 float64, __obf_0ebf159da3cfb2ab int, __obf_c5a1fa438321e828 []float64) (__obf_ca4f88e82278c739 map[string]int, __obf_af23b0de43c08000 error) {
	if __obf_0ebf159da3cfb2ab <= 0 {
		return nil, ErrIllegalNum
	}
	__obf_1429cb6a843c511b := // total, _ := Max(input)
		__obf_91d2ecb5786b6fc2(__obf_afeb6a6e72874971 / float64(__obf_0ebf159da3cfb2ab))
	__obf_2f32f638bd550f86 := func(__obf_5564062d548a8fe7 float64) string {
		for __obf_91e29e63b8ab7e94 := 1; __obf_91e29e63b8ab7e94 < __obf_0ebf159da3cfb2ab; __obf_91e29e63b8ab7e94++ {
			if __obf_5564062d548a8fe7 < float64(__obf_91e29e63b8ab7e94*__obf_1429cb6a843c511b) {
				return fmt.Sprintf("%d~%d", (__obf_91e29e63b8ab7e94-1)*__obf_1429cb6a843c511b, __obf_91e29e63b8ab7e94*__obf_1429cb6a843c511b)
			}
		}
		return fmt.Sprintf("%d~", (__obf_0ebf159da3cfb2ab-1)*__obf_1429cb6a843c511b)
	}
	__obf_ca4f88e82278c739 = make(map[string]int)
	for _, __obf_25a4e3c7bdbf84f4 := range __obf_c5a1fa438321e828 {
		__obf_ca4f88e82278c739[__obf_2f32f638bd550f86(__obf_25a4e3c7bdbf84f4)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_91e29e63b8ab7e94 uint) (__obf_bb96b0cd6c8fa7bc uint) {
	if __obf_91e29e63b8ab7e94 == 0 {
		return 1
	}
	for __obf_91e29e63b8ab7e94 != 0 {
		__obf_91e29e63b8ab7e94 /= 10
		__obf_bb96b0cd6c8fa7bc++
	}
	return __obf_bb96b0cd6c8fa7bc
}

// 计算阶乘: n!
func Factorial(__obf_cea35a19fa824f05 int) (__obf_604a6ceec2d312ba int) {
	if __obf_cea35a19fa824f05 > 0 {
		__obf_604a6ceec2d312ba = 1
		for __obf_91e29e63b8ab7e94 := 1; __obf_91e29e63b8ab7e94 <= __obf_cea35a19fa824f05; __obf_91e29e63b8ab7e94++ {
			__obf_604a6ceec2d312ba *= __obf_91e29e63b8ab7e94
		}
	}

	return __obf_604a6ceec2d312ba
}

// base 的 exp 次方
func IntPow(__obf_7ad95d72b923641e, __obf_9529397981ebe104 int) int {
	__obf_48461939034b8ae1 := 1
	for {
		if __obf_9529397981ebe104&1 == 1 {
			__obf_48461939034b8ae1 *= __obf_7ad95d72b923641e
		}
		__obf_9529397981ebe104 >>= 1
		if __obf_9529397981ebe104 == 0 {
			break
		}
		__obf_7ad95d72b923641e *= __obf_7ad95d72b923641e
	}

	return __obf_48461939034b8ae1
}
