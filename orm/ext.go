package __obf_86786288bdd26c8f

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
type FuncTx func(__obf_945c2e80cb402175 *sqlx.Tx, __obf_d51eb37f5f1cb276 any) error

func Open(__obf_c930766e485b834a, __obf_beb3eb0a7de124db string, __obf_f7c787543fd51ca7 time.Duration) (*DB, error) {
	__obf_019b2ae2e7d131c0, __obf_87239f12e0f48308 := sqlx.Open(__obf_c930766e485b834a, __obf_beb3eb0a7de124db)
	if __obf_87239f12e0f48308 != nil {
		return nil, __obf_87239f12e0f48308
	}
	__obf_8ed96d1577fd8cf6 := &DB{
		DB: __obf_019b2ae2e7d131c0, __obf_f7c787543fd51ca7: __obf_f7c787543fd51ca7,
	}
	__obf_8ed96d1577fd8cf6.__obf_deed23c41c242c9a.
		New = func() any {
		return __obf_8ed96d1577fd8cf6.__obf_d1c8a84bb4d4675d()
	}
	return __obf_8ed96d1577fd8cf6, nil
}

type DB struct {
	*sqlx.DB
	__obf_f7c787543fd51ca7 time.Duration
	__obf_deed23c41c242c9a sync.Pool
}

func (__obf_019b2ae2e7d131c0 *DB) __obf_d1c8a84bb4d4675d() *Context {
	return &Context{__obf_019b2ae2e7d131c0: __obf_019b2ae2e7d131c0}
}

// 获取一个`SQL`执行`Context`
func (__obf_019b2ae2e7d131c0 *DB) Acquire() *Context {
	__obf_ffae9d01a4fb83a6 := // 无需加锁，sync.Pool本身是线程安全的
		__obf_019b2ae2e7d131c0.__obf_deed23c41c242c9a.Get().(*Context)
	__obf_ffae9d01a4fb83a6.__obf_000b48902c5508d3()
	return __obf_ffae9d01a4fb83a6
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_019b2ae2e7d131c0 *DB) AcquireTx(__obf_945c2e80cb402175 *sqlx.Tx) *Context {
	__obf_ffae9d01a4fb83a6 := __obf_019b2ae2e7d131c0.Acquire()
	__obf_ffae9d01a4fb83a6.__obf_945c2e80cb402175 = __obf_945c2e80cb402175
	return __obf_ffae9d01a4fb83a6
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_019b2ae2e7d131c0 *DB) WithTx(__obf_9d06062984868572 FuncTx, __obf_d51eb37f5f1cb276 any) (__obf_87239f12e0f48308 error) {
	var __obf_945c2e80cb402175 *sqlx.Tx
	__obf_945c2e80cb402175, __obf_87239f12e0f48308 = __obf_019b2ae2e7d131c0.Beginx()
	if __obf_87239f12e0f48308 != nil {
		return
	}
	defer func() {
		if __obf_87239f12e0f48308 != nil && __obf_945c2e80cb402175 != nil {
			__obf_87239f12e0f48308 = __obf_945c2e80cb402175.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_87239f12e0f48308 = __obf_9d06062984868572(__obf_945c2e80cb402175, __obf_d51eb37f5f1cb276); __obf_87239f12e0f48308 != nil {
		return
	}
	__obf_87239f12e0f48308 = __obf_945c2e80cb402175.Commit()
	return
}

type Context struct {
	__obf_019b2ae2e7d131c0 *DB
	__obf_945c2e80cb402175 *sqlx.Tx //事务
	sql                    string
	__obf_005f563d12feeb25 string
	__obf_298548f7b2ed8c1a []string
	__obf_4210685111b131c8 []string
	__obf_b6b7d1d38659c765 string
	__obf_aa3113fef1a0d74d string
	__obf_7faf07cb163b0164 string
	__obf_e11a41cd614462e2 int64
	__obf_f2e5b45bab2a89f5 int64
	__obf_d51eb37f5f1cb276 []any
	__obf_159039e7777eb165 bool
	__obf_e771540f8f5111b7 bool //排他锁
	//共享锁
}

func (__obf_ffae9d01a4fb83a6 *Context) Name(__obf_005f563d12feeb25 string) *Context {
	__obf_ffae9d01a4fb83a6.__obf_005f563d12feeb25 = __obf_005f563d12feeb25
	return __obf_ffae9d01a4fb83a6
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_ffae9d01a4fb83a6 *Context) What(__obf_298548f7b2ed8c1a []string) *Context {
	__obf_ffae9d01a4fb83a6.__obf_298548f7b2ed8c1a = __obf_298548f7b2ed8c1a
	return __obf_ffae9d01a4fb83a6
}

func (__obf_ffae9d01a4fb83a6 *Context) Where(__obf_3cb9d69e303725bf string, __obf_d51eb37f5f1cb276 ...any) *Context {
	__obf_ffae9d01a4fb83a6.__obf_4210685111b131c8 = append(__obf_ffae9d01a4fb83a6.__obf_4210685111b131c8, __obf_3cb9d69e303725bf)
	__obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276 = append(__obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276,

		// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
		__obf_d51eb37f5f1cb276...)
	return __obf_ffae9d01a4fb83a6
}

func (__obf_ffae9d01a4fb83a6 *Context) WhereIn(__obf_89f6e57c8628366a string, __obf_d51eb37f5f1cb276 []any) *Context {
	__obf_7fac5a2246eb6f34 := len(__obf_d51eb37f5f1cb276)
	__obf_8044bd7bf175266e := make([]string, __obf_7fac5a2246eb6f34)
	for __obf_33fdeff971fca5a6 := 0; __obf_33fdeff971fca5a6 < __obf_7fac5a2246eb6f34; __obf_33fdeff971fca5a6++ {
		__obf_8044bd7bf175266e[__obf_33fdeff971fca5a6] = ParamMarker
	}
	__obf_a291118b285b2856 := fmt.Sprintf("%s in (%s)", __obf_89f6e57c8628366a, __obf_525fb93a93b88033(__obf_8044bd7bf175266e, SeqComma))
	return __obf_ffae9d01a4fb83a6.Where(__obf_a291118b285b2856, __obf_d51eb37f5f1cb276...)
}

func (__obf_ffae9d01a4fb83a6 *Context) Order(__obf_b6b7d1d38659c765 string) *Context {
	__obf_ffae9d01a4fb83a6.__obf_b6b7d1d38659c765 = __obf_b6b7d1d38659c765
	return __obf_ffae9d01a4fb83a6
}

func (__obf_ffae9d01a4fb83a6 *Context) Limit(__obf_e11a41cd614462e2 int64) *Context {
	__obf_ffae9d01a4fb83a6.__obf_e11a41cd614462e2 = __obf_e11a41cd614462e2
	return __obf_ffae9d01a4fb83a6
}

func (__obf_ffae9d01a4fb83a6 *Context) Offset(__obf_f2e5b45bab2a89f5 int64) *Context {
	__obf_ffae9d01a4fb83a6.__obf_f2e5b45bab2a89f5 = __obf_f2e5b45bab2a89f5
	return __obf_ffae9d01a4fb83a6
}

func (__obf_ffae9d01a4fb83a6 *Context) Group(__obf_aa3113fef1a0d74d string) *Context {
	__obf_ffae9d01a4fb83a6.__obf_aa3113fef1a0d74d = __obf_aa3113fef1a0d74d
	return __obf_ffae9d01a4fb83a6
}

func (__obf_ffae9d01a4fb83a6 *Context) Having(__obf_7faf07cb163b0164 string, __obf_d51eb37f5f1cb276 ...any) *Context {
	__obf_ffae9d01a4fb83a6.__obf_7faf07cb163b0164 = __obf_7faf07cb163b0164
	__obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276 = append(__obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276,

		// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
		__obf_d51eb37f5f1cb276...)
	return __obf_ffae9d01a4fb83a6
}

func (__obf_ffae9d01a4fb83a6 *Context) LockX() *Context {
	__obf_ffae9d01a4fb83a6.__obf_159039e7777eb165 = true
	return __obf_ffae9d01a4fb83a6
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_ffae9d01a4fb83a6 *Context) LockS() *Context {
	__obf_ffae9d01a4fb83a6.__obf_e771540f8f5111b7 = true
	return __obf_ffae9d01a4fb83a6
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_ffae9d01a4fb83a6 *Context) FindMany(__obf_a5b72982c262db32 any) error {
	return __obf_ffae9d01a4fb83a6.__obf_227db3f3cae441ff(__obf_a5b72982c262db32, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_ffae9d01a4fb83a6 *Context) FindById(__obf_a5b72982c262db32 any) error {
	return __obf_ffae9d01a4fb83a6.__obf_227db3f3cae441ff(__obf_a5b72982c262db32, SelectTypeOne)
}

// 插入
func (__obf_ffae9d01a4fb83a6 *Context) Insert(__obf_da9bca691bd3b236 map[string]any) (sql.Result, error) {
	var (
		__obf_44baf82a67e2ecb5 []string
		__obf_dc6f6069ded03706 []any
	)
	for __obf_299b79baf6e9ee26, __obf_0c15a3a65dfe2b86 := range __obf_da9bca691bd3b236 {
		__obf_44baf82a67e2ecb5 = append(__obf_44baf82a67e2ecb5, __obf_299b79baf6e9ee26)
		__obf_dc6f6069ded03706 = append(__obf_dc6f6069ded03706, __obf_0c15a3a65dfe2b86)
	}
	return __obf_ffae9d01a4fb83a6.InsertBatch(__obf_44baf82a67e2ecb5,

		// 批量插入
		__obf_dc6f6069ded03706)
}

func (__obf_ffae9d01a4fb83a6 *Context) InsertBatch(__obf_44baf82a67e2ecb5 []string, __obf_da9bca691bd3b236 ...[]any) (sql.Result, error) {
	var (
		__obf_dc6f6069ded03706 []any
		__obf_84080fb6d69d44ff []string
	)
	for _, __obf_e4548b7a468dd4b9 := range __obf_da9bca691bd3b236 {
		__obf_8044bd7bf175266e := make([]string, len(__obf_e4548b7a468dd4b9))
		for __obf_33fdeff971fca5a6, __obf_0c15a3a65dfe2b86 := range __obf_e4548b7a468dd4b9 {
			__obf_8044bd7bf175266e[__obf_33fdeff971fca5a6] = ParamMarker
			__obf_dc6f6069ded03706 = append(__obf_dc6f6069ded03706, __obf_0c15a3a65dfe2b86)
		}
		__obf_84080fb6d69d44ff = append(__obf_84080fb6d69d44ff, fmt.Sprintf("(%s)", __obf_525fb93a93b88033(__obf_8044bd7bf175266e, SeqComma)))
	}
	__obf_63a6a7575ede7589 := fmt.Sprintf("insert into %s (%s) values %s", __obf_ffae9d01a4fb83a6.__obf_005f563d12feeb25, __obf_525fb93a93b88033(__obf_44baf82a67e2ecb5, SeqComma), __obf_525fb93a93b88033(__obf_84080fb6d69d44ff, SeqComma))
	return __obf_ffae9d01a4fb83a6.__obf_5391411bdd39c82c(__obf_63a6a7575ede7589,

		// 使用map更新
		__obf_dc6f6069ded03706...)
}

func (__obf_ffae9d01a4fb83a6 *Context) UpdateMap(__obf_d51eb37f5f1cb276 map[string]any) (__obf_f8e432258242d186 int64, __obf_87239f12e0f48308 error) {
	var (
		__obf_dc6f6069ded03706 []any
		__obf_f34bc8a524409a60 []string
	)
	for __obf_299b79baf6e9ee26, __obf_0c15a3a65dfe2b86 := range __obf_d51eb37f5f1cb276 {
		__obf_dc6f6069ded03706 = append(__obf_dc6f6069ded03706, __obf_0c15a3a65dfe2b86)
		__obf_f34bc8a524409a60 = append(__obf_f34bc8a524409a60, fmt.Sprintf("%s=%s", __obf_299b79baf6e9ee26, ParamMarker))
	}
	__obf_97e3c5279b1f5477 := __obf_525fb93a93b88033(__obf_f34bc8a524409a60, SeqComma)
	__obf_f8e432258242d186, __obf_87239f12e0f48308 = __obf_ffae9d01a4fb83a6.Update(__obf_97e3c5279b1f5477, __obf_dc6f6069ded03706...)
	return
}

// 更新
func (__obf_ffae9d01a4fb83a6 *Context) Update(__obf_97e3c5279b1f5477 string, __obf_d51eb37f5f1cb276 ...any) (__obf_f8e432258242d186 int64, __obf_87239f12e0f48308 error) {
	__obf_948ed193ce2bb112 := "update %s set %s %s"
	__obf_3cb9d69e303725bf := __obf_7affa20e9841bed0(__obf_ffae9d01a4fb83a6.__obf_4210685111b131c8, Grouping)
	__obf_63a6a7575ede7589 := fmt.Sprintf(__obf_948ed193ce2bb112, __obf_ffae9d01a4fb83a6.__obf_005f563d12feeb25, __obf_97e3c5279b1f5477, __obf_3cb9d69e303725bf)
	__obf_dc6f6069ded03706 := append(__obf_d51eb37f5f1cb276, __obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276...)
	var __obf_d46b9a4ac0947f8a sql.Result
	__obf_d46b9a4ac0947f8a, __obf_87239f12e0f48308 = __obf_ffae9d01a4fb83a6.__obf_5391411bdd39c82c(__obf_63a6a7575ede7589, __obf_dc6f6069ded03706...)
	if __obf_87239f12e0f48308 != nil {
		return
	}
	__obf_f8e432258242d186, __obf_87239f12e0f48308 = __obf_d46b9a4ac0947f8a.RowsAffected()
	return
}

// 删除
func (__obf_ffae9d01a4fb83a6 *Context) Delete() (__obf_f8e432258242d186 int64, __obf_87239f12e0f48308 error) {
	__obf_948ed193ce2bb112 := "delete from %s %s"
	__obf_3cb9d69e303725bf := __obf_7affa20e9841bed0(__obf_ffae9d01a4fb83a6.__obf_4210685111b131c8, Grouping)
	__obf_63a6a7575ede7589 := fmt.Sprintf(__obf_948ed193ce2bb112, __obf_ffae9d01a4fb83a6.__obf_005f563d12feeb25, __obf_3cb9d69e303725bf)
	var __obf_d46b9a4ac0947f8a sql.Result
	__obf_d46b9a4ac0947f8a, __obf_87239f12e0f48308 = __obf_ffae9d01a4fb83a6.__obf_5391411bdd39c82c(__obf_63a6a7575ede7589, __obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276...)
	if __obf_87239f12e0f48308 != nil {
		return
	}
	__obf_f8e432258242d186, __obf_87239f12e0f48308 = __obf_d46b9a4ac0947f8a.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_ffae9d01a4fb83a6 *Context) Select(__obf_a5b72982c262db32 any, sql string, __obf_d51eb37f5f1cb276 ...any) error {
	__obf_ffae9d01a4fb83a6.
		sql = sql
	__obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276 = __obf_d51eb37f5f1cb276
	return __obf_ffae9d01a4fb83a6.__obf_227db3f3cae441ff(__obf_a5b72982c262db32, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_ffae9d01a4fb83a6 *Context) Get(__obf_a5b72982c262db32 any, sql string, __obf_d51eb37f5f1cb276 ...any) error {
	__obf_ffae9d01a4fb83a6.
		sql = sql
	__obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276 = __obf_d51eb37f5f1cb276
	return __obf_ffae9d01a4fb83a6.__obf_227db3f3cae441ff(__obf_a5b72982c262db32, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_ffae9d01a4fb83a6 *Context) Exec(sql string, __obf_d51eb37f5f1cb276 ...any) (sql.Result, error) {
	return __obf_ffae9d01a4fb83a6.__obf_5391411bdd39c82c(sql, __obf_d51eb37f5f1cb276...)
}

// 创建表
func (__obf_ffae9d01a4fb83a6 *Context) Create(sql string) (sql.Result, error) {
	return __obf_ffae9d01a4fb83a6.

		// 删除表
		__obf_5391411bdd39c82c(sql)
}

func (__obf_ffae9d01a4fb83a6 *Context) Drop() (sql.Result, error) {
	return __obf_ffae9d01a4fb83a6.__obf_5391411bdd39c82c(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_ffae9d01a4fb83a6.

		/////////////////////////private methods//////////////////////
		__obf_005f563d12feeb25))
}

// 重置Context
func (__obf_ffae9d01a4fb83a6 *Context) __obf_000b48902c5508d3() *Context {
	__obf_ffae9d01a4fb83a6.
		sql = ""
	__obf_ffae9d01a4fb83a6.__obf_005f563d12feeb25 = ""
	__obf_ffae9d01a4fb83a6.__obf_298548f7b2ed8c1a = []string{}
	__obf_ffae9d01a4fb83a6.__obf_4210685111b131c8 = []string{}
	__obf_ffae9d01a4fb83a6.__obf_b6b7d1d38659c765 = ""
	__obf_ffae9d01a4fb83a6.__obf_aa3113fef1a0d74d = ""
	__obf_ffae9d01a4fb83a6.__obf_7faf07cb163b0164 = ""
	__obf_ffae9d01a4fb83a6.__obf_e11a41cd614462e2 = 0
	__obf_ffae9d01a4fb83a6.__obf_f2e5b45bab2a89f5 = 0
	__obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276 = []any{}
	__obf_ffae9d01a4fb83a6.__obf_945c2e80cb402175 = nil
	__obf_ffae9d01a4fb83a6.__obf_e771540f8f5111b7 = false
	__obf_ffae9d01a4fb83a6.__obf_159039e7777eb165 = false
	return __obf_ffae9d01a4fb83a6
}

// 查询方法
func (__obf_ffae9d01a4fb83a6 *Context) __obf_227db3f3cae441ff(__obf_a5b72982c262db32 any, __obf_d81ce366008c8ac6 int) (__obf_87239f12e0f48308 error) {
	defer __obf_ffae9d01a4fb83a6.__obf_019b2ae2e7d131c0.__obf_deed23c41c242c9a.Put(__obf_ffae9d01a4fb83a6)
	__obf_cb97fa3dfe04df74, __obf_cc56d65a0fa2c0c7 := context.WithTimeout(context.Background(), __obf_ffae9d01a4fb83a6.__obf_019b2ae2e7d131c0.__obf_f7c787543fd51ca7)
	defer __obf_cc56d65a0fa2c0c7()
	if __obf_ffae9d01a4fb83a6.sql == "" {
		__obf_ffae9d01a4fb83a6.
			sql = __obf_ffae9d01a4fb83a6.__obf_4a8a9e0760f863cd(__obf_a5b72982c262db32)
	}
	switch __obf_d81ce366008c8ac6 {
	case SelectTypeOne:
		if __obf_ffae9d01a4fb83a6.__obf_945c2e80cb402175 != nil {
			__obf_87239f12e0f48308 = __obf_ffae9d01a4fb83a6.__obf_945c2e80cb402175.GetContext(__obf_cb97fa3dfe04df74, __obf_a5b72982c262db32, __obf_ffae9d01a4fb83a6.sql, __obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276...)
		} else {
			__obf_87239f12e0f48308 = __obf_ffae9d01a4fb83a6.__obf_019b2ae2e7d131c0.GetContext(__obf_cb97fa3dfe04df74, __obf_a5b72982c262db32, __obf_ffae9d01a4fb83a6.sql, __obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276...)
		}
	case SelectTypeMany:
		if __obf_ffae9d01a4fb83a6.__obf_945c2e80cb402175 != nil {
			__obf_87239f12e0f48308 = __obf_ffae9d01a4fb83a6.__obf_945c2e80cb402175.SelectContext(__obf_cb97fa3dfe04df74, __obf_a5b72982c262db32, __obf_ffae9d01a4fb83a6.sql, __obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276...)
		} else {
			__obf_87239f12e0f48308 = __obf_ffae9d01a4fb83a6.__obf_019b2ae2e7d131c0.SelectContext(__obf_cb97fa3dfe04df74, __obf_a5b72982c262db32, __obf_ffae9d01a4fb83a6.sql, __obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_ffae9d01a4fb83a6 *Context) __obf_5391411bdd39c82c(__obf_63a6a7575ede7589 string, __obf_d51eb37f5f1cb276 ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_63a6a7575ede7589, __obf_d51eb37f5f1cb276)
	defer __obf_ffae9d01a4fb83a6.__obf_019b2ae2e7d131c0.__obf_deed23c41c242c9a.Put(__obf_ffae9d01a4fb83a6)
	__obf_cb97fa3dfe04df74, __obf_cc56d65a0fa2c0c7 := context.WithTimeout(context.Background(), __obf_ffae9d01a4fb83a6.__obf_019b2ae2e7d131c0.__obf_f7c787543fd51ca7)
	defer __obf_cc56d65a0fa2c0c7()

	var __obf_0ef2c86cb2168815 sqlx.ExecerContext
	if __obf_ffae9d01a4fb83a6.__obf_945c2e80cb402175 != nil {
		__obf_0ef2c86cb2168815 = __obf_ffae9d01a4fb83a6.__obf_945c2e80cb402175
	} else {
		__obf_0ef2c86cb2168815 = __obf_ffae9d01a4fb83a6.__obf_019b2ae2e7d131c0
	}
	return __obf_0ef2c86cb2168815.ExecContext(__obf_cb97fa3dfe04df74, __obf_63a6a7575ede7589,

		// select查询语句的拼接
		__obf_d51eb37f5f1cb276...)
}

func (__obf_ffae9d01a4fb83a6 *Context) __obf_4a8a9e0760f863cd(__obf_a5b72982c262db32 any) string {
	var __obf_a0c3429734dfadae []string
	__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, "select")
	if len(__obf_ffae9d01a4fb83a6.__obf_298548f7b2ed8c1a) != 0 {
		__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, __obf_525fb93a93b88033(__obf_ffae9d01a4fb83a6.__obf_298548f7b2ed8c1a, SeqComma))
	} else {
		__obf_392c7cb25fd49ba1 := // 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
			// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
			__obf_28c0e19bf1017ee9(__obf_a5b72982c262db32)
		if len(__obf_392c7cb25fd49ba1) > 0 {
			__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, __obf_525fb93a93b88033(__obf_392c7cb25fd49ba1, SeqComma))
		} else {
			__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, "*")
		}
	}
	__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, "from "+__obf_ffae9d01a4fb83a6.__obf_005f563d12feeb25)
	if len(__obf_ffae9d01a4fb83a6.__obf_4210685111b131c8) != 0 {
		__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, __obf_7affa20e9841bed0(__obf_ffae9d01a4fb83a6.__obf_4210685111b131c8, Grouping))
	}

	if __obf_ffae9d01a4fb83a6.__obf_aa3113fef1a0d74d != "" {
		__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, "group by "+__obf_ffae9d01a4fb83a6.__obf_aa3113fef1a0d74d)
	}

	if __obf_ffae9d01a4fb83a6.__obf_7faf07cb163b0164 != "" {
		__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, "having "+__obf_ffae9d01a4fb83a6.__obf_7faf07cb163b0164)
	}

	if __obf_ffae9d01a4fb83a6.__obf_b6b7d1d38659c765 != "" {
		__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, "order by "+__obf_ffae9d01a4fb83a6.__obf_b6b7d1d38659c765)
	}

	if __obf_ffae9d01a4fb83a6.__obf_e11a41cd614462e2 != 0 {
		__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, fmt.Sprintf("limit %d, %d", __obf_ffae9d01a4fb83a6.__obf_f2e5b45bab2a89f5, __obf_ffae9d01a4fb83a6.__obf_e11a41cd614462e2))
	}
	if __obf_ffae9d01a4fb83a6.__obf_e771540f8f5111b7 {
		__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, "lock in share mode")
	}
	if __obf_ffae9d01a4fb83a6.__obf_159039e7777eb165 {
		__obf_a0c3429734dfadae = append(__obf_a0c3429734dfadae, "for update")
	}
	sql := __obf_525fb93a93b88033(__obf_a0c3429734dfadae, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_ffae9d01a4fb83a6.__obf_d51eb37f5f1cb276)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_7affa20e9841bed0(__obf_4210685111b131c8 []string, __obf_42ebc7294355baf2 string) string {
	if len(__obf_4210685111b131c8) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_4210685111b131c8, __obf_42ebc7294355baf2))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_525fb93a93b88033(__obf_d51eb37f5f1cb276 []string, __obf_98789205273002f8 string) string {
	return strings.Join(__obf_d51eb37f5f1cb276,

		// 解析对象中的 `db tag`
		// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
		__obf_98789205273002f8)
}

func __obf_28c0e19bf1017ee9(__obf_a5b72982c262db32 any) (__obf_44baf82a67e2ecb5 []string) {
	__obf_d131126320e759f6 := reflect.ValueOf(__obf_a5b72982c262db32)
	__obf_b7b9c101e0b5767e := __obf_d131126320e759f6.Type().Elem()
	var __obf_912a5b7033efc881 reflect.Type
	if __obf_b7b9c101e0b5767e.Kind() == reflect.Slice {
		__obf_912a5b7033efc881 = __obf_b7b9c101e0b5767e.Elem()
	} else {
		__obf_912a5b7033efc881 = __obf_b7b9c101e0b5767e
	}
	for __obf_33fdeff971fca5a6 := 0; __obf_33fdeff971fca5a6 < __obf_912a5b7033efc881.NumField(); __obf_33fdeff971fca5a6++ {
		__obf_90c653ed7289f13d := __obf_912a5b7033efc881.Field(__obf_33fdeff971fca5a6).Tag.Get(DBTag)
		if __obf_90c653ed7289f13d != "" {
			__obf_44baf82a67e2ecb5 = append(__obf_44baf82a67e2ecb5, __obf_90c653ed7289f13d)
		}
	}
	return
}

func CleanSQLXMap(__obf_2a0af36bd8f8ad61 map[string]any) map[string]any {
	for __obf_299b79baf6e9ee26, __obf_0c15a3a65dfe2b86 := range __obf_2a0af36bd8f8ad61 {
		switch __obf_935e5179094b0bc8 := __obf_0c15a3a65dfe2b86.(type) {
		case []byte:
			__obf_c13a94e56d44ee42 := string(__obf_935e5179094b0bc8)
			__obf_fd1bd346d24427a2 := // 尝试智能识别具体类型
				strings.TrimSpace(__obf_c13a94e56d44ee42)

			// bool
			if strings.EqualFold(__obf_fd1bd346d24427a2, "true") || strings.EqualFold(__obf_fd1bd346d24427a2, "false") {
				if __obf_9ad5a730c8ae4676, __obf_87239f12e0f48308 := strconv.ParseBool(__obf_fd1bd346d24427a2); __obf_87239f12e0f48308 == nil {
					__obf_2a0af36bd8f8ad61[__obf_299b79baf6e9ee26] = __obf_9ad5a730c8ae4676
					continue
				}
			}

			// int
			if __obf_33fdeff971fca5a6, __obf_87239f12e0f48308 := strconv.Atoi(__obf_fd1bd346d24427a2); __obf_87239f12e0f48308 == nil {
				__obf_2a0af36bd8f8ad61[__obf_299b79baf6e9ee26] = __obf_33fdeff971fca5a6
				continue
			}

			// float
			if __obf_ebada1fffc8c92ff, __obf_87239f12e0f48308 := strconv.ParseFloat(__obf_fd1bd346d24427a2, 64); __obf_87239f12e0f48308 == nil {
				__obf_2a0af36bd8f8ad61[__obf_299b79baf6e9ee26] = __obf_ebada1fffc8c92ff
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_fd1bd346d24427a2, "{") && strings.HasSuffix(__obf_fd1bd346d24427a2, "}")) ||
				(strings.HasPrefix(__obf_fd1bd346d24427a2, "[") && strings.HasSuffix(__obf_fd1bd346d24427a2, "]")) {
				var __obf_82fe9e5252f4629f any
				if __obf_87239f12e0f48308 := json.Unmarshal(__obf_935e5179094b0bc8, &__obf_82fe9e5252f4629f); __obf_87239f12e0f48308 == nil {
					__obf_2a0af36bd8f8ad61[__obf_299b79baf6e9ee26] = __obf_82fe9e5252f4629f
					continue
				}
			}
			__obf_2a0af36bd8f8ad61[ // 默认保留为 string
			__obf_299b79baf6e9ee26] = __obf_c13a94e56d44ee42
		default:
			__obf_2a0af36bd8f8ad61[__obf_299b79baf6e9ee26] = __obf_935e5179094b0bc8
		}
	}
	return __obf_2a0af36bd8f8ad61
}

func ToListMap(__obf_6751482a63a2d802 *sqlx.Rows) []map[string]any {
	var __obf_d46b9a4ac0947f8a []map[string]any
	for __obf_6751482a63a2d802.Next() {
		__obf_db864045c480f50e := make(map[string]any)
		if __obf_87239f12e0f48308 := __obf_6751482a63a2d802.MapScan(__obf_db864045c480f50e); __obf_87239f12e0f48308 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_87239f12e0f48308.Error()))
			continue
		}
		__obf_d46b9a4ac0947f8a = append(__obf_d46b9a4ac0947f8a, CleanSQLXMap(__obf_db864045c480f50e))
	}
	return __obf_d46b9a4ac0947f8a
}
