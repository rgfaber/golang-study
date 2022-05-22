package domain

import "clean-architecture/model"

type PersonService struct {
	Repo IRepo
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

func NewPersonService(repo IRepo) IPersonService {
	return &PersonService{
		Repo: repo,
	}
}
