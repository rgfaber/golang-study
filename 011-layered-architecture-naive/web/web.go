package web

import "layered-architecture/business"

type IPersonController interface {
}

type PersonController struct {
	Service business.IPersonService
}

func NewPersonController() IPersonController {
	return &PersonController{
		Service: business.NewPersonService(),
	}
}
