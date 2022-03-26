package schema

import (
	"fmt"
	"somnusorm/dialect"
	"testing"
)

type User struct {
	Name string `somnusorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	fmt.Println(schema.Name)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	fmt.Println(schema.GetFiled("Name").Name)
	if schema.GetFiled("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
