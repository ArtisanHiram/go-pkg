package __obf_62df4de078c8d208

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	redis "github.com/redis/go-redis/v9"
)

var __obf_9b31d4a4d4adfff9 = context.Background()

type Redis struct {
	__obf_bb6a956ebc49aa20 *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_a5ce663488889fe6 RdOption) *Redis {
	return &Redis{__obf_bb6a956ebc49aa20: redis.NewClient(&redis.Options{
		Addr:            __obf_a5ce663488889fe6.Host,
		Password:        __obf_a5ce663488889fe6.Password,
		DB:              __obf_a5ce663488889fe6.DB,
		MaxIdleConns:    __obf_a5ce663488889fe6.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_a5ce663488889fe6.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_2d21271215a2e06f *Redis) Delete(__obf_0c00b35780000392 string) {
	__obf_5b4b8302049facb2, __obf_0560c6b8f27080cc := __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.Keys(__obf_9b31d4a4d4adfff9, __obf_0c00b35780000392).Result()
	if __obf_0560c6b8f27080cc != nil {
		for _, __obf_c0676b3110071e57 := range __obf_5b4b8302049facb2 {
			__obf_2d21271215a2e06f.Remove(__obf_c0676b3110071e57)
		}
	}
}

func (__obf_2d21271215a2e06f *Redis) Add(__obf_4aecf3c737bbe5e8 string, __obf_676d75836c094b83 any) error {
	return __obf_2d21271215a2e06f.Set(__obf_4aecf3c737bbe5e8, __obf_676d75836c094b83, 0)
}

func (__obf_2d21271215a2e06f *Redis) Set(__obf_4aecf3c737bbe5e8 string, __obf_676d75836c094b83 any, __obf_24d92102aed9ce02 time.Duration) error {
	if __obf_298afaabecd63bb9, __obf_4bca5d4f1d82d040 := __obf_676d75836c094b83.([]byte); __obf_4bca5d4f1d82d040 {
		return __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.Set(__obf_9b31d4a4d4adfff9, __obf_4aecf3c737bbe5e8, __obf_298afaabecd63bb9, __obf_24d92102aed9ce02).Err()
	} else {
		__obf_0ff820cf70b68576, __obf_0560c6b8f27080cc := util.AnyToBytes(__obf_676d75836c094b83)
		if __obf_0560c6b8f27080cc != nil {
			return fmt.Errorf("serialize error: %s", __obf_0560c6b8f27080cc)
		}
		return __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.Set(__obf_9b31d4a4d4adfff9, __obf_4aecf3c737bbe5e8, __obf_0ff820cf70b68576, __obf_24d92102aed9ce02).Err()
	}
}

func (__obf_2d21271215a2e06f *Redis) Get(__obf_4aecf3c737bbe5e8 string) ([]byte, error) {
	return __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.Get(__obf_9b31d4a4d4adfff9, __obf_4aecf3c737bbe5e8).Bytes()
}

func (__obf_2d21271215a2e06f *Redis) Remove(__obf_4aecf3c737bbe5e8 string) error {
	return __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.Del(__obf_9b31d4a4d4adfff9, __obf_4aecf3c737bbe5e8).Err()
}

// ClearAll clears all cache in memcache.
func (__obf_2d21271215a2e06f *Redis) Clear() error {
	__obf_525bf67953d7093e := __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.FlushAll(__obf_9b31d4a4d4adfff9)
	return __obf_525bf67953d7093e.Err()
}

func (__obf_2d21271215a2e06f *Redis) Has(__obf_4aecf3c737bbe5e8 string) bool {
	return __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.Exists(__obf_9b31d4a4d4adfff9, __obf_4aecf3c737bbe5e8).Val() == 1
}

func (__obf_2d21271215a2e06f *Redis) Renew(__obf_4aecf3c737bbe5e8 string, __obf_24d92102aed9ce02 time.Duration) error {
	__obf_676d75836c094b83, __obf_0560c6b8f27080cc := __obf_2d21271215a2e06f.Get(__obf_4aecf3c737bbe5e8)
	if __obf_0560c6b8f27080cc != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_4aecf3c737bbe5e8, __obf_0560c6b8f27080cc)
	}
	return __obf_2d21271215a2e06f.Set(__obf_4aecf3c737bbe5e8, __obf_676d75836c094b83, __obf_24d92102aed9ce02)
}

func (__obf_2d21271215a2e06f *Redis) HSet(__obf_9b31d4a4d4adfff9 context.Context, __obf_4aecf3c737bbe5e8 string, __obf_80d8f8c3ba351a94 ...any) error {
	return __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.HSet(__obf_9b31d4a4d4adfff9, __obf_4aecf3c737bbe5e8, __obf_80d8f8c3ba351a94...).Err()
}

func (__obf_2d21271215a2e06f *Redis) HGet(__obf_9b31d4a4d4adfff9 context.Context, __obf_4aecf3c737bbe5e8, __obf_047ec2d0df1dcd44 string) error {
	return __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.HGet(__obf_9b31d4a4d4adfff9, __obf_4aecf3c737bbe5e8, __obf_047ec2d0df1dcd44).Err()
}

func (__obf_2d21271215a2e06f *Redis) HGetAll(__obf_9b31d4a4d4adfff9 context.Context, __obf_4aecf3c737bbe5e8 string) (map[string]string, error) {
	return __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.HGetAll(__obf_9b31d4a4d4adfff9, __obf_4aecf3c737bbe5e8).Result()
}

func (__obf_2d21271215a2e06f *Redis) TxPipeline() redis.Pipeliner {
	return __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.TxPipeline()
}

func (__obf_2d21271215a2e06f *Redis) HDel(__obf_9b31d4a4d4adfff9 context.Context, __obf_4aecf3c737bbe5e8, __obf_154f45380e7a7f70 string) error {
	return __obf_2d21271215a2e06f.__obf_bb6a956ebc49aa20.HDel(__obf_9b31d4a4d4adfff9, __obf_4aecf3c737bbe5e8, __obf_154f45380e7a7f70).Err()
}
