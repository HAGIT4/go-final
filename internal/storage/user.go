package storage

import (
	"context"

	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

func (st *BonusStorage) AddUser(req *modelStorage.AddUserRequest) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err = st.connection.Exec(ctx, "INSERT INTO user(username, passwdHash) VALUES($1, $2)",
		req.Username, req.PasswdHash,
	)
	if err != nil {
		return err
	}
	return nil
}

func (st *BonusStorage) GetUserByUsername(req *modelStorage.GetUserByUsernameRequest) (resp *modelStorage.GetUserByUsernameResponse, err error) {
	resp = &modelStorage.GetUserByUsernameResponse{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlResult, err := st.connection.Query(ctx, "SELECT * FROM user WHERE username=$1",
		req.Username,
	)
	if err != nil {
		return nil, err
	}
	defer sqlResult.Close()

	var userId int64
	for sqlResult.Next() {
		if err = sqlResult.Scan(&userId, &resp.Username, &resp.PasswdHash); err != nil {
			return nil, err
		}
	}
	resp.Found = true
	return resp, nil
}
