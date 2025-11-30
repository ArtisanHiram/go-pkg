package __obf_9e1ee87c6b054458

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/redis/go-redis/v9"
)

var __obf_587fbfd39d5c3d3d = context.Background()

type Redis struct {
	__obf_8d6b225703446deb *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_b1eb39bec362723a RdOption) *Redis {
	return &Redis{__obf_8d6b225703446deb: redis.NewClient(&redis.Options{
		Addr:            __obf_b1eb39bec362723a.Host,
		Password:        __obf_b1eb39bec362723a.Password,
		DB:              __obf_b1eb39bec362723a.DB,
		MaxIdleConns:    __obf_b1eb39bec362723a.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_b1eb39bec362723a.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_171f9f68d8286408 *Redis) Delete(__obf_14a292fb641d1b3d string) {
	__obf_361ee2f90c430d53, __obf_498673050542660c := __obf_171f9f68d8286408.__obf_8d6b225703446deb.Keys(__obf_587fbfd39d5c3d3d, __obf_14a292fb641d1b3d).Result()
	if __obf_498673050542660c != nil {
		for _, __obf_ffc285490d834437 := range __obf_361ee2f90c430d53 {
			__obf_171f9f68d8286408.
				Remove(__obf_ffc285490d834437)
		}
	}
}

func (__obf_171f9f68d8286408 *Redis) Add(__obf_3c13197612c6b39f string, __obf_7d0a67e130b6e9b4 any) error {
	return __obf_171f9f68d8286408.Set(__obf_3c13197612c6b39f, __obf_7d0a67e130b6e9b4, 0)
}

func (__obf_171f9f68d8286408 *Redis) Set(__obf_3c13197612c6b39f string, __obf_7d0a67e130b6e9b4 any, __obf_5a47a67359bdcb92 time.Duration) error {
	if __obf_aa24543fddec7b80, __obf_40eed10588aa6cec := __obf_7d0a67e130b6e9b4.([]byte); __obf_40eed10588aa6cec {
		return __obf_171f9f68d8286408.__obf_8d6b225703446deb.Set(__obf_587fbfd39d5c3d3d, __obf_3c13197612c6b39f, __obf_aa24543fddec7b80, __obf_5a47a67359bdcb92).Err()
	} else {
		__obf_c08337b0eb1cc1cf, __obf_498673050542660c := util.AnyToBytes(__obf_7d0a67e130b6e9b4)
		if __obf_498673050542660c != nil {
			return fmt.Errorf("serialize error: %s", __obf_498673050542660c)
		}
		return __obf_171f9f68d8286408.__obf_8d6b225703446deb.Set(__obf_587fbfd39d5c3d3d, __obf_3c13197612c6b39f, __obf_c08337b0eb1cc1cf, __obf_5a47a67359bdcb92).Err()
	}
}

func (__obf_171f9f68d8286408 *Redis) Get(__obf_3c13197612c6b39f string) ([]byte, error) {
	return __obf_171f9f68d8286408.__obf_8d6b225703446deb.Get(__obf_587fbfd39d5c3d3d, __obf_3c13197612c6b39f).Bytes()
}

func (__obf_171f9f68d8286408 *Redis) Remove(__obf_3c13197612c6b39f string) error {
	return __obf_171f9f68d8286408.__obf_8d6b225703446deb.Del(__obf_587fbfd39d5c3d3d,

		// ClearAll clears all cache in memcache.
		__obf_3c13197612c6b39f).Err()
}

func (__obf_171f9f68d8286408 *Redis) Clear() error {
	__obf_9b87ab3d195e3899 := __obf_171f9f68d8286408.__obf_8d6b225703446deb.FlushAll(__obf_587fbfd39d5c3d3d)
	return __obf_9b87ab3d195e3899.Err()
}

func (__obf_171f9f68d8286408 *Redis) Has(__obf_3c13197612c6b39f string) bool {
	return __obf_171f9f68d8286408.__obf_8d6b225703446deb.Exists(__obf_587fbfd39d5c3d3d, __obf_3c13197612c6b39f).Val() == 1
}

func (__obf_171f9f68d8286408 *Redis) Renew(__obf_3c13197612c6b39f string, __obf_5a47a67359bdcb92 time.Duration) error {
	__obf_7d0a67e130b6e9b4, __obf_498673050542660c := __obf_171f9f68d8286408.Get(__obf_3c13197612c6b39f)
	if __obf_498673050542660c != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_3c13197612c6b39f, __obf_498673050542660c)
	}
	return __obf_171f9f68d8286408.Set(__obf_3c13197612c6b39f, __obf_7d0a67e130b6e9b4, __obf_5a47a67359bdcb92)
}

func (__obf_171f9f68d8286408 *Redis) HSet(__obf_587fbfd39d5c3d3d context.Context, __obf_3c13197612c6b39f string, __obf_2238e5c9d845bc31 ...any) error {
	return __obf_171f9f68d8286408.__obf_8d6b225703446deb.HSet(__obf_587fbfd39d5c3d3d, __obf_3c13197612c6b39f, __obf_2238e5c9d845bc31...).Err()
}

func (__obf_171f9f68d8286408 *Redis) HGet(__obf_587fbfd39d5c3d3d context.Context, __obf_3c13197612c6b39f, __obf_8e77b02a8c9cdaf2 string) error {
	return __obf_171f9f68d8286408.__obf_8d6b225703446deb.HGet(__obf_587fbfd39d5c3d3d, __obf_3c13197612c6b39f, __obf_8e77b02a8c9cdaf2).Err()
}

func (__obf_171f9f68d8286408 *Redis) HGetAll(__obf_587fbfd39d5c3d3d context.Context, __obf_3c13197612c6b39f string) (map[string]string, error) {
	return __obf_171f9f68d8286408.__obf_8d6b225703446deb.HGetAll(__obf_587fbfd39d5c3d3d, __obf_3c13197612c6b39f).Result()
}

func (__obf_171f9f68d8286408 *Redis) TxPipeline() redis.Pipeliner {
	return __obf_171f9f68d8286408.__obf_8d6b225703446deb.TxPipeline()
}

func (__obf_171f9f68d8286408 *Redis) HDel(__obf_587fbfd39d5c3d3d context.Context, __obf_3c13197612c6b39f, __obf_d23c4c277bd011fe string) error {
	return __obf_171f9f68d8286408.__obf_8d6b225703446deb.HDel(__obf_587fbfd39d5c3d3d, __obf_3c13197612c6b39f, __obf_d23c4c277bd011fe).Err()
}
