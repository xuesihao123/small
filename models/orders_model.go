package models

import "small/dao"

type Order struct {
	OrderId int `gorm:"primary_key" json:"order_id"`
	UserId int `gorm:"index" json:"user_id"`
	HouseId int `gorm:"index" json:"house_id"`
	GoodID string `gorm:"type:int" json:"good_id"`
	OrderNum int `gorm:"type:int " json:"order_num"`
	OrderPrice float64 `gorm:"type:float " json:"'order_price'"`
	OrderCreateTime string `gorm:"type : date" json:"order_create_time"`
	OrderStatus int `gorm:"type:int " json:"order_status"`
	OrderTime string `gorm:"type : date" json:"order_time"`
}

type OrderR struct {
	OrderId int `json:"order_id"`
	UserId int `json:"user_id"`
	UserName string `json:"user_name"`
	HouseId int `json:"house_id"`
	HouseName string `json:"house_name"`
	GoodID string `json:"good_id"`
	GoodName string `json:"good_name"`
	OrderNum int `json:"order_num"`
	OrderPrice float64 `json:"order_price"`
	OrderCreateTime string `json:"order_create_time"`
	OrderStatus int `json:"order_status"`
	OrderTime string `json:"order_time"`
}

func a()  {
	dao.Db.Table("orders").Select("orders.order_id,orders.user_id,users.user_name,orders.house_id,houses.house_name,orders.good_id,goods.good_name,orders.order_num,orders.order_price,orders.order_create_time,orders.order_status,orders.order_time").Joins("left join users on orders.user_id = users.user_id").Joins("left join goods on orders.good_id = goods.good_id").Joins("left join houses on orders.house_id = houses.house_id")
}