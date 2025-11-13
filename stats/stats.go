package __obf_9d1e97a8e05fedb6

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_d4d18bbab73e7cfe []float64) (max float64, __obf_43faab21d338d2b9 error) {

	__obf_523a5e8263c3a631 := len(__obf_d4d18bbab73e7cfe)

	// Return an error if there are no numbers
	if __obf_523a5e8263c3a631 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_d4d18bbab73e7cfe[0]

	// Loop and replace higher values
	for __obf_2a83e31b68e40c3b := 1; __obf_2a83e31b68e40c3b < __obf_523a5e8263c3a631; __obf_2a83e31b68e40c3b++ {
		if __obf_d4d18bbab73e7cfe[__obf_2a83e31b68e40c3b] > max {
			max = __obf_d4d18bbab73e7cfe[__obf_2a83e31b68e40c3b]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_d4d18bbab73e7cfe []float64) (min float64, __obf_43faab21d338d2b9 error) {

	// Get the count of numbers in the slice
	__obf_523a5e8263c3a631 := len(__obf_d4d18bbab73e7cfe)

	// Return an error if there are no numbers
	if __obf_523a5e8263c3a631 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_d4d18bbab73e7cfe[0]

	// Iterate until done checking for a lower value
	for __obf_2a83e31b68e40c3b := 1; __obf_2a83e31b68e40c3b < __obf_523a5e8263c3a631; __obf_2a83e31b68e40c3b++ {
		if __obf_d4d18bbab73e7cfe[__obf_2a83e31b68e40c3b] < min {
			min = __obf_d4d18bbab73e7cfe[__obf_2a83e31b68e40c3b]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_d4d18bbab73e7cfe float64, __obf_fe13fd8904a82245 int) (__obf_28ed9bc1c602759a float64, __obf_43faab21d338d2b9 error) {

	// If the float is not a number
	if math.IsNaN(__obf_d4d18bbab73e7cfe) {
		return math.NaN(), ErrNaN
	}

	// Find out the actual sign and correct the input for later
	__obf_0fb54c3433b9bf2c := 1.0
	if __obf_d4d18bbab73e7cfe < 0 {
		__obf_0fb54c3433b9bf2c = -1
		__obf_d4d18bbab73e7cfe *= -1
	}

	// Use the places arg to get the amount of precision wanted
	__obf_b54026d0e151a335 := math.Pow(10, float64(__obf_fe13fd8904a82245))

	// Find the decimal place we are looking to round
	__obf_0c3b99d498661ae5 := __obf_d4d18bbab73e7cfe * __obf_b54026d0e151a335

	// Get the actual decimal number as a fraction to be compared
	_, __obf_6ca4775b6e9cbd92 := math.Modf(__obf_0c3b99d498661ae5)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_6ca4775b6e9cbd92 >= 0.5 {
		__obf_28ed9bc1c602759a = math.Ceil(__obf_0c3b99d498661ae5)
	} else {
		__obf_28ed9bc1c602759a = math.Floor(__obf_0c3b99d498661ae5)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_28ed9bc1c602759a / __obf_b54026d0e151a335 * __obf_0fb54c3433b9bf2c, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_d4d18bbab73e7cfe []float64) (float64, error) {

	__obf_523a5e8263c3a631 := len(__obf_d4d18bbab73e7cfe)
	if __obf_523a5e8263c3a631 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	__obf_c45cdf0bf16b659e, _ := Sum(__obf_d4d18bbab73e7cfe)

	return __obf_c45cdf0bf16b659e / float64(__obf_523a5e8263c3a631), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_d4d18bbab73e7cfe []float64) (__obf_4aff07b45e0bf4f2 float64, __obf_43faab21d338d2b9 error) {

	// Start by sorting a copy of the slice
	__obf_83f0eed854f2a6ce := __obf_4ab795f70adc8177(__obf_d4d18bbab73e7cfe)

	// No math is needed if there are no numbers
	// For even numbers we add the two middle numbers
	// and divide by two using the mean function above
	// For odd numbers we just use the middle number
	__obf_523a5e8263c3a631 := len(__obf_83f0eed854f2a6ce)
	if __obf_523a5e8263c3a631 == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_523a5e8263c3a631%2 == 0 {
		__obf_4aff07b45e0bf4f2, _ = Mean(__obf_83f0eed854f2a6ce[__obf_523a5e8263c3a631/2-1 : __obf_523a5e8263c3a631/2+1])
	} else {
		__obf_4aff07b45e0bf4f2 = __obf_83f0eed854f2a6ce[__obf_523a5e8263c3a631/2]
	}

	return __obf_4aff07b45e0bf4f2, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_d4d18bbab73e7cfe []float64) (__obf_c45cdf0bf16b659e float64, __obf_43faab21d338d2b9 error) {

	if len(__obf_d4d18bbab73e7cfe) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_269f64b4aa88ccec := range __obf_d4d18bbab73e7cfe {
		__obf_c45cdf0bf16b659e += __obf_269f64b4aa88ccec
	}

	return __obf_c45cdf0bf16b659e, nil
}

// 得到一个区间map
func RangeCount(__obf_74480f17180af927 float64, __obf_104e8fae2442066e int, __obf_d4d18bbab73e7cfe []float64) (__obf_8f8e33fc9ad7ab77 map[string]int, __obf_43faab21d338d2b9 error) {
	if __obf_104e8fae2442066e <= 0 {
		return nil, ErrIllegalNum
	}

	// total, _ := Max(input)
	__obf_8245eb3cf7fdf887 := __obf_27e7cc340c52f451(__obf_74480f17180af927 / float64(__obf_104e8fae2442066e))

	__obf_09bf4719d5e344e9 := func(__obf_b6dc3966d3f76245 float64) string {
		for __obf_2a83e31b68e40c3b := 1; __obf_2a83e31b68e40c3b < __obf_104e8fae2442066e; __obf_2a83e31b68e40c3b++ {
			if __obf_b6dc3966d3f76245 < float64(__obf_2a83e31b68e40c3b*__obf_8245eb3cf7fdf887) {
				return fmt.Sprintf("%d~%d", (__obf_2a83e31b68e40c3b-1)*__obf_8245eb3cf7fdf887, __obf_2a83e31b68e40c3b*__obf_8245eb3cf7fdf887)
			}
		}
		return fmt.Sprintf("%d~", (__obf_104e8fae2442066e-1)*__obf_8245eb3cf7fdf887)
	}

	__obf_8f8e33fc9ad7ab77 = make(map[string]int)
	for _, __obf_3078aa4307bdff01 := range __obf_d4d18bbab73e7cfe {
		__obf_8f8e33fc9ad7ab77[__obf_09bf4719d5e344e9(__obf_3078aa4307bdff01)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_2a83e31b68e40c3b uint) (__obf_83f0eed854f2a6ce uint) {
	if __obf_2a83e31b68e40c3b == 0 {
		return 1
	}
	for __obf_2a83e31b68e40c3b != 0 {
		__obf_2a83e31b68e40c3b /= 10
		__obf_83f0eed854f2a6ce++
	}
	return __obf_83f0eed854f2a6ce
}

// 计算阶乘: n!
func Factorial(__obf_269f64b4aa88ccec int) (__obf_36b5f10bd8d164ff int) {
	if __obf_269f64b4aa88ccec > 0 {
		__obf_36b5f10bd8d164ff = 1
		for __obf_2a83e31b68e40c3b := 1; __obf_2a83e31b68e40c3b <= __obf_269f64b4aa88ccec; __obf_2a83e31b68e40c3b++ {
			__obf_36b5f10bd8d164ff *= __obf_2a83e31b68e40c3b
		}
	}

	return __obf_36b5f10bd8d164ff
}

// base 的 exp 次方
func IntPow(__obf_0f4a563a40d9cde8, __obf_6baad7f8f629948d int) int {
	__obf_30bdbbe91fdd9a8a := 1
	for {
		if __obf_6baad7f8f629948d&1 == 1 {
			__obf_30bdbbe91fdd9a8a *= __obf_0f4a563a40d9cde8
		}
		__obf_6baad7f8f629948d >>= 1
		if __obf_6baad7f8f629948d == 0 {
			break
		}
		__obf_0f4a563a40d9cde8 *= __obf_0f4a563a40d9cde8
	}

	return __obf_30bdbbe91fdd9a8a
}
