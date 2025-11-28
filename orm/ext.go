package __obf_24a21e7f5375b619

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
type FuncTx func(__obf_5c86968691329dea *sqlx.Tx, __obf_377e2b8c06e9f32b any) error

func Open(__obf_f7910eccb2b73f02, __obf_c52ed46ee91c4a91 string, __obf_6f9817dd4df76249 time.Duration) (*DB, error) {
	__obf_80cd76b5279c3275, __obf_a687171eb2d5acc2 := sqlx.Open(__obf_f7910eccb2b73f02, __obf_c52ed46ee91c4a91)
	if __obf_a687171eb2d5acc2 != nil {
		return nil, __obf_a687171eb2d5acc2
	}
	__obf_5b7d2bf1f9566b41 := &DB{
		DB:                     __obf_80cd76b5279c3275,
		__obf_6f9817dd4df76249: __obf_6f9817dd4df76249,
	}
	__obf_5b7d2bf1f9566b41.__obf_e893d21d21a267ea.New = func() any {
		return __obf_5b7d2bf1f9566b41.__obf_6fc8cccc5c427da4()
	}
	return __obf_5b7d2bf1f9566b41, nil
}

type DB struct {
	*sqlx.DB
	__obf_6f9817dd4df76249 time.Duration
	__obf_e893d21d21a267ea sync.Pool
}

func (__obf_80cd76b5279c3275 *DB) __obf_6fc8cccc5c427da4() *Context {
	return &Context{__obf_80cd76b5279c3275: __obf_80cd76b5279c3275}
}

// 获取一个`SQL`执行`Context`
func (__obf_80cd76b5279c3275 *DB) Acquire() *Context {
	// 无需加锁，sync.Pool本身是线程安全的
	__obf_c4d331869f0df57e := __obf_80cd76b5279c3275.__obf_e893d21d21a267ea.Get().(*Context)
	__obf_c4d331869f0df57e.__obf_fecf8f1a3b683094()
	return __obf_c4d331869f0df57e
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_80cd76b5279c3275 *DB) AcquireTx(__obf_5c86968691329dea *sqlx.Tx) *Context {
	__obf_c4d331869f0df57e := __obf_80cd76b5279c3275.Acquire()
	__obf_c4d331869f0df57e.__obf_5c86968691329dea = __obf_5c86968691329dea
	return __obf_c4d331869f0df57e
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_80cd76b5279c3275 *DB) WithTx(__obf_239ebfc86ee26fe0 FuncTx, __obf_377e2b8c06e9f32b any) (__obf_a687171eb2d5acc2 error) {
	var __obf_5c86968691329dea *sqlx.Tx
	__obf_5c86968691329dea, __obf_a687171eb2d5acc2 = __obf_80cd76b5279c3275.Beginx()
	if __obf_a687171eb2d5acc2 != nil {
		return
	}
	defer func() {
		if __obf_a687171eb2d5acc2 != nil && __obf_5c86968691329dea != nil {
			__obf_a687171eb2d5acc2 = __obf_5c86968691329dea.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_a687171eb2d5acc2 = __obf_239ebfc86ee26fe0(__obf_5c86968691329dea, __obf_377e2b8c06e9f32b); __obf_a687171eb2d5acc2 != nil {
		return
	}

	__obf_a687171eb2d5acc2 = __obf_5c86968691329dea.Commit()
	return
}

type Context struct {
	__obf_80cd76b5279c3275 *DB
	__obf_5c86968691329dea *sqlx.Tx //事务
	sql                    string
	__obf_304b905d846b5b7f string
	__obf_4246d77f13872cda []string
	__obf_5272f005f4c01a04 []string
	__obf_bdd2021b6884d84c string
	__obf_ef86e824b155c84d string
	__obf_ac20db0fa03ecd0f string
	__obf_715bd2918fbe1802 int64
	__obf_240d9433b3d41ad5 int64
	__obf_377e2b8c06e9f32b []any
	__obf_8ea98b90e52e2eaf bool //排他锁
	__obf_39c96a1aa2130363 bool //共享锁
}

func (__obf_c4d331869f0df57e *Context) Name(__obf_304b905d846b5b7f string) *Context {
	__obf_c4d331869f0df57e.__obf_304b905d846b5b7f = __obf_304b905d846b5b7f
	return __obf_c4d331869f0df57e
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_c4d331869f0df57e *Context) What(__obf_4246d77f13872cda []string) *Context {
	__obf_c4d331869f0df57e.__obf_4246d77f13872cda = __obf_4246d77f13872cda
	return __obf_c4d331869f0df57e
}

func (__obf_c4d331869f0df57e *Context) Where(__obf_bfb42fff03b52307 string, __obf_377e2b8c06e9f32b ...any) *Context {
	__obf_c4d331869f0df57e.__obf_5272f005f4c01a04 = append(__obf_c4d331869f0df57e.__obf_5272f005f4c01a04, __obf_bfb42fff03b52307)
	__obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b = append(__obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b, __obf_377e2b8c06e9f32b...)
	return __obf_c4d331869f0df57e
}

// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
func (__obf_c4d331869f0df57e *Context) WhereIn(__obf_5b84b0458a74fd86 string, __obf_377e2b8c06e9f32b []any) *Context {
	__obf_1dcc69b9f3380452 := len(__obf_377e2b8c06e9f32b)
	__obf_230598266cd0f21e := make([]string, __obf_1dcc69b9f3380452)
	for __obf_f4c5dc7b1f70e707 := 0; __obf_f4c5dc7b1f70e707 < __obf_1dcc69b9f3380452; __obf_f4c5dc7b1f70e707++ {
		__obf_230598266cd0f21e[__obf_f4c5dc7b1f70e707] = ParamMarker
	}
	__obf_5c4b3ffdc06f43c5 := fmt.Sprintf("%s in (%s)", __obf_5b84b0458a74fd86, __obf_3a442ba935c75114(__obf_230598266cd0f21e, SeqComma))
	return __obf_c4d331869f0df57e.Where(__obf_5c4b3ffdc06f43c5, __obf_377e2b8c06e9f32b...)
}

func (__obf_c4d331869f0df57e *Context) Order(__obf_bdd2021b6884d84c string) *Context {
	__obf_c4d331869f0df57e.__obf_bdd2021b6884d84c = __obf_bdd2021b6884d84c
	return __obf_c4d331869f0df57e
}

func (__obf_c4d331869f0df57e *Context) Limit(__obf_715bd2918fbe1802 int64) *Context {
	__obf_c4d331869f0df57e.__obf_715bd2918fbe1802 = __obf_715bd2918fbe1802
	return __obf_c4d331869f0df57e
}

func (__obf_c4d331869f0df57e *Context) Offset(__obf_240d9433b3d41ad5 int64) *Context {
	__obf_c4d331869f0df57e.__obf_240d9433b3d41ad5 = __obf_240d9433b3d41ad5
	return __obf_c4d331869f0df57e
}

func (__obf_c4d331869f0df57e *Context) Group(__obf_ef86e824b155c84d string) *Context {
	__obf_c4d331869f0df57e.__obf_ef86e824b155c84d = __obf_ef86e824b155c84d
	return __obf_c4d331869f0df57e
}

func (__obf_c4d331869f0df57e *Context) Having(__obf_ac20db0fa03ecd0f string, __obf_377e2b8c06e9f32b ...any) *Context {
	__obf_c4d331869f0df57e.__obf_ac20db0fa03ecd0f = __obf_ac20db0fa03ecd0f
	__obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b = append(__obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b, __obf_377e2b8c06e9f32b...)
	return __obf_c4d331869f0df57e
}

// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
func (__obf_c4d331869f0df57e *Context) LockX() *Context {
	__obf_c4d331869f0df57e.__obf_8ea98b90e52e2eaf = true
	return __obf_c4d331869f0df57e
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_c4d331869f0df57e *Context) LockS() *Context {
	__obf_c4d331869f0df57e.__obf_39c96a1aa2130363 = true
	return __obf_c4d331869f0df57e
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_c4d331869f0df57e *Context) FindMany(__obf_3875f0f94f8915fe any) error {
	return __obf_c4d331869f0df57e.__obf_0e14147e4a510038(__obf_3875f0f94f8915fe, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_c4d331869f0df57e *Context) FindById(__obf_3875f0f94f8915fe any) error {
	return __obf_c4d331869f0df57e.__obf_0e14147e4a510038(__obf_3875f0f94f8915fe, SelectTypeOne)
}

// 插入
func (__obf_c4d331869f0df57e *Context) Insert(__obf_545d51601b10d5e3 map[string]any) (sql.Result, error) {
	var (
		__obf_2cf417fca3803f4a []string
		__obf_86274bb72e83b9d5 []any
	)
	for __obf_388eb12efe355036, __obf_bbb98599fdbc54eb := range __obf_545d51601b10d5e3 {
		__obf_2cf417fca3803f4a = append(__obf_2cf417fca3803f4a, __obf_388eb12efe355036)
		__obf_86274bb72e83b9d5 = append(__obf_86274bb72e83b9d5, __obf_bbb98599fdbc54eb)
	}
	return __obf_c4d331869f0df57e.InsertBatch(__obf_2cf417fca3803f4a, __obf_86274bb72e83b9d5)
}

// 批量插入
func (__obf_c4d331869f0df57e *Context) InsertBatch(__obf_2cf417fca3803f4a []string, __obf_545d51601b10d5e3 ...[]any) (sql.Result, error) {
	var (
		__obf_86274bb72e83b9d5 []any
		__obf_26ccd77c1a037af0 []string
	)
	for _, __obf_e004a1fb016bbae8 := range __obf_545d51601b10d5e3 {
		__obf_230598266cd0f21e := make([]string, len(__obf_e004a1fb016bbae8))
		for __obf_f4c5dc7b1f70e707, __obf_bbb98599fdbc54eb := range __obf_e004a1fb016bbae8 {
			__obf_230598266cd0f21e[__obf_f4c5dc7b1f70e707] = ParamMarker
			__obf_86274bb72e83b9d5 = append(__obf_86274bb72e83b9d5, __obf_bbb98599fdbc54eb)
		}
		__obf_26ccd77c1a037af0 = append(__obf_26ccd77c1a037af0, fmt.Sprintf("(%s)", __obf_3a442ba935c75114(__obf_230598266cd0f21e, SeqComma)))
	}

	__obf_9cea37a5fd264d1a := fmt.Sprintf("insert into %s (%s) values %s", __obf_c4d331869f0df57e.__obf_304b905d846b5b7f, __obf_3a442ba935c75114(__obf_2cf417fca3803f4a, SeqComma), __obf_3a442ba935c75114(__obf_26ccd77c1a037af0, SeqComma))
	return __obf_c4d331869f0df57e.__obf_10fb1abddf7bdca3(__obf_9cea37a5fd264d1a, __obf_86274bb72e83b9d5...)
}

// 使用map更新
func (__obf_c4d331869f0df57e *Context) UpdateMap(__obf_377e2b8c06e9f32b map[string]any) (__obf_9558a7b9c4d137e9 int64, __obf_a687171eb2d5acc2 error) {
	var (
		__obf_86274bb72e83b9d5 []any
		__obf_e43867295457f1b8 []string
	)
	for __obf_388eb12efe355036, __obf_bbb98599fdbc54eb := range __obf_377e2b8c06e9f32b {
		__obf_86274bb72e83b9d5 = append(__obf_86274bb72e83b9d5, __obf_bbb98599fdbc54eb)
		__obf_e43867295457f1b8 = append(__obf_e43867295457f1b8, fmt.Sprintf("%s=%s", __obf_388eb12efe355036, ParamMarker))
	}
	__obf_f5cf8260f4b8b24b := __obf_3a442ba935c75114(__obf_e43867295457f1b8, SeqComma)
	__obf_9558a7b9c4d137e9, __obf_a687171eb2d5acc2 = __obf_c4d331869f0df57e.Update(__obf_f5cf8260f4b8b24b, __obf_86274bb72e83b9d5...)
	return
}

// 更新
func (__obf_c4d331869f0df57e *Context) Update(__obf_f5cf8260f4b8b24b string, __obf_377e2b8c06e9f32b ...any) (__obf_9558a7b9c4d137e9 int64, __obf_a687171eb2d5acc2 error) {
	__obf_793a6b96a81d9d08 := "update %s set %s %s"
	__obf_bfb42fff03b52307 := __obf_5a65ff6f06a1c346(__obf_c4d331869f0df57e.__obf_5272f005f4c01a04, Grouping)
	__obf_9cea37a5fd264d1a := fmt.Sprintf(__obf_793a6b96a81d9d08, __obf_c4d331869f0df57e.__obf_304b905d846b5b7f, __obf_f5cf8260f4b8b24b, __obf_bfb42fff03b52307)
	__obf_86274bb72e83b9d5 := append(__obf_377e2b8c06e9f32b, __obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b...)
	var __obf_098085715c2fedd0 sql.Result
	__obf_098085715c2fedd0, __obf_a687171eb2d5acc2 = __obf_c4d331869f0df57e.__obf_10fb1abddf7bdca3(__obf_9cea37a5fd264d1a, __obf_86274bb72e83b9d5...)
	if __obf_a687171eb2d5acc2 != nil {
		return
	}
	__obf_9558a7b9c4d137e9, __obf_a687171eb2d5acc2 = __obf_098085715c2fedd0.RowsAffected()
	return
}

// 删除
func (__obf_c4d331869f0df57e *Context) Delete() (__obf_9558a7b9c4d137e9 int64, __obf_a687171eb2d5acc2 error) {
	__obf_793a6b96a81d9d08 := "delete from %s %s"
	__obf_bfb42fff03b52307 := __obf_5a65ff6f06a1c346(__obf_c4d331869f0df57e.__obf_5272f005f4c01a04, Grouping)

	__obf_9cea37a5fd264d1a := fmt.Sprintf(__obf_793a6b96a81d9d08, __obf_c4d331869f0df57e.__obf_304b905d846b5b7f, __obf_bfb42fff03b52307)
	var __obf_098085715c2fedd0 sql.Result
	__obf_098085715c2fedd0, __obf_a687171eb2d5acc2 = __obf_c4d331869f0df57e.__obf_10fb1abddf7bdca3(__obf_9cea37a5fd264d1a, __obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b...)
	if __obf_a687171eb2d5acc2 != nil {
		return
	}
	__obf_9558a7b9c4d137e9, __obf_a687171eb2d5acc2 = __obf_098085715c2fedd0.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_c4d331869f0df57e *Context) Select(__obf_3875f0f94f8915fe any, sql string, __obf_377e2b8c06e9f32b ...any) error {
	__obf_c4d331869f0df57e.sql = sql
	__obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b = __obf_377e2b8c06e9f32b
	return __obf_c4d331869f0df57e.__obf_0e14147e4a510038(__obf_3875f0f94f8915fe, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_c4d331869f0df57e *Context) Get(__obf_3875f0f94f8915fe any, sql string, __obf_377e2b8c06e9f32b ...any) error {
	__obf_c4d331869f0df57e.sql = sql
	__obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b = __obf_377e2b8c06e9f32b
	return __obf_c4d331869f0df57e.__obf_0e14147e4a510038(__obf_3875f0f94f8915fe, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_c4d331869f0df57e *Context) Exec(sql string, __obf_377e2b8c06e9f32b ...any) (sql.Result, error) {
	return __obf_c4d331869f0df57e.__obf_10fb1abddf7bdca3(sql, __obf_377e2b8c06e9f32b...)
}

// 创建表
func (__obf_c4d331869f0df57e *Context) Create(sql string) (sql.Result, error) {
	return __obf_c4d331869f0df57e.__obf_10fb1abddf7bdca3(sql)
}

// 删除表
func (__obf_c4d331869f0df57e *Context) Drop() (sql.Result, error) {
	return __obf_c4d331869f0df57e.__obf_10fb1abddf7bdca3(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_c4d331869f0df57e.__obf_304b905d846b5b7f))
}

/////////////////////////private methods//////////////////////

// 重置Context
func (__obf_c4d331869f0df57e *Context) __obf_fecf8f1a3b683094() *Context {
	__obf_c4d331869f0df57e.sql = ""
	__obf_c4d331869f0df57e.__obf_304b905d846b5b7f = ""
	__obf_c4d331869f0df57e.__obf_4246d77f13872cda = []string{}
	__obf_c4d331869f0df57e.__obf_5272f005f4c01a04 = []string{}
	__obf_c4d331869f0df57e.__obf_bdd2021b6884d84c = ""
	__obf_c4d331869f0df57e.__obf_ef86e824b155c84d = ""
	__obf_c4d331869f0df57e.__obf_ac20db0fa03ecd0f = ""
	__obf_c4d331869f0df57e.__obf_715bd2918fbe1802 = 0
	__obf_c4d331869f0df57e.__obf_240d9433b3d41ad5 = 0
	__obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b = []any{}
	__obf_c4d331869f0df57e.__obf_5c86968691329dea = nil
	__obf_c4d331869f0df57e.__obf_39c96a1aa2130363 = false
	__obf_c4d331869f0df57e.__obf_8ea98b90e52e2eaf = false
	return __obf_c4d331869f0df57e
}

// 查询方法
func (__obf_c4d331869f0df57e *Context) __obf_0e14147e4a510038(__obf_3875f0f94f8915fe any, __obf_3dd1f0b3e128ba4b int) (__obf_a687171eb2d5acc2 error) {
	defer __obf_c4d331869f0df57e.__obf_80cd76b5279c3275.__obf_e893d21d21a267ea.Put(__obf_c4d331869f0df57e)
	__obf_145c7d4acc74d21c, __obf_db39b093f1d564c9 := context.WithTimeout(context.Background(), __obf_c4d331869f0df57e.__obf_80cd76b5279c3275.__obf_6f9817dd4df76249)
	defer __obf_db39b093f1d564c9()
	if __obf_c4d331869f0df57e.sql == "" {
		__obf_c4d331869f0df57e.sql = __obf_c4d331869f0df57e.__obf_7c3c525bddc20a45(__obf_3875f0f94f8915fe)
	}
	switch __obf_3dd1f0b3e128ba4b {
	case SelectTypeOne:
		if __obf_c4d331869f0df57e.__obf_5c86968691329dea != nil {
			__obf_a687171eb2d5acc2 = __obf_c4d331869f0df57e.__obf_5c86968691329dea.GetContext(__obf_145c7d4acc74d21c, __obf_3875f0f94f8915fe, __obf_c4d331869f0df57e.sql, __obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b...)
		} else {
			__obf_a687171eb2d5acc2 = __obf_c4d331869f0df57e.__obf_80cd76b5279c3275.GetContext(__obf_145c7d4acc74d21c, __obf_3875f0f94f8915fe, __obf_c4d331869f0df57e.sql, __obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b...)
		}
	case SelectTypeMany:
		if __obf_c4d331869f0df57e.__obf_5c86968691329dea != nil {
			__obf_a687171eb2d5acc2 = __obf_c4d331869f0df57e.__obf_5c86968691329dea.SelectContext(__obf_145c7d4acc74d21c, __obf_3875f0f94f8915fe, __obf_c4d331869f0df57e.sql, __obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b...)
		} else {
			__obf_a687171eb2d5acc2 = __obf_c4d331869f0df57e.__obf_80cd76b5279c3275.SelectContext(__obf_145c7d4acc74d21c, __obf_3875f0f94f8915fe, __obf_c4d331869f0df57e.sql, __obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_c4d331869f0df57e *Context) __obf_10fb1abddf7bdca3(__obf_9cea37a5fd264d1a string, __obf_377e2b8c06e9f32b ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b)
	defer __obf_c4d331869f0df57e.__obf_80cd76b5279c3275.__obf_e893d21d21a267ea.Put(__obf_c4d331869f0df57e)
	__obf_145c7d4acc74d21c, __obf_db39b093f1d564c9 := context.WithTimeout(context.Background(), __obf_c4d331869f0df57e.__obf_80cd76b5279c3275.__obf_6f9817dd4df76249)
	defer __obf_db39b093f1d564c9()

	var __obf_2e0b2258551d2d49 sqlx.ExecerContext
	if __obf_c4d331869f0df57e.__obf_5c86968691329dea != nil {
		__obf_2e0b2258551d2d49 = __obf_c4d331869f0df57e.__obf_5c86968691329dea
	} else {
		__obf_2e0b2258551d2d49 = __obf_c4d331869f0df57e.__obf_80cd76b5279c3275
	}
	return __obf_2e0b2258551d2d49.ExecContext(__obf_145c7d4acc74d21c, __obf_9cea37a5fd264d1a, __obf_377e2b8c06e9f32b...)
}

// select查询语句的拼接
func (__obf_c4d331869f0df57e *Context) __obf_7c3c525bddc20a45(__obf_3875f0f94f8915fe any) string {
	var __obf_ba0052095ab92536 []string
	__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, "select")
	if len(__obf_c4d331869f0df57e.__obf_4246d77f13872cda) != 0 {
		__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, __obf_3a442ba935c75114(__obf_c4d331869f0df57e.__obf_4246d77f13872cda, SeqComma))
	} else {
		// 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
		// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
		__obf_ac212f52f066884d := __obf_9bcae970346d82f7(__obf_3875f0f94f8915fe)
		if len(__obf_ac212f52f066884d) > 0 {
			__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, __obf_3a442ba935c75114(__obf_ac212f52f066884d, SeqComma))
		} else {
			__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, "*")
		}
	}
	__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, "from "+__obf_c4d331869f0df57e.__obf_304b905d846b5b7f)
	if len(__obf_c4d331869f0df57e.__obf_5272f005f4c01a04) != 0 {
		__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, __obf_5a65ff6f06a1c346(__obf_c4d331869f0df57e.__obf_5272f005f4c01a04, Grouping))
	}

	if __obf_c4d331869f0df57e.__obf_ef86e824b155c84d != "" {
		__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, "group by "+__obf_c4d331869f0df57e.__obf_ef86e824b155c84d)
	}

	if __obf_c4d331869f0df57e.__obf_ac20db0fa03ecd0f != "" {
		__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, "having "+__obf_c4d331869f0df57e.__obf_ac20db0fa03ecd0f)
	}

	if __obf_c4d331869f0df57e.__obf_bdd2021b6884d84c != "" {
		__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, "order by "+__obf_c4d331869f0df57e.__obf_bdd2021b6884d84c)
	}

	if __obf_c4d331869f0df57e.__obf_715bd2918fbe1802 != 0 {
		__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, fmt.Sprintf("limit %d, %d", __obf_c4d331869f0df57e.__obf_240d9433b3d41ad5, __obf_c4d331869f0df57e.__obf_715bd2918fbe1802))
	}
	if __obf_c4d331869f0df57e.__obf_39c96a1aa2130363 {
		__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, "lock in share mode")
	}
	if __obf_c4d331869f0df57e.__obf_8ea98b90e52e2eaf {
		__obf_ba0052095ab92536 = append(__obf_ba0052095ab92536, "for update")
	}
	sql := __obf_3a442ba935c75114(__obf_ba0052095ab92536, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_c4d331869f0df57e.__obf_377e2b8c06e9f32b)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_5a65ff6f06a1c346(__obf_5272f005f4c01a04 []string, __obf_c24b9f9cbb068f40 string) string {
	if len(__obf_5272f005f4c01a04) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_5272f005f4c01a04, __obf_c24b9f9cbb068f40))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_3a442ba935c75114(__obf_377e2b8c06e9f32b []string, __obf_4644098f8cd6e0cc string) string {
	return strings.Join(__obf_377e2b8c06e9f32b, __obf_4644098f8cd6e0cc)
}

// 解析对象中的 `db tag`
// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
func __obf_9bcae970346d82f7(__obf_3875f0f94f8915fe any) (__obf_2cf417fca3803f4a []string) {
	__obf_0826efe2d8c95bbc := reflect.ValueOf(__obf_3875f0f94f8915fe)
	__obf_31ccffbd695f1d65 := __obf_0826efe2d8c95bbc.Type().Elem()
	var __obf_32c94480748aa05a reflect.Type
	if __obf_31ccffbd695f1d65.Kind() == reflect.Slice {
		__obf_32c94480748aa05a = __obf_31ccffbd695f1d65.Elem()
	} else {
		__obf_32c94480748aa05a = __obf_31ccffbd695f1d65
	}
	for __obf_f4c5dc7b1f70e707 := 0; __obf_f4c5dc7b1f70e707 < __obf_32c94480748aa05a.NumField(); __obf_f4c5dc7b1f70e707++ {
		__obf_576f44cbcd018f9b := __obf_32c94480748aa05a.Field(__obf_f4c5dc7b1f70e707).Tag.Get(DBTag)
		if __obf_576f44cbcd018f9b != "" {
			__obf_2cf417fca3803f4a = append(__obf_2cf417fca3803f4a, __obf_576f44cbcd018f9b)
		}
	}
	return
}

func CleanSQLXMap(__obf_02135a184b6b9e5d map[string]any) map[string]any {
	for __obf_388eb12efe355036, __obf_bbb98599fdbc54eb := range __obf_02135a184b6b9e5d {
		switch __obf_9e8e83e850d24f6c := __obf_bbb98599fdbc54eb.(type) {
		case []byte:
			__obf_75b913ceae9096c5 := string(__obf_9e8e83e850d24f6c)

			// 尝试智能识别具体类型
			__obf_ced0de90b9076fad := strings.TrimSpace(__obf_75b913ceae9096c5)

			// bool
			if strings.EqualFold(__obf_ced0de90b9076fad, "true") || strings.EqualFold(__obf_ced0de90b9076fad, "false") {
				if __obf_369e593d4ae86732, __obf_a687171eb2d5acc2 := strconv.ParseBool(__obf_ced0de90b9076fad); __obf_a687171eb2d5acc2 == nil {
					__obf_02135a184b6b9e5d[__obf_388eb12efe355036] = __obf_369e593d4ae86732
					continue
				}
			}

			// int
			if __obf_f4c5dc7b1f70e707, __obf_a687171eb2d5acc2 := strconv.Atoi(__obf_ced0de90b9076fad); __obf_a687171eb2d5acc2 == nil {
				__obf_02135a184b6b9e5d[__obf_388eb12efe355036] = __obf_f4c5dc7b1f70e707
				continue
			}

			// float
			if __obf_57b980f12a7a4617, __obf_a687171eb2d5acc2 := strconv.ParseFloat(__obf_ced0de90b9076fad, 64); __obf_a687171eb2d5acc2 == nil {
				__obf_02135a184b6b9e5d[__obf_388eb12efe355036] = __obf_57b980f12a7a4617
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_ced0de90b9076fad, "{") && strings.HasSuffix(__obf_ced0de90b9076fad, "}")) ||
				(strings.HasPrefix(__obf_ced0de90b9076fad, "[") && strings.HasSuffix(__obf_ced0de90b9076fad, "]")) {
				var __obf_cc50b5a135e751e6 any
				if __obf_a687171eb2d5acc2 := json.Unmarshal(__obf_9e8e83e850d24f6c, &__obf_cc50b5a135e751e6); __obf_a687171eb2d5acc2 == nil {
					__obf_02135a184b6b9e5d[__obf_388eb12efe355036] = __obf_cc50b5a135e751e6
					continue
				}
			}

			// 默认保留为 string
			__obf_02135a184b6b9e5d[__obf_388eb12efe355036] = __obf_75b913ceae9096c5
		default:
			__obf_02135a184b6b9e5d[__obf_388eb12efe355036] = __obf_9e8e83e850d24f6c
		}
	}
	return __obf_02135a184b6b9e5d
}

func ToListMap(__obf_d1a5f8eb962a8a6f *sqlx.Rows) []map[string]any {
	var __obf_098085715c2fedd0 []map[string]any
	for __obf_d1a5f8eb962a8a6f.Next() {
		__obf_82318075915b2428 := make(map[string]any)
		if __obf_a687171eb2d5acc2 := __obf_d1a5f8eb962a8a6f.MapScan(__obf_82318075915b2428); __obf_a687171eb2d5acc2 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_a687171eb2d5acc2.Error()))
			continue
		}

		__obf_098085715c2fedd0 = append(__obf_098085715c2fedd0, CleanSQLXMap(__obf_82318075915b2428))
	}
	return __obf_098085715c2fedd0
}
