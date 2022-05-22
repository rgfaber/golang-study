package repository

type IDb interface {
	Create(id string, data interface{}) (interface{}, error)
	Update(id string, data interface{}) (interface{}, error)
	Delete(id string) (interface{}, error)
	Select(id string) (interface{}, error)
	SelectStar() map[string]interface{}
}
