// 非线程安全set集合
package __obf_62eba4024f8fa381

import "fmt"

type Set[T comparable] map[T]__obf_f4e2d1257ba9de1e
type __obf_f4e2d1257ba9de1e struct{}

func NewSet[E comparable](__obf_e15be3f07fb220b3 []E) *Set[E] {
	// 获取Set的地址
	__obf_97bc815048dc3f5e := Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_f29ff4c234636d25 := range __obf_e15be3f07fb220b3 {
		// unsafe.Sizeof(void{}) // 结果为0
		__obf_97bc815048dc3f5e[__obf_f29ff4c234636d25] = __obf_f4e2d1257ba9de1e{}
	}
	return &__obf_97bc815048dc3f5e
}

func (__obf_97bc815048dc3f5e *Set[T]) Add(__obf_1376e2b9d2469c18 T) bool {
	if _, __obf_b47c9f51c1edd10c := (*__obf_97bc815048dc3f5e)[__obf_1376e2b9d2469c18]; __obf_b47c9f51c1edd10c {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_97bc815048dc3f5e)[__obf_1376e2b9d2469c18] = __obf_f4e2d1257ba9de1e{}
	return true
}

func (__obf_97bc815048dc3f5e *Set[T]) Contains(__obf_f29ff4c234636d25 T) bool {
	_, __obf_b47c9f51c1edd10c := (*__obf_97bc815048dc3f5e)[__obf_f29ff4c234636d25]
	return __obf_b47c9f51c1edd10c
}

func (__obf_97bc815048dc3f5e *Set[T]) Size() int {
	return len(*__obf_97bc815048dc3f5e)
}

func (__obf_97bc815048dc3f5e *Set[T]) Elem() []T {
	__obf_ea1c0c515eb08cda := make([]T, 0, len(*__obf_97bc815048dc3f5e))
	for __obf_f067a403d3b1b52a := range *__obf_97bc815048dc3f5e {
		__obf_ea1c0c515eb08cda = append(__obf_ea1c0c515eb08cda, __obf_f067a403d3b1b52a)
	}
	return __obf_ea1c0c515eb08cda
}

func (__obf_97bc815048dc3f5e *Set[T]) Remove(__obf_1376e2b9d2469c18 T) {
	delete(*__obf_97bc815048dc3f5e, __obf_1376e2b9d2469c18)
}

func (__obf_97bc815048dc3f5e *Set[T]) Clear() {
	*__obf_97bc815048dc3f5e = make(map[T]__obf_f4e2d1257ba9de1e)
}

// 相等
func (__obf_97bc815048dc3f5e *Set[T]) Equal(__obf_c2c35aac73c09bc5 *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_97bc815048dc3f5e.Size() != __obf_c2c35aac73c09bc5.Size() {
		return false
	}
	for __obf_3c95f4299d9b5c3d := range *__obf_97bc815048dc3f5e {
		if !__obf_c2c35aac73c09bc5.Contains(__obf_3c95f4299d9b5c3d) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_97bc815048dc3f5e *Set[T]) IsSubset(__obf_c2c35aac73c09bc5 *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_97bc815048dc3f5e.Size() > __obf_c2c35aac73c09bc5.Size() {
		return false
	}
	// 迭代遍历
	for __obf_df070ab4c712506c := range *__obf_97bc815048dc3f5e {
		if !__obf_c2c35aac73c09bc5.Contains(__obf_df070ab4c712506c) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_dee71b0d7df2b4cf *Set[T]) Union(__obf_c2c35aac73c09bc5 *Set[T]) *Set[T] {
	__obf_73f904d696ad2b3a := Set[T]{}
	for __obf_3c95f4299d9b5c3d := range *__obf_dee71b0d7df2b4cf {
		__obf_73f904d696ad2b3a.Add(__obf_3c95f4299d9b5c3d)
	}
	for __obf_3c95f4299d9b5c3d := range *__obf_c2c35aac73c09bc5 {
		__obf_73f904d696ad2b3a.Add(__obf_3c95f4299d9b5c3d)
	}
	return &__obf_73f904d696ad2b3a
}

// 交集
func (__obf_97bc815048dc3f5e *Set[T]) Intersect(__obf_c2c35aac73c09bc5 *Set[T]) *Set[T] {

	__obf_36f5a0a64fb05033 := Set[T]{}
	// loop over smaller set
	if __obf_97bc815048dc3f5e.Size() < __obf_c2c35aac73c09bc5.Size() {
		for __obf_3c95f4299d9b5c3d := range *__obf_97bc815048dc3f5e {
			if __obf_c2c35aac73c09bc5.Contains(__obf_3c95f4299d9b5c3d) {
				__obf_36f5a0a64fb05033.Add(__obf_3c95f4299d9b5c3d)
			}
		}
	} else {
		for __obf_3c95f4299d9b5c3d := range *__obf_c2c35aac73c09bc5 {
			if __obf_97bc815048dc3f5e.Contains(__obf_3c95f4299d9b5c3d) {
				__obf_36f5a0a64fb05033.Add(__obf_3c95f4299d9b5c3d)
			}
		}
	}
	return &__obf_36f5a0a64fb05033
}

// 差集
func (__obf_97bc815048dc3f5e *Set[T]) Difference(__obf_c2c35aac73c09bc5 *Set[T]) *Set[T] {

	__obf_8d7d331edf56059e := Set[T]{}
	for __obf_3c95f4299d9b5c3d := range *__obf_97bc815048dc3f5e {
		if !__obf_c2c35aac73c09bc5.Contains(__obf_3c95f4299d9b5c3d) {
			__obf_8d7d331edf56059e.Add(__obf_3c95f4299d9b5c3d)
		}
	}
	return &__obf_8d7d331edf56059e
}

// 函数遍历
func (__obf_97bc815048dc3f5e *Set[T]) Each(__obf_5a9895ffefb500dd func(any) bool) {
	for __obf_3c95f4299d9b5c3d := range *__obf_97bc815048dc3f5e {
		if __obf_5a9895ffefb500dd(__obf_3c95f4299d9b5c3d) {
			break
		}
	}
}

// 返回string数组
func (__obf_97bc815048dc3f5e *Set[T]) StringElem() []string {
	__obf_e15be3f07fb220b3 := make([]string, 0, len(*__obf_97bc815048dc3f5e))

	for __obf_3c95f4299d9b5c3d := range *__obf_97bc815048dc3f5e {
		__obf_e15be3f07fb220b3 = append(__obf_e15be3f07fb220b3, fmt.Sprintf("%v", __obf_3c95f4299d9b5c3d))
	}
	return __obf_e15be3f07fb220b3
}
