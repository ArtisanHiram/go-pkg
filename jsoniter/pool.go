package __obf_c7b79b12b33d8238

import (
	"io"
)

// IteratorPool a thread safe pool of iterators with same configuration
type IteratorPool interface {
	BorrowIterator(__obf_1d34d01a9b83218a []byte) *Iterator
	ReturnIterator(__obf_0da8c843dad13139 *Iterator)
}

// StreamPool a thread safe pool of streams with same configuration
type StreamPool interface {
	BorrowStream(__obf_29b48e4cc46aebd6 io.Writer) *Stream
	ReturnStream(__obf_d8f50bcfe87d8b47 *Stream)
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) BorrowStream(__obf_29b48e4cc46aebd6 io.Writer) *Stream {
	__obf_d8f50bcfe87d8b47 := __obf_c52e0bcfad4b8b71.__obf_4409df5a50bd7d1b.Get().(*Stream)
	__obf_d8f50bcfe87d8b47.
		Reset(__obf_29b48e4cc46aebd6)
	return __obf_d8f50bcfe87d8b47
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) ReturnStream(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.__obf_8b93306c8fe0d999 = nil
	__obf_d8f50bcfe87d8b47.
		Error = nil
	__obf_d8f50bcfe87d8b47.
		Attachment = nil
	__obf_c52e0bcfad4b8b71.__obf_4409df5a50bd7d1b.
		Put(__obf_d8f50bcfe87d8b47)
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) BorrowIterator(__obf_1d34d01a9b83218a []byte) *Iterator {
	__obf_0da8c843dad13139 := __obf_c52e0bcfad4b8b71.__obf_4d02f79ac1d4d36b.Get().(*Iterator)
	__obf_0da8c843dad13139.
		ResetBytes(__obf_1d34d01a9b83218a)
	return __obf_0da8c843dad13139
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) ReturnIterator(__obf_0da8c843dad13139 *Iterator) {
	__obf_0da8c843dad13139.
		Error = nil
	__obf_0da8c843dad13139.
		Attachment = nil
	__obf_c52e0bcfad4b8b71.__obf_4d02f79ac1d4d36b.
		Put(__obf_0da8c843dad13139)
}
