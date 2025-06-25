package repository

import (
	"context"
	"library_api/contract"
	"library_api/entity"

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

func (u *userRepo) InsertUser(ctx context.Context, user *entity.User) error {
	err := u.db.WithContext(ctx).Create(&user).Error
	return err
}
func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *userRepo) GetUserByID(id uint64) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *userRepo) UpdateUser(ctx context.Context, user *entity.User) error {
	err := u.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *userRepo) DeleteUser(ctx context.Context, id int) error {
	err := u.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *userRepo) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	err := u.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) IsEmailExists(email string) (bool, error) {
	var exists bool
	err := r.db.Raw("SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists).Error
	return exists, err
}
