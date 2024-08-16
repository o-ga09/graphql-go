package db

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/o-ga09/graphql-go/pkg/logger"
)

func Connect(ctx context.Context) *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	slog.Log(ctx, logger.SeverityInfo, "db connected")
	return db
}
