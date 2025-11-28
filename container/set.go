// 非线程安全set集合
package __obf_4f13ac5aa043b5ce

import "fmt"

type Set[T comparable] map[T]__obf_414b6649bee8dd86
type __obf_414b6649bee8dd86 struct{}

func NewSet[E comparable](__obf_99bd94e17fee651f []E) *Set[E] {
	// 获取Set的地址
	__obf_c558169dc557d909 := Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_121994bfe70e48a4 := range __obf_99bd94e17fee651f {
		// unsafe.Sizeof(void{}) // 结果为0
		__obf_c558169dc557d909[__obf_121994bfe70e48a4] = __obf_414b6649bee8dd86{}
	}
	return &__obf_c558169dc557d909
}

func (__obf_c558169dc557d909 *Set[T]) Add(__obf_e3e988c1360a57a4 T) bool {
	if _, __obf_c7b01e58066bbb7a := (*__obf_c558169dc557d909)[__obf_e3e988c1360a57a4]; __obf_c7b01e58066bbb7a {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_c558169dc557d909)[__obf_e3e988c1360a57a4] = __obf_414b6649bee8dd86{}
	return true
}

func (__obf_c558169dc557d909 *Set[T]) Contains(__obf_121994bfe70e48a4 T) bool {
	_, __obf_c7b01e58066bbb7a := (*__obf_c558169dc557d909)[__obf_121994bfe70e48a4]
	return __obf_c7b01e58066bbb7a
}

func (__obf_c558169dc557d909 *Set[T]) Size() int {
	return len(*__obf_c558169dc557d909)
}

func (__obf_c558169dc557d909 *Set[T]) Elem() []T {
	__obf_e7eafad979141f9b := make([]T, 0, len(*__obf_c558169dc557d909))
	for __obf_a0c6719477333085 := range *__obf_c558169dc557d909 {
		__obf_e7eafad979141f9b = append(__obf_e7eafad979141f9b, __obf_a0c6719477333085)
	}
	return __obf_e7eafad979141f9b
}

func (__obf_c558169dc557d909 *Set[T]) Remove(__obf_e3e988c1360a57a4 T) {
	delete(*__obf_c558169dc557d909, __obf_e3e988c1360a57a4)
}

func (__obf_c558169dc557d909 *Set[T]) Clear() {
	*__obf_c558169dc557d909 = make(map[T]__obf_414b6649bee8dd86)
}

// 相等
func (__obf_c558169dc557d909 *Set[T]) Equal(__obf_34f3b46746db8380 *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_c558169dc557d909.Size() != __obf_34f3b46746db8380.Size() {
		return false
	}
	for __obf_abcd0da910a815c1 := range *__obf_c558169dc557d909 {
		if !__obf_34f3b46746db8380.Contains(__obf_abcd0da910a815c1) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_c558169dc557d909 *Set[T]) IsSubset(__obf_34f3b46746db8380 *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_c558169dc557d909.Size() > __obf_34f3b46746db8380.Size() {
		return false
	}
	// 迭代遍历
	for __obf_797707a3dac0ebb7 := range *__obf_c558169dc557d909 {
		if !__obf_34f3b46746db8380.Contains(__obf_797707a3dac0ebb7) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_fc094ccf2d07faed *Set[T]) Union(__obf_34f3b46746db8380 *Set[T]) *Set[T] {
	__obf_629e14de6dfaa02b := Set[T]{}
	for __obf_abcd0da910a815c1 := range *__obf_fc094ccf2d07faed {
		__obf_629e14de6dfaa02b.Add(__obf_abcd0da910a815c1)
	}
	for __obf_abcd0da910a815c1 := range *__obf_34f3b46746db8380 {
		__obf_629e14de6dfaa02b.Add(__obf_abcd0da910a815c1)
	}
	return &__obf_629e14de6dfaa02b
}

// 交集
func (__obf_c558169dc557d909 *Set[T]) Intersect(__obf_34f3b46746db8380 *Set[T]) *Set[T] {

	__obf_8d2f03606629ef84 := Set[T]{}
	// loop over smaller set
	if __obf_c558169dc557d909.Size() < __obf_34f3b46746db8380.Size() {
		for __obf_abcd0da910a815c1 := range *__obf_c558169dc557d909 {
			if __obf_34f3b46746db8380.Contains(__obf_abcd0da910a815c1) {
				__obf_8d2f03606629ef84.Add(__obf_abcd0da910a815c1)
			}
		}
	} else {
		for __obf_abcd0da910a815c1 := range *__obf_34f3b46746db8380 {
			if __obf_c558169dc557d909.Contains(__obf_abcd0da910a815c1) {
				__obf_8d2f03606629ef84.Add(__obf_abcd0da910a815c1)
			}
		}
	}
	return &__obf_8d2f03606629ef84
}

// 差集
func (__obf_c558169dc557d909 *Set[T]) Difference(__obf_34f3b46746db8380 *Set[T]) *Set[T] {

	__obf_8ca5280c350fcf61 := Set[T]{}
	for __obf_abcd0da910a815c1 := range *__obf_c558169dc557d909 {
		if !__obf_34f3b46746db8380.Contains(__obf_abcd0da910a815c1) {
			__obf_8ca5280c350fcf61.Add(__obf_abcd0da910a815c1)
		}
	}
	return &__obf_8ca5280c350fcf61
}

// 函数遍历
func (__obf_c558169dc557d909 *Set[T]) Each(__obf_97bed5d44b6b723e func(any) bool) {
	for __obf_abcd0da910a815c1 := range *__obf_c558169dc557d909 {
		if __obf_97bed5d44b6b723e(__obf_abcd0da910a815c1) {
			break
		}
	}
}

// 返回string数组
func (__obf_c558169dc557d909 *Set[T]) StringElem() []string {
	__obf_99bd94e17fee651f := make([]string, 0, len(*__obf_c558169dc557d909))

	for __obf_abcd0da910a815c1 := range *__obf_c558169dc557d909 {
		__obf_99bd94e17fee651f = append(__obf_99bd94e17fee651f, fmt.Sprintf("%v", __obf_abcd0da910a815c1))
	}
	return __obf_99bd94e17fee651f
}
