package __obf_de86cdc8ae98b45b

import (
	"reflect"
	"sync"
)

var __obf_aae1f88805984eef struct {
	__obf_2a26e7104b4c4373 map[reflect.Type]chan reflect.Value
	sync.RWMutex
}

func __obf_cd14bf657d2f8fbe(__obf_b5900a65e8e7d9b5 reflect.Type) reflect.Value {
	__obf_aae1f88805984eef.
		RLock()
	__obf_bb5176e4e1e37d52 := __obf_aae1f88805984eef.__obf_2a26e7104b4c4373[__obf_b5900a65e8e7d9b5]
	__obf_aae1f88805984eef.
		RUnlock()
	if __obf_bb5176e4e1e37d52 != nil {
		return <-__obf_bb5176e4e1e37d52
	}
	__obf_aae1f88805984eef.
		Lock()
	defer __obf_aae1f88805984eef.Unlock()
	if __obf_bb5176e4e1e37d52 = __obf_aae1f88805984eef.__obf_2a26e7104b4c4373[__obf_b5900a65e8e7d9b5]; __obf_bb5176e4e1e37d52 != nil {
		return <-__obf_bb5176e4e1e37d52
	}
	__obf_bb5176e4e1e37d52 = make(chan reflect.Value, 256)
	go func() {
		for {
			__obf_bb5176e4e1e37d52 <- reflect.New(__obf_b5900a65e8e7d9b5)
		}
	}()
	if __obf_aae1f88805984eef.__obf_2a26e7104b4c4373 == nil {
		__obf_aae1f88805984eef.__obf_2a26e7104b4c4373 = make(map[reflect.Type]chan reflect.Value, 8)
	}
	__obf_aae1f88805984eef.__obf_2a26e7104b4c4373[__obf_b5900a65e8e7d9b5] = __obf_bb5176e4e1e37d52
	return <-__obf_bb5176e4e1e37d52
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_598c5a4b42040fb1(__obf_b5900a65e8e7d9b5 reflect.Type) reflect.Value {
	if __obf_dcaad1165bb07af9.__obf_a15a642720e39c3e&__obf_079ef68f50d79bcd == 0 {
		return reflect.New(__obf_b5900a65e8e7d9b5)
	}

	return __obf_cd14bf657d2f8fbe(__obf_b5900a65e8e7d9b5)
}
