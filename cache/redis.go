package __obf_6b306490bf7221d3

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/redis/go-redis/v9"
)

var __obf_2fd2ea48837086f1 = context.Background()

type Redis struct {
	__obf_e0fb86cbd4784ac3 *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_75e048f42715e227 RdOption) *Redis {
	return &Redis{__obf_e0fb86cbd4784ac3: redis.NewClient(&redis.Options{
		Addr:            __obf_75e048f42715e227.Host,
		Password:        __obf_75e048f42715e227.Password,
		DB:              __obf_75e048f42715e227.DB,
		MaxIdleConns:    __obf_75e048f42715e227.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_75e048f42715e227.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_3a5bbe1f7e27c46d *Redis) Delete(__obf_7e6d8e57607bcb1d string) {
	__obf_71b9221f67592dda, __obf_9ffdd5bbb9f1dc61 := __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.Keys(__obf_2fd2ea48837086f1, __obf_7e6d8e57607bcb1d).Result()
	if __obf_9ffdd5bbb9f1dc61 != nil {
		for _, __obf_afff83fa0e852ce0 := range __obf_71b9221f67592dda {
			__obf_3a5bbe1f7e27c46d.
				Remove(__obf_afff83fa0e852ce0)
		}
	}
}

func (__obf_3a5bbe1f7e27c46d *Redis) Add(__obf_fa3a380c35ada5d9 string, __obf_17a4af14df8bae4f any) error {
	return __obf_3a5bbe1f7e27c46d.Set(__obf_fa3a380c35ada5d9, __obf_17a4af14df8bae4f, 0)
}

func (__obf_3a5bbe1f7e27c46d *Redis) Set(__obf_fa3a380c35ada5d9 string, __obf_17a4af14df8bae4f any, __obf_a15e4803e6944fc5 time.Duration) error {
	if __obf_6208501794a51081, __obf_8761a3e94b6603ac := __obf_17a4af14df8bae4f.([]byte); __obf_8761a3e94b6603ac {
		return __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.Set(__obf_2fd2ea48837086f1, __obf_fa3a380c35ada5d9, __obf_6208501794a51081, __obf_a15e4803e6944fc5).Err()
	} else {
		__obf_53e4afb5fabb062f, __obf_9ffdd5bbb9f1dc61 := util.AnyToBytes(__obf_17a4af14df8bae4f)
		if __obf_9ffdd5bbb9f1dc61 != nil {
			return fmt.Errorf("serialize error: %s", __obf_9ffdd5bbb9f1dc61)
		}
		return __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.Set(__obf_2fd2ea48837086f1, __obf_fa3a380c35ada5d9, __obf_53e4afb5fabb062f, __obf_a15e4803e6944fc5).Err()
	}
}

func (__obf_3a5bbe1f7e27c46d *Redis) Get(__obf_fa3a380c35ada5d9 string) ([]byte, error) {
	return __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.Get(__obf_2fd2ea48837086f1, __obf_fa3a380c35ada5d9).Bytes()
}

func (__obf_3a5bbe1f7e27c46d *Redis) Remove(__obf_fa3a380c35ada5d9 string) error {
	return __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.Del(__obf_2fd2ea48837086f1,

		// ClearAll clears all cache in memcache.
		__obf_fa3a380c35ada5d9).Err()
}

func (__obf_3a5bbe1f7e27c46d *Redis) Clear() error {
	__obf_6527c64e373af750 := __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.FlushAll(__obf_2fd2ea48837086f1)
	return __obf_6527c64e373af750.Err()
}

func (__obf_3a5bbe1f7e27c46d *Redis) Has(__obf_fa3a380c35ada5d9 string) bool {
	return __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.Exists(__obf_2fd2ea48837086f1, __obf_fa3a380c35ada5d9).Val() == 1
}

func (__obf_3a5bbe1f7e27c46d *Redis) Renew(__obf_fa3a380c35ada5d9 string, __obf_a15e4803e6944fc5 time.Duration) error {
	__obf_17a4af14df8bae4f, __obf_9ffdd5bbb9f1dc61 := __obf_3a5bbe1f7e27c46d.Get(__obf_fa3a380c35ada5d9)
	if __obf_9ffdd5bbb9f1dc61 != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_fa3a380c35ada5d9, __obf_9ffdd5bbb9f1dc61)
	}
	return __obf_3a5bbe1f7e27c46d.Set(__obf_fa3a380c35ada5d9, __obf_17a4af14df8bae4f, __obf_a15e4803e6944fc5)
}

func (__obf_3a5bbe1f7e27c46d *Redis) HSet(__obf_2fd2ea48837086f1 context.Context, __obf_fa3a380c35ada5d9 string, __obf_d20fdd10a90faa06 ...any) error {
	return __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.HSet(__obf_2fd2ea48837086f1, __obf_fa3a380c35ada5d9, __obf_d20fdd10a90faa06...).Err()
}

func (__obf_3a5bbe1f7e27c46d *Redis) HGet(__obf_2fd2ea48837086f1 context.Context, __obf_fa3a380c35ada5d9, __obf_691b0a38ff587b34 string) error {
	return __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.HGet(__obf_2fd2ea48837086f1, __obf_fa3a380c35ada5d9, __obf_691b0a38ff587b34).Err()
}

func (__obf_3a5bbe1f7e27c46d *Redis) HGetAll(__obf_2fd2ea48837086f1 context.Context, __obf_fa3a380c35ada5d9 string) (map[string]string, error) {
	return __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.HGetAll(__obf_2fd2ea48837086f1, __obf_fa3a380c35ada5d9).Result()
}

func (__obf_3a5bbe1f7e27c46d *Redis) TxPipeline() redis.Pipeliner {
	return __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.TxPipeline()
}

func (__obf_3a5bbe1f7e27c46d *Redis) HDel(__obf_2fd2ea48837086f1 context.Context, __obf_fa3a380c35ada5d9, __obf_fd25b7b452cc8474 string) error {
	return __obf_3a5bbe1f7e27c46d.__obf_e0fb86cbd4784ac3.HDel(__obf_2fd2ea48837086f1, __obf_fa3a380c35ada5d9, __obf_fd25b7b452cc8474).Err()
}
