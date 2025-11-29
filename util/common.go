package __obf_d7b39e56b82f7f57

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
	__obf_6e1ddda33ce54ba3 = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_c2744e465ce68a8d any) ([]byte, error) {
	return __obf_6e1ddda33ce54ba3.Marshal(__obf_c2744e465ce68a8d)
}

func EncodeString(__obf_c2744e465ce68a8d any) string {
	if __obf_9513491cd9d9271e, __obf_3ccf8c32165eeb13 := __obf_6e1ddda33ce54ba3.Marshal(__obf_c2744e465ce68a8d); __obf_3ccf8c32165eeb13 == nil {
		return string(__obf_9513491cd9d9271e)
	}
	return ""
}

func Decode(__obf_d30bdb07cfba0735 string, __obf_c2744e465ce68a8d any) error {
	return __obf_6e1ddda33ce54ba3.UnmarshalFromString(__obf_d30bdb07cfba0735, __obf_c2744e465ce68a8d)
}

func DecodeByte(__obf_b2cb60631f053d14 []byte, __obf_c2744e465ce68a8d any) error {
	return __obf_6e1ddda33ce54ba3.Unmarshal(__obf_b2cb60631f053d14, __obf_c2744e465ce68a8d)
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

func PrefixInArray(__obf_75d228b82bae384c string, __obf_1461347d2e281479 []string) bool {
	for __obf_329f0eb878b8f7b5 := range __obf_1461347d2e281479 {
		if strings.HasPrefix(__obf_75d228b82bae384c, __obf_1461347d2e281479[__obf_329f0eb878b8f7b5]) {
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

var __obf_da33052e913c8215 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenRandom 获得随机字符串
func GenRandom(__obf_2688d2e7e665fec7 int) string {
	__obf_b2cb60631f053d14 := make([]byte, __obf_2688d2e7e665fec7)
	for __obf_329f0eb878b8f7b5 := range __obf_b2cb60631f053d14 {
		__obf_b2cb60631f053d14[__obf_329f0eb878b8f7b5] = __obf_da33052e913c8215[rand.IntN(len(__obf_da33052e913c8215))]
	}
	return string(__obf_b2cb60631f053d14)
}

func RemoveIndex(__obf_40384ad7789a1f57 []string, __obf_e694c4368fc1e60b int) []string {
	return append(__obf_40384ad7789a1f57[:__obf_e694c4368fc1e60b], __obf_40384ad7789a1f57[__obf_e694c4368fc1e60b+1:]...)
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
func ToCamel(__obf_d30bdb07cfba0735 string) (__obf_f5b0384ddb3578f7 string) {
	__obf_a8a19fd6fa65474d := []rune(__obf_d30bdb07cfba0735)
	__obf_f5b0384ddb3578f7 = __obf_d30bdb07cfba0735[0:1]
	if __obf_a8a19fd6fa65474d[0] >= 97 && __obf_a8a19fd6fa65474d[0] <= 122 {
		__obf_f5b0384ddb3578f7 = string(__obf_a8a19fd6fa65474d[0] - 32)
	}
	__obf_c21c8e0fb673015e := len(__obf_a8a19fd6fa65474d)
	for __obf_329f0eb878b8f7b5 := 1; __obf_329f0eb878b8f7b5 < __obf_c21c8e0fb673015e; __obf_329f0eb878b8f7b5++ {
		if __obf_a8a19fd6fa65474d[__obf_329f0eb878b8f7b5] == 95 && __obf_a8a19fd6fa65474d[__obf_329f0eb878b8f7b5+1] >= 97 && __obf_a8a19fd6fa65474d[__obf_329f0eb878b8f7b5+1] <= 122 {
			__obf_a8a19fd6fa65474d[ //过滤下划线
			__obf_329f0eb878b8f7b5+1] -= 32
		} else {
			__obf_f5b0384ddb3578f7 += string(__obf_a8a19fd6fa65474d[__obf_329f0eb878b8f7b5])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_d30bdb07cfba0735 string) (__obf_f5b0384ddb3578f7 string) {
	__obf_a8a19fd6fa65474d := []rune(__obf_d30bdb07cfba0735)
	__obf_f5b0384ddb3578f7 = __obf_d30bdb07cfba0735[0:1]
	if __obf_a8a19fd6fa65474d[0] >= 65 && __obf_a8a19fd6fa65474d[0] <= 90 {
		__obf_f5b0384ddb3578f7 = string(__obf_a8a19fd6fa65474d[0] + 32)
	}
	__obf_2688d2e7e665fec7 := len(__obf_a8a19fd6fa65474d)
	for __obf_329f0eb878b8f7b5 := 1; __obf_329f0eb878b8f7b5 < __obf_2688d2e7e665fec7; __obf_329f0eb878b8f7b5++ {
		if __obf_a8a19fd6fa65474d[__obf_329f0eb878b8f7b5] >= 65 && __obf_a8a19fd6fa65474d[__obf_329f0eb878b8f7b5] <= 90 {
			__obf_a8a19fd6fa65474d[ //大写变小写
			__obf_329f0eb878b8f7b5] += 32
			__obf_f5b0384ddb3578f7 += "_"
		}
		__obf_f5b0384ddb3578f7 += string(__obf_a8a19fd6fa65474d[__obf_329f0eb878b8f7b5])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_d30bdb07cfba0735 string, __obf_2e4aecebfa6b396a int, __obf_2688d2e7e665fec7 int) string {
	__obf_8dcbe605285dbcb3 := // 将字符串转换为rune切片，以正确处理多字节字符
		[]rune(__obf_d30bdb07cfba0735)
	__obf_3e19bf522273cec0 := len(__obf_8dcbe605285dbcb3)

	// 处理n为负数或0的无效情况
	if __obf_2688d2e7e665fec7 <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_2e4aecebfa6b396a < 0 {
		__obf_2e4aecebfa6b396a = __obf_3e19bf522273cec0 + __obf_2e4aecebfa6b396a
	}

	// 边界检查：确保start在有效范围内
	if __obf_2e4aecebfa6b396a < 0 || __obf_2e4aecebfa6b396a >= __obf_3e19bf522273cec0 {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_8dcbe605285dbcb3[__obf_2e4aecebfa6b396a:min(__obf_2e4aecebfa6b396a+__obf_2688d2e7e665fec7, __obf_3e19bf522273cec0)])
}

func ASCII(__obf_403793979a29000e rune) rune {
	switch {
	case 97 <= __obf_403793979a29000e && __obf_403793979a29000e <= 122:
		return __obf_403793979a29000e - 32
	case 65 <= __obf_403793979a29000e && __obf_403793979a29000e <= 90:
		return __obf_403793979a29000e + 32
	default:
		return __obf_403793979a29000e
	}
}

func IndexString(__obf_d30bdb07cfba0735 string, __obf_bbd2c7f8b6865e3b rune, __obf_d0c8c9897e8d68f1 int) string {
	__obf_c51bdd7900d51020 := []rune(__obf_d30bdb07cfba0735)
	var __obf_5be8a572dd692943 bytes.Buffer
	var __obf_2688d2e7e665fec7 int
	for __obf_329f0eb878b8f7b5, __obf_7ab77ee5165bad1b := 0, len(__obf_c51bdd7900d51020); __obf_329f0eb878b8f7b5 < __obf_7ab77ee5165bad1b; __obf_329f0eb878b8f7b5++ {
		if __obf_c51bdd7900d51020[__obf_329f0eb878b8f7b5] == __obf_bbd2c7f8b6865e3b {
			__obf_2688d2e7e665fec7 += 1
		}
		if __obf_2688d2e7e665fec7 == __obf_d0c8c9897e8d68f1 {
			break
		}
		__obf_5be8a572dd692943.
			WriteRune(__obf_c51bdd7900d51020[__obf_329f0eb878b8f7b5])
	}
	return __obf_5be8a572dd692943.String()
}

func LastIndexString(__obf_def516632b2b7340, __obf_73123f8d71e1bc14 string) string {
	__obf_40384ad7789a1f57 := strings.Split(__obf_def516632b2b7340, __obf_73123f8d71e1bc14)
	if __obf_2688d2e7e665fec7 := len(__obf_40384ad7789a1f57); __obf_2688d2e7e665fec7 > 1 {
		return __obf_40384ad7789a1f57[__obf_2688d2e7e665fec7-2]
	}
	return ""
}

func IsEmpty(__obf_f470d4efe5962aa7 any) bool {
	__obf_a30d33f1a4aa6e72 := reflect.ValueOf(__obf_f470d4efe5962aa7)
	if __obf_a30d33f1a4aa6e72.Kind() == reflect.Ptr {
		__obf_a30d33f1a4aa6e72 = __obf_a30d33f1a4aa6e72.Elem()
	}
	return __obf_a30d33f1a4aa6e72.Interface() == reflect.Zero(__obf_a30d33f1a4aa6e72.Type()).Interface()
}

func MillisecondToDateString(__obf_1a03adda1195e896 int64) string {
	return time.Unix(__obf_1a03adda1195e896, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_1a03adda1195e896 int64) string {
	return time.Unix(__obf_1a03adda1195e896, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_80c43ef83fa8226b string) (__obf_a65865c83781bbce, __obf_ebe53606f4d9bacf string) {
	__obf_80c43ef83fa8226b = filepath.Base(__obf_80c43ef83fa8226b)
	__obf_ebe53606f4d9bacf = filepath.Ext(__obf_80c43ef83fa8226b)
	__obf_a65865c83781bbce = strings.TrimSuffix(__obf_80c43ef83fa8226b, __obf_ebe53606f4d9bacf)
	return __obf_a65865c83781bbce,

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
		__obf_ebe53606f4d9bacf
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(__obf_11424d5d3daa102e string) (int64, error) {
	if __obf_1b64ea375184b406, __obf_3ccf8c32165eeb13 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_11424d5d3daa102e, Loc); __obf_3ccf8c32165eeb13 != nil {
		return 0, nil
	} else {
		__obf_1b64ea375184b406 = __obf_1b64ea375184b406.AddDate(0, 0, -__obf_1b64ea375184b406.Day()+1)
		return GetZeroTime(__obf_1b64ea375184b406).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_11424d5d3daa102e string) (int64, int64, error) {
	if __obf_1b64ea375184b406, __obf_3ccf8c32165eeb13 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_11424d5d3daa102e, Loc); __obf_3ccf8c32165eeb13 != nil {
		return 0, 0, nil
	} else {
		__obf_1b64ea375184b406 = __obf_1b64ea375184b406.AddDate(0, 0, -__obf_1b64ea375184b406.Day()+1)
		return GetZeroTime(__obf_1b64ea375184b406).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_1b64ea375184b406).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_fff126686fc07d47 time.Time) time.Time {
	return time.Date(__obf_fff126686fc07d47.Year(), __obf_fff126686fc07d47.Month(), __obf_fff126686fc07d47.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_1a03adda1195e896 int64) int64 {
	__obf_fd9e4954cd742ff3 := time.Unix(__obf_1a03adda1195e896, 0)
	__obf_d5208cec0fa78da6, __obf_979e66084e96cf8a, __obf_fff126686fc07d47 := __obf_fd9e4954cd742ff3.Date()
	return time.Date(__obf_d5208cec0fa78da6, __obf_979e66084e96cf8a, __obf_fff126686fc07d47, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_1a03adda1195e896 int64) int64 {
	__obf_fd9e4954cd742ff3 := time.Unix(__obf_1a03adda1195e896, 0)
	return __obf_fd9e4954cd742ff3.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_d5208cec0fa78da6, __obf_979e66084e96cf8a, __obf_fff126686fc07d47 := time.Now().Date()
	return time.Date(__obf_d5208cec0fa78da6, __obf_979e66084e96cf8a, __obf_fff126686fc07d47, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_11424d5d3daa102e string) (int64, int64) {
	__obf_780c484923edf5af, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_11424d5d3daa102e+" 00:00:00", Loc)
	min := __obf_780c484923edf5af.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_ec31fa56a2658129 int) (int64, int64) {
	__obf_11424d5d3daa102e := time.Now()
	__obf_fd9e4954cd742ff3 := __obf_11424d5d3daa102e.AddDate(0, 0, __obf_ec31fa56a2658129)
	__obf_d5208cec0fa78da6, __obf_979e66084e96cf8a, __obf_fff126686fc07d47 := __obf_fd9e4954cd742ff3.Date()
	min := time.Date(__obf_d5208cec0fa78da6, __obf_979e66084e96cf8a, __obf_fff126686fc07d47, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_11424d5d3daa102e.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_c38f6788e4938be4, __obf_aa33859662557178 string) bool {
	__obf_33c469af51fcadd1 := func(__obf_40384ad7789a1f57 string) string {
		return fmt.Sprintf(",%s,", __obf_40384ad7789a1f57)
	}
	return strings.Contains(__obf_33c469af51fcadd1(__obf_c38f6788e4938be4), __obf_33c469af51fcadd1(__obf_aa33859662557178))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_be36ee5ae4fcc638 string, __obf_efee66e877d6cb9a ...string) bool {
	if any {
		for _, __obf_36b2df5f73ce8ef9 := range __obf_efee66e877d6cb9a {
			if Contain(__obf_be36ee5ae4fcc638, __obf_36b2df5f73ce8ef9) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_36b2df5f73ce8ef9 := range __obf_efee66e877d6cb9a {
			if !Contain(__obf_be36ee5ae4fcc638, __obf_36b2df5f73ce8ef9) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_40384ad7789a1f57 []any, __obf_e694c4368fc1e60b int) []any {
	return slices.Delete(__obf_40384ad7789a1f57, __obf_e694c4368fc1e60b, __obf_e694c4368fc1e60b+1)
}

func String2Int8(__obf_d30bdb07cfba0735 string) int8 {
	__obf_b12c6e7ea34882c7, __obf_3ccf8c32165eeb13 := strconv.ParseInt(__obf_d30bdb07cfba0735, 10, 8)
	if __obf_3ccf8c32165eeb13 == nil {
		return int8(__obf_b12c6e7ea34882c7)
	}
	return 0
}

func String2Int32(__obf_d30bdb07cfba0735 string) int32 {
	__obf_b12c6e7ea34882c7, __obf_3ccf8c32165eeb13 := strconv.ParseInt(__obf_d30bdb07cfba0735, 10, 32)
	if __obf_3ccf8c32165eeb13 == nil {
		return int32(__obf_b12c6e7ea34882c7)
	}
	return 0
}

func String2Int64(__obf_d30bdb07cfba0735 string) int8 {
	__obf_b12c6e7ea34882c7, __obf_3ccf8c32165eeb13 := strconv.ParseInt(__obf_d30bdb07cfba0735, 10, 64)
	if __obf_3ccf8c32165eeb13 == nil {
		return int8(__obf_b12c6e7ea34882c7)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_51054a15aa6a7b13, __obf_2e984019fc8cf183 int) (__obf_2e4aecebfa6b396a, __obf_76f7f0f42f292af8 time.Time) {
	if __obf_51054a15aa6a7b13 > 0 && __obf_2e984019fc8cf183 > 0 {
		__obf_2e4aecebfa6b396a = time.Date(__obf_51054a15aa6a7b13, 1, 0, 0, 0, 0, 0, Loc)
		__obf_1edf50299740c806 := // 第一天是周几
			int(__obf_2e4aecebfa6b396a.AddDate(0, 0, 1).Weekday())
		__obf_2f076fd1684e1b96 := // 当年第一周有几天
			1
		if __obf_1edf50299740c806 != 0 {
			__obf_2f076fd1684e1b96 = 7 - __obf_1edf50299740c806 + 1
		}
		if __obf_2e984019fc8cf183 == 1 {
			__obf_76f7f0f42f292af8 = __obf_2e4aecebfa6b396a.AddDate(0, 0, __obf_2f076fd1684e1b96)
		} else {
			__obf_76f7f0f42f292af8 = __obf_2e4aecebfa6b396a.AddDate(0, 0, __obf_2f076fd1684e1b96+(__obf_2e984019fc8cf183-1)*7)
			__obf_2e4aecebfa6b396a = __obf_76f7f0f42f292af8.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_11424d5d3daa102e time.Time) (__obf_51054a15aa6a7b13, __obf_2e984019fc8cf183, __obf_6eadc5466a958f5b int) {
	__obf_51054a15aa6a7b13 = __obf_11424d5d3daa102e.Year()
	__obf_6eadc5466a958f5b = int(__obf_11424d5d3daa102e.Weekday())
	__obf_f375c2394c087ffc := __obf_11424d5d3daa102e.YearDay()
	__obf_33ce76dfebed9785 := __obf_11424d5d3daa102e.AddDate(0, 0, -__obf_f375c2394c087ffc+1)
	__obf_1edf50299740c806 := int(__obf_33ce76dfebed9785.Weekday())
	__obf_2f076fd1684e1b96 := // 当年第一周有几天
		1
	if __obf_1edf50299740c806 != 0 {
		__obf_2f076fd1684e1b96 = 7 - __obf_1edf50299740c806 + 1
	}
	if __obf_f375c2394c087ffc <= __obf_2f076fd1684e1b96 {
		__obf_2e984019fc8cf183 = 1
	} else {
		__obf_2e984019fc8cf183 = (__obf_f375c2394c087ffc-__obf_2f076fd1684e1b96)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_3b31a0fb4c87a405 []byte) map[string]string {
	__obf_1a6f2ab955d93240 := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_7731156423981cba := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_0e3b394c62afd88e []byte
	if __obf_0e3b394c62afd88e = __obf_1a6f2ab955d93240.ReplaceAll(__obf_3b31a0fb4c87a405, []byte("")); len(__obf_0e3b394c62afd88e) < 6 {
		return nil
	}
	__obf_f3eca6dca5babea5 := __obf_7731156423981cba.FindAllSubmatch(__obf_0e3b394c62afd88e[6:len(__obf_0e3b394c62afd88e)-7], -1)
	__obf_c8857c8f0f171829 := map[string]string{}
	if __obf_3d3bff4aaac2c6df := __obf_1a6f2ab955d93240.FindSubmatch(__obf_3b31a0fb4c87a405)[1]; len(__obf_3d3bff4aaac2c6df) != 0 {
		__obf_476ff7275dac41cc := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_3d3bff4aaac2c6df, -1)
		for _, __obf_a30d33f1a4aa6e72 := range __obf_476ff7275dac41cc {
			__obf_c8857c8f0f171829[string(__obf_a30d33f1a4aa6e72[1])] = string(__obf_a30d33f1a4aa6e72[2])
		}
	}

	for _, __obf_a30d33f1a4aa6e72 := range __obf_f3eca6dca5babea5 {
		__obf_c8857c8f0f171829[string(__obf_a30d33f1a4aa6e72[1])] = string(__obf_a30d33f1a4aa6e72[2])
	}
	return __obf_c8857c8f0f171829
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_6c455d2315059b4b time.Time) (int64, int64) {
	__obf_53f5dce3b8cda70f := GetZeroTime(__obf_6c455d2315059b4b).Unix() / 600
	return __obf_53f5dce3b8cda70f, __obf_53f5dce3b8cda70f + 143
}

func Abs(__obf_2688d2e7e665fec7 int64) int64 {
	__obf_d5208cec0fa78da6 := __obf_2688d2e7e665fec7 >> 63
	return (__obf_2688d2e7e665fec7 ^ __obf_d5208cec0fa78da6) - __obf_d5208cec0fa78da6
}

func Number2String(__obf_2688d2e7e665fec7 any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_2688d2e7e665fec7 := __obf_2688d2e7e665fec7.(type) {
	case int:
		return strconv.Itoa(__obf_2688d2e7e665fec7)
	case int32:
		return strconv.FormatInt(int64(__obf_2688d2e7e665fec7), 10)
	case int64:
		return strconv.FormatInt(__obf_2688d2e7e665fec7, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_2688d2e7e665fec7), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_2688d2e7e665fec7, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_d30bdb07cfba0735 any, __obf_bbd2c7f8b6865e3b string) string {
	if __obf_d30bdb07cfba0735 != nil && __obf_d30bdb07cfba0735.(string) != "" {
		// return sep + str + sep
		return __obf_d30bdb07cfba0735.(string)
	}
	return __obf_bbd2c7f8b6865e3b
}

func SortRange(__obf_979e66084e96cf8a map[string]any, __obf_1615e38c49fb8375 func(int, string)) {
	var __obf_c593e4f1f731ecd9 []string
	for __obf_9045b7561a7c656e := range __obf_979e66084e96cf8a {
		__obf_c593e4f1f731ecd9 = append(__obf_c593e4f1f731ecd9, __obf_9045b7561a7c656e)
	}
	sort.Strings(__obf_c593e4f1f731ecd9)
	for __obf_329f0eb878b8f7b5, __obf_9045b7561a7c656e := range __obf_c593e4f1f731ecd9 {
		__obf_1615e38c49fb8375(__obf_329f0eb878b8f7b5, __obf_9045b7561a7c656e)
	}
}

func HasField(__obf_3cd6b4e452a805a3 reflect.Value, __obf_a65865c83781bbce string) bool {

	if __obf_40384ad7789a1f57 := __obf_3cd6b4e452a805a3.FieldByNameFunc(func(__obf_2688d2e7e665fec7 string) bool {
		return strings.EqualFold(__obf_2688d2e7e665fec7, __obf_a65865c83781bbce)
	}); __obf_40384ad7789a1f57.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_3cd6b4e452a805a3 reflect.Value, __obf_a65865c83781bbce string) reflect.Value {
	if __obf_40384ad7789a1f57 := __obf_3cd6b4e452a805a3.FieldByNameFunc(func(__obf_2688d2e7e665fec7 string) bool {
		return strings.EqualFold(__obf_2688d2e7e665fec7, __obf_a65865c83781bbce)
	}); __obf_40384ad7789a1f57.IsValid() {
		return __obf_40384ad7789a1f57
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_3cd6b4e452a805a3 reflect.Value, __obf_a65865c83781bbce string, __obf_75d228b82bae384c any) bool {

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
	if __obf_40384ad7789a1f57 := __obf_3cd6b4e452a805a3.FieldByNameFunc(func(__obf_2688d2e7e665fec7 string) bool {
		return strings.EqualFold(__obf_2688d2e7e665fec7, __obf_a65865c83781bbce)
	}); __obf_40384ad7789a1f57.IsValid() {
		if __obf_429c22e552137cdf := __obf_40384ad7789a1f57.Type(); __obf_429c22e552137cdf == reflect.TypeOf(__obf_75d228b82bae384c) {
			__obf_40384ad7789a1f57.
				Set(reflect.ValueOf(__obf_75d228b82bae384c))
		} else {
			__obf_40384ad7789a1f57.
				Set(reflect.ValueOf(__obf_75d228b82bae384c).Convert(__obf_429c22e552137cdf))
		}
		return true
	}
	return false
}

func CopyMap(__obf_979e66084e96cf8a map[string]any) map[string]any {
	__obf_476ff7275dac41cc := make(map[string]any)
	for __obf_9045b7561a7c656e, __obf_a30d33f1a4aa6e72 := range __obf_979e66084e96cf8a {
		if __obf_1d12939c5d335950, __obf_6051cf6d55fb7047 := __obf_a30d33f1a4aa6e72.(map[string]any); __obf_6051cf6d55fb7047 {
			__obf_476ff7275dac41cc[__obf_9045b7561a7c656e] = CopyMap(__obf_1d12939c5d335950)
		} else {
			__obf_476ff7275dac41cc[__obf_9045b7561a7c656e] = __obf_a30d33f1a4aa6e72
		}
	}

	return __obf_476ff7275dac41cc
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_6517b70edd186159, __obf_d46b900dd28fb4b4 string, __obf_35cf1b44680f7a79 bool, __obf_41f41023f92700c0 ...any) {
	var __obf_5283a26e14d727c7 strings.Builder
	var __obf_4f6dfb1a90d2328d string
	var __obf_0e7c92f676694631 []string
	for _, __obf_85732cdadbba20c7 := range __obf_41f41023f92700c0 {
		__obf_40384ad7789a1f57 := reflect.TypeOf(__obf_85732cdadbba20c7)
		__obf_5283a26e14d727c7.
			WriteString(`CREATE TABLE `)
		__obf_5283a26e14d727c7.
			WriteString(__obf_40384ad7789a1f57.Name())
		__obf_5283a26e14d727c7.
			WriteString(" (\n")
		__obf_bcee7aab5bf7435b := __obf_40384ad7789a1f57.NumField()
		for __obf_329f0eb878b8f7b5 := 0; __obf_329f0eb878b8f7b5 < __obf_bcee7aab5bf7435b; __obf_329f0eb878b8f7b5++ {
			__obf_5283a26e14d727c7.
				WriteString("    ")
			__obf_0e7c92f676694631 = nil
			if __obf_c40517d9cd2f075e := string(__obf_40384ad7789a1f57.Field(__obf_329f0eb878b8f7b5).Tag.Get(__obf_6517b70edd186159)); __obf_c40517d9cd2f075e == "" {
				if __obf_35cf1b44680f7a79 {
					__obf_4f6dfb1a90d2328d = ToCamel(__obf_40384ad7789a1f57.Field(__obf_329f0eb878b8f7b5).Name)
				} else {
					__obf_4f6dfb1a90d2328d = __obf_40384ad7789a1f57.Field(__obf_329f0eb878b8f7b5).Name
				}
			} else {
				__obf_0e7c92f676694631 = strings.Split(__obf_c40517d9cd2f075e, __obf_d46b900dd28fb4b4)
				__obf_4f6dfb1a90d2328d = __obf_0e7c92f676694631[0]
			}
			__obf_5283a26e14d727c7.
				WriteString(__obf_4f6dfb1a90d2328d)
			__obf_5283a26e14d727c7.
				WriteString(" ")
			switch __obf_40384ad7789a1f57.Field(__obf_329f0eb878b8f7b5).Type.Name() {
			case "int8":
				__obf_5283a26e14d727c7.
					WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_5283a26e14d727c7.
					WriteString("INT")
			case "int64":
				__obf_5283a26e14d727c7.
					WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_5283a26e14d727c7.
					WriteString("VARCHAR(50)")
			}

			if len(__obf_0e7c92f676694631) > 1 {
				__obf_5283a26e14d727c7.
					WriteString(" ")
				__obf_5283a26e14d727c7.
					WriteString(strings.Join(__obf_0e7c92f676694631[1:], " "))
			}

			if __obf_329f0eb878b8f7b5+1 != __obf_bcee7aab5bf7435b {
				__obf_5283a26e14d727c7.
					WriteString(",")
			}
			__obf_5283a26e14d727c7.
				WriteString("\n")
		}
		__obf_5283a26e14d727c7.
			WriteString(");\n\n")
	}
	fmt.Println(__obf_5283a26e14d727c7.String())
}

func SaveImage(__obf_22b08ca1f395bd75 image.Image, __obf_0c22ca6cb41b6bd2 uint, __obf_650069cf7416165c string) error {
	__obf_ce72d0ac7203bc86, __obf_3ccf8c32165eeb13 := os.Create(__obf_650069cf7416165c)
	if __obf_3ccf8c32165eeb13 != nil {
		return __obf_3ccf8c32165eeb13
	}
	defer __obf_ce72d0ac7203bc86.Close()
	return jpeg.Encode(__obf_ce72d0ac7203bc86, resize.Resize(__obf_0c22ca6cb41b6bd2, 0, __obf_22b08ca1f395bd75, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_d30bdb07cfba0735 string) bool {
	return strings.Contains(__obf_d30bdb07cfba0735, ".") && net.ParseIP(__obf_d30bdb07cfba0735) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_d30bdb07cfba0735 string) bool {
	return strings.Contains(__obf_d30bdb07cfba0735, ":") && net.ParseIP(__obf_d30bdb07cfba0735) != nil
}

func AnyToBytes(__obf_a30d33f1a4aa6e72 any) ([]byte, error) {
	return msgpack.Marshal(__obf_a30d33f1a4aa6e72)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_75d228b82bae384c []byte) (__obf_c2744e465ce68a8d T, __obf_3ccf8c32165eeb13 error) {
	__obf_3ccf8c32165eeb13 = msgpack.Unmarshal(__obf_75d228b82bae384c, &__obf_c2744e465ce68a8d)
	return
}

func Loadyaml(__obf_eda86051d77a60ee string, __obf_f184a2774879f9c3 any) {
	__obf_aa33859662557178, __obf_3ccf8c32165eeb13 := os.ReadFile(__obf_eda86051d77a60ee)
	if __obf_3ccf8c32165eeb13 != nil {
		log.Fatalln(__obf_3ccf8c32165eeb13)
	}
	__obf_3ccf8c32165eeb13 = yaml.UnmarshalStrict(__obf_aa33859662557178, __obf_f184a2774879f9c3)
	if __obf_3ccf8c32165eeb13 != nil {
		log.Fatalln(__obf_3ccf8c32165eeb13)
	}
}

func ToAnyList[T any](__obf_0e1f1509062e69c4 []T) []any {
	__obf_d63f47e8f8b94e81 := make([]any, len(__obf_0e1f1509062e69c4))
	for __obf_329f0eb878b8f7b5, __obf_a30d33f1a4aa6e72 := range __obf_0e1f1509062e69c4 {
		__obf_d63f47e8f8b94e81[__obf_329f0eb878b8f7b5] = __obf_a30d33f1a4aa6e72
	}
	return __obf_d63f47e8f8b94e81
}

func TimeParse(__obf_1c09eb251c201b3c string) time.Time {
	__obf_1c09eb251c201b3c = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_1c09eb251c201b3c)
	__obf_fd9e4954cd742ff3, _ := time.ParseInLocation("2006-01-02 15:04", __obf_1c09eb251c201b3c, time.Local)
	return __obf_fd9e4954cd742ff3
}
