package __obf_62eba4024f8fa381

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_3ca757edadb508a4 map[string]*Element[V]
	__obf_2cee61f9518c65b8 __obf_40b36c6e023d30bc[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{
		__obf_3ca757edadb508a4: make(map[string]*Element[V]),
	}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_ca09540de29917b3 *OrderedMap[V]) Get(__obf_df070ab4c712506c string) (__obf_dbc685b0ad464e4d V, __obf_b47c9f51c1edd10c bool) {
	__obf_f69c1647001964ce, __obf_b47c9f51c1edd10c := __obf_ca09540de29917b3.__obf_3ca757edadb508a4[__obf_df070ab4c712506c]
	if __obf_b47c9f51c1edd10c {
		__obf_dbc685b0ad464e4d = __obf_f69c1647001964ce.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_ca09540de29917b3 *OrderedMap[V]) Set(__obf_df070ab4c712506c string, __obf_dbc685b0ad464e4d V) bool {
	_, __obf_ebe2022fc001ac5f := __obf_ca09540de29917b3.__obf_3ca757edadb508a4[__obf_df070ab4c712506c]
	if __obf_ebe2022fc001ac5f {
		__obf_ca09540de29917b3.__obf_3ca757edadb508a4[__obf_df070ab4c712506c].Value = __obf_dbc685b0ad464e4d
		return false
	}

	__obf_b5a80c614a2f738f := __obf_ca09540de29917b3.__obf_2cee61f9518c65b8.PushBack(__obf_df070ab4c712506c, __obf_dbc685b0ad464e4d)
	__obf_ca09540de29917b3.__obf_3ca757edadb508a4[__obf_df070ab4c712506c] = __obf_b5a80c614a2f738f
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_ca09540de29917b3 *OrderedMap[V]) GetOrDefault(__obf_df070ab4c712506c string, __obf_383ddce67550f7a1 V) V {
	if __obf_dbc685b0ad464e4d, __obf_b47c9f51c1edd10c := __obf_ca09540de29917b3.__obf_3ca757edadb508a4[__obf_df070ab4c712506c]; __obf_b47c9f51c1edd10c {
		return __obf_dbc685b0ad464e4d.Value
	}

	return __obf_383ddce67550f7a1
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_ca09540de29917b3 *OrderedMap[V]) GetElement(__obf_df070ab4c712506c string) *Element[V] {
	__obf_b5a80c614a2f738f, __obf_b47c9f51c1edd10c := __obf_ca09540de29917b3.__obf_3ca757edadb508a4[__obf_df070ab4c712506c]
	if __obf_b47c9f51c1edd10c {
		return __obf_b5a80c614a2f738f
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_ca09540de29917b3 *OrderedMap[V]) Len() int {
	return len(__obf_ca09540de29917b3.__obf_3ca757edadb508a4)
}

// Keys returns all of the keys in the order they were inserted. If a key was
// replaced it will retain the same position. To ensure most recently set keys
// are always at the end you must always Delete before Set.
func (__obf_ca09540de29917b3 *OrderedMap[V]) Keys() (__obf_ea1c0c515eb08cda []string) {
	__obf_ea1c0c515eb08cda = make([]string, 0, __obf_ca09540de29917b3.Len())
	for __obf_606ad3c69469507e := __obf_ca09540de29917b3.Front(); __obf_606ad3c69469507e != nil; __obf_606ad3c69469507e = __obf_606ad3c69469507e.Next() {
		__obf_ea1c0c515eb08cda = append(__obf_ea1c0c515eb08cda, __obf_606ad3c69469507e.Key)
	}
	return __obf_ea1c0c515eb08cda
}

// Values returns all of the values in the order they were inserted.
func (__obf_ca09540de29917b3 *OrderedMap[V]) Values() (__obf_cfff6b37b91c3027 []V) {
	__obf_cfff6b37b91c3027 = make([]V, 0, __obf_ca09540de29917b3.Len())
	for __obf_606ad3c69469507e := __obf_ca09540de29917b3.Front(); __obf_606ad3c69469507e != nil; __obf_606ad3c69469507e = __obf_606ad3c69469507e.Next() {
		__obf_cfff6b37b91c3027 = append(__obf_cfff6b37b91c3027, __obf_606ad3c69469507e.Value)
	}
	return __obf_cfff6b37b91c3027
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
func (__obf_ca09540de29917b3 *OrderedMap[V]) Delete(__obf_df070ab4c712506c string) (__obf_1ce39dd42c37bd04 bool) {
	__obf_b5a80c614a2f738f, __obf_b47c9f51c1edd10c := __obf_ca09540de29917b3.__obf_3ca757edadb508a4[__obf_df070ab4c712506c]
	if __obf_b47c9f51c1edd10c {
		__obf_ca09540de29917b3.__obf_2cee61f9518c65b8.Remove(__obf_b5a80c614a2f738f)
		delete(__obf_ca09540de29917b3.__obf_3ca757edadb508a4, __obf_df070ab4c712506c)
	}

	return __obf_b47c9f51c1edd10c
}

// Front will return the element that is the first (oldest Set element). If
// there are no elements this will return nil.
func (__obf_ca09540de29917b3 *OrderedMap[V]) Front() *Element[V] {
	return __obf_ca09540de29917b3.__obf_2cee61f9518c65b8.Front()
}

// Back will return the element that is the last (most recent Set element). If
// there are no elements this will return nil.
func (__obf_ca09540de29917b3 *OrderedMap[V]) Back() *Element[V] {
	return __obf_ca09540de29917b3.__obf_2cee61f9518c65b8.Back()
}

// Copy returns a new OrderedMap with the same elements.
// Using Copy while there are concurrent writes may mangle the result.
func (__obf_ca09540de29917b3 *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_e52cbcfde3d3e2da := NewOrderedMap[V]()
	for __obf_606ad3c69469507e := __obf_ca09540de29917b3.Front(); __obf_606ad3c69469507e != nil; __obf_606ad3c69469507e = __obf_606ad3c69469507e.Next() {
		__obf_e52cbcfde3d3e2da.Set(__obf_606ad3c69469507e.Key, __obf_606ad3c69469507e.Value)
	}
	return __obf_e52cbcfde3d3e2da
}

// SetFromJson set element from json string
func (__obf_ca09540de29917b3 *OrderedMap[V]) SetFromJson(__obf_1524bc1cc35c4ef3 string) (__obf_a4a3a0fe86b36626 error) {
	__obf_b97f8243a437c987 := make(map[string]V, 0)
	__obf_a4a3a0fe86b36626 = json.Unmarshal([]byte(__obf_1524bc1cc35c4ef3), &__obf_b97f8243a437c987)
	if __obf_a4a3a0fe86b36626 != nil {
		return
	}
	for __obf_f067a403d3b1b52a, __obf_f69c1647001964ce := range __obf_b97f8243a437c987 {
		__obf_ca09540de29917b3.Set(__obf_f067a403d3b1b52a, __obf_f69c1647001964ce)
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
