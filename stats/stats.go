package __obf_0c4a280c8ca42808

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_b97842801be35eb9 []float64) (max float64, __obf_9e05075c97d73684 error) {

	__obf_91f7bae1abca064c := len(__obf_b97842801be35eb9)

	// Return an error if there are no numbers
	if __obf_91f7bae1abca064c == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_b97842801be35eb9[0]

	// Loop and replace higher values
	for __obf_ce2be2f6b607bc4d := 1; __obf_ce2be2f6b607bc4d < __obf_91f7bae1abca064c; __obf_ce2be2f6b607bc4d++ {
		if __obf_b97842801be35eb9[__obf_ce2be2f6b607bc4d] > max {
			max = __obf_b97842801be35eb9[__obf_ce2be2f6b607bc4d]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_b97842801be35eb9 []float64) (min float64, __obf_9e05075c97d73684 error) {

	// Get the count of numbers in the slice
	__obf_91f7bae1abca064c := len(__obf_b97842801be35eb9)

	// Return an error if there are no numbers
	if __obf_91f7bae1abca064c == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_b97842801be35eb9[0]

	// Iterate until done checking for a lower value
	for __obf_ce2be2f6b607bc4d := 1; __obf_ce2be2f6b607bc4d < __obf_91f7bae1abca064c; __obf_ce2be2f6b607bc4d++ {
		if __obf_b97842801be35eb9[__obf_ce2be2f6b607bc4d] < min {
			min = __obf_b97842801be35eb9[__obf_ce2be2f6b607bc4d]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_b97842801be35eb9 float64, __obf_28188f452f979d75 int) (__obf_205be1d2456a15b5 float64, __obf_9e05075c97d73684 error) {

	// If the float is not a number
	if math.IsNaN(__obf_b97842801be35eb9) {
		return math.NaN(), ErrNaN
	}

	// Find out the actual sign and correct the input for later
	__obf_e86d6c74a1c1f918 := 1.0
	if __obf_b97842801be35eb9 < 0 {
		__obf_e86d6c74a1c1f918 = -1
		__obf_b97842801be35eb9 *= -1
	}

	// Use the places arg to get the amount of precision wanted
	__obf_73881694ec7e307f := math.Pow(10, float64(__obf_28188f452f979d75))

	// Find the decimal place we are looking to round
	__obf_1094f6b823185f99 := __obf_b97842801be35eb9 * __obf_73881694ec7e307f

	// Get the actual decimal number as a fraction to be compared
	_, __obf_749468df2b333bfd := math.Modf(__obf_1094f6b823185f99)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_749468df2b333bfd >= 0.5 {
		__obf_205be1d2456a15b5 = math.Ceil(__obf_1094f6b823185f99)
	} else {
		__obf_205be1d2456a15b5 = math.Floor(__obf_1094f6b823185f99)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_205be1d2456a15b5 / __obf_73881694ec7e307f * __obf_e86d6c74a1c1f918, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_b97842801be35eb9 []float64) (float64, error) {

	__obf_91f7bae1abca064c := len(__obf_b97842801be35eb9)
	if __obf_91f7bae1abca064c == 0 {
		return math.NaN(), ErrEmptyInput
	}

	__obf_bbe61034268cb8d0, _ := Sum(__obf_b97842801be35eb9)

	return __obf_bbe61034268cb8d0 / float64(__obf_91f7bae1abca064c), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_b97842801be35eb9 []float64) (__obf_98c24052aab12fca float64, __obf_9e05075c97d73684 error) {

	// Start by sorting a copy of the slice
	__obf_f59bd02bb57399bc := __obf_858f6951a9c2c144(__obf_b97842801be35eb9)

	// No math is needed if there are no numbers
	// For even numbers we add the two middle numbers
	// and divide by two using the mean function above
	// For odd numbers we just use the middle number
	__obf_91f7bae1abca064c := len(__obf_f59bd02bb57399bc)
	if __obf_91f7bae1abca064c == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_91f7bae1abca064c%2 == 0 {
		__obf_98c24052aab12fca, _ = Mean(__obf_f59bd02bb57399bc[__obf_91f7bae1abca064c/2-1 : __obf_91f7bae1abca064c/2+1])
	} else {
		__obf_98c24052aab12fca = __obf_f59bd02bb57399bc[__obf_91f7bae1abca064c/2]
	}

	return __obf_98c24052aab12fca, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_b97842801be35eb9 []float64) (__obf_bbe61034268cb8d0 float64, __obf_9e05075c97d73684 error) {

	if len(__obf_b97842801be35eb9) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_e922e3ea08dbc4be := range __obf_b97842801be35eb9 {
		__obf_bbe61034268cb8d0 += __obf_e922e3ea08dbc4be
	}

	return __obf_bbe61034268cb8d0, nil
}

// 得到一个区间map
func RangeCount(__obf_76a0c58e09f959c1 float64, __obf_3e6ae0e5ee88480b int, __obf_b97842801be35eb9 []float64) (__obf_3b0e9ea0a5a0d22e map[string]int, __obf_9e05075c97d73684 error) {
	if __obf_3e6ae0e5ee88480b <= 0 {
		return nil, ErrIllegalNum
	}

	// total, _ := Max(input)
	__obf_e8620161ee54d9cc := __obf_132364f3f9926888(__obf_76a0c58e09f959c1 / float64(__obf_3e6ae0e5ee88480b))

	__obf_fb72c11e5ad3d6f3 := func(__obf_7eda3435496c70df float64) string {
		for __obf_ce2be2f6b607bc4d := 1; __obf_ce2be2f6b607bc4d < __obf_3e6ae0e5ee88480b; __obf_ce2be2f6b607bc4d++ {
			if __obf_7eda3435496c70df < float64(__obf_ce2be2f6b607bc4d*__obf_e8620161ee54d9cc) {
				return fmt.Sprintf("%d~%d", (__obf_ce2be2f6b607bc4d-1)*__obf_e8620161ee54d9cc, __obf_ce2be2f6b607bc4d*__obf_e8620161ee54d9cc)
			}
		}
		return fmt.Sprintf("%d~", (__obf_3e6ae0e5ee88480b-1)*__obf_e8620161ee54d9cc)
	}

	__obf_3b0e9ea0a5a0d22e = make(map[string]int)
	for _, __obf_fb438f58750cdcd2 := range __obf_b97842801be35eb9 {
		__obf_3b0e9ea0a5a0d22e[__obf_fb72c11e5ad3d6f3(__obf_fb438f58750cdcd2)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_ce2be2f6b607bc4d uint) (__obf_f59bd02bb57399bc uint) {
	if __obf_ce2be2f6b607bc4d == 0 {
		return 1
	}
	for __obf_ce2be2f6b607bc4d != 0 {
		__obf_ce2be2f6b607bc4d /= 10
		__obf_f59bd02bb57399bc++
	}
	return __obf_f59bd02bb57399bc
}

// 计算阶乘: n!
func Factorial(__obf_e922e3ea08dbc4be int) (__obf_73b969aaa886bc0a int) {
	if __obf_e922e3ea08dbc4be > 0 {
		__obf_73b969aaa886bc0a = 1
		for __obf_ce2be2f6b607bc4d := 1; __obf_ce2be2f6b607bc4d <= __obf_e922e3ea08dbc4be; __obf_ce2be2f6b607bc4d++ {
			__obf_73b969aaa886bc0a *= __obf_ce2be2f6b607bc4d
		}
	}

	return __obf_73b969aaa886bc0a
}

// base 的 exp 次方
func IntPow(__obf_c724a8399fd3f399, __obf_51bd1ad8bf6480ba int) int {
	__obf_710fa70c00716d1c := 1
	for {
		if __obf_51bd1ad8bf6480ba&1 == 1 {
			__obf_710fa70c00716d1c *= __obf_c724a8399fd3f399
		}
		__obf_51bd1ad8bf6480ba >>= 1
		if __obf_51bd1ad8bf6480ba == 0 {
			break
		}
		__obf_c724a8399fd3f399 *= __obf_c724a8399fd3f399
	}

	return __obf_710fa70c00716d1c
}
