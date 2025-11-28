package __obf_d726bb43e85f2dfc

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
func (__obf_3f6b3a80155a822c *Str) Scan(__obf_27bbc0c6e11f098d any) error {
	if __obf_27bbc0c6e11f098d == nil {
		return nil
	}

	switch __obf_23f8e72519d9c06e := __obf_27bbc0c6e11f098d.(type) {
	case []byte:
		*__obf_3f6b3a80155a822c = Str(__obf_23f8e72519d9c06e)
	case string:
		*__obf_3f6b3a80155a822c = Str(__obf_23f8e72519d9c06e)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_27bbc0c6e11f098d)
	}
	return nil
}

func (__obf_3f6b3a80155a822c *Str) Str() string {
	if __obf_3f6b3a80155a822c == nil {
		return ""
	}
	return string(*__obf_3f6b3a80155a822c)
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

func (__obf_cbebdbf7ebca6a65 *Int) Scan(__obf_27bbc0c6e11f098d any) error {
	if __obf_27bbc0c6e11f098d == nil {
		return nil
	}

	switch __obf_23f8e72519d9c06e := __obf_27bbc0c6e11f098d.(type) {
	case []byte:
		__obf_145a2e7072310087, _ := strconv.ParseInt(string(__obf_23f8e72519d9c06e), 10, 64)
		*__obf_cbebdbf7ebca6a65 = Int(__obf_145a2e7072310087)
	case int64:
		*__obf_cbebdbf7ebca6a65 = Int(__obf_23f8e72519d9c06e)
	case uint64:
		*__obf_cbebdbf7ebca6a65 = Int(__obf_23f8e72519d9c06e)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_27bbc0c6e11f098d)
	}
	return nil
}

func (__obf_cbebdbf7ebca6a65 *Int) Int() int64 {
	if __obf_cbebdbf7ebca6a65 == nil {
		return 0
	}
	return int64(*__obf_cbebdbf7ebca6a65)
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

func (__obf_b0142d337517f48f *Float) Scan(__obf_27bbc0c6e11f098d any) error {
	if __obf_27bbc0c6e11f098d == nil {
		return nil
	}

	switch __obf_23f8e72519d9c06e := __obf_27bbc0c6e11f098d.(type) {
	case []byte:
		if __obf_3f6b3a80155a822c, __obf_f5c7d5fde3143a52 := strconv.ParseFloat(string(__obf_23f8e72519d9c06e), 64); __obf_f5c7d5fde3143a52 != nil {
			return __obf_f5c7d5fde3143a52
		} else {
			*__obf_b0142d337517f48f = Float(__obf_3f6b3a80155a822c)
		}
	case float64:
		*__obf_b0142d337517f48f = Float(__obf_23f8e72519d9c06e)
	default:
		return fmt.Errorf("unsupported type for: %T", __obf_27bbc0c6e11f098d)
	}

	return nil
}

func (__obf_cbebdbf7ebca6a65 *Float) Float() float64 {
	if __obf_cbebdbf7ebca6a65 == nil {
		return 0
	}
	return float64(*__obf_cbebdbf7ebca6a65)
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
func DoScan[T any](__obf_f4ddd31af644b83b T, __obf_145a2e7072310087 any) error {
	if __obf_145a2e7072310087 == nil {
		return nil
	}
	var __obf_0bf7d9ba888cf6cc []byte
	switch __obf_145a2e7072310087 := __obf_145a2e7072310087.(type) {
	case T:
		__obf_f4ddd31af644b83b = __obf_145a2e7072310087
		return nil
	case string:
		__obf_0bf7d9ba888cf6cc = []byte(__obf_145a2e7072310087)
	case []byte:
		__obf_0bf7d9ba888cf6cc = __obf_145a2e7072310087
	default:
		return fmt.Errorf("incompatible type for %T", __obf_145a2e7072310087)
	}

	if len(__obf_0bf7d9ba888cf6cc) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_0bf7d9ba888cf6cc, &__obf_f4ddd31af644b83b)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_23f8e72519d9c06e *JsonMap) Scan(__obf_14f652161c66608a any) error {
	return DoScan(__obf_23f8e72519d9c06e, __obf_14f652161c66608a)
}

func (__obf_23f8e72519d9c06e JsonMap) Value() (driver.Value, error) {
	if len(__obf_23f8e72519d9c06e) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_23f8e72519d9c06e)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_37e6167635151ccd *StringArray) Scan(__obf_14f652161c66608a any) error {
	return DoScan(__obf_37e6167635151ccd, __obf_14f652161c66608a)

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

func (__obf_23f8e72519d9c06e StringArray) Value() (driver.Value, error) {
	if len(__obf_23f8e72519d9c06e) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_23f8e72519d9c06e)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_23f8e72519d9c06e *IntArray) Scan(__obf_14f652161c66608a any) error {
	return DoScan(__obf_23f8e72519d9c06e, __obf_14f652161c66608a)
}

func (__obf_23f8e72519d9c06e IntArray) Value() (driver.Value, error) {
	if len(__obf_23f8e72519d9c06e) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_23f8e72519d9c06e)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_23f8e72519d9c06e *MapArray) Scan(__obf_14f652161c66608a any) error {
	return DoScan(__obf_23f8e72519d9c06e, __obf_14f652161c66608a)
}

func (__obf_23f8e72519d9c06e MapArray) Value() (driver.Value, error) {
	if len(__obf_23f8e72519d9c06e) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_23f8e72519d9c06e)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_23f8e72519d9c06e *CaseTask) Scan(__obf_14f652161c66608a any) error {
	return DoScan(__obf_23f8e72519d9c06e, __obf_14f652161c66608a)
}

func (__obf_23f8e72519d9c06e CaseTask) Value() (driver.Value, error) {
	if len(__obf_23f8e72519d9c06e) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_23f8e72519d9c06e)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_23f8e72519d9c06e *Timeline) Scan(__obf_14f652161c66608a any) error {
	return DoScan(__obf_23f8e72519d9c06e, __obf_14f652161c66608a)
}

func (__obf_23f8e72519d9c06e Timeline) Value() (driver.Value, error) {
	if len(__obf_23f8e72519d9c06e) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_23f8e72519d9c06e)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_23f8e72519d9c06e Timeline) TimeLeft() int64 {
	__obf_00a40ddda76aea5f := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_7faf168446410591 := range __obf_23f8e72519d9c06e {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_00a40ddda76aea5f >= __obf_7faf168446410591.Start && __obf_00a40ddda76aea5f < __obf_7faf168446410591.End {
			return __obf_7faf168446410591.End - __obf_00a40ddda76aea5f
		}
	}

	return 0
}

func (__obf_23f8e72519d9c06e Timeline) IsUpcoming() bool {
	__obf_00a40ddda76aea5f := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_7faf168446410591 := range __obf_23f8e72519d9c06e {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_7faf168446410591.End > __obf_00a40ddda76aea5f {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_20b73465be975bb2 RawJson) Message() json.RawMessage {
	if len(__obf_20b73465be975bb2) == 0 {
		return nil
	}
	return json.RawMessage(__obf_20b73465be975bb2)
}

// Scan implements the Scanner interface.
func (__obf_20b73465be975bb2 *RawJson) Scan(__obf_14f652161c66608a any) error {
	if __obf_14f652161c66608a == nil {
		*__obf_20b73465be975bb2 = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_39f20c87d32bf8d1, __obf_5dddaed8aec227de := __obf_14f652161c66608a.([]byte)
	if !__obf_5dddaed8aec227de {
		// Handle other types if necessary, e.g., string
		__obf_d5f4f8cfb4727077, __obf_5dddaed8aec227de := __obf_14f652161c66608a.(string)
		if __obf_5dddaed8aec227de {
			__obf_39f20c87d32bf8d1 = []byte(__obf_d5f4f8cfb4727077)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_14f652161c66608a)
		}
	}
	*__obf_20b73465be975bb2 = RawJson(__obf_39f20c87d32bf8d1)
	return nil
}

// Value implements the Valuer interface.
func (__obf_20b73465be975bb2 RawJson) Value() (driver.Value, error) {
	if len(__obf_20b73465be975bb2) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_20b73465be975bb2), nil
}
