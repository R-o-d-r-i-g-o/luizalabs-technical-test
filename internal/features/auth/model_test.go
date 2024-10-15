package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUserEntity(t *testing.T) {
	payload := &PostRegisterPayload{
		Email:    "test@example.com",
		Password: "securepassword",
	}

	userEntity := payload.ToUserEntity()

	assert.Equal(t, payload.Email, userEntity.Email, "Expected email to match")
	assert.Equal(t, payload.Password, userEntity.Password, "Expected password to match")
}

// TestToPostLoginPayloadToInput tests the ToPostLoginPayloadToInput method of PostLoginPayload.
func TestToPostLoginPayloadToInput(t *testing.T) {
	payload := &PostLoginPayload{
		Email:    "test@example.com",
		Password: "securepassword",
	}

	loginInput := payload.ToPostLoginPayloadToInput()

	assert.Equal(t, payload.Email, loginInput.Email, "Expected email to match")
	assert.Equal(t, payload.Password, loginInput.Password, "Expected password to match")
}
