package storage

import (
	storageModels "github.com/HAGIT4/go-final/pkg/storage/models"
)

type BonusStorageInterface interface {
	AddUser(req *storageModels.AddUserRequest) (err error)
	GetUserByUsername(req *storageModels.GetUserByUsernameRequest) (resp *storageModels.GetUserByUsernameResponse, err error)
}
