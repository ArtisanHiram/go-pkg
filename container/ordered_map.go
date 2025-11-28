package __obf_4f13ac5aa043b5ce

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_aa52d12db26ef244 map[string]*Element[V]
	__obf_c451bc9b9d85eff5 __obf_25770552d978c33e[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{
		__obf_aa52d12db26ef244: make(map[string]*Element[V]),
	}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_2403dfddde3862ad *OrderedMap[V]) Get(__obf_797707a3dac0ebb7 string) (__obf_77dcf3eed779629a V, __obf_c7b01e58066bbb7a bool) {
	__obf_1e75fd93d2a55aa7, __obf_c7b01e58066bbb7a := __obf_2403dfddde3862ad.__obf_aa52d12db26ef244[__obf_797707a3dac0ebb7]
	if __obf_c7b01e58066bbb7a {
		__obf_77dcf3eed779629a = __obf_1e75fd93d2a55aa7.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_2403dfddde3862ad *OrderedMap[V]) Set(__obf_797707a3dac0ebb7 string, __obf_77dcf3eed779629a V) bool {
	_, __obf_b6a38de6d8c295dc := __obf_2403dfddde3862ad.__obf_aa52d12db26ef244[__obf_797707a3dac0ebb7]
	if __obf_b6a38de6d8c295dc {
		__obf_2403dfddde3862ad.__obf_aa52d12db26ef244[__obf_797707a3dac0ebb7].Value = __obf_77dcf3eed779629a
		return false
	}

	__obf_8338fba3de46db62 := __obf_2403dfddde3862ad.__obf_c451bc9b9d85eff5.PushBack(__obf_797707a3dac0ebb7, __obf_77dcf3eed779629a)
	__obf_2403dfddde3862ad.__obf_aa52d12db26ef244[__obf_797707a3dac0ebb7] = __obf_8338fba3de46db62
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_2403dfddde3862ad *OrderedMap[V]) GetOrDefault(__obf_797707a3dac0ebb7 string, __obf_98302c3c43e74590 V) V {
	if __obf_77dcf3eed779629a, __obf_c7b01e58066bbb7a := __obf_2403dfddde3862ad.__obf_aa52d12db26ef244[__obf_797707a3dac0ebb7]; __obf_c7b01e58066bbb7a {
		return __obf_77dcf3eed779629a.Value
	}

	return __obf_98302c3c43e74590
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_2403dfddde3862ad *OrderedMap[V]) GetElement(__obf_797707a3dac0ebb7 string) *Element[V] {
	__obf_8338fba3de46db62, __obf_c7b01e58066bbb7a := __obf_2403dfddde3862ad.__obf_aa52d12db26ef244[__obf_797707a3dac0ebb7]
	if __obf_c7b01e58066bbb7a {
		return __obf_8338fba3de46db62
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_2403dfddde3862ad *OrderedMap[V]) Len() int {
	return len(__obf_2403dfddde3862ad.__obf_aa52d12db26ef244)
}

// Keys returns all of the keys in the order they were inserted. If a key was
// replaced it will retain the same position. To ensure most recently set keys
// are always at the end you must always Delete before Set.
func (__obf_2403dfddde3862ad *OrderedMap[V]) Keys() (__obf_e7eafad979141f9b []string) {
	__obf_e7eafad979141f9b = make([]string, 0, __obf_2403dfddde3862ad.Len())
	for __obf_23f1e93b8e98e9d5 := __obf_2403dfddde3862ad.Front(); __obf_23f1e93b8e98e9d5 != nil; __obf_23f1e93b8e98e9d5 = __obf_23f1e93b8e98e9d5.Next() {
		__obf_e7eafad979141f9b = append(__obf_e7eafad979141f9b, __obf_23f1e93b8e98e9d5.Key)
	}
	return __obf_e7eafad979141f9b
}

// Values returns all of the values in the order they were inserted.
func (__obf_2403dfddde3862ad *OrderedMap[V]) Values() (__obf_ad12e06d8e8771c7 []V) {
	__obf_ad12e06d8e8771c7 = make([]V, 0, __obf_2403dfddde3862ad.Len())
	for __obf_23f1e93b8e98e9d5 := __obf_2403dfddde3862ad.Front(); __obf_23f1e93b8e98e9d5 != nil; __obf_23f1e93b8e98e9d5 = __obf_23f1e93b8e98e9d5.Next() {
		__obf_ad12e06d8e8771c7 = append(__obf_ad12e06d8e8771c7, __obf_23f1e93b8e98e9d5.Value)
	}
	return __obf_ad12e06d8e8771c7
}

// MarshalJSON implements type json.Marshaler interface, so can be called in json.Marshal(om)
// func (m *OrderedMap[V]) MarshalJSON() (res []byte, err error) {
// 	res = append(res, '{')
// 	front, back := m.ll.Front(), m.ll.Back()
// 	for e := front; e != nil; e = e.Next() {
// 		k := e.Key
// 		res = append(res, fmt.Sprintf("%q:", k)...)
// 		var b []byte
// 		b, err = json.Marshal(m.kv[k])
// 		if err != nil {
// 			return
// 		}
// 		res = append(res, b...)
// 		if e != back {
// 			res = append(res, ',')
// 		}
// 	}
// 	res = append(res, '}')
// 	// fmt.Printf("marshalled: %v: %#v\n", res, res)
// 	return
// }

// UnmarshalJSON implements type json.Unmarshaler interface, so can be called in json.Unmarshal(data, om)
// func (m *OrderedMap[V]) UnmarshalJSON(data []byte) error {
// 	dec := json.NewDecoder(bytes.NewReader(data))
// 	dec.UseNumber()

// 	// must open with a delim token '{'
// 	t, err := dec.Token()
// 	if err != nil {
// 		return err
// 	}
// 	if delim, ok := t.(json.Delim); !ok || delim != '{' {
// 		return fmt.Errorf("expect JSON object open with '{'")
// 	}

// 	err = m.parseobject(dec)
// 	if err != nil {
// 		return err
// 	}

// 	t, err = dec.Token()
// 	if err != io.EOF {
// 		return fmt.Errorf("expect end of JSON object but got more token: %T: %v or err: %v", t, t, err)
// 	}

// 	return nil
// }

// Delete will remove a key from the map. It will return true if the key was
// removed (the key did exist).
func (__obf_2403dfddde3862ad *OrderedMap[V]) Delete(__obf_797707a3dac0ebb7 string) (__obf_beedee7d7bd2c208 bool) {
	__obf_8338fba3de46db62, __obf_c7b01e58066bbb7a := __obf_2403dfddde3862ad.__obf_aa52d12db26ef244[__obf_797707a3dac0ebb7]
	if __obf_c7b01e58066bbb7a {
		__obf_2403dfddde3862ad.__obf_c451bc9b9d85eff5.Remove(__obf_8338fba3de46db62)
		delete(__obf_2403dfddde3862ad.__obf_aa52d12db26ef244, __obf_797707a3dac0ebb7)
	}

	return __obf_c7b01e58066bbb7a
}

// Front will return the element that is the first (oldest Set element). If
// there are no elements this will return nil.
func (__obf_2403dfddde3862ad *OrderedMap[V]) Front() *Element[V] {
	return __obf_2403dfddde3862ad.__obf_c451bc9b9d85eff5.Front()
}

// Back will return the element that is the last (most recent Set element). If
// there are no elements this will return nil.
func (__obf_2403dfddde3862ad *OrderedMap[V]) Back() *Element[V] {
	return __obf_2403dfddde3862ad.__obf_c451bc9b9d85eff5.Back()
}

// Copy returns a new OrderedMap with the same elements.
// Using Copy while there are concurrent writes may mangle the result.
func (__obf_2403dfddde3862ad *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_44aa38149f008e1d := NewOrderedMap[V]()
	for __obf_23f1e93b8e98e9d5 := __obf_2403dfddde3862ad.Front(); __obf_23f1e93b8e98e9d5 != nil; __obf_23f1e93b8e98e9d5 = __obf_23f1e93b8e98e9d5.Next() {
		__obf_44aa38149f008e1d.Set(__obf_23f1e93b8e98e9d5.Key, __obf_23f1e93b8e98e9d5.Value)
	}
	return __obf_44aa38149f008e1d
}

// SetFromJson set element from json string
func (__obf_2403dfddde3862ad *OrderedMap[V]) SetFromJson(__obf_ecf61ed96b8ba849 string) (__obf_52943f8b1bce7293 error) {
	__obf_faac983e1e79a41e := make(map[string]V, 0)
	__obf_52943f8b1bce7293 = json.Unmarshal([]byte(__obf_ecf61ed96b8ba849), &__obf_faac983e1e79a41e)
	if __obf_52943f8b1bce7293 != nil {
		return
	}
	for __obf_a0c6719477333085, __obf_1e75fd93d2a55aa7 := range __obf_faac983e1e79a41e {
		__obf_2403dfddde3862ad.Set(__obf_a0c6719477333085, __obf_1e75fd93d2a55aa7)
	}
	return
}

// func (m *OrderedMap[V]) parseobject(dec *json.Decoder) (err error) {
// 	var t json.Token
// 	for dec.More() {
// 		t, err = dec.Token()
// 		if err != nil {
// 			return err
// 		}

// 		key, ok := t.(string)
// 		if !ok {
// 			return fmt.Errorf("expecting JSON key should be always a string: %T: %v", t, t)
// 		}

// 		t, err = dec.Token()
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			return
// 		}

// 		var value V

// 		if delim, ok := t.(json.Delim); ok {
// 			if delim != '{' {
// 				return fmt.Errorf("unexpected delimiter: %q", delim)
// 			}
// 			om2 := NewOrderedMap[V]()
// 			err = om2.parseobject(dec)
// 			if err != nil {
// 				return
// 			}
// 			value = om2
// 		} else {
// 			value = t
// 		}

// 		// m.keys = append(m.keys, key)
// 		m.kv[key] = m.ll.PushBack(key, value)
// 		// m.m[key] = value
// 	}

// 	t, err = dec.Token()
// 	if err != nil {
// 		return
// 	}
// 	if delim, ok := t.(json.Delim); !ok || delim != '}' {
// 		return fmt.Errorf("expect JSON object closeany with '}'")
// 	}

// 	return
// }
