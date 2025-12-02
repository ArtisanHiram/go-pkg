package __obf_a935eb7f548271a4

import (
	"reflect"
	"sync"
)

var __obf_cd207c0ce60bce24 struct {
	__obf_26d12ef0df36a324 map[reflect.Type]chan reflect.Value
	sync.RWMutex
}

func __obf_14360e03a3b5da30(__obf_14c99752d39cbe18 reflect.Type) reflect.Value {
	__obf_cd207c0ce60bce24.
		RLock()
	__obf_5016540e60b4e742 := __obf_cd207c0ce60bce24.__obf_26d12ef0df36a324[__obf_14c99752d39cbe18]
	__obf_cd207c0ce60bce24.
		RUnlock()
	if __obf_5016540e60b4e742 != nil {
		return <-__obf_5016540e60b4e742
	}
	__obf_cd207c0ce60bce24.
		Lock()
	defer __obf_cd207c0ce60bce24.Unlock()
	if __obf_5016540e60b4e742 = __obf_cd207c0ce60bce24.__obf_26d12ef0df36a324[__obf_14c99752d39cbe18]; __obf_5016540e60b4e742 != nil {
		return <-__obf_5016540e60b4e742
	}
	__obf_5016540e60b4e742 = make(chan reflect.Value, 256)
	go func() {
		for {
			__obf_5016540e60b4e742 <- reflect.New(__obf_14c99752d39cbe18)
		}
	}()
	if __obf_cd207c0ce60bce24.__obf_26d12ef0df36a324 == nil {
		__obf_cd207c0ce60bce24.__obf_26d12ef0df36a324 = make(map[reflect.Type]chan reflect.Value, 8)
	}
	__obf_cd207c0ce60bce24.__obf_26d12ef0df36a324[__obf_14c99752d39cbe18] = __obf_5016540e60b4e742
	return <-__obf_5016540e60b4e742
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_1fc1190bc702b3d7(__obf_14c99752d39cbe18 reflect.Type) reflect.Value {
	if __obf_a21885da2425f2b2.__obf_d617f3769ce16c47&__obf_4cf2a1ef1c20b040 == 0 {
		return reflect.New(__obf_14c99752d39cbe18)
	}

	return __obf_14360e03a3b5da30(__obf_14c99752d39cbe18)
}
