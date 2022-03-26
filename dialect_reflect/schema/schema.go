package schema

import (
	"go/ast"
	"reflect"
	"somnusorm/dialect"
)

//字段名 Name、类型 Type、和约束条件 Tag
//表中的一行
type Field struct {
	Name string
	Type string
	Tag  string
}

//被映射的对象 Model、表名 Name 和字段 Fields
type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetFiled(name string) *Field {
	return schema.fieldMap[name]
}

//将任意的对象解析为dialect
func Parse(dest interface{}, d dialect.Dialect) (schema *Schema) {
	//Indirect指向指针所对应的实例
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema = &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("somnusorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
