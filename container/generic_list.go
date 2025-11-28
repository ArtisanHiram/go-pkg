package __obf_62eba4024f8fa381

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_a88f24b4f460aa9c, __obf_fc44d56dec1421d9 *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_d038a85f302379f5 *Element[V]) Next() *Element[V] {
	return __obf_d038a85f302379f5.__obf_a88f24b4f460aa9c
}

// Prev returns the previous list element or nil.
func (__obf_d038a85f302379f5 *Element[V]) Prev() *Element[V] {
	return __obf_d038a85f302379f5.__obf_fc44d56dec1421d9
}

// list represents a null terminated (non circular) intrusive doubly linked list.
// The list is immediately usable after instantiation without the need of a dedicated initialization.
type __obf_40b36c6e023d30bc[V any] struct {
	__obf_6e004345cda28879 Element[V] // list head and tail
}

func (__obf_252798d37d9ac024 *__obf_40b36c6e023d30bc[V]) IsEmpty() bool {
	return __obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_a88f24b4f460aa9c == nil
}

// Front returns the first element of list l or nil if the list is empty.
func (__obf_252798d37d9ac024 *__obf_40b36c6e023d30bc[V]) Front() *Element[V] {
	return __obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_a88f24b4f460aa9c
}

// Back returns the last element of list l or nil if the list is empty.
func (__obf_252798d37d9ac024 *__obf_40b36c6e023d30bc[V]) Back() *Element[V] {
	return __obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_fc44d56dec1421d9
}

// Remove removes e from its list
func (__obf_252798d37d9ac024 *__obf_40b36c6e023d30bc[V]) Remove(__obf_d038a85f302379f5 *Element[V]) {
	if __obf_d038a85f302379f5.__obf_fc44d56dec1421d9 == nil {
		__obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_a88f24b4f460aa9c = __obf_d038a85f302379f5.__obf_a88f24b4f460aa9c
	} else {
		__obf_d038a85f302379f5.__obf_fc44d56dec1421d9.__obf_a88f24b4f460aa9c = __obf_d038a85f302379f5.__obf_a88f24b4f460aa9c
	}
	if __obf_d038a85f302379f5.__obf_a88f24b4f460aa9c == nil {
		__obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_fc44d56dec1421d9 = __obf_d038a85f302379f5.__obf_fc44d56dec1421d9
	} else {
		__obf_d038a85f302379f5.__obf_a88f24b4f460aa9c.__obf_fc44d56dec1421d9 = __obf_d038a85f302379f5.__obf_fc44d56dec1421d9
	}
	__obf_d038a85f302379f5.__obf_a88f24b4f460aa9c = nil // avoid memory leaks
	__obf_d038a85f302379f5.__obf_fc44d56dec1421d9 = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_252798d37d9ac024 *__obf_40b36c6e023d30bc[V]) PushFront(__obf_df070ab4c712506c string, __obf_dbc685b0ad464e4d V) *Element[V] {
	__obf_d038a85f302379f5 := &Element[V]{Key: __obf_df070ab4c712506c, Value: __obf_dbc685b0ad464e4d}
	if __obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_a88f24b4f460aa9c == nil {
		// It's the first element
		__obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_a88f24b4f460aa9c = __obf_d038a85f302379f5
		__obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_fc44d56dec1421d9 = __obf_d038a85f302379f5
		return __obf_d038a85f302379f5
	}

	__obf_d038a85f302379f5.__obf_a88f24b4f460aa9c = __obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_a88f24b4f460aa9c
	__obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_a88f24b4f460aa9c.__obf_fc44d56dec1421d9 = __obf_d038a85f302379f5
	__obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_a88f24b4f460aa9c = __obf_d038a85f302379f5
	return __obf_d038a85f302379f5
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (__obf_252798d37d9ac024 *__obf_40b36c6e023d30bc[V]) PushBack(__obf_df070ab4c712506c string, __obf_dbc685b0ad464e4d V) *Element[V] {
	__obf_d038a85f302379f5 := &Element[V]{Key: __obf_df070ab4c712506c, Value: __obf_dbc685b0ad464e4d}
	if __obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_fc44d56dec1421d9 == nil {
		// It's the first element
		__obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_a88f24b4f460aa9c = __obf_d038a85f302379f5
		__obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_fc44d56dec1421d9 = __obf_d038a85f302379f5
		return __obf_d038a85f302379f5
	}

	__obf_d038a85f302379f5.__obf_fc44d56dec1421d9 = __obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_fc44d56dec1421d9
	__obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_fc44d56dec1421d9.__obf_a88f24b4f460aa9c = __obf_d038a85f302379f5
	__obf_252798d37d9ac024.__obf_6e004345cda28879.__obf_fc44d56dec1421d9 = __obf_d038a85f302379f5
	return __obf_d038a85f302379f5
}
