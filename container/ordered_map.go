package __obf_b0bebe5eb45b8ad6

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_df1d37ee00c6f23a map[string]*Element[V]
	__obf_2e450f069f424dfb __obf_04c038c56fdc3d0a[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{__obf_df1d37ee00c6f23a: make(map[string]*Element[V])}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) Get(__obf_9c15798bcb95be3e string) (__obf_72f1369f2f75ff87 V, __obf_e8c1fb9f7287beef bool) {
	__obf_10522304adc21daf, __obf_e8c1fb9f7287beef := __obf_0f92cd4dcce1b272.__obf_df1d37ee00c6f23a[__obf_9c15798bcb95be3e]
	if __obf_e8c1fb9f7287beef {
		__obf_72f1369f2f75ff87 = __obf_10522304adc21daf.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) Set(__obf_9c15798bcb95be3e string, __obf_72f1369f2f75ff87 V) bool {
	_, __obf_664af2bbc77aa109 := __obf_0f92cd4dcce1b272.__obf_df1d37ee00c6f23a[__obf_9c15798bcb95be3e]
	if __obf_664af2bbc77aa109 {
		__obf_0f92cd4dcce1b272.__obf_df1d37ee00c6f23a[__obf_9c15798bcb95be3e].Value = __obf_72f1369f2f75ff87
		return false
	}
	__obf_4128aedf8b3867f4 := __obf_0f92cd4dcce1b272.__obf_2e450f069f424dfb.PushBack(__obf_9c15798bcb95be3e, __obf_72f1369f2f75ff87)
	__obf_0f92cd4dcce1b272.__obf_df1d37ee00c6f23a[__obf_9c15798bcb95be3e] = __obf_4128aedf8b3867f4
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) GetOrDefault(__obf_9c15798bcb95be3e string, __obf_b90751624c5e511f V) V {
	if __obf_72f1369f2f75ff87, __obf_e8c1fb9f7287beef := __obf_0f92cd4dcce1b272.__obf_df1d37ee00c6f23a[__obf_9c15798bcb95be3e]; __obf_e8c1fb9f7287beef {
		return __obf_72f1369f2f75ff87.Value
	}

	return __obf_b90751624c5e511f
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) GetElement(__obf_9c15798bcb95be3e string) *Element[V] {
	__obf_4128aedf8b3867f4, __obf_e8c1fb9f7287beef := __obf_0f92cd4dcce1b272.__obf_df1d37ee00c6f23a[__obf_9c15798bcb95be3e]
	if __obf_e8c1fb9f7287beef {
		return __obf_4128aedf8b3867f4
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) Len() int {
	return len(__obf_0f92cd4dcce1b272.

		// Keys returns all of the keys in the order they were inserted. If a key was
		// replaced it will retain the same position. To ensure most recently set keys
		// are always at the end you must always Delete before Set.
		__obf_df1d37ee00c6f23a)
}

func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) Keys() (__obf_1e824d5b9fe4a6dc []string) {
	__obf_1e824d5b9fe4a6dc = make([]string, 0, __obf_0f92cd4dcce1b272.Len())
	for __obf_790870c9c9e4a7c4 := __obf_0f92cd4dcce1b272.Front(); __obf_790870c9c9e4a7c4 != nil; __obf_790870c9c9e4a7c4 = __obf_790870c9c9e4a7c4.Next() {
		__obf_1e824d5b9fe4a6dc = append(__obf_1e824d5b9fe4a6dc, __obf_790870c9c9e4a7c4.Key)
	}
	return __obf_1e824d5b9fe4a6dc
}

// Values returns all of the values in the order they were inserted.
func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) Values() (__obf_d054d4311affb29a []V) {
	__obf_d054d4311affb29a = make([]V, 0, __obf_0f92cd4dcce1b272.Len())
	for __obf_790870c9c9e4a7c4 := __obf_0f92cd4dcce1b272.Front(); __obf_790870c9c9e4a7c4 != nil; __obf_790870c9c9e4a7c4 = __obf_790870c9c9e4a7c4.Next() {
		__obf_d054d4311affb29a = append(__obf_d054d4311affb29a, __obf_790870c9c9e4a7c4.Value)
	}
	return __obf_d054d4311affb29a
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
func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) Delete(__obf_9c15798bcb95be3e string) (__obf_6492b88d6c6aa22e bool) {
	__obf_4128aedf8b3867f4, __obf_e8c1fb9f7287beef := __obf_0f92cd4dcce1b272.__obf_df1d37ee00c6f23a[__obf_9c15798bcb95be3e]
	if __obf_e8c1fb9f7287beef {
		__obf_0f92cd4dcce1b272.__obf_2e450f069f424dfb.
			Remove(__obf_4128aedf8b3867f4)
		delete(__obf_0f92cd4dcce1b272.__obf_df1d37ee00c6f23a,

			// Front will return the element that is the first (oldest Set element). If
			// there are no elements this will return nil.
			__obf_9c15798bcb95be3e)
	}

	return __obf_e8c1fb9f7287beef
}

func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) Front() *Element[V] {
	return __obf_0f92cd4dcce1b272.

		// Back will return the element that is the last (most recent Set element). If
		// there are no elements this will return nil.
		__obf_2e450f069f424dfb.Front()
}

func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) Back() *Element[V] {
	return __obf_0f92cd4dcce1b272.

		// Copy returns a new OrderedMap with the same elements.
		// Using Copy while there are concurrent writes may mangle the result.
		__obf_2e450f069f424dfb.Back()
}

func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_83e90d803a06d38f := NewOrderedMap[V]()
	for __obf_790870c9c9e4a7c4 := __obf_0f92cd4dcce1b272.Front(); __obf_790870c9c9e4a7c4 != nil; __obf_790870c9c9e4a7c4 = __obf_790870c9c9e4a7c4.Next() {
		__obf_83e90d803a06d38f.
			Set(__obf_790870c9c9e4a7c4.Key, __obf_790870c9c9e4a7c4.Value)
	}
	return __obf_83e90d803a06d38f
}

// SetFromJson set element from json string
func (__obf_0f92cd4dcce1b272 *OrderedMap[V]) SetFromJson(__obf_63b8421601ca2033 string) (__obf_805eb0b5dd52f7ba error) {
	__obf_c2fed62bb29fed2b := make(map[string]V, 0)
	__obf_805eb0b5dd52f7ba = json.Unmarshal([]byte(__obf_63b8421601ca2033), &__obf_c2fed62bb29fed2b)
	if __obf_805eb0b5dd52f7ba != nil {
		return
	}
	for __obf_dca1c5db9aad12e8, __obf_10522304adc21daf := range __obf_c2fed62bb29fed2b {
		__obf_0f92cd4dcce1b272.
			Set(__obf_dca1c5db9aad12e8,

				// func (m *OrderedMap[V]) parseobject(dec *json.Decoder) (err error) {
				// 	var t json.Token
				// 	for dec.More() {
				// 		t, err = dec.Token()
				// 		if err != nil {
				// 			return err
				// 		}
				__obf_10522304adc21daf)
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
