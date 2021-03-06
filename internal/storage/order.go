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
	_, err = st.connection.Exec(ctx, sqlStmt, req.Number, req.Status, req.Accural, req.UserID, req.UploadedAt)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (st *BonusStorage) GetOrderByOrderID(req *modelStorage.GetOrderByOrderIDRequest) (resp *modelStorage.GetOrderByOrderIDResponse, err error) {
	resp = &modelStorage.GetOrderByOrderIDResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `SELECT status, user_id, uploaded_at FROM bonus.order WHERE number=$1`
	sqlResult, err := st.connection.Query(ctx, sqlStmt, req.OrderID)
	if err != nil {
		return nil, err
	}
	defer sqlResult.Close()

	var status string
	var userID int
	var uploadedAt time.Time
	for sqlResult.Next() {
		if err = sqlResult.Scan(&status, &userID, &uploadedAt); err != nil {
			return nil, err
		}
	}
	err = sqlResult.Err()
	if err != nil {
		return nil, err
	}

	resp = &modelStorage.GetOrderByOrderIDResponse{
		UserID:     userID,
		Status:     status,
		UploadedAt: uploadedAt,
	}
	return resp, nil
}

func (st *BonusStorage) GetAllOrdersFromUser(req *modelStorage.GetAllOrdersFromUserRequest) (resp *modelStorage.GetAllOrdersFromUserResponse, err error) {
	resp = &modelStorage.GetAllOrdersFromUserResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `SELECT number, status, accural, uploaded_at FROM bonus.order WHERE user_id=$1
		ORDER BY uploaded_at ASC`
	sqlResult, err := st.connection.Query(ctx, sqlStmt, req.UserID)
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

func (st *BonusStorage) GetOrdersForProcess(req *modelStorage.GetOrdersForProcessRequest) (resp *modelStorage.GetOrdersForProcessResponse, err error) {
	resp = &modelStorage.GetOrdersForProcessResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `SELECT number, user_id FROM bonus.order WHERE status='NEW' OR status='PROCESSING'`
	sqlResult, err := st.connection.Query(ctx, sqlStmt)
	if err != nil {
		return nil, err
	}
	defer sqlResult.Close()

	var orders []modelStorage.ProcessedOrder
	for sqlResult.Next() {
		var orderNumber, orderUserID int
		if err = sqlResult.Scan(&orderNumber, &orderUserID); err != nil {
			return nil, err
		}
		order := modelStorage.ProcessedOrder{
			Number: orderNumber,
			UserID: orderUserID,
		}
		orders = append(orders, order)
	}
	err = sqlResult.Err()
	if err != nil {
		return nil, err
	}
	resp.ProcessedOrders = orders
	return resp, nil
}

func (st *BonusStorage) MarkNewWithProcessing(req *modelStorage.MarkNewWithProcessingRequest) (resp *modelStorage.MarkNewWithProcessingResponse, err error) {
	resp = &modelStorage.MarkNewWithProcessingResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `UPDATE bonus.order SET status = 'PROCESSING' WHERE status = 'NEW'`
	_, err = st.connection.Exec(ctx, sqlStmt)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (st *BonusStorage) UpdateOrder(req *modelStorage.UpdateOrderRequest) (resp *modelStorage.UpdateOrderResponse, err error) {
	resp = &modelStorage.UpdateOrderResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `UPDATE bonus.order SET status = $1, accural = $2 WHERE number = $3`
	_, err = st.connection.Exec(ctx, sqlStmt, req.Status, req.Accural, req.Number)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
