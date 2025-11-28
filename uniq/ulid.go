package __obf_8fe28252c1b01dfe

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
	MonotonicRead(__obf_a51f50affa0720aa uint64, __obf_507583319a255883 []byte) error
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
func New(__obf_a51f50affa0720aa uint64, __obf_7fc7e7bc83026231 io.Reader) (__obf_b6d2ddd191a19ef1 ULID, __obf_1082dd6e56192e3a error) {
	if __obf_1082dd6e56192e3a = __obf_b6d2ddd191a19ef1.SetTime(__obf_a51f50affa0720aa); __obf_1082dd6e56192e3a != nil {
		return __obf_b6d2ddd191a19ef1, __obf_1082dd6e56192e3a
	}

	switch __obf_6d1c9d77673ade75 := __obf_7fc7e7bc83026231.(type) {
	case nil:
		return __obf_b6d2ddd191a19ef1, __obf_1082dd6e56192e3a
	case MonotonicReader:
		__obf_1082dd6e56192e3a = __obf_6d1c9d77673ade75.MonotonicRead(__obf_a51f50affa0720aa, __obf_b6d2ddd191a19ef1[6:])
	default:
		_, __obf_1082dd6e56192e3a = io.ReadFull(__obf_6d1c9d77673ade75, __obf_b6d2ddd191a19ef1[6:])
	}

	return __obf_b6d2ddd191a19ef1, __obf_1082dd6e56192e3a
}

// MustNew is a convenience function equivalent to New that panics on failure
// instead of returning an error.
func MustNew(__obf_a51f50affa0720aa uint64, __obf_7fc7e7bc83026231 io.Reader) ULID {
	__obf_b6d2ddd191a19ef1, __obf_1082dd6e56192e3a := New(__obf_a51f50affa0720aa, __obf_7fc7e7bc83026231)
	if __obf_1082dd6e56192e3a != nil {
		panic(__obf_1082dd6e56192e3a)
	}
	return __obf_b6d2ddd191a19ef1
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(__obf_3fbf3c379fcfc016 string) (__obf_b6d2ddd191a19ef1 ULID, __obf_1082dd6e56192e3a error) {
	return __obf_b6d2ddd191a19ef1, __obf_9e258196bd35a28d([]byte(__obf_3fbf3c379fcfc016), false, &__obf_b6d2ddd191a19ef1)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(__obf_3fbf3c379fcfc016 string) (__obf_b6d2ddd191a19ef1 ULID, __obf_1082dd6e56192e3a error) {
	return __obf_b6d2ddd191a19ef1, __obf_9e258196bd35a28d([]byte(__obf_3fbf3c379fcfc016), true, &__obf_b6d2ddd191a19ef1)
}

func __obf_9e258196bd35a28d(__obf_3b548eea5a5f82ce []byte, __obf_c279ac8413a6c8aa bool, __obf_b6d2ddd191a19ef1 *ULID) error {
	// Check if a base32 encoded ULID is the right length.
	if len(__obf_3b548eea5a5f82ce) != EncodedSize {
		return ErrDataSize
	}

	// Check if all the characters in a base32 encoded ULID are part of the
	// expected base32 character set.
	if __obf_c279ac8413a6c8aa &&
		(__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[0]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[1]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[2]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[3]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[4]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[5]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[6]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[7]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[8]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[9]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[10]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[11]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[12]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[13]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[14]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[15]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[16]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[17]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[18]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[19]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[20]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[21]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[22]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[23]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[24]] == 0xFF ||
			__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[25]] == 0xFF) {
		return ErrInvalidCharacters
	}

	// Check if the first character in a base32 encoded ULID will overflow. This
	// happens because the base32 representation encodes 130 bits, while the
	// ULID is only 128 bits.
	//
	// See https://github.com/oklog/ulid/issues/9 for details.
	if __obf_3b548eea5a5f82ce[0] > '7' {
		return ErrOverflow
	}

	// Use an optimized unrolled loop (from https://github.com/RobThree/NUlid)
	// to decode a base32 ULID.

	// 6 bytes timestamp (48 bits)
	(*__obf_b6d2ddd191a19ef1)[0] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[0]] << 5) | __obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[1]]
	(*__obf_b6d2ddd191a19ef1)[1] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[2]] << 3) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[3]] >> 2)
	(*__obf_b6d2ddd191a19ef1)[2] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[3]] << 6) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[4]] << 1) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[5]] >> 4)
	(*__obf_b6d2ddd191a19ef1)[3] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[5]] << 4) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[6]] >> 1)
	(*__obf_b6d2ddd191a19ef1)[4] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[6]] << 7) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[7]] << 2) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[8]] >> 3)
	(*__obf_b6d2ddd191a19ef1)[5] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[8]] << 5) | __obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[9]]

	// 10 bytes of entropy (80 bits)
	(*__obf_b6d2ddd191a19ef1)[6] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[10]] << 3) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[11]] >> 2)
	(*__obf_b6d2ddd191a19ef1)[7] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[11]] << 6) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[12]] << 1) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[13]] >> 4)
	(*__obf_b6d2ddd191a19ef1)[8] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[13]] << 4) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[14]] >> 1)
	(*__obf_b6d2ddd191a19ef1)[9] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[14]] << 7) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[15]] << 2) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[16]] >> 3)
	(*__obf_b6d2ddd191a19ef1)[10] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[16]] << 5) | __obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[17]]
	(*__obf_b6d2ddd191a19ef1)[11] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[18]] << 3) | __obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[19]]>>2
	(*__obf_b6d2ddd191a19ef1)[12] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[19]] << 6) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[20]] << 1) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[21]] >> 4)
	(*__obf_b6d2ddd191a19ef1)[13] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[21]] << 4) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[22]] >> 1)
	(*__obf_b6d2ddd191a19ef1)[14] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[22]] << 7) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[23]] << 2) | (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[24]] >> 3)
	(*__obf_b6d2ddd191a19ef1)[15] = (__obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[24]] << 5) | __obf_5a65ade5e75208e0[__obf_3b548eea5a5f82ce[25]]

	return nil
}

// MustParse is a convenience function equivalent to Parse that panics on failure
// instead of returning an error.
func MustParse(__obf_3fbf3c379fcfc016 string) ULID {
	__obf_b6d2ddd191a19ef1, __obf_1082dd6e56192e3a := Parse(__obf_3fbf3c379fcfc016)
	if __obf_1082dd6e56192e3a != nil {
		panic(__obf_1082dd6e56192e3a)
	}
	return __obf_b6d2ddd191a19ef1
}

// MustParseStrict is a convenience function equivalent to ParseStrict that
// panics on failure instead of returning an error.
func MustParseStrict(__obf_3fbf3c379fcfc016 string) ULID {
	__obf_b6d2ddd191a19ef1, __obf_1082dd6e56192e3a := ParseStrict(__obf_3fbf3c379fcfc016)
	if __obf_1082dd6e56192e3a != nil {
		panic(__obf_1082dd6e56192e3a)
	}
	return __obf_b6d2ddd191a19ef1
}

// Bytes returns bytes slice representation of ULID.
func (__obf_6a86d57e2bf1634f ULID) Bytes() []byte {
	return __obf_6a86d57e2bf1634f[:]
}

// String returns a lexicographically sortable string encoded ULID
// (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy
func (__obf_b6d2ddd191a19ef1 ULID) String() string {
	__obf_3fbf3c379fcfc016 := make([]byte, EncodedSize)
	_ = __obf_b6d2ddd191a19ef1.MarshalTextTo(__obf_3fbf3c379fcfc016)
	return string(__obf_3fbf3c379fcfc016)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (__obf_b6d2ddd191a19ef1 ULID) MarshalBinary() ([]byte, error) {
	__obf_3fbf3c379fcfc016 := make([]byte, len(__obf_b6d2ddd191a19ef1))
	return __obf_3fbf3c379fcfc016, __obf_b6d2ddd191a19ef1.MarshalBinaryTo(__obf_3fbf3c379fcfc016)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (__obf_b6d2ddd191a19ef1 ULID) MarshalBinaryTo(__obf_fe6fc3b0b1a9c057 []byte) error {
	if len(__obf_fe6fc3b0b1a9c057) != len(__obf_b6d2ddd191a19ef1) {
		return ErrBufferSize
	}

	copy(__obf_fe6fc3b0b1a9c057, __obf_b6d2ddd191a19ef1[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (__obf_b6d2ddd191a19ef1 *ULID) UnmarshalBinary(__obf_cd3f2eae66228888 []byte) error {
	if len(__obf_cd3f2eae66228888) != len(*__obf_b6d2ddd191a19ef1) {
		return ErrDataSize
	}

	copy((*__obf_b6d2ddd191a19ef1)[:], __obf_cd3f2eae66228888)
	return nil
}

// Encoding is the base 32 encoding alphabet used in ULID strings.
const Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (__obf_b6d2ddd191a19ef1 ULID) MarshalText() ([]byte, error) {
	__obf_3fbf3c379fcfc016 := make([]byte, EncodedSize)
	return __obf_3fbf3c379fcfc016, __obf_b6d2ddd191a19ef1.MarshalTextTo(__obf_3fbf3c379fcfc016)
}

// MarshalTextTo writes the ULID as a string to the given buffer.
// ErrBufferSize is returned when the len(dst) != 26.
func (__obf_b6d2ddd191a19ef1 ULID) MarshalTextTo(__obf_fe6fc3b0b1a9c057 []byte) error {
	// Optimized unrolled loop ahead.
	// From https://github.com/RobThree/NUlid

	if len(__obf_fe6fc3b0b1a9c057) != EncodedSize {
		return ErrBufferSize
	}

	// 10 byte timestamp
	__obf_fe6fc3b0b1a9c057[0] = Encoding[(__obf_b6d2ddd191a19ef1[0]&224)>>5]
	__obf_fe6fc3b0b1a9c057[1] = Encoding[__obf_b6d2ddd191a19ef1[0]&31]
	__obf_fe6fc3b0b1a9c057[2] = Encoding[(__obf_b6d2ddd191a19ef1[1]&248)>>3]
	__obf_fe6fc3b0b1a9c057[3] = Encoding[((__obf_b6d2ddd191a19ef1[1]&7)<<2)|((__obf_b6d2ddd191a19ef1[2]&192)>>6)]
	__obf_fe6fc3b0b1a9c057[4] = Encoding[(__obf_b6d2ddd191a19ef1[2]&62)>>1]
	__obf_fe6fc3b0b1a9c057[5] = Encoding[((__obf_b6d2ddd191a19ef1[2]&1)<<4)|((__obf_b6d2ddd191a19ef1[3]&240)>>4)]
	__obf_fe6fc3b0b1a9c057[6] = Encoding[((__obf_b6d2ddd191a19ef1[3]&15)<<1)|((__obf_b6d2ddd191a19ef1[4]&128)>>7)]
	__obf_fe6fc3b0b1a9c057[7] = Encoding[(__obf_b6d2ddd191a19ef1[4]&124)>>2]
	__obf_fe6fc3b0b1a9c057[8] = Encoding[((__obf_b6d2ddd191a19ef1[4]&3)<<3)|((__obf_b6d2ddd191a19ef1[5]&224)>>5)]
	__obf_fe6fc3b0b1a9c057[9] = Encoding[__obf_b6d2ddd191a19ef1[5]&31]

	// 16 bytes of entropy
	__obf_fe6fc3b0b1a9c057[10] = Encoding[(__obf_b6d2ddd191a19ef1[6]&248)>>3]
	__obf_fe6fc3b0b1a9c057[11] = Encoding[((__obf_b6d2ddd191a19ef1[6]&7)<<2)|((__obf_b6d2ddd191a19ef1[7]&192)>>6)]
	__obf_fe6fc3b0b1a9c057[12] = Encoding[(__obf_b6d2ddd191a19ef1[7]&62)>>1]
	__obf_fe6fc3b0b1a9c057[13] = Encoding[((__obf_b6d2ddd191a19ef1[7]&1)<<4)|((__obf_b6d2ddd191a19ef1[8]&240)>>4)]
	__obf_fe6fc3b0b1a9c057[14] = Encoding[((__obf_b6d2ddd191a19ef1[8]&15)<<1)|((__obf_b6d2ddd191a19ef1[9]&128)>>7)]
	__obf_fe6fc3b0b1a9c057[15] = Encoding[(__obf_b6d2ddd191a19ef1[9]&124)>>2]
	__obf_fe6fc3b0b1a9c057[16] = Encoding[((__obf_b6d2ddd191a19ef1[9]&3)<<3)|((__obf_b6d2ddd191a19ef1[10]&224)>>5)]
	__obf_fe6fc3b0b1a9c057[17] = Encoding[__obf_b6d2ddd191a19ef1[10]&31]
	__obf_fe6fc3b0b1a9c057[18] = Encoding[(__obf_b6d2ddd191a19ef1[11]&248)>>3]
	__obf_fe6fc3b0b1a9c057[19] = Encoding[((__obf_b6d2ddd191a19ef1[11]&7)<<2)|((__obf_b6d2ddd191a19ef1[12]&192)>>6)]
	__obf_fe6fc3b0b1a9c057[20] = Encoding[(__obf_b6d2ddd191a19ef1[12]&62)>>1]
	__obf_fe6fc3b0b1a9c057[21] = Encoding[((__obf_b6d2ddd191a19ef1[12]&1)<<4)|((__obf_b6d2ddd191a19ef1[13]&240)>>4)]
	__obf_fe6fc3b0b1a9c057[22] = Encoding[((__obf_b6d2ddd191a19ef1[13]&15)<<1)|((__obf_b6d2ddd191a19ef1[14]&128)>>7)]
	__obf_fe6fc3b0b1a9c057[23] = Encoding[(__obf_b6d2ddd191a19ef1[14]&124)>>2]
	__obf_fe6fc3b0b1a9c057[24] = Encoding[((__obf_b6d2ddd191a19ef1[14]&3)<<3)|((__obf_b6d2ddd191a19ef1[15]&224)>>5)]
	__obf_fe6fc3b0b1a9c057[25] = Encoding[__obf_b6d2ddd191a19ef1[15]&31]

	return nil
}

// Byte to index table for O(1) lookups when unmarshaling.
// We use 0xFF as sentinel value for invalid indexes.
var __obf_5a65ade5e75208e0 = [...]byte{
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
func (__obf_b6d2ddd191a19ef1 *ULID) UnmarshalText(__obf_3b548eea5a5f82ce []byte) error {
	return __obf_9e258196bd35a28d(__obf_3b548eea5a5f82ce, false, __obf_b6d2ddd191a19ef1)
}

// Time returns the Unix time in milliseconds encoded in the ULID.
// Use the top level Time function to convert the returned value to
// a time.Time.
func (__obf_b6d2ddd191a19ef1 ULID) Time() uint64 {
	return uint64(__obf_b6d2ddd191a19ef1[5]) | uint64(__obf_b6d2ddd191a19ef1[4])<<8 |
		uint64(__obf_b6d2ddd191a19ef1[3])<<16 | uint64(__obf_b6d2ddd191a19ef1[2])<<24 |
		uint64(__obf_b6d2ddd191a19ef1[1])<<32 | uint64(__obf_b6d2ddd191a19ef1[0])<<40
}

// maxTime is the maximum Unix time in milliseconds that can be
// represented in an ULID.
var __obf_22cdee49b4638678 = ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}.Time()

// MaxTime returns the maximum Unix time in milliseconds that
// can be encoded in an ULID.
func MaxTime() uint64 { return __obf_22cdee49b4638678 }

// Now is a convenience function that returns the current
// UTC time in Unix milliseconds. Equivalent to:
//
//	Timestamp(time.Now().UTC())
func Now() uint64 { return Timestamp(time.Now().UTC()) }

// Timestamp converts a time.Time to Unix milliseconds.
//
// Because of the way ULID stores time, times from the year
// 10889 produces undefined results.
func Timestamp(__obf_8bbd0f9c7d55c1cc time.Time) uint64 {
	return uint64(__obf_8bbd0f9c7d55c1cc.Unix())*1000 +
		uint64(__obf_8bbd0f9c7d55c1cc.Nanosecond()/int(time.Millisecond))
}

// Time converts Unix milliseconds in the format
// returned by the Timestamp function to a time.Time.
func Time(__obf_a51f50affa0720aa uint64) time.Time {
	__obf_3cfeefb840f75d7a := int64(__obf_a51f50affa0720aa / 1e3)
	__obf_93a1f12a05f18ef2 := int64((__obf_a51f50affa0720aa % 1e3) * 1e6)
	return time.Unix(__obf_3cfeefb840f75d7a, __obf_93a1f12a05f18ef2)
}

// SetTime sets the time component of the ULID to the given Unix time
// in milliseconds.
func (__obf_b6d2ddd191a19ef1 *ULID) SetTime(__obf_a51f50affa0720aa uint64) error {
	if __obf_a51f50affa0720aa > __obf_22cdee49b4638678 {
		return ErrBigTime
	}

	(*__obf_b6d2ddd191a19ef1)[0] = byte(__obf_a51f50affa0720aa >> 40)
	(*__obf_b6d2ddd191a19ef1)[1] = byte(__obf_a51f50affa0720aa >> 32)
	(*__obf_b6d2ddd191a19ef1)[2] = byte(__obf_a51f50affa0720aa >> 24)
	(*__obf_b6d2ddd191a19ef1)[3] = byte(__obf_a51f50affa0720aa >> 16)
	(*__obf_b6d2ddd191a19ef1)[4] = byte(__obf_a51f50affa0720aa >> 8)
	(*__obf_b6d2ddd191a19ef1)[5] = byte(__obf_a51f50affa0720aa)

	return nil
}

// Entropy returns the entropy from the ULID.
func (__obf_b6d2ddd191a19ef1 ULID) Entropy() []byte {
	__obf_6d1c9d77673ade75 := make([]byte, 10)
	copy(__obf_6d1c9d77673ade75, __obf_b6d2ddd191a19ef1[6:])
	return __obf_6d1c9d77673ade75
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (__obf_b6d2ddd191a19ef1 *ULID) SetEntropy(__obf_6d1c9d77673ade75 []byte) error {
	if len(__obf_6d1c9d77673ade75) != 10 {
		return ErrDataSize
	}

	copy((*__obf_b6d2ddd191a19ef1)[6:], __obf_6d1c9d77673ade75)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (__obf_b6d2ddd191a19ef1 ULID) Compare(__obf_127840d41413faf8 ULID) int {
	return bytes.Compare(__obf_b6d2ddd191a19ef1[:], __obf_127840d41413faf8[:])
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (__obf_b6d2ddd191a19ef1 *ULID) Scan(__obf_01204dd49c46ed5c any) error {
	switch __obf_cf72d1dd2ebb40bf := __obf_01204dd49c46ed5c.(type) {
	case nil:
		return nil
	case string:
		return __obf_b6d2ddd191a19ef1.UnmarshalText([]byte(__obf_cf72d1dd2ebb40bf))
	case []byte:
		return __obf_b6d2ddd191a19ef1.UnmarshalBinary(__obf_cf72d1dd2ebb40bf)
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
func (__obf_b6d2ddd191a19ef1 ULID) Value() (driver.Value, error) {
	return __obf_b6d2ddd191a19ef1.MarshalBinary()
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
func Monotonic(__obf_7fc7e7bc83026231 io.Reader, __obf_aa4ca6f6e152a9bb uint64) *MonotonicEntropy {
	__obf_38d1dfd4bb39d54e := MonotonicEntropy{
		Reader:                 bufio.NewReader(__obf_7fc7e7bc83026231),
		__obf_aa4ca6f6e152a9bb: __obf_aa4ca6f6e152a9bb,
	}

	if __obf_38d1dfd4bb39d54e.__obf_aa4ca6f6e152a9bb == 0 {
		__obf_38d1dfd4bb39d54e.__obf_aa4ca6f6e152a9bb = math.MaxUint32
	}

	if __obf_be567146b6a23b53, __obf_a3279a70b6b813b3 := __obf_7fc7e7bc83026231.(*rand.Rand); __obf_a3279a70b6b813b3 {
		__obf_38d1dfd4bb39d54e.__obf_be567146b6a23b53 = __obf_be567146b6a23b53
	}

	return &__obf_38d1dfd4bb39d54e
}

// MonotonicEntropy is an opaque type that provides monotonic entropy.
type MonotonicEntropy struct {
	io.Reader
	__obf_a51f50affa0720aa uint64
	__obf_aa4ca6f6e152a9bb uint64
	__obf_7fc7e7bc83026231 __obf_0d7415f175403a34
	rand                   [8]byte
	__obf_be567146b6a23b53 *rand.Rand
}

// MonotonicRead implements the MonotonicReader interface.
func (__obf_38d1dfd4bb39d54e *MonotonicEntropy) MonotonicRead(__obf_a51f50affa0720aa uint64, __obf_7fc7e7bc83026231 []byte) (__obf_1082dd6e56192e3a error) {
	if !__obf_38d1dfd4bb39d54e.__obf_7fc7e7bc83026231.IsZero() && __obf_38d1dfd4bb39d54e.__obf_a51f50affa0720aa == __obf_a51f50affa0720aa {
		__obf_1082dd6e56192e3a = __obf_38d1dfd4bb39d54e.__obf_d5ff09d176130f8d()
		__obf_38d1dfd4bb39d54e.__obf_7fc7e7bc83026231.AppendTo(__obf_7fc7e7bc83026231)
	} else if _, __obf_1082dd6e56192e3a = io.ReadFull(__obf_38d1dfd4bb39d54e.Reader, __obf_7fc7e7bc83026231); __obf_1082dd6e56192e3a == nil {
		__obf_38d1dfd4bb39d54e.__obf_a51f50affa0720aa = __obf_a51f50affa0720aa
		__obf_38d1dfd4bb39d54e.__obf_7fc7e7bc83026231.SetBytes(__obf_7fc7e7bc83026231)
	}
	return __obf_1082dd6e56192e3a
}

// increment the previous entropy number with a random number
// of up to m.inc (inclusive).
func (__obf_38d1dfd4bb39d54e *MonotonicEntropy) __obf_d5ff09d176130f8d() error {
	if __obf_aa4ca6f6e152a9bb, __obf_1082dd6e56192e3a := __obf_38d1dfd4bb39d54e.__obf_89b9433e8f830ddf(); __obf_1082dd6e56192e3a != nil {
		return __obf_1082dd6e56192e3a
	} else if __obf_38d1dfd4bb39d54e.__obf_7fc7e7bc83026231.Add(__obf_aa4ca6f6e152a9bb) {
		return ErrMonotonicOverflow
	}
	return nil
}

// random returns a uniform random value in [1, m.inc), reading entropy
// from m.Reader. When m.inc == 0 || m.inc == 1, it returns 1.
// Adapted from: https://golang.org/pkg/crypto/rand/#Int
func (__obf_38d1dfd4bb39d54e *MonotonicEntropy) __obf_89b9433e8f830ddf() (__obf_aa4ca6f6e152a9bb uint64, __obf_1082dd6e56192e3a error) {
	if __obf_38d1dfd4bb39d54e.__obf_aa4ca6f6e152a9bb <= 1 {
		return 1, nil
	}

	// Fast path for using a underlying rand.Rand directly.
	if __obf_38d1dfd4bb39d54e.__obf_be567146b6a23b53 != nil {
		// Range: [1, m.inc)
		return 1 + uint64(__obf_38d1dfd4bb39d54e.__obf_be567146b6a23b53.Int63n(int64(__obf_38d1dfd4bb39d54e.__obf_aa4ca6f6e152a9bb))), nil
	}

	// bitLen is the maximum bit length needed to encode a value < m.inc.
	__obf_98263cd28c33156c := bits.Len64(__obf_38d1dfd4bb39d54e.__obf_aa4ca6f6e152a9bb)

	// byteLen is the maximum byte length needed to encode a value < m.inc.
	__obf_961f116a27dfda69 := uint(__obf_98263cd28c33156c+7) / 8

	// msbitLen is the number of bits in the most significant byte of m.inc-1.
	__obf_c1c7c622f77a89a3 := uint(__obf_98263cd28c33156c % 8)
	if __obf_c1c7c622f77a89a3 == 0 {
		__obf_c1c7c622f77a89a3 = 8
	}

	for __obf_aa4ca6f6e152a9bb == 0 || __obf_aa4ca6f6e152a9bb >= __obf_38d1dfd4bb39d54e.__obf_aa4ca6f6e152a9bb {
		if _, __obf_1082dd6e56192e3a = io.ReadFull(__obf_38d1dfd4bb39d54e.Reader, __obf_38d1dfd4bb39d54e.rand[:__obf_961f116a27dfda69]); __obf_1082dd6e56192e3a != nil {
			return 0, __obf_1082dd6e56192e3a
		}

		// Clear bits in the first byte to increase the probability
		// that the candidate is < m.inc.
		__obf_38d1dfd4bb39d54e.rand[0] &= uint8(int(1<<__obf_c1c7c622f77a89a3) - 1)

		// Convert the read bytes into an uint64 with byteLen
		// Optimized unrolled loop.
		switch __obf_961f116a27dfda69 {
		case 1:
			__obf_aa4ca6f6e152a9bb = uint64(__obf_38d1dfd4bb39d54e.rand[0])
		case 2:
			__obf_aa4ca6f6e152a9bb = uint64(binary.LittleEndian.Uint16(__obf_38d1dfd4bb39d54e.rand[:2]))
		case 3, 4:
			__obf_aa4ca6f6e152a9bb = uint64(binary.LittleEndian.Uint32(__obf_38d1dfd4bb39d54e.rand[:4]))
		case 5, 6, 7, 8:
			__obf_aa4ca6f6e152a9bb = uint64(binary.LittleEndian.Uint64(__obf_38d1dfd4bb39d54e.rand[:8]))
		}
	}

	// Range: [1, m.inc)
	return 1 + __obf_aa4ca6f6e152a9bb, nil
}

type __obf_0d7415f175403a34 struct {
	Hi uint16
	Lo uint64
}

func (__obf_6a86d57e2bf1634f *__obf_0d7415f175403a34) SetBytes(__obf_68e5fd3bec3c113a []byte) {
	__obf_6a86d57e2bf1634f.Hi = binary.BigEndian.Uint16(__obf_68e5fd3bec3c113a[:2])
	__obf_6a86d57e2bf1634f.Lo = binary.BigEndian.Uint64(__obf_68e5fd3bec3c113a[2:])
}

func (__obf_6a86d57e2bf1634f *__obf_0d7415f175403a34) AppendTo(__obf_68e5fd3bec3c113a []byte) {
	binary.BigEndian.PutUint16(__obf_68e5fd3bec3c113a[:2], __obf_6a86d57e2bf1634f.Hi)
	binary.BigEndian.PutUint64(__obf_68e5fd3bec3c113a[2:], __obf_6a86d57e2bf1634f.Lo)
}

func (__obf_6a86d57e2bf1634f *__obf_0d7415f175403a34) Add(__obf_20d0dbd4aeeff396 uint64) (__obf_dd90dd4128137061 bool) {
	__obf_32ff16f0290e1981, __obf_c3eb9dbd241a94e9 := __obf_6a86d57e2bf1634f.Lo, __obf_6a86d57e2bf1634f.Hi
	if __obf_6a86d57e2bf1634f.Lo += __obf_20d0dbd4aeeff396; __obf_6a86d57e2bf1634f.Lo < __obf_32ff16f0290e1981 {
		__obf_6a86d57e2bf1634f.Hi++
	}
	return __obf_6a86d57e2bf1634f.Hi < __obf_c3eb9dbd241a94e9
}

func (__obf_6a86d57e2bf1634f __obf_0d7415f175403a34) IsZero() bool {
	return __obf_6a86d57e2bf1634f.Hi == 0 && __obf_6a86d57e2bf1634f.Lo == 0
}
