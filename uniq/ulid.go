package __obf_4d7a51f8e2f0494a

import (
	"bufio"
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"io"
	"math"
	"math/bits"
	"math/rand"
	"time"
)

/*
An ULID is a 16 byte Universally Unique Lexicographically Sortable Identifier

	The components are encoded as 16 octets.
	Each component is encoded with the MSB first (network byte order).
	0                   1                   2                   3
	0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|                      32_bit_uint_time_high                    |
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|     16_bit_uint_time_low      |       16_bit_uint_random      |
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|                       32_bit_uint_random                      |
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|                       32_bit_uint_random                      |
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
*/
type ULID [16]byte

var (
	// ErrDataSize is returned when parsing or unmarshaling ULIDs with the wrong
	// data size.
	ErrDataSize = errors.New("ulid: bad data size when unmarshaling")

	// ErrInvalidCharacters is returned when parsing or unmarshaling ULIDs with
	// invalid Base32 encodings.
	ErrInvalidCharacters = errors.New("ulid: bad data characters when unmarshaling")

	// ErrBufferSize is returned when marshalling ULIDs to a buffer of insufficient
	// size.
	ErrBufferSize = errors.New("ulid: bad buffer size when marshaling")

	// ErrBigTime is returned when constructing an ULID with a time that is larger
	// than MaxTime.
	ErrBigTime = errors.New("ulid: time too big")

	// ErrOverflow is returned when unmarshaling a ULID whose first character is
	// larger than 7, thereby exceeding the valid bit depth of 128.
	ErrOverflow = errors.New("ulid: overflow when unmarshaling")

	// ErrMonotonicOverflow is returned by a Monotonic entropy source when
	// incrementing the previous ULID's entropy bytes would result in overflow.
	ErrMonotonicOverflow = errors.New("ulid: monotonic entropy overflow")

	// ErrScanValue is returned when the value passed to scan cannot be unmarshaled
	// into the ULID.
	ErrScanValue = errors.New("ulid: source value must be a string or byte slice")
)

// MonotonicReader is an interface that should yield monotonically increasing
// entropy into the provided slice for all calls with the same ms parameter. If
// a MonotonicReader is provided to the New constructor, its MonotonicRead
// method will be used instead of Read.
type MonotonicReader interface {
	io.Reader
	MonotonicRead(__obf_3d28b9e3414c1273 uint64, __obf_8eaa13c5aab512f0 []byte) error
}

// New returns an ULID with the given Unix milliseconds timestamp and an
// optional entropy source. Use the Timestamp function to convert
// a time.Time to Unix milliseconds.
//
// ErrBigTime is returned when passing a timestamp bigger than MaxTime.
// Reading from the entropy source may also return an error.
//
// Safety for concurrent use is only dependent on the safety of the
// entropy source.
func New(__obf_3d28b9e3414c1273 uint64, __obf_e867f2cef457c33d io.Reader) (__obf_721283c42292359f ULID, __obf_38c5d446257f8d9a error) {
	if __obf_38c5d446257f8d9a = __obf_721283c42292359f.SetTime(__obf_3d28b9e3414c1273); __obf_38c5d446257f8d9a != nil {
		return __obf_721283c42292359f, __obf_38c5d446257f8d9a
	}

	switch __obf_9637d6eccc44ff3f := __obf_e867f2cef457c33d.(type) {
	case nil:
		return __obf_721283c42292359f, __obf_38c5d446257f8d9a
	case MonotonicReader:
		__obf_38c5d446257f8d9a = __obf_9637d6eccc44ff3f.MonotonicRead(__obf_3d28b9e3414c1273, __obf_721283c42292359f[6:])
	default:
		_, __obf_38c5d446257f8d9a = io.ReadFull(__obf_9637d6eccc44ff3f, __obf_721283c42292359f[6:])
	}

	return __obf_721283c42292359f,

		// MustNew is a convenience function equivalent to New that panics on failure
		// instead of returning an error.
		__obf_38c5d446257f8d9a
}

func MustNew(__obf_3d28b9e3414c1273 uint64, __obf_e867f2cef457c33d io.Reader) ULID {
	__obf_721283c42292359f, __obf_38c5d446257f8d9a := New(__obf_3d28b9e3414c1273, __obf_e867f2cef457c33d)
	if __obf_38c5d446257f8d9a != nil {
		panic(__obf_38c5d446257f8d9a)
	}
	return __obf_721283c42292359f
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(__obf_bc767bd04b3a0a5e string) (__obf_721283c42292359f ULID, __obf_38c5d446257f8d9a error) {
	return __obf_721283c42292359f, __obf_bb1718f494bf684d([]byte(__obf_bc767bd04b3a0a5e), false, &__obf_721283c42292359f)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(__obf_bc767bd04b3a0a5e string) (__obf_721283c42292359f ULID, __obf_38c5d446257f8d9a error) {
	return __obf_721283c42292359f, __obf_bb1718f494bf684d([]byte(__obf_bc767bd04b3a0a5e), true, &__obf_721283c42292359f)
}

func __obf_bb1718f494bf684d(__obf_ec9ea9e2ee15569a []byte, __obf_c1953c9f576b3bf3 bool, __obf_721283c42292359f *ULID) error {
	// Check if a base32 encoded ULID is the right length.
	if len(__obf_ec9ea9e2ee15569a) != EncodedSize {
		return ErrDataSize
	}

	// Check if all the characters in a base32 encoded ULID are part of the
	// expected base32 character set.
	if __obf_c1953c9f576b3bf3 &&
		(__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[0]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[1]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[2]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[3]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[4]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[5]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[6]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[7]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[8]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[9]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[10]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[11]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[12]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[13]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[14]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[15]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[16]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[17]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[18]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[19]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[20]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[21]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[22]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[23]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[24]] == 0xFF || __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[25]] == 0xFF) {
		return ErrInvalidCharacters
	}

	// Check if the first character in a base32 encoded ULID will overflow. This
	// happens because the base32 representation encodes 130 bits, while the
	// ULID is only 128 bits.
	//
	// See https://github.com/oklog/ulid/issues/9 for details.
	if __obf_ec9ea9e2ee15569a[0] > '7' {
		return ErrOverflow
	}

	// Use an optimized unrolled loop (from https://github.com/RobThree/NUlid)
	// to decode a base32 ULID.

	// 6 bytes timestamp (48 bits)
	(*__obf_721283c42292359f)[0] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[0]] << 5) | __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[1]]
	(*__obf_721283c42292359f)[1] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[2]] << 3) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[3]] >> 2)
	(*__obf_721283c42292359f)[2] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[3]] << 6) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[4]] << 1) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[5]] >> 4)
	(*__obf_721283c42292359f)[3] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[5]] << 4) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[6]] >> 1)
	(*__obf_721283c42292359f)[4] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[6]] << 7) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[7]] << 2) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[8]] >> 3)
	(*__obf_721283c42292359f)[5] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[8]] << 5) | __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[9]]

	// 10 bytes of entropy (80 bits)
	(*__obf_721283c42292359f)[6] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[10]] << 3) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[11]] >> 2)
	(*__obf_721283c42292359f)[7] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[11]] << 6) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[12]] << 1) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[13]] >> 4)
	(*__obf_721283c42292359f)[8] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[13]] << 4) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[14]] >> 1)
	(*__obf_721283c42292359f)[9] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[14]] << 7) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[15]] << 2) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[16]] >> 3)
	(*__obf_721283c42292359f)[10] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[16]] << 5) | __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[17]]
	(*__obf_721283c42292359f)[11] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[18]] << 3) | __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[19]]>>2
	(*__obf_721283c42292359f)[12] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[19]] << 6) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[20]] << 1) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[21]] >> 4)
	(*__obf_721283c42292359f)[13] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[21]] << 4) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[22]] >> 1)
	(*__obf_721283c42292359f)[14] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[22]] << 7) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[23]] << 2) | (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[24]] >> 3)
	(*__obf_721283c42292359f)[15] = (__obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[24]] << 5) | __obf_c289e5e7223a899b[__obf_ec9ea9e2ee15569a[25]]

	return nil
}

// MustParse is a convenience function equivalent to Parse that panics on failure
// instead of returning an error.
func MustParse(__obf_bc767bd04b3a0a5e string) ULID {
	__obf_721283c42292359f, __obf_38c5d446257f8d9a := Parse(__obf_bc767bd04b3a0a5e)
	if __obf_38c5d446257f8d9a != nil {
		panic(__obf_38c5d446257f8d9a)
	}
	return __obf_721283c42292359f
}

// MustParseStrict is a convenience function equivalent to ParseStrict that
// panics on failure instead of returning an error.
func MustParseStrict(__obf_bc767bd04b3a0a5e string) ULID {
	__obf_721283c42292359f, __obf_38c5d446257f8d9a := ParseStrict(__obf_bc767bd04b3a0a5e)
	if __obf_38c5d446257f8d9a != nil {
		panic(__obf_38c5d446257f8d9a)
	}
	return __obf_721283c42292359f
}

// Bytes returns bytes slice representation of ULID.
func (__obf_51362cca9c9190a3 ULID) Bytes() []byte {
	return __obf_51362cca9c9190a3[:]
}

// String returns a lexicographically sortable string encoded ULID
// (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy
func (__obf_721283c42292359f ULID) String() string {
	__obf_bc767bd04b3a0a5e := make([]byte, EncodedSize)
	_ = __obf_721283c42292359f.MarshalTextTo(__obf_bc767bd04b3a0a5e)
	return string(__obf_bc767bd04b3a0a5e)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (__obf_721283c42292359f ULID) MarshalBinary() ([]byte, error) {
	__obf_bc767bd04b3a0a5e := make([]byte, len(__obf_721283c42292359f))
	return __obf_bc767bd04b3a0a5e, __obf_721283c42292359f.MarshalBinaryTo(__obf_bc767bd04b3a0a5e)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (__obf_721283c42292359f ULID) MarshalBinaryTo(__obf_487aa814dfac9f67 []byte) error {
	if len(__obf_487aa814dfac9f67) != len(__obf_721283c42292359f) {
		return ErrBufferSize
	}

	copy(__obf_487aa814dfac9f67, __obf_721283c42292359f[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (__obf_721283c42292359f *ULID) UnmarshalBinary(__obf_78210b9635f4f64b []byte) error {
	if len(__obf_78210b9635f4f64b) != len(*__obf_721283c42292359f) {
		return ErrDataSize
	}

	copy((*__obf_721283c42292359f)[:], __obf_78210b9635f4f64b)
	return nil
}

// Encoding is the base 32 encoding alphabet used in ULID strings.
const Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (__obf_721283c42292359f ULID) MarshalText() ([]byte, error) {
	__obf_bc767bd04b3a0a5e := make([]byte, EncodedSize)
	return __obf_bc767bd04b3a0a5e, __obf_721283c42292359f.MarshalTextTo(__obf_bc767bd04b3a0a5e)
}

// MarshalTextTo writes the ULID as a string to the given buffer.
// ErrBufferSize is returned when the len(dst) != 26.
func (__obf_721283c42292359f ULID) MarshalTextTo(__obf_487aa814dfac9f67 []byte) error {
	// Optimized unrolled loop ahead.
	// From https://github.com/RobThree/NUlid

	if len(__obf_487aa814dfac9f67) != EncodedSize {
		return ErrBufferSize
	}
	__obf_487aa814dfac9f67[ // 10 byte timestamp
	0] = Encoding[(__obf_721283c42292359f[0]&224)>>5]
	__obf_487aa814dfac9f67[1] = Encoding[__obf_721283c42292359f[0]&31]
	__obf_487aa814dfac9f67[2] = Encoding[(__obf_721283c42292359f[1]&248)>>3]
	__obf_487aa814dfac9f67[3] = Encoding[((__obf_721283c42292359f[1]&7)<<2)|((__obf_721283c42292359f[2]&192)>>6)]
	__obf_487aa814dfac9f67[4] = Encoding[(__obf_721283c42292359f[2]&62)>>1]
	__obf_487aa814dfac9f67[5] = Encoding[((__obf_721283c42292359f[2]&1)<<4)|((__obf_721283c42292359f[3]&240)>>4)]
	__obf_487aa814dfac9f67[6] = Encoding[((__obf_721283c42292359f[3]&15)<<1)|((__obf_721283c42292359f[4]&128)>>7)]
	__obf_487aa814dfac9f67[7] = Encoding[(__obf_721283c42292359f[4]&124)>>2]
	__obf_487aa814dfac9f67[8] = Encoding[((__obf_721283c42292359f[4]&3)<<3)|((__obf_721283c42292359f[5]&224)>>5)]
	__obf_487aa814dfac9f67[9] = Encoding[__obf_721283c42292359f[5]&31]
	__obf_487aa814dfac9f67[ // 16 bytes of entropy
	10] = Encoding[(__obf_721283c42292359f[6]&248)>>3]
	__obf_487aa814dfac9f67[11] = Encoding[((__obf_721283c42292359f[6]&7)<<2)|((__obf_721283c42292359f[7]&192)>>6)]
	__obf_487aa814dfac9f67[12] = Encoding[(__obf_721283c42292359f[7]&62)>>1]
	__obf_487aa814dfac9f67[13] = Encoding[((__obf_721283c42292359f[7]&1)<<4)|((__obf_721283c42292359f[8]&240)>>4)]
	__obf_487aa814dfac9f67[14] = Encoding[((__obf_721283c42292359f[8]&15)<<1)|((__obf_721283c42292359f[9]&128)>>7)]
	__obf_487aa814dfac9f67[15] = Encoding[(__obf_721283c42292359f[9]&124)>>2]
	__obf_487aa814dfac9f67[16] = Encoding[((__obf_721283c42292359f[9]&3)<<3)|((__obf_721283c42292359f[10]&224)>>5)]
	__obf_487aa814dfac9f67[17] = Encoding[__obf_721283c42292359f[10]&31]
	__obf_487aa814dfac9f67[18] = Encoding[(__obf_721283c42292359f[11]&248)>>3]
	__obf_487aa814dfac9f67[19] = Encoding[((__obf_721283c42292359f[11]&7)<<2)|((__obf_721283c42292359f[12]&192)>>6)]
	__obf_487aa814dfac9f67[20] = Encoding[(__obf_721283c42292359f[12]&62)>>1]
	__obf_487aa814dfac9f67[21] = Encoding[((__obf_721283c42292359f[12]&1)<<4)|((__obf_721283c42292359f[13]&240)>>4)]
	__obf_487aa814dfac9f67[22] = Encoding[((__obf_721283c42292359f[13]&15)<<1)|((__obf_721283c42292359f[14]&128)>>7)]
	__obf_487aa814dfac9f67[23] = Encoding[(__obf_721283c42292359f[14]&124)>>2]
	__obf_487aa814dfac9f67[24] = Encoding[((__obf_721283c42292359f[14]&3)<<3)|((__obf_721283c42292359f[15]&224)>>5)]
	__obf_487aa814dfac9f67[25] = Encoding[__obf_721283c42292359f[15]&31]

	return nil
}

// Byte to index table for O(1) lookups when unmarshaling.
// We use 0xFF as sentinel value for invalid indexes.
var __obf_c289e5e7223a899b = [...]byte{
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x01,
	0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
	0x0F, 0x10, 0x11, 0xFF, 0x12, 0x13, 0xFF, 0x14, 0x15, 0xFF,
	0x16, 0x17, 0x18, 0x19, 0x1A, 0xFF, 0x1B, 0x1C, 0x1D, 0x1E,
	0x1F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0A, 0x0B, 0x0C,
	0x0D, 0x0E, 0x0F, 0x10, 0x11, 0xFF, 0x12, 0x13, 0xFF, 0x14,
	0x15, 0xFF, 0x16, 0x17, 0x18, 0x19, 0x1A, 0xFF, 0x1B, 0x1C,
	0x1D, 0x1E, 0x1F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
}

// EncodedSize is the length of a text encoded ULID.
const EncodedSize = 26

// UnmarshalText implements the encoding.TextUnmarshaler interface by
// parsing the data as string encoded ULID.
//
// ErrDataSize is returned if the len(v) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs.
func (__obf_721283c42292359f *ULID) UnmarshalText(__obf_ec9ea9e2ee15569a []byte) error {
	return __obf_bb1718f494bf684d(__obf_ec9ea9e2ee15569a, false, __obf_721283c42292359f)
}

// Time returns the Unix time in milliseconds encoded in the ULID.
// Use the top level Time function to convert the returned value to
// a time.Time.
func (__obf_721283c42292359f ULID) Time() uint64 {
	return uint64(__obf_721283c42292359f[5]) | uint64(__obf_721283c42292359f[4])<<8 |
		uint64(__obf_721283c42292359f[3])<<16 | uint64(__obf_721283c42292359f[2])<<24 |
		uint64(__obf_721283c42292359f[1])<<32 | uint64(__obf_721283c42292359f[0])<<40
}

// maxTime is the maximum Unix time in milliseconds that can be
// represented in an ULID.
var __obf_6bda0af3d1882741 = ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}.Time()

// MaxTime returns the maximum Unix time in milliseconds that
// can be encoded in an ULID.
func MaxTime() uint64 { return __obf_6bda0af3d1882741 }

// Now is a convenience function that returns the current
// UTC time in Unix milliseconds. Equivalent to:
//
//	Timestamp(time.Now().UTC())
func Now() uint64 { return Timestamp(time.Now().UTC()) }

// Timestamp converts a time.Time to Unix milliseconds.
//
// Because of the way ULID stores time, times from the year
// 10889 produces undefined results.
func Timestamp(__obf_d5aaddb5c5a3e313 time.Time) uint64 {
	return uint64(__obf_d5aaddb5c5a3e313.Unix())*1000 +
		uint64(__obf_d5aaddb5c5a3e313.Nanosecond()/int(time.Millisecond))
}

// Time converts Unix milliseconds in the format
// returned by the Timestamp function to a time.Time.
func Time(__obf_3d28b9e3414c1273 uint64) time.Time {
	__obf_cf2faa899ba218b9 := int64(__obf_3d28b9e3414c1273 / 1e3)
	__obf_c4a8aa5127f9d983 := int64((__obf_3d28b9e3414c1273 % 1e3) * 1e6)
	return time.Unix(__obf_cf2faa899ba218b9,

		// SetTime sets the time component of the ULID to the given Unix time
		// in milliseconds.
		__obf_c4a8aa5127f9d983)
}

func (__obf_721283c42292359f *ULID) SetTime(__obf_3d28b9e3414c1273 uint64) error {
	if __obf_3d28b9e3414c1273 > __obf_6bda0af3d1882741 {
		return ErrBigTime
	}

	(*__obf_721283c42292359f)[0] = byte(__obf_3d28b9e3414c1273 >> 40)
	(*__obf_721283c42292359f)[1] = byte(__obf_3d28b9e3414c1273 >> 32)
	(*__obf_721283c42292359f)[2] = byte(__obf_3d28b9e3414c1273 >> 24)
	(*__obf_721283c42292359f)[3] = byte(__obf_3d28b9e3414c1273 >> 16)
	(*__obf_721283c42292359f)[4] = byte(__obf_3d28b9e3414c1273 >> 8)
	(*__obf_721283c42292359f)[5] = byte(__obf_3d28b9e3414c1273)

	return nil
}

// Entropy returns the entropy from the ULID.
func (__obf_721283c42292359f ULID) Entropy() []byte {
	__obf_9637d6eccc44ff3f := make([]byte, 10)
	copy(__obf_9637d6eccc44ff3f, __obf_721283c42292359f[6:])
	return __obf_9637d6eccc44ff3f
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (__obf_721283c42292359f *ULID) SetEntropy(__obf_9637d6eccc44ff3f []byte) error {
	if len(__obf_9637d6eccc44ff3f) != 10 {
		return ErrDataSize
	}

	copy((*__obf_721283c42292359f)[6:], __obf_9637d6eccc44ff3f)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (__obf_721283c42292359f ULID) Compare(__obf_9dd52565732de30c ULID) int {
	return bytes.Compare(__obf_721283c42292359f[:], __obf_9dd52565732de30c[:])
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (__obf_721283c42292359f *ULID) Scan(__obf_88dec941a5f92534 any) error {
	switch __obf_2a103e2ab03dbd0a := __obf_88dec941a5f92534.(type) {
	case nil:
		return nil
	case string:
		return __obf_721283c42292359f.UnmarshalText([]byte(__obf_2a103e2ab03dbd0a))
	case []byte:
		return __obf_721283c42292359f.UnmarshalBinary(__obf_2a103e2ab03dbd0a)
	}

	return ErrScanValue
}

// Value implements the sql/driver.Valuer interface, returning the ULID as a
// slice of bytes, by invoking MarshalBinary. If your use case requires a string
// representation instead, you can create a wrapper type that calls String()
// instead.
//
//	type stringValuer ulid.ULID
//
//	func (v stringValuer) Value() (driver.Value, error) {
//	    return ulid.ULID(v).String(), nil
//	}
//
//	// Example usage.
//	db.Exec("...", stringValuer(id))
//
// All valid ULIDs, including zero-value ULIDs, return a valid Value with a nil
// error. If your use case requires zero-value ULIDs to return a non-nil error,
// you can create a wrapper type that special-cases this behavior.
//
//	var zeroValueULID ulid.ULID
//
//	type invalidZeroValuer ulid.ULID
//
//	func (v invalidZeroValuer) Value() (driver.Value, error) {
//	    if ulid.ULID(v).Compare(zeroValueULID) == 0 {
//	        return nil, fmt.Errorf("zero value")
//	    }
//	    return ulid.ULID(v).Value()
//	}
//
//	// Example usage.
//	db.Exec("...", invalidZeroValuer(id))
func (__obf_721283c42292359f ULID) Value() (driver.Value, error) {
	return __obf_721283c42292359f.MarshalBinary()
}

// Monotonic returns an entropy source that is guaranteed to yield
// strictly increasing entropy bytes for the same ULID timestamp.
// On conflicts, the previous ULID entropy is incremented with a
// random number between 1 and `inc` (inclusive).
//
// The provided entropy source must actually yield random bytes or else
// monotonic reads are not guaranteed to terminate, since there isn't
// enough randomness to compute an increment number.
//
// When `inc == 0`, it'll be set to a secure default of `math.MaxUint32`.
// The lower the value of `inc`, the easier the next ULID within the
// same millisecond is to guess. If your code depends on ULIDs having
// secure entropy bytes, then don't go under this default unless you know
// what you're doing.
//
// The returned type isn't safe for concurrent use.
func Monotonic(__obf_e867f2cef457c33d io.Reader, __obf_813f4a5e8bfdcc92 uint64) *MonotonicEntropy {
	__obf_50a171730f398461 := MonotonicEntropy{
		Reader: bufio.NewReader(__obf_e867f2cef457c33d), __obf_813f4a5e8bfdcc92: __obf_813f4a5e8bfdcc92,
	}

	if __obf_50a171730f398461.__obf_813f4a5e8bfdcc92 == 0 {
		__obf_50a171730f398461.__obf_813f4a5e8bfdcc92 = math.MaxUint32
	}

	if __obf_c2e9304dc1461d7c, __obf_a596699de7db8dcb := __obf_e867f2cef457c33d.(*rand.Rand); __obf_a596699de7db8dcb {
		__obf_50a171730f398461.__obf_c2e9304dc1461d7c = __obf_c2e9304dc1461d7c
	}

	return &__obf_50a171730f398461
}

// MonotonicEntropy is an opaque type that provides monotonic entropy.
type MonotonicEntropy struct {
	io.Reader
	__obf_3d28b9e3414c1273 uint64
	__obf_813f4a5e8bfdcc92 uint64
	__obf_e867f2cef457c33d __obf_aaee73615e1f0098

	rand                   [8]byte
	__obf_c2e9304dc1461d7c *rand.Rand
}

// MonotonicRead implements the MonotonicReader interface.
func (__obf_50a171730f398461 *MonotonicEntropy) MonotonicRead(__obf_3d28b9e3414c1273 uint64, __obf_e867f2cef457c33d []byte) (__obf_38c5d446257f8d9a error) {
	if !__obf_50a171730f398461.__obf_e867f2cef457c33d.IsZero() && __obf_50a171730f398461.__obf_3d28b9e3414c1273 == __obf_3d28b9e3414c1273 {
		__obf_38c5d446257f8d9a = __obf_50a171730f398461.__obf_d69b3128f1e6be0a()
		__obf_50a171730f398461.__obf_e867f2cef457c33d.
			AppendTo(__obf_e867f2cef457c33d)
	} else if _, __obf_38c5d446257f8d9a = io.ReadFull(__obf_50a171730f398461.Reader, __obf_e867f2cef457c33d); __obf_38c5d446257f8d9a == nil {
		__obf_50a171730f398461.__obf_3d28b9e3414c1273 = __obf_3d28b9e3414c1273
		__obf_50a171730f398461.__obf_e867f2cef457c33d.
			SetBytes(__obf_e867f2cef457c33d)
	}
	return __obf_38c5d446257f8d9a
}

// increment the previous entropy number with a random number
// of up to m.inc (inclusive).
func (__obf_50a171730f398461 *MonotonicEntropy) __obf_d69b3128f1e6be0a() error {
	if __obf_813f4a5e8bfdcc92, __obf_38c5d446257f8d9a := __obf_50a171730f398461.__obf_5d74ef60a5d6abf5(); __obf_38c5d446257f8d9a != nil {
		return __obf_38c5d446257f8d9a
	} else if __obf_50a171730f398461.__obf_e867f2cef457c33d.Add(__obf_813f4a5e8bfdcc92) {
		return ErrMonotonicOverflow
	}
	return nil
}

// random returns a uniform random value in [1, m.inc), reading entropy
// from m.Reader. When m.inc == 0 || m.inc == 1, it returns 1.
// Adapted from: https://golang.org/pkg/crypto/rand/#Int
func (__obf_50a171730f398461 *MonotonicEntropy) __obf_5d74ef60a5d6abf5() (__obf_813f4a5e8bfdcc92 uint64, __obf_38c5d446257f8d9a error) {
	if __obf_50a171730f398461.__obf_813f4a5e8bfdcc92 <= 1 {
		return 1, nil
	}

	// Fast path for using a underlying rand.Rand directly.
	if __obf_50a171730f398461.
		// Range: [1, m.inc)
		__obf_c2e9304dc1461d7c != nil {

		return 1 + uint64(__obf_50a171730f398461.__obf_c2e9304dc1461d7c.Int63n(int64(__obf_50a171730f398461.

			// bitLen is the maximum bit length needed to encode a value < m.inc.
			__obf_813f4a5e8bfdcc92))), nil
	}
	__obf_fc0cb7e67c928032 := bits.Len64(__obf_50a171730f398461.

		// byteLen is the maximum byte length needed to encode a value < m.inc.
		__obf_813f4a5e8bfdcc92)
	__obf_30ac0267f15bc9d8 := uint(__obf_fc0cb7e67c928032+7) / 8
	__obf_51ca6d228d4f2d60 := // msbitLen is the number of bits in the most significant byte of m.inc-1.
		uint(__obf_fc0cb7e67c928032 % 8)
	if __obf_51ca6d228d4f2d60 == 0 {
		__obf_51ca6d228d4f2d60 = 8
	}

	for __obf_813f4a5e8bfdcc92 == 0 || __obf_813f4a5e8bfdcc92 >= __obf_50a171730f398461.__obf_813f4a5e8bfdcc92 {
		if _, __obf_38c5d446257f8d9a = io.ReadFull(__obf_50a171730f398461.Reader, __obf_50a171730f398461.rand[:__obf_30ac0267f15bc9d8]); __obf_38c5d446257f8d9a != nil {
			return 0, __obf_38c5d446257f8d9a
		}
		__obf_50a171730f398461.

			// Clear bits in the first byte to increase the probability
			// that the candidate is < m.inc.
			rand[0] &= uint8(int(1<<__obf_51ca6d228d4f2d60) - 1)

		// Convert the read bytes into an uint64 with byteLen
		// Optimized unrolled loop.
		switch __obf_30ac0267f15bc9d8 {
		case 1:
			__obf_813f4a5e8bfdcc92 = uint64(__obf_50a171730f398461.rand[0])
		case 2:
			__obf_813f4a5e8bfdcc92 = uint64(binary.LittleEndian.Uint16(__obf_50a171730f398461.rand[:2]))
		case 3, 4:
			__obf_813f4a5e8bfdcc92 = uint64(binary.LittleEndian.Uint32(__obf_50a171730f398461.rand[:4]))
		case 5, 6, 7, 8:
			__obf_813f4a5e8bfdcc92 = uint64(binary.LittleEndian.Uint64(__obf_50a171730f398461.rand[:8]))
		}
	}

	// Range: [1, m.inc)
	return 1 + __obf_813f4a5e8bfdcc92, nil
}

type __obf_aaee73615e1f0098 struct {
	Hi uint16
	Lo uint64
}

func (__obf_51362cca9c9190a3 *__obf_aaee73615e1f0098) SetBytes(__obf_47df265e3f88b2a8 []byte) {
	__obf_51362cca9c9190a3.
		Hi = binary.BigEndian.Uint16(__obf_47df265e3f88b2a8[:2])
	__obf_51362cca9c9190a3.
		Lo = binary.BigEndian.Uint64(__obf_47df265e3f88b2a8[2:])
}

func (__obf_51362cca9c9190a3 *__obf_aaee73615e1f0098) AppendTo(__obf_47df265e3f88b2a8 []byte) {
	binary.BigEndian.PutUint16(__obf_47df265e3f88b2a8[:2], __obf_51362cca9c9190a3.Hi)
	binary.BigEndian.PutUint64(__obf_47df265e3f88b2a8[2:], __obf_51362cca9c9190a3.Lo)
}

func (__obf_51362cca9c9190a3 *__obf_aaee73615e1f0098) Add(__obf_da28677edd5fd8f1 uint64) (__obf_604245f4fc43fc14 bool) {
	__obf_586ee02affee1f48, __obf_aae6b42a22cb6ba1 := __obf_51362cca9c9190a3.Lo, __obf_51362cca9c9190a3.Hi
	if __obf_51362cca9c9190a3.Lo += __obf_da28677edd5fd8f1; __obf_51362cca9c9190a3.Lo < __obf_586ee02affee1f48 {
		__obf_51362cca9c9190a3.
			Hi++
	}
	return __obf_51362cca9c9190a3.Hi < __obf_aae6b42a22cb6ba1
}

func (__obf_51362cca9c9190a3 __obf_aaee73615e1f0098) IsZero() bool {
	return __obf_51362cca9c9190a3.Hi == 0 && __obf_51362cca9c9190a3.Lo == 0
}
