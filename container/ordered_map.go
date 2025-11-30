package __obf_e72ce603d10d02a1

import (
	"encoding/json"
)

type OrderedMap[V any] struct {
	__obf_521026317ed93586 map[string]*Element[V]
	__obf_a85613776068d346 __obf_e84590c269d065c9[V]
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{__obf_521026317ed93586: make(map[string]*Element[V])}
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (__obf_a2cdbb539e506980 *OrderedMap[V]) Get(__obf_9dd08479fe906662 string) (__obf_75ac60bfa194af6b V, __obf_b95641e1ff7de824 bool) {
	__obf_ebe2552a06e91009, __obf_b95641e1ff7de824 := __obf_a2cdbb539e506980.__obf_521026317ed93586[__obf_9dd08479fe906662]
	if __obf_b95641e1ff7de824 {
		__obf_75ac60bfa194af6b = __obf_ebe2552a06e91009.Value
	}

	return
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (__obf_a2cdbb539e506980 *OrderedMap[V]) Set(__obf_9dd08479fe906662 string, __obf_75ac60bfa194af6b V) bool {
	_, __obf_6dd14ab4119424f9 := __obf_a2cdbb539e506980.__obf_521026317ed93586[__obf_9dd08479fe906662]
	if __obf_6dd14ab4119424f9 {
		__obf_a2cdbb539e506980.__obf_521026317ed93586[__obf_9dd08479fe906662].Value = __obf_75ac60bfa194af6b
		return false
	}
	__obf_dd6147f749207af9 := __obf_a2cdbb539e506980.__obf_a85613776068d346.PushBack(__obf_9dd08479fe906662, __obf_75ac60bfa194af6b)
	__obf_a2cdbb539e506980.__obf_521026317ed93586[__obf_9dd08479fe906662] = __obf_dd6147f749207af9
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (__obf_a2cdbb539e506980 *OrderedMap[V]) GetOrDefault(__obf_9dd08479fe906662 string, __obf_469aa405ef5d0008 V) V {
	if __obf_75ac60bfa194af6b, __obf_b95641e1ff7de824 := __obf_a2cdbb539e506980.__obf_521026317ed93586[__obf_9dd08479fe906662]; __obf_b95641e1ff7de824 {
		return __obf_75ac60bfa194af6b.Value
	}

	return __obf_469aa405ef5d0008
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (__obf_a2cdbb539e506980 *OrderedMap[V]) GetElement(__obf_9dd08479fe906662 string) *Element[V] {
	__obf_dd6147f749207af9, __obf_b95641e1ff7de824 := __obf_a2cdbb539e506980.__obf_521026317ed93586[__obf_9dd08479fe906662]
	if __obf_b95641e1ff7de824 {
		return __obf_dd6147f749207af9
	}

	return nil
}

// Len returns the number of elements in the map.
func (__obf_a2cdbb539e506980 *OrderedMap[V]) Len() int {
	return len(__obf_a2cdbb539e506980.

		// Keys returns all of the keys in the order they were inserted. If a key was
		// replaced it will retain the same position. To ensure most recently set keys
		// are always at the end you must always Delete before Set.
		__obf_521026317ed93586)
}

func (__obf_a2cdbb539e506980 *OrderedMap[V]) Keys() (__obf_e4afa361e4970b8e []string) {
	__obf_e4afa361e4970b8e = make([]string, 0, __obf_a2cdbb539e506980.Len())
	for __obf_ff2030bc68a5c959 := __obf_a2cdbb539e506980.Front(); __obf_ff2030bc68a5c959 != nil; __obf_ff2030bc68a5c959 = __obf_ff2030bc68a5c959.Next() {
		__obf_e4afa361e4970b8e = append(__obf_e4afa361e4970b8e, __obf_ff2030bc68a5c959.Key)
	}
	return __obf_e4afa361e4970b8e
}

// Values returns all of the values in the order they were inserted.
func (__obf_a2cdbb539e506980 *OrderedMap[V]) Values() (__obf_7395e77052e9e60b []V) {
	__obf_7395e77052e9e60b = make([]V, 0, __obf_a2cdbb539e506980.Len())
	for __obf_ff2030bc68a5c959 := __obf_a2cdbb539e506980.Front(); __obf_ff2030bc68a5c959 != nil; __obf_ff2030bc68a5c959 = __obf_ff2030bc68a5c959.Next() {
		__obf_7395e77052e9e60b = append(__obf_7395e77052e9e60b, __obf_ff2030bc68a5c959.Value)
	}
	return __obf_7395e77052e9e60b
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
func (__obf_a2cdbb539e506980 *OrderedMap[V]) Delete(__obf_9dd08479fe906662 string) (__obf_6c0e677480126a91 bool) {
	__obf_dd6147f749207af9, __obf_b95641e1ff7de824 := __obf_a2cdbb539e506980.__obf_521026317ed93586[__obf_9dd08479fe906662]
	if __obf_b95641e1ff7de824 {
		__obf_a2cdbb539e506980.__obf_a85613776068d346.
			Remove(__obf_dd6147f749207af9)
		delete(__obf_a2cdbb539e506980.__obf_521026317ed93586,

			// Front will return the element that is the first (oldest Set element). If
			// there are no elements this will return nil.
			__obf_9dd08479fe906662)
	}

	return __obf_b95641e1ff7de824
}

func (__obf_a2cdbb539e506980 *OrderedMap[V]) Front() *Element[V] {
	return __obf_a2cdbb539e506980.

		// Back will return the element that is the last (most recent Set element). If
		// there are no elements this will return nil.
		__obf_a85613776068d346.Front()
}

func (__obf_a2cdbb539e506980 *OrderedMap[V]) Back() *Element[V] {
	return __obf_a2cdbb539e506980.

		// Copy returns a new OrderedMap with the same elements.
		// Using Copy while there are concurrent writes may mangle the result.
		__obf_a85613776068d346.Back()
}

func (__obf_a2cdbb539e506980 *OrderedMap[V]) Copy() *OrderedMap[V] {
	__obf_545ab43a458a3254 := NewOrderedMap[V]()
	for __obf_ff2030bc68a5c959 := __obf_a2cdbb539e506980.Front(); __obf_ff2030bc68a5c959 != nil; __obf_ff2030bc68a5c959 = __obf_ff2030bc68a5c959.Next() {
		__obf_545ab43a458a3254.
			Set(__obf_ff2030bc68a5c959.Key, __obf_ff2030bc68a5c959.Value)
	}
	return __obf_545ab43a458a3254
}

// SetFromJson set element from json string
func (__obf_a2cdbb539e506980 *OrderedMap[V]) SetFromJson(__obf_8100de325da9ea48 string) (__obf_3674305c6f99212d error) {
	__obf_6bd52d4fff1b8480 := make(map[string]V, 0)
	__obf_3674305c6f99212d = json.Unmarshal([]byte(__obf_8100de325da9ea48), &__obf_6bd52d4fff1b8480)
	if __obf_3674305c6f99212d != nil {
		return
	}
	for __obf_cce6a084cb2eec1f, __obf_ebe2552a06e91009 := range __obf_6bd52d4fff1b8480 {
		__obf_a2cdbb539e506980.
			Set(__obf_cce6a084cb2eec1f,

				// func (m *OrderedMap[V]) parseobject(dec *json.Decoder) (err error) {
				// 	var t json.Token
				// 	for dec.More() {
				// 		t, err = dec.Token()
				// 		if err != nil {
				// 			return err
				// 		}
				__obf_ebe2552a06e91009)
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
