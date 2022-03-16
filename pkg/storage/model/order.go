package model

import (
	"time"
)

type UploadOrderRequest struct {
	Number     int
	Status     string
	Accural    int
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

type GetAllOrdersFromUserRequest struct {
	UserId int
}

type Order struct {
	Number     int
	Status     string
	Accural    int
	UploadedAt time.Time
}

type GetAllOrdersFromUserResponse struct {
	Orders []Order
}

type GetOrdersForProcessRequest struct {
}

type ProcessedOrder struct {
	Number int
	UserId int
}

type GetOrdersForProcessResponse struct {
	ProcessedOrders []ProcessedOrder
}

type MarkNewWithProcessingRequest struct{}

type MarkNewWithProcessingResponse struct{}
