package __obf_5b802ce8d9ba56d6

import (
	"io"
	"unsafe"
)

type __obf_dff8f3718df12fe4 struct {
	__obf_fb75d4e4562132ae
	__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf
	__obf_9fc06d9180f0daca []byte
	__obf_6d071d48f3b321ad error
}

func (any *__obf_dff8f3718df12fe4) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_dff8f3718df12fe4) MustBeValid() Any {
	return any
}

func (any *__obf_dff8f3718df12fe4) LastError() error {
	return any.__obf_6d071d48f3b321ad
}

func (any *__obf_dff8f3718df12fe4) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_dff8f3718df12fe4) ToInt() int {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.ReadInt()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		any.__obf_6d071d48f3b321ad = __obf_67008a6a9e5ba828.Error
	}
	return __obf_5406b11e33b9d1d3
}

func (any *__obf_dff8f3718df12fe4) ToInt32() int32 {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.ReadInt32()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		any.__obf_6d071d48f3b321ad = __obf_67008a6a9e5ba828.Error
	}
	return __obf_5406b11e33b9d1d3
}

func (any *__obf_dff8f3718df12fe4) ToInt64() int64 {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.ReadInt64()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		any.__obf_6d071d48f3b321ad = __obf_67008a6a9e5ba828.Error
	}
	return __obf_5406b11e33b9d1d3
}

func (any *__obf_dff8f3718df12fe4) ToUint() uint {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.ReadUint()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		any.__obf_6d071d48f3b321ad = __obf_67008a6a9e5ba828.Error
	}
	return __obf_5406b11e33b9d1d3
}

func (any *__obf_dff8f3718df12fe4) ToUint32() uint32 {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.ReadUint32()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		any.__obf_6d071d48f3b321ad = __obf_67008a6a9e5ba828.Error
	}
	return __obf_5406b11e33b9d1d3
}

func (any *__obf_dff8f3718df12fe4) ToUint64() uint64 {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.ReadUint64()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		any.__obf_6d071d48f3b321ad = __obf_67008a6a9e5ba828.Error
	}
	return __obf_5406b11e33b9d1d3
}

func (any *__obf_dff8f3718df12fe4) ToFloat32() float32 {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.ReadFloat32()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		any.__obf_6d071d48f3b321ad = __obf_67008a6a9e5ba828.Error
	}
	return __obf_5406b11e33b9d1d3
}

func (any *__obf_dff8f3718df12fe4) ToFloat64() float64 {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_5406b11e33b9d1d3 := __obf_67008a6a9e5ba828.ReadFloat64()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		any.__obf_6d071d48f3b321ad = __obf_67008a6a9e5ba828.Error
	}
	return __obf_5406b11e33b9d1d3
}

func (any *__obf_dff8f3718df12fe4) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_9fc06d9180f0daca))
}

func (any *__obf_dff8f3718df12fe4) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		Write(any.__obf_9fc06d9180f0daca)
}

func (any *__obf_dff8f3718df12fe4) GetInterface() any {
	__obf_67008a6a9e5ba828 := any.__obf_dca106e1cdf85392.BorrowIterator(any.__obf_9fc06d9180f0daca)
	defer any.__obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	return __obf_67008a6a9e5ba828.Read()
}
