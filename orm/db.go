package __obf_df9e37b4dd16fa57

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

func NewORM(__obf_8beef8d164b02caa *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_8beef8d164b02caa.Driver, __obf_8beef8d164b02caa.DataSource)

	switch __obf_8beef8d164b02caa.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_8beef8d164b02caa.Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_8beef8d164b02caa.Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_8beef8d164b02caa.Placeholder = ":"
	case "sqlserver":
		__obf_8beef8d164b02caa.Placeholder = "@"
	}

	if __obf_4795d167371f16d4 := DB.Ping(); __obf_4795d167371f16d4 != nil {
		panic(__obf_4795d167371f16d4)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_8beef8d164b02caa,
		}, func() {
			DB.Close()
		}
}

func (__obf_ed3a9a8432b2d457 *ORM) ListMap(__obf_4f052f7555da44c1 *sqlx.Rows, __obf_938e189e8ea815b2 func(map[string]any) (string, string)) (__obf_7480d8040044dcbd []map[string]any, __obf_4795d167371f16d4 error) {
	for __obf_4f052f7555da44c1.Next() {
		__obf_bf519e886ecf0c6d := make(map[string]any)
		if __obf_4795d167371f16d4 = __obf_4f052f7555da44c1.MapScan(__obf_bf519e886ecf0c6d); __obf_4795d167371f16d4 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_4795d167371f16d4)
		}
		for __obf_6c0498eabb04b3fb, __obf_9573d75ff89816de := range __obf_bf519e886ecf0c6d {
			switch __obf_9573d75ff89816de := __obf_9573d75ff89816de.(type) {
			case []uint8:
				__obf_bf519e886ecf0c6d[__obf_6c0498eabb04b3fb] = string(__obf_9573d75ff89816de)
			case int64:
				__obf_bf519e886ecf0c6d[__obf_6c0498eabb04b3fb] = int64(__obf_9573d75ff89816de)
			case uint64:
				__obf_bf519e886ecf0c6d[__obf_6c0498eabb04b3fb] = uint64(__obf_9573d75ff89816de)
			}
		}
		if __obf_938e189e8ea815b2 != nil {
			__obf_ba0b180d5e50af65, __obf_8baa416e653862dd := __obf_938e189e8ea815b2(__obf_bf519e886ecf0c6d)
			__obf_bf519e886ecf0c6d[__obf_ba0b180d5e50af65] = __obf_8baa416e653862dd
		}
		__obf_7480d8040044dcbd = append(__obf_7480d8040044dcbd, __obf_bf519e886ecf0c6d)
	}
	return
}

func (__obf_ed3a9a8432b2d457 *ORM) Exists(__obf_c8be92edb09e8bee, __obf_511a4b0e43c6c0f2 string, __obf_3cef69abfd4b5fed ...any) bool {
	var __obf_16c35632de64bb7b bool
	_ = __obf_ed3a9a8432b2d457.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_c8be92edb09e8bee, __obf_511a4b0e43c6c0f2), __obf_3cef69abfd4b5fed...).Scan(&__obf_16c35632de64bb7b)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_16c35632de64bb7b
}

func (__obf_ed3a9a8432b2d457 *ORM) SaveModel(__obf_84f7492381c65d26 any, __obf_abd231560064ce3e string) (string, error) {
	__obf_9573d75ff89816de := reflect.ValueOf(__obf_84f7492381c65d26)
	if __obf_9573d75ff89816de.Kind() == reflect.Pointer {
		__obf_9573d75ff89816de = __obf_9573d75ff89816de.Elem()
	}
	var (
		__obf_6270f0d6af680a4f = __obf_9573d75ff89816de.FieldByName("Id")
		__obf_f996fba6555022f9 = util.ToSnake(reflect.Indirect(__obf_9573d75ff89816de).Type().Name())
		__obf_3b4f9deb3c30de53 string
		__obf_3cef69abfd4b5fed map[string]any
		__obf_4795d167371f16d4 error
	)

	if __obf_6270f0d6af680a4f.IsZero() {
		// insert
		__obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed, __obf_4795d167371f16d4 = BuildInsertSQL(__obf_84f7492381c65d26, __obf_f996fba6555022f9)
		if __obf_4795d167371f16d4 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_4795d167371f16d4.Error())
		}
		if __obf_abd231560064ce3e == "" {
			__obf_abd231560064ce3e = util.StringUUID()
		}

		__obf_3cef69abfd4b5fed["id"] = __obf_abd231560064ce3e
		__obf_48e721f3f91e157d := time.Now().Unix()
		if _, __obf_1b7dc4184c531bcd := __obf_3cef69abfd4b5fed[__obf_ed3a9a8432b2d457.Option.Reserved.CreateTime]; __obf_1b7dc4184c531bcd {
			__obf_3cef69abfd4b5fed[__obf_ed3a9a8432b2d457.Option.Reserved.CreateTime] = __obf_48e721f3f91e157d
		}
		if _, __obf_1b7dc4184c531bcd := __obf_3cef69abfd4b5fed[__obf_ed3a9a8432b2d457.Option.Reserved.UpdateTime]; __obf_1b7dc4184c531bcd {
			__obf_3cef69abfd4b5fed[__obf_ed3a9a8432b2d457.Option.Reserved.UpdateTime] = __obf_48e721f3f91e157d
		}
	} else {
		// update
		__obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed, __obf_4795d167371f16d4 = BuildUpdateSQL(__obf_84f7492381c65d26, __obf_f996fba6555022f9, __obf_ed3a9a8432b2d457.Option.Reserved.UpdateTime, []string{__obf_ed3a9a8432b2d457.Option.Reserved.ID}, __obf_ed3a9a8432b2d457.Option.Reserved.CreateTime, __obf_ed3a9a8432b2d457.Option.Reserved.NO)
		if __obf_4795d167371f16d4 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_4795d167371f16d4.Error())
		}
	}

	if __obf_3b4f9deb3c30de53 != "" {
		_, __obf_4795d167371f16d4 := __obf_ed3a9a8432b2d457.DB.NamedExec(__obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed)
		if __obf_4795d167371f16d4 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_4795d167371f16d4.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_3b4f9deb3c30de53, "args", __obf_3cef69abfd4b5fed)
	}

	return __obf_abd231560064ce3e, nil

}

func (__obf_ed3a9a8432b2d457 *ORM) SaveModelWidthAutoID(__obf_84f7492381c65d26 any, __obf_abd231560064ce3e int64) (int64, error) {
	// query, args := modelBind(db.Option, m, insertId)
	__obf_9573d75ff89816de := reflect.ValueOf(__obf_84f7492381c65d26)
	if __obf_9573d75ff89816de.Kind() == reflect.Pointer {
		__obf_9573d75ff89816de = __obf_9573d75ff89816de.Elem()
	}
	__obf_6270f0d6af680a4f := __obf_9573d75ff89816de.FieldByName("Id")
	__obf_f996fba6555022f9 := util.ToSnake(reflect.Indirect(__obf_9573d75ff89816de).Type().Name())
	var (
		__obf_3b4f9deb3c30de53 string
		__obf_3cef69abfd4b5fed map[string]any
		__obf_4795d167371f16d4 error
	)

	if __obf_6270f0d6af680a4f.IsZero() {
		// insert
		__obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed, __obf_4795d167371f16d4 = BuildInsertSQL(__obf_84f7492381c65d26, __obf_f996fba6555022f9)
		if __obf_4795d167371f16d4 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_4795d167371f16d4.Error())
		}
		if __obf_abd231560064ce3e != 0 {
			__obf_3cef69abfd4b5fed["id"] = __obf_abd231560064ce3e
		}
		__obf_48e721f3f91e157d := time.Now().Unix()
		if _, __obf_1b7dc4184c531bcd := __obf_3cef69abfd4b5fed[__obf_ed3a9a8432b2d457.Option.Reserved.CreateTime]; __obf_1b7dc4184c531bcd {
			__obf_3cef69abfd4b5fed[__obf_ed3a9a8432b2d457.Option.Reserved.CreateTime] = __obf_48e721f3f91e157d
		}
		if _, __obf_1b7dc4184c531bcd := __obf_3cef69abfd4b5fed[__obf_ed3a9a8432b2d457.Option.Reserved.UpdateTime]; __obf_1b7dc4184c531bcd {
			__obf_3cef69abfd4b5fed[__obf_ed3a9a8432b2d457.Option.Reserved.UpdateTime] = __obf_48e721f3f91e157d
		}
	} else {
		// update
		__obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed, __obf_4795d167371f16d4 = BuildUpdateSQL(__obf_84f7492381c65d26, __obf_f996fba6555022f9, __obf_ed3a9a8432b2d457.Option.Reserved.UpdateTime, []string{__obf_ed3a9a8432b2d457.Option.Reserved.ID}, __obf_ed3a9a8432b2d457.Option.Reserved.CreateTime, __obf_ed3a9a8432b2d457.Option.Reserved.NO)
		if __obf_4795d167371f16d4 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_4795d167371f16d4.Error())
		}
	}

	if __obf_3b4f9deb3c30de53 != "" {
		__obf_8baa416e653862dd, __obf_4795d167371f16d4 := __obf_ed3a9a8432b2d457.DB.NamedExec(__obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed)
		if __obf_4795d167371f16d4 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_4795d167371f16d4.Error())
		}
		if __obf_abd231560064ce3e == 0 {
			return __obf_8baa416e653862dd.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_3b4f9deb3c30de53, "args", __obf_3cef69abfd4b5fed)
	}

	return __obf_abd231560064ce3e, nil
}

func (__obf_ed3a9a8432b2d457 *ORM) SaveTxModel(__obf_4e1da10c2afe56d6 *sqlx.Tx, __obf_84f7492381c65d26 any, __obf_abd231560064ce3e string) error {
	var __obf_3b4f9deb3c30de53 string
	var __obf_3cef69abfd4b5fed map[string]any
	__obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed = __obf_a2d849acae211810(__obf_ed3a9a8432b2d457.Option, __obf_84f7492381c65d26, __obf_abd231560064ce3e)

	if __obf_3b4f9deb3c30de53 != "" {
		if _, __obf_4795d167371f16d4 := __obf_4e1da10c2afe56d6.NamedExec(__obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed); __obf_4795d167371f16d4 != nil {
			_ = __obf_4e1da10c2afe56d6.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_4795d167371f16d4)
		}

		slog.Debug("SaveTxModel", "query", __obf_3b4f9deb3c30de53, "args", __obf_3cef69abfd4b5fed)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_a2d849acae211810[RecordID int | string](__obf_8beef8d164b02caa *Options, __obf_84f7492381c65d26 any, __obf_abd231560064ce3e RecordID) (__obf_d76b5957a220f5b5 string, __obf_3cef69abfd4b5fed map[string]any) {
	__obf_9573d75ff89816de := reflect.ValueOf(__obf_84f7492381c65d26)
	if __obf_9573d75ff89816de.Kind() == reflect.Pointer {
		__obf_9573d75ff89816de = __obf_9573d75ff89816de.Elem()
	}
	var __obf_ee187ca87657359f []string
	__obf_3cef69abfd4b5fed = make(map[string]any)

	__obf_6270f0d6af680a4f := __obf_9573d75ff89816de.FieldByName("Id")
	if __obf_6270f0d6af680a4f.IsZero() {
		// 判断inserId是否为零值
		if __obf_abd231560064ce3e != *new(RecordID) {
			__obf_ee187ca87657359f = append(__obf_ee187ca87657359f, __obf_8beef8d164b02caa.Reserved.ID)
			__obf_3cef69abfd4b5fed[__obf_8beef8d164b02caa.Reserved.ID] = __obf_abd231560064ce3e
		} else {
			if __obf_6270f0d6af680a4f.Kind() == reflect.String {
				__obf_ee187ca87657359f = append(__obf_ee187ca87657359f, __obf_8beef8d164b02caa.Reserved.ID)
				__obf_3cef69abfd4b5fed[__obf_8beef8d164b02caa.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}

		__obf_11620555d3ba3b92 := reflect.Indirect(__obf_9573d75ff89816de).Type()
		for __obf_2a7388295d503598 := range __obf_11620555d3ba3b92.NumField() {
			if __obf_2e7920ea71b44a3b := __obf_11620555d3ba3b92.Field(__obf_2a7388295d503598); __obf_2e7920ea71b44a3b.Tag.Get("db") != "-" {
				switch __obf_7abde0e183c9f282 := util.ToSnake(__obf_2e7920ea71b44a3b.Name); __obf_7abde0e183c9f282 {
				case __obf_8beef8d164b02caa.Reserved.CreateTime, __obf_8beef8d164b02caa.Reserved.UpdateTime:
					__obf_ee187ca87657359f = append(__obf_ee187ca87657359f, __obf_7abde0e183c9f282)
					__obf_3cef69abfd4b5fed[__obf_7abde0e183c9f282] = time.Now().Unix()
				default:
					__obf_a0400a0a546202a1 := __obf_9573d75ff89816de.Field(__obf_2a7388295d503598).Interface()
					if strings.HasPrefix(__obf_7abde0e183c9f282, "is_") || __obf_a0400a0a546202a1 != reflect.Zero(__obf_2e7920ea71b44a3b.Type).Interface() {
						__obf_ee187ca87657359f = append(__obf_ee187ca87657359f, __obf_7abde0e183c9f282)
						__obf_3cef69abfd4b5fed[__obf_7abde0e183c9f282] = __obf_a0400a0a546202a1
					}
				}
			}
		}
		__obf_d76b5957a220f5b5 = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_11620555d3ba3b92.Name()), strings.Join(__obf_ee187ca87657359f, ","), strings.Join(__obf_ee187ca87657359f, ",:"))
	} else {
		__obf_11620555d3ba3b92 := reflect.Indirect(__obf_9573d75ff89816de).Type()
		for __obf_2a7388295d503598 := range __obf_11620555d3ba3b92.NumField() {
			if __obf_2e7920ea71b44a3b := __obf_11620555d3ba3b92.Field(__obf_2a7388295d503598); __obf_2e7920ea71b44a3b.Tag.Get("db") != "-" {
				__obf_a0400a0a546202a1 := __obf_9573d75ff89816de.Field(__obf_2a7388295d503598).Interface()
				if __obf_7abde0e183c9f282 := util.ToSnake(__obf_2e7920ea71b44a3b.Name); strings.HasPrefix(__obf_7abde0e183c9f282, "is_") || __obf_a0400a0a546202a1 != reflect.Zero(__obf_2e7920ea71b44a3b.Type).Interface() {
					__obf_ee187ca87657359f = append(__obf_ee187ca87657359f, fmt.Sprintf("%s=:%s", __obf_7abde0e183c9f282, __obf_7abde0e183c9f282))
					__obf_3cef69abfd4b5fed[__obf_7abde0e183c9f282] = __obf_a0400a0a546202a1
				}
			}
		}

		if util.HasField(__obf_9573d75ff89816de, util.ToCamel(__obf_8beef8d164b02caa.Reserved.UpdateTime)) {
			__obf_2187a7168f44f1dc := __obf_3cef69abfd4b5fed[__obf_8beef8d164b02caa.Reserved.UpdateTime]
			__obf_3cef69abfd4b5fed[__obf_8beef8d164b02caa.Reserved.UpdateTime] = time.Now().Unix()
			__obf_d76b5957a220f5b5 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_11620555d3ba3b92.Name()), strings.Join(__obf_ee187ca87657359f, ","), __obf_8beef8d164b02caa.Reserved.ID, __obf_8beef8d164b02caa.Reserved.ID, __obf_8beef8d164b02caa.Reserved.UpdateTime, __obf_2187a7168f44f1dc)
		} else {
			__obf_d76b5957a220f5b5 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_11620555d3ba3b92.Name()), strings.Join(__obf_ee187ca87657359f, ","), __obf_8beef8d164b02caa.Reserved.ID, __obf_8beef8d164b02caa.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_6611c297ec3770df any, __obf_f996fba6555022f9 string) (string, map[string]any, error) {
	if __obf_f996fba6555022f9 == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_3cef69abfd4b5fed := make(map[string]any)
	__obf_ee187ca87657359f := []string{}

	__obf_58d2d2aec8430e57 := func(__obf_a5ff5bba00286eca string, __obf_fe4dab033b4bbf22 any, __obf_e539bff9ccf80b81 bool) {
		__obf_079fbbfa5dfd6fd3 := __obf_fe4dab033b4bbf22 == nil || reflect.DeepEqual(__obf_fe4dab033b4bbf22, reflect.Zero(reflect.TypeOf(__obf_fe4dab033b4bbf22)).Interface())
		if __obf_e539bff9ccf80b81 && __obf_079fbbfa5dfd6fd3 {
			return
		}
		__obf_ee187ca87657359f = append(__obf_ee187ca87657359f, __obf_a5ff5bba00286eca)
		__obf_3cef69abfd4b5fed[__obf_a5ff5bba00286eca] = __obf_fe4dab033b4bbf22
	}

	switch __obf_a0400a0a546202a1 := __obf_6611c297ec3770df.(type) {
	case map[string]any:
		for __obf_6c0498eabb04b3fb, __obf_9573d75ff89816de := range __obf_a0400a0a546202a1 {
			__obf_58d2d2aec8430e57(__obf_6c0498eabb04b3fb, __obf_9573d75ff89816de, false)
		}

	default:
		__obf_fb7e0dccfd572f81 := reflect.TypeOf(__obf_6611c297ec3770df)
		if __obf_fb7e0dccfd572f81.Kind() == reflect.Pointer {
			__obf_fb7e0dccfd572f81 = __obf_fb7e0dccfd572f81.Elem()
		}
		__obf_9573d75ff89816de := reflect.ValueOf(__obf_6611c297ec3770df)
		if __obf_9573d75ff89816de.Kind() == reflect.Pointer {
			__obf_9573d75ff89816de = __obf_9573d75ff89816de.Elem()
		}

		if __obf_fb7e0dccfd572f81.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_2a7388295d503598 := 0; __obf_2a7388295d503598 < __obf_fb7e0dccfd572f81.NumField(); __obf_2a7388295d503598++ {
			__obf_3baee9d6ee92faf1 := __obf_fb7e0dccfd572f81.Field(__obf_2a7388295d503598).Tag.Get("db")
			if __obf_3baee9d6ee92faf1 == "" || __obf_3baee9d6ee92faf1 == "-" {
				continue
			}

			__obf_09eb0ef71367f7be := strings.Split(__obf_3baee9d6ee92faf1, ",")
			__obf_58d2d2aec8430e57(__obf_09eb0ef71367f7be[0], __obf_9573d75ff89816de.Field(__obf_2a7388295d503598).Interface(), len(__obf_09eb0ef71367f7be) > 1 && __obf_09eb0ef71367f7be[1] == "omitempty")
		}
	}

	if len(__obf_ee187ca87657359f) == 0 {
		return "", nil, errors.New("no fields to insert")
	}

	__obf_3b4f9deb3c30de53 := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)",
		__obf_f996fba6555022f9,
		strings.Join(__obf_ee187ca87657359f, ","),
		strings.Join(__obf_ee187ca87657359f, ",:"),
	)

	return __obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed, nil
}

// pkField 主键字段
// lockField 乐观锁字段(默认为updated_at)
// constField 固定不变字段
func BuildUpdateSQL(
	__obf_6611c297ec3770df any,
	__obf_f996fba6555022f9, __obf_264f6cb271285128 string,
	__obf_d8385d8ced3e92df []string,
	__obf_278ac72f75b28e87 ...string,
) (string, map[string]any, error) {
	if __obf_f996fba6555022f9 == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_0b2057eea2744cbc := make(map[string]struct{})
	for _, __obf_506233e4e9256668 := range __obf_d8385d8ced3e92df {
		__obf_0b2057eea2744cbc[__obf_506233e4e9256668] = struct{}{}
	}

	__obf_3cef69abfd4b5fed := make(map[string]any)
	__obf_ee187ca87657359f := []string{}
	__obf_531d42817e463ca1 := []string{}

	__obf_58d2d2aec8430e57 := func(__obf_a5ff5bba00286eca string, __obf_fe4dab033b4bbf22 any, __obf_e539bff9ccf80b81 bool) {
		__obf_079fbbfa5dfd6fd3 := __obf_fe4dab033b4bbf22 == nil || reflect.DeepEqual(__obf_fe4dab033b4bbf22, reflect.Zero(reflect.TypeOf(__obf_fe4dab033b4bbf22)).Interface())

		switch {
		case func() bool {
			_, __obf_1b7dc4184c531bcd := __obf_0b2057eea2744cbc[__obf_a5ff5bba00286eca]
			return __obf_1b7dc4184c531bcd
		}():
			__obf_531d42817e463ca1 = append(__obf_531d42817e463ca1, fmt.Sprintf("%s=:%s", __obf_a5ff5bba00286eca, __obf_a5ff5bba00286eca))
			__obf_3cef69abfd4b5fed[__obf_a5ff5bba00286eca] = __obf_fe4dab033b4bbf22

		case slices.Contains(__obf_278ac72f75b28e87, __obf_a5ff5bba00286eca):
			// skip constField
		case __obf_a5ff5bba00286eca == __obf_264f6cb271285128:
			const __obf_b52248a69e662e62 = "RESERVED_LOCK"
			__obf_ee187ca87657359f = append(__obf_ee187ca87657359f, fmt.Sprintf("%s=:%s", __obf_a5ff5bba00286eca, __obf_a5ff5bba00286eca))
			__obf_531d42817e463ca1 = append(__obf_531d42817e463ca1, fmt.Sprintf("%s=:%s", __obf_a5ff5bba00286eca, __obf_b52248a69e662e62))
			__obf_3cef69abfd4b5fed[__obf_a5ff5bba00286eca] = time.Now().Unix()
			__obf_3cef69abfd4b5fed[__obf_b52248a69e662e62] = __obf_fe4dab033b4bbf22

		default:
			if __obf_e539bff9ccf80b81 && __obf_079fbbfa5dfd6fd3 {
				return
			}
			__obf_ee187ca87657359f = append(__obf_ee187ca87657359f, fmt.Sprintf("%s=:%s", __obf_a5ff5bba00286eca, __obf_a5ff5bba00286eca))
			__obf_3cef69abfd4b5fed[__obf_a5ff5bba00286eca] = __obf_fe4dab033b4bbf22
		}
	}

	switch __obf_a0400a0a546202a1 := __obf_6611c297ec3770df.(type) {
	case map[string]any:
		for __obf_6c0498eabb04b3fb, __obf_9573d75ff89816de := range __obf_a0400a0a546202a1 {
			__obf_58d2d2aec8430e57(__obf_6c0498eabb04b3fb, __obf_9573d75ff89816de, false)
		}

	default:

		__obf_fb7e0dccfd572f81 := reflect.TypeOf(__obf_6611c297ec3770df)
		if __obf_fb7e0dccfd572f81.Kind() == reflect.Pointer {
			__obf_fb7e0dccfd572f81 = __obf_fb7e0dccfd572f81.Elem()
		}
		__obf_9573d75ff89816de := reflect.ValueOf(__obf_6611c297ec3770df)
		if __obf_9573d75ff89816de.Kind() == reflect.Pointer {
			__obf_9573d75ff89816de = __obf_9573d75ff89816de.Elem()
		}
		if __obf_fb7e0dccfd572f81.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_2a7388295d503598 := 0; __obf_2a7388295d503598 < __obf_fb7e0dccfd572f81.NumField(); __obf_2a7388295d503598++ {
			__obf_3baee9d6ee92faf1 := __obf_fb7e0dccfd572f81.Field(__obf_2a7388295d503598).Tag.Get("db")
			if __obf_3baee9d6ee92faf1 == "" || __obf_3baee9d6ee92faf1 == "-" {
				continue
			}

			__obf_09eb0ef71367f7be := strings.Split(__obf_3baee9d6ee92faf1, ",")
			__obf_58d2d2aec8430e57(__obf_09eb0ef71367f7be[0], __obf_9573d75ff89816de.Field(__obf_2a7388295d503598).Interface(), len(__obf_09eb0ef71367f7be) > 1 && __obf_09eb0ef71367f7be[1] == "omitempty")
		}
	}

	if len(__obf_d8385d8ced3e92df) == 0 || len(__obf_531d42817e463ca1) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_ee187ca87657359f) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	// if lockField != "" && !hasLockField {
	// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
	// }

	__obf_3b4f9deb3c30de53 := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		__obf_f996fba6555022f9,
		strings.Join(__obf_ee187ca87657359f, ","),
		strings.Join(__obf_531d42817e463ca1, " AND "),
	)

	return __obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed, nil
}
