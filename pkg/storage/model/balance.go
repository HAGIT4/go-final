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

type GetAllWithdrawalsByUserIDRequest struct {
	UserID int
}

type Withdrawal struct {
	OrderID     int
	Sum         int
	UserID      int
	ProcessedAt time.Time
}

type GetAllWithdrawalsByUserIDResponse struct {
	Withdrawals []Withdrawal
}

type SetUserBalanceByUserIDRequest struct {
	UserID    int
	Current   int
	Withdrawn int
}

type SetUserBalanceByUserIDResponse struct {
}

type AddWithdrawalRequest struct {
	UserID      int
	Current     int
	Withdrawn   int
	Sum         int
	OrderID     int
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
