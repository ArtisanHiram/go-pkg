package __obf_1fda7fbdeda52f1e

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_ad77162cacd39bac map[string]*Element[V]
	__obf_3c15e43fbd1a9166 __obf_5275791a072a11d1[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{
		__obf_ad77162cacd39bac: make(map[string]*Element[V]),
	}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_7c84111c9ad51184 *OrderedMap[V]) Get(__obf_95c4c09a56c1a070 string) (__obf_6d18c47092a65f53 V, __obf_d24c61455fa2e4db bool) {
	__obf_9371496cccb864c5, __obf_d24c61455fa2e4db := __obf_7c84111c9ad51184.__obf_ad77162cacd39bac[__obf_95c4c09a56c1a070]
	if __obf_d24c61455fa2e4db {
		__obf_6d18c47092a65f53 = __obf_9371496cccb864c5.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_7c84111c9ad51184 *OrderedMap[V]) Set(__obf_95c4c09a56c1a070 string, __obf_6d18c47092a65f53 V) bool {
	_, __obf_0bcee1844fb8e67d := __obf_7c84111c9ad51184.__obf_ad77162cacd39bac[__obf_95c4c09a56c1a070]
	if __obf_0bcee1844fb8e67d {
		__obf_7c84111c9ad51184.__obf_ad77162cacd39bac[__obf_95c4c09a56c1a070].Value = __obf_6d18c47092a65f53
		return false
	}

	__obf_c205dffe0337941a := __obf_7c84111c9ad51184.__obf_3c15e43fbd1a9166.PushBack(__obf_95c4c09a56c1a070, __obf_6d18c47092a65f53)
	__obf_7c84111c9ad51184.__obf_ad77162cacd39bac[__obf_95c4c09a56c1a070] = __obf_c205dffe0337941a
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_7c84111c9ad51184 *OrderedMap[V]) GetOrDefault(__obf_95c4c09a56c1a070 string, __obf_b3e6a7c1a508f73f V) V {
	if __obf_6d18c47092a65f53, __obf_d24c61455fa2e4db := __obf_7c84111c9ad51184.__obf_ad77162cacd39bac[__obf_95c4c09a56c1a070]; __obf_d24c61455fa2e4db {
		return __obf_6d18c47092a65f53.Value
	}

	return __obf_b3e6a7c1a508f73f
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_7c84111c9ad51184 *OrderedMap[V]) GetElement(__obf_95c4c09a56c1a070 string) *Element[V] {
	__obf_c205dffe0337941a, __obf_d24c61455fa2e4db := __obf_7c84111c9ad51184.__obf_ad77162cacd39bac[__obf_95c4c09a56c1a070]
	if __obf_d24c61455fa2e4db {
		return __obf_c205dffe0337941a
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_7c84111c9ad51184 *OrderedMap[V]) Len() int {
	return len(__obf_7c84111c9ad51184.__obf_ad77162cacd39bac)
}

// Keys returns all of the keys in the order they were inserted. If a key was
// replaced it will retain the same position. To ensure most recently set keys
// are always at the end you must always Delete before Set.
func (__obf_7c84111c9ad51184 *OrderedMap[V]) Keys() (__obf_f33bf750ed8ea535 []string) {
	__obf_f33bf750ed8ea535 = make([]string, 0, __obf_7c84111c9ad51184.Len())
	for __obf_dc39c698a857082c := __obf_7c84111c9ad51184.Front(); __obf_dc39c698a857082c != nil; __obf_dc39c698a857082c = __obf_dc39c698a857082c.Next() {
		__obf_f33bf750ed8ea535 = append(__obf_f33bf750ed8ea535, __obf_dc39c698a857082c.Key)
	}
	return __obf_f33bf750ed8ea535
}

// Values returns all of the values in the order they were inserted.
func (__obf_7c84111c9ad51184 *OrderedMap[V]) Values() (__obf_c6f1eaada6fb59d9 []V) {
	__obf_c6f1eaada6fb59d9 = make([]V, 0, __obf_7c84111c9ad51184.Len())
	for __obf_dc39c698a857082c := __obf_7c84111c9ad51184.Front(); __obf_dc39c698a857082c != nil; __obf_dc39c698a857082c = __obf_dc39c698a857082c.Next() {
		__obf_c6f1eaada6fb59d9 = append(__obf_c6f1eaada6fb59d9, __obf_dc39c698a857082c.Value)
	}
	return __obf_c6f1eaada6fb59d9
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
func (__obf_7c84111c9ad51184 *OrderedMap[V]) Delete(__obf_95c4c09a56c1a070 string) (__obf_69fd4db032b1b822 bool) {
	__obf_c205dffe0337941a, __obf_d24c61455fa2e4db := __obf_7c84111c9ad51184.__obf_ad77162cacd39bac[__obf_95c4c09a56c1a070]
	if __obf_d24c61455fa2e4db {
		__obf_7c84111c9ad51184.__obf_3c15e43fbd1a9166.Remove(__obf_c205dffe0337941a)
		delete(__obf_7c84111c9ad51184.__obf_ad77162cacd39bac, __obf_95c4c09a56c1a070)
	}

	return __obf_d24c61455fa2e4db
}

// Front will return the element that is the first (oldest Set element). If
// there are no elements this will return nil.
func (__obf_7c84111c9ad51184 *OrderedMap[V]) Front() *Element[V] {
	return __obf_7c84111c9ad51184.__obf_3c15e43fbd1a9166.Front()
}

// Back will return the element that is the last (most recent Set element). If
// there are no elements this will return nil.
func (__obf_7c84111c9ad51184 *OrderedMap[V]) Back() *Element[V] {
	return __obf_7c84111c9ad51184.__obf_3c15e43fbd1a9166.Back()
}

// Copy returns a new OrderedMap with the same elements.
// Using Copy while there are concurrent writes may mangle the result.
func (__obf_7c84111c9ad51184 *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_46d313f9ed0bfa48 := NewOrderedMap[V]()
	for __obf_dc39c698a857082c := __obf_7c84111c9ad51184.Front(); __obf_dc39c698a857082c != nil; __obf_dc39c698a857082c = __obf_dc39c698a857082c.Next() {
		__obf_46d313f9ed0bfa48.Set(__obf_dc39c698a857082c.Key, __obf_dc39c698a857082c.Value)
	}
	return __obf_46d313f9ed0bfa48
}

// SetFromJson set element from json string
func (__obf_7c84111c9ad51184 *OrderedMap[V]) SetFromJson(__obf_713d33afc78e8029 string) (__obf_483c2f6b078721a6 error) {
	__obf_73f80e4a6e7babfe := make(map[string]V, 0)
	__obf_483c2f6b078721a6 = json.Unmarshal([]byte(__obf_713d33afc78e8029), &__obf_73f80e4a6e7babfe)
	if __obf_483c2f6b078721a6 != nil {
		return
	}
	for __obf_9058efdba9a0277b, __obf_9371496cccb864c5 := range __obf_73f80e4a6e7babfe {
		__obf_7c84111c9ad51184.Set(__obf_9058efdba9a0277b, __obf_9371496cccb864c5)
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
