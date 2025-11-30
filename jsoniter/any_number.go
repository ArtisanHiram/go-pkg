package __obf_703d23ba520c3295

import (
	"io"
	"unsafe"
)

type __obf_2444f00b77eded26 struct {
	__obf_b2a2e31336581ab8
	__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78
	__obf_a065f8e0da5f5952 []byte
	__obf_e56742d6e52f642e error
}

func (any *__obf_2444f00b77eded26) ValueType() ValueType {
	return NumberValue
}

func (any *__obf_2444f00b77eded26) MustBeValid() Any {
	return any
}

func (any *__obf_2444f00b77eded26) LastError() error {
	return any.__obf_e56742d6e52f642e
}

func (any *__obf_2444f00b77eded26) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *__obf_2444f00b77eded26) ToInt() int {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.ReadInt()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		any.__obf_e56742d6e52f642e = __obf_47edab4c16a2d88d.Error
	}
	return __obf_b924538fffe5fd64
}

func (any *__obf_2444f00b77eded26) ToInt32() int32 {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.ReadInt32()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		any.__obf_e56742d6e52f642e = __obf_47edab4c16a2d88d.Error
	}
	return __obf_b924538fffe5fd64
}

func (any *__obf_2444f00b77eded26) ToInt64() int64 {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.ReadInt64()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		any.__obf_e56742d6e52f642e = __obf_47edab4c16a2d88d.Error
	}
	return __obf_b924538fffe5fd64
}

func (any *__obf_2444f00b77eded26) ToUint() uint {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.ReadUint()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		any.__obf_e56742d6e52f642e = __obf_47edab4c16a2d88d.Error
	}
	return __obf_b924538fffe5fd64
}

func (any *__obf_2444f00b77eded26) ToUint32() uint32 {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.ReadUint32()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		any.__obf_e56742d6e52f642e = __obf_47edab4c16a2d88d.Error
	}
	return __obf_b924538fffe5fd64
}

func (any *__obf_2444f00b77eded26) ToUint64() uint64 {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.ReadUint64()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		any.__obf_e56742d6e52f642e = __obf_47edab4c16a2d88d.Error
	}
	return __obf_b924538fffe5fd64
}

func (any *__obf_2444f00b77eded26) ToFloat32() float32 {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.ReadFloat32()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		any.__obf_e56742d6e52f642e = __obf_47edab4c16a2d88d.Error
	}
	return __obf_b924538fffe5fd64
}

func (any *__obf_2444f00b77eded26) ToFloat64() float64 {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_b924538fffe5fd64 := __obf_47edab4c16a2d88d.ReadFloat64()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		any.__obf_e56742d6e52f642e = __obf_47edab4c16a2d88d.Error
	}
	return __obf_b924538fffe5fd64
}

func (any *__obf_2444f00b77eded26) ToString() string {
	return *(*string)(unsafe.Pointer(&any.__obf_a065f8e0da5f5952))
}

func (any *__obf_2444f00b77eded26) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		Write(any.__obf_a065f8e0da5f5952)
}

func (any *__obf_2444f00b77eded26) GetInterface() any {
	__obf_47edab4c16a2d88d := any.__obf_25bd4f754a37b862.BorrowIterator(any.__obf_a065f8e0da5f5952)
	defer any.__obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	return __obf_47edab4c16a2d88d.Read()
}
