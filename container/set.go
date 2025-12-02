// 非线程安全set集合
package __obf_4a16ef421fb74992

import "fmt"

type Set[T comparable] map[T]__obf_b6fe4aa13f2799ab
type __obf_b6fe4aa13f2799ab struct{}

func NewSet[E comparable](__obf_9efe90729745e8ee []E) *Set[E] {
	__obf_0c76e51a510deb3d :=// 获取Set的地址
	Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_b56f3324ed629e96 := range __obf_9efe90729745e8ee {
		__obf_0c76e51a510deb3d[// unsafe.Sizeof(void{}) // 结果为0
		__obf_b56f3324ed629e96] = __obf_b6fe4aa13f2799ab{}
	}
	return &__obf_0c76e51a510deb3d
}

func (__obf_0c76e51a510deb3d *Set[T]) Add(__obf_79cf1c44489eaf01 T) bool {
	if _, __obf_0a7b144ecfd706f0 := (*__obf_0c76e51a510deb3d)[__obf_79cf1c44489eaf01]; __obf_0a7b144ecfd706f0 {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_0c76e51a510deb3d)[__obf_79cf1c44489eaf01] = __obf_b6fe4aa13f2799ab{}
	return true
}

func (__obf_0c76e51a510deb3d *Set[T]) Contains(__obf_b56f3324ed629e96 T) bool {
	_, __obf_0a7b144ecfd706f0 := (*__obf_0c76e51a510deb3d)[__obf_b56f3324ed629e96]
	return __obf_0a7b144ecfd706f0
}

func (__obf_0c76e51a510deb3d *Set[T]) Size() int {
	return len(*__obf_0c76e51a510deb3d)
}

func (__obf_0c76e51a510deb3d *Set[T]) Elem() []T {
	__obf_e937fcb7c91671b3 := make([]T, 0, len(*__obf_0c76e51a510deb3d))
	for __obf_91fb3c93595a08ac := range *__obf_0c76e51a510deb3d {
		__obf_e937fcb7c91671b3 = append(__obf_e937fcb7c91671b3, __obf_91fb3c93595a08ac)
	}
	return __obf_e937fcb7c91671b3
}

func (__obf_0c76e51a510deb3d *Set[T]) Remove(__obf_79cf1c44489eaf01 T) {
	delete(*__obf_0c76e51a510deb3d, __obf_79cf1c44489eaf01)
}

func (__obf_0c76e51a510deb3d *Set[T]) Clear() {
	*__obf_0c76e51a510deb3d = make(map[T]__obf_b6fe4aa13f2799ab)
}

// 相等
func (__obf_0c76e51a510deb3d *Set[T]) Equal(__obf_5753b9555f5a94ca *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_0c76e51a510deb3d.Size() != __obf_5753b9555f5a94ca.Size() {
		return false
	}
	for __obf_ef0e2650e203f28e := range *__obf_0c76e51a510deb3d {
		if !__obf_5753b9555f5a94ca.Contains(__obf_ef0e2650e203f28e) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_0c76e51a510deb3d *Set[T]) IsSubset(__obf_5753b9555f5a94ca *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_0c76e51a510deb3d.Size() > __obf_5753b9555f5a94ca.Size() {
		return false
	}
	// 迭代遍历
	for __obf_d1a1f98ae0fb119e := range *__obf_0c76e51a510deb3d {
		if !__obf_5753b9555f5a94ca.Contains(__obf_d1a1f98ae0fb119e) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_a4a6b352f61a966b *Set[T]) Union(__obf_5753b9555f5a94ca *Set[T]) *Set[T] {
	__obf_4b393d5800b36870 := Set[T]{}
	for __obf_ef0e2650e203f28e := range *__obf_a4a6b352f61a966b {
		__obf_4b393d5800b36870.
			Add(__obf_ef0e2650e203f28e)
	}
	for __obf_ef0e2650e203f28e := range *__obf_5753b9555f5a94ca {
		__obf_4b393d5800b36870.
			Add(__obf_ef0e2650e203f28e)
	}
	return &__obf_4b393d5800b36870
}

// 交集
func (__obf_0c76e51a510deb3d *Set[T]) Intersect(__obf_5753b9555f5a94ca *Set[T]) *Set[T] {
	__obf_bacd7beb0c75888a := Set[T]{}
	// loop over smaller set
	if __obf_0c76e51a510deb3d.Size() < __obf_5753b9555f5a94ca.Size() {
		for __obf_ef0e2650e203f28e := range *__obf_0c76e51a510deb3d {
			if __obf_5753b9555f5a94ca.Contains(__obf_ef0e2650e203f28e) {
				__obf_bacd7beb0c75888a.
					Add(__obf_ef0e2650e203f28e)
			}
		}
	} else {
		for __obf_ef0e2650e203f28e := range *__obf_5753b9555f5a94ca {
			if __obf_0c76e51a510deb3d.Contains(__obf_ef0e2650e203f28e) {
				__obf_bacd7beb0c75888a.
					Add(__obf_ef0e2650e203f28e)
			}
		}
	}
	return &__obf_bacd7beb0c75888a
}

// 差集
func (__obf_0c76e51a510deb3d *Set[T]) Difference(__obf_5753b9555f5a94ca *Set[T]) *Set[T] {
	__obf_af1d4f8de96ed6e6 := Set[T]{}
	for __obf_ef0e2650e203f28e := range *__obf_0c76e51a510deb3d {
		if !__obf_5753b9555f5a94ca.Contains(__obf_ef0e2650e203f28e) {
			__obf_af1d4f8de96ed6e6.
				Add(__obf_ef0e2650e203f28e)
		}
	}
	return &__obf_af1d4f8de96ed6e6
}

// 函数遍历
func (__obf_0c76e51a510deb3d *Set[T]) Each(__obf_82fb80bca9432eb9 func(any) bool) {
	for __obf_ef0e2650e203f28e := range *__obf_0c76e51a510deb3d {
		if __obf_82fb80bca9432eb9(__obf_ef0e2650e203f28e) {
			break
		}
	}
}

// 返回string数组
func (__obf_0c76e51a510deb3d *Set[T]) StringElem() []string {
	__obf_9efe90729745e8ee := make([]string, 0, len(*__obf_0c76e51a510deb3d))

	for __obf_ef0e2650e203f28e := range *__obf_0c76e51a510deb3d {
		__obf_9efe90729745e8ee = append(__obf_9efe90729745e8ee, fmt.Sprintf("%v", __obf_ef0e2650e203f28e))
	}
	return __obf_9efe90729745e8ee
}
