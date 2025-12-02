// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// A Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package __obf_34ce7ee87a5aa6b7

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
	__obf_6f81ce04737c6bd5 *sync.Mutex
	__obf_309afc0b99b84762 int64
	__obf_e66900ce9e2ed50b int64
	__obf_22fa72eb809abbdc uint16
	__obf_34be90aa0845421d uint16
}

var (
	ErrStartTimeAhead   = errors.New("start time is ahead of now")
	ErrNoPrivateAddress = errors.New("no private ip address")
	ErrOverTimeLimit    = errors.New("over the time limit")
	ErrInvalidMachineID = errors.New("invalid machine id")
)

var __obf_9ca00d6957ade7c1 = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func NewSonyflake(__obf_a880e15c7070a691 Settings) (*Sonyflake, error) {
	if __obf_a880e15c7070a691.StartTime.After(time.Now()) {
		return nil, ErrStartTimeAhead
	}
	__obf_5b23d2a7003cadc7 := new(Sonyflake)
	__obf_5b23d2a7003cadc7.__obf_6f81ce04737c6bd5 = new(sync.Mutex)
	__obf_5b23d2a7003cadc7.__obf_22fa72eb809abbdc = uint16(1<<BitLenSequence - 1)

	if __obf_a880e15c7070a691.StartTime.IsZero() {
		__obf_5b23d2a7003cadc7.__obf_309afc0b99b84762 = __obf_e11c722b2a444110(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		__obf_5b23d2a7003cadc7.__obf_309afc0b99b84762 = __obf_e11c722b2a444110(__obf_a880e15c7070a691.StartTime)
	}

	var __obf_0a414fcc6889bc71 error
	if __obf_a880e15c7070a691.MachineID == nil {
		__obf_5b23d2a7003cadc7.__obf_34be90aa0845421d, __obf_0a414fcc6889bc71 = __obf_bedef3cec536fabf(__obf_9ca00d6957ade7c1)
	} else {
		__obf_5b23d2a7003cadc7.__obf_34be90aa0845421d, __obf_0a414fcc6889bc71 = __obf_a880e15c7070a691.MachineID()
	}
	if __obf_0a414fcc6889bc71 != nil {
		return nil, __obf_0a414fcc6889bc71
	}

	if __obf_a880e15c7070a691.CheckMachineID != nil && !__obf_a880e15c7070a691.CheckMachineID(__obf_5b23d2a7003cadc7.__obf_34be90aa0845421d) {
		return nil, ErrInvalidMachineID
	}

	return __obf_5b23d2a7003cadc7, nil
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
func (__obf_5b23d2a7003cadc7 *Sonyflake) NextID() (string, error) {
	const __obf_a696d1a1f61acc0a = uint16(1<<BitLenSequence - 1)
	__obf_5b23d2a7003cadc7.__obf_6f81ce04737c6bd5.
		Lock()
	defer __obf_5b23d2a7003cadc7.__obf_6f81ce04737c6bd5.Unlock()
	__obf_a95a9deeb2be55b3 := __obf_847c5a77e7a52193(__obf_5b23d2a7003cadc7.__obf_309afc0b99b84762)
	if __obf_5b23d2a7003cadc7.__obf_e66900ce9e2ed50b < __obf_a95a9deeb2be55b3 {
		__obf_5b23d2a7003cadc7.__obf_e66900ce9e2ed50b = __obf_a95a9deeb2be55b3
		__obf_5b23d2a7003cadc7.

			// sf.elapsedTime >= current
			__obf_22fa72eb809abbdc = 0
	} else {
		__obf_5b23d2a7003cadc7.__obf_22fa72eb809abbdc = (__obf_5b23d2a7003cadc7.__obf_22fa72eb809abbdc + 1) & __obf_a696d1a1f61acc0a
		if __obf_5b23d2a7003cadc7.__obf_22fa72eb809abbdc == 0 {
			__obf_5b23d2a7003cadc7.__obf_e66900ce9e2ed50b++
			__obf_3a2d71bbf09b008d := __obf_5b23d2a7003cadc7.__obf_e66900ce9e2ed50b - __obf_a95a9deeb2be55b3
			time.Sleep(__obf_c180152fa85d5659((__obf_3a2d71bbf09b008d)))
		}
	}

	return __obf_5b23d2a7003cadc7.__obf_3d035dc64b7eb663()
}

const __obf_416c9887773b83f1 = 1e7 // nsec, i.e. 10 msec

func __obf_e11c722b2a444110(__obf_438b07c9285724d4 time.Time) int64 {
	return __obf_438b07c9285724d4.UTC().UnixNano() / __obf_416c9887773b83f1
}

func __obf_847c5a77e7a52193(__obf_309afc0b99b84762 int64) int64 {
	return __obf_e11c722b2a444110(time.Now()) - __obf_309afc0b99b84762
}

func __obf_c180152fa85d5659(__obf_3a2d71bbf09b008d int64) time.Duration {
	return time.Duration(__obf_3a2d71bbf09b008d*__obf_416c9887773b83f1) -
		time.Duration(time.Now().UTC().UnixNano()%__obf_416c9887773b83f1)
}

func (__obf_5b23d2a7003cadc7 *Sonyflake) __obf_3d035dc64b7eb663() (string, error) {
	if __obf_5b23d2a7003cadc7.__obf_e66900ce9e2ed50b >= 1<<BitLenTime {
		return "", ErrOverTimeLimit
	}

	return fmt.Sprintf("%X", uint64(__obf_5b23d2a7003cadc7.__obf_e66900ce9e2ed50b)<<(BitLenSequence+BitLenMachineID)|
		uint64(__obf_5b23d2a7003cadc7.__obf_22fa72eb809abbdc)<<BitLenMachineID|
		uint64(__obf_5b23d2a7003cadc7.__obf_34be90aa0845421d)), nil
}

func __obf_199c85d57143e184(__obf_a58e221375b19d5b InterfaceAddrs) (net.IP, error) {
	__obf_840f743f8d0cd2c9, __obf_0a414fcc6889bc71 := __obf_a58e221375b19d5b()
	if __obf_0a414fcc6889bc71 != nil {
		return nil, __obf_0a414fcc6889bc71
	}

	for _, __obf_712740027bb61501 := range __obf_840f743f8d0cd2c9 {
		__obf_303c45c4553a6cdc, __obf_04415e5e48fee761 := __obf_712740027bb61501.(*net.IPNet)
		if !__obf_04415e5e48fee761 || __obf_303c45c4553a6cdc.IP.IsLoopback() {
			continue
		}
		__obf_03f650ded10700a4 := __obf_303c45c4553a6cdc.IP.To4()
		if __obf_591cbfa764f640b9(__obf_03f650ded10700a4) {
			return __obf_03f650ded10700a4, nil
		}
	}
	return nil, ErrNoPrivateAddress
}

func __obf_591cbfa764f640b9(__obf_03f650ded10700a4 net.IP) bool {
	return __obf_03f650ded10700a4 != nil &&
		(__obf_03f650ded10700a4[0] == 10 || __obf_03f650ded10700a4[0] == 172 && (__obf_03f650ded10700a4[1] >= 16 && __obf_03f650ded10700a4[1] < 32) || __obf_03f650ded10700a4[0] == 192 && __obf_03f650ded10700a4[1] == 168)
}

func __obf_bedef3cec536fabf(__obf_a58e221375b19d5b InterfaceAddrs) (uint16, error) {
	__obf_03f650ded10700a4, __obf_0a414fcc6889bc71 := __obf_199c85d57143e184(__obf_a58e221375b19d5b)
	if __obf_0a414fcc6889bc71 != nil {
		return 0, __obf_0a414fcc6889bc71
	}

	return uint16(__obf_03f650ded10700a4[2])<<8 + uint16(__obf_03f650ded10700a4[3]), nil
}

// ElapsedTime returns the elapsed time when the given Sonyflake ID was generated.
func ElapsedTime(__obf_f60c7174b4657d83 uint64) time.Duration {
	return time.Duration(__obf_e66900ce9e2ed50b(__obf_f60c7174b4657d83) * __obf_416c9887773b83f1)
}

func __obf_e66900ce9e2ed50b(__obf_f60c7174b4657d83 uint64) uint64 {
	return __obf_f60c7174b4657d83 >> (BitLenSequence + BitLenMachineID)
}

// SequenceNumber returns the sequence number of a Sonyflake ID.
func SequenceNumber(__obf_f60c7174b4657d83 uint64) uint64 {
	const __obf_a696d1a1f61acc0a = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	return __obf_f60c7174b4657d83 & __obf_a696d1a1f61acc0a >> BitLenMachineID
}

// MachineID returns the machine ID of a Sonyflake ID.
func MachineID(__obf_f60c7174b4657d83 uint64) uint64 {
	const __obf_8b164caeac9d09d6 = uint64(1<<BitLenMachineID - 1)
	return __obf_f60c7174b4657d83 & __obf_8b164caeac9d09d6
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(__obf_f60c7174b4657d83 uint64) map[string]uint64 {
	__obf_17506d3739354def := __obf_f60c7174b4657d83 >> 63
	time := __obf_e66900ce9e2ed50b(__obf_f60c7174b4657d83)
	__obf_22fa72eb809abbdc := SequenceNumber(__obf_f60c7174b4657d83)
	__obf_34be90aa0845421d := MachineID(__obf_f60c7174b4657d83)
	return map[string]uint64{
		"id":         __obf_f60c7174b4657d83,
		"msb":        __obf_17506d3739354def,
		"time":       time,
		"sequence":   __obf_22fa72eb809abbdc,
		"machine-id": __obf_34be90aa0845421d,
	}
}
