package __obf_2f5c14e012cec42e

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/redis/go-redis/v9"
)

var __obf_00d0ac22bf72f865 = context.Background()

type Redis struct {
	__obf_fcdc75f6519255b2 *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_e5cf37a08b98986e RdOption) *Redis {
	return &Redis{__obf_fcdc75f6519255b2: redis.NewClient(&redis.Options{
		Addr:            __obf_e5cf37a08b98986e.Host,
		Password:        __obf_e5cf37a08b98986e.Password,
		DB:              __obf_e5cf37a08b98986e.DB,
		MaxIdleConns:    __obf_e5cf37a08b98986e.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_e5cf37a08b98986e.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_79efa7595252dd83 *Redis) Delete(__obf_68047b361b381cb2 string) {
	__obf_b0c0c10fcec6ff52, __obf_956c4015cf7da152 := __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.Keys(__obf_00d0ac22bf72f865, __obf_68047b361b381cb2).Result()
	if __obf_956c4015cf7da152 != nil {
		for _, __obf_ffd22490bdfba93e := range __obf_b0c0c10fcec6ff52 {
			__obf_79efa7595252dd83.Remove(__obf_ffd22490bdfba93e)
		}
	}
}

func (__obf_79efa7595252dd83 *Redis) Add(__obf_f0b3ebeba048b5e4 string, __obf_aee38dd09d94cdd5 any) error {
	return __obf_79efa7595252dd83.Set(__obf_f0b3ebeba048b5e4, __obf_aee38dd09d94cdd5, 0)
}

func (__obf_79efa7595252dd83 *Redis) Set(__obf_f0b3ebeba048b5e4 string, __obf_aee38dd09d94cdd5 any, __obf_6459304da6a62a40 time.Duration) error {
	if __obf_48da227d69fb34be, __obf_a39fb2a67596cf35 := __obf_aee38dd09d94cdd5.([]byte); __obf_a39fb2a67596cf35 {
		return __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.Set(__obf_00d0ac22bf72f865, __obf_f0b3ebeba048b5e4, __obf_48da227d69fb34be, __obf_6459304da6a62a40).Err()
	} else {
		__obf_297f986b313c8a07, __obf_956c4015cf7da152 := util.AnyToBytes(__obf_aee38dd09d94cdd5)
		if __obf_956c4015cf7da152 != nil {
			return fmt.Errorf("serialize error: %s", __obf_956c4015cf7da152)
		}
		return __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.Set(__obf_00d0ac22bf72f865, __obf_f0b3ebeba048b5e4, __obf_297f986b313c8a07, __obf_6459304da6a62a40).Err()
	}
}

func (__obf_79efa7595252dd83 *Redis) Get(__obf_f0b3ebeba048b5e4 string) ([]byte, error) {
	return __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.Get(__obf_00d0ac22bf72f865, __obf_f0b3ebeba048b5e4).Bytes()
}

func (__obf_79efa7595252dd83 *Redis) Remove(__obf_f0b3ebeba048b5e4 string) error {
	return __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.Del(__obf_00d0ac22bf72f865, __obf_f0b3ebeba048b5e4).Err()
}

// ClearAll clears all cache in memcache.
func (__obf_79efa7595252dd83 *Redis) Clear() error {
	__obf_387c4b2460e31a19 := __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.FlushAll(__obf_00d0ac22bf72f865)
	return __obf_387c4b2460e31a19.Err()
}

func (__obf_79efa7595252dd83 *Redis) Has(__obf_f0b3ebeba048b5e4 string) bool {
	return __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.Exists(__obf_00d0ac22bf72f865, __obf_f0b3ebeba048b5e4).Val() == 1
}

func (__obf_79efa7595252dd83 *Redis) Renew(__obf_f0b3ebeba048b5e4 string, __obf_6459304da6a62a40 time.Duration) error {
	__obf_aee38dd09d94cdd5, __obf_956c4015cf7da152 := __obf_79efa7595252dd83.Get(__obf_f0b3ebeba048b5e4)
	if __obf_956c4015cf7da152 != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_f0b3ebeba048b5e4, __obf_956c4015cf7da152)
	}
	return __obf_79efa7595252dd83.Set(__obf_f0b3ebeba048b5e4, __obf_aee38dd09d94cdd5, __obf_6459304da6a62a40)
}

func (__obf_79efa7595252dd83 *Redis) HSet(__obf_00d0ac22bf72f865 context.Context, __obf_f0b3ebeba048b5e4 string, __obf_678cb72ab649dfba ...any) error {
	return __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.HSet(__obf_00d0ac22bf72f865, __obf_f0b3ebeba048b5e4, __obf_678cb72ab649dfba...).Err()
}

func (__obf_79efa7595252dd83 *Redis) HGet(__obf_00d0ac22bf72f865 context.Context, __obf_f0b3ebeba048b5e4, __obf_778160f2edeece20 string) error {
	return __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.HGet(__obf_00d0ac22bf72f865, __obf_f0b3ebeba048b5e4, __obf_778160f2edeece20).Err()
}

func (__obf_79efa7595252dd83 *Redis) HGetAll(__obf_00d0ac22bf72f865 context.Context, __obf_f0b3ebeba048b5e4 string) (map[string]string, error) {
	return __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.HGetAll(__obf_00d0ac22bf72f865, __obf_f0b3ebeba048b5e4).Result()
}

func (__obf_79efa7595252dd83 *Redis) TxPipeline() redis.Pipeliner {
	return __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.TxPipeline()
}

func (__obf_79efa7595252dd83 *Redis) HDel(__obf_00d0ac22bf72f865 context.Context, __obf_f0b3ebeba048b5e4, __obf_05404e1a9e7a502e string) error {
	return __obf_79efa7595252dd83.__obf_fcdc75f6519255b2.HDel(__obf_00d0ac22bf72f865, __obf_f0b3ebeba048b5e4, __obf_05404e1a9e7a502e).Err()
}
