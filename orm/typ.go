package __obf_e471fc5ea32e9ac0

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
func (__obf_153c6aff24c72d24 *Str) Scan(__obf_3e3889a3ca3aa32a any) error {
	if __obf_3e3889a3ca3aa32a == nil {
		return nil
	}

	switch __obf_82ab67b57d43a79b := __obf_3e3889a3ca3aa32a.(type) {
	case []byte:
		*__obf_153c6aff24c72d24 = Str(__obf_82ab67b57d43a79b)
	case string:
		*__obf_153c6aff24c72d24 = Str(__obf_82ab67b57d43a79b)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_3e3889a3ca3aa32a)
	}
	return nil
}

func (__obf_153c6aff24c72d24 *Str) Str() string {
	if __obf_153c6aff24c72d24 == nil {
		return ""
	}
	return string(*__obf_153c6aff24c72d24)
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

func (__obf_be83ea6d875ef30b *Int) Scan(__obf_3e3889a3ca3aa32a any) error {
	if __obf_3e3889a3ca3aa32a == nil {
		return nil
	}

	switch __obf_82ab67b57d43a79b := __obf_3e3889a3ca3aa32a.(type) {
	case []byte:
		__obf_b6bf8b8bc11918d7, _ := strconv.ParseInt(string(__obf_82ab67b57d43a79b), 10, 64)
		*__obf_be83ea6d875ef30b = Int(__obf_b6bf8b8bc11918d7)
	case int64:
		*__obf_be83ea6d875ef30b = Int(__obf_82ab67b57d43a79b)
	case uint64:
		*__obf_be83ea6d875ef30b = Int(__obf_82ab67b57d43a79b)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_3e3889a3ca3aa32a)
	}
	return nil
}

func (__obf_be83ea6d875ef30b *Int) Int() int64 {
	if __obf_be83ea6d875ef30b == nil {
		return 0
	}
	return int64(*__obf_be83ea6d875ef30b)
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

func (__obf_715a8e43dc438a16 *Float) Scan(__obf_3e3889a3ca3aa32a any) error {
	if __obf_3e3889a3ca3aa32a == nil {
		return nil
	}

	switch __obf_82ab67b57d43a79b := __obf_3e3889a3ca3aa32a.(type) {
	case []byte:
		if __obf_153c6aff24c72d24, __obf_c5f1094aec829fc9 := strconv.ParseFloat(string(__obf_82ab67b57d43a79b), 64); __obf_c5f1094aec829fc9 != nil {
			return __obf_c5f1094aec829fc9
		} else {
			*__obf_715a8e43dc438a16 = Float(__obf_153c6aff24c72d24)
		}
	case float64:
		*__obf_715a8e43dc438a16 = Float(__obf_82ab67b57d43a79b)
	default:
		return fmt.Errorf("unsupported type for: %T", __obf_3e3889a3ca3aa32a)
	}

	return nil
}

func (__obf_be83ea6d875ef30b *Float) Float() float64 {
	if __obf_be83ea6d875ef30b == nil {
		return 0
	}
	return float64(*__obf_be83ea6d875ef30b)
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
func DoScan[T any](__obf_f0093308c4041e65 T, __obf_b6bf8b8bc11918d7 any) error {
	if __obf_b6bf8b8bc11918d7 == nil {
		return nil
	}
	var __obf_28d88d9264dbbdc8 []byte
	switch __obf_b6bf8b8bc11918d7 := __obf_b6bf8b8bc11918d7.(type) {
	case T:
		__obf_f0093308c4041e65 = __obf_b6bf8b8bc11918d7
		return nil
	case string:
		__obf_28d88d9264dbbdc8 = []byte(__obf_b6bf8b8bc11918d7)
	case []byte:
		__obf_28d88d9264dbbdc8 = __obf_b6bf8b8bc11918d7
	default:
		return fmt.Errorf("incompatible type for %T", __obf_b6bf8b8bc11918d7)
	}

	if len(__obf_28d88d9264dbbdc8) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_28d88d9264dbbdc8, &__obf_f0093308c4041e65)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_82ab67b57d43a79b *JsonMap) Scan(__obf_a12d79aa3dd7e532 any) error {
	return DoScan(__obf_82ab67b57d43a79b, __obf_a12d79aa3dd7e532)
}

func (__obf_82ab67b57d43a79b JsonMap) Value() (driver.Value, error) {
	if len(__obf_82ab67b57d43a79b) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_82ab67b57d43a79b)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_d51865a84a0adff4 *StringArray) Scan(__obf_a12d79aa3dd7e532 any) error {
	return DoScan(__obf_d51865a84a0adff4, __obf_a12d79aa3dd7e532)

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

func (__obf_82ab67b57d43a79b StringArray) Value() (driver.Value, error) {
	if len(__obf_82ab67b57d43a79b) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_82ab67b57d43a79b)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_82ab67b57d43a79b *IntArray) Scan(__obf_a12d79aa3dd7e532 any) error {
	return DoScan(__obf_82ab67b57d43a79b, __obf_a12d79aa3dd7e532)
}

func (__obf_82ab67b57d43a79b IntArray) Value() (driver.Value, error) {
	if len(__obf_82ab67b57d43a79b) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_82ab67b57d43a79b)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_82ab67b57d43a79b *MapArray) Scan(__obf_a12d79aa3dd7e532 any) error {
	return DoScan(__obf_82ab67b57d43a79b, __obf_a12d79aa3dd7e532)
}

func (__obf_82ab67b57d43a79b MapArray) Value() (driver.Value, error) {
	if len(__obf_82ab67b57d43a79b) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_82ab67b57d43a79b)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_82ab67b57d43a79b *CaseTask) Scan(__obf_a12d79aa3dd7e532 any) error {
	return DoScan(__obf_82ab67b57d43a79b, __obf_a12d79aa3dd7e532)
}

func (__obf_82ab67b57d43a79b CaseTask) Value() (driver.Value, error) {
	if len(__obf_82ab67b57d43a79b) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_82ab67b57d43a79b)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_82ab67b57d43a79b *Timeline) Scan(__obf_a12d79aa3dd7e532 any) error {
	return DoScan(__obf_82ab67b57d43a79b, __obf_a12d79aa3dd7e532)
}

func (__obf_82ab67b57d43a79b Timeline) Value() (driver.Value, error) {
	if len(__obf_82ab67b57d43a79b) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_82ab67b57d43a79b)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_82ab67b57d43a79b Timeline) TimeLeft() int64 {
	__obf_e72f8e2f354cf695 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_dd18fa1f94c40d1f := range __obf_82ab67b57d43a79b {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_e72f8e2f354cf695 >= __obf_dd18fa1f94c40d1f.Start && __obf_e72f8e2f354cf695 < __obf_dd18fa1f94c40d1f.End {
			return __obf_dd18fa1f94c40d1f.End - __obf_e72f8e2f354cf695
		}
	}

	return 0
}

func (__obf_82ab67b57d43a79b Timeline) IsUpcoming() bool {
	__obf_e72f8e2f354cf695 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_dd18fa1f94c40d1f := range __obf_82ab67b57d43a79b {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_dd18fa1f94c40d1f.End > __obf_e72f8e2f354cf695 {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_b1bac5403586174c RawJson) Message() json.RawMessage {
	if len(__obf_b1bac5403586174c) == 0 {
		return nil
	}
	return json.RawMessage(__obf_b1bac5403586174c)
}

// Scan implements the Scanner interface.
func (__obf_b1bac5403586174c *RawJson) Scan(__obf_a12d79aa3dd7e532 any) error {
	if __obf_a12d79aa3dd7e532 == nil {
		*__obf_b1bac5403586174c = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_45765a4c81a9a40e, __obf_f65665999cfd53f1 := __obf_a12d79aa3dd7e532.([]byte)
	if !__obf_f65665999cfd53f1 {
		// Handle other types if necessary, e.g., string
		__obf_0c9898604d2c7d32, __obf_f65665999cfd53f1 := __obf_a12d79aa3dd7e532.(string)
		if __obf_f65665999cfd53f1 {
			__obf_45765a4c81a9a40e = []byte(__obf_0c9898604d2c7d32)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_a12d79aa3dd7e532)
		}
	}
	*__obf_b1bac5403586174c = RawJson(__obf_45765a4c81a9a40e)
	return nil
}

// Value implements the Valuer interface.
func (__obf_b1bac5403586174c RawJson) Value() (driver.Value, error) {
	if len(__obf_b1bac5403586174c) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_b1bac5403586174c), nil
}
