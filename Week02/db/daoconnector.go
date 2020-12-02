package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	//TODO 这里需要替换自己的mysql连接
	db, err := gorm.Open("mysql", "xxxxxxxx")
	if err != nil {
		panic(err)
	}

	Db = db

}
