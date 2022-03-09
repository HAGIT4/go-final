package service

import (
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
