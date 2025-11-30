package __obf_a51a64e21f311927

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
	__obf_7017cfd2525a0d6a sync. // DEPRECATED: the below four variables will be removed in a future release.
				Mutex
	__obf_ee3b6745686ec441 int64 = -1 ^ (-1 << NodeBits)
	__obf_21db2839b3dd8ae6       = __obf_ee3b6745686ec441 <<
		StepBits
	__obf_d92f7e0b93bd02e3 int64 = -1 ^ (-1 << StepBits)
	__obf_f42e34b441867698       = NodeBits + StepBits
	__obf_908533553631ec3b       = StepBits
)

const __obf_f439734ad9ce9646 = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_60c8d4237ef5f2f0 [256]byte

const __obf_1f8cedac51ff93e8 = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_dbd47d27df61fe35 [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_4258cf00a45a4fc3 []byte }

func (__obf_a59a886208dd2746 JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_a59a886208dd2746.

		// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
		__obf_4258cf00a45a4fc3))
}

var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func init() {

	for __obf_c022f98255ac9ddc := 0; __obf_c022f98255ac9ddc < len(__obf_dbd47d27df61fe35); __obf_c022f98255ac9ddc++ {
		__obf_dbd47d27df61fe35[__obf_c022f98255ac9ddc] = 0xFF
	}

	for __obf_c022f98255ac9ddc := 0; __obf_c022f98255ac9ddc < len(__obf_1f8cedac51ff93e8); __obf_c022f98255ac9ddc++ {
		__obf_dbd47d27df61fe35[__obf_1f8cedac51ff93e8[__obf_c022f98255ac9ddc]] = byte(__obf_c022f98255ac9ddc)
	}

	for __obf_c022f98255ac9ddc := 0; __obf_c022f98255ac9ddc < len(__obf_60c8d4237ef5f2f0); __obf_c022f98255ac9ddc++ {
		__obf_60c8d4237ef5f2f0[__obf_c022f98255ac9ddc] = 0xFF
	}

	for __obf_c022f98255ac9ddc := 0; __obf_c022f98255ac9ddc < len(__obf_f439734ad9ce9646); __obf_c022f98255ac9ddc++ {
		__obf_60c8d4237ef5f2f0[__obf_f439734ad9ce9646[__obf_c022f98255ac9ddc]] = byte(__obf_c022f98255ac9ddc)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_7017cfd2525a0d6a sync.Mutex
	__obf_80c6d826465b4a8c time.Time
	time                   int64
	__obf_4be2dca4bfc31874 int64
	__obf_4cbed4c914721d88 int64
	__obf_ee3b6745686ec441 int64
	__obf_21db2839b3dd8ae6 int64
	__obf_d92f7e0b93bd02e3 int64
	__obf_f42e34b441867698 uint8
	__obf_908533553631ec3b uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_4be2dca4bfc31874 int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	__obf_7017cfd2525a0d6a.
		// re-calc in case custom NodeBits or StepBits were set
		// DEPRECATED: the below block will be removed in a future release.
		Lock()
	__obf_ee3b6745686ec441 = -1 ^ (-1 << NodeBits)
	__obf_21db2839b3dd8ae6 = __obf_ee3b6745686ec441 << StepBits
	__obf_d92f7e0b93bd02e3 = -1 ^ (-1 << StepBits)
	__obf_f42e34b441867698 = NodeBits + StepBits
	__obf_908533553631ec3b = StepBits
	__obf_7017cfd2525a0d6a.
		Unlock()
	__obf_9f6decc0e2d1ee43 := Node{}
	__obf_9f6decc0e2d1ee43.__obf_4be2dca4bfc31874 = __obf_4be2dca4bfc31874
	__obf_9f6decc0e2d1ee43.__obf_ee3b6745686ec441 = -1 ^ (-1 << NodeBits)
	__obf_9f6decc0e2d1ee43.__obf_21db2839b3dd8ae6 = __obf_9f6decc0e2d1ee43.__obf_ee3b6745686ec441 << StepBits
	__obf_9f6decc0e2d1ee43.__obf_d92f7e0b93bd02e3 = -1 ^ (-1 << StepBits)
	__obf_9f6decc0e2d1ee43.__obf_f42e34b441867698 = NodeBits + StepBits
	__obf_9f6decc0e2d1ee43.__obf_908533553631ec3b = StepBits

	if __obf_9f6decc0e2d1ee43.__obf_4be2dca4bfc31874 < 0 || __obf_9f6decc0e2d1ee43.__obf_4be2dca4bfc31874 > __obf_9f6decc0e2d1ee43.__obf_ee3b6745686ec441 {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_9f6decc0e2d1ee43.__obf_ee3b6745686ec441, 10))
	}

	var __obf_d886c2ecbb505ea7 = time.Now()
	__obf_9f6decc0e2d1ee43.
		// add time.Duration to curTime to make sure we use the monotonic clock if available
		__obf_80c6d826465b4a8c = __obf_d886c2ecbb505ea7.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_d886c2ecbb505ea7))

	return &__obf_9f6decc0e2d1ee43, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_9f6decc0e2d1ee43 *Node) Generate() ID {
	__obf_9f6decc0e2d1ee43.__obf_7017cfd2525a0d6a.
		Lock()
	defer __obf_9f6decc0e2d1ee43.__obf_7017cfd2525a0d6a.Unlock()
	__obf_759d5fd6611493a3 := time.Since(__obf_9f6decc0e2d1ee43.__obf_80c6d826465b4a8c).Milliseconds()

	if __obf_759d5fd6611493a3 == __obf_9f6decc0e2d1ee43.time {
		__obf_9f6decc0e2d1ee43.__obf_4cbed4c914721d88 = (__obf_9f6decc0e2d1ee43.__obf_4cbed4c914721d88 + 1) & __obf_9f6decc0e2d1ee43.__obf_d92f7e0b93bd02e3

		if __obf_9f6decc0e2d1ee43.__obf_4cbed4c914721d88 == 0 {
			for __obf_759d5fd6611493a3 <= __obf_9f6decc0e2d1ee43.time {
				__obf_759d5fd6611493a3 = time.Since(__obf_9f6decc0e2d1ee43.__obf_80c6d826465b4a8c).Milliseconds()
			}
		}
	} else {
		__obf_9f6decc0e2d1ee43.__obf_4cbed4c914721d88 = 0
	}
	__obf_9f6decc0e2d1ee43.
		time = __obf_759d5fd6611493a3
	__obf_5e311d3f7620aa0c := ID((__obf_759d5fd6611493a3)<<__obf_9f6decc0e2d1ee43.__obf_f42e34b441867698 |
		(__obf_9f6decc0e2d1ee43.__obf_4be2dca4bfc31874 << __obf_9f6decc0e2d1ee43.__obf_908533553631ec3b) |
		(__obf_9f6decc0e2d1ee43.__obf_4cbed4c914721d88),
	)

	return __obf_5e311d3f7620aa0c
}

// Int64 returns an int64 of the snowflake ID
func (__obf_c0d7ae038ed58ced ID) Int64() int64 {
	return int64(__obf_c0d7ae038ed58ced)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_4c112818117e5211 int64) ID {
	return ID(__obf_4c112818117e5211)
}

// String returns a string of the snowflake ID
func (__obf_c0d7ae038ed58ced ID) String() string {
	return strconv.FormatInt(int64(__obf_c0d7ae038ed58ced), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_4c112818117e5211 string) (ID, error) {
	__obf_c022f98255ac9ddc, __obf_886df8acc7f50b76 := strconv.ParseInt(__obf_4c112818117e5211, 10, 64)
	return ID(__obf_c022f98255ac9ddc), __obf_886df8acc7f50b76

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_c0d7ae038ed58ced ID) Base2() string {
	return strconv.FormatInt(int64(__obf_c0d7ae038ed58ced), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_4c112818117e5211 string) (ID, error) {
	__obf_c022f98255ac9ddc, __obf_886df8acc7f50b76 := strconv.ParseInt(__obf_4c112818117e5211, 2, 64)
	return ID(__obf_c022f98255ac9ddc), __obf_886df8acc7f50b76
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_c0d7ae038ed58ced ID) Base32() string {

	if __obf_c0d7ae038ed58ced < 32 {
		return string(__obf_f439734ad9ce9646[__obf_c0d7ae038ed58ced])
	}
	__obf_3a8317949773578f := make([]byte, 0, 12)
	for __obf_c0d7ae038ed58ced >= 32 {
		__obf_3a8317949773578f = append(__obf_3a8317949773578f, __obf_f439734ad9ce9646[__obf_c0d7ae038ed58ced%32])
		__obf_c0d7ae038ed58ced /= 32
	}
	__obf_3a8317949773578f = append(__obf_3a8317949773578f, __obf_f439734ad9ce9646[__obf_c0d7ae038ed58ced])

	for __obf_137cb83870c72e80, __obf_cfa5f85c1dc1f775 := 0, len(__obf_3a8317949773578f)-1; __obf_137cb83870c72e80 < __obf_cfa5f85c1dc1f775; __obf_137cb83870c72e80, __obf_cfa5f85c1dc1f775 = __obf_137cb83870c72e80+1, __obf_cfa5f85c1dc1f775-1 {
		__obf_3a8317949773578f[__obf_137cb83870c72e80], __obf_3a8317949773578f[__obf_cfa5f85c1dc1f775] = __obf_3a8317949773578f[__obf_cfa5f85c1dc1f775], __obf_3a8317949773578f[__obf_137cb83870c72e80]
	}

	return string(__obf_3a8317949773578f)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_3a8317949773578f []byte) (ID, error) {

	var __obf_4c112818117e5211 int64

	for __obf_c022f98255ac9ddc := range __obf_3a8317949773578f {
		if __obf_60c8d4237ef5f2f0[__obf_3a8317949773578f[__obf_c022f98255ac9ddc]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_4c112818117e5211 = __obf_4c112818117e5211*32 + int64(__obf_60c8d4237ef5f2f0[__obf_3a8317949773578f[__obf_c022f98255ac9ddc]])
	}

	return ID(__obf_4c112818117e5211), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_c0d7ae038ed58ced ID) Base36() string {
	return strconv.FormatInt(int64(__obf_c0d7ae038ed58ced), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_4c112818117e5211 string) (ID, error) {
	__obf_c022f98255ac9ddc, __obf_886df8acc7f50b76 := strconv.ParseInt(__obf_4c112818117e5211, 36, 64)
	return ID(__obf_c022f98255ac9ddc), __obf_886df8acc7f50b76
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_c0d7ae038ed58ced ID) Base58() string {

	if __obf_c0d7ae038ed58ced < 58 {
		return string(__obf_1f8cedac51ff93e8[__obf_c0d7ae038ed58ced])
	}
	__obf_3a8317949773578f := make([]byte, 0, 11)
	for __obf_c0d7ae038ed58ced >= 58 {
		__obf_3a8317949773578f = append(__obf_3a8317949773578f, __obf_1f8cedac51ff93e8[__obf_c0d7ae038ed58ced%58])
		__obf_c0d7ae038ed58ced /= 58
	}
	__obf_3a8317949773578f = append(__obf_3a8317949773578f, __obf_1f8cedac51ff93e8[__obf_c0d7ae038ed58ced])

	for __obf_137cb83870c72e80, __obf_cfa5f85c1dc1f775 := 0, len(__obf_3a8317949773578f)-1; __obf_137cb83870c72e80 < __obf_cfa5f85c1dc1f775; __obf_137cb83870c72e80, __obf_cfa5f85c1dc1f775 = __obf_137cb83870c72e80+1, __obf_cfa5f85c1dc1f775-1 {
		__obf_3a8317949773578f[__obf_137cb83870c72e80], __obf_3a8317949773578f[__obf_cfa5f85c1dc1f775] = __obf_3a8317949773578f[__obf_cfa5f85c1dc1f775], __obf_3a8317949773578f[__obf_137cb83870c72e80]
	}

	return string(__obf_3a8317949773578f)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_3a8317949773578f []byte) (ID, error) {

	var __obf_4c112818117e5211 int64

	for __obf_c022f98255ac9ddc := range __obf_3a8317949773578f {
		if __obf_dbd47d27df61fe35[__obf_3a8317949773578f[__obf_c022f98255ac9ddc]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_4c112818117e5211 = __obf_4c112818117e5211*58 + int64(__obf_dbd47d27df61fe35[__obf_3a8317949773578f[__obf_c022f98255ac9ddc]])
	}

	return ID(__obf_4c112818117e5211), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_c0d7ae038ed58ced ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_c0d7ae038ed58ced.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_4c112818117e5211 string) (ID, error) {
	__obf_3a8317949773578f, __obf_886df8acc7f50b76 := base64.StdEncoding.DecodeString(__obf_4c112818117e5211)
	if __obf_886df8acc7f50b76 != nil {
		return -1, __obf_886df8acc7f50b76
	}
	return ParseBytes(__obf_3a8317949773578f)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_c0d7ae038ed58ced ID) Bytes() []byte {
	return []byte(__obf_c0d7ae038ed58ced.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_4c112818117e5211 []byte) (ID, error) {
	__obf_c022f98255ac9ddc, __obf_886df8acc7f50b76 := strconv.ParseInt(string(__obf_4c112818117e5211), 10, 64)
	return ID(__obf_c022f98255ac9ddc), __obf_886df8acc7f50b76
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_c0d7ae038ed58ced ID) IntBytes() [8]byte {
	var __obf_3a8317949773578f [8]byte
	binary.BigEndian.PutUint64(__obf_3a8317949773578f[:], uint64(__obf_c0d7ae038ed58ced))
	return __obf_3a8317949773578f
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_4c112818117e5211 [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_4c112818117e5211[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_c0d7ae038ed58ced ID) Time() int64 {
	return (int64(__obf_c0d7ae038ed58ced) >> __obf_f42e34b441867698) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_c0d7ae038ed58ced ID) Node() int64 {
	return int64(__obf_c0d7ae038ed58ced) & __obf_21db2839b3dd8ae6 >> __obf_908533553631ec3b
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_c0d7ae038ed58ced ID) Step() int64 {
	return int64(__obf_c0d7ae038ed58ced) & __obf_d92f7e0b93bd02e3
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_c0d7ae038ed58ced ID) MarshalJSON() ([]byte, error) {
	__obf_98bc57b6b2327cd5 := make([]byte, 0, 22)
	__obf_98bc57b6b2327cd5 = append(__obf_98bc57b6b2327cd5, '"')
	__obf_98bc57b6b2327cd5 = strconv.AppendInt(__obf_98bc57b6b2327cd5, int64(__obf_c0d7ae038ed58ced), 10)
	__obf_98bc57b6b2327cd5 = append(__obf_98bc57b6b2327cd5, '"')
	return __obf_98bc57b6b2327cd5, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_c0d7ae038ed58ced *ID) UnmarshalJSON(__obf_3a8317949773578f []byte) error {
	if len(__obf_3a8317949773578f) < 3 || __obf_3a8317949773578f[0] != '"' || __obf_3a8317949773578f[len(__obf_3a8317949773578f)-1] != '"' {
		return JSONSyntaxError{__obf_3a8317949773578f}
	}
	__obf_c022f98255ac9ddc, __obf_886df8acc7f50b76 := strconv.ParseInt(string(__obf_3a8317949773578f[1:len(__obf_3a8317949773578f)-1]), 10, 64)
	if __obf_886df8acc7f50b76 != nil {
		return __obf_886df8acc7f50b76
	}

	*__obf_c0d7ae038ed58ced = ID(__obf_c022f98255ac9ddc)
	return nil
}
