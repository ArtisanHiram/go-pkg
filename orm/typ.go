package __obf_52bf13aa40478808

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
func (__obf_37836fc6d365a20e *Str) Scan(__obf_6a3e0978a15ae7c4 any) error {
	if __obf_6a3e0978a15ae7c4 == nil {
		return nil
	}
	if __obf_bcc644fb72af6d31, __obf_b5d20856a702a4bb := __obf_6a3e0978a15ae7c4.([]byte); __obf_b5d20856a702a4bb {
		*__obf_37836fc6d365a20e = Str(__obf_bcc644fb72af6d31)
		return nil
	}
	return fmt.Errorf("unsupported type for String: %T", __obf_6a3e0978a15ae7c4)
}

func (__obf_37836fc6d365a20e *Str) Str() string {
	if __obf_37836fc6d365a20e == nil {
		return ""
	}
	return string(*__obf_37836fc6d365a20e)
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

func (__obf_935763ab96242d2f *Int) Scan(__obf_6a3e0978a15ae7c4 any) error {
	if __obf_6a3e0978a15ae7c4 == nil {
		return nil
	}

	switch __obf_c157396a957f0d01 := __obf_6a3e0978a15ae7c4.(type) {
	case []byte:
		__obf_1115ae3c80598454, _ := strconv.ParseInt(string(__obf_c157396a957f0d01), 10, 64)
		*__obf_935763ab96242d2f = Int(__obf_1115ae3c80598454)
	case int64:
		*__obf_935763ab96242d2f = Int(__obf_c157396a957f0d01)
	case uint64:
		*__obf_935763ab96242d2f = Int(__obf_c157396a957f0d01)
	default:
		return fmt.Errorf("unsupported type for Int: %T", __obf_6a3e0978a15ae7c4)
	}
	return nil
}

func (__obf_935763ab96242d2f *Int) Int() int64 {
	if __obf_935763ab96242d2f == nil {
		return 0
	}
	return int64(*__obf_935763ab96242d2f)
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

func (__obf_5a07a2990ad85783 *Float) Scan(__obf_6a3e0978a15ae7c4 any) error {
	if __obf_6a3e0978a15ae7c4 == nil {
		return nil
	}

	switch __obf_c157396a957f0d01 := __obf_6a3e0978a15ae7c4.(type) {
	case []byte:
		if __obf_37836fc6d365a20e, __obf_e6c24855347eb3d0 := strconv.ParseFloat(string(__obf_c157396a957f0d01), 64); __obf_e6c24855347eb3d0 != nil {
			return __obf_e6c24855347eb3d0
		} else {
			*__obf_5a07a2990ad85783 = Float(__obf_37836fc6d365a20e)
		}
	case float64:
		*__obf_5a07a2990ad85783 = Float(__obf_c157396a957f0d01)
	default:
		return fmt.Errorf("unsupported type for Float: %T", __obf_6a3e0978a15ae7c4)
	}

	return nil
}

func (__obf_935763ab96242d2f *Float) Float() float64 {
	if __obf_935763ab96242d2f == nil {
		return 0
	}
	return float64(*__obf_935763ab96242d2f)
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
// 		return fmt.Errorf("Unsupported type for Bool: %T", value)
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
func DoScan[T any](__obf_269eb4d75fdec1b8 T, __obf_1115ae3c80598454 any) error {
	if __obf_1115ae3c80598454 == nil {
		return nil
	}
	var __obf_55ba18dc3bcffccb []byte
	switch __obf_1115ae3c80598454 := __obf_1115ae3c80598454.(type) {
	case T:
		__obf_269eb4d75fdec1b8 = __obf_1115ae3c80598454
		return nil
	case string:
		__obf_55ba18dc3bcffccb = []byte(__obf_1115ae3c80598454)
	case []byte:
		__obf_55ba18dc3bcffccb = __obf_1115ae3c80598454
	default:
		return fmt.Errorf("incompatible type for %T", __obf_1115ae3c80598454)
	}

	if len(__obf_55ba18dc3bcffccb) == 0 {
		return nil
	}
	return json.Unmarshal(__obf_55ba18dc3bcffccb, &__obf_269eb4d75fdec1b8)
}

type JsonMap map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_c157396a957f0d01 *JsonMap) Scan(__obf_bcc644fb72af6d31 any) error {
	return DoScan(__obf_c157396a957f0d01, __obf_bcc644fb72af6d31)
}

func (__obf_c157396a957f0d01 JsonMap) Value() (driver.Value, error) {
	if len(__obf_c157396a957f0d01) == 0 {
		return []byte("{}"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_c157396a957f0d01)
}

type StringArray []string

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_9e71cf23e62be43c *StringArray) Scan(__obf_bcc644fb72af6d31 any) error {
	return DoScan(__obf_9e71cf23e62be43c, __obf_bcc644fb72af6d31)

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

func (__obf_c157396a957f0d01 StringArray) Value() (driver.Value, error) {
	if len(__obf_c157396a957f0d01) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_c157396a957f0d01)
}

type IntArray []int

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_c157396a957f0d01 *IntArray) Scan(__obf_bcc644fb72af6d31 any) error {
	return DoScan(__obf_c157396a957f0d01, __obf_bcc644fb72af6d31)
}

func (__obf_c157396a957f0d01 IntArray) Value() (driver.Value, error) {
	if len(__obf_c157396a957f0d01) == 0 {
		return []byte("[]"), nil // 空slice存为JSON空数组
	}
	return json.Marshal(__obf_c157396a957f0d01)
}

type MapArray []map[string]any

// Scan implements the sql.Scanner interface.
// It unmarshals the JSON data from the database into the JsonMap.
func (__obf_c157396a957f0d01 *MapArray) Scan(__obf_bcc644fb72af6d31 any) error {
	return DoScan(__obf_c157396a957f0d01, __obf_bcc644fb72af6d31)
}

func (__obf_c157396a957f0d01 MapArray) Value() (driver.Value, error) {
	if len(__obf_c157396a957f0d01) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_c157396a957f0d01)
}

type CaseTask []struct {
	Biz    int     `json:"biz,omitempty" db:"biz,omitempty"`
	Name   string  `json:"name" db:"name"`
	Post   string  `json:"post,omitempty" db:"post,omitempty"`
	Action string  `json:"action,omitempty" db:"action,omitempty"`
	Score  float32 `json:"score,omitempty" db:"score,omitempty"`
}

func (__obf_c157396a957f0d01 *CaseTask) Scan(__obf_bcc644fb72af6d31 any) error {
	return DoScan(__obf_c157396a957f0d01, __obf_bcc644fb72af6d31)
}

func (__obf_c157396a957f0d01 CaseTask) Value() (driver.Value, error) {
	if len(__obf_c157396a957f0d01) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_c157396a957f0d01)
}

type Timeline []struct {
	Start int64  `json:"start,omitempty" db:"start,omitempty"`
	End   int64  `json:"end,omitempty" db:"end,omitempty"`
	Type  string `json:"type,omitempty" db:"type,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
}

func (__obf_c157396a957f0d01 *Timeline) Scan(__obf_bcc644fb72af6d31 any) error {
	return DoScan(__obf_c157396a957f0d01, __obf_bcc644fb72af6d31)
}

func (__obf_c157396a957f0d01 Timeline) Value() (driver.Value, error) {
	if len(__obf_c157396a957f0d01) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(__obf_c157396a957f0d01)
}

// TimeLeft 当前场次考试剩余时间
func (__obf_c157396a957f0d01 Timeline) TimeLeft() int64 {
	__obf_53ff6beaba571cdd := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_595e74b93a7019df := range __obf_c157396a957f0d01 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_53ff6beaba571cdd >= __obf_595e74b93a7019df.Start && __obf_53ff6beaba571cdd < __obf_595e74b93a7019df.End {
			return __obf_595e74b93a7019df.End - __obf_53ff6beaba571cdd
		}
	}

	return 0
}

func (__obf_c157396a957f0d01 Timeline) IsUpcoming() bool {
	__obf_53ff6beaba571cdd := time.Now().Unix()
	// 遍历所有时间段
	for _, __obf_595e74b93a7019df := range __obf_c157396a957f0d01 {
		// 检查当前时间是否在某个时间段的开始和结束之间（包含开始和结束时间）
		if __obf_595e74b93a7019df.End > __obf_53ff6beaba571cdd {
			return true
		}
	}
	return false
}

type RawJson json.RawMessage

func (__obf_5eaa53a3e67e2a06 RawJson) Message() json.RawMessage {
	if len(__obf_5eaa53a3e67e2a06) == 0 {
		return nil
	}
	return json.RawMessage(__obf_5eaa53a3e67e2a06)
}

// Scan implements the Scanner interface.
func (__obf_5eaa53a3e67e2a06 *RawJson) Scan(__obf_bcc644fb72af6d31 any) error {
	if __obf_bcc644fb72af6d31 == nil {
		*__obf_5eaa53a3e67e2a06 = nil // Set to nil if the database value is NULL
		return nil
	}
	__obf_c856331b817a3b80, __obf_b5d20856a702a4bb := __obf_bcc644fb72af6d31.([]byte)
	if !__obf_b5d20856a702a4bb {
		// Handle other types if necessary, e.g., string
		__obf_335331934bfc4a58, __obf_b5d20856a702a4bb := __obf_bcc644fb72af6d31.(string)
		if __obf_b5d20856a702a4bb {
			__obf_c856331b817a3b80 = []byte(__obf_335331934bfc4a58)
		} else {
			return fmt.Errorf("RawJson.Scan: unsupported type %T", __obf_bcc644fb72af6d31)
		}
	}
	*__obf_5eaa53a3e67e2a06 = RawJson(__obf_c856331b817a3b80)
	return nil
}

// Value implements the Valuer interface.
func (__obf_5eaa53a3e67e2a06 RawJson) Value() (driver.Value, error) {
	if len(__obf_5eaa53a3e67e2a06) == 0 {
		return nil, nil // Return NULL for empty or nil RawMessage
	}
	return []byte(__obf_5eaa53a3e67e2a06), nil
}
