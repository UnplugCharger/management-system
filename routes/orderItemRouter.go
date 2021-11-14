package routes

import (
	"github.com/byron/rest/controllers"
	"github.com/gin-gonic/gin"
)



func OrderItemRoutes(incommingRoutes *gin.Engine)  {
	incommingRoutes.GET("/orderItems",controllers.GetOrderItems())
	incommingRoutes.GET("/orderItems/:orderItem_id",controllers.GetOrderItem())
	incommingRoutes.PATCH("/orderItem/:orderItem_id", controllers.UpdateOrderItem())
	incommingRoutes.POST("/orderItems", controllers.CreateOrderItems())

}