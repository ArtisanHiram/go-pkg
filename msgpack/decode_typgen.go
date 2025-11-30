package __obf_fd770cb9675903b0

import (
	"reflect"
	"sync"
)

var __obf_39b6b1495e073fde struct {
	__obf_777489ec8e6b2044 map[reflect.Type]chan reflect.Value
	sync.RWMutex
}

func __obf_ed787a5d7a207071(__obf_0eb9b9619deca594 reflect.Type) reflect.Value {
	__obf_39b6b1495e073fde.
		RLock()
	__obf_69ed3f4ceaf4de65 := __obf_39b6b1495e073fde.__obf_777489ec8e6b2044[__obf_0eb9b9619deca594]
	__obf_39b6b1495e073fde.
		RUnlock()
	if __obf_69ed3f4ceaf4de65 != nil {
		return <-__obf_69ed3f4ceaf4de65
	}
	__obf_39b6b1495e073fde.
		Lock()
	defer __obf_39b6b1495e073fde.Unlock()
	if __obf_69ed3f4ceaf4de65 = __obf_39b6b1495e073fde.__obf_777489ec8e6b2044[__obf_0eb9b9619deca594]; __obf_69ed3f4ceaf4de65 != nil {
		return <-__obf_69ed3f4ceaf4de65
	}
	__obf_69ed3f4ceaf4de65 = make(chan reflect.Value, 256)
	go func() {
		for {
			__obf_69ed3f4ceaf4de65 <- reflect.New(__obf_0eb9b9619deca594)
		}
	}()
	if __obf_39b6b1495e073fde.__obf_777489ec8e6b2044 == nil {
		__obf_39b6b1495e073fde.__obf_777489ec8e6b2044 = make(map[reflect.Type]chan reflect.Value, 8)
	}
	__obf_39b6b1495e073fde.__obf_777489ec8e6b2044[__obf_0eb9b9619deca594] = __obf_69ed3f4ceaf4de65
	return <-__obf_69ed3f4ceaf4de65
}

func (__obf_5d389b660eefb08c *Decoder) __obf_9eb3ea28cca5995b(__obf_0eb9b9619deca594 reflect.Type) reflect.Value {
	if __obf_5d389b660eefb08c.__obf_7185cb62de7af792&__obf_759b68f674456e1c == 0 {
		return reflect.New(__obf_0eb9b9619deca594)
	}

	return __obf_ed787a5d7a207071(__obf_0eb9b9619deca594)
}
