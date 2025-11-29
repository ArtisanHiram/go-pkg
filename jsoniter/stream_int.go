package __obf_5b802ce8d9ba56d6

var __obf_4f6cd78097d2d97b []uint32

func init() {
	__obf_4f6cd78097d2d97b = make([]uint32, 1000)
	for __obf_2deec7c38ffb6ae3 := uint32(0); __obf_2deec7c38ffb6ae3 < 1000; __obf_2deec7c38ffb6ae3++ {
		__obf_4f6cd78097d2d97b[__obf_2deec7c38ffb6ae3] = (((__obf_2deec7c38ffb6ae3 / 100) + '0') << 16) + ((((__obf_2deec7c38ffb6ae3 / 10) % 10) + '0') << 8) + __obf_2deec7c38ffb6ae3%10 + '0'
		if __obf_2deec7c38ffb6ae3 < 10 {
			__obf_4f6cd78097d2d97b[__obf_2deec7c38ffb6ae3] += 2 << 24
		} else if __obf_2deec7c38ffb6ae3 < 100 {
			__obf_4f6cd78097d2d97b[__obf_2deec7c38ffb6ae3] += 1 << 24
		}
	}
}

func __obf_69863fd65526dc7a(__obf_1239dd43ead32055 []byte, __obf_91fc460a138dcae7 uint32) []byte {
	__obf_4773198b4b8f9524 := __obf_91fc460a138dcae7 >> 24
	switch __obf_4773198b4b8f9524 {
	case 0:
		__obf_1239dd43ead32055 = append(__obf_1239dd43ead32055, byte(__obf_91fc460a138dcae7>>16), byte(__obf_91fc460a138dcae7>>8))
	case 1:
		__obf_1239dd43ead32055 = append(__obf_1239dd43ead32055, byte(__obf_91fc460a138dcae7>>8))
	}
	__obf_1239dd43ead32055 = append(__obf_1239dd43ead32055, byte(__obf_91fc460a138dcae7))
	return __obf_1239dd43ead32055
}

func __obf_ed34828e637d2894(__obf_9fc06d9180f0daca []byte, __obf_91fc460a138dcae7 uint32) []byte {
	return append(__obf_9fc06d9180f0daca, byte(__obf_91fc460a138dcae7>>16), byte(__obf_91fc460a138dcae7>>8), byte(__obf_91fc460a138dcae7))
}

// WriteUint8 write uint8 to stream
func (__obf_00fc491caa967a74 *Stream) WriteUint8(__obf_5406b11e33b9d1d3 uint8) {
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca,

	// WriteInt8 write int8 to stream
	__obf_4f6cd78097d2d97b[__obf_5406b11e33b9d1d3])
}

func (__obf_00fc491caa967a74 *Stream) WriteInt8(__obf_8a4b27453b1536a5 int8) {
	var __obf_5406b11e33b9d1d3 uint8
	if __obf_8a4b27453b1536a5 < 0 {
		__obf_5406b11e33b9d1d3 = uint8(-__obf_8a4b27453b1536a5)
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, '-')
	} else {
		__obf_5406b11e33b9d1d3 = uint8(__obf_8a4b27453b1536a5)
	}
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca,

	// WriteUint16 write uint16 to stream
	__obf_4f6cd78097d2d97b[__obf_5406b11e33b9d1d3])
}

func (__obf_00fc491caa967a74 *Stream) WriteUint16(__obf_5406b11e33b9d1d3 uint16) {
	__obf_41ef579bef2baad9 := __obf_5406b11e33b9d1d3 / 1000
	if __obf_41ef579bef2baad9 == 0 {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_5406b11e33b9d1d3])
		return
	}
	__obf_3bedbb7e2d458433 := __obf_5406b11e33b9d1d3 - __obf_41ef579bef2baad9*1000
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_41ef579bef2baad9])
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca,

	// WriteInt16 write int16 to stream
	__obf_4f6cd78097d2d97b[__obf_3bedbb7e2d458433])
}

func (__obf_00fc491caa967a74 *Stream) WriteInt16(__obf_8a4b27453b1536a5 int16) {
	var __obf_5406b11e33b9d1d3 uint16
	if __obf_8a4b27453b1536a5 < 0 {
		__obf_5406b11e33b9d1d3 = uint16(-__obf_8a4b27453b1536a5)
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, '-')
	} else {
		__obf_5406b11e33b9d1d3 = uint16(__obf_8a4b27453b1536a5)
	}
	__obf_00fc491caa967a74.
		WriteUint16(__obf_5406b11e33b9d1d3)
}

// WriteUint32 write uint32 to stream
func (__obf_00fc491caa967a74 *Stream) WriteUint32(__obf_5406b11e33b9d1d3 uint32) {
	__obf_41ef579bef2baad9 := __obf_5406b11e33b9d1d3 / 1000
	if __obf_41ef579bef2baad9 == 0 {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_5406b11e33b9d1d3])
		return
	}
	__obf_3bedbb7e2d458433 := __obf_5406b11e33b9d1d3 - __obf_41ef579bef2baad9*1000
	__obf_3bddd139cbcc9a1d := __obf_41ef579bef2baad9 / 1000
	if __obf_3bddd139cbcc9a1d == 0 {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_41ef579bef2baad9])
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_3bedbb7e2d458433])
		return
	}
	__obf_63efaa8efa743d97 := __obf_41ef579bef2baad9 - __obf_3bddd139cbcc9a1d*1000
	__obf_180314ee60be84f6 := __obf_3bddd139cbcc9a1d / 1000
	if __obf_180314ee60be84f6 == 0 {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_3bddd139cbcc9a1d])
	} else {
		__obf_a8ea47959997eab3 := __obf_3bddd139cbcc9a1d - __obf_180314ee60be84f6*1000
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, byte(__obf_180314ee60be84f6+'0'))
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_a8ea47959997eab3])
	}
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_63efaa8efa743d97])
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca,

	// WriteInt32 write int32 to stream
	__obf_4f6cd78097d2d97b[__obf_3bedbb7e2d458433])
}

func (__obf_00fc491caa967a74 *Stream) WriteInt32(__obf_8a4b27453b1536a5 int32) {
	var __obf_5406b11e33b9d1d3 uint32
	if __obf_8a4b27453b1536a5 < 0 {
		__obf_5406b11e33b9d1d3 = uint32(-__obf_8a4b27453b1536a5)
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, '-')
	} else {
		__obf_5406b11e33b9d1d3 = uint32(__obf_8a4b27453b1536a5)
	}
	__obf_00fc491caa967a74.
		WriteUint32(__obf_5406b11e33b9d1d3)
}

// WriteUint64 write uint64 to stream
func (__obf_00fc491caa967a74 *Stream) WriteUint64(__obf_5406b11e33b9d1d3 uint64) {
	__obf_41ef579bef2baad9 := __obf_5406b11e33b9d1d3 / 1000
	if __obf_41ef579bef2baad9 == 0 {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_5406b11e33b9d1d3])
		return
	}
	__obf_3bedbb7e2d458433 := __obf_5406b11e33b9d1d3 - __obf_41ef579bef2baad9*1000
	__obf_3bddd139cbcc9a1d := __obf_41ef579bef2baad9 / 1000
	if __obf_3bddd139cbcc9a1d == 0 {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_41ef579bef2baad9])
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_3bedbb7e2d458433])
		return
	}
	__obf_63efaa8efa743d97 := __obf_41ef579bef2baad9 - __obf_3bddd139cbcc9a1d*1000
	__obf_180314ee60be84f6 := __obf_3bddd139cbcc9a1d / 1000
	if __obf_180314ee60be84f6 == 0 {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_3bddd139cbcc9a1d])
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_63efaa8efa743d97])
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_3bedbb7e2d458433])
		return
	}
	__obf_a8ea47959997eab3 := __obf_3bddd139cbcc9a1d - __obf_180314ee60be84f6*1000
	__obf_5d8c00765f91eda2 := __obf_180314ee60be84f6 / 1000
	if __obf_5d8c00765f91eda2 == 0 {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_180314ee60be84f6])
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_a8ea47959997eab3])
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_63efaa8efa743d97])
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_3bedbb7e2d458433])
		return
	}
	__obf_ddee3385a56125fc := __obf_180314ee60be84f6 - __obf_5d8c00765f91eda2*1000
	__obf_1abc4c5eb5023458 := __obf_5d8c00765f91eda2 / 1000
	if __obf_1abc4c5eb5023458 == 0 {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_5d8c00765f91eda2])
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_ddee3385a56125fc])
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_a8ea47959997eab3])
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_63efaa8efa743d97])
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_3bedbb7e2d458433])
		return
	}
	__obf_ec6beea5553b5e9e := __obf_5d8c00765f91eda2 - __obf_1abc4c5eb5023458*1000
	__obf_b4527f4fb5848286 := __obf_1abc4c5eb5023458 / 1000
	if __obf_b4527f4fb5848286 == 0 {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_1abc4c5eb5023458])
	} else {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_69863fd65526dc7a(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_b4527f4fb5848286])
		__obf_e20c16bf1593b095 := __obf_1abc4c5eb5023458 - __obf_b4527f4fb5848286*1000
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_e20c16bf1593b095])
	}
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_ec6beea5553b5e9e])
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_ddee3385a56125fc])
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_a8ea47959997eab3])
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_4f6cd78097d2d97b[__obf_63efaa8efa743d97])
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_ed34828e637d2894(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca,

	// WriteInt64 write int64 to stream
	__obf_4f6cd78097d2d97b[__obf_3bedbb7e2d458433])
}

func (__obf_00fc491caa967a74 *Stream) WriteInt64(__obf_8a4b27453b1536a5 int64) {
	var __obf_5406b11e33b9d1d3 uint64
	if __obf_8a4b27453b1536a5 < 0 {
		__obf_5406b11e33b9d1d3 = uint64(-__obf_8a4b27453b1536a5)
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, '-')
	} else {
		__obf_5406b11e33b9d1d3 = uint64(__obf_8a4b27453b1536a5)
	}
	__obf_00fc491caa967a74.
		WriteUint64(__obf_5406b11e33b9d1d3)
}

// WriteInt write int to stream
func (__obf_00fc491caa967a74 *Stream) WriteInt(__obf_5406b11e33b9d1d3 int) {
	__obf_00fc491caa967a74.
		WriteInt64(int64(__obf_5406b11e33b9d1d3))
}

// WriteUint write uint to stream
func (__obf_00fc491caa967a74 *Stream) WriteUint(__obf_5406b11e33b9d1d3 uint) {
	__obf_00fc491caa967a74.
		WriteUint64(uint64(__obf_5406b11e33b9d1d3))
}
