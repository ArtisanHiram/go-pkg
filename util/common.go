package __obf_85fcc7bd65d30170

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
	__obf_9c2b993d609fe396 = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_8b4207081337153d any) ([]byte, error) {
	return __obf_9c2b993d609fe396.Marshal(__obf_8b4207081337153d)
}

func EncodeString(__obf_8b4207081337153d any) string {
	if __obf_a88a18b8d1bc6b68, __obf_80a78f1bfaca43d3 := __obf_9c2b993d609fe396.Marshal(__obf_8b4207081337153d); __obf_80a78f1bfaca43d3 == nil {
		return string(__obf_a88a18b8d1bc6b68)
	}
	return ""
}

func Decode(__obf_fc15ddec64e94329 string, __obf_8b4207081337153d any) error {
	return __obf_9c2b993d609fe396.UnmarshalFromString(__obf_fc15ddec64e94329, __obf_8b4207081337153d)
}

func DecodeByte(__obf_7fcfc8a8085d93d2 []byte, __obf_8b4207081337153d any) error {
	return __obf_9c2b993d609fe396.Unmarshal(__obf_7fcfc8a8085d93d2, __obf_8b4207081337153d)
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

func PrefixInArray(__obf_5857e227a46c2375 string, __obf_366f285a106c8f28 []string) bool {
	for __obf_80f7ae497d9fa3a1 := range __obf_366f285a106c8f28 {
		if strings.HasPrefix(__obf_5857e227a46c2375, __obf_366f285a106c8f28[__obf_80f7ae497d9fa3a1]) {
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

var __obf_6ede3b63a5061f9c = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var __obf_5a0766ec6c9c5678 = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenRandom 获得随机字符串
func GenRandom(__obf_0ccb3bbe367656b6 int) string {
	__obf_7fcfc8a8085d93d2 := make([]byte, __obf_0ccb3bbe367656b6)
	for __obf_80f7ae497d9fa3a1 := range __obf_7fcfc8a8085d93d2 {
		__obf_7fcfc8a8085d93d2[__obf_80f7ae497d9fa3a1] = __obf_6ede3b63a5061f9c[__obf_5a0766ec6c9c5678.Intn(len(__obf_6ede3b63a5061f9c))]
	}
	return string(__obf_7fcfc8a8085d93d2)
}

func RemoveIndex(__obf_d90c40a5354bf79c []string, __obf_9dff7bde685a3eb4 int) []string {
	return append(__obf_d90c40a5354bf79c[:__obf_9dff7bde685a3eb4], __obf_d90c40a5354bf79c[__obf_9dff7bde685a3eb4+1:]...)
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
func ToCamel(__obf_fc15ddec64e94329 string) (__obf_521605711569ff86 string) {
	__obf_e2c8f7c0c56f452b := []rune(__obf_fc15ddec64e94329)
	__obf_521605711569ff86 = __obf_fc15ddec64e94329[0:1]
	if __obf_e2c8f7c0c56f452b[0] >= 97 && __obf_e2c8f7c0c56f452b[0] <= 122 {
		__obf_521605711569ff86 = string(__obf_e2c8f7c0c56f452b[0] - 32)
	}

	__obf_9c478247f32e3f14 := len(__obf_e2c8f7c0c56f452b)
	for __obf_80f7ae497d9fa3a1 := 1; __obf_80f7ae497d9fa3a1 < __obf_9c478247f32e3f14; __obf_80f7ae497d9fa3a1++ {
		if __obf_e2c8f7c0c56f452b[__obf_80f7ae497d9fa3a1] == 95 && __obf_e2c8f7c0c56f452b[__obf_80f7ae497d9fa3a1+1] >= 97 && __obf_e2c8f7c0c56f452b[__obf_80f7ae497d9fa3a1+1] <= 122 { //过滤下划线
			__obf_e2c8f7c0c56f452b[__obf_80f7ae497d9fa3a1+1] -= 32
		} else {
			__obf_521605711569ff86 += string(__obf_e2c8f7c0c56f452b[__obf_80f7ae497d9fa3a1])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_fc15ddec64e94329 string) (__obf_521605711569ff86 string) {
	__obf_e2c8f7c0c56f452b := []rune(__obf_fc15ddec64e94329)
	__obf_521605711569ff86 = __obf_fc15ddec64e94329[0:1]
	if __obf_e2c8f7c0c56f452b[0] >= 65 && __obf_e2c8f7c0c56f452b[0] <= 90 {
		__obf_521605711569ff86 = string(__obf_e2c8f7c0c56f452b[0] + 32)
	}

	__obf_0ccb3bbe367656b6 := len(__obf_e2c8f7c0c56f452b)
	for __obf_80f7ae497d9fa3a1 := 1; __obf_80f7ae497d9fa3a1 < __obf_0ccb3bbe367656b6; __obf_80f7ae497d9fa3a1++ {
		if __obf_e2c8f7c0c56f452b[__obf_80f7ae497d9fa3a1] >= 65 && __obf_e2c8f7c0c56f452b[__obf_80f7ae497d9fa3a1] <= 90 { //大写变小写
			__obf_e2c8f7c0c56f452b[__obf_80f7ae497d9fa3a1] += 32
			__obf_521605711569ff86 += "_"
		}
		__obf_521605711569ff86 += string(__obf_e2c8f7c0c56f452b[__obf_80f7ae497d9fa3a1])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_fc15ddec64e94329 string, __obf_33967b5b791d3d39 int, __obf_0ccb3bbe367656b6 int) string {
	// 将字符串转换为rune切片，以正确处理多字节字符
	__obf_2f6a7e4e091be9e6 := []rune(__obf_fc15ddec64e94329)
	__obf_11bffd4ed6d8c009 := len(__obf_2f6a7e4e091be9e6)

	// 处理n为负数或0的无效情况
	if __obf_0ccb3bbe367656b6 <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_33967b5b791d3d39 < 0 {
		__obf_33967b5b791d3d39 = __obf_11bffd4ed6d8c009 + __obf_33967b5b791d3d39
	}

	// 边界检查：确保start在有效范围内
	if __obf_33967b5b791d3d39 < 0 || __obf_33967b5b791d3d39 >= __obf_11bffd4ed6d8c009 {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_2f6a7e4e091be9e6[__obf_33967b5b791d3d39:min(__obf_33967b5b791d3d39+__obf_0ccb3bbe367656b6, __obf_11bffd4ed6d8c009)])
}

func ASCII(__obf_4edebc331d8e8496 rune) rune {
	switch {
	case 97 <= __obf_4edebc331d8e8496 && __obf_4edebc331d8e8496 <= 122:
		return __obf_4edebc331d8e8496 - 32
	case 65 <= __obf_4edebc331d8e8496 && __obf_4edebc331d8e8496 <= 90:
		return __obf_4edebc331d8e8496 + 32
	default:
		return __obf_4edebc331d8e8496
	}
}

func IndexString(__obf_fc15ddec64e94329 string, __obf_c573de5b483263b2 rune, __obf_49eb49aa6ddbdfda int) string {
	__obf_072c32c6b24bf10b := []rune(__obf_fc15ddec64e94329)
	var __obf_9a6e2ed415d6972a bytes.Buffer
	var __obf_0ccb3bbe367656b6 int
	for __obf_80f7ae497d9fa3a1, __obf_f7032d20c157a404 := 0, len(__obf_072c32c6b24bf10b); __obf_80f7ae497d9fa3a1 < __obf_f7032d20c157a404; __obf_80f7ae497d9fa3a1++ {
		if __obf_072c32c6b24bf10b[__obf_80f7ae497d9fa3a1] == __obf_c573de5b483263b2 {
			__obf_0ccb3bbe367656b6 += 1
		}
		if __obf_0ccb3bbe367656b6 == __obf_49eb49aa6ddbdfda {
			break
		}
		__obf_9a6e2ed415d6972a.WriteRune(__obf_072c32c6b24bf10b[__obf_80f7ae497d9fa3a1])
	}
	return __obf_9a6e2ed415d6972a.String()
}

func LastIndexString(__obf_58ac4ace6141f6fb, __obf_cc0855de1f28bfc5 string) string {
	__obf_d90c40a5354bf79c := strings.Split(__obf_58ac4ace6141f6fb, __obf_cc0855de1f28bfc5)
	if __obf_0ccb3bbe367656b6 := len(__obf_d90c40a5354bf79c); __obf_0ccb3bbe367656b6 > 1 {
		return __obf_d90c40a5354bf79c[__obf_0ccb3bbe367656b6-2]
	}
	return ""
}

func IsEmpty(__obf_c82c631722b73b7c any) bool {
	__obf_37a13444ba0aee54 := reflect.ValueOf(__obf_c82c631722b73b7c)
	if __obf_37a13444ba0aee54.Kind() == reflect.Ptr {
		__obf_37a13444ba0aee54 = __obf_37a13444ba0aee54.Elem()
	}
	return __obf_37a13444ba0aee54.Interface() == reflect.Zero(__obf_37a13444ba0aee54.Type()).Interface()
}

func MillisecondToDateString(__obf_71c6347a9ef45c05 int64) string {
	return time.Unix(__obf_71c6347a9ef45c05, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_71c6347a9ef45c05 int64) string {
	return time.Unix(__obf_71c6347a9ef45c05, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_3057041959274750 string) (__obf_ef2d507a631b5565, __obf_32fbedcf5c5eff40 string) {
	__obf_3057041959274750 = filepath.Base(__obf_3057041959274750)
	__obf_32fbedcf5c5eff40 = filepath.Ext(__obf_3057041959274750)
	__obf_ef2d507a631b5565 = strings.TrimSuffix(__obf_3057041959274750, __obf_32fbedcf5c5eff40)
	return __obf_ef2d507a631b5565, __obf_32fbedcf5c5eff40
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
func GetFirstDateOfMonth(__obf_69160e8fc5e99f94 string) (int64, error) {
	if __obf_3921d02e782dfc70, __obf_80a78f1bfaca43d3 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_69160e8fc5e99f94, Loc); __obf_80a78f1bfaca43d3 != nil {
		return 0, nil
	} else {
		__obf_3921d02e782dfc70 = __obf_3921d02e782dfc70.AddDate(0, 0, -__obf_3921d02e782dfc70.Day()+1)
		return GetZeroTime(__obf_3921d02e782dfc70).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_69160e8fc5e99f94 string) (int64, int64, error) {
	if __obf_3921d02e782dfc70, __obf_80a78f1bfaca43d3 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_69160e8fc5e99f94, Loc); __obf_80a78f1bfaca43d3 != nil {
		return 0, 0, nil
	} else {
		__obf_3921d02e782dfc70 = __obf_3921d02e782dfc70.AddDate(0, 0, -__obf_3921d02e782dfc70.Day()+1)
		return GetZeroTime(__obf_3921d02e782dfc70).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_3921d02e782dfc70).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_34214d422ca61c59 time.Time) time.Time {
	return time.Date(__obf_34214d422ca61c59.Year(), __obf_34214d422ca61c59.Month(), __obf_34214d422ca61c59.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_71c6347a9ef45c05 int64) int64 {
	__obf_f6fc3080233182df := time.Unix(__obf_71c6347a9ef45c05, 0)
	__obf_00f0d86e90291bae, __obf_a3c9e3ca5d5e25fb, __obf_34214d422ca61c59 := __obf_f6fc3080233182df.Date()
	return time.Date(__obf_00f0d86e90291bae, __obf_a3c9e3ca5d5e25fb, __obf_34214d422ca61c59, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_71c6347a9ef45c05 int64) int64 {
	__obf_f6fc3080233182df := time.Unix(__obf_71c6347a9ef45c05, 0)
	return __obf_f6fc3080233182df.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_00f0d86e90291bae, __obf_a3c9e3ca5d5e25fb, __obf_34214d422ca61c59 := time.Now().Date()
	return time.Date(__obf_00f0d86e90291bae, __obf_a3c9e3ca5d5e25fb, __obf_34214d422ca61c59, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_69160e8fc5e99f94 string) (int64, int64) {
	__obf_d20ee76104e30889, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_69160e8fc5e99f94+" 00:00:00", Loc)
	min := __obf_d20ee76104e30889.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_e5afc0f27abc6071 int) (int64, int64) {
	__obf_69160e8fc5e99f94 := time.Now()
	__obf_f6fc3080233182df := __obf_69160e8fc5e99f94.AddDate(0, 0, __obf_e5afc0f27abc6071)
	__obf_00f0d86e90291bae, __obf_a3c9e3ca5d5e25fb, __obf_34214d422ca61c59 := __obf_f6fc3080233182df.Date()
	min := time.Date(__obf_00f0d86e90291bae, __obf_a3c9e3ca5d5e25fb, __obf_34214d422ca61c59, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_69160e8fc5e99f94.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_6214b0247bea1810, __obf_26b3f552622ea0fd string) bool {
	__obf_977d8bf3eaa51b6d := func(__obf_d90c40a5354bf79c string) string {
		return fmt.Sprintf(",%s,", __obf_d90c40a5354bf79c)
	}
	return strings.Contains(__obf_977d8bf3eaa51b6d(__obf_6214b0247bea1810), __obf_977d8bf3eaa51b6d(__obf_26b3f552622ea0fd))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_218a5e46a460896c string, __obf_bc9b81c2db13e119 ...string) bool {
	if any {
		for _, __obf_75c519d5845e5d1e := range __obf_bc9b81c2db13e119 {
			if Contain(__obf_218a5e46a460896c, __obf_75c519d5845e5d1e) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_75c519d5845e5d1e := range __obf_bc9b81c2db13e119 {
			if !Contain(__obf_218a5e46a460896c, __obf_75c519d5845e5d1e) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_d90c40a5354bf79c []any, __obf_9dff7bde685a3eb4 int) []any {
	return slices.Delete(__obf_d90c40a5354bf79c, __obf_9dff7bde685a3eb4, __obf_9dff7bde685a3eb4+1)
}

func String2Int8(__obf_fc15ddec64e94329 string) int8 {
	__obf_eecd32e01852d137, __obf_80a78f1bfaca43d3 := strconv.ParseInt(__obf_fc15ddec64e94329, 10, 8)
	if __obf_80a78f1bfaca43d3 == nil {
		return int8(__obf_eecd32e01852d137)
	}
	return 0
}

func String2Int32(__obf_fc15ddec64e94329 string) int32 {
	__obf_eecd32e01852d137, __obf_80a78f1bfaca43d3 := strconv.ParseInt(__obf_fc15ddec64e94329, 10, 32)
	if __obf_80a78f1bfaca43d3 == nil {
		return int32(__obf_eecd32e01852d137)
	}
	return 0
}

func String2Int64(__obf_fc15ddec64e94329 string) int8 {
	__obf_eecd32e01852d137, __obf_80a78f1bfaca43d3 := strconv.ParseInt(__obf_fc15ddec64e94329, 10, 64)
	if __obf_80a78f1bfaca43d3 == nil {
		return int8(__obf_eecd32e01852d137)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_6048110e567fd1d2, __obf_b818f74c5fb64b6c int) (__obf_33967b5b791d3d39, __obf_1df4d36989f556b2 time.Time) {
	if __obf_6048110e567fd1d2 > 0 && __obf_b818f74c5fb64b6c > 0 {
		__obf_33967b5b791d3d39 = time.Date(__obf_6048110e567fd1d2, 1, 0, 0, 0, 0, 0, Loc)
		// 第一天是周几
		__obf_c210e4b350f74fea := int(__obf_33967b5b791d3d39.AddDate(0, 0, 1).Weekday())
		// 当年第一周有几天
		__obf_e42b35dfdc5b80dd := 1
		if __obf_c210e4b350f74fea != 0 {
			__obf_e42b35dfdc5b80dd = 7 - __obf_c210e4b350f74fea + 1
		}
		if __obf_b818f74c5fb64b6c == 1 {
			__obf_1df4d36989f556b2 = __obf_33967b5b791d3d39.AddDate(0, 0, __obf_e42b35dfdc5b80dd)
		} else {
			__obf_1df4d36989f556b2 = __obf_33967b5b791d3d39.AddDate(0, 0, __obf_e42b35dfdc5b80dd+(__obf_b818f74c5fb64b6c-1)*7)
			__obf_33967b5b791d3d39 = __obf_1df4d36989f556b2.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_69160e8fc5e99f94 time.Time) (__obf_6048110e567fd1d2, __obf_b818f74c5fb64b6c, __obf_3efe1295735cbfca int) {
	__obf_6048110e567fd1d2 = __obf_69160e8fc5e99f94.Year()
	__obf_3efe1295735cbfca = int(__obf_69160e8fc5e99f94.Weekday())
	__obf_787c1f7d15a43e57 := __obf_69160e8fc5e99f94.YearDay()
	__obf_7272d1f5d3e9a7e3 := __obf_69160e8fc5e99f94.AddDate(0, 0, -__obf_787c1f7d15a43e57+1)
	__obf_c210e4b350f74fea := int(__obf_7272d1f5d3e9a7e3.Weekday())
	// 当年第一周有几天
	__obf_e42b35dfdc5b80dd := 1
	if __obf_c210e4b350f74fea != 0 {
		__obf_e42b35dfdc5b80dd = 7 - __obf_c210e4b350f74fea + 1
	}
	if __obf_787c1f7d15a43e57 <= __obf_e42b35dfdc5b80dd {
		__obf_b818f74c5fb64b6c = 1
	} else {
		__obf_b818f74c5fb64b6c = (__obf_787c1f7d15a43e57-__obf_e42b35dfdc5b80dd)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_513e39b1608cc49d []byte) map[string]string {
	__obf_24bec09acffa1760 := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_1c36a2ccd9680ebc := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_44f26af6faf0bf2a []byte
	if __obf_44f26af6faf0bf2a = __obf_24bec09acffa1760.ReplaceAll(__obf_513e39b1608cc49d, []byte("")); len(__obf_44f26af6faf0bf2a) < 6 {
		return nil
	}

	__obf_41c312ac426f65ee := __obf_1c36a2ccd9680ebc.FindAllSubmatch(__obf_44f26af6faf0bf2a[6:len(__obf_44f26af6faf0bf2a)-7], -1)
	__obf_63f30b6213b186fd := map[string]string{}
	if __obf_4d76358808ce8f81 := __obf_24bec09acffa1760.FindSubmatch(__obf_513e39b1608cc49d)[1]; len(__obf_4d76358808ce8f81) != 0 {
		__obf_5d3aaeecc993323e := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_4d76358808ce8f81, -1)
		for _, __obf_37a13444ba0aee54 := range __obf_5d3aaeecc993323e {
			__obf_63f30b6213b186fd[string(__obf_37a13444ba0aee54[1])] = string(__obf_37a13444ba0aee54[2])
		}
	}

	for _, __obf_37a13444ba0aee54 := range __obf_41c312ac426f65ee {
		__obf_63f30b6213b186fd[string(__obf_37a13444ba0aee54[1])] = string(__obf_37a13444ba0aee54[2])
	}
	return __obf_63f30b6213b186fd
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_e197c5c0649a49f9 time.Time) (int64, int64) {
	__obf_e795e69980d61900 := GetZeroTime(__obf_e197c5c0649a49f9).Unix() / 600
	return __obf_e795e69980d61900, __obf_e795e69980d61900 + 143
}

func Abs(__obf_0ccb3bbe367656b6 int64) int64 {
	__obf_00f0d86e90291bae := __obf_0ccb3bbe367656b6 >> 63
	return (__obf_0ccb3bbe367656b6 ^ __obf_00f0d86e90291bae) - __obf_00f0d86e90291bae
}

func Number2String(__obf_0ccb3bbe367656b6 any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_0ccb3bbe367656b6 := __obf_0ccb3bbe367656b6.(type) {
	case int:
		return strconv.Itoa(__obf_0ccb3bbe367656b6)
	case int32:
		return strconv.FormatInt(int64(__obf_0ccb3bbe367656b6), 10)
	case int64:
		return strconv.FormatInt(__obf_0ccb3bbe367656b6, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_0ccb3bbe367656b6), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_0ccb3bbe367656b6, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_fc15ddec64e94329 any, __obf_c573de5b483263b2 string) string {
	if __obf_fc15ddec64e94329 != nil && __obf_fc15ddec64e94329.(string) != "" {
		// return sep + str + sep
		return __obf_fc15ddec64e94329.(string)
	}
	return __obf_c573de5b483263b2
}

func SortRange(__obf_a3c9e3ca5d5e25fb map[string]any, __obf_b3f9ba974f202dc7 func(int, string)) {
	var __obf_7aa67adad889ca5c []string
	for __obf_3b9db108b96306b1 := range __obf_a3c9e3ca5d5e25fb {
		__obf_7aa67adad889ca5c = append(__obf_7aa67adad889ca5c, __obf_3b9db108b96306b1)
	}
	sort.Strings(__obf_7aa67adad889ca5c)
	for __obf_80f7ae497d9fa3a1, __obf_3b9db108b96306b1 := range __obf_7aa67adad889ca5c {
		__obf_b3f9ba974f202dc7(__obf_80f7ae497d9fa3a1, __obf_3b9db108b96306b1)
	}
}

func HasField(__obf_16ddb1c6b06afb00 reflect.Value, __obf_ef2d507a631b5565 string) bool {

	if __obf_d90c40a5354bf79c := __obf_16ddb1c6b06afb00.FieldByNameFunc(func(__obf_0ccb3bbe367656b6 string) bool {
		return strings.EqualFold(__obf_0ccb3bbe367656b6, __obf_ef2d507a631b5565)
	}); __obf_d90c40a5354bf79c.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_16ddb1c6b06afb00 reflect.Value, __obf_ef2d507a631b5565 string) reflect.Value {
	if __obf_d90c40a5354bf79c := __obf_16ddb1c6b06afb00.FieldByNameFunc(func(__obf_0ccb3bbe367656b6 string) bool {
		return strings.EqualFold(__obf_0ccb3bbe367656b6, __obf_ef2d507a631b5565)
	}); __obf_d90c40a5354bf79c.IsValid() {
		return __obf_d90c40a5354bf79c
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_16ddb1c6b06afb00 reflect.Value, __obf_ef2d507a631b5565 string, __obf_5857e227a46c2375 any) bool {

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
	if __obf_d90c40a5354bf79c := __obf_16ddb1c6b06afb00.FieldByNameFunc(func(__obf_0ccb3bbe367656b6 string) bool {
		return strings.EqualFold(__obf_0ccb3bbe367656b6, __obf_ef2d507a631b5565)
	}); __obf_d90c40a5354bf79c.IsValid() {
		if __obf_15cd743ecc57fe17 := __obf_d90c40a5354bf79c.Type(); __obf_15cd743ecc57fe17 == reflect.TypeOf(__obf_5857e227a46c2375) {
			__obf_d90c40a5354bf79c.Set(reflect.ValueOf(__obf_5857e227a46c2375))
		} else {
			__obf_d90c40a5354bf79c.Set(reflect.ValueOf(__obf_5857e227a46c2375).Convert(__obf_15cd743ecc57fe17))
		}
		return true
	}
	return false
}

func CopyMap(__obf_a3c9e3ca5d5e25fb map[string]any) map[string]any {
	__obf_5d3aaeecc993323e := make(map[string]any)
	for __obf_3b9db108b96306b1, __obf_37a13444ba0aee54 := range __obf_a3c9e3ca5d5e25fb {
		if __obf_46c4aac445aa01dd, __obf_60ca2467b0c209af := __obf_37a13444ba0aee54.(map[string]any); __obf_60ca2467b0c209af {
			__obf_5d3aaeecc993323e[__obf_3b9db108b96306b1] = CopyMap(__obf_46c4aac445aa01dd)
		} else {
			__obf_5d3aaeecc993323e[__obf_3b9db108b96306b1] = __obf_37a13444ba0aee54
		}
	}

	return __obf_5d3aaeecc993323e
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_592ca3f7d124f0fd, __obf_cb79a6cda0eff58a string, __obf_58dc1db25a87887b bool, __obf_03f3434e00584061 ...any) {
	var __obf_b95eab86bd44e9d7 strings.Builder
	var __obf_45e362c871e3553a string
	var __obf_d772d88c43f01e15 []string
	for _, __obf_0a983f7fc387b193 := range __obf_03f3434e00584061 {
		__obf_d90c40a5354bf79c := reflect.TypeOf(__obf_0a983f7fc387b193)
		__obf_b95eab86bd44e9d7.WriteString(`CREATE TABLE `)
		__obf_b95eab86bd44e9d7.WriteString(__obf_d90c40a5354bf79c.Name())
		__obf_b95eab86bd44e9d7.WriteString(" (\n")
		__obf_38fe022446507a42 := __obf_d90c40a5354bf79c.NumField()
		for __obf_80f7ae497d9fa3a1 := 0; __obf_80f7ae497d9fa3a1 < __obf_38fe022446507a42; __obf_80f7ae497d9fa3a1++ {
			__obf_b95eab86bd44e9d7.WriteString("    ")
			__obf_d772d88c43f01e15 = nil
			if __obf_b98a9aed159a5f3b := string(__obf_d90c40a5354bf79c.Field(__obf_80f7ae497d9fa3a1).Tag.Get(__obf_592ca3f7d124f0fd)); __obf_b98a9aed159a5f3b == "" {
				if __obf_58dc1db25a87887b {
					__obf_45e362c871e3553a = ToCamel(__obf_d90c40a5354bf79c.Field(__obf_80f7ae497d9fa3a1).Name)
				} else {
					__obf_45e362c871e3553a = __obf_d90c40a5354bf79c.Field(__obf_80f7ae497d9fa3a1).Name
				}
			} else {
				__obf_d772d88c43f01e15 = strings.Split(__obf_b98a9aed159a5f3b, __obf_cb79a6cda0eff58a)
				__obf_45e362c871e3553a = __obf_d772d88c43f01e15[0]
			}
			__obf_b95eab86bd44e9d7.WriteString(__obf_45e362c871e3553a)
			__obf_b95eab86bd44e9d7.WriteString(" ")
			switch __obf_d90c40a5354bf79c.Field(__obf_80f7ae497d9fa3a1).Type.Name() {
			case "int8":
				__obf_b95eab86bd44e9d7.WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_b95eab86bd44e9d7.WriteString("INT")
			case "int64":
				__obf_b95eab86bd44e9d7.WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_b95eab86bd44e9d7.WriteString("VARCHAR(50)")
			}

			if len(__obf_d772d88c43f01e15) > 1 {
				__obf_b95eab86bd44e9d7.WriteString(" ")
				__obf_b95eab86bd44e9d7.WriteString(strings.Join(__obf_d772d88c43f01e15[1:], " "))
			}

			if __obf_80f7ae497d9fa3a1+1 != __obf_38fe022446507a42 {
				__obf_b95eab86bd44e9d7.WriteString(",")
			}
			__obf_b95eab86bd44e9d7.WriteString("\n")
		}
		__obf_b95eab86bd44e9d7.WriteString(");\n\n")
	}
	fmt.Println(__obf_b95eab86bd44e9d7.String())
}

func SaveImage(__obf_619043788c5bf5ec image.Image, __obf_182cac9fcafb119f uint, __obf_bc39b3ccd7b68cf9 string) error {
	__obf_5c3414c0ed9225bf, __obf_80a78f1bfaca43d3 := os.Create(__obf_bc39b3ccd7b68cf9)
	if __obf_80a78f1bfaca43d3 != nil {
		return __obf_80a78f1bfaca43d3
	}
	defer __obf_5c3414c0ed9225bf.Close()
	return jpeg.Encode(__obf_5c3414c0ed9225bf, resize.Resize(__obf_182cac9fcafb119f, 0, __obf_619043788c5bf5ec, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_fc15ddec64e94329 string) bool {
	return strings.Contains(__obf_fc15ddec64e94329, ".") && net.ParseIP(__obf_fc15ddec64e94329) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_fc15ddec64e94329 string) bool {
	return strings.Contains(__obf_fc15ddec64e94329, ":") && net.ParseIP(__obf_fc15ddec64e94329) != nil
}

func AnyToBytes(__obf_37a13444ba0aee54 any) ([]byte, error) {
	return msgpack.Marshal(__obf_37a13444ba0aee54)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_5857e227a46c2375 []byte) (__obf_8b4207081337153d T, __obf_80a78f1bfaca43d3 error) {
	__obf_80a78f1bfaca43d3 = msgpack.Unmarshal(__obf_5857e227a46c2375, &__obf_8b4207081337153d)
	return
}

func Loadyaml(__obf_7f52cafc75007067 string, __obf_89359ea67a29add2 any) {
	__obf_26b3f552622ea0fd, __obf_80a78f1bfaca43d3 := os.ReadFile(__obf_7f52cafc75007067)
	if __obf_80a78f1bfaca43d3 != nil {
		log.Fatalln(__obf_80a78f1bfaca43d3)
	}
	__obf_80a78f1bfaca43d3 = yaml.UnmarshalStrict(__obf_26b3f552622ea0fd, __obf_89359ea67a29add2)
	if __obf_80a78f1bfaca43d3 != nil {
		log.Fatalln(__obf_80a78f1bfaca43d3)
	}
}

func ToAnyList[T any](__obf_3498d57201eca6fd []T) []any {
	__obf_c92f0eca123acd75 := make([]any, len(__obf_3498d57201eca6fd))
	for __obf_80f7ae497d9fa3a1, __obf_37a13444ba0aee54 := range __obf_3498d57201eca6fd {
		__obf_c92f0eca123acd75[__obf_80f7ae497d9fa3a1] = __obf_37a13444ba0aee54
	}
	return __obf_c92f0eca123acd75
}

func TimeParse(__obf_b0a528d37794a90e string) time.Time {
	__obf_b0a528d37794a90e = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_b0a528d37794a90e)
	__obf_f6fc3080233182df, _ := time.ParseInLocation("2006-01-02 15:04", __obf_b0a528d37794a90e, time.Local)
	return __obf_f6fc3080233182df
}
