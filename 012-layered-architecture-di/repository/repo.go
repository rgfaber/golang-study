package repository

import (
	"layered-architecture/database"
	"sync"
)

var mutex = &sync.Mutex{}

type Repo struct {
	Db database.IDb
}

func (r *Repo) GetAll() map[string]interface{} {
	return r.Db.SelectStar()
}

func NewRepo(db database.IDb) IRepo {
	return &Repo{
		Db: db,
	}
}

func (r *Repo) Save(id string, data interface{}) interface{} {
	it, err := r.Db.Update(id, data)
	if err != nil {
		it, _ = r.Db.Create(id, data)
	}
	return it
}

func (r *Repo) Remove(id string) interface{} {
	it, err := r.Db.Delete(id)
	if err != nil {
		return nil
	}
	return it
}

func (r *Repo) Get(id string) interface{} {
	it, err := r.Db.Select(id)
	if err != nil {
		return nil
	}
	return it
}
