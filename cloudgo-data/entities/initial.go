package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var mydb *xorm.Engine

func init() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	mydb = engine
	/*
		if err = mydb.Sync2(new(UserInfo)); err != nil {
			panic(err)
		}
	*/
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
