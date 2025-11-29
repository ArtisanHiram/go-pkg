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

package __obf_2f51f7d26a2bcdf8

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_3a1953b683ad0773 = uint(8)                                         // 随机因子二进制位数
const __obf_3611283a7c1f54e0 = uint(8)                                         // 随机因子移位数
const __obf_9f1a4d3069d8c622 = __obf_3a1953b683ad0773 + __obf_3611283a7c1f54e0 // 自增数移位数

type Mist struct {
	sync.Mutex
	__obf_77b66b29d882350d int64 // 互斥锁
	__obf_bc700a4880de0308 int64 // 自增数
	__obf_217043232bb70beb int64 // 随机因子一
	// 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_77b66b29d882350d int64) *Mist {
	__obf_1682710a7974862c := Mist{__obf_77b66b29d882350d: __obf_77b66b29d882350d}
	return &__obf_1682710a7974862c
}

/* 生成唯一编号 */
func (__obf_3570df32dc4aa0c0 *Mist) Generate() int64 {
	__obf_3570df32dc4aa0c0.
		Lock()
	__obf_3570df32dc4aa0c0.

		// 获取随机因子数值 ｜ 使用真随机函数提高性能
		__obf_77b66b29d882350d++
	__obf_5d1a2a52f2d85b29, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_3570df32dc4aa0c0.__obf_bc700a4880de0308 = __obf_5d1a2a52f2d85b29.Int64()
	__obf_e7c0950e717f2c91, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_3570df32dc4aa0c0.__obf_217043232bb70beb = __obf_e7c0950e717f2c91.Int64()
	__obf_1682710a7974862c := // 通过位运算实现自动占位
		int64((__obf_3570df32dc4aa0c0.__obf_77b66b29d882350d << __obf_9f1a4d3069d8c622) | (__obf_3570df32dc4aa0c0.__obf_bc700a4880de0308 << __obf_3611283a7c1f54e0) | __obf_3570df32dc4aa0c0.__obf_217043232bb70beb)
	__obf_3570df32dc4aa0c0.
		Unlock()
	return __obf_1682710a7974862c
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
