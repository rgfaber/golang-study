package web

import (
	"github.com/stretchr/testify/assert"
	"layered-architecture/business"
	"layered-architecture/database"
	mocks_business "layered-architecture/mocks/business"
	mocks_db "layered-architecture/mocks/database"
	mocks_repo "layered-architecture/mocks/repository"
	"layered-architecture/repository"
	"testing"
)

func Test_NewPersonController(t *testing.T) {
	// Given
	svc := mocks_business.NewIPersonService(t)
	// When
	pc := NewPersonController(svc)
	// Then
	assert.NotNil(t, pc)
}

func Test_Create_Service_From_Web(t *testing.T) {
	// Given
	repo := mocks_repo.NewIRepo(t)
	// When
	svc := business.NewPersonService(repo)
	// Then
	assert.NotNil(t, svc)
}

func Test_Create_Repo_From_Web(t *testing.T) {
	// Given
	db := mocks_db.NewIDb(t)
	// When
	repo := repository.NewRepo(db)
	// Then
	assert.NotNil(t, repo)
}

func Test_Create_Db_From_Web(t *testing.T) {
	// Given
	// When
	db := database.SingletonDB()
	// Then
	assert.NotNil(t, db)
}
