package __obf_4f13ac5aa043b5ce

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_5a8afb345a76ad34 = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_e6532267fbea040a []any
	__obf_abd8b6f41210aeeb int
	__obf_3fc6be8a6a6e1dcc int
	__obf_7a971de362a6bf24 int
	__obf_d6362f25e2fab14a int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_fa2b5b41e5a12cb2 *Deque) Len() int {
	return __obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24
}

// PushBack appends an element to the back of the queue.  Implements FIFO when
// elements are removed with PopFront(), and LIFO when elements are removed
// with PopBack().
func (__obf_fa2b5b41e5a12cb2 *Deque) PushBack(__obf_abcd0da910a815c1 any) {
	__obf_fa2b5b41e5a12cb2.__obf_7ad1f25f4dc337cf()

	__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc] = __obf_abcd0da910a815c1
	// Calculate new tail position.
	__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc = __obf_fa2b5b41e5a12cb2.__obf_ab34c2f73a61331f(__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc)
	__obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24++
}

// PushFront prepends an element to the front of the queue.
func (__obf_fa2b5b41e5a12cb2 *Deque) PushFront(__obf_abcd0da910a815c1 any) {
	__obf_fa2b5b41e5a12cb2.__obf_7ad1f25f4dc337cf()

	// Calculate new head position.
	__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb = __obf_fa2b5b41e5a12cb2.__obf_5c809ff977b880f2(__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb)
	__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb] = __obf_abcd0da910a815c1
	__obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24++
}

// PopFront removes and returns the element from the front of the queue.
// Implements FIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_fa2b5b41e5a12cb2 *Deque) PopFront() any {
	if __obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24 <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_752ac253da10725b := __obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb]
	__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb] = nil
	// Calculate new head position.
	__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb = __obf_fa2b5b41e5a12cb2.__obf_ab34c2f73a61331f(__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb)
	__obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24--

	__obf_fa2b5b41e5a12cb2.__obf_a8575288f80b3db7()
	return __obf_752ac253da10725b
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_fa2b5b41e5a12cb2 *Deque) PopBack() any {
	if __obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24 <= 0 {
		panic("deque: PopBack() called on empty queue")
	}

	// Calculate new tail position
	__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc = __obf_fa2b5b41e5a12cb2.__obf_5c809ff977b880f2(__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc)

	// Remove value at tail.
	__obf_752ac253da10725b := __obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc]
	__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc] = nil
	__obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24--

	__obf_fa2b5b41e5a12cb2.__obf_a8575288f80b3db7()
	return __obf_752ac253da10725b
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_fa2b5b41e5a12cb2 *Deque) Front() any {
	if __obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24 <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb]
}

// Back returns the element at the back of the queue.  This is the element
// that would be returned by PopBack().  This call panics if the queue is
// empty.
func (__obf_fa2b5b41e5a12cb2 *Deque) Back() any {
	if __obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24 <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_5c809ff977b880f2(__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc)]
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
func (__obf_fa2b5b41e5a12cb2 *Deque) At(__obf_e3e988c1360a57a4 int) any {
	if __obf_e3e988c1360a57a4 < 0 || __obf_e3e988c1360a57a4 >= __obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24 {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[(__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb+__obf_e3e988c1360a57a4)&(len(__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a)-1)]
}

// Clear removes all elements from the queue, but retains the current capacity.
// This is useful when repeatedly reusing the queue at high frequency to avoid
// GC during reuse.  The queue will not be resized smaller as long as items are
// only added.  Only when items are removed is the queue subject to getting
// resized smaller.
func (__obf_fa2b5b41e5a12cb2 *Deque) Clear() {
	// bitwise modulus
	__obf_6f88a05709b2919b := len(__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a) - 1
	for __obf_b1fecdf21015160d := __obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb; __obf_b1fecdf21015160d != __obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc; __obf_b1fecdf21015160d = (__obf_b1fecdf21015160d + 1) & __obf_6f88a05709b2919b {
		__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_b1fecdf21015160d] = nil
	}
	__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb = 0
	__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc = 0
	__obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24 = 0
}

// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
// back-to-front.  Having Deque provide Rotate() avoids resizing that could
// happen if implementing rotation using only Pop and Push methods.
func (__obf_fa2b5b41e5a12cb2 *Deque) Rotate(__obf_9178770692fd13ed int) {
	if __obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24 <= 1 {
		return
	}
	// Rotating a multiple of q.count is same as no rotation.
	__obf_9178770692fd13ed %= __obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24
	if __obf_9178770692fd13ed == 0 {
		return
	}

	__obf_6f88a05709b2919b := len(__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a) - 1
	// If no empty space in buffer, only move head and tail indexes.
	if __obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb == __obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc {
		// Calculate new head and tail using bitwise modulus.
		__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb = (__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb + __obf_9178770692fd13ed) & __obf_6f88a05709b2919b
		__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc = (__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc + __obf_9178770692fd13ed) & __obf_6f88a05709b2919b
		return
	}

	if __obf_9178770692fd13ed < 0 {
		// Rotate back to front.
		for ; __obf_9178770692fd13ed < 0; __obf_9178770692fd13ed++ {
			// Calculate new head and tail using bitwise modulus.
			__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb = (__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb - 1) & __obf_6f88a05709b2919b
			__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc = (__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc - 1) & __obf_6f88a05709b2919b
			// Put tail value at head and remove value at tail.
			__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb] = __obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc]
			__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_9178770692fd13ed > 0; __obf_9178770692fd13ed-- {
		// Put head value at tail and remove value at head.
		__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc] = __obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb]
		__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb] = nil
		// Calculate new head and tail using bitwise modulus.
		__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb = (__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb + 1) & __obf_6f88a05709b2919b
		__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc = (__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc + 1) & __obf_6f88a05709b2919b
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_fa2b5b41e5a12cb2 *Deque) SetMinCapacity(__obf_086d71c831a1efe1 uint) {
	if 1<<__obf_086d71c831a1efe1 > __obf_5a8afb345a76ad34 {
		__obf_fa2b5b41e5a12cb2.__obf_d6362f25e2fab14a = 1 << __obf_086d71c831a1efe1
	} else {
		__obf_fa2b5b41e5a12cb2.__obf_d6362f25e2fab14a = __obf_5a8afb345a76ad34
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_fa2b5b41e5a12cb2 *Deque) __obf_5c809ff977b880f2(__obf_e3e988c1360a57a4 int) int {
	return (__obf_e3e988c1360a57a4 - 1) & (len(__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a) - 1) // bitwise modulus
}

// next returns the next buffer position wrapping around buffer.
func (__obf_fa2b5b41e5a12cb2 *Deque) __obf_ab34c2f73a61331f(__obf_e3e988c1360a57a4 int) int {
	return (__obf_e3e988c1360a57a4 + 1) & (len(__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a) - 1) // bitwise modulus
}

// growIfFull resizes up if the buffer is full.
func (__obf_fa2b5b41e5a12cb2 *Deque) __obf_7ad1f25f4dc337cf() {
	if len(__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a) == 0 {
		if __obf_fa2b5b41e5a12cb2.__obf_d6362f25e2fab14a == 0 {
			__obf_fa2b5b41e5a12cb2.__obf_d6362f25e2fab14a = __obf_5a8afb345a76ad34
		}
		__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a = make([]any, __obf_fa2b5b41e5a12cb2.__obf_d6362f25e2fab14a)
		return
	}
	if __obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24 == len(__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a) {
		__obf_fa2b5b41e5a12cb2.__obf_d698abc7d8370538()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_fa2b5b41e5a12cb2 *Deque) __obf_a8575288f80b3db7() {
	if len(__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a) > __obf_fa2b5b41e5a12cb2.__obf_d6362f25e2fab14a && (__obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24<<2) == len(__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a) {
		__obf_fa2b5b41e5a12cb2.__obf_d698abc7d8370538()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_fa2b5b41e5a12cb2 *Deque) __obf_d698abc7d8370538() {
	__obf_1bf0865c48b8fcdf := make([]any, __obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24<<1)
	if __obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc > __obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb {
		copy(__obf_1bf0865c48b8fcdf, __obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb:__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc])
	} else {
		__obf_9178770692fd13ed := copy(__obf_1bf0865c48b8fcdf, __obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb:])
		copy(__obf_1bf0865c48b8fcdf[__obf_9178770692fd13ed:], __obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a[:__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc])
	}

	__obf_fa2b5b41e5a12cb2.__obf_abd8b6f41210aeeb = 0
	__obf_fa2b5b41e5a12cb2.__obf_3fc6be8a6a6e1dcc = __obf_fa2b5b41e5a12cb2.__obf_7a971de362a6bf24
	__obf_fa2b5b41e5a12cb2.__obf_e6532267fbea040a = __obf_1bf0865c48b8fcdf
}
