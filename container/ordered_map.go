package __obf_76599ab96a208083

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_bf17359d9f0869c4 map[string]*Element[V]
	__obf_fbed25d2488ac00f __obf_3bdbb0fb35d34874[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{
		__obf_bf17359d9f0869c4: make(map[string]*Element[V]),
	}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_6859967c4af30821 *OrderedMap[V]) Get(__obf_4368715d0a6d4f0b string) (__obf_209abc4c4955f6d3 V, __obf_d47b43f0df7ca990 bool) {
	__obf_93d90cff8a0252e8, __obf_d47b43f0df7ca990 := __obf_6859967c4af30821.__obf_bf17359d9f0869c4[__obf_4368715d0a6d4f0b]
	if __obf_d47b43f0df7ca990 {
		__obf_209abc4c4955f6d3 = __obf_93d90cff8a0252e8.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_6859967c4af30821 *OrderedMap[V]) Set(__obf_4368715d0a6d4f0b string, __obf_209abc4c4955f6d3 V) bool {
	_, __obf_5fdfb83510f7dd73 := __obf_6859967c4af30821.__obf_bf17359d9f0869c4[__obf_4368715d0a6d4f0b]
	if __obf_5fdfb83510f7dd73 {
		__obf_6859967c4af30821.__obf_bf17359d9f0869c4[__obf_4368715d0a6d4f0b].Value = __obf_209abc4c4955f6d3
		return false
	}

	__obf_d2b7bc1f73131cd3 := __obf_6859967c4af30821.__obf_fbed25d2488ac00f.PushBack(__obf_4368715d0a6d4f0b, __obf_209abc4c4955f6d3)
	__obf_6859967c4af30821.__obf_bf17359d9f0869c4[__obf_4368715d0a6d4f0b] = __obf_d2b7bc1f73131cd3
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_6859967c4af30821 *OrderedMap[V]) GetOrDefault(__obf_4368715d0a6d4f0b string, __obf_41086f4887f5f3a5 V) V {
	if __obf_209abc4c4955f6d3, __obf_d47b43f0df7ca990 := __obf_6859967c4af30821.__obf_bf17359d9f0869c4[__obf_4368715d0a6d4f0b]; __obf_d47b43f0df7ca990 {
		return __obf_209abc4c4955f6d3.Value
	}

	return __obf_41086f4887f5f3a5
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_6859967c4af30821 *OrderedMap[V]) GetElement(__obf_4368715d0a6d4f0b string) *Element[V] {
	__obf_d2b7bc1f73131cd3, __obf_d47b43f0df7ca990 := __obf_6859967c4af30821.__obf_bf17359d9f0869c4[__obf_4368715d0a6d4f0b]
	if __obf_d47b43f0df7ca990 {
		return __obf_d2b7bc1f73131cd3
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_6859967c4af30821 *OrderedMap[V]) Len() int {
	return len(__obf_6859967c4af30821.__obf_bf17359d9f0869c4)
}

// Keys returns all of the keys in the order they were inserted. If a key was
// replaced it will retain the same position. To ensure most recently set keys
// are always at the end you must always Delete before Set.
func (__obf_6859967c4af30821 *OrderedMap[V]) Keys() (__obf_cfa0937a05242cbc []string) {
	__obf_cfa0937a05242cbc = make([]string, 0, __obf_6859967c4af30821.Len())
	for __obf_3cc72389c8ded9a2 := __obf_6859967c4af30821.Front(); __obf_3cc72389c8ded9a2 != nil; __obf_3cc72389c8ded9a2 = __obf_3cc72389c8ded9a2.Next() {
		__obf_cfa0937a05242cbc = append(__obf_cfa0937a05242cbc, __obf_3cc72389c8ded9a2.Key)
	}
	return __obf_cfa0937a05242cbc
}

// Values returns all of the values in the order they were inserted.
func (__obf_6859967c4af30821 *OrderedMap[V]) Values() (__obf_f2c3ba527a3a24ba []V) {
	__obf_f2c3ba527a3a24ba = make([]V, 0, __obf_6859967c4af30821.Len())
	for __obf_3cc72389c8ded9a2 := __obf_6859967c4af30821.Front(); __obf_3cc72389c8ded9a2 != nil; __obf_3cc72389c8ded9a2 = __obf_3cc72389c8ded9a2.Next() {
		__obf_f2c3ba527a3a24ba = append(__obf_f2c3ba527a3a24ba, __obf_3cc72389c8ded9a2.Value)
	}
	return __obf_f2c3ba527a3a24ba
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
func (__obf_6859967c4af30821 *OrderedMap[V]) Delete(__obf_4368715d0a6d4f0b string) (__obf_32f95d58072d8869 bool) {
	__obf_d2b7bc1f73131cd3, __obf_d47b43f0df7ca990 := __obf_6859967c4af30821.__obf_bf17359d9f0869c4[__obf_4368715d0a6d4f0b]
	if __obf_d47b43f0df7ca990 {
		__obf_6859967c4af30821.__obf_fbed25d2488ac00f.Remove(__obf_d2b7bc1f73131cd3)
		delete(__obf_6859967c4af30821.__obf_bf17359d9f0869c4, __obf_4368715d0a6d4f0b)
	}

	return __obf_d47b43f0df7ca990
}

// Front will return the element that is the first (oldest Set element). If
// there are no elements this will return nil.
func (__obf_6859967c4af30821 *OrderedMap[V]) Front() *Element[V] {
	return __obf_6859967c4af30821.__obf_fbed25d2488ac00f.Front()
}

// Back will return the element that is the last (most recent Set element). If
// there are no elements this will return nil.
func (__obf_6859967c4af30821 *OrderedMap[V]) Back() *Element[V] {
	return __obf_6859967c4af30821.__obf_fbed25d2488ac00f.Back()
}

// Copy returns a new OrderedMap with the same elements.
// Using Copy while there are concurrent writes may mangle the result.
func (__obf_6859967c4af30821 *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_ced0528d9bc49e08 := NewOrderedMap[V]()
	for __obf_3cc72389c8ded9a2 := __obf_6859967c4af30821.Front(); __obf_3cc72389c8ded9a2 != nil; __obf_3cc72389c8ded9a2 = __obf_3cc72389c8ded9a2.Next() {
		__obf_ced0528d9bc49e08.Set(__obf_3cc72389c8ded9a2.Key, __obf_3cc72389c8ded9a2.Value)
	}
	return __obf_ced0528d9bc49e08
}

// SetFromJson set element from json string
func (__obf_6859967c4af30821 *OrderedMap[V]) SetFromJson(__obf_81bddcf8753e1c07 string) (__obf_791dcfd55f252396 error) {
	__obf_efd9c2e0133924bf := make(map[string]V, 0)
	__obf_791dcfd55f252396 = json.Unmarshal([]byte(__obf_81bddcf8753e1c07), &__obf_efd9c2e0133924bf)
	if __obf_791dcfd55f252396 != nil {
		return
	}
	for __obf_2971938b3f770fd1, __obf_93d90cff8a0252e8 := range __obf_efd9c2e0133924bf {
		__obf_6859967c4af30821.Set(__obf_2971938b3f770fd1, __obf_93d90cff8a0252e8)
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
