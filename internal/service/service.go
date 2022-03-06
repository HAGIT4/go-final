package service

import (
	storage "github.com/HAGIT4/go-final/internal/storage"
	pkgService "github.com/HAGIT4/go-final/pkg/service"
)

const (
	bcryptCost int = 8
)

type BonusService struct {
	storage     storage.BonusStorageInterface
	authService *authService
}

var _ BonusServiceInterface = (*BonusService)(nil)

func NewBonusService(cfg *pkgService.BonusServiceConfig) (sv *BonusService, err error) {
	asv := NewAuthService()
	sv = &BonusService{
		storage:     cfg.Storage,
		authService: asv,
	}
	return sv, nil
}
