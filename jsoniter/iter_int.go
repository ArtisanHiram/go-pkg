package __obf_703d23ba520c3295

import (
	"math"
	"strconv"
)

var __obf_ec71c909fd507c89 []int8

const __obf_f32cee63cb58e892 = uint32(0xffffffff)/10 - 1
const __obf_470c0669890e2df4 = uint64(0xffffffffffffffff)/10 - 1
const __obf_df0c999320bb22b2 = 1<<53 - 1

func init() {
	__obf_ec71c909fd507c89 = make([]int8, 256)
	for __obf_b0a5d2bd48690f1d := 0; __obf_b0a5d2bd48690f1d < len(__obf_ec71c909fd507c89); __obf_b0a5d2bd48690f1d++ {
		__obf_ec71c909fd507c89[__obf_b0a5d2bd48690f1d] = __obf_0712b1d6d67b066b
	}
	for __obf_b0a5d2bd48690f1d := int8('0'); __obf_b0a5d2bd48690f1d <= int8('9'); __obf_b0a5d2bd48690f1d++ {
		__obf_ec71c909fd507c89[__obf_b0a5d2bd48690f1d] = __obf_b0a5d2bd48690f1d - int8('0')
	}
}

// ReadUint read uint
func (__obf_47edab4c16a2d88d *Iterator) ReadUint() uint {
	if strconv.IntSize == 32 {
		return uint(__obf_47edab4c16a2d88d.ReadUint32())
	}
	return uint(__obf_47edab4c16a2d88d.ReadUint64())
}

// ReadInt read int
func (__obf_47edab4c16a2d88d *Iterator) ReadInt() int {
	if strconv.IntSize == 32 {
		return int(__obf_47edab4c16a2d88d.ReadInt32())
	}
	return int(__obf_47edab4c16a2d88d.ReadInt64())
}

// ReadInt8 read int8
func (__obf_47edab4c16a2d88d *Iterator) ReadInt8() (__obf_551cbec38242ce0c int8) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == '-' {
		__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.__obf_d8d78c45df386ff3(__obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3())
		if __obf_b924538fffe5fd64 > math.MaxInt8+1 {
			__obf_47edab4c16a2d88d.
				ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_b924538fffe5fd64), 10))
			return
		}
		return -int8(__obf_b924538fffe5fd64)
	}
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.__obf_d8d78c45df386ff3(__obf_bd08f5161fff294a)
	if __obf_b924538fffe5fd64 > math.MaxInt8 {
		__obf_47edab4c16a2d88d.
			ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_b924538fffe5fd64), 10))
		return
	}
	return int8(__obf_b924538fffe5fd64)
}

// ReadUint8 read uint8
func (__obf_47edab4c16a2d88d *Iterator) ReadUint8() (__obf_551cbec38242ce0c uint8) {
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.__obf_d8d78c45df386ff3(__obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec())
	if __obf_b924538fffe5fd64 > math.MaxUint8 {
		__obf_47edab4c16a2d88d.
			ReportError("ReadUint8", "overflow: "+strconv.FormatInt(int64(__obf_b924538fffe5fd64), 10))
		return
	}
	return uint8(__obf_b924538fffe5fd64)
}

// ReadInt16 read int16
func (__obf_47edab4c16a2d88d *Iterator) ReadInt16() (__obf_551cbec38242ce0c int16) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == '-' {
		__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.__obf_d8d78c45df386ff3(__obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3())
		if __obf_b924538fffe5fd64 > math.MaxInt16+1 {
			__obf_47edab4c16a2d88d.
				ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_b924538fffe5fd64), 10))
			return
		}
		return -int16(__obf_b924538fffe5fd64)
	}
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.__obf_d8d78c45df386ff3(__obf_bd08f5161fff294a)
	if __obf_b924538fffe5fd64 > math.MaxInt16 {
		__obf_47edab4c16a2d88d.
			ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_b924538fffe5fd64), 10))
		return
	}
	return int16(__obf_b924538fffe5fd64)
}

// ReadUint16 read uint16
func (__obf_47edab4c16a2d88d *Iterator) ReadUint16() (__obf_551cbec38242ce0c uint16) {
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.__obf_d8d78c45df386ff3(__obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec())
	if __obf_b924538fffe5fd64 > math.MaxUint16 {
		__obf_47edab4c16a2d88d.
			ReportError("ReadUint16", "overflow: "+strconv.FormatInt(int64(__obf_b924538fffe5fd64), 10))
		return
	}
	return uint16(__obf_b924538fffe5fd64)
}

// ReadInt32 read int32
func (__obf_47edab4c16a2d88d *Iterator) ReadInt32() (__obf_551cbec38242ce0c int32) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == '-' {
		__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.__obf_d8d78c45df386ff3(__obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3())
		if __obf_b924538fffe5fd64 > math.MaxInt32+1 {
			__obf_47edab4c16a2d88d.
				ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_b924538fffe5fd64), 10))
			return
		}
		return -int32(__obf_b924538fffe5fd64)
	}
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.__obf_d8d78c45df386ff3(__obf_bd08f5161fff294a)
	if __obf_b924538fffe5fd64 > math.MaxInt32 {
		__obf_47edab4c16a2d88d.
			ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_b924538fffe5fd64), 10))
		return
	}
	return int32(__obf_b924538fffe5fd64)
}

// ReadUint32 read uint32
func (__obf_47edab4c16a2d88d *Iterator) ReadUint32() (__obf_551cbec38242ce0c uint32) {
	return __obf_47edab4c16a2d88d.__obf_d8d78c45df386ff3(__obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec())
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_d8d78c45df386ff3(__obf_bd08f5161fff294a byte) (__obf_551cbec38242ce0c uint32) {
	__obf_562d42737cc3cc4f := __obf_ec71c909fd507c89[__obf_bd08f5161fff294a]
	if __obf_562d42737cc3cc4f == 0 {
		__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
		return 0 // single zero
	}
	if __obf_562d42737cc3cc4f == __obf_0712b1d6d67b066b {
		__obf_47edab4c16a2d88d.
			ReportError("readUint32", "unexpected character: "+string([]byte{byte(__obf_562d42737cc3cc4f)}))
		return
	}
	__obf_ccb7d8cb07a5350f := uint32(__obf_562d42737cc3cc4f)
	if __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5-__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 > 10 {
		__obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11
		__obf_328476aaeaf4a16b := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_328476aaeaf4a16b == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f
		}
		__obf_b0a5d2bd48690f1d++
		__obf_ba441167005168ac := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_ba441167005168ac == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*10 + uint32(__obf_328476aaeaf4a16b)
		}
		__obf_b0a5d2bd48690f1d++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_e021b5c97986a208 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_e021b5c97986a208 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*100 + uint32(__obf_328476aaeaf4a16b)*10 + uint32(__obf_ba441167005168ac)
		}
		__obf_b0a5d2bd48690f1d++
		__obf_872da0e063d17a41 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_872da0e063d17a41 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*1000 + uint32(__obf_328476aaeaf4a16b)*100 + uint32(__obf_ba441167005168ac)*10 + uint32(__obf_e021b5c97986a208)
		}
		__obf_b0a5d2bd48690f1d++
		__obf_c547f94c64471e70 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_c547f94c64471e70 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*10000 + uint32(__obf_328476aaeaf4a16b)*1000 + uint32(__obf_ba441167005168ac)*100 + uint32(__obf_e021b5c97986a208)*10 + uint32(__obf_872da0e063d17a41)
		}
		__obf_b0a5d2bd48690f1d++
		__obf_a3e875ed7062d3f5 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_a3e875ed7062d3f5 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*100000 + uint32(__obf_328476aaeaf4a16b)*10000 + uint32(__obf_ba441167005168ac)*1000 + uint32(__obf_e021b5c97986a208)*100 + uint32(__obf_872da0e063d17a41)*10 + uint32(__obf_c547f94c64471e70)
		}
		__obf_b0a5d2bd48690f1d++
		__obf_87230d676bf41984 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_87230d676bf41984 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*1000000 + uint32(__obf_328476aaeaf4a16b)*100000 + uint32(__obf_ba441167005168ac)*10000 + uint32(__obf_e021b5c97986a208)*1000 + uint32(__obf_872da0e063d17a41)*100 + uint32(__obf_c547f94c64471e70)*10 + uint32(__obf_a3e875ed7062d3f5)
		}
		__obf_b0a5d2bd48690f1d++
		__obf_4f005b0882be4b84 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		__obf_ccb7d8cb07a5350f = __obf_ccb7d8cb07a5350f*10000000 + uint32(__obf_328476aaeaf4a16b)*1000000 + uint32(__obf_ba441167005168ac)*100000 + uint32(__obf_e021b5c97986a208)*10000 + uint32(__obf_872da0e063d17a41)*1000 + uint32(__obf_c547f94c64471e70)*100 + uint32(__obf_a3e875ed7062d3f5)*10 + uint32(__obf_87230d676bf41984)
		__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
		if __obf_4f005b0882be4b84 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f
		}
	}
	for {
		for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
			__obf_562d42737cc3cc4f = __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
			if __obf_562d42737cc3cc4f == __obf_0712b1d6d67b066b {
				__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
				__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
				return __obf_ccb7d8cb07a5350f
			}
			if __obf_ccb7d8cb07a5350f > __obf_f32cee63cb58e892 {
				__obf_5b5d2a16fdb3750e := (__obf_ccb7d8cb07a5350f << 3) + (__obf_ccb7d8cb07a5350f << 1) + uint32(__obf_562d42737cc3cc4f)
				if __obf_5b5d2a16fdb3750e < __obf_ccb7d8cb07a5350f {
					__obf_47edab4c16a2d88d.
						ReportError("readUint32", "overflow")
					return
				}
				__obf_ccb7d8cb07a5350f = __obf_5b5d2a16fdb3750e
				continue
			}
			__obf_ccb7d8cb07a5350f = (__obf_ccb7d8cb07a5350f << 3) + (__obf_ccb7d8cb07a5350f << 1) + uint32(__obf_562d42737cc3cc4f)
		}
		if !__obf_47edab4c16a2d88d.__obf_e147dd47ce22b671() {
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f
		}
	}
}

// ReadInt64 read int64
func (__obf_47edab4c16a2d88d *Iterator) ReadInt64() (__obf_551cbec38242ce0c int64) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == '-' {
		__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.__obf_c59d2fcd957143fb(__obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3())
		if __obf_b924538fffe5fd64 > math.MaxInt64+1 {
			__obf_47edab4c16a2d88d.
				ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_b924538fffe5fd64), 10))
			return
		}
		return -int64(__obf_b924538fffe5fd64)
	}
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.__obf_c59d2fcd957143fb(__obf_bd08f5161fff294a)
	if __obf_b924538fffe5fd64 > math.MaxInt64 {
		__obf_47edab4c16a2d88d.
			ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_b924538fffe5fd64), 10))
		return
	}
	return int64(__obf_b924538fffe5fd64)
}

// ReadUint64 read uint64
func (__obf_47edab4c16a2d88d *Iterator) ReadUint64() uint64 {
	return __obf_47edab4c16a2d88d.__obf_c59d2fcd957143fb(__obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec())
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_c59d2fcd957143fb(__obf_bd08f5161fff294a byte) (__obf_551cbec38242ce0c uint64) {
	__obf_562d42737cc3cc4f := __obf_ec71c909fd507c89[__obf_bd08f5161fff294a]
	if __obf_562d42737cc3cc4f == 0 {
		__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
		return 0 // single zero
	}
	if __obf_562d42737cc3cc4f == __obf_0712b1d6d67b066b {
		__obf_47edab4c16a2d88d.
			ReportError("readUint64", "unexpected character: "+string([]byte{byte(__obf_562d42737cc3cc4f)}))
		return
	}
	__obf_ccb7d8cb07a5350f := uint64(__obf_562d42737cc3cc4f)
	if __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5-__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 > 10 {
		__obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11
		__obf_328476aaeaf4a16b := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_328476aaeaf4a16b == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f
		}
		__obf_b0a5d2bd48690f1d++
		__obf_ba441167005168ac := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_ba441167005168ac == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*10 + uint64(__obf_328476aaeaf4a16b)
		}
		__obf_b0a5d2bd48690f1d++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_e021b5c97986a208 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_e021b5c97986a208 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*100 + uint64(__obf_328476aaeaf4a16b)*10 + uint64(__obf_ba441167005168ac)
		}
		__obf_b0a5d2bd48690f1d++
		__obf_872da0e063d17a41 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_872da0e063d17a41 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*1000 + uint64(__obf_328476aaeaf4a16b)*100 + uint64(__obf_ba441167005168ac)*10 + uint64(__obf_e021b5c97986a208)
		}
		__obf_b0a5d2bd48690f1d++
		__obf_c547f94c64471e70 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_c547f94c64471e70 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*10000 + uint64(__obf_328476aaeaf4a16b)*1000 + uint64(__obf_ba441167005168ac)*100 + uint64(__obf_e021b5c97986a208)*10 + uint64(__obf_872da0e063d17a41)
		}
		__obf_b0a5d2bd48690f1d++
		__obf_a3e875ed7062d3f5 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_a3e875ed7062d3f5 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*100000 + uint64(__obf_328476aaeaf4a16b)*10000 + uint64(__obf_ba441167005168ac)*1000 + uint64(__obf_e021b5c97986a208)*100 + uint64(__obf_872da0e063d17a41)*10 + uint64(__obf_c547f94c64471e70)
		}
		__obf_b0a5d2bd48690f1d++
		__obf_87230d676bf41984 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		if __obf_87230d676bf41984 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f*1000000 + uint64(__obf_328476aaeaf4a16b)*100000 + uint64(__obf_ba441167005168ac)*10000 + uint64(__obf_e021b5c97986a208)*1000 + uint64(__obf_872da0e063d17a41)*100 + uint64(__obf_c547f94c64471e70)*10 + uint64(__obf_a3e875ed7062d3f5)
		}
		__obf_b0a5d2bd48690f1d++
		__obf_4f005b0882be4b84 := __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
		__obf_ccb7d8cb07a5350f = __obf_ccb7d8cb07a5350f*10000000 + uint64(__obf_328476aaeaf4a16b)*1000000 + uint64(__obf_ba441167005168ac)*100000 + uint64(__obf_e021b5c97986a208)*10000 + uint64(__obf_872da0e063d17a41)*1000 + uint64(__obf_c547f94c64471e70)*100 + uint64(__obf_a3e875ed7062d3f5)*10 + uint64(__obf_87230d676bf41984)
		__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
		if __obf_4f005b0882be4b84 == __obf_0712b1d6d67b066b {
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f
		}
	}
	for {
		for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
			__obf_562d42737cc3cc4f = __obf_ec71c909fd507c89[__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]]
			if __obf_562d42737cc3cc4f == __obf_0712b1d6d67b066b {
				__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
				__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
				return __obf_ccb7d8cb07a5350f
			}
			if __obf_ccb7d8cb07a5350f > __obf_470c0669890e2df4 {
				__obf_5b5d2a16fdb3750e := (__obf_ccb7d8cb07a5350f << 3) + (__obf_ccb7d8cb07a5350f << 1) + uint64(__obf_562d42737cc3cc4f)
				if __obf_5b5d2a16fdb3750e < __obf_ccb7d8cb07a5350f {
					__obf_47edab4c16a2d88d.
						ReportError("readUint64", "overflow")
					return
				}
				__obf_ccb7d8cb07a5350f = __obf_5b5d2a16fdb3750e
				continue
			}
			__obf_ccb7d8cb07a5350f = (__obf_ccb7d8cb07a5350f << 3) + (__obf_ccb7d8cb07a5350f << 1) + uint64(__obf_562d42737cc3cc4f)
		}
		if !__obf_47edab4c16a2d88d.__obf_e147dd47ce22b671() {
			__obf_47edab4c16a2d88d.__obf_21da22c62f856f82()
			return __obf_ccb7d8cb07a5350f
		}
	}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_21da22c62f856f82() {
	if __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 && __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11] == '.' {
		__obf_47edab4c16a2d88d.
			ReportError("assertInteger", "can not decode float as int")
	}
}
