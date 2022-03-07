package service

import (
	pkgService "github.com/HAGIT4/go-final/pkg/service"
	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
	"golang.org/x/crypto/bcrypt"
)

func (sv *BonusService) Register(req pkgService.RegisterRequest) (resp *pkgService.RegisterResponse, err error) {
	dbReq := &modelStorage.GetUserByUsernameRequest{
		Username: req.Login,
	}
	userInDB, err := sv.storage.GetUserByUsername(dbReq)
	if err != nil {
		return nil, err
	}
	if !userInDB.Found {
		return nil, newBonusServiceLoginInUseError()
	}

	passwdHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	dbAddReq := &modelStorage.AddUserRequest{
		User: modelStorage.User{
			Username:   req.Login,
			PasswdHash: string(passwdHash),
		},
	}
	err = sv.storage.AddUser(dbAddReq)
	if err != nil {
		return nil, err
	}
	return &pkgService.RegisterResponse{}, nil
}
