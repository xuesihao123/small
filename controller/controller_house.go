package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"small/dao"
	"small/models"
	"small/service"
	"github.com/gin-contrib/sessions"
)


func FindHouse(c *gin.Context)  {
	err,houses := models.FindH()
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":err,
			"data":nil,
			"code":223,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status":true,
			"data":houses,
			"err":nil,
			"code":200,
		})
	}
}
func CreateHouse(c *gin.Context){
	S := sessions.Default(c)

	var house models.House
	c.ShouldBindJSON(&house)
	house.UserId = c.GetInt("userid")
	house.HousePic = S.Get("filename").(string)
	S.Set("filename","")
	S.Save()
	err := models.CreateH(&house)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"创建失败",
			"code":223,
			"data":nil,
		})
	}else {
		c.JSON(http.StatusOK,ok)
	}
}
func UpdateHouse(c *gin.Context){
	//S := sessions.Default(c)
	var house models.House
	c.ShouldBindJSON(&house)
	err,_:= models.GetIdH(house.HouseId)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"仓库不存在",
			"code":224,
			"data":nil,
		})
	}else {
		house.HousePic ,err = dao.RDb.Get("filename").Result()
		if err != nil {
			house.HousePic = ""
		}
		dao.RDb.Set("filename","",0)

		err = models.UpdateH(&house)
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
func DeleteHouse(c *gin.Context){
	var house models.House
	c.ShouldBindJSON(&house)
	err := models.DeleteH(house.HouseId)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"删除失败",
			"code":223,
			"data":nil,
		})
	}else {
		c.JSON(http.StatusOK,ok)
	}
}

//func SelectHouseIdGoods(c *gin.Context){
//	var house models.House
//	c.ShouldBindJSON(&house)
//	err,goods := models.GetHouserId(house.HouseId)
//	if err != nil {
//		c.JSON(http.StatusBadRequest,err)
//	}else{
//		c.JSON(http.StatusOK,goods)
//	}
//}

func SelectNameHouse(c *gin.Context){//模糊查询
	var house models.House
	c.ShouldBindJSON(&house)
	strsql := service.ChinaWords(house.HouseName)
	h,err := models.SelectName(strsql)
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
			"data":h,
			"code":200,
			"err":nil,
		})
	}
}