package web

import (
	"github.com/stretchr/testify/assert"
	"layered-architecture/business"
	"layered-architecture/database"
	"layered-architecture/repository"
	"testing"
)

func Test_NewPersonController(t *testing.T) {
	// Given
	// When
	pc := NewPersonController()
	// Then
	assert.NotNil(t, pc)
}

func Test_Create_Service_From_Web(t *testing.T) {
	// Given
	// When
	svc := business.NewPersonService()
	// Then
	assert.NotNil(t, svc)
}

func Test_Create_Repo_From_Web(t *testing.T) {
	// Given
	// When
	repo := repository.NewRepo()
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
