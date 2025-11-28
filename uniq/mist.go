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

package __obf_0f0e0d1ff72f3ff0

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_e62d5758b88003c1 = uint(8)                                         // 随机因子二进制位数
const __obf_1245a67c64e25af6 = uint(8)                                         // 随机因子移位数
const __obf_33c9fa92c99131da = __obf_e62d5758b88003c1 + __obf_1245a67c64e25af6 // 自增数移位数

type Mist struct {
	sync.Mutex                   // 互斥锁
	__obf_f662c63ba0ecfb47 int64 // 自增数
	__obf_3ba2504b5de91909 int64 // 随机因子一
	__obf_a390e2bba4729246 int64 // 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_f662c63ba0ecfb47 int64) *Mist {
	__obf_045126a5d8c6d66c := Mist{__obf_f662c63ba0ecfb47: __obf_f662c63ba0ecfb47}
	return &__obf_045126a5d8c6d66c
}

/* 生成唯一编号 */
func (__obf_9a4001ed8b724993 *Mist) Generate() int64 {
	__obf_9a4001ed8b724993.Lock()
	__obf_9a4001ed8b724993.__obf_f662c63ba0ecfb47++
	// 获取随机因子数值 ｜ 使用真随机函数提高性能
	__obf_226a82f0fc08409c, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_9a4001ed8b724993.__obf_3ba2504b5de91909 = __obf_226a82f0fc08409c.Int64()
	__obf_a95970eb52484f2b, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_9a4001ed8b724993.__obf_a390e2bba4729246 = __obf_a95970eb52484f2b.Int64()
	// 通过位运算实现自动占位
	__obf_045126a5d8c6d66c := int64((__obf_9a4001ed8b724993.__obf_f662c63ba0ecfb47 << __obf_33c9fa92c99131da) | (__obf_9a4001ed8b724993.__obf_3ba2504b5de91909 << __obf_1245a67c64e25af6) | __obf_9a4001ed8b724993.__obf_a390e2bba4729246)
	__obf_9a4001ed8b724993.Unlock()
	return __obf_045126a5d8c6d66c
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
