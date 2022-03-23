package storage

import (
	"context"

	modelStorage "github.com/HAGIT4/go-final/pkg/storage/model"
)

func (st *BonusStorage) AddUser(req *modelStorage.AddUserRequest) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err = st.connection.Exec(ctx, "INSERT INTO bonus.user(username, passwd_hash) VALUES($1, $2)",
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

	sqlResult, err := st.connection.Query(ctx, "SELECT * FROM bonus.user WHERE username=$1",
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
	if userId == 0 {
		resp = &modelStorage.GetUserByUsernameResponse{
			Found: false,
		}
	} else {
		resp.Found = true // will it work?
	}

	return resp, nil
}

func (st *BonusStorage) GetUserIDByUsername(req *modelStorage.GetUserIdByUsernameRequest) (resp *modelStorage.GetUserIdByUsernameResponse, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlResult, err := st.connection.Query(ctx, "SELECT id FROM bonus.user WHERE username=$1", req.Username)
	if err != nil {
		return nil, err
	}
	defer sqlResult.Close()

	var userId int
	for sqlResult.Next() {
		if err = sqlResult.Scan(&userId); err != nil {
			return nil, err
		}
	}
	err = sqlResult.Err()
	if err != nil {
		return nil, err
	}

	resp = &modelStorage.GetUserIdByUsernameResponse{}
	if userId == 0 {
		resp = &modelStorage.GetUserIdByUsernameResponse{
			UserId: 0,
			Found:  false,
		}
	} else {
		resp = &modelStorage.GetUserIdByUsernameResponse{
			UserId: userId,
			Found:  true,
		}
	}
	return resp, nil
}
