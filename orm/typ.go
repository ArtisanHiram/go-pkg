package __obf_f47aac06a08e5dbb

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
func (__obf_42c5c2ee071be035 *Str) Scan(__obf_e78ca1d20994615b any) error {
	if __obf_e78ca1d20994615b == nil {
		return nil
	}

	switch __obf_c3d7728bbd58f361 := __obf_e78ca1d20994615b.(type) {
	case []byte:
		*__obf_42c5c2ee071be035 = Str(__obf_c3d7728bbd58f361)
	case string:
		*__obf_42c5c2ee071be035 = Str(__obf_c3d7728bbd58f361)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_e78ca1d20994615b)
	}
	return nil
}

func (__obf_42c5c2ee071be035 *Str) Str() string {
	if __obf_42c5c2ee071be035 == nil {
		return ""
	}
	return string(*__obf_42c5c2ee071be035)
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

func (__obf_89306d663fb5edd8 *Int) Scan(__obf_e78ca1d20994615b any) error {
	if __obf_e78ca1d20994615b == nil {
		return nil
	}

	switch __obf_c3d7728bbd58f361 := __obf_e78ca1d20994615b.(type) {
	case []byte:
		__obf_613929daa0d28687, _ := strconv.ParseInt(string(__obf_c3d7728bbd58f361), 10, 64)
		*__obf_89306d663fb5edd8 = Int(__obf_613929daa0d28687)
	case int64:
		*__obf_89306d663fb5edd8 = Int(__obf_c3d7728bbd58f361)
	case uint64:
		*__obf_89306d663fb5edd8 = Int(__obf_c3d7728bbd58f361)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_e78ca1d20994615b)
	}
	return nil
}

func (__obf_89306d663fb5edd8 *Int) Int() int64 {
	if __obf_89306d663fb5edd8 == nil {
		return 0
	}
	return int64(*__obf_89306d663fb5edd8)
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

func (__obf_c9ab6b5901a38cd0 *Float) Scan(__obf_e78ca1d20994615b any) error {
	if __obf_e78ca1d20994615b == nil {
		return nil
	}

	switch __obf_c3d7728bbd58f361 := __obf_e78ca1d20994615b.(type) {
	case []byte:
		if __obf_42c5c2ee071be035, __obf_9ac6018edd4832e7 := strconv.ParseFloat(string(__obf_c3d7728bbd58f361), 64); __obf_9ac6018edd4832e7 != nil {
			return __obf_9ac6018edd4832e7
		} else {
			*__obf_c9ab6b5901a38cd0 = Float(__obf_42c5c2ee071be035)
		}
	case float64:
		*__obf_c9ab6b5901a38cd0 = Float(__obf_c3d7728bbd58f361)
	default:
		return fmt.Errorf("unsupported type for: %T", __obf_e78ca1d20994615b)
	}

	return nil
}

func (__obf_89306d663fb5edd8 *Float) Float() float64 {
	if __obf_89306d663fb5edd8 == nil {
		return 0
	}
	return float64(*__obf_89306d663fb5edd8)
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
func DoScan[T any](__obf_55c3f06639459f3a T, __obf_613929daa0d28687 any) error {
	if __obf_613929daa0d28687 == nil {
		return nil
	}
	var __obf_cf1876786a770d2d []byte
	switch __obf_613929daa0d28687 := __obf_613929daa0d28687.(type) {
	case T:
		__obf_55c3f06639459f3a = __obf_613929daa0d28687
		return nil
	case string:
		__obf_cf1876786a770d2d = []byte(__obf_613929daa0d28687)
	case []byte:
		__obf_cf1876786a770d2d = __obf_613929daa0d28687
	default:
		return fmt.Errorf("incompatible type for %T", __obf_613929daa0d28687)
	}

	if len(__obf_cf1876786a770d2d) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_cf1876786a770d2d, &__obf_55c3f06639459f3a)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_c3d7728bbd58f361 *JsonMap) Scan(__obf_e51512e07c73036f any) error {
	return DoScan(__obf_c3d7728bbd58f361, __obf_e51512e07c73036f)
}

func (__obf_c3d7728bbd58f361 JsonMap) Value() (driver.Value, error) {
	if len(__obf_c3d7728bbd58f361) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_c3d7728bbd58f361)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_1ac659f9858199ca *StringArray) Scan(__obf_e51512e07c73036f any) error {
	return DoScan(__obf_1ac659f9858199ca,

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
		__obf_e51512e07c73036f)

	// return json.Unmarshal(bytes, &ja)
}

func (__obf_c3d7728bbd58f361 StringArray) Value() (driver.Value, error) {
	if len(__obf_c3d7728bbd58f361) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_c3d7728bbd58f361)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_c3d7728bbd58f361 *IntArray) Scan(__obf_e51512e07c73036f any) error {
	return DoScan(__obf_c3d7728bbd58f361, __obf_e51512e07c73036f)
}

func (__obf_c3d7728bbd58f361 IntArray) Value() (driver.Value, error) {
	if len(__obf_c3d7728bbd58f361) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_c3d7728bbd58f361)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_c3d7728bbd58f361 *MapArray) Scan(__obf_e51512e07c73036f any) error {
	return DoScan(__obf_c3d7728bbd58f361, __obf_e51512e07c73036f)
}

func (__obf_c3d7728bbd58f361 MapArray) Value() (driver.Value, error) {
	if len(__obf_c3d7728bbd58f361) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_c3d7728bbd58f361)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_c3d7728bbd58f361 *CaseTask) Scan(__obf_e51512e07c73036f any) error {
	return DoScan(__obf_c3d7728bbd58f361, __obf_e51512e07c73036f)
}

func (__obf_c3d7728bbd58f361 CaseTask) Value() (driver.Value, error) {
	if len(__obf_c3d7728bbd58f361) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_c3d7728bbd58f361)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_c3d7728bbd58f361 *Timeline) Scan(__obf_e51512e07c73036f any) error {
	return DoScan(__obf_c3d7728bbd58f361, __obf_e51512e07c73036f)
}

func (__obf_c3d7728bbd58f361 Timeline) Value() (driver.Value, error) {
	if len(__obf_c3d7728bbd58f361) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_c3d7728bbd58f361)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_c3d7728bbd58f361 Timeline) TimeLeft() int64 {
	__obf_7884129cc9b29982 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_0049191641544cff := range __obf_c3d7728bbd58f361 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_7884129cc9b29982 >= __obf_0049191641544cff.Start && __obf_7884129cc9b29982 < __obf_0049191641544cff.End {
			return __obf_0049191641544cff.End - __obf_7884129cc9b29982
		}
	}

	return 0
}

func (__obf_c3d7728bbd58f361 Timeline) IsUpcoming() bool {
	__obf_7884129cc9b29982 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_0049191641544cff := range __obf_c3d7728bbd58f361 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_0049191641544cff.End > __obf_7884129cc9b29982 {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_6db17027e7f3f06b RawJson) Message() json.RawMessage {
	if len(__obf_6db17027e7f3f06b) == 0 {
		return nil
	}
	return json.RawMessage(__obf_6db17027e7f3f06b)
}

// Scan implements the Scanner interface.
func (__obf_6db17027e7f3f06b *RawJson) Scan(__obf_e51512e07c73036f any) error {
	if __obf_e51512e07c73036f == nil {
		*__obf_6db17027e7f3f06b = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_df36da84e6549012, __obf_410061093b631faf := __obf_e51512e07c73036f.([]byte)
	if !__obf_410061093b631faf {
		__obf_06abf9ba262d4c9a,
			// Handle other types if necessary, e.g., string
			__obf_410061093b631faf := __obf_e51512e07c73036f.(string)
		if __obf_410061093b631faf {
			__obf_df36da84e6549012 = []byte(__obf_06abf9ba262d4c9a)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_e51512e07c73036f)
		}
	}
	*__obf_6db17027e7f3f06b = RawJson(__obf_df36da84e6549012)
	return nil
}

// Value implements the Valuer interface.
func (__obf_6db17027e7f3f06b RawJson) Value() (driver.Value, error) {
	if len(__obf_6db17027e7f3f06b) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_6db17027e7f3f06b), nil
}
