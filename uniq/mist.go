/*
* 薄雾算法
*
* 1      2                                                     48         56       64
* +------+-----------------------------------------------------+----------+----------+
* retain | increas                                             | salt     | sequence |
* +------+-----------------------------------------------------+----------+----------+
* 0      | 0000000000 0000000000 0000000000 0000000000 0000000 | 00000000 | 00000000 |
* +------+-----------------------------------------------------+------------+--------+
*
* 0. 最高位，占 1 位，保持为 0，使得值永远为正数；
* 1. 自增数，占 47 位，自增数在高位能保证结果值呈递增态势，遂低位可以为所欲为；
* 2. 随机因子一，占 8 位，上限数值 255，使结果值不可预测；
* 3. 随机因子二，占 8 位，上限数值 255，使结果值不可预测；
*
* 编号上限为百万亿级，上限值计算为 140737488355327 即 int64(1 << 47 - 1)，假设每天取值 10 亿，能使用 385+ 年
 */

package __obf_3747a7e09ff475ee

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_26eaccd1f9c37fff = uint(8)                                         // 随机因子二进制位数
const __obf_4996829b7316fc10 = uint(8)                                         // 随机因子移位数
const __obf_a71ee0255499caef = __obf_26eaccd1f9c37fff + __obf_4996829b7316fc10 // 自增数移位数

type Mist struct {
	sync.Mutex
	__obf_da46e9ec79c837e3 int64 // 互斥锁
	__obf_1d724e7d5a8dc4bd int64 // 自增数
	__obf_bac241d769d99a1e int64 // 随机因子一
	// 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_da46e9ec79c837e3 int64) *Mist {
	__obf_3582d6ff5c97b868 := Mist{__obf_da46e9ec79c837e3: __obf_da46e9ec79c837e3}
	return &__obf_3582d6ff5c97b868
}

/* 生成唯一编号 */
func (__obf_4cd50e267441590a *Mist) Generate() int64 {
	__obf_4cd50e267441590a.
		Lock()
	__obf_4cd50e267441590a.

		// 获取随机因子数值 ｜ 使用真随机函数提高性能
		__obf_da46e9ec79c837e3++
	__obf_2e4d3954348b13c4, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_4cd50e267441590a.__obf_1d724e7d5a8dc4bd = __obf_2e4d3954348b13c4.Int64()
	__obf_d617771c09e9a1ca, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_4cd50e267441590a.__obf_bac241d769d99a1e = __obf_d617771c09e9a1ca.Int64()
	__obf_3582d6ff5c97b868 := // 通过位运算实现自动占位
		int64((__obf_4cd50e267441590a.__obf_da46e9ec79c837e3 << __obf_a71ee0255499caef) | (__obf_4cd50e267441590a.__obf_1d724e7d5a8dc4bd << __obf_4996829b7316fc10) | __obf_4cd50e267441590a.__obf_bac241d769d99a1e)
	__obf_4cd50e267441590a.
		Unlock()
	return __obf_3582d6ff5c97b868
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
