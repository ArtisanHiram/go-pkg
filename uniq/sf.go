package __obf_2f51f7d26a2bcdf8

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
	__obf_54867e7d58cb9039 sync. // DEPRECATED: the below four variables will be removed in a future release.
				Mutex
	__obf_0c27fd5797be7365 int64 = -1 ^ (-1 << NodeBits)
	__obf_3fac3a080d5e1605       = __obf_0c27fd5797be7365 <<
		StepBits
	__obf_d6aad654a8400fc3 int64 = -1 ^ (-1 << StepBits)
	__obf_d391aada5b93f61b       = NodeBits + StepBits
	__obf_17ecb59970d4387f       = StepBits
)

const __obf_4529985d9efef43d = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_bfff0cc1a11dc29a [256]byte

const __obf_15ae07e01e73f253 = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_f55c1f939827e050 [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_fbad881a9432eb4d []byte }

func (__obf_d616bb3ceeb71621 JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_d616bb3ceeb71621.

		// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
		__obf_fbad881a9432eb4d))
}

var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func init() {

	for __obf_9a4933bb8fba4943 := 0; __obf_9a4933bb8fba4943 < len(__obf_f55c1f939827e050); __obf_9a4933bb8fba4943++ {
		__obf_f55c1f939827e050[__obf_9a4933bb8fba4943] = 0xFF
	}

	for __obf_9a4933bb8fba4943 := 0; __obf_9a4933bb8fba4943 < len(__obf_15ae07e01e73f253); __obf_9a4933bb8fba4943++ {
		__obf_f55c1f939827e050[__obf_15ae07e01e73f253[__obf_9a4933bb8fba4943]] = byte(__obf_9a4933bb8fba4943)
	}

	for __obf_9a4933bb8fba4943 := 0; __obf_9a4933bb8fba4943 < len(__obf_bfff0cc1a11dc29a); __obf_9a4933bb8fba4943++ {
		__obf_bfff0cc1a11dc29a[__obf_9a4933bb8fba4943] = 0xFF
	}

	for __obf_9a4933bb8fba4943 := 0; __obf_9a4933bb8fba4943 < len(__obf_4529985d9efef43d); __obf_9a4933bb8fba4943++ {
		__obf_bfff0cc1a11dc29a[__obf_4529985d9efef43d[__obf_9a4933bb8fba4943]] = byte(__obf_9a4933bb8fba4943)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_54867e7d58cb9039 sync.Mutex
	__obf_3eace83c308ae04b time.Time
	time                   int64
	__obf_832cfdd9f675e3b5 int64
	__obf_0a7d59539d2ddda8 int64
	__obf_0c27fd5797be7365 int64
	__obf_3fac3a080d5e1605 int64
	__obf_d6aad654a8400fc3 int64
	__obf_d391aada5b93f61b uint8
	__obf_17ecb59970d4387f uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_832cfdd9f675e3b5 int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	__obf_54867e7d58cb9039.
		// re-calc in case custom NodeBits or StepBits were set
		// DEPRECATED: the below block will be removed in a future release.
		Lock()
	__obf_0c27fd5797be7365 = -1 ^ (-1 << NodeBits)
	__obf_3fac3a080d5e1605 = __obf_0c27fd5797be7365 << StepBits
	__obf_d6aad654a8400fc3 = -1 ^ (-1 << StepBits)
	__obf_d391aada5b93f61b = NodeBits + StepBits
	__obf_17ecb59970d4387f = StepBits
	__obf_54867e7d58cb9039.
		Unlock()
	__obf_f88ad57451625a6d := Node{}
	__obf_f88ad57451625a6d.__obf_832cfdd9f675e3b5 = __obf_832cfdd9f675e3b5
	__obf_f88ad57451625a6d.__obf_0c27fd5797be7365 = -1 ^ (-1 << NodeBits)
	__obf_f88ad57451625a6d.__obf_3fac3a080d5e1605 = __obf_f88ad57451625a6d.__obf_0c27fd5797be7365 << StepBits
	__obf_f88ad57451625a6d.__obf_d6aad654a8400fc3 = -1 ^ (-1 << StepBits)
	__obf_f88ad57451625a6d.__obf_d391aada5b93f61b = NodeBits + StepBits
	__obf_f88ad57451625a6d.__obf_17ecb59970d4387f = StepBits

	if __obf_f88ad57451625a6d.__obf_832cfdd9f675e3b5 < 0 || __obf_f88ad57451625a6d.__obf_832cfdd9f675e3b5 > __obf_f88ad57451625a6d.__obf_0c27fd5797be7365 {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_f88ad57451625a6d.__obf_0c27fd5797be7365, 10))
	}

	var __obf_e1ebea7b7a8675e4 = time.Now()
	__obf_f88ad57451625a6d.
		// add time.Duration to curTime to make sure we use the monotonic clock if available
		__obf_3eace83c308ae04b = __obf_e1ebea7b7a8675e4.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_e1ebea7b7a8675e4))

	return &__obf_f88ad57451625a6d, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_f88ad57451625a6d *Node) Generate() ID {
	__obf_f88ad57451625a6d.__obf_54867e7d58cb9039.
		Lock()
	defer __obf_f88ad57451625a6d.__obf_54867e7d58cb9039.Unlock()
	__obf_6a9c64860dbb489c := time.Since(__obf_f88ad57451625a6d.__obf_3eace83c308ae04b).Milliseconds()

	if __obf_6a9c64860dbb489c == __obf_f88ad57451625a6d.time {
		__obf_f88ad57451625a6d.__obf_0a7d59539d2ddda8 = (__obf_f88ad57451625a6d.__obf_0a7d59539d2ddda8 + 1) & __obf_f88ad57451625a6d.__obf_d6aad654a8400fc3

		if __obf_f88ad57451625a6d.__obf_0a7d59539d2ddda8 == 0 {
			for __obf_6a9c64860dbb489c <= __obf_f88ad57451625a6d.time {
				__obf_6a9c64860dbb489c = time.Since(__obf_f88ad57451625a6d.__obf_3eace83c308ae04b).Milliseconds()
			}
		}
	} else {
		__obf_f88ad57451625a6d.__obf_0a7d59539d2ddda8 = 0
	}
	__obf_f88ad57451625a6d.
		time = __obf_6a9c64860dbb489c
	__obf_3b7a9d8df60cb6ee := ID((__obf_6a9c64860dbb489c)<<__obf_f88ad57451625a6d.__obf_d391aada5b93f61b |
		(__obf_f88ad57451625a6d.__obf_832cfdd9f675e3b5 << __obf_f88ad57451625a6d.__obf_17ecb59970d4387f) |
		(__obf_f88ad57451625a6d.__obf_0a7d59539d2ddda8),
	)

	return __obf_3b7a9d8df60cb6ee
}

// Int64 returns an int64 of the snowflake ID
func (__obf_c16b13459661e274 ID) Int64() int64 {
	return int64(__obf_c16b13459661e274)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_8b246d7b3659af37 int64) ID {
	return ID(__obf_8b246d7b3659af37)
}

// String returns a string of the snowflake ID
func (__obf_c16b13459661e274 ID) String() string {
	return strconv.FormatInt(int64(__obf_c16b13459661e274), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_8b246d7b3659af37 string) (ID, error) {
	__obf_9a4933bb8fba4943, __obf_37e8e388755917c0 := strconv.ParseInt(__obf_8b246d7b3659af37, 10, 64)
	return ID(__obf_9a4933bb8fba4943), __obf_37e8e388755917c0

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_c16b13459661e274 ID) Base2() string {
	return strconv.FormatInt(int64(__obf_c16b13459661e274), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_8b246d7b3659af37 string) (ID, error) {
	__obf_9a4933bb8fba4943, __obf_37e8e388755917c0 := strconv.ParseInt(__obf_8b246d7b3659af37, 2, 64)
	return ID(__obf_9a4933bb8fba4943), __obf_37e8e388755917c0
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_c16b13459661e274 ID) Base32() string {

	if __obf_c16b13459661e274 < 32 {
		return string(__obf_4529985d9efef43d[__obf_c16b13459661e274])
	}
	__obf_e0252f35dbdc2381 := make([]byte, 0, 12)
	for __obf_c16b13459661e274 >= 32 {
		__obf_e0252f35dbdc2381 = append(__obf_e0252f35dbdc2381, __obf_4529985d9efef43d[__obf_c16b13459661e274%32])
		__obf_c16b13459661e274 /= 32
	}
	__obf_e0252f35dbdc2381 = append(__obf_e0252f35dbdc2381, __obf_4529985d9efef43d[__obf_c16b13459661e274])

	for __obf_caddb1840dda593a, __obf_2016ba9ef586aed9 := 0, len(__obf_e0252f35dbdc2381)-1; __obf_caddb1840dda593a < __obf_2016ba9ef586aed9; __obf_caddb1840dda593a, __obf_2016ba9ef586aed9 = __obf_caddb1840dda593a+1, __obf_2016ba9ef586aed9-1 {
		__obf_e0252f35dbdc2381[__obf_caddb1840dda593a], __obf_e0252f35dbdc2381[__obf_2016ba9ef586aed9] = __obf_e0252f35dbdc2381[__obf_2016ba9ef586aed9], __obf_e0252f35dbdc2381[__obf_caddb1840dda593a]
	}

	return string(__obf_e0252f35dbdc2381)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_e0252f35dbdc2381 []byte) (ID, error) {

	var __obf_8b246d7b3659af37 int64

	for __obf_9a4933bb8fba4943 := range __obf_e0252f35dbdc2381 {
		if __obf_bfff0cc1a11dc29a[__obf_e0252f35dbdc2381[__obf_9a4933bb8fba4943]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_8b246d7b3659af37 = __obf_8b246d7b3659af37*32 + int64(__obf_bfff0cc1a11dc29a[__obf_e0252f35dbdc2381[__obf_9a4933bb8fba4943]])
	}

	return ID(__obf_8b246d7b3659af37), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_c16b13459661e274 ID) Base36() string {
	return strconv.FormatInt(int64(__obf_c16b13459661e274), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_8b246d7b3659af37 string) (ID, error) {
	__obf_9a4933bb8fba4943, __obf_37e8e388755917c0 := strconv.ParseInt(__obf_8b246d7b3659af37, 36, 64)
	return ID(__obf_9a4933bb8fba4943), __obf_37e8e388755917c0
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_c16b13459661e274 ID) Base58() string {

	if __obf_c16b13459661e274 < 58 {
		return string(__obf_15ae07e01e73f253[__obf_c16b13459661e274])
	}
	__obf_e0252f35dbdc2381 := make([]byte, 0, 11)
	for __obf_c16b13459661e274 >= 58 {
		__obf_e0252f35dbdc2381 = append(__obf_e0252f35dbdc2381, __obf_15ae07e01e73f253[__obf_c16b13459661e274%58])
		__obf_c16b13459661e274 /= 58
	}
	__obf_e0252f35dbdc2381 = append(__obf_e0252f35dbdc2381, __obf_15ae07e01e73f253[__obf_c16b13459661e274])

	for __obf_caddb1840dda593a, __obf_2016ba9ef586aed9 := 0, len(__obf_e0252f35dbdc2381)-1; __obf_caddb1840dda593a < __obf_2016ba9ef586aed9; __obf_caddb1840dda593a, __obf_2016ba9ef586aed9 = __obf_caddb1840dda593a+1, __obf_2016ba9ef586aed9-1 {
		__obf_e0252f35dbdc2381[__obf_caddb1840dda593a], __obf_e0252f35dbdc2381[__obf_2016ba9ef586aed9] = __obf_e0252f35dbdc2381[__obf_2016ba9ef586aed9], __obf_e0252f35dbdc2381[__obf_caddb1840dda593a]
	}

	return string(__obf_e0252f35dbdc2381)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_e0252f35dbdc2381 []byte) (ID, error) {

	var __obf_8b246d7b3659af37 int64

	for __obf_9a4933bb8fba4943 := range __obf_e0252f35dbdc2381 {
		if __obf_f55c1f939827e050[__obf_e0252f35dbdc2381[__obf_9a4933bb8fba4943]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_8b246d7b3659af37 = __obf_8b246d7b3659af37*58 + int64(__obf_f55c1f939827e050[__obf_e0252f35dbdc2381[__obf_9a4933bb8fba4943]])
	}

	return ID(__obf_8b246d7b3659af37), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_c16b13459661e274 ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_c16b13459661e274.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_8b246d7b3659af37 string) (ID, error) {
	__obf_e0252f35dbdc2381, __obf_37e8e388755917c0 := base64.StdEncoding.DecodeString(__obf_8b246d7b3659af37)
	if __obf_37e8e388755917c0 != nil {
		return -1, __obf_37e8e388755917c0
	}
	return ParseBytes(__obf_e0252f35dbdc2381)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_c16b13459661e274 ID) Bytes() []byte {
	return []byte(__obf_c16b13459661e274.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_8b246d7b3659af37 []byte) (ID, error) {
	__obf_9a4933bb8fba4943, __obf_37e8e388755917c0 := strconv.ParseInt(string(__obf_8b246d7b3659af37), 10, 64)
	return ID(__obf_9a4933bb8fba4943), __obf_37e8e388755917c0
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_c16b13459661e274 ID) IntBytes() [8]byte {
	var __obf_e0252f35dbdc2381 [8]byte
	binary.BigEndian.PutUint64(__obf_e0252f35dbdc2381[:], uint64(__obf_c16b13459661e274))
	return __obf_e0252f35dbdc2381
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_8b246d7b3659af37 [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_8b246d7b3659af37[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_c16b13459661e274 ID) Time() int64 {
	return (int64(__obf_c16b13459661e274) >> __obf_d391aada5b93f61b) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_c16b13459661e274 ID) Node() int64 {
	return int64(__obf_c16b13459661e274) & __obf_3fac3a080d5e1605 >> __obf_17ecb59970d4387f
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_c16b13459661e274 ID) Step() int64 {
	return int64(__obf_c16b13459661e274) & __obf_d6aad654a8400fc3
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_c16b13459661e274 ID) MarshalJSON() ([]byte, error) {
	__obf_58aef99dff4a32d5 := make([]byte, 0, 22)
	__obf_58aef99dff4a32d5 = append(__obf_58aef99dff4a32d5, '"')
	__obf_58aef99dff4a32d5 = strconv.AppendInt(__obf_58aef99dff4a32d5, int64(__obf_c16b13459661e274), 10)
	__obf_58aef99dff4a32d5 = append(__obf_58aef99dff4a32d5, '"')
	return __obf_58aef99dff4a32d5, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_c16b13459661e274 *ID) UnmarshalJSON(__obf_e0252f35dbdc2381 []byte) error {
	if len(__obf_e0252f35dbdc2381) < 3 || __obf_e0252f35dbdc2381[0] != '"' || __obf_e0252f35dbdc2381[len(__obf_e0252f35dbdc2381)-1] != '"' {
		return JSONSyntaxError{__obf_e0252f35dbdc2381}
	}
	__obf_9a4933bb8fba4943, __obf_37e8e388755917c0 := strconv.ParseInt(string(__obf_e0252f35dbdc2381[1:len(__obf_e0252f35dbdc2381)-1]), 10, 64)
	if __obf_37e8e388755917c0 != nil {
		return __obf_37e8e388755917c0
	}

	*__obf_c16b13459661e274 = ID(__obf_9a4933bb8fba4943)
	return nil
}
