package __obf_07f0876faa0cf68e

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_03ad2d4376081c35 int64
	__obf_a9bc1d870de2ef59 int64
	__obf_dd9e42b79b3ea456 int64
	__obf_9b7630cb86161dd3 int64
	__obf_1757885ca6bb9ea3 *sync.Mutex
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
func NewIdWorker(__obf_85c539bcaa75c159 int64) (__obf_51e204f219d1ded9 *IDWorker, __obf_116d5a689b79a523 error) {
	__obf_51e204f219d1ded9 = new(IDWorker)
	__obf_51e204f219d1ded9.__obf_9b7630cb86161dd3 = __obf_0404352b53e35b8e()

	if __obf_85c539bcaa75c159 > __obf_51e204f219d1ded9.__obf_9b7630cb86161dd3 || __obf_85c539bcaa75c159 < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_51e204f219d1ded9.__obf_03ad2d4376081c35 = __obf_85c539bcaa75c159
	__obf_51e204f219d1ded9.__obf_a9bc1d870de2ef59 = -1
	__obf_51e204f219d1ded9.__obf_dd9e42b79b3ea456 = 0
	__obf_51e204f219d1ded9.__obf_1757885ca6bb9ea3 = new(sync.Mutex)
	return __obf_51e204f219d1ded9, nil
}

func __obf_0404352b53e35b8e() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_66118011e0872635() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_51e204f219d1ded9 *IDWorker) __obf_7c07ff272d5b7f5d() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_51e204f219d1ded9 *IDWorker) __obf_f5e8870147112714(__obf_0843fde565cc42fc int64) int64 {
	__obf_b57f28b88e5243ec := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_b57f28b88e5243ec <= __obf_0843fde565cc42fc {
			__obf_b57f28b88e5243ec = __obf_51e204f219d1ded9.__obf_7c07ff272d5b7f5d()
		} else {
			break
		}
	}
	return __obf_b57f28b88e5243ec
}

// NewId Func: Generate next id
func (__obf_51e204f219d1ded9 *IDWorker) NextID() (__obf_28d51b2675aee39d string, __obf_116d5a689b79a523 error) {
	__obf_51e204f219d1ded9.__obf_1757885ca6bb9ea3.
		Lock()
	defer __obf_51e204f219d1ded9.__obf_1757885ca6bb9ea3.Unlock()
	__obf_b57f28b88e5243ec := __obf_51e204f219d1ded9.__obf_7c07ff272d5b7f5d()
	if __obf_b57f28b88e5243ec == __obf_51e204f219d1ded9.__obf_a9bc1d870de2ef59 {
		__obf_51e204f219d1ded9.__obf_dd9e42b79b3ea456 = (__obf_51e204f219d1ded9.__obf_dd9e42b79b3ea456 + 1) & CSequenceMask
		if __obf_51e204f219d1ded9.__obf_dd9e42b79b3ea456 == 0 {
			__obf_b57f28b88e5243ec = __obf_51e204f219d1ded9.__obf_f5e8870147112714(__obf_b57f28b88e5243ec)
		}
	} else {
		__obf_51e204f219d1ded9.__obf_dd9e42b79b3ea456 = 0
	}

	if __obf_b57f28b88e5243ec < __obf_51e204f219d1ded9.__obf_a9bc1d870de2ef59 {
		__obf_116d5a689b79a523 = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_116d5a689b79a523
	}
	__obf_51e204f219d1ded9.__obf_a9bc1d870de2ef59 = __obf_b57f28b88e5243ec
	__obf_b57f28b88e5243ec = (__obf_b57f28b88e5243ec-CEpoch)<<CTimeStampShift | __obf_51e204f219d1ded9.__obf_03ad2d4376081c35<<CWorkerIdShift | __obf_51e204f219d1ded9.__obf_dd9e42b79b3ea456
	return fmt.Sprintf("%X", __obf_b57f28b88e5243ec), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_28d51b2675aee39d int64) (__obf_13047ef7aa466d6e time.Time, __obf_b57f28b88e5243ec int64, __obf_03ad2d4376081c35 int64, __obf_3e8b135c5bd99e27 int64) {
	__obf_3e8b135c5bd99e27 = __obf_28d51b2675aee39d & CSequenceMask
	__obf_03ad2d4376081c35 = (__obf_28d51b2675aee39d >> CWorkerIdShift) & CMaxWorker
	__obf_b57f28b88e5243ec = (__obf_28d51b2675aee39d >> CTimeStampShift) + CEpoch
	__obf_13047ef7aa466d6e = time.Unix(__obf_b57f28b88e5243ec/1000, (__obf_b57f28b88e5243ec%1000)*1000000)
	return
}
