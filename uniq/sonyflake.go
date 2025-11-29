// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_2f51f7d26a2bcdf8

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
	__obf_cc0198966ba2d14a *sync.Mutex
	__obf_bbef56ae138e51f0 int64
	__obf_6fa3133ce482ee57 int64
	__obf_692ac7e4829ff716 uint16
	__obf_db1e1151c51a61a0 uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_22d3dfe2cd174acd = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_c01c48bd849199f7 Settings) (*Sonyflake, error) {
	if __obf_c01c48bd849199f7.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}
	__obf_75de03075dcea8a5 := new(Sonyflake)
	__obf_75de03075dcea8a5.__obf_cc0198966ba2d14a = new(sync.Mutex)
	__obf_75de03075dcea8a5.__obf_692ac7e4829ff716 = uint16(1<<BitLenSequence - 1)

	if __obf_c01c48bd849199f7.StartTime.IsZero() {
		__obf_75de03075dcea8a5.__obf_bbef56ae138e51f0 = __obf_7bd7dab331651d0e(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_75de03075dcea8a5.__obf_bbef56ae138e51f0 = __obf_7bd7dab331651d0e(__obf_c01c48bd849199f7.StartTime)
	}

	var __obf_37e8e388755917c0 error
	if __obf_c01c48bd849199f7.MachineID == nil {
		__obf_75de03075dcea8a5.__obf_db1e1151c51a61a0, __obf_37e8e388755917c0 = __obf_c234d3b31e286643(__obf_22d3dfe2cd174acd)
	} else {
		__obf_75de03075dcea8a5.__obf_db1e1151c51a61a0, __obf_37e8e388755917c0 = __obf_c01c48bd849199f7.MachineID()
	}
	if __obf_37e8e388755917c0 != nil {
		return nil, __obf_37e8e388755917c0
	}

	if __obf_c01c48bd849199f7.CheckMachineID != nil && !__obf_c01c48bd849199f7.CheckMachineID(__obf_75de03075dcea8a5.__obf_db1e1151c51a61a0) {
		return nil, ErrInvalidMachineID
	}

	return __obf_75de03075dcea8a5, nil
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
func (__obf_75de03075dcea8a5 *Sonyflake) NextID() (string, error) {
	const __obf_5214c014ce657c36 = uint16(1<<BitLenSequence - 1)
	__obf_75de03075dcea8a5.__obf_cc0198966ba2d14a.
		Lock()
	defer __obf_75de03075dcea8a5.__obf_cc0198966ba2d14a.Unlock()
	__obf_aee10a7ede9ebb37 := __obf_b4d0522d76bfe039(__obf_75de03075dcea8a5.__obf_bbef56ae138e51f0)
	if __obf_75de03075dcea8a5.__obf_6fa3133ce482ee57 < __obf_aee10a7ede9ebb37 {
		__obf_75de03075dcea8a5.__obf_6fa3133ce482ee57 = __obf_aee10a7ede9ebb37
		__obf_75de03075dcea8a5.

			// sf.elapsedTime >= current
			__obf_692ac7e4829ff716 = 0
	} else {
		__obf_75de03075dcea8a5.__obf_692ac7e4829ff716 = (__obf_75de03075dcea8a5.__obf_692ac7e4829ff716 + 1) & __obf_5214c014ce657c36
		if __obf_75de03075dcea8a5.__obf_692ac7e4829ff716 == 0 {
			__obf_75de03075dcea8a5.__obf_6fa3133ce482ee57++
			__obf_a09ddbea1f9d4ee9 := __obf_75de03075dcea8a5.__obf_6fa3133ce482ee57 - __obf_aee10a7ede9ebb37
			time.Sleep(__obf_ad576fb9be5fccf8((__obf_a09ddbea1f9d4ee9)))
		}
	}

	return __obf_75de03075dcea8a5.__obf_68de03ef47887d8b()
}

const __obf_6ffcd44dd2119f0c = 1e7 // nsec, i.e. 10 msec

func __obf_7bd7dab331651d0e(__obf_61139133e991449d time.Time) int64 {
	return __obf_61139133e991449d.UTC().UnixNano() / __obf_6ffcd44dd2119f0c
}

func __obf_b4d0522d76bfe039(__obf_bbef56ae138e51f0 int64) int64 {
	return __obf_7bd7dab331651d0e(time.Now()) - __obf_bbef56ae138e51f0
}

func __obf_ad576fb9be5fccf8(__obf_a09ddbea1f9d4ee9 int64) time.Duration {
	return time.Duration(__obf_a09ddbea1f9d4ee9*__obf_6ffcd44dd2119f0c) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_6ffcd44dd2119f0c)
}

func (__obf_75de03075dcea8a5 *Sonyflake) __obf_68de03ef47887d8b() (string, error) {
	if __obf_75de03075dcea8a5.__obf_6fa3133ce482ee57 >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_75de03075dcea8a5.__obf_6fa3133ce482ee57)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_75de03075dcea8a5.__obf_692ac7e4829ff716)<<BitLenMachineID|
		uint64(__obf_75de03075dcea8a5.__obf_db1e1151c51a61a0)), nil
}

func __obf_307e7b7269f6950c(__obf_316240db66c76575 InterfaceAddrs) (net.IP, error) {
	__obf_8ef78ed1dff688a7, __obf_37e8e388755917c0 := __obf_316240db66c76575()
	if __obf_37e8e388755917c0 != nil {
		return nil, __obf_37e8e388755917c0
	}

	for _, __obf_1e177394b5c6680f := range __obf_8ef78ed1dff688a7 {
		__obf_40f52ea4d1f02c80, __obf_149b4fd56461a8a0 := __obf_1e177394b5c6680f.(*net.IPNet)
		if !__obf_149b4fd56461a8a0 || __obf_40f52ea4d1f02c80.IP.IsLoopback() {
			continue
		}
		__obf_a39bde65abb63d08 := __obf_40f52ea4d1f02c80.IP.To4()
		if __obf_1d1a654873246c7d(__obf_a39bde65abb63d08) {
			return __obf_a39bde65abb63d08, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_1d1a654873246c7d(__obf_a39bde65abb63d08 net.IP) bool {
	return __obf_a39bde65abb63d08 != nil &&
		(__obf_a39bde65abb63d08[0] == 10 || __obf_a39bde65abb63d08[0] == 172 && (__obf_a39bde65abb63d08[1] >= 16 && __obf_a39bde65abb63d08[1] < 32) || __obf_a39bde65abb63d08[0] == 192 && __obf_a39bde65abb63d08[1] == 168)
}

func __obf_c234d3b31e286643(__obf_316240db66c76575 InterfaceAddrs) (uint16, error) {
	__obf_a39bde65abb63d08, __obf_37e8e388755917c0 := __obf_307e7b7269f6950c(__obf_316240db66c76575)
	if __obf_37e8e388755917c0 != nil {
		return 0, __obf_37e8e388755917c0
	}

	return uint16(__obf_a39bde65abb63d08[2])<<8 + uint16(__obf_a39bde65abb63d08[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_8b246d7b3659af37 uint64) time.Duration {
	return time.Duration(__obf_6fa3133ce482ee57(__obf_8b246d7b3659af37) * __obf_6ffcd44dd2119f0c)
}

func __obf_6fa3133ce482ee57(__obf_8b246d7b3659af37 uint64) uint64 {
	return __obf_8b246d7b3659af37 >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_8b246d7b3659af37 uint64) uint64 {
	const __obf_5214c014ce657c36 = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_8b246d7b3659af37 & __obf_5214c014ce657c36 >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_8b246d7b3659af37 uint64) uint64 {
	const __obf_a87b5be844a6a85b = uint64(1<<BitLenMachineID - 1)
	return __obf_8b246d7b3659af37 & __obf_a87b5be844a6a85b
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_8b246d7b3659af37 uint64) map[string]uint64 {
	__obf_d535b42c78041d18 := __obf_8b246d7b3659af37 >> 63
	time := __obf_6fa3133ce482ee57(__obf_8b246d7b3659af37)
	__obf_692ac7e4829ff716 := SequenceNumber(__obf_8b246d7b3659af37)
	__obf_db1e1151c51a61a0 := MachineID(__obf_8b246d7b3659af37)
	return map[string]uint64{
		"id":         __obf_8b246d7b3659af37,
		"msb":        __obf_d535b42c78041d18,
		"time":       time,
		"sequence":   __obf_692ac7e4829ff716,
		"machine-id": __obf_db1e1151c51a61a0,
	}
}
