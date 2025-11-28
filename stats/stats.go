package __obf_13f6e310b0abe7e3

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_f8ab7230e9a1c193 []float64) (max float64, __obf_9ee970d69bb6dfc4 error) {

	__obf_b2b5a40a94c0d4bd := len(__obf_f8ab7230e9a1c193)

	// Return an error if there are no numbers
	if __obf_b2b5a40a94c0d4bd == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_f8ab7230e9a1c193[0]

	// Loop and replace higher values
	for __obf_689941e39eb27aec := 1; __obf_689941e39eb27aec < __obf_b2b5a40a94c0d4bd; __obf_689941e39eb27aec++ {
		if __obf_f8ab7230e9a1c193[__obf_689941e39eb27aec] > max {
			max = __obf_f8ab7230e9a1c193[__obf_689941e39eb27aec]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_f8ab7230e9a1c193 []float64) (min float64, __obf_9ee970d69bb6dfc4 error) {

	// Get the count of numbers in the slice
	__obf_b2b5a40a94c0d4bd := len(__obf_f8ab7230e9a1c193)

	// Return an error if there are no numbers
	if __obf_b2b5a40a94c0d4bd == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_f8ab7230e9a1c193[0]

	// Iterate until done checking for a lower value
	for __obf_689941e39eb27aec := 1; __obf_689941e39eb27aec < __obf_b2b5a40a94c0d4bd; __obf_689941e39eb27aec++ {
		if __obf_f8ab7230e9a1c193[__obf_689941e39eb27aec] < min {
			min = __obf_f8ab7230e9a1c193[__obf_689941e39eb27aec]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_f8ab7230e9a1c193 float64, __obf_f75b23bdb55fddd0 int) (__obf_4de6940e79897a3b float64, __obf_9ee970d69bb6dfc4 error) {

	// If the float is not a number
	if math.IsNaN(__obf_f8ab7230e9a1c193) {
		return math.NaN(), ErrNaN
	}

	// Find out the actual sign and correct the input for later
	__obf_eff20dac878351d1 := 1.0
	if __obf_f8ab7230e9a1c193 < 0 {
		__obf_eff20dac878351d1 = -1
		__obf_f8ab7230e9a1c193 *= -1
	}

	// Use the places arg to get the amount of precision wanted
	__obf_fbd0301f1718de90 := math.Pow(10, float64(__obf_f75b23bdb55fddd0))

	// Find the decimal place we are looking to round
	__obf_8401f16bc44555a8 := __obf_f8ab7230e9a1c193 * __obf_fbd0301f1718de90

	// Get the actual decimal number as a fraction to be compared
	_, __obf_47fec10041a54ef8 := math.Modf(__obf_8401f16bc44555a8)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_47fec10041a54ef8 >= 0.5 {
		__obf_4de6940e79897a3b = math.Ceil(__obf_8401f16bc44555a8)
	} else {
		__obf_4de6940e79897a3b = math.Floor(__obf_8401f16bc44555a8)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_4de6940e79897a3b / __obf_fbd0301f1718de90 * __obf_eff20dac878351d1, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_f8ab7230e9a1c193 []float64) (float64, error) {

	__obf_b2b5a40a94c0d4bd := len(__obf_f8ab7230e9a1c193)
	if __obf_b2b5a40a94c0d4bd == 0 {
		return math.NaN(), ErrEmptyInput
	}

	__obf_283b88fac29bfacf, _ := Sum(__obf_f8ab7230e9a1c193)

	return __obf_283b88fac29bfacf / float64(__obf_b2b5a40a94c0d4bd), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_f8ab7230e9a1c193 []float64) (__obf_3f0be44cea3f5c5c float64, __obf_9ee970d69bb6dfc4 error) {

	// Start by sorting a copy of the slice
	__obf_09de9ede322a8b24 := __obf_1591f280eef6f330(__obf_f8ab7230e9a1c193)

	// No math is needed if there are no numbers
	// For even numbers we add the two middle numbers
	// and divide by two using the mean function above
	// For odd numbers we just use the middle number
	__obf_b2b5a40a94c0d4bd := len(__obf_09de9ede322a8b24)
	if __obf_b2b5a40a94c0d4bd == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_b2b5a40a94c0d4bd%2 == 0 {
		__obf_3f0be44cea3f5c5c, _ = Mean(__obf_09de9ede322a8b24[__obf_b2b5a40a94c0d4bd/2-1 : __obf_b2b5a40a94c0d4bd/2+1])
	} else {
		__obf_3f0be44cea3f5c5c = __obf_09de9ede322a8b24[__obf_b2b5a40a94c0d4bd/2]
	}

	return __obf_3f0be44cea3f5c5c, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_f8ab7230e9a1c193 []float64) (__obf_283b88fac29bfacf float64, __obf_9ee970d69bb6dfc4 error) {

	if len(__obf_f8ab7230e9a1c193) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_0f3e9bacc4234331 := range __obf_f8ab7230e9a1c193 {
		__obf_283b88fac29bfacf += __obf_0f3e9bacc4234331
	}

	return __obf_283b88fac29bfacf, nil
}

// 得到一个区间map
func RangeCount(__obf_78101d22cc5d89c5 float64, __obf_e15c62f4276b009d int, __obf_f8ab7230e9a1c193 []float64) (__obf_055bb5d33ffb069d map[string]int, __obf_9ee970d69bb6dfc4 error) {
	if __obf_e15c62f4276b009d <= 0 {
		return nil, ErrIllegalNum
	}

	// total, _ := Max(input)
	__obf_d7716d8bbdbbdc7a := __obf_a1c4c9e1b18391fa(__obf_78101d22cc5d89c5 / float64(__obf_e15c62f4276b009d))

	__obf_2c9856bba79c4c55 := func(__obf_c390a0ccb56bb441 float64) string {
		for __obf_689941e39eb27aec := 1; __obf_689941e39eb27aec < __obf_e15c62f4276b009d; __obf_689941e39eb27aec++ {
			if __obf_c390a0ccb56bb441 < float64(__obf_689941e39eb27aec*__obf_d7716d8bbdbbdc7a) {
				return fmt.Sprintf("%d~%d", (__obf_689941e39eb27aec-1)*__obf_d7716d8bbdbbdc7a, __obf_689941e39eb27aec*__obf_d7716d8bbdbbdc7a)
			}
		}
		return fmt.Sprintf("%d~", (__obf_e15c62f4276b009d-1)*__obf_d7716d8bbdbbdc7a)
	}

	__obf_055bb5d33ffb069d = make(map[string]int)
	for _, __obf_ff5a871bbcb5a10f := range __obf_f8ab7230e9a1c193 {
		__obf_055bb5d33ffb069d[__obf_2c9856bba79c4c55(__obf_ff5a871bbcb5a10f)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_689941e39eb27aec uint) (__obf_09de9ede322a8b24 uint) {
	if __obf_689941e39eb27aec == 0 {
		return 1
	}
	for __obf_689941e39eb27aec != 0 {
		__obf_689941e39eb27aec /= 10
		__obf_09de9ede322a8b24++
	}
	return __obf_09de9ede322a8b24
}

// 计算阶乘: n!
func Factorial(__obf_0f3e9bacc4234331 int) (__obf_724c4079ca0c30ad int) {
	if __obf_0f3e9bacc4234331 > 0 {
		__obf_724c4079ca0c30ad = 1
		for __obf_689941e39eb27aec := 1; __obf_689941e39eb27aec <= __obf_0f3e9bacc4234331; __obf_689941e39eb27aec++ {
			__obf_724c4079ca0c30ad *= __obf_689941e39eb27aec
		}
	}

	return __obf_724c4079ca0c30ad
}

// base 的 exp 次方
func IntPow(__obf_afd3f14106113783, __obf_9fdbe20ca0c2c501 int) int {
	__obf_8d09d1c7b0e8a028 := 1
	for {
		if __obf_9fdbe20ca0c2c501&1 == 1 {
			__obf_8d09d1c7b0e8a028 *= __obf_afd3f14106113783
		}
		__obf_9fdbe20ca0c2c501 >>= 1
		if __obf_9fdbe20ca0c2c501 == 0 {
			break
		}
		__obf_afd3f14106113783 *= __obf_afd3f14106113783
	}

	return __obf_8d09d1c7b0e8a028
}
