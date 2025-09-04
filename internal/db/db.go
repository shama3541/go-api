package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr string, maxopenConns, maxIdleConns int, maxIdletime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxopenConns)
	db.SetMaxIdleConns(maxIdleConns)
	duration, _ := time.ParseDuration(maxIdletime)
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
