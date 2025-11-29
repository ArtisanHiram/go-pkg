package __obf_a05682be1c6caf18

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/redis/go-redis/v9"
)

var __obf_20ad0223e112e1c1 = context.Background()

type Redis struct {
	__obf_21fd2f3c74fafdf1 *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_0a44d99b97432d49 RdOption) *Redis {
	return &Redis{__obf_21fd2f3c74fafdf1: redis.NewClient(&redis.Options{
		Addr:            __obf_0a44d99b97432d49.Host,
		Password:        __obf_0a44d99b97432d49.Password,
		DB:              __obf_0a44d99b97432d49.DB,
		MaxIdleConns:    __obf_0a44d99b97432d49.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_0a44d99b97432d49.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_3f3723f674c63e34 *Redis) Delete(__obf_715bf114d0b286fb string) {
	__obf_3a3212e21950211c, __obf_94de95cb433b50f1 := __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.Keys(__obf_20ad0223e112e1c1, __obf_715bf114d0b286fb).Result()
	if __obf_94de95cb433b50f1 != nil {
		for _, __obf_b0c62418c29009a4 := range __obf_3a3212e21950211c {
			__obf_3f3723f674c63e34.
				Remove(__obf_b0c62418c29009a4)
		}
	}
}

func (__obf_3f3723f674c63e34 *Redis) Add(__obf_a0e56915a05e8d99 string, __obf_6030b68d8a95172f any) error {
	return __obf_3f3723f674c63e34.Set(__obf_a0e56915a05e8d99, __obf_6030b68d8a95172f, 0)
}

func (__obf_3f3723f674c63e34 *Redis) Set(__obf_a0e56915a05e8d99 string, __obf_6030b68d8a95172f any, __obf_c8947601a564de1b time.Duration) error {
	if __obf_40394ecac45b9c79, __obf_19211ff80aaad9b4 := __obf_6030b68d8a95172f.([]byte); __obf_19211ff80aaad9b4 {
		return __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.Set(__obf_20ad0223e112e1c1, __obf_a0e56915a05e8d99, __obf_40394ecac45b9c79, __obf_c8947601a564de1b).Err()
	} else {
		__obf_cc8497e7c2fb7f28, __obf_94de95cb433b50f1 := util.AnyToBytes(__obf_6030b68d8a95172f)
		if __obf_94de95cb433b50f1 != nil {
			return fmt.Errorf("serialize error: %s", __obf_94de95cb433b50f1)
		}
		return __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.Set(__obf_20ad0223e112e1c1, __obf_a0e56915a05e8d99, __obf_cc8497e7c2fb7f28, __obf_c8947601a564de1b).Err()
	}
}

func (__obf_3f3723f674c63e34 *Redis) Get(__obf_a0e56915a05e8d99 string) ([]byte, error) {
	return __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.Get(__obf_20ad0223e112e1c1, __obf_a0e56915a05e8d99).Bytes()
}

func (__obf_3f3723f674c63e34 *Redis) Remove(__obf_a0e56915a05e8d99 string) error {
	return __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.Del(__obf_20ad0223e112e1c1,

		// ClearAll clears all cache in memcache.
		__obf_a0e56915a05e8d99).Err()
}

func (__obf_3f3723f674c63e34 *Redis) Clear() error {
	__obf_020fc4bcc3246438 := __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.FlushAll(__obf_20ad0223e112e1c1)
	return __obf_020fc4bcc3246438.Err()
}

func (__obf_3f3723f674c63e34 *Redis) Has(__obf_a0e56915a05e8d99 string) bool {
	return __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.Exists(__obf_20ad0223e112e1c1, __obf_a0e56915a05e8d99).Val() == 1
}

func (__obf_3f3723f674c63e34 *Redis) Renew(__obf_a0e56915a05e8d99 string, __obf_c8947601a564de1b time.Duration) error {
	__obf_6030b68d8a95172f, __obf_94de95cb433b50f1 := __obf_3f3723f674c63e34.Get(__obf_a0e56915a05e8d99)
	if __obf_94de95cb433b50f1 != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_a0e56915a05e8d99, __obf_94de95cb433b50f1)
	}
	return __obf_3f3723f674c63e34.Set(__obf_a0e56915a05e8d99, __obf_6030b68d8a95172f, __obf_c8947601a564de1b)
}

func (__obf_3f3723f674c63e34 *Redis) HSet(__obf_20ad0223e112e1c1 context.Context, __obf_a0e56915a05e8d99 string, __obf_2bafbfda1c74db8f ...any) error {
	return __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.HSet(__obf_20ad0223e112e1c1, __obf_a0e56915a05e8d99, __obf_2bafbfda1c74db8f...).Err()
}

func (__obf_3f3723f674c63e34 *Redis) HGet(__obf_20ad0223e112e1c1 context.Context, __obf_a0e56915a05e8d99, __obf_8c6292045772d5f1 string) error {
	return __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.HGet(__obf_20ad0223e112e1c1, __obf_a0e56915a05e8d99, __obf_8c6292045772d5f1).Err()
}

func (__obf_3f3723f674c63e34 *Redis) HGetAll(__obf_20ad0223e112e1c1 context.Context, __obf_a0e56915a05e8d99 string) (map[string]string, error) {
	return __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.HGetAll(__obf_20ad0223e112e1c1, __obf_a0e56915a05e8d99).Result()
}

func (__obf_3f3723f674c63e34 *Redis) TxPipeline() redis.Pipeliner {
	return __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.TxPipeline()
}

func (__obf_3f3723f674c63e34 *Redis) HDel(__obf_20ad0223e112e1c1 context.Context, __obf_a0e56915a05e8d99, __obf_a440ced657a2a041 string) error {
	return __obf_3f3723f674c63e34.__obf_21fd2f3c74fafdf1.HDel(__obf_20ad0223e112e1c1, __obf_a0e56915a05e8d99, __obf_a440ced657a2a041).Err()
}
