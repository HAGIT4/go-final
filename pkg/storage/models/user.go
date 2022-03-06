package models

type User struct {
	Username   string
	PasswdHash string
}

type AddUserRequest struct {
	User
}

type GetUserByUsernameRequest struct {
	Username string
}

type GetUserByUsernameResponse struct {
	User
	Found bool
}
