package __obf_038560a94647875f

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_29dc09b5c146beea map[string]*Element[V]
	__obf_ff381ded265d4673 __obf_295482420737c0b5[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{__obf_29dc09b5c146beea: make(map[string]*Element[V])}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_e933cbaa93bd62dd *OrderedMap[V]) Get(__obf_52e73bb48e810dd2 string) (__obf_3270ab9e70995438 V, __obf_5904187971b4cb86 bool) {
	__obf_5ec12fe59cf82410, __obf_5904187971b4cb86 := __obf_e933cbaa93bd62dd.__obf_29dc09b5c146beea[__obf_52e73bb48e810dd2]
	if __obf_5904187971b4cb86 {
		__obf_3270ab9e70995438 = __obf_5ec12fe59cf82410.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_e933cbaa93bd62dd *OrderedMap[V]) Set(__obf_52e73bb48e810dd2 string, __obf_3270ab9e70995438 V) bool {
	_, __obf_1f25e5130c0bcc82 := __obf_e933cbaa93bd62dd.__obf_29dc09b5c146beea[__obf_52e73bb48e810dd2]
	if __obf_1f25e5130c0bcc82 {
		__obf_e933cbaa93bd62dd.__obf_29dc09b5c146beea[__obf_52e73bb48e810dd2].Value = __obf_3270ab9e70995438
		return false
	}
	__obf_05776320602743d5 := __obf_e933cbaa93bd62dd.__obf_ff381ded265d4673.PushBack(__obf_52e73bb48e810dd2, __obf_3270ab9e70995438)
	__obf_e933cbaa93bd62dd.__obf_29dc09b5c146beea[__obf_52e73bb48e810dd2] = __obf_05776320602743d5
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_e933cbaa93bd62dd *OrderedMap[V]) GetOrDefault(__obf_52e73bb48e810dd2 string, __obf_e880daed8c1fdd92 V) V {
	if __obf_3270ab9e70995438, __obf_5904187971b4cb86 := __obf_e933cbaa93bd62dd.__obf_29dc09b5c146beea[__obf_52e73bb48e810dd2]; __obf_5904187971b4cb86 {
		return __obf_3270ab9e70995438.Value
	}

	return __obf_e880daed8c1fdd92
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_e933cbaa93bd62dd *OrderedMap[V]) GetElement(__obf_52e73bb48e810dd2 string) *Element[V] {
	__obf_05776320602743d5, __obf_5904187971b4cb86 := __obf_e933cbaa93bd62dd.__obf_29dc09b5c146beea[__obf_52e73bb48e810dd2]
	if __obf_5904187971b4cb86 {
		return __obf_05776320602743d5
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_e933cbaa93bd62dd *OrderedMap[V]) Len() int {
	return len(__obf_e933cbaa93bd62dd.

		// Keys returns all of the keys in the order they were inserted. If a key was
		// replaced it will retain the same position. To ensure most recently set keys
		// are always at the end you must always Delete before Set.
		__obf_29dc09b5c146beea)
}

func (__obf_e933cbaa93bd62dd *OrderedMap[V]) Keys() (__obf_0414b3afaf21f9fa []string) {
	__obf_0414b3afaf21f9fa = make([]string, 0, __obf_e933cbaa93bd62dd.Len())
	for __obf_ed6ca3e0dc7f65ea := __obf_e933cbaa93bd62dd.Front(); __obf_ed6ca3e0dc7f65ea != nil; __obf_ed6ca3e0dc7f65ea = __obf_ed6ca3e0dc7f65ea.Next() {
		__obf_0414b3afaf21f9fa = append(__obf_0414b3afaf21f9fa, __obf_ed6ca3e0dc7f65ea.Key)
	}
	return __obf_0414b3afaf21f9fa
}

// Values returns all of the values in the order they were inserted.
func (__obf_e933cbaa93bd62dd *OrderedMap[V]) Values() (__obf_fcda9cff02fa0ecd []V) {
	__obf_fcda9cff02fa0ecd = make([]V, 0, __obf_e933cbaa93bd62dd.Len())
	for __obf_ed6ca3e0dc7f65ea := __obf_e933cbaa93bd62dd.Front(); __obf_ed6ca3e0dc7f65ea != nil; __obf_ed6ca3e0dc7f65ea = __obf_ed6ca3e0dc7f65ea.Next() {
		__obf_fcda9cff02fa0ecd = append(__obf_fcda9cff02fa0ecd, __obf_ed6ca3e0dc7f65ea.Value)
	}
	return __obf_fcda9cff02fa0ecd
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
func (__obf_e933cbaa93bd62dd *OrderedMap[V]) Delete(__obf_52e73bb48e810dd2 string) (__obf_a0adfeaff3b425b9 bool) {
	__obf_05776320602743d5, __obf_5904187971b4cb86 := __obf_e933cbaa93bd62dd.__obf_29dc09b5c146beea[__obf_52e73bb48e810dd2]
	if __obf_5904187971b4cb86 {
		__obf_e933cbaa93bd62dd.__obf_ff381ded265d4673.
			Remove(__obf_05776320602743d5)
		delete(__obf_e933cbaa93bd62dd.__obf_29dc09b5c146beea,

			// Front will return the element that is the first (oldest Set element). If
			// there are no elements this will return nil.
			__obf_52e73bb48e810dd2)
	}

	return __obf_5904187971b4cb86
}

func (__obf_e933cbaa93bd62dd *OrderedMap[V]) Front() *Element[V] {
	return __obf_e933cbaa93bd62dd.

		// Back will return the element that is the last (most recent Set element). If
		// there are no elements this will return nil.
		__obf_ff381ded265d4673.Front()
}

func (__obf_e933cbaa93bd62dd *OrderedMap[V]) Back() *Element[V] {
	return __obf_e933cbaa93bd62dd.

		// Copy returns a new OrderedMap with the same elements.
		// Using Copy while there are concurrent writes may mangle the result.
		__obf_ff381ded265d4673.Back()
}

func (__obf_e933cbaa93bd62dd *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_4509c54e4c75be7f := NewOrderedMap[V]()
	for __obf_ed6ca3e0dc7f65ea := __obf_e933cbaa93bd62dd.Front(); __obf_ed6ca3e0dc7f65ea != nil; __obf_ed6ca3e0dc7f65ea = __obf_ed6ca3e0dc7f65ea.Next() {
		__obf_4509c54e4c75be7f.
			Set(__obf_ed6ca3e0dc7f65ea.Key, __obf_ed6ca3e0dc7f65ea.Value)
	}
	return __obf_4509c54e4c75be7f
}

// SetFromJson set element from json string
func (__obf_e933cbaa93bd62dd *OrderedMap[V]) SetFromJson(__obf_94c99105455b1b61 string) (__obf_ad71453048e280e8 error) {
	__obf_60b27a69fbe91252 := make(map[string]V, 0)
	__obf_ad71453048e280e8 = json.Unmarshal([]byte(__obf_94c99105455b1b61), &__obf_60b27a69fbe91252)
	if __obf_ad71453048e280e8 != nil {
		return
	}
	for __obf_ed002d86aefb8e41, __obf_5ec12fe59cf82410 := range __obf_60b27a69fbe91252 {
		__obf_e933cbaa93bd62dd.
			Set(__obf_ed002d86aefb8e41,

				// func (m *OrderedMap[V]) parseobject(dec *json.Decoder) (err error) {
				// 	var t json.Token
				// 	for dec.More() {
				// 		t, err = dec.Token()
				// 		if err != nil {
				// 			return err
				// 		}
				__obf_5ec12fe59cf82410)
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
