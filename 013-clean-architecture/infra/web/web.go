package web

import "clean-architecture/domain"

type PersonController struct {
	Service domain.IPersonService
}

func NewPersonController(service domain.IPersonService) domain.IPersonController {
	return &PersonController{
		Service: service,
	}
}
