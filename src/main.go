package main

import (
	"fmt"
)

func main() {
	gohan := &Saiyan{
		Name: "Gohan",
		Power: 1000,
		Father: &Saiyan {
			Name: "Goku",
			Power: 9002,
			Father: nil,
		},
	}

	gohan.Father.Super()

	fmt.Println(gohan.Father.Power)
}

func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name: name,
		Power: power,
	}
}

type Saiyan struct {
	Name string
	Power int
	Father *Saiyan
}

func (s *Saiyan) Super() {
	s.Power += 10000
}
