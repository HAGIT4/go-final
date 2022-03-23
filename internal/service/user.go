package service

import (
	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

func (sv *BonusService) getUserIDByUsername(username string) (id int, found bool, err error) {
	dbReq := &modelStorage.GetUserIDByUsernameRequest{
		Username: username,
	}
	dbResp, err := sv.storage.GetUserIDByUsername(dbReq)
	if err != nil {
		return 0, false, err
	}
	return dbResp.UserID, dbResp.Found, nil
}

func (sv *BonusService) getUserByUsername(username string) (user *modelStorage.User, found bool, err error) {
	dbReq := &modelStorage.GetUserByUsernameRequest{
		Username: username,
	}
	dbResp, err := sv.storage.GetUserByUsername(dbReq)
	if err != nil {
		return nil, false, err
	}
	return &dbResp.User, dbResp.Found, nil
}
