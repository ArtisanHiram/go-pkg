package __obf_030d39f7a8de96c6

import (
	"io"
	"unsafe"
)

type __obf_0796d07d4885e149 struct {
	__obf_7834a13a259af712
	__obf_679611dc56ff465b *__obf_81639538813814ff
	__obf_0b1656d7ef4bd9df []byte
	__obf_fcc907dd69879566 error
}

func (any *__obf_0796d07d4885e149) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_0796d07d4885e149) MustBeValid() Any {
	return any
}

func (any *__obf_0796d07d4885e149) LastError() error {
	return any.__obf_fcc907dd69879566
}

func (any *__obf_0796d07d4885e149) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_0796d07d4885e149) ToInt() int {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.ReadInt()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		any.__obf_fcc907dd69879566 = __obf_4ab56a99965952e7.Error
	}
	return __obf_efaf2719b44ad8ac
}

func (any *__obf_0796d07d4885e149) ToInt32() int32 {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.ReadInt32()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		any.__obf_fcc907dd69879566 = __obf_4ab56a99965952e7.Error
	}
	return __obf_efaf2719b44ad8ac
}

func (any *__obf_0796d07d4885e149) ToInt64() int64 {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.ReadInt64()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		any.__obf_fcc907dd69879566 = __obf_4ab56a99965952e7.Error
	}
	return __obf_efaf2719b44ad8ac
}

func (any *__obf_0796d07d4885e149) ToUint() uint {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.ReadUint()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		any.__obf_fcc907dd69879566 = __obf_4ab56a99965952e7.Error
	}
	return __obf_efaf2719b44ad8ac
}

func (any *__obf_0796d07d4885e149) ToUint32() uint32 {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.ReadUint32()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		any.__obf_fcc907dd69879566 = __obf_4ab56a99965952e7.Error
	}
	return __obf_efaf2719b44ad8ac
}

func (any *__obf_0796d07d4885e149) ToUint64() uint64 {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.ReadUint64()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		any.__obf_fcc907dd69879566 = __obf_4ab56a99965952e7.Error
	}
	return __obf_efaf2719b44ad8ac
}

func (any *__obf_0796d07d4885e149) ToFloat32() float32 {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.ReadFloat32()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		any.__obf_fcc907dd69879566 = __obf_4ab56a99965952e7.Error
	}
	return __obf_efaf2719b44ad8ac
}

func (any *__obf_0796d07d4885e149) ToFloat64() float64 {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_efaf2719b44ad8ac := __obf_4ab56a99965952e7.ReadFloat64()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		any.__obf_fcc907dd69879566 = __obf_4ab56a99965952e7.Error
	}
	return __obf_efaf2719b44ad8ac
}

func (any *__obf_0796d07d4885e149) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_0b1656d7ef4bd9df))
}

func (any *__obf_0796d07d4885e149) WriteTo(__obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		Write(any.__obf_0b1656d7ef4bd9df)
}

func (any *__obf_0796d07d4885e149) GetInterface() any {
	__obf_4ab56a99965952e7 := any.__obf_679611dc56ff465b.BorrowIterator(any.__obf_0b1656d7ef4bd9df)
	defer any.__obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	return __obf_4ab56a99965952e7.Read()
}
