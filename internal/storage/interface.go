package storage

import (
	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

type BonusStorageInterface interface {
	AddUser(req *modelStorage.AddUserRequest) (err error)
	GetUserByUsername(req *modelStorage.GetUserByUsernameRequest) (resp *modelStorage.GetUserByUsernameResponse, err error)
	GetUserIdByUsername(req *modelStorage.GetUserIdByUsernameRequest) (resp *modelStorage.GetUserIdByUsernameResponse, err error)
	GetBalanceByUserId(req *modelStorage.GetBalanceByUserIdRequest) (resp *modelStorage.GetBalanceByUserIdResponse, err error)
	UploadOrder(req *modelStorage.UploadOrderRequest) (resp *modelStorage.UploadOrderResponse, err error)
	GetOrderByOrderId(req *modelStorage.GetOrderByOrderIdRequest) (resp *modelStorage.GetOrderByOrderIdResponse, err error)
	GetAllOrdersFromUser(req *modelStorage.GetAllOrdersFromUserRequest) (resp *modelStorage.GetAllOrdersFromUserResponse, err error)
	GetAllWithdrawalsByUserId(req *modelStorage.GetAllWithdrawalsByUserIdRequest) (resp *modelStorage.GetAllWithdrawalsByUserIdResponse, err error)
}
