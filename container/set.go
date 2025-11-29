// 非线程安全set集合
package __obf_90a4883a02d0b41c

import "fmt"

type Set[T comparable] map[T]__obf_f83d09a1d4f8aae2
type __obf_f83d09a1d4f8aae2 struct{}

func NewSet[E comparable](__obf_be3616a2ebcb3eaa []E) *Set[E] {
	__obf_669fe3d7fa25df0a :=// 获取Set的地址
	Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_94eeb0188a6c62c2 := range __obf_be3616a2ebcb3eaa {
		__obf_669fe3d7fa25df0a[// unsafe.Sizeof(void{}) // 结果为0
		__obf_94eeb0188a6c62c2] = __obf_f83d09a1d4f8aae2{}
	}
	return &__obf_669fe3d7fa25df0a
}

func (__obf_669fe3d7fa25df0a *Set[T]) Add(__obf_f315019af10902c2 T) bool {
	if _, __obf_c16a0dd3a76f078a := (*__obf_669fe3d7fa25df0a)[__obf_f315019af10902c2]; __obf_c16a0dd3a76f078a {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_669fe3d7fa25df0a)[__obf_f315019af10902c2] = __obf_f83d09a1d4f8aae2{}
	return true
}

func (__obf_669fe3d7fa25df0a *Set[T]) Contains(__obf_94eeb0188a6c62c2 T) bool {
	_, __obf_c16a0dd3a76f078a := (*__obf_669fe3d7fa25df0a)[__obf_94eeb0188a6c62c2]
	return __obf_c16a0dd3a76f078a
}

func (__obf_669fe3d7fa25df0a *Set[T]) Size() int {
	return len(*__obf_669fe3d7fa25df0a)
}

func (__obf_669fe3d7fa25df0a *Set[T]) Elem() []T {
	__obf_f72137bb13d161c8 := make([]T, 0, len(*__obf_669fe3d7fa25df0a))
	for __obf_7d788ca4b6a826e5 := range *__obf_669fe3d7fa25df0a {
		__obf_f72137bb13d161c8 = append(__obf_f72137bb13d161c8, __obf_7d788ca4b6a826e5)
	}
	return __obf_f72137bb13d161c8
}

func (__obf_669fe3d7fa25df0a *Set[T]) Remove(__obf_f315019af10902c2 T) {
	delete(*__obf_669fe3d7fa25df0a, __obf_f315019af10902c2)
}

func (__obf_669fe3d7fa25df0a *Set[T]) Clear() {
	*__obf_669fe3d7fa25df0a = make(map[T]__obf_f83d09a1d4f8aae2)
}

// 相等
func (__obf_669fe3d7fa25df0a *Set[T]) Equal(__obf_4644ca3cf6956164 *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_669fe3d7fa25df0a.Size() != __obf_4644ca3cf6956164.Size() {
		return false
	}
	for __obf_1afa695a0c002d93 := range *__obf_669fe3d7fa25df0a {
		if !__obf_4644ca3cf6956164.Contains(__obf_1afa695a0c002d93) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_669fe3d7fa25df0a *Set[T]) IsSubset(__obf_4644ca3cf6956164 *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_669fe3d7fa25df0a.Size() > __obf_4644ca3cf6956164.Size() {
		return false
	}
	// 迭代遍历
	for __obf_5bba24c1758bbf28 := range *__obf_669fe3d7fa25df0a {
		if !__obf_4644ca3cf6956164.Contains(__obf_5bba24c1758bbf28) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_7b74e0720a699650 *Set[T]) Union(__obf_4644ca3cf6956164 *Set[T]) *Set[T] {
	__obf_567904532ea69b70 := Set[T]{}
	for __obf_1afa695a0c002d93 := range *__obf_7b74e0720a699650 {
		__obf_567904532ea69b70.
			Add(__obf_1afa695a0c002d93)
	}
	for __obf_1afa695a0c002d93 := range *__obf_4644ca3cf6956164 {
		__obf_567904532ea69b70.
			Add(__obf_1afa695a0c002d93)
	}
	return &__obf_567904532ea69b70
}

// 交集
func (__obf_669fe3d7fa25df0a *Set[T]) Intersect(__obf_4644ca3cf6956164 *Set[T]) *Set[T] {
	__obf_38b591bac25ca4a0 := Set[T]{}
	// loop over smaller set
	if __obf_669fe3d7fa25df0a.Size() < __obf_4644ca3cf6956164.Size() {
		for __obf_1afa695a0c002d93 := range *__obf_669fe3d7fa25df0a {
			if __obf_4644ca3cf6956164.Contains(__obf_1afa695a0c002d93) {
				__obf_38b591bac25ca4a0.
					Add(__obf_1afa695a0c002d93)
			}
		}
	} else {
		for __obf_1afa695a0c002d93 := range *__obf_4644ca3cf6956164 {
			if __obf_669fe3d7fa25df0a.Contains(__obf_1afa695a0c002d93) {
				__obf_38b591bac25ca4a0.
					Add(__obf_1afa695a0c002d93)
			}
		}
	}
	return &__obf_38b591bac25ca4a0
}

// 差集
func (__obf_669fe3d7fa25df0a *Set[T]) Difference(__obf_4644ca3cf6956164 *Set[T]) *Set[T] {
	__obf_da21fa1afeb33b60 := Set[T]{}
	for __obf_1afa695a0c002d93 := range *__obf_669fe3d7fa25df0a {
		if !__obf_4644ca3cf6956164.Contains(__obf_1afa695a0c002d93) {
			__obf_da21fa1afeb33b60.
				Add(__obf_1afa695a0c002d93)
		}
	}
	return &__obf_da21fa1afeb33b60
}

// 函数遍历
func (__obf_669fe3d7fa25df0a *Set[T]) Each(__obf_2290134cd72fb10a func(any) bool) {
	for __obf_1afa695a0c002d93 := range *__obf_669fe3d7fa25df0a {
		if __obf_2290134cd72fb10a(__obf_1afa695a0c002d93) {
			break
		}
	}
}

// 返回string数组
func (__obf_669fe3d7fa25df0a *Set[T]) StringElem() []string {
	__obf_be3616a2ebcb3eaa := make([]string, 0, len(*__obf_669fe3d7fa25df0a))

	for __obf_1afa695a0c002d93 := range *__obf_669fe3d7fa25df0a {
		__obf_be3616a2ebcb3eaa = append(__obf_be3616a2ebcb3eaa, fmt.Sprintf("%v", __obf_1afa695a0c002d93))
	}
	return __obf_be3616a2ebcb3eaa
}
