package storage

import (
	"context"
	"time"

	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

func (st *BonusStorage) GetBalanceByUserId(req *modelStorage.GetBalanceByUserIdRequest) (resp *modelStorage.GetBalanceByUserIdResponse, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlResult, err := st.connection.Query(ctx, "SELECT current, withdrawn FROM bonus.balance WHERE user_id=$1", req.UserId)
	if err != nil {
		return nil, err
	}
	defer sqlResult.Close()

	var current, withdrawn int
	for sqlResult.Next() {
		if err = sqlResult.Scan(&current, &withdrawn); err != nil {
			return nil, err
		}
	}
	err = sqlResult.Err()
	if err != nil {
		return nil, err
	}

	resp = &modelStorage.GetBalanceByUserIdResponse{
		Balance: modelStorage.Balance{
			Current:   current,
			Withdrawn: withdrawn,
		},
		Found: true,
	}
	return resp, nil
}

func (st *BonusStorage) GetAllWithdrawalsByUserId(req *modelStorage.GetAllWithdrawalsByUserIdRequest) (resp *modelStorage.GetAllWithdrawalsByUserIdResponse, err error) {
	resp = &modelStorage.GetAllWithdrawalsByUserIdResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `SELECT order_id, sum, user_id, processed_at FROM bonus.withdrawal WHERE user_id=$1`
	sqlResult, err := st.connection.Query(ctx, sqlStmt, req.UserId)
	if err != nil {
		return nil, err
	}
	defer sqlResult.Close()

	var orderId int
	var sum int
	var userId int
	var processedAt time.Time

	var withdrawals []modelStorage.Withdrawal
	for sqlResult.Next() {
		if err = sqlResult.Scan(&orderId, &sum, &userId, &processedAt); err != nil {
			return nil, err
		}
		withdrawal := modelStorage.Withdrawal{
			OrderId:     orderId,
			Sum:         sum,
			UserId:      userId,
			ProcessedAt: processedAt,
		}
		withdrawals = append(withdrawals, withdrawal)
	}
	if err = sqlResult.Err(); err != nil {
		return nil, err
	}

	resp.Withdrawals = withdrawals
	return resp, nil
}
