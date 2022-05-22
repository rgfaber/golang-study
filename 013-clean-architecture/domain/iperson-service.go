package domain

import "clean-architecture/model"

type IPersonService interface {
	GetAllPersons() []*model.Person
}
