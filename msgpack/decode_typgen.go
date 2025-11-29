package __obf_a4aad98aaf3178f4

import (
	"reflect"
	"sync"
)

var __obf_e504b6890bfb0de7 struct {
	__obf_66cc4b26e3c9cdbb map[reflect.Type]chan reflect.Value
	sync.RWMutex
}

func __obf_5f0c2ccecff4d570(__obf_8698f5d73996ef26 reflect.Type) reflect.Value {
	__obf_e504b6890bfb0de7.
		RLock()
	__obf_da1f18cf380debe4 := __obf_e504b6890bfb0de7.__obf_66cc4b26e3c9cdbb[__obf_8698f5d73996ef26]
	__obf_e504b6890bfb0de7.
		RUnlock()
	if __obf_da1f18cf380debe4 != nil {
		return <-__obf_da1f18cf380debe4
	}
	__obf_e504b6890bfb0de7.
		Lock()
	defer __obf_e504b6890bfb0de7.Unlock()
	if __obf_da1f18cf380debe4 = __obf_e504b6890bfb0de7.__obf_66cc4b26e3c9cdbb[__obf_8698f5d73996ef26]; __obf_da1f18cf380debe4 != nil {
		return <-__obf_da1f18cf380debe4
	}
	__obf_da1f18cf380debe4 = make(chan reflect.Value, 256)
	go func() {
		for {
			__obf_da1f18cf380debe4 <- reflect.New(__obf_8698f5d73996ef26)
		}
	}()
	if __obf_e504b6890bfb0de7.__obf_66cc4b26e3c9cdbb == nil {
		__obf_e504b6890bfb0de7.__obf_66cc4b26e3c9cdbb = make(map[reflect.Type]chan reflect.Value, 8)
	}
	__obf_e504b6890bfb0de7.__obf_66cc4b26e3c9cdbb[__obf_8698f5d73996ef26] = __obf_da1f18cf380debe4
	return <-__obf_da1f18cf380debe4
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_7a63f73d6b3d6972(__obf_8698f5d73996ef26 reflect.Type) reflect.Value {
	if __obf_613397fefdec0ed0.__obf_a8400e70a712e9fa&__obf_97ba1004800fc0a8 == 0 {
		return reflect.New(__obf_8698f5d73996ef26)
	}

	return __obf_5f0c2ccecff4d570(__obf_8698f5d73996ef26)
}
