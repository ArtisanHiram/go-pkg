package __obf_af42fb6cde2beed6

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_a5e6e278a2b2fda1 map[string]*Element[V]
	__obf_4090089c7d247118 __obf_543c4aff1fda48f3[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{
		__obf_a5e6e278a2b2fda1: make(map[string]*Element[V]),
	}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_88bdf57825ea07fa *OrderedMap[V]) Get(__obf_55cde42f6d47c5be string) (__obf_9ed1ee5d2b51057a V, __obf_021756766e1db5da bool) {
	__obf_289dd915247ca25d, __obf_021756766e1db5da := __obf_88bdf57825ea07fa.__obf_a5e6e278a2b2fda1[__obf_55cde42f6d47c5be]
	if __obf_021756766e1db5da {
		__obf_9ed1ee5d2b51057a = __obf_289dd915247ca25d.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_88bdf57825ea07fa *OrderedMap[V]) Set(__obf_55cde42f6d47c5be string, __obf_9ed1ee5d2b51057a V) bool {
	_, __obf_5e8979ecf9b9d2dc := __obf_88bdf57825ea07fa.__obf_a5e6e278a2b2fda1[__obf_55cde42f6d47c5be]
	if __obf_5e8979ecf9b9d2dc {
		__obf_88bdf57825ea07fa.__obf_a5e6e278a2b2fda1[__obf_55cde42f6d47c5be].Value = __obf_9ed1ee5d2b51057a
		return false
	}

	__obf_20aeab0189b186c0 := __obf_88bdf57825ea07fa.__obf_4090089c7d247118.PushBack(__obf_55cde42f6d47c5be, __obf_9ed1ee5d2b51057a)
	__obf_88bdf57825ea07fa.__obf_a5e6e278a2b2fda1[__obf_55cde42f6d47c5be] = __obf_20aeab0189b186c0
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_88bdf57825ea07fa *OrderedMap[V]) GetOrDefault(__obf_55cde42f6d47c5be string, __obf_9e442b6a8851748b V) V {
	if __obf_9ed1ee5d2b51057a, __obf_021756766e1db5da := __obf_88bdf57825ea07fa.__obf_a5e6e278a2b2fda1[__obf_55cde42f6d47c5be]; __obf_021756766e1db5da {
		return __obf_9ed1ee5d2b51057a.Value
	}

	return __obf_9e442b6a8851748b
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_88bdf57825ea07fa *OrderedMap[V]) GetElement(__obf_55cde42f6d47c5be string) *Element[V] {
	__obf_20aeab0189b186c0, __obf_021756766e1db5da := __obf_88bdf57825ea07fa.__obf_a5e6e278a2b2fda1[__obf_55cde42f6d47c5be]
	if __obf_021756766e1db5da {
		return __obf_20aeab0189b186c0
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_88bdf57825ea07fa *OrderedMap[V]) Len() int {
	return len(__obf_88bdf57825ea07fa.__obf_a5e6e278a2b2fda1)
}

// Keys returns all of the keys in the order they were inserted. If a key was
// replaced it will retain the same position. To ensure most recently set keys
// are always at the end you must always Delete before Set.
func (__obf_88bdf57825ea07fa *OrderedMap[V]) Keys() (__obf_a0cee0a97b92008e []string) {
	__obf_a0cee0a97b92008e = make([]string, 0, __obf_88bdf57825ea07fa.Len())
	for __obf_91551c548d1b0a9e := __obf_88bdf57825ea07fa.Front(); __obf_91551c548d1b0a9e != nil; __obf_91551c548d1b0a9e = __obf_91551c548d1b0a9e.Next() {
		__obf_a0cee0a97b92008e = append(__obf_a0cee0a97b92008e, __obf_91551c548d1b0a9e.Key)
	}
	return __obf_a0cee0a97b92008e
}

// Values returns all of the values in the order they were inserted.
func (__obf_88bdf57825ea07fa *OrderedMap[V]) Values() (__obf_ecffc655bb698ea6 []V) {
	__obf_ecffc655bb698ea6 = make([]V, 0, __obf_88bdf57825ea07fa.Len())
	for __obf_91551c548d1b0a9e := __obf_88bdf57825ea07fa.Front(); __obf_91551c548d1b0a9e != nil; __obf_91551c548d1b0a9e = __obf_91551c548d1b0a9e.Next() {
		__obf_ecffc655bb698ea6 = append(__obf_ecffc655bb698ea6, __obf_91551c548d1b0a9e.Value)
	}
	return __obf_ecffc655bb698ea6
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
func (__obf_88bdf57825ea07fa *OrderedMap[V]) Delete(__obf_55cde42f6d47c5be string) (__obf_6934236d0776cd3a bool) {
	__obf_20aeab0189b186c0, __obf_021756766e1db5da := __obf_88bdf57825ea07fa.__obf_a5e6e278a2b2fda1[__obf_55cde42f6d47c5be]
	if __obf_021756766e1db5da {
		__obf_88bdf57825ea07fa.__obf_4090089c7d247118.Remove(__obf_20aeab0189b186c0)
		delete(__obf_88bdf57825ea07fa.__obf_a5e6e278a2b2fda1, __obf_55cde42f6d47c5be)
	}

	return __obf_021756766e1db5da
}

// Front will return the element that is the first (oldest Set element). If
// there are no elements this will return nil.
func (__obf_88bdf57825ea07fa *OrderedMap[V]) Front() *Element[V] {
	return __obf_88bdf57825ea07fa.__obf_4090089c7d247118.Front()
}

// Back will return the element that is the last (most recent Set element). If
// there are no elements this will return nil.
func (__obf_88bdf57825ea07fa *OrderedMap[V]) Back() *Element[V] {
	return __obf_88bdf57825ea07fa.__obf_4090089c7d247118.Back()
}

// Copy returns a new OrderedMap with the same elements.
// Using Copy while there are concurrent writes may mangle the result.
func (__obf_88bdf57825ea07fa *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_253bbe7abf639fb6 := NewOrderedMap[V]()
	for __obf_91551c548d1b0a9e := __obf_88bdf57825ea07fa.Front(); __obf_91551c548d1b0a9e != nil; __obf_91551c548d1b0a9e = __obf_91551c548d1b0a9e.Next() {
		__obf_253bbe7abf639fb6.Set(__obf_91551c548d1b0a9e.Key, __obf_91551c548d1b0a9e.Value)
	}
	return __obf_253bbe7abf639fb6
}

// SetFromJson set element from json string
func (__obf_88bdf57825ea07fa *OrderedMap[V]) SetFromJson(__obf_2f7f050797461426 string) (__obf_24456bd3f3e29186 error) {
	__obf_e9274782a8b581db := make(map[string]V, 0)
	__obf_24456bd3f3e29186 = json.Unmarshal([]byte(__obf_2f7f050797461426), &__obf_e9274782a8b581db)
	if __obf_24456bd3f3e29186 != nil {
		return
	}
	for __obf_43d363f3f5d57e86, __obf_289dd915247ca25d := range __obf_e9274782a8b581db {
		__obf_88bdf57825ea07fa.Set(__obf_43d363f3f5d57e86, __obf_289dd915247ca25d)
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
