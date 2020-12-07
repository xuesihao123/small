package models

import "small/dao"

func CreateO(order *Order)(err error)  {
	if err = dao.Db.Create(order).Error;err != nil {
		return err
	}
	return
}

func GetIdO(ID int)(err error,order *Order)  {
	order = new(Order)
	if err = dao.Db.Where("order_id = ?",ID).First(order).Error;err != nil {
		return err,nil
	}
	return err,order
}

func UpdateO(order *Order)(err error)  {
	if err = dao.Db.Save(order).Where("order_id = ?" , order.OrderId).Error;err != nil {
		return err
	}
	return
}

func UpdateOBuy(OrderId int)(err error)  {
	if err = dao.Db.Model(Order{}).Where("order_id = ?" , OrderId).Update("order_status = ?", 2).Error;err != nil {
		return err
	}
	return
}

func UpdateOBuyALL(OrderIdALL []int)(err error)  {
	tx := dao.Db.Begin()
	err = tx.Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Where("order_id in (?)" ,OrderIdALL).Update("order_status = ?",2).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func DeleteO(ID int)(err error)  {
	if err = dao.Db.Where("order_id = ?",ID).Delete(Order{}).Error;err != nil {
		return err
	}
	return
}
func OrderUserId(UserId int)(err error,orders []*Order)  {
	if err = dao.Db.Table("orders").Select("orders.order_id,orders.user_id,users.user_name,orders.house_id,houses.house_name,orders.good_id,goods.good_name,orders.order_num,orders.order_price,orders.order_create_time,orders.order_status,orders.order_time").Joins("left join users on orders.user_id = users.user_id").Joins("left join goods on orders.good_id = goods.good_id").Joins("left join houses on orders.house_id = houses.house_id").Where("user_id = ?",UserId).Find(&orders).Error;err!=nil{
		return err,nil
	}
	return err,orders
}

func OrderHouseId(HouseId int)(err error,orders []*OrderR)  {
	if err = dao.Db.Table("orders").Select("orders.order_id,orders.user_id,users.user_name,orders.house_id,houses.house_name,orders.good_id,goods.good_name,orders.order_num,orders.order_price,orders.order_create_time,orders.order_status,orders.order_time").Joins("left join users on orders.user_id = users.user_id").Joins("left join goods on orders.good_id = goods.good_id").Joins("left join houses on orders.house_id = houses.house_id").Where("house_id",HouseId).Find(&orders).Error;err!=nil{
		return err,nil
	}
	return err,orders
}

func GetIdAndStatusO(ID int,status int)(err error,order *OrderR)  {
	order = new(OrderR)
	if err = dao.Db.Table("orders").Select("orders.order_id,orders.user_id,users.user_name,orders.house_id,houses.house_name,orders.good_id,goods.good_name,orders.order_num,orders.order_price,orders.order_create_time,orders.order_status,orders.order_time").Joins("left join users on orders.user_id = users.user_id").Joins("left join goods on orders.good_id = goods.good_id").Joins("left join houses on orders.house_id = houses.house_id").Where("user_id = ? AND order_status = ?",ID,status).First(order).Error;err != nil {
		return err,nil
	}
	return err,order
}