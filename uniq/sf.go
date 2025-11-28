package __obf_8fe28252c1b01dfe

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
	__obf_26212263d98d9893 sync.Mutex
	__obf_bf43347003457cd1 int64 = -1 ^ (-1 << NodeBits)
	__obf_38f01199dd482734       = __obf_bf43347003457cd1 << StepBits
	__obf_e496ab64c263f602 int64 = -1 ^ (-1 << StepBits)
	__obf_95b2a2010196fdc2       = NodeBits + StepBits
	__obf_db68f943ae362041       = StepBits
)

const __obf_849d08ceb4404870 = "ybndrfg8ejkmcpqxot1uwisza345h769"

var __obf_7bb03111697929f8 [256]byte

const __obf_f1d3de1161e30064 = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var __obf_4af0ed504f13b712 [256]byte

// A JSONSyntaxError is returned from UnmarshalJSON if an invalid ID is provided.
type JSONSyntaxError struct{ __obf_cceb3697baa69df2 []byte }

func (__obf_da5e58bcd4150864 JSONSyntaxError) Error() string {
	return fmt.Sprintf("invalid snowflake ID %q", string(__obf_da5e58bcd4150864.__obf_cceb3697baa69df2))
}

// ErrInvalidBase58 is returned by ParseBase58 when given an invalid []byte
var ErrInvalidBase58 = errors.New("invalid base58")

// ErrInvalidBase32 is returned by ParseBase32 when given an invalid []byte
var ErrInvalidBase32 = errors.New("invalid base32")

// Create maps for decoding Base58/Base32.
// This speeds up the process tremendously.
func __obf_78249b7796f99b00() {

	for __obf_a6dd0a7023f6fe3f := 0; __obf_a6dd0a7023f6fe3f < len(__obf_4af0ed504f13b712); __obf_a6dd0a7023f6fe3f++ {
		__obf_4af0ed504f13b712[__obf_a6dd0a7023f6fe3f] = 0xFF
	}

	for __obf_a6dd0a7023f6fe3f := 0; __obf_a6dd0a7023f6fe3f < len(__obf_f1d3de1161e30064); __obf_a6dd0a7023f6fe3f++ {
		__obf_4af0ed504f13b712[__obf_f1d3de1161e30064[__obf_a6dd0a7023f6fe3f]] = byte(__obf_a6dd0a7023f6fe3f)
	}

	for __obf_a6dd0a7023f6fe3f := 0; __obf_a6dd0a7023f6fe3f < len(__obf_7bb03111697929f8); __obf_a6dd0a7023f6fe3f++ {
		__obf_7bb03111697929f8[__obf_a6dd0a7023f6fe3f] = 0xFF
	}

	for __obf_a6dd0a7023f6fe3f := 0; __obf_a6dd0a7023f6fe3f < len(__obf_849d08ceb4404870); __obf_a6dd0a7023f6fe3f++ {
		__obf_7bb03111697929f8[__obf_849d08ceb4404870[__obf_a6dd0a7023f6fe3f]] = byte(__obf_a6dd0a7023f6fe3f)
	}
}

// A Node struct holds the basic information needed for a snowflake generator
// node
type Node struct {
	__obf_26212263d98d9893 sync.Mutex
	__obf_bcb824e4f9282ca5 time.Time
	time                   int64
	__obf_561ced1d8cce088a int64
	__obf_02067b47387f16a4 int64

	__obf_bf43347003457cd1 int64
	__obf_38f01199dd482734 int64
	__obf_e496ab64c263f602 int64
	__obf_95b2a2010196fdc2 uint8
	__obf_db68f943ae362041 uint8
}

// An ID is a custom type used for a snowflake ID.  This is used so we can
// attach methods onto the ID.
type ID int64

// NewNode returns a new snowflake node that can be used to generate snowflake
// IDs
func NewNode(__obf_561ced1d8cce088a int64) (*Node, error) {

	if NodeBits+StepBits > 22 {
		return nil, errors.New("remember, you have a total 22 bits to share between Node/Step")
	}
	// re-calc in case custom NodeBits or StepBits were set
	// DEPRECATED: the below block will be removed in a future release.
	__obf_26212263d98d9893.Lock()
	__obf_bf43347003457cd1 = -1 ^ (-1 << NodeBits)
	__obf_38f01199dd482734 = __obf_bf43347003457cd1 << StepBits
	__obf_e496ab64c263f602 = -1 ^ (-1 << StepBits)
	__obf_95b2a2010196fdc2 = NodeBits + StepBits
	__obf_db68f943ae362041 = StepBits
	__obf_26212263d98d9893.Unlock()

	__obf_20d0dbd4aeeff396 := Node{}
	__obf_20d0dbd4aeeff396.__obf_561ced1d8cce088a = __obf_561ced1d8cce088a
	__obf_20d0dbd4aeeff396.__obf_bf43347003457cd1 = -1 ^ (-1 << NodeBits)
	__obf_20d0dbd4aeeff396.__obf_38f01199dd482734 = __obf_20d0dbd4aeeff396.__obf_bf43347003457cd1 << StepBits
	__obf_20d0dbd4aeeff396.__obf_e496ab64c263f602 = -1 ^ (-1 << StepBits)
	__obf_20d0dbd4aeeff396.__obf_95b2a2010196fdc2 = NodeBits + StepBits
	__obf_20d0dbd4aeeff396.__obf_db68f943ae362041 = StepBits

	if __obf_20d0dbd4aeeff396.__obf_561ced1d8cce088a < 0 || __obf_20d0dbd4aeeff396.__obf_561ced1d8cce088a > __obf_20d0dbd4aeeff396.__obf_bf43347003457cd1 {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(__obf_20d0dbd4aeeff396.__obf_bf43347003457cd1, 10))
	}

	var __obf_b5b81f6b21d0e642 = time.Now()
	// add time.Duration to curTime to make sure we use the monotonic clock if available
	__obf_20d0dbd4aeeff396.__obf_bcb824e4f9282ca5 = __obf_b5b81f6b21d0e642.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(__obf_b5b81f6b21d0e642))

	return &__obf_20d0dbd4aeeff396, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (__obf_20d0dbd4aeeff396 *Node) Generate() ID {

	__obf_20d0dbd4aeeff396.__obf_26212263d98d9893.Lock()
	defer __obf_20d0dbd4aeeff396.__obf_26212263d98d9893.Unlock()

	__obf_194d6051ead2e2f1 := time.Since(__obf_20d0dbd4aeeff396.__obf_bcb824e4f9282ca5).Milliseconds()

	if __obf_194d6051ead2e2f1 == __obf_20d0dbd4aeeff396.time {
		__obf_20d0dbd4aeeff396.__obf_02067b47387f16a4 = (__obf_20d0dbd4aeeff396.__obf_02067b47387f16a4 + 1) & __obf_20d0dbd4aeeff396.__obf_e496ab64c263f602

		if __obf_20d0dbd4aeeff396.__obf_02067b47387f16a4 == 0 {
			for __obf_194d6051ead2e2f1 <= __obf_20d0dbd4aeeff396.time {
				__obf_194d6051ead2e2f1 = time.Since(__obf_20d0dbd4aeeff396.__obf_bcb824e4f9282ca5).Milliseconds()
			}
		}
	} else {
		__obf_20d0dbd4aeeff396.__obf_02067b47387f16a4 = 0
	}

	__obf_20d0dbd4aeeff396.time = __obf_194d6051ead2e2f1

	__obf_41b6a03d01820ee6 := ID((__obf_194d6051ead2e2f1)<<__obf_20d0dbd4aeeff396.__obf_95b2a2010196fdc2 |
		(__obf_20d0dbd4aeeff396.__obf_561ced1d8cce088a << __obf_20d0dbd4aeeff396.__obf_db68f943ae362041) |
		(__obf_20d0dbd4aeeff396.__obf_02067b47387f16a4),
	)

	return __obf_41b6a03d01820ee6
}

// Int64 returns an int64 of the snowflake ID
func (__obf_be6bc8fa11f8e866 ID) Int64() int64 {
	return int64(__obf_be6bc8fa11f8e866)
}

// ParseInt64 converts an int64 into a snowflake ID
func ParseInt64(__obf_b6d2ddd191a19ef1 int64) ID {
	return ID(__obf_b6d2ddd191a19ef1)
}

// String returns a string of the snowflake ID
func (__obf_be6bc8fa11f8e866 ID) String() string {
	return strconv.FormatInt(int64(__obf_be6bc8fa11f8e866), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(__obf_b6d2ddd191a19ef1 string) (ID, error) {
	__obf_a6dd0a7023f6fe3f, __obf_1082dd6e56192e3a := strconv.ParseInt(__obf_b6d2ddd191a19ef1, 10, 64)
	return ID(__obf_a6dd0a7023f6fe3f), __obf_1082dd6e56192e3a

}

// Base2 returns a string base2 of the snowflake ID
func (__obf_be6bc8fa11f8e866 ID) Base2() string {
	return strconv.FormatInt(int64(__obf_be6bc8fa11f8e866), 2)
}

// ParseBase2 converts a Base2 string into a snowflake ID
func ParseBase2(__obf_b6d2ddd191a19ef1 string) (ID, error) {
	__obf_a6dd0a7023f6fe3f, __obf_1082dd6e56192e3a := strconv.ParseInt(__obf_b6d2ddd191a19ef1, 2, 64)
	return ID(__obf_a6dd0a7023f6fe3f), __obf_1082dd6e56192e3a
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func (__obf_be6bc8fa11f8e866 ID) Base32() string {

	if __obf_be6bc8fa11f8e866 < 32 {
		return string(__obf_849d08ceb4404870[__obf_be6bc8fa11f8e866])
	}

	__obf_6d2dd22fb831afa6 := make([]byte, 0, 12)
	for __obf_be6bc8fa11f8e866 >= 32 {
		__obf_6d2dd22fb831afa6 = append(__obf_6d2dd22fb831afa6, __obf_849d08ceb4404870[__obf_be6bc8fa11f8e866%32])
		__obf_be6bc8fa11f8e866 /= 32
	}
	__obf_6d2dd22fb831afa6 = append(__obf_6d2dd22fb831afa6, __obf_849d08ceb4404870[__obf_be6bc8fa11f8e866])

	for __obf_cf72d1dd2ebb40bf, __obf_77cc778ff2548c44 := 0, len(__obf_6d2dd22fb831afa6)-1; __obf_cf72d1dd2ebb40bf < __obf_77cc778ff2548c44; __obf_cf72d1dd2ebb40bf, __obf_77cc778ff2548c44 = __obf_cf72d1dd2ebb40bf+1, __obf_77cc778ff2548c44-1 {
		__obf_6d2dd22fb831afa6[__obf_cf72d1dd2ebb40bf], __obf_6d2dd22fb831afa6[__obf_77cc778ff2548c44] = __obf_6d2dd22fb831afa6[__obf_77cc778ff2548c44], __obf_6d2dd22fb831afa6[__obf_cf72d1dd2ebb40bf]
	}

	return string(__obf_6d2dd22fb831afa6)
}

// ParseBase32 parses a base32 []byte into a snowflake ID
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation.
func ParseBase32(__obf_6d2dd22fb831afa6 []byte) (ID, error) {

	var __obf_b6d2ddd191a19ef1 int64

	for __obf_a6dd0a7023f6fe3f := range __obf_6d2dd22fb831afa6 {
		if __obf_7bb03111697929f8[__obf_6d2dd22fb831afa6[__obf_a6dd0a7023f6fe3f]] == 0xFF {
			return -1, ErrInvalidBase32
		}
		__obf_b6d2ddd191a19ef1 = __obf_b6d2ddd191a19ef1*32 + int64(__obf_7bb03111697929f8[__obf_6d2dd22fb831afa6[__obf_a6dd0a7023f6fe3f]])
	}

	return ID(__obf_b6d2ddd191a19ef1), nil
}

// Base36 returns a base36 string of the snowflake ID
func (__obf_be6bc8fa11f8e866 ID) Base36() string {
	return strconv.FormatInt(int64(__obf_be6bc8fa11f8e866), 36)
}

// ParseBase36 converts a Base36 string into a snowflake ID
func ParseBase36(__obf_b6d2ddd191a19ef1 string) (ID, error) {
	__obf_a6dd0a7023f6fe3f, __obf_1082dd6e56192e3a := strconv.ParseInt(__obf_b6d2ddd191a19ef1, 36, 64)
	return ID(__obf_a6dd0a7023f6fe3f), __obf_1082dd6e56192e3a
}

// Base58 returns a base58 string of the snowflake ID
func (__obf_be6bc8fa11f8e866 ID) Base58() string {

	if __obf_be6bc8fa11f8e866 < 58 {
		return string(__obf_f1d3de1161e30064[__obf_be6bc8fa11f8e866])
	}

	__obf_6d2dd22fb831afa6 := make([]byte, 0, 11)
	for __obf_be6bc8fa11f8e866 >= 58 {
		__obf_6d2dd22fb831afa6 = append(__obf_6d2dd22fb831afa6, __obf_f1d3de1161e30064[__obf_be6bc8fa11f8e866%58])
		__obf_be6bc8fa11f8e866 /= 58
	}
	__obf_6d2dd22fb831afa6 = append(__obf_6d2dd22fb831afa6, __obf_f1d3de1161e30064[__obf_be6bc8fa11f8e866])

	for __obf_cf72d1dd2ebb40bf, __obf_77cc778ff2548c44 := 0, len(__obf_6d2dd22fb831afa6)-1; __obf_cf72d1dd2ebb40bf < __obf_77cc778ff2548c44; __obf_cf72d1dd2ebb40bf, __obf_77cc778ff2548c44 = __obf_cf72d1dd2ebb40bf+1, __obf_77cc778ff2548c44-1 {
		__obf_6d2dd22fb831afa6[__obf_cf72d1dd2ebb40bf], __obf_6d2dd22fb831afa6[__obf_77cc778ff2548c44] = __obf_6d2dd22fb831afa6[__obf_77cc778ff2548c44], __obf_6d2dd22fb831afa6[__obf_cf72d1dd2ebb40bf]
	}

	return string(__obf_6d2dd22fb831afa6)
}

// ParseBase58 parses a base58 []byte into a snowflake ID
func ParseBase58(__obf_6d2dd22fb831afa6 []byte) (ID, error) {

	var __obf_b6d2ddd191a19ef1 int64

	for __obf_a6dd0a7023f6fe3f := range __obf_6d2dd22fb831afa6 {
		if __obf_4af0ed504f13b712[__obf_6d2dd22fb831afa6[__obf_a6dd0a7023f6fe3f]] == 0xFF {
			return -1, ErrInvalidBase58
		}
		__obf_b6d2ddd191a19ef1 = __obf_b6d2ddd191a19ef1*58 + int64(__obf_4af0ed504f13b712[__obf_6d2dd22fb831afa6[__obf_a6dd0a7023f6fe3f]])
	}

	return ID(__obf_b6d2ddd191a19ef1), nil
}

// Base64 returns a base64 string of the snowflake ID
func (__obf_be6bc8fa11f8e866 ID) Base64() string {
	return base64.StdEncoding.EncodeToString(__obf_be6bc8fa11f8e866.Bytes())
}

// ParseBase64 converts a base64 string into a snowflake ID
func ParseBase64(__obf_b6d2ddd191a19ef1 string) (ID, error) {
	__obf_6d2dd22fb831afa6, __obf_1082dd6e56192e3a := base64.StdEncoding.DecodeString(__obf_b6d2ddd191a19ef1)
	if __obf_1082dd6e56192e3a != nil {
		return -1, __obf_1082dd6e56192e3a
	}
	return ParseBytes(__obf_6d2dd22fb831afa6)

}

// Bytes returns a byte slice of the snowflake ID
func (__obf_be6bc8fa11f8e866 ID) Bytes() []byte {
	return []byte(__obf_be6bc8fa11f8e866.String())
}

// ParseBytes converts a byte slice into a snowflake ID
func ParseBytes(__obf_b6d2ddd191a19ef1 []byte) (ID, error) {
	__obf_a6dd0a7023f6fe3f, __obf_1082dd6e56192e3a := strconv.ParseInt(string(__obf_b6d2ddd191a19ef1), 10, 64)
	return ID(__obf_a6dd0a7023f6fe3f), __obf_1082dd6e56192e3a
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (__obf_be6bc8fa11f8e866 ID) IntBytes() [8]byte {
	var __obf_6d2dd22fb831afa6 [8]byte
	binary.BigEndian.PutUint64(__obf_6d2dd22fb831afa6[:], uint64(__obf_be6bc8fa11f8e866))
	return __obf_6d2dd22fb831afa6
}

// ParseIntBytes converts an array of bytes encoded as big endian integer as
// a snowflake ID
func ParseIntBytes(__obf_b6d2ddd191a19ef1 [8]byte) ID {
	return ID(int64(binary.BigEndian.Uint64(__obf_b6d2ddd191a19ef1[:])))
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (__obf_be6bc8fa11f8e866 ID) Time() int64 {
	return (int64(__obf_be6bc8fa11f8e866) >> __obf_95b2a2010196fdc2) + Epoch
}

// Node returns an int64 of the snowflake ID node number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_be6bc8fa11f8e866 ID) Node() int64 {
	return int64(__obf_be6bc8fa11f8e866) & __obf_38f01199dd482734 >> __obf_db68f943ae362041
}

// Step returns an int64 of the snowflake step (or sequence) number
// DEPRECATED: the below function will be removed in a future release.
func (__obf_be6bc8fa11f8e866 ID) Step() int64 {
	return int64(__obf_be6bc8fa11f8e866) & __obf_e496ab64c263f602
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (__obf_be6bc8fa11f8e866 ID) MarshalJSON() ([]byte, error) {
	__obf_7a64b48aab618f9d := make([]byte, 0, 22)
	__obf_7a64b48aab618f9d = append(__obf_7a64b48aab618f9d, '"')
	__obf_7a64b48aab618f9d = strconv.AppendInt(__obf_7a64b48aab618f9d, int64(__obf_be6bc8fa11f8e866), 10)
	__obf_7a64b48aab618f9d = append(__obf_7a64b48aab618f9d, '"')
	return __obf_7a64b48aab618f9d, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (__obf_be6bc8fa11f8e866 *ID) UnmarshalJSON(__obf_6d2dd22fb831afa6 []byte) error {
	if len(__obf_6d2dd22fb831afa6) < 3 || __obf_6d2dd22fb831afa6[0] != '"' || __obf_6d2dd22fb831afa6[len(__obf_6d2dd22fb831afa6)-1] != '"' {
		return JSONSyntaxError{__obf_6d2dd22fb831afa6}
	}

	__obf_a6dd0a7023f6fe3f, __obf_1082dd6e56192e3a := strconv.ParseInt(string(__obf_6d2dd22fb831afa6[1:len(__obf_6d2dd22fb831afa6)-1]), 10, 64)
	if __obf_1082dd6e56192e3a != nil {
		return __obf_1082dd6e56192e3a
	}

	*__obf_be6bc8fa11f8e866 = ID(__obf_a6dd0a7023f6fe3f)
	return nil
}
