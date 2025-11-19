package __obf_9861fa13140c30a3

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_56088f21b87f829d map[string]*Element[V]
	__obf_44c765adef5b5421 __obf_9da43f38a5013de2[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{
		__obf_56088f21b87f829d: make(map[string]*Element[V]),
	}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_49189a7b5399252f *OrderedMap[V]) Get(__obf_43194ec765d86867 string) (__obf_3966a3e46995974f V, __obf_da2280290c7dd975 bool) {
	__obf_478af39a6b826f97, __obf_da2280290c7dd975 := __obf_49189a7b5399252f.__obf_56088f21b87f829d[__obf_43194ec765d86867]
	if __obf_da2280290c7dd975 {
		__obf_3966a3e46995974f = __obf_478af39a6b826f97.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_49189a7b5399252f *OrderedMap[V]) Set(__obf_43194ec765d86867 string, __obf_3966a3e46995974f V) bool {
	_, __obf_6f105a24391c9014 := __obf_49189a7b5399252f.__obf_56088f21b87f829d[__obf_43194ec765d86867]
	if __obf_6f105a24391c9014 {
		__obf_49189a7b5399252f.__obf_56088f21b87f829d[__obf_43194ec765d86867].Value = __obf_3966a3e46995974f
		return false
	}

	__obf_c4037c31e5ebe4ad := __obf_49189a7b5399252f.__obf_44c765adef5b5421.PushBack(__obf_43194ec765d86867, __obf_3966a3e46995974f)
	__obf_49189a7b5399252f.__obf_56088f21b87f829d[__obf_43194ec765d86867] = __obf_c4037c31e5ebe4ad
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_49189a7b5399252f *OrderedMap[V]) GetOrDefault(__obf_43194ec765d86867 string, __obf_f3db3694538d7a1d V) V {
	if __obf_3966a3e46995974f, __obf_da2280290c7dd975 := __obf_49189a7b5399252f.__obf_56088f21b87f829d[__obf_43194ec765d86867]; __obf_da2280290c7dd975 {
		return __obf_3966a3e46995974f.Value
	}

	return __obf_f3db3694538d7a1d
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_49189a7b5399252f *OrderedMap[V]) GetElement(__obf_43194ec765d86867 string) *Element[V] {
	__obf_c4037c31e5ebe4ad, __obf_da2280290c7dd975 := __obf_49189a7b5399252f.__obf_56088f21b87f829d[__obf_43194ec765d86867]
	if __obf_da2280290c7dd975 {
		return __obf_c4037c31e5ebe4ad
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_49189a7b5399252f *OrderedMap[V]) Len() int {
	return len(__obf_49189a7b5399252f.__obf_56088f21b87f829d)
}

// Keys returns all of the keys in the order they were inserted. If a key was
// replaced it will retain the same position. To ensure most recently set keys
// are always at the end you must always Delete before Set.
func (__obf_49189a7b5399252f *OrderedMap[V]) Keys() (__obf_6c59a1f64b802da6 []string) {
	__obf_6c59a1f64b802da6 = make([]string, 0, __obf_49189a7b5399252f.Len())
	for __obf_654277b99a28ad3c := __obf_49189a7b5399252f.Front(); __obf_654277b99a28ad3c != nil; __obf_654277b99a28ad3c = __obf_654277b99a28ad3c.Next() {
		__obf_6c59a1f64b802da6 = append(__obf_6c59a1f64b802da6, __obf_654277b99a28ad3c.Key)
	}
	return __obf_6c59a1f64b802da6
}

// Values returns all of the values in the order they were inserted.
func (__obf_49189a7b5399252f *OrderedMap[V]) Values() (__obf_010e40a589eedb1e []V) {
	__obf_010e40a589eedb1e = make([]V, 0, __obf_49189a7b5399252f.Len())
	for __obf_654277b99a28ad3c := __obf_49189a7b5399252f.Front(); __obf_654277b99a28ad3c != nil; __obf_654277b99a28ad3c = __obf_654277b99a28ad3c.Next() {
		__obf_010e40a589eedb1e = append(__obf_010e40a589eedb1e, __obf_654277b99a28ad3c.Value)
	}
	return __obf_010e40a589eedb1e
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
func (__obf_49189a7b5399252f *OrderedMap[V]) Delete(__obf_43194ec765d86867 string) (__obf_17313b8a32b145d9 bool) {
	__obf_c4037c31e5ebe4ad, __obf_da2280290c7dd975 := __obf_49189a7b5399252f.__obf_56088f21b87f829d[__obf_43194ec765d86867]
	if __obf_da2280290c7dd975 {
		__obf_49189a7b5399252f.__obf_44c765adef5b5421.Remove(__obf_c4037c31e5ebe4ad)
		delete(__obf_49189a7b5399252f.__obf_56088f21b87f829d, __obf_43194ec765d86867)
	}

	return __obf_da2280290c7dd975
}

// Front will return the element that is the first (oldest Set element). If
// there are no elements this will return nil.
func (__obf_49189a7b5399252f *OrderedMap[V]) Front() *Element[V] {
	return __obf_49189a7b5399252f.__obf_44c765adef5b5421.Front()
}

// Back will return the element that is the last (most recent Set element). If
// there are no elements this will return nil.
func (__obf_49189a7b5399252f *OrderedMap[V]) Back() *Element[V] {
	return __obf_49189a7b5399252f.__obf_44c765adef5b5421.Back()
}

// Copy returns a new OrderedMap with the same elements.
// Using Copy while there are concurrent writes may mangle the result.
func (__obf_49189a7b5399252f *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_38e7e492480d933e := NewOrderedMap[V]()
	for __obf_654277b99a28ad3c := __obf_49189a7b5399252f.Front(); __obf_654277b99a28ad3c != nil; __obf_654277b99a28ad3c = __obf_654277b99a28ad3c.Next() {
		__obf_38e7e492480d933e.Set(__obf_654277b99a28ad3c.Key, __obf_654277b99a28ad3c.Value)
	}
	return __obf_38e7e492480d933e
}

// SetFromJson set element from json string
func (__obf_49189a7b5399252f *OrderedMap[V]) SetFromJson(__obf_db9e15114372a098 string) (__obf_d372e4dd7b4813ca error) {
	__obf_6993879924ef45ba := make(map[string]V, 0)
	__obf_d372e4dd7b4813ca = json.Unmarshal([]byte(__obf_db9e15114372a098), &__obf_6993879924ef45ba)
	if __obf_d372e4dd7b4813ca != nil {
		return
	}
	for __obf_299b9ac8f6041f1b, __obf_478af39a6b826f97 := range __obf_6993879924ef45ba {
		__obf_49189a7b5399252f.Set(__obf_299b9ac8f6041f1b, __obf_478af39a6b826f97)
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
