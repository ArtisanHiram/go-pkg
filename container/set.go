// 非线程安全set集合
package __obf_038560a94647875f

import "fmt"

type Set[T comparable] map[T]__obf_75885acf59b596f4
type __obf_75885acf59b596f4 struct{}

func NewSet[E comparable](__obf_218b8d5ae68c5552 []E) *Set[E] {
	__obf_06832686dcef705f :=// 获取Set的地址
	Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_e8eb1854ef7cf459 := range __obf_218b8d5ae68c5552 {
		__obf_06832686dcef705f[// unsafe.Sizeof(void{}) // 结果为0
		__obf_e8eb1854ef7cf459] = __obf_75885acf59b596f4{}
	}
	return &__obf_06832686dcef705f
}

func (__obf_06832686dcef705f *Set[T]) Add(__obf_d2ea251c77215677 T) bool {
	if _, __obf_5904187971b4cb86 := (*__obf_06832686dcef705f)[__obf_d2ea251c77215677]; __obf_5904187971b4cb86 {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_06832686dcef705f)[__obf_d2ea251c77215677] = __obf_75885acf59b596f4{}
	return true
}

func (__obf_06832686dcef705f *Set[T]) Contains(__obf_e8eb1854ef7cf459 T) bool {
	_, __obf_5904187971b4cb86 := (*__obf_06832686dcef705f)[__obf_e8eb1854ef7cf459]
	return __obf_5904187971b4cb86
}

func (__obf_06832686dcef705f *Set[T]) Size() int {
	return len(*__obf_06832686dcef705f)
}

func (__obf_06832686dcef705f *Set[T]) Elem() []T {
	__obf_0414b3afaf21f9fa := make([]T, 0, len(*__obf_06832686dcef705f))
	for __obf_ed002d86aefb8e41 := range *__obf_06832686dcef705f {
		__obf_0414b3afaf21f9fa = append(__obf_0414b3afaf21f9fa, __obf_ed002d86aefb8e41)
	}
	return __obf_0414b3afaf21f9fa
}

func (__obf_06832686dcef705f *Set[T]) Remove(__obf_d2ea251c77215677 T) {
	delete(*__obf_06832686dcef705f, __obf_d2ea251c77215677)
}

func (__obf_06832686dcef705f *Set[T]) Clear() {
	*__obf_06832686dcef705f = make(map[T]__obf_75885acf59b596f4)
}

// 相等
func (__obf_06832686dcef705f *Set[T]) Equal(__obf_8045619291d56510 *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_06832686dcef705f.Size() != __obf_8045619291d56510.Size() {
		return false
	}
	for __obf_3c623453454fedc7 := range *__obf_06832686dcef705f {
		if !__obf_8045619291d56510.Contains(__obf_3c623453454fedc7) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_06832686dcef705f *Set[T]) IsSubset(__obf_8045619291d56510 *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_06832686dcef705f.Size() > __obf_8045619291d56510.Size() {
		return false
	}
	// 迭代遍历
	for __obf_52e73bb48e810dd2 := range *__obf_06832686dcef705f {
		if !__obf_8045619291d56510.Contains(__obf_52e73bb48e810dd2) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_096e79c377ff1b0d *Set[T]) Union(__obf_8045619291d56510 *Set[T]) *Set[T] {
	__obf_71dc1136337b167c := Set[T]{}
	for __obf_3c623453454fedc7 := range *__obf_096e79c377ff1b0d {
		__obf_71dc1136337b167c.
			Add(__obf_3c623453454fedc7)
	}
	for __obf_3c623453454fedc7 := range *__obf_8045619291d56510 {
		__obf_71dc1136337b167c.
			Add(__obf_3c623453454fedc7)
	}
	return &__obf_71dc1136337b167c
}

// 交集
func (__obf_06832686dcef705f *Set[T]) Intersect(__obf_8045619291d56510 *Set[T]) *Set[T] {
	__obf_7f6ceb9c2052a400 := Set[T]{}
	// loop over smaller set
	if __obf_06832686dcef705f.Size() < __obf_8045619291d56510.Size() {
		for __obf_3c623453454fedc7 := range *__obf_06832686dcef705f {
			if __obf_8045619291d56510.Contains(__obf_3c623453454fedc7) {
				__obf_7f6ceb9c2052a400.
					Add(__obf_3c623453454fedc7)
			}
		}
	} else {
		for __obf_3c623453454fedc7 := range *__obf_8045619291d56510 {
			if __obf_06832686dcef705f.Contains(__obf_3c623453454fedc7) {
				__obf_7f6ceb9c2052a400.
					Add(__obf_3c623453454fedc7)
			}
		}
	}
	return &__obf_7f6ceb9c2052a400
}

// 差集
func (__obf_06832686dcef705f *Set[T]) Difference(__obf_8045619291d56510 *Set[T]) *Set[T] {
	__obf_c4102e5a023963b5 := Set[T]{}
	for __obf_3c623453454fedc7 := range *__obf_06832686dcef705f {
		if !__obf_8045619291d56510.Contains(__obf_3c623453454fedc7) {
			__obf_c4102e5a023963b5.
				Add(__obf_3c623453454fedc7)
		}
	}
	return &__obf_c4102e5a023963b5
}

// 函数遍历
func (__obf_06832686dcef705f *Set[T]) Each(__obf_df5c7b9ab33daa5f func(any) bool) {
	for __obf_3c623453454fedc7 := range *__obf_06832686dcef705f {
		if __obf_df5c7b9ab33daa5f(__obf_3c623453454fedc7) {
			break
		}
	}
}

// 返回string数组
func (__obf_06832686dcef705f *Set[T]) StringElem() []string {
	__obf_218b8d5ae68c5552 := make([]string, 0, len(*__obf_06832686dcef705f))

	for __obf_3c623453454fedc7 := range *__obf_06832686dcef705f {
		__obf_218b8d5ae68c5552 = append(__obf_218b8d5ae68c5552, fmt.Sprintf("%v", __obf_3c623453454fedc7))
	}
	return __obf_218b8d5ae68c5552
}
