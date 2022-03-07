package service

import (
	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

func (sv *BonusService) getUserIdByUsername(username string) (id int, found bool, err error) {
	dbReq := &modelStorage.GetUserIdByUsernameRequest{
		Username: username,
	}
	dbResp, err := sv.storage.GetUserIdByUsername(dbReq)
	if err != nil {
		return 0, false, err
	}
	return dbResp.UserId, dbResp.Found, nil
}
