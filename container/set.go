// 非线程安全set集合
package __obf_af42fb6cde2beed6

import "fmt"

type Set[T comparable] map[T]__obf_74d757c652d3e569
type __obf_74d757c652d3e569 struct{}

func NewSet[E comparable](__obf_1f87755a03016aac []E) *Set[E] {
	// 获取Set的地址
	__obf_ef12820de42f2cbf := Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_626ebcb2efd33e98 := range __obf_1f87755a03016aac {
		// unsafe.Sizeof(void{}) // 结果为0
		__obf_ef12820de42f2cbf[__obf_626ebcb2efd33e98] = __obf_74d757c652d3e569{}
	}
	return &__obf_ef12820de42f2cbf
}

func (__obf_ef12820de42f2cbf *Set[T]) Add(__obf_d6db326dc08be53b T) bool {
	if _, __obf_021756766e1db5da := (*__obf_ef12820de42f2cbf)[__obf_d6db326dc08be53b]; __obf_021756766e1db5da {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_ef12820de42f2cbf)[__obf_d6db326dc08be53b] = __obf_74d757c652d3e569{}
	return true
}

func (__obf_ef12820de42f2cbf *Set[T]) Contains(__obf_626ebcb2efd33e98 T) bool {
	_, __obf_021756766e1db5da := (*__obf_ef12820de42f2cbf)[__obf_626ebcb2efd33e98]
	return __obf_021756766e1db5da
}

func (__obf_ef12820de42f2cbf *Set[T]) Size() int {
	return len(*__obf_ef12820de42f2cbf)
}

func (__obf_ef12820de42f2cbf *Set[T]) Elem() []T {
	__obf_a0cee0a97b92008e := make([]T, 0, len(*__obf_ef12820de42f2cbf))
	for __obf_43d363f3f5d57e86 := range *__obf_ef12820de42f2cbf {
		__obf_a0cee0a97b92008e = append(__obf_a0cee0a97b92008e, __obf_43d363f3f5d57e86)
	}
	return __obf_a0cee0a97b92008e
}

func (__obf_ef12820de42f2cbf *Set[T]) Remove(__obf_d6db326dc08be53b T) {
	delete(*__obf_ef12820de42f2cbf, __obf_d6db326dc08be53b)
}

func (__obf_ef12820de42f2cbf *Set[T]) Clear() {
	*__obf_ef12820de42f2cbf = make(map[T]__obf_74d757c652d3e569)
}

// 相等
func (__obf_ef12820de42f2cbf *Set[T]) Equal(__obf_14cd44a380175d47 *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_ef12820de42f2cbf.Size() != __obf_14cd44a380175d47.Size() {
		return false
	}
	for __obf_c58444540a03feb7 := range *__obf_ef12820de42f2cbf {
		if !__obf_14cd44a380175d47.Contains(__obf_c58444540a03feb7) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_ef12820de42f2cbf *Set[T]) IsSubset(__obf_14cd44a380175d47 *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_ef12820de42f2cbf.Size() > __obf_14cd44a380175d47.Size() {
		return false
	}
	// 迭代遍历
	for __obf_55cde42f6d47c5be := range *__obf_ef12820de42f2cbf {
		if !__obf_14cd44a380175d47.Contains(__obf_55cde42f6d47c5be) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_4e962dbd0255751a *Set[T]) Union(__obf_14cd44a380175d47 *Set[T]) *Set[T] {
	__obf_eed9bef3faf94405 := Set[T]{}
	for __obf_c58444540a03feb7 := range *__obf_4e962dbd0255751a {
		__obf_eed9bef3faf94405.Add(__obf_c58444540a03feb7)
	}
	for __obf_c58444540a03feb7 := range *__obf_14cd44a380175d47 {
		__obf_eed9bef3faf94405.Add(__obf_c58444540a03feb7)
	}
	return &__obf_eed9bef3faf94405
}

// 交集
func (__obf_ef12820de42f2cbf *Set[T]) Intersect(__obf_14cd44a380175d47 *Set[T]) *Set[T] {

	__obf_a0c11f316bf81745 := Set[T]{}
	// loop over smaller set
	if __obf_ef12820de42f2cbf.Size() < __obf_14cd44a380175d47.Size() {
		for __obf_c58444540a03feb7 := range *__obf_ef12820de42f2cbf {
			if __obf_14cd44a380175d47.Contains(__obf_c58444540a03feb7) {
				__obf_a0c11f316bf81745.Add(__obf_c58444540a03feb7)
			}
		}
	} else {
		for __obf_c58444540a03feb7 := range *__obf_14cd44a380175d47 {
			if __obf_ef12820de42f2cbf.Contains(__obf_c58444540a03feb7) {
				__obf_a0c11f316bf81745.Add(__obf_c58444540a03feb7)
			}
		}
	}
	return &__obf_a0c11f316bf81745
}

// 差集
func (__obf_ef12820de42f2cbf *Set[T]) Difference(__obf_14cd44a380175d47 *Set[T]) *Set[T] {

	__obf_5ae874ddfab62f4a := Set[T]{}
	for __obf_c58444540a03feb7 := range *__obf_ef12820de42f2cbf {
		if !__obf_14cd44a380175d47.Contains(__obf_c58444540a03feb7) {
			__obf_5ae874ddfab62f4a.Add(__obf_c58444540a03feb7)
		}
	}
	return &__obf_5ae874ddfab62f4a
}

// 函数遍历
func (__obf_ef12820de42f2cbf *Set[T]) Each(__obf_4f34168fb196a9f5 func(any) bool) {
	for __obf_c58444540a03feb7 := range *__obf_ef12820de42f2cbf {
		if __obf_4f34168fb196a9f5(__obf_c58444540a03feb7) {
			break
		}
	}
}

// 返回string数组
func (__obf_ef12820de42f2cbf *Set[T]) StringElem() []string {
	__obf_1f87755a03016aac := make([]string, 0, len(*__obf_ef12820de42f2cbf))

	for __obf_c58444540a03feb7 := range *__obf_ef12820de42f2cbf {
		__obf_1f87755a03016aac = append(__obf_1f87755a03016aac, fmt.Sprintf("%v", __obf_c58444540a03feb7))
	}
	return __obf_1f87755a03016aac
}
