package __obf_2f51f7d26a2bcdf8

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
	MonotonicRead(__obf_4450a2d1c70a6bed uint64, __obf_3cd6b42eee80d806 []byte) error
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
func New(__obf_4450a2d1c70a6bed uint64, __obf_64686fca0f92edfe io.Reader) (__obf_8b246d7b3659af37 ULID, __obf_37e8e388755917c0 error) {
	if __obf_37e8e388755917c0 = __obf_8b246d7b3659af37.SetTime(__obf_4450a2d1c70a6bed); __obf_37e8e388755917c0 != nil {
		return __obf_8b246d7b3659af37, __obf_37e8e388755917c0
	}

	switch __obf_8b3554b161bd95b9 := __obf_64686fca0f92edfe.(type) {
	case nil:
		return __obf_8b246d7b3659af37, __obf_37e8e388755917c0
	case MonotonicReader:
		__obf_37e8e388755917c0 = __obf_8b3554b161bd95b9.MonotonicRead(__obf_4450a2d1c70a6bed, __obf_8b246d7b3659af37[6:])
	default:
		_, __obf_37e8e388755917c0 = io.ReadFull(__obf_8b3554b161bd95b9, __obf_8b246d7b3659af37[6:])
	}

	return __obf_8b246d7b3659af37,

		// MustNew is a convenience function equivalent to New that panics on failure
		// instead of returning an error.
		__obf_37e8e388755917c0
}

func MustNew(__obf_4450a2d1c70a6bed uint64, __obf_64686fca0f92edfe io.Reader) ULID {
	__obf_8b246d7b3659af37, __obf_37e8e388755917c0 := New(__obf_4450a2d1c70a6bed, __obf_64686fca0f92edfe)
	if __obf_37e8e388755917c0 != nil {
		panic(__obf_37e8e388755917c0)
	}
	return __obf_8b246d7b3659af37
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(__obf_ac3eac88e099349d string) (__obf_8b246d7b3659af37 ULID, __obf_37e8e388755917c0 error) {
	return __obf_8b246d7b3659af37, __obf_c28a75a15814fc86([]byte(__obf_ac3eac88e099349d), false, &__obf_8b246d7b3659af37)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(__obf_ac3eac88e099349d string) (__obf_8b246d7b3659af37 ULID, __obf_37e8e388755917c0 error) {
	return __obf_8b246d7b3659af37, __obf_c28a75a15814fc86([]byte(__obf_ac3eac88e099349d), true, &__obf_8b246d7b3659af37)
}

func __obf_c28a75a15814fc86(__obf_1caf1c7883b88b75 []byte, __obf_6eda45070955415e bool, __obf_8b246d7b3659af37 *ULID) error {
	// Check if a base32 encoded ULID is the right length.
	if len(__obf_1caf1c7883b88b75) != EncodedSize {
		return ErrDataSize
	}

	// Check if all the characters in a base32 encoded ULID are part of the
	// expected base32 character set.
	if __obf_6eda45070955415e &&
		(__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[0]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[1]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[2]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[3]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[4]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[5]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[6]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[7]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[8]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[9]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[10]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[11]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[12]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[13]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[14]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[15]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[16]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[17]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[18]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[19]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[20]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[21]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[22]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[23]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[24]] == 0xFF || __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[25]] == 0xFF) {
		return ErrInvalidCharacters
	}

	// Check if the first character in a base32 encoded ULID will overflow. This
	// happens because the base32 representation encodes 130 bits, while the
	// ULID is only 128 bits.
	//
	// See https://github.com/oklog/ulid/issues/9 for details.
	if __obf_1caf1c7883b88b75[0] > '7' {
		return ErrOverflow
	}

	// Use an optimized unrolled loop (from https://github.com/RobThree/NUlid)
	// to decode a base32 ULID.

	// 6 bytes timestamp (48 bits)
	(*__obf_8b246d7b3659af37)[0] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[0]] << 5) | __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[1]]
	(*__obf_8b246d7b3659af37)[1] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[2]] << 3) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[3]] >> 2)
	(*__obf_8b246d7b3659af37)[2] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[3]] << 6) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[4]] << 1) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[5]] >> 4)
	(*__obf_8b246d7b3659af37)[3] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[5]] << 4) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[6]] >> 1)
	(*__obf_8b246d7b3659af37)[4] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[6]] << 7) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[7]] << 2) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[8]] >> 3)
	(*__obf_8b246d7b3659af37)[5] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[8]] << 5) | __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[9]]

	// 10 bytes of entropy (80 bits)
	(*__obf_8b246d7b3659af37)[6] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[10]] << 3) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[11]] >> 2)
	(*__obf_8b246d7b3659af37)[7] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[11]] << 6) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[12]] << 1) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[13]] >> 4)
	(*__obf_8b246d7b3659af37)[8] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[13]] << 4) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[14]] >> 1)
	(*__obf_8b246d7b3659af37)[9] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[14]] << 7) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[15]] << 2) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[16]] >> 3)
	(*__obf_8b246d7b3659af37)[10] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[16]] << 5) | __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[17]]
	(*__obf_8b246d7b3659af37)[11] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[18]] << 3) | __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[19]]>>2
	(*__obf_8b246d7b3659af37)[12] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[19]] << 6) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[20]] << 1) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[21]] >> 4)
	(*__obf_8b246d7b3659af37)[13] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[21]] << 4) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[22]] >> 1)
	(*__obf_8b246d7b3659af37)[14] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[22]] << 7) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[23]] << 2) | (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[24]] >> 3)
	(*__obf_8b246d7b3659af37)[15] = (__obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[24]] << 5) | __obf_c91e2f0f66e52747[__obf_1caf1c7883b88b75[25]]

	return nil
}

// MustParse is a convenience function equivalent to Parse that panics on failure
// instead of returning an error.
func MustParse(__obf_ac3eac88e099349d string) ULID {
	__obf_8b246d7b3659af37, __obf_37e8e388755917c0 := Parse(__obf_ac3eac88e099349d)
	if __obf_37e8e388755917c0 != nil {
		panic(__obf_37e8e388755917c0)
	}
	return __obf_8b246d7b3659af37
}

// MustParseStrict is a convenience function equivalent to ParseStrict that
// panics on failure instead of returning an error.
func MustParseStrict(__obf_ac3eac88e099349d string) ULID {
	__obf_8b246d7b3659af37, __obf_37e8e388755917c0 := ParseStrict(__obf_ac3eac88e099349d)
	if __obf_37e8e388755917c0 != nil {
		panic(__obf_37e8e388755917c0)
	}
	return __obf_8b246d7b3659af37
}

// Bytes returns bytes slice representation of ULID.
func (__obf_feb0d3963f34862a ULID) Bytes() []byte {
	return __obf_feb0d3963f34862a[:]
}

// String returns a lexicographically sortable string encoded ULID
// (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy
func (__obf_8b246d7b3659af37 ULID) String() string {
	__obf_ac3eac88e099349d := make([]byte, EncodedSize)
	_ = __obf_8b246d7b3659af37.MarshalTextTo(__obf_ac3eac88e099349d)
	return string(__obf_ac3eac88e099349d)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (__obf_8b246d7b3659af37 ULID) MarshalBinary() ([]byte, error) {
	__obf_ac3eac88e099349d := make([]byte, len(__obf_8b246d7b3659af37))
	return __obf_ac3eac88e099349d, __obf_8b246d7b3659af37.MarshalBinaryTo(__obf_ac3eac88e099349d)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (__obf_8b246d7b3659af37 ULID) MarshalBinaryTo(__obf_db477556a08e916c []byte) error {
	if len(__obf_db477556a08e916c) != len(__obf_8b246d7b3659af37) {
		return ErrBufferSize
	}

	copy(__obf_db477556a08e916c, __obf_8b246d7b3659af37[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (__obf_8b246d7b3659af37 *ULID) UnmarshalBinary(__obf_820f902ebc42d671 []byte) error {
	if len(__obf_820f902ebc42d671) != len(*__obf_8b246d7b3659af37) {
		return ErrDataSize
	}

	copy((*__obf_8b246d7b3659af37)[:], __obf_820f902ebc42d671)
	return nil
}

// Encoding is the base 32 encoding alphabet used in ULID strings.
const Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (__obf_8b246d7b3659af37 ULID) MarshalText() ([]byte, error) {
	__obf_ac3eac88e099349d := make([]byte, EncodedSize)
	return __obf_ac3eac88e099349d, __obf_8b246d7b3659af37.MarshalTextTo(__obf_ac3eac88e099349d)
}

// MarshalTextTo writes the ULID as a string to the given buffer.
// ErrBufferSize is returned when the len(dst) != 26.
func (__obf_8b246d7b3659af37 ULID) MarshalTextTo(__obf_db477556a08e916c []byte) error {
	// Optimized unrolled loop ahead.
	// From https://github.com/RobThree/NUlid

	if len(__obf_db477556a08e916c) != EncodedSize {
		return ErrBufferSize
	}
	__obf_db477556a08e916c[ // 10 byte timestamp
	0] = Encoding[(__obf_8b246d7b3659af37[0]&224)>>5]
	__obf_db477556a08e916c[1] = Encoding[__obf_8b246d7b3659af37[0]&31]
	__obf_db477556a08e916c[2] = Encoding[(__obf_8b246d7b3659af37[1]&248)>>3]
	__obf_db477556a08e916c[3] = Encoding[((__obf_8b246d7b3659af37[1]&7)<<2)|((__obf_8b246d7b3659af37[2]&192)>>6)]
	__obf_db477556a08e916c[4] = Encoding[(__obf_8b246d7b3659af37[2]&62)>>1]
	__obf_db477556a08e916c[5] = Encoding[((__obf_8b246d7b3659af37[2]&1)<<4)|((__obf_8b246d7b3659af37[3]&240)>>4)]
	__obf_db477556a08e916c[6] = Encoding[((__obf_8b246d7b3659af37[3]&15)<<1)|((__obf_8b246d7b3659af37[4]&128)>>7)]
	__obf_db477556a08e916c[7] = Encoding[(__obf_8b246d7b3659af37[4]&124)>>2]
	__obf_db477556a08e916c[8] = Encoding[((__obf_8b246d7b3659af37[4]&3)<<3)|((__obf_8b246d7b3659af37[5]&224)>>5)]
	__obf_db477556a08e916c[9] = Encoding[__obf_8b246d7b3659af37[5]&31]
	__obf_db477556a08e916c[ // 16 bytes of entropy
	10] = Encoding[(__obf_8b246d7b3659af37[6]&248)>>3]
	__obf_db477556a08e916c[11] = Encoding[((__obf_8b246d7b3659af37[6]&7)<<2)|((__obf_8b246d7b3659af37[7]&192)>>6)]
	__obf_db477556a08e916c[12] = Encoding[(__obf_8b246d7b3659af37[7]&62)>>1]
	__obf_db477556a08e916c[13] = Encoding[((__obf_8b246d7b3659af37[7]&1)<<4)|((__obf_8b246d7b3659af37[8]&240)>>4)]
	__obf_db477556a08e916c[14] = Encoding[((__obf_8b246d7b3659af37[8]&15)<<1)|((__obf_8b246d7b3659af37[9]&128)>>7)]
	__obf_db477556a08e916c[15] = Encoding[(__obf_8b246d7b3659af37[9]&124)>>2]
	__obf_db477556a08e916c[16] = Encoding[((__obf_8b246d7b3659af37[9]&3)<<3)|((__obf_8b246d7b3659af37[10]&224)>>5)]
	__obf_db477556a08e916c[17] = Encoding[__obf_8b246d7b3659af37[10]&31]
	__obf_db477556a08e916c[18] = Encoding[(__obf_8b246d7b3659af37[11]&248)>>3]
	__obf_db477556a08e916c[19] = Encoding[((__obf_8b246d7b3659af37[11]&7)<<2)|((__obf_8b246d7b3659af37[12]&192)>>6)]
	__obf_db477556a08e916c[20] = Encoding[(__obf_8b246d7b3659af37[12]&62)>>1]
	__obf_db477556a08e916c[21] = Encoding[((__obf_8b246d7b3659af37[12]&1)<<4)|((__obf_8b246d7b3659af37[13]&240)>>4)]
	__obf_db477556a08e916c[22] = Encoding[((__obf_8b246d7b3659af37[13]&15)<<1)|((__obf_8b246d7b3659af37[14]&128)>>7)]
	__obf_db477556a08e916c[23] = Encoding[(__obf_8b246d7b3659af37[14]&124)>>2]
	__obf_db477556a08e916c[24] = Encoding[((__obf_8b246d7b3659af37[14]&3)<<3)|((__obf_8b246d7b3659af37[15]&224)>>5)]
	__obf_db477556a08e916c[25] = Encoding[__obf_8b246d7b3659af37[15]&31]

	return nil
}

// Byte to index table for O(1) lookups when unmarshaling.
// We use 0xFF as sentinel value for invalid indexes.
var __obf_c91e2f0f66e52747 = [...]byte{
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
func (__obf_8b246d7b3659af37 *ULID) UnmarshalText(__obf_1caf1c7883b88b75 []byte) error {
	return __obf_c28a75a15814fc86(__obf_1caf1c7883b88b75, false, __obf_8b246d7b3659af37)
}

// Time returns the Unix time in milliseconds encoded in the ULID.
// Use the top level Time function to convert the returned value to
// a time.Time.
func (__obf_8b246d7b3659af37 ULID) Time() uint64 {
	return uint64(__obf_8b246d7b3659af37[5]) | uint64(__obf_8b246d7b3659af37[4])<<8 |
		uint64(__obf_8b246d7b3659af37[3])<<16 | uint64(__obf_8b246d7b3659af37[2])<<24 |
		uint64(__obf_8b246d7b3659af37[1])<<32 | uint64(__obf_8b246d7b3659af37[0])<<40
}

// maxTime is the maximum Unix time in milliseconds that can be
// represented in an ULID.
var __obf_ff744fe375e276cc = ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}.Time()

// MaxTime returns the maximum Unix time in milliseconds that
// can be encoded in an ULID.
func MaxTime() uint64 { return __obf_ff744fe375e276cc }

// Now is a convenience function that returns the current
// UTC time in Unix milliseconds. Equivalent to:
//
//	Timestamp(time.Now().UTC())
func Now() uint64 { return Timestamp(time.Now().UTC()) }

// Timestamp converts a time.Time to Unix milliseconds.
//
// Because of the way ULID stores time, times from the year
// 10889 produces undefined results.
func Timestamp(__obf_61139133e991449d time.Time) uint64 {
	return uint64(__obf_61139133e991449d.Unix())*1000 +
		uint64(__obf_61139133e991449d.Nanosecond()/int(time.Millisecond))
}

// Time converts Unix milliseconds in the format
// returned by the Timestamp function to a time.Time.
func Time(__obf_4450a2d1c70a6bed uint64) time.Time {
	__obf_80ba33a2bd7c6fc1 := int64(__obf_4450a2d1c70a6bed / 1e3)
	__obf_42066fe5b74e2d37 := int64((__obf_4450a2d1c70a6bed % 1e3) * 1e6)
	return time.Unix(__obf_80ba33a2bd7c6fc1,

		// SetTime sets the time component of the ULID to the given Unix time
		// in milliseconds.
		__obf_42066fe5b74e2d37)
}

func (__obf_8b246d7b3659af37 *ULID) SetTime(__obf_4450a2d1c70a6bed uint64) error {
	if __obf_4450a2d1c70a6bed > __obf_ff744fe375e276cc {
		return ErrBigTime
	}

	(*__obf_8b246d7b3659af37)[0] = byte(__obf_4450a2d1c70a6bed >> 40)
	(*__obf_8b246d7b3659af37)[1] = byte(__obf_4450a2d1c70a6bed >> 32)
	(*__obf_8b246d7b3659af37)[2] = byte(__obf_4450a2d1c70a6bed >> 24)
	(*__obf_8b246d7b3659af37)[3] = byte(__obf_4450a2d1c70a6bed >> 16)
	(*__obf_8b246d7b3659af37)[4] = byte(__obf_4450a2d1c70a6bed >> 8)
	(*__obf_8b246d7b3659af37)[5] = byte(__obf_4450a2d1c70a6bed)

	return nil
}

// Entropy returns the entropy from the ULID.
func (__obf_8b246d7b3659af37 ULID) Entropy() []byte {
	__obf_8b3554b161bd95b9 := make([]byte, 10)
	copy(__obf_8b3554b161bd95b9, __obf_8b246d7b3659af37[6:])
	return __obf_8b3554b161bd95b9
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (__obf_8b246d7b3659af37 *ULID) SetEntropy(__obf_8b3554b161bd95b9 []byte) error {
	if len(__obf_8b3554b161bd95b9) != 10 {
		return ErrDataSize
	}

	copy((*__obf_8b246d7b3659af37)[6:], __obf_8b3554b161bd95b9)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (__obf_8b246d7b3659af37 ULID) Compare(__obf_510d0067596555aa ULID) int {
	return bytes.Compare(__obf_8b246d7b3659af37[:], __obf_510d0067596555aa[:])
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (__obf_8b246d7b3659af37 *ULID) Scan(__obf_ba9a29a8774d35c1 any) error {
	switch __obf_caddb1840dda593a := __obf_ba9a29a8774d35c1.(type) {
	case nil:
		return nil
	case string:
		return __obf_8b246d7b3659af37.UnmarshalText([]byte(__obf_caddb1840dda593a))
	case []byte:
		return __obf_8b246d7b3659af37.UnmarshalBinary(__obf_caddb1840dda593a)
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
func (__obf_8b246d7b3659af37 ULID) Value() (driver.Value, error) {
	return __obf_8b246d7b3659af37.MarshalBinary()
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
func Monotonic(__obf_64686fca0f92edfe io.Reader, __obf_1259d0a556654a55 uint64) *MonotonicEntropy {
	__obf_2088693844b40541 := MonotonicEntropy{
		Reader: bufio.NewReader(__obf_64686fca0f92edfe), __obf_1259d0a556654a55: __obf_1259d0a556654a55,
	}

	if __obf_2088693844b40541.__obf_1259d0a556654a55 == 0 {
		__obf_2088693844b40541.__obf_1259d0a556654a55 = math.MaxUint32
	}

	if __obf_477df503d01c41d7, __obf_149b4fd56461a8a0 := __obf_64686fca0f92edfe.(*rand.Rand); __obf_149b4fd56461a8a0 {
		__obf_2088693844b40541.__obf_477df503d01c41d7 = __obf_477df503d01c41d7
	}

	return &__obf_2088693844b40541
}

// MonotonicEntropy is an opaque type that provides monotonic entropy.
type MonotonicEntropy struct {
	io.Reader
	__obf_4450a2d1c70a6bed uint64
	__obf_1259d0a556654a55 uint64
	__obf_64686fca0f92edfe __obf_018656cd02300c4e

	rand                   [8]byte
	__obf_477df503d01c41d7 *rand.Rand
}

// MonotonicRead implements the MonotonicReader interface.
func (__obf_2088693844b40541 *MonotonicEntropy) MonotonicRead(__obf_4450a2d1c70a6bed uint64, __obf_64686fca0f92edfe []byte) (__obf_37e8e388755917c0 error) {
	if !__obf_2088693844b40541.__obf_64686fca0f92edfe.IsZero() && __obf_2088693844b40541.__obf_4450a2d1c70a6bed == __obf_4450a2d1c70a6bed {
		__obf_37e8e388755917c0 = __obf_2088693844b40541.__obf_73f0e7d4747a1ce2()
		__obf_2088693844b40541.__obf_64686fca0f92edfe.
			AppendTo(__obf_64686fca0f92edfe)
	} else if _, __obf_37e8e388755917c0 = io.ReadFull(__obf_2088693844b40541.Reader, __obf_64686fca0f92edfe); __obf_37e8e388755917c0 == nil {
		__obf_2088693844b40541.__obf_4450a2d1c70a6bed = __obf_4450a2d1c70a6bed
		__obf_2088693844b40541.__obf_64686fca0f92edfe.
			SetBytes(__obf_64686fca0f92edfe)
	}
	return __obf_37e8e388755917c0
}

// increment the previous entropy number with a random number
// of up to m.inc (inclusive).
func (__obf_2088693844b40541 *MonotonicEntropy) __obf_73f0e7d4747a1ce2() error {
	if __obf_1259d0a556654a55, __obf_37e8e388755917c0 := __obf_2088693844b40541.__obf_f64daf515aab3567(); __obf_37e8e388755917c0 != nil {
		return __obf_37e8e388755917c0
	} else if __obf_2088693844b40541.__obf_64686fca0f92edfe.Add(__obf_1259d0a556654a55) {
		return ErrMonotonicOverflow
	}
	return nil
}

// random returns a uniform random value in [1, m.inc), reading entropy
// from m.Reader. When m.inc == 0 || m.inc == 1, it returns 1.
// Adapted from: https://golang.org/pkg/crypto/rand/#Int
func (__obf_2088693844b40541 *MonotonicEntropy) __obf_f64daf515aab3567() (__obf_1259d0a556654a55 uint64, __obf_37e8e388755917c0 error) {
	if __obf_2088693844b40541.__obf_1259d0a556654a55 <= 1 {
		return 1, nil
	}

	// Fast path for using a underlying rand.Rand directly.
	if __obf_2088693844b40541.
		// Range: [1, m.inc)
		__obf_477df503d01c41d7 != nil {

		return 1 + uint64(__obf_2088693844b40541.__obf_477df503d01c41d7.Int63n(int64(__obf_2088693844b40541.

			// bitLen is the maximum bit length needed to encode a value < m.inc.
			__obf_1259d0a556654a55))), nil
	}
	__obf_bb4d522222052bcb := bits.Len64(__obf_2088693844b40541.

		// byteLen is the maximum byte length needed to encode a value < m.inc.
		__obf_1259d0a556654a55)
	__obf_8f980283ac42ad50 := uint(__obf_bb4d522222052bcb+7) / 8
	__obf_9c48ec32fa656963 := // msbitLen is the number of bits in the most significant byte of m.inc-1.
		uint(__obf_bb4d522222052bcb % 8)
	if __obf_9c48ec32fa656963 == 0 {
		__obf_9c48ec32fa656963 = 8
	}

	for __obf_1259d0a556654a55 == 0 || __obf_1259d0a556654a55 >= __obf_2088693844b40541.__obf_1259d0a556654a55 {
		if _, __obf_37e8e388755917c0 = io.ReadFull(__obf_2088693844b40541.Reader, __obf_2088693844b40541.rand[:__obf_8f980283ac42ad50]); __obf_37e8e388755917c0 != nil {
			return 0, __obf_37e8e388755917c0
		}
		__obf_2088693844b40541.

			// Clear bits in the first byte to increase the probability
			// that the candidate is < m.inc.
			rand[0] &= uint8(int(1<<__obf_9c48ec32fa656963) - 1)

		// Convert the read bytes into an uint64 with byteLen
		// Optimized unrolled loop.
		switch __obf_8f980283ac42ad50 {
		case 1:
			__obf_1259d0a556654a55 = uint64(__obf_2088693844b40541.rand[0])
		case 2:
			__obf_1259d0a556654a55 = uint64(binary.LittleEndian.Uint16(__obf_2088693844b40541.rand[:2]))
		case 3, 4:
			__obf_1259d0a556654a55 = uint64(binary.LittleEndian.Uint32(__obf_2088693844b40541.rand[:4]))
		case 5, 6, 7, 8:
			__obf_1259d0a556654a55 = uint64(binary.LittleEndian.Uint64(__obf_2088693844b40541.rand[:8]))
		}
	}

	// Range: [1, m.inc)
	return 1 + __obf_1259d0a556654a55, nil
}

type __obf_018656cd02300c4e struct {
	Hi uint16
	Lo uint64
}

func (__obf_feb0d3963f34862a *__obf_018656cd02300c4e) SetBytes(__obf_7b61674af7e516fb []byte) {
	__obf_feb0d3963f34862a.
		Hi = binary.BigEndian.Uint16(__obf_7b61674af7e516fb[:2])
	__obf_feb0d3963f34862a.
		Lo = binary.BigEndian.Uint64(__obf_7b61674af7e516fb[2:])
}

func (__obf_feb0d3963f34862a *__obf_018656cd02300c4e) AppendTo(__obf_7b61674af7e516fb []byte) {
	binary.BigEndian.PutUint16(__obf_7b61674af7e516fb[:2], __obf_feb0d3963f34862a.Hi)
	binary.BigEndian.PutUint64(__obf_7b61674af7e516fb[2:], __obf_feb0d3963f34862a.Lo)
}

func (__obf_feb0d3963f34862a *__obf_018656cd02300c4e) Add(__obf_f88ad57451625a6d uint64) (__obf_c2ef1995f360a122 bool) {
	__obf_741b485afc35cf32, __obf_6554b5a1c945bf01 := __obf_feb0d3963f34862a.Lo, __obf_feb0d3963f34862a.Hi
	if __obf_feb0d3963f34862a.Lo += __obf_f88ad57451625a6d; __obf_feb0d3963f34862a.Lo < __obf_741b485afc35cf32 {
		__obf_feb0d3963f34862a.
			Hi++
	}
	return __obf_feb0d3963f34862a.Hi < __obf_6554b5a1c945bf01
}

func (__obf_feb0d3963f34862a __obf_018656cd02300c4e) IsZero() bool {
	return __obf_feb0d3963f34862a.Hi == 0 && __obf_feb0d3963f34862a.Lo == 0
}
