package __obf_a51a64e21f311927

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_fd14b2689ab4f0d7 int64
	__obf_09e0c19d9eab9bb3 int64
	__obf_97426f7ecb5d7ca6 int64
	__obf_41a95c1e40e67035 int64
	__obf_6b1f7f3376bfd1d2 *sync.Mutex
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
func NewIdWorker(__obf_741b6fa555c343ee int64) (__obf_1e118d648d470158 *IDWorker, __obf_886df8acc7f50b76 error) {
	__obf_1e118d648d470158 = new(IDWorker)
	__obf_1e118d648d470158.__obf_41a95c1e40e67035 = __obf_589dd7a566364483()

	if __obf_741b6fa555c343ee > __obf_1e118d648d470158.__obf_41a95c1e40e67035 || __obf_741b6fa555c343ee < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_1e118d648d470158.__obf_fd14b2689ab4f0d7 = __obf_741b6fa555c343ee
	__obf_1e118d648d470158.__obf_09e0c19d9eab9bb3 = -1
	__obf_1e118d648d470158.__obf_97426f7ecb5d7ca6 = 0
	__obf_1e118d648d470158.__obf_6b1f7f3376bfd1d2 = new(sync.Mutex)
	return __obf_1e118d648d470158, nil
}

func __obf_589dd7a566364483() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_3b0d4d538ada96e4() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_1e118d648d470158 *IDWorker) __obf_5ed3be5021b105fa() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_1e118d648d470158 *IDWorker) __obf_bc8ca894e29cd584(__obf_357a600734311b17 int64) int64 {
	__obf_26c99594a4c81bf5 := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_26c99594a4c81bf5 <= __obf_357a600734311b17 {
			__obf_26c99594a4c81bf5 = __obf_1e118d648d470158.__obf_5ed3be5021b105fa()
		} else {
			break
		}
	}
	return __obf_26c99594a4c81bf5
}

// NewId Func: Generate next id
func (__obf_1e118d648d470158 *IDWorker) NextID() (__obf_4c112818117e5211 string, __obf_886df8acc7f50b76 error) {
	__obf_1e118d648d470158.__obf_6b1f7f3376bfd1d2.
		Lock()
	defer __obf_1e118d648d470158.__obf_6b1f7f3376bfd1d2.Unlock()
	__obf_26c99594a4c81bf5 := __obf_1e118d648d470158.__obf_5ed3be5021b105fa()
	if __obf_26c99594a4c81bf5 == __obf_1e118d648d470158.__obf_09e0c19d9eab9bb3 {
		__obf_1e118d648d470158.__obf_97426f7ecb5d7ca6 = (__obf_1e118d648d470158.__obf_97426f7ecb5d7ca6 + 1) & CSequenceMask
		if __obf_1e118d648d470158.__obf_97426f7ecb5d7ca6 == 0 {
			__obf_26c99594a4c81bf5 = __obf_1e118d648d470158.__obf_bc8ca894e29cd584(__obf_26c99594a4c81bf5)
		}
	} else {
		__obf_1e118d648d470158.__obf_97426f7ecb5d7ca6 = 0
	}

	if __obf_26c99594a4c81bf5 < __obf_1e118d648d470158.__obf_09e0c19d9eab9bb3 {
		__obf_886df8acc7f50b76 = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_886df8acc7f50b76
	}
	__obf_1e118d648d470158.__obf_09e0c19d9eab9bb3 = __obf_26c99594a4c81bf5
	__obf_26c99594a4c81bf5 = (__obf_26c99594a4c81bf5-CEpoch)<<CTimeStampShift | __obf_1e118d648d470158.__obf_fd14b2689ab4f0d7<<CWorkerIdShift | __obf_1e118d648d470158.__obf_97426f7ecb5d7ca6
	return fmt.Sprintf("%X", __obf_26c99594a4c81bf5), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_4c112818117e5211 int64) (__obf_c5969bc7b424c7a8 time.Time, __obf_26c99594a4c81bf5 int64, __obf_fd14b2689ab4f0d7 int64, __obf_6176ddf512a087e0 int64) {
	__obf_6176ddf512a087e0 = __obf_4c112818117e5211 & CSequenceMask
	__obf_fd14b2689ab4f0d7 = (__obf_4c112818117e5211 >> CWorkerIdShift) & CMaxWorker
	__obf_26c99594a4c81bf5 = (__obf_4c112818117e5211 >> CTimeStampShift) + CEpoch
	__obf_c5969bc7b424c7a8 = time.Unix(__obf_26c99594a4c81bf5/1000, (__obf_26c99594a4c81bf5%1000)*1000000)
	return
}
