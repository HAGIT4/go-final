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
