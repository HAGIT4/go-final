package model

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

type GetUserIdByUsernameRequest struct {
	Username string
}

type GetUserIdByUsernameResponse struct {
	UserId int
	Found  bool
}
