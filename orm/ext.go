package __obf_f47aac06a08e5dbb

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	Grouping    = " and " //where条件拼接
	ParamMarker = "?"     //参数占位符
	DBTag       = "db"    //数据库tag
	SeqComma    = ", "    //字段分隔符
	SeqSpace    = " "     //空格分隔符
)

const (
	SelectTypeOne = iota
	SelectTypeMany
)

// 用单参数，函数内部调用自行转换类型，否则没办法传递，很烦
type FuncTx func(__obf_07b88b883eeed833 *sqlx.Tx, __obf_0bdb4f0ba995788b any) error

func Open(__obf_f7806770e1e234d2, __obf_c3e0af02fcf92606 string, __obf_03d062ef502ff660 time.Duration) (*DB, error) {
	__obf_ab97c1322d9ea2f2, __obf_9ac6018edd4832e7 := sqlx.Open(__obf_f7806770e1e234d2, __obf_c3e0af02fcf92606)
	if __obf_9ac6018edd4832e7 != nil {
		return nil, __obf_9ac6018edd4832e7
	}
	__obf_9554751b39677f0e := &DB{
		DB: __obf_ab97c1322d9ea2f2, __obf_03d062ef502ff660: __obf_03d062ef502ff660,
	}
	__obf_9554751b39677f0e.__obf_5f1f08b7d43e5ae9.
		New = func() any {
		return __obf_9554751b39677f0e.__obf_625ac88ebd2f624e()
	}
	return __obf_9554751b39677f0e, nil
}

type DB struct {
	*sqlx.DB
	__obf_03d062ef502ff660 time.Duration
	__obf_5f1f08b7d43e5ae9 sync.Pool
}

func (__obf_ab97c1322d9ea2f2 *DB) __obf_625ac88ebd2f624e() *Context {
	return &Context{__obf_ab97c1322d9ea2f2: __obf_ab97c1322d9ea2f2}
}

// 获取一个`SQL`执行`Context`
func (__obf_ab97c1322d9ea2f2 *DB) Acquire() *Context {
	__obf_098eeac6ede2fc65 := // 无需加锁，sync.Pool本身是线程安全的
		__obf_ab97c1322d9ea2f2.__obf_5f1f08b7d43e5ae9.Get().(*Context)
	__obf_098eeac6ede2fc65.__obf_41aaef56d88a426a()
	return __obf_098eeac6ede2fc65
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_ab97c1322d9ea2f2 *DB) AcquireTx(__obf_07b88b883eeed833 *sqlx.Tx) *Context {
	__obf_098eeac6ede2fc65 := __obf_ab97c1322d9ea2f2.Acquire()
	__obf_098eeac6ede2fc65.__obf_07b88b883eeed833 = __obf_07b88b883eeed833
	return __obf_098eeac6ede2fc65
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_ab97c1322d9ea2f2 *DB) WithTx(__obf_620bf407981a530b FuncTx, __obf_0bdb4f0ba995788b any) (__obf_9ac6018edd4832e7 error) {
	var __obf_07b88b883eeed833 *sqlx.Tx
	__obf_07b88b883eeed833, __obf_9ac6018edd4832e7 = __obf_ab97c1322d9ea2f2.Beginx()
	if __obf_9ac6018edd4832e7 != nil {
		return
	}
	defer func() {
		if __obf_9ac6018edd4832e7 != nil && __obf_07b88b883eeed833 != nil {
			__obf_9ac6018edd4832e7 = __obf_07b88b883eeed833.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_9ac6018edd4832e7 = __obf_620bf407981a530b(__obf_07b88b883eeed833, __obf_0bdb4f0ba995788b); __obf_9ac6018edd4832e7 != nil {
		return
	}
	__obf_9ac6018edd4832e7 = __obf_07b88b883eeed833.Commit()
	return
}

type Context struct {
	__obf_ab97c1322d9ea2f2 *DB
	__obf_07b88b883eeed833 *sqlx.Tx //事务
	sql                    string
	__obf_87a0df0fc6580d03 string
	__obf_8f7b83fdb4eae7ee []string
	__obf_b07f8c87122c571f []string
	__obf_653d68f29adf06d8 string
	__obf_eec5bdbc706cb503 string
	__obf_a527be6019096155 string
	__obf_04b7b643418a3e38 int64
	__obf_3a959f0103dcd624 int64
	__obf_0bdb4f0ba995788b []any
	__obf_2cb967dab9501e7b bool
	__obf_079e86c701ea5ab8 bool //排他锁
	//共享锁
}

func (__obf_098eeac6ede2fc65 *Context) Name(__obf_87a0df0fc6580d03 string) *Context {
	__obf_098eeac6ede2fc65.__obf_87a0df0fc6580d03 = __obf_87a0df0fc6580d03
	return __obf_098eeac6ede2fc65
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_098eeac6ede2fc65 *Context) What(__obf_8f7b83fdb4eae7ee []string) *Context {
	__obf_098eeac6ede2fc65.__obf_8f7b83fdb4eae7ee = __obf_8f7b83fdb4eae7ee
	return __obf_098eeac6ede2fc65
}

func (__obf_098eeac6ede2fc65 *Context) Where(__obf_23d3ae8f01b6a47d string, __obf_0bdb4f0ba995788b ...any) *Context {
	__obf_098eeac6ede2fc65.__obf_b07f8c87122c571f = append(__obf_098eeac6ede2fc65.__obf_b07f8c87122c571f, __obf_23d3ae8f01b6a47d)
	__obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b = append(__obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b,

		// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
		__obf_0bdb4f0ba995788b...)
	return __obf_098eeac6ede2fc65
}

func (__obf_098eeac6ede2fc65 *Context) WhereIn(__obf_40948bc1feaa6823 string, __obf_0bdb4f0ba995788b []any) *Context {
	__obf_2c8977f13139637e := len(__obf_0bdb4f0ba995788b)
	__obf_e5385f30d33f6bde := make([]string, __obf_2c8977f13139637e)
	for __obf_89306d663fb5edd8 := 0; __obf_89306d663fb5edd8 < __obf_2c8977f13139637e; __obf_89306d663fb5edd8++ {
		__obf_e5385f30d33f6bde[__obf_89306d663fb5edd8] = ParamMarker
	}
	__obf_45ec347c52f47e44 := fmt.Sprintf("%s in (%s)", __obf_40948bc1feaa6823, __obf_fcd3b5c493b54843(__obf_e5385f30d33f6bde, SeqComma))
	return __obf_098eeac6ede2fc65.Where(__obf_45ec347c52f47e44, __obf_0bdb4f0ba995788b...)
}

func (__obf_098eeac6ede2fc65 *Context) Order(__obf_653d68f29adf06d8 string) *Context {
	__obf_098eeac6ede2fc65.__obf_653d68f29adf06d8 = __obf_653d68f29adf06d8
	return __obf_098eeac6ede2fc65
}

func (__obf_098eeac6ede2fc65 *Context) Limit(__obf_04b7b643418a3e38 int64) *Context {
	__obf_098eeac6ede2fc65.__obf_04b7b643418a3e38 = __obf_04b7b643418a3e38
	return __obf_098eeac6ede2fc65
}

func (__obf_098eeac6ede2fc65 *Context) Offset(__obf_3a959f0103dcd624 int64) *Context {
	__obf_098eeac6ede2fc65.__obf_3a959f0103dcd624 = __obf_3a959f0103dcd624
	return __obf_098eeac6ede2fc65
}

func (__obf_098eeac6ede2fc65 *Context) Group(__obf_eec5bdbc706cb503 string) *Context {
	__obf_098eeac6ede2fc65.__obf_eec5bdbc706cb503 = __obf_eec5bdbc706cb503
	return __obf_098eeac6ede2fc65
}

func (__obf_098eeac6ede2fc65 *Context) Having(__obf_a527be6019096155 string, __obf_0bdb4f0ba995788b ...any) *Context {
	__obf_098eeac6ede2fc65.__obf_a527be6019096155 = __obf_a527be6019096155
	__obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b = append(__obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b,

		// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
		__obf_0bdb4f0ba995788b...)
	return __obf_098eeac6ede2fc65
}

func (__obf_098eeac6ede2fc65 *Context) LockX() *Context {
	__obf_098eeac6ede2fc65.__obf_2cb967dab9501e7b = true
	return __obf_098eeac6ede2fc65
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_098eeac6ede2fc65 *Context) LockS() *Context {
	__obf_098eeac6ede2fc65.__obf_079e86c701ea5ab8 = true
	return __obf_098eeac6ede2fc65
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_098eeac6ede2fc65 *Context) FindMany(__obf_122697469ced05d2 any) error {
	return __obf_098eeac6ede2fc65.__obf_39512023d5849c75(__obf_122697469ced05d2, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_098eeac6ede2fc65 *Context) FindById(__obf_122697469ced05d2 any) error {
	return __obf_098eeac6ede2fc65.__obf_39512023d5849c75(__obf_122697469ced05d2, SelectTypeOne)
}

// 插入
func (__obf_098eeac6ede2fc65 *Context) Insert(__obf_c9a5e2cfc6cf7a99 map[string]any) (sql.Result, error) {
	var (
		__obf_4a4fa4815616bb49 []string
		__obf_20b20d9125646929 []any
	)
	for __obf_5103bf8cf045836b, __obf_d6e1d0a7185fed66 := range __obf_c9a5e2cfc6cf7a99 {
		__obf_4a4fa4815616bb49 = append(__obf_4a4fa4815616bb49, __obf_5103bf8cf045836b)
		__obf_20b20d9125646929 = append(__obf_20b20d9125646929, __obf_d6e1d0a7185fed66)
	}
	return __obf_098eeac6ede2fc65.InsertBatch(__obf_4a4fa4815616bb49,

		// 批量插入
		__obf_20b20d9125646929)
}

func (__obf_098eeac6ede2fc65 *Context) InsertBatch(__obf_4a4fa4815616bb49 []string, __obf_c9a5e2cfc6cf7a99 ...[]any) (sql.Result, error) {
	var (
		__obf_20b20d9125646929 []any
		__obf_bbb04da4d82674ec []string
	)
	for _, __obf_6dc424237c51fb36 := range __obf_c9a5e2cfc6cf7a99 {
		__obf_e5385f30d33f6bde := make([]string, len(__obf_6dc424237c51fb36))
		for __obf_89306d663fb5edd8, __obf_d6e1d0a7185fed66 := range __obf_6dc424237c51fb36 {
			__obf_e5385f30d33f6bde[__obf_89306d663fb5edd8] = ParamMarker
			__obf_20b20d9125646929 = append(__obf_20b20d9125646929, __obf_d6e1d0a7185fed66)
		}
		__obf_bbb04da4d82674ec = append(__obf_bbb04da4d82674ec, fmt.Sprintf("(%s)", __obf_fcd3b5c493b54843(__obf_e5385f30d33f6bde, SeqComma)))
	}
	__obf_b7fc1611623cb763 := fmt.Sprintf("insert into %s (%s) values %s", __obf_098eeac6ede2fc65.__obf_87a0df0fc6580d03, __obf_fcd3b5c493b54843(__obf_4a4fa4815616bb49, SeqComma), __obf_fcd3b5c493b54843(__obf_bbb04da4d82674ec, SeqComma))
	return __obf_098eeac6ede2fc65.__obf_e5edb8c4f78dbea0(__obf_b7fc1611623cb763,

		// 使用map更新
		__obf_20b20d9125646929...)
}

func (__obf_098eeac6ede2fc65 *Context) UpdateMap(__obf_0bdb4f0ba995788b map[string]any) (__obf_974f2ba70ed94346 int64, __obf_9ac6018edd4832e7 error) {
	var (
		__obf_20b20d9125646929 []any
		__obf_267186dd1b490532 []string
	)
	for __obf_5103bf8cf045836b, __obf_d6e1d0a7185fed66 := range __obf_0bdb4f0ba995788b {
		__obf_20b20d9125646929 = append(__obf_20b20d9125646929, __obf_d6e1d0a7185fed66)
		__obf_267186dd1b490532 = append(__obf_267186dd1b490532, fmt.Sprintf("%s=%s", __obf_5103bf8cf045836b, ParamMarker))
	}
	__obf_869b92bc5d3ef3db := __obf_fcd3b5c493b54843(__obf_267186dd1b490532, SeqComma)
	__obf_974f2ba70ed94346, __obf_9ac6018edd4832e7 = __obf_098eeac6ede2fc65.Update(__obf_869b92bc5d3ef3db, __obf_20b20d9125646929...)
	return
}

// 更新
func (__obf_098eeac6ede2fc65 *Context) Update(__obf_869b92bc5d3ef3db string, __obf_0bdb4f0ba995788b ...any) (__obf_974f2ba70ed94346 int64, __obf_9ac6018edd4832e7 error) {
	__obf_c84964fdc0f6733f := "update %s set %s %s"
	__obf_23d3ae8f01b6a47d := __obf_8d9d0550fc248026(__obf_098eeac6ede2fc65.__obf_b07f8c87122c571f, Grouping)
	__obf_b7fc1611623cb763 := fmt.Sprintf(__obf_c84964fdc0f6733f, __obf_098eeac6ede2fc65.__obf_87a0df0fc6580d03, __obf_869b92bc5d3ef3db, __obf_23d3ae8f01b6a47d)
	__obf_20b20d9125646929 := append(__obf_0bdb4f0ba995788b, __obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b...)
	var __obf_c565a64719fc9ebe sql.Result
	__obf_c565a64719fc9ebe, __obf_9ac6018edd4832e7 = __obf_098eeac6ede2fc65.__obf_e5edb8c4f78dbea0(__obf_b7fc1611623cb763, __obf_20b20d9125646929...)
	if __obf_9ac6018edd4832e7 != nil {
		return
	}
	__obf_974f2ba70ed94346, __obf_9ac6018edd4832e7 = __obf_c565a64719fc9ebe.RowsAffected()
	return
}

// 删除
func (__obf_098eeac6ede2fc65 *Context) Delete() (__obf_974f2ba70ed94346 int64, __obf_9ac6018edd4832e7 error) {
	__obf_c84964fdc0f6733f := "delete from %s %s"
	__obf_23d3ae8f01b6a47d := __obf_8d9d0550fc248026(__obf_098eeac6ede2fc65.__obf_b07f8c87122c571f, Grouping)
	__obf_b7fc1611623cb763 := fmt.Sprintf(__obf_c84964fdc0f6733f, __obf_098eeac6ede2fc65.__obf_87a0df0fc6580d03, __obf_23d3ae8f01b6a47d)
	var __obf_c565a64719fc9ebe sql.Result
	__obf_c565a64719fc9ebe, __obf_9ac6018edd4832e7 = __obf_098eeac6ede2fc65.__obf_e5edb8c4f78dbea0(__obf_b7fc1611623cb763, __obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b...)
	if __obf_9ac6018edd4832e7 != nil {
		return
	}
	__obf_974f2ba70ed94346, __obf_9ac6018edd4832e7 = __obf_c565a64719fc9ebe.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_098eeac6ede2fc65 *Context) Select(__obf_122697469ced05d2 any, sql string, __obf_0bdb4f0ba995788b ...any) error {
	__obf_098eeac6ede2fc65.
		sql = sql
	__obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b = __obf_0bdb4f0ba995788b
	return __obf_098eeac6ede2fc65.__obf_39512023d5849c75(__obf_122697469ced05d2, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_098eeac6ede2fc65 *Context) Get(__obf_122697469ced05d2 any, sql string, __obf_0bdb4f0ba995788b ...any) error {
	__obf_098eeac6ede2fc65.
		sql = sql
	__obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b = __obf_0bdb4f0ba995788b
	return __obf_098eeac6ede2fc65.__obf_39512023d5849c75(__obf_122697469ced05d2, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_098eeac6ede2fc65 *Context) Exec(sql string, __obf_0bdb4f0ba995788b ...any) (sql.Result, error) {
	return __obf_098eeac6ede2fc65.__obf_e5edb8c4f78dbea0(sql, __obf_0bdb4f0ba995788b...)
}

// 创建表
func (__obf_098eeac6ede2fc65 *Context) Create(sql string) (sql.Result, error) {
	return __obf_098eeac6ede2fc65.

		// 删除表
		__obf_e5edb8c4f78dbea0(sql)
}

func (__obf_098eeac6ede2fc65 *Context) Drop() (sql.Result, error) {
	return __obf_098eeac6ede2fc65.__obf_e5edb8c4f78dbea0(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_098eeac6ede2fc65.

		/////////////////////////private methods//////////////////////
		__obf_87a0df0fc6580d03))
}

// 重置Context
func (__obf_098eeac6ede2fc65 *Context) __obf_41aaef56d88a426a() *Context {
	__obf_098eeac6ede2fc65.
		sql = ""
	__obf_098eeac6ede2fc65.__obf_87a0df0fc6580d03 = ""
	__obf_098eeac6ede2fc65.__obf_8f7b83fdb4eae7ee = []string{}
	__obf_098eeac6ede2fc65.__obf_b07f8c87122c571f = []string{}
	__obf_098eeac6ede2fc65.__obf_653d68f29adf06d8 = ""
	__obf_098eeac6ede2fc65.__obf_eec5bdbc706cb503 = ""
	__obf_098eeac6ede2fc65.__obf_a527be6019096155 = ""
	__obf_098eeac6ede2fc65.__obf_04b7b643418a3e38 = 0
	__obf_098eeac6ede2fc65.__obf_3a959f0103dcd624 = 0
	__obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b = []any{}
	__obf_098eeac6ede2fc65.__obf_07b88b883eeed833 = nil
	__obf_098eeac6ede2fc65.__obf_079e86c701ea5ab8 = false
	__obf_098eeac6ede2fc65.__obf_2cb967dab9501e7b = false
	return __obf_098eeac6ede2fc65
}

// 查询方法
func (__obf_098eeac6ede2fc65 *Context) __obf_39512023d5849c75(__obf_122697469ced05d2 any, __obf_e8f614714033efee int) (__obf_9ac6018edd4832e7 error) {
	defer __obf_098eeac6ede2fc65.__obf_ab97c1322d9ea2f2.__obf_5f1f08b7d43e5ae9.Put(__obf_098eeac6ede2fc65)
	__obf_9ae90e49d70c6aba, __obf_e605a74b56be205c := context.WithTimeout(context.Background(), __obf_098eeac6ede2fc65.__obf_ab97c1322d9ea2f2.__obf_03d062ef502ff660)
	defer __obf_e605a74b56be205c()
	if __obf_098eeac6ede2fc65.sql == "" {
		__obf_098eeac6ede2fc65.
			sql = __obf_098eeac6ede2fc65.__obf_ba5f183c67fbe1a3(__obf_122697469ced05d2)
	}
	switch __obf_e8f614714033efee {
	case SelectTypeOne:
		if __obf_098eeac6ede2fc65.__obf_07b88b883eeed833 != nil {
			__obf_9ac6018edd4832e7 = __obf_098eeac6ede2fc65.__obf_07b88b883eeed833.GetContext(__obf_9ae90e49d70c6aba, __obf_122697469ced05d2, __obf_098eeac6ede2fc65.sql, __obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b...)
		} else {
			__obf_9ac6018edd4832e7 = __obf_098eeac6ede2fc65.__obf_ab97c1322d9ea2f2.GetContext(__obf_9ae90e49d70c6aba, __obf_122697469ced05d2, __obf_098eeac6ede2fc65.sql, __obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b...)
		}
	case SelectTypeMany:
		if __obf_098eeac6ede2fc65.__obf_07b88b883eeed833 != nil {
			__obf_9ac6018edd4832e7 = __obf_098eeac6ede2fc65.__obf_07b88b883eeed833.SelectContext(__obf_9ae90e49d70c6aba, __obf_122697469ced05d2, __obf_098eeac6ede2fc65.sql, __obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b...)
		} else {
			__obf_9ac6018edd4832e7 = __obf_098eeac6ede2fc65.__obf_ab97c1322d9ea2f2.SelectContext(__obf_9ae90e49d70c6aba, __obf_122697469ced05d2, __obf_098eeac6ede2fc65.sql, __obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_098eeac6ede2fc65 *Context) __obf_e5edb8c4f78dbea0(__obf_b7fc1611623cb763 string, __obf_0bdb4f0ba995788b ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_b7fc1611623cb763, __obf_0bdb4f0ba995788b)
	defer __obf_098eeac6ede2fc65.__obf_ab97c1322d9ea2f2.__obf_5f1f08b7d43e5ae9.Put(__obf_098eeac6ede2fc65)
	__obf_9ae90e49d70c6aba, __obf_e605a74b56be205c := context.WithTimeout(context.Background(), __obf_098eeac6ede2fc65.__obf_ab97c1322d9ea2f2.__obf_03d062ef502ff660)
	defer __obf_e605a74b56be205c()

	var __obf_06143766db88766d sqlx.ExecerContext
	if __obf_098eeac6ede2fc65.__obf_07b88b883eeed833 != nil {
		__obf_06143766db88766d = __obf_098eeac6ede2fc65.__obf_07b88b883eeed833
	} else {
		__obf_06143766db88766d = __obf_098eeac6ede2fc65.__obf_ab97c1322d9ea2f2
	}
	return __obf_06143766db88766d.ExecContext(__obf_9ae90e49d70c6aba, __obf_b7fc1611623cb763,

		// select查询语句的拼接
		__obf_0bdb4f0ba995788b...)
}

func (__obf_098eeac6ede2fc65 *Context) __obf_ba5f183c67fbe1a3(__obf_122697469ced05d2 any) string {
	var __obf_0e6586f9a74f4aac []string
	__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, "select")
	if len(__obf_098eeac6ede2fc65.__obf_8f7b83fdb4eae7ee) != 0 {
		__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, __obf_fcd3b5c493b54843(__obf_098eeac6ede2fc65.__obf_8f7b83fdb4eae7ee, SeqComma))
	} else {
		__obf_b79b6029a1ef0e9e := // 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
			// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
			__obf_3ac154eba465ad9d(__obf_122697469ced05d2)
		if len(__obf_b79b6029a1ef0e9e) > 0 {
			__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, __obf_fcd3b5c493b54843(__obf_b79b6029a1ef0e9e, SeqComma))
		} else {
			__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, "*")
		}
	}
	__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, "from "+__obf_098eeac6ede2fc65.__obf_87a0df0fc6580d03)
	if len(__obf_098eeac6ede2fc65.__obf_b07f8c87122c571f) != 0 {
		__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, __obf_8d9d0550fc248026(__obf_098eeac6ede2fc65.__obf_b07f8c87122c571f, Grouping))
	}

	if __obf_098eeac6ede2fc65.__obf_eec5bdbc706cb503 != "" {
		__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, "group by "+__obf_098eeac6ede2fc65.__obf_eec5bdbc706cb503)
	}

	if __obf_098eeac6ede2fc65.__obf_a527be6019096155 != "" {
		__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, "having "+__obf_098eeac6ede2fc65.__obf_a527be6019096155)
	}

	if __obf_098eeac6ede2fc65.__obf_653d68f29adf06d8 != "" {
		__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, "order by "+__obf_098eeac6ede2fc65.__obf_653d68f29adf06d8)
	}

	if __obf_098eeac6ede2fc65.__obf_04b7b643418a3e38 != 0 {
		__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, fmt.Sprintf("limit %d, %d", __obf_098eeac6ede2fc65.__obf_3a959f0103dcd624, __obf_098eeac6ede2fc65.__obf_04b7b643418a3e38))
	}
	if __obf_098eeac6ede2fc65.__obf_079e86c701ea5ab8 {
		__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, "lock in share mode")
	}
	if __obf_098eeac6ede2fc65.__obf_2cb967dab9501e7b {
		__obf_0e6586f9a74f4aac = append(__obf_0e6586f9a74f4aac, "for update")
	}
	sql := __obf_fcd3b5c493b54843(__obf_0e6586f9a74f4aac, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_098eeac6ede2fc65.__obf_0bdb4f0ba995788b)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_8d9d0550fc248026(__obf_b07f8c87122c571f []string, __obf_3ecf61c000dac4fe string) string {
	if len(__obf_b07f8c87122c571f) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_b07f8c87122c571f, __obf_3ecf61c000dac4fe))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_fcd3b5c493b54843(__obf_0bdb4f0ba995788b []string, __obf_cbe55135fbbed312 string) string {
	return strings.Join(__obf_0bdb4f0ba995788b,

		// 解析对象中的 `db tag`
		// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
		__obf_cbe55135fbbed312)
}

func __obf_3ac154eba465ad9d(__obf_122697469ced05d2 any) (__obf_4a4fa4815616bb49 []string) {
	__obf_e51512e07c73036f := reflect.ValueOf(__obf_122697469ced05d2)
	__obf_acdd9ac22bdc242c := __obf_e51512e07c73036f.Type().Elem()
	var __obf_b55f2919dea16633 reflect.Type
	if __obf_acdd9ac22bdc242c.Kind() == reflect.Slice {
		__obf_b55f2919dea16633 = __obf_acdd9ac22bdc242c.Elem()
	} else {
		__obf_b55f2919dea16633 = __obf_acdd9ac22bdc242c
	}
	for __obf_89306d663fb5edd8 := 0; __obf_89306d663fb5edd8 < __obf_b55f2919dea16633.NumField(); __obf_89306d663fb5edd8++ {
		__obf_79243d884889b8d7 := __obf_b55f2919dea16633.Field(__obf_89306d663fb5edd8).Tag.Get(DBTag)
		if __obf_79243d884889b8d7 != "" {
			__obf_4a4fa4815616bb49 = append(__obf_4a4fa4815616bb49, __obf_79243d884889b8d7)
		}
	}
	return
}

func CleanSQLXMap(__obf_0c5e7937e9481423 map[string]any) map[string]any {
	for __obf_5103bf8cf045836b, __obf_d6e1d0a7185fed66 := range __obf_0c5e7937e9481423 {
		switch __obf_613929daa0d28687 := __obf_d6e1d0a7185fed66.(type) {
		case []byte:
			__obf_42c5c2ee071be035 := string(__obf_613929daa0d28687)
			__obf_e52e01ea79a846a7 := // 尝试智能识别具体类型
				strings.TrimSpace(__obf_42c5c2ee071be035)

			// bool
			if strings.EqualFold(__obf_e52e01ea79a846a7, "true") || strings.EqualFold(__obf_e52e01ea79a846a7, "false") {
				if __obf_10d761e27c62da28, __obf_9ac6018edd4832e7 := strconv.ParseBool(__obf_e52e01ea79a846a7); __obf_9ac6018edd4832e7 == nil {
					__obf_0c5e7937e9481423[__obf_5103bf8cf045836b] = __obf_10d761e27c62da28
					continue
				}
			}

			// int
			if __obf_89306d663fb5edd8, __obf_9ac6018edd4832e7 := strconv.Atoi(__obf_e52e01ea79a846a7); __obf_9ac6018edd4832e7 == nil {
				__obf_0c5e7937e9481423[__obf_5103bf8cf045836b] = __obf_89306d663fb5edd8
				continue
			}

			// float
			if __obf_c9ab6b5901a38cd0, __obf_9ac6018edd4832e7 := strconv.ParseFloat(__obf_e52e01ea79a846a7, 64); __obf_9ac6018edd4832e7 == nil {
				__obf_0c5e7937e9481423[__obf_5103bf8cf045836b] = __obf_c9ab6b5901a38cd0
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_e52e01ea79a846a7, "{") && strings.HasSuffix(__obf_e52e01ea79a846a7, "}")) ||
				(strings.HasPrefix(__obf_e52e01ea79a846a7, "[") && strings.HasSuffix(__obf_e52e01ea79a846a7, "]")) {
				var __obf_8d75eec75bc9880a any
				if __obf_9ac6018edd4832e7 := json.Unmarshal(__obf_613929daa0d28687, &__obf_8d75eec75bc9880a); __obf_9ac6018edd4832e7 == nil {
					__obf_0c5e7937e9481423[__obf_5103bf8cf045836b] = __obf_8d75eec75bc9880a
					continue
				}
			}
			__obf_0c5e7937e9481423[ // 默认保留为 string
			__obf_5103bf8cf045836b] = __obf_42c5c2ee071be035
		default:
			__obf_0c5e7937e9481423[__obf_5103bf8cf045836b] = __obf_613929daa0d28687
		}
	}
	return __obf_0c5e7937e9481423
}

func ToListMap(__obf_22a32b38c71f2a5d *sqlx.Rows) []map[string]any {
	var __obf_c565a64719fc9ebe []map[string]any
	for __obf_22a32b38c71f2a5d.Next() {
		__obf_0debd2347ef1c7dc := make(map[string]any)
		if __obf_9ac6018edd4832e7 := __obf_22a32b38c71f2a5d.MapScan(__obf_0debd2347ef1c7dc); __obf_9ac6018edd4832e7 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_9ac6018edd4832e7.Error()))
			continue
		}
		__obf_c565a64719fc9ebe = append(__obf_c565a64719fc9ebe, CleanSQLXMap(__obf_0debd2347ef1c7dc))
	}
	return __obf_c565a64719fc9ebe
}
