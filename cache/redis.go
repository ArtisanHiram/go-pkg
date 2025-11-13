package __obf_f4d8558b35981340

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/redis/go-redis/v9"
)

var __obf_6671ac5b52b0c724 = context.Background()

type Redis struct {
	__obf_4dd04dc312bc73e5 *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_b7caa38976d31343 RdOption) *Redis {
	return &Redis{__obf_4dd04dc312bc73e5: redis.NewClient(&redis.Options{
		Addr:            __obf_b7caa38976d31343.Host,
		Password:        __obf_b7caa38976d31343.Password,
		DB:              __obf_b7caa38976d31343.DB,
		MaxIdleConns:    __obf_b7caa38976d31343.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_b7caa38976d31343.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_ed249696e6d62c67 *Redis) Delete(__obf_1261d18b5941c6bd string) {
	__obf_f9e6675b21adb870, __obf_2d5e4b42e300bedd := __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.Keys(__obf_6671ac5b52b0c724, __obf_1261d18b5941c6bd).Result()
	if __obf_2d5e4b42e300bedd != nil {
		for _, __obf_f038738651cdb77b := range __obf_f9e6675b21adb870 {
			__obf_ed249696e6d62c67.Remove(__obf_f038738651cdb77b)
		}
	}
}

func (__obf_ed249696e6d62c67 *Redis) Add(__obf_6208471da03258dd string, __obf_20a54be727e70f3c any) error {
	return __obf_ed249696e6d62c67.Set(__obf_6208471da03258dd, __obf_20a54be727e70f3c, 0)
}

func (__obf_ed249696e6d62c67 *Redis) Set(__obf_6208471da03258dd string, __obf_20a54be727e70f3c any, __obf_1e67295112a2a62d time.Duration) error {
	if __obf_4153159a00c5e3bb, __obf_ed10eafa7bc79f22 := __obf_20a54be727e70f3c.([]byte); __obf_ed10eafa7bc79f22 {
		return __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.Set(__obf_6671ac5b52b0c724, __obf_6208471da03258dd, __obf_4153159a00c5e3bb, __obf_1e67295112a2a62d).Err()
	} else {
		__obf_f93ff3e37a3461c7, __obf_2d5e4b42e300bedd := util.AnyToBytes(__obf_20a54be727e70f3c)
		if __obf_2d5e4b42e300bedd != nil {
			return fmt.Errorf("serialize error: %s", __obf_2d5e4b42e300bedd)
		}
		return __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.Set(__obf_6671ac5b52b0c724, __obf_6208471da03258dd, __obf_f93ff3e37a3461c7, __obf_1e67295112a2a62d).Err()
	}
}

func (__obf_ed249696e6d62c67 *Redis) Get(__obf_6208471da03258dd string) ([]byte, error) {
	return __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.Get(__obf_6671ac5b52b0c724, __obf_6208471da03258dd).Bytes()
}

func (__obf_ed249696e6d62c67 *Redis) Remove(__obf_6208471da03258dd string) error {
	return __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.Del(__obf_6671ac5b52b0c724, __obf_6208471da03258dd).Err()
}

// ClearAll clears all cache in memcache.
func (__obf_ed249696e6d62c67 *Redis) Clear() error {
	__obf_4d64b34bc51708da := __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.FlushAll(__obf_6671ac5b52b0c724)
	return __obf_4d64b34bc51708da.Err()
}

func (__obf_ed249696e6d62c67 *Redis) Has(__obf_6208471da03258dd string) bool {
	return __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.Exists(__obf_6671ac5b52b0c724, __obf_6208471da03258dd).Val() == 1
}

func (__obf_ed249696e6d62c67 *Redis) Renew(__obf_6208471da03258dd string, __obf_1e67295112a2a62d time.Duration) error {
	__obf_20a54be727e70f3c, __obf_2d5e4b42e300bedd := __obf_ed249696e6d62c67.Get(__obf_6208471da03258dd)
	if __obf_2d5e4b42e300bedd != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_6208471da03258dd, __obf_2d5e4b42e300bedd)
	}
	return __obf_ed249696e6d62c67.Set(__obf_6208471da03258dd, __obf_20a54be727e70f3c, __obf_1e67295112a2a62d)
}

func (__obf_ed249696e6d62c67 *Redis) HSet(__obf_6671ac5b52b0c724 context.Context, __obf_6208471da03258dd string, __obf_c2a58d61cdf2c64d ...any) error {
	return __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.HSet(__obf_6671ac5b52b0c724, __obf_6208471da03258dd, __obf_c2a58d61cdf2c64d...).Err()
}

func (__obf_ed249696e6d62c67 *Redis) HGet(__obf_6671ac5b52b0c724 context.Context, __obf_6208471da03258dd, __obf_d837b1d0099f9abf string) error {
	return __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.HGet(__obf_6671ac5b52b0c724, __obf_6208471da03258dd, __obf_d837b1d0099f9abf).Err()
}

func (__obf_ed249696e6d62c67 *Redis) HGetAll(__obf_6671ac5b52b0c724 context.Context, __obf_6208471da03258dd string) (map[string]string, error) {
	return __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.HGetAll(__obf_6671ac5b52b0c724, __obf_6208471da03258dd).Result()
}

func (__obf_ed249696e6d62c67 *Redis) TxPipeline() redis.Pipeliner {
	return __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.TxPipeline()
}

func (__obf_ed249696e6d62c67 *Redis) HDel(__obf_6671ac5b52b0c724 context.Context, __obf_6208471da03258dd, __obf_00b319ccfe0d17c5 string) error {
	return __obf_ed249696e6d62c67.__obf_4dd04dc312bc73e5.HDel(__obf_6671ac5b52b0c724, __obf_6208471da03258dd, __obf_00b319ccfe0d17c5).Err()
}
