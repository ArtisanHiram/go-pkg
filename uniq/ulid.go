package __obf_3747a7e09ff475ee

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
	MonotonicRead(__obf_b65816d8a2627aaa uint64, __obf_51c25155de9addfb []byte) error
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
func New(__obf_b65816d8a2627aaa uint64, __obf_921ce66321acc48d io.Reader) (__obf_413d3a76ff8ab79e ULID, __obf_57cef2f010b4ce9a error) {
	if __obf_57cef2f010b4ce9a = __obf_413d3a76ff8ab79e.SetTime(__obf_b65816d8a2627aaa); __obf_57cef2f010b4ce9a != nil {
		return __obf_413d3a76ff8ab79e, __obf_57cef2f010b4ce9a
	}

	switch __obf_b9b0ed2418aa36c1 := __obf_921ce66321acc48d.(type) {
	case nil:
		return __obf_413d3a76ff8ab79e, __obf_57cef2f010b4ce9a
	case MonotonicReader:
		__obf_57cef2f010b4ce9a = __obf_b9b0ed2418aa36c1.MonotonicRead(__obf_b65816d8a2627aaa, __obf_413d3a76ff8ab79e[6:])
	default:
		_, __obf_57cef2f010b4ce9a = io.ReadFull(__obf_b9b0ed2418aa36c1, __obf_413d3a76ff8ab79e[6:])
	}

	return __obf_413d3a76ff8ab79e,

		// MustNew is a convenience function equivalent to New that panics on failure
		// instead of returning an error.
		__obf_57cef2f010b4ce9a
}

func MustNew(__obf_b65816d8a2627aaa uint64, __obf_921ce66321acc48d io.Reader) ULID {
	__obf_413d3a76ff8ab79e, __obf_57cef2f010b4ce9a := New(__obf_b65816d8a2627aaa, __obf_921ce66321acc48d)
	if __obf_57cef2f010b4ce9a != nil {
		panic(__obf_57cef2f010b4ce9a)
	}
	return __obf_413d3a76ff8ab79e
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(__obf_d0183ab9cd977f49 string) (__obf_413d3a76ff8ab79e ULID, __obf_57cef2f010b4ce9a error) {
	return __obf_413d3a76ff8ab79e, __obf_2db033fe6028a6f6([]byte(__obf_d0183ab9cd977f49), false, &__obf_413d3a76ff8ab79e)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(__obf_d0183ab9cd977f49 string) (__obf_413d3a76ff8ab79e ULID, __obf_57cef2f010b4ce9a error) {
	return __obf_413d3a76ff8ab79e, __obf_2db033fe6028a6f6([]byte(__obf_d0183ab9cd977f49), true, &__obf_413d3a76ff8ab79e)
}

func __obf_2db033fe6028a6f6(__obf_5ed0238ba74cbb91 []byte, __obf_85b3e483f779e7c9 bool, __obf_413d3a76ff8ab79e *ULID) error {
	// Check if a base32 encoded ULID is the right length.
	if len(__obf_5ed0238ba74cbb91) != EncodedSize {
		return ErrDataSize
	}

	// Check if all the characters in a base32 encoded ULID are part of the
	// expected base32 character set.
	if __obf_85b3e483f779e7c9 &&
		(__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[0]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[1]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[2]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[3]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[4]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[5]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[6]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[7]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[8]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[9]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[10]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[11]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[12]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[13]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[14]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[15]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[16]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[17]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[18]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[19]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[20]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[21]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[22]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[23]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[24]] == 0xFF || __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[25]] == 0xFF) {
		return ErrInvalidCharacters
	}

	// Check if the first character in a base32 encoded ULID will overflow. This
	// happens because the base32 representation encodes 130 bits, while the
	// ULID is only 128 bits.
	//
	// See https://github.com/oklog/ulid/issues/9 for details.
	if __obf_5ed0238ba74cbb91[0] > '7' {
		return ErrOverflow
	}

	// Use an optimized unrolled loop (from https://github.com/RobThree/NUlid)
	// to decode a base32 ULID.

	// 6 bytes timestamp (48 bits)
	(*__obf_413d3a76ff8ab79e)[0] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[0]] << 5) | __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[1]]
	(*__obf_413d3a76ff8ab79e)[1] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[2]] << 3) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[3]] >> 2)
	(*__obf_413d3a76ff8ab79e)[2] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[3]] << 6) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[4]] << 1) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[5]] >> 4)
	(*__obf_413d3a76ff8ab79e)[3] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[5]] << 4) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[6]] >> 1)
	(*__obf_413d3a76ff8ab79e)[4] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[6]] << 7) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[7]] << 2) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[8]] >> 3)
	(*__obf_413d3a76ff8ab79e)[5] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[8]] << 5) | __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[9]]

	// 10 bytes of entropy (80 bits)
	(*__obf_413d3a76ff8ab79e)[6] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[10]] << 3) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[11]] >> 2)
	(*__obf_413d3a76ff8ab79e)[7] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[11]] << 6) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[12]] << 1) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[13]] >> 4)
	(*__obf_413d3a76ff8ab79e)[8] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[13]] << 4) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[14]] >> 1)
	(*__obf_413d3a76ff8ab79e)[9] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[14]] << 7) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[15]] << 2) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[16]] >> 3)
	(*__obf_413d3a76ff8ab79e)[10] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[16]] << 5) | __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[17]]
	(*__obf_413d3a76ff8ab79e)[11] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[18]] << 3) | __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[19]]>>2
	(*__obf_413d3a76ff8ab79e)[12] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[19]] << 6) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[20]] << 1) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[21]] >> 4)
	(*__obf_413d3a76ff8ab79e)[13] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[21]] << 4) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[22]] >> 1)
	(*__obf_413d3a76ff8ab79e)[14] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[22]] << 7) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[23]] << 2) | (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[24]] >> 3)
	(*__obf_413d3a76ff8ab79e)[15] = (__obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[24]] << 5) | __obf_651dce1378d9c1c8[__obf_5ed0238ba74cbb91[25]]

	return nil
}

// MustParse is a convenience function equivalent to Parse that panics on failure
// instead of returning an error.
func MustParse(__obf_d0183ab9cd977f49 string) ULID {
	__obf_413d3a76ff8ab79e, __obf_57cef2f010b4ce9a := Parse(__obf_d0183ab9cd977f49)
	if __obf_57cef2f010b4ce9a != nil {
		panic(__obf_57cef2f010b4ce9a)
	}
	return __obf_413d3a76ff8ab79e
}

// MustParseStrict is a convenience function equivalent to ParseStrict that
// panics on failure instead of returning an error.
func MustParseStrict(__obf_d0183ab9cd977f49 string) ULID {
	__obf_413d3a76ff8ab79e, __obf_57cef2f010b4ce9a := ParseStrict(__obf_d0183ab9cd977f49)
	if __obf_57cef2f010b4ce9a != nil {
		panic(__obf_57cef2f010b4ce9a)
	}
	return __obf_413d3a76ff8ab79e
}

// Bytes returns bytes slice representation of ULID.
func (__obf_b4dcaafcadeeef58 ULID) Bytes() []byte {
	return __obf_b4dcaafcadeeef58[:]
}

// String returns a lexicographically sortable string encoded ULID
// (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy
func (__obf_413d3a76ff8ab79e ULID) String() string {
	__obf_d0183ab9cd977f49 := make([]byte, EncodedSize)
	_ = __obf_413d3a76ff8ab79e.MarshalTextTo(__obf_d0183ab9cd977f49)
	return string(__obf_d0183ab9cd977f49)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (__obf_413d3a76ff8ab79e ULID) MarshalBinary() ([]byte, error) {
	__obf_d0183ab9cd977f49 := make([]byte, len(__obf_413d3a76ff8ab79e))
	return __obf_d0183ab9cd977f49, __obf_413d3a76ff8ab79e.MarshalBinaryTo(__obf_d0183ab9cd977f49)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (__obf_413d3a76ff8ab79e ULID) MarshalBinaryTo(__obf_5018ed930160ab87 []byte) error {
	if len(__obf_5018ed930160ab87) != len(__obf_413d3a76ff8ab79e) {
		return ErrBufferSize
	}

	copy(__obf_5018ed930160ab87, __obf_413d3a76ff8ab79e[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (__obf_413d3a76ff8ab79e *ULID) UnmarshalBinary(__obf_bf3b3775c48456e5 []byte) error {
	if len(__obf_bf3b3775c48456e5) != len(*__obf_413d3a76ff8ab79e) {
		return ErrDataSize
	}

	copy((*__obf_413d3a76ff8ab79e)[:], __obf_bf3b3775c48456e5)
	return nil
}

// Encoding is the base 32 encoding alphabet used in ULID strings.
const Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (__obf_413d3a76ff8ab79e ULID) MarshalText() ([]byte, error) {
	__obf_d0183ab9cd977f49 := make([]byte, EncodedSize)
	return __obf_d0183ab9cd977f49, __obf_413d3a76ff8ab79e.MarshalTextTo(__obf_d0183ab9cd977f49)
}

// MarshalTextTo writes the ULID as a string to the given buffer.
// ErrBufferSize is returned when the len(dst) != 26.
func (__obf_413d3a76ff8ab79e ULID) MarshalTextTo(__obf_5018ed930160ab87 []byte) error {
	// Optimized unrolled loop ahead.
	// From https://github.com/RobThree/NUlid

	if len(__obf_5018ed930160ab87) != EncodedSize {
		return ErrBufferSize
	}
	__obf_5018ed930160ab87[ // 10 byte timestamp
	0] = Encoding[(__obf_413d3a76ff8ab79e[0]&224)>>5]
	__obf_5018ed930160ab87[1] = Encoding[__obf_413d3a76ff8ab79e[0]&31]
	__obf_5018ed930160ab87[2] = Encoding[(__obf_413d3a76ff8ab79e[1]&248)>>3]
	__obf_5018ed930160ab87[3] = Encoding[((__obf_413d3a76ff8ab79e[1]&7)<<2)|((__obf_413d3a76ff8ab79e[2]&192)>>6)]
	__obf_5018ed930160ab87[4] = Encoding[(__obf_413d3a76ff8ab79e[2]&62)>>1]
	__obf_5018ed930160ab87[5] = Encoding[((__obf_413d3a76ff8ab79e[2]&1)<<4)|((__obf_413d3a76ff8ab79e[3]&240)>>4)]
	__obf_5018ed930160ab87[6] = Encoding[((__obf_413d3a76ff8ab79e[3]&15)<<1)|((__obf_413d3a76ff8ab79e[4]&128)>>7)]
	__obf_5018ed930160ab87[7] = Encoding[(__obf_413d3a76ff8ab79e[4]&124)>>2]
	__obf_5018ed930160ab87[8] = Encoding[((__obf_413d3a76ff8ab79e[4]&3)<<3)|((__obf_413d3a76ff8ab79e[5]&224)>>5)]
	__obf_5018ed930160ab87[9] = Encoding[__obf_413d3a76ff8ab79e[5]&31]
	__obf_5018ed930160ab87[ // 16 bytes of entropy
	10] = Encoding[(__obf_413d3a76ff8ab79e[6]&248)>>3]
	__obf_5018ed930160ab87[11] = Encoding[((__obf_413d3a76ff8ab79e[6]&7)<<2)|((__obf_413d3a76ff8ab79e[7]&192)>>6)]
	__obf_5018ed930160ab87[12] = Encoding[(__obf_413d3a76ff8ab79e[7]&62)>>1]
	__obf_5018ed930160ab87[13] = Encoding[((__obf_413d3a76ff8ab79e[7]&1)<<4)|((__obf_413d3a76ff8ab79e[8]&240)>>4)]
	__obf_5018ed930160ab87[14] = Encoding[((__obf_413d3a76ff8ab79e[8]&15)<<1)|((__obf_413d3a76ff8ab79e[9]&128)>>7)]
	__obf_5018ed930160ab87[15] = Encoding[(__obf_413d3a76ff8ab79e[9]&124)>>2]
	__obf_5018ed930160ab87[16] = Encoding[((__obf_413d3a76ff8ab79e[9]&3)<<3)|((__obf_413d3a76ff8ab79e[10]&224)>>5)]
	__obf_5018ed930160ab87[17] = Encoding[__obf_413d3a76ff8ab79e[10]&31]
	__obf_5018ed930160ab87[18] = Encoding[(__obf_413d3a76ff8ab79e[11]&248)>>3]
	__obf_5018ed930160ab87[19] = Encoding[((__obf_413d3a76ff8ab79e[11]&7)<<2)|((__obf_413d3a76ff8ab79e[12]&192)>>6)]
	__obf_5018ed930160ab87[20] = Encoding[(__obf_413d3a76ff8ab79e[12]&62)>>1]
	__obf_5018ed930160ab87[21] = Encoding[((__obf_413d3a76ff8ab79e[12]&1)<<4)|((__obf_413d3a76ff8ab79e[13]&240)>>4)]
	__obf_5018ed930160ab87[22] = Encoding[((__obf_413d3a76ff8ab79e[13]&15)<<1)|((__obf_413d3a76ff8ab79e[14]&128)>>7)]
	__obf_5018ed930160ab87[23] = Encoding[(__obf_413d3a76ff8ab79e[14]&124)>>2]
	__obf_5018ed930160ab87[24] = Encoding[((__obf_413d3a76ff8ab79e[14]&3)<<3)|((__obf_413d3a76ff8ab79e[15]&224)>>5)]
	__obf_5018ed930160ab87[25] = Encoding[__obf_413d3a76ff8ab79e[15]&31]

	return nil
}

// Byte to index table for O(1) lookups when unmarshaling.
// We use 0xFF as sentinel value for invalid indexes.
var __obf_651dce1378d9c1c8 = [...]byte{
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
func (__obf_413d3a76ff8ab79e *ULID) UnmarshalText(__obf_5ed0238ba74cbb91 []byte) error {
	return __obf_2db033fe6028a6f6(__obf_5ed0238ba74cbb91, false, __obf_413d3a76ff8ab79e)
}

// Time returns the Unix time in milliseconds encoded in the ULID.
// Use the top level Time function to convert the returned value to
// a time.Time.
func (__obf_413d3a76ff8ab79e ULID) Time() uint64 {
	return uint64(__obf_413d3a76ff8ab79e[5]) | uint64(__obf_413d3a76ff8ab79e[4])<<8 |
		uint64(__obf_413d3a76ff8ab79e[3])<<16 | uint64(__obf_413d3a76ff8ab79e[2])<<24 |
		uint64(__obf_413d3a76ff8ab79e[1])<<32 | uint64(__obf_413d3a76ff8ab79e[0])<<40
}

// maxTime is the maximum Unix time in milliseconds that can be
// represented in an ULID.
var __obf_747edd5745186d02 = ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}.Time()

// MaxTime returns the maximum Unix time in milliseconds that
// can be encoded in an ULID.
func MaxTime() uint64 { return __obf_747edd5745186d02 }

// Now is a convenience function that returns the current
// UTC time in Unix milliseconds. Equivalent to:
//
//	Timestamp(time.Now().UTC())
func Now() uint64 { return Timestamp(time.Now().UTC()) }

// Timestamp converts a time.Time to Unix milliseconds.
//
// Because of the way ULID stores time, times from the year
// 10889 produces undefined results.
func Timestamp(__obf_b5849d1d191d006d time.Time) uint64 {
	return uint64(__obf_b5849d1d191d006d.Unix())*1000 +
		uint64(__obf_b5849d1d191d006d.Nanosecond()/int(time.Millisecond))
}

// Time converts Unix milliseconds in the format
// returned by the Timestamp function to a time.Time.
func Time(__obf_b65816d8a2627aaa uint64) time.Time {
	__obf_37226e04c5d29f60 := int64(__obf_b65816d8a2627aaa / 1e3)
	__obf_7fa0523073142870 := int64((__obf_b65816d8a2627aaa % 1e3) * 1e6)
	return time.Unix(__obf_37226e04c5d29f60,

		// SetTime sets the time component of the ULID to the given Unix time
		// in milliseconds.
		__obf_7fa0523073142870)
}

func (__obf_413d3a76ff8ab79e *ULID) SetTime(__obf_b65816d8a2627aaa uint64) error {
	if __obf_b65816d8a2627aaa > __obf_747edd5745186d02 {
		return ErrBigTime
	}

	(*__obf_413d3a76ff8ab79e)[0] = byte(__obf_b65816d8a2627aaa >> 40)
	(*__obf_413d3a76ff8ab79e)[1] = byte(__obf_b65816d8a2627aaa >> 32)
	(*__obf_413d3a76ff8ab79e)[2] = byte(__obf_b65816d8a2627aaa >> 24)
	(*__obf_413d3a76ff8ab79e)[3] = byte(__obf_b65816d8a2627aaa >> 16)
	(*__obf_413d3a76ff8ab79e)[4] = byte(__obf_b65816d8a2627aaa >> 8)
	(*__obf_413d3a76ff8ab79e)[5] = byte(__obf_b65816d8a2627aaa)

	return nil
}

// Entropy returns the entropy from the ULID.
func (__obf_413d3a76ff8ab79e ULID) Entropy() []byte {
	__obf_b9b0ed2418aa36c1 := make([]byte, 10)
	copy(__obf_b9b0ed2418aa36c1, __obf_413d3a76ff8ab79e[6:])
	return __obf_b9b0ed2418aa36c1
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (__obf_413d3a76ff8ab79e *ULID) SetEntropy(__obf_b9b0ed2418aa36c1 []byte) error {
	if len(__obf_b9b0ed2418aa36c1) != 10 {
		return ErrDataSize
	}

	copy((*__obf_413d3a76ff8ab79e)[6:], __obf_b9b0ed2418aa36c1)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (__obf_413d3a76ff8ab79e ULID) Compare(__obf_7fa65aa0966513bb ULID) int {
	return bytes.Compare(__obf_413d3a76ff8ab79e[:], __obf_7fa65aa0966513bb[:])
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (__obf_413d3a76ff8ab79e *ULID) Scan(__obf_3a1f6937fa218095 any) error {
	switch __obf_d2ffb8381b91d2b8 := __obf_3a1f6937fa218095.(type) {
	case nil:
		return nil
	case string:
		return __obf_413d3a76ff8ab79e.UnmarshalText([]byte(__obf_d2ffb8381b91d2b8))
	case []byte:
		return __obf_413d3a76ff8ab79e.UnmarshalBinary(__obf_d2ffb8381b91d2b8)
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
func (__obf_413d3a76ff8ab79e ULID) Value() (driver.Value, error) {
	return __obf_413d3a76ff8ab79e.MarshalBinary()
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
func Monotonic(__obf_921ce66321acc48d io.Reader, __obf_9616a04bb3f6fe65 uint64) *MonotonicEntropy {
	__obf_17bcb19ef8ded267 := MonotonicEntropy{
		Reader: bufio.NewReader(__obf_921ce66321acc48d), __obf_9616a04bb3f6fe65: __obf_9616a04bb3f6fe65,
	}

	if __obf_17bcb19ef8ded267.__obf_9616a04bb3f6fe65 == 0 {
		__obf_17bcb19ef8ded267.__obf_9616a04bb3f6fe65 = math.MaxUint32
	}

	if __obf_52673c5d7080ff8a, __obf_0941b7586b921890 := __obf_921ce66321acc48d.(*rand.Rand); __obf_0941b7586b921890 {
		__obf_17bcb19ef8ded267.__obf_52673c5d7080ff8a = __obf_52673c5d7080ff8a
	}

	return &__obf_17bcb19ef8ded267
}

// MonotonicEntropy is an opaque type that provides monotonic entropy.
type MonotonicEntropy struct {
	io.Reader
	__obf_b65816d8a2627aaa uint64
	__obf_9616a04bb3f6fe65 uint64
	__obf_921ce66321acc48d __obf_e2c8dee1f909a32d

	rand                   [8]byte
	__obf_52673c5d7080ff8a *rand.Rand
}

// MonotonicRead implements the MonotonicReader interface.
func (__obf_17bcb19ef8ded267 *MonotonicEntropy) MonotonicRead(__obf_b65816d8a2627aaa uint64, __obf_921ce66321acc48d []byte) (__obf_57cef2f010b4ce9a error) {
	if !__obf_17bcb19ef8ded267.__obf_921ce66321acc48d.IsZero() && __obf_17bcb19ef8ded267.__obf_b65816d8a2627aaa == __obf_b65816d8a2627aaa {
		__obf_57cef2f010b4ce9a = __obf_17bcb19ef8ded267.__obf_12b1a81dd37525c7()
		__obf_17bcb19ef8ded267.__obf_921ce66321acc48d.
			AppendTo(__obf_921ce66321acc48d)
	} else if _, __obf_57cef2f010b4ce9a = io.ReadFull(__obf_17bcb19ef8ded267.Reader, __obf_921ce66321acc48d); __obf_57cef2f010b4ce9a == nil {
		__obf_17bcb19ef8ded267.__obf_b65816d8a2627aaa = __obf_b65816d8a2627aaa
		__obf_17bcb19ef8ded267.__obf_921ce66321acc48d.
			SetBytes(__obf_921ce66321acc48d)
	}
	return __obf_57cef2f010b4ce9a
}

// increment the previous entropy number with a random number
// of up to m.inc (inclusive).
func (__obf_17bcb19ef8ded267 *MonotonicEntropy) __obf_12b1a81dd37525c7() error {
	if __obf_9616a04bb3f6fe65, __obf_57cef2f010b4ce9a := __obf_17bcb19ef8ded267.__obf_b0c0f88e772863b0(); __obf_57cef2f010b4ce9a != nil {
		return __obf_57cef2f010b4ce9a
	} else if __obf_17bcb19ef8ded267.__obf_921ce66321acc48d.Add(__obf_9616a04bb3f6fe65) {
		return ErrMonotonicOverflow
	}
	return nil
}

// random returns a uniform random value in [1, m.inc), reading entropy
// from m.Reader. When m.inc == 0 || m.inc == 1, it returns 1.
// Adapted from: https://golang.org/pkg/crypto/rand/#Int
func (__obf_17bcb19ef8ded267 *MonotonicEntropy) __obf_b0c0f88e772863b0() (__obf_9616a04bb3f6fe65 uint64, __obf_57cef2f010b4ce9a error) {
	if __obf_17bcb19ef8ded267.__obf_9616a04bb3f6fe65 <= 1 {
		return 1, nil
	}

	// Fast path for using a underlying rand.Rand directly.
	if __obf_17bcb19ef8ded267.
		// Range: [1, m.inc)
		__obf_52673c5d7080ff8a != nil {

		return 1 + uint64(__obf_17bcb19ef8ded267.__obf_52673c5d7080ff8a.Int63n(int64(__obf_17bcb19ef8ded267.

			// bitLen is the maximum bit length needed to encode a value < m.inc.
			__obf_9616a04bb3f6fe65))), nil
	}
	__obf_d17fcd4e80d87820 := bits.Len64(__obf_17bcb19ef8ded267.

		// byteLen is the maximum byte length needed to encode a value < m.inc.
		__obf_9616a04bb3f6fe65)
	__obf_2272ec6fda939c17 := uint(__obf_d17fcd4e80d87820+7) / 8
	__obf_2498c95fe7e151f8 := // msbitLen is the number of bits in the most significant byte of m.inc-1.
		uint(__obf_d17fcd4e80d87820 % 8)
	if __obf_2498c95fe7e151f8 == 0 {
		__obf_2498c95fe7e151f8 = 8
	}

	for __obf_9616a04bb3f6fe65 == 0 || __obf_9616a04bb3f6fe65 >= __obf_17bcb19ef8ded267.__obf_9616a04bb3f6fe65 {
		if _, __obf_57cef2f010b4ce9a = io.ReadFull(__obf_17bcb19ef8ded267.Reader, __obf_17bcb19ef8ded267.rand[:__obf_2272ec6fda939c17]); __obf_57cef2f010b4ce9a != nil {
			return 0, __obf_57cef2f010b4ce9a
		}
		__obf_17bcb19ef8ded267.

			// Clear bits in the first byte to increase the probability
			// that the candidate is < m.inc.
			rand[0] &= uint8(int(1<<__obf_2498c95fe7e151f8) - 1)

		// Convert the read bytes into an uint64 with byteLen
		// Optimized unrolled loop.
		switch __obf_2272ec6fda939c17 {
		case 1:
			__obf_9616a04bb3f6fe65 = uint64(__obf_17bcb19ef8ded267.rand[0])
		case 2:
			__obf_9616a04bb3f6fe65 = uint64(binary.LittleEndian.Uint16(__obf_17bcb19ef8ded267.rand[:2]))
		case 3, 4:
			__obf_9616a04bb3f6fe65 = uint64(binary.LittleEndian.Uint32(__obf_17bcb19ef8ded267.rand[:4]))
		case 5, 6, 7, 8:
			__obf_9616a04bb3f6fe65 = uint64(binary.LittleEndian.Uint64(__obf_17bcb19ef8ded267.rand[:8]))
		}
	}

	// Range: [1, m.inc)
	return 1 + __obf_9616a04bb3f6fe65, nil
}

type __obf_e2c8dee1f909a32d struct {
	Hi uint16
	Lo uint64
}

func (__obf_b4dcaafcadeeef58 *__obf_e2c8dee1f909a32d) SetBytes(__obf_01962d395bcb9642 []byte) {
	__obf_b4dcaafcadeeef58.
		Hi = binary.BigEndian.Uint16(__obf_01962d395bcb9642[:2])
	__obf_b4dcaafcadeeef58.
		Lo = binary.BigEndian.Uint64(__obf_01962d395bcb9642[2:])
}

func (__obf_b4dcaafcadeeef58 *__obf_e2c8dee1f909a32d) AppendTo(__obf_01962d395bcb9642 []byte) {
	binary.BigEndian.PutUint16(__obf_01962d395bcb9642[:2], __obf_b4dcaafcadeeef58.Hi)
	binary.BigEndian.PutUint64(__obf_01962d395bcb9642[2:], __obf_b4dcaafcadeeef58.Lo)
}

func (__obf_b4dcaafcadeeef58 *__obf_e2c8dee1f909a32d) Add(__obf_368df23032f058d2 uint64) (__obf_774b6119dc7ccab3 bool) {
	__obf_16003a1faba02566, __obf_292472b6f9a11213 := __obf_b4dcaafcadeeef58.Lo, __obf_b4dcaafcadeeef58.Hi
	if __obf_b4dcaafcadeeef58.Lo += __obf_368df23032f058d2; __obf_b4dcaafcadeeef58.Lo < __obf_16003a1faba02566 {
		__obf_b4dcaafcadeeef58.
			Hi++
	}
	return __obf_b4dcaafcadeeef58.Hi < __obf_292472b6f9a11213
}

func (__obf_b4dcaafcadeeef58 __obf_e2c8dee1f909a32d) IsZero() bool {
	return __obf_b4dcaafcadeeef58.Hi == 0 && __obf_b4dcaafcadeeef58.Lo == 0
}
