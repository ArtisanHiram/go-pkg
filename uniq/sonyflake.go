// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_7d8ac56be2e11a40

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

type InterfaceAddrs func() ([]net.Addr, error)

// These constants are the bit lengths of Sonyflake ID parts.
const (
	BitLenTime      = 39                               // bit length of time
	BitLenSequence  = 8                                // bit length of sequence number
	BitLenMachineID = 63 - BitLenTime - BitLenSequence // bit length of machine id
)

// Settings configures Sonyflake:
//
// StartTime is the time since which the Sonyflake time is defined as the elapsed time.
// If StartTime is 0, the start time of the Sonyflake is set to "2014-09-01 00:00:00 +0000 UTC".
// If StartTime is ahead of the current time, Sonyflake is not created.
//
// MachineID returns the unique ID of the Sonyflake instance.
// If MachineID returns an error, Sonyflake is not created.
// If MachineID is nil, default MachineID is used.
// Default MachineID returns the lower 16 bits of the private IP address.
//
// CheckMachineID validates the uniqueness of the machine ID.
// If CheckMachineID returns false, Sonyflake is not created.
// If CheckMachineID is nil, no validation is done.
type Settings struct {
	StartTime      time.Time
	MachineID      func() (uint16, error)
	CheckMachineID func(uint16) bool
}

// Sonyflake is a distributed unique ID generator.
type Sonyflake struct {
	__obf_2a77795831a05760 *sync.Mutex
	__obf_f21b2a32022eb176 int64
	__obf_f6d5ee9f6e5d9e13 int64
	__obf_f06c729260d81c2a uint16
	__obf_4d1c08ffc8910429 uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_b3415558c7d02e0c = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_fde2be03904403c6 Settings) (*Sonyflake, error) {
	if __obf_fde2be03904403c6.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}

	__obf_f98895c99a53c78c := new(Sonyflake)
	__obf_f98895c99a53c78c.__obf_2a77795831a05760 = new(sync.Mutex)
	__obf_f98895c99a53c78c.__obf_f06c729260d81c2a = uint16(1<<BitLenSequence - 1)

	if __obf_fde2be03904403c6.StartTime.IsZero() {
		__obf_f98895c99a53c78c.__obf_f21b2a32022eb176 = __obf_9bcec0b049916358(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_f98895c99a53c78c.__obf_f21b2a32022eb176 = __obf_9bcec0b049916358(__obf_fde2be03904403c6.StartTime)
	}

	var __obf_3f3f828d9e46d714 error
	if __obf_fde2be03904403c6.MachineID == nil {
		__obf_f98895c99a53c78c.__obf_4d1c08ffc8910429, __obf_3f3f828d9e46d714 = __obf_ea31609dad392b47(__obf_b3415558c7d02e0c)
	} else {
		__obf_f98895c99a53c78c.__obf_4d1c08ffc8910429, __obf_3f3f828d9e46d714 = __obf_fde2be03904403c6.MachineID()
	}
	if __obf_3f3f828d9e46d714 != nil {
		return nil, __obf_3f3f828d9e46d714
	}

	if __obf_fde2be03904403c6.CheckMachineID != nil && !__obf_fde2be03904403c6.CheckMachineID(__obf_f98895c99a53c78c.__obf_4d1c08ffc8910429) {
		return nil, ErrInvalidMachineID
	}

	return __obf_f98895c99a53c78c, nil
}

// NewSonyflake returns a new Sonyflake configured with the given Settings.
// NewSonyflake returns nil in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
// func NewSonyflake(st Settings) *Sonyflake {
// 	sf, _ := New(st)
// 	return sf
// }

// NextID generates a next unique ID.
// After the Sonyflake time overflows, NextID returns an error.
func (__obf_f98895c99a53c78c *Sonyflake) NextID() (string, error) {
	const __obf_3c6ae78a37915017 = uint16(1<<BitLenSequence - 1)

	__obf_f98895c99a53c78c.__obf_2a77795831a05760.Lock()
	defer __obf_f98895c99a53c78c.__obf_2a77795831a05760.Unlock()

	__obf_836cc6c5c7295bd7 := __obf_e77edc3b28908c56(__obf_f98895c99a53c78c.__obf_f21b2a32022eb176)
	if __obf_f98895c99a53c78c.__obf_f6d5ee9f6e5d9e13 < __obf_836cc6c5c7295bd7 {
		__obf_f98895c99a53c78c.__obf_f6d5ee9f6e5d9e13 = __obf_836cc6c5c7295bd7
		__obf_f98895c99a53c78c.__obf_f06c729260d81c2a = 0
	} else { // sf.elapsedTime >= current
		__obf_f98895c99a53c78c.__obf_f06c729260d81c2a = (__obf_f98895c99a53c78c.__obf_f06c729260d81c2a + 1) & __obf_3c6ae78a37915017
		if __obf_f98895c99a53c78c.__obf_f06c729260d81c2a == 0 {
			__obf_f98895c99a53c78c.__obf_f6d5ee9f6e5d9e13++
			__obf_39c91f0f3ebde38c := __obf_f98895c99a53c78c.__obf_f6d5ee9f6e5d9e13 - __obf_836cc6c5c7295bd7
			time.Sleep(__obf_75ce1e7d327c099e((__obf_39c91f0f3ebde38c)))
		}
	}

	return __obf_f98895c99a53c78c.__obf_babd7f936b625c1c()
}

const __obf_91623586805fb9e3 = 1e7 // nsec, i.e. 10 msec

func __obf_9bcec0b049916358(__obf_680f17a8e930eb64 time.Time) int64 {
	return __obf_680f17a8e930eb64.UTC().UnixNano() / __obf_91623586805fb9e3
}

func __obf_e77edc3b28908c56(__obf_f21b2a32022eb176 int64) int64 {
	return __obf_9bcec0b049916358(time.Now()) - __obf_f21b2a32022eb176
}

func __obf_75ce1e7d327c099e(__obf_39c91f0f3ebde38c int64) time.Duration {
	return time.Duration(__obf_39c91f0f3ebde38c*__obf_91623586805fb9e3) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_91623586805fb9e3)
}

func (__obf_f98895c99a53c78c *Sonyflake) __obf_babd7f936b625c1c() (string, error) {
	if __obf_f98895c99a53c78c.__obf_f6d5ee9f6e5d9e13 >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_f98895c99a53c78c.__obf_f6d5ee9f6e5d9e13)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_f98895c99a53c78c.__obf_f06c729260d81c2a)<<BitLenMachineID|
		uint64(__obf_f98895c99a53c78c.__obf_4d1c08ffc8910429)), nil
}

func __obf_4164cfc1ef9b086c(__obf_cc1a900781f051a4 InterfaceAddrs) (net.IP, error) {
	__obf_4b14adae374e468b, __obf_3f3f828d9e46d714 := __obf_cc1a900781f051a4()
	if __obf_3f3f828d9e46d714 != nil {
		return nil, __obf_3f3f828d9e46d714
	}

	for _, __obf_e5925bb152d76493 := range __obf_4b14adae374e468b {
		__obf_11a3d16d31a23dc3, __obf_4f3a1b70936ef093 := __obf_e5925bb152d76493.(*net.IPNet)
		if !__obf_4f3a1b70936ef093 || __obf_11a3d16d31a23dc3.IP.IsLoopback() {
			continue
		}

		__obf_7bf263a4fe6e7532 := __obf_11a3d16d31a23dc3.IP.To4()
		if __obf_999531fe877a2c45(__obf_7bf263a4fe6e7532) {
			return __obf_7bf263a4fe6e7532, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_999531fe877a2c45(__obf_7bf263a4fe6e7532 net.IP) bool {
	return __obf_7bf263a4fe6e7532 != nil &&
		(__obf_7bf263a4fe6e7532[0] == 10 || __obf_7bf263a4fe6e7532[0] == 172 && (__obf_7bf263a4fe6e7532[1] >= 16 && __obf_7bf263a4fe6e7532[1] < 32) || __obf_7bf263a4fe6e7532[0] == 192 && __obf_7bf263a4fe6e7532[1] == 168)
}

func __obf_ea31609dad392b47(__obf_cc1a900781f051a4 InterfaceAddrs) (uint16, error) {
	__obf_7bf263a4fe6e7532, __obf_3f3f828d9e46d714 := __obf_4164cfc1ef9b086c(__obf_cc1a900781f051a4)
	if __obf_3f3f828d9e46d714 != nil {
		return 0, __obf_3f3f828d9e46d714
	}

	return uint16(__obf_7bf263a4fe6e7532[2])<<8 + uint16(__obf_7bf263a4fe6e7532[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_6cfbdbb264317300 uint64) time.Duration {
	return time.Duration(__obf_f6d5ee9f6e5d9e13(__obf_6cfbdbb264317300) * __obf_91623586805fb9e3)
}

func __obf_f6d5ee9f6e5d9e13(__obf_6cfbdbb264317300 uint64) uint64 {
	return __obf_6cfbdbb264317300 >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_6cfbdbb264317300 uint64) uint64 {
	const __obf_3c6ae78a37915017 = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_6cfbdbb264317300 & __obf_3c6ae78a37915017 >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_6cfbdbb264317300 uint64) uint64 {
	const __obf_72e8049067844a62 = uint64(1<<BitLenMachineID - 1)
	return __obf_6cfbdbb264317300 & __obf_72e8049067844a62
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_6cfbdbb264317300 uint64) map[string]uint64 {
	__obf_a8aeb649935f9ce2 := __obf_6cfbdbb264317300 >> 63
	time := __obf_f6d5ee9f6e5d9e13(__obf_6cfbdbb264317300)
	__obf_f06c729260d81c2a := SequenceNumber(__obf_6cfbdbb264317300)
	__obf_4d1c08ffc8910429 := MachineID(__obf_6cfbdbb264317300)
	return map[string]uint64{
		"id":         __obf_6cfbdbb264317300,
		"msb":        __obf_a8aeb649935f9ce2,
		"time":       time,
		"sequence":   __obf_f06c729260d81c2a,
		"machine-id": __obf_4d1c08ffc8910429,
	}
}
