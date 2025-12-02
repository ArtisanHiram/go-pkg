package __obf_a935eb7f548271a4

var (
	PosFixedNumHigh byte = 0x7f
	NegFixedNumLow  byte = 0xe0

	Nil byte = 0xc0

	False byte = 0xc2
	True  byte = 0xc3

	Float  byte = 0xca
	Double byte = 0xcb

	Uint8  byte = 0xcc
	Uint16 byte = 0xcd
	Uint32 byte = 0xce
	Uint64 byte = 0xcf

	Int8  byte = 0xd0
	Int16 byte = 0xd1
	Int32 byte = 0xd2
	Int64 byte = 0xd3

	FixedStrLow  byte = 0xa0
	FixedStrHigh byte = 0xbf
	FixedStrMask byte = 0x1f
	Str8         byte = 0xd9
	Str16        byte = 0xda
	Str32        byte = 0xdb

	Bin8  byte = 0xc4
	Bin16 byte = 0xc5
	Bin32 byte = 0xc6

	FixedArrayLow  byte = 0x90
	FixedArrayHigh byte = 0x9f
	FixedArrayMask byte = 0xf
	Array16        byte = 0xdc
	Array32        byte = 0xdd

	FixedMapLow  byte = 0x80
	FixedMapHigh byte = 0x8f
	FixedMapMask byte = 0xf
	Map16        byte = 0xde
	Map32        byte = 0xdf

	FixExt1  byte = 0xd4
	FixExt2  byte = 0xd5
	FixExt4  byte = 0xd6
	FixExt8  byte = 0xd7
	FixExt16 byte = 0xd8
	Ext8     byte = 0xc7
	Ext16    byte = 0xc8
	Ext32    byte = 0xc9
)

func IsFixedNum(__obf_f5df560f4d67421b byte) bool {
	return __obf_f5df560f4d67421b <= PosFixedNumHigh || __obf_f5df560f4d67421b >= NegFixedNumLow
}

func IsFixedMap(__obf_f5df560f4d67421b byte) bool {
	return __obf_f5df560f4d67421b >= FixedMapLow && __obf_f5df560f4d67421b <= FixedMapHigh
}

func IsFixedArray(__obf_f5df560f4d67421b byte) bool {
	return __obf_f5df560f4d67421b >= FixedArrayLow && __obf_f5df560f4d67421b <= FixedArrayHigh
}

func IsFixedString(__obf_f5df560f4d67421b byte) bool {
	return __obf_f5df560f4d67421b >= FixedStrLow && __obf_f5df560f4d67421b <= FixedStrHigh
}

func IsString(__obf_f5df560f4d67421b byte) bool {
	return IsFixedString(__obf_f5df560f4d67421b) || __obf_f5df560f4d67421b == Str8 || __obf_f5df560f4d67421b == Str16 || __obf_f5df560f4d67421b == Str32
}

func IsBin(__obf_f5df560f4d67421b byte) bool {
	return __obf_f5df560f4d67421b == Bin8 || __obf_f5df560f4d67421b == Bin16 || __obf_f5df560f4d67421b == Bin32
}

func IsFixedExt(__obf_f5df560f4d67421b byte) bool {
	return __obf_f5df560f4d67421b >= FixExt1 && __obf_f5df560f4d67421b <= FixExt16
}

func IsExt(__obf_f5df560f4d67421b byte) bool {
	return IsFixedExt(__obf_f5df560f4d67421b) || __obf_f5df560f4d67421b == Ext8 || __obf_f5df560f4d67421b == Ext16 || __obf_f5df560f4d67421b == Ext32
}
