package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"small/dao"
	"small/models"
)

// @Summary 查询所有物品
// @Description 接口详细描述信息
// @Tags 物品信息
// @accept json
// @Produce  json
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {re} Res {"status":200,"data":null,"msg":""}
// @Router /test/{id} [get]
func FindGoods(c *gin.Context)  {
	err,goods:= models.FindG()
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
			"data":goods,
			"code":200,
			"err":nil,
		})
	}
}
func CreateGoods(c *gin.Context){
	//S := sessions.Default(c)

	var good models.Goods
	c.ShouldBindJSON(&good)
	good.UserId = c.GetInt("userid")
	good.GoodPic,_ = dao.RDb.Get("filename").Result()

	dao.RDb.Set("filename","",0)
	err := models.CreateG(&good)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"创建失败",
			"code":223,
			"data":nil,
		})
	}else{
		c.JSON(http.StatusOK,ok)
	}
}
func UpdateGoods(c *gin.Context){
	S := sessions.Default(c)

	var good models.Goods
	c.ShouldBindJSON(&good)
	err,_:= models.GetIdG(good.GoodId)
	good.GoodPic = S.Get("filename").(string)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"商品不存在",
			"code":224,
			"data":nil,
		})
	}else {
		err = models.UpdateG(&good)
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
func DeleteGoods(c *gin.Context){
	var good models.Goods
	c.ShouldBindJSON(&good)
	err := models.DeleteG(good.GoodId)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":"删除失败",
			"code":223,
			"data":nil,
		})
	}else{
		c.JSON(http.StatusOK,ok)
	}
}

func FindGoodsHouserId(c *gin.Context){
	var house models.House
	c.ShouldBindJSON(&house)
	err , goods:=models.GetHouserId(house.HouseId)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":err,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status":true,
			"data":goods,
			"code":200,
			"err":nil,
		})
	}
}
func FindGoodsUserId(c *gin.Context){
	UserId := c.GetInt("userid")
	err , goods:=models.GetUserId(UserId)
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
			"data":goods,
			"code":200,
			"err":nil,
		})
	}
}

func UpdateGoodsNumber(c *gin.Context){

	var good models.Goods
	c.ShouldBindJSON(&good)
	err,_:= models.GetIdG(good.GoodId)
	if good.GoodNumber <= 0{
		if err1 := models.DeleteG(good.GoodId);err1 != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":false,
				"err":"商品不存在",
				"code":224,
				"data":nil,
			})
		}else {
			c.JSON(http.StatusOK,ok)
		}
	}else{
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":false,
				"err":"修改失败",
				"code":223,
				"data":nil,
			})
		}else {
			err = models.UpdateG(&good)
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
}