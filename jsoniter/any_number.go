package __obf_91620b895eeff9ed

import (
	"io"
	"unsafe"
)

type __obf_37d3c66ca0437bb5 struct {
	__obf_58563642f42f4a04
	__obf_892632c77e859037 *__obf_972c2293804d141c
	__obf_184433571fa55237 []byte
	__obf_62967739bca1d11d error
}

func (any *__obf_37d3c66ca0437bb5) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_37d3c66ca0437bb5) MustBeValid() Any {
	return any
}

func (any *__obf_37d3c66ca0437bb5) LastError() error {
	return any.__obf_62967739bca1d11d
}

func (any *__obf_37d3c66ca0437bb5) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_37d3c66ca0437bb5) ToInt() int {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.ReadInt()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		any.__obf_62967739bca1d11d = __obf_1bb30e8a74ed8233.Error
	}
	return __obf_bbfd2ac8618a6f0c
}

func (any *__obf_37d3c66ca0437bb5) ToInt32() int32 {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.ReadInt32()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		any.__obf_62967739bca1d11d = __obf_1bb30e8a74ed8233.Error
	}
	return __obf_bbfd2ac8618a6f0c
}

func (any *__obf_37d3c66ca0437bb5) ToInt64() int64 {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.ReadInt64()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		any.__obf_62967739bca1d11d = __obf_1bb30e8a74ed8233.Error
	}
	return __obf_bbfd2ac8618a6f0c
}

func (any *__obf_37d3c66ca0437bb5) ToUint() uint {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.ReadUint()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		any.__obf_62967739bca1d11d = __obf_1bb30e8a74ed8233.Error
	}
	return __obf_bbfd2ac8618a6f0c
}

func (any *__obf_37d3c66ca0437bb5) ToUint32() uint32 {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.ReadUint32()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		any.__obf_62967739bca1d11d = __obf_1bb30e8a74ed8233.Error
	}
	return __obf_bbfd2ac8618a6f0c
}

func (any *__obf_37d3c66ca0437bb5) ToUint64() uint64 {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.ReadUint64()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		any.__obf_62967739bca1d11d = __obf_1bb30e8a74ed8233.Error
	}
	return __obf_bbfd2ac8618a6f0c
}

func (any *__obf_37d3c66ca0437bb5) ToFloat32() float32 {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.ReadFloat32()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		any.__obf_62967739bca1d11d = __obf_1bb30e8a74ed8233.Error
	}
	return __obf_bbfd2ac8618a6f0c
}

func (any *__obf_37d3c66ca0437bb5) ToFloat64() float64 {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_bbfd2ac8618a6f0c := __obf_1bb30e8a74ed8233.ReadFloat64()
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		any.__obf_62967739bca1d11d = __obf_1bb30e8a74ed8233.Error
	}
	return __obf_bbfd2ac8618a6f0c
}

func (any *__obf_37d3c66ca0437bb5) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_184433571fa55237))
}

func (any *__obf_37d3c66ca0437bb5) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		Write(any.__obf_184433571fa55237)
}

func (any *__obf_37d3c66ca0437bb5) GetInterface() any {
	__obf_1bb30e8a74ed8233 := any.__obf_892632c77e859037.BorrowIterator(any.__obf_184433571fa55237)
	defer any.__obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	return __obf_1bb30e8a74ed8233.Read()
}
