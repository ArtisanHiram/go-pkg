package __obf_c7b79b12b33d8238

import (
	"io"
	"unsafe"
)

type __obf_7ab06fd17433b644 struct {
	__obf_3a44d9c5204c2e12
	__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0
	__obf_684d738bc3ac851a []byte
	__obf_ea0680f8b461a85b error
}

func (any *__obf_7ab06fd17433b644) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_7ab06fd17433b644) MustBeValid() Any {
	return any
}

func (any *__obf_7ab06fd17433b644) LastError() error {
	return any.__obf_ea0680f8b461a85b
}

func (any *__obf_7ab06fd17433b644) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_7ab06fd17433b644) ToInt() int {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.ReadInt()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		any.__obf_ea0680f8b461a85b = __obf_0da8c843dad13139.Error
	}
	return __obf_35accd096612ebe4
}

func (any *__obf_7ab06fd17433b644) ToInt32() int32 {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.ReadInt32()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		any.__obf_ea0680f8b461a85b = __obf_0da8c843dad13139.Error
	}
	return __obf_35accd096612ebe4
}

func (any *__obf_7ab06fd17433b644) ToInt64() int64 {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.ReadInt64()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		any.__obf_ea0680f8b461a85b = __obf_0da8c843dad13139.Error
	}
	return __obf_35accd096612ebe4
}

func (any *__obf_7ab06fd17433b644) ToUint() uint {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.ReadUint()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		any.__obf_ea0680f8b461a85b = __obf_0da8c843dad13139.Error
	}
	return __obf_35accd096612ebe4
}

func (any *__obf_7ab06fd17433b644) ToUint32() uint32 {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.ReadUint32()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		any.__obf_ea0680f8b461a85b = __obf_0da8c843dad13139.Error
	}
	return __obf_35accd096612ebe4
}

func (any *__obf_7ab06fd17433b644) ToUint64() uint64 {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.ReadUint64()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		any.__obf_ea0680f8b461a85b = __obf_0da8c843dad13139.Error
	}
	return __obf_35accd096612ebe4
}

func (any *__obf_7ab06fd17433b644) ToFloat32() float32 {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.ReadFloat32()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		any.__obf_ea0680f8b461a85b = __obf_0da8c843dad13139.Error
	}
	return __obf_35accd096612ebe4
}

func (any *__obf_7ab06fd17433b644) ToFloat64() float64 {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_35accd096612ebe4 := __obf_0da8c843dad13139.ReadFloat64()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		any.__obf_ea0680f8b461a85b = __obf_0da8c843dad13139.Error
	}
	return __obf_35accd096612ebe4
}

func (any *__obf_7ab06fd17433b644) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_684d738bc3ac851a))
}

func (any *__obf_7ab06fd17433b644) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		Write(any.__obf_684d738bc3ac851a)
}

func (any *__obf_7ab06fd17433b644) GetInterface() any {
	__obf_0da8c843dad13139 := any.__obf_c52e0bcfad4b8b71.BorrowIterator(any.__obf_684d738bc3ac851a)
	defer any.__obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	return __obf_0da8c843dad13139.Read()
}
