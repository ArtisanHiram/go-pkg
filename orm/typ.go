package __obf_214d1e062aee9185

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
func (__obf_39f34c7e1dead46b *Str) Scan(__obf_3027e28da59d2939 any) error {
	if __obf_3027e28da59d2939 == nil {
		return nil
	}

	switch __obf_e7e84afe571f9ceb := __obf_3027e28da59d2939.(type) {
	case []byte:
		*__obf_39f34c7e1dead46b = Str(__obf_e7e84afe571f9ceb)
	case string:
		*__obf_39f34c7e1dead46b = Str(__obf_e7e84afe571f9ceb)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_3027e28da59d2939)
	}
	return nil
}

func (__obf_39f34c7e1dead46b *Str) Str() string {
	if __obf_39f34c7e1dead46b == nil {
		return ""
	}
	return string(*__obf_39f34c7e1dead46b)
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

func (__obf_7c6c9e1512e37985 *Int) Scan(__obf_3027e28da59d2939 any) error {
	if __obf_3027e28da59d2939 == nil {
		return nil
	}

	switch __obf_e7e84afe571f9ceb := __obf_3027e28da59d2939.(type) {
	case []byte:
		__obf_a73e9633161d85ba, _ := strconv.ParseInt(string(__obf_e7e84afe571f9ceb), 10, 64)
		*__obf_7c6c9e1512e37985 = Int(__obf_a73e9633161d85ba)
	case int64:
		*__obf_7c6c9e1512e37985 = Int(__obf_e7e84afe571f9ceb)
	case uint64:
		*__obf_7c6c9e1512e37985 = Int(__obf_e7e84afe571f9ceb)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_3027e28da59d2939)
	}
	return nil
}

func (__obf_7c6c9e1512e37985 *Int) Int() int64 {
	if __obf_7c6c9e1512e37985 == nil {
		return 0
	}
	return int64(*__obf_7c6c9e1512e37985)
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

func (__obf_e242c9b073a9f4fe *Float) Scan(__obf_3027e28da59d2939 any) error {
	if __obf_3027e28da59d2939 == nil {
		return nil
	}

	switch __obf_e7e84afe571f9ceb := __obf_3027e28da59d2939.(type) {
	case []byte:
		if __obf_39f34c7e1dead46b, __obf_63349b4c89eea887 := strconv.ParseFloat(string(__obf_e7e84afe571f9ceb), 64); __obf_63349b4c89eea887 != nil {
			return __obf_63349b4c89eea887
		} else {
			*__obf_e242c9b073a9f4fe = Float(__obf_39f34c7e1dead46b)
		}
	case float64:
		*__obf_e242c9b073a9f4fe = Float(__obf_e7e84afe571f9ceb)
	default:
		return fmt.Errorf("unsupported type for: %T", __obf_3027e28da59d2939)
	}

	return nil
}

func (__obf_7c6c9e1512e37985 *Float) Float() float64 {
	if __obf_7c6c9e1512e37985 == nil {
		return 0
	}
	return float64(*__obf_7c6c9e1512e37985)
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
func DoScan[T any](__obf_ef85a2cf30bcbb02 T, __obf_a73e9633161d85ba any) error {
	if __obf_a73e9633161d85ba == nil {
		return nil
	}
	var __obf_3ca90ccbf2253f6f []byte
	switch __obf_a73e9633161d85ba := __obf_a73e9633161d85ba.(type) {
	case T:
		__obf_ef85a2cf30bcbb02 = __obf_a73e9633161d85ba
		return nil
	case string:
		__obf_3ca90ccbf2253f6f = []byte(__obf_a73e9633161d85ba)
	case []byte:
		__obf_3ca90ccbf2253f6f = __obf_a73e9633161d85ba
	default:
		return fmt.Errorf("incompatible type for %T", __obf_a73e9633161d85ba)
	}

	if len(__obf_3ca90ccbf2253f6f) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_3ca90ccbf2253f6f, &__obf_ef85a2cf30bcbb02)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_e7e84afe571f9ceb *JsonMap) Scan(__obf_5dbd5d199912b6ea any) error {
	return DoScan(__obf_e7e84afe571f9ceb, __obf_5dbd5d199912b6ea)
}

func (__obf_e7e84afe571f9ceb JsonMap) Value() (driver.Value, error) {
	if len(__obf_e7e84afe571f9ceb) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_e7e84afe571f9ceb)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_5749cb79bd3e35c1 *StringArray) Scan(__obf_5dbd5d199912b6ea any) error {
	return DoScan(__obf_5749cb79bd3e35c1,

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
		__obf_5dbd5d199912b6ea)

	// return json.Unmarshal(bytes, &ja)
}

func (__obf_e7e84afe571f9ceb StringArray) Value() (driver.Value, error) {
	if len(__obf_e7e84afe571f9ceb) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_e7e84afe571f9ceb)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_e7e84afe571f9ceb *IntArray) Scan(__obf_5dbd5d199912b6ea any) error {
	return DoScan(__obf_e7e84afe571f9ceb, __obf_5dbd5d199912b6ea)
}

func (__obf_e7e84afe571f9ceb IntArray) Value() (driver.Value, error) {
	if len(__obf_e7e84afe571f9ceb) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_e7e84afe571f9ceb)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_e7e84afe571f9ceb *MapArray) Scan(__obf_5dbd5d199912b6ea any) error {
	return DoScan(__obf_e7e84afe571f9ceb, __obf_5dbd5d199912b6ea)
}

func (__obf_e7e84afe571f9ceb MapArray) Value() (driver.Value, error) {
	if len(__obf_e7e84afe571f9ceb) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_e7e84afe571f9ceb)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_e7e84afe571f9ceb *CaseTask) Scan(__obf_5dbd5d199912b6ea any) error {
	return DoScan(__obf_e7e84afe571f9ceb, __obf_5dbd5d199912b6ea)
}

func (__obf_e7e84afe571f9ceb CaseTask) Value() (driver.Value, error) {
	if len(__obf_e7e84afe571f9ceb) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_e7e84afe571f9ceb)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_e7e84afe571f9ceb *Timeline) Scan(__obf_5dbd5d199912b6ea any) error {
	return DoScan(__obf_e7e84afe571f9ceb, __obf_5dbd5d199912b6ea)
}

func (__obf_e7e84afe571f9ceb Timeline) Value() (driver.Value, error) {
	if len(__obf_e7e84afe571f9ceb) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_e7e84afe571f9ceb)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_e7e84afe571f9ceb Timeline) TimeLeft() int64 {
	__obf_cfe49a2e057394a9 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_f8fb2d3721af23b4 := range __obf_e7e84afe571f9ceb {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_cfe49a2e057394a9 >= __obf_f8fb2d3721af23b4.Start && __obf_cfe49a2e057394a9 < __obf_f8fb2d3721af23b4.End {
			return __obf_f8fb2d3721af23b4.End - __obf_cfe49a2e057394a9
		}
	}

	return 0
}

func (__obf_e7e84afe571f9ceb Timeline) IsUpcoming() bool {
	__obf_cfe49a2e057394a9 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_f8fb2d3721af23b4 := range __obf_e7e84afe571f9ceb {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_f8fb2d3721af23b4.End > __obf_cfe49a2e057394a9 {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_c946acaed56b4864 RawJson) Message() json.RawMessage {
	if len(__obf_c946acaed56b4864) == 0 {
		return nil
	}
	return json.RawMessage(__obf_c946acaed56b4864)
}

// Scan implements the Scanner interface.
func (__obf_c946acaed56b4864 *RawJson) Scan(__obf_5dbd5d199912b6ea any) error {
	if __obf_5dbd5d199912b6ea == nil {
		*__obf_c946acaed56b4864 = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_de4c5fac44bd0b92, __obf_65bc54d69fb5f329 := __obf_5dbd5d199912b6ea.([]byte)
	if !__obf_65bc54d69fb5f329 {
		__obf_8342a8c0dcd918b7,
			// Handle other types if necessary, e.g., string
			__obf_65bc54d69fb5f329 := __obf_5dbd5d199912b6ea.(string)
		if __obf_65bc54d69fb5f329 {
			__obf_de4c5fac44bd0b92 = []byte(__obf_8342a8c0dcd918b7)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_5dbd5d199912b6ea)
		}
	}
	*__obf_c946acaed56b4864 = RawJson(__obf_de4c5fac44bd0b92)
	return nil
}

// Value implements the Valuer interface.
func (__obf_c946acaed56b4864 RawJson) Value() (driver.Value, error) {
	if len(__obf_c946acaed56b4864) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_c946acaed56b4864), nil
}
