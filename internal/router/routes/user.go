package routes

import (
	"net/http"

	service "github.com/HAGIT4/go-final/internal/service"
	modelRouter "github.com/HAGIT4/go-final/pkg/router/model"
	pkgService "github.com/HAGIT4/go-final/pkg/service"
	gin "github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup, sv service.BonusServiceInterface) {
	rg.POST("/register", registerHandler(sv))
	rg.POST("/login", loginHandler(sv))
}

func registerHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {
		contentHeader := c.Request.Header.Get("Content-Type")
		if contentHeader != "application/json" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		rtReq := &modelRouter.RouterRegisterRequest{}
		if err := c.BindJSON(rtReq); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		svReq := &pkgService.RegisterRequest{
			Login:    rtReq.Login,
			Password: rtReq.Password,
		}
		svResp := sv.Register(svReq)
		switch svResp.GetStatus() {
		case pkgService.RegisterResponse_OK:
			c.Header("Authorization", svResp.AuthToken)
			c.Status(http.StatusOK)
			return
		case pkgService.RegisterResponse_LOGIN_IN_USE:
			c.AbortWithStatus(http.StatusConflict)
			return
		case pkgService.RegisterResponse_INTERNAL_SERVEL_ERROR:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	return
}

func loginHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {
		contentHeader := c.Request.Header.Get("Content-Type")
		if contentHeader != "application/json" {
			c.AbortWithStatus(http.StatusBadRequest) // add error
			return
		}

		rtReq := &modelRouter.RouterLoginRequest{}
		if err := c.BindJSON(rtReq); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		svReq := &pkgService.LoginRequest{
			Login:    rtReq.Login,
			Password: rtReq.Password,
		}
		svResp := sv.Login(svReq)
		switch svResp.GetStatus() {
		case pkgService.LoginResponse_OK:
			c.Header("Authorization", svResp.AuthToken)
			c.Status(http.StatusOK)
			return
		case pkgService.LoginResponse_UNAUTHORIZED:
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		case pkgService.LoginResponse_INTERNAL_SERVER_ERROR:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	return
}
