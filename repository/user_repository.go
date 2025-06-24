package repository

import (
	"library_api/contract"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func implUserRepository(db *gorm.DB) contract.UserRepository {
	return &userRepo{
		db: db,
	}
}
