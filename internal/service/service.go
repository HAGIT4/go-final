package service

import (
	storage "github.com/HAGIT4/go-final/internal/storage"
	pkgService "github.com/HAGIT4/go-final/pkg/service"
)

type BonusService struct {
	storage storage.BonusStorageInterface
}

var _ BonusServiceInterface = (*BonusService)(nil)

func NewBonusService(cfg *pkgService.BonusServiceConfig) (sv *BonusService, err error) {
	sv = &BonusService{
		storage: cfg.Storage,
	}
	return sv, nil
}
