package __obf_52bf13aa40478808

import (
	"errors"
	"fmt"
	"log/slog"
	"reflect"
	"slices"
	"strings"
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

func NewORM(__obf_2d84a01fd78b7ed8 *Options) (*ORM, func()) {
	DB := sqlx.MustConnect(__obf_2d84a01fd78b7ed8.Driver, __obf_2d84a01fd78b7ed8.DataSource)

	switch __obf_2d84a01fd78b7ed8.Driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql":
		__obf_2d84a01fd78b7ed8.Placeholder = "$"
	case "mysql", "sqlite3":
		__obf_2d84a01fd78b7ed8.Placeholder = "?"
	case "oci8", "ora", "goracle", "godror":
		__obf_2d84a01fd78b7ed8.Placeholder = ":"
	case "sqlserver":
		__obf_2d84a01fd78b7ed8.Placeholder = "@"
	}

	if __obf_e6c24855347eb3d0 := DB.Ping(); __obf_e6c24855347eb3d0 != nil {
		panic(__obf_e6c24855347eb3d0)
	}

	return &ORM{
			DB:     DB,
			Option: __obf_2d84a01fd78b7ed8,
		}, func() {
			DB.Close()
		}
}

func (__obf_86bb12f77b537935 *ORM) ListMap(__obf_85accf2129aa6cd7 *sqlx.Rows, __obf_a2982f1b4e442783 func(map[string]any) (string, string)) (__obf_c560cfbfddbf86d2 []map[string]any, __obf_e6c24855347eb3d0 error) {
	for __obf_85accf2129aa6cd7.Next() {
		__obf_f68079606d680005 := make(map[string]any)
		if __obf_e6c24855347eb3d0 = __obf_85accf2129aa6cd7.MapScan(__obf_f68079606d680005); __obf_e6c24855347eb3d0 != nil {
			return nil, fmt.Errorf("ListMap: %s", __obf_e6c24855347eb3d0)
		}
		for __obf_13037cdadbf8d35e, __obf_306dddc11fc3d690 := range __obf_f68079606d680005 {
			switch __obf_306dddc11fc3d690 := __obf_306dddc11fc3d690.(type) {
			case []uint8:
				__obf_f68079606d680005[__obf_13037cdadbf8d35e] = string(__obf_306dddc11fc3d690)
			case int64:
				__obf_f68079606d680005[__obf_13037cdadbf8d35e] = int64(__obf_306dddc11fc3d690)
			case uint64:
				__obf_f68079606d680005[__obf_13037cdadbf8d35e] = uint64(__obf_306dddc11fc3d690)
			}
		}
		if __obf_a2982f1b4e442783 != nil {
			__obf_9e0957fc4409fee3, __obf_6ab9f0211d0f4779 := __obf_a2982f1b4e442783(__obf_f68079606d680005)
			__obf_f68079606d680005[__obf_9e0957fc4409fee3] = __obf_6ab9f0211d0f4779
		}
		__obf_c560cfbfddbf86d2 = append(__obf_c560cfbfddbf86d2, __obf_f68079606d680005)
	}
	return
}

func (__obf_86bb12f77b537935 *ORM) Exists(__obf_348159e5665379e3, __obf_8062258fcb6f5cc5 string, __obf_cd0ddd46582953ac ...any) bool {
	var __obf_acc0cdc7582f3511 bool
	_ = __obf_86bb12f77b537935.DB.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s LIMIT 1)", __obf_348159e5665379e3, __obf_8062258fcb6f5cc5), __obf_cd0ddd46582953ac...).Scan(&__obf_acc0cdc7582f3511)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	return false
	// }
	return __obf_acc0cdc7582f3511
}

func (__obf_86bb12f77b537935 *ORM) SaveModel(__obf_f2992f05d328343c any, __obf_3d258fae35428ba8 string) (string, error) {
	__obf_306dddc11fc3d690 := reflect.ValueOf(__obf_f2992f05d328343c)
	if __obf_306dddc11fc3d690.Kind() == reflect.Pointer {
		__obf_306dddc11fc3d690 = __obf_306dddc11fc3d690.Elem()
	}
	var (
		__obf_d24f398beafeb263 = __obf_306dddc11fc3d690.FieldByName("Id")
		__obf_82ed77b51afe2acc = util.ToSnake(reflect.Indirect(__obf_306dddc11fc3d690).Type().Name())
		__obf_a856c5ef2c6a7852 string
		__obf_cd0ddd46582953ac map[string]any
		__obf_e6c24855347eb3d0 error
	)

	if __obf_d24f398beafeb263.IsZero() {
		// insert
		__obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac, __obf_e6c24855347eb3d0 = BuildInsertSQL(__obf_f2992f05d328343c, __obf_82ed77b51afe2acc)
		if __obf_e6c24855347eb3d0 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build insert sql: %s", __obf_e6c24855347eb3d0.Error())
		}
		if __obf_3d258fae35428ba8 == "" {
			__obf_3d258fae35428ba8 = util.StringUUID()
		}

		__obf_cd0ddd46582953ac["id"] = __obf_3d258fae35428ba8
		__obf_10f23379b4bda68b := time.Now().Unix()
		if _, __obf_b5d20856a702a4bb := __obf_cd0ddd46582953ac[__obf_86bb12f77b537935.Option.Reserved.CreateTime]; __obf_b5d20856a702a4bb {
			__obf_cd0ddd46582953ac[__obf_86bb12f77b537935.Option.Reserved.CreateTime] = __obf_10f23379b4bda68b
		}
		if _, __obf_b5d20856a702a4bb := __obf_cd0ddd46582953ac[__obf_86bb12f77b537935.Option.Reserved.UpdateTime]; __obf_b5d20856a702a4bb {
			__obf_cd0ddd46582953ac[__obf_86bb12f77b537935.Option.Reserved.UpdateTime] = __obf_10f23379b4bda68b
		}
	} else {
		// update
		__obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac, __obf_e6c24855347eb3d0 = BuildUpdateSQL(__obf_f2992f05d328343c, __obf_82ed77b51afe2acc, __obf_86bb12f77b537935.Option.Reserved.UpdateTime, []string{__obf_86bb12f77b537935.Option.Reserved.ID}, __obf_86bb12f77b537935.Option.Reserved.CreateTime, __obf_86bb12f77b537935.Option.Reserved.NO)
		if __obf_e6c24855347eb3d0 != nil {
			return "", fmt.Errorf("[SaveModel] failed to build update sql: %s", __obf_e6c24855347eb3d0.Error())
		}
	}

	if __obf_a856c5ef2c6a7852 != "" {
		_, __obf_e6c24855347eb3d0 := __obf_86bb12f77b537935.DB.NamedExec(__obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac)
		if __obf_e6c24855347eb3d0 != nil {
			return "", fmt.Errorf("[SaveModel]: %s", __obf_e6c24855347eb3d0.Error())
		}
		slog.Debug("[SaveModel]", "query", __obf_a856c5ef2c6a7852, "args", __obf_cd0ddd46582953ac)
	}

	return __obf_3d258fae35428ba8, nil

}

func (__obf_86bb12f77b537935 *ORM) SaveModelWidthAutoID(__obf_f2992f05d328343c any, __obf_3d258fae35428ba8 int64) (int64, error) {
	// query, args := modelBind(db.Option, m, insertId)
	__obf_306dddc11fc3d690 := reflect.ValueOf(__obf_f2992f05d328343c)
	if __obf_306dddc11fc3d690.Kind() == reflect.Pointer {
		__obf_306dddc11fc3d690 = __obf_306dddc11fc3d690.Elem()
	}
	__obf_d24f398beafeb263 := __obf_306dddc11fc3d690.FieldByName("Id")
	__obf_82ed77b51afe2acc := util.ToSnake(reflect.Indirect(__obf_306dddc11fc3d690).Type().Name())
	var (
		__obf_a856c5ef2c6a7852 string
		__obf_cd0ddd46582953ac map[string]any
		__obf_e6c24855347eb3d0 error
	)

	if __obf_d24f398beafeb263.IsZero() {
		// insert
		__obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac, __obf_e6c24855347eb3d0 = BuildInsertSQL(__obf_f2992f05d328343c, __obf_82ed77b51afe2acc)
		if __obf_e6c24855347eb3d0 != nil {
			return 0, fmt.Errorf("failed to build insert sql: %s", __obf_e6c24855347eb3d0.Error())
		}
		if __obf_3d258fae35428ba8 != 0 {
			__obf_cd0ddd46582953ac["id"] = __obf_3d258fae35428ba8
		}
		__obf_10f23379b4bda68b := time.Now().Unix()
		if _, __obf_b5d20856a702a4bb := __obf_cd0ddd46582953ac[__obf_86bb12f77b537935.Option.Reserved.CreateTime]; __obf_b5d20856a702a4bb {
			__obf_cd0ddd46582953ac[__obf_86bb12f77b537935.Option.Reserved.CreateTime] = __obf_10f23379b4bda68b
		}
		if _, __obf_b5d20856a702a4bb := __obf_cd0ddd46582953ac[__obf_86bb12f77b537935.Option.Reserved.UpdateTime]; __obf_b5d20856a702a4bb {
			__obf_cd0ddd46582953ac[__obf_86bb12f77b537935.Option.Reserved.UpdateTime] = __obf_10f23379b4bda68b
		}
	} else {
		// update
		__obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac, __obf_e6c24855347eb3d0 = BuildUpdateSQL(__obf_f2992f05d328343c, __obf_82ed77b51afe2acc, __obf_86bb12f77b537935.Option.Reserved.UpdateTime, []string{__obf_86bb12f77b537935.Option.Reserved.ID}, __obf_86bb12f77b537935.Option.Reserved.CreateTime, __obf_86bb12f77b537935.Option.Reserved.NO)
		if __obf_e6c24855347eb3d0 != nil {
			return 0, fmt.Errorf("failed to build update sql: %s", __obf_e6c24855347eb3d0.Error())
		}
	}

	if __obf_a856c5ef2c6a7852 != "" {
		__obf_6ab9f0211d0f4779, __obf_e6c24855347eb3d0 := __obf_86bb12f77b537935.DB.NamedExec(__obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac)
		if __obf_e6c24855347eb3d0 != nil {
			return 0, fmt.Errorf("[SaveModelWidthAutoID]: %s", __obf_e6c24855347eb3d0.Error())
		}
		if __obf_3d258fae35428ba8 == 0 {
			return __obf_6ab9f0211d0f4779.LastInsertId()
		}

		slog.Debug("SaveModelWidthAutoID", "query", __obf_a856c5ef2c6a7852, "args", __obf_cd0ddd46582953ac)
	}

	return __obf_3d258fae35428ba8, nil
}

func (__obf_86bb12f77b537935 *ORM) SaveTxModel(__obf_f1097d4e39bff719 *sqlx.Tx, __obf_f2992f05d328343c any, __obf_3d258fae35428ba8 string) error {
	var __obf_a856c5ef2c6a7852 string
	var __obf_cd0ddd46582953ac map[string]any
	__obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac = __obf_16f30d616c25b523(__obf_86bb12f77b537935.Option, __obf_f2992f05d328343c, __obf_3d258fae35428ba8)

	if __obf_a856c5ef2c6a7852 != "" {
		if _, __obf_e6c24855347eb3d0 := __obf_f1097d4e39bff719.NamedExec(__obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac); __obf_e6c24855347eb3d0 != nil {
			_ = __obf_f1097d4e39bff719.Rollback()
			return fmt.Errorf("SaveTxModel: %v", __obf_e6c24855347eb3d0)
		}

		slog.Debug("SaveTxModel", "query", __obf_a856c5ef2c6a7852, "args", __obf_cd0ddd46582953ac)

	}

	return nil
}

// 参数insertId用于在新增时，自己传入的ID，以方便作外键使用
func __obf_16f30d616c25b523[RecordID int | string](__obf_2d84a01fd78b7ed8 *Options, __obf_f2992f05d328343c any, __obf_3d258fae35428ba8 RecordID) (__obf_1b70b52240d0a975 string, __obf_cd0ddd46582953ac map[string]any) {
	__obf_306dddc11fc3d690 := reflect.ValueOf(__obf_f2992f05d328343c)
	if __obf_306dddc11fc3d690.Kind() == reflect.Pointer {
		__obf_306dddc11fc3d690 = __obf_306dddc11fc3d690.Elem()
	}
	var __obf_aa54a30bd96457fa []string
	__obf_cd0ddd46582953ac = make(map[string]any)

	__obf_d24f398beafeb263 := __obf_306dddc11fc3d690.FieldByName("Id")
	if __obf_d24f398beafeb263.IsZero() {
		// 判断inserId是否为零值
		if __obf_3d258fae35428ba8 != *new(RecordID) {
			__obf_aa54a30bd96457fa = append(__obf_aa54a30bd96457fa, __obf_2d84a01fd78b7ed8.Reserved.ID)
			__obf_cd0ddd46582953ac[__obf_2d84a01fd78b7ed8.Reserved.ID] = __obf_3d258fae35428ba8
		} else {
			if __obf_d24f398beafeb263.Kind() == reflect.String {
				__obf_aa54a30bd96457fa = append(__obf_aa54a30bd96457fa, __obf_2d84a01fd78b7ed8.Reserved.ID)
				__obf_cd0ddd46582953ac[__obf_2d84a01fd78b7ed8.Reserved.ID] = util.StringUUID()
				// idField.Set(reflect.ValueOf(insertId))
			}
			// 如果ID为int型, 视为数据库自增
		}

		__obf_e5d4c6d6d92e4353 := reflect.Indirect(__obf_306dddc11fc3d690).Type()
		for __obf_935763ab96242d2f := range __obf_e5d4c6d6d92e4353.NumField() {
			if __obf_5a07a2990ad85783 := __obf_e5d4c6d6d92e4353.Field(__obf_935763ab96242d2f); __obf_5a07a2990ad85783.Tag.Get("db") != "-" {
				switch __obf_c804ee49b6774cf8 := util.ToSnake(__obf_5a07a2990ad85783.Name); __obf_c804ee49b6774cf8 {
				case __obf_2d84a01fd78b7ed8.Reserved.CreateTime, __obf_2d84a01fd78b7ed8.Reserved.UpdateTime:
					__obf_aa54a30bd96457fa = append(__obf_aa54a30bd96457fa, __obf_c804ee49b6774cf8)
					__obf_cd0ddd46582953ac[__obf_c804ee49b6774cf8] = time.Now().Unix()
				default:
					__obf_1115ae3c80598454 := __obf_306dddc11fc3d690.Field(__obf_935763ab96242d2f).Interface()
					if strings.HasPrefix(__obf_c804ee49b6774cf8, "is_") || __obf_1115ae3c80598454 != reflect.Zero(__obf_5a07a2990ad85783.Type).Interface() {
						__obf_aa54a30bd96457fa = append(__obf_aa54a30bd96457fa, __obf_c804ee49b6774cf8)
						__obf_cd0ddd46582953ac[__obf_c804ee49b6774cf8] = __obf_1115ae3c80598454
					}
				}
			}
		}
		__obf_1b70b52240d0a975 = fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)", util.ToSnake(__obf_e5d4c6d6d92e4353.Name()), strings.Join(__obf_aa54a30bd96457fa, ","), strings.Join(__obf_aa54a30bd96457fa, ",:"))
	} else {
		__obf_e5d4c6d6d92e4353 := reflect.Indirect(__obf_306dddc11fc3d690).Type()
		for __obf_935763ab96242d2f := range __obf_e5d4c6d6d92e4353.NumField() {
			if __obf_5a07a2990ad85783 := __obf_e5d4c6d6d92e4353.Field(__obf_935763ab96242d2f); __obf_5a07a2990ad85783.Tag.Get("db") != "-" {
				__obf_1115ae3c80598454 := __obf_306dddc11fc3d690.Field(__obf_935763ab96242d2f).Interface()
				if __obf_c804ee49b6774cf8 := util.ToSnake(__obf_5a07a2990ad85783.Name); strings.HasPrefix(__obf_c804ee49b6774cf8, "is_") || __obf_1115ae3c80598454 != reflect.Zero(__obf_5a07a2990ad85783.Type).Interface() {
					__obf_aa54a30bd96457fa = append(__obf_aa54a30bd96457fa, fmt.Sprintf("%s=:%s", __obf_c804ee49b6774cf8, __obf_c804ee49b6774cf8))
					__obf_cd0ddd46582953ac[__obf_c804ee49b6774cf8] = __obf_1115ae3c80598454
				}
			}
		}

		if util.HasField(__obf_306dddc11fc3d690, util.ToCamel(__obf_2d84a01fd78b7ed8.Reserved.UpdateTime)) {
			__obf_9b1f0a5fbd2939dc := __obf_cd0ddd46582953ac[__obf_2d84a01fd78b7ed8.Reserved.UpdateTime]
			__obf_cd0ddd46582953ac[__obf_2d84a01fd78b7ed8.Reserved.UpdateTime] = time.Now().Unix()
			__obf_1b70b52240d0a975 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s AND %s=%v",
				util.ToSnake(__obf_e5d4c6d6d92e4353.Name()), strings.Join(__obf_aa54a30bd96457fa, ","), __obf_2d84a01fd78b7ed8.Reserved.ID, __obf_2d84a01fd78b7ed8.Reserved.ID, __obf_2d84a01fd78b7ed8.Reserved.UpdateTime, __obf_9b1f0a5fbd2939dc)
		} else {
			__obf_1b70b52240d0a975 = fmt.Sprintf("UPDATE %s SET %s WHERE %s=:%s", util.ToSnake(__obf_e5d4c6d6d92e4353.Name()), strings.Join(__obf_aa54a30bd96457fa, ","), __obf_2d84a01fd78b7ed8.Reserved.ID, __obf_2d84a01fd78b7ed8.Reserved.ID)
		}
	}
	return
}

func BuildInsertSQL(__obf_4fc3c32163040478 any, __obf_82ed77b51afe2acc string) (string, map[string]any, error) {
	if __obf_82ed77b51afe2acc == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_cd0ddd46582953ac := make(map[string]any)
	__obf_aa54a30bd96457fa := []string{}

	__obf_40538d34494d25da := func(__obf_9ad5b2ee11618c02 string, __obf_bcc644fb72af6d31 any, __obf_9ac10b6ff8c844bc bool) {
		__obf_7b3ba33e69f9dcf8 := __obf_bcc644fb72af6d31 == nil || reflect.DeepEqual(__obf_bcc644fb72af6d31, reflect.Zero(reflect.TypeOf(__obf_bcc644fb72af6d31)).Interface())
		if __obf_9ac10b6ff8c844bc && __obf_7b3ba33e69f9dcf8 {
			return
		}
		__obf_aa54a30bd96457fa = append(__obf_aa54a30bd96457fa, __obf_9ad5b2ee11618c02)
		__obf_cd0ddd46582953ac[__obf_9ad5b2ee11618c02] = __obf_bcc644fb72af6d31
	}

	switch __obf_1115ae3c80598454 := __obf_4fc3c32163040478.(type) {
	case map[string]any:
		for __obf_13037cdadbf8d35e, __obf_306dddc11fc3d690 := range __obf_1115ae3c80598454 {
			__obf_40538d34494d25da(__obf_13037cdadbf8d35e, __obf_306dddc11fc3d690, false)
		}

	default:
		__obf_c157396a957f0d01 := reflect.TypeOf(__obf_4fc3c32163040478)
		if __obf_c157396a957f0d01.Kind() == reflect.Pointer {
			__obf_c157396a957f0d01 = __obf_c157396a957f0d01.Elem()
		}
		__obf_306dddc11fc3d690 := reflect.ValueOf(__obf_4fc3c32163040478)
		if __obf_306dddc11fc3d690.Kind() == reflect.Pointer {
			__obf_306dddc11fc3d690 = __obf_306dddc11fc3d690.Elem()
		}

		if __obf_c157396a957f0d01.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_935763ab96242d2f := 0; __obf_935763ab96242d2f < __obf_c157396a957f0d01.NumField(); __obf_935763ab96242d2f++ {
			__obf_861fd84f1f79b8e6 := __obf_c157396a957f0d01.Field(__obf_935763ab96242d2f).Tag.Get("db")
			if __obf_861fd84f1f79b8e6 == "" || __obf_861fd84f1f79b8e6 == "-" {
				continue
			}

			__obf_8f5c84d0ab8f0fb8 := strings.Split(__obf_861fd84f1f79b8e6, ",")
			__obf_40538d34494d25da(__obf_8f5c84d0ab8f0fb8[0], __obf_306dddc11fc3d690.Field(__obf_935763ab96242d2f).Interface(), len(__obf_8f5c84d0ab8f0fb8) > 1 && __obf_8f5c84d0ab8f0fb8[1] == "omitempty")
		}
	}

	if len(__obf_aa54a30bd96457fa) == 0 {
		return "", nil, errors.New("no fields to insert")
	}

	__obf_a856c5ef2c6a7852 := fmt.Sprintf("INSERT INTO %s(%s)VALUES(:%s)",
		__obf_82ed77b51afe2acc,
		strings.Join(__obf_aa54a30bd96457fa, ","),
		strings.Join(__obf_aa54a30bd96457fa, ",:"),
	)

	return __obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac, nil
}

// pkField 主键字段
// lockField 乐观锁字段(默认为updated_at)
// constField 固定不变字段
func BuildUpdateSQL(
	__obf_4fc3c32163040478 any,
	__obf_82ed77b51afe2acc, __obf_b58b462373896d85 string,
	__obf_7895e480634c87cc []string,
	__obf_da32d0938efd699f ...string,
) (string, map[string]any, error) {
	if __obf_82ed77b51afe2acc == "" {
		return "", nil, errors.New("table name is required")
	}

	__obf_d08b4495dccb2aa1 := make(map[string]struct{})
	for _, __obf_a50fe7d91bea881e := range __obf_7895e480634c87cc {
		__obf_d08b4495dccb2aa1[__obf_a50fe7d91bea881e] = struct{}{}
	}

	__obf_cd0ddd46582953ac := make(map[string]any)
	__obf_aa54a30bd96457fa := []string{}
	__obf_af2ae7a6a98c4ac7 := []string{}

	__obf_40538d34494d25da := func(__obf_9ad5b2ee11618c02 string, __obf_bcc644fb72af6d31 any, __obf_9ac10b6ff8c844bc bool) {
		__obf_7b3ba33e69f9dcf8 := __obf_bcc644fb72af6d31 == nil || reflect.DeepEqual(__obf_bcc644fb72af6d31, reflect.Zero(reflect.TypeOf(__obf_bcc644fb72af6d31)).Interface())

		switch {
		case func() bool {
			_, __obf_b5d20856a702a4bb := __obf_d08b4495dccb2aa1[__obf_9ad5b2ee11618c02]
			return __obf_b5d20856a702a4bb
		}():
			__obf_af2ae7a6a98c4ac7 = append(__obf_af2ae7a6a98c4ac7, fmt.Sprintf("%s=:%s", __obf_9ad5b2ee11618c02, __obf_9ad5b2ee11618c02))
			__obf_cd0ddd46582953ac[__obf_9ad5b2ee11618c02] = __obf_bcc644fb72af6d31

		case slices.Contains(__obf_da32d0938efd699f, __obf_9ad5b2ee11618c02):
			// skip constField
		case __obf_9ad5b2ee11618c02 == __obf_b58b462373896d85:
			const __obf_1f7717855065cd3d = "RESERVED_LOCK"
			__obf_aa54a30bd96457fa = append(__obf_aa54a30bd96457fa, fmt.Sprintf("%s=:%s", __obf_9ad5b2ee11618c02, __obf_9ad5b2ee11618c02))
			__obf_af2ae7a6a98c4ac7 = append(__obf_af2ae7a6a98c4ac7, fmt.Sprintf("%s=:%s", __obf_9ad5b2ee11618c02, __obf_1f7717855065cd3d))
			__obf_cd0ddd46582953ac[__obf_9ad5b2ee11618c02] = time.Now().Unix()
			__obf_cd0ddd46582953ac[__obf_1f7717855065cd3d] = __obf_bcc644fb72af6d31

		default:
			if __obf_9ac10b6ff8c844bc && __obf_7b3ba33e69f9dcf8 {
				return
			}
			__obf_aa54a30bd96457fa = append(__obf_aa54a30bd96457fa, fmt.Sprintf("%s=:%s", __obf_9ad5b2ee11618c02, __obf_9ad5b2ee11618c02))
			__obf_cd0ddd46582953ac[__obf_9ad5b2ee11618c02] = __obf_bcc644fb72af6d31
		}
	}

	switch __obf_1115ae3c80598454 := __obf_4fc3c32163040478.(type) {
	case map[string]any:
		for __obf_13037cdadbf8d35e, __obf_306dddc11fc3d690 := range __obf_1115ae3c80598454 {
			__obf_40538d34494d25da(__obf_13037cdadbf8d35e, __obf_306dddc11fc3d690, false)
		}

	default:

		__obf_c157396a957f0d01 := reflect.TypeOf(__obf_4fc3c32163040478)
		if __obf_c157396a957f0d01.Kind() == reflect.Pointer {
			__obf_c157396a957f0d01 = __obf_c157396a957f0d01.Elem()
		}
		__obf_306dddc11fc3d690 := reflect.ValueOf(__obf_4fc3c32163040478)
		if __obf_306dddc11fc3d690.Kind() == reflect.Pointer {
			__obf_306dddc11fc3d690 = __obf_306dddc11fc3d690.Elem()
		}
		if __obf_c157396a957f0d01.Kind() != reflect.Struct {
			return "", nil, errors.New("input must be a struct or map")
		}

		for __obf_935763ab96242d2f := 0; __obf_935763ab96242d2f < __obf_c157396a957f0d01.NumField(); __obf_935763ab96242d2f++ {
			__obf_861fd84f1f79b8e6 := __obf_c157396a957f0d01.Field(__obf_935763ab96242d2f).Tag.Get("db")
			if __obf_861fd84f1f79b8e6 == "" || __obf_861fd84f1f79b8e6 == "-" {
				continue
			}

			__obf_8f5c84d0ab8f0fb8 := strings.Split(__obf_861fd84f1f79b8e6, ",")
			__obf_40538d34494d25da(__obf_8f5c84d0ab8f0fb8[0], __obf_306dddc11fc3d690.Field(__obf_935763ab96242d2f).Interface(), len(__obf_8f5c84d0ab8f0fb8) > 1 && __obf_8f5c84d0ab8f0fb8[1] == "omitempty")
		}
	}

	if len(__obf_7895e480634c87cc) == 0 || len(__obf_af2ae7a6a98c4ac7) == 0 {
		return "", nil, errors.New("at least one primary key is required")
	}
	if len(__obf_aa54a30bd96457fa) == 0 {
		return "", nil, errors.New("no fields to update")
	}
	// if lockField != "" && !hasLockField {
	// 	return "", nil, fmt.Errorf("updatedAt field '%s' not found or zero", lockField)
	// }

	__obf_a856c5ef2c6a7852 := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		__obf_82ed77b51afe2acc,
		strings.Join(__obf_aa54a30bd96457fa, ","),
		strings.Join(__obf_af2ae7a6a98c4ac7, " AND "),
	)

	return __obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac, nil
}
