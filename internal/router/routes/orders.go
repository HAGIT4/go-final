package routes

import (
	"github.com/HAGIT4/go-final/internal/service"
	"github.com/gin-gonic/gin"
)

func AddOrdersRoutes(rg *gin.RouterGroup, sv service.BonusServiceInterface) {
	rg.POST("/orders", uploadOrderHandler(sv))
	rg.GET("/orders", getOrderListHandler(sv))
}

func uploadOrderHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {

	}
	return
}

func getOrderListHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {

	}
	return
}
