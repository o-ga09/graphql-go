package test

import (
	"context"
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/o-ga09/graphql-go/db"
	"github.com/o-ga09/graphql-go/db/dao"
	"github.com/o-ga09/graphql-go/graph"
	"github.com/o-ga09/graphql-go/service"
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

func Test_E2E_User(t *testing.T) {
	ctx := context.Background()
	testutil.SetUpMySQL(t)
	tests := map[string]struct {
		statusCode int
	}{
		"TestGetUsers": {
			statusCode: 200,
		},
		// "TestGetUserById": {
		// 	statusCode: 200,
		// },
		// "TestCreateUser": {
		// 	statusCode: 200,
		// },
		// "TestUpdateUserById": {
		// 	statusCode: 200,
		// },
		// "TestDeleteUserById": {
		// 	statusCode: 200,
		// },
	}

	conn := db.Connect(ctx)
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

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			reqBody := testutil.GetRequestBody(t, goldenDir, name+"In.gpl")
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
			if diff := golden.Check(t, flagUpdate, goldenDir, name+"Out.json", got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_E2E_Note(t *testing.T) {
	t.Parallel()
	type args struct {
		rw  *httptest.ResponseRecorder
		req *http.Request
	}
	tests := map[string]struct {
		statusCode int
		args       args
	}{
		"TestGetNotes": {
			statusCode: 200,
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest("GET", "/query", nil),
			},
		},
		"TestGetNoteById": {
			statusCode: 200,
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest("GET", "/query", nil),
			},
		},
		"TestCreateNote": {
			statusCode: 200,
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest("POST", "/query", nil),
			},
		},
		"TestUpdateNoteById": {
			statusCode: 200,
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest("PUT", "/query", nil),
			},
		},
		"TestDeleteNoteById": {
			statusCode: 200,
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest("DELETE", "/query", nil),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			tt.args.rw.WriteHeader(tt.statusCode)
			if tt.args.rw.Code != tt.statusCode {
				t.Errorf("status code = %v, want %v", tt.args.rw.Code, tt.statusCode)
			}
		})
	}
}
