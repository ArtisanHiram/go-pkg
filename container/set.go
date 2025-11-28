// 非线程安全set集合
package __obf_1fda7fbdeda52f1e

import "fmt"

type Set[T comparable] map[T]__obf_0153603575849025
type __obf_0153603575849025 struct{}

func NewSet[E comparable](__obf_3d24c929f1a76b8b []E) *Set[E] {
	// 获取Set的地址
	__obf_7d38105b81b181cc := Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_86916786040a389c := range __obf_3d24c929f1a76b8b {
		// unsafe.Sizeof(void{}) // 结果为0
		__obf_7d38105b81b181cc[__obf_86916786040a389c] = __obf_0153603575849025{}
	}
	return &__obf_7d38105b81b181cc
}

func (__obf_7d38105b81b181cc *Set[T]) Add(__obf_457766c6d5154891 T) bool {
	if _, __obf_d24c61455fa2e4db := (*__obf_7d38105b81b181cc)[__obf_457766c6d5154891]; __obf_d24c61455fa2e4db {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_7d38105b81b181cc)[__obf_457766c6d5154891] = __obf_0153603575849025{}
	return true
}

func (__obf_7d38105b81b181cc *Set[T]) Contains(__obf_86916786040a389c T) bool {
	_, __obf_d24c61455fa2e4db := (*__obf_7d38105b81b181cc)[__obf_86916786040a389c]
	return __obf_d24c61455fa2e4db
}

func (__obf_7d38105b81b181cc *Set[T]) Size() int {
	return len(*__obf_7d38105b81b181cc)
}

func (__obf_7d38105b81b181cc *Set[T]) Elem() []T {
	__obf_f33bf750ed8ea535 := make([]T, 0, len(*__obf_7d38105b81b181cc))
	for __obf_9058efdba9a0277b := range *__obf_7d38105b81b181cc {
		__obf_f33bf750ed8ea535 = append(__obf_f33bf750ed8ea535, __obf_9058efdba9a0277b)
	}
	return __obf_f33bf750ed8ea535
}

func (__obf_7d38105b81b181cc *Set[T]) Remove(__obf_457766c6d5154891 T) {
	delete(*__obf_7d38105b81b181cc, __obf_457766c6d5154891)
}

func (__obf_7d38105b81b181cc *Set[T]) Clear() {
	*__obf_7d38105b81b181cc = make(map[T]__obf_0153603575849025)
}

// 相等
func (__obf_7d38105b81b181cc *Set[T]) Equal(__obf_94e6449148fc8624 *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_7d38105b81b181cc.Size() != __obf_94e6449148fc8624.Size() {
		return false
	}
	for __obf_5a12388a905e7c2e := range *__obf_7d38105b81b181cc {
		if !__obf_94e6449148fc8624.Contains(__obf_5a12388a905e7c2e) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_7d38105b81b181cc *Set[T]) IsSubset(__obf_94e6449148fc8624 *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_7d38105b81b181cc.Size() > __obf_94e6449148fc8624.Size() {
		return false
	}
	// 迭代遍历
	for __obf_95c4c09a56c1a070 := range *__obf_7d38105b81b181cc {
		if !__obf_94e6449148fc8624.Contains(__obf_95c4c09a56c1a070) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_b051f0cbd935a3ac *Set[T]) Union(__obf_94e6449148fc8624 *Set[T]) *Set[T] {
	__obf_f453da6aba100729 := Set[T]{}
	for __obf_5a12388a905e7c2e := range *__obf_b051f0cbd935a3ac {
		__obf_f453da6aba100729.Add(__obf_5a12388a905e7c2e)
	}
	for __obf_5a12388a905e7c2e := range *__obf_94e6449148fc8624 {
		__obf_f453da6aba100729.Add(__obf_5a12388a905e7c2e)
	}
	return &__obf_f453da6aba100729
}

// 交集
func (__obf_7d38105b81b181cc *Set[T]) Intersect(__obf_94e6449148fc8624 *Set[T]) *Set[T] {

	__obf_009adc3f49d583bb := Set[T]{}
	// loop over smaller set
	if __obf_7d38105b81b181cc.Size() < __obf_94e6449148fc8624.Size() {
		for __obf_5a12388a905e7c2e := range *__obf_7d38105b81b181cc {
			if __obf_94e6449148fc8624.Contains(__obf_5a12388a905e7c2e) {
				__obf_009adc3f49d583bb.Add(__obf_5a12388a905e7c2e)
			}
		}
	} else {
		for __obf_5a12388a905e7c2e := range *__obf_94e6449148fc8624 {
			if __obf_7d38105b81b181cc.Contains(__obf_5a12388a905e7c2e) {
				__obf_009adc3f49d583bb.Add(__obf_5a12388a905e7c2e)
			}
		}
	}
	return &__obf_009adc3f49d583bb
}

// 差集
func (__obf_7d38105b81b181cc *Set[T]) Difference(__obf_94e6449148fc8624 *Set[T]) *Set[T] {

	__obf_340cb9d18c761e30 := Set[T]{}
	for __obf_5a12388a905e7c2e := range *__obf_7d38105b81b181cc {
		if !__obf_94e6449148fc8624.Contains(__obf_5a12388a905e7c2e) {
			__obf_340cb9d18c761e30.Add(__obf_5a12388a905e7c2e)
		}
	}
	return &__obf_340cb9d18c761e30
}

// 函数遍历
func (__obf_7d38105b81b181cc *Set[T]) Each(__obf_240c606188546413 func(any) bool) {
	for __obf_5a12388a905e7c2e := range *__obf_7d38105b81b181cc {
		if __obf_240c606188546413(__obf_5a12388a905e7c2e) {
			break
		}
	}
}

// 返回string数组
func (__obf_7d38105b81b181cc *Set[T]) StringElem() []string {
	__obf_3d24c929f1a76b8b := make([]string, 0, len(*__obf_7d38105b81b181cc))

	for __obf_5a12388a905e7c2e := range *__obf_7d38105b81b181cc {
		__obf_3d24c929f1a76b8b = append(__obf_3d24c929f1a76b8b, fmt.Sprintf("%v", __obf_5a12388a905e7c2e))
	}
	return __obf_3d24c929f1a76b8b
}
