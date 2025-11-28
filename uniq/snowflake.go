package __obf_e2239f4853c61591

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_cae27ff8e22c5a84 int64
	__obf_693731192cf64b64 int64
	__obf_4a953d8c26aa0257 int64
	__obf_fb61d44e6080a92f int64
	__obf_5fbd0e3204eb4961 *sync.Mutex
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
func NewIdWorker(__obf_197f2a4e3d61109b int64) (__obf_2622d83baefeb562 *IDWorker, __obf_bb9dc4c4f445b22f error) {
	__obf_2622d83baefeb562 = new(IDWorker)

	__obf_2622d83baefeb562.__obf_fb61d44e6080a92f = __obf_e7b4092dda317667()

	if __obf_197f2a4e3d61109b > __obf_2622d83baefeb562.__obf_fb61d44e6080a92f || __obf_197f2a4e3d61109b < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_2622d83baefeb562.__obf_cae27ff8e22c5a84 = __obf_197f2a4e3d61109b
	__obf_2622d83baefeb562.__obf_693731192cf64b64 = -1
	__obf_2622d83baefeb562.__obf_4a953d8c26aa0257 = 0
	__obf_2622d83baefeb562.__obf_5fbd0e3204eb4961 = new(sync.Mutex)
	return __obf_2622d83baefeb562, nil
}

func __obf_e7b4092dda317667() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_264cad9c2def546d() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_2622d83baefeb562 *IDWorker) __obf_9902db9dfedf3d35() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_2622d83baefeb562 *IDWorker) __obf_fb6623be7931b1e1(__obf_6f2f4287d7522d7a int64) int64 {
	__obf_e2ad35a6406441f1 := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_e2ad35a6406441f1 <= __obf_6f2f4287d7522d7a {
			__obf_e2ad35a6406441f1 = __obf_2622d83baefeb562.__obf_9902db9dfedf3d35()
		} else {
			break
		}
	}
	return __obf_e2ad35a6406441f1
}

// NewId Func: Generate next id
func (__obf_2622d83baefeb562 *IDWorker) NextID() (__obf_cee82df7c007227d string, __obf_bb9dc4c4f445b22f error) {
	__obf_2622d83baefeb562.__obf_5fbd0e3204eb4961.Lock()
	defer __obf_2622d83baefeb562.__obf_5fbd0e3204eb4961.Unlock()
	__obf_e2ad35a6406441f1 := __obf_2622d83baefeb562.__obf_9902db9dfedf3d35()
	if __obf_e2ad35a6406441f1 == __obf_2622d83baefeb562.__obf_693731192cf64b64 {
		__obf_2622d83baefeb562.__obf_4a953d8c26aa0257 = (__obf_2622d83baefeb562.__obf_4a953d8c26aa0257 + 1) & CSequenceMask
		if __obf_2622d83baefeb562.__obf_4a953d8c26aa0257 == 0 {
			__obf_e2ad35a6406441f1 = __obf_2622d83baefeb562.__obf_fb6623be7931b1e1(__obf_e2ad35a6406441f1)
		}
	} else {
		__obf_2622d83baefeb562.__obf_4a953d8c26aa0257 = 0
	}

	if __obf_e2ad35a6406441f1 < __obf_2622d83baefeb562.__obf_693731192cf64b64 {
		__obf_bb9dc4c4f445b22f = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_bb9dc4c4f445b22f
	}
	__obf_2622d83baefeb562.__obf_693731192cf64b64 = __obf_e2ad35a6406441f1
	__obf_e2ad35a6406441f1 = (__obf_e2ad35a6406441f1-CEpoch)<<CTimeStampShift | __obf_2622d83baefeb562.__obf_cae27ff8e22c5a84<<CWorkerIdShift | __obf_2622d83baefeb562.__obf_4a953d8c26aa0257
	return fmt.Sprintf("%X", __obf_e2ad35a6406441f1), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_cee82df7c007227d int64) (__obf_de192fc01291fe85 time.Time, __obf_e2ad35a6406441f1 int64, __obf_cae27ff8e22c5a84 int64, __obf_64df1b95907fd3a2 int64) {
	__obf_64df1b95907fd3a2 = __obf_cee82df7c007227d & CSequenceMask
	__obf_cae27ff8e22c5a84 = (__obf_cee82df7c007227d >> CWorkerIdShift) & CMaxWorker
	__obf_e2ad35a6406441f1 = (__obf_cee82df7c007227d >> CTimeStampShift) + CEpoch
	__obf_de192fc01291fe85 = time.Unix(__obf_e2ad35a6406441f1/1000, (__obf_e2ad35a6406441f1%1000)*1000000)
	return
}
