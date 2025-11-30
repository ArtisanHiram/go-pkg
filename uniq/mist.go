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

package __obf_4d7a51f8e2f0494a

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_44f1c3252c6f9a3b = uint(8)                                         // 随机因子二进制位数
const __obf_78c523df66b0c934 = uint(8)                                         // 随机因子移位数
const __obf_61a698afe8519d32 = __obf_44f1c3252c6f9a3b + __obf_78c523df66b0c934 // 自增数移位数

type Mist struct {
	sync.Mutex
	__obf_1ef4ecfcc0d15f78 int64 // 互斥锁
	__obf_57d66e4fcb715f10 int64 // 自增数
	__obf_914f917f16bd537e int64 // 随机因子一
	// 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_1ef4ecfcc0d15f78 int64) *Mist {
	__obf_f627d657a2040ab1 := Mist{__obf_1ef4ecfcc0d15f78: __obf_1ef4ecfcc0d15f78}
	return &__obf_f627d657a2040ab1
}

/* 生成唯一编号 */
func (__obf_934c4d4a8dea62d8 *Mist) Generate() int64 {
	__obf_934c4d4a8dea62d8.
		Lock()
	__obf_934c4d4a8dea62d8.

		// 获取随机因子数值 ｜ 使用真随机函数提高性能
		__obf_1ef4ecfcc0d15f78++
	__obf_21495aebdf3464a8, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_934c4d4a8dea62d8.__obf_57d66e4fcb715f10 = __obf_21495aebdf3464a8.Int64()
	__obf_190f1265db10986e, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_934c4d4a8dea62d8.__obf_914f917f16bd537e = __obf_190f1265db10986e.Int64()
	__obf_f627d657a2040ab1 := // 通过位运算实现自动占位
		int64((__obf_934c4d4a8dea62d8.__obf_1ef4ecfcc0d15f78 << __obf_61a698afe8519d32) | (__obf_934c4d4a8dea62d8.__obf_57d66e4fcb715f10 << __obf_78c523df66b0c934) | __obf_934c4d4a8dea62d8.__obf_914f917f16bd537e)
	__obf_934c4d4a8dea62d8.
		Unlock()
	return __obf_f627d657a2040ab1
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
