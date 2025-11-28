package __obf_7d8ac56be2e11a40

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
	__obf_36f601912d1c90f0 sync.Mutex
	__obf_c77498d7b2bae5a1 int64 = -1 ^ (-1 << NodeBits)
	__obf_ec285ad98985cba7       = __obf_c77498d7b2bae5a1 << StepBits
	__obf_be10def969d4f133 int64 = -1 ^ (-1 << StepBits)
	__obf_8ffc9beda4e1c033       = NodeBits + StepBits
	__obf_8c1966e59b35632a       = StepBits
)

const __obf_5d45c909699c3dba = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_86c5cb0a5983378e [256]byte

const __obf_5a6613ccc08f0aaa = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_a1a7f8528082f543 [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_7bba4a5c21dfeb3f []byte }

func (__obf_b3f4d034b8411de1 JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_b3f4d034b8411de1.__obf_7bba4a5c21dfeb3f))
}

// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func __obf_ee09e752d0fe765e() {

	for __obf_5a3113a41c0f69a6 := 0; __obf_5a3113a41c0f69a6 < len(__obf_a1a7f8528082f543); __obf_5a3113a41c0f69a6++ {
		__obf_a1a7f8528082f543[__obf_5a3113a41c0f69a6] = 0xFF
	}

	for __obf_5a3113a41c0f69a6 := 0; __obf_5a3113a41c0f69a6 < len(__obf_5a6613ccc08f0aaa); __obf_5a3113a41c0f69a6++ {
		__obf_a1a7f8528082f543[__obf_5a6613ccc08f0aaa[__obf_5a3113a41c0f69a6]] = byte(__obf_5a3113a41c0f69a6)
	}

	for __obf_5a3113a41c0f69a6 := 0; __obf_5a3113a41c0f69a6 < len(__obf_86c5cb0a5983378e); __obf_5a3113a41c0f69a6++ {
		__obf_86c5cb0a5983378e[__obf_5a3113a41c0f69a6] = 0xFF
	}

	for __obf_5a3113a41c0f69a6 := 0; __obf_5a3113a41c0f69a6 < len(__obf_5d45c909699c3dba); __obf_5a3113a41c0f69a6++ {
		__obf_86c5cb0a5983378e[__obf_5d45c909699c3dba[__obf_5a3113a41c0f69a6]] = byte(__obf_5a3113a41c0f69a6)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_36f601912d1c90f0 sync.Mutex
	__obf_ef20ff41b68cd1e9 time.Time
	time                   int64
	__obf_fdba374c730f5b05 int64
	__obf_12554eca78ec7094 int64

	__obf_c77498d7b2bae5a1 int64
	__obf_ec285ad98985cba7 int64
	__obf_be10def969d4f133 int64
	__obf_8ffc9beda4e1c033 uint8
	__obf_8c1966e59b35632a uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_fdba374c730f5b05 int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	// re-calc in case custom NodeBits or StepBits were set
	// DEPRECATED: the below block will be removed in a future release.
	__obf_36f601912d1c90f0.Lock()
	__obf_c77498d7b2bae5a1 = -1 ^ (-1 << NodeBits)
	__obf_ec285ad98985cba7 = __obf_c77498d7b2bae5a1 << StepBits
	__obf_be10def969d4f133 = -1 ^ (-1 << StepBits)
	__obf_8ffc9beda4e1c033 = NodeBits + StepBits
	__obf_8c1966e59b35632a = StepBits
	__obf_36f601912d1c90f0.Unlock()

	__obf_eacf881057f264e4 := Node{}
	__obf_eacf881057f264e4.__obf_fdba374c730f5b05 = __obf_fdba374c730f5b05
	__obf_eacf881057f264e4.__obf_c77498d7b2bae5a1 = -1 ^ (-1 << NodeBits)
	__obf_eacf881057f264e4.__obf_ec285ad98985cba7 = __obf_eacf881057f264e4.__obf_c77498d7b2bae5a1 << StepBits
	__obf_eacf881057f264e4.__obf_be10def969d4f133 = -1 ^ (-1 << StepBits)
	__obf_eacf881057f264e4.__obf_8ffc9beda4e1c033 = NodeBits + StepBits
	__obf_eacf881057f264e4.__obf_8c1966e59b35632a = StepBits

	if __obf_eacf881057f264e4.__obf_fdba374c730f5b05 < 0 || __obf_eacf881057f264e4.__obf_fdba374c730f5b05 > __obf_eacf881057f264e4.__obf_c77498d7b2bae5a1 {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_eacf881057f264e4.__obf_c77498d7b2bae5a1, 10))
	}

	var __obf_3b9a1892c4040328 = time.Now()
	// add time.Duration to curTime to make sure we use the monotonic clock if available
	__obf_eacf881057f264e4.__obf_ef20ff41b68cd1e9 = __obf_3b9a1892c4040328.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_3b9a1892c4040328))

	return &__obf_eacf881057f264e4, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_eacf881057f264e4 *Node) Generate() ID {

	__obf_eacf881057f264e4.__obf_36f601912d1c90f0.Lock()
	defer __obf_eacf881057f264e4.__obf_36f601912d1c90f0.Unlock()

	__obf_3ff127c31e0cee83 := time.Since(__obf_eacf881057f264e4.__obf_ef20ff41b68cd1e9).Milliseconds()

	if __obf_3ff127c31e0cee83 == __obf_eacf881057f264e4.time {
		__obf_eacf881057f264e4.__obf_12554eca78ec7094 = (__obf_eacf881057f264e4.__obf_12554eca78ec7094 + 1) & __obf_eacf881057f264e4.__obf_be10def969d4f133

		if __obf_eacf881057f264e4.__obf_12554eca78ec7094 == 0 {
			for __obf_3ff127c31e0cee83 <= __obf_eacf881057f264e4.time {
				__obf_3ff127c31e0cee83 = time.Since(__obf_eacf881057f264e4.__obf_ef20ff41b68cd1e9).Milliseconds()
			}
		}
	} else {
		__obf_eacf881057f264e4.__obf_12554eca78ec7094 = 0
	}

	__obf_eacf881057f264e4.time = __obf_3ff127c31e0cee83

	__obf_009c7aa79bd322a7 := ID((__obf_3ff127c31e0cee83)<<__obf_eacf881057f264e4.__obf_8ffc9beda4e1c033 |
		(__obf_eacf881057f264e4.__obf_fdba374c730f5b05 << __obf_eacf881057f264e4.__obf_8c1966e59b35632a) |
		(__obf_eacf881057f264e4.__obf_12554eca78ec7094),
	)

	return __obf_009c7aa79bd322a7
}

// Int64 returns an int64 of the snowflake ID
func (__obf_de4dbc1e5d8476d4 ID) Int64() int64 {
	return int64(__obf_de4dbc1e5d8476d4)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_6cfbdbb264317300 int64) ID {
	return ID(__obf_6cfbdbb264317300)
}

// String returns a string of the snowflake ID
func (__obf_de4dbc1e5d8476d4 ID) String() string {
	return strconv.FormatInt(int64(__obf_de4dbc1e5d8476d4), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_6cfbdbb264317300 string) (ID, error) {
	__obf_5a3113a41c0f69a6, __obf_3f3f828d9e46d714 := strconv.ParseInt(__obf_6cfbdbb264317300, 10, 64)
	return ID(__obf_5a3113a41c0f69a6), __obf_3f3f828d9e46d714

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_de4dbc1e5d8476d4 ID) Base2() string {
	return strconv.FormatInt(int64(__obf_de4dbc1e5d8476d4), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_6cfbdbb264317300 string) (ID, error) {
	__obf_5a3113a41c0f69a6, __obf_3f3f828d9e46d714 := strconv.ParseInt(__obf_6cfbdbb264317300, 2, 64)
	return ID(__obf_5a3113a41c0f69a6), __obf_3f3f828d9e46d714
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_de4dbc1e5d8476d4 ID) Base32() string {

	if __obf_de4dbc1e5d8476d4 < 32 {
		return string(__obf_5d45c909699c3dba[__obf_de4dbc1e5d8476d4])
	}

	__obf_a10739fc051883a8 := make([]byte, 0, 12)
	for __obf_de4dbc1e5d8476d4 >= 32 {
		__obf_a10739fc051883a8 = append(__obf_a10739fc051883a8, __obf_5d45c909699c3dba[__obf_de4dbc1e5d8476d4%32])
		__obf_de4dbc1e5d8476d4 /= 32
	}
	__obf_a10739fc051883a8 = append(__obf_a10739fc051883a8, __obf_5d45c909699c3dba[__obf_de4dbc1e5d8476d4])

	for __obf_e8e0549a04264062, __obf_60f039670bd4075e := 0, len(__obf_a10739fc051883a8)-1; __obf_e8e0549a04264062 < __obf_60f039670bd4075e; __obf_e8e0549a04264062, __obf_60f039670bd4075e = __obf_e8e0549a04264062+1, __obf_60f039670bd4075e-1 {
		__obf_a10739fc051883a8[__obf_e8e0549a04264062], __obf_a10739fc051883a8[__obf_60f039670bd4075e] = __obf_a10739fc051883a8[__obf_60f039670bd4075e], __obf_a10739fc051883a8[__obf_e8e0549a04264062]
	}

	return string(__obf_a10739fc051883a8)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_a10739fc051883a8 []byte) (ID, error) {

	var __obf_6cfbdbb264317300 int64

	for __obf_5a3113a41c0f69a6 := range __obf_a10739fc051883a8 {
		if __obf_86c5cb0a5983378e[__obf_a10739fc051883a8[__obf_5a3113a41c0f69a6]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_6cfbdbb264317300 = __obf_6cfbdbb264317300*32 + int64(__obf_86c5cb0a5983378e[__obf_a10739fc051883a8[__obf_5a3113a41c0f69a6]])
	}

	return ID(__obf_6cfbdbb264317300), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_de4dbc1e5d8476d4 ID) Base36() string {
	return strconv.FormatInt(int64(__obf_de4dbc1e5d8476d4), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_6cfbdbb264317300 string) (ID, error) {
	__obf_5a3113a41c0f69a6, __obf_3f3f828d9e46d714 := strconv.ParseInt(__obf_6cfbdbb264317300, 36, 64)
	return ID(__obf_5a3113a41c0f69a6), __obf_3f3f828d9e46d714
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_de4dbc1e5d8476d4 ID) Base58() string {

	if __obf_de4dbc1e5d8476d4 < 58 {
		return string(__obf_5a6613ccc08f0aaa[__obf_de4dbc1e5d8476d4])
	}

	__obf_a10739fc051883a8 := make([]byte, 0, 11)
	for __obf_de4dbc1e5d8476d4 >= 58 {
		__obf_a10739fc051883a8 = append(__obf_a10739fc051883a8, __obf_5a6613ccc08f0aaa[__obf_de4dbc1e5d8476d4%58])
		__obf_de4dbc1e5d8476d4 /= 58
	}
	__obf_a10739fc051883a8 = append(__obf_a10739fc051883a8, __obf_5a6613ccc08f0aaa[__obf_de4dbc1e5d8476d4])

	for __obf_e8e0549a04264062, __obf_60f039670bd4075e := 0, len(__obf_a10739fc051883a8)-1; __obf_e8e0549a04264062 < __obf_60f039670bd4075e; __obf_e8e0549a04264062, __obf_60f039670bd4075e = __obf_e8e0549a04264062+1, __obf_60f039670bd4075e-1 {
		__obf_a10739fc051883a8[__obf_e8e0549a04264062], __obf_a10739fc051883a8[__obf_60f039670bd4075e] = __obf_a10739fc051883a8[__obf_60f039670bd4075e], __obf_a10739fc051883a8[__obf_e8e0549a04264062]
	}

	return string(__obf_a10739fc051883a8)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_a10739fc051883a8 []byte) (ID, error) {

	var __obf_6cfbdbb264317300 int64

	for __obf_5a3113a41c0f69a6 := range __obf_a10739fc051883a8 {
		if __obf_a1a7f8528082f543[__obf_a10739fc051883a8[__obf_5a3113a41c0f69a6]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_6cfbdbb264317300 = __obf_6cfbdbb264317300*58 + int64(__obf_a1a7f8528082f543[__obf_a10739fc051883a8[__obf_5a3113a41c0f69a6]])
	}

	return ID(__obf_6cfbdbb264317300), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_de4dbc1e5d8476d4 ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_de4dbc1e5d8476d4.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_6cfbdbb264317300 string) (ID, error) {
	__obf_a10739fc051883a8, __obf_3f3f828d9e46d714 := base64.StdEncoding.DecodeString(__obf_6cfbdbb264317300)
	if __obf_3f3f828d9e46d714 != nil {
		return -1, __obf_3f3f828d9e46d714
	}
	return ParseBytes(__obf_a10739fc051883a8)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_de4dbc1e5d8476d4 ID) Bytes() []byte {
	return []byte(__obf_de4dbc1e5d8476d4.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_6cfbdbb264317300 []byte) (ID, error) {
	__obf_5a3113a41c0f69a6, __obf_3f3f828d9e46d714 := strconv.ParseInt(string(__obf_6cfbdbb264317300), 10, 64)
	return ID(__obf_5a3113a41c0f69a6), __obf_3f3f828d9e46d714
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_de4dbc1e5d8476d4 ID) IntBytes() [8]byte {
	var __obf_a10739fc051883a8 [8]byte
	binary.BigEndian.PutUint64(__obf_a10739fc051883a8[:], uint64(__obf_de4dbc1e5d8476d4))
	return __obf_a10739fc051883a8
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_6cfbdbb264317300 [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_6cfbdbb264317300[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_de4dbc1e5d8476d4 ID) Time() int64 {
	return (int64(__obf_de4dbc1e5d8476d4) >> __obf_8ffc9beda4e1c033) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_de4dbc1e5d8476d4 ID) Node() int64 {
	return int64(__obf_de4dbc1e5d8476d4) & __obf_ec285ad98985cba7 >> __obf_8c1966e59b35632a
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_de4dbc1e5d8476d4 ID) Step() int64 {
	return int64(__obf_de4dbc1e5d8476d4) & __obf_be10def969d4f133
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_de4dbc1e5d8476d4 ID) MarshalJSON() ([]byte, error) {
	__obf_b6c8338d997815d8 := make([]byte, 0, 22)
	__obf_b6c8338d997815d8 = append(__obf_b6c8338d997815d8, '"')
	__obf_b6c8338d997815d8 = strconv.AppendInt(__obf_b6c8338d997815d8, int64(__obf_de4dbc1e5d8476d4), 10)
	__obf_b6c8338d997815d8 = append(__obf_b6c8338d997815d8, '"')
	return __obf_b6c8338d997815d8, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_de4dbc1e5d8476d4 *ID) UnmarshalJSON(__obf_a10739fc051883a8 []byte) error {
	if len(__obf_a10739fc051883a8) < 3 || __obf_a10739fc051883a8[0] != '"' || __obf_a10739fc051883a8[len(__obf_a10739fc051883a8)-1] != '"' {
		return JSONSyntaxError{__obf_a10739fc051883a8}
	}

	__obf_5a3113a41c0f69a6, __obf_3f3f828d9e46d714 := strconv.ParseInt(string(__obf_a10739fc051883a8[1:len(__obf_a10739fc051883a8)-1]), 10, 64)
	if __obf_3f3f828d9e46d714 != nil {
		return __obf_3f3f828d9e46d714
	}

	*__obf_de4dbc1e5d8476d4 = ID(__obf_5a3113a41c0f69a6)
	return nil
}
