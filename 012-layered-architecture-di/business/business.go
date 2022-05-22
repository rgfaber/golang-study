package business

import (
	"layered-architecture/model"
	"layered-architecture/repository"
)

type PersonService struct {
	Repo repository.IRepo
}

func (p *PersonService) GetAllPersons() []*model.Person {
	result := make([]*model.Person, 1)
	res := p.Repo.GetAll()
	for _, p := range res {
		pers := p.(*model.Person)
		result = append(result, pers)
	}
	return result
}

func NewPersonService(repo repository.IRepo) IPersonService {
	return &PersonService{
		Repo: repo,
	}
}
