package __obf_52bf13aa40478808

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
type FuncTx func(__obf_f1097d4e39bff719 *sqlx.Tx, __obf_cd0ddd46582953ac any) error

func Open(__obf_ae2c7f6c832213f8, __obf_00f6397308d34896 string, __obf_ac5442efe07d6f83 time.Duration) (*DB, error) {
	__obf_86bb12f77b537935, __obf_e6c24855347eb3d0 := sqlx.Open(__obf_ae2c7f6c832213f8, __obf_00f6397308d34896)
	if __obf_e6c24855347eb3d0 != nil {
		return nil, __obf_e6c24855347eb3d0
	}
	__obf_6ab9f0211d0f4779 := &DB{
		DB:                     __obf_86bb12f77b537935,
		__obf_ac5442efe07d6f83: __obf_ac5442efe07d6f83,
	}
	__obf_6ab9f0211d0f4779.__obf_a258c25df16a54e4.New = func() any {
		return __obf_6ab9f0211d0f4779.__obf_e3488b2a1a9d2bef()
	}
	return __obf_6ab9f0211d0f4779, nil
}

type DB struct {
	*sqlx.DB
	__obf_ac5442efe07d6f83 time.Duration
	__obf_a258c25df16a54e4 sync.Pool
}

func (__obf_86bb12f77b537935 *DB) __obf_e3488b2a1a9d2bef() *Context {
	return &Context{__obf_86bb12f77b537935: __obf_86bb12f77b537935}
}

// 获取一个`SQL`执行`Context`
func (__obf_86bb12f77b537935 *DB) Acquire() *Context {
	// 无需加锁，sync.Pool本身是线程安全的
	__obf_282740e215b7ba25 := __obf_86bb12f77b537935.__obf_a258c25df16a54e4.Get().(*Context)
	__obf_282740e215b7ba25.__obf_80ee30d1f66185a3()
	return __obf_282740e215b7ba25
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_86bb12f77b537935 *DB) AcquireTx(__obf_f1097d4e39bff719 *sqlx.Tx) *Context {
	__obf_282740e215b7ba25 := __obf_86bb12f77b537935.Acquire()
	__obf_282740e215b7ba25.__obf_f1097d4e39bff719 = __obf_f1097d4e39bff719
	return __obf_282740e215b7ba25
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_86bb12f77b537935 *DB) WithTx(__obf_9ce67657fbd9fcc9 FuncTx, __obf_cd0ddd46582953ac any) (__obf_e6c24855347eb3d0 error) {
	var __obf_f1097d4e39bff719 *sqlx.Tx
	__obf_f1097d4e39bff719, __obf_e6c24855347eb3d0 = __obf_86bb12f77b537935.Beginx()
	if __obf_e6c24855347eb3d0 != nil {
		return
	}
	defer func() {
		if __obf_e6c24855347eb3d0 != nil && __obf_f1097d4e39bff719 != nil {
			__obf_e6c24855347eb3d0 = __obf_f1097d4e39bff719.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_e6c24855347eb3d0 = __obf_9ce67657fbd9fcc9(__obf_f1097d4e39bff719, __obf_cd0ddd46582953ac); __obf_e6c24855347eb3d0 != nil {
		return
	}

	__obf_e6c24855347eb3d0 = __obf_f1097d4e39bff719.Commit()
	return
}

type Context struct {
	__obf_86bb12f77b537935 *DB
	__obf_f1097d4e39bff719 *sqlx.Tx //事务
	sql                    string
	__obf_1e594c2872df3ee6 string
	__obf_12ef4265b6c5952c []string
	__obf_fd58175a17e6db6a []string
	__obf_198495c974926bd5 string
	__obf_7b02799558bf0c2b string
	__obf_a126235a6d03d790 string
	__obf_cf8c3f8498fcaffa int64
	__obf_0837d1022b120201 int64
	__obf_cd0ddd46582953ac []any
	__obf_ff2346d785224919 bool //排他锁
	__obf_61706855cb19b018 bool //共享锁
}

func (__obf_282740e215b7ba25 *Context) Name(__obf_1e594c2872df3ee6 string) *Context {
	__obf_282740e215b7ba25.__obf_1e594c2872df3ee6 = __obf_1e594c2872df3ee6
	return __obf_282740e215b7ba25
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_282740e215b7ba25 *Context) What(__obf_12ef4265b6c5952c []string) *Context {
	__obf_282740e215b7ba25.__obf_12ef4265b6c5952c = __obf_12ef4265b6c5952c
	return __obf_282740e215b7ba25
}

func (__obf_282740e215b7ba25 *Context) Where(__obf_8062258fcb6f5cc5 string, __obf_cd0ddd46582953ac ...any) *Context {
	__obf_282740e215b7ba25.__obf_fd58175a17e6db6a = append(__obf_282740e215b7ba25.__obf_fd58175a17e6db6a, __obf_8062258fcb6f5cc5)
	__obf_282740e215b7ba25.__obf_cd0ddd46582953ac = append(__obf_282740e215b7ba25.__obf_cd0ddd46582953ac, __obf_cd0ddd46582953ac...)
	return __obf_282740e215b7ba25
}

// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
func (__obf_282740e215b7ba25 *Context) WhereIn(__obf_9ad5b2ee11618c02 string, __obf_cd0ddd46582953ac []any) *Context {
	__obf_d3e41717719fe753 := len(__obf_cd0ddd46582953ac)
	__obf_8b8cd4af3afe4a2e := make([]string, __obf_d3e41717719fe753)
	for __obf_935763ab96242d2f := 0; __obf_935763ab96242d2f < __obf_d3e41717719fe753; __obf_935763ab96242d2f++ {
		__obf_8b8cd4af3afe4a2e[__obf_935763ab96242d2f] = ParamMarker
	}
	__obf_d34fdc133fc71acf := fmt.Sprintf("%s in (%s)", __obf_9ad5b2ee11618c02, __obf_e0808ea7302c82cd(__obf_8b8cd4af3afe4a2e, SeqComma))
	return __obf_282740e215b7ba25.Where(__obf_d34fdc133fc71acf, __obf_cd0ddd46582953ac...)
}

func (__obf_282740e215b7ba25 *Context) Order(__obf_198495c974926bd5 string) *Context {
	__obf_282740e215b7ba25.__obf_198495c974926bd5 = __obf_198495c974926bd5
	return __obf_282740e215b7ba25
}

func (__obf_282740e215b7ba25 *Context) Limit(__obf_cf8c3f8498fcaffa int64) *Context {
	__obf_282740e215b7ba25.__obf_cf8c3f8498fcaffa = __obf_cf8c3f8498fcaffa
	return __obf_282740e215b7ba25
}

func (__obf_282740e215b7ba25 *Context) Offset(__obf_0837d1022b120201 int64) *Context {
	__obf_282740e215b7ba25.__obf_0837d1022b120201 = __obf_0837d1022b120201
	return __obf_282740e215b7ba25
}

func (__obf_282740e215b7ba25 *Context) Group(__obf_7b02799558bf0c2b string) *Context {
	__obf_282740e215b7ba25.__obf_7b02799558bf0c2b = __obf_7b02799558bf0c2b
	return __obf_282740e215b7ba25
}

func (__obf_282740e215b7ba25 *Context) Having(__obf_a126235a6d03d790 string, __obf_cd0ddd46582953ac ...any) *Context {
	__obf_282740e215b7ba25.__obf_a126235a6d03d790 = __obf_a126235a6d03d790
	__obf_282740e215b7ba25.__obf_cd0ddd46582953ac = append(__obf_282740e215b7ba25.__obf_cd0ddd46582953ac, __obf_cd0ddd46582953ac...)
	return __obf_282740e215b7ba25
}

// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
func (__obf_282740e215b7ba25 *Context) LockX() *Context {
	__obf_282740e215b7ba25.__obf_ff2346d785224919 = true
	return __obf_282740e215b7ba25
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_282740e215b7ba25 *Context) LockS() *Context {
	__obf_282740e215b7ba25.__obf_61706855cb19b018 = true
	return __obf_282740e215b7ba25
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_282740e215b7ba25 *Context) FindMany(__obf_b2810da3d9419758 any) error {
	return __obf_282740e215b7ba25.__obf_0aac3242027c1fa4(__obf_b2810da3d9419758, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_282740e215b7ba25 *Context) FindById(__obf_b2810da3d9419758 any) error {
	return __obf_282740e215b7ba25.__obf_0aac3242027c1fa4(__obf_b2810da3d9419758, SelectTypeOne)
}

// 插入
func (__obf_282740e215b7ba25 *Context) Insert(__obf_125c7c95577ed4b5 map[string]any) (sql.Result, error) {
	var (
		__obf_9dfd890a99ebcac9 []string
		__obf_beb42c5ddfaa3445 []any
	)
	for __obf_13037cdadbf8d35e, __obf_306dddc11fc3d690 := range __obf_125c7c95577ed4b5 {
		__obf_9dfd890a99ebcac9 = append(__obf_9dfd890a99ebcac9, __obf_13037cdadbf8d35e)
		__obf_beb42c5ddfaa3445 = append(__obf_beb42c5ddfaa3445, __obf_306dddc11fc3d690)
	}
	return __obf_282740e215b7ba25.InsertBatch(__obf_9dfd890a99ebcac9, __obf_beb42c5ddfaa3445)
}

// 批量插入
func (__obf_282740e215b7ba25 *Context) InsertBatch(__obf_9dfd890a99ebcac9 []string, __obf_125c7c95577ed4b5 ...[]any) (sql.Result, error) {
	var (
		__obf_beb42c5ddfaa3445 []any
		__obf_9808300ab0b1184a []string
	)
	for _, __obf_d1a1f5aa8ecd1af9 := range __obf_125c7c95577ed4b5 {
		__obf_8b8cd4af3afe4a2e := make([]string, len(__obf_d1a1f5aa8ecd1af9))
		for __obf_935763ab96242d2f, __obf_306dddc11fc3d690 := range __obf_d1a1f5aa8ecd1af9 {
			__obf_8b8cd4af3afe4a2e[__obf_935763ab96242d2f] = ParamMarker
			__obf_beb42c5ddfaa3445 = append(__obf_beb42c5ddfaa3445, __obf_306dddc11fc3d690)
		}
		__obf_9808300ab0b1184a = append(__obf_9808300ab0b1184a, fmt.Sprintf("(%s)", __obf_e0808ea7302c82cd(__obf_8b8cd4af3afe4a2e, SeqComma)))
	}

	__obf_a856c5ef2c6a7852 := fmt.Sprintf("insert into %s (%s) values %s", __obf_282740e215b7ba25.__obf_1e594c2872df3ee6, __obf_e0808ea7302c82cd(__obf_9dfd890a99ebcac9, SeqComma), __obf_e0808ea7302c82cd(__obf_9808300ab0b1184a, SeqComma))
	return __obf_282740e215b7ba25.__obf_81a79baaaf1b3f4c(__obf_a856c5ef2c6a7852, __obf_beb42c5ddfaa3445...)
}

// 使用map更新
func (__obf_282740e215b7ba25 *Context) UpdateMap(__obf_cd0ddd46582953ac map[string]any) (__obf_1fa25eb035b90351 int64, __obf_e6c24855347eb3d0 error) {
	var (
		__obf_beb42c5ddfaa3445 []any
		__obf_0db88687d5af63dc []string
	)
	for __obf_13037cdadbf8d35e, __obf_306dddc11fc3d690 := range __obf_cd0ddd46582953ac {
		__obf_beb42c5ddfaa3445 = append(__obf_beb42c5ddfaa3445, __obf_306dddc11fc3d690)
		__obf_0db88687d5af63dc = append(__obf_0db88687d5af63dc, fmt.Sprintf("%s=%s", __obf_13037cdadbf8d35e, ParamMarker))
	}
	__obf_35f96d8bcea24e7c := __obf_e0808ea7302c82cd(__obf_0db88687d5af63dc, SeqComma)
	__obf_1fa25eb035b90351, __obf_e6c24855347eb3d0 = __obf_282740e215b7ba25.Update(__obf_35f96d8bcea24e7c, __obf_beb42c5ddfaa3445...)
	return
}

// 更新
func (__obf_282740e215b7ba25 *Context) Update(__obf_35f96d8bcea24e7c string, __obf_cd0ddd46582953ac ...any) (__obf_1fa25eb035b90351 int64, __obf_e6c24855347eb3d0 error) {
	__obf_adfa3ae498d18f25 := "update %s set %s %s"
	__obf_8062258fcb6f5cc5 := __obf_a8bcd0981395934b(__obf_282740e215b7ba25.__obf_fd58175a17e6db6a, Grouping)
	__obf_a856c5ef2c6a7852 := fmt.Sprintf(__obf_adfa3ae498d18f25, __obf_282740e215b7ba25.__obf_1e594c2872df3ee6, __obf_35f96d8bcea24e7c, __obf_8062258fcb6f5cc5)
	__obf_beb42c5ddfaa3445 := append(__obf_cd0ddd46582953ac, __obf_282740e215b7ba25.__obf_cd0ddd46582953ac...)
	var __obf_c560cfbfddbf86d2 sql.Result
	__obf_c560cfbfddbf86d2, __obf_e6c24855347eb3d0 = __obf_282740e215b7ba25.__obf_81a79baaaf1b3f4c(__obf_a856c5ef2c6a7852, __obf_beb42c5ddfaa3445...)
	if __obf_e6c24855347eb3d0 != nil {
		return
	}
	__obf_1fa25eb035b90351, __obf_e6c24855347eb3d0 = __obf_c560cfbfddbf86d2.RowsAffected()
	return
}

// 删除
func (__obf_282740e215b7ba25 *Context) Delete() (__obf_1fa25eb035b90351 int64, __obf_e6c24855347eb3d0 error) {
	__obf_adfa3ae498d18f25 := "delete from %s %s"
	__obf_8062258fcb6f5cc5 := __obf_a8bcd0981395934b(__obf_282740e215b7ba25.__obf_fd58175a17e6db6a, Grouping)

	__obf_a856c5ef2c6a7852 := fmt.Sprintf(__obf_adfa3ae498d18f25, __obf_282740e215b7ba25.__obf_1e594c2872df3ee6, __obf_8062258fcb6f5cc5)
	var __obf_c560cfbfddbf86d2 sql.Result
	__obf_c560cfbfddbf86d2, __obf_e6c24855347eb3d0 = __obf_282740e215b7ba25.__obf_81a79baaaf1b3f4c(__obf_a856c5ef2c6a7852, __obf_282740e215b7ba25.__obf_cd0ddd46582953ac...)
	if __obf_e6c24855347eb3d0 != nil {
		return
	}
	__obf_1fa25eb035b90351, __obf_e6c24855347eb3d0 = __obf_c560cfbfddbf86d2.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_282740e215b7ba25 *Context) Select(__obf_b2810da3d9419758 any, sql string, __obf_cd0ddd46582953ac ...any) error {
	__obf_282740e215b7ba25.sql = sql
	__obf_282740e215b7ba25.__obf_cd0ddd46582953ac = __obf_cd0ddd46582953ac
	return __obf_282740e215b7ba25.__obf_0aac3242027c1fa4(__obf_b2810da3d9419758, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_282740e215b7ba25 *Context) Get(__obf_b2810da3d9419758 any, sql string, __obf_cd0ddd46582953ac ...any) error {
	__obf_282740e215b7ba25.sql = sql
	__obf_282740e215b7ba25.__obf_cd0ddd46582953ac = __obf_cd0ddd46582953ac
	return __obf_282740e215b7ba25.__obf_0aac3242027c1fa4(__obf_b2810da3d9419758, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_282740e215b7ba25 *Context) Exec(sql string, __obf_cd0ddd46582953ac ...any) (sql.Result, error) {
	return __obf_282740e215b7ba25.__obf_81a79baaaf1b3f4c(sql, __obf_cd0ddd46582953ac...)
}

// 创建表
func (__obf_282740e215b7ba25 *Context) Create(sql string) (sql.Result, error) {
	return __obf_282740e215b7ba25.__obf_81a79baaaf1b3f4c(sql)
}

// 删除表
func (__obf_282740e215b7ba25 *Context) Drop() (sql.Result, error) {
	return __obf_282740e215b7ba25.__obf_81a79baaaf1b3f4c(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_282740e215b7ba25.__obf_1e594c2872df3ee6))
}

/////////////////////////private methods//////////////////////

// 重置Context
func (__obf_282740e215b7ba25 *Context) __obf_80ee30d1f66185a3() *Context {
	__obf_282740e215b7ba25.sql = ""
	__obf_282740e215b7ba25.__obf_1e594c2872df3ee6 = ""
	__obf_282740e215b7ba25.__obf_12ef4265b6c5952c = []string{}
	__obf_282740e215b7ba25.__obf_fd58175a17e6db6a = []string{}
	__obf_282740e215b7ba25.__obf_198495c974926bd5 = ""
	__obf_282740e215b7ba25.__obf_7b02799558bf0c2b = ""
	__obf_282740e215b7ba25.__obf_a126235a6d03d790 = ""
	__obf_282740e215b7ba25.__obf_cf8c3f8498fcaffa = 0
	__obf_282740e215b7ba25.__obf_0837d1022b120201 = 0
	__obf_282740e215b7ba25.__obf_cd0ddd46582953ac = []any{}
	__obf_282740e215b7ba25.__obf_f1097d4e39bff719 = nil
	__obf_282740e215b7ba25.__obf_61706855cb19b018 = false
	__obf_282740e215b7ba25.__obf_ff2346d785224919 = false
	return __obf_282740e215b7ba25
}

// 查询方法
func (__obf_282740e215b7ba25 *Context) __obf_0aac3242027c1fa4(__obf_b2810da3d9419758 any, __obf_b61318ea43371396 int) (__obf_e6c24855347eb3d0 error) {
	defer __obf_282740e215b7ba25.__obf_86bb12f77b537935.__obf_a258c25df16a54e4.Put(__obf_282740e215b7ba25)
	__obf_04a526d9e5c7d2f5, __obf_a9be0f1a8cdef675 := context.WithTimeout(context.Background(), __obf_282740e215b7ba25.__obf_86bb12f77b537935.__obf_ac5442efe07d6f83)
	defer __obf_a9be0f1a8cdef675()
	if __obf_282740e215b7ba25.sql == "" {
		__obf_282740e215b7ba25.sql = __obf_282740e215b7ba25.__obf_197d72fb0c7a43d0(__obf_b2810da3d9419758)
	}
	switch __obf_b61318ea43371396 {
	case SelectTypeOne:
		if __obf_282740e215b7ba25.__obf_f1097d4e39bff719 != nil {
			__obf_e6c24855347eb3d0 = __obf_282740e215b7ba25.__obf_f1097d4e39bff719.GetContext(__obf_04a526d9e5c7d2f5, __obf_b2810da3d9419758, __obf_282740e215b7ba25.sql, __obf_282740e215b7ba25.__obf_cd0ddd46582953ac...)
		} else {
			__obf_e6c24855347eb3d0 = __obf_282740e215b7ba25.__obf_86bb12f77b537935.GetContext(__obf_04a526d9e5c7d2f5, __obf_b2810da3d9419758, __obf_282740e215b7ba25.sql, __obf_282740e215b7ba25.__obf_cd0ddd46582953ac...)
		}
	case SelectTypeMany:
		if __obf_282740e215b7ba25.__obf_f1097d4e39bff719 != nil {
			__obf_e6c24855347eb3d0 = __obf_282740e215b7ba25.__obf_f1097d4e39bff719.SelectContext(__obf_04a526d9e5c7d2f5, __obf_b2810da3d9419758, __obf_282740e215b7ba25.sql, __obf_282740e215b7ba25.__obf_cd0ddd46582953ac...)
		} else {
			__obf_e6c24855347eb3d0 = __obf_282740e215b7ba25.__obf_86bb12f77b537935.SelectContext(__obf_04a526d9e5c7d2f5, __obf_b2810da3d9419758, __obf_282740e215b7ba25.sql, __obf_282740e215b7ba25.__obf_cd0ddd46582953ac...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_282740e215b7ba25 *Context) __obf_81a79baaaf1b3f4c(__obf_a856c5ef2c6a7852 string, __obf_cd0ddd46582953ac ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac)
	defer __obf_282740e215b7ba25.__obf_86bb12f77b537935.__obf_a258c25df16a54e4.Put(__obf_282740e215b7ba25)
	__obf_04a526d9e5c7d2f5, __obf_a9be0f1a8cdef675 := context.WithTimeout(context.Background(), __obf_282740e215b7ba25.__obf_86bb12f77b537935.__obf_ac5442efe07d6f83)
	defer __obf_a9be0f1a8cdef675()

	var __obf_c2496db385a79b0d sqlx.ExecerContext
	if __obf_282740e215b7ba25.__obf_f1097d4e39bff719 != nil {
		__obf_c2496db385a79b0d = __obf_282740e215b7ba25.__obf_f1097d4e39bff719
	} else {
		__obf_c2496db385a79b0d = __obf_282740e215b7ba25.__obf_86bb12f77b537935
	}
	return __obf_c2496db385a79b0d.ExecContext(__obf_04a526d9e5c7d2f5, __obf_a856c5ef2c6a7852, __obf_cd0ddd46582953ac...)
}

// select查询语句的拼接
func (__obf_282740e215b7ba25 *Context) __obf_197d72fb0c7a43d0(__obf_b2810da3d9419758 any) string {
	var __obf_6318cebc5c46f8b3 []string
	__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, "select")
	if len(__obf_282740e215b7ba25.__obf_12ef4265b6c5952c) != 0 {
		__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, __obf_e0808ea7302c82cd(__obf_282740e215b7ba25.__obf_12ef4265b6c5952c, SeqComma))
	} else {
		// 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
		// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
		__obf_42777e8726c0ffe1 := __obf_0abe083e8c2cb9a0(__obf_b2810da3d9419758)
		if len(__obf_42777e8726c0ffe1) > 0 {
			__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, __obf_e0808ea7302c82cd(__obf_42777e8726c0ffe1, SeqComma))
		} else {
			__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, "*")
		}
	}
	__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, "from "+__obf_282740e215b7ba25.__obf_1e594c2872df3ee6)
	if len(__obf_282740e215b7ba25.__obf_fd58175a17e6db6a) != 0 {
		__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, __obf_a8bcd0981395934b(__obf_282740e215b7ba25.__obf_fd58175a17e6db6a, Grouping))
	}

	if __obf_282740e215b7ba25.__obf_7b02799558bf0c2b != "" {
		__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, "group by "+__obf_282740e215b7ba25.__obf_7b02799558bf0c2b)
	}

	if __obf_282740e215b7ba25.__obf_a126235a6d03d790 != "" {
		__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, "having "+__obf_282740e215b7ba25.__obf_a126235a6d03d790)
	}

	if __obf_282740e215b7ba25.__obf_198495c974926bd5 != "" {
		__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, "order by "+__obf_282740e215b7ba25.__obf_198495c974926bd5)
	}

	if __obf_282740e215b7ba25.__obf_cf8c3f8498fcaffa != 0 {
		__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, fmt.Sprintf("limit %d, %d", __obf_282740e215b7ba25.__obf_0837d1022b120201, __obf_282740e215b7ba25.__obf_cf8c3f8498fcaffa))
	}
	if __obf_282740e215b7ba25.__obf_61706855cb19b018 {
		__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, "lock in share mode")
	}
	if __obf_282740e215b7ba25.__obf_ff2346d785224919 {
		__obf_6318cebc5c46f8b3 = append(__obf_6318cebc5c46f8b3, "for update")
	}
	sql := __obf_e0808ea7302c82cd(__obf_6318cebc5c46f8b3, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_282740e215b7ba25.__obf_cd0ddd46582953ac)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_a8bcd0981395934b(__obf_fd58175a17e6db6a []string, __obf_f96baa606208d730 string) string {
	if len(__obf_fd58175a17e6db6a) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_fd58175a17e6db6a, __obf_f96baa606208d730))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_e0808ea7302c82cd(__obf_cd0ddd46582953ac []string, __obf_dd620cc3a08f7527 string) string {
	return strings.Join(__obf_cd0ddd46582953ac, __obf_dd620cc3a08f7527)
}

// 解析对象中的 `db tag`
// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
func __obf_0abe083e8c2cb9a0(__obf_b2810da3d9419758 any) (__obf_9dfd890a99ebcac9 []string) {
	__obf_bcc644fb72af6d31 := reflect.ValueOf(__obf_b2810da3d9419758)
	__obf_23ab6376879fa703 := __obf_bcc644fb72af6d31.Type().Elem()
	var __obf_bc14126bca80a890 reflect.Type
	if __obf_23ab6376879fa703.Kind() == reflect.Slice {
		__obf_bc14126bca80a890 = __obf_23ab6376879fa703.Elem()
	} else {
		__obf_bc14126bca80a890 = __obf_23ab6376879fa703
	}
	for __obf_935763ab96242d2f := 0; __obf_935763ab96242d2f < __obf_bc14126bca80a890.NumField(); __obf_935763ab96242d2f++ {
		__obf_97b2563ef492c97c := __obf_bc14126bca80a890.Field(__obf_935763ab96242d2f).Tag.Get(DBTag)
		if __obf_97b2563ef492c97c != "" {
			__obf_9dfd890a99ebcac9 = append(__obf_9dfd890a99ebcac9, __obf_97b2563ef492c97c)
		}
	}
	return
}

func CleanSQLXMap(__obf_99782ba885b4f4ff map[string]any) map[string]any {
	for __obf_13037cdadbf8d35e, __obf_306dddc11fc3d690 := range __obf_99782ba885b4f4ff {
		switch __obf_1115ae3c80598454 := __obf_306dddc11fc3d690.(type) {
		case []byte:
			__obf_37836fc6d365a20e := string(__obf_1115ae3c80598454)

			// 尝试智能识别具体类型
			__obf_008ae55685cb56f6 := strings.TrimSpace(__obf_37836fc6d365a20e)

			// bool
			if strings.EqualFold(__obf_008ae55685cb56f6, "true") || strings.EqualFold(__obf_008ae55685cb56f6, "false") {
				if __obf_da19938515b58d5b, __obf_e6c24855347eb3d0 := strconv.ParseBool(__obf_008ae55685cb56f6); __obf_e6c24855347eb3d0 == nil {
					__obf_99782ba885b4f4ff[__obf_13037cdadbf8d35e] = __obf_da19938515b58d5b
					continue
				}
			}

			// int
			if __obf_935763ab96242d2f, __obf_e6c24855347eb3d0 := strconv.Atoi(__obf_008ae55685cb56f6); __obf_e6c24855347eb3d0 == nil {
				__obf_99782ba885b4f4ff[__obf_13037cdadbf8d35e] = __obf_935763ab96242d2f
				continue
			}

			// float
			if __obf_5a07a2990ad85783, __obf_e6c24855347eb3d0 := strconv.ParseFloat(__obf_008ae55685cb56f6, 64); __obf_e6c24855347eb3d0 == nil {
				__obf_99782ba885b4f4ff[__obf_13037cdadbf8d35e] = __obf_5a07a2990ad85783
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_008ae55685cb56f6, "{") && strings.HasSuffix(__obf_008ae55685cb56f6, "}")) ||
				(strings.HasPrefix(__obf_008ae55685cb56f6, "[") && strings.HasSuffix(__obf_008ae55685cb56f6, "]")) {
				var __obf_82f3915011399cec any
				if __obf_e6c24855347eb3d0 := json.Unmarshal(__obf_1115ae3c80598454, &__obf_82f3915011399cec); __obf_e6c24855347eb3d0 == nil {
					__obf_99782ba885b4f4ff[__obf_13037cdadbf8d35e] = __obf_82f3915011399cec
					continue
				}
			}

			// 默认保留为 string
			__obf_99782ba885b4f4ff[__obf_13037cdadbf8d35e] = __obf_37836fc6d365a20e
		default:
			__obf_99782ba885b4f4ff[__obf_13037cdadbf8d35e] = __obf_1115ae3c80598454
		}
	}
	return __obf_99782ba885b4f4ff
}

func ToListMap(__obf_85accf2129aa6cd7 *sqlx.Rows) []map[string]any {
	var __obf_c560cfbfddbf86d2 []map[string]any
	for __obf_85accf2129aa6cd7.Next() {
		__obf_ed8a226a1985dce5 := make(map[string]any)
		if __obf_e6c24855347eb3d0 := __obf_85accf2129aa6cd7.MapScan(__obf_ed8a226a1985dce5); __obf_e6c24855347eb3d0 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_e6c24855347eb3d0.Error()))
			continue
		}

		__obf_c560cfbfddbf86d2 = append(__obf_c560cfbfddbf86d2, CleanSQLXMap(__obf_ed8a226a1985dce5))
	}
	return __obf_c560cfbfddbf86d2
}
