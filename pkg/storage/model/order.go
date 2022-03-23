package model

import (
	"time"
)

type UploadOrderRequest struct {
	Number     int
	Status     string
	Accural    int
	UserID     int
	UploadedAt time.Time
}

type UploadOrderResponse struct{}

type GetOrderByOrderIDRequest struct {
	OrderID int64
}

type GetOrderByOrderIDResponse struct {
	UserID     int
	Status     string
	UploadedAt time.Time
}

type GetAllOrdersFromUserRequest struct {
	UserID int
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
	UserID int
}

type GetOrdersForProcessResponse struct {
	ProcessedOrders []ProcessedOrder
}

type MarkNewWithProcessingRequest struct{}

type MarkNewWithProcessingResponse struct{}

type UpdateOrderRequest struct {
	Number  int
	Status  string
	Accural int
}

type UpdateOrderResponse struct {
}
