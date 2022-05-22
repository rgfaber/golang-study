package domain

type IRepo interface {
	Save(id string, data interface{}) interface{}
	Remove(id string) interface{}
	Get(id string) interface{}
	GetAll() map[string]interface{}
}
