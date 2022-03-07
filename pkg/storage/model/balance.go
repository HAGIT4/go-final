package model

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
