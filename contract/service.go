package contract

import (
	"context"
	"library_api/dto"
)

type Service struct {
	Auth    AuthService
	Book    BookService
	Lending LendingService
}

type AuthService interface {
	Register(ctx context.Context, payload *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, payload *dto.LoginRequest) (*dto.LoginResponse, error)
}

type BookService interface {
	GetAllBooks() (*dto.BookListResponse, error)
	GetBookByID(id uint64) (*dto.BookDetailResponse, error)
	InsertBook(payload *dto.BookRequest, pdfFileBytes []byte) (*dto.BookDetailResponse, error)
	GetBooksByGenre(genre string) (*dto.BookListResponse, error)
	SearchBooks(keyword string) (*dto.BookListResponse, error)
}

type LendingService interface {
	GetAllLendings() (*dto.LendingResponse, error)
	MakeLending(userID uint64, payload *dto.LendingRequest) (*dto.MakeLendingResponse, error)
	ReturnBook(lendingID, userID uint64, fileBytes []byte) (*dto.MakeLendingResponse, error)
	GetLendingsByStatus(status string) (*dto.LendingResponse, error)
	SearchLendings(keyword string) (*dto.LendingResponse, error)
}
