package entity

import (
	"gorm.io/gorm"
)

// TbUser defines the name of the table for the User entity in the PostgreSQL database.
const TbUser = "Tb_User"

// User represents a user in the system.
type User struct {
	gorm.Model
	Email    string `gorm:"size:100;uniqueIndex"`
	Password string `gorm:"size:100"`
}

// TableName returns the name of the table for the User model.
func (User) TableName() string {
	return TbUser
}

// ToJSONClaims formats a user entity to string mapper.
func (u *User) ToJSONClaims() map[string]interface{} {
	return map[string]interface{}{
		"ID":    u.ID,
		"Email": u.Email,
	}
}
