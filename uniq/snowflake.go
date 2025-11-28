package __obf_8fe28252c1b01dfe

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_108fc1a52cfb2f39 int64
	__obf_50f043c0a3d7a62e int64
	__obf_1e50fedd92e91188 int64
	__obf_3586593158f3c3a7 int64
	__obf_740b5a1f5e6fb2d0 *sync.Mutex
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
func NewIdWorker(__obf_f69b6787b790a911 int64) (__obf_46ef2c89b1192c8f *IDWorker, __obf_1082dd6e56192e3a error) {
	__obf_46ef2c89b1192c8f = new(IDWorker)

	__obf_46ef2c89b1192c8f.__obf_3586593158f3c3a7 = __obf_a88cb6ff01a15459()

	if __obf_f69b6787b790a911 > __obf_46ef2c89b1192c8f.__obf_3586593158f3c3a7 || __obf_f69b6787b790a911 < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_46ef2c89b1192c8f.__obf_108fc1a52cfb2f39 = __obf_f69b6787b790a911
	__obf_46ef2c89b1192c8f.__obf_50f043c0a3d7a62e = -1
	__obf_46ef2c89b1192c8f.__obf_1e50fedd92e91188 = 0
	__obf_46ef2c89b1192c8f.__obf_740b5a1f5e6fb2d0 = new(sync.Mutex)
	return __obf_46ef2c89b1192c8f, nil
}

func __obf_a88cb6ff01a15459() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_5931d64ec382971a() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_46ef2c89b1192c8f *IDWorker) __obf_248a0612c610a820() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_46ef2c89b1192c8f *IDWorker) __obf_fcea7a5b3679749c(__obf_a1064549ea5e3728 int64) int64 {
	__obf_c5cd88acb522108d := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_c5cd88acb522108d <= __obf_a1064549ea5e3728 {
			__obf_c5cd88acb522108d = __obf_46ef2c89b1192c8f.__obf_248a0612c610a820()
		} else {
			break
		}
	}
	return __obf_c5cd88acb522108d
}

// NewId Func: Generate next id
func (__obf_46ef2c89b1192c8f *IDWorker) NextID() (__obf_b6d2ddd191a19ef1 string, __obf_1082dd6e56192e3a error) {
	__obf_46ef2c89b1192c8f.__obf_740b5a1f5e6fb2d0.Lock()
	defer __obf_46ef2c89b1192c8f.__obf_740b5a1f5e6fb2d0.Unlock()
	__obf_c5cd88acb522108d := __obf_46ef2c89b1192c8f.__obf_248a0612c610a820()
	if __obf_c5cd88acb522108d == __obf_46ef2c89b1192c8f.__obf_50f043c0a3d7a62e {
		__obf_46ef2c89b1192c8f.__obf_1e50fedd92e91188 = (__obf_46ef2c89b1192c8f.__obf_1e50fedd92e91188 + 1) & CSequenceMask
		if __obf_46ef2c89b1192c8f.__obf_1e50fedd92e91188 == 0 {
			__obf_c5cd88acb522108d = __obf_46ef2c89b1192c8f.__obf_fcea7a5b3679749c(__obf_c5cd88acb522108d)
		}
	} else {
		__obf_46ef2c89b1192c8f.__obf_1e50fedd92e91188 = 0
	}

	if __obf_c5cd88acb522108d < __obf_46ef2c89b1192c8f.__obf_50f043c0a3d7a62e {
		__obf_1082dd6e56192e3a = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_1082dd6e56192e3a
	}
	__obf_46ef2c89b1192c8f.__obf_50f043c0a3d7a62e = __obf_c5cd88acb522108d
	__obf_c5cd88acb522108d = (__obf_c5cd88acb522108d-CEpoch)<<CTimeStampShift | __obf_46ef2c89b1192c8f.__obf_108fc1a52cfb2f39<<CWorkerIdShift | __obf_46ef2c89b1192c8f.__obf_1e50fedd92e91188
	return fmt.Sprintf("%X", __obf_c5cd88acb522108d), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_b6d2ddd191a19ef1 int64) (__obf_8bbd0f9c7d55c1cc time.Time, __obf_c5cd88acb522108d int64, __obf_108fc1a52cfb2f39 int64, __obf_919a740f61871f72 int64) {
	__obf_919a740f61871f72 = __obf_b6d2ddd191a19ef1 & CSequenceMask
	__obf_108fc1a52cfb2f39 = (__obf_b6d2ddd191a19ef1 >> CWorkerIdShift) & CMaxWorker
	__obf_c5cd88acb522108d = (__obf_b6d2ddd191a19ef1 >> CTimeStampShift) + CEpoch
	__obf_8bbd0f9c7d55c1cc = time.Unix(__obf_c5cd88acb522108d/1000, (__obf_c5cd88acb522108d%1000)*1000000)
	return
}
