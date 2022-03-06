package routes

import (
	"github.com/HAGIT4/go-final/internal/service"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup, sv service.BonusServiceInterface) {
	rg.POST("/register", registerHandler(sv))
	rg.POST("/login", loginHandler(sv))
}

func registerHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {

	}
	return
}

func loginHandler(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {

	}
	return
}
