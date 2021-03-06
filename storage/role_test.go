package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllRole(t *testing.T) {
	dbmap := InitTestDB(true)

	roles, err := GetAllRole(dbmap)
	checkTestErr(err)

	assert.Equal(t, 2, len(roles))
}
