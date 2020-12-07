package main

import (
	"small/dao"
	"small/models"
	"small/routers"
)

func main() {
	err := dao.InitMysql()
	if err != nil {
		return
	}

	//err = dao.InitRedis() //初始化redis，使用redis取代session
	if err != nil {
		return
	}
	defer dao.Db.Close()
	//映射结构体
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.User{})
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.House{})
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.Goods{})
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.Mail{})
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.Order{})

	r := routers.SetupRouters()
	r.Run(":9099")
}