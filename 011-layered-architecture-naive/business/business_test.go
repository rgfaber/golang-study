package business

import (
	"github.com/stretchr/testify/assert"
	"layered-architecture/database"
	"layered-architecture/repository"
	"testing"
)

func Test_NewPersonService(t *testing.T) {
	// Given
	// When
	srv := NewPersonService()
	// Then
	assert.NotNil(t, srv)
}

func Test_Create_Repo_From_Business(t *testing.T) {
	repo := repository.NewRepo()
	assert.NotNil(t, repo)
}

func Test_Create_DB_From_Business(t *testing.T) {
	db := database.SingletonDB()
	assert.NotNil(t, db)
}
