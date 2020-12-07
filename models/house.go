package models

import "small/dao"

// re的名 获取帖子列表query string参数
func CreateH(house *House)(err error)  {
	if err = dao.Db.Create(house).Error;err != nil {
		return err
	}
	return
}

func FindH()(err error,houses []*House){
	if err = dao.Db.Find(&houses).Error;err!=nil{
		return err,nil
	}
	return err,houses
}

func GetIdH(ID int)(err error,house *House)  {
	house = new(House)
	if err = dao.Db.Where("house_id = ?",ID).First(house).Error;err != nil {
		return err,nil
	}
	return err,house
}

func UpdateH(house *House)(err error)  {
	if err = dao.Db.Save(house).Where("house_id = ?" , house.HouseId).Error;err != nil {
		return err
	}
	return
}

func DeleteH(ID int)(err error)  {
	if err = dao.Db.Where("house_id = ?",ID).Delete(House{}).Error;err != nil {
		return err
	}
	return 	
}
func SelectName(name string)(houses []*House,err error)  {
	err = dao.Db.Where("house_name like %s",name).Find(&houses).Error
	if err != nil {
		return nil,err
	}
	return houses,err
}