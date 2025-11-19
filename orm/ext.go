package __obf_e471fc5ea32e9ac0

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
type FuncTx func(__obf_9e151d73a4e332b3 *sqlx.Tx, __obf_c2803eea3fa184fb any) error

func Open(__obf_47b1dc002f1a1832, __obf_ccdf18fec48c50fc string, __obf_e38852d65f365433 time.Duration) (*DB, error) {
	__obf_4fa371f8997ca48b, __obf_c5f1094aec829fc9 := sqlx.Open(__obf_47b1dc002f1a1832, __obf_ccdf18fec48c50fc)
	if __obf_c5f1094aec829fc9 != nil {
		return nil, __obf_c5f1094aec829fc9
	}
	__obf_57b28d01863df3c3 := &DB{
		DB:                     __obf_4fa371f8997ca48b,
		__obf_e38852d65f365433: __obf_e38852d65f365433,
	}
	__obf_57b28d01863df3c3.__obf_1781ce9da7fccb6b.New = func() any {
		return __obf_57b28d01863df3c3.__obf_590cd6326f521f9f()
	}
	return __obf_57b28d01863df3c3, nil
}

type DB struct {
	*sqlx.DB
	__obf_e38852d65f365433 time.Duration
	__obf_1781ce9da7fccb6b sync.Pool
}

func (__obf_4fa371f8997ca48b *DB) __obf_590cd6326f521f9f() *Context {
	return &Context{__obf_4fa371f8997ca48b: __obf_4fa371f8997ca48b}
}

// 获取一个`SQL`执行`Context`
func (__obf_4fa371f8997ca48b *DB) Acquire() *Context {
	// 无需加锁，sync.Pool本身是线程安全的
	__obf_0a997ae29dcd2151 := __obf_4fa371f8997ca48b.__obf_1781ce9da7fccb6b.Get().(*Context)
	__obf_0a997ae29dcd2151.__obf_710a3f1b888bba8f()
	return __obf_0a997ae29dcd2151
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_4fa371f8997ca48b *DB) AcquireTx(__obf_9e151d73a4e332b3 *sqlx.Tx) *Context {
	__obf_0a997ae29dcd2151 := __obf_4fa371f8997ca48b.Acquire()
	__obf_0a997ae29dcd2151.__obf_9e151d73a4e332b3 = __obf_9e151d73a4e332b3
	return __obf_0a997ae29dcd2151
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_4fa371f8997ca48b *DB) WithTx(__obf_10696ef0b8a89c02 FuncTx, __obf_c2803eea3fa184fb any) (__obf_c5f1094aec829fc9 error) {
	var __obf_9e151d73a4e332b3 *sqlx.Tx
	__obf_9e151d73a4e332b3, __obf_c5f1094aec829fc9 = __obf_4fa371f8997ca48b.Beginx()
	if __obf_c5f1094aec829fc9 != nil {
		return
	}
	defer func() {
		if __obf_c5f1094aec829fc9 != nil && __obf_9e151d73a4e332b3 != nil {
			__obf_c5f1094aec829fc9 = __obf_9e151d73a4e332b3.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_c5f1094aec829fc9 = __obf_10696ef0b8a89c02(__obf_9e151d73a4e332b3, __obf_c2803eea3fa184fb); __obf_c5f1094aec829fc9 != nil {
		return
	}

	__obf_c5f1094aec829fc9 = __obf_9e151d73a4e332b3.Commit()
	return
}

type Context struct {
	__obf_4fa371f8997ca48b *DB
	__obf_9e151d73a4e332b3 *sqlx.Tx //事务
	sql                    string
	__obf_cedaef9c76741497 string
	__obf_51b95c03f864a2d6 []string
	__obf_5af184fc0776704f []string
	__obf_fd894e0fc556a765 string
	__obf_7520905e1037a88e string
	__obf_c477a4ab86eee032 string
	__obf_aef9d2894b657a72 int64
	__obf_83afeab8f20ac9fa int64
	__obf_c2803eea3fa184fb []any
	__obf_6ffe93394ec6618f bool //排他锁
	__obf_9a6c8174eadbe8e2 bool //共享锁
}

func (__obf_0a997ae29dcd2151 *Context) Name(__obf_cedaef9c76741497 string) *Context {
	__obf_0a997ae29dcd2151.__obf_cedaef9c76741497 = __obf_cedaef9c76741497
	return __obf_0a997ae29dcd2151
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_0a997ae29dcd2151 *Context) What(__obf_51b95c03f864a2d6 []string) *Context {
	__obf_0a997ae29dcd2151.__obf_51b95c03f864a2d6 = __obf_51b95c03f864a2d6
	return __obf_0a997ae29dcd2151
}

func (__obf_0a997ae29dcd2151 *Context) Where(__obf_5acd8865f259d22f string, __obf_c2803eea3fa184fb ...any) *Context {
	__obf_0a997ae29dcd2151.__obf_5af184fc0776704f = append(__obf_0a997ae29dcd2151.__obf_5af184fc0776704f, __obf_5acd8865f259d22f)
	__obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb = append(__obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb, __obf_c2803eea3fa184fb...)
	return __obf_0a997ae29dcd2151
}

// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
func (__obf_0a997ae29dcd2151 *Context) WhereIn(__obf_fb05cce5b14da37d string, __obf_c2803eea3fa184fb []any) *Context {
	__obf_048d7fae75b6b953 := len(__obf_c2803eea3fa184fb)
	__obf_ab73b15c103ff08d := make([]string, __obf_048d7fae75b6b953)
	for __obf_be83ea6d875ef30b := 0; __obf_be83ea6d875ef30b < __obf_048d7fae75b6b953; __obf_be83ea6d875ef30b++ {
		__obf_ab73b15c103ff08d[__obf_be83ea6d875ef30b] = ParamMarker
	}
	__obf_88d171b55ce07691 := fmt.Sprintf("%s in (%s)", __obf_fb05cce5b14da37d, __obf_5774faa58902e978(__obf_ab73b15c103ff08d, SeqComma))
	return __obf_0a997ae29dcd2151.Where(__obf_88d171b55ce07691, __obf_c2803eea3fa184fb...)
}

func (__obf_0a997ae29dcd2151 *Context) Order(__obf_fd894e0fc556a765 string) *Context {
	__obf_0a997ae29dcd2151.__obf_fd894e0fc556a765 = __obf_fd894e0fc556a765
	return __obf_0a997ae29dcd2151
}

func (__obf_0a997ae29dcd2151 *Context) Limit(__obf_aef9d2894b657a72 int64) *Context {
	__obf_0a997ae29dcd2151.__obf_aef9d2894b657a72 = __obf_aef9d2894b657a72
	return __obf_0a997ae29dcd2151
}

func (__obf_0a997ae29dcd2151 *Context) Offset(__obf_83afeab8f20ac9fa int64) *Context {
	__obf_0a997ae29dcd2151.__obf_83afeab8f20ac9fa = __obf_83afeab8f20ac9fa
	return __obf_0a997ae29dcd2151
}

func (__obf_0a997ae29dcd2151 *Context) Group(__obf_7520905e1037a88e string) *Context {
	__obf_0a997ae29dcd2151.__obf_7520905e1037a88e = __obf_7520905e1037a88e
	return __obf_0a997ae29dcd2151
}

func (__obf_0a997ae29dcd2151 *Context) Having(__obf_c477a4ab86eee032 string, __obf_c2803eea3fa184fb ...any) *Context {
	__obf_0a997ae29dcd2151.__obf_c477a4ab86eee032 = __obf_c477a4ab86eee032
	__obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb = append(__obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb, __obf_c2803eea3fa184fb...)
	return __obf_0a997ae29dcd2151
}

// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
func (__obf_0a997ae29dcd2151 *Context) LockX() *Context {
	__obf_0a997ae29dcd2151.__obf_6ffe93394ec6618f = true
	return __obf_0a997ae29dcd2151
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_0a997ae29dcd2151 *Context) LockS() *Context {
	__obf_0a997ae29dcd2151.__obf_9a6c8174eadbe8e2 = true
	return __obf_0a997ae29dcd2151
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_0a997ae29dcd2151 *Context) FindMany(__obf_d1e55f38350f16b4 any) error {
	return __obf_0a997ae29dcd2151.__obf_306bedeb316ef5df(__obf_d1e55f38350f16b4, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_0a997ae29dcd2151 *Context) FindById(__obf_d1e55f38350f16b4 any) error {
	return __obf_0a997ae29dcd2151.__obf_306bedeb316ef5df(__obf_d1e55f38350f16b4, SelectTypeOne)
}

// 插入
func (__obf_0a997ae29dcd2151 *Context) Insert(__obf_1b3dac75d02767c2 map[string]any) (sql.Result, error) {
	var (
		__obf_95db1c3c210c8eea []string
		__obf_267f29887fa9a411 []any
	)
	for __obf_12c89da7507a8f69, __obf_be4148454b0decfa := range __obf_1b3dac75d02767c2 {
		__obf_95db1c3c210c8eea = append(__obf_95db1c3c210c8eea, __obf_12c89da7507a8f69)
		__obf_267f29887fa9a411 = append(__obf_267f29887fa9a411, __obf_be4148454b0decfa)
	}
	return __obf_0a997ae29dcd2151.InsertBatch(__obf_95db1c3c210c8eea, __obf_267f29887fa9a411)
}

// 批量插入
func (__obf_0a997ae29dcd2151 *Context) InsertBatch(__obf_95db1c3c210c8eea []string, __obf_1b3dac75d02767c2 ...[]any) (sql.Result, error) {
	var (
		__obf_267f29887fa9a411 []any
		__obf_1d4e8ec8e0874e31 []string
	)
	for _, __obf_dbec7ac16c1f18b5 := range __obf_1b3dac75d02767c2 {
		__obf_ab73b15c103ff08d := make([]string, len(__obf_dbec7ac16c1f18b5))
		for __obf_be83ea6d875ef30b, __obf_be4148454b0decfa := range __obf_dbec7ac16c1f18b5 {
			__obf_ab73b15c103ff08d[__obf_be83ea6d875ef30b] = ParamMarker
			__obf_267f29887fa9a411 = append(__obf_267f29887fa9a411, __obf_be4148454b0decfa)
		}
		__obf_1d4e8ec8e0874e31 = append(__obf_1d4e8ec8e0874e31, fmt.Sprintf("(%s)", __obf_5774faa58902e978(__obf_ab73b15c103ff08d, SeqComma)))
	}

	__obf_3d2a5afadabd5bfc := fmt.Sprintf("insert into %s (%s) values %s", __obf_0a997ae29dcd2151.__obf_cedaef9c76741497, __obf_5774faa58902e978(__obf_95db1c3c210c8eea, SeqComma), __obf_5774faa58902e978(__obf_1d4e8ec8e0874e31, SeqComma))
	return __obf_0a997ae29dcd2151.__obf_d2a72b6b2983288e(__obf_3d2a5afadabd5bfc, __obf_267f29887fa9a411...)
}

// 使用map更新
func (__obf_0a997ae29dcd2151 *Context) UpdateMap(__obf_c2803eea3fa184fb map[string]any) (__obf_82fe61e8ad802ae4 int64, __obf_c5f1094aec829fc9 error) {
	var (
		__obf_267f29887fa9a411 []any
		__obf_604b4698448952c4 []string
	)
	for __obf_12c89da7507a8f69, __obf_be4148454b0decfa := range __obf_c2803eea3fa184fb {
		__obf_267f29887fa9a411 = append(__obf_267f29887fa9a411, __obf_be4148454b0decfa)
		__obf_604b4698448952c4 = append(__obf_604b4698448952c4, fmt.Sprintf("%s=%s", __obf_12c89da7507a8f69, ParamMarker))
	}
	__obf_7b09e5f28b16fb96 := __obf_5774faa58902e978(__obf_604b4698448952c4, SeqComma)
	__obf_82fe61e8ad802ae4, __obf_c5f1094aec829fc9 = __obf_0a997ae29dcd2151.Update(__obf_7b09e5f28b16fb96, __obf_267f29887fa9a411...)
	return
}

// 更新
func (__obf_0a997ae29dcd2151 *Context) Update(__obf_7b09e5f28b16fb96 string, __obf_c2803eea3fa184fb ...any) (__obf_82fe61e8ad802ae4 int64, __obf_c5f1094aec829fc9 error) {
	__obf_e9cd5e409745af58 := "update %s set %s %s"
	__obf_5acd8865f259d22f := __obf_c042ceba83a0ce45(__obf_0a997ae29dcd2151.__obf_5af184fc0776704f, Grouping)
	__obf_3d2a5afadabd5bfc := fmt.Sprintf(__obf_e9cd5e409745af58, __obf_0a997ae29dcd2151.__obf_cedaef9c76741497, __obf_7b09e5f28b16fb96, __obf_5acd8865f259d22f)
	__obf_267f29887fa9a411 := append(__obf_c2803eea3fa184fb, __obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb...)
	var __obf_5237801d32b5499f sql.Result
	__obf_5237801d32b5499f, __obf_c5f1094aec829fc9 = __obf_0a997ae29dcd2151.__obf_d2a72b6b2983288e(__obf_3d2a5afadabd5bfc, __obf_267f29887fa9a411...)
	if __obf_c5f1094aec829fc9 != nil {
		return
	}
	__obf_82fe61e8ad802ae4, __obf_c5f1094aec829fc9 = __obf_5237801d32b5499f.RowsAffected()
	return
}

// 删除
func (__obf_0a997ae29dcd2151 *Context) Delete() (__obf_82fe61e8ad802ae4 int64, __obf_c5f1094aec829fc9 error) {
	__obf_e9cd5e409745af58 := "delete from %s %s"
	__obf_5acd8865f259d22f := __obf_c042ceba83a0ce45(__obf_0a997ae29dcd2151.__obf_5af184fc0776704f, Grouping)

	__obf_3d2a5afadabd5bfc := fmt.Sprintf(__obf_e9cd5e409745af58, __obf_0a997ae29dcd2151.__obf_cedaef9c76741497, __obf_5acd8865f259d22f)
	var __obf_5237801d32b5499f sql.Result
	__obf_5237801d32b5499f, __obf_c5f1094aec829fc9 = __obf_0a997ae29dcd2151.__obf_d2a72b6b2983288e(__obf_3d2a5afadabd5bfc, __obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb...)
	if __obf_c5f1094aec829fc9 != nil {
		return
	}
	__obf_82fe61e8ad802ae4, __obf_c5f1094aec829fc9 = __obf_5237801d32b5499f.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_0a997ae29dcd2151 *Context) Select(__obf_d1e55f38350f16b4 any, sql string, __obf_c2803eea3fa184fb ...any) error {
	__obf_0a997ae29dcd2151.sql = sql
	__obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb = __obf_c2803eea3fa184fb
	return __obf_0a997ae29dcd2151.__obf_306bedeb316ef5df(__obf_d1e55f38350f16b4, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_0a997ae29dcd2151 *Context) Get(__obf_d1e55f38350f16b4 any, sql string, __obf_c2803eea3fa184fb ...any) error {
	__obf_0a997ae29dcd2151.sql = sql
	__obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb = __obf_c2803eea3fa184fb
	return __obf_0a997ae29dcd2151.__obf_306bedeb316ef5df(__obf_d1e55f38350f16b4, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_0a997ae29dcd2151 *Context) Exec(sql string, __obf_c2803eea3fa184fb ...any) (sql.Result, error) {
	return __obf_0a997ae29dcd2151.__obf_d2a72b6b2983288e(sql, __obf_c2803eea3fa184fb...)
}

// 创建表
func (__obf_0a997ae29dcd2151 *Context) Create(sql string) (sql.Result, error) {
	return __obf_0a997ae29dcd2151.__obf_d2a72b6b2983288e(sql)
}

// 删除表
func (__obf_0a997ae29dcd2151 *Context) Drop() (sql.Result, error) {
	return __obf_0a997ae29dcd2151.__obf_d2a72b6b2983288e(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_0a997ae29dcd2151.__obf_cedaef9c76741497))
}

/////////////////////////private methods//////////////////////

// 重置Context
func (__obf_0a997ae29dcd2151 *Context) __obf_710a3f1b888bba8f() *Context {
	__obf_0a997ae29dcd2151.sql = ""
	__obf_0a997ae29dcd2151.__obf_cedaef9c76741497 = ""
	__obf_0a997ae29dcd2151.__obf_51b95c03f864a2d6 = []string{}
	__obf_0a997ae29dcd2151.__obf_5af184fc0776704f = []string{}
	__obf_0a997ae29dcd2151.__obf_fd894e0fc556a765 = ""
	__obf_0a997ae29dcd2151.__obf_7520905e1037a88e = ""
	__obf_0a997ae29dcd2151.__obf_c477a4ab86eee032 = ""
	__obf_0a997ae29dcd2151.__obf_aef9d2894b657a72 = 0
	__obf_0a997ae29dcd2151.__obf_83afeab8f20ac9fa = 0
	__obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb = []any{}
	__obf_0a997ae29dcd2151.__obf_9e151d73a4e332b3 = nil
	__obf_0a997ae29dcd2151.__obf_9a6c8174eadbe8e2 = false
	__obf_0a997ae29dcd2151.__obf_6ffe93394ec6618f = false
	return __obf_0a997ae29dcd2151
}

// 查询方法
func (__obf_0a997ae29dcd2151 *Context) __obf_306bedeb316ef5df(__obf_d1e55f38350f16b4 any, __obf_4a751e3a798d4a61 int) (__obf_c5f1094aec829fc9 error) {
	defer __obf_0a997ae29dcd2151.__obf_4fa371f8997ca48b.__obf_1781ce9da7fccb6b.Put(__obf_0a997ae29dcd2151)
	__obf_4a7d1ab77451dfb9, __obf_4f4c5ec1fff5e3d7 := context.WithTimeout(context.Background(), __obf_0a997ae29dcd2151.__obf_4fa371f8997ca48b.__obf_e38852d65f365433)
	defer __obf_4f4c5ec1fff5e3d7()
	if __obf_0a997ae29dcd2151.sql == "" {
		__obf_0a997ae29dcd2151.sql = __obf_0a997ae29dcd2151.__obf_7a76c31a0bf598e8(__obf_d1e55f38350f16b4)
	}
	switch __obf_4a751e3a798d4a61 {
	case SelectTypeOne:
		if __obf_0a997ae29dcd2151.__obf_9e151d73a4e332b3 != nil {
			__obf_c5f1094aec829fc9 = __obf_0a997ae29dcd2151.__obf_9e151d73a4e332b3.GetContext(__obf_4a7d1ab77451dfb9, __obf_d1e55f38350f16b4, __obf_0a997ae29dcd2151.sql, __obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb...)
		} else {
			__obf_c5f1094aec829fc9 = __obf_0a997ae29dcd2151.__obf_4fa371f8997ca48b.GetContext(__obf_4a7d1ab77451dfb9, __obf_d1e55f38350f16b4, __obf_0a997ae29dcd2151.sql, __obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb...)
		}
	case SelectTypeMany:
		if __obf_0a997ae29dcd2151.__obf_9e151d73a4e332b3 != nil {
			__obf_c5f1094aec829fc9 = __obf_0a997ae29dcd2151.__obf_9e151d73a4e332b3.SelectContext(__obf_4a7d1ab77451dfb9, __obf_d1e55f38350f16b4, __obf_0a997ae29dcd2151.sql, __obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb...)
		} else {
			__obf_c5f1094aec829fc9 = __obf_0a997ae29dcd2151.__obf_4fa371f8997ca48b.SelectContext(__obf_4a7d1ab77451dfb9, __obf_d1e55f38350f16b4, __obf_0a997ae29dcd2151.sql, __obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_0a997ae29dcd2151 *Context) __obf_d2a72b6b2983288e(__obf_3d2a5afadabd5bfc string, __obf_c2803eea3fa184fb ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb)
	defer __obf_0a997ae29dcd2151.__obf_4fa371f8997ca48b.__obf_1781ce9da7fccb6b.Put(__obf_0a997ae29dcd2151)
	__obf_4a7d1ab77451dfb9, __obf_4f4c5ec1fff5e3d7 := context.WithTimeout(context.Background(), __obf_0a997ae29dcd2151.__obf_4fa371f8997ca48b.__obf_e38852d65f365433)
	defer __obf_4f4c5ec1fff5e3d7()

	var __obf_c418bf05de267175 sqlx.ExecerContext
	if __obf_0a997ae29dcd2151.__obf_9e151d73a4e332b3 != nil {
		__obf_c418bf05de267175 = __obf_0a997ae29dcd2151.__obf_9e151d73a4e332b3
	} else {
		__obf_c418bf05de267175 = __obf_0a997ae29dcd2151.__obf_4fa371f8997ca48b
	}
	return __obf_c418bf05de267175.ExecContext(__obf_4a7d1ab77451dfb9, __obf_3d2a5afadabd5bfc, __obf_c2803eea3fa184fb...)
}

// select查询语句的拼接
func (__obf_0a997ae29dcd2151 *Context) __obf_7a76c31a0bf598e8(__obf_d1e55f38350f16b4 any) string {
	var __obf_d1664e864ff7f6d7 []string
	__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, "select")
	if len(__obf_0a997ae29dcd2151.__obf_51b95c03f864a2d6) != 0 {
		__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, __obf_5774faa58902e978(__obf_0a997ae29dcd2151.__obf_51b95c03f864a2d6, SeqComma))
	} else {
		// 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
		// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
		__obf_df9a1d2f4ee7ef3a := __obf_50aaccb6311abd8a(__obf_d1e55f38350f16b4)
		if len(__obf_df9a1d2f4ee7ef3a) > 0 {
			__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, __obf_5774faa58902e978(__obf_df9a1d2f4ee7ef3a, SeqComma))
		} else {
			__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, "*")
		}
	}
	__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, "from "+__obf_0a997ae29dcd2151.__obf_cedaef9c76741497)
	if len(__obf_0a997ae29dcd2151.__obf_5af184fc0776704f) != 0 {
		__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, __obf_c042ceba83a0ce45(__obf_0a997ae29dcd2151.__obf_5af184fc0776704f, Grouping))
	}

	if __obf_0a997ae29dcd2151.__obf_7520905e1037a88e != "" {
		__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, "group by "+__obf_0a997ae29dcd2151.__obf_7520905e1037a88e)
	}

	if __obf_0a997ae29dcd2151.__obf_c477a4ab86eee032 != "" {
		__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, "having "+__obf_0a997ae29dcd2151.__obf_c477a4ab86eee032)
	}

	if __obf_0a997ae29dcd2151.__obf_fd894e0fc556a765 != "" {
		__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, "order by "+__obf_0a997ae29dcd2151.__obf_fd894e0fc556a765)
	}

	if __obf_0a997ae29dcd2151.__obf_aef9d2894b657a72 != 0 {
		__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, fmt.Sprintf("limit %d, %d", __obf_0a997ae29dcd2151.__obf_83afeab8f20ac9fa, __obf_0a997ae29dcd2151.__obf_aef9d2894b657a72))
	}
	if __obf_0a997ae29dcd2151.__obf_9a6c8174eadbe8e2 {
		__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, "lock in share mode")
	}
	if __obf_0a997ae29dcd2151.__obf_6ffe93394ec6618f {
		__obf_d1664e864ff7f6d7 = append(__obf_d1664e864ff7f6d7, "for update")
	}
	sql := __obf_5774faa58902e978(__obf_d1664e864ff7f6d7, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_0a997ae29dcd2151.__obf_c2803eea3fa184fb)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_c042ceba83a0ce45(__obf_5af184fc0776704f []string, __obf_2505e29fc93d39ba string) string {
	if len(__obf_5af184fc0776704f) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_5af184fc0776704f, __obf_2505e29fc93d39ba))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_5774faa58902e978(__obf_c2803eea3fa184fb []string, __obf_c4e001fb48170fb0 string) string {
	return strings.Join(__obf_c2803eea3fa184fb, __obf_c4e001fb48170fb0)
}

// 解析对象中的 `db tag`
// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
func __obf_50aaccb6311abd8a(__obf_d1e55f38350f16b4 any) (__obf_95db1c3c210c8eea []string) {
	__obf_a12d79aa3dd7e532 := reflect.ValueOf(__obf_d1e55f38350f16b4)
	__obf_0749fc6f94fb5a7d := __obf_a12d79aa3dd7e532.Type().Elem()
	var __obf_0373d9c9d4051e69 reflect.Type
	if __obf_0749fc6f94fb5a7d.Kind() == reflect.Slice {
		__obf_0373d9c9d4051e69 = __obf_0749fc6f94fb5a7d.Elem()
	} else {
		__obf_0373d9c9d4051e69 = __obf_0749fc6f94fb5a7d
	}
	for __obf_be83ea6d875ef30b := 0; __obf_be83ea6d875ef30b < __obf_0373d9c9d4051e69.NumField(); __obf_be83ea6d875ef30b++ {
		__obf_7c8c104e88806e18 := __obf_0373d9c9d4051e69.Field(__obf_be83ea6d875ef30b).Tag.Get(DBTag)
		if __obf_7c8c104e88806e18 != "" {
			__obf_95db1c3c210c8eea = append(__obf_95db1c3c210c8eea, __obf_7c8c104e88806e18)
		}
	}
	return
}

func CleanSQLXMap(__obf_a6f100900b64e452 map[string]any) map[string]any {
	for __obf_12c89da7507a8f69, __obf_be4148454b0decfa := range __obf_a6f100900b64e452 {
		switch __obf_b6bf8b8bc11918d7 := __obf_be4148454b0decfa.(type) {
		case []byte:
			__obf_153c6aff24c72d24 := string(__obf_b6bf8b8bc11918d7)

			// 尝试智能识别具体类型
			__obf_0844f011b55aab1b := strings.TrimSpace(__obf_153c6aff24c72d24)

			// bool
			if strings.EqualFold(__obf_0844f011b55aab1b, "true") || strings.EqualFold(__obf_0844f011b55aab1b, "false") {
				if __obf_6bc7442d4e9fe88a, __obf_c5f1094aec829fc9 := strconv.ParseBool(__obf_0844f011b55aab1b); __obf_c5f1094aec829fc9 == nil {
					__obf_a6f100900b64e452[__obf_12c89da7507a8f69] = __obf_6bc7442d4e9fe88a
					continue
				}
			}

			// int
			if __obf_be83ea6d875ef30b, __obf_c5f1094aec829fc9 := strconv.Atoi(__obf_0844f011b55aab1b); __obf_c5f1094aec829fc9 == nil {
				__obf_a6f100900b64e452[__obf_12c89da7507a8f69] = __obf_be83ea6d875ef30b
				continue
			}

			// float
			if __obf_715a8e43dc438a16, __obf_c5f1094aec829fc9 := strconv.ParseFloat(__obf_0844f011b55aab1b, 64); __obf_c5f1094aec829fc9 == nil {
				__obf_a6f100900b64e452[__obf_12c89da7507a8f69] = __obf_715a8e43dc438a16
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_0844f011b55aab1b, "{") && strings.HasSuffix(__obf_0844f011b55aab1b, "}")) ||
				(strings.HasPrefix(__obf_0844f011b55aab1b, "[") && strings.HasSuffix(__obf_0844f011b55aab1b, "]")) {
				var __obf_3bc93023d0742fbb any
				if __obf_c5f1094aec829fc9 := json.Unmarshal(__obf_b6bf8b8bc11918d7, &__obf_3bc93023d0742fbb); __obf_c5f1094aec829fc9 == nil {
					__obf_a6f100900b64e452[__obf_12c89da7507a8f69] = __obf_3bc93023d0742fbb
					continue
				}
			}

			// 默认保留为 string
			__obf_a6f100900b64e452[__obf_12c89da7507a8f69] = __obf_153c6aff24c72d24
		default:
			__obf_a6f100900b64e452[__obf_12c89da7507a8f69] = __obf_b6bf8b8bc11918d7
		}
	}
	return __obf_a6f100900b64e452
}

func ToListMap(__obf_31a48fd558fdaaab *sqlx.Rows) []map[string]any {
	var __obf_5237801d32b5499f []map[string]any
	for __obf_31a48fd558fdaaab.Next() {
		__obf_1770879c3eb4faba := make(map[string]any)
		if __obf_c5f1094aec829fc9 := __obf_31a48fd558fdaaab.MapScan(__obf_1770879c3eb4faba); __obf_c5f1094aec829fc9 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_c5f1094aec829fc9.Error()))
			continue
		}

		__obf_5237801d32b5499f = append(__obf_5237801d32b5499f, CleanSQLXMap(__obf_1770879c3eb4faba))
	}
	return __obf_5237801d32b5499f
}
