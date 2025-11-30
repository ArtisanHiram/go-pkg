package __obf_8fd9db4a493b531f

import (
	msgpack "github.com/ArtisanHiram/go-pkg/msgpack"
	ds "google.golang.org/appengine/datastore"
	"reflect"
)

func init() {
	msgpack.Register((*ds.Key)(nil), __obf_94d227b6d60dc56c, __obf_b8d62cf18ea15481)
	msgpack.Register((*ds.Cursor)(nil), __obf_3659ca27e634fd20, __obf_14b633cd159f3698)
}

func EncodeDatastoreKey(__obf_713fbf9577717da8 *msgpack.Encoder, __obf_994cdbcd2230267d *ds.Key) error {
	if __obf_994cdbcd2230267d == nil {
		return __obf_713fbf9577717da8.EncodeNil()
	}
	return __obf_713fbf9577717da8.EncodeString(__obf_994cdbcd2230267d.Encode())
}

func __obf_94d227b6d60dc56c(__obf_713fbf9577717da8 *msgpack.Encoder, __obf_b15f1066f0f4f665 reflect.Value) error {
	__obf_994cdbcd2230267d := __obf_b15f1066f0f4f665.Interface().(*ds.Key)
	return EncodeDatastoreKey(__obf_713fbf9577717da8, __obf_994cdbcd2230267d)
}

func DecodeDatastoreKey(__obf_04c5b7ecde089247 *msgpack.Decoder) (*ds.Key, error) {
	__obf_b15f1066f0f4f665, __obf_dda9f7294d0c88e9 := __obf_04c5b7ecde089247.DecodeString()
	if __obf_dda9f7294d0c88e9 != nil {
		return nil, __obf_dda9f7294d0c88e9
	}
	if __obf_b15f1066f0f4f665 == "" {
		return nil, nil
	}
	return ds.DecodeKey(__obf_b15f1066f0f4f665)
}

func __obf_b8d62cf18ea15481(__obf_04c5b7ecde089247 *msgpack.Decoder, __obf_b15f1066f0f4f665 reflect.Value) error {
	__obf_994cdbcd2230267d, __obf_dda9f7294d0c88e9 := DecodeDatastoreKey(__obf_04c5b7ecde089247)
	if __obf_dda9f7294d0c88e9 != nil {
		return __obf_dda9f7294d0c88e9
	}
	__obf_b15f1066f0f4f665.
		Set(reflect.ValueOf(__obf_994cdbcd2230267d))
	return nil
}

func __obf_3659ca27e634fd20(__obf_713fbf9577717da8 *msgpack.Encoder, __obf_b15f1066f0f4f665 reflect.Value) error {
	__obf_ec9a1006b05d6f00 := __obf_b15f1066f0f4f665.Interface().(ds.Cursor)
	return __obf_713fbf9577717da8.Encode(__obf_ec9a1006b05d6f00.String())
}

func __obf_14b633cd159f3698(__obf_04c5b7ecde089247 *msgpack.Decoder, __obf_b15f1066f0f4f665 reflect.Value) error {
	__obf_040b711b12696910, __obf_dda9f7294d0c88e9 := __obf_04c5b7ecde089247.DecodeString()
	if __obf_dda9f7294d0c88e9 != nil {
		return __obf_dda9f7294d0c88e9
	}
	__obf_ec9a1006b05d6f00, __obf_dda9f7294d0c88e9 := ds.DecodeCursor(__obf_040b711b12696910)
	if __obf_dda9f7294d0c88e9 != nil {
		return __obf_dda9f7294d0c88e9
	}
	__obf_b15f1066f0f4f665.
		Set(reflect.ValueOf(__obf_ec9a1006b05d6f00))
	return nil
}
