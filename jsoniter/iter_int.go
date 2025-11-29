package __obf_91620b895eeff9ed

import (
	"math"
	"strconv"
)

var __obf_56c7464cf1f5a110 []int8

const __obf_631720670e71a3cf = uint32(0xffffffff)/10 - 1
const __obf_fea4c6f6b5068026 = uint64(0xffffffffffffffff)/10 - 1
const __obf_397fd3fe0942535f = 1<<53 - 1

func init() {
	__obf_56c7464cf1f5a110 = make([]int8, 256)
	for __obf_5aa5c8829b97f182 := 0; __obf_5aa5c8829b97f182 < len(__obf_56c7464cf1f5a110); __obf_5aa5c8829b97f182++ {
		__obf_56c7464cf1f5a110[__obf_5aa5c8829b97f182] = __obf_17acfd2d0998ee00
	}
	for __obf_5aa5c8829b97f182 := int8('0'); __obf_5aa5c8829b97f182 <= int8('9'); __obf_5aa5c8829b97f182++ {
		__obf_56c7464cf1f5a110[__obf_5aa5c8829b97f182] = __obf_5aa5c8829b97f182 - int8('0')
	}
}

// ReadUint read uint
func (__obf_1bb30e8a74ed8233 *Iterator) ReadUint() uint {
	if strconv.IntSize == 32 {
		return uint(__obf_1bb30e8a74ed8233.ReadUint32())
	}
	return uint(__obf_1bb30e8a74ed8233.ReadUint64())
}

// ReadInt read int
func (__obf_1bb30e8a74ed8233 *Iterator) ReadInt() int {
	if strconv.IntSize == 32 {
		return int(__obf_1bb30e8a74ed8233.ReadInt32())
	}
	return int(__obf_1bb30e8a74ed8233.ReadInt64())
}

// ReadInt8 read int8
func (__obf_1bb30e8a74ed8233 *Iterator) ReadInt8() (__obf_e46f5fe3db5036fe int8) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == '-' {
		__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.__obf_f735ee00ef3fae0b(__obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc())
		if __obf_bbfd2ac8618a6f0c > math.MaxInt8+1 {
			__obf_1bb30e8a74ed8233.
				ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_bbfd2ac8618a6f0c), 10))
			return
		}
		return -int8(__obf_bbfd2ac8618a6f0c)
	}
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.__obf_f735ee00ef3fae0b(__obf_f16b4157911bc9af)
	if __obf_bbfd2ac8618a6f0c > math.MaxInt8 {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_bbfd2ac8618a6f0c), 10))
		return
	}
	return int8(__obf_bbfd2ac8618a6f0c)
}

// ReadUint8 read uint8
func (__obf_1bb30e8a74ed8233 *Iterator) ReadUint8() (__obf_e46f5fe3db5036fe uint8) {
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.__obf_f735ee00ef3fae0b(__obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049())
	if __obf_bbfd2ac8618a6f0c > math.MaxUint8 {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadUint8", "overflow: "+strconv.FormatInt(int64(__obf_bbfd2ac8618a6f0c), 10))
		return
	}
	return uint8(__obf_bbfd2ac8618a6f0c)
}

// ReadInt16 read int16
func (__obf_1bb30e8a74ed8233 *Iterator) ReadInt16() (__obf_e46f5fe3db5036fe int16) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == '-' {
		__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.__obf_f735ee00ef3fae0b(__obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc())
		if __obf_bbfd2ac8618a6f0c > math.MaxInt16+1 {
			__obf_1bb30e8a74ed8233.
				ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_bbfd2ac8618a6f0c), 10))
			return
		}
		return -int16(__obf_bbfd2ac8618a6f0c)
	}
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.__obf_f735ee00ef3fae0b(__obf_f16b4157911bc9af)
	if __obf_bbfd2ac8618a6f0c > math.MaxInt16 {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_bbfd2ac8618a6f0c), 10))
		return
	}
	return int16(__obf_bbfd2ac8618a6f0c)
}

// ReadUint16 read uint16
func (__obf_1bb30e8a74ed8233 *Iterator) ReadUint16() (__obf_e46f5fe3db5036fe uint16) {
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.__obf_f735ee00ef3fae0b(__obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049())
	if __obf_bbfd2ac8618a6f0c > math.MaxUint16 {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadUint16", "overflow: "+strconv.FormatInt(int64(__obf_bbfd2ac8618a6f0c), 10))
		return
	}
	return uint16(__obf_bbfd2ac8618a6f0c)
}

// ReadInt32 read int32
func (__obf_1bb30e8a74ed8233 *Iterator) ReadInt32() (__obf_e46f5fe3db5036fe int32) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == '-' {
		__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.__obf_f735ee00ef3fae0b(__obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc())
		if __obf_bbfd2ac8618a6f0c > math.MaxInt32+1 {
			__obf_1bb30e8a74ed8233.
				ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_bbfd2ac8618a6f0c), 10))
			return
		}
		return -int32(__obf_bbfd2ac8618a6f0c)
	}
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.__obf_f735ee00ef3fae0b(__obf_f16b4157911bc9af)
	if __obf_bbfd2ac8618a6f0c > math.MaxInt32 {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_bbfd2ac8618a6f0c), 10))
		return
	}
	return int32(__obf_bbfd2ac8618a6f0c)
}

// ReadUint32 read uint32
func (__obf_1bb30e8a74ed8233 *Iterator) ReadUint32() (__obf_e46f5fe3db5036fe uint32) {
	return __obf_1bb30e8a74ed8233.__obf_f735ee00ef3fae0b(__obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049())
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_f735ee00ef3fae0b(__obf_f16b4157911bc9af byte) (__obf_e46f5fe3db5036fe uint32) {
	__obf_2f4419eff5cde989 := __obf_56c7464cf1f5a110[__obf_f16b4157911bc9af]
	if __obf_2f4419eff5cde989 == 0 {
		__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
		return 0 // single zero
	}
	if __obf_2f4419eff5cde989 == __obf_17acfd2d0998ee00 {
		__obf_1bb30e8a74ed8233.
			ReportError("readUint32", "unexpected character: "+string([]byte{byte(__obf_2f4419eff5cde989)}))
		return
	}
	__obf_4724d1b596d6a0c0 := uint32(__obf_2f4419eff5cde989)
	if __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae-__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 > 10 {
		__obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21
		__obf_876313056e7be50f := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_876313056e7be50f == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0
		}
		__obf_5aa5c8829b97f182++
		__obf_05a3bdb6b937ea2f := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_05a3bdb6b937ea2f == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*10 + uint32(__obf_876313056e7be50f)
		}
		__obf_5aa5c8829b97f182++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_ecb6354aaf3b0c49 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_ecb6354aaf3b0c49 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*100 + uint32(__obf_876313056e7be50f)*10 + uint32(__obf_05a3bdb6b937ea2f)
		}
		__obf_5aa5c8829b97f182++
		__obf_229d756183516940 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_229d756183516940 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*1000 + uint32(__obf_876313056e7be50f)*100 + uint32(__obf_05a3bdb6b937ea2f)*10 + uint32(__obf_ecb6354aaf3b0c49)
		}
		__obf_5aa5c8829b97f182++
		__obf_a7f3336a2b12c963 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_a7f3336a2b12c963 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*10000 + uint32(__obf_876313056e7be50f)*1000 + uint32(__obf_05a3bdb6b937ea2f)*100 + uint32(__obf_ecb6354aaf3b0c49)*10 + uint32(__obf_229d756183516940)
		}
		__obf_5aa5c8829b97f182++
		__obf_048af3fc3830ebf6 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_048af3fc3830ebf6 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*100000 + uint32(__obf_876313056e7be50f)*10000 + uint32(__obf_05a3bdb6b937ea2f)*1000 + uint32(__obf_ecb6354aaf3b0c49)*100 + uint32(__obf_229d756183516940)*10 + uint32(__obf_a7f3336a2b12c963)
		}
		__obf_5aa5c8829b97f182++
		__obf_32c364a9a036e296 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_32c364a9a036e296 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*1000000 + uint32(__obf_876313056e7be50f)*100000 + uint32(__obf_05a3bdb6b937ea2f)*10000 + uint32(__obf_ecb6354aaf3b0c49)*1000 + uint32(__obf_229d756183516940)*100 + uint32(__obf_a7f3336a2b12c963)*10 + uint32(__obf_048af3fc3830ebf6)
		}
		__obf_5aa5c8829b97f182++
		__obf_f396c5cb97dac746 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		__obf_4724d1b596d6a0c0 = __obf_4724d1b596d6a0c0*10000000 + uint32(__obf_876313056e7be50f)*1000000 + uint32(__obf_05a3bdb6b937ea2f)*100000 + uint32(__obf_ecb6354aaf3b0c49)*10000 + uint32(__obf_229d756183516940)*1000 + uint32(__obf_a7f3336a2b12c963)*100 + uint32(__obf_048af3fc3830ebf6)*10 + uint32(__obf_32c364a9a036e296)
		__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
		if __obf_f396c5cb97dac746 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0
		}
	}
	for {
		for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
			__obf_2f4419eff5cde989 = __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
			if __obf_2f4419eff5cde989 == __obf_17acfd2d0998ee00 {
				__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
				__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
				return __obf_4724d1b596d6a0c0
			}
			if __obf_4724d1b596d6a0c0 > __obf_631720670e71a3cf {
				__obf_9b1fe1f4d908a86d := (__obf_4724d1b596d6a0c0 << 3) + (__obf_4724d1b596d6a0c0 << 1) + uint32(__obf_2f4419eff5cde989)
				if __obf_9b1fe1f4d908a86d < __obf_4724d1b596d6a0c0 {
					__obf_1bb30e8a74ed8233.
						ReportError("readUint32", "overflow")
					return
				}
				__obf_4724d1b596d6a0c0 = __obf_9b1fe1f4d908a86d
				continue
			}
			__obf_4724d1b596d6a0c0 = (__obf_4724d1b596d6a0c0 << 3) + (__obf_4724d1b596d6a0c0 << 1) + uint32(__obf_2f4419eff5cde989)
		}
		if !__obf_1bb30e8a74ed8233.__obf_e927802c539fc09d() {
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0
		}
	}
}

// ReadInt64 read int64
func (__obf_1bb30e8a74ed8233 *Iterator) ReadInt64() (__obf_e46f5fe3db5036fe int64) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == '-' {
		__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.__obf_8fe731619a0b623b(__obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc())
		if __obf_bbfd2ac8618a6f0c > math.MaxInt64+1 {
			__obf_1bb30e8a74ed8233.
				ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_bbfd2ac8618a6f0c), 10))
			return
		}
		return -int64(__obf_bbfd2ac8618a6f0c)
	}
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.__obf_8fe731619a0b623b(__obf_f16b4157911bc9af)
	if __obf_bbfd2ac8618a6f0c > math.MaxInt64 {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_bbfd2ac8618a6f0c), 10))
		return
	}
	return int64(__obf_bbfd2ac8618a6f0c)
}

// ReadUint64 read uint64
func (__obf_1bb30e8a74ed8233 *Iterator) ReadUint64() uint64 {
	return __obf_1bb30e8a74ed8233.__obf_8fe731619a0b623b(__obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049())
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_8fe731619a0b623b(__obf_f16b4157911bc9af byte) (__obf_e46f5fe3db5036fe uint64) {
	__obf_2f4419eff5cde989 := __obf_56c7464cf1f5a110[__obf_f16b4157911bc9af]
	if __obf_2f4419eff5cde989 == 0 {
		__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
		return 0 // single zero
	}
	if __obf_2f4419eff5cde989 == __obf_17acfd2d0998ee00 {
		__obf_1bb30e8a74ed8233.
			ReportError("readUint64", "unexpected character: "+string([]byte{byte(__obf_2f4419eff5cde989)}))
		return
	}
	__obf_4724d1b596d6a0c0 := uint64(__obf_2f4419eff5cde989)
	if __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae-__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 > 10 {
		__obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21
		__obf_876313056e7be50f := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_876313056e7be50f == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0
		}
		__obf_5aa5c8829b97f182++
		__obf_05a3bdb6b937ea2f := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_05a3bdb6b937ea2f == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*10 + uint64(__obf_876313056e7be50f)
		}
		__obf_5aa5c8829b97f182++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_ecb6354aaf3b0c49 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_ecb6354aaf3b0c49 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*100 + uint64(__obf_876313056e7be50f)*10 + uint64(__obf_05a3bdb6b937ea2f)
		}
		__obf_5aa5c8829b97f182++
		__obf_229d756183516940 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_229d756183516940 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*1000 + uint64(__obf_876313056e7be50f)*100 + uint64(__obf_05a3bdb6b937ea2f)*10 + uint64(__obf_ecb6354aaf3b0c49)
		}
		__obf_5aa5c8829b97f182++
		__obf_a7f3336a2b12c963 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_a7f3336a2b12c963 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*10000 + uint64(__obf_876313056e7be50f)*1000 + uint64(__obf_05a3bdb6b937ea2f)*100 + uint64(__obf_ecb6354aaf3b0c49)*10 + uint64(__obf_229d756183516940)
		}
		__obf_5aa5c8829b97f182++
		__obf_048af3fc3830ebf6 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_048af3fc3830ebf6 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*100000 + uint64(__obf_876313056e7be50f)*10000 + uint64(__obf_05a3bdb6b937ea2f)*1000 + uint64(__obf_ecb6354aaf3b0c49)*100 + uint64(__obf_229d756183516940)*10 + uint64(__obf_a7f3336a2b12c963)
		}
		__obf_5aa5c8829b97f182++
		__obf_32c364a9a036e296 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		if __obf_32c364a9a036e296 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0*1000000 + uint64(__obf_876313056e7be50f)*100000 + uint64(__obf_05a3bdb6b937ea2f)*10000 + uint64(__obf_ecb6354aaf3b0c49)*1000 + uint64(__obf_229d756183516940)*100 + uint64(__obf_a7f3336a2b12c963)*10 + uint64(__obf_048af3fc3830ebf6)
		}
		__obf_5aa5c8829b97f182++
		__obf_f396c5cb97dac746 := __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
		__obf_4724d1b596d6a0c0 = __obf_4724d1b596d6a0c0*10000000 + uint64(__obf_876313056e7be50f)*1000000 + uint64(__obf_05a3bdb6b937ea2f)*100000 + uint64(__obf_ecb6354aaf3b0c49)*10000 + uint64(__obf_229d756183516940)*1000 + uint64(__obf_a7f3336a2b12c963)*100 + uint64(__obf_048af3fc3830ebf6)*10 + uint64(__obf_32c364a9a036e296)
		__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
		if __obf_f396c5cb97dac746 == __obf_17acfd2d0998ee00 {
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0
		}
	}
	for {
		for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
			__obf_2f4419eff5cde989 = __obf_56c7464cf1f5a110[__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]]
			if __obf_2f4419eff5cde989 == __obf_17acfd2d0998ee00 {
				__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
				__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
				return __obf_4724d1b596d6a0c0
			}
			if __obf_4724d1b596d6a0c0 > __obf_fea4c6f6b5068026 {
				__obf_9b1fe1f4d908a86d := (__obf_4724d1b596d6a0c0 << 3) + (__obf_4724d1b596d6a0c0 << 1) + uint64(__obf_2f4419eff5cde989)
				if __obf_9b1fe1f4d908a86d < __obf_4724d1b596d6a0c0 {
					__obf_1bb30e8a74ed8233.
						ReportError("readUint64", "overflow")
					return
				}
				__obf_4724d1b596d6a0c0 = __obf_9b1fe1f4d908a86d
				continue
			}
			__obf_4724d1b596d6a0c0 = (__obf_4724d1b596d6a0c0 << 3) + (__obf_4724d1b596d6a0c0 << 1) + uint64(__obf_2f4419eff5cde989)
		}
		if !__obf_1bb30e8a74ed8233.__obf_e927802c539fc09d() {
			__obf_1bb30e8a74ed8233.__obf_403107a39c491f28()
			return __obf_4724d1b596d6a0c0
		}
	}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_403107a39c491f28() {
	if __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae && __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21] == '.' {
		__obf_1bb30e8a74ed8233.
			ReportError("assertInteger", "can not decode float as int")
	}
}
