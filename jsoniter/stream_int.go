package __obf_c7b79b12b33d8238

var __obf_8a8ef744eac3c7b4 []uint32

func init() {
	__obf_8a8ef744eac3c7b4 = make([]uint32, 1000)
	for __obf_a051269af3a541bb := uint32(0); __obf_a051269af3a541bb < 1000; __obf_a051269af3a541bb++ {
		__obf_8a8ef744eac3c7b4[__obf_a051269af3a541bb] = (((__obf_a051269af3a541bb / 100) + '0') << 16) + ((((__obf_a051269af3a541bb / 10) % 10) + '0') << 8) + __obf_a051269af3a541bb%10 + '0'
		if __obf_a051269af3a541bb < 10 {
			__obf_8a8ef744eac3c7b4[__obf_a051269af3a541bb] += 2 << 24
		} else if __obf_a051269af3a541bb < 100 {
			__obf_8a8ef744eac3c7b4[__obf_a051269af3a541bb] += 1 << 24
		}
	}
}

func __obf_d0931d9327448c4e(__obf_9421762504b7c038 []byte, __obf_bc7196b82ffbf1d3 uint32) []byte {
	__obf_28e7d8a603535072 := __obf_bc7196b82ffbf1d3 >> 24
	switch __obf_28e7d8a603535072 {
	case 0:
		__obf_9421762504b7c038 = append(__obf_9421762504b7c038, byte(__obf_bc7196b82ffbf1d3>>16), byte(__obf_bc7196b82ffbf1d3>>8))
	case 1:
		__obf_9421762504b7c038 = append(__obf_9421762504b7c038, byte(__obf_bc7196b82ffbf1d3>>8))
	}
	__obf_9421762504b7c038 = append(__obf_9421762504b7c038, byte(__obf_bc7196b82ffbf1d3))
	return __obf_9421762504b7c038
}

func __obf_9af927d7780d9de5(__obf_684d738bc3ac851a []byte, __obf_bc7196b82ffbf1d3 uint32) []byte {
	return append(__obf_684d738bc3ac851a, byte(__obf_bc7196b82ffbf1d3>>16), byte(__obf_bc7196b82ffbf1d3>>8), byte(__obf_bc7196b82ffbf1d3))
}

// WriteUint8 write uint8 to stream
func (__obf_d8f50bcfe87d8b47 *Stream) WriteUint8(__obf_35accd096612ebe4 uint8) {
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a,

	// WriteInt8 write int8 to stream
	__obf_8a8ef744eac3c7b4[__obf_35accd096612ebe4])
}

func (__obf_d8f50bcfe87d8b47 *Stream) WriteInt8(__obf_4f2c0b216dd14ef3 int8) {
	var __obf_35accd096612ebe4 uint8
	if __obf_4f2c0b216dd14ef3 < 0 {
		__obf_35accd096612ebe4 = uint8(-__obf_4f2c0b216dd14ef3)
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, '-')
	} else {
		__obf_35accd096612ebe4 = uint8(__obf_4f2c0b216dd14ef3)
	}
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a,

	// WriteUint16 write uint16 to stream
	__obf_8a8ef744eac3c7b4[__obf_35accd096612ebe4])
}

func (__obf_d8f50bcfe87d8b47 *Stream) WriteUint16(__obf_35accd096612ebe4 uint16) {
	__obf_6fa367c50b007a7a := __obf_35accd096612ebe4 / 1000
	if __obf_6fa367c50b007a7a == 0 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_35accd096612ebe4])
		return
	}
	__obf_be71319e64817e51 := __obf_35accd096612ebe4 - __obf_6fa367c50b007a7a*1000
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_6fa367c50b007a7a])
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a,

	// WriteInt16 write int16 to stream
	__obf_8a8ef744eac3c7b4[__obf_be71319e64817e51])
}

func (__obf_d8f50bcfe87d8b47 *Stream) WriteInt16(__obf_4f2c0b216dd14ef3 int16) {
	var __obf_35accd096612ebe4 uint16
	if __obf_4f2c0b216dd14ef3 < 0 {
		__obf_35accd096612ebe4 = uint16(-__obf_4f2c0b216dd14ef3)
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, '-')
	} else {
		__obf_35accd096612ebe4 = uint16(__obf_4f2c0b216dd14ef3)
	}
	__obf_d8f50bcfe87d8b47.
		WriteUint16(__obf_35accd096612ebe4)
}

// WriteUint32 write uint32 to stream
func (__obf_d8f50bcfe87d8b47 *Stream) WriteUint32(__obf_35accd096612ebe4 uint32) {
	__obf_6fa367c50b007a7a := __obf_35accd096612ebe4 / 1000
	if __obf_6fa367c50b007a7a == 0 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_35accd096612ebe4])
		return
	}
	__obf_be71319e64817e51 := __obf_35accd096612ebe4 - __obf_6fa367c50b007a7a*1000
	__obf_e1e9dc13ea735118 := __obf_6fa367c50b007a7a / 1000
	if __obf_e1e9dc13ea735118 == 0 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_6fa367c50b007a7a])
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_be71319e64817e51])
		return
	}
	__obf_a818cce3414b2179 := __obf_6fa367c50b007a7a - __obf_e1e9dc13ea735118*1000
	__obf_3d0b3efc00ea84bc := __obf_e1e9dc13ea735118 / 1000
	if __obf_3d0b3efc00ea84bc == 0 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_e1e9dc13ea735118])
	} else {
		__obf_fd77516aa9628ec3 := __obf_e1e9dc13ea735118 - __obf_3d0b3efc00ea84bc*1000
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, byte(__obf_3d0b3efc00ea84bc+'0'))
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_fd77516aa9628ec3])
	}
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_a818cce3414b2179])
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a,

	// WriteInt32 write int32 to stream
	__obf_8a8ef744eac3c7b4[__obf_be71319e64817e51])
}

func (__obf_d8f50bcfe87d8b47 *Stream) WriteInt32(__obf_4f2c0b216dd14ef3 int32) {
	var __obf_35accd096612ebe4 uint32
	if __obf_4f2c0b216dd14ef3 < 0 {
		__obf_35accd096612ebe4 = uint32(-__obf_4f2c0b216dd14ef3)
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, '-')
	} else {
		__obf_35accd096612ebe4 = uint32(__obf_4f2c0b216dd14ef3)
	}
	__obf_d8f50bcfe87d8b47.
		WriteUint32(__obf_35accd096612ebe4)
}

// WriteUint64 write uint64 to stream
func (__obf_d8f50bcfe87d8b47 *Stream) WriteUint64(__obf_35accd096612ebe4 uint64) {
	__obf_6fa367c50b007a7a := __obf_35accd096612ebe4 / 1000
	if __obf_6fa367c50b007a7a == 0 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_35accd096612ebe4])
		return
	}
	__obf_be71319e64817e51 := __obf_35accd096612ebe4 - __obf_6fa367c50b007a7a*1000
	__obf_e1e9dc13ea735118 := __obf_6fa367c50b007a7a / 1000
	if __obf_e1e9dc13ea735118 == 0 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_6fa367c50b007a7a])
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_be71319e64817e51])
		return
	}
	__obf_a818cce3414b2179 := __obf_6fa367c50b007a7a - __obf_e1e9dc13ea735118*1000
	__obf_3d0b3efc00ea84bc := __obf_e1e9dc13ea735118 / 1000
	if __obf_3d0b3efc00ea84bc == 0 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_e1e9dc13ea735118])
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_a818cce3414b2179])
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_be71319e64817e51])
		return
	}
	__obf_fd77516aa9628ec3 := __obf_e1e9dc13ea735118 - __obf_3d0b3efc00ea84bc*1000
	__obf_ba7bcb270c6d9334 := __obf_3d0b3efc00ea84bc / 1000
	if __obf_ba7bcb270c6d9334 == 0 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_3d0b3efc00ea84bc])
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_fd77516aa9628ec3])
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_a818cce3414b2179])
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_be71319e64817e51])
		return
	}
	__obf_356d95099b726ab3 := __obf_3d0b3efc00ea84bc - __obf_ba7bcb270c6d9334*1000
	__obf_a3e05a484914233d := __obf_ba7bcb270c6d9334 / 1000
	if __obf_a3e05a484914233d == 0 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_ba7bcb270c6d9334])
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_356d95099b726ab3])
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_fd77516aa9628ec3])
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_a818cce3414b2179])
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_be71319e64817e51])
		return
	}
	__obf_64c462aed55117eb := __obf_ba7bcb270c6d9334 - __obf_a3e05a484914233d*1000
	__obf_104ed5ca1648ea56 := __obf_a3e05a484914233d / 1000
	if __obf_104ed5ca1648ea56 == 0 {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_a3e05a484914233d])
	} else {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d0931d9327448c4e(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_104ed5ca1648ea56])
		__obf_27a594eb79e40256 := __obf_a3e05a484914233d - __obf_104ed5ca1648ea56*1000
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_27a594eb79e40256])
	}
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_64c462aed55117eb])
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_356d95099b726ab3])
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_fd77516aa9628ec3])
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_8a8ef744eac3c7b4[__obf_a818cce3414b2179])
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_9af927d7780d9de5(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a,

	// WriteInt64 write int64 to stream
	__obf_8a8ef744eac3c7b4[__obf_be71319e64817e51])
}

func (__obf_d8f50bcfe87d8b47 *Stream) WriteInt64(__obf_4f2c0b216dd14ef3 int64) {
	var __obf_35accd096612ebe4 uint64
	if __obf_4f2c0b216dd14ef3 < 0 {
		__obf_35accd096612ebe4 = uint64(-__obf_4f2c0b216dd14ef3)
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, '-')
	} else {
		__obf_35accd096612ebe4 = uint64(__obf_4f2c0b216dd14ef3)
	}
	__obf_d8f50bcfe87d8b47.
		WriteUint64(__obf_35accd096612ebe4)
}

// WriteInt write int to stream
func (__obf_d8f50bcfe87d8b47 *Stream) WriteInt(__obf_35accd096612ebe4 int) {
	__obf_d8f50bcfe87d8b47.
		WriteInt64(int64(__obf_35accd096612ebe4))
}

// WriteUint write uint to stream
func (__obf_d8f50bcfe87d8b47 *Stream) WriteUint(__obf_35accd096612ebe4 uint) {
	__obf_d8f50bcfe87d8b47.
		WriteUint64(uint64(__obf_35accd096612ebe4))
}
