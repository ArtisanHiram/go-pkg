package __obf_9004b47202f9204b

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	__obf_d5e61e639158b14b,
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_fabc805cdf3486a7 *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_5a4873a915966668 *Element[V]) Next() *Element[V] {
	return __obf_5a4873a915966668.

	// Prev returns the previous list element or nil.
	__obf_d5e61e639158b14b
}

func (__obf_5a4873a915966668 *Element[V]) Prev() *Element[V] {
	return __obf_5a4873a915966668.

	// list represents a null terminated (non circular) intrusive doubly linked list.
	// The list is immediately usable after instantiation without the need of a dedicated initialization.
	__obf_fabc805cdf3486a7
}

type __obf_72ca0e09d2929881[V any] struct {
	__obf_59e90690d56f2236 Element[V] // list head and tail
}

func (__obf_99b3eb7c1d62277d *__obf_72ca0e09d2929881[V]) IsEmpty() bool {
	return __obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.

	// Front returns the first element of list l or nil if the list is empty.
	__obf_d5e61e639158b14b == nil
}

func (__obf_99b3eb7c1d62277d *__obf_72ca0e09d2929881[V]) Front() *Element[V] {
	return __obf_99b3eb7c1d62277d.

	// Back returns the last element of list l or nil if the list is empty.
	__obf_59e90690d56f2236.__obf_d5e61e639158b14b
}

func (__obf_99b3eb7c1d62277d *__obf_72ca0e09d2929881[V]) Back() *Element[V] {
	return __obf_99b3eb7c1d62277d.

	// Remove removes e from its list
	__obf_59e90690d56f2236.__obf_fabc805cdf3486a7
}

func (__obf_99b3eb7c1d62277d *__obf_72ca0e09d2929881[V]) Remove(__obf_5a4873a915966668 *Element[V]) {
	if __obf_5a4873a915966668.__obf_fabc805cdf3486a7 == nil {
		__obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_d5e61e639158b14b = __obf_5a4873a915966668.__obf_d5e61e639158b14b
	} else {
		__obf_5a4873a915966668.__obf_fabc805cdf3486a7.__obf_d5e61e639158b14b = __obf_5a4873a915966668.__obf_d5e61e639158b14b
	}
	if __obf_5a4873a915966668.__obf_d5e61e639158b14b == nil {
		__obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_fabc805cdf3486a7 = __obf_5a4873a915966668.__obf_fabc805cdf3486a7
	} else {
		__obf_5a4873a915966668.__obf_d5e61e639158b14b.

		// avoid memory leaks
		__obf_fabc805cdf3486a7 = __obf_5a4873a915966668.__obf_fabc805cdf3486a7
	}
	__obf_5a4873a915966668.__obf_d5e61e639158b14b = nil
	__obf_5a4873a915966668.__obf_fabc805cdf3486a7 = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_99b3eb7c1d62277d *__obf_72ca0e09d2929881[V]) PushFront(__obf_355e7c922e433678 string, __obf_7269ae284cf88b6b V) *Element[V] {
	__obf_5a4873a915966668 := &Element[V]{Key: __obf_355e7c922e433678, Value: __obf_7269ae284cf88b6b}
	if __obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.
	// It's the first element
	__obf_d5e61e639158b14b == nil {
		__obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_d5e61e639158b14b = __obf_5a4873a915966668
		__obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_fabc805cdf3486a7 = __obf_5a4873a915966668
		return __obf_5a4873a915966668
	}
	__obf_5a4873a915966668.__obf_d5e61e639158b14b = __obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_d5e61e639158b14b

	// PushBack inserts a new element e with value v at the back of list l and returns e.
	__obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_d5e61e639158b14b.__obf_fabc805cdf3486a7 = __obf_5a4873a915966668
	__obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_d5e61e639158b14b = __obf_5a4873a915966668
	return __obf_5a4873a915966668
}

func (__obf_99b3eb7c1d62277d *__obf_72ca0e09d2929881[V]) PushBack(__obf_355e7c922e433678 string, __obf_7269ae284cf88b6b V) *Element[V] {
	__obf_5a4873a915966668 := &Element[V]{Key: __obf_355e7c922e433678, Value: __obf_7269ae284cf88b6b}
	if __obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.
	// It's the first element
	__obf_fabc805cdf3486a7 == nil {
		__obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_d5e61e639158b14b = __obf_5a4873a915966668
		__obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_fabc805cdf3486a7 = __obf_5a4873a915966668
		return __obf_5a4873a915966668
	}
	__obf_5a4873a915966668.__obf_fabc805cdf3486a7 = __obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_fabc805cdf3486a7
	__obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_fabc805cdf3486a7.__obf_d5e61e639158b14b = __obf_5a4873a915966668
	__obf_99b3eb7c1d62277d.__obf_59e90690d56f2236.__obf_fabc805cdf3486a7 = __obf_5a4873a915966668
	return __obf_5a4873a915966668
}
