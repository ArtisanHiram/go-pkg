package __obf_c1f2c3d265c98f25

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
type FuncTx func(__obf_b210b34065b441df *sqlx.Tx, __obf_c2322d6ab4e7aa7f any) error

func Open(__obf_1eb74688b4ba3e57, __obf_195d2a6cfa835cda string, __obf_cc30e8549e22914a time.Duration) (*DB, error) {
	__obf_2e89f553fec85399, __obf_ab5192d9e0103059 := sqlx.Open(__obf_1eb74688b4ba3e57, __obf_195d2a6cfa835cda)
	if __obf_ab5192d9e0103059 != nil {
		return nil, __obf_ab5192d9e0103059
	}
	__obf_d806e474f8bd0d56 := &DB{
		DB: __obf_2e89f553fec85399, __obf_cc30e8549e22914a: __obf_cc30e8549e22914a,
	}
	__obf_d806e474f8bd0d56.__obf_ab4cebcf1acecdf0.
		New = func() any {
		return __obf_d806e474f8bd0d56.__obf_dbf347f70b029ffe()
	}
	return __obf_d806e474f8bd0d56, nil
}

type DB struct {
	*sqlx.DB
	__obf_cc30e8549e22914a time.Duration
	__obf_ab4cebcf1acecdf0 sync.Pool
}

func (__obf_2e89f553fec85399 *DB) __obf_dbf347f70b029ffe() *Context {
	return &Context{__obf_2e89f553fec85399: __obf_2e89f553fec85399}
}

// 获取一个`SQL`执行`Context`
func (__obf_2e89f553fec85399 *DB) Acquire() *Context {
	__obf_03d12f1ecab719e1 := // 无需加锁，sync.Pool本身是线程安全的
		__obf_2e89f553fec85399.__obf_ab4cebcf1acecdf0.Get().(*Context)
	__obf_03d12f1ecab719e1.__obf_68e8df12e48b03e0()
	return __obf_03d12f1ecab719e1
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_2e89f553fec85399 *DB) AcquireTx(__obf_b210b34065b441df *sqlx.Tx) *Context {
	__obf_03d12f1ecab719e1 := __obf_2e89f553fec85399.Acquire()
	__obf_03d12f1ecab719e1.__obf_b210b34065b441df = __obf_b210b34065b441df
	return __obf_03d12f1ecab719e1
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_2e89f553fec85399 *DB) WithTx(__obf_5b197e25ab264988 FuncTx, __obf_c2322d6ab4e7aa7f any) (__obf_ab5192d9e0103059 error) {
	var __obf_b210b34065b441df *sqlx.Tx
	__obf_b210b34065b441df, __obf_ab5192d9e0103059 = __obf_2e89f553fec85399.Beginx()
	if __obf_ab5192d9e0103059 != nil {
		return
	}
	defer func() {
		if __obf_ab5192d9e0103059 != nil && __obf_b210b34065b441df != nil {
			__obf_ab5192d9e0103059 = __obf_b210b34065b441df.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_ab5192d9e0103059 = __obf_5b197e25ab264988(__obf_b210b34065b441df, __obf_c2322d6ab4e7aa7f); __obf_ab5192d9e0103059 != nil {
		return
	}
	__obf_ab5192d9e0103059 = __obf_b210b34065b441df.Commit()
	return
}

type Context struct {
	__obf_2e89f553fec85399 *DB
	__obf_b210b34065b441df *sqlx.Tx //事务
	sql                    string
	__obf_1ebef8f2a981f0f4 string
	__obf_739f909e439feddb []string
	__obf_e8b5a700a5355b56 []string
	__obf_fa570685b7794b0d string
	__obf_bb5c00d71cf92f10 string
	__obf_36ce77795c55c871 string
	__obf_48eaaf4227b33422 int64
	__obf_d3025b9543b73c0b int64
	__obf_c2322d6ab4e7aa7f []any
	__obf_5a5f66d6b091294d bool
	__obf_6a22a5c046530e35 bool //排他锁
	//共享锁
}

func (__obf_03d12f1ecab719e1 *Context) Name(__obf_1ebef8f2a981f0f4 string) *Context {
	__obf_03d12f1ecab719e1.__obf_1ebef8f2a981f0f4 = __obf_1ebef8f2a981f0f4
	return __obf_03d12f1ecab719e1
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_03d12f1ecab719e1 *Context) What(__obf_739f909e439feddb []string) *Context {
	__obf_03d12f1ecab719e1.__obf_739f909e439feddb = __obf_739f909e439feddb
	return __obf_03d12f1ecab719e1
}

func (__obf_03d12f1ecab719e1 *Context) Where(__obf_b1004fa5302ae44f string, __obf_c2322d6ab4e7aa7f ...any) *Context {
	__obf_03d12f1ecab719e1.__obf_e8b5a700a5355b56 = append(__obf_03d12f1ecab719e1.__obf_e8b5a700a5355b56, __obf_b1004fa5302ae44f)
	__obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f = append(__obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f,

		// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
		__obf_c2322d6ab4e7aa7f...)
	return __obf_03d12f1ecab719e1
}

func (__obf_03d12f1ecab719e1 *Context) WhereIn(__obf_fbae619a1e6618fb string, __obf_c2322d6ab4e7aa7f []any) *Context {
	__obf_40db83c90826dbbf := len(__obf_c2322d6ab4e7aa7f)
	__obf_7ca3d9ff14eb867e := make([]string, __obf_40db83c90826dbbf)
	for __obf_739dac239fb8287e := 0; __obf_739dac239fb8287e < __obf_40db83c90826dbbf; __obf_739dac239fb8287e++ {
		__obf_7ca3d9ff14eb867e[__obf_739dac239fb8287e] = ParamMarker
	}
	__obf_562c8458e88f084e := fmt.Sprintf("%s in (%s)", __obf_fbae619a1e6618fb, __obf_6c1ae2aa470953c2(__obf_7ca3d9ff14eb867e, SeqComma))
	return __obf_03d12f1ecab719e1.Where(__obf_562c8458e88f084e, __obf_c2322d6ab4e7aa7f...)
}

func (__obf_03d12f1ecab719e1 *Context) Order(__obf_fa570685b7794b0d string) *Context {
	__obf_03d12f1ecab719e1.__obf_fa570685b7794b0d = __obf_fa570685b7794b0d
	return __obf_03d12f1ecab719e1
}

func (__obf_03d12f1ecab719e1 *Context) Limit(__obf_48eaaf4227b33422 int64) *Context {
	__obf_03d12f1ecab719e1.__obf_48eaaf4227b33422 = __obf_48eaaf4227b33422
	return __obf_03d12f1ecab719e1
}

func (__obf_03d12f1ecab719e1 *Context) Offset(__obf_d3025b9543b73c0b int64) *Context {
	__obf_03d12f1ecab719e1.__obf_d3025b9543b73c0b = __obf_d3025b9543b73c0b
	return __obf_03d12f1ecab719e1
}

func (__obf_03d12f1ecab719e1 *Context) Group(__obf_bb5c00d71cf92f10 string) *Context {
	__obf_03d12f1ecab719e1.__obf_bb5c00d71cf92f10 = __obf_bb5c00d71cf92f10
	return __obf_03d12f1ecab719e1
}

func (__obf_03d12f1ecab719e1 *Context) Having(__obf_36ce77795c55c871 string, __obf_c2322d6ab4e7aa7f ...any) *Context {
	__obf_03d12f1ecab719e1.__obf_36ce77795c55c871 = __obf_36ce77795c55c871
	__obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f = append(__obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f,

		// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
		__obf_c2322d6ab4e7aa7f...)
	return __obf_03d12f1ecab719e1
}

func (__obf_03d12f1ecab719e1 *Context) LockX() *Context {
	__obf_03d12f1ecab719e1.__obf_5a5f66d6b091294d = true
	return __obf_03d12f1ecab719e1
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_03d12f1ecab719e1 *Context) LockS() *Context {
	__obf_03d12f1ecab719e1.__obf_6a22a5c046530e35 = true
	return __obf_03d12f1ecab719e1
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_03d12f1ecab719e1 *Context) FindMany(__obf_f67dd9e9f9b7d2f9 any) error {
	return __obf_03d12f1ecab719e1.__obf_2fe92164135afc29(__obf_f67dd9e9f9b7d2f9, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_03d12f1ecab719e1 *Context) FindById(__obf_f67dd9e9f9b7d2f9 any) error {
	return __obf_03d12f1ecab719e1.__obf_2fe92164135afc29(__obf_f67dd9e9f9b7d2f9, SelectTypeOne)
}

// 插入
func (__obf_03d12f1ecab719e1 *Context) Insert(__obf_df6ac6b94da14379 map[string]any) (sql.Result, error) {
	var (
		__obf_75d4ce8525b58df8 []string
		__obf_05cde18c0f67f8c5 []any
	)
	for __obf_02cffdb2b73b0b57, __obf_8ec2f1c7b43e366d := range __obf_df6ac6b94da14379 {
		__obf_75d4ce8525b58df8 = append(__obf_75d4ce8525b58df8, __obf_02cffdb2b73b0b57)
		__obf_05cde18c0f67f8c5 = append(__obf_05cde18c0f67f8c5, __obf_8ec2f1c7b43e366d)
	}
	return __obf_03d12f1ecab719e1.InsertBatch(__obf_75d4ce8525b58df8,

		// 批量插入
		__obf_05cde18c0f67f8c5)
}

func (__obf_03d12f1ecab719e1 *Context) InsertBatch(__obf_75d4ce8525b58df8 []string, __obf_df6ac6b94da14379 ...[]any) (sql.Result, error) {
	var (
		__obf_05cde18c0f67f8c5 []any
		__obf_20565d58e11e8173 []string
	)
	for _, __obf_defcc67f3ad4f964 := range __obf_df6ac6b94da14379 {
		__obf_7ca3d9ff14eb867e := make([]string, len(__obf_defcc67f3ad4f964))
		for __obf_739dac239fb8287e, __obf_8ec2f1c7b43e366d := range __obf_defcc67f3ad4f964 {
			__obf_7ca3d9ff14eb867e[__obf_739dac239fb8287e] = ParamMarker
			__obf_05cde18c0f67f8c5 = append(__obf_05cde18c0f67f8c5, __obf_8ec2f1c7b43e366d)
		}
		__obf_20565d58e11e8173 = append(__obf_20565d58e11e8173, fmt.Sprintf("(%s)", __obf_6c1ae2aa470953c2(__obf_7ca3d9ff14eb867e, SeqComma)))
	}
	__obf_8d780c51a7e1aba8 := fmt.Sprintf("insert into %s (%s) values %s", __obf_03d12f1ecab719e1.__obf_1ebef8f2a981f0f4, __obf_6c1ae2aa470953c2(__obf_75d4ce8525b58df8, SeqComma), __obf_6c1ae2aa470953c2(__obf_20565d58e11e8173, SeqComma))
	return __obf_03d12f1ecab719e1.__obf_9f33bf964bd83b55(__obf_8d780c51a7e1aba8,

		// 使用map更新
		__obf_05cde18c0f67f8c5...)
}

func (__obf_03d12f1ecab719e1 *Context) UpdateMap(__obf_c2322d6ab4e7aa7f map[string]any) (__obf_5758290ee6e3c6af int64, __obf_ab5192d9e0103059 error) {
	var (
		__obf_05cde18c0f67f8c5 []any
		__obf_ee118c327c04d6e0 []string
	)
	for __obf_02cffdb2b73b0b57, __obf_8ec2f1c7b43e366d := range __obf_c2322d6ab4e7aa7f {
		__obf_05cde18c0f67f8c5 = append(__obf_05cde18c0f67f8c5, __obf_8ec2f1c7b43e366d)
		__obf_ee118c327c04d6e0 = append(__obf_ee118c327c04d6e0, fmt.Sprintf("%s=%s", __obf_02cffdb2b73b0b57, ParamMarker))
	}
	__obf_19a5599cb1cfd681 := __obf_6c1ae2aa470953c2(__obf_ee118c327c04d6e0, SeqComma)
	__obf_5758290ee6e3c6af, __obf_ab5192d9e0103059 = __obf_03d12f1ecab719e1.Update(__obf_19a5599cb1cfd681, __obf_05cde18c0f67f8c5...)
	return
}

// 更新
func (__obf_03d12f1ecab719e1 *Context) Update(__obf_19a5599cb1cfd681 string, __obf_c2322d6ab4e7aa7f ...any) (__obf_5758290ee6e3c6af int64, __obf_ab5192d9e0103059 error) {
	__obf_6cd274412975cdee := "update %s set %s %s"
	__obf_b1004fa5302ae44f := __obf_b9a1feb76538f56d(__obf_03d12f1ecab719e1.__obf_e8b5a700a5355b56, Grouping)
	__obf_8d780c51a7e1aba8 := fmt.Sprintf(__obf_6cd274412975cdee, __obf_03d12f1ecab719e1.__obf_1ebef8f2a981f0f4, __obf_19a5599cb1cfd681, __obf_b1004fa5302ae44f)
	__obf_05cde18c0f67f8c5 := append(__obf_c2322d6ab4e7aa7f, __obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f...)
	var __obf_fda7a11c3287a559 sql.Result
	__obf_fda7a11c3287a559, __obf_ab5192d9e0103059 = __obf_03d12f1ecab719e1.__obf_9f33bf964bd83b55(__obf_8d780c51a7e1aba8, __obf_05cde18c0f67f8c5...)
	if __obf_ab5192d9e0103059 != nil {
		return
	}
	__obf_5758290ee6e3c6af, __obf_ab5192d9e0103059 = __obf_fda7a11c3287a559.RowsAffected()
	return
}

// 删除
func (__obf_03d12f1ecab719e1 *Context) Delete() (__obf_5758290ee6e3c6af int64, __obf_ab5192d9e0103059 error) {
	__obf_6cd274412975cdee := "delete from %s %s"
	__obf_b1004fa5302ae44f := __obf_b9a1feb76538f56d(__obf_03d12f1ecab719e1.__obf_e8b5a700a5355b56, Grouping)
	__obf_8d780c51a7e1aba8 := fmt.Sprintf(__obf_6cd274412975cdee, __obf_03d12f1ecab719e1.__obf_1ebef8f2a981f0f4, __obf_b1004fa5302ae44f)
	var __obf_fda7a11c3287a559 sql.Result
	__obf_fda7a11c3287a559, __obf_ab5192d9e0103059 = __obf_03d12f1ecab719e1.__obf_9f33bf964bd83b55(__obf_8d780c51a7e1aba8, __obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f...)
	if __obf_ab5192d9e0103059 != nil {
		return
	}
	__obf_5758290ee6e3c6af, __obf_ab5192d9e0103059 = __obf_fda7a11c3287a559.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_03d12f1ecab719e1 *Context) Select(__obf_f67dd9e9f9b7d2f9 any, sql string, __obf_c2322d6ab4e7aa7f ...any) error {
	__obf_03d12f1ecab719e1.
		sql = sql
	__obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f = __obf_c2322d6ab4e7aa7f
	return __obf_03d12f1ecab719e1.__obf_2fe92164135afc29(__obf_f67dd9e9f9b7d2f9, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_03d12f1ecab719e1 *Context) Get(__obf_f67dd9e9f9b7d2f9 any, sql string, __obf_c2322d6ab4e7aa7f ...any) error {
	__obf_03d12f1ecab719e1.
		sql = sql
	__obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f = __obf_c2322d6ab4e7aa7f
	return __obf_03d12f1ecab719e1.__obf_2fe92164135afc29(__obf_f67dd9e9f9b7d2f9, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_03d12f1ecab719e1 *Context) Exec(sql string, __obf_c2322d6ab4e7aa7f ...any) (sql.Result, error) {
	return __obf_03d12f1ecab719e1.__obf_9f33bf964bd83b55(sql, __obf_c2322d6ab4e7aa7f...)
}

// 创建表
func (__obf_03d12f1ecab719e1 *Context) Create(sql string) (sql.Result, error) {
	return __obf_03d12f1ecab719e1.

		// 删除表
		__obf_9f33bf964bd83b55(sql)
}

func (__obf_03d12f1ecab719e1 *Context) Drop() (sql.Result, error) {
	return __obf_03d12f1ecab719e1.__obf_9f33bf964bd83b55(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_03d12f1ecab719e1.

		/////////////////////////private methods//////////////////////
		__obf_1ebef8f2a981f0f4))
}

// 重置Context
func (__obf_03d12f1ecab719e1 *Context) __obf_68e8df12e48b03e0() *Context {
	__obf_03d12f1ecab719e1.
		sql = ""
	__obf_03d12f1ecab719e1.__obf_1ebef8f2a981f0f4 = ""
	__obf_03d12f1ecab719e1.__obf_739f909e439feddb = []string{}
	__obf_03d12f1ecab719e1.__obf_e8b5a700a5355b56 = []string{}
	__obf_03d12f1ecab719e1.__obf_fa570685b7794b0d = ""
	__obf_03d12f1ecab719e1.__obf_bb5c00d71cf92f10 = ""
	__obf_03d12f1ecab719e1.__obf_36ce77795c55c871 = ""
	__obf_03d12f1ecab719e1.__obf_48eaaf4227b33422 = 0
	__obf_03d12f1ecab719e1.__obf_d3025b9543b73c0b = 0
	__obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f = []any{}
	__obf_03d12f1ecab719e1.__obf_b210b34065b441df = nil
	__obf_03d12f1ecab719e1.__obf_6a22a5c046530e35 = false
	__obf_03d12f1ecab719e1.__obf_5a5f66d6b091294d = false
	return __obf_03d12f1ecab719e1
}

// 查询方法
func (__obf_03d12f1ecab719e1 *Context) __obf_2fe92164135afc29(__obf_f67dd9e9f9b7d2f9 any, __obf_fe876443fb8485de int) (__obf_ab5192d9e0103059 error) {
	defer __obf_03d12f1ecab719e1.__obf_2e89f553fec85399.__obf_ab4cebcf1acecdf0.Put(__obf_03d12f1ecab719e1)
	__obf_518565c480c5e7ee, __obf_e64b2eeb6dee0780 := context.WithTimeout(context.Background(), __obf_03d12f1ecab719e1.__obf_2e89f553fec85399.__obf_cc30e8549e22914a)
	defer __obf_e64b2eeb6dee0780()
	if __obf_03d12f1ecab719e1.sql == "" {
		__obf_03d12f1ecab719e1.
			sql = __obf_03d12f1ecab719e1.__obf_810f5cf0a82ef44a(__obf_f67dd9e9f9b7d2f9)
	}
	switch __obf_fe876443fb8485de {
	case SelectTypeOne:
		if __obf_03d12f1ecab719e1.__obf_b210b34065b441df != nil {
			__obf_ab5192d9e0103059 = __obf_03d12f1ecab719e1.__obf_b210b34065b441df.GetContext(__obf_518565c480c5e7ee, __obf_f67dd9e9f9b7d2f9, __obf_03d12f1ecab719e1.sql, __obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f...)
		} else {
			__obf_ab5192d9e0103059 = __obf_03d12f1ecab719e1.__obf_2e89f553fec85399.GetContext(__obf_518565c480c5e7ee, __obf_f67dd9e9f9b7d2f9, __obf_03d12f1ecab719e1.sql, __obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f...)
		}
	case SelectTypeMany:
		if __obf_03d12f1ecab719e1.__obf_b210b34065b441df != nil {
			__obf_ab5192d9e0103059 = __obf_03d12f1ecab719e1.__obf_b210b34065b441df.SelectContext(__obf_518565c480c5e7ee, __obf_f67dd9e9f9b7d2f9, __obf_03d12f1ecab719e1.sql, __obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f...)
		} else {
			__obf_ab5192d9e0103059 = __obf_03d12f1ecab719e1.__obf_2e89f553fec85399.SelectContext(__obf_518565c480c5e7ee, __obf_f67dd9e9f9b7d2f9, __obf_03d12f1ecab719e1.sql, __obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_03d12f1ecab719e1 *Context) __obf_9f33bf964bd83b55(__obf_8d780c51a7e1aba8 string, __obf_c2322d6ab4e7aa7f ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_8d780c51a7e1aba8, __obf_c2322d6ab4e7aa7f)
	defer __obf_03d12f1ecab719e1.__obf_2e89f553fec85399.__obf_ab4cebcf1acecdf0.Put(__obf_03d12f1ecab719e1)
	__obf_518565c480c5e7ee, __obf_e64b2eeb6dee0780 := context.WithTimeout(context.Background(), __obf_03d12f1ecab719e1.__obf_2e89f553fec85399.__obf_cc30e8549e22914a)
	defer __obf_e64b2eeb6dee0780()

	var __obf_adbbd0f779d860d6 sqlx.ExecerContext
	if __obf_03d12f1ecab719e1.__obf_b210b34065b441df != nil {
		__obf_adbbd0f779d860d6 = __obf_03d12f1ecab719e1.__obf_b210b34065b441df
	} else {
		__obf_adbbd0f779d860d6 = __obf_03d12f1ecab719e1.__obf_2e89f553fec85399
	}
	return __obf_adbbd0f779d860d6.ExecContext(__obf_518565c480c5e7ee, __obf_8d780c51a7e1aba8,

		// select查询语句的拼接
		__obf_c2322d6ab4e7aa7f...)
}

func (__obf_03d12f1ecab719e1 *Context) __obf_810f5cf0a82ef44a(__obf_f67dd9e9f9b7d2f9 any) string {
	var __obf_80ebdb73e162d19b []string
	__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, "select")
	if len(__obf_03d12f1ecab719e1.__obf_739f909e439feddb) != 0 {
		__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, __obf_6c1ae2aa470953c2(__obf_03d12f1ecab719e1.__obf_739f909e439feddb, SeqComma))
	} else {
		__obf_0b4fc2b7f749adb2 := // 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
			// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
			__obf_a643453be0e8a805(__obf_f67dd9e9f9b7d2f9)
		if len(__obf_0b4fc2b7f749adb2) > 0 {
			__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, __obf_6c1ae2aa470953c2(__obf_0b4fc2b7f749adb2, SeqComma))
		} else {
			__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, "*")
		}
	}
	__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, "from "+__obf_03d12f1ecab719e1.__obf_1ebef8f2a981f0f4)
	if len(__obf_03d12f1ecab719e1.__obf_e8b5a700a5355b56) != 0 {
		__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, __obf_b9a1feb76538f56d(__obf_03d12f1ecab719e1.__obf_e8b5a700a5355b56, Grouping))
	}

	if __obf_03d12f1ecab719e1.__obf_bb5c00d71cf92f10 != "" {
		__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, "group by "+__obf_03d12f1ecab719e1.__obf_bb5c00d71cf92f10)
	}

	if __obf_03d12f1ecab719e1.__obf_36ce77795c55c871 != "" {
		__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, "having "+__obf_03d12f1ecab719e1.__obf_36ce77795c55c871)
	}

	if __obf_03d12f1ecab719e1.__obf_fa570685b7794b0d != "" {
		__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, "order by "+__obf_03d12f1ecab719e1.__obf_fa570685b7794b0d)
	}

	if __obf_03d12f1ecab719e1.__obf_48eaaf4227b33422 != 0 {
		__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, fmt.Sprintf("limit %d, %d", __obf_03d12f1ecab719e1.__obf_d3025b9543b73c0b, __obf_03d12f1ecab719e1.__obf_48eaaf4227b33422))
	}
	if __obf_03d12f1ecab719e1.__obf_6a22a5c046530e35 {
		__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, "lock in share mode")
	}
	if __obf_03d12f1ecab719e1.__obf_5a5f66d6b091294d {
		__obf_80ebdb73e162d19b = append(__obf_80ebdb73e162d19b, "for update")
	}
	sql := __obf_6c1ae2aa470953c2(__obf_80ebdb73e162d19b, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_03d12f1ecab719e1.__obf_c2322d6ab4e7aa7f)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_b9a1feb76538f56d(__obf_e8b5a700a5355b56 []string, __obf_359f8c9cb1992d2c string) string {
	if len(__obf_e8b5a700a5355b56) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_e8b5a700a5355b56, __obf_359f8c9cb1992d2c))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_6c1ae2aa470953c2(__obf_c2322d6ab4e7aa7f []string, __obf_0d4ebd50003021e8 string) string {
	return strings.Join(__obf_c2322d6ab4e7aa7f,

		// 解析对象中的 `db tag`
		// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
		__obf_0d4ebd50003021e8)
}

func __obf_a643453be0e8a805(__obf_f67dd9e9f9b7d2f9 any) (__obf_75d4ce8525b58df8 []string) {
	__obf_1fe888a8765085d6 := reflect.ValueOf(__obf_f67dd9e9f9b7d2f9)
	__obf_b9334f423898a2a9 := __obf_1fe888a8765085d6.Type().Elem()
	var __obf_dc6c153502324171 reflect.Type
	if __obf_b9334f423898a2a9.Kind() == reflect.Slice {
		__obf_dc6c153502324171 = __obf_b9334f423898a2a9.Elem()
	} else {
		__obf_dc6c153502324171 = __obf_b9334f423898a2a9
	}
	for __obf_739dac239fb8287e := 0; __obf_739dac239fb8287e < __obf_dc6c153502324171.NumField(); __obf_739dac239fb8287e++ {
		__obf_46112455848804f7 := __obf_dc6c153502324171.Field(__obf_739dac239fb8287e).Tag.Get(DBTag)
		if __obf_46112455848804f7 != "" {
			__obf_75d4ce8525b58df8 = append(__obf_75d4ce8525b58df8, __obf_46112455848804f7)
		}
	}
	return
}

func CleanSQLXMap(__obf_2cd8f1355aa77531 map[string]any) map[string]any {
	for __obf_02cffdb2b73b0b57, __obf_8ec2f1c7b43e366d := range __obf_2cd8f1355aa77531 {
		switch __obf_e87e6f05cc491080 := __obf_8ec2f1c7b43e366d.(type) {
		case []byte:
			__obf_07574ea869b3a910 := string(__obf_e87e6f05cc491080)
			__obf_7acfcf46a1124f5c := // 尝试智能识别具体类型
				strings.TrimSpace(__obf_07574ea869b3a910)

			// bool
			if strings.EqualFold(__obf_7acfcf46a1124f5c, "true") || strings.EqualFold(__obf_7acfcf46a1124f5c, "false") {
				if __obf_052a8b210b7d905e, __obf_ab5192d9e0103059 := strconv.ParseBool(__obf_7acfcf46a1124f5c); __obf_ab5192d9e0103059 == nil {
					__obf_2cd8f1355aa77531[__obf_02cffdb2b73b0b57] = __obf_052a8b210b7d905e
					continue
				}
			}

			// int
			if __obf_739dac239fb8287e, __obf_ab5192d9e0103059 := strconv.Atoi(__obf_7acfcf46a1124f5c); __obf_ab5192d9e0103059 == nil {
				__obf_2cd8f1355aa77531[__obf_02cffdb2b73b0b57] = __obf_739dac239fb8287e
				continue
			}

			// float
			if __obf_717e0741d033c9e0, __obf_ab5192d9e0103059 := strconv.ParseFloat(__obf_7acfcf46a1124f5c, 64); __obf_ab5192d9e0103059 == nil {
				__obf_2cd8f1355aa77531[__obf_02cffdb2b73b0b57] = __obf_717e0741d033c9e0
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_7acfcf46a1124f5c, "{") && strings.HasSuffix(__obf_7acfcf46a1124f5c, "}")) ||
				(strings.HasPrefix(__obf_7acfcf46a1124f5c, "[") && strings.HasSuffix(__obf_7acfcf46a1124f5c, "]")) {
				var __obf_6ff847ea9cbd4820 any
				if __obf_ab5192d9e0103059 := json.Unmarshal(__obf_e87e6f05cc491080, &__obf_6ff847ea9cbd4820); __obf_ab5192d9e0103059 == nil {
					__obf_2cd8f1355aa77531[__obf_02cffdb2b73b0b57] = __obf_6ff847ea9cbd4820
					continue
				}
			}
			__obf_2cd8f1355aa77531[ // 默认保留为 string
			__obf_02cffdb2b73b0b57] = __obf_07574ea869b3a910
		default:
			__obf_2cd8f1355aa77531[__obf_02cffdb2b73b0b57] = __obf_e87e6f05cc491080
		}
	}
	return __obf_2cd8f1355aa77531
}

func ToListMap(__obf_ce52679823fca375 *sqlx.Rows) []map[string]any {
	var __obf_fda7a11c3287a559 []map[string]any
	for __obf_ce52679823fca375.Next() {
		__obf_32c3755ee3b9dc45 := make(map[string]any)
		if __obf_ab5192d9e0103059 := __obf_ce52679823fca375.MapScan(__obf_32c3755ee3b9dc45); __obf_ab5192d9e0103059 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_ab5192d9e0103059.Error()))
			continue
		}
		__obf_fda7a11c3287a559 = append(__obf_fda7a11c3287a559, CleanSQLXMap(__obf_32c3755ee3b9dc45))
	}
	return __obf_fda7a11c3287a559
}
