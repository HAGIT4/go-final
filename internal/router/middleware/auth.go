package middleware

import (
	"net/http"

	service "github.com/HAGIT4/go-final/internal/service"
	pkgService "github.com/HAGIT4/go-final/pkg/service"
	gin "github.com/gin-gonic/gin"
)

func AuthenticateUserMiddleware(sv service.BonusServiceInterface) (h gin.HandlerFunc) {
	h = func(c *gin.Context) {
		authHeader := c.Request.Header["Authorization"][0]
		authReq := &pkgService.AuthRequest{
			Token: authHeader,
		}
		svResp := sv.Authenticate(authReq)
		switch svResp.Status {
		case pkgService.AuthResponse_INTERNAL_SERVER_ERROR:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		case pkgService.AuthResponse_UNAUTHORIZED:
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		case pkgService.AuthResponse_OK:
			c.Set("username", svResp.Username)
			return
		}
	}
	return
}
