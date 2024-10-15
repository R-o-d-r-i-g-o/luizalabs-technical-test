package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestTableName(t *testing.T) {
	var (
		user      User
		tableName = user.TableName()
	)

	assert.Equal(t, TbUser, tableName)
}

func TestToJSONClaims(t *testing.T) {
	user := &User{
		Model:    gorm.Model{ID: 1},
		Email:    "test@example.com",
		Password: "password123",
	}
	expectedClaims := map[string]interface{}{
		"ID":    user.ID,
		"Email": user.Email,
	}

	claims := user.ToJSONClaims()
	assert.Equal(t, expectedClaims, claims)
}
