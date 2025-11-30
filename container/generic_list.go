package __obf_038560a94647875f

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	__obf_612bb80bc1fd44c4,
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_9432ad7a4ca8c46e *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_55c18423f3794a21 *Element[V]) Next() *Element[V] {
	return __obf_55c18423f3794a21.

	// Prev returns the previous list element or nil.
	__obf_612bb80bc1fd44c4
}

func (__obf_55c18423f3794a21 *Element[V]) Prev() *Element[V] {
	return __obf_55c18423f3794a21.

	// list represents a null terminated (non circular) intrusive doubly linked list.
	// The list is immediately usable after instantiation without the need of a dedicated initialization.
	__obf_9432ad7a4ca8c46e
}

type __obf_295482420737c0b5[V any] struct {
	__obf_d5c554c694036b84 Element[V] // list head and tail
}

func (__obf_d0d441a66e79a49d *__obf_295482420737c0b5[V]) IsEmpty() bool {
	return __obf_d0d441a66e79a49d.__obf_d5c554c694036b84.

	// Front returns the first element of list l or nil if the list is empty.
	__obf_612bb80bc1fd44c4 == nil
}

func (__obf_d0d441a66e79a49d *__obf_295482420737c0b5[V]) Front() *Element[V] {
	return __obf_d0d441a66e79a49d.

	// Back returns the last element of list l or nil if the list is empty.
	__obf_d5c554c694036b84.__obf_612bb80bc1fd44c4
}

func (__obf_d0d441a66e79a49d *__obf_295482420737c0b5[V]) Back() *Element[V] {
	return __obf_d0d441a66e79a49d.

	// Remove removes e from its list
	__obf_d5c554c694036b84.__obf_9432ad7a4ca8c46e
}

func (__obf_d0d441a66e79a49d *__obf_295482420737c0b5[V]) Remove(__obf_55c18423f3794a21 *Element[V]) {
	if __obf_55c18423f3794a21.__obf_9432ad7a4ca8c46e == nil {
		__obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_612bb80bc1fd44c4 = __obf_55c18423f3794a21.__obf_612bb80bc1fd44c4
	} else {
		__obf_55c18423f3794a21.__obf_9432ad7a4ca8c46e.__obf_612bb80bc1fd44c4 = __obf_55c18423f3794a21.__obf_612bb80bc1fd44c4
	}
	if __obf_55c18423f3794a21.__obf_612bb80bc1fd44c4 == nil {
		__obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_9432ad7a4ca8c46e = __obf_55c18423f3794a21.__obf_9432ad7a4ca8c46e
	} else {
		__obf_55c18423f3794a21.__obf_612bb80bc1fd44c4.

		// avoid memory leaks
		__obf_9432ad7a4ca8c46e = __obf_55c18423f3794a21.__obf_9432ad7a4ca8c46e
	}
	__obf_55c18423f3794a21.__obf_612bb80bc1fd44c4 = nil
	__obf_55c18423f3794a21.__obf_9432ad7a4ca8c46e = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_d0d441a66e79a49d *__obf_295482420737c0b5[V]) PushFront(__obf_52e73bb48e810dd2 string, __obf_3270ab9e70995438 V) *Element[V] {
	__obf_55c18423f3794a21 := &Element[V]{Key: __obf_52e73bb48e810dd2, Value: __obf_3270ab9e70995438}
	if __obf_d0d441a66e79a49d.__obf_d5c554c694036b84.
	// It's the first element
	__obf_612bb80bc1fd44c4 == nil {
		__obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_612bb80bc1fd44c4 = __obf_55c18423f3794a21
		__obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_9432ad7a4ca8c46e = __obf_55c18423f3794a21
		return __obf_55c18423f3794a21
	}
	__obf_55c18423f3794a21.__obf_612bb80bc1fd44c4 = __obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_612bb80bc1fd44c4

	// PushBack inserts a new element e with value v at the back of list l and returns e.
	__obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_612bb80bc1fd44c4.__obf_9432ad7a4ca8c46e = __obf_55c18423f3794a21
	__obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_612bb80bc1fd44c4 = __obf_55c18423f3794a21
	return __obf_55c18423f3794a21
}

func (__obf_d0d441a66e79a49d *__obf_295482420737c0b5[V]) PushBack(__obf_52e73bb48e810dd2 string, __obf_3270ab9e70995438 V) *Element[V] {
	__obf_55c18423f3794a21 := &Element[V]{Key: __obf_52e73bb48e810dd2, Value: __obf_3270ab9e70995438}
	if __obf_d0d441a66e79a49d.__obf_d5c554c694036b84.
	// It's the first element
	__obf_9432ad7a4ca8c46e == nil {
		__obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_612bb80bc1fd44c4 = __obf_55c18423f3794a21
		__obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_9432ad7a4ca8c46e = __obf_55c18423f3794a21
		return __obf_55c18423f3794a21
	}
	__obf_55c18423f3794a21.__obf_9432ad7a4ca8c46e = __obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_9432ad7a4ca8c46e
	__obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_9432ad7a4ca8c46e.__obf_612bb80bc1fd44c4 = __obf_55c18423f3794a21
	__obf_d0d441a66e79a49d.__obf_d5c554c694036b84.__obf_9432ad7a4ca8c46e = __obf_55c18423f3794a21
	return __obf_55c18423f3794a21
}
