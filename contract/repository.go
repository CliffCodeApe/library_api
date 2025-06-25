package contract

import (
	"context"
	"library_api/entity"
)

type Repository struct {
	User UserRepository
	Book BookRepository
}
type UserRepository interface {
	InsertUser(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByID(id uint64) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
	IsEmailExists(email string) (bool, error)
}

type BookRepository interface {
	GetAllBooks() ([]entity.Book, error)
	GetBookByID(id uint64) (*entity.Book, error)
	InsertBook(book *entity.Book) error
}
