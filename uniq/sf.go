package __obf_3747a7e09ff475ee

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
	__obf_1d87674837946b18 sync. // DEPRECATED: the below four variables will be removed in a future release.
				Mutex
	__obf_fd2a505fc157bc5b int64 = -1 ^ (-1 << NodeBits)
	__obf_bbc1c63b768b03cc       = __obf_fd2a505fc157bc5b <<
		StepBits
	__obf_4bb5a3804ded0828 int64 = -1 ^ (-1 << StepBits)
	__obf_444c3eed855a9bba       = NodeBits + StepBits
	__obf_e19cb3423c33812f       = StepBits
)

const __obf_bd74d6975be021f8 = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_ab01d6580a95ebfe [256]byte

const __obf_e642f800dcf673ec = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_07cd93227dd89c33 [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_a03bca2ee0b29d36 []byte }

func (__obf_f198df20fe4f77ea JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_f198df20fe4f77ea.

		// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
		__obf_a03bca2ee0b29d36))
}

var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func init() {

	for __obf_a4042898d22f95f6 := 0; __obf_a4042898d22f95f6 < len(__obf_07cd93227dd89c33); __obf_a4042898d22f95f6++ {
		__obf_07cd93227dd89c33[__obf_a4042898d22f95f6] = 0xFF
	}

	for __obf_a4042898d22f95f6 := 0; __obf_a4042898d22f95f6 < len(__obf_e642f800dcf673ec); __obf_a4042898d22f95f6++ {
		__obf_07cd93227dd89c33[__obf_e642f800dcf673ec[__obf_a4042898d22f95f6]] = byte(__obf_a4042898d22f95f6)
	}

	for __obf_a4042898d22f95f6 := 0; __obf_a4042898d22f95f6 < len(__obf_ab01d6580a95ebfe); __obf_a4042898d22f95f6++ {
		__obf_ab01d6580a95ebfe[__obf_a4042898d22f95f6] = 0xFF
	}

	for __obf_a4042898d22f95f6 := 0; __obf_a4042898d22f95f6 < len(__obf_bd74d6975be021f8); __obf_a4042898d22f95f6++ {
		__obf_ab01d6580a95ebfe[__obf_bd74d6975be021f8[__obf_a4042898d22f95f6]] = byte(__obf_a4042898d22f95f6)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_1d87674837946b18 sync.Mutex
	__obf_a6ed45b4c5339875 time.Time
	time                   int64
	__obf_487abfe58b7ebec3 int64
	__obf_b68e8c48d58e1dca int64
	__obf_fd2a505fc157bc5b int64
	__obf_bbc1c63b768b03cc int64
	__obf_4bb5a3804ded0828 int64
	__obf_444c3eed855a9bba uint8
	__obf_e19cb3423c33812f uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_487abfe58b7ebec3 int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	__obf_1d87674837946b18.
		// re-calc in case custom NodeBits or StepBits were set
		// DEPRECATED: the below block will be removed in a future release.
		Lock()
	__obf_fd2a505fc157bc5b = -1 ^ (-1 << NodeBits)
	__obf_bbc1c63b768b03cc = __obf_fd2a505fc157bc5b << StepBits
	__obf_4bb5a3804ded0828 = -1 ^ (-1 << StepBits)
	__obf_444c3eed855a9bba = NodeBits + StepBits
	__obf_e19cb3423c33812f = StepBits
	__obf_1d87674837946b18.
		Unlock()
	__obf_368df23032f058d2 := Node{}
	__obf_368df23032f058d2.__obf_487abfe58b7ebec3 = __obf_487abfe58b7ebec3
	__obf_368df23032f058d2.__obf_fd2a505fc157bc5b = -1 ^ (-1 << NodeBits)
	__obf_368df23032f058d2.__obf_bbc1c63b768b03cc = __obf_368df23032f058d2.__obf_fd2a505fc157bc5b << StepBits
	__obf_368df23032f058d2.__obf_4bb5a3804ded0828 = -1 ^ (-1 << StepBits)
	__obf_368df23032f058d2.__obf_444c3eed855a9bba = NodeBits + StepBits
	__obf_368df23032f058d2.__obf_e19cb3423c33812f = StepBits

	if __obf_368df23032f058d2.__obf_487abfe58b7ebec3 < 0 || __obf_368df23032f058d2.__obf_487abfe58b7ebec3 > __obf_368df23032f058d2.__obf_fd2a505fc157bc5b {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_368df23032f058d2.__obf_fd2a505fc157bc5b, 10))
	}

	var __obf_37931928ce9fa001 = time.Now()
	__obf_368df23032f058d2.
		// add time.Duration to curTime to make sure we use the monotonic clock if available
		__obf_a6ed45b4c5339875 = __obf_37931928ce9fa001.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_37931928ce9fa001))

	return &__obf_368df23032f058d2, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_368df23032f058d2 *Node) Generate() ID {
	__obf_368df23032f058d2.__obf_1d87674837946b18.
		Lock()
	defer __obf_368df23032f058d2.__obf_1d87674837946b18.Unlock()
	__obf_6df2298cc6241599 := time.Since(__obf_368df23032f058d2.__obf_a6ed45b4c5339875).Milliseconds()

	if __obf_6df2298cc6241599 == __obf_368df23032f058d2.time {
		__obf_368df23032f058d2.__obf_b68e8c48d58e1dca = (__obf_368df23032f058d2.__obf_b68e8c48d58e1dca + 1) & __obf_368df23032f058d2.__obf_4bb5a3804ded0828

		if __obf_368df23032f058d2.__obf_b68e8c48d58e1dca == 0 {
			for __obf_6df2298cc6241599 <= __obf_368df23032f058d2.time {
				__obf_6df2298cc6241599 = time.Since(__obf_368df23032f058d2.__obf_a6ed45b4c5339875).Milliseconds()
			}
		}
	} else {
		__obf_368df23032f058d2.__obf_b68e8c48d58e1dca = 0
	}
	__obf_368df23032f058d2.
		time = __obf_6df2298cc6241599
	__obf_64990aa27d4bceaa := ID((__obf_6df2298cc6241599)<<__obf_368df23032f058d2.__obf_444c3eed855a9bba |
		(__obf_368df23032f058d2.__obf_487abfe58b7ebec3 << __obf_368df23032f058d2.__obf_e19cb3423c33812f) |
		(__obf_368df23032f058d2.__obf_b68e8c48d58e1dca),
	)

	return __obf_64990aa27d4bceaa
}

// Int64 returns an int64 of the snowflake ID
func (__obf_5f66c6dacd23d45f ID) Int64() int64 {
	return int64(__obf_5f66c6dacd23d45f)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_413d3a76ff8ab79e int64) ID {
	return ID(__obf_413d3a76ff8ab79e)
}

// String returns a string of the snowflake ID
func (__obf_5f66c6dacd23d45f ID) String() string {
	return strconv.FormatInt(int64(__obf_5f66c6dacd23d45f), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_413d3a76ff8ab79e string) (ID, error) {
	__obf_a4042898d22f95f6, __obf_57cef2f010b4ce9a := strconv.ParseInt(__obf_413d3a76ff8ab79e, 10, 64)
	return ID(__obf_a4042898d22f95f6), __obf_57cef2f010b4ce9a

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_5f66c6dacd23d45f ID) Base2() string {
	return strconv.FormatInt(int64(__obf_5f66c6dacd23d45f), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_413d3a76ff8ab79e string) (ID, error) {
	__obf_a4042898d22f95f6, __obf_57cef2f010b4ce9a := strconv.ParseInt(__obf_413d3a76ff8ab79e, 2, 64)
	return ID(__obf_a4042898d22f95f6), __obf_57cef2f010b4ce9a
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_5f66c6dacd23d45f ID) Base32() string {

	if __obf_5f66c6dacd23d45f < 32 {
		return string(__obf_bd74d6975be021f8[__obf_5f66c6dacd23d45f])
	}
	__obf_9b18a3f8d8063793 := make([]byte, 0, 12)
	for __obf_5f66c6dacd23d45f >= 32 {
		__obf_9b18a3f8d8063793 = append(__obf_9b18a3f8d8063793, __obf_bd74d6975be021f8[__obf_5f66c6dacd23d45f%32])
		__obf_5f66c6dacd23d45f /= 32
	}
	__obf_9b18a3f8d8063793 = append(__obf_9b18a3f8d8063793, __obf_bd74d6975be021f8[__obf_5f66c6dacd23d45f])

	for __obf_d2ffb8381b91d2b8, __obf_9bd3c328cc69f819 := 0, len(__obf_9b18a3f8d8063793)-1; __obf_d2ffb8381b91d2b8 < __obf_9bd3c328cc69f819; __obf_d2ffb8381b91d2b8, __obf_9bd3c328cc69f819 = __obf_d2ffb8381b91d2b8+1, __obf_9bd3c328cc69f819-1 {
		__obf_9b18a3f8d8063793[__obf_d2ffb8381b91d2b8], __obf_9b18a3f8d8063793[__obf_9bd3c328cc69f819] = __obf_9b18a3f8d8063793[__obf_9bd3c328cc69f819], __obf_9b18a3f8d8063793[__obf_d2ffb8381b91d2b8]
	}

	return string(__obf_9b18a3f8d8063793)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_9b18a3f8d8063793 []byte) (ID, error) {

	var __obf_413d3a76ff8ab79e int64

	for __obf_a4042898d22f95f6 := range __obf_9b18a3f8d8063793 {
		if __obf_ab01d6580a95ebfe[__obf_9b18a3f8d8063793[__obf_a4042898d22f95f6]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_413d3a76ff8ab79e = __obf_413d3a76ff8ab79e*32 + int64(__obf_ab01d6580a95ebfe[__obf_9b18a3f8d8063793[__obf_a4042898d22f95f6]])
	}

	return ID(__obf_413d3a76ff8ab79e), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_5f66c6dacd23d45f ID) Base36() string {
	return strconv.FormatInt(int64(__obf_5f66c6dacd23d45f), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_413d3a76ff8ab79e string) (ID, error) {
	__obf_a4042898d22f95f6, __obf_57cef2f010b4ce9a := strconv.ParseInt(__obf_413d3a76ff8ab79e, 36, 64)
	return ID(__obf_a4042898d22f95f6), __obf_57cef2f010b4ce9a
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_5f66c6dacd23d45f ID) Base58() string {

	if __obf_5f66c6dacd23d45f < 58 {
		return string(__obf_e642f800dcf673ec[__obf_5f66c6dacd23d45f])
	}
	__obf_9b18a3f8d8063793 := make([]byte, 0, 11)
	for __obf_5f66c6dacd23d45f >= 58 {
		__obf_9b18a3f8d8063793 = append(__obf_9b18a3f8d8063793, __obf_e642f800dcf673ec[__obf_5f66c6dacd23d45f%58])
		__obf_5f66c6dacd23d45f /= 58
	}
	__obf_9b18a3f8d8063793 = append(__obf_9b18a3f8d8063793, __obf_e642f800dcf673ec[__obf_5f66c6dacd23d45f])

	for __obf_d2ffb8381b91d2b8, __obf_9bd3c328cc69f819 := 0, len(__obf_9b18a3f8d8063793)-1; __obf_d2ffb8381b91d2b8 < __obf_9bd3c328cc69f819; __obf_d2ffb8381b91d2b8, __obf_9bd3c328cc69f819 = __obf_d2ffb8381b91d2b8+1, __obf_9bd3c328cc69f819-1 {
		__obf_9b18a3f8d8063793[__obf_d2ffb8381b91d2b8], __obf_9b18a3f8d8063793[__obf_9bd3c328cc69f819] = __obf_9b18a3f8d8063793[__obf_9bd3c328cc69f819], __obf_9b18a3f8d8063793[__obf_d2ffb8381b91d2b8]
	}

	return string(__obf_9b18a3f8d8063793)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_9b18a3f8d8063793 []byte) (ID, error) {

	var __obf_413d3a76ff8ab79e int64

	for __obf_a4042898d22f95f6 := range __obf_9b18a3f8d8063793 {
		if __obf_07cd93227dd89c33[__obf_9b18a3f8d8063793[__obf_a4042898d22f95f6]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_413d3a76ff8ab79e = __obf_413d3a76ff8ab79e*58 + int64(__obf_07cd93227dd89c33[__obf_9b18a3f8d8063793[__obf_a4042898d22f95f6]])
	}

	return ID(__obf_413d3a76ff8ab79e), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_5f66c6dacd23d45f ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_5f66c6dacd23d45f.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_413d3a76ff8ab79e string) (ID, error) {
	__obf_9b18a3f8d8063793, __obf_57cef2f010b4ce9a := base64.StdEncoding.DecodeString(__obf_413d3a76ff8ab79e)
	if __obf_57cef2f010b4ce9a != nil {
		return -1, __obf_57cef2f010b4ce9a
	}
	return ParseBytes(__obf_9b18a3f8d8063793)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_5f66c6dacd23d45f ID) Bytes() []byte {
	return []byte(__obf_5f66c6dacd23d45f.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_413d3a76ff8ab79e []byte) (ID, error) {
	__obf_a4042898d22f95f6, __obf_57cef2f010b4ce9a := strconv.ParseInt(string(__obf_413d3a76ff8ab79e), 10, 64)
	return ID(__obf_a4042898d22f95f6), __obf_57cef2f010b4ce9a
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_5f66c6dacd23d45f ID) IntBytes() [8]byte {
	var __obf_9b18a3f8d8063793 [8]byte
	binary.BigEndian.PutUint64(__obf_9b18a3f8d8063793[:], uint64(__obf_5f66c6dacd23d45f))
	return __obf_9b18a3f8d8063793
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_413d3a76ff8ab79e [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_413d3a76ff8ab79e[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_5f66c6dacd23d45f ID) Time() int64 {
	return (int64(__obf_5f66c6dacd23d45f) >> __obf_444c3eed855a9bba) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_5f66c6dacd23d45f ID) Node() int64 {
	return int64(__obf_5f66c6dacd23d45f) & __obf_bbc1c63b768b03cc >> __obf_e19cb3423c33812f
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_5f66c6dacd23d45f ID) Step() int64 {
	return int64(__obf_5f66c6dacd23d45f) & __obf_4bb5a3804ded0828
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_5f66c6dacd23d45f ID) MarshalJSON() ([]byte, error) {
	__obf_2e4f0199afd4749f := make([]byte, 0, 22)
	__obf_2e4f0199afd4749f = append(__obf_2e4f0199afd4749f, '"')
	__obf_2e4f0199afd4749f = strconv.AppendInt(__obf_2e4f0199afd4749f, int64(__obf_5f66c6dacd23d45f), 10)
	__obf_2e4f0199afd4749f = append(__obf_2e4f0199afd4749f, '"')
	return __obf_2e4f0199afd4749f, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_5f66c6dacd23d45f *ID) UnmarshalJSON(__obf_9b18a3f8d8063793 []byte) error {
	if len(__obf_9b18a3f8d8063793) < 3 || __obf_9b18a3f8d8063793[0] != '"' || __obf_9b18a3f8d8063793[len(__obf_9b18a3f8d8063793)-1] != '"' {
		return JSONSyntaxError{__obf_9b18a3f8d8063793}
	}
	__obf_a4042898d22f95f6, __obf_57cef2f010b4ce9a := strconv.ParseInt(string(__obf_9b18a3f8d8063793[1:len(__obf_9b18a3f8d8063793)-1]), 10, 64)
	if __obf_57cef2f010b4ce9a != nil {
		return __obf_57cef2f010b4ce9a
	}

	*__obf_5f66c6dacd23d45f = ID(__obf_a4042898d22f95f6)
	return nil
}
