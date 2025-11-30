package __obf_3b8640e918b7e3ff

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/redis/go-redis/v9"
)

var __obf_5646268b50e6cca3 = context.Background()

type Redis struct {
	__obf_ba155ff6eaecf5d8 *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_379860bacbefcca8 RdOption) *Redis {
	return &Redis{__obf_ba155ff6eaecf5d8: redis.NewClient(&redis.Options{
		Addr:            __obf_379860bacbefcca8.Host,
		Password:        __obf_379860bacbefcca8.Password,
		DB:              __obf_379860bacbefcca8.DB,
		MaxIdleConns:    __obf_379860bacbefcca8.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_379860bacbefcca8.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_f3a80f3795c3e8d5 *Redis) Delete(__obf_4298344c079705aa string) {
	__obf_f3126ec576b86993, __obf_8c6b39ef87b4061f := __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.Keys(__obf_5646268b50e6cca3, __obf_4298344c079705aa).Result()
	if __obf_8c6b39ef87b4061f != nil {
		for _, __obf_ebdd4526b9b834d1 := range __obf_f3126ec576b86993 {
			__obf_f3a80f3795c3e8d5.
				Remove(__obf_ebdd4526b9b834d1)
		}
	}
}

func (__obf_f3a80f3795c3e8d5 *Redis) Add(__obf_3405c14f70aaa4d0 string, __obf_13458218654a7f13 any) error {
	return __obf_f3a80f3795c3e8d5.Set(__obf_3405c14f70aaa4d0, __obf_13458218654a7f13, 0)
}

func (__obf_f3a80f3795c3e8d5 *Redis) Set(__obf_3405c14f70aaa4d0 string, __obf_13458218654a7f13 any, __obf_433bae995cf2f1ae time.Duration) error {
	if __obf_fe4eccc509e2ba8c, __obf_cddbfb0aefdf4145 := __obf_13458218654a7f13.([]byte); __obf_cddbfb0aefdf4145 {
		return __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.Set(__obf_5646268b50e6cca3, __obf_3405c14f70aaa4d0, __obf_fe4eccc509e2ba8c, __obf_433bae995cf2f1ae).Err()
	} else {
		__obf_12b0b98285fac904, __obf_8c6b39ef87b4061f := util.AnyToBytes(__obf_13458218654a7f13)
		if __obf_8c6b39ef87b4061f != nil {
			return fmt.Errorf("serialize error: %s", __obf_8c6b39ef87b4061f)
		}
		return __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.Set(__obf_5646268b50e6cca3, __obf_3405c14f70aaa4d0, __obf_12b0b98285fac904, __obf_433bae995cf2f1ae).Err()
	}
}

func (__obf_f3a80f3795c3e8d5 *Redis) Get(__obf_3405c14f70aaa4d0 string) ([]byte, error) {
	return __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.Get(__obf_5646268b50e6cca3, __obf_3405c14f70aaa4d0).Bytes()
}

func (__obf_f3a80f3795c3e8d5 *Redis) Remove(__obf_3405c14f70aaa4d0 string) error {
	return __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.Del(__obf_5646268b50e6cca3,

		// ClearAll clears all cache in memcache.
		__obf_3405c14f70aaa4d0).Err()
}

func (__obf_f3a80f3795c3e8d5 *Redis) Clear() error {
	__obf_2148c66210a82787 := __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.FlushAll(__obf_5646268b50e6cca3)
	return __obf_2148c66210a82787.Err()
}

func (__obf_f3a80f3795c3e8d5 *Redis) Has(__obf_3405c14f70aaa4d0 string) bool {
	return __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.Exists(__obf_5646268b50e6cca3, __obf_3405c14f70aaa4d0).Val() == 1
}

func (__obf_f3a80f3795c3e8d5 *Redis) Renew(__obf_3405c14f70aaa4d0 string, __obf_433bae995cf2f1ae time.Duration) error {
	__obf_13458218654a7f13, __obf_8c6b39ef87b4061f := __obf_f3a80f3795c3e8d5.Get(__obf_3405c14f70aaa4d0)
	if __obf_8c6b39ef87b4061f != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_3405c14f70aaa4d0, __obf_8c6b39ef87b4061f)
	}
	return __obf_f3a80f3795c3e8d5.Set(__obf_3405c14f70aaa4d0, __obf_13458218654a7f13, __obf_433bae995cf2f1ae)
}

func (__obf_f3a80f3795c3e8d5 *Redis) HSet(__obf_5646268b50e6cca3 context.Context, __obf_3405c14f70aaa4d0 string, __obf_48493cfe7e66a294 ...any) error {
	return __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.HSet(__obf_5646268b50e6cca3, __obf_3405c14f70aaa4d0, __obf_48493cfe7e66a294...).Err()
}

func (__obf_f3a80f3795c3e8d5 *Redis) HGet(__obf_5646268b50e6cca3 context.Context, __obf_3405c14f70aaa4d0, __obf_931e6622ad04c5ff string) error {
	return __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.HGet(__obf_5646268b50e6cca3, __obf_3405c14f70aaa4d0, __obf_931e6622ad04c5ff).Err()
}

func (__obf_f3a80f3795c3e8d5 *Redis) HGetAll(__obf_5646268b50e6cca3 context.Context, __obf_3405c14f70aaa4d0 string) (map[string]string, error) {
	return __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.HGetAll(__obf_5646268b50e6cca3, __obf_3405c14f70aaa4d0).Result()
}

func (__obf_f3a80f3795c3e8d5 *Redis) TxPipeline() redis.Pipeliner {
	return __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.TxPipeline()
}

func (__obf_f3a80f3795c3e8d5 *Redis) HDel(__obf_5646268b50e6cca3 context.Context, __obf_3405c14f70aaa4d0, __obf_e733d110f0d170ba string) error {
	return __obf_f3a80f3795c3e8d5.__obf_ba155ff6eaecf5d8.HDel(__obf_5646268b50e6cca3, __obf_3405c14f70aaa4d0, __obf_e733d110f0d170ba).Err()
}
