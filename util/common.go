package __obf_426da37e60cac670

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
	__obf_c77efeb3502803e0 = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_96894de1e88392eb any) ([]byte, error) {
	return __obf_c77efeb3502803e0.Marshal(__obf_96894de1e88392eb)
}

func EncodeString(__obf_96894de1e88392eb any) string {
	if __obf_57ffa0fa03aa449d, __obf_74916b80241ef1ff := __obf_c77efeb3502803e0.Marshal(__obf_96894de1e88392eb); __obf_74916b80241ef1ff == nil {
		return string(__obf_57ffa0fa03aa449d)
	}
	return ""
}

func Decode(__obf_a1d01074eaf8ad7a string, __obf_96894de1e88392eb any) error {
	return __obf_c77efeb3502803e0.UnmarshalFromString(__obf_a1d01074eaf8ad7a, __obf_96894de1e88392eb)
}

func DecodeByte(__obf_1871a36bb793083c []byte, __obf_96894de1e88392eb any) error {
	return __obf_c77efeb3502803e0.Unmarshal(__obf_1871a36bb793083c, __obf_96894de1e88392eb)
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

func PrefixInArray(__obf_78314a8a6680c2e3 string, __obf_21a32d21b3b8839f []string) bool {
	for __obf_9f418fe88df9aaef := range __obf_21a32d21b3b8839f {
		if strings.HasPrefix(__obf_78314a8a6680c2e3, __obf_21a32d21b3b8839f[__obf_9f418fe88df9aaef]) {
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

var __obf_847703e4ad3133e1 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenRandom 获得随机字符串
func GenRandom(__obf_f7eefc38acf35467 int) string {
	__obf_1871a36bb793083c := make([]byte, __obf_f7eefc38acf35467)
	for __obf_9f418fe88df9aaef := range __obf_1871a36bb793083c {
		__obf_1871a36bb793083c[__obf_9f418fe88df9aaef] = __obf_847703e4ad3133e1[rand.IntN(len(__obf_847703e4ad3133e1))]
	}
	return string(__obf_1871a36bb793083c)
}

func RemoveIndex(__obf_b69dd8ffbe131543 []string, __obf_6350743dad5dd760 int) []string {
	return append(__obf_b69dd8ffbe131543[:__obf_6350743dad5dd760], __obf_b69dd8ffbe131543[__obf_6350743dad5dd760+1:]...)
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
func ToCamel(__obf_a1d01074eaf8ad7a string) (__obf_9088057b0091059d string) {
	__obf_637aa4ca65b63b8f := []rune(__obf_a1d01074eaf8ad7a)
	__obf_9088057b0091059d = __obf_a1d01074eaf8ad7a[0:1]
	if __obf_637aa4ca65b63b8f[0] >= 97 && __obf_637aa4ca65b63b8f[0] <= 122 {
		__obf_9088057b0091059d = string(__obf_637aa4ca65b63b8f[0] - 32)
	}
	__obf_57ed846508d4cdd2 := len(__obf_637aa4ca65b63b8f)
	for __obf_9f418fe88df9aaef := 1; __obf_9f418fe88df9aaef < __obf_57ed846508d4cdd2; __obf_9f418fe88df9aaef++ {
		if __obf_637aa4ca65b63b8f[__obf_9f418fe88df9aaef] == 95 && __obf_637aa4ca65b63b8f[__obf_9f418fe88df9aaef+1] >= 97 && __obf_637aa4ca65b63b8f[__obf_9f418fe88df9aaef+1] <= 122 {
			__obf_637aa4ca65b63b8f[ //过滤下划线
			__obf_9f418fe88df9aaef+1] -= 32
		} else {
			__obf_9088057b0091059d += string(__obf_637aa4ca65b63b8f[__obf_9f418fe88df9aaef])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_a1d01074eaf8ad7a string) (__obf_9088057b0091059d string) {
	__obf_637aa4ca65b63b8f := []rune(__obf_a1d01074eaf8ad7a)
	__obf_9088057b0091059d = __obf_a1d01074eaf8ad7a[0:1]
	if __obf_637aa4ca65b63b8f[0] >= 65 && __obf_637aa4ca65b63b8f[0] <= 90 {
		__obf_9088057b0091059d = string(__obf_637aa4ca65b63b8f[0] + 32)
	}
	__obf_f7eefc38acf35467 := len(__obf_637aa4ca65b63b8f)
	for __obf_9f418fe88df9aaef := 1; __obf_9f418fe88df9aaef < __obf_f7eefc38acf35467; __obf_9f418fe88df9aaef++ {
		if __obf_637aa4ca65b63b8f[__obf_9f418fe88df9aaef] >= 65 && __obf_637aa4ca65b63b8f[__obf_9f418fe88df9aaef] <= 90 {
			__obf_637aa4ca65b63b8f[ //大写变小写
			__obf_9f418fe88df9aaef] += 32
			__obf_9088057b0091059d += "_"
		}
		__obf_9088057b0091059d += string(__obf_637aa4ca65b63b8f[__obf_9f418fe88df9aaef])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_a1d01074eaf8ad7a string, __obf_45e86ccd431992bf int, __obf_f7eefc38acf35467 int) string {
	__obf_72976309e7b69f1c := // 将字符串转换为rune切片，以正确处理多字节字符
		[]rune(__obf_a1d01074eaf8ad7a)
	__obf_a5eeba35e2de3b47 := len(__obf_72976309e7b69f1c)

	// 处理n为负数或0的无效情况
	if __obf_f7eefc38acf35467 <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_45e86ccd431992bf < 0 {
		__obf_45e86ccd431992bf = __obf_a5eeba35e2de3b47 + __obf_45e86ccd431992bf
	}

	// 边界检查：确保start在有效范围内
	if __obf_45e86ccd431992bf < 0 || __obf_45e86ccd431992bf >= __obf_a5eeba35e2de3b47 {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_72976309e7b69f1c[__obf_45e86ccd431992bf:min(__obf_45e86ccd431992bf+__obf_f7eefc38acf35467, __obf_a5eeba35e2de3b47)])
}

func ASCII(__obf_3be342a2e8ac84dc rune) rune {
	switch {
	case 97 <= __obf_3be342a2e8ac84dc && __obf_3be342a2e8ac84dc <= 122:
		return __obf_3be342a2e8ac84dc - 32
	case 65 <= __obf_3be342a2e8ac84dc && __obf_3be342a2e8ac84dc <= 90:
		return __obf_3be342a2e8ac84dc + 32
	default:
		return __obf_3be342a2e8ac84dc
	}
}

func IndexString(__obf_a1d01074eaf8ad7a string, __obf_f442da269b639645 rune, __obf_41a5e3e6e2373371 int) string {
	__obf_803a89718335a817 := []rune(__obf_a1d01074eaf8ad7a)
	var __obf_bf188cb97948b3b2 bytes.Buffer
	var __obf_f7eefc38acf35467 int
	for __obf_9f418fe88df9aaef, __obf_96869096dab539b3 := 0, len(__obf_803a89718335a817); __obf_9f418fe88df9aaef < __obf_96869096dab539b3; __obf_9f418fe88df9aaef++ {
		if __obf_803a89718335a817[__obf_9f418fe88df9aaef] == __obf_f442da269b639645 {
			__obf_f7eefc38acf35467 += 1
		}
		if __obf_f7eefc38acf35467 == __obf_41a5e3e6e2373371 {
			break
		}
		__obf_bf188cb97948b3b2.
			WriteRune(__obf_803a89718335a817[__obf_9f418fe88df9aaef])
	}
	return __obf_bf188cb97948b3b2.String()
}

func LastIndexString(__obf_18b7ce7a2be7b8cb, __obf_fe0f0c5862f9518f string) string {
	__obf_b69dd8ffbe131543 := strings.Split(__obf_18b7ce7a2be7b8cb, __obf_fe0f0c5862f9518f)
	if __obf_f7eefc38acf35467 := len(__obf_b69dd8ffbe131543); __obf_f7eefc38acf35467 > 1 {
		return __obf_b69dd8ffbe131543[__obf_f7eefc38acf35467-2]
	}
	return ""
}

func IsEmpty(__obf_327df22ecbea62b2 any) bool {
	__obf_c76af479674ee3d0 := reflect.ValueOf(__obf_327df22ecbea62b2)
	if __obf_c76af479674ee3d0.Kind() == reflect.Ptr {
		__obf_c76af479674ee3d0 = __obf_c76af479674ee3d0.Elem()
	}
	return __obf_c76af479674ee3d0.Interface() == reflect.Zero(__obf_c76af479674ee3d0.Type()).Interface()
}

func MillisecondToDateString(__obf_efb05265093293d7 int64) string {
	return time.Unix(__obf_efb05265093293d7, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_efb05265093293d7 int64) string {
	return time.Unix(__obf_efb05265093293d7, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_dcf00bef8e19f10e string) (__obf_25ec75064e81d6e1, __obf_4dea8813aa560f0c string) {
	__obf_dcf00bef8e19f10e = filepath.Base(__obf_dcf00bef8e19f10e)
	__obf_4dea8813aa560f0c = filepath.Ext(__obf_dcf00bef8e19f10e)
	__obf_25ec75064e81d6e1 = strings.TrimSuffix(__obf_dcf00bef8e19f10e, __obf_4dea8813aa560f0c)
	return __obf_25ec75064e81d6e1,

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
		__obf_4dea8813aa560f0c
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(__obf_98f0c646fd6de3f5 string) (int64, error) {
	if __obf_cde6665e254e6154, __obf_74916b80241ef1ff := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_98f0c646fd6de3f5, Loc); __obf_74916b80241ef1ff != nil {
		return 0, nil
	} else {
		__obf_cde6665e254e6154 = __obf_cde6665e254e6154.AddDate(0, 0, -__obf_cde6665e254e6154.Day()+1)
		return GetZeroTime(__obf_cde6665e254e6154).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_98f0c646fd6de3f5 string) (int64, int64, error) {
	if __obf_cde6665e254e6154, __obf_74916b80241ef1ff := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_98f0c646fd6de3f5, Loc); __obf_74916b80241ef1ff != nil {
		return 0, 0, nil
	} else {
		__obf_cde6665e254e6154 = __obf_cde6665e254e6154.AddDate(0, 0, -__obf_cde6665e254e6154.Day()+1)
		return GetZeroTime(__obf_cde6665e254e6154).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_cde6665e254e6154).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_ddc716cb6518a16a time.Time) time.Time {
	return time.Date(__obf_ddc716cb6518a16a.Year(), __obf_ddc716cb6518a16a.Month(), __obf_ddc716cb6518a16a.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_efb05265093293d7 int64) int64 {
	__obf_3c9398873f01bdde := time.Unix(__obf_efb05265093293d7, 0)
	__obf_f7b31fbdd4a98dbe, __obf_9274d0cb16fd12a1, __obf_ddc716cb6518a16a := __obf_3c9398873f01bdde.Date()
	return time.Date(__obf_f7b31fbdd4a98dbe, __obf_9274d0cb16fd12a1, __obf_ddc716cb6518a16a, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_efb05265093293d7 int64) int64 {
	__obf_3c9398873f01bdde := time.Unix(__obf_efb05265093293d7, 0)
	return __obf_3c9398873f01bdde.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_f7b31fbdd4a98dbe, __obf_9274d0cb16fd12a1, __obf_ddc716cb6518a16a := time.Now().Date()
	return time.Date(__obf_f7b31fbdd4a98dbe, __obf_9274d0cb16fd12a1, __obf_ddc716cb6518a16a, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_98f0c646fd6de3f5 string) (int64, int64) {
	__obf_946868c73a452eff, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_98f0c646fd6de3f5+" 00:00:00", Loc)
	min := __obf_946868c73a452eff.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_e07b9e7cc32b5d8d int) (int64, int64) {
	__obf_98f0c646fd6de3f5 := time.Now()
	__obf_3c9398873f01bdde := __obf_98f0c646fd6de3f5.AddDate(0, 0, __obf_e07b9e7cc32b5d8d)
	__obf_f7b31fbdd4a98dbe, __obf_9274d0cb16fd12a1, __obf_ddc716cb6518a16a := __obf_3c9398873f01bdde.Date()
	min := time.Date(__obf_f7b31fbdd4a98dbe, __obf_9274d0cb16fd12a1, __obf_ddc716cb6518a16a, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_98f0c646fd6de3f5.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_7fd1bad11015f313, __obf_4f5a11ae844c76f9 string) bool {
	__obf_2d816b4393c15d4d := func(__obf_b69dd8ffbe131543 string) string {
		return fmt.Sprintf(",%s,", __obf_b69dd8ffbe131543)
	}
	return strings.Contains(__obf_2d816b4393c15d4d(__obf_7fd1bad11015f313), __obf_2d816b4393c15d4d(__obf_4f5a11ae844c76f9))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_2cac35b215db05d3 string, __obf_cbcce9b9437b8990 ...string) bool {
	if any {
		for _, __obf_8f69d04a4f68fa13 := range __obf_cbcce9b9437b8990 {
			if Contain(__obf_2cac35b215db05d3, __obf_8f69d04a4f68fa13) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_8f69d04a4f68fa13 := range __obf_cbcce9b9437b8990 {
			if !Contain(__obf_2cac35b215db05d3, __obf_8f69d04a4f68fa13) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_b69dd8ffbe131543 []any, __obf_6350743dad5dd760 int) []any {
	return slices.Delete(__obf_b69dd8ffbe131543, __obf_6350743dad5dd760, __obf_6350743dad5dd760+1)
}

func String2Int8(__obf_a1d01074eaf8ad7a string) int8 {
	__obf_44c91a9f865a4768, __obf_74916b80241ef1ff := strconv.ParseInt(__obf_a1d01074eaf8ad7a, 10, 8)
	if __obf_74916b80241ef1ff == nil {
		return int8(__obf_44c91a9f865a4768)
	}
	return 0
}

func String2Int32(__obf_a1d01074eaf8ad7a string) int32 {
	__obf_44c91a9f865a4768, __obf_74916b80241ef1ff := strconv.ParseInt(__obf_a1d01074eaf8ad7a, 10, 32)
	if __obf_74916b80241ef1ff == nil {
		return int32(__obf_44c91a9f865a4768)
	}
	return 0
}

func String2Int64(__obf_a1d01074eaf8ad7a string) int8 {
	__obf_44c91a9f865a4768, __obf_74916b80241ef1ff := strconv.ParseInt(__obf_a1d01074eaf8ad7a, 10, 64)
	if __obf_74916b80241ef1ff == nil {
		return int8(__obf_44c91a9f865a4768)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_3f917c41dcbf3ae7, __obf_c3124209773b67aa int) (__obf_45e86ccd431992bf, __obf_94b5cd395cab417f time.Time) {
	if __obf_3f917c41dcbf3ae7 > 0 && __obf_c3124209773b67aa > 0 {
		__obf_45e86ccd431992bf = time.Date(__obf_3f917c41dcbf3ae7, 1, 0, 0, 0, 0, 0, Loc)
		__obf_5164a56e71c3db22 := // 第一天是周几
			int(__obf_45e86ccd431992bf.AddDate(0, 0, 1).Weekday())
		__obf_819ba529e0670acc := // 当年第一周有几天
			1
		if __obf_5164a56e71c3db22 != 0 {
			__obf_819ba529e0670acc = 7 - __obf_5164a56e71c3db22 + 1
		}
		if __obf_c3124209773b67aa == 1 {
			__obf_94b5cd395cab417f = __obf_45e86ccd431992bf.AddDate(0, 0, __obf_819ba529e0670acc)
		} else {
			__obf_94b5cd395cab417f = __obf_45e86ccd431992bf.AddDate(0, 0, __obf_819ba529e0670acc+(__obf_c3124209773b67aa-1)*7)
			__obf_45e86ccd431992bf = __obf_94b5cd395cab417f.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_98f0c646fd6de3f5 time.Time) (__obf_3f917c41dcbf3ae7, __obf_c3124209773b67aa, __obf_f95caaad420abe18 int) {
	__obf_3f917c41dcbf3ae7 = __obf_98f0c646fd6de3f5.Year()
	__obf_f95caaad420abe18 = int(__obf_98f0c646fd6de3f5.Weekday())
	__obf_38af294d2489b79b := __obf_98f0c646fd6de3f5.YearDay()
	__obf_ea01e2e791b89151 := __obf_98f0c646fd6de3f5.AddDate(0, 0, -__obf_38af294d2489b79b+1)
	__obf_5164a56e71c3db22 := int(__obf_ea01e2e791b89151.Weekday())
	__obf_819ba529e0670acc := // 当年第一周有几天
		1
	if __obf_5164a56e71c3db22 != 0 {
		__obf_819ba529e0670acc = 7 - __obf_5164a56e71c3db22 + 1
	}
	if __obf_38af294d2489b79b <= __obf_819ba529e0670acc {
		__obf_c3124209773b67aa = 1
	} else {
		__obf_c3124209773b67aa = (__obf_38af294d2489b79b-__obf_819ba529e0670acc)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_db93bd2077589795 []byte) map[string]string {
	__obf_838642ef6b42f30e := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_0ce2aab0e27fa8bd := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_02d160405c076cce []byte
	if __obf_02d160405c076cce = __obf_838642ef6b42f30e.ReplaceAll(__obf_db93bd2077589795, []byte("")); len(__obf_02d160405c076cce) < 6 {
		return nil
	}
	__obf_90cf9b241b498b5d := __obf_0ce2aab0e27fa8bd.FindAllSubmatch(__obf_02d160405c076cce[6:len(__obf_02d160405c076cce)-7], -1)
	__obf_55dfd9101ff64ab1 := map[string]string{}
	if __obf_aff97ba640720ba1 := __obf_838642ef6b42f30e.FindSubmatch(__obf_db93bd2077589795)[1]; len(__obf_aff97ba640720ba1) != 0 {
		__obf_3eb3aeef3b56277e := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_aff97ba640720ba1, -1)
		for _, __obf_c76af479674ee3d0 := range __obf_3eb3aeef3b56277e {
			__obf_55dfd9101ff64ab1[string(__obf_c76af479674ee3d0[1])] = string(__obf_c76af479674ee3d0[2])
		}
	}

	for _, __obf_c76af479674ee3d0 := range __obf_90cf9b241b498b5d {
		__obf_55dfd9101ff64ab1[string(__obf_c76af479674ee3d0[1])] = string(__obf_c76af479674ee3d0[2])
	}
	return __obf_55dfd9101ff64ab1
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_626e07c32cf9bcfc time.Time) (int64, int64) {
	__obf_81ad6a8b349852e4 := GetZeroTime(__obf_626e07c32cf9bcfc).Unix() / 600
	return __obf_81ad6a8b349852e4, __obf_81ad6a8b349852e4 + 143
}

func Abs(__obf_f7eefc38acf35467 int64) int64 {
	__obf_f7b31fbdd4a98dbe := __obf_f7eefc38acf35467 >> 63
	return (__obf_f7eefc38acf35467 ^ __obf_f7b31fbdd4a98dbe) - __obf_f7b31fbdd4a98dbe
}

func Number2String(__obf_f7eefc38acf35467 any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_f7eefc38acf35467 := __obf_f7eefc38acf35467.(type) {
	case int:
		return strconv.Itoa(__obf_f7eefc38acf35467)
	case int32:
		return strconv.FormatInt(int64(__obf_f7eefc38acf35467), 10)
	case int64:
		return strconv.FormatInt(__obf_f7eefc38acf35467, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_f7eefc38acf35467), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_f7eefc38acf35467, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_a1d01074eaf8ad7a any, __obf_f442da269b639645 string) string {
	if __obf_a1d01074eaf8ad7a != nil && __obf_a1d01074eaf8ad7a.(string) != "" {
		// return sep + str + sep
		return __obf_a1d01074eaf8ad7a.(string)
	}
	return __obf_f442da269b639645
}

func SortRange(__obf_9274d0cb16fd12a1 map[string]any, __obf_ea8da91a172a8512 func(int, string)) {
	var __obf_487ebd9caa583a61 []string
	for __obf_35fc606f1e48ca40 := range __obf_9274d0cb16fd12a1 {
		__obf_487ebd9caa583a61 = append(__obf_487ebd9caa583a61, __obf_35fc606f1e48ca40)
	}
	sort.Strings(__obf_487ebd9caa583a61)
	for __obf_9f418fe88df9aaef, __obf_35fc606f1e48ca40 := range __obf_487ebd9caa583a61 {
		__obf_ea8da91a172a8512(__obf_9f418fe88df9aaef, __obf_35fc606f1e48ca40)
	}
}

func HasField(__obf_4092577a5bf1e979 reflect.Value, __obf_25ec75064e81d6e1 string) bool {

	if __obf_b69dd8ffbe131543 := __obf_4092577a5bf1e979.FieldByNameFunc(func(__obf_f7eefc38acf35467 string) bool {
		return strings.EqualFold(__obf_f7eefc38acf35467, __obf_25ec75064e81d6e1)
	}); __obf_b69dd8ffbe131543.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_4092577a5bf1e979 reflect.Value, __obf_25ec75064e81d6e1 string) reflect.Value {
	if __obf_b69dd8ffbe131543 := __obf_4092577a5bf1e979.FieldByNameFunc(func(__obf_f7eefc38acf35467 string) bool {
		return strings.EqualFold(__obf_f7eefc38acf35467, __obf_25ec75064e81d6e1)
	}); __obf_b69dd8ffbe131543.IsValid() {
		return __obf_b69dd8ffbe131543
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_4092577a5bf1e979 reflect.Value, __obf_25ec75064e81d6e1 string, __obf_78314a8a6680c2e3 any) bool {

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
	if __obf_b69dd8ffbe131543 := __obf_4092577a5bf1e979.FieldByNameFunc(func(__obf_f7eefc38acf35467 string) bool {
		return strings.EqualFold(__obf_f7eefc38acf35467, __obf_25ec75064e81d6e1)
	}); __obf_b69dd8ffbe131543.IsValid() {
		if __obf_8913b14b67eccc1e := __obf_b69dd8ffbe131543.Type(); __obf_8913b14b67eccc1e == reflect.TypeOf(__obf_78314a8a6680c2e3) {
			__obf_b69dd8ffbe131543.
				Set(reflect.ValueOf(__obf_78314a8a6680c2e3))
		} else {
			__obf_b69dd8ffbe131543.
				Set(reflect.ValueOf(__obf_78314a8a6680c2e3).Convert(__obf_8913b14b67eccc1e))
		}
		return true
	}
	return false
}

func CopyMap(__obf_9274d0cb16fd12a1 map[string]any) map[string]any {
	__obf_3eb3aeef3b56277e := make(map[string]any)
	for __obf_35fc606f1e48ca40, __obf_c76af479674ee3d0 := range __obf_9274d0cb16fd12a1 {
		if __obf_619e8baf21e4ea36, __obf_f5603d2a4f06faee := __obf_c76af479674ee3d0.(map[string]any); __obf_f5603d2a4f06faee {
			__obf_3eb3aeef3b56277e[__obf_35fc606f1e48ca40] = CopyMap(__obf_619e8baf21e4ea36)
		} else {
			__obf_3eb3aeef3b56277e[__obf_35fc606f1e48ca40] = __obf_c76af479674ee3d0
		}
	}

	return __obf_3eb3aeef3b56277e
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_fe32fb81ebf474d0, __obf_68cfc5f67cac1404 string, __obf_e0d28168269174e4 bool, __obf_c7e3e17fbaf4078a ...any) {
	var __obf_5691d905931ddccd strings.Builder
	var __obf_b3fed84695b938c9 string
	var __obf_df0a7558b9c581fd []string
	for _, __obf_da1a70d8c0dc1912 := range __obf_c7e3e17fbaf4078a {
		__obf_b69dd8ffbe131543 := reflect.TypeOf(__obf_da1a70d8c0dc1912)
		__obf_5691d905931ddccd.
			WriteString(`CREATE TABLE `)
		__obf_5691d905931ddccd.
			WriteString(__obf_b69dd8ffbe131543.Name())
		__obf_5691d905931ddccd.
			WriteString(" (\n")
		__obf_bb71ca1556579102 := __obf_b69dd8ffbe131543.NumField()
		for __obf_9f418fe88df9aaef := 0; __obf_9f418fe88df9aaef < __obf_bb71ca1556579102; __obf_9f418fe88df9aaef++ {
			__obf_5691d905931ddccd.
				WriteString("    ")
			__obf_df0a7558b9c581fd = nil
			if __obf_8be1afe49afe1e06 := string(__obf_b69dd8ffbe131543.Field(__obf_9f418fe88df9aaef).Tag.Get(__obf_fe32fb81ebf474d0)); __obf_8be1afe49afe1e06 == "" {
				if __obf_e0d28168269174e4 {
					__obf_b3fed84695b938c9 = ToCamel(__obf_b69dd8ffbe131543.Field(__obf_9f418fe88df9aaef).Name)
				} else {
					__obf_b3fed84695b938c9 = __obf_b69dd8ffbe131543.Field(__obf_9f418fe88df9aaef).Name
				}
			} else {
				__obf_df0a7558b9c581fd = strings.Split(__obf_8be1afe49afe1e06, __obf_68cfc5f67cac1404)
				__obf_b3fed84695b938c9 = __obf_df0a7558b9c581fd[0]
			}
			__obf_5691d905931ddccd.
				WriteString(__obf_b3fed84695b938c9)
			__obf_5691d905931ddccd.
				WriteString(" ")
			switch __obf_b69dd8ffbe131543.Field(__obf_9f418fe88df9aaef).Type.Name() {
			case "int8":
				__obf_5691d905931ddccd.
					WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_5691d905931ddccd.
					WriteString("INT")
			case "int64":
				__obf_5691d905931ddccd.
					WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_5691d905931ddccd.
					WriteString("VARCHAR(50)")
			}

			if len(__obf_df0a7558b9c581fd) > 1 {
				__obf_5691d905931ddccd.
					WriteString(" ")
				__obf_5691d905931ddccd.
					WriteString(strings.Join(__obf_df0a7558b9c581fd[1:], " "))
			}

			if __obf_9f418fe88df9aaef+1 != __obf_bb71ca1556579102 {
				__obf_5691d905931ddccd.
					WriteString(",")
			}
			__obf_5691d905931ddccd.
				WriteString("\n")
		}
		__obf_5691d905931ddccd.
			WriteString(");\n\n")
	}
	fmt.Println(__obf_5691d905931ddccd.String())
}

func SaveImage(__obf_60cdefdebc1234da image.Image, __obf_c93615f02c5dca64 uint, __obf_b9a84a7d99601d2c string) error {
	__obf_edf1039d3be58eda, __obf_74916b80241ef1ff := os.Create(__obf_b9a84a7d99601d2c)
	if __obf_74916b80241ef1ff != nil {
		return __obf_74916b80241ef1ff
	}
	defer __obf_edf1039d3be58eda.Close()
	return jpeg.Encode(__obf_edf1039d3be58eda, resize.Resize(__obf_c93615f02c5dca64, 0, __obf_60cdefdebc1234da, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_a1d01074eaf8ad7a string) bool {
	return strings.Contains(__obf_a1d01074eaf8ad7a, ".") && net.ParseIP(__obf_a1d01074eaf8ad7a) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_a1d01074eaf8ad7a string) bool {
	return strings.Contains(__obf_a1d01074eaf8ad7a, ":") && net.ParseIP(__obf_a1d01074eaf8ad7a) != nil
}

func AnyToBytes(__obf_c76af479674ee3d0 any) ([]byte, error) {
	return msgpack.Marshal(__obf_c76af479674ee3d0)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_78314a8a6680c2e3 []byte) (__obf_96894de1e88392eb T, __obf_74916b80241ef1ff error) {
	__obf_74916b80241ef1ff = msgpack.Unmarshal(__obf_78314a8a6680c2e3, &__obf_96894de1e88392eb)
	return
}

func Loadyaml(__obf_8838aea65f47be65 string, __obf_4af871183799e309 any) {
	__obf_4f5a11ae844c76f9, __obf_74916b80241ef1ff := os.ReadFile(__obf_8838aea65f47be65)
	if __obf_74916b80241ef1ff != nil {
		log.Fatalln(__obf_74916b80241ef1ff)
	}
	__obf_74916b80241ef1ff = yaml.UnmarshalStrict(__obf_4f5a11ae844c76f9, __obf_4af871183799e309)
	if __obf_74916b80241ef1ff != nil {
		log.Fatalln(__obf_74916b80241ef1ff)
	}
}

func ToAnyList[T any](__obf_fbc6f697bebfa471 []T) []any {
	__obf_9ea8079c893a3f73 := make([]any, len(__obf_fbc6f697bebfa471))
	for __obf_9f418fe88df9aaef, __obf_c76af479674ee3d0 := range __obf_fbc6f697bebfa471 {
		__obf_9ea8079c893a3f73[__obf_9f418fe88df9aaef] = __obf_c76af479674ee3d0
	}
	return __obf_9ea8079c893a3f73
}

func TimeParse(__obf_38ee8bbd43ec6aa3 string) time.Time {
	__obf_38ee8bbd43ec6aa3 = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_38ee8bbd43ec6aa3)
	__obf_3c9398873f01bdde, _ := time.ParseInLocation("2006-01-02 15:04", __obf_38ee8bbd43ec6aa3, time.Local)
	return __obf_3c9398873f01bdde
}
