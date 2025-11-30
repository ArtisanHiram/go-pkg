package __obf_c3cd04a15c56f32f

import (
	"math"
	"strconv"
)

var __obf_2bbeb874031ba0db []int8

const __obf_e9c077bf3a5fb89f = uint32(0xffffffff)/10 - 1
const __obf_61c0bbfc8fa559d6 = uint64(0xffffffffffffffff)/10 - 1
const __obf_8d3d71fc19d51467 = 1<<53 - 1

func init() {
	__obf_2bbeb874031ba0db = make([]int8, 256)
	for __obf_28d099df85f083a8 := 0; __obf_28d099df85f083a8 < len(__obf_2bbeb874031ba0db); __obf_28d099df85f083a8++ {
		__obf_2bbeb874031ba0db[__obf_28d099df85f083a8] = __obf_234bd47da6a58289
	}
	for __obf_28d099df85f083a8 := int8('0'); __obf_28d099df85f083a8 <= int8('9'); __obf_28d099df85f083a8++ {
		__obf_2bbeb874031ba0db[__obf_28d099df85f083a8] = __obf_28d099df85f083a8 - int8('0')
	}
}

// ReadUint read uint
func (__obf_edd9fe48d39445e4 *Iterator) ReadUint() uint {
	if strconv.IntSize == 32 {
		return uint(__obf_edd9fe48d39445e4.ReadUint32())
	}
	return uint(__obf_edd9fe48d39445e4.ReadUint64())
}

// ReadInt read int
func (__obf_edd9fe48d39445e4 *Iterator) ReadInt() int {
	if strconv.IntSize == 32 {
		return int(__obf_edd9fe48d39445e4.ReadInt32())
	}
	return int(__obf_edd9fe48d39445e4.ReadInt64())
}

// ReadInt8 read int8
func (__obf_edd9fe48d39445e4 *Iterator) ReadInt8() (__obf_316af79472661247 int8) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == '-' {
		__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.__obf_ab9d31c9c0a0a86d(__obf_edd9fe48d39445e4.__obf_70c673c91de4f883())
		if __obf_d59813f23ad740a8 > math.MaxInt8+1 {
			__obf_edd9fe48d39445e4.
				ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_d59813f23ad740a8), 10))
			return
		}
		return -int8(__obf_d59813f23ad740a8)
	}
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.__obf_ab9d31c9c0a0a86d(__obf_0c1bc1e511a43120)
	if __obf_d59813f23ad740a8 > math.MaxInt8 {
		__obf_edd9fe48d39445e4.
			ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_d59813f23ad740a8), 10))
		return
	}
	return int8(__obf_d59813f23ad740a8)
}

// ReadUint8 read uint8
func (__obf_edd9fe48d39445e4 *Iterator) ReadUint8() (__obf_316af79472661247 uint8) {
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.__obf_ab9d31c9c0a0a86d(__obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11())
	if __obf_d59813f23ad740a8 > math.MaxUint8 {
		__obf_edd9fe48d39445e4.
			ReportError("ReadUint8", "overflow: "+strconv.FormatInt(int64(__obf_d59813f23ad740a8), 10))
		return
	}
	return uint8(__obf_d59813f23ad740a8)
}

// ReadInt16 read int16
func (__obf_edd9fe48d39445e4 *Iterator) ReadInt16() (__obf_316af79472661247 int16) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == '-' {
		__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.__obf_ab9d31c9c0a0a86d(__obf_edd9fe48d39445e4.__obf_70c673c91de4f883())
		if __obf_d59813f23ad740a8 > math.MaxInt16+1 {
			__obf_edd9fe48d39445e4.
				ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_d59813f23ad740a8), 10))
			return
		}
		return -int16(__obf_d59813f23ad740a8)
	}
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.__obf_ab9d31c9c0a0a86d(__obf_0c1bc1e511a43120)
	if __obf_d59813f23ad740a8 > math.MaxInt16 {
		__obf_edd9fe48d39445e4.
			ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_d59813f23ad740a8), 10))
		return
	}
	return int16(__obf_d59813f23ad740a8)
}

// ReadUint16 read uint16
func (__obf_edd9fe48d39445e4 *Iterator) ReadUint16() (__obf_316af79472661247 uint16) {
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.__obf_ab9d31c9c0a0a86d(__obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11())
	if __obf_d59813f23ad740a8 > math.MaxUint16 {
		__obf_edd9fe48d39445e4.
			ReportError("ReadUint16", "overflow: "+strconv.FormatInt(int64(__obf_d59813f23ad740a8), 10))
		return
	}
	return uint16(__obf_d59813f23ad740a8)
}

// ReadInt32 read int32
func (__obf_edd9fe48d39445e4 *Iterator) ReadInt32() (__obf_316af79472661247 int32) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == '-' {
		__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.__obf_ab9d31c9c0a0a86d(__obf_edd9fe48d39445e4.__obf_70c673c91de4f883())
		if __obf_d59813f23ad740a8 > math.MaxInt32+1 {
			__obf_edd9fe48d39445e4.
				ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_d59813f23ad740a8), 10))
			return
		}
		return -int32(__obf_d59813f23ad740a8)
	}
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.__obf_ab9d31c9c0a0a86d(__obf_0c1bc1e511a43120)
	if __obf_d59813f23ad740a8 > math.MaxInt32 {
		__obf_edd9fe48d39445e4.
			ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_d59813f23ad740a8), 10))
		return
	}
	return int32(__obf_d59813f23ad740a8)
}

// ReadUint32 read uint32
func (__obf_edd9fe48d39445e4 *Iterator) ReadUint32() (__obf_316af79472661247 uint32) {
	return __obf_edd9fe48d39445e4.__obf_ab9d31c9c0a0a86d(__obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11())
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_ab9d31c9c0a0a86d(__obf_0c1bc1e511a43120 byte) (__obf_316af79472661247 uint32) {
	__obf_5e075812c416735d := __obf_2bbeb874031ba0db[__obf_0c1bc1e511a43120]
	if __obf_5e075812c416735d == 0 {
		__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
		return 0 // single zero
	}
	if __obf_5e075812c416735d == __obf_234bd47da6a58289 {
		__obf_edd9fe48d39445e4.
			ReportError("readUint32", "unexpected character: "+string([]byte{byte(__obf_5e075812c416735d)}))
		return
	}
	__obf_1a534fe904baaf4a := uint32(__obf_5e075812c416735d)
	if __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57-__obf_edd9fe48d39445e4.__obf_edd3c3885447229b > 10 {
		__obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b
		__obf_b4ad9dc1a0ab5358 := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_b4ad9dc1a0ab5358 == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a
		}
		__obf_28d099df85f083a8++
		__obf_0a922a725c232bef := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_0a922a725c232bef == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*10 + uint32(__obf_b4ad9dc1a0ab5358)
		}
		__obf_28d099df85f083a8++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_6fcb7f366f6b511c := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_6fcb7f366f6b511c == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*100 + uint32(__obf_b4ad9dc1a0ab5358)*10 + uint32(__obf_0a922a725c232bef)
		}
		__obf_28d099df85f083a8++
		__obf_5035bccc47dbd7dd := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_5035bccc47dbd7dd == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*1000 + uint32(__obf_b4ad9dc1a0ab5358)*100 + uint32(__obf_0a922a725c232bef)*10 + uint32(__obf_6fcb7f366f6b511c)
		}
		__obf_28d099df85f083a8++
		__obf_6eb5666211b89940 := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_6eb5666211b89940 == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*10000 + uint32(__obf_b4ad9dc1a0ab5358)*1000 + uint32(__obf_0a922a725c232bef)*100 + uint32(__obf_6fcb7f366f6b511c)*10 + uint32(__obf_5035bccc47dbd7dd)
		}
		__obf_28d099df85f083a8++
		__obf_be11d974c2752ddf := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_be11d974c2752ddf == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*100000 + uint32(__obf_b4ad9dc1a0ab5358)*10000 + uint32(__obf_0a922a725c232bef)*1000 + uint32(__obf_6fcb7f366f6b511c)*100 + uint32(__obf_5035bccc47dbd7dd)*10 + uint32(__obf_6eb5666211b89940)
		}
		__obf_28d099df85f083a8++
		__obf_bbcd0db27325b612 := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_bbcd0db27325b612 == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*1000000 + uint32(__obf_b4ad9dc1a0ab5358)*100000 + uint32(__obf_0a922a725c232bef)*10000 + uint32(__obf_6fcb7f366f6b511c)*1000 + uint32(__obf_5035bccc47dbd7dd)*100 + uint32(__obf_6eb5666211b89940)*10 + uint32(__obf_be11d974c2752ddf)
		}
		__obf_28d099df85f083a8++
		__obf_039f0de76bae99e3 := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		__obf_1a534fe904baaf4a = __obf_1a534fe904baaf4a*10000000 + uint32(__obf_b4ad9dc1a0ab5358)*1000000 + uint32(__obf_0a922a725c232bef)*100000 + uint32(__obf_6fcb7f366f6b511c)*10000 + uint32(__obf_5035bccc47dbd7dd)*1000 + uint32(__obf_6eb5666211b89940)*100 + uint32(__obf_be11d974c2752ddf)*10 + uint32(__obf_bbcd0db27325b612)
		__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
		if __obf_039f0de76bae99e3 == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a
		}
	}
	for {
		for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
			__obf_5e075812c416735d = __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
			if __obf_5e075812c416735d == __obf_234bd47da6a58289 {
				__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
				__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
				return __obf_1a534fe904baaf4a
			}
			if __obf_1a534fe904baaf4a > __obf_e9c077bf3a5fb89f {
				__obf_06f1b354459bf29f := (__obf_1a534fe904baaf4a << 3) + (__obf_1a534fe904baaf4a << 1) + uint32(__obf_5e075812c416735d)
				if __obf_06f1b354459bf29f < __obf_1a534fe904baaf4a {
					__obf_edd9fe48d39445e4.
						ReportError("readUint32", "overflow")
					return
				}
				__obf_1a534fe904baaf4a = __obf_06f1b354459bf29f
				continue
			}
			__obf_1a534fe904baaf4a = (__obf_1a534fe904baaf4a << 3) + (__obf_1a534fe904baaf4a << 1) + uint32(__obf_5e075812c416735d)
		}
		if !__obf_edd9fe48d39445e4.__obf_35e704c169a88a74() {
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a
		}
	}
}

// ReadInt64 read int64
func (__obf_edd9fe48d39445e4 *Iterator) ReadInt64() (__obf_316af79472661247 int64) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == '-' {
		__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.__obf_6ba9369cd0a94b2b(__obf_edd9fe48d39445e4.__obf_70c673c91de4f883())
		if __obf_d59813f23ad740a8 > math.MaxInt64+1 {
			__obf_edd9fe48d39445e4.
				ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_d59813f23ad740a8), 10))
			return
		}
		return -int64(__obf_d59813f23ad740a8)
	}
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.__obf_6ba9369cd0a94b2b(__obf_0c1bc1e511a43120)
	if __obf_d59813f23ad740a8 > math.MaxInt64 {
		__obf_edd9fe48d39445e4.
			ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_d59813f23ad740a8), 10))
		return
	}
	return int64(__obf_d59813f23ad740a8)
}

// ReadUint64 read uint64
func (__obf_edd9fe48d39445e4 *Iterator) ReadUint64() uint64 {
	return __obf_edd9fe48d39445e4.__obf_6ba9369cd0a94b2b(__obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11())
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_6ba9369cd0a94b2b(__obf_0c1bc1e511a43120 byte) (__obf_316af79472661247 uint64) {
	__obf_5e075812c416735d := __obf_2bbeb874031ba0db[__obf_0c1bc1e511a43120]
	if __obf_5e075812c416735d == 0 {
		__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
		return 0 // single zero
	}
	if __obf_5e075812c416735d == __obf_234bd47da6a58289 {
		__obf_edd9fe48d39445e4.
			ReportError("readUint64", "unexpected character: "+string([]byte{byte(__obf_5e075812c416735d)}))
		return
	}
	__obf_1a534fe904baaf4a := uint64(__obf_5e075812c416735d)
	if __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57-__obf_edd9fe48d39445e4.__obf_edd3c3885447229b > 10 {
		__obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b
		__obf_b4ad9dc1a0ab5358 := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_b4ad9dc1a0ab5358 == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a
		}
		__obf_28d099df85f083a8++
		__obf_0a922a725c232bef := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_0a922a725c232bef == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*10 + uint64(__obf_b4ad9dc1a0ab5358)
		}
		__obf_28d099df85f083a8++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_6fcb7f366f6b511c := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_6fcb7f366f6b511c == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*100 + uint64(__obf_b4ad9dc1a0ab5358)*10 + uint64(__obf_0a922a725c232bef)
		}
		__obf_28d099df85f083a8++
		__obf_5035bccc47dbd7dd := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_5035bccc47dbd7dd == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*1000 + uint64(__obf_b4ad9dc1a0ab5358)*100 + uint64(__obf_0a922a725c232bef)*10 + uint64(__obf_6fcb7f366f6b511c)
		}
		__obf_28d099df85f083a8++
		__obf_6eb5666211b89940 := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_6eb5666211b89940 == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*10000 + uint64(__obf_b4ad9dc1a0ab5358)*1000 + uint64(__obf_0a922a725c232bef)*100 + uint64(__obf_6fcb7f366f6b511c)*10 + uint64(__obf_5035bccc47dbd7dd)
		}
		__obf_28d099df85f083a8++
		__obf_be11d974c2752ddf := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_be11d974c2752ddf == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*100000 + uint64(__obf_b4ad9dc1a0ab5358)*10000 + uint64(__obf_0a922a725c232bef)*1000 + uint64(__obf_6fcb7f366f6b511c)*100 + uint64(__obf_5035bccc47dbd7dd)*10 + uint64(__obf_6eb5666211b89940)
		}
		__obf_28d099df85f083a8++
		__obf_bbcd0db27325b612 := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		if __obf_bbcd0db27325b612 == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a*1000000 + uint64(__obf_b4ad9dc1a0ab5358)*100000 + uint64(__obf_0a922a725c232bef)*10000 + uint64(__obf_6fcb7f366f6b511c)*1000 + uint64(__obf_5035bccc47dbd7dd)*100 + uint64(__obf_6eb5666211b89940)*10 + uint64(__obf_be11d974c2752ddf)
		}
		__obf_28d099df85f083a8++
		__obf_039f0de76bae99e3 := __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
		__obf_1a534fe904baaf4a = __obf_1a534fe904baaf4a*10000000 + uint64(__obf_b4ad9dc1a0ab5358)*1000000 + uint64(__obf_0a922a725c232bef)*100000 + uint64(__obf_6fcb7f366f6b511c)*10000 + uint64(__obf_5035bccc47dbd7dd)*1000 + uint64(__obf_6eb5666211b89940)*100 + uint64(__obf_be11d974c2752ddf)*10 + uint64(__obf_bbcd0db27325b612)
		__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
		if __obf_039f0de76bae99e3 == __obf_234bd47da6a58289 {
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a
		}
	}
	for {
		for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
			__obf_5e075812c416735d = __obf_2bbeb874031ba0db[__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]]
			if __obf_5e075812c416735d == __obf_234bd47da6a58289 {
				__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
				__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
				return __obf_1a534fe904baaf4a
			}
			if __obf_1a534fe904baaf4a > __obf_61c0bbfc8fa559d6 {
				__obf_06f1b354459bf29f := (__obf_1a534fe904baaf4a << 3) + (__obf_1a534fe904baaf4a << 1) + uint64(__obf_5e075812c416735d)
				if __obf_06f1b354459bf29f < __obf_1a534fe904baaf4a {
					__obf_edd9fe48d39445e4.
						ReportError("readUint64", "overflow")
					return
				}
				__obf_1a534fe904baaf4a = __obf_06f1b354459bf29f
				continue
			}
			__obf_1a534fe904baaf4a = (__obf_1a534fe904baaf4a << 3) + (__obf_1a534fe904baaf4a << 1) + uint64(__obf_5e075812c416735d)
		}
		if !__obf_edd9fe48d39445e4.__obf_35e704c169a88a74() {
			__obf_edd9fe48d39445e4.__obf_9776b470236f6734()
			return __obf_1a534fe904baaf4a
		}
	}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_9776b470236f6734() {
	if __obf_edd9fe48d39445e4.__obf_edd3c3885447229b < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 && __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_edd9fe48d39445e4.__obf_edd3c3885447229b] == '.' {
		__obf_edd9fe48d39445e4.
			ReportError("assertInteger", "can not decode float as int")
	}
}
