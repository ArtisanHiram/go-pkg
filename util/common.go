package __obf_d984cff8712b1ee6

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"math/rand/v2"
	"net"
	"path/filepath"

	"log"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"slices"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/nfnt/resize"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/yaml.v2"
)

var (
	Loc, _                 = time.LoadLocation("Asia/Shanghai")
	__obf_4f9564fbc13d6c68 = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_fff9dcb5e6c38578 any) ([]byte, error) {
	return __obf_4f9564fbc13d6c68.Marshal(__obf_fff9dcb5e6c38578)
}

func EncodeString(__obf_fff9dcb5e6c38578 any) string {
	if __obf_18cced84a245ac5a, __obf_44ebf351b5f3fef8 := __obf_4f9564fbc13d6c68.Marshal(__obf_fff9dcb5e6c38578); __obf_44ebf351b5f3fef8 == nil {
		return string(__obf_18cced84a245ac5a)
	}
	return ""
}

func Decode(__obf_5adf03aa4942297c string, __obf_fff9dcb5e6c38578 any) error {
	return __obf_4f9564fbc13d6c68.UnmarshalFromString(__obf_5adf03aa4942297c, __obf_fff9dcb5e6c38578)
}

func DecodeByte(__obf_f7b8f4634cd465ff []byte, __obf_fff9dcb5e6c38578 any) error {
	return __obf_4f9564fbc13d6c68.Unmarshal(__obf_f7b8f4634cd465ff, __obf_fff9dcb5e6c38578)
}

func StringUUID() string {

	return strings.ReplaceAll(uuid.New().String(), "-", "")

}

func IntUUID() uint32 {
	return uuid.New().ID()
}

// func InArray(val string, array []string) bool {
// 	for i := 0; i < len(array); i++ {
// 		// if strings.HasPrefix(val, array[i]) {
// 		if val == array[i] {
// 			return true
// 		}
// 	}
// 	return false
// }

func PrefixInArray(__obf_9c2e06b2d4c9e8a8 string, __obf_bc5450daf0fc5e5c []string) bool {
	for __obf_3682d6b25621f3a1 := range __obf_bc5450daf0fc5e5c {
		if strings.HasPrefix(__obf_9c2e06b2d4c9e8a8, __obf_bc5450daf0fc5e5c[__obf_3682d6b25621f3a1]) {
			return true
		}
	}
	return false
}

// func FindStrIndex(slice []string, val string) int {
// 	for i, item := range slice {
// 		if item == val {
// 			return i
// 		}
// 	}
// 	return -1
// }

//	func FindIntIndex(slice []int64, val int64) int {
//		for i, item := range slice {
//			if item == val {
//				return i
//			}
//		}
//		return -1
//	}

var __obf_a96fc69ab52c98e6 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenRandom 获得随机字符串
func GenRandom(__obf_1b64762461a96f61 int) string {
	__obf_f7b8f4634cd465ff := make([]byte, __obf_1b64762461a96f61)
	for __obf_3682d6b25621f3a1 := range __obf_f7b8f4634cd465ff {
		__obf_f7b8f4634cd465ff[__obf_3682d6b25621f3a1] = __obf_a96fc69ab52c98e6[rand.IntN(len(__obf_a96fc69ab52c98e6))]
	}
	return string(__obf_f7b8f4634cd465ff)
}

func RemoveIndex(__obf_1dbb0366e8ff201b []string, __obf_a7146f772dfd34d7 int) []string {
	return append(__obf_1dbb0366e8ff201b[:__obf_a7146f772dfd34d7], __obf_1dbb0366e8ff201b[__obf_a7146f772dfd34d7+1:]...)
}

// func FindIndex(slice []string, val string) int {
// 	for i, item := range slice {
// 		if item == val {
// 			return i
// 		}
// 	}
// 	return -1
// }

// 驼峰式
func ToCamel(__obf_5adf03aa4942297c string) (__obf_21e1d580f0acd4d1 string) {
	__obf_f130823467be9dd2 := []rune(__obf_5adf03aa4942297c)
	__obf_21e1d580f0acd4d1 = __obf_5adf03aa4942297c[0:1]
	if __obf_f130823467be9dd2[0] >= 97 && __obf_f130823467be9dd2[0] <= 122 {
		__obf_21e1d580f0acd4d1 = string(__obf_f130823467be9dd2[0] - 32)
	}

	__obf_18b51e009d13a45a := len(__obf_f130823467be9dd2)
	for __obf_3682d6b25621f3a1 := 1; __obf_3682d6b25621f3a1 < __obf_18b51e009d13a45a; __obf_3682d6b25621f3a1++ {
		if __obf_f130823467be9dd2[__obf_3682d6b25621f3a1] == 95 && __obf_f130823467be9dd2[__obf_3682d6b25621f3a1+1] >= 97 && __obf_f130823467be9dd2[__obf_3682d6b25621f3a1+1] <= 122 { //过滤下划线
			__obf_f130823467be9dd2[__obf_3682d6b25621f3a1+1] -= 32
		} else {
			__obf_21e1d580f0acd4d1 += string(__obf_f130823467be9dd2[__obf_3682d6b25621f3a1])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_5adf03aa4942297c string) (__obf_21e1d580f0acd4d1 string) {
	__obf_f130823467be9dd2 := []rune(__obf_5adf03aa4942297c)
	__obf_21e1d580f0acd4d1 = __obf_5adf03aa4942297c[0:1]
	if __obf_f130823467be9dd2[0] >= 65 && __obf_f130823467be9dd2[0] <= 90 {
		__obf_21e1d580f0acd4d1 = string(__obf_f130823467be9dd2[0] + 32)
	}

	__obf_1b64762461a96f61 := len(__obf_f130823467be9dd2)
	for __obf_3682d6b25621f3a1 := 1; __obf_3682d6b25621f3a1 < __obf_1b64762461a96f61; __obf_3682d6b25621f3a1++ {
		if __obf_f130823467be9dd2[__obf_3682d6b25621f3a1] >= 65 && __obf_f130823467be9dd2[__obf_3682d6b25621f3a1] <= 90 { //大写变小写
			__obf_f130823467be9dd2[__obf_3682d6b25621f3a1] += 32
			__obf_21e1d580f0acd4d1 += "_"
		}
		__obf_21e1d580f0acd4d1 += string(__obf_f130823467be9dd2[__obf_3682d6b25621f3a1])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_5adf03aa4942297c string, __obf_0ec31877d2939bc5 int, __obf_1b64762461a96f61 int) string {
	// 将字符串转换为rune切片，以正确处理多字节字符
	__obf_8691952f2125f68e := []rune(__obf_5adf03aa4942297c)
	__obf_910e5a981464d071 := len(__obf_8691952f2125f68e)

	// 处理n为负数或0的无效情况
	if __obf_1b64762461a96f61 <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_0ec31877d2939bc5 < 0 {
		__obf_0ec31877d2939bc5 = __obf_910e5a981464d071 + __obf_0ec31877d2939bc5
	}

	// 边界检查：确保start在有效范围内
	if __obf_0ec31877d2939bc5 < 0 || __obf_0ec31877d2939bc5 >= __obf_910e5a981464d071 {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_8691952f2125f68e[__obf_0ec31877d2939bc5:min(__obf_0ec31877d2939bc5+__obf_1b64762461a96f61, __obf_910e5a981464d071)])
}

func ASCII(__obf_d7f59788250a1a1e rune) rune {
	switch {
	case 97 <= __obf_d7f59788250a1a1e && __obf_d7f59788250a1a1e <= 122:
		return __obf_d7f59788250a1a1e - 32
	case 65 <= __obf_d7f59788250a1a1e && __obf_d7f59788250a1a1e <= 90:
		return __obf_d7f59788250a1a1e + 32
	default:
		return __obf_d7f59788250a1a1e
	}
}

func IndexString(__obf_5adf03aa4942297c string, __obf_88ce94801855d048 rune, __obf_c42dfd5cec13c5da int) string {
	__obf_8cb70b99086f6db0 := []rune(__obf_5adf03aa4942297c)
	var __obf_3e61a5d40e7b8f7f bytes.Buffer
	var __obf_1b64762461a96f61 int
	for __obf_3682d6b25621f3a1, __obf_d886e0eb58bccdd7 := 0, len(__obf_8cb70b99086f6db0); __obf_3682d6b25621f3a1 < __obf_d886e0eb58bccdd7; __obf_3682d6b25621f3a1++ {
		if __obf_8cb70b99086f6db0[__obf_3682d6b25621f3a1] == __obf_88ce94801855d048 {
			__obf_1b64762461a96f61 += 1
		}
		if __obf_1b64762461a96f61 == __obf_c42dfd5cec13c5da {
			break
		}
		__obf_3e61a5d40e7b8f7f.WriteRune(__obf_8cb70b99086f6db0[__obf_3682d6b25621f3a1])
	}
	return __obf_3e61a5d40e7b8f7f.String()
}

func LastIndexString(__obf_43053eac9d2321a5, __obf_ef6990ef7d8b1eac string) string {
	__obf_1dbb0366e8ff201b := strings.Split(__obf_43053eac9d2321a5, __obf_ef6990ef7d8b1eac)
	if __obf_1b64762461a96f61 := len(__obf_1dbb0366e8ff201b); __obf_1b64762461a96f61 > 1 {
		return __obf_1dbb0366e8ff201b[__obf_1b64762461a96f61-2]
	}
	return ""
}

func IsEmpty(__obf_2d5bbd84b5709790 any) bool {
	__obf_5ca54f845825fcc4 := reflect.ValueOf(__obf_2d5bbd84b5709790)
	if __obf_5ca54f845825fcc4.Kind() == reflect.Ptr {
		__obf_5ca54f845825fcc4 = __obf_5ca54f845825fcc4.Elem()
	}
	return __obf_5ca54f845825fcc4.Interface() == reflect.Zero(__obf_5ca54f845825fcc4.Type()).Interface()
}

func MillisecondToDateString(__obf_df8eface920851bd int64) string {
	return time.Unix(__obf_df8eface920851bd, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_df8eface920851bd int64) string {
	return time.Unix(__obf_df8eface920851bd, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_94dd7c2a55e630bc string) (__obf_0a5084207611bf1f, __obf_19447c1a68711a24 string) {
	__obf_94dd7c2a55e630bc = filepath.Base(__obf_94dd7c2a55e630bc)
	__obf_19447c1a68711a24 = filepath.Ext(__obf_94dd7c2a55e630bc)
	__obf_0a5084207611bf1f = strings.TrimSuffix(__obf_94dd7c2a55e630bc, __obf_19447c1a68711a24)
	return __obf_0a5084207611bf1f, __obf_19447c1a68711a24
}

// func ListMap(rows *orm.Rows, call func(map[string]any) (string, string)) (result []map[string]any) {
// 	for rows.Next() {
// 		tmp := make(map[string]any)
// 		rows.MapScan(tmp)
// 		for k, encoded := range tmp {
// 			switch encoded.(type) {
// 			case []byte:
// 				tmp[k] = string(encoded.([]byte))
// 			}
// 		}
// 		if call != nil {
// 			key, res := call(tmp)
// 			tmp[key] = res
// 		}
// 		result = append(result, tmp)
// 	}
// 	return
// }

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(__obf_4c17d1c34384fde6 string) (int64, error) {
	if __obf_710d20992ab66874, __obf_44ebf351b5f3fef8 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_4c17d1c34384fde6, Loc); __obf_44ebf351b5f3fef8 != nil {
		return 0, nil
	} else {
		__obf_710d20992ab66874 = __obf_710d20992ab66874.AddDate(0, 0, -__obf_710d20992ab66874.Day()+1)
		return GetZeroTime(__obf_710d20992ab66874).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_4c17d1c34384fde6 string) (int64, int64, error) {
	if __obf_710d20992ab66874, __obf_44ebf351b5f3fef8 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_4c17d1c34384fde6, Loc); __obf_44ebf351b5f3fef8 != nil {
		return 0, 0, nil
	} else {
		__obf_710d20992ab66874 = __obf_710d20992ab66874.AddDate(0, 0, -__obf_710d20992ab66874.Day()+1)
		return GetZeroTime(__obf_710d20992ab66874).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_710d20992ab66874).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_968eb1e1c0ee776f time.Time) time.Time {
	return time.Date(__obf_968eb1e1c0ee776f.Year(), __obf_968eb1e1c0ee776f.Month(), __obf_968eb1e1c0ee776f.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_df8eface920851bd int64) int64 {
	__obf_c3555481aee3fa53 := time.Unix(__obf_df8eface920851bd, 0)
	__obf_9b629708071c5ab1, __obf_678669a0a436a3a1, __obf_968eb1e1c0ee776f := __obf_c3555481aee3fa53.Date()
	return time.Date(__obf_9b629708071c5ab1, __obf_678669a0a436a3a1, __obf_968eb1e1c0ee776f, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_df8eface920851bd int64) int64 {
	__obf_c3555481aee3fa53 := time.Unix(__obf_df8eface920851bd, 0)
	return __obf_c3555481aee3fa53.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_9b629708071c5ab1, __obf_678669a0a436a3a1, __obf_968eb1e1c0ee776f := time.Now().Date()
	return time.Date(__obf_9b629708071c5ab1, __obf_678669a0a436a3a1, __obf_968eb1e1c0ee776f, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_4c17d1c34384fde6 string) (int64, int64) {
	__obf_4105fb6d4f8fcfc1, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_4c17d1c34384fde6+" 00:00:00", Loc)
	min := __obf_4105fb6d4f8fcfc1.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_11d5635b909d6faf int) (int64, int64) {
	__obf_4c17d1c34384fde6 := time.Now()
	__obf_c3555481aee3fa53 := __obf_4c17d1c34384fde6.AddDate(0, 0, __obf_11d5635b909d6faf)
	__obf_9b629708071c5ab1, __obf_678669a0a436a3a1, __obf_968eb1e1c0ee776f := __obf_c3555481aee3fa53.Date()
	min := time.Date(__obf_9b629708071c5ab1, __obf_678669a0a436a3a1, __obf_968eb1e1c0ee776f, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_4c17d1c34384fde6.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_2f2b19d1a1ae8ca2, __obf_1b0eb151c5924058 string) bool {
	__obf_c4c84e4a662f4335 := func(__obf_1dbb0366e8ff201b string) string {
		return fmt.Sprintf(",%s,", __obf_1dbb0366e8ff201b)
	}
	return strings.Contains(__obf_c4c84e4a662f4335(__obf_2f2b19d1a1ae8ca2), __obf_c4c84e4a662f4335(__obf_1b0eb151c5924058))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_b854c3d4cccac277 string, __obf_b12dc4479c319c8d ...string) bool {
	if any {
		for _, __obf_1a08474ebb9e961a := range __obf_b12dc4479c319c8d {
			if Contain(__obf_b854c3d4cccac277, __obf_1a08474ebb9e961a) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_1a08474ebb9e961a := range __obf_b12dc4479c319c8d {
			if !Contain(__obf_b854c3d4cccac277, __obf_1a08474ebb9e961a) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_1dbb0366e8ff201b []any, __obf_a7146f772dfd34d7 int) []any {
	return slices.Delete(__obf_1dbb0366e8ff201b, __obf_a7146f772dfd34d7, __obf_a7146f772dfd34d7+1)
}

func String2Int8(__obf_5adf03aa4942297c string) int8 {
	__obf_3144b29142edab41, __obf_44ebf351b5f3fef8 := strconv.ParseInt(__obf_5adf03aa4942297c, 10, 8)
	if __obf_44ebf351b5f3fef8 == nil {
		return int8(__obf_3144b29142edab41)
	}
	return 0
}

func String2Int32(__obf_5adf03aa4942297c string) int32 {
	__obf_3144b29142edab41, __obf_44ebf351b5f3fef8 := strconv.ParseInt(__obf_5adf03aa4942297c, 10, 32)
	if __obf_44ebf351b5f3fef8 == nil {
		return int32(__obf_3144b29142edab41)
	}
	return 0
}

func String2Int64(__obf_5adf03aa4942297c string) int8 {
	__obf_3144b29142edab41, __obf_44ebf351b5f3fef8 := strconv.ParseInt(__obf_5adf03aa4942297c, 10, 64)
	if __obf_44ebf351b5f3fef8 == nil {
		return int8(__obf_3144b29142edab41)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_da3bcec2717f0afb, __obf_8a03e0b735b13a49 int) (__obf_0ec31877d2939bc5, __obf_f091b7d51b0a4f61 time.Time) {
	if __obf_da3bcec2717f0afb > 0 && __obf_8a03e0b735b13a49 > 0 {
		__obf_0ec31877d2939bc5 = time.Date(__obf_da3bcec2717f0afb, 1, 0, 0, 0, 0, 0, Loc)
		// 第一天是周几
		__obf_22e1f58a20aca581 := int(__obf_0ec31877d2939bc5.AddDate(0, 0, 1).Weekday())
		// 当年第一周有几天
		__obf_43e02d16fe8c24a4 := 1
		if __obf_22e1f58a20aca581 != 0 {
			__obf_43e02d16fe8c24a4 = 7 - __obf_22e1f58a20aca581 + 1
		}
		if __obf_8a03e0b735b13a49 == 1 {
			__obf_f091b7d51b0a4f61 = __obf_0ec31877d2939bc5.AddDate(0, 0, __obf_43e02d16fe8c24a4)
		} else {
			__obf_f091b7d51b0a4f61 = __obf_0ec31877d2939bc5.AddDate(0, 0, __obf_43e02d16fe8c24a4+(__obf_8a03e0b735b13a49-1)*7)
			__obf_0ec31877d2939bc5 = __obf_f091b7d51b0a4f61.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_4c17d1c34384fde6 time.Time) (__obf_da3bcec2717f0afb, __obf_8a03e0b735b13a49, __obf_3f18ac9798a3ba17 int) {
	__obf_da3bcec2717f0afb = __obf_4c17d1c34384fde6.Year()
	__obf_3f18ac9798a3ba17 = int(__obf_4c17d1c34384fde6.Weekday())
	__obf_a465890916fcc88b := __obf_4c17d1c34384fde6.YearDay()
	__obf_2864397921be5be7 := __obf_4c17d1c34384fde6.AddDate(0, 0, -__obf_a465890916fcc88b+1)
	__obf_22e1f58a20aca581 := int(__obf_2864397921be5be7.Weekday())
	// 当年第一周有几天
	__obf_43e02d16fe8c24a4 := 1
	if __obf_22e1f58a20aca581 != 0 {
		__obf_43e02d16fe8c24a4 = 7 - __obf_22e1f58a20aca581 + 1
	}
	if __obf_a465890916fcc88b <= __obf_43e02d16fe8c24a4 {
		__obf_8a03e0b735b13a49 = 1
	} else {
		__obf_8a03e0b735b13a49 = (__obf_a465890916fcc88b-__obf_43e02d16fe8c24a4)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_d3d682d97f2d0435 []byte) map[string]string {
	__obf_179b07ce95dad5f2 := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_5e0284ee0f24efaf := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_3f70185f70481cd8 []byte
	if __obf_3f70185f70481cd8 = __obf_179b07ce95dad5f2.ReplaceAll(__obf_d3d682d97f2d0435, []byte("")); len(__obf_3f70185f70481cd8) < 6 {
		return nil
	}

	__obf_85c708c9f4c300f3 := __obf_5e0284ee0f24efaf.FindAllSubmatch(__obf_3f70185f70481cd8[6:len(__obf_3f70185f70481cd8)-7], -1)
	__obf_fbb4247cebecffb7 := map[string]string{}
	if __obf_c8503a247a8a4d14 := __obf_179b07ce95dad5f2.FindSubmatch(__obf_d3d682d97f2d0435)[1]; len(__obf_c8503a247a8a4d14) != 0 {
		__obf_df2ea973c799028e := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_c8503a247a8a4d14, -1)
		for _, __obf_5ca54f845825fcc4 := range __obf_df2ea973c799028e {
			__obf_fbb4247cebecffb7[string(__obf_5ca54f845825fcc4[1])] = string(__obf_5ca54f845825fcc4[2])
		}
	}

	for _, __obf_5ca54f845825fcc4 := range __obf_85c708c9f4c300f3 {
		__obf_fbb4247cebecffb7[string(__obf_5ca54f845825fcc4[1])] = string(__obf_5ca54f845825fcc4[2])
	}
	return __obf_fbb4247cebecffb7
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_423972d9be0c283a time.Time) (int64, int64) {
	__obf_f82890c826838dd5 := GetZeroTime(__obf_423972d9be0c283a).Unix() / 600
	return __obf_f82890c826838dd5, __obf_f82890c826838dd5 + 143
}

func Abs(__obf_1b64762461a96f61 int64) int64 {
	__obf_9b629708071c5ab1 := __obf_1b64762461a96f61 >> 63
	return (__obf_1b64762461a96f61 ^ __obf_9b629708071c5ab1) - __obf_9b629708071c5ab1
}

func Number2String(__obf_1b64762461a96f61 any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_1b64762461a96f61 := __obf_1b64762461a96f61.(type) {
	case int:
		return strconv.Itoa(__obf_1b64762461a96f61)
	case int32:
		return strconv.FormatInt(int64(__obf_1b64762461a96f61), 10)
	case int64:
		return strconv.FormatInt(__obf_1b64762461a96f61, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_1b64762461a96f61), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_1b64762461a96f61, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_5adf03aa4942297c any, __obf_88ce94801855d048 string) string {
	if __obf_5adf03aa4942297c != nil && __obf_5adf03aa4942297c.(string) != "" {
		// return sep + str + sep
		return __obf_5adf03aa4942297c.(string)
	}
	return __obf_88ce94801855d048
}

func SortRange(__obf_678669a0a436a3a1 map[string]any, __obf_4621ced8fc91df0c func(int, string)) {
	var __obf_44acc96b7f70fef6 []string
	for __obf_7bff7adf95417505 := range __obf_678669a0a436a3a1 {
		__obf_44acc96b7f70fef6 = append(__obf_44acc96b7f70fef6, __obf_7bff7adf95417505)
	}
	sort.Strings(__obf_44acc96b7f70fef6)
	for __obf_3682d6b25621f3a1, __obf_7bff7adf95417505 := range __obf_44acc96b7f70fef6 {
		__obf_4621ced8fc91df0c(__obf_3682d6b25621f3a1, __obf_7bff7adf95417505)
	}
}

func HasField(__obf_d85aaabfaa1dbbe1 reflect.Value, __obf_0a5084207611bf1f string) bool {

	if __obf_1dbb0366e8ff201b := __obf_d85aaabfaa1dbbe1.FieldByNameFunc(func(__obf_1b64762461a96f61 string) bool {
		return strings.EqualFold(__obf_1b64762461a96f61, __obf_0a5084207611bf1f)
	}); __obf_1dbb0366e8ff201b.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_d85aaabfaa1dbbe1 reflect.Value, __obf_0a5084207611bf1f string) reflect.Value {
	if __obf_1dbb0366e8ff201b := __obf_d85aaabfaa1dbbe1.FieldByNameFunc(func(__obf_1b64762461a96f61 string) bool {
		return strings.EqualFold(__obf_1b64762461a96f61, __obf_0a5084207611bf1f)
	}); __obf_1dbb0366e8ff201b.IsValid() {
		return __obf_1dbb0366e8ff201b
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_d85aaabfaa1dbbe1 reflect.Value, __obf_0a5084207611bf1f string, __obf_9c2e06b2d4c9e8a8 any) bool {

	// if st.Kind() == reflect.Ptr {
	// 	st = st.Elem()
	// 	st = reflect.Indirect(st)
	// 	if s := st.FieldByNameFunc(func(n string) bool { return strings.EqualFold(n, name) }); s.IsValid() {
	// 		if stype := s.Type(); stype == reflect.TypeOf(val) {
	// 			s.Set(reflect.ValueOf(val))
	// 		} else {
	// 			s.Set(reflect.ValueOf(val).Convert(stype))
	// 		}
	// 		return true
	// 	}
	// }
	if __obf_1dbb0366e8ff201b := __obf_d85aaabfaa1dbbe1.FieldByNameFunc(func(__obf_1b64762461a96f61 string) bool {
		return strings.EqualFold(__obf_1b64762461a96f61, __obf_0a5084207611bf1f)
	}); __obf_1dbb0366e8ff201b.IsValid() {
		if __obf_3574d01a2eb0f1a5 := __obf_1dbb0366e8ff201b.Type(); __obf_3574d01a2eb0f1a5 == reflect.TypeOf(__obf_9c2e06b2d4c9e8a8) {
			__obf_1dbb0366e8ff201b.Set(reflect.ValueOf(__obf_9c2e06b2d4c9e8a8))
		} else {
			__obf_1dbb0366e8ff201b.Set(reflect.ValueOf(__obf_9c2e06b2d4c9e8a8).Convert(__obf_3574d01a2eb0f1a5))
		}
		return true
	}
	return false
}

func CopyMap(__obf_678669a0a436a3a1 map[string]any) map[string]any {
	__obf_df2ea973c799028e := make(map[string]any)
	for __obf_7bff7adf95417505, __obf_5ca54f845825fcc4 := range __obf_678669a0a436a3a1 {
		if __obf_6da269bfb5fb6725, __obf_78b5927d79871a32 := __obf_5ca54f845825fcc4.(map[string]any); __obf_78b5927d79871a32 {
			__obf_df2ea973c799028e[__obf_7bff7adf95417505] = CopyMap(__obf_6da269bfb5fb6725)
		} else {
			__obf_df2ea973c799028e[__obf_7bff7adf95417505] = __obf_5ca54f845825fcc4
		}
	}

	return __obf_df2ea973c799028e
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_1c173f4352ac0b13, __obf_bd708ec7bc22375a string, __obf_73a842bdad45bd70 bool, __obf_dcd616eb94d8d77d ...any) {
	var __obf_701cc0b946f4427a strings.Builder
	var __obf_3f38da3ab9450507 string
	var __obf_5ed673cfc1f74f0e []string
	for _, __obf_bc4b9d7b176483de := range __obf_dcd616eb94d8d77d {
		__obf_1dbb0366e8ff201b := reflect.TypeOf(__obf_bc4b9d7b176483de)
		__obf_701cc0b946f4427a.WriteString(`CREATE TABLE `)
		__obf_701cc0b946f4427a.WriteString(__obf_1dbb0366e8ff201b.Name())
		__obf_701cc0b946f4427a.WriteString(" (\n")
		__obf_8e33fd09eeede116 := __obf_1dbb0366e8ff201b.NumField()
		for __obf_3682d6b25621f3a1 := 0; __obf_3682d6b25621f3a1 < __obf_8e33fd09eeede116; __obf_3682d6b25621f3a1++ {
			__obf_701cc0b946f4427a.WriteString("    ")
			__obf_5ed673cfc1f74f0e = nil
			if __obf_15ac000269ca7ff4 := string(__obf_1dbb0366e8ff201b.Field(__obf_3682d6b25621f3a1).Tag.Get(__obf_1c173f4352ac0b13)); __obf_15ac000269ca7ff4 == "" {
				if __obf_73a842bdad45bd70 {
					__obf_3f38da3ab9450507 = ToCamel(__obf_1dbb0366e8ff201b.Field(__obf_3682d6b25621f3a1).Name)
				} else {
					__obf_3f38da3ab9450507 = __obf_1dbb0366e8ff201b.Field(__obf_3682d6b25621f3a1).Name
				}
			} else {
				__obf_5ed673cfc1f74f0e = strings.Split(__obf_15ac000269ca7ff4, __obf_bd708ec7bc22375a)
				__obf_3f38da3ab9450507 = __obf_5ed673cfc1f74f0e[0]
			}
			__obf_701cc0b946f4427a.WriteString(__obf_3f38da3ab9450507)
			__obf_701cc0b946f4427a.WriteString(" ")
			switch __obf_1dbb0366e8ff201b.Field(__obf_3682d6b25621f3a1).Type.Name() {
			case "int8":
				__obf_701cc0b946f4427a.WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_701cc0b946f4427a.WriteString("INT")
			case "int64":
				__obf_701cc0b946f4427a.WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_701cc0b946f4427a.WriteString("VARCHAR(50)")
			}

			if len(__obf_5ed673cfc1f74f0e) > 1 {
				__obf_701cc0b946f4427a.WriteString(" ")
				__obf_701cc0b946f4427a.WriteString(strings.Join(__obf_5ed673cfc1f74f0e[1:], " "))
			}

			if __obf_3682d6b25621f3a1+1 != __obf_8e33fd09eeede116 {
				__obf_701cc0b946f4427a.WriteString(",")
			}
			__obf_701cc0b946f4427a.WriteString("\n")
		}
		__obf_701cc0b946f4427a.WriteString(");\n\n")
	}
	fmt.Println(__obf_701cc0b946f4427a.String())
}

func SaveImage(__obf_f02c63132b94ef63 image.Image, __obf_d9f9e056cb642245 uint, __obf_85a9e60e4e37ed46 string) error {
	__obf_2fd8ae35bd667e7f, __obf_44ebf351b5f3fef8 := os.Create(__obf_85a9e60e4e37ed46)
	if __obf_44ebf351b5f3fef8 != nil {
		return __obf_44ebf351b5f3fef8
	}
	defer __obf_2fd8ae35bd667e7f.Close()
	return jpeg.Encode(__obf_2fd8ae35bd667e7f, resize.Resize(__obf_d9f9e056cb642245, 0, __obf_f02c63132b94ef63, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_5adf03aa4942297c string) bool {
	return strings.Contains(__obf_5adf03aa4942297c, ".") && net.ParseIP(__obf_5adf03aa4942297c) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_5adf03aa4942297c string) bool {
	return strings.Contains(__obf_5adf03aa4942297c, ":") && net.ParseIP(__obf_5adf03aa4942297c) != nil
}

func AnyToBytes(__obf_5ca54f845825fcc4 any) ([]byte, error) {
	return msgpack.Marshal(__obf_5ca54f845825fcc4)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_9c2e06b2d4c9e8a8 []byte) (__obf_fff9dcb5e6c38578 T, __obf_44ebf351b5f3fef8 error) {
	__obf_44ebf351b5f3fef8 = msgpack.Unmarshal(__obf_9c2e06b2d4c9e8a8, &__obf_fff9dcb5e6c38578)
	return
}

func Loadyaml(__obf_b8ed48d9d038b3a0 string, __obf_447647fb43e6bcba any) {
	__obf_1b0eb151c5924058, __obf_44ebf351b5f3fef8 := os.ReadFile(__obf_b8ed48d9d038b3a0)
	if __obf_44ebf351b5f3fef8 != nil {
		log.Fatalln(__obf_44ebf351b5f3fef8)
	}
	__obf_44ebf351b5f3fef8 = yaml.UnmarshalStrict(__obf_1b0eb151c5924058, __obf_447647fb43e6bcba)
	if __obf_44ebf351b5f3fef8 != nil {
		log.Fatalln(__obf_44ebf351b5f3fef8)
	}
}

func ToAnyList[T any](__obf_75398679a8e79c45 []T) []any {
	__obf_9b80fe0bc429c869 := make([]any, len(__obf_75398679a8e79c45))
	for __obf_3682d6b25621f3a1, __obf_5ca54f845825fcc4 := range __obf_75398679a8e79c45 {
		__obf_9b80fe0bc429c869[__obf_3682d6b25621f3a1] = __obf_5ca54f845825fcc4
	}
	return __obf_9b80fe0bc429c869
}

func TimeParse(__obf_169b98a6a29b09c7 string) time.Time {
	__obf_169b98a6a29b09c7 = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_169b98a6a29b09c7)
	__obf_c3555481aee3fa53, _ := time.ParseInLocation("2006-01-02 15:04", __obf_169b98a6a29b09c7, time.Local)
	return __obf_c3555481aee3fa53
}
