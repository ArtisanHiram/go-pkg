package __obf_f1e346e3aa5cc554

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"net"
	"path/filepath"

	"log"
	"math/rand"
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
	__obf_f3d2a2a7fb490096 = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_e8e5c333aaebfc88 any) ([]byte, error) {
	return __obf_f3d2a2a7fb490096.Marshal(__obf_e8e5c333aaebfc88)
}

func EncodeString(__obf_e8e5c333aaebfc88 any) string {
	if __obf_40a6f6db3e00dcb3, __obf_eec784b359ebf42f := __obf_f3d2a2a7fb490096.Marshal(__obf_e8e5c333aaebfc88); __obf_eec784b359ebf42f == nil {
		return string(__obf_40a6f6db3e00dcb3)
	}
	return ""
}

func Decode(__obf_d8f9dc4f17edd08b string, __obf_e8e5c333aaebfc88 any) error {
	return __obf_f3d2a2a7fb490096.UnmarshalFromString(__obf_d8f9dc4f17edd08b, __obf_e8e5c333aaebfc88)
}

func DecodeByte(__obf_44c24162abb79539 []byte, __obf_e8e5c333aaebfc88 any) error {
	return __obf_f3d2a2a7fb490096.Unmarshal(__obf_44c24162abb79539, __obf_e8e5c333aaebfc88)
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

func PrefixInArray(__obf_8f9b1d9676963689 string, __obf_21346c1d4428aa6a []string) bool {
	for __obf_c6095a9cdc385fed := range __obf_21346c1d4428aa6a {
		if strings.HasPrefix(__obf_8f9b1d9676963689, __obf_21346c1d4428aa6a[__obf_c6095a9cdc385fed]) {
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

var __obf_70f5a02e9d90ad35 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var __obf_1958cb0c2eaa20dc = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenRandom 获得随机字符串
func GenRandom(__obf_719365bda976491c int) string {
	__obf_44c24162abb79539 := make([]byte, __obf_719365bda976491c)
	for __obf_c6095a9cdc385fed := range __obf_44c24162abb79539 {
		__obf_44c24162abb79539[__obf_c6095a9cdc385fed] = __obf_70f5a02e9d90ad35[__obf_1958cb0c2eaa20dc.Intn(len(__obf_70f5a02e9d90ad35))]
	}
	return string(__obf_44c24162abb79539)
}

func RemoveIndex(__obf_0115cb118880283a []string, __obf_eed4872d51469f90 int) []string {
	return append(__obf_0115cb118880283a[:__obf_eed4872d51469f90], __obf_0115cb118880283a[__obf_eed4872d51469f90+1:]...)
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
func ToCamel(__obf_d8f9dc4f17edd08b string) (__obf_969aede8eee9c705 string) {
	__obf_2161c6b3dabeccdb := []rune(__obf_d8f9dc4f17edd08b)
	__obf_969aede8eee9c705 = __obf_d8f9dc4f17edd08b[0:1]
	if __obf_2161c6b3dabeccdb[0] >= 97 && __obf_2161c6b3dabeccdb[0] <= 122 {
		__obf_969aede8eee9c705 = string(__obf_2161c6b3dabeccdb[0] - 32)
	}

	__obf_f7e0b18db9cbbf01 := len(__obf_2161c6b3dabeccdb)
	for __obf_c6095a9cdc385fed := 1; __obf_c6095a9cdc385fed < __obf_f7e0b18db9cbbf01; __obf_c6095a9cdc385fed++ {
		if __obf_2161c6b3dabeccdb[__obf_c6095a9cdc385fed] == 95 && __obf_2161c6b3dabeccdb[__obf_c6095a9cdc385fed+1] >= 97 && __obf_2161c6b3dabeccdb[__obf_c6095a9cdc385fed+1] <= 122 { //过滤下划线
			__obf_2161c6b3dabeccdb[__obf_c6095a9cdc385fed+1] -= 32
		} else {
			__obf_969aede8eee9c705 += string(__obf_2161c6b3dabeccdb[__obf_c6095a9cdc385fed])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_d8f9dc4f17edd08b string) (__obf_969aede8eee9c705 string) {
	__obf_2161c6b3dabeccdb := []rune(__obf_d8f9dc4f17edd08b)
	__obf_969aede8eee9c705 = __obf_d8f9dc4f17edd08b[0:1]
	if __obf_2161c6b3dabeccdb[0] >= 65 && __obf_2161c6b3dabeccdb[0] <= 90 {
		__obf_969aede8eee9c705 = string(__obf_2161c6b3dabeccdb[0] + 32)
	}

	__obf_719365bda976491c := len(__obf_2161c6b3dabeccdb)
	for __obf_c6095a9cdc385fed := 1; __obf_c6095a9cdc385fed < __obf_719365bda976491c; __obf_c6095a9cdc385fed++ {
		if __obf_2161c6b3dabeccdb[__obf_c6095a9cdc385fed] >= 65 && __obf_2161c6b3dabeccdb[__obf_c6095a9cdc385fed] <= 90 { //大写变小写
			__obf_2161c6b3dabeccdb[__obf_c6095a9cdc385fed] += 32
			__obf_969aede8eee9c705 += "_"
		}
		__obf_969aede8eee9c705 += string(__obf_2161c6b3dabeccdb[__obf_c6095a9cdc385fed])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_d8f9dc4f17edd08b string, __obf_a654a3441d92acdd int, __obf_719365bda976491c int) string {
	// 将字符串转换为rune切片，以正确处理多字节字符
	__obf_126def84f40c23ed := []rune(__obf_d8f9dc4f17edd08b)
	__obf_60e733d31111c611 := len(__obf_126def84f40c23ed)

	// 处理n为负数或0的无效情况
	if __obf_719365bda976491c <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_a654a3441d92acdd < 0 {
		__obf_a654a3441d92acdd = __obf_60e733d31111c611 + __obf_a654a3441d92acdd
	}

	// 边界检查：确保start在有效范围内
	if __obf_a654a3441d92acdd < 0 || __obf_a654a3441d92acdd >= __obf_60e733d31111c611 {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_126def84f40c23ed[__obf_a654a3441d92acdd:min(__obf_a654a3441d92acdd+__obf_719365bda976491c, __obf_60e733d31111c611)])
}

func ASCII(__obf_72fc7f85fed6cac3 rune) rune {
	switch {
	case 97 <= __obf_72fc7f85fed6cac3 && __obf_72fc7f85fed6cac3 <= 122:
		return __obf_72fc7f85fed6cac3 - 32
	case 65 <= __obf_72fc7f85fed6cac3 && __obf_72fc7f85fed6cac3 <= 90:
		return __obf_72fc7f85fed6cac3 + 32
	default:
		return __obf_72fc7f85fed6cac3
	}
}

func IndexString(__obf_d8f9dc4f17edd08b string, __obf_41f2064bbc5b2f8c rune, __obf_c9aacbd3702f33b4 int) string {
	__obf_c2d2bd16d37ee8a9 := []rune(__obf_d8f9dc4f17edd08b)
	var __obf_85ccd31995966a71 bytes.Buffer
	var __obf_719365bda976491c int
	for __obf_c6095a9cdc385fed, __obf_1619ce9f476ca3fb := 0, len(__obf_c2d2bd16d37ee8a9); __obf_c6095a9cdc385fed < __obf_1619ce9f476ca3fb; __obf_c6095a9cdc385fed++ {
		if __obf_c2d2bd16d37ee8a9[__obf_c6095a9cdc385fed] == __obf_41f2064bbc5b2f8c {
			__obf_719365bda976491c += 1
		}
		if __obf_719365bda976491c == __obf_c9aacbd3702f33b4 {
			break
		}
		__obf_85ccd31995966a71.WriteRune(__obf_c2d2bd16d37ee8a9[__obf_c6095a9cdc385fed])
	}
	return __obf_85ccd31995966a71.String()
}

func LastIndexString(__obf_ac1f80c958d4d9f8, __obf_8e7035aa779a6038 string) string {
	__obf_0115cb118880283a := strings.Split(__obf_ac1f80c958d4d9f8, __obf_8e7035aa779a6038)
	if __obf_719365bda976491c := len(__obf_0115cb118880283a); __obf_719365bda976491c > 1 {
		return __obf_0115cb118880283a[__obf_719365bda976491c-2]
	}
	return ""
}

func IsEmpty(__obf_d8b2f7111f17bddd any) bool {
	__obf_29f3385acfb70f69 := reflect.ValueOf(__obf_d8b2f7111f17bddd)
	if __obf_29f3385acfb70f69.Kind() == reflect.Ptr {
		__obf_29f3385acfb70f69 = __obf_29f3385acfb70f69.Elem()
	}
	return __obf_29f3385acfb70f69.Interface() == reflect.Zero(__obf_29f3385acfb70f69.Type()).Interface()
}

func MillisecondToDateString(__obf_2fa96c2bf75184ec int64) string {
	return time.Unix(__obf_2fa96c2bf75184ec, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_2fa96c2bf75184ec int64) string {
	return time.Unix(__obf_2fa96c2bf75184ec, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_9f8049147b4183e3 string) (__obf_1d62aee7a77ffb9c, __obf_3fbb1b761e3f8517 string) {
	__obf_9f8049147b4183e3 = filepath.Base(__obf_9f8049147b4183e3)
	__obf_3fbb1b761e3f8517 = filepath.Ext(__obf_9f8049147b4183e3)
	__obf_1d62aee7a77ffb9c = strings.TrimSuffix(__obf_9f8049147b4183e3, __obf_3fbb1b761e3f8517)
	return __obf_1d62aee7a77ffb9c, __obf_3fbb1b761e3f8517
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
func GetFirstDateOfMonth(__obf_d34c56e6ecc42b40 string) (int64, error) {
	if __obf_ecb27042443d0be0, __obf_eec784b359ebf42f := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_d34c56e6ecc42b40, Loc); __obf_eec784b359ebf42f != nil {
		return 0, nil
	} else {
		__obf_ecb27042443d0be0 = __obf_ecb27042443d0be0.AddDate(0, 0, -__obf_ecb27042443d0be0.Day()+1)
		return GetZeroTime(__obf_ecb27042443d0be0).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_d34c56e6ecc42b40 string) (int64, int64, error) {
	if __obf_ecb27042443d0be0, __obf_eec784b359ebf42f := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_d34c56e6ecc42b40, Loc); __obf_eec784b359ebf42f != nil {
		return 0, 0, nil
	} else {
		__obf_ecb27042443d0be0 = __obf_ecb27042443d0be0.AddDate(0, 0, -__obf_ecb27042443d0be0.Day()+1)
		return GetZeroTime(__obf_ecb27042443d0be0).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_ecb27042443d0be0).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_8beab787517c6111 time.Time) time.Time {
	return time.Date(__obf_8beab787517c6111.Year(), __obf_8beab787517c6111.Month(), __obf_8beab787517c6111.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_2fa96c2bf75184ec int64) int64 {
	__obf_96b180d9b3b77b6c := time.Unix(__obf_2fa96c2bf75184ec, 0)
	__obf_aa280eefe020afb3, __obf_941ea73b441588af, __obf_8beab787517c6111 := __obf_96b180d9b3b77b6c.Date()
	return time.Date(__obf_aa280eefe020afb3, __obf_941ea73b441588af, __obf_8beab787517c6111, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_2fa96c2bf75184ec int64) int64 {
	__obf_96b180d9b3b77b6c := time.Unix(__obf_2fa96c2bf75184ec, 0)
	return __obf_96b180d9b3b77b6c.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_aa280eefe020afb3, __obf_941ea73b441588af, __obf_8beab787517c6111 := time.Now().Date()
	return time.Date(__obf_aa280eefe020afb3, __obf_941ea73b441588af, __obf_8beab787517c6111, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_d34c56e6ecc42b40 string) (int64, int64) {
	__obf_6f98439ac0eecf99, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_d34c56e6ecc42b40+" 00:00:00", Loc)
	min := __obf_6f98439ac0eecf99.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_0b5d40def09b1ee6 int) (int64, int64) {
	__obf_d34c56e6ecc42b40 := time.Now()
	__obf_96b180d9b3b77b6c := __obf_d34c56e6ecc42b40.AddDate(0, 0, __obf_0b5d40def09b1ee6)
	__obf_aa280eefe020afb3, __obf_941ea73b441588af, __obf_8beab787517c6111 := __obf_96b180d9b3b77b6c.Date()
	min := time.Date(__obf_aa280eefe020afb3, __obf_941ea73b441588af, __obf_8beab787517c6111, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_d34c56e6ecc42b40.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_3c54a8c9814b0d37, __obf_638b4e1cada6af75 string) bool {
	__obf_1863bf8c7cda7588 := func(__obf_0115cb118880283a string) string {
		return fmt.Sprintf(",%s,", __obf_0115cb118880283a)
	}
	return strings.Contains(__obf_1863bf8c7cda7588(__obf_3c54a8c9814b0d37), __obf_1863bf8c7cda7588(__obf_638b4e1cada6af75))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_18ed008c1e57ff24 string, __obf_d391f92f6135e2c7 ...string) bool {
	if any {
		for _, __obf_d4bd147170d96001 := range __obf_d391f92f6135e2c7 {
			if Contain(__obf_18ed008c1e57ff24, __obf_d4bd147170d96001) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_d4bd147170d96001 := range __obf_d391f92f6135e2c7 {
			if !Contain(__obf_18ed008c1e57ff24, __obf_d4bd147170d96001) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_0115cb118880283a []any, __obf_eed4872d51469f90 int) []any {
	return slices.Delete(__obf_0115cb118880283a, __obf_eed4872d51469f90, __obf_eed4872d51469f90+1)
}

func String2Int8(__obf_d8f9dc4f17edd08b string) int8 {
	__obf_6c3aa62f4935182a, __obf_eec784b359ebf42f := strconv.ParseInt(__obf_d8f9dc4f17edd08b, 10, 8)
	if __obf_eec784b359ebf42f == nil {
		return int8(__obf_6c3aa62f4935182a)
	}
	return 0
}

func String2Int32(__obf_d8f9dc4f17edd08b string) int32 {
	__obf_6c3aa62f4935182a, __obf_eec784b359ebf42f := strconv.ParseInt(__obf_d8f9dc4f17edd08b, 10, 32)
	if __obf_eec784b359ebf42f == nil {
		return int32(__obf_6c3aa62f4935182a)
	}
	return 0
}

func String2Int64(__obf_d8f9dc4f17edd08b string) int8 {
	__obf_6c3aa62f4935182a, __obf_eec784b359ebf42f := strconv.ParseInt(__obf_d8f9dc4f17edd08b, 10, 64)
	if __obf_eec784b359ebf42f == nil {
		return int8(__obf_6c3aa62f4935182a)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_17c36c29cba6b866, __obf_552577b95f9f9701 int) (__obf_a654a3441d92acdd, __obf_5f149c5c3d87e633 time.Time) {
	if __obf_17c36c29cba6b866 > 0 && __obf_552577b95f9f9701 > 0 {
		__obf_a654a3441d92acdd = time.Date(__obf_17c36c29cba6b866, 1, 0, 0, 0, 0, 0, Loc)
		// 第一天是周几
		__obf_fdafe8cf69b7b7e4 := int(__obf_a654a3441d92acdd.AddDate(0, 0, 1).Weekday())
		// 当年第一周有几天
		__obf_e399bf9c2aeee0cb := 1
		if __obf_fdafe8cf69b7b7e4 != 0 {
			__obf_e399bf9c2aeee0cb = 7 - __obf_fdafe8cf69b7b7e4 + 1
		}
		if __obf_552577b95f9f9701 == 1 {
			__obf_5f149c5c3d87e633 = __obf_a654a3441d92acdd.AddDate(0, 0, __obf_e399bf9c2aeee0cb)
		} else {
			__obf_5f149c5c3d87e633 = __obf_a654a3441d92acdd.AddDate(0, 0, __obf_e399bf9c2aeee0cb+(__obf_552577b95f9f9701-1)*7)
			__obf_a654a3441d92acdd = __obf_5f149c5c3d87e633.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_d34c56e6ecc42b40 time.Time) (__obf_17c36c29cba6b866, __obf_552577b95f9f9701, __obf_4b84d52d757f30fb int) {
	__obf_17c36c29cba6b866 = __obf_d34c56e6ecc42b40.Year()
	__obf_4b84d52d757f30fb = int(__obf_d34c56e6ecc42b40.Weekday())
	__obf_8dbfb62fe1037f8f := __obf_d34c56e6ecc42b40.YearDay()
	__obf_88ae18592c30b28a := __obf_d34c56e6ecc42b40.AddDate(0, 0, -__obf_8dbfb62fe1037f8f+1)
	__obf_fdafe8cf69b7b7e4 := int(__obf_88ae18592c30b28a.Weekday())
	// 当年第一周有几天
	__obf_e399bf9c2aeee0cb := 1
	if __obf_fdafe8cf69b7b7e4 != 0 {
		__obf_e399bf9c2aeee0cb = 7 - __obf_fdafe8cf69b7b7e4 + 1
	}
	if __obf_8dbfb62fe1037f8f <= __obf_e399bf9c2aeee0cb {
		__obf_552577b95f9f9701 = 1
	} else {
		__obf_552577b95f9f9701 = (__obf_8dbfb62fe1037f8f-__obf_e399bf9c2aeee0cb)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_ef338491beae81ce []byte) map[string]string {
	__obf_a1fac298fbd8b132 := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_926886db3a68f8c3 := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_ef24bb248dd834d9 []byte
	if __obf_ef24bb248dd834d9 = __obf_a1fac298fbd8b132.ReplaceAll(__obf_ef338491beae81ce, []byte("")); len(__obf_ef24bb248dd834d9) < 6 {
		return nil
	}

	__obf_0399d349c9773b80 := __obf_926886db3a68f8c3.FindAllSubmatch(__obf_ef24bb248dd834d9[6:len(__obf_ef24bb248dd834d9)-7], -1)
	__obf_bcbbc5bec8036d2c := map[string]string{}
	if __obf_51cb0633104fd84e := __obf_a1fac298fbd8b132.FindSubmatch(__obf_ef338491beae81ce)[1]; len(__obf_51cb0633104fd84e) != 0 {
		__obf_d8903c7e20cc4ec5 := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_51cb0633104fd84e, -1)
		for _, __obf_29f3385acfb70f69 := range __obf_d8903c7e20cc4ec5 {
			__obf_bcbbc5bec8036d2c[string(__obf_29f3385acfb70f69[1])] = string(__obf_29f3385acfb70f69[2])
		}
	}

	for _, __obf_29f3385acfb70f69 := range __obf_0399d349c9773b80 {
		__obf_bcbbc5bec8036d2c[string(__obf_29f3385acfb70f69[1])] = string(__obf_29f3385acfb70f69[2])
	}
	return __obf_bcbbc5bec8036d2c
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_b43668ac5189b4ef time.Time) (int64, int64) {
	__obf_fbe52828a0cdfed1 := GetZeroTime(__obf_b43668ac5189b4ef).Unix() / 600
	return __obf_fbe52828a0cdfed1, __obf_fbe52828a0cdfed1 + 143
}

func Abs(__obf_719365bda976491c int64) int64 {
	__obf_aa280eefe020afb3 := __obf_719365bda976491c >> 63
	return (__obf_719365bda976491c ^ __obf_aa280eefe020afb3) - __obf_aa280eefe020afb3
}

func Number2String(__obf_719365bda976491c any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_719365bda976491c := __obf_719365bda976491c.(type) {
	case int:
		return strconv.Itoa(__obf_719365bda976491c)
	case int32:
		return strconv.FormatInt(int64(__obf_719365bda976491c), 10)
	case int64:
		return strconv.FormatInt(__obf_719365bda976491c, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_719365bda976491c), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_719365bda976491c, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_d8f9dc4f17edd08b any, __obf_41f2064bbc5b2f8c string) string {
	if __obf_d8f9dc4f17edd08b != nil && __obf_d8f9dc4f17edd08b.(string) != "" {
		// return sep + str + sep
		return __obf_d8f9dc4f17edd08b.(string)
	}
	return __obf_41f2064bbc5b2f8c
}

func SortRange(__obf_941ea73b441588af map[string]any, __obf_1027c1dd0a81527d func(int, string)) {
	var __obf_d22ef00394e112b3 []string
	for __obf_6126311d2a895ba1 := range __obf_941ea73b441588af {
		__obf_d22ef00394e112b3 = append(__obf_d22ef00394e112b3, __obf_6126311d2a895ba1)
	}
	sort.Strings(__obf_d22ef00394e112b3)
	for __obf_c6095a9cdc385fed, __obf_6126311d2a895ba1 := range __obf_d22ef00394e112b3 {
		__obf_1027c1dd0a81527d(__obf_c6095a9cdc385fed, __obf_6126311d2a895ba1)
	}
}

func HasField(__obf_388199f0c9224db2 reflect.Value, __obf_1d62aee7a77ffb9c string) bool {

	if __obf_0115cb118880283a := __obf_388199f0c9224db2.FieldByNameFunc(func(__obf_719365bda976491c string) bool {
		return strings.EqualFold(__obf_719365bda976491c, __obf_1d62aee7a77ffb9c)
	}); __obf_0115cb118880283a.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_388199f0c9224db2 reflect.Value, __obf_1d62aee7a77ffb9c string) reflect.Value {
	if __obf_0115cb118880283a := __obf_388199f0c9224db2.FieldByNameFunc(func(__obf_719365bda976491c string) bool {
		return strings.EqualFold(__obf_719365bda976491c, __obf_1d62aee7a77ffb9c)
	}); __obf_0115cb118880283a.IsValid() {
		return __obf_0115cb118880283a
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_388199f0c9224db2 reflect.Value, __obf_1d62aee7a77ffb9c string, __obf_8f9b1d9676963689 any) bool {

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
	if __obf_0115cb118880283a := __obf_388199f0c9224db2.FieldByNameFunc(func(__obf_719365bda976491c string) bool {
		return strings.EqualFold(__obf_719365bda976491c, __obf_1d62aee7a77ffb9c)
	}); __obf_0115cb118880283a.IsValid() {
		if __obf_a85583fa8a8104e3 := __obf_0115cb118880283a.Type(); __obf_a85583fa8a8104e3 == reflect.TypeOf(__obf_8f9b1d9676963689) {
			__obf_0115cb118880283a.Set(reflect.ValueOf(__obf_8f9b1d9676963689))
		} else {
			__obf_0115cb118880283a.Set(reflect.ValueOf(__obf_8f9b1d9676963689).Convert(__obf_a85583fa8a8104e3))
		}
		return true
	}
	return false
}

func CopyMap(__obf_941ea73b441588af map[string]any) map[string]any {
	__obf_d8903c7e20cc4ec5 := make(map[string]any)
	for __obf_6126311d2a895ba1, __obf_29f3385acfb70f69 := range __obf_941ea73b441588af {
		if __obf_187271a5698e04ec, __obf_f6f920d3af14ffc4 := __obf_29f3385acfb70f69.(map[string]any); __obf_f6f920d3af14ffc4 {
			__obf_d8903c7e20cc4ec5[__obf_6126311d2a895ba1] = CopyMap(__obf_187271a5698e04ec)
		} else {
			__obf_d8903c7e20cc4ec5[__obf_6126311d2a895ba1] = __obf_29f3385acfb70f69
		}
	}

	return __obf_d8903c7e20cc4ec5
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_9a20564fc99ec5cb, __obf_47c11abcf013771d string, __obf_c951347d8141d7c6 bool, __obf_c5118b339bc30c90 ...any) {
	var __obf_b8db8c99c5bd923b strings.Builder
	var __obf_8c260306a04b43c6 string
	var __obf_b4e4b2412efcbc7a []string
	for _, __obf_876da6167aa49686 := range __obf_c5118b339bc30c90 {
		__obf_0115cb118880283a := reflect.TypeOf(__obf_876da6167aa49686)
		__obf_b8db8c99c5bd923b.WriteString(`CREATE TABLE `)
		__obf_b8db8c99c5bd923b.WriteString(__obf_0115cb118880283a.Name())
		__obf_b8db8c99c5bd923b.WriteString(" (\n")
		__obf_5e8046ffa6b5fdc7 := __obf_0115cb118880283a.NumField()
		for __obf_c6095a9cdc385fed := 0; __obf_c6095a9cdc385fed < __obf_5e8046ffa6b5fdc7; __obf_c6095a9cdc385fed++ {
			__obf_b8db8c99c5bd923b.WriteString("    ")
			__obf_b4e4b2412efcbc7a = nil
			if __obf_8960e1b1e842e1df := string(__obf_0115cb118880283a.Field(__obf_c6095a9cdc385fed).Tag.Get(__obf_9a20564fc99ec5cb)); __obf_8960e1b1e842e1df == "" {
				if __obf_c951347d8141d7c6 {
					__obf_8c260306a04b43c6 = ToCamel(__obf_0115cb118880283a.Field(__obf_c6095a9cdc385fed).Name)
				} else {
					__obf_8c260306a04b43c6 = __obf_0115cb118880283a.Field(__obf_c6095a9cdc385fed).Name
				}
			} else {
				__obf_b4e4b2412efcbc7a = strings.Split(__obf_8960e1b1e842e1df, __obf_47c11abcf013771d)
				__obf_8c260306a04b43c6 = __obf_b4e4b2412efcbc7a[0]
			}
			__obf_b8db8c99c5bd923b.WriteString(__obf_8c260306a04b43c6)
			__obf_b8db8c99c5bd923b.WriteString(" ")
			switch __obf_0115cb118880283a.Field(__obf_c6095a9cdc385fed).Type.Name() {
			case "int8":
				__obf_b8db8c99c5bd923b.WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_b8db8c99c5bd923b.WriteString("INT")
			case "int64":
				__obf_b8db8c99c5bd923b.WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_b8db8c99c5bd923b.WriteString("VARCHAR(50)")
			}

			if len(__obf_b4e4b2412efcbc7a) > 1 {
				__obf_b8db8c99c5bd923b.WriteString(" ")
				__obf_b8db8c99c5bd923b.WriteString(strings.Join(__obf_b4e4b2412efcbc7a[1:], " "))
			}

			if __obf_c6095a9cdc385fed+1 != __obf_5e8046ffa6b5fdc7 {
				__obf_b8db8c99c5bd923b.WriteString(",")
			}
			__obf_b8db8c99c5bd923b.WriteString("\n")
		}
		__obf_b8db8c99c5bd923b.WriteString(");\n\n")
	}
	fmt.Println(__obf_b8db8c99c5bd923b.String())
}

func SaveImage(__obf_779fc0f42e26bffb image.Image, __obf_23944a10381541b6 uint, __obf_87217cc53d764936 string) error {
	__obf_a10af74ab7e40165, __obf_eec784b359ebf42f := os.Create(__obf_87217cc53d764936)
	if __obf_eec784b359ebf42f != nil {
		return __obf_eec784b359ebf42f
	}
	defer __obf_a10af74ab7e40165.Close()
	return jpeg.Encode(__obf_a10af74ab7e40165, resize.Resize(__obf_23944a10381541b6, 0, __obf_779fc0f42e26bffb, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_d8f9dc4f17edd08b string) bool {
	return strings.Contains(__obf_d8f9dc4f17edd08b, ".") && net.ParseIP(__obf_d8f9dc4f17edd08b) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_d8f9dc4f17edd08b string) bool {
	return strings.Contains(__obf_d8f9dc4f17edd08b, ":") && net.ParseIP(__obf_d8f9dc4f17edd08b) != nil
}

func AnyToBytes(__obf_29f3385acfb70f69 any) ([]byte, error) {
	return msgpack.Marshal(__obf_29f3385acfb70f69)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_8f9b1d9676963689 []byte) (__obf_e8e5c333aaebfc88 T, __obf_eec784b359ebf42f error) {
	__obf_eec784b359ebf42f = msgpack.Unmarshal(__obf_8f9b1d9676963689, &__obf_e8e5c333aaebfc88)
	return
}

func Loadyaml(__obf_cc63b9e89e6f82ad string, __obf_7dd83a2b1dec4622 any) {
	__obf_638b4e1cada6af75, __obf_eec784b359ebf42f := os.ReadFile(__obf_cc63b9e89e6f82ad)
	if __obf_eec784b359ebf42f != nil {
		log.Fatalln(__obf_eec784b359ebf42f)
	}
	__obf_eec784b359ebf42f = yaml.UnmarshalStrict(__obf_638b4e1cada6af75, __obf_7dd83a2b1dec4622)
	if __obf_eec784b359ebf42f != nil {
		log.Fatalln(__obf_eec784b359ebf42f)
	}
}

func ToAnyList[T any](__obf_c1e323470bd8f42f []T) []any {
	__obf_2c5bb84cb28204b7 := make([]any, len(__obf_c1e323470bd8f42f))
	for __obf_c6095a9cdc385fed, __obf_29f3385acfb70f69 := range __obf_c1e323470bd8f42f {
		__obf_2c5bb84cb28204b7[__obf_c6095a9cdc385fed] = __obf_29f3385acfb70f69
	}
	return __obf_2c5bb84cb28204b7
}

func TimeParse(__obf_4207745024457aaf string) time.Time {
	__obf_4207745024457aaf = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_4207745024457aaf)
	__obf_96b180d9b3b77b6c, _ := time.ParseInLocation("2006-01-02 15:04", __obf_4207745024457aaf, time.Local)
	return __obf_96b180d9b3b77b6c
}
