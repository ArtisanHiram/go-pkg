// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_a51a64e21f311927

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
	__obf_e99a29e636f342d2 *sync.Mutex
	__obf_2986d383f596e408 int64
	__obf_5c366c6a8c901659 int64
	__obf_97426f7ecb5d7ca6 uint16
	__obf_5ea4c9790d43f858 uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_9ce1cd479bcd9330 = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_dbec924b7d374c89 Settings) (*Sonyflake, error) {
	if __obf_dbec924b7d374c89.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}
	__obf_6f8fc99019dffd83 := new(Sonyflake)
	__obf_6f8fc99019dffd83.__obf_e99a29e636f342d2 = new(sync.Mutex)
	__obf_6f8fc99019dffd83.__obf_97426f7ecb5d7ca6 = uint16(1<<BitLenSequence - 1)

	if __obf_dbec924b7d374c89.StartTime.IsZero() {
		__obf_6f8fc99019dffd83.__obf_2986d383f596e408 = __obf_6d304f8b393d3200(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_6f8fc99019dffd83.__obf_2986d383f596e408 = __obf_6d304f8b393d3200(__obf_dbec924b7d374c89.StartTime)
	}

	var __obf_886df8acc7f50b76 error
	if __obf_dbec924b7d374c89.MachineID == nil {
		__obf_6f8fc99019dffd83.__obf_5ea4c9790d43f858, __obf_886df8acc7f50b76 = __obf_4dd11ae87126fc3a(__obf_9ce1cd479bcd9330)
	} else {
		__obf_6f8fc99019dffd83.__obf_5ea4c9790d43f858, __obf_886df8acc7f50b76 = __obf_dbec924b7d374c89.MachineID()
	}
	if __obf_886df8acc7f50b76 != nil {
		return nil, __obf_886df8acc7f50b76
	}

	if __obf_dbec924b7d374c89.CheckMachineID != nil && !__obf_dbec924b7d374c89.CheckMachineID(__obf_6f8fc99019dffd83.__obf_5ea4c9790d43f858) {
		return nil, ErrInvalidMachineID
	}

	return __obf_6f8fc99019dffd83, nil
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
func (__obf_6f8fc99019dffd83 *Sonyflake) NextID() (string, error) {
	const __obf_2e5191743c9a4bb1 = uint16(1<<BitLenSequence - 1)
	__obf_6f8fc99019dffd83.__obf_e99a29e636f342d2.
		Lock()
	defer __obf_6f8fc99019dffd83.__obf_e99a29e636f342d2.Unlock()
	__obf_06a075bfa90b1eff := __obf_be0e0229b8f012f9(__obf_6f8fc99019dffd83.__obf_2986d383f596e408)
	if __obf_6f8fc99019dffd83.__obf_5c366c6a8c901659 < __obf_06a075bfa90b1eff {
		__obf_6f8fc99019dffd83.__obf_5c366c6a8c901659 = __obf_06a075bfa90b1eff
		__obf_6f8fc99019dffd83.

			// sf.elapsedTime >= current
			__obf_97426f7ecb5d7ca6 = 0
	} else {
		__obf_6f8fc99019dffd83.__obf_97426f7ecb5d7ca6 = (__obf_6f8fc99019dffd83.__obf_97426f7ecb5d7ca6 + 1) & __obf_2e5191743c9a4bb1
		if __obf_6f8fc99019dffd83.__obf_97426f7ecb5d7ca6 == 0 {
			__obf_6f8fc99019dffd83.__obf_5c366c6a8c901659++
			__obf_bd16f8f55a057103 := __obf_6f8fc99019dffd83.__obf_5c366c6a8c901659 - __obf_06a075bfa90b1eff
			time.Sleep(__obf_21c57d5d22111af7((__obf_bd16f8f55a057103)))
		}
	}

	return __obf_6f8fc99019dffd83.__obf_1ee32388ea6f478f()
}

const __obf_a487d140d3d61c1f = 1e7 // nsec, i.e. 10 msec

func __obf_6d304f8b393d3200(__obf_c5969bc7b424c7a8 time.Time) int64 {
	return __obf_c5969bc7b424c7a8.UTC().UnixNano() / __obf_a487d140d3d61c1f
}

func __obf_be0e0229b8f012f9(__obf_2986d383f596e408 int64) int64 {
	return __obf_6d304f8b393d3200(time.Now()) - __obf_2986d383f596e408
}

func __obf_21c57d5d22111af7(__obf_bd16f8f55a057103 int64) time.Duration {
	return time.Duration(__obf_bd16f8f55a057103*__obf_a487d140d3d61c1f) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_a487d140d3d61c1f)
}

func (__obf_6f8fc99019dffd83 *Sonyflake) __obf_1ee32388ea6f478f() (string, error) {
	if __obf_6f8fc99019dffd83.__obf_5c366c6a8c901659 >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_6f8fc99019dffd83.__obf_5c366c6a8c901659)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_6f8fc99019dffd83.__obf_97426f7ecb5d7ca6)<<BitLenMachineID|
		uint64(__obf_6f8fc99019dffd83.__obf_5ea4c9790d43f858)), nil
}

func __obf_46d025c149aa3839(__obf_b17342c1de076612 InterfaceAddrs) (net.IP, error) {
	__obf_2b10b721c3035ad7, __obf_886df8acc7f50b76 := __obf_b17342c1de076612()
	if __obf_886df8acc7f50b76 != nil {
		return nil, __obf_886df8acc7f50b76
	}

	for _, __obf_ea86e12fa8969a56 := range __obf_2b10b721c3035ad7 {
		__obf_0a906bc5c9d558fe, __obf_80eec96aa83c2efa := __obf_ea86e12fa8969a56.(*net.IPNet)
		if !__obf_80eec96aa83c2efa || __obf_0a906bc5c9d558fe.IP.IsLoopback() {
			continue
		}
		__obf_9646fce6b17fd688 := __obf_0a906bc5c9d558fe.IP.To4()
		if __obf_6cd4301a0d851c60(__obf_9646fce6b17fd688) {
			return __obf_9646fce6b17fd688, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_6cd4301a0d851c60(__obf_9646fce6b17fd688 net.IP) bool {
	return __obf_9646fce6b17fd688 != nil &&
		(__obf_9646fce6b17fd688[0] == 10 || __obf_9646fce6b17fd688[0] == 172 && (__obf_9646fce6b17fd688[1] >= 16 && __obf_9646fce6b17fd688[1] < 32) || __obf_9646fce6b17fd688[0] == 192 && __obf_9646fce6b17fd688[1] == 168)
}

func __obf_4dd11ae87126fc3a(__obf_b17342c1de076612 InterfaceAddrs) (uint16, error) {
	__obf_9646fce6b17fd688, __obf_886df8acc7f50b76 := __obf_46d025c149aa3839(__obf_b17342c1de076612)
	if __obf_886df8acc7f50b76 != nil {
		return 0, __obf_886df8acc7f50b76
	}

	return uint16(__obf_9646fce6b17fd688[2])<<8 + uint16(__obf_9646fce6b17fd688[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_4c112818117e5211 uint64) time.Duration {
	return time.Duration(__obf_5c366c6a8c901659(__obf_4c112818117e5211) * __obf_a487d140d3d61c1f)
}

func __obf_5c366c6a8c901659(__obf_4c112818117e5211 uint64) uint64 {
	return __obf_4c112818117e5211 >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_4c112818117e5211 uint64) uint64 {
	const __obf_2e5191743c9a4bb1 = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_4c112818117e5211 & __obf_2e5191743c9a4bb1 >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_4c112818117e5211 uint64) uint64 {
	const __obf_09f3cd7197eabaf4 = uint64(1<<BitLenMachineID - 1)
	return __obf_4c112818117e5211 & __obf_09f3cd7197eabaf4
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_4c112818117e5211 uint64) map[string]uint64 {
	__obf_45d65e9e48357fef := __obf_4c112818117e5211 >> 63
	time := __obf_5c366c6a8c901659(__obf_4c112818117e5211)
	__obf_97426f7ecb5d7ca6 := SequenceNumber(__obf_4c112818117e5211)
	__obf_5ea4c9790d43f858 := MachineID(__obf_4c112818117e5211)
	return map[string]uint64{
		"id":         __obf_4c112818117e5211,
		"msb":        __obf_45d65e9e48357fef,
		"time":       time,
		"sequence":   __obf_97426f7ecb5d7ca6,
		"machine-id": __obf_5ea4c9790d43f858,
	}
}
