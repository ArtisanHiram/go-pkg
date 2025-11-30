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

package __obf_a51a64e21f311927

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_96487cda11ef49ee = uint(8)                                         // 随机因子二进制位数
const __obf_071bed805a6fbd07 = uint(8)                                         // 随机因子移位数
const __obf_02030bf8121e2ff4 = __obf_96487cda11ef49ee + __obf_071bed805a6fbd07 // 自增数移位数

type Mist struct {
	sync.Mutex
	__obf_5876679f517beca4 int64 // 互斥锁
	__obf_f8c2a310b39d9ec5 int64 // 自增数
	__obf_97921e76b1c1ea56 int64 // 随机因子一
	// 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_5876679f517beca4 int64) *Mist {
	__obf_5092acf7edc003bf := Mist{__obf_5876679f517beca4: __obf_5876679f517beca4}
	return &__obf_5092acf7edc003bf
}

/* 生成唯一编号 */
func (__obf_8931c216d973ece1 *Mist) Generate() int64 {
	__obf_8931c216d973ece1.
		Lock()
	__obf_8931c216d973ece1.

		// 获取随机因子数值 ｜ 使用真随机函数提高性能
		__obf_5876679f517beca4++
	__obf_d406a3b8d94cb2a1, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_8931c216d973ece1.__obf_f8c2a310b39d9ec5 = __obf_d406a3b8d94cb2a1.Int64()
	__obf_8256665a03878275, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_8931c216d973ece1.__obf_97921e76b1c1ea56 = __obf_8256665a03878275.Int64()
	__obf_5092acf7edc003bf := // 通过位运算实现自动占位
		int64((__obf_8931c216d973ece1.__obf_5876679f517beca4 << __obf_02030bf8121e2ff4) | (__obf_8931c216d973ece1.__obf_f8c2a310b39d9ec5 << __obf_071bed805a6fbd07) | __obf_8931c216d973ece1.__obf_97921e76b1c1ea56)
	__obf_8931c216d973ece1.
		Unlock()
	return __obf_5092acf7edc003bf
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
