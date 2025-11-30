package __obf_030d39f7a8de96c6

import (
	"fmt"
	"math"
	"strconv"
)

var __obf_faa2ce64bf143410 []uint64

func init() {
	__obf_faa2ce64bf143410 = []uint64{1, 10, 100, 1000, 10000, 100000, 1000000}
}

// WriteFloat32 write float32 to stream
func (__obf_8a2c51fe22775655 *Stream) WriteFloat32(__obf_efaf2719b44ad8ac float32) {
	if math.IsInf(float64(__obf_efaf2719b44ad8ac), 0) || math.IsNaN(float64(__obf_efaf2719b44ad8ac)) {
		__obf_8a2c51fe22775655.
			Error = fmt.Errorf("unsupported value: %f", __obf_efaf2719b44ad8ac)
		return
	}
	__obf_aaded2753c66e327 := math.Abs(float64(__obf_efaf2719b44ad8ac))
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_aaded2753c66e327 != 0 {
		if float32(__obf_aaded2753c66e327) < 1e-6 || float32(__obf_aaded2753c66e327) >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = strconv.AppendFloat(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, float64(__obf_efaf2719b44ad8ac), fmt, -1, 32)
	if fmt == 'e' {
		__obf_e40b3963a92ca4a0 := // clean up e-09 to e-9
			len(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df)
		if __obf_e40b3963a92ca4a0 >= 4 && __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[__obf_e40b3963a92ca4a0-4] == 'e' && __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[__obf_e40b3963a92ca4a0-3] == '-' && __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[__obf_e40b3963a92ca4a0-2] == '0' {
			__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[__obf_e40b3963a92ca4a0-2] = __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[__obf_e40b3963a92ca4a0-1]
			__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[:__obf_e40b3963a92ca4a0-1]
		}
	}
}

// WriteFloat32Lossy write float32 to stream with ONLY 6 digits precision although much much faster
func (__obf_8a2c51fe22775655 *Stream) WriteFloat32Lossy(__obf_efaf2719b44ad8ac float32) {
	if math.IsInf(float64(__obf_efaf2719b44ad8ac), 0) || math.IsNaN(float64(__obf_efaf2719b44ad8ac)) {
		__obf_8a2c51fe22775655.
			Error = fmt.Errorf("unsupported value: %f", __obf_efaf2719b44ad8ac)
		return
	}
	if __obf_efaf2719b44ad8ac < 0 {
		__obf_8a2c51fe22775655.__obf_41130daa346c0e53('-')
		__obf_efaf2719b44ad8ac = -__obf_efaf2719b44ad8ac
	}
	if __obf_efaf2719b44ad8ac > 0x4ffffff {
		__obf_8a2c51fe22775655.
			WriteFloat32(__obf_efaf2719b44ad8ac)
		return
	}
	__obf_95dfcd037d5464a2 := 6
	__obf_da498eac17ea869c := uint64(1000000)
	__obf_a629f4b0bf305ec3 := // 6
		uint64(float64(__obf_efaf2719b44ad8ac)*float64(__obf_da498eac17ea869c) + 0.5)
	__obf_8a2c51fe22775655.
		WriteUint64(__obf_a629f4b0bf305ec3 / __obf_da498eac17ea869c)
	__obf_3b4d64f48a6770cb := __obf_a629f4b0bf305ec3 % __obf_da498eac17ea869c
	if __obf_3b4d64f48a6770cb == 0 {
		return
	}
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('.')
	for __obf_23ceea01a3d73f6a := __obf_95dfcd037d5464a2 - 1; __obf_23ceea01a3d73f6a > 0 && __obf_3b4d64f48a6770cb < __obf_faa2ce64bf143410[__obf_23ceea01a3d73f6a]; __obf_23ceea01a3d73f6a-- {
		__obf_8a2c51fe22775655.__obf_41130daa346c0e53('0')
	}
	__obf_8a2c51fe22775655.
		WriteUint64(__obf_3b4d64f48a6770cb)
	for __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[len(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df)-1] == '0' {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[:len(__obf_8a2c51fe22775655.

			// WriteFloat64 write float64 to stream
			__obf_0b1656d7ef4bd9df)-1]
	}
}

func (__obf_8a2c51fe22775655 *Stream) WriteFloat64(__obf_efaf2719b44ad8ac float64) {
	if math.IsInf(__obf_efaf2719b44ad8ac, 0) || math.IsNaN(__obf_efaf2719b44ad8ac) {
		__obf_8a2c51fe22775655.
			Error = fmt.Errorf("unsupported value: %f", __obf_efaf2719b44ad8ac)
		return
	}
	__obf_aaded2753c66e327 := math.Abs(__obf_efaf2719b44ad8ac)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if __obf_aaded2753c66e327 != 0 {
		if __obf_aaded2753c66e327 < 1e-6 || __obf_aaded2753c66e327 >= 1e21 {
			fmt = 'e'
		}
	}
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = strconv.AppendFloat(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, float64(__obf_efaf2719b44ad8ac), fmt, -1, 64)
	if fmt == 'e' {
		__obf_e40b3963a92ca4a0 := // clean up e-09 to e-9
			len(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df)
		if __obf_e40b3963a92ca4a0 >= 4 && __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[__obf_e40b3963a92ca4a0-4] == 'e' && __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[__obf_e40b3963a92ca4a0-3] == '-' && __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[__obf_e40b3963a92ca4a0-2] == '0' {
			__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[__obf_e40b3963a92ca4a0-2] = __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[__obf_e40b3963a92ca4a0-1]
			__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[:__obf_e40b3963a92ca4a0-1]
		}
	}
}

// WriteFloat64Lossy write float64 to stream with ONLY 6 digits precision although much much faster
func (__obf_8a2c51fe22775655 *Stream) WriteFloat64Lossy(__obf_efaf2719b44ad8ac float64) {
	if math.IsInf(__obf_efaf2719b44ad8ac, 0) || math.IsNaN(__obf_efaf2719b44ad8ac) {
		__obf_8a2c51fe22775655.
			Error = fmt.Errorf("unsupported value: %f", __obf_efaf2719b44ad8ac)
		return
	}
	if __obf_efaf2719b44ad8ac < 0 {
		__obf_8a2c51fe22775655.__obf_41130daa346c0e53('-')
		__obf_efaf2719b44ad8ac = -__obf_efaf2719b44ad8ac
	}
	if __obf_efaf2719b44ad8ac > 0x4ffffff {
		__obf_8a2c51fe22775655.
			WriteFloat64(__obf_efaf2719b44ad8ac)
		return
	}
	__obf_95dfcd037d5464a2 := 6
	__obf_da498eac17ea869c := uint64(1000000)
	__obf_a629f4b0bf305ec3 := // 6
		uint64(__obf_efaf2719b44ad8ac*float64(__obf_da498eac17ea869c) + 0.5)
	__obf_8a2c51fe22775655.
		WriteUint64(__obf_a629f4b0bf305ec3 / __obf_da498eac17ea869c)
	__obf_3b4d64f48a6770cb := __obf_a629f4b0bf305ec3 % __obf_da498eac17ea869c
	if __obf_3b4d64f48a6770cb == 0 {
		return
	}
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('.')
	for __obf_23ceea01a3d73f6a := __obf_95dfcd037d5464a2 - 1; __obf_23ceea01a3d73f6a > 0 && __obf_3b4d64f48a6770cb < __obf_faa2ce64bf143410[__obf_23ceea01a3d73f6a]; __obf_23ceea01a3d73f6a-- {
		__obf_8a2c51fe22775655.__obf_41130daa346c0e53('0')
	}
	__obf_8a2c51fe22775655.
		WriteUint64(__obf_3b4d64f48a6770cb)
	for __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[len(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df)-1] == '0' {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[:len(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df)-1]
	}
}
