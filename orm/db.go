package __obf_24a21e7f5375b619

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

func NewORM(__obf_b7f68c096f8695ad *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_b7f68c096f8695ad.Driver, __obf_b7f68c096f8695ad.DataSource)

	switch __obf_b7f68c096f8695ad.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_b7f68c096f8695ad.Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_b7f68c096f8695ad.Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_b7f68c096f8695ad.Placeholder = ":"
	case "sqlserver":
		__obf_b7f68c096f8695ad.Placeholder = "@"
	}

	if __obf_a687171eb2d5acc2 := DB.Ping(); __obf_a687171eb2d5acc2 != nil {
		panic(__obf_a687171eb2d5acc2)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_b7f68c096f8695ad,
		}, func() {
			DB.Close()
		}
}

func (__obf_80cd76b5279c3275 *ORM) ListMap(__obf_d1a5f8eb962a8a6f *sqlx.Rows, __obf_52add46828e4c3dd func(map[string]any) (string, string)) (__obf_098085715c2fedd0 []map[string]any, __obf_a687171eb2d5acc2 error) {
	for __obf_d1a5f8eb962a8a6f.Next() {
		__obf_d4893020749434f1 := make(map[string]any)
		if __obf_a687171eb2d5acc2 = __obf_d1a5f8eb962a8a6f.MapScan(__obf_d4893020749434f1); __obf_a687171eb2d5acc2 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_a687171eb2d5acc2)
		}
		for __obf_388eb12efe355036, __obf_bbb98599fdbc54eb := range __obf_d4893020749434f1 {
			switch __obf_bbb98599fdbc54eb := __obf_bbb98599fdbc54eb.(type) {
			case []uint8:
				__obf_d4893020749434f1[__obf_388eb12efe355036] = string(__obf_bbb98599fdbc54eb)
			case int64:
				__obf_d4893020749434f1[__obf_388eb12efe355036] = int64(__obf_bbb98599fdbc54eb)
			case uint64:
				__obf_d4893020749434f1[__obf_388eb12efe355036] = uint64(__obf_bbb98599fdbc54eb)
			}
		}
		if __obf_52add46828e4c3dd != nil {
			__obf_52c60063fb98b213, __obf_5b7d2bf1f9566b41 := __obf_52add46828e4c3dd(__obf_d4893020749434f1)
			__obf_d4893020749434f1[__obf_52c60063fb98b213] = __obf_5b7d2bf1f9566b41
		}
		__obf_098085715c2fedd0 = append(__obf_098085715c2fedd0, __obf_d4893020749434f1)
	}
	return
}

func (__obf_80cd76b5279c3275 *ORM) Exists(__obf_1376e33d76faeb06, __obf_bfb42fff03b52307 string, __obf_377e2b8c06e9f32b ...any) bool {
	var __obf_ab7465756ef50050 bool
	_ = __obf_80cd76b5279c3275.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_1376e33d76faeb06, __obf_bfb42fff03b52307), __obf_377e2b8c06e9f32b...).Scan(&__obf_ab7465756ef50050)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_ab7465756ef50050
}

func (__obf_80cd76b5279c3275 *ORM) SaveModel(__obf_e1556accd205ae51 any, __obf_71bda281e3129a0c string) (string, error) {
	__obf_bbb98599fdbc54eb := reflect.ValueOf(__obf_e1556accd205ae51)
	if __obf_bbb98599fdbc54eb.Kind() == reflect.Pointer {
		__obf_bbb98599fdbc54eb = __obf_bbb98599fdbc54eb.Elem()
	}
	var (
		__obf_8e4e3643ee163fff = __obf_bbb98599fdbc54eb.FieldByName("Id")
		__obf_4fc2ce20fd217a61 = util.ToSnake(reflect.Indirect(__obf_bbb98599fdbc54eb).Type().Name())
		__obf_9cea37a5fd264d1a string
		__obf_377e2b8c06e9f32b map[string]any
		__obf_a687171eb2d5acc2 error
	)

	if __obf_8e4e3643ee163fff.IsZero() {
		// insert
		__obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b, __obf_a687171eb2d5acc2 = BuildInsertSQL(__obf_e1556accd205ae51, __obf_4fc2ce20fd217a61)
		if __obf_a687171eb2d5acc2 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_a687171eb2d5acc2.Error())
		}
		if __obf_71bda281e3129a0c == "" {
			__obf_71bda281e3129a0c = util.StringUUID()
		}

		__obf_377e2b8c06e9f32b["id"] = __obf_71bda281e3129a0c
		__obf_b7274a2e7d910c2c := time.Now().Unix()
		if _, __obf_bfcf4076c7714374 := __obf_377e2b8c06e9f32b[__obf_80cd76b5279c3275.Option.Reserved.CreateTime]; __obf_bfcf4076c7714374 {
			__obf_377e2b8c06e9f32b[__obf_80cd76b5279c3275.Option.Reserved.CreateTime] = __obf_b7274a2e7d910c2c
		}
		if _, __obf_bfcf4076c7714374 := __obf_377e2b8c06e9f32b[__obf_80cd76b5279c3275.Option.Reserved.UpdateTime]; __obf_bfcf4076c7714374 {
			__obf_377e2b8c06e9f32b[__obf_80cd76b5279c3275.Option.Reserved.UpdateTime] = __obf_b7274a2e7d910c2c
		}
	} else {
		// update
		__obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b, __obf_a687171eb2d5acc2 = BuildUpdateSQL(__obf_e1556accd205ae51, __obf_4fc2ce20fd217a61, __obf_80cd76b5279c3275.Option.Reserved.UpdateTime, []string{__obf_80cd76b5279c3275.Option.Reserved.ID}, __obf_80cd76b5279c3275.Option.Reserved.CreateTime, __obf_80cd76b5279c3275.Option.Reserved.NO)
		if __obf_a687171eb2d5acc2 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_a687171eb2d5acc2.Error())
		}
	}

	if __obf_9cea37a5fd264d1a != "" {
		_, __obf_a687171eb2d5acc2 := __obf_80cd76b5279c3275.DB.NamedExec(__obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b)
		if __obf_a687171eb2d5acc2 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_a687171eb2d5acc2.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_9cea37a5fd264d1a, "args", __obf_377e2b8c06e9f32b)
	}

	return __obf_71bda281e3129a0c, nil

}

func (__obf_80cd76b5279c3275 *ORM) SaveModelWidthAutoID(__obf_e1556accd205ae51 any, __obf_71bda281e3129a0c int64) (int64, error) {
	// query, args := modelBind(db.Option, m, insertId)
	__obf_bbb98599fdbc54eb := reflect.ValueOf(__obf_e1556accd205ae51)
	if __obf_bbb98599fdbc54eb.Kind() == reflect.Pointer {
		__obf_bbb98599fdbc54eb = __obf_bbb98599fdbc54eb.Elem()
	}
	__obf_8e4e3643ee163fff := __obf_bbb98599fdbc54eb.FieldByName("Id")
	__obf_4fc2ce20fd217a61 := util.ToSnake(reflect.Indirect(__obf_bbb98599fdbc54eb).Type().Name())
	var (
		__obf_9cea37a5fd264d1a string
		__obf_377e2b8c06e9f32b map[string]any
		__obf_a687171eb2d5acc2 error
	)

	if __obf_8e4e3643ee163fff.IsZero() {
		// insert
		__obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b, __obf_a687171eb2d5acc2 = BuildInsertSQL(__obf_e1556accd205ae51, __obf_4fc2ce20fd217a61)
		if __obf_a687171eb2d5acc2 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_a687171eb2d5acc2.Error())
		}
		if __obf_71bda281e3129a0c != 0 {
			__obf_377e2b8c06e9f32b["id"] = __obf_71bda281e3129a0c
		}
		__obf_b7274a2e7d910c2c := time.Now().Unix()
		if _, __obf_bfcf4076c7714374 := __obf_377e2b8c06e9f32b[__obf_80cd76b5279c3275.Option.Reserved.CreateTime]; __obf_bfcf4076c7714374 {
			__obf_377e2b8c06e9f32b[__obf_80cd76b5279c3275.Option.Reserved.CreateTime] = __obf_b7274a2e7d910c2c
		}
		if _, __obf_bfcf4076c7714374 := __obf_377e2b8c06e9f32b[__obf_80cd76b5279c3275.Option.Reserved.UpdateTime]; __obf_bfcf4076c7714374 {
			__obf_377e2b8c06e9f32b[__obf_80cd76b5279c3275.Option.Reserved.UpdateTime] = __obf_b7274a2e7d910c2c
		}
	} else {
		// update
		__obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b, __obf_a687171eb2d5acc2 = BuildUpdateSQL(__obf_e1556accd205ae51, __obf_4fc2ce20fd217a61, __obf_80cd76b5279c3275.Option.Reserved.UpdateTime, []string{__obf_80cd76b5279c3275.Option.Reserved.ID}, __obf_80cd76b5279c3275.Option.Reserved.CreateTime, __obf_80cd76b5279c3275.Option.Reserved.NO)
		if __obf_a687171eb2d5acc2 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_a687171eb2d5acc2.Error())
		}
	}

	if __obf_9cea37a5fd264d1a != "" {
		__obf_5b7d2bf1f9566b41, __obf_a687171eb2d5acc2 := __obf_80cd76b5279c3275.DB.NamedExec(__obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b)
		if __obf_a687171eb2d5acc2 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_a687171eb2d5acc2.Error())
		}
		if __obf_71bda281e3129a0c == 0 {
			return __obf_5b7d2bf1f9566b41.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_9cea37a5fd264d1a, "args", __obf_377e2b8c06e9f32b)
	}

	return __obf_71bda281e3129a0c, nil
}

func (__obf_80cd76b5279c3275 *ORM) SaveTxModel(__obf_5c86968691329dea *sqlx.Tx, __obf_e1556accd205ae51 any, __obf_71bda281e3129a0c string) error {
	var __obf_9cea37a5fd264d1a string
	var __obf_377e2b8c06e9f32b map[string]any
	__obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b = __obf_9411f810cd49f6ef(__obf_80cd76b5279c3275.Option, __obf_e1556accd205ae51, __obf_71bda281e3129a0c)

	if __obf_9cea37a5fd264d1a != "" {
		if _, __obf_a687171eb2d5acc2 := __obf_5c86968691329dea.NamedExec(__obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b); __obf_a687171eb2d5acc2 != nil {
			_ = __obf_5c86968691329dea.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_a687171eb2d5acc2)
		}

		slog.Debug("SaveTxModel", "query", __obf_9cea37a5fd264d1a, "args", __obf_377e2b8c06e9f32b)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_9411f810cd49f6ef[RecordID int | string](__obf_b7f68c096f8695ad *Options, __obf_e1556accd205ae51 any, __obf_71bda281e3129a0c RecordID) (__obf_595f5ea64f010d17 string, __obf_377e2b8c06e9f32b map[string]any) {
	__obf_bbb98599fdbc54eb := reflect.ValueOf(__obf_e1556accd205ae51)
	if __obf_bbb98599fdbc54eb.Kind() == reflect.Pointer {
		__obf_bbb98599fdbc54eb = __obf_bbb98599fdbc54eb.Elem()
	}
	var __obf_83867815ff46bde4 []string
	__obf_377e2b8c06e9f32b = make(map[string]any)

	__obf_8e4e3643ee163fff := __obf_bbb98599fdbc54eb.FieldByName("Id")
	if __obf_8e4e3643ee163fff.IsZero() {
		// 判断inserId是否为零值
		if __obf_71bda281e3129a0c != *new(RecordID) {
			__obf_83867815ff46bde4 = append(__obf_83867815ff46bde4, __obf_b7f68c096f8695ad.Reserved.ID)
			__obf_377e2b8c06e9f32b[__obf_b7f68c096f8695ad.Reserved.ID] = __obf_71bda281e3129a0c
		} else {
			if __obf_8e4e3643ee163fff.Kind() == reflect.String {
				__obf_83867815ff46bde4 = append(__obf_83867815ff46bde4, __obf_b7f68c096f8695ad.Reserved.ID)
				__obf_377e2b8c06e9f32b[__obf_b7f68c096f8695ad.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}

		__obf_7e2d9a58f8ca6d57 := reflect.Indirect(__obf_bbb98599fdbc54eb).Type()
		for __obf_f4c5dc7b1f70e707 := range __obf_7e2d9a58f8ca6d57.NumField() {
			if __obf_57b980f12a7a4617 := __obf_7e2d9a58f8ca6d57.Field(__obf_f4c5dc7b1f70e707); __obf_57b980f12a7a4617.Tag.Get("db") != "-" {
				switch __obf_b51f4581f1849e6f := util.ToSnake(__obf_57b980f12a7a4617.Name); __obf_b51f4581f1849e6f {
				case __obf_b7f68c096f8695ad.Reserved.CreateTime, __obf_b7f68c096f8695ad.Reserved.UpdateTime:
					__obf_83867815ff46bde4 = append(__obf_83867815ff46bde4, __obf_b51f4581f1849e6f)
					__obf_377e2b8c06e9f32b[__obf_b51f4581f1849e6f] = time.Now().Unix()
				default:
					__obf_9e8e83e850d24f6c := __obf_bbb98599fdbc54eb.Field(__obf_f4c5dc7b1f70e707).Interface()
					if strings.HasPrefix(__obf_b51f4581f1849e6f, "is_") || __obf_9e8e83e850d24f6c != reflect.Zero(__obf_57b980f12a7a4617.Type).Interface() {
						__obf_83867815ff46bde4 = append(__obf_83867815ff46bde4, __obf_b51f4581f1849e6f)
						__obf_377e2b8c06e9f32b[__obf_b51f4581f1849e6f] = __obf_9e8e83e850d24f6c
					}
				}
			}
		}
		__obf_595f5ea64f010d17 = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_7e2d9a58f8ca6d57.Name()), strings.Join(__obf_83867815ff46bde4, ","), strings.Join(__obf_83867815ff46bde4, ",:"))
	} else {
		__obf_7e2d9a58f8ca6d57 := reflect.Indirect(__obf_bbb98599fdbc54eb).Type()
		for __obf_f4c5dc7b1f70e707 := range __obf_7e2d9a58f8ca6d57.NumField() {
			if __obf_57b980f12a7a4617 := __obf_7e2d9a58f8ca6d57.Field(__obf_f4c5dc7b1f70e707); __obf_57b980f12a7a4617.Tag.Get("db") != "-" {
				__obf_9e8e83e850d24f6c := __obf_bbb98599fdbc54eb.Field(__obf_f4c5dc7b1f70e707).Interface()
				if __obf_b51f4581f1849e6f := util.ToSnake(__obf_57b980f12a7a4617.Name); strings.HasPrefix(__obf_b51f4581f1849e6f, "is_") || __obf_9e8e83e850d24f6c != reflect.Zero(__obf_57b980f12a7a4617.Type).Interface() {
					__obf_83867815ff46bde4 = append(__obf_83867815ff46bde4, fmt.Sprintf("%s=:%s", __obf_b51f4581f1849e6f, __obf_b51f4581f1849e6f))
					__obf_377e2b8c06e9f32b[__obf_b51f4581f1849e6f] = __obf_9e8e83e850d24f6c
				}
			}
		}

		if util.HasField(__obf_bbb98599fdbc54eb, util.ToCamel(__obf_b7f68c096f8695ad.Reserved.UpdateTime)) {
			__obf_956608398fafc1da := __obf_377e2b8c06e9f32b[__obf_b7f68c096f8695ad.Reserved.UpdateTime]
			__obf_377e2b8c06e9f32b[__obf_b7f68c096f8695ad.Reserved.UpdateTime] = time.Now().Unix()
			__obf_595f5ea64f010d17 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_7e2d9a58f8ca6d57.Name()), strings.Join(__obf_83867815ff46bde4, ","), __obf_b7f68c096f8695ad.Reserved.ID, __obf_b7f68c096f8695ad.Reserved.ID, __obf_b7f68c096f8695ad.Reserved.UpdateTime, __obf_956608398fafc1da)
		} else {
			__obf_595f5ea64f010d17 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_7e2d9a58f8ca6d57.Name()), strings.Join(__obf_83867815ff46bde4, ","), __obf_b7f68c096f8695ad.Reserved.ID, __obf_b7f68c096f8695ad.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_2380b67aaabc10f7 any, __obf_4fc2ce20fd217a61 string) (string, map[string]any, error) {
	if __obf_4fc2ce20fd217a61 == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_377e2b8c06e9f32b := make(map[string]any)
	__obf_83867815ff46bde4 := []string{}

	__obf_c3f0599695f99aad := func(__obf_5b84b0458a74fd86 string, __obf_0826efe2d8c95bbc any, __obf_0e757268af578238 bool) {
		__obf_84a5c6250dec847a := __obf_0826efe2d8c95bbc == nil || reflect.DeepEqual(__obf_0826efe2d8c95bbc, reflect.Zero(reflect.TypeOf(__obf_0826efe2d8c95bbc)).Interface())
		if __obf_0e757268af578238 && __obf_84a5c6250dec847a {
			return
		}
		__obf_83867815ff46bde4 = append(__obf_83867815ff46bde4, __obf_5b84b0458a74fd86)
		__obf_377e2b8c06e9f32b[__obf_5b84b0458a74fd86] = __obf_0826efe2d8c95bbc
	}

	switch __obf_9e8e83e850d24f6c := __obf_2380b67aaabc10f7.(type) {
	case map[string]any:
		for __obf_388eb12efe355036, __obf_bbb98599fdbc54eb := range __obf_9e8e83e850d24f6c {
			__obf_c3f0599695f99aad(__obf_388eb12efe355036, __obf_bbb98599fdbc54eb, false)
		}

	default:
		__obf_4224178bd2feb4fe := reflect.TypeOf(__obf_2380b67aaabc10f7)
		if __obf_4224178bd2feb4fe.Kind() == reflect.Pointer {
			__obf_4224178bd2feb4fe = __obf_4224178bd2feb4fe.Elem()
		}
		__obf_bbb98599fdbc54eb := reflect.ValueOf(__obf_2380b67aaabc10f7)
		if __obf_bbb98599fdbc54eb.Kind() == reflect.Pointer {
			__obf_bbb98599fdbc54eb = __obf_bbb98599fdbc54eb.Elem()
		}

		if __obf_4224178bd2feb4fe.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_f4c5dc7b1f70e707 := 0; __obf_f4c5dc7b1f70e707 < __obf_4224178bd2feb4fe.NumField(); __obf_f4c5dc7b1f70e707++ {
			__obf_01af7303ccf65b29 := __obf_4224178bd2feb4fe.Field(__obf_f4c5dc7b1f70e707).Tag.Get("db")
			if __obf_01af7303ccf65b29 == "" || __obf_01af7303ccf65b29 == "-" {
				continue
			}

			__obf_165c1158355d6041 := strings.Split(__obf_01af7303ccf65b29, ",")
			__obf_c3f0599695f99aad(__obf_165c1158355d6041[0], __obf_bbb98599fdbc54eb.Field(__obf_f4c5dc7b1f70e707).Interface(), len(__obf_165c1158355d6041) > 1 && __obf_165c1158355d6041[1] == "omitempty")
		}
	}

	if len(__obf_83867815ff46bde4) == 0 {
		return "", nil, errors.New("no fields to insert")
	}

	__obf_9cea37a5fd264d1a := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)",
		__obf_4fc2ce20fd217a61,
		strings.Join(__obf_83867815ff46bde4, ","),
		strings.Join(__obf_83867815ff46bde4, ",:"),
	)

	return __obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b, nil
}

// pkField 主键字段
// lockField 乐观锁字段(默认为updated_at)
// constField 固定不变字段
func BuildUpdateSQL(
	__obf_2380b67aaabc10f7 any,
	__obf_4fc2ce20fd217a61, __obf_0a30b833ce8b56dd string,
	__obf_06fc2a99fcbf225a []string,
	__obf_e5e9862cceb34156 ...string,
) (string, map[string]any, error) {
	if __obf_4fc2ce20fd217a61 == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_b177bf6bfa8fb1fb := make(map[string]struct{})
	for _, __obf_1a41b2262c731def := range __obf_06fc2a99fcbf225a {
		__obf_b177bf6bfa8fb1fb[__obf_1a41b2262c731def] = struct{}{}
	}

	__obf_377e2b8c06e9f32b := make(map[string]any)
	__obf_83867815ff46bde4 := []string{}
	__obf_cbb673ebbc8bffd4 := []string{}

	__obf_c3f0599695f99aad := func(__obf_5b84b0458a74fd86 string, __obf_0826efe2d8c95bbc any, __obf_0e757268af578238 bool) {
		__obf_84a5c6250dec847a := __obf_0826efe2d8c95bbc == nil || reflect.DeepEqual(__obf_0826efe2d8c95bbc, reflect.Zero(reflect.TypeOf(__obf_0826efe2d8c95bbc)).Interface())

		switch {
		case func() bool {
			_, __obf_bfcf4076c7714374 := __obf_b177bf6bfa8fb1fb[__obf_5b84b0458a74fd86]
			return __obf_bfcf4076c7714374
		}():
			__obf_cbb673ebbc8bffd4 = append(__obf_cbb673ebbc8bffd4, fmt.Sprintf("%s=:%s", __obf_5b84b0458a74fd86, __obf_5b84b0458a74fd86))
			__obf_377e2b8c06e9f32b[__obf_5b84b0458a74fd86] = __obf_0826efe2d8c95bbc

		case slices.Contains(__obf_e5e9862cceb34156, __obf_5b84b0458a74fd86):
			// skip constField
		case __obf_5b84b0458a74fd86 == __obf_0a30b833ce8b56dd:
			const __obf_fd5acab55827e085 = "RESERVED_LOCK"
			__obf_83867815ff46bde4 = append(__obf_83867815ff46bde4, fmt.Sprintf("%s=:%s", __obf_5b84b0458a74fd86, __obf_5b84b0458a74fd86))
			__obf_cbb673ebbc8bffd4 = append(__obf_cbb673ebbc8bffd4, fmt.Sprintf("%s=:%s", __obf_5b84b0458a74fd86, __obf_fd5acab55827e085))
			__obf_377e2b8c06e9f32b[__obf_5b84b0458a74fd86] = time.Now().Unix()
			__obf_377e2b8c06e9f32b[__obf_fd5acab55827e085] = __obf_0826efe2d8c95bbc

		default:
			if __obf_0e757268af578238 && __obf_84a5c6250dec847a {
				return
			}
			__obf_83867815ff46bde4 = append(__obf_83867815ff46bde4, fmt.Sprintf("%s=:%s", __obf_5b84b0458a74fd86, __obf_5b84b0458a74fd86))
			__obf_377e2b8c06e9f32b[__obf_5b84b0458a74fd86] = __obf_0826efe2d8c95bbc
		}
	}

	switch __obf_9e8e83e850d24f6c := __obf_2380b67aaabc10f7.(type) {
	case map[string]any:
		for __obf_388eb12efe355036, __obf_bbb98599fdbc54eb := range __obf_9e8e83e850d24f6c {
			__obf_c3f0599695f99aad(__obf_388eb12efe355036, __obf_bbb98599fdbc54eb, false)
		}

	default:

		__obf_4224178bd2feb4fe := reflect.TypeOf(__obf_2380b67aaabc10f7)
		if __obf_4224178bd2feb4fe.Kind() == reflect.Pointer {
			__obf_4224178bd2feb4fe = __obf_4224178bd2feb4fe.Elem()
		}
		__obf_bbb98599fdbc54eb := reflect.ValueOf(__obf_2380b67aaabc10f7)
		if __obf_bbb98599fdbc54eb.Kind() == reflect.Pointer {
			__obf_bbb98599fdbc54eb = __obf_bbb98599fdbc54eb.Elem()
		}
		if __obf_4224178bd2feb4fe.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_f4c5dc7b1f70e707 := 0; __obf_f4c5dc7b1f70e707 < __obf_4224178bd2feb4fe.NumField(); __obf_f4c5dc7b1f70e707++ {
			__obf_01af7303ccf65b29 := __obf_4224178bd2feb4fe.Field(__obf_f4c5dc7b1f70e707).Tag.Get("db")
			if __obf_01af7303ccf65b29 == "" || __obf_01af7303ccf65b29 == "-" {
				continue
			}

			__obf_165c1158355d6041 := strings.Split(__obf_01af7303ccf65b29, ",")
			__obf_c3f0599695f99aad(__obf_165c1158355d6041[0], __obf_bbb98599fdbc54eb.Field(__obf_f4c5dc7b1f70e707).Interface(), len(__obf_165c1158355d6041) > 1 && __obf_165c1158355d6041[1] == "omitempty")
		}
	}

	if len(__obf_06fc2a99fcbf225a) == 0 || len(__obf_cbb673ebbc8bffd4) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_83867815ff46bde4) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	// if lockField != "" && !hasLockField {
	// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
	// }

	__obf_9cea37a5fd264d1a := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		__obf_4fc2ce20fd217a61,
		strings.Join(__obf_83867815ff46bde4, ","),
		strings.Join(__obf_cbb673ebbc8bffd4, " AND "),
	)

	return __obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b, nil
}
