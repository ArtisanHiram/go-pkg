package __obf_88e12bac16d31f5b

import (
	msgpack "github.com/ArtisanHiram/go-pkg/msgpack"
	ds "google.golang.org/appengine/datastore"
	"reflect"
)

func init() {
	msgpack.Register((*ds.Key)(nil), __obf_3f2fe07338efb3e1, __obf_78d28a837cf6fee6)
	msgpack.Register((*ds.Cursor)(nil), __obf_7b3caed5111b15c1, __obf_db624dea98d302b4)
}

func EncodeDatastoreKey(__obf_1e7adefbf4ed0863 *msgpack.Encoder, __obf_b90f0ea3693d8b5b *ds.Key) error {
	if __obf_b90f0ea3693d8b5b == nil {
		return __obf_1e7adefbf4ed0863.EncodeNil()
	}
	return __obf_1e7adefbf4ed0863.EncodeString(__obf_b90f0ea3693d8b5b.Encode())
}

func __obf_3f2fe07338efb3e1(__obf_1e7adefbf4ed0863 *msgpack.Encoder, __obf_f5976b3105c91416 reflect.Value) error {
	__obf_b90f0ea3693d8b5b := __obf_f5976b3105c91416.Interface().(*ds.Key)
	return EncodeDatastoreKey(__obf_1e7adefbf4ed0863, __obf_b90f0ea3693d8b5b)
}

func DecodeDatastoreKey(__obf_a5d5503ec6188b3d *msgpack.Decoder) (*ds.Key, error) {
	__obf_f5976b3105c91416, __obf_c622fa5d95d390a5 := __obf_a5d5503ec6188b3d.DecodeString()
	if __obf_c622fa5d95d390a5 != nil {
		return nil, __obf_c622fa5d95d390a5
	}
	if __obf_f5976b3105c91416 == "" {
		return nil, nil
	}
	return ds.DecodeKey(__obf_f5976b3105c91416)
}

func __obf_78d28a837cf6fee6(__obf_a5d5503ec6188b3d *msgpack.Decoder, __obf_f5976b3105c91416 reflect.Value) error {
	__obf_b90f0ea3693d8b5b, __obf_c622fa5d95d390a5 := DecodeDatastoreKey(__obf_a5d5503ec6188b3d)
	if __obf_c622fa5d95d390a5 != nil {
		return __obf_c622fa5d95d390a5
	}
	__obf_f5976b3105c91416.
		Set(reflect.ValueOf(__obf_b90f0ea3693d8b5b))
	return nil
}

func __obf_7b3caed5111b15c1(__obf_1e7adefbf4ed0863 *msgpack.Encoder, __obf_f5976b3105c91416 reflect.Value) error {
	__obf_e863fdad66f088c2 := __obf_f5976b3105c91416.Interface().(ds.Cursor)
	return __obf_1e7adefbf4ed0863.Encode(__obf_e863fdad66f088c2.String())
}

func __obf_db624dea98d302b4(__obf_a5d5503ec6188b3d *msgpack.Decoder, __obf_f5976b3105c91416 reflect.Value) error {
	__obf_fd5a035cfdb1508a, __obf_c622fa5d95d390a5 := __obf_a5d5503ec6188b3d.DecodeString()
	if __obf_c622fa5d95d390a5 != nil {
		return __obf_c622fa5d95d390a5
	}
	__obf_e863fdad66f088c2, __obf_c622fa5d95d390a5 := ds.DecodeCursor(__obf_fd5a035cfdb1508a)
	if __obf_c622fa5d95d390a5 != nil {
		return __obf_c622fa5d95d390a5
	}
	__obf_f5976b3105c91416.
		Set(reflect.ValueOf(__obf_e863fdad66f088c2))
	return nil
}
