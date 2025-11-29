package __obf_90a4883a02d0b41c

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_b7a331088c2af1df map[string]*Element[V]
	__obf_61fee515419381f8 __obf_3947ec1119a8bcb1[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{__obf_b7a331088c2af1df: make(map[string]*Element[V])}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_d75fab052d636194 *OrderedMap[V]) Get(__obf_5bba24c1758bbf28 string) (__obf_934b4769b75de0d4 V, __obf_c16a0dd3a76f078a bool) {
	__obf_1c35129ed9e54186, __obf_c16a0dd3a76f078a := __obf_d75fab052d636194.__obf_b7a331088c2af1df[__obf_5bba24c1758bbf28]
	if __obf_c16a0dd3a76f078a {
		__obf_934b4769b75de0d4 = __obf_1c35129ed9e54186.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_d75fab052d636194 *OrderedMap[V]) Set(__obf_5bba24c1758bbf28 string, __obf_934b4769b75de0d4 V) bool {
	_, __obf_fc145ba90508da80 := __obf_d75fab052d636194.__obf_b7a331088c2af1df[__obf_5bba24c1758bbf28]
	if __obf_fc145ba90508da80 {
		__obf_d75fab052d636194.__obf_b7a331088c2af1df[__obf_5bba24c1758bbf28].Value = __obf_934b4769b75de0d4
		return false
	}
	__obf_f469067f729b9098 := __obf_d75fab052d636194.__obf_61fee515419381f8.PushBack(__obf_5bba24c1758bbf28, __obf_934b4769b75de0d4)
	__obf_d75fab052d636194.__obf_b7a331088c2af1df[__obf_5bba24c1758bbf28] = __obf_f469067f729b9098
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_d75fab052d636194 *OrderedMap[V]) GetOrDefault(__obf_5bba24c1758bbf28 string, __obf_d3dd2d8281806f3e V) V {
	if __obf_934b4769b75de0d4, __obf_c16a0dd3a76f078a := __obf_d75fab052d636194.__obf_b7a331088c2af1df[__obf_5bba24c1758bbf28]; __obf_c16a0dd3a76f078a {
		return __obf_934b4769b75de0d4.Value
	}

	return __obf_d3dd2d8281806f3e
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_d75fab052d636194 *OrderedMap[V]) GetElement(__obf_5bba24c1758bbf28 string) *Element[V] {
	__obf_f469067f729b9098, __obf_c16a0dd3a76f078a := __obf_d75fab052d636194.__obf_b7a331088c2af1df[__obf_5bba24c1758bbf28]
	if __obf_c16a0dd3a76f078a {
		return __obf_f469067f729b9098
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_d75fab052d636194 *OrderedMap[V]) Len() int {
	return len(__obf_d75fab052d636194.

		// Keys returns all of the keys in the order they were inserted. If a key was
		// replaced it will retain the same position. To ensure most recently set keys
		// are always at the end you must always Delete before Set.
		__obf_b7a331088c2af1df)
}

func (__obf_d75fab052d636194 *OrderedMap[V]) Keys() (__obf_f72137bb13d161c8 []string) {
	__obf_f72137bb13d161c8 = make([]string, 0, __obf_d75fab052d636194.Len())
	for __obf_8084077fee117687 := __obf_d75fab052d636194.Front(); __obf_8084077fee117687 != nil; __obf_8084077fee117687 = __obf_8084077fee117687.Next() {
		__obf_f72137bb13d161c8 = append(__obf_f72137bb13d161c8, __obf_8084077fee117687.Key)
	}
	return __obf_f72137bb13d161c8
}

// Values returns all of the values in the order they were inserted.
func (__obf_d75fab052d636194 *OrderedMap[V]) Values() (__obf_c20d68262da83750 []V) {
	__obf_c20d68262da83750 = make([]V, 0, __obf_d75fab052d636194.Len())
	for __obf_8084077fee117687 := __obf_d75fab052d636194.Front(); __obf_8084077fee117687 != nil; __obf_8084077fee117687 = __obf_8084077fee117687.Next() {
		__obf_c20d68262da83750 = append(__obf_c20d68262da83750, __obf_8084077fee117687.Value)
	}
	return __obf_c20d68262da83750
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
func (__obf_d75fab052d636194 *OrderedMap[V]) Delete(__obf_5bba24c1758bbf28 string) (__obf_23dab5fef08a0f9b bool) {
	__obf_f469067f729b9098, __obf_c16a0dd3a76f078a := __obf_d75fab052d636194.__obf_b7a331088c2af1df[__obf_5bba24c1758bbf28]
	if __obf_c16a0dd3a76f078a {
		__obf_d75fab052d636194.__obf_61fee515419381f8.
			Remove(__obf_f469067f729b9098)
		delete(__obf_d75fab052d636194.__obf_b7a331088c2af1df,

			// Front will return the element that is the first (oldest Set element). If
			// there are no elements this will return nil.
			__obf_5bba24c1758bbf28)
	}

	return __obf_c16a0dd3a76f078a
}

func (__obf_d75fab052d636194 *OrderedMap[V]) Front() *Element[V] {
	return __obf_d75fab052d636194.

		// Back will return the element that is the last (most recent Set element). If
		// there are no elements this will return nil.
		__obf_61fee515419381f8.Front()
}

func (__obf_d75fab052d636194 *OrderedMap[V]) Back() *Element[V] {
	return __obf_d75fab052d636194.

		// Copy returns a new OrderedMap with the same elements.
		// Using Copy while there are concurrent writes may mangle the result.
		__obf_61fee515419381f8.Back()
}

func (__obf_d75fab052d636194 *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_bda76c937068ccd0 := NewOrderedMap[V]()
	for __obf_8084077fee117687 := __obf_d75fab052d636194.Front(); __obf_8084077fee117687 != nil; __obf_8084077fee117687 = __obf_8084077fee117687.Next() {
		__obf_bda76c937068ccd0.
			Set(__obf_8084077fee117687.Key, __obf_8084077fee117687.Value)
	}
	return __obf_bda76c937068ccd0
}

// SetFromJson set element from json string
func (__obf_d75fab052d636194 *OrderedMap[V]) SetFromJson(__obf_04f88f426f5352e2 string) (__obf_f14c805fc0e38747 error) {
	__obf_560ae7f29c595dfe := make(map[string]V, 0)
	__obf_f14c805fc0e38747 = json.Unmarshal([]byte(__obf_04f88f426f5352e2), &__obf_560ae7f29c595dfe)
	if __obf_f14c805fc0e38747 != nil {
		return
	}
	for __obf_7d788ca4b6a826e5, __obf_1c35129ed9e54186 := range __obf_560ae7f29c595dfe {
		__obf_d75fab052d636194.
			Set(__obf_7d788ca4b6a826e5,

				// func (m *OrderedMap[V]) parseobject(dec *json.Decoder) (err error) {
				// 	var t json.Token
				// 	for dec.More() {
				// 		t, err = dec.Token()
				// 		if err != nil {
				// 			return err
				// 		}
				__obf_1c35129ed9e54186)
	}
	return
}

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
