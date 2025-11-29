// 非线程安全set集合
package __obf_b0bebe5eb45b8ad6

import "fmt"

type Set[T comparable] map[T]__obf_a77b6ec582e8e9a7
type __obf_a77b6ec582e8e9a7 struct{}

func NewSet[E comparable](__obf_c2c6fb2808a047a0 []E) *Set[E] {
	__obf_c161febd615faf6b :=// 获取Set的地址
	Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_162cd0c81f138e95 := range __obf_c2c6fb2808a047a0 {
		__obf_c161febd615faf6b[// unsafe.Sizeof(void{}) // 结果为0
		__obf_162cd0c81f138e95] = __obf_a77b6ec582e8e9a7{}
	}
	return &__obf_c161febd615faf6b
}

func (__obf_c161febd615faf6b *Set[T]) Add(__obf_095a614fd95fc26e T) bool {
	if _, __obf_e8c1fb9f7287beef := (*__obf_c161febd615faf6b)[__obf_095a614fd95fc26e]; __obf_e8c1fb9f7287beef {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_c161febd615faf6b)[__obf_095a614fd95fc26e] = __obf_a77b6ec582e8e9a7{}
	return true
}

func (__obf_c161febd615faf6b *Set[T]) Contains(__obf_162cd0c81f138e95 T) bool {
	_, __obf_e8c1fb9f7287beef := (*__obf_c161febd615faf6b)[__obf_162cd0c81f138e95]
	return __obf_e8c1fb9f7287beef
}

func (__obf_c161febd615faf6b *Set[T]) Size() int {
	return len(*__obf_c161febd615faf6b)
}

func (__obf_c161febd615faf6b *Set[T]) Elem() []T {
	__obf_1e824d5b9fe4a6dc := make([]T, 0, len(*__obf_c161febd615faf6b))
	for __obf_dca1c5db9aad12e8 := range *__obf_c161febd615faf6b {
		__obf_1e824d5b9fe4a6dc = append(__obf_1e824d5b9fe4a6dc, __obf_dca1c5db9aad12e8)
	}
	return __obf_1e824d5b9fe4a6dc
}

func (__obf_c161febd615faf6b *Set[T]) Remove(__obf_095a614fd95fc26e T) {
	delete(*__obf_c161febd615faf6b, __obf_095a614fd95fc26e)
}

func (__obf_c161febd615faf6b *Set[T]) Clear() {
	*__obf_c161febd615faf6b = make(map[T]__obf_a77b6ec582e8e9a7)
}

// 相等
func (__obf_c161febd615faf6b *Set[T]) Equal(__obf_84359bf54a2548e8 *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_c161febd615faf6b.Size() != __obf_84359bf54a2548e8.Size() {
		return false
	}
	for __obf_326d32b136dc0852 := range *__obf_c161febd615faf6b {
		if !__obf_84359bf54a2548e8.Contains(__obf_326d32b136dc0852) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_c161febd615faf6b *Set[T]) IsSubset(__obf_84359bf54a2548e8 *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_c161febd615faf6b.Size() > __obf_84359bf54a2548e8.Size() {
		return false
	}
	// 迭代遍历
	for __obf_9c15798bcb95be3e := range *__obf_c161febd615faf6b {
		if !__obf_84359bf54a2548e8.Contains(__obf_9c15798bcb95be3e) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_ff686dbd8069ab9c *Set[T]) Union(__obf_84359bf54a2548e8 *Set[T]) *Set[T] {
	__obf_8161088f8f490ae7 := Set[T]{}
	for __obf_326d32b136dc0852 := range *__obf_ff686dbd8069ab9c {
		__obf_8161088f8f490ae7.
			Add(__obf_326d32b136dc0852)
	}
	for __obf_326d32b136dc0852 := range *__obf_84359bf54a2548e8 {
		__obf_8161088f8f490ae7.
			Add(__obf_326d32b136dc0852)
	}
	return &__obf_8161088f8f490ae7
}

// 交集
func (__obf_c161febd615faf6b *Set[T]) Intersect(__obf_84359bf54a2548e8 *Set[T]) *Set[T] {
	__obf_8a2aedfe93db4e5c := Set[T]{}
	// loop over smaller set
	if __obf_c161febd615faf6b.Size() < __obf_84359bf54a2548e8.Size() {
		for __obf_326d32b136dc0852 := range *__obf_c161febd615faf6b {
			if __obf_84359bf54a2548e8.Contains(__obf_326d32b136dc0852) {
				__obf_8a2aedfe93db4e5c.
					Add(__obf_326d32b136dc0852)
			}
		}
	} else {
		for __obf_326d32b136dc0852 := range *__obf_84359bf54a2548e8 {
			if __obf_c161febd615faf6b.Contains(__obf_326d32b136dc0852) {
				__obf_8a2aedfe93db4e5c.
					Add(__obf_326d32b136dc0852)
			}
		}
	}
	return &__obf_8a2aedfe93db4e5c
}

// 差集
func (__obf_c161febd615faf6b *Set[T]) Difference(__obf_84359bf54a2548e8 *Set[T]) *Set[T] {
	__obf_d167bea51b1f752c := Set[T]{}
	for __obf_326d32b136dc0852 := range *__obf_c161febd615faf6b {
		if !__obf_84359bf54a2548e8.Contains(__obf_326d32b136dc0852) {
			__obf_d167bea51b1f752c.
				Add(__obf_326d32b136dc0852)
		}
	}
	return &__obf_d167bea51b1f752c
}

// 函数遍历
func (__obf_c161febd615faf6b *Set[T]) Each(__obf_22a0a28267af8af2 func(any) bool) {
	for __obf_326d32b136dc0852 := range *__obf_c161febd615faf6b {
		if __obf_22a0a28267af8af2(__obf_326d32b136dc0852) {
			break
		}
	}
}

// 返回string数组
func (__obf_c161febd615faf6b *Set[T]) StringElem() []string {
	__obf_c2c6fb2808a047a0 := make([]string, 0, len(*__obf_c161febd615faf6b))

	for __obf_326d32b136dc0852 := range *__obf_c161febd615faf6b {
		__obf_c2c6fb2808a047a0 = append(__obf_c2c6fb2808a047a0, fmt.Sprintf("%v", __obf_326d32b136dc0852))
	}
	return __obf_c2c6fb2808a047a0
}
