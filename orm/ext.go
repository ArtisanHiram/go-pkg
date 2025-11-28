package __obf_d726bb43e85f2dfc

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
type FuncTx func(__obf_5a0e0f8d9e4d8c5a *sqlx.Tx, __obf_9f04002b942a6c82 any) error

func Open(__obf_b13a38acdc4c6456, __obf_814566703bd71c4a string, __obf_86a431316c2a08bb time.Duration) (*DB, error) {
	__obf_a1d3ecfbcf6e018c, __obf_f5c7d5fde3143a52 := sqlx.Open(__obf_b13a38acdc4c6456, __obf_814566703bd71c4a)
	if __obf_f5c7d5fde3143a52 != nil {
		return nil, __obf_f5c7d5fde3143a52
	}
	__obf_4fceda9cabeb1c4f := &DB{
		DB:                     __obf_a1d3ecfbcf6e018c,
		__obf_86a431316c2a08bb: __obf_86a431316c2a08bb,
	}
	__obf_4fceda9cabeb1c4f.__obf_f7c8d1217439ddcb.New = func() any {
		return __obf_4fceda9cabeb1c4f.__obf_4365ebe7e3a1906e()
	}
	return __obf_4fceda9cabeb1c4f, nil
}

type DB struct {
	*sqlx.DB
	__obf_86a431316c2a08bb time.Duration
	__obf_f7c8d1217439ddcb sync.Pool
}

func (__obf_a1d3ecfbcf6e018c *DB) __obf_4365ebe7e3a1906e() *Context {
	return &Context{__obf_a1d3ecfbcf6e018c: __obf_a1d3ecfbcf6e018c}
}

// 获取一个`SQL`执行`Context`
func (__obf_a1d3ecfbcf6e018c *DB) Acquire() *Context {
	// 无需加锁，sync.Pool本身是线程安全的
	__obf_6ad564fb26485960 := __obf_a1d3ecfbcf6e018c.__obf_f7c8d1217439ddcb.Get().(*Context)
	__obf_6ad564fb26485960.__obf_434576f3bc8f78b5()
	return __obf_6ad564fb26485960
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_a1d3ecfbcf6e018c *DB) AcquireTx(__obf_5a0e0f8d9e4d8c5a *sqlx.Tx) *Context {
	__obf_6ad564fb26485960 := __obf_a1d3ecfbcf6e018c.Acquire()
	__obf_6ad564fb26485960.__obf_5a0e0f8d9e4d8c5a = __obf_5a0e0f8d9e4d8c5a
	return __obf_6ad564fb26485960
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_a1d3ecfbcf6e018c *DB) WithTx(__obf_7f624f13a8018b46 FuncTx, __obf_9f04002b942a6c82 any) (__obf_f5c7d5fde3143a52 error) {
	var __obf_5a0e0f8d9e4d8c5a *sqlx.Tx
	__obf_5a0e0f8d9e4d8c5a, __obf_f5c7d5fde3143a52 = __obf_a1d3ecfbcf6e018c.Beginx()
	if __obf_f5c7d5fde3143a52 != nil {
		return
	}
	defer func() {
		if __obf_f5c7d5fde3143a52 != nil && __obf_5a0e0f8d9e4d8c5a != nil {
			__obf_f5c7d5fde3143a52 = __obf_5a0e0f8d9e4d8c5a.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_f5c7d5fde3143a52 = __obf_7f624f13a8018b46(__obf_5a0e0f8d9e4d8c5a, __obf_9f04002b942a6c82); __obf_f5c7d5fde3143a52 != nil {
		return
	}

	__obf_f5c7d5fde3143a52 = __obf_5a0e0f8d9e4d8c5a.Commit()
	return
}

type Context struct {
	__obf_a1d3ecfbcf6e018c *DB
	__obf_5a0e0f8d9e4d8c5a *sqlx.Tx //事务
	sql                    string
	__obf_ba86bc56ac53a079 string
	__obf_d66d975edbed92c2 []string
	__obf_f61602f68808826e []string
	__obf_0ca96ac0748bcd5c string
	__obf_6ba7cc100bd4652e string
	__obf_e97f0303555e084f string
	__obf_227961e18089a5c0 int64
	__obf_156b3d3e86f94f51 int64
	__obf_9f04002b942a6c82 []any
	__obf_a01889b0267b3e2d bool //排他锁
	__obf_ed4d600785f949ec bool //共享锁
}

func (__obf_6ad564fb26485960 *Context) Name(__obf_ba86bc56ac53a079 string) *Context {
	__obf_6ad564fb26485960.__obf_ba86bc56ac53a079 = __obf_ba86bc56ac53a079
	return __obf_6ad564fb26485960
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_6ad564fb26485960 *Context) What(__obf_d66d975edbed92c2 []string) *Context {
	__obf_6ad564fb26485960.__obf_d66d975edbed92c2 = __obf_d66d975edbed92c2
	return __obf_6ad564fb26485960
}

func (__obf_6ad564fb26485960 *Context) Where(__obf_62fa9048d8479e91 string, __obf_9f04002b942a6c82 ...any) *Context {
	__obf_6ad564fb26485960.__obf_f61602f68808826e = append(__obf_6ad564fb26485960.__obf_f61602f68808826e, __obf_62fa9048d8479e91)
	__obf_6ad564fb26485960.__obf_9f04002b942a6c82 = append(__obf_6ad564fb26485960.__obf_9f04002b942a6c82, __obf_9f04002b942a6c82...)
	return __obf_6ad564fb26485960
}

// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
func (__obf_6ad564fb26485960 *Context) WhereIn(__obf_ef67ad2b7525fc31 string, __obf_9f04002b942a6c82 []any) *Context {
	__obf_41826dfbd0f4290a := len(__obf_9f04002b942a6c82)
	__obf_59b0318c61ce2409 := make([]string, __obf_41826dfbd0f4290a)
	for __obf_cbebdbf7ebca6a65 := 0; __obf_cbebdbf7ebca6a65 < __obf_41826dfbd0f4290a; __obf_cbebdbf7ebca6a65++ {
		__obf_59b0318c61ce2409[__obf_cbebdbf7ebca6a65] = ParamMarker
	}
	__obf_03568fbe66fb2483 := fmt.Sprintf("%s in (%s)", __obf_ef67ad2b7525fc31, __obf_3be8ae796ed41ca9(__obf_59b0318c61ce2409, SeqComma))
	return __obf_6ad564fb26485960.Where(__obf_03568fbe66fb2483, __obf_9f04002b942a6c82...)
}

func (__obf_6ad564fb26485960 *Context) Order(__obf_0ca96ac0748bcd5c string) *Context {
	__obf_6ad564fb26485960.__obf_0ca96ac0748bcd5c = __obf_0ca96ac0748bcd5c
	return __obf_6ad564fb26485960
}

func (__obf_6ad564fb26485960 *Context) Limit(__obf_227961e18089a5c0 int64) *Context {
	__obf_6ad564fb26485960.__obf_227961e18089a5c0 = __obf_227961e18089a5c0
	return __obf_6ad564fb26485960
}

func (__obf_6ad564fb26485960 *Context) Offset(__obf_156b3d3e86f94f51 int64) *Context {
	__obf_6ad564fb26485960.__obf_156b3d3e86f94f51 = __obf_156b3d3e86f94f51
	return __obf_6ad564fb26485960
}

func (__obf_6ad564fb26485960 *Context) Group(__obf_6ba7cc100bd4652e string) *Context {
	__obf_6ad564fb26485960.__obf_6ba7cc100bd4652e = __obf_6ba7cc100bd4652e
	return __obf_6ad564fb26485960
}

func (__obf_6ad564fb26485960 *Context) Having(__obf_e97f0303555e084f string, __obf_9f04002b942a6c82 ...any) *Context {
	__obf_6ad564fb26485960.__obf_e97f0303555e084f = __obf_e97f0303555e084f
	__obf_6ad564fb26485960.__obf_9f04002b942a6c82 = append(__obf_6ad564fb26485960.__obf_9f04002b942a6c82, __obf_9f04002b942a6c82...)
	return __obf_6ad564fb26485960
}

// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
func (__obf_6ad564fb26485960 *Context) LockX() *Context {
	__obf_6ad564fb26485960.__obf_a01889b0267b3e2d = true
	return __obf_6ad564fb26485960
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_6ad564fb26485960 *Context) LockS() *Context {
	__obf_6ad564fb26485960.__obf_ed4d600785f949ec = true
	return __obf_6ad564fb26485960
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_6ad564fb26485960 *Context) FindMany(__obf_bef1033cd918553a any) error {
	return __obf_6ad564fb26485960.__obf_5b9f31e22e15f84c(__obf_bef1033cd918553a, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_6ad564fb26485960 *Context) FindById(__obf_bef1033cd918553a any) error {
	return __obf_6ad564fb26485960.__obf_5b9f31e22e15f84c(__obf_bef1033cd918553a, SelectTypeOne)
}

// 插入
func (__obf_6ad564fb26485960 *Context) Insert(__obf_592e5973c519a4e9 map[string]any) (sql.Result, error) {
	var (
		__obf_d84f4fa139de2567 []string
		__obf_97a6dff53cb2815a []any
	)
	for __obf_3f7031f75dab588c, __obf_53bf414fffa45e91 := range __obf_592e5973c519a4e9 {
		__obf_d84f4fa139de2567 = append(__obf_d84f4fa139de2567, __obf_3f7031f75dab588c)
		__obf_97a6dff53cb2815a = append(__obf_97a6dff53cb2815a, __obf_53bf414fffa45e91)
	}
	return __obf_6ad564fb26485960.InsertBatch(__obf_d84f4fa139de2567, __obf_97a6dff53cb2815a)
}

// 批量插入
func (__obf_6ad564fb26485960 *Context) InsertBatch(__obf_d84f4fa139de2567 []string, __obf_592e5973c519a4e9 ...[]any) (sql.Result, error) {
	var (
		__obf_97a6dff53cb2815a []any
		__obf_5445c6dd9aee0d6c []string
	)
	for _, __obf_efb0c5ad0e0f8bad := range __obf_592e5973c519a4e9 {
		__obf_59b0318c61ce2409 := make([]string, len(__obf_efb0c5ad0e0f8bad))
		for __obf_cbebdbf7ebca6a65, __obf_53bf414fffa45e91 := range __obf_efb0c5ad0e0f8bad {
			__obf_59b0318c61ce2409[__obf_cbebdbf7ebca6a65] = ParamMarker
			__obf_97a6dff53cb2815a = append(__obf_97a6dff53cb2815a, __obf_53bf414fffa45e91)
		}
		__obf_5445c6dd9aee0d6c = append(__obf_5445c6dd9aee0d6c, fmt.Sprintf("(%s)", __obf_3be8ae796ed41ca9(__obf_59b0318c61ce2409, SeqComma)))
	}

	__obf_16dbf8405d413a99 := fmt.Sprintf("insert into %s (%s) values %s", __obf_6ad564fb26485960.__obf_ba86bc56ac53a079, __obf_3be8ae796ed41ca9(__obf_d84f4fa139de2567, SeqComma), __obf_3be8ae796ed41ca9(__obf_5445c6dd9aee0d6c, SeqComma))
	return __obf_6ad564fb26485960.__obf_ae3c302743cfed79(__obf_16dbf8405d413a99, __obf_97a6dff53cb2815a...)
}

// 使用map更新
func (__obf_6ad564fb26485960 *Context) UpdateMap(__obf_9f04002b942a6c82 map[string]any) (__obf_9f54832607302719 int64, __obf_f5c7d5fde3143a52 error) {
	var (
		__obf_97a6dff53cb2815a []any
		__obf_fe9395904379ca31 []string
	)
	for __obf_3f7031f75dab588c, __obf_53bf414fffa45e91 := range __obf_9f04002b942a6c82 {
		__obf_97a6dff53cb2815a = append(__obf_97a6dff53cb2815a, __obf_53bf414fffa45e91)
		__obf_fe9395904379ca31 = append(__obf_fe9395904379ca31, fmt.Sprintf("%s=%s", __obf_3f7031f75dab588c, ParamMarker))
	}
	__obf_ff7bb7ce78ec8b77 := __obf_3be8ae796ed41ca9(__obf_fe9395904379ca31, SeqComma)
	__obf_9f54832607302719, __obf_f5c7d5fde3143a52 = __obf_6ad564fb26485960.Update(__obf_ff7bb7ce78ec8b77, __obf_97a6dff53cb2815a...)
	return
}

// 更新
func (__obf_6ad564fb26485960 *Context) Update(__obf_ff7bb7ce78ec8b77 string, __obf_9f04002b942a6c82 ...any) (__obf_9f54832607302719 int64, __obf_f5c7d5fde3143a52 error) {
	__obf_05fce86926e5327d := "update %s set %s %s"
	__obf_62fa9048d8479e91 := __obf_7c756a01eb80767e(__obf_6ad564fb26485960.__obf_f61602f68808826e, Grouping)
	__obf_16dbf8405d413a99 := fmt.Sprintf(__obf_05fce86926e5327d, __obf_6ad564fb26485960.__obf_ba86bc56ac53a079, __obf_ff7bb7ce78ec8b77, __obf_62fa9048d8479e91)
	__obf_97a6dff53cb2815a := append(__obf_9f04002b942a6c82, __obf_6ad564fb26485960.__obf_9f04002b942a6c82...)
	var __obf_3357f1ed7884ec23 sql.Result
	__obf_3357f1ed7884ec23, __obf_f5c7d5fde3143a52 = __obf_6ad564fb26485960.__obf_ae3c302743cfed79(__obf_16dbf8405d413a99, __obf_97a6dff53cb2815a...)
	if __obf_f5c7d5fde3143a52 != nil {
		return
	}
	__obf_9f54832607302719, __obf_f5c7d5fde3143a52 = __obf_3357f1ed7884ec23.RowsAffected()
	return
}

// 删除
func (__obf_6ad564fb26485960 *Context) Delete() (__obf_9f54832607302719 int64, __obf_f5c7d5fde3143a52 error) {
	__obf_05fce86926e5327d := "delete from %s %s"
	__obf_62fa9048d8479e91 := __obf_7c756a01eb80767e(__obf_6ad564fb26485960.__obf_f61602f68808826e, Grouping)

	__obf_16dbf8405d413a99 := fmt.Sprintf(__obf_05fce86926e5327d, __obf_6ad564fb26485960.__obf_ba86bc56ac53a079, __obf_62fa9048d8479e91)
	var __obf_3357f1ed7884ec23 sql.Result
	__obf_3357f1ed7884ec23, __obf_f5c7d5fde3143a52 = __obf_6ad564fb26485960.__obf_ae3c302743cfed79(__obf_16dbf8405d413a99, __obf_6ad564fb26485960.__obf_9f04002b942a6c82...)
	if __obf_f5c7d5fde3143a52 != nil {
		return
	}
	__obf_9f54832607302719, __obf_f5c7d5fde3143a52 = __obf_3357f1ed7884ec23.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_6ad564fb26485960 *Context) Select(__obf_bef1033cd918553a any, sql string, __obf_9f04002b942a6c82 ...any) error {
	__obf_6ad564fb26485960.sql = sql
	__obf_6ad564fb26485960.__obf_9f04002b942a6c82 = __obf_9f04002b942a6c82
	return __obf_6ad564fb26485960.__obf_5b9f31e22e15f84c(__obf_bef1033cd918553a, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_6ad564fb26485960 *Context) Get(__obf_bef1033cd918553a any, sql string, __obf_9f04002b942a6c82 ...any) error {
	__obf_6ad564fb26485960.sql = sql
	__obf_6ad564fb26485960.__obf_9f04002b942a6c82 = __obf_9f04002b942a6c82
	return __obf_6ad564fb26485960.__obf_5b9f31e22e15f84c(__obf_bef1033cd918553a, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_6ad564fb26485960 *Context) Exec(sql string, __obf_9f04002b942a6c82 ...any) (sql.Result, error) {
	return __obf_6ad564fb26485960.__obf_ae3c302743cfed79(sql, __obf_9f04002b942a6c82...)
}

// 创建表
func (__obf_6ad564fb26485960 *Context) Create(sql string) (sql.Result, error) {
	return __obf_6ad564fb26485960.__obf_ae3c302743cfed79(sql)
}

// 删除表
func (__obf_6ad564fb26485960 *Context) Drop() (sql.Result, error) {
	return __obf_6ad564fb26485960.__obf_ae3c302743cfed79(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_6ad564fb26485960.__obf_ba86bc56ac53a079))
}

/////////////////////////private methods//////////////////////

// 重置Context
func (__obf_6ad564fb26485960 *Context) __obf_434576f3bc8f78b5() *Context {
	__obf_6ad564fb26485960.sql = ""
	__obf_6ad564fb26485960.__obf_ba86bc56ac53a079 = ""
	__obf_6ad564fb26485960.__obf_d66d975edbed92c2 = []string{}
	__obf_6ad564fb26485960.__obf_f61602f68808826e = []string{}
	__obf_6ad564fb26485960.__obf_0ca96ac0748bcd5c = ""
	__obf_6ad564fb26485960.__obf_6ba7cc100bd4652e = ""
	__obf_6ad564fb26485960.__obf_e97f0303555e084f = ""
	__obf_6ad564fb26485960.__obf_227961e18089a5c0 = 0
	__obf_6ad564fb26485960.__obf_156b3d3e86f94f51 = 0
	__obf_6ad564fb26485960.__obf_9f04002b942a6c82 = []any{}
	__obf_6ad564fb26485960.__obf_5a0e0f8d9e4d8c5a = nil
	__obf_6ad564fb26485960.__obf_ed4d600785f949ec = false
	__obf_6ad564fb26485960.__obf_a01889b0267b3e2d = false
	return __obf_6ad564fb26485960
}

// 查询方法
func (__obf_6ad564fb26485960 *Context) __obf_5b9f31e22e15f84c(__obf_bef1033cd918553a any, __obf_aff9bde975e2c3a8 int) (__obf_f5c7d5fde3143a52 error) {
	defer __obf_6ad564fb26485960.__obf_a1d3ecfbcf6e018c.__obf_f7c8d1217439ddcb.Put(__obf_6ad564fb26485960)
	__obf_510315e46c820672, __obf_eef5f0c83dd1b75c := context.WithTimeout(context.Background(), __obf_6ad564fb26485960.__obf_a1d3ecfbcf6e018c.__obf_86a431316c2a08bb)
	defer __obf_eef5f0c83dd1b75c()
	if __obf_6ad564fb26485960.sql == "" {
		__obf_6ad564fb26485960.sql = __obf_6ad564fb26485960.__obf_8f69b866e1b041de(__obf_bef1033cd918553a)
	}
	switch __obf_aff9bde975e2c3a8 {
	case SelectTypeOne:
		if __obf_6ad564fb26485960.__obf_5a0e0f8d9e4d8c5a != nil {
			__obf_f5c7d5fde3143a52 = __obf_6ad564fb26485960.__obf_5a0e0f8d9e4d8c5a.GetContext(__obf_510315e46c820672, __obf_bef1033cd918553a, __obf_6ad564fb26485960.sql, __obf_6ad564fb26485960.__obf_9f04002b942a6c82...)
		} else {
			__obf_f5c7d5fde3143a52 = __obf_6ad564fb26485960.__obf_a1d3ecfbcf6e018c.GetContext(__obf_510315e46c820672, __obf_bef1033cd918553a, __obf_6ad564fb26485960.sql, __obf_6ad564fb26485960.__obf_9f04002b942a6c82...)
		}
	case SelectTypeMany:
		if __obf_6ad564fb26485960.__obf_5a0e0f8d9e4d8c5a != nil {
			__obf_f5c7d5fde3143a52 = __obf_6ad564fb26485960.__obf_5a0e0f8d9e4d8c5a.SelectContext(__obf_510315e46c820672, __obf_bef1033cd918553a, __obf_6ad564fb26485960.sql, __obf_6ad564fb26485960.__obf_9f04002b942a6c82...)
		} else {
			__obf_f5c7d5fde3143a52 = __obf_6ad564fb26485960.__obf_a1d3ecfbcf6e018c.SelectContext(__obf_510315e46c820672, __obf_bef1033cd918553a, __obf_6ad564fb26485960.sql, __obf_6ad564fb26485960.__obf_9f04002b942a6c82...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_6ad564fb26485960 *Context) __obf_ae3c302743cfed79(__obf_16dbf8405d413a99 string, __obf_9f04002b942a6c82 ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_16dbf8405d413a99, __obf_9f04002b942a6c82)
	defer __obf_6ad564fb26485960.__obf_a1d3ecfbcf6e018c.__obf_f7c8d1217439ddcb.Put(__obf_6ad564fb26485960)
	__obf_510315e46c820672, __obf_eef5f0c83dd1b75c := context.WithTimeout(context.Background(), __obf_6ad564fb26485960.__obf_a1d3ecfbcf6e018c.__obf_86a431316c2a08bb)
	defer __obf_eef5f0c83dd1b75c()

	var __obf_0c36519de712be70 sqlx.ExecerContext
	if __obf_6ad564fb26485960.__obf_5a0e0f8d9e4d8c5a != nil {
		__obf_0c36519de712be70 = __obf_6ad564fb26485960.__obf_5a0e0f8d9e4d8c5a
	} else {
		__obf_0c36519de712be70 = __obf_6ad564fb26485960.__obf_a1d3ecfbcf6e018c
	}
	return __obf_0c36519de712be70.ExecContext(__obf_510315e46c820672, __obf_16dbf8405d413a99, __obf_9f04002b942a6c82...)
}

// select查询语句的拼接
func (__obf_6ad564fb26485960 *Context) __obf_8f69b866e1b041de(__obf_bef1033cd918553a any) string {
	var __obf_949038cdfbc174ac []string
	__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, "select")
	if len(__obf_6ad564fb26485960.__obf_d66d975edbed92c2) != 0 {
		__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, __obf_3be8ae796ed41ca9(__obf_6ad564fb26485960.__obf_d66d975edbed92c2, SeqComma))
	} else {
		// 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
		// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
		__obf_c0aa4c84673baa94 := __obf_d8909b96abd573e1(__obf_bef1033cd918553a)
		if len(__obf_c0aa4c84673baa94) > 0 {
			__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, __obf_3be8ae796ed41ca9(__obf_c0aa4c84673baa94, SeqComma))
		} else {
			__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, "*")
		}
	}
	__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, "from "+__obf_6ad564fb26485960.__obf_ba86bc56ac53a079)
	if len(__obf_6ad564fb26485960.__obf_f61602f68808826e) != 0 {
		__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, __obf_7c756a01eb80767e(__obf_6ad564fb26485960.__obf_f61602f68808826e, Grouping))
	}

	if __obf_6ad564fb26485960.__obf_6ba7cc100bd4652e != "" {
		__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, "group by "+__obf_6ad564fb26485960.__obf_6ba7cc100bd4652e)
	}

	if __obf_6ad564fb26485960.__obf_e97f0303555e084f != "" {
		__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, "having "+__obf_6ad564fb26485960.__obf_e97f0303555e084f)
	}

	if __obf_6ad564fb26485960.__obf_0ca96ac0748bcd5c != "" {
		__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, "order by "+__obf_6ad564fb26485960.__obf_0ca96ac0748bcd5c)
	}

	if __obf_6ad564fb26485960.__obf_227961e18089a5c0 != 0 {
		__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, fmt.Sprintf("limit %d, %d", __obf_6ad564fb26485960.__obf_156b3d3e86f94f51, __obf_6ad564fb26485960.__obf_227961e18089a5c0))
	}
	if __obf_6ad564fb26485960.__obf_ed4d600785f949ec {
		__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, "lock in share mode")
	}
	if __obf_6ad564fb26485960.__obf_a01889b0267b3e2d {
		__obf_949038cdfbc174ac = append(__obf_949038cdfbc174ac, "for update")
	}
	sql := __obf_3be8ae796ed41ca9(__obf_949038cdfbc174ac, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_6ad564fb26485960.__obf_9f04002b942a6c82)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_7c756a01eb80767e(__obf_f61602f68808826e []string, __obf_f8d5a93ab58b6574 string) string {
	if len(__obf_f61602f68808826e) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_f61602f68808826e, __obf_f8d5a93ab58b6574))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_3be8ae796ed41ca9(__obf_9f04002b942a6c82 []string, __obf_608fc9324cceb524 string) string {
	return strings.Join(__obf_9f04002b942a6c82, __obf_608fc9324cceb524)
}

// 解析对象中的 `db tag`
// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
func __obf_d8909b96abd573e1(__obf_bef1033cd918553a any) (__obf_d84f4fa139de2567 []string) {
	__obf_14f652161c66608a := reflect.ValueOf(__obf_bef1033cd918553a)
	__obf_dd4144155edd98e5 := __obf_14f652161c66608a.Type().Elem()
	var __obf_c9231ca39f190b5e reflect.Type
	if __obf_dd4144155edd98e5.Kind() == reflect.Slice {
		__obf_c9231ca39f190b5e = __obf_dd4144155edd98e5.Elem()
	} else {
		__obf_c9231ca39f190b5e = __obf_dd4144155edd98e5
	}
	for __obf_cbebdbf7ebca6a65 := 0; __obf_cbebdbf7ebca6a65 < __obf_c9231ca39f190b5e.NumField(); __obf_cbebdbf7ebca6a65++ {
		__obf_9f57bc1c737b05d1 := __obf_c9231ca39f190b5e.Field(__obf_cbebdbf7ebca6a65).Tag.Get(DBTag)
		if __obf_9f57bc1c737b05d1 != "" {
			__obf_d84f4fa139de2567 = append(__obf_d84f4fa139de2567, __obf_9f57bc1c737b05d1)
		}
	}
	return
}

func CleanSQLXMap(__obf_f2df2e8729071d08 map[string]any) map[string]any {
	for __obf_3f7031f75dab588c, __obf_53bf414fffa45e91 := range __obf_f2df2e8729071d08 {
		switch __obf_145a2e7072310087 := __obf_53bf414fffa45e91.(type) {
		case []byte:
			__obf_3f6b3a80155a822c := string(__obf_145a2e7072310087)

			// 尝试智能识别具体类型
			__obf_6b2dd2a6d824aea5 := strings.TrimSpace(__obf_3f6b3a80155a822c)

			// bool
			if strings.EqualFold(__obf_6b2dd2a6d824aea5, "true") || strings.EqualFold(__obf_6b2dd2a6d824aea5, "false") {
				if __obf_afe0e796e0b12d8f, __obf_f5c7d5fde3143a52 := strconv.ParseBool(__obf_6b2dd2a6d824aea5); __obf_f5c7d5fde3143a52 == nil {
					__obf_f2df2e8729071d08[__obf_3f7031f75dab588c] = __obf_afe0e796e0b12d8f
					continue
				}
			}

			// int
			if __obf_cbebdbf7ebca6a65, __obf_f5c7d5fde3143a52 := strconv.Atoi(__obf_6b2dd2a6d824aea5); __obf_f5c7d5fde3143a52 == nil {
				__obf_f2df2e8729071d08[__obf_3f7031f75dab588c] = __obf_cbebdbf7ebca6a65
				continue
			}

			// float
			if __obf_b0142d337517f48f, __obf_f5c7d5fde3143a52 := strconv.ParseFloat(__obf_6b2dd2a6d824aea5, 64); __obf_f5c7d5fde3143a52 == nil {
				__obf_f2df2e8729071d08[__obf_3f7031f75dab588c] = __obf_b0142d337517f48f
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_6b2dd2a6d824aea5, "{") && strings.HasSuffix(__obf_6b2dd2a6d824aea5, "}")) ||
				(strings.HasPrefix(__obf_6b2dd2a6d824aea5, "[") && strings.HasSuffix(__obf_6b2dd2a6d824aea5, "]")) {
				var __obf_af17a6e3080cb22b any
				if __obf_f5c7d5fde3143a52 := json.Unmarshal(__obf_145a2e7072310087, &__obf_af17a6e3080cb22b); __obf_f5c7d5fde3143a52 == nil {
					__obf_f2df2e8729071d08[__obf_3f7031f75dab588c] = __obf_af17a6e3080cb22b
					continue
				}
			}

			// 默认保留为 string
			__obf_f2df2e8729071d08[__obf_3f7031f75dab588c] = __obf_3f6b3a80155a822c
		default:
			__obf_f2df2e8729071d08[__obf_3f7031f75dab588c] = __obf_145a2e7072310087
		}
	}
	return __obf_f2df2e8729071d08
}

func ToListMap(__obf_af10750736e4963e *sqlx.Rows) []map[string]any {
	var __obf_3357f1ed7884ec23 []map[string]any
	for __obf_af10750736e4963e.Next() {
		__obf_25565284ad74e3ee := make(map[string]any)
		if __obf_f5c7d5fde3143a52 := __obf_af10750736e4963e.MapScan(__obf_25565284ad74e3ee); __obf_f5c7d5fde3143a52 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_f5c7d5fde3143a52.Error()))
			continue
		}

		__obf_3357f1ed7884ec23 = append(__obf_3357f1ed7884ec23, CleanSQLXMap(__obf_25565284ad74e3ee))
	}
	return __obf_3357f1ed7884ec23
}
