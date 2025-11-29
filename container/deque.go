package __obf_b0bebe5eb45b8ad6

// minCapacity is the smallest capacity that deque may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const __obf_b1f1d7b073d05429 = 16

// Deque represents a single instance of the deque data structure.
type Deque struct {
	__obf_ae7e338e9e1ae58c []any
	__obf_60f8632f9380ab0b int
	__obf_b8b64c4ed5238f00 int
	__obf_39f45b145b5eac3d int
	__obf_eae3ac34c1634c15 int
}

// Len returns the number of elements currently stored in the queue.
func (__obf_67fdd157e2264dfc *Deque) Len() int {
	return __obf_67fdd157e2264dfc.

	// PushBack appends an element to the back of the queue.  Implements FIFO when
	// elements are removed with PopFront(), and LIFO when elements are removed
	// with PopBack().
	__obf_39f45b145b5eac3d
}

func (__obf_67fdd157e2264dfc *Deque) PushBack(__obf_326d32b136dc0852 any) {
	__obf_67fdd157e2264dfc.__obf_179d26c9212610c2()
	__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.
	// Calculate new tail position.
	__obf_b8b64c4ed5238f00] = __obf_326d32b136dc0852
	__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00 = __obf_67fdd157e2264dfc.__obf_1cab4cec8958fb57(__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00)
	__obf_67fdd157e2264dfc.

	// PushFront prepends an element to the front of the queue.
	__obf_39f45b145b5eac3d++
}

func (__obf_67fdd157e2264dfc *Deque) PushFront(__obf_326d32b136dc0852 any) {
	__obf_67fdd157e2264dfc.

	// Calculate new head position.
	__obf_179d26c9212610c2()
	__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b = __obf_67fdd157e2264dfc.__obf_1222f1e21057fa76(__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b)
	__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b] = __obf_326d32b136dc0852
	__obf_67fdd157e2264dfc.

	// PopFront removes and returns the element from the front of the queue.
	// Implements FIFO when used with PushBack().  If the queue is empty, the call
	// panics.
	__obf_39f45b145b5eac3d++
}

func (__obf_67fdd157e2264dfc *Deque) PopFront() any {
	if __obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	__obf_8aa26b35d18913ca := __obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b]
	__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.
	// Calculate new head position.
	__obf_60f8632f9380ab0b] = nil
	__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b = __obf_67fdd157e2264dfc.__obf_1cab4cec8958fb57(__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b)
	__obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d--
	__obf_67fdd157e2264dfc.__obf_786c21ae7857ba7b()
	return __obf_8aa26b35d18913ca
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack().  If the queue is empty, the call
// panics.
func (__obf_67fdd157e2264dfc *Deque) PopBack() any {
	if __obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d <= 0 {
		panic("deque: PopBack() called on empty queue")
	}
	__obf_67fdd157e2264dfc.

	// Calculate new tail position
	__obf_b8b64c4ed5238f00 = __obf_67fdd157e2264dfc.

	// Remove value at tail.
	__obf_1222f1e21057fa76(__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00)
	__obf_8aa26b35d18913ca := __obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00]
	__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00] = nil
	__obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d--
	__obf_67fdd157e2264dfc.__obf_786c21ae7857ba7b()
	return __obf_8aa26b35d18913ca
}

// Front returns the element at the front of the queue.  This is the element
// that would be returned by PopFront().  This call panics if the queue is
// empty.
func (__obf_67fdd157e2264dfc *Deque) Front() any {
	if __obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d <= 0 {
		panic("deque: Front() called when empty")
	}
	return __obf_67fdd157e2264dfc.

	// Back returns the element at the back of the queue.  This is the element
	// that would be returned by PopBack().  This call panics if the queue is
	// empty.
	__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b]
}

func (__obf_67fdd157e2264dfc *Deque) Back() any {
	if __obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d <= 0 {
		panic("deque: Back() called when empty")
	}
	return __obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.

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
	__obf_1222f1e21057fa76(__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00)]
}

func (__obf_67fdd157e2264dfc *Deque) At(__obf_095a614fd95fc26e int) any {
	if __obf_095a614fd95fc26e < 0 || __obf_095a614fd95fc26e >= __obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d {
		panic("deque: At() called with index out of range")
	}
	// bitwise modulus
	return __obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[(__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b+__obf_095a614fd95fc26e)&(len(__obf_67fdd157e2264dfc.

	// Clear removes all elements from the queue, but retains the current capacity.
	// This is useful when repeatedly reusing the queue at high frequency to avoid
	// GC during reuse.  The queue will not be resized smaller as long as items are
	// only added.  Only when items are removed is the queue subject to getting
	// resized smaller.
	__obf_ae7e338e9e1ae58c)-1)]
}

func (__obf_67fdd157e2264dfc *Deque) Clear() {
	__obf_32b2c7f5efdd76be :=// bitwise modulus
	len(__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c) - 1
	for __obf_98871087bafa5bd1 := __obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b; __obf_98871087bafa5bd1 != __obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00; __obf_98871087bafa5bd1 = (__obf_98871087bafa5bd1 + 1) & __obf_32b2c7f5efdd76be {
		__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_98871087bafa5bd1] = nil
	}
	__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b = 0
	__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00 = 0
	__obf_67fdd157e2264dfc.

	// Rotate rotates the deque n steps front-to-back.  If n is negative, rotates
	// back-to-front.  Having Deque provide Rotate() avoids resizing that could
	// happen if implementing rotation using only Pop and Push methods.
	__obf_39f45b145b5eac3d = 0
}

func (__obf_67fdd157e2264dfc *Deque) Rotate(__obf_fc5555a986252f47 int) {
	if __obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d <= 1 {
		return
	}
	__obf_fc5555a986252f47 %=// Rotating a multiple of q.count is same as no rotation.
	__obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d
	if __obf_fc5555a986252f47 == 0 {
		return
	}
	__obf_32b2c7f5efdd76be := len(__obf_67fdd157e2264dfc.
	// If no empty space in buffer, only move head and tail indexes.
	__obf_ae7e338e9e1ae58c) - 1

	if __obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b == __obf_67fdd157e2264dfc.
	// Calculate new head and tail using bitwise modulus.
	__obf_b8b64c4ed5238f00 {
		__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b = (__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b + __obf_fc5555a986252f47) & __obf_32b2c7f5efdd76be
		__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00 = (__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00 + __obf_fc5555a986252f47) & __obf_32b2c7f5efdd76be
		return
	}

	if __obf_fc5555a986252f47 < 0 {
		// Rotate back to front.
		for ; __obf_fc5555a986252f47 < 0; __obf_fc5555a986252f47++ {
			__obf_67fdd157e2264dfc.
			// Calculate new head and tail using bitwise modulus.
			__obf_60f8632f9380ab0b = (__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b - 1) & __obf_32b2c7f5efdd76be
			__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00 = (__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00 - 1) & __obf_32b2c7f5efdd76be
			// Put tail value at head and remove value at tail.
			__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b] = __obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00]
			__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00] = nil
		}
		return
	}

	// Rotate front to back.
	for ; __obf_fc5555a986252f47 > 0; __obf_fc5555a986252f47-- {
		__obf_67fdd157e2264dfc.
		// Put head value at tail and remove value at head.
		__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00] = __obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b]
		__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.
		// Calculate new head and tail using bitwise modulus.
		__obf_60f8632f9380ab0b] = nil
		__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b = (__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b + 1) & __obf_32b2c7f5efdd76be
		__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00 = (__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00 + 1) & __obf_32b2c7f5efdd76be
	}
}

// SetMinCapacity sets a minimum capacity of 2^minCapacityExp.  If the value of
// the minimum capacity is less than or equal to the minimum allowed, then
// capacity is set to the minimum allowed.  This may be called at anytime to
// set a new minimum capacity.
//
// Setting a larger minimum capacity may be used to prevent resizing when the
// number of stored items changes frequently across a wide range.
func (__obf_67fdd157e2264dfc *Deque) SetMinCapacity(__obf_418c2249490f8735 uint) {
	if 1<<__obf_418c2249490f8735 > __obf_b1f1d7b073d05429 {
		__obf_67fdd157e2264dfc.__obf_eae3ac34c1634c15 = 1 << __obf_418c2249490f8735
	} else {
		__obf_67fdd157e2264dfc.__obf_eae3ac34c1634c15 = __obf_b1f1d7b073d05429
	}
}

// prev returns the previous buffer position wrapping around buffer.
func (__obf_67fdd157e2264dfc *Deque) __obf_1222f1e21057fa76(__obf_095a614fd95fc26e int) int {
	return (__obf_095a614fd95fc26e - 1) & (len(__obf_67fdd157e2264dfc. // bitwise modulus
	__obf_ae7e338e9e1ae58c) - 1)
}

// next returns the next buffer position wrapping around buffer.
func (__obf_67fdd157e2264dfc *Deque) __obf_1cab4cec8958fb57(__obf_095a614fd95fc26e int) int {
	return (__obf_095a614fd95fc26e + 1) & (len(__obf_67fdd157e2264dfc. // bitwise modulus
	__obf_ae7e338e9e1ae58c) - 1)
}

// growIfFull resizes up if the buffer is full.
func (__obf_67fdd157e2264dfc *Deque) __obf_179d26c9212610c2() {
	if len(__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c) == 0 {
		if __obf_67fdd157e2264dfc.__obf_eae3ac34c1634c15 == 0 {
			__obf_67fdd157e2264dfc.__obf_eae3ac34c1634c15 = __obf_b1f1d7b073d05429
		}
		__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c = make([]any, __obf_67fdd157e2264dfc.__obf_eae3ac34c1634c15)
		return
	}
	if __obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d == len(__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c) {
		__obf_67fdd157e2264dfc.__obf_ee21ea9658bea692()
	}
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (__obf_67fdd157e2264dfc *Deque) __obf_786c21ae7857ba7b() {
	if len(__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c) > __obf_67fdd157e2264dfc.__obf_eae3ac34c1634c15 && (__obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d<<2) == len(__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c) {
		__obf_67fdd157e2264dfc.__obf_ee21ea9658bea692()
	}
}

// resize resizes the deque to fit exactly twice its current contents.  This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (__obf_67fdd157e2264dfc *Deque) __obf_ee21ea9658bea692() {
	__obf_2e0b8165c628f0ff := make([]any, __obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d<<1)
	if __obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00 > __obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b {
		copy(__obf_2e0b8165c628f0ff, __obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b:__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00])
	} else {
		__obf_fc5555a986252f47 := copy(__obf_2e0b8165c628f0ff, __obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b:])
		copy(__obf_2e0b8165c628f0ff[__obf_fc5555a986252f47:], __obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c[:__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00])
	}
	__obf_67fdd157e2264dfc.__obf_60f8632f9380ab0b = 0
	__obf_67fdd157e2264dfc.__obf_b8b64c4ed5238f00 = __obf_67fdd157e2264dfc.__obf_39f45b145b5eac3d
	__obf_67fdd157e2264dfc.__obf_ae7e338e9e1ae58c = __obf_2e0b8165c628f0ff
}
