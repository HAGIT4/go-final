package model

type RouterRegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RouterLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
