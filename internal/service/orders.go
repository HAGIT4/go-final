package service

import (
	"time"

	pkgService "github.com/HAGIT4/go-final/pkg/service"
	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

func (sv *BonusService) UploadOrder(req *pkgService.UploadOrderRequest) (resp *pkgService.UploadOrderResponse) {
	resp = &pkgService.UploadOrderResponse{}

	// TODO: Add order number validate (Luhn)

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
		if ordersUserId != int(req.Order) {
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

func (sv *BonusService) getOrderByOrderId(orderId int64) (ordersUserId int, orderFound bool, err error) {
	dbReq := &modelStorage.GetOrderByOrderIdRequest{
		OrderId: int(orderId),
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
