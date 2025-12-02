package __obf_4a16ef421fb74992

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	__obf_2c113f310b7cc228,
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_7b497de6e96482de *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_0c42fea19843f159 *Element[V]) Next() *Element[V] {
	return __obf_0c42fea19843f159.

	// Prev returns the previous list element or nil.
	__obf_2c113f310b7cc228
}

func (__obf_0c42fea19843f159 *Element[V]) Prev() *Element[V] {
	return __obf_0c42fea19843f159.

	// list represents a null terminated (non circular) intrusive doubly linked list.
	// The list is immediately usable after instantiation without the need of a dedicated initialization.
	__obf_7b497de6e96482de
}

type __obf_333361aa59ab2d50[V any] struct {
	__obf_7ddae20cea12111b Element[V] // list head and tail
}

func (__obf_3ce53f1812062902 *__obf_333361aa59ab2d50[V]) IsEmpty() bool {
	return __obf_3ce53f1812062902.__obf_7ddae20cea12111b.

	// Front returns the first element of list l or nil if the list is empty.
	__obf_2c113f310b7cc228 == nil
}

func (__obf_3ce53f1812062902 *__obf_333361aa59ab2d50[V]) Front() *Element[V] {
	return __obf_3ce53f1812062902.

	// Back returns the last element of list l or nil if the list is empty.
	__obf_7ddae20cea12111b.__obf_2c113f310b7cc228
}

func (__obf_3ce53f1812062902 *__obf_333361aa59ab2d50[V]) Back() *Element[V] {
	return __obf_3ce53f1812062902.

	// Remove removes e from its list
	__obf_7ddae20cea12111b.__obf_7b497de6e96482de
}

func (__obf_3ce53f1812062902 *__obf_333361aa59ab2d50[V]) Remove(__obf_0c42fea19843f159 *Element[V]) {
	if __obf_0c42fea19843f159.__obf_7b497de6e96482de == nil {
		__obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_2c113f310b7cc228 = __obf_0c42fea19843f159.__obf_2c113f310b7cc228
	} else {
		__obf_0c42fea19843f159.__obf_7b497de6e96482de.__obf_2c113f310b7cc228 = __obf_0c42fea19843f159.__obf_2c113f310b7cc228
	}
	if __obf_0c42fea19843f159.__obf_2c113f310b7cc228 == nil {
		__obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_7b497de6e96482de = __obf_0c42fea19843f159.__obf_7b497de6e96482de
	} else {
		__obf_0c42fea19843f159.__obf_2c113f310b7cc228.

		// avoid memory leaks
		__obf_7b497de6e96482de = __obf_0c42fea19843f159.__obf_7b497de6e96482de
	}
	__obf_0c42fea19843f159.__obf_2c113f310b7cc228 = nil
	__obf_0c42fea19843f159.__obf_7b497de6e96482de = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_3ce53f1812062902 *__obf_333361aa59ab2d50[V]) PushFront(__obf_d1a1f98ae0fb119e string, __obf_06ad079a003b5526 V) *Element[V] {
	__obf_0c42fea19843f159 := &Element[V]{Key: __obf_d1a1f98ae0fb119e, Value: __obf_06ad079a003b5526}
	if __obf_3ce53f1812062902.__obf_7ddae20cea12111b.
	// It's the first element
	__obf_2c113f310b7cc228 == nil {
		__obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_2c113f310b7cc228 = __obf_0c42fea19843f159
		__obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_7b497de6e96482de = __obf_0c42fea19843f159
		return __obf_0c42fea19843f159
	}
	__obf_0c42fea19843f159.__obf_2c113f310b7cc228 = __obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_2c113f310b7cc228

	// PushBack inserts a new element e with value v at the back of list l and returns e.
	__obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_2c113f310b7cc228.__obf_7b497de6e96482de = __obf_0c42fea19843f159
	__obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_2c113f310b7cc228 = __obf_0c42fea19843f159
	return __obf_0c42fea19843f159
}

func (__obf_3ce53f1812062902 *__obf_333361aa59ab2d50[V]) PushBack(__obf_d1a1f98ae0fb119e string, __obf_06ad079a003b5526 V) *Element[V] {
	__obf_0c42fea19843f159 := &Element[V]{Key: __obf_d1a1f98ae0fb119e, Value: __obf_06ad079a003b5526}
	if __obf_3ce53f1812062902.__obf_7ddae20cea12111b.
	// It's the first element
	__obf_7b497de6e96482de == nil {
		__obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_2c113f310b7cc228 = __obf_0c42fea19843f159
		__obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_7b497de6e96482de = __obf_0c42fea19843f159
		return __obf_0c42fea19843f159
	}
	__obf_0c42fea19843f159.__obf_7b497de6e96482de = __obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_7b497de6e96482de
	__obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_7b497de6e96482de.__obf_2c113f310b7cc228 = __obf_0c42fea19843f159
	__obf_3ce53f1812062902.__obf_7ddae20cea12111b.__obf_7b497de6e96482de = __obf_0c42fea19843f159
	return __obf_0c42fea19843f159
}
