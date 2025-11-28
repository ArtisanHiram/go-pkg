package __obf_79e7502d8563d901

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/redis/go-redis/v9"
)

var __obf_b6a9556406cbf421 = context.Background()

type Redis struct {
	__obf_52ee466a12dae311 *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_05ffe4ff8bbb9a88 RdOption) *Redis {
	return &Redis{__obf_52ee466a12dae311: redis.NewClient(&redis.Options{
		Addr:            __obf_05ffe4ff8bbb9a88.Host,
		Password:        __obf_05ffe4ff8bbb9a88.Password,
		DB:              __obf_05ffe4ff8bbb9a88.DB,
		MaxIdleConns:    __obf_05ffe4ff8bbb9a88.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_05ffe4ff8bbb9a88.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_3038774b5f5f3925 *Redis) Delete(__obf_bac226386ec983d7 string) {
	__obf_c9071e56d1cc094b, __obf_8d46af2525fab46a := __obf_3038774b5f5f3925.__obf_52ee466a12dae311.Keys(__obf_b6a9556406cbf421, __obf_bac226386ec983d7).Result()
	if __obf_8d46af2525fab46a != nil {
		for _, __obf_08ceb9044b83b8fd := range __obf_c9071e56d1cc094b {
			__obf_3038774b5f5f3925.Remove(__obf_08ceb9044b83b8fd)
		}
	}
}

func (__obf_3038774b5f5f3925 *Redis) Add(__obf_50994613b7653a88 string, __obf_e043381e0771feca any) error {
	return __obf_3038774b5f5f3925.Set(__obf_50994613b7653a88, __obf_e043381e0771feca, 0)
}

func (__obf_3038774b5f5f3925 *Redis) Set(__obf_50994613b7653a88 string, __obf_e043381e0771feca any, __obf_4ff6d8752a4fea92 time.Duration) error {
	if __obf_bfb4c8a1c77523af, __obf_7500628aeb1f47ab := __obf_e043381e0771feca.([]byte); __obf_7500628aeb1f47ab {
		return __obf_3038774b5f5f3925.__obf_52ee466a12dae311.Set(__obf_b6a9556406cbf421, __obf_50994613b7653a88, __obf_bfb4c8a1c77523af, __obf_4ff6d8752a4fea92).Err()
	} else {
		__obf_f39a48d17f9e52dc, __obf_8d46af2525fab46a := util.AnyToBytes(__obf_e043381e0771feca)
		if __obf_8d46af2525fab46a != nil {
			return fmt.Errorf("serialize error: %s", __obf_8d46af2525fab46a)
		}
		return __obf_3038774b5f5f3925.__obf_52ee466a12dae311.Set(__obf_b6a9556406cbf421, __obf_50994613b7653a88, __obf_f39a48d17f9e52dc, __obf_4ff6d8752a4fea92).Err()
	}
}

func (__obf_3038774b5f5f3925 *Redis) Get(__obf_50994613b7653a88 string) ([]byte, error) {
	return __obf_3038774b5f5f3925.__obf_52ee466a12dae311.Get(__obf_b6a9556406cbf421, __obf_50994613b7653a88).Bytes()
}

func (__obf_3038774b5f5f3925 *Redis) Remove(__obf_50994613b7653a88 string) error {
	return __obf_3038774b5f5f3925.__obf_52ee466a12dae311.Del(__obf_b6a9556406cbf421, __obf_50994613b7653a88).Err()
}

// ClearAll clears all cache in memcache.
func (__obf_3038774b5f5f3925 *Redis) Clear() error {
	__obf_9e1420876a56fd96 := __obf_3038774b5f5f3925.__obf_52ee466a12dae311.FlushAll(__obf_b6a9556406cbf421)
	return __obf_9e1420876a56fd96.Err()
}

func (__obf_3038774b5f5f3925 *Redis) Has(__obf_50994613b7653a88 string) bool {
	return __obf_3038774b5f5f3925.__obf_52ee466a12dae311.Exists(__obf_b6a9556406cbf421, __obf_50994613b7653a88).Val() == 1
}

func (__obf_3038774b5f5f3925 *Redis) Renew(__obf_50994613b7653a88 string, __obf_4ff6d8752a4fea92 time.Duration) error {
	__obf_e043381e0771feca, __obf_8d46af2525fab46a := __obf_3038774b5f5f3925.Get(__obf_50994613b7653a88)
	if __obf_8d46af2525fab46a != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_50994613b7653a88, __obf_8d46af2525fab46a)
	}
	return __obf_3038774b5f5f3925.Set(__obf_50994613b7653a88, __obf_e043381e0771feca, __obf_4ff6d8752a4fea92)
}

func (__obf_3038774b5f5f3925 *Redis) HSet(__obf_b6a9556406cbf421 context.Context, __obf_50994613b7653a88 string, __obf_9dfee4d60b382155 ...any) error {
	return __obf_3038774b5f5f3925.__obf_52ee466a12dae311.HSet(__obf_b6a9556406cbf421, __obf_50994613b7653a88, __obf_9dfee4d60b382155...).Err()
}

func (__obf_3038774b5f5f3925 *Redis) HGet(__obf_b6a9556406cbf421 context.Context, __obf_50994613b7653a88, __obf_0160bad1f35f16be string) error {
	return __obf_3038774b5f5f3925.__obf_52ee466a12dae311.HGet(__obf_b6a9556406cbf421, __obf_50994613b7653a88, __obf_0160bad1f35f16be).Err()
}

func (__obf_3038774b5f5f3925 *Redis) HGetAll(__obf_b6a9556406cbf421 context.Context, __obf_50994613b7653a88 string) (map[string]string, error) {
	return __obf_3038774b5f5f3925.__obf_52ee466a12dae311.HGetAll(__obf_b6a9556406cbf421, __obf_50994613b7653a88).Result()
}

func (__obf_3038774b5f5f3925 *Redis) TxPipeline() redis.Pipeliner {
	return __obf_3038774b5f5f3925.__obf_52ee466a12dae311.TxPipeline()
}

func (__obf_3038774b5f5f3925 *Redis) HDel(__obf_b6a9556406cbf421 context.Context, __obf_50994613b7653a88, __obf_1809595297598689 string) error {
	return __obf_3038774b5f5f3925.__obf_52ee466a12dae311.HDel(__obf_b6a9556406cbf421, __obf_50994613b7653a88, __obf_1809595297598689).Err()
}
