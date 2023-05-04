package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTestDB struct {
	mock.Mock
}

func (m *mockTestDB) Save(person *Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func TestPersonServiceImpl_CreatePerson_NoError(t *testing.T) {
	testDb := new(mockTestDB)
	p := &Person{
		ID:   1,
		Name: "John D",
	}
	testDb.On("Save", p).Return(nil)
	personServiceImpl := new(PersonServiceImpl)
	personServiceImpl.db = testDb
	err := personServiceImpl.CreatePerson(p)
	assert.NoError(t, err)
}

func TestPersonServiceImpl_CreatePerson_Error(t *testing.T) {
	testDb := new(mockTestDB)
	p := &Person{
		ID:   1,
		Name: "John D",
	}
	testDb.On("Save", p).Return(fmt.Errorf("save operation failed"))
	personServiceImpl := new(PersonServiceImpl)
	personServiceImpl.db = testDb
	err := personServiceImpl.CreatePerson(p)
	assert.Error(t, err)
}
