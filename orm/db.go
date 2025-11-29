package __obf_214d1e062aee9185

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

func NewORM(__obf_a3eb6a9918aa41c6 *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_a3eb6a9918aa41c6.Driver, __obf_a3eb6a9918aa41c6.DataSource)

	switch __obf_a3eb6a9918aa41c6.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_a3eb6a9918aa41c6.
			Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_a3eb6a9918aa41c6.
			Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_a3eb6a9918aa41c6.
			Placeholder = ":"
	case "sqlserver":
		__obf_a3eb6a9918aa41c6.
			Placeholder = "@"
	}

	if __obf_63349b4c89eea887 := DB.Ping(); __obf_63349b4c89eea887 != nil {
		panic(__obf_63349b4c89eea887)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_a3eb6a9918aa41c6,
		}, func() {
			DB.Close()
		}
}

func (__obf_9858352bdc58a6fc *ORM) ListMap(__obf_97c7de7ee94125bc *sqlx.Rows, __obf_e838cc4961464fb4 func(map[string]any) (string, string)) (__obf_1576eeccdf72c56c []map[string]any, __obf_63349b4c89eea887 error) {
	for __obf_97c7de7ee94125bc.Next() {
		__obf_26037241a6982718 := make(map[string]any)
		if __obf_63349b4c89eea887 = __obf_97c7de7ee94125bc.MapScan(__obf_26037241a6982718); __obf_63349b4c89eea887 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_63349b4c89eea887)
		}
		for __obf_9b67b54491941e78, __obf_cf9e8016bdd91675 := range __obf_26037241a6982718 {
			switch __obf_cf9e8016bdd91675 := __obf_cf9e8016bdd91675.(type) {
			case []uint8:
				__obf_26037241a6982718[__obf_9b67b54491941e78] = string(__obf_cf9e8016bdd91675)
			case int64:
				__obf_26037241a6982718[__obf_9b67b54491941e78] = int64(__obf_cf9e8016bdd91675)
			case uint64:
				__obf_26037241a6982718[__obf_9b67b54491941e78] = uint64(__obf_cf9e8016bdd91675)
			}
		}
		if __obf_e838cc4961464fb4 != nil {
			__obf_16fdb1dd68ada711, __obf_c19d7b5d3b7a82bb := __obf_e838cc4961464fb4(__obf_26037241a6982718)
			__obf_26037241a6982718[__obf_16fdb1dd68ada711] = __obf_c19d7b5d3b7a82bb
		}
		__obf_1576eeccdf72c56c = append(__obf_1576eeccdf72c56c, __obf_26037241a6982718)
	}
	return
}

func (__obf_9858352bdc58a6fc *ORM) Exists(__obf_b025d80df9cc3143, __obf_9c0f3cdc51a1c12b string, __obf_ede102380a19106b ...any) bool {
	var __obf_646ad5fc24b0b8a2 bool
	_ = __obf_9858352bdc58a6fc.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_b025d80df9cc3143, __obf_9c0f3cdc51a1c12b), __obf_ede102380a19106b...).Scan(&__obf_646ad5fc24b0b8a2)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_646ad5fc24b0b8a2
}

func (__obf_9858352bdc58a6fc *ORM) SaveModel(__obf_cc69e7efd3b8b168 any, __obf_e4298885b2647a3d string) (string, error) {
	__obf_cf9e8016bdd91675 := reflect.ValueOf(__obf_cc69e7efd3b8b168)
	if __obf_cf9e8016bdd91675.Kind() == reflect.Pointer {
		__obf_cf9e8016bdd91675 = __obf_cf9e8016bdd91675.Elem()
	}
	var (
		__obf_0f7e3dac05988774 = __obf_cf9e8016bdd91675.
					FieldByName("Id")
		__obf_343096a0d9708548 = util.ToSnake(reflect.Indirect(__obf_cf9e8016bdd91675).Type().Name())
		__obf_8350072420eebe14 string
		__obf_ede102380a19106b map[string]any
		__obf_63349b4c89eea887 error
	)

	if __obf_0f7e3dac05988774.IsZero() {
		__obf_8350072420eebe14,
			// insert
			__obf_ede102380a19106b, __obf_63349b4c89eea887 = BuildInsertSQL(__obf_cc69e7efd3b8b168, __obf_343096a0d9708548)
		if __obf_63349b4c89eea887 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_63349b4c89eea887.Error())
		}
		if __obf_e4298885b2647a3d == "" {
			__obf_e4298885b2647a3d = util.StringUUID()
		}
		__obf_ede102380a19106b["id"] = __obf_e4298885b2647a3d
		__obf_3df3ac99436393c8 := time.Now().Unix()
		if _, __obf_65bc54d69fb5f329 := __obf_ede102380a19106b[__obf_9858352bdc58a6fc.Option.Reserved.CreateTime]; __obf_65bc54d69fb5f329 {
			__obf_ede102380a19106b[__obf_9858352bdc58a6fc.Option.Reserved.CreateTime] = __obf_3df3ac99436393c8
		}
		if _, __obf_65bc54d69fb5f329 := __obf_ede102380a19106b[__obf_9858352bdc58a6fc.Option.Reserved.UpdateTime]; __obf_65bc54d69fb5f329 {
			__obf_ede102380a19106b[__obf_9858352bdc58a6fc.Option.Reserved.UpdateTime] = __obf_3df3ac99436393c8
		}
	} else {
		__obf_8350072420eebe14,
			// update
			__obf_ede102380a19106b, __obf_63349b4c89eea887 = BuildUpdateSQL(__obf_cc69e7efd3b8b168, __obf_343096a0d9708548, __obf_9858352bdc58a6fc.Option.Reserved.UpdateTime, []string{__obf_9858352bdc58a6fc.Option.Reserved.ID}, __obf_9858352bdc58a6fc.Option.Reserved.CreateTime, __obf_9858352bdc58a6fc.Option.Reserved.NO)
		if __obf_63349b4c89eea887 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_63349b4c89eea887.Error())
		}
	}

	if __obf_8350072420eebe14 != "" {
		_, __obf_63349b4c89eea887 := __obf_9858352bdc58a6fc.DB.NamedExec(__obf_8350072420eebe14, __obf_ede102380a19106b)
		if __obf_63349b4c89eea887 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_63349b4c89eea887.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_8350072420eebe14, "args", __obf_ede102380a19106b)
	}

	return __obf_e4298885b2647a3d, nil

}

func (__obf_9858352bdc58a6fc *ORM) SaveModelWidthAutoID(__obf_cc69e7efd3b8b168 any, __obf_e4298885b2647a3d int64) (int64, error) {
	__obf_cf9e8016bdd91675 := // query, args := modelBind(db.Option, m, insertId)
		reflect.ValueOf(__obf_cc69e7efd3b8b168)
	if __obf_cf9e8016bdd91675.Kind() == reflect.Pointer {
		__obf_cf9e8016bdd91675 = __obf_cf9e8016bdd91675.Elem()
	}
	__obf_0f7e3dac05988774 := __obf_cf9e8016bdd91675.FieldByName("Id")
	__obf_343096a0d9708548 := util.ToSnake(reflect.Indirect(__obf_cf9e8016bdd91675).Type().Name())
	var (
		__obf_8350072420eebe14 string
		__obf_ede102380a19106b map[string]any
		__obf_63349b4c89eea887 error
	)

	if __obf_0f7e3dac05988774.IsZero() {
		__obf_8350072420eebe14,
			// insert
			__obf_ede102380a19106b, __obf_63349b4c89eea887 = BuildInsertSQL(__obf_cc69e7efd3b8b168, __obf_343096a0d9708548)
		if __obf_63349b4c89eea887 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_63349b4c89eea887.Error())
		}
		if __obf_e4298885b2647a3d != 0 {
			__obf_ede102380a19106b["id"] = __obf_e4298885b2647a3d
		}
		__obf_3df3ac99436393c8 := time.Now().Unix()
		if _, __obf_65bc54d69fb5f329 := __obf_ede102380a19106b[__obf_9858352bdc58a6fc.Option.Reserved.CreateTime]; __obf_65bc54d69fb5f329 {
			__obf_ede102380a19106b[__obf_9858352bdc58a6fc.Option.Reserved.CreateTime] = __obf_3df3ac99436393c8
		}
		if _, __obf_65bc54d69fb5f329 := __obf_ede102380a19106b[__obf_9858352bdc58a6fc.Option.Reserved.UpdateTime]; __obf_65bc54d69fb5f329 {
			__obf_ede102380a19106b[__obf_9858352bdc58a6fc.Option.Reserved.UpdateTime] = __obf_3df3ac99436393c8
		}
	} else {
		__obf_8350072420eebe14,
			// update
			__obf_ede102380a19106b, __obf_63349b4c89eea887 = BuildUpdateSQL(__obf_cc69e7efd3b8b168, __obf_343096a0d9708548, __obf_9858352bdc58a6fc.Option.Reserved.UpdateTime, []string{__obf_9858352bdc58a6fc.Option.Reserved.ID}, __obf_9858352bdc58a6fc.Option.Reserved.CreateTime, __obf_9858352bdc58a6fc.Option.Reserved.NO)
		if __obf_63349b4c89eea887 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_63349b4c89eea887.Error())
		}
	}

	if __obf_8350072420eebe14 != "" {
		__obf_c19d7b5d3b7a82bb, __obf_63349b4c89eea887 := __obf_9858352bdc58a6fc.DB.NamedExec(__obf_8350072420eebe14, __obf_ede102380a19106b)
		if __obf_63349b4c89eea887 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_63349b4c89eea887.Error())
		}
		if __obf_e4298885b2647a3d == 0 {
			return __obf_c19d7b5d3b7a82bb.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_8350072420eebe14, "args", __obf_ede102380a19106b)
	}

	return __obf_e4298885b2647a3d, nil
}

func (__obf_9858352bdc58a6fc *ORM) SaveTxModel(__obf_300af38d91fabd70 *sqlx.Tx, __obf_cc69e7efd3b8b168 any, __obf_e4298885b2647a3d string) error {
	var __obf_8350072420eebe14 string
	var __obf_ede102380a19106b map[string]any
	__obf_8350072420eebe14, __obf_ede102380a19106b = __obf_812aee5c26a337d6(__obf_9858352bdc58a6fc.Option, __obf_cc69e7efd3b8b168, __obf_e4298885b2647a3d)

	if __obf_8350072420eebe14 != "" {
		if _, __obf_63349b4c89eea887 := __obf_300af38d91fabd70.NamedExec(__obf_8350072420eebe14, __obf_ede102380a19106b); __obf_63349b4c89eea887 != nil {
			_ = __obf_300af38d91fabd70.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_63349b4c89eea887)
		}

		slog.Debug("SaveTxModel", "query", __obf_8350072420eebe14, "args", __obf_ede102380a19106b)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_812aee5c26a337d6[RecordID int | string](__obf_a3eb6a9918aa41c6 *Options, __obf_cc69e7efd3b8b168 any, __obf_e4298885b2647a3d RecordID) (__obf_a6e0fc98d079e0db string, __obf_ede102380a19106b map[string]any) {
	__obf_cf9e8016bdd91675 := reflect.ValueOf(__obf_cc69e7efd3b8b168)
	if __obf_cf9e8016bdd91675.Kind() == reflect.Pointer {
		__obf_cf9e8016bdd91675 = __obf_cf9e8016bdd91675.Elem()
	}
	var __obf_d6674f7e971548e5 []string
	__obf_ede102380a19106b = make(map[string]any)
	__obf_0f7e3dac05988774 := __obf_cf9e8016bdd91675.FieldByName("Id")
	if __obf_0f7e3dac05988774.IsZero() {
		// 判断inserId是否为零值
		if __obf_e4298885b2647a3d != *new(RecordID) {
			__obf_d6674f7e971548e5 = append(__obf_d6674f7e971548e5, __obf_a3eb6a9918aa41c6.Reserved.ID)
			__obf_ede102380a19106b[__obf_a3eb6a9918aa41c6.Reserved.ID] = __obf_e4298885b2647a3d
		} else {
			if __obf_0f7e3dac05988774.Kind() == reflect.String {
				__obf_d6674f7e971548e5 = append(__obf_d6674f7e971548e5, __obf_a3eb6a9918aa41c6.Reserved.ID)
				__obf_ede102380a19106b[__obf_a3eb6a9918aa41c6.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}
		__obf_f963d4d02d5b78b3 := reflect.Indirect(__obf_cf9e8016bdd91675).Type()
		for __obf_7c6c9e1512e37985 := range __obf_f963d4d02d5b78b3.NumField() {
			if __obf_e242c9b073a9f4fe := __obf_f963d4d02d5b78b3.Field(__obf_7c6c9e1512e37985); __obf_e242c9b073a9f4fe.Tag.Get("db") != "-" {
				switch __obf_5d3cc48af55ee91d := util.ToSnake(__obf_e242c9b073a9f4fe.Name); __obf_5d3cc48af55ee91d {
				case __obf_a3eb6a9918aa41c6.Reserved.CreateTime, __obf_a3eb6a9918aa41c6.Reserved.UpdateTime:
					__obf_d6674f7e971548e5 = append(__obf_d6674f7e971548e5, __obf_5d3cc48af55ee91d)
					__obf_ede102380a19106b[__obf_5d3cc48af55ee91d] = time.Now().Unix()
				default:
					__obf_a73e9633161d85ba := __obf_cf9e8016bdd91675.Field(__obf_7c6c9e1512e37985).Interface()
					if strings.HasPrefix(__obf_5d3cc48af55ee91d, "is_") || __obf_a73e9633161d85ba != reflect.Zero(__obf_e242c9b073a9f4fe.Type).Interface() {
						__obf_d6674f7e971548e5 = append(__obf_d6674f7e971548e5, __obf_5d3cc48af55ee91d)
						__obf_ede102380a19106b[__obf_5d3cc48af55ee91d] = __obf_a73e9633161d85ba
					}
				}
			}
		}
		__obf_a6e0fc98d079e0db = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_f963d4d02d5b78b3.Name()), strings.Join(__obf_d6674f7e971548e5, ","), strings.Join(__obf_d6674f7e971548e5, ",:"))
	} else {
		__obf_f963d4d02d5b78b3 := reflect.Indirect(__obf_cf9e8016bdd91675).Type()
		for __obf_7c6c9e1512e37985 := range __obf_f963d4d02d5b78b3.NumField() {
			if __obf_e242c9b073a9f4fe := __obf_f963d4d02d5b78b3.Field(__obf_7c6c9e1512e37985); __obf_e242c9b073a9f4fe.Tag.Get("db") != "-" {
				__obf_a73e9633161d85ba := __obf_cf9e8016bdd91675.Field(__obf_7c6c9e1512e37985).Interface()
				if __obf_5d3cc48af55ee91d := util.ToSnake(__obf_e242c9b073a9f4fe.Name); strings.HasPrefix(__obf_5d3cc48af55ee91d, "is_") || __obf_a73e9633161d85ba != reflect.Zero(__obf_e242c9b073a9f4fe.Type).Interface() {
					__obf_d6674f7e971548e5 = append(__obf_d6674f7e971548e5, fmt.Sprintf("%s=:%s", __obf_5d3cc48af55ee91d, __obf_5d3cc48af55ee91d))
					__obf_ede102380a19106b[__obf_5d3cc48af55ee91d] = __obf_a73e9633161d85ba
				}
			}
		}

		if util.HasField(__obf_cf9e8016bdd91675, util.ToCamel(__obf_a3eb6a9918aa41c6.Reserved.UpdateTime)) {
			__obf_1c2d6be5ab7dfe99 := __obf_ede102380a19106b[__obf_a3eb6a9918aa41c6.Reserved.UpdateTime]
			__obf_ede102380a19106b[__obf_a3eb6a9918aa41c6.Reserved.UpdateTime] = time.Now().Unix()
			__obf_a6e0fc98d079e0db = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_f963d4d02d5b78b3.Name()), strings.Join(__obf_d6674f7e971548e5, ","), __obf_a3eb6a9918aa41c6.Reserved.ID, __obf_a3eb6a9918aa41c6.Reserved.ID, __obf_a3eb6a9918aa41c6.Reserved.UpdateTime, __obf_1c2d6be5ab7dfe99)
		} else {
			__obf_a6e0fc98d079e0db = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_f963d4d02d5b78b3.Name()), strings.Join(__obf_d6674f7e971548e5, ","), __obf_a3eb6a9918aa41c6.Reserved.ID, __obf_a3eb6a9918aa41c6.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_e3cb62ac43162778 any, __obf_343096a0d9708548 string) (string, map[string]any, error) {
	if __obf_343096a0d9708548 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_ede102380a19106b := make(map[string]any)
	__obf_d6674f7e971548e5 := []string{}
	__obf_dd169095c0809faa := func(__obf_96d2ca3394c04fbe string, __obf_5dbd5d199912b6ea any, __obf_b8112b7ff7db28b0 bool) {
		__obf_970ea51bcd3f1683 := __obf_5dbd5d199912b6ea == nil || reflect.DeepEqual(__obf_5dbd5d199912b6ea, reflect.Zero(reflect.TypeOf(__obf_5dbd5d199912b6ea)).Interface())
		if __obf_b8112b7ff7db28b0 && __obf_970ea51bcd3f1683 {
			return
		}
		__obf_d6674f7e971548e5 = append(__obf_d6674f7e971548e5, __obf_96d2ca3394c04fbe)
		__obf_ede102380a19106b[__obf_96d2ca3394c04fbe] = __obf_5dbd5d199912b6ea
	}

	switch __obf_a73e9633161d85ba := __obf_e3cb62ac43162778.(type) {
	case map[string]any:
		for __obf_9b67b54491941e78, __obf_cf9e8016bdd91675 := range __obf_a73e9633161d85ba {
			__obf_dd169095c0809faa(__obf_9b67b54491941e78, __obf_cf9e8016bdd91675, false)
		}

	default:
		__obf_e7e84afe571f9ceb := reflect.TypeOf(__obf_e3cb62ac43162778)
		if __obf_e7e84afe571f9ceb.Kind() == reflect.Pointer {
			__obf_e7e84afe571f9ceb = __obf_e7e84afe571f9ceb.Elem()
		}
		__obf_cf9e8016bdd91675 := reflect.ValueOf(__obf_e3cb62ac43162778)
		if __obf_cf9e8016bdd91675.Kind() == reflect.Pointer {
			__obf_cf9e8016bdd91675 = __obf_cf9e8016bdd91675.Elem()
		}

		if __obf_e7e84afe571f9ceb.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_7c6c9e1512e37985 := 0; __obf_7c6c9e1512e37985 < __obf_e7e84afe571f9ceb.NumField(); __obf_7c6c9e1512e37985++ {
			__obf_4e0a6144ac5e256f := __obf_e7e84afe571f9ceb.Field(__obf_7c6c9e1512e37985).Tag.Get("db")
			if __obf_4e0a6144ac5e256f == "" || __obf_4e0a6144ac5e256f == "-" {
				continue
			}
			__obf_2e882d19f83a48fc := strings.Split(__obf_4e0a6144ac5e256f, ",")
			__obf_dd169095c0809faa(__obf_2e882d19f83a48fc[0], __obf_cf9e8016bdd91675.Field(__obf_7c6c9e1512e37985).Interface(), len(__obf_2e882d19f83a48fc) > 1 && __obf_2e882d19f83a48fc[1] == "omitempty")
		}
	}

	if len(__obf_d6674f7e971548e5) == 0 {
		return "", nil, errors.New("no fields to insert")
	}
	__obf_8350072420eebe14 := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", __obf_343096a0d9708548, strings.Join(__obf_d6674f7e971548e5, ","),
		strings.Join(__obf_d6674f7e971548e5, ",:"),
	)

	return __obf_8350072420eebe14,

		// pkField 主键字段
		// lockField 乐观锁字段(默认为updated_at)
		// constField 固定不变字段
		__obf_ede102380a19106b, nil
}

func BuildUpdateSQL(__obf_e3cb62ac43162778 any, __obf_343096a0d9708548, __obf_41bfad66b2c23ecc string, __obf_267a255f1b2e6f8c []string, __obf_8dbab9533b2596eb ...string,
) (string, map[string]any, error) {
	if __obf_343096a0d9708548 == "" {
		return "", nil, errors.New("table name is required")
	}
	__obf_9c1ee22bfa5de806 := make(map[string]struct{})
	for _, __obf_1f9a3ad791c5bcd7 := range __obf_267a255f1b2e6f8c {
		__obf_9c1ee22bfa5de806[__obf_1f9a3ad791c5bcd7] = struct{}{}
	}
	__obf_ede102380a19106b := make(map[string]any)
	__obf_d6674f7e971548e5 := []string{}
	__obf_849f045b6c32c7f3 := []string{}
	__obf_dd169095c0809faa := func(__obf_96d2ca3394c04fbe string, __obf_5dbd5d199912b6ea any, __obf_b8112b7ff7db28b0 bool) {
		__obf_970ea51bcd3f1683 := __obf_5dbd5d199912b6ea == nil || reflect.DeepEqual(__obf_5dbd5d199912b6ea, reflect.Zero(reflect.TypeOf(__obf_5dbd5d199912b6ea)).Interface())

		switch {
		case func() bool {
			_, __obf_65bc54d69fb5f329 := __obf_9c1ee22bfa5de806[__obf_96d2ca3394c04fbe]
			return __obf_65bc54d69fb5f329
		}():
			__obf_849f045b6c32c7f3 = append(__obf_849f045b6c32c7f3, fmt.Sprintf("%s=:%s", __obf_96d2ca3394c04fbe, __obf_96d2ca3394c04fbe))
			__obf_ede102380a19106b[__obf_96d2ca3394c04fbe] = __obf_5dbd5d199912b6ea

		case slices.Contains(__obf_8dbab9533b2596eb, __obf_96d2ca3394c04fbe):
			// skip constField
		case __obf_96d2ca3394c04fbe == __obf_41bfad66b2c23ecc:
			const __obf_baa6aa082ab8b296 = "RESERVED_LOCK"
			__obf_d6674f7e971548e5 = append(__obf_d6674f7e971548e5, fmt.Sprintf("%s=:%s", __obf_96d2ca3394c04fbe, __obf_96d2ca3394c04fbe))
			__obf_849f045b6c32c7f3 = append(__obf_849f045b6c32c7f3, fmt.Sprintf("%s=:%s", __obf_96d2ca3394c04fbe, __obf_baa6aa082ab8b296))
			__obf_ede102380a19106b[__obf_96d2ca3394c04fbe] = time.Now().Unix()
			__obf_ede102380a19106b[__obf_baa6aa082ab8b296] = __obf_5dbd5d199912b6ea

		default:
			if __obf_b8112b7ff7db28b0 && __obf_970ea51bcd3f1683 {
				return
			}
			__obf_d6674f7e971548e5 = append(__obf_d6674f7e971548e5, fmt.Sprintf("%s=:%s", __obf_96d2ca3394c04fbe, __obf_96d2ca3394c04fbe))
			__obf_ede102380a19106b[__obf_96d2ca3394c04fbe] = __obf_5dbd5d199912b6ea
		}
	}

	switch __obf_a73e9633161d85ba := __obf_e3cb62ac43162778.(type) {
	case map[string]any:
		for __obf_9b67b54491941e78, __obf_cf9e8016bdd91675 := range __obf_a73e9633161d85ba {
			__obf_dd169095c0809faa(__obf_9b67b54491941e78, __obf_cf9e8016bdd91675, false)
		}

	default:
		__obf_e7e84afe571f9ceb := reflect.TypeOf(__obf_e3cb62ac43162778)
		if __obf_e7e84afe571f9ceb.Kind() == reflect.Pointer {
			__obf_e7e84afe571f9ceb = __obf_e7e84afe571f9ceb.Elem()
		}
		__obf_cf9e8016bdd91675 := reflect.ValueOf(__obf_e3cb62ac43162778)
		if __obf_cf9e8016bdd91675.Kind() == reflect.Pointer {
			__obf_cf9e8016bdd91675 = __obf_cf9e8016bdd91675.Elem()
		}
		if __obf_e7e84afe571f9ceb.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_7c6c9e1512e37985 := 0; __obf_7c6c9e1512e37985 < __obf_e7e84afe571f9ceb.NumField(); __obf_7c6c9e1512e37985++ {
			__obf_4e0a6144ac5e256f := __obf_e7e84afe571f9ceb.Field(__obf_7c6c9e1512e37985).Tag.Get("db")
			if __obf_4e0a6144ac5e256f == "" || __obf_4e0a6144ac5e256f == "-" {
				continue
			}
			__obf_2e882d19f83a48fc := strings.Split(__obf_4e0a6144ac5e256f, ",")
			__obf_dd169095c0809faa(__obf_2e882d19f83a48fc[0], __obf_cf9e8016bdd91675.Field(__obf_7c6c9e1512e37985).Interface(), len(__obf_2e882d19f83a48fc) > 1 && __obf_2e882d19f83a48fc[1] == "omitempty")
		}
	}

	if len(__obf_267a255f1b2e6f8c) == 0 || len(__obf_849f045b6c32c7f3) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_d6674f7e971548e5) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	__obf_8350072420eebe14 := // if lockField != "" && !hasLockField {
		// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
		// }

		fmt.Sprintf("UPDATE %s SET %s WHERE %s", __obf_343096a0d9708548, strings.Join(__obf_d6674f7e971548e5, ","),
			strings.Join(__obf_849f045b6c32c7f3, " AND "),
		)

	return __obf_8350072420eebe14, __obf_ede102380a19106b, nil
}
