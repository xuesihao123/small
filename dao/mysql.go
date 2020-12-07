package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql")

var (
	Db *gorm.DB
)

func InitMysql()(err error){
	Db , err = gorm.Open("mysql", "root:123456@(127.0.0.1:3307)/small?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		return
	}
	return Db.DB().Ping()
}
