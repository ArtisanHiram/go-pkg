package __obf_e13b701bec415b48

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
	__obf_8a0276aba040f633 = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_1e49a709823fcdda any) ([]byte, error) {
	return __obf_8a0276aba040f633.Marshal(__obf_1e49a709823fcdda)
}

func EncodeString(__obf_1e49a709823fcdda any) string {
	if __obf_de4fc3eb6573a952, __obf_b93c04d14e5c503f := __obf_8a0276aba040f633.Marshal(__obf_1e49a709823fcdda); __obf_b93c04d14e5c503f == nil {
		return string(__obf_de4fc3eb6573a952)
	}
	return ""
}

func Decode(__obf_db7cf183ccadbefa string, __obf_1e49a709823fcdda any) error {
	return __obf_8a0276aba040f633.UnmarshalFromString(__obf_db7cf183ccadbefa, __obf_1e49a709823fcdda)
}

func DecodeByte(__obf_2f1ab84c85030098 []byte, __obf_1e49a709823fcdda any) error {
	return __obf_8a0276aba040f633.Unmarshal(__obf_2f1ab84c85030098, __obf_1e49a709823fcdda)
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

func PrefixInArray(__obf_aae3b6f1d4dc6396 string, __obf_f8e4738ab2360c9d []string) bool {
	for __obf_7cd64e5d17b513e6 := range __obf_f8e4738ab2360c9d {
		if strings.HasPrefix(__obf_aae3b6f1d4dc6396, __obf_f8e4738ab2360c9d[__obf_7cd64e5d17b513e6]) {
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

var __obf_b858c136bc0ad6f7 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenRandom 获得随机字符串
func GenRandom(__obf_550d8612d04f5667 int) string {
	__obf_2f1ab84c85030098 := make([]byte, __obf_550d8612d04f5667)
	for __obf_7cd64e5d17b513e6 := range __obf_2f1ab84c85030098 {
		__obf_2f1ab84c85030098[__obf_7cd64e5d17b513e6] = __obf_b858c136bc0ad6f7[rand.IntN(len(__obf_b858c136bc0ad6f7))]
	}
	return string(__obf_2f1ab84c85030098)
}

func RemoveIndex(__obf_a9c9200b03759058 []string, __obf_db0a1d5b34e76561 int) []string {
	return append(__obf_a9c9200b03759058[:__obf_db0a1d5b34e76561], __obf_a9c9200b03759058[__obf_db0a1d5b34e76561+1:]...)
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
func ToCamel(__obf_db7cf183ccadbefa string) (__obf_d230249dd0c37358 string) {
	__obf_713ffcf459444e84 := []rune(__obf_db7cf183ccadbefa)
	__obf_d230249dd0c37358 = __obf_db7cf183ccadbefa[0:1]
	if __obf_713ffcf459444e84[0] >= 97 && __obf_713ffcf459444e84[0] <= 122 {
		__obf_d230249dd0c37358 = string(__obf_713ffcf459444e84[0] - 32)
	}
	__obf_94249fcac3df5177 := len(__obf_713ffcf459444e84)
	for __obf_7cd64e5d17b513e6 := 1; __obf_7cd64e5d17b513e6 < __obf_94249fcac3df5177; __obf_7cd64e5d17b513e6++ {
		if __obf_713ffcf459444e84[__obf_7cd64e5d17b513e6] == 95 && __obf_713ffcf459444e84[__obf_7cd64e5d17b513e6+1] >= 97 && __obf_713ffcf459444e84[__obf_7cd64e5d17b513e6+1] <= 122 {
			__obf_713ffcf459444e84[ //过滤下划线
			__obf_7cd64e5d17b513e6+1] -= 32
		} else {
			__obf_d230249dd0c37358 += string(__obf_713ffcf459444e84[__obf_7cd64e5d17b513e6])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_db7cf183ccadbefa string) (__obf_d230249dd0c37358 string) {
	__obf_713ffcf459444e84 := []rune(__obf_db7cf183ccadbefa)
	__obf_d230249dd0c37358 = __obf_db7cf183ccadbefa[0:1]
	if __obf_713ffcf459444e84[0] >= 65 && __obf_713ffcf459444e84[0] <= 90 {
		__obf_d230249dd0c37358 = string(__obf_713ffcf459444e84[0] + 32)
	}
	__obf_550d8612d04f5667 := len(__obf_713ffcf459444e84)
	for __obf_7cd64e5d17b513e6 := 1; __obf_7cd64e5d17b513e6 < __obf_550d8612d04f5667; __obf_7cd64e5d17b513e6++ {
		if __obf_713ffcf459444e84[__obf_7cd64e5d17b513e6] >= 65 && __obf_713ffcf459444e84[__obf_7cd64e5d17b513e6] <= 90 {
			__obf_713ffcf459444e84[ //大写变小写
			__obf_7cd64e5d17b513e6] += 32
			__obf_d230249dd0c37358 += "_"
		}
		__obf_d230249dd0c37358 += string(__obf_713ffcf459444e84[__obf_7cd64e5d17b513e6])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_db7cf183ccadbefa string, __obf_e14e1a6f3f62b9ca int, __obf_550d8612d04f5667 int) string {
	__obf_9c045a5bdeba092c := // 将字符串转换为rune切片，以正确处理多字节字符
		[]rune(__obf_db7cf183ccadbefa)
	__obf_fdba24dfa958082c := len(__obf_9c045a5bdeba092c)

	// 处理n为负数或0的无效情况
	if __obf_550d8612d04f5667 <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_e14e1a6f3f62b9ca < 0 {
		__obf_e14e1a6f3f62b9ca = __obf_fdba24dfa958082c + __obf_e14e1a6f3f62b9ca
	}

	// 边界检查：确保start在有效范围内
	if __obf_e14e1a6f3f62b9ca < 0 || __obf_e14e1a6f3f62b9ca >= __obf_fdba24dfa958082c {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_9c045a5bdeba092c[__obf_e14e1a6f3f62b9ca:min(__obf_e14e1a6f3f62b9ca+__obf_550d8612d04f5667, __obf_fdba24dfa958082c)])
}

func ASCII(__obf_af65a69017eda959 rune) rune {
	switch {
	case 97 <= __obf_af65a69017eda959 && __obf_af65a69017eda959 <= 122:
		return __obf_af65a69017eda959 - 32
	case 65 <= __obf_af65a69017eda959 && __obf_af65a69017eda959 <= 90:
		return __obf_af65a69017eda959 + 32
	default:
		return __obf_af65a69017eda959
	}
}

func IndexString(__obf_db7cf183ccadbefa string, __obf_991c6908209e9369 rune, __obf_325d680a436661fb int) string {
	__obf_327403b74941ba85 := []rune(__obf_db7cf183ccadbefa)
	var __obf_9f429377bc76d224 bytes.Buffer
	var __obf_550d8612d04f5667 int
	for __obf_7cd64e5d17b513e6, __obf_3f47a2a6cbc52427 := 0, len(__obf_327403b74941ba85); __obf_7cd64e5d17b513e6 < __obf_3f47a2a6cbc52427; __obf_7cd64e5d17b513e6++ {
		if __obf_327403b74941ba85[__obf_7cd64e5d17b513e6] == __obf_991c6908209e9369 {
			__obf_550d8612d04f5667 += 1
		}
		if __obf_550d8612d04f5667 == __obf_325d680a436661fb {
			break
		}
		__obf_9f429377bc76d224.
			WriteRune(__obf_327403b74941ba85[__obf_7cd64e5d17b513e6])
	}
	return __obf_9f429377bc76d224.String()
}

func LastIndexString(__obf_f7fcca366fccd9d2, __obf_c4b530ee77e93631 string) string {
	__obf_a9c9200b03759058 := strings.Split(__obf_f7fcca366fccd9d2, __obf_c4b530ee77e93631)
	if __obf_550d8612d04f5667 := len(__obf_a9c9200b03759058); __obf_550d8612d04f5667 > 1 {
		return __obf_a9c9200b03759058[__obf_550d8612d04f5667-2]
	}
	return ""
}

func IsEmpty(__obf_e138a084fc0b1411 any) bool {
	__obf_99fad6cd23d30931 := reflect.ValueOf(__obf_e138a084fc0b1411)
	if __obf_99fad6cd23d30931.Kind() == reflect.Ptr {
		__obf_99fad6cd23d30931 = __obf_99fad6cd23d30931.Elem()
	}
	return __obf_99fad6cd23d30931.Interface() == reflect.Zero(__obf_99fad6cd23d30931.Type()).Interface()
}

func MillisecondToDateString(__obf_6fc209c9da6e885e int64) string {
	return time.Unix(__obf_6fc209c9da6e885e, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_6fc209c9da6e885e int64) string {
	return time.Unix(__obf_6fc209c9da6e885e, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_7334ad1da5a6b0ba string) (__obf_a5ab9a007e4ee13b, __obf_03292ef9a6626456 string) {
	__obf_7334ad1da5a6b0ba = filepath.Base(__obf_7334ad1da5a6b0ba)
	__obf_03292ef9a6626456 = filepath.Ext(__obf_7334ad1da5a6b0ba)
	__obf_a5ab9a007e4ee13b = strings.TrimSuffix(__obf_7334ad1da5a6b0ba, __obf_03292ef9a6626456)
	return __obf_a5ab9a007e4ee13b,

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
		__obf_03292ef9a6626456
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(__obf_9a100b7bea67aa51 string) (int64, error) {
	if __obf_a6ecbcbe0bab3d3f, __obf_b93c04d14e5c503f := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_9a100b7bea67aa51, Loc); __obf_b93c04d14e5c503f != nil {
		return 0, nil
	} else {
		__obf_a6ecbcbe0bab3d3f = __obf_a6ecbcbe0bab3d3f.AddDate(0, 0, -__obf_a6ecbcbe0bab3d3f.Day()+1)
		return GetZeroTime(__obf_a6ecbcbe0bab3d3f).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_9a100b7bea67aa51 string) (int64, int64, error) {
	if __obf_a6ecbcbe0bab3d3f, __obf_b93c04d14e5c503f := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_9a100b7bea67aa51, Loc); __obf_b93c04d14e5c503f != nil {
		return 0, 0, nil
	} else {
		__obf_a6ecbcbe0bab3d3f = __obf_a6ecbcbe0bab3d3f.AddDate(0, 0, -__obf_a6ecbcbe0bab3d3f.Day()+1)
		return GetZeroTime(__obf_a6ecbcbe0bab3d3f).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_a6ecbcbe0bab3d3f).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_6ea3b2d60a754d22 time.Time) time.Time {
	return time.Date(__obf_6ea3b2d60a754d22.Year(), __obf_6ea3b2d60a754d22.Month(), __obf_6ea3b2d60a754d22.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_6fc209c9da6e885e int64) int64 {
	__obf_0d88914889a7ae68 := time.Unix(__obf_6fc209c9da6e885e, 0)
	__obf_522d092ea1cebb9f, __obf_7385df5914c40b35, __obf_6ea3b2d60a754d22 := __obf_0d88914889a7ae68.Date()
	return time.Date(__obf_522d092ea1cebb9f, __obf_7385df5914c40b35, __obf_6ea3b2d60a754d22, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_6fc209c9da6e885e int64) int64 {
	__obf_0d88914889a7ae68 := time.Unix(__obf_6fc209c9da6e885e, 0)
	return __obf_0d88914889a7ae68.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_522d092ea1cebb9f, __obf_7385df5914c40b35, __obf_6ea3b2d60a754d22 := time.Now().Date()
	return time.Date(__obf_522d092ea1cebb9f, __obf_7385df5914c40b35, __obf_6ea3b2d60a754d22, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_9a100b7bea67aa51 string) (int64, int64) {
	__obf_7902b4035cb96c5a, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_9a100b7bea67aa51+" 00:00:00", Loc)
	min := __obf_7902b4035cb96c5a.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_aabb9850a989068e int) (int64, int64) {
	__obf_9a100b7bea67aa51 := time.Now()
	__obf_0d88914889a7ae68 := __obf_9a100b7bea67aa51.AddDate(0, 0, __obf_aabb9850a989068e)
	__obf_522d092ea1cebb9f, __obf_7385df5914c40b35, __obf_6ea3b2d60a754d22 := __obf_0d88914889a7ae68.Date()
	min := time.Date(__obf_522d092ea1cebb9f, __obf_7385df5914c40b35, __obf_6ea3b2d60a754d22, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_9a100b7bea67aa51.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_df8f4f570e6f0155, __obf_c01945cae81f90dd string) bool {
	__obf_19d89907187ffe7e := func(__obf_a9c9200b03759058 string) string {
		return fmt.Sprintf(",%s,", __obf_a9c9200b03759058)
	}
	return strings.Contains(__obf_19d89907187ffe7e(__obf_df8f4f570e6f0155), __obf_19d89907187ffe7e(__obf_c01945cae81f90dd))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_56faa3389f87d5d9 string, __obf_ff9bd987644c25c3 ...string) bool {
	if any {
		for _, __obf_24a198bd8c59d97f := range __obf_ff9bd987644c25c3 {
			if Contain(__obf_56faa3389f87d5d9, __obf_24a198bd8c59d97f) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_24a198bd8c59d97f := range __obf_ff9bd987644c25c3 {
			if !Contain(__obf_56faa3389f87d5d9, __obf_24a198bd8c59d97f) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_a9c9200b03759058 []any, __obf_db0a1d5b34e76561 int) []any {
	return slices.Delete(__obf_a9c9200b03759058, __obf_db0a1d5b34e76561, __obf_db0a1d5b34e76561+1)
}

func String2Int8(__obf_db7cf183ccadbefa string) int8 {
	__obf_70b73c82a66e7803, __obf_b93c04d14e5c503f := strconv.ParseInt(__obf_db7cf183ccadbefa, 10, 8)
	if __obf_b93c04d14e5c503f == nil {
		return int8(__obf_70b73c82a66e7803)
	}
	return 0
}

func String2Int32(__obf_db7cf183ccadbefa string) int32 {
	__obf_70b73c82a66e7803, __obf_b93c04d14e5c503f := strconv.ParseInt(__obf_db7cf183ccadbefa, 10, 32)
	if __obf_b93c04d14e5c503f == nil {
		return int32(__obf_70b73c82a66e7803)
	}
	return 0
}

func String2Int64(__obf_db7cf183ccadbefa string) int8 {
	__obf_70b73c82a66e7803, __obf_b93c04d14e5c503f := strconv.ParseInt(__obf_db7cf183ccadbefa, 10, 64)
	if __obf_b93c04d14e5c503f == nil {
		return int8(__obf_70b73c82a66e7803)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_968bfd24570c9e4f, __obf_bb62b27fc6949736 int) (__obf_e14e1a6f3f62b9ca, __obf_9693f4b33d8e4720 time.Time) {
	if __obf_968bfd24570c9e4f > 0 && __obf_bb62b27fc6949736 > 0 {
		__obf_e14e1a6f3f62b9ca = time.Date(__obf_968bfd24570c9e4f, 1, 0, 0, 0, 0, 0, Loc)
		__obf_95c443f5ea661756 := // 第一天是周几
			int(__obf_e14e1a6f3f62b9ca.AddDate(0, 0, 1).Weekday())
		__obf_a4512d2dec2774db := // 当年第一周有几天
			1
		if __obf_95c443f5ea661756 != 0 {
			__obf_a4512d2dec2774db = 7 - __obf_95c443f5ea661756 + 1
		}
		if __obf_bb62b27fc6949736 == 1 {
			__obf_9693f4b33d8e4720 = __obf_e14e1a6f3f62b9ca.AddDate(0, 0, __obf_a4512d2dec2774db)
		} else {
			__obf_9693f4b33d8e4720 = __obf_e14e1a6f3f62b9ca.AddDate(0, 0, __obf_a4512d2dec2774db+(__obf_bb62b27fc6949736-1)*7)
			__obf_e14e1a6f3f62b9ca = __obf_9693f4b33d8e4720.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_9a100b7bea67aa51 time.Time) (__obf_968bfd24570c9e4f, __obf_bb62b27fc6949736, __obf_5805677c75e8cc73 int) {
	__obf_968bfd24570c9e4f = __obf_9a100b7bea67aa51.Year()
	__obf_5805677c75e8cc73 = int(__obf_9a100b7bea67aa51.Weekday())
	__obf_0d6527c9b2d7ed70 := __obf_9a100b7bea67aa51.YearDay()
	__obf_fe87dc7297b389e4 := __obf_9a100b7bea67aa51.AddDate(0, 0, -__obf_0d6527c9b2d7ed70+1)
	__obf_95c443f5ea661756 := int(__obf_fe87dc7297b389e4.Weekday())
	__obf_a4512d2dec2774db := // 当年第一周有几天
		1
	if __obf_95c443f5ea661756 != 0 {
		__obf_a4512d2dec2774db = 7 - __obf_95c443f5ea661756 + 1
	}
	if __obf_0d6527c9b2d7ed70 <= __obf_a4512d2dec2774db {
		__obf_bb62b27fc6949736 = 1
	} else {
		__obf_bb62b27fc6949736 = (__obf_0d6527c9b2d7ed70-__obf_a4512d2dec2774db)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_20542c6d126443f0 []byte) map[string]string {
	__obf_6dd1f4a4d9667146 := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_cfcc5fffe1599d23 := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_c6ab0f7e490eb768 []byte
	if __obf_c6ab0f7e490eb768 = __obf_6dd1f4a4d9667146.ReplaceAll(__obf_20542c6d126443f0, []byte("")); len(__obf_c6ab0f7e490eb768) < 6 {
		return nil
	}
	__obf_9ea53a08e18f0c98 := __obf_cfcc5fffe1599d23.FindAllSubmatch(__obf_c6ab0f7e490eb768[6:len(__obf_c6ab0f7e490eb768)-7], -1)
	__obf_150ab95fa43a1756 := map[string]string{}
	if __obf_abf083c5b2db9a29 := __obf_6dd1f4a4d9667146.FindSubmatch(__obf_20542c6d126443f0)[1]; len(__obf_abf083c5b2db9a29) != 0 {
		__obf_e82b8d42891ab355 := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_abf083c5b2db9a29, -1)
		for _, __obf_99fad6cd23d30931 := range __obf_e82b8d42891ab355 {
			__obf_150ab95fa43a1756[string(__obf_99fad6cd23d30931[1])] = string(__obf_99fad6cd23d30931[2])
		}
	}

	for _, __obf_99fad6cd23d30931 := range __obf_9ea53a08e18f0c98 {
		__obf_150ab95fa43a1756[string(__obf_99fad6cd23d30931[1])] = string(__obf_99fad6cd23d30931[2])
	}
	return __obf_150ab95fa43a1756
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_bb094403b67a0653 time.Time) (int64, int64) {
	__obf_0e58ede3954a23f3 := GetZeroTime(__obf_bb094403b67a0653).Unix() / 600
	return __obf_0e58ede3954a23f3, __obf_0e58ede3954a23f3 + 143
}

func Abs(__obf_550d8612d04f5667 int64) int64 {
	__obf_522d092ea1cebb9f := __obf_550d8612d04f5667 >> 63
	return (__obf_550d8612d04f5667 ^ __obf_522d092ea1cebb9f) - __obf_522d092ea1cebb9f
}

func Number2String(__obf_550d8612d04f5667 any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_550d8612d04f5667 := __obf_550d8612d04f5667.(type) {
	case int:
		return strconv.Itoa(__obf_550d8612d04f5667)
	case int32:
		return strconv.FormatInt(int64(__obf_550d8612d04f5667), 10)
	case int64:
		return strconv.FormatInt(__obf_550d8612d04f5667, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_550d8612d04f5667), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_550d8612d04f5667, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_db7cf183ccadbefa any, __obf_991c6908209e9369 string) string {
	if __obf_db7cf183ccadbefa != nil && __obf_db7cf183ccadbefa.(string) != "" {
		// return sep + str + sep
		return __obf_db7cf183ccadbefa.(string)
	}
	return __obf_991c6908209e9369
}

func SortRange(__obf_7385df5914c40b35 map[string]any, __obf_34a93e1eca3762b2 func(int, string)) {
	var __obf_6500911ba3ec32ef []string
	for __obf_ccfb9e5b04e5d653 := range __obf_7385df5914c40b35 {
		__obf_6500911ba3ec32ef = append(__obf_6500911ba3ec32ef, __obf_ccfb9e5b04e5d653)
	}
	sort.Strings(__obf_6500911ba3ec32ef)
	for __obf_7cd64e5d17b513e6, __obf_ccfb9e5b04e5d653 := range __obf_6500911ba3ec32ef {
		__obf_34a93e1eca3762b2(__obf_7cd64e5d17b513e6, __obf_ccfb9e5b04e5d653)
	}
}

func HasField(__obf_4fbfecd6f5873e53 reflect.Value, __obf_a5ab9a007e4ee13b string) bool {

	if __obf_a9c9200b03759058 := __obf_4fbfecd6f5873e53.FieldByNameFunc(func(__obf_550d8612d04f5667 string) bool {
		return strings.EqualFold(__obf_550d8612d04f5667, __obf_a5ab9a007e4ee13b)
	}); __obf_a9c9200b03759058.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_4fbfecd6f5873e53 reflect.Value, __obf_a5ab9a007e4ee13b string) reflect.Value {
	if __obf_a9c9200b03759058 := __obf_4fbfecd6f5873e53.FieldByNameFunc(func(__obf_550d8612d04f5667 string) bool {
		return strings.EqualFold(__obf_550d8612d04f5667, __obf_a5ab9a007e4ee13b)
	}); __obf_a9c9200b03759058.IsValid() {
		return __obf_a9c9200b03759058
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_4fbfecd6f5873e53 reflect.Value, __obf_a5ab9a007e4ee13b string, __obf_aae3b6f1d4dc6396 any) bool {

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
	if __obf_a9c9200b03759058 := __obf_4fbfecd6f5873e53.FieldByNameFunc(func(__obf_550d8612d04f5667 string) bool {
		return strings.EqualFold(__obf_550d8612d04f5667, __obf_a5ab9a007e4ee13b)
	}); __obf_a9c9200b03759058.IsValid() {
		if __obf_4222bb1aae8e3238 := __obf_a9c9200b03759058.Type(); __obf_4222bb1aae8e3238 == reflect.TypeOf(__obf_aae3b6f1d4dc6396) {
			__obf_a9c9200b03759058.
				Set(reflect.ValueOf(__obf_aae3b6f1d4dc6396))
		} else {
			__obf_a9c9200b03759058.
				Set(reflect.ValueOf(__obf_aae3b6f1d4dc6396).Convert(__obf_4222bb1aae8e3238))
		}
		return true
	}
	return false
}

func CopyMap(__obf_7385df5914c40b35 map[string]any) map[string]any {
	__obf_e82b8d42891ab355 := make(map[string]any)
	for __obf_ccfb9e5b04e5d653, __obf_99fad6cd23d30931 := range __obf_7385df5914c40b35 {
		if __obf_29dae70446e56336, __obf_faea3f7f8b9b1fca := __obf_99fad6cd23d30931.(map[string]any); __obf_faea3f7f8b9b1fca {
			__obf_e82b8d42891ab355[__obf_ccfb9e5b04e5d653] = CopyMap(__obf_29dae70446e56336)
		} else {
			__obf_e82b8d42891ab355[__obf_ccfb9e5b04e5d653] = __obf_99fad6cd23d30931
		}
	}

	return __obf_e82b8d42891ab355
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_88ef59ecffef7a5b, __obf_a47b74f2a6ceb840 string, __obf_090faa7bcdaa57cc bool, __obf_c246f4d56531b17b ...any) {
	var __obf_26eca02e820f7a05 strings.Builder
	var __obf_848afd96ea64cc2e string
	var __obf_6b83773df08ead52 []string
	for _, __obf_845665d36914759a := range __obf_c246f4d56531b17b {
		__obf_a9c9200b03759058 := reflect.TypeOf(__obf_845665d36914759a)
		__obf_26eca02e820f7a05.
			WriteString(`CREATE TABLE `)
		__obf_26eca02e820f7a05.
			WriteString(__obf_a9c9200b03759058.Name())
		__obf_26eca02e820f7a05.
			WriteString(" (\n")
		__obf_a312094afb0f9c12 := __obf_a9c9200b03759058.NumField()
		for __obf_7cd64e5d17b513e6 := 0; __obf_7cd64e5d17b513e6 < __obf_a312094afb0f9c12; __obf_7cd64e5d17b513e6++ {
			__obf_26eca02e820f7a05.
				WriteString("    ")
			__obf_6b83773df08ead52 = nil
			if __obf_70e31da1ed254365 := string(__obf_a9c9200b03759058.Field(__obf_7cd64e5d17b513e6).Tag.Get(__obf_88ef59ecffef7a5b)); __obf_70e31da1ed254365 == "" {
				if __obf_090faa7bcdaa57cc {
					__obf_848afd96ea64cc2e = ToCamel(__obf_a9c9200b03759058.Field(__obf_7cd64e5d17b513e6).Name)
				} else {
					__obf_848afd96ea64cc2e = __obf_a9c9200b03759058.Field(__obf_7cd64e5d17b513e6).Name
				}
			} else {
				__obf_6b83773df08ead52 = strings.Split(__obf_70e31da1ed254365, __obf_a47b74f2a6ceb840)
				__obf_848afd96ea64cc2e = __obf_6b83773df08ead52[0]
			}
			__obf_26eca02e820f7a05.
				WriteString(__obf_848afd96ea64cc2e)
			__obf_26eca02e820f7a05.
				WriteString(" ")
			switch __obf_a9c9200b03759058.Field(__obf_7cd64e5d17b513e6).Type.Name() {
			case "int8":
				__obf_26eca02e820f7a05.
					WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_26eca02e820f7a05.
					WriteString("INT")
			case "int64":
				__obf_26eca02e820f7a05.
					WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_26eca02e820f7a05.
					WriteString("VARCHAR(50)")
			}

			if len(__obf_6b83773df08ead52) > 1 {
				__obf_26eca02e820f7a05.
					WriteString(" ")
				__obf_26eca02e820f7a05.
					WriteString(strings.Join(__obf_6b83773df08ead52[1:], " "))
			}

			if __obf_7cd64e5d17b513e6+1 != __obf_a312094afb0f9c12 {
				__obf_26eca02e820f7a05.
					WriteString(",")
			}
			__obf_26eca02e820f7a05.
				WriteString("\n")
		}
		__obf_26eca02e820f7a05.
			WriteString(");\n\n")
	}
	fmt.Println(__obf_26eca02e820f7a05.String())
}

func SaveImage(__obf_63d6317535e4f76a image.Image, __obf_0fc1396105fd157a uint, __obf_6012080f068353fe string) error {
	__obf_146921211cb71664, __obf_b93c04d14e5c503f := os.Create(__obf_6012080f068353fe)
	if __obf_b93c04d14e5c503f != nil {
		return __obf_b93c04d14e5c503f
	}
	defer __obf_146921211cb71664.Close()
	return jpeg.Encode(__obf_146921211cb71664, resize.Resize(__obf_0fc1396105fd157a, 0, __obf_63d6317535e4f76a, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_db7cf183ccadbefa string) bool {
	return strings.Contains(__obf_db7cf183ccadbefa, ".") && net.ParseIP(__obf_db7cf183ccadbefa) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_db7cf183ccadbefa string) bool {
	return strings.Contains(__obf_db7cf183ccadbefa, ":") && net.ParseIP(__obf_db7cf183ccadbefa) != nil
}

func AnyToBytes(__obf_99fad6cd23d30931 any) ([]byte, error) {
	return msgpack.Marshal(__obf_99fad6cd23d30931)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_aae3b6f1d4dc6396 []byte) (__obf_1e49a709823fcdda T, __obf_b93c04d14e5c503f error) {
	__obf_b93c04d14e5c503f = msgpack.Unmarshal(__obf_aae3b6f1d4dc6396, &__obf_1e49a709823fcdda)
	return
}

func Loadyaml(__obf_8fc23618e16a26b2 string, __obf_d44218ba8addea2d any) {
	__obf_c01945cae81f90dd, __obf_b93c04d14e5c503f := os.ReadFile(__obf_8fc23618e16a26b2)
	if __obf_b93c04d14e5c503f != nil {
		log.Fatalln(__obf_b93c04d14e5c503f)
	}
	__obf_b93c04d14e5c503f = yaml.UnmarshalStrict(__obf_c01945cae81f90dd, __obf_d44218ba8addea2d)
	if __obf_b93c04d14e5c503f != nil {
		log.Fatalln(__obf_b93c04d14e5c503f)
	}
}

func ToAnyList[T any](__obf_92d1693a77de7b25 []T) []any {
	__obf_6849fbc26f98886a := make([]any, len(__obf_92d1693a77de7b25))
	for __obf_7cd64e5d17b513e6, __obf_99fad6cd23d30931 := range __obf_92d1693a77de7b25 {
		__obf_6849fbc26f98886a[__obf_7cd64e5d17b513e6] = __obf_99fad6cd23d30931
	}
	return __obf_6849fbc26f98886a
}

func TimeParse(__obf_0ead2aab6221d9d9 string) time.Time {
	__obf_0ead2aab6221d9d9 = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_0ead2aab6221d9d9)
	__obf_0d88914889a7ae68, _ := time.ParseInLocation("2006-01-02 15:04", __obf_0ead2aab6221d9d9, time.Local)
	return __obf_0d88914889a7ae68
}
