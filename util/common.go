package __obf_083c8deafa73f533

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
	yaml "gopkg.in/yaml.v2"
)

var (
	Loc, _                 = time.LoadLocation("Asia/Shanghai")
	__obf_283076f580f99f88 = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_9cf562b2e3294b16 any) ([]byte, error) {
	return __obf_283076f580f99f88.Marshal(__obf_9cf562b2e3294b16)
}

func EncodeString(__obf_9cf562b2e3294b16 any) string {
	if __obf_71e3b50a7a885567, __obf_ab078048114898aa := __obf_283076f580f99f88.Marshal(__obf_9cf562b2e3294b16); __obf_ab078048114898aa == nil {
		return string(__obf_71e3b50a7a885567)
	}
	return ""
}

func Decode(__obf_6bd9d99d372748de string, __obf_9cf562b2e3294b16 any) error {
	return __obf_283076f580f99f88.UnmarshalFromString(__obf_6bd9d99d372748de, __obf_9cf562b2e3294b16)
}

func DecodeByte(__obf_b1e85a310b2a8288 []byte, __obf_9cf562b2e3294b16 any) error {
	return __obf_283076f580f99f88.Unmarshal(__obf_b1e85a310b2a8288, __obf_9cf562b2e3294b16)
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

func PrefixInArray(__obf_739cf405cd68540d string, __obf_9f213c7be5779923 []string) bool {
	for __obf_444299a68c7d4da5 := range __obf_9f213c7be5779923 {
		if strings.HasPrefix(__obf_739cf405cd68540d, __obf_9f213c7be5779923[__obf_444299a68c7d4da5]) {
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

var __obf_2d3f59469f1b5efd = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var __obf_f5d6a86ee1482d66 = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenRandom 获得随机字符串
func GenRandom(__obf_d5bf5930f448d12c int) string {
	__obf_b1e85a310b2a8288 := make([]byte, __obf_d5bf5930f448d12c)
	for __obf_444299a68c7d4da5 := range __obf_b1e85a310b2a8288 {
		__obf_b1e85a310b2a8288[__obf_444299a68c7d4da5] = __obf_2d3f59469f1b5efd[__obf_f5d6a86ee1482d66.Intn(len(__obf_2d3f59469f1b5efd))]
	}
	return string(__obf_b1e85a310b2a8288)
}

func RemoveIndex(__obf_116809f90a65bc1a []string, __obf_75a85ed31fa6b562 int) []string {
	return append(__obf_116809f90a65bc1a[:__obf_75a85ed31fa6b562], __obf_116809f90a65bc1a[__obf_75a85ed31fa6b562+1:]...)
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
func ToCamel(__obf_6bd9d99d372748de string) (__obf_fae510501b39e125 string) {
	__obf_2d34f2deef5a5092 := []rune(__obf_6bd9d99d372748de)
	__obf_fae510501b39e125 = __obf_6bd9d99d372748de[0:1]
	if __obf_2d34f2deef5a5092[0] >= 97 && __obf_2d34f2deef5a5092[0] <= 122 {
		__obf_fae510501b39e125 = string(__obf_2d34f2deef5a5092[0] - 32)
	}

	__obf_2d9138371caba2c0 := len(__obf_2d34f2deef5a5092)
	for __obf_444299a68c7d4da5 := 1; __obf_444299a68c7d4da5 < __obf_2d9138371caba2c0; __obf_444299a68c7d4da5++ {
		if __obf_2d34f2deef5a5092[__obf_444299a68c7d4da5] == 95 && __obf_2d34f2deef5a5092[__obf_444299a68c7d4da5+1] >= 97 && __obf_2d34f2deef5a5092[__obf_444299a68c7d4da5+1] <= 122 { //过滤下划线
			__obf_2d34f2deef5a5092[__obf_444299a68c7d4da5+1] -= 32
		} else {
			__obf_fae510501b39e125 += string(__obf_2d34f2deef5a5092[__obf_444299a68c7d4da5])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_6bd9d99d372748de string) (__obf_fae510501b39e125 string) {
	__obf_2d34f2deef5a5092 := []rune(__obf_6bd9d99d372748de)
	__obf_fae510501b39e125 = __obf_6bd9d99d372748de[0:1]
	if __obf_2d34f2deef5a5092[0] >= 65 && __obf_2d34f2deef5a5092[0] <= 90 {
		__obf_fae510501b39e125 = string(__obf_2d34f2deef5a5092[0] + 32)
	}

	__obf_d5bf5930f448d12c := len(__obf_2d34f2deef5a5092)
	for __obf_444299a68c7d4da5 := 1; __obf_444299a68c7d4da5 < __obf_d5bf5930f448d12c; __obf_444299a68c7d4da5++ {
		if __obf_2d34f2deef5a5092[__obf_444299a68c7d4da5] >= 65 && __obf_2d34f2deef5a5092[__obf_444299a68c7d4da5] <= 90 { //大写变小写
			__obf_2d34f2deef5a5092[__obf_444299a68c7d4da5] += 32
			__obf_fae510501b39e125 += "_"
		}
		__obf_fae510501b39e125 += string(__obf_2d34f2deef5a5092[__obf_444299a68c7d4da5])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_6bd9d99d372748de string, __obf_29f7dda144cd58fd int, __obf_d5bf5930f448d12c int) string {
	// 将字符串转换为rune切片，以正确处理多字节字符
	__obf_8d507cebae11c9a0 := []rune(__obf_6bd9d99d372748de)
	__obf_48154335eab5cfd3 := len(__obf_8d507cebae11c9a0)

	// 处理n为负数或0的无效情况
	if __obf_d5bf5930f448d12c <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_29f7dda144cd58fd < 0 {
		__obf_29f7dda144cd58fd = __obf_48154335eab5cfd3 + __obf_29f7dda144cd58fd
	}

	// 边界检查：确保start在有效范围内
	if __obf_29f7dda144cd58fd < 0 || __obf_29f7dda144cd58fd >= __obf_48154335eab5cfd3 {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_8d507cebae11c9a0[__obf_29f7dda144cd58fd:min(__obf_29f7dda144cd58fd+__obf_d5bf5930f448d12c, __obf_48154335eab5cfd3)])
}

func ASCII(__obf_c094a893233159c5 rune) rune {
	switch {
	case 97 <= __obf_c094a893233159c5 && __obf_c094a893233159c5 <= 122:
		return __obf_c094a893233159c5 - 32
	case 65 <= __obf_c094a893233159c5 && __obf_c094a893233159c5 <= 90:
		return __obf_c094a893233159c5 + 32
	default:
		return __obf_c094a893233159c5
	}
}

func IndexString(__obf_6bd9d99d372748de string, __obf_f80570083423b2d7 rune, __obf_bb9292c8c650d5da int) string {
	__obf_4e548d75f9c18617 := []rune(__obf_6bd9d99d372748de)
	var __obf_fcdc351f804b7c83 bytes.Buffer
	var __obf_d5bf5930f448d12c int
	for __obf_444299a68c7d4da5, __obf_caca1be9d402f5e9 := 0, len(__obf_4e548d75f9c18617); __obf_444299a68c7d4da5 < __obf_caca1be9d402f5e9; __obf_444299a68c7d4da5++ {
		if __obf_4e548d75f9c18617[__obf_444299a68c7d4da5] == __obf_f80570083423b2d7 {
			__obf_d5bf5930f448d12c += 1
		}
		if __obf_d5bf5930f448d12c == __obf_bb9292c8c650d5da {
			break
		}
		__obf_fcdc351f804b7c83.WriteRune(__obf_4e548d75f9c18617[__obf_444299a68c7d4da5])
	}
	return __obf_fcdc351f804b7c83.String()
}

func LastIndexString(__obf_641c50999baea69d, __obf_0fe2304383195552 string) string {
	__obf_116809f90a65bc1a := strings.Split(__obf_641c50999baea69d, __obf_0fe2304383195552)
	if __obf_d5bf5930f448d12c := len(__obf_116809f90a65bc1a); __obf_d5bf5930f448d12c > 1 {
		return __obf_116809f90a65bc1a[__obf_d5bf5930f448d12c-2]
	}
	return ""
}

func IsEmpty(__obf_5795522b6759d31f any) bool {
	__obf_bc97935e99909ed0 := reflect.ValueOf(__obf_5795522b6759d31f)
	if __obf_bc97935e99909ed0.Kind() == reflect.Ptr {
		__obf_bc97935e99909ed0 = __obf_bc97935e99909ed0.Elem()
	}
	return __obf_bc97935e99909ed0.Interface() == reflect.Zero(__obf_bc97935e99909ed0.Type()).Interface()
}

func MillisecondToDateString(__obf_846c2f29b699842b int64) string {
	return time.Unix(__obf_846c2f29b699842b, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_846c2f29b699842b int64) string {
	return time.Unix(__obf_846c2f29b699842b, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_3a0f763a7e299ba3 string) (__obf_018da406c2219d4e, __obf_016e92df841897ab string) {
	__obf_3a0f763a7e299ba3 = filepath.Base(__obf_3a0f763a7e299ba3)
	__obf_016e92df841897ab = filepath.Ext(__obf_3a0f763a7e299ba3)
	__obf_018da406c2219d4e = strings.TrimSuffix(__obf_3a0f763a7e299ba3, __obf_016e92df841897ab)
	return __obf_018da406c2219d4e, __obf_016e92df841897ab
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
func GetFirstDateOfMonth(__obf_28debac1646d6b76 string) (int64, error) {
	if __obf_0e7639b2451a2a94, __obf_ab078048114898aa := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_28debac1646d6b76, Loc); __obf_ab078048114898aa != nil {
		return 0, nil
	} else {
		__obf_0e7639b2451a2a94 = __obf_0e7639b2451a2a94.AddDate(0, 0, -__obf_0e7639b2451a2a94.Day()+1)
		return GetZeroTime(__obf_0e7639b2451a2a94).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_28debac1646d6b76 string) (int64, int64, error) {
	if __obf_0e7639b2451a2a94, __obf_ab078048114898aa := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_28debac1646d6b76, Loc); __obf_ab078048114898aa != nil {
		return 0, 0, nil
	} else {
		__obf_0e7639b2451a2a94 = __obf_0e7639b2451a2a94.AddDate(0, 0, -__obf_0e7639b2451a2a94.Day()+1)
		return GetZeroTime(__obf_0e7639b2451a2a94).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_0e7639b2451a2a94).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_66f560b1a3148c0c time.Time) time.Time {
	return time.Date(__obf_66f560b1a3148c0c.Year(), __obf_66f560b1a3148c0c.Month(), __obf_66f560b1a3148c0c.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_846c2f29b699842b int64) int64 {
	__obf_95a64e15cfc35236 := time.Unix(__obf_846c2f29b699842b, 0)
	__obf_df2b5a160989211b, __obf_4669ae2244ec2e3f, __obf_66f560b1a3148c0c := __obf_95a64e15cfc35236.Date()
	return time.Date(__obf_df2b5a160989211b, __obf_4669ae2244ec2e3f, __obf_66f560b1a3148c0c, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_846c2f29b699842b int64) int64 {
	__obf_95a64e15cfc35236 := time.Unix(__obf_846c2f29b699842b, 0)
	return __obf_95a64e15cfc35236.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_df2b5a160989211b, __obf_4669ae2244ec2e3f, __obf_66f560b1a3148c0c := time.Now().Date()
	return time.Date(__obf_df2b5a160989211b, __obf_4669ae2244ec2e3f, __obf_66f560b1a3148c0c, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_28debac1646d6b76 string) (int64, int64) {
	__obf_edf28eb0ee09b9ed, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_28debac1646d6b76+" 00:00:00", Loc)
	min := __obf_edf28eb0ee09b9ed.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_637003f8db19dec4 int) (int64, int64) {
	__obf_28debac1646d6b76 := time.Now()
	__obf_95a64e15cfc35236 := __obf_28debac1646d6b76.AddDate(0, 0, __obf_637003f8db19dec4)
	__obf_df2b5a160989211b, __obf_4669ae2244ec2e3f, __obf_66f560b1a3148c0c := __obf_95a64e15cfc35236.Date()
	min := time.Date(__obf_df2b5a160989211b, __obf_4669ae2244ec2e3f, __obf_66f560b1a3148c0c, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_28debac1646d6b76.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_6215ccc7dee16652, __obf_4ca251a835e266e7 string) bool {
	__obf_c9703863e2234b70 := func(__obf_116809f90a65bc1a string) string {
		return fmt.Sprintf(",%s,", __obf_116809f90a65bc1a)
	}
	return strings.Contains(__obf_c9703863e2234b70(__obf_6215ccc7dee16652), __obf_c9703863e2234b70(__obf_4ca251a835e266e7))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_92aa3e3ad38c2f14 string, __obf_572438716c1612d0 ...string) bool {
	if any {
		for _, __obf_90ec48ccc47fd0ef := range __obf_572438716c1612d0 {
			if Contain(__obf_92aa3e3ad38c2f14, __obf_90ec48ccc47fd0ef) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_90ec48ccc47fd0ef := range __obf_572438716c1612d0 {
			if !Contain(__obf_92aa3e3ad38c2f14, __obf_90ec48ccc47fd0ef) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_116809f90a65bc1a []any, __obf_75a85ed31fa6b562 int) []any {
	return slices.Delete(__obf_116809f90a65bc1a, __obf_75a85ed31fa6b562, __obf_75a85ed31fa6b562+1)
}

func String2Int8(__obf_6bd9d99d372748de string) int8 {
	__obf_e69f70c4c637db82, __obf_ab078048114898aa := strconv.ParseInt(__obf_6bd9d99d372748de, 10, 8)
	if __obf_ab078048114898aa == nil {
		return int8(__obf_e69f70c4c637db82)
	}
	return 0
}

func String2Int32(__obf_6bd9d99d372748de string) int32 {
	__obf_e69f70c4c637db82, __obf_ab078048114898aa := strconv.ParseInt(__obf_6bd9d99d372748de, 10, 32)
	if __obf_ab078048114898aa == nil {
		return int32(__obf_e69f70c4c637db82)
	}
	return 0
}

func String2Int64(__obf_6bd9d99d372748de string) int8 {
	__obf_e69f70c4c637db82, __obf_ab078048114898aa := strconv.ParseInt(__obf_6bd9d99d372748de, 10, 64)
	if __obf_ab078048114898aa == nil {
		return int8(__obf_e69f70c4c637db82)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_dc46f6172cafb849, __obf_3963b4ff5af56288 int) (__obf_29f7dda144cd58fd, __obf_110312f138e70cd2 time.Time) {
	if __obf_dc46f6172cafb849 > 0 && __obf_3963b4ff5af56288 > 0 {
		__obf_29f7dda144cd58fd = time.Date(__obf_dc46f6172cafb849, 1, 0, 0, 0, 0, 0, Loc)
		// 第一天是周几
		__obf_3653f16248fe7a02 := int(__obf_29f7dda144cd58fd.AddDate(0, 0, 1).Weekday())
		// 当年第一周有几天
		__obf_48b1bd304a110f7f := 1
		if __obf_3653f16248fe7a02 != 0 {
			__obf_48b1bd304a110f7f = 7 - __obf_3653f16248fe7a02 + 1
		}
		if __obf_3963b4ff5af56288 == 1 {
			__obf_110312f138e70cd2 = __obf_29f7dda144cd58fd.AddDate(0, 0, __obf_48b1bd304a110f7f)
		} else {
			__obf_110312f138e70cd2 = __obf_29f7dda144cd58fd.AddDate(0, 0, __obf_48b1bd304a110f7f+(__obf_3963b4ff5af56288-1)*7)
			__obf_29f7dda144cd58fd = __obf_110312f138e70cd2.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_28debac1646d6b76 time.Time) (__obf_dc46f6172cafb849, __obf_3963b4ff5af56288, __obf_1d0bab01ae8a3e63 int) {
	__obf_dc46f6172cafb849 = __obf_28debac1646d6b76.Year()
	__obf_1d0bab01ae8a3e63 = int(__obf_28debac1646d6b76.Weekday())
	__obf_d96e13c81f4a9733 := __obf_28debac1646d6b76.YearDay()
	__obf_c37e39854335d141 := __obf_28debac1646d6b76.AddDate(0, 0, -__obf_d96e13c81f4a9733+1)
	__obf_3653f16248fe7a02 := int(__obf_c37e39854335d141.Weekday())
	// 当年第一周有几天
	__obf_48b1bd304a110f7f := 1
	if __obf_3653f16248fe7a02 != 0 {
		__obf_48b1bd304a110f7f = 7 - __obf_3653f16248fe7a02 + 1
	}
	if __obf_d96e13c81f4a9733 <= __obf_48b1bd304a110f7f {
		__obf_3963b4ff5af56288 = 1
	} else {
		__obf_3963b4ff5af56288 = (__obf_d96e13c81f4a9733-__obf_48b1bd304a110f7f)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_ee7b252aee995dfe []byte) map[string]string {
	__obf_648b793921b3b85b := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_7a15a2bc7ed220d9 := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_ae9bd39d8d529ecc []byte
	if __obf_ae9bd39d8d529ecc = __obf_648b793921b3b85b.ReplaceAll(__obf_ee7b252aee995dfe, []byte("")); len(__obf_ae9bd39d8d529ecc) < 6 {
		return nil
	}

	__obf_48b14d384a77d133 := __obf_7a15a2bc7ed220d9.FindAllSubmatch(__obf_ae9bd39d8d529ecc[6:len(__obf_ae9bd39d8d529ecc)-7], -1)
	__obf_ecde037f3f1bf6ca := map[string]string{}
	if __obf_dcbb2d0631b6c243 := __obf_648b793921b3b85b.FindSubmatch(__obf_ee7b252aee995dfe)[1]; len(__obf_dcbb2d0631b6c243) != 0 {
		__obf_d860e982c2675848 := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_dcbb2d0631b6c243, -1)
		for _, __obf_bc97935e99909ed0 := range __obf_d860e982c2675848 {
			__obf_ecde037f3f1bf6ca[string(__obf_bc97935e99909ed0[1])] = string(__obf_bc97935e99909ed0[2])
		}
	}

	for _, __obf_bc97935e99909ed0 := range __obf_48b14d384a77d133 {
		__obf_ecde037f3f1bf6ca[string(__obf_bc97935e99909ed0[1])] = string(__obf_bc97935e99909ed0[2])
	}
	return __obf_ecde037f3f1bf6ca
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_47eddb7c1b7c7a75 time.Time) (int64, int64) {
	__obf_6ad2990c75a97ab3 := GetZeroTime(__obf_47eddb7c1b7c7a75).Unix() / 600
	return __obf_6ad2990c75a97ab3, __obf_6ad2990c75a97ab3 + 143
}

func Abs(__obf_d5bf5930f448d12c int64) int64 {
	__obf_df2b5a160989211b := __obf_d5bf5930f448d12c >> 63
	return (__obf_d5bf5930f448d12c ^ __obf_df2b5a160989211b) - __obf_df2b5a160989211b
}

func Number2String(__obf_d5bf5930f448d12c any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_d5bf5930f448d12c := __obf_d5bf5930f448d12c.(type) {
	case int:
		return strconv.Itoa(__obf_d5bf5930f448d12c)
	case int32:
		return strconv.FormatInt(int64(__obf_d5bf5930f448d12c), 10)
	case int64:
		return strconv.FormatInt(__obf_d5bf5930f448d12c, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_d5bf5930f448d12c), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_d5bf5930f448d12c, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_6bd9d99d372748de any, __obf_f80570083423b2d7 string) string {
	if __obf_6bd9d99d372748de != nil && __obf_6bd9d99d372748de.(string) != "" {
		// return sep + str + sep
		return __obf_6bd9d99d372748de.(string)
	}
	return __obf_f80570083423b2d7
}

func SortRange(__obf_4669ae2244ec2e3f map[string]any, __obf_2f555a3e46769181 func(int, string)) {
	var __obf_387eac597d6e2597 []string
	for __obf_4f6630328e99b83a := range __obf_4669ae2244ec2e3f {
		__obf_387eac597d6e2597 = append(__obf_387eac597d6e2597, __obf_4f6630328e99b83a)
	}
	sort.Strings(__obf_387eac597d6e2597)
	for __obf_444299a68c7d4da5, __obf_4f6630328e99b83a := range __obf_387eac597d6e2597 {
		__obf_2f555a3e46769181(__obf_444299a68c7d4da5, __obf_4f6630328e99b83a)
	}
}

func HasField(__obf_0a7cf6bab01114b0 reflect.Value, __obf_018da406c2219d4e string) bool {

	if __obf_116809f90a65bc1a := __obf_0a7cf6bab01114b0.FieldByNameFunc(func(__obf_d5bf5930f448d12c string) bool {
		return strings.EqualFold(__obf_d5bf5930f448d12c, __obf_018da406c2219d4e)
	}); __obf_116809f90a65bc1a.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_0a7cf6bab01114b0 reflect.Value, __obf_018da406c2219d4e string) reflect.Value {
	if __obf_116809f90a65bc1a := __obf_0a7cf6bab01114b0.FieldByNameFunc(func(__obf_d5bf5930f448d12c string) bool {
		return strings.EqualFold(__obf_d5bf5930f448d12c, __obf_018da406c2219d4e)
	}); __obf_116809f90a65bc1a.IsValid() {
		return __obf_116809f90a65bc1a
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_0a7cf6bab01114b0 reflect.Value, __obf_018da406c2219d4e string, __obf_739cf405cd68540d any) bool {

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
	if __obf_116809f90a65bc1a := __obf_0a7cf6bab01114b0.FieldByNameFunc(func(__obf_d5bf5930f448d12c string) bool {
		return strings.EqualFold(__obf_d5bf5930f448d12c, __obf_018da406c2219d4e)
	}); __obf_116809f90a65bc1a.IsValid() {
		if __obf_1bb1f83d5285f08a := __obf_116809f90a65bc1a.Type(); __obf_1bb1f83d5285f08a == reflect.TypeOf(__obf_739cf405cd68540d) {
			__obf_116809f90a65bc1a.Set(reflect.ValueOf(__obf_739cf405cd68540d))
		} else {
			__obf_116809f90a65bc1a.Set(reflect.ValueOf(__obf_739cf405cd68540d).Convert(__obf_1bb1f83d5285f08a))
		}
		return true
	}
	return false
}

func CopyMap(__obf_4669ae2244ec2e3f map[string]any) map[string]any {
	__obf_d860e982c2675848 := make(map[string]any)
	for __obf_4f6630328e99b83a, __obf_bc97935e99909ed0 := range __obf_4669ae2244ec2e3f {
		if __obf_2e46c8ca0f20a304, __obf_e2038eb4c2d652db := __obf_bc97935e99909ed0.(map[string]any); __obf_e2038eb4c2d652db {
			__obf_d860e982c2675848[__obf_4f6630328e99b83a] = CopyMap(__obf_2e46c8ca0f20a304)
		} else {
			__obf_d860e982c2675848[__obf_4f6630328e99b83a] = __obf_bc97935e99909ed0
		}
	}

	return __obf_d860e982c2675848
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_5e4139c6b01db295, __obf_5c683eee1ca9760f string, __obf_6e95edebc3dd01e6 bool, __obf_befee03d48538c4c ...any) {
	var __obf_c358e9b12f630ac9 strings.Builder
	var __obf_be40ef4f3c35643c string
	var __obf_b10a7c5a291c1214 []string
	for _, __obf_4782a38b05b82360 := range __obf_befee03d48538c4c {
		__obf_116809f90a65bc1a := reflect.TypeOf(__obf_4782a38b05b82360)
		__obf_c358e9b12f630ac9.WriteString(`CREATE TABLE `)
		__obf_c358e9b12f630ac9.WriteString(__obf_116809f90a65bc1a.Name())
		__obf_c358e9b12f630ac9.WriteString(" (\n")
		__obf_2eb3acbd73d3de8c := __obf_116809f90a65bc1a.NumField()
		for __obf_444299a68c7d4da5 := 0; __obf_444299a68c7d4da5 < __obf_2eb3acbd73d3de8c; __obf_444299a68c7d4da5++ {
			__obf_c358e9b12f630ac9.WriteString("    ")
			__obf_b10a7c5a291c1214 = nil
			if __obf_c0228d5c5424628e := string(__obf_116809f90a65bc1a.Field(__obf_444299a68c7d4da5).Tag.Get(__obf_5e4139c6b01db295)); __obf_c0228d5c5424628e == "" {
				if __obf_6e95edebc3dd01e6 {
					__obf_be40ef4f3c35643c = ToCamel(__obf_116809f90a65bc1a.Field(__obf_444299a68c7d4da5).Name)
				} else {
					__obf_be40ef4f3c35643c = __obf_116809f90a65bc1a.Field(__obf_444299a68c7d4da5).Name
				}
			} else {
				__obf_b10a7c5a291c1214 = strings.Split(__obf_c0228d5c5424628e, __obf_5c683eee1ca9760f)
				__obf_be40ef4f3c35643c = __obf_b10a7c5a291c1214[0]
			}
			__obf_c358e9b12f630ac9.WriteString(__obf_be40ef4f3c35643c)
			__obf_c358e9b12f630ac9.WriteString(" ")
			switch __obf_116809f90a65bc1a.Field(__obf_444299a68c7d4da5).Type.Name() {
			case "int8":
				__obf_c358e9b12f630ac9.WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_c358e9b12f630ac9.WriteString("INT")
			case "int64":
				__obf_c358e9b12f630ac9.WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_c358e9b12f630ac9.WriteString("VARCHAR(50)")
			}

			if len(__obf_b10a7c5a291c1214) > 1 {
				__obf_c358e9b12f630ac9.WriteString(" ")
				__obf_c358e9b12f630ac9.WriteString(strings.Join(__obf_b10a7c5a291c1214[1:], " "))
			}

			if __obf_444299a68c7d4da5+1 != __obf_2eb3acbd73d3de8c {
				__obf_c358e9b12f630ac9.WriteString(",")
			}
			__obf_c358e9b12f630ac9.WriteString("\n")
		}
		__obf_c358e9b12f630ac9.WriteString(");\n\n")
	}
	fmt.Println(__obf_c358e9b12f630ac9.String())
}

func SaveImage(__obf_4c842c5037c33e79 image.Image, __obf_8ea2c5197b468b20 uint, __obf_3fbe1bd9391f9a80 string) error {
	__obf_c81702feb6e0fc1c, __obf_ab078048114898aa := os.Create(__obf_3fbe1bd9391f9a80)
	if __obf_ab078048114898aa != nil {
		return __obf_ab078048114898aa
	}
	defer __obf_c81702feb6e0fc1c.Close()
	return jpeg.Encode(__obf_c81702feb6e0fc1c, resize.Resize(__obf_8ea2c5197b468b20, 0, __obf_4c842c5037c33e79, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_6bd9d99d372748de string) bool {
	return strings.Contains(__obf_6bd9d99d372748de, ".") && net.ParseIP(__obf_6bd9d99d372748de) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_6bd9d99d372748de string) bool {
	return strings.Contains(__obf_6bd9d99d372748de, ":") && net.ParseIP(__obf_6bd9d99d372748de) != nil
}

func AnyToBytes(__obf_bc97935e99909ed0 any) ([]byte, error) {
	return msgpack.Marshal(__obf_bc97935e99909ed0)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_739cf405cd68540d []byte) (__obf_9cf562b2e3294b16 T, __obf_ab078048114898aa error) {
	__obf_ab078048114898aa = msgpack.Unmarshal(__obf_739cf405cd68540d, &__obf_9cf562b2e3294b16)
	return
}

func Loadyaml(__obf_a6ecb2b8155710d2 string, __obf_40d1536861f2859f any) {
	__obf_4ca251a835e266e7, __obf_ab078048114898aa := os.ReadFile(__obf_a6ecb2b8155710d2)
	if __obf_ab078048114898aa != nil {
		log.Fatalln(__obf_ab078048114898aa)
	}
	__obf_ab078048114898aa = yaml.UnmarshalStrict(__obf_4ca251a835e266e7, __obf_40d1536861f2859f)
	if __obf_ab078048114898aa != nil {
		log.Fatalln(__obf_ab078048114898aa)
	}
}

func ToAnyList[T any](__obf_3db537b9c7db992a []T) []any {
	__obf_cda4827c2cfb820e := make([]any, len(__obf_3db537b9c7db992a))
	for __obf_444299a68c7d4da5, __obf_bc97935e99909ed0 := range __obf_3db537b9c7db992a {
		__obf_cda4827c2cfb820e[__obf_444299a68c7d4da5] = __obf_bc97935e99909ed0
	}
	return __obf_cda4827c2cfb820e
}

func TimeParse(__obf_613a063577a9e117 string) time.Time {
	__obf_613a063577a9e117 = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_613a063577a9e117)
	__obf_95a64e15cfc35236, _ := time.ParseInLocation("2006-01-02 15:04", __obf_613a063577a9e117, time.Local)
	return __obf_95a64e15cfc35236
}
