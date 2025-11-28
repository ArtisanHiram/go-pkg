package __obf_e2239f4853c61591

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	// Epoch is set to the twitter snowflake epoch of Nov 04 2010 01:42:54 UTC in milliseconds
	// You may customize this to set a different epoch for your application.
	Epoch int64 = 1288834974657

	// NodeBits holds the number of bits to use for Node
	// Remember, you have a total 22 bits to share between Node/Step
	NodeBits uint8 = 10

	// StepBits holds the number of bits to use for Step
	// Remember, you have a total 22 bits to share between Node/Step
	StepBits uint8 = 12

	// DEPRECATED: the below four variables will be removed in a future release.
	__obf_19980aa09194e26b sync.Mutex
	__obf_d72082f7b564cda4 int64 = -1 ^ (-1 << NodeBits)
	__obf_3ffd8d4153eb1fda       = __obf_d72082f7b564cda4 << StepBits
	__obf_dd92903e5d032f20 int64 = -1 ^ (-1 << StepBits)
	__obf_c63a3a171b2fa23b       = NodeBits + StepBits
	__obf_fc5608b28ac2080a       = StepBits
)

const __obf_669ed185c1c8d1d4 = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_097a84d1abdb0d3a [256]byte

const __obf_7fbc23b6463890cc = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_b68b0eed645c6256 [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_7dc361783814a1a0 []byte }

func (__obf_2715144e561f317a JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_2715144e561f317a.__obf_7dc361783814a1a0))
}

// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func __obf_4f66ddc68d38d55f() {

	for __obf_38a57c7e430e9465 := 0; __obf_38a57c7e430e9465 < len(__obf_b68b0eed645c6256); __obf_38a57c7e430e9465++ {
		__obf_b68b0eed645c6256[__obf_38a57c7e430e9465] = 0xFF
	}

	for __obf_38a57c7e430e9465 := 0; __obf_38a57c7e430e9465 < len(__obf_7fbc23b6463890cc); __obf_38a57c7e430e9465++ {
		__obf_b68b0eed645c6256[__obf_7fbc23b6463890cc[__obf_38a57c7e430e9465]] = byte(__obf_38a57c7e430e9465)
	}

	for __obf_38a57c7e430e9465 := 0; __obf_38a57c7e430e9465 < len(__obf_097a84d1abdb0d3a); __obf_38a57c7e430e9465++ {
		__obf_097a84d1abdb0d3a[__obf_38a57c7e430e9465] = 0xFF
	}

	for __obf_38a57c7e430e9465 := 0; __obf_38a57c7e430e9465 < len(__obf_669ed185c1c8d1d4); __obf_38a57c7e430e9465++ {
		__obf_097a84d1abdb0d3a[__obf_669ed185c1c8d1d4[__obf_38a57c7e430e9465]] = byte(__obf_38a57c7e430e9465)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_19980aa09194e26b sync.Mutex
	__obf_f257f6963f8260bd time.Time
	time                   int64
	__obf_8e8cb6321097d26f int64
	__obf_dafeaf4f95060dcc int64

	__obf_d72082f7b564cda4 int64
	__obf_3ffd8d4153eb1fda int64
	__obf_dd92903e5d032f20 int64
	__obf_c63a3a171b2fa23b uint8
	__obf_fc5608b28ac2080a uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_8e8cb6321097d26f int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	// re-calc in case custom NodeBits or StepBits were set
	// DEPRECATED: the below block will be removed in a future release.
	__obf_19980aa09194e26b.Lock()
	__obf_d72082f7b564cda4 = -1 ^ (-1 << NodeBits)
	__obf_3ffd8d4153eb1fda = __obf_d72082f7b564cda4 << StepBits
	__obf_dd92903e5d032f20 = -1 ^ (-1 << StepBits)
	__obf_c63a3a171b2fa23b = NodeBits + StepBits
	__obf_fc5608b28ac2080a = StepBits
	__obf_19980aa09194e26b.Unlock()

	__obf_1758a8beead18d7e := Node{}
	__obf_1758a8beead18d7e.__obf_8e8cb6321097d26f = __obf_8e8cb6321097d26f
	__obf_1758a8beead18d7e.__obf_d72082f7b564cda4 = -1 ^ (-1 << NodeBits)
	__obf_1758a8beead18d7e.__obf_3ffd8d4153eb1fda = __obf_1758a8beead18d7e.__obf_d72082f7b564cda4 << StepBits
	__obf_1758a8beead18d7e.__obf_dd92903e5d032f20 = -1 ^ (-1 << StepBits)
	__obf_1758a8beead18d7e.__obf_c63a3a171b2fa23b = NodeBits + StepBits
	__obf_1758a8beead18d7e.__obf_fc5608b28ac2080a = StepBits

	if __obf_1758a8beead18d7e.__obf_8e8cb6321097d26f < 0 || __obf_1758a8beead18d7e.__obf_8e8cb6321097d26f > __obf_1758a8beead18d7e.__obf_d72082f7b564cda4 {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_1758a8beead18d7e.__obf_d72082f7b564cda4, 10))
	}

	var __obf_0f2705024ab0fbb4 = time.Now()
	// add time.Duration to curTime to make sure we use the monotonic clock if available
	__obf_1758a8beead18d7e.__obf_f257f6963f8260bd = __obf_0f2705024ab0fbb4.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_0f2705024ab0fbb4))

	return &__obf_1758a8beead18d7e, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_1758a8beead18d7e *Node) Generate() ID {

	__obf_1758a8beead18d7e.__obf_19980aa09194e26b.Lock()
	defer __obf_1758a8beead18d7e.__obf_19980aa09194e26b.Unlock()

	__obf_4e31f2cac37e1712 := time.Since(__obf_1758a8beead18d7e.__obf_f257f6963f8260bd).Milliseconds()

	if __obf_4e31f2cac37e1712 == __obf_1758a8beead18d7e.time {
		__obf_1758a8beead18d7e.__obf_dafeaf4f95060dcc = (__obf_1758a8beead18d7e.__obf_dafeaf4f95060dcc + 1) & __obf_1758a8beead18d7e.__obf_dd92903e5d032f20

		if __obf_1758a8beead18d7e.__obf_dafeaf4f95060dcc == 0 {
			for __obf_4e31f2cac37e1712 <= __obf_1758a8beead18d7e.time {
				__obf_4e31f2cac37e1712 = time.Since(__obf_1758a8beead18d7e.__obf_f257f6963f8260bd).Milliseconds()
			}
		}
	} else {
		__obf_1758a8beead18d7e.__obf_dafeaf4f95060dcc = 0
	}

	__obf_1758a8beead18d7e.time = __obf_4e31f2cac37e1712

	__obf_8f477f5e0816a289 := ID((__obf_4e31f2cac37e1712)<<__obf_1758a8beead18d7e.__obf_c63a3a171b2fa23b |
		(__obf_1758a8beead18d7e.__obf_8e8cb6321097d26f << __obf_1758a8beead18d7e.__obf_fc5608b28ac2080a) |
		(__obf_1758a8beead18d7e.__obf_dafeaf4f95060dcc),
	)

	return __obf_8f477f5e0816a289
}

// Int64 returns an int64 of the snowflake ID
func (__obf_217bce253f99f19f ID) Int64() int64 {
	return int64(__obf_217bce253f99f19f)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_cee82df7c007227d int64) ID {
	return ID(__obf_cee82df7c007227d)
}

// String returns a string of the snowflake ID
func (__obf_217bce253f99f19f ID) String() string {
	return strconv.FormatInt(int64(__obf_217bce253f99f19f), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_cee82df7c007227d string) (ID, error) {
	__obf_38a57c7e430e9465, __obf_bb9dc4c4f445b22f := strconv.ParseInt(__obf_cee82df7c007227d, 10, 64)
	return ID(__obf_38a57c7e430e9465), __obf_bb9dc4c4f445b22f

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_217bce253f99f19f ID) Base2() string {
	return strconv.FormatInt(int64(__obf_217bce253f99f19f), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_cee82df7c007227d string) (ID, error) {
	__obf_38a57c7e430e9465, __obf_bb9dc4c4f445b22f := strconv.ParseInt(__obf_cee82df7c007227d, 2, 64)
	return ID(__obf_38a57c7e430e9465), __obf_bb9dc4c4f445b22f
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_217bce253f99f19f ID) Base32() string {

	if __obf_217bce253f99f19f < 32 {
		return string(__obf_669ed185c1c8d1d4[__obf_217bce253f99f19f])
	}

	__obf_861323918c1ee08b := make([]byte, 0, 12)
	for __obf_217bce253f99f19f >= 32 {
		__obf_861323918c1ee08b = append(__obf_861323918c1ee08b, __obf_669ed185c1c8d1d4[__obf_217bce253f99f19f%32])
		__obf_217bce253f99f19f /= 32
	}
	__obf_861323918c1ee08b = append(__obf_861323918c1ee08b, __obf_669ed185c1c8d1d4[__obf_217bce253f99f19f])

	for __obf_87fbd978922103e3, __obf_00056d588e338871 := 0, len(__obf_861323918c1ee08b)-1; __obf_87fbd978922103e3 < __obf_00056d588e338871; __obf_87fbd978922103e3, __obf_00056d588e338871 = __obf_87fbd978922103e3+1, __obf_00056d588e338871-1 {
		__obf_861323918c1ee08b[__obf_87fbd978922103e3], __obf_861323918c1ee08b[__obf_00056d588e338871] = __obf_861323918c1ee08b[__obf_00056d588e338871], __obf_861323918c1ee08b[__obf_87fbd978922103e3]
	}

	return string(__obf_861323918c1ee08b)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_861323918c1ee08b []byte) (ID, error) {

	var __obf_cee82df7c007227d int64

	for __obf_38a57c7e430e9465 := range __obf_861323918c1ee08b {
		if __obf_097a84d1abdb0d3a[__obf_861323918c1ee08b[__obf_38a57c7e430e9465]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_cee82df7c007227d = __obf_cee82df7c007227d*32 + int64(__obf_097a84d1abdb0d3a[__obf_861323918c1ee08b[__obf_38a57c7e430e9465]])
	}

	return ID(__obf_cee82df7c007227d), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_217bce253f99f19f ID) Base36() string {
	return strconv.FormatInt(int64(__obf_217bce253f99f19f), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_cee82df7c007227d string) (ID, error) {
	__obf_38a57c7e430e9465, __obf_bb9dc4c4f445b22f := strconv.ParseInt(__obf_cee82df7c007227d, 36, 64)
	return ID(__obf_38a57c7e430e9465), __obf_bb9dc4c4f445b22f
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_217bce253f99f19f ID) Base58() string {

	if __obf_217bce253f99f19f < 58 {
		return string(__obf_7fbc23b6463890cc[__obf_217bce253f99f19f])
	}

	__obf_861323918c1ee08b := make([]byte, 0, 11)
	for __obf_217bce253f99f19f >= 58 {
		__obf_861323918c1ee08b = append(__obf_861323918c1ee08b, __obf_7fbc23b6463890cc[__obf_217bce253f99f19f%58])
		__obf_217bce253f99f19f /= 58
	}
	__obf_861323918c1ee08b = append(__obf_861323918c1ee08b, __obf_7fbc23b6463890cc[__obf_217bce253f99f19f])

	for __obf_87fbd978922103e3, __obf_00056d588e338871 := 0, len(__obf_861323918c1ee08b)-1; __obf_87fbd978922103e3 < __obf_00056d588e338871; __obf_87fbd978922103e3, __obf_00056d588e338871 = __obf_87fbd978922103e3+1, __obf_00056d588e338871-1 {
		__obf_861323918c1ee08b[__obf_87fbd978922103e3], __obf_861323918c1ee08b[__obf_00056d588e338871] = __obf_861323918c1ee08b[__obf_00056d588e338871], __obf_861323918c1ee08b[__obf_87fbd978922103e3]
	}

	return string(__obf_861323918c1ee08b)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_861323918c1ee08b []byte) (ID, error) {

	var __obf_cee82df7c007227d int64

	for __obf_38a57c7e430e9465 := range __obf_861323918c1ee08b {
		if __obf_b68b0eed645c6256[__obf_861323918c1ee08b[__obf_38a57c7e430e9465]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_cee82df7c007227d = __obf_cee82df7c007227d*58 + int64(__obf_b68b0eed645c6256[__obf_861323918c1ee08b[__obf_38a57c7e430e9465]])
	}

	return ID(__obf_cee82df7c007227d), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_217bce253f99f19f ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_217bce253f99f19f.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_cee82df7c007227d string) (ID, error) {
	__obf_861323918c1ee08b, __obf_bb9dc4c4f445b22f := base64.StdEncoding.DecodeString(__obf_cee82df7c007227d)
	if __obf_bb9dc4c4f445b22f != nil {
		return -1, __obf_bb9dc4c4f445b22f
	}
	return ParseBytes(__obf_861323918c1ee08b)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_217bce253f99f19f ID) Bytes() []byte {
	return []byte(__obf_217bce253f99f19f.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_cee82df7c007227d []byte) (ID, error) {
	__obf_38a57c7e430e9465, __obf_bb9dc4c4f445b22f := strconv.ParseInt(string(__obf_cee82df7c007227d), 10, 64)
	return ID(__obf_38a57c7e430e9465), __obf_bb9dc4c4f445b22f
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_217bce253f99f19f ID) IntBytes() [8]byte {
	var __obf_861323918c1ee08b [8]byte
	binary.BigEndian.PutUint64(__obf_861323918c1ee08b[:], uint64(__obf_217bce253f99f19f))
	return __obf_861323918c1ee08b
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_cee82df7c007227d [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_cee82df7c007227d[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_217bce253f99f19f ID) Time() int64 {
	return (int64(__obf_217bce253f99f19f) >> __obf_c63a3a171b2fa23b) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_217bce253f99f19f ID) Node() int64 {
	return int64(__obf_217bce253f99f19f) & __obf_3ffd8d4153eb1fda >> __obf_fc5608b28ac2080a
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_217bce253f99f19f ID) Step() int64 {
	return int64(__obf_217bce253f99f19f) & __obf_dd92903e5d032f20
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_217bce253f99f19f ID) MarshalJSON() ([]byte, error) {
	__obf_cc9f249f75df880f := make([]byte, 0, 22)
	__obf_cc9f249f75df880f = append(__obf_cc9f249f75df880f, '"')
	__obf_cc9f249f75df880f = strconv.AppendInt(__obf_cc9f249f75df880f, int64(__obf_217bce253f99f19f), 10)
	__obf_cc9f249f75df880f = append(__obf_cc9f249f75df880f, '"')
	return __obf_cc9f249f75df880f, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_217bce253f99f19f *ID) UnmarshalJSON(__obf_861323918c1ee08b []byte) error {
	if len(__obf_861323918c1ee08b) < 3 || __obf_861323918c1ee08b[0] != '"' || __obf_861323918c1ee08b[len(__obf_861323918c1ee08b)-1] != '"' {
		return JSONSyntaxError{__obf_861323918c1ee08b}
	}

	__obf_38a57c7e430e9465, __obf_bb9dc4c4f445b22f := strconv.ParseInt(string(__obf_861323918c1ee08b[1:len(__obf_861323918c1ee08b)-1]), 10, 64)
	if __obf_bb9dc4c4f445b22f != nil {
		return __obf_bb9dc4c4f445b22f
	}

	*__obf_217bce253f99f19f = ID(__obf_38a57c7e430e9465)
	return nil
}
