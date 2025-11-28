// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_8fe28252c1b01dfe

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
	__obf_acf835f51a66de41 *sync.Mutex
	__obf_68dc3c14620a0639 int64
	__obf_04e1395853ab8744 int64
	__obf_1e50fedd92e91188 uint16
	__obf_6b5c579c35cb3c44 uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_1778f2ce04e37e2f = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_2b20260661c57163 Settings) (*Sonyflake, error) {
	if __obf_2b20260661c57163.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}

	__obf_a28805d2857e9e17 := new(Sonyflake)
	__obf_a28805d2857e9e17.__obf_acf835f51a66de41 = new(sync.Mutex)
	__obf_a28805d2857e9e17.__obf_1e50fedd92e91188 = uint16(1<<BitLenSequence - 1)

	if __obf_2b20260661c57163.StartTime.IsZero() {
		__obf_a28805d2857e9e17.__obf_68dc3c14620a0639 = __obf_f9a8784837e6ee54(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_a28805d2857e9e17.__obf_68dc3c14620a0639 = __obf_f9a8784837e6ee54(__obf_2b20260661c57163.StartTime)
	}

	var __obf_1082dd6e56192e3a error
	if __obf_2b20260661c57163.MachineID == nil {
		__obf_a28805d2857e9e17.__obf_6b5c579c35cb3c44, __obf_1082dd6e56192e3a = __obf_ca8637483b1868e5(__obf_1778f2ce04e37e2f)
	} else {
		__obf_a28805d2857e9e17.__obf_6b5c579c35cb3c44, __obf_1082dd6e56192e3a = __obf_2b20260661c57163.MachineID()
	}
	if __obf_1082dd6e56192e3a != nil {
		return nil, __obf_1082dd6e56192e3a
	}

	if __obf_2b20260661c57163.CheckMachineID != nil && !__obf_2b20260661c57163.CheckMachineID(__obf_a28805d2857e9e17.__obf_6b5c579c35cb3c44) {
		return nil, ErrInvalidMachineID
	}

	return __obf_a28805d2857e9e17, nil
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
func (__obf_a28805d2857e9e17 *Sonyflake) NextID() (string, error) {
	const __obf_2a3a242797dcb190 = uint16(1<<BitLenSequence - 1)

	__obf_a28805d2857e9e17.__obf_acf835f51a66de41.Lock()
	defer __obf_a28805d2857e9e17.__obf_acf835f51a66de41.Unlock()

	__obf_291b6d36f046db8e := __obf_c7f76b0e8e76d077(__obf_a28805d2857e9e17.__obf_68dc3c14620a0639)
	if __obf_a28805d2857e9e17.__obf_04e1395853ab8744 < __obf_291b6d36f046db8e {
		__obf_a28805d2857e9e17.__obf_04e1395853ab8744 = __obf_291b6d36f046db8e
		__obf_a28805d2857e9e17.__obf_1e50fedd92e91188 = 0
	} else { // sf.elapsedTime >= current
		__obf_a28805d2857e9e17.__obf_1e50fedd92e91188 = (__obf_a28805d2857e9e17.__obf_1e50fedd92e91188 + 1) & __obf_2a3a242797dcb190
		if __obf_a28805d2857e9e17.__obf_1e50fedd92e91188 == 0 {
			__obf_a28805d2857e9e17.__obf_04e1395853ab8744++
			__obf_dc2458d62d7a43d1 := __obf_a28805d2857e9e17.__obf_04e1395853ab8744 - __obf_291b6d36f046db8e
			time.Sleep(__obf_1f96f47640973191((__obf_dc2458d62d7a43d1)))
		}
	}

	return __obf_a28805d2857e9e17.__obf_ab13c9af6d6c7187()
}

const __obf_390a45458ea27b4f = 1e7 // nsec, i.e. 10 msec

func __obf_f9a8784837e6ee54(__obf_8bbd0f9c7d55c1cc time.Time) int64 {
	return __obf_8bbd0f9c7d55c1cc.UTC().UnixNano() / __obf_390a45458ea27b4f
}

func __obf_c7f76b0e8e76d077(__obf_68dc3c14620a0639 int64) int64 {
	return __obf_f9a8784837e6ee54(time.Now()) - __obf_68dc3c14620a0639
}

func __obf_1f96f47640973191(__obf_dc2458d62d7a43d1 int64) time.Duration {
	return time.Duration(__obf_dc2458d62d7a43d1*__obf_390a45458ea27b4f) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_390a45458ea27b4f)
}

func (__obf_a28805d2857e9e17 *Sonyflake) __obf_ab13c9af6d6c7187() (string, error) {
	if __obf_a28805d2857e9e17.__obf_04e1395853ab8744 >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_a28805d2857e9e17.__obf_04e1395853ab8744)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_a28805d2857e9e17.__obf_1e50fedd92e91188)<<BitLenMachineID|
		uint64(__obf_a28805d2857e9e17.__obf_6b5c579c35cb3c44)), nil
}

func __obf_8cb2ff83785d0596(__obf_40f7a78a142552eb InterfaceAddrs) (net.IP, error) {
	__obf_be47f6f5ceaac641, __obf_1082dd6e56192e3a := __obf_40f7a78a142552eb()
	if __obf_1082dd6e56192e3a != nil {
		return nil, __obf_1082dd6e56192e3a
	}

	for _, __obf_193a3f4eef5d295c := range __obf_be47f6f5ceaac641 {
		__obf_5e5b8d7cb393005c, __obf_a3279a70b6b813b3 := __obf_193a3f4eef5d295c.(*net.IPNet)
		if !__obf_a3279a70b6b813b3 || __obf_5e5b8d7cb393005c.IP.IsLoopback() {
			continue
		}

		__obf_41e183106fc42283 := __obf_5e5b8d7cb393005c.IP.To4()
		if __obf_fdaf152ac6432929(__obf_41e183106fc42283) {
			return __obf_41e183106fc42283, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_fdaf152ac6432929(__obf_41e183106fc42283 net.IP) bool {
	return __obf_41e183106fc42283 != nil &&
		(__obf_41e183106fc42283[0] == 10 || __obf_41e183106fc42283[0] == 172 && (__obf_41e183106fc42283[1] >= 16 && __obf_41e183106fc42283[1] < 32) || __obf_41e183106fc42283[0] == 192 && __obf_41e183106fc42283[1] == 168)
}

func __obf_ca8637483b1868e5(__obf_40f7a78a142552eb InterfaceAddrs) (uint16, error) {
	__obf_41e183106fc42283, __obf_1082dd6e56192e3a := __obf_8cb2ff83785d0596(__obf_40f7a78a142552eb)
	if __obf_1082dd6e56192e3a != nil {
		return 0, __obf_1082dd6e56192e3a
	}

	return uint16(__obf_41e183106fc42283[2])<<8 + uint16(__obf_41e183106fc42283[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_b6d2ddd191a19ef1 uint64) time.Duration {
	return time.Duration(__obf_04e1395853ab8744(__obf_b6d2ddd191a19ef1) * __obf_390a45458ea27b4f)
}

func __obf_04e1395853ab8744(__obf_b6d2ddd191a19ef1 uint64) uint64 {
	return __obf_b6d2ddd191a19ef1 >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_b6d2ddd191a19ef1 uint64) uint64 {
	const __obf_2a3a242797dcb190 = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_b6d2ddd191a19ef1 & __obf_2a3a242797dcb190 >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_b6d2ddd191a19ef1 uint64) uint64 {
	const __obf_2983b4de02f1846b = uint64(1<<BitLenMachineID - 1)
	return __obf_b6d2ddd191a19ef1 & __obf_2983b4de02f1846b
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_b6d2ddd191a19ef1 uint64) map[string]uint64 {
	__obf_f7b293a9f8ecf409 := __obf_b6d2ddd191a19ef1 >> 63
	time := __obf_04e1395853ab8744(__obf_b6d2ddd191a19ef1)
	__obf_1e50fedd92e91188 := SequenceNumber(__obf_b6d2ddd191a19ef1)
	__obf_6b5c579c35cb3c44 := MachineID(__obf_b6d2ddd191a19ef1)
	return map[string]uint64{
		"id":         __obf_b6d2ddd191a19ef1,
		"msb":        __obf_f7b293a9f8ecf409,
		"time":       time,
		"sequence":   __obf_1e50fedd92e91188,
		"machine-id": __obf_6b5c579c35cb3c44,
	}
}
