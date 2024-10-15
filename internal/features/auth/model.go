package auth

import "luizalabs-technical-test/internal/pkg/entity"

// PostRegisterPayload represents the payload for register a user in database.
type PostRegisterPayload struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// PostLoginPayload represents the payload for user login requests.
type PostLoginPayload struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthenticateUserInput represents the input structure in
// service layer for autentication of user login.
type AuthenticateUserInput struct {
	Email    string
	Password string
}

// AuthenticateUserResponse represents the response structure
// containing a JWT token upon successful user authentication.
type AuthenticateUserResponse struct {
	JWTToken string `json:"token"`
}

// GetUserFilter represents the filter criteria for querying users.
type GetUserFilter struct {
	Email string
}

// ToUserEntity converts the PostRegisterPayload to a User database entity.
func (p *PostRegisterPayload) ToUserEntity() entity.User {
	return entity.User{
		Email:    p.Email,
		Password: p.Password,
	}
}

// ToPostLoginPayloadToInput maps PostLoginPayload to PostLoginInput.
func (p *PostLoginPayload) ToPostLoginPayloadToInput() AuthenticateUserInput {
	return AuthenticateUserInput{
		Email:    p.Email,
		Password: p.Password,
	}
}

// ToPostLoginInputToFilter maps PostLoginInput to GetUserFilter.
func (i *AuthenticateUserInput) ToPostLoginInputToFilter() GetUserFilter {
	return GetUserFilter{
		Email: i.Email,
	}
}
