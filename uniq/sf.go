package __obf_07f0876faa0cf68e

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
	StepBits               uint8 = 12
	__obf_15719f79f826bdc4 sync. // DEPRECATED: the below four variables will be removed in a future release.
				Mutex
	__obf_fa14daecff57d827 int64 = -1 ^ (-1 << NodeBits)
	__obf_8a3309aeba4c3e95       = __obf_fa14daecff57d827 <<
		StepBits
	__obf_c23664664a7762d5 int64 = -1 ^ (-1 << StepBits)
	__obf_462febdd1e69694f       = NodeBits + StepBits
	__obf_6df89a1a9eab32eb       = StepBits
)

const __obf_3de1fb88ae0b4aee = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_99dec5fdbabbf1cc [256]byte

const __obf_e969c0efc0d267e0 = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_5bbf872fd814e2b3 [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_80a5ab3fcf95cc43 []byte }

func (__obf_d9a2c13a100d8195 JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_d9a2c13a100d8195.

		// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
		__obf_80a5ab3fcf95cc43))
}

var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func init() {

	for __obf_1ef2f5d835c97294 := 0; __obf_1ef2f5d835c97294 < len(__obf_5bbf872fd814e2b3); __obf_1ef2f5d835c97294++ {
		__obf_5bbf872fd814e2b3[__obf_1ef2f5d835c97294] = 0xFF
	}

	for __obf_1ef2f5d835c97294 := 0; __obf_1ef2f5d835c97294 < len(__obf_e969c0efc0d267e0); __obf_1ef2f5d835c97294++ {
		__obf_5bbf872fd814e2b3[__obf_e969c0efc0d267e0[__obf_1ef2f5d835c97294]] = byte(__obf_1ef2f5d835c97294)
	}

	for __obf_1ef2f5d835c97294 := 0; __obf_1ef2f5d835c97294 < len(__obf_99dec5fdbabbf1cc); __obf_1ef2f5d835c97294++ {
		__obf_99dec5fdbabbf1cc[__obf_1ef2f5d835c97294] = 0xFF
	}

	for __obf_1ef2f5d835c97294 := 0; __obf_1ef2f5d835c97294 < len(__obf_3de1fb88ae0b4aee); __obf_1ef2f5d835c97294++ {
		__obf_99dec5fdbabbf1cc[__obf_3de1fb88ae0b4aee[__obf_1ef2f5d835c97294]] = byte(__obf_1ef2f5d835c97294)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_15719f79f826bdc4 sync.Mutex
	__obf_c61d4e1ec93a0d69 time.Time
	time                   int64
	__obf_0203258bc566ede9 int64
	__obf_b3c91dc2f5c97abd int64
	__obf_fa14daecff57d827 int64
	__obf_8a3309aeba4c3e95 int64
	__obf_c23664664a7762d5 int64
	__obf_462febdd1e69694f uint8
	__obf_6df89a1a9eab32eb uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_0203258bc566ede9 int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	__obf_15719f79f826bdc4.
		// re-calc in case custom NodeBits or StepBits were set
		// DEPRECATED: the below block will be removed in a future release.
		Lock()
	__obf_fa14daecff57d827 = -1 ^ (-1 << NodeBits)
	__obf_8a3309aeba4c3e95 = __obf_fa14daecff57d827 << StepBits
	__obf_c23664664a7762d5 = -1 ^ (-1 << StepBits)
	__obf_462febdd1e69694f = NodeBits + StepBits
	__obf_6df89a1a9eab32eb = StepBits
	__obf_15719f79f826bdc4.
		Unlock()
	__obf_479576c1f28f3885 := Node{}
	__obf_479576c1f28f3885.__obf_0203258bc566ede9 = __obf_0203258bc566ede9
	__obf_479576c1f28f3885.__obf_fa14daecff57d827 = -1 ^ (-1 << NodeBits)
	__obf_479576c1f28f3885.__obf_8a3309aeba4c3e95 = __obf_479576c1f28f3885.__obf_fa14daecff57d827 << StepBits
	__obf_479576c1f28f3885.__obf_c23664664a7762d5 = -1 ^ (-1 << StepBits)
	__obf_479576c1f28f3885.__obf_462febdd1e69694f = NodeBits + StepBits
	__obf_479576c1f28f3885.__obf_6df89a1a9eab32eb = StepBits

	if __obf_479576c1f28f3885.__obf_0203258bc566ede9 < 0 || __obf_479576c1f28f3885.__obf_0203258bc566ede9 > __obf_479576c1f28f3885.__obf_fa14daecff57d827 {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_479576c1f28f3885.__obf_fa14daecff57d827, 10))
	}

	var __obf_9146c4d9abb93522 = time.Now()
	__obf_479576c1f28f3885.
		// add time.Duration to curTime to make sure we use the monotonic clock if available
		__obf_c61d4e1ec93a0d69 = __obf_9146c4d9abb93522.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_9146c4d9abb93522))

	return &__obf_479576c1f28f3885, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_479576c1f28f3885 *Node) Generate() ID {
	__obf_479576c1f28f3885.__obf_15719f79f826bdc4.
		Lock()
	defer __obf_479576c1f28f3885.__obf_15719f79f826bdc4.Unlock()
	__obf_88d7ac29d5972512 := time.Since(__obf_479576c1f28f3885.__obf_c61d4e1ec93a0d69).Milliseconds()

	if __obf_88d7ac29d5972512 == __obf_479576c1f28f3885.time {
		__obf_479576c1f28f3885.__obf_b3c91dc2f5c97abd = (__obf_479576c1f28f3885.__obf_b3c91dc2f5c97abd + 1) & __obf_479576c1f28f3885.__obf_c23664664a7762d5

		if __obf_479576c1f28f3885.__obf_b3c91dc2f5c97abd == 0 {
			for __obf_88d7ac29d5972512 <= __obf_479576c1f28f3885.time {
				__obf_88d7ac29d5972512 = time.Since(__obf_479576c1f28f3885.__obf_c61d4e1ec93a0d69).Milliseconds()
			}
		}
	} else {
		__obf_479576c1f28f3885.__obf_b3c91dc2f5c97abd = 0
	}
	__obf_479576c1f28f3885.
		time = __obf_88d7ac29d5972512
	__obf_4c01cf6df5b9c963 := ID((__obf_88d7ac29d5972512)<<__obf_479576c1f28f3885.__obf_462febdd1e69694f |
		(__obf_479576c1f28f3885.__obf_0203258bc566ede9 << __obf_479576c1f28f3885.__obf_6df89a1a9eab32eb) |
		(__obf_479576c1f28f3885.__obf_b3c91dc2f5c97abd),
	)

	return __obf_4c01cf6df5b9c963
}

// Int64 returns an int64 of the snowflake ID
func (__obf_7d188f436128d50d ID) Int64() int64 {
	return int64(__obf_7d188f436128d50d)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_28d51b2675aee39d int64) ID {
	return ID(__obf_28d51b2675aee39d)
}

// String returns a string of the snowflake ID
func (__obf_7d188f436128d50d ID) String() string {
	return strconv.FormatInt(int64(__obf_7d188f436128d50d), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_28d51b2675aee39d string) (ID, error) {
	__obf_1ef2f5d835c97294, __obf_116d5a689b79a523 := strconv.ParseInt(__obf_28d51b2675aee39d, 10, 64)
	return ID(__obf_1ef2f5d835c97294), __obf_116d5a689b79a523

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_7d188f436128d50d ID) Base2() string {
	return strconv.FormatInt(int64(__obf_7d188f436128d50d), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_28d51b2675aee39d string) (ID, error) {
	__obf_1ef2f5d835c97294, __obf_116d5a689b79a523 := strconv.ParseInt(__obf_28d51b2675aee39d, 2, 64)
	return ID(__obf_1ef2f5d835c97294), __obf_116d5a689b79a523
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_7d188f436128d50d ID) Base32() string {

	if __obf_7d188f436128d50d < 32 {
		return string(__obf_3de1fb88ae0b4aee[__obf_7d188f436128d50d])
	}
	__obf_bd937547bf71516d := make([]byte, 0, 12)
	for __obf_7d188f436128d50d >= 32 {
		__obf_bd937547bf71516d = append(__obf_bd937547bf71516d, __obf_3de1fb88ae0b4aee[__obf_7d188f436128d50d%32])
		__obf_7d188f436128d50d /= 32
	}
	__obf_bd937547bf71516d = append(__obf_bd937547bf71516d, __obf_3de1fb88ae0b4aee[__obf_7d188f436128d50d])

	for __obf_d139d546d54aa8c8, __obf_ebb354505391a30e := 0, len(__obf_bd937547bf71516d)-1; __obf_d139d546d54aa8c8 < __obf_ebb354505391a30e; __obf_d139d546d54aa8c8, __obf_ebb354505391a30e = __obf_d139d546d54aa8c8+1, __obf_ebb354505391a30e-1 {
		__obf_bd937547bf71516d[__obf_d139d546d54aa8c8], __obf_bd937547bf71516d[__obf_ebb354505391a30e] = __obf_bd937547bf71516d[__obf_ebb354505391a30e], __obf_bd937547bf71516d[__obf_d139d546d54aa8c8]
	}

	return string(__obf_bd937547bf71516d)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_bd937547bf71516d []byte) (ID, error) {

	var __obf_28d51b2675aee39d int64

	for __obf_1ef2f5d835c97294 := range __obf_bd937547bf71516d {
		if __obf_99dec5fdbabbf1cc[__obf_bd937547bf71516d[__obf_1ef2f5d835c97294]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_28d51b2675aee39d = __obf_28d51b2675aee39d*32 + int64(__obf_99dec5fdbabbf1cc[__obf_bd937547bf71516d[__obf_1ef2f5d835c97294]])
	}

	return ID(__obf_28d51b2675aee39d), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_7d188f436128d50d ID) Base36() string {
	return strconv.FormatInt(int64(__obf_7d188f436128d50d), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_28d51b2675aee39d string) (ID, error) {
	__obf_1ef2f5d835c97294, __obf_116d5a689b79a523 := strconv.ParseInt(__obf_28d51b2675aee39d, 36, 64)
	return ID(__obf_1ef2f5d835c97294), __obf_116d5a689b79a523
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_7d188f436128d50d ID) Base58() string {

	if __obf_7d188f436128d50d < 58 {
		return string(__obf_e969c0efc0d267e0[__obf_7d188f436128d50d])
	}
	__obf_bd937547bf71516d := make([]byte, 0, 11)
	for __obf_7d188f436128d50d >= 58 {
		__obf_bd937547bf71516d = append(__obf_bd937547bf71516d, __obf_e969c0efc0d267e0[__obf_7d188f436128d50d%58])
		__obf_7d188f436128d50d /= 58
	}
	__obf_bd937547bf71516d = append(__obf_bd937547bf71516d, __obf_e969c0efc0d267e0[__obf_7d188f436128d50d])

	for __obf_d139d546d54aa8c8, __obf_ebb354505391a30e := 0, len(__obf_bd937547bf71516d)-1; __obf_d139d546d54aa8c8 < __obf_ebb354505391a30e; __obf_d139d546d54aa8c8, __obf_ebb354505391a30e = __obf_d139d546d54aa8c8+1, __obf_ebb354505391a30e-1 {
		__obf_bd937547bf71516d[__obf_d139d546d54aa8c8], __obf_bd937547bf71516d[__obf_ebb354505391a30e] = __obf_bd937547bf71516d[__obf_ebb354505391a30e], __obf_bd937547bf71516d[__obf_d139d546d54aa8c8]
	}

	return string(__obf_bd937547bf71516d)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_bd937547bf71516d []byte) (ID, error) {

	var __obf_28d51b2675aee39d int64

	for __obf_1ef2f5d835c97294 := range __obf_bd937547bf71516d {
		if __obf_5bbf872fd814e2b3[__obf_bd937547bf71516d[__obf_1ef2f5d835c97294]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_28d51b2675aee39d = __obf_28d51b2675aee39d*58 + int64(__obf_5bbf872fd814e2b3[__obf_bd937547bf71516d[__obf_1ef2f5d835c97294]])
	}

	return ID(__obf_28d51b2675aee39d), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_7d188f436128d50d ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_7d188f436128d50d.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_28d51b2675aee39d string) (ID, error) {
	__obf_bd937547bf71516d, __obf_116d5a689b79a523 := base64.StdEncoding.DecodeString(__obf_28d51b2675aee39d)
	if __obf_116d5a689b79a523 != nil {
		return -1, __obf_116d5a689b79a523
	}
	return ParseBytes(__obf_bd937547bf71516d)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_7d188f436128d50d ID) Bytes() []byte {
	return []byte(__obf_7d188f436128d50d.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_28d51b2675aee39d []byte) (ID, error) {
	__obf_1ef2f5d835c97294, __obf_116d5a689b79a523 := strconv.ParseInt(string(__obf_28d51b2675aee39d), 10, 64)
	return ID(__obf_1ef2f5d835c97294), __obf_116d5a689b79a523
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_7d188f436128d50d ID) IntBytes() [8]byte {
	var __obf_bd937547bf71516d [8]byte
	binary.BigEndian.PutUint64(__obf_bd937547bf71516d[:], uint64(__obf_7d188f436128d50d))
	return __obf_bd937547bf71516d
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_28d51b2675aee39d [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_28d51b2675aee39d[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_7d188f436128d50d ID) Time() int64 {
	return (int64(__obf_7d188f436128d50d) >> __obf_462febdd1e69694f) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_7d188f436128d50d ID) Node() int64 {
	return int64(__obf_7d188f436128d50d) & __obf_8a3309aeba4c3e95 >> __obf_6df89a1a9eab32eb
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_7d188f436128d50d ID) Step() int64 {
	return int64(__obf_7d188f436128d50d) & __obf_c23664664a7762d5
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_7d188f436128d50d ID) MarshalJSON() ([]byte, error) {
	__obf_a97f5b468fdea060 := make([]byte, 0, 22)
	__obf_a97f5b468fdea060 = append(__obf_a97f5b468fdea060, '"')
	__obf_a97f5b468fdea060 = strconv.AppendInt(__obf_a97f5b468fdea060, int64(__obf_7d188f436128d50d), 10)
	__obf_a97f5b468fdea060 = append(__obf_a97f5b468fdea060, '"')
	return __obf_a97f5b468fdea060, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_7d188f436128d50d *ID) UnmarshalJSON(__obf_bd937547bf71516d []byte) error {
	if len(__obf_bd937547bf71516d) < 3 || __obf_bd937547bf71516d[0] != '"' || __obf_bd937547bf71516d[len(__obf_bd937547bf71516d)-1] != '"' {
		return JSONSyntaxError{__obf_bd937547bf71516d}
	}
	__obf_1ef2f5d835c97294, __obf_116d5a689b79a523 := strconv.ParseInt(string(__obf_bd937547bf71516d[1:len(__obf_bd937547bf71516d)-1]), 10, 64)
	if __obf_116d5a689b79a523 != nil {
		return __obf_116d5a689b79a523
	}

	*__obf_7d188f436128d50d = ID(__obf_1ef2f5d835c97294)
	return nil
}
