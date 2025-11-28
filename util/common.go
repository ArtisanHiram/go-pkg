package __obf_077bcefb098a89af

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
	__obf_24b8a17d9ed5e394 = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Encode(__obf_f8bbf5dce647ada9 any) ([]byte, error) {
	return __obf_24b8a17d9ed5e394.Marshal(__obf_f8bbf5dce647ada9)
}

func EncodeString(__obf_f8bbf5dce647ada9 any) string {
	if __obf_7929a0a4e5454dde, __obf_224415a75b186177 := __obf_24b8a17d9ed5e394.Marshal(__obf_f8bbf5dce647ada9); __obf_224415a75b186177 == nil {
		return string(__obf_7929a0a4e5454dde)
	}
	return ""
}

func Decode(__obf_937cd323c9259c13 string, __obf_f8bbf5dce647ada9 any) error {
	return __obf_24b8a17d9ed5e394.UnmarshalFromString(__obf_937cd323c9259c13, __obf_f8bbf5dce647ada9)
}

func DecodeByte(__obf_3811c39ee1400927 []byte, __obf_f8bbf5dce647ada9 any) error {
	return __obf_24b8a17d9ed5e394.Unmarshal(__obf_3811c39ee1400927, __obf_f8bbf5dce647ada9)
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

func PrefixInArray(__obf_cd007b9a47849e9a string, __obf_bf6a6eefbec13048 []string) bool {
	for __obf_ff7abde39552bac6 := range __obf_bf6a6eefbec13048 {
		if strings.HasPrefix(__obf_cd007b9a47849e9a, __obf_bf6a6eefbec13048[__obf_ff7abde39552bac6]) {
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

var __obf_74e52b43cdf023f9 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var __obf_945a7c712333bc4f = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenRandom 获得随机字符串
func GenRandom(__obf_c705c49030e00b6f int) string {
	__obf_3811c39ee1400927 := make([]byte, __obf_c705c49030e00b6f)
	for __obf_ff7abde39552bac6 := range __obf_3811c39ee1400927 {
		__obf_3811c39ee1400927[__obf_ff7abde39552bac6] = __obf_74e52b43cdf023f9[__obf_945a7c712333bc4f.Intn(len(__obf_74e52b43cdf023f9))]
	}
	return string(__obf_3811c39ee1400927)
}

func RemoveIndex(__obf_2b69c849afdb21c3 []string, __obf_1510b85c5b64ad72 int) []string {
	return append(__obf_2b69c849afdb21c3[:__obf_1510b85c5b64ad72], __obf_2b69c849afdb21c3[__obf_1510b85c5b64ad72+1:]...)
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
func ToCamel(__obf_937cd323c9259c13 string) (__obf_84489a76bf4c5164 string) {
	__obf_d091e2fa54e96939 := []rune(__obf_937cd323c9259c13)
	__obf_84489a76bf4c5164 = __obf_937cd323c9259c13[0:1]
	if __obf_d091e2fa54e96939[0] >= 97 && __obf_d091e2fa54e96939[0] <= 122 {
		__obf_84489a76bf4c5164 = string(__obf_d091e2fa54e96939[0] - 32)
	}

	__obf_e089d723128d4dc7 := len(__obf_d091e2fa54e96939)
	for __obf_ff7abde39552bac6 := 1; __obf_ff7abde39552bac6 < __obf_e089d723128d4dc7; __obf_ff7abde39552bac6++ {
		if __obf_d091e2fa54e96939[__obf_ff7abde39552bac6] == 95 && __obf_d091e2fa54e96939[__obf_ff7abde39552bac6+1] >= 97 && __obf_d091e2fa54e96939[__obf_ff7abde39552bac6+1] <= 122 { //过滤下划线
			__obf_d091e2fa54e96939[__obf_ff7abde39552bac6+1] -= 32
		} else {
			__obf_84489a76bf4c5164 += string(__obf_d091e2fa54e96939[__obf_ff7abde39552bac6])
		}
	}
	return
}

// 下划线式
func ToSnake(__obf_937cd323c9259c13 string) (__obf_84489a76bf4c5164 string) {
	__obf_d091e2fa54e96939 := []rune(__obf_937cd323c9259c13)
	__obf_84489a76bf4c5164 = __obf_937cd323c9259c13[0:1]
	if __obf_d091e2fa54e96939[0] >= 65 && __obf_d091e2fa54e96939[0] <= 90 {
		__obf_84489a76bf4c5164 = string(__obf_d091e2fa54e96939[0] + 32)
	}

	__obf_c705c49030e00b6f := len(__obf_d091e2fa54e96939)
	for __obf_ff7abde39552bac6 := 1; __obf_ff7abde39552bac6 < __obf_c705c49030e00b6f; __obf_ff7abde39552bac6++ {
		if __obf_d091e2fa54e96939[__obf_ff7abde39552bac6] >= 65 && __obf_d091e2fa54e96939[__obf_ff7abde39552bac6] <= 90 { //大写变小写
			__obf_d091e2fa54e96939[__obf_ff7abde39552bac6] += 32
			__obf_84489a76bf4c5164 += "_"
		}
		__obf_84489a76bf4c5164 += string(__obf_d091e2fa54e96939[__obf_ff7abde39552bac6])
	}
	return
}

// Substr 函数用于截取字符串
// str: 原始字符串
// start: 起始索引。若为负数，则从尾部第|start|位开始。
// n: 截取长度
func Substr(__obf_937cd323c9259c13 string, __obf_7d64ca592a1d7345 int, __obf_c705c49030e00b6f int) string {
	// 将字符串转换为rune切片，以正确处理多字节字符
	__obf_4737d1d739735508 := []rune(__obf_937cd323c9259c13)
	__obf_85c14d42f927cfdf := len(__obf_4737d1d739735508)

	// 处理n为负数或0的无效情况
	if __obf_c705c49030e00b6f <= 0 {
		return ""
	}

	// 处理负数起始索引
	if __obf_7d64ca592a1d7345 < 0 {
		__obf_7d64ca592a1d7345 = __obf_85c14d42f927cfdf + __obf_7d64ca592a1d7345
	}

	// 边界检查：确保start在有效范围内
	if __obf_7d64ca592a1d7345 < 0 || __obf_7d64ca592a1d7345 >= __obf_85c14d42f927cfdf {
		return ""
	}

	// 执行截取并返回结果
	return string(__obf_4737d1d739735508[__obf_7d64ca592a1d7345:min(__obf_7d64ca592a1d7345+__obf_c705c49030e00b6f, __obf_85c14d42f927cfdf)])
}

func ASCII(__obf_c07a5c3f92c3565e rune) rune {
	switch {
	case 97 <= __obf_c07a5c3f92c3565e && __obf_c07a5c3f92c3565e <= 122:
		return __obf_c07a5c3f92c3565e - 32
	case 65 <= __obf_c07a5c3f92c3565e && __obf_c07a5c3f92c3565e <= 90:
		return __obf_c07a5c3f92c3565e + 32
	default:
		return __obf_c07a5c3f92c3565e
	}
}

func IndexString(__obf_937cd323c9259c13 string, __obf_53aa887d3a1a5d3b rune, __obf_7f1a5b6b87be8b24 int) string {
	__obf_d7fc655e8eec739e := []rune(__obf_937cd323c9259c13)
	var __obf_8932d3234367d407 bytes.Buffer
	var __obf_c705c49030e00b6f int
	for __obf_ff7abde39552bac6, __obf_9fec9a5165303447 := 0, len(__obf_d7fc655e8eec739e); __obf_ff7abde39552bac6 < __obf_9fec9a5165303447; __obf_ff7abde39552bac6++ {
		if __obf_d7fc655e8eec739e[__obf_ff7abde39552bac6] == __obf_53aa887d3a1a5d3b {
			__obf_c705c49030e00b6f += 1
		}
		if __obf_c705c49030e00b6f == __obf_7f1a5b6b87be8b24 {
			break
		}
		__obf_8932d3234367d407.WriteRune(__obf_d7fc655e8eec739e[__obf_ff7abde39552bac6])
	}
	return __obf_8932d3234367d407.String()
}

func LastIndexString(__obf_1da8943d4f351337, __obf_cd8085fda584726d string) string {
	__obf_2b69c849afdb21c3 := strings.Split(__obf_1da8943d4f351337, __obf_cd8085fda584726d)
	if __obf_c705c49030e00b6f := len(__obf_2b69c849afdb21c3); __obf_c705c49030e00b6f > 1 {
		return __obf_2b69c849afdb21c3[__obf_c705c49030e00b6f-2]
	}
	return ""
}

func IsEmpty(__obf_d7db53b4d4d757d4 any) bool {
	__obf_1c4edfc2acf87521 := reflect.ValueOf(__obf_d7db53b4d4d757d4)
	if __obf_1c4edfc2acf87521.Kind() == reflect.Ptr {
		__obf_1c4edfc2acf87521 = __obf_1c4edfc2acf87521.Elem()
	}
	return __obf_1c4edfc2acf87521.Interface() == reflect.Zero(__obf_1c4edfc2acf87521.Type()).Interface()
}

func MillisecondToDateString(__obf_3ff996fecf3ee8ea int64) string {
	return time.Unix(__obf_3ff996fecf3ee8ea, 0).Format("2006-01-02")
}

func MillisecondToDateHMS(__obf_3ff996fecf3ee8ea int64) string {
	return time.Unix(__obf_3ff996fecf3ee8ea, 0).Format(FORMAT_ISO8601_DATE_TIME)
}

func SplitFilename(__obf_bc6f3ed71ce8766a string) (__obf_1e523056a8de9a9c, __obf_7f0a8d8dc99aad34 string) {
	__obf_bc6f3ed71ce8766a = filepath.Base(__obf_bc6f3ed71ce8766a)
	__obf_7f0a8d8dc99aad34 = filepath.Ext(__obf_bc6f3ed71ce8766a)
	__obf_1e523056a8de9a9c = strings.TrimSuffix(__obf_bc6f3ed71ce8766a, __obf_7f0a8d8dc99aad34)
	return __obf_1e523056a8de9a9c, __obf_7f0a8d8dc99aad34
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
func GetFirstDateOfMonth(__obf_0a2edfa7a86ee54e string) (int64, error) {
	if __obf_28e6a5ad48ab70d5, __obf_224415a75b186177 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_0a2edfa7a86ee54e, Loc); __obf_224415a75b186177 != nil {
		return 0, nil
	} else {
		__obf_28e6a5ad48ab70d5 = __obf_28e6a5ad48ab70d5.AddDate(0, 0, -__obf_28e6a5ad48ab70d5.Day()+1)
		return GetZeroTime(__obf_28e6a5ad48ab70d5).Unix(), nil
	}
}

// 获取传入的时间所在月份的第一天凌晨与最后一天23点59分59秒
func GetFirstAndLastDateOfMonth(__obf_0a2edfa7a86ee54e string) (int64, int64, error) {
	if __obf_28e6a5ad48ab70d5, __obf_224415a75b186177 := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_0a2edfa7a86ee54e, Loc); __obf_224415a75b186177 != nil {
		return 0, 0, nil
	} else {
		__obf_28e6a5ad48ab70d5 = __obf_28e6a5ad48ab70d5.AddDate(0, 0, -__obf_28e6a5ad48ab70d5.Day()+1)
		return GetZeroTime(__obf_28e6a5ad48ab70d5).Unix(), DAY_TS_DIFF + GetZeroTime(__obf_28e6a5ad48ab70d5).AddDate(0, 1, -1).Unix(), nil
	}
}

// 获取某一天的0点时间
func GetZeroTime(__obf_ba749bcf2f186fe5 time.Time) time.Time {
	return time.Date(__obf_ba749bcf2f186fe5.Year(), __obf_ba749bcf2f186fe5.Month(), __obf_ba749bcf2f186fe5.Day(), 0, 0, 0, 0, Loc)
}

func GetZeroTSByTS(__obf_3ff996fecf3ee8ea int64) int64 {
	__obf_f439046764df8d86 := time.Unix(__obf_3ff996fecf3ee8ea, 0)
	__obf_5af83d31b7fcc5d6, __obf_1bc902c81fe07f6f, __obf_ba749bcf2f186fe5 := __obf_f439046764df8d86.Date()
	return time.Date(__obf_5af83d31b7fcc5d6, __obf_1bc902c81fe07f6f, __obf_ba749bcf2f186fe5, 0, 0, 0, 0, Loc).Unix()
}

func GetMonthEndTSByTS(__obf_3ff996fecf3ee8ea int64) int64 {
	__obf_f439046764df8d86 := time.Unix(__obf_3ff996fecf3ee8ea, 0)
	return __obf_f439046764df8d86.AddDate(0, 1, -1).Unix() + DAY_TS_DIFF
}

func GetZeroTS() int64 {
	__obf_5af83d31b7fcc5d6, __obf_1bc902c81fe07f6f, __obf_ba749bcf2f186fe5 := time.Now().Date()
	return time.Date(__obf_5af83d31b7fcc5d6, __obf_1bc902c81fe07f6f, __obf_ba749bcf2f186fe5, 0, 0, 0, 0, Loc).Unix()
}

func GetTS(__obf_0a2edfa7a86ee54e string) (int64, int64) {
	__obf_535c6a3f8240b2a1, _ := time.ParseInLocation(FORMAT_ISO8601_DATE_TIME, __obf_0a2edfa7a86ee54e+" 00:00:00", Loc)
	min := __obf_535c6a3f8240b2a1.Unix()
	return min, min + DAY_TS_DIFF
}

func GetDiffTS(__obf_62f5286119abd293 int) (int64, int64) {
	__obf_0a2edfa7a86ee54e := time.Now()
	__obf_f439046764df8d86 := __obf_0a2edfa7a86ee54e.AddDate(0, 0, __obf_62f5286119abd293)
	__obf_5af83d31b7fcc5d6, __obf_1bc902c81fe07f6f, __obf_ba749bcf2f186fe5 := __obf_f439046764df8d86.Date()
	min := time.Date(__obf_5af83d31b7fcc5d6, __obf_1bc902c81fe07f6f, __obf_ba749bcf2f186fe5, 0, 0, 0, 0, Loc).Unix()
	return min, __obf_0a2edfa7a86ee54e.Unix()
}

// p 为以逗号隔开的字符串
func Contain(__obf_4f222390be4477dc, __obf_19ee06338066807d string) bool {
	__obf_8916ea5a9becb604 := func(__obf_2b69c849afdb21c3 string) string {
		return fmt.Sprintf(",%s,", __obf_2b69c849afdb21c3)
	}
	return strings.Contains(__obf_8916ea5a9becb604(__obf_4f222390be4477dc), __obf_8916ea5a9becb604(__obf_19ee06338066807d))
}

// s1是否包含s2; 其中s1为逗号隔开的字符串，s2为可能出现在s1中，逗号分隔的某个子串;
// any 是否只包含其中一个即可，还是必须包含所有
func SubContain(any bool, __obf_8a8f593f1268fbf6 string, __obf_73fc59cae8519e6c ...string) bool {
	if any {
		for _, __obf_e05175036b1dc5fa := range __obf_73fc59cae8519e6c {
			if Contain(__obf_8a8f593f1268fbf6, __obf_e05175036b1dc5fa) {
				return true
			}
			continue
		}
		return false
	} else {
		for _, __obf_e05175036b1dc5fa := range __obf_73fc59cae8519e6c {
			if !Contain(__obf_8a8f593f1268fbf6, __obf_e05175036b1dc5fa) {
				return false
			}
			continue
		}
		return true
	}
}

func SliceRemove(__obf_2b69c849afdb21c3 []any, __obf_1510b85c5b64ad72 int) []any {
	return slices.Delete(__obf_2b69c849afdb21c3, __obf_1510b85c5b64ad72, __obf_1510b85c5b64ad72+1)
}

func String2Int8(__obf_937cd323c9259c13 string) int8 {
	__obf_97fc9da0ff9422b3, __obf_224415a75b186177 := strconv.ParseInt(__obf_937cd323c9259c13, 10, 8)
	if __obf_224415a75b186177 == nil {
		return int8(__obf_97fc9da0ff9422b3)
	}
	return 0
}

func String2Int32(__obf_937cd323c9259c13 string) int32 {
	__obf_97fc9da0ff9422b3, __obf_224415a75b186177 := strconv.ParseInt(__obf_937cd323c9259c13, 10, 32)
	if __obf_224415a75b186177 == nil {
		return int32(__obf_97fc9da0ff9422b3)
	}
	return 0
}

func String2Int64(__obf_937cd323c9259c13 string) int8 {
	__obf_97fc9da0ff9422b3, __obf_224415a75b186177 := strconv.ParseInt(__obf_937cd323c9259c13, 10, 64)
	if __obf_224415a75b186177 == nil {
		return int8(__obf_97fc9da0ff9422b3)
	}
	return 0
}

// OrdinalWeekPeriod 获取某年某周的起始日期
func OrdinalWeekPeriod(__obf_a8f7d5ed3e036323, __obf_758f44425645ced1 int) (__obf_7d64ca592a1d7345, __obf_00f9377132be8963 time.Time) {
	if __obf_a8f7d5ed3e036323 > 0 && __obf_758f44425645ced1 > 0 {
		__obf_7d64ca592a1d7345 = time.Date(__obf_a8f7d5ed3e036323, 1, 0, 0, 0, 0, 0, Loc)
		// 第一天是周几
		__obf_e4b0c86c00202b7d := int(__obf_7d64ca592a1d7345.AddDate(0, 0, 1).Weekday())
		// 当年第一周有几天
		__obf_edcbfc71c5dc013a := 1
		if __obf_e4b0c86c00202b7d != 0 {
			__obf_edcbfc71c5dc013a = 7 - __obf_e4b0c86c00202b7d + 1
		}
		if __obf_758f44425645ced1 == 1 {
			__obf_00f9377132be8963 = __obf_7d64ca592a1d7345.AddDate(0, 0, __obf_edcbfc71c5dc013a)
		} else {
			__obf_00f9377132be8963 = __obf_7d64ca592a1d7345.AddDate(0, 0, __obf_edcbfc71c5dc013a+(__obf_758f44425645ced1-1)*7)
			__obf_7d64ca592a1d7345 = __obf_00f9377132be8963.AddDate(0, 0, -7)
		}
	}
	return
}

// OrdinalWeek 当前日期为该年第几周周几，星期天为0
func OrdinalWeek(__obf_0a2edfa7a86ee54e time.Time) (__obf_a8f7d5ed3e036323, __obf_758f44425645ced1, __obf_a4b6e41aefcd1bb1 int) {
	__obf_a8f7d5ed3e036323 = __obf_0a2edfa7a86ee54e.Year()
	__obf_a4b6e41aefcd1bb1 = int(__obf_0a2edfa7a86ee54e.Weekday())
	__obf_21da77af459ec992 := __obf_0a2edfa7a86ee54e.YearDay()
	__obf_1271dd7ffec4684a := __obf_0a2edfa7a86ee54e.AddDate(0, 0, -__obf_21da77af459ec992+1)
	__obf_e4b0c86c00202b7d := int(__obf_1271dd7ffec4684a.Weekday())
	// 当年第一周有几天
	__obf_edcbfc71c5dc013a := 1
	if __obf_e4b0c86c00202b7d != 0 {
		__obf_edcbfc71c5dc013a = 7 - __obf_e4b0c86c00202b7d + 1
	}
	if __obf_21da77af459ec992 <= __obf_edcbfc71c5dc013a {
		__obf_758f44425645ced1 = 1
	} else {
		__obf_758f44425645ced1 = (__obf_21da77af459ec992-__obf_edcbfc71c5dc013a)/7 + 2
	}
	return
}

func HJ212RegExtract(__obf_c8d1b27c4126c87a []byte) map[string]string {
	__obf_c71fa0dae2d6b7ee := regexp.MustCompile(`CP=&&(.+)&&`)
	__obf_3be1e5b02c4c59f1 := regexp.MustCompile(`([^=]+)=([^;]+);?`)

	var __obf_713b20de5431218f []byte
	if __obf_713b20de5431218f = __obf_c71fa0dae2d6b7ee.ReplaceAll(__obf_c8d1b27c4126c87a, []byte("")); len(__obf_713b20de5431218f) < 6 {
		return nil
	}

	__obf_e3d74a0be2f4b93c := __obf_3be1e5b02c4c59f1.FindAllSubmatch(__obf_713b20de5431218f[6:len(__obf_713b20de5431218f)-7], -1)
	__obf_62a1f9b412c2cf69 := map[string]string{}
	if __obf_2a4d06608733608b := __obf_c71fa0dae2d6b7ee.FindSubmatch(__obf_c8d1b27c4126c87a)[1]; len(__obf_2a4d06608733608b) != 0 {
		__obf_46c26a6ecd59a116 := regexp.MustCompile(`([^=]+)=([^;,]+)[;,]?`).FindAllSubmatch(__obf_2a4d06608733608b, -1)
		for _, __obf_1c4edfc2acf87521 := range __obf_46c26a6ecd59a116 {
			__obf_62a1f9b412c2cf69[string(__obf_1c4edfc2acf87521[1])] = string(__obf_1c4edfc2acf87521[2])
		}
	}

	for _, __obf_1c4edfc2acf87521 := range __obf_e3d74a0be2f4b93c {
		__obf_62a1f9b412c2cf69[string(__obf_1c4edfc2acf87521[1])] = string(__obf_1c4edfc2acf87521[2])
	}
	return __obf_62a1f9b412c2cf69
}

// GetDiffTen 获取某一天00:00:00至23:59:59，以10min为单位的时间戳起止序列
func GetDiffTen(__obf_97a23963c801b013 time.Time) (int64, int64) {
	__obf_b6be0d242c1a676c := GetZeroTime(__obf_97a23963c801b013).Unix() / 600
	return __obf_b6be0d242c1a676c, __obf_b6be0d242c1a676c + 143
}

func Abs(__obf_c705c49030e00b6f int64) int64 {
	__obf_5af83d31b7fcc5d6 := __obf_c705c49030e00b6f >> 63
	return (__obf_c705c49030e00b6f ^ __obf_5af83d31b7fcc5d6) - __obf_5af83d31b7fcc5d6
}

func Number2String(__obf_c705c49030e00b6f any) string {
	// fmt.Println(reflect.TypeOf(n))
	switch __obf_c705c49030e00b6f := __obf_c705c49030e00b6f.(type) {
	case int:
		return strconv.Itoa(__obf_c705c49030e00b6f)
	case int32:
		return strconv.FormatInt(int64(__obf_c705c49030e00b6f), 10)
	case int64:
		return strconv.FormatInt(__obf_c705c49030e00b6f, 10)
	case float32:
		//return fmt.Sprint(n.(float32))
		return strconv.FormatFloat(float64(__obf_c705c49030e00b6f), 'f', -1, 32)
	case float64:
		//return fmt.Sprint(n.(float64))
		return strconv.FormatFloat(__obf_c705c49030e00b6f, 'f', -1, 64)
	default:
		return ""
	}
}

func IfNull(__obf_937cd323c9259c13 any, __obf_53aa887d3a1a5d3b string) string {
	if __obf_937cd323c9259c13 != nil && __obf_937cd323c9259c13.(string) != "" {
		// return sep + str + sep
		return __obf_937cd323c9259c13.(string)
	}
	return __obf_53aa887d3a1a5d3b
}

func SortRange(__obf_1bc902c81fe07f6f map[string]any, __obf_f5af2cf557c7f764 func(int, string)) {
	var __obf_f5f5a2b2454c1415 []string
	for __obf_1be1c6a485485b05 := range __obf_1bc902c81fe07f6f {
		__obf_f5f5a2b2454c1415 = append(__obf_f5f5a2b2454c1415, __obf_1be1c6a485485b05)
	}
	sort.Strings(__obf_f5f5a2b2454c1415)
	for __obf_ff7abde39552bac6, __obf_1be1c6a485485b05 := range __obf_f5f5a2b2454c1415 {
		__obf_f5af2cf557c7f764(__obf_ff7abde39552bac6, __obf_1be1c6a485485b05)
	}
}

func HasField(__obf_5f69b2844b820706 reflect.Value, __obf_1e523056a8de9a9c string) bool {

	if __obf_2b69c849afdb21c3 := __obf_5f69b2844b820706.FieldByNameFunc(func(__obf_c705c49030e00b6f string) bool {
		return strings.EqualFold(__obf_c705c49030e00b6f, __obf_1e523056a8de9a9c)
	}); __obf_2b69c849afdb21c3.IsValid() {
		return true
	}
	return false
}

func FieldByName(__obf_5f69b2844b820706 reflect.Value, __obf_1e523056a8de9a9c string) reflect.Value {
	if __obf_2b69c849afdb21c3 := __obf_5f69b2844b820706.FieldByNameFunc(func(__obf_c705c49030e00b6f string) bool {
		return strings.EqualFold(__obf_c705c49030e00b6f, __obf_1e523056a8de9a9c)
	}); __obf_2b69c849afdb21c3.IsValid() {
		return __obf_2b69c849afdb21c3
	}
	return reflect.Value{}
}

func SetFieldByName(__obf_5f69b2844b820706 reflect.Value, __obf_1e523056a8de9a9c string, __obf_cd007b9a47849e9a any) bool {

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
	if __obf_2b69c849afdb21c3 := __obf_5f69b2844b820706.FieldByNameFunc(func(__obf_c705c49030e00b6f string) bool {
		return strings.EqualFold(__obf_c705c49030e00b6f, __obf_1e523056a8de9a9c)
	}); __obf_2b69c849afdb21c3.IsValid() {
		if __obf_8f3fb0622e94d4b3 := __obf_2b69c849afdb21c3.Type(); __obf_8f3fb0622e94d4b3 == reflect.TypeOf(__obf_cd007b9a47849e9a) {
			__obf_2b69c849afdb21c3.Set(reflect.ValueOf(__obf_cd007b9a47849e9a))
		} else {
			__obf_2b69c849afdb21c3.Set(reflect.ValueOf(__obf_cd007b9a47849e9a).Convert(__obf_8f3fb0622e94d4b3))
		}
		return true
	}
	return false
}

func CopyMap(__obf_1bc902c81fe07f6f map[string]any) map[string]any {
	__obf_46c26a6ecd59a116 := make(map[string]any)
	for __obf_1be1c6a485485b05, __obf_1c4edfc2acf87521 := range __obf_1bc902c81fe07f6f {
		if __obf_4b41674a3353d967, __obf_481d3c2dc6d82397 := __obf_1c4edfc2acf87521.(map[string]any); __obf_481d3c2dc6d82397 {
			__obf_46c26a6ecd59a116[__obf_1be1c6a485485b05] = CopyMap(__obf_4b41674a3353d967)
		} else {
			__obf_46c26a6ecd59a116[__obf_1be1c6a485485b05] = __obf_1c4edfc2acf87521
		}
	}

	return __obf_46c26a6ecd59a116
}

// key: tag中的键值，delimiter: tag中的分隔符，structs：要生成表的结构体
func StructToTable(__obf_0af462578a0d3820, __obf_bd6ee9841979bf4e string, __obf_aa6bcf3410c5469d bool, __obf_0f669fd88fbb534c ...any) {
	var __obf_12444a4c09585469 strings.Builder
	var __obf_087dfca7b7bf4c83 string
	var __obf_8f2f13bf82dc08b1 []string
	for _, __obf_49e3bd562308b16c := range __obf_0f669fd88fbb534c {
		__obf_2b69c849afdb21c3 := reflect.TypeOf(__obf_49e3bd562308b16c)
		__obf_12444a4c09585469.WriteString(`CREATE TABLE `)
		__obf_12444a4c09585469.WriteString(__obf_2b69c849afdb21c3.Name())
		__obf_12444a4c09585469.WriteString(" (\n")
		__obf_874a7131aaa226ba := __obf_2b69c849afdb21c3.NumField()
		for __obf_ff7abde39552bac6 := 0; __obf_ff7abde39552bac6 < __obf_874a7131aaa226ba; __obf_ff7abde39552bac6++ {
			__obf_12444a4c09585469.WriteString("    ")
			__obf_8f2f13bf82dc08b1 = nil
			if __obf_72e0b94128874712 := string(__obf_2b69c849afdb21c3.Field(__obf_ff7abde39552bac6).Tag.Get(__obf_0af462578a0d3820)); __obf_72e0b94128874712 == "" {
				if __obf_aa6bcf3410c5469d {
					__obf_087dfca7b7bf4c83 = ToCamel(__obf_2b69c849afdb21c3.Field(__obf_ff7abde39552bac6).Name)
				} else {
					__obf_087dfca7b7bf4c83 = __obf_2b69c849afdb21c3.Field(__obf_ff7abde39552bac6).Name
				}
			} else {
				__obf_8f2f13bf82dc08b1 = strings.Split(__obf_72e0b94128874712, __obf_bd6ee9841979bf4e)
				__obf_087dfca7b7bf4c83 = __obf_8f2f13bf82dc08b1[0]
			}
			__obf_12444a4c09585469.WriteString(__obf_087dfca7b7bf4c83)
			__obf_12444a4c09585469.WriteString(" ")
			switch __obf_2b69c849afdb21c3.Field(__obf_ff7abde39552bac6).Type.Name() {
			case "int8":
				__obf_12444a4c09585469.WriteString("TINYINT")
			case "int", "int16", "int32":
				__obf_12444a4c09585469.WriteString("INT")
			case "int64":
				__obf_12444a4c09585469.WriteString("BIGINT")
			case "string":
				fallthrough
			default:
				__obf_12444a4c09585469.WriteString("VARCHAR(50)")
			}

			if len(__obf_8f2f13bf82dc08b1) > 1 {
				__obf_12444a4c09585469.WriteString(" ")
				__obf_12444a4c09585469.WriteString(strings.Join(__obf_8f2f13bf82dc08b1[1:], " "))
			}

			if __obf_ff7abde39552bac6+1 != __obf_874a7131aaa226ba {
				__obf_12444a4c09585469.WriteString(",")
			}
			__obf_12444a4c09585469.WriteString("\n")
		}
		__obf_12444a4c09585469.WriteString(");\n\n")
	}
	fmt.Println(__obf_12444a4c09585469.String())
}

func SaveImage(__obf_2698f45ab60c9a81 image.Image, __obf_20f803c1f7864c94 uint, __obf_9bb393ef5a0ca4c9 string) error {
	__obf_a2b808c1e5ef30e8, __obf_224415a75b186177 := os.Create(__obf_9bb393ef5a0ca4c9)
	if __obf_224415a75b186177 != nil {
		return __obf_224415a75b186177
	}
	defer __obf_a2b808c1e5ef30e8.Close()
	return jpeg.Encode(__obf_a2b808c1e5ef30e8, resize.Resize(__obf_20f803c1f7864c94, 0, __obf_2698f45ab60c9a81, resize.Bilinear), nil)
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(__obf_937cd323c9259c13 string) bool {
	return strings.Contains(__obf_937cd323c9259c13, ".") && net.ParseIP(__obf_937cd323c9259c13) != nil
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(__obf_937cd323c9259c13 string) bool {
	return strings.Contains(__obf_937cd323c9259c13, ":") && net.ParseIP(__obf_937cd323c9259c13) != nil
}

func AnyToBytes(__obf_1c4edfc2acf87521 any) ([]byte, error) {
	return msgpack.Marshal(__obf_1c4edfc2acf87521)
}

// 返回 any 该为某个结构体的指针类型
// func bytesToAny(val []byte, data any) (err error) {
// 	buf := new(bytes.Buffer)
// 	dec := gob.NewDecoder(buf)
// 	err = dec.Decode(data)
// 	return
// }

func BytesToAny[T any](__obf_cd007b9a47849e9a []byte) (__obf_f8bbf5dce647ada9 T, __obf_224415a75b186177 error) {
	__obf_224415a75b186177 = msgpack.Unmarshal(__obf_cd007b9a47849e9a, &__obf_f8bbf5dce647ada9)
	return
}

func Loadyaml(__obf_94e2fb9ca306266c string, __obf_7fe107eb1b6afbc7 any) {
	__obf_19ee06338066807d, __obf_224415a75b186177 := os.ReadFile(__obf_94e2fb9ca306266c)
	if __obf_224415a75b186177 != nil {
		log.Fatalln(__obf_224415a75b186177)
	}
	__obf_224415a75b186177 = yaml.UnmarshalStrict(__obf_19ee06338066807d, __obf_7fe107eb1b6afbc7)
	if __obf_224415a75b186177 != nil {
		log.Fatalln(__obf_224415a75b186177)
	}
}

func ToAnyList[T any](__obf_4697a31f998a0b74 []T) []any {
	__obf_df7a3e1932b09833 := make([]any, len(__obf_4697a31f998a0b74))
	for __obf_ff7abde39552bac6, __obf_1c4edfc2acf87521 := range __obf_4697a31f998a0b74 {
		__obf_df7a3e1932b09833[__obf_ff7abde39552bac6] = __obf_1c4edfc2acf87521
	}
	return __obf_df7a3e1932b09833
}

func TimeParse(__obf_fde675a93c2a2010 string) time.Time {
	__obf_fde675a93c2a2010 = fmt.Sprintf("%s %s", time.Now().Format(time.DateOnly), __obf_fde675a93c2a2010)
	__obf_f439046764df8d86, _ := time.ParseInLocation("2006-01-02 15:04", __obf_fde675a93c2a2010, time.Local)
	return __obf_f439046764df8d86
}
