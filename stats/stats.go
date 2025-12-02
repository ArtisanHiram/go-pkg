package __obf_e7d13a84deed4934

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_31fb5f86ef39e603 []float64) (max float64, __obf_de3ff84835a3e8a9 error) {
	__obf_5877a99da19d0c54 := len(__obf_31fb5f86ef39e603)

	// Return an error if there are no numbers
	if __obf_5877a99da19d0c54 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_31fb5f86ef39e603[0]

	// Loop and replace higher values
	for __obf_ad36f3e0e87931f7 := 1; __obf_ad36f3e0e87931f7 < __obf_5877a99da19d0c54; __obf_ad36f3e0e87931f7++ {
		if __obf_31fb5f86ef39e603[__obf_ad36f3e0e87931f7] > max {
			max = __obf_31fb5f86ef39e603[__obf_ad36f3e0e87931f7]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_31fb5f86ef39e603 []float64) (min float64, __obf_de3ff84835a3e8a9 error) {
	__obf_5877a99da19d0c54 := // Get the count of numbers in the slice
		len(__obf_31fb5f86ef39e603)

	// Return an error if there are no numbers
	if __obf_5877a99da19d0c54 == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_31fb5f86ef39e603[0]

	// Iterate until done checking for a lower value
	for __obf_ad36f3e0e87931f7 := 1; __obf_ad36f3e0e87931f7 < __obf_5877a99da19d0c54; __obf_ad36f3e0e87931f7++ {
		if __obf_31fb5f86ef39e603[__obf_ad36f3e0e87931f7] < min {
			min = __obf_31fb5f86ef39e603[__obf_ad36f3e0e87931f7]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_31fb5f86ef39e603 float64, __obf_0d89d3043ed00fbf int) (__obf_dc09ddf9e3af7175 float64, __obf_de3ff84835a3e8a9 error) {

	// If the float is not a number
	if math.IsNaN(__obf_31fb5f86ef39e603) {
		return math.NaN(), ErrNaN
	}
	__obf_281d42886714e22f := // Find out the actual sign and correct the input for later
		1.0
	if __obf_31fb5f86ef39e603 < 0 {
		__obf_281d42886714e22f = -1
		__obf_31fb5f86ef39e603 *= -1
	}
	__obf_7478aec7fba5d3b9 := // Use the places arg to get the amount of precision wanted
		math.Pow(10, float64(__obf_0d89d3043ed00fbf))
	__obf_91e2d4659398f658 := // Find the decimal place we are looking to round
		__obf_31fb5f86ef39e603 * __obf_7478aec7fba5d3b9

	// Get the actual decimal number as a fraction to be compared
	_, __obf_dbabfcccf2c1ba2a := math.Modf(__obf_91e2d4659398f658)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_dbabfcccf2c1ba2a >= 0.5 {
		__obf_dc09ddf9e3af7175 = math.Ceil(__obf_91e2d4659398f658)
	} else {
		__obf_dc09ddf9e3af7175 = math.Floor(__obf_91e2d4659398f658)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_dc09ddf9e3af7175 / __obf_7478aec7fba5d3b9 * __obf_281d42886714e22f, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_31fb5f86ef39e603 []float64) (float64, error) {
	__obf_5877a99da19d0c54 := len(__obf_31fb5f86ef39e603)
	if __obf_5877a99da19d0c54 == 0 {
		return math.NaN(), ErrEmptyInput
	}
	__obf_9254b8e383e48d4e, _ := Sum(__obf_31fb5f86ef39e603)

	return __obf_9254b8e383e48d4e / float64(__obf_5877a99da19d0c54), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_31fb5f86ef39e603 []float64) (__obf_3c8d043b7cc9bf07 float64, __obf_de3ff84835a3e8a9 error) {
	__obf_a3b84026c41e5030 := // Start by sorting a copy of the slice
		__obf_a9fa2c6c1f72c723(__obf_31fb5f86ef39e603)
	__obf_5877a99da19d0c54 := // No math is needed if there are no numbers
		// For even numbers we add the two middle numbers
		// and divide by two using the mean function above
		// For odd numbers we just use the middle number
		len(__obf_a3b84026c41e5030)
	if __obf_5877a99da19d0c54 == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_5877a99da19d0c54%2 == 0 {
		__obf_3c8d043b7cc9bf07, _ = Mean(__obf_a3b84026c41e5030[__obf_5877a99da19d0c54/2-1 : __obf_5877a99da19d0c54/2+1])
	} else {
		__obf_3c8d043b7cc9bf07 = __obf_a3b84026c41e5030[__obf_5877a99da19d0c54/2]
	}

	return __obf_3c8d043b7cc9bf07, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_31fb5f86ef39e603 []float64) (__obf_9254b8e383e48d4e float64, __obf_de3ff84835a3e8a9 error) {

	if len(__obf_31fb5f86ef39e603) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_20a589ec2142db8c := range __obf_31fb5f86ef39e603 {
		__obf_9254b8e383e48d4e += __obf_20a589ec2142db8c
	}

	return __obf_9254b8e383e48d4e, nil
}

// 得到一个区间map
func RangeCount(__obf_929e94bc7d76fd6a float64, __obf_9804678316d9acea int, __obf_31fb5f86ef39e603 []float64) (__obf_a7e3064b3033ab17 map[string]int, __obf_de3ff84835a3e8a9 error) {
	if __obf_9804678316d9acea <= 0 {
		return nil, ErrIllegalNum
	}
	__obf_76c5675bb34befc3 := // total, _ := Max(input)
		__obf_c967c3a35008f1c8(__obf_929e94bc7d76fd6a / float64(__obf_9804678316d9acea))
	__obf_866594f0994b9ffb := func(__obf_dda978585b5eecc4 float64) string {
		for __obf_ad36f3e0e87931f7 := 1; __obf_ad36f3e0e87931f7 < __obf_9804678316d9acea; __obf_ad36f3e0e87931f7++ {
			if __obf_dda978585b5eecc4 < float64(__obf_ad36f3e0e87931f7*__obf_76c5675bb34befc3) {
				return fmt.Sprintf("%d~%d", (__obf_ad36f3e0e87931f7-1)*__obf_76c5675bb34befc3, __obf_ad36f3e0e87931f7*__obf_76c5675bb34befc3)
			}
		}
		return fmt.Sprintf("%d~", (__obf_9804678316d9acea-1)*__obf_76c5675bb34befc3)
	}
	__obf_a7e3064b3033ab17 = make(map[string]int)
	for _, __obf_e0927dd094c67f4f := range __obf_31fb5f86ef39e603 {
		__obf_a7e3064b3033ab17[__obf_866594f0994b9ffb(__obf_e0927dd094c67f4f)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_ad36f3e0e87931f7 uint) (__obf_a3b84026c41e5030 uint) {
	if __obf_ad36f3e0e87931f7 == 0 {
		return 1
	}
	for __obf_ad36f3e0e87931f7 != 0 {
		__obf_ad36f3e0e87931f7 /= 10
		__obf_a3b84026c41e5030++
	}
	return __obf_a3b84026c41e5030
}

// 计算阶乘: n!
func Factorial(__obf_20a589ec2142db8c int) (__obf_93ce884d99da5505 int) {
	if __obf_20a589ec2142db8c > 0 {
		__obf_93ce884d99da5505 = 1
		for __obf_ad36f3e0e87931f7 := 1; __obf_ad36f3e0e87931f7 <= __obf_20a589ec2142db8c; __obf_ad36f3e0e87931f7++ {
			__obf_93ce884d99da5505 *= __obf_ad36f3e0e87931f7
		}
	}

	return __obf_93ce884d99da5505
}

// base 的 exp 次方
func IntPow(__obf_d5cf0aa7d134be23, __obf_d2dbdecfaed64e8a int) int {
	__obf_d9a44870416f840c := 1
	for {
		if __obf_d2dbdecfaed64e8a&1 == 1 {
			__obf_d9a44870416f840c *= __obf_d5cf0aa7d134be23
		}
		__obf_d2dbdecfaed64e8a >>= 1
		if __obf_d2dbdecfaed64e8a == 0 {
			break
		}
		__obf_d5cf0aa7d134be23 *= __obf_d5cf0aa7d134be23
	}

	return __obf_d9a44870416f840c
}
