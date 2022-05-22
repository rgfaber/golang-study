package domain

import (
	"clean-architecture/infra/database"
	"clean-architecture/infra/repository"
	mocks_domain "clean-architecture/mocks/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewPersonService(t *testing.T) {
	// Given
	repo := mocks_domain.NewIRepo(t)
	// When
	srv := NewPersonService(repo)
	// Then
	assert.NotNil(t, srv)
}

func Test_Inject_Concrete_Repo_To_Business(t *testing.T) {
	// Given
	db := database.SingletonDB()
	repo := repository.NewRepo(db)
	// When
	srv := NewPersonService(repo)
	// Then
	assert.NotNil(t, srv)
}

func Test_Create_DB_Directly_From_Business(t *testing.T) {
	// Given
	// When
	db := database.SingletonDB()
	// Then
	assert.NotNil(t, db)
}
