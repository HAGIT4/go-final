package routes

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	middleware "github.com/HAGIT4/go-final/internal/router/middleware"
	service "github.com/HAGIT4/go-final/internal/service"
	pkgService "github.com/HAGIT4/go-final/pkg/service"
	gin "github.com/gin-gonic/gin"
)

func AddOrdersRoutes(rg *gin.RouterGroup, sv service.BonusServiceInterface) {
	rg.POST("/orders", middleware.AuthenticateUserMiddleware(sv), uploadOrderHandler(sv))
	rg.GET("/orders", getOrderListHandler(sv))
}

func uploadOrderHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {
		username := c.GetString("username")
		orderNumberBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		orderNumber, err := strconv.Atoi(string(orderNumberBytes))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		svReq := &pkgService.UploadOrderRequest{
			Order:    int64(orderNumber),
			Username: username,
		}
		svResp := sv.UploadOrder(svReq)
		fmt.Print(svResp.Status)
	}
	return
}

func getOrderListHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {

	}
	return
}
