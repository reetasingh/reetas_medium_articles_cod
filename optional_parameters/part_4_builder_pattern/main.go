package main

import "fmt"

type Person struct {
	name  string
	phone string
	email string
	// optional fields
	twitterHandle string
}

func NewPerson(name string, phone string, email string) Person {
	p := Person{}
	p.name = name
	p.phone = phone
	p.email = email
	return p
}

func (p Person) WithTwitterHandle(twitterHandle string) Person {
	p.twitterHandle = twitterHandle
	return p
}

func main() {
	person_1 := NewPerson("rohit", "1739203", "rohit@gmail.com").WithTwitterHandle("ro1")
	fmt.Println(person_1)
	person_2 := NewPerson("jeremy", "48848", "jeremy@gmail.com")
	fmt.Println(person_2)
	person_3 := NewPerson("martin", "28483", "martin@outlook.com")
	fmt.Println(person_3)
}
