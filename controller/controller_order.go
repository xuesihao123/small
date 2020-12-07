package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"small/models"
	"strconv"
)


// @Summary 按照用户id查询用户所有的订单
// @Description 按照用户id查询用户所有的订单
// @Tags 订单
// @accept json
// @Produce  json
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /user_order [get]
func SelectUserOrder(c *gin.Context)  {
	ID := c.GetInt("userid")
	err , orders := models.OrderUserId(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err" : "查询失败",
			"data" : nil,
			"code" : 223,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":true,
		"data":orders,
		"err":nil,
		"code":200,
	})
}

// @Summary 按照仓库id查询仓库所有的订单
// @Description 按照仓库id查询仓库所有的订单
// @Tags 订单
// @accept json
// @Produce  json
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /house_order/{house_id} [get]
func SelectHouseOrder(c *gin.Context)  {
	ID := c.Param("house_id")
	Id , _ := strconv.Atoi(ID)
	err , orders := models.OrderHouseId(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err" : "查询失败",
			"data" : nil,
			"code":223,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":true,
		"data":orders,
		"err":nil,
		"code":200,
	})
}



// @Summary 按照用户id和订单状态查询订单
// @Description 按照用户id和订单状态查询订单
// @Tags 订单
// @accept json
// @Produce  json
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /user_order/{status} [get]
func SelectUserStatusOrder(c *gin.Context)  {
	ID := c.GetInt("userid")
	Status := c.Param("status")
	status ,_:= strconv.Atoi(Status)
	err , orders := models.GetIdAndStatusO(ID , status)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err" :"查询失败",
			"data" : nil,
			"code":223,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":true,
		"data":orders,
		"err":nil,
		"code":200,
	})
}

// @Summary 按照用户id和订单状态查询订单
// @Description 按照用户id和订单状态查询订单
// @Tags 订单
// @accept json
// @Produce  json
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /orderdelete/{orderId} [post]
func SelectDeleteOrder(c *gin.Context)  {
	ID := c.Param("orderId")
	Id , _ := strconv.Atoi(ID)
	err := models.DeleteO(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err" : "查询失败",
			"data" : nil,
			"code":223,
		})
		return
	}
	c.JSON(http.StatusOK,ok)
}


// @Summary 支付完成后修改订单的状态
// @Description 支付完成后修改订单的状态
// @Tags 订单
// @accept json
// @Produce  json
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /orderdelete/{orderId} [post]
func UpdateNoBuyOrder(c *gin.Context)  {
	var order models.Order
	c.ShouldBindJSON(&order)
	err := models.UpdateO(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err" : "修改失败",
			"data" : nil,
			"code" : 223,
		})
		return
	}
	er,UpdateOrder := models.GetIdO(order.OrderId)
	if er != nil {
		c.JSON(http.StatusNotImplemented,gin.H{
			"status":false,
			"err" : "请重新获取数据",
			"data" : nil,
			"code":225,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":true,
		"data": UpdateOrder,
		"err": nil,
		"code":200,
	})
}

// @Summary 批量处理未支付订单
// @Description 批量处理未支付订单
// @Tags 订单
// @accept json
// @Produce  json
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /orderdelete/{orderId} [post]
func UpdateBuyOrder(c *gin.Context)  {
	ID := c.Param("orderId")
	Id , _ := strconv.Atoi(ID)
	err := models.UpdateOBuy(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"code":223,
			"err" : "更新失败",
			"data" : nil,
		})
		return
	}
	c.JSON(http.StatusOK,ok)
}


// @Summary 修改订单内容（仅在订单未支付的时候可以修改所有的订单）
// @Description 修改订单内容（仅在订单未支付的时候可以修改所有的订单）
// @Tags 订单
// @accept json
// @Produce  json
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /orderdelete/{orderId} [post]
func ALLUpdateBuyOrder(c *gin.Context)  {
	type Arr struct {
		OrderId []int `json:"order_id"`
	}
	var AllOrder Arr
	c.ShouldBindJSON(&AllOrder)
	err := models.UpdateOBuyALL(AllOrder.OrderId)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"code":223,
			"err" : "批量添加失败",
			"data" : nil,
		})
		return
	}
	c.JSON(http.StatusOK,ok)
}

// @Summary 修改订单内容（仅在订单未支付的时候可以修改所有的订单）
// @Description 修改订单内容（仅在订单未支付的时候可以修改所有的订单）
// @Tags 订单
// @accept json
// @Produce  json
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /orderdelete/{orderId} [post]
func CreateBuyOrder(c *gin.Context)  {
	var Order *models.Order
	c.ShouldBindJSON(Order)
	err := models.CreateO(Order)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err" : "创建失败",
			"code":223,
			"data" : nil,
		})
		return
	}
	c.JSON(http.StatusOK,ok)
}
