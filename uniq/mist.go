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

package __obf_7913809aab6c8423

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_f49c0b7806b5df8c = uint(8)                                         // 随机因子二进制位数
const __obf_ba14d07f06ecd2e8 = uint(8)                                         // 随机因子移位数
const __obf_2b414c258ca5b6c5 = __obf_f49c0b7806b5df8c + __obf_ba14d07f06ecd2e8 // 自增数移位数

type Mist struct {
	sync.Mutex                   // 互斥锁
	__obf_df106466c6f3363d int64 // 自增数
	__obf_fb2b70906751d9c1 int64 // 随机因子一
	__obf_f2973be10b3afa9b int64 // 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_df106466c6f3363d int64) *Mist {
	__obf_c055351b29015730 := Mist{__obf_df106466c6f3363d: __obf_df106466c6f3363d}
	return &__obf_c055351b29015730
}

/* 生成唯一编号 */
func (__obf_fa0b1524a946e7a1 *Mist) Generate() int64 {
	__obf_fa0b1524a946e7a1.Lock()
	__obf_fa0b1524a946e7a1.__obf_df106466c6f3363d++
	// 获取随机因子数值 ｜ 使用真随机函数提高性能
	__obf_9eca525418099c33, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_fa0b1524a946e7a1.__obf_fb2b70906751d9c1 = __obf_9eca525418099c33.Int64()
	__obf_1f10fdc99d59962b, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_fa0b1524a946e7a1.__obf_f2973be10b3afa9b = __obf_1f10fdc99d59962b.Int64()
	// 通过位运算实现自动占位
	__obf_c055351b29015730 := int64((__obf_fa0b1524a946e7a1.__obf_df106466c6f3363d << __obf_2b414c258ca5b6c5) | (__obf_fa0b1524a946e7a1.__obf_fb2b70906751d9c1 << __obf_ba14d07f06ecd2e8) | __obf_fa0b1524a946e7a1.__obf_f2973be10b3afa9b)
	__obf_fa0b1524a946e7a1.Unlock()
	return __obf_c055351b29015730
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
