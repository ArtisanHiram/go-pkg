package __obf_ce59006f265ba976

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_5884ff1bb2ed3b04 []float64) (max float64, __obf_2a4fcc8899864600 error) {

	__obf_0da71cfecaf7d1fa := len(__obf_5884ff1bb2ed3b04)

	// Return an error if there are no numbers
	if __obf_0da71cfecaf7d1fa == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_5884ff1bb2ed3b04[0]

	// Loop and replace higher values
	for __obf_9554e949c69e0962 := 1; __obf_9554e949c69e0962 < __obf_0da71cfecaf7d1fa; __obf_9554e949c69e0962++ {
		if __obf_5884ff1bb2ed3b04[__obf_9554e949c69e0962] > max {
			max = __obf_5884ff1bb2ed3b04[__obf_9554e949c69e0962]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_5884ff1bb2ed3b04 []float64) (min float64, __obf_2a4fcc8899864600 error) {

	// Get the count of numbers in the slice
	__obf_0da71cfecaf7d1fa := len(__obf_5884ff1bb2ed3b04)

	// Return an error if there are no numbers
	if __obf_0da71cfecaf7d1fa == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_5884ff1bb2ed3b04[0]

	// Iterate until done checking for a lower value
	for __obf_9554e949c69e0962 := 1; __obf_9554e949c69e0962 < __obf_0da71cfecaf7d1fa; __obf_9554e949c69e0962++ {
		if __obf_5884ff1bb2ed3b04[__obf_9554e949c69e0962] < min {
			min = __obf_5884ff1bb2ed3b04[__obf_9554e949c69e0962]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_5884ff1bb2ed3b04 float64, __obf_02f3f445a51096c6 int) (__obf_c87c765a3c6706b9 float64, __obf_2a4fcc8899864600 error) {

	// If the float is not a number
	if math.IsNaN(__obf_5884ff1bb2ed3b04) {
		return math.NaN(), ErrNaN
	}

	// Find out the actual sign and correct the input for later
	__obf_6b0f56b67fc6aad8 := 1.0
	if __obf_5884ff1bb2ed3b04 < 0 {
		__obf_6b0f56b67fc6aad8 = -1
		__obf_5884ff1bb2ed3b04 *= -1
	}

	// Use the places arg to get the amount of precision wanted
	__obf_33e3302a6ce0fcbb := math.Pow(10, float64(__obf_02f3f445a51096c6))

	// Find the decimal place we are looking to round
	__obf_11789b6a343ba6a5 := __obf_5884ff1bb2ed3b04 * __obf_33e3302a6ce0fcbb

	// Get the actual decimal number as a fraction to be compared
	_, __obf_6bd65c5980555e1c := math.Modf(__obf_11789b6a343ba6a5)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_6bd65c5980555e1c >= 0.5 {
		__obf_c87c765a3c6706b9 = math.Ceil(__obf_11789b6a343ba6a5)
	} else {
		__obf_c87c765a3c6706b9 = math.Floor(__obf_11789b6a343ba6a5)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_c87c765a3c6706b9 / __obf_33e3302a6ce0fcbb * __obf_6b0f56b67fc6aad8, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_5884ff1bb2ed3b04 []float64) (float64, error) {

	__obf_0da71cfecaf7d1fa := len(__obf_5884ff1bb2ed3b04)
	if __obf_0da71cfecaf7d1fa == 0 {
		return math.NaN(), ErrEmptyInput
	}

	__obf_c68e61bd450dc65a, _ := Sum(__obf_5884ff1bb2ed3b04)

	return __obf_c68e61bd450dc65a / float64(__obf_0da71cfecaf7d1fa), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_5884ff1bb2ed3b04 []float64) (__obf_2a87d655b42827ad float64, __obf_2a4fcc8899864600 error) {

	// Start by sorting a copy of the slice
	__obf_232c105f7cb968c1 := __obf_cd4216a3bd4147b5(__obf_5884ff1bb2ed3b04)

	// No math is needed if there are no numbers
	// For even numbers we add the two middle numbers
	// and divide by two using the mean function above
	// For odd numbers we just use the middle number
	__obf_0da71cfecaf7d1fa := len(__obf_232c105f7cb968c1)
	if __obf_0da71cfecaf7d1fa == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_0da71cfecaf7d1fa%2 == 0 {
		__obf_2a87d655b42827ad, _ = Mean(__obf_232c105f7cb968c1[__obf_0da71cfecaf7d1fa/2-1 : __obf_0da71cfecaf7d1fa/2+1])
	} else {
		__obf_2a87d655b42827ad = __obf_232c105f7cb968c1[__obf_0da71cfecaf7d1fa/2]
	}

	return __obf_2a87d655b42827ad, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_5884ff1bb2ed3b04 []float64) (__obf_c68e61bd450dc65a float64, __obf_2a4fcc8899864600 error) {

	if len(__obf_5884ff1bb2ed3b04) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_d1237cc562272f8f := range __obf_5884ff1bb2ed3b04 {
		__obf_c68e61bd450dc65a += __obf_d1237cc562272f8f
	}

	return __obf_c68e61bd450dc65a, nil
}

// 得到一个区间map
func RangeCount(__obf_f706afb7050f846b float64, __obf_33f1c36ef655011d int, __obf_5884ff1bb2ed3b04 []float64) (__obf_c150666d09288129 map[string]int, __obf_2a4fcc8899864600 error) {
	if __obf_33f1c36ef655011d <= 0 {
		return nil, ErrIllegalNum
	}

	// total, _ := Max(input)
	__obf_6bb6ec24db37c9cf := __obf_d98fcee495fc36b5(__obf_f706afb7050f846b / float64(__obf_33f1c36ef655011d))

	__obf_ce0563fdc03f090a := func(__obf_e73118b0b4ec69d9 float64) string {
		for __obf_9554e949c69e0962 := 1; __obf_9554e949c69e0962 < __obf_33f1c36ef655011d; __obf_9554e949c69e0962++ {
			if __obf_e73118b0b4ec69d9 < float64(__obf_9554e949c69e0962*__obf_6bb6ec24db37c9cf) {
				return fmt.Sprintf("%d~%d", (__obf_9554e949c69e0962-1)*__obf_6bb6ec24db37c9cf, __obf_9554e949c69e0962*__obf_6bb6ec24db37c9cf)
			}
		}
		return fmt.Sprintf("%d~", (__obf_33f1c36ef655011d-1)*__obf_6bb6ec24db37c9cf)
	}

	__obf_c150666d09288129 = make(map[string]int)
	for _, __obf_d8aef033f1d7caa2 := range __obf_5884ff1bb2ed3b04 {
		__obf_c150666d09288129[__obf_ce0563fdc03f090a(__obf_d8aef033f1d7caa2)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_9554e949c69e0962 uint) (__obf_232c105f7cb968c1 uint) {
	if __obf_9554e949c69e0962 == 0 {
		return 1
	}
	for __obf_9554e949c69e0962 != 0 {
		__obf_9554e949c69e0962 /= 10
		__obf_232c105f7cb968c1++
	}
	return __obf_232c105f7cb968c1
}

// 计算阶乘: n!
func Factorial(__obf_d1237cc562272f8f int) (__obf_e88a53799ffa2488 int) {
	if __obf_d1237cc562272f8f > 0 {
		__obf_e88a53799ffa2488 = 1
		for __obf_9554e949c69e0962 := 1; __obf_9554e949c69e0962 <= __obf_d1237cc562272f8f; __obf_9554e949c69e0962++ {
			__obf_e88a53799ffa2488 *= __obf_9554e949c69e0962
		}
	}

	return __obf_e88a53799ffa2488
}

// base 的 exp 次方
func IntPow(__obf_fdab6c0803c1c655, __obf_16925749356aa619 int) int {
	__obf_fbc776edc3e22695 := 1
	for {
		if __obf_16925749356aa619&1 == 1 {
			__obf_fbc776edc3e22695 *= __obf_fdab6c0803c1c655
		}
		__obf_16925749356aa619 >>= 1
		if __obf_16925749356aa619 == 0 {
			break
		}
		__obf_fdab6c0803c1c655 *= __obf_fdab6c0803c1c655
	}

	return __obf_fbc776edc3e22695
}
