package __obf_7d8ac56be2e11a40

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
	MonotonicRead(__obf_ac65a7a697c3f660 uint64, __obf_fd897694e595b152 []byte) error
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
func New(__obf_ac65a7a697c3f660 uint64, __obf_5e512fd52092c516 io.Reader) (__obf_6cfbdbb264317300 ULID, __obf_3f3f828d9e46d714 error) {
	if __obf_3f3f828d9e46d714 = __obf_6cfbdbb264317300.SetTime(__obf_ac65a7a697c3f660); __obf_3f3f828d9e46d714 != nil {
		return __obf_6cfbdbb264317300, __obf_3f3f828d9e46d714
	}

	switch __obf_7ccc04e728b458af := __obf_5e512fd52092c516.(type) {
	case nil:
		return __obf_6cfbdbb264317300, __obf_3f3f828d9e46d714
	case MonotonicReader:
		__obf_3f3f828d9e46d714 = __obf_7ccc04e728b458af.MonotonicRead(__obf_ac65a7a697c3f660, __obf_6cfbdbb264317300[6:])
	default:
		_, __obf_3f3f828d9e46d714 = io.ReadFull(__obf_7ccc04e728b458af, __obf_6cfbdbb264317300[6:])
	}

	return __obf_6cfbdbb264317300, __obf_3f3f828d9e46d714
}

// MustNew is a convenience function equivalent to New that panics on failure
// instead of returning an error.
func MustNew(__obf_ac65a7a697c3f660 uint64, __obf_5e512fd52092c516 io.Reader) ULID {
	__obf_6cfbdbb264317300, __obf_3f3f828d9e46d714 := New(__obf_ac65a7a697c3f660, __obf_5e512fd52092c516)
	if __obf_3f3f828d9e46d714 != nil {
		panic(__obf_3f3f828d9e46d714)
	}
	return __obf_6cfbdbb264317300
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(__obf_b821f5606a98dbea string) (__obf_6cfbdbb264317300 ULID, __obf_3f3f828d9e46d714 error) {
	return __obf_6cfbdbb264317300, __obf_8fd26739f58667a6([]byte(__obf_b821f5606a98dbea), false, &__obf_6cfbdbb264317300)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(__obf_b821f5606a98dbea string) (__obf_6cfbdbb264317300 ULID, __obf_3f3f828d9e46d714 error) {
	return __obf_6cfbdbb264317300, __obf_8fd26739f58667a6([]byte(__obf_b821f5606a98dbea), true, &__obf_6cfbdbb264317300)
}

func __obf_8fd26739f58667a6(__obf_88dfec73463ba77a []byte, __obf_97ac6d1cf225a2e0 bool, __obf_6cfbdbb264317300 *ULID) error {
	// Check if a base32 encoded ULID is the right length.
	if len(__obf_88dfec73463ba77a) != EncodedSize {
		return ErrDataSize
	}

	// Check if all the characters in a base32 encoded ULID are part of the
	// expected base32 character set.
	if __obf_97ac6d1cf225a2e0 &&
		(__obf_235b330c52312d16[__obf_88dfec73463ba77a[0]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[1]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[2]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[3]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[4]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[5]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[6]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[7]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[8]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[9]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[10]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[11]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[12]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[13]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[14]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[15]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[16]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[17]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[18]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[19]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[20]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[21]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[22]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[23]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[24]] == 0xFF ||
			__obf_235b330c52312d16[__obf_88dfec73463ba77a[25]] == 0xFF) {
		return ErrInvalidCharacters
	}

	// Check if the first character in a base32 encoded ULID will overflow. This
	// happens because the base32 representation encodes 130 bits, while the
	// ULID is only 128 bits.
	//
	// See https://github.com/oklog/ulid/issues/9 for details.
	if __obf_88dfec73463ba77a[0] > '7' {
		return ErrOverflow
	}

	// Use an optimized unrolled loop (from https://github.com/RobThree/NUlid)
	// to decode a base32 ULID.

	// 6 bytes timestamp (48 bits)
	(*__obf_6cfbdbb264317300)[0] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[0]] << 5) | __obf_235b330c52312d16[__obf_88dfec73463ba77a[1]]
	(*__obf_6cfbdbb264317300)[1] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[2]] << 3) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[3]] >> 2)
	(*__obf_6cfbdbb264317300)[2] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[3]] << 6) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[4]] << 1) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[5]] >> 4)
	(*__obf_6cfbdbb264317300)[3] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[5]] << 4) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[6]] >> 1)
	(*__obf_6cfbdbb264317300)[4] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[6]] << 7) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[7]] << 2) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[8]] >> 3)
	(*__obf_6cfbdbb264317300)[5] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[8]] << 5) | __obf_235b330c52312d16[__obf_88dfec73463ba77a[9]]

	// 10 bytes of entropy (80 bits)
	(*__obf_6cfbdbb264317300)[6] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[10]] << 3) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[11]] >> 2)
	(*__obf_6cfbdbb264317300)[7] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[11]] << 6) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[12]] << 1) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[13]] >> 4)
	(*__obf_6cfbdbb264317300)[8] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[13]] << 4) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[14]] >> 1)
	(*__obf_6cfbdbb264317300)[9] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[14]] << 7) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[15]] << 2) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[16]] >> 3)
	(*__obf_6cfbdbb264317300)[10] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[16]] << 5) | __obf_235b330c52312d16[__obf_88dfec73463ba77a[17]]
	(*__obf_6cfbdbb264317300)[11] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[18]] << 3) | __obf_235b330c52312d16[__obf_88dfec73463ba77a[19]]>>2
	(*__obf_6cfbdbb264317300)[12] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[19]] << 6) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[20]] << 1) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[21]] >> 4)
	(*__obf_6cfbdbb264317300)[13] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[21]] << 4) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[22]] >> 1)
	(*__obf_6cfbdbb264317300)[14] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[22]] << 7) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[23]] << 2) | (__obf_235b330c52312d16[__obf_88dfec73463ba77a[24]] >> 3)
	(*__obf_6cfbdbb264317300)[15] = (__obf_235b330c52312d16[__obf_88dfec73463ba77a[24]] << 5) | __obf_235b330c52312d16[__obf_88dfec73463ba77a[25]]

	return nil
}

// MustParse is a convenience function equivalent to Parse that panics on failure
// instead of returning an error.
func MustParse(__obf_b821f5606a98dbea string) ULID {
	__obf_6cfbdbb264317300, __obf_3f3f828d9e46d714 := Parse(__obf_b821f5606a98dbea)
	if __obf_3f3f828d9e46d714 != nil {
		panic(__obf_3f3f828d9e46d714)
	}
	return __obf_6cfbdbb264317300
}

// MustParseStrict is a convenience function equivalent to ParseStrict that
// panics on failure instead of returning an error.
func MustParseStrict(__obf_b821f5606a98dbea string) ULID {
	__obf_6cfbdbb264317300, __obf_3f3f828d9e46d714 := ParseStrict(__obf_b821f5606a98dbea)
	if __obf_3f3f828d9e46d714 != nil {
		panic(__obf_3f3f828d9e46d714)
	}
	return __obf_6cfbdbb264317300
}

// Bytes returns bytes slice representation of ULID.
func (__obf_841c11dc9eb6471a ULID) Bytes() []byte {
	return __obf_841c11dc9eb6471a[:]
}

// String returns a lexicographically sortable string encoded ULID
// (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy
func (__obf_6cfbdbb264317300 ULID) String() string {
	__obf_b821f5606a98dbea := make([]byte, EncodedSize)
	_ = __obf_6cfbdbb264317300.MarshalTextTo(__obf_b821f5606a98dbea)
	return string(__obf_b821f5606a98dbea)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (__obf_6cfbdbb264317300 ULID) MarshalBinary() ([]byte, error) {
	__obf_b821f5606a98dbea := make([]byte, len(__obf_6cfbdbb264317300))
	return __obf_b821f5606a98dbea, __obf_6cfbdbb264317300.MarshalBinaryTo(__obf_b821f5606a98dbea)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (__obf_6cfbdbb264317300 ULID) MarshalBinaryTo(__obf_6014b5249e82ec1f []byte) error {
	if len(__obf_6014b5249e82ec1f) != len(__obf_6cfbdbb264317300) {
		return ErrBufferSize
	}

	copy(__obf_6014b5249e82ec1f, __obf_6cfbdbb264317300[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (__obf_6cfbdbb264317300 *ULID) UnmarshalBinary(__obf_06461dd3b32c24bd []byte) error {
	if len(__obf_06461dd3b32c24bd) != len(*__obf_6cfbdbb264317300) {
		return ErrDataSize
	}

	copy((*__obf_6cfbdbb264317300)[:], __obf_06461dd3b32c24bd)
	return nil
}

// Encoding is the base 32 encoding alphabet used in ULID strings.
const Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (__obf_6cfbdbb264317300 ULID) MarshalText() ([]byte, error) {
	__obf_b821f5606a98dbea := make([]byte, EncodedSize)
	return __obf_b821f5606a98dbea, __obf_6cfbdbb264317300.MarshalTextTo(__obf_b821f5606a98dbea)
}

// MarshalTextTo writes the ULID as a string to the given buffer.
// ErrBufferSize is returned when the len(dst) != 26.
func (__obf_6cfbdbb264317300 ULID) MarshalTextTo(__obf_6014b5249e82ec1f []byte) error {
	// Optimized unrolled loop ahead.
	// From https://github.com/RobThree/NUlid

	if len(__obf_6014b5249e82ec1f) != EncodedSize {
		return ErrBufferSize
	}

	// 10 byte timestamp
	__obf_6014b5249e82ec1f[0] = Encoding[(__obf_6cfbdbb264317300[0]&224)>>5]
	__obf_6014b5249e82ec1f[1] = Encoding[__obf_6cfbdbb264317300[0]&31]
	__obf_6014b5249e82ec1f[2] = Encoding[(__obf_6cfbdbb264317300[1]&248)>>3]
	__obf_6014b5249e82ec1f[3] = Encoding[((__obf_6cfbdbb264317300[1]&7)<<2)|((__obf_6cfbdbb264317300[2]&192)>>6)]
	__obf_6014b5249e82ec1f[4] = Encoding[(__obf_6cfbdbb264317300[2]&62)>>1]
	__obf_6014b5249e82ec1f[5] = Encoding[((__obf_6cfbdbb264317300[2]&1)<<4)|((__obf_6cfbdbb264317300[3]&240)>>4)]
	__obf_6014b5249e82ec1f[6] = Encoding[((__obf_6cfbdbb264317300[3]&15)<<1)|((__obf_6cfbdbb264317300[4]&128)>>7)]
	__obf_6014b5249e82ec1f[7] = Encoding[(__obf_6cfbdbb264317300[4]&124)>>2]
	__obf_6014b5249e82ec1f[8] = Encoding[((__obf_6cfbdbb264317300[4]&3)<<3)|((__obf_6cfbdbb264317300[5]&224)>>5)]
	__obf_6014b5249e82ec1f[9] = Encoding[__obf_6cfbdbb264317300[5]&31]

	// 16 bytes of entropy
	__obf_6014b5249e82ec1f[10] = Encoding[(__obf_6cfbdbb264317300[6]&248)>>3]
	__obf_6014b5249e82ec1f[11] = Encoding[((__obf_6cfbdbb264317300[6]&7)<<2)|((__obf_6cfbdbb264317300[7]&192)>>6)]
	__obf_6014b5249e82ec1f[12] = Encoding[(__obf_6cfbdbb264317300[7]&62)>>1]
	__obf_6014b5249e82ec1f[13] = Encoding[((__obf_6cfbdbb264317300[7]&1)<<4)|((__obf_6cfbdbb264317300[8]&240)>>4)]
	__obf_6014b5249e82ec1f[14] = Encoding[((__obf_6cfbdbb264317300[8]&15)<<1)|((__obf_6cfbdbb264317300[9]&128)>>7)]
	__obf_6014b5249e82ec1f[15] = Encoding[(__obf_6cfbdbb264317300[9]&124)>>2]
	__obf_6014b5249e82ec1f[16] = Encoding[((__obf_6cfbdbb264317300[9]&3)<<3)|((__obf_6cfbdbb264317300[10]&224)>>5)]
	__obf_6014b5249e82ec1f[17] = Encoding[__obf_6cfbdbb264317300[10]&31]
	__obf_6014b5249e82ec1f[18] = Encoding[(__obf_6cfbdbb264317300[11]&248)>>3]
	__obf_6014b5249e82ec1f[19] = Encoding[((__obf_6cfbdbb264317300[11]&7)<<2)|((__obf_6cfbdbb264317300[12]&192)>>6)]
	__obf_6014b5249e82ec1f[20] = Encoding[(__obf_6cfbdbb264317300[12]&62)>>1]
	__obf_6014b5249e82ec1f[21] = Encoding[((__obf_6cfbdbb264317300[12]&1)<<4)|((__obf_6cfbdbb264317300[13]&240)>>4)]
	__obf_6014b5249e82ec1f[22] = Encoding[((__obf_6cfbdbb264317300[13]&15)<<1)|((__obf_6cfbdbb264317300[14]&128)>>7)]
	__obf_6014b5249e82ec1f[23] = Encoding[(__obf_6cfbdbb264317300[14]&124)>>2]
	__obf_6014b5249e82ec1f[24] = Encoding[((__obf_6cfbdbb264317300[14]&3)<<3)|((__obf_6cfbdbb264317300[15]&224)>>5)]
	__obf_6014b5249e82ec1f[25] = Encoding[__obf_6cfbdbb264317300[15]&31]

	return nil
}

// Byte to index table for O(1) lookups when unmarshaling.
// We use 0xFF as sentinel value for invalid indexes.
var __obf_235b330c52312d16 = [...]byte{
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
func (__obf_6cfbdbb264317300 *ULID) UnmarshalText(__obf_88dfec73463ba77a []byte) error {
	return __obf_8fd26739f58667a6(__obf_88dfec73463ba77a, false, __obf_6cfbdbb264317300)
}

// Time returns the Unix time in milliseconds encoded in the ULID.
// Use the top level Time function to convert the returned value to
// a time.Time.
func (__obf_6cfbdbb264317300 ULID) Time() uint64 {
	return uint64(__obf_6cfbdbb264317300[5]) | uint64(__obf_6cfbdbb264317300[4])<<8 |
		uint64(__obf_6cfbdbb264317300[3])<<16 | uint64(__obf_6cfbdbb264317300[2])<<24 |
		uint64(__obf_6cfbdbb264317300[1])<<32 | uint64(__obf_6cfbdbb264317300[0])<<40
}

// maxTime is the maximum Unix time in milliseconds that can be
// represented in an ULID.
var __obf_0f001d53a3ccbcd2 = ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}.Time()

// MaxTime returns the maximum Unix time in milliseconds that
// can be encoded in an ULID.
func MaxTime() uint64 { return __obf_0f001d53a3ccbcd2 }

// Now is a convenience function that returns the current
// UTC time in Unix milliseconds. Equivalent to:
//
//	Timestamp(time.Now().UTC())
func Now() uint64 { return Timestamp(time.Now().UTC()) }

// Timestamp converts a time.Time to Unix milliseconds.
//
// Because of the way ULID stores time, times from the year
// 10889 produces undefined results.
func Timestamp(__obf_680f17a8e930eb64 time.Time) uint64 {
	return uint64(__obf_680f17a8e930eb64.Unix())*1000 +
		uint64(__obf_680f17a8e930eb64.Nanosecond()/int(time.Millisecond))
}

// Time converts Unix milliseconds in the format
// returned by the Timestamp function to a time.Time.
func Time(__obf_ac65a7a697c3f660 uint64) time.Time {
	__obf_72b3d0170844a599 := int64(__obf_ac65a7a697c3f660 / 1e3)
	__obf_146f175924d13464 := int64((__obf_ac65a7a697c3f660 % 1e3) * 1e6)
	return time.Unix(__obf_72b3d0170844a599, __obf_146f175924d13464)
}

// SetTime sets the time component of the ULID to the given Unix time
// in milliseconds.
func (__obf_6cfbdbb264317300 *ULID) SetTime(__obf_ac65a7a697c3f660 uint64) error {
	if __obf_ac65a7a697c3f660 > __obf_0f001d53a3ccbcd2 {
		return ErrBigTime
	}

	(*__obf_6cfbdbb264317300)[0] = byte(__obf_ac65a7a697c3f660 >> 40)
	(*__obf_6cfbdbb264317300)[1] = byte(__obf_ac65a7a697c3f660 >> 32)
	(*__obf_6cfbdbb264317300)[2] = byte(__obf_ac65a7a697c3f660 >> 24)
	(*__obf_6cfbdbb264317300)[3] = byte(__obf_ac65a7a697c3f660 >> 16)
	(*__obf_6cfbdbb264317300)[4] = byte(__obf_ac65a7a697c3f660 >> 8)
	(*__obf_6cfbdbb264317300)[5] = byte(__obf_ac65a7a697c3f660)

	return nil
}

// Entropy returns the entropy from the ULID.
func (__obf_6cfbdbb264317300 ULID) Entropy() []byte {
	__obf_7ccc04e728b458af := make([]byte, 10)
	copy(__obf_7ccc04e728b458af, __obf_6cfbdbb264317300[6:])
	return __obf_7ccc04e728b458af
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (__obf_6cfbdbb264317300 *ULID) SetEntropy(__obf_7ccc04e728b458af []byte) error {
	if len(__obf_7ccc04e728b458af) != 10 {
		return ErrDataSize
	}

	copy((*__obf_6cfbdbb264317300)[6:], __obf_7ccc04e728b458af)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (__obf_6cfbdbb264317300 ULID) Compare(__obf_d0b85127c7257422 ULID) int {
	return bytes.Compare(__obf_6cfbdbb264317300[:], __obf_d0b85127c7257422[:])
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (__obf_6cfbdbb264317300 *ULID) Scan(__obf_b2e4d1651f2fb10e any) error {
	switch __obf_e8e0549a04264062 := __obf_b2e4d1651f2fb10e.(type) {
	case nil:
		return nil
	case string:
		return __obf_6cfbdbb264317300.UnmarshalText([]byte(__obf_e8e0549a04264062))
	case []byte:
		return __obf_6cfbdbb264317300.UnmarshalBinary(__obf_e8e0549a04264062)
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
func (__obf_6cfbdbb264317300 ULID) Value() (driver.Value, error) {
	return __obf_6cfbdbb264317300.MarshalBinary()
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
func Monotonic(__obf_5e512fd52092c516 io.Reader, __obf_625fbdfd5c25e64f uint64) *MonotonicEntropy {
	__obf_76024f9c3840a3ba := MonotonicEntropy{
		Reader:                 bufio.NewReader(__obf_5e512fd52092c516),
		__obf_625fbdfd5c25e64f: __obf_625fbdfd5c25e64f,
	}

	if __obf_76024f9c3840a3ba.__obf_625fbdfd5c25e64f == 0 {
		__obf_76024f9c3840a3ba.__obf_625fbdfd5c25e64f = math.MaxUint32
	}

	if __obf_03f20a63f1ecffc1, __obf_4f3a1b70936ef093 := __obf_5e512fd52092c516.(*rand.Rand); __obf_4f3a1b70936ef093 {
		__obf_76024f9c3840a3ba.__obf_03f20a63f1ecffc1 = __obf_03f20a63f1ecffc1
	}

	return &__obf_76024f9c3840a3ba
}

// MonotonicEntropy is an opaque type that provides monotonic entropy.
type MonotonicEntropy struct {
	io.Reader
	__obf_ac65a7a697c3f660 uint64
	__obf_625fbdfd5c25e64f uint64
	__obf_5e512fd52092c516 __obf_44d191ad4e3fa0a3
	rand                   [8]byte
	__obf_03f20a63f1ecffc1 *rand.Rand
}

// MonotonicRead implements the MonotonicReader interface.
func (__obf_76024f9c3840a3ba *MonotonicEntropy) MonotonicRead(__obf_ac65a7a697c3f660 uint64, __obf_5e512fd52092c516 []byte) (__obf_3f3f828d9e46d714 error) {
	if !__obf_76024f9c3840a3ba.__obf_5e512fd52092c516.IsZero() && __obf_76024f9c3840a3ba.__obf_ac65a7a697c3f660 == __obf_ac65a7a697c3f660 {
		__obf_3f3f828d9e46d714 = __obf_76024f9c3840a3ba.__obf_1e7b87d851e13054()
		__obf_76024f9c3840a3ba.__obf_5e512fd52092c516.AppendTo(__obf_5e512fd52092c516)
	} else if _, __obf_3f3f828d9e46d714 = io.ReadFull(__obf_76024f9c3840a3ba.Reader, __obf_5e512fd52092c516); __obf_3f3f828d9e46d714 == nil {
		__obf_76024f9c3840a3ba.__obf_ac65a7a697c3f660 = __obf_ac65a7a697c3f660
		__obf_76024f9c3840a3ba.__obf_5e512fd52092c516.SetBytes(__obf_5e512fd52092c516)
	}
	return __obf_3f3f828d9e46d714
}

// increment the previous entropy number with a random number
// of up to m.inc (inclusive).
func (__obf_76024f9c3840a3ba *MonotonicEntropy) __obf_1e7b87d851e13054() error {
	if __obf_625fbdfd5c25e64f, __obf_3f3f828d9e46d714 := __obf_76024f9c3840a3ba.__obf_8c0bf1b2f07165dc(); __obf_3f3f828d9e46d714 != nil {
		return __obf_3f3f828d9e46d714
	} else if __obf_76024f9c3840a3ba.__obf_5e512fd52092c516.Add(__obf_625fbdfd5c25e64f) {
		return ErrMonotonicOverflow
	}
	return nil
}

// random returns a uniform random value in [1, m.inc), reading entropy
// from m.Reader. When m.inc == 0 || m.inc == 1, it returns 1.
// Adapted from: https://golang.org/pkg/crypto/rand/#Int
func (__obf_76024f9c3840a3ba *MonotonicEntropy) __obf_8c0bf1b2f07165dc() (__obf_625fbdfd5c25e64f uint64, __obf_3f3f828d9e46d714 error) {
	if __obf_76024f9c3840a3ba.__obf_625fbdfd5c25e64f <= 1 {
		return 1, nil
	}

	// Fast path for using a underlying rand.Rand directly.
	if __obf_76024f9c3840a3ba.__obf_03f20a63f1ecffc1 != nil {
		// Range: [1, m.inc)
		return 1 + uint64(__obf_76024f9c3840a3ba.__obf_03f20a63f1ecffc1.Int63n(int64(__obf_76024f9c3840a3ba.__obf_625fbdfd5c25e64f))), nil
	}

	// bitLen is the maximum bit length needed to encode a value < m.inc.
	__obf_c86bac20c2c288fd := bits.Len64(__obf_76024f9c3840a3ba.__obf_625fbdfd5c25e64f)

	// byteLen is the maximum byte length needed to encode a value < m.inc.
	__obf_4253a37bd6df4e24 := uint(__obf_c86bac20c2c288fd+7) / 8

	// msbitLen is the number of bits in the most significant byte of m.inc-1.
	__obf_0b7210bf6c2ff45a := uint(__obf_c86bac20c2c288fd % 8)
	if __obf_0b7210bf6c2ff45a == 0 {
		__obf_0b7210bf6c2ff45a = 8
	}

	for __obf_625fbdfd5c25e64f == 0 || __obf_625fbdfd5c25e64f >= __obf_76024f9c3840a3ba.__obf_625fbdfd5c25e64f {
		if _, __obf_3f3f828d9e46d714 = io.ReadFull(__obf_76024f9c3840a3ba.Reader, __obf_76024f9c3840a3ba.rand[:__obf_4253a37bd6df4e24]); __obf_3f3f828d9e46d714 != nil {
			return 0, __obf_3f3f828d9e46d714
		}

		// Clear bits in the first byte to increase the probability
		// that the candidate is < m.inc.
		__obf_76024f9c3840a3ba.rand[0] &= uint8(int(1<<__obf_0b7210bf6c2ff45a) - 1)

		// Convert the read bytes into an uint64 with byteLen
		// Optimized unrolled loop.
		switch __obf_4253a37bd6df4e24 {
		case 1:
			__obf_625fbdfd5c25e64f = uint64(__obf_76024f9c3840a3ba.rand[0])
		case 2:
			__obf_625fbdfd5c25e64f = uint64(binary.LittleEndian.Uint16(__obf_76024f9c3840a3ba.rand[:2]))
		case 3, 4:
			__obf_625fbdfd5c25e64f = uint64(binary.LittleEndian.Uint32(__obf_76024f9c3840a3ba.rand[:4]))
		case 5, 6, 7, 8:
			__obf_625fbdfd5c25e64f = uint64(binary.LittleEndian.Uint64(__obf_76024f9c3840a3ba.rand[:8]))
		}
	}

	// Range: [1, m.inc)
	return 1 + __obf_625fbdfd5c25e64f, nil
}

type __obf_44d191ad4e3fa0a3 struct {
	Hi uint16
	Lo uint64
}

func (__obf_841c11dc9eb6471a *__obf_44d191ad4e3fa0a3) SetBytes(__obf_095a7f48d7bcc318 []byte) {
	__obf_841c11dc9eb6471a.Hi = binary.BigEndian.Uint16(__obf_095a7f48d7bcc318[:2])
	__obf_841c11dc9eb6471a.Lo = binary.BigEndian.Uint64(__obf_095a7f48d7bcc318[2:])
}

func (__obf_841c11dc9eb6471a *__obf_44d191ad4e3fa0a3) AppendTo(__obf_095a7f48d7bcc318 []byte) {
	binary.BigEndian.PutUint16(__obf_095a7f48d7bcc318[:2], __obf_841c11dc9eb6471a.Hi)
	binary.BigEndian.PutUint64(__obf_095a7f48d7bcc318[2:], __obf_841c11dc9eb6471a.Lo)
}

func (__obf_841c11dc9eb6471a *__obf_44d191ad4e3fa0a3) Add(__obf_eacf881057f264e4 uint64) (__obf_93f8f159779c465a bool) {
	__obf_c165a24267382889, __obf_9e3e190a11a6782d := __obf_841c11dc9eb6471a.Lo, __obf_841c11dc9eb6471a.Hi
	if __obf_841c11dc9eb6471a.Lo += __obf_eacf881057f264e4; __obf_841c11dc9eb6471a.Lo < __obf_c165a24267382889 {
		__obf_841c11dc9eb6471a.Hi++
	}
	return __obf_841c11dc9eb6471a.Hi < __obf_9e3e190a11a6782d
}

func (__obf_841c11dc9eb6471a __obf_44d191ad4e3fa0a3) IsZero() bool {
	return __obf_841c11dc9eb6471a.Hi == 0 && __obf_841c11dc9eb6471a.Lo == 0
}
