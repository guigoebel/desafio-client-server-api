package dbsqlite

import (
	"context"
	"database/sql"

	"github.com/guigoebel/desafio-client-server-api/cotation"
)

type SQLite struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *SQLite {
	return &SQLite{
		db: db,
	}
}

func (s *SQLite) Store(ctx context.Context, cotation *cotation.Cotation) (string, error) {
	//insert into table cotation
	//return error if any

	stmt, err := s.db.PrepareContext(ctx, "INSERT INTO cotation (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, cotation.USDBRL.Code, cotation.USDBRL.Codein, cotation.USDBRL.Name, cotation.USDBRL.High, cotation.USDBRL.Low, cotation.USDBRL.VarBid, cotation.USDBRL.PctChange, cotation.USDBRL.Bid, cotation.USDBRL.Ask, cotation.USDBRL.Timestamp, cotation.USDBRL.CreateDate)
	if err != nil {
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return string(id), nil
}
