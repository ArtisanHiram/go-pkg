package __obf_face5ef149f4f55f

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_0f75e398d64caf9c []float64) (max float64, __obf_c7c3606d4e4bdb79 error) {
	__obf_04e6dee297b5993c := len(__obf_0f75e398d64caf9c)

	// Return an error if there are no numbers
	if __obf_04e6dee297b5993c == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_0f75e398d64caf9c[0]

	// Loop and replace higher values
	for __obf_fd659d86c467cd75 := 1; __obf_fd659d86c467cd75 < __obf_04e6dee297b5993c; __obf_fd659d86c467cd75++ {
		if __obf_0f75e398d64caf9c[__obf_fd659d86c467cd75] > max {
			max = __obf_0f75e398d64caf9c[__obf_fd659d86c467cd75]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_0f75e398d64caf9c []float64) (min float64, __obf_c7c3606d4e4bdb79 error) {
	__obf_04e6dee297b5993c := // Get the count of numbers in the slice
		len(__obf_0f75e398d64caf9c)

	// Return an error if there are no numbers
	if __obf_04e6dee297b5993c == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_0f75e398d64caf9c[0]

	// Iterate until done checking for a lower value
	for __obf_fd659d86c467cd75 := 1; __obf_fd659d86c467cd75 < __obf_04e6dee297b5993c; __obf_fd659d86c467cd75++ {
		if __obf_0f75e398d64caf9c[__obf_fd659d86c467cd75] < min {
			min = __obf_0f75e398d64caf9c[__obf_fd659d86c467cd75]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_0f75e398d64caf9c float64, __obf_ed40b3d2615cefef int) (__obf_68432e77f5456161 float64, __obf_c7c3606d4e4bdb79 error) {

	// If the float is not a number
	if math.IsNaN(__obf_0f75e398d64caf9c) {
		return math.NaN(), ErrNaN
	}
	__obf_c1c56b992986daca := // Find out the actual sign and correct the input for later
		1.0
	if __obf_0f75e398d64caf9c < 0 {
		__obf_c1c56b992986daca = -1
		__obf_0f75e398d64caf9c *= -1
	}
	__obf_005738fec85e6776 := // Use the places arg to get the amount of precision wanted
		math.Pow(10, float64(__obf_ed40b3d2615cefef))
	__obf_cd96a55387352d1c := // Find the decimal place we are looking to round
		__obf_0f75e398d64caf9c * __obf_005738fec85e6776

	// Get the actual decimal number as a fraction to be compared
	_, __obf_a88fbb39941e6072 := math.Modf(__obf_cd96a55387352d1c)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_a88fbb39941e6072 >= 0.5 {
		__obf_68432e77f5456161 = math.Ceil(__obf_cd96a55387352d1c)
	} else {
		__obf_68432e77f5456161 = math.Floor(__obf_cd96a55387352d1c)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_68432e77f5456161 / __obf_005738fec85e6776 * __obf_c1c56b992986daca, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_0f75e398d64caf9c []float64) (float64, error) {
	__obf_04e6dee297b5993c := len(__obf_0f75e398d64caf9c)
	if __obf_04e6dee297b5993c == 0 {
		return math.NaN(), ErrEmptyInput
	}
	__obf_0ecc87d22d503881, _ := Sum(__obf_0f75e398d64caf9c)

	return __obf_0ecc87d22d503881 / float64(__obf_04e6dee297b5993c), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_0f75e398d64caf9c []float64) (__obf_0cec47034428f739 float64, __obf_c7c3606d4e4bdb79 error) {
	__obf_20ef9d4f82088cee := // Start by sorting a copy of the slice
		__obf_c61afc841612b113(__obf_0f75e398d64caf9c)
	__obf_04e6dee297b5993c := // No math is needed if there are no numbers
		// For even numbers we add the two middle numbers
		// and divide by two using the mean function above
		// For odd numbers we just use the middle number
		len(__obf_20ef9d4f82088cee)
	if __obf_04e6dee297b5993c == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_04e6dee297b5993c%2 == 0 {
		__obf_0cec47034428f739, _ = Mean(__obf_20ef9d4f82088cee[__obf_04e6dee297b5993c/2-1 : __obf_04e6dee297b5993c/2+1])
	} else {
		__obf_0cec47034428f739 = __obf_20ef9d4f82088cee[__obf_04e6dee297b5993c/2]
	}

	return __obf_0cec47034428f739, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_0f75e398d64caf9c []float64) (__obf_0ecc87d22d503881 float64, __obf_c7c3606d4e4bdb79 error) {

	if len(__obf_0f75e398d64caf9c) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_13a8cab6aefe35a1 := range __obf_0f75e398d64caf9c {
		__obf_0ecc87d22d503881 += __obf_13a8cab6aefe35a1
	}

	return __obf_0ecc87d22d503881, nil
}

// 得到一个区间map
func RangeCount(__obf_c17dfa68396ed1a6 float64, __obf_c0e57f69a80bc112 int, __obf_0f75e398d64caf9c []float64) (__obf_9d76a35150fa2888 map[string]int, __obf_c7c3606d4e4bdb79 error) {
	if __obf_c0e57f69a80bc112 <= 0 {
		return nil, ErrIllegalNum
	}
	__obf_b8c1e6450c17d58f := // total, _ := Max(input)
		__obf_fa1f409c06ca1f1e(__obf_c17dfa68396ed1a6 / float64(__obf_c0e57f69a80bc112))
	__obf_d7b7f9b945c2e44b := func(__obf_ad443e34657c5f48 float64) string {
		for __obf_fd659d86c467cd75 := 1; __obf_fd659d86c467cd75 < __obf_c0e57f69a80bc112; __obf_fd659d86c467cd75++ {
			if __obf_ad443e34657c5f48 < float64(__obf_fd659d86c467cd75*__obf_b8c1e6450c17d58f) {
				return fmt.Sprintf("%d~%d", (__obf_fd659d86c467cd75-1)*__obf_b8c1e6450c17d58f, __obf_fd659d86c467cd75*__obf_b8c1e6450c17d58f)
			}
		}
		return fmt.Sprintf("%d~", (__obf_c0e57f69a80bc112-1)*__obf_b8c1e6450c17d58f)
	}
	__obf_9d76a35150fa2888 = make(map[string]int)
	for _, __obf_085a824077a252ce := range __obf_0f75e398d64caf9c {
		__obf_9d76a35150fa2888[__obf_d7b7f9b945c2e44b(__obf_085a824077a252ce)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_fd659d86c467cd75 uint) (__obf_20ef9d4f82088cee uint) {
	if __obf_fd659d86c467cd75 == 0 {
		return 1
	}
	for __obf_fd659d86c467cd75 != 0 {
		__obf_fd659d86c467cd75 /= 10
		__obf_20ef9d4f82088cee++
	}
	return __obf_20ef9d4f82088cee
}

// 计算阶乘: n!
func Factorial(__obf_13a8cab6aefe35a1 int) (__obf_10a1b949a432b9a0 int) {
	if __obf_13a8cab6aefe35a1 > 0 {
		__obf_10a1b949a432b9a0 = 1
		for __obf_fd659d86c467cd75 := 1; __obf_fd659d86c467cd75 <= __obf_13a8cab6aefe35a1; __obf_fd659d86c467cd75++ {
			__obf_10a1b949a432b9a0 *= __obf_fd659d86c467cd75
		}
	}

	return __obf_10a1b949a432b9a0
}

// base 的 exp 次方
func IntPow(__obf_64f1aa652602626d, __obf_3eec4b5700d26783 int) int {
	__obf_c0afeebfb3688bca := 1
	for {
		if __obf_3eec4b5700d26783&1 == 1 {
			__obf_c0afeebfb3688bca *= __obf_64f1aa652602626d
		}
		__obf_3eec4b5700d26783 >>= 1
		if __obf_3eec4b5700d26783 == 0 {
			break
		}
		__obf_64f1aa652602626d *= __obf_64f1aa652602626d
	}

	return __obf_c0afeebfb3688bca
}
