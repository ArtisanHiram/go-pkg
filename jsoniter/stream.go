package __obf_91620b895eeff9ed

import (
	"io"
)

// stream is a io.Writer like object, with JSON specific write functions.
// Error is not returned as return value, but stored as Error member on this stream instance.
type Stream struct {
	__obf_892632c77e859037 *__obf_972c2293804d141c
	__obf_feb7fc81d385b509 io.Writer
	__obf_184433571fa55237 []byte
	Error                  error
	__obf_2dfa32334fb1aa21 int
	Attachment             any // open for customized encoder
}

// NewStream create new stream instance.
// cfg can be jsoniter.ConfigDefault.
// out can be nil if write to internal buffer.
// bufSize is the initial size for the internal buffer in bytes.
func NewStream(__obf_892632c77e859037 API, __obf_feb7fc81d385b509 io.Writer, __obf_0355f91790eba345 int) *Stream {
	return &Stream{__obf_892632c77e859037: __obf_892632c77e859037.(*__obf_972c2293804d141c), __obf_feb7fc81d385b509: __obf_feb7fc81d385b509, __obf_184433571fa55237: make([]byte, 0, __obf_0355f91790eba345), Error: nil, __obf_2dfa32334fb1aa21: 0}
}

// Pool returns a pool can provide more stream with same configuration
func (__obf_850a7457bb739a32 *Stream) Pool() StreamPool {
	return __obf_850a7457bb739a32.

		// Reset reuse this stream instance by assign a new writer
		__obf_892632c77e859037
}

func (__obf_850a7457bb739a32 *Stream) Reset(__obf_feb7fc81d385b509 io.Writer) {
	__obf_850a7457bb739a32.__obf_feb7fc81d385b509 = __obf_feb7fc81d385b509
	__obf_850a7457bb739a32.

		// Available returns how many bytes are unused in the buffer.
		__obf_184433571fa55237 = __obf_850a7457bb739a32.__obf_184433571fa55237[:0]
}

func (__obf_850a7457bb739a32 *Stream) Available() int {
	return cap(__obf_850a7457bb739a32.__obf_184433571fa55237) - len(__obf_850a7457bb739a32.

		// Buffered returns the number of bytes that have been written into the current buffer.
		__obf_184433571fa55237)
}

func (__obf_850a7457bb739a32 *Stream) Buffered() int {
	return len(__obf_850a7457bb739a32.

		// Buffer if writer is nil, use this method to take the result
		__obf_184433571fa55237)
}

func (__obf_850a7457bb739a32 *Stream) Buffer() []byte {
	return __obf_850a7457bb739a32.

		// SetBuffer allows to append to the internal buffer directly
		__obf_184433571fa55237
}

func (__obf_850a7457bb739a32 *Stream) SetBuffer(__obf_184433571fa55237 []byte) {
	__obf_850a7457bb739a32.

		// Write writes the contents of p into the buffer.
		// It returns the number of bytes written.
		// If nn < len(p), it also returns an error explaining
		// why the write is short.
		__obf_184433571fa55237 = __obf_184433571fa55237
}

func (__obf_850a7457bb739a32 *Stream) Write(__obf_bcd7bec668625a79 []byte) (__obf_7b8a0d21934b8993 int, __obf_62967739bca1d11d error) {
	__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_bcd7bec668625a79...)
	if __obf_850a7457bb739a32.__obf_feb7fc81d385b509 != nil {
		__obf_7b8a0d21934b8993, __obf_62967739bca1d11d = __obf_850a7457bb739a32.__obf_feb7fc81d385b509.Write(__obf_850a7457bb739a32.__obf_184433571fa55237)
		__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_850a7457bb739a32.__obf_184433571fa55237[__obf_7b8a0d21934b8993:]
		return
	}
	return len(__obf_bcd7bec668625a79), nil
}

// WriteByte writes a single byte.
func (__obf_850a7457bb739a32 *Stream) __obf_72837f6fe737f078(__obf_f16b4157911bc9af byte) {
	__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_f16b4157911bc9af)
}

func (__obf_850a7457bb739a32 *Stream) __obf_56747a2b5f26ca3c(__obf_4c9093fc2f2347b2 byte, __obf_375d4d212fbe3891 byte) {
	__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_4c9093fc2f2347b2, __obf_375d4d212fbe3891)
}

func (__obf_850a7457bb739a32 *Stream) __obf_acbc9a70f419e9d2(__obf_4c9093fc2f2347b2 byte, __obf_375d4d212fbe3891 byte, __obf_a80f226bd39ca29c byte) {
	__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_4c9093fc2f2347b2, __obf_375d4d212fbe3891, __obf_a80f226bd39ca29c)
}

func (__obf_850a7457bb739a32 *Stream) __obf_624b2ddd28fff892(__obf_4c9093fc2f2347b2 byte, __obf_375d4d212fbe3891 byte, __obf_a80f226bd39ca29c byte, __obf_c0ae8fc634deb1cc byte) {
	__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_4c9093fc2f2347b2, __obf_375d4d212fbe3891, __obf_a80f226bd39ca29c, __obf_c0ae8fc634deb1cc)
}

func (__obf_850a7457bb739a32 *Stream) __obf_8e7dd11cd8538a58(__obf_4c9093fc2f2347b2 byte, __obf_375d4d212fbe3891 byte, __obf_a80f226bd39ca29c byte, __obf_c0ae8fc634deb1cc byte, __obf_a5bce925f39761e9 byte) {
	__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237,

		// Flush writes any buffered data to the underlying io.Writer.
		__obf_4c9093fc2f2347b2, __obf_375d4d212fbe3891, __obf_a80f226bd39ca29c, __obf_c0ae8fc634deb1cc, __obf_a5bce925f39761e9)
}

func (__obf_850a7457bb739a32 *Stream) Flush() error {
	if __obf_850a7457bb739a32.__obf_feb7fc81d385b509 == nil {
		return nil
	}
	if __obf_850a7457bb739a32.Error != nil {
		return __obf_850a7457bb739a32.Error
	}
	_, __obf_62967739bca1d11d := __obf_850a7457bb739a32.__obf_feb7fc81d385b509.Write(__obf_850a7457bb739a32.__obf_184433571fa55237)
	if __obf_62967739bca1d11d != nil {
		if __obf_850a7457bb739a32.Error == nil {
			__obf_850a7457bb739a32.
				Error = __obf_62967739bca1d11d
		}
		return __obf_62967739bca1d11d
	}
	__obf_850a7457bb739a32.__obf_184433571fa55237 = __obf_850a7457bb739a32.__obf_184433571fa55237[:0]
	return nil
}

// WriteRaw write string out without quotes, just like []byte
func (__obf_850a7457bb739a32 *Stream) WriteRaw(__obf_6f3495a136f456ab string) {
	__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.

		// WriteNil write null to stream
		__obf_184433571fa55237, __obf_6f3495a136f456ab...)
}

func (__obf_850a7457bb739a32 *Stream) WriteNil() {
	__obf_850a7457bb739a32.__obf_624b2ddd28fff892('n', 'u', 'l', 'l')
}

// WriteTrue write true to stream
func (__obf_850a7457bb739a32 *Stream) WriteTrue() {
	__obf_850a7457bb739a32.__obf_624b2ddd28fff892('t', 'r', 'u', 'e')
}

// WriteFalse write false to stream
func (__obf_850a7457bb739a32 *Stream) WriteFalse() {
	__obf_850a7457bb739a32.__obf_8e7dd11cd8538a58('f', 'a', 'l', 's', 'e')
}

// WriteBool write true or false into stream
func (__obf_850a7457bb739a32 *Stream) WriteBool(__obf_bbfd2ac8618a6f0c bool) {
	if __obf_bbfd2ac8618a6f0c {
		__obf_850a7457bb739a32.
			WriteTrue()
	} else {
		__obf_850a7457bb739a32.
			WriteFalse()
	}
}

// WriteObjectStart write { with possible indention
func (__obf_850a7457bb739a32 *Stream) WriteObjectStart() {
	__obf_850a7457bb739a32.__obf_2dfa32334fb1aa21 += __obf_850a7457bb739a32.__obf_892632c77e859037.__obf_f186c052066ae54a
	__obf_850a7457bb739a32.

		// WriteObjectField write "field": with possible indention
		__obf_72837f6fe737f078('{')
	__obf_850a7457bb739a32.__obf_937228d15c09cb5e(0)
}

func (__obf_850a7457bb739a32 *Stream) WriteObjectField(__obf_7e01b5b4651074d4 string) {
	__obf_850a7457bb739a32.
		WriteString(__obf_7e01b5b4651074d4)
	if __obf_850a7457bb739a32.__obf_2dfa32334fb1aa21 > 0 {
		__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c(':', ' ')
	} else {
		__obf_850a7457bb739a32.__obf_72837f6fe737f078(':')
	}
}

// WriteObjectEnd write } with possible indention
func (__obf_850a7457bb739a32 *Stream) WriteObjectEnd() {
	__obf_850a7457bb739a32.__obf_937228d15c09cb5e(__obf_850a7457bb739a32.__obf_892632c77e859037.__obf_f186c052066ae54a)
	__obf_850a7457bb739a32.__obf_2dfa32334fb1aa21 -= __obf_850a7457bb739a32.__obf_892632c77e859037.__obf_f186c052066ae54a

	// WriteEmptyObject write {}
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('}')
}

func (__obf_850a7457bb739a32 *Stream) WriteEmptyObject() {
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('{')
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('}')
}

// WriteMore write , with possible indention
func (__obf_850a7457bb739a32 *Stream) WriteMore() {
	__obf_850a7457bb739a32.__obf_72837f6fe737f078(',')
	__obf_850a7457bb739a32.__obf_937228d15c09cb5e(0)
}

// WriteArrayStart write [ with possible indention
func (__obf_850a7457bb739a32 *Stream) WriteArrayStart() {
	__obf_850a7457bb739a32.__obf_2dfa32334fb1aa21 += __obf_850a7457bb739a32.__obf_892632c77e859037.__obf_f186c052066ae54a
	__obf_850a7457bb739a32.

		// WriteEmptyArray write []
		__obf_72837f6fe737f078('[')
	__obf_850a7457bb739a32.__obf_937228d15c09cb5e(0)
}

func (__obf_850a7457bb739a32 *Stream) WriteEmptyArray() {
	__obf_850a7457bb739a32.__obf_56747a2b5f26ca3c('[', ']')
}

// WriteArrayEnd write ] with possible indention
func (__obf_850a7457bb739a32 *Stream) WriteArrayEnd() {
	__obf_850a7457bb739a32.__obf_937228d15c09cb5e(__obf_850a7457bb739a32.__obf_892632c77e859037.__obf_f186c052066ae54a)
	__obf_850a7457bb739a32.__obf_2dfa32334fb1aa21 -= __obf_850a7457bb739a32.__obf_892632c77e859037.__obf_f186c052066ae54a
	__obf_850a7457bb739a32.__obf_72837f6fe737f078(']')
}

func (__obf_850a7457bb739a32 *Stream) __obf_937228d15c09cb5e(__obf_356be6e6f0b0fb2d int) {
	if __obf_850a7457bb739a32.__obf_2dfa32334fb1aa21 == 0 {
		return
	}
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('\n')
	__obf_e456bd59ee15de3b := __obf_850a7457bb739a32.__obf_2dfa32334fb1aa21 - __obf_356be6e6f0b0fb2d
	for __obf_5aa5c8829b97f182 := 0; __obf_5aa5c8829b97f182 < __obf_e456bd59ee15de3b; __obf_5aa5c8829b97f182++ {
		__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, ' ')
	}
}
