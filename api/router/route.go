package router

import "github.com/gin-gonic/gin"

type Router struct{}

func (r Router) InitDefault() {
	eng := gin.Default()
	userRouter := UserRouter{E: eng}
	userRouter.route()
	productRouter := ProductRouter{E: eng}
	productRouter.route()
	tradeRouter := TradeRouter{E: eng}
	tradeRouter.route()
}
