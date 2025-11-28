package __obf_66da9ed7fbc1ec9e

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
func (__obf_a3ec10d2be1cb29c *Str) Scan(__obf_d175223e413b5539 any) error {
	if __obf_d175223e413b5539 == nil {
		return nil
	}

	switch __obf_7d1378dcb9668016 := __obf_d175223e413b5539.(type) {
	case []byte:
		*__obf_a3ec10d2be1cb29c = Str(__obf_7d1378dcb9668016)
	case string:
		*__obf_a3ec10d2be1cb29c = Str(__obf_7d1378dcb9668016)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_d175223e413b5539)
	}
	return nil
}

func (__obf_a3ec10d2be1cb29c *Str) Str() string {
	if __obf_a3ec10d2be1cb29c == nil {
		return ""
	}
	return string(*__obf_a3ec10d2be1cb29c)
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

func (__obf_4ceb93679168ba44 *Int) Scan(__obf_d175223e413b5539 any) error {
	if __obf_d175223e413b5539 == nil {
		return nil
	}

	switch __obf_7d1378dcb9668016 := __obf_d175223e413b5539.(type) {
	case []byte:
		__obf_bc4d09a84501c296, _ := strconv.ParseInt(string(__obf_7d1378dcb9668016), 10, 64)
		*__obf_4ceb93679168ba44 = Int(__obf_bc4d09a84501c296)
	case int64:
		*__obf_4ceb93679168ba44 = Int(__obf_7d1378dcb9668016)
	case uint64:
		*__obf_4ceb93679168ba44 = Int(__obf_7d1378dcb9668016)
	default:
		return fmt.Errorf("unsupported type for %T", __obf_d175223e413b5539)
	}
	return nil
}

func (__obf_4ceb93679168ba44 *Int) Int() int64 {
	if __obf_4ceb93679168ba44 == nil {
		return 0
	}
	return int64(*__obf_4ceb93679168ba44)
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

func (__obf_bb16074987cbe782 *Float) Scan(__obf_d175223e413b5539 any) error {
	if __obf_d175223e413b5539 == nil {
		return nil
	}

	switch __obf_7d1378dcb9668016 := __obf_d175223e413b5539.(type) {
	case []byte:
		if __obf_a3ec10d2be1cb29c, __obf_1229ad25944902c2 := strconv.ParseFloat(string(__obf_7d1378dcb9668016), 64); __obf_1229ad25944902c2 != nil {
			return __obf_1229ad25944902c2
		} else {
			*__obf_bb16074987cbe782 = Float(__obf_a3ec10d2be1cb29c)
		}
	case float64:
		*__obf_bb16074987cbe782 = Float(__obf_7d1378dcb9668016)
	default:
		return fmt.Errorf("unsupported type for: %T", __obf_d175223e413b5539)
	}

	return nil
}

func (__obf_4ceb93679168ba44 *Float) Float() float64 {
	if __obf_4ceb93679168ba44 == nil {
		return 0
	}
	return float64(*__obf_4ceb93679168ba44)
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
func DoScan[T any](__obf_41a9e3ce7b95a912 T, __obf_bc4d09a84501c296 any) error {
	if __obf_bc4d09a84501c296 == nil {
		return nil
	}
	var __obf_c4c751dcd43e50c7 []byte
	switch __obf_bc4d09a84501c296 := __obf_bc4d09a84501c296.(type) {
	case T:
		__obf_41a9e3ce7b95a912 = __obf_bc4d09a84501c296
		return nil
	case string:
		__obf_c4c751dcd43e50c7 = []byte(__obf_bc4d09a84501c296)
	case []byte:
		__obf_c4c751dcd43e50c7 = __obf_bc4d09a84501c296
	default:
		return fmt.Errorf("incompatible type for %T", __obf_bc4d09a84501c296)
	}

	if len(__obf_c4c751dcd43e50c7) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_c4c751dcd43e50c7, &__obf_41a9e3ce7b95a912)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_7d1378dcb9668016 *JsonMap) Scan(__obf_830db7a4c207dec8 any) error {
	return DoScan(__obf_7d1378dcb9668016, __obf_830db7a4c207dec8)
}

func (__obf_7d1378dcb9668016 JsonMap) Value() (driver.Value, error) {
	if len(__obf_7d1378dcb9668016) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_7d1378dcb9668016)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_af84c77f65a8b954 *StringArray) Scan(__obf_830db7a4c207dec8 any) error {
	return DoScan(__obf_af84c77f65a8b954, __obf_830db7a4c207dec8)

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

func (__obf_7d1378dcb9668016 StringArray) Value() (driver.Value, error) {
	if len(__obf_7d1378dcb9668016) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_7d1378dcb9668016)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_7d1378dcb9668016 *IntArray) Scan(__obf_830db7a4c207dec8 any) error {
	return DoScan(__obf_7d1378dcb9668016, __obf_830db7a4c207dec8)
}

func (__obf_7d1378dcb9668016 IntArray) Value() (driver.Value, error) {
	if len(__obf_7d1378dcb9668016) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_7d1378dcb9668016)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_7d1378dcb9668016 *MapArray) Scan(__obf_830db7a4c207dec8 any) error {
	return DoScan(__obf_7d1378dcb9668016, __obf_830db7a4c207dec8)
}

func (__obf_7d1378dcb9668016 MapArray) Value() (driver.Value, error) {
	if len(__obf_7d1378dcb9668016) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_7d1378dcb9668016)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_7d1378dcb9668016 *CaseTask) Scan(__obf_830db7a4c207dec8 any) error {
	return DoScan(__obf_7d1378dcb9668016, __obf_830db7a4c207dec8)
}

func (__obf_7d1378dcb9668016 CaseTask) Value() (driver.Value, error) {
	if len(__obf_7d1378dcb9668016) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_7d1378dcb9668016)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_7d1378dcb9668016 *Timeline) Scan(__obf_830db7a4c207dec8 any) error {
	return DoScan(__obf_7d1378dcb9668016, __obf_830db7a4c207dec8)
}

func (__obf_7d1378dcb9668016 Timeline) Value() (driver.Value, error) {
	if len(__obf_7d1378dcb9668016) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_7d1378dcb9668016)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_7d1378dcb9668016 Timeline) TimeLeft() int64 {
	__obf_c1c73aeb1df55de2 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_bc568ad8d15e6e37 := range __obf_7d1378dcb9668016 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_c1c73aeb1df55de2 >= __obf_bc568ad8d15e6e37.Start && __obf_c1c73aeb1df55de2 < __obf_bc568ad8d15e6e37.End {
			return __obf_bc568ad8d15e6e37.End - __obf_c1c73aeb1df55de2
		}
	}

	return 0
}

func (__obf_7d1378dcb9668016 Timeline) IsUpcoming() bool {
	__obf_c1c73aeb1df55de2 := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_bc568ad8d15e6e37 := range __obf_7d1378dcb9668016 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_bc568ad8d15e6e37.End > __obf_c1c73aeb1df55de2 {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_37e0648504b076b2 RawJson) Message() json.RawMessage {
	if len(__obf_37e0648504b076b2) == 0 {
		return nil
	}
	return json.RawMessage(__obf_37e0648504b076b2)
}

// Scan implements the Scanner interface.
func (__obf_37e0648504b076b2 *RawJson) Scan(__obf_830db7a4c207dec8 any) error {
	if __obf_830db7a4c207dec8 == nil {
		*__obf_37e0648504b076b2 = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_254b8879ad8d5437, __obf_184ad66d1ce62d27 := __obf_830db7a4c207dec8.([]byte)
	if !__obf_184ad66d1ce62d27 {
		// Handle other types if necessary, e.g., string
		__obf_1ddafd4e3fc11361, __obf_184ad66d1ce62d27 := __obf_830db7a4c207dec8.(string)
		if __obf_184ad66d1ce62d27 {
			__obf_254b8879ad8d5437 = []byte(__obf_1ddafd4e3fc11361)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_830db7a4c207dec8)
		}
	}
	*__obf_37e0648504b076b2 = RawJson(__obf_254b8879ad8d5437)
	return nil
}

// Value implements the Valuer interface.
func (__obf_37e0648504b076b2 RawJson) Value() (driver.Value, error) {
	if len(__obf_37e0648504b076b2) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_37e0648504b076b2), nil
}
