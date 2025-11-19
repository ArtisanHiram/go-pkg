package __obf_9861fa13140c30a3

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_9c578e473cd91a70, __obf_9e64c21ddd522887 *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_bd79c4ec5cbfc921 *Element[V]) Next() *Element[V] {
	return __obf_bd79c4ec5cbfc921.__obf_9c578e473cd91a70
}

// Prev returns the previous list element or nil.
func (__obf_bd79c4ec5cbfc921 *Element[V]) Prev() *Element[V] {
	return __obf_bd79c4ec5cbfc921.__obf_9e64c21ddd522887
}

// list represents a null terminated (non circular) intrusive doubly linked list.
// The list is immediately usable after instantiation without the need of a dedicated initialization.
type __obf_9da43f38a5013de2[V any] struct {
	__obf_272bfa9ba362b63a Element[V] // list head and tail
}

func (__obf_e6e08b693d421fab *__obf_9da43f38a5013de2[V]) IsEmpty() bool {
	return __obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9c578e473cd91a70 == nil
}

// Front returns the first element of list l or nil if the list is empty.
func (__obf_e6e08b693d421fab *__obf_9da43f38a5013de2[V]) Front() *Element[V] {
	return __obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9c578e473cd91a70
}

// Back returns the last element of list l or nil if the list is empty.
func (__obf_e6e08b693d421fab *__obf_9da43f38a5013de2[V]) Back() *Element[V] {
	return __obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9e64c21ddd522887
}

// Remove removes e from its list
func (__obf_e6e08b693d421fab *__obf_9da43f38a5013de2[V]) Remove(__obf_bd79c4ec5cbfc921 *Element[V]) {
	if __obf_bd79c4ec5cbfc921.__obf_9e64c21ddd522887 == nil {
		__obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9c578e473cd91a70 = __obf_bd79c4ec5cbfc921.__obf_9c578e473cd91a70
	} else {
		__obf_bd79c4ec5cbfc921.__obf_9e64c21ddd522887.__obf_9c578e473cd91a70 = __obf_bd79c4ec5cbfc921.__obf_9c578e473cd91a70
	}
	if __obf_bd79c4ec5cbfc921.__obf_9c578e473cd91a70 == nil {
		__obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9e64c21ddd522887 = __obf_bd79c4ec5cbfc921.__obf_9e64c21ddd522887
	} else {
		__obf_bd79c4ec5cbfc921.__obf_9c578e473cd91a70.__obf_9e64c21ddd522887 = __obf_bd79c4ec5cbfc921.__obf_9e64c21ddd522887
	}
	__obf_bd79c4ec5cbfc921.__obf_9c578e473cd91a70 = nil // avoid memory leaks
	__obf_bd79c4ec5cbfc921.__obf_9e64c21ddd522887 = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_e6e08b693d421fab *__obf_9da43f38a5013de2[V]) PushFront(__obf_43194ec765d86867 string, __obf_3966a3e46995974f V) *Element[V] {
	__obf_bd79c4ec5cbfc921 := &Element[V]{Key: __obf_43194ec765d86867, Value: __obf_3966a3e46995974f}
	if __obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9c578e473cd91a70 == nil {
		// It's the first element
		__obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9c578e473cd91a70 = __obf_bd79c4ec5cbfc921
		__obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9e64c21ddd522887 = __obf_bd79c4ec5cbfc921
		return __obf_bd79c4ec5cbfc921
	}

	__obf_bd79c4ec5cbfc921.__obf_9c578e473cd91a70 = __obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9c578e473cd91a70
	__obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9c578e473cd91a70.__obf_9e64c21ddd522887 = __obf_bd79c4ec5cbfc921
	__obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9c578e473cd91a70 = __obf_bd79c4ec5cbfc921
	return __obf_bd79c4ec5cbfc921
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (__obf_e6e08b693d421fab *__obf_9da43f38a5013de2[V]) PushBack(__obf_43194ec765d86867 string, __obf_3966a3e46995974f V) *Element[V] {
	__obf_bd79c4ec5cbfc921 := &Element[V]{Key: __obf_43194ec765d86867, Value: __obf_3966a3e46995974f}
	if __obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9e64c21ddd522887 == nil {
		// It's the first element
		__obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9c578e473cd91a70 = __obf_bd79c4ec5cbfc921
		__obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9e64c21ddd522887 = __obf_bd79c4ec5cbfc921
		return __obf_bd79c4ec5cbfc921
	}

	__obf_bd79c4ec5cbfc921.__obf_9e64c21ddd522887 = __obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9e64c21ddd522887
	__obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9e64c21ddd522887.__obf_9c578e473cd91a70 = __obf_bd79c4ec5cbfc921
	__obf_e6e08b693d421fab.__obf_272bfa9ba362b63a.__obf_9e64c21ddd522887 = __obf_bd79c4ec5cbfc921
	return __obf_bd79c4ec5cbfc921
}
