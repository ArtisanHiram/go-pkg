package __obf_66da9ed7fbc1ec9e

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

func NewORM(__obf_523686877a0cd4ad *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_523686877a0cd4ad.Driver, __obf_523686877a0cd4ad.DataSource)

	switch __obf_523686877a0cd4ad.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_523686877a0cd4ad.Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_523686877a0cd4ad.Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_523686877a0cd4ad.Placeholder = ":"
	case "sqlserver":
		__obf_523686877a0cd4ad.Placeholder = "@"
	}

	if __obf_1229ad25944902c2 := DB.Ping(); __obf_1229ad25944902c2 != nil {
		panic(__obf_1229ad25944902c2)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_523686877a0cd4ad,
		}, func() {
			DB.Close()
		}
}

func (__obf_2e13600f19148084 *ORM) ListMap(__obf_ce506c59e72daf0b *sqlx.Rows, __obf_9a3e00206a1ffa07 func(map[string]any) (string, string)) (__obf_8f096a0267f35515 []map[string]any, __obf_1229ad25944902c2 error) {
	for __obf_ce506c59e72daf0b.Next() {
		__obf_81c5c86298ada95f := make(map[string]any)
		if __obf_1229ad25944902c2 = __obf_ce506c59e72daf0b.MapScan(__obf_81c5c86298ada95f); __obf_1229ad25944902c2 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_1229ad25944902c2)
		}
		for __obf_22dbe2e4f410cd44, __obf_ac35cf98df641116 := range __obf_81c5c86298ada95f {
			switch __obf_ac35cf98df641116 := __obf_ac35cf98df641116.(type) {
			case []uint8:
				__obf_81c5c86298ada95f[__obf_22dbe2e4f410cd44] = string(__obf_ac35cf98df641116)
			case int64:
				__obf_81c5c86298ada95f[__obf_22dbe2e4f410cd44] = int64(__obf_ac35cf98df641116)
			case uint64:
				__obf_81c5c86298ada95f[__obf_22dbe2e4f410cd44] = uint64(__obf_ac35cf98df641116)
			}
		}
		if __obf_9a3e00206a1ffa07 != nil {
			__obf_98e9face90be6a75, __obf_1d423fa3821811d8 := __obf_9a3e00206a1ffa07(__obf_81c5c86298ada95f)
			__obf_81c5c86298ada95f[__obf_98e9face90be6a75] = __obf_1d423fa3821811d8
		}
		__obf_8f096a0267f35515 = append(__obf_8f096a0267f35515, __obf_81c5c86298ada95f)
	}
	return
}

func (__obf_2e13600f19148084 *ORM) Exists(__obf_d2e84a17d961b34a, __obf_f818ec1012d81c56 string, __obf_66ab003bba546d1d ...any) bool {
	var __obf_24ed991c54d9f99f bool
	_ = __obf_2e13600f19148084.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_d2e84a17d961b34a, __obf_f818ec1012d81c56), __obf_66ab003bba546d1d...).Scan(&__obf_24ed991c54d9f99f)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_24ed991c54d9f99f
}

func (__obf_2e13600f19148084 *ORM) SaveModel(__obf_cd921d275317807f any, __obf_85ea3cf83c7c53a0 string) (string, error) {
	__obf_ac35cf98df641116 := reflect.ValueOf(__obf_cd921d275317807f)
	if __obf_ac35cf98df641116.Kind() == reflect.Pointer {
		__obf_ac35cf98df641116 = __obf_ac35cf98df641116.Elem()
	}
	var (
		__obf_f7322eb09388974c = __obf_ac35cf98df641116.FieldByName("Id")
		__obf_0e689b7bb7c6300b = util.ToSnake(reflect.Indirect(__obf_ac35cf98df641116).Type().Name())
		__obf_8c21f68e9672c4f9 string
		__obf_66ab003bba546d1d map[string]any
		__obf_1229ad25944902c2 error
	)

	if __obf_f7322eb09388974c.IsZero() {
		// insert
		__obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d, __obf_1229ad25944902c2 = BuildInsertSQL(__obf_cd921d275317807f, __obf_0e689b7bb7c6300b)
		if __obf_1229ad25944902c2 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_1229ad25944902c2.Error())
		}
		if __obf_85ea3cf83c7c53a0 == "" {
			__obf_85ea3cf83c7c53a0 = util.StringUUID()
		}

		__obf_66ab003bba546d1d["id"] = __obf_85ea3cf83c7c53a0
		__obf_28b57be2daaed4f8 := time.Now().Unix()
		if _, __obf_184ad66d1ce62d27 := __obf_66ab003bba546d1d[__obf_2e13600f19148084.Option.Reserved.CreateTime]; __obf_184ad66d1ce62d27 {
			__obf_66ab003bba546d1d[__obf_2e13600f19148084.Option.Reserved.CreateTime] = __obf_28b57be2daaed4f8
		}
		if _, __obf_184ad66d1ce62d27 := __obf_66ab003bba546d1d[__obf_2e13600f19148084.Option.Reserved.UpdateTime]; __obf_184ad66d1ce62d27 {
			__obf_66ab003bba546d1d[__obf_2e13600f19148084.Option.Reserved.UpdateTime] = __obf_28b57be2daaed4f8
		}
	} else {
		// update
		__obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d, __obf_1229ad25944902c2 = BuildUpdateSQL(__obf_cd921d275317807f, __obf_0e689b7bb7c6300b, __obf_2e13600f19148084.Option.Reserved.UpdateTime, []string{__obf_2e13600f19148084.Option.Reserved.ID}, __obf_2e13600f19148084.Option.Reserved.CreateTime, __obf_2e13600f19148084.Option.Reserved.NO)
		if __obf_1229ad25944902c2 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_1229ad25944902c2.Error())
		}
	}

	if __obf_8c21f68e9672c4f9 != "" {
		_, __obf_1229ad25944902c2 := __obf_2e13600f19148084.DB.NamedExec(__obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d)
		if __obf_1229ad25944902c2 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_1229ad25944902c2.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_8c21f68e9672c4f9, "args", __obf_66ab003bba546d1d)
	}

	return __obf_85ea3cf83c7c53a0, nil

}

func (__obf_2e13600f19148084 *ORM) SaveModelWidthAutoID(__obf_cd921d275317807f any, __obf_85ea3cf83c7c53a0 int64) (int64, error) {
	// query, args := modelBind(db.Option, m, insertId)
	__obf_ac35cf98df641116 := reflect.ValueOf(__obf_cd921d275317807f)
	if __obf_ac35cf98df641116.Kind() == reflect.Pointer {
		__obf_ac35cf98df641116 = __obf_ac35cf98df641116.Elem()
	}
	__obf_f7322eb09388974c := __obf_ac35cf98df641116.FieldByName("Id")
	__obf_0e689b7bb7c6300b := util.ToSnake(reflect.Indirect(__obf_ac35cf98df641116).Type().Name())
	var (
		__obf_8c21f68e9672c4f9 string
		__obf_66ab003bba546d1d map[string]any
		__obf_1229ad25944902c2 error
	)

	if __obf_f7322eb09388974c.IsZero() {
		// insert
		__obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d, __obf_1229ad25944902c2 = BuildInsertSQL(__obf_cd921d275317807f, __obf_0e689b7bb7c6300b)
		if __obf_1229ad25944902c2 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_1229ad25944902c2.Error())
		}
		if __obf_85ea3cf83c7c53a0 != 0 {
			__obf_66ab003bba546d1d["id"] = __obf_85ea3cf83c7c53a0
		}
		__obf_28b57be2daaed4f8 := time.Now().Unix()
		if _, __obf_184ad66d1ce62d27 := __obf_66ab003bba546d1d[__obf_2e13600f19148084.Option.Reserved.CreateTime]; __obf_184ad66d1ce62d27 {
			__obf_66ab003bba546d1d[__obf_2e13600f19148084.Option.Reserved.CreateTime] = __obf_28b57be2daaed4f8
		}
		if _, __obf_184ad66d1ce62d27 := __obf_66ab003bba546d1d[__obf_2e13600f19148084.Option.Reserved.UpdateTime]; __obf_184ad66d1ce62d27 {
			__obf_66ab003bba546d1d[__obf_2e13600f19148084.Option.Reserved.UpdateTime] = __obf_28b57be2daaed4f8
		}
	} else {
		// update
		__obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d, __obf_1229ad25944902c2 = BuildUpdateSQL(__obf_cd921d275317807f, __obf_0e689b7bb7c6300b, __obf_2e13600f19148084.Option.Reserved.UpdateTime, []string{__obf_2e13600f19148084.Option.Reserved.ID}, __obf_2e13600f19148084.Option.Reserved.CreateTime, __obf_2e13600f19148084.Option.Reserved.NO)
		if __obf_1229ad25944902c2 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_1229ad25944902c2.Error())
		}
	}

	if __obf_8c21f68e9672c4f9 != "" {
		__obf_1d423fa3821811d8, __obf_1229ad25944902c2 := __obf_2e13600f19148084.DB.NamedExec(__obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d)
		if __obf_1229ad25944902c2 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_1229ad25944902c2.Error())
		}
		if __obf_85ea3cf83c7c53a0 == 0 {
			return __obf_1d423fa3821811d8.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_8c21f68e9672c4f9, "args", __obf_66ab003bba546d1d)
	}

	return __obf_85ea3cf83c7c53a0, nil
}

func (__obf_2e13600f19148084 *ORM) SaveTxModel(__obf_181488bf1961013b *sqlx.Tx, __obf_cd921d275317807f any, __obf_85ea3cf83c7c53a0 string) error {
	var __obf_8c21f68e9672c4f9 string
	var __obf_66ab003bba546d1d map[string]any
	__obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d = __obf_0b37041106d33529(__obf_2e13600f19148084.Option, __obf_cd921d275317807f, __obf_85ea3cf83c7c53a0)

	if __obf_8c21f68e9672c4f9 != "" {
		if _, __obf_1229ad25944902c2 := __obf_181488bf1961013b.NamedExec(__obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d); __obf_1229ad25944902c2 != nil {
			_ = __obf_181488bf1961013b.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_1229ad25944902c2)
		}

		slog.Debug("SaveTxModel", "query", __obf_8c21f68e9672c4f9, "args", __obf_66ab003bba546d1d)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_0b37041106d33529[RecordID int | string](__obf_523686877a0cd4ad *Options, __obf_cd921d275317807f any, __obf_85ea3cf83c7c53a0 RecordID) (__obf_9aa3173ef0f4c2e5 string, __obf_66ab003bba546d1d map[string]any) {
	__obf_ac35cf98df641116 := reflect.ValueOf(__obf_cd921d275317807f)
	if __obf_ac35cf98df641116.Kind() == reflect.Pointer {
		__obf_ac35cf98df641116 = __obf_ac35cf98df641116.Elem()
	}
	var __obf_a0021a865e903262 []string
	__obf_66ab003bba546d1d = make(map[string]any)

	__obf_f7322eb09388974c := __obf_ac35cf98df641116.FieldByName("Id")
	if __obf_f7322eb09388974c.IsZero() {
		// 判断inserId是否为零值
		if __obf_85ea3cf83c7c53a0 != *new(RecordID) {
			__obf_a0021a865e903262 = append(__obf_a0021a865e903262, __obf_523686877a0cd4ad.Reserved.ID)
			__obf_66ab003bba546d1d[__obf_523686877a0cd4ad.Reserved.ID] = __obf_85ea3cf83c7c53a0
		} else {
			if __obf_f7322eb09388974c.Kind() == reflect.String {
				__obf_a0021a865e903262 = append(__obf_a0021a865e903262, __obf_523686877a0cd4ad.Reserved.ID)
				__obf_66ab003bba546d1d[__obf_523686877a0cd4ad.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}

		__obf_d9df25f8620bc5a4 := reflect.Indirect(__obf_ac35cf98df641116).Type()
		for __obf_4ceb93679168ba44 := range __obf_d9df25f8620bc5a4.NumField() {
			if __obf_bb16074987cbe782 := __obf_d9df25f8620bc5a4.Field(__obf_4ceb93679168ba44); __obf_bb16074987cbe782.Tag.Get("db") != "-" {
				switch __obf_6afb544cca5ed47d := util.ToSnake(__obf_bb16074987cbe782.Name); __obf_6afb544cca5ed47d {
				case __obf_523686877a0cd4ad.Reserved.CreateTime, __obf_523686877a0cd4ad.Reserved.UpdateTime:
					__obf_a0021a865e903262 = append(__obf_a0021a865e903262, __obf_6afb544cca5ed47d)
					__obf_66ab003bba546d1d[__obf_6afb544cca5ed47d] = time.Now().Unix()
				default:
					__obf_bc4d09a84501c296 := __obf_ac35cf98df641116.Field(__obf_4ceb93679168ba44).Interface()
					if strings.HasPrefix(__obf_6afb544cca5ed47d, "is_") || __obf_bc4d09a84501c296 != reflect.Zero(__obf_bb16074987cbe782.Type).Interface() {
						__obf_a0021a865e903262 = append(__obf_a0021a865e903262, __obf_6afb544cca5ed47d)
						__obf_66ab003bba546d1d[__obf_6afb544cca5ed47d] = __obf_bc4d09a84501c296
					}
				}
			}
		}
		__obf_9aa3173ef0f4c2e5 = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_d9df25f8620bc5a4.Name()), strings.Join(__obf_a0021a865e903262, ","), strings.Join(__obf_a0021a865e903262, ",:"))
	} else {
		__obf_d9df25f8620bc5a4 := reflect.Indirect(__obf_ac35cf98df641116).Type()
		for __obf_4ceb93679168ba44 := range __obf_d9df25f8620bc5a4.NumField() {
			if __obf_bb16074987cbe782 := __obf_d9df25f8620bc5a4.Field(__obf_4ceb93679168ba44); __obf_bb16074987cbe782.Tag.Get("db") != "-" {
				__obf_bc4d09a84501c296 := __obf_ac35cf98df641116.Field(__obf_4ceb93679168ba44).Interface()
				if __obf_6afb544cca5ed47d := util.ToSnake(__obf_bb16074987cbe782.Name); strings.HasPrefix(__obf_6afb544cca5ed47d, "is_") || __obf_bc4d09a84501c296 != reflect.Zero(__obf_bb16074987cbe782.Type).Interface() {
					__obf_a0021a865e903262 = append(__obf_a0021a865e903262, fmt.Sprintf("%s=:%s", __obf_6afb544cca5ed47d, __obf_6afb544cca5ed47d))
					__obf_66ab003bba546d1d[__obf_6afb544cca5ed47d] = __obf_bc4d09a84501c296
				}
			}
		}

		if util.HasField(__obf_ac35cf98df641116, util.ToCamel(__obf_523686877a0cd4ad.Reserved.UpdateTime)) {
			__obf_a7fbaa3425c3355a := __obf_66ab003bba546d1d[__obf_523686877a0cd4ad.Reserved.UpdateTime]
			__obf_66ab003bba546d1d[__obf_523686877a0cd4ad.Reserved.UpdateTime] = time.Now().Unix()
			__obf_9aa3173ef0f4c2e5 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_d9df25f8620bc5a4.Name()), strings.Join(__obf_a0021a865e903262, ","), __obf_523686877a0cd4ad.Reserved.ID, __obf_523686877a0cd4ad.Reserved.ID, __obf_523686877a0cd4ad.Reserved.UpdateTime, __obf_a7fbaa3425c3355a)
		} else {
			__obf_9aa3173ef0f4c2e5 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_d9df25f8620bc5a4.Name()), strings.Join(__obf_a0021a865e903262, ","), __obf_523686877a0cd4ad.Reserved.ID, __obf_523686877a0cd4ad.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_6c2ef9edf771bc1c any, __obf_0e689b7bb7c6300b string) (string, map[string]any, error) {
	if __obf_0e689b7bb7c6300b == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_66ab003bba546d1d := make(map[string]any)
	__obf_a0021a865e903262 := []string{}

	__obf_b17117511740e68d := func(__obf_7527be360ce64e08 string, __obf_830db7a4c207dec8 any, __obf_7f5b9bc7589dc19f bool) {
		__obf_9cc545a40f46a1b4 := __obf_830db7a4c207dec8 == nil || reflect.DeepEqual(__obf_830db7a4c207dec8, reflect.Zero(reflect.TypeOf(__obf_830db7a4c207dec8)).Interface())
		if __obf_7f5b9bc7589dc19f && __obf_9cc545a40f46a1b4 {
			return
		}
		__obf_a0021a865e903262 = append(__obf_a0021a865e903262, __obf_7527be360ce64e08)
		__obf_66ab003bba546d1d[__obf_7527be360ce64e08] = __obf_830db7a4c207dec8
	}

	switch __obf_bc4d09a84501c296 := __obf_6c2ef9edf771bc1c.(type) {
	case map[string]any:
		for __obf_22dbe2e4f410cd44, __obf_ac35cf98df641116 := range __obf_bc4d09a84501c296 {
			__obf_b17117511740e68d(__obf_22dbe2e4f410cd44, __obf_ac35cf98df641116, false)
		}

	default:
		__obf_7d1378dcb9668016 := reflect.TypeOf(__obf_6c2ef9edf771bc1c)
		if __obf_7d1378dcb9668016.Kind() == reflect.Pointer {
			__obf_7d1378dcb9668016 = __obf_7d1378dcb9668016.Elem()
		}
		__obf_ac35cf98df641116 := reflect.ValueOf(__obf_6c2ef9edf771bc1c)
		if __obf_ac35cf98df641116.Kind() == reflect.Pointer {
			__obf_ac35cf98df641116 = __obf_ac35cf98df641116.Elem()
		}

		if __obf_7d1378dcb9668016.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_4ceb93679168ba44 := 0; __obf_4ceb93679168ba44 < __obf_7d1378dcb9668016.NumField(); __obf_4ceb93679168ba44++ {
			__obf_a95d5a81a7fcb1d4 := __obf_7d1378dcb9668016.Field(__obf_4ceb93679168ba44).Tag.Get("db")
			if __obf_a95d5a81a7fcb1d4 == "" || __obf_a95d5a81a7fcb1d4 == "-" {
				continue
			}

			__obf_be58c044629c4868 := strings.Split(__obf_a95d5a81a7fcb1d4, ",")
			__obf_b17117511740e68d(__obf_be58c044629c4868[0], __obf_ac35cf98df641116.Field(__obf_4ceb93679168ba44).Interface(), len(__obf_be58c044629c4868) > 1 && __obf_be58c044629c4868[1] == "omitempty")
		}
	}

	if len(__obf_a0021a865e903262) == 0 {
		return "", nil, errors.New("no fields to insert")
	}

	__obf_8c21f68e9672c4f9 := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)",
		__obf_0e689b7bb7c6300b,
		strings.Join(__obf_a0021a865e903262, ","),
		strings.Join(__obf_a0021a865e903262, ",:"),
	)

	return __obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d, nil
}

// pkField 主键字段
// lockField 乐观锁字段(默认为updated_at)
// constField 固定不变字段
func BuildUpdateSQL(
	__obf_6c2ef9edf771bc1c any,
	__obf_0e689b7bb7c6300b, __obf_16877c611e62efad string,
	__obf_ae6896a443a247f8 []string,
	__obf_d5f5c32d83b6aecd ...string,
) (string, map[string]any, error) {
	if __obf_0e689b7bb7c6300b == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_5af4b472a5d2e445 := make(map[string]struct{})
	for _, __obf_5e805ca92bfd3fe7 := range __obf_ae6896a443a247f8 {
		__obf_5af4b472a5d2e445[__obf_5e805ca92bfd3fe7] = struct{}{}
	}

	__obf_66ab003bba546d1d := make(map[string]any)
	__obf_a0021a865e903262 := []string{}
	__obf_3ece7ffbc92b3e59 := []string{}

	__obf_b17117511740e68d := func(__obf_7527be360ce64e08 string, __obf_830db7a4c207dec8 any, __obf_7f5b9bc7589dc19f bool) {
		__obf_9cc545a40f46a1b4 := __obf_830db7a4c207dec8 == nil || reflect.DeepEqual(__obf_830db7a4c207dec8, reflect.Zero(reflect.TypeOf(__obf_830db7a4c207dec8)).Interface())

		switch {
		case func() bool {
			_, __obf_184ad66d1ce62d27 := __obf_5af4b472a5d2e445[__obf_7527be360ce64e08]
			return __obf_184ad66d1ce62d27
		}():
			__obf_3ece7ffbc92b3e59 = append(__obf_3ece7ffbc92b3e59, fmt.Sprintf("%s=:%s", __obf_7527be360ce64e08, __obf_7527be360ce64e08))
			__obf_66ab003bba546d1d[__obf_7527be360ce64e08] = __obf_830db7a4c207dec8

		case slices.Contains(__obf_d5f5c32d83b6aecd, __obf_7527be360ce64e08):
			// skip constField
		case __obf_7527be360ce64e08 == __obf_16877c611e62efad:
			const __obf_40d1a5a618836b70 = "RESERVED_LOCK"
			__obf_a0021a865e903262 = append(__obf_a0021a865e903262, fmt.Sprintf("%s=:%s", __obf_7527be360ce64e08, __obf_7527be360ce64e08))
			__obf_3ece7ffbc92b3e59 = append(__obf_3ece7ffbc92b3e59, fmt.Sprintf("%s=:%s", __obf_7527be360ce64e08, __obf_40d1a5a618836b70))
			__obf_66ab003bba546d1d[__obf_7527be360ce64e08] = time.Now().Unix()
			__obf_66ab003bba546d1d[__obf_40d1a5a618836b70] = __obf_830db7a4c207dec8

		default:
			if __obf_7f5b9bc7589dc19f && __obf_9cc545a40f46a1b4 {
				return
			}
			__obf_a0021a865e903262 = append(__obf_a0021a865e903262, fmt.Sprintf("%s=:%s", __obf_7527be360ce64e08, __obf_7527be360ce64e08))
			__obf_66ab003bba546d1d[__obf_7527be360ce64e08] = __obf_830db7a4c207dec8
		}
	}

	switch __obf_bc4d09a84501c296 := __obf_6c2ef9edf771bc1c.(type) {
	case map[string]any:
		for __obf_22dbe2e4f410cd44, __obf_ac35cf98df641116 := range __obf_bc4d09a84501c296 {
			__obf_b17117511740e68d(__obf_22dbe2e4f410cd44, __obf_ac35cf98df641116, false)
		}

	default:

		__obf_7d1378dcb9668016 := reflect.TypeOf(__obf_6c2ef9edf771bc1c)
		if __obf_7d1378dcb9668016.Kind() == reflect.Pointer {
			__obf_7d1378dcb9668016 = __obf_7d1378dcb9668016.Elem()
		}
		__obf_ac35cf98df641116 := reflect.ValueOf(__obf_6c2ef9edf771bc1c)
		if __obf_ac35cf98df641116.Kind() == reflect.Pointer {
			__obf_ac35cf98df641116 = __obf_ac35cf98df641116.Elem()
		}
		if __obf_7d1378dcb9668016.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_4ceb93679168ba44 := 0; __obf_4ceb93679168ba44 < __obf_7d1378dcb9668016.NumField(); __obf_4ceb93679168ba44++ {
			__obf_a95d5a81a7fcb1d4 := __obf_7d1378dcb9668016.Field(__obf_4ceb93679168ba44).Tag.Get("db")
			if __obf_a95d5a81a7fcb1d4 == "" || __obf_a95d5a81a7fcb1d4 == "-" {
				continue
			}

			__obf_be58c044629c4868 := strings.Split(__obf_a95d5a81a7fcb1d4, ",")
			__obf_b17117511740e68d(__obf_be58c044629c4868[0], __obf_ac35cf98df641116.Field(__obf_4ceb93679168ba44).Interface(), len(__obf_be58c044629c4868) > 1 && __obf_be58c044629c4868[1] == "omitempty")
		}
	}

	if len(__obf_ae6896a443a247f8) == 0 || len(__obf_3ece7ffbc92b3e59) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_a0021a865e903262) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	// if lockField != "" && !hasLockField {
	// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
	// }

	__obf_8c21f68e9672c4f9 := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		__obf_0e689b7bb7c6300b,
		strings.Join(__obf_a0021a865e903262, ","),
		strings.Join(__obf_3ece7ffbc92b3e59, " AND "),
	)

	return __obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d, nil
}
