package config

import (
	"context"
	database "github.com/MuhamadAndre/auth-service/internal/db"
	"github.com/MuhamadAndre/auth-service/internal/utils"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func OpenConnectionDB(log *zap.Logger) *database.Queries {
	conn, err := pgx.Connect(context.Background(), utils.Env.DSN)
	if err != nil {
		log.Error("Error connecting to database", zap.Error(err))
		panic(err)
	}

	return database.New(conn)
}
