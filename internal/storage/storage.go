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

	_, err = conn.Exec(ctx, "CREATE SCHEMA IF NOT EXISTS bonus")
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS bonus.user (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		passwd_hash TEXT NOT NULL
		)`)
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS bonus.balance (
		id SERIAL PRIMARY KEY,
		user_id INTEGER UNIQUE,
		current INTEGER,
		withdrawn INTEGER
	)`)
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(ctx, `CREATE TYPE order_status as ENUM ('NEW', 'PROCESSING', 'INVALID', 'PROCESSED');
		CREATE TABLE IF NOT EXISTS bonus.order (
		id SERIAL PRIMARY KEY,
		number INTEGER UNIQUE NOT NULL,
		status order_status NOT NULL,
		user_id INTEGER,
		uploaded_at TIMESTAMPTZ NOT NULL
	)`)

	_, err = conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS bonus.withdrawal (
		id SERIAL PRIMARY KEY,
		order_id INTEGER UNIQUE NOT NULL,
		sum INTEGER NOT NULL,
		user_id INTEGER,
		processed_at TIMESTAMPTZ NOT NULL
	)`)
	if err != nil {
		return nil, err
	}

	st = &BonusStorage{
		connectionString: cfg.ConnectionString,
		connection:       conn,
	}
	return st, nil
}
