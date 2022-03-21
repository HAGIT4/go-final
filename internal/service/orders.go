package service

import (
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

	userId, userFound, err := sv.getUserIdByUsername(req.Username)
	if err != nil {
		resp.Status = pkgService.UploadOrderResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !userFound {
		resp.Status = pkgService.UploadOrderResponse_UNAUTHORIZED
		return resp
	}

	ordersUserId, orderFound, err := sv.getOrderByOrderId(req.Order)
	if err != nil {
		resp.Status = pkgService.UploadOrderResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if orderFound {
		if ordersUserId != userId {
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
		UserId:     userId,
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
	userId, userFound, err := sv.getUserIdByUsername(req.Username)
	if err != nil {
		resp.Status = pkgService.GetOrderListResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if !userFound {
		resp.Status = pkgService.GetOrderListResponse_UNAUTHORIZED
		return resp
	}

	orders, err := sv.getOrderListByUser(userId)
	if err != nil {
		resp.Status = pkgService.GetOrderListResponse_INTERNAL_SERVER_ERROR
		return resp
	}
	if len(orders) == 0 {
		resp.Status = pkgService.GetOrderListResponse_NO_DATA
		return resp
	}
	resp.OrderInfo = orders
	resp.Status = pkgService.GetOrderListResponse_OK
	return resp
}

func (sv *BonusService) getOrderListByUser(userId int) (orders []*pkgService.OrderInfo, err error) {
	dbReq := &modelStorage.GetAllOrdersFromUserRequest{
		UserId: userId,
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
			Accural:    newAccural,
			UploadedAt: newTime,
		}
		orders = append(orders, &newOrder)
	}
	return orders, nil
}

func (sv *BonusService) getOrderByOrderId(orderId int64) (ordersUserId int, orderFound bool, err error) {
	dbReq := &modelStorage.GetOrderByOrderIdRequest{
		OrderId: int64(orderId),
	}
	dbResp, err := sv.storage.GetOrderByOrderId(dbReq)
	if err != nil {
		return 0, false, err
	}
	if dbResp.UserId != 0 {
		orderFound = true
	} else {
		orderFound = false
	}
	return dbResp.UserId, orderFound, nil

}
