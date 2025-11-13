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

package __obf_417508f5ae215d0a

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const __obf_ab9bbc1a2fa61e59 = uint(8)                                         // 随机因子二进制位数
const __obf_db6050471bb435f4 = uint(8)                                         // 随机因子移位数
const __obf_c8c8683fc8fb7b2a = __obf_ab9bbc1a2fa61e59 + __obf_db6050471bb435f4 // 自增数移位数

type Mist struct {
	sync.Mutex                   // 互斥锁
	__obf_e0fe3e0c7d71a499 int64 // 自增数
	__obf_2fe60be9de1d9a69 int64 // 随机因子一
	__obf_5ec1e4c8af460406 int64 // 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist(__obf_e0fe3e0c7d71a499 int64) *Mist {
	__obf_5685540bcad65f9c := Mist{__obf_e0fe3e0c7d71a499: __obf_e0fe3e0c7d71a499}
	return &__obf_5685540bcad65f9c
}

/* 生成唯一编号 */
func (__obf_fbaf9c1c0dab70a7 *Mist) Generate() int64 {
	__obf_fbaf9c1c0dab70a7.Lock()
	__obf_fbaf9c1c0dab70a7.__obf_e0fe3e0c7d71a499++
	// 获取随机因子数值 ｜ 使用真随机函数提高性能
	__obf_a7c1679a1a7d3cbc, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_fbaf9c1c0dab70a7.__obf_2fe60be9de1d9a69 = __obf_a7c1679a1a7d3cbc.Int64()
	__obf_cd240e80aaa4cd80, _ := rand.Int(rand.Reader, big.NewInt(255))
	__obf_fbaf9c1c0dab70a7.__obf_5ec1e4c8af460406 = __obf_cd240e80aaa4cd80.Int64()
	// 通过位运算实现自动占位
	__obf_5685540bcad65f9c := int64((__obf_fbaf9c1c0dab70a7.__obf_e0fe3e0c7d71a499 << __obf_c8c8683fc8fb7b2a) | (__obf_fbaf9c1c0dab70a7.__obf_2fe60be9de1d9a69 << __obf_db6050471bb435f4) | __obf_fbaf9c1c0dab70a7.__obf_5ec1e4c8af460406)
	__obf_fbaf9c1c0dab70a7.Unlock()
	return __obf_5685540bcad65f9c
}

// func main() {
// 	// 使用方法
// 	mist := NewMist()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(mist.Generate())
// 	}
// }
