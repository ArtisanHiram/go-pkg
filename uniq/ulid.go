package __obf_0f0e0d1ff72f3ff0

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
	MonotonicRead(__obf_244c56be1ba55e39 uint64, __obf_2ba781a6f45d6c2f []byte) error
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
func New(__obf_244c56be1ba55e39 uint64, __obf_70a59ee32a73e7a9 io.Reader) (__obf_a7593a31c69c5d40 ULID, __obf_f24ec7ae16791add error) {
	if __obf_f24ec7ae16791add = __obf_a7593a31c69c5d40.SetTime(__obf_244c56be1ba55e39); __obf_f24ec7ae16791add != nil {
		return __obf_a7593a31c69c5d40, __obf_f24ec7ae16791add
	}

	switch __obf_c1a5e325c87c081f := __obf_70a59ee32a73e7a9.(type) {
	case nil:
		return __obf_a7593a31c69c5d40, __obf_f24ec7ae16791add
	case MonotonicReader:
		__obf_f24ec7ae16791add = __obf_c1a5e325c87c081f.MonotonicRead(__obf_244c56be1ba55e39, __obf_a7593a31c69c5d40[6:])
	default:
		_, __obf_f24ec7ae16791add = io.ReadFull(__obf_c1a5e325c87c081f, __obf_a7593a31c69c5d40[6:])
	}

	return __obf_a7593a31c69c5d40, __obf_f24ec7ae16791add
}

// MustNew is a convenience function equivalent to New that panics on failure
// instead of returning an error.
func MustNew(__obf_244c56be1ba55e39 uint64, __obf_70a59ee32a73e7a9 io.Reader) ULID {
	__obf_a7593a31c69c5d40, __obf_f24ec7ae16791add := New(__obf_244c56be1ba55e39, __obf_70a59ee32a73e7a9)
	if __obf_f24ec7ae16791add != nil {
		panic(__obf_f24ec7ae16791add)
	}
	return __obf_a7593a31c69c5d40
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(__obf_c0e4b2d2a07c675d string) (__obf_a7593a31c69c5d40 ULID, __obf_f24ec7ae16791add error) {
	return __obf_a7593a31c69c5d40, __obf_28dbd2755f0948c8([]byte(__obf_c0e4b2d2a07c675d), false, &__obf_a7593a31c69c5d40)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(__obf_c0e4b2d2a07c675d string) (__obf_a7593a31c69c5d40 ULID, __obf_f24ec7ae16791add error) {
	return __obf_a7593a31c69c5d40, __obf_28dbd2755f0948c8([]byte(__obf_c0e4b2d2a07c675d), true, &__obf_a7593a31c69c5d40)
}

func __obf_28dbd2755f0948c8(__obf_ffc5d2d602752af4 []byte, __obf_fc0be1da03ed855a bool, __obf_a7593a31c69c5d40 *ULID) error {
	// Check if a base32 encoded ULID is the right length.
	if len(__obf_ffc5d2d602752af4) != EncodedSize {
		return ErrDataSize
	}

	// Check if all the characters in a base32 encoded ULID are part of the
	// expected base32 character set.
	if __obf_fc0be1da03ed855a &&
		(__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[0]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[1]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[2]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[3]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[4]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[5]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[6]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[7]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[8]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[9]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[10]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[11]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[12]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[13]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[14]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[15]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[16]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[17]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[18]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[19]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[20]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[21]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[22]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[23]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[24]] == 0xFF ||
			__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[25]] == 0xFF) {
		return ErrInvalidCharacters
	}

	// Check if the first character in a base32 encoded ULID will overflow. This
	// happens because the base32 representation encodes 130 bits, while the
	// ULID is only 128 bits.
	//
	// See https://github.com/oklog/ulid/issues/9 for details.
	if __obf_ffc5d2d602752af4[0] > '7' {
		return ErrOverflow
	}

	// Use an optimized unrolled loop (from https://github.com/RobThree/NUlid)
	// to decode a base32 ULID.

	// 6 bytes timestamp (48 bits)
	(*__obf_a7593a31c69c5d40)[0] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[0]] << 5) | __obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[1]]
	(*__obf_a7593a31c69c5d40)[1] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[2]] << 3) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[3]] >> 2)
	(*__obf_a7593a31c69c5d40)[2] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[3]] << 6) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[4]] << 1) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[5]] >> 4)
	(*__obf_a7593a31c69c5d40)[3] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[5]] << 4) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[6]] >> 1)
	(*__obf_a7593a31c69c5d40)[4] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[6]] << 7) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[7]] << 2) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[8]] >> 3)
	(*__obf_a7593a31c69c5d40)[5] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[8]] << 5) | __obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[9]]

	// 10 bytes of entropy (80 bits)
	(*__obf_a7593a31c69c5d40)[6] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[10]] << 3) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[11]] >> 2)
	(*__obf_a7593a31c69c5d40)[7] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[11]] << 6) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[12]] << 1) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[13]] >> 4)
	(*__obf_a7593a31c69c5d40)[8] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[13]] << 4) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[14]] >> 1)
	(*__obf_a7593a31c69c5d40)[9] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[14]] << 7) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[15]] << 2) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[16]] >> 3)
	(*__obf_a7593a31c69c5d40)[10] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[16]] << 5) | __obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[17]]
	(*__obf_a7593a31c69c5d40)[11] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[18]] << 3) | __obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[19]]>>2
	(*__obf_a7593a31c69c5d40)[12] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[19]] << 6) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[20]] << 1) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[21]] >> 4)
	(*__obf_a7593a31c69c5d40)[13] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[21]] << 4) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[22]] >> 1)
	(*__obf_a7593a31c69c5d40)[14] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[22]] << 7) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[23]] << 2) | (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[24]] >> 3)
	(*__obf_a7593a31c69c5d40)[15] = (__obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[24]] << 5) | __obf_02e2cbf089bd81b5[__obf_ffc5d2d602752af4[25]]

	return nil
}

// MustParse is a convenience function equivalent to Parse that panics on failure
// instead of returning an error.
func MustParse(__obf_c0e4b2d2a07c675d string) ULID {
	__obf_a7593a31c69c5d40, __obf_f24ec7ae16791add := Parse(__obf_c0e4b2d2a07c675d)
	if __obf_f24ec7ae16791add != nil {
		panic(__obf_f24ec7ae16791add)
	}
	return __obf_a7593a31c69c5d40
}

// MustParseStrict is a convenience function equivalent to ParseStrict that
// panics on failure instead of returning an error.
func MustParseStrict(__obf_c0e4b2d2a07c675d string) ULID {
	__obf_a7593a31c69c5d40, __obf_f24ec7ae16791add := ParseStrict(__obf_c0e4b2d2a07c675d)
	if __obf_f24ec7ae16791add != nil {
		panic(__obf_f24ec7ae16791add)
	}
	return __obf_a7593a31c69c5d40
}

// Bytes returns bytes slice representation of ULID.
func (__obf_c5c91f357c087ad4 ULID) Bytes() []byte {
	return __obf_c5c91f357c087ad4[:]
}

// String returns a lexicographically sortable string encoded ULID
// (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy
func (__obf_a7593a31c69c5d40 ULID) String() string {
	__obf_c0e4b2d2a07c675d := make([]byte, EncodedSize)
	_ = __obf_a7593a31c69c5d40.MarshalTextTo(__obf_c0e4b2d2a07c675d)
	return string(__obf_c0e4b2d2a07c675d)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (__obf_a7593a31c69c5d40 ULID) MarshalBinary() ([]byte, error) {
	__obf_c0e4b2d2a07c675d := make([]byte, len(__obf_a7593a31c69c5d40))
	return __obf_c0e4b2d2a07c675d, __obf_a7593a31c69c5d40.MarshalBinaryTo(__obf_c0e4b2d2a07c675d)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (__obf_a7593a31c69c5d40 ULID) MarshalBinaryTo(__obf_7a9042af3f7b15ad []byte) error {
	if len(__obf_7a9042af3f7b15ad) != len(__obf_a7593a31c69c5d40) {
		return ErrBufferSize
	}

	copy(__obf_7a9042af3f7b15ad, __obf_a7593a31c69c5d40[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (__obf_a7593a31c69c5d40 *ULID) UnmarshalBinary(__obf_98cef0dd2de60ca8 []byte) error {
	if len(__obf_98cef0dd2de60ca8) != len(*__obf_a7593a31c69c5d40) {
		return ErrDataSize
	}

	copy((*__obf_a7593a31c69c5d40)[:], __obf_98cef0dd2de60ca8)
	return nil
}

// Encoding is the base 32 encoding alphabet used in ULID strings.
const Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (__obf_a7593a31c69c5d40 ULID) MarshalText() ([]byte, error) {
	__obf_c0e4b2d2a07c675d := make([]byte, EncodedSize)
	return __obf_c0e4b2d2a07c675d, __obf_a7593a31c69c5d40.MarshalTextTo(__obf_c0e4b2d2a07c675d)
}

// MarshalTextTo writes the ULID as a string to the given buffer.
// ErrBufferSize is returned when the len(dst) != 26.
func (__obf_a7593a31c69c5d40 ULID) MarshalTextTo(__obf_7a9042af3f7b15ad []byte) error {
	// Optimized unrolled loop ahead.
	// From https://github.com/RobThree/NUlid

	if len(__obf_7a9042af3f7b15ad) != EncodedSize {
		return ErrBufferSize
	}

	// 10 byte timestamp
	__obf_7a9042af3f7b15ad[0] = Encoding[(__obf_a7593a31c69c5d40[0]&224)>>5]
	__obf_7a9042af3f7b15ad[1] = Encoding[__obf_a7593a31c69c5d40[0]&31]
	__obf_7a9042af3f7b15ad[2] = Encoding[(__obf_a7593a31c69c5d40[1]&248)>>3]
	__obf_7a9042af3f7b15ad[3] = Encoding[((__obf_a7593a31c69c5d40[1]&7)<<2)|((__obf_a7593a31c69c5d40[2]&192)>>6)]
	__obf_7a9042af3f7b15ad[4] = Encoding[(__obf_a7593a31c69c5d40[2]&62)>>1]
	__obf_7a9042af3f7b15ad[5] = Encoding[((__obf_a7593a31c69c5d40[2]&1)<<4)|((__obf_a7593a31c69c5d40[3]&240)>>4)]
	__obf_7a9042af3f7b15ad[6] = Encoding[((__obf_a7593a31c69c5d40[3]&15)<<1)|((__obf_a7593a31c69c5d40[4]&128)>>7)]
	__obf_7a9042af3f7b15ad[7] = Encoding[(__obf_a7593a31c69c5d40[4]&124)>>2]
	__obf_7a9042af3f7b15ad[8] = Encoding[((__obf_a7593a31c69c5d40[4]&3)<<3)|((__obf_a7593a31c69c5d40[5]&224)>>5)]
	__obf_7a9042af3f7b15ad[9] = Encoding[__obf_a7593a31c69c5d40[5]&31]

	// 16 bytes of entropy
	__obf_7a9042af3f7b15ad[10] = Encoding[(__obf_a7593a31c69c5d40[6]&248)>>3]
	__obf_7a9042af3f7b15ad[11] = Encoding[((__obf_a7593a31c69c5d40[6]&7)<<2)|((__obf_a7593a31c69c5d40[7]&192)>>6)]
	__obf_7a9042af3f7b15ad[12] = Encoding[(__obf_a7593a31c69c5d40[7]&62)>>1]
	__obf_7a9042af3f7b15ad[13] = Encoding[((__obf_a7593a31c69c5d40[7]&1)<<4)|((__obf_a7593a31c69c5d40[8]&240)>>4)]
	__obf_7a9042af3f7b15ad[14] = Encoding[((__obf_a7593a31c69c5d40[8]&15)<<1)|((__obf_a7593a31c69c5d40[9]&128)>>7)]
	__obf_7a9042af3f7b15ad[15] = Encoding[(__obf_a7593a31c69c5d40[9]&124)>>2]
	__obf_7a9042af3f7b15ad[16] = Encoding[((__obf_a7593a31c69c5d40[9]&3)<<3)|((__obf_a7593a31c69c5d40[10]&224)>>5)]
	__obf_7a9042af3f7b15ad[17] = Encoding[__obf_a7593a31c69c5d40[10]&31]
	__obf_7a9042af3f7b15ad[18] = Encoding[(__obf_a7593a31c69c5d40[11]&248)>>3]
	__obf_7a9042af3f7b15ad[19] = Encoding[((__obf_a7593a31c69c5d40[11]&7)<<2)|((__obf_a7593a31c69c5d40[12]&192)>>6)]
	__obf_7a9042af3f7b15ad[20] = Encoding[(__obf_a7593a31c69c5d40[12]&62)>>1]
	__obf_7a9042af3f7b15ad[21] = Encoding[((__obf_a7593a31c69c5d40[12]&1)<<4)|((__obf_a7593a31c69c5d40[13]&240)>>4)]
	__obf_7a9042af3f7b15ad[22] = Encoding[((__obf_a7593a31c69c5d40[13]&15)<<1)|((__obf_a7593a31c69c5d40[14]&128)>>7)]
	__obf_7a9042af3f7b15ad[23] = Encoding[(__obf_a7593a31c69c5d40[14]&124)>>2]
	__obf_7a9042af3f7b15ad[24] = Encoding[((__obf_a7593a31c69c5d40[14]&3)<<3)|((__obf_a7593a31c69c5d40[15]&224)>>5)]
	__obf_7a9042af3f7b15ad[25] = Encoding[__obf_a7593a31c69c5d40[15]&31]

	return nil
}

// Byte to index table for O(1) lookups when unmarshaling.
// We use 0xFF as sentinel value for invalid indexes.
var __obf_02e2cbf089bd81b5 = [...]byte{
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
func (__obf_a7593a31c69c5d40 *ULID) UnmarshalText(__obf_ffc5d2d602752af4 []byte) error {
	return __obf_28dbd2755f0948c8(__obf_ffc5d2d602752af4, false, __obf_a7593a31c69c5d40)
}

// Time returns the Unix time in milliseconds encoded in the ULID.
// Use the top level Time function to convert the returned value to
// a time.Time.
func (__obf_a7593a31c69c5d40 ULID) Time() uint64 {
	return uint64(__obf_a7593a31c69c5d40[5]) | uint64(__obf_a7593a31c69c5d40[4])<<8 |
		uint64(__obf_a7593a31c69c5d40[3])<<16 | uint64(__obf_a7593a31c69c5d40[2])<<24 |
		uint64(__obf_a7593a31c69c5d40[1])<<32 | uint64(__obf_a7593a31c69c5d40[0])<<40
}

// maxTime is the maximum Unix time in milliseconds that can be
// represented in an ULID.
var __obf_172f50847afdadbe = ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}.Time()

// MaxTime returns the maximum Unix time in milliseconds that
// can be encoded in an ULID.
func MaxTime() uint64 { return __obf_172f50847afdadbe }

// Now is a convenience function that returns the current
// UTC time in Unix milliseconds. Equivalent to:
//
//	Timestamp(time.Now().UTC())
func Now() uint64 { return Timestamp(time.Now().UTC()) }

// Timestamp converts a time.Time to Unix milliseconds.
//
// Because of the way ULID stores time, times from the year
// 10889 produces undefined results.
func Timestamp(__obf_3c7f4fcc21cb9f0d time.Time) uint64 {
	return uint64(__obf_3c7f4fcc21cb9f0d.Unix())*1000 +
		uint64(__obf_3c7f4fcc21cb9f0d.Nanosecond()/int(time.Millisecond))
}

// Time converts Unix milliseconds in the format
// returned by the Timestamp function to a time.Time.
func Time(__obf_244c56be1ba55e39 uint64) time.Time {
	__obf_d7da0d0479be92d1 := int64(__obf_244c56be1ba55e39 / 1e3)
	__obf_ea033946bfbc5404 := int64((__obf_244c56be1ba55e39 % 1e3) * 1e6)
	return time.Unix(__obf_d7da0d0479be92d1, __obf_ea033946bfbc5404)
}

// SetTime sets the time component of the ULID to the given Unix time
// in milliseconds.
func (__obf_a7593a31c69c5d40 *ULID) SetTime(__obf_244c56be1ba55e39 uint64) error {
	if __obf_244c56be1ba55e39 > __obf_172f50847afdadbe {
		return ErrBigTime
	}

	(*__obf_a7593a31c69c5d40)[0] = byte(__obf_244c56be1ba55e39 >> 40)
	(*__obf_a7593a31c69c5d40)[1] = byte(__obf_244c56be1ba55e39 >> 32)
	(*__obf_a7593a31c69c5d40)[2] = byte(__obf_244c56be1ba55e39 >> 24)
	(*__obf_a7593a31c69c5d40)[3] = byte(__obf_244c56be1ba55e39 >> 16)
	(*__obf_a7593a31c69c5d40)[4] = byte(__obf_244c56be1ba55e39 >> 8)
	(*__obf_a7593a31c69c5d40)[5] = byte(__obf_244c56be1ba55e39)

	return nil
}

// Entropy returns the entropy from the ULID.
func (__obf_a7593a31c69c5d40 ULID) Entropy() []byte {
	__obf_c1a5e325c87c081f := make([]byte, 10)
	copy(__obf_c1a5e325c87c081f, __obf_a7593a31c69c5d40[6:])
	return __obf_c1a5e325c87c081f
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (__obf_a7593a31c69c5d40 *ULID) SetEntropy(__obf_c1a5e325c87c081f []byte) error {
	if len(__obf_c1a5e325c87c081f) != 10 {
		return ErrDataSize
	}

	copy((*__obf_a7593a31c69c5d40)[6:], __obf_c1a5e325c87c081f)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (__obf_a7593a31c69c5d40 ULID) Compare(__obf_fd36e20437fb2a26 ULID) int {
	return bytes.Compare(__obf_a7593a31c69c5d40[:], __obf_fd36e20437fb2a26[:])
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (__obf_a7593a31c69c5d40 *ULID) Scan(__obf_51e714206ac706ad any) error {
	switch __obf_09cf9c38ec453181 := __obf_51e714206ac706ad.(type) {
	case nil:
		return nil
	case string:
		return __obf_a7593a31c69c5d40.UnmarshalText([]byte(__obf_09cf9c38ec453181))
	case []byte:
		return __obf_a7593a31c69c5d40.UnmarshalBinary(__obf_09cf9c38ec453181)
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
func (__obf_a7593a31c69c5d40 ULID) Value() (driver.Value, error) {
	return __obf_a7593a31c69c5d40.MarshalBinary()
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
func Monotonic(__obf_70a59ee32a73e7a9 io.Reader, __obf_ec8d1078da56dbff uint64) *MonotonicEntropy {
	__obf_f0ff2b61c246135b := MonotonicEntropy{
		Reader:                 bufio.NewReader(__obf_70a59ee32a73e7a9),
		__obf_ec8d1078da56dbff: __obf_ec8d1078da56dbff,
	}

	if __obf_f0ff2b61c246135b.__obf_ec8d1078da56dbff == 0 {
		__obf_f0ff2b61c246135b.__obf_ec8d1078da56dbff = math.MaxUint32
	}

	if __obf_b5931328e9a79d42, __obf_b08eebaa4bd2dee3 := __obf_70a59ee32a73e7a9.(*rand.Rand); __obf_b08eebaa4bd2dee3 {
		__obf_f0ff2b61c246135b.__obf_b5931328e9a79d42 = __obf_b5931328e9a79d42
	}

	return &__obf_f0ff2b61c246135b
}

// MonotonicEntropy is an opaque type that provides monotonic entropy.
type MonotonicEntropy struct {
	io.Reader
	__obf_244c56be1ba55e39 uint64
	__obf_ec8d1078da56dbff uint64
	__obf_70a59ee32a73e7a9 __obf_82cd85fd15cb8e6a
	rand                   [8]byte
	__obf_b5931328e9a79d42 *rand.Rand
}

// MonotonicRead implements the MonotonicReader interface.
func (__obf_f0ff2b61c246135b *MonotonicEntropy) MonotonicRead(__obf_244c56be1ba55e39 uint64, __obf_70a59ee32a73e7a9 []byte) (__obf_f24ec7ae16791add error) {
	if !__obf_f0ff2b61c246135b.__obf_70a59ee32a73e7a9.IsZero() && __obf_f0ff2b61c246135b.__obf_244c56be1ba55e39 == __obf_244c56be1ba55e39 {
		__obf_f24ec7ae16791add = __obf_f0ff2b61c246135b.__obf_8967a2c6f920600b()
		__obf_f0ff2b61c246135b.__obf_70a59ee32a73e7a9.AppendTo(__obf_70a59ee32a73e7a9)
	} else if _, __obf_f24ec7ae16791add = io.ReadFull(__obf_f0ff2b61c246135b.Reader, __obf_70a59ee32a73e7a9); __obf_f24ec7ae16791add == nil {
		__obf_f0ff2b61c246135b.__obf_244c56be1ba55e39 = __obf_244c56be1ba55e39
		__obf_f0ff2b61c246135b.__obf_70a59ee32a73e7a9.SetBytes(__obf_70a59ee32a73e7a9)
	}
	return __obf_f24ec7ae16791add
}

// increment the previous entropy number with a random number
// of up to m.inc (inclusive).
func (__obf_f0ff2b61c246135b *MonotonicEntropy) __obf_8967a2c6f920600b() error {
	if __obf_ec8d1078da56dbff, __obf_f24ec7ae16791add := __obf_f0ff2b61c246135b.__obf_943c7074017a21a5(); __obf_f24ec7ae16791add != nil {
		return __obf_f24ec7ae16791add
	} else if __obf_f0ff2b61c246135b.__obf_70a59ee32a73e7a9.Add(__obf_ec8d1078da56dbff) {
		return ErrMonotonicOverflow
	}
	return nil
}

// random returns a uniform random value in [1, m.inc), reading entropy
// from m.Reader. When m.inc == 0 || m.inc == 1, it returns 1.
// Adapted from: https://golang.org/pkg/crypto/rand/#Int
func (__obf_f0ff2b61c246135b *MonotonicEntropy) __obf_943c7074017a21a5() (__obf_ec8d1078da56dbff uint64, __obf_f24ec7ae16791add error) {
	if __obf_f0ff2b61c246135b.__obf_ec8d1078da56dbff <= 1 {
		return 1, nil
	}

	// Fast path for using a underlying rand.Rand directly.
	if __obf_f0ff2b61c246135b.__obf_b5931328e9a79d42 != nil {
		// Range: [1, m.inc)
		return 1 + uint64(__obf_f0ff2b61c246135b.__obf_b5931328e9a79d42.Int63n(int64(__obf_f0ff2b61c246135b.__obf_ec8d1078da56dbff))), nil
	}

	// bitLen is the maximum bit length needed to encode a value < m.inc.
	__obf_555a139c311bfb42 := bits.Len64(__obf_f0ff2b61c246135b.__obf_ec8d1078da56dbff)

	// byteLen is the maximum byte length needed to encode a value < m.inc.
	__obf_7638dffa57f3d747 := uint(__obf_555a139c311bfb42+7) / 8

	// msbitLen is the number of bits in the most significant byte of m.inc-1.
	__obf_d4d803e6caaeff89 := uint(__obf_555a139c311bfb42 % 8)
	if __obf_d4d803e6caaeff89 == 0 {
		__obf_d4d803e6caaeff89 = 8
	}

	for __obf_ec8d1078da56dbff == 0 || __obf_ec8d1078da56dbff >= __obf_f0ff2b61c246135b.__obf_ec8d1078da56dbff {
		if _, __obf_f24ec7ae16791add = io.ReadFull(__obf_f0ff2b61c246135b.Reader, __obf_f0ff2b61c246135b.rand[:__obf_7638dffa57f3d747]); __obf_f24ec7ae16791add != nil {
			return 0, __obf_f24ec7ae16791add
		}

		// Clear bits in the first byte to increase the probability
		// that the candidate is < m.inc.
		__obf_f0ff2b61c246135b.rand[0] &= uint8(int(1<<__obf_d4d803e6caaeff89) - 1)

		// Convert the read bytes into an uint64 with byteLen
		// Optimized unrolled loop.
		switch __obf_7638dffa57f3d747 {
		case 1:
			__obf_ec8d1078da56dbff = uint64(__obf_f0ff2b61c246135b.rand[0])
		case 2:
			__obf_ec8d1078da56dbff = uint64(binary.LittleEndian.Uint16(__obf_f0ff2b61c246135b.rand[:2]))
		case 3, 4:
			__obf_ec8d1078da56dbff = uint64(binary.LittleEndian.Uint32(__obf_f0ff2b61c246135b.rand[:4]))
		case 5, 6, 7, 8:
			__obf_ec8d1078da56dbff = uint64(binary.LittleEndian.Uint64(__obf_f0ff2b61c246135b.rand[:8]))
		}
	}

	// Range: [1, m.inc)
	return 1 + __obf_ec8d1078da56dbff, nil
}

type __obf_82cd85fd15cb8e6a struct {
	Hi uint16
	Lo uint64
}

func (__obf_c5c91f357c087ad4 *__obf_82cd85fd15cb8e6a) SetBytes(__obf_6068ae77f3b0e747 []byte) {
	__obf_c5c91f357c087ad4.Hi = binary.BigEndian.Uint16(__obf_6068ae77f3b0e747[:2])
	__obf_c5c91f357c087ad4.Lo = binary.BigEndian.Uint64(__obf_6068ae77f3b0e747[2:])
}

func (__obf_c5c91f357c087ad4 *__obf_82cd85fd15cb8e6a) AppendTo(__obf_6068ae77f3b0e747 []byte) {
	binary.BigEndian.PutUint16(__obf_6068ae77f3b0e747[:2], __obf_c5c91f357c087ad4.Hi)
	binary.BigEndian.PutUint64(__obf_6068ae77f3b0e747[2:], __obf_c5c91f357c087ad4.Lo)
}

func (__obf_c5c91f357c087ad4 *__obf_82cd85fd15cb8e6a) Add(__obf_037bbba5aacdc33e uint64) (__obf_d2c82fd1ba6a2096 bool) {
	__obf_5398f72d7431fc5b, __obf_a24dd7a0f3b8ca96 := __obf_c5c91f357c087ad4.Lo, __obf_c5c91f357c087ad4.Hi
	if __obf_c5c91f357c087ad4.Lo += __obf_037bbba5aacdc33e; __obf_c5c91f357c087ad4.Lo < __obf_5398f72d7431fc5b {
		__obf_c5c91f357c087ad4.Hi++
	}
	return __obf_c5c91f357c087ad4.Hi < __obf_a24dd7a0f3b8ca96
}

func (__obf_c5c91f357c087ad4 __obf_82cd85fd15cb8e6a) IsZero() bool {
	return __obf_c5c91f357c087ad4.Hi == 0 && __obf_c5c91f357c087ad4.Lo == 0
}
