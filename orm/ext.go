package __obf_df9e37b4dd16fa57

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
type FuncTx func(__obf_4e1da10c2afe56d6 *sqlx.Tx, __obf_3cef69abfd4b5fed any) error

func Open(__obf_3e1e1a445db62cdf, __obf_bc181e5ba2602fb2 string, __obf_435469c73588f1a2 time.Duration) (*DB, error) {
	__obf_ed3a9a8432b2d457, __obf_4795d167371f16d4 := sqlx.Open(__obf_3e1e1a445db62cdf, __obf_bc181e5ba2602fb2)
	if __obf_4795d167371f16d4 != nil {
		return nil, __obf_4795d167371f16d4
	}
	__obf_8baa416e653862dd := &DB{
		DB:                     __obf_ed3a9a8432b2d457,
		__obf_435469c73588f1a2: __obf_435469c73588f1a2,
	}
	__obf_8baa416e653862dd.__obf_d6fbd616396c8e75.New = func() any {
		return __obf_8baa416e653862dd.__obf_c6402afb25903cc1()
	}
	return __obf_8baa416e653862dd, nil
}

type DB struct {
	*sqlx.DB
	__obf_435469c73588f1a2 time.Duration
	__obf_d6fbd616396c8e75 sync.Pool
}

func (__obf_ed3a9a8432b2d457 *DB) __obf_c6402afb25903cc1() *Context {
	return &Context{__obf_ed3a9a8432b2d457: __obf_ed3a9a8432b2d457}
}

// 获取一个`SQL`执行`Context`
func (__obf_ed3a9a8432b2d457 *DB) Acquire() *Context {
	// 无需加锁，sync.Pool本身是线程安全的
	__obf_036c6b778079b76b := __obf_ed3a9a8432b2d457.__obf_d6fbd616396c8e75.Get().(*Context)
	__obf_036c6b778079b76b.__obf_79a134d3318127c7()
	return __obf_036c6b778079b76b
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_ed3a9a8432b2d457 *DB) AcquireTx(__obf_4e1da10c2afe56d6 *sqlx.Tx) *Context {
	__obf_036c6b778079b76b := __obf_ed3a9a8432b2d457.Acquire()
	__obf_036c6b778079b76b.__obf_4e1da10c2afe56d6 = __obf_4e1da10c2afe56d6
	return __obf_036c6b778079b76b
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_ed3a9a8432b2d457 *DB) WithTx(__obf_0c6f3b0568984f99 FuncTx, __obf_3cef69abfd4b5fed any) (__obf_4795d167371f16d4 error) {
	var __obf_4e1da10c2afe56d6 *sqlx.Tx
	__obf_4e1da10c2afe56d6, __obf_4795d167371f16d4 = __obf_ed3a9a8432b2d457.Beginx()
	if __obf_4795d167371f16d4 != nil {
		return
	}
	defer func() {
		if __obf_4795d167371f16d4 != nil && __obf_4e1da10c2afe56d6 != nil {
			__obf_4795d167371f16d4 = __obf_4e1da10c2afe56d6.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_4795d167371f16d4 = __obf_0c6f3b0568984f99(__obf_4e1da10c2afe56d6, __obf_3cef69abfd4b5fed); __obf_4795d167371f16d4 != nil {
		return
	}

	__obf_4795d167371f16d4 = __obf_4e1da10c2afe56d6.Commit()
	return
}

type Context struct {
	__obf_ed3a9a8432b2d457 *DB
	__obf_4e1da10c2afe56d6 *sqlx.Tx //事务
	sql                    string
	__obf_5bf373809aa21efc string
	__obf_375b71252e6f26c6 []string
	__obf_8649156a558d0fb4 []string
	__obf_72aa9c9e212e03fe string
	__obf_79a5b98b6cf49502 string
	__obf_3901bcf69ddf5fea string
	__obf_159c787549b63f29 int64
	__obf_b9906ca51574c61e int64
	__obf_3cef69abfd4b5fed []any
	__obf_04522ec63d23e18d bool //排他锁
	__obf_dacd8adf705f369a bool //共享锁
}

func (__obf_036c6b778079b76b *Context) Name(__obf_5bf373809aa21efc string) *Context {
	__obf_036c6b778079b76b.__obf_5bf373809aa21efc = __obf_5bf373809aa21efc
	return __obf_036c6b778079b76b
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_036c6b778079b76b *Context) What(__obf_375b71252e6f26c6 []string) *Context {
	__obf_036c6b778079b76b.__obf_375b71252e6f26c6 = __obf_375b71252e6f26c6
	return __obf_036c6b778079b76b
}

func (__obf_036c6b778079b76b *Context) Where(__obf_511a4b0e43c6c0f2 string, __obf_3cef69abfd4b5fed ...any) *Context {
	__obf_036c6b778079b76b.__obf_8649156a558d0fb4 = append(__obf_036c6b778079b76b.__obf_8649156a558d0fb4, __obf_511a4b0e43c6c0f2)
	__obf_036c6b778079b76b.__obf_3cef69abfd4b5fed = append(__obf_036c6b778079b76b.__obf_3cef69abfd4b5fed, __obf_3cef69abfd4b5fed...)
	return __obf_036c6b778079b76b
}

// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
func (__obf_036c6b778079b76b *Context) WhereIn(__obf_a5ff5bba00286eca string, __obf_3cef69abfd4b5fed []any) *Context {
	__obf_b30c16b0f3bef92f := len(__obf_3cef69abfd4b5fed)
	__obf_4108d94ccc3925fa := make([]string, __obf_b30c16b0f3bef92f)
	for __obf_2a7388295d503598 := 0; __obf_2a7388295d503598 < __obf_b30c16b0f3bef92f; __obf_2a7388295d503598++ {
		__obf_4108d94ccc3925fa[__obf_2a7388295d503598] = ParamMarker
	}
	__obf_832ab91dc904a7c6 := fmt.Sprintf("%s in (%s)", __obf_a5ff5bba00286eca, __obf_b3ceeeee12719402(__obf_4108d94ccc3925fa, SeqComma))
	return __obf_036c6b778079b76b.Where(__obf_832ab91dc904a7c6, __obf_3cef69abfd4b5fed...)
}

func (__obf_036c6b778079b76b *Context) Order(__obf_72aa9c9e212e03fe string) *Context {
	__obf_036c6b778079b76b.__obf_72aa9c9e212e03fe = __obf_72aa9c9e212e03fe
	return __obf_036c6b778079b76b
}

func (__obf_036c6b778079b76b *Context) Limit(__obf_159c787549b63f29 int64) *Context {
	__obf_036c6b778079b76b.__obf_159c787549b63f29 = __obf_159c787549b63f29
	return __obf_036c6b778079b76b
}

func (__obf_036c6b778079b76b *Context) Offset(__obf_b9906ca51574c61e int64) *Context {
	__obf_036c6b778079b76b.__obf_b9906ca51574c61e = __obf_b9906ca51574c61e
	return __obf_036c6b778079b76b
}

func (__obf_036c6b778079b76b *Context) Group(__obf_79a5b98b6cf49502 string) *Context {
	__obf_036c6b778079b76b.__obf_79a5b98b6cf49502 = __obf_79a5b98b6cf49502
	return __obf_036c6b778079b76b
}

func (__obf_036c6b778079b76b *Context) Having(__obf_3901bcf69ddf5fea string, __obf_3cef69abfd4b5fed ...any) *Context {
	__obf_036c6b778079b76b.__obf_3901bcf69ddf5fea = __obf_3901bcf69ddf5fea
	__obf_036c6b778079b76b.__obf_3cef69abfd4b5fed = append(__obf_036c6b778079b76b.__obf_3cef69abfd4b5fed, __obf_3cef69abfd4b5fed...)
	return __obf_036c6b778079b76b
}

// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
func (__obf_036c6b778079b76b *Context) LockX() *Context {
	__obf_036c6b778079b76b.__obf_04522ec63d23e18d = true
	return __obf_036c6b778079b76b
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_036c6b778079b76b *Context) LockS() *Context {
	__obf_036c6b778079b76b.__obf_dacd8adf705f369a = true
	return __obf_036c6b778079b76b
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_036c6b778079b76b *Context) FindMany(__obf_ecfa1ba7846c04d3 any) error {
	return __obf_036c6b778079b76b.__obf_fde655c5896dd5ee(__obf_ecfa1ba7846c04d3, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_036c6b778079b76b *Context) FindById(__obf_ecfa1ba7846c04d3 any) error {
	return __obf_036c6b778079b76b.__obf_fde655c5896dd5ee(__obf_ecfa1ba7846c04d3, SelectTypeOne)
}

// 插入
func (__obf_036c6b778079b76b *Context) Insert(__obf_2b62c53fec3ac5a3 map[string]any) (sql.Result, error) {
	var (
		__obf_cabff7cb0e419137 []string
		__obf_b4f45d1add3ff7a2 []any
	)
	for __obf_6c0498eabb04b3fb, __obf_9573d75ff89816de := range __obf_2b62c53fec3ac5a3 {
		__obf_cabff7cb0e419137 = append(__obf_cabff7cb0e419137, __obf_6c0498eabb04b3fb)
		__obf_b4f45d1add3ff7a2 = append(__obf_b4f45d1add3ff7a2, __obf_9573d75ff89816de)
	}
	return __obf_036c6b778079b76b.InsertBatch(__obf_cabff7cb0e419137, __obf_b4f45d1add3ff7a2)
}

// 批量插入
func (__obf_036c6b778079b76b *Context) InsertBatch(__obf_cabff7cb0e419137 []string, __obf_2b62c53fec3ac5a3 ...[]any) (sql.Result, error) {
	var (
		__obf_b4f45d1add3ff7a2 []any
		__obf_0787ed292dc128be []string
	)
	for _, __obf_2591249cec989e72 := range __obf_2b62c53fec3ac5a3 {
		__obf_4108d94ccc3925fa := make([]string, len(__obf_2591249cec989e72))
		for __obf_2a7388295d503598, __obf_9573d75ff89816de := range __obf_2591249cec989e72 {
			__obf_4108d94ccc3925fa[__obf_2a7388295d503598] = ParamMarker
			__obf_b4f45d1add3ff7a2 = append(__obf_b4f45d1add3ff7a2, __obf_9573d75ff89816de)
		}
		__obf_0787ed292dc128be = append(__obf_0787ed292dc128be, fmt.Sprintf("(%s)", __obf_b3ceeeee12719402(__obf_4108d94ccc3925fa, SeqComma)))
	}

	__obf_3b4f9deb3c30de53 := fmt.Sprintf("insert into %s (%s) values %s", __obf_036c6b778079b76b.__obf_5bf373809aa21efc, __obf_b3ceeeee12719402(__obf_cabff7cb0e419137, SeqComma), __obf_b3ceeeee12719402(__obf_0787ed292dc128be, SeqComma))
	return __obf_036c6b778079b76b.__obf_d5f9a0ffa9af380a(__obf_3b4f9deb3c30de53, __obf_b4f45d1add3ff7a2...)
}

// 使用map更新
func (__obf_036c6b778079b76b *Context) UpdateMap(__obf_3cef69abfd4b5fed map[string]any) (__obf_179a7402a13f78d3 int64, __obf_4795d167371f16d4 error) {
	var (
		__obf_b4f45d1add3ff7a2 []any
		__obf_85d298bd6a284f1d []string
	)
	for __obf_6c0498eabb04b3fb, __obf_9573d75ff89816de := range __obf_3cef69abfd4b5fed {
		__obf_b4f45d1add3ff7a2 = append(__obf_b4f45d1add3ff7a2, __obf_9573d75ff89816de)
		__obf_85d298bd6a284f1d = append(__obf_85d298bd6a284f1d, fmt.Sprintf("%s=%s", __obf_6c0498eabb04b3fb, ParamMarker))
	}
	__obf_94276ff5367398a4 := __obf_b3ceeeee12719402(__obf_85d298bd6a284f1d, SeqComma)
	__obf_179a7402a13f78d3, __obf_4795d167371f16d4 = __obf_036c6b778079b76b.Update(__obf_94276ff5367398a4, __obf_b4f45d1add3ff7a2...)
	return
}

// 更新
func (__obf_036c6b778079b76b *Context) Update(__obf_94276ff5367398a4 string, __obf_3cef69abfd4b5fed ...any) (__obf_179a7402a13f78d3 int64, __obf_4795d167371f16d4 error) {
	__obf_a7c1276d711ccf62 := "update %s set %s %s"
	__obf_511a4b0e43c6c0f2 := __obf_507ca9459d85fdd1(__obf_036c6b778079b76b.__obf_8649156a558d0fb4, Grouping)
	__obf_3b4f9deb3c30de53 := fmt.Sprintf(__obf_a7c1276d711ccf62, __obf_036c6b778079b76b.__obf_5bf373809aa21efc, __obf_94276ff5367398a4, __obf_511a4b0e43c6c0f2)
	__obf_b4f45d1add3ff7a2 := append(__obf_3cef69abfd4b5fed, __obf_036c6b778079b76b.__obf_3cef69abfd4b5fed...)
	var __obf_7480d8040044dcbd sql.Result
	__obf_7480d8040044dcbd, __obf_4795d167371f16d4 = __obf_036c6b778079b76b.__obf_d5f9a0ffa9af380a(__obf_3b4f9deb3c30de53, __obf_b4f45d1add3ff7a2...)
	if __obf_4795d167371f16d4 != nil {
		return
	}
	__obf_179a7402a13f78d3, __obf_4795d167371f16d4 = __obf_7480d8040044dcbd.RowsAffected()
	return
}

// 删除
func (__obf_036c6b778079b76b *Context) Delete() (__obf_179a7402a13f78d3 int64, __obf_4795d167371f16d4 error) {
	__obf_a7c1276d711ccf62 := "delete from %s %s"
	__obf_511a4b0e43c6c0f2 := __obf_507ca9459d85fdd1(__obf_036c6b778079b76b.__obf_8649156a558d0fb4, Grouping)

	__obf_3b4f9deb3c30de53 := fmt.Sprintf(__obf_a7c1276d711ccf62, __obf_036c6b778079b76b.__obf_5bf373809aa21efc, __obf_511a4b0e43c6c0f2)
	var __obf_7480d8040044dcbd sql.Result
	__obf_7480d8040044dcbd, __obf_4795d167371f16d4 = __obf_036c6b778079b76b.__obf_d5f9a0ffa9af380a(__obf_3b4f9deb3c30de53, __obf_036c6b778079b76b.__obf_3cef69abfd4b5fed...)
	if __obf_4795d167371f16d4 != nil {
		return
	}
	__obf_179a7402a13f78d3, __obf_4795d167371f16d4 = __obf_7480d8040044dcbd.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_036c6b778079b76b *Context) Select(__obf_ecfa1ba7846c04d3 any, sql string, __obf_3cef69abfd4b5fed ...any) error {
	__obf_036c6b778079b76b.sql = sql
	__obf_036c6b778079b76b.__obf_3cef69abfd4b5fed = __obf_3cef69abfd4b5fed
	return __obf_036c6b778079b76b.__obf_fde655c5896dd5ee(__obf_ecfa1ba7846c04d3, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_036c6b778079b76b *Context) Get(__obf_ecfa1ba7846c04d3 any, sql string, __obf_3cef69abfd4b5fed ...any) error {
	__obf_036c6b778079b76b.sql = sql
	__obf_036c6b778079b76b.__obf_3cef69abfd4b5fed = __obf_3cef69abfd4b5fed
	return __obf_036c6b778079b76b.__obf_fde655c5896dd5ee(__obf_ecfa1ba7846c04d3, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_036c6b778079b76b *Context) Exec(sql string, __obf_3cef69abfd4b5fed ...any) (sql.Result, error) {
	return __obf_036c6b778079b76b.__obf_d5f9a0ffa9af380a(sql, __obf_3cef69abfd4b5fed...)
}

// 创建表
func (__obf_036c6b778079b76b *Context) Create(sql string) (sql.Result, error) {
	return __obf_036c6b778079b76b.__obf_d5f9a0ffa9af380a(sql)
}

// 删除表
func (__obf_036c6b778079b76b *Context) Drop() (sql.Result, error) {
	return __obf_036c6b778079b76b.__obf_d5f9a0ffa9af380a(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_036c6b778079b76b.__obf_5bf373809aa21efc))
}

/////////////////////////private methods//////////////////////

// 重置Context
func (__obf_036c6b778079b76b *Context) __obf_79a134d3318127c7() *Context {
	__obf_036c6b778079b76b.sql = ""
	__obf_036c6b778079b76b.__obf_5bf373809aa21efc = ""
	__obf_036c6b778079b76b.__obf_375b71252e6f26c6 = []string{}
	__obf_036c6b778079b76b.__obf_8649156a558d0fb4 = []string{}
	__obf_036c6b778079b76b.__obf_72aa9c9e212e03fe = ""
	__obf_036c6b778079b76b.__obf_79a5b98b6cf49502 = ""
	__obf_036c6b778079b76b.__obf_3901bcf69ddf5fea = ""
	__obf_036c6b778079b76b.__obf_159c787549b63f29 = 0
	__obf_036c6b778079b76b.__obf_b9906ca51574c61e = 0
	__obf_036c6b778079b76b.__obf_3cef69abfd4b5fed = []any{}
	__obf_036c6b778079b76b.__obf_4e1da10c2afe56d6 = nil
	__obf_036c6b778079b76b.__obf_dacd8adf705f369a = false
	__obf_036c6b778079b76b.__obf_04522ec63d23e18d = false
	return __obf_036c6b778079b76b
}

// 查询方法
func (__obf_036c6b778079b76b *Context) __obf_fde655c5896dd5ee(__obf_ecfa1ba7846c04d3 any, __obf_547a935d85a68218 int) (__obf_4795d167371f16d4 error) {
	defer __obf_036c6b778079b76b.__obf_ed3a9a8432b2d457.__obf_d6fbd616396c8e75.Put(__obf_036c6b778079b76b)
	__obf_1f4eede373123ce9, __obf_b2f227a0dc11c0d8 := context.WithTimeout(context.Background(), __obf_036c6b778079b76b.__obf_ed3a9a8432b2d457.__obf_435469c73588f1a2)
	defer __obf_b2f227a0dc11c0d8()
	if __obf_036c6b778079b76b.sql == "" {
		__obf_036c6b778079b76b.sql = __obf_036c6b778079b76b.__obf_13932d5d7fde5bcd(__obf_ecfa1ba7846c04d3)
	}
	switch __obf_547a935d85a68218 {
	case SelectTypeOne:
		if __obf_036c6b778079b76b.__obf_4e1da10c2afe56d6 != nil {
			__obf_4795d167371f16d4 = __obf_036c6b778079b76b.__obf_4e1da10c2afe56d6.GetContext(__obf_1f4eede373123ce9, __obf_ecfa1ba7846c04d3, __obf_036c6b778079b76b.sql, __obf_036c6b778079b76b.__obf_3cef69abfd4b5fed...)
		} else {
			__obf_4795d167371f16d4 = __obf_036c6b778079b76b.__obf_ed3a9a8432b2d457.GetContext(__obf_1f4eede373123ce9, __obf_ecfa1ba7846c04d3, __obf_036c6b778079b76b.sql, __obf_036c6b778079b76b.__obf_3cef69abfd4b5fed...)
		}
	case SelectTypeMany:
		if __obf_036c6b778079b76b.__obf_4e1da10c2afe56d6 != nil {
			__obf_4795d167371f16d4 = __obf_036c6b778079b76b.__obf_4e1da10c2afe56d6.SelectContext(__obf_1f4eede373123ce9, __obf_ecfa1ba7846c04d3, __obf_036c6b778079b76b.sql, __obf_036c6b778079b76b.__obf_3cef69abfd4b5fed...)
		} else {
			__obf_4795d167371f16d4 = __obf_036c6b778079b76b.__obf_ed3a9a8432b2d457.SelectContext(__obf_1f4eede373123ce9, __obf_ecfa1ba7846c04d3, __obf_036c6b778079b76b.sql, __obf_036c6b778079b76b.__obf_3cef69abfd4b5fed...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_036c6b778079b76b *Context) __obf_d5f9a0ffa9af380a(__obf_3b4f9deb3c30de53 string, __obf_3cef69abfd4b5fed ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed)
	defer __obf_036c6b778079b76b.__obf_ed3a9a8432b2d457.__obf_d6fbd616396c8e75.Put(__obf_036c6b778079b76b)
	__obf_1f4eede373123ce9, __obf_b2f227a0dc11c0d8 := context.WithTimeout(context.Background(), __obf_036c6b778079b76b.__obf_ed3a9a8432b2d457.__obf_435469c73588f1a2)
	defer __obf_b2f227a0dc11c0d8()

	var __obf_eeff625af9405efe sqlx.ExecerContext
	if __obf_036c6b778079b76b.__obf_4e1da10c2afe56d6 != nil {
		__obf_eeff625af9405efe = __obf_036c6b778079b76b.__obf_4e1da10c2afe56d6
	} else {
		__obf_eeff625af9405efe = __obf_036c6b778079b76b.__obf_ed3a9a8432b2d457
	}
	return __obf_eeff625af9405efe.ExecContext(__obf_1f4eede373123ce9, __obf_3b4f9deb3c30de53, __obf_3cef69abfd4b5fed...)
}

// select查询语句的拼接
func (__obf_036c6b778079b76b *Context) __obf_13932d5d7fde5bcd(__obf_ecfa1ba7846c04d3 any) string {
	var __obf_ad9413ed6cce9d39 []string
	__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, "select")
	if len(__obf_036c6b778079b76b.__obf_375b71252e6f26c6) != 0 {
		__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, __obf_b3ceeeee12719402(__obf_036c6b778079b76b.__obf_375b71252e6f26c6, SeqComma))
	} else {
		// 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
		// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
		__obf_72b4fbe76d2b2179 := __obf_1c4ef1753f0a87ce(__obf_ecfa1ba7846c04d3)
		if len(__obf_72b4fbe76d2b2179) > 0 {
			__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, __obf_b3ceeeee12719402(__obf_72b4fbe76d2b2179, SeqComma))
		} else {
			__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, "*")
		}
	}
	__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, "from "+__obf_036c6b778079b76b.__obf_5bf373809aa21efc)
	if len(__obf_036c6b778079b76b.__obf_8649156a558d0fb4) != 0 {
		__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, __obf_507ca9459d85fdd1(__obf_036c6b778079b76b.__obf_8649156a558d0fb4, Grouping))
	}

	if __obf_036c6b778079b76b.__obf_79a5b98b6cf49502 != "" {
		__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, "group by "+__obf_036c6b778079b76b.__obf_79a5b98b6cf49502)
	}

	if __obf_036c6b778079b76b.__obf_3901bcf69ddf5fea != "" {
		__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, "having "+__obf_036c6b778079b76b.__obf_3901bcf69ddf5fea)
	}

	if __obf_036c6b778079b76b.__obf_72aa9c9e212e03fe != "" {
		__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, "order by "+__obf_036c6b778079b76b.__obf_72aa9c9e212e03fe)
	}

	if __obf_036c6b778079b76b.__obf_159c787549b63f29 != 0 {
		__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, fmt.Sprintf("limit %d, %d", __obf_036c6b778079b76b.__obf_b9906ca51574c61e, __obf_036c6b778079b76b.__obf_159c787549b63f29))
	}
	if __obf_036c6b778079b76b.__obf_dacd8adf705f369a {
		__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, "lock in share mode")
	}
	if __obf_036c6b778079b76b.__obf_04522ec63d23e18d {
		__obf_ad9413ed6cce9d39 = append(__obf_ad9413ed6cce9d39, "for update")
	}
	sql := __obf_b3ceeeee12719402(__obf_ad9413ed6cce9d39, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_036c6b778079b76b.__obf_3cef69abfd4b5fed)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_507ca9459d85fdd1(__obf_8649156a558d0fb4 []string, __obf_77b849016a88d61a string) string {
	if len(__obf_8649156a558d0fb4) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_8649156a558d0fb4, __obf_77b849016a88d61a))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_b3ceeeee12719402(__obf_3cef69abfd4b5fed []string, __obf_cfa7c05ebdd73e4f string) string {
	return strings.Join(__obf_3cef69abfd4b5fed, __obf_cfa7c05ebdd73e4f)
}

// 解析对象中的 `db tag`
// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
func __obf_1c4ef1753f0a87ce(__obf_ecfa1ba7846c04d3 any) (__obf_cabff7cb0e419137 []string) {
	__obf_fe4dab033b4bbf22 := reflect.ValueOf(__obf_ecfa1ba7846c04d3)
	__obf_98a4c1858d15f6fd := __obf_fe4dab033b4bbf22.Type().Elem()
	var __obf_95c95ab62fed5eaa reflect.Type
	if __obf_98a4c1858d15f6fd.Kind() == reflect.Slice {
		__obf_95c95ab62fed5eaa = __obf_98a4c1858d15f6fd.Elem()
	} else {
		__obf_95c95ab62fed5eaa = __obf_98a4c1858d15f6fd
	}
	for __obf_2a7388295d503598 := 0; __obf_2a7388295d503598 < __obf_95c95ab62fed5eaa.NumField(); __obf_2a7388295d503598++ {
		__obf_52fff51e27b377b0 := __obf_95c95ab62fed5eaa.Field(__obf_2a7388295d503598).Tag.Get(DBTag)
		if __obf_52fff51e27b377b0 != "" {
			__obf_cabff7cb0e419137 = append(__obf_cabff7cb0e419137, __obf_52fff51e27b377b0)
		}
	}
	return
}

func CleanSQLXMap(__obf_7644650c72213f91 map[string]any) map[string]any {
	for __obf_6c0498eabb04b3fb, __obf_9573d75ff89816de := range __obf_7644650c72213f91 {
		switch __obf_a0400a0a546202a1 := __obf_9573d75ff89816de.(type) {
		case []byte:
			__obf_2eac47834e11c59d := string(__obf_a0400a0a546202a1)

			// 尝试智能识别具体类型
			__obf_d8a74ef9a4d6ebf6 := strings.TrimSpace(__obf_2eac47834e11c59d)

			// bool
			if strings.EqualFold(__obf_d8a74ef9a4d6ebf6, "true") || strings.EqualFold(__obf_d8a74ef9a4d6ebf6, "false") {
				if __obf_9e1db65f077c8102, __obf_4795d167371f16d4 := strconv.ParseBool(__obf_d8a74ef9a4d6ebf6); __obf_4795d167371f16d4 == nil {
					__obf_7644650c72213f91[__obf_6c0498eabb04b3fb] = __obf_9e1db65f077c8102
					continue
				}
			}

			// int
			if __obf_2a7388295d503598, __obf_4795d167371f16d4 := strconv.Atoi(__obf_d8a74ef9a4d6ebf6); __obf_4795d167371f16d4 == nil {
				__obf_7644650c72213f91[__obf_6c0498eabb04b3fb] = __obf_2a7388295d503598
				continue
			}

			// float
			if __obf_2e7920ea71b44a3b, __obf_4795d167371f16d4 := strconv.ParseFloat(__obf_d8a74ef9a4d6ebf6, 64); __obf_4795d167371f16d4 == nil {
				__obf_7644650c72213f91[__obf_6c0498eabb04b3fb] = __obf_2e7920ea71b44a3b
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_d8a74ef9a4d6ebf6, "{") && strings.HasSuffix(__obf_d8a74ef9a4d6ebf6, "}")) ||
				(strings.HasPrefix(__obf_d8a74ef9a4d6ebf6, "[") && strings.HasSuffix(__obf_d8a74ef9a4d6ebf6, "]")) {
				var __obf_2199dbb4ae4933fb any
				if __obf_4795d167371f16d4 := json.Unmarshal(__obf_a0400a0a546202a1, &__obf_2199dbb4ae4933fb); __obf_4795d167371f16d4 == nil {
					__obf_7644650c72213f91[__obf_6c0498eabb04b3fb] = __obf_2199dbb4ae4933fb
					continue
				}
			}

			// 默认保留为 string
			__obf_7644650c72213f91[__obf_6c0498eabb04b3fb] = __obf_2eac47834e11c59d
		default:
			__obf_7644650c72213f91[__obf_6c0498eabb04b3fb] = __obf_a0400a0a546202a1
		}
	}
	return __obf_7644650c72213f91
}

func ToListMap(__obf_4f052f7555da44c1 *sqlx.Rows) []map[string]any {
	var __obf_7480d8040044dcbd []map[string]any
	for __obf_4f052f7555da44c1.Next() {
		__obf_9c2998b49b552a7b := make(map[string]any)
		if __obf_4795d167371f16d4 := __obf_4f052f7555da44c1.MapScan(__obf_9c2998b49b552a7b); __obf_4795d167371f16d4 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_4795d167371f16d4.Error()))
			continue
		}

		__obf_7480d8040044dcbd = append(__obf_7480d8040044dcbd, CleanSQLXMap(__obf_9c2998b49b552a7b))
	}
	return __obf_7480d8040044dcbd
}
