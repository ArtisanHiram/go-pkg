package __obf_b81118ac905f398e

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

	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	msgpack "github.com/ArtisanHiram/go-pkg/msgpack"
	resize "github.com/ArtisanHiram/go-pkg/resize"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
	"slices"
)

var (
	Loc, _                 = time.LoadLocation("Asia/Shanghai")
	__obf_da4a17fce5fcc04c = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_703a701ac3bc0aab any) ([]byte, error) {
	return __obf_da4a17fce5fcc04c.Marshal(__obf_703a701ac3bc0aab)
}

func EncodeString(__obf_703a701ac3bc0aab any) string {
	if __obf_099caa5a26f40901, __obf_65d3bb18bf1bd15f := __obf_da4a17fce5fcc04c.Marshal(__obf_703a701ac3bc0aab); __obf_65d3bb18bf1bd15f == nil {
		return string(__obf_099caa5a26f40901)
	}
	return ""
}

func Decode(__obf_37d7f27c5f97b901 string, __obf_703a701ac3bc0aab any) error {
	return __obf_da4a17fce5fcc04c.UnmarshalFromString(__obf_37d7f27c5f97b901, __obf_703a701ac3bc0aab)
}

func DecodeByte(__obf_f3483784ccf89cd0 []byte, __obf_703a701ac3bc0aab any) error {
	return __obf_da4a17fce5fcc04c.Unmarshal(__obf_f3483784ccf89cd0, __obf_703a701ac3bc0aab)
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

func PrefixInArray(__obf_263270b555073815 string, __obf_365edeac8cf477e0 []string) bool {
	for __obf_2482e14692f1453a := range __obf_365edeac8cf477e0 {
		if strings.HasPrefix(__obf_263270b555073815, __obf_365edeac8cf477e0[__obf_2482e14692f1453a]) {
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

var __obf_ae348b0e62e51cf3 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenRandom 获得随机字符串
func GenRandom(__obf_cfef5ea6a21678e8 int) string {
	__obf_f3483784ccf89cd0 := make([]byte, __obf_cfef5ea6a21678e8)
	for __obf_2482e14692f1453a := range __obf_f3483784ccf89cd0 {
		__obf_f3483784ccf89cd0[__obf_2482e14692f1453a] = __obf_ae348b0e62e51cf3[rand.IntN(len(__obf_ae348b0e62e51cf3))]
	}
	return string(__obf_f3483784ccf89cd0)
}

func RemoveIndex(__obf_c0c7a6993b194cda []string, __obf_8050ca2da4eede19 int) []string {
	return append(__obf_c0c7a6993b194cda[:__obf_8050ca2da4eede19], __obf_c0c7a6993b194cda[__obf_8050ca2da4eede19+1:]...)
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
func ToCamel(__obf_37d7f27c5f97b901 string) (__obf_e8db7e51ba1ec27d string) {
	__obf_41e51334b498bc59 := []rune(__obf_37d7f27c5f97b901)
	__obf_e8db7e51ba1ec27d = __obf_37d7f27c5f97b901[0:1]
	if __obf_41e51334b498bc59[0] >= 97 && __obf_41e51334b498bc59[0] <= 122 {
		__obf_e8db7e51ba1ec27d = string(__obf_41e51334b498bc59[0] - 32)
	}
	__obf_6b8c74141f887409 := len(__obf_41e51334b498bc59)
	for __obf_2482e14692f1453a := 1; __obf_2482e14692f1453a < __obf_6b8c74141f887409; __obf_2482e14692f1453a++ {
		if __obf_41e51334b498bc59[__obf_2482e14692f1453a] == 95 && __obf_41e51334b498bc59[__obf_2482e14692f1453a+1] >= 97 && __obf_41e51334b498bc59[__obf_2482e14692f1453a+1] <= 122 {
			__obf_41e51334b498bc59[ //过滤下划线
			__obf_2482e14692f1453a+1] -= 32
		} else {
			__obf_e8db7e51ba1ec27d += string(__obf_41e51334b498bc59[__obf_2482e14692f1453a])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_37d7f27c5f97b901 string) (__obf_e8db7e51ba1ec27d string) {
	__obf_41e51334b498bc59 := []rune(__obf_37d7f27c5f97b901)
	__obf_e8db7e51ba1ec27d = __obf_37d7f27c5f97b901[0:1]
	if __obf_41e51334b498bc59[0] >= 65 && __obf_41e51334b498bc59[0] <= 90 {
		__obf_e8db7e51ba1ec27d = string(__obf_41e51334b498bc59[0] + 32)
	}
	__obf_cfef5ea6a21678e8 := len(__obf_41e51334b498bc59)
	for __obf_2482e14692f1453a := 1; __obf_2482e14692f1453a < __obf_cfef5ea6a21678e8; __obf_2482e14692f1453a++ {
		if __obf_41e51334b498bc59[__obf_2482e14692f1453a] >= 65 && __obf_41e51334b498bc59[__obf_2482e14692f1453a] <= 90 {
			__obf_41e51334b498bc59[ //大写变小写
			__obf_2482e14692f1453a] += 32
			__obf_e8db7e51ba1ec27d += "_"
		}
		__obf_e8db7e51ba1ec27d += string(__obf_41e51334b498bc59[__obf_2482e14692f1453a])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_37d7f27c5f97b901 string, __obf_c62fc6dd77a64e94 int, __obf_cfef5ea6a21678e8 int) string {
	__obf_4a96ffe2ef2965a8 := // 将字符串转换为rune切片，以正确处理多字节字符
		[]rune(__obf_37d7f27c5f97b901)
	__obf_e4668c42490030e3 := len(__obf_4a96ffe2ef2965a8)

	// 处理n为负数或0的无效情况
	if __obf_cfef5ea6a21678e8 <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_c62fc6dd77a64e94 < 0 {
		__obf_c62fc6dd77a64e94 = __obf_e4668c42490030e3 + __obf_c62fc6dd77a64e94
	}

	// 边界检查：确保start在有效范围内
	if __obf_c62fc6dd77a64e94 < 0 || __obf_c62fc6dd77a64e94 >= __obf_e4668c42490030e3 {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_4a96ffe2ef2965a8[__obf_c62fc6dd77a64e94:min(__obf_c62fc6dd77a64e94+__obf_cfef5ea6a21678e8, __obf_e4668c42490030e3)])
}

func ASCII(__obf_8e8b497458f714e7 rune) rune {
	switch {
	case 97 <= __obf_8e8b497458f714e7 && __obf_8e8b497458f714e7 <= 122:
		return __obf_8e8b497458f714e7 - 32
	case 65 <= __obf_8e8b497458f714e7 && __obf_8e8b497458f714e7 <= 90:
		return __obf_8e8b497458f714e7 + 32
	default:
		return __obf_8e8b497458f714e7
	}
}

func IndexString(__obf_37d7f27c5f97b901 string, __obf_001d80c65dc4d93c rune, __obf_8cabfcfa02b67c38 int) string {
	__obf_ce270cc9a231dcec := []rune(__obf_37d7f27c5f97b901)
	var __obf_818532ea20c7a33b bytes.Buffer
	var __obf_cfef5ea6a21678e8 int
	for __obf_2482e14692f1453a, __obf_53f4d293022844f8 := 0, len(__obf_ce270cc9a231dcec); __obf_2482e14692f1453a < __obf_53f4d293022844f8; __obf_2482e14692f1453a++ {
		if __obf_ce270cc9a231dcec[__obf_2482e14692f1453a] == __obf_001d80c65dc4d93c {
			__obf_cfef5ea6a21678e8 += 1
		}
		if __obf_cfef5ea6a21678e8 == __obf_8cabfcfa02b67c38 {
			break
		}
		__obf_818532ea20c7a33b.
			WriteRune(__obf_ce270cc9a231dcec[__obf_2482e14692f1453a])
	}
	return __obf_818532ea20c7a33b.String()
}

func LastIndexString(__obf_16b08a5a1d6464d0, __obf_f7192c976570e9b9 string) string {
	__obf_c0c7a6993b194cda := strings.Split(__obf_16b08a5a1d6464d0, __obf_f7192c976570e9b9)
	if __obf_cfef5ea6a21678e8 := len(__obf_c0c7a6993b194cda); __obf_cfef5ea6a21678e8 > 1 {
		return __obf_c0c7a6993b194cda[__obf_cfef5ea6a21678e8-2]
	}
	return ""
}

func IsEmpty(__obf_2029f950af19a455 any) bool {
	__obf_194e82db8912d873 := reflect.ValueOf(__obf_2029f950af19a455)
	if __obf_194e82db8912d873.Kind() == reflect.Ptr {
		__obf_194e82db8912d873 = __obf_194e82db8912d873.Elem()
	}
	return __obf_194e82db8912d873.Interface() == reflect.Zero(__obf_194e82db8912d873.Type()).Interface()
}

func MillisecondToDateString(__obf_d73de86ebb26b023 int64) string {
	return time.Unix(__obf_d73de86ebb26b023, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_d73de86ebb26b023 int64) string {
	return time.Unix(__obf_d73de86ebb26b023, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_d122bd101d831be6 string) (__obf_783842bb1ca67ba6, __obf_9b2075572417bb43 string) {
	__obf_d122bd101d831be6 = filepath.Base(__obf_d122bd101d831be6)
	__obf_9b2075572417bb43 = filepath.Ext(__obf_d122bd101d831be6)
	__obf_783842bb1ca67ba6 = strings.TrimSuffix(__obf_d122bd101d831be6, __obf_9b2075572417bb43)
	return __obf_783842bb1ca67ba6,

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
		__obf_9b2075572417bb43
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(__obf_bba455645ee33c38 string) (int64, error) {
	if __obf_b3495adf9a275ab6, __obf_65d3bb18bf1bd15f := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_bba455645ee33c38, Loc); __obf_65d3bb18bf1bd15f != nil {
		return 0, nil
	} else {
		__obf_b3495adf9a275ab6 = __obf_b3495adf9a275ab6.AddDate(0, 0, -__obf_b3495adf9a275ab6.Day()+1)
		return GetZeroTime(__obf_b3495adf9a275ab6).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_bba455645ee33c38 string) (int64, int64, error) {
	if __obf_b3495adf9a275ab6, __obf_65d3bb18bf1bd15f := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_bba455645ee33c38, Loc); __obf_65d3bb18bf1bd15f != nil {
		return 0, 0, nil
	} else {
		__obf_b3495adf9a275ab6 = __obf_b3495adf9a275ab6.AddDate(0, 0, -__obf_b3495adf9a275ab6.Day()+1)
		return GetZeroTime(__obf_b3495adf9a275ab6).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_b3495adf9a275ab6).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_1bd7c4bbab843f7b time.Time) time.Time {
	return time.Date(__obf_1bd7c4bbab843f7b.Year(), __obf_1bd7c4bbab843f7b.Month(), __obf_1bd7c4bbab843f7b.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_d73de86ebb26b023 int64) int64 {
	__obf_d881f8a0ff5d4e10 := time.Unix(__obf_d73de86ebb26b023, 0)
	__obf_f217601aafabfccb, __obf_8d5a9085886e6ba0, __obf_1bd7c4bbab843f7b := __obf_d881f8a0ff5d4e10.Date()
	return time.Date(__obf_f217601aafabfccb, __obf_8d5a9085886e6ba0, __obf_1bd7c4bbab843f7b, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_d73de86ebb26b023 int64) int64 {
	__obf_d881f8a0ff5d4e10 := time.Unix(__obf_d73de86ebb26b023, 0)
	return __obf_d881f8a0ff5d4e10.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_f217601aafabfccb, __obf_8d5a9085886e6ba0, __obf_1bd7c4bbab843f7b := time.Now().Date()
	return time.Date(__obf_f217601aafabfccb, __obf_8d5a9085886e6ba0, __obf_1bd7c4bbab843f7b, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_bba455645ee33c38 string) (int64, int64) {
	__obf_c26cd28f49ffaa09, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_bba455645ee33c38+" 00:00:00", Loc)
	min := __obf_c26cd28f49ffaa09.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_f84dafda7509b1b6 int) (int64, int64) {
	__obf_bba455645ee33c38 := time.Now()
	__obf_d881f8a0ff5d4e10 := __obf_bba455645ee33c38.AddDate(0, 0, __obf_f84dafda7509b1b6)
	__obf_f217601aafabfccb, __obf_8d5a9085886e6ba0, __obf_1bd7c4bbab843f7b := __obf_d881f8a0ff5d4e10.Date()
	min := time.Date(__obf_f217601aafabfccb, __obf_8d5a9085886e6ba0, __obf_1bd7c4bbab843f7b, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_bba455645ee33c38.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_655b7c7207361c09, __obf_00e3488b89a95768 string) bool {
	__obf_8708c8ce2b934c26 := func(__obf_c0c7a6993b194cda string) string {
		return fmt.Sprintf(",%s,", __obf_c0c7a6993b194cda)
	}
	return strings.Contains(__obf_8708c8ce2b934c26(__obf_655b7c7207361c09), __obf_8708c8ce2b934c26(__obf_00e3488b89a95768))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_d65ec76415fc3c57 string, __obf_e1c7f7959d083409 ...string) bool {
	if any {
		for _, __obf_c9c59dd2cdb7bc0f := range __obf_e1c7f7959d083409 {
			if Contain(__obf_d65ec76415fc3c57, __obf_c9c59dd2cdb7bc0f) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_c9c59dd2cdb7bc0f := range __obf_e1c7f7959d083409 {
			if !Contain(__obf_d65ec76415fc3c57, __obf_c9c59dd2cdb7bc0f) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_c0c7a6993b194cda []any, __obf_8050ca2da4eede19 int) []any {
	return slices.Delete(__obf_c0c7a6993b194cda, __obf_8050ca2da4eede19, __obf_8050ca2da4eede19+1)
}

func String2Int8(__obf_37d7f27c5f97b901 string) int8 {
	__obf_a57f1c550f889d64, __obf_65d3bb18bf1bd15f := strconv.ParseInt(__obf_37d7f27c5f97b901, 10, 8)
	if __obf_65d3bb18bf1bd15f == nil {
		return int8(__obf_a57f1c550f889d64)
	}
	return 0
}

func String2Int32(__obf_37d7f27c5f97b901 string) int32 {
	__obf_a57f1c550f889d64, __obf_65d3bb18bf1bd15f := strconv.ParseInt(__obf_37d7f27c5f97b901, 10, 32)
	if __obf_65d3bb18bf1bd15f == nil {
		return int32(__obf_a57f1c550f889d64)
	}
	return 0
}

func String2Int64(__obf_37d7f27c5f97b901 string) int8 {
	__obf_a57f1c550f889d64, __obf_65d3bb18bf1bd15f := strconv.ParseInt(__obf_37d7f27c5f97b901, 10, 64)
	if __obf_65d3bb18bf1bd15f == nil {
		return int8(__obf_a57f1c550f889d64)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_54113281f7ba56e0, __obf_df4bb7e6b4e8ecfe int) (__obf_c62fc6dd77a64e94, __obf_493963029ded18fd time.Time) {
	if __obf_54113281f7ba56e0 > 0 && __obf_df4bb7e6b4e8ecfe > 0 {
		__obf_c62fc6dd77a64e94 = time.Date(__obf_54113281f7ba56e0, 1, 0, 0, 0, 0, 0, Loc)
		__obf_5f1ce3a2c9feed10 := // 第一天是周几
			int(__obf_c62fc6dd77a64e94.AddDate(0, 0, 1).Weekday())
		__obf_958ac8c7d809ed9b := // 当年第一周有几天
			1
		if __obf_5f1ce3a2c9feed10 != 0 {
			__obf_958ac8c7d809ed9b = 7 - __obf_5f1ce3a2c9feed10 + 1
		}
		if __obf_df4bb7e6b4e8ecfe == 1 {
			__obf_493963029ded18fd = __obf_c62fc6dd77a64e94.AddDate(0, 0, __obf_958ac8c7d809ed9b)
		} else {
			__obf_493963029ded18fd = __obf_c62fc6dd77a64e94.AddDate(0, 0, __obf_958ac8c7d809ed9b+(__obf_df4bb7e6b4e8ecfe-1)*7)
			__obf_c62fc6dd77a64e94 = __obf_493963029ded18fd.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_bba455645ee33c38 time.Time) (__obf_54113281f7ba56e0, __obf_df4bb7e6b4e8ecfe, __obf_758e60f9234123d6 int) {
	__obf_54113281f7ba56e0 = __obf_bba455645ee33c38.Year()
	__obf_758e60f9234123d6 = int(__obf_bba455645ee33c38.Weekday())
	__obf_56f3f176de53f41d := __obf_bba455645ee33c38.YearDay()
	__obf_b117471e7b55bf76 := __obf_bba455645ee33c38.AddDate(0, 0, -__obf_56f3f176de53f41d+1)
	__obf_5f1ce3a2c9feed10 := int(__obf_b117471e7b55bf76.Weekday())
	__obf_958ac8c7d809ed9b := // 当年第一周有几天
		1
	if __obf_5f1ce3a2c9feed10 != 0 {
		__obf_958ac8c7d809ed9b = 7 - __obf_5f1ce3a2c9feed10 + 1
	}
	if __obf_56f3f176de53f41d <= __obf_958ac8c7d809ed9b {
		__obf_df4bb7e6b4e8ecfe = 1
	} else {
		__obf_df4bb7e6b4e8ecfe = (__obf_56f3f176de53f41d-__obf_958ac8c7d809ed9b)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_9699bbcf8255c645 []byte) map[string]string {
	__obf_c8f7bcceb3dd3f2e := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_c8a4b10f42bc2688 := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_e9e92b3bba50740f []byte
	if __obf_e9e92b3bba50740f = __obf_c8f7bcceb3dd3f2e.ReplaceAll(__obf_9699bbcf8255c645, []byte("")); len(__obf_e9e92b3bba50740f) < 6 {
		return nil
	}
	__obf_b377fc7201569f7c := __obf_c8a4b10f42bc2688.FindAllSubmatch(__obf_e9e92b3bba50740f[6:len(__obf_e9e92b3bba50740f)-7], -1)
	__obf_f7e1171900eebc57 := map[string]string{}
	if __obf_d3ea6d146f5c2bf6 := __obf_c8f7bcceb3dd3f2e.FindSubmatch(__obf_9699bbcf8255c645)[1]; len(__obf_d3ea6d146f5c2bf6) != 0 {
		__obf_6a80f9b3275ed3f4 := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_d3ea6d146f5c2bf6, -1)
		for _, __obf_194e82db8912d873 := range __obf_6a80f9b3275ed3f4 {
			__obf_f7e1171900eebc57[string(__obf_194e82db8912d873[1])] = string(__obf_194e82db8912d873[2])
		}
	}

	for _, __obf_194e82db8912d873 := range __obf_b377fc7201569f7c {
		__obf_f7e1171900eebc57[string(__obf_194e82db8912d873[1])] = string(__obf_194e82db8912d873[2])
	}
	return __obf_f7e1171900eebc57
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_90233fb629c2318b time.Time) (int64, int64) {
	__obf_97402de7d80215b7 := GetZeroTime(__obf_90233fb629c2318b).Unix() / 600
	return __obf_97402de7d80215b7, __obf_97402de7d80215b7 + 143
}

func Abs(__obf_cfef5ea6a21678e8 int64) int64 {
	__obf_f217601aafabfccb := __obf_cfef5ea6a21678e8 >> 63
	return (__obf_cfef5ea6a21678e8 ^ __obf_f217601aafabfccb) - __obf_f217601aafabfccb
}

func Number2String(__obf_cfef5ea6a21678e8 any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_cfef5ea6a21678e8 := __obf_cfef5ea6a21678e8.(type) {
	case int:
		return strconv.Itoa(__obf_cfef5ea6a21678e8)
	case int32:
		return strconv.FormatInt(int64(__obf_cfef5ea6a21678e8), 10)
	case int64:
		return strconv.FormatInt(__obf_cfef5ea6a21678e8, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_cfef5ea6a21678e8), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_cfef5ea6a21678e8, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_37d7f27c5f97b901 any, __obf_001d80c65dc4d93c string) string {
	if __obf_37d7f27c5f97b901 != nil && __obf_37d7f27c5f97b901.(string) != "" {
		// return sep + str + sep
		return __obf_37d7f27c5f97b901.(string)
	}
	return __obf_001d80c65dc4d93c
}

func SortRange(__obf_8d5a9085886e6ba0 map[string]any, __obf_919144fdd1f21680 func(int, string)) {
	var __obf_aa28447075c8b10d []string
	for __obf_b29ccd2992d5240a := range __obf_8d5a9085886e6ba0 {
		__obf_aa28447075c8b10d = append(__obf_aa28447075c8b10d, __obf_b29ccd2992d5240a)
	}
	sort.Strings(__obf_aa28447075c8b10d)
	for __obf_2482e14692f1453a, __obf_b29ccd2992d5240a := range __obf_aa28447075c8b10d {
		__obf_919144fdd1f21680(__obf_2482e14692f1453a, __obf_b29ccd2992d5240a)
	}
}

func HasField(__obf_9b7828bb7f6f66d1 reflect.Value, __obf_783842bb1ca67ba6 string) bool {

	if __obf_c0c7a6993b194cda := __obf_9b7828bb7f6f66d1.FieldByNameFunc(func(__obf_cfef5ea6a21678e8 string) bool {
		return strings.EqualFold(__obf_cfef5ea6a21678e8, __obf_783842bb1ca67ba6)
	}); __obf_c0c7a6993b194cda.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_9b7828bb7f6f66d1 reflect.Value, __obf_783842bb1ca67ba6 string) reflect.Value {
	if __obf_c0c7a6993b194cda := __obf_9b7828bb7f6f66d1.FieldByNameFunc(func(__obf_cfef5ea6a21678e8 string) bool {
		return strings.EqualFold(__obf_cfef5ea6a21678e8, __obf_783842bb1ca67ba6)
	}); __obf_c0c7a6993b194cda.IsValid() {
		return __obf_c0c7a6993b194cda
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_9b7828bb7f6f66d1 reflect.Value, __obf_783842bb1ca67ba6 string, __obf_263270b555073815 any) bool {

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
	if __obf_c0c7a6993b194cda := __obf_9b7828bb7f6f66d1.FieldByNameFunc(func(__obf_cfef5ea6a21678e8 string) bool {
		return strings.EqualFold(__obf_cfef5ea6a21678e8, __obf_783842bb1ca67ba6)
	}); __obf_c0c7a6993b194cda.IsValid() {
		if __obf_fd890a00033ef84e := __obf_c0c7a6993b194cda.Type(); __obf_fd890a00033ef84e == reflect.TypeOf(__obf_263270b555073815) {
			__obf_c0c7a6993b194cda.
				Set(reflect.ValueOf(__obf_263270b555073815))
		} else {
			__obf_c0c7a6993b194cda.
				Set(reflect.ValueOf(__obf_263270b555073815).Convert(__obf_fd890a00033ef84e))
		}
		return true
	}
	return false
}

func CopyMap(__obf_8d5a9085886e6ba0 map[string]any) map[string]any {
	__obf_6a80f9b3275ed3f4 := make(map[string]any)
	for __obf_b29ccd2992d5240a, __obf_194e82db8912d873 := range __obf_8d5a9085886e6ba0 {
		if __obf_fd6002e099f2cb9e, __obf_b87c1b9200311a15 := __obf_194e82db8912d873.(map[string]any); __obf_b87c1b9200311a15 {
			__obf_6a80f9b3275ed3f4[__obf_b29ccd2992d5240a] = CopyMap(__obf_fd6002e099f2cb9e)
		} else {
			__obf_6a80f9b3275ed3f4[__obf_b29ccd2992d5240a] = __obf_194e82db8912d873
		}
	}

	return __obf_6a80f9b3275ed3f4
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_826deb59f8429f3a, __obf_54903ba9cbfd194d string, __obf_0431fdd11337483e bool, __obf_d3bca3fcf260e8a4 ...any) {
	var __obf_30869df9454700d0 strings.Builder
	var __obf_ed9c6c15c2447101 string
	var __obf_70cb6882cccf43fe []string
	for _, __obf_2592cb4e1c2b1d14 := range __obf_d3bca3fcf260e8a4 {
		__obf_c0c7a6993b194cda := reflect.TypeOf(__obf_2592cb4e1c2b1d14)
		__obf_30869df9454700d0.
			WriteString(`CREATE TABLE `)
		__obf_30869df9454700d0.
			WriteString(__obf_c0c7a6993b194cda.Name())
		__obf_30869df9454700d0.
			WriteString(" (\n")
		__obf_910124f43434be9f := __obf_c0c7a6993b194cda.NumField()
		for __obf_2482e14692f1453a := 0; __obf_2482e14692f1453a < __obf_910124f43434be9f; __obf_2482e14692f1453a++ {
			__obf_30869df9454700d0.
				WriteString("    ")
			__obf_70cb6882cccf43fe = nil
			if __obf_d2b7bcd6d96d3ada := string(__obf_c0c7a6993b194cda.Field(__obf_2482e14692f1453a).Tag.Get(__obf_826deb59f8429f3a)); __obf_d2b7bcd6d96d3ada == "" {
				if __obf_0431fdd11337483e {
					__obf_ed9c6c15c2447101 = ToCamel(__obf_c0c7a6993b194cda.Field(__obf_2482e14692f1453a).Name)
				} else {
					__obf_ed9c6c15c2447101 = __obf_c0c7a6993b194cda.Field(__obf_2482e14692f1453a).Name
				}
			} else {
				__obf_70cb6882cccf43fe = strings.Split(__obf_d2b7bcd6d96d3ada, __obf_54903ba9cbfd194d)
				__obf_ed9c6c15c2447101 = __obf_70cb6882cccf43fe[0]
			}
			__obf_30869df9454700d0.
				WriteString(__obf_ed9c6c15c2447101)
			__obf_30869df9454700d0.
				WriteString(" ")
			switch __obf_c0c7a6993b194cda.Field(__obf_2482e14692f1453a).Type.Name() {
			case "int8":
				__obf_30869df9454700d0.
					WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_30869df9454700d0.
					WriteString("INT")
			case "int64":
				__obf_30869df9454700d0.
					WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_30869df9454700d0.
					WriteString("VARCHAR(50)")
			}

			if len(__obf_70cb6882cccf43fe) > 1 {
				__obf_30869df9454700d0.
					WriteString(" ")
				__obf_30869df9454700d0.
					WriteString(strings.Join(__obf_70cb6882cccf43fe[1:], " "))
			}

			if __obf_2482e14692f1453a+1 != __obf_910124f43434be9f {
				__obf_30869df9454700d0.
					WriteString(",")
			}
			__obf_30869df9454700d0.
				WriteString("\n")
		}
		__obf_30869df9454700d0.
			WriteString(");\n\n")
	}
	fmt.Println(__obf_30869df9454700d0.String())
}

func SaveImage(__obf_8f6e8f558dcdfd32 image.Image, __obf_ff723ca022862d6a uint, __obf_82cdac0270bd6a95 string) error {
	__obf_e68b24230f2d353e, __obf_65d3bb18bf1bd15f := os.Create(__obf_82cdac0270bd6a95)
	if __obf_65d3bb18bf1bd15f != nil {
		return __obf_65d3bb18bf1bd15f
	}
	defer __obf_e68b24230f2d353e.Close()
	return jpeg.Encode(__obf_e68b24230f2d353e, resize.Resize(__obf_ff723ca022862d6a, 0, __obf_8f6e8f558dcdfd32, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_37d7f27c5f97b901 string) bool {
	return strings.Contains(__obf_37d7f27c5f97b901, ".") && net.ParseIP(__obf_37d7f27c5f97b901) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_37d7f27c5f97b901 string) bool {
	return strings.Contains(__obf_37d7f27c5f97b901, ":") && net.ParseIP(__obf_37d7f27c5f97b901) != nil
}

func AnyToBytes(__obf_194e82db8912d873 any) ([]byte, error) {
	return msgpack.Marshal(__obf_194e82db8912d873)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_263270b555073815 []byte) (__obf_703a701ac3bc0aab T, __obf_65d3bb18bf1bd15f error) {
	__obf_65d3bb18bf1bd15f = msgpack.Unmarshal(__obf_263270b555073815, &__obf_703a701ac3bc0aab)
	return
}

func Loadyaml(__obf_2bc9e2fd7368c566 string, __obf_215d6b6fa4441f92 any) {
	__obf_00e3488b89a95768, __obf_65d3bb18bf1bd15f := os.ReadFile(__obf_2bc9e2fd7368c566)
	if __obf_65d3bb18bf1bd15f != nil {
		log.Fatalln(__obf_65d3bb18bf1bd15f)
	}
	__obf_65d3bb18bf1bd15f = yaml.UnmarshalStrict(__obf_00e3488b89a95768, __obf_215d6b6fa4441f92)
	if __obf_65d3bb18bf1bd15f != nil {
		log.Fatalln(__obf_65d3bb18bf1bd15f)
	}
}

func ToAnyList[T any](__obf_b4a97c22d49a2f64 []T) []any {
	__obf_7f856d40f1b20cb7 := make([]any, len(__obf_b4a97c22d49a2f64))
	for __obf_2482e14692f1453a, __obf_194e82db8912d873 := range __obf_b4a97c22d49a2f64 {
		__obf_7f856d40f1b20cb7[__obf_2482e14692f1453a] = __obf_194e82db8912d873
	}
	return __obf_7f856d40f1b20cb7
}

func TimeParse(__obf_b194262236753edc string) time.Time {
	__obf_b194262236753edc = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_b194262236753edc)
	__obf_d881f8a0ff5d4e10, _ := time.ParseInLocation("2006-01-02 15:04", __obf_b194262236753edc, time.Local)
	return __obf_d881f8a0ff5d4e10
}
