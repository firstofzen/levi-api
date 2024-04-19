package router

import (
	"github.com/gin-gonic/gin"
	"levi-api/api/handler"
)

type TradeRouter struct {
	E *gin.Engine
}

func (r *TradeRouter) route() {
	var th handler.TradeHandler
	r.E.GET("/order/get/id", func(c *gin.Context) {
		th.GetOrderById(c)
	})
	r.E.GET("/order/get/id_user", func(c *gin.Context) {
		th.GetOrderByIdUser(c)
	})
	r.E.GET("/order/add", func(c *gin.Context) {
		th.AddOrder(c)
	})
	r.E.PATCH("/order/update/status", func(c *gin.Context) {
		th.UpdateStatusOrder(c)
	})
	r.E.DELETE("/order/delete", func(c *gin.Context) {
		th.DeleteOrder(c)
	})
	r.E.GET("/order/ship/get", func(c *gin.Context) {
		th.GetShipOrder(c)
	})
	r.E.PATCH("/order/ship/update/state", func(c *gin.Context) {
		th.UpdateStateShipOrder(c)
	})
}
