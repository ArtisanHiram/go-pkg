package __obf_90a4883a02d0b41c

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_b12808a307454062 = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_bdeb577f5f8651b9 []any
	__obf_d4747c905c4d989c int
	__obf_30fd5c9692a3bb01 int
	__obf_ee99c442f9d25bf4 int
	__obf_e0201e1587be4a15 int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_d0b5b6f649a98c27 *Deque) Len() int {
	return __obf_d0b5b6f649a98c27.

	// PushBack appends an element to the back of the queue.  Implements FIFO when
	// elements are removed with PopFront(), and LIFO when elements are removed
	// with PopBack().
	__obf_ee99c442f9d25bf4
}

func (__obf_d0b5b6f649a98c27 *Deque) PushBack(__obf_1afa695a0c002d93 any) {
	__obf_d0b5b6f649a98c27.__obf_871ea368bc5734fe()
	__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.
	// Calculate new tail position.
	__obf_30fd5c9692a3bb01] = __obf_1afa695a0c002d93
	__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01 = __obf_d0b5b6f649a98c27.__obf_32b235e6da4cbc13(__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01)
	__obf_d0b5b6f649a98c27.

	// PushFront prepends an element to the front of the queue.
	__obf_ee99c442f9d25bf4++
}

func (__obf_d0b5b6f649a98c27 *Deque) PushFront(__obf_1afa695a0c002d93 any) {
	__obf_d0b5b6f649a98c27.

	// Calculate new head position.
	__obf_871ea368bc5734fe()
	__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c = __obf_d0b5b6f649a98c27.__obf_e3286a9ca021156f(__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c)
	__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c] = __obf_1afa695a0c002d93
	__obf_d0b5b6f649a98c27.

	// PopFront removes and returns the element from the front of the queue.
	// Implements FIFO when used with PushBack().  If the queue is empty, the call
	// panics.
	__obf_ee99c442f9d25bf4++
}

func (__obf_d0b5b6f649a98c27 *Deque) PopFront() any {
	if __obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4 <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_b828fde9e3ad6671 := __obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c]
	__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.
	// Calculate new head position.
	__obf_d4747c905c4d989c] = nil
	__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c = __obf_d0b5b6f649a98c27.__obf_32b235e6da4cbc13(__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c)
	__obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4--
	__obf_d0b5b6f649a98c27.__obf_250415b7024be355()
	return __obf_b828fde9e3ad6671
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_d0b5b6f649a98c27 *Deque) PopBack() any {
	if __obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4 <= 0 {
		panic("deque: PopBack() called on empty queue")
	}
	__obf_d0b5b6f649a98c27.

	// Calculate new tail position
	__obf_30fd5c9692a3bb01 = __obf_d0b5b6f649a98c27.

	// Remove value at tail.
	__obf_e3286a9ca021156f(__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01)
	__obf_b828fde9e3ad6671 := __obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01]
	__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01] = nil
	__obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4--
	__obf_d0b5b6f649a98c27.__obf_250415b7024be355()
	return __obf_b828fde9e3ad6671
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_d0b5b6f649a98c27 *Deque) Front() any {
	if __obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4 <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_d0b5b6f649a98c27.

	// Back returns the element at the back of the queue.  This is the element
	// that would be returned by PopBack().  This call panics if the queue is
	// empty.
	__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c]
}

func (__obf_d0b5b6f649a98c27 *Deque) Back() any {
	if __obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4 <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.

	// At returns the element at index i in the queue without removing the element
	// from the queue.  This method accepts only non-negative index values.  At(0)
	// refers to the first element and is the same as Front().  At(Len()-1) refers
	// to the last element and is the same as Back().  If the index is invalid, the
	// call panics.
	//
	// The purpose of At is to allow Deque to serve as a more general purpose
	// circular buffer, where items are only added to and removed from the ends of
	// the deque, but may be read from any place within the deque.  Consider the
	// case of a fixed-size circular log buffer: A new entry is pushed onto one end
	// and when full the oldest is popped from the other end.  All the log entries
	// in the buffer must be readable without altering the buffer contents.
	__obf_e3286a9ca021156f(__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01)]
}

func (__obf_d0b5b6f649a98c27 *Deque) At(__obf_f315019af10902c2 int) any {
	if __obf_f315019af10902c2 < 0 || __obf_f315019af10902c2 >= __obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4 {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[(__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c+__obf_f315019af10902c2)&(len(__obf_d0b5b6f649a98c27.

	// Clear removes all elements from the queue, but retains the current capacity.
	// This is useful when repeatedly reusing the queue at high frequency to avoid
	// GC during reuse.  The queue will not be resized smaller as long as items are
	// only added.  Only when items are removed is the queue subject to getting
	// resized smaller.
	__obf_bdeb577f5f8651b9)-1)]
}

func (__obf_d0b5b6f649a98c27 *Deque) Clear() {
	__obf_b65c862793db56e1 :=// bitwise modulus
	len(__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9) - 1
	for __obf_ce6856a901600875 := __obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c; __obf_ce6856a901600875 != __obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01; __obf_ce6856a901600875 = (__obf_ce6856a901600875 + 1) & __obf_b65c862793db56e1 {
		__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_ce6856a901600875] = nil
	}
	__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c = 0
	__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01 = 0
	__obf_d0b5b6f649a98c27.

	// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
	// back-to-front.  Having Deque provide Rotate() avoids resizing that could
	// happen if implementing rotation using only Pop and Push methods.
	__obf_ee99c442f9d25bf4 = 0
}

func (__obf_d0b5b6f649a98c27 *Deque) Rotate(__obf_12d9f59bac8fe254 int) {
	if __obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4 <= 1 {
		return
	}
	__obf_12d9f59bac8fe254 %=// Rotating a multiple of q.count is same as no rotation.
	__obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4
	if __obf_12d9f59bac8fe254 == 0 {
		return
	}
	__obf_b65c862793db56e1 := len(__obf_d0b5b6f649a98c27.
	// If no empty space in buffer, only move head and tail indexes.
	__obf_bdeb577f5f8651b9) - 1

	if __obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c == __obf_d0b5b6f649a98c27.
	// Calculate new head and tail using bitwise modulus.
	__obf_30fd5c9692a3bb01 {
		__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c = (__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c + __obf_12d9f59bac8fe254) & __obf_b65c862793db56e1
		__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01 = (__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01 + __obf_12d9f59bac8fe254) & __obf_b65c862793db56e1
		return
	}

	if __obf_12d9f59bac8fe254 < 0 {
		// Rotate back to front.
		for ; __obf_12d9f59bac8fe254 < 0; __obf_12d9f59bac8fe254++ {
			__obf_d0b5b6f649a98c27.
			// Calculate new head and tail using bitwise modulus.
			__obf_d4747c905c4d989c = (__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c - 1) & __obf_b65c862793db56e1
			__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01 = (__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01 - 1) & __obf_b65c862793db56e1
			// Put tail value at head and remove value at tail.
			__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c] = __obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01]
			__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_12d9f59bac8fe254 > 0; __obf_12d9f59bac8fe254-- {
		__obf_d0b5b6f649a98c27.
		// Put head value at tail and remove value at head.
		__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01] = __obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c]
		__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.
		// Calculate new head and tail using bitwise modulus.
		__obf_d4747c905c4d989c] = nil
		__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c = (__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c + 1) & __obf_b65c862793db56e1
		__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01 = (__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01 + 1) & __obf_b65c862793db56e1
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_d0b5b6f649a98c27 *Deque) SetMinCapacity(__obf_c9704716748d08b1 uint) {
	if 1<<__obf_c9704716748d08b1 > __obf_b12808a307454062 {
		__obf_d0b5b6f649a98c27.__obf_e0201e1587be4a15 = 1 << __obf_c9704716748d08b1
	} else {
		__obf_d0b5b6f649a98c27.__obf_e0201e1587be4a15 = __obf_b12808a307454062
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_d0b5b6f649a98c27 *Deque) __obf_e3286a9ca021156f(__obf_f315019af10902c2 int) int {
	return (__obf_f315019af10902c2 - 1) & (len(__obf_d0b5b6f649a98c27. // bitwise modulus
	__obf_bdeb577f5f8651b9) - 1)
}

// next returns the next buffer position wrapping around buffer.
func (__obf_d0b5b6f649a98c27 *Deque) __obf_32b235e6da4cbc13(__obf_f315019af10902c2 int) int {
	return (__obf_f315019af10902c2 + 1) & (len(__obf_d0b5b6f649a98c27. // bitwise modulus
	__obf_bdeb577f5f8651b9) - 1)
}

// growIfFull resizes up if the buffer is full.
func (__obf_d0b5b6f649a98c27 *Deque) __obf_871ea368bc5734fe() {
	if len(__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9) == 0 {
		if __obf_d0b5b6f649a98c27.__obf_e0201e1587be4a15 == 0 {
			__obf_d0b5b6f649a98c27.__obf_e0201e1587be4a15 = __obf_b12808a307454062
		}
		__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9 = make([]any, __obf_d0b5b6f649a98c27.__obf_e0201e1587be4a15)
		return
	}
	if __obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4 == len(__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9) {
		__obf_d0b5b6f649a98c27.__obf_9a83f87899818fc5()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_d0b5b6f649a98c27 *Deque) __obf_250415b7024be355() {
	if len(__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9) > __obf_d0b5b6f649a98c27.__obf_e0201e1587be4a15 && (__obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4<<2) == len(__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9) {
		__obf_d0b5b6f649a98c27.__obf_9a83f87899818fc5()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_d0b5b6f649a98c27 *Deque) __obf_9a83f87899818fc5() {
	__obf_f8984508fb457f45 := make([]any, __obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4<<1)
	if __obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01 > __obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c {
		copy(__obf_f8984508fb457f45, __obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c:__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01])
	} else {
		__obf_12d9f59bac8fe254 := copy(__obf_f8984508fb457f45, __obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c:])
		copy(__obf_f8984508fb457f45[__obf_12d9f59bac8fe254:], __obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9[:__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01])
	}
	__obf_d0b5b6f649a98c27.__obf_d4747c905c4d989c = 0
	__obf_d0b5b6f649a98c27.__obf_30fd5c9692a3bb01 = __obf_d0b5b6f649a98c27.__obf_ee99c442f9d25bf4
	__obf_d0b5b6f649a98c27.__obf_bdeb577f5f8651b9 = __obf_f8984508fb457f45
}
