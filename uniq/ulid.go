package __obf_417508f5ae215d0a

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
	MonotonicRead(__obf_443a50eb68e39a29 uint64, __obf_23d8d7aa7bc9ec29 []byte) error
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
func New(__obf_443a50eb68e39a29 uint64, __obf_b43269b69bf1ec16 io.Reader) (__obf_48218c2a1a645d91 ULID, __obf_b99f62eaceb52d40 error) {
	if __obf_b99f62eaceb52d40 = __obf_48218c2a1a645d91.SetTime(__obf_443a50eb68e39a29); __obf_b99f62eaceb52d40 != nil {
		return __obf_48218c2a1a645d91, __obf_b99f62eaceb52d40
	}

	switch __obf_fee66209667b9075 := __obf_b43269b69bf1ec16.(type) {
	case nil:
		return __obf_48218c2a1a645d91, __obf_b99f62eaceb52d40
	case MonotonicReader:
		__obf_b99f62eaceb52d40 = __obf_fee66209667b9075.MonotonicRead(__obf_443a50eb68e39a29, __obf_48218c2a1a645d91[6:])
	default:
		_, __obf_b99f62eaceb52d40 = io.ReadFull(__obf_fee66209667b9075, __obf_48218c2a1a645d91[6:])
	}

	return __obf_48218c2a1a645d91, __obf_b99f62eaceb52d40
}

// MustNew is a convenience function equivalent to New that panics on failure
// instead of returning an error.
func MustNew(__obf_443a50eb68e39a29 uint64, __obf_b43269b69bf1ec16 io.Reader) ULID {
	__obf_48218c2a1a645d91, __obf_b99f62eaceb52d40 := New(__obf_443a50eb68e39a29, __obf_b43269b69bf1ec16)
	if __obf_b99f62eaceb52d40 != nil {
		panic(__obf_b99f62eaceb52d40)
	}
	return __obf_48218c2a1a645d91
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(__obf_5d0699c9c5840b8a string) (__obf_48218c2a1a645d91 ULID, __obf_b99f62eaceb52d40 error) {
	return __obf_48218c2a1a645d91, __obf_712654f03aa64404([]byte(__obf_5d0699c9c5840b8a), false, &__obf_48218c2a1a645d91)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(__obf_5d0699c9c5840b8a string) (__obf_48218c2a1a645d91 ULID, __obf_b99f62eaceb52d40 error) {
	return __obf_48218c2a1a645d91, __obf_712654f03aa64404([]byte(__obf_5d0699c9c5840b8a), true, &__obf_48218c2a1a645d91)
}

func __obf_712654f03aa64404(__obf_a9527a68ae1dfba0 []byte, __obf_46e8e5ddecc53788 bool, __obf_48218c2a1a645d91 *ULID) error {
	// Check if a base32 encoded ULID is the right length.
	if len(__obf_a9527a68ae1dfba0) != EncodedSize {
		return ErrDataSize
	}

	// Check if all the characters in a base32 encoded ULID are part of the
	// expected base32 character set.
	if __obf_46e8e5ddecc53788 &&
		(__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[0]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[1]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[2]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[3]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[4]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[5]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[6]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[7]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[8]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[9]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[10]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[11]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[12]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[13]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[14]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[15]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[16]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[17]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[18]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[19]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[20]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[21]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[22]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[23]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[24]] == 0xFF ||
			__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[25]] == 0xFF) {
		return ErrInvalidCharacters
	}

	// Check if the first character in a base32 encoded ULID will overflow. This
	// happens because the base32 representation encodes 130 bits, while the
	// ULID is only 128 bits.
	//
	// See https://github.com/oklog/ulid/issues/9 for details.
	if __obf_a9527a68ae1dfba0[0] > '7' {
		return ErrOverflow
	}

	// Use an optimized unrolled loop (from https://github.com/RobThree/NUlid)
	// to decode a base32 ULID.

	// 6 bytes timestamp (48 bits)
	(*__obf_48218c2a1a645d91)[0] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[0]] << 5) | __obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[1]]
	(*__obf_48218c2a1a645d91)[1] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[2]] << 3) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[3]] >> 2)
	(*__obf_48218c2a1a645d91)[2] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[3]] << 6) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[4]] << 1) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[5]] >> 4)
	(*__obf_48218c2a1a645d91)[3] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[5]] << 4) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[6]] >> 1)
	(*__obf_48218c2a1a645d91)[4] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[6]] << 7) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[7]] << 2) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[8]] >> 3)
	(*__obf_48218c2a1a645d91)[5] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[8]] << 5) | __obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[9]]

	// 10 bytes of entropy (80 bits)
	(*__obf_48218c2a1a645d91)[6] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[10]] << 3) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[11]] >> 2)
	(*__obf_48218c2a1a645d91)[7] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[11]] << 6) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[12]] << 1) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[13]] >> 4)
	(*__obf_48218c2a1a645d91)[8] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[13]] << 4) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[14]] >> 1)
	(*__obf_48218c2a1a645d91)[9] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[14]] << 7) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[15]] << 2) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[16]] >> 3)
	(*__obf_48218c2a1a645d91)[10] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[16]] << 5) | __obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[17]]
	(*__obf_48218c2a1a645d91)[11] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[18]] << 3) | __obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[19]]>>2
	(*__obf_48218c2a1a645d91)[12] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[19]] << 6) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[20]] << 1) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[21]] >> 4)
	(*__obf_48218c2a1a645d91)[13] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[21]] << 4) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[22]] >> 1)
	(*__obf_48218c2a1a645d91)[14] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[22]] << 7) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[23]] << 2) | (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[24]] >> 3)
	(*__obf_48218c2a1a645d91)[15] = (__obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[24]] << 5) | __obf_0b2b68c9f1500f60[__obf_a9527a68ae1dfba0[25]]

	return nil
}

// MustParse is a convenience function equivalent to Parse that panics on failure
// instead of returning an error.
func MustParse(__obf_5d0699c9c5840b8a string) ULID {
	__obf_48218c2a1a645d91, __obf_b99f62eaceb52d40 := Parse(__obf_5d0699c9c5840b8a)
	if __obf_b99f62eaceb52d40 != nil {
		panic(__obf_b99f62eaceb52d40)
	}
	return __obf_48218c2a1a645d91
}

// MustParseStrict is a convenience function equivalent to ParseStrict that
// panics on failure instead of returning an error.
func MustParseStrict(__obf_5d0699c9c5840b8a string) ULID {
	__obf_48218c2a1a645d91, __obf_b99f62eaceb52d40 := ParseStrict(__obf_5d0699c9c5840b8a)
	if __obf_b99f62eaceb52d40 != nil {
		panic(__obf_b99f62eaceb52d40)
	}
	return __obf_48218c2a1a645d91
}

// Bytes returns bytes slice representation of ULID.
func (__obf_feb8e5a89f005fb4 ULID) Bytes() []byte {
	return __obf_feb8e5a89f005fb4[:]
}

// String returns a lexicographically sortable string encoded ULID
// (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy
func (__obf_48218c2a1a645d91 ULID) String() string {
	__obf_5d0699c9c5840b8a := make([]byte, EncodedSize)
	_ = __obf_48218c2a1a645d91.MarshalTextTo(__obf_5d0699c9c5840b8a)
	return string(__obf_5d0699c9c5840b8a)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (__obf_48218c2a1a645d91 ULID) MarshalBinary() ([]byte, error) {
	__obf_5d0699c9c5840b8a := make([]byte, len(__obf_48218c2a1a645d91))
	return __obf_5d0699c9c5840b8a, __obf_48218c2a1a645d91.MarshalBinaryTo(__obf_5d0699c9c5840b8a)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (__obf_48218c2a1a645d91 ULID) MarshalBinaryTo(__obf_c7745a97c2612fba []byte) error {
	if len(__obf_c7745a97c2612fba) != len(__obf_48218c2a1a645d91) {
		return ErrBufferSize
	}

	copy(__obf_c7745a97c2612fba, __obf_48218c2a1a645d91[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (__obf_48218c2a1a645d91 *ULID) UnmarshalBinary(__obf_b7ac0354993a604d []byte) error {
	if len(__obf_b7ac0354993a604d) != len(*__obf_48218c2a1a645d91) {
		return ErrDataSize
	}

	copy((*__obf_48218c2a1a645d91)[:], __obf_b7ac0354993a604d)
	return nil
}

// Encoding is the base 32 encoding alphabet used in ULID strings.
const Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (__obf_48218c2a1a645d91 ULID) MarshalText() ([]byte, error) {
	__obf_5d0699c9c5840b8a := make([]byte, EncodedSize)
	return __obf_5d0699c9c5840b8a, __obf_48218c2a1a645d91.MarshalTextTo(__obf_5d0699c9c5840b8a)
}

// MarshalTextTo writes the ULID as a string to the given buffer.
// ErrBufferSize is returned when the len(dst) != 26.
func (__obf_48218c2a1a645d91 ULID) MarshalTextTo(__obf_c7745a97c2612fba []byte) error {
	// Optimized unrolled loop ahead.
	// From https://github.com/RobThree/NUlid

	if len(__obf_c7745a97c2612fba) != EncodedSize {
		return ErrBufferSize
	}

	// 10 byte timestamp
	__obf_c7745a97c2612fba[0] = Encoding[(__obf_48218c2a1a645d91[0]&224)>>5]
	__obf_c7745a97c2612fba[1] = Encoding[__obf_48218c2a1a645d91[0]&31]
	__obf_c7745a97c2612fba[2] = Encoding[(__obf_48218c2a1a645d91[1]&248)>>3]
	__obf_c7745a97c2612fba[3] = Encoding[((__obf_48218c2a1a645d91[1]&7)<<2)|((__obf_48218c2a1a645d91[2]&192)>>6)]
	__obf_c7745a97c2612fba[4] = Encoding[(__obf_48218c2a1a645d91[2]&62)>>1]
	__obf_c7745a97c2612fba[5] = Encoding[((__obf_48218c2a1a645d91[2]&1)<<4)|((__obf_48218c2a1a645d91[3]&240)>>4)]
	__obf_c7745a97c2612fba[6] = Encoding[((__obf_48218c2a1a645d91[3]&15)<<1)|((__obf_48218c2a1a645d91[4]&128)>>7)]
	__obf_c7745a97c2612fba[7] = Encoding[(__obf_48218c2a1a645d91[4]&124)>>2]
	__obf_c7745a97c2612fba[8] = Encoding[((__obf_48218c2a1a645d91[4]&3)<<3)|((__obf_48218c2a1a645d91[5]&224)>>5)]
	__obf_c7745a97c2612fba[9] = Encoding[__obf_48218c2a1a645d91[5]&31]

	// 16 bytes of entropy
	__obf_c7745a97c2612fba[10] = Encoding[(__obf_48218c2a1a645d91[6]&248)>>3]
	__obf_c7745a97c2612fba[11] = Encoding[((__obf_48218c2a1a645d91[6]&7)<<2)|((__obf_48218c2a1a645d91[7]&192)>>6)]
	__obf_c7745a97c2612fba[12] = Encoding[(__obf_48218c2a1a645d91[7]&62)>>1]
	__obf_c7745a97c2612fba[13] = Encoding[((__obf_48218c2a1a645d91[7]&1)<<4)|((__obf_48218c2a1a645d91[8]&240)>>4)]
	__obf_c7745a97c2612fba[14] = Encoding[((__obf_48218c2a1a645d91[8]&15)<<1)|((__obf_48218c2a1a645d91[9]&128)>>7)]
	__obf_c7745a97c2612fba[15] = Encoding[(__obf_48218c2a1a645d91[9]&124)>>2]
	__obf_c7745a97c2612fba[16] = Encoding[((__obf_48218c2a1a645d91[9]&3)<<3)|((__obf_48218c2a1a645d91[10]&224)>>5)]
	__obf_c7745a97c2612fba[17] = Encoding[__obf_48218c2a1a645d91[10]&31]
	__obf_c7745a97c2612fba[18] = Encoding[(__obf_48218c2a1a645d91[11]&248)>>3]
	__obf_c7745a97c2612fba[19] = Encoding[((__obf_48218c2a1a645d91[11]&7)<<2)|((__obf_48218c2a1a645d91[12]&192)>>6)]
	__obf_c7745a97c2612fba[20] = Encoding[(__obf_48218c2a1a645d91[12]&62)>>1]
	__obf_c7745a97c2612fba[21] = Encoding[((__obf_48218c2a1a645d91[12]&1)<<4)|((__obf_48218c2a1a645d91[13]&240)>>4)]
	__obf_c7745a97c2612fba[22] = Encoding[((__obf_48218c2a1a645d91[13]&15)<<1)|((__obf_48218c2a1a645d91[14]&128)>>7)]
	__obf_c7745a97c2612fba[23] = Encoding[(__obf_48218c2a1a645d91[14]&124)>>2]
	__obf_c7745a97c2612fba[24] = Encoding[((__obf_48218c2a1a645d91[14]&3)<<3)|((__obf_48218c2a1a645d91[15]&224)>>5)]
	__obf_c7745a97c2612fba[25] = Encoding[__obf_48218c2a1a645d91[15]&31]

	return nil
}

// Byte to index table for O(1) lookups when unmarshaling.
// We use 0xFF as sentinel value for invalid indexes.
var __obf_0b2b68c9f1500f60 = [...]byte{
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
func (__obf_48218c2a1a645d91 *ULID) UnmarshalText(__obf_a9527a68ae1dfba0 []byte) error {
	return __obf_712654f03aa64404(__obf_a9527a68ae1dfba0, false, __obf_48218c2a1a645d91)
}

// Time returns the Unix time in milliseconds encoded in the ULID.
// Use the top level Time function to convert the returned value to
// a time.Time.
func (__obf_48218c2a1a645d91 ULID) Time() uint64 {
	return uint64(__obf_48218c2a1a645d91[5]) | uint64(__obf_48218c2a1a645d91[4])<<8 |
		uint64(__obf_48218c2a1a645d91[3])<<16 | uint64(__obf_48218c2a1a645d91[2])<<24 |
		uint64(__obf_48218c2a1a645d91[1])<<32 | uint64(__obf_48218c2a1a645d91[0])<<40
}

// maxTime is the maximum Unix time in milliseconds that can be
// represented in an ULID.
var __obf_cbe820ab04ac2dee = ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}.Time()

// MaxTime returns the maximum Unix time in milliseconds that
// can be encoded in an ULID.
func MaxTime() uint64 { return __obf_cbe820ab04ac2dee }

// Now is a convenience function that returns the current
// UTC time in Unix milliseconds. Equivalent to:
//
//	Timestamp(time.Now().UTC())
func Now() uint64 { return Timestamp(time.Now().UTC()) }

// Timestamp converts a time.Time to Unix milliseconds.
//
// Because of the way ULID stores time, times from the year
// 10889 produces undefined results.
func Timestamp(__obf_59b4996dba32a5cd time.Time) uint64 {
	return uint64(__obf_59b4996dba32a5cd.Unix())*1000 +
		uint64(__obf_59b4996dba32a5cd.Nanosecond()/int(time.Millisecond))
}

// Time converts Unix milliseconds in the format
// returned by the Timestamp function to a time.Time.
func Time(__obf_443a50eb68e39a29 uint64) time.Time {
	__obf_eecb45230a7a2ebb := int64(__obf_443a50eb68e39a29 / 1e3)
	__obf_186a9aa7b7c0f043 := int64((__obf_443a50eb68e39a29 % 1e3) * 1e6)
	return time.Unix(__obf_eecb45230a7a2ebb, __obf_186a9aa7b7c0f043)
}

// SetTime sets the time component of the ULID to the given Unix time
// in milliseconds.
func (__obf_48218c2a1a645d91 *ULID) SetTime(__obf_443a50eb68e39a29 uint64) error {
	if __obf_443a50eb68e39a29 > __obf_cbe820ab04ac2dee {
		return ErrBigTime
	}

	(*__obf_48218c2a1a645d91)[0] = byte(__obf_443a50eb68e39a29 >> 40)
	(*__obf_48218c2a1a645d91)[1] = byte(__obf_443a50eb68e39a29 >> 32)
	(*__obf_48218c2a1a645d91)[2] = byte(__obf_443a50eb68e39a29 >> 24)
	(*__obf_48218c2a1a645d91)[3] = byte(__obf_443a50eb68e39a29 >> 16)
	(*__obf_48218c2a1a645d91)[4] = byte(__obf_443a50eb68e39a29 >> 8)
	(*__obf_48218c2a1a645d91)[5] = byte(__obf_443a50eb68e39a29)

	return nil
}

// Entropy returns the entropy from the ULID.
func (__obf_48218c2a1a645d91 ULID) Entropy() []byte {
	__obf_fee66209667b9075 := make([]byte, 10)
	copy(__obf_fee66209667b9075, __obf_48218c2a1a645d91[6:])
	return __obf_fee66209667b9075
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (__obf_48218c2a1a645d91 *ULID) SetEntropy(__obf_fee66209667b9075 []byte) error {
	if len(__obf_fee66209667b9075) != 10 {
		return ErrDataSize
	}

	copy((*__obf_48218c2a1a645d91)[6:], __obf_fee66209667b9075)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (__obf_48218c2a1a645d91 ULID) Compare(__obf_cff0bd4ca6cf2a25 ULID) int {
	return bytes.Compare(__obf_48218c2a1a645d91[:], __obf_cff0bd4ca6cf2a25[:])
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (__obf_48218c2a1a645d91 *ULID) Scan(__obf_379586b936fdd04e any) error {
	switch __obf_46a933db83592c20 := __obf_379586b936fdd04e.(type) {
	case nil:
		return nil
	case string:
		return __obf_48218c2a1a645d91.UnmarshalText([]byte(__obf_46a933db83592c20))
	case []byte:
		return __obf_48218c2a1a645d91.UnmarshalBinary(__obf_46a933db83592c20)
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
func (__obf_48218c2a1a645d91 ULID) Value() (driver.Value, error) {
	return __obf_48218c2a1a645d91.MarshalBinary()
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
func Monotonic(__obf_b43269b69bf1ec16 io.Reader, __obf_534044727099f444 uint64) *MonotonicEntropy {
	__obf_0db4e7f4c3797edd := MonotonicEntropy{
		Reader:                 bufio.NewReader(__obf_b43269b69bf1ec16),
		__obf_534044727099f444: __obf_534044727099f444,
	}

	if __obf_0db4e7f4c3797edd.__obf_534044727099f444 == 0 {
		__obf_0db4e7f4c3797edd.__obf_534044727099f444 = math.MaxUint32
	}

	if __obf_0a92aebefacfe4a3, __obf_0bd9d9196038949b := __obf_b43269b69bf1ec16.(*rand.Rand); __obf_0bd9d9196038949b {
		__obf_0db4e7f4c3797edd.__obf_0a92aebefacfe4a3 = __obf_0a92aebefacfe4a3
	}

	return &__obf_0db4e7f4c3797edd
}

// MonotonicEntropy is an opaque type that provides monotonic entropy.
type MonotonicEntropy struct {
	io.Reader
	__obf_443a50eb68e39a29 uint64
	__obf_534044727099f444 uint64
	__obf_b43269b69bf1ec16 __obf_87b078d1bfc5f077
	rand                   [8]byte
	__obf_0a92aebefacfe4a3 *rand.Rand
}

// MonotonicRead implements the MonotonicReader interface.
func (__obf_0db4e7f4c3797edd *MonotonicEntropy) MonotonicRead(__obf_443a50eb68e39a29 uint64, __obf_b43269b69bf1ec16 []byte) (__obf_b99f62eaceb52d40 error) {
	if !__obf_0db4e7f4c3797edd.__obf_b43269b69bf1ec16.IsZero() && __obf_0db4e7f4c3797edd.__obf_443a50eb68e39a29 == __obf_443a50eb68e39a29 {
		__obf_b99f62eaceb52d40 = __obf_0db4e7f4c3797edd.__obf_510b83a3e924d48e()
		__obf_0db4e7f4c3797edd.__obf_b43269b69bf1ec16.AppendTo(__obf_b43269b69bf1ec16)
	} else if _, __obf_b99f62eaceb52d40 = io.ReadFull(__obf_0db4e7f4c3797edd.Reader, __obf_b43269b69bf1ec16); __obf_b99f62eaceb52d40 == nil {
		__obf_0db4e7f4c3797edd.__obf_443a50eb68e39a29 = __obf_443a50eb68e39a29
		__obf_0db4e7f4c3797edd.__obf_b43269b69bf1ec16.SetBytes(__obf_b43269b69bf1ec16)
	}
	return __obf_b99f62eaceb52d40
}

// increment the previous entropy number with a random number
// of up to m.inc (inclusive).
func (__obf_0db4e7f4c3797edd *MonotonicEntropy) __obf_510b83a3e924d48e() error {
	if __obf_534044727099f444, __obf_b99f62eaceb52d40 := __obf_0db4e7f4c3797edd.__obf_bf00bacba2926418(); __obf_b99f62eaceb52d40 != nil {
		return __obf_b99f62eaceb52d40
	} else if __obf_0db4e7f4c3797edd.__obf_b43269b69bf1ec16.Add(__obf_534044727099f444) {
		return ErrMonotonicOverflow
	}
	return nil
}

// random returns a uniform random value in [1, m.inc), reading entropy
// from m.Reader. When m.inc == 0 || m.inc == 1, it returns 1.
// Adapted from: https://golang.org/pkg/crypto/rand/#Int
func (__obf_0db4e7f4c3797edd *MonotonicEntropy) __obf_bf00bacba2926418() (__obf_534044727099f444 uint64, __obf_b99f62eaceb52d40 error) {
	if __obf_0db4e7f4c3797edd.__obf_534044727099f444 <= 1 {
		return 1, nil
	}

	// Fast path for using a underlying rand.Rand directly.
	if __obf_0db4e7f4c3797edd.__obf_0a92aebefacfe4a3 != nil {
		// Range: [1, m.inc)
		return 1 + uint64(__obf_0db4e7f4c3797edd.__obf_0a92aebefacfe4a3.Int63n(int64(__obf_0db4e7f4c3797edd.__obf_534044727099f444))), nil
	}

	// bitLen is the maximum bit length needed to encode a value < m.inc.
	__obf_d3c0aba9e12addcc := bits.Len64(__obf_0db4e7f4c3797edd.__obf_534044727099f444)

	// byteLen is the maximum byte length needed to encode a value < m.inc.
	__obf_33994bd469669ccd := uint(__obf_d3c0aba9e12addcc+7) / 8

	// msbitLen is the number of bits in the most significant byte of m.inc-1.
	__obf_e99dc5ec007fcced := uint(__obf_d3c0aba9e12addcc % 8)
	if __obf_e99dc5ec007fcced == 0 {
		__obf_e99dc5ec007fcced = 8
	}

	for __obf_534044727099f444 == 0 || __obf_534044727099f444 >= __obf_0db4e7f4c3797edd.__obf_534044727099f444 {
		if _, __obf_b99f62eaceb52d40 = io.ReadFull(__obf_0db4e7f4c3797edd.Reader, __obf_0db4e7f4c3797edd.rand[:__obf_33994bd469669ccd]); __obf_b99f62eaceb52d40 != nil {
			return 0, __obf_b99f62eaceb52d40
		}

		// Clear bits in the first byte to increase the probability
		// that the candidate is < m.inc.
		__obf_0db4e7f4c3797edd.rand[0] &= uint8(int(1<<__obf_e99dc5ec007fcced) - 1)

		// Convert the read bytes into an uint64 with byteLen
		// Optimized unrolled loop.
		switch __obf_33994bd469669ccd {
		case 1:
			__obf_534044727099f444 = uint64(__obf_0db4e7f4c3797edd.rand[0])
		case 2:
			__obf_534044727099f444 = uint64(binary.LittleEndian.Uint16(__obf_0db4e7f4c3797edd.rand[:2]))
		case 3, 4:
			__obf_534044727099f444 = uint64(binary.LittleEndian.Uint32(__obf_0db4e7f4c3797edd.rand[:4]))
		case 5, 6, 7, 8:
			__obf_534044727099f444 = uint64(binary.LittleEndian.Uint64(__obf_0db4e7f4c3797edd.rand[:8]))
		}
	}

	// Range: [1, m.inc)
	return 1 + __obf_534044727099f444, nil
}

type __obf_87b078d1bfc5f077 struct {
	Hi uint16
	Lo uint64
}

func (__obf_feb8e5a89f005fb4 *__obf_87b078d1bfc5f077) SetBytes(__obf_51bf1d51a16f7b1c []byte) {
	__obf_feb8e5a89f005fb4.Hi = binary.BigEndian.Uint16(__obf_51bf1d51a16f7b1c[:2])
	__obf_feb8e5a89f005fb4.Lo = binary.BigEndian.Uint64(__obf_51bf1d51a16f7b1c[2:])
}

func (__obf_feb8e5a89f005fb4 *__obf_87b078d1bfc5f077) AppendTo(__obf_51bf1d51a16f7b1c []byte) {
	binary.BigEndian.PutUint16(__obf_51bf1d51a16f7b1c[:2], __obf_feb8e5a89f005fb4.Hi)
	binary.BigEndian.PutUint64(__obf_51bf1d51a16f7b1c[2:], __obf_feb8e5a89f005fb4.Lo)
}

func (__obf_feb8e5a89f005fb4 *__obf_87b078d1bfc5f077) Add(__obf_097b9edd67ed12ac uint64) (__obf_87dec720da48b649 bool) {
	__obf_3a6fd7e784d8cc35, __obf_f19bd0b19e98565e := __obf_feb8e5a89f005fb4.Lo, __obf_feb8e5a89f005fb4.Hi
	if __obf_feb8e5a89f005fb4.Lo += __obf_097b9edd67ed12ac; __obf_feb8e5a89f005fb4.Lo < __obf_3a6fd7e784d8cc35 {
		__obf_feb8e5a89f005fb4.Hi++
	}
	return __obf_feb8e5a89f005fb4.Hi < __obf_f19bd0b19e98565e
}

func (__obf_feb8e5a89f005fb4 __obf_87b078d1bfc5f077) IsZero() bool {
	return __obf_feb8e5a89f005fb4.Hi == 0 && __obf_feb8e5a89f005fb4.Lo == 0
}
