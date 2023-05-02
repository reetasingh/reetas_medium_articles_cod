package main

import "fmt"

type Person struct {
	ID   int
	Name string
}

type PersonDB struct {
	// Database connection and operations
}

func (db *PersonDB) Save(person *Person) error {
	// Implementation for saving a person to the database
	fmt.Printf("Saving person: %+v\n", person)

	return nil
}

type PersonService struct {
	db *PersonDB
}

func (s *PersonService) CreatePerson(person *Person) error {
	// Business logic for creating a person

	err := s.db.Save(person)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	db := &PersonDB{}
	service := &PersonService{db: db}

	person := &Person{
		ID:   101,
		Name: "John B",
	}

	err := service.CreatePerson(person)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
