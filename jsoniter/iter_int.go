package __obf_c7b79b12b33d8238

import (
	"math"
	"strconv"
)

var __obf_93d9f69558bcb6dd []int8

const __obf_4dd169b585e67502 = uint32(0xffffffff)/10 - 1
const __obf_888f28ae8b5c39dc = uint64(0xffffffffffffffff)/10 - 1
const __obf_a86df43efd02f0dc = 1<<53 - 1

func init() {
	__obf_93d9f69558bcb6dd = make([]int8, 256)
	for __obf_a051269af3a541bb := 0; __obf_a051269af3a541bb < len(__obf_93d9f69558bcb6dd); __obf_a051269af3a541bb++ {
		__obf_93d9f69558bcb6dd[__obf_a051269af3a541bb] = __obf_905dd362f7428491
	}
	for __obf_a051269af3a541bb := int8('0'); __obf_a051269af3a541bb <= int8('9'); __obf_a051269af3a541bb++ {
		__obf_93d9f69558bcb6dd[__obf_a051269af3a541bb] = __obf_a051269af3a541bb - int8('0')
	}
}

// ReadUint read uint
func (__obf_0da8c843dad13139 *Iterator) ReadUint() uint {
	if strconv.IntSize == 32 {
		return uint(__obf_0da8c843dad13139.ReadUint32())
	}
	return uint(__obf_0da8c843dad13139.ReadUint64())
}

// ReadInt read int
func (__obf_0da8c843dad13139 *Iterator) ReadInt() int {
	if strconv.IntSize == 32 {
		return int(__obf_0da8c843dad13139.ReadInt32())
	}
	return int(__obf_0da8c843dad13139.ReadInt64())
}

// ReadInt8 read int8
func (__obf_0da8c843dad13139 *Iterator) ReadInt8() (__obf_9a8638dac9e1c083 int8) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == '-' {
		__obf_35accd096612ebe4 := __obf_0da8c843dad13139.__obf_6271157db3787fb7(__obf_0da8c843dad13139.__obf_fe749dd9dadf32f8())
		if __obf_35accd096612ebe4 > math.MaxInt8+1 {
			__obf_0da8c843dad13139.
				ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_35accd096612ebe4), 10))
			return
		}
		return -int8(__obf_35accd096612ebe4)
	}
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.__obf_6271157db3787fb7(__obf_12577c03a12f0d2c)
	if __obf_35accd096612ebe4 > math.MaxInt8 {
		__obf_0da8c843dad13139.
			ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_35accd096612ebe4), 10))
		return
	}
	return int8(__obf_35accd096612ebe4)
}

// ReadUint8 read uint8
func (__obf_0da8c843dad13139 *Iterator) ReadUint8() (__obf_9a8638dac9e1c083 uint8) {
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.__obf_6271157db3787fb7(__obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d())
	if __obf_35accd096612ebe4 > math.MaxUint8 {
		__obf_0da8c843dad13139.
			ReportError("ReadUint8", "overflow: "+strconv.FormatInt(int64(__obf_35accd096612ebe4), 10))
		return
	}
	return uint8(__obf_35accd096612ebe4)
}

// ReadInt16 read int16
func (__obf_0da8c843dad13139 *Iterator) ReadInt16() (__obf_9a8638dac9e1c083 int16) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == '-' {
		__obf_35accd096612ebe4 := __obf_0da8c843dad13139.__obf_6271157db3787fb7(__obf_0da8c843dad13139.__obf_fe749dd9dadf32f8())
		if __obf_35accd096612ebe4 > math.MaxInt16+1 {
			__obf_0da8c843dad13139.
				ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_35accd096612ebe4), 10))
			return
		}
		return -int16(__obf_35accd096612ebe4)
	}
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.__obf_6271157db3787fb7(__obf_12577c03a12f0d2c)
	if __obf_35accd096612ebe4 > math.MaxInt16 {
		__obf_0da8c843dad13139.
			ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_35accd096612ebe4), 10))
		return
	}
	return int16(__obf_35accd096612ebe4)
}

// ReadUint16 read uint16
func (__obf_0da8c843dad13139 *Iterator) ReadUint16() (__obf_9a8638dac9e1c083 uint16) {
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.__obf_6271157db3787fb7(__obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d())
	if __obf_35accd096612ebe4 > math.MaxUint16 {
		__obf_0da8c843dad13139.
			ReportError("ReadUint16", "overflow: "+strconv.FormatInt(int64(__obf_35accd096612ebe4), 10))
		return
	}
	return uint16(__obf_35accd096612ebe4)
}

// ReadInt32 read int32
func (__obf_0da8c843dad13139 *Iterator) ReadInt32() (__obf_9a8638dac9e1c083 int32) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == '-' {
		__obf_35accd096612ebe4 := __obf_0da8c843dad13139.__obf_6271157db3787fb7(__obf_0da8c843dad13139.__obf_fe749dd9dadf32f8())
		if __obf_35accd096612ebe4 > math.MaxInt32+1 {
			__obf_0da8c843dad13139.
				ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_35accd096612ebe4), 10))
			return
		}
		return -int32(__obf_35accd096612ebe4)
	}
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.__obf_6271157db3787fb7(__obf_12577c03a12f0d2c)
	if __obf_35accd096612ebe4 > math.MaxInt32 {
		__obf_0da8c843dad13139.
			ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_35accd096612ebe4), 10))
		return
	}
	return int32(__obf_35accd096612ebe4)
}

// ReadUint32 read uint32
func (__obf_0da8c843dad13139 *Iterator) ReadUint32() (__obf_9a8638dac9e1c083 uint32) {
	return __obf_0da8c843dad13139.__obf_6271157db3787fb7(__obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d())
}

func (__obf_0da8c843dad13139 *Iterator) __obf_6271157db3787fb7(__obf_12577c03a12f0d2c byte) (__obf_9a8638dac9e1c083 uint32) {
	__obf_83600c60799be31c := __obf_93d9f69558bcb6dd[__obf_12577c03a12f0d2c]
	if __obf_83600c60799be31c == 0 {
		__obf_0da8c843dad13139.__obf_cedfe64867925b60()
		return 0 // single zero
	}
	if __obf_83600c60799be31c == __obf_905dd362f7428491 {
		__obf_0da8c843dad13139.
			ReportError("readUint32", "unexpected character: "+string([]byte{byte(__obf_83600c60799be31c)}))
		return
	}
	__obf_52321dce0d8db989 := uint32(__obf_83600c60799be31c)
	if __obf_0da8c843dad13139.__obf_840246080559848c-__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 > 10 {
		__obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2
		__obf_f1203786d413d1c1 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_f1203786d413d1c1 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989
		}
		__obf_a051269af3a541bb++
		__obf_0d1c817ab9a323c0 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_0d1c817ab9a323c0 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*10 + uint32(__obf_f1203786d413d1c1)
		}
		__obf_a051269af3a541bb++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_9ad0450152566eb5 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_9ad0450152566eb5 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*100 + uint32(__obf_f1203786d413d1c1)*10 + uint32(__obf_0d1c817ab9a323c0)
		}
		__obf_a051269af3a541bb++
		__obf_f3c6c47b2303d824 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_f3c6c47b2303d824 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*1000 + uint32(__obf_f1203786d413d1c1)*100 + uint32(__obf_0d1c817ab9a323c0)*10 + uint32(__obf_9ad0450152566eb5)
		}
		__obf_a051269af3a541bb++
		__obf_1e81010584ae24b0 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_1e81010584ae24b0 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*10000 + uint32(__obf_f1203786d413d1c1)*1000 + uint32(__obf_0d1c817ab9a323c0)*100 + uint32(__obf_9ad0450152566eb5)*10 + uint32(__obf_f3c6c47b2303d824)
		}
		__obf_a051269af3a541bb++
		__obf_132505af8610b167 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_132505af8610b167 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*100000 + uint32(__obf_f1203786d413d1c1)*10000 + uint32(__obf_0d1c817ab9a323c0)*1000 + uint32(__obf_9ad0450152566eb5)*100 + uint32(__obf_f3c6c47b2303d824)*10 + uint32(__obf_1e81010584ae24b0)
		}
		__obf_a051269af3a541bb++
		__obf_080315d80dd66cd7 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_080315d80dd66cd7 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*1000000 + uint32(__obf_f1203786d413d1c1)*100000 + uint32(__obf_0d1c817ab9a323c0)*10000 + uint32(__obf_9ad0450152566eb5)*1000 + uint32(__obf_f3c6c47b2303d824)*100 + uint32(__obf_1e81010584ae24b0)*10 + uint32(__obf_132505af8610b167)
		}
		__obf_a051269af3a541bb++
		__obf_5c0a12b1837a92a1 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		__obf_52321dce0d8db989 = __obf_52321dce0d8db989*10000000 + uint32(__obf_f1203786d413d1c1)*1000000 + uint32(__obf_0d1c817ab9a323c0)*100000 + uint32(__obf_9ad0450152566eb5)*10000 + uint32(__obf_f3c6c47b2303d824)*1000 + uint32(__obf_1e81010584ae24b0)*100 + uint32(__obf_132505af8610b167)*10 + uint32(__obf_080315d80dd66cd7)
		__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
		if __obf_5c0a12b1837a92a1 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989
		}
	}
	for {
		for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
			__obf_83600c60799be31c = __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
			if __obf_83600c60799be31c == __obf_905dd362f7428491 {
				__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
				__obf_0da8c843dad13139.__obf_cedfe64867925b60()
				return __obf_52321dce0d8db989
			}
			if __obf_52321dce0d8db989 > __obf_4dd169b585e67502 {
				__obf_82148c9b913ae6cd := (__obf_52321dce0d8db989 << 3) + (__obf_52321dce0d8db989 << 1) + uint32(__obf_83600c60799be31c)
				if __obf_82148c9b913ae6cd < __obf_52321dce0d8db989 {
					__obf_0da8c843dad13139.
						ReportError("readUint32", "overflow")
					return
				}
				__obf_52321dce0d8db989 = __obf_82148c9b913ae6cd
				continue
			}
			__obf_52321dce0d8db989 = (__obf_52321dce0d8db989 << 3) + (__obf_52321dce0d8db989 << 1) + uint32(__obf_83600c60799be31c)
		}
		if !__obf_0da8c843dad13139.__obf_baaf768e13842431() {
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989
		}
	}
}

// ReadInt64 read int64
func (__obf_0da8c843dad13139 *Iterator) ReadInt64() (__obf_9a8638dac9e1c083 int64) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == '-' {
		__obf_35accd096612ebe4 := __obf_0da8c843dad13139.__obf_a5eb7a835524fb08(__obf_0da8c843dad13139.__obf_fe749dd9dadf32f8())
		if __obf_35accd096612ebe4 > math.MaxInt64+1 {
			__obf_0da8c843dad13139.
				ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_35accd096612ebe4), 10))
			return
		}
		return -int64(__obf_35accd096612ebe4)
	}
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.__obf_a5eb7a835524fb08(__obf_12577c03a12f0d2c)
	if __obf_35accd096612ebe4 > math.MaxInt64 {
		__obf_0da8c843dad13139.
			ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_35accd096612ebe4), 10))
		return
	}
	return int64(__obf_35accd096612ebe4)
}

// ReadUint64 read uint64
func (__obf_0da8c843dad13139 *Iterator) ReadUint64() uint64 {
	return __obf_0da8c843dad13139.__obf_a5eb7a835524fb08(__obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d())
}

func (__obf_0da8c843dad13139 *Iterator) __obf_a5eb7a835524fb08(__obf_12577c03a12f0d2c byte) (__obf_9a8638dac9e1c083 uint64) {
	__obf_83600c60799be31c := __obf_93d9f69558bcb6dd[__obf_12577c03a12f0d2c]
	if __obf_83600c60799be31c == 0 {
		__obf_0da8c843dad13139.__obf_cedfe64867925b60()
		return 0 // single zero
	}
	if __obf_83600c60799be31c == __obf_905dd362f7428491 {
		__obf_0da8c843dad13139.
			ReportError("readUint64", "unexpected character: "+string([]byte{byte(__obf_83600c60799be31c)}))
		return
	}
	__obf_52321dce0d8db989 := uint64(__obf_83600c60799be31c)
	if __obf_0da8c843dad13139.__obf_840246080559848c-__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 > 10 {
		__obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2
		__obf_f1203786d413d1c1 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_f1203786d413d1c1 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989
		}
		__obf_a051269af3a541bb++
		__obf_0d1c817ab9a323c0 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_0d1c817ab9a323c0 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*10 + uint64(__obf_f1203786d413d1c1)
		}
		__obf_a051269af3a541bb++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_9ad0450152566eb5 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_9ad0450152566eb5 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*100 + uint64(__obf_f1203786d413d1c1)*10 + uint64(__obf_0d1c817ab9a323c0)
		}
		__obf_a051269af3a541bb++
		__obf_f3c6c47b2303d824 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_f3c6c47b2303d824 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*1000 + uint64(__obf_f1203786d413d1c1)*100 + uint64(__obf_0d1c817ab9a323c0)*10 + uint64(__obf_9ad0450152566eb5)
		}
		__obf_a051269af3a541bb++
		__obf_1e81010584ae24b0 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_1e81010584ae24b0 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*10000 + uint64(__obf_f1203786d413d1c1)*1000 + uint64(__obf_0d1c817ab9a323c0)*100 + uint64(__obf_9ad0450152566eb5)*10 + uint64(__obf_f3c6c47b2303d824)
		}
		__obf_a051269af3a541bb++
		__obf_132505af8610b167 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_132505af8610b167 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*100000 + uint64(__obf_f1203786d413d1c1)*10000 + uint64(__obf_0d1c817ab9a323c0)*1000 + uint64(__obf_9ad0450152566eb5)*100 + uint64(__obf_f3c6c47b2303d824)*10 + uint64(__obf_1e81010584ae24b0)
		}
		__obf_a051269af3a541bb++
		__obf_080315d80dd66cd7 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		if __obf_080315d80dd66cd7 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989*1000000 + uint64(__obf_f1203786d413d1c1)*100000 + uint64(__obf_0d1c817ab9a323c0)*10000 + uint64(__obf_9ad0450152566eb5)*1000 + uint64(__obf_f3c6c47b2303d824)*100 + uint64(__obf_1e81010584ae24b0)*10 + uint64(__obf_132505af8610b167)
		}
		__obf_a051269af3a541bb++
		__obf_5c0a12b1837a92a1 := __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
		__obf_52321dce0d8db989 = __obf_52321dce0d8db989*10000000 + uint64(__obf_f1203786d413d1c1)*1000000 + uint64(__obf_0d1c817ab9a323c0)*100000 + uint64(__obf_9ad0450152566eb5)*10000 + uint64(__obf_f3c6c47b2303d824)*1000 + uint64(__obf_1e81010584ae24b0)*100 + uint64(__obf_132505af8610b167)*10 + uint64(__obf_080315d80dd66cd7)
		__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
		if __obf_5c0a12b1837a92a1 == __obf_905dd362f7428491 {
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989
		}
	}
	for {
		for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
			__obf_83600c60799be31c = __obf_93d9f69558bcb6dd[__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]]
			if __obf_83600c60799be31c == __obf_905dd362f7428491 {
				__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
				__obf_0da8c843dad13139.__obf_cedfe64867925b60()
				return __obf_52321dce0d8db989
			}
			if __obf_52321dce0d8db989 > __obf_888f28ae8b5c39dc {
				__obf_82148c9b913ae6cd := (__obf_52321dce0d8db989 << 3) + (__obf_52321dce0d8db989 << 1) + uint64(__obf_83600c60799be31c)
				if __obf_82148c9b913ae6cd < __obf_52321dce0d8db989 {
					__obf_0da8c843dad13139.
						ReportError("readUint64", "overflow")
					return
				}
				__obf_52321dce0d8db989 = __obf_82148c9b913ae6cd
				continue
			}
			__obf_52321dce0d8db989 = (__obf_52321dce0d8db989 << 3) + (__obf_52321dce0d8db989 << 1) + uint64(__obf_83600c60799be31c)
		}
		if !__obf_0da8c843dad13139.__obf_baaf768e13842431() {
			__obf_0da8c843dad13139.__obf_cedfe64867925b60()
			return __obf_52321dce0d8db989
		}
	}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_cedfe64867925b60() {
	if __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 < __obf_0da8c843dad13139.__obf_840246080559848c && __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2] == '.' {
		__obf_0da8c843dad13139.
			ReportError("assertInteger", "can not decode float as int")
	}
}
