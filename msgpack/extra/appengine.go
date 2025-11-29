package __obf_8f725c4dd8f38531

import (
	msgpack "github.com/ArtisanHiram/go-pkg/msgpack"
	ds "google.golang.org/appengine/datastore"
	"reflect"
)

func init() {
	msgpack.Register((*ds.Key)(nil), __obf_7ca5bbab01136a74, __obf_95a30443b1ad64cc)
	msgpack.Register((*ds.Cursor)(nil), __obf_5081dea10a71e7da, __obf_069377e75c5d3c3d)
}

func EncodeDatastoreKey(__obf_a8ad6b0d863e3198 *msgpack.Encoder, __obf_2173e29e8dad0c0b *ds.Key) error {
	if __obf_2173e29e8dad0c0b == nil {
		return __obf_a8ad6b0d863e3198.EncodeNil()
	}
	return __obf_a8ad6b0d863e3198.EncodeString(__obf_2173e29e8dad0c0b.Encode())
}

func __obf_7ca5bbab01136a74(__obf_a8ad6b0d863e3198 *msgpack.Encoder, __obf_370b9adf405669e7 reflect.Value) error {
	__obf_2173e29e8dad0c0b := __obf_370b9adf405669e7.Interface().(*ds.Key)
	return EncodeDatastoreKey(__obf_a8ad6b0d863e3198, __obf_2173e29e8dad0c0b)
}

func DecodeDatastoreKey(__obf_3b9e094fe11c6458 *msgpack.Decoder) (*ds.Key, error) {
	__obf_370b9adf405669e7, __obf_1a66af63cf1f7f5f := __obf_3b9e094fe11c6458.DecodeString()
	if __obf_1a66af63cf1f7f5f != nil {
		return nil, __obf_1a66af63cf1f7f5f
	}
	if __obf_370b9adf405669e7 == "" {
		return nil, nil
	}
	return ds.DecodeKey(__obf_370b9adf405669e7)
}

func __obf_95a30443b1ad64cc(__obf_3b9e094fe11c6458 *msgpack.Decoder, __obf_370b9adf405669e7 reflect.Value) error {
	__obf_2173e29e8dad0c0b, __obf_1a66af63cf1f7f5f := DecodeDatastoreKey(__obf_3b9e094fe11c6458)
	if __obf_1a66af63cf1f7f5f != nil {
		return __obf_1a66af63cf1f7f5f
	}
	__obf_370b9adf405669e7.
		Set(reflect.ValueOf(__obf_2173e29e8dad0c0b))
	return nil
}

func __obf_5081dea10a71e7da(__obf_a8ad6b0d863e3198 *msgpack.Encoder, __obf_370b9adf405669e7 reflect.Value) error {
	__obf_d4c7e523a9062e26 := __obf_370b9adf405669e7.Interface().(ds.Cursor)
	return __obf_a8ad6b0d863e3198.Encode(__obf_d4c7e523a9062e26.String())
}

func __obf_069377e75c5d3c3d(__obf_3b9e094fe11c6458 *msgpack.Decoder, __obf_370b9adf405669e7 reflect.Value) error {
	__obf_5366ed31edfb2b22, __obf_1a66af63cf1f7f5f := __obf_3b9e094fe11c6458.DecodeString()
	if __obf_1a66af63cf1f7f5f != nil {
		return __obf_1a66af63cf1f7f5f
	}
	__obf_d4c7e523a9062e26, __obf_1a66af63cf1f7f5f := ds.DecodeCursor(__obf_5366ed31edfb2b22)
	if __obf_1a66af63cf1f7f5f != nil {
		return __obf_1a66af63cf1f7f5f
	}
	__obf_370b9adf405669e7.
		Set(reflect.ValueOf(__obf_d4c7e523a9062e26))
	return nil
}
