package __obf_e471fc5ea32e9ac0

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

func NewORM(__obf_389b86a233c48607 *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_389b86a233c48607.Driver, __obf_389b86a233c48607.DataSource)

	switch __obf_389b86a233c48607.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_389b86a233c48607.Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_389b86a233c48607.Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_389b86a233c48607.Placeholder = ":"
	case "sqlserver":
		__obf_389b86a233c48607.Placeholder = "@"
	}

	if __obf_c5f1094aec829fc9 := DB.Ping(); __obf_c5f1094aec829fc9 != nil {
		panic(__obf_c5f1094aec829fc9)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_389b86a233c48607,
		}, func() {
			DB.Close()
		}
}

func (__obf_4fa371f8997ca48b *ORM) ListMap(__obf_31a48fd558fdaaab *sqlx.Rows, __obf_437bbdb792244c8c func(map[string]any) (string, string)) (__obf_5237801d32b5499f []map[string]any, __obf_c5f1094aec829fc9 error) {
	for __obf_31a48fd558fdaaab.Next() {
		__obf_1953229484593f0a := make(map[string]any)
		if __obf_c5f1094aec829fc9 = __obf_31a48fd558fdaaab.MapScan(__obf_1953229484593f0a); __obf_c5f1094aec829fc9 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_c5f1094aec829fc9)
		}
		for __obf_12c89da7507a8f69, __obf_be4148454b0decfa := range __obf_1953229484593f0a {
			switch __obf_be4148454b0decfa := __obf_be4148454b0decfa.(type) {
			case []uint8:
				__obf_1953229484593f0a[__obf_12c89da7507a8f69] = string(__obf_be4148454b0decfa)
			case int64:
				__obf_1953229484593f0a[__obf_12c89da7507a8f69] = int64(__obf_be4148454b0decfa)
			case uint64:
				__obf_1953229484593f0a[__obf_12c89da7507a8f69] = uint64(__obf_be4148454b0decfa)
			}
		}
		if __obf_437bbdb792244c8c != nil {
			__obf_8696f27aedea956f, __obf_57b28d01863df3c3 := __obf_437bbdb792244c8c(__obf_1953229484593f0a)
			__obf_1953229484593f0a[__obf_8696f27aedea956f] = __obf_57b28d01863df3c3
		}
		__obf_5237801d32b5499f = append(__obf_5237801d32b5499f, __obf_1953229484593f0a)
	}
	return
}

func (__obf_4fa371f8997ca48b *ORM) Exists(__obf_6ea7a22c2dfa5a70, __obf_5acd8865f259d22f string, __obf_c2803eea3fa184fb ...any) bool {
	var __obf_1fa9eefc7b96e746 bool
	_ = __obf_4fa371f8997ca48b.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_6ea7a22c2dfa5a70, __obf_5acd8865f259d22f), __obf_c2803eea3fa184fb...).Scan(&__obf_1fa9eefc7b96e746)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_1fa9eefc7b96e746
}

func (__obf_4fa371f8997ca48b *ORM) SaveModel(__obf_4ae26b2f515a7b6b any, __obf_8998a9a956e35282 string) (string, error) {
	__obf_be4148454b0decfa := reflect.ValueOf(__obf_4ae26b2f515a7b6b)
	if __obf_be4148454b0decfa.Kind() == reflect.Pointer {
		__obf_be4148454b0decfa = __obf_be4148454b0decfa.Elem()
	}
	var (
		__obf_aab98d0ca2054f7d = __obf_be4148454b0decfa.FieldByName("Id")
		__obf_6908758718a006fd = util.ToSnake(reflect.Indirect(__obf_be4148454b0decfa).Type().Name())
		__obf_3d2a5afadabd5bfc string
		__obf_c2803eea3fa184fb map[string]any
		__obf_c5f1094aec829fc9 error
	)

	if __obf_aab98d0ca2054f7d.IsZero() {
		// insert
		__obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb, __obf_c5f1094aec829fc9 = BuildInsertSQL(__obf_4ae26b2f515a7b6b, __obf_6908758718a006fd)
		if __obf_c5f1094aec829fc9 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_c5f1094aec829fc9.Error())
		}
		if __obf_8998a9a956e35282 == "" {
			__obf_8998a9a956e35282 = util.StringUUID()
		}

		__obf_c2803eea3fa184fb["id"] = __obf_8998a9a956e35282
		__obf_d84635f328e5442b := time.Now().Unix()
		if _, __obf_f65665999cfd53f1 := __obf_c2803eea3fa184fb[__obf_4fa371f8997ca48b.Option.Reserved.CreateTime]; __obf_f65665999cfd53f1 {
			__obf_c2803eea3fa184fb[__obf_4fa371f8997ca48b.Option.Reserved.CreateTime] = __obf_d84635f328e5442b
		}
		if _, __obf_f65665999cfd53f1 := __obf_c2803eea3fa184fb[__obf_4fa371f8997ca48b.Option.Reserved.UpdateTime]; __obf_f65665999cfd53f1 {
			__obf_c2803eea3fa184fb[__obf_4fa371f8997ca48b.Option.Reserved.UpdateTime] = __obf_d84635f328e5442b
		}
	} else {
		// update
		__obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb, __obf_c5f1094aec829fc9 = BuildUpdateSQL(__obf_4ae26b2f515a7b6b, __obf_6908758718a006fd, __obf_4fa371f8997ca48b.Option.Reserved.UpdateTime, []string{__obf_4fa371f8997ca48b.Option.Reserved.ID}, __obf_4fa371f8997ca48b.Option.Reserved.CreateTime, __obf_4fa371f8997ca48b.Option.Reserved.NO)
		if __obf_c5f1094aec829fc9 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_c5f1094aec829fc9.Error())
		}
	}

	if __obf_3d2a5afadabd5bfc != "" {
		_, __obf_c5f1094aec829fc9 := __obf_4fa371f8997ca48b.DB.NamedExec(__obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb)
		if __obf_c5f1094aec829fc9 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_c5f1094aec829fc9.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_3d2a5afadabd5bfc, "args", __obf_c2803eea3fa184fb)
	}

	return __obf_8998a9a956e35282, nil

}

func (__obf_4fa371f8997ca48b *ORM) SaveModelWidthAutoID(__obf_4ae26b2f515a7b6b any, __obf_8998a9a956e35282 int64) (int64, error) {
	// query, args := modelBind(db.Option, m, insertId)
	__obf_be4148454b0decfa := reflect.ValueOf(__obf_4ae26b2f515a7b6b)
	if __obf_be4148454b0decfa.Kind() == reflect.Pointer {
		__obf_be4148454b0decfa = __obf_be4148454b0decfa.Elem()
	}
	__obf_aab98d0ca2054f7d := __obf_be4148454b0decfa.FieldByName("Id")
	__obf_6908758718a006fd := util.ToSnake(reflect.Indirect(__obf_be4148454b0decfa).Type().Name())
	var (
		__obf_3d2a5afadabd5bfc string
		__obf_c2803eea3fa184fb map[string]any
		__obf_c5f1094aec829fc9 error
	)

	if __obf_aab98d0ca2054f7d.IsZero() {
		// insert
		__obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb, __obf_c5f1094aec829fc9 = BuildInsertSQL(__obf_4ae26b2f515a7b6b, __obf_6908758718a006fd)
		if __obf_c5f1094aec829fc9 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_c5f1094aec829fc9.Error())
		}
		if __obf_8998a9a956e35282 != 0 {
			__obf_c2803eea3fa184fb["id"] = __obf_8998a9a956e35282
		}
		__obf_d84635f328e5442b := time.Now().Unix()
		if _, __obf_f65665999cfd53f1 := __obf_c2803eea3fa184fb[__obf_4fa371f8997ca48b.Option.Reserved.CreateTime]; __obf_f65665999cfd53f1 {
			__obf_c2803eea3fa184fb[__obf_4fa371f8997ca48b.Option.Reserved.CreateTime] = __obf_d84635f328e5442b
		}
		if _, __obf_f65665999cfd53f1 := __obf_c2803eea3fa184fb[__obf_4fa371f8997ca48b.Option.Reserved.UpdateTime]; __obf_f65665999cfd53f1 {
			__obf_c2803eea3fa184fb[__obf_4fa371f8997ca48b.Option.Reserved.UpdateTime] = __obf_d84635f328e5442b
		}
	} else {
		// update
		__obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb, __obf_c5f1094aec829fc9 = BuildUpdateSQL(__obf_4ae26b2f515a7b6b, __obf_6908758718a006fd, __obf_4fa371f8997ca48b.Option.Reserved.UpdateTime, []string{__obf_4fa371f8997ca48b.Option.Reserved.ID}, __obf_4fa371f8997ca48b.Option.Reserved.CreateTime, __obf_4fa371f8997ca48b.Option.Reserved.NO)
		if __obf_c5f1094aec829fc9 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_c5f1094aec829fc9.Error())
		}
	}

	if __obf_3d2a5afadabd5bfc != "" {
		__obf_57b28d01863df3c3, __obf_c5f1094aec829fc9 := __obf_4fa371f8997ca48b.DB.NamedExec(__obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb)
		if __obf_c5f1094aec829fc9 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_c5f1094aec829fc9.Error())
		}
		if __obf_8998a9a956e35282 == 0 {
			return __obf_57b28d01863df3c3.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_3d2a5afadabd5bfc, "args", __obf_c2803eea3fa184fb)
	}

	return __obf_8998a9a956e35282, nil
}

func (__obf_4fa371f8997ca48b *ORM) SaveTxModel(__obf_9e151d73a4e332b3 *sqlx.Tx, __obf_4ae26b2f515a7b6b any, __obf_8998a9a956e35282 string) error {
	var __obf_3d2a5afadabd5bfc string
	var __obf_c2803eea3fa184fb map[string]any
	__obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb = __obf_7985de248899c6f1(__obf_4fa371f8997ca48b.Option, __obf_4ae26b2f515a7b6b, __obf_8998a9a956e35282)

	if __obf_3d2a5afadabd5bfc != "" {
		if _, __obf_c5f1094aec829fc9 := __obf_9e151d73a4e332b3.NamedExec(__obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb); __obf_c5f1094aec829fc9 != nil {
			_ = __obf_9e151d73a4e332b3.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_c5f1094aec829fc9)
		}

		slog.Debug("SaveTxModel", "query", __obf_3d2a5afadabd5bfc, "args", __obf_c2803eea3fa184fb)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_7985de248899c6f1[RecordID int | string](__obf_389b86a233c48607 *Options, __obf_4ae26b2f515a7b6b any, __obf_8998a9a956e35282 RecordID) (__obf_b0877d3c4b761dfe string, __obf_c2803eea3fa184fb map[string]any) {
	__obf_be4148454b0decfa := reflect.ValueOf(__obf_4ae26b2f515a7b6b)
	if __obf_be4148454b0decfa.Kind() == reflect.Pointer {
		__obf_be4148454b0decfa = __obf_be4148454b0decfa.Elem()
	}
	var __obf_05d6f2cd4d407fe8 []string
	__obf_c2803eea3fa184fb = make(map[string]any)

	__obf_aab98d0ca2054f7d := __obf_be4148454b0decfa.FieldByName("Id")
	if __obf_aab98d0ca2054f7d.IsZero() {
		// 判断inserId是否为零值
		if __obf_8998a9a956e35282 != *new(RecordID) {
			__obf_05d6f2cd4d407fe8 = append(__obf_05d6f2cd4d407fe8, __obf_389b86a233c48607.Reserved.ID)
			__obf_c2803eea3fa184fb[__obf_389b86a233c48607.Reserved.ID] = __obf_8998a9a956e35282
		} else {
			if __obf_aab98d0ca2054f7d.Kind() == reflect.String {
				__obf_05d6f2cd4d407fe8 = append(__obf_05d6f2cd4d407fe8, __obf_389b86a233c48607.Reserved.ID)
				__obf_c2803eea3fa184fb[__obf_389b86a233c48607.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}

		__obf_6b48d44aa2f8b519 := reflect.Indirect(__obf_be4148454b0decfa).Type()
		for __obf_be83ea6d875ef30b := range __obf_6b48d44aa2f8b519.NumField() {
			if __obf_715a8e43dc438a16 := __obf_6b48d44aa2f8b519.Field(__obf_be83ea6d875ef30b); __obf_715a8e43dc438a16.Tag.Get("db") != "-" {
				switch __obf_2f18c866a1c51bcd := util.ToSnake(__obf_715a8e43dc438a16.Name); __obf_2f18c866a1c51bcd {
				case __obf_389b86a233c48607.Reserved.CreateTime, __obf_389b86a233c48607.Reserved.UpdateTime:
					__obf_05d6f2cd4d407fe8 = append(__obf_05d6f2cd4d407fe8, __obf_2f18c866a1c51bcd)
					__obf_c2803eea3fa184fb[__obf_2f18c866a1c51bcd] = time.Now().Unix()
				default:
					__obf_b6bf8b8bc11918d7 := __obf_be4148454b0decfa.Field(__obf_be83ea6d875ef30b).Interface()
					if strings.HasPrefix(__obf_2f18c866a1c51bcd, "is_") || __obf_b6bf8b8bc11918d7 != reflect.Zero(__obf_715a8e43dc438a16.Type).Interface() {
						__obf_05d6f2cd4d407fe8 = append(__obf_05d6f2cd4d407fe8, __obf_2f18c866a1c51bcd)
						__obf_c2803eea3fa184fb[__obf_2f18c866a1c51bcd] = __obf_b6bf8b8bc11918d7
					}
				}
			}
		}
		__obf_b0877d3c4b761dfe = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_6b48d44aa2f8b519.Name()), strings.Join(__obf_05d6f2cd4d407fe8, ","), strings.Join(__obf_05d6f2cd4d407fe8, ",:"))
	} else {
		__obf_6b48d44aa2f8b519 := reflect.Indirect(__obf_be4148454b0decfa).Type()
		for __obf_be83ea6d875ef30b := range __obf_6b48d44aa2f8b519.NumField() {
			if __obf_715a8e43dc438a16 := __obf_6b48d44aa2f8b519.Field(__obf_be83ea6d875ef30b); __obf_715a8e43dc438a16.Tag.Get("db") != "-" {
				__obf_b6bf8b8bc11918d7 := __obf_be4148454b0decfa.Field(__obf_be83ea6d875ef30b).Interface()
				if __obf_2f18c866a1c51bcd := util.ToSnake(__obf_715a8e43dc438a16.Name); strings.HasPrefix(__obf_2f18c866a1c51bcd, "is_") || __obf_b6bf8b8bc11918d7 != reflect.Zero(__obf_715a8e43dc438a16.Type).Interface() {
					__obf_05d6f2cd4d407fe8 = append(__obf_05d6f2cd4d407fe8, fmt.Sprintf("%s=:%s", __obf_2f18c866a1c51bcd, __obf_2f18c866a1c51bcd))
					__obf_c2803eea3fa184fb[__obf_2f18c866a1c51bcd] = __obf_b6bf8b8bc11918d7
				}
			}
		}

		if util.HasField(__obf_be4148454b0decfa, util.ToCamel(__obf_389b86a233c48607.Reserved.UpdateTime)) {
			__obf_c93b7055487aa339 := __obf_c2803eea3fa184fb[__obf_389b86a233c48607.Reserved.UpdateTime]
			__obf_c2803eea3fa184fb[__obf_389b86a233c48607.Reserved.UpdateTime] = time.Now().Unix()
			__obf_b0877d3c4b761dfe = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_6b48d44aa2f8b519.Name()), strings.Join(__obf_05d6f2cd4d407fe8, ","), __obf_389b86a233c48607.Reserved.ID, __obf_389b86a233c48607.Reserved.ID, __obf_389b86a233c48607.Reserved.UpdateTime, __obf_c93b7055487aa339)
		} else {
			__obf_b0877d3c4b761dfe = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_6b48d44aa2f8b519.Name()), strings.Join(__obf_05d6f2cd4d407fe8, ","), __obf_389b86a233c48607.Reserved.ID, __obf_389b86a233c48607.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_53999f33a514460f any, __obf_6908758718a006fd string) (string, map[string]any, error) {
	if __obf_6908758718a006fd == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_c2803eea3fa184fb := make(map[string]any)
	__obf_05d6f2cd4d407fe8 := []string{}

	__obf_44912f234fa2b844 := func(__obf_fb05cce5b14da37d string, __obf_a12d79aa3dd7e532 any, __obf_c73e1b232bb6aa12 bool) {
		__obf_a2a740eb043b5663 := __obf_a12d79aa3dd7e532 == nil || reflect.DeepEqual(__obf_a12d79aa3dd7e532, reflect.Zero(reflect.TypeOf(__obf_a12d79aa3dd7e532)).Interface())
		if __obf_c73e1b232bb6aa12 && __obf_a2a740eb043b5663 {
			return
		}
		__obf_05d6f2cd4d407fe8 = append(__obf_05d6f2cd4d407fe8, __obf_fb05cce5b14da37d)
		__obf_c2803eea3fa184fb[__obf_fb05cce5b14da37d] = __obf_a12d79aa3dd7e532
	}

	switch __obf_b6bf8b8bc11918d7 := __obf_53999f33a514460f.(type) {
	case map[string]any:
		for __obf_12c89da7507a8f69, __obf_be4148454b0decfa := range __obf_b6bf8b8bc11918d7 {
			__obf_44912f234fa2b844(__obf_12c89da7507a8f69, __obf_be4148454b0decfa, false)
		}

	default:
		__obf_82ab67b57d43a79b := reflect.TypeOf(__obf_53999f33a514460f)
		if __obf_82ab67b57d43a79b.Kind() == reflect.Pointer {
			__obf_82ab67b57d43a79b = __obf_82ab67b57d43a79b.Elem()
		}
		__obf_be4148454b0decfa := reflect.ValueOf(__obf_53999f33a514460f)
		if __obf_be4148454b0decfa.Kind() == reflect.Pointer {
			__obf_be4148454b0decfa = __obf_be4148454b0decfa.Elem()
		}

		if __obf_82ab67b57d43a79b.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_be83ea6d875ef30b := 0; __obf_be83ea6d875ef30b < __obf_82ab67b57d43a79b.NumField(); __obf_be83ea6d875ef30b++ {
			__obf_0f2119e3d4aa79b4 := __obf_82ab67b57d43a79b.Field(__obf_be83ea6d875ef30b).Tag.Get("db")
			if __obf_0f2119e3d4aa79b4 == "" || __obf_0f2119e3d4aa79b4 == "-" {
				continue
			}

			__obf_1ef5ba4adf4d1e3e := strings.Split(__obf_0f2119e3d4aa79b4, ",")
			__obf_44912f234fa2b844(__obf_1ef5ba4adf4d1e3e[0], __obf_be4148454b0decfa.Field(__obf_be83ea6d875ef30b).Interface(), len(__obf_1ef5ba4adf4d1e3e) > 1 && __obf_1ef5ba4adf4d1e3e[1] == "omitempty")
		}
	}

	if len(__obf_05d6f2cd4d407fe8) == 0 {
		return "", nil, errors.New("no fields to insert")
	}

	__obf_3d2a5afadabd5bfc := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)",
		__obf_6908758718a006fd,
		strings.Join(__obf_05d6f2cd4d407fe8, ","),
		strings.Join(__obf_05d6f2cd4d407fe8, ",:"),
	)

	return __obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb, nil
}

// pkField 主键字段
// lockField 乐观锁字段(默认为updated_at)
// constField 固定不变字段
func BuildUpdateSQL(
	__obf_53999f33a514460f any,
	__obf_6908758718a006fd, __obf_62adcca9754ea074 string,
	__obf_37dc7e92146f2125 []string,
	__obf_80db84fa4b7eedb5 ...string,
) (string, map[string]any, error) {
	if __obf_6908758718a006fd == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_e0ea3a85cfb5152f := make(map[string]struct{})
	for _, __obf_e4f32635d56b8af0 := range __obf_37dc7e92146f2125 {
		__obf_e0ea3a85cfb5152f[__obf_e4f32635d56b8af0] = struct{}{}
	}

	__obf_c2803eea3fa184fb := make(map[string]any)
	__obf_05d6f2cd4d407fe8 := []string{}
	__obf_339d916ec3e9b735 := []string{}

	__obf_44912f234fa2b844 := func(__obf_fb05cce5b14da37d string, __obf_a12d79aa3dd7e532 any, __obf_c73e1b232bb6aa12 bool) {
		__obf_a2a740eb043b5663 := __obf_a12d79aa3dd7e532 == nil || reflect.DeepEqual(__obf_a12d79aa3dd7e532, reflect.Zero(reflect.TypeOf(__obf_a12d79aa3dd7e532)).Interface())

		switch {
		case func() bool {
			_, __obf_f65665999cfd53f1 := __obf_e0ea3a85cfb5152f[__obf_fb05cce5b14da37d]
			return __obf_f65665999cfd53f1
		}():
			__obf_339d916ec3e9b735 = append(__obf_339d916ec3e9b735, fmt.Sprintf("%s=:%s", __obf_fb05cce5b14da37d, __obf_fb05cce5b14da37d))
			__obf_c2803eea3fa184fb[__obf_fb05cce5b14da37d] = __obf_a12d79aa3dd7e532

		case slices.Contains(__obf_80db84fa4b7eedb5, __obf_fb05cce5b14da37d):
			// skip constField
		case __obf_fb05cce5b14da37d == __obf_62adcca9754ea074:
			const __obf_fd67da5d99bab6f9 = "RESERVED_LOCK"
			__obf_05d6f2cd4d407fe8 = append(__obf_05d6f2cd4d407fe8, fmt.Sprintf("%s=:%s", __obf_fb05cce5b14da37d, __obf_fb05cce5b14da37d))
			__obf_339d916ec3e9b735 = append(__obf_339d916ec3e9b735, fmt.Sprintf("%s=:%s", __obf_fb05cce5b14da37d, __obf_fd67da5d99bab6f9))
			__obf_c2803eea3fa184fb[__obf_fb05cce5b14da37d] = time.Now().Unix()
			__obf_c2803eea3fa184fb[__obf_fd67da5d99bab6f9] = __obf_a12d79aa3dd7e532

		default:
			if __obf_c73e1b232bb6aa12 && __obf_a2a740eb043b5663 {
				return
			}
			__obf_05d6f2cd4d407fe8 = append(__obf_05d6f2cd4d407fe8, fmt.Sprintf("%s=:%s", __obf_fb05cce5b14da37d, __obf_fb05cce5b14da37d))
			__obf_c2803eea3fa184fb[__obf_fb05cce5b14da37d] = __obf_a12d79aa3dd7e532
		}
	}

	switch __obf_b6bf8b8bc11918d7 := __obf_53999f33a514460f.(type) {
	case map[string]any:
		for __obf_12c89da7507a8f69, __obf_be4148454b0decfa := range __obf_b6bf8b8bc11918d7 {
			__obf_44912f234fa2b844(__obf_12c89da7507a8f69, __obf_be4148454b0decfa, false)
		}

	default:

		__obf_82ab67b57d43a79b := reflect.TypeOf(__obf_53999f33a514460f)
		if __obf_82ab67b57d43a79b.Kind() == reflect.Pointer {
			__obf_82ab67b57d43a79b = __obf_82ab67b57d43a79b.Elem()
		}
		__obf_be4148454b0decfa := reflect.ValueOf(__obf_53999f33a514460f)
		if __obf_be4148454b0decfa.Kind() == reflect.Pointer {
			__obf_be4148454b0decfa = __obf_be4148454b0decfa.Elem()
		}
		if __obf_82ab67b57d43a79b.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_be83ea6d875ef30b := 0; __obf_be83ea6d875ef30b < __obf_82ab67b57d43a79b.NumField(); __obf_be83ea6d875ef30b++ {
			__obf_0f2119e3d4aa79b4 := __obf_82ab67b57d43a79b.Field(__obf_be83ea6d875ef30b).Tag.Get("db")
			if __obf_0f2119e3d4aa79b4 == "" || __obf_0f2119e3d4aa79b4 == "-" {
				continue
			}

			__obf_1ef5ba4adf4d1e3e := strings.Split(__obf_0f2119e3d4aa79b4, ",")
			__obf_44912f234fa2b844(__obf_1ef5ba4adf4d1e3e[0], __obf_be4148454b0decfa.Field(__obf_be83ea6d875ef30b).Interface(), len(__obf_1ef5ba4adf4d1e3e) > 1 && __obf_1ef5ba4adf4d1e3e[1] == "omitempty")
		}
	}

	if len(__obf_37dc7e92146f2125) == 0 || len(__obf_339d916ec3e9b735) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_05d6f2cd4d407fe8) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	// if lockField != "" && !hasLockField {
	// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
	// }

	__obf_3d2a5afadabd5bfc := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		__obf_6908758718a006fd,
		strings.Join(__obf_05d6f2cd4d407fe8, ","),
		strings.Join(__obf_339d916ec3e9b735, " AND "),
	)

	return __obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb, nil
}
