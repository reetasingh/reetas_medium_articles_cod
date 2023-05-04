package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockTestDB struct {
}

func (m *mockTestDB) Save(person *Person) error {
	// just a mock, don't do anything
	return nil
}

func (m *mockTestDB) Save(person *Person) error {
	// just a mock, don't do anything
	return fmt.Errorf("error in save")
}

func TestPersonServiceImpl_CreatePerson(t *testing.T) {
	p := Person{
		ID:  1,
		err: nil,
	}
	personServiceImpl := new(PersonServiceImpl)
	personServiceImpl.db = mockTestDB{}
	err := personServiceImpl.CreatePerson(p)
	assert.NoError(t, err)
}
