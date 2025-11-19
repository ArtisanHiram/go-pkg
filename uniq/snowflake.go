package __obf_7913809aab6c8423

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_60e6c20fe32b0bb5 int64
	__obf_448967efa85ecf26 int64
	__obf_007ec6572a283fb6 int64
	__obf_92b61aa378459ea7 int64
	__obf_b05d58c82e8ad833 *sync.Mutex
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
func NewIdWorker(__obf_f851380ef2c87315 int64) (__obf_e62433baa09fafa8 *IDWorker, __obf_e62cd9e417a87ee7 error) {
	__obf_e62433baa09fafa8 = new(IDWorker)

	__obf_e62433baa09fafa8.__obf_92b61aa378459ea7 = __obf_ba213b7c556b7e4f()

	if __obf_f851380ef2c87315 > __obf_e62433baa09fafa8.__obf_92b61aa378459ea7 || __obf_f851380ef2c87315 < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_e62433baa09fafa8.__obf_60e6c20fe32b0bb5 = __obf_f851380ef2c87315
	__obf_e62433baa09fafa8.__obf_448967efa85ecf26 = -1
	__obf_e62433baa09fafa8.__obf_007ec6572a283fb6 = 0
	__obf_e62433baa09fafa8.__obf_b05d58c82e8ad833 = new(sync.Mutex)
	return __obf_e62433baa09fafa8, nil
}

func __obf_ba213b7c556b7e4f() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_24922beb5801f1f9() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_e62433baa09fafa8 *IDWorker) __obf_d2e526e4f19a1a37() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_e62433baa09fafa8 *IDWorker) __obf_f0c7aac45cb8b0d0(__obf_93b3847cb15feb85 int64) int64 {
	__obf_6369016dedfed10f := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_6369016dedfed10f <= __obf_93b3847cb15feb85 {
			__obf_6369016dedfed10f = __obf_e62433baa09fafa8.__obf_d2e526e4f19a1a37()
		} else {
			break
		}
	}
	return __obf_6369016dedfed10f
}

// NewId Func: Generate next id
func (__obf_e62433baa09fafa8 *IDWorker) NextID() (__obf_1dfe314328e20778 string, __obf_e62cd9e417a87ee7 error) {
	__obf_e62433baa09fafa8.__obf_b05d58c82e8ad833.Lock()
	defer __obf_e62433baa09fafa8.__obf_b05d58c82e8ad833.Unlock()
	__obf_6369016dedfed10f := __obf_e62433baa09fafa8.__obf_d2e526e4f19a1a37()
	if __obf_6369016dedfed10f == __obf_e62433baa09fafa8.__obf_448967efa85ecf26 {
		__obf_e62433baa09fafa8.__obf_007ec6572a283fb6 = (__obf_e62433baa09fafa8.__obf_007ec6572a283fb6 + 1) & CSequenceMask
		if __obf_e62433baa09fafa8.__obf_007ec6572a283fb6 == 0 {
			__obf_6369016dedfed10f = __obf_e62433baa09fafa8.__obf_f0c7aac45cb8b0d0(__obf_6369016dedfed10f)
		}
	} else {
		__obf_e62433baa09fafa8.__obf_007ec6572a283fb6 = 0
	}

	if __obf_6369016dedfed10f < __obf_e62433baa09fafa8.__obf_448967efa85ecf26 {
		__obf_e62cd9e417a87ee7 = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_e62cd9e417a87ee7
	}
	__obf_e62433baa09fafa8.__obf_448967efa85ecf26 = __obf_6369016dedfed10f
	__obf_6369016dedfed10f = (__obf_6369016dedfed10f-CEpoch)<<CTimeStampShift | __obf_e62433baa09fafa8.__obf_60e6c20fe32b0bb5<<CWorkerIdShift | __obf_e62433baa09fafa8.__obf_007ec6572a283fb6
	return fmt.Sprintf("%X", __obf_6369016dedfed10f), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_1dfe314328e20778 int64) (__obf_f154b37f4911a720 time.Time, __obf_6369016dedfed10f int64, __obf_60e6c20fe32b0bb5 int64, __obf_cde251f9028a615e int64) {
	__obf_cde251f9028a615e = __obf_1dfe314328e20778 & CSequenceMask
	__obf_60e6c20fe32b0bb5 = (__obf_1dfe314328e20778 >> CWorkerIdShift) & CMaxWorker
	__obf_6369016dedfed10f = (__obf_1dfe314328e20778 >> CTimeStampShift) + CEpoch
	__obf_f154b37f4911a720 = time.Unix(__obf_6369016dedfed10f/1000, (__obf_6369016dedfed10f%1000)*1000000)
	return
}
