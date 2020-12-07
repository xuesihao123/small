package models

type Mail struct {
	MailID int `json:"mail_id"`
	MailTime int64 `gorm:"type:long" json:"mail_time"`
	MailRind string	`json:"mail_rind"`
	MailUserId int `json:"mail_user_id"`
	MailMail string `json:"mail_mail"`
}

