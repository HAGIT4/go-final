package storage

import (
	"context"

	pkgStorage "github.com/HAGIT4/go-final/pkg/storage"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type BonusStorage struct {
	connectionString string
	connection       *pgxpool.Pool
}

var _ BonusStorageInterface = (*BonusStorage)(nil)

func NewBonusStorage(cfg *pkgStorage.BonusStorageConfig) (st *BonusStorage, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	connCfg, err := pgxpool.ParseConfig(cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.ConnectConfig(ctx, connCfg)
	if err != nil {
		return nil, err
	}

	// conn, err := pgx.ConnectConfig(ctx, connCfg)
	// if err != nil {
	// 	return nil, err
	// }

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
		user_id BIGINT UNIQUE NOT NULL,
		current BIGINT NOT NULL,
		withdrawn BIGINT NOT NULL
	)`)
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(ctx, `CREATE TYPE order_status as ENUM ('NEW', 'PROCESSING', 'INVALID', 'PROCESSED');
		CREATE TABLE IF NOT EXISTS bonus.order (
		id SERIAL PRIMARY KEY,
		number BIGINT UNIQUE NOT NULL,
		status order_status NOT NULL,
		accural BIGINT,
		user_id BIGINT NOT NULL,
		uploaded_at TIMESTAMPTZ NOT NULL
	)`)

	_, err = conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS bonus.withdrawal (
		id SERIAL PRIMARY KEY,
		order_id BIGINT UNIQUE NOT NULL,
		sum BIGINT NOT NULL,
		user_id BIGINT NOT NULL,
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
