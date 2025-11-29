package __obf_6fd4fe34e3f784df

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/redis/go-redis/v9"
)

var __obf_71755847ac967507 = context.Background()

type Redis struct {
	__obf_394e0737b83973f0 *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_b9532de317f57626 RdOption) *Redis {
	return &Redis{__obf_394e0737b83973f0: redis.NewClient(&redis.Options{
		Addr:            __obf_b9532de317f57626.Host,
		Password:        __obf_b9532de317f57626.Password,
		DB:              __obf_b9532de317f57626.DB,
		MaxIdleConns:    __obf_b9532de317f57626.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_b9532de317f57626.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_129832cddf20a553 *Redis) Delete(__obf_c53a458d6b02aeb9 string) {
	__obf_007ba5e3d2b495b2, __obf_7f8f5f5c173fa58e := __obf_129832cddf20a553.__obf_394e0737b83973f0.Keys(__obf_71755847ac967507, __obf_c53a458d6b02aeb9).Result()
	if __obf_7f8f5f5c173fa58e != nil {
		for _, __obf_2a0493c028b94053 := range __obf_007ba5e3d2b495b2 {
			__obf_129832cddf20a553.
				Remove(__obf_2a0493c028b94053)
		}
	}
}

func (__obf_129832cddf20a553 *Redis) Add(__obf_ca986a38d1f8fbbb string, __obf_e1d0258259217620 any) error {
	return __obf_129832cddf20a553.Set(__obf_ca986a38d1f8fbbb, __obf_e1d0258259217620, 0)
}

func (__obf_129832cddf20a553 *Redis) Set(__obf_ca986a38d1f8fbbb string, __obf_e1d0258259217620 any, __obf_bfbeaaea582ee60a time.Duration) error {
	if __obf_b1ecf81133de0ff6, __obf_949c761ffd90d5fa := __obf_e1d0258259217620.([]byte); __obf_949c761ffd90d5fa {
		return __obf_129832cddf20a553.__obf_394e0737b83973f0.Set(__obf_71755847ac967507, __obf_ca986a38d1f8fbbb, __obf_b1ecf81133de0ff6, __obf_bfbeaaea582ee60a).Err()
	} else {
		__obf_6e42d0b5efdcdf62, __obf_7f8f5f5c173fa58e := util.AnyToBytes(__obf_e1d0258259217620)
		if __obf_7f8f5f5c173fa58e != nil {
			return fmt.Errorf("serialize error: %s", __obf_7f8f5f5c173fa58e)
		}
		return __obf_129832cddf20a553.__obf_394e0737b83973f0.Set(__obf_71755847ac967507, __obf_ca986a38d1f8fbbb, __obf_6e42d0b5efdcdf62, __obf_bfbeaaea582ee60a).Err()
	}
}

func (__obf_129832cddf20a553 *Redis) Get(__obf_ca986a38d1f8fbbb string) ([]byte, error) {
	return __obf_129832cddf20a553.__obf_394e0737b83973f0.Get(__obf_71755847ac967507, __obf_ca986a38d1f8fbbb).Bytes()
}

func (__obf_129832cddf20a553 *Redis) Remove(__obf_ca986a38d1f8fbbb string) error {
	return __obf_129832cddf20a553.__obf_394e0737b83973f0.Del(__obf_71755847ac967507,

		// ClearAll clears all cache in memcache.
		__obf_ca986a38d1f8fbbb).Err()
}

func (__obf_129832cddf20a553 *Redis) Clear() error {
	__obf_c9f1ce6b30a488ef := __obf_129832cddf20a553.__obf_394e0737b83973f0.FlushAll(__obf_71755847ac967507)
	return __obf_c9f1ce6b30a488ef.Err()
}

func (__obf_129832cddf20a553 *Redis) Has(__obf_ca986a38d1f8fbbb string) bool {
	return __obf_129832cddf20a553.__obf_394e0737b83973f0.Exists(__obf_71755847ac967507, __obf_ca986a38d1f8fbbb).Val() == 1
}

func (__obf_129832cddf20a553 *Redis) Renew(__obf_ca986a38d1f8fbbb string, __obf_bfbeaaea582ee60a time.Duration) error {
	__obf_e1d0258259217620, __obf_7f8f5f5c173fa58e := __obf_129832cddf20a553.Get(__obf_ca986a38d1f8fbbb)
	if __obf_7f8f5f5c173fa58e != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_ca986a38d1f8fbbb, __obf_7f8f5f5c173fa58e)
	}
	return __obf_129832cddf20a553.Set(__obf_ca986a38d1f8fbbb, __obf_e1d0258259217620, __obf_bfbeaaea582ee60a)
}

func (__obf_129832cddf20a553 *Redis) HSet(__obf_71755847ac967507 context.Context, __obf_ca986a38d1f8fbbb string, __obf_415a2d23340fabda ...any) error {
	return __obf_129832cddf20a553.__obf_394e0737b83973f0.HSet(__obf_71755847ac967507, __obf_ca986a38d1f8fbbb, __obf_415a2d23340fabda...).Err()
}

func (__obf_129832cddf20a553 *Redis) HGet(__obf_71755847ac967507 context.Context, __obf_ca986a38d1f8fbbb, __obf_28a3c2b56e22b8f1 string) error {
	return __obf_129832cddf20a553.__obf_394e0737b83973f0.HGet(__obf_71755847ac967507, __obf_ca986a38d1f8fbbb, __obf_28a3c2b56e22b8f1).Err()
}

func (__obf_129832cddf20a553 *Redis) HGetAll(__obf_71755847ac967507 context.Context, __obf_ca986a38d1f8fbbb string) (map[string]string, error) {
	return __obf_129832cddf20a553.__obf_394e0737b83973f0.HGetAll(__obf_71755847ac967507, __obf_ca986a38d1f8fbbb).Result()
}

func (__obf_129832cddf20a553 *Redis) TxPipeline() redis.Pipeliner {
	return __obf_129832cddf20a553.__obf_394e0737b83973f0.TxPipeline()
}

func (__obf_129832cddf20a553 *Redis) HDel(__obf_71755847ac967507 context.Context, __obf_ca986a38d1f8fbbb, __obf_b9e5891ea6c77c07 string) error {
	return __obf_129832cddf20a553.__obf_394e0737b83973f0.HDel(__obf_71755847ac967507, __obf_ca986a38d1f8fbbb, __obf_b9e5891ea6c77c07).Err()
}
