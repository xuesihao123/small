package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"small/dao"
	"small/models"
	"small/service"
	"strings"
	"time"
)

var ok gin.H = gin.H{
	"status":true,
	"err":nil,
	"code":200,
	"data":nil,
}

func FindUser(c *gin.Context)  {
	err,users:= models.FindU()
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"查询失败",
			"code":223,
			"data":nil,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status":true,
			"err":nil,
			"code":200,
			"data":users,
		})
	}
}
func CreateUser(c *gin.Context){
	var user models.User
	c.ShouldBindJSON(&user)
	if models.GetName(user.UserName){
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"code": 222,
			"err":"用户名重复",
			"data":nil,
		})
	}
	user.UserPassword = service.Md5Password(user.UserPassword)
	err := models.CreateU(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"code":223,
			"err":"创建失败",
			"data":nil,
		})
	}else{
		c.JSON(http.StatusOK,ok)
	}
}
func UpdateUser(c *gin.Context){//这里不提供修改密码
	var user models.User
	c.ShouldBindJSON(&user)
	err,_:= models.GetU(user.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest,err)
	}else {
		err = models.UpdateU(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":false,
				"err":"更新失败",
				"code":223,
				"data":nil,
			})
		}else{
			c.JSON(http.StatusOK,ok)
		}
	}
}

func UpdateUserPassword(c *gin.Context){//提供修改密码
	var user models.User
	c.ShouldBindJSON(&user)
	err,_:= models.GetU(user.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"code":224,
			"err":"不存在这个用户",
			"data":nil,
		})
	}else {
		user.UserPassword = service.Md5Password(user.UserPassword)
		err = models.UpdateU(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":false,
				"code":223,
				"err":err,
				"data":nil,
			})
		}else{
			c.JSON(http.StatusOK,ok)
		}
	}
}

func DeleteUser(c *gin.Context){
	var user models.User
	c.ShouldBindJSON(&user)
	err := models.DeleteU(user.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"code":223,
			"err":"删除失败",
			"data":nil,
		})
	}else{
		c.JSON(http.StatusOK,ok)
	}
}
func IndexUser(c *gin.Context){
	var user models.User
	c.ShouldBindJSON(&user)
	user.UserPassword = service.Md5Password(user.UserPassword)
	err,u:= models.GetNameAndPassword(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"code":201,
			"err":"账号或密码错误",
			"data":nil,
		})
	}else{
		tokenString, _ := service.GenToken(*u)
		c.JSON(http.StatusOK,gin.H{
			"token":tokenString,
			"status":true,
			"code":200,
			"err":nil,
			"ok":u.UserStatus,
		})
	}
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := service.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Set("userid",mc.ID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func SetMail(c *gin.Context){
	var mail []string
	var  m  service.Mail
	var M models.Mail
	c.ShouldBindJSON(&m)
	mail = append(mail,m.AMail)
	rind := models.Rind()
	My,err := models.GetEmail(m.AMail)
	//c.JSON(http.StatusOK,My)
	if err != nil {
		c.JSON(http.StatusBadRequest,err)
	}
	M.MailUserId = My.UserId
	//c.Set("userId",My.UserId)
	M.MailMail = m.AMail
	M.MailRind = rind
	M.MailTime = time.Now().Unix()+600

	err = models.CreateMail(&M)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"发送失败",
			"code":223,
			"data":nil,
		})
	}else {
		//c.JSON(http.StatusOK,mail)
		err := service.SetAMail(mail, rind)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":false,
				"err":"发送失败",
				"code":223,
				"data":nil,
			})
		} else {
			c.JSON(http.StatusOK, ok)
		}
	}
}

func TrueRind(c *gin.Context){
	var mail models.Mail
	c.ShouldBindJSON(&mail)
	My,err := models.GetEmail(mail.MailMail)
	mail.MailUserId = My.UserId
	_,err = models.SetIdMail(&mail)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"验证错误",
			"code":223,
			"data":nil,
		})
	}else{
		c.JSON(http.StatusOK,ok)
	}
}

func RePassword(c *gin.Context){//提供修改密码
	var user models.User
	c.ShouldBindJSON(&user)
	u,err:= models.GetEmail(user.UserEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"邮箱不存在",
			"code":224,
			"data":nil,
		})
	}else {
		u.UserPassword = service.Md5Password(user.UserPassword)
		err = models.UpdateU(u)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":false,
				"err":"修改失败",
				"code":223,
				"data":nil,
			})
		}else{
			c.JSON(http.StatusOK,ok)
		}
	}
}

func SetMailRegister(c *gin.Context){
	var mail []string
	var  m  service.Mail
	var M models.Mail
	c.ShouldBindJSON(&m)
	mail = append(mail,m.AMail)
	rind := models.Rind()
	My,err := models.GetEmail(m.AMail)
	if My != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"已经被注册了",
			"code":201,
			"data":nil,
		})
	}
	//c.Set("userId",My.UserId)
	M.MailMail = m.AMail
	M.MailRind = rind
	M.MailTime = time.Now().Unix()+600

	err = models.CreateMail(&M)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"失败",
			"code":223,
			"data":nil,
		})
	}else {
		//c.JSON(http.StatusOK,mail)
		e  := service.SetAMail(mail, rind)
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":false,
				"err":"失败",
				"code":201,
				"data":nil,
			})
		} else {
			c.JSON(http.StatusOK, ok)
		}
	}
}

func TrueRindRegister(c *gin.Context){
	var mail models.Mail
	c.ShouldBindJSON(&mail)
	_,err := models.SetIdMail(&mail)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"验证失败",
			"code":223,
			"data":nil,
		})
	}else{
		c.JSON(http.StatusOK,ok)
	}
}

func PostFile(c *gin.Context) {
	//S := sessions.Default(c)

	file, err := c.FormFile("imgfile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"err":"失败",
			"code":223,
			"data":nil,
		})
		return
	}
	//获取文件名
	fileName := file.Filename
	filet := "./pic/" + fileName
	dst := "47.99.181.139:9099/"+"pic/" + fileName
	dao.RDb.Set("filename",dst,0)
	//S.Set("filename",dst)
	//S.Save()
	//fmt.Println("文件名：", fileName)
	//保存文件到服务器本地
	//SaveUploadedFile(文件头，保存路径)q
	if err := c.SaveUploadedFile(file, filet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":false,
			"err":"保存失败",
			"code":223,
			"data":nil,
		})
		return
	}
	c.JSON(http.StatusOK, ok)
}
