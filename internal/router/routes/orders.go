package routes

import (
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
	rg.GET("/orders", middleware.AuthenticateUserMiddleware(sv), getOrderListHandler(sv))
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
		switch svResp.GetStatus() {
		case pkgService.UploadOrderResponse_ALREADY_UPLOADED_BY_THIS_USER:
			c.Status(http.StatusOK)
			return
		case pkgService.UploadOrderResponse_OK:
			c.Status(http.StatusAccepted)
			return
		case pkgService.UploadOrderResponse_BAD_REQUEST:
			c.Status(http.StatusBadRequest)
			return
		case pkgService.UploadOrderResponse_BAD_ORDER_NUMBER:
			c.Status(http.StatusUnprocessableEntity)
			return
		case pkgService.UploadOrderResponse_UNAUTHORIZED:
			c.Status(http.StatusUnauthorized)
			return
		case pkgService.UploadOrderResponse_ALREADY_UPLOADED_BY_ANOTHER_USER:
			c.Status(http.StatusUnprocessableEntity)
			return
		case pkgService.UploadOrderResponse_INTERNAL_SERVER_ERROR:
			c.Status(http.StatusInternalServerError)
			return
		default:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	return
}

func getOrderListHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {
		username := c.GetString("username")
		svReq := &pkgService.GetOrderListRequest{
			Username: username,
		}
		svResp := sv.GetAllOrdersFromUser(svReq)
		switch svResp.Status {
		case pkgService.GetOrderListResponse_INTERNAL_SERVER_ERROR:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		case pkgService.GetOrderListResponse_NO_DATA:
			c.Status(http.StatusNoContent)
			return
		case pkgService.GetOrderListResponse_OK:
			c.JSON(http.StatusOK, svResp.OrderInfo)
			return
		}
	}
	return
}
