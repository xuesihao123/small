package models

import "small/dao"


func CreateU(user *User)(err error){
	if err := dao.Db.Create(user).Error;err!=nil{
		return err
	}
	return
}

func FindU()(err error, users []*User){
	if err = dao.Db.Find(&users).Error;err!=nil{
		return err,nil
	}
	return err,users
}

func GetU(ID int)(err error,user *User){
	user = new(User)
	if err = dao.Db.Where("user_id = ?",ID).First(user).Error;err != nil{
		return err,nil
	}
	return err,user
}

func UpdateU(user *User)(err error){
	if err = dao.Db.Save(user).Where("user_id = ?",user.UserId).Error;err != nil{
		return err
	}
	return
}

func DeleteU(ID int)(err error){
	 if err = dao.Db.Where("user_id = ?",ID).Delete(User{}).Error;err!=nil{
	 	return err
	}
	return
}

func GetNameAndPassword(user *User)(error,*User){
	//上面的Password值要在logic中重新用md5加密
	//var u *User
	if err := dao.Db.Debug().Where("user_name = ? AND user_password = ?",user.UserName,user.UserPassword).First(user).Error;err!=nil{
		return err,nil
	}
	return nil,user
}

func GetEmail(email string)(u *User,err error){
	//上面的Password值要在logic中重新用md5加密
	u = new(User)
	if err = dao.Db.Debug().Where("user_email = ? ",email).First(u).Error;err!=nil{
		return nil,err
	}
	return u,err
}

func GetName(UserName string)bool{
	//上面的Password值要在logic中重新用md5加密
	//var u *User
	if err := dao.Db.Debug().Where("user_name = ? ",UserName).First(&User{}).Error;err!=nil{
		return false
	}
	return true
}