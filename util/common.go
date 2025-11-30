package __obf_69af70389b6622a3

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
	__obf_bd9eacaebdd04f5a = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_2801af39b525757c any) ([]byte, error) {
	return __obf_bd9eacaebdd04f5a.Marshal(__obf_2801af39b525757c)
}

func EncodeString(__obf_2801af39b525757c any) string {
	if __obf_335770de14796486, __obf_e812cfd1203ed4f3 := __obf_bd9eacaebdd04f5a.Marshal(__obf_2801af39b525757c); __obf_e812cfd1203ed4f3 == nil {
		return string(__obf_335770de14796486)
	}
	return ""
}

func Decode(__obf_f5886586954a86fc string, __obf_2801af39b525757c any) error {
	return __obf_bd9eacaebdd04f5a.UnmarshalFromString(__obf_f5886586954a86fc, __obf_2801af39b525757c)
}

func DecodeByte(__obf_55286557d75f8aca []byte, __obf_2801af39b525757c any) error {
	return __obf_bd9eacaebdd04f5a.Unmarshal(__obf_55286557d75f8aca, __obf_2801af39b525757c)
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

func PrefixInArray(__obf_9050e2c8d15b9a49 string, __obf_a8eda86efbef790d []string) bool {
	for __obf_962fdbe99abe4185 := range __obf_a8eda86efbef790d {
		if strings.HasPrefix(__obf_9050e2c8d15b9a49, __obf_a8eda86efbef790d[__obf_962fdbe99abe4185]) {
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

var __obf_68bd14ca736bfb71 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenRandom 获得随机字符串
func GenRandom(__obf_5b1c1c5422ae12c5 int) string {
	__obf_55286557d75f8aca := make([]byte, __obf_5b1c1c5422ae12c5)
	for __obf_962fdbe99abe4185 := range __obf_55286557d75f8aca {
		__obf_55286557d75f8aca[__obf_962fdbe99abe4185] = __obf_68bd14ca736bfb71[rand.IntN(len(__obf_68bd14ca736bfb71))]
	}
	return string(__obf_55286557d75f8aca)
}

func RemoveIndex(__obf_5331eba287d21166 []string, __obf_374ed272b06bc80f int) []string {
	return append(__obf_5331eba287d21166[:__obf_374ed272b06bc80f], __obf_5331eba287d21166[__obf_374ed272b06bc80f+1:]...)
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
func ToCamel(__obf_f5886586954a86fc string) (__obf_f7c4b5f8a87c287c string) {
	__obf_abec29f1b4efada5 := []rune(__obf_f5886586954a86fc)
	__obf_f7c4b5f8a87c287c = __obf_f5886586954a86fc[0:1]
	if __obf_abec29f1b4efada5[0] >= 97 && __obf_abec29f1b4efada5[0] <= 122 {
		__obf_f7c4b5f8a87c287c = string(__obf_abec29f1b4efada5[0] - 32)
	}
	__obf_bd52445dbe90030c := len(__obf_abec29f1b4efada5)
	for __obf_962fdbe99abe4185 := 1; __obf_962fdbe99abe4185 < __obf_bd52445dbe90030c; __obf_962fdbe99abe4185++ {
		if __obf_abec29f1b4efada5[__obf_962fdbe99abe4185] == 95 && __obf_abec29f1b4efada5[__obf_962fdbe99abe4185+1] >= 97 && __obf_abec29f1b4efada5[__obf_962fdbe99abe4185+1] <= 122 {
			__obf_abec29f1b4efada5[ //过滤下划线
			__obf_962fdbe99abe4185+1] -= 32
		} else {
			__obf_f7c4b5f8a87c287c += string(__obf_abec29f1b4efada5[__obf_962fdbe99abe4185])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_f5886586954a86fc string) (__obf_f7c4b5f8a87c287c string) {
	__obf_abec29f1b4efada5 := []rune(__obf_f5886586954a86fc)
	__obf_f7c4b5f8a87c287c = __obf_f5886586954a86fc[0:1]
	if __obf_abec29f1b4efada5[0] >= 65 && __obf_abec29f1b4efada5[0] <= 90 {
		__obf_f7c4b5f8a87c287c = string(__obf_abec29f1b4efada5[0] + 32)
	}
	__obf_5b1c1c5422ae12c5 := len(__obf_abec29f1b4efada5)
	for __obf_962fdbe99abe4185 := 1; __obf_962fdbe99abe4185 < __obf_5b1c1c5422ae12c5; __obf_962fdbe99abe4185++ {
		if __obf_abec29f1b4efada5[__obf_962fdbe99abe4185] >= 65 && __obf_abec29f1b4efada5[__obf_962fdbe99abe4185] <= 90 {
			__obf_abec29f1b4efada5[ //大写变小写
			__obf_962fdbe99abe4185] += 32
			__obf_f7c4b5f8a87c287c += "_"
		}
		__obf_f7c4b5f8a87c287c += string(__obf_abec29f1b4efada5[__obf_962fdbe99abe4185])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_f5886586954a86fc string, __obf_ace71eb4992cd39a int, __obf_5b1c1c5422ae12c5 int) string {
	__obf_c1d462f8a1cf81b3 := // 将字符串转换为rune切片，以正确处理多字节字符
		[]rune(__obf_f5886586954a86fc)
	__obf_9020047b0bb52202 := len(__obf_c1d462f8a1cf81b3)

	// 处理n为负数或0的无效情况
	if __obf_5b1c1c5422ae12c5 <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_ace71eb4992cd39a < 0 {
		__obf_ace71eb4992cd39a = __obf_9020047b0bb52202 + __obf_ace71eb4992cd39a
	}

	// 边界检查：确保start在有效范围内
	if __obf_ace71eb4992cd39a < 0 || __obf_ace71eb4992cd39a >= __obf_9020047b0bb52202 {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_c1d462f8a1cf81b3[__obf_ace71eb4992cd39a:min(__obf_ace71eb4992cd39a+__obf_5b1c1c5422ae12c5, __obf_9020047b0bb52202)])
}

func ASCII(__obf_675dd561f1bbbbff rune) rune {
	switch {
	case 97 <= __obf_675dd561f1bbbbff && __obf_675dd561f1bbbbff <= 122:
		return __obf_675dd561f1bbbbff - 32
	case 65 <= __obf_675dd561f1bbbbff && __obf_675dd561f1bbbbff <= 90:
		return __obf_675dd561f1bbbbff + 32
	default:
		return __obf_675dd561f1bbbbff
	}
}

func IndexString(__obf_f5886586954a86fc string, __obf_5b24539688807f6c rune, __obf_3de4900122b4e648 int) string {
	__obf_c0f01b651627a935 := []rune(__obf_f5886586954a86fc)
	var __obf_81432d275fcb9c98 bytes.Buffer
	var __obf_5b1c1c5422ae12c5 int
	for __obf_962fdbe99abe4185, __obf_2d66d34609586e65 := 0, len(__obf_c0f01b651627a935); __obf_962fdbe99abe4185 < __obf_2d66d34609586e65; __obf_962fdbe99abe4185++ {
		if __obf_c0f01b651627a935[__obf_962fdbe99abe4185] == __obf_5b24539688807f6c {
			__obf_5b1c1c5422ae12c5 += 1
		}
		if __obf_5b1c1c5422ae12c5 == __obf_3de4900122b4e648 {
			break
		}
		__obf_81432d275fcb9c98.
			WriteRune(__obf_c0f01b651627a935[__obf_962fdbe99abe4185])
	}
	return __obf_81432d275fcb9c98.String()
}

func LastIndexString(__obf_6fb8112d553c288e, __obf_4246f670989ac9f3 string) string {
	__obf_5331eba287d21166 := strings.Split(__obf_6fb8112d553c288e, __obf_4246f670989ac9f3)
	if __obf_5b1c1c5422ae12c5 := len(__obf_5331eba287d21166); __obf_5b1c1c5422ae12c5 > 1 {
		return __obf_5331eba287d21166[__obf_5b1c1c5422ae12c5-2]
	}
	return ""
}

func IsEmpty(__obf_2afbf591382ea161 any) bool {
	__obf_b1dd41bcd3423053 := reflect.ValueOf(__obf_2afbf591382ea161)
	if __obf_b1dd41bcd3423053.Kind() == reflect.Ptr {
		__obf_b1dd41bcd3423053 = __obf_b1dd41bcd3423053.Elem()
	}
	return __obf_b1dd41bcd3423053.Interface() == reflect.Zero(__obf_b1dd41bcd3423053.Type()).Interface()
}

func MillisecondToDateString(__obf_62b38bc1ef31c0de int64) string {
	return time.Unix(__obf_62b38bc1ef31c0de, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_62b38bc1ef31c0de int64) string {
	return time.Unix(__obf_62b38bc1ef31c0de, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_9d8bcc9943fd7a19 string) (__obf_04d9a2c744f64089, __obf_4c8e5967c43667cd string) {
	__obf_9d8bcc9943fd7a19 = filepath.Base(__obf_9d8bcc9943fd7a19)
	__obf_4c8e5967c43667cd = filepath.Ext(__obf_9d8bcc9943fd7a19)
	__obf_04d9a2c744f64089 = strings.TrimSuffix(__obf_9d8bcc9943fd7a19, __obf_4c8e5967c43667cd)
	return __obf_04d9a2c744f64089,

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
		__obf_4c8e5967c43667cd
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(__obf_1b56f2857581b365 string) (int64, error) {
	if __obf_1036735836ccf8a0, __obf_e812cfd1203ed4f3 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_1b56f2857581b365, Loc); __obf_e812cfd1203ed4f3 != nil {
		return 0, nil
	} else {
		__obf_1036735836ccf8a0 = __obf_1036735836ccf8a0.AddDate(0, 0, -__obf_1036735836ccf8a0.Day()+1)
		return GetZeroTime(__obf_1036735836ccf8a0).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_1b56f2857581b365 string) (int64, int64, error) {
	if __obf_1036735836ccf8a0, __obf_e812cfd1203ed4f3 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_1b56f2857581b365, Loc); __obf_e812cfd1203ed4f3 != nil {
		return 0, 0, nil
	} else {
		__obf_1036735836ccf8a0 = __obf_1036735836ccf8a0.AddDate(0, 0, -__obf_1036735836ccf8a0.Day()+1)
		return GetZeroTime(__obf_1036735836ccf8a0).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_1036735836ccf8a0).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_3cbd1ac0144c98a4 time.Time) time.Time {
	return time.Date(__obf_3cbd1ac0144c98a4.Year(), __obf_3cbd1ac0144c98a4.Month(), __obf_3cbd1ac0144c98a4.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_62b38bc1ef31c0de int64) int64 {
	__obf_9152cf6c829268b7 := time.Unix(__obf_62b38bc1ef31c0de, 0)
	__obf_285c0277db982a36, __obf_4680696b4dfeb451, __obf_3cbd1ac0144c98a4 := __obf_9152cf6c829268b7.Date()
	return time.Date(__obf_285c0277db982a36, __obf_4680696b4dfeb451, __obf_3cbd1ac0144c98a4, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_62b38bc1ef31c0de int64) int64 {
	__obf_9152cf6c829268b7 := time.Unix(__obf_62b38bc1ef31c0de, 0)
	return __obf_9152cf6c829268b7.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_285c0277db982a36, __obf_4680696b4dfeb451, __obf_3cbd1ac0144c98a4 := time.Now().Date()
	return time.Date(__obf_285c0277db982a36, __obf_4680696b4dfeb451, __obf_3cbd1ac0144c98a4, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_1b56f2857581b365 string) (int64, int64) {
	__obf_3d8f132c9b95b085, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_1b56f2857581b365+" 00:00:00", Loc)
	min := __obf_3d8f132c9b95b085.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_24f6061d6b123806 int) (int64, int64) {
	__obf_1b56f2857581b365 := time.Now()
	__obf_9152cf6c829268b7 := __obf_1b56f2857581b365.AddDate(0, 0, __obf_24f6061d6b123806)
	__obf_285c0277db982a36, __obf_4680696b4dfeb451, __obf_3cbd1ac0144c98a4 := __obf_9152cf6c829268b7.Date()
	min := time.Date(__obf_285c0277db982a36, __obf_4680696b4dfeb451, __obf_3cbd1ac0144c98a4, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_1b56f2857581b365.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_b08329a85edaf58a, __obf_11ca1896a3d6e7a0 string) bool {
	__obf_3a7b30fbd6a98558 := func(__obf_5331eba287d21166 string) string {
		return fmt.Sprintf(",%s,", __obf_5331eba287d21166)
	}
	return strings.Contains(__obf_3a7b30fbd6a98558(__obf_b08329a85edaf58a), __obf_3a7b30fbd6a98558(__obf_11ca1896a3d6e7a0))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_515594e624ce23f9 string, __obf_6e1b88907deb29ac ...string) bool {
	if any {
		for _, __obf_3ffba1b5be1ebf4c := range __obf_6e1b88907deb29ac {
			if Contain(__obf_515594e624ce23f9, __obf_3ffba1b5be1ebf4c) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_3ffba1b5be1ebf4c := range __obf_6e1b88907deb29ac {
			if !Contain(__obf_515594e624ce23f9, __obf_3ffba1b5be1ebf4c) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_5331eba287d21166 []any, __obf_374ed272b06bc80f int) []any {
	return slices.Delete(__obf_5331eba287d21166, __obf_374ed272b06bc80f, __obf_374ed272b06bc80f+1)
}

func String2Int8(__obf_f5886586954a86fc string) int8 {
	__obf_4b6befa82615345a, __obf_e812cfd1203ed4f3 := strconv.ParseInt(__obf_f5886586954a86fc, 10, 8)
	if __obf_e812cfd1203ed4f3 == nil {
		return int8(__obf_4b6befa82615345a)
	}
	return 0
}

func String2Int32(__obf_f5886586954a86fc string) int32 {
	__obf_4b6befa82615345a, __obf_e812cfd1203ed4f3 := strconv.ParseInt(__obf_f5886586954a86fc, 10, 32)
	if __obf_e812cfd1203ed4f3 == nil {
		return int32(__obf_4b6befa82615345a)
	}
	return 0
}

func String2Int64(__obf_f5886586954a86fc string) int8 {
	__obf_4b6befa82615345a, __obf_e812cfd1203ed4f3 := strconv.ParseInt(__obf_f5886586954a86fc, 10, 64)
	if __obf_e812cfd1203ed4f3 == nil {
		return int8(__obf_4b6befa82615345a)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_925f09fcc373a370, __obf_c5b17201c45cf553 int) (__obf_ace71eb4992cd39a, __obf_3e837396e16bfe91 time.Time) {
	if __obf_925f09fcc373a370 > 0 && __obf_c5b17201c45cf553 > 0 {
		__obf_ace71eb4992cd39a = time.Date(__obf_925f09fcc373a370, 1, 0, 0, 0, 0, 0, Loc)
		__obf_1a514bbac0ebaa8e := // 第一天是周几
			int(__obf_ace71eb4992cd39a.AddDate(0, 0, 1).Weekday())
		__obf_a56c9a86720c1549 := // 当年第一周有几天
			1
		if __obf_1a514bbac0ebaa8e != 0 {
			__obf_a56c9a86720c1549 = 7 - __obf_1a514bbac0ebaa8e + 1
		}
		if __obf_c5b17201c45cf553 == 1 {
			__obf_3e837396e16bfe91 = __obf_ace71eb4992cd39a.AddDate(0, 0, __obf_a56c9a86720c1549)
		} else {
			__obf_3e837396e16bfe91 = __obf_ace71eb4992cd39a.AddDate(0, 0, __obf_a56c9a86720c1549+(__obf_c5b17201c45cf553-1)*7)
			__obf_ace71eb4992cd39a = __obf_3e837396e16bfe91.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_1b56f2857581b365 time.Time) (__obf_925f09fcc373a370, __obf_c5b17201c45cf553, __obf_1aae24e96f6336af int) {
	__obf_925f09fcc373a370 = __obf_1b56f2857581b365.Year()
	__obf_1aae24e96f6336af = int(__obf_1b56f2857581b365.Weekday())
	__obf_b50d1f63571c1b57 := __obf_1b56f2857581b365.YearDay()
	__obf_65f87c16f84fe99c := __obf_1b56f2857581b365.AddDate(0, 0, -__obf_b50d1f63571c1b57+1)
	__obf_1a514bbac0ebaa8e := int(__obf_65f87c16f84fe99c.Weekday())
	__obf_a56c9a86720c1549 := // 当年第一周有几天
		1
	if __obf_1a514bbac0ebaa8e != 0 {
		__obf_a56c9a86720c1549 = 7 - __obf_1a514bbac0ebaa8e + 1
	}
	if __obf_b50d1f63571c1b57 <= __obf_a56c9a86720c1549 {
		__obf_c5b17201c45cf553 = 1
	} else {
		__obf_c5b17201c45cf553 = (__obf_b50d1f63571c1b57-__obf_a56c9a86720c1549)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_8a3ca2bc8c37dd4f []byte) map[string]string {
	__obf_a5b1fab1d315e928 := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_df2ed56aff87b185 := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_dba2e213085c8291 []byte
	if __obf_dba2e213085c8291 = __obf_a5b1fab1d315e928.ReplaceAll(__obf_8a3ca2bc8c37dd4f, []byte("")); len(__obf_dba2e213085c8291) < 6 {
		return nil
	}
	__obf_7ea83908ffc18d2b := __obf_df2ed56aff87b185.FindAllSubmatch(__obf_dba2e213085c8291[6:len(__obf_dba2e213085c8291)-7], -1)
	__obf_f30e36c0445c2b95 := map[string]string{}
	if __obf_cfe8d8e3f2d7a0a7 := __obf_a5b1fab1d315e928.FindSubmatch(__obf_8a3ca2bc8c37dd4f)[1]; len(__obf_cfe8d8e3f2d7a0a7) != 0 {
		__obf_3fae786f9530b9b4 := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_cfe8d8e3f2d7a0a7, -1)
		for _, __obf_b1dd41bcd3423053 := range __obf_3fae786f9530b9b4 {
			__obf_f30e36c0445c2b95[string(__obf_b1dd41bcd3423053[1])] = string(__obf_b1dd41bcd3423053[2])
		}
	}

	for _, __obf_b1dd41bcd3423053 := range __obf_7ea83908ffc18d2b {
		__obf_f30e36c0445c2b95[string(__obf_b1dd41bcd3423053[1])] = string(__obf_b1dd41bcd3423053[2])
	}
	return __obf_f30e36c0445c2b95
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_c160b716b26035e0 time.Time) (int64, int64) {
	__obf_b0260f32620f0b26 := GetZeroTime(__obf_c160b716b26035e0).Unix() / 600
	return __obf_b0260f32620f0b26, __obf_b0260f32620f0b26 + 143
}

func Abs(__obf_5b1c1c5422ae12c5 int64) int64 {
	__obf_285c0277db982a36 := __obf_5b1c1c5422ae12c5 >> 63
	return (__obf_5b1c1c5422ae12c5 ^ __obf_285c0277db982a36) - __obf_285c0277db982a36
}

func Number2String(__obf_5b1c1c5422ae12c5 any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_5b1c1c5422ae12c5 := __obf_5b1c1c5422ae12c5.(type) {
	case int:
		return strconv.Itoa(__obf_5b1c1c5422ae12c5)
	case int32:
		return strconv.FormatInt(int64(__obf_5b1c1c5422ae12c5), 10)
	case int64:
		return strconv.FormatInt(__obf_5b1c1c5422ae12c5, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_5b1c1c5422ae12c5), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_5b1c1c5422ae12c5, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_f5886586954a86fc any, __obf_5b24539688807f6c string) string {
	if __obf_f5886586954a86fc != nil && __obf_f5886586954a86fc.(string) != "" {
		// return sep + str + sep
		return __obf_f5886586954a86fc.(string)
	}
	return __obf_5b24539688807f6c
}

func SortRange(__obf_4680696b4dfeb451 map[string]any, __obf_fc0b0dc970c2bda4 func(int, string)) {
	var __obf_91a4205e41193433 []string
	for __obf_5028ac00fc46e50f := range __obf_4680696b4dfeb451 {
		__obf_91a4205e41193433 = append(__obf_91a4205e41193433, __obf_5028ac00fc46e50f)
	}
	sort.Strings(__obf_91a4205e41193433)
	for __obf_962fdbe99abe4185, __obf_5028ac00fc46e50f := range __obf_91a4205e41193433 {
		__obf_fc0b0dc970c2bda4(__obf_962fdbe99abe4185, __obf_5028ac00fc46e50f)
	}
}

func HasField(__obf_312f7e38daac66bc reflect.Value, __obf_04d9a2c744f64089 string) bool {

	if __obf_5331eba287d21166 := __obf_312f7e38daac66bc.FieldByNameFunc(func(__obf_5b1c1c5422ae12c5 string) bool {
		return strings.EqualFold(__obf_5b1c1c5422ae12c5, __obf_04d9a2c744f64089)
	}); __obf_5331eba287d21166.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_312f7e38daac66bc reflect.Value, __obf_04d9a2c744f64089 string) reflect.Value {
	if __obf_5331eba287d21166 := __obf_312f7e38daac66bc.FieldByNameFunc(func(__obf_5b1c1c5422ae12c5 string) bool {
		return strings.EqualFold(__obf_5b1c1c5422ae12c5, __obf_04d9a2c744f64089)
	}); __obf_5331eba287d21166.IsValid() {
		return __obf_5331eba287d21166
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_312f7e38daac66bc reflect.Value, __obf_04d9a2c744f64089 string, __obf_9050e2c8d15b9a49 any) bool {

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
	if __obf_5331eba287d21166 := __obf_312f7e38daac66bc.FieldByNameFunc(func(__obf_5b1c1c5422ae12c5 string) bool {
		return strings.EqualFold(__obf_5b1c1c5422ae12c5, __obf_04d9a2c744f64089)
	}); __obf_5331eba287d21166.IsValid() {
		if __obf_a5d8a62eba5f7894 := __obf_5331eba287d21166.Type(); __obf_a5d8a62eba5f7894 == reflect.TypeOf(__obf_9050e2c8d15b9a49) {
			__obf_5331eba287d21166.
				Set(reflect.ValueOf(__obf_9050e2c8d15b9a49))
		} else {
			__obf_5331eba287d21166.
				Set(reflect.ValueOf(__obf_9050e2c8d15b9a49).Convert(__obf_a5d8a62eba5f7894))
		}
		return true
	}
	return false
}

func CopyMap(__obf_4680696b4dfeb451 map[string]any) map[string]any {
	__obf_3fae786f9530b9b4 := make(map[string]any)
	for __obf_5028ac00fc46e50f, __obf_b1dd41bcd3423053 := range __obf_4680696b4dfeb451 {
		if __obf_2795ab03f162bbad, __obf_91e9ef828bbf26ae := __obf_b1dd41bcd3423053.(map[string]any); __obf_91e9ef828bbf26ae {
			__obf_3fae786f9530b9b4[__obf_5028ac00fc46e50f] = CopyMap(__obf_2795ab03f162bbad)
		} else {
			__obf_3fae786f9530b9b4[__obf_5028ac00fc46e50f] = __obf_b1dd41bcd3423053
		}
	}

	return __obf_3fae786f9530b9b4
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_10d6974a5cfc5e1f, __obf_75a5f35ff86574f6 string, __obf_4be41d142d507141 bool, __obf_15b948c206c47ea0 ...any) {
	var __obf_d1c950bb70005551 strings.Builder
	var __obf_d8208a6fbc77ea40 string
	var __obf_ab61c9da0e11f8ab []string
	for _, __obf_d64da226cc1e7d3a := range __obf_15b948c206c47ea0 {
		__obf_5331eba287d21166 := reflect.TypeOf(__obf_d64da226cc1e7d3a)
		__obf_d1c950bb70005551.
			WriteString(`CREATE TABLE `)
		__obf_d1c950bb70005551.
			WriteString(__obf_5331eba287d21166.Name())
		__obf_d1c950bb70005551.
			WriteString(" (\n")
		__obf_dc26918a1f990b05 := __obf_5331eba287d21166.NumField()
		for __obf_962fdbe99abe4185 := 0; __obf_962fdbe99abe4185 < __obf_dc26918a1f990b05; __obf_962fdbe99abe4185++ {
			__obf_d1c950bb70005551.
				WriteString("    ")
			__obf_ab61c9da0e11f8ab = nil
			if __obf_62124934e94086b6 := string(__obf_5331eba287d21166.Field(__obf_962fdbe99abe4185).Tag.Get(__obf_10d6974a5cfc5e1f)); __obf_62124934e94086b6 == "" {
				if __obf_4be41d142d507141 {
					__obf_d8208a6fbc77ea40 = ToCamel(__obf_5331eba287d21166.Field(__obf_962fdbe99abe4185).Name)
				} else {
					__obf_d8208a6fbc77ea40 = __obf_5331eba287d21166.Field(__obf_962fdbe99abe4185).Name
				}
			} else {
				__obf_ab61c9da0e11f8ab = strings.Split(__obf_62124934e94086b6, __obf_75a5f35ff86574f6)
				__obf_d8208a6fbc77ea40 = __obf_ab61c9da0e11f8ab[0]
			}
			__obf_d1c950bb70005551.
				WriteString(__obf_d8208a6fbc77ea40)
			__obf_d1c950bb70005551.
				WriteString(" ")
			switch __obf_5331eba287d21166.Field(__obf_962fdbe99abe4185).Type.Name() {
			case "int8":
				__obf_d1c950bb70005551.
					WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_d1c950bb70005551.
					WriteString("INT")
			case "int64":
				__obf_d1c950bb70005551.
					WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_d1c950bb70005551.
					WriteString("VARCHAR(50)")
			}

			if len(__obf_ab61c9da0e11f8ab) > 1 {
				__obf_d1c950bb70005551.
					WriteString(" ")
				__obf_d1c950bb70005551.
					WriteString(strings.Join(__obf_ab61c9da0e11f8ab[1:], " "))
			}

			if __obf_962fdbe99abe4185+1 != __obf_dc26918a1f990b05 {
				__obf_d1c950bb70005551.
					WriteString(",")
			}
			__obf_d1c950bb70005551.
				WriteString("\n")
		}
		__obf_d1c950bb70005551.
			WriteString(");\n\n")
	}
	fmt.Println(__obf_d1c950bb70005551.String())
}

func SaveImage(__obf_c181a9f654116be0 image.Image, __obf_1a34f67aa85ee7fa uint, __obf_1887085ce0ebe227 string) error {
	__obf_a021358f9f12d664, __obf_e812cfd1203ed4f3 := os.Create(__obf_1887085ce0ebe227)
	if __obf_e812cfd1203ed4f3 != nil {
		return __obf_e812cfd1203ed4f3
	}
	defer __obf_a021358f9f12d664.Close()
	return jpeg.Encode(__obf_a021358f9f12d664, resize.Resize(__obf_1a34f67aa85ee7fa, 0, __obf_c181a9f654116be0, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_f5886586954a86fc string) bool {
	return strings.Contains(__obf_f5886586954a86fc, ".") && net.ParseIP(__obf_f5886586954a86fc) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_f5886586954a86fc string) bool {
	return strings.Contains(__obf_f5886586954a86fc, ":") && net.ParseIP(__obf_f5886586954a86fc) != nil
}

func AnyToBytes(__obf_b1dd41bcd3423053 any) ([]byte, error) {
	return msgpack.Marshal(__obf_b1dd41bcd3423053)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_9050e2c8d15b9a49 []byte) (__obf_2801af39b525757c T, __obf_e812cfd1203ed4f3 error) {
	__obf_e812cfd1203ed4f3 = msgpack.Unmarshal(__obf_9050e2c8d15b9a49, &__obf_2801af39b525757c)
	return
}

func Loadyaml(__obf_9ede601aecfc34a8 string, __obf_4c7f9c9625b44b6a any) {
	__obf_11ca1896a3d6e7a0, __obf_e812cfd1203ed4f3 := os.ReadFile(__obf_9ede601aecfc34a8)
	if __obf_e812cfd1203ed4f3 != nil {
		log.Fatalln(__obf_e812cfd1203ed4f3)
	}
	__obf_e812cfd1203ed4f3 = yaml.UnmarshalStrict(__obf_11ca1896a3d6e7a0, __obf_4c7f9c9625b44b6a)
	if __obf_e812cfd1203ed4f3 != nil {
		log.Fatalln(__obf_e812cfd1203ed4f3)
	}
}

func ToAnyList[T any](__obf_f86f24cda5f73d84 []T) []any {
	__obf_bcb8eebd0c2b98f1 := make([]any, len(__obf_f86f24cda5f73d84))
	for __obf_962fdbe99abe4185, __obf_b1dd41bcd3423053 := range __obf_f86f24cda5f73d84 {
		__obf_bcb8eebd0c2b98f1[__obf_962fdbe99abe4185] = __obf_b1dd41bcd3423053
	}
	return __obf_bcb8eebd0c2b98f1
}

func TimeParse(__obf_50362072120dbd92 string) time.Time {
	__obf_50362072120dbd92 = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_50362072120dbd92)
	__obf_9152cf6c829268b7, _ := time.ParseInLocation("2006-01-02 15:04", __obf_50362072120dbd92, time.Local)
	return __obf_9152cf6c829268b7
}
