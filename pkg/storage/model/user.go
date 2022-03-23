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

type GetUserIDByUsernameRequest struct {
	Username string
}

type GetUserIDByUsernameResponse struct {
	UserID int
	Found  bool
}
