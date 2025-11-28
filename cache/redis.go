package __obf_32d927a1e06b7e71

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/redis/go-redis/v9"
)

var __obf_2bd6fb55a3ccae1f = context.Background()

type Redis struct {
	__obf_29f6424fa0c73eda *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_3bb2dfd7d27563f9 RdOption) *Redis {
	return &Redis{__obf_29f6424fa0c73eda: redis.NewClient(&redis.Options{
		Addr:            __obf_3bb2dfd7d27563f9.Host,
		Password:        __obf_3bb2dfd7d27563f9.Password,
		DB:              __obf_3bb2dfd7d27563f9.DB,
		MaxIdleConns:    __obf_3bb2dfd7d27563f9.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_3bb2dfd7d27563f9.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_bb9a2535b3ac3a55 *Redis) Delete(__obf_e5c9d9e65a3fa384 string) {
	__obf_444e1925194b0bc3, __obf_19fb5c4c25ff452a := __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.Keys(__obf_2bd6fb55a3ccae1f, __obf_e5c9d9e65a3fa384).Result()
	if __obf_19fb5c4c25ff452a != nil {
		for _, __obf_4be74b34d9a5767c := range __obf_444e1925194b0bc3 {
			__obf_bb9a2535b3ac3a55.Remove(__obf_4be74b34d9a5767c)
		}
	}
}

func (__obf_bb9a2535b3ac3a55 *Redis) Add(__obf_13113b328a6972ad string, __obf_a967d4ebf1531f4c any) error {
	return __obf_bb9a2535b3ac3a55.Set(__obf_13113b328a6972ad, __obf_a967d4ebf1531f4c, 0)
}

func (__obf_bb9a2535b3ac3a55 *Redis) Set(__obf_13113b328a6972ad string, __obf_a967d4ebf1531f4c any, __obf_481cbd2caaded2ed time.Duration) error {
	if __obf_cf59be873d0aaa34, __obf_542c0401f22c9aec := __obf_a967d4ebf1531f4c.([]byte); __obf_542c0401f22c9aec {
		return __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.Set(__obf_2bd6fb55a3ccae1f, __obf_13113b328a6972ad, __obf_cf59be873d0aaa34, __obf_481cbd2caaded2ed).Err()
	} else {
		__obf_9fc6c45711ce2370, __obf_19fb5c4c25ff452a := util.AnyToBytes(__obf_a967d4ebf1531f4c)
		if __obf_19fb5c4c25ff452a != nil {
			return fmt.Errorf("serialize error: %s", __obf_19fb5c4c25ff452a)
		}
		return __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.Set(__obf_2bd6fb55a3ccae1f, __obf_13113b328a6972ad, __obf_9fc6c45711ce2370, __obf_481cbd2caaded2ed).Err()
	}
}

func (__obf_bb9a2535b3ac3a55 *Redis) Get(__obf_13113b328a6972ad string) ([]byte, error) {
	return __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.Get(__obf_2bd6fb55a3ccae1f, __obf_13113b328a6972ad).Bytes()
}

func (__obf_bb9a2535b3ac3a55 *Redis) Remove(__obf_13113b328a6972ad string) error {
	return __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.Del(__obf_2bd6fb55a3ccae1f, __obf_13113b328a6972ad).Err()
}

// ClearAll clears all cache in memcache.
func (__obf_bb9a2535b3ac3a55 *Redis) Clear() error {
	__obf_978c7980bab17b1b := __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.FlushAll(__obf_2bd6fb55a3ccae1f)
	return __obf_978c7980bab17b1b.Err()
}

func (__obf_bb9a2535b3ac3a55 *Redis) Has(__obf_13113b328a6972ad string) bool {
	return __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.Exists(__obf_2bd6fb55a3ccae1f, __obf_13113b328a6972ad).Val() == 1
}

func (__obf_bb9a2535b3ac3a55 *Redis) Renew(__obf_13113b328a6972ad string, __obf_481cbd2caaded2ed time.Duration) error {
	__obf_a967d4ebf1531f4c, __obf_19fb5c4c25ff452a := __obf_bb9a2535b3ac3a55.Get(__obf_13113b328a6972ad)
	if __obf_19fb5c4c25ff452a != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_13113b328a6972ad, __obf_19fb5c4c25ff452a)
	}
	return __obf_bb9a2535b3ac3a55.Set(__obf_13113b328a6972ad, __obf_a967d4ebf1531f4c, __obf_481cbd2caaded2ed)
}

func (__obf_bb9a2535b3ac3a55 *Redis) HSet(__obf_2bd6fb55a3ccae1f context.Context, __obf_13113b328a6972ad string, __obf_7b0000b39332c436 ...any) error {
	return __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.HSet(__obf_2bd6fb55a3ccae1f, __obf_13113b328a6972ad, __obf_7b0000b39332c436...).Err()
}

func (__obf_bb9a2535b3ac3a55 *Redis) HGet(__obf_2bd6fb55a3ccae1f context.Context, __obf_13113b328a6972ad, __obf_f610592e11b0aaf7 string) error {
	return __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.HGet(__obf_2bd6fb55a3ccae1f, __obf_13113b328a6972ad, __obf_f610592e11b0aaf7).Err()
}

func (__obf_bb9a2535b3ac3a55 *Redis) HGetAll(__obf_2bd6fb55a3ccae1f context.Context, __obf_13113b328a6972ad string) (map[string]string, error) {
	return __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.HGetAll(__obf_2bd6fb55a3ccae1f, __obf_13113b328a6972ad).Result()
}

func (__obf_bb9a2535b3ac3a55 *Redis) TxPipeline() redis.Pipeliner {
	return __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.TxPipeline()
}

func (__obf_bb9a2535b3ac3a55 *Redis) HDel(__obf_2bd6fb55a3ccae1f context.Context, __obf_13113b328a6972ad, __obf_fd7825f5851ff090 string) error {
	return __obf_bb9a2535b3ac3a55.__obf_29f6424fa0c73eda.HDel(__obf_2bd6fb55a3ccae1f, __obf_13113b328a6972ad, __obf_fd7825f5851ff090).Err()
}
