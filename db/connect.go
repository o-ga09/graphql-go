package db

import (
	"context"
	"database/sql"
	"fmt"
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
	fmt.Println("called")
	slog.Log(ctx, logger.SeverityInfo, "db connected")
	return db
}
