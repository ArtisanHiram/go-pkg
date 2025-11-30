package __obf_030d39f7a8de96c6

import (
	"math"
	"strconv"
)

var __obf_03403867a04aca29 []int8

const __obf_b87c581cb94dded6 = uint32(0xffffffff)/10 - 1
const __obf_92013c25906febd6 = uint64(0xffffffffffffffff)/10 - 1
const __obf_636d63e04d6021f5 = 1<<53 - 1

func init() {
	__obf_03403867a04aca29 = make([]int8, 256)
	for __obf_82c6e05b9d226c58 := 0; __obf_82c6e05b9d226c58 < len(__obf_03403867a04aca29); __obf_82c6e05b9d226c58++ {
		__obf_03403867a04aca29[__obf_82c6e05b9d226c58] = __obf_78b317d8e70f5d5a
	}
	for __obf_82c6e05b9d226c58 := int8('0'); __obf_82c6e05b9d226c58 <= int8('9'); __obf_82c6e05b9d226c58++ {
		__obf_03403867a04aca29[__obf_82c6e05b9d226c58] = __obf_82c6e05b9d226c58 - int8('0')
	}
}

// ReadUint read uint
func (__obf_4ab56a99965952e7 *Iterator) ReadUint() uint {
	if strconv.IntSize == 32 {
		return uint(__obf_4ab56a99965952e7.ReadUint32())
	}
	return uint(__obf_4ab56a99965952e7.ReadUint64())
}

// ReadInt read int
func (__obf_4ab56a99965952e7 *Iterator) ReadInt() int {
	if strconv.IntSize == 32 {
		return int(__obf_4ab56a99965952e7.ReadInt32())
	}
	return int(__obf_4ab56a99965952e7.ReadInt64())
}

// ReadInt8 read int8
func (__obf_4ab56a99965952e7 *Iterator) ReadInt8() (__obf_59c72400ec1a6110 int8) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == '-' {
		__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.__obf_dc44880eb3cc451d(__obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68())
		if __obf_efaf2719b44ad8ac > math.MaxInt8+1 {
			__obf_4ab56a99965952e7.
				ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_efaf2719b44ad8ac), 10))
			return
		}
		return -int8(__obf_efaf2719b44ad8ac)
	}
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.__obf_dc44880eb3cc451d(__obf_24309b3b0ff9ba22)
	if __obf_efaf2719b44ad8ac > math.MaxInt8 {
		__obf_4ab56a99965952e7.
			ReportError("ReadInt8", "overflow: "+strconv.FormatInt(int64(__obf_efaf2719b44ad8ac), 10))
		return
	}
	return int8(__obf_efaf2719b44ad8ac)
}

// ReadUint8 read uint8
func (__obf_4ab56a99965952e7 *Iterator) ReadUint8() (__obf_59c72400ec1a6110 uint8) {
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.__obf_dc44880eb3cc451d(__obf_4ab56a99965952e7.__obf_61df301d1f67ad73())
	if __obf_efaf2719b44ad8ac > math.MaxUint8 {
		__obf_4ab56a99965952e7.
			ReportError("ReadUint8", "overflow: "+strconv.FormatInt(int64(__obf_efaf2719b44ad8ac), 10))
		return
	}
	return uint8(__obf_efaf2719b44ad8ac)
}

// ReadInt16 read int16
func (__obf_4ab56a99965952e7 *Iterator) ReadInt16() (__obf_59c72400ec1a6110 int16) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == '-' {
		__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.__obf_dc44880eb3cc451d(__obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68())
		if __obf_efaf2719b44ad8ac > math.MaxInt16+1 {
			__obf_4ab56a99965952e7.
				ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_efaf2719b44ad8ac), 10))
			return
		}
		return -int16(__obf_efaf2719b44ad8ac)
	}
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.__obf_dc44880eb3cc451d(__obf_24309b3b0ff9ba22)
	if __obf_efaf2719b44ad8ac > math.MaxInt16 {
		__obf_4ab56a99965952e7.
			ReportError("ReadInt16", "overflow: "+strconv.FormatInt(int64(__obf_efaf2719b44ad8ac), 10))
		return
	}
	return int16(__obf_efaf2719b44ad8ac)
}

// ReadUint16 read uint16
func (__obf_4ab56a99965952e7 *Iterator) ReadUint16() (__obf_59c72400ec1a6110 uint16) {
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.__obf_dc44880eb3cc451d(__obf_4ab56a99965952e7.__obf_61df301d1f67ad73())
	if __obf_efaf2719b44ad8ac > math.MaxUint16 {
		__obf_4ab56a99965952e7.
			ReportError("ReadUint16", "overflow: "+strconv.FormatInt(int64(__obf_efaf2719b44ad8ac), 10))
		return
	}
	return uint16(__obf_efaf2719b44ad8ac)
}

// ReadInt32 read int32
func (__obf_4ab56a99965952e7 *Iterator) ReadInt32() (__obf_59c72400ec1a6110 int32) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == '-' {
		__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.__obf_dc44880eb3cc451d(__obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68())
		if __obf_efaf2719b44ad8ac > math.MaxInt32+1 {
			__obf_4ab56a99965952e7.
				ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_efaf2719b44ad8ac), 10))
			return
		}
		return -int32(__obf_efaf2719b44ad8ac)
	}
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.__obf_dc44880eb3cc451d(__obf_24309b3b0ff9ba22)
	if __obf_efaf2719b44ad8ac > math.MaxInt32 {
		__obf_4ab56a99965952e7.
			ReportError("ReadInt32", "overflow: "+strconv.FormatInt(int64(__obf_efaf2719b44ad8ac), 10))
		return
	}
	return int32(__obf_efaf2719b44ad8ac)
}

// ReadUint32 read uint32
func (__obf_4ab56a99965952e7 *Iterator) ReadUint32() (__obf_59c72400ec1a6110 uint32) {
	return __obf_4ab56a99965952e7.__obf_dc44880eb3cc451d(__obf_4ab56a99965952e7.__obf_61df301d1f67ad73())
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_dc44880eb3cc451d(__obf_24309b3b0ff9ba22 byte) (__obf_59c72400ec1a6110 uint32) {
	__obf_0d7ccf72a2e94911 := __obf_03403867a04aca29[__obf_24309b3b0ff9ba22]
	if __obf_0d7ccf72a2e94911 == 0 {
		__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
		return 0 // single zero
	}
	if __obf_0d7ccf72a2e94911 == __obf_78b317d8e70f5d5a {
		__obf_4ab56a99965952e7.
			ReportError("readUint32", "unexpected character: "+string([]byte{byte(__obf_0d7ccf72a2e94911)}))
		return
	}
	__obf_b80d5217e1232943 := uint32(__obf_0d7ccf72a2e94911)
	if __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc-__obf_4ab56a99965952e7.__obf_5e22d6270d27491f > 10 {
		__obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f
		__obf_8b3d3761235ee860 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_8b3d3761235ee860 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943
		}
		__obf_82c6e05b9d226c58++
		__obf_bd7a44b446c0f2eb := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_bd7a44b446c0f2eb == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*10 + uint32(__obf_8b3d3761235ee860)
		}
		__obf_82c6e05b9d226c58++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_bc5725ebafdfb1b8 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_bc5725ebafdfb1b8 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*100 + uint32(__obf_8b3d3761235ee860)*10 + uint32(__obf_bd7a44b446c0f2eb)
		}
		__obf_82c6e05b9d226c58++
		__obf_5c788bf8b20b2b60 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_5c788bf8b20b2b60 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*1000 + uint32(__obf_8b3d3761235ee860)*100 + uint32(__obf_bd7a44b446c0f2eb)*10 + uint32(__obf_bc5725ebafdfb1b8)
		}
		__obf_82c6e05b9d226c58++
		__obf_79941b1d34ce7165 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_79941b1d34ce7165 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*10000 + uint32(__obf_8b3d3761235ee860)*1000 + uint32(__obf_bd7a44b446c0f2eb)*100 + uint32(__obf_bc5725ebafdfb1b8)*10 + uint32(__obf_5c788bf8b20b2b60)
		}
		__obf_82c6e05b9d226c58++
		__obf_9639592f55bb3c69 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_9639592f55bb3c69 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*100000 + uint32(__obf_8b3d3761235ee860)*10000 + uint32(__obf_bd7a44b446c0f2eb)*1000 + uint32(__obf_bc5725ebafdfb1b8)*100 + uint32(__obf_5c788bf8b20b2b60)*10 + uint32(__obf_79941b1d34ce7165)
		}
		__obf_82c6e05b9d226c58++
		__obf_fab340f78ca73f45 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_fab340f78ca73f45 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*1000000 + uint32(__obf_8b3d3761235ee860)*100000 + uint32(__obf_bd7a44b446c0f2eb)*10000 + uint32(__obf_bc5725ebafdfb1b8)*1000 + uint32(__obf_5c788bf8b20b2b60)*100 + uint32(__obf_79941b1d34ce7165)*10 + uint32(__obf_9639592f55bb3c69)
		}
		__obf_82c6e05b9d226c58++
		__obf_f64f5fdbade2b834 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		__obf_b80d5217e1232943 = __obf_b80d5217e1232943*10000000 + uint32(__obf_8b3d3761235ee860)*1000000 + uint32(__obf_bd7a44b446c0f2eb)*100000 + uint32(__obf_bc5725ebafdfb1b8)*10000 + uint32(__obf_5c788bf8b20b2b60)*1000 + uint32(__obf_79941b1d34ce7165)*100 + uint32(__obf_9639592f55bb3c69)*10 + uint32(__obf_fab340f78ca73f45)
		__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
		if __obf_f64f5fdbade2b834 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943
		}
	}
	for {
		for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
			__obf_0d7ccf72a2e94911 = __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
			if __obf_0d7ccf72a2e94911 == __obf_78b317d8e70f5d5a {
				__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
				__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
				return __obf_b80d5217e1232943
			}
			if __obf_b80d5217e1232943 > __obf_b87c581cb94dded6 {
				__obf_50aaa7aa198ad0ad := (__obf_b80d5217e1232943 << 3) + (__obf_b80d5217e1232943 << 1) + uint32(__obf_0d7ccf72a2e94911)
				if __obf_50aaa7aa198ad0ad < __obf_b80d5217e1232943 {
					__obf_4ab56a99965952e7.
						ReportError("readUint32", "overflow")
					return
				}
				__obf_b80d5217e1232943 = __obf_50aaa7aa198ad0ad
				continue
			}
			__obf_b80d5217e1232943 = (__obf_b80d5217e1232943 << 3) + (__obf_b80d5217e1232943 << 1) + uint32(__obf_0d7ccf72a2e94911)
		}
		if !__obf_4ab56a99965952e7.__obf_08a0b850abb3abbd() {
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943
		}
	}
}

// ReadInt64 read int64
func (__obf_4ab56a99965952e7 *Iterator) ReadInt64() (__obf_59c72400ec1a6110 int64) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == '-' {
		__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.__obf_bc7d04a45a032c40(__obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68())
		if __obf_efaf2719b44ad8ac > math.MaxInt64+1 {
			__obf_4ab56a99965952e7.
				ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_efaf2719b44ad8ac), 10))
			return
		}
		return -int64(__obf_efaf2719b44ad8ac)
	}
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.__obf_bc7d04a45a032c40(__obf_24309b3b0ff9ba22)
	if __obf_efaf2719b44ad8ac > math.MaxInt64 {
		__obf_4ab56a99965952e7.
			ReportError("ReadInt64", "overflow: "+strconv.FormatUint(uint64(__obf_efaf2719b44ad8ac), 10))
		return
	}
	return int64(__obf_efaf2719b44ad8ac)
}

// ReadUint64 read uint64
func (__obf_4ab56a99965952e7 *Iterator) ReadUint64() uint64 {
	return __obf_4ab56a99965952e7.__obf_bc7d04a45a032c40(__obf_4ab56a99965952e7.__obf_61df301d1f67ad73())
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_bc7d04a45a032c40(__obf_24309b3b0ff9ba22 byte) (__obf_59c72400ec1a6110 uint64) {
	__obf_0d7ccf72a2e94911 := __obf_03403867a04aca29[__obf_24309b3b0ff9ba22]
	if __obf_0d7ccf72a2e94911 == 0 {
		__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
		return 0 // single zero
	}
	if __obf_0d7ccf72a2e94911 == __obf_78b317d8e70f5d5a {
		__obf_4ab56a99965952e7.
			ReportError("readUint64", "unexpected character: "+string([]byte{byte(__obf_0d7ccf72a2e94911)}))
		return
	}
	__obf_b80d5217e1232943 := uint64(__obf_0d7ccf72a2e94911)
	if __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc-__obf_4ab56a99965952e7.__obf_5e22d6270d27491f > 10 {
		__obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f
		__obf_8b3d3761235ee860 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_8b3d3761235ee860 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943
		}
		__obf_82c6e05b9d226c58++
		__obf_bd7a44b446c0f2eb := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_bd7a44b446c0f2eb == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*10 + uint64(__obf_8b3d3761235ee860)
		}
		__obf_82c6e05b9d226c58++ //iter.head = i + 1
		//value = value * 100 + uint32(ind2) * 10 + uint32(ind3)
		__obf_bc5725ebafdfb1b8 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_bc5725ebafdfb1b8 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*100 + uint64(__obf_8b3d3761235ee860)*10 + uint64(__obf_bd7a44b446c0f2eb)
		}
		__obf_82c6e05b9d226c58++
		__obf_5c788bf8b20b2b60 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_5c788bf8b20b2b60 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*1000 + uint64(__obf_8b3d3761235ee860)*100 + uint64(__obf_bd7a44b446c0f2eb)*10 + uint64(__obf_bc5725ebafdfb1b8)
		}
		__obf_82c6e05b9d226c58++
		__obf_79941b1d34ce7165 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_79941b1d34ce7165 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*10000 + uint64(__obf_8b3d3761235ee860)*1000 + uint64(__obf_bd7a44b446c0f2eb)*100 + uint64(__obf_bc5725ebafdfb1b8)*10 + uint64(__obf_5c788bf8b20b2b60)
		}
		__obf_82c6e05b9d226c58++
		__obf_9639592f55bb3c69 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_9639592f55bb3c69 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*100000 + uint64(__obf_8b3d3761235ee860)*10000 + uint64(__obf_bd7a44b446c0f2eb)*1000 + uint64(__obf_bc5725ebafdfb1b8)*100 + uint64(__obf_5c788bf8b20b2b60)*10 + uint64(__obf_79941b1d34ce7165)
		}
		__obf_82c6e05b9d226c58++
		__obf_fab340f78ca73f45 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		if __obf_fab340f78ca73f45 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943*1000000 + uint64(__obf_8b3d3761235ee860)*100000 + uint64(__obf_bd7a44b446c0f2eb)*10000 + uint64(__obf_bc5725ebafdfb1b8)*1000 + uint64(__obf_5c788bf8b20b2b60)*100 + uint64(__obf_79941b1d34ce7165)*10 + uint64(__obf_9639592f55bb3c69)
		}
		__obf_82c6e05b9d226c58++
		__obf_f64f5fdbade2b834 := __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
		__obf_b80d5217e1232943 = __obf_b80d5217e1232943*10000000 + uint64(__obf_8b3d3761235ee860)*1000000 + uint64(__obf_bd7a44b446c0f2eb)*100000 + uint64(__obf_bc5725ebafdfb1b8)*10000 + uint64(__obf_5c788bf8b20b2b60)*1000 + uint64(__obf_79941b1d34ce7165)*100 + uint64(__obf_9639592f55bb3c69)*10 + uint64(__obf_fab340f78ca73f45)
		__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
		if __obf_f64f5fdbade2b834 == __obf_78b317d8e70f5d5a {
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943
		}
	}
	for {
		for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
			__obf_0d7ccf72a2e94911 = __obf_03403867a04aca29[__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]]
			if __obf_0d7ccf72a2e94911 == __obf_78b317d8e70f5d5a {
				__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
				__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
				return __obf_b80d5217e1232943
			}
			if __obf_b80d5217e1232943 > __obf_92013c25906febd6 {
				__obf_50aaa7aa198ad0ad := (__obf_b80d5217e1232943 << 3) + (__obf_b80d5217e1232943 << 1) + uint64(__obf_0d7ccf72a2e94911)
				if __obf_50aaa7aa198ad0ad < __obf_b80d5217e1232943 {
					__obf_4ab56a99965952e7.
						ReportError("readUint64", "overflow")
					return
				}
				__obf_b80d5217e1232943 = __obf_50aaa7aa198ad0ad
				continue
			}
			__obf_b80d5217e1232943 = (__obf_b80d5217e1232943 << 3) + (__obf_b80d5217e1232943 << 1) + uint64(__obf_0d7ccf72a2e94911)
		}
		if !__obf_4ab56a99965952e7.__obf_08a0b850abb3abbd() {
			__obf_4ab56a99965952e7.__obf_f779acc81d83903b()
			return __obf_b80d5217e1232943
		}
	}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_f779acc81d83903b() {
	if __obf_4ab56a99965952e7.__obf_5e22d6270d27491f < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc && __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_4ab56a99965952e7.__obf_5e22d6270d27491f] == '.' {
		__obf_4ab56a99965952e7.
			ReportError("assertInteger", "can not decode float as int")
	}
}
