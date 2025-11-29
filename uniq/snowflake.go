package __obf_2f51f7d26a2bcdf8

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_572374c23004839f int64
	__obf_177a20ea7c90b986 int64
	__obf_692ac7e4829ff716 int64
	__obf_b2522fc1f399259e int64
	__obf_3de34439d6aac6d0 *sync.Mutex
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
func NewIdWorker(__obf_29a15d1d5a46cf6e int64) (__obf_94e1eb77f8ee7884 *IDWorker, __obf_37e8e388755917c0 error) {
	__obf_94e1eb77f8ee7884 = new(IDWorker)
	__obf_94e1eb77f8ee7884.__obf_b2522fc1f399259e = __obf_4b15d5f9ad0f6a9d()

	if __obf_29a15d1d5a46cf6e > __obf_94e1eb77f8ee7884.__obf_b2522fc1f399259e || __obf_29a15d1d5a46cf6e < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_94e1eb77f8ee7884.__obf_572374c23004839f = __obf_29a15d1d5a46cf6e
	__obf_94e1eb77f8ee7884.__obf_177a20ea7c90b986 = -1
	__obf_94e1eb77f8ee7884.__obf_692ac7e4829ff716 = 0
	__obf_94e1eb77f8ee7884.__obf_3de34439d6aac6d0 = new(sync.Mutex)
	return __obf_94e1eb77f8ee7884, nil
}

func __obf_4b15d5f9ad0f6a9d() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_5676cee0f26429b1() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_94e1eb77f8ee7884 *IDWorker) __obf_9fa6ef7c8cfb58ed() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_94e1eb77f8ee7884 *IDWorker) __obf_136a041c3099ed85(__obf_bb82aaec650e053e int64) int64 {
	__obf_84e84031cdf6f5e0 := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_84e84031cdf6f5e0 <= __obf_bb82aaec650e053e {
			__obf_84e84031cdf6f5e0 = __obf_94e1eb77f8ee7884.__obf_9fa6ef7c8cfb58ed()
		} else {
			break
		}
	}
	return __obf_84e84031cdf6f5e0
}

// NewId Func: Generate next id
func (__obf_94e1eb77f8ee7884 *IDWorker) NextID() (__obf_8b246d7b3659af37 string, __obf_37e8e388755917c0 error) {
	__obf_94e1eb77f8ee7884.__obf_3de34439d6aac6d0.
		Lock()
	defer __obf_94e1eb77f8ee7884.__obf_3de34439d6aac6d0.Unlock()
	__obf_84e84031cdf6f5e0 := __obf_94e1eb77f8ee7884.__obf_9fa6ef7c8cfb58ed()
	if __obf_84e84031cdf6f5e0 == __obf_94e1eb77f8ee7884.__obf_177a20ea7c90b986 {
		__obf_94e1eb77f8ee7884.__obf_692ac7e4829ff716 = (__obf_94e1eb77f8ee7884.__obf_692ac7e4829ff716 + 1) & CSequenceMask
		if __obf_94e1eb77f8ee7884.__obf_692ac7e4829ff716 == 0 {
			__obf_84e84031cdf6f5e0 = __obf_94e1eb77f8ee7884.__obf_136a041c3099ed85(__obf_84e84031cdf6f5e0)
		}
	} else {
		__obf_94e1eb77f8ee7884.__obf_692ac7e4829ff716 = 0
	}

	if __obf_84e84031cdf6f5e0 < __obf_94e1eb77f8ee7884.__obf_177a20ea7c90b986 {
		__obf_37e8e388755917c0 = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_37e8e388755917c0
	}
	__obf_94e1eb77f8ee7884.__obf_177a20ea7c90b986 = __obf_84e84031cdf6f5e0
	__obf_84e84031cdf6f5e0 = (__obf_84e84031cdf6f5e0-CEpoch)<<CTimeStampShift | __obf_94e1eb77f8ee7884.__obf_572374c23004839f<<CWorkerIdShift | __obf_94e1eb77f8ee7884.__obf_692ac7e4829ff716
	return fmt.Sprintf("%X", __obf_84e84031cdf6f5e0), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_8b246d7b3659af37 int64) (__obf_61139133e991449d time.Time, __obf_84e84031cdf6f5e0 int64, __obf_572374c23004839f int64, __obf_9b13dc8d6dee3e68 int64) {
	__obf_9b13dc8d6dee3e68 = __obf_8b246d7b3659af37 & CSequenceMask
	__obf_572374c23004839f = (__obf_8b246d7b3659af37 >> CWorkerIdShift) & CMaxWorker
	__obf_84e84031cdf6f5e0 = (__obf_8b246d7b3659af37 >> CTimeStampShift) + CEpoch
	__obf_61139133e991449d = time.Unix(__obf_84e84031cdf6f5e0/1000, (__obf_84e84031cdf6f5e0%1000)*1000000)
	return
}
