package config

import (
	"context"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func OpenConnectionDB(log *zap.Logger) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), Env.DSN)
	if err != nil {
		log.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}
	
	return conn
}
