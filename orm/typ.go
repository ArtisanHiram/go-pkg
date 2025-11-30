package __obf_86786288bdd26c8f

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Str string

// var nullBytes = []byte("")

// Scan implements the Scanner interface for String
func (__obf_c13a94e56d44ee42 *Str) Scan(__obf_8fa7f78d807b0d0f any) error {
	if __obf_8fa7f78d807b0d0f == nil {
		return nil
	}

	switch __obf_374fb3284e6cd2d2 := __obf_8fa7f78d807b0d0f.(type) {
	case []byte:
		*__obf_c13a94e56d44ee42 = Str(__obf_374fb3284e6cd2d2)
	case string:
		*__obf_c13a94e56d44ee42 = Str(__obf_374fb3284e6cd2d2)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_8fa7f78d807b0d0f)
	}
	return nil
}

func (__obf_c13a94e56d44ee42 *Str) Str() string {
	if __obf_c13a94e56d44ee42 == nil {
		return ""
	}
	return string(*__obf_c13a94e56d44ee42)
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string and null input. Blank string input does not produce a null String.
// func (s *String) UnmarshalJSON(data []byte) error {
// 	if bytes.Equal(data, nullBytes) {
// 		return nil
// 	}

// 	*s = String(data)
// 	return nil
// }

// MarshalJSON implements json.Marshaler.
// It will encode null if this String is null.
// func (s String) MarshalJSON() ([]byte, error) {
// 	if s == "" {
// 		return nullBytes, nil
// 	}
// 	return []byte(string(s)), nil
// 	// return []byte(fmt.Sprintf(`"%s"`, s)), nil
// 	// return json.Marshal(s)
// }

type Int int64

func (__obf_33fdeff971fca5a6 *Int) Scan(__obf_8fa7f78d807b0d0f any) error {
	if __obf_8fa7f78d807b0d0f == nil {
		return nil
	}

	switch __obf_374fb3284e6cd2d2 := __obf_8fa7f78d807b0d0f.(type) {
	case []byte:
		__obf_935e5179094b0bc8, _ := strconv.ParseInt(string(__obf_374fb3284e6cd2d2), 10, 64)
		*__obf_33fdeff971fca5a6 = Int(__obf_935e5179094b0bc8)
	case int64:
		*__obf_33fdeff971fca5a6 = Int(__obf_374fb3284e6cd2d2)
	case uint64:
		*__obf_33fdeff971fca5a6 = Int(__obf_374fb3284e6cd2d2)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_8fa7f78d807b0d0f)
	}
	return nil
}

func (__obf_33fdeff971fca5a6 *Int) Int() int64 {
	if __obf_33fdeff971fca5a6 == nil {
		return 0
	}
	return int64(*__obf_33fdeff971fca5a6)
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports number, string, and null input.
// 0 will not be considered a null Int.
// func (i *Int) UnmarshalJSON(data []byte) error {
// 	if bytes.Equal(data, nullBytes) {
// 		return nil
// 	}

// 	if err := json.Unmarshal(data, &i); err != nil {
// 		var typeError *json.UnmarshalTypeError
// 		if errors.As(err, &typeError) {
// 			// special case: accept string input
// 			if typeError.Value != "string" {
// 				return fmt.Errorf("JSON input is invalid type (need int or string): %w", err)
// 			}
// 			var str string
// 			if err := json.Unmarshal(data, &str); err != nil {
// 				return fmt.Errorf("couldn't unmarshal number string: %w", err)
// 			}
// 			n, err := strconv.ParseInt(str, 10, 64)
// 			if err != nil {
// 				return fmt.Errorf("couldn't convert string to int: %w", err)
// 			}
// 			*i = Int(n)
// 			return nil
// 		}
// 		return fmt.Errorf("couldn't unmarshal JSON: %w", err)
// 	}

// 	return nil
// }

// MarshalJSON implements json.Marshaler.
// It will encode null if this Int is null.
// func (i Int) MarshalJSON() ([]byte, error) {
// 	if i == 0 {
// 		return nullBytes, nil
// 	}
// 	return []byte(strconv.FormatInt(int64(i), 10)), nil
// }

type Float float64

func (__obf_ebada1fffc8c92ff *Float) Scan(__obf_8fa7f78d807b0d0f any) error {
	if __obf_8fa7f78d807b0d0f == nil {
		return nil
	}

	switch __obf_374fb3284e6cd2d2 := __obf_8fa7f78d807b0d0f.(type) {
	case []byte:
		if __obf_c13a94e56d44ee42, __obf_87239f12e0f48308 := strconv.ParseFloat(string(__obf_374fb3284e6cd2d2), 64); __obf_87239f12e0f48308 != nil {
			return __obf_87239f12e0f48308
		} else {
			*__obf_ebada1fffc8c92ff = Float(__obf_c13a94e56d44ee42)
		}
	case float64:
		*__obf_ebada1fffc8c92ff = Float(__obf_374fb3284e6cd2d2)
	default:
		return fmt.Errorf("unsupported type for: %T", __obf_8fa7f78d807b0d0f)
	}

	return nil
}

func (__obf_33fdeff971fca5a6 *Float) Float() float64 {
	if __obf_33fdeff971fca5a6 == nil {
		return 0
	}
	return float64(*__obf_33fdeff971fca5a6)
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will not be considered a null Float.
// func (f *Float) UnmarshalJSON(data []byte) error {
// 	if bytes.Equal(data, nullBytes) {
// 		return nil
// 	}

// 	if err := json.Unmarshal(data, &f); err != nil {
// 		var typeError *json.UnmarshalTypeError
// 		if errors.As(err, &typeError) {
// 			// special case: accept string input
// 			if typeError.Value != "string" {
// 				return fmt.Errorf("JSON input is invalid type (need float or string): %w", err)
// 			}
// 			var str string
// 			if err := json.Unmarshal(data, &str); err != nil {
// 				return fmt.Errorf("couldn't unmarshal number string: %w", err)
// 			}
// 			n, err := strconv.ParseFloat(str, 64)
// 			if err != nil {
// 				return fmt.Errorf("couldn't convert string to float: %w", err)
// 			}
// 			*f = Float(n)
// 			return nil
// 		}
// 		return fmt.Errorf("couldn't unmarshal JSON: %w", err)
// 	}
// 	return nil
// }

// MarshalJSON implements json.Marshaler.
// It will encode null if this Float is null.
// func (f Float) MarshalJSON() ([]byte, error) {

// 	if math.IsInf(float64(f), 0) || math.IsNaN(float64(f)) {
// 		return nil, &json.UnsupportedValueError{
// 			Value: reflect.ValueOf(f),
// 			Str:   strconv.FormatFloat(float64(f), 'g', -1, 64),
// 		}
// 	}
// 	return []byte(strconv.FormatFloat(float64(f), 'f', -1, 64)), nil
// }

type Bool bool

// Scan implements the Scanner interface.
// func (b *Bool) Scan(value any) error {
// 	if value == nil {
// 		*b = false
// 		return nil
// 	}
// 	switch t := value.(type) {
// 	case []byte:
// 		if bl, err := strconv.ParseBool(string(t)); err != nil {
// 			return err
// 		} else {
// 			*b = Bool(bl)
// 		}
// 	case string:
// 		if bl, err := strconv.ParseBool(t); err != nil {
// 			return err
// 		} else {
// 			*b = Bool(bl)
// 		}
// 	default:
// 		return fmt.Errorf("Unsupported type for: %T", value)
// 	}
// 	return nil
// }

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will not be considered a null Bool.
// func (b *Bool) UnmarshalJSON(data []byte) error {
// 	if bytes.Equal(data, nullBytes) {
// 		*b = false
// 		return nil
// 	}

// 	if err := json.Unmarshal(data, &b); err != nil {
// 		return fmt.Errorf("couldn't unmarshal JSON: %w", err)
// 	}
// 	return nil
// }

// MarshalJSON implements json.Marshaler.
// It will encode null if this Bool is null.
// func (b Bool) MarshalJSON() ([]byte, error) {
// 	if !b {
// 		return []byte("false"), nil
// 	}
// 	return []byte("true"), nil
// }

// obj 为指针类型
func DoScan[T any](__obf_f62847cc5b375971 T, __obf_935e5179094b0bc8 any) error {
	if __obf_935e5179094b0bc8 == nil {
		return nil
	}
	var __obf_fa4aef4fcdf73680 []byte
	switch __obf_935e5179094b0bc8 := __obf_935e5179094b0bc8.(type) {
	case T:
		__obf_f62847cc5b375971 = __obf_935e5179094b0bc8
		return nil
	case string:
		__obf_fa4aef4fcdf73680 = []byte(__obf_935e5179094b0bc8)
	case []byte:
		__obf_fa4aef4fcdf73680 = __obf_935e5179094b0bc8
	default:
		return fmt.Errorf("incompatible type for %T", __obf_935e5179094b0bc8)
	}

	if len(__obf_fa4aef4fcdf73680) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_fa4aef4fcdf73680, &__obf_f62847cc5b375971)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_374fb3284e6cd2d2 *JsonMap) Scan(__obf_d131126320e759f6 any) error {
	return DoScan(__obf_374fb3284e6cd2d2, __obf_d131126320e759f6)
}

func (__obf_374fb3284e6cd2d2 JsonMap) Value() (driver.Value, error) {
	if len(__obf_374fb3284e6cd2d2) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_374fb3284e6cd2d2)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_ba4d2184c3bb4f63 *StringArray) Scan(__obf_d131126320e759f6 any) error {
	return DoScan(__obf_ba4d2184c3bb4f63,

		// var bytes []byte
		// switch value := value.(type) {
		// case []uint8:
		// 	if len(value) == 0 {
		// 		return nil
		// 	}
		// 	bytes = value
		// default:
		// 	return fmt.Errorf("StringArray: scan source was not []byte")
		// }
		__obf_d131126320e759f6)

	// return json.Unmarshal(bytes, &ja)
}

func (__obf_374fb3284e6cd2d2 StringArray) Value() (driver.Value, error) {
	if len(__obf_374fb3284e6cd2d2) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_374fb3284e6cd2d2)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_374fb3284e6cd2d2 *IntArray) Scan(__obf_d131126320e759f6 any) error {
	return DoScan(__obf_374fb3284e6cd2d2, __obf_d131126320e759f6)
}

func (__obf_374fb3284e6cd2d2 IntArray) Value() (driver.Value, error) {
	if len(__obf_374fb3284e6cd2d2) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_374fb3284e6cd2d2)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_374fb3284e6cd2d2 *MapArray) Scan(__obf_d131126320e759f6 any) error {
	return DoScan(__obf_374fb3284e6cd2d2, __obf_d131126320e759f6)
}

func (__obf_374fb3284e6cd2d2 MapArray) Value() (driver.Value, error) {
	if len(__obf_374fb3284e6cd2d2) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_374fb3284e6cd2d2)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_374fb3284e6cd2d2 *CaseTask) Scan(__obf_d131126320e759f6 any) error {
	return DoScan(__obf_374fb3284e6cd2d2, __obf_d131126320e759f6)
}

func (__obf_374fb3284e6cd2d2 CaseTask) Value() (driver.Value, error) {
	if len(__obf_374fb3284e6cd2d2) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_374fb3284e6cd2d2)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_374fb3284e6cd2d2 *Timeline) Scan(__obf_d131126320e759f6 any) error {
	return DoScan(__obf_374fb3284e6cd2d2, __obf_d131126320e759f6)
}

func (__obf_374fb3284e6cd2d2 Timeline) Value() (driver.Value, error) {
	if len(__obf_374fb3284e6cd2d2) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_374fb3284e6cd2d2)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_374fb3284e6cd2d2 Timeline) TimeLeft() int64 {
	__obf_262febc9abe2d314 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_a0deb2126f98b25c := range __obf_374fb3284e6cd2d2 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_262febc9abe2d314 >= __obf_a0deb2126f98b25c.Start && __obf_262febc9abe2d314 < __obf_a0deb2126f98b25c.End {
			return __obf_a0deb2126f98b25c.End - __obf_262febc9abe2d314
		}
	}

	return 0
}

func (__obf_374fb3284e6cd2d2 Timeline) IsUpcoming() bool {
	__obf_262febc9abe2d314 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_a0deb2126f98b25c := range __obf_374fb3284e6cd2d2 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_a0deb2126f98b25c.End > __obf_262febc9abe2d314 {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_943530d0ac11f87b RawJson) Message() json.RawMessage {
	if len(__obf_943530d0ac11f87b) == 0 {
		return nil
	}
	return json.RawMessage(__obf_943530d0ac11f87b)
}

// Scan implements the Scanner interface.
func (__obf_943530d0ac11f87b *RawJson) Scan(__obf_d131126320e759f6 any) error {
	if __obf_d131126320e759f6 == nil {
		*__obf_943530d0ac11f87b = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_6a21857d8446d266, __obf_bfedcf8573204366 := __obf_d131126320e759f6.([]byte)
	if !__obf_bfedcf8573204366 {
		__obf_77504f613d36108a,
			// Handle other types if necessary, e.g., string
			__obf_bfedcf8573204366 := __obf_d131126320e759f6.(string)
		if __obf_bfedcf8573204366 {
			__obf_6a21857d8446d266 = []byte(__obf_77504f613d36108a)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_d131126320e759f6)
		}
	}
	*__obf_943530d0ac11f87b = RawJson(__obf_6a21857d8446d266)
	return nil
}

// Value implements the Valuer interface.
func (__obf_943530d0ac11f87b RawJson) Value() (driver.Value, error) {
	if len(__obf_943530d0ac11f87b) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_943530d0ac11f87b), nil
}
