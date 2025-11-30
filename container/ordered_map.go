package __obf_9004b47202f9204b

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_ab7eb424d0a0826e map[string]*Element[V]
	__obf_73ff9ed62f0ef2e4 __obf_72ca0e09d2929881[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{__obf_ab7eb424d0a0826e: make(map[string]*Element[V])}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) Get(__obf_355e7c922e433678 string) (__obf_7269ae284cf88b6b V, __obf_545047dd3a6c3c8d bool) {
	__obf_8616b2f4d0fa5a6a, __obf_545047dd3a6c3c8d := __obf_ff1f15dd12dee7a5.__obf_ab7eb424d0a0826e[__obf_355e7c922e433678]
	if __obf_545047dd3a6c3c8d {
		__obf_7269ae284cf88b6b = __obf_8616b2f4d0fa5a6a.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) Set(__obf_355e7c922e433678 string, __obf_7269ae284cf88b6b V) bool {
	_, __obf_0505051c0ff014ac := __obf_ff1f15dd12dee7a5.__obf_ab7eb424d0a0826e[__obf_355e7c922e433678]
	if __obf_0505051c0ff014ac {
		__obf_ff1f15dd12dee7a5.__obf_ab7eb424d0a0826e[__obf_355e7c922e433678].Value = __obf_7269ae284cf88b6b
		return false
	}
	__obf_b1b4a7d9151cbc43 := __obf_ff1f15dd12dee7a5.__obf_73ff9ed62f0ef2e4.PushBack(__obf_355e7c922e433678, __obf_7269ae284cf88b6b)
	__obf_ff1f15dd12dee7a5.__obf_ab7eb424d0a0826e[__obf_355e7c922e433678] = __obf_b1b4a7d9151cbc43
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) GetOrDefault(__obf_355e7c922e433678 string, __obf_f3cae0e75eebbfe6 V) V {
	if __obf_7269ae284cf88b6b, __obf_545047dd3a6c3c8d := __obf_ff1f15dd12dee7a5.__obf_ab7eb424d0a0826e[__obf_355e7c922e433678]; __obf_545047dd3a6c3c8d {
		return __obf_7269ae284cf88b6b.Value
	}

	return __obf_f3cae0e75eebbfe6
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) GetElement(__obf_355e7c922e433678 string) *Element[V] {
	__obf_b1b4a7d9151cbc43, __obf_545047dd3a6c3c8d := __obf_ff1f15dd12dee7a5.__obf_ab7eb424d0a0826e[__obf_355e7c922e433678]
	if __obf_545047dd3a6c3c8d {
		return __obf_b1b4a7d9151cbc43
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) Len() int {
	return len(__obf_ff1f15dd12dee7a5.

		// Keys returns all of the keys in the order they were inserted. If a key was
		// replaced it will retain the same position. To ensure most recently set keys
		// are always at the end you must always Delete before Set.
		__obf_ab7eb424d0a0826e)
}

func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) Keys() (__obf_6bb2e295ddff69c0 []string) {
	__obf_6bb2e295ddff69c0 = make([]string, 0, __obf_ff1f15dd12dee7a5.Len())
	for __obf_6ec5d1c972f42683 := __obf_ff1f15dd12dee7a5.Front(); __obf_6ec5d1c972f42683 != nil; __obf_6ec5d1c972f42683 = __obf_6ec5d1c972f42683.Next() {
		__obf_6bb2e295ddff69c0 = append(__obf_6bb2e295ddff69c0, __obf_6ec5d1c972f42683.Key)
	}
	return __obf_6bb2e295ddff69c0
}

// Values returns all of the values in the order they were inserted.
func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) Values() (__obf_489f8053148d8299 []V) {
	__obf_489f8053148d8299 = make([]V, 0, __obf_ff1f15dd12dee7a5.Len())
	for __obf_6ec5d1c972f42683 := __obf_ff1f15dd12dee7a5.Front(); __obf_6ec5d1c972f42683 != nil; __obf_6ec5d1c972f42683 = __obf_6ec5d1c972f42683.Next() {
		__obf_489f8053148d8299 = append(__obf_489f8053148d8299, __obf_6ec5d1c972f42683.Value)
	}
	return __obf_489f8053148d8299
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
func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) Delete(__obf_355e7c922e433678 string) (__obf_11bbfcb0242f92ea bool) {
	__obf_b1b4a7d9151cbc43, __obf_545047dd3a6c3c8d := __obf_ff1f15dd12dee7a5.__obf_ab7eb424d0a0826e[__obf_355e7c922e433678]
	if __obf_545047dd3a6c3c8d {
		__obf_ff1f15dd12dee7a5.__obf_73ff9ed62f0ef2e4.
			Remove(__obf_b1b4a7d9151cbc43)
		delete(__obf_ff1f15dd12dee7a5.__obf_ab7eb424d0a0826e,

			// Front will return the element that is the first (oldest Set element). If
			// there are no elements this will return nil.
			__obf_355e7c922e433678)
	}

	return __obf_545047dd3a6c3c8d
}

func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) Front() *Element[V] {
	return __obf_ff1f15dd12dee7a5.

		// Back will return the element that is the last (most recent Set element). If
		// there are no elements this will return nil.
		__obf_73ff9ed62f0ef2e4.Front()
}

func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) Back() *Element[V] {
	return __obf_ff1f15dd12dee7a5.

		// Copy returns a new OrderedMap with the same elements.
		// Using Copy while there are concurrent writes may mangle the result.
		__obf_73ff9ed62f0ef2e4.Back()
}

func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_d210d82acdb2977a := NewOrderedMap[V]()
	for __obf_6ec5d1c972f42683 := __obf_ff1f15dd12dee7a5.Front(); __obf_6ec5d1c972f42683 != nil; __obf_6ec5d1c972f42683 = __obf_6ec5d1c972f42683.Next() {
		__obf_d210d82acdb2977a.
			Set(__obf_6ec5d1c972f42683.Key, __obf_6ec5d1c972f42683.Value)
	}
	return __obf_d210d82acdb2977a
}

// SetFromJson set element from json string
func (__obf_ff1f15dd12dee7a5 *OrderedMap[V]) SetFromJson(__obf_a2fde4b7c6372925 string) (__obf_4a0f6bb7d9302e64 error) {
	__obf_8f6731dd48ca32ec := make(map[string]V, 0)
	__obf_4a0f6bb7d9302e64 = json.Unmarshal([]byte(__obf_a2fde4b7c6372925), &__obf_8f6731dd48ca32ec)
	if __obf_4a0f6bb7d9302e64 != nil {
		return
	}
	for __obf_cc18ca82ac13013e, __obf_8616b2f4d0fa5a6a := range __obf_8f6731dd48ca32ec {
		__obf_ff1f15dd12dee7a5.
			Set(__obf_cc18ca82ac13013e,

				// func (m *OrderedMap[V]) parseobject(dec *json.Decoder) (err error) {
				// 	var t json.Token
				// 	for dec.More() {
				// 		t, err = dec.Token()
				// 		if err != nil {
				// 			return err
				// 		}
				__obf_8616b2f4d0fa5a6a)
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
