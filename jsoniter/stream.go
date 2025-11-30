package __obf_030d39f7a8de96c6

import (
	"io"
)

// stream is a io.Writer like object, with JSON specific write functions.
// Error is not returned as return value, but stored as Error member on this stream instance.
type Stream struct {
	__obf_679611dc56ff465b *__obf_81639538813814ff
	__obf_238de3ec07c7e9da io.Writer
	__obf_0b1656d7ef4bd9df []byte
	Error                  error
	__obf_508619ef3f08fbb3 int
	Attachment             any // open for customized encoder
}

// NewStream create new stream instance.
// cfg can be jsoniter.ConfigDefault.
// out can be nil if write to internal buffer.
// bufSize is the initial size for the internal buffer in bytes.
func NewStream(__obf_679611dc56ff465b API, __obf_238de3ec07c7e9da io.Writer, __obf_8e50793d0b2df740 int) *Stream {
	return &Stream{__obf_679611dc56ff465b: __obf_679611dc56ff465b.(*__obf_81639538813814ff), __obf_238de3ec07c7e9da: __obf_238de3ec07c7e9da, __obf_0b1656d7ef4bd9df: make([]byte, 0, __obf_8e50793d0b2df740), Error: nil, __obf_508619ef3f08fbb3: 0}
}

// Pool returns a pool can provide more stream with same configuration
func (__obf_8a2c51fe22775655 *Stream) Pool() StreamPool {
	return __obf_8a2c51fe22775655.

		// Reset reuse this stream instance by assign a new writer
		__obf_679611dc56ff465b
}

func (__obf_8a2c51fe22775655 *Stream) Reset(__obf_238de3ec07c7e9da io.Writer) {
	__obf_8a2c51fe22775655.__obf_238de3ec07c7e9da = __obf_238de3ec07c7e9da
	__obf_8a2c51fe22775655.

		// Available returns how many bytes are unused in the buffer.
		__obf_0b1656d7ef4bd9df = __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[:0]
}

func (__obf_8a2c51fe22775655 *Stream) Available() int {
	return cap(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df) - len(__obf_8a2c51fe22775655.

		// Buffered returns the number of bytes that have been written into the current buffer.
		__obf_0b1656d7ef4bd9df)
}

func (__obf_8a2c51fe22775655 *Stream) Buffered() int {
	return len(__obf_8a2c51fe22775655.

		// Buffer if writer is nil, use this method to take the result
		__obf_0b1656d7ef4bd9df)
}

func (__obf_8a2c51fe22775655 *Stream) Buffer() []byte {
	return __obf_8a2c51fe22775655.

		// SetBuffer allows to append to the internal buffer directly
		__obf_0b1656d7ef4bd9df
}

func (__obf_8a2c51fe22775655 *Stream) SetBuffer(__obf_0b1656d7ef4bd9df []byte) {
	__obf_8a2c51fe22775655.

		// Write writes the contents of p into the buffer.
		// It returns the number of bytes written.
		// If nn < len(p), it also returns an error explaining
		// why the write is short.
		__obf_0b1656d7ef4bd9df = __obf_0b1656d7ef4bd9df
}

func (__obf_8a2c51fe22775655 *Stream) Write(__obf_23ceea01a3d73f6a []byte) (__obf_b04582d9cbeee4a8 int, __obf_fcc907dd69879566 error) {
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_23ceea01a3d73f6a...)
	if __obf_8a2c51fe22775655.__obf_238de3ec07c7e9da != nil {
		__obf_b04582d9cbeee4a8, __obf_fcc907dd69879566 = __obf_8a2c51fe22775655.__obf_238de3ec07c7e9da.Write(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df)
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[__obf_b04582d9cbeee4a8:]
		return
	}
	return len(__obf_23ceea01a3d73f6a), nil
}

// WriteByte writes a single byte.
func (__obf_8a2c51fe22775655 *Stream) __obf_41130daa346c0e53(__obf_24309b3b0ff9ba22 byte) {
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_24309b3b0ff9ba22)
}

func (__obf_8a2c51fe22775655 *Stream) __obf_3635bad429be5857(__obf_184b1324310379fa byte, __obf_b4edefea91791a88 byte) {
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_184b1324310379fa, __obf_b4edefea91791a88)
}

func (__obf_8a2c51fe22775655 *Stream) __obf_061952646293ff8b(__obf_184b1324310379fa byte, __obf_b4edefea91791a88 byte, __obf_bb721a792780c665 byte) {
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_184b1324310379fa, __obf_b4edefea91791a88, __obf_bb721a792780c665)
}

func (__obf_8a2c51fe22775655 *Stream) __obf_f4ed7f2eb4d35873(__obf_184b1324310379fa byte, __obf_b4edefea91791a88 byte, __obf_bb721a792780c665 byte, __obf_758cb31de84bfed4 byte) {
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_184b1324310379fa, __obf_b4edefea91791a88, __obf_bb721a792780c665, __obf_758cb31de84bfed4)
}

func (__obf_8a2c51fe22775655 *Stream) __obf_f6c006a09024d4ad(__obf_184b1324310379fa byte, __obf_b4edefea91791a88 byte, __obf_bb721a792780c665 byte, __obf_758cb31de84bfed4 byte, __obf_b1b3cbc5fe38ee1c byte) {
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df,

		// Flush writes any buffered data to the underlying io.Writer.
		__obf_184b1324310379fa, __obf_b4edefea91791a88, __obf_bb721a792780c665, __obf_758cb31de84bfed4, __obf_b1b3cbc5fe38ee1c)
}

func (__obf_8a2c51fe22775655 *Stream) Flush() error {
	if __obf_8a2c51fe22775655.__obf_238de3ec07c7e9da == nil {
		return nil
	}
	if __obf_8a2c51fe22775655.Error != nil {
		return __obf_8a2c51fe22775655.Error
	}
	_, __obf_fcc907dd69879566 := __obf_8a2c51fe22775655.__obf_238de3ec07c7e9da.Write(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df)
	if __obf_fcc907dd69879566 != nil {
		if __obf_8a2c51fe22775655.Error == nil {
			__obf_8a2c51fe22775655.
				Error = __obf_fcc907dd69879566
		}
		return __obf_fcc907dd69879566
	}
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = __obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df[:0]
	return nil
}

// WriteRaw write string out without quotes, just like []byte
func (__obf_8a2c51fe22775655 *Stream) WriteRaw(__obf_1d7ed4e74380ec76 string) {
	__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.

		// WriteNil write null to stream
		__obf_0b1656d7ef4bd9df, __obf_1d7ed4e74380ec76...)
}

func (__obf_8a2c51fe22775655 *Stream) WriteNil() {
	__obf_8a2c51fe22775655.__obf_f4ed7f2eb4d35873('n', 'u', 'l', 'l')
}

// WriteTrue write true to stream
func (__obf_8a2c51fe22775655 *Stream) WriteTrue() {
	__obf_8a2c51fe22775655.__obf_f4ed7f2eb4d35873('t', 'r', 'u', 'e')
}

// WriteFalse write false to stream
func (__obf_8a2c51fe22775655 *Stream) WriteFalse() {
	__obf_8a2c51fe22775655.__obf_f6c006a09024d4ad('f', 'a', 'l', 's', 'e')
}

// WriteBool write true or false into stream
func (__obf_8a2c51fe22775655 *Stream) WriteBool(__obf_efaf2719b44ad8ac bool) {
	if __obf_efaf2719b44ad8ac {
		__obf_8a2c51fe22775655.
			WriteTrue()
	} else {
		__obf_8a2c51fe22775655.
			WriteFalse()
	}
}

// WriteObjectStart write { with possible indention
func (__obf_8a2c51fe22775655 *Stream) WriteObjectStart() {
	__obf_8a2c51fe22775655.__obf_508619ef3f08fbb3 += __obf_8a2c51fe22775655.__obf_679611dc56ff465b.__obf_063460864bd2e9ef
	__obf_8a2c51fe22775655.

		// WriteObjectField write "field": with possible indention
		__obf_41130daa346c0e53('{')
	__obf_8a2c51fe22775655.__obf_6724d2a71fd358ab(0)
}

func (__obf_8a2c51fe22775655 *Stream) WriteObjectField(__obf_cd4d02f264b18d55 string) {
	__obf_8a2c51fe22775655.
		WriteString(__obf_cd4d02f264b18d55)
	if __obf_8a2c51fe22775655.__obf_508619ef3f08fbb3 > 0 {
		__obf_8a2c51fe22775655.__obf_3635bad429be5857(':', ' ')
	} else {
		__obf_8a2c51fe22775655.__obf_41130daa346c0e53(':')
	}
}

// WriteObjectEnd write } with possible indention
func (__obf_8a2c51fe22775655 *Stream) WriteObjectEnd() {
	__obf_8a2c51fe22775655.__obf_6724d2a71fd358ab(__obf_8a2c51fe22775655.__obf_679611dc56ff465b.__obf_063460864bd2e9ef)
	__obf_8a2c51fe22775655.__obf_508619ef3f08fbb3 -= __obf_8a2c51fe22775655.__obf_679611dc56ff465b.__obf_063460864bd2e9ef

	// WriteEmptyObject write {}
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('}')
}

func (__obf_8a2c51fe22775655 *Stream) WriteEmptyObject() {
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('{')
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('}')
}

// WriteMore write , with possible indention
func (__obf_8a2c51fe22775655 *Stream) WriteMore() {
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53(',')
	__obf_8a2c51fe22775655.__obf_6724d2a71fd358ab(0)
}

// WriteArrayStart write [ with possible indention
func (__obf_8a2c51fe22775655 *Stream) WriteArrayStart() {
	__obf_8a2c51fe22775655.__obf_508619ef3f08fbb3 += __obf_8a2c51fe22775655.__obf_679611dc56ff465b.__obf_063460864bd2e9ef
	__obf_8a2c51fe22775655.

		// WriteEmptyArray write []
		__obf_41130daa346c0e53('[')
	__obf_8a2c51fe22775655.__obf_6724d2a71fd358ab(0)
}

func (__obf_8a2c51fe22775655 *Stream) WriteEmptyArray() {
	__obf_8a2c51fe22775655.__obf_3635bad429be5857('[', ']')
}

// WriteArrayEnd write ] with possible indention
func (__obf_8a2c51fe22775655 *Stream) WriteArrayEnd() {
	__obf_8a2c51fe22775655.__obf_6724d2a71fd358ab(__obf_8a2c51fe22775655.__obf_679611dc56ff465b.__obf_063460864bd2e9ef)
	__obf_8a2c51fe22775655.__obf_508619ef3f08fbb3 -= __obf_8a2c51fe22775655.__obf_679611dc56ff465b.__obf_063460864bd2e9ef
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53(']')
}

func (__obf_8a2c51fe22775655 *Stream) __obf_6724d2a71fd358ab(__obf_5531c2c40f0b8547 int) {
	if __obf_8a2c51fe22775655.__obf_508619ef3f08fbb3 == 0 {
		return
	}
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('\n')
	__obf_bd39984e4d73aa96 := __obf_8a2c51fe22775655.__obf_508619ef3f08fbb3 - __obf_5531c2c40f0b8547
	for __obf_82c6e05b9d226c58 := 0; __obf_82c6e05b9d226c58 < __obf_bd39984e4d73aa96; __obf_82c6e05b9d226c58++ {
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, ' ')
	}
}
