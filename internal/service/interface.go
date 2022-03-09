package service

import pkgService "github.com/HAGIT4/go-final/pkg/service"

type BonusServiceInterface interface {
	Register(req *pkgService.RegisterRequest) (resp *pkgService.RegisterResponse)
	Login(req *pkgService.LoginRequest) (resp *pkgService.LoginResponse)
	GetUserBalance(req *pkgService.GetUserBalanceRequest) (resp *pkgService.GetUserBalanceResponse)
}
