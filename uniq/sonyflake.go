// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_4d7a51f8e2f0494a

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
	__obf_b1d5db65fb023892 *sync.Mutex
	__obf_1139230f93bac2b4 int64
	__obf_86f0900e85ed242a int64
	__obf_de8c2d4f1a1d02ed uint16
	__obf_b4613fc9fdd7e73f uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_266445fc230d550f = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_610f1e648b94aa5b Settings) (*Sonyflake, error) {
	if __obf_610f1e648b94aa5b.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}
	__obf_8687e17f8ca5af91 := new(Sonyflake)
	__obf_8687e17f8ca5af91.__obf_b1d5db65fb023892 = new(sync.Mutex)
	__obf_8687e17f8ca5af91.__obf_de8c2d4f1a1d02ed = uint16(1<<BitLenSequence - 1)

	if __obf_610f1e648b94aa5b.StartTime.IsZero() {
		__obf_8687e17f8ca5af91.__obf_1139230f93bac2b4 = __obf_199363dc396f25d0(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_8687e17f8ca5af91.__obf_1139230f93bac2b4 = __obf_199363dc396f25d0(__obf_610f1e648b94aa5b.StartTime)
	}

	var __obf_38c5d446257f8d9a error
	if __obf_610f1e648b94aa5b.MachineID == nil {
		__obf_8687e17f8ca5af91.__obf_b4613fc9fdd7e73f, __obf_38c5d446257f8d9a = __obf_af327f01cb1e0b2f(__obf_266445fc230d550f)
	} else {
		__obf_8687e17f8ca5af91.__obf_b4613fc9fdd7e73f, __obf_38c5d446257f8d9a = __obf_610f1e648b94aa5b.MachineID()
	}
	if __obf_38c5d446257f8d9a != nil {
		return nil, __obf_38c5d446257f8d9a
	}

	if __obf_610f1e648b94aa5b.CheckMachineID != nil && !__obf_610f1e648b94aa5b.CheckMachineID(__obf_8687e17f8ca5af91.__obf_b4613fc9fdd7e73f) {
		return nil, ErrInvalidMachineID
	}

	return __obf_8687e17f8ca5af91, nil
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
func (__obf_8687e17f8ca5af91 *Sonyflake) NextID() (string, error) {
	const __obf_cf91983a8cf9928e = uint16(1<<BitLenSequence - 1)
	__obf_8687e17f8ca5af91.__obf_b1d5db65fb023892.
		Lock()
	defer __obf_8687e17f8ca5af91.__obf_b1d5db65fb023892.Unlock()
	__obf_f467f02c3d108ad1 := __obf_255a1707821bb9ca(__obf_8687e17f8ca5af91.__obf_1139230f93bac2b4)
	if __obf_8687e17f8ca5af91.__obf_86f0900e85ed242a < __obf_f467f02c3d108ad1 {
		__obf_8687e17f8ca5af91.__obf_86f0900e85ed242a = __obf_f467f02c3d108ad1
		__obf_8687e17f8ca5af91.

			// sf.elapsedTime >= current
			__obf_de8c2d4f1a1d02ed = 0
	} else {
		__obf_8687e17f8ca5af91.__obf_de8c2d4f1a1d02ed = (__obf_8687e17f8ca5af91.__obf_de8c2d4f1a1d02ed + 1) & __obf_cf91983a8cf9928e
		if __obf_8687e17f8ca5af91.__obf_de8c2d4f1a1d02ed == 0 {
			__obf_8687e17f8ca5af91.__obf_86f0900e85ed242a++
			__obf_2e00240033b15699 := __obf_8687e17f8ca5af91.__obf_86f0900e85ed242a - __obf_f467f02c3d108ad1
			time.Sleep(__obf_c9103bf6571831c5((__obf_2e00240033b15699)))
		}
	}

	return __obf_8687e17f8ca5af91.__obf_4e0039d408814931()
}

const __obf_76c124ec3d092534 = 1e7 // nsec, i.e. 10 msec

func __obf_199363dc396f25d0(__obf_d5aaddb5c5a3e313 time.Time) int64 {
	return __obf_d5aaddb5c5a3e313.UTC().UnixNano() / __obf_76c124ec3d092534
}

func __obf_255a1707821bb9ca(__obf_1139230f93bac2b4 int64) int64 {
	return __obf_199363dc396f25d0(time.Now()) - __obf_1139230f93bac2b4
}

func __obf_c9103bf6571831c5(__obf_2e00240033b15699 int64) time.Duration {
	return time.Duration(__obf_2e00240033b15699*__obf_76c124ec3d092534) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_76c124ec3d092534)
}

func (__obf_8687e17f8ca5af91 *Sonyflake) __obf_4e0039d408814931() (string, error) {
	if __obf_8687e17f8ca5af91.__obf_86f0900e85ed242a >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_8687e17f8ca5af91.__obf_86f0900e85ed242a)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_8687e17f8ca5af91.__obf_de8c2d4f1a1d02ed)<<BitLenMachineID|
		uint64(__obf_8687e17f8ca5af91.__obf_b4613fc9fdd7e73f)), nil
}

func __obf_4be930c1af0f2a1d(__obf_c2842fd4d079c9f1 InterfaceAddrs) (net.IP, error) {
	__obf_e80721a0c199c67a, __obf_38c5d446257f8d9a := __obf_c2842fd4d079c9f1()
	if __obf_38c5d446257f8d9a != nil {
		return nil, __obf_38c5d446257f8d9a
	}

	for _, __obf_c5d6005604272752 := range __obf_e80721a0c199c67a {
		__obf_954347b5d578aa3d, __obf_a596699de7db8dcb := __obf_c5d6005604272752.(*net.IPNet)
		if !__obf_a596699de7db8dcb || __obf_954347b5d578aa3d.IP.IsLoopback() {
			continue
		}
		__obf_8aa7d939e74a1ed8 := __obf_954347b5d578aa3d.IP.To4()
		if __obf_3df71e08fd6688a1(__obf_8aa7d939e74a1ed8) {
			return __obf_8aa7d939e74a1ed8, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_3df71e08fd6688a1(__obf_8aa7d939e74a1ed8 net.IP) bool {
	return __obf_8aa7d939e74a1ed8 != nil &&
		(__obf_8aa7d939e74a1ed8[0] == 10 || __obf_8aa7d939e74a1ed8[0] == 172 && (__obf_8aa7d939e74a1ed8[1] >= 16 && __obf_8aa7d939e74a1ed8[1] < 32) || __obf_8aa7d939e74a1ed8[0] == 192 && __obf_8aa7d939e74a1ed8[1] == 168)
}

func __obf_af327f01cb1e0b2f(__obf_c2842fd4d079c9f1 InterfaceAddrs) (uint16, error) {
	__obf_8aa7d939e74a1ed8, __obf_38c5d446257f8d9a := __obf_4be930c1af0f2a1d(__obf_c2842fd4d079c9f1)
	if __obf_38c5d446257f8d9a != nil {
		return 0, __obf_38c5d446257f8d9a
	}

	return uint16(__obf_8aa7d939e74a1ed8[2])<<8 + uint16(__obf_8aa7d939e74a1ed8[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_721283c42292359f uint64) time.Duration {
	return time.Duration(__obf_86f0900e85ed242a(__obf_721283c42292359f) * __obf_76c124ec3d092534)
}

func __obf_86f0900e85ed242a(__obf_721283c42292359f uint64) uint64 {
	return __obf_721283c42292359f >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_721283c42292359f uint64) uint64 {
	const __obf_cf91983a8cf9928e = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_721283c42292359f & __obf_cf91983a8cf9928e >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_721283c42292359f uint64) uint64 {
	const __obf_0991bb70b5b1fb67 = uint64(1<<BitLenMachineID - 1)
	return __obf_721283c42292359f & __obf_0991bb70b5b1fb67
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_721283c42292359f uint64) map[string]uint64 {
	__obf_75344f88b5e57c8c := __obf_721283c42292359f >> 63
	time := __obf_86f0900e85ed242a(__obf_721283c42292359f)
	__obf_de8c2d4f1a1d02ed := SequenceNumber(__obf_721283c42292359f)
	__obf_b4613fc9fdd7e73f := MachineID(__obf_721283c42292359f)
	return map[string]uint64{
		"id":         __obf_721283c42292359f,
		"msb":        __obf_75344f88b5e57c8c,
		"time":       time,
		"sequence":   __obf_de8c2d4f1a1d02ed,
		"machine-id": __obf_b4613fc9fdd7e73f,
	}
}
