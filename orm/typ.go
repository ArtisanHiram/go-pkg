package __obf_c1f2c3d265c98f25

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
func (__obf_07574ea869b3a910 *Str) Scan(__obf_31430ff340f1102d any) error {
	if __obf_31430ff340f1102d == nil {
		return nil
	}

	switch __obf_4829520f0a5c1ebb := __obf_31430ff340f1102d.(type) {
	case []byte:
		*__obf_07574ea869b3a910 = Str(__obf_4829520f0a5c1ebb)
	case string:
		*__obf_07574ea869b3a910 = Str(__obf_4829520f0a5c1ebb)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_31430ff340f1102d)
	}
	return nil
}

func (__obf_07574ea869b3a910 *Str) Str() string {
	if __obf_07574ea869b3a910 == nil {
		return ""
	}
	return string(*__obf_07574ea869b3a910)
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

func (__obf_739dac239fb8287e *Int) Scan(__obf_31430ff340f1102d any) error {
	if __obf_31430ff340f1102d == nil {
		return nil
	}

	switch __obf_4829520f0a5c1ebb := __obf_31430ff340f1102d.(type) {
	case []byte:
		__obf_e87e6f05cc491080, _ := strconv.ParseInt(string(__obf_4829520f0a5c1ebb), 10, 64)
		*__obf_739dac239fb8287e = Int(__obf_e87e6f05cc491080)
	case int64:
		*__obf_739dac239fb8287e = Int(__obf_4829520f0a5c1ebb)
	case uint64:
		*__obf_739dac239fb8287e = Int(__obf_4829520f0a5c1ebb)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_31430ff340f1102d)
	}
	return nil
}

func (__obf_739dac239fb8287e *Int) Int() int64 {
	if __obf_739dac239fb8287e == nil {
		return 0
	}
	return int64(*__obf_739dac239fb8287e)
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

func (__obf_717e0741d033c9e0 *Float) Scan(__obf_31430ff340f1102d any) error {
	if __obf_31430ff340f1102d == nil {
		return nil
	}

	switch __obf_4829520f0a5c1ebb := __obf_31430ff340f1102d.(type) {
	case []byte:
		if __obf_07574ea869b3a910, __obf_ab5192d9e0103059 := strconv.ParseFloat(string(__obf_4829520f0a5c1ebb), 64); __obf_ab5192d9e0103059 != nil {
			return __obf_ab5192d9e0103059
		} else {
			*__obf_717e0741d033c9e0 = Float(__obf_07574ea869b3a910)
		}
	case float64:
		*__obf_717e0741d033c9e0 = Float(__obf_4829520f0a5c1ebb)
	default:
		return fmt.Errorf("unsupported type for: %T", __obf_31430ff340f1102d)
	}

	return nil
}

func (__obf_739dac239fb8287e *Float) Float() float64 {
	if __obf_739dac239fb8287e == nil {
		return 0
	}
	return float64(*__obf_739dac239fb8287e)
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
func DoScan[T any](__obf_608e74048025a546 T, __obf_e87e6f05cc491080 any) error {
	if __obf_e87e6f05cc491080 == nil {
		return nil
	}
	var __obf_5d73b38aa9aa107c []byte
	switch __obf_e87e6f05cc491080 := __obf_e87e6f05cc491080.(type) {
	case T:
		__obf_608e74048025a546 = __obf_e87e6f05cc491080
		return nil
	case string:
		__obf_5d73b38aa9aa107c = []byte(__obf_e87e6f05cc491080)
	case []byte:
		__obf_5d73b38aa9aa107c = __obf_e87e6f05cc491080
	default:
		return fmt.Errorf("incompatible type for %T", __obf_e87e6f05cc491080)
	}

	if len(__obf_5d73b38aa9aa107c) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_5d73b38aa9aa107c, &__obf_608e74048025a546)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_4829520f0a5c1ebb *JsonMap) Scan(__obf_1fe888a8765085d6 any) error {
	return DoScan(__obf_4829520f0a5c1ebb, __obf_1fe888a8765085d6)
}

func (__obf_4829520f0a5c1ebb JsonMap) Value() (driver.Value, error) {
	if len(__obf_4829520f0a5c1ebb) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_4829520f0a5c1ebb)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_01146d46b84522cc *StringArray) Scan(__obf_1fe888a8765085d6 any) error {
	return DoScan(__obf_01146d46b84522cc,

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
		__obf_1fe888a8765085d6)

	// return json.Unmarshal(bytes, &ja)
}

func (__obf_4829520f0a5c1ebb StringArray) Value() (driver.Value, error) {
	if len(__obf_4829520f0a5c1ebb) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_4829520f0a5c1ebb)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_4829520f0a5c1ebb *IntArray) Scan(__obf_1fe888a8765085d6 any) error {
	return DoScan(__obf_4829520f0a5c1ebb, __obf_1fe888a8765085d6)
}

func (__obf_4829520f0a5c1ebb IntArray) Value() (driver.Value, error) {
	if len(__obf_4829520f0a5c1ebb) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_4829520f0a5c1ebb)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_4829520f0a5c1ebb *MapArray) Scan(__obf_1fe888a8765085d6 any) error {
	return DoScan(__obf_4829520f0a5c1ebb, __obf_1fe888a8765085d6)
}

func (__obf_4829520f0a5c1ebb MapArray) Value() (driver.Value, error) {
	if len(__obf_4829520f0a5c1ebb) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_4829520f0a5c1ebb)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_4829520f0a5c1ebb *CaseTask) Scan(__obf_1fe888a8765085d6 any) error {
	return DoScan(__obf_4829520f0a5c1ebb, __obf_1fe888a8765085d6)
}

func (__obf_4829520f0a5c1ebb CaseTask) Value() (driver.Value, error) {
	if len(__obf_4829520f0a5c1ebb) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_4829520f0a5c1ebb)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_4829520f0a5c1ebb *Timeline) Scan(__obf_1fe888a8765085d6 any) error {
	return DoScan(__obf_4829520f0a5c1ebb, __obf_1fe888a8765085d6)
}

func (__obf_4829520f0a5c1ebb Timeline) Value() (driver.Value, error) {
	if len(__obf_4829520f0a5c1ebb) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_4829520f0a5c1ebb)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_4829520f0a5c1ebb Timeline) TimeLeft() int64 {
	__obf_0a01708e6f0e2d5d := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_94c2c3cbd44de900 := range __obf_4829520f0a5c1ebb {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_0a01708e6f0e2d5d >= __obf_94c2c3cbd44de900.Start && __obf_0a01708e6f0e2d5d < __obf_94c2c3cbd44de900.End {
			return __obf_94c2c3cbd44de900.End - __obf_0a01708e6f0e2d5d
		}
	}

	return 0
}

func (__obf_4829520f0a5c1ebb Timeline) IsUpcoming() bool {
	__obf_0a01708e6f0e2d5d := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_94c2c3cbd44de900 := range __obf_4829520f0a5c1ebb {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_94c2c3cbd44de900.End > __obf_0a01708e6f0e2d5d {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_7f062a273c75bfd9 RawJson) Message() json.RawMessage {
	if len(__obf_7f062a273c75bfd9) == 0 {
		return nil
	}
	return json.RawMessage(__obf_7f062a273c75bfd9)
}

// Scan implements the Scanner interface.
func (__obf_7f062a273c75bfd9 *RawJson) Scan(__obf_1fe888a8765085d6 any) error {
	if __obf_1fe888a8765085d6 == nil {
		*__obf_7f062a273c75bfd9 = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_b077537eb1b549b3, __obf_0d4343a67e152792 := __obf_1fe888a8765085d6.([]byte)
	if !__obf_0d4343a67e152792 {
		__obf_6ded9ac88ab54720,
			// Handle other types if necessary, e.g., string
			__obf_0d4343a67e152792 := __obf_1fe888a8765085d6.(string)
		if __obf_0d4343a67e152792 {
			__obf_b077537eb1b549b3 = []byte(__obf_6ded9ac88ab54720)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_1fe888a8765085d6)
		}
	}
	*__obf_7f062a273c75bfd9 = RawJson(__obf_b077537eb1b549b3)
	return nil
}

// Value implements the Valuer interface.
func (__obf_7f062a273c75bfd9 RawJson) Value() (driver.Value, error) {
	if len(__obf_7f062a273c75bfd9) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_7f062a273c75bfd9), nil
}
