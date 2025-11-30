// 非线程安全set集合
package __obf_e72ce603d10d02a1

import "fmt"

type Set[T comparable] map[T]__obf_6b8a0e2afb80d958
type __obf_6b8a0e2afb80d958 struct{}

func NewSet[E comparable](__obf_84503acd3fedd4ef []E) *Set[E] {
	__obf_bdd29c369052a035 :=// 获取Set的地址
	Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_c480cb5c2a20f5dd := range __obf_84503acd3fedd4ef {
		__obf_bdd29c369052a035[// unsafe.Sizeof(void{}) // 结果为0
		__obf_c480cb5c2a20f5dd] = __obf_6b8a0e2afb80d958{}
	}
	return &__obf_bdd29c369052a035
}

func (__obf_bdd29c369052a035 *Set[T]) Add(__obf_c03cf6504cc3cf88 T) bool {
	if _, __obf_b95641e1ff7de824 := (*__obf_bdd29c369052a035)[__obf_c03cf6504cc3cf88]; __obf_b95641e1ff7de824 {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_bdd29c369052a035)[__obf_c03cf6504cc3cf88] = __obf_6b8a0e2afb80d958{}
	return true
}

func (__obf_bdd29c369052a035 *Set[T]) Contains(__obf_c480cb5c2a20f5dd T) bool {
	_, __obf_b95641e1ff7de824 := (*__obf_bdd29c369052a035)[__obf_c480cb5c2a20f5dd]
	return __obf_b95641e1ff7de824
}

func (__obf_bdd29c369052a035 *Set[T]) Size() int {
	return len(*__obf_bdd29c369052a035)
}

func (__obf_bdd29c369052a035 *Set[T]) Elem() []T {
	__obf_e4afa361e4970b8e := make([]T, 0, len(*__obf_bdd29c369052a035))
	for __obf_cce6a084cb2eec1f := range *__obf_bdd29c369052a035 {
		__obf_e4afa361e4970b8e = append(__obf_e4afa361e4970b8e, __obf_cce6a084cb2eec1f)
	}
	return __obf_e4afa361e4970b8e
}

func (__obf_bdd29c369052a035 *Set[T]) Remove(__obf_c03cf6504cc3cf88 T) {
	delete(*__obf_bdd29c369052a035, __obf_c03cf6504cc3cf88)
}

func (__obf_bdd29c369052a035 *Set[T]) Clear() {
	*__obf_bdd29c369052a035 = make(map[T]__obf_6b8a0e2afb80d958)
}

// 相等
func (__obf_bdd29c369052a035 *Set[T]) Equal(__obf_20fdf4ef89f645eb *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_bdd29c369052a035.Size() != __obf_20fdf4ef89f645eb.Size() {
		return false
	}
	for __obf_9148a001224202b7 := range *__obf_bdd29c369052a035 {
		if !__obf_20fdf4ef89f645eb.Contains(__obf_9148a001224202b7) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_bdd29c369052a035 *Set[T]) IsSubset(__obf_20fdf4ef89f645eb *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_bdd29c369052a035.Size() > __obf_20fdf4ef89f645eb.Size() {
		return false
	}
	// 迭代遍历
	for __obf_9dd08479fe906662 := range *__obf_bdd29c369052a035 {
		if !__obf_20fdf4ef89f645eb.Contains(__obf_9dd08479fe906662) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_1ebe440f7503b18d *Set[T]) Union(__obf_20fdf4ef89f645eb *Set[T]) *Set[T] {
	__obf_c41a381a52126db3 := Set[T]{}
	for __obf_9148a001224202b7 := range *__obf_1ebe440f7503b18d {
		__obf_c41a381a52126db3.
			Add(__obf_9148a001224202b7)
	}
	for __obf_9148a001224202b7 := range *__obf_20fdf4ef89f645eb {
		__obf_c41a381a52126db3.
			Add(__obf_9148a001224202b7)
	}
	return &__obf_c41a381a52126db3
}

// 交集
func (__obf_bdd29c369052a035 *Set[T]) Intersect(__obf_20fdf4ef89f645eb *Set[T]) *Set[T] {
	__obf_5510f7d21107236b := Set[T]{}
	// loop over smaller set
	if __obf_bdd29c369052a035.Size() < __obf_20fdf4ef89f645eb.Size() {
		for __obf_9148a001224202b7 := range *__obf_bdd29c369052a035 {
			if __obf_20fdf4ef89f645eb.Contains(__obf_9148a001224202b7) {
				__obf_5510f7d21107236b.
					Add(__obf_9148a001224202b7)
			}
		}
	} else {
		for __obf_9148a001224202b7 := range *__obf_20fdf4ef89f645eb {
			if __obf_bdd29c369052a035.Contains(__obf_9148a001224202b7) {
				__obf_5510f7d21107236b.
					Add(__obf_9148a001224202b7)
			}
		}
	}
	return &__obf_5510f7d21107236b
}

// 差集
func (__obf_bdd29c369052a035 *Set[T]) Difference(__obf_20fdf4ef89f645eb *Set[T]) *Set[T] {
	__obf_22cd859a9dc1a5a8 := Set[T]{}
	for __obf_9148a001224202b7 := range *__obf_bdd29c369052a035 {
		if !__obf_20fdf4ef89f645eb.Contains(__obf_9148a001224202b7) {
			__obf_22cd859a9dc1a5a8.
				Add(__obf_9148a001224202b7)
		}
	}
	return &__obf_22cd859a9dc1a5a8
}

// 函数遍历
func (__obf_bdd29c369052a035 *Set[T]) Each(__obf_966f331a674c59ef func(any) bool) {
	for __obf_9148a001224202b7 := range *__obf_bdd29c369052a035 {
		if __obf_966f331a674c59ef(__obf_9148a001224202b7) {
			break
		}
	}
}

// 返回string数组
func (__obf_bdd29c369052a035 *Set[T]) StringElem() []string {
	__obf_84503acd3fedd4ef := make([]string, 0, len(*__obf_bdd29c369052a035))

	for __obf_9148a001224202b7 := range *__obf_bdd29c369052a035 {
		__obf_84503acd3fedd4ef = append(__obf_84503acd3fedd4ef, fmt.Sprintf("%v", __obf_9148a001224202b7))
	}
	return __obf_84503acd3fedd4ef
}
