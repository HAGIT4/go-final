package router

import "github.com/HAGIT4/go-final/internal/service"

type BonusRouterConfig struct {
	Address string
	Service service.BonusServiceInterface
}
