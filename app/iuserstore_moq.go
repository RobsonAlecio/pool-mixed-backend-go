// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package app

import (
	"sync"
)

var (
	lockIUserStoreMockFindOne sync.RWMutex
	lockIUserStoreMockSave    sync.RWMutex
)

// IUserStoreMock is a mock implementation of IUserStore.
//
//     func TestSomethingThatUsesIUserStore(t *testing.T) {
//
//         // make and configure a mocked IUserStore
//         mockedIUserStore := &IUserStoreMock{
//             FindOneFunc: func(q *UserQuery) (*User, error) {
// 	               panic("mock out the FindOne method")
//             },
//             SaveFunc: func(record *User) (bool, error) {
// 	               panic("mock out the Save method")
//             },
//         }
//
//         // use mockedIUserStore in code that requires IUserStore
//         // and then make assertions.
//
//     }
type IUserStoreMock struct {
	// FindOneFunc mocks the FindOne method.
	FindOneFunc func(q *UserQuery) (*User, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(record *User) (bool, error)

	// calls tracks calls to the methods.
	calls struct {
		// FindOne holds details about calls to the FindOne method.
		FindOne []struct {
			// Q is the q argument value.
			Q *UserQuery
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Record is the record argument value.
			Record *User
		}
	}
}

// FindOne calls FindOneFunc.
func (mock *IUserStoreMock) FindOne(q *UserQuery) (*User, error) {
	if mock.FindOneFunc == nil {
		panic("IUserStoreMock.FindOneFunc: method is nil but IUserStore.FindOne was just called")
	}
	callInfo := struct {
		Q *UserQuery
	}{
		Q: q,
	}
	lockIUserStoreMockFindOne.Lock()
	mock.calls.FindOne = append(mock.calls.FindOne, callInfo)
	lockIUserStoreMockFindOne.Unlock()
	return mock.FindOneFunc(q)
}

// FindOneCalls gets all the calls that were made to FindOne.
// Check the length with:
//     len(mockedIUserStore.FindOneCalls())
func (mock *IUserStoreMock) FindOneCalls() []struct {
	Q *UserQuery
} {
	var calls []struct {
		Q *UserQuery
	}
	lockIUserStoreMockFindOne.RLock()
	calls = mock.calls.FindOne
	lockIUserStoreMockFindOne.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *IUserStoreMock) Save(record *User) (bool, error) {
	if mock.SaveFunc == nil {
		panic("IUserStoreMock.SaveFunc: method is nil but IUserStore.Save was just called")
	}
	callInfo := struct {
		Record *User
	}{
		Record: record,
	}
	lockIUserStoreMockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	lockIUserStoreMockSave.Unlock()
	return mock.SaveFunc(record)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//     len(mockedIUserStore.SaveCalls())
func (mock *IUserStoreMock) SaveCalls() []struct {
	Record *User
} {
	var calls []struct {
		Record *User
	}
	lockIUserStoreMockSave.RLock()
	calls = mock.calls.Save
	lockIUserStoreMockSave.RUnlock()
	return calls
}