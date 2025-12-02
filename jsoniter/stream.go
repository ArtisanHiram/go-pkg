package __obf_c7b79b12b33d8238

import (
	"io"
)

// stream is a io.Writer like object, with JSON specific write functions.
// Error is not returned as return value, but stored as Error member on this stream instance.
type Stream struct {
	__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0
	__obf_8b93306c8fe0d999 io.Writer
	__obf_684d738bc3ac851a []byte
	Error                  error
	__obf_4f8e10a1467d205b int
	Attachment             any // open for customized encoder
}

// NewStream create new stream instance.
// cfg can be jsoniter.ConfigDefault.
// out can be nil if write to internal buffer.
// bufSize is the initial size for the internal buffer in bytes.
func NewStream(__obf_c52e0bcfad4b8b71 API, __obf_8b93306c8fe0d999 io.Writer, __obf_90520bf89c30ab5c int) *Stream {
	return &Stream{__obf_c52e0bcfad4b8b71: __obf_c52e0bcfad4b8b71.(*__obf_30fe5c95cabd69c0), __obf_8b93306c8fe0d999: __obf_8b93306c8fe0d999, __obf_684d738bc3ac851a: make([]byte, 0, __obf_90520bf89c30ab5c), Error: nil, __obf_4f8e10a1467d205b: 0}
}

// Pool returns a pool can provide more stream with same configuration
func (__obf_d8f50bcfe87d8b47 *Stream) Pool() StreamPool {
	return __obf_d8f50bcfe87d8b47.

		// Reset reuse this stream instance by assign a new writer
		__obf_c52e0bcfad4b8b71
}

func (__obf_d8f50bcfe87d8b47 *Stream) Reset(__obf_8b93306c8fe0d999 io.Writer) {
	__obf_d8f50bcfe87d8b47.__obf_8b93306c8fe0d999 = __obf_8b93306c8fe0d999
	__obf_d8f50bcfe87d8b47.

		// Available returns how many bytes are unused in the buffer.
		__obf_684d738bc3ac851a = __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[:0]
}

func (__obf_d8f50bcfe87d8b47 *Stream) Available() int {
	return cap(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a) - len(__obf_d8f50bcfe87d8b47.

		// Buffered returns the number of bytes that have been written into the current buffer.
		__obf_684d738bc3ac851a)
}

func (__obf_d8f50bcfe87d8b47 *Stream) Buffered() int {
	return len(__obf_d8f50bcfe87d8b47.

		// Buffer if writer is nil, use this method to take the result
		__obf_684d738bc3ac851a)
}

func (__obf_d8f50bcfe87d8b47 *Stream) Buffer() []byte {
	return __obf_d8f50bcfe87d8b47.

		// SetBuffer allows to append to the internal buffer directly
		__obf_684d738bc3ac851a
}

func (__obf_d8f50bcfe87d8b47 *Stream) SetBuffer(__obf_684d738bc3ac851a []byte) {
	__obf_d8f50bcfe87d8b47.

		// Write writes the contents of p into the buffer.
		// It returns the number of bytes written.
		// If nn < len(p), it also returns an error explaining
		// why the write is short.
		__obf_684d738bc3ac851a = __obf_684d738bc3ac851a
}

func (__obf_d8f50bcfe87d8b47 *Stream) Write(__obf_c03fbc322c1a9109 []byte) (__obf_9d996b7bbb5a7fa7 int, __obf_ea0680f8b461a85b error) {
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_c03fbc322c1a9109...)
	if __obf_d8f50bcfe87d8b47.__obf_8b93306c8fe0d999 != nil {
		__obf_9d996b7bbb5a7fa7, __obf_ea0680f8b461a85b = __obf_d8f50bcfe87d8b47.__obf_8b93306c8fe0d999.Write(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a)
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[__obf_9d996b7bbb5a7fa7:]
		return
	}
	return len(__obf_c03fbc322c1a9109), nil
}

// WriteByte writes a single byte.
func (__obf_d8f50bcfe87d8b47 *Stream) __obf_7340077d41996822(__obf_12577c03a12f0d2c byte) {
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_12577c03a12f0d2c)
}

func (__obf_d8f50bcfe87d8b47 *Stream) __obf_99147f1a9b9d5ffc(__obf_81b1c9a9db5bdfab byte, __obf_9e81083082bf048f byte) {
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_81b1c9a9db5bdfab, __obf_9e81083082bf048f)
}

func (__obf_d8f50bcfe87d8b47 *Stream) __obf_656069488f52bbbb(__obf_81b1c9a9db5bdfab byte, __obf_9e81083082bf048f byte, __obf_0281717e51d6b416 byte) {
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_81b1c9a9db5bdfab, __obf_9e81083082bf048f, __obf_0281717e51d6b416)
}

func (__obf_d8f50bcfe87d8b47 *Stream) __obf_d4ac6556d988c1e8(__obf_81b1c9a9db5bdfab byte, __obf_9e81083082bf048f byte, __obf_0281717e51d6b416 byte, __obf_29c9b2e02c64ffae byte) {
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_81b1c9a9db5bdfab, __obf_9e81083082bf048f, __obf_0281717e51d6b416, __obf_29c9b2e02c64ffae)
}

func (__obf_d8f50bcfe87d8b47 *Stream) __obf_7dbda4e6e4979498(__obf_81b1c9a9db5bdfab byte, __obf_9e81083082bf048f byte, __obf_0281717e51d6b416 byte, __obf_29c9b2e02c64ffae byte, __obf_a012b487b925fe50 byte) {
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a,

		// Flush writes any buffered data to the underlying io.Writer.
		__obf_81b1c9a9db5bdfab, __obf_9e81083082bf048f, __obf_0281717e51d6b416, __obf_29c9b2e02c64ffae, __obf_a012b487b925fe50)
}

func (__obf_d8f50bcfe87d8b47 *Stream) Flush() error {
	if __obf_d8f50bcfe87d8b47.__obf_8b93306c8fe0d999 == nil {
		return nil
	}
	if __obf_d8f50bcfe87d8b47.Error != nil {
		return __obf_d8f50bcfe87d8b47.Error
	}
	_, __obf_ea0680f8b461a85b := __obf_d8f50bcfe87d8b47.__obf_8b93306c8fe0d999.Write(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a)
	if __obf_ea0680f8b461a85b != nil {
		if __obf_d8f50bcfe87d8b47.Error == nil {
			__obf_d8f50bcfe87d8b47.
				Error = __obf_ea0680f8b461a85b
		}
		return __obf_ea0680f8b461a85b
	}
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = __obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a[:0]
	return nil
}

// WriteRaw write string out without quotes, just like []byte
func (__obf_d8f50bcfe87d8b47 *Stream) WriteRaw(__obf_057e9641d10a1246 string) {
	__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.

		// WriteNil write null to stream
		__obf_684d738bc3ac851a, __obf_057e9641d10a1246...)
}

func (__obf_d8f50bcfe87d8b47 *Stream) WriteNil() {
	__obf_d8f50bcfe87d8b47.__obf_d4ac6556d988c1e8('n', 'u', 'l', 'l')
}

// WriteTrue write true to stream
func (__obf_d8f50bcfe87d8b47 *Stream) WriteTrue() {
	__obf_d8f50bcfe87d8b47.__obf_d4ac6556d988c1e8('t', 'r', 'u', 'e')
}

// WriteFalse write false to stream
func (__obf_d8f50bcfe87d8b47 *Stream) WriteFalse() {
	__obf_d8f50bcfe87d8b47.__obf_7dbda4e6e4979498('f', 'a', 'l', 's', 'e')
}

// WriteBool write true or false into stream
func (__obf_d8f50bcfe87d8b47 *Stream) WriteBool(__obf_35accd096612ebe4 bool) {
	if __obf_35accd096612ebe4 {
		__obf_d8f50bcfe87d8b47.
			WriteTrue()
	} else {
		__obf_d8f50bcfe87d8b47.
			WriteFalse()
	}
}

// WriteObjectStart write { with possible indention
func (__obf_d8f50bcfe87d8b47 *Stream) WriteObjectStart() {
	__obf_d8f50bcfe87d8b47.__obf_4f8e10a1467d205b += __obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.__obf_9f90ab8046926590
	__obf_d8f50bcfe87d8b47.

		// WriteObjectField write "field": with possible indention
		__obf_7340077d41996822('{')
	__obf_d8f50bcfe87d8b47.__obf_545fe7d3561de810(0)
}

func (__obf_d8f50bcfe87d8b47 *Stream) WriteObjectField(__obf_ad34f8a6de2b01cb string) {
	__obf_d8f50bcfe87d8b47.
		WriteString(__obf_ad34f8a6de2b01cb)
	if __obf_d8f50bcfe87d8b47.__obf_4f8e10a1467d205b > 0 {
		__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc(':', ' ')
	} else {
		__obf_d8f50bcfe87d8b47.__obf_7340077d41996822(':')
	}
}

// WriteObjectEnd write } with possible indention
func (__obf_d8f50bcfe87d8b47 *Stream) WriteObjectEnd() {
	__obf_d8f50bcfe87d8b47.__obf_545fe7d3561de810(__obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.__obf_9f90ab8046926590)
	__obf_d8f50bcfe87d8b47.__obf_4f8e10a1467d205b -= __obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.__obf_9f90ab8046926590

	// WriteEmptyObject write {}
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('}')
}

func (__obf_d8f50bcfe87d8b47 *Stream) WriteEmptyObject() {
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('{')
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('}')
}

// WriteMore write , with possible indention
func (__obf_d8f50bcfe87d8b47 *Stream) WriteMore() {
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822(',')
	__obf_d8f50bcfe87d8b47.__obf_545fe7d3561de810(0)
}

// WriteArrayStart write [ with possible indention
func (__obf_d8f50bcfe87d8b47 *Stream) WriteArrayStart() {
	__obf_d8f50bcfe87d8b47.__obf_4f8e10a1467d205b += __obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.__obf_9f90ab8046926590
	__obf_d8f50bcfe87d8b47.

		// WriteEmptyArray write []
		__obf_7340077d41996822('[')
	__obf_d8f50bcfe87d8b47.__obf_545fe7d3561de810(0)
}

func (__obf_d8f50bcfe87d8b47 *Stream) WriteEmptyArray() {
	__obf_d8f50bcfe87d8b47.__obf_99147f1a9b9d5ffc('[', ']')
}

// WriteArrayEnd write ] with possible indention
func (__obf_d8f50bcfe87d8b47 *Stream) WriteArrayEnd() {
	__obf_d8f50bcfe87d8b47.__obf_545fe7d3561de810(__obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.__obf_9f90ab8046926590)
	__obf_d8f50bcfe87d8b47.__obf_4f8e10a1467d205b -= __obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.__obf_9f90ab8046926590
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822(']')
}

func (__obf_d8f50bcfe87d8b47 *Stream) __obf_545fe7d3561de810(__obf_8a954554a1a653c5 int) {
	if __obf_d8f50bcfe87d8b47.__obf_4f8e10a1467d205b == 0 {
		return
	}
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('\n')
	__obf_133bd33d89f669ce := __obf_d8f50bcfe87d8b47.__obf_4f8e10a1467d205b - __obf_8a954554a1a653c5
	for __obf_a051269af3a541bb := 0; __obf_a051269af3a541bb < __obf_133bd33d89f669ce; __obf_a051269af3a541bb++ {
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, ' ')
	}
}
