package __obf_c3cd04a15c56f32f

import (
	"io"
)

// stream is a io.Writer like object, with JSON specific write functions.
// Error is not returned as return value, but stored as Error member on this stream instance.
type Stream struct {
	__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb
	__obf_743aba00c5d1ef26 io.Writer
	__obf_ace979f6622823f3 []byte
	Error                  error
	__obf_7ccde3f02c81cd81 int
	Attachment             any // open for customized encoder
}

// NewStream create new stream instance.
// cfg can be jsoniter.ConfigDefault.
// out can be nil if write to internal buffer.
// bufSize is the initial size for the internal buffer in bytes.
func NewStream(__obf_0c8065c219ccf0ab API, __obf_743aba00c5d1ef26 io.Writer, __obf_5ab0791f20d1d4db int) *Stream {
	return &Stream{__obf_0c8065c219ccf0ab: __obf_0c8065c219ccf0ab.(*__obf_3a25fb4c9a8342bb), __obf_743aba00c5d1ef26: __obf_743aba00c5d1ef26, __obf_ace979f6622823f3: make([]byte, 0, __obf_5ab0791f20d1d4db), Error: nil, __obf_7ccde3f02c81cd81: 0}
}

// Pool returns a pool can provide more stream with same configuration
func (__obf_2361f5214e84c60f *Stream) Pool() StreamPool {
	return __obf_2361f5214e84c60f.

		// Reset reuse this stream instance by assign a new writer
		__obf_0c8065c219ccf0ab
}

func (__obf_2361f5214e84c60f *Stream) Reset(__obf_743aba00c5d1ef26 io.Writer) {
	__obf_2361f5214e84c60f.__obf_743aba00c5d1ef26 = __obf_743aba00c5d1ef26
	__obf_2361f5214e84c60f.

		// Available returns how many bytes are unused in the buffer.
		__obf_ace979f6622823f3 = __obf_2361f5214e84c60f.__obf_ace979f6622823f3[:0]
}

func (__obf_2361f5214e84c60f *Stream) Available() int {
	return cap(__obf_2361f5214e84c60f.__obf_ace979f6622823f3) - len(__obf_2361f5214e84c60f.

		// Buffered returns the number of bytes that have been written into the current buffer.
		__obf_ace979f6622823f3)
}

func (__obf_2361f5214e84c60f *Stream) Buffered() int {
	return len(__obf_2361f5214e84c60f.

		// Buffer if writer is nil, use this method to take the result
		__obf_ace979f6622823f3)
}

func (__obf_2361f5214e84c60f *Stream) Buffer() []byte {
	return __obf_2361f5214e84c60f.

		// SetBuffer allows to append to the internal buffer directly
		__obf_ace979f6622823f3
}

func (__obf_2361f5214e84c60f *Stream) SetBuffer(__obf_ace979f6622823f3 []byte) {
	__obf_2361f5214e84c60f.

		// Write writes the contents of p into the buffer.
		// It returns the number of bytes written.
		// If nn < len(p), it also returns an error explaining
		// why the write is short.
		__obf_ace979f6622823f3 = __obf_ace979f6622823f3
}

func (__obf_2361f5214e84c60f *Stream) Write(__obf_420b1bf30275e403 []byte) (__obf_4547628c501616ee int, __obf_5966ecc5edb9b17e error) {
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_420b1bf30275e403...)
	if __obf_2361f5214e84c60f.__obf_743aba00c5d1ef26 != nil {
		__obf_4547628c501616ee, __obf_5966ecc5edb9b17e = __obf_2361f5214e84c60f.__obf_743aba00c5d1ef26.Write(__obf_2361f5214e84c60f.__obf_ace979f6622823f3)
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_2361f5214e84c60f.__obf_ace979f6622823f3[__obf_4547628c501616ee:]
		return
	}
	return len(__obf_420b1bf30275e403), nil
}

// WriteByte writes a single byte.
func (__obf_2361f5214e84c60f *Stream) __obf_c4fec0edfb3875ad(__obf_0c1bc1e511a43120 byte) {
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_0c1bc1e511a43120)
}

func (__obf_2361f5214e84c60f *Stream) __obf_5e728551f00598e5(__obf_b8cc99c90d8b529f byte, __obf_56db9d416fbe7586 byte) {
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_b8cc99c90d8b529f, __obf_56db9d416fbe7586)
}

func (__obf_2361f5214e84c60f *Stream) __obf_d4910575fd3711ca(__obf_b8cc99c90d8b529f byte, __obf_56db9d416fbe7586 byte, __obf_c03d116d77d6e4b5 byte) {
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_b8cc99c90d8b529f, __obf_56db9d416fbe7586, __obf_c03d116d77d6e4b5)
}

func (__obf_2361f5214e84c60f *Stream) __obf_06499ccc8898b24b(__obf_b8cc99c90d8b529f byte, __obf_56db9d416fbe7586 byte, __obf_c03d116d77d6e4b5 byte, __obf_0fa25a84f0fccf03 byte) {
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_b8cc99c90d8b529f, __obf_56db9d416fbe7586, __obf_c03d116d77d6e4b5, __obf_0fa25a84f0fccf03)
}

func (__obf_2361f5214e84c60f *Stream) __obf_54c3f9e691e28ed6(__obf_b8cc99c90d8b529f byte, __obf_56db9d416fbe7586 byte, __obf_c03d116d77d6e4b5 byte, __obf_0fa25a84f0fccf03 byte, __obf_e5f212f092815fc0 byte) {
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3,

		// Flush writes any buffered data to the underlying io.Writer.
		__obf_b8cc99c90d8b529f, __obf_56db9d416fbe7586, __obf_c03d116d77d6e4b5, __obf_0fa25a84f0fccf03, __obf_e5f212f092815fc0)
}

func (__obf_2361f5214e84c60f *Stream) Flush() error {
	if __obf_2361f5214e84c60f.__obf_743aba00c5d1ef26 == nil {
		return nil
	}
	if __obf_2361f5214e84c60f.Error != nil {
		return __obf_2361f5214e84c60f.Error
	}
	_, __obf_5966ecc5edb9b17e := __obf_2361f5214e84c60f.__obf_743aba00c5d1ef26.Write(__obf_2361f5214e84c60f.__obf_ace979f6622823f3)
	if __obf_5966ecc5edb9b17e != nil {
		if __obf_2361f5214e84c60f.Error == nil {
			__obf_2361f5214e84c60f.
				Error = __obf_5966ecc5edb9b17e
		}
		return __obf_5966ecc5edb9b17e
	}
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = __obf_2361f5214e84c60f.__obf_ace979f6622823f3[:0]
	return nil
}

// WriteRaw write string out without quotes, just like []byte
func (__obf_2361f5214e84c60f *Stream) WriteRaw(__obf_20e65aaba6bfc813 string) {
	__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.

		// WriteNil write null to stream
		__obf_ace979f6622823f3, __obf_20e65aaba6bfc813...)
}

func (__obf_2361f5214e84c60f *Stream) WriteNil() {
	__obf_2361f5214e84c60f.__obf_06499ccc8898b24b('n', 'u', 'l', 'l')
}

// WriteTrue write true to stream
func (__obf_2361f5214e84c60f *Stream) WriteTrue() {
	__obf_2361f5214e84c60f.__obf_06499ccc8898b24b('t', 'r', 'u', 'e')
}

// WriteFalse write false to stream
func (__obf_2361f5214e84c60f *Stream) WriteFalse() {
	__obf_2361f5214e84c60f.__obf_54c3f9e691e28ed6('f', 'a', 'l', 's', 'e')
}

// WriteBool write true or false into stream
func (__obf_2361f5214e84c60f *Stream) WriteBool(__obf_d59813f23ad740a8 bool) {
	if __obf_d59813f23ad740a8 {
		__obf_2361f5214e84c60f.
			WriteTrue()
	} else {
		__obf_2361f5214e84c60f.
			WriteFalse()
	}
}

// WriteObjectStart write { with possible indention
func (__obf_2361f5214e84c60f *Stream) WriteObjectStart() {
	__obf_2361f5214e84c60f.__obf_7ccde3f02c81cd81 += __obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.__obf_1ee92c16ad6ec329
	__obf_2361f5214e84c60f.

		// WriteObjectField write "field": with possible indention
		__obf_c4fec0edfb3875ad('{')
	__obf_2361f5214e84c60f.__obf_01752a60ae0b165d(0)
}

func (__obf_2361f5214e84c60f *Stream) WriteObjectField(__obf_f992c91036745ca0 string) {
	__obf_2361f5214e84c60f.
		WriteString(__obf_f992c91036745ca0)
	if __obf_2361f5214e84c60f.__obf_7ccde3f02c81cd81 > 0 {
		__obf_2361f5214e84c60f.__obf_5e728551f00598e5(':', ' ')
	} else {
		__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad(':')
	}
}

// WriteObjectEnd write } with possible indention
func (__obf_2361f5214e84c60f *Stream) WriteObjectEnd() {
	__obf_2361f5214e84c60f.__obf_01752a60ae0b165d(__obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.__obf_1ee92c16ad6ec329)
	__obf_2361f5214e84c60f.__obf_7ccde3f02c81cd81 -= __obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.__obf_1ee92c16ad6ec329

	// WriteEmptyObject write {}
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('}')
}

func (__obf_2361f5214e84c60f *Stream) WriteEmptyObject() {
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('{')
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('}')
}

// WriteMore write , with possible indention
func (__obf_2361f5214e84c60f *Stream) WriteMore() {
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad(',')
	__obf_2361f5214e84c60f.__obf_01752a60ae0b165d(0)
}

// WriteArrayStart write [ with possible indention
func (__obf_2361f5214e84c60f *Stream) WriteArrayStart() {
	__obf_2361f5214e84c60f.__obf_7ccde3f02c81cd81 += __obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.__obf_1ee92c16ad6ec329
	__obf_2361f5214e84c60f.

		// WriteEmptyArray write []
		__obf_c4fec0edfb3875ad('[')
	__obf_2361f5214e84c60f.__obf_01752a60ae0b165d(0)
}

func (__obf_2361f5214e84c60f *Stream) WriteEmptyArray() {
	__obf_2361f5214e84c60f.__obf_5e728551f00598e5('[', ']')
}

// WriteArrayEnd write ] with possible indention
func (__obf_2361f5214e84c60f *Stream) WriteArrayEnd() {
	__obf_2361f5214e84c60f.__obf_01752a60ae0b165d(__obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.__obf_1ee92c16ad6ec329)
	__obf_2361f5214e84c60f.__obf_7ccde3f02c81cd81 -= __obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.__obf_1ee92c16ad6ec329
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad(']')
}

func (__obf_2361f5214e84c60f *Stream) __obf_01752a60ae0b165d(__obf_b05997550f8ad591 int) {
	if __obf_2361f5214e84c60f.__obf_7ccde3f02c81cd81 == 0 {
		return
	}
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('\n')
	__obf_3973c7cdf4a4e02b := __obf_2361f5214e84c60f.__obf_7ccde3f02c81cd81 - __obf_b05997550f8ad591
	for __obf_28d099df85f083a8 := 0; __obf_28d099df85f083a8 < __obf_3973c7cdf4a4e02b; __obf_28d099df85f083a8++ {
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, ' ')
	}
}
