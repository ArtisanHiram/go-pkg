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

package __obf_7d8ac56be2e11a40

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_24b16b9e86e8e024 = uint(8)                                         // 随机因子二进制位数
const __obf_64ec0e2de9ffaa7e = uint(8)                                         // 随机因子移位数
const __obf_90e3d424190d773a = __obf_24b16b9e86e8e024 + __obf_64ec0e2de9ffaa7e // 自增数移位数

type Mist struct {
	sync.Mutex                   // 互斥锁
	__obf_463dd678f4e81058 int64 // 自增数
	__obf_dc57cc895815ae92 int64 // 随机因子一
	__obf_2895a03d50944a35 int64 // 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_463dd678f4e81058 int64) *Mist {
	__obf_2a9e5a6fc49363db := Mist{__obf_463dd678f4e81058: __obf_463dd678f4e81058}
	return &__obf_2a9e5a6fc49363db
}

/* 生成唯一编号 */
func (__obf_1d6c073ee1ba2a2b *Mist) Generate() int64 {
	__obf_1d6c073ee1ba2a2b.Lock()
	__obf_1d6c073ee1ba2a2b.__obf_463dd678f4e81058++
	// 获取随机因子数值 ｜ 使用真随机函数提高性能
	__obf_27ef4ba288b97d95, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_1d6c073ee1ba2a2b.__obf_dc57cc895815ae92 = __obf_27ef4ba288b97d95.Int64()
	__obf_6cfe9c217144f182, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_1d6c073ee1ba2a2b.__obf_2895a03d50944a35 = __obf_6cfe9c217144f182.Int64()
	// 通过位运算实现自动占位
	__obf_2a9e5a6fc49363db := int64((__obf_1d6c073ee1ba2a2b.__obf_463dd678f4e81058 << __obf_90e3d424190d773a) | (__obf_1d6c073ee1ba2a2b.__obf_dc57cc895815ae92 << __obf_64ec0e2de9ffaa7e) | __obf_1d6c073ee1ba2a2b.__obf_2895a03d50944a35)
	__obf_1d6c073ee1ba2a2b.Unlock()
	return __obf_2a9e5a6fc49363db
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
