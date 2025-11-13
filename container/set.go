// 非线程安全set集合
package __obf_76599ab96a208083

import "fmt"

type Set[T comparable] map[T]__obf_b3b73112d0168107
type __obf_b3b73112d0168107 struct{}

func NewSet[E comparable](__obf_faaaaee941ca04ee []E) *Set[E] {
	// 获取Set的地址
	__obf_95818091fafad7d8 := Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_152b2d48928938f0 := range __obf_faaaaee941ca04ee {
		// unsafe.Sizeof(void{}) // 结果为0
		__obf_95818091fafad7d8[__obf_152b2d48928938f0] = __obf_b3b73112d0168107{}
	}
	return &__obf_95818091fafad7d8
}

func (__obf_95818091fafad7d8 *Set[T]) Add(__obf_8d4edb68382a63c4 T) bool {
	if _, __obf_d47b43f0df7ca990 := (*__obf_95818091fafad7d8)[__obf_8d4edb68382a63c4]; __obf_d47b43f0df7ca990 {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_95818091fafad7d8)[__obf_8d4edb68382a63c4] = __obf_b3b73112d0168107{}
	return true
}

func (__obf_95818091fafad7d8 *Set[T]) Contains(__obf_152b2d48928938f0 T) bool {
	_, __obf_d47b43f0df7ca990 := (*__obf_95818091fafad7d8)[__obf_152b2d48928938f0]
	return __obf_d47b43f0df7ca990
}

func (__obf_95818091fafad7d8 *Set[T]) Size() int {
	return len(*__obf_95818091fafad7d8)
}

func (__obf_95818091fafad7d8 *Set[T]) Elem() []T {
	__obf_cfa0937a05242cbc := make([]T, 0, len(*__obf_95818091fafad7d8))
	for __obf_2971938b3f770fd1 := range *__obf_95818091fafad7d8 {
		__obf_cfa0937a05242cbc = append(__obf_cfa0937a05242cbc, __obf_2971938b3f770fd1)
	}
	return __obf_cfa0937a05242cbc
}

func (__obf_95818091fafad7d8 *Set[T]) Remove(__obf_8d4edb68382a63c4 T) {
	delete(*__obf_95818091fafad7d8, __obf_8d4edb68382a63c4)
}

func (__obf_95818091fafad7d8 *Set[T]) Clear() {
	*__obf_95818091fafad7d8 = make(map[T]__obf_b3b73112d0168107)
}

// 相等
func (__obf_95818091fafad7d8 *Set[T]) Equal(__obf_d80176ae183ef97a *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_95818091fafad7d8.Size() != __obf_d80176ae183ef97a.Size() {
		return false
	}
	for __obf_820b59466c6fe31c := range *__obf_95818091fafad7d8 {
		if !__obf_d80176ae183ef97a.Contains(__obf_820b59466c6fe31c) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_95818091fafad7d8 *Set[T]) IsSubset(__obf_d80176ae183ef97a *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_95818091fafad7d8.Size() > __obf_d80176ae183ef97a.Size() {
		return false
	}
	// 迭代遍历
	for __obf_4368715d0a6d4f0b := range *__obf_95818091fafad7d8 {
		if !__obf_d80176ae183ef97a.Contains(__obf_4368715d0a6d4f0b) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_ba565784786f5941 *Set[T]) Union(__obf_d80176ae183ef97a *Set[T]) *Set[T] {
	__obf_6c2dc064f259dbfe := Set[T]{}
	for __obf_820b59466c6fe31c := range *__obf_ba565784786f5941 {
		__obf_6c2dc064f259dbfe.Add(__obf_820b59466c6fe31c)
	}
	for __obf_820b59466c6fe31c := range *__obf_d80176ae183ef97a {
		__obf_6c2dc064f259dbfe.Add(__obf_820b59466c6fe31c)
	}
	return &__obf_6c2dc064f259dbfe
}

// 交集
func (__obf_95818091fafad7d8 *Set[T]) Intersect(__obf_d80176ae183ef97a *Set[T]) *Set[T] {

	__obf_def2ace2e870fe4b := Set[T]{}
	// loop over smaller set
	if __obf_95818091fafad7d8.Size() < __obf_d80176ae183ef97a.Size() {
		for __obf_820b59466c6fe31c := range *__obf_95818091fafad7d8 {
			if __obf_d80176ae183ef97a.Contains(__obf_820b59466c6fe31c) {
				__obf_def2ace2e870fe4b.Add(__obf_820b59466c6fe31c)
			}
		}
	} else {
		for __obf_820b59466c6fe31c := range *__obf_d80176ae183ef97a {
			if __obf_95818091fafad7d8.Contains(__obf_820b59466c6fe31c) {
				__obf_def2ace2e870fe4b.Add(__obf_820b59466c6fe31c)
			}
		}
	}
	return &__obf_def2ace2e870fe4b
}

// 差集
func (__obf_95818091fafad7d8 *Set[T]) Difference(__obf_d80176ae183ef97a *Set[T]) *Set[T] {

	__obf_7c61e8deb9965419 := Set[T]{}
	for __obf_820b59466c6fe31c := range *__obf_95818091fafad7d8 {
		if !__obf_d80176ae183ef97a.Contains(__obf_820b59466c6fe31c) {
			__obf_7c61e8deb9965419.Add(__obf_820b59466c6fe31c)
		}
	}
	return &__obf_7c61e8deb9965419
}

// 函数遍历
func (__obf_95818091fafad7d8 *Set[T]) Each(__obf_b284d67e9c3c6ba1 func(any) bool) {
	for __obf_820b59466c6fe31c := range *__obf_95818091fafad7d8 {
		if __obf_b284d67e9c3c6ba1(__obf_820b59466c6fe31c) {
			break
		}
	}
}

// 返回string数组
func (__obf_95818091fafad7d8 *Set[T]) StringElem() []string {
	__obf_faaaaee941ca04ee := make([]string, 0, len(*__obf_95818091fafad7d8))

	for __obf_820b59466c6fe31c := range *__obf_95818091fafad7d8 {
		__obf_faaaaee941ca04ee = append(__obf_faaaaee941ca04ee, fmt.Sprintf("%v", __obf_820b59466c6fe31c))
	}
	return __obf_faaaaee941ca04ee
}
