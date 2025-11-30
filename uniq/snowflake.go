package __obf_3747a7e09ff475ee

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_99a2aa55cbd819ac int64
	__obf_0c18bdba049ba1e6 int64
	__obf_32d285c02de19ff8 int64
	__obf_ed2e11c24e76761c int64
	__obf_218ba8e1456a7869 *sync.Mutex
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
func NewIdWorker(__obf_ef6a9c9e46ca2e5b int64) (__obf_7f538276625ed09b *IDWorker, __obf_57cef2f010b4ce9a error) {
	__obf_7f538276625ed09b = new(IDWorker)
	__obf_7f538276625ed09b.__obf_ed2e11c24e76761c = __obf_fc91aa62255a86f1()

	if __obf_ef6a9c9e46ca2e5b > __obf_7f538276625ed09b.__obf_ed2e11c24e76761c || __obf_ef6a9c9e46ca2e5b < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_7f538276625ed09b.__obf_99a2aa55cbd819ac = __obf_ef6a9c9e46ca2e5b
	__obf_7f538276625ed09b.__obf_0c18bdba049ba1e6 = -1
	__obf_7f538276625ed09b.__obf_32d285c02de19ff8 = 0
	__obf_7f538276625ed09b.__obf_218ba8e1456a7869 = new(sync.Mutex)
	return __obf_7f538276625ed09b, nil
}

func __obf_fc91aa62255a86f1() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_07a18b56d9b26260() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_7f538276625ed09b *IDWorker) __obf_04d7f434df35909c() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_7f538276625ed09b *IDWorker) __obf_2c1ebd77380eb657(__obf_03fa9ff55948b1c1 int64) int64 {
	__obf_5fa8637e08c90073 := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_5fa8637e08c90073 <= __obf_03fa9ff55948b1c1 {
			__obf_5fa8637e08c90073 = __obf_7f538276625ed09b.__obf_04d7f434df35909c()
		} else {
			break
		}
	}
	return __obf_5fa8637e08c90073
}

// NewId Func: Generate next id
func (__obf_7f538276625ed09b *IDWorker) NextID() (__obf_413d3a76ff8ab79e string, __obf_57cef2f010b4ce9a error) {
	__obf_7f538276625ed09b.__obf_218ba8e1456a7869.
		Lock()
	defer __obf_7f538276625ed09b.__obf_218ba8e1456a7869.Unlock()
	__obf_5fa8637e08c90073 := __obf_7f538276625ed09b.__obf_04d7f434df35909c()
	if __obf_5fa8637e08c90073 == __obf_7f538276625ed09b.__obf_0c18bdba049ba1e6 {
		__obf_7f538276625ed09b.__obf_32d285c02de19ff8 = (__obf_7f538276625ed09b.__obf_32d285c02de19ff8 + 1) & CSequenceMask
		if __obf_7f538276625ed09b.__obf_32d285c02de19ff8 == 0 {
			__obf_5fa8637e08c90073 = __obf_7f538276625ed09b.__obf_2c1ebd77380eb657(__obf_5fa8637e08c90073)
		}
	} else {
		__obf_7f538276625ed09b.__obf_32d285c02de19ff8 = 0
	}

	if __obf_5fa8637e08c90073 < __obf_7f538276625ed09b.__obf_0c18bdba049ba1e6 {
		__obf_57cef2f010b4ce9a = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_57cef2f010b4ce9a
	}
	__obf_7f538276625ed09b.__obf_0c18bdba049ba1e6 = __obf_5fa8637e08c90073
	__obf_5fa8637e08c90073 = (__obf_5fa8637e08c90073-CEpoch)<<CTimeStampShift | __obf_7f538276625ed09b.__obf_99a2aa55cbd819ac<<CWorkerIdShift | __obf_7f538276625ed09b.__obf_32d285c02de19ff8
	return fmt.Sprintf("%X", __obf_5fa8637e08c90073), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_413d3a76ff8ab79e int64) (__obf_b5849d1d191d006d time.Time, __obf_5fa8637e08c90073 int64, __obf_99a2aa55cbd819ac int64, __obf_692a1478e50c31af int64) {
	__obf_692a1478e50c31af = __obf_413d3a76ff8ab79e & CSequenceMask
	__obf_99a2aa55cbd819ac = (__obf_413d3a76ff8ab79e >> CWorkerIdShift) & CMaxWorker
	__obf_5fa8637e08c90073 = (__obf_413d3a76ff8ab79e >> CTimeStampShift) + CEpoch
	__obf_b5849d1d191d006d = time.Unix(__obf_5fa8637e08c90073/1000, (__obf_5fa8637e08c90073%1000)*1000000)
	return
}
