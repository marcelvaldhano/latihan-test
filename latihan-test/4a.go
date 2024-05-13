package main

import (
	"strings"
)

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{name: name}
}

func (p *Person) GetFullName() string {
	return strings.Title(p.name)
}

type FullNamePerson struct {
	*Person
}

func NewFullNamePerson(name string) *FullNamePerson {
	return &FullNamePerson{Person: NewPerson(name)}
}

func (p *FullNamePerson) GetFullName() string {
	return strings.Title(p.name)
}
