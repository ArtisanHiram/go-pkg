package __obf_c1f2c3d265c98f25

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

func NewORM(__obf_0cf6bdc83db39235 *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_0cf6bdc83db39235.Driver, __obf_0cf6bdc83db39235.DataSource)

	switch __obf_0cf6bdc83db39235.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_0cf6bdc83db39235.
			Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_0cf6bdc83db39235.
			Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_0cf6bdc83db39235.
			Placeholder = ":"
	case "sqlserver":
		__obf_0cf6bdc83db39235.
			Placeholder = "@"
	}

	if __obf_ab5192d9e0103059 := DB.Ping(); __obf_ab5192d9e0103059 != nil {
		panic(__obf_ab5192d9e0103059)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_0cf6bdc83db39235,
		}, func() {
			DB.Close()
		}
}

func (__obf_2e89f553fec85399 *ORM) ListMap(__obf_ce52679823fca375 *sqlx.Rows, __obf_9daef98e75d10801 func(map[string]any) (string, string)) (__obf_fda7a11c3287a559 []map[string]any, __obf_ab5192d9e0103059 error) {
	for __obf_ce52679823fca375.Next() {
		__obf_7b8ef7e43b17cb62 := make(map[string]any)
		if __obf_ab5192d9e0103059 = __obf_ce52679823fca375.MapScan(__obf_7b8ef7e43b17cb62); __obf_ab5192d9e0103059 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_ab5192d9e0103059)
		}
		for __obf_02cffdb2b73b0b57, __obf_8ec2f1c7b43e366d := range __obf_7b8ef7e43b17cb62 {
			switch __obf_8ec2f1c7b43e366d := __obf_8ec2f1c7b43e366d.(type) {
			case []uint8:
				__obf_7b8ef7e43b17cb62[__obf_02cffdb2b73b0b57] = string(__obf_8ec2f1c7b43e366d)
			case int64:
				__obf_7b8ef7e43b17cb62[__obf_02cffdb2b73b0b57] = int64(__obf_8ec2f1c7b43e366d)
			case uint64:
				__obf_7b8ef7e43b17cb62[__obf_02cffdb2b73b0b57] = uint64(__obf_8ec2f1c7b43e366d)
			}
		}
		if __obf_9daef98e75d10801 != nil {
			__obf_7dc1aeb3ab4ef0f1, __obf_d806e474f8bd0d56 := __obf_9daef98e75d10801(__obf_7b8ef7e43b17cb62)
			__obf_7b8ef7e43b17cb62[__obf_7dc1aeb3ab4ef0f1] = __obf_d806e474f8bd0d56
		}
		__obf_fda7a11c3287a559 = append(__obf_fda7a11c3287a559, __obf_7b8ef7e43b17cb62)
	}
	return
}

func (__obf_2e89f553fec85399 *ORM) Exists(__obf_9e542c5aed914f7c, __obf_b1004fa5302ae44f string, __obf_c2322d6ab4e7aa7f ...any) bool {
	var __obf_2c544b331e921525 bool
	_ = __obf_2e89f553fec85399.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_9e542c5aed914f7c, __obf_b1004fa5302ae44f), __obf_c2322d6ab4e7aa7f...).Scan(&__obf_2c544b331e921525)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_2c544b331e921525
}

func (__obf_2e89f553fec85399 *ORM) SaveModel(__obf_e099ce1bfb59c86e any, __obf_dd17ce38b1f29e8e string) (string, error) {
	__obf_8ec2f1c7b43e366d := reflect.ValueOf(__obf_e099ce1bfb59c86e)
	if __obf_8ec2f1c7b43e366d.Kind() == reflect.Pointer {
		__obf_8ec2f1c7b43e366d = __obf_8ec2f1c7b43e366d.Elem()
	}
	var (
		__obf_8c80580816dda1a1 = __obf_8ec2f1c7b43e366d.
					FieldByName("Id")
		__obf_18745a511b5910a3 = util.ToSnake(reflect.Indirect(__obf_8ec2f1c7b43e366d).Type().Name())
		__obf_8d780c51a7e1aba8 string
		__obf_c2322d6ab4e7aa7f map[string]any
		__obf_ab5192d9e0103059 error
	)

	if __obf_8c80580816dda1a1.IsZero() {
		__obf_8d780c51a7e1aba8,
			// insert
			__obf_c2322d6ab4e7aa7f, __obf_ab5192d9e0103059 = BuildInsertSQL(__obf_e099ce1bfb59c86e, __obf_18745a511b5910a3)
		if __obf_ab5192d9e0103059 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_ab5192d9e0103059.Error())
		}
		if __obf_dd17ce38b1f29e8e == "" {
			__obf_dd17ce38b1f29e8e = util.StringUUID()
		}
		__obf_c2322d6ab4e7aa7f["id"] = __obf_dd17ce38b1f29e8e
		__obf_eae50469245bb62e := time.Now().Unix()
		if _, __obf_0d4343a67e152792 := __obf_c2322d6ab4e7aa7f[__obf_2e89f553fec85399.Option.Reserved.CreateTime]; __obf_0d4343a67e152792 {
			__obf_c2322d6ab4e7aa7f[__obf_2e89f553fec85399.Option.Reserved.CreateTime] = __obf_eae50469245bb62e
		}
		if _, __obf_0d4343a67e152792 := __obf_c2322d6ab4e7aa7f[__obf_2e89f553fec85399.Option.Reserved.UpdateTime]; __obf_0d4343a67e152792 {
			__obf_c2322d6ab4e7aa7f[__obf_2e89f553fec85399.Option.Reserved.UpdateTime] = __obf_eae50469245bb62e
		}
	} else {
		__obf_8d780c51a7e1aba8,
			// update
			__obf_c2322d6ab4e7aa7f, __obf_ab5192d9e0103059 = BuildUpdateSQL(__obf_e099ce1bfb59c86e, __obf_18745a511b5910a3, __obf_2e89f553fec85399.Option.Reserved.UpdateTime, []string{__obf_2e89f553fec85399.Option.Reserved.ID}, __obf_2e89f553fec85399.Option.Reserved.CreateTime, __obf_2e89f553fec85399.Option.Reserved.NO)
		if __obf_ab5192d9e0103059 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_ab5192d9e0103059.Error())
		}
	}

	if __obf_8d780c51a7e1aba8 != "" {
		_, __obf_ab5192d9e0103059 := __obf_2e89f553fec85399.DB.NamedExec(__obf_8d780c51a7e1aba8, __obf_c2322d6ab4e7aa7f)
		if __obf_ab5192d9e0103059 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_ab5192d9e0103059.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_8d780c51a7e1aba8, "args", __obf_c2322d6ab4e7aa7f)
	}

	return __obf_dd17ce38b1f29e8e, nil

}

func (__obf_2e89f553fec85399 *ORM) SaveModelWidthAutoID(__obf_e099ce1bfb59c86e any, __obf_dd17ce38b1f29e8e int64) (int64, error) {
	__obf_8ec2f1c7b43e366d := // query, args := modelBind(db.Option, m, insertId)
		reflect.ValueOf(__obf_e099ce1bfb59c86e)
	if __obf_8ec2f1c7b43e366d.Kind() == reflect.Pointer {
		__obf_8ec2f1c7b43e366d = __obf_8ec2f1c7b43e366d.Elem()
	}
	__obf_8c80580816dda1a1 := __obf_8ec2f1c7b43e366d.FieldByName("Id")
	__obf_18745a511b5910a3 := util.ToSnake(reflect.Indirect(__obf_8ec2f1c7b43e366d).Type().Name())
	var (
		__obf_8d780c51a7e1aba8 string
		__obf_c2322d6ab4e7aa7f map[string]any
		__obf_ab5192d9e0103059 error
	)

	if __obf_8c80580816dda1a1.IsZero() {
		__obf_8d780c51a7e1aba8,
			// insert
			__obf_c2322d6ab4e7aa7f, __obf_ab5192d9e0103059 = BuildInsertSQL(__obf_e099ce1bfb59c86e, __obf_18745a511b5910a3)
		if __obf_ab5192d9e0103059 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_ab5192d9e0103059.Error())
		}
		if __obf_dd17ce38b1f29e8e != 0 {
			__obf_c2322d6ab4e7aa7f["id"] = __obf_dd17ce38b1f29e8e
		}
		__obf_eae50469245bb62e := time.Now().Unix()
		if _, __obf_0d4343a67e152792 := __obf_c2322d6ab4e7aa7f[__obf_2e89f553fec85399.Option.Reserved.CreateTime]; __obf_0d4343a67e152792 {
			__obf_c2322d6ab4e7aa7f[__obf_2e89f553fec85399.Option.Reserved.CreateTime] = __obf_eae50469245bb62e
		}
		if _, __obf_0d4343a67e152792 := __obf_c2322d6ab4e7aa7f[__obf_2e89f553fec85399.Option.Reserved.UpdateTime]; __obf_0d4343a67e152792 {
			__obf_c2322d6ab4e7aa7f[__obf_2e89f553fec85399.Option.Reserved.UpdateTime] = __obf_eae50469245bb62e
		}
	} else {
		__obf_8d780c51a7e1aba8,
			// update
			__obf_c2322d6ab4e7aa7f, __obf_ab5192d9e0103059 = BuildUpdateSQL(__obf_e099ce1bfb59c86e, __obf_18745a511b5910a3, __obf_2e89f553fec85399.Option.Reserved.UpdateTime, []string{__obf_2e89f553fec85399.Option.Reserved.ID}, __obf_2e89f553fec85399.Option.Reserved.CreateTime, __obf_2e89f553fec85399.Option.Reserved.NO)
		if __obf_ab5192d9e0103059 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_ab5192d9e0103059.Error())
		}
	}

	if __obf_8d780c51a7e1aba8 != "" {
		__obf_d806e474f8bd0d56, __obf_ab5192d9e0103059 := __obf_2e89f553fec85399.DB.NamedExec(__obf_8d780c51a7e1aba8, __obf_c2322d6ab4e7aa7f)
		if __obf_ab5192d9e0103059 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_ab5192d9e0103059.Error())
		}
		if __obf_dd17ce38b1f29e8e == 0 {
			return __obf_d806e474f8bd0d56.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_8d780c51a7e1aba8, "args", __obf_c2322d6ab4e7aa7f)
	}

	return __obf_dd17ce38b1f29e8e, nil
}

func (__obf_2e89f553fec85399 *ORM) SaveTxModel(__obf_b210b34065b441df *sqlx.Tx, __obf_e099ce1bfb59c86e any, __obf_dd17ce38b1f29e8e string) error {
	var __obf_8d780c51a7e1aba8 string
	var __obf_c2322d6ab4e7aa7f map[string]any
	__obf_8d780c51a7e1aba8, __obf_c2322d6ab4e7aa7f = __obf_35fd5f843c38abde(__obf_2e89f553fec85399.Option, __obf_e099ce1bfb59c86e, __obf_dd17ce38b1f29e8e)

	if __obf_8d780c51a7e1aba8 != "" {
		if _, __obf_ab5192d9e0103059 := __obf_b210b34065b441df.NamedExec(__obf_8d780c51a7e1aba8, __obf_c2322d6ab4e7aa7f); __obf_ab5192d9e0103059 != nil {
			_ = __obf_b210b34065b441df.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_ab5192d9e0103059)
		}

		slog.Debug("SaveTxModel", "query", __obf_8d780c51a7e1aba8, "args", __obf_c2322d6ab4e7aa7f)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_35fd5f843c38abde[RecordID int | string](__obf_0cf6bdc83db39235 *Options, __obf_e099ce1bfb59c86e any, __obf_dd17ce38b1f29e8e RecordID) (__obf_effa7aaceac77c64 string, __obf_c2322d6ab4e7aa7f map[string]any) {
	__obf_8ec2f1c7b43e366d := reflect.ValueOf(__obf_e099ce1bfb59c86e)
	if __obf_8ec2f1c7b43e366d.Kind() == reflect.Pointer {
		__obf_8ec2f1c7b43e366d = __obf_8ec2f1c7b43e366d.Elem()
	}
	var __obf_a9cb12cc8c31ba71 []string
	__obf_c2322d6ab4e7aa7f = make(map[string]any)
	__obf_8c80580816dda1a1 := __obf_8ec2f1c7b43e366d.FieldByName("Id")
	if __obf_8c80580816dda1a1.IsZero() {
		// 判断inserId是否为零值
		if __obf_dd17ce38b1f29e8e != *new(RecordID) {
			__obf_a9cb12cc8c31ba71 = append(__obf_a9cb12cc8c31ba71, __obf_0cf6bdc83db39235.Reserved.ID)
			__obf_c2322d6ab4e7aa7f[__obf_0cf6bdc83db39235.Reserved.ID] = __obf_dd17ce38b1f29e8e
		} else {
			if __obf_8c80580816dda1a1.Kind() == reflect.String {
				__obf_a9cb12cc8c31ba71 = append(__obf_a9cb12cc8c31ba71, __obf_0cf6bdc83db39235.Reserved.ID)
				__obf_c2322d6ab4e7aa7f[__obf_0cf6bdc83db39235.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}
		__obf_3db23f3d8ae3b235 := reflect.Indirect(__obf_8ec2f1c7b43e366d).Type()
		for __obf_739dac239fb8287e := range __obf_3db23f3d8ae3b235.NumField() {
			if __obf_717e0741d033c9e0 := __obf_3db23f3d8ae3b235.Field(__obf_739dac239fb8287e); __obf_717e0741d033c9e0.Tag.Get("db") != "-" {
				switch __obf_2541904edefc8de7 := util.ToSnake(__obf_717e0741d033c9e0.Name); __obf_2541904edefc8de7 {
				case __obf_0cf6bdc83db39235.Reserved.CreateTime, __obf_0cf6bdc83db39235.Reserved.UpdateTime:
					__obf_a9cb12cc8c31ba71 = append(__obf_a9cb12cc8c31ba71, __obf_2541904edefc8de7)
					__obf_c2322d6ab4e7aa7f[__obf_2541904edefc8de7] = time.Now().Unix()
				default:
					__obf_e87e6f05cc491080 := __obf_8ec2f1c7b43e366d.Field(__obf_739dac239fb8287e).Interface()
					if strings.HasPrefix(__obf_2541904edefc8de7, "is_") || __obf_e87e6f05cc491080 != reflect.Zero(__obf_717e0741d033c9e0.Type).Interface() {
						__obf_a9cb12cc8c31ba71 = append(__obf_a9cb12cc8c31ba71, __obf_2541904edefc8de7)
						__obf_c2322d6ab4e7aa7f[__obf_2541904edefc8de7] = __obf_e87e6f05cc491080
					}
				}
			}
		}
		__obf_effa7aaceac77c64 = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_3db23f3d8ae3b235.Name()), strings.Join(__obf_a9cb12cc8c31ba71, ","), strings.Join(__obf_a9cb12cc8c31ba71, ",:"))
	} else {
		__obf_3db23f3d8ae3b235 := reflect.Indirect(__obf_8ec2f1c7b43e366d).Type()
		for __obf_739dac239fb8287e := range __obf_3db23f3d8ae3b235.NumField() {
			if __obf_717e0741d033c9e0 := __obf_3db23f3d8ae3b235.Field(__obf_739dac239fb8287e); __obf_717e0741d033c9e0.Tag.Get("db") != "-" {
				__obf_e87e6f05cc491080 := __obf_8ec2f1c7b43e366d.Field(__obf_739dac239fb8287e).Interface()
				if __obf_2541904edefc8de7 := util.ToSnake(__obf_717e0741d033c9e0.Name); strings.HasPrefix(__obf_2541904edefc8de7, "is_") || __obf_e87e6f05cc491080 != reflect.Zero(__obf_717e0741d033c9e0.Type).Interface() {
					__obf_a9cb12cc8c31ba71 = append(__obf_a9cb12cc8c31ba71, fmt.Sprintf("%s=:%s", __obf_2541904edefc8de7, __obf_2541904edefc8de7))
					__obf_c2322d6ab4e7aa7f[__obf_2541904edefc8de7] = __obf_e87e6f05cc491080
				}
			}
		}

		if util.HasField(__obf_8ec2f1c7b43e366d, util.ToCamel(__obf_0cf6bdc83db39235.Reserved.UpdateTime)) {
			__obf_482a5c678866ceca := __obf_c2322d6ab4e7aa7f[__obf_0cf6bdc83db39235.Reserved.UpdateTime]
			__obf_c2322d6ab4e7aa7f[__obf_0cf6bdc83db39235.Reserved.UpdateTime] = time.Now().Unix()
			__obf_effa7aaceac77c64 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_3db23f3d8ae3b235.Name()), strings.Join(__obf_a9cb12cc8c31ba71, ","), __obf_0cf6bdc83db39235.Reserved.ID, __obf_0cf6bdc83db39235.Reserved.ID, __obf_0cf6bdc83db39235.Reserved.UpdateTime, __obf_482a5c678866ceca)
		} else {
			__obf_effa7aaceac77c64 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_3db23f3d8ae3b235.Name()), strings.Join(__obf_a9cb12cc8c31ba71, ","), __obf_0cf6bdc83db39235.Reserved.ID, __obf_0cf6bdc83db39235.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_015cf9d133c2185a any, __obf_18745a511b5910a3 string) (string, map[string]any, error) {
	if __obf_18745a511b5910a3 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_c2322d6ab4e7aa7f := make(map[string]any)
	__obf_a9cb12cc8c31ba71 := []string{}
	__obf_aecb6e2ba975aa08 := func(__obf_fbae619a1e6618fb string, __obf_1fe888a8765085d6 any, __obf_dc797d536b8c67e2 bool) {
		__obf_bbbe764efe28f7c7 := __obf_1fe888a8765085d6 == nil || reflect.DeepEqual(__obf_1fe888a8765085d6, reflect.Zero(reflect.TypeOf(__obf_1fe888a8765085d6)).Interface())
		if __obf_dc797d536b8c67e2 && __obf_bbbe764efe28f7c7 {
			return
		}
		__obf_a9cb12cc8c31ba71 = append(__obf_a9cb12cc8c31ba71, __obf_fbae619a1e6618fb)
		__obf_c2322d6ab4e7aa7f[__obf_fbae619a1e6618fb] = __obf_1fe888a8765085d6
	}

	switch __obf_e87e6f05cc491080 := __obf_015cf9d133c2185a.(type) {
	case map[string]any:
		for __obf_02cffdb2b73b0b57, __obf_8ec2f1c7b43e366d := range __obf_e87e6f05cc491080 {
			__obf_aecb6e2ba975aa08(__obf_02cffdb2b73b0b57, __obf_8ec2f1c7b43e366d, false)
		}

	default:
		__obf_4829520f0a5c1ebb := reflect.TypeOf(__obf_015cf9d133c2185a)
		if __obf_4829520f0a5c1ebb.Kind() == reflect.Pointer {
			__obf_4829520f0a5c1ebb = __obf_4829520f0a5c1ebb.Elem()
		}
		__obf_8ec2f1c7b43e366d := reflect.ValueOf(__obf_015cf9d133c2185a)
		if __obf_8ec2f1c7b43e366d.Kind() == reflect.Pointer {
			__obf_8ec2f1c7b43e366d = __obf_8ec2f1c7b43e366d.Elem()
		}

		if __obf_4829520f0a5c1ebb.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_739dac239fb8287e := 0; __obf_739dac239fb8287e < __obf_4829520f0a5c1ebb.NumField(); __obf_739dac239fb8287e++ {
			__obf_dfe1702d60a62240 := __obf_4829520f0a5c1ebb.Field(__obf_739dac239fb8287e).Tag.Get("db")
			if __obf_dfe1702d60a62240 == "" || __obf_dfe1702d60a62240 == "-" {
				continue
			}
			__obf_a3fc24a5357e10c2 := strings.Split(__obf_dfe1702d60a62240, ",")
			__obf_aecb6e2ba975aa08(__obf_a3fc24a5357e10c2[0], __obf_8ec2f1c7b43e366d.Field(__obf_739dac239fb8287e).Interface(), len(__obf_a3fc24a5357e10c2) > 1 && __obf_a3fc24a5357e10c2[1] == "omitempty")
		}
	}

	if len(__obf_a9cb12cc8c31ba71) == 0 {
		return "", nil, errors.New("no fields to insert")
	}
	__obf_8d780c51a7e1aba8 := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", __obf_18745a511b5910a3, strings.Join(__obf_a9cb12cc8c31ba71, ","),
		strings.Join(__obf_a9cb12cc8c31ba71, ",:"),
	)

	return __obf_8d780c51a7e1aba8,

		// pkField 主键字段
		// lockField 乐观锁字段(默认为updated_at)
		// constField 固定不变字段
		__obf_c2322d6ab4e7aa7f, nil
}

func BuildUpdateSQL(__obf_015cf9d133c2185a any, __obf_18745a511b5910a3, __obf_1cee4673a45ca930 string, __obf_e3229b351ecffb7e []string, __obf_9bd67cbce193e2f8 ...string,
) (string, map[string]any, error) {
	if __obf_18745a511b5910a3 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_30626c8d607dde04 := make(map[string]struct{})
	for _, __obf_eebcb9589b8371f6 := range __obf_e3229b351ecffb7e {
		__obf_30626c8d607dde04[__obf_eebcb9589b8371f6] = struct{}{}
	}
	__obf_c2322d6ab4e7aa7f := make(map[string]any)
	__obf_a9cb12cc8c31ba71 := []string{}
	__obf_ade89aa779c34fa6 := []string{}
	__obf_aecb6e2ba975aa08 := func(__obf_fbae619a1e6618fb string, __obf_1fe888a8765085d6 any, __obf_dc797d536b8c67e2 bool) {
		__obf_bbbe764efe28f7c7 := __obf_1fe888a8765085d6 == nil || reflect.DeepEqual(__obf_1fe888a8765085d6, reflect.Zero(reflect.TypeOf(__obf_1fe888a8765085d6)).Interface())

		switch {
		case func() bool {
			_, __obf_0d4343a67e152792 := __obf_30626c8d607dde04[__obf_fbae619a1e6618fb]
			return __obf_0d4343a67e152792
		}():
			__obf_ade89aa779c34fa6 = append(__obf_ade89aa779c34fa6, fmt.Sprintf("%s=:%s", __obf_fbae619a1e6618fb, __obf_fbae619a1e6618fb))
			__obf_c2322d6ab4e7aa7f[__obf_fbae619a1e6618fb] = __obf_1fe888a8765085d6

		case slices.Contains(__obf_9bd67cbce193e2f8, __obf_fbae619a1e6618fb):
			// skip constField
		case __obf_fbae619a1e6618fb == __obf_1cee4673a45ca930:
			const __obf_6491589c0e591377 = "RESERVED_LOCK"
			__obf_a9cb12cc8c31ba71 = append(__obf_a9cb12cc8c31ba71, fmt.Sprintf("%s=:%s", __obf_fbae619a1e6618fb, __obf_fbae619a1e6618fb))
			__obf_ade89aa779c34fa6 = append(__obf_ade89aa779c34fa6, fmt.Sprintf("%s=:%s", __obf_fbae619a1e6618fb, __obf_6491589c0e591377))
			__obf_c2322d6ab4e7aa7f[__obf_fbae619a1e6618fb] = time.Now().Unix()
			__obf_c2322d6ab4e7aa7f[__obf_6491589c0e591377] = __obf_1fe888a8765085d6

		default:
			if __obf_dc797d536b8c67e2 && __obf_bbbe764efe28f7c7 {
				return
			}
			__obf_a9cb12cc8c31ba71 = append(__obf_a9cb12cc8c31ba71, fmt.Sprintf("%s=:%s", __obf_fbae619a1e6618fb, __obf_fbae619a1e6618fb))
			__obf_c2322d6ab4e7aa7f[__obf_fbae619a1e6618fb] = __obf_1fe888a8765085d6
		}
	}

	switch __obf_e87e6f05cc491080 := __obf_015cf9d133c2185a.(type) {
	case map[string]any:
		for __obf_02cffdb2b73b0b57, __obf_8ec2f1c7b43e366d := range __obf_e87e6f05cc491080 {
			__obf_aecb6e2ba975aa08(__obf_02cffdb2b73b0b57, __obf_8ec2f1c7b43e366d, false)
		}

	default:
		__obf_4829520f0a5c1ebb := reflect.TypeOf(__obf_015cf9d133c2185a)
		if __obf_4829520f0a5c1ebb.Kind() == reflect.Pointer {
			__obf_4829520f0a5c1ebb = __obf_4829520f0a5c1ebb.Elem()
		}
		__obf_8ec2f1c7b43e366d := reflect.ValueOf(__obf_015cf9d133c2185a)
		if __obf_8ec2f1c7b43e366d.Kind() == reflect.Pointer {
			__obf_8ec2f1c7b43e366d = __obf_8ec2f1c7b43e366d.Elem()
		}
		if __obf_4829520f0a5c1ebb.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_739dac239fb8287e := 0; __obf_739dac239fb8287e < __obf_4829520f0a5c1ebb.NumField(); __obf_739dac239fb8287e++ {
			__obf_dfe1702d60a62240 := __obf_4829520f0a5c1ebb.Field(__obf_739dac239fb8287e).Tag.Get("db")
			if __obf_dfe1702d60a62240 == "" || __obf_dfe1702d60a62240 == "-" {
				continue
			}
			__obf_a3fc24a5357e10c2 := strings.Split(__obf_dfe1702d60a62240, ",")
			__obf_aecb6e2ba975aa08(__obf_a3fc24a5357e10c2[0], __obf_8ec2f1c7b43e366d.Field(__obf_739dac239fb8287e).Interface(), len(__obf_a3fc24a5357e10c2) > 1 && __obf_a3fc24a5357e10c2[1] == "omitempty")
		}
	}

	if len(__obf_e3229b351ecffb7e) == 0 || len(__obf_ade89aa779c34fa6) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_a9cb12cc8c31ba71) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	__obf_8d780c51a7e1aba8 := // if lockField != "" && !hasLockField {
		// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
		// }

		fmt.Sprintf("UPDATE %s SET %s WHERE %s", __obf_18745a511b5910a3, strings.Join(__obf_a9cb12cc8c31ba71, ","),
			strings.Join(__obf_ade89aa779c34fa6, " AND "),
		)

	return __obf_8d780c51a7e1aba8, __obf_c2322d6ab4e7aa7f, nil
}
