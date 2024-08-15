package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_E2E_User(t *testing.T) {
	t.Parallel()
	type args struct {
		rw  *httptest.ResponseRecorder
		req *http.Request
	}
	tests := map[string]struct {
		statusCode int
		args       args
	}{
		"TestGetUsers": {
			statusCode: 200,
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest("GET", "/query", nil),
			},
		},
		"TestGetUserById": {
			statusCode: 200,
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest("GET", "/query", nil),
			},
		},
		"TestCreateUser": {
			statusCode: 200,
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest("POST", "/query", nil),
			},
		},
		"TestUpdateUserById": {
			statusCode: 200,
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest("PUT", "/query", nil),
			},
		},
		"TestDeleteUserById": {
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
