package __obf_ea545e4bdf748fd2

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
func (__obf_4c5a866ce1415830 *Str) Scan(__obf_3233962bdf9d4ee1 any) error {
	if __obf_3233962bdf9d4ee1 == nil {
		return nil
	}

	switch __obf_5f939ebac6d18417 := __obf_3233962bdf9d4ee1.(type) {
	case []byte:
		*__obf_4c5a866ce1415830 = Str(__obf_5f939ebac6d18417)
	case string:
		*__obf_4c5a866ce1415830 = Str(__obf_5f939ebac6d18417)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_3233962bdf9d4ee1)
	}
	return nil
}

func (__obf_4c5a866ce1415830 *Str) Str() string {
	if __obf_4c5a866ce1415830 == nil {
		return ""
	}
	return string(*__obf_4c5a866ce1415830)
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

func (__obf_5f1ac33f5a6c1104 *Int) Scan(__obf_3233962bdf9d4ee1 any) error {
	if __obf_3233962bdf9d4ee1 == nil {
		return nil
	}

	switch __obf_5f939ebac6d18417 := __obf_3233962bdf9d4ee1.(type) {
	case []byte:
		__obf_abd89eb645093144, _ := strconv.ParseInt(string(__obf_5f939ebac6d18417), 10, 64)
		*__obf_5f1ac33f5a6c1104 = Int(__obf_abd89eb645093144)
	case int64:
		*__obf_5f1ac33f5a6c1104 = Int(__obf_5f939ebac6d18417)
	case uint64:
		*__obf_5f1ac33f5a6c1104 = Int(__obf_5f939ebac6d18417)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_3233962bdf9d4ee1)
	}
	return nil
}

func (__obf_5f1ac33f5a6c1104 *Int) Int() int64 {
	if __obf_5f1ac33f5a6c1104 == nil {
		return 0
	}
	return int64(*__obf_5f1ac33f5a6c1104)
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

func (__obf_af66979cc7b3eb24 *Float) Scan(__obf_3233962bdf9d4ee1 any) error {
	if __obf_3233962bdf9d4ee1 == nil {
		return nil
	}

	switch __obf_5f939ebac6d18417 := __obf_3233962bdf9d4ee1.(type) {
	case []byte:
		if __obf_4c5a866ce1415830, __obf_3220344ba911fc77 := strconv.ParseFloat(string(__obf_5f939ebac6d18417), 64); __obf_3220344ba911fc77 != nil {
			return __obf_3220344ba911fc77
		} else {
			*__obf_af66979cc7b3eb24 = Float(__obf_4c5a866ce1415830)
		}
	case float64:
		*__obf_af66979cc7b3eb24 = Float(__obf_5f939ebac6d18417)
	default:
		return fmt.Errorf("unsupported type for: %T", __obf_3233962bdf9d4ee1)
	}

	return nil
}

func (__obf_5f1ac33f5a6c1104 *Float) Float() float64 {
	if __obf_5f1ac33f5a6c1104 == nil {
		return 0
	}
	return float64(*__obf_5f1ac33f5a6c1104)
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
func DoScan[T any](__obf_bdb37dd60be9d858 T, __obf_abd89eb645093144 any) error {
	if __obf_abd89eb645093144 == nil {
		return nil
	}
	var __obf_1d87051673d8e599 []byte
	switch __obf_abd89eb645093144 := __obf_abd89eb645093144.(type) {
	case T:
		__obf_bdb37dd60be9d858 = __obf_abd89eb645093144
		return nil
	case string:
		__obf_1d87051673d8e599 = []byte(__obf_abd89eb645093144)
	case []byte:
		__obf_1d87051673d8e599 = __obf_abd89eb645093144
	default:
		return fmt.Errorf("incompatible type for %T", __obf_abd89eb645093144)
	}

	if len(__obf_1d87051673d8e599) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_1d87051673d8e599, &__obf_bdb37dd60be9d858)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_5f939ebac6d18417 *JsonMap) Scan(__obf_43a03d8b7535a35f any) error {
	return DoScan(__obf_5f939ebac6d18417, __obf_43a03d8b7535a35f)
}

func (__obf_5f939ebac6d18417 JsonMap) Value() (driver.Value, error) {
	if len(__obf_5f939ebac6d18417) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_5f939ebac6d18417)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_2f9ae813edcaa1b8 *StringArray) Scan(__obf_43a03d8b7535a35f any) error {
	return DoScan(__obf_2f9ae813edcaa1b8,

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
		__obf_43a03d8b7535a35f)

	// return json.Unmarshal(bytes, &ja)
}

func (__obf_5f939ebac6d18417 StringArray) Value() (driver.Value, error) {
	if len(__obf_5f939ebac6d18417) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_5f939ebac6d18417)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_5f939ebac6d18417 *IntArray) Scan(__obf_43a03d8b7535a35f any) error {
	return DoScan(__obf_5f939ebac6d18417, __obf_43a03d8b7535a35f)
}

func (__obf_5f939ebac6d18417 IntArray) Value() (driver.Value, error) {
	if len(__obf_5f939ebac6d18417) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_5f939ebac6d18417)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_5f939ebac6d18417 *MapArray) Scan(__obf_43a03d8b7535a35f any) error {
	return DoScan(__obf_5f939ebac6d18417, __obf_43a03d8b7535a35f)
}

func (__obf_5f939ebac6d18417 MapArray) Value() (driver.Value, error) {
	if len(__obf_5f939ebac6d18417) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_5f939ebac6d18417)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_5f939ebac6d18417 *CaseTask) Scan(__obf_43a03d8b7535a35f any) error {
	return DoScan(__obf_5f939ebac6d18417, __obf_43a03d8b7535a35f)
}

func (__obf_5f939ebac6d18417 CaseTask) Value() (driver.Value, error) {
	if len(__obf_5f939ebac6d18417) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_5f939ebac6d18417)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_5f939ebac6d18417 *Timeline) Scan(__obf_43a03d8b7535a35f any) error {
	return DoScan(__obf_5f939ebac6d18417, __obf_43a03d8b7535a35f)
}

func (__obf_5f939ebac6d18417 Timeline) Value() (driver.Value, error) {
	if len(__obf_5f939ebac6d18417) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_5f939ebac6d18417)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_5f939ebac6d18417 Timeline) TimeLeft() int64 {
	__obf_d7cc9de281829818 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_cd315e33a9f6fe4e := range __obf_5f939ebac6d18417 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_d7cc9de281829818 >= __obf_cd315e33a9f6fe4e.Start && __obf_d7cc9de281829818 < __obf_cd315e33a9f6fe4e.End {
			return __obf_cd315e33a9f6fe4e.End - __obf_d7cc9de281829818
		}
	}

	return 0
}

func (__obf_5f939ebac6d18417 Timeline) IsUpcoming() bool {
	__obf_d7cc9de281829818 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_cd315e33a9f6fe4e := range __obf_5f939ebac6d18417 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_cd315e33a9f6fe4e.End > __obf_d7cc9de281829818 {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_c10bce80e1d2cd9e RawJson) Message() json.RawMessage {
	if len(__obf_c10bce80e1d2cd9e) == 0 {
		return nil
	}
	return json.RawMessage(__obf_c10bce80e1d2cd9e)
}

// Scan implements the Scanner interface.
func (__obf_c10bce80e1d2cd9e *RawJson) Scan(__obf_43a03d8b7535a35f any) error {
	if __obf_43a03d8b7535a35f == nil {
		*__obf_c10bce80e1d2cd9e = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_1e8e0c77f709c270, __obf_fcb98d9ec4b8d44c := __obf_43a03d8b7535a35f.([]byte)
	if !__obf_fcb98d9ec4b8d44c {
		__obf_1b9c3122b7c0647b,
			// Handle other types if necessary, e.g., string
			__obf_fcb98d9ec4b8d44c := __obf_43a03d8b7535a35f.(string)
		if __obf_fcb98d9ec4b8d44c {
			__obf_1e8e0c77f709c270 = []byte(__obf_1b9c3122b7c0647b)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_43a03d8b7535a35f)
		}
	}
	*__obf_c10bce80e1d2cd9e = RawJson(__obf_1e8e0c77f709c270)
	return nil
}

// Value implements the Valuer interface.
func (__obf_c10bce80e1d2cd9e RawJson) Value() (driver.Value, error) {
	if len(__obf_c10bce80e1d2cd9e) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_c10bce80e1d2cd9e), nil
}
