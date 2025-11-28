package __obf_8281010a3a6d2ab0

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_d5b5da4c935838ef *memcache.Client
}

func (__obf_0bb6a155e4080f19 *Memcache) Has(__obf_805b9182f4a01a43 string) bool {
	_, __obf_fd03130c6793cb3b := __obf_0bb6a155e4080f19.__obf_d5b5da4c935838ef.Get(__obf_805b9182f4a01a43)
	return __obf_fd03130c6793cb3b == nil
}

// Delete implements Cache.
func (__obf_0bb6a155e4080f19 *Memcache) Delete(__obf_c2739d84a72f2e49 string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_686e4b10aabc1c6f []string) Cache {
	return &Memcache{memcache.New(__obf_686e4b10aabc1c6f...)}
}

// Get get value from memcache.
func (__obf_0bb6a155e4080f19 *Memcache) Get(__obf_805b9182f4a01a43 string) (__obf_7abe52d80ce16a1e []byte, __obf_fd03130c6793cb3b error) {
	if __obf_c3889543117ef1cd, __obf_fd03130c6793cb3b := __obf_0bb6a155e4080f19.__obf_d5b5da4c935838ef.Get(__obf_805b9182f4a01a43); __obf_fd03130c6793cb3b == nil {

		return __obf_c3889543117ef1cd.Value, nil
	} else {
		return nil, __obf_fd03130c6793cb3b
	}
}

func (__obf_0bb6a155e4080f19 *Memcache) Set(__obf_805b9182f4a01a43 string, __obf_a31843a6764aecf9 any, __obf_e2e393700839b297 time.Duration) error {
	__obf_c3889543117ef1cd := memcache.Item{Key: __obf_805b9182f4a01a43, Expiration: int32(__obf_e2e393700839b297 / time.Second)}
	var __obf_fd03130c6793cb3b error
	__obf_c3889543117ef1cd.Value, __obf_fd03130c6793cb3b = util.AnyToBytes(__obf_a31843a6764aecf9)
	if __obf_fd03130c6793cb3b != nil {
		return __obf_fd03130c6793cb3b
	}
	return __obf_0bb6a155e4080f19.__obf_d5b5da4c935838ef.Set(&__obf_c3889543117ef1cd)
}

func (__obf_0bb6a155e4080f19 *Memcache) Add(__obf_805b9182f4a01a43 string, __obf_a31843a6764aecf9 any) error {
	__obf_c3889543117ef1cd := memcache.Item{Key: __obf_805b9182f4a01a43}

	var __obf_fd03130c6793cb3b error
	__obf_c3889543117ef1cd.Value, __obf_fd03130c6793cb3b = util.AnyToBytes(__obf_a31843a6764aecf9)
	if __obf_fd03130c6793cb3b != nil {
		return __obf_fd03130c6793cb3b
	}

	return __obf_0bb6a155e4080f19.__obf_d5b5da4c935838ef.Set(&__obf_c3889543117ef1cd)
}

// ClearAll clears all cache in memcache.
func (__obf_0bb6a155e4080f19 *Memcache) Clear() error {
	return __obf_0bb6a155e4080f19.__obf_d5b5da4c935838ef.FlushAll()
}

// Incr increases counter.
func (__obf_0bb6a155e4080f19 *Memcache) Incr(__obf_805b9182f4a01a43 string) error {
	_, __obf_fd03130c6793cb3b := __obf_0bb6a155e4080f19.__obf_d5b5da4c935838ef.Increment(__obf_805b9182f4a01a43, 1)
	return __obf_fd03130c6793cb3b
}

// Decr decreases counter.
func (__obf_0bb6a155e4080f19 *Memcache) Decr(__obf_805b9182f4a01a43 string) error {
	_, __obf_fd03130c6793cb3b := __obf_0bb6a155e4080f19.__obf_d5b5da4c935838ef.Decrement(__obf_805b9182f4a01a43, 1)
	return __obf_fd03130c6793cb3b
}

func (__obf_0bb6a155e4080f19 *Memcache) Renew(__obf_805b9182f4a01a43 string, __obf_e2e393700839b297 time.Duration) error {
	if __obf_c3889543117ef1cd, __obf_fd03130c6793cb3b := __obf_0bb6a155e4080f19.__obf_d5b5da4c935838ef.Get(__obf_805b9182f4a01a43); __obf_fd03130c6793cb3b == nil {
		return __obf_0bb6a155e4080f19.Set(__obf_805b9182f4a01a43, __obf_c3889543117ef1cd, __obf_e2e393700839b297)
	} else {
		return __obf_fd03130c6793cb3b
	}
}

func (__obf_0bb6a155e4080f19 *Memcache) IsExist(__obf_805b9182f4a01a43 string) bool {
	_, __obf_fd03130c6793cb3b := __obf_0bb6a155e4080f19.__obf_d5b5da4c935838ef.Get(__obf_805b9182f4a01a43)
	return __obf_fd03130c6793cb3b == nil
}

func (__obf_0bb6a155e4080f19 *Memcache) Remove(__obf_805b9182f4a01a43 string) error {
	return __obf_0bb6a155e4080f19.__obf_d5b5da4c935838ef.Delete(__obf_805b9182f4a01a43)
}
