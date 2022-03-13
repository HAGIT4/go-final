package service

import (
	"strconv"

	pkgService "github.com/HAGIT4/go-final/pkg/service"
	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

func (sv *BonusService) GetUserBalance(req *pkgService.GetUserBalanceRequest) (resp *pkgService.GetUserBalanceResponse) {
	resp = &pkgService.GetUserBalanceResponse{}
	username := req.GetUsername()
	userId, found, err := sv.getUserIdByUsername(username)
	if err != nil {
		resp.Status = pkgService.GetUserBalanceResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !found {
		resp.Status = pkgService.GetUserBalanceResponse_UNAUTHORIZED
		return resp
	}

	current, withdrawn, foundBalance, err := sv.getBalanceByUserId(userId)
	if err != nil {
		resp.Status = pkgService.GetUserBalanceResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !foundBalance {
		resp.Status = pkgService.GetUserBalanceResponse_UNAUTHORIZED
		return resp
	}
	resp = &pkgService.GetUserBalanceResponse{
		Status:    pkgService.GetUserBalanceResponse_OK,
		Current:   current,
		Withdrawn: withdrawn,
	}
	return resp
}

func (sv *BonusService) GetUserWithdrawals(req *pkgService.GetAllWithdrawalsByUserRequest) (resp *pkgService.GetAllWithdrawalsByUserResponse) {
	resp = &pkgService.GetAllWithdrawalsByUserResponse{}
	username := req.Username
	userId, found, err := sv.getUserIdByUsername(username)
	if err != nil {
		resp.Status = pkgService.GetAllWithdrawalsByUserResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !found {
		resp.Status = pkgService.GetAllWithdrawalsByUserResponse_UNAUTHORIZED
		return resp
	}

	withdrawalsList, err := sv.getWithdrawalsByUserId(userId)
	if err != nil {
		resp.Status = pkgService.GetAllWithdrawalsByUserResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if len(withdrawalsList) == 0 {
		resp.Status = pkgService.GetAllWithdrawalsByUserResponse_NO_DATA
		return resp
	}
	resp.WithdrawalInfo = withdrawalsList
	return resp
}

func (sv *BonusService) getWithdrawalsByUserId(userId int) (witdrawalsList []*pkgService.Withdrawal, err error) {
	dbReq := &modelStorage.GetAllWithdrawalsByUserIdRequest{
		UserId: userId,
	}
	dbResp, err := sv.storage.GetAllWithdrawalsByUserId(dbReq)
	if err != nil {
		return nil, err
	}
	for _, dbWithdrawal := range dbResp.Withdrawals {
		svWithdrawal := &pkgService.Withdrawal{
			Order:       strconv.Itoa(dbWithdrawal.OrderId),
			Sum:         float32(dbWithdrawal.Sum) / 100,
			ProcessedAt: dbWithdrawal.ProcessedAt.String(),
		}
		witdrawalsList = append(witdrawalsList, svWithdrawal)
	}
	return witdrawalsList, nil
}

func (sv *BonusService) getBalanceByUserId(userId int) (current float32, withdrawn float32, found bool, err error) {
	dbReq := &modelStorage.GetBalanceByUserIdRequest{
		UserId: userId,
	}
	dbResp, err := sv.storage.GetBalanceByUserId(dbReq)
	if err != nil {
		return 0, 0, false, err
	}
	if dbResp.Found {
		current = float32(dbResp.Current) / 100
		withdrawn = float32(dbResp.Withdrawn) / 100
	}
	return current, withdrawn, dbResp.Found, nil
}
