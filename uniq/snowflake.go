package __obf_0f0e0d1ff72f3ff0

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_9deced313e158bdd int64
	__obf_9f6945a1feaa0cfa int64
	__obf_60bc88779bdc777d int64
	__obf_346220989aa73748 int64
	__obf_fd5e7e85c5fc0569 *sync.Mutex
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
func NewIdWorker(__obf_baefde1bf1519d77 int64) (__obf_7801b238940b2c47 *IDWorker, __obf_f24ec7ae16791add error) {
	__obf_7801b238940b2c47 = new(IDWorker)

	__obf_7801b238940b2c47.__obf_346220989aa73748 = __obf_c1085025470bee0b()

	if __obf_baefde1bf1519d77 > __obf_7801b238940b2c47.__obf_346220989aa73748 || __obf_baefde1bf1519d77 < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_7801b238940b2c47.__obf_9deced313e158bdd = __obf_baefde1bf1519d77
	__obf_7801b238940b2c47.__obf_9f6945a1feaa0cfa = -1
	__obf_7801b238940b2c47.__obf_60bc88779bdc777d = 0
	__obf_7801b238940b2c47.__obf_fd5e7e85c5fc0569 = new(sync.Mutex)
	return __obf_7801b238940b2c47, nil
}

func __obf_c1085025470bee0b() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_7995c020c62f17c3() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_7801b238940b2c47 *IDWorker) __obf_16f8dd9532a6fb76() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_7801b238940b2c47 *IDWorker) __obf_98c65466b54133ff(__obf_82dc7e5fcb802101 int64) int64 {
	__obf_d33d13bb08a6244e := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_d33d13bb08a6244e <= __obf_82dc7e5fcb802101 {
			__obf_d33d13bb08a6244e = __obf_7801b238940b2c47.__obf_16f8dd9532a6fb76()
		} else {
			break
		}
	}
	return __obf_d33d13bb08a6244e
}

// NewId Func: Generate next id
func (__obf_7801b238940b2c47 *IDWorker) NextID() (__obf_a7593a31c69c5d40 string, __obf_f24ec7ae16791add error) {
	__obf_7801b238940b2c47.__obf_fd5e7e85c5fc0569.Lock()
	defer __obf_7801b238940b2c47.__obf_fd5e7e85c5fc0569.Unlock()
	__obf_d33d13bb08a6244e := __obf_7801b238940b2c47.__obf_16f8dd9532a6fb76()
	if __obf_d33d13bb08a6244e == __obf_7801b238940b2c47.__obf_9f6945a1feaa0cfa {
		__obf_7801b238940b2c47.__obf_60bc88779bdc777d = (__obf_7801b238940b2c47.__obf_60bc88779bdc777d + 1) & CSequenceMask
		if __obf_7801b238940b2c47.__obf_60bc88779bdc777d == 0 {
			__obf_d33d13bb08a6244e = __obf_7801b238940b2c47.__obf_98c65466b54133ff(__obf_d33d13bb08a6244e)
		}
	} else {
		__obf_7801b238940b2c47.__obf_60bc88779bdc777d = 0
	}

	if __obf_d33d13bb08a6244e < __obf_7801b238940b2c47.__obf_9f6945a1feaa0cfa {
		__obf_f24ec7ae16791add = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_f24ec7ae16791add
	}
	__obf_7801b238940b2c47.__obf_9f6945a1feaa0cfa = __obf_d33d13bb08a6244e
	__obf_d33d13bb08a6244e = (__obf_d33d13bb08a6244e-CEpoch)<<CTimeStampShift | __obf_7801b238940b2c47.__obf_9deced313e158bdd<<CWorkerIdShift | __obf_7801b238940b2c47.__obf_60bc88779bdc777d
	return fmt.Sprintf("%X", __obf_d33d13bb08a6244e), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_a7593a31c69c5d40 int64) (__obf_3c7f4fcc21cb9f0d time.Time, __obf_d33d13bb08a6244e int64, __obf_9deced313e158bdd int64, __obf_5e6d15cf509517a5 int64) {
	__obf_5e6d15cf509517a5 = __obf_a7593a31c69c5d40 & CSequenceMask
	__obf_9deced313e158bdd = (__obf_a7593a31c69c5d40 >> CWorkerIdShift) & CMaxWorker
	__obf_d33d13bb08a6244e = (__obf_a7593a31c69c5d40 >> CTimeStampShift) + CEpoch
	__obf_3c7f4fcc21cb9f0d = time.Unix(__obf_d33d13bb08a6244e/1000, (__obf_d33d13bb08a6244e%1000)*1000000)
	return
}
