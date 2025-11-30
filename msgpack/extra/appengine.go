package __obf_7e36788fbb5a0b79

import (
	msgpack "github.com/ArtisanHiram/go-pkg/msgpack"
	ds "google.golang.org/appengine/datastore"
	"reflect"
)

func init() {
	msgpack.Register((*ds.Key)(nil), __obf_129f5532ad76ad77, __obf_075569303cc24bf1)
	msgpack.Register((*ds.Cursor)(nil), __obf_148512d35a5d3dd6, __obf_947db997759bade7)
}

func EncodeDatastoreKey(__obf_95675f385e1b5ed2 *msgpack.Encoder, __obf_f1b95a7f649e8810 *ds.Key) error {
	if __obf_f1b95a7f649e8810 == nil {
		return __obf_95675f385e1b5ed2.EncodeNil()
	}
	return __obf_95675f385e1b5ed2.EncodeString(__obf_f1b95a7f649e8810.Encode())
}

func __obf_129f5532ad76ad77(__obf_95675f385e1b5ed2 *msgpack.Encoder, __obf_7a99ed6469499f7b reflect.Value) error {
	__obf_f1b95a7f649e8810 := __obf_7a99ed6469499f7b.Interface().(*ds.Key)
	return EncodeDatastoreKey(__obf_95675f385e1b5ed2, __obf_f1b95a7f649e8810)
}

func DecodeDatastoreKey(__obf_f1968e9edc9b5015 *msgpack.Decoder) (*ds.Key, error) {
	__obf_7a99ed6469499f7b, __obf_4452d599739d773a := __obf_f1968e9edc9b5015.DecodeString()
	if __obf_4452d599739d773a != nil {
		return nil, __obf_4452d599739d773a
	}
	if __obf_7a99ed6469499f7b == "" {
		return nil, nil
	}
	return ds.DecodeKey(__obf_7a99ed6469499f7b)
}

func __obf_075569303cc24bf1(__obf_f1968e9edc9b5015 *msgpack.Decoder, __obf_7a99ed6469499f7b reflect.Value) error {
	__obf_f1b95a7f649e8810, __obf_4452d599739d773a := DecodeDatastoreKey(__obf_f1968e9edc9b5015)
	if __obf_4452d599739d773a != nil {
		return __obf_4452d599739d773a
	}
	__obf_7a99ed6469499f7b.
		Set(reflect.ValueOf(__obf_f1b95a7f649e8810))
	return nil
}

func __obf_148512d35a5d3dd6(__obf_95675f385e1b5ed2 *msgpack.Encoder, __obf_7a99ed6469499f7b reflect.Value) error {
	__obf_34959fafc7e4dd6a := __obf_7a99ed6469499f7b.Interface().(ds.Cursor)
	return __obf_95675f385e1b5ed2.Encode(__obf_34959fafc7e4dd6a.String())
}

func __obf_947db997759bade7(__obf_f1968e9edc9b5015 *msgpack.Decoder, __obf_7a99ed6469499f7b reflect.Value) error {
	__obf_21ab7f8853f2d79d, __obf_4452d599739d773a := __obf_f1968e9edc9b5015.DecodeString()
	if __obf_4452d599739d773a != nil {
		return __obf_4452d599739d773a
	}
	__obf_34959fafc7e4dd6a, __obf_4452d599739d773a := ds.DecodeCursor(__obf_21ab7f8853f2d79d)
	if __obf_4452d599739d773a != nil {
		return __obf_4452d599739d773a
	}
	__obf_7a99ed6469499f7b.
		Set(reflect.ValueOf(__obf_34959fafc7e4dd6a))
	return nil
}
