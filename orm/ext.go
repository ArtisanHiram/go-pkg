package __obf_ed189c965cdcd132

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
type FuncTx func(__obf_83eebeb3bc1b8de9 *sqlx.Tx, __obf_8670d196c9bebfda any) error

func Open(__obf_d620e87a38f058f1, __obf_3471b74c716ab7f3 string, __obf_aa022e947d9d8192 time.Duration) (*DB, error) {
	__obf_e17d8f0e9734b0de, __obf_9f5a64d352151704 := sqlx.Open(__obf_d620e87a38f058f1, __obf_3471b74c716ab7f3)
	if __obf_9f5a64d352151704 != nil {
		return nil, __obf_9f5a64d352151704
	}
	__obf_e1e03f4986eb2a8d := &DB{
		DB: __obf_e17d8f0e9734b0de, __obf_aa022e947d9d8192: __obf_aa022e947d9d8192,
	}
	__obf_e1e03f4986eb2a8d.__obf_daa5f07e389ba355.
		New = func() any {
		return __obf_e1e03f4986eb2a8d.__obf_40cd0666316949f8()
	}
	return __obf_e1e03f4986eb2a8d, nil
}

type DB struct {
	*sqlx.DB
	__obf_aa022e947d9d8192 time.Duration
	__obf_daa5f07e389ba355 sync.Pool
}

func (__obf_e17d8f0e9734b0de *DB) __obf_40cd0666316949f8() *Context {
	return &Context{__obf_e17d8f0e9734b0de: __obf_e17d8f0e9734b0de}
}

// 获取一个`SQL`执行`Context`
func (__obf_e17d8f0e9734b0de *DB) Acquire() *Context {
	__obf_4aa3842b7a1e7983 := // 无需加锁，sync.Pool本身是线程安全的
		__obf_e17d8f0e9734b0de.__obf_daa5f07e389ba355.Get().(*Context)
	__obf_4aa3842b7a1e7983.__obf_5b1c2133075e9601()
	return __obf_4aa3842b7a1e7983
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_e17d8f0e9734b0de *DB) AcquireTx(__obf_83eebeb3bc1b8de9 *sqlx.Tx) *Context {
	__obf_4aa3842b7a1e7983 := __obf_e17d8f0e9734b0de.Acquire()
	__obf_4aa3842b7a1e7983.__obf_83eebeb3bc1b8de9 = __obf_83eebeb3bc1b8de9
	return __obf_4aa3842b7a1e7983
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_e17d8f0e9734b0de *DB) WithTx(__obf_339417fab022652b FuncTx, __obf_8670d196c9bebfda any) (__obf_9f5a64d352151704 error) {
	var __obf_83eebeb3bc1b8de9 *sqlx.Tx
	__obf_83eebeb3bc1b8de9, __obf_9f5a64d352151704 = __obf_e17d8f0e9734b0de.Beginx()
	if __obf_9f5a64d352151704 != nil {
		return
	}
	defer func() {
		if __obf_9f5a64d352151704 != nil && __obf_83eebeb3bc1b8de9 != nil {
			__obf_9f5a64d352151704 = __obf_83eebeb3bc1b8de9.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_9f5a64d352151704 = __obf_339417fab022652b(__obf_83eebeb3bc1b8de9, __obf_8670d196c9bebfda); __obf_9f5a64d352151704 != nil {
		return
	}
	__obf_9f5a64d352151704 = __obf_83eebeb3bc1b8de9.Commit()
	return
}

type Context struct {
	__obf_e17d8f0e9734b0de *DB
	__obf_83eebeb3bc1b8de9 *sqlx.Tx //事务
	sql                    string
	__obf_584cc9bde8d28fef string
	__obf_5ff8657915367f7f []string
	__obf_75fe7f4f4288a958 []string
	__obf_360859ed4cfca4e4 string
	__obf_7440e1fcb72dea01 string
	__obf_da981eaa39b23de2 string
	__obf_9fd9365a03a7c2e3 int64
	__obf_99d5fa2d15da6b60 int64
	__obf_8670d196c9bebfda []any
	__obf_6df264e7b2ee8673 bool
	__obf_f7855e2674bef899 bool //排他锁
	//共享锁
}

func (__obf_4aa3842b7a1e7983 *Context) Name(__obf_584cc9bde8d28fef string) *Context {
	__obf_4aa3842b7a1e7983.__obf_584cc9bde8d28fef = __obf_584cc9bde8d28fef
	return __obf_4aa3842b7a1e7983
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_4aa3842b7a1e7983 *Context) What(__obf_5ff8657915367f7f []string) *Context {
	__obf_4aa3842b7a1e7983.__obf_5ff8657915367f7f = __obf_5ff8657915367f7f
	return __obf_4aa3842b7a1e7983
}

func (__obf_4aa3842b7a1e7983 *Context) Where(__obf_055a197727e3c792 string, __obf_8670d196c9bebfda ...any) *Context {
	__obf_4aa3842b7a1e7983.__obf_75fe7f4f4288a958 = append(__obf_4aa3842b7a1e7983.__obf_75fe7f4f4288a958, __obf_055a197727e3c792)
	__obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda = append(__obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda,

		// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
		__obf_8670d196c9bebfda...)
	return __obf_4aa3842b7a1e7983
}

func (__obf_4aa3842b7a1e7983 *Context) WhereIn(__obf_8e630cdb3c3b799f string, __obf_8670d196c9bebfda []any) *Context {
	__obf_0560ee77d623dc5f := len(__obf_8670d196c9bebfda)
	__obf_33cf795b6273cb25 := make([]string, __obf_0560ee77d623dc5f)
	for __obf_c7e5bacf18525a19 := 0; __obf_c7e5bacf18525a19 < __obf_0560ee77d623dc5f; __obf_c7e5bacf18525a19++ {
		__obf_33cf795b6273cb25[__obf_c7e5bacf18525a19] = ParamMarker
	}
	__obf_b0264f01d82aed13 := fmt.Sprintf("%s in (%s)", __obf_8e630cdb3c3b799f, __obf_6823befc7747cba6(__obf_33cf795b6273cb25, SeqComma))
	return __obf_4aa3842b7a1e7983.Where(__obf_b0264f01d82aed13, __obf_8670d196c9bebfda...)
}

func (__obf_4aa3842b7a1e7983 *Context) Order(__obf_360859ed4cfca4e4 string) *Context {
	__obf_4aa3842b7a1e7983.__obf_360859ed4cfca4e4 = __obf_360859ed4cfca4e4
	return __obf_4aa3842b7a1e7983
}

func (__obf_4aa3842b7a1e7983 *Context) Limit(__obf_9fd9365a03a7c2e3 int64) *Context {
	__obf_4aa3842b7a1e7983.__obf_9fd9365a03a7c2e3 = __obf_9fd9365a03a7c2e3
	return __obf_4aa3842b7a1e7983
}

func (__obf_4aa3842b7a1e7983 *Context) Offset(__obf_99d5fa2d15da6b60 int64) *Context {
	__obf_4aa3842b7a1e7983.__obf_99d5fa2d15da6b60 = __obf_99d5fa2d15da6b60
	return __obf_4aa3842b7a1e7983
}

func (__obf_4aa3842b7a1e7983 *Context) Group(__obf_7440e1fcb72dea01 string) *Context {
	__obf_4aa3842b7a1e7983.__obf_7440e1fcb72dea01 = __obf_7440e1fcb72dea01
	return __obf_4aa3842b7a1e7983
}

func (__obf_4aa3842b7a1e7983 *Context) Having(__obf_da981eaa39b23de2 string, __obf_8670d196c9bebfda ...any) *Context {
	__obf_4aa3842b7a1e7983.__obf_da981eaa39b23de2 = __obf_da981eaa39b23de2
	__obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda = append(__obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda,

		// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
		__obf_8670d196c9bebfda...)
	return __obf_4aa3842b7a1e7983
}

func (__obf_4aa3842b7a1e7983 *Context) LockX() *Context {
	__obf_4aa3842b7a1e7983.__obf_6df264e7b2ee8673 = true
	return __obf_4aa3842b7a1e7983
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_4aa3842b7a1e7983 *Context) LockS() *Context {
	__obf_4aa3842b7a1e7983.__obf_f7855e2674bef899 = true
	return __obf_4aa3842b7a1e7983
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_4aa3842b7a1e7983 *Context) FindMany(__obf_86d96f1d3a893832 any) error {
	return __obf_4aa3842b7a1e7983.__obf_f170de3deba02c2c(__obf_86d96f1d3a893832, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_4aa3842b7a1e7983 *Context) FindById(__obf_86d96f1d3a893832 any) error {
	return __obf_4aa3842b7a1e7983.__obf_f170de3deba02c2c(__obf_86d96f1d3a893832, SelectTypeOne)
}

// 插入
func (__obf_4aa3842b7a1e7983 *Context) Insert(__obf_01c2af9e93ca6bb9 map[string]any) (sql.Result, error) {
	var (
		__obf_9b0267b79e4b328b []string
		__obf_69403baf4045890c []any
	)
	for __obf_cb6df4213d1d222d, __obf_604eba2487c8ea9b := range __obf_01c2af9e93ca6bb9 {
		__obf_9b0267b79e4b328b = append(__obf_9b0267b79e4b328b, __obf_cb6df4213d1d222d)
		__obf_69403baf4045890c = append(__obf_69403baf4045890c, __obf_604eba2487c8ea9b)
	}
	return __obf_4aa3842b7a1e7983.InsertBatch(__obf_9b0267b79e4b328b,

		// 批量插入
		__obf_69403baf4045890c)
}

func (__obf_4aa3842b7a1e7983 *Context) InsertBatch(__obf_9b0267b79e4b328b []string, __obf_01c2af9e93ca6bb9 ...[]any) (sql.Result, error) {
	var (
		__obf_69403baf4045890c []any
		__obf_e33aad30a146f7b6 []string
	)
	for _, __obf_2bf1e56fb1284f35 := range __obf_01c2af9e93ca6bb9 {
		__obf_33cf795b6273cb25 := make([]string, len(__obf_2bf1e56fb1284f35))
		for __obf_c7e5bacf18525a19, __obf_604eba2487c8ea9b := range __obf_2bf1e56fb1284f35 {
			__obf_33cf795b6273cb25[__obf_c7e5bacf18525a19] = ParamMarker
			__obf_69403baf4045890c = append(__obf_69403baf4045890c, __obf_604eba2487c8ea9b)
		}
		__obf_e33aad30a146f7b6 = append(__obf_e33aad30a146f7b6, fmt.Sprintf("(%s)", __obf_6823befc7747cba6(__obf_33cf795b6273cb25, SeqComma)))
	}
	__obf_91b87182c3b48e5a := fmt.Sprintf("insert into %s (%s) values %s", __obf_4aa3842b7a1e7983.__obf_584cc9bde8d28fef, __obf_6823befc7747cba6(__obf_9b0267b79e4b328b, SeqComma), __obf_6823befc7747cba6(__obf_e33aad30a146f7b6, SeqComma))
	return __obf_4aa3842b7a1e7983.__obf_ab853f04525c61e1(__obf_91b87182c3b48e5a,

		// 使用map更新
		__obf_69403baf4045890c...)
}

func (__obf_4aa3842b7a1e7983 *Context) UpdateMap(__obf_8670d196c9bebfda map[string]any) (__obf_f9aa917e28271006 int64, __obf_9f5a64d352151704 error) {
	var (
		__obf_69403baf4045890c []any
		__obf_4a9aaeaadbee4340 []string
	)
	for __obf_cb6df4213d1d222d, __obf_604eba2487c8ea9b := range __obf_8670d196c9bebfda {
		__obf_69403baf4045890c = append(__obf_69403baf4045890c, __obf_604eba2487c8ea9b)
		__obf_4a9aaeaadbee4340 = append(__obf_4a9aaeaadbee4340, fmt.Sprintf("%s=%s", __obf_cb6df4213d1d222d, ParamMarker))
	}
	__obf_ebec791607f6c641 := __obf_6823befc7747cba6(__obf_4a9aaeaadbee4340, SeqComma)
	__obf_f9aa917e28271006, __obf_9f5a64d352151704 = __obf_4aa3842b7a1e7983.Update(__obf_ebec791607f6c641, __obf_69403baf4045890c...)
	return
}

// 更新
func (__obf_4aa3842b7a1e7983 *Context) Update(__obf_ebec791607f6c641 string, __obf_8670d196c9bebfda ...any) (__obf_f9aa917e28271006 int64, __obf_9f5a64d352151704 error) {
	__obf_f4591c07b714b6bd := "update %s set %s %s"
	__obf_055a197727e3c792 := __obf_8ce3e91b3e5ad2b2(__obf_4aa3842b7a1e7983.__obf_75fe7f4f4288a958, Grouping)
	__obf_91b87182c3b48e5a := fmt.Sprintf(__obf_f4591c07b714b6bd, __obf_4aa3842b7a1e7983.__obf_584cc9bde8d28fef, __obf_ebec791607f6c641, __obf_055a197727e3c792)
	__obf_69403baf4045890c := append(__obf_8670d196c9bebfda, __obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda...)
	var __obf_3bf64be4135c172e sql.Result
	__obf_3bf64be4135c172e, __obf_9f5a64d352151704 = __obf_4aa3842b7a1e7983.__obf_ab853f04525c61e1(__obf_91b87182c3b48e5a, __obf_69403baf4045890c...)
	if __obf_9f5a64d352151704 != nil {
		return
	}
	__obf_f9aa917e28271006, __obf_9f5a64d352151704 = __obf_3bf64be4135c172e.RowsAffected()
	return
}

// 删除
func (__obf_4aa3842b7a1e7983 *Context) Delete() (__obf_f9aa917e28271006 int64, __obf_9f5a64d352151704 error) {
	__obf_f4591c07b714b6bd := "delete from %s %s"
	__obf_055a197727e3c792 := __obf_8ce3e91b3e5ad2b2(__obf_4aa3842b7a1e7983.__obf_75fe7f4f4288a958, Grouping)
	__obf_91b87182c3b48e5a := fmt.Sprintf(__obf_f4591c07b714b6bd, __obf_4aa3842b7a1e7983.__obf_584cc9bde8d28fef, __obf_055a197727e3c792)
	var __obf_3bf64be4135c172e sql.Result
	__obf_3bf64be4135c172e, __obf_9f5a64d352151704 = __obf_4aa3842b7a1e7983.__obf_ab853f04525c61e1(__obf_91b87182c3b48e5a, __obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda...)
	if __obf_9f5a64d352151704 != nil {
		return
	}
	__obf_f9aa917e28271006, __obf_9f5a64d352151704 = __obf_3bf64be4135c172e.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_4aa3842b7a1e7983 *Context) Select(__obf_86d96f1d3a893832 any, sql string, __obf_8670d196c9bebfda ...any) error {
	__obf_4aa3842b7a1e7983.
		sql = sql
	__obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda = __obf_8670d196c9bebfda
	return __obf_4aa3842b7a1e7983.__obf_f170de3deba02c2c(__obf_86d96f1d3a893832, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_4aa3842b7a1e7983 *Context) Get(__obf_86d96f1d3a893832 any, sql string, __obf_8670d196c9bebfda ...any) error {
	__obf_4aa3842b7a1e7983.
		sql = sql
	__obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda = __obf_8670d196c9bebfda
	return __obf_4aa3842b7a1e7983.__obf_f170de3deba02c2c(__obf_86d96f1d3a893832, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_4aa3842b7a1e7983 *Context) Exec(sql string, __obf_8670d196c9bebfda ...any) (sql.Result, error) {
	return __obf_4aa3842b7a1e7983.__obf_ab853f04525c61e1(sql, __obf_8670d196c9bebfda...)
}

// 创建表
func (__obf_4aa3842b7a1e7983 *Context) Create(sql string) (sql.Result, error) {
	return __obf_4aa3842b7a1e7983.

		// 删除表
		__obf_ab853f04525c61e1(sql)
}

func (__obf_4aa3842b7a1e7983 *Context) Drop() (sql.Result, error) {
	return __obf_4aa3842b7a1e7983.__obf_ab853f04525c61e1(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_4aa3842b7a1e7983.

		/////////////////////////private methods//////////////////////
		__obf_584cc9bde8d28fef))
}

// 重置Context
func (__obf_4aa3842b7a1e7983 *Context) __obf_5b1c2133075e9601() *Context {
	__obf_4aa3842b7a1e7983.
		sql = ""
	__obf_4aa3842b7a1e7983.__obf_584cc9bde8d28fef = ""
	__obf_4aa3842b7a1e7983.__obf_5ff8657915367f7f = []string{}
	__obf_4aa3842b7a1e7983.__obf_75fe7f4f4288a958 = []string{}
	__obf_4aa3842b7a1e7983.__obf_360859ed4cfca4e4 = ""
	__obf_4aa3842b7a1e7983.__obf_7440e1fcb72dea01 = ""
	__obf_4aa3842b7a1e7983.__obf_da981eaa39b23de2 = ""
	__obf_4aa3842b7a1e7983.__obf_9fd9365a03a7c2e3 = 0
	__obf_4aa3842b7a1e7983.__obf_99d5fa2d15da6b60 = 0
	__obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda = []any{}
	__obf_4aa3842b7a1e7983.__obf_83eebeb3bc1b8de9 = nil
	__obf_4aa3842b7a1e7983.__obf_f7855e2674bef899 = false
	__obf_4aa3842b7a1e7983.__obf_6df264e7b2ee8673 = false
	return __obf_4aa3842b7a1e7983
}

// 查询方法
func (__obf_4aa3842b7a1e7983 *Context) __obf_f170de3deba02c2c(__obf_86d96f1d3a893832 any, __obf_0a0a6c3cec55b68b int) (__obf_9f5a64d352151704 error) {
	defer __obf_4aa3842b7a1e7983.__obf_e17d8f0e9734b0de.__obf_daa5f07e389ba355.Put(__obf_4aa3842b7a1e7983)
	__obf_e1b66d66c214138a, __obf_890a73ac1a1bf9b8 := context.WithTimeout(context.Background(), __obf_4aa3842b7a1e7983.__obf_e17d8f0e9734b0de.__obf_aa022e947d9d8192)
	defer __obf_890a73ac1a1bf9b8()
	if __obf_4aa3842b7a1e7983.sql == "" {
		__obf_4aa3842b7a1e7983.
			sql = __obf_4aa3842b7a1e7983.__obf_e780a86bacce22ab(__obf_86d96f1d3a893832)
	}
	switch __obf_0a0a6c3cec55b68b {
	case SelectTypeOne:
		if __obf_4aa3842b7a1e7983.__obf_83eebeb3bc1b8de9 != nil {
			__obf_9f5a64d352151704 = __obf_4aa3842b7a1e7983.__obf_83eebeb3bc1b8de9.GetContext(__obf_e1b66d66c214138a, __obf_86d96f1d3a893832, __obf_4aa3842b7a1e7983.sql, __obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda...)
		} else {
			__obf_9f5a64d352151704 = __obf_4aa3842b7a1e7983.__obf_e17d8f0e9734b0de.GetContext(__obf_e1b66d66c214138a, __obf_86d96f1d3a893832, __obf_4aa3842b7a1e7983.sql, __obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda...)
		}
	case SelectTypeMany:
		if __obf_4aa3842b7a1e7983.__obf_83eebeb3bc1b8de9 != nil {
			__obf_9f5a64d352151704 = __obf_4aa3842b7a1e7983.__obf_83eebeb3bc1b8de9.SelectContext(__obf_e1b66d66c214138a, __obf_86d96f1d3a893832, __obf_4aa3842b7a1e7983.sql, __obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda...)
		} else {
			__obf_9f5a64d352151704 = __obf_4aa3842b7a1e7983.__obf_e17d8f0e9734b0de.SelectContext(__obf_e1b66d66c214138a, __obf_86d96f1d3a893832, __obf_4aa3842b7a1e7983.sql, __obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_4aa3842b7a1e7983 *Context) __obf_ab853f04525c61e1(__obf_91b87182c3b48e5a string, __obf_8670d196c9bebfda ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_91b87182c3b48e5a, __obf_8670d196c9bebfda)
	defer __obf_4aa3842b7a1e7983.__obf_e17d8f0e9734b0de.__obf_daa5f07e389ba355.Put(__obf_4aa3842b7a1e7983)
	__obf_e1b66d66c214138a, __obf_890a73ac1a1bf9b8 := context.WithTimeout(context.Background(), __obf_4aa3842b7a1e7983.__obf_e17d8f0e9734b0de.__obf_aa022e947d9d8192)
	defer __obf_890a73ac1a1bf9b8()

	var __obf_6d4c7ae02e7c89dc sqlx.ExecerContext
	if __obf_4aa3842b7a1e7983.__obf_83eebeb3bc1b8de9 != nil {
		__obf_6d4c7ae02e7c89dc = __obf_4aa3842b7a1e7983.__obf_83eebeb3bc1b8de9
	} else {
		__obf_6d4c7ae02e7c89dc = __obf_4aa3842b7a1e7983.__obf_e17d8f0e9734b0de
	}
	return __obf_6d4c7ae02e7c89dc.ExecContext(__obf_e1b66d66c214138a, __obf_91b87182c3b48e5a,

		// select查询语句的拼接
		__obf_8670d196c9bebfda...)
}

func (__obf_4aa3842b7a1e7983 *Context) __obf_e780a86bacce22ab(__obf_86d96f1d3a893832 any) string {
	var __obf_446e01e526d9fb8e []string
	__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, "select")
	if len(__obf_4aa3842b7a1e7983.__obf_5ff8657915367f7f) != 0 {
		__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, __obf_6823befc7747cba6(__obf_4aa3842b7a1e7983.__obf_5ff8657915367f7f, SeqComma))
	} else {
		__obf_d1b1526538353aa1 := // 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
			// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
			__obf_a15ae184bc2fc634(__obf_86d96f1d3a893832)
		if len(__obf_d1b1526538353aa1) > 0 {
			__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, __obf_6823befc7747cba6(__obf_d1b1526538353aa1, SeqComma))
		} else {
			__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, "*")
		}
	}
	__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, "from "+__obf_4aa3842b7a1e7983.__obf_584cc9bde8d28fef)
	if len(__obf_4aa3842b7a1e7983.__obf_75fe7f4f4288a958) != 0 {
		__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, __obf_8ce3e91b3e5ad2b2(__obf_4aa3842b7a1e7983.__obf_75fe7f4f4288a958, Grouping))
	}

	if __obf_4aa3842b7a1e7983.__obf_7440e1fcb72dea01 != "" {
		__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, "group by "+__obf_4aa3842b7a1e7983.__obf_7440e1fcb72dea01)
	}

	if __obf_4aa3842b7a1e7983.__obf_da981eaa39b23de2 != "" {
		__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, "having "+__obf_4aa3842b7a1e7983.__obf_da981eaa39b23de2)
	}

	if __obf_4aa3842b7a1e7983.__obf_360859ed4cfca4e4 != "" {
		__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, "order by "+__obf_4aa3842b7a1e7983.__obf_360859ed4cfca4e4)
	}

	if __obf_4aa3842b7a1e7983.__obf_9fd9365a03a7c2e3 != 0 {
		__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, fmt.Sprintf("limit %d, %d", __obf_4aa3842b7a1e7983.__obf_99d5fa2d15da6b60, __obf_4aa3842b7a1e7983.__obf_9fd9365a03a7c2e3))
	}
	if __obf_4aa3842b7a1e7983.__obf_f7855e2674bef899 {
		__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, "lock in share mode")
	}
	if __obf_4aa3842b7a1e7983.__obf_6df264e7b2ee8673 {
		__obf_446e01e526d9fb8e = append(__obf_446e01e526d9fb8e, "for update")
	}
	sql := __obf_6823befc7747cba6(__obf_446e01e526d9fb8e, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_4aa3842b7a1e7983.__obf_8670d196c9bebfda)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_8ce3e91b3e5ad2b2(__obf_75fe7f4f4288a958 []string, __obf_5ba64c9a63c055e3 string) string {
	if len(__obf_75fe7f4f4288a958) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_75fe7f4f4288a958, __obf_5ba64c9a63c055e3))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_6823befc7747cba6(__obf_8670d196c9bebfda []string, __obf_0fc6378e37985bd9 string) string {
	return strings.Join(__obf_8670d196c9bebfda,

		// 解析对象中的 `db tag`
		// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
		__obf_0fc6378e37985bd9)
}

func __obf_a15ae184bc2fc634(__obf_86d96f1d3a893832 any) (__obf_9b0267b79e4b328b []string) {
	__obf_9c99160fd6d9fa7d := reflect.ValueOf(__obf_86d96f1d3a893832)
	__obf_696c68042f724dfb := __obf_9c99160fd6d9fa7d.Type().Elem()
	var __obf_b9de2d8368b11e30 reflect.Type
	if __obf_696c68042f724dfb.Kind() == reflect.Slice {
		__obf_b9de2d8368b11e30 = __obf_696c68042f724dfb.Elem()
	} else {
		__obf_b9de2d8368b11e30 = __obf_696c68042f724dfb
	}
	for __obf_c7e5bacf18525a19 := 0; __obf_c7e5bacf18525a19 < __obf_b9de2d8368b11e30.NumField(); __obf_c7e5bacf18525a19++ {
		__obf_6422965122b91db6 := __obf_b9de2d8368b11e30.Field(__obf_c7e5bacf18525a19).Tag.Get(DBTag)
		if __obf_6422965122b91db6 != "" {
			__obf_9b0267b79e4b328b = append(__obf_9b0267b79e4b328b, __obf_6422965122b91db6)
		}
	}
	return
}

func CleanSQLXMap(__obf_8f80623a98f06b7b map[string]any) map[string]any {
	for __obf_cb6df4213d1d222d, __obf_604eba2487c8ea9b := range __obf_8f80623a98f06b7b {
		switch __obf_f9cbd174629de499 := __obf_604eba2487c8ea9b.(type) {
		case []byte:
			__obf_1e3af5b8f467d03d := string(__obf_f9cbd174629de499)
			__obf_d0ea745c4fdfa83f := // 尝试智能识别具体类型
				strings.TrimSpace(__obf_1e3af5b8f467d03d)

			// bool
			if strings.EqualFold(__obf_d0ea745c4fdfa83f, "true") || strings.EqualFold(__obf_d0ea745c4fdfa83f, "false") {
				if __obf_0064b82837abd18e, __obf_9f5a64d352151704 := strconv.ParseBool(__obf_d0ea745c4fdfa83f); __obf_9f5a64d352151704 == nil {
					__obf_8f80623a98f06b7b[__obf_cb6df4213d1d222d] = __obf_0064b82837abd18e
					continue
				}
			}

			// int
			if __obf_c7e5bacf18525a19, __obf_9f5a64d352151704 := strconv.Atoi(__obf_d0ea745c4fdfa83f); __obf_9f5a64d352151704 == nil {
				__obf_8f80623a98f06b7b[__obf_cb6df4213d1d222d] = __obf_c7e5bacf18525a19
				continue
			}

			// float
			if __obf_626ed35c11f71017, __obf_9f5a64d352151704 := strconv.ParseFloat(__obf_d0ea745c4fdfa83f, 64); __obf_9f5a64d352151704 == nil {
				__obf_8f80623a98f06b7b[__obf_cb6df4213d1d222d] = __obf_626ed35c11f71017
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_d0ea745c4fdfa83f, "{") && strings.HasSuffix(__obf_d0ea745c4fdfa83f, "}")) ||
				(strings.HasPrefix(__obf_d0ea745c4fdfa83f, "[") && strings.HasSuffix(__obf_d0ea745c4fdfa83f, "]")) {
				var __obf_2c02a256ed79e82c any
				if __obf_9f5a64d352151704 := json.Unmarshal(__obf_f9cbd174629de499, &__obf_2c02a256ed79e82c); __obf_9f5a64d352151704 == nil {
					__obf_8f80623a98f06b7b[__obf_cb6df4213d1d222d] = __obf_2c02a256ed79e82c
					continue
				}
			}
			__obf_8f80623a98f06b7b[ // 默认保留为 string
			__obf_cb6df4213d1d222d] = __obf_1e3af5b8f467d03d
		default:
			__obf_8f80623a98f06b7b[__obf_cb6df4213d1d222d] = __obf_f9cbd174629de499
		}
	}
	return __obf_8f80623a98f06b7b
}

func ToListMap(__obf_b9da22b0e5936af1 *sqlx.Rows) []map[string]any {
	var __obf_3bf64be4135c172e []map[string]any
	for __obf_b9da22b0e5936af1.Next() {
		__obf_8a777ea989d024ab := make(map[string]any)
		if __obf_9f5a64d352151704 := __obf_b9da22b0e5936af1.MapScan(__obf_8a777ea989d024ab); __obf_9f5a64d352151704 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_9f5a64d352151704.Error()))
			continue
		}
		__obf_3bf64be4135c172e = append(__obf_3bf64be4135c172e, CleanSQLXMap(__obf_8a777ea989d024ab))
	}
	return __obf_3bf64be4135c172e
}
