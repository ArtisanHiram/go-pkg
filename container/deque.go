package __obf_4a16ef421fb74992

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_50d5bde4e3410864 = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_561436743d5f7317 []any
	__obf_e0a663f0cac8e1ef int
	__obf_00349c0f12dbb91a int
	__obf_1d7629104d39aeed int
	__obf_c8228109fcc7e7db int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_94b2a7432790879c *Deque) Len() int {
	return __obf_94b2a7432790879c.

	// PushBack appends an element to the back of the queue.  Implements FIFO when
	// elements are removed with PopFront(), and LIFO when elements are removed
	// with PopBack().
	__obf_1d7629104d39aeed
}

func (__obf_94b2a7432790879c *Deque) PushBack(__obf_ef0e2650e203f28e any) {
	__obf_94b2a7432790879c.__obf_b67e81b7577c3ca7()
	__obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.
	// Calculate new tail position.
	__obf_00349c0f12dbb91a] = __obf_ef0e2650e203f28e
	__obf_94b2a7432790879c.__obf_00349c0f12dbb91a = __obf_94b2a7432790879c.__obf_2c113f310b7cc228(__obf_94b2a7432790879c.__obf_00349c0f12dbb91a)
	__obf_94b2a7432790879c.

	// PushFront prepends an element to the front of the queue.
	__obf_1d7629104d39aeed++
}

func (__obf_94b2a7432790879c *Deque) PushFront(__obf_ef0e2650e203f28e any) {
	__obf_94b2a7432790879c.

	// Calculate new head position.
	__obf_b67e81b7577c3ca7()
	__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef = __obf_94b2a7432790879c.__obf_7b497de6e96482de(__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef)
	__obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef] = __obf_ef0e2650e203f28e
	__obf_94b2a7432790879c.

	// PopFront removes and returns the element from the front of the queue.
	// Implements FIFO when used with PushBack().  If the queue is empty, the call
	// panics.
	__obf_1d7629104d39aeed++
}

func (__obf_94b2a7432790879c *Deque) PopFront() any {
	if __obf_94b2a7432790879c.__obf_1d7629104d39aeed <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_b820bc1fe8762c40 := __obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef]
	__obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.
	// Calculate new head position.
	__obf_e0a663f0cac8e1ef] = nil
	__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef = __obf_94b2a7432790879c.__obf_2c113f310b7cc228(__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef)
	__obf_94b2a7432790879c.__obf_1d7629104d39aeed--
	__obf_94b2a7432790879c.__obf_009a01887696e6e2()
	return __obf_b820bc1fe8762c40
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_94b2a7432790879c *Deque) PopBack() any {
	if __obf_94b2a7432790879c.__obf_1d7629104d39aeed <= 0 {
		panic("deque: PopBack() called on empty queue")
	}
	__obf_94b2a7432790879c.

	// Calculate new tail position
	__obf_00349c0f12dbb91a = __obf_94b2a7432790879c.

	// Remove value at tail.
	__obf_7b497de6e96482de(__obf_94b2a7432790879c.__obf_00349c0f12dbb91a)
	__obf_b820bc1fe8762c40 := __obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_00349c0f12dbb91a]
	__obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_00349c0f12dbb91a] = nil
	__obf_94b2a7432790879c.__obf_1d7629104d39aeed--
	__obf_94b2a7432790879c.__obf_009a01887696e6e2()
	return __obf_b820bc1fe8762c40
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_94b2a7432790879c *Deque) Front() any {
	if __obf_94b2a7432790879c.__obf_1d7629104d39aeed <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_94b2a7432790879c.

	// Back returns the element at the back of the queue.  This is the element
	// that would be returned by PopBack().  This call panics if the queue is
	// empty.
	__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef]
}

func (__obf_94b2a7432790879c *Deque) Back() any {
	if __obf_94b2a7432790879c.__obf_1d7629104d39aeed <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.

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
	__obf_7b497de6e96482de(__obf_94b2a7432790879c.__obf_00349c0f12dbb91a)]
}

func (__obf_94b2a7432790879c *Deque) At(__obf_79cf1c44489eaf01 int) any {
	if __obf_79cf1c44489eaf01 < 0 || __obf_79cf1c44489eaf01 >= __obf_94b2a7432790879c.__obf_1d7629104d39aeed {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_94b2a7432790879c.__obf_561436743d5f7317[(__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef+__obf_79cf1c44489eaf01)&(len(__obf_94b2a7432790879c.

	// Clear removes all elements from the queue, but retains the current capacity.
	// This is useful when repeatedly reusing the queue at high frequency to avoid
	// GC during reuse.  The queue will not be resized smaller as long as items are
	// only added.  Only when items are removed is the queue subject to getting
	// resized smaller.
	__obf_561436743d5f7317)-1)]
}

func (__obf_94b2a7432790879c *Deque) Clear() {
	__obf_a0937d202098727c :=// bitwise modulus
	len(__obf_94b2a7432790879c.__obf_561436743d5f7317) - 1
	for __obf_55fe93b5d3e19f79 := __obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef; __obf_55fe93b5d3e19f79 != __obf_94b2a7432790879c.__obf_00349c0f12dbb91a; __obf_55fe93b5d3e19f79 = (__obf_55fe93b5d3e19f79 + 1) & __obf_a0937d202098727c {
		__obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_55fe93b5d3e19f79] = nil
	}
	__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef = 0
	__obf_94b2a7432790879c.__obf_00349c0f12dbb91a = 0
	__obf_94b2a7432790879c.

	// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
	// back-to-front.  Having Deque provide Rotate() avoids resizing that could
	// happen if implementing rotation using only Pop and Push methods.
	__obf_1d7629104d39aeed = 0
}

func (__obf_94b2a7432790879c *Deque) Rotate(__obf_41509bed52f90987 int) {
	if __obf_94b2a7432790879c.__obf_1d7629104d39aeed <= 1 {
		return
	}
	__obf_41509bed52f90987 %=// Rotating a multiple of q.count is same as no rotation.
	__obf_94b2a7432790879c.__obf_1d7629104d39aeed
	if __obf_41509bed52f90987 == 0 {
		return
	}
	__obf_a0937d202098727c := len(__obf_94b2a7432790879c.
	// If no empty space in buffer, only move head and tail indexes.
	__obf_561436743d5f7317) - 1

	if __obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef == __obf_94b2a7432790879c.
	// Calculate new head and tail using bitwise modulus.
	__obf_00349c0f12dbb91a {
		__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef = (__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef + __obf_41509bed52f90987) & __obf_a0937d202098727c
		__obf_94b2a7432790879c.__obf_00349c0f12dbb91a = (__obf_94b2a7432790879c.__obf_00349c0f12dbb91a + __obf_41509bed52f90987) & __obf_a0937d202098727c
		return
	}

	if __obf_41509bed52f90987 < 0 {
		// Rotate back to front.
		for ; __obf_41509bed52f90987 < 0; __obf_41509bed52f90987++ {
			__obf_94b2a7432790879c.
			// Calculate new head and tail using bitwise modulus.
			__obf_e0a663f0cac8e1ef = (__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef - 1) & __obf_a0937d202098727c
			__obf_94b2a7432790879c.__obf_00349c0f12dbb91a = (__obf_94b2a7432790879c.__obf_00349c0f12dbb91a - 1) & __obf_a0937d202098727c
			// Put tail value at head and remove value at tail.
			__obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef] = __obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_00349c0f12dbb91a]
			__obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_00349c0f12dbb91a] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_41509bed52f90987 > 0; __obf_41509bed52f90987-- {
		__obf_94b2a7432790879c.
		// Put head value at tail and remove value at head.
		__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_00349c0f12dbb91a] = __obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef]
		__obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.
		// Calculate new head and tail using bitwise modulus.
		__obf_e0a663f0cac8e1ef] = nil
		__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef = (__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef + 1) & __obf_a0937d202098727c
		__obf_94b2a7432790879c.__obf_00349c0f12dbb91a = (__obf_94b2a7432790879c.__obf_00349c0f12dbb91a + 1) & __obf_a0937d202098727c
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_94b2a7432790879c *Deque) SetMinCapacity(__obf_98e35237c6553725 uint) {
	if 1<<__obf_98e35237c6553725 > __obf_50d5bde4e3410864 {
		__obf_94b2a7432790879c.__obf_c8228109fcc7e7db = 1 << __obf_98e35237c6553725
	} else {
		__obf_94b2a7432790879c.__obf_c8228109fcc7e7db = __obf_50d5bde4e3410864
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_94b2a7432790879c *Deque) __obf_7b497de6e96482de(__obf_79cf1c44489eaf01 int) int {
	return (__obf_79cf1c44489eaf01 - 1) & (len(__obf_94b2a7432790879c. // bitwise modulus
	__obf_561436743d5f7317) - 1)
}

// next returns the next buffer position wrapping around buffer.
func (__obf_94b2a7432790879c *Deque) __obf_2c113f310b7cc228(__obf_79cf1c44489eaf01 int) int {
	return (__obf_79cf1c44489eaf01 + 1) & (len(__obf_94b2a7432790879c. // bitwise modulus
	__obf_561436743d5f7317) - 1)
}

// growIfFull resizes up if the buffer is full.
func (__obf_94b2a7432790879c *Deque) __obf_b67e81b7577c3ca7() {
	if len(__obf_94b2a7432790879c.__obf_561436743d5f7317) == 0 {
		if __obf_94b2a7432790879c.__obf_c8228109fcc7e7db == 0 {
			__obf_94b2a7432790879c.__obf_c8228109fcc7e7db = __obf_50d5bde4e3410864
		}
		__obf_94b2a7432790879c.__obf_561436743d5f7317 = make([]any, __obf_94b2a7432790879c.__obf_c8228109fcc7e7db)
		return
	}
	if __obf_94b2a7432790879c.__obf_1d7629104d39aeed == len(__obf_94b2a7432790879c.__obf_561436743d5f7317) {
		__obf_94b2a7432790879c.__obf_f821a6689f0ff3d9()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_94b2a7432790879c *Deque) __obf_009a01887696e6e2() {
	if len(__obf_94b2a7432790879c.__obf_561436743d5f7317) > __obf_94b2a7432790879c.__obf_c8228109fcc7e7db && (__obf_94b2a7432790879c.__obf_1d7629104d39aeed<<2) == len(__obf_94b2a7432790879c.__obf_561436743d5f7317) {
		__obf_94b2a7432790879c.__obf_f821a6689f0ff3d9()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_94b2a7432790879c *Deque) __obf_f821a6689f0ff3d9() {
	__obf_b1dc4cf28b0bb502 := make([]any, __obf_94b2a7432790879c.__obf_1d7629104d39aeed<<1)
	if __obf_94b2a7432790879c.__obf_00349c0f12dbb91a > __obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef {
		copy(__obf_b1dc4cf28b0bb502, __obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef:__obf_94b2a7432790879c.__obf_00349c0f12dbb91a])
	} else {
		__obf_41509bed52f90987 := copy(__obf_b1dc4cf28b0bb502, __obf_94b2a7432790879c.__obf_561436743d5f7317[__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef:])
		copy(__obf_b1dc4cf28b0bb502[__obf_41509bed52f90987:], __obf_94b2a7432790879c.__obf_561436743d5f7317[:__obf_94b2a7432790879c.__obf_00349c0f12dbb91a])
	}
	__obf_94b2a7432790879c.__obf_e0a663f0cac8e1ef = 0
	__obf_94b2a7432790879c.__obf_00349c0f12dbb91a = __obf_94b2a7432790879c.__obf_1d7629104d39aeed
	__obf_94b2a7432790879c.__obf_561436743d5f7317 = __obf_b1dc4cf28b0bb502
}
