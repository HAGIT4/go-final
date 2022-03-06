package storage

import (
	"context"

	storageModels "github.com/HAGIT4/go-final/pkg/storage/models"
)

func (st *BonusStorage) AddUser(req *storageModels.AddUserRequest) (err error) {
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

func (st *BonusStorage) GetUserByUsername(req *storageModels.GetUserByUsernameRequest) (resp *storageModels.GetUserByUsernameResponse, err error) {
	resp = &storageModels.GetUserByUsernameResponse{}
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
