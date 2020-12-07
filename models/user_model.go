package models

type User struct {
	UserId int `gorm:"primary_key" json:"user_id"`
	UserName string	`gorm:"type:varchar(20)" json:"user_name"`
	UserPassword string `gorm:"type:varchar(50)" json:"user_password"`
	UserEmail string	`gorm:"type:varchar(20);unique_index" json:"user_email"`
	UserStatus int `gorm:"type:int;default:1" json:"user_status"`
}