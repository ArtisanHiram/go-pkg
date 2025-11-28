package __obf_66da9ed7fbc1ec9e

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
type FuncTx func(__obf_181488bf1961013b *sqlx.Tx, __obf_66ab003bba546d1d any) error

func Open(__obf_9e03da570445de8c, __obf_b3d2e1eb0cdd272b string, __obf_f6d4ed25fadbb031 time.Duration) (*DB, error) {
	__obf_2e13600f19148084, __obf_1229ad25944902c2 := sqlx.Open(__obf_9e03da570445de8c, __obf_b3d2e1eb0cdd272b)
	if __obf_1229ad25944902c2 != nil {
		return nil, __obf_1229ad25944902c2
	}
	__obf_1d423fa3821811d8 := &DB{
		DB:                     __obf_2e13600f19148084,
		__obf_f6d4ed25fadbb031: __obf_f6d4ed25fadbb031,
	}
	__obf_1d423fa3821811d8.__obf_cd09a97e77526b71.New = func() any {
		return __obf_1d423fa3821811d8.__obf_c3fbd30db852bdaf()
	}
	return __obf_1d423fa3821811d8, nil
}

type DB struct {
	*sqlx.DB
	__obf_f6d4ed25fadbb031 time.Duration
	__obf_cd09a97e77526b71 sync.Pool
}

func (__obf_2e13600f19148084 *DB) __obf_c3fbd30db852bdaf() *Context {
	return &Context{__obf_2e13600f19148084: __obf_2e13600f19148084}
}

// 获取一个`SQL`执行`Context`
func (__obf_2e13600f19148084 *DB) Acquire() *Context {
	// 无需加锁，sync.Pool本身是线程安全的
	__obf_def785cbde1dd5f0 := __obf_2e13600f19148084.__obf_cd09a97e77526b71.Get().(*Context)
	__obf_def785cbde1dd5f0.__obf_90518a6ab8d87fe8()
	return __obf_def785cbde1dd5f0
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_2e13600f19148084 *DB) AcquireTx(__obf_181488bf1961013b *sqlx.Tx) *Context {
	__obf_def785cbde1dd5f0 := __obf_2e13600f19148084.Acquire()
	__obf_def785cbde1dd5f0.__obf_181488bf1961013b = __obf_181488bf1961013b
	return __obf_def785cbde1dd5f0
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_2e13600f19148084 *DB) WithTx(__obf_9c7c78fc0a38ac8d FuncTx, __obf_66ab003bba546d1d any) (__obf_1229ad25944902c2 error) {
	var __obf_181488bf1961013b *sqlx.Tx
	__obf_181488bf1961013b, __obf_1229ad25944902c2 = __obf_2e13600f19148084.Beginx()
	if __obf_1229ad25944902c2 != nil {
		return
	}
	defer func() {
		if __obf_1229ad25944902c2 != nil && __obf_181488bf1961013b != nil {
			__obf_1229ad25944902c2 = __obf_181488bf1961013b.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_1229ad25944902c2 = __obf_9c7c78fc0a38ac8d(__obf_181488bf1961013b, __obf_66ab003bba546d1d); __obf_1229ad25944902c2 != nil {
		return
	}

	__obf_1229ad25944902c2 = __obf_181488bf1961013b.Commit()
	return
}

type Context struct {
	__obf_2e13600f19148084 *DB
	__obf_181488bf1961013b *sqlx.Tx //事务
	sql                    string
	__obf_f477196bed539112 string
	__obf_1421765dbeb61db5 []string
	__obf_097c32c9294d6428 []string
	__obf_5e44ba8d82e98713 string
	__obf_f53196fb47cd5ca4 string
	__obf_62eeef7f2a3ed712 string
	__obf_6164bcde7de4e49c int64
	__obf_d807a2811f7abbfc int64
	__obf_66ab003bba546d1d []any
	__obf_605a9fc0aebcb352 bool //排他锁
	__obf_74e2706a3d008fb5 bool //共享锁
}

func (__obf_def785cbde1dd5f0 *Context) Name(__obf_f477196bed539112 string) *Context {
	__obf_def785cbde1dd5f0.__obf_f477196bed539112 = __obf_f477196bed539112
	return __obf_def785cbde1dd5f0
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_def785cbde1dd5f0 *Context) What(__obf_1421765dbeb61db5 []string) *Context {
	__obf_def785cbde1dd5f0.__obf_1421765dbeb61db5 = __obf_1421765dbeb61db5
	return __obf_def785cbde1dd5f0
}

func (__obf_def785cbde1dd5f0 *Context) Where(__obf_f818ec1012d81c56 string, __obf_66ab003bba546d1d ...any) *Context {
	__obf_def785cbde1dd5f0.__obf_097c32c9294d6428 = append(__obf_def785cbde1dd5f0.__obf_097c32c9294d6428, __obf_f818ec1012d81c56)
	__obf_def785cbde1dd5f0.__obf_66ab003bba546d1d = append(__obf_def785cbde1dd5f0.__obf_66ab003bba546d1d, __obf_66ab003bba546d1d...)
	return __obf_def785cbde1dd5f0
}

// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
func (__obf_def785cbde1dd5f0 *Context) WhereIn(__obf_7527be360ce64e08 string, __obf_66ab003bba546d1d []any) *Context {
	__obf_a964a53ba0b068c6 := len(__obf_66ab003bba546d1d)
	__obf_827735f3ab9e435a := make([]string, __obf_a964a53ba0b068c6)
	for __obf_4ceb93679168ba44 := 0; __obf_4ceb93679168ba44 < __obf_a964a53ba0b068c6; __obf_4ceb93679168ba44++ {
		__obf_827735f3ab9e435a[__obf_4ceb93679168ba44] = ParamMarker
	}
	__obf_8c7d5e2a1cb7efb0 := fmt.Sprintf("%s in (%s)", __obf_7527be360ce64e08, __obf_d5759b3288e8c167(__obf_827735f3ab9e435a, SeqComma))
	return __obf_def785cbde1dd5f0.Where(__obf_8c7d5e2a1cb7efb0, __obf_66ab003bba546d1d...)
}

func (__obf_def785cbde1dd5f0 *Context) Order(__obf_5e44ba8d82e98713 string) *Context {
	__obf_def785cbde1dd5f0.__obf_5e44ba8d82e98713 = __obf_5e44ba8d82e98713
	return __obf_def785cbde1dd5f0
}

func (__obf_def785cbde1dd5f0 *Context) Limit(__obf_6164bcde7de4e49c int64) *Context {
	__obf_def785cbde1dd5f0.__obf_6164bcde7de4e49c = __obf_6164bcde7de4e49c
	return __obf_def785cbde1dd5f0
}

func (__obf_def785cbde1dd5f0 *Context) Offset(__obf_d807a2811f7abbfc int64) *Context {
	__obf_def785cbde1dd5f0.__obf_d807a2811f7abbfc = __obf_d807a2811f7abbfc
	return __obf_def785cbde1dd5f0
}

func (__obf_def785cbde1dd5f0 *Context) Group(__obf_f53196fb47cd5ca4 string) *Context {
	__obf_def785cbde1dd5f0.__obf_f53196fb47cd5ca4 = __obf_f53196fb47cd5ca4
	return __obf_def785cbde1dd5f0
}

func (__obf_def785cbde1dd5f0 *Context) Having(__obf_62eeef7f2a3ed712 string, __obf_66ab003bba546d1d ...any) *Context {
	__obf_def785cbde1dd5f0.__obf_62eeef7f2a3ed712 = __obf_62eeef7f2a3ed712
	__obf_def785cbde1dd5f0.__obf_66ab003bba546d1d = append(__obf_def785cbde1dd5f0.__obf_66ab003bba546d1d, __obf_66ab003bba546d1d...)
	return __obf_def785cbde1dd5f0
}

// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
func (__obf_def785cbde1dd5f0 *Context) LockX() *Context {
	__obf_def785cbde1dd5f0.__obf_605a9fc0aebcb352 = true
	return __obf_def785cbde1dd5f0
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_def785cbde1dd5f0 *Context) LockS() *Context {
	__obf_def785cbde1dd5f0.__obf_74e2706a3d008fb5 = true
	return __obf_def785cbde1dd5f0
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_def785cbde1dd5f0 *Context) FindMany(__obf_888dae6fde0def77 any) error {
	return __obf_def785cbde1dd5f0.__obf_7762dc532a04150f(__obf_888dae6fde0def77, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_def785cbde1dd5f0 *Context) FindById(__obf_888dae6fde0def77 any) error {
	return __obf_def785cbde1dd5f0.__obf_7762dc532a04150f(__obf_888dae6fde0def77, SelectTypeOne)
}

// 插入
func (__obf_def785cbde1dd5f0 *Context) Insert(__obf_7542478dd197efb5 map[string]any) (sql.Result, error) {
	var (
		__obf_61e6b3ecbb290190 []string
		__obf_0f506337693ba0e6 []any
	)
	for __obf_22dbe2e4f410cd44, __obf_ac35cf98df641116 := range __obf_7542478dd197efb5 {
		__obf_61e6b3ecbb290190 = append(__obf_61e6b3ecbb290190, __obf_22dbe2e4f410cd44)
		__obf_0f506337693ba0e6 = append(__obf_0f506337693ba0e6, __obf_ac35cf98df641116)
	}
	return __obf_def785cbde1dd5f0.InsertBatch(__obf_61e6b3ecbb290190, __obf_0f506337693ba0e6)
}

// 批量插入
func (__obf_def785cbde1dd5f0 *Context) InsertBatch(__obf_61e6b3ecbb290190 []string, __obf_7542478dd197efb5 ...[]any) (sql.Result, error) {
	var (
		__obf_0f506337693ba0e6 []any
		__obf_544cb58d8e7c4f15 []string
	)
	for _, __obf_3ec97e021e4afb4b := range __obf_7542478dd197efb5 {
		__obf_827735f3ab9e435a := make([]string, len(__obf_3ec97e021e4afb4b))
		for __obf_4ceb93679168ba44, __obf_ac35cf98df641116 := range __obf_3ec97e021e4afb4b {
			__obf_827735f3ab9e435a[__obf_4ceb93679168ba44] = ParamMarker
			__obf_0f506337693ba0e6 = append(__obf_0f506337693ba0e6, __obf_ac35cf98df641116)
		}
		__obf_544cb58d8e7c4f15 = append(__obf_544cb58d8e7c4f15, fmt.Sprintf("(%s)", __obf_d5759b3288e8c167(__obf_827735f3ab9e435a, SeqComma)))
	}

	__obf_8c21f68e9672c4f9 := fmt.Sprintf("insert into %s (%s) values %s", __obf_def785cbde1dd5f0.__obf_f477196bed539112, __obf_d5759b3288e8c167(__obf_61e6b3ecbb290190, SeqComma), __obf_d5759b3288e8c167(__obf_544cb58d8e7c4f15, SeqComma))
	return __obf_def785cbde1dd5f0.__obf_5512d8558b9b1498(__obf_8c21f68e9672c4f9, __obf_0f506337693ba0e6...)
}

// 使用map更新
func (__obf_def785cbde1dd5f0 *Context) UpdateMap(__obf_66ab003bba546d1d map[string]any) (__obf_a21601ecbe05773b int64, __obf_1229ad25944902c2 error) {
	var (
		__obf_0f506337693ba0e6 []any
		__obf_babc3866b4d26829 []string
	)
	for __obf_22dbe2e4f410cd44, __obf_ac35cf98df641116 := range __obf_66ab003bba546d1d {
		__obf_0f506337693ba0e6 = append(__obf_0f506337693ba0e6, __obf_ac35cf98df641116)
		__obf_babc3866b4d26829 = append(__obf_babc3866b4d26829, fmt.Sprintf("%s=%s", __obf_22dbe2e4f410cd44, ParamMarker))
	}
	__obf_d6fa7f60b49d0ca3 := __obf_d5759b3288e8c167(__obf_babc3866b4d26829, SeqComma)
	__obf_a21601ecbe05773b, __obf_1229ad25944902c2 = __obf_def785cbde1dd5f0.Update(__obf_d6fa7f60b49d0ca3, __obf_0f506337693ba0e6...)
	return
}

// 更新
func (__obf_def785cbde1dd5f0 *Context) Update(__obf_d6fa7f60b49d0ca3 string, __obf_66ab003bba546d1d ...any) (__obf_a21601ecbe05773b int64, __obf_1229ad25944902c2 error) {
	__obf_c59032a111760ff4 := "update %s set %s %s"
	__obf_f818ec1012d81c56 := __obf_314aea75e959ddeb(__obf_def785cbde1dd5f0.__obf_097c32c9294d6428, Grouping)
	__obf_8c21f68e9672c4f9 := fmt.Sprintf(__obf_c59032a111760ff4, __obf_def785cbde1dd5f0.__obf_f477196bed539112, __obf_d6fa7f60b49d0ca3, __obf_f818ec1012d81c56)
	__obf_0f506337693ba0e6 := append(__obf_66ab003bba546d1d, __obf_def785cbde1dd5f0.__obf_66ab003bba546d1d...)
	var __obf_8f096a0267f35515 sql.Result
	__obf_8f096a0267f35515, __obf_1229ad25944902c2 = __obf_def785cbde1dd5f0.__obf_5512d8558b9b1498(__obf_8c21f68e9672c4f9, __obf_0f506337693ba0e6...)
	if __obf_1229ad25944902c2 != nil {
		return
	}
	__obf_a21601ecbe05773b, __obf_1229ad25944902c2 = __obf_8f096a0267f35515.RowsAffected()
	return
}

// 删除
func (__obf_def785cbde1dd5f0 *Context) Delete() (__obf_a21601ecbe05773b int64, __obf_1229ad25944902c2 error) {
	__obf_c59032a111760ff4 := "delete from %s %s"
	__obf_f818ec1012d81c56 := __obf_314aea75e959ddeb(__obf_def785cbde1dd5f0.__obf_097c32c9294d6428, Grouping)

	__obf_8c21f68e9672c4f9 := fmt.Sprintf(__obf_c59032a111760ff4, __obf_def785cbde1dd5f0.__obf_f477196bed539112, __obf_f818ec1012d81c56)
	var __obf_8f096a0267f35515 sql.Result
	__obf_8f096a0267f35515, __obf_1229ad25944902c2 = __obf_def785cbde1dd5f0.__obf_5512d8558b9b1498(__obf_8c21f68e9672c4f9, __obf_def785cbde1dd5f0.__obf_66ab003bba546d1d...)
	if __obf_1229ad25944902c2 != nil {
		return
	}
	__obf_a21601ecbe05773b, __obf_1229ad25944902c2 = __obf_8f096a0267f35515.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_def785cbde1dd5f0 *Context) Select(__obf_888dae6fde0def77 any, sql string, __obf_66ab003bba546d1d ...any) error {
	__obf_def785cbde1dd5f0.sql = sql
	__obf_def785cbde1dd5f0.__obf_66ab003bba546d1d = __obf_66ab003bba546d1d
	return __obf_def785cbde1dd5f0.__obf_7762dc532a04150f(__obf_888dae6fde0def77, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_def785cbde1dd5f0 *Context) Get(__obf_888dae6fde0def77 any, sql string, __obf_66ab003bba546d1d ...any) error {
	__obf_def785cbde1dd5f0.sql = sql
	__obf_def785cbde1dd5f0.__obf_66ab003bba546d1d = __obf_66ab003bba546d1d
	return __obf_def785cbde1dd5f0.__obf_7762dc532a04150f(__obf_888dae6fde0def77, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_def785cbde1dd5f0 *Context) Exec(sql string, __obf_66ab003bba546d1d ...any) (sql.Result, error) {
	return __obf_def785cbde1dd5f0.__obf_5512d8558b9b1498(sql, __obf_66ab003bba546d1d...)
}

// 创建表
func (__obf_def785cbde1dd5f0 *Context) Create(sql string) (sql.Result, error) {
	return __obf_def785cbde1dd5f0.__obf_5512d8558b9b1498(sql)
}

// 删除表
func (__obf_def785cbde1dd5f0 *Context) Drop() (sql.Result, error) {
	return __obf_def785cbde1dd5f0.__obf_5512d8558b9b1498(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_def785cbde1dd5f0.__obf_f477196bed539112))
}

/////////////////////////private methods//////////////////////

// 重置Context
func (__obf_def785cbde1dd5f0 *Context) __obf_90518a6ab8d87fe8() *Context {
	__obf_def785cbde1dd5f0.sql = ""
	__obf_def785cbde1dd5f0.__obf_f477196bed539112 = ""
	__obf_def785cbde1dd5f0.__obf_1421765dbeb61db5 = []string{}
	__obf_def785cbde1dd5f0.__obf_097c32c9294d6428 = []string{}
	__obf_def785cbde1dd5f0.__obf_5e44ba8d82e98713 = ""
	__obf_def785cbde1dd5f0.__obf_f53196fb47cd5ca4 = ""
	__obf_def785cbde1dd5f0.__obf_62eeef7f2a3ed712 = ""
	__obf_def785cbde1dd5f0.__obf_6164bcde7de4e49c = 0
	__obf_def785cbde1dd5f0.__obf_d807a2811f7abbfc = 0
	__obf_def785cbde1dd5f0.__obf_66ab003bba546d1d = []any{}
	__obf_def785cbde1dd5f0.__obf_181488bf1961013b = nil
	__obf_def785cbde1dd5f0.__obf_74e2706a3d008fb5 = false
	__obf_def785cbde1dd5f0.__obf_605a9fc0aebcb352 = false
	return __obf_def785cbde1dd5f0
}

// 查询方法
func (__obf_def785cbde1dd5f0 *Context) __obf_7762dc532a04150f(__obf_888dae6fde0def77 any, __obf_fbca303bff678977 int) (__obf_1229ad25944902c2 error) {
	defer __obf_def785cbde1dd5f0.__obf_2e13600f19148084.__obf_cd09a97e77526b71.Put(__obf_def785cbde1dd5f0)
	__obf_b904a54e22a0591d, __obf_bf4713cdd156b365 := context.WithTimeout(context.Background(), __obf_def785cbde1dd5f0.__obf_2e13600f19148084.__obf_f6d4ed25fadbb031)
	defer __obf_bf4713cdd156b365()
	if __obf_def785cbde1dd5f0.sql == "" {
		__obf_def785cbde1dd5f0.sql = __obf_def785cbde1dd5f0.__obf_f14a9c2c3e275e87(__obf_888dae6fde0def77)
	}
	switch __obf_fbca303bff678977 {
	case SelectTypeOne:
		if __obf_def785cbde1dd5f0.__obf_181488bf1961013b != nil {
			__obf_1229ad25944902c2 = __obf_def785cbde1dd5f0.__obf_181488bf1961013b.GetContext(__obf_b904a54e22a0591d, __obf_888dae6fde0def77, __obf_def785cbde1dd5f0.sql, __obf_def785cbde1dd5f0.__obf_66ab003bba546d1d...)
		} else {
			__obf_1229ad25944902c2 = __obf_def785cbde1dd5f0.__obf_2e13600f19148084.GetContext(__obf_b904a54e22a0591d, __obf_888dae6fde0def77, __obf_def785cbde1dd5f0.sql, __obf_def785cbde1dd5f0.__obf_66ab003bba546d1d...)
		}
	case SelectTypeMany:
		if __obf_def785cbde1dd5f0.__obf_181488bf1961013b != nil {
			__obf_1229ad25944902c2 = __obf_def785cbde1dd5f0.__obf_181488bf1961013b.SelectContext(__obf_b904a54e22a0591d, __obf_888dae6fde0def77, __obf_def785cbde1dd5f0.sql, __obf_def785cbde1dd5f0.__obf_66ab003bba546d1d...)
		} else {
			__obf_1229ad25944902c2 = __obf_def785cbde1dd5f0.__obf_2e13600f19148084.SelectContext(__obf_b904a54e22a0591d, __obf_888dae6fde0def77, __obf_def785cbde1dd5f0.sql, __obf_def785cbde1dd5f0.__obf_66ab003bba546d1d...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_def785cbde1dd5f0 *Context) __obf_5512d8558b9b1498(__obf_8c21f68e9672c4f9 string, __obf_66ab003bba546d1d ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d)
	defer __obf_def785cbde1dd5f0.__obf_2e13600f19148084.__obf_cd09a97e77526b71.Put(__obf_def785cbde1dd5f0)
	__obf_b904a54e22a0591d, __obf_bf4713cdd156b365 := context.WithTimeout(context.Background(), __obf_def785cbde1dd5f0.__obf_2e13600f19148084.__obf_f6d4ed25fadbb031)
	defer __obf_bf4713cdd156b365()

	var __obf_4ed23a295d93b84a sqlx.ExecerContext
	if __obf_def785cbde1dd5f0.__obf_181488bf1961013b != nil {
		__obf_4ed23a295d93b84a = __obf_def785cbde1dd5f0.__obf_181488bf1961013b
	} else {
		__obf_4ed23a295d93b84a = __obf_def785cbde1dd5f0.__obf_2e13600f19148084
	}
	return __obf_4ed23a295d93b84a.ExecContext(__obf_b904a54e22a0591d, __obf_8c21f68e9672c4f9, __obf_66ab003bba546d1d...)
}

// select查询语句的拼接
func (__obf_def785cbde1dd5f0 *Context) __obf_f14a9c2c3e275e87(__obf_888dae6fde0def77 any) string {
	var __obf_2946fd754152c36e []string
	__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, "select")
	if len(__obf_def785cbde1dd5f0.__obf_1421765dbeb61db5) != 0 {
		__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, __obf_d5759b3288e8c167(__obf_def785cbde1dd5f0.__obf_1421765dbeb61db5, SeqComma))
	} else {
		// 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
		// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
		__obf_0f189bb6d7d1ca5c := __obf_bf495dcf9bbf3e4b(__obf_888dae6fde0def77)
		if len(__obf_0f189bb6d7d1ca5c) > 0 {
			__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, __obf_d5759b3288e8c167(__obf_0f189bb6d7d1ca5c, SeqComma))
		} else {
			__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, "*")
		}
	}
	__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, "from "+__obf_def785cbde1dd5f0.__obf_f477196bed539112)
	if len(__obf_def785cbde1dd5f0.__obf_097c32c9294d6428) != 0 {
		__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, __obf_314aea75e959ddeb(__obf_def785cbde1dd5f0.__obf_097c32c9294d6428, Grouping))
	}

	if __obf_def785cbde1dd5f0.__obf_f53196fb47cd5ca4 != "" {
		__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, "group by "+__obf_def785cbde1dd5f0.__obf_f53196fb47cd5ca4)
	}

	if __obf_def785cbde1dd5f0.__obf_62eeef7f2a3ed712 != "" {
		__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, "having "+__obf_def785cbde1dd5f0.__obf_62eeef7f2a3ed712)
	}

	if __obf_def785cbde1dd5f0.__obf_5e44ba8d82e98713 != "" {
		__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, "order by "+__obf_def785cbde1dd5f0.__obf_5e44ba8d82e98713)
	}

	if __obf_def785cbde1dd5f0.__obf_6164bcde7de4e49c != 0 {
		__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, fmt.Sprintf("limit %d, %d", __obf_def785cbde1dd5f0.__obf_d807a2811f7abbfc, __obf_def785cbde1dd5f0.__obf_6164bcde7de4e49c))
	}
	if __obf_def785cbde1dd5f0.__obf_74e2706a3d008fb5 {
		__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, "lock in share mode")
	}
	if __obf_def785cbde1dd5f0.__obf_605a9fc0aebcb352 {
		__obf_2946fd754152c36e = append(__obf_2946fd754152c36e, "for update")
	}
	sql := __obf_d5759b3288e8c167(__obf_2946fd754152c36e, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_def785cbde1dd5f0.__obf_66ab003bba546d1d)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_314aea75e959ddeb(__obf_097c32c9294d6428 []string, __obf_e38deba7eb0a0e9f string) string {
	if len(__obf_097c32c9294d6428) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_097c32c9294d6428, __obf_e38deba7eb0a0e9f))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_d5759b3288e8c167(__obf_66ab003bba546d1d []string, __obf_b4e2a90785c77a27 string) string {
	return strings.Join(__obf_66ab003bba546d1d, __obf_b4e2a90785c77a27)
}

// 解析对象中的 `db tag`
// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
func __obf_bf495dcf9bbf3e4b(__obf_888dae6fde0def77 any) (__obf_61e6b3ecbb290190 []string) {
	__obf_830db7a4c207dec8 := reflect.ValueOf(__obf_888dae6fde0def77)
	__obf_b7e3d0c96c05802f := __obf_830db7a4c207dec8.Type().Elem()
	var __obf_978cab3718b5cd5d reflect.Type
	if __obf_b7e3d0c96c05802f.Kind() == reflect.Slice {
		__obf_978cab3718b5cd5d = __obf_b7e3d0c96c05802f.Elem()
	} else {
		__obf_978cab3718b5cd5d = __obf_b7e3d0c96c05802f
	}
	for __obf_4ceb93679168ba44 := 0; __obf_4ceb93679168ba44 < __obf_978cab3718b5cd5d.NumField(); __obf_4ceb93679168ba44++ {
		__obf_b1a624be540b7a6b := __obf_978cab3718b5cd5d.Field(__obf_4ceb93679168ba44).Tag.Get(DBTag)
		if __obf_b1a624be540b7a6b != "" {
			__obf_61e6b3ecbb290190 = append(__obf_61e6b3ecbb290190, __obf_b1a624be540b7a6b)
		}
	}
	return
}

func CleanSQLXMap(__obf_90b9e55b34956f13 map[string]any) map[string]any {
	for __obf_22dbe2e4f410cd44, __obf_ac35cf98df641116 := range __obf_90b9e55b34956f13 {
		switch __obf_bc4d09a84501c296 := __obf_ac35cf98df641116.(type) {
		case []byte:
			__obf_a3ec10d2be1cb29c := string(__obf_bc4d09a84501c296)

			// 尝试智能识别具体类型
			__obf_5bc0a86d054bc94e := strings.TrimSpace(__obf_a3ec10d2be1cb29c)

			// bool
			if strings.EqualFold(__obf_5bc0a86d054bc94e, "true") || strings.EqualFold(__obf_5bc0a86d054bc94e, "false") {
				if __obf_b150671df762e3f7, __obf_1229ad25944902c2 := strconv.ParseBool(__obf_5bc0a86d054bc94e); __obf_1229ad25944902c2 == nil {
					__obf_90b9e55b34956f13[__obf_22dbe2e4f410cd44] = __obf_b150671df762e3f7
					continue
				}
			}

			// int
			if __obf_4ceb93679168ba44, __obf_1229ad25944902c2 := strconv.Atoi(__obf_5bc0a86d054bc94e); __obf_1229ad25944902c2 == nil {
				__obf_90b9e55b34956f13[__obf_22dbe2e4f410cd44] = __obf_4ceb93679168ba44
				continue
			}

			// float
			if __obf_bb16074987cbe782, __obf_1229ad25944902c2 := strconv.ParseFloat(__obf_5bc0a86d054bc94e, 64); __obf_1229ad25944902c2 == nil {
				__obf_90b9e55b34956f13[__obf_22dbe2e4f410cd44] = __obf_bb16074987cbe782
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_5bc0a86d054bc94e, "{") && strings.HasSuffix(__obf_5bc0a86d054bc94e, "}")) ||
				(strings.HasPrefix(__obf_5bc0a86d054bc94e, "[") && strings.HasSuffix(__obf_5bc0a86d054bc94e, "]")) {
				var __obf_2ebc4f322d0ea57b any
				if __obf_1229ad25944902c2 := json.Unmarshal(__obf_bc4d09a84501c296, &__obf_2ebc4f322d0ea57b); __obf_1229ad25944902c2 == nil {
					__obf_90b9e55b34956f13[__obf_22dbe2e4f410cd44] = __obf_2ebc4f322d0ea57b
					continue
				}
			}

			// 默认保留为 string
			__obf_90b9e55b34956f13[__obf_22dbe2e4f410cd44] = __obf_a3ec10d2be1cb29c
		default:
			__obf_90b9e55b34956f13[__obf_22dbe2e4f410cd44] = __obf_bc4d09a84501c296
		}
	}
	return __obf_90b9e55b34956f13
}

func ToListMap(__obf_ce506c59e72daf0b *sqlx.Rows) []map[string]any {
	var __obf_8f096a0267f35515 []map[string]any
	for __obf_ce506c59e72daf0b.Next() {
		__obf_69d25d4c49e4c4af := make(map[string]any)
		if __obf_1229ad25944902c2 := __obf_ce506c59e72daf0b.MapScan(__obf_69d25d4c49e4c4af); __obf_1229ad25944902c2 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_1229ad25944902c2.Error()))
			continue
		}

		__obf_8f096a0267f35515 = append(__obf_8f096a0267f35515, CleanSQLXMap(__obf_69d25d4c49e4c4af))
	}
	return __obf_8f096a0267f35515
}
