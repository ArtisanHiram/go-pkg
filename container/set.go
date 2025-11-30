// 非线程安全set集合
package __obf_9004b47202f9204b

import "fmt"

type Set[T comparable] map[T]__obf_1690e2572402fb23
type __obf_1690e2572402fb23 struct{}

func NewSet[E comparable](__obf_b8a39b8ef99213e5 []E) *Set[E] {
	__obf_1b9b96f861f8239a :=// 获取Set的地址
	Set[E]{}
	// 声明map类型的数据结构
	for _, __obf_7b75a5e4c702d1ce := range __obf_b8a39b8ef99213e5 {
		__obf_1b9b96f861f8239a[// unsafe.Sizeof(void{}) // 结果为0
		__obf_7b75a5e4c702d1ce] = __obf_1690e2572402fb23{}
	}
	return &__obf_1b9b96f861f8239a
}

func (__obf_1b9b96f861f8239a *Set[T]) Add(__obf_c97cf0c2d79c5ab6 T) bool {
	if _, __obf_545047dd3a6c3c8d := (*__obf_1b9b96f861f8239a)[__obf_c97cf0c2d79c5ab6]; __obf_545047dd3a6c3c8d {
		return false
	}

	// unsafe.Sizeof(void{}) // 结果为0
	(*__obf_1b9b96f861f8239a)[__obf_c97cf0c2d79c5ab6] = __obf_1690e2572402fb23{}
	return true
}

func (__obf_1b9b96f861f8239a *Set[T]) Contains(__obf_7b75a5e4c702d1ce T) bool {
	_, __obf_545047dd3a6c3c8d := (*__obf_1b9b96f861f8239a)[__obf_7b75a5e4c702d1ce]
	return __obf_545047dd3a6c3c8d
}

func (__obf_1b9b96f861f8239a *Set[T]) Size() int {
	return len(*__obf_1b9b96f861f8239a)
}

func (__obf_1b9b96f861f8239a *Set[T]) Elem() []T {
	__obf_6bb2e295ddff69c0 := make([]T, 0, len(*__obf_1b9b96f861f8239a))
	for __obf_cc18ca82ac13013e := range *__obf_1b9b96f861f8239a {
		__obf_6bb2e295ddff69c0 = append(__obf_6bb2e295ddff69c0, __obf_cc18ca82ac13013e)
	}
	return __obf_6bb2e295ddff69c0
}

func (__obf_1b9b96f861f8239a *Set[T]) Remove(__obf_c97cf0c2d79c5ab6 T) {
	delete(*__obf_1b9b96f861f8239a, __obf_c97cf0c2d79c5ab6)
}

func (__obf_1b9b96f861f8239a *Set[T]) Clear() {
	*__obf_1b9b96f861f8239a = make(map[T]__obf_1690e2572402fb23)
}

// 相等
func (__obf_1b9b96f861f8239a *Set[T]) Equal(__obf_a34645c64dea2a87 *Set[T]) bool {
	// _ = other.(*Set)
	// 如果两者Size不相等，就不用比较了
	if __obf_1b9b96f861f8239a.Size() != __obf_a34645c64dea2a87.Size() {
		return false
	}
	for __obf_74717af4bf261f66 := range *__obf_1b9b96f861f8239a {
		if !__obf_a34645c64dea2a87.Contains(__obf_74717af4bf261f66) {
			return false
		}
	}
	return true

}

// 子集
func (__obf_1b9b96f861f8239a *Set[T]) IsSubset(__obf_a34645c64dea2a87 *Set[T]) bool {
	// s的size长于other，不用说了
	if __obf_1b9b96f861f8239a.Size() > __obf_a34645c64dea2a87.Size() {
		return false
	}
	// 迭代遍历
	for __obf_355e7c922e433678 := range *__obf_1b9b96f861f8239a {
		if !__obf_a34645c64dea2a87.Contains(__obf_355e7c922e433678) {
			return false
		}
	}
	return true
}

// 并集
func (__obf_992d569fad654d85 *Set[T]) Union(__obf_a34645c64dea2a87 *Set[T]) *Set[T] {
	__obf_a3630fbeee0d0e14 := Set[T]{}
	for __obf_74717af4bf261f66 := range *__obf_992d569fad654d85 {
		__obf_a3630fbeee0d0e14.
			Add(__obf_74717af4bf261f66)
	}
	for __obf_74717af4bf261f66 := range *__obf_a34645c64dea2a87 {
		__obf_a3630fbeee0d0e14.
			Add(__obf_74717af4bf261f66)
	}
	return &__obf_a3630fbeee0d0e14
}

// 交集
func (__obf_1b9b96f861f8239a *Set[T]) Intersect(__obf_a34645c64dea2a87 *Set[T]) *Set[T] {
	__obf_3e73e4c11758ca75 := Set[T]{}
	// loop over smaller set
	if __obf_1b9b96f861f8239a.Size() < __obf_a34645c64dea2a87.Size() {
		for __obf_74717af4bf261f66 := range *__obf_1b9b96f861f8239a {
			if __obf_a34645c64dea2a87.Contains(__obf_74717af4bf261f66) {
				__obf_3e73e4c11758ca75.
					Add(__obf_74717af4bf261f66)
			}
		}
	} else {
		for __obf_74717af4bf261f66 := range *__obf_a34645c64dea2a87 {
			if __obf_1b9b96f861f8239a.Contains(__obf_74717af4bf261f66) {
				__obf_3e73e4c11758ca75.
					Add(__obf_74717af4bf261f66)
			}
		}
	}
	return &__obf_3e73e4c11758ca75
}

// 差集
func (__obf_1b9b96f861f8239a *Set[T]) Difference(__obf_a34645c64dea2a87 *Set[T]) *Set[T] {
	__obf_40c753bd16a3dd5e := Set[T]{}
	for __obf_74717af4bf261f66 := range *__obf_1b9b96f861f8239a {
		if !__obf_a34645c64dea2a87.Contains(__obf_74717af4bf261f66) {
			__obf_40c753bd16a3dd5e.
				Add(__obf_74717af4bf261f66)
		}
	}
	return &__obf_40c753bd16a3dd5e
}

// 函数遍历
func (__obf_1b9b96f861f8239a *Set[T]) Each(__obf_874c0807fd1a46ad func(any) bool) {
	for __obf_74717af4bf261f66 := range *__obf_1b9b96f861f8239a {
		if __obf_874c0807fd1a46ad(__obf_74717af4bf261f66) {
			break
		}
	}
}

// 返回string数组
func (__obf_1b9b96f861f8239a *Set[T]) StringElem() []string {
	__obf_b8a39b8ef99213e5 := make([]string, 0, len(*__obf_1b9b96f861f8239a))

	for __obf_74717af4bf261f66 := range *__obf_1b9b96f861f8239a {
		__obf_b8a39b8ef99213e5 = append(__obf_b8a39b8ef99213e5, fmt.Sprintf("%v", __obf_74717af4bf261f66))
	}
	return __obf_b8a39b8ef99213e5
}
