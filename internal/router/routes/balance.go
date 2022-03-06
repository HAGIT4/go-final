package routes

import (
	"github.com/HAGIT4/go-final/internal/service"
	"github.com/gin-gonic/gin"
)

func AddBalanceRoutes(rg *gin.RouterGroup, sv service.BonusServiceInterface) {
	rg.GET("/balance", getBalanceHandler(sv))
	rg.POST("/balance/withdraw", withdrawHandler(sv))
	rg.POST("/balance/withdrawals", getWithdrawalsInfoHandler(sv))
}

func getBalanceHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {

	}
	return
}

func withdrawHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {

	}
	return
}

func getWithdrawalsInfoHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {

	}
	return
}
