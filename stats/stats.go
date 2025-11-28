package __obf_c357e2bf526d00f9

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_60f27e022365452d []float64) (max float64, __obf_4a00730613aa7357 error) {

	__obf_605e91a34b590aeb := len(__obf_60f27e022365452d)

	// Return an error if there are no numbers
	if __obf_605e91a34b590aeb == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_60f27e022365452d[0]

	// Loop and replace higher values
	for __obf_1d46e896b0445c7b := 1; __obf_1d46e896b0445c7b < __obf_605e91a34b590aeb; __obf_1d46e896b0445c7b++ {
		if __obf_60f27e022365452d[__obf_1d46e896b0445c7b] > max {
			max = __obf_60f27e022365452d[__obf_1d46e896b0445c7b]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_60f27e022365452d []float64) (min float64, __obf_4a00730613aa7357 error) {

	// Get the count of numbers in the slice
	__obf_605e91a34b590aeb := len(__obf_60f27e022365452d)

	// Return an error if there are no numbers
	if __obf_605e91a34b590aeb == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_60f27e022365452d[0]

	// Iterate until done checking for a lower value
	for __obf_1d46e896b0445c7b := 1; __obf_1d46e896b0445c7b < __obf_605e91a34b590aeb; __obf_1d46e896b0445c7b++ {
		if __obf_60f27e022365452d[__obf_1d46e896b0445c7b] < min {
			min = __obf_60f27e022365452d[__obf_1d46e896b0445c7b]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_60f27e022365452d float64, __obf_6e7a46d95ea288a4 int) (__obf_06feb888d34ddcfc float64, __obf_4a00730613aa7357 error) {

	// If the float is not a number
	if math.IsNaN(__obf_60f27e022365452d) {
		return math.NaN(), ErrNaN
	}

	// Find out the actual sign and correct the input for later
	__obf_9827661664545ae9 := 1.0
	if __obf_60f27e022365452d < 0 {
		__obf_9827661664545ae9 = -1
		__obf_60f27e022365452d *= -1
	}

	// Use the places arg to get the amount of precision wanted
	__obf_78ac38dd590d0de2 := math.Pow(10, float64(__obf_6e7a46d95ea288a4))

	// Find the decimal place we are looking to round
	__obf_f5bba65e31e19aa8 := __obf_60f27e022365452d * __obf_78ac38dd590d0de2

	// Get the actual decimal number as a fraction to be compared
	_, __obf_25f20098d257022e := math.Modf(__obf_f5bba65e31e19aa8)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_25f20098d257022e >= 0.5 {
		__obf_06feb888d34ddcfc = math.Ceil(__obf_f5bba65e31e19aa8)
	} else {
		__obf_06feb888d34ddcfc = math.Floor(__obf_f5bba65e31e19aa8)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_06feb888d34ddcfc / __obf_78ac38dd590d0de2 * __obf_9827661664545ae9, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_60f27e022365452d []float64) (float64, error) {

	__obf_605e91a34b590aeb := len(__obf_60f27e022365452d)
	if __obf_605e91a34b590aeb == 0 {
		return math.NaN(), ErrEmptyInput
	}

	__obf_63c11d7ad215cf21, _ := Sum(__obf_60f27e022365452d)

	return __obf_63c11d7ad215cf21 / float64(__obf_605e91a34b590aeb), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_60f27e022365452d []float64) (__obf_059be6e20f69c186 float64, __obf_4a00730613aa7357 error) {

	// Start by sorting a copy of the slice
	__obf_78755b923a41f260 := __obf_a9d139b4fa16b92f(__obf_60f27e022365452d)

	// No math is needed if there are no numbers
	// For even numbers we add the two middle numbers
	// and divide by two using the mean function above
	// For odd numbers we just use the middle number
	__obf_605e91a34b590aeb := len(__obf_78755b923a41f260)
	if __obf_605e91a34b590aeb == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_605e91a34b590aeb%2 == 0 {
		__obf_059be6e20f69c186, _ = Mean(__obf_78755b923a41f260[__obf_605e91a34b590aeb/2-1 : __obf_605e91a34b590aeb/2+1])
	} else {
		__obf_059be6e20f69c186 = __obf_78755b923a41f260[__obf_605e91a34b590aeb/2]
	}

	return __obf_059be6e20f69c186, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_60f27e022365452d []float64) (__obf_63c11d7ad215cf21 float64, __obf_4a00730613aa7357 error) {

	if len(__obf_60f27e022365452d) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_0a086816eb79cf22 := range __obf_60f27e022365452d {
		__obf_63c11d7ad215cf21 += __obf_0a086816eb79cf22
	}

	return __obf_63c11d7ad215cf21, nil
}

// 得到一个区间map
func RangeCount(__obf_0a41750ce07f7444 float64, __obf_3d33d9682dd00a73 int, __obf_60f27e022365452d []float64) (__obf_f05c59b02c1ae9e1 map[string]int, __obf_4a00730613aa7357 error) {
	if __obf_3d33d9682dd00a73 <= 0 {
		return nil, ErrIllegalNum
	}

	// total, _ := Max(input)
	__obf_4c727b8681e271b8 := __obf_5a5f26f165c78990(__obf_0a41750ce07f7444 / float64(__obf_3d33d9682dd00a73))

	__obf_bff898bb1b5a5087 := func(__obf_c0fc9ed37536b3f5 float64) string {
		for __obf_1d46e896b0445c7b := 1; __obf_1d46e896b0445c7b < __obf_3d33d9682dd00a73; __obf_1d46e896b0445c7b++ {
			if __obf_c0fc9ed37536b3f5 < float64(__obf_1d46e896b0445c7b*__obf_4c727b8681e271b8) {
				return fmt.Sprintf("%d~%d", (__obf_1d46e896b0445c7b-1)*__obf_4c727b8681e271b8, __obf_1d46e896b0445c7b*__obf_4c727b8681e271b8)
			}
		}
		return fmt.Sprintf("%d~", (__obf_3d33d9682dd00a73-1)*__obf_4c727b8681e271b8)
	}

	__obf_f05c59b02c1ae9e1 = make(map[string]int)
	for _, __obf_d95cb01312def2a4 := range __obf_60f27e022365452d {
		__obf_f05c59b02c1ae9e1[__obf_bff898bb1b5a5087(__obf_d95cb01312def2a4)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_1d46e896b0445c7b uint) (__obf_78755b923a41f260 uint) {
	if __obf_1d46e896b0445c7b == 0 {
		return 1
	}
	for __obf_1d46e896b0445c7b != 0 {
		__obf_1d46e896b0445c7b /= 10
		__obf_78755b923a41f260++
	}
	return __obf_78755b923a41f260
}

// 计算阶乘: n!
func Factorial(__obf_0a086816eb79cf22 int) (__obf_08824ae0b0d34472 int) {
	if __obf_0a086816eb79cf22 > 0 {
		__obf_08824ae0b0d34472 = 1
		for __obf_1d46e896b0445c7b := 1; __obf_1d46e896b0445c7b <= __obf_0a086816eb79cf22; __obf_1d46e896b0445c7b++ {
			__obf_08824ae0b0d34472 *= __obf_1d46e896b0445c7b
		}
	}

	return __obf_08824ae0b0d34472
}

// base 的 exp 次方
func IntPow(__obf_d6922ae321e41ae3, __obf_b47b60d30be9b011 int) int {
	__obf_44b87c4ef225ea5e := 1
	for {
		if __obf_b47b60d30be9b011&1 == 1 {
			__obf_44b87c4ef225ea5e *= __obf_d6922ae321e41ae3
		}
		__obf_b47b60d30be9b011 >>= 1
		if __obf_b47b60d30be9b011 == 0 {
			break
		}
		__obf_d6922ae321e41ae3 *= __obf_d6922ae321e41ae3
	}

	return __obf_44b87c4ef225ea5e
}
