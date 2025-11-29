package __obf_91620b895eeff9ed

import "fmt"

type __obf_a4f671452051c765 struct {
	__obf_58563642f42f4a04
	__obf_62967739bca1d11d error
}

func __obf_227499e666173bc3(__obf_82c50fcbfc9ab432 []any) *__obf_a4f671452051c765 {
	return &__obf_a4f671452051c765{__obf_58563642f42f4a04{}, fmt.Errorf("%v not found", __obf_82c50fcbfc9ab432)}
}

func (any *__obf_a4f671452051c765,) LastError() error {
	return any.__obf_62967739bca1d11d
}

func (any *__obf_a4f671452051c765,) ValueType() ValueType {
	return InvalidValue
}

func (any *__obf_a4f671452051c765,) MustBeValid() Any {
	panic(any.__obf_62967739bca1d11d)
}

func (any *__obf_a4f671452051c765,) ToBool() bool {
	return false
}

func (any *__obf_a4f671452051c765,) ToInt() int {
	return 0
}

func (any *__obf_a4f671452051c765,) ToInt32() int32 {
	return 0
}

func (any *__obf_a4f671452051c765,) ToInt64() int64 {
	return 0
}

func (any *__obf_a4f671452051c765,) ToUint() uint {
	return 0
}

func (any *__obf_a4f671452051c765,) ToUint32() uint32 {
	return 0
}

func (any *__obf_a4f671452051c765,) ToUint64() uint64 {
	return 0
}

func (any *__obf_a4f671452051c765,) ToFloat32() float32 {
	return 0
}

func (any *__obf_a4f671452051c765,) ToFloat64() float64 {
	return 0
}

func (any *__obf_a4f671452051c765,) ToString() string {
	return ""
}

func (any *__obf_a4f671452051c765,) WriteTo(__obf_850a7457bb739a32 *Stream) {
}

func (any *__obf_a4f671452051c765,) Get(__obf_82c50fcbfc9ab432 ...any) Any {
	if any.__obf_62967739bca1d11d == nil {
		return &__obf_a4f671452051c765{__obf_58563642f42f4a04{}, fmt.Errorf("get %v from invalid", __obf_82c50fcbfc9ab432)}
	}
	return &__obf_a4f671452051c765{__obf_58563642f42f4a04{}, fmt.Errorf("%v, get %v from invalid", any.__obf_62967739bca1d11d, __obf_82c50fcbfc9ab432)}
}

func (any *__obf_a4f671452051c765,) Parse() *Iterator {
	return nil
}

func (any *__obf_a4f671452051c765,) GetInterface() any {
	return nil
}
