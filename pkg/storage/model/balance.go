package model

import "time"

type Balance struct {
	Current   int
	Withdrawn int
}

type GetBalanceByUserIDRequest struct {
	UserID int
}

type GetBalanceByUserIDResponse struct {
	Balance
	Found bool
}

type GetAllWithdrawalsByUserIdRequest struct {
	UserID int
}

type Withdrawal struct {
	OrderId     int
	Sum         int
	UserID      int
	ProcessedAt time.Time
}

type GetAllWithdrawalsByUserIdResponse struct {
	Withdrawals []Withdrawal
}

type SetUserBalanceByUserIdRequest struct {
	UserID    int
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
	UserID int
}

type AddUserBalanceResponse struct {
}

type AddSumToUserBalanceRequest struct {
	UserID int
	Sum    int
}

type AddSumToUserBalanceResponse struct {
}
