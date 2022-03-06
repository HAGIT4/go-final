package service

import (
	pkgService "github.com/HAGIT4/go-final/pkg/service"
	modelsStorage "github.com/HAGIT4/go-final/pkg/storage/models"
	"golang.org/x/crypto/bcrypt"
)

func (sv *BonusService) Register(req pkgService.RegisterRequest) (resp *pkgService.RegisterResponse, err error) {
	dbReq := &modelsStorage.GetUserByUsernameRequest{
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
	dbAddReq := &modelsStorage.AddUserRequest{
		User: modelsStorage.User{
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
