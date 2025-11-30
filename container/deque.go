package __obf_038560a94647875f

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_922a40d0a2c2c23c = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_d962d9abd42389c3 []any
	__obf_9670a3ef26fbfec6 int
	__obf_3569b496ae727405 int
	__obf_8b99337c97e3c5e6 int
	__obf_e678fb92ec87264a int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_70fc50e9effd2311 *Deque) Len() int {
	return __obf_70fc50e9effd2311.

	// PushBack appends an element to the back of the queue.  Implements FIFO when
	// elements are removed with PopFront(), and LIFO when elements are removed
	// with PopBack().
	__obf_8b99337c97e3c5e6
}

func (__obf_70fc50e9effd2311 *Deque) PushBack(__obf_3c623453454fedc7 any) {
	__obf_70fc50e9effd2311.__obf_25b58779b8f662ee()
	__obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.
	// Calculate new tail position.
	__obf_3569b496ae727405] = __obf_3c623453454fedc7
	__obf_70fc50e9effd2311.__obf_3569b496ae727405 = __obf_70fc50e9effd2311.__obf_612bb80bc1fd44c4(__obf_70fc50e9effd2311.__obf_3569b496ae727405)
	__obf_70fc50e9effd2311.

	// PushFront prepends an element to the front of the queue.
	__obf_8b99337c97e3c5e6++
}

func (__obf_70fc50e9effd2311 *Deque) PushFront(__obf_3c623453454fedc7 any) {
	__obf_70fc50e9effd2311.

	// Calculate new head position.
	__obf_25b58779b8f662ee()
	__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6 = __obf_70fc50e9effd2311.__obf_9432ad7a4ca8c46e(__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6)
	__obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6] = __obf_3c623453454fedc7
	__obf_70fc50e9effd2311.

	// PopFront removes and returns the element from the front of the queue.
	// Implements FIFO when used with PushBack().  If the queue is empty, the call
	// panics.
	__obf_8b99337c97e3c5e6++
}

func (__obf_70fc50e9effd2311 *Deque) PopFront() any {
	if __obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6 <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_e66bdf08a044ad76 := __obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6]
	__obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.
	// Calculate new head position.
	__obf_9670a3ef26fbfec6] = nil
	__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6 = __obf_70fc50e9effd2311.__obf_612bb80bc1fd44c4(__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6)
	__obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6--
	__obf_70fc50e9effd2311.__obf_f0e76b59f1a6e2d5()
	return __obf_e66bdf08a044ad76
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_70fc50e9effd2311 *Deque) PopBack() any {
	if __obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6 <= 0 {
		panic("deque: PopBack() called on empty queue")
	}
	__obf_70fc50e9effd2311.

	// Calculate new tail position
	__obf_3569b496ae727405 = __obf_70fc50e9effd2311.

	// Remove value at tail.
	__obf_9432ad7a4ca8c46e(__obf_70fc50e9effd2311.__obf_3569b496ae727405)
	__obf_e66bdf08a044ad76 := __obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_3569b496ae727405]
	__obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_3569b496ae727405] = nil
	__obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6--
	__obf_70fc50e9effd2311.__obf_f0e76b59f1a6e2d5()
	return __obf_e66bdf08a044ad76
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_70fc50e9effd2311 *Deque) Front() any {
	if __obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6 <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_70fc50e9effd2311.

	// Back returns the element at the back of the queue.  This is the element
	// that would be returned by PopBack().  This call panics if the queue is
	// empty.
	__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6]
}

func (__obf_70fc50e9effd2311 *Deque) Back() any {
	if __obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6 <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.

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
	__obf_9432ad7a4ca8c46e(__obf_70fc50e9effd2311.__obf_3569b496ae727405)]
}

func (__obf_70fc50e9effd2311 *Deque) At(__obf_d2ea251c77215677 int) any {
	if __obf_d2ea251c77215677 < 0 || __obf_d2ea251c77215677 >= __obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6 {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_70fc50e9effd2311.__obf_d962d9abd42389c3[(__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6+__obf_d2ea251c77215677)&(len(__obf_70fc50e9effd2311.

	// Clear removes all elements from the queue, but retains the current capacity.
	// This is useful when repeatedly reusing the queue at high frequency to avoid
	// GC during reuse.  The queue will not be resized smaller as long as items are
	// only added.  Only when items are removed is the queue subject to getting
	// resized smaller.
	__obf_d962d9abd42389c3)-1)]
}

func (__obf_70fc50e9effd2311 *Deque) Clear() {
	__obf_742e131028675b3c :=// bitwise modulus
	len(__obf_70fc50e9effd2311.__obf_d962d9abd42389c3) - 1
	for __obf_85e761a4baea4bf9 := __obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6; __obf_85e761a4baea4bf9 != __obf_70fc50e9effd2311.__obf_3569b496ae727405; __obf_85e761a4baea4bf9 = (__obf_85e761a4baea4bf9 + 1) & __obf_742e131028675b3c {
		__obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_85e761a4baea4bf9] = nil
	}
	__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6 = 0
	__obf_70fc50e9effd2311.__obf_3569b496ae727405 = 0
	__obf_70fc50e9effd2311.

	// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
	// back-to-front.  Having Deque provide Rotate() avoids resizing that could
	// happen if implementing rotation using only Pop and Push methods.
	__obf_8b99337c97e3c5e6 = 0
}

func (__obf_70fc50e9effd2311 *Deque) Rotate(__obf_406632683c089fe0 int) {
	if __obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6 <= 1 {
		return
	}
	__obf_406632683c089fe0 %=// Rotating a multiple of q.count is same as no rotation.
	__obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6
	if __obf_406632683c089fe0 == 0 {
		return
	}
	__obf_742e131028675b3c := len(__obf_70fc50e9effd2311.
	// If no empty space in buffer, only move head and tail indexes.
	__obf_d962d9abd42389c3) - 1

	if __obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6 == __obf_70fc50e9effd2311.
	// Calculate new head and tail using bitwise modulus.
	__obf_3569b496ae727405 {
		__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6 = (__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6 + __obf_406632683c089fe0) & __obf_742e131028675b3c
		__obf_70fc50e9effd2311.__obf_3569b496ae727405 = (__obf_70fc50e9effd2311.__obf_3569b496ae727405 + __obf_406632683c089fe0) & __obf_742e131028675b3c
		return
	}

	if __obf_406632683c089fe0 < 0 {
		// Rotate back to front.
		for ; __obf_406632683c089fe0 < 0; __obf_406632683c089fe0++ {
			__obf_70fc50e9effd2311.
			// Calculate new head and tail using bitwise modulus.
			__obf_9670a3ef26fbfec6 = (__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6 - 1) & __obf_742e131028675b3c
			__obf_70fc50e9effd2311.__obf_3569b496ae727405 = (__obf_70fc50e9effd2311.__obf_3569b496ae727405 - 1) & __obf_742e131028675b3c
			// Put tail value at head and remove value at tail.
			__obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6] = __obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_3569b496ae727405]
			__obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_3569b496ae727405] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_406632683c089fe0 > 0; __obf_406632683c089fe0-- {
		__obf_70fc50e9effd2311.
		// Put head value at tail and remove value at head.
		__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_3569b496ae727405] = __obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6]
		__obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.
		// Calculate new head and tail using bitwise modulus.
		__obf_9670a3ef26fbfec6] = nil
		__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6 = (__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6 + 1) & __obf_742e131028675b3c
		__obf_70fc50e9effd2311.__obf_3569b496ae727405 = (__obf_70fc50e9effd2311.__obf_3569b496ae727405 + 1) & __obf_742e131028675b3c
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_70fc50e9effd2311 *Deque) SetMinCapacity(__obf_5fa8ca8310e7c164 uint) {
	if 1<<__obf_5fa8ca8310e7c164 > __obf_922a40d0a2c2c23c {
		__obf_70fc50e9effd2311.__obf_e678fb92ec87264a = 1 << __obf_5fa8ca8310e7c164
	} else {
		__obf_70fc50e9effd2311.__obf_e678fb92ec87264a = __obf_922a40d0a2c2c23c
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_70fc50e9effd2311 *Deque) __obf_9432ad7a4ca8c46e(__obf_d2ea251c77215677 int) int {
	return (__obf_d2ea251c77215677 - 1) & (len(__obf_70fc50e9effd2311. // bitwise modulus
	__obf_d962d9abd42389c3) - 1)
}

// next returns the next buffer position wrapping around buffer.
func (__obf_70fc50e9effd2311 *Deque) __obf_612bb80bc1fd44c4(__obf_d2ea251c77215677 int) int {
	return (__obf_d2ea251c77215677 + 1) & (len(__obf_70fc50e9effd2311. // bitwise modulus
	__obf_d962d9abd42389c3) - 1)
}

// growIfFull resizes up if the buffer is full.
func (__obf_70fc50e9effd2311 *Deque) __obf_25b58779b8f662ee() {
	if len(__obf_70fc50e9effd2311.__obf_d962d9abd42389c3) == 0 {
		if __obf_70fc50e9effd2311.__obf_e678fb92ec87264a == 0 {
			__obf_70fc50e9effd2311.__obf_e678fb92ec87264a = __obf_922a40d0a2c2c23c
		}
		__obf_70fc50e9effd2311.__obf_d962d9abd42389c3 = make([]any, __obf_70fc50e9effd2311.__obf_e678fb92ec87264a)
		return
	}
	if __obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6 == len(__obf_70fc50e9effd2311.__obf_d962d9abd42389c3) {
		__obf_70fc50e9effd2311.__obf_00a0361f8bf3931f()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_70fc50e9effd2311 *Deque) __obf_f0e76b59f1a6e2d5() {
	if len(__obf_70fc50e9effd2311.__obf_d962d9abd42389c3) > __obf_70fc50e9effd2311.__obf_e678fb92ec87264a && (__obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6<<2) == len(__obf_70fc50e9effd2311.__obf_d962d9abd42389c3) {
		__obf_70fc50e9effd2311.__obf_00a0361f8bf3931f()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_70fc50e9effd2311 *Deque) __obf_00a0361f8bf3931f() {
	__obf_728960870055d86a := make([]any, __obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6<<1)
	if __obf_70fc50e9effd2311.__obf_3569b496ae727405 > __obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6 {
		copy(__obf_728960870055d86a, __obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6:__obf_70fc50e9effd2311.__obf_3569b496ae727405])
	} else {
		__obf_406632683c089fe0 := copy(__obf_728960870055d86a, __obf_70fc50e9effd2311.__obf_d962d9abd42389c3[__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6:])
		copy(__obf_728960870055d86a[__obf_406632683c089fe0:], __obf_70fc50e9effd2311.__obf_d962d9abd42389c3[:__obf_70fc50e9effd2311.__obf_3569b496ae727405])
	}
	__obf_70fc50e9effd2311.__obf_9670a3ef26fbfec6 = 0
	__obf_70fc50e9effd2311.__obf_3569b496ae727405 = __obf_70fc50e9effd2311.__obf_8b99337c97e3c5e6
	__obf_70fc50e9effd2311.__obf_d962d9abd42389c3 = __obf_728960870055d86a
}
