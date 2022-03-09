package routes

import (
	"net/http"

	service "github.com/HAGIT4/go-final/internal/service"
	pkgService "github.com/HAGIT4/go-final/pkg/service"
	"github.com/gin-gonic/gin"
)

func AddBalanceRoutes(rg *gin.RouterGroup, sv service.BonusServiceInterface) {
	rg.GET("/balance", getBalanceHandler(sv))
	rg.POST("/balance/withdraw", withdrawHandler(sv))
	rg.POST("/balance/withdrawals", getWithdrawalsInfoHandler(sv))
}

func getBalanceHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {
		svReq := &pkgService.GetUserBalanceRequest{
			Username: "test",
		}
		svResp := sv.GetUserBalance(svReq)
		switch svResp.GetStatus() {
		case pkgService.GetUserBalanceResponse_OK:
			svResp.Status = 0
			c.JSON(http.StatusOK, *svResp)
			return
		case pkgService.GetUserBalanceResponse_UNAUTHORIZED:
			svResp.Status = 0 // maybe new error type
			c.AbortWithStatus(http.StatusUnauthorized)
		case pkgService.GetUserBalanceResponse_INTERNAL_SERVER_ERROR:
			err := NewBonusRouterInternalServerError()
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
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
