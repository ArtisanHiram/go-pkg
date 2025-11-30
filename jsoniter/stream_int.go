package __obf_703d23ba520c3295

var __obf_45e84cd317d1e682 []uint32

func init() {
	__obf_45e84cd317d1e682 = make([]uint32, 1000)
	for __obf_b0a5d2bd48690f1d := uint32(0); __obf_b0a5d2bd48690f1d < 1000; __obf_b0a5d2bd48690f1d++ {
		__obf_45e84cd317d1e682[__obf_b0a5d2bd48690f1d] = (((__obf_b0a5d2bd48690f1d / 100) + '0') << 16) + ((((__obf_b0a5d2bd48690f1d / 10) % 10) + '0') << 8) + __obf_b0a5d2bd48690f1d%10 + '0'
		if __obf_b0a5d2bd48690f1d < 10 {
			__obf_45e84cd317d1e682[__obf_b0a5d2bd48690f1d] += 2 << 24
		} else if __obf_b0a5d2bd48690f1d < 100 {
			__obf_45e84cd317d1e682[__obf_b0a5d2bd48690f1d] += 1 << 24
		}
	}
}

func __obf_d32fe4a1a5bf1668(__obf_7cf9eb0269a38f30 []byte, __obf_9cfb140618a077d4 uint32) []byte {
	__obf_bb31006da87bf9e2 := __obf_9cfb140618a077d4 >> 24
	switch __obf_bb31006da87bf9e2 {
	case 0:
		__obf_7cf9eb0269a38f30 = append(__obf_7cf9eb0269a38f30, byte(__obf_9cfb140618a077d4>>16), byte(__obf_9cfb140618a077d4>>8))
	case 1:
		__obf_7cf9eb0269a38f30 = append(__obf_7cf9eb0269a38f30, byte(__obf_9cfb140618a077d4>>8))
	}
	__obf_7cf9eb0269a38f30 = append(__obf_7cf9eb0269a38f30, byte(__obf_9cfb140618a077d4))
	return __obf_7cf9eb0269a38f30
}

func __obf_5574ec74d0150a75(__obf_a065f8e0da5f5952 []byte, __obf_9cfb140618a077d4 uint32) []byte {
	return append(__obf_a065f8e0da5f5952, byte(__obf_9cfb140618a077d4>>16), byte(__obf_9cfb140618a077d4>>8), byte(__obf_9cfb140618a077d4))
}

// WriteUint8 write uint8 to stream
func (__obf_9a34f62857fb1d1d *Stream) WriteUint8(__obf_b924538fffe5fd64 uint8) {
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952,

	// WriteInt8 write int8 to stream
	__obf_45e84cd317d1e682[__obf_b924538fffe5fd64])
}

func (__obf_9a34f62857fb1d1d *Stream) WriteInt8(__obf_c237dd482419bbf8 int8) {
	var __obf_b924538fffe5fd64 uint8
	if __obf_c237dd482419bbf8 < 0 {
		__obf_b924538fffe5fd64 = uint8(-__obf_c237dd482419bbf8)
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, '-')
	} else {
		__obf_b924538fffe5fd64 = uint8(__obf_c237dd482419bbf8)
	}
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952,

	// WriteUint16 write uint16 to stream
	__obf_45e84cd317d1e682[__obf_b924538fffe5fd64])
}

func (__obf_9a34f62857fb1d1d *Stream) WriteUint16(__obf_b924538fffe5fd64 uint16) {
	__obf_a3ddb1d12db95c12 := __obf_b924538fffe5fd64 / 1000
	if __obf_a3ddb1d12db95c12 == 0 {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_b924538fffe5fd64])
		return
	}
	__obf_96717f4412ba442f := __obf_b924538fffe5fd64 - __obf_a3ddb1d12db95c12*1000
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_a3ddb1d12db95c12])
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952,

	// WriteInt16 write int16 to stream
	__obf_45e84cd317d1e682[__obf_96717f4412ba442f])
}

func (__obf_9a34f62857fb1d1d *Stream) WriteInt16(__obf_c237dd482419bbf8 int16) {
	var __obf_b924538fffe5fd64 uint16
	if __obf_c237dd482419bbf8 < 0 {
		__obf_b924538fffe5fd64 = uint16(-__obf_c237dd482419bbf8)
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, '-')
	} else {
		__obf_b924538fffe5fd64 = uint16(__obf_c237dd482419bbf8)
	}
	__obf_9a34f62857fb1d1d.
		WriteUint16(__obf_b924538fffe5fd64)
}

// WriteUint32 write uint32 to stream
func (__obf_9a34f62857fb1d1d *Stream) WriteUint32(__obf_b924538fffe5fd64 uint32) {
	__obf_a3ddb1d12db95c12 := __obf_b924538fffe5fd64 / 1000
	if __obf_a3ddb1d12db95c12 == 0 {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_b924538fffe5fd64])
		return
	}
	__obf_96717f4412ba442f := __obf_b924538fffe5fd64 - __obf_a3ddb1d12db95c12*1000
	__obf_365f8cd086612431 := __obf_a3ddb1d12db95c12 / 1000
	if __obf_365f8cd086612431 == 0 {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_a3ddb1d12db95c12])
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_96717f4412ba442f])
		return
	}
	__obf_f2b2cb07840ef812 := __obf_a3ddb1d12db95c12 - __obf_365f8cd086612431*1000
	__obf_583ef97182359d8a := __obf_365f8cd086612431 / 1000
	if __obf_583ef97182359d8a == 0 {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_365f8cd086612431])
	} else {
		__obf_fce4bfb518ff9ce7 := __obf_365f8cd086612431 - __obf_583ef97182359d8a*1000
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, byte(__obf_583ef97182359d8a+'0'))
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_fce4bfb518ff9ce7])
	}
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_f2b2cb07840ef812])
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952,

	// WriteInt32 write int32 to stream
	__obf_45e84cd317d1e682[__obf_96717f4412ba442f])
}

func (__obf_9a34f62857fb1d1d *Stream) WriteInt32(__obf_c237dd482419bbf8 int32) {
	var __obf_b924538fffe5fd64 uint32
	if __obf_c237dd482419bbf8 < 0 {
		__obf_b924538fffe5fd64 = uint32(-__obf_c237dd482419bbf8)
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, '-')
	} else {
		__obf_b924538fffe5fd64 = uint32(__obf_c237dd482419bbf8)
	}
	__obf_9a34f62857fb1d1d.
		WriteUint32(__obf_b924538fffe5fd64)
}

// WriteUint64 write uint64 to stream
func (__obf_9a34f62857fb1d1d *Stream) WriteUint64(__obf_b924538fffe5fd64 uint64) {
	__obf_a3ddb1d12db95c12 := __obf_b924538fffe5fd64 / 1000
	if __obf_a3ddb1d12db95c12 == 0 {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_b924538fffe5fd64])
		return
	}
	__obf_96717f4412ba442f := __obf_b924538fffe5fd64 - __obf_a3ddb1d12db95c12*1000
	__obf_365f8cd086612431 := __obf_a3ddb1d12db95c12 / 1000
	if __obf_365f8cd086612431 == 0 {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_a3ddb1d12db95c12])
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_96717f4412ba442f])
		return
	}
	__obf_f2b2cb07840ef812 := __obf_a3ddb1d12db95c12 - __obf_365f8cd086612431*1000
	__obf_583ef97182359d8a := __obf_365f8cd086612431 / 1000
	if __obf_583ef97182359d8a == 0 {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_365f8cd086612431])
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_f2b2cb07840ef812])
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_96717f4412ba442f])
		return
	}
	__obf_fce4bfb518ff9ce7 := __obf_365f8cd086612431 - __obf_583ef97182359d8a*1000
	__obf_a9316c328e9a4740 := __obf_583ef97182359d8a / 1000
	if __obf_a9316c328e9a4740 == 0 {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_583ef97182359d8a])
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_fce4bfb518ff9ce7])
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_f2b2cb07840ef812])
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_96717f4412ba442f])
		return
	}
	__obf_3e7b59c731816097 := __obf_583ef97182359d8a - __obf_a9316c328e9a4740*1000
	__obf_f1361eba9f271c82 := __obf_a9316c328e9a4740 / 1000
	if __obf_f1361eba9f271c82 == 0 {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_a9316c328e9a4740])
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_3e7b59c731816097])
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_fce4bfb518ff9ce7])
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_f2b2cb07840ef812])
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_96717f4412ba442f])
		return
	}
	__obf_0272f291ec53ad4e := __obf_a9316c328e9a4740 - __obf_f1361eba9f271c82*1000
	__obf_67cb174585423d76 := __obf_f1361eba9f271c82 / 1000
	if __obf_67cb174585423d76 == 0 {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_f1361eba9f271c82])
	} else {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_d32fe4a1a5bf1668(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_67cb174585423d76])
		__obf_c1f9624e6bfd42c6 := __obf_f1361eba9f271c82 - __obf_67cb174585423d76*1000
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_c1f9624e6bfd42c6])
	}
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_0272f291ec53ad4e])
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_3e7b59c731816097])
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_fce4bfb518ff9ce7])
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_45e84cd317d1e682[__obf_f2b2cb07840ef812])
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_5574ec74d0150a75(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952,

	// WriteInt64 write int64 to stream
	__obf_45e84cd317d1e682[__obf_96717f4412ba442f])
}

func (__obf_9a34f62857fb1d1d *Stream) WriteInt64(__obf_c237dd482419bbf8 int64) {
	var __obf_b924538fffe5fd64 uint64
	if __obf_c237dd482419bbf8 < 0 {
		__obf_b924538fffe5fd64 = uint64(-__obf_c237dd482419bbf8)
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, '-')
	} else {
		__obf_b924538fffe5fd64 = uint64(__obf_c237dd482419bbf8)
	}
	__obf_9a34f62857fb1d1d.
		WriteUint64(__obf_b924538fffe5fd64)
}

// WriteInt write int to stream
func (__obf_9a34f62857fb1d1d *Stream) WriteInt(__obf_b924538fffe5fd64 int) {
	__obf_9a34f62857fb1d1d.
		WriteInt64(int64(__obf_b924538fffe5fd64))
}

// WriteUint write uint to stream
func (__obf_9a34f62857fb1d1d *Stream) WriteUint(__obf_b924538fffe5fd64 uint) {
	__obf_9a34f62857fb1d1d.
		WriteUint64(uint64(__obf_b924538fffe5fd64))
}
