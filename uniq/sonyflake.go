// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_07f0876faa0cf68e

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
	__obf_5510d5383b3817f2 *sync.Mutex
	__obf_ecebf086735260cb int64
	__obf_f1683acb750e61fd int64
	__obf_dd9e42b79b3ea456 uint16
	__obf_19f0013c66fe68c2 uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_2eca3c0efc069313 = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_350f7ec9f530868a Settings) (*Sonyflake, error) {
	if __obf_350f7ec9f530868a.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}
	__obf_5c1217b1c2c9bef1 := new(Sonyflake)
	__obf_5c1217b1c2c9bef1.__obf_5510d5383b3817f2 = new(sync.Mutex)
	__obf_5c1217b1c2c9bef1.__obf_dd9e42b79b3ea456 = uint16(1<<BitLenSequence - 1)

	if __obf_350f7ec9f530868a.StartTime.IsZero() {
		__obf_5c1217b1c2c9bef1.__obf_ecebf086735260cb = __obf_7a7a0d8c9e5f5340(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_5c1217b1c2c9bef1.__obf_ecebf086735260cb = __obf_7a7a0d8c9e5f5340(__obf_350f7ec9f530868a.StartTime)
	}

	var __obf_116d5a689b79a523 error
	if __obf_350f7ec9f530868a.MachineID == nil {
		__obf_5c1217b1c2c9bef1.__obf_19f0013c66fe68c2, __obf_116d5a689b79a523 = __obf_4b65cce0206d235e(__obf_2eca3c0efc069313)
	} else {
		__obf_5c1217b1c2c9bef1.__obf_19f0013c66fe68c2, __obf_116d5a689b79a523 = __obf_350f7ec9f530868a.MachineID()
	}
	if __obf_116d5a689b79a523 != nil {
		return nil, __obf_116d5a689b79a523
	}

	if __obf_350f7ec9f530868a.CheckMachineID != nil && !__obf_350f7ec9f530868a.CheckMachineID(__obf_5c1217b1c2c9bef1.__obf_19f0013c66fe68c2) {
		return nil, ErrInvalidMachineID
	}

	return __obf_5c1217b1c2c9bef1, nil
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
func (__obf_5c1217b1c2c9bef1 *Sonyflake) NextID() (string, error) {
	const __obf_47657b9abc1f81fc = uint16(1<<BitLenSequence - 1)
	__obf_5c1217b1c2c9bef1.__obf_5510d5383b3817f2.
		Lock()
	defer __obf_5c1217b1c2c9bef1.__obf_5510d5383b3817f2.Unlock()
	__obf_1a7aacc55917eb0d := __obf_81d07c789ae1135f(__obf_5c1217b1c2c9bef1.__obf_ecebf086735260cb)
	if __obf_5c1217b1c2c9bef1.__obf_f1683acb750e61fd < __obf_1a7aacc55917eb0d {
		__obf_5c1217b1c2c9bef1.__obf_f1683acb750e61fd = __obf_1a7aacc55917eb0d
		__obf_5c1217b1c2c9bef1.

			// sf.elapsedTime >= current
			__obf_dd9e42b79b3ea456 = 0
	} else {
		__obf_5c1217b1c2c9bef1.__obf_dd9e42b79b3ea456 = (__obf_5c1217b1c2c9bef1.__obf_dd9e42b79b3ea456 + 1) & __obf_47657b9abc1f81fc
		if __obf_5c1217b1c2c9bef1.__obf_dd9e42b79b3ea456 == 0 {
			__obf_5c1217b1c2c9bef1.__obf_f1683acb750e61fd++
			__obf_c8cc602db905ffce := __obf_5c1217b1c2c9bef1.__obf_f1683acb750e61fd - __obf_1a7aacc55917eb0d
			time.Sleep(__obf_5ade0e22ae342c86((__obf_c8cc602db905ffce)))
		}
	}

	return __obf_5c1217b1c2c9bef1.__obf_718b38b06cc64baf()
}

const __obf_94c85fca9c150861 = 1e7 // nsec, i.e. 10 msec

func __obf_7a7a0d8c9e5f5340(__obf_13047ef7aa466d6e time.Time) int64 {
	return __obf_13047ef7aa466d6e.UTC().UnixNano() / __obf_94c85fca9c150861
}

func __obf_81d07c789ae1135f(__obf_ecebf086735260cb int64) int64 {
	return __obf_7a7a0d8c9e5f5340(time.Now()) - __obf_ecebf086735260cb
}

func __obf_5ade0e22ae342c86(__obf_c8cc602db905ffce int64) time.Duration {
	return time.Duration(__obf_c8cc602db905ffce*__obf_94c85fca9c150861) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_94c85fca9c150861)
}

func (__obf_5c1217b1c2c9bef1 *Sonyflake) __obf_718b38b06cc64baf() (string, error) {
	if __obf_5c1217b1c2c9bef1.__obf_f1683acb750e61fd >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_5c1217b1c2c9bef1.__obf_f1683acb750e61fd)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_5c1217b1c2c9bef1.__obf_dd9e42b79b3ea456)<<BitLenMachineID|
		uint64(__obf_5c1217b1c2c9bef1.__obf_19f0013c66fe68c2)), nil
}

func __obf_fdd1b2a2f6a7b71d(__obf_9ca35b090bb5a88e InterfaceAddrs) (net.IP, error) {
	__obf_5fae203cba181304, __obf_116d5a689b79a523 := __obf_9ca35b090bb5a88e()
	if __obf_116d5a689b79a523 != nil {
		return nil, __obf_116d5a689b79a523
	}

	for _, __obf_63951c6162ec7efe := range __obf_5fae203cba181304 {
		__obf_3a3297f11894932b, __obf_820762e8a994bf2b := __obf_63951c6162ec7efe.(*net.IPNet)
		if !__obf_820762e8a994bf2b || __obf_3a3297f11894932b.IP.IsLoopback() {
			continue
		}
		__obf_0beb1b759290e625 := __obf_3a3297f11894932b.IP.To4()
		if __obf_ad8e3d029757c1f8(__obf_0beb1b759290e625) {
			return __obf_0beb1b759290e625, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_ad8e3d029757c1f8(__obf_0beb1b759290e625 net.IP) bool {
	return __obf_0beb1b759290e625 != nil &&
		(__obf_0beb1b759290e625[0] == 10 || __obf_0beb1b759290e625[0] == 172 && (__obf_0beb1b759290e625[1] >= 16 && __obf_0beb1b759290e625[1] < 32) || __obf_0beb1b759290e625[0] == 192 && __obf_0beb1b759290e625[1] == 168)
}

func __obf_4b65cce0206d235e(__obf_9ca35b090bb5a88e InterfaceAddrs) (uint16, error) {
	__obf_0beb1b759290e625, __obf_116d5a689b79a523 := __obf_fdd1b2a2f6a7b71d(__obf_9ca35b090bb5a88e)
	if __obf_116d5a689b79a523 != nil {
		return 0, __obf_116d5a689b79a523
	}

	return uint16(__obf_0beb1b759290e625[2])<<8 + uint16(__obf_0beb1b759290e625[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_28d51b2675aee39d uint64) time.Duration {
	return time.Duration(__obf_f1683acb750e61fd(__obf_28d51b2675aee39d) * __obf_94c85fca9c150861)
}

func __obf_f1683acb750e61fd(__obf_28d51b2675aee39d uint64) uint64 {
	return __obf_28d51b2675aee39d >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_28d51b2675aee39d uint64) uint64 {
	const __obf_47657b9abc1f81fc = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_28d51b2675aee39d & __obf_47657b9abc1f81fc >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_28d51b2675aee39d uint64) uint64 {
	const __obf_0b15c53bd8aaedaa = uint64(1<<BitLenMachineID - 1)
	return __obf_28d51b2675aee39d & __obf_0b15c53bd8aaedaa
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_28d51b2675aee39d uint64) map[string]uint64 {
	__obf_90aba5c355ff1317 := __obf_28d51b2675aee39d >> 63
	time := __obf_f1683acb750e61fd(__obf_28d51b2675aee39d)
	__obf_dd9e42b79b3ea456 := SequenceNumber(__obf_28d51b2675aee39d)
	__obf_19f0013c66fe68c2 := MachineID(__obf_28d51b2675aee39d)
	return map[string]uint64{
		"id":         __obf_28d51b2675aee39d,
		"msb":        __obf_90aba5c355ff1317,
		"time":       time,
		"sequence":   __obf_dd9e42b79b3ea456,
		"machine-id": __obf_19f0013c66fe68c2,
	}
}
