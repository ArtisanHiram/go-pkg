package __obf_34ce7ee87a5aa6b7

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
	__obf_ca950ed529feb0aa sync. // DEPRECATED: the below four variables will be removed in a future release.
				Mutex
	__obf_36317282a4c93d63 int64 = -1 ^ (-1 << NodeBits)
	__obf_de05b40f137f7d76       = __obf_36317282a4c93d63 <<
		StepBits
	__obf_ad800cdd0090678f int64 = -1 ^ (-1 << StepBits)
	__obf_a9abb3327ca8ff6a       = NodeBits + StepBits
	__obf_a6ed5a5633ab91ee       = StepBits
)

const __obf_c4afde87e000c829 = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_e6f0a8129d87632c [256]byte

const __obf_a8b98f8545ad9f19 = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_134274a634842e71 [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_faf682486813ab45 []byte }

func (__obf_b1ad146e6310b209 JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_b1ad146e6310b209.

		// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
		__obf_faf682486813ab45))
}

var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func init() {

	for __obf_5c1ed62ff2bd2f1a := 0; __obf_5c1ed62ff2bd2f1a < len(__obf_134274a634842e71); __obf_5c1ed62ff2bd2f1a++ {
		__obf_134274a634842e71[__obf_5c1ed62ff2bd2f1a] = 0xFF
	}

	for __obf_5c1ed62ff2bd2f1a := 0; __obf_5c1ed62ff2bd2f1a < len(__obf_a8b98f8545ad9f19); __obf_5c1ed62ff2bd2f1a++ {
		__obf_134274a634842e71[__obf_a8b98f8545ad9f19[__obf_5c1ed62ff2bd2f1a]] = byte(__obf_5c1ed62ff2bd2f1a)
	}

	for __obf_5c1ed62ff2bd2f1a := 0; __obf_5c1ed62ff2bd2f1a < len(__obf_e6f0a8129d87632c); __obf_5c1ed62ff2bd2f1a++ {
		__obf_e6f0a8129d87632c[__obf_5c1ed62ff2bd2f1a] = 0xFF
	}

	for __obf_5c1ed62ff2bd2f1a := 0; __obf_5c1ed62ff2bd2f1a < len(__obf_c4afde87e000c829); __obf_5c1ed62ff2bd2f1a++ {
		__obf_e6f0a8129d87632c[__obf_c4afde87e000c829[__obf_5c1ed62ff2bd2f1a]] = byte(__obf_5c1ed62ff2bd2f1a)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_ca950ed529feb0aa sync.Mutex
	__obf_ff33511dd60894e7 time.Time
	time                   int64
	__obf_e90657cd6335ee96 int64
	__obf_75199e701b498833 int64
	__obf_36317282a4c93d63 int64
	__obf_de05b40f137f7d76 int64
	__obf_ad800cdd0090678f int64
	__obf_a9abb3327ca8ff6a uint8
	__obf_a6ed5a5633ab91ee uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_e90657cd6335ee96 int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	__obf_ca950ed529feb0aa.
		// re-calc in case custom NodeBits or StepBits were set
		// DEPRECATED: the below block will be removed in a future release.
		Lock()
	__obf_36317282a4c93d63 = -1 ^ (-1 << NodeBits)
	__obf_de05b40f137f7d76 = __obf_36317282a4c93d63 << StepBits
	__obf_ad800cdd0090678f = -1 ^ (-1 << StepBits)
	__obf_a9abb3327ca8ff6a = NodeBits + StepBits
	__obf_a6ed5a5633ab91ee = StepBits
	__obf_ca950ed529feb0aa.
		Unlock()
	__obf_0d08165857356423 := Node{}
	__obf_0d08165857356423.__obf_e90657cd6335ee96 = __obf_e90657cd6335ee96
	__obf_0d08165857356423.__obf_36317282a4c93d63 = -1 ^ (-1 << NodeBits)
	__obf_0d08165857356423.__obf_de05b40f137f7d76 = __obf_0d08165857356423.__obf_36317282a4c93d63 << StepBits
	__obf_0d08165857356423.__obf_ad800cdd0090678f = -1 ^ (-1 << StepBits)
	__obf_0d08165857356423.__obf_a9abb3327ca8ff6a = NodeBits + StepBits
	__obf_0d08165857356423.__obf_a6ed5a5633ab91ee = StepBits

	if __obf_0d08165857356423.__obf_e90657cd6335ee96 < 0 || __obf_0d08165857356423.__obf_e90657cd6335ee96 > __obf_0d08165857356423.__obf_36317282a4c93d63 {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_0d08165857356423.__obf_36317282a4c93d63, 10))
	}

	var __obf_d8ba6c1a13ba402f = time.Now()
	__obf_0d08165857356423.
		// add time.Duration to curTime to make sure we use the monotonic clock if available
		__obf_ff33511dd60894e7 = __obf_d8ba6c1a13ba402f.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_d8ba6c1a13ba402f))

	return &__obf_0d08165857356423, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_0d08165857356423 *Node) Generate() ID {
	__obf_0d08165857356423.__obf_ca950ed529feb0aa.
		Lock()
	defer __obf_0d08165857356423.__obf_ca950ed529feb0aa.Unlock()
	__obf_58a59a6cbcc05a7a := time.Since(__obf_0d08165857356423.__obf_ff33511dd60894e7).Milliseconds()

	if __obf_58a59a6cbcc05a7a == __obf_0d08165857356423.time {
		__obf_0d08165857356423.__obf_75199e701b498833 = (__obf_0d08165857356423.__obf_75199e701b498833 + 1) & __obf_0d08165857356423.__obf_ad800cdd0090678f

		if __obf_0d08165857356423.__obf_75199e701b498833 == 0 {
			for __obf_58a59a6cbcc05a7a <= __obf_0d08165857356423.time {
				__obf_58a59a6cbcc05a7a = time.Since(__obf_0d08165857356423.__obf_ff33511dd60894e7).Milliseconds()
			}
		}
	} else {
		__obf_0d08165857356423.__obf_75199e701b498833 = 0
	}
	__obf_0d08165857356423.
		time = __obf_58a59a6cbcc05a7a
	__obf_360e62cded5e2997 := ID((__obf_58a59a6cbcc05a7a)<<__obf_0d08165857356423.__obf_a9abb3327ca8ff6a |
		(__obf_0d08165857356423.__obf_e90657cd6335ee96 << __obf_0d08165857356423.__obf_a6ed5a5633ab91ee) |
		(__obf_0d08165857356423.__obf_75199e701b498833),
	)

	return __obf_360e62cded5e2997
}

// Int64 returns an int64 of the snowflake ID
func (__obf_03dfdc1b988e232e ID) Int64() int64 {
	return int64(__obf_03dfdc1b988e232e)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_f60c7174b4657d83 int64) ID {
	return ID(__obf_f60c7174b4657d83)
}

// String returns a string of the snowflake ID
func (__obf_03dfdc1b988e232e ID) String() string {
	return strconv.FormatInt(int64(__obf_03dfdc1b988e232e), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_f60c7174b4657d83 string) (ID, error) {
	__obf_5c1ed62ff2bd2f1a, __obf_0a414fcc6889bc71 := strconv.ParseInt(__obf_f60c7174b4657d83, 10, 64)
	return ID(__obf_5c1ed62ff2bd2f1a), __obf_0a414fcc6889bc71

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_03dfdc1b988e232e ID) Base2() string {
	return strconv.FormatInt(int64(__obf_03dfdc1b988e232e), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_f60c7174b4657d83 string) (ID, error) {
	__obf_5c1ed62ff2bd2f1a, __obf_0a414fcc6889bc71 := strconv.ParseInt(__obf_f60c7174b4657d83, 2, 64)
	return ID(__obf_5c1ed62ff2bd2f1a), __obf_0a414fcc6889bc71
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_03dfdc1b988e232e ID) Base32() string {

	if __obf_03dfdc1b988e232e < 32 {
		return string(__obf_c4afde87e000c829[__obf_03dfdc1b988e232e])
	}
	__obf_09d5a686a504734f := make([]byte, 0, 12)
	for __obf_03dfdc1b988e232e >= 32 {
		__obf_09d5a686a504734f = append(__obf_09d5a686a504734f, __obf_c4afde87e000c829[__obf_03dfdc1b988e232e%32])
		__obf_03dfdc1b988e232e /= 32
	}
	__obf_09d5a686a504734f = append(__obf_09d5a686a504734f, __obf_c4afde87e000c829[__obf_03dfdc1b988e232e])

	for __obf_50f57aac3b0e8f3e, __obf_bd0f90f24f33616c := 0, len(__obf_09d5a686a504734f)-1; __obf_50f57aac3b0e8f3e < __obf_bd0f90f24f33616c; __obf_50f57aac3b0e8f3e, __obf_bd0f90f24f33616c = __obf_50f57aac3b0e8f3e+1, __obf_bd0f90f24f33616c-1 {
		__obf_09d5a686a504734f[__obf_50f57aac3b0e8f3e], __obf_09d5a686a504734f[__obf_bd0f90f24f33616c] = __obf_09d5a686a504734f[__obf_bd0f90f24f33616c], __obf_09d5a686a504734f[__obf_50f57aac3b0e8f3e]
	}

	return string(__obf_09d5a686a504734f)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_09d5a686a504734f []byte) (ID, error) {

	var __obf_f60c7174b4657d83 int64

	for __obf_5c1ed62ff2bd2f1a := range __obf_09d5a686a504734f {
		if __obf_e6f0a8129d87632c[__obf_09d5a686a504734f[__obf_5c1ed62ff2bd2f1a]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_f60c7174b4657d83 = __obf_f60c7174b4657d83*32 + int64(__obf_e6f0a8129d87632c[__obf_09d5a686a504734f[__obf_5c1ed62ff2bd2f1a]])
	}

	return ID(__obf_f60c7174b4657d83), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_03dfdc1b988e232e ID) Base36() string {
	return strconv.FormatInt(int64(__obf_03dfdc1b988e232e), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_f60c7174b4657d83 string) (ID, error) {
	__obf_5c1ed62ff2bd2f1a, __obf_0a414fcc6889bc71 := strconv.ParseInt(__obf_f60c7174b4657d83, 36, 64)
	return ID(__obf_5c1ed62ff2bd2f1a), __obf_0a414fcc6889bc71
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_03dfdc1b988e232e ID) Base58() string {

	if __obf_03dfdc1b988e232e < 58 {
		return string(__obf_a8b98f8545ad9f19[__obf_03dfdc1b988e232e])
	}
	__obf_09d5a686a504734f := make([]byte, 0, 11)
	for __obf_03dfdc1b988e232e >= 58 {
		__obf_09d5a686a504734f = append(__obf_09d5a686a504734f, __obf_a8b98f8545ad9f19[__obf_03dfdc1b988e232e%58])
		__obf_03dfdc1b988e232e /= 58
	}
	__obf_09d5a686a504734f = append(__obf_09d5a686a504734f, __obf_a8b98f8545ad9f19[__obf_03dfdc1b988e232e])

	for __obf_50f57aac3b0e8f3e, __obf_bd0f90f24f33616c := 0, len(__obf_09d5a686a504734f)-1; __obf_50f57aac3b0e8f3e < __obf_bd0f90f24f33616c; __obf_50f57aac3b0e8f3e, __obf_bd0f90f24f33616c = __obf_50f57aac3b0e8f3e+1, __obf_bd0f90f24f33616c-1 {
		__obf_09d5a686a504734f[__obf_50f57aac3b0e8f3e], __obf_09d5a686a504734f[__obf_bd0f90f24f33616c] = __obf_09d5a686a504734f[__obf_bd0f90f24f33616c], __obf_09d5a686a504734f[__obf_50f57aac3b0e8f3e]
	}

	return string(__obf_09d5a686a504734f)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_09d5a686a504734f []byte) (ID, error) {

	var __obf_f60c7174b4657d83 int64

	for __obf_5c1ed62ff2bd2f1a := range __obf_09d5a686a504734f {
		if __obf_134274a634842e71[__obf_09d5a686a504734f[__obf_5c1ed62ff2bd2f1a]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_f60c7174b4657d83 = __obf_f60c7174b4657d83*58 + int64(__obf_134274a634842e71[__obf_09d5a686a504734f[__obf_5c1ed62ff2bd2f1a]])
	}

	return ID(__obf_f60c7174b4657d83), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_03dfdc1b988e232e ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_03dfdc1b988e232e.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_f60c7174b4657d83 string) (ID, error) {
	__obf_09d5a686a504734f, __obf_0a414fcc6889bc71 := base64.StdEncoding.DecodeString(__obf_f60c7174b4657d83)
	if __obf_0a414fcc6889bc71 != nil {
		return -1, __obf_0a414fcc6889bc71
	}
	return ParseBytes(__obf_09d5a686a504734f)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_03dfdc1b988e232e ID) Bytes() []byte {
	return []byte(__obf_03dfdc1b988e232e.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_f60c7174b4657d83 []byte) (ID, error) {
	__obf_5c1ed62ff2bd2f1a, __obf_0a414fcc6889bc71 := strconv.ParseInt(string(__obf_f60c7174b4657d83), 10, 64)
	return ID(__obf_5c1ed62ff2bd2f1a), __obf_0a414fcc6889bc71
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_03dfdc1b988e232e ID) IntBytes() [8]byte {
	var __obf_09d5a686a504734f [8]byte
	binary.BigEndian.PutUint64(__obf_09d5a686a504734f[:], uint64(__obf_03dfdc1b988e232e))
	return __obf_09d5a686a504734f
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_f60c7174b4657d83 [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_f60c7174b4657d83[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_03dfdc1b988e232e ID) Time() int64 {
	return (int64(__obf_03dfdc1b988e232e) >> __obf_a9abb3327ca8ff6a) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_03dfdc1b988e232e ID) Node() int64 {
	return int64(__obf_03dfdc1b988e232e) & __obf_de05b40f137f7d76 >> __obf_a6ed5a5633ab91ee
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_03dfdc1b988e232e ID) Step() int64 {
	return int64(__obf_03dfdc1b988e232e) & __obf_ad800cdd0090678f
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_03dfdc1b988e232e ID) MarshalJSON() ([]byte, error) {
	__obf_f6f694676f89a8b8 := make([]byte, 0, 22)
	__obf_f6f694676f89a8b8 = append(__obf_f6f694676f89a8b8, '"')
	__obf_f6f694676f89a8b8 = strconv.AppendInt(__obf_f6f694676f89a8b8, int64(__obf_03dfdc1b988e232e), 10)
	__obf_f6f694676f89a8b8 = append(__obf_f6f694676f89a8b8, '"')
	return __obf_f6f694676f89a8b8, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_03dfdc1b988e232e *ID) UnmarshalJSON(__obf_09d5a686a504734f []byte) error {
	if len(__obf_09d5a686a504734f) < 3 || __obf_09d5a686a504734f[0] != '"' || __obf_09d5a686a504734f[len(__obf_09d5a686a504734f)-1] != '"' {
		return JSONSyntaxError{__obf_09d5a686a504734f}
	}
	__obf_5c1ed62ff2bd2f1a, __obf_0a414fcc6889bc71 := strconv.ParseInt(string(__obf_09d5a686a504734f[1:len(__obf_09d5a686a504734f)-1]), 10, 64)
	if __obf_0a414fcc6889bc71 != nil {
		return __obf_0a414fcc6889bc71
	}

	*__obf_03dfdc1b988e232e = ID(__obf_5c1ed62ff2bd2f1a)
	return nil
}
