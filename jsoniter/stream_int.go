package __obf_030d39f7a8de96c6

var __obf_464ebeb7353addf6 []uint32

func init() {
	__obf_464ebeb7353addf6 = make([]uint32, 1000)
	for __obf_82c6e05b9d226c58 := uint32(0); __obf_82c6e05b9d226c58 < 1000; __obf_82c6e05b9d226c58++ {
		__obf_464ebeb7353addf6[__obf_82c6e05b9d226c58] = (((__obf_82c6e05b9d226c58 / 100) + '0') << 16) + ((((__obf_82c6e05b9d226c58 / 10) % 10) + '0') << 8) + __obf_82c6e05b9d226c58%10 + '0'
		if __obf_82c6e05b9d226c58 < 10 {
			__obf_464ebeb7353addf6[__obf_82c6e05b9d226c58] += 2 << 24
		} else if __obf_82c6e05b9d226c58 < 100 {
			__obf_464ebeb7353addf6[__obf_82c6e05b9d226c58] += 1 << 24
		}
	}
}

func __obf_1f72d91ae95720f7(__obf_b65253d54f46c9b6 []byte, __obf_41b184145cfea25b uint32) []byte {
	__obf_447a3399d55be0d6 := __obf_41b184145cfea25b >> 24
	switch __obf_447a3399d55be0d6 {
	case 0:
		__obf_b65253d54f46c9b6 = append(__obf_b65253d54f46c9b6, byte(__obf_41b184145cfea25b>>16), byte(__obf_41b184145cfea25b>>8))
	case 1:
		__obf_b65253d54f46c9b6 = append(__obf_b65253d54f46c9b6, byte(__obf_41b184145cfea25b>>8))
	}
	__obf_b65253d54f46c9b6 = append(__obf_b65253d54f46c9b6, byte(__obf_41b184145cfea25b))
	return __obf_b65253d54f46c9b6
}

func __obf_243624d0b4594522(__obf_0b1656d7ef4bd9df []byte, __obf_41b184145cfea25b uint32) []byte {
	return append(__obf_0b1656d7ef4bd9df, byte(__obf_41b184145cfea25b>>16), byte(__obf_41b184145cfea25b>>8), byte(__obf_41b184145cfea25b))
}

// WriteUint8 write uint8 to stream
func (__obf_8a2c51fe22775655 *Stream) WriteUint8(__obf_efaf2719b44ad8ac uint8) {
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df,

	// WriteInt8 write int8 to stream
	__obf_464ebeb7353addf6[__obf_efaf2719b44ad8ac])
}

func (__obf_8a2c51fe22775655 *Stream) WriteInt8(__obf_259f5ba48cea8307 int8) {
	var __obf_efaf2719b44ad8ac uint8
	if __obf_259f5ba48cea8307 < 0 {
		__obf_efaf2719b44ad8ac = uint8(-__obf_259f5ba48cea8307)
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, '-')
	} else {
		__obf_efaf2719b44ad8ac = uint8(__obf_259f5ba48cea8307)
	}
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df,

	// WriteUint16 write uint16 to stream
	__obf_464ebeb7353addf6[__obf_efaf2719b44ad8ac])
}

func (__obf_8a2c51fe22775655 *Stream) WriteUint16(__obf_efaf2719b44ad8ac uint16) {
	__obf_287f7db6eeb4df8e := __obf_efaf2719b44ad8ac / 1000
	if __obf_287f7db6eeb4df8e == 0 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_efaf2719b44ad8ac])
		return
	}
	__obf_61ea5bd6abf20460 := __obf_efaf2719b44ad8ac - __obf_287f7db6eeb4df8e*1000
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_287f7db6eeb4df8e])
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df,

	// WriteInt16 write int16 to stream
	__obf_464ebeb7353addf6[__obf_61ea5bd6abf20460])
}

func (__obf_8a2c51fe22775655 *Stream) WriteInt16(__obf_259f5ba48cea8307 int16) {
	var __obf_efaf2719b44ad8ac uint16
	if __obf_259f5ba48cea8307 < 0 {
		__obf_efaf2719b44ad8ac = uint16(-__obf_259f5ba48cea8307)
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, '-')
	} else {
		__obf_efaf2719b44ad8ac = uint16(__obf_259f5ba48cea8307)
	}
	__obf_8a2c51fe22775655.
		WriteUint16(__obf_efaf2719b44ad8ac)
}

// WriteUint32 write uint32 to stream
func (__obf_8a2c51fe22775655 *Stream) WriteUint32(__obf_efaf2719b44ad8ac uint32) {
	__obf_287f7db6eeb4df8e := __obf_efaf2719b44ad8ac / 1000
	if __obf_287f7db6eeb4df8e == 0 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_efaf2719b44ad8ac])
		return
	}
	__obf_61ea5bd6abf20460 := __obf_efaf2719b44ad8ac - __obf_287f7db6eeb4df8e*1000
	__obf_a0159114a8d19416 := __obf_287f7db6eeb4df8e / 1000
	if __obf_a0159114a8d19416 == 0 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_287f7db6eeb4df8e])
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_61ea5bd6abf20460])
		return
	}
	__obf_12837937a8e321f5 := __obf_287f7db6eeb4df8e - __obf_a0159114a8d19416*1000
	__obf_f0ef2ed2641d5f86 := __obf_a0159114a8d19416 / 1000
	if __obf_f0ef2ed2641d5f86 == 0 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_a0159114a8d19416])
	} else {
		__obf_80b1fe72791a634f := __obf_a0159114a8d19416 - __obf_f0ef2ed2641d5f86*1000
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, byte(__obf_f0ef2ed2641d5f86+'0'))
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_80b1fe72791a634f])
	}
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_12837937a8e321f5])
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df,

	// WriteInt32 write int32 to stream
	__obf_464ebeb7353addf6[__obf_61ea5bd6abf20460])
}

func (__obf_8a2c51fe22775655 *Stream) WriteInt32(__obf_259f5ba48cea8307 int32) {
	var __obf_efaf2719b44ad8ac uint32
	if __obf_259f5ba48cea8307 < 0 {
		__obf_efaf2719b44ad8ac = uint32(-__obf_259f5ba48cea8307)
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, '-')
	} else {
		__obf_efaf2719b44ad8ac = uint32(__obf_259f5ba48cea8307)
	}
	__obf_8a2c51fe22775655.
		WriteUint32(__obf_efaf2719b44ad8ac)
}

// WriteUint64 write uint64 to stream
func (__obf_8a2c51fe22775655 *Stream) WriteUint64(__obf_efaf2719b44ad8ac uint64) {
	__obf_287f7db6eeb4df8e := __obf_efaf2719b44ad8ac / 1000
	if __obf_287f7db6eeb4df8e == 0 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_efaf2719b44ad8ac])
		return
	}
	__obf_61ea5bd6abf20460 := __obf_efaf2719b44ad8ac - __obf_287f7db6eeb4df8e*1000
	__obf_a0159114a8d19416 := __obf_287f7db6eeb4df8e / 1000
	if __obf_a0159114a8d19416 == 0 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_287f7db6eeb4df8e])
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_61ea5bd6abf20460])
		return
	}
	__obf_12837937a8e321f5 := __obf_287f7db6eeb4df8e - __obf_a0159114a8d19416*1000
	__obf_f0ef2ed2641d5f86 := __obf_a0159114a8d19416 / 1000
	if __obf_f0ef2ed2641d5f86 == 0 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_a0159114a8d19416])
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_12837937a8e321f5])
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_61ea5bd6abf20460])
		return
	}
	__obf_80b1fe72791a634f := __obf_a0159114a8d19416 - __obf_f0ef2ed2641d5f86*1000
	__obf_a69d10b5d056a772 := __obf_f0ef2ed2641d5f86 / 1000
	if __obf_a69d10b5d056a772 == 0 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_f0ef2ed2641d5f86])
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_80b1fe72791a634f])
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_12837937a8e321f5])
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_61ea5bd6abf20460])
		return
	}
	__obf_753a386e6d3b45f5 := __obf_f0ef2ed2641d5f86 - __obf_a69d10b5d056a772*1000
	__obf_8148908e1985bfc5 := __obf_a69d10b5d056a772 / 1000
	if __obf_8148908e1985bfc5 == 0 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_a69d10b5d056a772])
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_753a386e6d3b45f5])
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_80b1fe72791a634f])
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_12837937a8e321f5])
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_61ea5bd6abf20460])
		return
	}
	__obf_79e21e051f66c60d := __obf_a69d10b5d056a772 - __obf_8148908e1985bfc5*1000
	__obf_602bd55c8bf90c2d := __obf_8148908e1985bfc5 / 1000
	if __obf_602bd55c8bf90c2d == 0 {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_8148908e1985bfc5])
	} else {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_1f72d91ae95720f7(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_602bd55c8bf90c2d])
		__obf_3182067e2ec7ab18 := __obf_8148908e1985bfc5 - __obf_602bd55c8bf90c2d*1000
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_3182067e2ec7ab18])
	}
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_79e21e051f66c60d])
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_753a386e6d3b45f5])
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_80b1fe72791a634f])
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_464ebeb7353addf6[__obf_12837937a8e321f5])
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_243624d0b4594522(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df,

	// WriteInt64 write int64 to stream
	__obf_464ebeb7353addf6[__obf_61ea5bd6abf20460])
}

func (__obf_8a2c51fe22775655 *Stream) WriteInt64(__obf_259f5ba48cea8307 int64) {
	var __obf_efaf2719b44ad8ac uint64
	if __obf_259f5ba48cea8307 < 0 {
		__obf_efaf2719b44ad8ac = uint64(-__obf_259f5ba48cea8307)
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, '-')
	} else {
		__obf_efaf2719b44ad8ac = uint64(__obf_259f5ba48cea8307)
	}
	__obf_8a2c51fe22775655.
		WriteUint64(__obf_efaf2719b44ad8ac)
}

// WriteInt write int to stream
func (__obf_8a2c51fe22775655 *Stream) WriteInt(__obf_efaf2719b44ad8ac int) {
	__obf_8a2c51fe22775655.
		WriteInt64(int64(__obf_efaf2719b44ad8ac))
}

// WriteUint write uint to stream
func (__obf_8a2c51fe22775655 *Stream) WriteUint(__obf_efaf2719b44ad8ac uint) {
	__obf_8a2c51fe22775655.
		WriteUint64(uint64(__obf_efaf2719b44ad8ac))
}
