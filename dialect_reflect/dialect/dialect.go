package dialect

import "reflect"

//管理所有注册了dialect映射的数据库
var dialectMap = map[string]Dialect{}

type Dialect interface {
	//通过go得到sql的类型
	DataTypeOf(typ reflect.Value) string
	//判断某个表是否存在的sql语句
	TableExistSQL(tableName string) (string, []interface{})
}

//如果增加了对某个数据库的支持，需要注册进入Map
func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

//得到某个数据库的dialect
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
