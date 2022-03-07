package router

import (
	routes "github.com/HAGIT4/go-final/internal/router/routes"
	routerCfg "github.com/HAGIT4/go-final/pkg/router"
	gin "github.com/gin-gonic/gin"
)

type bonusRouter struct {
	address string
	mux     *gin.Engine
}

var _ BonusRouterInterface = (*bonusRouter)(nil)

func NewBonusRouter(cfg *routerCfg.BonusRouterConfig) (r *bonusRouter, err error) {
	mux := gin.Default()
	apiUserGroup := mux.Group("api/user")
	routes.AddUserRoutes(apiUserGroup, cfg.Service)
	routes.AddOrdersRoutes(apiUserGroup, cfg.Service)
	routes.AddBalanceRoutes(apiUserGroup, cfg.Service)

	r = &bonusRouter{
		address: cfg.Address,
		mux:     mux,
	}
	return r, nil
}

func (rt *bonusRouter) Run() (err error) {
	if err = rt.mux.Run(rt.address); err != nil {
		return err
	}
	return nil
}
