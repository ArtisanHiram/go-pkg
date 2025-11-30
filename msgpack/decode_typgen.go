package __obf_3e0c303610a19bd4

import (
	"reflect"
	"sync"
)

var __obf_8685a2e4ec8e36a2 struct {
	__obf_b416b6a4555be5c8 map[reflect.Type]chan reflect.Value
	sync.RWMutex
}

func __obf_bcf79cb8d15fbf1f(__obf_cb646455b7fd3ba7 reflect.Type) reflect.Value {
	__obf_8685a2e4ec8e36a2.
		RLock()
	__obf_57f9be11ae42c8bc := __obf_8685a2e4ec8e36a2.__obf_b416b6a4555be5c8[__obf_cb646455b7fd3ba7]
	__obf_8685a2e4ec8e36a2.
		RUnlock()
	if __obf_57f9be11ae42c8bc != nil {
		return <-__obf_57f9be11ae42c8bc
	}
	__obf_8685a2e4ec8e36a2.
		Lock()
	defer __obf_8685a2e4ec8e36a2.Unlock()
	if __obf_57f9be11ae42c8bc = __obf_8685a2e4ec8e36a2.__obf_b416b6a4555be5c8[__obf_cb646455b7fd3ba7]; __obf_57f9be11ae42c8bc != nil {
		return <-__obf_57f9be11ae42c8bc
	}
	__obf_57f9be11ae42c8bc = make(chan reflect.Value, 256)
	go func() {
		for {
			__obf_57f9be11ae42c8bc <- reflect.New(__obf_cb646455b7fd3ba7)
		}
	}()
	if __obf_8685a2e4ec8e36a2.__obf_b416b6a4555be5c8 == nil {
		__obf_8685a2e4ec8e36a2.__obf_b416b6a4555be5c8 = make(map[reflect.Type]chan reflect.Value, 8)
	}
	__obf_8685a2e4ec8e36a2.__obf_b416b6a4555be5c8[__obf_cb646455b7fd3ba7] = __obf_57f9be11ae42c8bc
	return <-__obf_57f9be11ae42c8bc
}

func (__obf_dc35117108ba8439 *Decoder) __obf_94f521ecbd0b4afa(__obf_cb646455b7fd3ba7 reflect.Type) reflect.Value {
	if __obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb&__obf_875c65698843f8e6 == 0 {
		return reflect.New(__obf_cb646455b7fd3ba7)
	}

	return __obf_bcf79cb8d15fbf1f(__obf_cb646455b7fd3ba7)
}
