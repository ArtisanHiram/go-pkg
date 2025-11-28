// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_0f0e0d1ff72f3ff0

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
	__obf_63f6c289520bd7e2 *sync.Mutex
	__obf_bc63f3e0335cd8ba int64
	__obf_c648d99dd6a74354 int64
	__obf_60bc88779bdc777d uint16
	__obf_91ff39c23c7fcbf1 uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_160beb80d20992ed = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_10f225029ba7d295 Settings) (*Sonyflake, error) {
	if __obf_10f225029ba7d295.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}

	__obf_93c464fcd213d377 := new(Sonyflake)
	__obf_93c464fcd213d377.__obf_63f6c289520bd7e2 = new(sync.Mutex)
	__obf_93c464fcd213d377.__obf_60bc88779bdc777d = uint16(1<<BitLenSequence - 1)

	if __obf_10f225029ba7d295.StartTime.IsZero() {
		__obf_93c464fcd213d377.__obf_bc63f3e0335cd8ba = __obf_7e4c1daef2ced091(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_93c464fcd213d377.__obf_bc63f3e0335cd8ba = __obf_7e4c1daef2ced091(__obf_10f225029ba7d295.StartTime)
	}

	var __obf_f24ec7ae16791add error
	if __obf_10f225029ba7d295.MachineID == nil {
		__obf_93c464fcd213d377.__obf_91ff39c23c7fcbf1, __obf_f24ec7ae16791add = __obf_bbf4e99a718bfaf5(__obf_160beb80d20992ed)
	} else {
		__obf_93c464fcd213d377.__obf_91ff39c23c7fcbf1, __obf_f24ec7ae16791add = __obf_10f225029ba7d295.MachineID()
	}
	if __obf_f24ec7ae16791add != nil {
		return nil, __obf_f24ec7ae16791add
	}

	if __obf_10f225029ba7d295.CheckMachineID != nil && !__obf_10f225029ba7d295.CheckMachineID(__obf_93c464fcd213d377.__obf_91ff39c23c7fcbf1) {
		return nil, ErrInvalidMachineID
	}

	return __obf_93c464fcd213d377, nil
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
func (__obf_93c464fcd213d377 *Sonyflake) NextID() (string, error) {
	const __obf_52ef6e2da5c40865 = uint16(1<<BitLenSequence - 1)

	__obf_93c464fcd213d377.__obf_63f6c289520bd7e2.Lock()
	defer __obf_93c464fcd213d377.__obf_63f6c289520bd7e2.Unlock()

	__obf_5d6a616846e72812 := __obf_92d75f1dc8129463(__obf_93c464fcd213d377.__obf_bc63f3e0335cd8ba)
	if __obf_93c464fcd213d377.__obf_c648d99dd6a74354 < __obf_5d6a616846e72812 {
		__obf_93c464fcd213d377.__obf_c648d99dd6a74354 = __obf_5d6a616846e72812
		__obf_93c464fcd213d377.__obf_60bc88779bdc777d = 0
	} else { // sf.elapsedTime >= current
		__obf_93c464fcd213d377.__obf_60bc88779bdc777d = (__obf_93c464fcd213d377.__obf_60bc88779bdc777d + 1) & __obf_52ef6e2da5c40865
		if __obf_93c464fcd213d377.__obf_60bc88779bdc777d == 0 {
			__obf_93c464fcd213d377.__obf_c648d99dd6a74354++
			__obf_11bf7d3195bd594b := __obf_93c464fcd213d377.__obf_c648d99dd6a74354 - __obf_5d6a616846e72812
			time.Sleep(__obf_57c82dd4f7ef85e4((__obf_11bf7d3195bd594b)))
		}
	}

	return __obf_93c464fcd213d377.__obf_de9d3d4479b0867d()
}

const __obf_a6e3ddfca53f1c41 = 1e7 // nsec, i.e. 10 msec

func __obf_7e4c1daef2ced091(__obf_3c7f4fcc21cb9f0d time.Time) int64 {
	return __obf_3c7f4fcc21cb9f0d.UTC().UnixNano() / __obf_a6e3ddfca53f1c41
}

func __obf_92d75f1dc8129463(__obf_bc63f3e0335cd8ba int64) int64 {
	return __obf_7e4c1daef2ced091(time.Now()) - __obf_bc63f3e0335cd8ba
}

func __obf_57c82dd4f7ef85e4(__obf_11bf7d3195bd594b int64) time.Duration {
	return time.Duration(__obf_11bf7d3195bd594b*__obf_a6e3ddfca53f1c41) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_a6e3ddfca53f1c41)
}

func (__obf_93c464fcd213d377 *Sonyflake) __obf_de9d3d4479b0867d() (string, error) {
	if __obf_93c464fcd213d377.__obf_c648d99dd6a74354 >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_93c464fcd213d377.__obf_c648d99dd6a74354)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_93c464fcd213d377.__obf_60bc88779bdc777d)<<BitLenMachineID|
		uint64(__obf_93c464fcd213d377.__obf_91ff39c23c7fcbf1)), nil
}

func __obf_cd1a65bbacac6a62(__obf_99d7ff6bde6d7540 InterfaceAddrs) (net.IP, error) {
	__obf_243666a925f1dff1, __obf_f24ec7ae16791add := __obf_99d7ff6bde6d7540()
	if __obf_f24ec7ae16791add != nil {
		return nil, __obf_f24ec7ae16791add
	}

	for _, __obf_92d04951402c828c := range __obf_243666a925f1dff1 {
		__obf_ddb9513e7b6bc4b7, __obf_b08eebaa4bd2dee3 := __obf_92d04951402c828c.(*net.IPNet)
		if !__obf_b08eebaa4bd2dee3 || __obf_ddb9513e7b6bc4b7.IP.IsLoopback() {
			continue
		}

		__obf_61706468d58f484e := __obf_ddb9513e7b6bc4b7.IP.To4()
		if __obf_aabc836bce84c46e(__obf_61706468d58f484e) {
			return __obf_61706468d58f484e, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_aabc836bce84c46e(__obf_61706468d58f484e net.IP) bool {
	return __obf_61706468d58f484e != nil &&
		(__obf_61706468d58f484e[0] == 10 || __obf_61706468d58f484e[0] == 172 && (__obf_61706468d58f484e[1] >= 16 && __obf_61706468d58f484e[1] < 32) || __obf_61706468d58f484e[0] == 192 && __obf_61706468d58f484e[1] == 168)
}

func __obf_bbf4e99a718bfaf5(__obf_99d7ff6bde6d7540 InterfaceAddrs) (uint16, error) {
	__obf_61706468d58f484e, __obf_f24ec7ae16791add := __obf_cd1a65bbacac6a62(__obf_99d7ff6bde6d7540)
	if __obf_f24ec7ae16791add != nil {
		return 0, __obf_f24ec7ae16791add
	}

	return uint16(__obf_61706468d58f484e[2])<<8 + uint16(__obf_61706468d58f484e[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_a7593a31c69c5d40 uint64) time.Duration {
	return time.Duration(__obf_c648d99dd6a74354(__obf_a7593a31c69c5d40) * __obf_a6e3ddfca53f1c41)
}

func __obf_c648d99dd6a74354(__obf_a7593a31c69c5d40 uint64) uint64 {
	return __obf_a7593a31c69c5d40 >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_a7593a31c69c5d40 uint64) uint64 {
	const __obf_52ef6e2da5c40865 = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_a7593a31c69c5d40 & __obf_52ef6e2da5c40865 >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_a7593a31c69c5d40 uint64) uint64 {
	const __obf_d4c0b0343970b293 = uint64(1<<BitLenMachineID - 1)
	return __obf_a7593a31c69c5d40 & __obf_d4c0b0343970b293
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_a7593a31c69c5d40 uint64) map[string]uint64 {
	__obf_32edf9de6f654b3b := __obf_a7593a31c69c5d40 >> 63
	time := __obf_c648d99dd6a74354(__obf_a7593a31c69c5d40)
	__obf_60bc88779bdc777d := SequenceNumber(__obf_a7593a31c69c5d40)
	__obf_91ff39c23c7fcbf1 := MachineID(__obf_a7593a31c69c5d40)
	return map[string]uint64{
		"id":         __obf_a7593a31c69c5d40,
		"msb":        __obf_32edf9de6f654b3b,
		"time":       time,
		"sequence":   __obf_60bc88779bdc777d,
		"machine-id": __obf_91ff39c23c7fcbf1,
	}
}
