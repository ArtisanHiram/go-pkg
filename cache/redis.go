package __obf_72d962f6a40bc95f

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/redis/go-redis/v9"
)

var __obf_c5de9c245d72f9ef = context.Background()

type Redis struct {
	__obf_05fb9187a6182d24 *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_eb2ab5af97863704 RdOption) *Redis {
	return &Redis{__obf_05fb9187a6182d24: redis.NewClient(&redis.Options{
		Addr:            __obf_eb2ab5af97863704.Host,
		Password:        __obf_eb2ab5af97863704.Password,
		DB:              __obf_eb2ab5af97863704.DB,
		MaxIdleConns:    __obf_eb2ab5af97863704.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_eb2ab5af97863704.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_d20240e4761bf874 *Redis) Delete(__obf_939173229b00ad97 string) {
	__obf_9c9e5eba310dd965, __obf_2498adaec5f4c8f1 := __obf_d20240e4761bf874.__obf_05fb9187a6182d24.Keys(__obf_c5de9c245d72f9ef, __obf_939173229b00ad97).Result()
	if __obf_2498adaec5f4c8f1 != nil {
		for _, __obf_41bce9a58f43c231 := range __obf_9c9e5eba310dd965 {
			__obf_d20240e4761bf874.
				Remove(__obf_41bce9a58f43c231)
		}
	}
}

func (__obf_d20240e4761bf874 *Redis) Add(__obf_73cb8148cbf55699 string, __obf_1dbdf9c3e13789b0 any) error {
	return __obf_d20240e4761bf874.Set(__obf_73cb8148cbf55699, __obf_1dbdf9c3e13789b0, 0)
}

func (__obf_d20240e4761bf874 *Redis) Set(__obf_73cb8148cbf55699 string, __obf_1dbdf9c3e13789b0 any, __obf_c6c24327f124e623 time.Duration) error {
	if __obf_b33efcc537be1290, __obf_97faa3bbfe732802 := __obf_1dbdf9c3e13789b0.([]byte); __obf_97faa3bbfe732802 {
		return __obf_d20240e4761bf874.__obf_05fb9187a6182d24.Set(__obf_c5de9c245d72f9ef, __obf_73cb8148cbf55699, __obf_b33efcc537be1290, __obf_c6c24327f124e623).Err()
	} else {
		__obf_a75c0778d463b37a, __obf_2498adaec5f4c8f1 := util.AnyToBytes(__obf_1dbdf9c3e13789b0)
		if __obf_2498adaec5f4c8f1 != nil {
			return fmt.Errorf("serialize error: %s", __obf_2498adaec5f4c8f1)
		}
		return __obf_d20240e4761bf874.__obf_05fb9187a6182d24.Set(__obf_c5de9c245d72f9ef, __obf_73cb8148cbf55699, __obf_a75c0778d463b37a, __obf_c6c24327f124e623).Err()
	}
}

func (__obf_d20240e4761bf874 *Redis) Get(__obf_73cb8148cbf55699 string) ([]byte, error) {
	return __obf_d20240e4761bf874.__obf_05fb9187a6182d24.Get(__obf_c5de9c245d72f9ef, __obf_73cb8148cbf55699).Bytes()
}

func (__obf_d20240e4761bf874 *Redis) Remove(__obf_73cb8148cbf55699 string) error {
	return __obf_d20240e4761bf874.__obf_05fb9187a6182d24.Del(__obf_c5de9c245d72f9ef,

		// ClearAll clears all cache in memcache.
		__obf_73cb8148cbf55699).Err()
}

func (__obf_d20240e4761bf874 *Redis) Clear() error {
	__obf_d42a9567d5e56708 := __obf_d20240e4761bf874.__obf_05fb9187a6182d24.FlushAll(__obf_c5de9c245d72f9ef)
	return __obf_d42a9567d5e56708.Err()
}

func (__obf_d20240e4761bf874 *Redis) Has(__obf_73cb8148cbf55699 string) bool {
	return __obf_d20240e4761bf874.__obf_05fb9187a6182d24.Exists(__obf_c5de9c245d72f9ef, __obf_73cb8148cbf55699).Val() == 1
}

func (__obf_d20240e4761bf874 *Redis) Renew(__obf_73cb8148cbf55699 string, __obf_c6c24327f124e623 time.Duration) error {
	__obf_1dbdf9c3e13789b0, __obf_2498adaec5f4c8f1 := __obf_d20240e4761bf874.Get(__obf_73cb8148cbf55699)
	if __obf_2498adaec5f4c8f1 != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_73cb8148cbf55699, __obf_2498adaec5f4c8f1)
	}
	return __obf_d20240e4761bf874.Set(__obf_73cb8148cbf55699, __obf_1dbdf9c3e13789b0, __obf_c6c24327f124e623)
}

func (__obf_d20240e4761bf874 *Redis) HSet(__obf_c5de9c245d72f9ef context.Context, __obf_73cb8148cbf55699 string, __obf_e7d3ec0f2282d56d ...any) error {
	return __obf_d20240e4761bf874.__obf_05fb9187a6182d24.HSet(__obf_c5de9c245d72f9ef, __obf_73cb8148cbf55699, __obf_e7d3ec0f2282d56d...).Err()
}

func (__obf_d20240e4761bf874 *Redis) HGet(__obf_c5de9c245d72f9ef context.Context, __obf_73cb8148cbf55699, __obf_efe0d1706e03f1dc string) error {
	return __obf_d20240e4761bf874.__obf_05fb9187a6182d24.HGet(__obf_c5de9c245d72f9ef, __obf_73cb8148cbf55699, __obf_efe0d1706e03f1dc).Err()
}

func (__obf_d20240e4761bf874 *Redis) HGetAll(__obf_c5de9c245d72f9ef context.Context, __obf_73cb8148cbf55699 string) (map[string]string, error) {
	return __obf_d20240e4761bf874.__obf_05fb9187a6182d24.HGetAll(__obf_c5de9c245d72f9ef, __obf_73cb8148cbf55699).Result()
}

func (__obf_d20240e4761bf874 *Redis) TxPipeline() redis.Pipeliner {
	return __obf_d20240e4761bf874.__obf_05fb9187a6182d24.TxPipeline()
}

func (__obf_d20240e4761bf874 *Redis) HDel(__obf_c5de9c245d72f9ef context.Context, __obf_73cb8148cbf55699, __obf_d422b2de8ed53c8f string) error {
	return __obf_d20240e4761bf874.__obf_05fb9187a6182d24.HDel(__obf_c5de9c245d72f9ef, __obf_73cb8148cbf55699, __obf_d422b2de8ed53c8f).Err()
}
