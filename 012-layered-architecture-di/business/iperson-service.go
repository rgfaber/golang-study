package business

import "layered-architecture/model"

type IPersonService interface {
	GetAllPersons() []*model.Person
}
