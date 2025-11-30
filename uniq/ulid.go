package __obf_a51a64e21f311927

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
	MonotonicRead(__obf_bcfd4e83be51ceb0 uint64, __obf_e16341475febf1e2 []byte) error
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
func New(__obf_bcfd4e83be51ceb0 uint64, __obf_039024cb6088a1c4 io.Reader) (__obf_4c112818117e5211 ULID, __obf_886df8acc7f50b76 error) {
	if __obf_886df8acc7f50b76 = __obf_4c112818117e5211.SetTime(__obf_bcfd4e83be51ceb0); __obf_886df8acc7f50b76 != nil {
		return __obf_4c112818117e5211, __obf_886df8acc7f50b76
	}

	switch __obf_2b3a8603fb1a3e4d := __obf_039024cb6088a1c4.(type) {
	case nil:
		return __obf_4c112818117e5211, __obf_886df8acc7f50b76
	case MonotonicReader:
		__obf_886df8acc7f50b76 = __obf_2b3a8603fb1a3e4d.MonotonicRead(__obf_bcfd4e83be51ceb0, __obf_4c112818117e5211[6:])
	default:
		_, __obf_886df8acc7f50b76 = io.ReadFull(__obf_2b3a8603fb1a3e4d, __obf_4c112818117e5211[6:])
	}

	return __obf_4c112818117e5211,

		// MustNew is a convenience function equivalent to New that panics on failure
		// instead of returning an error.
		__obf_886df8acc7f50b76
}

func MustNew(__obf_bcfd4e83be51ceb0 uint64, __obf_039024cb6088a1c4 io.Reader) ULID {
	__obf_4c112818117e5211, __obf_886df8acc7f50b76 := New(__obf_bcfd4e83be51ceb0, __obf_039024cb6088a1c4)
	if __obf_886df8acc7f50b76 != nil {
		panic(__obf_886df8acc7f50b76)
	}
	return __obf_4c112818117e5211
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(__obf_177446b81f659d7f string) (__obf_4c112818117e5211 ULID, __obf_886df8acc7f50b76 error) {
	return __obf_4c112818117e5211, __obf_e5cd0ae1df7f7a23([]byte(__obf_177446b81f659d7f), false, &__obf_4c112818117e5211)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(__obf_177446b81f659d7f string) (__obf_4c112818117e5211 ULID, __obf_886df8acc7f50b76 error) {
	return __obf_4c112818117e5211, __obf_e5cd0ae1df7f7a23([]byte(__obf_177446b81f659d7f), true, &__obf_4c112818117e5211)
}

func __obf_e5cd0ae1df7f7a23(__obf_8c882dce5ca83eeb []byte, __obf_e0c244445314d9ba bool, __obf_4c112818117e5211 *ULID) error {
	// Check if a base32 encoded ULID is the right length.
	if len(__obf_8c882dce5ca83eeb) != EncodedSize {
		return ErrDataSize
	}

	// Check if all the characters in a base32 encoded ULID are part of the
	// expected base32 character set.
	if __obf_e0c244445314d9ba &&
		(__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[0]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[1]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[2]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[3]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[4]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[5]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[6]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[7]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[8]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[9]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[10]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[11]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[12]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[13]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[14]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[15]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[16]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[17]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[18]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[19]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[20]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[21]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[22]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[23]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[24]] == 0xFF || __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[25]] == 0xFF) {
		return ErrInvalidCharacters
	}

	// Check if the first character in a base32 encoded ULID will overflow. This
	// happens because the base32 representation encodes 130 bits, while the
	// ULID is only 128 bits.
	//
	// See https://github.com/oklog/ulid/issues/9 for details.
	if __obf_8c882dce5ca83eeb[0] > '7' {
		return ErrOverflow
	}

	// Use an optimized unrolled loop (from https://github.com/RobThree/NUlid)
	// to decode a base32 ULID.

	// 6 bytes timestamp (48 bits)
	(*__obf_4c112818117e5211)[0] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[0]] << 5) | __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[1]]
	(*__obf_4c112818117e5211)[1] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[2]] << 3) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[3]] >> 2)
	(*__obf_4c112818117e5211)[2] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[3]] << 6) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[4]] << 1) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[5]] >> 4)
	(*__obf_4c112818117e5211)[3] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[5]] << 4) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[6]] >> 1)
	(*__obf_4c112818117e5211)[4] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[6]] << 7) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[7]] << 2) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[8]] >> 3)
	(*__obf_4c112818117e5211)[5] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[8]] << 5) | __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[9]]

	// 10 bytes of entropy (80 bits)
	(*__obf_4c112818117e5211)[6] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[10]] << 3) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[11]] >> 2)
	(*__obf_4c112818117e5211)[7] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[11]] << 6) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[12]] << 1) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[13]] >> 4)
	(*__obf_4c112818117e5211)[8] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[13]] << 4) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[14]] >> 1)
	(*__obf_4c112818117e5211)[9] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[14]] << 7) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[15]] << 2) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[16]] >> 3)
	(*__obf_4c112818117e5211)[10] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[16]] << 5) | __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[17]]
	(*__obf_4c112818117e5211)[11] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[18]] << 3) | __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[19]]>>2
	(*__obf_4c112818117e5211)[12] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[19]] << 6) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[20]] << 1) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[21]] >> 4)
	(*__obf_4c112818117e5211)[13] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[21]] << 4) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[22]] >> 1)
	(*__obf_4c112818117e5211)[14] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[22]] << 7) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[23]] << 2) | (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[24]] >> 3)
	(*__obf_4c112818117e5211)[15] = (__obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[24]] << 5) | __obf_0bf1be72e921a9de[__obf_8c882dce5ca83eeb[25]]

	return nil
}

// MustParse is a convenience function equivalent to Parse that panics on failure
// instead of returning an error.
func MustParse(__obf_177446b81f659d7f string) ULID {
	__obf_4c112818117e5211, __obf_886df8acc7f50b76 := Parse(__obf_177446b81f659d7f)
	if __obf_886df8acc7f50b76 != nil {
		panic(__obf_886df8acc7f50b76)
	}
	return __obf_4c112818117e5211
}

// MustParseStrict is a convenience function equivalent to ParseStrict that
// panics on failure instead of returning an error.
func MustParseStrict(__obf_177446b81f659d7f string) ULID {
	__obf_4c112818117e5211, __obf_886df8acc7f50b76 := ParseStrict(__obf_177446b81f659d7f)
	if __obf_886df8acc7f50b76 != nil {
		panic(__obf_886df8acc7f50b76)
	}
	return __obf_4c112818117e5211
}

// Bytes returns bytes slice representation of ULID.
func (__obf_d5f5a3a4687208bc ULID) Bytes() []byte {
	return __obf_d5f5a3a4687208bc[:]
}

// String returns a lexicographically sortable string encoded ULID
// (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy
func (__obf_4c112818117e5211 ULID) String() string {
	__obf_177446b81f659d7f := make([]byte, EncodedSize)
	_ = __obf_4c112818117e5211.MarshalTextTo(__obf_177446b81f659d7f)
	return string(__obf_177446b81f659d7f)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (__obf_4c112818117e5211 ULID) MarshalBinary() ([]byte, error) {
	__obf_177446b81f659d7f := make([]byte, len(__obf_4c112818117e5211))
	return __obf_177446b81f659d7f, __obf_4c112818117e5211.MarshalBinaryTo(__obf_177446b81f659d7f)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (__obf_4c112818117e5211 ULID) MarshalBinaryTo(__obf_d166c1de1b2ffc41 []byte) error {
	if len(__obf_d166c1de1b2ffc41) != len(__obf_4c112818117e5211) {
		return ErrBufferSize
	}

	copy(__obf_d166c1de1b2ffc41, __obf_4c112818117e5211[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (__obf_4c112818117e5211 *ULID) UnmarshalBinary(__obf_82f1b077b939fcec []byte) error {
	if len(__obf_82f1b077b939fcec) != len(*__obf_4c112818117e5211) {
		return ErrDataSize
	}

	copy((*__obf_4c112818117e5211)[:], __obf_82f1b077b939fcec)
	return nil
}

// Encoding is the base 32 encoding alphabet used in ULID strings.
const Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (__obf_4c112818117e5211 ULID) MarshalText() ([]byte, error) {
	__obf_177446b81f659d7f := make([]byte, EncodedSize)
	return __obf_177446b81f659d7f, __obf_4c112818117e5211.MarshalTextTo(__obf_177446b81f659d7f)
}

// MarshalTextTo writes the ULID as a string to the given buffer.
// ErrBufferSize is returned when the len(dst) != 26.
func (__obf_4c112818117e5211 ULID) MarshalTextTo(__obf_d166c1de1b2ffc41 []byte) error {
	// Optimized unrolled loop ahead.
	// From https://github.com/RobThree/NUlid

	if len(__obf_d166c1de1b2ffc41) != EncodedSize {
		return ErrBufferSize
	}
	__obf_d166c1de1b2ffc41[ // 10 byte timestamp
	0] = Encoding[(__obf_4c112818117e5211[0]&224)>>5]
	__obf_d166c1de1b2ffc41[1] = Encoding[__obf_4c112818117e5211[0]&31]
	__obf_d166c1de1b2ffc41[2] = Encoding[(__obf_4c112818117e5211[1]&248)>>3]
	__obf_d166c1de1b2ffc41[3] = Encoding[((__obf_4c112818117e5211[1]&7)<<2)|((__obf_4c112818117e5211[2]&192)>>6)]
	__obf_d166c1de1b2ffc41[4] = Encoding[(__obf_4c112818117e5211[2]&62)>>1]
	__obf_d166c1de1b2ffc41[5] = Encoding[((__obf_4c112818117e5211[2]&1)<<4)|((__obf_4c112818117e5211[3]&240)>>4)]
	__obf_d166c1de1b2ffc41[6] = Encoding[((__obf_4c112818117e5211[3]&15)<<1)|((__obf_4c112818117e5211[4]&128)>>7)]
	__obf_d166c1de1b2ffc41[7] = Encoding[(__obf_4c112818117e5211[4]&124)>>2]
	__obf_d166c1de1b2ffc41[8] = Encoding[((__obf_4c112818117e5211[4]&3)<<3)|((__obf_4c112818117e5211[5]&224)>>5)]
	__obf_d166c1de1b2ffc41[9] = Encoding[__obf_4c112818117e5211[5]&31]
	__obf_d166c1de1b2ffc41[ // 16 bytes of entropy
	10] = Encoding[(__obf_4c112818117e5211[6]&248)>>3]
	__obf_d166c1de1b2ffc41[11] = Encoding[((__obf_4c112818117e5211[6]&7)<<2)|((__obf_4c112818117e5211[7]&192)>>6)]
	__obf_d166c1de1b2ffc41[12] = Encoding[(__obf_4c112818117e5211[7]&62)>>1]
	__obf_d166c1de1b2ffc41[13] = Encoding[((__obf_4c112818117e5211[7]&1)<<4)|((__obf_4c112818117e5211[8]&240)>>4)]
	__obf_d166c1de1b2ffc41[14] = Encoding[((__obf_4c112818117e5211[8]&15)<<1)|((__obf_4c112818117e5211[9]&128)>>7)]
	__obf_d166c1de1b2ffc41[15] = Encoding[(__obf_4c112818117e5211[9]&124)>>2]
	__obf_d166c1de1b2ffc41[16] = Encoding[((__obf_4c112818117e5211[9]&3)<<3)|((__obf_4c112818117e5211[10]&224)>>5)]
	__obf_d166c1de1b2ffc41[17] = Encoding[__obf_4c112818117e5211[10]&31]
	__obf_d166c1de1b2ffc41[18] = Encoding[(__obf_4c112818117e5211[11]&248)>>3]
	__obf_d166c1de1b2ffc41[19] = Encoding[((__obf_4c112818117e5211[11]&7)<<2)|((__obf_4c112818117e5211[12]&192)>>6)]
	__obf_d166c1de1b2ffc41[20] = Encoding[(__obf_4c112818117e5211[12]&62)>>1]
	__obf_d166c1de1b2ffc41[21] = Encoding[((__obf_4c112818117e5211[12]&1)<<4)|((__obf_4c112818117e5211[13]&240)>>4)]
	__obf_d166c1de1b2ffc41[22] = Encoding[((__obf_4c112818117e5211[13]&15)<<1)|((__obf_4c112818117e5211[14]&128)>>7)]
	__obf_d166c1de1b2ffc41[23] = Encoding[(__obf_4c112818117e5211[14]&124)>>2]
	__obf_d166c1de1b2ffc41[24] = Encoding[((__obf_4c112818117e5211[14]&3)<<3)|((__obf_4c112818117e5211[15]&224)>>5)]
	__obf_d166c1de1b2ffc41[25] = Encoding[__obf_4c112818117e5211[15]&31]

	return nil
}

// Byte to index table for O(1) lookups when unmarshaling.
// We use 0xFF as sentinel value for invalid indexes.
var __obf_0bf1be72e921a9de = [...]byte{
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
func (__obf_4c112818117e5211 *ULID) UnmarshalText(__obf_8c882dce5ca83eeb []byte) error {
	return __obf_e5cd0ae1df7f7a23(__obf_8c882dce5ca83eeb, false, __obf_4c112818117e5211)
}

// Time returns the Unix time in milliseconds encoded in the ULID.
// Use the top level Time function to convert the returned value to
// a time.Time.
func (__obf_4c112818117e5211 ULID) Time() uint64 {
	return uint64(__obf_4c112818117e5211[5]) | uint64(__obf_4c112818117e5211[4])<<8 |
		uint64(__obf_4c112818117e5211[3])<<16 | uint64(__obf_4c112818117e5211[2])<<24 |
		uint64(__obf_4c112818117e5211[1])<<32 | uint64(__obf_4c112818117e5211[0])<<40
}

// maxTime is the maximum Unix time in milliseconds that can be
// represented in an ULID.
var __obf_50c93da969d97251 = ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}.Time()

// MaxTime returns the maximum Unix time in milliseconds that
// can be encoded in an ULID.
func MaxTime() uint64 { return __obf_50c93da969d97251 }

// Now is a convenience function that returns the current
// UTC time in Unix milliseconds. Equivalent to:
//
//	Timestamp(time.Now().UTC())
func Now() uint64 { return Timestamp(time.Now().UTC()) }

// Timestamp converts a time.Time to Unix milliseconds.
//
// Because of the way ULID stores time, times from the year
// 10889 produces undefined results.
func Timestamp(__obf_c5969bc7b424c7a8 time.Time) uint64 {
	return uint64(__obf_c5969bc7b424c7a8.Unix())*1000 +
		uint64(__obf_c5969bc7b424c7a8.Nanosecond()/int(time.Millisecond))
}

// Time converts Unix milliseconds in the format
// returned by the Timestamp function to a time.Time.
func Time(__obf_bcfd4e83be51ceb0 uint64) time.Time {
	__obf_c51d45361c8f2458 := int64(__obf_bcfd4e83be51ceb0 / 1e3)
	__obf_3ec6d76544365bf7 := int64((__obf_bcfd4e83be51ceb0 % 1e3) * 1e6)
	return time.Unix(__obf_c51d45361c8f2458,

		// SetTime sets the time component of the ULID to the given Unix time
		// in milliseconds.
		__obf_3ec6d76544365bf7)
}

func (__obf_4c112818117e5211 *ULID) SetTime(__obf_bcfd4e83be51ceb0 uint64) error {
	if __obf_bcfd4e83be51ceb0 > __obf_50c93da969d97251 {
		return ErrBigTime
	}

	(*__obf_4c112818117e5211)[0] = byte(__obf_bcfd4e83be51ceb0 >> 40)
	(*__obf_4c112818117e5211)[1] = byte(__obf_bcfd4e83be51ceb0 >> 32)
	(*__obf_4c112818117e5211)[2] = byte(__obf_bcfd4e83be51ceb0 >> 24)
	(*__obf_4c112818117e5211)[3] = byte(__obf_bcfd4e83be51ceb0 >> 16)
	(*__obf_4c112818117e5211)[4] = byte(__obf_bcfd4e83be51ceb0 >> 8)
	(*__obf_4c112818117e5211)[5] = byte(__obf_bcfd4e83be51ceb0)

	return nil
}

// Entropy returns the entropy from the ULID.
func (__obf_4c112818117e5211 ULID) Entropy() []byte {
	__obf_2b3a8603fb1a3e4d := make([]byte, 10)
	copy(__obf_2b3a8603fb1a3e4d, __obf_4c112818117e5211[6:])
	return __obf_2b3a8603fb1a3e4d
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (__obf_4c112818117e5211 *ULID) SetEntropy(__obf_2b3a8603fb1a3e4d []byte) error {
	if len(__obf_2b3a8603fb1a3e4d) != 10 {
		return ErrDataSize
	}

	copy((*__obf_4c112818117e5211)[6:], __obf_2b3a8603fb1a3e4d)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (__obf_4c112818117e5211 ULID) Compare(__obf_4ab294a4b121dfd9 ULID) int {
	return bytes.Compare(__obf_4c112818117e5211[:], __obf_4ab294a4b121dfd9[:])
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (__obf_4c112818117e5211 *ULID) Scan(__obf_d5b52944909f5a29 any) error {
	switch __obf_137cb83870c72e80 := __obf_d5b52944909f5a29.(type) {
	case nil:
		return nil
	case string:
		return __obf_4c112818117e5211.UnmarshalText([]byte(__obf_137cb83870c72e80))
	case []byte:
		return __obf_4c112818117e5211.UnmarshalBinary(__obf_137cb83870c72e80)
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
func (__obf_4c112818117e5211 ULID) Value() (driver.Value, error) {
	return __obf_4c112818117e5211.MarshalBinary()
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
func Monotonic(__obf_039024cb6088a1c4 io.Reader, __obf_c514709dcb281e4f uint64) *MonotonicEntropy {
	__obf_5c83ea5b977f90b8 := MonotonicEntropy{
		Reader: bufio.NewReader(__obf_039024cb6088a1c4), __obf_c514709dcb281e4f: __obf_c514709dcb281e4f,
	}

	if __obf_5c83ea5b977f90b8.__obf_c514709dcb281e4f == 0 {
		__obf_5c83ea5b977f90b8.__obf_c514709dcb281e4f = math.MaxUint32
	}

	if __obf_2f0dd74324bcccf3, __obf_80eec96aa83c2efa := __obf_039024cb6088a1c4.(*rand.Rand); __obf_80eec96aa83c2efa {
		__obf_5c83ea5b977f90b8.__obf_2f0dd74324bcccf3 = __obf_2f0dd74324bcccf3
	}

	return &__obf_5c83ea5b977f90b8
}

// MonotonicEntropy is an opaque type that provides monotonic entropy.
type MonotonicEntropy struct {
	io.Reader
	__obf_bcfd4e83be51ceb0 uint64
	__obf_c514709dcb281e4f uint64
	__obf_039024cb6088a1c4 __obf_041e1e5a50f97db5

	rand                   [8]byte
	__obf_2f0dd74324bcccf3 *rand.Rand
}

// MonotonicRead implements the MonotonicReader interface.
func (__obf_5c83ea5b977f90b8 *MonotonicEntropy) MonotonicRead(__obf_bcfd4e83be51ceb0 uint64, __obf_039024cb6088a1c4 []byte) (__obf_886df8acc7f50b76 error) {
	if !__obf_5c83ea5b977f90b8.__obf_039024cb6088a1c4.IsZero() && __obf_5c83ea5b977f90b8.__obf_bcfd4e83be51ceb0 == __obf_bcfd4e83be51ceb0 {
		__obf_886df8acc7f50b76 = __obf_5c83ea5b977f90b8.__obf_fa71afed06c562e4()
		__obf_5c83ea5b977f90b8.__obf_039024cb6088a1c4.
			AppendTo(__obf_039024cb6088a1c4)
	} else if _, __obf_886df8acc7f50b76 = io.ReadFull(__obf_5c83ea5b977f90b8.Reader, __obf_039024cb6088a1c4); __obf_886df8acc7f50b76 == nil {
		__obf_5c83ea5b977f90b8.__obf_bcfd4e83be51ceb0 = __obf_bcfd4e83be51ceb0
		__obf_5c83ea5b977f90b8.__obf_039024cb6088a1c4.
			SetBytes(__obf_039024cb6088a1c4)
	}
	return __obf_886df8acc7f50b76
}

// increment the previous entropy number with a random number
// of up to m.inc (inclusive).
func (__obf_5c83ea5b977f90b8 *MonotonicEntropy) __obf_fa71afed06c562e4() error {
	if __obf_c514709dcb281e4f, __obf_886df8acc7f50b76 := __obf_5c83ea5b977f90b8.__obf_aaaca78a945f310c(); __obf_886df8acc7f50b76 != nil {
		return __obf_886df8acc7f50b76
	} else if __obf_5c83ea5b977f90b8.__obf_039024cb6088a1c4.Add(__obf_c514709dcb281e4f) {
		return ErrMonotonicOverflow
	}
	return nil
}

// random returns a uniform random value in [1, m.inc), reading entropy
// from m.Reader. When m.inc == 0 || m.inc == 1, it returns 1.
// Adapted from: https://golang.org/pkg/crypto/rand/#Int
func (__obf_5c83ea5b977f90b8 *MonotonicEntropy) __obf_aaaca78a945f310c() (__obf_c514709dcb281e4f uint64, __obf_886df8acc7f50b76 error) {
	if __obf_5c83ea5b977f90b8.__obf_c514709dcb281e4f <= 1 {
		return 1, nil
	}

	// Fast path for using a underlying rand.Rand directly.
	if __obf_5c83ea5b977f90b8.
		// Range: [1, m.inc)
		__obf_2f0dd74324bcccf3 != nil {

		return 1 + uint64(__obf_5c83ea5b977f90b8.__obf_2f0dd74324bcccf3.Int63n(int64(__obf_5c83ea5b977f90b8.

			// bitLen is the maximum bit length needed to encode a value < m.inc.
			__obf_c514709dcb281e4f))), nil
	}
	__obf_a363c8ab67490f8d := bits.Len64(__obf_5c83ea5b977f90b8.

		// byteLen is the maximum byte length needed to encode a value < m.inc.
		__obf_c514709dcb281e4f)
	__obf_8efaba1c990e3912 := uint(__obf_a363c8ab67490f8d+7) / 8
	__obf_f343e8a2689d1ac4 := // msbitLen is the number of bits in the most significant byte of m.inc-1.
		uint(__obf_a363c8ab67490f8d % 8)
	if __obf_f343e8a2689d1ac4 == 0 {
		__obf_f343e8a2689d1ac4 = 8
	}

	for __obf_c514709dcb281e4f == 0 || __obf_c514709dcb281e4f >= __obf_5c83ea5b977f90b8.__obf_c514709dcb281e4f {
		if _, __obf_886df8acc7f50b76 = io.ReadFull(__obf_5c83ea5b977f90b8.Reader, __obf_5c83ea5b977f90b8.rand[:__obf_8efaba1c990e3912]); __obf_886df8acc7f50b76 != nil {
			return 0, __obf_886df8acc7f50b76
		}
		__obf_5c83ea5b977f90b8.

			// Clear bits in the first byte to increase the probability
			// that the candidate is < m.inc.
			rand[0] &= uint8(int(1<<__obf_f343e8a2689d1ac4) - 1)

		// Convert the read bytes into an uint64 with byteLen
		// Optimized unrolled loop.
		switch __obf_8efaba1c990e3912 {
		case 1:
			__obf_c514709dcb281e4f = uint64(__obf_5c83ea5b977f90b8.rand[0])
		case 2:
			__obf_c514709dcb281e4f = uint64(binary.LittleEndian.Uint16(__obf_5c83ea5b977f90b8.rand[:2]))
		case 3, 4:
			__obf_c514709dcb281e4f = uint64(binary.LittleEndian.Uint32(__obf_5c83ea5b977f90b8.rand[:4]))
		case 5, 6, 7, 8:
			__obf_c514709dcb281e4f = uint64(binary.LittleEndian.Uint64(__obf_5c83ea5b977f90b8.rand[:8]))
		}
	}

	// Range: [1, m.inc)
	return 1 + __obf_c514709dcb281e4f, nil
}

type __obf_041e1e5a50f97db5 struct {
	Hi uint16
	Lo uint64
}

func (__obf_d5f5a3a4687208bc *__obf_041e1e5a50f97db5) SetBytes(__obf_46b17ccab1d7b33b []byte) {
	__obf_d5f5a3a4687208bc.
		Hi = binary.BigEndian.Uint16(__obf_46b17ccab1d7b33b[:2])
	__obf_d5f5a3a4687208bc.
		Lo = binary.BigEndian.Uint64(__obf_46b17ccab1d7b33b[2:])
}

func (__obf_d5f5a3a4687208bc *__obf_041e1e5a50f97db5) AppendTo(__obf_46b17ccab1d7b33b []byte) {
	binary.BigEndian.PutUint16(__obf_46b17ccab1d7b33b[:2], __obf_d5f5a3a4687208bc.Hi)
	binary.BigEndian.PutUint64(__obf_46b17ccab1d7b33b[2:], __obf_d5f5a3a4687208bc.Lo)
}

func (__obf_d5f5a3a4687208bc *__obf_041e1e5a50f97db5) Add(__obf_9f6decc0e2d1ee43 uint64) (__obf_ff37a93d038c14a3 bool) {
	__obf_d999262136d1a9d1, __obf_803d6346d4a4a7b3 := __obf_d5f5a3a4687208bc.Lo, __obf_d5f5a3a4687208bc.Hi
	if __obf_d5f5a3a4687208bc.Lo += __obf_9f6decc0e2d1ee43; __obf_d5f5a3a4687208bc.Lo < __obf_d999262136d1a9d1 {
		__obf_d5f5a3a4687208bc.
			Hi++
	}
	return __obf_d5f5a3a4687208bc.Hi < __obf_803d6346d4a4a7b3
}

func (__obf_d5f5a3a4687208bc __obf_041e1e5a50f97db5) IsZero() bool {
	return __obf_d5f5a3a4687208bc.Hi == 0 && __obf_d5f5a3a4687208bc.Lo == 0
}
