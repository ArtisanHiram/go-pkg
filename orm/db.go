package __obf_d726bb43e85f2dfc

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

func NewORM(__obf_c9efa06b0a7790ff *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_c9efa06b0a7790ff.Driver, __obf_c9efa06b0a7790ff.DataSource)

	switch __obf_c9efa06b0a7790ff.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_c9efa06b0a7790ff.Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_c9efa06b0a7790ff.Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_c9efa06b0a7790ff.Placeholder = ":"
	case "sqlserver":
		__obf_c9efa06b0a7790ff.Placeholder = "@"
	}

	if __obf_f5c7d5fde3143a52 := DB.Ping(); __obf_f5c7d5fde3143a52 != nil {
		panic(__obf_f5c7d5fde3143a52)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_c9efa06b0a7790ff,
		}, func() {
			DB.Close()
		}
}

func (__obf_a1d3ecfbcf6e018c *ORM) ListMap(__obf_af10750736e4963e *sqlx.Rows, __obf_055a22efeea8af0f func(map[string]any) (string, string)) (__obf_3357f1ed7884ec23 []map[string]any, __obf_f5c7d5fde3143a52 error) {
	for __obf_af10750736e4963e.Next() {
		__obf_2708b8f15d03d7e5 := make(map[string]any)
		if __obf_f5c7d5fde3143a52 = __obf_af10750736e4963e.MapScan(__obf_2708b8f15d03d7e5); __obf_f5c7d5fde3143a52 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_f5c7d5fde3143a52)
		}
		for __obf_3f7031f75dab588c, __obf_53bf414fffa45e91 := range __obf_2708b8f15d03d7e5 {
			switch __obf_53bf414fffa45e91 := __obf_53bf414fffa45e91.(type) {
			case []uint8:
				__obf_2708b8f15d03d7e5[__obf_3f7031f75dab588c] = string(__obf_53bf414fffa45e91)
			case int64:
				__obf_2708b8f15d03d7e5[__obf_3f7031f75dab588c] = int64(__obf_53bf414fffa45e91)
			case uint64:
				__obf_2708b8f15d03d7e5[__obf_3f7031f75dab588c] = uint64(__obf_53bf414fffa45e91)
			}
		}
		if __obf_055a22efeea8af0f != nil {
			__obf_3ead1b8be5f3d59b, __obf_4fceda9cabeb1c4f := __obf_055a22efeea8af0f(__obf_2708b8f15d03d7e5)
			__obf_2708b8f15d03d7e5[__obf_3ead1b8be5f3d59b] = __obf_4fceda9cabeb1c4f
		}
		__obf_3357f1ed7884ec23 = append(__obf_3357f1ed7884ec23, __obf_2708b8f15d03d7e5)
	}
	return
}

func (__obf_a1d3ecfbcf6e018c *ORM) Exists(__obf_8d3f22b980db2ac8, __obf_62fa9048d8479e91 string, __obf_9f04002b942a6c82 ...any) bool {
	var __obf_a0da796cbdcaac9c bool
	_ = __obf_a1d3ecfbcf6e018c.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_8d3f22b980db2ac8, __obf_62fa9048d8479e91), __obf_9f04002b942a6c82...).Scan(&__obf_a0da796cbdcaac9c)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_a0da796cbdcaac9c
}

func (__obf_a1d3ecfbcf6e018c *ORM) SaveModel(__obf_23bfa2fbe722cc6b any, __obf_d9b8e4b64be2f21e string) (string, error) {
	__obf_53bf414fffa45e91 := reflect.ValueOf(__obf_23bfa2fbe722cc6b)
	if __obf_53bf414fffa45e91.Kind() == reflect.Pointer {
		__obf_53bf414fffa45e91 = __obf_53bf414fffa45e91.Elem()
	}
	var (
		__obf_694937eb9026f025 = __obf_53bf414fffa45e91.FieldByName("Id")
		__obf_ce07a7d9c547c556 = util.ToSnake(reflect.Indirect(__obf_53bf414fffa45e91).Type().Name())
		__obf_16dbf8405d413a99 string
		__obf_9f04002b942a6c82 map[string]any
		__obf_f5c7d5fde3143a52 error
	)

	if __obf_694937eb9026f025.IsZero() {
		// insert
		__obf_16dbf8405d413a99, __obf_9f04002b942a6c82, __obf_f5c7d5fde3143a52 = BuildInsertSQL(__obf_23bfa2fbe722cc6b, __obf_ce07a7d9c547c556)
		if __obf_f5c7d5fde3143a52 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_f5c7d5fde3143a52.Error())
		}
		if __obf_d9b8e4b64be2f21e == "" {
			__obf_d9b8e4b64be2f21e = util.StringUUID()
		}

		__obf_9f04002b942a6c82["id"] = __obf_d9b8e4b64be2f21e
		__obf_7fcc1d8079af052b := time.Now().Unix()
		if _, __obf_5dddaed8aec227de := __obf_9f04002b942a6c82[__obf_a1d3ecfbcf6e018c.Option.Reserved.CreateTime]; __obf_5dddaed8aec227de {
			__obf_9f04002b942a6c82[__obf_a1d3ecfbcf6e018c.Option.Reserved.CreateTime] = __obf_7fcc1d8079af052b
		}
		if _, __obf_5dddaed8aec227de := __obf_9f04002b942a6c82[__obf_a1d3ecfbcf6e018c.Option.Reserved.UpdateTime]; __obf_5dddaed8aec227de {
			__obf_9f04002b942a6c82[__obf_a1d3ecfbcf6e018c.Option.Reserved.UpdateTime] = __obf_7fcc1d8079af052b
		}
	} else {
		// update
		__obf_16dbf8405d413a99, __obf_9f04002b942a6c82, __obf_f5c7d5fde3143a52 = BuildUpdateSQL(__obf_23bfa2fbe722cc6b, __obf_ce07a7d9c547c556, __obf_a1d3ecfbcf6e018c.Option.Reserved.UpdateTime, []string{__obf_a1d3ecfbcf6e018c.Option.Reserved.ID}, __obf_a1d3ecfbcf6e018c.Option.Reserved.CreateTime, __obf_a1d3ecfbcf6e018c.Option.Reserved.NO)
		if __obf_f5c7d5fde3143a52 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_f5c7d5fde3143a52.Error())
		}
	}

	if __obf_16dbf8405d413a99 != "" {
		_, __obf_f5c7d5fde3143a52 := __obf_a1d3ecfbcf6e018c.DB.NamedExec(__obf_16dbf8405d413a99, __obf_9f04002b942a6c82)
		if __obf_f5c7d5fde3143a52 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_f5c7d5fde3143a52.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_16dbf8405d413a99, "args", __obf_9f04002b942a6c82)
	}

	return __obf_d9b8e4b64be2f21e, nil

}

func (__obf_a1d3ecfbcf6e018c *ORM) SaveModelWidthAutoID(__obf_23bfa2fbe722cc6b any, __obf_d9b8e4b64be2f21e int64) (int64, error) {
	// query, args := modelBind(db.Option, m, insertId)
	__obf_53bf414fffa45e91 := reflect.ValueOf(__obf_23bfa2fbe722cc6b)
	if __obf_53bf414fffa45e91.Kind() == reflect.Pointer {
		__obf_53bf414fffa45e91 = __obf_53bf414fffa45e91.Elem()
	}
	__obf_694937eb9026f025 := __obf_53bf414fffa45e91.FieldByName("Id")
	__obf_ce07a7d9c547c556 := util.ToSnake(reflect.Indirect(__obf_53bf414fffa45e91).Type().Name())
	var (
		__obf_16dbf8405d413a99 string
		__obf_9f04002b942a6c82 map[string]any
		__obf_f5c7d5fde3143a52 error
	)

	if __obf_694937eb9026f025.IsZero() {
		// insert
		__obf_16dbf8405d413a99, __obf_9f04002b942a6c82, __obf_f5c7d5fde3143a52 = BuildInsertSQL(__obf_23bfa2fbe722cc6b, __obf_ce07a7d9c547c556)
		if __obf_f5c7d5fde3143a52 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_f5c7d5fde3143a52.Error())
		}
		if __obf_d9b8e4b64be2f21e != 0 {
			__obf_9f04002b942a6c82["id"] = __obf_d9b8e4b64be2f21e
		}
		__obf_7fcc1d8079af052b := time.Now().Unix()
		if _, __obf_5dddaed8aec227de := __obf_9f04002b942a6c82[__obf_a1d3ecfbcf6e018c.Option.Reserved.CreateTime]; __obf_5dddaed8aec227de {
			__obf_9f04002b942a6c82[__obf_a1d3ecfbcf6e018c.Option.Reserved.CreateTime] = __obf_7fcc1d8079af052b
		}
		if _, __obf_5dddaed8aec227de := __obf_9f04002b942a6c82[__obf_a1d3ecfbcf6e018c.Option.Reserved.UpdateTime]; __obf_5dddaed8aec227de {
			__obf_9f04002b942a6c82[__obf_a1d3ecfbcf6e018c.Option.Reserved.UpdateTime] = __obf_7fcc1d8079af052b
		}
	} else {
		// update
		__obf_16dbf8405d413a99, __obf_9f04002b942a6c82, __obf_f5c7d5fde3143a52 = BuildUpdateSQL(__obf_23bfa2fbe722cc6b, __obf_ce07a7d9c547c556, __obf_a1d3ecfbcf6e018c.Option.Reserved.UpdateTime, []string{__obf_a1d3ecfbcf6e018c.Option.Reserved.ID}, __obf_a1d3ecfbcf6e018c.Option.Reserved.CreateTime, __obf_a1d3ecfbcf6e018c.Option.Reserved.NO)
		if __obf_f5c7d5fde3143a52 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_f5c7d5fde3143a52.Error())
		}
	}

	if __obf_16dbf8405d413a99 != "" {
		__obf_4fceda9cabeb1c4f, __obf_f5c7d5fde3143a52 := __obf_a1d3ecfbcf6e018c.DB.NamedExec(__obf_16dbf8405d413a99, __obf_9f04002b942a6c82)
		if __obf_f5c7d5fde3143a52 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_f5c7d5fde3143a52.Error())
		}
		if __obf_d9b8e4b64be2f21e == 0 {
			return __obf_4fceda9cabeb1c4f.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_16dbf8405d413a99, "args", __obf_9f04002b942a6c82)
	}

	return __obf_d9b8e4b64be2f21e, nil
}

func (__obf_a1d3ecfbcf6e018c *ORM) SaveTxModel(__obf_5a0e0f8d9e4d8c5a *sqlx.Tx, __obf_23bfa2fbe722cc6b any, __obf_d9b8e4b64be2f21e string) error {
	var __obf_16dbf8405d413a99 string
	var __obf_9f04002b942a6c82 map[string]any
	__obf_16dbf8405d413a99, __obf_9f04002b942a6c82 = __obf_e6250472d56428fb(__obf_a1d3ecfbcf6e018c.Option, __obf_23bfa2fbe722cc6b, __obf_d9b8e4b64be2f21e)

	if __obf_16dbf8405d413a99 != "" {
		if _, __obf_f5c7d5fde3143a52 := __obf_5a0e0f8d9e4d8c5a.NamedExec(__obf_16dbf8405d413a99, __obf_9f04002b942a6c82); __obf_f5c7d5fde3143a52 != nil {
			_ = __obf_5a0e0f8d9e4d8c5a.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_f5c7d5fde3143a52)
		}

		slog.Debug("SaveTxModel", "query", __obf_16dbf8405d413a99, "args", __obf_9f04002b942a6c82)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_e6250472d56428fb[RecordID int | string](__obf_c9efa06b0a7790ff *Options, __obf_23bfa2fbe722cc6b any, __obf_d9b8e4b64be2f21e RecordID) (__obf_d7756dd017912e73 string, __obf_9f04002b942a6c82 map[string]any) {
	__obf_53bf414fffa45e91 := reflect.ValueOf(__obf_23bfa2fbe722cc6b)
	if __obf_53bf414fffa45e91.Kind() == reflect.Pointer {
		__obf_53bf414fffa45e91 = __obf_53bf414fffa45e91.Elem()
	}
	var __obf_051befe23d6a222b []string
	__obf_9f04002b942a6c82 = make(map[string]any)

	__obf_694937eb9026f025 := __obf_53bf414fffa45e91.FieldByName("Id")
	if __obf_694937eb9026f025.IsZero() {
		// 判断inserId是否为零值
		if __obf_d9b8e4b64be2f21e != *new(RecordID) {
			__obf_051befe23d6a222b = append(__obf_051befe23d6a222b, __obf_c9efa06b0a7790ff.Reserved.ID)
			__obf_9f04002b942a6c82[__obf_c9efa06b0a7790ff.Reserved.ID] = __obf_d9b8e4b64be2f21e
		} else {
			if __obf_694937eb9026f025.Kind() == reflect.String {
				__obf_051befe23d6a222b = append(__obf_051befe23d6a222b, __obf_c9efa06b0a7790ff.Reserved.ID)
				__obf_9f04002b942a6c82[__obf_c9efa06b0a7790ff.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}

		__obf_a8e58f2688d30de0 := reflect.Indirect(__obf_53bf414fffa45e91).Type()
		for __obf_cbebdbf7ebca6a65 := range __obf_a8e58f2688d30de0.NumField() {
			if __obf_b0142d337517f48f := __obf_a8e58f2688d30de0.Field(__obf_cbebdbf7ebca6a65); __obf_b0142d337517f48f.Tag.Get("db") != "-" {
				switch __obf_0381d5dba75b9958 := util.ToSnake(__obf_b0142d337517f48f.Name); __obf_0381d5dba75b9958 {
				case __obf_c9efa06b0a7790ff.Reserved.CreateTime, __obf_c9efa06b0a7790ff.Reserved.UpdateTime:
					__obf_051befe23d6a222b = append(__obf_051befe23d6a222b, __obf_0381d5dba75b9958)
					__obf_9f04002b942a6c82[__obf_0381d5dba75b9958] = time.Now().Unix()
				default:
					__obf_145a2e7072310087 := __obf_53bf414fffa45e91.Field(__obf_cbebdbf7ebca6a65).Interface()
					if strings.HasPrefix(__obf_0381d5dba75b9958, "is_") || __obf_145a2e7072310087 != reflect.Zero(__obf_b0142d337517f48f.Type).Interface() {
						__obf_051befe23d6a222b = append(__obf_051befe23d6a222b, __obf_0381d5dba75b9958)
						__obf_9f04002b942a6c82[__obf_0381d5dba75b9958] = __obf_145a2e7072310087
					}
				}
			}
		}
		__obf_d7756dd017912e73 = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_a8e58f2688d30de0.Name()), strings.Join(__obf_051befe23d6a222b, ","), strings.Join(__obf_051befe23d6a222b, ",:"))
	} else {
		__obf_a8e58f2688d30de0 := reflect.Indirect(__obf_53bf414fffa45e91).Type()
		for __obf_cbebdbf7ebca6a65 := range __obf_a8e58f2688d30de0.NumField() {
			if __obf_b0142d337517f48f := __obf_a8e58f2688d30de0.Field(__obf_cbebdbf7ebca6a65); __obf_b0142d337517f48f.Tag.Get("db") != "-" {
				__obf_145a2e7072310087 := __obf_53bf414fffa45e91.Field(__obf_cbebdbf7ebca6a65).Interface()
				if __obf_0381d5dba75b9958 := util.ToSnake(__obf_b0142d337517f48f.Name); strings.HasPrefix(__obf_0381d5dba75b9958, "is_") || __obf_145a2e7072310087 != reflect.Zero(__obf_b0142d337517f48f.Type).Interface() {
					__obf_051befe23d6a222b = append(__obf_051befe23d6a222b, fmt.Sprintf("%s=:%s", __obf_0381d5dba75b9958, __obf_0381d5dba75b9958))
					__obf_9f04002b942a6c82[__obf_0381d5dba75b9958] = __obf_145a2e7072310087
				}
			}
		}

		if util.HasField(__obf_53bf414fffa45e91, util.ToCamel(__obf_c9efa06b0a7790ff.Reserved.UpdateTime)) {
			__obf_9db70216f8cd8df8 := __obf_9f04002b942a6c82[__obf_c9efa06b0a7790ff.Reserved.UpdateTime]
			__obf_9f04002b942a6c82[__obf_c9efa06b0a7790ff.Reserved.UpdateTime] = time.Now().Unix()
			__obf_d7756dd017912e73 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_a8e58f2688d30de0.Name()), strings.Join(__obf_051befe23d6a222b, ","), __obf_c9efa06b0a7790ff.Reserved.ID, __obf_c9efa06b0a7790ff.Reserved.ID, __obf_c9efa06b0a7790ff.Reserved.UpdateTime, __obf_9db70216f8cd8df8)
		} else {
			__obf_d7756dd017912e73 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_a8e58f2688d30de0.Name()), strings.Join(__obf_051befe23d6a222b, ","), __obf_c9efa06b0a7790ff.Reserved.ID, __obf_c9efa06b0a7790ff.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_e6d9d0f4617d24a3 any, __obf_ce07a7d9c547c556 string) (string, map[string]any, error) {
	if __obf_ce07a7d9c547c556 == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_9f04002b942a6c82 := make(map[string]any)
	__obf_051befe23d6a222b := []string{}

	__obf_f8002c9b449a4966 := func(__obf_ef67ad2b7525fc31 string, __obf_14f652161c66608a any, __obf_bf9d50e59ca4c405 bool) {
		__obf_ccdb10839eb41d10 := __obf_14f652161c66608a == nil || reflect.DeepEqual(__obf_14f652161c66608a, reflect.Zero(reflect.TypeOf(__obf_14f652161c66608a)).Interface())
		if __obf_bf9d50e59ca4c405 && __obf_ccdb10839eb41d10 {
			return
		}
		__obf_051befe23d6a222b = append(__obf_051befe23d6a222b, __obf_ef67ad2b7525fc31)
		__obf_9f04002b942a6c82[__obf_ef67ad2b7525fc31] = __obf_14f652161c66608a
	}

	switch __obf_145a2e7072310087 := __obf_e6d9d0f4617d24a3.(type) {
	case map[string]any:
		for __obf_3f7031f75dab588c, __obf_53bf414fffa45e91 := range __obf_145a2e7072310087 {
			__obf_f8002c9b449a4966(__obf_3f7031f75dab588c, __obf_53bf414fffa45e91, false)
		}

	default:
		__obf_23f8e72519d9c06e := reflect.TypeOf(__obf_e6d9d0f4617d24a3)
		if __obf_23f8e72519d9c06e.Kind() == reflect.Pointer {
			__obf_23f8e72519d9c06e = __obf_23f8e72519d9c06e.Elem()
		}
		__obf_53bf414fffa45e91 := reflect.ValueOf(__obf_e6d9d0f4617d24a3)
		if __obf_53bf414fffa45e91.Kind() == reflect.Pointer {
			__obf_53bf414fffa45e91 = __obf_53bf414fffa45e91.Elem()
		}

		if __obf_23f8e72519d9c06e.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_cbebdbf7ebca6a65 := 0; __obf_cbebdbf7ebca6a65 < __obf_23f8e72519d9c06e.NumField(); __obf_cbebdbf7ebca6a65++ {
			__obf_da784e30d17ffae2 := __obf_23f8e72519d9c06e.Field(__obf_cbebdbf7ebca6a65).Tag.Get("db")
			if __obf_da784e30d17ffae2 == "" || __obf_da784e30d17ffae2 == "-" {
				continue
			}

			__obf_edd59ccb67c86dc5 := strings.Split(__obf_da784e30d17ffae2, ",")
			__obf_f8002c9b449a4966(__obf_edd59ccb67c86dc5[0], __obf_53bf414fffa45e91.Field(__obf_cbebdbf7ebca6a65).Interface(), len(__obf_edd59ccb67c86dc5) > 1 && __obf_edd59ccb67c86dc5[1] == "omitempty")
		}
	}

	if len(__obf_051befe23d6a222b) == 0 {
		return "", nil, errors.New("no fields to insert")
	}

	__obf_16dbf8405d413a99 := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)",
		__obf_ce07a7d9c547c556,
		strings.Join(__obf_051befe23d6a222b, ","),
		strings.Join(__obf_051befe23d6a222b, ",:"),
	)

	return __obf_16dbf8405d413a99, __obf_9f04002b942a6c82, nil
}

// pkField 主键字段
// lockField 乐观锁字段(默认为updated_at)
// constField 固定不变字段
func BuildUpdateSQL(
	__obf_e6d9d0f4617d24a3 any,
	__obf_ce07a7d9c547c556, __obf_3546e3887b8d4f54 string,
	__obf_4ad6f7bcfc201a17 []string,
	__obf_e9f483fbcffceda7 ...string,
) (string, map[string]any, error) {
	if __obf_ce07a7d9c547c556 == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_bb1772b8842c9397 := make(map[string]struct{})
	for _, __obf_c91c281845ae8c3a := range __obf_4ad6f7bcfc201a17 {
		__obf_bb1772b8842c9397[__obf_c91c281845ae8c3a] = struct{}{}
	}

	__obf_9f04002b942a6c82 := make(map[string]any)
	__obf_051befe23d6a222b := []string{}
	__obf_047a1835b6720a89 := []string{}

	__obf_f8002c9b449a4966 := func(__obf_ef67ad2b7525fc31 string, __obf_14f652161c66608a any, __obf_bf9d50e59ca4c405 bool) {
		__obf_ccdb10839eb41d10 := __obf_14f652161c66608a == nil || reflect.DeepEqual(__obf_14f652161c66608a, reflect.Zero(reflect.TypeOf(__obf_14f652161c66608a)).Interface())

		switch {
		case func() bool {
			_, __obf_5dddaed8aec227de := __obf_bb1772b8842c9397[__obf_ef67ad2b7525fc31]
			return __obf_5dddaed8aec227de
		}():
			__obf_047a1835b6720a89 = append(__obf_047a1835b6720a89, fmt.Sprintf("%s=:%s", __obf_ef67ad2b7525fc31, __obf_ef67ad2b7525fc31))
			__obf_9f04002b942a6c82[__obf_ef67ad2b7525fc31] = __obf_14f652161c66608a

		case slices.Contains(__obf_e9f483fbcffceda7, __obf_ef67ad2b7525fc31):
			// skip constField
		case __obf_ef67ad2b7525fc31 == __obf_3546e3887b8d4f54:
			const __obf_4a526e8ab11c4811 = "RESERVED_LOCK"
			__obf_051befe23d6a222b = append(__obf_051befe23d6a222b, fmt.Sprintf("%s=:%s", __obf_ef67ad2b7525fc31, __obf_ef67ad2b7525fc31))
			__obf_047a1835b6720a89 = append(__obf_047a1835b6720a89, fmt.Sprintf("%s=:%s", __obf_ef67ad2b7525fc31, __obf_4a526e8ab11c4811))
			__obf_9f04002b942a6c82[__obf_ef67ad2b7525fc31] = time.Now().Unix()
			__obf_9f04002b942a6c82[__obf_4a526e8ab11c4811] = __obf_14f652161c66608a

		default:
			if __obf_bf9d50e59ca4c405 && __obf_ccdb10839eb41d10 {
				return
			}
			__obf_051befe23d6a222b = append(__obf_051befe23d6a222b, fmt.Sprintf("%s=:%s", __obf_ef67ad2b7525fc31, __obf_ef67ad2b7525fc31))
			__obf_9f04002b942a6c82[__obf_ef67ad2b7525fc31] = __obf_14f652161c66608a
		}
	}

	switch __obf_145a2e7072310087 := __obf_e6d9d0f4617d24a3.(type) {
	case map[string]any:
		for __obf_3f7031f75dab588c, __obf_53bf414fffa45e91 := range __obf_145a2e7072310087 {
			__obf_f8002c9b449a4966(__obf_3f7031f75dab588c, __obf_53bf414fffa45e91, false)
		}

	default:

		__obf_23f8e72519d9c06e := reflect.TypeOf(__obf_e6d9d0f4617d24a3)
		if __obf_23f8e72519d9c06e.Kind() == reflect.Pointer {
			__obf_23f8e72519d9c06e = __obf_23f8e72519d9c06e.Elem()
		}
		__obf_53bf414fffa45e91 := reflect.ValueOf(__obf_e6d9d0f4617d24a3)
		if __obf_53bf414fffa45e91.Kind() == reflect.Pointer {
			__obf_53bf414fffa45e91 = __obf_53bf414fffa45e91.Elem()
		}
		if __obf_23f8e72519d9c06e.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_cbebdbf7ebca6a65 := 0; __obf_cbebdbf7ebca6a65 < __obf_23f8e72519d9c06e.NumField(); __obf_cbebdbf7ebca6a65++ {
			__obf_da784e30d17ffae2 := __obf_23f8e72519d9c06e.Field(__obf_cbebdbf7ebca6a65).Tag.Get("db")
			if __obf_da784e30d17ffae2 == "" || __obf_da784e30d17ffae2 == "-" {
				continue
			}

			__obf_edd59ccb67c86dc5 := strings.Split(__obf_da784e30d17ffae2, ",")
			__obf_f8002c9b449a4966(__obf_edd59ccb67c86dc5[0], __obf_53bf414fffa45e91.Field(__obf_cbebdbf7ebca6a65).Interface(), len(__obf_edd59ccb67c86dc5) > 1 && __obf_edd59ccb67c86dc5[1] == "omitempty")
		}
	}

	if len(__obf_4ad6f7bcfc201a17) == 0 || len(__obf_047a1835b6720a89) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_051befe23d6a222b) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	// if lockField != "" && !hasLockField {
	// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
	// }

	__obf_16dbf8405d413a99 := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		__obf_ce07a7d9c547c556,
		strings.Join(__obf_051befe23d6a222b, ","),
		strings.Join(__obf_047a1835b6720a89, " AND "),
	)

	return __obf_16dbf8405d413a99, __obf_9f04002b942a6c82, nil
}
