package service

import (
	"strconv"
	"time"

	pkgService "github.com/HAGIT4/go-final/pkg/service"
	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

func (sv *BonusService) GetUserBalance(req *pkgService.GetUserBalanceRequest) (resp *pkgService.GetUserBalanceResponse) {
	resp = &pkgService.GetUserBalanceResponse{}
	username := req.GetUsername()
	userID, found, err := sv.getUserIdByUsername(username)
	if err != nil {
		resp.Status = pkgService.GetUserBalanceResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !found {
		resp.Status = pkgService.GetUserBalanceResponse_UNAUTHORIZED
		return resp
	}

	current, withdrawn, foundBalance, err := sv.getBalanceByUserID(userID)
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

func (sv *BonusService) addSumToUserBalance(userID int, accrual float32) (err error) {

	balance, _, _, err := sv.getBalanceByUserID(userID)
	if err != nil {
		return err
	}
	dbReq := &modelStorage.AddSumToUserBalanceRequest{
		UserID: userID,
		Sum:    int(balance*100) + int(accrual*100),
	}
	if _, err = sv.storage.AddSumToUserBalance(dbReq); err != nil {
		return err
	}
	return nil
}

func (sv *BonusService) GetUserWithdrawals(req *pkgService.GetAllWithdrawalsByUserRequest) (resp *pkgService.GetAllWithdrawalsByUserResponse) {
	resp = &pkgService.GetAllWithdrawalsByUserResponse{}
	username := req.Username
	userID, found, err := sv.getUserIdByUsername(username)
	if err != nil {
		resp.Status = pkgService.GetAllWithdrawalsByUserResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !found {
		resp.Status = pkgService.GetAllWithdrawalsByUserResponse_UNAUTHORIZED
		return resp
	}

	withdrawalsList, err := sv.getWithdrawalsByUserID(userID)
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

func (sv *BonusService) Withdraw(req *pkgService.WithdrawRequest) (resp *pkgService.WithdrawResponse) {
	resp = &pkgService.WithdrawResponse{}
	username := req.Username
	userID, found, err := sv.getUserIdByUsername(username)
	if err != nil {
		resp.Status = pkgService.WithdrawResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !found {
		resp.Status = pkgService.WithdrawResponse_UNAUTHORIZED
		return resp
	}

	userCurrent, userWithdrawn, balanceFound, err := sv.getBalanceByUserID(userID)
	if err != nil {
		resp.Status = pkgService.WithdrawResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !balanceFound {
		resp.Status = pkgService.WithdrawResponse_UNAUTHORIZED
		return resp
	}

	if req.Sum > userCurrent {
		resp.Status = pkgService.WithdrawResponse_INSUFFICIENT_FUNDS
		return resp
	} else {
		newUserCurrent := userCurrent - req.Sum
		newUserWithdrawn := userWithdrawn + req.Sum
		orderID, err := strconv.Atoi(req.Order)
		if err != nil {
			resp.Status = pkgService.WithdrawResponse_BAD_ORDER_NUMBER
			return resp
		}
		if err = sv.AddWithdrawal(newUserCurrent, newUserWithdrawn, req.Sum, userID, orderID); err != nil {
			resp.Status = pkgService.WithdrawResponse_INTERNAL_SERVER_ERROR
			return resp
		}
		resp.Status = pkgService.WithdrawResponse_OK
		return resp
	}
}

func (sv *BonusService) AddWithdrawal(current float32, withdrawn float32, sum float32, userID int, orderId int) (err error) {
	dbReq := &modelStorage.AddWithdrawalRequest{
		UserId:      userID,
		Current:     int(current * 100),
		Withdrawn:   int(withdrawn * 100),
		Sum:         int(sum * 100),
		ProcessedAt: time.Now(),
	}
	_, err = sv.storage.AddWithdrawal(dbReq)
	if err != nil {
		return err
	}
	return nil
}

func (sv *BonusService) getWithdrawalsByUserID(userID int) (witdrawalsList []*pkgService.Withdrawal, err error) {
	dbReq := &modelStorage.GetAllWithdrawalsByUserIdRequest{
		UserId: userID,
	}
	dbResp, err := sv.storage.GetAllWithdrawalsByUserID(dbReq)
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

func (sv *BonusService) getBalanceByUserID(userId int) (current float32, withdrawn float32, found bool, err error) {
	dbReq := &modelStorage.GetBalanceByUserIdRequest{
		UserID: userId,
	}
	dbResp, err := sv.storage.GetBalanceByUserID(dbReq)
	if err != nil {
		return 0, 0, false, err
	}
	if dbResp.Found {
		current = float32(dbResp.Current) / 100
		withdrawn = float32(dbResp.Withdrawn) / 100
	}
	return current, withdrawn, dbResp.Found, nil
}
