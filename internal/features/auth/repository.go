package auth

import (
	"luizalabs-technical-test/internal/pkg/entity"

	"gorm.io/gorm"
)

// RepositoryImp defines the interface for the repository layer,
// which abstracts data access operations.
type RepositoryImp interface {
	RegisterUser(user entity.User) error
	GetUser(filter GetUserFilter) (*entity.User, error)
}

// repository struct implements the repositoryImp interface,
// that interacts with external entities such as databases or external APIs.
type repository struct {
	db *gorm.DB
}

// NewRepository creates and returns a new instance of the repository.
func NewRepository(db *gorm.DB) RepositoryImp {
	return &repository{db}
}

// RegisterUser adds a new user to the database.
func (r *repository) RegisterUser(user entity.User) error {
	tx := r.db.Table(entity.TbUser).Create(&user)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

// GetUser retrieves a user by email from the database.
func (r *repository) GetUser(filter GetUserFilter) (*entity.User, error) {
	fetchedUser := new(entity.User)

	tx := r.db.Where("Email = ?", filter.Email).First(&fetchedUser)
	if err := tx.Error; err != nil {
		return nil, err
	}

	return fetchedUser, nil
}
