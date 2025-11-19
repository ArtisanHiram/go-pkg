// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_7913809aab6c8423

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
	__obf_27be151129f99025 *sync.Mutex
	__obf_fc0ec6f645c1b8bd int64
	__obf_07ab9ade19346688 int64
	__obf_007ec6572a283fb6 uint16
	__obf_a1d6b2555ea864dd uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_14e45a783661b383 = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_9aa5ddd833df9a52 Settings) (*Sonyflake, error) {
	if __obf_9aa5ddd833df9a52.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}

	__obf_609e1f690a9bdb71 := new(Sonyflake)
	__obf_609e1f690a9bdb71.__obf_27be151129f99025 = new(sync.Mutex)
	__obf_609e1f690a9bdb71.__obf_007ec6572a283fb6 = uint16(1<<BitLenSequence - 1)

	if __obf_9aa5ddd833df9a52.StartTime.IsZero() {
		__obf_609e1f690a9bdb71.__obf_fc0ec6f645c1b8bd = __obf_cc2f44abacb0408b(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_609e1f690a9bdb71.__obf_fc0ec6f645c1b8bd = __obf_cc2f44abacb0408b(__obf_9aa5ddd833df9a52.StartTime)
	}

	var __obf_e62cd9e417a87ee7 error
	if __obf_9aa5ddd833df9a52.MachineID == nil {
		__obf_609e1f690a9bdb71.__obf_a1d6b2555ea864dd, __obf_e62cd9e417a87ee7 = __obf_d9ec635ee8c6d754(__obf_14e45a783661b383)
	} else {
		__obf_609e1f690a9bdb71.__obf_a1d6b2555ea864dd, __obf_e62cd9e417a87ee7 = __obf_9aa5ddd833df9a52.MachineID()
	}
	if __obf_e62cd9e417a87ee7 != nil {
		return nil, __obf_e62cd9e417a87ee7
	}

	if __obf_9aa5ddd833df9a52.CheckMachineID != nil && !__obf_9aa5ddd833df9a52.CheckMachineID(__obf_609e1f690a9bdb71.__obf_a1d6b2555ea864dd) {
		return nil, ErrInvalidMachineID
	}

	return __obf_609e1f690a9bdb71, nil
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
func (__obf_609e1f690a9bdb71 *Sonyflake) NextID() (string, error) {
	const __obf_ab062e813ac77875 = uint16(1<<BitLenSequence - 1)

	__obf_609e1f690a9bdb71.__obf_27be151129f99025.Lock()
	defer __obf_609e1f690a9bdb71.__obf_27be151129f99025.Unlock()

	__obf_c5695d044cde6996 := __obf_c01ebe7dd2704a3d(__obf_609e1f690a9bdb71.__obf_fc0ec6f645c1b8bd)
	if __obf_609e1f690a9bdb71.__obf_07ab9ade19346688 < __obf_c5695d044cde6996 {
		__obf_609e1f690a9bdb71.__obf_07ab9ade19346688 = __obf_c5695d044cde6996
		__obf_609e1f690a9bdb71.__obf_007ec6572a283fb6 = 0
	} else { // sf.elapsedTime >= current
		__obf_609e1f690a9bdb71.__obf_007ec6572a283fb6 = (__obf_609e1f690a9bdb71.__obf_007ec6572a283fb6 + 1) & __obf_ab062e813ac77875
		if __obf_609e1f690a9bdb71.__obf_007ec6572a283fb6 == 0 {
			__obf_609e1f690a9bdb71.__obf_07ab9ade19346688++
			__obf_118efe86013deedf := __obf_609e1f690a9bdb71.__obf_07ab9ade19346688 - __obf_c5695d044cde6996
			time.Sleep(__obf_2b50c20e0ad2a012((__obf_118efe86013deedf)))
		}
	}

	return __obf_609e1f690a9bdb71.__obf_65554ed04c6fa6bc()
}

const __obf_13dc99f55cf6618f = 1e7 // nsec, i.e. 10 msec

func __obf_cc2f44abacb0408b(__obf_f154b37f4911a720 time.Time) int64 {
	return __obf_f154b37f4911a720.UTC().UnixNano() / __obf_13dc99f55cf6618f
}

func __obf_c01ebe7dd2704a3d(__obf_fc0ec6f645c1b8bd int64) int64 {
	return __obf_cc2f44abacb0408b(time.Now()) - __obf_fc0ec6f645c1b8bd
}

func __obf_2b50c20e0ad2a012(__obf_118efe86013deedf int64) time.Duration {
	return time.Duration(__obf_118efe86013deedf*__obf_13dc99f55cf6618f) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_13dc99f55cf6618f)
}

func (__obf_609e1f690a9bdb71 *Sonyflake) __obf_65554ed04c6fa6bc() (string, error) {
	if __obf_609e1f690a9bdb71.__obf_07ab9ade19346688 >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_609e1f690a9bdb71.__obf_07ab9ade19346688)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_609e1f690a9bdb71.__obf_007ec6572a283fb6)<<BitLenMachineID|
		uint64(__obf_609e1f690a9bdb71.__obf_a1d6b2555ea864dd)), nil
}

func __obf_804575fb9d8d64b6(__obf_c6787a8a9759149d InterfaceAddrs) (net.IP, error) {
	__obf_663788bf5979f5d2, __obf_e62cd9e417a87ee7 := __obf_c6787a8a9759149d()
	if __obf_e62cd9e417a87ee7 != nil {
		return nil, __obf_e62cd9e417a87ee7
	}

	for _, __obf_3da24acaf7d1fffd := range __obf_663788bf5979f5d2 {
		__obf_59e528b477c09677, __obf_c8f04275fcb5c686 := __obf_3da24acaf7d1fffd.(*net.IPNet)
		if !__obf_c8f04275fcb5c686 || __obf_59e528b477c09677.IP.IsLoopback() {
			continue
		}

		__obf_63e8f9372174d7f0 := __obf_59e528b477c09677.IP.To4()
		if __obf_43c4258a47d63b90(__obf_63e8f9372174d7f0) {
			return __obf_63e8f9372174d7f0, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_43c4258a47d63b90(__obf_63e8f9372174d7f0 net.IP) bool {
	return __obf_63e8f9372174d7f0 != nil &&
		(__obf_63e8f9372174d7f0[0] == 10 || __obf_63e8f9372174d7f0[0] == 172 && (__obf_63e8f9372174d7f0[1] >= 16 && __obf_63e8f9372174d7f0[1] < 32) || __obf_63e8f9372174d7f0[0] == 192 && __obf_63e8f9372174d7f0[1] == 168)
}

func __obf_d9ec635ee8c6d754(__obf_c6787a8a9759149d InterfaceAddrs) (uint16, error) {
	__obf_63e8f9372174d7f0, __obf_e62cd9e417a87ee7 := __obf_804575fb9d8d64b6(__obf_c6787a8a9759149d)
	if __obf_e62cd9e417a87ee7 != nil {
		return 0, __obf_e62cd9e417a87ee7
	}

	return uint16(__obf_63e8f9372174d7f0[2])<<8 + uint16(__obf_63e8f9372174d7f0[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_1dfe314328e20778 uint64) time.Duration {
	return time.Duration(__obf_07ab9ade19346688(__obf_1dfe314328e20778) * __obf_13dc99f55cf6618f)
}

func __obf_07ab9ade19346688(__obf_1dfe314328e20778 uint64) uint64 {
	return __obf_1dfe314328e20778 >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_1dfe314328e20778 uint64) uint64 {
	const __obf_ab062e813ac77875 = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_1dfe314328e20778 & __obf_ab062e813ac77875 >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_1dfe314328e20778 uint64) uint64 {
	const __obf_d65150b72dcc040d = uint64(1<<BitLenMachineID - 1)
	return __obf_1dfe314328e20778 & __obf_d65150b72dcc040d
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_1dfe314328e20778 uint64) map[string]uint64 {
	__obf_9ec5dd2e93ae8261 := __obf_1dfe314328e20778 >> 63
	time := __obf_07ab9ade19346688(__obf_1dfe314328e20778)
	__obf_007ec6572a283fb6 := SequenceNumber(__obf_1dfe314328e20778)
	__obf_a1d6b2555ea864dd := MachineID(__obf_1dfe314328e20778)
	return map[string]uint64{
		"id":         __obf_1dfe314328e20778,
		"msb":        __obf_9ec5dd2e93ae8261,
		"time":       time,
		"sequence":   __obf_007ec6572a283fb6,
		"machine-id": __obf_a1d6b2555ea864dd,
	}
}
