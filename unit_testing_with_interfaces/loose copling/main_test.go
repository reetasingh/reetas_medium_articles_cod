package main

import "testing"

type mockTestDB struct {

}

func (m * mockTestDB)Save(person *Person) error {
	// just a mock, don't do enything
	return nil
}

func TestPersonServiceImpl_CreatePerson(t *testing.T) {
	p := Person{
		ID: 1,
		err: nil
	}
	personServiceImpl := new(PersonServiceImpl)
	personServiceImpl.db := new(mockTestDB)
	err := personServiceImpl.CreatePerson(p)
	assert.NoError(err)
}
