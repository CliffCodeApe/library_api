package repository

import (
	"library_api/contract"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *contract.Repository {
	return &contract.Repository{
		User: implUserRepository(db),
		// Code here
		// Example:
		// Example: implExampleRepository(db),
	}
}
