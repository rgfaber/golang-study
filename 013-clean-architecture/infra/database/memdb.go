package database

import (
	"clean-architecture/infra/repository"
	"fmt"
	"sync"
)

type MemDB struct {
	Data map[string]interface{}
}

var sdb *MemDB
var mutex = &sync.Mutex{}

func SingletonDB() repository.IDb {
	if sdb == nil {
		mutex.Lock()
		defer mutex.Unlock()
		sdb = &MemDB{
			Data: make(map[string]interface{}, 1),
		}
	}
	return sdb
}

func (db *MemDB) Create(id string, data interface{}) (interface{}, error) {
	_, found := db.Data[id]
	if found {
		err := fmt.Errorf("element with id [%+v] already exists", id)
		return nil, err
	}
	db.Data[id] = data
	return db.Data[id], nil
}

func (db *MemDB) Update(id string, data interface{}) (interface{}, error) {
	_, found := db.Data[id]
	if !found {
		err := fmt.Errorf("id [%+v] not found", id)
		return nil, err
	}
	db.Data[id] = data
	return db.Data[id], nil
}

func (db *MemDB) Delete(id string) (interface{}, error) {
	it, found := db.Data[id]
	if !found {
		err := fmt.Errorf("item [%+v] was not found", id)
		return nil, err
	}
	delete(db.Data, id)
	return it, nil
}

func (db *MemDB) Select(id string) (interface{}, error) {
	it, found := db.Data[id]
	if !found {
		err := fmt.Errorf("item [%+v] not found", id)
		return nil, err
	}
	return it, nil
}

func (db *MemDB) SelectStar() map[string]interface{} {
	return db.Data
}
