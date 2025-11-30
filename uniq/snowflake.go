package __obf_4d7a51f8e2f0494a

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_2103bbb91aba9b10 int64
	__obf_45f6ec9d07383a8c int64
	__obf_de8c2d4f1a1d02ed int64
	__obf_5d51e76f7973f04b int64
	__obf_f1cabaa526197d15 *sync.Mutex
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
func NewIdWorker(__obf_b01cd8c5178ba3dd int64) (__obf_6bac90a21c725eff *IDWorker, __obf_38c5d446257f8d9a error) {
	__obf_6bac90a21c725eff = new(IDWorker)
	__obf_6bac90a21c725eff.__obf_5d51e76f7973f04b = __obf_ce7052d902a49d7a()

	if __obf_b01cd8c5178ba3dd > __obf_6bac90a21c725eff.__obf_5d51e76f7973f04b || __obf_b01cd8c5178ba3dd < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_6bac90a21c725eff.__obf_2103bbb91aba9b10 = __obf_b01cd8c5178ba3dd
	__obf_6bac90a21c725eff.__obf_45f6ec9d07383a8c = -1
	__obf_6bac90a21c725eff.__obf_de8c2d4f1a1d02ed = 0
	__obf_6bac90a21c725eff.__obf_f1cabaa526197d15 = new(sync.Mutex)
	return __obf_6bac90a21c725eff, nil
}

func __obf_ce7052d902a49d7a() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_1b67fc205a49905d() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_6bac90a21c725eff *IDWorker) __obf_d4a5a084d81da9aa() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_6bac90a21c725eff *IDWorker) __obf_f79ebadf6d8f59e6(__obf_e44f0a7bde3d82a4 int64) int64 {
	__obf_8a68cc6e94529446 := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_8a68cc6e94529446 <= __obf_e44f0a7bde3d82a4 {
			__obf_8a68cc6e94529446 = __obf_6bac90a21c725eff.__obf_d4a5a084d81da9aa()
		} else {
			break
		}
	}
	return __obf_8a68cc6e94529446
}

// NewId Func: Generate next id
func (__obf_6bac90a21c725eff *IDWorker) NextID() (__obf_721283c42292359f string, __obf_38c5d446257f8d9a error) {
	__obf_6bac90a21c725eff.__obf_f1cabaa526197d15.
		Lock()
	defer __obf_6bac90a21c725eff.__obf_f1cabaa526197d15.Unlock()
	__obf_8a68cc6e94529446 := __obf_6bac90a21c725eff.__obf_d4a5a084d81da9aa()
	if __obf_8a68cc6e94529446 == __obf_6bac90a21c725eff.__obf_45f6ec9d07383a8c {
		__obf_6bac90a21c725eff.__obf_de8c2d4f1a1d02ed = (__obf_6bac90a21c725eff.__obf_de8c2d4f1a1d02ed + 1) & CSequenceMask
		if __obf_6bac90a21c725eff.__obf_de8c2d4f1a1d02ed == 0 {
			__obf_8a68cc6e94529446 = __obf_6bac90a21c725eff.__obf_f79ebadf6d8f59e6(__obf_8a68cc6e94529446)
		}
	} else {
		__obf_6bac90a21c725eff.__obf_de8c2d4f1a1d02ed = 0
	}

	if __obf_8a68cc6e94529446 < __obf_6bac90a21c725eff.__obf_45f6ec9d07383a8c {
		__obf_38c5d446257f8d9a = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_38c5d446257f8d9a
	}
	__obf_6bac90a21c725eff.__obf_45f6ec9d07383a8c = __obf_8a68cc6e94529446
	__obf_8a68cc6e94529446 = (__obf_8a68cc6e94529446-CEpoch)<<CTimeStampShift | __obf_6bac90a21c725eff.__obf_2103bbb91aba9b10<<CWorkerIdShift | __obf_6bac90a21c725eff.__obf_de8c2d4f1a1d02ed
	return fmt.Sprintf("%X", __obf_8a68cc6e94529446), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_721283c42292359f int64) (__obf_d5aaddb5c5a3e313 time.Time, __obf_8a68cc6e94529446 int64, __obf_2103bbb91aba9b10 int64, __obf_4e18f31cbea3d556 int64) {
	__obf_4e18f31cbea3d556 = __obf_721283c42292359f & CSequenceMask
	__obf_2103bbb91aba9b10 = (__obf_721283c42292359f >> CWorkerIdShift) & CMaxWorker
	__obf_8a68cc6e94529446 = (__obf_721283c42292359f >> CTimeStampShift) + CEpoch
	__obf_d5aaddb5c5a3e313 = time.Unix(__obf_8a68cc6e94529446/1000, (__obf_8a68cc6e94529446%1000)*1000000)
	return
}
