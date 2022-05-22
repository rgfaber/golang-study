package business

import (
	"github.com/stretchr/testify/assert"
	"layered-architecture/database"
	mocks_db "layered-architecture/mocks/database"
	mocks_repo "layered-architecture/mocks/repository"
	"layered-architecture/repository"
	"testing"
)

func Test_NewPersonService(t *testing.T) {
	// Given
	repo := mocks_repo.NewIRepo(t)
	// When
	srv := NewPersonService(repo)
	// Then
	assert.NotNil(t, srv)
}

func Test_Create_Repo_From_Business(t *testing.T) {
	// Given
	db := mocks_db.NewIDb(t)
	// When
	repo := repository.NewRepo(db)
	// Then
	assert.NotNil(t, repo)
}

func Test_Create_DB_From_Business(t *testing.T) {
	// Given
	// When
	db := database.SingletonDB()
	// Then
	assert.NotNil(t, db)
}
