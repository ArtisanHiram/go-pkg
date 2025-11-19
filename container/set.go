// 非线程安全set集合
package __obf_9861fa13140c30a3

import "fmt"

type Set[T comparable] map[T]__obf_a16343c056e40b04
type __obf_a16343c056e40b04 struct{}

func NewSet[E comparable](__obf_218c469bce080450 []E) *Set[E] {
	// 获取Set的地址
	__obf_82e5a2bc6d3e754a := Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_e201ec52e8e19314 := range __obf_218c469bce080450 {
		// unsafe.Sizeof(void{}) // 结果为0
		__obf_82e5a2bc6d3e754a[__obf_e201ec52e8e19314] = __obf_a16343c056e40b04{}
	}
	return &__obf_82e5a2bc6d3e754a
}

func (__obf_82e5a2bc6d3e754a *Set[T]) Add(__obf_35fc60e92e4b05ba T) bool {
	if _, __obf_da2280290c7dd975 := (*__obf_82e5a2bc6d3e754a)[__obf_35fc60e92e4b05ba]; __obf_da2280290c7dd975 {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_82e5a2bc6d3e754a)[__obf_35fc60e92e4b05ba] = __obf_a16343c056e40b04{}
	return true
}

func (__obf_82e5a2bc6d3e754a *Set[T]) Contains(__obf_e201ec52e8e19314 T) bool {
	_, __obf_da2280290c7dd975 := (*__obf_82e5a2bc6d3e754a)[__obf_e201ec52e8e19314]
	return __obf_da2280290c7dd975
}

func (__obf_82e5a2bc6d3e754a *Set[T]) Size() int {
	return len(*__obf_82e5a2bc6d3e754a)
}

func (__obf_82e5a2bc6d3e754a *Set[T]) Elem() []T {
	__obf_6c59a1f64b802da6 := make([]T, 0, len(*__obf_82e5a2bc6d3e754a))
	for __obf_299b9ac8f6041f1b := range *__obf_82e5a2bc6d3e754a {
		__obf_6c59a1f64b802da6 = append(__obf_6c59a1f64b802da6, __obf_299b9ac8f6041f1b)
	}
	return __obf_6c59a1f64b802da6
}

func (__obf_82e5a2bc6d3e754a *Set[T]) Remove(__obf_35fc60e92e4b05ba T) {
	delete(*__obf_82e5a2bc6d3e754a, __obf_35fc60e92e4b05ba)
}

func (__obf_82e5a2bc6d3e754a *Set[T]) Clear() {
	*__obf_82e5a2bc6d3e754a = make(map[T]__obf_a16343c056e40b04)
}

// 相等
func (__obf_82e5a2bc6d3e754a *Set[T]) Equal(__obf_c6caadad3910b4a2 *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_82e5a2bc6d3e754a.Size() != __obf_c6caadad3910b4a2.Size() {
		return false
	}
	for __obf_1cfa18c298665e4b := range *__obf_82e5a2bc6d3e754a {
		if !__obf_c6caadad3910b4a2.Contains(__obf_1cfa18c298665e4b) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_82e5a2bc6d3e754a *Set[T]) IsSubset(__obf_c6caadad3910b4a2 *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_82e5a2bc6d3e754a.Size() > __obf_c6caadad3910b4a2.Size() {
		return false
	}
	// 迭代遍历
	for __obf_43194ec765d86867 := range *__obf_82e5a2bc6d3e754a {
		if !__obf_c6caadad3910b4a2.Contains(__obf_43194ec765d86867) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_c181c777f41b4768 *Set[T]) Union(__obf_c6caadad3910b4a2 *Set[T]) *Set[T] {
	__obf_820339352e1641a3 := Set[T]{}
	for __obf_1cfa18c298665e4b := range *__obf_c181c777f41b4768 {
		__obf_820339352e1641a3.Add(__obf_1cfa18c298665e4b)
	}
	for __obf_1cfa18c298665e4b := range *__obf_c6caadad3910b4a2 {
		__obf_820339352e1641a3.Add(__obf_1cfa18c298665e4b)
	}
	return &__obf_820339352e1641a3
}

// 交集
func (__obf_82e5a2bc6d3e754a *Set[T]) Intersect(__obf_c6caadad3910b4a2 *Set[T]) *Set[T] {

	__obf_bbe2385ee84f0917 := Set[T]{}
	// loop over smaller set
	if __obf_82e5a2bc6d3e754a.Size() < __obf_c6caadad3910b4a2.Size() {
		for __obf_1cfa18c298665e4b := range *__obf_82e5a2bc6d3e754a {
			if __obf_c6caadad3910b4a2.Contains(__obf_1cfa18c298665e4b) {
				__obf_bbe2385ee84f0917.Add(__obf_1cfa18c298665e4b)
			}
		}
	} else {
		for __obf_1cfa18c298665e4b := range *__obf_c6caadad3910b4a2 {
			if __obf_82e5a2bc6d3e754a.Contains(__obf_1cfa18c298665e4b) {
				__obf_bbe2385ee84f0917.Add(__obf_1cfa18c298665e4b)
			}
		}
	}
	return &__obf_bbe2385ee84f0917
}

// 差集
func (__obf_82e5a2bc6d3e754a *Set[T]) Difference(__obf_c6caadad3910b4a2 *Set[T]) *Set[T] {

	__obf_a4782b31520a8d6a := Set[T]{}
	for __obf_1cfa18c298665e4b := range *__obf_82e5a2bc6d3e754a {
		if !__obf_c6caadad3910b4a2.Contains(__obf_1cfa18c298665e4b) {
			__obf_a4782b31520a8d6a.Add(__obf_1cfa18c298665e4b)
		}
	}
	return &__obf_a4782b31520a8d6a
}

// 函数遍历
func (__obf_82e5a2bc6d3e754a *Set[T]) Each(__obf_71ab4c5b41541869 func(any) bool) {
	for __obf_1cfa18c298665e4b := range *__obf_82e5a2bc6d3e754a {
		if __obf_71ab4c5b41541869(__obf_1cfa18c298665e4b) {
			break
		}
	}
}

// 返回string数组
func (__obf_82e5a2bc6d3e754a *Set[T]) StringElem() []string {
	__obf_218c469bce080450 := make([]string, 0, len(*__obf_82e5a2bc6d3e754a))

	for __obf_1cfa18c298665e4b := range *__obf_82e5a2bc6d3e754a {
		__obf_218c469bce080450 = append(__obf_218c469bce080450, fmt.Sprintf("%v", __obf_1cfa18c298665e4b))
	}
	return __obf_218c469bce080450
}
