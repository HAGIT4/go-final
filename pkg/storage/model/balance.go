package model

import "time"

type Balance struct {
	Current   int
	Withdrawn int
}

type GetBalanceByUserIdRequest struct {
	UserId int
}

type GetBalanceByUserIdResponse struct {
	Balance
	Found bool
}

type GetAllWithdrawalsByUserIdRequest struct {
	UserId int
}

type Withdrawal struct {
	OrderId     int
	Sum         int
	UserId      int
	ProcessedAt time.Time
}

type GetAllWithdrawalsByUserIdResponse struct {
	Withdrawals []Withdrawal
}

type SetUserBalanceByUserIdRequest struct {
	UserId    int
	Current   int
	Withdrawn int
}

type SetUserBalanceByUserIdResponse struct {
}

type AddWithdrawalRequest struct {
	UserId      int
	Current     int
	Withdrawn   int
	Sum         int
	OrderId     int
	ProcessedAt time.Time
}

type AddWithdrawalResponse struct {
}

type AddUserBalanceRequest struct {
	UserId int
}

type AddUserBalanceResponse struct {
}
