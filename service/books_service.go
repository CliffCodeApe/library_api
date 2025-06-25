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

func (b *BookService) GetAllBooks() (*dto.BookListResponse, error) {
	books, err := b.bookRepository.GetAllBooks()
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

func (b *BookService) GetBookByID(id uint64) (*dto.BookDetailResponse, error) {
	book, err := b.bookRepository.GetBookByID(id)
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
			Stock:       book.Stock,
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

func (b *BookService) GetBooksByGenre(genre string) (*dto.BookListResponse, error) {
	books, err := b.bookRepository.GetBooksByGenre(genre)
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

func (b *BookService) InsertBook(payload *dto.BookRequest, pdfFileBytes []byte) (*dto.BookDetailResponse, error) {
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

	if err := os.MkdirAll(pdfDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create pdf directory: %w", err)
	}
	pdfPath := filepath.Join(pdfDir, pdfFileName)
	if err := os.WriteFile(pdfPath, pdfFileBytes, 0644); err != nil {
		return nil, fmt.Errorf("failed to save pdf: %w", err)
	}

	if err := os.MkdirAll(thumbnailDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create thumbnail directory: %w", err)
	}
	thumbnailPath := filepath.Join(thumbnailDir, thumbFileName)

	doc, err := fitz.New(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open pdf for thumbnail: %w", err)
	}
	defer doc.Close()

	img, err := doc.Image(0)
	if err != nil {
		return nil, fmt.Errorf("failed to render pdf cover: %w", err)
	}

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
		Stock:       payload.Stock,
		Description: payload.Description,
		Publisher:   payload.Publisher,
		ISBN:        payload.ISBN,
		Language:    payload.Language,
		Pages:       payload.Pages,
		Thumbnail:   thumbURL,
		FilePath:    pdfURL,
	}
	err = b.bookRepository.InsertBook(book)
	if err != nil {
		return nil, err
	}

	response := &dto.BookDetailResponse{
		StatusCode: http.StatusCreated,
		Message:    "Buku berhasil diupload",
		Data: dto.BookData{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Year:        book.Year,
			Genre:       book.Genre,
			Stock:       book.Stock,
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

func (b *BookService) SearchBooks(keyword string) (*dto.BookListResponse, error) {
	books, err := b.bookRepository.SearchBooks(keyword)
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
