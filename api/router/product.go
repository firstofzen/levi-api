package router

import (
	"github.com/gin-gonic/gin"
	"levi-api/api/handler"
)

type ProductRouter struct {
	E *gin.Engine
}

var ph handler.ProductHandler

func (r *ProductRouter) route() {
	r.E.GET("/product/get/category", func(c *gin.Context) {
		ph.GetProductByCategory(c)
	})
	r.E.GET("/product/get_top/category", func(c *gin.Context) {
		ph.GetTopProductByCategory(c)
	})
	r.E.PATCH("/product/update", func(c *gin.Context) {
		ph.UpdateProduct(c)
	})
	r.E.DELETE("/product/delete", func(c *gin.Context) {
		ph.DeleteProduct(c)
	})
	r.E.GET("/product/rating/get", func(c *gin.Context) {
		ph.GetRatingProduct(c)
	})
	r.E.PUT("/product/rating/add", func(c *gin.Context) {
		ph.UpdateRatingProduct(c)
	})
	r.E.DELETE("/product/rating/delete", func(c *gin.Context) {
		ph.DeleteRatingProduct(c)
	})
	r.E.PUT("/product/sale/add", func(c *gin.Context) {
		ph.AddSaleProduct(c)
	})
}
