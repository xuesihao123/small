package models

import (
	"fmt"
	"math/rand"
	"small/dao"
	"time"
)


func CreateMail(mail *Mail)(err error)  {
	//SetIdMail(mail)
	if err = dao.Db.Create(mail).Error;err != nil {
		return err
	}
	return
}

func SetIdMail(mail *Mail)(m *Mail,err error){

	if err = dao.Db.Debug().Where("mail_user_id = ? ", mail.MailUserId).First(mail).Error;err != nil {
		return nil,err
	}
		DeleteMail(mail)
		return mail,err
}

func DeleteMail(mail *Mail)(err error){
	if err = dao.Db.Where("mail_id = ?",mail.MailID).Delete(Mail{}).Error;err != nil {
		return err
	}
	return
}

func Rind()string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}