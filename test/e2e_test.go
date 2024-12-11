package test

import (
	"context"
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/o-ga09/graphql-go/graph"
	"github.com/o-ga09/graphql-go/internal/db"
	"github.com/o-ga09/graphql-go/internal/db/dao"

	"github.com/o-ga09/graphql-go/internal/service"
	"github.com/o-ga09/graphql-go/test/testutil"
	"github.com/tenntenn/golden"
)

var (
	flagUpdate bool
	goldenDir  string = "./testdata/golden/"
)

func init() {
	flag.BoolVar(&flagUpdate, "update", false, "update golden files")
}

func resetDatabase(t *testing.T) {
	t.Helper()
	testutil.SetUpMySQL(t)
}

func Test_E2E_User(t *testing.T) {
	t.Setenv("ENV", "test")
	ctx := context.Background()
	tests := map[string]struct {
		statusCode int
	}{
		"TestGetUsers": {
			statusCode: 200,
		},
		"TestGetUserById": {
			statusCode: 200,
		},
		"TestCreateUser": {
			statusCode: 200,
		},
		"TestUpdateUserById": {
			statusCode: 200,
		},
		"TestDeleteUserById": {
			statusCode: 200,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			resetDatabase(t)

			conn, _ := db.Connect(ctx)
			noteRepo := dao.NewNoteDao(conn)
			userRepo := dao.NewUserDao(conn)
			noteService := service.NewNoteService(noteRepo)
			userService := service.NewUserService(userRepo)
			srv := httptest.NewServer(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
				NoteService: *noteService,
				UserService: *userService,
			}})),
			)
			t.Cleanup(func() { srv.Close() })

			reqBody := testutil.GetRequestBody(t, goldenDir, "user/"+name+"In.gpl")
			req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, srv.URL, reqBody)
			if err != nil {
				t.Fatal("error new request", err)
			}
			req.Header.Add("Content-Type", "application/json")
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal("error request", err)
			}
			t.Cleanup(func() { res.Body.Close() })

			if res.StatusCode != tt.statusCode {
				t.Errorf("status code = %v, want %v", res.StatusCode, tt.statusCode)
			}
			got := testutil.GetResponseBody(t, res)
			if diff := golden.Check(t, flagUpdate, goldenDir, "user/"+name+"Out.json", got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
				t.Log(got)
			}
		})
	}
}

func Test_E2E_Note(t *testing.T) {
	t.Setenv("ENV", "test")
	ctx := context.Background()
	tests := map[string]struct {
		statusCode int
	}{
		"TestGetNotes": {
			statusCode: 200,
		},
		"TestGetNoteAll": {
			statusCode: 200,
		},
		"TestGetNoteById": {
			statusCode: 200,
		},
		"TestCreateNote": {
			statusCode: 200,
		},
		"TestUpdateNoteById": {
			statusCode: 200,
		},
		"TestDeleteNoteById": {
			statusCode: 200,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			resetDatabase(t)

			conn, _ := db.Connect(ctx)
			noteRepo := dao.NewNoteDao(conn)
			userRepo := dao.NewUserDao(conn)
			noteService := service.NewNoteService(noteRepo)
			userService := service.NewUserService(userRepo)
			srv := httptest.NewServer(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
				NoteService: *noteService,
				UserService: *userService,
			}})),
			)
			t.Cleanup(func() { srv.Close() })

			reqBody := testutil.GetRequestBody(t, goldenDir, "note/"+name+"In.gpl")
			req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, srv.URL, reqBody)
			if err != nil {
				t.Fatal("error new request", err)
			}

			req.Header.Add("Content-Type", "application/json")
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal("error request", err)
			}
			t.Cleanup(func() { res.Body.Close() })
			if res.StatusCode != tt.statusCode {
				t.Errorf("status code = %v, want %v", res.StatusCode, tt.statusCode)
			}
			got := testutil.GetResponseBody(t, res)
			if diff := golden.Check(t, flagUpdate, goldenDir, "note/"+name+"Out.json", got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
				t.Log(got)
			}
		})
	}
}
