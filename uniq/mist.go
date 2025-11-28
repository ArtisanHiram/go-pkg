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

package __obf_e2239f4853c61591

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_7e65d5b564449dec = uint(8)                                         // 随机因子二进制位数
const __obf_6557737c899102c3 = uint(8)                                         // 随机因子移位数
const __obf_b656c71d1e9689b2 = __obf_7e65d5b564449dec + __obf_6557737c899102c3 // 自增数移位数

type Mist struct {
	sync.Mutex                   // 互斥锁
	__obf_1d593ab487e100dc int64 // 自增数
	__obf_fb9b483135e83ec5 int64 // 随机因子一
	__obf_5d3c50f9e8a89895 int64 // 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_1d593ab487e100dc int64) *Mist {
	__obf_fc65aba223ce84d3 := Mist{__obf_1d593ab487e100dc: __obf_1d593ab487e100dc}
	return &__obf_fc65aba223ce84d3
}

/* 生成唯一编号 */
func (__obf_ca9a2d6567626676 *Mist) Generate() int64 {
	__obf_ca9a2d6567626676.Lock()
	__obf_ca9a2d6567626676.__obf_1d593ab487e100dc++
	// 获取随机因子数值 ｜ 使用真随机函数提高性能
	__obf_5dbf190dac9c8818, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_ca9a2d6567626676.__obf_fb9b483135e83ec5 = __obf_5dbf190dac9c8818.Int64()
	__obf_5cfd922e7c1f9385, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_ca9a2d6567626676.__obf_5d3c50f9e8a89895 = __obf_5cfd922e7c1f9385.Int64()
	// 通过位运算实现自动占位
	__obf_fc65aba223ce84d3 := int64((__obf_ca9a2d6567626676.__obf_1d593ab487e100dc << __obf_b656c71d1e9689b2) | (__obf_ca9a2d6567626676.__obf_fb9b483135e83ec5 << __obf_6557737c899102c3) | __obf_ca9a2d6567626676.__obf_5d3c50f9e8a89895)
	__obf_ca9a2d6567626676.Unlock()
	return __obf_fc65aba223ce84d3
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
