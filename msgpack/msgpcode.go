package __obf_a4aad98aaf3178f4

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

func IsFixedNum(__obf_6dbe86b3d9d9b037 byte) bool {
	return __obf_6dbe86b3d9d9b037 <= PosFixedNumHigh || __obf_6dbe86b3d9d9b037 >= NegFixedNumLow
}

func IsFixedMap(__obf_6dbe86b3d9d9b037 byte) bool {
	return __obf_6dbe86b3d9d9b037 >= FixedMapLow && __obf_6dbe86b3d9d9b037 <= FixedMapHigh
}

func IsFixedArray(__obf_6dbe86b3d9d9b037 byte) bool {
	return __obf_6dbe86b3d9d9b037 >= FixedArrayLow && __obf_6dbe86b3d9d9b037 <= FixedArrayHigh
}

func IsFixedString(__obf_6dbe86b3d9d9b037 byte) bool {
	return __obf_6dbe86b3d9d9b037 >= FixedStrLow && __obf_6dbe86b3d9d9b037 <= FixedStrHigh
}

func IsString(__obf_6dbe86b3d9d9b037 byte) bool {
	return IsFixedString(__obf_6dbe86b3d9d9b037) || __obf_6dbe86b3d9d9b037 == Str8 || __obf_6dbe86b3d9d9b037 == Str16 || __obf_6dbe86b3d9d9b037 == Str32
}

func IsBin(__obf_6dbe86b3d9d9b037 byte) bool {
	return __obf_6dbe86b3d9d9b037 == Bin8 || __obf_6dbe86b3d9d9b037 == Bin16 || __obf_6dbe86b3d9d9b037 == Bin32
}

func IsFixedExt(__obf_6dbe86b3d9d9b037 byte) bool {
	return __obf_6dbe86b3d9d9b037 >= FixExt1 && __obf_6dbe86b3d9d9b037 <= FixExt16
}

func IsExt(__obf_6dbe86b3d9d9b037 byte) bool {
	return IsFixedExt(__obf_6dbe86b3d9d9b037) || __obf_6dbe86b3d9d9b037 == Ext8 || __obf_6dbe86b3d9d9b037 == Ext16 || __obf_6dbe86b3d9d9b037 == Ext32
}
