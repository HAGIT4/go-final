package storage

import (
	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

type BonusStorageInterface interface {
	AddUser(req *modelStorage.AddUserRequest) (err error)
	AddUserBalance(req *modelStorage.AddUserBalanceRequest) (resp *modelStorage.AddUserBalanceResponse, err error)
	GetUserByUsername(req *modelStorage.GetUserByUsernameRequest) (resp *modelStorage.GetUserByUsernameResponse, err error)
	GetUserIDByUsername(req *modelStorage.GetUserIDByUsernameRequest) (resp *modelStorage.GetUserIDByUsernameResponse, err error)
	GetBalanceByUserID(req *modelStorage.GetBalanceByUserIDRequest) (resp *modelStorage.GetBalanceByUserIDResponse, err error)
	UploadOrder(req *modelStorage.UploadOrderRequest) (resp *modelStorage.UploadOrderResponse, err error)
	GetOrderByOrderID(req *modelStorage.GetOrderByOrderIDRequest) (resp *modelStorage.GetOrderByOrderIDResponse, err error)
	GetAllOrdersFromUser(req *modelStorage.GetAllOrdersFromUserRequest) (resp *modelStorage.GetAllOrdersFromUserResponse, err error)
	GetAllWithdrawalsByUserID(req *modelStorage.GetAllWithdrawalsByUserIDRequest) (resp *modelStorage.GetAllWithdrawalsByUserIDResponse, err error)
	AddWithdrawal(req *modelStorage.AddWithdrawalRequest) (resp *modelStorage.AddWithdrawalResponse, err error)
	GetOrdersForProcess(req *modelStorage.GetOrdersForProcessRequest) (resp *modelStorage.GetOrdersForProcessResponse, err error)
	MarkNewWithProcessing(req *modelStorage.MarkNewWithProcessingRequest) (resp *modelStorage.MarkNewWithProcessingResponse, err error)
	UpdateOrder(req *modelStorage.UpdateOrderRequest) (resp *modelStorage.UpdateOrderResponse, err error)
	AddSumToUserBalance(req *modelStorage.AddSumToUserBalanceRequest) (resp *modelStorage.AddSumToUserBalanceResponse, err error)
}
