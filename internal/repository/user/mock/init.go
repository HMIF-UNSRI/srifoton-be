// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	"github.com/HMIF-UNSRI/srifoton-be/internal/repository/user"
	"sync"
)

// Ensure, that RepositoryMock does implement user.Repository.
// If this is not the case, regenerate this file with moq.
var _ user.Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of user.Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked user.Repository
// 		mockedRepository := &RepositoryMock{
// 			FindAllFunc: func(ctx context.Context) (userDomain.User, error) {
// 				panic("mock out the FindAll method")
// 			},
// 			FindByEmailFunc: func(ctx context.Context, email string) (userDomain.User, error) {
// 				panic("mock out the FindByEmail method")
// 			},
// 			FindByIDFunc: func(ctx context.Context, id string) (userDomain.User, error) {
// 				panic("mock out the FindByID method")
// 			},
// 			InsertFunc: func(ctx context.Context, user userDomain.User) (string, error) {
// 				panic("mock out the Insert method")
// 			},
// 			InsertFileFunc: func(ctx context.Context) (string, error) {
// 				panic("mock out the InsertFile method")
// 			},
// 			UpdateVerifiedEmailFunc: func(ctx context.Context, id string) (string, error) {
// 				panic("mock out the UpdateVerifiedEmail method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires user.Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// FindAllFunc mocks the FindAll method.
	FindAllFunc func(ctx context.Context) (userDomain.User, error)

	// FindByEmailFunc mocks the FindByEmail method.
	FindByEmailFunc func(ctx context.Context, email string) (userDomain.User, error)

	// FindByIDFunc mocks the FindByID method.
	FindByIDFunc func(ctx context.Context, id string) (userDomain.User, error)

	// InsertFunc mocks the Insert method.
	InsertFunc func(ctx context.Context, user userDomain.User) (string, error)

	// InsertFileFunc mocks the InsertFile method.
	InsertFileFunc func(ctx context.Context) (string, error)

	// UpdateVerifiedEmailFunc mocks the UpdateVerifiedEmail method.
	UpdateVerifiedEmailFunc func(ctx context.Context, id string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// FindAll holds details about calls to the FindAll method.
		FindAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// FindByEmail holds details about calls to the FindByEmail method.
		FindByEmail []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Email is the email argument value.
			Email string
		}
		// FindByID holds details about calls to the FindByID method.
		FindByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// Insert holds details about calls to the Insert method.
		Insert []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// User is the user argument value.
			User userDomain.User
		}
		// InsertFile holds details about calls to the InsertFile method.
		InsertFile []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// UpdateVerifiedEmail holds details about calls to the UpdateVerifiedEmail method.
		UpdateVerifiedEmail []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
	}
	lockFindAll             sync.RWMutex
	lockFindByEmail         sync.RWMutex
	lockFindByID            sync.RWMutex
	lockInsert              sync.RWMutex
	lockInsertFile          sync.RWMutex
	lockUpdateVerifiedEmail sync.RWMutex
}

// FindAll calls FindAllFunc.
func (mock *RepositoryMock) FindAll(ctx context.Context) (userDomain.User, error) {
	if mock.FindAllFunc == nil {
		panic("RepositoryMock.FindAllFunc: method is nil but Repository.FindAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockFindAll.Lock()
	mock.calls.FindAll = append(mock.calls.FindAll, callInfo)
	mock.lockFindAll.Unlock()
	return mock.FindAllFunc(ctx)
}

// FindAllCalls gets all the calls that were made to FindAll.
// Check the length with:
//     len(mockedRepository.FindAllCalls())
func (mock *RepositoryMock) FindAllCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockFindAll.RLock()
	calls = mock.calls.FindAll
	mock.lockFindAll.RUnlock()
	return calls
}

// FindByEmail calls FindByEmailFunc.
func (mock *RepositoryMock) FindByEmail(ctx context.Context, email string) (userDomain.User, error) {
	if mock.FindByEmailFunc == nil {
		panic("RepositoryMock.FindByEmailFunc: method is nil but Repository.FindByEmail was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Email string
	}{
		Ctx:   ctx,
		Email: email,
	}
	mock.lockFindByEmail.Lock()
	mock.calls.FindByEmail = append(mock.calls.FindByEmail, callInfo)
	mock.lockFindByEmail.Unlock()
	return mock.FindByEmailFunc(ctx, email)
}

// FindByEmailCalls gets all the calls that were made to FindByEmail.
// Check the length with:
//     len(mockedRepository.FindByEmailCalls())
func (mock *RepositoryMock) FindByEmailCalls() []struct {
	Ctx   context.Context
	Email string
} {
	var calls []struct {
		Ctx   context.Context
		Email string
	}
	mock.lockFindByEmail.RLock()
	calls = mock.calls.FindByEmail
	mock.lockFindByEmail.RUnlock()
	return calls
}

// FindByID calls FindByIDFunc.
func (mock *RepositoryMock) FindByID(ctx context.Context, id string) (userDomain.User, error) {
	if mock.FindByIDFunc == nil {
		panic("RepositoryMock.FindByIDFunc: method is nil but Repository.FindByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFindByID.Lock()
	mock.calls.FindByID = append(mock.calls.FindByID, callInfo)
	mock.lockFindByID.Unlock()
	return mock.FindByIDFunc(ctx, id)
}

// FindByIDCalls gets all the calls that were made to FindByID.
// Check the length with:
//     len(mockedRepository.FindByIDCalls())
func (mock *RepositoryMock) FindByIDCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockFindByID.RLock()
	calls = mock.calls.FindByID
	mock.lockFindByID.RUnlock()
	return calls
}

// Insert calls InsertFunc.
func (mock *RepositoryMock) Insert(ctx context.Context, user userDomain.User) (string, error) {
	if mock.InsertFunc == nil {
		panic("RepositoryMock.InsertFunc: method is nil but Repository.Insert was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		User userDomain.User
	}{
		Ctx:  ctx,
		User: user,
	}
	mock.lockInsert.Lock()
	mock.calls.Insert = append(mock.calls.Insert, callInfo)
	mock.lockInsert.Unlock()
	return mock.InsertFunc(ctx, user)
}

// InsertCalls gets all the calls that were made to Insert.
// Check the length with:
//     len(mockedRepository.InsertCalls())
func (mock *RepositoryMock) InsertCalls() []struct {
	Ctx  context.Context
	User userDomain.User
} {
	var calls []struct {
		Ctx  context.Context
		User userDomain.User
	}
	mock.lockInsert.RLock()
	calls = mock.calls.Insert
	mock.lockInsert.RUnlock()
	return calls
}

// InsertFile calls InsertFileFunc.
func (mock *RepositoryMock) InsertFile(ctx context.Context) (string, error) {
	if mock.InsertFileFunc == nil {
		panic("RepositoryMock.InsertFileFunc: method is nil but Repository.InsertFile was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockInsertFile.Lock()
	mock.calls.InsertFile = append(mock.calls.InsertFile, callInfo)
	mock.lockInsertFile.Unlock()
	return mock.InsertFileFunc(ctx)
}

// InsertFileCalls gets all the calls that were made to InsertFile.
// Check the length with:
//     len(mockedRepository.InsertFileCalls())
func (mock *RepositoryMock) InsertFileCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockInsertFile.RLock()
	calls = mock.calls.InsertFile
	mock.lockInsertFile.RUnlock()
	return calls
}

// UpdateVerifiedEmail calls UpdateVerifiedEmailFunc.
func (mock *RepositoryMock) UpdateVerifiedEmail(ctx context.Context, id string) (string, error) {
	if mock.UpdateVerifiedEmailFunc == nil {
		panic("RepositoryMock.UpdateVerifiedEmailFunc: method is nil but Repository.UpdateVerifiedEmail was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockUpdateVerifiedEmail.Lock()
	mock.calls.UpdateVerifiedEmail = append(mock.calls.UpdateVerifiedEmail, callInfo)
	mock.lockUpdateVerifiedEmail.Unlock()
	return mock.UpdateVerifiedEmailFunc(ctx, id)
}

// UpdateVerifiedEmailCalls gets all the calls that were made to UpdateVerifiedEmail.
// Check the length with:
//     len(mockedRepository.UpdateVerifiedEmailCalls())
func (mock *RepositoryMock) UpdateVerifiedEmailCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockUpdateVerifiedEmail.RLock()
	calls = mock.calls.UpdateVerifiedEmail
	mock.lockUpdateVerifiedEmail.RUnlock()
	return calls
}
