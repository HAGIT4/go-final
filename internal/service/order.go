package service

import (
	"fmt"
	"strconv"
	"time"

	pkgService "github.com/HAGIT4/go-final/pkg/service"
	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
	goluhn "github.com/ShiraazMoollatjie/goluhn"
)

func (sv *BonusService) UploadOrder(req *pkgService.UploadOrderRequest) (resp *pkgService.UploadOrderResponse) {
	resp = &pkgService.UploadOrderResponse{}

	order := strconv.FormatInt(req.Order, 10)
	err := goluhn.Validate(order)
	if err != nil {
		resp.Status = pkgService.UploadOrderResponse_BAD_ORDER_NUMBER
		return resp
	}

	userID, userFound, err := sv.getUserIDByUsername(req.Username)
	if err != nil {
		resp.Status = pkgService.UploadOrderResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !userFound {
		resp.Status = pkgService.UploadOrderResponse_UNAUTHORIZED
		return resp
	}

	ordersUserID, orderFound, err := sv.getOrderByOrderID(req.Order)
	if err != nil {
		resp.Status = pkgService.UploadOrderResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if orderFound {
		if ordersUserID != userID {
			resp.Status = pkgService.UploadOrderResponse_ALREADY_UPLOADED_BY_ANOTHER_USER
			return resp
		} else {
			resp.Status = pkgService.UploadOrderResponse_ALREADY_UPLOADED_BY_THIS_USER
			return resp
		}
	}

	dbReq := &modelStorage.UploadOrderRequest{
		Number:     int(req.Order),
		Status:     "NEW",
		Accural:    0,
		UserID:     userID,
		UploadedAt: time.Now(),
	}
	_, err = sv.storage.UploadOrder(dbReq)
	if err != nil {
		resp.Status = pkgService.UploadOrderResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	resp.Status = pkgService.UploadOrderResponse_OK
	return resp
}

func (sv *BonusService) GetAllOrdersFromUser(req *pkgService.GetOrderListRequest) (resp *pkgService.GetOrderListResponse) {
	resp = &pkgService.GetOrderListResponse{}
	userID, userFound, err := sv.getUserIDByUsername(req.Username)
	if err != nil {
		resp.Status = pkgService.GetOrderListResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !userFound {
		resp.Status = pkgService.GetOrderListResponse_UNAUTHORIZED
		return resp
	}

	orders, err := sv.getOrderListByUser(userID)
	if err != nil {
		resp.Status = pkgService.GetOrderListResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if len(orders) == 0 {
		resp.Status = pkgService.GetOrderListResponse_NO_DATA
		return resp
	}
	resp.OrderInfo = orders
	fmt.Println(orders)
	resp.Status = pkgService.GetOrderListResponse_OK
	return resp
}

func (sv *BonusService) getOrderListByUser(userID int) (orders []*pkgService.OrderInfo, err error) {
	dbReq := &modelStorage.GetAllOrdersFromUserRequest{
		UserID: userID,
	}
	dbResp, err := sv.storage.GetAllOrdersFromUser(dbReq)
	if err != nil {
		return nil, err
	}
	orders = []*pkgService.OrderInfo{}
	for _, order := range dbResp.Orders {
		newNumber := strconv.Itoa(order.Number)
		newAccural := float32(order.Accural) / 100
		newTime := order.UploadedAt.Format(time.RFC3339)
		newOrder := pkgService.OrderInfo{
			Number:     newNumber,
			Status:     order.Status,
			Accrual:    newAccural,
			UploadedAt: newTime,
		}
		orders = append(orders, &newOrder)
	}
	return orders, nil
}

func (sv *BonusService) getOrderByOrderID(orderID int64) (ordersUserID int, orderFound bool, err error) {
	dbReq := &modelStorage.GetOrderByOrderIDRequest{
		OrderID: int64(orderID),
	}
	dbResp, err := sv.storage.GetOrderByOrderID(dbReq)
	if err != nil {
		return 0, false, err
	}
	if dbResp.UserID != 0 {
		orderFound = true
	} else {
		orderFound = false
	}
	return dbResp.UserID, orderFound, nil

}
