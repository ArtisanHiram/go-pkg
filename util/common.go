package __obf_813b65e0329aecbf

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
	__obf_006321724ad61b24 = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_30267a212643ccd8 any) ([]byte, error) {
	return __obf_006321724ad61b24.Marshal(__obf_30267a212643ccd8)
}

func EncodeString(__obf_30267a212643ccd8 any) string {
	if __obf_1ca36eea06f99b46, __obf_54101b1325683820 := __obf_006321724ad61b24.Marshal(__obf_30267a212643ccd8); __obf_54101b1325683820 == nil {
		return string(__obf_1ca36eea06f99b46)
	}
	return ""
}

func Decode(__obf_b4cea1e6bfe66ebc string, __obf_30267a212643ccd8 any) error {
	return __obf_006321724ad61b24.UnmarshalFromString(__obf_b4cea1e6bfe66ebc, __obf_30267a212643ccd8)
}

func DecodeByte(__obf_243fff1d60c134b7 []byte, __obf_30267a212643ccd8 any) error {
	return __obf_006321724ad61b24.Unmarshal(__obf_243fff1d60c134b7, __obf_30267a212643ccd8)
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

func PrefixInArray(__obf_12bd4ec2dc5e8341 string, __obf_5db463f850b93b74 []string) bool {
	for __obf_0e59933515fc0d1b := range __obf_5db463f850b93b74 {
		if strings.HasPrefix(__obf_12bd4ec2dc5e8341, __obf_5db463f850b93b74[__obf_0e59933515fc0d1b]) {
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

var __obf_50efa867b83d75f0 = time.Now().UnixNano()

// GenRandom 获得随机字符串
func GenRandom(__obf_a42bf05a2506c569 int) string {
	bytes := []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	__obf_0aa760b4b83b758d := []byte{}
	__obf_138ad5c24fa33841 := rand.New(rand.NewSource(__obf_50efa867b83d75f0))
	for range __obf_a42bf05a2506c569 {
		__obf_0aa760b4b83b758d = append(__obf_0aa760b4b83b758d, bytes[__obf_138ad5c24fa33841.Intn(len(bytes))])
	}
	return string(__obf_0aa760b4b83b758d)
}

func RemoveIndex(__obf_3b0cfd8394bed5fb []string, __obf_3115636dc95f69a7 int) []string {
	return append(__obf_3b0cfd8394bed5fb[:__obf_3115636dc95f69a7], __obf_3b0cfd8394bed5fb[__obf_3115636dc95f69a7+1:]...)
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
func ToCamel(__obf_b4cea1e6bfe66ebc string) (__obf_0aa760b4b83b758d string) {
	__obf_e3c267d9c9848351 := []rune(__obf_b4cea1e6bfe66ebc)
	__obf_0aa760b4b83b758d = __obf_b4cea1e6bfe66ebc[0:1]
	if __obf_e3c267d9c9848351[0] >= 97 && __obf_e3c267d9c9848351[0] <= 122 {
		__obf_0aa760b4b83b758d = string(__obf_e3c267d9c9848351[0] - 32)
	}

	__obf_a42bf05a2506c569 := len(__obf_e3c267d9c9848351)
	for __obf_0e59933515fc0d1b := 1; __obf_0e59933515fc0d1b < __obf_a42bf05a2506c569; __obf_0e59933515fc0d1b++ {
		if __obf_e3c267d9c9848351[__obf_0e59933515fc0d1b] == 95 && __obf_e3c267d9c9848351[__obf_0e59933515fc0d1b+1] >= 97 && __obf_e3c267d9c9848351[__obf_0e59933515fc0d1b+1] <= 122 { //过滤下划线
			__obf_e3c267d9c9848351[__obf_0e59933515fc0d1b+1] -= 32
		} else {
			__obf_0aa760b4b83b758d += string(__obf_e3c267d9c9848351[__obf_0e59933515fc0d1b])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_b4cea1e6bfe66ebc string) (__obf_0aa760b4b83b758d string) {
	__obf_e3c267d9c9848351 := []rune(__obf_b4cea1e6bfe66ebc)
	__obf_0aa760b4b83b758d = __obf_b4cea1e6bfe66ebc[0:1]
	if __obf_e3c267d9c9848351[0] >= 65 && __obf_e3c267d9c9848351[0] <= 90 {
		__obf_0aa760b4b83b758d = string(__obf_e3c267d9c9848351[0] + 32)
	}

	__obf_2c6b6b0f61b708f9 := len(__obf_e3c267d9c9848351)
	for __obf_0e59933515fc0d1b := 1; __obf_0e59933515fc0d1b < __obf_2c6b6b0f61b708f9; __obf_0e59933515fc0d1b++ {
		if __obf_e3c267d9c9848351[__obf_0e59933515fc0d1b] >= 65 && __obf_e3c267d9c9848351[__obf_0e59933515fc0d1b] <= 90 { //大写变小写
			__obf_e3c267d9c9848351[__obf_0e59933515fc0d1b] += 32
			__obf_0aa760b4b83b758d += "_"
		}
		__obf_0aa760b4b83b758d += string(__obf_e3c267d9c9848351[__obf_0e59933515fc0d1b])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_b4cea1e6bfe66ebc string, __obf_08983e6a5ae33637 int, __obf_2c6b6b0f61b708f9 int) string {
	// 将字符串转换为rune切片，以正确处理多字节字符
	__obf_9e605cb2d7e06da4 := []rune(__obf_b4cea1e6bfe66ebc)
	__obf_d38fc4dc3ee3e9cd := len(__obf_9e605cb2d7e06da4)

	// 处理n为负数或0的无效情况
	if __obf_2c6b6b0f61b708f9 <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_08983e6a5ae33637 < 0 {
		__obf_08983e6a5ae33637 = __obf_d38fc4dc3ee3e9cd + __obf_08983e6a5ae33637
	}

	// 边界检查：确保start在有效范围内
	if __obf_08983e6a5ae33637 < 0 || __obf_08983e6a5ae33637 >= __obf_d38fc4dc3ee3e9cd {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_9e605cb2d7e06da4[__obf_08983e6a5ae33637:min(__obf_08983e6a5ae33637+__obf_2c6b6b0f61b708f9, __obf_d38fc4dc3ee3e9cd)])
}

func ASCII(__obf_138ad5c24fa33841 rune) rune {
	switch {
	case 97 <= __obf_138ad5c24fa33841 && __obf_138ad5c24fa33841 <= 122:
		return __obf_138ad5c24fa33841 - 32
	case 65 <= __obf_138ad5c24fa33841 && __obf_138ad5c24fa33841 <= 90:
		return __obf_138ad5c24fa33841 + 32
	default:
		return __obf_138ad5c24fa33841
	}
}

func IndexString(__obf_b4cea1e6bfe66ebc string, __obf_ea8c99a9b7b45ad7 rune, __obf_f82af4a4520be29d int) string {
	__obf_689adb7165737f05 := []rune(__obf_b4cea1e6bfe66ebc)
	var __obf_d7599df897b8418f bytes.Buffer
	var __obf_2c6b6b0f61b708f9 int
	for __obf_0e59933515fc0d1b, __obf_1034afbba6fe168e := 0, len(__obf_689adb7165737f05); __obf_0e59933515fc0d1b < __obf_1034afbba6fe168e; __obf_0e59933515fc0d1b++ {
		if __obf_689adb7165737f05[__obf_0e59933515fc0d1b] == __obf_ea8c99a9b7b45ad7 {
			__obf_2c6b6b0f61b708f9 += 1
		}
		if __obf_2c6b6b0f61b708f9 == __obf_f82af4a4520be29d {
			break
		}
		__obf_d7599df897b8418f.WriteRune(__obf_689adb7165737f05[__obf_0e59933515fc0d1b])
	}
	return __obf_d7599df897b8418f.String()
}

func LastIndexString(__obf_36f511fc7589036d, __obf_a1cb2c06dba4a4ba string) string {
	__obf_3b0cfd8394bed5fb := strings.Split(__obf_36f511fc7589036d, __obf_a1cb2c06dba4a4ba)
	if __obf_2c6b6b0f61b708f9 := len(__obf_3b0cfd8394bed5fb); __obf_2c6b6b0f61b708f9 > 1 {
		return __obf_3b0cfd8394bed5fb[__obf_2c6b6b0f61b708f9-2]
	}
	return ""
}

func IsEmpty(__obf_a0052bfb73bc2626 any) bool {
	__obf_2dfe4b44e04cc8de := reflect.ValueOf(__obf_a0052bfb73bc2626)
	if __obf_2dfe4b44e04cc8de.Kind() == reflect.Ptr {
		__obf_2dfe4b44e04cc8de = __obf_2dfe4b44e04cc8de.Elem()
	}
	return __obf_2dfe4b44e04cc8de.Interface() == reflect.Zero(__obf_2dfe4b44e04cc8de.Type()).Interface()
}

func MillisecondToDateString(__obf_832fddf98e38c8bd int64) string {
	return time.Unix(__obf_832fddf98e38c8bd, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_832fddf98e38c8bd int64) string {
	return time.Unix(__obf_832fddf98e38c8bd, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_965820536ba77e1c string) (__obf_794ec8367f651eb7, __obf_7018c93569bd8dcc string) {
	__obf_965820536ba77e1c = filepath.Base(__obf_965820536ba77e1c)
	__obf_7018c93569bd8dcc = filepath.Ext(__obf_965820536ba77e1c)
	__obf_794ec8367f651eb7 = strings.TrimSuffix(__obf_965820536ba77e1c, __obf_7018c93569bd8dcc)
	return __obf_794ec8367f651eb7, __obf_7018c93569bd8dcc
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
func GetFirstDateOfMonth(__obf_79dfacec4dccf298 string) (int64, error) {
	if __obf_eb498caf3bdfabea, __obf_54101b1325683820 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_79dfacec4dccf298, Loc); __obf_54101b1325683820 != nil {
		return 0, nil
	} else {
		__obf_eb498caf3bdfabea = __obf_eb498caf3bdfabea.AddDate(0, 0, -__obf_eb498caf3bdfabea.Day()+1)
		return GetZeroTime(__obf_eb498caf3bdfabea).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_79dfacec4dccf298 string) (int64, int64, error) {
	if __obf_eb498caf3bdfabea, __obf_54101b1325683820 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_79dfacec4dccf298, Loc); __obf_54101b1325683820 != nil {
		return 0, 0, nil
	} else {
		__obf_eb498caf3bdfabea = __obf_eb498caf3bdfabea.AddDate(0, 0, -__obf_eb498caf3bdfabea.Day()+1)
		return GetZeroTime(__obf_eb498caf3bdfabea).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_eb498caf3bdfabea).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_44b44226049e1da4 time.Time) time.Time {
	return time.Date(__obf_44b44226049e1da4.Year(), __obf_44b44226049e1da4.Month(), __obf_44b44226049e1da4.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_832fddf98e38c8bd int64) int64 {
	__obf_f8bbb4ad6d831056 := time.Unix(__obf_832fddf98e38c8bd, 0)
	__obf_657b6f43e1c8acad, __obf_5964dd98fffcb9f3, __obf_44b44226049e1da4 := __obf_f8bbb4ad6d831056.Date()
	return time.Date(__obf_657b6f43e1c8acad, __obf_5964dd98fffcb9f3, __obf_44b44226049e1da4, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_832fddf98e38c8bd int64) int64 {
	__obf_f8bbb4ad6d831056 := time.Unix(__obf_832fddf98e38c8bd, 0)
	return __obf_f8bbb4ad6d831056.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_657b6f43e1c8acad, __obf_5964dd98fffcb9f3, __obf_44b44226049e1da4 := time.Now().Date()
	return time.Date(__obf_657b6f43e1c8acad, __obf_5964dd98fffcb9f3, __obf_44b44226049e1da4, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_79dfacec4dccf298 string) (int64, int64) {
	__obf_b4efed7a717da46e, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_79dfacec4dccf298+" 00:00:00", Loc)
	min := __obf_b4efed7a717da46e.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_00e936f025aac32c int) (int64, int64) {
	__obf_79dfacec4dccf298 := time.Now()
	__obf_f8bbb4ad6d831056 := __obf_79dfacec4dccf298.AddDate(0, 0, __obf_00e936f025aac32c)
	__obf_657b6f43e1c8acad, __obf_5964dd98fffcb9f3, __obf_44b44226049e1da4 := __obf_f8bbb4ad6d831056.Date()
	min := time.Date(__obf_657b6f43e1c8acad, __obf_5964dd98fffcb9f3, __obf_44b44226049e1da4, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_79dfacec4dccf298.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_8c5741ba25882166, __obf_5ddb24a8c119c8f3 string) bool {
	__obf_c6daebe0ed645fe1 := func(__obf_3b0cfd8394bed5fb string) string {
		return fmt.Sprintf(",%s,", __obf_3b0cfd8394bed5fb)
	}
	return strings.Contains(__obf_c6daebe0ed645fe1(__obf_8c5741ba25882166), __obf_c6daebe0ed645fe1(__obf_5ddb24a8c119c8f3))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_e7bcf7e17b1c81ef string, __obf_a1ec806ff9b79c28 ...string) bool {
	if any {
		for _, __obf_fc343dd99409448f := range __obf_a1ec806ff9b79c28 {
			if Contain(__obf_e7bcf7e17b1c81ef, __obf_fc343dd99409448f) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_fc343dd99409448f := range __obf_a1ec806ff9b79c28 {
			if !Contain(__obf_e7bcf7e17b1c81ef, __obf_fc343dd99409448f) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_3b0cfd8394bed5fb []any, __obf_3115636dc95f69a7 int) []any {
	return slices.Delete(__obf_3b0cfd8394bed5fb, __obf_3115636dc95f69a7, __obf_3115636dc95f69a7+1)
}

func String2Int8(__obf_b4cea1e6bfe66ebc string) int8 {
	__obf_5ea4835781916df5, __obf_54101b1325683820 := strconv.ParseInt(__obf_b4cea1e6bfe66ebc, 10, 8)
	if __obf_54101b1325683820 == nil {
		return int8(__obf_5ea4835781916df5)
	}
	return 0
}

func String2Int32(__obf_b4cea1e6bfe66ebc string) int32 {
	__obf_5ea4835781916df5, __obf_54101b1325683820 := strconv.ParseInt(__obf_b4cea1e6bfe66ebc, 10, 32)
	if __obf_54101b1325683820 == nil {
		return int32(__obf_5ea4835781916df5)
	}
	return 0
}

func String2Int64(__obf_b4cea1e6bfe66ebc string) int8 {
	__obf_5ea4835781916df5, __obf_54101b1325683820 := strconv.ParseInt(__obf_b4cea1e6bfe66ebc, 10, 64)
	if __obf_54101b1325683820 == nil {
		return int8(__obf_5ea4835781916df5)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_d1f313c046d571ac, __obf_eca361ae9c975a47 int) (__obf_08983e6a5ae33637, __obf_ae945ab2670ea1ba time.Time) {
	if __obf_d1f313c046d571ac > 0 && __obf_eca361ae9c975a47 > 0 {
		__obf_08983e6a5ae33637 = time.Date(__obf_d1f313c046d571ac, 1, 0, 0, 0, 0, 0, Loc)
		// 第一天是周几
		__obf_884e099f204779de := int(__obf_08983e6a5ae33637.AddDate(0, 0, 1).Weekday())
		// 当年第一周有几天
		__obf_1dabbff8910ce740 := 1
		if __obf_884e099f204779de != 0 {
			__obf_1dabbff8910ce740 = 7 - __obf_884e099f204779de + 1
		}
		if __obf_eca361ae9c975a47 == 1 {
			__obf_ae945ab2670ea1ba = __obf_08983e6a5ae33637.AddDate(0, 0, __obf_1dabbff8910ce740)
		} else {
			__obf_ae945ab2670ea1ba = __obf_08983e6a5ae33637.AddDate(0, 0, __obf_1dabbff8910ce740+(__obf_eca361ae9c975a47-1)*7)
			__obf_08983e6a5ae33637 = __obf_ae945ab2670ea1ba.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_79dfacec4dccf298 time.Time) (__obf_d1f313c046d571ac, __obf_eca361ae9c975a47, __obf_0000832e7b664061 int) {
	__obf_d1f313c046d571ac = __obf_79dfacec4dccf298.Year()
	__obf_0000832e7b664061 = int(__obf_79dfacec4dccf298.Weekday())
	__obf_36822964ec44a1d5 := __obf_79dfacec4dccf298.YearDay()
	__obf_381304047efb4987 := __obf_79dfacec4dccf298.AddDate(0, 0, -__obf_36822964ec44a1d5+1)
	__obf_884e099f204779de := int(__obf_381304047efb4987.Weekday())
	// 当年第一周有几天
	__obf_1dabbff8910ce740 := 1
	if __obf_884e099f204779de != 0 {
		__obf_1dabbff8910ce740 = 7 - __obf_884e099f204779de + 1
	}
	if __obf_36822964ec44a1d5 <= __obf_1dabbff8910ce740 {
		__obf_eca361ae9c975a47 = 1
	} else {
		__obf_eca361ae9c975a47 = (__obf_36822964ec44a1d5-__obf_1dabbff8910ce740)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_d96770f955ae576e []byte) map[string]string {
	__obf_e59cae2fda8ae3e6 := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_dcc28264ae7b30fa := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_b7fa727bbd78163e []byte
	if __obf_b7fa727bbd78163e = __obf_e59cae2fda8ae3e6.ReplaceAll(__obf_d96770f955ae576e, []byte("")); len(__obf_b7fa727bbd78163e) < 6 {
		return nil
	}

	__obf_62fb8424e69827f0 := __obf_dcc28264ae7b30fa.FindAllSubmatch(__obf_b7fa727bbd78163e[6:len(__obf_b7fa727bbd78163e)-7], -1)
	__obf_8e7481452a149366 := map[string]string{}
	if __obf_c6715213a0957dd5 := __obf_e59cae2fda8ae3e6.FindSubmatch(__obf_d96770f955ae576e)[1]; len(__obf_c6715213a0957dd5) != 0 {
		__obf_04714e13b3070ece := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_c6715213a0957dd5, -1)
		for _, __obf_2dfe4b44e04cc8de := range __obf_04714e13b3070ece {
			__obf_8e7481452a149366[string(__obf_2dfe4b44e04cc8de[1])] = string(__obf_2dfe4b44e04cc8de[2])
		}
	}

	for _, __obf_2dfe4b44e04cc8de := range __obf_62fb8424e69827f0 {
		__obf_8e7481452a149366[string(__obf_2dfe4b44e04cc8de[1])] = string(__obf_2dfe4b44e04cc8de[2])
	}
	return __obf_8e7481452a149366
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_93814122851407a3 time.Time) (int64, int64) {
	__obf_4af796bc9d4d79d5 := GetZeroTime(__obf_93814122851407a3).Unix() / 600
	return __obf_4af796bc9d4d79d5, __obf_4af796bc9d4d79d5 + 143
}

func Abs(__obf_2c6b6b0f61b708f9 int64) int64 {
	__obf_657b6f43e1c8acad := __obf_2c6b6b0f61b708f9 >> 63
	return (__obf_2c6b6b0f61b708f9 ^ __obf_657b6f43e1c8acad) - __obf_657b6f43e1c8acad
}

func Number2String(__obf_2c6b6b0f61b708f9 any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_2c6b6b0f61b708f9 := __obf_2c6b6b0f61b708f9.(type) {
	case int:
		return strconv.Itoa(__obf_2c6b6b0f61b708f9)
	case int32:
		return strconv.FormatInt(int64(__obf_2c6b6b0f61b708f9), 10)
	case int64:
		return strconv.FormatInt(__obf_2c6b6b0f61b708f9, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_2c6b6b0f61b708f9), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_2c6b6b0f61b708f9, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_b4cea1e6bfe66ebc any, __obf_ea8c99a9b7b45ad7 string) string {
	if __obf_b4cea1e6bfe66ebc != nil && __obf_b4cea1e6bfe66ebc.(string) != "" {
		// return sep + str + sep
		return __obf_b4cea1e6bfe66ebc.(string)
	}
	return __obf_ea8c99a9b7b45ad7
}

func SortRange(__obf_5964dd98fffcb9f3 map[string]any, __obf_49f027f6b6c03687 func(int, string)) {
	var __obf_f4ac6abfa1ce2dd6 []string
	for __obf_12f59ca81d2169f7 := range __obf_5964dd98fffcb9f3 {
		__obf_f4ac6abfa1ce2dd6 = append(__obf_f4ac6abfa1ce2dd6, __obf_12f59ca81d2169f7)
	}
	sort.Strings(__obf_f4ac6abfa1ce2dd6)
	for __obf_0e59933515fc0d1b, __obf_12f59ca81d2169f7 := range __obf_f4ac6abfa1ce2dd6 {
		__obf_49f027f6b6c03687(__obf_0e59933515fc0d1b, __obf_12f59ca81d2169f7)
	}
}

func HasField(__obf_b3e8a2f7ed12a09a reflect.Value, __obf_794ec8367f651eb7 string) bool {

	if __obf_3b0cfd8394bed5fb := __obf_b3e8a2f7ed12a09a.FieldByNameFunc(func(__obf_2c6b6b0f61b708f9 string) bool {
		return strings.EqualFold(__obf_2c6b6b0f61b708f9, __obf_794ec8367f651eb7)
	}); __obf_3b0cfd8394bed5fb.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_b3e8a2f7ed12a09a reflect.Value, __obf_794ec8367f651eb7 string) reflect.Value {
	if __obf_3b0cfd8394bed5fb := __obf_b3e8a2f7ed12a09a.FieldByNameFunc(func(__obf_2c6b6b0f61b708f9 string) bool {
		return strings.EqualFold(__obf_2c6b6b0f61b708f9, __obf_794ec8367f651eb7)
	}); __obf_3b0cfd8394bed5fb.IsValid() {
		return __obf_3b0cfd8394bed5fb
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_b3e8a2f7ed12a09a reflect.Value, __obf_794ec8367f651eb7 string, __obf_12bd4ec2dc5e8341 any) bool {

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
	if __obf_3b0cfd8394bed5fb := __obf_b3e8a2f7ed12a09a.FieldByNameFunc(func(__obf_2c6b6b0f61b708f9 string) bool {
		return strings.EqualFold(__obf_2c6b6b0f61b708f9, __obf_794ec8367f651eb7)
	}); __obf_3b0cfd8394bed5fb.IsValid() {
		if __obf_d5eaab91e3686e67 := __obf_3b0cfd8394bed5fb.Type(); __obf_d5eaab91e3686e67 == reflect.TypeOf(__obf_12bd4ec2dc5e8341) {
			__obf_3b0cfd8394bed5fb.Set(reflect.ValueOf(__obf_12bd4ec2dc5e8341))
		} else {
			__obf_3b0cfd8394bed5fb.Set(reflect.ValueOf(__obf_12bd4ec2dc5e8341).Convert(__obf_d5eaab91e3686e67))
		}
		return true
	}
	return false
}

func CopyMap(__obf_5964dd98fffcb9f3 map[string]any) map[string]any {
	__obf_04714e13b3070ece := make(map[string]any)
	for __obf_12f59ca81d2169f7, __obf_2dfe4b44e04cc8de := range __obf_5964dd98fffcb9f3 {
		if __obf_32667feeeae2bccb, __obf_f5eff000e29052bc := __obf_2dfe4b44e04cc8de.(map[string]any); __obf_f5eff000e29052bc {
			__obf_04714e13b3070ece[__obf_12f59ca81d2169f7] = CopyMap(__obf_32667feeeae2bccb)
		} else {
			__obf_04714e13b3070ece[__obf_12f59ca81d2169f7] = __obf_2dfe4b44e04cc8de
		}
	}

	return __obf_04714e13b3070ece
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_3610eaa580a77615, __obf_975c81dd7540d49b string, __obf_225ccc6630ccb7d4 bool, __obf_52c55f1023dc619c ...any) {
	var __obf_efec31cff5a1de41 strings.Builder
	var __obf_2c87f76bb1a9dad2 string
	var __obf_c9871a7a12b62fa1 []string
	for _, __obf_80ba63e5ffd3b5b7 := range __obf_52c55f1023dc619c {
		__obf_3b0cfd8394bed5fb := reflect.TypeOf(__obf_80ba63e5ffd3b5b7)
		__obf_efec31cff5a1de41.WriteString(`CREATE TABLE `)
		__obf_efec31cff5a1de41.WriteString(__obf_3b0cfd8394bed5fb.Name())
		__obf_efec31cff5a1de41.WriteString(" (\n")
		__obf_c242410c9053bac2 := __obf_3b0cfd8394bed5fb.NumField()
		for __obf_0e59933515fc0d1b := 0; __obf_0e59933515fc0d1b < __obf_c242410c9053bac2; __obf_0e59933515fc0d1b++ {
			__obf_efec31cff5a1de41.WriteString("    ")
			__obf_c9871a7a12b62fa1 = nil
			if __obf_8f7055e828f0ec42 := string(__obf_3b0cfd8394bed5fb.Field(__obf_0e59933515fc0d1b).Tag.Get(__obf_3610eaa580a77615)); __obf_8f7055e828f0ec42 == "" {
				if __obf_225ccc6630ccb7d4 {
					__obf_2c87f76bb1a9dad2 = ToCamel(__obf_3b0cfd8394bed5fb.Field(__obf_0e59933515fc0d1b).Name)
				} else {
					__obf_2c87f76bb1a9dad2 = __obf_3b0cfd8394bed5fb.Field(__obf_0e59933515fc0d1b).Name
				}
			} else {
				__obf_c9871a7a12b62fa1 = strings.Split(__obf_8f7055e828f0ec42, __obf_975c81dd7540d49b)
				__obf_2c87f76bb1a9dad2 = __obf_c9871a7a12b62fa1[0]
			}
			__obf_efec31cff5a1de41.WriteString(__obf_2c87f76bb1a9dad2)
			__obf_efec31cff5a1de41.WriteString(" ")
			switch __obf_3b0cfd8394bed5fb.Field(__obf_0e59933515fc0d1b).Type.Name() {
			case "int8":
				__obf_efec31cff5a1de41.WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_efec31cff5a1de41.WriteString("INT")
			case "int64":
				__obf_efec31cff5a1de41.WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_efec31cff5a1de41.WriteString("VARCHAR(50)")
			}

			if len(__obf_c9871a7a12b62fa1) > 1 {
				__obf_efec31cff5a1de41.WriteString(" ")
				__obf_efec31cff5a1de41.WriteString(strings.Join(__obf_c9871a7a12b62fa1[1:], " "))
			}

			if __obf_0e59933515fc0d1b+1 != __obf_c242410c9053bac2 {
				__obf_efec31cff5a1de41.WriteString(",")
			}
			__obf_efec31cff5a1de41.WriteString("\n")
		}
		__obf_efec31cff5a1de41.WriteString(");\n\n")
	}
	fmt.Println(__obf_efec31cff5a1de41.String())
}

func SaveImage(__obf_e94e7ecd04636acf image.Image, __obf_1695017cbaa232e4 uint, __obf_6adfe472ddfbb40a string) error {
	__obf_89f7b6b34ee5d2a1, __obf_54101b1325683820 := os.Create(__obf_6adfe472ddfbb40a)
	if __obf_54101b1325683820 != nil {
		return __obf_54101b1325683820
	}
	defer __obf_89f7b6b34ee5d2a1.Close()
	return jpeg.Encode(__obf_89f7b6b34ee5d2a1, resize.Resize(__obf_1695017cbaa232e4, 0, __obf_e94e7ecd04636acf, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_b4cea1e6bfe66ebc string) bool {
	return strings.Contains(__obf_b4cea1e6bfe66ebc, ".") && net.ParseIP(__obf_b4cea1e6bfe66ebc) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_b4cea1e6bfe66ebc string) bool {
	return strings.Contains(__obf_b4cea1e6bfe66ebc, ":") && net.ParseIP(__obf_b4cea1e6bfe66ebc) != nil
}

func AnyToBytes(__obf_2dfe4b44e04cc8de any) ([]byte, error) {
	return msgpack.Marshal(__obf_2dfe4b44e04cc8de)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_12bd4ec2dc5e8341 []byte) (__obf_30267a212643ccd8 T, __obf_54101b1325683820 error) {
	__obf_54101b1325683820 = msgpack.Unmarshal(__obf_12bd4ec2dc5e8341, &__obf_30267a212643ccd8)
	return
}

func Loadyaml(__obf_140ad13c43198ecd string, __obf_07477fb8448749e9 any) {
	__obf_5ddb24a8c119c8f3, __obf_54101b1325683820 := os.ReadFile(__obf_140ad13c43198ecd)
	if __obf_54101b1325683820 != nil {
		log.Fatalln(__obf_54101b1325683820)
	}
	__obf_54101b1325683820 = yaml.UnmarshalStrict(__obf_5ddb24a8c119c8f3, __obf_07477fb8448749e9)
	if __obf_54101b1325683820 != nil {
		log.Fatalln(__obf_54101b1325683820)
	}
}

func ToAnyList[T any](__obf_5c5b650afabd1f25 []T) []any {
	__obf_5809bf6741492ebb := make([]any, len(__obf_5c5b650afabd1f25))
	for __obf_0e59933515fc0d1b, __obf_2dfe4b44e04cc8de := range __obf_5c5b650afabd1f25 {
		__obf_5809bf6741492ebb[__obf_0e59933515fc0d1b] = __obf_2dfe4b44e04cc8de
	}
	return __obf_5809bf6741492ebb
}

func TimeParse(__obf_d3ea24771c4e7656 string) time.Time {
	__obf_d3ea24771c4e7656 = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_d3ea24771c4e7656)
	__obf_f8bbb4ad6d831056, _ := time.ParseInLocation("2006-01-02 15:04", __obf_d3ea24771c4e7656, time.Local)
	return __obf_f8bbb4ad6d831056
}
