package __obf_24a21e7f5375b619

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
func (__obf_75b913ceae9096c5 *Str) Scan(__obf_9493dc7eb40a414f any) error {
	if __obf_9493dc7eb40a414f == nil {
		return nil
	}

	switch __obf_4224178bd2feb4fe := __obf_9493dc7eb40a414f.(type) {
	case []byte:
		*__obf_75b913ceae9096c5 = Str(__obf_4224178bd2feb4fe)
	case string:
		*__obf_75b913ceae9096c5 = Str(__obf_4224178bd2feb4fe)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_9493dc7eb40a414f)
	}
	return nil
}

func (__obf_75b913ceae9096c5 *Str) Str() string {
	if __obf_75b913ceae9096c5 == nil {
		return ""
	}
	return string(*__obf_75b913ceae9096c5)
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

func (__obf_f4c5dc7b1f70e707 *Int) Scan(__obf_9493dc7eb40a414f any) error {
	if __obf_9493dc7eb40a414f == nil {
		return nil
	}

	switch __obf_4224178bd2feb4fe := __obf_9493dc7eb40a414f.(type) {
	case []byte:
		__obf_9e8e83e850d24f6c, _ := strconv.ParseInt(string(__obf_4224178bd2feb4fe), 10, 64)
		*__obf_f4c5dc7b1f70e707 = Int(__obf_9e8e83e850d24f6c)
	case int64:
		*__obf_f4c5dc7b1f70e707 = Int(__obf_4224178bd2feb4fe)
	case uint64:
		*__obf_f4c5dc7b1f70e707 = Int(__obf_4224178bd2feb4fe)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_9493dc7eb40a414f)
	}
	return nil
}

func (__obf_f4c5dc7b1f70e707 *Int) Int() int64 {
	if __obf_f4c5dc7b1f70e707 == nil {
		return 0
	}
	return int64(*__obf_f4c5dc7b1f70e707)
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

func (__obf_57b980f12a7a4617 *Float) Scan(__obf_9493dc7eb40a414f any) error {
	if __obf_9493dc7eb40a414f == nil {
		return nil
	}

	switch __obf_4224178bd2feb4fe := __obf_9493dc7eb40a414f.(type) {
	case []byte:
		if __obf_75b913ceae9096c5, __obf_a687171eb2d5acc2 := strconv.ParseFloat(string(__obf_4224178bd2feb4fe), 64); __obf_a687171eb2d5acc2 != nil {
			return __obf_a687171eb2d5acc2
		} else {
			*__obf_57b980f12a7a4617 = Float(__obf_75b913ceae9096c5)
		}
	case float64:
		*__obf_57b980f12a7a4617 = Float(__obf_4224178bd2feb4fe)
	default:
		return fmt.Errorf("unsupported type for: %T", __obf_9493dc7eb40a414f)
	}

	return nil
}

func (__obf_f4c5dc7b1f70e707 *Float) Float() float64 {
	if __obf_f4c5dc7b1f70e707 == nil {
		return 0
	}
	return float64(*__obf_f4c5dc7b1f70e707)
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
func DoScan[T any](__obf_2dcd500b1e634a96 T, __obf_9e8e83e850d24f6c any) error {
	if __obf_9e8e83e850d24f6c == nil {
		return nil
	}
	var __obf_fa6f408520f47863 []byte
	switch __obf_9e8e83e850d24f6c := __obf_9e8e83e850d24f6c.(type) {
	case T:
		__obf_2dcd500b1e634a96 = __obf_9e8e83e850d24f6c
		return nil
	case string:
		__obf_fa6f408520f47863 = []byte(__obf_9e8e83e850d24f6c)
	case []byte:
		__obf_fa6f408520f47863 = __obf_9e8e83e850d24f6c
	default:
		return fmt.Errorf("incompatible type for %T", __obf_9e8e83e850d24f6c)
	}

	if len(__obf_fa6f408520f47863) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_fa6f408520f47863, &__obf_2dcd500b1e634a96)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_4224178bd2feb4fe *JsonMap) Scan(__obf_0826efe2d8c95bbc any) error {
	return DoScan(__obf_4224178bd2feb4fe, __obf_0826efe2d8c95bbc)
}

func (__obf_4224178bd2feb4fe JsonMap) Value() (driver.Value, error) {
	if len(__obf_4224178bd2feb4fe) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_4224178bd2feb4fe)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_373636a8a19c2ec4 *StringArray) Scan(__obf_0826efe2d8c95bbc any) error {
	return DoScan(__obf_373636a8a19c2ec4, __obf_0826efe2d8c95bbc)

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

func (__obf_4224178bd2feb4fe StringArray) Value() (driver.Value, error) {
	if len(__obf_4224178bd2feb4fe) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_4224178bd2feb4fe)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_4224178bd2feb4fe *IntArray) Scan(__obf_0826efe2d8c95bbc any) error {
	return DoScan(__obf_4224178bd2feb4fe, __obf_0826efe2d8c95bbc)
}

func (__obf_4224178bd2feb4fe IntArray) Value() (driver.Value, error) {
	if len(__obf_4224178bd2feb4fe) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_4224178bd2feb4fe)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_4224178bd2feb4fe *MapArray) Scan(__obf_0826efe2d8c95bbc any) error {
	return DoScan(__obf_4224178bd2feb4fe, __obf_0826efe2d8c95bbc)
}

func (__obf_4224178bd2feb4fe MapArray) Value() (driver.Value, error) {
	if len(__obf_4224178bd2feb4fe) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_4224178bd2feb4fe)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_4224178bd2feb4fe *CaseTask) Scan(__obf_0826efe2d8c95bbc any) error {
	return DoScan(__obf_4224178bd2feb4fe, __obf_0826efe2d8c95bbc)
}

func (__obf_4224178bd2feb4fe CaseTask) Value() (driver.Value, error) {
	if len(__obf_4224178bd2feb4fe) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_4224178bd2feb4fe)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_4224178bd2feb4fe *Timeline) Scan(__obf_0826efe2d8c95bbc any) error {
	return DoScan(__obf_4224178bd2feb4fe, __obf_0826efe2d8c95bbc)
}

func (__obf_4224178bd2feb4fe Timeline) Value() (driver.Value, error) {
	if len(__obf_4224178bd2feb4fe) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_4224178bd2feb4fe)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_4224178bd2feb4fe Timeline) TimeLeft() int64 {
	__obf_dc3f2f3951ec09dd := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_9e960b23bd5e3da8 := range __obf_4224178bd2feb4fe {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_dc3f2f3951ec09dd >= __obf_9e960b23bd5e3da8.Start && __obf_dc3f2f3951ec09dd < __obf_9e960b23bd5e3da8.End {
			return __obf_9e960b23bd5e3da8.End - __obf_dc3f2f3951ec09dd
		}
	}

	return 0
}

func (__obf_4224178bd2feb4fe Timeline) IsUpcoming() bool {
	__obf_dc3f2f3951ec09dd := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_9e960b23bd5e3da8 := range __obf_4224178bd2feb4fe {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_9e960b23bd5e3da8.End > __obf_dc3f2f3951ec09dd {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_9df0c4a0a0b4bef1 RawJson) Message() json.RawMessage {
	if len(__obf_9df0c4a0a0b4bef1) == 0 {
		return nil
	}
	return json.RawMessage(__obf_9df0c4a0a0b4bef1)
}

// Scan implements the Scanner interface.
func (__obf_9df0c4a0a0b4bef1 *RawJson) Scan(__obf_0826efe2d8c95bbc any) error {
	if __obf_0826efe2d8c95bbc == nil {
		*__obf_9df0c4a0a0b4bef1 = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_08d236129dbb0df2, __obf_bfcf4076c7714374 := __obf_0826efe2d8c95bbc.([]byte)
	if !__obf_bfcf4076c7714374 {
		// Handle other types if necessary, e.g., string
		__obf_d2b22995c880514e, __obf_bfcf4076c7714374 := __obf_0826efe2d8c95bbc.(string)
		if __obf_bfcf4076c7714374 {
			__obf_08d236129dbb0df2 = []byte(__obf_d2b22995c880514e)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_0826efe2d8c95bbc)
		}
	}
	*__obf_9df0c4a0a0b4bef1 = RawJson(__obf_08d236129dbb0df2)
	return nil
}

// Value implements the Valuer interface.
func (__obf_9df0c4a0a0b4bef1 RawJson) Value() (driver.Value, error) {
	if len(__obf_9df0c4a0a0b4bef1) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_9df0c4a0a0b4bef1), nil
}
