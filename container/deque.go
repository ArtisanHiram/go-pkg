package __obf_af42fb6cde2beed6

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_7d1a583453f8afa5 = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_26ef969b00424002 []any
	__obf_946b869070fbffee int
	__obf_8e599326d6c7e858 int
	__obf_b4fd98c1e0169c13 int
	__obf_36b025bd303d8a94 int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_ce73249ea3672f2d *Deque) Len() int {
	return __obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13
}

// PushBack appends an element to the back of the queue.  Implements FIFO when
// elements are removed with PopFront(), and LIFO when elements are removed
// with PopBack().
func (__obf_ce73249ea3672f2d *Deque) PushBack(__obf_c58444540a03feb7 any) {
	__obf_ce73249ea3672f2d.__obf_d88ce6f9569de89b()

	__obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858] = __obf_c58444540a03feb7
	// Calculate new tail position.
	__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 = __obf_ce73249ea3672f2d.__obf_3c691a3fec1b7fd4(__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858)
	__obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13++
}

// PushFront prepends an element to the front of the queue.
func (__obf_ce73249ea3672f2d *Deque) PushFront(__obf_c58444540a03feb7 any) {
	__obf_ce73249ea3672f2d.__obf_d88ce6f9569de89b()

	// Calculate new head position.
	__obf_ce73249ea3672f2d.__obf_946b869070fbffee = __obf_ce73249ea3672f2d.__obf_29579fcd202167dd(__obf_ce73249ea3672f2d.__obf_946b869070fbffee)
	__obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_946b869070fbffee] = __obf_c58444540a03feb7
	__obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13++
}

// PopFront removes and returns the element from the front of the queue.
// Implements FIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_ce73249ea3672f2d *Deque) PopFront() any {
	if __obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13 <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_5e29731448e4d5d2 := __obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_946b869070fbffee]
	__obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_946b869070fbffee] = nil
	// Calculate new head position.
	__obf_ce73249ea3672f2d.__obf_946b869070fbffee = __obf_ce73249ea3672f2d.__obf_3c691a3fec1b7fd4(__obf_ce73249ea3672f2d.__obf_946b869070fbffee)
	__obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13--

	__obf_ce73249ea3672f2d.__obf_ec5a0d4f45a10e05()
	return __obf_5e29731448e4d5d2
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_ce73249ea3672f2d *Deque) PopBack() any {
	if __obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13 <= 0 {
		panic("deque: PopBack() called on empty queue")
	}

	// Calculate new tail position
	__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 = __obf_ce73249ea3672f2d.__obf_29579fcd202167dd(__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858)

	// Remove value at tail.
	__obf_5e29731448e4d5d2 := __obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858]
	__obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858] = nil
	__obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13--

	__obf_ce73249ea3672f2d.__obf_ec5a0d4f45a10e05()
	return __obf_5e29731448e4d5d2
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_ce73249ea3672f2d *Deque) Front() any {
	if __obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13 <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_946b869070fbffee]
}

// Back returns the element at the back of the queue.  This is the element
// that would be returned by PopBack().  This call panics if the queue is
// empty.
func (__obf_ce73249ea3672f2d *Deque) Back() any {
	if __obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13 <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_29579fcd202167dd(__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858)]
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
func (__obf_ce73249ea3672f2d *Deque) At(__obf_d6db326dc08be53b int) any {
	if __obf_d6db326dc08be53b < 0 || __obf_d6db326dc08be53b >= __obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13 {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_ce73249ea3672f2d.__obf_26ef969b00424002[(__obf_ce73249ea3672f2d.__obf_946b869070fbffee+__obf_d6db326dc08be53b)&(len(__obf_ce73249ea3672f2d.__obf_26ef969b00424002)-1)]
}

// Clear removes all elements from the queue, but retains the current capacity.
// This is useful when repeatedly reusing the queue at high frequency to avoid
// GC during reuse.  The queue will not be resized smaller as long as items are
// only added.  Only when items are removed is the queue subject to getting
// resized smaller.
func (__obf_ce73249ea3672f2d *Deque) Clear() {
	// bitwise modulus
	__obf_c85f8fa339e17ab6 := len(__obf_ce73249ea3672f2d.__obf_26ef969b00424002) - 1
	for __obf_146e082d3ec5bb18 := __obf_ce73249ea3672f2d.__obf_946b869070fbffee; __obf_146e082d3ec5bb18 != __obf_ce73249ea3672f2d.__obf_8e599326d6c7e858; __obf_146e082d3ec5bb18 = (__obf_146e082d3ec5bb18 + 1) & __obf_c85f8fa339e17ab6 {
		__obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_146e082d3ec5bb18] = nil
	}
	__obf_ce73249ea3672f2d.__obf_946b869070fbffee = 0
	__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 = 0
	__obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13 = 0
}

// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
// back-to-front.  Having Deque provide Rotate() avoids resizing that could
// happen if implementing rotation using only Pop and Push methods.
func (__obf_ce73249ea3672f2d *Deque) Rotate(__obf_da62d4b37cd624ff int) {
	if __obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13 <= 1 {
		return
	}
	// Rotating a multiple of q.count is same as no rotation.
	__obf_da62d4b37cd624ff %= __obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13
	if __obf_da62d4b37cd624ff == 0 {
		return
	}

	__obf_c85f8fa339e17ab6 := len(__obf_ce73249ea3672f2d.__obf_26ef969b00424002) - 1
	// If no empty space in buffer, only move head and tail indexes.
	if __obf_ce73249ea3672f2d.__obf_946b869070fbffee == __obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 {
		// Calculate new head and tail using bitwise modulus.
		__obf_ce73249ea3672f2d.__obf_946b869070fbffee = (__obf_ce73249ea3672f2d.__obf_946b869070fbffee + __obf_da62d4b37cd624ff) & __obf_c85f8fa339e17ab6
		__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 = (__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 + __obf_da62d4b37cd624ff) & __obf_c85f8fa339e17ab6
		return
	}

	if __obf_da62d4b37cd624ff < 0 {
		// Rotate back to front.
		for ; __obf_da62d4b37cd624ff < 0; __obf_da62d4b37cd624ff++ {
			// Calculate new head and tail using bitwise modulus.
			__obf_ce73249ea3672f2d.__obf_946b869070fbffee = (__obf_ce73249ea3672f2d.__obf_946b869070fbffee - 1) & __obf_c85f8fa339e17ab6
			__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 = (__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 - 1) & __obf_c85f8fa339e17ab6
			// Put tail value at head and remove value at tail.
			__obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_946b869070fbffee] = __obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858]
			__obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_da62d4b37cd624ff > 0; __obf_da62d4b37cd624ff-- {
		// Put head value at tail and remove value at head.
		__obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858] = __obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_946b869070fbffee]
		__obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_946b869070fbffee] = nil
		// Calculate new head and tail using bitwise modulus.
		__obf_ce73249ea3672f2d.__obf_946b869070fbffee = (__obf_ce73249ea3672f2d.__obf_946b869070fbffee + 1) & __obf_c85f8fa339e17ab6
		__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 = (__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 + 1) & __obf_c85f8fa339e17ab6
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_ce73249ea3672f2d *Deque) SetMinCapacity(__obf_4f1835b4db4b385a uint) {
	if 1<<__obf_4f1835b4db4b385a > __obf_7d1a583453f8afa5 {
		__obf_ce73249ea3672f2d.__obf_36b025bd303d8a94 = 1 << __obf_4f1835b4db4b385a
	} else {
		__obf_ce73249ea3672f2d.__obf_36b025bd303d8a94 = __obf_7d1a583453f8afa5
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_ce73249ea3672f2d *Deque) __obf_29579fcd202167dd(__obf_d6db326dc08be53b int) int {
	return (__obf_d6db326dc08be53b - 1) & (len(__obf_ce73249ea3672f2d.__obf_26ef969b00424002) - 1) // bitwise modulus
}

// next returns the next buffer position wrapping around buffer.
func (__obf_ce73249ea3672f2d *Deque) __obf_3c691a3fec1b7fd4(__obf_d6db326dc08be53b int) int {
	return (__obf_d6db326dc08be53b + 1) & (len(__obf_ce73249ea3672f2d.__obf_26ef969b00424002) - 1) // bitwise modulus
}

// growIfFull resizes up if the buffer is full.
func (__obf_ce73249ea3672f2d *Deque) __obf_d88ce6f9569de89b() {
	if len(__obf_ce73249ea3672f2d.__obf_26ef969b00424002) == 0 {
		if __obf_ce73249ea3672f2d.__obf_36b025bd303d8a94 == 0 {
			__obf_ce73249ea3672f2d.__obf_36b025bd303d8a94 = __obf_7d1a583453f8afa5
		}
		__obf_ce73249ea3672f2d.__obf_26ef969b00424002 = make([]any, __obf_ce73249ea3672f2d.__obf_36b025bd303d8a94)
		return
	}
	if __obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13 == len(__obf_ce73249ea3672f2d.__obf_26ef969b00424002) {
		__obf_ce73249ea3672f2d.__obf_8e68d9f20c615fa9()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_ce73249ea3672f2d *Deque) __obf_ec5a0d4f45a10e05() {
	if len(__obf_ce73249ea3672f2d.__obf_26ef969b00424002) > __obf_ce73249ea3672f2d.__obf_36b025bd303d8a94 && (__obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13<<2) == len(__obf_ce73249ea3672f2d.__obf_26ef969b00424002) {
		__obf_ce73249ea3672f2d.__obf_8e68d9f20c615fa9()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_ce73249ea3672f2d *Deque) __obf_8e68d9f20c615fa9() {
	__obf_29dbf13f0d545224 := make([]any, __obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13<<1)
	if __obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 > __obf_ce73249ea3672f2d.__obf_946b869070fbffee {
		copy(__obf_29dbf13f0d545224, __obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_946b869070fbffee:__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858])
	} else {
		__obf_da62d4b37cd624ff := copy(__obf_29dbf13f0d545224, __obf_ce73249ea3672f2d.__obf_26ef969b00424002[__obf_ce73249ea3672f2d.__obf_946b869070fbffee:])
		copy(__obf_29dbf13f0d545224[__obf_da62d4b37cd624ff:], __obf_ce73249ea3672f2d.__obf_26ef969b00424002[:__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858])
	}

	__obf_ce73249ea3672f2d.__obf_946b869070fbffee = 0
	__obf_ce73249ea3672f2d.__obf_8e599326d6c7e858 = __obf_ce73249ea3672f2d.__obf_b4fd98c1e0169c13
	__obf_ce73249ea3672f2d.__obf_26ef969b00424002 = __obf_29dbf13f0d545224
}
