package repository

import (
	"clean-architecture/infra/database"
	mocks_repo "clean-architecture/mocks/infra/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewRepository(t *testing.T) {
	// Given
	db := mocks_repo.NewIDb(t)
	// When
	repo := NewRepo(db)
	// Then
	assert.NotNil(t, repo)
}

func Test_Create_Db_Directly_From_Repo(t *testing.T) {
	// Given
	// When
	db := database.SingletonDB()
	// Then
	assert.NotNil(t, db)
}
