package __obf_8281010a3a6d2ab0

import (
	"context"
	"fmt"

	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/redis/go-redis/v9"
)

var __obf_e4e530f1d4f267bf = context.Background()

type Redis struct {
	__obf_b3f92554e3dc80e1 *redis.Client
}

type RdOption struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"auth"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max-idle"`
	IdleTimeout int    `yaml:"idle-timeout"`
}

func NewRedis(__obf_1b32cb33dbcfe9e7 RdOption) *Redis {
	return &Redis{__obf_b3f92554e3dc80e1: redis.NewClient(&redis.Options{
		Addr:            __obf_1b32cb33dbcfe9e7.Host,
		Password:        __obf_1b32cb33dbcfe9e7.Password,
		DB:              __obf_1b32cb33dbcfe9e7.DB,
		MaxIdleConns:    __obf_1b32cb33dbcfe9e7.MaxIdle,
		ConnMaxIdleTime: time.Duration(__obf_1b32cb33dbcfe9e7.IdleTimeout) * time.Minute,
	})}
}

// func (r *Redis) Keys(key string) ([]string, error) {
// 	return r.cli.Keys(ctx, key).Result()
// }

func (__obf_bcf0de439a6c6928 *Redis) Delete(__obf_c2739d84a72f2e49 string) {
	__obf_8d5c4673989a442a, __obf_fd03130c6793cb3b := __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.Keys(__obf_e4e530f1d4f267bf, __obf_c2739d84a72f2e49).Result()
	if __obf_fd03130c6793cb3b != nil {
		for _, __obf_f4db856e39cfcf16 := range __obf_8d5c4673989a442a {
			__obf_bcf0de439a6c6928.Remove(__obf_f4db856e39cfcf16)
		}
	}
}

func (__obf_bcf0de439a6c6928 *Redis) Add(__obf_805b9182f4a01a43 string, __obf_a31843a6764aecf9 any) error {
	return __obf_bcf0de439a6c6928.Set(__obf_805b9182f4a01a43, __obf_a31843a6764aecf9, 0)
}

func (__obf_bcf0de439a6c6928 *Redis) Set(__obf_805b9182f4a01a43 string, __obf_a31843a6764aecf9 any, __obf_e2e393700839b297 time.Duration) error {
	if __obf_51fc6ca2aca54eed, __obf_102da286c8b8a571 := __obf_a31843a6764aecf9.([]byte); __obf_102da286c8b8a571 {
		return __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.Set(__obf_e4e530f1d4f267bf, __obf_805b9182f4a01a43, __obf_51fc6ca2aca54eed, __obf_e2e393700839b297).Err()
	} else {
		__obf_fff87ad6ee2b8d0c, __obf_fd03130c6793cb3b := util.AnyToBytes(__obf_a31843a6764aecf9)
		if __obf_fd03130c6793cb3b != nil {
			return fmt.Errorf("serialize error: %s", __obf_fd03130c6793cb3b)
		}
		return __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.Set(__obf_e4e530f1d4f267bf, __obf_805b9182f4a01a43, __obf_fff87ad6ee2b8d0c, __obf_e2e393700839b297).Err()
	}
}

func (__obf_bcf0de439a6c6928 *Redis) Get(__obf_805b9182f4a01a43 string) ([]byte, error) {
	return __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.Get(__obf_e4e530f1d4f267bf, __obf_805b9182f4a01a43).Bytes()
}

func (__obf_bcf0de439a6c6928 *Redis) Remove(__obf_805b9182f4a01a43 string) error {
	return __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.Del(__obf_e4e530f1d4f267bf, __obf_805b9182f4a01a43).Err()
}

// ClearAll clears all cache in memcache.
func (__obf_bcf0de439a6c6928 *Redis) Clear() error {
	__obf_ac128080011a3a76 := __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.FlushAll(__obf_e4e530f1d4f267bf)
	return __obf_ac128080011a3a76.Err()
}

func (__obf_bcf0de439a6c6928 *Redis) Has(__obf_805b9182f4a01a43 string) bool {
	return __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.Exists(__obf_e4e530f1d4f267bf, __obf_805b9182f4a01a43).Val() == 1
}

func (__obf_bcf0de439a6c6928 *Redis) Renew(__obf_805b9182f4a01a43 string, __obf_e2e393700839b297 time.Duration) error {
	__obf_a31843a6764aecf9, __obf_fd03130c6793cb3b := __obf_bcf0de439a6c6928.Get(__obf_805b9182f4a01a43)
	if __obf_fd03130c6793cb3b != nil {
		return fmt.Errorf("'%s' not exists: %s", __obf_805b9182f4a01a43, __obf_fd03130c6793cb3b)
	}
	return __obf_bcf0de439a6c6928.Set(__obf_805b9182f4a01a43, __obf_a31843a6764aecf9, __obf_e2e393700839b297)
}

func (__obf_bcf0de439a6c6928 *Redis) HSet(__obf_e4e530f1d4f267bf context.Context, __obf_805b9182f4a01a43 string, __obf_5d8b5d5365fdd31e ...any) error {
	return __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.HSet(__obf_e4e530f1d4f267bf, __obf_805b9182f4a01a43, __obf_5d8b5d5365fdd31e...).Err()
}

func (__obf_bcf0de439a6c6928 *Redis) HGet(__obf_e4e530f1d4f267bf context.Context, __obf_805b9182f4a01a43, __obf_50bb432740f74fc9 string) error {
	return __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.HGet(__obf_e4e530f1d4f267bf, __obf_805b9182f4a01a43, __obf_50bb432740f74fc9).Err()
}

func (__obf_bcf0de439a6c6928 *Redis) HGetAll(__obf_e4e530f1d4f267bf context.Context, __obf_805b9182f4a01a43 string) (map[string]string, error) {
	return __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.HGetAll(__obf_e4e530f1d4f267bf, __obf_805b9182f4a01a43).Result()
}

func (__obf_bcf0de439a6c6928 *Redis) TxPipeline() redis.Pipeliner {
	return __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.TxPipeline()
}

func (__obf_bcf0de439a6c6928 *Redis) HDel(__obf_e4e530f1d4f267bf context.Context, __obf_805b9182f4a01a43, __obf_23a67683d7195dae string) error {
	return __obf_bcf0de439a6c6928.__obf_b3f92554e3dc80e1.HDel(__obf_e4e530f1d4f267bf, __obf_805b9182f4a01a43, __obf_23a67683d7195dae).Err()
}
