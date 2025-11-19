package __obf_7913809aab6c8423

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
	__obf_cb0ab65220100005 sync.Mutex
	__obf_2926f7bcd2fe9a1a int64 = -1 ^ (-1 << NodeBits)
	__obf_b400a3f5e3bd647a       = __obf_2926f7bcd2fe9a1a << StepBits
	__obf_9013192fe9be0196 int64 = -1 ^ (-1 << StepBits)
	__obf_c5d4900a9c6cc2e4       = NodeBits + StepBits
	__obf_34d81b8109e4fbbe       = StepBits
)

const __obf_9a3a0ff0a8d85a41 = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_d0fe23fa361e5b45 [256]byte

const __obf_c57e20bde062975a = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_3de06d1876362eaa [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_ad385a4c5ce86824 []byte }

func (__obf_8114f600bfa0279c JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_8114f600bfa0279c.__obf_ad385a4c5ce86824))
}

// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func __obf_49ce724bf20a31b4() {

	for __obf_c2f487c3e92f8673 := 0; __obf_c2f487c3e92f8673 < len(__obf_3de06d1876362eaa); __obf_c2f487c3e92f8673++ {
		__obf_3de06d1876362eaa[__obf_c2f487c3e92f8673] = 0xFF
	}

	for __obf_c2f487c3e92f8673 := 0; __obf_c2f487c3e92f8673 < len(__obf_c57e20bde062975a); __obf_c2f487c3e92f8673++ {
		__obf_3de06d1876362eaa[__obf_c57e20bde062975a[__obf_c2f487c3e92f8673]] = byte(__obf_c2f487c3e92f8673)
	}

	for __obf_c2f487c3e92f8673 := 0; __obf_c2f487c3e92f8673 < len(__obf_d0fe23fa361e5b45); __obf_c2f487c3e92f8673++ {
		__obf_d0fe23fa361e5b45[__obf_c2f487c3e92f8673] = 0xFF
	}

	for __obf_c2f487c3e92f8673 := 0; __obf_c2f487c3e92f8673 < len(__obf_9a3a0ff0a8d85a41); __obf_c2f487c3e92f8673++ {
		__obf_d0fe23fa361e5b45[__obf_9a3a0ff0a8d85a41[__obf_c2f487c3e92f8673]] = byte(__obf_c2f487c3e92f8673)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_cb0ab65220100005 sync.Mutex
	__obf_acb5fc709f43643b time.Time
	time                   int64
	__obf_e9aab4dfc7a97570 int64
	__obf_b7de5ab06ea39213 int64

	__obf_2926f7bcd2fe9a1a int64
	__obf_b400a3f5e3bd647a int64
	__obf_9013192fe9be0196 int64
	__obf_c5d4900a9c6cc2e4 uint8
	__obf_34d81b8109e4fbbe uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_e9aab4dfc7a97570 int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	// re-calc in case custom NodeBits or StepBits were set
	// DEPRECATED: the below block will be removed in a future release.
	__obf_cb0ab65220100005.Lock()
	__obf_2926f7bcd2fe9a1a = -1 ^ (-1 << NodeBits)
	__obf_b400a3f5e3bd647a = __obf_2926f7bcd2fe9a1a << StepBits
	__obf_9013192fe9be0196 = -1 ^ (-1 << StepBits)
	__obf_c5d4900a9c6cc2e4 = NodeBits + StepBits
	__obf_34d81b8109e4fbbe = StepBits
	__obf_cb0ab65220100005.Unlock()

	__obf_61da857b954dcf36 := Node{}
	__obf_61da857b954dcf36.__obf_e9aab4dfc7a97570 = __obf_e9aab4dfc7a97570
	__obf_61da857b954dcf36.__obf_2926f7bcd2fe9a1a = -1 ^ (-1 << NodeBits)
	__obf_61da857b954dcf36.__obf_b400a3f5e3bd647a = __obf_61da857b954dcf36.__obf_2926f7bcd2fe9a1a << StepBits
	__obf_61da857b954dcf36.__obf_9013192fe9be0196 = -1 ^ (-1 << StepBits)
	__obf_61da857b954dcf36.__obf_c5d4900a9c6cc2e4 = NodeBits + StepBits
	__obf_61da857b954dcf36.__obf_34d81b8109e4fbbe = StepBits

	if __obf_61da857b954dcf36.__obf_e9aab4dfc7a97570 < 0 || __obf_61da857b954dcf36.__obf_e9aab4dfc7a97570 > __obf_61da857b954dcf36.__obf_2926f7bcd2fe9a1a {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_61da857b954dcf36.__obf_2926f7bcd2fe9a1a, 10))
	}

	var __obf_b3b1d82801237c28 = time.Now()
	// add time.Duration to curTime to make sure we use the monotonic clock if available
	__obf_61da857b954dcf36.__obf_acb5fc709f43643b = __obf_b3b1d82801237c28.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_b3b1d82801237c28))

	return &__obf_61da857b954dcf36, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_61da857b954dcf36 *Node) Generate() ID {

	__obf_61da857b954dcf36.__obf_cb0ab65220100005.Lock()
	defer __obf_61da857b954dcf36.__obf_cb0ab65220100005.Unlock()

	__obf_0cb19e0f723e3d0d := time.Since(__obf_61da857b954dcf36.__obf_acb5fc709f43643b).Milliseconds()

	if __obf_0cb19e0f723e3d0d == __obf_61da857b954dcf36.time {
		__obf_61da857b954dcf36.__obf_b7de5ab06ea39213 = (__obf_61da857b954dcf36.__obf_b7de5ab06ea39213 + 1) & __obf_61da857b954dcf36.__obf_9013192fe9be0196

		if __obf_61da857b954dcf36.__obf_b7de5ab06ea39213 == 0 {
			for __obf_0cb19e0f723e3d0d <= __obf_61da857b954dcf36.time {
				__obf_0cb19e0f723e3d0d = time.Since(__obf_61da857b954dcf36.__obf_acb5fc709f43643b).Milliseconds()
			}
		}
	} else {
		__obf_61da857b954dcf36.__obf_b7de5ab06ea39213 = 0
	}

	__obf_61da857b954dcf36.time = __obf_0cb19e0f723e3d0d

	__obf_8b640e335ac0bd44 := ID((__obf_0cb19e0f723e3d0d)<<__obf_61da857b954dcf36.__obf_c5d4900a9c6cc2e4 |
		(__obf_61da857b954dcf36.__obf_e9aab4dfc7a97570 << __obf_61da857b954dcf36.__obf_34d81b8109e4fbbe) |
		(__obf_61da857b954dcf36.__obf_b7de5ab06ea39213),
	)

	return __obf_8b640e335ac0bd44
}

// Int64 returns an int64 of the snowflake ID
func (__obf_7a4baec89f9c7e29 ID) Int64() int64 {
	return int64(__obf_7a4baec89f9c7e29)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_1dfe314328e20778 int64) ID {
	return ID(__obf_1dfe314328e20778)
}

// String returns a string of the snowflake ID
func (__obf_7a4baec89f9c7e29 ID) String() string {
	return strconv.FormatInt(int64(__obf_7a4baec89f9c7e29), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_1dfe314328e20778 string) (ID, error) {
	__obf_c2f487c3e92f8673, __obf_e62cd9e417a87ee7 := strconv.ParseInt(__obf_1dfe314328e20778, 10, 64)
	return ID(__obf_c2f487c3e92f8673), __obf_e62cd9e417a87ee7

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_7a4baec89f9c7e29 ID) Base2() string {
	return strconv.FormatInt(int64(__obf_7a4baec89f9c7e29), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_1dfe314328e20778 string) (ID, error) {
	__obf_c2f487c3e92f8673, __obf_e62cd9e417a87ee7 := strconv.ParseInt(__obf_1dfe314328e20778, 2, 64)
	return ID(__obf_c2f487c3e92f8673), __obf_e62cd9e417a87ee7
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_7a4baec89f9c7e29 ID) Base32() string {

	if __obf_7a4baec89f9c7e29 < 32 {
		return string(__obf_9a3a0ff0a8d85a41[__obf_7a4baec89f9c7e29])
	}

	__obf_a502f0db8c3fed48 := make([]byte, 0, 12)
	for __obf_7a4baec89f9c7e29 >= 32 {
		__obf_a502f0db8c3fed48 = append(__obf_a502f0db8c3fed48, __obf_9a3a0ff0a8d85a41[__obf_7a4baec89f9c7e29%32])
		__obf_7a4baec89f9c7e29 /= 32
	}
	__obf_a502f0db8c3fed48 = append(__obf_a502f0db8c3fed48, __obf_9a3a0ff0a8d85a41[__obf_7a4baec89f9c7e29])

	for __obf_ffa353604979097c, __obf_7c87814e4324c879 := 0, len(__obf_a502f0db8c3fed48)-1; __obf_ffa353604979097c < __obf_7c87814e4324c879; __obf_ffa353604979097c, __obf_7c87814e4324c879 = __obf_ffa353604979097c+1, __obf_7c87814e4324c879-1 {
		__obf_a502f0db8c3fed48[__obf_ffa353604979097c], __obf_a502f0db8c3fed48[__obf_7c87814e4324c879] = __obf_a502f0db8c3fed48[__obf_7c87814e4324c879], __obf_a502f0db8c3fed48[__obf_ffa353604979097c]
	}

	return string(__obf_a502f0db8c3fed48)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_a502f0db8c3fed48 []byte) (ID, error) {

	var __obf_1dfe314328e20778 int64

	for __obf_c2f487c3e92f8673 := range __obf_a502f0db8c3fed48 {
		if __obf_d0fe23fa361e5b45[__obf_a502f0db8c3fed48[__obf_c2f487c3e92f8673]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_1dfe314328e20778 = __obf_1dfe314328e20778*32 + int64(__obf_d0fe23fa361e5b45[__obf_a502f0db8c3fed48[__obf_c2f487c3e92f8673]])
	}

	return ID(__obf_1dfe314328e20778), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_7a4baec89f9c7e29 ID) Base36() string {
	return strconv.FormatInt(int64(__obf_7a4baec89f9c7e29), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_1dfe314328e20778 string) (ID, error) {
	__obf_c2f487c3e92f8673, __obf_e62cd9e417a87ee7 := strconv.ParseInt(__obf_1dfe314328e20778, 36, 64)
	return ID(__obf_c2f487c3e92f8673), __obf_e62cd9e417a87ee7
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_7a4baec89f9c7e29 ID) Base58() string {

	if __obf_7a4baec89f9c7e29 < 58 {
		return string(__obf_c57e20bde062975a[__obf_7a4baec89f9c7e29])
	}

	__obf_a502f0db8c3fed48 := make([]byte, 0, 11)
	for __obf_7a4baec89f9c7e29 >= 58 {
		__obf_a502f0db8c3fed48 = append(__obf_a502f0db8c3fed48, __obf_c57e20bde062975a[__obf_7a4baec89f9c7e29%58])
		__obf_7a4baec89f9c7e29 /= 58
	}
	__obf_a502f0db8c3fed48 = append(__obf_a502f0db8c3fed48, __obf_c57e20bde062975a[__obf_7a4baec89f9c7e29])

	for __obf_ffa353604979097c, __obf_7c87814e4324c879 := 0, len(__obf_a502f0db8c3fed48)-1; __obf_ffa353604979097c < __obf_7c87814e4324c879; __obf_ffa353604979097c, __obf_7c87814e4324c879 = __obf_ffa353604979097c+1, __obf_7c87814e4324c879-1 {
		__obf_a502f0db8c3fed48[__obf_ffa353604979097c], __obf_a502f0db8c3fed48[__obf_7c87814e4324c879] = __obf_a502f0db8c3fed48[__obf_7c87814e4324c879], __obf_a502f0db8c3fed48[__obf_ffa353604979097c]
	}

	return string(__obf_a502f0db8c3fed48)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_a502f0db8c3fed48 []byte) (ID, error) {

	var __obf_1dfe314328e20778 int64

	for __obf_c2f487c3e92f8673 := range __obf_a502f0db8c3fed48 {
		if __obf_3de06d1876362eaa[__obf_a502f0db8c3fed48[__obf_c2f487c3e92f8673]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_1dfe314328e20778 = __obf_1dfe314328e20778*58 + int64(__obf_3de06d1876362eaa[__obf_a502f0db8c3fed48[__obf_c2f487c3e92f8673]])
	}

	return ID(__obf_1dfe314328e20778), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_7a4baec89f9c7e29 ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_7a4baec89f9c7e29.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_1dfe314328e20778 string) (ID, error) {
	__obf_a502f0db8c3fed48, __obf_e62cd9e417a87ee7 := base64.StdEncoding.DecodeString(__obf_1dfe314328e20778)
	if __obf_e62cd9e417a87ee7 != nil {
		return -1, __obf_e62cd9e417a87ee7
	}
	return ParseBytes(__obf_a502f0db8c3fed48)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_7a4baec89f9c7e29 ID) Bytes() []byte {
	return []byte(__obf_7a4baec89f9c7e29.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_1dfe314328e20778 []byte) (ID, error) {
	__obf_c2f487c3e92f8673, __obf_e62cd9e417a87ee7 := strconv.ParseInt(string(__obf_1dfe314328e20778), 10, 64)
	return ID(__obf_c2f487c3e92f8673), __obf_e62cd9e417a87ee7
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_7a4baec89f9c7e29 ID) IntBytes() [8]byte {
	var __obf_a502f0db8c3fed48 [8]byte
	binary.BigEndian.PutUint64(__obf_a502f0db8c3fed48[:], uint64(__obf_7a4baec89f9c7e29))
	return __obf_a502f0db8c3fed48
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_1dfe314328e20778 [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_1dfe314328e20778[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_7a4baec89f9c7e29 ID) Time() int64 {
	return (int64(__obf_7a4baec89f9c7e29) >> __obf_c5d4900a9c6cc2e4) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_7a4baec89f9c7e29 ID) Node() int64 {
	return int64(__obf_7a4baec89f9c7e29) & __obf_b400a3f5e3bd647a >> __obf_34d81b8109e4fbbe
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_7a4baec89f9c7e29 ID) Step() int64 {
	return int64(__obf_7a4baec89f9c7e29) & __obf_9013192fe9be0196
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_7a4baec89f9c7e29 ID) MarshalJSON() ([]byte, error) {
	__obf_1e43a5fc987cfd28 := make([]byte, 0, 22)
	__obf_1e43a5fc987cfd28 = append(__obf_1e43a5fc987cfd28, '"')
	__obf_1e43a5fc987cfd28 = strconv.AppendInt(__obf_1e43a5fc987cfd28, int64(__obf_7a4baec89f9c7e29), 10)
	__obf_1e43a5fc987cfd28 = append(__obf_1e43a5fc987cfd28, '"')
	return __obf_1e43a5fc987cfd28, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_7a4baec89f9c7e29 *ID) UnmarshalJSON(__obf_a502f0db8c3fed48 []byte) error {
	if len(__obf_a502f0db8c3fed48) < 3 || __obf_a502f0db8c3fed48[0] != '"' || __obf_a502f0db8c3fed48[len(__obf_a502f0db8c3fed48)-1] != '"' {
		return JSONSyntaxError{__obf_a502f0db8c3fed48}
	}

	__obf_c2f487c3e92f8673, __obf_e62cd9e417a87ee7 := strconv.ParseInt(string(__obf_a502f0db8c3fed48[1:len(__obf_a502f0db8c3fed48)-1]), 10, 64)
	if __obf_e62cd9e417a87ee7 != nil {
		return __obf_e62cd9e417a87ee7
	}

	*__obf_7a4baec89f9c7e29 = ID(__obf_c2f487c3e92f8673)
	return nil
}
