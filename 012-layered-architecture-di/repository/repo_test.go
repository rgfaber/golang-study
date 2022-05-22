package repository

import (
	"github.com/stretchr/testify/assert"
	"layered-architecture/database"
	mocks "layered-architecture/mocks/database"
	"testing"
)

func Test_NewRepository(t *testing.T) {
	// Given
	db := mocks.NewIDb(t)
	// When
	repo := NewRepo(db)
	// Then
	assert.NotNil(t, repo)
}

func Test_Create_Db_From_Repo(t *testing.T) {
	// Given
	// When
	db := database.SingletonDB()
	// Then
	assert.NotNil(t, db)
}
