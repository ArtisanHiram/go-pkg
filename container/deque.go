package __obf_62eba4024f8fa381

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_f2030f8383eac9da = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_feefcd00f16d85f9 []any
	__obf_1fbf00a2407f816f int
	__obf_16dadd0fd1fc0113 int
	__obf_6c61f1671393b933 int
	__obf_749f3e5b121d5f03 int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_77299e995ace9e00 *Deque) Len() int {
	return __obf_77299e995ace9e00.__obf_6c61f1671393b933
}

// PushBack appends an element to the back of the queue.  Implements FIFO when
// elements are removed with PopFront(), and LIFO when elements are removed
// with PopBack().
func (__obf_77299e995ace9e00 *Deque) PushBack(__obf_3c95f4299d9b5c3d any) {
	__obf_77299e995ace9e00.__obf_7f95ace05b715888()

	__obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113] = __obf_3c95f4299d9b5c3d
	// Calculate new tail position.
	__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 = __obf_77299e995ace9e00.__obf_a88f24b4f460aa9c(__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113)
	__obf_77299e995ace9e00.__obf_6c61f1671393b933++
}

// PushFront prepends an element to the front of the queue.
func (__obf_77299e995ace9e00 *Deque) PushFront(__obf_3c95f4299d9b5c3d any) {
	__obf_77299e995ace9e00.__obf_7f95ace05b715888()

	// Calculate new head position.
	__obf_77299e995ace9e00.__obf_1fbf00a2407f816f = __obf_77299e995ace9e00.__obf_fc44d56dec1421d9(__obf_77299e995ace9e00.__obf_1fbf00a2407f816f)
	__obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_1fbf00a2407f816f] = __obf_3c95f4299d9b5c3d
	__obf_77299e995ace9e00.__obf_6c61f1671393b933++
}

// PopFront removes and returns the element from the front of the queue.
// Implements FIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_77299e995ace9e00 *Deque) PopFront() any {
	if __obf_77299e995ace9e00.__obf_6c61f1671393b933 <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_c9ba1390dfa534da := __obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_1fbf00a2407f816f]
	__obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_1fbf00a2407f816f] = nil
	// Calculate new head position.
	__obf_77299e995ace9e00.__obf_1fbf00a2407f816f = __obf_77299e995ace9e00.__obf_a88f24b4f460aa9c(__obf_77299e995ace9e00.__obf_1fbf00a2407f816f)
	__obf_77299e995ace9e00.__obf_6c61f1671393b933--

	__obf_77299e995ace9e00.__obf_e04075e6bde01db5()
	return __obf_c9ba1390dfa534da
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_77299e995ace9e00 *Deque) PopBack() any {
	if __obf_77299e995ace9e00.__obf_6c61f1671393b933 <= 0 {
		panic("deque: PopBack() called on empty queue")
	}

	// Calculate new tail position
	__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 = __obf_77299e995ace9e00.__obf_fc44d56dec1421d9(__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113)

	// Remove value at tail.
	__obf_c9ba1390dfa534da := __obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113]
	__obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113] = nil
	__obf_77299e995ace9e00.__obf_6c61f1671393b933--

	__obf_77299e995ace9e00.__obf_e04075e6bde01db5()
	return __obf_c9ba1390dfa534da
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_77299e995ace9e00 *Deque) Front() any {
	if __obf_77299e995ace9e00.__obf_6c61f1671393b933 <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_1fbf00a2407f816f]
}

// Back returns the element at the back of the queue.  This is the element
// that would be returned by PopBack().  This call panics if the queue is
// empty.
func (__obf_77299e995ace9e00 *Deque) Back() any {
	if __obf_77299e995ace9e00.__obf_6c61f1671393b933 <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_fc44d56dec1421d9(__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113)]
}

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
func (__obf_77299e995ace9e00 *Deque) At(__obf_1376e2b9d2469c18 int) any {
	if __obf_1376e2b9d2469c18 < 0 || __obf_1376e2b9d2469c18 >= __obf_77299e995ace9e00.__obf_6c61f1671393b933 {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_77299e995ace9e00.__obf_feefcd00f16d85f9[(__obf_77299e995ace9e00.__obf_1fbf00a2407f816f+__obf_1376e2b9d2469c18)&(len(__obf_77299e995ace9e00.__obf_feefcd00f16d85f9)-1)]
}

// Clear removes all elements from the queue, but retains the current capacity.
// This is useful when repeatedly reusing the queue at high frequency to avoid
// GC during reuse.  The queue will not be resized smaller as long as items are
// only added.  Only when items are removed is the queue subject to getting
// resized smaller.
func (__obf_77299e995ace9e00 *Deque) Clear() {
	// bitwise modulus
	__obf_1141f6e829eb7fd5 := len(__obf_77299e995ace9e00.__obf_feefcd00f16d85f9) - 1
	for __obf_1271301575c5c2f5 := __obf_77299e995ace9e00.__obf_1fbf00a2407f816f; __obf_1271301575c5c2f5 != __obf_77299e995ace9e00.__obf_16dadd0fd1fc0113; __obf_1271301575c5c2f5 = (__obf_1271301575c5c2f5 + 1) & __obf_1141f6e829eb7fd5 {
		__obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_1271301575c5c2f5] = nil
	}
	__obf_77299e995ace9e00.__obf_1fbf00a2407f816f = 0
	__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 = 0
	__obf_77299e995ace9e00.__obf_6c61f1671393b933 = 0
}

// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
// back-to-front.  Having Deque provide Rotate() avoids resizing that could
// happen if implementing rotation using only Pop and Push methods.
func (__obf_77299e995ace9e00 *Deque) Rotate(__obf_409580b446b55be1 int) {
	if __obf_77299e995ace9e00.__obf_6c61f1671393b933 <= 1 {
		return
	}
	// Rotating a multiple of q.count is same as no rotation.
	__obf_409580b446b55be1 %= __obf_77299e995ace9e00.__obf_6c61f1671393b933
	if __obf_409580b446b55be1 == 0 {
		return
	}

	__obf_1141f6e829eb7fd5 := len(__obf_77299e995ace9e00.__obf_feefcd00f16d85f9) - 1
	// If no empty space in buffer, only move head and tail indexes.
	if __obf_77299e995ace9e00.__obf_1fbf00a2407f816f == __obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 {
		// Calculate new head and tail using bitwise modulus.
		__obf_77299e995ace9e00.__obf_1fbf00a2407f816f = (__obf_77299e995ace9e00.__obf_1fbf00a2407f816f + __obf_409580b446b55be1) & __obf_1141f6e829eb7fd5
		__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 = (__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 + __obf_409580b446b55be1) & __obf_1141f6e829eb7fd5
		return
	}

	if __obf_409580b446b55be1 < 0 {
		// Rotate back to front.
		for ; __obf_409580b446b55be1 < 0; __obf_409580b446b55be1++ {
			// Calculate new head and tail using bitwise modulus.
			__obf_77299e995ace9e00.__obf_1fbf00a2407f816f = (__obf_77299e995ace9e00.__obf_1fbf00a2407f816f - 1) & __obf_1141f6e829eb7fd5
			__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 = (__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 - 1) & __obf_1141f6e829eb7fd5
			// Put tail value at head and remove value at tail.
			__obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_1fbf00a2407f816f] = __obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113]
			__obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_409580b446b55be1 > 0; __obf_409580b446b55be1-- {
		// Put head value at tail and remove value at head.
		__obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113] = __obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_1fbf00a2407f816f]
		__obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_1fbf00a2407f816f] = nil
		// Calculate new head and tail using bitwise modulus.
		__obf_77299e995ace9e00.__obf_1fbf00a2407f816f = (__obf_77299e995ace9e00.__obf_1fbf00a2407f816f + 1) & __obf_1141f6e829eb7fd5
		__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 = (__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 + 1) & __obf_1141f6e829eb7fd5
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_77299e995ace9e00 *Deque) SetMinCapacity(__obf_df8bc4e1aa5a7ae0 uint) {
	if 1<<__obf_df8bc4e1aa5a7ae0 > __obf_f2030f8383eac9da {
		__obf_77299e995ace9e00.__obf_749f3e5b121d5f03 = 1 << __obf_df8bc4e1aa5a7ae0
	} else {
		__obf_77299e995ace9e00.__obf_749f3e5b121d5f03 = __obf_f2030f8383eac9da
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_77299e995ace9e00 *Deque) __obf_fc44d56dec1421d9(__obf_1376e2b9d2469c18 int) int {
	return (__obf_1376e2b9d2469c18 - 1) & (len(__obf_77299e995ace9e00.__obf_feefcd00f16d85f9) - 1) // bitwise modulus
}

// next returns the next buffer position wrapping around buffer.
func (__obf_77299e995ace9e00 *Deque) __obf_a88f24b4f460aa9c(__obf_1376e2b9d2469c18 int) int {
	return (__obf_1376e2b9d2469c18 + 1) & (len(__obf_77299e995ace9e00.__obf_feefcd00f16d85f9) - 1) // bitwise modulus
}

// growIfFull resizes up if the buffer is full.
func (__obf_77299e995ace9e00 *Deque) __obf_7f95ace05b715888() {
	if len(__obf_77299e995ace9e00.__obf_feefcd00f16d85f9) == 0 {
		if __obf_77299e995ace9e00.__obf_749f3e5b121d5f03 == 0 {
			__obf_77299e995ace9e00.__obf_749f3e5b121d5f03 = __obf_f2030f8383eac9da
		}
		__obf_77299e995ace9e00.__obf_feefcd00f16d85f9 = make([]any, __obf_77299e995ace9e00.__obf_749f3e5b121d5f03)
		return
	}
	if __obf_77299e995ace9e00.__obf_6c61f1671393b933 == len(__obf_77299e995ace9e00.__obf_feefcd00f16d85f9) {
		__obf_77299e995ace9e00.__obf_667730052a024a51()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_77299e995ace9e00 *Deque) __obf_e04075e6bde01db5() {
	if len(__obf_77299e995ace9e00.__obf_feefcd00f16d85f9) > __obf_77299e995ace9e00.__obf_749f3e5b121d5f03 && (__obf_77299e995ace9e00.__obf_6c61f1671393b933<<2) == len(__obf_77299e995ace9e00.__obf_feefcd00f16d85f9) {
		__obf_77299e995ace9e00.__obf_667730052a024a51()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_77299e995ace9e00 *Deque) __obf_667730052a024a51() {
	__obf_80fbdaaafe7715c6 := make([]any, __obf_77299e995ace9e00.__obf_6c61f1671393b933<<1)
	if __obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 > __obf_77299e995ace9e00.__obf_1fbf00a2407f816f {
		copy(__obf_80fbdaaafe7715c6, __obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_1fbf00a2407f816f:__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113])
	} else {
		__obf_409580b446b55be1 := copy(__obf_80fbdaaafe7715c6, __obf_77299e995ace9e00.__obf_feefcd00f16d85f9[__obf_77299e995ace9e00.__obf_1fbf00a2407f816f:])
		copy(__obf_80fbdaaafe7715c6[__obf_409580b446b55be1:], __obf_77299e995ace9e00.__obf_feefcd00f16d85f9[:__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113])
	}

	__obf_77299e995ace9e00.__obf_1fbf00a2407f816f = 0
	__obf_77299e995ace9e00.__obf_16dadd0fd1fc0113 = __obf_77299e995ace9e00.__obf_6c61f1671393b933
	__obf_77299e995ace9e00.__obf_feefcd00f16d85f9 = __obf_80fbdaaafe7715c6
}
