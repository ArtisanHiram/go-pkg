package __obf_8cdcb71eb89a0b1a

// https://github.com/montanaflynn/stats
import (
	"fmt"
	"math"
)

// Max finds the highest number in a slice
func Max(__obf_12053c429f32e343 []float64) (max float64, __obf_79a8b5a43b28cb22 error) {
	__obf_e6a935bbf21d899c := len(__obf_12053c429f32e343)

	// Return an error if there are no numbers
	if __obf_e6a935bbf21d899c == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	max = __obf_12053c429f32e343[0]

	// Loop and replace higher values
	for __obf_4888a164dbcc1e45 := 1; __obf_4888a164dbcc1e45 < __obf_e6a935bbf21d899c; __obf_4888a164dbcc1e45++ {
		if __obf_12053c429f32e343[__obf_4888a164dbcc1e45] > max {
			max = __obf_12053c429f32e343[__obf_4888a164dbcc1e45]
		}
	}

	return max, nil
}

// Min finds the lowest number in a set of data
func Min(__obf_12053c429f32e343 []float64) (min float64, __obf_79a8b5a43b28cb22 error) {
	__obf_e6a935bbf21d899c := // Get the count of numbers in the slice
		len(__obf_12053c429f32e343)

	// Return an error if there are no numbers
	if __obf_e6a935bbf21d899c == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Get the first value as the starting point
	min = __obf_12053c429f32e343[0]

	// Iterate until done checking for a lower value
	for __obf_4888a164dbcc1e45 := 1; __obf_4888a164dbcc1e45 < __obf_e6a935bbf21d899c; __obf_4888a164dbcc1e45++ {
		if __obf_12053c429f32e343[__obf_4888a164dbcc1e45] < min {
			min = __obf_12053c429f32e343[__obf_4888a164dbcc1e45]
		}
	}
	return min, nil
}

// Round a float to a specific decimal place or precision
func Round(__obf_12053c429f32e343 float64, __obf_c909abea34f35eb6 int) (__obf_6ff69525d975343b float64, __obf_79a8b5a43b28cb22 error) {

	// If the float is not a number
	if math.IsNaN(__obf_12053c429f32e343) {
		return math.NaN(), ErrNaN
	}
	__obf_42d11585ce83b129 := // Find out the actual sign and correct the input for later
		1.0
	if __obf_12053c429f32e343 < 0 {
		__obf_42d11585ce83b129 = -1
		__obf_12053c429f32e343 *= -1
	}
	__obf_346100edf34209dc := // Use the places arg to get the amount of precision wanted
		math.Pow(10, float64(__obf_c909abea34f35eb6))
	__obf_f24f6798e38d2124 := // Find the decimal place we are looking to round
		__obf_12053c429f32e343 * __obf_346100edf34209dc

	// Get the actual decimal number as a fraction to be compared
	_, __obf_8eac8091210bcd8d := math.Modf(__obf_f24f6798e38d2124)

	// If the decimal is less than .5 we round down otherwise up
	if __obf_8eac8091210bcd8d >= 0.5 {
		__obf_6ff69525d975343b = math.Ceil(__obf_f24f6798e38d2124)
	} else {
		__obf_6ff69525d975343b = math.Floor(__obf_f24f6798e38d2124)
	}

	// Finally we do the math to actually create a rounded number
	return __obf_6ff69525d975343b / __obf_346100edf34209dc * __obf_42d11585ce83b129, nil
}

// Mean gets the average of a slice of numbers
func Mean(__obf_12053c429f32e343 []float64) (float64, error) {
	__obf_e6a935bbf21d899c := len(__obf_12053c429f32e343)
	if __obf_e6a935bbf21d899c == 0 {
		return math.NaN(), ErrEmptyInput
	}
	__obf_d17d123a74652e66, _ := Sum(__obf_12053c429f32e343)

	return __obf_d17d123a74652e66 / float64(__obf_e6a935bbf21d899c), nil
}

// Median gets the median number in a slice of numbers
func Median(__obf_12053c429f32e343 []float64) (__obf_69da0450328dbbf5 float64, __obf_79a8b5a43b28cb22 error) {
	__obf_a8b8990b0ba5866c := // Start by sorting a copy of the slice
		__obf_c06efca49e526820(__obf_12053c429f32e343)
	__obf_e6a935bbf21d899c := // No math is needed if there are no numbers
		// For even numbers we add the two middle numbers
		// and divide by two using the mean function above
		// For odd numbers we just use the middle number
		len(__obf_a8b8990b0ba5866c)
	if __obf_e6a935bbf21d899c == 0 {
		return math.NaN(), ErrEmptyInput
	} else if __obf_e6a935bbf21d899c%2 == 0 {
		__obf_69da0450328dbbf5, _ = Mean(__obf_a8b8990b0ba5866c[__obf_e6a935bbf21d899c/2-1 : __obf_e6a935bbf21d899c/2+1])
	} else {
		__obf_69da0450328dbbf5 = __obf_a8b8990b0ba5866c[__obf_e6a935bbf21d899c/2]
	}

	return __obf_69da0450328dbbf5, nil
}

// Sum adds all the numbers of a slice together
func Sum(__obf_12053c429f32e343 []float64) (__obf_d17d123a74652e66 float64, __obf_79a8b5a43b28cb22 error) {

	if len(__obf_12053c429f32e343) == 0 {
		return math.NaN(), ErrEmptyInput
	}

	// Add em up
	for _, __obf_e71572e6e26a93c9 := range __obf_12053c429f32e343 {
		__obf_d17d123a74652e66 += __obf_e71572e6e26a93c9
	}

	return __obf_d17d123a74652e66, nil
}

// 得到一个区间map
func RangeCount(__obf_3873afa06b013162 float64, __obf_23b8b0e1f131257e int, __obf_12053c429f32e343 []float64) (__obf_c18da580d416039d map[string]int, __obf_79a8b5a43b28cb22 error) {
	if __obf_23b8b0e1f131257e <= 0 {
		return nil, ErrIllegalNum
	}
	__obf_b10ec94b26b175e0 := // total, _ := Max(input)
		__obf_f090a72098d77103(__obf_3873afa06b013162 / float64(__obf_23b8b0e1f131257e))
	__obf_881c8ebda7baec3f := func(__obf_1e38cf5ca4344d72 float64) string {
		for __obf_4888a164dbcc1e45 := 1; __obf_4888a164dbcc1e45 < __obf_23b8b0e1f131257e; __obf_4888a164dbcc1e45++ {
			if __obf_1e38cf5ca4344d72 < float64(__obf_4888a164dbcc1e45*__obf_b10ec94b26b175e0) {
				return fmt.Sprintf("%d~%d", (__obf_4888a164dbcc1e45-1)*__obf_b10ec94b26b175e0, __obf_4888a164dbcc1e45*__obf_b10ec94b26b175e0)
			}
		}
		return fmt.Sprintf("%d~", (__obf_23b8b0e1f131257e-1)*__obf_b10ec94b26b175e0)
	}
	__obf_c18da580d416039d = make(map[string]int)
	for _, __obf_38a4509bfc560c01 := range __obf_12053c429f32e343 {
		__obf_c18da580d416039d[__obf_881c8ebda7baec3f(__obf_38a4509bfc560c01)] += 1
	}

	return
}

// 计算整型数多少位
func GetDigits(__obf_4888a164dbcc1e45 uint) (__obf_a8b8990b0ba5866c uint) {
	if __obf_4888a164dbcc1e45 == 0 {
		return 1
	}
	for __obf_4888a164dbcc1e45 != 0 {
		__obf_4888a164dbcc1e45 /= 10
		__obf_a8b8990b0ba5866c++
	}
	return __obf_a8b8990b0ba5866c
}

// 计算阶乘: n!
func Factorial(__obf_e71572e6e26a93c9 int) (__obf_172cb9ba35af351f int) {
	if __obf_e71572e6e26a93c9 > 0 {
		__obf_172cb9ba35af351f = 1
		for __obf_4888a164dbcc1e45 := 1; __obf_4888a164dbcc1e45 <= __obf_e71572e6e26a93c9; __obf_4888a164dbcc1e45++ {
			__obf_172cb9ba35af351f *= __obf_4888a164dbcc1e45
		}
	}

	return __obf_172cb9ba35af351f
}

// base 的 exp 次方
func IntPow(__obf_b7b57c599b13fb17, __obf_decd0af701f3ff68 int) int {
	__obf_e6cfb7b5f80ea49f := 1
	for {
		if __obf_decd0af701f3ff68&1 == 1 {
			__obf_e6cfb7b5f80ea49f *= __obf_b7b57c599b13fb17
		}
		__obf_decd0af701f3ff68 >>= 1
		if __obf_decd0af701f3ff68 == 0 {
			break
		}
		__obf_b7b57c599b13fb17 *= __obf_b7b57c599b13fb17
	}

	return __obf_e6cfb7b5f80ea49f
}
