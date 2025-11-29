package __obf_90a4883a02d0b41c

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	__obf_32b235e6da4cbc13,
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_e3286a9ca021156f *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_40c59de1daa8c7e5 *Element[V]) Next() *Element[V] {
	return __obf_40c59de1daa8c7e5.

	// Prev returns the previous list element or nil.
	__obf_32b235e6da4cbc13
}

func (__obf_40c59de1daa8c7e5 *Element[V]) Prev() *Element[V] {
	return __obf_40c59de1daa8c7e5.

	// list represents a null terminated (non circular) intrusive doubly linked list.
	// The list is immediately usable after instantiation without the need of a dedicated initialization.
	__obf_e3286a9ca021156f
}

type __obf_3947ec1119a8bcb1[V any] struct {
	__obf_46382cb1db4867be Element[V] // list head and tail
}

func (__obf_4725aec49db1b350 *__obf_3947ec1119a8bcb1[V]) IsEmpty() bool {
	return __obf_4725aec49db1b350.__obf_46382cb1db4867be.

	// Front returns the first element of list l or nil if the list is empty.
	__obf_32b235e6da4cbc13 == nil
}

func (__obf_4725aec49db1b350 *__obf_3947ec1119a8bcb1[V]) Front() *Element[V] {
	return __obf_4725aec49db1b350.

	// Back returns the last element of list l or nil if the list is empty.
	__obf_46382cb1db4867be.__obf_32b235e6da4cbc13
}

func (__obf_4725aec49db1b350 *__obf_3947ec1119a8bcb1[V]) Back() *Element[V] {
	return __obf_4725aec49db1b350.

	// Remove removes e from its list
	__obf_46382cb1db4867be.__obf_e3286a9ca021156f
}

func (__obf_4725aec49db1b350 *__obf_3947ec1119a8bcb1[V]) Remove(__obf_40c59de1daa8c7e5 *Element[V]) {
	if __obf_40c59de1daa8c7e5.__obf_e3286a9ca021156f == nil {
		__obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_32b235e6da4cbc13 = __obf_40c59de1daa8c7e5.__obf_32b235e6da4cbc13
	} else {
		__obf_40c59de1daa8c7e5.__obf_e3286a9ca021156f.__obf_32b235e6da4cbc13 = __obf_40c59de1daa8c7e5.__obf_32b235e6da4cbc13
	}
	if __obf_40c59de1daa8c7e5.__obf_32b235e6da4cbc13 == nil {
		__obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_e3286a9ca021156f = __obf_40c59de1daa8c7e5.__obf_e3286a9ca021156f
	} else {
		__obf_40c59de1daa8c7e5.__obf_32b235e6da4cbc13.

		// avoid memory leaks
		__obf_e3286a9ca021156f = __obf_40c59de1daa8c7e5.__obf_e3286a9ca021156f
	}
	__obf_40c59de1daa8c7e5.__obf_32b235e6da4cbc13 = nil
	__obf_40c59de1daa8c7e5.__obf_e3286a9ca021156f = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_4725aec49db1b350 *__obf_3947ec1119a8bcb1[V]) PushFront(__obf_5bba24c1758bbf28 string, __obf_934b4769b75de0d4 V) *Element[V] {
	__obf_40c59de1daa8c7e5 := &Element[V]{Key: __obf_5bba24c1758bbf28, Value: __obf_934b4769b75de0d4}
	if __obf_4725aec49db1b350.__obf_46382cb1db4867be.
	// It's the first element
	__obf_32b235e6da4cbc13 == nil {
		__obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_32b235e6da4cbc13 = __obf_40c59de1daa8c7e5
		__obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_e3286a9ca021156f = __obf_40c59de1daa8c7e5
		return __obf_40c59de1daa8c7e5
	}
	__obf_40c59de1daa8c7e5.__obf_32b235e6da4cbc13 = __obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_32b235e6da4cbc13

	// PushBack inserts a new element e with value v at the back of list l and returns e.
	__obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_32b235e6da4cbc13.__obf_e3286a9ca021156f = __obf_40c59de1daa8c7e5
	__obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_32b235e6da4cbc13 = __obf_40c59de1daa8c7e5
	return __obf_40c59de1daa8c7e5
}

func (__obf_4725aec49db1b350 *__obf_3947ec1119a8bcb1[V]) PushBack(__obf_5bba24c1758bbf28 string, __obf_934b4769b75de0d4 V) *Element[V] {
	__obf_40c59de1daa8c7e5 := &Element[V]{Key: __obf_5bba24c1758bbf28, Value: __obf_934b4769b75de0d4}
	if __obf_4725aec49db1b350.__obf_46382cb1db4867be.
	// It's the first element
	__obf_e3286a9ca021156f == nil {
		__obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_32b235e6da4cbc13 = __obf_40c59de1daa8c7e5
		__obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_e3286a9ca021156f = __obf_40c59de1daa8c7e5
		return __obf_40c59de1daa8c7e5
	}
	__obf_40c59de1daa8c7e5.__obf_e3286a9ca021156f = __obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_e3286a9ca021156f
	__obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_e3286a9ca021156f.__obf_32b235e6da4cbc13 = __obf_40c59de1daa8c7e5
	__obf_4725aec49db1b350.__obf_46382cb1db4867be.__obf_e3286a9ca021156f = __obf_40c59de1daa8c7e5
	return __obf_40c59de1daa8c7e5
}
