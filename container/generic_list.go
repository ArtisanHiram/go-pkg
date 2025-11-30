package __obf_e72ce603d10d02a1

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[V any] struct {
	__obf_dbfc801eeafa9b57,
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	__obf_69e5dfbc702ecf34 *Element[V]

	// The key that corresponds to this element in the ordered map.
	Key string

	// The value stored with this element.
	Value V
}

// Next returns the next list element or nil.
func (__obf_acdf9ddfac58a9c9 *Element[V]) Next() *Element[V] {
	return __obf_acdf9ddfac58a9c9.

	// Prev returns the previous list element or nil.
	__obf_dbfc801eeafa9b57
}

func (__obf_acdf9ddfac58a9c9 *Element[V]) Prev() *Element[V] {
	return __obf_acdf9ddfac58a9c9.

	// list represents a null terminated (non circular) intrusive doubly linked list.
	// The list is immediately usable after instantiation without the need of a dedicated initialization.
	__obf_69e5dfbc702ecf34
}

type __obf_e84590c269d065c9[V any] struct {
	__obf_f07ead6617abfc58 Element[V] // list head and tail
}

func (__obf_82d6abd2a338508b *__obf_e84590c269d065c9[V]) IsEmpty() bool {
	return __obf_82d6abd2a338508b.__obf_f07ead6617abfc58.

	// Front returns the first element of list l or nil if the list is empty.
	__obf_dbfc801eeafa9b57 == nil
}

func (__obf_82d6abd2a338508b *__obf_e84590c269d065c9[V]) Front() *Element[V] {
	return __obf_82d6abd2a338508b.

	// Back returns the last element of list l or nil if the list is empty.
	__obf_f07ead6617abfc58.__obf_dbfc801eeafa9b57
}

func (__obf_82d6abd2a338508b *__obf_e84590c269d065c9[V]) Back() *Element[V] {
	return __obf_82d6abd2a338508b.

	// Remove removes e from its list
	__obf_f07ead6617abfc58.__obf_69e5dfbc702ecf34
}

func (__obf_82d6abd2a338508b *__obf_e84590c269d065c9[V]) Remove(__obf_acdf9ddfac58a9c9 *Element[V]) {
	if __obf_acdf9ddfac58a9c9.__obf_69e5dfbc702ecf34 == nil {
		__obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_dbfc801eeafa9b57 = __obf_acdf9ddfac58a9c9.__obf_dbfc801eeafa9b57
	} else {
		__obf_acdf9ddfac58a9c9.__obf_69e5dfbc702ecf34.__obf_dbfc801eeafa9b57 = __obf_acdf9ddfac58a9c9.__obf_dbfc801eeafa9b57
	}
	if __obf_acdf9ddfac58a9c9.__obf_dbfc801eeafa9b57 == nil {
		__obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_69e5dfbc702ecf34 = __obf_acdf9ddfac58a9c9.__obf_69e5dfbc702ecf34
	} else {
		__obf_acdf9ddfac58a9c9.__obf_dbfc801eeafa9b57.

		// avoid memory leaks
		__obf_69e5dfbc702ecf34 = __obf_acdf9ddfac58a9c9.__obf_69e5dfbc702ecf34
	}
	__obf_acdf9ddfac58a9c9.__obf_dbfc801eeafa9b57 = nil
	__obf_acdf9ddfac58a9c9.__obf_69e5dfbc702ecf34 = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (__obf_82d6abd2a338508b *__obf_e84590c269d065c9[V]) PushFront(__obf_9dd08479fe906662 string, __obf_75ac60bfa194af6b V) *Element[V] {
	__obf_acdf9ddfac58a9c9 := &Element[V]{Key: __obf_9dd08479fe906662, Value: __obf_75ac60bfa194af6b}
	if __obf_82d6abd2a338508b.__obf_f07ead6617abfc58.
	// It's the first element
	__obf_dbfc801eeafa9b57 == nil {
		__obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_dbfc801eeafa9b57 = __obf_acdf9ddfac58a9c9
		__obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_69e5dfbc702ecf34 = __obf_acdf9ddfac58a9c9
		return __obf_acdf9ddfac58a9c9
	}
	__obf_acdf9ddfac58a9c9.__obf_dbfc801eeafa9b57 = __obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_dbfc801eeafa9b57

	// PushBack inserts a new element e with value v at the back of list l and returns e.
	__obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_dbfc801eeafa9b57.__obf_69e5dfbc702ecf34 = __obf_acdf9ddfac58a9c9
	__obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_dbfc801eeafa9b57 = __obf_acdf9ddfac58a9c9
	return __obf_acdf9ddfac58a9c9
}

func (__obf_82d6abd2a338508b *__obf_e84590c269d065c9[V]) PushBack(__obf_9dd08479fe906662 string, __obf_75ac60bfa194af6b V) *Element[V] {
	__obf_acdf9ddfac58a9c9 := &Element[V]{Key: __obf_9dd08479fe906662, Value: __obf_75ac60bfa194af6b}
	if __obf_82d6abd2a338508b.__obf_f07ead6617abfc58.
	// It's the first element
	__obf_69e5dfbc702ecf34 == nil {
		__obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_dbfc801eeafa9b57 = __obf_acdf9ddfac58a9c9
		__obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_69e5dfbc702ecf34 = __obf_acdf9ddfac58a9c9
		return __obf_acdf9ddfac58a9c9
	}
	__obf_acdf9ddfac58a9c9.__obf_69e5dfbc702ecf34 = __obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_69e5dfbc702ecf34
	__obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_69e5dfbc702ecf34.__obf_dbfc801eeafa9b57 = __obf_acdf9ddfac58a9c9
	__obf_82d6abd2a338508b.__obf_f07ead6617abfc58.__obf_69e5dfbc702ecf34 = __obf_acdf9ddfac58a9c9
	return __obf_acdf9ddfac58a9c9
}
