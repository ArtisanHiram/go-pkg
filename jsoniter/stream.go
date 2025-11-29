package __obf_5b802ce8d9ba56d6

import (
	"io"
)

// stream is a io.Writer like object, with JSON specific write functions.
// Error is not returned as return value, but stored as Error member on this stream instance.
type Stream struct {
	__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf
	__obf_3d1a004850652c9c io.Writer
	__obf_9fc06d9180f0daca []byte
	Error                  error
	__obf_e68cefffff828d90 int
	Attachment             any // open for customized encoder
}

// NewStream create new stream instance.
// cfg can be jsoniter.ConfigDefault.
// out can be nil if write to internal buffer.
// bufSize is the initial size for the internal buffer in bytes.
func NewStream(__obf_dca106e1cdf85392 API, __obf_3d1a004850652c9c io.Writer, __obf_3f8315b01fa921cf int) *Stream {
	return &Stream{__obf_dca106e1cdf85392: __obf_dca106e1cdf85392.(*__obf_5d13d7f3bd06c6cf), __obf_3d1a004850652c9c: __obf_3d1a004850652c9c, __obf_9fc06d9180f0daca: make([]byte, 0, __obf_3f8315b01fa921cf), Error: nil, __obf_e68cefffff828d90: 0}
}

// Pool returns a pool can provide more stream with same configuration
func (__obf_00fc491caa967a74 *Stream) Pool() StreamPool {
	return __obf_00fc491caa967a74.

		// Reset reuse this stream instance by assign a new writer
		__obf_dca106e1cdf85392
}

func (__obf_00fc491caa967a74 *Stream) Reset(__obf_3d1a004850652c9c io.Writer) {
	__obf_00fc491caa967a74.__obf_3d1a004850652c9c = __obf_3d1a004850652c9c
	__obf_00fc491caa967a74.

		// Available returns how many bytes are unused in the buffer.
		__obf_9fc06d9180f0daca = __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[:0]
}

func (__obf_00fc491caa967a74 *Stream) Available() int {
	return cap(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca) - len(__obf_00fc491caa967a74.

		// Buffered returns the number of bytes that have been written into the current buffer.
		__obf_9fc06d9180f0daca)
}

func (__obf_00fc491caa967a74 *Stream) Buffered() int {
	return len(__obf_00fc491caa967a74.

		// Buffer if writer is nil, use this method to take the result
		__obf_9fc06d9180f0daca)
}

func (__obf_00fc491caa967a74 *Stream) Buffer() []byte {
	return __obf_00fc491caa967a74.

		// SetBuffer allows to append to the internal buffer directly
		__obf_9fc06d9180f0daca
}

func (__obf_00fc491caa967a74 *Stream) SetBuffer(__obf_9fc06d9180f0daca []byte) {
	__obf_00fc491caa967a74.

		// Write writes the contents of p into the buffer.
		// It returns the number of bytes written.
		// If nn < len(p), it also returns an error explaining
		// why the write is short.
		__obf_9fc06d9180f0daca = __obf_9fc06d9180f0daca
}

func (__obf_00fc491caa967a74 *Stream) Write(__obf_82ddba9040ff8d8c []byte) (__obf_4eed202611921696 int, __obf_6d071d48f3b321ad error) {
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_82ddba9040ff8d8c...)
	if __obf_00fc491caa967a74.__obf_3d1a004850652c9c != nil {
		__obf_4eed202611921696, __obf_6d071d48f3b321ad = __obf_00fc491caa967a74.__obf_3d1a004850652c9c.Write(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca)
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[__obf_4eed202611921696:]
		return
	}
	return len(__obf_82ddba9040ff8d8c), nil
}

// WriteByte writes a single byte.
func (__obf_00fc491caa967a74 *Stream) __obf_487ee8035c7dd8f6(__obf_dab9baaadfa7c8c2 byte) {
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_dab9baaadfa7c8c2)
}

func (__obf_00fc491caa967a74 *Stream) __obf_dd59782cf69113a3(__obf_eb472fd1daa794d1 byte, __obf_6b316bd1116d017a byte) {
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_eb472fd1daa794d1, __obf_6b316bd1116d017a)
}

func (__obf_00fc491caa967a74 *Stream) __obf_3b5ca53029ea9496(__obf_eb472fd1daa794d1 byte, __obf_6b316bd1116d017a byte, __obf_db2ca00fe147906b byte) {
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_eb472fd1daa794d1, __obf_6b316bd1116d017a, __obf_db2ca00fe147906b)
}

func (__obf_00fc491caa967a74 *Stream) __obf_d39fc8b8cf9e8086(__obf_eb472fd1daa794d1 byte, __obf_6b316bd1116d017a byte, __obf_db2ca00fe147906b byte, __obf_f5c02955c675744e byte) {
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_eb472fd1daa794d1, __obf_6b316bd1116d017a, __obf_db2ca00fe147906b, __obf_f5c02955c675744e)
}

func (__obf_00fc491caa967a74 *Stream) __obf_16cd09460493426a(__obf_eb472fd1daa794d1 byte, __obf_6b316bd1116d017a byte, __obf_db2ca00fe147906b byte, __obf_f5c02955c675744e byte, __obf_3b887b6c4c185da4 byte) {
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca,

		// Flush writes any buffered data to the underlying io.Writer.
		__obf_eb472fd1daa794d1, __obf_6b316bd1116d017a, __obf_db2ca00fe147906b, __obf_f5c02955c675744e, __obf_3b887b6c4c185da4)
}

func (__obf_00fc491caa967a74 *Stream) Flush() error {
	if __obf_00fc491caa967a74.__obf_3d1a004850652c9c == nil {
		return nil
	}
	if __obf_00fc491caa967a74.Error != nil {
		return __obf_00fc491caa967a74.Error
	}
	_, __obf_6d071d48f3b321ad := __obf_00fc491caa967a74.__obf_3d1a004850652c9c.Write(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca)
	if __obf_6d071d48f3b321ad != nil {
		if __obf_00fc491caa967a74.Error == nil {
			__obf_00fc491caa967a74.
				Error = __obf_6d071d48f3b321ad
		}
		return __obf_6d071d48f3b321ad
	}
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = __obf_00fc491caa967a74.__obf_9fc06d9180f0daca[:0]
	return nil
}

// WriteRaw write string out without quotes, just like []byte
func (__obf_00fc491caa967a74 *Stream) WriteRaw(__obf_5ba76a0156a3a826 string) {
	__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.

		// WriteNil write null to stream
		__obf_9fc06d9180f0daca, __obf_5ba76a0156a3a826...)
}

func (__obf_00fc491caa967a74 *Stream) WriteNil() {
	__obf_00fc491caa967a74.__obf_d39fc8b8cf9e8086('n', 'u', 'l', 'l')
}

// WriteTrue write true to stream
func (__obf_00fc491caa967a74 *Stream) WriteTrue() {
	__obf_00fc491caa967a74.__obf_d39fc8b8cf9e8086('t', 'r', 'u', 'e')
}

// WriteFalse write false to stream
func (__obf_00fc491caa967a74 *Stream) WriteFalse() {
	__obf_00fc491caa967a74.__obf_16cd09460493426a('f', 'a', 'l', 's', 'e')
}

// WriteBool write true or false into stream
func (__obf_00fc491caa967a74 *Stream) WriteBool(__obf_5406b11e33b9d1d3 bool) {
	if __obf_5406b11e33b9d1d3 {
		__obf_00fc491caa967a74.
			WriteTrue()
	} else {
		__obf_00fc491caa967a74.
			WriteFalse()
	}
}

// WriteObjectStart write { with possible indention
func (__obf_00fc491caa967a74 *Stream) WriteObjectStart() {
	__obf_00fc491caa967a74.__obf_e68cefffff828d90 += __obf_00fc491caa967a74.__obf_dca106e1cdf85392.__obf_ee9b4b613c13fea6
	__obf_00fc491caa967a74.

		// WriteObjectField write "field": with possible indention
		__obf_487ee8035c7dd8f6('{')
	__obf_00fc491caa967a74.__obf_f56f94fa6609002e(0)
}

func (__obf_00fc491caa967a74 *Stream) WriteObjectField(__obf_61998fb083387c3c string) {
	__obf_00fc491caa967a74.
		WriteString(__obf_61998fb083387c3c)
	if __obf_00fc491caa967a74.__obf_e68cefffff828d90 > 0 {
		__obf_00fc491caa967a74.__obf_dd59782cf69113a3(':', ' ')
	} else {
		__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6(':')
	}
}

// WriteObjectEnd write } with possible indention
func (__obf_00fc491caa967a74 *Stream) WriteObjectEnd() {
	__obf_00fc491caa967a74.__obf_f56f94fa6609002e(__obf_00fc491caa967a74.__obf_dca106e1cdf85392.__obf_ee9b4b613c13fea6)
	__obf_00fc491caa967a74.__obf_e68cefffff828d90 -= __obf_00fc491caa967a74.__obf_dca106e1cdf85392.__obf_ee9b4b613c13fea6

	// WriteEmptyObject write {}
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('}')
}

func (__obf_00fc491caa967a74 *Stream) WriteEmptyObject() {
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('{')
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('}')
}

// WriteMore write , with possible indention
func (__obf_00fc491caa967a74 *Stream) WriteMore() {
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6(',')
	__obf_00fc491caa967a74.__obf_f56f94fa6609002e(0)
}

// WriteArrayStart write [ with possible indention
func (__obf_00fc491caa967a74 *Stream) WriteArrayStart() {
	__obf_00fc491caa967a74.__obf_e68cefffff828d90 += __obf_00fc491caa967a74.__obf_dca106e1cdf85392.__obf_ee9b4b613c13fea6
	__obf_00fc491caa967a74.

		// WriteEmptyArray write []
		__obf_487ee8035c7dd8f6('[')
	__obf_00fc491caa967a74.__obf_f56f94fa6609002e(0)
}

func (__obf_00fc491caa967a74 *Stream) WriteEmptyArray() {
	__obf_00fc491caa967a74.__obf_dd59782cf69113a3('[', ']')
}

// WriteArrayEnd write ] with possible indention
func (__obf_00fc491caa967a74 *Stream) WriteArrayEnd() {
	__obf_00fc491caa967a74.__obf_f56f94fa6609002e(__obf_00fc491caa967a74.__obf_dca106e1cdf85392.__obf_ee9b4b613c13fea6)
	__obf_00fc491caa967a74.__obf_e68cefffff828d90 -= __obf_00fc491caa967a74.__obf_dca106e1cdf85392.__obf_ee9b4b613c13fea6
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6(']')
}

func (__obf_00fc491caa967a74 *Stream) __obf_f56f94fa6609002e(__obf_8e3421b8ffbdf57b int) {
	if __obf_00fc491caa967a74.__obf_e68cefffff828d90 == 0 {
		return
	}
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('\n')
	__obf_185f31fc8313f634 := __obf_00fc491caa967a74.__obf_e68cefffff828d90 - __obf_8e3421b8ffbdf57b
	for __obf_2deec7c38ffb6ae3 := 0; __obf_2deec7c38ffb6ae3 < __obf_185f31fc8313f634; __obf_2deec7c38ffb6ae3++ {
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, ' ')
	}
}
