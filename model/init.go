package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(connstring string) {
	fmt.Println("connstring", connstring)
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		fmt.Println(err)
		panic("Mysql数据库连接错误")
	}
	fmt.Println("success")
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)                       //表名不加s，gorm默认是在表名后加s
	db.DB().SetMaxIdleConns(20)                  //设置连接池
	db.DB().SetMaxOpenConns(100)                 //设置最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) //连接时间
	DB = db
	migration()
}
