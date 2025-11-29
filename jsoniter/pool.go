package __obf_91620b895eeff9ed

import (
	"io"
)

// IteratorPool a thread safe pool of iterators with same configuration
type IteratorPool interface {
	BorrowIterator(__obf_2894589095c323b4 []byte) *Iterator
	ReturnIterator(__obf_1bb30e8a74ed8233 *Iterator)
}

// StreamPool a thread safe pool of streams with same configuration
type StreamPool interface {
	BorrowStream(__obf_e4132b2c3ef9e14e io.Writer) *Stream
	ReturnStream(__obf_850a7457bb739a32 *Stream)
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) BorrowStream(__obf_e4132b2c3ef9e14e io.Writer) *Stream {
	__obf_850a7457bb739a32 := __obf_892632c77e859037.__obf_9e7b77251a4cf7c4.Get().(*Stream)
	__obf_850a7457bb739a32.
		Reset(__obf_e4132b2c3ef9e14e)
	return __obf_850a7457bb739a32
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) ReturnStream(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.__obf_feb7fc81d385b509 = nil
	__obf_850a7457bb739a32.
		Error = nil
	__obf_850a7457bb739a32.
		Attachment = nil
	__obf_892632c77e859037.__obf_9e7b77251a4cf7c4.
		Put(__obf_850a7457bb739a32)
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) BorrowIterator(__obf_2894589095c323b4 []byte) *Iterator {
	__obf_1bb30e8a74ed8233 := __obf_892632c77e859037.__obf_145970d48796be8b.Get().(*Iterator)
	__obf_1bb30e8a74ed8233.
		ResetBytes(__obf_2894589095c323b4)
	return __obf_1bb30e8a74ed8233
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) ReturnIterator(__obf_1bb30e8a74ed8233 *Iterator) {
	__obf_1bb30e8a74ed8233.
		Error = nil
	__obf_1bb30e8a74ed8233.
		Attachment = nil
	__obf_892632c77e859037.__obf_145970d48796be8b.
		Put(__obf_1bb30e8a74ed8233)
}
