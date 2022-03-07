package main

import (
	"log"

	api "github.com/HAGIT4/go-final/internal/api"
	pkgApi "github.com/HAGIT4/go-final/pkg/api"
)

func main() {
	cfg, err := pkgApi.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	bonusServer, err := api.NewBonusServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := bonusServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
