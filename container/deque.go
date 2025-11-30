package __obf_e72ce603d10d02a1

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_e999e7fa52c788c3 = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_2244b235a0b5bae4 []any
	__obf_5ff7aa83a91eea77 int
	__obf_25b94aad123f5c8d int
	__obf_34ac5ddd5470ea6e int
	__obf_98c87c740af863d4 int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_5e9593da6f276918 *Deque) Len() int {
	return __obf_5e9593da6f276918.

	// PushBack appends an element to the back of the queue.  Implements FIFO when
	// elements are removed with PopFront(), and LIFO when elements are removed
	// with PopBack().
	__obf_34ac5ddd5470ea6e
}

func (__obf_5e9593da6f276918 *Deque) PushBack(__obf_9148a001224202b7 any) {
	__obf_5e9593da6f276918.__obf_11cdf8b9ba6217d3()
	__obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.
	// Calculate new tail position.
	__obf_25b94aad123f5c8d] = __obf_9148a001224202b7
	__obf_5e9593da6f276918.__obf_25b94aad123f5c8d = __obf_5e9593da6f276918.__obf_dbfc801eeafa9b57(__obf_5e9593da6f276918.__obf_25b94aad123f5c8d)
	__obf_5e9593da6f276918.

	// PushFront prepends an element to the front of the queue.
	__obf_34ac5ddd5470ea6e++
}

func (__obf_5e9593da6f276918 *Deque) PushFront(__obf_9148a001224202b7 any) {
	__obf_5e9593da6f276918.

	// Calculate new head position.
	__obf_11cdf8b9ba6217d3()
	__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77 = __obf_5e9593da6f276918.__obf_69e5dfbc702ecf34(__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77)
	__obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77] = __obf_9148a001224202b7
	__obf_5e9593da6f276918.

	// PopFront removes and returns the element from the front of the queue.
	// Implements FIFO when used with PushBack().  If the queue is empty, the call
	// panics.
	__obf_34ac5ddd5470ea6e++
}

func (__obf_5e9593da6f276918 *Deque) PopFront() any {
	if __obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_826b47218ad141a4 := __obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77]
	__obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.
	// Calculate new head position.
	__obf_5ff7aa83a91eea77] = nil
	__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77 = __obf_5e9593da6f276918.__obf_dbfc801eeafa9b57(__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77)
	__obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e--
	__obf_5e9593da6f276918.__obf_639e88611397eb90()
	return __obf_826b47218ad141a4
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_5e9593da6f276918 *Deque) PopBack() any {
	if __obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e <= 0 {
		panic("deque: PopBack() called on empty queue")
	}
	__obf_5e9593da6f276918.

	// Calculate new tail position
	__obf_25b94aad123f5c8d = __obf_5e9593da6f276918.

	// Remove value at tail.
	__obf_69e5dfbc702ecf34(__obf_5e9593da6f276918.__obf_25b94aad123f5c8d)
	__obf_826b47218ad141a4 := __obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_25b94aad123f5c8d]
	__obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_25b94aad123f5c8d] = nil
	__obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e--
	__obf_5e9593da6f276918.__obf_639e88611397eb90()
	return __obf_826b47218ad141a4
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_5e9593da6f276918 *Deque) Front() any {
	if __obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_5e9593da6f276918.

	// Back returns the element at the back of the queue.  This is the element
	// that would be returned by PopBack().  This call panics if the queue is
	// empty.
	__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77]
}

func (__obf_5e9593da6f276918 *Deque) Back() any {
	if __obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.

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
	__obf_69e5dfbc702ecf34(__obf_5e9593da6f276918.__obf_25b94aad123f5c8d)]
}

func (__obf_5e9593da6f276918 *Deque) At(__obf_c03cf6504cc3cf88 int) any {
	if __obf_c03cf6504cc3cf88 < 0 || __obf_c03cf6504cc3cf88 >= __obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_5e9593da6f276918.__obf_2244b235a0b5bae4[(__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77+__obf_c03cf6504cc3cf88)&(len(__obf_5e9593da6f276918.

	// Clear removes all elements from the queue, but retains the current capacity.
	// This is useful when repeatedly reusing the queue at high frequency to avoid
	// GC during reuse.  The queue will not be resized smaller as long as items are
	// only added.  Only when items are removed is the queue subject to getting
	// resized smaller.
	__obf_2244b235a0b5bae4)-1)]
}

func (__obf_5e9593da6f276918 *Deque) Clear() {
	__obf_7415ae619987504f :=// bitwise modulus
	len(__obf_5e9593da6f276918.__obf_2244b235a0b5bae4) - 1
	for __obf_41a6f945d68bc803 := __obf_5e9593da6f276918.__obf_5ff7aa83a91eea77; __obf_41a6f945d68bc803 != __obf_5e9593da6f276918.__obf_25b94aad123f5c8d; __obf_41a6f945d68bc803 = (__obf_41a6f945d68bc803 + 1) & __obf_7415ae619987504f {
		__obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_41a6f945d68bc803] = nil
	}
	__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77 = 0
	__obf_5e9593da6f276918.__obf_25b94aad123f5c8d = 0
	__obf_5e9593da6f276918.

	// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
	// back-to-front.  Having Deque provide Rotate() avoids resizing that could
	// happen if implementing rotation using only Pop and Push methods.
	__obf_34ac5ddd5470ea6e = 0
}

func (__obf_5e9593da6f276918 *Deque) Rotate(__obf_59eea0bd834c59bb int) {
	if __obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e <= 1 {
		return
	}
	__obf_59eea0bd834c59bb %=// Rotating a multiple of q.count is same as no rotation.
	__obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e
	if __obf_59eea0bd834c59bb == 0 {
		return
	}
	__obf_7415ae619987504f := len(__obf_5e9593da6f276918.
	// If no empty space in buffer, only move head and tail indexes.
	__obf_2244b235a0b5bae4) - 1

	if __obf_5e9593da6f276918.__obf_5ff7aa83a91eea77 == __obf_5e9593da6f276918.
	// Calculate new head and tail using bitwise modulus.
	__obf_25b94aad123f5c8d {
		__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77 = (__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77 + __obf_59eea0bd834c59bb) & __obf_7415ae619987504f
		__obf_5e9593da6f276918.__obf_25b94aad123f5c8d = (__obf_5e9593da6f276918.__obf_25b94aad123f5c8d + __obf_59eea0bd834c59bb) & __obf_7415ae619987504f
		return
	}

	if __obf_59eea0bd834c59bb < 0 {
		// Rotate back to front.
		for ; __obf_59eea0bd834c59bb < 0; __obf_59eea0bd834c59bb++ {
			__obf_5e9593da6f276918.
			// Calculate new head and tail using bitwise modulus.
			__obf_5ff7aa83a91eea77 = (__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77 - 1) & __obf_7415ae619987504f
			__obf_5e9593da6f276918.__obf_25b94aad123f5c8d = (__obf_5e9593da6f276918.__obf_25b94aad123f5c8d - 1) & __obf_7415ae619987504f
			// Put tail value at head and remove value at tail.
			__obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77] = __obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_25b94aad123f5c8d]
			__obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_25b94aad123f5c8d] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_59eea0bd834c59bb > 0; __obf_59eea0bd834c59bb-- {
		__obf_5e9593da6f276918.
		// Put head value at tail and remove value at head.
		__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_25b94aad123f5c8d] = __obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77]
		__obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.
		// Calculate new head and tail using bitwise modulus.
		__obf_5ff7aa83a91eea77] = nil
		__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77 = (__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77 + 1) & __obf_7415ae619987504f
		__obf_5e9593da6f276918.__obf_25b94aad123f5c8d = (__obf_5e9593da6f276918.__obf_25b94aad123f5c8d + 1) & __obf_7415ae619987504f
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_5e9593da6f276918 *Deque) SetMinCapacity(__obf_64d5c8ea4738f042 uint) {
	if 1<<__obf_64d5c8ea4738f042 > __obf_e999e7fa52c788c3 {
		__obf_5e9593da6f276918.__obf_98c87c740af863d4 = 1 << __obf_64d5c8ea4738f042
	} else {
		__obf_5e9593da6f276918.__obf_98c87c740af863d4 = __obf_e999e7fa52c788c3
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_5e9593da6f276918 *Deque) __obf_69e5dfbc702ecf34(__obf_c03cf6504cc3cf88 int) int {
	return (__obf_c03cf6504cc3cf88 - 1) & (len(__obf_5e9593da6f276918. // bitwise modulus
	__obf_2244b235a0b5bae4) - 1)
}

// next returns the next buffer position wrapping around buffer.
func (__obf_5e9593da6f276918 *Deque) __obf_dbfc801eeafa9b57(__obf_c03cf6504cc3cf88 int) int {
	return (__obf_c03cf6504cc3cf88 + 1) & (len(__obf_5e9593da6f276918. // bitwise modulus
	__obf_2244b235a0b5bae4) - 1)
}

// growIfFull resizes up if the buffer is full.
func (__obf_5e9593da6f276918 *Deque) __obf_11cdf8b9ba6217d3() {
	if len(__obf_5e9593da6f276918.__obf_2244b235a0b5bae4) == 0 {
		if __obf_5e9593da6f276918.__obf_98c87c740af863d4 == 0 {
			__obf_5e9593da6f276918.__obf_98c87c740af863d4 = __obf_e999e7fa52c788c3
		}
		__obf_5e9593da6f276918.__obf_2244b235a0b5bae4 = make([]any, __obf_5e9593da6f276918.__obf_98c87c740af863d4)
		return
	}
	if __obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e == len(__obf_5e9593da6f276918.__obf_2244b235a0b5bae4) {
		__obf_5e9593da6f276918.__obf_5c2503c5604068da()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_5e9593da6f276918 *Deque) __obf_639e88611397eb90() {
	if len(__obf_5e9593da6f276918.__obf_2244b235a0b5bae4) > __obf_5e9593da6f276918.__obf_98c87c740af863d4 && (__obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e<<2) == len(__obf_5e9593da6f276918.__obf_2244b235a0b5bae4) {
		__obf_5e9593da6f276918.__obf_5c2503c5604068da()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_5e9593da6f276918 *Deque) __obf_5c2503c5604068da() {
	__obf_9d9512acba1024d1 := make([]any, __obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e<<1)
	if __obf_5e9593da6f276918.__obf_25b94aad123f5c8d > __obf_5e9593da6f276918.__obf_5ff7aa83a91eea77 {
		copy(__obf_9d9512acba1024d1, __obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77:__obf_5e9593da6f276918.__obf_25b94aad123f5c8d])
	} else {
		__obf_59eea0bd834c59bb := copy(__obf_9d9512acba1024d1, __obf_5e9593da6f276918.__obf_2244b235a0b5bae4[__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77:])
		copy(__obf_9d9512acba1024d1[__obf_59eea0bd834c59bb:], __obf_5e9593da6f276918.__obf_2244b235a0b5bae4[:__obf_5e9593da6f276918.__obf_25b94aad123f5c8d])
	}
	__obf_5e9593da6f276918.__obf_5ff7aa83a91eea77 = 0
	__obf_5e9593da6f276918.__obf_25b94aad123f5c8d = __obf_5e9593da6f276918.__obf_34ac5ddd5470ea6e
	__obf_5e9593da6f276918.__obf_2244b235a0b5bae4 = __obf_9d9512acba1024d1
}
