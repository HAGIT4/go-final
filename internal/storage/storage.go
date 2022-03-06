package storage

import (
	"context"

	pkgStorage "github.com/HAGIT4/go-final/pkg/storage"
	pgx "github.com/jackc/pgx/v4"
)

type BonusStorage struct {
	connectionString string
	connection       *pgx.Conn
}

var _ BonusStorageInterface = (*BonusStorage)(nil)

func NewBonusStorage(cfg *pkgStorage.BonusStorageConfig) (st *BonusStorage, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	connCfg, err := pgx.ParseConfig(cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(ctx, connCfg)
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS user (id SERIAL PRIMARY KEY, username TEXT, passwdHash)") // add unique
	if err != nil {
		return nil, err
	}

	st = &BonusStorage{
		connectionString: cfg.ConnectionString,
		connection:       conn,
	}
	return st, nil
}
