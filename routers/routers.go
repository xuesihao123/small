package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"small/controller"
	_ "small/docs"
	"small/service"
)

// @title //项目名
// @version 1.0
// @description //项目简介

// @contact.name
// @contact.url

// @license.name //服务器名

// @host //主页url
// @BasePath
func SetupRouters()*gin.Engine  {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))

	r.StaticFS("/pic",http.Dir("./pic"))
	//路由上加入session中间件
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(sessions.Sessions("mysession", store))

	r.Use(service.Cors())
	r.POST("/mailregister",controller.SetMailRegister)
	r.POST("/trueregister",controller.TrueRindRegister)
	r.POST("/index",controller.IndexUser)//登录
	r.POST("/register",controller.CreateUser)//注册
	r.POST("/Mail",controller.SetMail)//发送邮件
	r.POST("/true",controller.TrueRind)//验证验证码
	r.POST("/repassword",controller.RePassword)//修改密码
	//路由组
	r.POST("/pic",controller.PostFile)

	userGroup :=r.Group("small")//括号中的路径是/small/...
	userGroup.Use(controller.JWTAuthMiddleware())
	{
		//用户
		userGroup.GET("/user",controller.FindUser)//返回所有用户
		userGroup.POST("/user",)
		userGroup.PUT("/user",controller.UpdateUser)//更改一个用户信息,但不能修改密码
		userGroup.PUT("/password",controller.UpdateUserPassword)//修改密码
		userGroup.POST("/user/delete",controller.DeleteUser)//删除一个用户
		//仓库
		userGroup.GET("/house",controller.FindHouse)//返回所有仓库
		userGroup.POST("/house",controller.CreateHouse)//添加一个仓库
		userGroup.PUT("/house",controller.UpdateHouse)//修改一个仓库
		userGroup.POST("/house/delete",controller.DeleteHouse)//删除一个仓库
		userGroup.POST("/houseId",controller.FindGoodsHouserId)//用仓库ID返回所有商品
		//物品
		userGroup.GET("/goods",controller.FindGoods)//返回所有商品
		userGroup.POST("/goods",controller.CreateGoods)//创建一个新的商品
		userGroup.PUT("/goods",controller.UpdateGoods)//修改
		userGroup.POST("/goods/delete",controller.DeleteGoods)//删除
		userGroup.GET("/gooduser",controller.FindGoodsUserId)
		userGroup.PUT("/goodsNumber",controller.UpdateGoodsNumber)
		userGroup.POST("/pic",controller.PostFile)
	}


	return r
}
