package main

import "fmt"

type Person struct {
	ID   int
	Name string
}

type PersonDB interface {
	Save(person *Person) error
}

type Database struct {
	// Database connection and operations
}

func (db *Database) Save(person *Person) error {
	// Implementation for saving a person to the database
	fmt.Printf("Saving person: %+v\n", person)

	return nil
}

type PersonService interface {
	CreatePerson(person *Person) error
}

type PersonServiceImpl struct {
	db PersonDB
}

func (s *PersonServiceImpl) CreatePerson(person *Person) error {
	// Business logic for creating a person

	err := s.db.Save(person)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	db := &Database{}
	service := &PersonServiceImpl{db: db}

	person := &Person{
		ID:   101,
		Name: "John D",
	}

	err := service.CreatePerson(person)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
