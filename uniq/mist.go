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

package __obf_34ce7ee87a5aa6b7

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_dd022c6af3ef1400 = uint(8)                                         // 随机因子二进制位数
const __obf_9ebd0ca46fa5ef74 = uint(8)                                         // 随机因子移位数
const __obf_0ad3ef4c2f5adc6d = __obf_dd022c6af3ef1400 + __obf_9ebd0ca46fa5ef74 // 自增数移位数

type Mist struct {
	sync.Mutex
	__obf_7eb1cfd52b425b32 int64 // 互斥锁
	__obf_3862b136142cf80c int64 // 自增数
	__obf_0b4db9eb87354ff8 int64 // 随机因子一
	// 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_7eb1cfd52b425b32 int64) *Mist {
	__obf_0a77b1482a5f3b17 := Mist{__obf_7eb1cfd52b425b32: __obf_7eb1cfd52b425b32}
	return &__obf_0a77b1482a5f3b17
}

/* 生成唯一编号 */
func (__obf_ec84b9ba40ae72c7 *Mist) Generate() int64 {
	__obf_ec84b9ba40ae72c7.
		Lock()
	__obf_ec84b9ba40ae72c7.

		// 获取随机因子数值 ｜ 使用真随机函数提高性能
		__obf_7eb1cfd52b425b32++
	__obf_b7482ccc53f96a43, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_ec84b9ba40ae72c7.__obf_3862b136142cf80c = __obf_b7482ccc53f96a43.Int64()
	__obf_05d9a63be494b445, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_ec84b9ba40ae72c7.__obf_0b4db9eb87354ff8 = __obf_05d9a63be494b445.Int64()
	__obf_0a77b1482a5f3b17 := // 通过位运算实现自动占位
		int64((__obf_ec84b9ba40ae72c7.__obf_7eb1cfd52b425b32 << __obf_0ad3ef4c2f5adc6d) | (__obf_ec84b9ba40ae72c7.__obf_3862b136142cf80c << __obf_9ebd0ca46fa5ef74) | __obf_ec84b9ba40ae72c7.__obf_0b4db9eb87354ff8)
	__obf_ec84b9ba40ae72c7.
		Unlock()
	return __obf_0a77b1482a5f3b17
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
