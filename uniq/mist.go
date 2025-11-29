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

package __obf_07f0876faa0cf68e

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_4308c69bb16c764b = uint(8)                                         // 随机因子二进制位数
const __obf_526ea20e1484e8e8 = uint(8)                                         // 随机因子移位数
const __obf_def3b06d49a67cfe = __obf_4308c69bb16c764b + __obf_526ea20e1484e8e8 // 自增数移位数

type Mist struct {
	sync.Mutex
	__obf_49dd365a2ad62f67 int64 // 互斥锁
	__obf_b9e62152cda8d0bf int64 // 自增数
	__obf_640d1128ca7eac31 int64 // 随机因子一
	// 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_49dd365a2ad62f67 int64) *Mist {
	__obf_9b69a355e8a81435 := Mist{__obf_49dd365a2ad62f67: __obf_49dd365a2ad62f67}
	return &__obf_9b69a355e8a81435
}

/* 生成唯一编号 */
func (__obf_4a4db1b60f8d9fdf *Mist) Generate() int64 {
	__obf_4a4db1b60f8d9fdf.
		Lock()
	__obf_4a4db1b60f8d9fdf.

		// 获取随机因子数值 ｜ 使用真随机函数提高性能
		__obf_49dd365a2ad62f67++
	__obf_de39ecf25f737bdd, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_4a4db1b60f8d9fdf.__obf_b9e62152cda8d0bf = __obf_de39ecf25f737bdd.Int64()
	__obf_f8418a74f5331029, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_4a4db1b60f8d9fdf.__obf_640d1128ca7eac31 = __obf_f8418a74f5331029.Int64()
	__obf_9b69a355e8a81435 := // 通过位运算实现自动占位
		int64((__obf_4a4db1b60f8d9fdf.__obf_49dd365a2ad62f67 << __obf_def3b06d49a67cfe) | (__obf_4a4db1b60f8d9fdf.__obf_b9e62152cda8d0bf << __obf_526ea20e1484e8e8) | __obf_4a4db1b60f8d9fdf.__obf_640d1128ca7eac31)
	__obf_4a4db1b60f8d9fdf.
		Unlock()
	return __obf_9b69a355e8a81435
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
