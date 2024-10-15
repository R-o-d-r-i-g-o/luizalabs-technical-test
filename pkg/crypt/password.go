package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher defines the interface for password hashing operations.
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

// passwordHasher is an implementation of the PasswordHasher interface using bcrypt.
type passwordHasher struct{}

// NewPasswordHasher creates a new instance of BcryptPasswordHasher.
func NewPasswordHasher() PasswordHasher {
	return &passwordHasher{}
}

// HashPassword hashes a password using bcrypt.
func (b *passwordHasher) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// CheckPasswordHash verifies if the provided password matches the hash.
func (b *passwordHasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
