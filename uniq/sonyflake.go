// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_3747a7e09ff475ee

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
	__obf_0f8a4579ce651036 *sync.Mutex
	__obf_2e783d19789e8927 int64
	__obf_a2778e244d60c065 int64
	__obf_32d285c02de19ff8 uint16
	__obf_825d516c59f4195a uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_0e22c6d498150d08 = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_7b794e5d598e8593 Settings) (*Sonyflake, error) {
	if __obf_7b794e5d598e8593.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}
	__obf_a0f1aecfa77b8f05 := new(Sonyflake)
	__obf_a0f1aecfa77b8f05.__obf_0f8a4579ce651036 = new(sync.Mutex)
	__obf_a0f1aecfa77b8f05.__obf_32d285c02de19ff8 = uint16(1<<BitLenSequence - 1)

	if __obf_7b794e5d598e8593.StartTime.IsZero() {
		__obf_a0f1aecfa77b8f05.__obf_2e783d19789e8927 = __obf_a5a2229875b333b2(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_a0f1aecfa77b8f05.__obf_2e783d19789e8927 = __obf_a5a2229875b333b2(__obf_7b794e5d598e8593.StartTime)
	}

	var __obf_57cef2f010b4ce9a error
	if __obf_7b794e5d598e8593.MachineID == nil {
		__obf_a0f1aecfa77b8f05.__obf_825d516c59f4195a, __obf_57cef2f010b4ce9a = __obf_148a164a645089a6(__obf_0e22c6d498150d08)
	} else {
		__obf_a0f1aecfa77b8f05.__obf_825d516c59f4195a, __obf_57cef2f010b4ce9a = __obf_7b794e5d598e8593.MachineID()
	}
	if __obf_57cef2f010b4ce9a != nil {
		return nil, __obf_57cef2f010b4ce9a
	}

	if __obf_7b794e5d598e8593.CheckMachineID != nil && !__obf_7b794e5d598e8593.CheckMachineID(__obf_a0f1aecfa77b8f05.__obf_825d516c59f4195a) {
		return nil, ErrInvalidMachineID
	}

	return __obf_a0f1aecfa77b8f05, nil
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
func (__obf_a0f1aecfa77b8f05 *Sonyflake) NextID() (string, error) {
	const __obf_39a40eb8f9f68696 = uint16(1<<BitLenSequence - 1)
	__obf_a0f1aecfa77b8f05.__obf_0f8a4579ce651036.
		Lock()
	defer __obf_a0f1aecfa77b8f05.__obf_0f8a4579ce651036.Unlock()
	__obf_ea96c3b58a417c7d := __obf_8f4064454d16d18a(__obf_a0f1aecfa77b8f05.__obf_2e783d19789e8927)
	if __obf_a0f1aecfa77b8f05.__obf_a2778e244d60c065 < __obf_ea96c3b58a417c7d {
		__obf_a0f1aecfa77b8f05.__obf_a2778e244d60c065 = __obf_ea96c3b58a417c7d
		__obf_a0f1aecfa77b8f05.

			// sf.elapsedTime >= current
			__obf_32d285c02de19ff8 = 0
	} else {
		__obf_a0f1aecfa77b8f05.__obf_32d285c02de19ff8 = (__obf_a0f1aecfa77b8f05.__obf_32d285c02de19ff8 + 1) & __obf_39a40eb8f9f68696
		if __obf_a0f1aecfa77b8f05.__obf_32d285c02de19ff8 == 0 {
			__obf_a0f1aecfa77b8f05.__obf_a2778e244d60c065++
			__obf_14275282e4c90688 := __obf_a0f1aecfa77b8f05.__obf_a2778e244d60c065 - __obf_ea96c3b58a417c7d
			time.Sleep(__obf_6ff83f94e804510a((__obf_14275282e4c90688)))
		}
	}

	return __obf_a0f1aecfa77b8f05.__obf_c2cc79b71132a5df()
}

const __obf_2f8ad3861a12d75e = 1e7 // nsec, i.e. 10 msec

func __obf_a5a2229875b333b2(__obf_b5849d1d191d006d time.Time) int64 {
	return __obf_b5849d1d191d006d.UTC().UnixNano() / __obf_2f8ad3861a12d75e
}

func __obf_8f4064454d16d18a(__obf_2e783d19789e8927 int64) int64 {
	return __obf_a5a2229875b333b2(time.Now()) - __obf_2e783d19789e8927
}

func __obf_6ff83f94e804510a(__obf_14275282e4c90688 int64) time.Duration {
	return time.Duration(__obf_14275282e4c90688*__obf_2f8ad3861a12d75e) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_2f8ad3861a12d75e)
}

func (__obf_a0f1aecfa77b8f05 *Sonyflake) __obf_c2cc79b71132a5df() (string, error) {
	if __obf_a0f1aecfa77b8f05.__obf_a2778e244d60c065 >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_a0f1aecfa77b8f05.__obf_a2778e244d60c065)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_a0f1aecfa77b8f05.__obf_32d285c02de19ff8)<<BitLenMachineID|
		uint64(__obf_a0f1aecfa77b8f05.__obf_825d516c59f4195a)), nil
}

func __obf_dc9f389f94af361f(__obf_876f04b960c8f60c InterfaceAddrs) (net.IP, error) {
	__obf_ce9c4d540cef3f19, __obf_57cef2f010b4ce9a := __obf_876f04b960c8f60c()
	if __obf_57cef2f010b4ce9a != nil {
		return nil, __obf_57cef2f010b4ce9a
	}

	for _, __obf_2ea177ce50a9da52 := range __obf_ce9c4d540cef3f19 {
		__obf_6c69ece45c49c4a9, __obf_0941b7586b921890 := __obf_2ea177ce50a9da52.(*net.IPNet)
		if !__obf_0941b7586b921890 || __obf_6c69ece45c49c4a9.IP.IsLoopback() {
			continue
		}
		__obf_fc2de974197327a1 := __obf_6c69ece45c49c4a9.IP.To4()
		if __obf_c370799b2a669262(__obf_fc2de974197327a1) {
			return __obf_fc2de974197327a1, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_c370799b2a669262(__obf_fc2de974197327a1 net.IP) bool {
	return __obf_fc2de974197327a1 != nil &&
		(__obf_fc2de974197327a1[0] == 10 || __obf_fc2de974197327a1[0] == 172 && (__obf_fc2de974197327a1[1] >= 16 && __obf_fc2de974197327a1[1] < 32) || __obf_fc2de974197327a1[0] == 192 && __obf_fc2de974197327a1[1] == 168)
}

func __obf_148a164a645089a6(__obf_876f04b960c8f60c InterfaceAddrs) (uint16, error) {
	__obf_fc2de974197327a1, __obf_57cef2f010b4ce9a := __obf_dc9f389f94af361f(__obf_876f04b960c8f60c)
	if __obf_57cef2f010b4ce9a != nil {
		return 0, __obf_57cef2f010b4ce9a
	}

	return uint16(__obf_fc2de974197327a1[2])<<8 + uint16(__obf_fc2de974197327a1[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_413d3a76ff8ab79e uint64) time.Duration {
	return time.Duration(__obf_a2778e244d60c065(__obf_413d3a76ff8ab79e) * __obf_2f8ad3861a12d75e)
}

func __obf_a2778e244d60c065(__obf_413d3a76ff8ab79e uint64) uint64 {
	return __obf_413d3a76ff8ab79e >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_413d3a76ff8ab79e uint64) uint64 {
	const __obf_39a40eb8f9f68696 = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_413d3a76ff8ab79e & __obf_39a40eb8f9f68696 >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_413d3a76ff8ab79e uint64) uint64 {
	const __obf_d24beb0725029562 = uint64(1<<BitLenMachineID - 1)
	return __obf_413d3a76ff8ab79e & __obf_d24beb0725029562
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_413d3a76ff8ab79e uint64) map[string]uint64 {
	__obf_0d51d746f64bb81b := __obf_413d3a76ff8ab79e >> 63
	time := __obf_a2778e244d60c065(__obf_413d3a76ff8ab79e)
	__obf_32d285c02de19ff8 := SequenceNumber(__obf_413d3a76ff8ab79e)
	__obf_825d516c59f4195a := MachineID(__obf_413d3a76ff8ab79e)
	return map[string]uint64{
		"id":         __obf_413d3a76ff8ab79e,
		"msb":        __obf_0d51d746f64bb81b,
		"time":       time,
		"sequence":   __obf_32d285c02de19ff8,
		"machine-id": __obf_825d516c59f4195a,
	}
}
