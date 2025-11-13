package __obf_417508f5ae215d0a

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// IdWorker Struct
type IDWorker struct {
	__obf_99c6557d706f015c int64
	__obf_79a4b2c02807dc4e int64
	__obf_cc3a3b06b6be52c3 int64
	__obf_ed1448eab7d488b1 int64
	__obf_7480d635725fedc4 *sync.Mutex
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
func NewIdWorker(__obf_077ba605a7b371f8 int64) (__obf_6c06962f14ee1cef *IDWorker, __obf_b99f62eaceb52d40 error) {
	__obf_6c06962f14ee1cef = new(IDWorker)

	__obf_6c06962f14ee1cef.__obf_ed1448eab7d488b1 = __obf_85fa4bea1bc6e4d1()

	if __obf_077ba605a7b371f8 > __obf_6c06962f14ee1cef.__obf_ed1448eab7d488b1 || __obf_077ba605a7b371f8 < 0 {
		return nil, errors.New("worker not fit")
	}
	__obf_6c06962f14ee1cef.__obf_99c6557d706f015c = __obf_077ba605a7b371f8
	__obf_6c06962f14ee1cef.__obf_79a4b2c02807dc4e = -1
	__obf_6c06962f14ee1cef.__obf_cc3a3b06b6be52c3 = 0
	__obf_6c06962f14ee1cef.__obf_7480d635725fedc4 = new(sync.Mutex)
	return __obf_6c06962f14ee1cef, nil
}

func __obf_85fa4bea1bc6e4d1() int64 {
	return -1 ^ -1<<CWorkerIdBits
}

func __obf_1423ae10358e09c4() int64 {
	return -1 ^ -1<<CSenquenceBits
}

// return in ms
func (__obf_6c06962f14ee1cef *IDWorker) __obf_27c44d3de9117743() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (__obf_6c06962f14ee1cef *IDWorker) __obf_e54064cdeb97dd8e(__obf_c8d252a40f6eff2a int64) int64 {
	__obf_9b6047abad04e0d9 := time.Now().UnixNano() / 1000 / 1000
	for {
		if __obf_9b6047abad04e0d9 <= __obf_c8d252a40f6eff2a {
			__obf_9b6047abad04e0d9 = __obf_6c06962f14ee1cef.__obf_27c44d3de9117743()
		} else {
			break
		}
	}
	return __obf_9b6047abad04e0d9
}

// NewId Func: Generate next id
func (__obf_6c06962f14ee1cef *IDWorker) NextID() (__obf_48218c2a1a645d91 string, __obf_b99f62eaceb52d40 error) {
	__obf_6c06962f14ee1cef.__obf_7480d635725fedc4.Lock()
	defer __obf_6c06962f14ee1cef.__obf_7480d635725fedc4.Unlock()
	__obf_9b6047abad04e0d9 := __obf_6c06962f14ee1cef.__obf_27c44d3de9117743()
	if __obf_9b6047abad04e0d9 == __obf_6c06962f14ee1cef.__obf_79a4b2c02807dc4e {
		__obf_6c06962f14ee1cef.__obf_cc3a3b06b6be52c3 = (__obf_6c06962f14ee1cef.__obf_cc3a3b06b6be52c3 + 1) & CSequenceMask
		if __obf_6c06962f14ee1cef.__obf_cc3a3b06b6be52c3 == 0 {
			__obf_9b6047abad04e0d9 = __obf_6c06962f14ee1cef.__obf_e54064cdeb97dd8e(__obf_9b6047abad04e0d9)
		}
	} else {
		__obf_6c06962f14ee1cef.__obf_cc3a3b06b6be52c3 = 0
	}

	if __obf_9b6047abad04e0d9 < __obf_6c06962f14ee1cef.__obf_79a4b2c02807dc4e {
		__obf_b99f62eaceb52d40 = errors.New("clock moved backwards, refuse gen id")
		return "", __obf_b99f62eaceb52d40
	}
	__obf_6c06962f14ee1cef.__obf_79a4b2c02807dc4e = __obf_9b6047abad04e0d9
	__obf_9b6047abad04e0d9 = (__obf_9b6047abad04e0d9-CEpoch)<<CTimeStampShift | __obf_6c06962f14ee1cef.__obf_99c6557d706f015c<<CWorkerIdShift | __obf_6c06962f14ee1cef.__obf_cc3a3b06b6be52c3
	return fmt.Sprintf("%X", __obf_9b6047abad04e0d9), nil
}

// ParseId Func: reverse uid to timestamp, workid, seq
func ParseId(__obf_48218c2a1a645d91 int64) (__obf_59b4996dba32a5cd time.Time, __obf_9b6047abad04e0d9 int64, __obf_99c6557d706f015c int64, __obf_74c84b6c5c790cf4 int64) {
	__obf_74c84b6c5c790cf4 = __obf_48218c2a1a645d91 & CSequenceMask
	__obf_99c6557d706f015c = (__obf_48218c2a1a645d91 >> CWorkerIdShift) & CMaxWorker
	__obf_9b6047abad04e0d9 = (__obf_48218c2a1a645d91 >> CTimeStampShift) + CEpoch
	__obf_59b4996dba32a5cd = time.Unix(__obf_9b6047abad04e0d9/1000, (__obf_9b6047abad04e0d9%1000)*1000000)
	return
}
