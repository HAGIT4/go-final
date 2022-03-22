package api

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	router "github.com/HAGIT4/go-final/internal/router"
	service "github.com/HAGIT4/go-final/internal/service"
	storage "github.com/HAGIT4/go-final/internal/storage"
	pkgApi "github.com/HAGIT4/go-final/pkg/api"
	pkgRouter "github.com/HAGIT4/go-final/pkg/router"
	pkgService "github.com/HAGIT4/go-final/pkg/service"
	pkgStorage "github.com/HAGIT4/go-final/pkg/storage"
)

type bonusServer struct {
	router  router.BonusRouterInterface
	service service.BonusServiceInterface
	storage storage.BonusStorageInterface
}

func NewBonusServer(cfg *pkgApi.APIConfig) (bs *bonusServer, err error) {
	dbCfg := &pkgStorage.BonusStorageConfig{
		ConnectionString: cfg.DatabaseUri,
	}
	st, err := storage.NewBonusStorage(dbCfg)
	if err != nil {
		return nil, err
	}

	svCfg := &pkgService.BonusServiceConfig{
		Storage:              st,
		AccuralSystemAddress: cfg.AccuralSystemAddress,
	}
	fmt.Println("Service config:", svCfg)
	sv, err := service.NewBonusService(svCfg)
	if err != nil {
		return nil, err
	}

	rtCfg := &pkgRouter.BonusRouterConfig{
		Address: cfg.RunAddress,
		Service: sv,
	}
	rt, err := router.NewBonusRouter(rtCfg)
	if err != nil {
		return nil, err
	}

	bs = &bonusServer{
		router:  rt,
		service: sv,
		storage: st,
	}
	return bs, nil
}

func (bs *bonusServer) ListenAndServe() (err error) {
	go func() {
		if err := bs.router.Run(); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if err := bs.service.ProcessOrders(); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-quit
	log.Println("Server shutdown...")
	return nil
}
