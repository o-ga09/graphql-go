package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"log/slog"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/o-ga09/graphql-go/graph"
	"github.com/o-ga09/graphql-go/internal/db"
	"github.com/o-ga09/graphql-go/internal/db/dao"
	"github.com/o-ga09/graphql-go/internal/service"
	"github.com/o-ga09/graphql-go/pkg/logger"
	"github.com/o-ga09/graphql-go/pkg/middleware"
)

const defaultPort = "8080"

func main() {
	ctx := context.Background()
	logger.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conn, err := db.Connect(ctx)
	if err != nil {
		slog.Log(ctx, logger.SeverityError, "failed to connect db", "err msg", err.Error())
		return
	}

	result, err := conn.Query("SELECT 1")
	if err != nil {
		slog.Log(ctx, logger.SeverityError, "failed to ping db", "err msg", err.Error())
		return
	}
	defer result.Close()

	slog.Log(ctx, logger.SeverityInfo, "connected to db")

	noteRepo := dao.NewNoteDao(conn)
	userRepo := dao.NewUserDao(conn)
	noteService := service.NewNoteService(noteRepo)
	userService := service.NewUserService(userRepo)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		NoteService: *noteService,
		UserService: *userService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.CorsMiddleware(srv))

	slog.Log(ctx, logger.SeverityInfo, fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
