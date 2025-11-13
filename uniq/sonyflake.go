// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_417508f5ae215d0a

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
	__obf_19f13c576fc455e1 *sync.Mutex
	__obf_ae6a78dd199cf619 int64
	__obf_e07e1c4829f6ccfd int64
	__obf_cc3a3b06b6be52c3 uint16
	__obf_92bb3f86f6f3e834 uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_3979cf40c3d000a4 = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_a74305e868c9f8cd Settings) (*Sonyflake, error) {
	if __obf_a74305e868c9f8cd.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}

	__obf_1dc1d5851f78a480 := new(Sonyflake)
	__obf_1dc1d5851f78a480.__obf_19f13c576fc455e1 = new(sync.Mutex)
	__obf_1dc1d5851f78a480.__obf_cc3a3b06b6be52c3 = uint16(1<<BitLenSequence - 1)

	if __obf_a74305e868c9f8cd.StartTime.IsZero() {
		__obf_1dc1d5851f78a480.__obf_ae6a78dd199cf619 = __obf_c8012e152bf746c3(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_1dc1d5851f78a480.__obf_ae6a78dd199cf619 = __obf_c8012e152bf746c3(__obf_a74305e868c9f8cd.StartTime)
	}

	var __obf_b99f62eaceb52d40 error
	if __obf_a74305e868c9f8cd.MachineID == nil {
		__obf_1dc1d5851f78a480.__obf_92bb3f86f6f3e834, __obf_b99f62eaceb52d40 = __obf_2e0275912c1a5893(__obf_3979cf40c3d000a4)
	} else {
		__obf_1dc1d5851f78a480.__obf_92bb3f86f6f3e834, __obf_b99f62eaceb52d40 = __obf_a74305e868c9f8cd.MachineID()
	}
	if __obf_b99f62eaceb52d40 != nil {
		return nil, __obf_b99f62eaceb52d40
	}

	if __obf_a74305e868c9f8cd.CheckMachineID != nil && !__obf_a74305e868c9f8cd.CheckMachineID(__obf_1dc1d5851f78a480.__obf_92bb3f86f6f3e834) {
		return nil, ErrInvalidMachineID
	}

	return __obf_1dc1d5851f78a480, nil
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
func (__obf_1dc1d5851f78a480 *Sonyflake) NextID() (string, error) {
	const __obf_da1ed67f1939196d = uint16(1<<BitLenSequence - 1)

	__obf_1dc1d5851f78a480.__obf_19f13c576fc455e1.Lock()
	defer __obf_1dc1d5851f78a480.__obf_19f13c576fc455e1.Unlock()

	__obf_b37711c5035db5b3 := __obf_3c5db6f43a172b87(__obf_1dc1d5851f78a480.__obf_ae6a78dd199cf619)
	if __obf_1dc1d5851f78a480.__obf_e07e1c4829f6ccfd < __obf_b37711c5035db5b3 {
		__obf_1dc1d5851f78a480.__obf_e07e1c4829f6ccfd = __obf_b37711c5035db5b3
		__obf_1dc1d5851f78a480.__obf_cc3a3b06b6be52c3 = 0
	} else { // sf.elapsedTime >= current
		__obf_1dc1d5851f78a480.__obf_cc3a3b06b6be52c3 = (__obf_1dc1d5851f78a480.__obf_cc3a3b06b6be52c3 + 1) & __obf_da1ed67f1939196d
		if __obf_1dc1d5851f78a480.__obf_cc3a3b06b6be52c3 == 0 {
			__obf_1dc1d5851f78a480.__obf_e07e1c4829f6ccfd++
			__obf_72b14c8a86c8b3a1 := __obf_1dc1d5851f78a480.__obf_e07e1c4829f6ccfd - __obf_b37711c5035db5b3
			time.Sleep(__obf_9b058eda855353e0((__obf_72b14c8a86c8b3a1)))
		}
	}

	return __obf_1dc1d5851f78a480.__obf_11ff538be92b70bf()
}

const __obf_d2b2a7ef999145d4 = 1e7 // nsec, i.e. 10 msec

func __obf_c8012e152bf746c3(__obf_59b4996dba32a5cd time.Time) int64 {
	return __obf_59b4996dba32a5cd.UTC().UnixNano() / __obf_d2b2a7ef999145d4
}

func __obf_3c5db6f43a172b87(__obf_ae6a78dd199cf619 int64) int64 {
	return __obf_c8012e152bf746c3(time.Now()) - __obf_ae6a78dd199cf619
}

func __obf_9b058eda855353e0(__obf_72b14c8a86c8b3a1 int64) time.Duration {
	return time.Duration(__obf_72b14c8a86c8b3a1*__obf_d2b2a7ef999145d4) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_d2b2a7ef999145d4)
}

func (__obf_1dc1d5851f78a480 *Sonyflake) __obf_11ff538be92b70bf() (string, error) {
	if __obf_1dc1d5851f78a480.__obf_e07e1c4829f6ccfd >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_1dc1d5851f78a480.__obf_e07e1c4829f6ccfd)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_1dc1d5851f78a480.__obf_cc3a3b06b6be52c3)<<BitLenMachineID|
		uint64(__obf_1dc1d5851f78a480.__obf_92bb3f86f6f3e834)), nil
}

func __obf_1106b5cac7dae757(__obf_dc8eeab52f029c53 InterfaceAddrs) (net.IP, error) {
	__obf_12901e53ebdc1216, __obf_b99f62eaceb52d40 := __obf_dc8eeab52f029c53()
	if __obf_b99f62eaceb52d40 != nil {
		return nil, __obf_b99f62eaceb52d40
	}

	for _, __obf_7b6aa415a6efafa3 := range __obf_12901e53ebdc1216 {
		__obf_9e00f7369e3a101b, __obf_0bd9d9196038949b := __obf_7b6aa415a6efafa3.(*net.IPNet)
		if !__obf_0bd9d9196038949b || __obf_9e00f7369e3a101b.IP.IsLoopback() {
			continue
		}

		__obf_9f01940a1b996679 := __obf_9e00f7369e3a101b.IP.To4()
		if __obf_2959cdd997af6f83(__obf_9f01940a1b996679) {
			return __obf_9f01940a1b996679, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_2959cdd997af6f83(__obf_9f01940a1b996679 net.IP) bool {
	return __obf_9f01940a1b996679 != nil &&
		(__obf_9f01940a1b996679[0] == 10 || __obf_9f01940a1b996679[0] == 172 && (__obf_9f01940a1b996679[1] >= 16 && __obf_9f01940a1b996679[1] < 32) || __obf_9f01940a1b996679[0] == 192 && __obf_9f01940a1b996679[1] == 168)
}

func __obf_2e0275912c1a5893(__obf_dc8eeab52f029c53 InterfaceAddrs) (uint16, error) {
	__obf_9f01940a1b996679, __obf_b99f62eaceb52d40 := __obf_1106b5cac7dae757(__obf_dc8eeab52f029c53)
	if __obf_b99f62eaceb52d40 != nil {
		return 0, __obf_b99f62eaceb52d40
	}

	return uint16(__obf_9f01940a1b996679[2])<<8 + uint16(__obf_9f01940a1b996679[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_48218c2a1a645d91 uint64) time.Duration {
	return time.Duration(__obf_e07e1c4829f6ccfd(__obf_48218c2a1a645d91) * __obf_d2b2a7ef999145d4)
}

func __obf_e07e1c4829f6ccfd(__obf_48218c2a1a645d91 uint64) uint64 {
	return __obf_48218c2a1a645d91 >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_48218c2a1a645d91 uint64) uint64 {
	const __obf_da1ed67f1939196d = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_48218c2a1a645d91 & __obf_da1ed67f1939196d >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_48218c2a1a645d91 uint64) uint64 {
	const __obf_c72f3c905cbb26f6 = uint64(1<<BitLenMachineID - 1)
	return __obf_48218c2a1a645d91 & __obf_c72f3c905cbb26f6
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_48218c2a1a645d91 uint64) map[string]uint64 {
	__obf_be064e94779ee2c2 := __obf_48218c2a1a645d91 >> 63
	time := __obf_e07e1c4829f6ccfd(__obf_48218c2a1a645d91)
	__obf_cc3a3b06b6be52c3 := SequenceNumber(__obf_48218c2a1a645d91)
	__obf_92bb3f86f6f3e834 := MachineID(__obf_48218c2a1a645d91)
	return map[string]uint64{
		"id":         __obf_48218c2a1a645d91,
		"msb":        __obf_be064e94779ee2c2,
		"time":       time,
		"sequence":   __obf_cc3a3b06b6be52c3,
		"machine-id": __obf_92bb3f86f6f3e834,
	}
}
