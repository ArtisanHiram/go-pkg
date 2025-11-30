package __obf_3e0c303610a19bd4

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

func IsFixedNum(__obf_e46289218af709bf byte) bool {
	return __obf_e46289218af709bf <= PosFixedNumHigh || __obf_e46289218af709bf >= NegFixedNumLow
}

func IsFixedMap(__obf_e46289218af709bf byte) bool {
	return __obf_e46289218af709bf >= FixedMapLow && __obf_e46289218af709bf <= FixedMapHigh
}

func IsFixedArray(__obf_e46289218af709bf byte) bool {
	return __obf_e46289218af709bf >= FixedArrayLow && __obf_e46289218af709bf <= FixedArrayHigh
}

func IsFixedString(__obf_e46289218af709bf byte) bool {
	return __obf_e46289218af709bf >= FixedStrLow && __obf_e46289218af709bf <= FixedStrHigh
}

func IsString(__obf_e46289218af709bf byte) bool {
	return IsFixedString(__obf_e46289218af709bf) || __obf_e46289218af709bf == Str8 || __obf_e46289218af709bf == Str16 || __obf_e46289218af709bf == Str32
}

func IsBin(__obf_e46289218af709bf byte) bool {
	return __obf_e46289218af709bf == Bin8 || __obf_e46289218af709bf == Bin16 || __obf_e46289218af709bf == Bin32
}

func IsFixedExt(__obf_e46289218af709bf byte) bool {
	return __obf_e46289218af709bf >= FixExt1 && __obf_e46289218af709bf <= FixExt16
}

func IsExt(__obf_e46289218af709bf byte) bool {
	return IsFixedExt(__obf_e46289218af709bf) || __obf_e46289218af709bf == Ext8 || __obf_e46289218af709bf == Ext16 || __obf_e46289218af709bf == Ext32
}
