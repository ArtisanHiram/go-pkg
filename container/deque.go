package __obf_9004b47202f9204b

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_7db7a03f4e37eb16 = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_00070cd910684634 []any
	__obf_8b1fbd32e67b6492 int
	__obf_41d05da135f9faea int
	__obf_4dd2f2e7ef923bcc int
	__obf_5734e8c88484ab71 int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_a3e1a258c167a26d *Deque) Len() int {
	return __obf_a3e1a258c167a26d.

	// PushBack appends an element to the back of the queue.  Implements FIFO when
	// elements are removed with PopFront(), and LIFO when elements are removed
	// with PopBack().
	__obf_4dd2f2e7ef923bcc
}

func (__obf_a3e1a258c167a26d *Deque) PushBack(__obf_74717af4bf261f66 any) {
	__obf_a3e1a258c167a26d.__obf_54c60e961e23de86()
	__obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.
	// Calculate new tail position.
	__obf_41d05da135f9faea] = __obf_74717af4bf261f66
	__obf_a3e1a258c167a26d.__obf_41d05da135f9faea = __obf_a3e1a258c167a26d.__obf_d5e61e639158b14b(__obf_a3e1a258c167a26d.__obf_41d05da135f9faea)
	__obf_a3e1a258c167a26d.

	// PushFront prepends an element to the front of the queue.
	__obf_4dd2f2e7ef923bcc++
}

func (__obf_a3e1a258c167a26d *Deque) PushFront(__obf_74717af4bf261f66 any) {
	__obf_a3e1a258c167a26d.

	// Calculate new head position.
	__obf_54c60e961e23de86()
	__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492 = __obf_a3e1a258c167a26d.__obf_fabc805cdf3486a7(__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492)
	__obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492] = __obf_74717af4bf261f66
	__obf_a3e1a258c167a26d.

	// PopFront removes and returns the element from the front of the queue.
	// Implements FIFO when used with PushBack().  If the queue is empty, the call
	// panics.
	__obf_4dd2f2e7ef923bcc++
}

func (__obf_a3e1a258c167a26d *Deque) PopFront() any {
	if __obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_c4b9466bc144d45d := __obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492]
	__obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.
	// Calculate new head position.
	__obf_8b1fbd32e67b6492] = nil
	__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492 = __obf_a3e1a258c167a26d.__obf_d5e61e639158b14b(__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492)
	__obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc--
	__obf_a3e1a258c167a26d.__obf_215d800001842452()
	return __obf_c4b9466bc144d45d
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_a3e1a258c167a26d *Deque) PopBack() any {
	if __obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc <= 0 {
		panic("deque: PopBack() called on empty queue")
	}
	__obf_a3e1a258c167a26d.

	// Calculate new tail position
	__obf_41d05da135f9faea = __obf_a3e1a258c167a26d.

	// Remove value at tail.
	__obf_fabc805cdf3486a7(__obf_a3e1a258c167a26d.__obf_41d05da135f9faea)
	__obf_c4b9466bc144d45d := __obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_41d05da135f9faea]
	__obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_41d05da135f9faea] = nil
	__obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc--
	__obf_a3e1a258c167a26d.__obf_215d800001842452()
	return __obf_c4b9466bc144d45d
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_a3e1a258c167a26d *Deque) Front() any {
	if __obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_a3e1a258c167a26d.

	// Back returns the element at the back of the queue.  This is the element
	// that would be returned by PopBack().  This call panics if the queue is
	// empty.
	__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492]
}

func (__obf_a3e1a258c167a26d *Deque) Back() any {
	if __obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.

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
	__obf_fabc805cdf3486a7(__obf_a3e1a258c167a26d.__obf_41d05da135f9faea)]
}

func (__obf_a3e1a258c167a26d *Deque) At(__obf_c97cf0c2d79c5ab6 int) any {
	if __obf_c97cf0c2d79c5ab6 < 0 || __obf_c97cf0c2d79c5ab6 >= __obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_a3e1a258c167a26d.__obf_00070cd910684634[(__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492+__obf_c97cf0c2d79c5ab6)&(len(__obf_a3e1a258c167a26d.

	// Clear removes all elements from the queue, but retains the current capacity.
	// This is useful when repeatedly reusing the queue at high frequency to avoid
	// GC during reuse.  The queue will not be resized smaller as long as items are
	// only added.  Only when items are removed is the queue subject to getting
	// resized smaller.
	__obf_00070cd910684634)-1)]
}

func (__obf_a3e1a258c167a26d *Deque) Clear() {
	__obf_5f184550f8c33981 :=// bitwise modulus
	len(__obf_a3e1a258c167a26d.__obf_00070cd910684634) - 1
	for __obf_e64b3bb2d46916d3 := __obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492; __obf_e64b3bb2d46916d3 != __obf_a3e1a258c167a26d.__obf_41d05da135f9faea; __obf_e64b3bb2d46916d3 = (__obf_e64b3bb2d46916d3 + 1) & __obf_5f184550f8c33981 {
		__obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_e64b3bb2d46916d3] = nil
	}
	__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492 = 0
	__obf_a3e1a258c167a26d.__obf_41d05da135f9faea = 0
	__obf_a3e1a258c167a26d.

	// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
	// back-to-front.  Having Deque provide Rotate() avoids resizing that could
	// happen if implementing rotation using only Pop and Push methods.
	__obf_4dd2f2e7ef923bcc = 0
}

func (__obf_a3e1a258c167a26d *Deque) Rotate(__obf_d2d1bdfa577ba7e8 int) {
	if __obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc <= 1 {
		return
	}
	__obf_d2d1bdfa577ba7e8 %=// Rotating a multiple of q.count is same as no rotation.
	__obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc
	if __obf_d2d1bdfa577ba7e8 == 0 {
		return
	}
	__obf_5f184550f8c33981 := len(__obf_a3e1a258c167a26d.
	// If no empty space in buffer, only move head and tail indexes.
	__obf_00070cd910684634) - 1

	if __obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492 == __obf_a3e1a258c167a26d.
	// Calculate new head and tail using bitwise modulus.
	__obf_41d05da135f9faea {
		__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492 = (__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492 + __obf_d2d1bdfa577ba7e8) & __obf_5f184550f8c33981
		__obf_a3e1a258c167a26d.__obf_41d05da135f9faea = (__obf_a3e1a258c167a26d.__obf_41d05da135f9faea + __obf_d2d1bdfa577ba7e8) & __obf_5f184550f8c33981
		return
	}

	if __obf_d2d1bdfa577ba7e8 < 0 {
		// Rotate back to front.
		for ; __obf_d2d1bdfa577ba7e8 < 0; __obf_d2d1bdfa577ba7e8++ {
			__obf_a3e1a258c167a26d.
			// Calculate new head and tail using bitwise modulus.
			__obf_8b1fbd32e67b6492 = (__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492 - 1) & __obf_5f184550f8c33981
			__obf_a3e1a258c167a26d.__obf_41d05da135f9faea = (__obf_a3e1a258c167a26d.__obf_41d05da135f9faea - 1) & __obf_5f184550f8c33981
			// Put tail value at head and remove value at tail.
			__obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492] = __obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_41d05da135f9faea]
			__obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_41d05da135f9faea] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_d2d1bdfa577ba7e8 > 0; __obf_d2d1bdfa577ba7e8-- {
		__obf_a3e1a258c167a26d.
		// Put head value at tail and remove value at head.
		__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_41d05da135f9faea] = __obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492]
		__obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.
		// Calculate new head and tail using bitwise modulus.
		__obf_8b1fbd32e67b6492] = nil
		__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492 = (__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492 + 1) & __obf_5f184550f8c33981
		__obf_a3e1a258c167a26d.__obf_41d05da135f9faea = (__obf_a3e1a258c167a26d.__obf_41d05da135f9faea + 1) & __obf_5f184550f8c33981
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_a3e1a258c167a26d *Deque) SetMinCapacity(__obf_1c4f65f0866cc1c1 uint) {
	if 1<<__obf_1c4f65f0866cc1c1 > __obf_7db7a03f4e37eb16 {
		__obf_a3e1a258c167a26d.__obf_5734e8c88484ab71 = 1 << __obf_1c4f65f0866cc1c1
	} else {
		__obf_a3e1a258c167a26d.__obf_5734e8c88484ab71 = __obf_7db7a03f4e37eb16
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_a3e1a258c167a26d *Deque) __obf_fabc805cdf3486a7(__obf_c97cf0c2d79c5ab6 int) int {
	return (__obf_c97cf0c2d79c5ab6 - 1) & (len(__obf_a3e1a258c167a26d. // bitwise modulus
	__obf_00070cd910684634) - 1)
}

// next returns the next buffer position wrapping around buffer.
func (__obf_a3e1a258c167a26d *Deque) __obf_d5e61e639158b14b(__obf_c97cf0c2d79c5ab6 int) int {
	return (__obf_c97cf0c2d79c5ab6 + 1) & (len(__obf_a3e1a258c167a26d. // bitwise modulus
	__obf_00070cd910684634) - 1)
}

// growIfFull resizes up if the buffer is full.
func (__obf_a3e1a258c167a26d *Deque) __obf_54c60e961e23de86() {
	if len(__obf_a3e1a258c167a26d.__obf_00070cd910684634) == 0 {
		if __obf_a3e1a258c167a26d.__obf_5734e8c88484ab71 == 0 {
			__obf_a3e1a258c167a26d.__obf_5734e8c88484ab71 = __obf_7db7a03f4e37eb16
		}
		__obf_a3e1a258c167a26d.__obf_00070cd910684634 = make([]any, __obf_a3e1a258c167a26d.__obf_5734e8c88484ab71)
		return
	}
	if __obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc == len(__obf_a3e1a258c167a26d.__obf_00070cd910684634) {
		__obf_a3e1a258c167a26d.__obf_5c35a17ebaf7070a()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_a3e1a258c167a26d *Deque) __obf_215d800001842452() {
	if len(__obf_a3e1a258c167a26d.__obf_00070cd910684634) > __obf_a3e1a258c167a26d.__obf_5734e8c88484ab71 && (__obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc<<2) == len(__obf_a3e1a258c167a26d.__obf_00070cd910684634) {
		__obf_a3e1a258c167a26d.__obf_5c35a17ebaf7070a()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_a3e1a258c167a26d *Deque) __obf_5c35a17ebaf7070a() {
	__obf_19c469dd95ea871d := make([]any, __obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc<<1)
	if __obf_a3e1a258c167a26d.__obf_41d05da135f9faea > __obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492 {
		copy(__obf_19c469dd95ea871d, __obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492:__obf_a3e1a258c167a26d.__obf_41d05da135f9faea])
	} else {
		__obf_d2d1bdfa577ba7e8 := copy(__obf_19c469dd95ea871d, __obf_a3e1a258c167a26d.__obf_00070cd910684634[__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492:])
		copy(__obf_19c469dd95ea871d[__obf_d2d1bdfa577ba7e8:], __obf_a3e1a258c167a26d.__obf_00070cd910684634[:__obf_a3e1a258c167a26d.__obf_41d05da135f9faea])
	}
	__obf_a3e1a258c167a26d.__obf_8b1fbd32e67b6492 = 0
	__obf_a3e1a258c167a26d.__obf_41d05da135f9faea = __obf_a3e1a258c167a26d.__obf_4dd2f2e7ef923bcc
	__obf_a3e1a258c167a26d.__obf_00070cd910684634 = __obf_19c469dd95ea871d
}
