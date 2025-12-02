package __obf_36f06d078c5cf16f

import (
	msgpack "github.com/ArtisanHiram/go-pkg/msgpack"
	ds "google.golang.org/appengine/datastore"
	"reflect"
)

func init() {
	msgpack.Register((*ds.Key)(nil), __obf_c42aaa45510b7cdc, __obf_87a05698befe4693)
	msgpack.Register((*ds.Cursor)(nil), __obf_c2f88ad7cfb13871, __obf_5dd8db9729ec977b)
}

func EncodeDatastoreKey(__obf_f0d8c07be0e293ee *msgpack.Encoder, __obf_78ae89912cd184b9 *ds.Key) error {
	if __obf_78ae89912cd184b9 == nil {
		return __obf_f0d8c07be0e293ee.EncodeNil()
	}
	return __obf_f0d8c07be0e293ee.EncodeString(__obf_78ae89912cd184b9.Encode())
}

func __obf_c42aaa45510b7cdc(__obf_f0d8c07be0e293ee *msgpack.Encoder, __obf_6cb61f9cf11d3853 reflect.Value) error {
	__obf_78ae89912cd184b9 := __obf_6cb61f9cf11d3853.Interface().(*ds.Key)
	return EncodeDatastoreKey(__obf_f0d8c07be0e293ee, __obf_78ae89912cd184b9)
}

func DecodeDatastoreKey(__obf_9f7795339dd8959b *msgpack.Decoder) (*ds.Key, error) {
	__obf_6cb61f9cf11d3853, __obf_2f89ec15ad731e29 := __obf_9f7795339dd8959b.DecodeString()
	if __obf_2f89ec15ad731e29 != nil {
		return nil, __obf_2f89ec15ad731e29
	}
	if __obf_6cb61f9cf11d3853 == "" {
		return nil, nil
	}
	return ds.DecodeKey(__obf_6cb61f9cf11d3853)
}

func __obf_87a05698befe4693(__obf_9f7795339dd8959b *msgpack.Decoder, __obf_6cb61f9cf11d3853 reflect.Value) error {
	__obf_78ae89912cd184b9, __obf_2f89ec15ad731e29 := DecodeDatastoreKey(__obf_9f7795339dd8959b)
	if __obf_2f89ec15ad731e29 != nil {
		return __obf_2f89ec15ad731e29
	}
	__obf_6cb61f9cf11d3853.
		Set(reflect.ValueOf(__obf_78ae89912cd184b9))
	return nil
}

func __obf_c2f88ad7cfb13871(__obf_f0d8c07be0e293ee *msgpack.Encoder, __obf_6cb61f9cf11d3853 reflect.Value) error {
	__obf_5fee15608b2d4718 := __obf_6cb61f9cf11d3853.Interface().(ds.Cursor)
	return __obf_f0d8c07be0e293ee.Encode(__obf_5fee15608b2d4718.String())
}

func __obf_5dd8db9729ec977b(__obf_9f7795339dd8959b *msgpack.Decoder, __obf_6cb61f9cf11d3853 reflect.Value) error {
	__obf_43566c009e3caf91, __obf_2f89ec15ad731e29 := __obf_9f7795339dd8959b.DecodeString()
	if __obf_2f89ec15ad731e29 != nil {
		return __obf_2f89ec15ad731e29
	}
	__obf_5fee15608b2d4718, __obf_2f89ec15ad731e29 := ds.DecodeCursor(__obf_43566c009e3caf91)
	if __obf_2f89ec15ad731e29 != nil {
		return __obf_2f89ec15ad731e29
	}
	__obf_6cb61f9cf11d3853.
		Set(reflect.ValueOf(__obf_5fee15608b2d4718))
	return nil
}
