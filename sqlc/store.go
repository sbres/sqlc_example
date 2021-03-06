package sqlc

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Store interface {
	Querier

	ExecTx(context.Context, func(Querier) error) error
}

type SQLStore struct {
	db *sql.DB

	*Queries
}

func InitDb(dsn string) (Store, error) {
	conn, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, errors.Wrapf(err, "Problem calling to the database")
	}
	conn.SetMaxOpenConns(0)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(time.Hour)

	err = conn.Ping()
	if err != nil {
		return nil, errors.Wrapf(err, "Problem connecting to the database")
	}
	db := New(conn)
	return &SQLStore{db: conn, Queries: db}, nil
}

// ExecTx will execute the function f inside a transaction
// It will rollback in case something goes wrong
func (s *SQLStore) ExecTx(ctx context.Context, f func(Querier) error) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return err
	}

	qTx := s.Queries.WithTx(tx)

	err = f(qTx)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			log.WithError(rbErr).WithField("txErr", err).Error("Problem rolling back TX")
			return rbErr
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.WithError(err).Error("Problem commiting tx")
		return err
	}

	return nil
}
