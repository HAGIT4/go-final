package service

import pkgService "github.com/HAGIT4/go-final/pkg/service"

type BonusServiceInterface interface {
	Register(req *pkgService.RegisterRequest) (resp *pkgService.RegisterResponse)
	Login(req *pkgService.LoginRequest) (resp *pkgService.LoginResponse)
	Authenticate(req *pkgService.AuthRequest) (resp *pkgService.AuthResponse)
	GetUserBalance(req *pkgService.GetUserBalanceRequest) (resp *pkgService.GetUserBalanceResponse)
	UploadOrder(req *pkgService.UploadOrderRequest) (resp *pkgService.UploadOrderResponse)
	GetAllOrdersFromUser(req *pkgService.GetOrderListRequest) (resp *pkgService.GetOrderListResponse)
	GetUserWithdrawals(req *pkgService.GetAllWithdrawalsByUserRequest) (resp *pkgService.GetAllWithdrawalsByUserResponse)
	Withdraw(req *pkgService.WithdrawRequest) (resp *pkgService.WithdrawResponse)
	ProcessOrders() (err error)
}
