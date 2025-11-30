package __obf_86786288bdd26c8f

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

func NewORM(__obf_49a8263c0b23a89d *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_49a8263c0b23a89d.Driver, __obf_49a8263c0b23a89d.DataSource)

	switch __obf_49a8263c0b23a89d.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_49a8263c0b23a89d.
			Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_49a8263c0b23a89d.
			Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_49a8263c0b23a89d.
			Placeholder = ":"
	case "sqlserver":
		__obf_49a8263c0b23a89d.
			Placeholder = "@"
	}

	if __obf_87239f12e0f48308 := DB.Ping(); __obf_87239f12e0f48308 != nil {
		panic(__obf_87239f12e0f48308)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_49a8263c0b23a89d,
		}, func() {
			DB.Close()
		}
}

func (__obf_019b2ae2e7d131c0 *ORM) ListMap(__obf_6751482a63a2d802 *sqlx.Rows, __obf_125393a2b2645f41 func(map[string]any) (string, string)) (__obf_d46b9a4ac0947f8a []map[string]any, __obf_87239f12e0f48308 error) {
	for __obf_6751482a63a2d802.Next() {
		__obf_eedffc917008218d := make(map[string]any)
		if __obf_87239f12e0f48308 = __obf_6751482a63a2d802.MapScan(__obf_eedffc917008218d); __obf_87239f12e0f48308 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_87239f12e0f48308)
		}
		for __obf_299b79baf6e9ee26, __obf_0c15a3a65dfe2b86 := range __obf_eedffc917008218d {
			switch __obf_0c15a3a65dfe2b86 := __obf_0c15a3a65dfe2b86.(type) {
			case []uint8:
				__obf_eedffc917008218d[__obf_299b79baf6e9ee26] = string(__obf_0c15a3a65dfe2b86)
			case int64:
				__obf_eedffc917008218d[__obf_299b79baf6e9ee26] = int64(__obf_0c15a3a65dfe2b86)
			case uint64:
				__obf_eedffc917008218d[__obf_299b79baf6e9ee26] = uint64(__obf_0c15a3a65dfe2b86)
			}
		}
		if __obf_125393a2b2645f41 != nil {
			__obf_69d08bfacbecaaaa, __obf_8ed96d1577fd8cf6 := __obf_125393a2b2645f41(__obf_eedffc917008218d)
			__obf_eedffc917008218d[__obf_69d08bfacbecaaaa] = __obf_8ed96d1577fd8cf6
		}
		__obf_d46b9a4ac0947f8a = append(__obf_d46b9a4ac0947f8a, __obf_eedffc917008218d)
	}
	return
}

func (__obf_019b2ae2e7d131c0 *ORM) Exists(__obf_9545f0864bb74775, __obf_3cb9d69e303725bf string, __obf_d51eb37f5f1cb276 ...any) bool {
	var __obf_353197c423c1b70c bool
	_ = __obf_019b2ae2e7d131c0.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_9545f0864bb74775, __obf_3cb9d69e303725bf), __obf_d51eb37f5f1cb276...).Scan(&__obf_353197c423c1b70c)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_353197c423c1b70c
}

func (__obf_019b2ae2e7d131c0 *ORM) SaveModel(__obf_e6ccc30272d0143b any, __obf_e61148b831519ca1 string) (string, error) {
	__obf_0c15a3a65dfe2b86 := reflect.ValueOf(__obf_e6ccc30272d0143b)
	if __obf_0c15a3a65dfe2b86.Kind() == reflect.Pointer {
		__obf_0c15a3a65dfe2b86 = __obf_0c15a3a65dfe2b86.Elem()
	}
	var (
		__obf_1696e8696422ed34 = __obf_0c15a3a65dfe2b86.
					FieldByName("Id")
		__obf_84df7616189ea291 = util.ToSnake(reflect.Indirect(__obf_0c15a3a65dfe2b86).Type().Name())
		__obf_63a6a7575ede7589 string
		__obf_d51eb37f5f1cb276 map[string]any
		__obf_87239f12e0f48308 error
	)

	if __obf_1696e8696422ed34.IsZero() {
		__obf_63a6a7575ede7589,
			// insert
			__obf_d51eb37f5f1cb276, __obf_87239f12e0f48308 = BuildInsertSQL(__obf_e6ccc30272d0143b, __obf_84df7616189ea291)
		if __obf_87239f12e0f48308 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_87239f12e0f48308.Error())
		}
		if __obf_e61148b831519ca1 == "" {
			__obf_e61148b831519ca1 = util.StringUUID()
		}
		__obf_d51eb37f5f1cb276["id"] = __obf_e61148b831519ca1
		__obf_fc5b65fde87caf19 := time.Now().Unix()
		if _, __obf_bfedcf8573204366 := __obf_d51eb37f5f1cb276[__obf_019b2ae2e7d131c0.Option.Reserved.CreateTime]; __obf_bfedcf8573204366 {
			__obf_d51eb37f5f1cb276[__obf_019b2ae2e7d131c0.Option.Reserved.CreateTime] = __obf_fc5b65fde87caf19
		}
		if _, __obf_bfedcf8573204366 := __obf_d51eb37f5f1cb276[__obf_019b2ae2e7d131c0.Option.Reserved.UpdateTime]; __obf_bfedcf8573204366 {
			__obf_d51eb37f5f1cb276[__obf_019b2ae2e7d131c0.Option.Reserved.UpdateTime] = __obf_fc5b65fde87caf19
		}
	} else {
		__obf_63a6a7575ede7589,
			// update
			__obf_d51eb37f5f1cb276, __obf_87239f12e0f48308 = BuildUpdateSQL(__obf_e6ccc30272d0143b, __obf_84df7616189ea291, __obf_019b2ae2e7d131c0.Option.Reserved.UpdateTime, []string{__obf_019b2ae2e7d131c0.Option.Reserved.ID}, __obf_019b2ae2e7d131c0.Option.Reserved.CreateTime, __obf_019b2ae2e7d131c0.Option.Reserved.NO)
		if __obf_87239f12e0f48308 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_87239f12e0f48308.Error())
		}
	}

	if __obf_63a6a7575ede7589 != "" {
		_, __obf_87239f12e0f48308 := __obf_019b2ae2e7d131c0.DB.NamedExec(__obf_63a6a7575ede7589, __obf_d51eb37f5f1cb276)
		if __obf_87239f12e0f48308 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_87239f12e0f48308.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_63a6a7575ede7589, "args", __obf_d51eb37f5f1cb276)
	}

	return __obf_e61148b831519ca1, nil

}

func (__obf_019b2ae2e7d131c0 *ORM) SaveModelWidthAutoID(__obf_e6ccc30272d0143b any, __obf_e61148b831519ca1 int64) (int64, error) {
	__obf_0c15a3a65dfe2b86 := // query, args := modelBind(db.Option, m, insertId)
		reflect.ValueOf(__obf_e6ccc30272d0143b)
	if __obf_0c15a3a65dfe2b86.Kind() == reflect.Pointer {
		__obf_0c15a3a65dfe2b86 = __obf_0c15a3a65dfe2b86.Elem()
	}
	__obf_1696e8696422ed34 := __obf_0c15a3a65dfe2b86.FieldByName("Id")
	__obf_84df7616189ea291 := util.ToSnake(reflect.Indirect(__obf_0c15a3a65dfe2b86).Type().Name())
	var (
		__obf_63a6a7575ede7589 string
		__obf_d51eb37f5f1cb276 map[string]any
		__obf_87239f12e0f48308 error
	)

	if __obf_1696e8696422ed34.IsZero() {
		__obf_63a6a7575ede7589,
			// insert
			__obf_d51eb37f5f1cb276, __obf_87239f12e0f48308 = BuildInsertSQL(__obf_e6ccc30272d0143b, __obf_84df7616189ea291)
		if __obf_87239f12e0f48308 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_87239f12e0f48308.Error())
		}
		if __obf_e61148b831519ca1 != 0 {
			__obf_d51eb37f5f1cb276["id"] = __obf_e61148b831519ca1
		}
		__obf_fc5b65fde87caf19 := time.Now().Unix()
		if _, __obf_bfedcf8573204366 := __obf_d51eb37f5f1cb276[__obf_019b2ae2e7d131c0.Option.Reserved.CreateTime]; __obf_bfedcf8573204366 {
			__obf_d51eb37f5f1cb276[__obf_019b2ae2e7d131c0.Option.Reserved.CreateTime] = __obf_fc5b65fde87caf19
		}
		if _, __obf_bfedcf8573204366 := __obf_d51eb37f5f1cb276[__obf_019b2ae2e7d131c0.Option.Reserved.UpdateTime]; __obf_bfedcf8573204366 {
			__obf_d51eb37f5f1cb276[__obf_019b2ae2e7d131c0.Option.Reserved.UpdateTime] = __obf_fc5b65fde87caf19
		}
	} else {
		__obf_63a6a7575ede7589,
			// update
			__obf_d51eb37f5f1cb276, __obf_87239f12e0f48308 = BuildUpdateSQL(__obf_e6ccc30272d0143b, __obf_84df7616189ea291, __obf_019b2ae2e7d131c0.Option.Reserved.UpdateTime, []string{__obf_019b2ae2e7d131c0.Option.Reserved.ID}, __obf_019b2ae2e7d131c0.Option.Reserved.CreateTime, __obf_019b2ae2e7d131c0.Option.Reserved.NO)
		if __obf_87239f12e0f48308 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_87239f12e0f48308.Error())
		}
	}

	if __obf_63a6a7575ede7589 != "" {
		__obf_8ed96d1577fd8cf6, __obf_87239f12e0f48308 := __obf_019b2ae2e7d131c0.DB.NamedExec(__obf_63a6a7575ede7589, __obf_d51eb37f5f1cb276)
		if __obf_87239f12e0f48308 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_87239f12e0f48308.Error())
		}
		if __obf_e61148b831519ca1 == 0 {
			return __obf_8ed96d1577fd8cf6.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_63a6a7575ede7589, "args", __obf_d51eb37f5f1cb276)
	}

	return __obf_e61148b831519ca1, nil
}

func (__obf_019b2ae2e7d131c0 *ORM) SaveTxModel(__obf_945c2e80cb402175 *sqlx.Tx, __obf_e6ccc30272d0143b any, __obf_e61148b831519ca1 string) error {
	var __obf_63a6a7575ede7589 string
	var __obf_d51eb37f5f1cb276 map[string]any
	__obf_63a6a7575ede7589, __obf_d51eb37f5f1cb276 = __obf_3621642bce364936(__obf_019b2ae2e7d131c0.Option, __obf_e6ccc30272d0143b, __obf_e61148b831519ca1)

	if __obf_63a6a7575ede7589 != "" {
		if _, __obf_87239f12e0f48308 := __obf_945c2e80cb402175.NamedExec(__obf_63a6a7575ede7589, __obf_d51eb37f5f1cb276); __obf_87239f12e0f48308 != nil {
			_ = __obf_945c2e80cb402175.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_87239f12e0f48308)
		}

		slog.Debug("SaveTxModel", "query", __obf_63a6a7575ede7589, "args", __obf_d51eb37f5f1cb276)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_3621642bce364936[RecordID int | string](__obf_49a8263c0b23a89d *Options, __obf_e6ccc30272d0143b any, __obf_e61148b831519ca1 RecordID) (__obf_f2daa417656eadde string, __obf_d51eb37f5f1cb276 map[string]any) {
	__obf_0c15a3a65dfe2b86 := reflect.ValueOf(__obf_e6ccc30272d0143b)
	if __obf_0c15a3a65dfe2b86.Kind() == reflect.Pointer {
		__obf_0c15a3a65dfe2b86 = __obf_0c15a3a65dfe2b86.Elem()
	}
	var __obf_040be2e4fc761967 []string
	__obf_d51eb37f5f1cb276 = make(map[string]any)
	__obf_1696e8696422ed34 := __obf_0c15a3a65dfe2b86.FieldByName("Id")
	if __obf_1696e8696422ed34.IsZero() {
		// 判断inserId是否为零值
		if __obf_e61148b831519ca1 != *new(RecordID) {
			__obf_040be2e4fc761967 = append(__obf_040be2e4fc761967, __obf_49a8263c0b23a89d.Reserved.ID)
			__obf_d51eb37f5f1cb276[__obf_49a8263c0b23a89d.Reserved.ID] = __obf_e61148b831519ca1
		} else {
			if __obf_1696e8696422ed34.Kind() == reflect.String {
				__obf_040be2e4fc761967 = append(__obf_040be2e4fc761967, __obf_49a8263c0b23a89d.Reserved.ID)
				__obf_d51eb37f5f1cb276[__obf_49a8263c0b23a89d.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}
		__obf_e9c7b039c67e532b := reflect.Indirect(__obf_0c15a3a65dfe2b86).Type()
		for __obf_33fdeff971fca5a6 := range __obf_e9c7b039c67e532b.NumField() {
			if __obf_ebada1fffc8c92ff := __obf_e9c7b039c67e532b.Field(__obf_33fdeff971fca5a6); __obf_ebada1fffc8c92ff.Tag.Get("db") != "-" {
				switch __obf_d8f4c85893097807 := util.ToSnake(__obf_ebada1fffc8c92ff.Name); __obf_d8f4c85893097807 {
				case __obf_49a8263c0b23a89d.Reserved.CreateTime, __obf_49a8263c0b23a89d.Reserved.UpdateTime:
					__obf_040be2e4fc761967 = append(__obf_040be2e4fc761967, __obf_d8f4c85893097807)
					__obf_d51eb37f5f1cb276[__obf_d8f4c85893097807] = time.Now().Unix()
				default:
					__obf_935e5179094b0bc8 := __obf_0c15a3a65dfe2b86.Field(__obf_33fdeff971fca5a6).Interface()
					if strings.HasPrefix(__obf_d8f4c85893097807, "is_") || __obf_935e5179094b0bc8 != reflect.Zero(__obf_ebada1fffc8c92ff.Type).Interface() {
						__obf_040be2e4fc761967 = append(__obf_040be2e4fc761967, __obf_d8f4c85893097807)
						__obf_d51eb37f5f1cb276[__obf_d8f4c85893097807] = __obf_935e5179094b0bc8
					}
				}
			}
		}
		__obf_f2daa417656eadde = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_e9c7b039c67e532b.Name()), strings.Join(__obf_040be2e4fc761967, ","), strings.Join(__obf_040be2e4fc761967, ",:"))
	} else {
		__obf_e9c7b039c67e532b := reflect.Indirect(__obf_0c15a3a65dfe2b86).Type()
		for __obf_33fdeff971fca5a6 := range __obf_e9c7b039c67e532b.NumField() {
			if __obf_ebada1fffc8c92ff := __obf_e9c7b039c67e532b.Field(__obf_33fdeff971fca5a6); __obf_ebada1fffc8c92ff.Tag.Get("db") != "-" {
				__obf_935e5179094b0bc8 := __obf_0c15a3a65dfe2b86.Field(__obf_33fdeff971fca5a6).Interface()
				if __obf_d8f4c85893097807 := util.ToSnake(__obf_ebada1fffc8c92ff.Name); strings.HasPrefix(__obf_d8f4c85893097807, "is_") || __obf_935e5179094b0bc8 != reflect.Zero(__obf_ebada1fffc8c92ff.Type).Interface() {
					__obf_040be2e4fc761967 = append(__obf_040be2e4fc761967, fmt.Sprintf("%s=:%s", __obf_d8f4c85893097807, __obf_d8f4c85893097807))
					__obf_d51eb37f5f1cb276[__obf_d8f4c85893097807] = __obf_935e5179094b0bc8
				}
			}
		}

		if util.HasField(__obf_0c15a3a65dfe2b86, util.ToCamel(__obf_49a8263c0b23a89d.Reserved.UpdateTime)) {
			__obf_385dffa8492ed9aa := __obf_d51eb37f5f1cb276[__obf_49a8263c0b23a89d.Reserved.UpdateTime]
			__obf_d51eb37f5f1cb276[__obf_49a8263c0b23a89d.Reserved.UpdateTime] = time.Now().Unix()
			__obf_f2daa417656eadde = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_e9c7b039c67e532b.Name()), strings.Join(__obf_040be2e4fc761967, ","), __obf_49a8263c0b23a89d.Reserved.ID, __obf_49a8263c0b23a89d.Reserved.ID, __obf_49a8263c0b23a89d.Reserved.UpdateTime, __obf_385dffa8492ed9aa)
		} else {
			__obf_f2daa417656eadde = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_e9c7b039c67e532b.Name()), strings.Join(__obf_040be2e4fc761967, ","), __obf_49a8263c0b23a89d.Reserved.ID, __obf_49a8263c0b23a89d.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_60e2f85c9b390c47 any, __obf_84df7616189ea291 string) (string, map[string]any, error) {
	if __obf_84df7616189ea291 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_d51eb37f5f1cb276 := make(map[string]any)
	__obf_040be2e4fc761967 := []string{}
	__obf_a2a9526ae5cbbb9d := func(__obf_89f6e57c8628366a string, __obf_d131126320e759f6 any, __obf_6726e9ede7f4ba3c bool) {
		__obf_9cc396279448fd9c := __obf_d131126320e759f6 == nil || reflect.DeepEqual(__obf_d131126320e759f6, reflect.Zero(reflect.TypeOf(__obf_d131126320e759f6)).Interface())
		if __obf_6726e9ede7f4ba3c && __obf_9cc396279448fd9c {
			return
		}
		__obf_040be2e4fc761967 = append(__obf_040be2e4fc761967, __obf_89f6e57c8628366a)
		__obf_d51eb37f5f1cb276[__obf_89f6e57c8628366a] = __obf_d131126320e759f6
	}

	switch __obf_935e5179094b0bc8 := __obf_60e2f85c9b390c47.(type) {
	case map[string]any:
		for __obf_299b79baf6e9ee26, __obf_0c15a3a65dfe2b86 := range __obf_935e5179094b0bc8 {
			__obf_a2a9526ae5cbbb9d(__obf_299b79baf6e9ee26, __obf_0c15a3a65dfe2b86, false)
		}

	default:
		__obf_374fb3284e6cd2d2 := reflect.TypeOf(__obf_60e2f85c9b390c47)
		if __obf_374fb3284e6cd2d2.Kind() == reflect.Pointer {
			__obf_374fb3284e6cd2d2 = __obf_374fb3284e6cd2d2.Elem()
		}
		__obf_0c15a3a65dfe2b86 := reflect.ValueOf(__obf_60e2f85c9b390c47)
		if __obf_0c15a3a65dfe2b86.Kind() == reflect.Pointer {
			__obf_0c15a3a65dfe2b86 = __obf_0c15a3a65dfe2b86.Elem()
		}

		if __obf_374fb3284e6cd2d2.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_33fdeff971fca5a6 := 0; __obf_33fdeff971fca5a6 < __obf_374fb3284e6cd2d2.NumField(); __obf_33fdeff971fca5a6++ {
			__obf_35b8be3340c5326a := __obf_374fb3284e6cd2d2.Field(__obf_33fdeff971fca5a6).Tag.Get("db")
			if __obf_35b8be3340c5326a == "" || __obf_35b8be3340c5326a == "-" {
				continue
			}
			__obf_19ffdc7506b410d5 := strings.Split(__obf_35b8be3340c5326a, ",")
			__obf_a2a9526ae5cbbb9d(__obf_19ffdc7506b410d5[0], __obf_0c15a3a65dfe2b86.Field(__obf_33fdeff971fca5a6).Interface(), len(__obf_19ffdc7506b410d5) > 1 && __obf_19ffdc7506b410d5[1] == "omitempty")
		}
	}

	if len(__obf_040be2e4fc761967) == 0 {
		return "", nil, errors.New("no fields to insert")
	}
	__obf_63a6a7575ede7589 := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", __obf_84df7616189ea291, strings.Join(__obf_040be2e4fc761967, ","),
		strings.Join(__obf_040be2e4fc761967, ",:"),
	)

	return __obf_63a6a7575ede7589,

		// pkField 主键字段
		// lockField 乐观锁字段(默认为updated_at)
		// constField 固定不变字段
		__obf_d51eb37f5f1cb276, nil
}

func BuildUpdateSQL(__obf_60e2f85c9b390c47 any, __obf_84df7616189ea291, __obf_27ac2a047202375e string, __obf_ae21c87ed6c7798a []string, __obf_cccb9a1f7321ba10 ...string,
) (string, map[string]any, error) {
	if __obf_84df7616189ea291 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_1caf7d72c7e57bec := make(map[string]struct{})
	for _, __obf_cc75e9f1b9bf03b9 := range __obf_ae21c87ed6c7798a {
		__obf_1caf7d72c7e57bec[__obf_cc75e9f1b9bf03b9] = struct{}{}
	}
	__obf_d51eb37f5f1cb276 := make(map[string]any)
	__obf_040be2e4fc761967 := []string{}
	__obf_1a6610f17bbff768 := []string{}
	__obf_a2a9526ae5cbbb9d := func(__obf_89f6e57c8628366a string, __obf_d131126320e759f6 any, __obf_6726e9ede7f4ba3c bool) {
		__obf_9cc396279448fd9c := __obf_d131126320e759f6 == nil || reflect.DeepEqual(__obf_d131126320e759f6, reflect.Zero(reflect.TypeOf(__obf_d131126320e759f6)).Interface())

		switch {
		case func() bool {
			_, __obf_bfedcf8573204366 := __obf_1caf7d72c7e57bec[__obf_89f6e57c8628366a]
			return __obf_bfedcf8573204366
		}():
			__obf_1a6610f17bbff768 = append(__obf_1a6610f17bbff768, fmt.Sprintf("%s=:%s", __obf_89f6e57c8628366a, __obf_89f6e57c8628366a))
			__obf_d51eb37f5f1cb276[__obf_89f6e57c8628366a] = __obf_d131126320e759f6

		case slices.Contains(__obf_cccb9a1f7321ba10, __obf_89f6e57c8628366a):
			// skip constField
		case __obf_89f6e57c8628366a == __obf_27ac2a047202375e:
			const __obf_1cdb1d48fb73d9ff = "RESERVED_LOCK"
			__obf_040be2e4fc761967 = append(__obf_040be2e4fc761967, fmt.Sprintf("%s=:%s", __obf_89f6e57c8628366a, __obf_89f6e57c8628366a))
			__obf_1a6610f17bbff768 = append(__obf_1a6610f17bbff768, fmt.Sprintf("%s=:%s", __obf_89f6e57c8628366a, __obf_1cdb1d48fb73d9ff))
			__obf_d51eb37f5f1cb276[__obf_89f6e57c8628366a] = time.Now().Unix()
			__obf_d51eb37f5f1cb276[__obf_1cdb1d48fb73d9ff] = __obf_d131126320e759f6

		default:
			if __obf_6726e9ede7f4ba3c && __obf_9cc396279448fd9c {
				return
			}
			__obf_040be2e4fc761967 = append(__obf_040be2e4fc761967, fmt.Sprintf("%s=:%s", __obf_89f6e57c8628366a, __obf_89f6e57c8628366a))
			__obf_d51eb37f5f1cb276[__obf_89f6e57c8628366a] = __obf_d131126320e759f6
		}
	}

	switch __obf_935e5179094b0bc8 := __obf_60e2f85c9b390c47.(type) {
	case map[string]any:
		for __obf_299b79baf6e9ee26, __obf_0c15a3a65dfe2b86 := range __obf_935e5179094b0bc8 {
			__obf_a2a9526ae5cbbb9d(__obf_299b79baf6e9ee26, __obf_0c15a3a65dfe2b86, false)
		}

	default:
		__obf_374fb3284e6cd2d2 := reflect.TypeOf(__obf_60e2f85c9b390c47)
		if __obf_374fb3284e6cd2d2.Kind() == reflect.Pointer {
			__obf_374fb3284e6cd2d2 = __obf_374fb3284e6cd2d2.Elem()
		}
		__obf_0c15a3a65dfe2b86 := reflect.ValueOf(__obf_60e2f85c9b390c47)
		if __obf_0c15a3a65dfe2b86.Kind() == reflect.Pointer {
			__obf_0c15a3a65dfe2b86 = __obf_0c15a3a65dfe2b86.Elem()
		}
		if __obf_374fb3284e6cd2d2.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_33fdeff971fca5a6 := 0; __obf_33fdeff971fca5a6 < __obf_374fb3284e6cd2d2.NumField(); __obf_33fdeff971fca5a6++ {
			__obf_35b8be3340c5326a := __obf_374fb3284e6cd2d2.Field(__obf_33fdeff971fca5a6).Tag.Get("db")
			if __obf_35b8be3340c5326a == "" || __obf_35b8be3340c5326a == "-" {
				continue
			}
			__obf_19ffdc7506b410d5 := strings.Split(__obf_35b8be3340c5326a, ",")
			__obf_a2a9526ae5cbbb9d(__obf_19ffdc7506b410d5[0], __obf_0c15a3a65dfe2b86.Field(__obf_33fdeff971fca5a6).Interface(), len(__obf_19ffdc7506b410d5) > 1 && __obf_19ffdc7506b410d5[1] == "omitempty")
		}
	}

	if len(__obf_ae21c87ed6c7798a) == 0 || len(__obf_1a6610f17bbff768) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_040be2e4fc761967) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	__obf_63a6a7575ede7589 := // if lockField != "" && !hasLockField {
		// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
		// }

		fmt.Sprintf("UPDATE %s SET %s WHERE %s", __obf_84df7616189ea291, strings.Join(__obf_040be2e4fc761967, ","),
			strings.Join(__obf_1a6610f17bbff768, " AND "),
		)

	return __obf_63a6a7575ede7589, __obf_d51eb37f5f1cb276, nil
}
