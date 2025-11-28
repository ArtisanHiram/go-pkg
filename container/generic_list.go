package __obf_af42fb6cde2beed6

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_3c691a3fec1b7fd4, __obf_29579fcd202167dd *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_49d94ed18268f4d8 *Element[V]) Next() *Element[V] {
	return __obf_49d94ed18268f4d8.__obf_3c691a3fec1b7fd4
}

// Prev returns the previous list element or nil.
func (__obf_49d94ed18268f4d8 *Element[V]) Prev() *Element[V] {
	return __obf_49d94ed18268f4d8.__obf_29579fcd202167dd
}

// list represents a null terminated (non circular) intrusive doubly linked list.
// The list is immediately usable after instantiation without the need of a dedicated initialization.
type __obf_543c4aff1fda48f3[V any] struct {
	__obf_08879f490b6b8864 Element[V] // list head and tail
}

func (__obf_b56a8243feef1dee *__obf_543c4aff1fda48f3[V]) IsEmpty() bool {
	return __obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_3c691a3fec1b7fd4 == nil
}

// Front returns the first element of list l or nil if the list is empty.
func (__obf_b56a8243feef1dee *__obf_543c4aff1fda48f3[V]) Front() *Element[V] {
	return __obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_3c691a3fec1b7fd4
}

// Back returns the last element of list l or nil if the list is empty.
func (__obf_b56a8243feef1dee *__obf_543c4aff1fda48f3[V]) Back() *Element[V] {
	return __obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_29579fcd202167dd
}

// Remove removes e from its list
func (__obf_b56a8243feef1dee *__obf_543c4aff1fda48f3[V]) Remove(__obf_49d94ed18268f4d8 *Element[V]) {
	if __obf_49d94ed18268f4d8.__obf_29579fcd202167dd == nil {
		__obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_3c691a3fec1b7fd4 = __obf_49d94ed18268f4d8.__obf_3c691a3fec1b7fd4
	} else {
		__obf_49d94ed18268f4d8.__obf_29579fcd202167dd.__obf_3c691a3fec1b7fd4 = __obf_49d94ed18268f4d8.__obf_3c691a3fec1b7fd4
	}
	if __obf_49d94ed18268f4d8.__obf_3c691a3fec1b7fd4 == nil {
		__obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_29579fcd202167dd = __obf_49d94ed18268f4d8.__obf_29579fcd202167dd
	} else {
		__obf_49d94ed18268f4d8.__obf_3c691a3fec1b7fd4.__obf_29579fcd202167dd = __obf_49d94ed18268f4d8.__obf_29579fcd202167dd
	}
	__obf_49d94ed18268f4d8.__obf_3c691a3fec1b7fd4 = nil // avoid memory leaks
	__obf_49d94ed18268f4d8.__obf_29579fcd202167dd = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_b56a8243feef1dee *__obf_543c4aff1fda48f3[V]) PushFront(__obf_55cde42f6d47c5be string, __obf_9ed1ee5d2b51057a V) *Element[V] {
	__obf_49d94ed18268f4d8 := &Element[V]{Key: __obf_55cde42f6d47c5be, Value: __obf_9ed1ee5d2b51057a}
	if __obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_3c691a3fec1b7fd4 == nil {
		// It's the first element
		__obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_3c691a3fec1b7fd4 = __obf_49d94ed18268f4d8
		__obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_29579fcd202167dd = __obf_49d94ed18268f4d8
		return __obf_49d94ed18268f4d8
	}

	__obf_49d94ed18268f4d8.__obf_3c691a3fec1b7fd4 = __obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_3c691a3fec1b7fd4
	__obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_3c691a3fec1b7fd4.__obf_29579fcd202167dd = __obf_49d94ed18268f4d8
	__obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_3c691a3fec1b7fd4 = __obf_49d94ed18268f4d8
	return __obf_49d94ed18268f4d8
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (__obf_b56a8243feef1dee *__obf_543c4aff1fda48f3[V]) PushBack(__obf_55cde42f6d47c5be string, __obf_9ed1ee5d2b51057a V) *Element[V] {
	__obf_49d94ed18268f4d8 := &Element[V]{Key: __obf_55cde42f6d47c5be, Value: __obf_9ed1ee5d2b51057a}
	if __obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_29579fcd202167dd == nil {
		// It's the first element
		__obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_3c691a3fec1b7fd4 = __obf_49d94ed18268f4d8
		__obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_29579fcd202167dd = __obf_49d94ed18268f4d8
		return __obf_49d94ed18268f4d8
	}

	__obf_49d94ed18268f4d8.__obf_29579fcd202167dd = __obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_29579fcd202167dd
	__obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_29579fcd202167dd.__obf_3c691a3fec1b7fd4 = __obf_49d94ed18268f4d8
	__obf_b56a8243feef1dee.__obf_08879f490b6b8864.__obf_29579fcd202167dd = __obf_49d94ed18268f4d8
	return __obf_49d94ed18268f4d8
}
