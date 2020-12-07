package models

type Goods struct {
	GoodId int `gorm:"primary_key" json:"good_id"`
	UserId int `gorm:"index" json:"user_id"`
	HouseId int `gorm:"index" json:"house_id"`
	GoodName string `gorm:"type:varchar(20)" json:"good_name"`
	GoodMessage string `gorm:"type:varchar(500)" json:"good_message"`
	GoodNumber int `gorm:"type:int" json:"good_number"`
	GoodStatus string `gorm:"type:varchar(11)" json:"good_status"`
	GoodPic string `gorm:"type:mediumtext" json:"good_pic"`
	GoodType string `gorm:"type:varchar(20)" json:"good_type"`
}

type HouseGood struct {
	HouseName string
	GoodId int `gorm:"primary_key" json:"good_id"`
	UserId int `gorm:"index" json:"user_id"`
	HouseId int `gorm:"index" json:"house_id"`
	GoodName string `gorm:"type:varchar(20)" json:"good_name"`
	GoodMessage string `gorm:"type:varchar(500)" json:"good_message"`
	GoodNumber int `gorm:"type:int" json:"good_number"`
	GoodStatus string `gorm:"type:int" json:"good_status"`
	GoodPic string `gorm:"type:mediumtext" json:"good_pic"`
	GoodType string `gorm:"type:varchar(20)" json:"good_type"`
}