package __obf_5b802ce8d9ba56d6

import (
	"math"
	"strconv"
)

var __obf_268306cf92be8be7 []int8

const __obf_42a1f733df7ba815 = uint32(0xffffffff)/10 - 1
const __obf_c6c8cab5bfa05789 = uint64(0xffffffffffffffff)/10 - 1
const __obf_d004373e847514ce = 1<<53 - 1

func init() {
	__obf_268306cf92be8be7 = make([]int8, 256)
	for __obf_2deec7c38ffb6ae3 := 0; __obf_2deec7c38ffb6ae3 < len(__obf_268306cf92be8be7); __obf_2deec7c38ffb6ae3++ {
		__obf_268306cf92be8be7[__obf_2deec7c38ffb6ae3] = __obf_22cd644f84c3de2e
	}
	for __obf_2deec7c38ffb6ae3 := int8('0'); __obf_2deec7c38ffb6ae3 <= int8('9'); __obf_2deec7c38ffb6ae3++ {
		__obf_268306cf92be8be7[__obf_2deec7c38ffb6ae3] = __obf_2deec7c38ffb6ae3 - int8('0')
	}
}

// ReadUint read uint
func (__obf_67008a6a9e5ba828 *Iterator) ReadUint() uint {
	if strconv.IntSize == 32 {
		return uint(__obf_67008a6a9e5ba828.ReadUint32())
	}
	return uint(__obf_67008a6a9e5ba828.ReadUint64())
}

// ReadInt read int
func (__obf_67008a6a9e5ba828 *Iterator) ReadInt() int {
	if strconv.IntSize == 32 {
		return int(__obf_67008a6a9e5ba828.ReadInt32())
	}
	return int(__obf_67008a6a9e5ba828.ReadInt64())
}

// ReadInt8 read int8
func (__obf_67008a6a9e5ba828 *Iterator) ReadInt8() (__obf_5dabcdfee5097ed6 int8) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == '-' {
		__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.__obf_4b424a27cc271ec1(__obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb())
		if __obf_5406b11e33b9d1d3 > math.MaxInt8+1 {
			__obf_67008a6a9e5ba828.
				ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_5406b11e33b9d1d3), 10))
			return
		}
		return -int8(__obf_5406b11e33b9d1d3)
	}
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.__obf_4b424a27cc271ec1(__obf_dab9baaadfa7c8c2)
	if __obf_5406b11e33b9d1d3 > math.MaxInt8 {
		__obf_67008a6a9e5ba828.
			ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_5406b11e33b9d1d3), 10))
		return
	}
	return int8(__obf_5406b11e33b9d1d3)
}

// ReadUint8 read uint8
func (__obf_67008a6a9e5ba828 *Iterator) ReadUint8() (__obf_5dabcdfee5097ed6 uint8) {
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.__obf_4b424a27cc271ec1(__obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490())
	if __obf_5406b11e33b9d1d3 > math.MaxUint8 {
		__obf_67008a6a9e5ba828.
			ReportError("ReadUint8", "overflow: "+strconv.FormatInt(int64(__obf_5406b11e33b9d1d3), 10))
		return
	}
	return uint8(__obf_5406b11e33b9d1d3)
}

// ReadInt16 read int16
func (__obf_67008a6a9e5ba828 *Iterator) ReadInt16() (__obf_5dabcdfee5097ed6 int16) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == '-' {
		__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.__obf_4b424a27cc271ec1(__obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb())
		if __obf_5406b11e33b9d1d3 > math.MaxInt16+1 {
			__obf_67008a6a9e5ba828.
				ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_5406b11e33b9d1d3), 10))
			return
		}
		return -int16(__obf_5406b11e33b9d1d3)
	}
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.__obf_4b424a27cc271ec1(__obf_dab9baaadfa7c8c2)
	if __obf_5406b11e33b9d1d3 > math.MaxInt16 {
		__obf_67008a6a9e5ba828.
			ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_5406b11e33b9d1d3), 10))
		return
	}
	return int16(__obf_5406b11e33b9d1d3)
}

// ReadUint16 read uint16
func (__obf_67008a6a9e5ba828 *Iterator) ReadUint16() (__obf_5dabcdfee5097ed6 uint16) {
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.__obf_4b424a27cc271ec1(__obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490())
	if __obf_5406b11e33b9d1d3 > math.MaxUint16 {
		__obf_67008a6a9e5ba828.
			ReportError("ReadUint16", "overflow: "+strconv.FormatInt(int64(__obf_5406b11e33b9d1d3), 10))
		return
	}
	return uint16(__obf_5406b11e33b9d1d3)
}

// ReadInt32 read int32
func (__obf_67008a6a9e5ba828 *Iterator) ReadInt32() (__obf_5dabcdfee5097ed6 int32) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == '-' {
		__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.__obf_4b424a27cc271ec1(__obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb())
		if __obf_5406b11e33b9d1d3 > math.MaxInt32+1 {
			__obf_67008a6a9e5ba828.
				ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_5406b11e33b9d1d3), 10))
			return
		}
		return -int32(__obf_5406b11e33b9d1d3)
	}
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.__obf_4b424a27cc271ec1(__obf_dab9baaadfa7c8c2)
	if __obf_5406b11e33b9d1d3 > math.MaxInt32 {
		__obf_67008a6a9e5ba828.
			ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_5406b11e33b9d1d3), 10))
		return
	}
	return int32(__obf_5406b11e33b9d1d3)
}

// ReadUint32 read uint32
func (__obf_67008a6a9e5ba828 *Iterator) ReadUint32() (__obf_5dabcdfee5097ed6 uint32) {
	return __obf_67008a6a9e5ba828.__obf_4b424a27cc271ec1(__obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490())
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_4b424a27cc271ec1(__obf_dab9baaadfa7c8c2 byte) (__obf_5dabcdfee5097ed6 uint32) {
	__obf_0664edf071899d0b := __obf_268306cf92be8be7[__obf_dab9baaadfa7c8c2]
	if __obf_0664edf071899d0b == 0 {
		__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
		return 0 // single zero
	}
	if __obf_0664edf071899d0b == __obf_22cd644f84c3de2e {
		__obf_67008a6a9e5ba828.
			ReportError("readUint32", "unexpected character: "+string([]byte{byte(__obf_0664edf071899d0b)}))
		return
	}
	__obf_c949477a4af2082d := uint32(__obf_0664edf071899d0b)
	if __obf_67008a6a9e5ba828.__obf_3a36550914545c79-__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 > 10 {
		__obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36
		__obf_4f1e52c7f025e5c6 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_4f1e52c7f025e5c6 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d
		}
		__obf_2deec7c38ffb6ae3++
		__obf_1b7da151b9b467d3 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_1b7da151b9b467d3 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*10 + uint32(__obf_4f1e52c7f025e5c6)
		}
		__obf_2deec7c38ffb6ae3++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_9088051f0136bec7 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_9088051f0136bec7 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*100 + uint32(__obf_4f1e52c7f025e5c6)*10 + uint32(__obf_1b7da151b9b467d3)
		}
		__obf_2deec7c38ffb6ae3++
		__obf_e0ae2c842c6d3400 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_e0ae2c842c6d3400 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*1000 + uint32(__obf_4f1e52c7f025e5c6)*100 + uint32(__obf_1b7da151b9b467d3)*10 + uint32(__obf_9088051f0136bec7)
		}
		__obf_2deec7c38ffb6ae3++
		__obf_917017677c93405d := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_917017677c93405d == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*10000 + uint32(__obf_4f1e52c7f025e5c6)*1000 + uint32(__obf_1b7da151b9b467d3)*100 + uint32(__obf_9088051f0136bec7)*10 + uint32(__obf_e0ae2c842c6d3400)
		}
		__obf_2deec7c38ffb6ae3++
		__obf_19d911a3e7a134b7 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_19d911a3e7a134b7 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*100000 + uint32(__obf_4f1e52c7f025e5c6)*10000 + uint32(__obf_1b7da151b9b467d3)*1000 + uint32(__obf_9088051f0136bec7)*100 + uint32(__obf_e0ae2c842c6d3400)*10 + uint32(__obf_917017677c93405d)
		}
		__obf_2deec7c38ffb6ae3++
		__obf_947928a3f46bbdf8 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_947928a3f46bbdf8 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*1000000 + uint32(__obf_4f1e52c7f025e5c6)*100000 + uint32(__obf_1b7da151b9b467d3)*10000 + uint32(__obf_9088051f0136bec7)*1000 + uint32(__obf_e0ae2c842c6d3400)*100 + uint32(__obf_917017677c93405d)*10 + uint32(__obf_19d911a3e7a134b7)
		}
		__obf_2deec7c38ffb6ae3++
		__obf_3b32bd1953670d04 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		__obf_c949477a4af2082d = __obf_c949477a4af2082d*10000000 + uint32(__obf_4f1e52c7f025e5c6)*1000000 + uint32(__obf_1b7da151b9b467d3)*100000 + uint32(__obf_9088051f0136bec7)*10000 + uint32(__obf_e0ae2c842c6d3400)*1000 + uint32(__obf_917017677c93405d)*100 + uint32(__obf_19d911a3e7a134b7)*10 + uint32(__obf_947928a3f46bbdf8)
		__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
		if __obf_3b32bd1953670d04 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d
		}
	}
	for {
		for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
			__obf_0664edf071899d0b = __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
			if __obf_0664edf071899d0b == __obf_22cd644f84c3de2e {
				__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
				__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
				return __obf_c949477a4af2082d
			}
			if __obf_c949477a4af2082d > __obf_42a1f733df7ba815 {
				__obf_92be7ee8cfb531c6 := (__obf_c949477a4af2082d << 3) + (__obf_c949477a4af2082d << 1) + uint32(__obf_0664edf071899d0b)
				if __obf_92be7ee8cfb531c6 < __obf_c949477a4af2082d {
					__obf_67008a6a9e5ba828.
						ReportError("readUint32", "overflow")
					return
				}
				__obf_c949477a4af2082d = __obf_92be7ee8cfb531c6
				continue
			}
			__obf_c949477a4af2082d = (__obf_c949477a4af2082d << 3) + (__obf_c949477a4af2082d << 1) + uint32(__obf_0664edf071899d0b)
		}
		if !__obf_67008a6a9e5ba828.__obf_538f3d3337e0395f() {
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d
		}
	}
}

// ReadInt64 read int64
func (__obf_67008a6a9e5ba828 *Iterator) ReadInt64() (__obf_5dabcdfee5097ed6 int64) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == '-' {
		__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.__obf_7d605b53e8331536(__obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb())
		if __obf_5406b11e33b9d1d3 > math.MaxInt64+1 {
			__obf_67008a6a9e5ba828.
				ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_5406b11e33b9d1d3), 10))
			return
		}
		return -int64(__obf_5406b11e33b9d1d3)
	}
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.__obf_7d605b53e8331536(__obf_dab9baaadfa7c8c2)
	if __obf_5406b11e33b9d1d3 > math.MaxInt64 {
		__obf_67008a6a9e5ba828.
			ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_5406b11e33b9d1d3), 10))
		return
	}
	return int64(__obf_5406b11e33b9d1d3)
}

// ReadUint64 read uint64
func (__obf_67008a6a9e5ba828 *Iterator) ReadUint64() uint64 {
	return __obf_67008a6a9e5ba828.__obf_7d605b53e8331536(__obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490())
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_7d605b53e8331536(__obf_dab9baaadfa7c8c2 byte) (__obf_5dabcdfee5097ed6 uint64) {
	__obf_0664edf071899d0b := __obf_268306cf92be8be7[__obf_dab9baaadfa7c8c2]
	if __obf_0664edf071899d0b == 0 {
		__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
		return 0 // single zero
	}
	if __obf_0664edf071899d0b == __obf_22cd644f84c3de2e {
		__obf_67008a6a9e5ba828.
			ReportError("readUint64", "unexpected character: "+string([]byte{byte(__obf_0664edf071899d0b)}))
		return
	}
	__obf_c949477a4af2082d := uint64(__obf_0664edf071899d0b)
	if __obf_67008a6a9e5ba828.__obf_3a36550914545c79-__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 > 10 {
		__obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36
		__obf_4f1e52c7f025e5c6 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_4f1e52c7f025e5c6 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d
		}
		__obf_2deec7c38ffb6ae3++
		__obf_1b7da151b9b467d3 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_1b7da151b9b467d3 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*10 + uint64(__obf_4f1e52c7f025e5c6)
		}
		__obf_2deec7c38ffb6ae3++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_9088051f0136bec7 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_9088051f0136bec7 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*100 + uint64(__obf_4f1e52c7f025e5c6)*10 + uint64(__obf_1b7da151b9b467d3)
		}
		__obf_2deec7c38ffb6ae3++
		__obf_e0ae2c842c6d3400 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_e0ae2c842c6d3400 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*1000 + uint64(__obf_4f1e52c7f025e5c6)*100 + uint64(__obf_1b7da151b9b467d3)*10 + uint64(__obf_9088051f0136bec7)
		}
		__obf_2deec7c38ffb6ae3++
		__obf_917017677c93405d := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_917017677c93405d == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*10000 + uint64(__obf_4f1e52c7f025e5c6)*1000 + uint64(__obf_1b7da151b9b467d3)*100 + uint64(__obf_9088051f0136bec7)*10 + uint64(__obf_e0ae2c842c6d3400)
		}
		__obf_2deec7c38ffb6ae3++
		__obf_19d911a3e7a134b7 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_19d911a3e7a134b7 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*100000 + uint64(__obf_4f1e52c7f025e5c6)*10000 + uint64(__obf_1b7da151b9b467d3)*1000 + uint64(__obf_9088051f0136bec7)*100 + uint64(__obf_e0ae2c842c6d3400)*10 + uint64(__obf_917017677c93405d)
		}
		__obf_2deec7c38ffb6ae3++
		__obf_947928a3f46bbdf8 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		if __obf_947928a3f46bbdf8 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d*1000000 + uint64(__obf_4f1e52c7f025e5c6)*100000 + uint64(__obf_1b7da151b9b467d3)*10000 + uint64(__obf_9088051f0136bec7)*1000 + uint64(__obf_e0ae2c842c6d3400)*100 + uint64(__obf_917017677c93405d)*10 + uint64(__obf_19d911a3e7a134b7)
		}
		__obf_2deec7c38ffb6ae3++
		__obf_3b32bd1953670d04 := __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
		__obf_c949477a4af2082d = __obf_c949477a4af2082d*10000000 + uint64(__obf_4f1e52c7f025e5c6)*1000000 + uint64(__obf_1b7da151b9b467d3)*100000 + uint64(__obf_9088051f0136bec7)*10000 + uint64(__obf_e0ae2c842c6d3400)*1000 + uint64(__obf_917017677c93405d)*100 + uint64(__obf_19d911a3e7a134b7)*10 + uint64(__obf_947928a3f46bbdf8)
		__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
		if __obf_3b32bd1953670d04 == __obf_22cd644f84c3de2e {
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d
		}
	}
	for {
		for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
			__obf_0664edf071899d0b = __obf_268306cf92be8be7[__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]]
			if __obf_0664edf071899d0b == __obf_22cd644f84c3de2e {
				__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
				__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
				return __obf_c949477a4af2082d
			}
			if __obf_c949477a4af2082d > __obf_c6c8cab5bfa05789 {
				__obf_92be7ee8cfb531c6 := (__obf_c949477a4af2082d << 3) + (__obf_c949477a4af2082d << 1) + uint64(__obf_0664edf071899d0b)
				if __obf_92be7ee8cfb531c6 < __obf_c949477a4af2082d {
					__obf_67008a6a9e5ba828.
						ReportError("readUint64", "overflow")
					return
				}
				__obf_c949477a4af2082d = __obf_92be7ee8cfb531c6
				continue
			}
			__obf_c949477a4af2082d = (__obf_c949477a4af2082d << 3) + (__obf_c949477a4af2082d << 1) + uint64(__obf_0664edf071899d0b)
		}
		if !__obf_67008a6a9e5ba828.__obf_538f3d3337e0395f() {
			__obf_67008a6a9e5ba828.__obf_66c3e3ce61a617f3()
			return __obf_c949477a4af2082d
		}
	}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_66c3e3ce61a617f3() {
	if __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79 && __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36] == '.' {
		__obf_67008a6a9e5ba828.
			ReportError("assertInteger", "can not decode float as int")
	}
}
