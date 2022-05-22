package web

import (
	"clean-architecture/domain"
	"clean-architecture/infra/database"
	"clean-architecture/infra/repository"
	mocks_domain "clean-architecture/mocks/domain"
	mocks_repo "clean-architecture/mocks/infra/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewPersonController(t *testing.T) {
	// Given
	db := database.SingletonDB()
	repo := repository.NewRepo(db)
	svc := domain.NewPersonService(repo)
	// When
	pc := NewPersonController(svc)
	// Then
	assert.NotNil(t, pc)
}

func Test_Create_Service_From_Web(t *testing.T) {
	// Given
	repo := mocks_domain.NewIRepo(t)
	// When
	svc := domain.NewPersonService(repo)
	// Then
	assert.NotNil(t, svc)
}

func Test_Create_Repo_From_Web(t *testing.T) {
	// Given
	db := mocks_repo.NewIDb(t)
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
