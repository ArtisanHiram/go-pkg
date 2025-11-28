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

package __obf_8fe28252c1b01dfe

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_aea0874139c14d2f = uint(8)                                         // 随机因子二进制位数
const __obf_62eec86316328bb9 = uint(8)                                         // 随机因子移位数
const __obf_1cbc06981c0dfe39 = __obf_aea0874139c14d2f + __obf_62eec86316328bb9 // 自增数移位数

type Mist struct {
	sync.Mutex                   // 互斥锁
	__obf_b9499d2271541130 int64 // 自增数
	__obf_53b76bd9b299c2ac int64 // 随机因子一
	__obf_debf396c5ea99c4c int64 // 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_b9499d2271541130 int64) *Mist {
	__obf_0a3c52831183e6a4 := Mist{__obf_b9499d2271541130: __obf_b9499d2271541130}
	return &__obf_0a3c52831183e6a4
}

/* 生成唯一编号 */
func (__obf_1aa68404a6855ad9 *Mist) Generate() int64 {
	__obf_1aa68404a6855ad9.Lock()
	__obf_1aa68404a6855ad9.__obf_b9499d2271541130++
	// 获取随机因子数值 ｜ 使用真随机函数提高性能
	__obf_8fc983a450b8859a, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_1aa68404a6855ad9.__obf_53b76bd9b299c2ac = __obf_8fc983a450b8859a.Int64()
	__obf_28e268951765b770, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_1aa68404a6855ad9.__obf_debf396c5ea99c4c = __obf_28e268951765b770.Int64()
	// 通过位运算实现自动占位
	__obf_0a3c52831183e6a4 := int64((__obf_1aa68404a6855ad9.__obf_b9499d2271541130 << __obf_1cbc06981c0dfe39) | (__obf_1aa68404a6855ad9.__obf_53b76bd9b299c2ac << __obf_62eec86316328bb9) | __obf_1aa68404a6855ad9.__obf_debf396c5ea99c4c)
	__obf_1aa68404a6855ad9.Unlock()
	return __obf_0a3c52831183e6a4
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
