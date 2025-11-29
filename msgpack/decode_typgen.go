package __obf_3edaa49e853afa16

import (
	"reflect"
	"sync"
)

var __obf_9fe6bb7e98c0c037 struct {
	__obf_c26f5adfb3152475 map[reflect.Type]chan reflect.Value
	sync.RWMutex
}

func __obf_c3a387a7995b1678(__obf_862c8bdd067ceabe reflect.Type) reflect.Value {
	__obf_9fe6bb7e98c0c037.
		RLock()
	__obf_46eefac4e0b76040 := __obf_9fe6bb7e98c0c037.__obf_c26f5adfb3152475[__obf_862c8bdd067ceabe]
	__obf_9fe6bb7e98c0c037.
		RUnlock()
	if __obf_46eefac4e0b76040 != nil {
		return <-__obf_46eefac4e0b76040
	}
	__obf_9fe6bb7e98c0c037.
		Lock()
	defer __obf_9fe6bb7e98c0c037.Unlock()
	if __obf_46eefac4e0b76040 = __obf_9fe6bb7e98c0c037.__obf_c26f5adfb3152475[__obf_862c8bdd067ceabe]; __obf_46eefac4e0b76040 != nil {
		return <-__obf_46eefac4e0b76040
	}
	__obf_46eefac4e0b76040 = make(chan reflect.Value, 256)
	go func() {
		for {
			__obf_46eefac4e0b76040 <- reflect.New(__obf_862c8bdd067ceabe)
		}
	}()
	if __obf_9fe6bb7e98c0c037.__obf_c26f5adfb3152475 == nil {
		__obf_9fe6bb7e98c0c037.__obf_c26f5adfb3152475 = make(map[reflect.Type]chan reflect.Value, 8)
	}
	__obf_9fe6bb7e98c0c037.__obf_c26f5adfb3152475[__obf_862c8bdd067ceabe] = __obf_46eefac4e0b76040
	return <-__obf_46eefac4e0b76040
}

func (__obf_015afbee33862a01 *Decoder) __obf_4c87e5dcce642cc7(__obf_862c8bdd067ceabe reflect.Type) reflect.Value {
	if __obf_015afbee33862a01.__obf_6aa0efd537415192&__obf_5157ad164f309c0c == 0 {
		return reflect.New(__obf_862c8bdd067ceabe)
	}

	return __obf_c3a387a7995b1678(__obf_862c8bdd067ceabe)
}
