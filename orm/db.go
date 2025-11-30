package __obf_f47aac06a08e5dbb

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

func NewORM(__obf_71da7205c5561eb4 *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_71da7205c5561eb4.Driver, __obf_71da7205c5561eb4.DataSource)

	switch __obf_71da7205c5561eb4.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_71da7205c5561eb4.
			Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_71da7205c5561eb4.
			Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_71da7205c5561eb4.
			Placeholder = ":"
	case "sqlserver":
		__obf_71da7205c5561eb4.
			Placeholder = "@"
	}

	if __obf_9ac6018edd4832e7 := DB.Ping(); __obf_9ac6018edd4832e7 != nil {
		panic(__obf_9ac6018edd4832e7)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_71da7205c5561eb4,
		}, func() {
			DB.Close()
		}
}

func (__obf_ab97c1322d9ea2f2 *ORM) ListMap(__obf_22a32b38c71f2a5d *sqlx.Rows, __obf_3447427d651270c4 func(map[string]any) (string, string)) (__obf_c565a64719fc9ebe []map[string]any, __obf_9ac6018edd4832e7 error) {
	for __obf_22a32b38c71f2a5d.Next() {
		__obf_5639e56d717dc56e := make(map[string]any)
		if __obf_9ac6018edd4832e7 = __obf_22a32b38c71f2a5d.MapScan(__obf_5639e56d717dc56e); __obf_9ac6018edd4832e7 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_9ac6018edd4832e7)
		}
		for __obf_5103bf8cf045836b, __obf_d6e1d0a7185fed66 := range __obf_5639e56d717dc56e {
			switch __obf_d6e1d0a7185fed66 := __obf_d6e1d0a7185fed66.(type) {
			case []uint8:
				__obf_5639e56d717dc56e[__obf_5103bf8cf045836b] = string(__obf_d6e1d0a7185fed66)
			case int64:
				__obf_5639e56d717dc56e[__obf_5103bf8cf045836b] = int64(__obf_d6e1d0a7185fed66)
			case uint64:
				__obf_5639e56d717dc56e[__obf_5103bf8cf045836b] = uint64(__obf_d6e1d0a7185fed66)
			}
		}
		if __obf_3447427d651270c4 != nil {
			__obf_af49a21815e3d222, __obf_9554751b39677f0e := __obf_3447427d651270c4(__obf_5639e56d717dc56e)
			__obf_5639e56d717dc56e[__obf_af49a21815e3d222] = __obf_9554751b39677f0e
		}
		__obf_c565a64719fc9ebe = append(__obf_c565a64719fc9ebe, __obf_5639e56d717dc56e)
	}
	return
}

func (__obf_ab97c1322d9ea2f2 *ORM) Exists(__obf_c0d7776bdf9e184b, __obf_23d3ae8f01b6a47d string, __obf_0bdb4f0ba995788b ...any) bool {
	var __obf_ab33d7b8ddfe6d34 bool
	_ = __obf_ab97c1322d9ea2f2.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_c0d7776bdf9e184b, __obf_23d3ae8f01b6a47d), __obf_0bdb4f0ba995788b...).Scan(&__obf_ab33d7b8ddfe6d34)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_ab33d7b8ddfe6d34
}

func (__obf_ab97c1322d9ea2f2 *ORM) SaveModel(__obf_ecf40ee198ffe028 any, __obf_3f212fc13bf1eb54 string) (string, error) {
	__obf_d6e1d0a7185fed66 := reflect.ValueOf(__obf_ecf40ee198ffe028)
	if __obf_d6e1d0a7185fed66.Kind() == reflect.Pointer {
		__obf_d6e1d0a7185fed66 = __obf_d6e1d0a7185fed66.Elem()
	}
	var (
		__obf_50223a5e658ece82 = __obf_d6e1d0a7185fed66.
					FieldByName("Id")
		__obf_bce62f3d54639633 = util.ToSnake(reflect.Indirect(__obf_d6e1d0a7185fed66).Type().Name())
		__obf_b7fc1611623cb763 string
		__obf_0bdb4f0ba995788b map[string]any
		__obf_9ac6018edd4832e7 error
	)

	if __obf_50223a5e658ece82.IsZero() {
		__obf_b7fc1611623cb763,
			// insert
			__obf_0bdb4f0ba995788b, __obf_9ac6018edd4832e7 = BuildInsertSQL(__obf_ecf40ee198ffe028, __obf_bce62f3d54639633)
		if __obf_9ac6018edd4832e7 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_9ac6018edd4832e7.Error())
		}
		if __obf_3f212fc13bf1eb54 == "" {
			__obf_3f212fc13bf1eb54 = util.StringUUID()
		}
		__obf_0bdb4f0ba995788b["id"] = __obf_3f212fc13bf1eb54
		__obf_1ee88a996f678cc2 := time.Now().Unix()
		if _, __obf_410061093b631faf := __obf_0bdb4f0ba995788b[__obf_ab97c1322d9ea2f2.Option.Reserved.CreateTime]; __obf_410061093b631faf {
			__obf_0bdb4f0ba995788b[__obf_ab97c1322d9ea2f2.Option.Reserved.CreateTime] = __obf_1ee88a996f678cc2
		}
		if _, __obf_410061093b631faf := __obf_0bdb4f0ba995788b[__obf_ab97c1322d9ea2f2.Option.Reserved.UpdateTime]; __obf_410061093b631faf {
			__obf_0bdb4f0ba995788b[__obf_ab97c1322d9ea2f2.Option.Reserved.UpdateTime] = __obf_1ee88a996f678cc2
		}
	} else {
		__obf_b7fc1611623cb763,
			// update
			__obf_0bdb4f0ba995788b, __obf_9ac6018edd4832e7 = BuildUpdateSQL(__obf_ecf40ee198ffe028, __obf_bce62f3d54639633, __obf_ab97c1322d9ea2f2.Option.Reserved.UpdateTime, []string{__obf_ab97c1322d9ea2f2.Option.Reserved.ID}, __obf_ab97c1322d9ea2f2.Option.Reserved.CreateTime, __obf_ab97c1322d9ea2f2.Option.Reserved.NO)
		if __obf_9ac6018edd4832e7 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_9ac6018edd4832e7.Error())
		}
	}

	if __obf_b7fc1611623cb763 != "" {
		_, __obf_9ac6018edd4832e7 := __obf_ab97c1322d9ea2f2.DB.NamedExec(__obf_b7fc1611623cb763, __obf_0bdb4f0ba995788b)
		if __obf_9ac6018edd4832e7 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_9ac6018edd4832e7.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_b7fc1611623cb763, "args", __obf_0bdb4f0ba995788b)
	}

	return __obf_3f212fc13bf1eb54, nil

}

func (__obf_ab97c1322d9ea2f2 *ORM) SaveModelWidthAutoID(__obf_ecf40ee198ffe028 any, __obf_3f212fc13bf1eb54 int64) (int64, error) {
	__obf_d6e1d0a7185fed66 := // query, args := modelBind(db.Option, m, insertId)
		reflect.ValueOf(__obf_ecf40ee198ffe028)
	if __obf_d6e1d0a7185fed66.Kind() == reflect.Pointer {
		__obf_d6e1d0a7185fed66 = __obf_d6e1d0a7185fed66.Elem()
	}
	__obf_50223a5e658ece82 := __obf_d6e1d0a7185fed66.FieldByName("Id")
	__obf_bce62f3d54639633 := util.ToSnake(reflect.Indirect(__obf_d6e1d0a7185fed66).Type().Name())
	var (
		__obf_b7fc1611623cb763 string
		__obf_0bdb4f0ba995788b map[string]any
		__obf_9ac6018edd4832e7 error
	)

	if __obf_50223a5e658ece82.IsZero() {
		__obf_b7fc1611623cb763,
			// insert
			__obf_0bdb4f0ba995788b, __obf_9ac6018edd4832e7 = BuildInsertSQL(__obf_ecf40ee198ffe028, __obf_bce62f3d54639633)
		if __obf_9ac6018edd4832e7 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_9ac6018edd4832e7.Error())
		}
		if __obf_3f212fc13bf1eb54 != 0 {
			__obf_0bdb4f0ba995788b["id"] = __obf_3f212fc13bf1eb54
		}
		__obf_1ee88a996f678cc2 := time.Now().Unix()
		if _, __obf_410061093b631faf := __obf_0bdb4f0ba995788b[__obf_ab97c1322d9ea2f2.Option.Reserved.CreateTime]; __obf_410061093b631faf {
			__obf_0bdb4f0ba995788b[__obf_ab97c1322d9ea2f2.Option.Reserved.CreateTime] = __obf_1ee88a996f678cc2
		}
		if _, __obf_410061093b631faf := __obf_0bdb4f0ba995788b[__obf_ab97c1322d9ea2f2.Option.Reserved.UpdateTime]; __obf_410061093b631faf {
			__obf_0bdb4f0ba995788b[__obf_ab97c1322d9ea2f2.Option.Reserved.UpdateTime] = __obf_1ee88a996f678cc2
		}
	} else {
		__obf_b7fc1611623cb763,
			// update
			__obf_0bdb4f0ba995788b, __obf_9ac6018edd4832e7 = BuildUpdateSQL(__obf_ecf40ee198ffe028, __obf_bce62f3d54639633, __obf_ab97c1322d9ea2f2.Option.Reserved.UpdateTime, []string{__obf_ab97c1322d9ea2f2.Option.Reserved.ID}, __obf_ab97c1322d9ea2f2.Option.Reserved.CreateTime, __obf_ab97c1322d9ea2f2.Option.Reserved.NO)
		if __obf_9ac6018edd4832e7 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_9ac6018edd4832e7.Error())
		}
	}

	if __obf_b7fc1611623cb763 != "" {
		__obf_9554751b39677f0e, __obf_9ac6018edd4832e7 := __obf_ab97c1322d9ea2f2.DB.NamedExec(__obf_b7fc1611623cb763, __obf_0bdb4f0ba995788b)
		if __obf_9ac6018edd4832e7 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_9ac6018edd4832e7.Error())
		}
		if __obf_3f212fc13bf1eb54 == 0 {
			return __obf_9554751b39677f0e.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_b7fc1611623cb763, "args", __obf_0bdb4f0ba995788b)
	}

	return __obf_3f212fc13bf1eb54, nil
}

func (__obf_ab97c1322d9ea2f2 *ORM) SaveTxModel(__obf_07b88b883eeed833 *sqlx.Tx, __obf_ecf40ee198ffe028 any, __obf_3f212fc13bf1eb54 string) error {
	var __obf_b7fc1611623cb763 string
	var __obf_0bdb4f0ba995788b map[string]any
	__obf_b7fc1611623cb763, __obf_0bdb4f0ba995788b = __obf_5c62521ec4c2772e(__obf_ab97c1322d9ea2f2.Option, __obf_ecf40ee198ffe028, __obf_3f212fc13bf1eb54)

	if __obf_b7fc1611623cb763 != "" {
		if _, __obf_9ac6018edd4832e7 := __obf_07b88b883eeed833.NamedExec(__obf_b7fc1611623cb763, __obf_0bdb4f0ba995788b); __obf_9ac6018edd4832e7 != nil {
			_ = __obf_07b88b883eeed833.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_9ac6018edd4832e7)
		}

		slog.Debug("SaveTxModel", "query", __obf_b7fc1611623cb763, "args", __obf_0bdb4f0ba995788b)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_5c62521ec4c2772e[RecordID int | string](__obf_71da7205c5561eb4 *Options, __obf_ecf40ee198ffe028 any, __obf_3f212fc13bf1eb54 RecordID) (__obf_1c5eddb24bf91307 string, __obf_0bdb4f0ba995788b map[string]any) {
	__obf_d6e1d0a7185fed66 := reflect.ValueOf(__obf_ecf40ee198ffe028)
	if __obf_d6e1d0a7185fed66.Kind() == reflect.Pointer {
		__obf_d6e1d0a7185fed66 = __obf_d6e1d0a7185fed66.Elem()
	}
	var __obf_79789738de9b87b0 []string
	__obf_0bdb4f0ba995788b = make(map[string]any)
	__obf_50223a5e658ece82 := __obf_d6e1d0a7185fed66.FieldByName("Id")
	if __obf_50223a5e658ece82.IsZero() {
		// 判断inserId是否为零值
		if __obf_3f212fc13bf1eb54 != *new(RecordID) {
			__obf_79789738de9b87b0 = append(__obf_79789738de9b87b0, __obf_71da7205c5561eb4.Reserved.ID)
			__obf_0bdb4f0ba995788b[__obf_71da7205c5561eb4.Reserved.ID] = __obf_3f212fc13bf1eb54
		} else {
			if __obf_50223a5e658ece82.Kind() == reflect.String {
				__obf_79789738de9b87b0 = append(__obf_79789738de9b87b0, __obf_71da7205c5561eb4.Reserved.ID)
				__obf_0bdb4f0ba995788b[__obf_71da7205c5561eb4.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}
		__obf_5e46d635fee1c260 := reflect.Indirect(__obf_d6e1d0a7185fed66).Type()
		for __obf_89306d663fb5edd8 := range __obf_5e46d635fee1c260.NumField() {
			if __obf_c9ab6b5901a38cd0 := __obf_5e46d635fee1c260.Field(__obf_89306d663fb5edd8); __obf_c9ab6b5901a38cd0.Tag.Get("db") != "-" {
				switch __obf_600298a232583a7c := util.ToSnake(__obf_c9ab6b5901a38cd0.Name); __obf_600298a232583a7c {
				case __obf_71da7205c5561eb4.Reserved.CreateTime, __obf_71da7205c5561eb4.Reserved.UpdateTime:
					__obf_79789738de9b87b0 = append(__obf_79789738de9b87b0, __obf_600298a232583a7c)
					__obf_0bdb4f0ba995788b[__obf_600298a232583a7c] = time.Now().Unix()
				default:
					__obf_613929daa0d28687 := __obf_d6e1d0a7185fed66.Field(__obf_89306d663fb5edd8).Interface()
					if strings.HasPrefix(__obf_600298a232583a7c, "is_") || __obf_613929daa0d28687 != reflect.Zero(__obf_c9ab6b5901a38cd0.Type).Interface() {
						__obf_79789738de9b87b0 = append(__obf_79789738de9b87b0, __obf_600298a232583a7c)
						__obf_0bdb4f0ba995788b[__obf_600298a232583a7c] = __obf_613929daa0d28687
					}
				}
			}
		}
		__obf_1c5eddb24bf91307 = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_5e46d635fee1c260.Name()), strings.Join(__obf_79789738de9b87b0, ","), strings.Join(__obf_79789738de9b87b0, ",:"))
	} else {
		__obf_5e46d635fee1c260 := reflect.Indirect(__obf_d6e1d0a7185fed66).Type()
		for __obf_89306d663fb5edd8 := range __obf_5e46d635fee1c260.NumField() {
			if __obf_c9ab6b5901a38cd0 := __obf_5e46d635fee1c260.Field(__obf_89306d663fb5edd8); __obf_c9ab6b5901a38cd0.Tag.Get("db") != "-" {
				__obf_613929daa0d28687 := __obf_d6e1d0a7185fed66.Field(__obf_89306d663fb5edd8).Interface()
				if __obf_600298a232583a7c := util.ToSnake(__obf_c9ab6b5901a38cd0.Name); strings.HasPrefix(__obf_600298a232583a7c, "is_") || __obf_613929daa0d28687 != reflect.Zero(__obf_c9ab6b5901a38cd0.Type).Interface() {
					__obf_79789738de9b87b0 = append(__obf_79789738de9b87b0, fmt.Sprintf("%s=:%s", __obf_600298a232583a7c, __obf_600298a232583a7c))
					__obf_0bdb4f0ba995788b[__obf_600298a232583a7c] = __obf_613929daa0d28687
				}
			}
		}

		if util.HasField(__obf_d6e1d0a7185fed66, util.ToCamel(__obf_71da7205c5561eb4.Reserved.UpdateTime)) {
			__obf_d74ed11548a8bbcf := __obf_0bdb4f0ba995788b[__obf_71da7205c5561eb4.Reserved.UpdateTime]
			__obf_0bdb4f0ba995788b[__obf_71da7205c5561eb4.Reserved.UpdateTime] = time.Now().Unix()
			__obf_1c5eddb24bf91307 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_5e46d635fee1c260.Name()), strings.Join(__obf_79789738de9b87b0, ","), __obf_71da7205c5561eb4.Reserved.ID, __obf_71da7205c5561eb4.Reserved.ID, __obf_71da7205c5561eb4.Reserved.UpdateTime, __obf_d74ed11548a8bbcf)
		} else {
			__obf_1c5eddb24bf91307 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_5e46d635fee1c260.Name()), strings.Join(__obf_79789738de9b87b0, ","), __obf_71da7205c5561eb4.Reserved.ID, __obf_71da7205c5561eb4.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_9dd9bcccbfb124d6 any, __obf_bce62f3d54639633 string) (string, map[string]any, error) {
	if __obf_bce62f3d54639633 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_0bdb4f0ba995788b := make(map[string]any)
	__obf_79789738de9b87b0 := []string{}
	__obf_998d47ed8188db1d := func(__obf_40948bc1feaa6823 string, __obf_e51512e07c73036f any, __obf_458a7bee86e0d236 bool) {
		__obf_c2b0f37aae6a25b6 := __obf_e51512e07c73036f == nil || reflect.DeepEqual(__obf_e51512e07c73036f, reflect.Zero(reflect.TypeOf(__obf_e51512e07c73036f)).Interface())
		if __obf_458a7bee86e0d236 && __obf_c2b0f37aae6a25b6 {
			return
		}
		__obf_79789738de9b87b0 = append(__obf_79789738de9b87b0, __obf_40948bc1feaa6823)
		__obf_0bdb4f0ba995788b[__obf_40948bc1feaa6823] = __obf_e51512e07c73036f
	}

	switch __obf_613929daa0d28687 := __obf_9dd9bcccbfb124d6.(type) {
	case map[string]any:
		for __obf_5103bf8cf045836b, __obf_d6e1d0a7185fed66 := range __obf_613929daa0d28687 {
			__obf_998d47ed8188db1d(__obf_5103bf8cf045836b, __obf_d6e1d0a7185fed66, false)
		}

	default:
		__obf_c3d7728bbd58f361 := reflect.TypeOf(__obf_9dd9bcccbfb124d6)
		if __obf_c3d7728bbd58f361.Kind() == reflect.Pointer {
			__obf_c3d7728bbd58f361 = __obf_c3d7728bbd58f361.Elem()
		}
		__obf_d6e1d0a7185fed66 := reflect.ValueOf(__obf_9dd9bcccbfb124d6)
		if __obf_d6e1d0a7185fed66.Kind() == reflect.Pointer {
			__obf_d6e1d0a7185fed66 = __obf_d6e1d0a7185fed66.Elem()
		}

		if __obf_c3d7728bbd58f361.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_89306d663fb5edd8 := 0; __obf_89306d663fb5edd8 < __obf_c3d7728bbd58f361.NumField(); __obf_89306d663fb5edd8++ {
			__obf_3566929bef409e47 := __obf_c3d7728bbd58f361.Field(__obf_89306d663fb5edd8).Tag.Get("db")
			if __obf_3566929bef409e47 == "" || __obf_3566929bef409e47 == "-" {
				continue
			}
			__obf_232fe921aef9eb04 := strings.Split(__obf_3566929bef409e47, ",")
			__obf_998d47ed8188db1d(__obf_232fe921aef9eb04[0], __obf_d6e1d0a7185fed66.Field(__obf_89306d663fb5edd8).Interface(), len(__obf_232fe921aef9eb04) > 1 && __obf_232fe921aef9eb04[1] == "omitempty")
		}
	}

	if len(__obf_79789738de9b87b0) == 0 {
		return "", nil, errors.New("no fields to insert")
	}
	__obf_b7fc1611623cb763 := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", __obf_bce62f3d54639633, strings.Join(__obf_79789738de9b87b0, ","),
		strings.Join(__obf_79789738de9b87b0, ",:"),
	)

	return __obf_b7fc1611623cb763,

		// pkField 主键字段
		// lockField 乐观锁字段(默认为updated_at)
		// constField 固定不变字段
		__obf_0bdb4f0ba995788b, nil
}

func BuildUpdateSQL(__obf_9dd9bcccbfb124d6 any, __obf_bce62f3d54639633, __obf_6a5cbac3c1c52635 string, __obf_22c40ebe50e6ac76 []string, __obf_a1fa340133a93181 ...string,
) (string, map[string]any, error) {
	if __obf_bce62f3d54639633 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_6ecec620e55275ab := make(map[string]struct{})
	for _, __obf_5093d710cc493d23 := range __obf_22c40ebe50e6ac76 {
		__obf_6ecec620e55275ab[__obf_5093d710cc493d23] = struct{}{}
	}
	__obf_0bdb4f0ba995788b := make(map[string]any)
	__obf_79789738de9b87b0 := []string{}
	__obf_99af445a437db062 := []string{}
	__obf_998d47ed8188db1d := func(__obf_40948bc1feaa6823 string, __obf_e51512e07c73036f any, __obf_458a7bee86e0d236 bool) {
		__obf_c2b0f37aae6a25b6 := __obf_e51512e07c73036f == nil || reflect.DeepEqual(__obf_e51512e07c73036f, reflect.Zero(reflect.TypeOf(__obf_e51512e07c73036f)).Interface())

		switch {
		case func() bool {
			_, __obf_410061093b631faf := __obf_6ecec620e55275ab[__obf_40948bc1feaa6823]
			return __obf_410061093b631faf
		}():
			__obf_99af445a437db062 = append(__obf_99af445a437db062, fmt.Sprintf("%s=:%s", __obf_40948bc1feaa6823, __obf_40948bc1feaa6823))
			__obf_0bdb4f0ba995788b[__obf_40948bc1feaa6823] = __obf_e51512e07c73036f

		case slices.Contains(__obf_a1fa340133a93181, __obf_40948bc1feaa6823):
			// skip constField
		case __obf_40948bc1feaa6823 == __obf_6a5cbac3c1c52635:
			const __obf_f9f577fca04b7873 = "RESERVED_LOCK"
			__obf_79789738de9b87b0 = append(__obf_79789738de9b87b0, fmt.Sprintf("%s=:%s", __obf_40948bc1feaa6823, __obf_40948bc1feaa6823))
			__obf_99af445a437db062 = append(__obf_99af445a437db062, fmt.Sprintf("%s=:%s", __obf_40948bc1feaa6823, __obf_f9f577fca04b7873))
			__obf_0bdb4f0ba995788b[__obf_40948bc1feaa6823] = time.Now().Unix()
			__obf_0bdb4f0ba995788b[__obf_f9f577fca04b7873] = __obf_e51512e07c73036f

		default:
			if __obf_458a7bee86e0d236 && __obf_c2b0f37aae6a25b6 {
				return
			}
			__obf_79789738de9b87b0 = append(__obf_79789738de9b87b0, fmt.Sprintf("%s=:%s", __obf_40948bc1feaa6823, __obf_40948bc1feaa6823))
			__obf_0bdb4f0ba995788b[__obf_40948bc1feaa6823] = __obf_e51512e07c73036f
		}
	}

	switch __obf_613929daa0d28687 := __obf_9dd9bcccbfb124d6.(type) {
	case map[string]any:
		for __obf_5103bf8cf045836b, __obf_d6e1d0a7185fed66 := range __obf_613929daa0d28687 {
			__obf_998d47ed8188db1d(__obf_5103bf8cf045836b, __obf_d6e1d0a7185fed66, false)
		}

	default:
		__obf_c3d7728bbd58f361 := reflect.TypeOf(__obf_9dd9bcccbfb124d6)
		if __obf_c3d7728bbd58f361.Kind() == reflect.Pointer {
			__obf_c3d7728bbd58f361 = __obf_c3d7728bbd58f361.Elem()
		}
		__obf_d6e1d0a7185fed66 := reflect.ValueOf(__obf_9dd9bcccbfb124d6)
		if __obf_d6e1d0a7185fed66.Kind() == reflect.Pointer {
			__obf_d6e1d0a7185fed66 = __obf_d6e1d0a7185fed66.Elem()
		}
		if __obf_c3d7728bbd58f361.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_89306d663fb5edd8 := 0; __obf_89306d663fb5edd8 < __obf_c3d7728bbd58f361.NumField(); __obf_89306d663fb5edd8++ {
			__obf_3566929bef409e47 := __obf_c3d7728bbd58f361.Field(__obf_89306d663fb5edd8).Tag.Get("db")
			if __obf_3566929bef409e47 == "" || __obf_3566929bef409e47 == "-" {
				continue
			}
			__obf_232fe921aef9eb04 := strings.Split(__obf_3566929bef409e47, ",")
			__obf_998d47ed8188db1d(__obf_232fe921aef9eb04[0], __obf_d6e1d0a7185fed66.Field(__obf_89306d663fb5edd8).Interface(), len(__obf_232fe921aef9eb04) > 1 && __obf_232fe921aef9eb04[1] == "omitempty")
		}
	}

	if len(__obf_22c40ebe50e6ac76) == 0 || len(__obf_99af445a437db062) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_79789738de9b87b0) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	__obf_b7fc1611623cb763 := // if lockField != "" && !hasLockField {
		// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
		// }

		fmt.Sprintf("UPDATE %s SET %s WHERE %s", __obf_bce62f3d54639633, strings.Join(__obf_79789738de9b87b0, ","),
			strings.Join(__obf_99af445a437db062, " AND "),
		)

	return __obf_b7fc1611623cb763, __obf_0bdb4f0ba995788b, nil
}
