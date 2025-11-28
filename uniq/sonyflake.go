// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_e2239f4853c61591

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
	__obf_14955da4c545dc5e *sync.Mutex
	__obf_f2964e8024ee677e int64
	__obf_abb9417b21679785 int64
	__obf_4a953d8c26aa0257 uint16
	__obf_7a7df53f3536f4b5 uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_f9d2048a569a8081 = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_626eedced3bcc321 Settings) (*Sonyflake, error) {
	if __obf_626eedced3bcc321.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}

	__obf_0cc8b14a978ef1b7 := new(Sonyflake)
	__obf_0cc8b14a978ef1b7.__obf_14955da4c545dc5e = new(sync.Mutex)
	__obf_0cc8b14a978ef1b7.__obf_4a953d8c26aa0257 = uint16(1<<BitLenSequence - 1)

	if __obf_626eedced3bcc321.StartTime.IsZero() {
		__obf_0cc8b14a978ef1b7.__obf_f2964e8024ee677e = __obf_3083090acf13e42e(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_0cc8b14a978ef1b7.__obf_f2964e8024ee677e = __obf_3083090acf13e42e(__obf_626eedced3bcc321.StartTime)
	}

	var __obf_bb9dc4c4f445b22f error
	if __obf_626eedced3bcc321.MachineID == nil {
		__obf_0cc8b14a978ef1b7.__obf_7a7df53f3536f4b5, __obf_bb9dc4c4f445b22f = __obf_78f4aaaed7ade9d6(__obf_f9d2048a569a8081)
	} else {
		__obf_0cc8b14a978ef1b7.__obf_7a7df53f3536f4b5, __obf_bb9dc4c4f445b22f = __obf_626eedced3bcc321.MachineID()
	}
	if __obf_bb9dc4c4f445b22f != nil {
		return nil, __obf_bb9dc4c4f445b22f
	}

	if __obf_626eedced3bcc321.CheckMachineID != nil && !__obf_626eedced3bcc321.CheckMachineID(__obf_0cc8b14a978ef1b7.__obf_7a7df53f3536f4b5) {
		return nil, ErrInvalidMachineID
	}

	return __obf_0cc8b14a978ef1b7, nil
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
func (__obf_0cc8b14a978ef1b7 *Sonyflake) NextID() (string, error) {
	const __obf_1c8fe20b293bca9b = uint16(1<<BitLenSequence - 1)

	__obf_0cc8b14a978ef1b7.__obf_14955da4c545dc5e.Lock()
	defer __obf_0cc8b14a978ef1b7.__obf_14955da4c545dc5e.Unlock()

	__obf_d086261bae877500 := __obf_8d6d6f5355a2662d(__obf_0cc8b14a978ef1b7.__obf_f2964e8024ee677e)
	if __obf_0cc8b14a978ef1b7.__obf_abb9417b21679785 < __obf_d086261bae877500 {
		__obf_0cc8b14a978ef1b7.__obf_abb9417b21679785 = __obf_d086261bae877500
		__obf_0cc8b14a978ef1b7.__obf_4a953d8c26aa0257 = 0
	} else { // sf.elapsedTime >= current
		__obf_0cc8b14a978ef1b7.__obf_4a953d8c26aa0257 = (__obf_0cc8b14a978ef1b7.__obf_4a953d8c26aa0257 + 1) & __obf_1c8fe20b293bca9b
		if __obf_0cc8b14a978ef1b7.__obf_4a953d8c26aa0257 == 0 {
			__obf_0cc8b14a978ef1b7.__obf_abb9417b21679785++
			__obf_5726766afca9a7f3 := __obf_0cc8b14a978ef1b7.__obf_abb9417b21679785 - __obf_d086261bae877500
			time.Sleep(__obf_e4e06342d78811a5((__obf_5726766afca9a7f3)))
		}
	}

	return __obf_0cc8b14a978ef1b7.__obf_8e5e99e3a4354ce9()
}

const __obf_d0218570b2c0bc43 = 1e7 // nsec, i.e. 10 msec

func __obf_3083090acf13e42e(__obf_de192fc01291fe85 time.Time) int64 {
	return __obf_de192fc01291fe85.UTC().UnixNano() / __obf_d0218570b2c0bc43
}

func __obf_8d6d6f5355a2662d(__obf_f2964e8024ee677e int64) int64 {
	return __obf_3083090acf13e42e(time.Now()) - __obf_f2964e8024ee677e
}

func __obf_e4e06342d78811a5(__obf_5726766afca9a7f3 int64) time.Duration {
	return time.Duration(__obf_5726766afca9a7f3*__obf_d0218570b2c0bc43) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_d0218570b2c0bc43)
}

func (__obf_0cc8b14a978ef1b7 *Sonyflake) __obf_8e5e99e3a4354ce9() (string, error) {
	if __obf_0cc8b14a978ef1b7.__obf_abb9417b21679785 >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_0cc8b14a978ef1b7.__obf_abb9417b21679785)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_0cc8b14a978ef1b7.__obf_4a953d8c26aa0257)<<BitLenMachineID|
		uint64(__obf_0cc8b14a978ef1b7.__obf_7a7df53f3536f4b5)), nil
}

func __obf_f0d2664204670974(__obf_c86c2ca6156ba7aa InterfaceAddrs) (net.IP, error) {
	__obf_efb146f587effe0d, __obf_bb9dc4c4f445b22f := __obf_c86c2ca6156ba7aa()
	if __obf_bb9dc4c4f445b22f != nil {
		return nil, __obf_bb9dc4c4f445b22f
	}

	for _, __obf_f570211fa82d5006 := range __obf_efb146f587effe0d {
		__obf_4c9a03fdc80f97ba, __obf_c8f6ef41918857c4 := __obf_f570211fa82d5006.(*net.IPNet)
		if !__obf_c8f6ef41918857c4 || __obf_4c9a03fdc80f97ba.IP.IsLoopback() {
			continue
		}

		__obf_2dd9106aa4c0bbc9 := __obf_4c9a03fdc80f97ba.IP.To4()
		if __obf_baf78644606f4204(__obf_2dd9106aa4c0bbc9) {
			return __obf_2dd9106aa4c0bbc9, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_baf78644606f4204(__obf_2dd9106aa4c0bbc9 net.IP) bool {
	return __obf_2dd9106aa4c0bbc9 != nil &&
		(__obf_2dd9106aa4c0bbc9[0] == 10 || __obf_2dd9106aa4c0bbc9[0] == 172 && (__obf_2dd9106aa4c0bbc9[1] >= 16 && __obf_2dd9106aa4c0bbc9[1] < 32) || __obf_2dd9106aa4c0bbc9[0] == 192 && __obf_2dd9106aa4c0bbc9[1] == 168)
}

func __obf_78f4aaaed7ade9d6(__obf_c86c2ca6156ba7aa InterfaceAddrs) (uint16, error) {
	__obf_2dd9106aa4c0bbc9, __obf_bb9dc4c4f445b22f := __obf_f0d2664204670974(__obf_c86c2ca6156ba7aa)
	if __obf_bb9dc4c4f445b22f != nil {
		return 0, __obf_bb9dc4c4f445b22f
	}

	return uint16(__obf_2dd9106aa4c0bbc9[2])<<8 + uint16(__obf_2dd9106aa4c0bbc9[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_cee82df7c007227d uint64) time.Duration {
	return time.Duration(__obf_abb9417b21679785(__obf_cee82df7c007227d) * __obf_d0218570b2c0bc43)
}

func __obf_abb9417b21679785(__obf_cee82df7c007227d uint64) uint64 {
	return __obf_cee82df7c007227d >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_cee82df7c007227d uint64) uint64 {
	const __obf_1c8fe20b293bca9b = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_cee82df7c007227d & __obf_1c8fe20b293bca9b >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_cee82df7c007227d uint64) uint64 {
	const __obf_64363f64a5dfa787 = uint64(1<<BitLenMachineID - 1)
	return __obf_cee82df7c007227d & __obf_64363f64a5dfa787
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_cee82df7c007227d uint64) map[string]uint64 {
	__obf_05fc973619505216 := __obf_cee82df7c007227d >> 63
	time := __obf_abb9417b21679785(__obf_cee82df7c007227d)
	__obf_4a953d8c26aa0257 := SequenceNumber(__obf_cee82df7c007227d)
	__obf_7a7df53f3536f4b5 := MachineID(__obf_cee82df7c007227d)
	return map[string]uint64{
		"id":         __obf_cee82df7c007227d,
		"msb":        __obf_05fc973619505216,
		"time":       time,
		"sequence":   __obf_4a953d8c26aa0257,
		"machine-id": __obf_7a7df53f3536f4b5,
	}
}
