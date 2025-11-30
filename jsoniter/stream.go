package __obf_703d23ba520c3295

import (
	"io"
)

// stream is a io.Writer like object, with JSON specific write functions.
// Error is not returned as return value, but stored as Error member on this stream instance.
type Stream struct {
	__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78
	__obf_1746319fdfc0f229 io.Writer
	__obf_a065f8e0da5f5952 []byte
	Error                  error
	__obf_1dd0566c8f02217c int
	Attachment             any // open for customized encoder
}

// NewStream create new stream instance.
// cfg can be jsoniter.ConfigDefault.
// out can be nil if write to internal buffer.
// bufSize is the initial size for the internal buffer in bytes.
func NewStream(__obf_25bd4f754a37b862 API, __obf_1746319fdfc0f229 io.Writer, __obf_7da20755ab0b0a83 int) *Stream {
	return &Stream{__obf_25bd4f754a37b862: __obf_25bd4f754a37b862.(*__obf_ef74248d8ae9ba78), __obf_1746319fdfc0f229: __obf_1746319fdfc0f229, __obf_a065f8e0da5f5952: make([]byte, 0, __obf_7da20755ab0b0a83), Error: nil, __obf_1dd0566c8f02217c: 0}
}

// Pool returns a pool can provide more stream with same configuration
func (__obf_9a34f62857fb1d1d *Stream) Pool() StreamPool {
	return __obf_9a34f62857fb1d1d.

		// Reset reuse this stream instance by assign a new writer
		__obf_25bd4f754a37b862
}

func (__obf_9a34f62857fb1d1d *Stream) Reset(__obf_1746319fdfc0f229 io.Writer) {
	__obf_9a34f62857fb1d1d.__obf_1746319fdfc0f229 = __obf_1746319fdfc0f229
	__obf_9a34f62857fb1d1d.

		// Available returns how many bytes are unused in the buffer.
		__obf_a065f8e0da5f5952 = __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[:0]
}

func (__obf_9a34f62857fb1d1d *Stream) Available() int {
	return cap(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952) - len(__obf_9a34f62857fb1d1d.

		// Buffered returns the number of bytes that have been written into the current buffer.
		__obf_a065f8e0da5f5952)
}

func (__obf_9a34f62857fb1d1d *Stream) Buffered() int {
	return len(__obf_9a34f62857fb1d1d.

		// Buffer if writer is nil, use this method to take the result
		__obf_a065f8e0da5f5952)
}

func (__obf_9a34f62857fb1d1d *Stream) Buffer() []byte {
	return __obf_9a34f62857fb1d1d.

		// SetBuffer allows to append to the internal buffer directly
		__obf_a065f8e0da5f5952
}

func (__obf_9a34f62857fb1d1d *Stream) SetBuffer(__obf_a065f8e0da5f5952 []byte) {
	__obf_9a34f62857fb1d1d.

		// Write writes the contents of p into the buffer.
		// It returns the number of bytes written.
		// If nn < len(p), it also returns an error explaining
		// why the write is short.
		__obf_a065f8e0da5f5952 = __obf_a065f8e0da5f5952
}

func (__obf_9a34f62857fb1d1d *Stream) Write(__obf_f9ed235be8f52eaf []byte) (__obf_ee0269ea5a8b7c47 int, __obf_e56742d6e52f642e error) {
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_f9ed235be8f52eaf...)
	if __obf_9a34f62857fb1d1d.__obf_1746319fdfc0f229 != nil {
		__obf_ee0269ea5a8b7c47, __obf_e56742d6e52f642e = __obf_9a34f62857fb1d1d.__obf_1746319fdfc0f229.Write(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952)
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[__obf_ee0269ea5a8b7c47:]
		return
	}
	return len(__obf_f9ed235be8f52eaf), nil
}

// WriteByte writes a single byte.
func (__obf_9a34f62857fb1d1d *Stream) __obf_f7df2a5224462d19(__obf_bd08f5161fff294a byte) {
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_bd08f5161fff294a)
}

func (__obf_9a34f62857fb1d1d *Stream) __obf_e3076a8aafc9c7f0(__obf_a15aedb89f57d7fe byte, __obf_b8ab05f6da0f7a77 byte) {
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_a15aedb89f57d7fe, __obf_b8ab05f6da0f7a77)
}

func (__obf_9a34f62857fb1d1d *Stream) __obf_d273855e5396c495(__obf_a15aedb89f57d7fe byte, __obf_b8ab05f6da0f7a77 byte, __obf_a833a257d56ab74c byte) {
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_a15aedb89f57d7fe, __obf_b8ab05f6da0f7a77, __obf_a833a257d56ab74c)
}

func (__obf_9a34f62857fb1d1d *Stream) __obf_e330632ecc74fc6b(__obf_a15aedb89f57d7fe byte, __obf_b8ab05f6da0f7a77 byte, __obf_a833a257d56ab74c byte, __obf_a9120b28350c9ca6 byte) {
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_a15aedb89f57d7fe, __obf_b8ab05f6da0f7a77, __obf_a833a257d56ab74c, __obf_a9120b28350c9ca6)
}

func (__obf_9a34f62857fb1d1d *Stream) __obf_4e279bd7bbc88748(__obf_a15aedb89f57d7fe byte, __obf_b8ab05f6da0f7a77 byte, __obf_a833a257d56ab74c byte, __obf_a9120b28350c9ca6 byte, __obf_6771eef086336c37 byte) {
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952,

		// Flush writes any buffered data to the underlying io.Writer.
		__obf_a15aedb89f57d7fe, __obf_b8ab05f6da0f7a77, __obf_a833a257d56ab74c, __obf_a9120b28350c9ca6, __obf_6771eef086336c37)
}

func (__obf_9a34f62857fb1d1d *Stream) Flush() error {
	if __obf_9a34f62857fb1d1d.__obf_1746319fdfc0f229 == nil {
		return nil
	}
	if __obf_9a34f62857fb1d1d.Error != nil {
		return __obf_9a34f62857fb1d1d.Error
	}
	_, __obf_e56742d6e52f642e := __obf_9a34f62857fb1d1d.__obf_1746319fdfc0f229.Write(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952)
	if __obf_e56742d6e52f642e != nil {
		if __obf_9a34f62857fb1d1d.Error == nil {
			__obf_9a34f62857fb1d1d.
				Error = __obf_e56742d6e52f642e
		}
		return __obf_e56742d6e52f642e
	}
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = __obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952[:0]
	return nil
}

// WriteRaw write string out without quotes, just like []byte
func (__obf_9a34f62857fb1d1d *Stream) WriteRaw(__obf_61541a8d4689037b string) {
	__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.

		// WriteNil write null to stream
		__obf_a065f8e0da5f5952, __obf_61541a8d4689037b...)
}

func (__obf_9a34f62857fb1d1d *Stream) WriteNil() {
	__obf_9a34f62857fb1d1d.__obf_e330632ecc74fc6b('n', 'u', 'l', 'l')
}

// WriteTrue write true to stream
func (__obf_9a34f62857fb1d1d *Stream) WriteTrue() {
	__obf_9a34f62857fb1d1d.__obf_e330632ecc74fc6b('t', 'r', 'u', 'e')
}

// WriteFalse write false to stream
func (__obf_9a34f62857fb1d1d *Stream) WriteFalse() {
	__obf_9a34f62857fb1d1d.__obf_4e279bd7bbc88748('f', 'a', 'l', 's', 'e')
}

// WriteBool write true or false into stream
func (__obf_9a34f62857fb1d1d *Stream) WriteBool(__obf_b924538fffe5fd64 bool) {
	if __obf_b924538fffe5fd64 {
		__obf_9a34f62857fb1d1d.
			WriteTrue()
	} else {
		__obf_9a34f62857fb1d1d.
			WriteFalse()
	}
}

// WriteObjectStart write { with possible indention
func (__obf_9a34f62857fb1d1d *Stream) WriteObjectStart() {
	__obf_9a34f62857fb1d1d.__obf_1dd0566c8f02217c += __obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.__obf_f5558fa1bf957433
	__obf_9a34f62857fb1d1d.

		// WriteObjectField write "field": with possible indention
		__obf_f7df2a5224462d19('{')
	__obf_9a34f62857fb1d1d.__obf_cca1920d9fb159c4(0)
}

func (__obf_9a34f62857fb1d1d *Stream) WriteObjectField(__obf_13f6440419533990 string) {
	__obf_9a34f62857fb1d1d.
		WriteString(__obf_13f6440419533990)
	if __obf_9a34f62857fb1d1d.__obf_1dd0566c8f02217c > 0 {
		__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0(':', ' ')
	} else {
		__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19(':')
	}
}

// WriteObjectEnd write } with possible indention
func (__obf_9a34f62857fb1d1d *Stream) WriteObjectEnd() {
	__obf_9a34f62857fb1d1d.__obf_cca1920d9fb159c4(__obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.__obf_f5558fa1bf957433)
	__obf_9a34f62857fb1d1d.__obf_1dd0566c8f02217c -= __obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.__obf_f5558fa1bf957433

	// WriteEmptyObject write {}
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('}')
}

func (__obf_9a34f62857fb1d1d *Stream) WriteEmptyObject() {
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('{')
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('}')
}

// WriteMore write , with possible indention
func (__obf_9a34f62857fb1d1d *Stream) WriteMore() {
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19(',')
	__obf_9a34f62857fb1d1d.__obf_cca1920d9fb159c4(0)
}

// WriteArrayStart write [ with possible indention
func (__obf_9a34f62857fb1d1d *Stream) WriteArrayStart() {
	__obf_9a34f62857fb1d1d.__obf_1dd0566c8f02217c += __obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.__obf_f5558fa1bf957433
	__obf_9a34f62857fb1d1d.

		// WriteEmptyArray write []
		__obf_f7df2a5224462d19('[')
	__obf_9a34f62857fb1d1d.__obf_cca1920d9fb159c4(0)
}

func (__obf_9a34f62857fb1d1d *Stream) WriteEmptyArray() {
	__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0('[', ']')
}

// WriteArrayEnd write ] with possible indention
func (__obf_9a34f62857fb1d1d *Stream) WriteArrayEnd() {
	__obf_9a34f62857fb1d1d.__obf_cca1920d9fb159c4(__obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.__obf_f5558fa1bf957433)
	__obf_9a34f62857fb1d1d.__obf_1dd0566c8f02217c -= __obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.__obf_f5558fa1bf957433
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19(']')
}

func (__obf_9a34f62857fb1d1d *Stream) __obf_cca1920d9fb159c4(__obf_4935b4e0cf94755b int) {
	if __obf_9a34f62857fb1d1d.__obf_1dd0566c8f02217c == 0 {
		return
	}
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('\n')
	__obf_8043e44f5f0decfb := __obf_9a34f62857fb1d1d.__obf_1dd0566c8f02217c - __obf_4935b4e0cf94755b
	for __obf_b0a5d2bd48690f1d := 0; __obf_b0a5d2bd48690f1d < __obf_8043e44f5f0decfb; __obf_b0a5d2bd48690f1d++ {
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, ' ')
	}
}
