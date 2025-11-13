package __obf_76599ab96a208083

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_54005283f515b2f3 = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_9fe268d736364e9f []any
	__obf_4d89c0af2f36d012 int
	__obf_8bc2d90b334ac328 int
	__obf_c5d600e67da7e0a9 int
	__obf_0b2b4cb4a4eac34c int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_11028cff7d6f8d69 *Deque) Len() int {
	return __obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9
}

// PushBack appends an element to the back of the queue.  Implements FIFO when
// elements are removed with PopFront(), and LIFO when elements are removed
// with PopBack().
func (__obf_11028cff7d6f8d69 *Deque) PushBack(__obf_820b59466c6fe31c any) {
	__obf_11028cff7d6f8d69.__obf_2013021a3470d347()

	__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328] = __obf_820b59466c6fe31c
	// Calculate new tail position.
	__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 = __obf_11028cff7d6f8d69.__obf_b7a3ba8614d8f9d8(__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328)
	__obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9++
}

// PushFront prepends an element to the front of the queue.
func (__obf_11028cff7d6f8d69 *Deque) PushFront(__obf_820b59466c6fe31c any) {
	__obf_11028cff7d6f8d69.__obf_2013021a3470d347()

	// Calculate new head position.
	__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 = __obf_11028cff7d6f8d69.__obf_6b94d13992080dfb(__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012)
	__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012] = __obf_820b59466c6fe31c
	__obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9++
}

// PopFront removes and returns the element from the front of the queue.
// Implements FIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_11028cff7d6f8d69 *Deque) PopFront() any {
	if __obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9 <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_12bd3b3b2da39ede := __obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012]
	__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012] = nil
	// Calculate new head position.
	__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 = __obf_11028cff7d6f8d69.__obf_b7a3ba8614d8f9d8(__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012)
	__obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9--

	__obf_11028cff7d6f8d69.__obf_232fd657786529e3()
	return __obf_12bd3b3b2da39ede
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_11028cff7d6f8d69 *Deque) PopBack() any {
	if __obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9 <= 0 {
		panic("deque: PopBack() called on empty queue")
	}

	// Calculate new tail position
	__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 = __obf_11028cff7d6f8d69.__obf_6b94d13992080dfb(__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328)

	// Remove value at tail.
	__obf_12bd3b3b2da39ede := __obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328]
	__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328] = nil
	__obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9--

	__obf_11028cff7d6f8d69.__obf_232fd657786529e3()
	return __obf_12bd3b3b2da39ede
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_11028cff7d6f8d69 *Deque) Front() any {
	if __obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9 <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012]
}

// Back returns the element at the back of the queue.  This is the element
// that would be returned by PopBack().  This call panics if the queue is
// empty.
func (__obf_11028cff7d6f8d69 *Deque) Back() any {
	if __obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9 <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_6b94d13992080dfb(__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328)]
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
func (__obf_11028cff7d6f8d69 *Deque) At(__obf_8d4edb68382a63c4 int) any {
	if __obf_8d4edb68382a63c4 < 0 || __obf_8d4edb68382a63c4 >= __obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9 {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[(__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012+__obf_8d4edb68382a63c4)&(len(__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f)-1)]
}

// Clear removes all elements from the queue, but retains the current capacity.
// This is useful when repeatedly reusing the queue at high frequency to avoid
// GC during reuse.  The queue will not be resized smaller as long as items are
// only added.  Only when items are removed is the queue subject to getting
// resized smaller.
func (__obf_11028cff7d6f8d69 *Deque) Clear() {
	// bitwise modulus
	__obf_c6970df5fc1f25be := len(__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f) - 1
	for __obf_703df7aac448a8d0 := __obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012; __obf_703df7aac448a8d0 != __obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328; __obf_703df7aac448a8d0 = (__obf_703df7aac448a8d0 + 1) & __obf_c6970df5fc1f25be {
		__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_703df7aac448a8d0] = nil
	}
	__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 = 0
	__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 = 0
	__obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9 = 0
}

// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
// back-to-front.  Having Deque provide Rotate() avoids resizing that could
// happen if implementing rotation using only Pop and Push methods.
func (__obf_11028cff7d6f8d69 *Deque) Rotate(__obf_18a54e2632d73807 int) {
	if __obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9 <= 1 {
		return
	}
	// Rotating a multiple of q.count is same as no rotation.
	__obf_18a54e2632d73807 %= __obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9
	if __obf_18a54e2632d73807 == 0 {
		return
	}

	__obf_c6970df5fc1f25be := len(__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f) - 1
	// If no empty space in buffer, only move head and tail indexes.
	if __obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 == __obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 {
		// Calculate new head and tail using bitwise modulus.
		__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 = (__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 + __obf_18a54e2632d73807) & __obf_c6970df5fc1f25be
		__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 = (__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 + __obf_18a54e2632d73807) & __obf_c6970df5fc1f25be
		return
	}

	if __obf_18a54e2632d73807 < 0 {
		// Rotate back to front.
		for ; __obf_18a54e2632d73807 < 0; __obf_18a54e2632d73807++ {
			// Calculate new head and tail using bitwise modulus.
			__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 = (__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 - 1) & __obf_c6970df5fc1f25be
			__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 = (__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 - 1) & __obf_c6970df5fc1f25be
			// Put tail value at head and remove value at tail.
			__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012] = __obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328]
			__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_18a54e2632d73807 > 0; __obf_18a54e2632d73807-- {
		// Put head value at tail and remove value at head.
		__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328] = __obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012]
		__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012] = nil
		// Calculate new head and tail using bitwise modulus.
		__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 = (__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 + 1) & __obf_c6970df5fc1f25be
		__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 = (__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 + 1) & __obf_c6970df5fc1f25be
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_11028cff7d6f8d69 *Deque) SetMinCapacity(__obf_acc83aca276429bf uint) {
	if 1<<__obf_acc83aca276429bf > __obf_54005283f515b2f3 {
		__obf_11028cff7d6f8d69.__obf_0b2b4cb4a4eac34c = 1 << __obf_acc83aca276429bf
	} else {
		__obf_11028cff7d6f8d69.__obf_0b2b4cb4a4eac34c = __obf_54005283f515b2f3
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_11028cff7d6f8d69 *Deque) __obf_6b94d13992080dfb(__obf_8d4edb68382a63c4 int) int {
	return (__obf_8d4edb68382a63c4 - 1) & (len(__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f) - 1) // bitwise modulus
}

// next returns the next buffer position wrapping around buffer.
func (__obf_11028cff7d6f8d69 *Deque) __obf_b7a3ba8614d8f9d8(__obf_8d4edb68382a63c4 int) int {
	return (__obf_8d4edb68382a63c4 + 1) & (len(__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f) - 1) // bitwise modulus
}

// growIfFull resizes up if the buffer is full.
func (__obf_11028cff7d6f8d69 *Deque) __obf_2013021a3470d347() {
	if len(__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f) == 0 {
		if __obf_11028cff7d6f8d69.__obf_0b2b4cb4a4eac34c == 0 {
			__obf_11028cff7d6f8d69.__obf_0b2b4cb4a4eac34c = __obf_54005283f515b2f3
		}
		__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f = make([]any, __obf_11028cff7d6f8d69.__obf_0b2b4cb4a4eac34c)
		return
	}
	if __obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9 == len(__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f) {
		__obf_11028cff7d6f8d69.__obf_b364840a737f7895()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_11028cff7d6f8d69 *Deque) __obf_232fd657786529e3() {
	if len(__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f) > __obf_11028cff7d6f8d69.__obf_0b2b4cb4a4eac34c && (__obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9<<2) == len(__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f) {
		__obf_11028cff7d6f8d69.__obf_b364840a737f7895()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_11028cff7d6f8d69 *Deque) __obf_b364840a737f7895() {
	__obf_9d02c6602c4d3519 := make([]any, __obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9<<1)
	if __obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 > __obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 {
		copy(__obf_9d02c6602c4d3519, __obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012:__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328])
	} else {
		__obf_18a54e2632d73807 := copy(__obf_9d02c6602c4d3519, __obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012:])
		copy(__obf_9d02c6602c4d3519[__obf_18a54e2632d73807:], __obf_11028cff7d6f8d69.__obf_9fe268d736364e9f[:__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328])
	}

	__obf_11028cff7d6f8d69.__obf_4d89c0af2f36d012 = 0
	__obf_11028cff7d6f8d69.__obf_8bc2d90b334ac328 = __obf_11028cff7d6f8d69.__obf_c5d600e67da7e0a9
	__obf_11028cff7d6f8d69.__obf_9fe268d736364e9f = __obf_9d02c6602c4d3519
}
