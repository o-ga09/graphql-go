package dao

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/o-ga09/graphql-go/internal/db/db"
	"github.com/o-ga09/graphql-go/internal/domain"
)

func TestNewUserDao(t *testing.T) {
	t.Parallel()
	dbmock, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		d *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *userDao
	}{
		{name: "TestNewUserDao", args: args{d: dbmock}, want: &userDao{query: db.New(dbmock)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewUserDao(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserDao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userDao_GetUsers(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	res := []*domain.User{
		{ID: "1", FirstName: "name1", LastName: "name1", Email: "email1", Address: "address1", Sex: 1, BirthDay: "2024-08-15", Password: "password1", CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC)},
	}

	expected := db.User{
		ID:        1,
		UserID:    "1",
		Name:      "name1 name1",
		Email:     "email1",
		Address:   "address1",
		Birthday:  "2024-08-15",
		Password:  "password1",
		Sex:       1,
		CreatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "address", "email", "password", "sex", "birthday", "created_at", "updated_at"})
	rows.AddRow(expected.ID, expected.UserID, expected.Name, expected.Address, expected.Email, expected.Password, expected.Sex, expected.Birthday, expected.CreatedAt.Time, expected.UpdatedAt.Time)
	mock.ExpectQuery("SELECT id, user_id, name, address, email, password, sex, birthday, created_at, updated_at FROM users ORDER BY created_at DESC").WillReturnRows(rows)

	type fields struct {
		query *db.Queries
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.User
		wantErr bool
	}{
		{name: "TestGetUsers", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx}, want: res, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &userDao{
				query: tt.fields.query,
			}
			got, err := u.GetUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userDao.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userDao.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userDao_GetUserById(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	res := []*domain.User{
		{ID: "1", FirstName: "name1", LastName: "name1", Email: "email1", Address: "address1", Sex: 1, BirthDay: "2024-08-15", Password: "password1", CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC)},
	}

	expected := db.User{
		ID:        1,
		UserID:    "1",
		Name:      "name1 name1",
		Email:     "email1",
		Address:   "address1",
		Birthday:  "2024-08-15",
		Password:  "password1",
		Sex:       1,
		CreatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "address", "email", "password", "sex", "birthday", "created_at", "updated_at"})
	rows.AddRow(expected.ID, expected.UserID, expected.Name, expected.Address, expected.Email, expected.Password, expected.Sex, expected.Birthday, expected.CreatedAt.Time, expected.UpdatedAt.Time)
	mock.ExpectQuery("SELECT id, user_id, name, address, email, password, sex, birthday, created_at, updated_at FROM users WHERE user_id = \\? LIMIT 1").WithArgs("1").WillReturnRows(rows)

	type fields struct {
		query *db.Queries
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{name: "TestGetUserById", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx, id: "1"}, want: res[0], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &userDao{
				query: tt.fields.query,
			}
			got, err := u.GetUserById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userDao.GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userDao.GetUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userDao_CreateUser(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	res := []*domain.User{
		{ID: "1", FirstName: "name1", LastName: "name1", Email: "email1", Address: "address1", Sex: 1, BirthDay: "2024-08-15", Password: "password1", CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC)},
	}

	mock.ExpectExec("INSERT INTO users").WithArgs("1", "name1 name1", "email1", "address1", 1, "2024-08-15", "password1").WillReturnResult(sqlmock.NewResult(1, 1))

	type fields struct {
		query *db.Queries
	}
	type args struct {
		ctx  context.Context
		user *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestCreateUser", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx, user: res[0]}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &userDao{
				query: tt.fields.query,
			}
			if err := u.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userDao.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userDao_UpdateUserById(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	res := []*domain.User{
		{ID: "1", FirstName: "name1", LastName: "name1", Email: "email1", Address: "address1", Sex: 1, BirthDay: "2024-08-15", Password: "password1", CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC)},
	}

	arg := db.UpdateUserParams{
		Name:     "name1 name1",
		Email:    "email1",
		Address:  "address1",
		Sex:      1,
		Birthday: "2024-08-15",
		Password: "password1",
		UserID:   "1",
	}

	mock.ExpectExec("UPDATE users").
		WithArgs(
			arg.Name,
			arg.Email,
			arg.Address,
			arg.Sex,
			arg.Birthday,
			arg.Password,
			arg.UserID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	type fields struct {
		query *db.Queries
	}
	type args struct {
		ctx  context.Context
		id   string
		user *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestUpdateUserById", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx, id: "1", user: res[0]}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &userDao{
				query: tt.fields.query,
			}
			if err := u.UpdateUserById(tt.args.ctx, tt.args.id, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userDao.UpdateUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userDao_DeleteUserById(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectExec("DELETE FROM users").WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1))

	type fields struct {
		query *db.Queries
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestDeleteUserById", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx, id: "1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &userDao{
				query: tt.fields.query,
			}
			if err := u.DeleteUserById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("userDao.DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
