package models

type House struct {
	HouseId int `gorm:"primary_key" json:"house_id"`
	UserId int	`gorm:"type:int" json:"user_id"`
	HouseName string `gorm:"type:varchar(20)" json:"house_name"`
	HouseSize string `gorm:"type:int" json:"house_size"`
	HousePrice string `gorm:"type:float" json:"house_price"`
	HousePic string `gorm:"type:mediumtext" json:"house_pic"`
	HouseStatus int `gorm:"type:int" json:"house_status"`
}
