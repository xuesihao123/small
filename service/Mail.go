package service

import (
	"gopkg.in/gomail.v2"
)


// MailboxConf 邮箱配置
type MailboxConf struct {
	// 邮件标题
	Title string
	// 邮件内容
	Body string
	// 收件人列表
	RecipientList []string
	// 发件人账号
	Sender string
	// 发件人密码，QQ邮箱这里配置授权码
	SPassword string
	// SMTP 服务器地址， QQ邮箱是smtp.qq.com
	SMTPAddr string
	// SMTP端口 QQ邮箱是25
	SMTPPort int
}

type Mail struct {
	AMail string `json:"user_email"`
}

func SetAMail(mail []string,rind string) (err error) {
	var mailConf MailboxConf
	mailConf.Title = "小仓存储验证码"
	mailConf.Body = "<html>亲爱的用户：您好！<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;您收到这封这封电子邮件是给您发送一个验证码。假如这不是您本人所申请, 请不用理会这封电子邮件, 但是如果您持续收到这类的信件骚扰, 请您尽快联络管理员。<br><br>验证码为"+ rind +" &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;注意:请您在收到邮件10分钟内使用，否则该验证码将会失效。<br><br><center> &copy;  小仓存储服务中心</center></html>"
	mailConf.RecipientList = mail
	mailConf.Sender = `754769243@qq.com`
	mailConf.SPassword = "lgscdkuwsznibdgi"
	mailConf.SMTPAddr = `smtp.qq.com`
	mailConf.SMTPPort = 465

	m := gomail.NewMessage()
	m.SetHeader(`From`, mailConf.Sender)
	m.SetHeader(`To`, mailConf.RecipientList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, mailConf.Body)
	m.Attach("./Dockerfile")   //添加附件
	err = gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
	//err := d.DialAndSend(m)
	//log.Printf("Send Email Success")
	return nil
}
