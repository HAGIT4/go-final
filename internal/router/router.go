package router

import (
	routes "github.com/HAGIT4/go-final/internal/router/routes"
	routerCfg "github.com/HAGIT4/go-final/pkg/router"
	gin "github.com/gin-gonic/gin"
)

type bonusRouter struct {
	mux *gin.Engine
}

var _ BonusRouterInterface = (*bonusRouter)(nil)

func NewBonusRouter(cfg *routerCfg.BonusRouterConfig) (r *bonusRouter, err error) {
	mux := gin.Default()
	apiUserGroup := mux.Group("api/user")
	routes.AddUserRoutes(apiUserGroup, cfg.Service)
	routes.AddOrdersRoutes(apiUserGroup, cfg.Service)
	routes.AddBalanceRoutes(apiUserGroup, cfg.Service)

	r = &bonusRouter{
		mux: mux,
	}
	return r, nil
}
