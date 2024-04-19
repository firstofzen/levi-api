package router

import (
	"github.com/gin-gonic/gin"
	"levi-api/api/handler"
)

type UserRouter struct {
	E *gin.Engine
}

var uh handler.UserHandler

func (r *UserRouter) route() {
	r.E.GET("/login", func(c *gin.Context) {
		uh.Login(c)
	})
	r.E.GET("/login/google/callback", func(c *gin.Context) {
		uh.LoginCallback(c)
	})
	r.E.GET("/login/access-token", func(c *gin.Context) {
		uh.LoginAccessToken(c)
	})
	r.E.GET("/user/account/add", func(c *gin.Context) {
		uh.AddAccount(c)
	})
	r.E.GET("/user/detail/update", func(c *gin.Context) {
		uh.UserDetailUpdate(c)
	})
	r.E.GET("/user/account/delete", func(c *gin.Context) {
		uh.AccountDelete(c)
	})
	r.E.GET("/cart/get/id", func(c *gin.Context) {
		uh.GetCart(c)
	})
	r.E.GET("/booth/get/id", func(c *gin.Context) {
		uh.GetBooth(c)
	})
}
