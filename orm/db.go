package __obf_ea545e4bdf748fd2

import (
	"errors"
	"fmt"
	util "github.com/ArtisanHiram/go-pkg/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"reflect"
	"slices"
	"strings"
	"time"
)

type Options struct {
	Reserved struct {
		ID         string `yaml:"id"`
		NO         string `yaml:"no"`
		CreateTime string `yaml:"create-time"`
		UpdateTime string `yaml:"update-time"`
	} `yaml:"reserved"`
	Driver     string `yaml:"driver"`
	DataSource string `yaml:"datasource"`
	Object     struct {
		Table map[uint8]string `yaml:"table"`
		Query map[uint8]string `yaml:"query"`
	} `yaml:"object"`
	// Debug       bool   `yaml:"-"`
	Placeholder string `yaml:"-"`
}

type ORM struct {
	DB     *sqlx.DB
	Option *Options
}

func NewORM(__obf_79c864729015d67e *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_79c864729015d67e.Driver, __obf_79c864729015d67e.DataSource)

	switch __obf_79c864729015d67e.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_79c864729015d67e.
			Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_79c864729015d67e.
			Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_79c864729015d67e.
			Placeholder = ":"
	case "sqlserver":
		__obf_79c864729015d67e.
			Placeholder = "@"
	}

	if __obf_3220344ba911fc77 := DB.Ping(); __obf_3220344ba911fc77 != nil {
		panic(__obf_3220344ba911fc77)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_79c864729015d67e,
		}, func() {
			DB.Close()
		}
}

func (__obf_d74aacde370b8778 *ORM) ListMap(__obf_84494ca656481f28 *sqlx.Rows, __obf_cf04d29817b50981 func(map[string]any) (string, string)) (__obf_b0679735c3db1659 []map[string]any, __obf_3220344ba911fc77 error) {
	for __obf_84494ca656481f28.Next() {
		__obf_761281bbc31a730f := make(map[string]any)
		if __obf_3220344ba911fc77 = __obf_84494ca656481f28.MapScan(__obf_761281bbc31a730f); __obf_3220344ba911fc77 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_3220344ba911fc77)
		}
		for __obf_5304b9e5e1368ecf, __obf_8a2e311a616cdf4b := range __obf_761281bbc31a730f {
			switch __obf_8a2e311a616cdf4b := __obf_8a2e311a616cdf4b.(type) {
			case []uint8:
				__obf_761281bbc31a730f[__obf_5304b9e5e1368ecf] = string(__obf_8a2e311a616cdf4b)
			case int64:
				__obf_761281bbc31a730f[__obf_5304b9e5e1368ecf] = int64(__obf_8a2e311a616cdf4b)
			case uint64:
				__obf_761281bbc31a730f[__obf_5304b9e5e1368ecf] = uint64(__obf_8a2e311a616cdf4b)
			}
		}
		if __obf_cf04d29817b50981 != nil {
			__obf_dc86f67160aae392, __obf_a6490b2f4d4d2871 := __obf_cf04d29817b50981(__obf_761281bbc31a730f)
			__obf_761281bbc31a730f[__obf_dc86f67160aae392] = __obf_a6490b2f4d4d2871
		}
		__obf_b0679735c3db1659 = append(__obf_b0679735c3db1659, __obf_761281bbc31a730f)
	}
	return
}

func (__obf_d74aacde370b8778 *ORM) Exists(__obf_a70fdc9edb43b1fa, __obf_d1a4a9e58b41ed33 string, __obf_15a4a0514baf9e2c ...any) bool {
	var __obf_668cb41a9b069b90 bool
	_ = __obf_d74aacde370b8778.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_a70fdc9edb43b1fa, __obf_d1a4a9e58b41ed33), __obf_15a4a0514baf9e2c...).Scan(&__obf_668cb41a9b069b90)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_668cb41a9b069b90
}

func (__obf_d74aacde370b8778 *ORM) SaveModel(__obf_61e43608656dfcc0 any, __obf_525801f57a5d856f string) (string, error) {
	__obf_8a2e311a616cdf4b := reflect.ValueOf(__obf_61e43608656dfcc0)
	if __obf_8a2e311a616cdf4b.Kind() == reflect.Pointer {
		__obf_8a2e311a616cdf4b = __obf_8a2e311a616cdf4b.Elem()
	}
	var (
		__obf_d78ea253231b6890 = __obf_8a2e311a616cdf4b.
					FieldByName("Id")
		__obf_dadf54047061b513 = util.ToSnake(reflect.Indirect(__obf_8a2e311a616cdf4b).Type().Name())
		__obf_3c80b73d77602a41 string
		__obf_15a4a0514baf9e2c map[string]any
		__obf_3220344ba911fc77 error
	)

	if __obf_d78ea253231b6890.IsZero() {
		__obf_3c80b73d77602a41,
			// insert
			__obf_15a4a0514baf9e2c, __obf_3220344ba911fc77 = BuildInsertSQL(__obf_61e43608656dfcc0, __obf_dadf54047061b513)
		if __obf_3220344ba911fc77 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_3220344ba911fc77.Error())
		}
		if __obf_525801f57a5d856f == "" {
			__obf_525801f57a5d856f = util.StringUUID()
		}
		__obf_15a4a0514baf9e2c["id"] = __obf_525801f57a5d856f
		__obf_43dd99531253543c := time.Now().Unix()
		if _, __obf_fcb98d9ec4b8d44c := __obf_15a4a0514baf9e2c[__obf_d74aacde370b8778.Option.Reserved.CreateTime]; __obf_fcb98d9ec4b8d44c {
			__obf_15a4a0514baf9e2c[__obf_d74aacde370b8778.Option.Reserved.CreateTime] = __obf_43dd99531253543c
		}
		if _, __obf_fcb98d9ec4b8d44c := __obf_15a4a0514baf9e2c[__obf_d74aacde370b8778.Option.Reserved.UpdateTime]; __obf_fcb98d9ec4b8d44c {
			__obf_15a4a0514baf9e2c[__obf_d74aacde370b8778.Option.Reserved.UpdateTime] = __obf_43dd99531253543c
		}
	} else {
		__obf_3c80b73d77602a41,
			// update
			__obf_15a4a0514baf9e2c, __obf_3220344ba911fc77 = BuildUpdateSQL(__obf_61e43608656dfcc0, __obf_dadf54047061b513, __obf_d74aacde370b8778.Option.Reserved.UpdateTime, []string{__obf_d74aacde370b8778.Option.Reserved.ID}, __obf_d74aacde370b8778.Option.Reserved.CreateTime, __obf_d74aacde370b8778.Option.Reserved.NO)
		if __obf_3220344ba911fc77 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_3220344ba911fc77.Error())
		}
	}

	if __obf_3c80b73d77602a41 != "" {
		_, __obf_3220344ba911fc77 := __obf_d74aacde370b8778.DB.NamedExec(__obf_3c80b73d77602a41, __obf_15a4a0514baf9e2c)
		if __obf_3220344ba911fc77 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_3220344ba911fc77.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_3c80b73d77602a41, "args", __obf_15a4a0514baf9e2c)
	}

	return __obf_525801f57a5d856f, nil

}

func (__obf_d74aacde370b8778 *ORM) SaveModelWidthAutoID(__obf_61e43608656dfcc0 any, __obf_525801f57a5d856f int64) (int64, error) {
	__obf_8a2e311a616cdf4b := // query, args := modelBind(db.Option, m, insertId)
		reflect.ValueOf(__obf_61e43608656dfcc0)
	if __obf_8a2e311a616cdf4b.Kind() == reflect.Pointer {
		__obf_8a2e311a616cdf4b = __obf_8a2e311a616cdf4b.Elem()
	}
	__obf_d78ea253231b6890 := __obf_8a2e311a616cdf4b.FieldByName("Id")
	__obf_dadf54047061b513 := util.ToSnake(reflect.Indirect(__obf_8a2e311a616cdf4b).Type().Name())
	var (
		__obf_3c80b73d77602a41 string
		__obf_15a4a0514baf9e2c map[string]any
		__obf_3220344ba911fc77 error
	)

	if __obf_d78ea253231b6890.IsZero() {
		__obf_3c80b73d77602a41,
			// insert
			__obf_15a4a0514baf9e2c, __obf_3220344ba911fc77 = BuildInsertSQL(__obf_61e43608656dfcc0, __obf_dadf54047061b513)
		if __obf_3220344ba911fc77 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_3220344ba911fc77.Error())
		}
		if __obf_525801f57a5d856f != 0 {
			__obf_15a4a0514baf9e2c["id"] = __obf_525801f57a5d856f
		}
		__obf_43dd99531253543c := time.Now().Unix()
		if _, __obf_fcb98d9ec4b8d44c := __obf_15a4a0514baf9e2c[__obf_d74aacde370b8778.Option.Reserved.CreateTime]; __obf_fcb98d9ec4b8d44c {
			__obf_15a4a0514baf9e2c[__obf_d74aacde370b8778.Option.Reserved.CreateTime] = __obf_43dd99531253543c
		}
		if _, __obf_fcb98d9ec4b8d44c := __obf_15a4a0514baf9e2c[__obf_d74aacde370b8778.Option.Reserved.UpdateTime]; __obf_fcb98d9ec4b8d44c {
			__obf_15a4a0514baf9e2c[__obf_d74aacde370b8778.Option.Reserved.UpdateTime] = __obf_43dd99531253543c
		}
	} else {
		__obf_3c80b73d77602a41,
			// update
			__obf_15a4a0514baf9e2c, __obf_3220344ba911fc77 = BuildUpdateSQL(__obf_61e43608656dfcc0, __obf_dadf54047061b513, __obf_d74aacde370b8778.Option.Reserved.UpdateTime, []string{__obf_d74aacde370b8778.Option.Reserved.ID}, __obf_d74aacde370b8778.Option.Reserved.CreateTime, __obf_d74aacde370b8778.Option.Reserved.NO)
		if __obf_3220344ba911fc77 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_3220344ba911fc77.Error())
		}
	}

	if __obf_3c80b73d77602a41 != "" {
		__obf_a6490b2f4d4d2871, __obf_3220344ba911fc77 := __obf_d74aacde370b8778.DB.NamedExec(__obf_3c80b73d77602a41, __obf_15a4a0514baf9e2c)
		if __obf_3220344ba911fc77 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_3220344ba911fc77.Error())
		}
		if __obf_525801f57a5d856f == 0 {
			return __obf_a6490b2f4d4d2871.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_3c80b73d77602a41, "args", __obf_15a4a0514baf9e2c)
	}

	return __obf_525801f57a5d856f, nil
}

func (__obf_d74aacde370b8778 *ORM) SaveTxModel(__obf_2623653de66b5673 *sqlx.Tx, __obf_61e43608656dfcc0 any, __obf_525801f57a5d856f string) error {
	var __obf_3c80b73d77602a41 string
	var __obf_15a4a0514baf9e2c map[string]any
	__obf_3c80b73d77602a41, __obf_15a4a0514baf9e2c = __obf_24e8a2f2d64fbd4d(__obf_d74aacde370b8778.Option, __obf_61e43608656dfcc0, __obf_525801f57a5d856f)

	if __obf_3c80b73d77602a41 != "" {
		if _, __obf_3220344ba911fc77 := __obf_2623653de66b5673.NamedExec(__obf_3c80b73d77602a41, __obf_15a4a0514baf9e2c); __obf_3220344ba911fc77 != nil {
			_ = __obf_2623653de66b5673.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_3220344ba911fc77)
		}

		slog.Debug("SaveTxModel", "query", __obf_3c80b73d77602a41, "args", __obf_15a4a0514baf9e2c)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_24e8a2f2d64fbd4d[RecordID int | string](__obf_79c864729015d67e *Options, __obf_61e43608656dfcc0 any, __obf_525801f57a5d856f RecordID) (__obf_1f70f2f94054a64f string, __obf_15a4a0514baf9e2c map[string]any) {
	__obf_8a2e311a616cdf4b := reflect.ValueOf(__obf_61e43608656dfcc0)
	if __obf_8a2e311a616cdf4b.Kind() == reflect.Pointer {
		__obf_8a2e311a616cdf4b = __obf_8a2e311a616cdf4b.Elem()
	}
	var __obf_1a4fdd0ca9581fec []string
	__obf_15a4a0514baf9e2c = make(map[string]any)
	__obf_d78ea253231b6890 := __obf_8a2e311a616cdf4b.FieldByName("Id")
	if __obf_d78ea253231b6890.IsZero() {
		// 判断inserId是否为零值
		if __obf_525801f57a5d856f != *new(RecordID) {
			__obf_1a4fdd0ca9581fec = append(__obf_1a4fdd0ca9581fec, __obf_79c864729015d67e.Reserved.ID)
			__obf_15a4a0514baf9e2c[__obf_79c864729015d67e.Reserved.ID] = __obf_525801f57a5d856f
		} else {
			if __obf_d78ea253231b6890.Kind() == reflect.String {
				__obf_1a4fdd0ca9581fec = append(__obf_1a4fdd0ca9581fec, __obf_79c864729015d67e.Reserved.ID)
				__obf_15a4a0514baf9e2c[__obf_79c864729015d67e.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}
		__obf_d8b9ea30b41bc92c := reflect.Indirect(__obf_8a2e311a616cdf4b).Type()
		for __obf_5f1ac33f5a6c1104 := range __obf_d8b9ea30b41bc92c.NumField() {
			if __obf_af66979cc7b3eb24 := __obf_d8b9ea30b41bc92c.Field(__obf_5f1ac33f5a6c1104); __obf_af66979cc7b3eb24.Tag.Get("db") != "-" {
				switch __obf_89fbeda44e0c76d8 := util.ToSnake(__obf_af66979cc7b3eb24.Name); __obf_89fbeda44e0c76d8 {
				case __obf_79c864729015d67e.Reserved.CreateTime, __obf_79c864729015d67e.Reserved.UpdateTime:
					__obf_1a4fdd0ca9581fec = append(__obf_1a4fdd0ca9581fec, __obf_89fbeda44e0c76d8)
					__obf_15a4a0514baf9e2c[__obf_89fbeda44e0c76d8] = time.Now().Unix()
				default:
					__obf_abd89eb645093144 := __obf_8a2e311a616cdf4b.Field(__obf_5f1ac33f5a6c1104).Interface()
					if strings.HasPrefix(__obf_89fbeda44e0c76d8, "is_") || __obf_abd89eb645093144 != reflect.Zero(__obf_af66979cc7b3eb24.Type).Interface() {
						__obf_1a4fdd0ca9581fec = append(__obf_1a4fdd0ca9581fec, __obf_89fbeda44e0c76d8)
						__obf_15a4a0514baf9e2c[__obf_89fbeda44e0c76d8] = __obf_abd89eb645093144
					}
				}
			}
		}
		__obf_1f70f2f94054a64f = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_d8b9ea30b41bc92c.Name()), strings.Join(__obf_1a4fdd0ca9581fec, ","), strings.Join(__obf_1a4fdd0ca9581fec, ",:"))
	} else {
		__obf_d8b9ea30b41bc92c := reflect.Indirect(__obf_8a2e311a616cdf4b).Type()
		for __obf_5f1ac33f5a6c1104 := range __obf_d8b9ea30b41bc92c.NumField() {
			if __obf_af66979cc7b3eb24 := __obf_d8b9ea30b41bc92c.Field(__obf_5f1ac33f5a6c1104); __obf_af66979cc7b3eb24.Tag.Get("db") != "-" {
				__obf_abd89eb645093144 := __obf_8a2e311a616cdf4b.Field(__obf_5f1ac33f5a6c1104).Interface()
				if __obf_89fbeda44e0c76d8 := util.ToSnake(__obf_af66979cc7b3eb24.Name); strings.HasPrefix(__obf_89fbeda44e0c76d8, "is_") || __obf_abd89eb645093144 != reflect.Zero(__obf_af66979cc7b3eb24.Type).Interface() {
					__obf_1a4fdd0ca9581fec = append(__obf_1a4fdd0ca9581fec, fmt.Sprintf("%s=:%s", __obf_89fbeda44e0c76d8, __obf_89fbeda44e0c76d8))
					__obf_15a4a0514baf9e2c[__obf_89fbeda44e0c76d8] = __obf_abd89eb645093144
				}
			}
		}

		if util.HasField(__obf_8a2e311a616cdf4b, util.ToCamel(__obf_79c864729015d67e.Reserved.UpdateTime)) {
			__obf_9092e32882a9d698 := __obf_15a4a0514baf9e2c[__obf_79c864729015d67e.Reserved.UpdateTime]
			__obf_15a4a0514baf9e2c[__obf_79c864729015d67e.Reserved.UpdateTime] = time.Now().Unix()
			__obf_1f70f2f94054a64f = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_d8b9ea30b41bc92c.Name()), strings.Join(__obf_1a4fdd0ca9581fec, ","), __obf_79c864729015d67e.Reserved.ID, __obf_79c864729015d67e.Reserved.ID, __obf_79c864729015d67e.Reserved.UpdateTime, __obf_9092e32882a9d698)
		} else {
			__obf_1f70f2f94054a64f = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_d8b9ea30b41bc92c.Name()), strings.Join(__obf_1a4fdd0ca9581fec, ","), __obf_79c864729015d67e.Reserved.ID, __obf_79c864729015d67e.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_86248f54b9553448 any, __obf_dadf54047061b513 string) (string, map[string]any, error) {
	if __obf_dadf54047061b513 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_15a4a0514baf9e2c := make(map[string]any)
	__obf_1a4fdd0ca9581fec := []string{}
	__obf_6f0d1f91568c2367 := func(__obf_e144436f864ee67b string, __obf_43a03d8b7535a35f any, __obf_cf0f8be53ffc786c bool) {
		__obf_cfb38cc815e93ce1 := __obf_43a03d8b7535a35f == nil || reflect.DeepEqual(__obf_43a03d8b7535a35f, reflect.Zero(reflect.TypeOf(__obf_43a03d8b7535a35f)).Interface())
		if __obf_cf0f8be53ffc786c && __obf_cfb38cc815e93ce1 {
			return
		}
		__obf_1a4fdd0ca9581fec = append(__obf_1a4fdd0ca9581fec, __obf_e144436f864ee67b)
		__obf_15a4a0514baf9e2c[__obf_e144436f864ee67b] = __obf_43a03d8b7535a35f
	}

	switch __obf_abd89eb645093144 := __obf_86248f54b9553448.(type) {
	case map[string]any:
		for __obf_5304b9e5e1368ecf, __obf_8a2e311a616cdf4b := range __obf_abd89eb645093144 {
			__obf_6f0d1f91568c2367(__obf_5304b9e5e1368ecf, __obf_8a2e311a616cdf4b, false)
		}

	default:
		__obf_5f939ebac6d18417 := reflect.TypeOf(__obf_86248f54b9553448)
		if __obf_5f939ebac6d18417.Kind() == reflect.Pointer {
			__obf_5f939ebac6d18417 = __obf_5f939ebac6d18417.Elem()
		}
		__obf_8a2e311a616cdf4b := reflect.ValueOf(__obf_86248f54b9553448)
		if __obf_8a2e311a616cdf4b.Kind() == reflect.Pointer {
			__obf_8a2e311a616cdf4b = __obf_8a2e311a616cdf4b.Elem()
		}

		if __obf_5f939ebac6d18417.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_5f1ac33f5a6c1104 := 0; __obf_5f1ac33f5a6c1104 < __obf_5f939ebac6d18417.NumField(); __obf_5f1ac33f5a6c1104++ {
			__obf_2c9e17466249e406 := __obf_5f939ebac6d18417.Field(__obf_5f1ac33f5a6c1104).Tag.Get("db")
			if __obf_2c9e17466249e406 == "" || __obf_2c9e17466249e406 == "-" {
				continue
			}
			__obf_0fde6f11c19c1386 := strings.Split(__obf_2c9e17466249e406, ",")
			__obf_6f0d1f91568c2367(__obf_0fde6f11c19c1386[0], __obf_8a2e311a616cdf4b.Field(__obf_5f1ac33f5a6c1104).Interface(), len(__obf_0fde6f11c19c1386) > 1 && __obf_0fde6f11c19c1386[1] == "omitempty")
		}
	}

	if len(__obf_1a4fdd0ca9581fec) == 0 {
		return "", nil, errors.New("no fields to insert")
	}
	__obf_3c80b73d77602a41 := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", __obf_dadf54047061b513, strings.Join(__obf_1a4fdd0ca9581fec, ","),
		strings.Join(__obf_1a4fdd0ca9581fec, ",:"),
	)

	return __obf_3c80b73d77602a41,

		// pkField 主键字段
		// lockField 乐观锁字段(默认为updated_at)
		// constField 固定不变字段
		__obf_15a4a0514baf9e2c, nil
}

func BuildUpdateSQL(__obf_86248f54b9553448 any, __obf_dadf54047061b513, __obf_89b212035215da6d string, __obf_e9f5502e55c87eac []string, __obf_6af07d3e30f354d6 ...string,
) (string, map[string]any, error) {
	if __obf_dadf54047061b513 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_366c99879a3fb838 := make(map[string]struct{})
	for _, __obf_792444f86a2bdc4c := range __obf_e9f5502e55c87eac {
		__obf_366c99879a3fb838[__obf_792444f86a2bdc4c] = struct{}{}
	}
	__obf_15a4a0514baf9e2c := make(map[string]any)
	__obf_1a4fdd0ca9581fec := []string{}
	__obf_c3a8e8feffdf7b76 := []string{}
	__obf_6f0d1f91568c2367 := func(__obf_e144436f864ee67b string, __obf_43a03d8b7535a35f any, __obf_cf0f8be53ffc786c bool) {
		__obf_cfb38cc815e93ce1 := __obf_43a03d8b7535a35f == nil || reflect.DeepEqual(__obf_43a03d8b7535a35f, reflect.Zero(reflect.TypeOf(__obf_43a03d8b7535a35f)).Interface())

		switch {
		case func() bool {
			_, __obf_fcb98d9ec4b8d44c := __obf_366c99879a3fb838[__obf_e144436f864ee67b]
			return __obf_fcb98d9ec4b8d44c
		}():
			__obf_c3a8e8feffdf7b76 = append(__obf_c3a8e8feffdf7b76, fmt.Sprintf("%s=:%s", __obf_e144436f864ee67b, __obf_e144436f864ee67b))
			__obf_15a4a0514baf9e2c[__obf_e144436f864ee67b] = __obf_43a03d8b7535a35f

		case slices.Contains(__obf_6af07d3e30f354d6, __obf_e144436f864ee67b):
			// skip constField
		case __obf_e144436f864ee67b == __obf_89b212035215da6d:
			const __obf_c707dccc84b0b62f = "RESERVED_LOCK"
			__obf_1a4fdd0ca9581fec = append(__obf_1a4fdd0ca9581fec, fmt.Sprintf("%s=:%s", __obf_e144436f864ee67b, __obf_e144436f864ee67b))
			__obf_c3a8e8feffdf7b76 = append(__obf_c3a8e8feffdf7b76, fmt.Sprintf("%s=:%s", __obf_e144436f864ee67b, __obf_c707dccc84b0b62f))
			__obf_15a4a0514baf9e2c[__obf_e144436f864ee67b] = time.Now().Unix()
			__obf_15a4a0514baf9e2c[__obf_c707dccc84b0b62f] = __obf_43a03d8b7535a35f

		default:
			if __obf_cf0f8be53ffc786c && __obf_cfb38cc815e93ce1 {
				return
			}
			__obf_1a4fdd0ca9581fec = append(__obf_1a4fdd0ca9581fec, fmt.Sprintf("%s=:%s", __obf_e144436f864ee67b, __obf_e144436f864ee67b))
			__obf_15a4a0514baf9e2c[__obf_e144436f864ee67b] = __obf_43a03d8b7535a35f
		}
	}

	switch __obf_abd89eb645093144 := __obf_86248f54b9553448.(type) {
	case map[string]any:
		for __obf_5304b9e5e1368ecf, __obf_8a2e311a616cdf4b := range __obf_abd89eb645093144 {
			__obf_6f0d1f91568c2367(__obf_5304b9e5e1368ecf, __obf_8a2e311a616cdf4b, false)
		}

	default:
		__obf_5f939ebac6d18417 := reflect.TypeOf(__obf_86248f54b9553448)
		if __obf_5f939ebac6d18417.Kind() == reflect.Pointer {
			__obf_5f939ebac6d18417 = __obf_5f939ebac6d18417.Elem()
		}
		__obf_8a2e311a616cdf4b := reflect.ValueOf(__obf_86248f54b9553448)
		if __obf_8a2e311a616cdf4b.Kind() == reflect.Pointer {
			__obf_8a2e311a616cdf4b = __obf_8a2e311a616cdf4b.Elem()
		}
		if __obf_5f939ebac6d18417.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_5f1ac33f5a6c1104 := 0; __obf_5f1ac33f5a6c1104 < __obf_5f939ebac6d18417.NumField(); __obf_5f1ac33f5a6c1104++ {
			__obf_2c9e17466249e406 := __obf_5f939ebac6d18417.Field(__obf_5f1ac33f5a6c1104).Tag.Get("db")
			if __obf_2c9e17466249e406 == "" || __obf_2c9e17466249e406 == "-" {
				continue
			}
			__obf_0fde6f11c19c1386 := strings.Split(__obf_2c9e17466249e406, ",")
			__obf_6f0d1f91568c2367(__obf_0fde6f11c19c1386[0], __obf_8a2e311a616cdf4b.Field(__obf_5f1ac33f5a6c1104).Interface(), len(__obf_0fde6f11c19c1386) > 1 && __obf_0fde6f11c19c1386[1] == "omitempty")
		}
	}

	if len(__obf_e9f5502e55c87eac) == 0 || len(__obf_c3a8e8feffdf7b76) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_1a4fdd0ca9581fec) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	__obf_3c80b73d77602a41 := // if lockField != "" && !hasLockField {
		// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
		// }

		fmt.Sprintf("UPDATE %s SET %s WHERE %s", __obf_dadf54047061b513, strings.Join(__obf_1a4fdd0ca9581fec, ","),
			strings.Join(__obf_c3a8e8feffdf7b76, " AND "),
		)

	return __obf_3c80b73d77602a41, __obf_15a4a0514baf9e2c, nil
}
