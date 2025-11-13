package __obf_417508f5ae215d0a

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
	__obf_d87314b4d756b35e sync.Mutex
	__obf_88880bdccf057e46 int64 = -1 ^ (-1 << NodeBits)
	__obf_4dcf51b16f65205c       = __obf_88880bdccf057e46 << StepBits
	__obf_212be91272526480 int64 = -1 ^ (-1 << StepBits)
	__obf_95dbd4ae72aeb1ce       = NodeBits + StepBits
	__obf_6f56b015fc9958a8       = StepBits
)

const __obf_6d49a67a3e92cc8b = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_ac613c1da05ec054 [256]byte

const __obf_ae414d8fba2d2484 = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_3ef6bd8cce3e5ada [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_4fb336162d4747b4 []byte }

func (__obf_5b9e1f6fc6e21c16 JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_5b9e1f6fc6e21c16.__obf_4fb336162d4747b4))
}

// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func __obf_474da4075ef60947() {

	for __obf_6e993b9e217ddd26 := 0; __obf_6e993b9e217ddd26 < len(__obf_3ef6bd8cce3e5ada); __obf_6e993b9e217ddd26++ {
		__obf_3ef6bd8cce3e5ada[__obf_6e993b9e217ddd26] = 0xFF
	}

	for __obf_6e993b9e217ddd26 := 0; __obf_6e993b9e217ddd26 < len(__obf_ae414d8fba2d2484); __obf_6e993b9e217ddd26++ {
		__obf_3ef6bd8cce3e5ada[__obf_ae414d8fba2d2484[__obf_6e993b9e217ddd26]] = byte(__obf_6e993b9e217ddd26)
	}

	for __obf_6e993b9e217ddd26 := 0; __obf_6e993b9e217ddd26 < len(__obf_ac613c1da05ec054); __obf_6e993b9e217ddd26++ {
		__obf_ac613c1da05ec054[__obf_6e993b9e217ddd26] = 0xFF
	}

	for __obf_6e993b9e217ddd26 := 0; __obf_6e993b9e217ddd26 < len(__obf_6d49a67a3e92cc8b); __obf_6e993b9e217ddd26++ {
		__obf_ac613c1da05ec054[__obf_6d49a67a3e92cc8b[__obf_6e993b9e217ddd26]] = byte(__obf_6e993b9e217ddd26)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_d87314b4d756b35e sync.Mutex
	__obf_53a0b8ccab269061 time.Time
	time                   int64
	__obf_a5fd510dab391272 int64
	__obf_810107a662769211 int64

	__obf_88880bdccf057e46 int64
	__obf_4dcf51b16f65205c int64
	__obf_212be91272526480 int64
	__obf_95dbd4ae72aeb1ce uint8
	__obf_6f56b015fc9958a8 uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_a5fd510dab391272 int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	// re-calc in case custom NodeBits or StepBits were set
	// DEPRECATED: the below block will be removed in a future release.
	__obf_d87314b4d756b35e.Lock()
	__obf_88880bdccf057e46 = -1 ^ (-1 << NodeBits)
	__obf_4dcf51b16f65205c = __obf_88880bdccf057e46 << StepBits
	__obf_212be91272526480 = -1 ^ (-1 << StepBits)
	__obf_95dbd4ae72aeb1ce = NodeBits + StepBits
	__obf_6f56b015fc9958a8 = StepBits
	__obf_d87314b4d756b35e.Unlock()

	__obf_097b9edd67ed12ac := Node{}
	__obf_097b9edd67ed12ac.__obf_a5fd510dab391272 = __obf_a5fd510dab391272
	__obf_097b9edd67ed12ac.__obf_88880bdccf057e46 = -1 ^ (-1 << NodeBits)
	__obf_097b9edd67ed12ac.__obf_4dcf51b16f65205c = __obf_097b9edd67ed12ac.__obf_88880bdccf057e46 << StepBits
	__obf_097b9edd67ed12ac.__obf_212be91272526480 = -1 ^ (-1 << StepBits)
	__obf_097b9edd67ed12ac.__obf_95dbd4ae72aeb1ce = NodeBits + StepBits
	__obf_097b9edd67ed12ac.__obf_6f56b015fc9958a8 = StepBits

	if __obf_097b9edd67ed12ac.__obf_a5fd510dab391272 < 0 || __obf_097b9edd67ed12ac.__obf_a5fd510dab391272 > __obf_097b9edd67ed12ac.__obf_88880bdccf057e46 {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_097b9edd67ed12ac.__obf_88880bdccf057e46, 10))
	}

	var __obf_cceb6d51c6cedb86 = time.Now()
	// add time.Duration to curTime to make sure we use the monotonic clock if available
	__obf_097b9edd67ed12ac.__obf_53a0b8ccab269061 = __obf_cceb6d51c6cedb86.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_cceb6d51c6cedb86))

	return &__obf_097b9edd67ed12ac, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_097b9edd67ed12ac *Node) Generate() ID {

	__obf_097b9edd67ed12ac.__obf_d87314b4d756b35e.Lock()
	defer __obf_097b9edd67ed12ac.__obf_d87314b4d756b35e.Unlock()

	__obf_cdfc766fcfeefd70 := time.Since(__obf_097b9edd67ed12ac.__obf_53a0b8ccab269061).Milliseconds()

	if __obf_cdfc766fcfeefd70 == __obf_097b9edd67ed12ac.time {
		__obf_097b9edd67ed12ac.__obf_810107a662769211 = (__obf_097b9edd67ed12ac.__obf_810107a662769211 + 1) & __obf_097b9edd67ed12ac.__obf_212be91272526480

		if __obf_097b9edd67ed12ac.__obf_810107a662769211 == 0 {
			for __obf_cdfc766fcfeefd70 <= __obf_097b9edd67ed12ac.time {
				__obf_cdfc766fcfeefd70 = time.Since(__obf_097b9edd67ed12ac.__obf_53a0b8ccab269061).Milliseconds()
			}
		}
	} else {
		__obf_097b9edd67ed12ac.__obf_810107a662769211 = 0
	}

	__obf_097b9edd67ed12ac.time = __obf_cdfc766fcfeefd70

	__obf_93999fd0eeb6db63 := ID((__obf_cdfc766fcfeefd70)<<__obf_097b9edd67ed12ac.__obf_95dbd4ae72aeb1ce |
		(__obf_097b9edd67ed12ac.__obf_a5fd510dab391272 << __obf_097b9edd67ed12ac.__obf_6f56b015fc9958a8) |
		(__obf_097b9edd67ed12ac.__obf_810107a662769211),
	)

	return __obf_93999fd0eeb6db63
}

// Int64 returns an int64 of the snowflake ID
func (__obf_bf59baeaaf4309a2 ID) Int64() int64 {
	return int64(__obf_bf59baeaaf4309a2)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_48218c2a1a645d91 int64) ID {
	return ID(__obf_48218c2a1a645d91)
}

// String returns a string of the snowflake ID
func (__obf_bf59baeaaf4309a2 ID) String() string {
	return strconv.FormatInt(int64(__obf_bf59baeaaf4309a2), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_48218c2a1a645d91 string) (ID, error) {
	__obf_6e993b9e217ddd26, __obf_b99f62eaceb52d40 := strconv.ParseInt(__obf_48218c2a1a645d91, 10, 64)
	return ID(__obf_6e993b9e217ddd26), __obf_b99f62eaceb52d40

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_bf59baeaaf4309a2 ID) Base2() string {
	return strconv.FormatInt(int64(__obf_bf59baeaaf4309a2), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_48218c2a1a645d91 string) (ID, error) {
	__obf_6e993b9e217ddd26, __obf_b99f62eaceb52d40 := strconv.ParseInt(__obf_48218c2a1a645d91, 2, 64)
	return ID(__obf_6e993b9e217ddd26), __obf_b99f62eaceb52d40
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_bf59baeaaf4309a2 ID) Base32() string {

	if __obf_bf59baeaaf4309a2 < 32 {
		return string(__obf_6d49a67a3e92cc8b[__obf_bf59baeaaf4309a2])
	}

	__obf_d5835f2896f288ef := make([]byte, 0, 12)
	for __obf_bf59baeaaf4309a2 >= 32 {
		__obf_d5835f2896f288ef = append(__obf_d5835f2896f288ef, __obf_6d49a67a3e92cc8b[__obf_bf59baeaaf4309a2%32])
		__obf_bf59baeaaf4309a2 /= 32
	}
	__obf_d5835f2896f288ef = append(__obf_d5835f2896f288ef, __obf_6d49a67a3e92cc8b[__obf_bf59baeaaf4309a2])

	for __obf_46a933db83592c20, __obf_0587a9ad79713bb3 := 0, len(__obf_d5835f2896f288ef)-1; __obf_46a933db83592c20 < __obf_0587a9ad79713bb3; __obf_46a933db83592c20, __obf_0587a9ad79713bb3 = __obf_46a933db83592c20+1, __obf_0587a9ad79713bb3-1 {
		__obf_d5835f2896f288ef[__obf_46a933db83592c20], __obf_d5835f2896f288ef[__obf_0587a9ad79713bb3] = __obf_d5835f2896f288ef[__obf_0587a9ad79713bb3], __obf_d5835f2896f288ef[__obf_46a933db83592c20]
	}

	return string(__obf_d5835f2896f288ef)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_d5835f2896f288ef []byte) (ID, error) {

	var __obf_48218c2a1a645d91 int64

	for __obf_6e993b9e217ddd26 := range __obf_d5835f2896f288ef {
		if __obf_ac613c1da05ec054[__obf_d5835f2896f288ef[__obf_6e993b9e217ddd26]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_48218c2a1a645d91 = __obf_48218c2a1a645d91*32 + int64(__obf_ac613c1da05ec054[__obf_d5835f2896f288ef[__obf_6e993b9e217ddd26]])
	}

	return ID(__obf_48218c2a1a645d91), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_bf59baeaaf4309a2 ID) Base36() string {
	return strconv.FormatInt(int64(__obf_bf59baeaaf4309a2), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_48218c2a1a645d91 string) (ID, error) {
	__obf_6e993b9e217ddd26, __obf_b99f62eaceb52d40 := strconv.ParseInt(__obf_48218c2a1a645d91, 36, 64)
	return ID(__obf_6e993b9e217ddd26), __obf_b99f62eaceb52d40
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_bf59baeaaf4309a2 ID) Base58() string {

	if __obf_bf59baeaaf4309a2 < 58 {
		return string(__obf_ae414d8fba2d2484[__obf_bf59baeaaf4309a2])
	}

	__obf_d5835f2896f288ef := make([]byte, 0, 11)
	for __obf_bf59baeaaf4309a2 >= 58 {
		__obf_d5835f2896f288ef = append(__obf_d5835f2896f288ef, __obf_ae414d8fba2d2484[__obf_bf59baeaaf4309a2%58])
		__obf_bf59baeaaf4309a2 /= 58
	}
	__obf_d5835f2896f288ef = append(__obf_d5835f2896f288ef, __obf_ae414d8fba2d2484[__obf_bf59baeaaf4309a2])

	for __obf_46a933db83592c20, __obf_0587a9ad79713bb3 := 0, len(__obf_d5835f2896f288ef)-1; __obf_46a933db83592c20 < __obf_0587a9ad79713bb3; __obf_46a933db83592c20, __obf_0587a9ad79713bb3 = __obf_46a933db83592c20+1, __obf_0587a9ad79713bb3-1 {
		__obf_d5835f2896f288ef[__obf_46a933db83592c20], __obf_d5835f2896f288ef[__obf_0587a9ad79713bb3] = __obf_d5835f2896f288ef[__obf_0587a9ad79713bb3], __obf_d5835f2896f288ef[__obf_46a933db83592c20]
	}

	return string(__obf_d5835f2896f288ef)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_d5835f2896f288ef []byte) (ID, error) {

	var __obf_48218c2a1a645d91 int64

	for __obf_6e993b9e217ddd26 := range __obf_d5835f2896f288ef {
		if __obf_3ef6bd8cce3e5ada[__obf_d5835f2896f288ef[__obf_6e993b9e217ddd26]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_48218c2a1a645d91 = __obf_48218c2a1a645d91*58 + int64(__obf_3ef6bd8cce3e5ada[__obf_d5835f2896f288ef[__obf_6e993b9e217ddd26]])
	}

	return ID(__obf_48218c2a1a645d91), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_bf59baeaaf4309a2 ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_bf59baeaaf4309a2.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_48218c2a1a645d91 string) (ID, error) {
	__obf_d5835f2896f288ef, __obf_b99f62eaceb52d40 := base64.StdEncoding.DecodeString(__obf_48218c2a1a645d91)
	if __obf_b99f62eaceb52d40 != nil {
		return -1, __obf_b99f62eaceb52d40
	}
	return ParseBytes(__obf_d5835f2896f288ef)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_bf59baeaaf4309a2 ID) Bytes() []byte {
	return []byte(__obf_bf59baeaaf4309a2.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_48218c2a1a645d91 []byte) (ID, error) {
	__obf_6e993b9e217ddd26, __obf_b99f62eaceb52d40 := strconv.ParseInt(string(__obf_48218c2a1a645d91), 10, 64)
	return ID(__obf_6e993b9e217ddd26), __obf_b99f62eaceb52d40
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_bf59baeaaf4309a2 ID) IntBytes() [8]byte {
	var __obf_d5835f2896f288ef [8]byte
	binary.BigEndian.PutUint64(__obf_d5835f2896f288ef[:], uint64(__obf_bf59baeaaf4309a2))
	return __obf_d5835f2896f288ef
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_48218c2a1a645d91 [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_48218c2a1a645d91[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_bf59baeaaf4309a2 ID) Time() int64 {
	return (int64(__obf_bf59baeaaf4309a2) >> __obf_95dbd4ae72aeb1ce) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_bf59baeaaf4309a2 ID) Node() int64 {
	return int64(__obf_bf59baeaaf4309a2) & __obf_4dcf51b16f65205c >> __obf_6f56b015fc9958a8
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_bf59baeaaf4309a2 ID) Step() int64 {
	return int64(__obf_bf59baeaaf4309a2) & __obf_212be91272526480
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_bf59baeaaf4309a2 ID) MarshalJSON() ([]byte, error) {
	__obf_aaea6784b2f0c22f := make([]byte, 0, 22)
	__obf_aaea6784b2f0c22f = append(__obf_aaea6784b2f0c22f, '"')
	__obf_aaea6784b2f0c22f = strconv.AppendInt(__obf_aaea6784b2f0c22f, int64(__obf_bf59baeaaf4309a2), 10)
	__obf_aaea6784b2f0c22f = append(__obf_aaea6784b2f0c22f, '"')
	return __obf_aaea6784b2f0c22f, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_bf59baeaaf4309a2 *ID) UnmarshalJSON(__obf_d5835f2896f288ef []byte) error {
	if len(__obf_d5835f2896f288ef) < 3 || __obf_d5835f2896f288ef[0] != '"' || __obf_d5835f2896f288ef[len(__obf_d5835f2896f288ef)-1] != '"' {
		return JSONSyntaxError{__obf_d5835f2896f288ef}
	}

	__obf_6e993b9e217ddd26, __obf_b99f62eaceb52d40 := strconv.ParseInt(string(__obf_d5835f2896f288ef[1:len(__obf_d5835f2896f288ef)-1]), 10, 64)
	if __obf_b99f62eaceb52d40 != nil {
		return __obf_b99f62eaceb52d40
	}

	*__obf_bf59baeaaf4309a2 = ID(__obf_6e993b9e217ddd26)
	return nil
}
