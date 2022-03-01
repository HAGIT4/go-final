package models

type AddUserRequest struct {
	Username   string
	PasswdHash string
}

type GetUserByUsernameRequest struct{}

type GetUserByUsernameResponse struct{}
