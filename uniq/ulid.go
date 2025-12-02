package __obf_34ce7ee87a5aa6b7

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
	MonotonicRead(__obf_89e5a47f9cb41daa uint64, __obf_13faf8933f955cd7 []byte) error
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
func New(__obf_89e5a47f9cb41daa uint64, __obf_f158f67066af7014 io.Reader) (__obf_f60c7174b4657d83 ULID, __obf_0a414fcc6889bc71 error) {
	if __obf_0a414fcc6889bc71 = __obf_f60c7174b4657d83.SetTime(__obf_89e5a47f9cb41daa); __obf_0a414fcc6889bc71 != nil {
		return __obf_f60c7174b4657d83, __obf_0a414fcc6889bc71
	}

	switch __obf_5663b6f25ada6000 := __obf_f158f67066af7014.(type) {
	case nil:
		return __obf_f60c7174b4657d83, __obf_0a414fcc6889bc71
	case MonotonicReader:
		__obf_0a414fcc6889bc71 = __obf_5663b6f25ada6000.MonotonicRead(__obf_89e5a47f9cb41daa, __obf_f60c7174b4657d83[6:])
	default:
		_, __obf_0a414fcc6889bc71 = io.ReadFull(__obf_5663b6f25ada6000, __obf_f60c7174b4657d83[6:])
	}

	return __obf_f60c7174b4657d83,

		// MustNew is a convenience function equivalent to New that panics on failure
		// instead of returning an error.
		__obf_0a414fcc6889bc71
}

func MustNew(__obf_89e5a47f9cb41daa uint64, __obf_f158f67066af7014 io.Reader) ULID {
	__obf_f60c7174b4657d83, __obf_0a414fcc6889bc71 := New(__obf_89e5a47f9cb41daa, __obf_f158f67066af7014)
	if __obf_0a414fcc6889bc71 != nil {
		panic(__obf_0a414fcc6889bc71)
	}
	return __obf_f60c7174b4657d83
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(__obf_37cd3c8f59b7876c string) (__obf_f60c7174b4657d83 ULID, __obf_0a414fcc6889bc71 error) {
	return __obf_f60c7174b4657d83, __obf_9b2c04010506db63([]byte(__obf_37cd3c8f59b7876c), false, &__obf_f60c7174b4657d83)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(__obf_37cd3c8f59b7876c string) (__obf_f60c7174b4657d83 ULID, __obf_0a414fcc6889bc71 error) {
	return __obf_f60c7174b4657d83, __obf_9b2c04010506db63([]byte(__obf_37cd3c8f59b7876c), true, &__obf_f60c7174b4657d83)
}

func __obf_9b2c04010506db63(__obf_29b927ea91ca599d []byte, __obf_f6b83c7df5095609 bool, __obf_f60c7174b4657d83 *ULID) error {
	// Check if a base32 encoded ULID is the right length.
	if len(__obf_29b927ea91ca599d) != EncodedSize {
		return ErrDataSize
	}

	// Check if all the characters in a base32 encoded ULID are part of the
	// expected base32 character set.
	if __obf_f6b83c7df5095609 &&
		(__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[0]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[1]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[2]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[3]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[4]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[5]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[6]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[7]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[8]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[9]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[10]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[11]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[12]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[13]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[14]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[15]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[16]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[17]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[18]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[19]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[20]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[21]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[22]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[23]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[24]] == 0xFF || __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[25]] == 0xFF) {
		return ErrInvalidCharacters
	}

	// Check if the first character in a base32 encoded ULID will overflow. This
	// happens because the base32 representation encodes 130 bits, while the
	// ULID is only 128 bits.
	//
	// See https://github.com/oklog/ulid/issues/9 for details.
	if __obf_29b927ea91ca599d[0] > '7' {
		return ErrOverflow
	}

	// Use an optimized unrolled loop (from https://github.com/RobThree/NUlid)
	// to decode a base32 ULID.

	// 6 bytes timestamp (48 bits)
	(*__obf_f60c7174b4657d83)[0] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[0]] << 5) | __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[1]]
	(*__obf_f60c7174b4657d83)[1] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[2]] << 3) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[3]] >> 2)
	(*__obf_f60c7174b4657d83)[2] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[3]] << 6) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[4]] << 1) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[5]] >> 4)
	(*__obf_f60c7174b4657d83)[3] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[5]] << 4) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[6]] >> 1)
	(*__obf_f60c7174b4657d83)[4] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[6]] << 7) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[7]] << 2) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[8]] >> 3)
	(*__obf_f60c7174b4657d83)[5] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[8]] << 5) | __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[9]]

	// 10 bytes of entropy (80 bits)
	(*__obf_f60c7174b4657d83)[6] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[10]] << 3) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[11]] >> 2)
	(*__obf_f60c7174b4657d83)[7] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[11]] << 6) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[12]] << 1) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[13]] >> 4)
	(*__obf_f60c7174b4657d83)[8] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[13]] << 4) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[14]] >> 1)
	(*__obf_f60c7174b4657d83)[9] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[14]] << 7) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[15]] << 2) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[16]] >> 3)
	(*__obf_f60c7174b4657d83)[10] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[16]] << 5) | __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[17]]
	(*__obf_f60c7174b4657d83)[11] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[18]] << 3) | __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[19]]>>2
	(*__obf_f60c7174b4657d83)[12] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[19]] << 6) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[20]] << 1) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[21]] >> 4)
	(*__obf_f60c7174b4657d83)[13] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[21]] << 4) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[22]] >> 1)
	(*__obf_f60c7174b4657d83)[14] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[22]] << 7) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[23]] << 2) | (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[24]] >> 3)
	(*__obf_f60c7174b4657d83)[15] = (__obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[24]] << 5) | __obf_f5b9c3ec9e628a53[__obf_29b927ea91ca599d[25]]

	return nil
}

// MustParse is a convenience function equivalent to Parse that panics on failure
// instead of returning an error.
func MustParse(__obf_37cd3c8f59b7876c string) ULID {
	__obf_f60c7174b4657d83, __obf_0a414fcc6889bc71 := Parse(__obf_37cd3c8f59b7876c)
	if __obf_0a414fcc6889bc71 != nil {
		panic(__obf_0a414fcc6889bc71)
	}
	return __obf_f60c7174b4657d83
}

// MustParseStrict is a convenience function equivalent to ParseStrict that
// panics on failure instead of returning an error.
func MustParseStrict(__obf_37cd3c8f59b7876c string) ULID {
	__obf_f60c7174b4657d83, __obf_0a414fcc6889bc71 := ParseStrict(__obf_37cd3c8f59b7876c)
	if __obf_0a414fcc6889bc71 != nil {
		panic(__obf_0a414fcc6889bc71)
	}
	return __obf_f60c7174b4657d83
}

// Bytes returns bytes slice representation of ULID.
func (__obf_e722acd2b7e3a316 ULID) Bytes() []byte {
	return __obf_e722acd2b7e3a316[:]
}

// String returns a lexicographically sortable string encoded ULID
// (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy
func (__obf_f60c7174b4657d83 ULID) String() string {
	__obf_37cd3c8f59b7876c := make([]byte, EncodedSize)
	_ = __obf_f60c7174b4657d83.MarshalTextTo(__obf_37cd3c8f59b7876c)
	return string(__obf_37cd3c8f59b7876c)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (__obf_f60c7174b4657d83 ULID) MarshalBinary() ([]byte, error) {
	__obf_37cd3c8f59b7876c := make([]byte, len(__obf_f60c7174b4657d83))
	return __obf_37cd3c8f59b7876c, __obf_f60c7174b4657d83.MarshalBinaryTo(__obf_37cd3c8f59b7876c)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (__obf_f60c7174b4657d83 ULID) MarshalBinaryTo(__obf_dadfec7f8fe0a296 []byte) error {
	if len(__obf_dadfec7f8fe0a296) != len(__obf_f60c7174b4657d83) {
		return ErrBufferSize
	}

	copy(__obf_dadfec7f8fe0a296, __obf_f60c7174b4657d83[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (__obf_f60c7174b4657d83 *ULID) UnmarshalBinary(__obf_80d90954bfed35d9 []byte) error {
	if len(__obf_80d90954bfed35d9) != len(*__obf_f60c7174b4657d83) {
		return ErrDataSize
	}

	copy((*__obf_f60c7174b4657d83)[:], __obf_80d90954bfed35d9)
	return nil
}

// Encoding is the base 32 encoding alphabet used in ULID strings.
const Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (__obf_f60c7174b4657d83 ULID) MarshalText() ([]byte, error) {
	__obf_37cd3c8f59b7876c := make([]byte, EncodedSize)
	return __obf_37cd3c8f59b7876c, __obf_f60c7174b4657d83.MarshalTextTo(__obf_37cd3c8f59b7876c)
}

// MarshalTextTo writes the ULID as a string to the given buffer.
// ErrBufferSize is returned when the len(dst) != 26.
func (__obf_f60c7174b4657d83 ULID) MarshalTextTo(__obf_dadfec7f8fe0a296 []byte) error {
	// Optimized unrolled loop ahead.
	// From https://github.com/RobThree/NUlid

	if len(__obf_dadfec7f8fe0a296) != EncodedSize {
		return ErrBufferSize
	}
	__obf_dadfec7f8fe0a296[ // 10 byte timestamp
	0] = Encoding[(__obf_f60c7174b4657d83[0]&224)>>5]
	__obf_dadfec7f8fe0a296[1] = Encoding[__obf_f60c7174b4657d83[0]&31]
	__obf_dadfec7f8fe0a296[2] = Encoding[(__obf_f60c7174b4657d83[1]&248)>>3]
	__obf_dadfec7f8fe0a296[3] = Encoding[((__obf_f60c7174b4657d83[1]&7)<<2)|((__obf_f60c7174b4657d83[2]&192)>>6)]
	__obf_dadfec7f8fe0a296[4] = Encoding[(__obf_f60c7174b4657d83[2]&62)>>1]
	__obf_dadfec7f8fe0a296[5] = Encoding[((__obf_f60c7174b4657d83[2]&1)<<4)|((__obf_f60c7174b4657d83[3]&240)>>4)]
	__obf_dadfec7f8fe0a296[6] = Encoding[((__obf_f60c7174b4657d83[3]&15)<<1)|((__obf_f60c7174b4657d83[4]&128)>>7)]
	__obf_dadfec7f8fe0a296[7] = Encoding[(__obf_f60c7174b4657d83[4]&124)>>2]
	__obf_dadfec7f8fe0a296[8] = Encoding[((__obf_f60c7174b4657d83[4]&3)<<3)|((__obf_f60c7174b4657d83[5]&224)>>5)]
	__obf_dadfec7f8fe0a296[9] = Encoding[__obf_f60c7174b4657d83[5]&31]
	__obf_dadfec7f8fe0a296[ // 16 bytes of entropy
	10] = Encoding[(__obf_f60c7174b4657d83[6]&248)>>3]
	__obf_dadfec7f8fe0a296[11] = Encoding[((__obf_f60c7174b4657d83[6]&7)<<2)|((__obf_f60c7174b4657d83[7]&192)>>6)]
	__obf_dadfec7f8fe0a296[12] = Encoding[(__obf_f60c7174b4657d83[7]&62)>>1]
	__obf_dadfec7f8fe0a296[13] = Encoding[((__obf_f60c7174b4657d83[7]&1)<<4)|((__obf_f60c7174b4657d83[8]&240)>>4)]
	__obf_dadfec7f8fe0a296[14] = Encoding[((__obf_f60c7174b4657d83[8]&15)<<1)|((__obf_f60c7174b4657d83[9]&128)>>7)]
	__obf_dadfec7f8fe0a296[15] = Encoding[(__obf_f60c7174b4657d83[9]&124)>>2]
	__obf_dadfec7f8fe0a296[16] = Encoding[((__obf_f60c7174b4657d83[9]&3)<<3)|((__obf_f60c7174b4657d83[10]&224)>>5)]
	__obf_dadfec7f8fe0a296[17] = Encoding[__obf_f60c7174b4657d83[10]&31]
	__obf_dadfec7f8fe0a296[18] = Encoding[(__obf_f60c7174b4657d83[11]&248)>>3]
	__obf_dadfec7f8fe0a296[19] = Encoding[((__obf_f60c7174b4657d83[11]&7)<<2)|((__obf_f60c7174b4657d83[12]&192)>>6)]
	__obf_dadfec7f8fe0a296[20] = Encoding[(__obf_f60c7174b4657d83[12]&62)>>1]
	__obf_dadfec7f8fe0a296[21] = Encoding[((__obf_f60c7174b4657d83[12]&1)<<4)|((__obf_f60c7174b4657d83[13]&240)>>4)]
	__obf_dadfec7f8fe0a296[22] = Encoding[((__obf_f60c7174b4657d83[13]&15)<<1)|((__obf_f60c7174b4657d83[14]&128)>>7)]
	__obf_dadfec7f8fe0a296[23] = Encoding[(__obf_f60c7174b4657d83[14]&124)>>2]
	__obf_dadfec7f8fe0a296[24] = Encoding[((__obf_f60c7174b4657d83[14]&3)<<3)|((__obf_f60c7174b4657d83[15]&224)>>5)]
	__obf_dadfec7f8fe0a296[25] = Encoding[__obf_f60c7174b4657d83[15]&31]

	return nil
}

// Byte to index table for O(1) lookups when unmarshaling.
// We use 0xFF as sentinel value for invalid indexes.
var __obf_f5b9c3ec9e628a53 = [...]byte{
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
func (__obf_f60c7174b4657d83 *ULID) UnmarshalText(__obf_29b927ea91ca599d []byte) error {
	return __obf_9b2c04010506db63(__obf_29b927ea91ca599d, false, __obf_f60c7174b4657d83)
}

// Time returns the Unix time in milliseconds encoded in the ULID.
// Use the top level Time function to convert the returned value to
// a time.Time.
func (__obf_f60c7174b4657d83 ULID) Time() uint64 {
	return uint64(__obf_f60c7174b4657d83[5]) | uint64(__obf_f60c7174b4657d83[4])<<8 |
		uint64(__obf_f60c7174b4657d83[3])<<16 | uint64(__obf_f60c7174b4657d83[2])<<24 |
		uint64(__obf_f60c7174b4657d83[1])<<32 | uint64(__obf_f60c7174b4657d83[0])<<40
}

// maxTime is the maximum Unix time in milliseconds that can be
// represented in an ULID.
var __obf_0290cc7f1ef93d1b = ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}.Time()

// MaxTime returns the maximum Unix time in milliseconds that
// can be encoded in an ULID.
func MaxTime() uint64 { return __obf_0290cc7f1ef93d1b }

// Now is a convenience function that returns the current
// UTC time in Unix milliseconds. Equivalent to:
//
//	Timestamp(time.Now().UTC())
func Now() uint64 { return Timestamp(time.Now().UTC()) }

// Timestamp converts a time.Time to Unix milliseconds.
//
// Because of the way ULID stores time, times from the year
// 10889 produces undefined results.
func Timestamp(__obf_438b07c9285724d4 time.Time) uint64 {
	return uint64(__obf_438b07c9285724d4.Unix())*1000 +
		uint64(__obf_438b07c9285724d4.Nanosecond()/int(time.Millisecond))
}

// Time converts Unix milliseconds in the format
// returned by the Timestamp function to a time.Time.
func Time(__obf_89e5a47f9cb41daa uint64) time.Time {
	__obf_b5366bb34e6ca8a0 := int64(__obf_89e5a47f9cb41daa / 1e3)
	__obf_53b4c9a6bc10b3ce := int64((__obf_89e5a47f9cb41daa % 1e3) * 1e6)
	return time.Unix(__obf_b5366bb34e6ca8a0,

		// SetTime sets the time component of the ULID to the given Unix time
		// in milliseconds.
		__obf_53b4c9a6bc10b3ce)
}

func (__obf_f60c7174b4657d83 *ULID) SetTime(__obf_89e5a47f9cb41daa uint64) error {
	if __obf_89e5a47f9cb41daa > __obf_0290cc7f1ef93d1b {
		return ErrBigTime
	}

	(*__obf_f60c7174b4657d83)[0] = byte(__obf_89e5a47f9cb41daa >> 40)
	(*__obf_f60c7174b4657d83)[1] = byte(__obf_89e5a47f9cb41daa >> 32)
	(*__obf_f60c7174b4657d83)[2] = byte(__obf_89e5a47f9cb41daa >> 24)
	(*__obf_f60c7174b4657d83)[3] = byte(__obf_89e5a47f9cb41daa >> 16)
	(*__obf_f60c7174b4657d83)[4] = byte(__obf_89e5a47f9cb41daa >> 8)
	(*__obf_f60c7174b4657d83)[5] = byte(__obf_89e5a47f9cb41daa)

	return nil
}

// Entropy returns the entropy from the ULID.
func (__obf_f60c7174b4657d83 ULID) Entropy() []byte {
	__obf_5663b6f25ada6000 := make([]byte, 10)
	copy(__obf_5663b6f25ada6000, __obf_f60c7174b4657d83[6:])
	return __obf_5663b6f25ada6000
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (__obf_f60c7174b4657d83 *ULID) SetEntropy(__obf_5663b6f25ada6000 []byte) error {
	if len(__obf_5663b6f25ada6000) != 10 {
		return ErrDataSize
	}

	copy((*__obf_f60c7174b4657d83)[6:], __obf_5663b6f25ada6000)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (__obf_f60c7174b4657d83 ULID) Compare(__obf_5a9e2c4080fdae80 ULID) int {
	return bytes.Compare(__obf_f60c7174b4657d83[:], __obf_5a9e2c4080fdae80[:])
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (__obf_f60c7174b4657d83 *ULID) Scan(__obf_1e0008c183aadaf4 any) error {
	switch __obf_50f57aac3b0e8f3e := __obf_1e0008c183aadaf4.(type) {
	case nil:
		return nil
	case string:
		return __obf_f60c7174b4657d83.UnmarshalText([]byte(__obf_50f57aac3b0e8f3e))
	case []byte:
		return __obf_f60c7174b4657d83.UnmarshalBinary(__obf_50f57aac3b0e8f3e)
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
func (__obf_f60c7174b4657d83 ULID) Value() (driver.Value, error) {
	return __obf_f60c7174b4657d83.MarshalBinary()
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
func Monotonic(__obf_f158f67066af7014 io.Reader, __obf_d019268c430a693c uint64) *MonotonicEntropy {
	__obf_a7062b4c5c8129e1 := MonotonicEntropy{
		Reader: bufio.NewReader(__obf_f158f67066af7014), __obf_d019268c430a693c: __obf_d019268c430a693c,
	}

	if __obf_a7062b4c5c8129e1.__obf_d019268c430a693c == 0 {
		__obf_a7062b4c5c8129e1.__obf_d019268c430a693c = math.MaxUint32
	}

	if __obf_427e7770a91e7063, __obf_04415e5e48fee761 := __obf_f158f67066af7014.(*rand.Rand); __obf_04415e5e48fee761 {
		__obf_a7062b4c5c8129e1.__obf_427e7770a91e7063 = __obf_427e7770a91e7063
	}

	return &__obf_a7062b4c5c8129e1
}

// MonotonicEntropy is an opaque type that provides monotonic entropy.
type MonotonicEntropy struct {
	io.Reader
	__obf_89e5a47f9cb41daa uint64
	__obf_d019268c430a693c uint64
	__obf_f158f67066af7014 __obf_d8d3c5d9223730c1

	rand                   [8]byte
	__obf_427e7770a91e7063 *rand.Rand
}

// MonotonicRead implements the MonotonicReader interface.
func (__obf_a7062b4c5c8129e1 *MonotonicEntropy) MonotonicRead(__obf_89e5a47f9cb41daa uint64, __obf_f158f67066af7014 []byte) (__obf_0a414fcc6889bc71 error) {
	if !__obf_a7062b4c5c8129e1.__obf_f158f67066af7014.IsZero() && __obf_a7062b4c5c8129e1.__obf_89e5a47f9cb41daa == __obf_89e5a47f9cb41daa {
		__obf_0a414fcc6889bc71 = __obf_a7062b4c5c8129e1.__obf_ecce1ad0ec0fb048()
		__obf_a7062b4c5c8129e1.__obf_f158f67066af7014.
			AppendTo(__obf_f158f67066af7014)
	} else if _, __obf_0a414fcc6889bc71 = io.ReadFull(__obf_a7062b4c5c8129e1.Reader, __obf_f158f67066af7014); __obf_0a414fcc6889bc71 == nil {
		__obf_a7062b4c5c8129e1.__obf_89e5a47f9cb41daa = __obf_89e5a47f9cb41daa
		__obf_a7062b4c5c8129e1.__obf_f158f67066af7014.
			SetBytes(__obf_f158f67066af7014)
	}
	return __obf_0a414fcc6889bc71
}

// increment the previous entropy number with a random number
// of up to m.inc (inclusive).
func (__obf_a7062b4c5c8129e1 *MonotonicEntropy) __obf_ecce1ad0ec0fb048() error {
	if __obf_d019268c430a693c, __obf_0a414fcc6889bc71 := __obf_a7062b4c5c8129e1.__obf_83b908495b56e469(); __obf_0a414fcc6889bc71 != nil {
		return __obf_0a414fcc6889bc71
	} else if __obf_a7062b4c5c8129e1.__obf_f158f67066af7014.Add(__obf_d019268c430a693c) {
		return ErrMonotonicOverflow
	}
	return nil
}

// random returns a uniform random value in [1, m.inc), reading entropy
// from m.Reader. When m.inc == 0 || m.inc == 1, it returns 1.
// Adapted from: https://golang.org/pkg/crypto/rand/#Int
func (__obf_a7062b4c5c8129e1 *MonotonicEntropy) __obf_83b908495b56e469() (__obf_d019268c430a693c uint64, __obf_0a414fcc6889bc71 error) {
	if __obf_a7062b4c5c8129e1.__obf_d019268c430a693c <= 1 {
		return 1, nil
	}

	// Fast path for using a underlying rand.Rand directly.
	if __obf_a7062b4c5c8129e1.
		// Range: [1, m.inc)
		__obf_427e7770a91e7063 != nil {

		return 1 + uint64(__obf_a7062b4c5c8129e1.__obf_427e7770a91e7063.Int63n(int64(__obf_a7062b4c5c8129e1.

			// bitLen is the maximum bit length needed to encode a value < m.inc.
			__obf_d019268c430a693c))), nil
	}
	__obf_27a28c9913482463 := bits.Len64(__obf_a7062b4c5c8129e1.

		// byteLen is the maximum byte length needed to encode a value < m.inc.
		__obf_d019268c430a693c)
	__obf_c7eaf344aa788ae1 := uint(__obf_27a28c9913482463+7) / 8
	__obf_1d356b8a3b21dfdd := // msbitLen is the number of bits in the most significant byte of m.inc-1.
		uint(__obf_27a28c9913482463 % 8)
	if __obf_1d356b8a3b21dfdd == 0 {
		__obf_1d356b8a3b21dfdd = 8
	}

	for __obf_d019268c430a693c == 0 || __obf_d019268c430a693c >= __obf_a7062b4c5c8129e1.__obf_d019268c430a693c {
		if _, __obf_0a414fcc6889bc71 = io.ReadFull(__obf_a7062b4c5c8129e1.Reader, __obf_a7062b4c5c8129e1.rand[:__obf_c7eaf344aa788ae1]); __obf_0a414fcc6889bc71 != nil {
			return 0, __obf_0a414fcc6889bc71
		}
		__obf_a7062b4c5c8129e1.

			// Clear bits in the first byte to increase the probability
			// that the candidate is < m.inc.
			rand[0] &= uint8(int(1<<__obf_1d356b8a3b21dfdd) - 1)

		// Convert the read bytes into an uint64 with byteLen
		// Optimized unrolled loop.
		switch __obf_c7eaf344aa788ae1 {
		case 1:
			__obf_d019268c430a693c = uint64(__obf_a7062b4c5c8129e1.rand[0])
		case 2:
			__obf_d019268c430a693c = uint64(binary.LittleEndian.Uint16(__obf_a7062b4c5c8129e1.rand[:2]))
		case 3, 4:
			__obf_d019268c430a693c = uint64(binary.LittleEndian.Uint32(__obf_a7062b4c5c8129e1.rand[:4]))
		case 5, 6, 7, 8:
			__obf_d019268c430a693c = uint64(binary.LittleEndian.Uint64(__obf_a7062b4c5c8129e1.rand[:8]))
		}
	}

	// Range: [1, m.inc)
	return 1 + __obf_d019268c430a693c, nil
}

type __obf_d8d3c5d9223730c1 struct {
	Hi uint16
	Lo uint64
}

func (__obf_e722acd2b7e3a316 *__obf_d8d3c5d9223730c1) SetBytes(__obf_78da9f19c233d32f []byte) {
	__obf_e722acd2b7e3a316.
		Hi = binary.BigEndian.Uint16(__obf_78da9f19c233d32f[:2])
	__obf_e722acd2b7e3a316.
		Lo = binary.BigEndian.Uint64(__obf_78da9f19c233d32f[2:])
}

func (__obf_e722acd2b7e3a316 *__obf_d8d3c5d9223730c1) AppendTo(__obf_78da9f19c233d32f []byte) {
	binary.BigEndian.PutUint16(__obf_78da9f19c233d32f[:2], __obf_e722acd2b7e3a316.Hi)
	binary.BigEndian.PutUint64(__obf_78da9f19c233d32f[2:], __obf_e722acd2b7e3a316.Lo)
}

func (__obf_e722acd2b7e3a316 *__obf_d8d3c5d9223730c1) Add(__obf_0d08165857356423 uint64) (__obf_10058363e32de1ec bool) {
	__obf_f657a52dfbcac426, __obf_d36e40b0dda774aa := __obf_e722acd2b7e3a316.Lo, __obf_e722acd2b7e3a316.Hi
	if __obf_e722acd2b7e3a316.Lo += __obf_0d08165857356423; __obf_e722acd2b7e3a316.Lo < __obf_f657a52dfbcac426 {
		__obf_e722acd2b7e3a316.
			Hi++
	}
	return __obf_e722acd2b7e3a316.Hi < __obf_d36e40b0dda774aa
}

func (__obf_e722acd2b7e3a316 __obf_d8d3c5d9223730c1) IsZero() bool {
	return __obf_e722acd2b7e3a316.Hi == 0 && __obf_e722acd2b7e3a316.Lo == 0
}
