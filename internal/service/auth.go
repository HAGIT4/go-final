package service

import (
	pkgService "github.com/HAGIT4/go-final/pkg/service"
	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
	"golang.org/x/crypto/bcrypt"
)

func (sv *BonusService) Register(req *pkgService.RegisterRequest) (resp *pkgService.RegisterResponse) {
	resp = &pkgService.RegisterResponse{}
	dbReq := &modelStorage.GetUserByUsernameRequest{
		Username: req.Login,
	}
	userInDB, err := sv.storage.GetUserByUsername(dbReq)
	if err != nil {
		resp.Status = pkgService.RegisterResponse_INTERNAL_SERVEL_ERROR
		return resp
	}
	if userInDB.Found {
		resp.Status = pkgService.RegisterResponse_LOGIN_IN_USE
		return resp
	}

	passwdHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcryptCost)
	if err != nil {
		resp.Status = pkgService.RegisterResponse_INTERNAL_SERVEL_ERROR
		return resp
	}

	dbAddReq := &modelStorage.AddUserRequest{
		User: modelStorage.User{
			Username:   req.Login,
			PasswdHash: string(passwdHash),
		},
	}
	err = sv.storage.AddUser(dbAddReq)
	if err != nil {
		resp.Status = pkgService.RegisterResponse_INTERNAL_SERVEL_ERROR
		return resp
	}

	userID, found, err := sv.getUserIDByUsername(req.Login)
	if err != nil || !found {
		resp.Status = pkgService.RegisterResponse_INTERNAL_SERVEL_ERROR
		return resp
	}

	dbAddBalanceReq := &modelStorage.AddUserBalanceRequest{
		UserId: userID,
	}
	_, err = sv.storage.AddUserBalance(dbAddBalanceReq)
	if err != nil {
		resp.Status = pkgService.RegisterResponse_INTERNAL_SERVEL_ERROR
		return resp
	}

	token, err := sv.authService.GenerateToken(req.Login)
	if err != nil {
		resp.Status = pkgService.RegisterResponse_INTERNAL_SERVEL_ERROR
		return resp
	}
	resp = &pkgService.RegisterResponse{
		Status:    pkgService.RegisterResponse_OK,
		AuthToken: token,
	}

	return resp
}

func (sv *BonusService) Login(req *pkgService.LoginRequest) (resp *pkgService.LoginResponse) {
	resp = &pkgService.LoginResponse{}
	user, userFound, err := sv.getUserByUsername(req.Login)
	if err != nil {
		resp.Status = pkgService.LoginResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !userFound {
		resp.Status = pkgService.LoginResponse_UNAUTHORIZED
		return resp
	}

	dbHash := []byte(user.PasswdHash)

	if err := bcrypt.CompareHashAndPassword(dbHash, []byte(req.Password)); err != nil {
		resp.Status = pkgService.LoginResponse_UNAUTHORIZED
		return resp
	} else {
		token, err := sv.authService.GenerateToken(req.Login)
		if err != nil {
			resp.Status = pkgService.LoginResponse_INTERNAL_SERVER_ERROR
			return resp
		}
		resp = &pkgService.LoginResponse{
			Status:    pkgService.LoginResponse_OK,
			AuthToken: token,
		}
	}
	return resp
}

func (sv *BonusService) Authenticate(req *pkgService.AuthRequest) (resp *pkgService.AuthResponse) {
	resp = &pkgService.AuthResponse{}
	token, err := sv.authService.ValidateToken(req.Token)
	if err != nil {
		resp.Status = pkgService.AuthResponse_UNAUTHORIZED
		return resp
	}
	claims, ok := token.Claims.(*authClaims)
	if !ok {
		resp.Status = pkgService.AuthResponse_INTERNAL_SERVER_ERROR
		return resp
	} else {
		resp.Status = pkgService.AuthResponse_OK
		resp.Username = claims.Username
		return resp
	}
}
