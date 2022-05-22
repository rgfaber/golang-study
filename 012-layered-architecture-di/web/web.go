package web

import "layered-architecture/business"

type IPersonController interface {
}

type PersonController struct {
	Service business.IPersonService
}

func NewPersonController(service business.IPersonService) IPersonController {
	return &PersonController{
		Service: service,
	}
}
