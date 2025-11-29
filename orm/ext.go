package __obf_ea545e4bdf748fd2

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
type FuncTx func(__obf_2623653de66b5673 *sqlx.Tx, __obf_15a4a0514baf9e2c any) error

func Open(__obf_72011884f9850b2c, __obf_1d4852f0ccf3e6d7 string, __obf_d09ad108b12e2148 time.Duration) (*DB, error) {
	__obf_d74aacde370b8778, __obf_3220344ba911fc77 := sqlx.Open(__obf_72011884f9850b2c, __obf_1d4852f0ccf3e6d7)
	if __obf_3220344ba911fc77 != nil {
		return nil, __obf_3220344ba911fc77
	}
	__obf_a6490b2f4d4d2871 := &DB{
		DB: __obf_d74aacde370b8778, __obf_d09ad108b12e2148: __obf_d09ad108b12e2148,
	}
	__obf_a6490b2f4d4d2871.__obf_3c74824899d4164e.
		New = func() any {
		return __obf_a6490b2f4d4d2871.__obf_b198426d4ba50371()
	}
	return __obf_a6490b2f4d4d2871, nil
}

type DB struct {
	*sqlx.DB
	__obf_d09ad108b12e2148 time.Duration
	__obf_3c74824899d4164e sync.Pool
}

func (__obf_d74aacde370b8778 *DB) __obf_b198426d4ba50371() *Context {
	return &Context{__obf_d74aacde370b8778: __obf_d74aacde370b8778}
}

// 获取一个`SQL`执行`Context`
func (__obf_d74aacde370b8778 *DB) Acquire() *Context {
	__obf_af5f7ab0755f3d15 := // 无需加锁，sync.Pool本身是线程安全的
		__obf_d74aacde370b8778.__obf_3c74824899d4164e.Get().(*Context)
	__obf_af5f7ab0755f3d15.__obf_12f6bb1fada500d2()
	return __obf_af5f7ab0755f3d15
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_d74aacde370b8778 *DB) AcquireTx(__obf_2623653de66b5673 *sqlx.Tx) *Context {
	__obf_af5f7ab0755f3d15 := __obf_d74aacde370b8778.Acquire()
	__obf_af5f7ab0755f3d15.__obf_2623653de66b5673 = __obf_2623653de66b5673
	return __obf_af5f7ab0755f3d15
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_d74aacde370b8778 *DB) WithTx(__obf_4b280fd46d014d73 FuncTx, __obf_15a4a0514baf9e2c any) (__obf_3220344ba911fc77 error) {
	var __obf_2623653de66b5673 *sqlx.Tx
	__obf_2623653de66b5673, __obf_3220344ba911fc77 = __obf_d74aacde370b8778.Beginx()
	if __obf_3220344ba911fc77 != nil {
		return
	}
	defer func() {
		if __obf_3220344ba911fc77 != nil && __obf_2623653de66b5673 != nil {
			__obf_3220344ba911fc77 = __obf_2623653de66b5673.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_3220344ba911fc77 = __obf_4b280fd46d014d73(__obf_2623653de66b5673, __obf_15a4a0514baf9e2c); __obf_3220344ba911fc77 != nil {
		return
	}
	__obf_3220344ba911fc77 = __obf_2623653de66b5673.Commit()
	return
}

type Context struct {
	__obf_d74aacde370b8778 *DB
	__obf_2623653de66b5673 *sqlx.Tx //事务
	sql                    string
	__obf_a91af539ddf108b8 string
	__obf_5e1fc84a9e7eebf4 []string
	__obf_e7224cd4fac6a299 []string
	__obf_d11208ea43059898 string
	__obf_abc2468c822e6c02 string
	__obf_19ae4cc7956ee352 string
	__obf_ccf0c4588a92d2e9 int64
	__obf_6ff5a973ed6715e5 int64
	__obf_15a4a0514baf9e2c []any
	__obf_9ffae2f186bfe4f7 bool
	__obf_6770b9463d1d7ff5 bool //排他锁
	//共享锁
}

func (__obf_af5f7ab0755f3d15 *Context) Name(__obf_a91af539ddf108b8 string) *Context {
	__obf_af5f7ab0755f3d15.__obf_a91af539ddf108b8 = __obf_a91af539ddf108b8
	return __obf_af5f7ab0755f3d15
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_af5f7ab0755f3d15 *Context) What(__obf_5e1fc84a9e7eebf4 []string) *Context {
	__obf_af5f7ab0755f3d15.__obf_5e1fc84a9e7eebf4 = __obf_5e1fc84a9e7eebf4
	return __obf_af5f7ab0755f3d15
}

func (__obf_af5f7ab0755f3d15 *Context) Where(__obf_d1a4a9e58b41ed33 string, __obf_15a4a0514baf9e2c ...any) *Context {
	__obf_af5f7ab0755f3d15.__obf_e7224cd4fac6a299 = append(__obf_af5f7ab0755f3d15.__obf_e7224cd4fac6a299, __obf_d1a4a9e58b41ed33)
	__obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c = append(__obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c,

		// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
		__obf_15a4a0514baf9e2c...)
	return __obf_af5f7ab0755f3d15
}

func (__obf_af5f7ab0755f3d15 *Context) WhereIn(__obf_e144436f864ee67b string, __obf_15a4a0514baf9e2c []any) *Context {
	__obf_27278b36a4aa8831 := len(__obf_15a4a0514baf9e2c)
	__obf_0d0b70c914abb3fa := make([]string, __obf_27278b36a4aa8831)
	for __obf_5f1ac33f5a6c1104 := 0; __obf_5f1ac33f5a6c1104 < __obf_27278b36a4aa8831; __obf_5f1ac33f5a6c1104++ {
		__obf_0d0b70c914abb3fa[__obf_5f1ac33f5a6c1104] = ParamMarker
	}
	__obf_eac03782c96b3640 := fmt.Sprintf("%s in (%s)", __obf_e144436f864ee67b, __obf_f3ea9e2199d82ed5(__obf_0d0b70c914abb3fa, SeqComma))
	return __obf_af5f7ab0755f3d15.Where(__obf_eac03782c96b3640, __obf_15a4a0514baf9e2c...)
}

func (__obf_af5f7ab0755f3d15 *Context) Order(__obf_d11208ea43059898 string) *Context {
	__obf_af5f7ab0755f3d15.__obf_d11208ea43059898 = __obf_d11208ea43059898
	return __obf_af5f7ab0755f3d15
}

func (__obf_af5f7ab0755f3d15 *Context) Limit(__obf_ccf0c4588a92d2e9 int64) *Context {
	__obf_af5f7ab0755f3d15.__obf_ccf0c4588a92d2e9 = __obf_ccf0c4588a92d2e9
	return __obf_af5f7ab0755f3d15
}

func (__obf_af5f7ab0755f3d15 *Context) Offset(__obf_6ff5a973ed6715e5 int64) *Context {
	__obf_af5f7ab0755f3d15.__obf_6ff5a973ed6715e5 = __obf_6ff5a973ed6715e5
	return __obf_af5f7ab0755f3d15
}

func (__obf_af5f7ab0755f3d15 *Context) Group(__obf_abc2468c822e6c02 string) *Context {
	__obf_af5f7ab0755f3d15.__obf_abc2468c822e6c02 = __obf_abc2468c822e6c02
	return __obf_af5f7ab0755f3d15
}

func (__obf_af5f7ab0755f3d15 *Context) Having(__obf_19ae4cc7956ee352 string, __obf_15a4a0514baf9e2c ...any) *Context {
	__obf_af5f7ab0755f3d15.__obf_19ae4cc7956ee352 = __obf_19ae4cc7956ee352
	__obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c = append(__obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c,

		// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
		__obf_15a4a0514baf9e2c...)
	return __obf_af5f7ab0755f3d15
}

func (__obf_af5f7ab0755f3d15 *Context) LockX() *Context {
	__obf_af5f7ab0755f3d15.__obf_9ffae2f186bfe4f7 = true
	return __obf_af5f7ab0755f3d15
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_af5f7ab0755f3d15 *Context) LockS() *Context {
	__obf_af5f7ab0755f3d15.__obf_6770b9463d1d7ff5 = true
	return __obf_af5f7ab0755f3d15
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_af5f7ab0755f3d15 *Context) FindMany(__obf_64d00c39ca426312 any) error {
	return __obf_af5f7ab0755f3d15.__obf_d27b101008e873ee(__obf_64d00c39ca426312, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_af5f7ab0755f3d15 *Context) FindById(__obf_64d00c39ca426312 any) error {
	return __obf_af5f7ab0755f3d15.__obf_d27b101008e873ee(__obf_64d00c39ca426312, SelectTypeOne)
}

// 插入
func (__obf_af5f7ab0755f3d15 *Context) Insert(__obf_b38a4b3f3a26b406 map[string]any) (sql.Result, error) {
	var (
		__obf_d503c65f4de0ce59 []string
		__obf_3191fd1aa9d8da0d []any
	)
	for __obf_5304b9e5e1368ecf, __obf_8a2e311a616cdf4b := range __obf_b38a4b3f3a26b406 {
		__obf_d503c65f4de0ce59 = append(__obf_d503c65f4de0ce59, __obf_5304b9e5e1368ecf)
		__obf_3191fd1aa9d8da0d = append(__obf_3191fd1aa9d8da0d, __obf_8a2e311a616cdf4b)
	}
	return __obf_af5f7ab0755f3d15.InsertBatch(__obf_d503c65f4de0ce59,

		// 批量插入
		__obf_3191fd1aa9d8da0d)
}

func (__obf_af5f7ab0755f3d15 *Context) InsertBatch(__obf_d503c65f4de0ce59 []string, __obf_b38a4b3f3a26b406 ...[]any) (sql.Result, error) {
	var (
		__obf_3191fd1aa9d8da0d []any
		__obf_798dd58c3394544f []string
	)
	for _, __obf_b02ba11dfb557d98 := range __obf_b38a4b3f3a26b406 {
		__obf_0d0b70c914abb3fa := make([]string, len(__obf_b02ba11dfb557d98))
		for __obf_5f1ac33f5a6c1104, __obf_8a2e311a616cdf4b := range __obf_b02ba11dfb557d98 {
			__obf_0d0b70c914abb3fa[__obf_5f1ac33f5a6c1104] = ParamMarker
			__obf_3191fd1aa9d8da0d = append(__obf_3191fd1aa9d8da0d, __obf_8a2e311a616cdf4b)
		}
		__obf_798dd58c3394544f = append(__obf_798dd58c3394544f, fmt.Sprintf("(%s)", __obf_f3ea9e2199d82ed5(__obf_0d0b70c914abb3fa, SeqComma)))
	}
	__obf_3c80b73d77602a41 := fmt.Sprintf("insert into %s (%s) values %s", __obf_af5f7ab0755f3d15.__obf_a91af539ddf108b8, __obf_f3ea9e2199d82ed5(__obf_d503c65f4de0ce59, SeqComma), __obf_f3ea9e2199d82ed5(__obf_798dd58c3394544f, SeqComma))
	return __obf_af5f7ab0755f3d15.__obf_e6a7261be0d8a638(__obf_3c80b73d77602a41,

		// 使用map更新
		__obf_3191fd1aa9d8da0d...)
}

func (__obf_af5f7ab0755f3d15 *Context) UpdateMap(__obf_15a4a0514baf9e2c map[string]any) (__obf_b8f841866c38d6bf int64, __obf_3220344ba911fc77 error) {
	var (
		__obf_3191fd1aa9d8da0d []any
		__obf_e29244e8fe131909 []string
	)
	for __obf_5304b9e5e1368ecf, __obf_8a2e311a616cdf4b := range __obf_15a4a0514baf9e2c {
		__obf_3191fd1aa9d8da0d = append(__obf_3191fd1aa9d8da0d, __obf_8a2e311a616cdf4b)
		__obf_e29244e8fe131909 = append(__obf_e29244e8fe131909, fmt.Sprintf("%s=%s", __obf_5304b9e5e1368ecf, ParamMarker))
	}
	__obf_3e0763d4a97fcc22 := __obf_f3ea9e2199d82ed5(__obf_e29244e8fe131909, SeqComma)
	__obf_b8f841866c38d6bf, __obf_3220344ba911fc77 = __obf_af5f7ab0755f3d15.Update(__obf_3e0763d4a97fcc22, __obf_3191fd1aa9d8da0d...)
	return
}

// 更新
func (__obf_af5f7ab0755f3d15 *Context) Update(__obf_3e0763d4a97fcc22 string, __obf_15a4a0514baf9e2c ...any) (__obf_b8f841866c38d6bf int64, __obf_3220344ba911fc77 error) {
	__obf_b1a100942d3cf161 := "update %s set %s %s"
	__obf_d1a4a9e58b41ed33 := __obf_5a81c120a97f31bb(__obf_af5f7ab0755f3d15.__obf_e7224cd4fac6a299, Grouping)
	__obf_3c80b73d77602a41 := fmt.Sprintf(__obf_b1a100942d3cf161, __obf_af5f7ab0755f3d15.__obf_a91af539ddf108b8, __obf_3e0763d4a97fcc22, __obf_d1a4a9e58b41ed33)
	__obf_3191fd1aa9d8da0d := append(__obf_15a4a0514baf9e2c, __obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c...)
	var __obf_b0679735c3db1659 sql.Result
	__obf_b0679735c3db1659, __obf_3220344ba911fc77 = __obf_af5f7ab0755f3d15.__obf_e6a7261be0d8a638(__obf_3c80b73d77602a41, __obf_3191fd1aa9d8da0d...)
	if __obf_3220344ba911fc77 != nil {
		return
	}
	__obf_b8f841866c38d6bf, __obf_3220344ba911fc77 = __obf_b0679735c3db1659.RowsAffected()
	return
}

// 删除
func (__obf_af5f7ab0755f3d15 *Context) Delete() (__obf_b8f841866c38d6bf int64, __obf_3220344ba911fc77 error) {
	__obf_b1a100942d3cf161 := "delete from %s %s"
	__obf_d1a4a9e58b41ed33 := __obf_5a81c120a97f31bb(__obf_af5f7ab0755f3d15.__obf_e7224cd4fac6a299, Grouping)
	__obf_3c80b73d77602a41 := fmt.Sprintf(__obf_b1a100942d3cf161, __obf_af5f7ab0755f3d15.__obf_a91af539ddf108b8, __obf_d1a4a9e58b41ed33)
	var __obf_b0679735c3db1659 sql.Result
	__obf_b0679735c3db1659, __obf_3220344ba911fc77 = __obf_af5f7ab0755f3d15.__obf_e6a7261be0d8a638(__obf_3c80b73d77602a41, __obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c...)
	if __obf_3220344ba911fc77 != nil {
		return
	}
	__obf_b8f841866c38d6bf, __obf_3220344ba911fc77 = __obf_b0679735c3db1659.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_af5f7ab0755f3d15 *Context) Select(__obf_64d00c39ca426312 any, sql string, __obf_15a4a0514baf9e2c ...any) error {
	__obf_af5f7ab0755f3d15.
		sql = sql
	__obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c = __obf_15a4a0514baf9e2c
	return __obf_af5f7ab0755f3d15.__obf_d27b101008e873ee(__obf_64d00c39ca426312, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_af5f7ab0755f3d15 *Context) Get(__obf_64d00c39ca426312 any, sql string, __obf_15a4a0514baf9e2c ...any) error {
	__obf_af5f7ab0755f3d15.
		sql = sql
	__obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c = __obf_15a4a0514baf9e2c
	return __obf_af5f7ab0755f3d15.__obf_d27b101008e873ee(__obf_64d00c39ca426312, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_af5f7ab0755f3d15 *Context) Exec(sql string, __obf_15a4a0514baf9e2c ...any) (sql.Result, error) {
	return __obf_af5f7ab0755f3d15.__obf_e6a7261be0d8a638(sql, __obf_15a4a0514baf9e2c...)
}

// 创建表
func (__obf_af5f7ab0755f3d15 *Context) Create(sql string) (sql.Result, error) {
	return __obf_af5f7ab0755f3d15.

		// 删除表
		__obf_e6a7261be0d8a638(sql)
}

func (__obf_af5f7ab0755f3d15 *Context) Drop() (sql.Result, error) {
	return __obf_af5f7ab0755f3d15.__obf_e6a7261be0d8a638(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_af5f7ab0755f3d15.

		/////////////////////////private methods//////////////////////
		__obf_a91af539ddf108b8))
}

// 重置Context
func (__obf_af5f7ab0755f3d15 *Context) __obf_12f6bb1fada500d2() *Context {
	__obf_af5f7ab0755f3d15.
		sql = ""
	__obf_af5f7ab0755f3d15.__obf_a91af539ddf108b8 = ""
	__obf_af5f7ab0755f3d15.__obf_5e1fc84a9e7eebf4 = []string{}
	__obf_af5f7ab0755f3d15.__obf_e7224cd4fac6a299 = []string{}
	__obf_af5f7ab0755f3d15.__obf_d11208ea43059898 = ""
	__obf_af5f7ab0755f3d15.__obf_abc2468c822e6c02 = ""
	__obf_af5f7ab0755f3d15.__obf_19ae4cc7956ee352 = ""
	__obf_af5f7ab0755f3d15.__obf_ccf0c4588a92d2e9 = 0
	__obf_af5f7ab0755f3d15.__obf_6ff5a973ed6715e5 = 0
	__obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c = []any{}
	__obf_af5f7ab0755f3d15.__obf_2623653de66b5673 = nil
	__obf_af5f7ab0755f3d15.__obf_6770b9463d1d7ff5 = false
	__obf_af5f7ab0755f3d15.__obf_9ffae2f186bfe4f7 = false
	return __obf_af5f7ab0755f3d15
}

// 查询方法
func (__obf_af5f7ab0755f3d15 *Context) __obf_d27b101008e873ee(__obf_64d00c39ca426312 any, __obf_351c042f57f61b5c int) (__obf_3220344ba911fc77 error) {
	defer __obf_af5f7ab0755f3d15.__obf_d74aacde370b8778.__obf_3c74824899d4164e.Put(__obf_af5f7ab0755f3d15)
	__obf_4c64a930a6a98f4c, __obf_c84c07af9eb15722 := context.WithTimeout(context.Background(), __obf_af5f7ab0755f3d15.__obf_d74aacde370b8778.__obf_d09ad108b12e2148)
	defer __obf_c84c07af9eb15722()
	if __obf_af5f7ab0755f3d15.sql == "" {
		__obf_af5f7ab0755f3d15.
			sql = __obf_af5f7ab0755f3d15.__obf_8c5c1eebbebe15cc(__obf_64d00c39ca426312)
	}
	switch __obf_351c042f57f61b5c {
	case SelectTypeOne:
		if __obf_af5f7ab0755f3d15.__obf_2623653de66b5673 != nil {
			__obf_3220344ba911fc77 = __obf_af5f7ab0755f3d15.__obf_2623653de66b5673.GetContext(__obf_4c64a930a6a98f4c, __obf_64d00c39ca426312, __obf_af5f7ab0755f3d15.sql, __obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c...)
		} else {
			__obf_3220344ba911fc77 = __obf_af5f7ab0755f3d15.__obf_d74aacde370b8778.GetContext(__obf_4c64a930a6a98f4c, __obf_64d00c39ca426312, __obf_af5f7ab0755f3d15.sql, __obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c...)
		}
	case SelectTypeMany:
		if __obf_af5f7ab0755f3d15.__obf_2623653de66b5673 != nil {
			__obf_3220344ba911fc77 = __obf_af5f7ab0755f3d15.__obf_2623653de66b5673.SelectContext(__obf_4c64a930a6a98f4c, __obf_64d00c39ca426312, __obf_af5f7ab0755f3d15.sql, __obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c...)
		} else {
			__obf_3220344ba911fc77 = __obf_af5f7ab0755f3d15.__obf_d74aacde370b8778.SelectContext(__obf_4c64a930a6a98f4c, __obf_64d00c39ca426312, __obf_af5f7ab0755f3d15.sql, __obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_af5f7ab0755f3d15 *Context) __obf_e6a7261be0d8a638(__obf_3c80b73d77602a41 string, __obf_15a4a0514baf9e2c ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_3c80b73d77602a41, __obf_15a4a0514baf9e2c)
	defer __obf_af5f7ab0755f3d15.__obf_d74aacde370b8778.__obf_3c74824899d4164e.Put(__obf_af5f7ab0755f3d15)
	__obf_4c64a930a6a98f4c, __obf_c84c07af9eb15722 := context.WithTimeout(context.Background(), __obf_af5f7ab0755f3d15.__obf_d74aacde370b8778.__obf_d09ad108b12e2148)
	defer __obf_c84c07af9eb15722()

	var __obf_cfe874928ef3d912 sqlx.ExecerContext
	if __obf_af5f7ab0755f3d15.__obf_2623653de66b5673 != nil {
		__obf_cfe874928ef3d912 = __obf_af5f7ab0755f3d15.__obf_2623653de66b5673
	} else {
		__obf_cfe874928ef3d912 = __obf_af5f7ab0755f3d15.__obf_d74aacde370b8778
	}
	return __obf_cfe874928ef3d912.ExecContext(__obf_4c64a930a6a98f4c, __obf_3c80b73d77602a41,

		// select查询语句的拼接
		__obf_15a4a0514baf9e2c...)
}

func (__obf_af5f7ab0755f3d15 *Context) __obf_8c5c1eebbebe15cc(__obf_64d00c39ca426312 any) string {
	var __obf_2e743112b6ec0f89 []string
	__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, "select")
	if len(__obf_af5f7ab0755f3d15.__obf_5e1fc84a9e7eebf4) != 0 {
		__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, __obf_f3ea9e2199d82ed5(__obf_af5f7ab0755f3d15.__obf_5e1fc84a9e7eebf4, SeqComma))
	} else {
		__obf_fccd3b668e35de00 := // 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
			// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
			__obf_8b988f308b80fc1c(__obf_64d00c39ca426312)
		if len(__obf_fccd3b668e35de00) > 0 {
			__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, __obf_f3ea9e2199d82ed5(__obf_fccd3b668e35de00, SeqComma))
		} else {
			__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, "*")
		}
	}
	__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, "from "+__obf_af5f7ab0755f3d15.__obf_a91af539ddf108b8)
	if len(__obf_af5f7ab0755f3d15.__obf_e7224cd4fac6a299) != 0 {
		__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, __obf_5a81c120a97f31bb(__obf_af5f7ab0755f3d15.__obf_e7224cd4fac6a299, Grouping))
	}

	if __obf_af5f7ab0755f3d15.__obf_abc2468c822e6c02 != "" {
		__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, "group by "+__obf_af5f7ab0755f3d15.__obf_abc2468c822e6c02)
	}

	if __obf_af5f7ab0755f3d15.__obf_19ae4cc7956ee352 != "" {
		__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, "having "+__obf_af5f7ab0755f3d15.__obf_19ae4cc7956ee352)
	}

	if __obf_af5f7ab0755f3d15.__obf_d11208ea43059898 != "" {
		__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, "order by "+__obf_af5f7ab0755f3d15.__obf_d11208ea43059898)
	}

	if __obf_af5f7ab0755f3d15.__obf_ccf0c4588a92d2e9 != 0 {
		__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, fmt.Sprintf("limit %d, %d", __obf_af5f7ab0755f3d15.__obf_6ff5a973ed6715e5, __obf_af5f7ab0755f3d15.__obf_ccf0c4588a92d2e9))
	}
	if __obf_af5f7ab0755f3d15.__obf_6770b9463d1d7ff5 {
		__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, "lock in share mode")
	}
	if __obf_af5f7ab0755f3d15.__obf_9ffae2f186bfe4f7 {
		__obf_2e743112b6ec0f89 = append(__obf_2e743112b6ec0f89, "for update")
	}
	sql := __obf_f3ea9e2199d82ed5(__obf_2e743112b6ec0f89, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_af5f7ab0755f3d15.__obf_15a4a0514baf9e2c)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_5a81c120a97f31bb(__obf_e7224cd4fac6a299 []string, __obf_83f71c935a37bb33 string) string {
	if len(__obf_e7224cd4fac6a299) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_e7224cd4fac6a299, __obf_83f71c935a37bb33))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_f3ea9e2199d82ed5(__obf_15a4a0514baf9e2c []string, __obf_28a1a9880009eb85 string) string {
	return strings.Join(__obf_15a4a0514baf9e2c,

		// 解析对象中的 `db tag`
		// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
		__obf_28a1a9880009eb85)
}

func __obf_8b988f308b80fc1c(__obf_64d00c39ca426312 any) (__obf_d503c65f4de0ce59 []string) {
	__obf_43a03d8b7535a35f := reflect.ValueOf(__obf_64d00c39ca426312)
	__obf_ee551e5b6b258d3e := __obf_43a03d8b7535a35f.Type().Elem()
	var __obf_84a111e80a7c6126 reflect.Type
	if __obf_ee551e5b6b258d3e.Kind() == reflect.Slice {
		__obf_84a111e80a7c6126 = __obf_ee551e5b6b258d3e.Elem()
	} else {
		__obf_84a111e80a7c6126 = __obf_ee551e5b6b258d3e
	}
	for __obf_5f1ac33f5a6c1104 := 0; __obf_5f1ac33f5a6c1104 < __obf_84a111e80a7c6126.NumField(); __obf_5f1ac33f5a6c1104++ {
		__obf_04bbd256baa3a464 := __obf_84a111e80a7c6126.Field(__obf_5f1ac33f5a6c1104).Tag.Get(DBTag)
		if __obf_04bbd256baa3a464 != "" {
			__obf_d503c65f4de0ce59 = append(__obf_d503c65f4de0ce59, __obf_04bbd256baa3a464)
		}
	}
	return
}

func CleanSQLXMap(__obf_f4a7ec7378561e42 map[string]any) map[string]any {
	for __obf_5304b9e5e1368ecf, __obf_8a2e311a616cdf4b := range __obf_f4a7ec7378561e42 {
		switch __obf_abd89eb645093144 := __obf_8a2e311a616cdf4b.(type) {
		case []byte:
			__obf_4c5a866ce1415830 := string(__obf_abd89eb645093144)
			__obf_851c13248a42bdf6 := // 尝试智能识别具体类型
				strings.TrimSpace(__obf_4c5a866ce1415830)

			// bool
			if strings.EqualFold(__obf_851c13248a42bdf6, "true") || strings.EqualFold(__obf_851c13248a42bdf6, "false") {
				if __obf_00aa56e70a15e735, __obf_3220344ba911fc77 := strconv.ParseBool(__obf_851c13248a42bdf6); __obf_3220344ba911fc77 == nil {
					__obf_f4a7ec7378561e42[__obf_5304b9e5e1368ecf] = __obf_00aa56e70a15e735
					continue
				}
			}

			// int
			if __obf_5f1ac33f5a6c1104, __obf_3220344ba911fc77 := strconv.Atoi(__obf_851c13248a42bdf6); __obf_3220344ba911fc77 == nil {
				__obf_f4a7ec7378561e42[__obf_5304b9e5e1368ecf] = __obf_5f1ac33f5a6c1104
				continue
			}

			// float
			if __obf_af66979cc7b3eb24, __obf_3220344ba911fc77 := strconv.ParseFloat(__obf_851c13248a42bdf6, 64); __obf_3220344ba911fc77 == nil {
				__obf_f4a7ec7378561e42[__obf_5304b9e5e1368ecf] = __obf_af66979cc7b3eb24
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_851c13248a42bdf6, "{") && strings.HasSuffix(__obf_851c13248a42bdf6, "}")) ||
				(strings.HasPrefix(__obf_851c13248a42bdf6, "[") && strings.HasSuffix(__obf_851c13248a42bdf6, "]")) {
				var __obf_2d333a32224c9e9e any
				if __obf_3220344ba911fc77 := json.Unmarshal(__obf_abd89eb645093144, &__obf_2d333a32224c9e9e); __obf_3220344ba911fc77 == nil {
					__obf_f4a7ec7378561e42[__obf_5304b9e5e1368ecf] = __obf_2d333a32224c9e9e
					continue
				}
			}
			__obf_f4a7ec7378561e42[ // 默认保留为 string
			__obf_5304b9e5e1368ecf] = __obf_4c5a866ce1415830
		default:
			__obf_f4a7ec7378561e42[__obf_5304b9e5e1368ecf] = __obf_abd89eb645093144
		}
	}
	return __obf_f4a7ec7378561e42
}

func ToListMap(__obf_84494ca656481f28 *sqlx.Rows) []map[string]any {
	var __obf_b0679735c3db1659 []map[string]any
	for __obf_84494ca656481f28.Next() {
		__obf_02cdb615ee1e03b2 := make(map[string]any)
		if __obf_3220344ba911fc77 := __obf_84494ca656481f28.MapScan(__obf_02cdb615ee1e03b2); __obf_3220344ba911fc77 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_3220344ba911fc77.Error()))
			continue
		}
		__obf_b0679735c3db1659 = append(__obf_b0679735c3db1659, CleanSQLXMap(__obf_02cdb615ee1e03b2))
	}
	return __obf_b0679735c3db1659
}
