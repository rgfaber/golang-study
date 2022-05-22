package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SingletonDb(t *testing.T) {
	// Given
	// When
	db := SingletonDB()
	// Then
	assert.NotNil(t, db)

}
