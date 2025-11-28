package __obf_4f13ac5aa043b5ce

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_ab34c2f73a61331f, __obf_5c809ff977b880f2 *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_91355e6634178919 *Element[V]) Next() *Element[V] {
	return __obf_91355e6634178919.__obf_ab34c2f73a61331f
}

// Prev returns the previous list element or nil.
func (__obf_91355e6634178919 *Element[V]) Prev() *Element[V] {
	return __obf_91355e6634178919.__obf_5c809ff977b880f2
}

// list represents a null terminated (non circular) intrusive doubly linked list.
// The list is immediately usable after instantiation without the need of a dedicated initialization.
type __obf_25770552d978c33e[V any] struct {
	__obf_1f75909e0fb05321 Element[V] // list head and tail
}

func (__obf_ec251f9f79f12c3b *__obf_25770552d978c33e[V]) IsEmpty() bool {
	return __obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_ab34c2f73a61331f == nil
}

// Front returns the first element of list l or nil if the list is empty.
func (__obf_ec251f9f79f12c3b *__obf_25770552d978c33e[V]) Front() *Element[V] {
	return __obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_ab34c2f73a61331f
}

// Back returns the last element of list l or nil if the list is empty.
func (__obf_ec251f9f79f12c3b *__obf_25770552d978c33e[V]) Back() *Element[V] {
	return __obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_5c809ff977b880f2
}

// Remove removes e from its list
func (__obf_ec251f9f79f12c3b *__obf_25770552d978c33e[V]) Remove(__obf_91355e6634178919 *Element[V]) {
	if __obf_91355e6634178919.__obf_5c809ff977b880f2 == nil {
		__obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_ab34c2f73a61331f = __obf_91355e6634178919.__obf_ab34c2f73a61331f
	} else {
		__obf_91355e6634178919.__obf_5c809ff977b880f2.__obf_ab34c2f73a61331f = __obf_91355e6634178919.__obf_ab34c2f73a61331f
	}
	if __obf_91355e6634178919.__obf_ab34c2f73a61331f == nil {
		__obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_5c809ff977b880f2 = __obf_91355e6634178919.__obf_5c809ff977b880f2
	} else {
		__obf_91355e6634178919.__obf_ab34c2f73a61331f.__obf_5c809ff977b880f2 = __obf_91355e6634178919.__obf_5c809ff977b880f2
	}
	__obf_91355e6634178919.__obf_ab34c2f73a61331f = nil // avoid memory leaks
	__obf_91355e6634178919.__obf_5c809ff977b880f2 = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_ec251f9f79f12c3b *__obf_25770552d978c33e[V]) PushFront(__obf_797707a3dac0ebb7 string, __obf_77dcf3eed779629a V) *Element[V] {
	__obf_91355e6634178919 := &Element[V]{Key: __obf_797707a3dac0ebb7, Value: __obf_77dcf3eed779629a}
	if __obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_ab34c2f73a61331f == nil {
		// It's the first element
		__obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_ab34c2f73a61331f = __obf_91355e6634178919
		__obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_5c809ff977b880f2 = __obf_91355e6634178919
		return __obf_91355e6634178919
	}

	__obf_91355e6634178919.__obf_ab34c2f73a61331f = __obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_ab34c2f73a61331f
	__obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_ab34c2f73a61331f.__obf_5c809ff977b880f2 = __obf_91355e6634178919
	__obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_ab34c2f73a61331f = __obf_91355e6634178919
	return __obf_91355e6634178919
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (__obf_ec251f9f79f12c3b *__obf_25770552d978c33e[V]) PushBack(__obf_797707a3dac0ebb7 string, __obf_77dcf3eed779629a V) *Element[V] {
	__obf_91355e6634178919 := &Element[V]{Key: __obf_797707a3dac0ebb7, Value: __obf_77dcf3eed779629a}
	if __obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_5c809ff977b880f2 == nil {
		// It's the first element
		__obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_ab34c2f73a61331f = __obf_91355e6634178919
		__obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_5c809ff977b880f2 = __obf_91355e6634178919
		return __obf_91355e6634178919
	}

	__obf_91355e6634178919.__obf_5c809ff977b880f2 = __obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_5c809ff977b880f2
	__obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_5c809ff977b880f2.__obf_ab34c2f73a61331f = __obf_91355e6634178919
	__obf_ec251f9f79f12c3b.__obf_1f75909e0fb05321.__obf_5c809ff977b880f2 = __obf_91355e6634178919
	return __obf_91355e6634178919
}
