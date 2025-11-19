package __obf_9861fa13140c30a3

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_cc95479c3d76dafd = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_22291ba8f5a063e8 []any
	__obf_e7184c1cd4d9cb16 int
	__obf_30d187ec07bc80ec int
	__obf_7e6767c22d23f06d int
	__obf_c3af9f3aa40e7102 int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_cccb1f17f448628f *Deque) Len() int {
	return __obf_cccb1f17f448628f.__obf_7e6767c22d23f06d
}

// PushBack appends an element to the back of the queue.  Implements FIFO when
// elements are removed with PopFront(), and LIFO when elements are removed
// with PopBack().
func (__obf_cccb1f17f448628f *Deque) PushBack(__obf_1cfa18c298665e4b any) {
	__obf_cccb1f17f448628f.__obf_5c24d5ef59e77124()

	__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec] = __obf_1cfa18c298665e4b
	// Calculate new tail position.
	__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec = __obf_cccb1f17f448628f.__obf_9c578e473cd91a70(__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec)
	__obf_cccb1f17f448628f.__obf_7e6767c22d23f06d++
}

// PushFront prepends an element to the front of the queue.
func (__obf_cccb1f17f448628f *Deque) PushFront(__obf_1cfa18c298665e4b any) {
	__obf_cccb1f17f448628f.__obf_5c24d5ef59e77124()

	// Calculate new head position.
	__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 = __obf_cccb1f17f448628f.__obf_9e64c21ddd522887(__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16)
	__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16] = __obf_1cfa18c298665e4b
	__obf_cccb1f17f448628f.__obf_7e6767c22d23f06d++
}

// PopFront removes and returns the element from the front of the queue.
// Implements FIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_cccb1f17f448628f *Deque) PopFront() any {
	if __obf_cccb1f17f448628f.__obf_7e6767c22d23f06d <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_960335ceca7c79b8 := __obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16]
	__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16] = nil
	// Calculate new head position.
	__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 = __obf_cccb1f17f448628f.__obf_9c578e473cd91a70(__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16)
	__obf_cccb1f17f448628f.__obf_7e6767c22d23f06d--

	__obf_cccb1f17f448628f.__obf_e286abf03b28bd7a()
	return __obf_960335ceca7c79b8
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_cccb1f17f448628f *Deque) PopBack() any {
	if __obf_cccb1f17f448628f.__obf_7e6767c22d23f06d <= 0 {
		panic("deque: PopBack() called on empty queue")
	}

	// Calculate new tail position
	__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec = __obf_cccb1f17f448628f.__obf_9e64c21ddd522887(__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec)

	// Remove value at tail.
	__obf_960335ceca7c79b8 := __obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec]
	__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec] = nil
	__obf_cccb1f17f448628f.__obf_7e6767c22d23f06d--

	__obf_cccb1f17f448628f.__obf_e286abf03b28bd7a()
	return __obf_960335ceca7c79b8
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_cccb1f17f448628f *Deque) Front() any {
	if __obf_cccb1f17f448628f.__obf_7e6767c22d23f06d <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16]
}

// Back returns the element at the back of the queue.  This is the element
// that would be returned by PopBack().  This call panics if the queue is
// empty.
func (__obf_cccb1f17f448628f *Deque) Back() any {
	if __obf_cccb1f17f448628f.__obf_7e6767c22d23f06d <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_9e64c21ddd522887(__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec)]
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
func (__obf_cccb1f17f448628f *Deque) At(__obf_35fc60e92e4b05ba int) any {
	if __obf_35fc60e92e4b05ba < 0 || __obf_35fc60e92e4b05ba >= __obf_cccb1f17f448628f.__obf_7e6767c22d23f06d {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[(__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16+__obf_35fc60e92e4b05ba)&(len(__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8)-1)]
}

// Clear removes all elements from the queue, but retains the current capacity.
// This is useful when repeatedly reusing the queue at high frequency to avoid
// GC during reuse.  The queue will not be resized smaller as long as items are
// only added.  Only when items are removed is the queue subject to getting
// resized smaller.
func (__obf_cccb1f17f448628f *Deque) Clear() {
	// bitwise modulus
	__obf_cbf465f12a13d3c7 := len(__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8) - 1
	for __obf_aceb1679325404db := __obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16; __obf_aceb1679325404db != __obf_cccb1f17f448628f.__obf_30d187ec07bc80ec; __obf_aceb1679325404db = (__obf_aceb1679325404db + 1) & __obf_cbf465f12a13d3c7 {
		__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_aceb1679325404db] = nil
	}
	__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 = 0
	__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec = 0
	__obf_cccb1f17f448628f.__obf_7e6767c22d23f06d = 0
}

// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
// back-to-front.  Having Deque provide Rotate() avoids resizing that could
// happen if implementing rotation using only Pop and Push methods.
func (__obf_cccb1f17f448628f *Deque) Rotate(__obf_bdf68a51ae002dc2 int) {
	if __obf_cccb1f17f448628f.__obf_7e6767c22d23f06d <= 1 {
		return
	}
	// Rotating a multiple of q.count is same as no rotation.
	__obf_bdf68a51ae002dc2 %= __obf_cccb1f17f448628f.__obf_7e6767c22d23f06d
	if __obf_bdf68a51ae002dc2 == 0 {
		return
	}

	__obf_cbf465f12a13d3c7 := len(__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8) - 1
	// If no empty space in buffer, only move head and tail indexes.
	if __obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 == __obf_cccb1f17f448628f.__obf_30d187ec07bc80ec {
		// Calculate new head and tail using bitwise modulus.
		__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 = (__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 + __obf_bdf68a51ae002dc2) & __obf_cbf465f12a13d3c7
		__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec = (__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec + __obf_bdf68a51ae002dc2) & __obf_cbf465f12a13d3c7
		return
	}

	if __obf_bdf68a51ae002dc2 < 0 {
		// Rotate back to front.
		for ; __obf_bdf68a51ae002dc2 < 0; __obf_bdf68a51ae002dc2++ {
			// Calculate new head and tail using bitwise modulus.
			__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 = (__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 - 1) & __obf_cbf465f12a13d3c7
			__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec = (__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec - 1) & __obf_cbf465f12a13d3c7
			// Put tail value at head and remove value at tail.
			__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16] = __obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec]
			__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_bdf68a51ae002dc2 > 0; __obf_bdf68a51ae002dc2-- {
		// Put head value at tail and remove value at head.
		__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec] = __obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16]
		__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16] = nil
		// Calculate new head and tail using bitwise modulus.
		__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 = (__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 + 1) & __obf_cbf465f12a13d3c7
		__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec = (__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec + 1) & __obf_cbf465f12a13d3c7
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_cccb1f17f448628f *Deque) SetMinCapacity(__obf_e436a0606c4ddaec uint) {
	if 1<<__obf_e436a0606c4ddaec > __obf_cc95479c3d76dafd {
		__obf_cccb1f17f448628f.__obf_c3af9f3aa40e7102 = 1 << __obf_e436a0606c4ddaec
	} else {
		__obf_cccb1f17f448628f.__obf_c3af9f3aa40e7102 = __obf_cc95479c3d76dafd
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_cccb1f17f448628f *Deque) __obf_9e64c21ddd522887(__obf_35fc60e92e4b05ba int) int {
	return (__obf_35fc60e92e4b05ba - 1) & (len(__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8) - 1) // bitwise modulus
}

// next returns the next buffer position wrapping around buffer.
func (__obf_cccb1f17f448628f *Deque) __obf_9c578e473cd91a70(__obf_35fc60e92e4b05ba int) int {
	return (__obf_35fc60e92e4b05ba + 1) & (len(__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8) - 1) // bitwise modulus
}

// growIfFull resizes up if the buffer is full.
func (__obf_cccb1f17f448628f *Deque) __obf_5c24d5ef59e77124() {
	if len(__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8) == 0 {
		if __obf_cccb1f17f448628f.__obf_c3af9f3aa40e7102 == 0 {
			__obf_cccb1f17f448628f.__obf_c3af9f3aa40e7102 = __obf_cc95479c3d76dafd
		}
		__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8 = make([]any, __obf_cccb1f17f448628f.__obf_c3af9f3aa40e7102)
		return
	}
	if __obf_cccb1f17f448628f.__obf_7e6767c22d23f06d == len(__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8) {
		__obf_cccb1f17f448628f.__obf_5502207518564f83()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_cccb1f17f448628f *Deque) __obf_e286abf03b28bd7a() {
	if len(__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8) > __obf_cccb1f17f448628f.__obf_c3af9f3aa40e7102 && (__obf_cccb1f17f448628f.__obf_7e6767c22d23f06d<<2) == len(__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8) {
		__obf_cccb1f17f448628f.__obf_5502207518564f83()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_cccb1f17f448628f *Deque) __obf_5502207518564f83() {
	__obf_075d33b74bf64552 := make([]any, __obf_cccb1f17f448628f.__obf_7e6767c22d23f06d<<1)
	if __obf_cccb1f17f448628f.__obf_30d187ec07bc80ec > __obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 {
		copy(__obf_075d33b74bf64552, __obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16:__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec])
	} else {
		__obf_bdf68a51ae002dc2 := copy(__obf_075d33b74bf64552, __obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16:])
		copy(__obf_075d33b74bf64552[__obf_bdf68a51ae002dc2:], __obf_cccb1f17f448628f.__obf_22291ba8f5a063e8[:__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec])
	}

	__obf_cccb1f17f448628f.__obf_e7184c1cd4d9cb16 = 0
	__obf_cccb1f17f448628f.__obf_30d187ec07bc80ec = __obf_cccb1f17f448628f.__obf_7e6767c22d23f06d
	__obf_cccb1f17f448628f.__obf_22291ba8f5a063e8 = __obf_075d33b74bf64552
}
