package __obf_34ce7ee87a5aa6b7

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_2aaf4a5fa49bcfef int64
	__obf_b3d273a38238bf43 int64
	__obf_22fa72eb809abbdc int64
	__obf_799fd121f51aa374 int64
	__obf_d34d4a98122d480a *sync.Mutex
}

const (
	CEpoch         = 1474802888000
	CWorkerIdBits  = 10 // Num of WorkerId Bits
	CSenquenceBits = 12 // Num of Sequence Bits

	CWorkerIdShift  = 12
	CTimeStampShift = 22

	CSequenceMask = 0xfff // equal as getSequenceMask()
	CMaxWorker    = 0x3ff // equal as getMaxWorkerId()
)

// NewIdWorker Func: Generate NewIdWorker with Given workerid
func NewIdWorker(__obf_9522e8c0260c7361 int64) (__obf_101524b2cf02f2ee *IDWorker, __obf_0a414fcc6889bc71 error) {
	__obf_101524b2cf02f2ee = new(IDWorker)
	__obf_101524b2cf02f2ee.__obf_799fd121f51aa374 = __obf_6401801f2954e071()

	if __obf_9522e8c0260c7361 > __obf_101524b2cf02f2ee.__obf_799fd121f51aa374 || __obf_9522e8c0260c7361 < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_101524b2cf02f2ee.__obf_2aaf4a5fa49bcfef = __obf_9522e8c0260c7361
	__obf_101524b2cf02f2ee.__obf_b3d273a38238bf43 = -1
	__obf_101524b2cf02f2ee.__obf_22fa72eb809abbdc = 0
	__obf_101524b2cf02f2ee.__obf_d34d4a98122d480a = new(sync.Mutex)
	return __obf_101524b2cf02f2ee, nil
}

func __obf_6401801f2954e071() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_47f530478a8e72ef() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_101524b2cf02f2ee *IDWorker) __obf_8a430be2c9507dfc() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_101524b2cf02f2ee *IDWorker) __obf_09a1d0fcd5e28696(__obf_1407a7b2f67be97d int64) int64 {
	__obf_73efe4d53b8026a2 := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_73efe4d53b8026a2 <= __obf_1407a7b2f67be97d {
			__obf_73efe4d53b8026a2 = __obf_101524b2cf02f2ee.__obf_8a430be2c9507dfc()
		} else {
			break
		}
	}
	return __obf_73efe4d53b8026a2
}

// NewId Func: Generate next id
func (__obf_101524b2cf02f2ee *IDWorker) NextID() (__obf_f60c7174b4657d83 string, __obf_0a414fcc6889bc71 error) {
	__obf_101524b2cf02f2ee.__obf_d34d4a98122d480a.
		Lock()
	defer __obf_101524b2cf02f2ee.__obf_d34d4a98122d480a.Unlock()
	__obf_73efe4d53b8026a2 := __obf_101524b2cf02f2ee.__obf_8a430be2c9507dfc()
	if __obf_73efe4d53b8026a2 == __obf_101524b2cf02f2ee.__obf_b3d273a38238bf43 {
		__obf_101524b2cf02f2ee.__obf_22fa72eb809abbdc = (__obf_101524b2cf02f2ee.__obf_22fa72eb809abbdc + 1) & CSequenceMask
		if __obf_101524b2cf02f2ee.__obf_22fa72eb809abbdc == 0 {
			__obf_73efe4d53b8026a2 = __obf_101524b2cf02f2ee.__obf_09a1d0fcd5e28696(__obf_73efe4d53b8026a2)
		}
	} else {
		__obf_101524b2cf02f2ee.__obf_22fa72eb809abbdc = 0
	}

	if __obf_73efe4d53b8026a2 < __obf_101524b2cf02f2ee.__obf_b3d273a38238bf43 {
		__obf_0a414fcc6889bc71 = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_0a414fcc6889bc71
	}
	__obf_101524b2cf02f2ee.__obf_b3d273a38238bf43 = __obf_73efe4d53b8026a2
	__obf_73efe4d53b8026a2 = (__obf_73efe4d53b8026a2-CEpoch)<<CTimeStampShift | __obf_101524b2cf02f2ee.__obf_2aaf4a5fa49bcfef<<CWorkerIdShift | __obf_101524b2cf02f2ee.__obf_22fa72eb809abbdc
	return fmt.Sprintf("%X", __obf_73efe4d53b8026a2), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_f60c7174b4657d83 int64) (__obf_438b07c9285724d4 time.Time, __obf_73efe4d53b8026a2 int64, __obf_2aaf4a5fa49bcfef int64, __obf_11dd36c6a2360aee int64) {
	__obf_11dd36c6a2360aee = __obf_f60c7174b4657d83 & CSequenceMask
	__obf_2aaf4a5fa49bcfef = (__obf_f60c7174b4657d83 >> CWorkerIdShift) & CMaxWorker
	__obf_73efe4d53b8026a2 = (__obf_f60c7174b4657d83 >> CTimeStampShift) + CEpoch
	__obf_438b07c9285724d4 = time.Unix(__obf_73efe4d53b8026a2/1000, (__obf_73efe4d53b8026a2%1000)*1000000)
	return
}
