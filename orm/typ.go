package __obf_ed189c965cdcd132

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
func (__obf_1e3af5b8f467d03d *Str) Scan(__obf_b93e1c63757055cb any) error {
	if __obf_b93e1c63757055cb == nil {
		return nil
	}

	switch __obf_34511df5079d8a2e := __obf_b93e1c63757055cb.(type) {
	case []byte:
		*__obf_1e3af5b8f467d03d = Str(__obf_34511df5079d8a2e)
	case string:
		*__obf_1e3af5b8f467d03d = Str(__obf_34511df5079d8a2e)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_b93e1c63757055cb)
	}
	return nil
}

func (__obf_1e3af5b8f467d03d *Str) Str() string {
	if __obf_1e3af5b8f467d03d == nil {
		return ""
	}
	return string(*__obf_1e3af5b8f467d03d)
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

func (__obf_c7e5bacf18525a19 *Int) Scan(__obf_b93e1c63757055cb any) error {
	if __obf_b93e1c63757055cb == nil {
		return nil
	}

	switch __obf_34511df5079d8a2e := __obf_b93e1c63757055cb.(type) {
	case []byte:
		__obf_f9cbd174629de499, _ := strconv.ParseInt(string(__obf_34511df5079d8a2e), 10, 64)
		*__obf_c7e5bacf18525a19 = Int(__obf_f9cbd174629de499)
	case int64:
		*__obf_c7e5bacf18525a19 = Int(__obf_34511df5079d8a2e)
	case uint64:
		*__obf_c7e5bacf18525a19 = Int(__obf_34511df5079d8a2e)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_b93e1c63757055cb)
	}
	return nil
}

func (__obf_c7e5bacf18525a19 *Int) Int() int64 {
	if __obf_c7e5bacf18525a19 == nil {
		return 0
	}
	return int64(*__obf_c7e5bacf18525a19)
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

func (__obf_626ed35c11f71017 *Float) Scan(__obf_b93e1c63757055cb any) error {
	if __obf_b93e1c63757055cb == nil {
		return nil
	}

	switch __obf_34511df5079d8a2e := __obf_b93e1c63757055cb.(type) {
	case []byte:
		if __obf_1e3af5b8f467d03d, __obf_9f5a64d352151704 := strconv.ParseFloat(string(__obf_34511df5079d8a2e), 64); __obf_9f5a64d352151704 != nil {
			return __obf_9f5a64d352151704
		} else {
			*__obf_626ed35c11f71017 = Float(__obf_1e3af5b8f467d03d)
		}
	case float64:
		*__obf_626ed35c11f71017 = Float(__obf_34511df5079d8a2e)
	default:
		return fmt.Errorf("unsupported type for: %T", __obf_b93e1c63757055cb)
	}

	return nil
}

func (__obf_c7e5bacf18525a19 *Float) Float() float64 {
	if __obf_c7e5bacf18525a19 == nil {
		return 0
	}
	return float64(*__obf_c7e5bacf18525a19)
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
func DoScan[T any](__obf_fd49918f96bc36c2 T, __obf_f9cbd174629de499 any) error {
	if __obf_f9cbd174629de499 == nil {
		return nil
	}
	var __obf_2274f6bd07b2425a []byte
	switch __obf_f9cbd174629de499 := __obf_f9cbd174629de499.(type) {
	case T:
		__obf_fd49918f96bc36c2 = __obf_f9cbd174629de499
		return nil
	case string:
		__obf_2274f6bd07b2425a = []byte(__obf_f9cbd174629de499)
	case []byte:
		__obf_2274f6bd07b2425a = __obf_f9cbd174629de499
	default:
		return fmt.Errorf("incompatible type for %T", __obf_f9cbd174629de499)
	}

	if len(__obf_2274f6bd07b2425a) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_2274f6bd07b2425a, &__obf_fd49918f96bc36c2)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_34511df5079d8a2e *JsonMap) Scan(__obf_9c99160fd6d9fa7d any) error {
	return DoScan(__obf_34511df5079d8a2e, __obf_9c99160fd6d9fa7d)
}

func (__obf_34511df5079d8a2e JsonMap) Value() (driver.Value, error) {
	if len(__obf_34511df5079d8a2e) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_34511df5079d8a2e)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_ae85d8c4a4b980aa *StringArray) Scan(__obf_9c99160fd6d9fa7d any) error {
	return DoScan(__obf_ae85d8c4a4b980aa,

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
		__obf_9c99160fd6d9fa7d)

	// return json.Unmarshal(bytes, &ja)
}

func (__obf_34511df5079d8a2e StringArray) Value() (driver.Value, error) {
	if len(__obf_34511df5079d8a2e) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_34511df5079d8a2e)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_34511df5079d8a2e *IntArray) Scan(__obf_9c99160fd6d9fa7d any) error {
	return DoScan(__obf_34511df5079d8a2e, __obf_9c99160fd6d9fa7d)
}

func (__obf_34511df5079d8a2e IntArray) Value() (driver.Value, error) {
	if len(__obf_34511df5079d8a2e) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_34511df5079d8a2e)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_34511df5079d8a2e *MapArray) Scan(__obf_9c99160fd6d9fa7d any) error {
	return DoScan(__obf_34511df5079d8a2e, __obf_9c99160fd6d9fa7d)
}

func (__obf_34511df5079d8a2e MapArray) Value() (driver.Value, error) {
	if len(__obf_34511df5079d8a2e) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_34511df5079d8a2e)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_34511df5079d8a2e *CaseTask) Scan(__obf_9c99160fd6d9fa7d any) error {
	return DoScan(__obf_34511df5079d8a2e, __obf_9c99160fd6d9fa7d)
}

func (__obf_34511df5079d8a2e CaseTask) Value() (driver.Value, error) {
	if len(__obf_34511df5079d8a2e) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_34511df5079d8a2e)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_34511df5079d8a2e *Timeline) Scan(__obf_9c99160fd6d9fa7d any) error {
	return DoScan(__obf_34511df5079d8a2e, __obf_9c99160fd6d9fa7d)
}

func (__obf_34511df5079d8a2e Timeline) Value() (driver.Value, error) {
	if len(__obf_34511df5079d8a2e) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_34511df5079d8a2e)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_34511df5079d8a2e Timeline) TimeLeft() int64 {
	__obf_ddc58e1af8598691 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_f1985ff415da2a02 := range __obf_34511df5079d8a2e {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_ddc58e1af8598691 >= __obf_f1985ff415da2a02.Start && __obf_ddc58e1af8598691 < __obf_f1985ff415da2a02.End {
			return __obf_f1985ff415da2a02.End - __obf_ddc58e1af8598691
		}
	}

	return 0
}

func (__obf_34511df5079d8a2e Timeline) IsUpcoming() bool {
	__obf_ddc58e1af8598691 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_f1985ff415da2a02 := range __obf_34511df5079d8a2e {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_f1985ff415da2a02.End > __obf_ddc58e1af8598691 {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_6176f198a7590139 RawJson) Message() json.RawMessage {
	if len(__obf_6176f198a7590139) == 0 {
		return nil
	}
	return json.RawMessage(__obf_6176f198a7590139)
}

// Scan implements the Scanner interface.
func (__obf_6176f198a7590139 *RawJson) Scan(__obf_9c99160fd6d9fa7d any) error {
	if __obf_9c99160fd6d9fa7d == nil {
		*__obf_6176f198a7590139 = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_75b9ed80d8791993, __obf_27df3b3b715e7e91 := __obf_9c99160fd6d9fa7d.([]byte)
	if !__obf_27df3b3b715e7e91 {
		__obf_2ed5ed7d7e423c28,
			// Handle other types if necessary, e.g., string
			__obf_27df3b3b715e7e91 := __obf_9c99160fd6d9fa7d.(string)
		if __obf_27df3b3b715e7e91 {
			__obf_75b9ed80d8791993 = []byte(__obf_2ed5ed7d7e423c28)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_9c99160fd6d9fa7d)
		}
	}
	*__obf_6176f198a7590139 = RawJson(__obf_75b9ed80d8791993)
	return nil
}

// Value implements the Valuer interface.
func (__obf_6176f198a7590139 RawJson) Value() (driver.Value, error) {
	if len(__obf_6176f198a7590139) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_6176f198a7590139), nil
}
