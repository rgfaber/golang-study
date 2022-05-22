package main

import (
	"fmt"
	"layered-architecture/business"
)

func main() {
	ps := business.NewPersonService()
	ls := ps.GetAllPersons()
	for _, p := range ls {
		fmt.Println(p)

	}
}
