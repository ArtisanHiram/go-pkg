package __obf_4d7a51f8e2f0494a

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
	__obf_6d0dc3fbe6f4289e sync. // DEPRECATED: the below four variables will be removed in a future release.
				Mutex
	__obf_4017ad9c558adbee int64 = -1 ^ (-1 << NodeBits)
	__obf_b87a81d04e8ecdc5       = __obf_4017ad9c558adbee <<
		StepBits
	__obf_118ef6870c08dca5 int64 = -1 ^ (-1 << StepBits)
	__obf_43facb4bfd5cad5b       = NodeBits + StepBits
	__obf_6fde542eef241ce8       = StepBits
)

const __obf_c0cae8079f5416fe = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_ad3dbfde58f4baf0 [256]byte

const __obf_e93fc1cf96e54053 = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_487b5deba62869c7 [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_a453391550f5d71d []byte }

func (__obf_ccf2022677d69302 JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_ccf2022677d69302.

		// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
		__obf_a453391550f5d71d))
}

var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func init() {

	for __obf_d43488f90f0960a8 := 0; __obf_d43488f90f0960a8 < len(__obf_487b5deba62869c7); __obf_d43488f90f0960a8++ {
		__obf_487b5deba62869c7[__obf_d43488f90f0960a8] = 0xFF
	}

	for __obf_d43488f90f0960a8 := 0; __obf_d43488f90f0960a8 < len(__obf_e93fc1cf96e54053); __obf_d43488f90f0960a8++ {
		__obf_487b5deba62869c7[__obf_e93fc1cf96e54053[__obf_d43488f90f0960a8]] = byte(__obf_d43488f90f0960a8)
	}

	for __obf_d43488f90f0960a8 := 0; __obf_d43488f90f0960a8 < len(__obf_ad3dbfde58f4baf0); __obf_d43488f90f0960a8++ {
		__obf_ad3dbfde58f4baf0[__obf_d43488f90f0960a8] = 0xFF
	}

	for __obf_d43488f90f0960a8 := 0; __obf_d43488f90f0960a8 < len(__obf_c0cae8079f5416fe); __obf_d43488f90f0960a8++ {
		__obf_ad3dbfde58f4baf0[__obf_c0cae8079f5416fe[__obf_d43488f90f0960a8]] = byte(__obf_d43488f90f0960a8)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_6d0dc3fbe6f4289e sync.Mutex
	__obf_849a3f2c46cc8ab7 time.Time
	time                   int64
	__obf_55e3af91627ba314 int64
	__obf_c0107f4cae83b90d int64
	__obf_4017ad9c558adbee int64
	__obf_b87a81d04e8ecdc5 int64
	__obf_118ef6870c08dca5 int64
	__obf_43facb4bfd5cad5b uint8
	__obf_6fde542eef241ce8 uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_55e3af91627ba314 int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	__obf_6d0dc3fbe6f4289e.
		// re-calc in case custom NodeBits or StepBits were set
		// DEPRECATED: the below block will be removed in a future release.
		Lock()
	__obf_4017ad9c558adbee = -1 ^ (-1 << NodeBits)
	__obf_b87a81d04e8ecdc5 = __obf_4017ad9c558adbee << StepBits
	__obf_118ef6870c08dca5 = -1 ^ (-1 << StepBits)
	__obf_43facb4bfd5cad5b = NodeBits + StepBits
	__obf_6fde542eef241ce8 = StepBits
	__obf_6d0dc3fbe6f4289e.
		Unlock()
	__obf_da28677edd5fd8f1 := Node{}
	__obf_da28677edd5fd8f1.__obf_55e3af91627ba314 = __obf_55e3af91627ba314
	__obf_da28677edd5fd8f1.__obf_4017ad9c558adbee = -1 ^ (-1 << NodeBits)
	__obf_da28677edd5fd8f1.__obf_b87a81d04e8ecdc5 = __obf_da28677edd5fd8f1.__obf_4017ad9c558adbee << StepBits
	__obf_da28677edd5fd8f1.__obf_118ef6870c08dca5 = -1 ^ (-1 << StepBits)
	__obf_da28677edd5fd8f1.__obf_43facb4bfd5cad5b = NodeBits + StepBits
	__obf_da28677edd5fd8f1.__obf_6fde542eef241ce8 = StepBits

	if __obf_da28677edd5fd8f1.__obf_55e3af91627ba314 < 0 || __obf_da28677edd5fd8f1.__obf_55e3af91627ba314 > __obf_da28677edd5fd8f1.__obf_4017ad9c558adbee {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_da28677edd5fd8f1.__obf_4017ad9c558adbee, 10))
	}

	var __obf_a170573baa911e3e = time.Now()
	__obf_da28677edd5fd8f1.
		// add time.Duration to curTime to make sure we use the monotonic clock if available
		__obf_849a3f2c46cc8ab7 = __obf_a170573baa911e3e.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_a170573baa911e3e))

	return &__obf_da28677edd5fd8f1, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_da28677edd5fd8f1 *Node) Generate() ID {
	__obf_da28677edd5fd8f1.__obf_6d0dc3fbe6f4289e.
		Lock()
	defer __obf_da28677edd5fd8f1.__obf_6d0dc3fbe6f4289e.Unlock()
	__obf_ab7c81ad01d32d59 := time.Since(__obf_da28677edd5fd8f1.__obf_849a3f2c46cc8ab7).Milliseconds()

	if __obf_ab7c81ad01d32d59 == __obf_da28677edd5fd8f1.time {
		__obf_da28677edd5fd8f1.__obf_c0107f4cae83b90d = (__obf_da28677edd5fd8f1.__obf_c0107f4cae83b90d + 1) & __obf_da28677edd5fd8f1.__obf_118ef6870c08dca5

		if __obf_da28677edd5fd8f1.__obf_c0107f4cae83b90d == 0 {
			for __obf_ab7c81ad01d32d59 <= __obf_da28677edd5fd8f1.time {
				__obf_ab7c81ad01d32d59 = time.Since(__obf_da28677edd5fd8f1.__obf_849a3f2c46cc8ab7).Milliseconds()
			}
		}
	} else {
		__obf_da28677edd5fd8f1.__obf_c0107f4cae83b90d = 0
	}
	__obf_da28677edd5fd8f1.
		time = __obf_ab7c81ad01d32d59
	__obf_4c0fa134d0dffa75 := ID((__obf_ab7c81ad01d32d59)<<__obf_da28677edd5fd8f1.__obf_43facb4bfd5cad5b |
		(__obf_da28677edd5fd8f1.__obf_55e3af91627ba314 << __obf_da28677edd5fd8f1.__obf_6fde542eef241ce8) |
		(__obf_da28677edd5fd8f1.__obf_c0107f4cae83b90d),
	)

	return __obf_4c0fa134d0dffa75
}

// Int64 returns an int64 of the snowflake ID
func (__obf_045c7fea9195ede5 ID) Int64() int64 {
	return int64(__obf_045c7fea9195ede5)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_721283c42292359f int64) ID {
	return ID(__obf_721283c42292359f)
}

// String returns a string of the snowflake ID
func (__obf_045c7fea9195ede5 ID) String() string {
	return strconv.FormatInt(int64(__obf_045c7fea9195ede5), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_721283c42292359f string) (ID, error) {
	__obf_d43488f90f0960a8, __obf_38c5d446257f8d9a := strconv.ParseInt(__obf_721283c42292359f, 10, 64)
	return ID(__obf_d43488f90f0960a8), __obf_38c5d446257f8d9a

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_045c7fea9195ede5 ID) Base2() string {
	return strconv.FormatInt(int64(__obf_045c7fea9195ede5), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_721283c42292359f string) (ID, error) {
	__obf_d43488f90f0960a8, __obf_38c5d446257f8d9a := strconv.ParseInt(__obf_721283c42292359f, 2, 64)
	return ID(__obf_d43488f90f0960a8), __obf_38c5d446257f8d9a
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_045c7fea9195ede5 ID) Base32() string {

	if __obf_045c7fea9195ede5 < 32 {
		return string(__obf_c0cae8079f5416fe[__obf_045c7fea9195ede5])
	}
	__obf_25d0a77df4064330 := make([]byte, 0, 12)
	for __obf_045c7fea9195ede5 >= 32 {
		__obf_25d0a77df4064330 = append(__obf_25d0a77df4064330, __obf_c0cae8079f5416fe[__obf_045c7fea9195ede5%32])
		__obf_045c7fea9195ede5 /= 32
	}
	__obf_25d0a77df4064330 = append(__obf_25d0a77df4064330, __obf_c0cae8079f5416fe[__obf_045c7fea9195ede5])

	for __obf_2a103e2ab03dbd0a, __obf_dffb606f75e8b13c := 0, len(__obf_25d0a77df4064330)-1; __obf_2a103e2ab03dbd0a < __obf_dffb606f75e8b13c; __obf_2a103e2ab03dbd0a, __obf_dffb606f75e8b13c = __obf_2a103e2ab03dbd0a+1, __obf_dffb606f75e8b13c-1 {
		__obf_25d0a77df4064330[__obf_2a103e2ab03dbd0a], __obf_25d0a77df4064330[__obf_dffb606f75e8b13c] = __obf_25d0a77df4064330[__obf_dffb606f75e8b13c], __obf_25d0a77df4064330[__obf_2a103e2ab03dbd0a]
	}

	return string(__obf_25d0a77df4064330)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_25d0a77df4064330 []byte) (ID, error) {

	var __obf_721283c42292359f int64

	for __obf_d43488f90f0960a8 := range __obf_25d0a77df4064330 {
		if __obf_ad3dbfde58f4baf0[__obf_25d0a77df4064330[__obf_d43488f90f0960a8]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_721283c42292359f = __obf_721283c42292359f*32 + int64(__obf_ad3dbfde58f4baf0[__obf_25d0a77df4064330[__obf_d43488f90f0960a8]])
	}

	return ID(__obf_721283c42292359f), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_045c7fea9195ede5 ID) Base36() string {
	return strconv.FormatInt(int64(__obf_045c7fea9195ede5), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_721283c42292359f string) (ID, error) {
	__obf_d43488f90f0960a8, __obf_38c5d446257f8d9a := strconv.ParseInt(__obf_721283c42292359f, 36, 64)
	return ID(__obf_d43488f90f0960a8), __obf_38c5d446257f8d9a
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_045c7fea9195ede5 ID) Base58() string {

	if __obf_045c7fea9195ede5 < 58 {
		return string(__obf_e93fc1cf96e54053[__obf_045c7fea9195ede5])
	}
	__obf_25d0a77df4064330 := make([]byte, 0, 11)
	for __obf_045c7fea9195ede5 >= 58 {
		__obf_25d0a77df4064330 = append(__obf_25d0a77df4064330, __obf_e93fc1cf96e54053[__obf_045c7fea9195ede5%58])
		__obf_045c7fea9195ede5 /= 58
	}
	__obf_25d0a77df4064330 = append(__obf_25d0a77df4064330, __obf_e93fc1cf96e54053[__obf_045c7fea9195ede5])

	for __obf_2a103e2ab03dbd0a, __obf_dffb606f75e8b13c := 0, len(__obf_25d0a77df4064330)-1; __obf_2a103e2ab03dbd0a < __obf_dffb606f75e8b13c; __obf_2a103e2ab03dbd0a, __obf_dffb606f75e8b13c = __obf_2a103e2ab03dbd0a+1, __obf_dffb606f75e8b13c-1 {
		__obf_25d0a77df4064330[__obf_2a103e2ab03dbd0a], __obf_25d0a77df4064330[__obf_dffb606f75e8b13c] = __obf_25d0a77df4064330[__obf_dffb606f75e8b13c], __obf_25d0a77df4064330[__obf_2a103e2ab03dbd0a]
	}

	return string(__obf_25d0a77df4064330)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_25d0a77df4064330 []byte) (ID, error) {

	var __obf_721283c42292359f int64

	for __obf_d43488f90f0960a8 := range __obf_25d0a77df4064330 {
		if __obf_487b5deba62869c7[__obf_25d0a77df4064330[__obf_d43488f90f0960a8]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_721283c42292359f = __obf_721283c42292359f*58 + int64(__obf_487b5deba62869c7[__obf_25d0a77df4064330[__obf_d43488f90f0960a8]])
	}

	return ID(__obf_721283c42292359f), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_045c7fea9195ede5 ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_045c7fea9195ede5.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_721283c42292359f string) (ID, error) {
	__obf_25d0a77df4064330, __obf_38c5d446257f8d9a := base64.StdEncoding.DecodeString(__obf_721283c42292359f)
	if __obf_38c5d446257f8d9a != nil {
		return -1, __obf_38c5d446257f8d9a
	}
	return ParseBytes(__obf_25d0a77df4064330)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_045c7fea9195ede5 ID) Bytes() []byte {
	return []byte(__obf_045c7fea9195ede5.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_721283c42292359f []byte) (ID, error) {
	__obf_d43488f90f0960a8, __obf_38c5d446257f8d9a := strconv.ParseInt(string(__obf_721283c42292359f), 10, 64)
	return ID(__obf_d43488f90f0960a8), __obf_38c5d446257f8d9a
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_045c7fea9195ede5 ID) IntBytes() [8]byte {
	var __obf_25d0a77df4064330 [8]byte
	binary.BigEndian.PutUint64(__obf_25d0a77df4064330[:], uint64(__obf_045c7fea9195ede5))
	return __obf_25d0a77df4064330
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_721283c42292359f [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_721283c42292359f[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_045c7fea9195ede5 ID) Time() int64 {
	return (int64(__obf_045c7fea9195ede5) >> __obf_43facb4bfd5cad5b) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_045c7fea9195ede5 ID) Node() int64 {
	return int64(__obf_045c7fea9195ede5) & __obf_b87a81d04e8ecdc5 >> __obf_6fde542eef241ce8
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_045c7fea9195ede5 ID) Step() int64 {
	return int64(__obf_045c7fea9195ede5) & __obf_118ef6870c08dca5
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_045c7fea9195ede5 ID) MarshalJSON() ([]byte, error) {
	__obf_0d6981704968d856 := make([]byte, 0, 22)
	__obf_0d6981704968d856 = append(__obf_0d6981704968d856, '"')
	__obf_0d6981704968d856 = strconv.AppendInt(__obf_0d6981704968d856, int64(__obf_045c7fea9195ede5), 10)
	__obf_0d6981704968d856 = append(__obf_0d6981704968d856, '"')
	return __obf_0d6981704968d856, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_045c7fea9195ede5 *ID) UnmarshalJSON(__obf_25d0a77df4064330 []byte) error {
	if len(__obf_25d0a77df4064330) < 3 || __obf_25d0a77df4064330[0] != '"' || __obf_25d0a77df4064330[len(__obf_25d0a77df4064330)-1] != '"' {
		return JSONSyntaxError{__obf_25d0a77df4064330}
	}
	__obf_d43488f90f0960a8, __obf_38c5d446257f8d9a := strconv.ParseInt(string(__obf_25d0a77df4064330[1:len(__obf_25d0a77df4064330)-1]), 10, 64)
	if __obf_38c5d446257f8d9a != nil {
		return __obf_38c5d446257f8d9a
	}

	*__obf_045c7fea9195ede5 = ID(__obf_d43488f90f0960a8)
	return nil
}
