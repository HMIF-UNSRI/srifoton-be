// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	"github.com/HMIF-UNSRI/srifoton-be/internal/repository/user"
	"github.com/google/uuid"
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
// 			FindMemberByIDFunc: func(ctx context.Context, id string) (memberDomain.Member, error) {
// 				panic("mock out the FindMemberByID method")
// 			},
// 			FindTeamByIDFunc: func(ctx context.Context, id string) (teamDomain.Team, error) {
// 				panic("mock out the FindTeamByID method")
// 			},
// 			FindUserByNimFunc: func(ctx context.Context, nim string) (userDomain.User, error) {
// 				panic("mock out the FindUserByNim method")
// 			},
// 			InsertFileFunc: func(ctx context.Context, filename string) (string, error) {
// 				panic("mock out the InsertFile method")
// 			},
// 			InsertMemberFunc: func(ctx context.Context, member memberDomain.Member) (uuid.NullUUID, error) {
// 				panic("mock out the InsertMember method")
// 			},
// 			InsertTeamFunc: func(ctx context.Context, team teamDomain.Team) (string, error) {
// 				panic("mock out the InsertTeam method")
// 			},
// 			InsertUserFunc: func(ctx context.Context, user userDomain.User) (string, error) {
// 				panic("mock out the InsertUser method")
// 			},
// 			UpdatePasswordFunc: func(ctx context.Context, id string, password string) (string, error) {
// 				panic("mock out the UpdatePassword method")
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

	// FindMemberByIDFunc mocks the FindMemberByID method.
	FindMemberByIDFunc func(ctx context.Context, id string) (memberDomain.Member, error)

	// FindTeamByIDFunc mocks the FindTeamByID method.
	FindTeamByIDFunc func(ctx context.Context, id string) (teamDomain.Team, error)

	// FindUserByNimFunc mocks the FindUserByNim method.
	FindUserByNimFunc func(ctx context.Context, nim string) (userDomain.User, error)

	// InsertFileFunc mocks the InsertFile method.
	InsertFileFunc func(ctx context.Context, filename string) (string, error)

	// InsertMemberFunc mocks the InsertMember method.
	InsertMemberFunc func(ctx context.Context, member memberDomain.Member) (uuid.NullUUID, error)

	// InsertTeamFunc mocks the InsertTeam method.
	InsertTeamFunc func(ctx context.Context, team teamDomain.Team) (string, error)

	// InsertUserFunc mocks the InsertUser method.
	InsertUserFunc func(ctx context.Context, user userDomain.User) (string, error)

	// UpdatePasswordFunc mocks the UpdatePassword method.
	UpdatePasswordFunc func(ctx context.Context, id string, password string) (string, error)

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
		// FindMemberByID holds details about calls to the FindMemberByID method.
		FindMemberByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// FindTeamByID holds details about calls to the FindTeamByID method.
		FindTeamByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// FindUserByNim holds details about calls to the FindUserByNim method.
		FindUserByNim []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Nim is the nim argument value.
			Nim string
		}
		// InsertFile holds details about calls to the InsertFile method.
		InsertFile []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filename is the filename argument value.
			Filename string
		}
		// InsertMember holds details about calls to the InsertMember method.
		InsertMember []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Member is the member argument value.
			Member memberDomain.Member
		}
		// InsertTeam holds details about calls to the InsertTeam method.
		InsertTeam []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Team is the team argument value.
			Team teamDomain.Team
		}
		// InsertUser holds details about calls to the InsertUser method.
		InsertUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// User is the user argument value.
			User userDomain.User
		}
		// UpdatePassword holds details about calls to the UpdatePassword method.
		UpdatePassword []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
			// Password is the password argument value.
			Password string
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
	lockFindMemberByID      sync.RWMutex
	lockFindTeamByID        sync.RWMutex
	lockFindUserByNim       sync.RWMutex
	lockInsertFile          sync.RWMutex
	lockInsertMember        sync.RWMutex
	lockInsertTeam          sync.RWMutex
	lockInsertUser          sync.RWMutex
	lockUpdatePassword      sync.RWMutex
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

// FindMemberByID calls FindMemberByIDFunc.
func (mock *RepositoryMock) FindMemberByID(ctx context.Context, id string) (memberDomain.Member, error) {
	if mock.FindMemberByIDFunc == nil {
		panic("RepositoryMock.FindMemberByIDFunc: method is nil but Repository.FindMemberByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFindMemberByID.Lock()
	mock.calls.FindMemberByID = append(mock.calls.FindMemberByID, callInfo)
	mock.lockFindMemberByID.Unlock()
	return mock.FindMemberByIDFunc(ctx, id)
}

// FindMemberByIDCalls gets all the calls that were made to FindMemberByID.
// Check the length with:
//     len(mockedRepository.FindMemberByIDCalls())
func (mock *RepositoryMock) FindMemberByIDCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockFindMemberByID.RLock()
	calls = mock.calls.FindMemberByID
	mock.lockFindMemberByID.RUnlock()
	return calls
}

// FindTeamByID calls FindTeamByIDFunc.
func (mock *RepositoryMock) FindTeamByID(ctx context.Context, id string) (teamDomain.Team, error) {
	if mock.FindTeamByIDFunc == nil {
		panic("RepositoryMock.FindTeamByIDFunc: method is nil but Repository.FindTeamByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFindTeamByID.Lock()
	mock.calls.FindTeamByID = append(mock.calls.FindTeamByID, callInfo)
	mock.lockFindTeamByID.Unlock()
	return mock.FindTeamByIDFunc(ctx, id)
}

// FindTeamByIDCalls gets all the calls that were made to FindTeamByID.
// Check the length with:
//     len(mockedRepository.FindTeamByIDCalls())
func (mock *RepositoryMock) FindTeamByIDCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockFindTeamByID.RLock()
	calls = mock.calls.FindTeamByID
	mock.lockFindTeamByID.RUnlock()
	return calls
}

// FindUserByNim calls FindUserByNimFunc.
func (mock *RepositoryMock) FindUserByNim(ctx context.Context, nim string) (userDomain.User, error) {
	if mock.FindUserByNimFunc == nil {
		panic("RepositoryMock.FindUserByNimFunc: method is nil but Repository.FindUserByNim was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Nim string
	}{
		Ctx: ctx,
		Nim: nim,
	}
	mock.lockFindUserByNim.Lock()
	mock.calls.FindUserByNim = append(mock.calls.FindUserByNim, callInfo)
	mock.lockFindUserByNim.Unlock()
	return mock.FindUserByNimFunc(ctx, nim)
}

// FindUserByNimCalls gets all the calls that were made to FindUserByNim.
// Check the length with:
//     len(mockedRepository.FindUserByNimCalls())
func (mock *RepositoryMock) FindUserByNimCalls() []struct {
	Ctx context.Context
	Nim string
} {
	var calls []struct {
		Ctx context.Context
		Nim string
	}
	mock.lockFindUserByNim.RLock()
	calls = mock.calls.FindUserByNim
	mock.lockFindUserByNim.RUnlock()
	return calls
}

// InsertFile calls InsertFileFunc.
func (mock *RepositoryMock) InsertFile(ctx context.Context, filename string) (string, error) {
	if mock.InsertFileFunc == nil {
		panic("RepositoryMock.InsertFileFunc: method is nil but Repository.InsertFile was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Filename string
	}{
		Ctx:      ctx,
		Filename: filename,
	}
	mock.lockInsertFile.Lock()
	mock.calls.InsertFile = append(mock.calls.InsertFile, callInfo)
	mock.lockInsertFile.Unlock()
	return mock.InsertFileFunc(ctx, filename)
}

// InsertFileCalls gets all the calls that were made to InsertFile.
// Check the length with:
//     len(mockedRepository.InsertFileCalls())
func (mock *RepositoryMock) InsertFileCalls() []struct {
	Ctx      context.Context
	Filename string
} {
	var calls []struct {
		Ctx      context.Context
		Filename string
	}
	mock.lockInsertFile.RLock()
	calls = mock.calls.InsertFile
	mock.lockInsertFile.RUnlock()
	return calls
}

// InsertMember calls InsertMemberFunc.
func (mock *RepositoryMock) InsertMember(ctx context.Context, member memberDomain.Member) (uuid.NullUUID, error) {
	if mock.InsertMemberFunc == nil {
		panic("RepositoryMock.InsertMemberFunc: method is nil but Repository.InsertMember was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Member memberDomain.Member
	}{
		Ctx:    ctx,
		Member: member,
	}
	mock.lockInsertMember.Lock()
	mock.calls.InsertMember = append(mock.calls.InsertMember, callInfo)
	mock.lockInsertMember.Unlock()
	return mock.InsertMemberFunc(ctx, member)
}

// InsertMemberCalls gets all the calls that were made to InsertMember.
// Check the length with:
//     len(mockedRepository.InsertMemberCalls())
func (mock *RepositoryMock) InsertMemberCalls() []struct {
	Ctx    context.Context
	Member memberDomain.Member
} {
	var calls []struct {
		Ctx    context.Context
		Member memberDomain.Member
	}
	mock.lockInsertMember.RLock()
	calls = mock.calls.InsertMember
	mock.lockInsertMember.RUnlock()
	return calls
}

// InsertTeam calls InsertTeamFunc.
func (mock *RepositoryMock) InsertTeam(ctx context.Context, team teamDomain.Team) (string, error) {
	if mock.InsertTeamFunc == nil {
		panic("RepositoryMock.InsertTeamFunc: method is nil but Repository.InsertTeam was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Team teamDomain.Team
	}{
		Ctx:  ctx,
		Team: team,
	}
	mock.lockInsertTeam.Lock()
	mock.calls.InsertTeam = append(mock.calls.InsertTeam, callInfo)
	mock.lockInsertTeam.Unlock()
	return mock.InsertTeamFunc(ctx, team)
}

// InsertTeamCalls gets all the calls that were made to InsertTeam.
// Check the length with:
//     len(mockedRepository.InsertTeamCalls())
func (mock *RepositoryMock) InsertTeamCalls() []struct {
	Ctx  context.Context
	Team teamDomain.Team
} {
	var calls []struct {
		Ctx  context.Context
		Team teamDomain.Team
	}
	mock.lockInsertTeam.RLock()
	calls = mock.calls.InsertTeam
	mock.lockInsertTeam.RUnlock()
	return calls
}

// InsertUser calls InsertUserFunc.
func (mock *RepositoryMock) InsertUser(ctx context.Context, user userDomain.User) (string, error) {
	if mock.InsertUserFunc == nil {
		panic("RepositoryMock.InsertUserFunc: method is nil but Repository.InsertUser was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		User userDomain.User
	}{
		Ctx:  ctx,
		User: user,
	}
	mock.lockInsertUser.Lock()
	mock.calls.InsertUser = append(mock.calls.InsertUser, callInfo)
	mock.lockInsertUser.Unlock()
	return mock.InsertUserFunc(ctx, user)
}

// InsertUserCalls gets all the calls that were made to InsertUser.
// Check the length with:
//     len(mockedRepository.InsertUserCalls())
func (mock *RepositoryMock) InsertUserCalls() []struct {
	Ctx  context.Context
	User userDomain.User
} {
	var calls []struct {
		Ctx  context.Context
		User userDomain.User
	}
	mock.lockInsertUser.RLock()
	calls = mock.calls.InsertUser
	mock.lockInsertUser.RUnlock()
	return calls
}

// UpdatePassword calls UpdatePasswordFunc.
func (mock *RepositoryMock) UpdatePassword(ctx context.Context, id string, password string) (string, error) {
	if mock.UpdatePasswordFunc == nil {
		panic("RepositoryMock.UpdatePasswordFunc: method is nil but Repository.UpdatePassword was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ID       string
		Password string
	}{
		Ctx:      ctx,
		ID:       id,
		Password: password,
	}
	mock.lockUpdatePassword.Lock()
	mock.calls.UpdatePassword = append(mock.calls.UpdatePassword, callInfo)
	mock.lockUpdatePassword.Unlock()
	return mock.UpdatePasswordFunc(ctx, id, password)
}

// UpdatePasswordCalls gets all the calls that were made to UpdatePassword.
// Check the length with:
//     len(mockedRepository.UpdatePasswordCalls())
func (mock *RepositoryMock) UpdatePasswordCalls() []struct {
	Ctx      context.Context
	ID       string
	Password string
} {
	var calls []struct {
		Ctx      context.Context
		ID       string
		Password string
	}
	mock.lockUpdatePassword.RLock()
	calls = mock.calls.UpdatePassword
	mock.lockUpdatePassword.RUnlock()
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
