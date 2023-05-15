package domain

import "github.com/google/uuid"

type Excursion struct {
	Id          uuid.UUID `bun:",pk"`
	Name        string
	Description string
	Price       int
}

func NewExcursion(name string, description string, price int) *Excursion {
	return &Excursion{
		Id:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
	}
}
