package service

import (
	"fmt"
	"library_api/contract"
	"library_api/dto"
	"library_api/entity"
	"library_api/pkg/helpers"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gen2brain/go-fitz"
)

type BookService struct {
	bookRepository contract.BookRepository
}

func implBookService(repo *contract.Repository) contract.BookService {
	return &BookService{
		bookRepository: repo.Book,
	}
}

func (s *BookService) GetAllBooks() (*dto.BookListResponse, error) {
	books, err := s.bookRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}

	data := []dto.BookListData{}
	for _, book := range books {
		data = append(data, dto.BookListData{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Genre:  book.Genre,
		})
	}

	responses := &dto.BookListResponse{
		StatusCode: http.StatusOK,
		Message:    "List Buku berhasil diambil",
		Data:       data,
	}
	return responses, nil
}

func (s *BookService) GetBookByID(id uint64) (*dto.BookDetailResponse, error) {
	book, err := s.bookRepository.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.BookDetailResponse{
		StatusCode: http.StatusOK,
		Message:    "Buku berhasil diambil",
		Data: dto.BookData{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Year:        book.Year,
			Genre:       book.Genre,
			Description: book.Description,
			Publisher:   book.Publisher,
			ISBN:        book.ISBN,
			Language:    book.Language,
			Pages:       book.Pages,
			Thumbnail:   book.Thumbnail,
			FilePath:    book.FilePath,
		},
	}, nil
}

func (s *BookService) InsertBook(payload *dto.BookRequest, pdfFileBytes []byte) (*dto.BookDetailResponse, error) {
	validPayload := helpers.ValidateStruct(payload)
	if validPayload != nil {
		return nil, validPayload
	}

	baseURL := os.Getenv("BASE_URL")
	rand.Seed(time.Now().UnixNano())
	randomNum := strconv.FormatInt(rand.Int63(), 10)
	pdfFileName := randomNum + ".pdf"
	thumbFileName := randomNum + "_thumb.jpg"

	pdfDir := "assets/pdf"
	thumbnailDir := "assets/thumbnails"
	// 1. Save PDF file
	if err := os.MkdirAll(pdfDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create pdf directory: %w", err)
	}
	pdfPath := filepath.Join(pdfDir, pdfFileName)
	if err := os.WriteFile(pdfPath, pdfFileBytes, 0644); err != nil {
		return nil, fmt.Errorf("failed to save pdf: %w", err)
	}

	// 2. Generate thumbnail from PDF cover (first page)
	if err := os.MkdirAll(thumbnailDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create thumbnail directory: %w", err)
	}
	thumbnailPath := filepath.Join(thumbnailDir, thumbFileName+".jpg")

	doc, err := fitz.New(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open pdf for thumbnail: %w", err)
	}
	defer doc.Close()

	img, err := doc.Image(0)
	if err != nil {
		return nil, fmt.Errorf("failed to render pdf cover: %w", err)
	}

	// Optionally resize thumbnail
	thumbImg := imaging.Resize(img, 300, 0, imaging.Lanczos)
	if err := imaging.Save(thumbImg, thumbnailPath); err != nil {
		return nil, fmt.Errorf("failed to save thumbnail: %w", err)
	}

	pdfURL := fmt.Sprintf("%s/assets/pdf/%s", baseURL, pdfFileName)
	thumbURL := fmt.Sprintf("%s/assets/thumbnails/%s", baseURL, thumbFileName)

	book := &entity.Book{
		Title:       payload.Title,
		Author:      payload.Author,
		Year:        payload.Year,
		Genre:       payload.Genre,
		Description: payload.Description,
		Publisher:   payload.Publisher,
		ISBN:        payload.ISBN,
		Language:    payload.Language,
		Pages:       payload.Pages,
		Thumbnail:   thumbURL,
		FilePath:    pdfURL,
	}
	err = s.bookRepository.InsertBook(book)
	if err != nil {
		return nil, err
	}

	// 4. Prepare response
	response := &dto.BookDetailResponse{
		StatusCode: http.StatusCreated,
		Message:    "Buku berhasil diupload",
		Data: dto.BookData{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Year:        book.Year,
			Genre:       book.Genre,
			Description: book.Description,
			Publisher:   book.Publisher,
			ISBN:        book.ISBN,
			Language:    book.Language,
			Pages:       book.Pages,
			Thumbnail:   book.Thumbnail,
			FilePath:    book.FilePath,
		},
	}
	return response, nil
}
