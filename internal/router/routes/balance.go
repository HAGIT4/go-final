package routes

import (
	"net/http"

	"github.com/HAGIT4/go-final/internal/router/middleware"
	service "github.com/HAGIT4/go-final/internal/service"
	pkgService "github.com/HAGIT4/go-final/pkg/service"
	gin "github.com/gin-gonic/gin"
)

func AddBalanceRoutes(rg *gin.RouterGroup, sv service.BonusServiceInterface) {
	rg.GET("/balance", middleware.AuthenticateUserMiddleware(sv), getBalanceHandler(sv))
	rg.POST("/balance/withdraw", middleware.AuthenticateUserMiddleware(sv), withdrawHandler(sv))
	rg.POST("/balance/withdrawals", middleware.AuthenticateUserMiddleware(sv), getWithdrawalsInfoHandler(sv))
}

func getBalanceHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {
		username := c.GetString("username")
		svReq := &pkgService.GetUserBalanceRequest{
			Username: username,
		}
		svResp := sv.GetUserBalance(svReq)
		switch svResp.GetStatus() {
		case pkgService.GetUserBalanceResponse_OK:
			svResp.Status = 0
			c.JSON(http.StatusOK, *svResp)
			return
		case pkgService.GetUserBalanceResponse_UNAUTHORIZED:
			svResp.Status = 0
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
		username := c.GetString("username")
		svReq := &pkgService.WithdrawRequest{}
		c.BindJSON(svReq)
		svReq.Username = username
		svResp := sv.Withdraw(svReq)
		switch svResp.GetStatus() {
		case pkgService.WithdrawResponse_OK:
			c.Status(http.StatusOK)
			return
		case pkgService.WithdrawResponse_UNAUTHORIZED:
			c.Status(http.StatusUnauthorized)
			return
		case pkgService.WithdrawResponse_INSUFFICIENT_FUNDS:
			c.Status(http.StatusPaymentRequired)
			return
		case pkgService.WithdrawResponse_BAD_ORDER_NUMBER:
			c.Status(http.StatusUnprocessableEntity)
			return
		case pkgService.WithdrawResponse_INTERNAL_SERVER_ERROR:
			c.Status(http.StatusInternalServerError)
			return
		}
	}
	return
}

func getWithdrawalsInfoHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {
		username := c.GetString("username")
		svReq := &pkgService.GetAllWithdrawalsByUserRequest{
			Username: username,
		}
		svResp := sv.GetUserWithdrawals(svReq)
		switch svResp.GetStatus() {
		case pkgService.GetAllWithdrawalsByUserResponse_OK:
			c.JSON(http.StatusOK, svResp.GetWithdrawalInfo())
			return
		case pkgService.GetAllWithdrawalsByUserResponse_NO_DATA:
			c.Status(http.StatusNoContent)
			return
		case pkgService.GetAllWithdrawalsByUserResponse_UNAUTHORIZED:
			c.Status(http.StatusUnauthorized)
			return
		case pkgService.GetAllWithdrawalsByUserResponse_INTERNAL_SERVER_ERROR:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	return
}
