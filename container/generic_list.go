package __obf_76599ab96a208083

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_b7a3ba8614d8f9d8, __obf_6b94d13992080dfb *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_7eb1599d8c06f3ff *Element[V]) Next() *Element[V] {
	return __obf_7eb1599d8c06f3ff.__obf_b7a3ba8614d8f9d8
}

// Prev returns the previous list element or nil.
func (__obf_7eb1599d8c06f3ff *Element[V]) Prev() *Element[V] {
	return __obf_7eb1599d8c06f3ff.__obf_6b94d13992080dfb
}

// list represents a null terminated (non circular) intrusive doubly linked list.
// The list is immediately usable after instantiation without the need of a dedicated initialization.
type __obf_3bdbb0fb35d34874[V any] struct {
	__obf_33946b56d5c4af15 Element[V] // list head and tail
}

func (__obf_b9c6e3d4b172ac7a *__obf_3bdbb0fb35d34874[V]) IsEmpty() bool {
	return __obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_b7a3ba8614d8f9d8 == nil
}

// Front returns the first element of list l or nil if the list is empty.
func (__obf_b9c6e3d4b172ac7a *__obf_3bdbb0fb35d34874[V]) Front() *Element[V] {
	return __obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_b7a3ba8614d8f9d8
}

// Back returns the last element of list l or nil if the list is empty.
func (__obf_b9c6e3d4b172ac7a *__obf_3bdbb0fb35d34874[V]) Back() *Element[V] {
	return __obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_6b94d13992080dfb
}

// Remove removes e from its list
func (__obf_b9c6e3d4b172ac7a *__obf_3bdbb0fb35d34874[V]) Remove(__obf_7eb1599d8c06f3ff *Element[V]) {
	if __obf_7eb1599d8c06f3ff.__obf_6b94d13992080dfb == nil {
		__obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_b7a3ba8614d8f9d8 = __obf_7eb1599d8c06f3ff.__obf_b7a3ba8614d8f9d8
	} else {
		__obf_7eb1599d8c06f3ff.__obf_6b94d13992080dfb.__obf_b7a3ba8614d8f9d8 = __obf_7eb1599d8c06f3ff.__obf_b7a3ba8614d8f9d8
	}
	if __obf_7eb1599d8c06f3ff.__obf_b7a3ba8614d8f9d8 == nil {
		__obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_6b94d13992080dfb = __obf_7eb1599d8c06f3ff.__obf_6b94d13992080dfb
	} else {
		__obf_7eb1599d8c06f3ff.__obf_b7a3ba8614d8f9d8.__obf_6b94d13992080dfb = __obf_7eb1599d8c06f3ff.__obf_6b94d13992080dfb
	}
	__obf_7eb1599d8c06f3ff.__obf_b7a3ba8614d8f9d8 = nil // avoid memory leaks
	__obf_7eb1599d8c06f3ff.__obf_6b94d13992080dfb = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_b9c6e3d4b172ac7a *__obf_3bdbb0fb35d34874[V]) PushFront(__obf_4368715d0a6d4f0b string, __obf_209abc4c4955f6d3 V) *Element[V] {
	__obf_7eb1599d8c06f3ff := &Element[V]{Key: __obf_4368715d0a6d4f0b, Value: __obf_209abc4c4955f6d3}
	if __obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_b7a3ba8614d8f9d8 == nil {
		// It's the first element
		__obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_b7a3ba8614d8f9d8 = __obf_7eb1599d8c06f3ff
		__obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_6b94d13992080dfb = __obf_7eb1599d8c06f3ff
		return __obf_7eb1599d8c06f3ff
	}

	__obf_7eb1599d8c06f3ff.__obf_b7a3ba8614d8f9d8 = __obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_b7a3ba8614d8f9d8
	__obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_b7a3ba8614d8f9d8.__obf_6b94d13992080dfb = __obf_7eb1599d8c06f3ff
	__obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_b7a3ba8614d8f9d8 = __obf_7eb1599d8c06f3ff
	return __obf_7eb1599d8c06f3ff
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (__obf_b9c6e3d4b172ac7a *__obf_3bdbb0fb35d34874[V]) PushBack(__obf_4368715d0a6d4f0b string, __obf_209abc4c4955f6d3 V) *Element[V] {
	__obf_7eb1599d8c06f3ff := &Element[V]{Key: __obf_4368715d0a6d4f0b, Value: __obf_209abc4c4955f6d3}
	if __obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_6b94d13992080dfb == nil {
		// It's the first element
		__obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_b7a3ba8614d8f9d8 = __obf_7eb1599d8c06f3ff
		__obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_6b94d13992080dfb = __obf_7eb1599d8c06f3ff
		return __obf_7eb1599d8c06f3ff
	}

	__obf_7eb1599d8c06f3ff.__obf_6b94d13992080dfb = __obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_6b94d13992080dfb
	__obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_6b94d13992080dfb.__obf_b7a3ba8614d8f9d8 = __obf_7eb1599d8c06f3ff
	__obf_b9c6e3d4b172ac7a.__obf_33946b56d5c4af15.__obf_6b94d13992080dfb = __obf_7eb1599d8c06f3ff
	return __obf_7eb1599d8c06f3ff
}
