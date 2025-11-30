package __obf_c3cd04a15c56f32f

import (
	"io"
	"unsafe"
)

type __obf_76146a3e4f4d12ac struct {
	__obf_86279f81acb381f3
	__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb
	__obf_ace979f6622823f3 []byte
	__obf_5966ecc5edb9b17e error
}

func (any *__obf_76146a3e4f4d12ac) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_76146a3e4f4d12ac) MustBeValid() Any {
	return any
}

func (any *__obf_76146a3e4f4d12ac) LastError() error {
	return any.__obf_5966ecc5edb9b17e
}

func (any *__obf_76146a3e4f4d12ac) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_76146a3e4f4d12ac) ToInt() int {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.ReadInt()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		any.__obf_5966ecc5edb9b17e = __obf_edd9fe48d39445e4.Error
	}
	return __obf_d59813f23ad740a8
}

func (any *__obf_76146a3e4f4d12ac) ToInt32() int32 {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.ReadInt32()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		any.__obf_5966ecc5edb9b17e = __obf_edd9fe48d39445e4.Error
	}
	return __obf_d59813f23ad740a8
}

func (any *__obf_76146a3e4f4d12ac) ToInt64() int64 {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.ReadInt64()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		any.__obf_5966ecc5edb9b17e = __obf_edd9fe48d39445e4.Error
	}
	return __obf_d59813f23ad740a8
}

func (any *__obf_76146a3e4f4d12ac) ToUint() uint {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.ReadUint()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		any.__obf_5966ecc5edb9b17e = __obf_edd9fe48d39445e4.Error
	}
	return __obf_d59813f23ad740a8
}

func (any *__obf_76146a3e4f4d12ac) ToUint32() uint32 {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.ReadUint32()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		any.__obf_5966ecc5edb9b17e = __obf_edd9fe48d39445e4.Error
	}
	return __obf_d59813f23ad740a8
}

func (any *__obf_76146a3e4f4d12ac) ToUint64() uint64 {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.ReadUint64()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		any.__obf_5966ecc5edb9b17e = __obf_edd9fe48d39445e4.Error
	}
	return __obf_d59813f23ad740a8
}

func (any *__obf_76146a3e4f4d12ac) ToFloat32() float32 {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.ReadFloat32()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		any.__obf_5966ecc5edb9b17e = __obf_edd9fe48d39445e4.Error
	}
	return __obf_d59813f23ad740a8
}

func (any *__obf_76146a3e4f4d12ac) ToFloat64() float64 {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_d59813f23ad740a8 := __obf_edd9fe48d39445e4.ReadFloat64()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		any.__obf_5966ecc5edb9b17e = __obf_edd9fe48d39445e4.Error
	}
	return __obf_d59813f23ad740a8
}

func (any *__obf_76146a3e4f4d12ac) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_ace979f6622823f3))
}

func (any *__obf_76146a3e4f4d12ac) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		Write(any.__obf_ace979f6622823f3)
}

func (any *__obf_76146a3e4f4d12ac) GetInterface() any {
	__obf_edd9fe48d39445e4 := any.__obf_0c8065c219ccf0ab.BorrowIterator(any.__obf_ace979f6622823f3)
	defer any.__obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	return __obf_edd9fe48d39445e4.Read()
}
