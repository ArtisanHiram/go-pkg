package __obf_7913809aab6c8423

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
	MonotonicRead(__obf_efc47ff3f584b49b uint64, __obf_90c08d248ca839a3 []byte) error
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
func New(__obf_efc47ff3f584b49b uint64, __obf_67566d403e1ce10f io.Reader) (__obf_1dfe314328e20778 ULID, __obf_e62cd9e417a87ee7 error) {
	if __obf_e62cd9e417a87ee7 = __obf_1dfe314328e20778.SetTime(__obf_efc47ff3f584b49b); __obf_e62cd9e417a87ee7 != nil {
		return __obf_1dfe314328e20778, __obf_e62cd9e417a87ee7
	}

	switch __obf_e9c0741dc5032511 := __obf_67566d403e1ce10f.(type) {
	case nil:
		return __obf_1dfe314328e20778, __obf_e62cd9e417a87ee7
	case MonotonicReader:
		__obf_e62cd9e417a87ee7 = __obf_e9c0741dc5032511.MonotonicRead(__obf_efc47ff3f584b49b, __obf_1dfe314328e20778[6:])
	default:
		_, __obf_e62cd9e417a87ee7 = io.ReadFull(__obf_e9c0741dc5032511, __obf_1dfe314328e20778[6:])
	}

	return __obf_1dfe314328e20778, __obf_e62cd9e417a87ee7
}

// MustNew is a convenience function equivalent to New that panics on failure
// instead of returning an error.
func MustNew(__obf_efc47ff3f584b49b uint64, __obf_67566d403e1ce10f io.Reader) ULID {
	__obf_1dfe314328e20778, __obf_e62cd9e417a87ee7 := New(__obf_efc47ff3f584b49b, __obf_67566d403e1ce10f)
	if __obf_e62cd9e417a87ee7 != nil {
		panic(__obf_e62cd9e417a87ee7)
	}
	return __obf_1dfe314328e20778
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(__obf_bfd738fbd36ff3de string) (__obf_1dfe314328e20778 ULID, __obf_e62cd9e417a87ee7 error) {
	return __obf_1dfe314328e20778, __obf_46dce11e234e40b8([]byte(__obf_bfd738fbd36ff3de), false, &__obf_1dfe314328e20778)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(__obf_bfd738fbd36ff3de string) (__obf_1dfe314328e20778 ULID, __obf_e62cd9e417a87ee7 error) {
	return __obf_1dfe314328e20778, __obf_46dce11e234e40b8([]byte(__obf_bfd738fbd36ff3de), true, &__obf_1dfe314328e20778)
}

func __obf_46dce11e234e40b8(__obf_ce304ea72067f76f []byte, __obf_00f09f01e15c84f2 bool, __obf_1dfe314328e20778 *ULID) error {
	// Check if a base32 encoded ULID is the right length.
	if len(__obf_ce304ea72067f76f) != EncodedSize {
		return ErrDataSize
	}

	// Check if all the characters in a base32 encoded ULID are part of the
	// expected base32 character set.
	if __obf_00f09f01e15c84f2 &&
		(__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[0]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[1]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[2]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[3]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[4]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[5]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[6]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[7]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[8]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[9]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[10]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[11]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[12]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[13]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[14]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[15]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[16]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[17]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[18]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[19]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[20]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[21]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[22]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[23]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[24]] == 0xFF ||
			__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[25]] == 0xFF) {
		return ErrInvalidCharacters
	}

	// Check if the first character in a base32 encoded ULID will overflow. This
	// happens because the base32 representation encodes 130 bits, while the
	// ULID is only 128 bits.
	//
	// See https://github.com/oklog/ulid/issues/9 for details.
	if __obf_ce304ea72067f76f[0] > '7' {
		return ErrOverflow
	}

	// Use an optimized unrolled loop (from https://github.com/RobThree/NUlid)
	// to decode a base32 ULID.

	// 6 bytes timestamp (48 bits)
	(*__obf_1dfe314328e20778)[0] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[0]] << 5) | __obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[1]]
	(*__obf_1dfe314328e20778)[1] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[2]] << 3) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[3]] >> 2)
	(*__obf_1dfe314328e20778)[2] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[3]] << 6) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[4]] << 1) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[5]] >> 4)
	(*__obf_1dfe314328e20778)[3] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[5]] << 4) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[6]] >> 1)
	(*__obf_1dfe314328e20778)[4] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[6]] << 7) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[7]] << 2) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[8]] >> 3)
	(*__obf_1dfe314328e20778)[5] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[8]] << 5) | __obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[9]]

	// 10 bytes of entropy (80 bits)
	(*__obf_1dfe314328e20778)[6] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[10]] << 3) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[11]] >> 2)
	(*__obf_1dfe314328e20778)[7] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[11]] << 6) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[12]] << 1) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[13]] >> 4)
	(*__obf_1dfe314328e20778)[8] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[13]] << 4) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[14]] >> 1)
	(*__obf_1dfe314328e20778)[9] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[14]] << 7) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[15]] << 2) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[16]] >> 3)
	(*__obf_1dfe314328e20778)[10] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[16]] << 5) | __obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[17]]
	(*__obf_1dfe314328e20778)[11] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[18]] << 3) | __obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[19]]>>2
	(*__obf_1dfe314328e20778)[12] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[19]] << 6) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[20]] << 1) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[21]] >> 4)
	(*__obf_1dfe314328e20778)[13] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[21]] << 4) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[22]] >> 1)
	(*__obf_1dfe314328e20778)[14] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[22]] << 7) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[23]] << 2) | (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[24]] >> 3)
	(*__obf_1dfe314328e20778)[15] = (__obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[24]] << 5) | __obf_75ce8c07c15c5d3a[__obf_ce304ea72067f76f[25]]

	return nil
}

// MustParse is a convenience function equivalent to Parse that panics on failure
// instead of returning an error.
func MustParse(__obf_bfd738fbd36ff3de string) ULID {
	__obf_1dfe314328e20778, __obf_e62cd9e417a87ee7 := Parse(__obf_bfd738fbd36ff3de)
	if __obf_e62cd9e417a87ee7 != nil {
		panic(__obf_e62cd9e417a87ee7)
	}
	return __obf_1dfe314328e20778
}

// MustParseStrict is a convenience function equivalent to ParseStrict that
// panics on failure instead of returning an error.
func MustParseStrict(__obf_bfd738fbd36ff3de string) ULID {
	__obf_1dfe314328e20778, __obf_e62cd9e417a87ee7 := ParseStrict(__obf_bfd738fbd36ff3de)
	if __obf_e62cd9e417a87ee7 != nil {
		panic(__obf_e62cd9e417a87ee7)
	}
	return __obf_1dfe314328e20778
}

// Bytes returns bytes slice representation of ULID.
func (__obf_3d63dae1aea4bf71 ULID) Bytes() []byte {
	return __obf_3d63dae1aea4bf71[:]
}

// String returns a lexicographically sortable string encoded ULID
// (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy
func (__obf_1dfe314328e20778 ULID) String() string {
	__obf_bfd738fbd36ff3de := make([]byte, EncodedSize)
	_ = __obf_1dfe314328e20778.MarshalTextTo(__obf_bfd738fbd36ff3de)
	return string(__obf_bfd738fbd36ff3de)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (__obf_1dfe314328e20778 ULID) MarshalBinary() ([]byte, error) {
	__obf_bfd738fbd36ff3de := make([]byte, len(__obf_1dfe314328e20778))
	return __obf_bfd738fbd36ff3de, __obf_1dfe314328e20778.MarshalBinaryTo(__obf_bfd738fbd36ff3de)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (__obf_1dfe314328e20778 ULID) MarshalBinaryTo(__obf_8dabb278b0366623 []byte) error {
	if len(__obf_8dabb278b0366623) != len(__obf_1dfe314328e20778) {
		return ErrBufferSize
	}

	copy(__obf_8dabb278b0366623, __obf_1dfe314328e20778[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (__obf_1dfe314328e20778 *ULID) UnmarshalBinary(__obf_20b9c0d01b4df6a8 []byte) error {
	if len(__obf_20b9c0d01b4df6a8) != len(*__obf_1dfe314328e20778) {
		return ErrDataSize
	}

	copy((*__obf_1dfe314328e20778)[:], __obf_20b9c0d01b4df6a8)
	return nil
}

// Encoding is the base 32 encoding alphabet used in ULID strings.
const Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (__obf_1dfe314328e20778 ULID) MarshalText() ([]byte, error) {
	__obf_bfd738fbd36ff3de := make([]byte, EncodedSize)
	return __obf_bfd738fbd36ff3de, __obf_1dfe314328e20778.MarshalTextTo(__obf_bfd738fbd36ff3de)
}

// MarshalTextTo writes the ULID as a string to the given buffer.
// ErrBufferSize is returned when the len(dst) != 26.
func (__obf_1dfe314328e20778 ULID) MarshalTextTo(__obf_8dabb278b0366623 []byte) error {
	// Optimized unrolled loop ahead.
	// From https://github.com/RobThree/NUlid

	if len(__obf_8dabb278b0366623) != EncodedSize {
		return ErrBufferSize
	}

	// 10 byte timestamp
	__obf_8dabb278b0366623[0] = Encoding[(__obf_1dfe314328e20778[0]&224)>>5]
	__obf_8dabb278b0366623[1] = Encoding[__obf_1dfe314328e20778[0]&31]
	__obf_8dabb278b0366623[2] = Encoding[(__obf_1dfe314328e20778[1]&248)>>3]
	__obf_8dabb278b0366623[3] = Encoding[((__obf_1dfe314328e20778[1]&7)<<2)|((__obf_1dfe314328e20778[2]&192)>>6)]
	__obf_8dabb278b0366623[4] = Encoding[(__obf_1dfe314328e20778[2]&62)>>1]
	__obf_8dabb278b0366623[5] = Encoding[((__obf_1dfe314328e20778[2]&1)<<4)|((__obf_1dfe314328e20778[3]&240)>>4)]
	__obf_8dabb278b0366623[6] = Encoding[((__obf_1dfe314328e20778[3]&15)<<1)|((__obf_1dfe314328e20778[4]&128)>>7)]
	__obf_8dabb278b0366623[7] = Encoding[(__obf_1dfe314328e20778[4]&124)>>2]
	__obf_8dabb278b0366623[8] = Encoding[((__obf_1dfe314328e20778[4]&3)<<3)|((__obf_1dfe314328e20778[5]&224)>>5)]
	__obf_8dabb278b0366623[9] = Encoding[__obf_1dfe314328e20778[5]&31]

	// 16 bytes of entropy
	__obf_8dabb278b0366623[10] = Encoding[(__obf_1dfe314328e20778[6]&248)>>3]
	__obf_8dabb278b0366623[11] = Encoding[((__obf_1dfe314328e20778[6]&7)<<2)|((__obf_1dfe314328e20778[7]&192)>>6)]
	__obf_8dabb278b0366623[12] = Encoding[(__obf_1dfe314328e20778[7]&62)>>1]
	__obf_8dabb278b0366623[13] = Encoding[((__obf_1dfe314328e20778[7]&1)<<4)|((__obf_1dfe314328e20778[8]&240)>>4)]
	__obf_8dabb278b0366623[14] = Encoding[((__obf_1dfe314328e20778[8]&15)<<1)|((__obf_1dfe314328e20778[9]&128)>>7)]
	__obf_8dabb278b0366623[15] = Encoding[(__obf_1dfe314328e20778[9]&124)>>2]
	__obf_8dabb278b0366623[16] = Encoding[((__obf_1dfe314328e20778[9]&3)<<3)|((__obf_1dfe314328e20778[10]&224)>>5)]
	__obf_8dabb278b0366623[17] = Encoding[__obf_1dfe314328e20778[10]&31]
	__obf_8dabb278b0366623[18] = Encoding[(__obf_1dfe314328e20778[11]&248)>>3]
	__obf_8dabb278b0366623[19] = Encoding[((__obf_1dfe314328e20778[11]&7)<<2)|((__obf_1dfe314328e20778[12]&192)>>6)]
	__obf_8dabb278b0366623[20] = Encoding[(__obf_1dfe314328e20778[12]&62)>>1]
	__obf_8dabb278b0366623[21] = Encoding[((__obf_1dfe314328e20778[12]&1)<<4)|((__obf_1dfe314328e20778[13]&240)>>4)]
	__obf_8dabb278b0366623[22] = Encoding[((__obf_1dfe314328e20778[13]&15)<<1)|((__obf_1dfe314328e20778[14]&128)>>7)]
	__obf_8dabb278b0366623[23] = Encoding[(__obf_1dfe314328e20778[14]&124)>>2]
	__obf_8dabb278b0366623[24] = Encoding[((__obf_1dfe314328e20778[14]&3)<<3)|((__obf_1dfe314328e20778[15]&224)>>5)]
	__obf_8dabb278b0366623[25] = Encoding[__obf_1dfe314328e20778[15]&31]

	return nil
}

// Byte to index table for O(1) lookups when unmarshaling.
// We use 0xFF as sentinel value for invalid indexes.
var __obf_75ce8c07c15c5d3a = [...]byte{
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
func (__obf_1dfe314328e20778 *ULID) UnmarshalText(__obf_ce304ea72067f76f []byte) error {
	return __obf_46dce11e234e40b8(__obf_ce304ea72067f76f, false, __obf_1dfe314328e20778)
}

// Time returns the Unix time in milliseconds encoded in the ULID.
// Use the top level Time function to convert the returned value to
// a time.Time.
func (__obf_1dfe314328e20778 ULID) Time() uint64 {
	return uint64(__obf_1dfe314328e20778[5]) | uint64(__obf_1dfe314328e20778[4])<<8 |
		uint64(__obf_1dfe314328e20778[3])<<16 | uint64(__obf_1dfe314328e20778[2])<<24 |
		uint64(__obf_1dfe314328e20778[1])<<32 | uint64(__obf_1dfe314328e20778[0])<<40
}

// maxTime is the maximum Unix time in milliseconds that can be
// represented in an ULID.
var __obf_3a31d2b1b03aacda = ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}.Time()

// MaxTime returns the maximum Unix time in milliseconds that
// can be encoded in an ULID.
func MaxTime() uint64 { return __obf_3a31d2b1b03aacda }

// Now is a convenience function that returns the current
// UTC time in Unix milliseconds. Equivalent to:
//
//	Timestamp(time.Now().UTC())
func Now() uint64 { return Timestamp(time.Now().UTC()) }

// Timestamp converts a time.Time to Unix milliseconds.
//
// Because of the way ULID stores time, times from the year
// 10889 produces undefined results.
func Timestamp(__obf_f154b37f4911a720 time.Time) uint64 {
	return uint64(__obf_f154b37f4911a720.Unix())*1000 +
		uint64(__obf_f154b37f4911a720.Nanosecond()/int(time.Millisecond))
}

// Time converts Unix milliseconds in the format
// returned by the Timestamp function to a time.Time.
func Time(__obf_efc47ff3f584b49b uint64) time.Time {
	__obf_8fb7ccdb626087da := int64(__obf_efc47ff3f584b49b / 1e3)
	__obf_23c5a0253f6acbd5 := int64((__obf_efc47ff3f584b49b % 1e3) * 1e6)
	return time.Unix(__obf_8fb7ccdb626087da, __obf_23c5a0253f6acbd5)
}

// SetTime sets the time component of the ULID to the given Unix time
// in milliseconds.
func (__obf_1dfe314328e20778 *ULID) SetTime(__obf_efc47ff3f584b49b uint64) error {
	if __obf_efc47ff3f584b49b > __obf_3a31d2b1b03aacda {
		return ErrBigTime
	}

	(*__obf_1dfe314328e20778)[0] = byte(__obf_efc47ff3f584b49b >> 40)
	(*__obf_1dfe314328e20778)[1] = byte(__obf_efc47ff3f584b49b >> 32)
	(*__obf_1dfe314328e20778)[2] = byte(__obf_efc47ff3f584b49b >> 24)
	(*__obf_1dfe314328e20778)[3] = byte(__obf_efc47ff3f584b49b >> 16)
	(*__obf_1dfe314328e20778)[4] = byte(__obf_efc47ff3f584b49b >> 8)
	(*__obf_1dfe314328e20778)[5] = byte(__obf_efc47ff3f584b49b)

	return nil
}

// Entropy returns the entropy from the ULID.
func (__obf_1dfe314328e20778 ULID) Entropy() []byte {
	__obf_e9c0741dc5032511 := make([]byte, 10)
	copy(__obf_e9c0741dc5032511, __obf_1dfe314328e20778[6:])
	return __obf_e9c0741dc5032511
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (__obf_1dfe314328e20778 *ULID) SetEntropy(__obf_e9c0741dc5032511 []byte) error {
	if len(__obf_e9c0741dc5032511) != 10 {
		return ErrDataSize
	}

	copy((*__obf_1dfe314328e20778)[6:], __obf_e9c0741dc5032511)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (__obf_1dfe314328e20778 ULID) Compare(__obf_9a33b6319cbe29a4 ULID) int {
	return bytes.Compare(__obf_1dfe314328e20778[:], __obf_9a33b6319cbe29a4[:])
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (__obf_1dfe314328e20778 *ULID) Scan(__obf_4187606b16e416d3 any) error {
	switch __obf_ffa353604979097c := __obf_4187606b16e416d3.(type) {
	case nil:
		return nil
	case string:
		return __obf_1dfe314328e20778.UnmarshalText([]byte(__obf_ffa353604979097c))
	case []byte:
		return __obf_1dfe314328e20778.UnmarshalBinary(__obf_ffa353604979097c)
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
func (__obf_1dfe314328e20778 ULID) Value() (driver.Value, error) {
	return __obf_1dfe314328e20778.MarshalBinary()
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
func Monotonic(__obf_67566d403e1ce10f io.Reader, __obf_3edac16313ccf038 uint64) *MonotonicEntropy {
	__obf_262de5093f0a3916 := MonotonicEntropy{
		Reader:                 bufio.NewReader(__obf_67566d403e1ce10f),
		__obf_3edac16313ccf038: __obf_3edac16313ccf038,
	}

	if __obf_262de5093f0a3916.__obf_3edac16313ccf038 == 0 {
		__obf_262de5093f0a3916.__obf_3edac16313ccf038 = math.MaxUint32
	}

	if __obf_d77c8f4b6a147d22, __obf_c8f04275fcb5c686 := __obf_67566d403e1ce10f.(*rand.Rand); __obf_c8f04275fcb5c686 {
		__obf_262de5093f0a3916.__obf_d77c8f4b6a147d22 = __obf_d77c8f4b6a147d22
	}

	return &__obf_262de5093f0a3916
}

// MonotonicEntropy is an opaque type that provides monotonic entropy.
type MonotonicEntropy struct {
	io.Reader
	__obf_efc47ff3f584b49b uint64
	__obf_3edac16313ccf038 uint64
	__obf_67566d403e1ce10f __obf_33b5709b271bfc5c
	rand                   [8]byte
	__obf_d77c8f4b6a147d22 *rand.Rand
}

// MonotonicRead implements the MonotonicReader interface.
func (__obf_262de5093f0a3916 *MonotonicEntropy) MonotonicRead(__obf_efc47ff3f584b49b uint64, __obf_67566d403e1ce10f []byte) (__obf_e62cd9e417a87ee7 error) {
	if !__obf_262de5093f0a3916.__obf_67566d403e1ce10f.IsZero() && __obf_262de5093f0a3916.__obf_efc47ff3f584b49b == __obf_efc47ff3f584b49b {
		__obf_e62cd9e417a87ee7 = __obf_262de5093f0a3916.__obf_19e5fd6e43d0dad9()
		__obf_262de5093f0a3916.__obf_67566d403e1ce10f.AppendTo(__obf_67566d403e1ce10f)
	} else if _, __obf_e62cd9e417a87ee7 = io.ReadFull(__obf_262de5093f0a3916.Reader, __obf_67566d403e1ce10f); __obf_e62cd9e417a87ee7 == nil {
		__obf_262de5093f0a3916.__obf_efc47ff3f584b49b = __obf_efc47ff3f584b49b
		__obf_262de5093f0a3916.__obf_67566d403e1ce10f.SetBytes(__obf_67566d403e1ce10f)
	}
	return __obf_e62cd9e417a87ee7
}

// increment the previous entropy number with a random number
// of up to m.inc (inclusive).
func (__obf_262de5093f0a3916 *MonotonicEntropy) __obf_19e5fd6e43d0dad9() error {
	if __obf_3edac16313ccf038, __obf_e62cd9e417a87ee7 := __obf_262de5093f0a3916.__obf_09dda633b361dd59(); __obf_e62cd9e417a87ee7 != nil {
		return __obf_e62cd9e417a87ee7
	} else if __obf_262de5093f0a3916.__obf_67566d403e1ce10f.Add(__obf_3edac16313ccf038) {
		return ErrMonotonicOverflow
	}
	return nil
}

// random returns a uniform random value in [1, m.inc), reading entropy
// from m.Reader. When m.inc == 0 || m.inc == 1, it returns 1.
// Adapted from: https://golang.org/pkg/crypto/rand/#Int
func (__obf_262de5093f0a3916 *MonotonicEntropy) __obf_09dda633b361dd59() (__obf_3edac16313ccf038 uint64, __obf_e62cd9e417a87ee7 error) {
	if __obf_262de5093f0a3916.__obf_3edac16313ccf038 <= 1 {
		return 1, nil
	}

	// Fast path for using a underlying rand.Rand directly.
	if __obf_262de5093f0a3916.__obf_d77c8f4b6a147d22 != nil {
		// Range: [1, m.inc)
		return 1 + uint64(__obf_262de5093f0a3916.__obf_d77c8f4b6a147d22.Int63n(int64(__obf_262de5093f0a3916.__obf_3edac16313ccf038))), nil
	}

	// bitLen is the maximum bit length needed to encode a value < m.inc.
	__obf_f7ac5c42b11ede92 := bits.Len64(__obf_262de5093f0a3916.__obf_3edac16313ccf038)

	// byteLen is the maximum byte length needed to encode a value < m.inc.
	__obf_e37780da2b35c7c4 := uint(__obf_f7ac5c42b11ede92+7) / 8

	// msbitLen is the number of bits in the most significant byte of m.inc-1.
	__obf_41044651bce391e1 := uint(__obf_f7ac5c42b11ede92 % 8)
	if __obf_41044651bce391e1 == 0 {
		__obf_41044651bce391e1 = 8
	}

	for __obf_3edac16313ccf038 == 0 || __obf_3edac16313ccf038 >= __obf_262de5093f0a3916.__obf_3edac16313ccf038 {
		if _, __obf_e62cd9e417a87ee7 = io.ReadFull(__obf_262de5093f0a3916.Reader, __obf_262de5093f0a3916.rand[:__obf_e37780da2b35c7c4]); __obf_e62cd9e417a87ee7 != nil {
			return 0, __obf_e62cd9e417a87ee7
		}

		// Clear bits in the first byte to increase the probability
		// that the candidate is < m.inc.
		__obf_262de5093f0a3916.rand[0] &= uint8(int(1<<__obf_41044651bce391e1) - 1)

		// Convert the read bytes into an uint64 with byteLen
		// Optimized unrolled loop.
		switch __obf_e37780da2b35c7c4 {
		case 1:
			__obf_3edac16313ccf038 = uint64(__obf_262de5093f0a3916.rand[0])
		case 2:
			__obf_3edac16313ccf038 = uint64(binary.LittleEndian.Uint16(__obf_262de5093f0a3916.rand[:2]))
		case 3, 4:
			__obf_3edac16313ccf038 = uint64(binary.LittleEndian.Uint32(__obf_262de5093f0a3916.rand[:4]))
		case 5, 6, 7, 8:
			__obf_3edac16313ccf038 = uint64(binary.LittleEndian.Uint64(__obf_262de5093f0a3916.rand[:8]))
		}
	}

	// Range: [1, m.inc)
	return 1 + __obf_3edac16313ccf038, nil
}

type __obf_33b5709b271bfc5c struct {
	Hi uint16
	Lo uint64
}

func (__obf_3d63dae1aea4bf71 *__obf_33b5709b271bfc5c) SetBytes(__obf_3ab9bcf798b4a924 []byte) {
	__obf_3d63dae1aea4bf71.Hi = binary.BigEndian.Uint16(__obf_3ab9bcf798b4a924[:2])
	__obf_3d63dae1aea4bf71.Lo = binary.BigEndian.Uint64(__obf_3ab9bcf798b4a924[2:])
}

func (__obf_3d63dae1aea4bf71 *__obf_33b5709b271bfc5c) AppendTo(__obf_3ab9bcf798b4a924 []byte) {
	binary.BigEndian.PutUint16(__obf_3ab9bcf798b4a924[:2], __obf_3d63dae1aea4bf71.Hi)
	binary.BigEndian.PutUint64(__obf_3ab9bcf798b4a924[2:], __obf_3d63dae1aea4bf71.Lo)
}

func (__obf_3d63dae1aea4bf71 *__obf_33b5709b271bfc5c) Add(__obf_61da857b954dcf36 uint64) (__obf_1e68ceeca8b9cc8d bool) {
	__obf_6540fbdb215c3809, __obf_d4260ed0422ffb15 := __obf_3d63dae1aea4bf71.Lo, __obf_3d63dae1aea4bf71.Hi
	if __obf_3d63dae1aea4bf71.Lo += __obf_61da857b954dcf36; __obf_3d63dae1aea4bf71.Lo < __obf_6540fbdb215c3809 {
		__obf_3d63dae1aea4bf71.Hi++
	}
	return __obf_3d63dae1aea4bf71.Hi < __obf_d4260ed0422ffb15
}

func (__obf_3d63dae1aea4bf71 __obf_33b5709b271bfc5c) IsZero() bool {
	return __obf_3d63dae1aea4bf71.Hi == 0 && __obf_3d63dae1aea4bf71.Lo == 0
}
