package __obf_214d1e062aee9185

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
type FuncTx func(__obf_300af38d91fabd70 *sqlx.Tx, __obf_ede102380a19106b any) error

func Open(__obf_b0d95b286d0abce1, __obf_3fa7fdb986369d43 string, __obf_d40bafd7b95c975a time.Duration) (*DB, error) {
	__obf_9858352bdc58a6fc, __obf_63349b4c89eea887 := sqlx.Open(__obf_b0d95b286d0abce1, __obf_3fa7fdb986369d43)
	if __obf_63349b4c89eea887 != nil {
		return nil, __obf_63349b4c89eea887
	}
	__obf_c19d7b5d3b7a82bb := &DB{
		DB: __obf_9858352bdc58a6fc, __obf_d40bafd7b95c975a: __obf_d40bafd7b95c975a,
	}
	__obf_c19d7b5d3b7a82bb.__obf_5d8be8432bb33d38.
		New = func() any {
		return __obf_c19d7b5d3b7a82bb.__obf_f04f212d41574d84()
	}
	return __obf_c19d7b5d3b7a82bb, nil
}

type DB struct {
	*sqlx.DB
	__obf_d40bafd7b95c975a time.Duration
	__obf_5d8be8432bb33d38 sync.Pool
}

func (__obf_9858352bdc58a6fc *DB) __obf_f04f212d41574d84() *Context {
	return &Context{__obf_9858352bdc58a6fc: __obf_9858352bdc58a6fc}
}

// 获取一个`SQL`执行`Context`
func (__obf_9858352bdc58a6fc *DB) Acquire() *Context {
	__obf_1397b4a37b8546eb := // 无需加锁，sync.Pool本身是线程安全的
		__obf_9858352bdc58a6fc.__obf_5d8be8432bb33d38.Get().(*Context)
	__obf_1397b4a37b8546eb.__obf_6091f70bee9468bb()
	return __obf_1397b4a37b8546eb
}

// 获取一个带有事务`tx`的`SQL`执行`Context`
func (__obf_9858352bdc58a6fc *DB) AcquireTx(__obf_300af38d91fabd70 *sqlx.Tx) *Context {
	__obf_1397b4a37b8546eb := __obf_9858352bdc58a6fc.Acquire()
	__obf_1397b4a37b8546eb.__obf_300af38d91fabd70 = __obf_300af38d91fabd70
	return __obf_1397b4a37b8546eb
}

// 优雅的开启事务
// 只能用装饰器了，相当于注入了一个事务的上下文对象
// 除了可以统一处理开启事务的代码，好像也没看到啥好处，而且还限制了参数的传递，只能传递一个参数，所以多参数就弄成一个对象传递吧
// 返回值也就只有异常，所以如果需要返回什么数据的，就直接搞到异常里面吧，我也不知道怎么搞...
// 最后，不要搞嵌套事务
func (__obf_9858352bdc58a6fc *DB) WithTx(__obf_ea9a2e1d584dfb6c FuncTx, __obf_ede102380a19106b any) (__obf_63349b4c89eea887 error) {
	var __obf_300af38d91fabd70 *sqlx.Tx
	__obf_300af38d91fabd70, __obf_63349b4c89eea887 = __obf_9858352bdc58a6fc.Beginx()
	if __obf_63349b4c89eea887 != nil {
		return
	}
	defer func() {
		if __obf_63349b4c89eea887 != nil && __obf_300af38d91fabd70 != nil {
			__obf_63349b4c89eea887 = __obf_300af38d91fabd70.Rollback()
		}
	}()

	// 调用外部函数
	if __obf_63349b4c89eea887 = __obf_ea9a2e1d584dfb6c(__obf_300af38d91fabd70, __obf_ede102380a19106b); __obf_63349b4c89eea887 != nil {
		return
	}
	__obf_63349b4c89eea887 = __obf_300af38d91fabd70.Commit()
	return
}

type Context struct {
	__obf_9858352bdc58a6fc *DB
	__obf_300af38d91fabd70 *sqlx.Tx //事务
	sql                    string
	__obf_e6237083261f4fc2 string
	__obf_4939871f1153a0a4 []string
	__obf_502636b0963e1e47 []string
	__obf_e82b41174897944b string
	__obf_b002972f2f4e9511 string
	__obf_fff491dab590b800 string
	__obf_d9b389e6a9615f76 int64
	__obf_1726fa49ac928860 int64
	__obf_ede102380a19106b []any
	__obf_15a31c61a999fdce bool
	__obf_62fa3336f332da73 bool //排他锁
	//共享锁
}

func (__obf_1397b4a37b8546eb *Context) Name(__obf_e6237083261f4fc2 string) *Context {
	__obf_1397b4a37b8546eb.__obf_e6237083261f4fc2 = __obf_e6237083261f4fc2
	return __obf_1397b4a37b8546eb
}

// 查询字段
// 如果不指定查询字段，默认使用传递的对象中的标签`db`指定的字段，如果没有指定`db`标签则使用`*`代替
// 使用`*`以后增加数据库字段可能会导致老的查询出错，对兼容性不好，可能是`orm`这个库的问题
func (__obf_1397b4a37b8546eb *Context) What(__obf_4939871f1153a0a4 []string) *Context {
	__obf_1397b4a37b8546eb.__obf_4939871f1153a0a4 = __obf_4939871f1153a0a4
	return __obf_1397b4a37b8546eb
}

func (__obf_1397b4a37b8546eb *Context) Where(__obf_9c0f3cdc51a1c12b string, __obf_ede102380a19106b ...any) *Context {
	__obf_1397b4a37b8546eb.__obf_502636b0963e1e47 = append(__obf_1397b4a37b8546eb.__obf_502636b0963e1e47, __obf_9c0f3cdc51a1c12b)
	__obf_1397b4a37b8546eb.__obf_ede102380a19106b = append(__obf_1397b4a37b8546eb.__obf_ede102380a19106b,

		// 指定字段和字段的可取值，自动拼接成 `field in (?,?)` 形式，`args`必须是 `[]any`类型，"严格"的类型系统，蛤...
		__obf_ede102380a19106b...)
	return __obf_1397b4a37b8546eb
}

func (__obf_1397b4a37b8546eb *Context) WhereIn(__obf_96d2ca3394c04fbe string, __obf_ede102380a19106b []any) *Context {
	__obf_076fabfbf373d123 := len(__obf_ede102380a19106b)
	__obf_686b8d8da991702f := make([]string, __obf_076fabfbf373d123)
	for __obf_7c6c9e1512e37985 := 0; __obf_7c6c9e1512e37985 < __obf_076fabfbf373d123; __obf_7c6c9e1512e37985++ {
		__obf_686b8d8da991702f[__obf_7c6c9e1512e37985] = ParamMarker
	}
	__obf_d4bc350b0298740b := fmt.Sprintf("%s in (%s)", __obf_96d2ca3394c04fbe, __obf_8efd0bdb5b054c71(__obf_686b8d8da991702f, SeqComma))
	return __obf_1397b4a37b8546eb.Where(__obf_d4bc350b0298740b, __obf_ede102380a19106b...)
}

func (__obf_1397b4a37b8546eb *Context) Order(__obf_e82b41174897944b string) *Context {
	__obf_1397b4a37b8546eb.__obf_e82b41174897944b = __obf_e82b41174897944b
	return __obf_1397b4a37b8546eb
}

func (__obf_1397b4a37b8546eb *Context) Limit(__obf_d9b389e6a9615f76 int64) *Context {
	__obf_1397b4a37b8546eb.__obf_d9b389e6a9615f76 = __obf_d9b389e6a9615f76
	return __obf_1397b4a37b8546eb
}

func (__obf_1397b4a37b8546eb *Context) Offset(__obf_1726fa49ac928860 int64) *Context {
	__obf_1397b4a37b8546eb.__obf_1726fa49ac928860 = __obf_1726fa49ac928860
	return __obf_1397b4a37b8546eb
}

func (__obf_1397b4a37b8546eb *Context) Group(__obf_b002972f2f4e9511 string) *Context {
	__obf_1397b4a37b8546eb.__obf_b002972f2f4e9511 = __obf_b002972f2f4e9511
	return __obf_1397b4a37b8546eb
}

func (__obf_1397b4a37b8546eb *Context) Having(__obf_fff491dab590b800 string, __obf_ede102380a19106b ...any) *Context {
	__obf_1397b4a37b8546eb.__obf_fff491dab590b800 = __obf_fff491dab590b800
	__obf_1397b4a37b8546eb.__obf_ede102380a19106b = append(__obf_1397b4a37b8546eb.__obf_ede102380a19106b,

		// 加排他锁(X锁)，不保证与`LockS`互斥，自己保证
		__obf_ede102380a19106b...)
	return __obf_1397b4a37b8546eb
}

func (__obf_1397b4a37b8546eb *Context) LockX() *Context {
	__obf_1397b4a37b8546eb.__obf_15a31c61a999fdce = true
	return __obf_1397b4a37b8546eb
}

// 加共享锁(S锁)，不保证与`LockX`互斥，自己保证
func (__obf_1397b4a37b8546eb *Context) LockS() *Context {
	__obf_1397b4a37b8546eb.__obf_62fa3336f332da73 = true
	return __obf_1397b4a37b8546eb
}

// 查询多条记录，参数传入一个数组的指针，eg: &[]Little
func (__obf_1397b4a37b8546eb *Context) FindMany(__obf_02706219b5e5cf33 any) error {
	return __obf_1397b4a37b8546eb.__obf_66a128019ff47f8a(__obf_02706219b5e5cf33, SelectTypeMany)
}

// 查询一条记录，参数传入一个对象指针
func (__obf_1397b4a37b8546eb *Context) FindById(__obf_02706219b5e5cf33 any) error {
	return __obf_1397b4a37b8546eb.__obf_66a128019ff47f8a(__obf_02706219b5e5cf33, SelectTypeOne)
}

// 插入
func (__obf_1397b4a37b8546eb *Context) Insert(__obf_e681f73ad6c2f4f7 map[string]any) (sql.Result, error) {
	var (
		__obf_deabc6b2d5f62f57 []string
		__obf_afadf0e8122acde5 []any
	)
	for __obf_9b67b54491941e78, __obf_cf9e8016bdd91675 := range __obf_e681f73ad6c2f4f7 {
		__obf_deabc6b2d5f62f57 = append(__obf_deabc6b2d5f62f57, __obf_9b67b54491941e78)
		__obf_afadf0e8122acde5 = append(__obf_afadf0e8122acde5, __obf_cf9e8016bdd91675)
	}
	return __obf_1397b4a37b8546eb.InsertBatch(__obf_deabc6b2d5f62f57,

		// 批量插入
		__obf_afadf0e8122acde5)
}

func (__obf_1397b4a37b8546eb *Context) InsertBatch(__obf_deabc6b2d5f62f57 []string, __obf_e681f73ad6c2f4f7 ...[]any) (sql.Result, error) {
	var (
		__obf_afadf0e8122acde5 []any
		__obf_323f00d39aea4f98 []string
	)
	for _, __obf_e79e7bf9896a574e := range __obf_e681f73ad6c2f4f7 {
		__obf_686b8d8da991702f := make([]string, len(__obf_e79e7bf9896a574e))
		for __obf_7c6c9e1512e37985, __obf_cf9e8016bdd91675 := range __obf_e79e7bf9896a574e {
			__obf_686b8d8da991702f[__obf_7c6c9e1512e37985] = ParamMarker
			__obf_afadf0e8122acde5 = append(__obf_afadf0e8122acde5, __obf_cf9e8016bdd91675)
		}
		__obf_323f00d39aea4f98 = append(__obf_323f00d39aea4f98, fmt.Sprintf("(%s)", __obf_8efd0bdb5b054c71(__obf_686b8d8da991702f, SeqComma)))
	}
	__obf_8350072420eebe14 := fmt.Sprintf("insert into %s (%s) values %s", __obf_1397b4a37b8546eb.__obf_e6237083261f4fc2, __obf_8efd0bdb5b054c71(__obf_deabc6b2d5f62f57, SeqComma), __obf_8efd0bdb5b054c71(__obf_323f00d39aea4f98, SeqComma))
	return __obf_1397b4a37b8546eb.__obf_31c51f7c54ad4f2d(__obf_8350072420eebe14,

		// 使用map更新
		__obf_afadf0e8122acde5...)
}

func (__obf_1397b4a37b8546eb *Context) UpdateMap(__obf_ede102380a19106b map[string]any) (__obf_4b703263387a0dc8 int64, __obf_63349b4c89eea887 error) {
	var (
		__obf_afadf0e8122acde5 []any
		__obf_5e82aa24cb399c0f []string
	)
	for __obf_9b67b54491941e78, __obf_cf9e8016bdd91675 := range __obf_ede102380a19106b {
		__obf_afadf0e8122acde5 = append(__obf_afadf0e8122acde5, __obf_cf9e8016bdd91675)
		__obf_5e82aa24cb399c0f = append(__obf_5e82aa24cb399c0f, fmt.Sprintf("%s=%s", __obf_9b67b54491941e78, ParamMarker))
	}
	__obf_fc63a565b6280129 := __obf_8efd0bdb5b054c71(__obf_5e82aa24cb399c0f, SeqComma)
	__obf_4b703263387a0dc8, __obf_63349b4c89eea887 = __obf_1397b4a37b8546eb.Update(__obf_fc63a565b6280129, __obf_afadf0e8122acde5...)
	return
}

// 更新
func (__obf_1397b4a37b8546eb *Context) Update(__obf_fc63a565b6280129 string, __obf_ede102380a19106b ...any) (__obf_4b703263387a0dc8 int64, __obf_63349b4c89eea887 error) {
	__obf_0854773979a26c23 := "update %s set %s %s"
	__obf_9c0f3cdc51a1c12b := __obf_e8b95ee39de307d2(__obf_1397b4a37b8546eb.__obf_502636b0963e1e47, Grouping)
	__obf_8350072420eebe14 := fmt.Sprintf(__obf_0854773979a26c23, __obf_1397b4a37b8546eb.__obf_e6237083261f4fc2, __obf_fc63a565b6280129, __obf_9c0f3cdc51a1c12b)
	__obf_afadf0e8122acde5 := append(__obf_ede102380a19106b, __obf_1397b4a37b8546eb.__obf_ede102380a19106b...)
	var __obf_1576eeccdf72c56c sql.Result
	__obf_1576eeccdf72c56c, __obf_63349b4c89eea887 = __obf_1397b4a37b8546eb.__obf_31c51f7c54ad4f2d(__obf_8350072420eebe14, __obf_afadf0e8122acde5...)
	if __obf_63349b4c89eea887 != nil {
		return
	}
	__obf_4b703263387a0dc8, __obf_63349b4c89eea887 = __obf_1576eeccdf72c56c.RowsAffected()
	return
}

// 删除
func (__obf_1397b4a37b8546eb *Context) Delete() (__obf_4b703263387a0dc8 int64, __obf_63349b4c89eea887 error) {
	__obf_0854773979a26c23 := "delete from %s %s"
	__obf_9c0f3cdc51a1c12b := __obf_e8b95ee39de307d2(__obf_1397b4a37b8546eb.__obf_502636b0963e1e47, Grouping)
	__obf_8350072420eebe14 := fmt.Sprintf(__obf_0854773979a26c23, __obf_1397b4a37b8546eb.__obf_e6237083261f4fc2, __obf_9c0f3cdc51a1c12b)
	var __obf_1576eeccdf72c56c sql.Result
	__obf_1576eeccdf72c56c, __obf_63349b4c89eea887 = __obf_1397b4a37b8546eb.__obf_31c51f7c54ad4f2d(__obf_8350072420eebe14, __obf_1397b4a37b8546eb.__obf_ede102380a19106b...)
	if __obf_63349b4c89eea887 != nil {
		return
	}
	__obf_4b703263387a0dc8, __obf_63349b4c89eea887 = __obf_1576eeccdf72c56c.RowsAffected()
	return
}

// 查询多条记录，直接使用给定的`sql`和`args`
func (__obf_1397b4a37b8546eb *Context) Select(__obf_02706219b5e5cf33 any, sql string, __obf_ede102380a19106b ...any) error {
	__obf_1397b4a37b8546eb.
		sql = sql
	__obf_1397b4a37b8546eb.__obf_ede102380a19106b = __obf_ede102380a19106b
	return __obf_1397b4a37b8546eb.__obf_66a128019ff47f8a(__obf_02706219b5e5cf33, SelectTypeMany)
}

// 查询单条记录，直接使用给定的`sql`和`args`
func (__obf_1397b4a37b8546eb *Context) Get(__obf_02706219b5e5cf33 any, sql string, __obf_ede102380a19106b ...any) error {
	__obf_1397b4a37b8546eb.
		sql = sql
	__obf_1397b4a37b8546eb.__obf_ede102380a19106b = __obf_ede102380a19106b
	return __obf_1397b4a37b8546eb.__obf_66a128019ff47f8a(__obf_02706219b5e5cf33, SelectTypeOne)
}

// 直接执行操作，直接使用给定的`sql`和`args`
func (__obf_1397b4a37b8546eb *Context) Exec(sql string, __obf_ede102380a19106b ...any) (sql.Result, error) {
	return __obf_1397b4a37b8546eb.__obf_31c51f7c54ad4f2d(sql, __obf_ede102380a19106b...)
}

// 创建表
func (__obf_1397b4a37b8546eb *Context) Create(sql string) (sql.Result, error) {
	return __obf_1397b4a37b8546eb.

		// 删除表
		__obf_31c51f7c54ad4f2d(sql)
}

func (__obf_1397b4a37b8546eb *Context) Drop() (sql.Result, error) {
	return __obf_1397b4a37b8546eb.__obf_31c51f7c54ad4f2d(fmt.Sprintf("DROP TABLE IF EXISTS %s", __obf_1397b4a37b8546eb.

		/////////////////////////private methods//////////////////////
		__obf_e6237083261f4fc2))
}

// 重置Context
func (__obf_1397b4a37b8546eb *Context) __obf_6091f70bee9468bb() *Context {
	__obf_1397b4a37b8546eb.
		sql = ""
	__obf_1397b4a37b8546eb.__obf_e6237083261f4fc2 = ""
	__obf_1397b4a37b8546eb.__obf_4939871f1153a0a4 = []string{}
	__obf_1397b4a37b8546eb.__obf_502636b0963e1e47 = []string{}
	__obf_1397b4a37b8546eb.__obf_e82b41174897944b = ""
	__obf_1397b4a37b8546eb.__obf_b002972f2f4e9511 = ""
	__obf_1397b4a37b8546eb.__obf_fff491dab590b800 = ""
	__obf_1397b4a37b8546eb.__obf_d9b389e6a9615f76 = 0
	__obf_1397b4a37b8546eb.__obf_1726fa49ac928860 = 0
	__obf_1397b4a37b8546eb.__obf_ede102380a19106b = []any{}
	__obf_1397b4a37b8546eb.__obf_300af38d91fabd70 = nil
	__obf_1397b4a37b8546eb.__obf_62fa3336f332da73 = false
	__obf_1397b4a37b8546eb.__obf_15a31c61a999fdce = false
	return __obf_1397b4a37b8546eb
}

// 查询方法
func (__obf_1397b4a37b8546eb *Context) __obf_66a128019ff47f8a(__obf_02706219b5e5cf33 any, __obf_9683f87a2fc19cf8 int) (__obf_63349b4c89eea887 error) {
	defer __obf_1397b4a37b8546eb.__obf_9858352bdc58a6fc.__obf_5d8be8432bb33d38.Put(__obf_1397b4a37b8546eb)
	__obf_43f62dc6e0fe2c6f, __obf_49027310ec59880a := context.WithTimeout(context.Background(), __obf_1397b4a37b8546eb.__obf_9858352bdc58a6fc.__obf_d40bafd7b95c975a)
	defer __obf_49027310ec59880a()
	if __obf_1397b4a37b8546eb.sql == "" {
		__obf_1397b4a37b8546eb.
			sql = __obf_1397b4a37b8546eb.__obf_edaf9b1cb0d62126(__obf_02706219b5e5cf33)
	}
	switch __obf_9683f87a2fc19cf8 {
	case SelectTypeOne:
		if __obf_1397b4a37b8546eb.__obf_300af38d91fabd70 != nil {
			__obf_63349b4c89eea887 = __obf_1397b4a37b8546eb.__obf_300af38d91fabd70.GetContext(__obf_43f62dc6e0fe2c6f, __obf_02706219b5e5cf33, __obf_1397b4a37b8546eb.sql, __obf_1397b4a37b8546eb.__obf_ede102380a19106b...)
		} else {
			__obf_63349b4c89eea887 = __obf_1397b4a37b8546eb.__obf_9858352bdc58a6fc.GetContext(__obf_43f62dc6e0fe2c6f, __obf_02706219b5e5cf33, __obf_1397b4a37b8546eb.sql, __obf_1397b4a37b8546eb.__obf_ede102380a19106b...)
		}
	case SelectTypeMany:
		if __obf_1397b4a37b8546eb.__obf_300af38d91fabd70 != nil {
			__obf_63349b4c89eea887 = __obf_1397b4a37b8546eb.__obf_300af38d91fabd70.SelectContext(__obf_43f62dc6e0fe2c6f, __obf_02706219b5e5cf33, __obf_1397b4a37b8546eb.sql, __obf_1397b4a37b8546eb.__obf_ede102380a19106b...)
		} else {
			__obf_63349b4c89eea887 = __obf_1397b4a37b8546eb.__obf_9858352bdc58a6fc.SelectContext(__obf_43f62dc6e0fe2c6f, __obf_02706219b5e5cf33, __obf_1397b4a37b8546eb.sql, __obf_1397b4a37b8546eb.__obf_ede102380a19106b...)
		}
	default:
		panic("select type err")
	}
	return
}

// update,insert,delete方法
func (__obf_1397b4a37b8546eb *Context) __obf_31c51f7c54ad4f2d(__obf_8350072420eebe14 string, __obf_ede102380a19106b ...any) (sql.Result, error) {
	log.Printf("littleorm exec sql: <%s>, args: %#v", __obf_8350072420eebe14, __obf_ede102380a19106b)
	defer __obf_1397b4a37b8546eb.__obf_9858352bdc58a6fc.__obf_5d8be8432bb33d38.Put(__obf_1397b4a37b8546eb)
	__obf_43f62dc6e0fe2c6f, __obf_49027310ec59880a := context.WithTimeout(context.Background(), __obf_1397b4a37b8546eb.__obf_9858352bdc58a6fc.__obf_d40bafd7b95c975a)
	defer __obf_49027310ec59880a()

	var __obf_21b9d3de34eee523 sqlx.ExecerContext
	if __obf_1397b4a37b8546eb.__obf_300af38d91fabd70 != nil {
		__obf_21b9d3de34eee523 = __obf_1397b4a37b8546eb.__obf_300af38d91fabd70
	} else {
		__obf_21b9d3de34eee523 = __obf_1397b4a37b8546eb.__obf_9858352bdc58a6fc
	}
	return __obf_21b9d3de34eee523.ExecContext(__obf_43f62dc6e0fe2c6f, __obf_8350072420eebe14,

		// select查询语句的拼接
		__obf_ede102380a19106b...)
}

func (__obf_1397b4a37b8546eb *Context) __obf_edaf9b1cb0d62126(__obf_02706219b5e5cf33 any) string {
	var __obf_31e5142b7a1bc7dd []string
	__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, "select")
	if len(__obf_1397b4a37b8546eb.__obf_4939871f1153a0a4) != 0 {
		__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, __obf_8efd0bdb5b054c71(__obf_1397b4a37b8546eb.__obf_4939871f1153a0a4, SeqComma))
	} else {
		__obf_8ff238b4c2bb499f := // 如果不指定字段，取出目标对象的 tag 中的 db 全部填充了，
			// 不使用 * 来填充是因为 orm 解析时候如果对象中不包含数据库中全部字段会出现映射错误，会让以后增加数据库字段时候不兼容
			__obf_bf6eae1a72666267(__obf_02706219b5e5cf33)
		if len(__obf_8ff238b4c2bb499f) > 0 {
			__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, __obf_8efd0bdb5b054c71(__obf_8ff238b4c2bb499f, SeqComma))
		} else {
			__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, "*")
		}
	}
	__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, "from "+__obf_1397b4a37b8546eb.__obf_e6237083261f4fc2)
	if len(__obf_1397b4a37b8546eb.__obf_502636b0963e1e47) != 0 {
		__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, __obf_e8b95ee39de307d2(__obf_1397b4a37b8546eb.__obf_502636b0963e1e47, Grouping))
	}

	if __obf_1397b4a37b8546eb.__obf_b002972f2f4e9511 != "" {
		__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, "group by "+__obf_1397b4a37b8546eb.__obf_b002972f2f4e9511)
	}

	if __obf_1397b4a37b8546eb.__obf_fff491dab590b800 != "" {
		__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, "having "+__obf_1397b4a37b8546eb.__obf_fff491dab590b800)
	}

	if __obf_1397b4a37b8546eb.__obf_e82b41174897944b != "" {
		__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, "order by "+__obf_1397b4a37b8546eb.__obf_e82b41174897944b)
	}

	if __obf_1397b4a37b8546eb.__obf_d9b389e6a9615f76 != 0 {
		__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, fmt.Sprintf("limit %d, %d", __obf_1397b4a37b8546eb.__obf_1726fa49ac928860, __obf_1397b4a37b8546eb.__obf_d9b389e6a9615f76))
	}
	if __obf_1397b4a37b8546eb.__obf_62fa3336f332da73 {
		__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, "lock in share mode")
	}
	if __obf_1397b4a37b8546eb.__obf_15a31c61a999fdce {
		__obf_31e5142b7a1bc7dd = append(__obf_31e5142b7a1bc7dd, "for update")
	}
	sql := __obf_8efd0bdb5b054c71(__obf_31e5142b7a1bc7dd, SeqSpace)
	log.Printf("littleorm sql: <%v>, args: %#v", sql, __obf_1397b4a37b8546eb.__obf_ede102380a19106b)
	return sql
}

// /////////////////////////utils method/////////////////////////
// 拼接where条件
func __obf_e8b95ee39de307d2(__obf_502636b0963e1e47 []string, __obf_230d00ec78b1a178 string) string {
	if len(__obf_502636b0963e1e47) > 0 {
		return fmt.Sprintf("where %s", strings.Join(__obf_502636b0963e1e47, __obf_230d00ec78b1a178))
	} else {
		return ""
	}
}

// 拼接数组字符串
func __obf_8efd0bdb5b054c71(__obf_ede102380a19106b []string, __obf_801572388a12d075 string) string {
	return strings.Join(__obf_ede102380a19106b,

		// 解析对象中的 `db tag`
		// 参数只能指针，单个对象或者数组，eg: &little, &[]Little
		__obf_801572388a12d075)
}

func __obf_bf6eae1a72666267(__obf_02706219b5e5cf33 any) (__obf_deabc6b2d5f62f57 []string) {
	__obf_5dbd5d199912b6ea := reflect.ValueOf(__obf_02706219b5e5cf33)
	__obf_0a07f9a05d46fc0c := __obf_5dbd5d199912b6ea.Type().Elem()
	var __obf_3458801daaa16341 reflect.Type
	if __obf_0a07f9a05d46fc0c.Kind() == reflect.Slice {
		__obf_3458801daaa16341 = __obf_0a07f9a05d46fc0c.Elem()
	} else {
		__obf_3458801daaa16341 = __obf_0a07f9a05d46fc0c
	}
	for __obf_7c6c9e1512e37985 := 0; __obf_7c6c9e1512e37985 < __obf_3458801daaa16341.NumField(); __obf_7c6c9e1512e37985++ {
		__obf_8e4f70f32e1f241f := __obf_3458801daaa16341.Field(__obf_7c6c9e1512e37985).Tag.Get(DBTag)
		if __obf_8e4f70f32e1f241f != "" {
			__obf_deabc6b2d5f62f57 = append(__obf_deabc6b2d5f62f57, __obf_8e4f70f32e1f241f)
		}
	}
	return
}

func CleanSQLXMap(__obf_c253043d2fa709c1 map[string]any) map[string]any {
	for __obf_9b67b54491941e78, __obf_cf9e8016bdd91675 := range __obf_c253043d2fa709c1 {
		switch __obf_a73e9633161d85ba := __obf_cf9e8016bdd91675.(type) {
		case []byte:
			__obf_39f34c7e1dead46b := string(__obf_a73e9633161d85ba)
			__obf_cddd7715e53c449e := // 尝试智能识别具体类型
				strings.TrimSpace(__obf_39f34c7e1dead46b)

			// bool
			if strings.EqualFold(__obf_cddd7715e53c449e, "true") || strings.EqualFold(__obf_cddd7715e53c449e, "false") {
				if __obf_866808db5ec5296b, __obf_63349b4c89eea887 := strconv.ParseBool(__obf_cddd7715e53c449e); __obf_63349b4c89eea887 == nil {
					__obf_c253043d2fa709c1[__obf_9b67b54491941e78] = __obf_866808db5ec5296b
					continue
				}
			}

			// int
			if __obf_7c6c9e1512e37985, __obf_63349b4c89eea887 := strconv.Atoi(__obf_cddd7715e53c449e); __obf_63349b4c89eea887 == nil {
				__obf_c253043d2fa709c1[__obf_9b67b54491941e78] = __obf_7c6c9e1512e37985
				continue
			}

			// float
			if __obf_e242c9b073a9f4fe, __obf_63349b4c89eea887 := strconv.ParseFloat(__obf_cddd7715e53c449e, 64); __obf_63349b4c89eea887 == nil {
				__obf_c253043d2fa709c1[__obf_9b67b54491941e78] = __obf_e242c9b073a9f4fe
				continue
			}

			// JSON 对象或数组
			if (strings.HasPrefix(__obf_cddd7715e53c449e, "{") && strings.HasSuffix(__obf_cddd7715e53c449e, "}")) ||
				(strings.HasPrefix(__obf_cddd7715e53c449e, "[") && strings.HasSuffix(__obf_cddd7715e53c449e, "]")) {
				var __obf_d1ba51a229daf597 any
				if __obf_63349b4c89eea887 := json.Unmarshal(__obf_a73e9633161d85ba, &__obf_d1ba51a229daf597); __obf_63349b4c89eea887 == nil {
					__obf_c253043d2fa709c1[__obf_9b67b54491941e78] = __obf_d1ba51a229daf597
					continue
				}
			}
			__obf_c253043d2fa709c1[ // 默认保留为 string
			__obf_9b67b54491941e78] = __obf_39f34c7e1dead46b
		default:
			__obf_c253043d2fa709c1[__obf_9b67b54491941e78] = __obf_a73e9633161d85ba
		}
	}
	return __obf_c253043d2fa709c1
}

func ToListMap(__obf_97c7de7ee94125bc *sqlx.Rows) []map[string]any {
	var __obf_1576eeccdf72c56c []map[string]any
	for __obf_97c7de7ee94125bc.Next() {
		__obf_44e4bd3025765b4d := make(map[string]any)
		if __obf_63349b4c89eea887 := __obf_97c7de7ee94125bc.MapScan(__obf_44e4bd3025765b4d); __obf_63349b4c89eea887 != nil {
			slog.Error(fmt.Sprintf("[ToListMap] MapScan: %s", __obf_63349b4c89eea887.Error()))
			continue
		}
		__obf_1576eeccdf72c56c = append(__obf_1576eeccdf72c56c, CleanSQLXMap(__obf_44e4bd3025765b4d))
	}
	return __obf_1576eeccdf72c56c
}
