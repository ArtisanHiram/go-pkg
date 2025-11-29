package __obf_0e388cc3170ce2cd

import (
	msgpack "github.com/ArtisanHiram/go-pkg/msgpack"
	ds "google.golang.org/appengine/datastore"
	"reflect"
)

func init() {
	msgpack.Register((*ds.Key)(nil), __obf_93e0ec8b6323cbaf, __obf_4a87c8ef369a5296)
	msgpack.Register((*ds.Cursor)(nil), __obf_a2e9a4753e66c3f2, __obf_95d8917c61b2fec3)
}

func EncodeDatastoreKey(__obf_63e332877d8a7762 *msgpack.Encoder, __obf_6e4b24df668b6c53 *ds.Key) error {
	if __obf_6e4b24df668b6c53 == nil {
		return __obf_63e332877d8a7762.EncodeNil()
	}
	return __obf_63e332877d8a7762.EncodeString(__obf_6e4b24df668b6c53.Encode())
}

func __obf_93e0ec8b6323cbaf(__obf_63e332877d8a7762 *msgpack.Encoder, __obf_7651bcbf2fbe9702 reflect.Value) error {
	__obf_6e4b24df668b6c53 := __obf_7651bcbf2fbe9702.Interface().(*ds.Key)
	return EncodeDatastoreKey(__obf_63e332877d8a7762, __obf_6e4b24df668b6c53)
}

func DecodeDatastoreKey(__obf_bfd66a7c497b59e0 *msgpack.Decoder) (*ds.Key, error) {
	__obf_7651bcbf2fbe9702, __obf_0fdb31ee52bade7b := __obf_bfd66a7c497b59e0.DecodeString()
	if __obf_0fdb31ee52bade7b != nil {
		return nil, __obf_0fdb31ee52bade7b
	}
	if __obf_7651bcbf2fbe9702 == "" {
		return nil, nil
	}
	return ds.DecodeKey(__obf_7651bcbf2fbe9702)
}

func __obf_4a87c8ef369a5296(__obf_bfd66a7c497b59e0 *msgpack.Decoder, __obf_7651bcbf2fbe9702 reflect.Value) error {
	__obf_6e4b24df668b6c53, __obf_0fdb31ee52bade7b := DecodeDatastoreKey(__obf_bfd66a7c497b59e0)
	if __obf_0fdb31ee52bade7b != nil {
		return __obf_0fdb31ee52bade7b
	}
	__obf_7651bcbf2fbe9702.
		Set(reflect.ValueOf(__obf_6e4b24df668b6c53))
	return nil
}

func __obf_a2e9a4753e66c3f2(__obf_63e332877d8a7762 *msgpack.Encoder, __obf_7651bcbf2fbe9702 reflect.Value) error {
	__obf_9db1dc874255f6bb := __obf_7651bcbf2fbe9702.Interface().(ds.Cursor)
	return __obf_63e332877d8a7762.Encode(__obf_9db1dc874255f6bb.String())
}

func __obf_95d8917c61b2fec3(__obf_bfd66a7c497b59e0 *msgpack.Decoder, __obf_7651bcbf2fbe9702 reflect.Value) error {
	__obf_207d4422c98c177b, __obf_0fdb31ee52bade7b := __obf_bfd66a7c497b59e0.DecodeString()
	if __obf_0fdb31ee52bade7b != nil {
		return __obf_0fdb31ee52bade7b
	}
	__obf_9db1dc874255f6bb, __obf_0fdb31ee52bade7b := ds.DecodeCursor(__obf_207d4422c98c177b)
	if __obf_0fdb31ee52bade7b != nil {
		return __obf_0fdb31ee52bade7b
	}
	__obf_7651bcbf2fbe9702.
		Set(reflect.ValueOf(__obf_9db1dc874255f6bb))
	return nil
}
