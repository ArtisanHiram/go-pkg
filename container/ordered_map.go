package __obf_4a16ef421fb74992

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_264c7b83919c0aa6 map[string]*Element[V]
	__obf_7a1af2aea5f3c8de __obf_333361aa59ab2d50[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{__obf_264c7b83919c0aa6: make(map[string]*Element[V])}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) Get(__obf_d1a1f98ae0fb119e string) (__obf_06ad079a003b5526 V, __obf_0a7b144ecfd706f0 bool) {
	__obf_b3b7dced15e3b450, __obf_0a7b144ecfd706f0 := __obf_f830c75ef9a29ed2.__obf_264c7b83919c0aa6[__obf_d1a1f98ae0fb119e]
	if __obf_0a7b144ecfd706f0 {
		__obf_06ad079a003b5526 = __obf_b3b7dced15e3b450.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) Set(__obf_d1a1f98ae0fb119e string, __obf_06ad079a003b5526 V) bool {
	_, __obf_f764d2809139bc47 := __obf_f830c75ef9a29ed2.__obf_264c7b83919c0aa6[__obf_d1a1f98ae0fb119e]
	if __obf_f764d2809139bc47 {
		__obf_f830c75ef9a29ed2.__obf_264c7b83919c0aa6[__obf_d1a1f98ae0fb119e].Value = __obf_06ad079a003b5526
		return false
	}
	__obf_33ece9b6688716f4 := __obf_f830c75ef9a29ed2.__obf_7a1af2aea5f3c8de.PushBack(__obf_d1a1f98ae0fb119e, __obf_06ad079a003b5526)
	__obf_f830c75ef9a29ed2.__obf_264c7b83919c0aa6[__obf_d1a1f98ae0fb119e] = __obf_33ece9b6688716f4
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) GetOrDefault(__obf_d1a1f98ae0fb119e string, __obf_ac8457b9a9a9dd1b V) V {
	if __obf_06ad079a003b5526, __obf_0a7b144ecfd706f0 := __obf_f830c75ef9a29ed2.__obf_264c7b83919c0aa6[__obf_d1a1f98ae0fb119e]; __obf_0a7b144ecfd706f0 {
		return __obf_06ad079a003b5526.Value
	}

	return __obf_ac8457b9a9a9dd1b
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) GetElement(__obf_d1a1f98ae0fb119e string) *Element[V] {
	__obf_33ece9b6688716f4, __obf_0a7b144ecfd706f0 := __obf_f830c75ef9a29ed2.__obf_264c7b83919c0aa6[__obf_d1a1f98ae0fb119e]
	if __obf_0a7b144ecfd706f0 {
		return __obf_33ece9b6688716f4
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) Len() int {
	return len(__obf_f830c75ef9a29ed2.

		// Keys returns all of the keys in the order they were inserted. If a key was
		// replaced it will retain the same position. To ensure most recently set keys
		// are always at the end you must always Delete before Set.
		__obf_264c7b83919c0aa6)
}

func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) Keys() (__obf_e937fcb7c91671b3 []string) {
	__obf_e937fcb7c91671b3 = make([]string, 0, __obf_f830c75ef9a29ed2.Len())
	for __obf_8843fe4a92e1ec68 := __obf_f830c75ef9a29ed2.Front(); __obf_8843fe4a92e1ec68 != nil; __obf_8843fe4a92e1ec68 = __obf_8843fe4a92e1ec68.Next() {
		__obf_e937fcb7c91671b3 = append(__obf_e937fcb7c91671b3, __obf_8843fe4a92e1ec68.Key)
	}
	return __obf_e937fcb7c91671b3
}

// Values returns all of the values in the order they were inserted.
func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) Values() (__obf_4f59e8768782d4bd []V) {
	__obf_4f59e8768782d4bd = make([]V, 0, __obf_f830c75ef9a29ed2.Len())
	for __obf_8843fe4a92e1ec68 := __obf_f830c75ef9a29ed2.Front(); __obf_8843fe4a92e1ec68 != nil; __obf_8843fe4a92e1ec68 = __obf_8843fe4a92e1ec68.Next() {
		__obf_4f59e8768782d4bd = append(__obf_4f59e8768782d4bd, __obf_8843fe4a92e1ec68.Value)
	}
	return __obf_4f59e8768782d4bd
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
func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) Delete(__obf_d1a1f98ae0fb119e string) (__obf_814e0fc558c60017 bool) {
	__obf_33ece9b6688716f4, __obf_0a7b144ecfd706f0 := __obf_f830c75ef9a29ed2.__obf_264c7b83919c0aa6[__obf_d1a1f98ae0fb119e]
	if __obf_0a7b144ecfd706f0 {
		__obf_f830c75ef9a29ed2.__obf_7a1af2aea5f3c8de.
			Remove(__obf_33ece9b6688716f4)
		delete(__obf_f830c75ef9a29ed2.__obf_264c7b83919c0aa6,

			// Front will return the element that is the first (oldest Set element). If
			// there are no elements this will return nil.
			__obf_d1a1f98ae0fb119e)
	}

	return __obf_0a7b144ecfd706f0
}

func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) Front() *Element[V] {
	return __obf_f830c75ef9a29ed2.

		// Back will return the element that is the last (most recent Set element). If
		// there are no elements this will return nil.
		__obf_7a1af2aea5f3c8de.Front()
}

func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) Back() *Element[V] {
	return __obf_f830c75ef9a29ed2.

		// Copy returns a new OrderedMap with the same elements.
		// Using Copy while there are concurrent writes may mangle the result.
		__obf_7a1af2aea5f3c8de.Back()
}

func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_320c757560a0b907 := NewOrderedMap[V]()
	for __obf_8843fe4a92e1ec68 := __obf_f830c75ef9a29ed2.Front(); __obf_8843fe4a92e1ec68 != nil; __obf_8843fe4a92e1ec68 = __obf_8843fe4a92e1ec68.Next() {
		__obf_320c757560a0b907.
			Set(__obf_8843fe4a92e1ec68.Key, __obf_8843fe4a92e1ec68.Value)
	}
	return __obf_320c757560a0b907
}

// SetFromJson set element from json string
func (__obf_f830c75ef9a29ed2 *OrderedMap[V]) SetFromJson(__obf_0f4fa23fdfbce2d0 string) (__obf_fdc252da0d475ab8 error) {
	__obf_5af817edd5b5b2a1 := make(map[string]V, 0)
	__obf_fdc252da0d475ab8 = json.Unmarshal([]byte(__obf_0f4fa23fdfbce2d0), &__obf_5af817edd5b5b2a1)
	if __obf_fdc252da0d475ab8 != nil {
		return
	}
	for __obf_91fb3c93595a08ac, __obf_b3b7dced15e3b450 := range __obf_5af817edd5b5b2a1 {
		__obf_f830c75ef9a29ed2.
			Set(__obf_91fb3c93595a08ac,

				// func (m *OrderedMap[V]) parseobject(dec *json.Decoder) (err error) {
				// 	var t json.Token
				// 	for dec.More() {
				// 		t, err = dec.Token()
				// 		if err != nil {
				// 			return err
				// 		}
				__obf_b3b7dced15e3b450)
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
