package __obf_ed189c965cdcd132

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

func NewORM(__obf_986080d7afca28dc *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_986080d7afca28dc.Driver, __obf_986080d7afca28dc.DataSource)

	switch __obf_986080d7afca28dc.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_986080d7afca28dc.
			Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_986080d7afca28dc.
			Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_986080d7afca28dc.
			Placeholder = ":"
	case "sqlserver":
		__obf_986080d7afca28dc.
			Placeholder = "@"
	}

	if __obf_9f5a64d352151704 := DB.Ping(); __obf_9f5a64d352151704 != nil {
		panic(__obf_9f5a64d352151704)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_986080d7afca28dc,
		}, func() {
			DB.Close()
		}
}

func (__obf_e17d8f0e9734b0de *ORM) ListMap(__obf_b9da22b0e5936af1 *sqlx.Rows, __obf_394ce5329b6ab76e func(map[string]any) (string, string)) (__obf_3bf64be4135c172e []map[string]any, __obf_9f5a64d352151704 error) {
	for __obf_b9da22b0e5936af1.Next() {
		__obf_d027ed907492e50e := make(map[string]any)
		if __obf_9f5a64d352151704 = __obf_b9da22b0e5936af1.MapScan(__obf_d027ed907492e50e); __obf_9f5a64d352151704 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_9f5a64d352151704)
		}
		for __obf_cb6df4213d1d222d, __obf_604eba2487c8ea9b := range __obf_d027ed907492e50e {
			switch __obf_604eba2487c8ea9b := __obf_604eba2487c8ea9b.(type) {
			case []uint8:
				__obf_d027ed907492e50e[__obf_cb6df4213d1d222d] = string(__obf_604eba2487c8ea9b)
			case int64:
				__obf_d027ed907492e50e[__obf_cb6df4213d1d222d] = int64(__obf_604eba2487c8ea9b)
			case uint64:
				__obf_d027ed907492e50e[__obf_cb6df4213d1d222d] = uint64(__obf_604eba2487c8ea9b)
			}
		}
		if __obf_394ce5329b6ab76e != nil {
			__obf_7eed0b8644379c62, __obf_e1e03f4986eb2a8d := __obf_394ce5329b6ab76e(__obf_d027ed907492e50e)
			__obf_d027ed907492e50e[__obf_7eed0b8644379c62] = __obf_e1e03f4986eb2a8d
		}
		__obf_3bf64be4135c172e = append(__obf_3bf64be4135c172e, __obf_d027ed907492e50e)
	}
	return
}

func (__obf_e17d8f0e9734b0de *ORM) Exists(__obf_284385e29e7da903, __obf_055a197727e3c792 string, __obf_8670d196c9bebfda ...any) bool {
	var __obf_7a7bc1fc31993e57 bool
	_ = __obf_e17d8f0e9734b0de.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_284385e29e7da903, __obf_055a197727e3c792), __obf_8670d196c9bebfda...).Scan(&__obf_7a7bc1fc31993e57)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_7a7bc1fc31993e57
}

func (__obf_e17d8f0e9734b0de *ORM) SaveModel(__obf_fa84bbdca73f8f1c any, __obf_36129a25ed1ba406 string) (string, error) {
	__obf_604eba2487c8ea9b := reflect.ValueOf(__obf_fa84bbdca73f8f1c)
	if __obf_604eba2487c8ea9b.Kind() == reflect.Pointer {
		__obf_604eba2487c8ea9b = __obf_604eba2487c8ea9b.Elem()
	}
	var (
		__obf_b8adb094b68d7e2d = __obf_604eba2487c8ea9b.
					FieldByName("Id")
		__obf_4dd2ad913c3df0d6 = util.ToSnake(reflect.Indirect(__obf_604eba2487c8ea9b).Type().Name())
		__obf_91b87182c3b48e5a string
		__obf_8670d196c9bebfda map[string]any
		__obf_9f5a64d352151704 error
	)

	if __obf_b8adb094b68d7e2d.IsZero() {
		__obf_91b87182c3b48e5a,
			// insert
			__obf_8670d196c9bebfda, __obf_9f5a64d352151704 = BuildInsertSQL(__obf_fa84bbdca73f8f1c, __obf_4dd2ad913c3df0d6)
		if __obf_9f5a64d352151704 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_9f5a64d352151704.Error())
		}
		if __obf_36129a25ed1ba406 == "" {
			__obf_36129a25ed1ba406 = util.StringUUID()
		}
		__obf_8670d196c9bebfda["id"] = __obf_36129a25ed1ba406
		__obf_f1f28cbc561bd323 := time.Now().Unix()
		if _, __obf_27df3b3b715e7e91 := __obf_8670d196c9bebfda[__obf_e17d8f0e9734b0de.Option.Reserved.CreateTime]; __obf_27df3b3b715e7e91 {
			__obf_8670d196c9bebfda[__obf_e17d8f0e9734b0de.Option.Reserved.CreateTime] = __obf_f1f28cbc561bd323
		}
		if _, __obf_27df3b3b715e7e91 := __obf_8670d196c9bebfda[__obf_e17d8f0e9734b0de.Option.Reserved.UpdateTime]; __obf_27df3b3b715e7e91 {
			__obf_8670d196c9bebfda[__obf_e17d8f0e9734b0de.Option.Reserved.UpdateTime] = __obf_f1f28cbc561bd323
		}
	} else {
		__obf_91b87182c3b48e5a,
			// update
			__obf_8670d196c9bebfda, __obf_9f5a64d352151704 = BuildUpdateSQL(__obf_fa84bbdca73f8f1c, __obf_4dd2ad913c3df0d6, __obf_e17d8f0e9734b0de.Option.Reserved.UpdateTime, []string{__obf_e17d8f0e9734b0de.Option.Reserved.ID}, __obf_e17d8f0e9734b0de.Option.Reserved.CreateTime, __obf_e17d8f0e9734b0de.Option.Reserved.NO)
		if __obf_9f5a64d352151704 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_9f5a64d352151704.Error())
		}
	}

	if __obf_91b87182c3b48e5a != "" {
		_, __obf_9f5a64d352151704 := __obf_e17d8f0e9734b0de.DB.NamedExec(__obf_91b87182c3b48e5a, __obf_8670d196c9bebfda)
		if __obf_9f5a64d352151704 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_9f5a64d352151704.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_91b87182c3b48e5a, "args", __obf_8670d196c9bebfda)
	}

	return __obf_36129a25ed1ba406, nil

}

func (__obf_e17d8f0e9734b0de *ORM) SaveModelWidthAutoID(__obf_fa84bbdca73f8f1c any, __obf_36129a25ed1ba406 int64) (int64, error) {
	__obf_604eba2487c8ea9b := // query, args := modelBind(db.Option, m, insertId)
		reflect.ValueOf(__obf_fa84bbdca73f8f1c)
	if __obf_604eba2487c8ea9b.Kind() == reflect.Pointer {
		__obf_604eba2487c8ea9b = __obf_604eba2487c8ea9b.Elem()
	}
	__obf_b8adb094b68d7e2d := __obf_604eba2487c8ea9b.FieldByName("Id")
	__obf_4dd2ad913c3df0d6 := util.ToSnake(reflect.Indirect(__obf_604eba2487c8ea9b).Type().Name())
	var (
		__obf_91b87182c3b48e5a string
		__obf_8670d196c9bebfda map[string]any
		__obf_9f5a64d352151704 error
	)

	if __obf_b8adb094b68d7e2d.IsZero() {
		__obf_91b87182c3b48e5a,
			// insert
			__obf_8670d196c9bebfda, __obf_9f5a64d352151704 = BuildInsertSQL(__obf_fa84bbdca73f8f1c, __obf_4dd2ad913c3df0d6)
		if __obf_9f5a64d352151704 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_9f5a64d352151704.Error())
		}
		if __obf_36129a25ed1ba406 != 0 {
			__obf_8670d196c9bebfda["id"] = __obf_36129a25ed1ba406
		}
		__obf_f1f28cbc561bd323 := time.Now().Unix()
		if _, __obf_27df3b3b715e7e91 := __obf_8670d196c9bebfda[__obf_e17d8f0e9734b0de.Option.Reserved.CreateTime]; __obf_27df3b3b715e7e91 {
			__obf_8670d196c9bebfda[__obf_e17d8f0e9734b0de.Option.Reserved.CreateTime] = __obf_f1f28cbc561bd323
		}
		if _, __obf_27df3b3b715e7e91 := __obf_8670d196c9bebfda[__obf_e17d8f0e9734b0de.Option.Reserved.UpdateTime]; __obf_27df3b3b715e7e91 {
			__obf_8670d196c9bebfda[__obf_e17d8f0e9734b0de.Option.Reserved.UpdateTime] = __obf_f1f28cbc561bd323
		}
	} else {
		__obf_91b87182c3b48e5a,
			// update
			__obf_8670d196c9bebfda, __obf_9f5a64d352151704 = BuildUpdateSQL(__obf_fa84bbdca73f8f1c, __obf_4dd2ad913c3df0d6, __obf_e17d8f0e9734b0de.Option.Reserved.UpdateTime, []string{__obf_e17d8f0e9734b0de.Option.Reserved.ID}, __obf_e17d8f0e9734b0de.Option.Reserved.CreateTime, __obf_e17d8f0e9734b0de.Option.Reserved.NO)
		if __obf_9f5a64d352151704 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_9f5a64d352151704.Error())
		}
	}

	if __obf_91b87182c3b48e5a != "" {
		__obf_e1e03f4986eb2a8d, __obf_9f5a64d352151704 := __obf_e17d8f0e9734b0de.DB.NamedExec(__obf_91b87182c3b48e5a, __obf_8670d196c9bebfda)
		if __obf_9f5a64d352151704 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_9f5a64d352151704.Error())
		}
		if __obf_36129a25ed1ba406 == 0 {
			return __obf_e1e03f4986eb2a8d.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_91b87182c3b48e5a, "args", __obf_8670d196c9bebfda)
	}

	return __obf_36129a25ed1ba406, nil
}

func (__obf_e17d8f0e9734b0de *ORM) SaveTxModel(__obf_83eebeb3bc1b8de9 *sqlx.Tx, __obf_fa84bbdca73f8f1c any, __obf_36129a25ed1ba406 string) error {
	var __obf_91b87182c3b48e5a string
	var __obf_8670d196c9bebfda map[string]any
	__obf_91b87182c3b48e5a, __obf_8670d196c9bebfda = __obf_35be12b03b21ea30(__obf_e17d8f0e9734b0de.Option, __obf_fa84bbdca73f8f1c, __obf_36129a25ed1ba406)

	if __obf_91b87182c3b48e5a != "" {
		if _, __obf_9f5a64d352151704 := __obf_83eebeb3bc1b8de9.NamedExec(__obf_91b87182c3b48e5a, __obf_8670d196c9bebfda); __obf_9f5a64d352151704 != nil {
			_ = __obf_83eebeb3bc1b8de9.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_9f5a64d352151704)
		}

		slog.Debug("SaveTxModel", "query", __obf_91b87182c3b48e5a, "args", __obf_8670d196c9bebfda)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_35be12b03b21ea30[RecordID int | string](__obf_986080d7afca28dc *Options, __obf_fa84bbdca73f8f1c any, __obf_36129a25ed1ba406 RecordID) (__obf_b911737164f414fd string, __obf_8670d196c9bebfda map[string]any) {
	__obf_604eba2487c8ea9b := reflect.ValueOf(__obf_fa84bbdca73f8f1c)
	if __obf_604eba2487c8ea9b.Kind() == reflect.Pointer {
		__obf_604eba2487c8ea9b = __obf_604eba2487c8ea9b.Elem()
	}
	var __obf_9a8ea8a59523879e []string
	__obf_8670d196c9bebfda = make(map[string]any)
	__obf_b8adb094b68d7e2d := __obf_604eba2487c8ea9b.FieldByName("Id")
	if __obf_b8adb094b68d7e2d.IsZero() {
		// 判断inserId是否为零值
		if __obf_36129a25ed1ba406 != *new(RecordID) {
			__obf_9a8ea8a59523879e = append(__obf_9a8ea8a59523879e, __obf_986080d7afca28dc.Reserved.ID)
			__obf_8670d196c9bebfda[__obf_986080d7afca28dc.Reserved.ID] = __obf_36129a25ed1ba406
		} else {
			if __obf_b8adb094b68d7e2d.Kind() == reflect.String {
				__obf_9a8ea8a59523879e = append(__obf_9a8ea8a59523879e, __obf_986080d7afca28dc.Reserved.ID)
				__obf_8670d196c9bebfda[__obf_986080d7afca28dc.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}
		__obf_123438a31cb29c60 := reflect.Indirect(__obf_604eba2487c8ea9b).Type()
		for __obf_c7e5bacf18525a19 := range __obf_123438a31cb29c60.NumField() {
			if __obf_626ed35c11f71017 := __obf_123438a31cb29c60.Field(__obf_c7e5bacf18525a19); __obf_626ed35c11f71017.Tag.Get("db") != "-" {
				switch __obf_6063d381d441b5b3 := util.ToSnake(__obf_626ed35c11f71017.Name); __obf_6063d381d441b5b3 {
				case __obf_986080d7afca28dc.Reserved.CreateTime, __obf_986080d7afca28dc.Reserved.UpdateTime:
					__obf_9a8ea8a59523879e = append(__obf_9a8ea8a59523879e, __obf_6063d381d441b5b3)
					__obf_8670d196c9bebfda[__obf_6063d381d441b5b3] = time.Now().Unix()
				default:
					__obf_f9cbd174629de499 := __obf_604eba2487c8ea9b.Field(__obf_c7e5bacf18525a19).Interface()
					if strings.HasPrefix(__obf_6063d381d441b5b3, "is_") || __obf_f9cbd174629de499 != reflect.Zero(__obf_626ed35c11f71017.Type).Interface() {
						__obf_9a8ea8a59523879e = append(__obf_9a8ea8a59523879e, __obf_6063d381d441b5b3)
						__obf_8670d196c9bebfda[__obf_6063d381d441b5b3] = __obf_f9cbd174629de499
					}
				}
			}
		}
		__obf_b911737164f414fd = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_123438a31cb29c60.Name()), strings.Join(__obf_9a8ea8a59523879e, ","), strings.Join(__obf_9a8ea8a59523879e, ",:"))
	} else {
		__obf_123438a31cb29c60 := reflect.Indirect(__obf_604eba2487c8ea9b).Type()
		for __obf_c7e5bacf18525a19 := range __obf_123438a31cb29c60.NumField() {
			if __obf_626ed35c11f71017 := __obf_123438a31cb29c60.Field(__obf_c7e5bacf18525a19); __obf_626ed35c11f71017.Tag.Get("db") != "-" {
				__obf_f9cbd174629de499 := __obf_604eba2487c8ea9b.Field(__obf_c7e5bacf18525a19).Interface()
				if __obf_6063d381d441b5b3 := util.ToSnake(__obf_626ed35c11f71017.Name); strings.HasPrefix(__obf_6063d381d441b5b3, "is_") || __obf_f9cbd174629de499 != reflect.Zero(__obf_626ed35c11f71017.Type).Interface() {
					__obf_9a8ea8a59523879e = append(__obf_9a8ea8a59523879e, fmt.Sprintf("%s=:%s", __obf_6063d381d441b5b3, __obf_6063d381d441b5b3))
					__obf_8670d196c9bebfda[__obf_6063d381d441b5b3] = __obf_f9cbd174629de499
				}
			}
		}

		if util.HasField(__obf_604eba2487c8ea9b, util.ToCamel(__obf_986080d7afca28dc.Reserved.UpdateTime)) {
			__obf_5fb949d87f720239 := __obf_8670d196c9bebfda[__obf_986080d7afca28dc.Reserved.UpdateTime]
			__obf_8670d196c9bebfda[__obf_986080d7afca28dc.Reserved.UpdateTime] = time.Now().Unix()
			__obf_b911737164f414fd = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_123438a31cb29c60.Name()), strings.Join(__obf_9a8ea8a59523879e, ","), __obf_986080d7afca28dc.Reserved.ID, __obf_986080d7afca28dc.Reserved.ID, __obf_986080d7afca28dc.Reserved.UpdateTime, __obf_5fb949d87f720239)
		} else {
			__obf_b911737164f414fd = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_123438a31cb29c60.Name()), strings.Join(__obf_9a8ea8a59523879e, ","), __obf_986080d7afca28dc.Reserved.ID, __obf_986080d7afca28dc.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_3718213bd04eb8f2 any, __obf_4dd2ad913c3df0d6 string) (string, map[string]any, error) {
	if __obf_4dd2ad913c3df0d6 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_8670d196c9bebfda := make(map[string]any)
	__obf_9a8ea8a59523879e := []string{}
	__obf_aa34f0e7c2df88a3 := func(__obf_8e630cdb3c3b799f string, __obf_9c99160fd6d9fa7d any, __obf_8507ac958c0bc6cf bool) {
		__obf_a3cc5b40742c1900 := __obf_9c99160fd6d9fa7d == nil || reflect.DeepEqual(__obf_9c99160fd6d9fa7d, reflect.Zero(reflect.TypeOf(__obf_9c99160fd6d9fa7d)).Interface())
		if __obf_8507ac958c0bc6cf && __obf_a3cc5b40742c1900 {
			return
		}
		__obf_9a8ea8a59523879e = append(__obf_9a8ea8a59523879e, __obf_8e630cdb3c3b799f)
		__obf_8670d196c9bebfda[__obf_8e630cdb3c3b799f] = __obf_9c99160fd6d9fa7d
	}

	switch __obf_f9cbd174629de499 := __obf_3718213bd04eb8f2.(type) {
	case map[string]any:
		for __obf_cb6df4213d1d222d, __obf_604eba2487c8ea9b := range __obf_f9cbd174629de499 {
			__obf_aa34f0e7c2df88a3(__obf_cb6df4213d1d222d, __obf_604eba2487c8ea9b, false)
		}

	default:
		__obf_34511df5079d8a2e := reflect.TypeOf(__obf_3718213bd04eb8f2)
		if __obf_34511df5079d8a2e.Kind() == reflect.Pointer {
			__obf_34511df5079d8a2e = __obf_34511df5079d8a2e.Elem()
		}
		__obf_604eba2487c8ea9b := reflect.ValueOf(__obf_3718213bd04eb8f2)
		if __obf_604eba2487c8ea9b.Kind() == reflect.Pointer {
			__obf_604eba2487c8ea9b = __obf_604eba2487c8ea9b.Elem()
		}

		if __obf_34511df5079d8a2e.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_c7e5bacf18525a19 := 0; __obf_c7e5bacf18525a19 < __obf_34511df5079d8a2e.NumField(); __obf_c7e5bacf18525a19++ {
			__obf_a8c398c4c97f3c1b := __obf_34511df5079d8a2e.Field(__obf_c7e5bacf18525a19).Tag.Get("db")
			if __obf_a8c398c4c97f3c1b == "" || __obf_a8c398c4c97f3c1b == "-" {
				continue
			}
			__obf_6b25fa559885e4f1 := strings.Split(__obf_a8c398c4c97f3c1b, ",")
			__obf_aa34f0e7c2df88a3(__obf_6b25fa559885e4f1[0], __obf_604eba2487c8ea9b.Field(__obf_c7e5bacf18525a19).Interface(), len(__obf_6b25fa559885e4f1) > 1 && __obf_6b25fa559885e4f1[1] == "omitempty")
		}
	}

	if len(__obf_9a8ea8a59523879e) == 0 {
		return "", nil, errors.New("no fields to insert")
	}
	__obf_91b87182c3b48e5a := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", __obf_4dd2ad913c3df0d6, strings.Join(__obf_9a8ea8a59523879e, ","),
		strings.Join(__obf_9a8ea8a59523879e, ",:"),
	)

	return __obf_91b87182c3b48e5a,

		// pkField 主键字段
		// lockField 乐观锁字段(默认为updated_at)
		// constField 固定不变字段
		__obf_8670d196c9bebfda, nil
}

func BuildUpdateSQL(__obf_3718213bd04eb8f2 any, __obf_4dd2ad913c3df0d6, __obf_7bdcb455779c9cf1 string, __obf_b421a7979724ecd7 []string, __obf_74885e9d452f9ed3 ...string,
) (string, map[string]any, error) {
	if __obf_4dd2ad913c3df0d6 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_6352beb6380bc1ca := make(map[string]struct{})
	for _, __obf_a2b004a27590154a := range __obf_b421a7979724ecd7 {
		__obf_6352beb6380bc1ca[__obf_a2b004a27590154a] = struct{}{}
	}
	__obf_8670d196c9bebfda := make(map[string]any)
	__obf_9a8ea8a59523879e := []string{}
	__obf_4d6a41b54ff34740 := []string{}
	__obf_aa34f0e7c2df88a3 := func(__obf_8e630cdb3c3b799f string, __obf_9c99160fd6d9fa7d any, __obf_8507ac958c0bc6cf bool) {
		__obf_a3cc5b40742c1900 := __obf_9c99160fd6d9fa7d == nil || reflect.DeepEqual(__obf_9c99160fd6d9fa7d, reflect.Zero(reflect.TypeOf(__obf_9c99160fd6d9fa7d)).Interface())

		switch {
		case func() bool {
			_, __obf_27df3b3b715e7e91 := __obf_6352beb6380bc1ca[__obf_8e630cdb3c3b799f]
			return __obf_27df3b3b715e7e91
		}():
			__obf_4d6a41b54ff34740 = append(__obf_4d6a41b54ff34740, fmt.Sprintf("%s=:%s", __obf_8e630cdb3c3b799f, __obf_8e630cdb3c3b799f))
			__obf_8670d196c9bebfda[__obf_8e630cdb3c3b799f] = __obf_9c99160fd6d9fa7d

		case slices.Contains(__obf_74885e9d452f9ed3, __obf_8e630cdb3c3b799f):
			// skip constField
		case __obf_8e630cdb3c3b799f == __obf_7bdcb455779c9cf1:
			const __obf_67797b4886954b8e = "RESERVED_LOCK"
			__obf_9a8ea8a59523879e = append(__obf_9a8ea8a59523879e, fmt.Sprintf("%s=:%s", __obf_8e630cdb3c3b799f, __obf_8e630cdb3c3b799f))
			__obf_4d6a41b54ff34740 = append(__obf_4d6a41b54ff34740, fmt.Sprintf("%s=:%s", __obf_8e630cdb3c3b799f, __obf_67797b4886954b8e))
			__obf_8670d196c9bebfda[__obf_8e630cdb3c3b799f] = time.Now().Unix()
			__obf_8670d196c9bebfda[__obf_67797b4886954b8e] = __obf_9c99160fd6d9fa7d

		default:
			if __obf_8507ac958c0bc6cf && __obf_a3cc5b40742c1900 {
				return
			}
			__obf_9a8ea8a59523879e = append(__obf_9a8ea8a59523879e, fmt.Sprintf("%s=:%s", __obf_8e630cdb3c3b799f, __obf_8e630cdb3c3b799f))
			__obf_8670d196c9bebfda[__obf_8e630cdb3c3b799f] = __obf_9c99160fd6d9fa7d
		}
	}

	switch __obf_f9cbd174629de499 := __obf_3718213bd04eb8f2.(type) {
	case map[string]any:
		for __obf_cb6df4213d1d222d, __obf_604eba2487c8ea9b := range __obf_f9cbd174629de499 {
			__obf_aa34f0e7c2df88a3(__obf_cb6df4213d1d222d, __obf_604eba2487c8ea9b, false)
		}

	default:
		__obf_34511df5079d8a2e := reflect.TypeOf(__obf_3718213bd04eb8f2)
		if __obf_34511df5079d8a2e.Kind() == reflect.Pointer {
			__obf_34511df5079d8a2e = __obf_34511df5079d8a2e.Elem()
		}
		__obf_604eba2487c8ea9b := reflect.ValueOf(__obf_3718213bd04eb8f2)
		if __obf_604eba2487c8ea9b.Kind() == reflect.Pointer {
			__obf_604eba2487c8ea9b = __obf_604eba2487c8ea9b.Elem()
		}
		if __obf_34511df5079d8a2e.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_c7e5bacf18525a19 := 0; __obf_c7e5bacf18525a19 < __obf_34511df5079d8a2e.NumField(); __obf_c7e5bacf18525a19++ {
			__obf_a8c398c4c97f3c1b := __obf_34511df5079d8a2e.Field(__obf_c7e5bacf18525a19).Tag.Get("db")
			if __obf_a8c398c4c97f3c1b == "" || __obf_a8c398c4c97f3c1b == "-" {
				continue
			}
			__obf_6b25fa559885e4f1 := strings.Split(__obf_a8c398c4c97f3c1b, ",")
			__obf_aa34f0e7c2df88a3(__obf_6b25fa559885e4f1[0], __obf_604eba2487c8ea9b.Field(__obf_c7e5bacf18525a19).Interface(), len(__obf_6b25fa559885e4f1) > 1 && __obf_6b25fa559885e4f1[1] == "omitempty")
		}
	}

	if len(__obf_b421a7979724ecd7) == 0 || len(__obf_4d6a41b54ff34740) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_9a8ea8a59523879e) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	__obf_91b87182c3b48e5a := // if lockField != "" && !hasLockField {
		// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
		// }

		fmt.Sprintf("UPDATE %s SET %s WHERE %s", __obf_4dd2ad913c3df0d6, strings.Join(__obf_9a8ea8a59523879e, ","),
			strings.Join(__obf_4d6a41b54ff34740, " AND "),
		)

	return __obf_91b87182c3b48e5a, __obf_8670d196c9bebfda, nil
}
