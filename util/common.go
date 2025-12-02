package __obf_8b17832dd38bb602

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
	__obf_1688e890467fd727 = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_cab20f4d57b74374 any) ([]byte, error) {
	return __obf_1688e890467fd727.Marshal(__obf_cab20f4d57b74374)
}

func EncodeString(__obf_cab20f4d57b74374 any) string {
	if __obf_5266fec8afeae400, __obf_17881c470b1e6626 := __obf_1688e890467fd727.Marshal(__obf_cab20f4d57b74374); __obf_17881c470b1e6626 == nil {
		return string(__obf_5266fec8afeae400)
	}
	return ""
}

func Decode(__obf_33c594c934f4dc48 string, __obf_cab20f4d57b74374 any) error {
	return __obf_1688e890467fd727.UnmarshalFromString(__obf_33c594c934f4dc48, __obf_cab20f4d57b74374)
}

func DecodeByte(__obf_2304d1cf2f26db23 []byte, __obf_cab20f4d57b74374 any) error {
	return __obf_1688e890467fd727.Unmarshal(__obf_2304d1cf2f26db23, __obf_cab20f4d57b74374)
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

func PrefixInArray(__obf_63adfffcb1a3e166 string, __obf_d1dbe2c327eee3b2 []string) bool {
	for __obf_742320e62d970dab := range __obf_d1dbe2c327eee3b2 {
		if strings.HasPrefix(__obf_63adfffcb1a3e166, __obf_d1dbe2c327eee3b2[__obf_742320e62d970dab]) {
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

var __obf_c61d51864ab40015 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenRandom 获得随机字符串
func GenRandom(__obf_786bd10169423f93 int) string {
	__obf_2304d1cf2f26db23 := make([]byte, __obf_786bd10169423f93)
	for __obf_742320e62d970dab := range __obf_2304d1cf2f26db23 {
		__obf_2304d1cf2f26db23[__obf_742320e62d970dab] = __obf_c61d51864ab40015[rand.IntN(len(__obf_c61d51864ab40015))]
	}
	return string(__obf_2304d1cf2f26db23)
}

func RemoveIndex(__obf_d3f34fed89aa60a8 []string, __obf_923a5818dc332353 int) []string {
	return append(__obf_d3f34fed89aa60a8[:__obf_923a5818dc332353], __obf_d3f34fed89aa60a8[__obf_923a5818dc332353+1:]...)
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
func ToCamel(__obf_33c594c934f4dc48 string) (__obf_5d1c6c1ddd8f37a7 string) {
	__obf_f22d80c3ac83f3b3 := []rune(__obf_33c594c934f4dc48)
	__obf_5d1c6c1ddd8f37a7 = __obf_33c594c934f4dc48[0:1]
	if __obf_f22d80c3ac83f3b3[0] >= 97 && __obf_f22d80c3ac83f3b3[0] <= 122 {
		__obf_5d1c6c1ddd8f37a7 = string(__obf_f22d80c3ac83f3b3[0] - 32)
	}
	__obf_53d798bc9e5fd952 := len(__obf_f22d80c3ac83f3b3)
	for __obf_742320e62d970dab := 1; __obf_742320e62d970dab < __obf_53d798bc9e5fd952; __obf_742320e62d970dab++ {
		if __obf_f22d80c3ac83f3b3[__obf_742320e62d970dab] == 95 && __obf_f22d80c3ac83f3b3[__obf_742320e62d970dab+1] >= 97 && __obf_f22d80c3ac83f3b3[__obf_742320e62d970dab+1] <= 122 {
			__obf_f22d80c3ac83f3b3[ //过滤下划线
			__obf_742320e62d970dab+1] -= 32
		} else {
			__obf_5d1c6c1ddd8f37a7 += string(__obf_f22d80c3ac83f3b3[__obf_742320e62d970dab])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_33c594c934f4dc48 string) (__obf_5d1c6c1ddd8f37a7 string) {
	__obf_f22d80c3ac83f3b3 := []rune(__obf_33c594c934f4dc48)
	__obf_5d1c6c1ddd8f37a7 = __obf_33c594c934f4dc48[0:1]
	if __obf_f22d80c3ac83f3b3[0] >= 65 && __obf_f22d80c3ac83f3b3[0] <= 90 {
		__obf_5d1c6c1ddd8f37a7 = string(__obf_f22d80c3ac83f3b3[0] + 32)
	}
	__obf_786bd10169423f93 := len(__obf_f22d80c3ac83f3b3)
	for __obf_742320e62d970dab := 1; __obf_742320e62d970dab < __obf_786bd10169423f93; __obf_742320e62d970dab++ {
		if __obf_f22d80c3ac83f3b3[__obf_742320e62d970dab] >= 65 && __obf_f22d80c3ac83f3b3[__obf_742320e62d970dab] <= 90 {
			__obf_f22d80c3ac83f3b3[ //大写变小写
			__obf_742320e62d970dab] += 32
			__obf_5d1c6c1ddd8f37a7 += "_"
		}
		__obf_5d1c6c1ddd8f37a7 += string(__obf_f22d80c3ac83f3b3[__obf_742320e62d970dab])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_33c594c934f4dc48 string, __obf_49c7a422dfb3dbf1 int, __obf_786bd10169423f93 int) string {
	__obf_c36ab669f4057c25 := // 将字符串转换为rune切片，以正确处理多字节字符
		[]rune(__obf_33c594c934f4dc48)
	__obf_ad4bb67c6362e4c5 := len(__obf_c36ab669f4057c25)

	// 处理n为负数或0的无效情况
	if __obf_786bd10169423f93 <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_49c7a422dfb3dbf1 < 0 {
		__obf_49c7a422dfb3dbf1 = __obf_ad4bb67c6362e4c5 + __obf_49c7a422dfb3dbf1
	}

	// 边界检查：确保start在有效范围内
	if __obf_49c7a422dfb3dbf1 < 0 || __obf_49c7a422dfb3dbf1 >= __obf_ad4bb67c6362e4c5 {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_c36ab669f4057c25[__obf_49c7a422dfb3dbf1:min(__obf_49c7a422dfb3dbf1+__obf_786bd10169423f93, __obf_ad4bb67c6362e4c5)])
}

func ASCII(__obf_274b76b7499b5c52 rune) rune {
	switch {
	case 97 <= __obf_274b76b7499b5c52 && __obf_274b76b7499b5c52 <= 122:
		return __obf_274b76b7499b5c52 - 32
	case 65 <= __obf_274b76b7499b5c52 && __obf_274b76b7499b5c52 <= 90:
		return __obf_274b76b7499b5c52 + 32
	default:
		return __obf_274b76b7499b5c52
	}
}

func IndexString(__obf_33c594c934f4dc48 string, __obf_16a3a22302fb358d rune, __obf_edbc03f368ae983e int) string {
	__obf_65fad6b129baddb3 := []rune(__obf_33c594c934f4dc48)
	var __obf_1623f06f6e3c3665 bytes.Buffer
	var __obf_786bd10169423f93 int
	for __obf_742320e62d970dab, __obf_017e44f246a7cd42 := 0, len(__obf_65fad6b129baddb3); __obf_742320e62d970dab < __obf_017e44f246a7cd42; __obf_742320e62d970dab++ {
		if __obf_65fad6b129baddb3[__obf_742320e62d970dab] == __obf_16a3a22302fb358d {
			__obf_786bd10169423f93 += 1
		}
		if __obf_786bd10169423f93 == __obf_edbc03f368ae983e {
			break
		}
		__obf_1623f06f6e3c3665.
			WriteRune(__obf_65fad6b129baddb3[__obf_742320e62d970dab])
	}
	return __obf_1623f06f6e3c3665.String()
}

func LastIndexString(__obf_6e113e0dc133cb4d, __obf_477600825a2b06ef string) string {
	__obf_d3f34fed89aa60a8 := strings.Split(__obf_6e113e0dc133cb4d, __obf_477600825a2b06ef)
	if __obf_786bd10169423f93 := len(__obf_d3f34fed89aa60a8); __obf_786bd10169423f93 > 1 {
		return __obf_d3f34fed89aa60a8[__obf_786bd10169423f93-2]
	}
	return ""
}

func IsEmpty(__obf_15c44fbcb4636cae any) bool {
	__obf_d4f65aef7baa5a20 := reflect.ValueOf(__obf_15c44fbcb4636cae)
	if __obf_d4f65aef7baa5a20.Kind() == reflect.Ptr {
		__obf_d4f65aef7baa5a20 = __obf_d4f65aef7baa5a20.Elem()
	}
	return __obf_d4f65aef7baa5a20.Interface() == reflect.Zero(__obf_d4f65aef7baa5a20.Type()).Interface()
}

func MillisecondToDateString(__obf_c796d891b05bc83e int64) string {
	return time.Unix(__obf_c796d891b05bc83e, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_c796d891b05bc83e int64) string {
	return time.Unix(__obf_c796d891b05bc83e, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_bc4e96568bb84981 string) (__obf_b8bb78e4260e7c51, __obf_70e8f0ba458f9d48 string) {
	__obf_bc4e96568bb84981 = filepath.Base(__obf_bc4e96568bb84981)
	__obf_70e8f0ba458f9d48 = filepath.Ext(__obf_bc4e96568bb84981)
	__obf_b8bb78e4260e7c51 = strings.TrimSuffix(__obf_bc4e96568bb84981, __obf_70e8f0ba458f9d48)
	return __obf_b8bb78e4260e7c51,

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
		__obf_70e8f0ba458f9d48
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(__obf_8a67277ce1bc7098 string) (int64, error) {
	if __obf_8b7eeb066de83267, __obf_17881c470b1e6626 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_8a67277ce1bc7098, Loc); __obf_17881c470b1e6626 != nil {
		return 0, nil
	} else {
		__obf_8b7eeb066de83267 = __obf_8b7eeb066de83267.AddDate(0, 0, -__obf_8b7eeb066de83267.Day()+1)
		return GetZeroTime(__obf_8b7eeb066de83267).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_8a67277ce1bc7098 string) (int64, int64, error) {
	if __obf_8b7eeb066de83267, __obf_17881c470b1e6626 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_8a67277ce1bc7098, Loc); __obf_17881c470b1e6626 != nil {
		return 0, 0, nil
	} else {
		__obf_8b7eeb066de83267 = __obf_8b7eeb066de83267.AddDate(0, 0, -__obf_8b7eeb066de83267.Day()+1)
		return GetZeroTime(__obf_8b7eeb066de83267).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_8b7eeb066de83267).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_82958ae45e49795c time.Time) time.Time {
	return time.Date(__obf_82958ae45e49795c.Year(), __obf_82958ae45e49795c.Month(), __obf_82958ae45e49795c.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_c796d891b05bc83e int64) int64 {
	__obf_f3d928b64f4ec1af := time.Unix(__obf_c796d891b05bc83e, 0)
	__obf_84294d4631798a58, __obf_269bbc279a406c9b, __obf_82958ae45e49795c := __obf_f3d928b64f4ec1af.Date()
	return time.Date(__obf_84294d4631798a58, __obf_269bbc279a406c9b, __obf_82958ae45e49795c, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_c796d891b05bc83e int64) int64 {
	__obf_f3d928b64f4ec1af := time.Unix(__obf_c796d891b05bc83e, 0)
	return __obf_f3d928b64f4ec1af.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_84294d4631798a58, __obf_269bbc279a406c9b, __obf_82958ae45e49795c := time.Now().Date()
	return time.Date(__obf_84294d4631798a58, __obf_269bbc279a406c9b, __obf_82958ae45e49795c, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_8a67277ce1bc7098 string) (int64, int64) {
	__obf_305342afe7b47284, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_8a67277ce1bc7098+" 00:00:00", Loc)
	min := __obf_305342afe7b47284.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_8430ae6f0d40596c int) (int64, int64) {
	__obf_8a67277ce1bc7098 := time.Now()
	__obf_f3d928b64f4ec1af := __obf_8a67277ce1bc7098.AddDate(0, 0, __obf_8430ae6f0d40596c)
	__obf_84294d4631798a58, __obf_269bbc279a406c9b, __obf_82958ae45e49795c := __obf_f3d928b64f4ec1af.Date()
	min := time.Date(__obf_84294d4631798a58, __obf_269bbc279a406c9b, __obf_82958ae45e49795c, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_8a67277ce1bc7098.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_110a418716941920, __obf_6bd4634e68bc5c24 string) bool {
	__obf_2d9eb206a4c77357 := func(__obf_d3f34fed89aa60a8 string) string {
		return fmt.Sprintf(",%s,", __obf_d3f34fed89aa60a8)
	}
	return strings.Contains(__obf_2d9eb206a4c77357(__obf_110a418716941920), __obf_2d9eb206a4c77357(__obf_6bd4634e68bc5c24))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_89fe90fbb3aac87c string, __obf_044f89959e8c758a ...string) bool {
	if any {
		for _, __obf_abf6da56a0f5212f := range __obf_044f89959e8c758a {
			if Contain(__obf_89fe90fbb3aac87c, __obf_abf6da56a0f5212f) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_abf6da56a0f5212f := range __obf_044f89959e8c758a {
			if !Contain(__obf_89fe90fbb3aac87c, __obf_abf6da56a0f5212f) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_d3f34fed89aa60a8 []any, __obf_923a5818dc332353 int) []any {
	return slices.Delete(__obf_d3f34fed89aa60a8, __obf_923a5818dc332353, __obf_923a5818dc332353+1)
}

func String2Int8(__obf_33c594c934f4dc48 string) int8 {
	__obf_fa51c9d8aae1a6ad, __obf_17881c470b1e6626 := strconv.ParseInt(__obf_33c594c934f4dc48, 10, 8)
	if __obf_17881c470b1e6626 == nil {
		return int8(__obf_fa51c9d8aae1a6ad)
	}
	return 0
}

func String2Int32(__obf_33c594c934f4dc48 string) int32 {
	__obf_fa51c9d8aae1a6ad, __obf_17881c470b1e6626 := strconv.ParseInt(__obf_33c594c934f4dc48, 10, 32)
	if __obf_17881c470b1e6626 == nil {
		return int32(__obf_fa51c9d8aae1a6ad)
	}
	return 0
}

func String2Int64(__obf_33c594c934f4dc48 string) int8 {
	__obf_fa51c9d8aae1a6ad, __obf_17881c470b1e6626 := strconv.ParseInt(__obf_33c594c934f4dc48, 10, 64)
	if __obf_17881c470b1e6626 == nil {
		return int8(__obf_fa51c9d8aae1a6ad)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_6b6aa41165a051e0, __obf_93cf3f905ff25a88 int) (__obf_49c7a422dfb3dbf1, __obf_10244c94429d8d4a time.Time) {
	if __obf_6b6aa41165a051e0 > 0 && __obf_93cf3f905ff25a88 > 0 {
		__obf_49c7a422dfb3dbf1 = time.Date(__obf_6b6aa41165a051e0, 1, 0, 0, 0, 0, 0, Loc)
		__obf_eeee8842a90ba08c := // 第一天是周几
			int(__obf_49c7a422dfb3dbf1.AddDate(0, 0, 1).Weekday())
		__obf_862e6e258b9d9e29 := // 当年第一周有几天
			1
		if __obf_eeee8842a90ba08c != 0 {
			__obf_862e6e258b9d9e29 = 7 - __obf_eeee8842a90ba08c + 1
		}
		if __obf_93cf3f905ff25a88 == 1 {
			__obf_10244c94429d8d4a = __obf_49c7a422dfb3dbf1.AddDate(0, 0, __obf_862e6e258b9d9e29)
		} else {
			__obf_10244c94429d8d4a = __obf_49c7a422dfb3dbf1.AddDate(0, 0, __obf_862e6e258b9d9e29+(__obf_93cf3f905ff25a88-1)*7)
			__obf_49c7a422dfb3dbf1 = __obf_10244c94429d8d4a.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_8a67277ce1bc7098 time.Time) (__obf_6b6aa41165a051e0, __obf_93cf3f905ff25a88, __obf_187df9ed7fca0b91 int) {
	__obf_6b6aa41165a051e0 = __obf_8a67277ce1bc7098.Year()
	__obf_187df9ed7fca0b91 = int(__obf_8a67277ce1bc7098.Weekday())
	__obf_8a3309aff899d7e5 := __obf_8a67277ce1bc7098.YearDay()
	__obf_62a547b4f8552b52 := __obf_8a67277ce1bc7098.AddDate(0, 0, -__obf_8a3309aff899d7e5+1)
	__obf_eeee8842a90ba08c := int(__obf_62a547b4f8552b52.Weekday())
	__obf_862e6e258b9d9e29 := // 当年第一周有几天
		1
	if __obf_eeee8842a90ba08c != 0 {
		__obf_862e6e258b9d9e29 = 7 - __obf_eeee8842a90ba08c + 1
	}
	if __obf_8a3309aff899d7e5 <= __obf_862e6e258b9d9e29 {
		__obf_93cf3f905ff25a88 = 1
	} else {
		__obf_93cf3f905ff25a88 = (__obf_8a3309aff899d7e5-__obf_862e6e258b9d9e29)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_d712b0b707068372 []byte) map[string]string {
	__obf_536318b158f35ab2 := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_ab84113f6c00cb9f := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_6f314c1b215a5f24 []byte
	if __obf_6f314c1b215a5f24 = __obf_536318b158f35ab2.ReplaceAll(__obf_d712b0b707068372, []byte("")); len(__obf_6f314c1b215a5f24) < 6 {
		return nil
	}
	__obf_23296f9e58f5f8c7 := __obf_ab84113f6c00cb9f.FindAllSubmatch(__obf_6f314c1b215a5f24[6:len(__obf_6f314c1b215a5f24)-7], -1)
	__obf_575d8c52403ffeeb := map[string]string{}
	if __obf_b0ef85feaed7a04a := __obf_536318b158f35ab2.FindSubmatch(__obf_d712b0b707068372)[1]; len(__obf_b0ef85feaed7a04a) != 0 {
		__obf_1acfd3a1503ae678 := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_b0ef85feaed7a04a, -1)
		for _, __obf_d4f65aef7baa5a20 := range __obf_1acfd3a1503ae678 {
			__obf_575d8c52403ffeeb[string(__obf_d4f65aef7baa5a20[1])] = string(__obf_d4f65aef7baa5a20[2])
		}
	}

	for _, __obf_d4f65aef7baa5a20 := range __obf_23296f9e58f5f8c7 {
		__obf_575d8c52403ffeeb[string(__obf_d4f65aef7baa5a20[1])] = string(__obf_d4f65aef7baa5a20[2])
	}
	return __obf_575d8c52403ffeeb
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_f92a87eeeef0ec48 time.Time) (int64, int64) {
	__obf_229a2ee0a6ebbef1 := GetZeroTime(__obf_f92a87eeeef0ec48).Unix() / 600
	return __obf_229a2ee0a6ebbef1, __obf_229a2ee0a6ebbef1 + 143
}

func Abs(__obf_786bd10169423f93 int64) int64 {
	__obf_84294d4631798a58 := __obf_786bd10169423f93 >> 63
	return (__obf_786bd10169423f93 ^ __obf_84294d4631798a58) - __obf_84294d4631798a58
}

func Number2String(__obf_786bd10169423f93 any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_786bd10169423f93 := __obf_786bd10169423f93.(type) {
	case int:
		return strconv.Itoa(__obf_786bd10169423f93)
	case int32:
		return strconv.FormatInt(int64(__obf_786bd10169423f93), 10)
	case int64:
		return strconv.FormatInt(__obf_786bd10169423f93, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_786bd10169423f93), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_786bd10169423f93, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_33c594c934f4dc48 any, __obf_16a3a22302fb358d string) string {
	if __obf_33c594c934f4dc48 != nil && __obf_33c594c934f4dc48.(string) != "" {
		// return sep + str + sep
		return __obf_33c594c934f4dc48.(string)
	}
	return __obf_16a3a22302fb358d
}

func SortRange(__obf_269bbc279a406c9b map[string]any, __obf_6754aaa26fa71907 func(int, string)) {
	var __obf_16c0c97fd1384c05 []string
	for __obf_e97fdc8b2a44cc99 := range __obf_269bbc279a406c9b {
		__obf_16c0c97fd1384c05 = append(__obf_16c0c97fd1384c05, __obf_e97fdc8b2a44cc99)
	}
	sort.Strings(__obf_16c0c97fd1384c05)
	for __obf_742320e62d970dab, __obf_e97fdc8b2a44cc99 := range __obf_16c0c97fd1384c05 {
		__obf_6754aaa26fa71907(__obf_742320e62d970dab, __obf_e97fdc8b2a44cc99)
	}
}

func HasField(__obf_3a8f841809acb8f1 reflect.Value, __obf_b8bb78e4260e7c51 string) bool {

	if __obf_d3f34fed89aa60a8 := __obf_3a8f841809acb8f1.FieldByNameFunc(func(__obf_786bd10169423f93 string) bool {
		return strings.EqualFold(__obf_786bd10169423f93, __obf_b8bb78e4260e7c51)
	}); __obf_d3f34fed89aa60a8.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_3a8f841809acb8f1 reflect.Value, __obf_b8bb78e4260e7c51 string) reflect.Value {
	if __obf_d3f34fed89aa60a8 := __obf_3a8f841809acb8f1.FieldByNameFunc(func(__obf_786bd10169423f93 string) bool {
		return strings.EqualFold(__obf_786bd10169423f93, __obf_b8bb78e4260e7c51)
	}); __obf_d3f34fed89aa60a8.IsValid() {
		return __obf_d3f34fed89aa60a8
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_3a8f841809acb8f1 reflect.Value, __obf_b8bb78e4260e7c51 string, __obf_63adfffcb1a3e166 any) bool {

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
	if __obf_d3f34fed89aa60a8 := __obf_3a8f841809acb8f1.FieldByNameFunc(func(__obf_786bd10169423f93 string) bool {
		return strings.EqualFold(__obf_786bd10169423f93, __obf_b8bb78e4260e7c51)
	}); __obf_d3f34fed89aa60a8.IsValid() {
		if __obf_47907f86be3c83af := __obf_d3f34fed89aa60a8.Type(); __obf_47907f86be3c83af == reflect.TypeOf(__obf_63adfffcb1a3e166) {
			__obf_d3f34fed89aa60a8.
				Set(reflect.ValueOf(__obf_63adfffcb1a3e166))
		} else {
			__obf_d3f34fed89aa60a8.
				Set(reflect.ValueOf(__obf_63adfffcb1a3e166).Convert(__obf_47907f86be3c83af))
		}
		return true
	}
	return false
}

func CopyMap(__obf_269bbc279a406c9b map[string]any) map[string]any {
	__obf_1acfd3a1503ae678 := make(map[string]any)
	for __obf_e97fdc8b2a44cc99, __obf_d4f65aef7baa5a20 := range __obf_269bbc279a406c9b {
		if __obf_591ed0b3242de369, __obf_3238e128616a1cf5 := __obf_d4f65aef7baa5a20.(map[string]any); __obf_3238e128616a1cf5 {
			__obf_1acfd3a1503ae678[__obf_e97fdc8b2a44cc99] = CopyMap(__obf_591ed0b3242de369)
		} else {
			__obf_1acfd3a1503ae678[__obf_e97fdc8b2a44cc99] = __obf_d4f65aef7baa5a20
		}
	}

	return __obf_1acfd3a1503ae678
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_adf3e0476606c3da, __obf_04573e5dd51225c0 string, __obf_d5ceefcbdd63cca8 bool, __obf_6d18d44bef02b3f6 ...any) {
	var __obf_10d950b29180e5b5 strings.Builder
	var __obf_74d81b907b8c31af string
	var __obf_8ebe306ee2ce4647 []string
	for _, __obf_2b40febcc1e757dd := range __obf_6d18d44bef02b3f6 {
		__obf_d3f34fed89aa60a8 := reflect.TypeOf(__obf_2b40febcc1e757dd)
		__obf_10d950b29180e5b5.
			WriteString(`CREATE TABLE `)
		__obf_10d950b29180e5b5.
			WriteString(__obf_d3f34fed89aa60a8.Name())
		__obf_10d950b29180e5b5.
			WriteString(" (\n")
		__obf_e2a3ce374cbe9bff := __obf_d3f34fed89aa60a8.NumField()
		for __obf_742320e62d970dab := 0; __obf_742320e62d970dab < __obf_e2a3ce374cbe9bff; __obf_742320e62d970dab++ {
			__obf_10d950b29180e5b5.
				WriteString("    ")
			__obf_8ebe306ee2ce4647 = nil
			if __obf_737d6e964a81bc17 := string(__obf_d3f34fed89aa60a8.Field(__obf_742320e62d970dab).Tag.Get(__obf_adf3e0476606c3da)); __obf_737d6e964a81bc17 == "" {
				if __obf_d5ceefcbdd63cca8 {
					__obf_74d81b907b8c31af = ToCamel(__obf_d3f34fed89aa60a8.Field(__obf_742320e62d970dab).Name)
				} else {
					__obf_74d81b907b8c31af = __obf_d3f34fed89aa60a8.Field(__obf_742320e62d970dab).Name
				}
			} else {
				__obf_8ebe306ee2ce4647 = strings.Split(__obf_737d6e964a81bc17, __obf_04573e5dd51225c0)
				__obf_74d81b907b8c31af = __obf_8ebe306ee2ce4647[0]
			}
			__obf_10d950b29180e5b5.
				WriteString(__obf_74d81b907b8c31af)
			__obf_10d950b29180e5b5.
				WriteString(" ")
			switch __obf_d3f34fed89aa60a8.Field(__obf_742320e62d970dab).Type.Name() {
			case "int8":
				__obf_10d950b29180e5b5.
					WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_10d950b29180e5b5.
					WriteString("INT")
			case "int64":
				__obf_10d950b29180e5b5.
					WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_10d950b29180e5b5.
					WriteString("VARCHAR(50)")
			}

			if len(__obf_8ebe306ee2ce4647) > 1 {
				__obf_10d950b29180e5b5.
					WriteString(" ")
				__obf_10d950b29180e5b5.
					WriteString(strings.Join(__obf_8ebe306ee2ce4647[1:], " "))
			}

			if __obf_742320e62d970dab+1 != __obf_e2a3ce374cbe9bff {
				__obf_10d950b29180e5b5.
					WriteString(",")
			}
			__obf_10d950b29180e5b5.
				WriteString("\n")
		}
		__obf_10d950b29180e5b5.
			WriteString(");\n\n")
	}
	fmt.Println(__obf_10d950b29180e5b5.String())
}

func SaveImage(__obf_6c1a907cdf477407 image.Image, __obf_58e0c1dce698f81d uint, __obf_6f15f312caaae9ef string) error {
	__obf_32025fa9c4461895, __obf_17881c470b1e6626 := os.Create(__obf_6f15f312caaae9ef)
	if __obf_17881c470b1e6626 != nil {
		return __obf_17881c470b1e6626
	}
	defer __obf_32025fa9c4461895.Close()
	return jpeg.Encode(__obf_32025fa9c4461895, resize.Resize(__obf_58e0c1dce698f81d, 0, __obf_6c1a907cdf477407, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_33c594c934f4dc48 string) bool {
	return strings.Contains(__obf_33c594c934f4dc48, ".") && net.ParseIP(__obf_33c594c934f4dc48) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_33c594c934f4dc48 string) bool {
	return strings.Contains(__obf_33c594c934f4dc48, ":") && net.ParseIP(__obf_33c594c934f4dc48) != nil
}

func AnyToBytes(__obf_d4f65aef7baa5a20 any) ([]byte, error) {
	return msgpack.Marshal(__obf_d4f65aef7baa5a20)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_63adfffcb1a3e166 []byte) (__obf_cab20f4d57b74374 T, __obf_17881c470b1e6626 error) {
	__obf_17881c470b1e6626 = msgpack.Unmarshal(__obf_63adfffcb1a3e166, &__obf_cab20f4d57b74374)
	return
}

func Loadyaml(__obf_38f497afd9ab82ad string, __obf_99d1b7ba4ac13a9a any) {
	__obf_6bd4634e68bc5c24, __obf_17881c470b1e6626 := os.ReadFile(__obf_38f497afd9ab82ad)
	if __obf_17881c470b1e6626 != nil {
		log.Fatalln(__obf_17881c470b1e6626)
	}
	__obf_17881c470b1e6626 = yaml.UnmarshalStrict(__obf_6bd4634e68bc5c24, __obf_99d1b7ba4ac13a9a)
	if __obf_17881c470b1e6626 != nil {
		log.Fatalln(__obf_17881c470b1e6626)
	}
}

func ToAnyList[T any](__obf_dcd6cb876460107b []T) []any {
	__obf_cf82a447129e2a81 := make([]any, len(__obf_dcd6cb876460107b))
	for __obf_742320e62d970dab, __obf_d4f65aef7baa5a20 := range __obf_dcd6cb876460107b {
		__obf_cf82a447129e2a81[__obf_742320e62d970dab] = __obf_d4f65aef7baa5a20
	}
	return __obf_cf82a447129e2a81
}

func TimeParse(__obf_85f4e09df3f275e4 string) time.Time {
	__obf_85f4e09df3f275e4 = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_85f4e09df3f275e4)
	__obf_f3d928b64f4ec1af, _ := time.ParseInLocation("2006-01-02 15:04", __obf_85f4e09df3f275e4, time.Local)
	return __obf_f3d928b64f4ec1af
}
