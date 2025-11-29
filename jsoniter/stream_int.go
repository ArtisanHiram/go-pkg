package __obf_91620b895eeff9ed

var __obf_209c14c49f9e1dfb []uint32

func init() {
	__obf_209c14c49f9e1dfb = make([]uint32, 1000)
	for __obf_5aa5c8829b97f182 := uint32(0); __obf_5aa5c8829b97f182 < 1000; __obf_5aa5c8829b97f182++ {
		__obf_209c14c49f9e1dfb[__obf_5aa5c8829b97f182] = (((__obf_5aa5c8829b97f182 / 100) + '0') << 16) + ((((__obf_5aa5c8829b97f182 / 10) % 10) + '0') << 8) + __obf_5aa5c8829b97f182%10 + '0'
		if __obf_5aa5c8829b97f182 < 10 {
			__obf_209c14c49f9e1dfb[__obf_5aa5c8829b97f182] += 2 << 24
		} else if __obf_5aa5c8829b97f182 < 100 {
			__obf_209c14c49f9e1dfb[__obf_5aa5c8829b97f182] += 1 << 24
		}
	}
}

func __obf_fd59d706acca82ca(__obf_3b06e5d61898f33c []byte, __obf_97f02ddd0770463f uint32) []byte {
	__obf_924c0b8ced3b6bc7 := __obf_97f02ddd0770463f >> 24
	switch __obf_924c0b8ced3b6bc7 {
	case 0:
		__obf_3b06e5d61898f33c = append(__obf_3b06e5d61898f33c, byte(__obf_97f02ddd0770463f>>16), byte(__obf_97f02ddd0770463f>>8))
	case 1:
		__obf_3b06e5d61898f33c = append(__obf_3b06e5d61898f33c, byte(__obf_97f02ddd0770463f>>8))
	}
	__obf_3b06e5d61898f33c = append(__obf_3b06e5d61898f33c, byte(__obf_97f02ddd0770463f))
	return __obf_3b06e5d61898f33c
}

func __obf_56cc32e8e17354a8(__obf_184433571fa55237 []byte, __obf_97f02ddd0770463f uint32) []byte {
	return append(__obf_184433571fa55237, byte(__obf_97f02ddd0770463f>>16), byte(__obf_97f02ddd0770463f>>8), byte(__obf_97f02ddd0770463f))
}

// WriteUint8 write uint8 to stream
func (__obf_850a7457bb739a32 *Stream) WriteUint8(__obf_bbfd2ac8618a6f0c uint8) {
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237,

	// WriteInt8 write int8 to stream
	__obf_209c14c49f9e1dfb[__obf_bbfd2ac8618a6f0c])
}

func (__obf_850a7457bb739a32 *Stream) WriteInt8(__obf_993dfc2f74defa46 int8) {
	var __obf_bbfd2ac8618a6f0c uint8
	if __obf_993dfc2f74defa46 < 0 {
		__obf_bbfd2ac8618a6f0c = uint8(-__obf_993dfc2f74defa46)
		__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, '-')
	} else {
		__obf_bbfd2ac8618a6f0c = uint8(__obf_993dfc2f74defa46)
	}
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237,

	// WriteUint16 write uint16 to stream
	__obf_209c14c49f9e1dfb[__obf_bbfd2ac8618a6f0c])
}

func (__obf_850a7457bb739a32 *Stream) WriteUint16(__obf_bbfd2ac8618a6f0c uint16) {
	__obf_73c1195b655f767d := __obf_bbfd2ac8618a6f0c / 1000
	if __obf_73c1195b655f767d == 0 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_bbfd2ac8618a6f0c])
		return
	}
	__obf_bd67c2557d95db9e := __obf_bbfd2ac8618a6f0c - __obf_73c1195b655f767d*1000
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_73c1195b655f767d])
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237,

	// WriteInt16 write int16 to stream
	__obf_209c14c49f9e1dfb[__obf_bd67c2557d95db9e])
}

func (__obf_850a7457bb739a32 *Stream) WriteInt16(__obf_993dfc2f74defa46 int16) {
	var __obf_bbfd2ac8618a6f0c uint16
	if __obf_993dfc2f74defa46 < 0 {
		__obf_bbfd2ac8618a6f0c = uint16(-__obf_993dfc2f74defa46)
		__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, '-')
	} else {
		__obf_bbfd2ac8618a6f0c = uint16(__obf_993dfc2f74defa46)
	}
	__obf_850a7457bb739a32.
		WriteUint16(__obf_bbfd2ac8618a6f0c)
}

// WriteUint32 write uint32 to stream
func (__obf_850a7457bb739a32 *Stream) WriteUint32(__obf_bbfd2ac8618a6f0c uint32) {
	__obf_73c1195b655f767d := __obf_bbfd2ac8618a6f0c / 1000
	if __obf_73c1195b655f767d == 0 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_bbfd2ac8618a6f0c])
		return
	}
	__obf_bd67c2557d95db9e := __obf_bbfd2ac8618a6f0c - __obf_73c1195b655f767d*1000
	__obf_1d9e78caa969d04e := __obf_73c1195b655f767d / 1000
	if __obf_1d9e78caa969d04e == 0 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_73c1195b655f767d])
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_bd67c2557d95db9e])
		return
	}
	__obf_c3f7781dafa5b032 := __obf_73c1195b655f767d - __obf_1d9e78caa969d04e*1000
	__obf_df8f035bbd062d14 := __obf_1d9e78caa969d04e / 1000
	if __obf_df8f035bbd062d14 == 0 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_1d9e78caa969d04e])
	} else {
		__obf_c2b79e0e6e6f08b5 := __obf_1d9e78caa969d04e - __obf_df8f035bbd062d14*1000
		__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, byte(__obf_df8f035bbd062d14+'0'))
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_c2b79e0e6e6f08b5])
	}
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_c3f7781dafa5b032])
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237,

	// WriteInt32 write int32 to stream
	__obf_209c14c49f9e1dfb[__obf_bd67c2557d95db9e])
}

func (__obf_850a7457bb739a32 *Stream) WriteInt32(__obf_993dfc2f74defa46 int32) {
	var __obf_bbfd2ac8618a6f0c uint32
	if __obf_993dfc2f74defa46 < 0 {
		__obf_bbfd2ac8618a6f0c = uint32(-__obf_993dfc2f74defa46)
		__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, '-')
	} else {
		__obf_bbfd2ac8618a6f0c = uint32(__obf_993dfc2f74defa46)
	}
	__obf_850a7457bb739a32.
		WriteUint32(__obf_bbfd2ac8618a6f0c)
}

// WriteUint64 write uint64 to stream
func (__obf_850a7457bb739a32 *Stream) WriteUint64(__obf_bbfd2ac8618a6f0c uint64) {
	__obf_73c1195b655f767d := __obf_bbfd2ac8618a6f0c / 1000
	if __obf_73c1195b655f767d == 0 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_bbfd2ac8618a6f0c])
		return
	}
	__obf_bd67c2557d95db9e := __obf_bbfd2ac8618a6f0c - __obf_73c1195b655f767d*1000
	__obf_1d9e78caa969d04e := __obf_73c1195b655f767d / 1000
	if __obf_1d9e78caa969d04e == 0 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_73c1195b655f767d])
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_bd67c2557d95db9e])
		return
	}
	__obf_c3f7781dafa5b032 := __obf_73c1195b655f767d - __obf_1d9e78caa969d04e*1000
	__obf_df8f035bbd062d14 := __obf_1d9e78caa969d04e / 1000
	if __obf_df8f035bbd062d14 == 0 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_1d9e78caa969d04e])
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_c3f7781dafa5b032])
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_bd67c2557d95db9e])
		return
	}
	__obf_c2b79e0e6e6f08b5 := __obf_1d9e78caa969d04e - __obf_df8f035bbd062d14*1000
	__obf_63c29fc2df70bf68 := __obf_df8f035bbd062d14 / 1000
	if __obf_63c29fc2df70bf68 == 0 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_df8f035bbd062d14])
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_c2b79e0e6e6f08b5])
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_c3f7781dafa5b032])
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_bd67c2557d95db9e])
		return
	}
	__obf_28adab95c01a3b21 := __obf_df8f035bbd062d14 - __obf_63c29fc2df70bf68*1000
	__obf_b44e89116ba9abe2 := __obf_63c29fc2df70bf68 / 1000
	if __obf_b44e89116ba9abe2 == 0 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_63c29fc2df70bf68])
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_28adab95c01a3b21])
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_c2b79e0e6e6f08b5])
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_c3f7781dafa5b032])
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_bd67c2557d95db9e])
		return
	}
	__obf_c144e1a73562607f := __obf_63c29fc2df70bf68 - __obf_b44e89116ba9abe2*1000
	__obf_c40df0d651eb33f6 := __obf_b44e89116ba9abe2 / 1000
	if __obf_c40df0d651eb33f6 == 0 {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_b44e89116ba9abe2])
	} else {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_fd59d706acca82ca(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_c40df0d651eb33f6])
		__obf_18436ec4bf2fa8d4 := __obf_b44e89116ba9abe2 - __obf_c40df0d651eb33f6*1000
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_18436ec4bf2fa8d4])
	}
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_c144e1a73562607f])
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_28adab95c01a3b21])
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_c2b79e0e6e6f08b5])
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_209c14c49f9e1dfb[__obf_c3f7781dafa5b032])
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_56cc32e8e17354a8(__obf_850a7457bb739a32.__obf_184433571fa55237,

	// WriteInt64 write int64 to stream
	__obf_209c14c49f9e1dfb[__obf_bd67c2557d95db9e])
}

func (__obf_850a7457bb739a32 *Stream) WriteInt64(__obf_993dfc2f74defa46 int64) {
	var __obf_bbfd2ac8618a6f0c uint64
	if __obf_993dfc2f74defa46 < 0 {
		__obf_bbfd2ac8618a6f0c = uint64(-__obf_993dfc2f74defa46)
		__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, '-')
	} else {
		__obf_bbfd2ac8618a6f0c = uint64(__obf_993dfc2f74defa46)
	}
	__obf_850a7457bb739a32.
		WriteUint64(__obf_bbfd2ac8618a6f0c)
}

// WriteInt write int to stream
func (__obf_850a7457bb739a32 *Stream) WriteInt(__obf_bbfd2ac8618a6f0c int) {
	__obf_850a7457bb739a32.
		WriteInt64(int64(__obf_bbfd2ac8618a6f0c))
}

// WriteUint write uint to stream
func (__obf_850a7457bb739a32 *Stream) WriteUint(__obf_bbfd2ac8618a6f0c uint) {
	__obf_850a7457bb739a32.
		WriteUint64(uint64(__obf_bbfd2ac8618a6f0c))
}
