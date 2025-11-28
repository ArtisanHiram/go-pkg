package __obf_0d0d05e597ad9adf

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_d044dbbfe899ea30 []float64) (max float64, __obf_7e8938093f134bb5 error) {

	__obf_f84602ca17d8b0c0 := len(__obf_d044dbbfe899ea30)

	// Return an error if there are no numbers
	if __obf_f84602ca17d8b0c0 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_d044dbbfe899ea30[0]

	// Loop and replace higher values
	for __obf_f09d06d347565b79 := 1; __obf_f09d06d347565b79 < __obf_f84602ca17d8b0c0; __obf_f09d06d347565b79++ {
		if __obf_d044dbbfe899ea30[__obf_f09d06d347565b79] > max {
			max = __obf_d044dbbfe899ea30[__obf_f09d06d347565b79]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_d044dbbfe899ea30 []float64) (min float64, __obf_7e8938093f134bb5 error) {

	// Get the count of numbers in the slice
	__obf_f84602ca17d8b0c0 := len(__obf_d044dbbfe899ea30)

	// Return an error if there are no numbers
	if __obf_f84602ca17d8b0c0 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_d044dbbfe899ea30[0]

	// Iterate until done checking for a lower value
	for __obf_f09d06d347565b79 := 1; __obf_f09d06d347565b79 < __obf_f84602ca17d8b0c0; __obf_f09d06d347565b79++ {
		if __obf_d044dbbfe899ea30[__obf_f09d06d347565b79] < min {
			min = __obf_d044dbbfe899ea30[__obf_f09d06d347565b79]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_d044dbbfe899ea30 float64, __obf_290f79d775691b62 int) (__obf_3c69aa32951f4b36 float64, __obf_7e8938093f134bb5 error) {

	// If the float is not a number
	if math.IsNaN(__obf_d044dbbfe899ea30) {
		return math.NaN(), ErrNaN
	}

	// Find out the actual sign and correct the input for later
	__obf_313f2d3ff0b94e5c := 1.0
	if __obf_d044dbbfe899ea30 < 0 {
		__obf_313f2d3ff0b94e5c = -1
		__obf_d044dbbfe899ea30 *= -1
	}

	// Use the places arg to get the amount of precision wanted
	__obf_9a405e6b718be1bb := math.Pow(10, float64(__obf_290f79d775691b62))

	// Find the decimal place we are looking to round
	__obf_822d1ba03c266e6d := __obf_d044dbbfe899ea30 * __obf_9a405e6b718be1bb

	// Get the actual decimal number as a fraction to be compared
	_, __obf_25b4ef57678b0af5 := math.Modf(__obf_822d1ba03c266e6d)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_25b4ef57678b0af5 >= 0.5 {
		__obf_3c69aa32951f4b36 = math.Ceil(__obf_822d1ba03c266e6d)
	} else {
		__obf_3c69aa32951f4b36 = math.Floor(__obf_822d1ba03c266e6d)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_3c69aa32951f4b36 / __obf_9a405e6b718be1bb * __obf_313f2d3ff0b94e5c, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_d044dbbfe899ea30 []float64) (float64, error) {

	__obf_f84602ca17d8b0c0 := len(__obf_d044dbbfe899ea30)
	if __obf_f84602ca17d8b0c0 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	__obf_df920757ec08e8f0, _ := Sum(__obf_d044dbbfe899ea30)

	return __obf_df920757ec08e8f0 / float64(__obf_f84602ca17d8b0c0), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_d044dbbfe899ea30 []float64) (__obf_944c4cec735d1882 float64, __obf_7e8938093f134bb5 error) {

	// Start by sorting a copy of the slice
	__obf_baa0fb6d0e9e9a00 := __obf_67ec427ad1535581(__obf_d044dbbfe899ea30)

	// No math is needed if there are no numbers
	// For even numbers we add the two middle numbers
	// and divide by two using the mean function above
	// For odd numbers we just use the middle number
	__obf_f84602ca17d8b0c0 := len(__obf_baa0fb6d0e9e9a00)
	if __obf_f84602ca17d8b0c0 == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_f84602ca17d8b0c0%2 == 0 {
		__obf_944c4cec735d1882, _ = Mean(__obf_baa0fb6d0e9e9a00[__obf_f84602ca17d8b0c0/2-1 : __obf_f84602ca17d8b0c0/2+1])
	} else {
		__obf_944c4cec735d1882 = __obf_baa0fb6d0e9e9a00[__obf_f84602ca17d8b0c0/2]
	}

	return __obf_944c4cec735d1882, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_d044dbbfe899ea30 []float64) (__obf_df920757ec08e8f0 float64, __obf_7e8938093f134bb5 error) {

	if len(__obf_d044dbbfe899ea30) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_f74c2486e574f62d := range __obf_d044dbbfe899ea30 {
		__obf_df920757ec08e8f0 += __obf_f74c2486e574f62d
	}

	return __obf_df920757ec08e8f0, nil
}

// 得到一个区间map
func RangeCount(__obf_c046766b1dabd64a float64, __obf_16cb2381b4b6b609 int, __obf_d044dbbfe899ea30 []float64) (__obf_8efc924126050e9f map[string]int, __obf_7e8938093f134bb5 error) {
	if __obf_16cb2381b4b6b609 <= 0 {
		return nil, ErrIllegalNum
	}

	// total, _ := Max(input)
	__obf_9ccb59b7d0834097 := __obf_b3d3a21f3716a950(__obf_c046766b1dabd64a / float64(__obf_16cb2381b4b6b609))

	__obf_ebbe1dde3598a464 := func(__obf_c9ceef1a95cbf5c3 float64) string {
		for __obf_f09d06d347565b79 := 1; __obf_f09d06d347565b79 < __obf_16cb2381b4b6b609; __obf_f09d06d347565b79++ {
			if __obf_c9ceef1a95cbf5c3 < float64(__obf_f09d06d347565b79*__obf_9ccb59b7d0834097) {
				return fmt.Sprintf("%d~%d", (__obf_f09d06d347565b79-1)*__obf_9ccb59b7d0834097, __obf_f09d06d347565b79*__obf_9ccb59b7d0834097)
			}
		}
		return fmt.Sprintf("%d~", (__obf_16cb2381b4b6b609-1)*__obf_9ccb59b7d0834097)
	}

	__obf_8efc924126050e9f = make(map[string]int)
	for _, __obf_59a7386d0097c13c := range __obf_d044dbbfe899ea30 {
		__obf_8efc924126050e9f[__obf_ebbe1dde3598a464(__obf_59a7386d0097c13c)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_f09d06d347565b79 uint) (__obf_baa0fb6d0e9e9a00 uint) {
	if __obf_f09d06d347565b79 == 0 {
		return 1
	}
	for __obf_f09d06d347565b79 != 0 {
		__obf_f09d06d347565b79 /= 10
		__obf_baa0fb6d0e9e9a00++
	}
	return __obf_baa0fb6d0e9e9a00
}

// 计算阶乘: n!
func Factorial(__obf_f74c2486e574f62d int) (__obf_c0b40d8cd73dd2fb int) {
	if __obf_f74c2486e574f62d > 0 {
		__obf_c0b40d8cd73dd2fb = 1
		for __obf_f09d06d347565b79 := 1; __obf_f09d06d347565b79 <= __obf_f74c2486e574f62d; __obf_f09d06d347565b79++ {
			__obf_c0b40d8cd73dd2fb *= __obf_f09d06d347565b79
		}
	}

	return __obf_c0b40d8cd73dd2fb
}

// base 的 exp 次方
func IntPow(__obf_09e04d4ccdc9e5e1, __obf_bf5adf5d0d0bd4a2 int) int {
	__obf_f9293cc8a3314371 := 1
	for {
		if __obf_bf5adf5d0d0bd4a2&1 == 1 {
			__obf_f9293cc8a3314371 *= __obf_09e04d4ccdc9e5e1
		}
		__obf_bf5adf5d0d0bd4a2 >>= 1
		if __obf_bf5adf5d0d0bd4a2 == 0 {
			break
		}
		__obf_09e04d4ccdc9e5e1 *= __obf_09e04d4ccdc9e5e1
	}

	return __obf_f9293cc8a3314371
}
