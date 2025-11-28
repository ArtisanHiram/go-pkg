package __obf_1fda7fbdeda52f1e

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_3f38a216fb054660, __obf_475810ae089ac9c9 *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_673d5726ca470196 *Element[V]) Next() *Element[V] {
	return __obf_673d5726ca470196.__obf_3f38a216fb054660
}

// Prev returns the previous list element or nil.
func (__obf_673d5726ca470196 *Element[V]) Prev() *Element[V] {
	return __obf_673d5726ca470196.__obf_475810ae089ac9c9
}

// list represents a null terminated (non circular) intrusive doubly linked list.
// The list is immediately usable after instantiation without the need of a dedicated initialization.
type __obf_5275791a072a11d1[V any] struct {
	__obf_549d995cf5b98905 Element[V] // list head and tail
}

func (__obf_3f548a299e41833e *__obf_5275791a072a11d1[V]) IsEmpty() bool {
	return __obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_3f38a216fb054660 == nil
}

// Front returns the first element of list l or nil if the list is empty.
func (__obf_3f548a299e41833e *__obf_5275791a072a11d1[V]) Front() *Element[V] {
	return __obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_3f38a216fb054660
}

// Back returns the last element of list l or nil if the list is empty.
func (__obf_3f548a299e41833e *__obf_5275791a072a11d1[V]) Back() *Element[V] {
	return __obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_475810ae089ac9c9
}

// Remove removes e from its list
func (__obf_3f548a299e41833e *__obf_5275791a072a11d1[V]) Remove(__obf_673d5726ca470196 *Element[V]) {
	if __obf_673d5726ca470196.__obf_475810ae089ac9c9 == nil {
		__obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_3f38a216fb054660 = __obf_673d5726ca470196.__obf_3f38a216fb054660
	} else {
		__obf_673d5726ca470196.__obf_475810ae089ac9c9.__obf_3f38a216fb054660 = __obf_673d5726ca470196.__obf_3f38a216fb054660
	}
	if __obf_673d5726ca470196.__obf_3f38a216fb054660 == nil {
		__obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_475810ae089ac9c9 = __obf_673d5726ca470196.__obf_475810ae089ac9c9
	} else {
		__obf_673d5726ca470196.__obf_3f38a216fb054660.__obf_475810ae089ac9c9 = __obf_673d5726ca470196.__obf_475810ae089ac9c9
	}
	__obf_673d5726ca470196.__obf_3f38a216fb054660 = nil // avoid memory leaks
	__obf_673d5726ca470196.__obf_475810ae089ac9c9 = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_3f548a299e41833e *__obf_5275791a072a11d1[V]) PushFront(__obf_95c4c09a56c1a070 string, __obf_6d18c47092a65f53 V) *Element[V] {
	__obf_673d5726ca470196 := &Element[V]{Key: __obf_95c4c09a56c1a070, Value: __obf_6d18c47092a65f53}
	if __obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_3f38a216fb054660 == nil {
		// It's the first element
		__obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_3f38a216fb054660 = __obf_673d5726ca470196
		__obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_475810ae089ac9c9 = __obf_673d5726ca470196
		return __obf_673d5726ca470196
	}

	__obf_673d5726ca470196.__obf_3f38a216fb054660 = __obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_3f38a216fb054660
	__obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_3f38a216fb054660.__obf_475810ae089ac9c9 = __obf_673d5726ca470196
	__obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_3f38a216fb054660 = __obf_673d5726ca470196
	return __obf_673d5726ca470196
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (__obf_3f548a299e41833e *__obf_5275791a072a11d1[V]) PushBack(__obf_95c4c09a56c1a070 string, __obf_6d18c47092a65f53 V) *Element[V] {
	__obf_673d5726ca470196 := &Element[V]{Key: __obf_95c4c09a56c1a070, Value: __obf_6d18c47092a65f53}
	if __obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_475810ae089ac9c9 == nil {
		// It's the first element
		__obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_3f38a216fb054660 = __obf_673d5726ca470196
		__obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_475810ae089ac9c9 = __obf_673d5726ca470196
		return __obf_673d5726ca470196
	}

	__obf_673d5726ca470196.__obf_475810ae089ac9c9 = __obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_475810ae089ac9c9
	__obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_475810ae089ac9c9.__obf_3f38a216fb054660 = __obf_673d5726ca470196
	__obf_3f548a299e41833e.__obf_549d995cf5b98905.__obf_475810ae089ac9c9 = __obf_673d5726ca470196
	return __obf_673d5726ca470196
}
