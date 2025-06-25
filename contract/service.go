package contract

import (
	"context"
	"library_api/dto"
)

type Service struct {
	Auth AuthService
	Book BookService
}

type AuthService interface {
	Register(ctx context.Context, payload *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, payload *dto.LoginRequest) (*dto.LoginResponse, error)
}

type BookService interface {
	GetAllBooks() (*dto.BookListResponse, error)
	GetBookByID(id uint64) (*dto.BookDetailResponse, error)
	InsertBook(payload *dto.BookRequest, pdfFileBytes []byte) (*dto.BookDetailResponse, error)
}
