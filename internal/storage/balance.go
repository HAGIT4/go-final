package storage

import (
	"context"

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
