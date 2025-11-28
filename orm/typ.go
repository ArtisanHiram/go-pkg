package __obf_df9e37b4dd16fa57

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
func (__obf_2eac47834e11c59d *Str) Scan(__obf_b6d2b72f1fbb20ab any) error {
	if __obf_b6d2b72f1fbb20ab == nil {
		return nil
	}

	switch __obf_fb7e0dccfd572f81 := __obf_b6d2b72f1fbb20ab.(type) {
	case []byte:
		*__obf_2eac47834e11c59d = Str(__obf_fb7e0dccfd572f81)
	case string:
		*__obf_2eac47834e11c59d = Str(__obf_fb7e0dccfd572f81)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_b6d2b72f1fbb20ab)
	}
	return nil
}

func (__obf_2eac47834e11c59d *Str) Str() string {
	if __obf_2eac47834e11c59d == nil {
		return ""
	}
	return string(*__obf_2eac47834e11c59d)
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

func (__obf_2a7388295d503598 *Int) Scan(__obf_b6d2b72f1fbb20ab any) error {
	if __obf_b6d2b72f1fbb20ab == nil {
		return nil
	}

	switch __obf_fb7e0dccfd572f81 := __obf_b6d2b72f1fbb20ab.(type) {
	case []byte:
		__obf_a0400a0a546202a1, _ := strconv.ParseInt(string(__obf_fb7e0dccfd572f81), 10, 64)
		*__obf_2a7388295d503598 = Int(__obf_a0400a0a546202a1)
	case int64:
		*__obf_2a7388295d503598 = Int(__obf_fb7e0dccfd572f81)
	case uint64:
		*__obf_2a7388295d503598 = Int(__obf_fb7e0dccfd572f81)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_b6d2b72f1fbb20ab)
	}
	return nil
}

func (__obf_2a7388295d503598 *Int) Int() int64 {
	if __obf_2a7388295d503598 == nil {
		return 0
	}
	return int64(*__obf_2a7388295d503598)
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

func (__obf_2e7920ea71b44a3b *Float) Scan(__obf_b6d2b72f1fbb20ab any) error {
	if __obf_b6d2b72f1fbb20ab == nil {
		return nil
	}

	switch __obf_fb7e0dccfd572f81 := __obf_b6d2b72f1fbb20ab.(type) {
	case []byte:
		if __obf_2eac47834e11c59d, __obf_4795d167371f16d4 := strconv.ParseFloat(string(__obf_fb7e0dccfd572f81), 64); __obf_4795d167371f16d4 != nil {
			return __obf_4795d167371f16d4
		} else {
			*__obf_2e7920ea71b44a3b = Float(__obf_2eac47834e11c59d)
		}
	case float64:
		*__obf_2e7920ea71b44a3b = Float(__obf_fb7e0dccfd572f81)
	default:
		return fmt.Errorf("unsupported type for: %T", __obf_b6d2b72f1fbb20ab)
	}

	return nil
}

func (__obf_2a7388295d503598 *Float) Float() float64 {
	if __obf_2a7388295d503598 == nil {
		return 0
	}
	return float64(*__obf_2a7388295d503598)
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
func DoScan[T any](__obf_317b24c7665170a2 T, __obf_a0400a0a546202a1 any) error {
	if __obf_a0400a0a546202a1 == nil {
		return nil
	}
	var __obf_04e12d1b5f6c3392 []byte
	switch __obf_a0400a0a546202a1 := __obf_a0400a0a546202a1.(type) {
	case T:
		__obf_317b24c7665170a2 = __obf_a0400a0a546202a1
		return nil
	case string:
		__obf_04e12d1b5f6c3392 = []byte(__obf_a0400a0a546202a1)
	case []byte:
		__obf_04e12d1b5f6c3392 = __obf_a0400a0a546202a1
	default:
		return fmt.Errorf("incompatible type for %T", __obf_a0400a0a546202a1)
	}

	if len(__obf_04e12d1b5f6c3392) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_04e12d1b5f6c3392, &__obf_317b24c7665170a2)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_fb7e0dccfd572f81 *JsonMap) Scan(__obf_fe4dab033b4bbf22 any) error {
	return DoScan(__obf_fb7e0dccfd572f81, __obf_fe4dab033b4bbf22)
}

func (__obf_fb7e0dccfd572f81 JsonMap) Value() (driver.Value, error) {
	if len(__obf_fb7e0dccfd572f81) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_fb7e0dccfd572f81)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_388aede2a4740f62 *StringArray) Scan(__obf_fe4dab033b4bbf22 any) error {
	return DoScan(__obf_388aede2a4740f62, __obf_fe4dab033b4bbf22)

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

	// return json.Unmarshal(bytes, &ja)
}

func (__obf_fb7e0dccfd572f81 StringArray) Value() (driver.Value, error) {
	if len(__obf_fb7e0dccfd572f81) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_fb7e0dccfd572f81)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_fb7e0dccfd572f81 *IntArray) Scan(__obf_fe4dab033b4bbf22 any) error {
	return DoScan(__obf_fb7e0dccfd572f81, __obf_fe4dab033b4bbf22)
}

func (__obf_fb7e0dccfd572f81 IntArray) Value() (driver.Value, error) {
	if len(__obf_fb7e0dccfd572f81) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_fb7e0dccfd572f81)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_fb7e0dccfd572f81 *MapArray) Scan(__obf_fe4dab033b4bbf22 any) error {
	return DoScan(__obf_fb7e0dccfd572f81, __obf_fe4dab033b4bbf22)
}

func (__obf_fb7e0dccfd572f81 MapArray) Value() (driver.Value, error) {
	if len(__obf_fb7e0dccfd572f81) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_fb7e0dccfd572f81)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_fb7e0dccfd572f81 *CaseTask) Scan(__obf_fe4dab033b4bbf22 any) error {
	return DoScan(__obf_fb7e0dccfd572f81, __obf_fe4dab033b4bbf22)
}

func (__obf_fb7e0dccfd572f81 CaseTask) Value() (driver.Value, error) {
	if len(__obf_fb7e0dccfd572f81) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_fb7e0dccfd572f81)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_fb7e0dccfd572f81 *Timeline) Scan(__obf_fe4dab033b4bbf22 any) error {
	return DoScan(__obf_fb7e0dccfd572f81, __obf_fe4dab033b4bbf22)
}

func (__obf_fb7e0dccfd572f81 Timeline) Value() (driver.Value, error) {
	if len(__obf_fb7e0dccfd572f81) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_fb7e0dccfd572f81)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_fb7e0dccfd572f81 Timeline) TimeLeft() int64 {
	__obf_1c7fb5b88bfe52e9 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_9ffd29542ac27b0d := range __obf_fb7e0dccfd572f81 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_1c7fb5b88bfe52e9 >= __obf_9ffd29542ac27b0d.Start && __obf_1c7fb5b88bfe52e9 < __obf_9ffd29542ac27b0d.End {
			return __obf_9ffd29542ac27b0d.End - __obf_1c7fb5b88bfe52e9
		}
	}

	return 0
}

func (__obf_fb7e0dccfd572f81 Timeline) IsUpcoming() bool {
	__obf_1c7fb5b88bfe52e9 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_9ffd29542ac27b0d := range __obf_fb7e0dccfd572f81 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_9ffd29542ac27b0d.End > __obf_1c7fb5b88bfe52e9 {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_11d7a6474cb03a3b RawJson) Message() json.RawMessage {
	if len(__obf_11d7a6474cb03a3b) == 0 {
		return nil
	}
	return json.RawMessage(__obf_11d7a6474cb03a3b)
}

// Scan implements the Scanner interface.
func (__obf_11d7a6474cb03a3b *RawJson) Scan(__obf_fe4dab033b4bbf22 any) error {
	if __obf_fe4dab033b4bbf22 == nil {
		*__obf_11d7a6474cb03a3b = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_3f2b402cab2d09b4, __obf_1b7dc4184c531bcd := __obf_fe4dab033b4bbf22.([]byte)
	if !__obf_1b7dc4184c531bcd {
		// Handle other types if necessary, e.g., string
		__obf_624f6c8fbdafd38a, __obf_1b7dc4184c531bcd := __obf_fe4dab033b4bbf22.(string)
		if __obf_1b7dc4184c531bcd {
			__obf_3f2b402cab2d09b4 = []byte(__obf_624f6c8fbdafd38a)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_fe4dab033b4bbf22)
		}
	}
	*__obf_11d7a6474cb03a3b = RawJson(__obf_3f2b402cab2d09b4)
	return nil
}

// Value implements the Valuer interface.
func (__obf_11d7a6474cb03a3b RawJson) Value() (driver.Value, error) {
	if len(__obf_11d7a6474cb03a3b) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_11d7a6474cb03a3b), nil
}
