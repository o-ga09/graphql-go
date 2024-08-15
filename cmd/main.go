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
	"github.com/o-ga09/graphql-go/db"
	"github.com/o-ga09/graphql-go/db/dao"
	"github.com/o-ga09/graphql-go/graph"
	"github.com/o-ga09/graphql-go/pkg/logger"
	"github.com/o-ga09/graphql-go/service"
)

const defaultPort = "8080"

func main() {
	ctx := context.Background()
	logger.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conn := db.Connect(ctx)
	if conn == nil {
		slog.Log(ctx, logger.SeverityError, "failed to connect db")
		return
	}
	noteRepo := dao.NewNoteDao(conn)
	userRepo := dao.NewUserDao(conn)
	noteService := service.NewNoteService(noteRepo)
	userService := service.NewUserService(userRepo)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		NoteService: *noteService,
		UserService: *userService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	slog.Log(ctx, logger.SeverityInfo, fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
