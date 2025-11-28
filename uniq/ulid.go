package __obf_e2239f4853c61591

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
	MonotonicRead(__obf_f31ce73ea1a018b5 uint64, __obf_322b9fa6feb09105 []byte) error
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
func New(__obf_f31ce73ea1a018b5 uint64, __obf_a2c87a4c4e3dcf68 io.Reader) (__obf_cee82df7c007227d ULID, __obf_bb9dc4c4f445b22f error) {
	if __obf_bb9dc4c4f445b22f = __obf_cee82df7c007227d.SetTime(__obf_f31ce73ea1a018b5); __obf_bb9dc4c4f445b22f != nil {
		return __obf_cee82df7c007227d, __obf_bb9dc4c4f445b22f
	}

	switch __obf_512e06f31ef4b08b := __obf_a2c87a4c4e3dcf68.(type) {
	case nil:
		return __obf_cee82df7c007227d, __obf_bb9dc4c4f445b22f
	case MonotonicReader:
		__obf_bb9dc4c4f445b22f = __obf_512e06f31ef4b08b.MonotonicRead(__obf_f31ce73ea1a018b5, __obf_cee82df7c007227d[6:])
	default:
		_, __obf_bb9dc4c4f445b22f = io.ReadFull(__obf_512e06f31ef4b08b, __obf_cee82df7c007227d[6:])
	}

	return __obf_cee82df7c007227d, __obf_bb9dc4c4f445b22f
}

// MustNew is a convenience function equivalent to New that panics on failure
// instead of returning an error.
func MustNew(__obf_f31ce73ea1a018b5 uint64, __obf_a2c87a4c4e3dcf68 io.Reader) ULID {
	__obf_cee82df7c007227d, __obf_bb9dc4c4f445b22f := New(__obf_f31ce73ea1a018b5, __obf_a2c87a4c4e3dcf68)
	if __obf_bb9dc4c4f445b22f != nil {
		panic(__obf_bb9dc4c4f445b22f)
	}
	return __obf_cee82df7c007227d
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(__obf_d3da1d861f7e34aa string) (__obf_cee82df7c007227d ULID, __obf_bb9dc4c4f445b22f error) {
	return __obf_cee82df7c007227d, __obf_f2ad7698498214f6([]byte(__obf_d3da1d861f7e34aa), false, &__obf_cee82df7c007227d)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(__obf_d3da1d861f7e34aa string) (__obf_cee82df7c007227d ULID, __obf_bb9dc4c4f445b22f error) {
	return __obf_cee82df7c007227d, __obf_f2ad7698498214f6([]byte(__obf_d3da1d861f7e34aa), true, &__obf_cee82df7c007227d)
}

func __obf_f2ad7698498214f6(__obf_1ab57eab018cfbd0 []byte, __obf_824894c259e2e19e bool, __obf_cee82df7c007227d *ULID) error {
	// Check if a base32 encoded ULID is the right length.
	if len(__obf_1ab57eab018cfbd0) != EncodedSize {
		return ErrDataSize
	}

	// Check if all the characters in a base32 encoded ULID are part of the
	// expected base32 character set.
	if __obf_824894c259e2e19e &&
		(__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[0]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[1]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[2]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[3]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[4]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[5]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[6]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[7]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[8]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[9]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[10]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[11]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[12]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[13]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[14]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[15]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[16]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[17]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[18]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[19]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[20]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[21]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[22]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[23]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[24]] == 0xFF ||
			__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[25]] == 0xFF) {
		return ErrInvalidCharacters
	}

	// Check if the first character in a base32 encoded ULID will overflow. This
	// happens because the base32 representation encodes 130 bits, while the
	// ULID is only 128 bits.
	//
	// See https://github.com/oklog/ulid/issues/9 for details.
	if __obf_1ab57eab018cfbd0[0] > '7' {
		return ErrOverflow
	}

	// Use an optimized unrolled loop (from https://github.com/RobThree/NUlid)
	// to decode a base32 ULID.

	// 6 bytes timestamp (48 bits)
	(*__obf_cee82df7c007227d)[0] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[0]] << 5) | __obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[1]]
	(*__obf_cee82df7c007227d)[1] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[2]] << 3) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[3]] >> 2)
	(*__obf_cee82df7c007227d)[2] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[3]] << 6) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[4]] << 1) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[5]] >> 4)
	(*__obf_cee82df7c007227d)[3] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[5]] << 4) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[6]] >> 1)
	(*__obf_cee82df7c007227d)[4] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[6]] << 7) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[7]] << 2) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[8]] >> 3)
	(*__obf_cee82df7c007227d)[5] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[8]] << 5) | __obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[9]]

	// 10 bytes of entropy (80 bits)
	(*__obf_cee82df7c007227d)[6] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[10]] << 3) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[11]] >> 2)
	(*__obf_cee82df7c007227d)[7] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[11]] << 6) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[12]] << 1) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[13]] >> 4)
	(*__obf_cee82df7c007227d)[8] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[13]] << 4) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[14]] >> 1)
	(*__obf_cee82df7c007227d)[9] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[14]] << 7) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[15]] << 2) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[16]] >> 3)
	(*__obf_cee82df7c007227d)[10] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[16]] << 5) | __obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[17]]
	(*__obf_cee82df7c007227d)[11] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[18]] << 3) | __obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[19]]>>2
	(*__obf_cee82df7c007227d)[12] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[19]] << 6) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[20]] << 1) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[21]] >> 4)
	(*__obf_cee82df7c007227d)[13] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[21]] << 4) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[22]] >> 1)
	(*__obf_cee82df7c007227d)[14] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[22]] << 7) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[23]] << 2) | (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[24]] >> 3)
	(*__obf_cee82df7c007227d)[15] = (__obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[24]] << 5) | __obf_5cd44092418da1ed[__obf_1ab57eab018cfbd0[25]]

	return nil
}

// MustParse is a convenience function equivalent to Parse that panics on failure
// instead of returning an error.
func MustParse(__obf_d3da1d861f7e34aa string) ULID {
	__obf_cee82df7c007227d, __obf_bb9dc4c4f445b22f := Parse(__obf_d3da1d861f7e34aa)
	if __obf_bb9dc4c4f445b22f != nil {
		panic(__obf_bb9dc4c4f445b22f)
	}
	return __obf_cee82df7c007227d
}

// MustParseStrict is a convenience function equivalent to ParseStrict that
// panics on failure instead of returning an error.
func MustParseStrict(__obf_d3da1d861f7e34aa string) ULID {
	__obf_cee82df7c007227d, __obf_bb9dc4c4f445b22f := ParseStrict(__obf_d3da1d861f7e34aa)
	if __obf_bb9dc4c4f445b22f != nil {
		panic(__obf_bb9dc4c4f445b22f)
	}
	return __obf_cee82df7c007227d
}

// Bytes returns bytes slice representation of ULID.
func (__obf_ac971c48d174c3f7 ULID) Bytes() []byte {
	return __obf_ac971c48d174c3f7[:]
}

// String returns a lexicographically sortable string encoded ULID
// (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy
func (__obf_cee82df7c007227d ULID) String() string {
	__obf_d3da1d861f7e34aa := make([]byte, EncodedSize)
	_ = __obf_cee82df7c007227d.MarshalTextTo(__obf_d3da1d861f7e34aa)
	return string(__obf_d3da1d861f7e34aa)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (__obf_cee82df7c007227d ULID) MarshalBinary() ([]byte, error) {
	__obf_d3da1d861f7e34aa := make([]byte, len(__obf_cee82df7c007227d))
	return __obf_d3da1d861f7e34aa, __obf_cee82df7c007227d.MarshalBinaryTo(__obf_d3da1d861f7e34aa)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (__obf_cee82df7c007227d ULID) MarshalBinaryTo(__obf_b061caccdb252a3d []byte) error {
	if len(__obf_b061caccdb252a3d) != len(__obf_cee82df7c007227d) {
		return ErrBufferSize
	}

	copy(__obf_b061caccdb252a3d, __obf_cee82df7c007227d[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (__obf_cee82df7c007227d *ULID) UnmarshalBinary(__obf_84070dac9920cf3f []byte) error {
	if len(__obf_84070dac9920cf3f) != len(*__obf_cee82df7c007227d) {
		return ErrDataSize
	}

	copy((*__obf_cee82df7c007227d)[:], __obf_84070dac9920cf3f)
	return nil
}

// Encoding is the base 32 encoding alphabet used in ULID strings.
const Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (__obf_cee82df7c007227d ULID) MarshalText() ([]byte, error) {
	__obf_d3da1d861f7e34aa := make([]byte, EncodedSize)
	return __obf_d3da1d861f7e34aa, __obf_cee82df7c007227d.MarshalTextTo(__obf_d3da1d861f7e34aa)
}

// MarshalTextTo writes the ULID as a string to the given buffer.
// ErrBufferSize is returned when the len(dst) != 26.
func (__obf_cee82df7c007227d ULID) MarshalTextTo(__obf_b061caccdb252a3d []byte) error {
	// Optimized unrolled loop ahead.
	// From https://github.com/RobThree/NUlid

	if len(__obf_b061caccdb252a3d) != EncodedSize {
		return ErrBufferSize
	}

	// 10 byte timestamp
	__obf_b061caccdb252a3d[0] = Encoding[(__obf_cee82df7c007227d[0]&224)>>5]
	__obf_b061caccdb252a3d[1] = Encoding[__obf_cee82df7c007227d[0]&31]
	__obf_b061caccdb252a3d[2] = Encoding[(__obf_cee82df7c007227d[1]&248)>>3]
	__obf_b061caccdb252a3d[3] = Encoding[((__obf_cee82df7c007227d[1]&7)<<2)|((__obf_cee82df7c007227d[2]&192)>>6)]
	__obf_b061caccdb252a3d[4] = Encoding[(__obf_cee82df7c007227d[2]&62)>>1]
	__obf_b061caccdb252a3d[5] = Encoding[((__obf_cee82df7c007227d[2]&1)<<4)|((__obf_cee82df7c007227d[3]&240)>>4)]
	__obf_b061caccdb252a3d[6] = Encoding[((__obf_cee82df7c007227d[3]&15)<<1)|((__obf_cee82df7c007227d[4]&128)>>7)]
	__obf_b061caccdb252a3d[7] = Encoding[(__obf_cee82df7c007227d[4]&124)>>2]
	__obf_b061caccdb252a3d[8] = Encoding[((__obf_cee82df7c007227d[4]&3)<<3)|((__obf_cee82df7c007227d[5]&224)>>5)]
	__obf_b061caccdb252a3d[9] = Encoding[__obf_cee82df7c007227d[5]&31]

	// 16 bytes of entropy
	__obf_b061caccdb252a3d[10] = Encoding[(__obf_cee82df7c007227d[6]&248)>>3]
	__obf_b061caccdb252a3d[11] = Encoding[((__obf_cee82df7c007227d[6]&7)<<2)|((__obf_cee82df7c007227d[7]&192)>>6)]
	__obf_b061caccdb252a3d[12] = Encoding[(__obf_cee82df7c007227d[7]&62)>>1]
	__obf_b061caccdb252a3d[13] = Encoding[((__obf_cee82df7c007227d[7]&1)<<4)|((__obf_cee82df7c007227d[8]&240)>>4)]
	__obf_b061caccdb252a3d[14] = Encoding[((__obf_cee82df7c007227d[8]&15)<<1)|((__obf_cee82df7c007227d[9]&128)>>7)]
	__obf_b061caccdb252a3d[15] = Encoding[(__obf_cee82df7c007227d[9]&124)>>2]
	__obf_b061caccdb252a3d[16] = Encoding[((__obf_cee82df7c007227d[9]&3)<<3)|((__obf_cee82df7c007227d[10]&224)>>5)]
	__obf_b061caccdb252a3d[17] = Encoding[__obf_cee82df7c007227d[10]&31]
	__obf_b061caccdb252a3d[18] = Encoding[(__obf_cee82df7c007227d[11]&248)>>3]
	__obf_b061caccdb252a3d[19] = Encoding[((__obf_cee82df7c007227d[11]&7)<<2)|((__obf_cee82df7c007227d[12]&192)>>6)]
	__obf_b061caccdb252a3d[20] = Encoding[(__obf_cee82df7c007227d[12]&62)>>1]
	__obf_b061caccdb252a3d[21] = Encoding[((__obf_cee82df7c007227d[12]&1)<<4)|((__obf_cee82df7c007227d[13]&240)>>4)]
	__obf_b061caccdb252a3d[22] = Encoding[((__obf_cee82df7c007227d[13]&15)<<1)|((__obf_cee82df7c007227d[14]&128)>>7)]
	__obf_b061caccdb252a3d[23] = Encoding[(__obf_cee82df7c007227d[14]&124)>>2]
	__obf_b061caccdb252a3d[24] = Encoding[((__obf_cee82df7c007227d[14]&3)<<3)|((__obf_cee82df7c007227d[15]&224)>>5)]
	__obf_b061caccdb252a3d[25] = Encoding[__obf_cee82df7c007227d[15]&31]

	return nil
}

// Byte to index table for O(1) lookups when unmarshaling.
// We use 0xFF as sentinel value for invalid indexes.
var __obf_5cd44092418da1ed = [...]byte{
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
func (__obf_cee82df7c007227d *ULID) UnmarshalText(__obf_1ab57eab018cfbd0 []byte) error {
	return __obf_f2ad7698498214f6(__obf_1ab57eab018cfbd0, false, __obf_cee82df7c007227d)
}

// Time returns the Unix time in milliseconds encoded in the ULID.
// Use the top level Time function to convert the returned value to
// a time.Time.
func (__obf_cee82df7c007227d ULID) Time() uint64 {
	return uint64(__obf_cee82df7c007227d[5]) | uint64(__obf_cee82df7c007227d[4])<<8 |
		uint64(__obf_cee82df7c007227d[3])<<16 | uint64(__obf_cee82df7c007227d[2])<<24 |
		uint64(__obf_cee82df7c007227d[1])<<32 | uint64(__obf_cee82df7c007227d[0])<<40
}

// maxTime is the maximum Unix time in milliseconds that can be
// represented in an ULID.
var __obf_5c4134d4f0cf2787 = ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}.Time()

// MaxTime returns the maximum Unix time in milliseconds that
// can be encoded in an ULID.
func MaxTime() uint64 { return __obf_5c4134d4f0cf2787 }

// Now is a convenience function that returns the current
// UTC time in Unix milliseconds. Equivalent to:
//
//	Timestamp(time.Now().UTC())
func Now() uint64 { return Timestamp(time.Now().UTC()) }

// Timestamp converts a time.Time to Unix milliseconds.
//
// Because of the way ULID stores time, times from the year
// 10889 produces undefined results.
func Timestamp(__obf_de192fc01291fe85 time.Time) uint64 {
	return uint64(__obf_de192fc01291fe85.Unix())*1000 +
		uint64(__obf_de192fc01291fe85.Nanosecond()/int(time.Millisecond))
}

// Time converts Unix milliseconds in the format
// returned by the Timestamp function to a time.Time.
func Time(__obf_f31ce73ea1a018b5 uint64) time.Time {
	__obf_922f86e6d1baffa6 := int64(__obf_f31ce73ea1a018b5 / 1e3)
	__obf_1511395b4c64bfad := int64((__obf_f31ce73ea1a018b5 % 1e3) * 1e6)
	return time.Unix(__obf_922f86e6d1baffa6, __obf_1511395b4c64bfad)
}

// SetTime sets the time component of the ULID to the given Unix time
// in milliseconds.
func (__obf_cee82df7c007227d *ULID) SetTime(__obf_f31ce73ea1a018b5 uint64) error {
	if __obf_f31ce73ea1a018b5 > __obf_5c4134d4f0cf2787 {
		return ErrBigTime
	}

	(*__obf_cee82df7c007227d)[0] = byte(__obf_f31ce73ea1a018b5 >> 40)
	(*__obf_cee82df7c007227d)[1] = byte(__obf_f31ce73ea1a018b5 >> 32)
	(*__obf_cee82df7c007227d)[2] = byte(__obf_f31ce73ea1a018b5 >> 24)
	(*__obf_cee82df7c007227d)[3] = byte(__obf_f31ce73ea1a018b5 >> 16)
	(*__obf_cee82df7c007227d)[4] = byte(__obf_f31ce73ea1a018b5 >> 8)
	(*__obf_cee82df7c007227d)[5] = byte(__obf_f31ce73ea1a018b5)

	return nil
}

// Entropy returns the entropy from the ULID.
func (__obf_cee82df7c007227d ULID) Entropy() []byte {
	__obf_512e06f31ef4b08b := make([]byte, 10)
	copy(__obf_512e06f31ef4b08b, __obf_cee82df7c007227d[6:])
	return __obf_512e06f31ef4b08b
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (__obf_cee82df7c007227d *ULID) SetEntropy(__obf_512e06f31ef4b08b []byte) error {
	if len(__obf_512e06f31ef4b08b) != 10 {
		return ErrDataSize
	}

	copy((*__obf_cee82df7c007227d)[6:], __obf_512e06f31ef4b08b)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (__obf_cee82df7c007227d ULID) Compare(__obf_0d049dc4f010fae3 ULID) int {
	return bytes.Compare(__obf_cee82df7c007227d[:], __obf_0d049dc4f010fae3[:])
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (__obf_cee82df7c007227d *ULID) Scan(__obf_0aac115271097993 any) error {
	switch __obf_87fbd978922103e3 := __obf_0aac115271097993.(type) {
	case nil:
		return nil
	case string:
		return __obf_cee82df7c007227d.UnmarshalText([]byte(__obf_87fbd978922103e3))
	case []byte:
		return __obf_cee82df7c007227d.UnmarshalBinary(__obf_87fbd978922103e3)
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
func (__obf_cee82df7c007227d ULID) Value() (driver.Value, error) {
	return __obf_cee82df7c007227d.MarshalBinary()
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
func Monotonic(__obf_a2c87a4c4e3dcf68 io.Reader, __obf_4aa59cae06bddf01 uint64) *MonotonicEntropy {
	__obf_dcb140ccc793bff9 := MonotonicEntropy{
		Reader:                 bufio.NewReader(__obf_a2c87a4c4e3dcf68),
		__obf_4aa59cae06bddf01: __obf_4aa59cae06bddf01,
	}

	if __obf_dcb140ccc793bff9.__obf_4aa59cae06bddf01 == 0 {
		__obf_dcb140ccc793bff9.__obf_4aa59cae06bddf01 = math.MaxUint32
	}

	if __obf_6547fdd8df243eca, __obf_c8f6ef41918857c4 := __obf_a2c87a4c4e3dcf68.(*rand.Rand); __obf_c8f6ef41918857c4 {
		__obf_dcb140ccc793bff9.__obf_6547fdd8df243eca = __obf_6547fdd8df243eca
	}

	return &__obf_dcb140ccc793bff9
}

// MonotonicEntropy is an opaque type that provides monotonic entropy.
type MonotonicEntropy struct {
	io.Reader
	__obf_f31ce73ea1a018b5 uint64
	__obf_4aa59cae06bddf01 uint64
	__obf_a2c87a4c4e3dcf68 __obf_94c0f73d7bb212b7
	rand                   [8]byte
	__obf_6547fdd8df243eca *rand.Rand
}

// MonotonicRead implements the MonotonicReader interface.
func (__obf_dcb140ccc793bff9 *MonotonicEntropy) MonotonicRead(__obf_f31ce73ea1a018b5 uint64, __obf_a2c87a4c4e3dcf68 []byte) (__obf_bb9dc4c4f445b22f error) {
	if !__obf_dcb140ccc793bff9.__obf_a2c87a4c4e3dcf68.IsZero() && __obf_dcb140ccc793bff9.__obf_f31ce73ea1a018b5 == __obf_f31ce73ea1a018b5 {
		__obf_bb9dc4c4f445b22f = __obf_dcb140ccc793bff9.__obf_29a4c8e9e2c99b49()
		__obf_dcb140ccc793bff9.__obf_a2c87a4c4e3dcf68.AppendTo(__obf_a2c87a4c4e3dcf68)
	} else if _, __obf_bb9dc4c4f445b22f = io.ReadFull(__obf_dcb140ccc793bff9.Reader, __obf_a2c87a4c4e3dcf68); __obf_bb9dc4c4f445b22f == nil {
		__obf_dcb140ccc793bff9.__obf_f31ce73ea1a018b5 = __obf_f31ce73ea1a018b5
		__obf_dcb140ccc793bff9.__obf_a2c87a4c4e3dcf68.SetBytes(__obf_a2c87a4c4e3dcf68)
	}
	return __obf_bb9dc4c4f445b22f
}

// increment the previous entropy number with a random number
// of up to m.inc (inclusive).
func (__obf_dcb140ccc793bff9 *MonotonicEntropy) __obf_29a4c8e9e2c99b49() error {
	if __obf_4aa59cae06bddf01, __obf_bb9dc4c4f445b22f := __obf_dcb140ccc793bff9.__obf_d2a1b1c26c3e11a4(); __obf_bb9dc4c4f445b22f != nil {
		return __obf_bb9dc4c4f445b22f
	} else if __obf_dcb140ccc793bff9.__obf_a2c87a4c4e3dcf68.Add(__obf_4aa59cae06bddf01) {
		return ErrMonotonicOverflow
	}
	return nil
}

// random returns a uniform random value in [1, m.inc), reading entropy
// from m.Reader. When m.inc == 0 || m.inc == 1, it returns 1.
// Adapted from: https://golang.org/pkg/crypto/rand/#Int
func (__obf_dcb140ccc793bff9 *MonotonicEntropy) __obf_d2a1b1c26c3e11a4() (__obf_4aa59cae06bddf01 uint64, __obf_bb9dc4c4f445b22f error) {
	if __obf_dcb140ccc793bff9.__obf_4aa59cae06bddf01 <= 1 {
		return 1, nil
	}

	// Fast path for using a underlying rand.Rand directly.
	if __obf_dcb140ccc793bff9.__obf_6547fdd8df243eca != nil {
		// Range: [1, m.inc)
		return 1 + uint64(__obf_dcb140ccc793bff9.__obf_6547fdd8df243eca.Int63n(int64(__obf_dcb140ccc793bff9.__obf_4aa59cae06bddf01))), nil
	}

	// bitLen is the maximum bit length needed to encode a value < m.inc.
	__obf_9ac09b67d5da5690 := bits.Len64(__obf_dcb140ccc793bff9.__obf_4aa59cae06bddf01)

	// byteLen is the maximum byte length needed to encode a value < m.inc.
	__obf_2f10e4f847e403a4 := uint(__obf_9ac09b67d5da5690+7) / 8

	// msbitLen is the number of bits in the most significant byte of m.inc-1.
	__obf_f3f95609d6a06706 := uint(__obf_9ac09b67d5da5690 % 8)
	if __obf_f3f95609d6a06706 == 0 {
		__obf_f3f95609d6a06706 = 8
	}

	for __obf_4aa59cae06bddf01 == 0 || __obf_4aa59cae06bddf01 >= __obf_dcb140ccc793bff9.__obf_4aa59cae06bddf01 {
		if _, __obf_bb9dc4c4f445b22f = io.ReadFull(__obf_dcb140ccc793bff9.Reader, __obf_dcb140ccc793bff9.rand[:__obf_2f10e4f847e403a4]); __obf_bb9dc4c4f445b22f != nil {
			return 0, __obf_bb9dc4c4f445b22f
		}

		// Clear bits in the first byte to increase the probability
		// that the candidate is < m.inc.
		__obf_dcb140ccc793bff9.rand[0] &= uint8(int(1<<__obf_f3f95609d6a06706) - 1)

		// Convert the read bytes into an uint64 with byteLen
		// Optimized unrolled loop.
		switch __obf_2f10e4f847e403a4 {
		case 1:
			__obf_4aa59cae06bddf01 = uint64(__obf_dcb140ccc793bff9.rand[0])
		case 2:
			__obf_4aa59cae06bddf01 = uint64(binary.LittleEndian.Uint16(__obf_dcb140ccc793bff9.rand[:2]))
		case 3, 4:
			__obf_4aa59cae06bddf01 = uint64(binary.LittleEndian.Uint32(__obf_dcb140ccc793bff9.rand[:4]))
		case 5, 6, 7, 8:
			__obf_4aa59cae06bddf01 = uint64(binary.LittleEndian.Uint64(__obf_dcb140ccc793bff9.rand[:8]))
		}
	}

	// Range: [1, m.inc)
	return 1 + __obf_4aa59cae06bddf01, nil
}

type __obf_94c0f73d7bb212b7 struct {
	Hi uint16
	Lo uint64
}

func (__obf_ac971c48d174c3f7 *__obf_94c0f73d7bb212b7) SetBytes(__obf_8fb62a7e38414372 []byte) {
	__obf_ac971c48d174c3f7.Hi = binary.BigEndian.Uint16(__obf_8fb62a7e38414372[:2])
	__obf_ac971c48d174c3f7.Lo = binary.BigEndian.Uint64(__obf_8fb62a7e38414372[2:])
}

func (__obf_ac971c48d174c3f7 *__obf_94c0f73d7bb212b7) AppendTo(__obf_8fb62a7e38414372 []byte) {
	binary.BigEndian.PutUint16(__obf_8fb62a7e38414372[:2], __obf_ac971c48d174c3f7.Hi)
	binary.BigEndian.PutUint64(__obf_8fb62a7e38414372[2:], __obf_ac971c48d174c3f7.Lo)
}

func (__obf_ac971c48d174c3f7 *__obf_94c0f73d7bb212b7) Add(__obf_1758a8beead18d7e uint64) (__obf_03b137b0e634a154 bool) {
	__obf_43b830a21b264544, __obf_3200341b6550b128 := __obf_ac971c48d174c3f7.Lo, __obf_ac971c48d174c3f7.Hi
	if __obf_ac971c48d174c3f7.Lo += __obf_1758a8beead18d7e; __obf_ac971c48d174c3f7.Lo < __obf_43b830a21b264544 {
		__obf_ac971c48d174c3f7.Hi++
	}
	return __obf_ac971c48d174c3f7.Hi < __obf_3200341b6550b128
}

func (__obf_ac971c48d174c3f7 __obf_94c0f73d7bb212b7) IsZero() bool {
	return __obf_ac971c48d174c3f7.Hi == 0 && __obf_ac971c48d174c3f7.Lo == 0
}
