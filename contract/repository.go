package contract

import (
	"context"
	"library_api/entity"
)

type Repository struct {
	User    UserRepository
	Book    BookRepository
	Lending LendingRepository
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
	ChangeStock(id uint64, delta int) error
	GetBooksByGenre(genre string) ([]entity.Book, error)
	SearchBooks(keyword string) ([]entity.Book, error)
	GetBookByURL(url string) (*entity.Book, error)
}

type LendingRepository interface {
	GetAllLendings() ([]entity.Lending, error)
	GetLendingByID(id uint64) (*entity.Lending, error)
	MakeLending(lending *entity.Lending) error
	ChangeLendingStatus(id uint64, lending *entity.Lending) error
	GetLendingsByStatus(status string) ([]entity.Lending, error)
	SearchLendings(keyword string) ([]entity.Lending, error)
}
