package models

import "small/dao"


func CreateG(good *Goods)(err error){
	if err = dao.Db.Create(good).Error ; err!=nil{
		return err
	}
	return
}

func FindG()(err error,goods []*HouseGood)  {
	//if err = dao.Db.Debug().Table("goods").Select("houses.house_name as HouseName,goods.good_id as GoodId,goods.user_id as UserId,goods.house_id as HouseId,goods.good_name as GoodName,goods.good_message as GoodMessage,goods.good_number as GoodNumber,goods.good_status as GoodStatus,goods.good_pic as GoodPic,goods.good_type as GoodType").Joins("left join houses on houses.house_id = goods.house_id").Find(&goods).Error ; err!=nil{
	if err = dao.Db.Debug().Table("goods").Select("houses.house_name,goods.good_id , goods.user_id,goods.house_id ,goods.good_name ,goods.good_message ,goods.good_number ,goods.good_status ,goods.good_pic ,goods.good_type").Joins("left join houses on houses.house_id = goods.house_id").Find(&goods).Error ; err!=nil{
		return err,nil
	}
	return  err,goods
}

func GetIdG(ID int)(err error,good *Goods){
	good = new(Goods)
	if err = dao.Db.Where("good_id = ?",ID).First(good).Error ; err!=nil{
		return err,nil
	}
	return err,good
}

func UpdateG(good *Goods)(err error){
	if err = dao.Db.Save(good).Where("good_id = ?", good.GoodId).Error;err!=nil{
		return err
	}
	return
}

func DeleteG(ID int)(err error){
	if err = dao.Db.Where("good_id = ?",ID).Delete(Goods{}).Error;err!=nil{
		return err
	}
	return
}

func GetHouserId(ID int)(err error,goods []*HouseGood){
	err = dao.Db.Table("goods").Select("houses.house_name,goods.good_id , goods.user_id,goods.house_id ,goods.good_name ,goods.good_message ,goods.good_number ,goods.good_status ,goods.good_pic ,goods.good_type").Joins("left join houses on houses.house_id = goods.house_id").Where("house.house_id = ?",ID).Find(&goods).Error
	if err != nil {
		return err,nil
	}
	return
}

func GetUserId(ID int)(err error,goods []*HouseGood){
	err = dao.Db.Table("goods").Select("houses.house_name,goods.good_id , goods.user_id,goods.house_id ,goods.good_name ,goods.good_message ,goods.good_number ,goods.good_status ,goods.good_pic ,goods.good_type").Joins("left join houses on houses.house_id = goods.house_id").Where("goods.user_id = ?",ID).Find(&goods).Error
	if err != nil {
		return err,nil
	}
	return
}
