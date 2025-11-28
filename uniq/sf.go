package __obf_0f0e0d1ff72f3ff0

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
	__obf_8acdbb3d62aaf7ee sync.Mutex
	__obf_e550618e80e0b446 int64 = -1 ^ (-1 << NodeBits)
	__obf_0b5dc0d1b426e339       = __obf_e550618e80e0b446 << StepBits
	__obf_5f21123a81ec210f int64 = -1 ^ (-1 << StepBits)
	__obf_78b1b90cabc4cd8d       = NodeBits + StepBits
	__obf_39d60697045b4271       = StepBits
)

const __obf_42a8138f3694fef9 = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_9792f2a5eecd7068 [256]byte

const __obf_df7bfa4e2a125bdd = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_22660eac14c05ddd [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_bbdf738a9d50abc8 []byte }

func (__obf_a92847568986dcc0 JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_a92847568986dcc0.__obf_bbdf738a9d50abc8))
}

// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func __obf_98ff950639a7e2a9() {

	for __obf_0725dd8f5f5e1058 := 0; __obf_0725dd8f5f5e1058 < len(__obf_22660eac14c05ddd); __obf_0725dd8f5f5e1058++ {
		__obf_22660eac14c05ddd[__obf_0725dd8f5f5e1058] = 0xFF
	}

	for __obf_0725dd8f5f5e1058 := 0; __obf_0725dd8f5f5e1058 < len(__obf_df7bfa4e2a125bdd); __obf_0725dd8f5f5e1058++ {
		__obf_22660eac14c05ddd[__obf_df7bfa4e2a125bdd[__obf_0725dd8f5f5e1058]] = byte(__obf_0725dd8f5f5e1058)
	}

	for __obf_0725dd8f5f5e1058 := 0; __obf_0725dd8f5f5e1058 < len(__obf_9792f2a5eecd7068); __obf_0725dd8f5f5e1058++ {
		__obf_9792f2a5eecd7068[__obf_0725dd8f5f5e1058] = 0xFF
	}

	for __obf_0725dd8f5f5e1058 := 0; __obf_0725dd8f5f5e1058 < len(__obf_42a8138f3694fef9); __obf_0725dd8f5f5e1058++ {
		__obf_9792f2a5eecd7068[__obf_42a8138f3694fef9[__obf_0725dd8f5f5e1058]] = byte(__obf_0725dd8f5f5e1058)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_8acdbb3d62aaf7ee sync.Mutex
	__obf_2bd87f8d78b2aa83 time.Time
	time                   int64
	__obf_6b93e97530f1d07c int64
	__obf_485b2a90b2b0863a int64

	__obf_e550618e80e0b446 int64
	__obf_0b5dc0d1b426e339 int64
	__obf_5f21123a81ec210f int64
	__obf_78b1b90cabc4cd8d uint8
	__obf_39d60697045b4271 uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_6b93e97530f1d07c int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	// re-calc in case custom NodeBits or StepBits were set
	// DEPRECATED: the below block will be removed in a future release.
	__obf_8acdbb3d62aaf7ee.Lock()
	__obf_e550618e80e0b446 = -1 ^ (-1 << NodeBits)
	__obf_0b5dc0d1b426e339 = __obf_e550618e80e0b446 << StepBits
	__obf_5f21123a81ec210f = -1 ^ (-1 << StepBits)
	__obf_78b1b90cabc4cd8d = NodeBits + StepBits
	__obf_39d60697045b4271 = StepBits
	__obf_8acdbb3d62aaf7ee.Unlock()

	__obf_037bbba5aacdc33e := Node{}
	__obf_037bbba5aacdc33e.__obf_6b93e97530f1d07c = __obf_6b93e97530f1d07c
	__obf_037bbba5aacdc33e.__obf_e550618e80e0b446 = -1 ^ (-1 << NodeBits)
	__obf_037bbba5aacdc33e.__obf_0b5dc0d1b426e339 = __obf_037bbba5aacdc33e.__obf_e550618e80e0b446 << StepBits
	__obf_037bbba5aacdc33e.__obf_5f21123a81ec210f = -1 ^ (-1 << StepBits)
	__obf_037bbba5aacdc33e.__obf_78b1b90cabc4cd8d = NodeBits + StepBits
	__obf_037bbba5aacdc33e.__obf_39d60697045b4271 = StepBits

	if __obf_037bbba5aacdc33e.__obf_6b93e97530f1d07c < 0 || __obf_037bbba5aacdc33e.__obf_6b93e97530f1d07c > __obf_037bbba5aacdc33e.__obf_e550618e80e0b446 {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_037bbba5aacdc33e.__obf_e550618e80e0b446, 10))
	}

	var __obf_dd4c68b2fa2b52e3 = time.Now()
	// add time.Duration to curTime to make sure we use the monotonic clock if available
	__obf_037bbba5aacdc33e.__obf_2bd87f8d78b2aa83 = __obf_dd4c68b2fa2b52e3.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_dd4c68b2fa2b52e3))

	return &__obf_037bbba5aacdc33e, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_037bbba5aacdc33e *Node) Generate() ID {

	__obf_037bbba5aacdc33e.__obf_8acdbb3d62aaf7ee.Lock()
	defer __obf_037bbba5aacdc33e.__obf_8acdbb3d62aaf7ee.Unlock()

	__obf_2d87415af710f529 := time.Since(__obf_037bbba5aacdc33e.__obf_2bd87f8d78b2aa83).Milliseconds()

	if __obf_2d87415af710f529 == __obf_037bbba5aacdc33e.time {
		__obf_037bbba5aacdc33e.__obf_485b2a90b2b0863a = (__obf_037bbba5aacdc33e.__obf_485b2a90b2b0863a + 1) & __obf_037bbba5aacdc33e.__obf_5f21123a81ec210f

		if __obf_037bbba5aacdc33e.__obf_485b2a90b2b0863a == 0 {
			for __obf_2d87415af710f529 <= __obf_037bbba5aacdc33e.time {
				__obf_2d87415af710f529 = time.Since(__obf_037bbba5aacdc33e.__obf_2bd87f8d78b2aa83).Milliseconds()
			}
		}
	} else {
		__obf_037bbba5aacdc33e.__obf_485b2a90b2b0863a = 0
	}

	__obf_037bbba5aacdc33e.time = __obf_2d87415af710f529

	__obf_3aca7dca83f1e315 := ID((__obf_2d87415af710f529)<<__obf_037bbba5aacdc33e.__obf_78b1b90cabc4cd8d |
		(__obf_037bbba5aacdc33e.__obf_6b93e97530f1d07c << __obf_037bbba5aacdc33e.__obf_39d60697045b4271) |
		(__obf_037bbba5aacdc33e.__obf_485b2a90b2b0863a),
	)

	return __obf_3aca7dca83f1e315
}

// Int64 returns an int64 of the snowflake ID
func (__obf_2aa6953aa53f5f45 ID) Int64() int64 {
	return int64(__obf_2aa6953aa53f5f45)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_a7593a31c69c5d40 int64) ID {
	return ID(__obf_a7593a31c69c5d40)
}

// String returns a string of the snowflake ID
func (__obf_2aa6953aa53f5f45 ID) String() string {
	return strconv.FormatInt(int64(__obf_2aa6953aa53f5f45), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_a7593a31c69c5d40 string) (ID, error) {
	__obf_0725dd8f5f5e1058, __obf_f24ec7ae16791add := strconv.ParseInt(__obf_a7593a31c69c5d40, 10, 64)
	return ID(__obf_0725dd8f5f5e1058), __obf_f24ec7ae16791add

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_2aa6953aa53f5f45 ID) Base2() string {
	return strconv.FormatInt(int64(__obf_2aa6953aa53f5f45), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_a7593a31c69c5d40 string) (ID, error) {
	__obf_0725dd8f5f5e1058, __obf_f24ec7ae16791add := strconv.ParseInt(__obf_a7593a31c69c5d40, 2, 64)
	return ID(__obf_0725dd8f5f5e1058), __obf_f24ec7ae16791add
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_2aa6953aa53f5f45 ID) Base32() string {

	if __obf_2aa6953aa53f5f45 < 32 {
		return string(__obf_42a8138f3694fef9[__obf_2aa6953aa53f5f45])
	}

	__obf_880396ee93778b08 := make([]byte, 0, 12)
	for __obf_2aa6953aa53f5f45 >= 32 {
		__obf_880396ee93778b08 = append(__obf_880396ee93778b08, __obf_42a8138f3694fef9[__obf_2aa6953aa53f5f45%32])
		__obf_2aa6953aa53f5f45 /= 32
	}
	__obf_880396ee93778b08 = append(__obf_880396ee93778b08, __obf_42a8138f3694fef9[__obf_2aa6953aa53f5f45])

	for __obf_09cf9c38ec453181, __obf_c9762e566a14ef17 := 0, len(__obf_880396ee93778b08)-1; __obf_09cf9c38ec453181 < __obf_c9762e566a14ef17; __obf_09cf9c38ec453181, __obf_c9762e566a14ef17 = __obf_09cf9c38ec453181+1, __obf_c9762e566a14ef17-1 {
		__obf_880396ee93778b08[__obf_09cf9c38ec453181], __obf_880396ee93778b08[__obf_c9762e566a14ef17] = __obf_880396ee93778b08[__obf_c9762e566a14ef17], __obf_880396ee93778b08[__obf_09cf9c38ec453181]
	}

	return string(__obf_880396ee93778b08)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_880396ee93778b08 []byte) (ID, error) {

	var __obf_a7593a31c69c5d40 int64

	for __obf_0725dd8f5f5e1058 := range __obf_880396ee93778b08 {
		if __obf_9792f2a5eecd7068[__obf_880396ee93778b08[__obf_0725dd8f5f5e1058]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_a7593a31c69c5d40 = __obf_a7593a31c69c5d40*32 + int64(__obf_9792f2a5eecd7068[__obf_880396ee93778b08[__obf_0725dd8f5f5e1058]])
	}

	return ID(__obf_a7593a31c69c5d40), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_2aa6953aa53f5f45 ID) Base36() string {
	return strconv.FormatInt(int64(__obf_2aa6953aa53f5f45), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_a7593a31c69c5d40 string) (ID, error) {
	__obf_0725dd8f5f5e1058, __obf_f24ec7ae16791add := strconv.ParseInt(__obf_a7593a31c69c5d40, 36, 64)
	return ID(__obf_0725dd8f5f5e1058), __obf_f24ec7ae16791add
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_2aa6953aa53f5f45 ID) Base58() string {

	if __obf_2aa6953aa53f5f45 < 58 {
		return string(__obf_df7bfa4e2a125bdd[__obf_2aa6953aa53f5f45])
	}

	__obf_880396ee93778b08 := make([]byte, 0, 11)
	for __obf_2aa6953aa53f5f45 >= 58 {
		__obf_880396ee93778b08 = append(__obf_880396ee93778b08, __obf_df7bfa4e2a125bdd[__obf_2aa6953aa53f5f45%58])
		__obf_2aa6953aa53f5f45 /= 58
	}
	__obf_880396ee93778b08 = append(__obf_880396ee93778b08, __obf_df7bfa4e2a125bdd[__obf_2aa6953aa53f5f45])

	for __obf_09cf9c38ec453181, __obf_c9762e566a14ef17 := 0, len(__obf_880396ee93778b08)-1; __obf_09cf9c38ec453181 < __obf_c9762e566a14ef17; __obf_09cf9c38ec453181, __obf_c9762e566a14ef17 = __obf_09cf9c38ec453181+1, __obf_c9762e566a14ef17-1 {
		__obf_880396ee93778b08[__obf_09cf9c38ec453181], __obf_880396ee93778b08[__obf_c9762e566a14ef17] = __obf_880396ee93778b08[__obf_c9762e566a14ef17], __obf_880396ee93778b08[__obf_09cf9c38ec453181]
	}

	return string(__obf_880396ee93778b08)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_880396ee93778b08 []byte) (ID, error) {

	var __obf_a7593a31c69c5d40 int64

	for __obf_0725dd8f5f5e1058 := range __obf_880396ee93778b08 {
		if __obf_22660eac14c05ddd[__obf_880396ee93778b08[__obf_0725dd8f5f5e1058]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_a7593a31c69c5d40 = __obf_a7593a31c69c5d40*58 + int64(__obf_22660eac14c05ddd[__obf_880396ee93778b08[__obf_0725dd8f5f5e1058]])
	}

	return ID(__obf_a7593a31c69c5d40), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_2aa6953aa53f5f45 ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_2aa6953aa53f5f45.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_a7593a31c69c5d40 string) (ID, error) {
	__obf_880396ee93778b08, __obf_f24ec7ae16791add := base64.StdEncoding.DecodeString(__obf_a7593a31c69c5d40)
	if __obf_f24ec7ae16791add != nil {
		return -1, __obf_f24ec7ae16791add
	}
	return ParseBytes(__obf_880396ee93778b08)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_2aa6953aa53f5f45 ID) Bytes() []byte {
	return []byte(__obf_2aa6953aa53f5f45.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_a7593a31c69c5d40 []byte) (ID, error) {
	__obf_0725dd8f5f5e1058, __obf_f24ec7ae16791add := strconv.ParseInt(string(__obf_a7593a31c69c5d40), 10, 64)
	return ID(__obf_0725dd8f5f5e1058), __obf_f24ec7ae16791add
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_2aa6953aa53f5f45 ID) IntBytes() [8]byte {
	var __obf_880396ee93778b08 [8]byte
	binary.BigEndian.PutUint64(__obf_880396ee93778b08[:], uint64(__obf_2aa6953aa53f5f45))
	return __obf_880396ee93778b08
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_a7593a31c69c5d40 [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_a7593a31c69c5d40[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_2aa6953aa53f5f45 ID) Time() int64 {
	return (int64(__obf_2aa6953aa53f5f45) >> __obf_78b1b90cabc4cd8d) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_2aa6953aa53f5f45 ID) Node() int64 {
	return int64(__obf_2aa6953aa53f5f45) & __obf_0b5dc0d1b426e339 >> __obf_39d60697045b4271
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_2aa6953aa53f5f45 ID) Step() int64 {
	return int64(__obf_2aa6953aa53f5f45) & __obf_5f21123a81ec210f
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_2aa6953aa53f5f45 ID) MarshalJSON() ([]byte, error) {
	__obf_9ee4541142cff5fc := make([]byte, 0, 22)
	__obf_9ee4541142cff5fc = append(__obf_9ee4541142cff5fc, '"')
	__obf_9ee4541142cff5fc = strconv.AppendInt(__obf_9ee4541142cff5fc, int64(__obf_2aa6953aa53f5f45), 10)
	__obf_9ee4541142cff5fc = append(__obf_9ee4541142cff5fc, '"')
	return __obf_9ee4541142cff5fc, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_2aa6953aa53f5f45 *ID) UnmarshalJSON(__obf_880396ee93778b08 []byte) error {
	if len(__obf_880396ee93778b08) < 3 || __obf_880396ee93778b08[0] != '"' || __obf_880396ee93778b08[len(__obf_880396ee93778b08)-1] != '"' {
		return JSONSyntaxError{__obf_880396ee93778b08}
	}

	__obf_0725dd8f5f5e1058, __obf_f24ec7ae16791add := strconv.ParseInt(string(__obf_880396ee93778b08[1:len(__obf_880396ee93778b08)-1]), 10, 64)
	if __obf_f24ec7ae16791add != nil {
		return __obf_f24ec7ae16791add
	}

	*__obf_2aa6953aa53f5f45 = ID(__obf_0725dd8f5f5e1058)
	return nil
}
