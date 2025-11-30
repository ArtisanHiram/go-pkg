package __obf_c3cd04a15c56f32f

var __obf_c1c5e3ca99cf8f77 []uint32

func init() {
	__obf_c1c5e3ca99cf8f77 = make([]uint32, 1000)
	for __obf_28d099df85f083a8 := uint32(0); __obf_28d099df85f083a8 < 1000; __obf_28d099df85f083a8++ {
		__obf_c1c5e3ca99cf8f77[__obf_28d099df85f083a8] = (((__obf_28d099df85f083a8 / 100) + '0') << 16) + ((((__obf_28d099df85f083a8 / 10) % 10) + '0') << 8) + __obf_28d099df85f083a8%10 + '0'
		if __obf_28d099df85f083a8 < 10 {
			__obf_c1c5e3ca99cf8f77[__obf_28d099df85f083a8] += 2 << 24
		} else if __obf_28d099df85f083a8 < 100 {
			__obf_c1c5e3ca99cf8f77[__obf_28d099df85f083a8] += 1 << 24
		}
	}
}

func __obf_d838aa64c4a879e5(__obf_01960ab63e397c41 []byte, __obf_f967fc9e700d5e33 uint32) []byte {
	__obf_91c0b88ab32aec51 := __obf_f967fc9e700d5e33 >> 24
	switch __obf_91c0b88ab32aec51 {
	case 0:
		__obf_01960ab63e397c41 = append(__obf_01960ab63e397c41, byte(__obf_f967fc9e700d5e33>>16), byte(__obf_f967fc9e700d5e33>>8))
	case 1:
		__obf_01960ab63e397c41 = append(__obf_01960ab63e397c41, byte(__obf_f967fc9e700d5e33>>8))
	}
	__obf_01960ab63e397c41 = append(__obf_01960ab63e397c41, byte(__obf_f967fc9e700d5e33))
	return __obf_01960ab63e397c41
}

func __obf_6e246dbe68241c26(__obf_ace979f6622823f3 []byte, __obf_f967fc9e700d5e33 uint32) []byte {
	return append(__obf_ace979f6622823f3, byte(__obf_f967fc9e700d5e33>>16), byte(__obf_f967fc9e700d5e33>>8), byte(__obf_f967fc9e700d5e33))
}

// WriteUint8 write uint8 to stream
func (__obf_2361f5214e84c60f *Stream) WriteUint8(__obf_d59813f23ad740a8 uint8) {
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3,

	// WriteInt8 write int8 to stream
	__obf_c1c5e3ca99cf8f77[__obf_d59813f23ad740a8])
}

func (__obf_2361f5214e84c60f *Stream) WriteInt8(__obf_25abfc6f5c3d1493 int8) {
	var __obf_d59813f23ad740a8 uint8
	if __obf_25abfc6f5c3d1493 < 0 {
		__obf_d59813f23ad740a8 = uint8(-__obf_25abfc6f5c3d1493)
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, '-')
	} else {
		__obf_d59813f23ad740a8 = uint8(__obf_25abfc6f5c3d1493)
	}
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3,

	// WriteUint16 write uint16 to stream
	__obf_c1c5e3ca99cf8f77[__obf_d59813f23ad740a8])
}

func (__obf_2361f5214e84c60f *Stream) WriteUint16(__obf_d59813f23ad740a8 uint16) {
	__obf_80885b1182916b4b := __obf_d59813f23ad740a8 / 1000
	if __obf_80885b1182916b4b == 0 {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_d59813f23ad740a8])
		return
	}
	__obf_c82eb00c4b7bf382 := __obf_d59813f23ad740a8 - __obf_80885b1182916b4b*1000
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_80885b1182916b4b])
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3,

	// WriteInt16 write int16 to stream
	__obf_c1c5e3ca99cf8f77[__obf_c82eb00c4b7bf382])
}

func (__obf_2361f5214e84c60f *Stream) WriteInt16(__obf_25abfc6f5c3d1493 int16) {
	var __obf_d59813f23ad740a8 uint16
	if __obf_25abfc6f5c3d1493 < 0 {
		__obf_d59813f23ad740a8 = uint16(-__obf_25abfc6f5c3d1493)
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, '-')
	} else {
		__obf_d59813f23ad740a8 = uint16(__obf_25abfc6f5c3d1493)
	}
	__obf_2361f5214e84c60f.
		WriteUint16(__obf_d59813f23ad740a8)
}

// WriteUint32 write uint32 to stream
func (__obf_2361f5214e84c60f *Stream) WriteUint32(__obf_d59813f23ad740a8 uint32) {
	__obf_80885b1182916b4b := __obf_d59813f23ad740a8 / 1000
	if __obf_80885b1182916b4b == 0 {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_d59813f23ad740a8])
		return
	}
	__obf_c82eb00c4b7bf382 := __obf_d59813f23ad740a8 - __obf_80885b1182916b4b*1000
	__obf_c67883099db3e995 := __obf_80885b1182916b4b / 1000
	if __obf_c67883099db3e995 == 0 {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_80885b1182916b4b])
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_c82eb00c4b7bf382])
		return
	}
	__obf_695eff5f2dce5e40 := __obf_80885b1182916b4b - __obf_c67883099db3e995*1000
	__obf_fa9c8e2cdaf3b04d := __obf_c67883099db3e995 / 1000
	if __obf_fa9c8e2cdaf3b04d == 0 {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_c67883099db3e995])
	} else {
		__obf_40180a7409751e64 := __obf_c67883099db3e995 - __obf_fa9c8e2cdaf3b04d*1000
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, byte(__obf_fa9c8e2cdaf3b04d+'0'))
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_40180a7409751e64])
	}
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_695eff5f2dce5e40])
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3,

	// WriteInt32 write int32 to stream
	__obf_c1c5e3ca99cf8f77[__obf_c82eb00c4b7bf382])
}

func (__obf_2361f5214e84c60f *Stream) WriteInt32(__obf_25abfc6f5c3d1493 int32) {
	var __obf_d59813f23ad740a8 uint32
	if __obf_25abfc6f5c3d1493 < 0 {
		__obf_d59813f23ad740a8 = uint32(-__obf_25abfc6f5c3d1493)
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, '-')
	} else {
		__obf_d59813f23ad740a8 = uint32(__obf_25abfc6f5c3d1493)
	}
	__obf_2361f5214e84c60f.
		WriteUint32(__obf_d59813f23ad740a8)
}

// WriteUint64 write uint64 to stream
func (__obf_2361f5214e84c60f *Stream) WriteUint64(__obf_d59813f23ad740a8 uint64) {
	__obf_80885b1182916b4b := __obf_d59813f23ad740a8 / 1000
	if __obf_80885b1182916b4b == 0 {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_d59813f23ad740a8])
		return
	}
	__obf_c82eb00c4b7bf382 := __obf_d59813f23ad740a8 - __obf_80885b1182916b4b*1000
	__obf_c67883099db3e995 := __obf_80885b1182916b4b / 1000
	if __obf_c67883099db3e995 == 0 {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_80885b1182916b4b])
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_c82eb00c4b7bf382])
		return
	}
	__obf_695eff5f2dce5e40 := __obf_80885b1182916b4b - __obf_c67883099db3e995*1000
	__obf_fa9c8e2cdaf3b04d := __obf_c67883099db3e995 / 1000
	if __obf_fa9c8e2cdaf3b04d == 0 {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_c67883099db3e995])
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_695eff5f2dce5e40])
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_c82eb00c4b7bf382])
		return
	}
	__obf_40180a7409751e64 := __obf_c67883099db3e995 - __obf_fa9c8e2cdaf3b04d*1000
	__obf_f56717db9b83fee6 := __obf_fa9c8e2cdaf3b04d / 1000
	if __obf_f56717db9b83fee6 == 0 {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_fa9c8e2cdaf3b04d])
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_40180a7409751e64])
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_695eff5f2dce5e40])
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_c82eb00c4b7bf382])
		return
	}
	__obf_f4656c57e6ab361a := __obf_fa9c8e2cdaf3b04d - __obf_f56717db9b83fee6*1000
	__obf_4f40250b7a0f2053 := __obf_f56717db9b83fee6 / 1000
	if __obf_4f40250b7a0f2053 == 0 {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_f56717db9b83fee6])
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_f4656c57e6ab361a])
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_40180a7409751e64])
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_695eff5f2dce5e40])
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_c82eb00c4b7bf382])
		return
	}
	__obf_ec0e72611cf88bef := __obf_f56717db9b83fee6 - __obf_4f40250b7a0f2053*1000
	__obf_607b279d8be38c37 := __obf_4f40250b7a0f2053 / 1000
	if __obf_607b279d8be38c37 == 0 {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_4f40250b7a0f2053])
	} else {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_d838aa64c4a879e5(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_607b279d8be38c37])
		__obf_f9788c1b57f8a7c2 := __obf_4f40250b7a0f2053 - __obf_607b279d8be38c37*1000
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_f9788c1b57f8a7c2])
	}
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_ec0e72611cf88bef])
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_f4656c57e6ab361a])
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_40180a7409751e64])
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_c1c5e3ca99cf8f77[__obf_695eff5f2dce5e40])
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_6e246dbe68241c26(__obf_2361f5214e84c60f.__obf_ace979f6622823f3,

	// WriteInt64 write int64 to stream
	__obf_c1c5e3ca99cf8f77[__obf_c82eb00c4b7bf382])
}

func (__obf_2361f5214e84c60f *Stream) WriteInt64(__obf_25abfc6f5c3d1493 int64) {
	var __obf_d59813f23ad740a8 uint64
	if __obf_25abfc6f5c3d1493 < 0 {
		__obf_d59813f23ad740a8 = uint64(-__obf_25abfc6f5c3d1493)
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, '-')
	} else {
		__obf_d59813f23ad740a8 = uint64(__obf_25abfc6f5c3d1493)
	}
	__obf_2361f5214e84c60f.
		WriteUint64(__obf_d59813f23ad740a8)
}

// WriteInt write int to stream
func (__obf_2361f5214e84c60f *Stream) WriteInt(__obf_d59813f23ad740a8 int) {
	__obf_2361f5214e84c60f.
		WriteInt64(int64(__obf_d59813f23ad740a8))
}

// WriteUint write uint to stream
func (__obf_2361f5214e84c60f *Stream) WriteUint(__obf_d59813f23ad740a8 uint) {
	__obf_2361f5214e84c60f.
		WriteUint64(uint64(__obf_d59813f23ad740a8))
}
