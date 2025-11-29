package __obf_aa50dbb4389f0ad9

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_79ecf1e2364bfac6 []float64) (max float64, __obf_5cc22c21b7b4e9ff error) {
	__obf_327a83ba05aa82bd := len(__obf_79ecf1e2364bfac6)

	// Return an error if there are no numbers
	if __obf_327a83ba05aa82bd == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_79ecf1e2364bfac6[0]

	// Loop and replace higher values
	for __obf_5a121396de114aab := 1; __obf_5a121396de114aab < __obf_327a83ba05aa82bd; __obf_5a121396de114aab++ {
		if __obf_79ecf1e2364bfac6[__obf_5a121396de114aab] > max {
			max = __obf_79ecf1e2364bfac6[__obf_5a121396de114aab]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_79ecf1e2364bfac6 []float64) (min float64, __obf_5cc22c21b7b4e9ff error) {
	__obf_327a83ba05aa82bd := // Get the count of numbers in the slice
		len(__obf_79ecf1e2364bfac6)

	// Return an error if there are no numbers
	if __obf_327a83ba05aa82bd == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_79ecf1e2364bfac6[0]

	// Iterate until done checking for a lower value
	for __obf_5a121396de114aab := 1; __obf_5a121396de114aab < __obf_327a83ba05aa82bd; __obf_5a121396de114aab++ {
		if __obf_79ecf1e2364bfac6[__obf_5a121396de114aab] < min {
			min = __obf_79ecf1e2364bfac6[__obf_5a121396de114aab]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_79ecf1e2364bfac6 float64, __obf_ad739921ab60ee6b int) (__obf_8f2ec84d204b7892 float64, __obf_5cc22c21b7b4e9ff error) {

	// If the float is not a number
	if math.IsNaN(__obf_79ecf1e2364bfac6) {
		return math.NaN(), ErrNaN
	}
	__obf_e2e6ad696317d3b9 := // Find out the actual sign and correct the input for later
		1.0
	if __obf_79ecf1e2364bfac6 < 0 {
		__obf_e2e6ad696317d3b9 = -1
		__obf_79ecf1e2364bfac6 *= -1
	}
	__obf_aeb99608530addd9 := // Use the places arg to get the amount of precision wanted
		math.Pow(10, float64(__obf_ad739921ab60ee6b))
	__obf_5a90cd3f209a8b34 := // Find the decimal place we are looking to round
		__obf_79ecf1e2364bfac6 * __obf_aeb99608530addd9

	// Get the actual decimal number as a fraction to be compared
	_, __obf_b91f134e08f75d89 := math.Modf(__obf_5a90cd3f209a8b34)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_b91f134e08f75d89 >= 0.5 {
		__obf_8f2ec84d204b7892 = math.Ceil(__obf_5a90cd3f209a8b34)
	} else {
		__obf_8f2ec84d204b7892 = math.Floor(__obf_5a90cd3f209a8b34)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_8f2ec84d204b7892 / __obf_aeb99608530addd9 * __obf_e2e6ad696317d3b9, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_79ecf1e2364bfac6 []float64) (float64, error) {
	__obf_327a83ba05aa82bd := len(__obf_79ecf1e2364bfac6)
	if __obf_327a83ba05aa82bd == 0 {
		return math.NaN(), ErrEmptyInput
	}
	__obf_860bb2c3d69402d3, _ := Sum(__obf_79ecf1e2364bfac6)

	return __obf_860bb2c3d69402d3 / float64(__obf_327a83ba05aa82bd), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_79ecf1e2364bfac6 []float64) (__obf_ef2a6753af11ffea float64, __obf_5cc22c21b7b4e9ff error) {
	__obf_dc5d868b7e24cb30 := // Start by sorting a copy of the slice
		__obf_78a9e1790a5bd3a2(__obf_79ecf1e2364bfac6)
	__obf_327a83ba05aa82bd := // No math is needed if there are no numbers
		// For even numbers we add the two middle numbers
		// and divide by two using the mean function above
		// For odd numbers we just use the middle number
		len(__obf_dc5d868b7e24cb30)
	if __obf_327a83ba05aa82bd == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_327a83ba05aa82bd%2 == 0 {
		__obf_ef2a6753af11ffea, _ = Mean(__obf_dc5d868b7e24cb30[__obf_327a83ba05aa82bd/2-1 : __obf_327a83ba05aa82bd/2+1])
	} else {
		__obf_ef2a6753af11ffea = __obf_dc5d868b7e24cb30[__obf_327a83ba05aa82bd/2]
	}

	return __obf_ef2a6753af11ffea, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_79ecf1e2364bfac6 []float64) (__obf_860bb2c3d69402d3 float64, __obf_5cc22c21b7b4e9ff error) {

	if len(__obf_79ecf1e2364bfac6) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_78e5e22631997aa3 := range __obf_79ecf1e2364bfac6 {
		__obf_860bb2c3d69402d3 += __obf_78e5e22631997aa3
	}

	return __obf_860bb2c3d69402d3, nil
}

// 得到一个区间map
func RangeCount(__obf_d598bd84af74ce05 float64, __obf_3c28c342210e3a60 int, __obf_79ecf1e2364bfac6 []float64) (__obf_7191614bcfb6086b map[string]int, __obf_5cc22c21b7b4e9ff error) {
	if __obf_3c28c342210e3a60 <= 0 {
		return nil, ErrIllegalNum
	}
	__obf_744965df6f3fa626 := // total, _ := Max(input)
		__obf_dfb49e042673d42e(__obf_d598bd84af74ce05 / float64(__obf_3c28c342210e3a60))
	__obf_89aa2be738b658dd := func(__obf_c11705c7df178947 float64) string {
		for __obf_5a121396de114aab := 1; __obf_5a121396de114aab < __obf_3c28c342210e3a60; __obf_5a121396de114aab++ {
			if __obf_c11705c7df178947 < float64(__obf_5a121396de114aab*__obf_744965df6f3fa626) {
				return fmt.Sprintf("%d~%d", (__obf_5a121396de114aab-1)*__obf_744965df6f3fa626, __obf_5a121396de114aab*__obf_744965df6f3fa626)
			}
		}
		return fmt.Sprintf("%d~", (__obf_3c28c342210e3a60-1)*__obf_744965df6f3fa626)
	}
	__obf_7191614bcfb6086b = make(map[string]int)
	for _, __obf_9a182482973f940c := range __obf_79ecf1e2364bfac6 {
		__obf_7191614bcfb6086b[__obf_89aa2be738b658dd(__obf_9a182482973f940c)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_5a121396de114aab uint) (__obf_dc5d868b7e24cb30 uint) {
	if __obf_5a121396de114aab == 0 {
		return 1
	}
	for __obf_5a121396de114aab != 0 {
		__obf_5a121396de114aab /= 10
		__obf_dc5d868b7e24cb30++
	}
	return __obf_dc5d868b7e24cb30
}

// 计算阶乘: n!
func Factorial(__obf_78e5e22631997aa3 int) (__obf_5e8a55a14064a364 int) {
	if __obf_78e5e22631997aa3 > 0 {
		__obf_5e8a55a14064a364 = 1
		for __obf_5a121396de114aab := 1; __obf_5a121396de114aab <= __obf_78e5e22631997aa3; __obf_5a121396de114aab++ {
			__obf_5e8a55a14064a364 *= __obf_5a121396de114aab
		}
	}

	return __obf_5e8a55a14064a364
}

// base 的 exp 次方
func IntPow(__obf_328e1522ad41d047, __obf_00dec74cf741ff92 int) int {
	__obf_f0843ff72420340f := 1
	for {
		if __obf_00dec74cf741ff92&1 == 1 {
			__obf_f0843ff72420340f *= __obf_328e1522ad41d047
		}
		__obf_00dec74cf741ff92 >>= 1
		if __obf_00dec74cf741ff92 == 0 {
			break
		}
		__obf_328e1522ad41d047 *= __obf_328e1522ad41d047
	}

	return __obf_f0843ff72420340f
}
