package storage

import (
	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

type BonusStorageInterface interface {
	AddUser(req *modelStorage.AddUserRequest) (err error)
	AddUserBalance(req *modelStorage.AddUserBalanceRequest) (resp *modelStorage.AddUserBalanceResponse, err error)
	GetUserByUsername(req *modelStorage.GetUserByUsernameRequest) (resp *modelStorage.GetUserByUsernameResponse, err error)
	GetUserIdByUsername(req *modelStorage.GetUserIdByUsernameRequest) (resp *modelStorage.GetUserIdByUsernameResponse, err error)
	GetBalanceByUserId(req *modelStorage.GetBalanceByUserIdRequest) (resp *modelStorage.GetBalanceByUserIdResponse, err error)
	UploadOrder(req *modelStorage.UploadOrderRequest) (resp *modelStorage.UploadOrderResponse, err error)
	GetOrderByOrderId(req *modelStorage.GetOrderByOrderIdRequest) (resp *modelStorage.GetOrderByOrderIdResponse, err error)
	GetAllOrdersFromUser(req *modelStorage.GetAllOrdersFromUserRequest) (resp *modelStorage.GetAllOrdersFromUserResponse, err error)
	GetAllWithdrawalsByUserId(req *modelStorage.GetAllWithdrawalsByUserIdRequest) (resp *modelStorage.GetAllWithdrawalsByUserIdResponse, err error)
	AddWithdrawal(req *modelStorage.AddWithdrawalRequest) (resp *modelStorage.AddWithdrawalResponse, err error)
	GetOrdersForProcess(req *modelStorage.GetOrdersForProcessRequest) (resp *modelStorage.GetOrdersForProcessResponse, err error)
	MarkNewWithProcessing(req *modelStorage.MarkNewWithProcessingRequest) (resp *modelStorage.MarkNewWithProcessingResponse, err error)
	UpdateOrder(req *modelStorage.UpdateOrderRequest) (resp *modelStorage.UpdateOrderResponse, err error)
}
