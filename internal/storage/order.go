package storage

import (
	"context"
	"time"

	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

func (st *BonusStorage) UploadOrder(req *modelStorage.UploadOrderRequest) (resp *modelStorage.UploadOrderResponse, err error) {
	resp = &modelStorage.UploadOrderResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `INSERT INTO bonus.order(number, status, accural, user_id, uploaded_at) VALUES (
		$1, $2, $3, $4, $5)`
	_, err = st.connection.Exec(ctx, sqlStmt, req.Number, req.Status, req.Accural, req.UserId, req.UploadedAt)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (st *BonusStorage) GetOrderByOrderId(req *modelStorage.GetOrderByOrderIdRequest) (resp *modelStorage.GetOrderByOrderIdResponse, err error) {
	resp = &modelStorage.GetOrderByOrderIdResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `SELECT (status, user_id, uploaded_at) FROM bonus.order WHERE number=$1`
	sqlResult, err := st.connection.Query(ctx, sqlStmt, req.OrderId)
	if err != nil {
		return nil, err
	}
	defer sqlResult.Close()

	var status string
	var userId int
	var uploadedAt time.Time
	for sqlResult.Next() {
		if err = sqlResult.Scan(&status, &userId, &uploadedAt); err != nil {
			return nil, err
		}
	}
	err = sqlResult.Err()
	if err != nil {
		return nil, err
	}

	resp = &modelStorage.GetOrderByOrderIdResponse{
		UserId:     userId,
		Status:     status,
		UploadedAt: uploadedAt,
	}
	return resp, nil
}

func (st *BonusStorage) GetAllOrdersFromUser(req *modelStorage.GetAllOrdersFromUserRequest) (resp *modelStorage.GetAllOrdersFromUserResponse, err error) {
	resp = &modelStorage.GetAllOrdersFromUserResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `SELECT number, status, accural, uploaded_at FROM bonus.order WHERE user_id=$1`
	sqlResult, err := st.connection.Query(ctx, sqlStmt, req.UserId)
	if err != nil {
		return nil, err
	}
	defer sqlResult.Close()

	var orders []modelStorage.Order
	for sqlResult.Next() {
		var number int
		var status string
		var accural int
		var uploadedAt time.Time
		if err = sqlResult.Scan(&number, &status, &accural, &uploadedAt); err != nil {
			return nil, err
		}
		order := modelStorage.Order{
			Number:     number,
			Status:     status,
			Accural:    accural,
			UploadedAt: uploadedAt,
		}
		orders = append(orders, order)
	}
	err = sqlResult.Err()
	if err != nil {
		return nil, err
	}
	resp.Orders = orders
	return resp, nil
}
