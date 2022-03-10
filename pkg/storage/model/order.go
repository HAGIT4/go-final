package model

import "time"

type UploadOrderRequest struct {
	Number     int
	Status     string
	UserId     int
	UploadedAt time.Time
}

type UploadOrderResponse struct{}

type GetOrderByOrderIdRequest struct {
	OrderId int
}

type GetOrderByOrderIdResponse struct {
	UserId     int
	Status     string
	UploadedAt time.Time
}
