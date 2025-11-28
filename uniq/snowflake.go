package __obf_7d8ac56be2e11a40

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_1d83e10d310a545a int64
	__obf_09d52451782f93df int64
	__obf_f06c729260d81c2a int64
	__obf_166cfd7925898bb1 int64
	__obf_857655b064a2556d *sync.Mutex
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
func NewIdWorker(__obf_b92b5159b349c507 int64) (__obf_1dd25ae9f65a6d11 *IDWorker, __obf_3f3f828d9e46d714 error) {
	__obf_1dd25ae9f65a6d11 = new(IDWorker)

	__obf_1dd25ae9f65a6d11.__obf_166cfd7925898bb1 = __obf_b3fcb1129a86c6bf()

	if __obf_b92b5159b349c507 > __obf_1dd25ae9f65a6d11.__obf_166cfd7925898bb1 || __obf_b92b5159b349c507 < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_1dd25ae9f65a6d11.__obf_1d83e10d310a545a = __obf_b92b5159b349c507
	__obf_1dd25ae9f65a6d11.__obf_09d52451782f93df = -1
	__obf_1dd25ae9f65a6d11.__obf_f06c729260d81c2a = 0
	__obf_1dd25ae9f65a6d11.__obf_857655b064a2556d = new(sync.Mutex)
	return __obf_1dd25ae9f65a6d11, nil
}

func __obf_b3fcb1129a86c6bf() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_8423c407d94b6bf2() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_1dd25ae9f65a6d11 *IDWorker) __obf_cb6e39c621b81f90() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_1dd25ae9f65a6d11 *IDWorker) __obf_ba00570227a89851(__obf_16a4dac3d3ddde5b int64) int64 {
	__obf_0e785863e2da461b := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_0e785863e2da461b <= __obf_16a4dac3d3ddde5b {
			__obf_0e785863e2da461b = __obf_1dd25ae9f65a6d11.__obf_cb6e39c621b81f90()
		} else {
			break
		}
	}
	return __obf_0e785863e2da461b
}

// NewId Func: Generate next id
func (__obf_1dd25ae9f65a6d11 *IDWorker) NextID() (__obf_6cfbdbb264317300 string, __obf_3f3f828d9e46d714 error) {
	__obf_1dd25ae9f65a6d11.__obf_857655b064a2556d.Lock()
	defer __obf_1dd25ae9f65a6d11.__obf_857655b064a2556d.Unlock()
	__obf_0e785863e2da461b := __obf_1dd25ae9f65a6d11.__obf_cb6e39c621b81f90()
	if __obf_0e785863e2da461b == __obf_1dd25ae9f65a6d11.__obf_09d52451782f93df {
		__obf_1dd25ae9f65a6d11.__obf_f06c729260d81c2a = (__obf_1dd25ae9f65a6d11.__obf_f06c729260d81c2a + 1) & CSequenceMask
		if __obf_1dd25ae9f65a6d11.__obf_f06c729260d81c2a == 0 {
			__obf_0e785863e2da461b = __obf_1dd25ae9f65a6d11.__obf_ba00570227a89851(__obf_0e785863e2da461b)
		}
	} else {
		__obf_1dd25ae9f65a6d11.__obf_f06c729260d81c2a = 0
	}

	if __obf_0e785863e2da461b < __obf_1dd25ae9f65a6d11.__obf_09d52451782f93df {
		__obf_3f3f828d9e46d714 = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_3f3f828d9e46d714
	}
	__obf_1dd25ae9f65a6d11.__obf_09d52451782f93df = __obf_0e785863e2da461b
	__obf_0e785863e2da461b = (__obf_0e785863e2da461b-CEpoch)<<CTimeStampShift | __obf_1dd25ae9f65a6d11.__obf_1d83e10d310a545a<<CWorkerIdShift | __obf_1dd25ae9f65a6d11.__obf_f06c729260d81c2a
	return fmt.Sprintf("%X", __obf_0e785863e2da461b), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_6cfbdbb264317300 int64) (__obf_680f17a8e930eb64 time.Time, __obf_0e785863e2da461b int64, __obf_1d83e10d310a545a int64, __obf_d1fe27cbd63fb1b9 int64) {
	__obf_d1fe27cbd63fb1b9 = __obf_6cfbdbb264317300 & CSequenceMask
	__obf_1d83e10d310a545a = (__obf_6cfbdbb264317300 >> CWorkerIdShift) & CMaxWorker
	__obf_0e785863e2da461b = (__obf_6cfbdbb264317300 >> CTimeStampShift) + CEpoch
	__obf_680f17a8e930eb64 = time.Unix(__obf_0e785863e2da461b/1000, (__obf_0e785863e2da461b%1000)*1000000)
	return
}
