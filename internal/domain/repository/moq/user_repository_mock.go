// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/o-ga09/graphql-go/internal/domain"
	"github.com/o-ga09/graphql-go/internal/domain/repository"
	"sync"
)

// Ensure, that UserRepositoryMock does implement repository.UserRepository.
// If this is not the case, regenerate this file with moq.
var _ repository.UserRepository = &UserRepositoryMock{}

// UserRepositoryMock is a mock implementation of repository.UserRepository.
//
//	func TestSomethingThatUsesUserRepository(t *testing.T) {
//
//		// make and configure a mocked repository.UserRepository
//		mockedUserRepository := &UserRepositoryMock{
//			DeleteFunc: func(ctx context.Context, id string) error {
//				panic("mock out the Delete method")
//			},
//			GetUserByIDFunc: func(ctx context.Context, id string) (*domain.User, error) {
//				panic("mock out the GetUserByID method")
//			},
//			GetUsersFunc: func(ctx context.Context) ([]*domain.User, error) {
//				panic("mock out the GetUsers method")
//			},
//			SaveFunc: func(ctx context.Context, user *domain.User) error {
//				panic("mock out the Save method")
//			},
//		}
//
//		// use mockedUserRepository in code that requires repository.UserRepository
//		// and then make assertions.
//
//	}
type UserRepositoryMock struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, id string) error

	// GetUserByIDFunc mocks the GetUserByID method.
	GetUserByIDFunc func(ctx context.Context, id string) (*domain.User, error)

	// GetUsersFunc mocks the GetUsers method.
	GetUsersFunc func(ctx context.Context) ([]*domain.User, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(ctx context.Context, user *domain.User) error

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// GetUserByID holds details about calls to the GetUserByID method.
		GetUserByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// GetUsers holds details about calls to the GetUsers method.
		GetUsers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// User is the user argument value.
			User *domain.User
		}
	}
	lockDelete      sync.RWMutex
	lockGetUserByID sync.RWMutex
	lockGetUsers    sync.RWMutex
	lockSave        sync.RWMutex
}

// Delete calls DeleteFunc.
func (mock *UserRepositoryMock) Delete(ctx context.Context, id string) error {
	if mock.DeleteFunc == nil {
		panic("UserRepositoryMock.DeleteFunc: method is nil but UserRepository.Delete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedUserRepository.DeleteCalls())
func (mock *UserRepositoryMock) DeleteCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// GetUserByID calls GetUserByIDFunc.
func (mock *UserRepositoryMock) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	if mock.GetUserByIDFunc == nil {
		panic("UserRepositoryMock.GetUserByIDFunc: method is nil but UserRepository.GetUserByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetUserByID.Lock()
	mock.calls.GetUserByID = append(mock.calls.GetUserByID, callInfo)
	mock.lockGetUserByID.Unlock()
	return mock.GetUserByIDFunc(ctx, id)
}

// GetUserByIDCalls gets all the calls that were made to GetUserByID.
// Check the length with:
//
//	len(mockedUserRepository.GetUserByIDCalls())
func (mock *UserRepositoryMock) GetUserByIDCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetUserByID.RLock()
	calls = mock.calls.GetUserByID
	mock.lockGetUserByID.RUnlock()
	return calls
}

// GetUsers calls GetUsersFunc.
func (mock *UserRepositoryMock) GetUsers(ctx context.Context) ([]*domain.User, error) {
	if mock.GetUsersFunc == nil {
		panic("UserRepositoryMock.GetUsersFunc: method is nil but UserRepository.GetUsers was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetUsers.Lock()
	mock.calls.GetUsers = append(mock.calls.GetUsers, callInfo)
	mock.lockGetUsers.Unlock()
	return mock.GetUsersFunc(ctx)
}

// GetUsersCalls gets all the calls that were made to GetUsers.
// Check the length with:
//
//	len(mockedUserRepository.GetUsersCalls())
func (mock *UserRepositoryMock) GetUsersCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetUsers.RLock()
	calls = mock.calls.GetUsers
	mock.lockGetUsers.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *UserRepositoryMock) Save(ctx context.Context, user *domain.User) error {
	if mock.SaveFunc == nil {
		panic("UserRepositoryMock.SaveFunc: method is nil but UserRepository.Save was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		User *domain.User
	}{
		Ctx:  ctx,
		User: user,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	return mock.SaveFunc(ctx, user)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//
//	len(mockedUserRepository.SaveCalls())
func (mock *UserRepositoryMock) SaveCalls() []struct {
	Ctx  context.Context
	User *domain.User
} {
	var calls []struct {
		Ctx  context.Context
		User *domain.User
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}
