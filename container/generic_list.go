package __obf_b0bebe5eb45b8ad6

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	__obf_1cab4cec8958fb57,
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_1222f1e21057fa76 *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_341b1174cea927bd *Element[V]) Next() *Element[V] {
	return __obf_341b1174cea927bd.

	// Prev returns the previous list element or nil.
	__obf_1cab4cec8958fb57
}

func (__obf_341b1174cea927bd *Element[V]) Prev() *Element[V] {
	return __obf_341b1174cea927bd.

	// list represents a null terminated (non circular) intrusive doubly linked list.
	// The list is immediately usable after instantiation without the need of a dedicated initialization.
	__obf_1222f1e21057fa76
}

type __obf_04c038c56fdc3d0a[V any] struct {
	__obf_ec657759c74d4eef Element[V] // list head and tail
}

func (__obf_317b1918bcc86816 *__obf_04c038c56fdc3d0a[V]) IsEmpty() bool {
	return __obf_317b1918bcc86816.__obf_ec657759c74d4eef.

	// Front returns the first element of list l or nil if the list is empty.
	__obf_1cab4cec8958fb57 == nil
}

func (__obf_317b1918bcc86816 *__obf_04c038c56fdc3d0a[V]) Front() *Element[V] {
	return __obf_317b1918bcc86816.

	// Back returns the last element of list l or nil if the list is empty.
	__obf_ec657759c74d4eef.__obf_1cab4cec8958fb57
}

func (__obf_317b1918bcc86816 *__obf_04c038c56fdc3d0a[V]) Back() *Element[V] {
	return __obf_317b1918bcc86816.

	// Remove removes e from its list
	__obf_ec657759c74d4eef.__obf_1222f1e21057fa76
}

func (__obf_317b1918bcc86816 *__obf_04c038c56fdc3d0a[V]) Remove(__obf_341b1174cea927bd *Element[V]) {
	if __obf_341b1174cea927bd.__obf_1222f1e21057fa76 == nil {
		__obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1cab4cec8958fb57 = __obf_341b1174cea927bd.__obf_1cab4cec8958fb57
	} else {
		__obf_341b1174cea927bd.__obf_1222f1e21057fa76.__obf_1cab4cec8958fb57 = __obf_341b1174cea927bd.__obf_1cab4cec8958fb57
	}
	if __obf_341b1174cea927bd.__obf_1cab4cec8958fb57 == nil {
		__obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1222f1e21057fa76 = __obf_341b1174cea927bd.__obf_1222f1e21057fa76
	} else {
		__obf_341b1174cea927bd.__obf_1cab4cec8958fb57.

		// avoid memory leaks
		__obf_1222f1e21057fa76 = __obf_341b1174cea927bd.__obf_1222f1e21057fa76
	}
	__obf_341b1174cea927bd.__obf_1cab4cec8958fb57 = nil
	__obf_341b1174cea927bd.__obf_1222f1e21057fa76 = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_317b1918bcc86816 *__obf_04c038c56fdc3d0a[V]) PushFront(__obf_9c15798bcb95be3e string, __obf_72f1369f2f75ff87 V) *Element[V] {
	__obf_341b1174cea927bd := &Element[V]{Key: __obf_9c15798bcb95be3e, Value: __obf_72f1369f2f75ff87}
	if __obf_317b1918bcc86816.__obf_ec657759c74d4eef.
	// It's the first element
	__obf_1cab4cec8958fb57 == nil {
		__obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1cab4cec8958fb57 = __obf_341b1174cea927bd
		__obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1222f1e21057fa76 = __obf_341b1174cea927bd
		return __obf_341b1174cea927bd
	}
	__obf_341b1174cea927bd.__obf_1cab4cec8958fb57 = __obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1cab4cec8958fb57

	// PushBack inserts a new element e with value v at the back of list l and returns e.
	__obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1cab4cec8958fb57.__obf_1222f1e21057fa76 = __obf_341b1174cea927bd
	__obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1cab4cec8958fb57 = __obf_341b1174cea927bd
	return __obf_341b1174cea927bd
}

func (__obf_317b1918bcc86816 *__obf_04c038c56fdc3d0a[V]) PushBack(__obf_9c15798bcb95be3e string, __obf_72f1369f2f75ff87 V) *Element[V] {
	__obf_341b1174cea927bd := &Element[V]{Key: __obf_9c15798bcb95be3e, Value: __obf_72f1369f2f75ff87}
	if __obf_317b1918bcc86816.__obf_ec657759c74d4eef.
	// It's the first element
	__obf_1222f1e21057fa76 == nil {
		__obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1cab4cec8958fb57 = __obf_341b1174cea927bd
		__obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1222f1e21057fa76 = __obf_341b1174cea927bd
		return __obf_341b1174cea927bd
	}
	__obf_341b1174cea927bd.__obf_1222f1e21057fa76 = __obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1222f1e21057fa76
	__obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1222f1e21057fa76.__obf_1cab4cec8958fb57 = __obf_341b1174cea927bd
	__obf_317b1918bcc86816.__obf_ec657759c74d4eef.__obf_1222f1e21057fa76 = __obf_341b1174cea927bd
	return __obf_341b1174cea927bd
}
