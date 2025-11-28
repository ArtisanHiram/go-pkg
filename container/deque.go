package __obf_1fda7fbdeda52f1e

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_f0d6964507650a12 = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_9bcb11a51dfb8908 []any
	__obf_4fb23df1386611b2 int
	__obf_29adb9bd53a5143c int
	__obf_c29f343498177aa5 int
	__obf_774dfdb134f2399f int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_56903dfea862bbd5 *Deque) Len() int {
	return __obf_56903dfea862bbd5.__obf_c29f343498177aa5
}

// PushBack appends an element to the back of the queue.  Implements FIFO when
// elements are removed with PopFront(), and LIFO when elements are removed
// with PopBack().
func (__obf_56903dfea862bbd5 *Deque) PushBack(__obf_5a12388a905e7c2e any) {
	__obf_56903dfea862bbd5.__obf_33354f18cb57e134()

	__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c] = __obf_5a12388a905e7c2e
	// Calculate new tail position.
	__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c = __obf_56903dfea862bbd5.__obf_3f38a216fb054660(__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c)
	__obf_56903dfea862bbd5.__obf_c29f343498177aa5++
}

// PushFront prepends an element to the front of the queue.
func (__obf_56903dfea862bbd5 *Deque) PushFront(__obf_5a12388a905e7c2e any) {
	__obf_56903dfea862bbd5.__obf_33354f18cb57e134()

	// Calculate new head position.
	__obf_56903dfea862bbd5.__obf_4fb23df1386611b2 = __obf_56903dfea862bbd5.__obf_475810ae089ac9c9(__obf_56903dfea862bbd5.__obf_4fb23df1386611b2)
	__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_4fb23df1386611b2] = __obf_5a12388a905e7c2e
	__obf_56903dfea862bbd5.__obf_c29f343498177aa5++
}

// PopFront removes and returns the element from the front of the queue.
// Implements FIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_56903dfea862bbd5 *Deque) PopFront() any {
	if __obf_56903dfea862bbd5.__obf_c29f343498177aa5 <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_1e3d8fcab4c505b6 := __obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_4fb23df1386611b2]
	__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_4fb23df1386611b2] = nil
	// Calculate new head position.
	__obf_56903dfea862bbd5.__obf_4fb23df1386611b2 = __obf_56903dfea862bbd5.__obf_3f38a216fb054660(__obf_56903dfea862bbd5.__obf_4fb23df1386611b2)
	__obf_56903dfea862bbd5.__obf_c29f343498177aa5--

	__obf_56903dfea862bbd5.__obf_295fbeff55046851()
	return __obf_1e3d8fcab4c505b6
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_56903dfea862bbd5 *Deque) PopBack() any {
	if __obf_56903dfea862bbd5.__obf_c29f343498177aa5 <= 0 {
		panic("deque: PopBack() called on empty queue")
	}

	// Calculate new tail position
	__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c = __obf_56903dfea862bbd5.__obf_475810ae089ac9c9(__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c)

	// Remove value at tail.
	__obf_1e3d8fcab4c505b6 := __obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c]
	__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c] = nil
	__obf_56903dfea862bbd5.__obf_c29f343498177aa5--

	__obf_56903dfea862bbd5.__obf_295fbeff55046851()
	return __obf_1e3d8fcab4c505b6
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_56903dfea862bbd5 *Deque) Front() any {
	if __obf_56903dfea862bbd5.__obf_c29f343498177aa5 <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_4fb23df1386611b2]
}

// Back returns the element at the back of the queue.  This is the element
// that would be returned by PopBack().  This call panics if the queue is
// empty.
func (__obf_56903dfea862bbd5 *Deque) Back() any {
	if __obf_56903dfea862bbd5.__obf_c29f343498177aa5 <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_475810ae089ac9c9(__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c)]
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
func (__obf_56903dfea862bbd5 *Deque) At(__obf_457766c6d5154891 int) any {
	if __obf_457766c6d5154891 < 0 || __obf_457766c6d5154891 >= __obf_56903dfea862bbd5.__obf_c29f343498177aa5 {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[(__obf_56903dfea862bbd5.__obf_4fb23df1386611b2+__obf_457766c6d5154891)&(len(__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908)-1)]
}

// Clear removes all elements from the queue, but retains the current capacity.
// This is useful when repeatedly reusing the queue at high frequency to avoid
// GC during reuse.  The queue will not be resized smaller as long as items are
// only added.  Only when items are removed is the queue subject to getting
// resized smaller.
func (__obf_56903dfea862bbd5 *Deque) Clear() {
	// bitwise modulus
	__obf_59f61d945ae58240 := len(__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908) - 1
	for __obf_79bb9dc2d03a5e29 := __obf_56903dfea862bbd5.__obf_4fb23df1386611b2; __obf_79bb9dc2d03a5e29 != __obf_56903dfea862bbd5.__obf_29adb9bd53a5143c; __obf_79bb9dc2d03a5e29 = (__obf_79bb9dc2d03a5e29 + 1) & __obf_59f61d945ae58240 {
		__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_79bb9dc2d03a5e29] = nil
	}
	__obf_56903dfea862bbd5.__obf_4fb23df1386611b2 = 0
	__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c = 0
	__obf_56903dfea862bbd5.__obf_c29f343498177aa5 = 0
}

// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
// back-to-front.  Having Deque provide Rotate() avoids resizing that could
// happen if implementing rotation using only Pop and Push methods.
func (__obf_56903dfea862bbd5 *Deque) Rotate(__obf_6d56bf10a4deb62c int) {
	if __obf_56903dfea862bbd5.__obf_c29f343498177aa5 <= 1 {
		return
	}
	// Rotating a multiple of q.count is same as no rotation.
	__obf_6d56bf10a4deb62c %= __obf_56903dfea862bbd5.__obf_c29f343498177aa5
	if __obf_6d56bf10a4deb62c == 0 {
		return
	}

	__obf_59f61d945ae58240 := len(__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908) - 1
	// If no empty space in buffer, only move head and tail indexes.
	if __obf_56903dfea862bbd5.__obf_4fb23df1386611b2 == __obf_56903dfea862bbd5.__obf_29adb9bd53a5143c {
		// Calculate new head and tail using bitwise modulus.
		__obf_56903dfea862bbd5.__obf_4fb23df1386611b2 = (__obf_56903dfea862bbd5.__obf_4fb23df1386611b2 + __obf_6d56bf10a4deb62c) & __obf_59f61d945ae58240
		__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c = (__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c + __obf_6d56bf10a4deb62c) & __obf_59f61d945ae58240
		return
	}

	if __obf_6d56bf10a4deb62c < 0 {
		// Rotate back to front.
		for ; __obf_6d56bf10a4deb62c < 0; __obf_6d56bf10a4deb62c++ {
			// Calculate new head and tail using bitwise modulus.
			__obf_56903dfea862bbd5.__obf_4fb23df1386611b2 = (__obf_56903dfea862bbd5.__obf_4fb23df1386611b2 - 1) & __obf_59f61d945ae58240
			__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c = (__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c - 1) & __obf_59f61d945ae58240
			// Put tail value at head and remove value at tail.
			__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_4fb23df1386611b2] = __obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c]
			__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_6d56bf10a4deb62c > 0; __obf_6d56bf10a4deb62c-- {
		// Put head value at tail and remove value at head.
		__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c] = __obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_4fb23df1386611b2]
		__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_4fb23df1386611b2] = nil
		// Calculate new head and tail using bitwise modulus.
		__obf_56903dfea862bbd5.__obf_4fb23df1386611b2 = (__obf_56903dfea862bbd5.__obf_4fb23df1386611b2 + 1) & __obf_59f61d945ae58240
		__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c = (__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c + 1) & __obf_59f61d945ae58240
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_56903dfea862bbd5 *Deque) SetMinCapacity(__obf_c4ac147812953ac1 uint) {
	if 1<<__obf_c4ac147812953ac1 > __obf_f0d6964507650a12 {
		__obf_56903dfea862bbd5.__obf_774dfdb134f2399f = 1 << __obf_c4ac147812953ac1
	} else {
		__obf_56903dfea862bbd5.__obf_774dfdb134f2399f = __obf_f0d6964507650a12
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_56903dfea862bbd5 *Deque) __obf_475810ae089ac9c9(__obf_457766c6d5154891 int) int {
	return (__obf_457766c6d5154891 - 1) & (len(__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908) - 1) // bitwise modulus
}

// next returns the next buffer position wrapping around buffer.
func (__obf_56903dfea862bbd5 *Deque) __obf_3f38a216fb054660(__obf_457766c6d5154891 int) int {
	return (__obf_457766c6d5154891 + 1) & (len(__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908) - 1) // bitwise modulus
}

// growIfFull resizes up if the buffer is full.
func (__obf_56903dfea862bbd5 *Deque) __obf_33354f18cb57e134() {
	if len(__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908) == 0 {
		if __obf_56903dfea862bbd5.__obf_774dfdb134f2399f == 0 {
			__obf_56903dfea862bbd5.__obf_774dfdb134f2399f = __obf_f0d6964507650a12
		}
		__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908 = make([]any, __obf_56903dfea862bbd5.__obf_774dfdb134f2399f)
		return
	}
	if __obf_56903dfea862bbd5.__obf_c29f343498177aa5 == len(__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908) {
		__obf_56903dfea862bbd5.__obf_ca2e59faa4c2b873()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_56903dfea862bbd5 *Deque) __obf_295fbeff55046851() {
	if len(__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908) > __obf_56903dfea862bbd5.__obf_774dfdb134f2399f && (__obf_56903dfea862bbd5.__obf_c29f343498177aa5<<2) == len(__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908) {
		__obf_56903dfea862bbd5.__obf_ca2e59faa4c2b873()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_56903dfea862bbd5 *Deque) __obf_ca2e59faa4c2b873() {
	__obf_64251cbbe9a176c2 := make([]any, __obf_56903dfea862bbd5.__obf_c29f343498177aa5<<1)
	if __obf_56903dfea862bbd5.__obf_29adb9bd53a5143c > __obf_56903dfea862bbd5.__obf_4fb23df1386611b2 {
		copy(__obf_64251cbbe9a176c2, __obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_4fb23df1386611b2:__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c])
	} else {
		__obf_6d56bf10a4deb62c := copy(__obf_64251cbbe9a176c2, __obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[__obf_56903dfea862bbd5.__obf_4fb23df1386611b2:])
		copy(__obf_64251cbbe9a176c2[__obf_6d56bf10a4deb62c:], __obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908[:__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c])
	}

	__obf_56903dfea862bbd5.__obf_4fb23df1386611b2 = 0
	__obf_56903dfea862bbd5.__obf_29adb9bd53a5143c = __obf_56903dfea862bbd5.__obf_c29f343498177aa5
	__obf_56903dfea862bbd5.__obf_9bcb11a51dfb8908 = __obf_64251cbbe9a176c2
}
