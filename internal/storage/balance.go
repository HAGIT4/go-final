package storage

import (
	"context"
	"time"

	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

func (st *BonusStorage) GetBalanceByUserID(req *modelStorage.GetBalanceByUserIDRequest) (resp *modelStorage.GetBalanceByUserIDResponse, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlResult, err := st.connection.Query(ctx, "SELECT current, withdrawn FROM bonus.balance WHERE user_id=$1", req.UserID)
	if err != nil {
		return nil, err
	}
	defer sqlResult.Close()

	var current, withdrawn int
	var found bool
	for sqlResult.Next() {
		if err = sqlResult.Scan(&current, &withdrawn); err != nil {
			return nil, err
		}
		found = true
	}
	err = sqlResult.Err()
	if err != nil {
		return nil, err
	}

	resp = &modelStorage.GetBalanceByUserIDResponse{
		Balance: modelStorage.Balance{
			Current:   current,
			Withdrawn: withdrawn,
		},
		Found: found,
	}
	return resp, nil
}

func (st *BonusStorage) GetAllWithdrawalsByUserID(req *modelStorage.GetAllWithdrawalsByUserIDRequest) (resp *modelStorage.GetAllWithdrawalsByUserIDResponse, err error) {
	resp = &modelStorage.GetAllWithdrawalsByUserIDResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `SELECT order_id, sum, user_id, processed_at FROM bonus.withdrawal WHERE user_id=$1
		ORDER BY processed_at ASC`
	sqlResult, err := st.connection.Query(ctx, sqlStmt, req.UserID)
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
			OrderID:     orderId,
			Sum:         sum,
			UserID:      userId,
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

func (st *BonusStorage) SetUserBalanceByUserID(req *modelStorage.SetUserBalanceByUserIDRequest) (resp *modelStorage.SetUserBalanceByUserIDResponse, err error) {
	resp = &modelStorage.SetUserBalanceByUserIDResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `UPDATE bonus.balance SET current=$1, withdrawn=$2 WHERE user_id=$3`
	_, err = st.connection.Exec(ctx, sqlStmt, req.Current, req.Withdrawn, req.UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (st *BonusStorage) AddWithdrawal(req *modelStorage.AddWithdrawalRequest) (resp *modelStorage.AddWithdrawalResponse, err error) {
	resp = &modelStorage.AddWithdrawalResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := st.connection.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	sqlStmtBalance := `UPDATE bonus.balance SET current=$1, withdrawn=$2 WHERE user_id=$3`
	_, err = tx.Exec(ctx, sqlStmtBalance, req.Current, req.Withdrawn, req.UserID)
	if err != nil {
		return nil, err
	}

	sqlStmtWithdrawal := `INSERT INTO bonus.withdrawal (order_id, sum, user_id, processed_at)
		VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(ctx, sqlStmtWithdrawal, req.OrderID, req.Sum, req.OrderID, req.ProcessedAt)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (st *BonusStorage) AddUserBalance(req *modelStorage.AddUserBalanceRequest) (resp *modelStorage.AddUserBalanceResponse, err error) {
	resp = &modelStorage.AddUserBalanceResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `INSERT INTO bonus.balance (user_id, current, withdrawn) VALUES ($1, 0, 0)`
	_, err = st.connection.Exec(ctx, sqlStmt, req.UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (st *BonusStorage) AddSumToUserBalance(req *modelStorage.AddSumToUserBalanceRequest) (resp *modelStorage.AddSumToUserBalanceResponse, err error) {
	resp = &modelStorage.AddSumToUserBalanceResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlStmt := `UPDATE bonus.balance SET current=$1 WHERE user_id=$2`
	_, err = st.connection.Exec(ctx, sqlStmt, req.Sum, req.UserID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
